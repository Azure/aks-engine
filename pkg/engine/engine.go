// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/pkg/errors"

	_ "k8s.io/client-go/plugin/pkg/client/auth/azure" // register azure (AD) authentication plugin
)

var commonTemplateFiles = []string{agentOutputs, agentParams, masterOutputs, iaasOutputs, masterParams, windowsParams}
var dcosTemplateFiles = []string{dcosBaseFile, dcosAgentResourcesVMAS, dcosAgentResourcesVMSS, dcosAgentVars, dcosMasterResources, dcosMasterVars, dcosParams, dcosWindowsAgentResourcesVMAS, dcosWindowsAgentResourcesVMSS}
var dcos2TemplateFiles = []string{dcos2BaseFile, dcosAgentResourcesVMAS, dcosAgentResourcesVMSS, dcosAgentVars, dcos2MasterResources, dcos2BootstrapResources, dcos2MasterVars, dcosParams, dcosWindowsAgentResourcesVMAS, dcosWindowsAgentResourcesVMSS, dcos2BootstrapVars, dcos2BootstrapParams}
var kubernetesParamFiles = []string{armParameters, kubernetesParams, masterParams, agentParams, windowsParams}
var swarmTemplateFiles = []string{swarmBaseFile, swarmParams, swarmAgentResourcesVMAS, swarmAgentVars, swarmAgentResourcesVMSS, swarmBaseFile, swarmMasterResources, swarmMasterVars, swarmWinAgentResourcesVMAS, swarmWinAgentResourcesVMSS}
var swarmModeTemplateFiles = []string{swarmBaseFile, swarmParams, swarmAgentResourcesVMAS, swarmAgentVars, swarmAgentResourcesVMSS, swarmBaseFile, swarmMasterResources, swarmMasterVars, swarmWinAgentResourcesVMAS, swarmWinAgentResourcesVMSS}

var keyvaultSecretPathRe *regexp.Regexp

func init() {
	keyvaultSecretPathRe = regexp.MustCompile(`^(/subscriptions/\S+/resourceGroups/\S+/providers/Microsoft.KeyVault/vaults/\S+)/secrets/([^/\s]+)(/(\S+))?$`)
}

// GenerateKubeConfig returns a JSON string representing the KubeConfig
func GenerateKubeConfig(properties *api.Properties, location string) (string, error) {
	if properties == nil {
		return "", errors.New("Properties nil in GenerateKubeConfig")
	}
	if properties.CertificateProfile == nil {
		return "", errors.New("CertificateProfile property may not be nil in GenerateKubeConfig")
	}
	b, err := Asset(kubeConfigJSON)
	if err != nil {
		return "", errors.Wrapf(err, "error reading kube config template file %s", kubeConfigJSON)
	}
	kubeconfig := string(b)
	// variable replacement
	kubeconfig = strings.Replace(kubeconfig, "{{WrapAsVerbatim \"parameters('caCertificate')\"}}", base64.StdEncoding.EncodeToString([]byte(properties.CertificateProfile.CaCertificate)), -1)
	if properties.OrchestratorProfile != nil &&
		properties.OrchestratorProfile.KubernetesConfig != nil &&
		properties.OrchestratorProfile.KubernetesConfig.PrivateCluster != nil &&
		to.Bool(properties.OrchestratorProfile.KubernetesConfig.PrivateCluster.Enabled) {
		if properties.MasterProfile.HasMultipleNodes() {
			// more than 1 master, use the internal lb IP
			firstMasterIP := net.ParseIP(properties.MasterProfile.FirstConsecutiveStaticIP).To4()
			if firstMasterIP == nil {
				return "", errors.Errorf("MasterProfile.FirstConsecutiveStaticIP '%s' is an invalid IP address", properties.MasterProfile.FirstConsecutiveStaticIP)
			}
			lbIP := net.IP{firstMasterIP[0], firstMasterIP[1], firstMasterIP[2], firstMasterIP[3] + byte(DefaultInternalLbStaticIPOffset)}
			kubeconfig = strings.Replace(kubeconfig, "{{WrapAsVerbatim \"reference(concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))).dnsSettings.fqdn\"}}", lbIP.String(), -1)
		} else {
			// Master count is 1, use the master IP
			kubeconfig = strings.Replace(kubeconfig, "{{WrapAsVerbatim \"reference(concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))).dnsSettings.fqdn\"}}", properties.MasterProfile.FirstConsecutiveStaticIP, -1)
		}
	} else {
		kubeconfig = strings.Replace(kubeconfig, "{{WrapAsVerbatim \"reference(concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))).dnsSettings.fqdn\"}}", api.FormatProdFQDNByLocation(properties.MasterProfile.DNSPrefix, location, properties.GetCustomCloudName()), -1)
	}
	kubeconfig = strings.Replace(kubeconfig, "{{WrapAsVariable \"resourceGroup\"}}", properties.MasterProfile.DNSPrefix, -1)

	var authInfo string
	if properties.AADProfile == nil {
		authInfo = fmt.Sprintf("{\"client-certificate-data\":\"%v\",\"client-key-data\":\"%v\"}",
			base64.StdEncoding.EncodeToString([]byte(properties.CertificateProfile.KubeConfigCertificate)),
			base64.StdEncoding.EncodeToString([]byte(properties.CertificateProfile.KubeConfigPrivateKey)))
	} else {
		tenantID := properties.AADProfile.TenantID
		if len(tenantID) == 0 {
			tenantID = "common"
		}

		authInfo = fmt.Sprintf("{\"auth-provider\":{\"name\":\"azure\",\"config\":{\"environment\":\"%v\",\"tenant-id\":\"%v\",\"apiserver-id\":\"%v\",\"client-id\":\"%v\"}}}",
			helpers.GetTargetEnv(location, properties.GetCustomCloudName()),
			tenantID,
			properties.AADProfile.ServerAppID,
			properties.AADProfile.ClientAppID)
	}
	kubeconfig = strings.Replace(kubeconfig, "{{authInfo}}", authInfo, -1)

	return kubeconfig, nil
}

// validateDistro checks if the requested orchestrator type is supported on the requested Linux distro.
func validateDistro(cs *api.ContainerService) bool {
	// Check Master distro
	if cs.Properties.MasterProfile != nil && cs.Properties.MasterProfile.Distro == api.RHEL &&
		(cs.Properties.OrchestratorProfile.OrchestratorType != api.SwarmMode) {
		log.Printf("Orchestrator type %s not suported on RHEL Master", cs.Properties.OrchestratorProfile.OrchestratorType)
		return false
	}
	// Check Agent distros
	for _, agentProfile := range cs.Properties.AgentPoolProfiles {
		if agentProfile.Distro == api.RHEL &&
			(cs.Properties.OrchestratorProfile.OrchestratorType != api.SwarmMode) {
			log.Printf("Orchestrator type %s not suported on RHEL Agent", cs.Properties.OrchestratorProfile.OrchestratorType)
			return false
		}
	}
	return true
}

// generateConsecutiveIPsList takes a starting IP address and returns a string slice of length "count" of subsequent, consecutive IP addresses
func generateConsecutiveIPsList(count int, firstAddr string) ([]string, error) {
	ipaddr := net.ParseIP(firstAddr).To4()
	if ipaddr == nil {
		return nil, errors.Errorf("IPAddr '%s' is an invalid IP address", firstAddr)
	}
	if int(ipaddr[3])+count >= 255 {
		return nil, errors.Errorf("IPAddr '%s' + %d will overflow the fourth octet", firstAddr, count)
	}
	ret := make([]string, count)
	for i := 0; i < count; i++ {
		nextAddress := fmt.Sprintf("%d.%d.%d.%d", ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3]+byte(i))
		ipaddr := net.ParseIP(nextAddress).To4()
		if ipaddr == nil {
			return nil, errors.Errorf("IPAddr '%s' is an invalid IP address", nextAddress)
		}
		ret[i] = nextAddress
	}
	return ret, nil
}

func addValue(m paramsMap, k string, v interface{}) {
	m[k] = paramsMap{
		"value": v,
	}
}

func addKeyvaultReference(m paramsMap, k string, vaultID, secretName, secretVersion string) {
	m[k] = paramsMap{
		"reference": &KeyVaultRef{
			KeyVault: KeyVaultID{
				ID: vaultID,
			},
			SecretName:    secretName,
			SecretVersion: secretVersion,
		},
	}
}

func addSecret(m paramsMap, k string, v interface{}, encode bool) {
	str, ok := v.(string)
	if !ok {
		addValue(m, k, v)
		return
	}
	parts := keyvaultSecretPathRe.FindStringSubmatch(str)
	if parts == nil || len(parts) != 5 {
		if encode {
			addValue(m, k, base64.StdEncoding.EncodeToString([]byte(str)))
		} else {
			addValue(m, k, str)
		}
		return
	}
	addKeyvaultReference(m, k, parts[1], parts[2], parts[4])
}

func makeMasterExtensionScriptCommands(cs *api.ContainerService) string {
	curlCaCertOpt := ""
	if cs.Properties.IsAzureStackCloud() {
		curlCaCertOpt = fmt.Sprintf("--cacert %s", common.AzureStackCaCertLocation)
	}
	return makeExtensionScriptCommands(cs.Properties.MasterProfile.PreprovisionExtension,
		curlCaCertOpt, cs.Properties.ExtensionProfiles)
}

func makeAgentExtensionScriptCommands(cs *api.ContainerService, profile *api.AgentPoolProfile) string {
	if profile.OSType == api.Windows {
		return makeWindowsExtensionScriptCommands(profile.PreprovisionExtension,
			cs.Properties.ExtensionProfiles)
	}
	curlCaCertOpt := ""
	if cs.Properties.IsAzureStackCloud() {
		curlCaCertOpt = fmt.Sprintf("--cacert %s", common.AzureStackCaCertLocation)
	}
	return makeExtensionScriptCommands(profile.PreprovisionExtension,
		curlCaCertOpt, cs.Properties.ExtensionProfiles)
}

func makeExtensionScriptCommands(extension *api.Extension, curlCaCertOpt string, extensionProfiles []*api.ExtensionProfile) string {
	var extensionProfile *api.ExtensionProfile
	for _, eP := range extensionProfiles {
		if strings.EqualFold(eP.Name, extension.Name) {
			extensionProfile = eP
			break
		}
	}

	if extensionProfile == nil {
		panic(fmt.Sprintf("%s extension referenced was not found in the extension profile", extension.Name))
	}

	extensionsParameterReference := fmt.Sprintf("parameters('%sParameters')", extensionProfile.Name)
	scriptURL := getExtensionURL(extensionProfile.RootURL, extensionProfile.Name, extensionProfile.Version, extensionProfile.Script, extensionProfile.URLQuery)
	scriptFilePath := fmt.Sprintf("/opt/azure/containers/extensions/%s/%s", extensionProfile.Name, extensionProfile.Script)
	return fmt.Sprintf("- sudo /usr/bin/curl --retry 5 --retry-delay 10 --retry-max-time 30 -o %s --create-dirs %s \"%s\" \n- sudo /bin/chmod 744 %s \n- sudo %s ',%s,' > /var/log/%s-output.log",
		scriptFilePath, curlCaCertOpt, scriptURL, scriptFilePath, scriptFilePath, extensionsParameterReference, extensionProfile.Name)
}

func makeWindowsExtensionScriptCommands(extension *api.Extension, extensionProfiles []*api.ExtensionProfile) string {
	var extensionProfile *api.ExtensionProfile
	for _, eP := range extensionProfiles {
		if strings.EqualFold(eP.Name, extension.Name) {
			extensionProfile = eP
			break
		}
	}

	if extensionProfile == nil {
		panic(fmt.Sprintf("%s extension referenced was not found in the extension profile", extension.Name))
	}

	scriptURL := getExtensionURL(extensionProfile.RootURL, extensionProfile.Name, extensionProfile.Version, extensionProfile.Script, extensionProfile.URLQuery)
	scriptFileDir := fmt.Sprintf("$env:SystemDrive:/AzureData/extensions/%s", extensionProfile.Name)
	scriptFilePath := fmt.Sprintf("%s/%s", scriptFileDir, extensionProfile.Script)
	return fmt.Sprintf("New-Item -ItemType Directory -Force -Path \"%s\" ; Invoke-WebRequest -Uri \"%s\" -OutFile \"%s\" ; powershell \"%s `\"',parameters('%sParameters'),'`\"\"\n", scriptFileDir, scriptURL, scriptFilePath, scriptFilePath, extensionProfile.Name)
}

func getDCOSWindowsAgentPreprovisionParameters(cs *api.ContainerService, profile *api.AgentPoolProfile) string {
	extension := profile.PreprovisionExtension

	var extensionProfile *api.ExtensionProfile

	for _, eP := range cs.Properties.ExtensionProfiles {
		if strings.EqualFold(eP.Name, extension.Name) {
			extensionProfile = eP
			break
		}
	}

	parms := extensionProfile.ExtensionParameters
	return parms
}

func getDCOSDefaultBootstrapInstallerURL(profile *api.OrchestratorProfile) string {
	if profile.OrchestratorType == api.DCOS {
		switch profile.OrchestratorVersion {
		case common.DCOSVersion1Dot11Dot2:
			return "https://dcos-mirror.azureedge.net/dcos-1-11-2/dcos_generate_config.sh"
		case common.DCOSVersion1Dot11Dot0:
			return "https://dcos-mirror.azureedge.net/dcos-1-11-0/dcos_generate_config.sh"
		}
	}
	return ""
}

func getDCOSDefaultWindowsBootstrapInstallerURL(profile *api.OrchestratorProfile) string {
	if profile.OrchestratorType == api.DCOS {
		switch profile.OrchestratorVersion {
		case common.DCOSVersion1Dot11Dot2:
			return "https://dcos-mirror.azureedge.net/dcos-windows/1-11-2"
		case common.DCOSVersion1Dot11Dot0:
			return "https://dcos-mirror.azureedge.net/dcos-windows/1-11-0"
		}
	}
	return ""
}

func getDCOSDefaultProviderPackageGUID(orchestratorType string, orchestratorVersion string, masterCount int) string {
	if orchestratorType == api.DCOS {
		switch orchestratorVersion {
		case common.DCOSVersion1Dot10Dot0:
			switch masterCount {
			case 1:
				return "c4ec6210f396b8e435177b82e3280a2cef0ce721"
			case 3:
				return "08197947cb57d479eddb077a429fa15c139d7d20"
			case 5:
				return "f286ad9d3641da5abb622e4a8781f73ecd8492fa"
			}
		case common.DCOSVersion1Dot9Dot0:
			switch masterCount {
			case 1:
				return "bcc883b7a3191412cf41824bdee06c1142187a0b"
			case 3:
				return "dcff7e24c0c1827bebeb7f1a806f558054481b33"
			case 5:
				return "b41bfa84137a6374b2ff5eb1655364d7302bd257"
			}
		case common.DCOSVersion1Dot9Dot8:
			switch masterCount {
			case 1:
				return "e8b0e3fc4a16394dc6dd5b19fc54bf1543bff429"
			case 3:
				return "2d36c3f570d9dd7d187c699f9a322ed9d95e7dfa"
			case 5:
				return "c03c9587f88929f310b80af4f448b7b51654f1c8"
			}
		case common.DCOSVersion1Dot8Dot8:
			switch masterCount {
			case 1:
				return "441385ce2f5942df7e29075c12fb38fa5e92cbba"
			case 3:
				return "b1cd359287504efb780257bd12cc3a63704e42d4"
			case 5:
				return "d9b61156dfcc9383e014851529738aa550ef57d9"
			}
		}
	}
	return ""
}

func getDCOSDefaultRepositoryURL(orchestratorType string, orchestratorVersion string) string {
	if orchestratorType == api.DCOS {
		switch orchestratorVersion {
		case common.DCOSVersion1Dot10Dot0:
			return "https://dcosio.azureedge.net/dcos/stable/1.10.0"
		case common.DCOSVersion1Dot9Dot8:
			return "https://dcosio.azureedge.net/dcos/stable/1.9.8"
		default:
			return "https://dcosio.azureedge.net/dcos/stable"
		}
	}
	return ""
}

func getDCOSMasterCustomNodeLabels() string {
	// return empty string for DCOS since no attribtutes needed on master
	return ""
}

func getDCOSAgentCustomNodeLabels(profile *api.AgentPoolProfile) string {
	var buf bytes.Buffer
	var attrstring string
	buf.WriteString("")
	// always write MESOS_ATTRIBUTES because
	// the provision script will add FD/UD attributes
	// at node provisioning time
	if len(profile.OSType) > 0 {
		attrstring = fmt.Sprintf("MESOS_ATTRIBUTES=\"os:%s", profile.OSType)
	} else {
		attrstring = fmt.Sprintf("MESOS_ATTRIBUTES=\"os:%s", api.Linux)
	}

	if len(profile.Ports) > 0 {
		attrstring += ";public_ip:yes"
	}

	buf.WriteString(attrstring)
	if len(profile.CustomNodeLabels) > 0 {
		for k, v := range profile.CustomNodeLabels {
			buf.WriteString(fmt.Sprintf(";%s:%s", k, v))
		}
	}
	buf.WriteString("\"")
	return buf.String()
}

func getDCOSWindowsAgentCustomAttributes(profile *api.AgentPoolProfile) string {
	var buf bytes.Buffer
	var attrstring string
	buf.WriteString("")
	if len(profile.OSType) > 0 {
		attrstring = fmt.Sprintf("os:%s", profile.OSType)
	} else {
		attrstring = "os:windows"
	}
	if len(profile.Ports) > 0 {
		attrstring += ";public_ip:yes"
	}
	buf.WriteString(attrstring)
	if len(profile.CustomNodeLabels) > 0 {
		for k, v := range profile.CustomNodeLabels {
			buf.WriteString(fmt.Sprintf(";%s:%s", k, v))
		}
	}
	return buf.String()
}

func getVNETAddressPrefixes(properties *api.Properties) string {
	visitedSubnets := make(map[string]bool)
	var buf bytes.Buffer
	buf.WriteString(`"[variables('masterSubnet')]"`)
	visitedSubnets[properties.MasterProfile.Subnet] = true
	for _, profile := range properties.AgentPoolProfiles {
		if _, ok := visitedSubnets[profile.Subnet]; !ok {
			buf.WriteString(fmt.Sprintf(",\n            \"[variables('%sSubnet')]\"", profile.Name))
		}
	}
	return buf.String()
}

func getVNETSubnetDependencies(properties *api.Properties) string {
	agentString := `        "[concat('Microsoft.Network/networkSecurityGroups/', variables('%sNSGName'))]"`
	var buf bytes.Buffer
	for index, agentProfile := range properties.AgentPoolProfiles {
		if index > 0 {
			buf.WriteString(",\n")
		}
		buf.WriteString(fmt.Sprintf(agentString, agentProfile.Name))
	}
	return buf.String()
}

func getVNETSubnets(properties *api.Properties, addNSG bool) string {
	masterString := `{
            "name": "[variables('masterSubnetName')]",
            "properties": {
              "addressPrefix": "[variables('masterSubnet')]"
            }
          }`
	agentString := `          {
            "name": "[variables('%sSubnetName')]",
            "properties": {
              "addressPrefix": "[variables('%sSubnet')]"
            }
          }`
	agentStringNSG := `          {
            "name": "[variables('%sSubnetName')]",
            "properties": {
              "addressPrefix": "[variables('%sSubnet')]",
              "networkSecurityGroup": {
                "id": "[resourceId('Microsoft.Network/networkSecurityGroups', variables('%sNSGName'))]"
              }
            }
          }`
	var buf bytes.Buffer
	buf.WriteString(masterString)
	for _, agentProfile := range properties.AgentPoolProfiles {
		buf.WriteString(",\n")
		if addNSG {
			buf.WriteString(fmt.Sprintf(agentStringNSG, agentProfile.Name, agentProfile.Name, agentProfile.Name))
		} else {
			buf.WriteString(fmt.Sprintf(agentString, agentProfile.Name, agentProfile.Name))
		}

	}
	return buf.String()
}

func getLBRule(name string, port int) string {
	return fmt.Sprintf(`	          {
            "name": "LBRule%d",
            "properties": {
              "backendAddressPool": {
                "id": "[concat(variables('%sLbID'), '/backendAddressPools/', variables('%sLbBackendPoolName'))]"
              },
              "backendPort": %d,
              "enableFloatingIP": false,
              "frontendIPConfiguration": {
                "id": "[variables('%sLbIPConfigID')]"
              },
              "frontendPort": %d,
              "idleTimeoutInMinutes": 5,
              "loadDistribution": "Default",
              "probe": {
                "id": "[concat(variables('%sLbID'),'/probes/tcp%dProbe')]"
              },
              "protocol": "Tcp"
            }
          }`, port, name, name, port, name, port, name, port)
}

func getLBRules(name string, ports []int) string {
	var buf bytes.Buffer
	for index, port := range ports {
		if index > 0 {
			buf.WriteString(",\n")
		}
		buf.WriteString(getLBRule(name, port))
	}
	return buf.String()
}

func getProbe(port int) string {
	return fmt.Sprintf(`          {
            "name": "tcp%dProbe",
            "properties": {
              "intervalInSeconds": 5,
              "numberOfProbes": 2,
              "port": %d,
              "protocol": "Tcp"
            }
          }`, port, port)
}

func getProbes(ports []int) string {
	var buf bytes.Buffer
	for index, port := range ports {
		if index > 0 {
			buf.WriteString(",\n")
		}
		buf.WriteString(getProbe(port))
	}
	return buf.String()
}

func getSecurityRule(port int, portIndex int) string {
	// BaseLBPriority specifies the base lb priority.
	BaseLBPriority := 200
	return fmt.Sprintf(`          {
            "name": "Allow_%d",
            "properties": {
              "access": "Allow",
              "description": "Allow traffic from the Internet to port %d",
              "destinationAddressPrefix": "*",
              "destinationPortRange": "%d",
              "direction": "Inbound",
              "priority": %d,
              "protocol": "*",
              "sourceAddressPrefix": "Internet",
              "sourcePortRange": "*"
            }
          }`, port, port, port, BaseLBPriority+portIndex)
}

func getDataDisks(a *api.AgentPoolProfile) string {
	if !a.HasDisks() {
		return ""
	}
	var buf bytes.Buffer
	buf.WriteString("\"dataDisks\": [\n")
	dataDisks := `            {
              "createOption": "Empty",
              "diskSizeGB": "%d",
              "lun": %d,
              "caching": "ReadOnly",
              "name": "[concat(variables('%sVMNamePrefix'), copyIndex(),'-datadisk%d')]",
              "vhd": {
                "uri": "[concat('http://',variables('storageAccountPrefixes')[mod(add(add(div(copyIndex(),variables('maxVMsPerStorageAccount')),variables('%sStorageAccountOffset')),variables('dataStorageAccountPrefixSeed')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(add(div(copyIndex(),variables('maxVMsPerStorageAccount')),variables('%sStorageAccountOffset')),variables('dataStorageAccountPrefixSeed')),variables('storageAccountPrefixesCount'))],variables('%sDataAccountName'),'.blob.core.windows.net/vhds/',variables('%sVMNamePrefix'),copyIndex(), '--datadisk%d.vhd')]"
              }
            }`
	managedDataDisks := `            {
              "diskSizeGB": "%d",
              "lun": %d,
              "caching": "ReadOnly",
              "createOption": "Empty"
            }`
	for i, diskSize := range a.DiskSizesGB {
		if i > 0 {
			buf.WriteString(",\n")
		}
		if a.StorageProfile == api.StorageAccount {
			buf.WriteString(fmt.Sprintf(dataDisks, diskSize, i, a.Name, i, a.Name, a.Name, a.Name, a.Name, i))
		} else if a.StorageProfile == api.ManagedDisks {
			buf.WriteString(fmt.Sprintf(managedDataDisks, diskSize, i))
		}
	}
	buf.WriteString("\n          ],")
	return buf.String()
}

func getSecurityRules(ports []int) string {
	var buf bytes.Buffer
	for index, port := range ports {
		if index > 0 {
			buf.WriteString(",\n")
		}
		buf.WriteString(getSecurityRule(port, index))
	}
	return buf.String()
}

// getSingleLine returns the file as a single line
func (t *TemplateGenerator) getSingleLine(textFilename string, cs *api.ContainerService, profile interface{}) (string, error) {
	b, err := Asset(textFilename)
	if err != nil {
		return "", t.Translator.Errorf("yaml file %s does not exist", textFilename)
	}

	// use go templates to process the text filename
	templ := template.New("customdata template").Funcs(t.getTemplateFuncMap(cs))
	if _, err = templ.New(textFilename).Parse(string(b)); err != nil {
		return "", t.Translator.Errorf("error parsing file %s: %v", textFilename, err)
	}

	var buffer bytes.Buffer
	if err = templ.ExecuteTemplate(&buffer, textFilename, profile); err != nil {
		return "", t.Translator.Errorf("error executing template for file %s: %v", textFilename, err)
	}
	expandedTemplate := buffer.String()

	return expandedTemplate, nil
}

// getSingleLineForTemplate returns the file as a single line for embedding in an arm template
func (t *TemplateGenerator) getSingleLineForTemplate(textFilename string, cs *api.ContainerService, profile interface{}) (string, error) {
	expandedTemplate, err := t.getSingleLine(textFilename, cs, profile)
	if err != nil {
		return "", err
	}

	textStr := escapeSingleLine(expandedTemplate)

	return textStr, nil
}

func escapeSingleLine(escapedStr string) string {
	// template.JSEscapeString leaves undesirable chars that don't work with pretty print
	escapedStr = strings.Replace(escapedStr, "\\", "\\\\", -1)
	escapedStr = strings.Replace(escapedStr, "\r\n", "\\n", -1)
	escapedStr = strings.Replace(escapedStr, "\n", "\\n", -1)
	escapedStr = strings.Replace(escapedStr, "\"", "\\\"", -1)
	return escapedStr
}

// getBase64EncodedGzippedCustomScript will return a base64 of the CSE
func getBase64EncodedGzippedCustomScript(csFilename string, cs *api.ContainerService) string {
	b, err := Asset(csFilename)
	if err != nil {
		// this should never happen and this is a bug
		panic(fmt.Sprintf("BUG: %s", err.Error()))
	}
	// translate the parameters
	templ := template.New("ContainerService template").Funcs(getContainerServiceFuncMap(cs))
	_, err = templ.Parse(string(b))
	if err != nil {
		// this should never happen and this is a bug
		panic(fmt.Sprintf("BUG: %s", err.Error()))
	}
	var buffer bytes.Buffer
	_ = templ.Execute(&buffer, cs)
	csStr := buffer.String()
	csStr = strings.Replace(csStr, "\r\n", "\n", -1)
	return getBase64EncodedGzippedCustomScriptFromStr(csStr)
}

func getStringFromBase64(str string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(str)
	return string(decodedBytes), err
}

// getBase64EncodedGzippedCustomScriptFromStr will return a base64-encoded string of the gzip'd source data
func getBase64EncodedGzippedCustomScriptFromStr(str string) string {
	var gzipB bytes.Buffer
	w := gzip.NewWriter(&gzipB)
	_, _ = w.Write([]byte(str))
	w.Close()
	return base64.StdEncoding.EncodeToString(gzipB.Bytes())
}

func getComponentFuncMap(component api.KubernetesComponent, cs *api.ContainerService) template.FuncMap {
	ret := template.FuncMap{
		"ContainerImage": func(name string) string {
			if i := component.GetContainersIndexByName(name); i > -1 {
				return component.Containers[i].Image
			}
			return ""
		},
		"ContainerCPUReqs": func(name string) string {
			if i := component.GetContainersIndexByName(name); i > -1 {
				return component.Containers[i].CPURequests
			}
			return ""
		},
		"ContainerCPULimits": func(name string) string {
			if i := component.GetContainersIndexByName(name); i > -1 {
				return component.Containers[i].CPULimits
			}
			return ""
		},
		"ContainerMemReqs": func(name string) string {
			if i := component.GetContainersIndexByName(name); i > -1 {
				return component.Containers[i].MemoryRequests
			}
			return ""
		},
		"ContainerMemLimits": func(name string) string {
			if i := component.GetContainersIndexByName(name); i > -1 {
				return component.Containers[i].MemoryLimits
			}
			return ""
		},
		"ContainerConfig": func(name string) string {
			return component.Config[name]
		},
		"IsCustomCloudProfile": func() bool {
			return cs.Properties.IsCustomCloudProfile()
		},
		"IsAzureStackCloud": func() bool {
			return cs.Properties.IsAzureStackCloud()
		},
		"IsKubernetesVersionGe": func(version string) bool {
			return cs.Properties.OrchestratorProfile.IsKubernetes() && common.IsKubernetesVersionGe(cs.Properties.OrchestratorProfile.OrchestratorVersion, version)
		},
	}
	if component.Name == common.APIServerComponentName {
		ret["GetAPIServerArgs"] = func() string {
			return common.GetOrderedEscapedKeyValsString(cs.Properties.OrchestratorProfile.KubernetesConfig.APIServerConfig)
		}
	}
	if component.Name == common.ControllerManagerComponentName {
		ret["GetControllerManagerArgs"] = func() string {
			return common.GetOrderedEscapedKeyValsString(cs.Properties.OrchestratorProfile.KubernetesConfig.ControllerManagerConfig)
		}
	}
	if component.Name == common.SchedulerComponentName {
		ret["GetSchedulerArgs"] = func() string {
			return common.GetOrderedEscapedKeyValsString(cs.Properties.OrchestratorProfile.KubernetesConfig.SchedulerConfig)
		}
	}
	if component.Name == common.CloudControllerManagerComponentName {
		ret["GetCloudControllerManagerArgs"] = func() string {
			return common.GetOrderedEscapedKeyValsString(cs.Properties.OrchestratorProfile.KubernetesConfig.CloudControllerManagerConfig)
		}
	}
	return ret
}

func getAddonFuncMap(addon api.KubernetesAddon, cs *api.ContainerService) template.FuncMap {
	return template.FuncMap{
		"ContainerImage": func(name string) string {
			i := addon.GetAddonContainersIndexByName(name)
			return addon.Containers[i].Image
		},
		"ContainerCPUReqs": func(name string) string {
			i := addon.GetAddonContainersIndexByName(name)
			return addon.Containers[i].CPURequests
		},
		"ContainerCPULimits": func(name string) string {
			i := addon.GetAddonContainersIndexByName(name)
			return addon.Containers[i].CPULimits
		},
		"ContainerMemReqs": func(name string) string {
			i := addon.GetAddonContainersIndexByName(name)
			return addon.Containers[i].MemoryRequests
		},
		"ContainerMemLimits": func(name string) string {
			i := addon.GetAddonContainersIndexByName(name)
			return addon.Containers[i].MemoryLimits
		},
		"ContainerConfig": func(name string) string {
			return addon.Config[name]
		},
		"HasWindows": func() bool {
			return cs.Properties.HasWindows()
		},
		"IsCustomCloudProfile": func() bool {
			return cs.Properties.IsCustomCloudProfile()
		},
		"HasLinux": func() bool {
			return cs.Properties.AnyAgentIsLinux()
		},
		"IsAzureStackCloud": func() bool {
			return cs.Properties.IsAzureStackCloud()
		},
		"NeedsStorageAccountStorageClasses": func() bool {
			return len(cs.Properties.AgentPoolProfiles) > 0 && cs.Properties.AgentPoolProfiles[0].StorageProfile == api.StorageAccount
		},
		"NeedsManagedDiskStorageClasses": func() bool {
			return len(cs.Properties.AgentPoolProfiles) > 0 && cs.Properties.AgentPoolProfiles[0].StorageProfile == api.ManagedDisks
		},
		"UsesCloudControllerManager": func() bool {
			return to.Bool(cs.Properties.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager)
		},
		"HasAvailabilityZones": func() bool {
			return cs.Properties.HasAvailabilityZones()
		},
		"GetZones": func() string {
			if len(cs.Properties.AgentPoolProfiles) == 0 {
				return ""
			}

			var zones string
			for _, zone := range cs.Properties.AgentPoolProfiles[0].AvailabilityZones {
				zones += fmt.Sprintf("\n    - %s-%s", cs.Location, zone)
			}
			return zones
		},
		"CSIControllerReplicas": func() string {
			replicas := "2"
			if cs.Properties.HasWindows() && !cs.Properties.AnyAgentIsLinux() {
				replicas = "1"
			}
			return replicas
		},
		"ShouldEnableCSISnapshotFeature": func(csiDriverName string) bool {
			// Snapshot is not available for Windows clusters
			if cs.Properties.HasWindows() && !cs.Properties.AnyAgentIsLinux() {
				return false
			}

			switch csiDriverName {
			case common.AzureDiskCSIDriverAddonName:
				// Snapshot feature for Azure Disk CSI Driver is in beta, requiring K8s 1.17+
				return common.IsKubernetesVersionGe(cs.Properties.OrchestratorProfile.OrchestratorVersion, "1.17.0")
			case common.AzureFileCSIDriverAddonName:
				// Snapshot feature for Azure File CSI Driver is in alpha, requiring K8s 1.13-1.16
				return common.IsKubernetesVersionGe(cs.Properties.OrchestratorProfile.OrchestratorVersion, "1.13.0") &&
					!common.IsKubernetesVersionGe(cs.Properties.OrchestratorProfile.OrchestratorVersion, "1.17.0")
			}
			return false
		},
		"IsKubernetesVersionGe": func(version string) bool {
			return common.IsKubernetesVersionGe(cs.Properties.OrchestratorProfile.OrchestratorVersion, version)
		},
		"GetAADPodIdentityTaintKey": func() string {
			return common.AADPodIdentityTaintKey
		},
		"GetMode": func() string {
			return addon.Mode
		},
	}
}

func getClusterAutoscalerAddonFuncMap(addon api.KubernetesAddon, cs *api.ContainerService) template.FuncMap {
	return template.FuncMap{
		"ContainerImage": func(name string) string {
			i := addon.GetAddonContainersIndexByName(name)
			return addon.Containers[i].Image
		},

		"ContainerCPUReqs": func(name string) string {
			i := addon.GetAddonContainersIndexByName(name)
			return addon.Containers[i].CPURequests
		},

		"ContainerCPULimits": func(name string) string {
			i := addon.GetAddonContainersIndexByName(name)
			return addon.Containers[i].CPULimits
		},

		"ContainerMemReqs": func(name string) string {
			i := addon.GetAddonContainersIndexByName(name)
			return addon.Containers[i].MemoryRequests
		},

		"ContainerMemLimits": func(name string) string {
			i := addon.GetAddonContainersIndexByName(name)
			return addon.Containers[i].MemoryLimits
		},
		"ContainerConfig": func(name string) string {
			return addon.Config[name]
		},
		"GetMode": func() string {
			return addon.Mode
		},
		"GetClusterAutoscalerNodesConfig": func() string {
			return api.GetClusterAutoscalerNodesConfig(addon, cs)
		},
		"GetBase64EncodedVMType": func() string {
			return base64.StdEncoding.EncodeToString([]byte(cs.Properties.GetVMType()))
		},
		"GetVolumeMounts": func() string {
			if cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity {
				return "\n        - mountPath: /var/lib/waagent/\n          name: waagent\n          readOnly: true"
			}
			return ""
		},
		"GetVolumes": func() string {
			if cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity {
				return "\n      - hostPath:\n          path: /var/lib/waagent/\n        name: waagent"
			}
			return ""
		},
		"GetHostNetwork": func() string {
			if cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity {
				return "\n      hostNetwork: true"
			}
			return ""
		},
		"GetCloud": func() string {
			cloudSpecConfig := cs.GetCloudSpecConfig()
			return cloudSpecConfig.CloudName
		},
		"UseManagedIdentity": func() string {
			if cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity {
				return "true"
			}
			return "false"
		},
		"IsKubernetesVersionGe": func(version string) bool {
			return common.IsKubernetesVersionGe(cs.Properties.OrchestratorProfile.OrchestratorVersion, version)
		},
	}
}

func getComponentsString(cs *api.ContainerService, sourcePath string) string {
	properties := cs.Properties
	var result string
	settingsMap := kubernetesComponentSettingsInit(properties)

	var componentNames []string

	for componentName := range settingsMap {
		componentNames = append(componentNames, componentName)
	}

	sort.Strings(componentNames)

	for _, componentName := range componentNames {
		setting := settingsMap[componentName]
		if component, isEnabled := cs.Properties.OrchestratorProfile.KubernetesConfig.IsComponentEnabled(componentName); isEnabled {
			var input string
			if setting.base64Data != "" {
				var err error
				input, err = getStringFromBase64(setting.base64Data)
				if err != nil {
					return ""
				}
			} else if setting.sourceFile != "" {
				orchProfile := properties.OrchestratorProfile
				versions := strings.Split(orchProfile.OrchestratorVersion, ".")
				templ := template.New("component resolver template").Funcs(getComponentFuncMap(component, cs))
				componentFile := getCustomDataFilePath(setting.sourceFile, sourcePath, versions[0]+"."+versions[1])
				componentFileBytes, err := Asset(componentFile)
				if err != nil {
					return ""
				}
				_, err = templ.Parse(string(componentFileBytes))
				if err != nil {
					return ""
				}
				var buffer bytes.Buffer
				_ = templ.Execute(&buffer, component)
				input = buffer.String()
			}
			if componentName == common.ClusterInitComponentName {
				result += getComponentString(input, "/opt/azure/containers", setting.destinationFile)
			} else {
				result += getComponentString(input, "/etc/kubernetes/manifests", setting.destinationFile)
			}
		}
	}
	return result
}

func getAddonsString(cs *api.ContainerService, sourcePath string) string {
	properties := cs.Properties
	var result string
	settingsMap := kubernetesAddonSettingsInit(properties)

	var addonNames []string

	for addonName := range settingsMap {
		addonNames = append(addonNames, addonName)
	}

	sort.Strings(addonNames)

	for _, addonName := range addonNames {
		setting := settingsMap[addonName]
		if cs.Properties.OrchestratorProfile.KubernetesConfig.IsAddonEnabled(addonName) {
			var input string
			if setting.base64Data != "" {
				var err error
				input, err = getStringFromBase64(setting.base64Data)
				if err != nil {
					return ""
				}
			} else {
				orchProfile := properties.OrchestratorProfile
				versions := strings.Split(orchProfile.OrchestratorVersion, ".")
				addon := orchProfile.KubernetesConfig.GetAddonByName(addonName)
				var templ *template.Template
				switch addonName {
				case "cluster-autoscaler":
					templ = template.New("addon resolver template").Funcs(getClusterAutoscalerAddonFuncMap(addon, cs))
				default:
					templ = template.New("addon resolver template").Funcs(getAddonFuncMap(addon, cs))
				}
				addonFile := getCustomDataFilePath(setting.sourceFile, sourcePath, versions[0]+"."+versions[1])
				addonFileBytes, err := Asset(addonFile)
				if err != nil {
					return ""
				}
				_, err = templ.Parse(string(addonFileBytes))
				if err != nil {
					return ""
				}
				var buffer bytes.Buffer
				_ = templ.Execute(&buffer, addon)
				input = buffer.String()
			}
			result += getComponentString(input, "/etc/kubernetes/addons", setting.destinationFile)
		}
	}
	return result
}

func getDCOSMasterProvisionScript(orchProfile *api.OrchestratorProfile, bootstrapIP string) string {
	scriptname := dcos2Provision
	if orchProfile.DcosConfig == nil || orchProfile.DcosConfig.BootstrapProfile == nil {
		scriptname = dcosProvision
	}

	// add the provision script
	bp, err := Asset(scriptname)
	if err != nil {
		panic(fmt.Sprintf("BUG: %s", err.Error()))
	}

	provisionScript := string(bp)
	if strings.Contains(provisionScript, "'") {
		panic(fmt.Sprintf("BUG: %s may not contain character '", scriptname))
	}

	// the embedded roleFileContents
	roleFileContents := `touch /etc/mesosphere/roles/master
touch /etc/mesosphere/roles/azure_master`
	provisionScript = strings.Replace(provisionScript, "ROLESFILECONTENTS", roleFileContents, -1)
	provisionScript = strings.Replace(provisionScript, "BOOTSTRAP_IP", bootstrapIP, -1)

	var b bytes.Buffer
	b.WriteString(provisionScript)
	b.WriteString("\n")

	return strings.Replace(strings.Replace(b.String(), "\r\n", "\n", -1), "\n", "\n\n    ", -1)
}

func buildYamlFileWithWriteFiles(files []string, cs *api.ContainerService) string {
	clusterYamlFile := `#cloud-config

write_files:
%s
`
	writeFileBlock := ` -  encoding: gzip
    content: !!binary |
        %s
    path: /opt/azure/containers/%s
    permissions: "0744"
`

	filelines := ""
	for _, file := range files {
		b64GzipString := getBase64EncodedGzippedCustomScript(file, cs)
		fileNoPath := strings.TrimPrefix(file, "swarm/")
		filelines += fmt.Sprintf(writeFileBlock, b64GzipString, fileNoPath)
	}
	return fmt.Sprintf(clusterYamlFile, filelines)
}

func getKubernetesSubnets(properties *api.Properties) string {
	subnetString := `{
            "name": "podCIDR%d",
            "properties": {
              "addressPrefix": "10.244.%d.0/24",
              "networkSecurityGroup": {
                "id": "[variables('nsgID')]"
              },
              "routeTable": {
                "id": "[variables('routeTableID')]"
              }
            }
          }`
	var buf bytes.Buffer

	cidrIndex := getKubernetesPodStartIndex(properties)
	for _, agentProfile := range properties.AgentPoolProfiles {
		if agentProfile.OSType == api.Windows {
			for i := 0; i < agentProfile.Count; i++ {
				buf.WriteString(",\n")
				buf.WriteString(fmt.Sprintf(subnetString, cidrIndex, cidrIndex))
				cidrIndex++
			}
		}
	}
	return buf.String()
}

func getKubernetesPodStartIndex(properties *api.Properties) int {
	nodeCount := 0
	nodeCount += properties.MasterProfile.Count
	for _, agentProfile := range properties.AgentPoolProfiles {
		if agentProfile.OSType != api.Windows {
			nodeCount += agentProfile.Count
		}
	}

	return nodeCount + 1
}

func getMasterLinkedTemplateText(orchestratorType string, extensionProfile *api.ExtensionProfile, singleOrAll string) (string, error) {
	extTargetVMNamePrefix := "variables('masterVMNamePrefix')"

	loopCount := "[variables('masterCount')]"
	loopOffset := ""
	if orchestratorType == api.Kubernetes {
		// Due to upgrade k8s sometimes needs to install just some of the nodes.
		loopCount = "[sub(variables('masterCount'), variables('masterOffset'))]"
		loopOffset = "variables('masterOffset')"
	}

	if strings.EqualFold(singleOrAll, "single") {
		loopCount = "1"
	}
	return internalGetPoolLinkedTemplateText(extTargetVMNamePrefix, orchestratorType, loopCount,
		loopOffset, extensionProfile)
}

func getAgentPoolLinkedTemplateText(agentPoolProfile *api.AgentPoolProfile, orchestratorType string, extensionProfile *api.ExtensionProfile, singleOrAll string) (string, error) {
	extTargetVMNamePrefix := fmt.Sprintf("variables('%sVMNamePrefix')", agentPoolProfile.Name)
	loopCount := fmt.Sprintf("[variables('%sCount'))]", agentPoolProfile.Name)
	loopOffset := ""

	// Availability sets can have an offset since we don't redeploy vms.
	// So we don't want to rerun these extensions in scale up scenarios.
	if agentPoolProfile.IsAvailabilitySets() {
		loopCount = fmt.Sprintf("[sub(variables('%sCount'), variables('%sOffset'))]",
			agentPoolProfile.Name, agentPoolProfile.Name)
		loopOffset = fmt.Sprintf("variables('%sOffset')", agentPoolProfile.Name)
	}

	if strings.EqualFold(singleOrAll, "single") {
		loopCount = "1"
	}

	return internalGetPoolLinkedTemplateText(extTargetVMNamePrefix, orchestratorType, loopCount,
		loopOffset, extensionProfile)
}

func internalGetPoolLinkedTemplateText(extTargetVMNamePrefix, orchestratorType, loopCount, loopOffset string, extensionProfile *api.ExtensionProfile) (string, error) {
	dta, e := getLinkedTemplateTextForURL(extensionProfile.RootURL, orchestratorType, extensionProfile.Name, extensionProfile.Version, extensionProfile.URLQuery)
	if e != nil {
		return "", e
	}
	if strings.Contains(extTargetVMNamePrefix, "master") {
		dta = strings.Replace(dta, "EXTENSION_TARGET_VM_TYPE", "master", -1)
	} else {
		dta = strings.Replace(dta, "EXTENSION_TARGET_VM_TYPE", "agent", -1)
	}
	extensionsParameterReference := fmt.Sprintf("[parameters('%sParameters')]", extensionProfile.Name)
	dta = strings.Replace(dta, "EXTENSION_PARAMETERS_REPLACE", extensionsParameterReference, -1)
	dta = strings.Replace(dta, "EXTENSION_URL_REPLACE", extensionProfile.RootURL, -1)
	dta = strings.Replace(dta, "EXTENSION_TARGET_VM_NAME_PREFIX", extTargetVMNamePrefix, -1)
	if _, err := strconv.Atoi(loopCount); err == nil {
		dta = strings.Replace(dta, "\"EXTENSION_LOOP_COUNT\"", loopCount, -1)
	} else {
		dta = strings.Replace(dta, "EXTENSION_LOOP_COUNT", loopCount, -1)
	}

	dta = strings.Replace(dta, "EXTENSION_LOOP_OFFSET", loopOffset, -1)
	return dta, nil
}

func validateProfileOptedForExtension(extensionName string, profileExtensions []api.Extension) (bool, string) {
	for _, extension := range profileExtensions {
		if extensionName == extension.Name {
			return true, extension.SingleOrAll
		}
	}
	return false, ""
}

// getLinkedTemplateTextForURL returns the string data from
// template-link.json in the following directory:
// extensionsRootURL/extensions/extensionName/version
// It returns an error if the extension cannot be found
// or loaded.  getLinkedTemplateTextForURL provides the ability
// to pass a root extensions url for testing
func getLinkedTemplateTextForURL(rootURL, orchestrator, extensionName, version, query string) (string, error) {
	supportsExtension, err := orchestratorSupportsExtension(rootURL, orchestrator, extensionName, version, query)
	if !supportsExtension {
		return "", errors.Wrap(err, "Extension not supported for orchestrator")
	}

	templateLinkBytes, err := getExtensionResource(rootURL, extensionName, version, "template-link.json", query)
	if err != nil {
		return "", err
	}

	return string(templateLinkBytes), nil
}

func orchestratorSupportsExtension(rootURL, orchestrator, extensionName, version, query string) (bool, error) {
	orchestratorBytes, err := getExtensionResource(rootURL, extensionName, version, "supported-orchestrators.json", query)
	if err != nil {
		return false, err
	}

	var supportedOrchestrators []string
	err = json.Unmarshal(orchestratorBytes, &supportedOrchestrators)
	if err != nil {
		return false, errors.Errorf("Unable to parse supported-orchestrators.json for Extension %s Version %s", extensionName, version)
	}

	if !stringInSlice(orchestrator, supportedOrchestrators) {
		return false, errors.Errorf("Orchestrator: %s not in list of supported orchestrators for Extension: %s Version %s", orchestrator, extensionName, version)
	}

	return true, nil
}

func getExtensionResource(rootURL, extensionName, version, fileName, query string) ([]byte, error) {
	requestURL := getExtensionURL(rootURL, extensionName, version, fileName, query)

	res, err := http.Get(requestURL)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to GET extension resource for extension: %s with version %s with filename %s at URL: %s", extensionName, version, fileName, requestURL)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.Errorf("Unable to GET extension resource for extension: %s with version %s with filename %s at URL: %s StatusCode: %s: Status: %s", extensionName, version, fileName, requestURL, strconv.Itoa(res.StatusCode), res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to GET extension resource for extension: %s with version %s  with filename %s at URL: %s", extensionName, version, fileName, requestURL)
	}

	return body, nil
}

func getExtensionURL(rootURL, extensionName, version, fileName, query string) string {
	extensionsDir := "extensions"
	url := rootURL + extensionsDir + "/" + extensionName + "/" + version + "/" + fileName
	if query != "" {
		url += "?" + query
	}
	return url
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getSwarmVersions(orchestratorVersion, dockerComposeVersion string) string {
	return fmt.Sprintf("\"orchestratorVersion\": \"%s\",\n\"dockerComposeVersion\": \"%s\",\n", orchestratorVersion, dockerComposeVersion)
}

func wrapAsVariableObject(o, v string) string {
	return fmt.Sprintf("',variables('%s').%s,'", o, v)
}

func getSSHPublicKeysPowerShell(linuxProfile *api.LinuxProfile) string {
	str := ""
	if linuxProfile != nil {
		lastItem := len(linuxProfile.SSH.PublicKeys) - 1
		for i, publicKey := range linuxProfile.SSH.PublicKeys {
			str += `"` + strings.TrimSpace(publicKey.KeyData) + `"`
			if i < lastItem {
				str += ", "
			}
		}
	}
	return str
}

func getWindowsMasterSubnetARMParam(masterProfile *api.MasterProfile) string {
	if masterProfile != nil && masterProfile.IsCustomVNET() {
		return "',parameters('vnetCidr'),'"
	}
	return "',parameters('masterSubnet'),'"
}
