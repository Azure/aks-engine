// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"fmt"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// TemplateGenerator represents the object that performs the template generation.
type TemplateGenerator struct {
	Translator *i18n.Translator
}

// InitializeTemplateGenerator creates a new template generator object
func InitializeTemplateGenerator(ctx Context) (*TemplateGenerator, error) {
	t := &TemplateGenerator{
		Translator: ctx.Translator,
	}

	if err := t.verifyFiles(); err != nil {
		return nil, err
	}

	return t, nil
}

// GenerateTemplate generates the template from the API Model
func (t *TemplateGenerator) GenerateTemplate(containerService *api.ContainerService, generatorCode string, acsengineVersion string) (templateRaw string, parametersRaw string, err error) {
	// named return values are used in order to set err in case of a panic
	templateRaw = ""
	parametersRaw = ""
	err = nil

	var templ *template.Template

	properties := containerService.Properties
	// save the current orchestrator version and restore it after deploying.
	// this allows us to deploy agents on the most recent patch without updating the orchestator version in the object
	orchVersion := properties.OrchestratorProfile.OrchestratorVersion
	defer func() {
		properties.OrchestratorProfile.OrchestratorVersion = orchVersion
	}()

	templ = template.New("acs template").Funcs(t.getTemplateFuncMap(containerService))

	files, baseFile, e := t.prepareTemplateFiles(properties)
	if e != nil {
		return "", "", e
	}

	for _, file := range files {
		bytes, e := Asset(file)
		if e != nil {
			err = t.Translator.Errorf("Error reading file %s, Error: %s", file, e.Error())
			return templateRaw, parametersRaw, err
		}
		if _, err = templ.New(file).Parse(string(bytes)); err != nil {
			return templateRaw, parametersRaw, err
		}
	}
	// template generation may have panics in the called functions.  This catches those panics
	// and ensures the panic is returned as an error
	defer func() {
		if r := recover(); r != nil {
			s := debug.Stack()
			err = errors.Errorf("%v - %s", r, s)

			// invalidate the template and the parameters
			templateRaw = ""
			parametersRaw = ""
		}
	}()

	if !validateDistro(containerService) {
		return templateRaw, parametersRaw, errors.New("Invalid distro")
	}

	var b bytes.Buffer
	if err = templ.ExecuteTemplate(&b, baseFile, properties); err != nil {
		return templateRaw, parametersRaw, err
	}
	templateRaw = b.String()

	var parametersMap paramsMap
	if parametersMap, err = getParameters(containerService, generatorCode, acsengineVersion); err != nil {
		return templateRaw, parametersRaw, err
	}

	var parameterBytes []byte
	if parameterBytes, err = helpers.JSONMarshal(parametersMap, false); err != nil {
		return templateRaw, parametersRaw, err
	}
	parametersRaw = string(parameterBytes)

	return templateRaw, parametersRaw, err
}

func (t *TemplateGenerator) verifyFiles() error {
	allFiles := commonTemplateFiles
	allFiles = append(allFiles, dcosTemplateFiles...)
	allFiles = append(allFiles, dcos2TemplateFiles...)
	allFiles = append(allFiles, kubernetesTemplateFiles...)
	allFiles = append(allFiles, swarmTemplateFiles...)
	for _, file := range allFiles {
		if _, err := Asset(file); err != nil {
			return t.Translator.Errorf("template file %s does not exist", file)
		}
	}
	return nil
}

func (t *TemplateGenerator) prepareTemplateFiles(properties *api.Properties) ([]string, string, error) {
	var files []string
	var baseFile string
	switch properties.OrchestratorProfile.OrchestratorType {
	case api.DCOS:
		if properties.OrchestratorProfile.DcosConfig == nil || properties.OrchestratorProfile.DcosConfig.BootstrapProfile == nil {
			files = append(commonTemplateFiles, dcosTemplateFiles...)
			baseFile = dcosBaseFile
		} else {
			files = append(commonTemplateFiles, dcos2TemplateFiles...)
			baseFile = dcos2BaseFile
		}
	case api.Swarm:
		files = append(commonTemplateFiles, swarmTemplateFiles...)
		baseFile = swarmBaseFile
	case api.Kubernetes:
		files = append(commonTemplateFiles, kubernetesTemplateFiles...)
		baseFile = kubernetesBaseFile
	case api.SwarmMode:
		files = append(commonTemplateFiles, swarmModeTemplateFiles...)
		baseFile = swarmBaseFile
	default:
		return nil, "", t.Translator.Errorf("orchestrator '%s' is unsupported", properties.OrchestratorProfile.OrchestratorType)
	}

	return files, baseFile, nil
}

func (t *TemplateGenerator) getMasterCustomData(cs *api.ContainerService, textFilename string, profile *api.Properties) string {
	str, e := t.getSingleLineForTemplate(textFilename, cs, profile)
	if e != nil {
		panic(e)
	}

	// add manifests
	str = substituteConfigString(str,
		kubernetesManifestSettingsInit(profile),
		"k8s/manifests",
		"/etc/kubernetes/manifests",
		"MASTER_MANIFESTS_CONFIG_PLACEHOLDER",
		profile.OrchestratorProfile.OrchestratorVersion)

	// add artifacts
	str = substituteConfigString(str,
		kubernetesArtifactSettingsInitMaster(profile),
		"k8s/artifacts",
		"/etc/systemd/system",
		"MASTER_ARTIFACTS_CONFIG_PLACEHOLDER",
		profile.OrchestratorProfile.OrchestratorVersion)

	// add addons
	str = substituteConfigString(str,
		kubernetesAddonSettingsInit(profile),
		"k8s/addons",
		"/etc/kubernetes/addons",
		"MASTER_ADDONS_CONFIG_PLACEHOLDER",
		profile.OrchestratorProfile.OrchestratorVersion)

	// add custom files
	customFilesReader, err := customfilesIntoReaders(masterCustomFiles(profile))
	if err != nil {
		log.Fatalf("Could not read custom files: %s", err.Error())
	}
	str = substituteConfigStringCustomFiles(str,
		customFilesReader,
		"MASTER_CUSTOM_FILES_PLACEHOLDER")

	addonStr := getContainerAddonsString(cs.Properties, "k8s/containeraddons")

	str = strings.Replace(str, "MASTER_CONTAINER_ADDONS_PLACEHOLDER", addonStr, -1)

	// return the custom data
	return fmt.Sprintf("\"customData\": \"[base64(concat('%s'))]\",", str)
}

// getTemplateFuncMap returns all functions used in template generation
func (t *TemplateGenerator) getTemplateFuncMap(cs *api.ContainerService) template.FuncMap {
	return template.FuncMap{
		"IsMasterVirtualMachineScaleSets": func() bool {
			return cs.Properties.MasterProfile != nil && cs.Properties.MasterProfile.IsVirtualMachineScaleSets()
		},
		"IsHostedMaster": func() bool {
			return cs.Properties.IsHostedMasterProfile()
		},
		"IsDCOS19": func() bool {
			return cs.Properties.OrchestratorProfile.OrchestratorType == api.DCOS &&
				(cs.Properties.OrchestratorProfile.OrchestratorVersion == common.DCOSVersion1Dot9Dot0 ||
					cs.Properties.OrchestratorProfile.OrchestratorVersion == common.DCOSVersion1Dot9Dot8)
		},
		"IsKubernetesVersionGe": func(version string) bool {
			return cs.Properties.OrchestratorProfile.IsKubernetes() && common.IsKubernetesVersionGe(cs.Properties.OrchestratorProfile.OrchestratorVersion, version)
		},
		"IsKubernetesVersionLt": func(version string) bool {
			return cs.Properties.OrchestratorProfile.IsKubernetes() && !common.IsKubernetesVersionGe(cs.Properties.OrchestratorProfile.OrchestratorVersion, version)
		},
		"GetMasterKubernetesLabels": func(rg string) string {
			var buf bytes.Buffer
			buf.WriteString("kubernetes.io/role=master")
			buf.WriteString(",node-role.kubernetes.io/master=")
			buf.WriteString(fmt.Sprintf(",kubernetes.azure.com/cluster=%s", rg))
			return buf.String()
		},
		"GetAgentKubernetesLabels": func(profile *api.AgentPoolProfile, rg string) string {
			var buf bytes.Buffer
			buf.WriteString("node-role.kubernetes.io/agent=")
			buf.WriteString(fmt.Sprintf(",kubernetes.io/role=agent,agentpool=%s", profile.Name))
			if profile.StorageProfile == api.ManagedDisks {
				storagetier, _ := getStorageAccountType(profile.VMSize)
				buf.WriteString(fmt.Sprintf(",storageprofile=managed,storagetier=%s", storagetier))
			}
			if common.IsNvidiaEnabledSKU(profile.VMSize) {
				accelerator := "nvidia"
				buf.WriteString(fmt.Sprintf(",accelerator=%s", accelerator))
			}
			buf.WriteString(fmt.Sprintf(",kubernetes.azure.com/cluster=%s", rg))
			for k, v := range profile.CustomNodeLabels {
				buf.WriteString(fmt.Sprintf(",%s=%s", k, v))
			}
			return buf.String()
		},
		"GetKubeletConfigKeyVals": func(kc *api.KubernetesConfig) string {
			if kc == nil {
				return ""
			}
			kubeletConfig := cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
			if kc.KubeletConfig != nil {
				kubeletConfig = kc.KubeletConfig
			}
			// Order by key for consistency
			keys := []string{}
			for key := range kubeletConfig {
				keys = append(keys, key)
			}
			sort.Strings(keys)
			var buf bytes.Buffer
			for _, key := range keys {
				buf.WriteString(fmt.Sprintf("%s=%s ", key, kubeletConfig[key]))
			}
			return buf.String()
		},
		"GetKubeletConfigKeyValsPsh": func(kc *api.KubernetesConfig) string {
			if kc == nil {
				return ""
			}
			kubeletConfig := cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
			if kc.KubeletConfig != nil {
				kubeletConfig = kc.KubeletConfig
			}
			// Order by key for consistency
			keys := []string{}
			for key := range kubeletConfig {
				keys = append(keys, key)
			}
			sort.Strings(keys)
			var buf bytes.Buffer
			for _, key := range keys {
				buf.WriteString(fmt.Sprintf("\"%s=%s\", ", key, kubeletConfig[key]))
			}
			return strings.TrimSuffix(buf.String(), ", ")
		},
		"GetK8sRuntimeConfigKeyVals": func(config map[string]string) string {
			// Order by key for consistency
			keys := []string{}
			for key := range config {
				keys = append(keys, key)
			}
			sort.Strings(keys)
			var buf bytes.Buffer
			for _, key := range keys {
				buf.WriteString(fmt.Sprintf("\\\"%s=%s\\\", ", key, config[key]))
			}
			return strings.TrimSuffix(buf.String(), ", ")
		},
		"HasPrivateRegistry": func() bool {
			if cs.Properties.OrchestratorProfile.DcosConfig != nil {
				return len(cs.Properties.OrchestratorProfile.DcosConfig.Registry) > 0
			}
			return false
		},
		"RequiresFakeAgentOutput": func() bool {
			return cs.Properties.OrchestratorProfile.IsKubernetes()
		},
		"IsSwarmMode": func() bool {
			return cs.Properties.OrchestratorProfile.IsSwarmMode()
		},
		"IsKubernetes": func() bool {
			return cs.Properties.OrchestratorProfile.IsKubernetes()
		},
		"IsPublic": func(ports []int) bool {
			return len(ports) > 0
		},
		"IsAzureCNI": func() bool {
			return cs.Properties.OrchestratorProfile.IsAzureCNI()
		},
		"RequireRouteTable": func() bool {
			return cs.Properties.OrchestratorProfile.RequireRouteTable()
		},

		"IsPrivateCluster": func() bool {
			if !cs.Properties.OrchestratorProfile.IsKubernetes() {
				return false
			}
			return helpers.IsTrueBoolPointer(cs.Properties.OrchestratorProfile.KubernetesConfig.PrivateCluster.Enabled)
		},
		"ProvisionJumpbox": func() bool {
			return cs.Properties.OrchestratorProfile.KubernetesConfig.PrivateJumpboxProvision()
		},
		"JumpboxIsManagedDisks": func() bool {
			if cs.Properties.OrchestratorProfile.KubernetesConfig.PrivateJumpboxProvision() && cs.Properties.OrchestratorProfile.KubernetesConfig.PrivateCluster.JumpboxProfile.StorageProfile == api.ManagedDisks {
				return true
			}
			return false
		},
		"GetKubeConfig": func() string {
			kubeConfig, err := GenerateKubeConfig(cs.Properties, cs.Location)
			if err != nil {
				panic(err)
			}
			return escapeSingleLine(kubeConfig)
		},
		"UseManagedIdentity": func() bool {
			return cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity
		},
		"UserAssignedIDEnabled": func() bool {
			if cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity &&
				cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedID != "" {
				return true
			}
			return false
		},
		"UserAssignedID": func() string {
			if cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity &&
				cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedID != "" {
				return cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedID
			}
			return ""
		},
		"UserAssignedClientID": func() string {
			if cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity &&
				cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedClientID != "" {
				return cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedClientID
			}
			return ""
		},
		"UseAksExtension": func() bool {
			cloudSpecConfig := cs.GetCloudSpecConfig()
			return cloudSpecConfig.CloudName == api.AzurePublicCloud
		},
		"IsMooncake": func() bool {
			cloudSpecConfig := cs.GetCloudSpecConfig()
			return cloudSpecConfig.CloudName == api.AzureChinaCloud
		},
		"UseInstanceMetadata": func() bool {
			return helpers.IsTrueBoolPointer(cs.Properties.OrchestratorProfile.KubernetesConfig.UseInstanceMetadata)
		},
		"NeedsKubeDNSWithExecHealthz": func() bool {
			return cs.Properties.OrchestratorProfile.NeedsExecHealthz()
		},
		"LoadBalancerSku": func() string {
			return cs.Properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku
		},
		"ExcludeMasterFromStandardLB": func() bool {
			return helpers.IsTrueBoolPointer(cs.Properties.OrchestratorProfile.KubernetesConfig.ExcludeMasterFromStandardLB)
		},
		"GetVNETSubnetDependencies": func() string {
			return getVNETSubnetDependencies(cs.Properties)
		},
		"GetLBRules": func(name string, ports []int) string {
			return getLBRules(name, ports)
		},
		"GetProbes": func(ports []int) string {
			return getProbes(ports)
		},
		"GetSecurityRules": func(ports []int) string {
			return getSecurityRules(ports)
		},
		"GetUniqueNameSuffix": func() string {
			return cs.Properties.GetClusterID()
		},
		"GetVNETAddressPrefixes": func() string {
			return getVNETAddressPrefixes(cs.Properties)
		},
		"GetVNETSubnets": func(addNSG bool) string {
			return getVNETSubnets(cs.Properties, addNSG)
		},
		"GetDataDisks": func(profile *api.AgentPoolProfile) string {
			return getDataDisks(profile)
		},
		"HasBootstrap": func() bool {
			return cs.Properties.OrchestratorProfile.DcosConfig != nil && cs.Properties.OrchestratorProfile.DcosConfig.BootstrapProfile != nil
		},
		"HasBootstrapPublicIP": func() bool {
			return false
		},
		"IsHostedBootstrap": func() bool {
			return false
		},
		"IsFeatureEnabled": func(feature string) bool {
			return cs.Properties.FeatureFlags.IsFeatureEnabled(feature)
		},
		"GetDCOSBootstrapCustomData": func() string {
			masterIPList := generateIPList(cs.Properties.MasterProfile.Count, cs.Properties.MasterProfile.FirstConsecutiveStaticIP)
			for i, v := range masterIPList {
				masterIPList[i] = "    - " + v
			}

			str := getSingleLineDCOSCustomData(
				cs.Properties.OrchestratorProfile.OrchestratorType,
				dcos2BootstrapCustomdata, 0,
				map[string]string{
					"PROVISION_SOURCE_STR":    getDCOSProvisionScript(dcosProvisionSource),
					"PROVISION_STR":           getDCOSProvisionScript(dcos2BootstrapProvision),
					"MASTER_IP_LIST":          strings.Join(masterIPList, "\n"),
					"BOOTSTRAP_IP":            cs.Properties.OrchestratorProfile.DcosConfig.BootstrapProfile.StaticIP,
					"BOOTSTRAP_OAUTH_ENABLED": strconv.FormatBool(cs.Properties.OrchestratorProfile.DcosConfig.BootstrapProfile.OAuthEnabled)})

			return fmt.Sprintf("\"customData\": \"[base64(concat('#cloud-config\\n\\n', '%s'))]\",", str)
		},
		"GetDCOSMasterCustomData": func() string {
			masterAttributeContents := getDCOSMasterCustomNodeLabels()
			masterPreprovisionExtension := ""
			if cs.Properties.MasterProfile.PreprovisionExtension != nil {
				masterPreprovisionExtension += "\n"
				masterPreprovisionExtension += makeMasterExtensionScriptCommands(cs)
			}
			var bootstrapIP string
			if cs.Properties.OrchestratorProfile.DcosConfig != nil && cs.Properties.OrchestratorProfile.DcosConfig.BootstrapProfile != nil {
				bootstrapIP = cs.Properties.OrchestratorProfile.DcosConfig.BootstrapProfile.StaticIP
			}

			str := getSingleLineDCOSCustomData(
				cs.Properties.OrchestratorProfile.OrchestratorType,
				getDCOSCustomDataTemplate(cs.Properties.OrchestratorProfile.OrchestratorType, cs.Properties.OrchestratorProfile.OrchestratorVersion),
				cs.Properties.MasterProfile.Count,
				map[string]string{
					"PROVISION_SOURCE_STR":   getDCOSProvisionScript(dcosProvisionSource),
					"PROVISION_STR":          getDCOSMasterProvisionScript(cs.Properties.OrchestratorProfile, bootstrapIP),
					"ATTRIBUTES_STR":         masterAttributeContents,
					"PREPROVISION_EXTENSION": masterPreprovisionExtension,
					"ROLENAME":               "master"})

			return fmt.Sprintf("\"customData\": \"[base64(concat('#cloud-config\\n\\n', '%s'))]\",", str)
		},
		"GetDCOSAgentCustomData": func(profile *api.AgentPoolProfile) string {
			attributeContents := getDCOSAgentCustomNodeLabels(profile)
			agentPreprovisionExtension := ""
			if profile.PreprovisionExtension != nil {
				agentPreprovisionExtension += "\n"
				agentPreprovisionExtension += makeAgentExtensionScriptCommands(cs, profile)
			}
			var agentRoleName, bootstrapIP string
			if len(profile.Ports) > 0 {
				agentRoleName = "slave_public"
			} else {
				agentRoleName = "slave"
			}
			if cs.Properties.OrchestratorProfile.DcosConfig != nil && cs.Properties.OrchestratorProfile.DcosConfig.BootstrapProfile != nil {
				bootstrapIP = cs.Properties.OrchestratorProfile.DcosConfig.BootstrapProfile.StaticIP
			}

			str := getSingleLineDCOSCustomData(
				cs.Properties.OrchestratorProfile.OrchestratorType,
				getDCOSCustomDataTemplate(cs.Properties.OrchestratorProfile.OrchestratorType, cs.Properties.OrchestratorProfile.OrchestratorVersion),
				cs.Properties.MasterProfile.Count,
				map[string]string{
					"PROVISION_SOURCE_STR":   getDCOSProvisionScript(dcosProvisionSource),
					"PROVISION_STR":          getDCOSAgentProvisionScript(profile, cs.Properties.OrchestratorProfile, bootstrapIP),
					"ATTRIBUTES_STR":         attributeContents,
					"PREPROVISION_EXTENSION": agentPreprovisionExtension,
					"ROLENAME":               agentRoleName})

			return fmt.Sprintf("\"customData\": \"[base64(concat('#cloud-config\\n\\n', '%s'))]\",", str)
		},
		"GetDCOSWindowsAgentCustomData": func(profile *api.AgentPoolProfile) string {
			agentPreprovisionExtension := ""
			if profile.PreprovisionExtension != nil {
				agentPreprovisionExtension += "\n"
				agentPreprovisionExtension += makeAgentExtensionScriptCommands(cs, profile)
			}
			b, err := Asset(dcosWindowsProvision)
			if err != nil {
				// this should never happen and this is a bug
				panic(fmt.Sprintf("BUG: %s", err.Error()))
			}
			// translate the parameters
			csStr := string(b)
			csStr = strings.Replace(csStr, "PREPROVISION_EXTENSION", agentPreprovisionExtension, -1)
			csStr = strings.Replace(csStr, "\r\n", "\n", -1)
			str := getBase64CustomScriptFromStr(csStr)
			return fmt.Sprintf("\"customData\": \"%s\"", str)
		},
		"GetDCOSWindowsAgentCustomNodeAttributes": func(profile *api.AgentPoolProfile) string {
			return getDCOSWindowsAgentCustomAttributes(profile)
		},
		"GetDCOSWindowsAgentPreprovisionParameters": func(profile *api.AgentPoolProfile) string {
			agentPreprovisionExtensionParameters := ""
			if profile.PreprovisionExtension != nil {
				agentPreprovisionExtensionParameters = getDCOSWindowsAgentPreprovisionParameters(cs, profile)
			}
			return agentPreprovisionExtensionParameters
		},
		"GetMasterAllowedSizes": func() string {
			if cs.Properties.OrchestratorProfile.OrchestratorType == api.DCOS {
				return helpers.GetDCOSMasterAllowedSizes()
			}
			return helpers.GetKubernetesAllowedSizes()
		},
		"GetDefaultVNETCIDR": func() string {
			return DefaultVNETCIDR
		},
		"GetAgentAllowedSizes": func() string {
			return helpers.GetKubernetesAllowedSizes()
		},
		"getSwarmVersions": func() string {
			return getSwarmVersions(api.SwarmVersion, api.SwarmDockerComposeVersion)
		},
		"GetSwarmModeVersions": func() string {
			return getSwarmVersions(api.DockerCEVersion, api.DockerCEDockerComposeVersion)
		},
		"GetSizeMap": func() string {
			return helpers.GetSizeMap()
		},
		"Base64": func(s string) string {
			return base64.StdEncoding.EncodeToString([]byte(s))
		},
		"GetDefaultInternalLbStaticIPOffset": func() int {
			return DefaultInternalLbStaticIPOffset
		},
		"GetKubernetesMasterCustomData": func(profile *api.Properties) string {
			str := t.getMasterCustomData(cs, kubernetesMasterCustomDataYaml, profile)
			return str
		},
		"GetKubernetesAgentCustomData": func(profile *api.AgentPoolProfile) string {
			str, e := t.getSingleLineForTemplate(kubernetesAgentCustomDataYaml, cs, profile)

			if e != nil {
				panic(e)
			}

			// add artifacts
			str = substituteConfigString(str,
				kubernetesArtifactSettingsInitAgent(cs.Properties),
				"k8s/artifacts",
				"/etc/systemd/system",
				"AGENT_ARTIFACTS_CONFIG_PLACEHOLDER",
				cs.Properties.OrchestratorProfile.OrchestratorVersion)

			return fmt.Sprintf("\"customData\": \"[base64(concat('%s'))]\",", str)
		},
		"GetKubernetesJumpboxCustomData": func(p *api.Properties) string {
			str, err := t.getSingleLineForTemplate(kubernetesJumpboxCustomDataYaml, cs, p)

			if err != nil {
				panic(err)
			}

			return fmt.Sprintf("\"customData\": \"[base64(concat('%s'))]\",", str)
		},
		"WriteLinkedTemplatesForExtensions": func() string {
			extensions := getLinkedTemplatesForExtensions(cs.Properties)
			return extensions
		},
		"GetKubernetesB64Provision": func() string {
			return getBase64CustomScript(kubernetesCustomScript)
		},
		"GetKubernetesB64ProvisionSource": func() string {
			return getBase64CustomScript(kubernetesProvisionSourceScript)
		},
		"GetKubernetesB64HealthMonitorScript": func() string {
			return getBase64CustomScript(kubernetesHealthMonitorScript)
		},
		"GetKubernetesB64Installs": func() string {
			return getBase64CustomScript(kubernetesInstalls)
		},
		"GetKubernetesB64Configs": func() string {
			return getBase64CustomScript(kubernetesConfigurations)
		},
		"GetKubernetesB64Mountetcd": func() string {
			return getBase64CustomScript(kubernetesMountetcd)
		},
		"GetKubernetesB64CustomSearchDomainsScript": func() string {
			return getBase64CustomScript(kubernetesCustomSearchDomainsScript)
		},
		"GetKubernetesB64GenerateProxyCerts": func() string {
			return getBase64CustomScript(kubernetesMasterGenerateProxyCertsScript)
		},
		"GetB64sshdConfig": func() string {
			return getBase64CustomScript(sshdConfig)
		},
		"GetB64systemConf": func() string {
			return getBase64CustomScript(systemConf)
		},
		"GetKubernetesMasterPreprovisionYaml": func() string {
			str := ""
			if cs.Properties.MasterProfile.PreprovisionExtension != nil {
				str += "\n"
				str += makeMasterExtensionScriptCommands(cs)
			}
			return str
		},
		"GetKubernetesAgentPreprovisionYaml": func(profile *api.AgentPoolProfile) string {
			str := ""
			if profile.PreprovisionExtension != nil {
				str += "\n"
				str += makeAgentExtensionScriptCommands(cs, profile)
			}
			return str
		},
		"GetMasterSwarmCustomData": func() string {
			files := []string{swarmProvision}
			str := buildYamlFileWithWriteFiles(files)
			if cs.Properties.MasterProfile.PreprovisionExtension != nil {
				extensionStr := makeMasterExtensionScriptCommands(cs)
				str += "'runcmd:\n" + extensionStr + "\n\n'"
			}
			str = escapeSingleLine(str)
			return fmt.Sprintf("\"customData\": \"[base64(concat('%s'))]\",", str)
		},
		"GetAgentSwarmCustomData": func(profile *api.AgentPoolProfile) string {
			files := []string{swarmProvision}
			str := buildYamlFileWithWriteFiles(files)
			str = escapeSingleLine(str)
			return fmt.Sprintf("\"customData\": \"[base64(concat('%s',variables('%sRunCmdFile'),variables('%sRunCmd')))]\",", str, profile.Name, profile.Name)
		},
		"GetSwarmAgentPreprovisionExtensionCommands": func(profile *api.AgentPoolProfile) string {
			str := ""
			if profile.PreprovisionExtension != nil {
				makeAgentExtensionScriptCommands(cs, profile)
			}
			str = escapeSingleLine(str)
			return str
		},
		"GetLocation": func() string {
			return cs.Location
		},
		"GetWinAgentSwarmCustomData": func() string {
			str := getBase64CustomScript(swarmWindowsProvision)
			return fmt.Sprintf("\"customData\": \"%s\"", str)
		},
		"GetWinAgentSwarmModeCustomData": func() string {
			str := getBase64CustomScript(swarmModeWindowsProvision)
			return fmt.Sprintf("\"customData\": \"%s\"", str)
		},
		"GetKubernetesWindowsAgentFunctions": func() string {
			// Collect all the parts into a zip
			var parts = []string{
				kubernetesWindowsAgentFunctionsPS1,
				kubernetesWindowsConfigFunctionsPS1,
				kubernetesWindowsKubeletFunctionsPS1,
				kubernetesWindowsCniFunctionsPS1,
				kubernetesWindowsAzureCniFunctionsPS1}

			// Create a buffer, new zip
			buf := new(bytes.Buffer)
			zw := zip.NewWriter(buf)

			for _, part := range parts {
				f, err := zw.Create(part)
				if err != nil {
					panic(err)
				}
				partContents, err := Asset(part)
				if err != nil {
					panic(err)
				}
				_, err = f.Write([]byte(partContents))
				if err != nil {
					panic(err)
				}
			}
			err := zw.Close()
			if err != nil {
				panic(err)
			}
			return base64.StdEncoding.EncodeToString(buf.Bytes())
		},
		"GetKubernetesWindowsAgentCustomData": func(profile *api.AgentPoolProfile) string {
			str, e := t.getSingleLineForTemplate(kubernetesWindowsAgentCustomDataPS1, cs, profile)

			if e != nil {
				panic(e)
			}

			preprovisionCmd := ""

			if profile.PreprovisionExtension != nil {
				preprovisionCmd = makeAgentExtensionScriptCommands(cs, profile)
			}

			str = strings.Replace(str, "PREPROVISION_EXTENSION", escapeSingleLine(strings.TrimSpace(preprovisionCmd)), -1)

			return fmt.Sprintf("\"customData\": \"[base64(concat('%s'))]\",", str)
		},
		"GetMasterSwarmModeCustomData": func() string {
			files := []string{swarmModeProvision}
			str := buildYamlFileWithWriteFiles(files)
			if cs.Properties.MasterProfile.PreprovisionExtension != nil {
				extensionStr := makeMasterExtensionScriptCommands(cs)
				str += "runcmd:\n" + extensionStr + "\n\n"
			}
			str = escapeSingleLine(str)
			return fmt.Sprintf("\"customData\": \"[base64(concat('%s'))]\",", str)
		},
		"GetAgentSwarmModeCustomData": func(profile *api.AgentPoolProfile) string {
			files := []string{swarmModeProvision}
			str := buildYamlFileWithWriteFiles(files)
			str = escapeSingleLine(str)
			return fmt.Sprintf("\"customData\": \"[base64(concat('%s',variables('%sRunCmdFile'),variables('%sRunCmd')))]\",", str, profile.Name, profile.Name)
		},
		"GetKubernetesSubnets": func() string {
			return getKubernetesSubnets(cs.Properties)
		},
		"GetKubernetesPodStartIndex": func() string {
			return fmt.Sprintf("%d", getKubernetesPodStartIndex(cs.Properties))
		},
		"WrapAsVariable": func(s string) string {
			return fmt.Sprintf("',variables('%s'),'", s)
		},
		"WrapAsParameter": func(s string) string {
			return fmt.Sprintf("',parameters('%s'),'", s)
		},
		"WrapAsParameterObject": func(o, p string) string {
			return fmt.Sprintf("',parameters('%s').%s,'", o, p)
		},
		"WrapAsVerbatim": func(s string) string {
			return fmt.Sprintf("',%s,'", s)
		},
		"AnyAgentUsesAvailabilitySets": func() bool {
			for _, agentProfile := range cs.Properties.AgentPoolProfiles {
				if agentProfile.IsAvailabilitySets() {
					return true
				}
			}
			return false
		},
		"AnyAgentUsesVirtualMachineScaleSets": func() bool {
			for _, agentProfile := range cs.Properties.AgentPoolProfiles {
				if agentProfile.IsVirtualMachineScaleSets() {
					return true
				}
			}
			return false
		},
		"HasLinuxAgents": func() bool {
			for _, agentProfile := range cs.Properties.AgentPoolProfiles {
				if agentProfile.IsLinux() {
					return true
				}
			}
			return false
		},
		"IsNVIDIADevicePluginEnabled": func() bool {
			return cs.Properties.IsNVIDIADevicePluginEnabled()
		},
		"IsNSeriesSKU": func(profile *api.AgentPoolProfile) bool {
			return common.IsNvidiaEnabledSKU(profile.VMSize)
		},
		"UseSinglePlacementGroup": func(profile *api.AgentPoolProfile) bool {
			return *profile.SinglePlacementGroup
		},
		"HasAvailabilityZones": func(profile *api.AgentPoolProfile) bool {
			return profile.HasAvailabilityZones()
		},
		"HasLinuxSecrets": func() bool {
			return cs.Properties.LinuxProfile.HasSecrets()
		},
		"HasCustomSearchDomain": func() bool {
			return cs.Properties.LinuxProfile.HasSearchDomain()
		},
		"HasCustomNodesDNS": func() bool {
			return cs.Properties.LinuxProfile.HasCustomNodesDNS()
		},
		"HasWindowsSecrets": func() bool {
			return cs.Properties.WindowsProfile.HasSecrets()
		},
		"HasWindowsCustomImage": func() bool {
			return cs.Properties.WindowsProfile.HasCustomImage()
		},
		"GetConfigurationScriptRootURL": func() string {
			if cs.Properties.LinuxProfile.ScriptRootURL == "" {
				return DefaultConfigurationScriptRootURL
			}
			return cs.Properties.LinuxProfile.ScriptRootURL
		},
		"GetMasterOSImageOffer": func() string {
			cloudSpecConfig := cs.GetCloudSpecConfig()
			return fmt.Sprintf("\"%s\"", cloudSpecConfig.OSImageConfig[cs.Properties.MasterProfile.Distro].ImageOffer)
		},
		"GetMasterOSImagePublisher": func() string {
			cloudSpecConfig := cs.GetCloudSpecConfig()
			return fmt.Sprintf("\"%s\"", cloudSpecConfig.OSImageConfig[cs.Properties.MasterProfile.Distro].ImagePublisher)
		},
		"GetMasterOSImageSKU": func() string {
			cloudSpecConfig := cs.GetCloudSpecConfig()
			return fmt.Sprintf("\"%s\"", cloudSpecConfig.OSImageConfig[cs.Properties.MasterProfile.Distro].ImageSku)
		},
		"GetMasterOSImageVersion": func() string {
			cloudSpecConfig := cs.GetCloudSpecConfig()
			return fmt.Sprintf("\"%s\"", cloudSpecConfig.OSImageConfig[cs.Properties.MasterProfile.Distro].ImageVersion)
		},
		"GetAgentOSImageOffer": func(profile *api.AgentPoolProfile) string {
			cloudSpecConfig := cs.GetCloudSpecConfig()
			return fmt.Sprintf("\"%s\"", cloudSpecConfig.OSImageConfig[profile.Distro].ImageOffer)
		},
		"GetAgentOSImagePublisher": func(profile *api.AgentPoolProfile) string {
			cloudSpecConfig := cs.GetCloudSpecConfig()
			return fmt.Sprintf("\"%s\"", cloudSpecConfig.OSImageConfig[profile.Distro].ImagePublisher)
		},
		"GetAgentOSImageSKU": func(profile *api.AgentPoolProfile) string {
			cloudSpecConfig := cs.GetCloudSpecConfig()
			return fmt.Sprintf("\"%s\"", cloudSpecConfig.OSImageConfig[profile.Distro].ImageSku)
		},
		"GetAgentOSImageVersion": func(profile *api.AgentPoolProfile) string {
			cloudSpecConfig := cs.GetCloudSpecConfig()
			return fmt.Sprintf("\"%s\"", cloudSpecConfig.OSImageConfig[profile.Distro].ImageVersion)
		},
		"UseAgentCustomImage": func(profile *api.AgentPoolProfile) bool {
			imageRef := profile.ImageRef
			return imageRef != nil && len(imageRef.Name) > 0 && len(imageRef.ResourceGroup) > 0
		},
		"UseMasterCustomImage": func() bool {
			imageRef := cs.Properties.MasterProfile.ImageRef
			return imageRef != nil && len(imageRef.Name) > 0 && len(imageRef.ResourceGroup) > 0
		},
		"GetAgentVMPrefix": func(profile *api.AgentPoolProfile) string {
			return cs.Properties.GetAgentVMPrefix(profile)
		},
		"GetMasterVMPrefix": func() string {
			return cs.Properties.GetMasterVMPrefix()
		},
		"GetRouteTableName": func() string {
			return cs.Properties.GetRouteTableName()
		},
		"GetNSGName": func() string {
			return cs.Properties.GetNSGName()
		},
		"GetMasterEtcdServerPort": func() int {
			return DefaultMasterEtcdServerPort
		},
		"GetMasterEtcdClientPort": func() int {
			return DefaultMasterEtcdClientPort
		},
		"GetPrimaryAvailabilitySetName": func() string {
			return cs.Properties.GetPrimaryAvailabilitySetName()
		},
		"GetPrimaryScaleSetName": func() string {
			return cs.Properties.GetPrimaryScaleSetName()
		},
		"UseCloudControllerManager": func() bool {
			return cs.Properties.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager != nil && *cs.Properties.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager
		},
		"AdminGroupID": func() bool {
			return cs.Properties.AADProfile != nil && cs.Properties.AADProfile.AdminGroupID != ""
		},
		"EnableDataEncryptionAtRest": func() bool {
			return helpers.IsTrueBoolPointer(cs.Properties.OrchestratorProfile.KubernetesConfig.EnableDataEncryptionAtRest)
		},
		"EnableEncryptionWithExternalKms": func() bool {
			return helpers.IsTrueBoolPointer(cs.Properties.OrchestratorProfile.KubernetesConfig.EnableEncryptionWithExternalKms)
		},
		"EnableAggregatedAPIs": func() bool {
			if cs.Properties.OrchestratorProfile.KubernetesConfig.EnableAggregatedAPIs {
				return true
			} else if common.IsKubernetesVersionGe(cs.Properties.OrchestratorProfile.OrchestratorVersion, "1.9.0") {
				return true
			}
			return false
		},
		"EnablePodSecurityPolicy": func() bool {
			return helpers.IsTrueBoolPointer(cs.Properties.OrchestratorProfile.KubernetesConfig.EnablePodSecurityPolicy)
		},
		// inspired by http://stackoverflow.com/questions/18276173/calling-a-template-with-several-pipeline-parameters/18276968#18276968
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values)%2 != 0 {
				return nil, errors.New("invalid dict call")
			}
			dict := make(map[string]interface{}, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, errors.New("dict keys must be strings")
				}
				dict[key] = values[i+1]
			}
			return dict, nil
		},
		"loop": func(min, max int) []int {
			var s []int
			for i := min; i <= max; i++ {
				s = append(s, i)
			}
			return s
		},
		"subtract": func(a, b int) int {
			return a - b
		},
		"IsCustomVNET": func() bool {
			return cs.Properties.AreAgentProfilesCustomVNET()
		},
		"quote":      strconv.Quote,
		"shellQuote": helpers.ShellQuote,
	}
}
