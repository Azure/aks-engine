// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"runtime/debug"
	"sort"
	"strings"
	"text/template"

	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/Azure/aks-engine/pkg/telemetry"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type ARMTemplate struct {
	Schema         string      `json:"$schema,omitempty"`
	ContentVersion string      `json:"contentVersion,omitempty"`
	Parameters     interface{} `json:"parameters,omitempty"`
	Variables      interface{} `json:"variables,omitempty"`
	Resources      interface{} `json:"resources,omitempty"`
	Outputs        interface{} `json:"outputs,omitempty"`
}

// TemplateGenerator represents the object that performs the template generation.
type TemplateGenerator struct {
	Translator *i18n.Translator
}

// InitializeTemplateGenerator creates a new template generator object
func InitializeTemplateGenerator(ctx Context) (*TemplateGenerator, error) {
	t := &TemplateGenerator{
		Translator: ctx.Translator,
	}

	if t.Translator == nil {
		t.Translator = &i18n.Translator{}
	}

	if err := t.verifyFiles(); err != nil {
		return nil, err
	}

	return t, nil
}

// GenerateTemplate generates the template from the API Model
func (t *TemplateGenerator) GenerateTemplate(containerService *api.ContainerService, generatorCode string, aksEngineVersion string) (templateRaw string, parametersRaw string, err error) {
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

	var parametersMap = getParameters(containerService, generatorCode, aksEngineVersion)

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
	case api.SwarmMode:
		files = append(commonTemplateFiles, swarmModeTemplateFiles...)
		baseFile = swarmBaseFile
	default:
		return nil, "", t.Translator.Errorf("orchestrator '%s' is unsupported", properties.OrchestratorProfile.OrchestratorType)
	}

	return files, baseFile, nil
}

func (t *TemplateGenerator) GetJumpboxCustomDataJSON(cs *api.ContainerService) string {
	str, err := t.getSingleLineForTemplate(kubernetesJumpboxCustomDataYaml, cs, cs.Properties)

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("{\"customData\": \"[base64(concat('%s'))]\"}", str)
}

// GetMasterCustomDataJSONObject returns master customData JSON object in the form
// { "customData": "[base64(concat(<customData string>))]" }
func (t *TemplateGenerator) GetMasterCustomDataJSONObject(cs *api.ContainerService) string {
	profile := cs.Properties

	str, e := t.getSingleLineForTemplate(kubernetesMasterNodeCustomDataYaml, cs, profile)
	if e != nil {
		panic(e)
	}
	// add manifests
	componentStr := getComponentsString(cs, "k8s/manifests")
	str = strings.Replace(str, "MASTER_MANIFESTS_CONFIG_PLACEHOLDER", componentStr, -1)

	// add custom files
	customFilesReader, err := customfilesIntoReaders(masterCustomFiles(profile))
	if err != nil {
		log.Fatalf("Could not read custom files: %s", err.Error())
	}
	str = substituteConfigStringCustomFiles(str,
		customFilesReader,
		"MASTER_CUSTOM_FILES_PLACEHOLDER")

	addonStr := getAddonsString(cs, "k8s/addons")

	str = strings.Replace(str, "MASTER_CONTAINER_ADDONS_PLACEHOLDER", addonStr, -1)

	// return the custom data
	return fmt.Sprintf("{\"customData\": \"[base64(concat('%s'))]\"}", str)
}

// GetKubernetesLinuxNodeCustomDataJSONObject returns Linux customData JSON object in the form
// { "customData": "[base64(concat(<customData string>))]" }
func (t *TemplateGenerator) GetKubernetesLinuxNodeCustomDataJSONObject(cs *api.ContainerService, profile *api.AgentPoolProfile) string {
	str, e := t.getSingleLineForTemplate(kubernetesNodeCustomDataYaml, cs, profile)

	if e != nil {
		panic(e)
	}

	return fmt.Sprintf("{\"customData\": \"[base64(concat('%s'))]\"}", str)
}

// GetKubernetesWindowsNodeCustomDataJSONObject returns Windows customData JSON object in the form
// { "customData": "[base64(concat(<customData string>))]" }
func (t *TemplateGenerator) GetKubernetesWindowsNodeCustomDataJSONObject(cs *api.ContainerService, profile *api.AgentPoolProfile) string {
	str, e := t.getSingleLineForTemplate(kubernetesWindowsAgentCustomDataPS1, cs, profile)

	if e != nil {
		panic(e)
	}

	preprovisionCmd := ""

	if profile.PreprovisionExtension != nil {
		preprovisionCmd = makeAgentExtensionScriptCommands(cs, profile)
	}

	str = strings.Replace(str, "PREPROVISION_EXTENSION", escapeSingleLine(strings.TrimSpace(preprovisionCmd)), -1)

	return fmt.Sprintf("{\"customData\": \"[base64(concat('%s'))]\"}", str)
}

// getTemplateFuncMap returns the general purpose template func map from getContainerServiceFuncMap
func (t *TemplateGenerator) getTemplateFuncMap(cs *api.ContainerService) template.FuncMap {
	return getContainerServiceFuncMap(cs)
}

// getContainerServiceFuncMap returns all functions used in template generation
// These funcs are a thin wrapper for template generation operations,
// all business logic is implemented in the underlying func
func getContainerServiceFuncMap(cs *api.ContainerService) template.FuncMap {
	return template.FuncMap{
		"IsCustomCloudProfile": func() bool {
			return cs.Properties.IsCustomCloudProfile()
		},
		"GetCustomCloudRootCertificates": func() string {
			return cs.Properties.GetCustomCloudRootCertificates()
		},
		"GetCustomCloudSourcesList": func() string {
			return cs.Properties.GetCustomCloudSourcesList()
		},
		"IsAzureStackCloud": func() bool {
			return cs.Properties.IsAzureStackCloud()
		},
		"IsMultiMasterCluster": func() bool {
			return cs.Properties.MasterProfile != nil && cs.Properties.MasterProfile.HasMultipleNodes()
		},
		"IsMasterVirtualMachineScaleSets": func() bool {
			return cs.Properties.MasterProfile != nil && cs.Properties.MasterProfile.IsVirtualMachineScaleSets()
		},
		"IsVirtualMachineScaleSets": func(profile *api.AgentPoolProfile) bool {
			return profile.IsVirtualMachineScaleSets()
		},
		"IsHostedMaster": func() bool {
			return cs.Properties.IsHostedMasterProfile()
		},
		"IsIPMasqAgentEnabled": func() bool {
			return cs.Properties.IsIPMasqAgentEnabled()
		},
		"IsDCOS19": func() bool {
			return cs.Properties.OrchestratorProfile != nil && cs.Properties.OrchestratorProfile.IsDCOS19()
		},
		"IsKubernetesVersionGe": func(version string) bool {
			return cs.Properties.OrchestratorProfile.IsKubernetes() && common.IsKubernetesVersionGe(cs.Properties.OrchestratorProfile.OrchestratorVersion, version)
		},
		"IsKubernetesVersionLt": func(version string) bool {
			return cs.Properties.OrchestratorProfile.IsKubernetes() && !common.IsKubernetesVersionGe(cs.Properties.OrchestratorProfile.OrchestratorVersion, version)
		},
		"GetMasterKubernetesLabels": func(rg string) string {
			return common.GetMasterKubernetesLabels(rg, false)
		},
		"GetMasterKubernetesLabelsDeprecated": func(rg string) string {
			return common.GetMasterKubernetesLabels(rg, true)
		},
		"GetAgentKubernetesLabels": func(profile *api.AgentPoolProfile, rg string) string {
			return profile.GetKubernetesLabels(rg, false)
		},
		"GetAgentKubernetesLabelsDeprecated": func(profile *api.AgentPoolProfile, rg string) string {
			return profile.GetKubernetesLabels(rg, true)
		},
		"GetKubeletConfigKeyVals": func(kc *api.KubernetesConfig) string {
			if kc == nil {
				return ""
			}
			return kc.GetOrderedKubeletConfigString()
		},
		"GetKubeletConfigKeyValsPsh": func(kc *api.KubernetesConfig) string {
			if kc == nil {
				return ""
			}
			return kc.GetOrderedKubeletConfigStringForPowershell()
		},
		"HasKubeReservedCgroup": func() bool {
			kc := cs.Properties.OrchestratorProfile.KubernetesConfig
			return kc != nil && kc.KubeReservedCgroup != ""
		},

		"GetKubeReservedCgroup": func() string {
			kc := cs.Properties.OrchestratorProfile.KubernetesConfig
			if kc == nil {
				return ""
			}
			return kc.KubeReservedCgroup
		},
		"GetK8sRuntimeConfigKeyVals": func(config map[string]string) string {
			return common.GetOrderedEscapedKeyValsString(config)
		},
		"GetServiceCidr": func() string {
			return cs.Properties.OrchestratorProfile.KubernetesConfig.ServiceCIDR
		},
		"GetKubeProxyMode": func() string {
			return string(cs.Properties.OrchestratorProfile.KubernetesConfig.ProxyMode)
		},
		"HasPrivateRegistry": func() bool {
			if cs.Properties.OrchestratorProfile.DcosConfig != nil {
				return cs.Properties.OrchestratorProfile.DcosConfig.HasPrivateRegistry()
			}
			return false
		},
		"IsSwarmMode": func() bool {
			return cs.Properties.OrchestratorProfile.IsSwarmMode()
		},
		"IsKubernetes": func() bool {
			return cs.Properties.OrchestratorProfile.IsKubernetes()
		},
		"IsPublic": func(ports []int) bool {
			return common.SliceIntIsNonEmpty(ports)
		},
		"IsAzureCNI": func() bool {
			return cs.Properties.OrchestratorProfile.IsAzureCNI()
		},
		"HasClusterInitComponent": func() bool {
			_, enabled := cs.Properties.OrchestratorProfile.KubernetesConfig.IsComponentEnabled(common.ClusterInitComponentName)
			return enabled
		},
		"HasCosmosEtcd": func() bool {
			return cs.Properties.MasterProfile != nil && cs.Properties.MasterProfile.HasCosmosEtcd()
		},
		"GetCosmosEndPointUri": func() string {
			if cs.Properties.MasterProfile != nil {
				return cs.Properties.MasterProfile.GetCosmosEndPointURI()
			}
			return ""
		},
		"IsPrivateCluster": func() bool {
			return cs.Properties.OrchestratorProfile.IsPrivateCluster()
		},
		"ProvisionJumpbox": func() bool {
			return cs.Properties.OrchestratorProfile.KubernetesConfig.PrivateJumpboxProvision()
		},
		"UseManagedIdentity": func() bool {
			return cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity
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
			return cs.Properties.OrchestratorProfile.DcosConfig != nil && cs.Properties.OrchestratorProfile.DcosConfig.HasBootstrap()
		},
		"GetDCOSBootstrapCustomData": func() string {
			return getDCOSBootstrapCustomData(cs.Properties)
		},
		"GetDCOSMasterCustomData": func() string {
			return getDCOSMasterCustomData(cs)
		},
		"GetDCOSAgentCustomData": func(profile *api.AgentPoolProfile) string {
			return getDCOSAgentCustomData(cs, profile)
		},
		"GetDCOSWindowsAgentCustomData": func(profile *api.AgentPoolProfile) string {
			return getDCOSWindowsAgentCustomData(cs, profile)
		},
		"GetDCOSWindowsAgentCustomNodeAttributes": func(profile *api.AgentPoolProfile) string {
			return getDCOSWindowsAgentCustomAttributes(profile)
		},
		"GetDCOSWindowsAgentPreprovisionParameters": func(profile *api.AgentPoolProfile) string {
			if profile.PreprovisionExtension != nil {
				return getDCOSWindowsAgentPreprovisionParameters(cs, profile)
			}
			return ""
		},
		"GetMasterAllowedSizes": func() string {
			if cs.Properties.OrchestratorProfile.OrchestratorType == api.DCOS {
				return helpers.GetDCOSMasterAllowedSizes()
			}
			return helpers.GetKubernetesAllowedVMSKUs()
		},
		"GetDefaultVNETCIDR": func() string {
			return DefaultVNETCIDR
		},
		"GetKubernetesAllowedVMSKUs": func() string {
			return helpers.GetKubernetesAllowedVMSKUs()
		},
		"GetDefaultVNETCIDRIPv6": func() string {
			return DefaultVNETCIDRIPv6
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
		"WriteLinkedTemplatesForExtensions": func() string {
			return getLinkedTemplatesForExtensions(cs.Properties)
		},
		"GetSshPublicKeysPowerShell": func() string {
			return getSSHPublicKeysPowerShell(cs.Properties.LinuxProfile)
		},
		"GetWindowsMasterSubnetARMParam": func() string {
			return getWindowsMasterSubnetARMParam(cs.Properties.MasterProfile)
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
			str := buildYamlFileWithWriteFiles(files, cs)
			if cs.Properties.MasterProfile.PreprovisionExtension != nil {
				extensionStr := makeMasterExtensionScriptCommands(cs)
				str += "'runcmd:\n" + extensionStr + "\n\n'"
			}
			str = escapeSingleLine(str)
			return fmt.Sprintf("\"customData\": \"[base64(concat('%s'))]\",", str)
		},
		"GetAgentSwarmCustomData": func(profile *api.AgentPoolProfile) string {
			files := []string{swarmProvision}
			str := buildYamlFileWithWriteFiles(files, cs)
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
			str := getBase64EncodedGzippedCustomScript(swarmWindowsProvision, cs)
			return fmt.Sprintf("\"customData\": \"%s\"", str)
		},
		"GetWinAgentSwarmModeCustomData": func() string {
			str := getBase64EncodedGzippedCustomScript(swarmModeWindowsProvision, cs)
			return fmt.Sprintf("\"customData\": \"%s\"", str)
		},
		"GetKubernetesWindowsAgentFunctions": func() string {
			// Collect all the parts into a zip
			var parts = []string{
				kubernetesWindowsAgentFunctionsPS1,
				kubernetesWindowsConfigFunctionsPS1,
				kubernetesWindowsContainerdFunctionsPS1,
				kubernetesWindowsCsiProxyFunctionsPS1,
				kubernetesWindowsKubeletFunctionsPS1,
				kubernetesWindowsCniFunctionsPS1,
				kubernetesWindowsAzureCniFunctionsPS1,
				kubernetesWindowsLogsCleanupPS1,
				kubernetesWindowsNodeResetPS1,
				kubernetesWindowsOpenSSHFunctionPS1,
				kubeletStartPS1,
				kubeproxyStartPS1,
			}

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
				_, err = f.Write(partContents)
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
		"GetMasterSwarmModeCustomData": func() string {
			files := []string{swarmModeProvision}
			str := buildYamlFileWithWriteFiles(files, cs)
			if cs.Properties.MasterProfile.PreprovisionExtension != nil {
				extensionStr := makeMasterExtensionScriptCommands(cs)
				str += "runcmd:\n" + extensionStr + "\n\n"
			}
			str = escapeSingleLine(str)
			return fmt.Sprintf("\"customData\": \"[base64(concat('%s'))]\",", str)
		},
		"GetAgentSwarmModeCustomData": func(profile *api.AgentPoolProfile) string {
			files := []string{swarmModeProvision}
			str := buildYamlFileWithWriteFiles(files, cs)
			str = escapeSingleLine(str)
			return fmt.Sprintf("\"customData\": \"[base64(concat('%s',variables('%sRunCmdFile'),variables('%sRunCmd')))]\",", str, profile.Name, profile.Name)
		},
		"WrapAsVariable": func(s string) string {
			return common.WrapAsARMVariable(s)
		},
		"CloudInitData": func(s string) string {
			return wrapAsVariableObject("cloudInitFiles", s)
		},
		"WrapAsParameter": func(s string) string {
			return common.WrapAsParameter(s)
		},
		"WrapAsVerbatim": func(s string) string {
			return common.WrapAsVerbatim(s)
		},
		"HasVMASAgentPool": func() bool {
			return cs.Properties.HasVMASAgentPool()
		},
		"AnyAgentIsLinux": func() bool {
			return cs.Properties.AnyAgentIsLinux()
		},
		"IsNSeriesSKU": func(vmSKU string) bool {
			return common.IsNvidiaEnabledSKU(vmSKU)
		},
		"HasAvailabilityZones": func(profile *api.AgentPoolProfile) bool {
			return profile.HasAvailabilityZones()
		},
		"HasLinuxSecrets": func() bool {
			return cs.Properties.LinuxProfile.HasSecrets()
		},
		"HasCustomSearchDomain": func() bool {
			return cs.Properties.LinuxProfile != nil && cs.Properties.LinuxProfile.HasSearchDomain()
		},
		"GetSearchDomainName": func() string {
			if cs.Properties.LinuxProfile != nil && cs.Properties.LinuxProfile.HasSearchDomain() {
				return cs.Properties.LinuxProfile.CustomSearchDomain.Name
			}
			return ""
		},
		"GetSearchDomainRealmUser": func() string {
			if cs.Properties.LinuxProfile != nil && cs.Properties.LinuxProfile.HasSearchDomain() {
				return cs.Properties.LinuxProfile.CustomSearchDomain.RealmUser
			}
			return ""
		},
		"GetSearchDomainRealmPassword": func() string {
			if cs.Properties.LinuxProfile != nil && cs.Properties.LinuxProfile.HasSearchDomain() {
				return cs.Properties.LinuxProfile.CustomSearchDomain.RealmPassword
			}
			return ""
		},
		"HasCiliumNetworkPlugin": func() bool {
			return cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin == NetworkPluginCilium
		},
		"HasCiliumNetworkPolicy": func() bool {
			return cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy == NetworkPolicyCilium
		},
		"HasAntreaNetworkPolicy": func() bool {
			return cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy == NetworkPolicyAntrea
		},
		"HasFlannelNetworkPlugin": func() bool {
			return cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin == NetworkPluginFlannel
		},
		"HasCustomNodesDNS": func() bool {
			return cs.Properties.LinuxProfile != nil && cs.Properties.LinuxProfile.HasCustomNodesDNS()
		},
		"HasWindowsSecrets": func() bool {
			return cs.Properties.WindowsProfile.HasSecrets()
		},
		"HasWindowsCustomImage": func() bool {
			return cs.Properties.WindowsProfile.HasCustomImage()
		},
		"WindowsSSHEnabled": func() bool {
			return cs.Properties.WindowsProfile.GetSSHEnabled()
		},
		"GetConfigurationScriptRootURL": func() string {
			linuxProfile := cs.Properties.LinuxProfile
			if linuxProfile == nil || linuxProfile.ScriptRootURL == "" {
				return DefaultConfigurationScriptRootURL
			}
			return linuxProfile.ScriptRootURL
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
		"HasVHDDistroNodes": func() bool {
			return cs.Properties.HasVHDDistroNodes()
		},
		"IsVHDDistroForAllNodes": func() bool {
			return cs.Properties.IsVHDDistroForAllNodes()
		},
		"UseCloudControllerManager": func() bool {
			return cs.Properties.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager != nil && *cs.Properties.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager
		},
		"AdminGroupID": func() bool {
			return cs.Properties.AADProfile != nil && cs.Properties.AADProfile.AdminGroupID != ""
		},
		"EnableDataEncryptionAtRest": func() bool {
			return to.Bool(cs.Properties.OrchestratorProfile.KubernetesConfig.EnableDataEncryptionAtRest)
		},
		"EnableEncryptionWithExternalKms": func() bool {
			return to.Bool(cs.Properties.OrchestratorProfile.KubernetesConfig.EnableEncryptionWithExternalKms)
		},
		"EnableAggregatedAPIs": func() bool {
			if cs.Properties.OrchestratorProfile.KubernetesConfig.EnableAggregatedAPIs {
				return true
			} else if common.IsKubernetesVersionGe(cs.Properties.OrchestratorProfile.OrchestratorVersion, "1.9.0") {
				return true
			}
			return false
		},
		"IsCustomVNET": func() bool {
			return cs.Properties.AreAgentProfilesCustomVNET()
		},
		"IsIPv6DualStackFeatureEnabled": func() bool {
			return cs.Properties.FeatureFlags.IsFeatureEnabled("EnableIPv6DualStack")
		},
		"IsIPv6Enabled": func() bool {
			return cs.Properties.FeatureFlags.IsFeatureEnabled("EnableIPv6Only") || cs.Properties.FeatureFlags.IsFeatureEnabled("EnableIPv6DualStack")
		},
		"GetBase64EncodedEnvironmentJSON": func() string {
			customEnvironmentJSON, _ := cs.Properties.GetCustomEnvironmentJSON(false)
			return base64.StdEncoding.EncodeToString([]byte(customEnvironmentJSON))
		},
		"GetIdentitySystem": func() string {
			if cs.Properties.IsCustomCloudProfile() {
				return cs.Properties.CustomCloudProfile.IdentitySystem
			}

			return api.AzureADIdentitySystem
		},
		"GetPodInfraContainerSpec": func() string {
			return cs.Properties.OrchestratorProfile.GetPodInfraContainerSpec()
		},
		"IsKubenet": func() bool {
			return cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin == NetworkPluginKubenet
		},
		"NeedsContainerd": func() bool {
			return cs.Properties.OrchestratorProfile.KubernetesConfig.NeedsContainerd()
		},
		"IsDockerContainerRuntime": func() bool {
			return cs.Properties.OrchestratorProfile.KubernetesConfig.ContainerRuntime == api.Docker
		},
		"GetDockerConfig": func(hasGPU bool) string {
			val, err := getDockerConfig(cs, hasGPU)
			if err != nil {
				return ""
			}
			return val
		},
		"GetContainerdConfig": func() string {
			val, err := getContainerdConfig(cs)
			if err != nil {
				return ""
			}
			return val
		},
		"HasNSeriesSKU": func() bool {
			return cs.Properties.HasNSeriesSKU()
		},
		"HasDCSeriesSKU": func() bool {
			return cs.Properties.HasDCSeriesSKU()
		},
		"RequiresDocker": func() bool {
			return cs.Properties.OrchestratorProfile.KubernetesConfig.RequiresDocker()
		},
		"IsAzurePolicyAddonEnabled": func() bool {
			return cs.Properties.OrchestratorProfile.KubernetesConfig.IsAddonEnabled(common.AzurePolicyAddonName)
		},
		"IsACIConnectorAddonEnabled": func() bool {
			return cs.Properties.OrchestratorProfile.KubernetesConfig.IsAddonEnabled(common.ACIConnectorAddonName)
		},
		"IsClusterAutoscalerAddonEnabled": func() bool {
			return cs.Properties.OrchestratorProfile.KubernetesConfig.IsAddonEnabled(common.ClusterAutoscalerAddonName)
		},
		"IsAADPodIdentityAddonEnabled": func() bool {
			return cs.Properties.OrchestratorProfile.KubernetesConfig.IsAddonEnabled(common.AADPodIdentityAddonName)
		},
		"GetAADPodIdentityTaintKey": func() string {
			return common.AADPodIdentityTaintKey
		},
		"GetHyperkubeImageReference": func() string {
			hyperkubeImageBase := cs.Properties.OrchestratorProfile.KubernetesConfig.KubernetesImageBase
			k8sComponents := api.GetK8sComponentsByVersionMap(cs.Properties.OrchestratorProfile.KubernetesConfig)[cs.Properties.OrchestratorProfile.OrchestratorVersion]
			hyperkubeImage := hyperkubeImageBase + k8sComponents[common.Hyperkube]
			if cs.Properties.IsAzureStackCloud() {
				hyperkubeImage = hyperkubeImage + common.AzureStackSuffix
			}
			if cs.Properties.OrchestratorProfile.KubernetesConfig.CustomHyperkubeImage != "" {
				hyperkubeImage = cs.Properties.OrchestratorProfile.KubernetesConfig.CustomHyperkubeImage
			}
			return hyperkubeImage
		},
		"GetTargetEnvironment": func() string {
			return helpers.GetTargetEnv(cs.Location, cs.Properties.GetCustomCloudName())
		},
		"GetCustomCloudConfigCSEScriptFilepath": func() string {
			return customCloudConfigCSEScriptFilepath
		},
		"GetCSEHelpersScriptFilepath": func() string {
			return cseHelpersScriptFilepath
		},
		"GetCSEInstallScriptFilepath": func() string {
			return cseInstallScriptFilepath
		},
		"GetCSEConfigScriptFilepath": func() string {
			return cseConfigScriptFilepath
		},
		"GetCustomSearchDomainsCSEScriptFilepath": func() string {
			return customSearchDomainsCSEScriptFilepath
		},
		"GetDHCPv6ServiceCSEScriptFilepath": func() string {
			return dhcpV6ServiceCSEScriptFilepath
		},
		"GetDHCPv6ConfigCSEScriptFilepath": func() string {
			return dhcpV6ConfigCSEScriptFilepath
		},
		"HasPrivateAzureRegistryServer": func() bool {
			return cs.Properties.OrchestratorProfile.KubernetesConfig.PrivateAzureRegistryServer != ""
		},
		"GetPrivateAzureRegistryServer": func() string {
			return cs.Properties.OrchestratorProfile.KubernetesConfig.PrivateAzureRegistryServer
		},
		"HasTelemetryEnabled": func() bool {
			return cs.Properties.FeatureFlags != nil && cs.Properties.FeatureFlags.EnableTelemetry
		},
		"GetCSEErrorCode": func(errorType string) int {
			return GetCSEErrorCode(errorType)
		},
		"GetApplicationInsightsTelemetryKeys": func() string {
			userSuppliedAIKey := ""
			if cs.Properties.TelemetryProfile != nil {
				userSuppliedAIKey = cs.Properties.TelemetryProfile.ApplicationInsightsKey
			}

			possibleKeys := []string{
				telemetry.AKSEngineAppInsightsKey,
				userSuppliedAIKey,
			}

			var keys []string
			for _, key := range possibleKeys {
				if key != "" {
					keys = append(keys, key)
				}
			}

			return strings.Join(keys, ",")
		},
		"GetLinuxDefaultTelemetryTags": func() string {
			tags := map[string]string{
				"k8s_version":    cs.Properties.OrchestratorProfile.OrchestratorVersion,
				"network_plugin": cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin,
				"network_policy": cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy,
				"network_mode":   cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkMode,
				"cri":            cs.Properties.OrchestratorProfile.KubernetesConfig.ContainerRuntime,
				"cri_version":    cs.Properties.OrchestratorProfile.KubernetesConfig.ContainerdVersion,
				"distro":         string(cs.Properties.LinuxProfile.Distro),
				"os_image_sku":   cs.GetCloudSpecConfig().OSImageConfig[cs.Properties.LinuxProfile.Distro].ImageSku,
				"os_type":        "linux",
			}

			var kvs []string
			for k, v := range tags {
				if v != "" {
					kvs = append(kvs, fmt.Sprintf("%s=%s", k, v))
				}
			}
			sort.Strings(kvs)
			return strings.Join(kvs, ",")
		},
		"GetSysctlDConfigKeyVals": func(sysctlDConfig map[string]string) string {
			return common.GetOrderedNewlinedKeyValsStringForCloudInit(sysctlDConfig)
		},
		"OpenBraces": func() string {
			return "{{"
		},
		"CloseBraces": func() string {
			return "}}"
		},
		"IndentString": func(original string, spaces int) string {
			return common.IndentString(original, spaces)
		},
	}
}

func (t *TemplateGenerator) GenerateTemplateV2(containerService *api.ContainerService, generatorCode string, acsengineVersion string) (templateRaw string, parametersRaw string, err error) {

	armParams, _ := t.getParameterDescMap(containerService)
	armResources := GenerateARMResources(containerService)
	armVariables, err := GetKubernetesVariables(containerService)
	if err != nil {
		return "", "", err
	}
	armOutputs := GetKubernetesOutputs(containerService)

	armTemplate := ARMTemplate{
		Schema:         "https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
		ContentVersion: "1.0.0.0",
		Parameters:     armParams,
		Variables:      armVariables,
		Resources:      armResources,
		Outputs:        armOutputs,
	}

	var templBytes []byte
	if templBytes, err = json.Marshal(armTemplate); err != nil {
		return "", "", err
	}
	templateRaw = string(templBytes)

	var parametersMap = getParameters(containerService, generatorCode, acsengineVersion)

	var parameterBytes []byte
	if parameterBytes, err = helpers.JSONMarshal(parametersMap, false); err != nil {
		return "", "", err
	}
	parametersRaw = string(parameterBytes)

	return templateRaw, parametersRaw, err
}

func (t *TemplateGenerator) getParameterDescMap(containerService *api.ContainerService) (interface{}, error) {
	var templ *template.Template
	var paramsDescMap map[string]interface{}
	properties := containerService.Properties
	// save the current orchestrator version and restore it after deploying.
	// this allows us to deploy agents on the most recent patch without updating the orchestrator version in the object
	orchVersion := properties.OrchestratorProfile.OrchestratorVersion
	defer func() {
		properties.OrchestratorProfile.OrchestratorVersion = orchVersion
	}()

	templ = template.New("acs template").Funcs(t.getTemplateFuncMap(containerService))

	files, baseFile := kubernetesParamFiles, armParameters

	for _, file := range files {
		bytes, e := Asset(file)
		if e != nil {
			err := t.Translator.Errorf("Error reading file %s, Error: %s", file, e.Error())
			return nil, err
		}
		if _, err := templ.New(file).Parse(string(bytes)); err != nil {
			return nil, err
		}
	}

	var b bytes.Buffer
	if err := templ.ExecuteTemplate(&b, baseFile, properties); err != nil {
		return nil, err
	}

	err := json.Unmarshal(b.Bytes(), &paramsDescMap)

	if err != nil {
		return nil, err
	}

	return paramsDescMap["parameters"], nil
}

func generateUserAssignedIdentityClientIDParameter(isUserAssignedIdentity bool) string {
	if isUserAssignedIdentity {
		return "' USER_ASSIGNED_IDENTITY_ID=',reference(variables('userAssignedIDReference'), variables('apiVersionManagedIdentity')).clientId, ' '"
	}
	return "' USER_ASSIGNED_IDENTITY_ID=',' '"
}

func getDockerConfig(cs *api.ContainerService, hasGPU bool) (string, error) {
	var overrides []func(*common.DockerConfig) error

	if hasGPU {
		overrides = append(overrides, common.DockerNvidiaOverride)
	}

	val, err := common.GetDockerConfig(cs.Properties.OrchestratorProfile.KubernetesConfig.ContainerRuntimeConfig, overrides)
	if err != nil {
		return "", err
	}

	return val, nil
}

func getContainerdConfig(cs *api.ContainerService) (string, error) {
	var overrides = []func(*common.ContainerdConfig) error{
		common.ContainerdSandboxImageOverrider(cs.Properties.OrchestratorProfile.GetPodInfraContainerSpec()),
	}

	if cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin == NetworkPluginKubenet {
		overrides = append(overrides, common.ContainerdKubenetOverride)
	}

	val, err := common.GetContainerdConfig(cs.Properties.OrchestratorProfile.KubernetesConfig.ContainerRuntimeConfig, overrides)
	if err != nil {
		return "", err
	}

	return val, nil
}
