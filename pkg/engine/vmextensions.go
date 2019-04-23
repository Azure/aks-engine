// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/pkg/errors"
)

func CreateAKSBillingExtension(cs *api.ContainerService) VirtualMachineExtensionARM {
	location := "[variables('location')]"
	name := "[concat(variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')), '/computeAksLinuxBilling')]"
	dependentVM := ""
	dependentVM = "[concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')))]"

	return VirtualMachineExtensionARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionCompute')]",
			Copy: map[string]string{
				"count": "[sub(variables('masterCount'), variables('masterOffset'))]",
				"name":  "vmLoopNode",
			},
			DependsOn: []string{
				dependentVM,
			},
		},
		VirtualMachineExtension: compute.VirtualMachineExtension{
			Location: to.StringPtr(location),
			Name:     to.StringPtr(name),
			VirtualMachineExtensionProperties: &compute.VirtualMachineExtensionProperties{
				Publisher:               to.StringPtr("Microsoft.AKS"),
				Type:                    to.StringPtr("Compute.AKS-Engine.Linux.Billing"),
				TypeHandlerVersion:      to.StringPtr("1.0"),
				AutoUpgradeMinorVersion: to.BoolPtr(true),
				Settings:                &map[string]interface{}{},
			},
			Type: to.StringPtr("Microsoft.Compute/virtualMachines/extensions"),
			Tags: map[string]*string{},
		},
	}
}

func CreateCustomScriptExtension(cs *api.ContainerService) VirtualMachineExtensionARM {
	location := "[variables('location')]"
	name := "[concat(variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')),'/cse', '-master-', copyIndex(variables('masterOffset')))]"
	outBoundCmd := ""
	registry := ""
	ncBinary := "nc"
	if cs.Properties.MasterProfile != nil && cs.Properties.MasterProfile.IsCoreOS() {
		ncBinary = "ncat"
	}
	if !cs.Properties.FeatureFlags.IsFeatureEnabled("BlockOutboundInternet") {
		if cs.GetCloudSpecConfig().CloudName == api.AzureChinaCloud {
			registry = `gcr.azk8s.cn 80`
		} else {
			registry = `k8s.gcr.io 443 && retrycmd_if_failure 50 1 3 ` + ncBinary + ` -vz gcr.io 443 && retrycmd_if_failure 50 1 3 ` + ncBinary + ` -vz docker.io 443`
		}
		outBoundCmd = `ERR_OUTBOUND_CONN_FAIL=50; retrycmd_if_failure 50 1 3 ` + ncBinary + ` -vz ` + registry + ` || exit $ERR_OUTBOUND_CONN_FAIL;`
	}
	vmExtension := compute.VirtualMachineExtension{
		Location: to.StringPtr(location),
		Name:     to.StringPtr(name),
		VirtualMachineExtensionProperties: &compute.VirtualMachineExtensionProperties{
			Publisher:               to.StringPtr("Microsoft.Azure.Extensions"),
			Type:                    to.StringPtr("CustomScript"),
			TypeHandlerVersion:      to.StringPtr("2.0"),
			AutoUpgradeMinorVersion: to.BoolPtr(true),
			Settings:                &map[string]interface{}{},
			ProtectedSettings: &map[string]interface{}{
				"commandToExecute": "[concat('retrycmd_if_failure() { r=$1; w=$2; t=$3; shift && shift && shift; for i in $(seq 1 $r); do timeout $t ${@}; [ $? -eq 0  ] && break || if [ $i -eq $r ]; then return 1; else sleep $w; fi; done }; " + outBoundCmd + " for i in $(seq 1 1200); do if [ -f /opt/azure/containers/provision.sh ]; then break; fi; if [ $i -eq 1200 ]; then exit 100; else sleep 1; fi; done; ', variables('provisionScriptParametersCommon'),' ',variables('provisionScriptParametersMaster'), ' /usr/bin/nohup /bin/bash -c \"/bin/bash /opt/azure/containers/provision.sh >> /var/log/azure/cluster-provision.log 2>&1\"')]",
			},
		},
		Type: to.StringPtr("Microsoft.Compute/virtualMachines/extensions"),
		Tags: map[string]*string{},
	}
	return VirtualMachineExtensionARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionCompute')]",
			Copy: map[string]string{
				"count": "[sub(variables('masterCount'), variables('masterOffset'))]",
				"name":  "vmLoopNode",
			},
			DependsOn: []string{"[concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')))]"},
		},
		VirtualMachineExtension: vmExtension,
	}
}

func createAgentVMASCustomScriptExtension(cs *api.ContainerService, profile *api.AgentPoolProfile) VirtualMachineExtensionARM {
	location := "[variables('location')]"
	name := fmt.Sprintf("[concat(variables('%[1]sVMNamePrefix'), copyIndex(variables('%[1]sOffset')),'/cse', '-agent-', copyIndex(variables('%[1]sOffset')))]", profile.Name)
	outBoundCmd := ""
	registry := ""
	ncBinary := "nc"
	if profile.IsCoreOS() {
		ncBinary = "ncat"
	}
	if !cs.Properties.FeatureFlags.IsFeatureEnabled("BlockOutboundInternet") {
		if cs.GetCloudSpecConfig().CloudName == api.AzureChinaCloud {
			registry = `gcr.azk8s.cn 80`
		} else {
			registry = `k8s.gcr.io 443 && retrycmd_if_failure 50 1 3 ` + ncBinary + ` -vz gcr.io 443 && retrycmd_if_failure 50 1 3 ` + ncBinary + ` -vz docker.io 443`
		}
		outBoundCmd = `ERR_OUTBOUND_CONN_FAIL=50; retrycmd_if_failure 50 1 3 ` + ncBinary + ` -vz ` + registry + ` || exit $ERR_OUTBOUND_CONN_FAIL;`
	}

	runInBackground := ""

	if cs.Properties.FeatureFlags.IsFeatureEnabled("CSERunInBackground") {
		runInBackground = " &"
	}

	nVidiaEnabled := strconv.FormatBool(common.IsNvidiaEnabledSKU(profile.VMSize))
	sgxEnabled := strconv.FormatBool(common.IsSgxEnabledSKU(profile.VMSize))

	vmExtension := compute.VirtualMachineExtension{
		Location: to.StringPtr(location),
		Name:     to.StringPtr(name),
		VirtualMachineExtensionProperties: &compute.VirtualMachineExtensionProperties{
			AutoUpgradeMinorVersion: to.BoolPtr(true),
			Settings:                &map[string]interface{}{},
		},
		Type: to.StringPtr("Microsoft.Compute/virtualMachines/extensions"),
	}

	if profile.IsWindows() {
		vmExtension.Publisher = to.StringPtr("Microsoft.Compute")
		vmExtension.VirtualMachineExtensionProperties.Type = to.StringPtr("CustomScriptExtension")
		vmExtension.TypeHandlerVersion = to.StringPtr("1.8")
		vmExtension.ProtectedSettings = &map[string]interface{}{
			"commandToExecute": "[concat('powershell.exe -ExecutionPolicy Unrestricted -command \"', '$arguments = ', variables('singleQuote'),'-MasterIP ',variables('kubernetesAPIServerIP'),' -KubeDnsServiceIp ',parameters('kubeDnsServiceIp'),' -MasterFQDNPrefix ',variables('masterFqdnPrefix'),' -Location ',variables('location'),' -AgentKey ',parameters('clientPrivateKey'),' -AADClientId ',variables('servicePrincipalClientId'),' -AADClientSecret ',variables('singleQuote'),variables('singleQuote'),variables('servicePrincipalClientSecret'),variables('singleQuote'),variables('singleQuote'), ' ',variables('singleQuote'), ' ; ', variables('windowsCustomScriptSuffix'), '\" > %SYSTEMDRIVE%\\AzureData\\CustomDataSetupScript.log 2>&1')]",
		}
	} else {
		vmExtension.Publisher = to.StringPtr("Microsoft.Azure.Extensions")
		vmExtension.VirtualMachineExtensionProperties.Type = to.StringPtr("CustomScript")
		vmExtension.TypeHandlerVersion = to.StringPtr("2.0")
		commandExec := fmt.Sprintf("[concat('retrycmd_if_failure() { r=$1; w=$2; t=$3; shift && shift && shift; for i in $(seq 1 $r); do timeout $t ${@}; [ $? -eq 0  ] && break || if [ $i -eq $r ]; then return 1; else sleep $w; fi; done }; %s for i in $(seq 1 1200); do if [ -f /opt/azure/containers/provision.sh ]; then break; fi; if [ $i -eq 1200 ]; then exit 100; else sleep 1; fi; done; ', variables('provisionScriptParametersCommon'),' GPU_NODE=%s SGX_NODE=%s /usr/bin/nohup /bin/bash -c \"/bin/bash /opt/azure/containers/provision.sh >> /var/log/azure/cluster-provision.log 2>&1%s\"')]", outBoundCmd, nVidiaEnabled, sgxEnabled, runInBackground)
		vmExtension.ProtectedSettings = &map[string]interface{}{
			"commandToExecute": commandExec,
		}
	}

	dependency := fmt.Sprintf("[concat('Microsoft.Compute/virtualMachines/', variables('%[1]sVMNamePrefix'), copyIndex(variables('%[1]sOffset')))]", profile.Name)

	return VirtualMachineExtensionARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionCompute')]",
			Copy: map[string]string{
				"count": fmt.Sprintf("[sub(variables('%[1]sCount'), variables('%[1]sOffset'))]", profile.Name),
				"name":  "vmLoopNode",
			},
			DependsOn: []string{dependency},
		},
		VirtualMachineExtension: vmExtension,
	}
}

func CreateAgentVMASAKSBillingExtension(cs *api.ContainerService, profile *api.AgentPoolProfile) VirtualMachineExtensionARM {
	location := "[variables('location')]"
	name := fmt.Sprintf("[concat(variables('%[1]sVMNamePrefix'), copyIndex(variables('%[1]sOffset')), '/computeAksLinuxBilling')]", profile.Name)
	dependentVM := fmt.Sprintf("[concat('Microsoft.Compute/virtualMachines/', variables('%[1]sVMNamePrefix'), copyIndex(variables('%[1]sOffset')))]", profile.Name)

	vmExtension := compute.VirtualMachineExtension{
		Location: to.StringPtr(location),
		Name:     to.StringPtr(name),
		VirtualMachineExtensionProperties: &compute.VirtualMachineExtensionProperties{
			Publisher:               to.StringPtr("Microsoft.AKS"),
			TypeHandlerVersion:      to.StringPtr("1.0"),
			AutoUpgradeMinorVersion: to.BoolPtr(true),
			Settings:                &map[string]interface{}{},
		},
		Type: to.StringPtr("Microsoft.Compute/virtualMachines/extensions"),
	}

	if cs.Properties.IsHostedMasterProfile() {
		if profile.IsWindows() {
			vmExtension.VirtualMachineExtensionProperties.Type = to.StringPtr("Compute.AKS.Windows.Billing")
		} else {
			vmExtension.VirtualMachineExtensionProperties.Type = to.StringPtr("Compute.AKS.Linux.Billing")
		}
	} else {
		if profile.IsWindows() {
			vmExtension.VirtualMachineExtensionProperties.Type = to.StringPtr("Compute.AKS-Engine.Windows.Billing")
		} else {
			vmExtension.VirtualMachineExtensionProperties.Type = to.StringPtr("Compute.AKS-Engine.Linux.Billing")
		}
	}

	return VirtualMachineExtensionARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionCompute')]",
			Copy: map[string]string{
				"count": fmt.Sprintf("[sub(variables('%[1]sCount'), variables('%[1]sOffset'))]", profile.Name),
				"name":  "vmLoopNode",
			},
			DependsOn: []string{
				dependentVM,
			},
		},
		VirtualMachineExtension: vmExtension,
	}
}

// CreateLinkedTemplatesForExtensions returns the
// Microsoft.Resources/deployments for each extension
func CreateLinkedTemplatesForExtensions(properties *api.Properties) string {
	var result string

	extensions := properties.ExtensionProfiles
	masterProfileExtensions := properties.MasterProfile.Extensions
	orchestratorType := properties.OrchestratorProfile.OrchestratorType

	for err, extensionProfile := range extensions {
		_ = err

		masterOptedForExtension, singleOrAll := validateProfileOptedForExtension(extensionProfile.Name, masterProfileExtensions)
		if masterOptedForExtension {
			result += ","
			dta, e := getMasterLinkedTemplateText(properties.MasterProfile, orchestratorType, extensionProfile, singleOrAll)
			if e != nil {
				fmt.Println(e.Error())
				return ""
			}
			result += dta
		}

		for _, agentPoolProfile := range properties.AgentPoolProfiles {
			poolProfileExtensions := agentPoolProfile.Extensions
			poolOptedForExtension, singleOrAll := validateProfileOptedForExtension(extensionProfile.Name, poolProfileExtensions)
			if poolOptedForExtension {
				result += ","
				dta, e := getAgentPoolLinkedTemplateText(agentPoolProfile, orchestratorType, extensionProfile, singleOrAll)
				if e != nil {
					fmt.Println(e.Error())
					return ""
				}
				result += dta
			}

		}
	}

	return result
}

func getMasterLinkedTemplateText(masterProfile *api.MasterProfile, orchestratorType string, extensionProfile *api.ExtensionProfile, singleOrAll string) (string, error) {
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
