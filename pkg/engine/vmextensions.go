// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
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
	var userAssignedIDEnabled bool
	if cs.Properties.OrchestratorProfile != nil && cs.Properties.OrchestratorProfile.KubernetesConfig != nil {
		userAssignedIDEnabled = cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedIDEnabled()
	} else {
		userAssignedIDEnabled = false
	}
	isVHD := "false"
	if cs.Properties.MasterProfile != nil {
		isVHD = strconv.FormatBool(cs.Properties.MasterProfile.IsVHDDistro())
	}

	// TODO The AzureStack constraint has to be relaxed, it should only apply to *disconnected* instances
	if !cs.Properties.FeatureFlags.IsFeatureEnabled("BlockOutboundInternet") && !cs.Properties.IsAzureStackCloud() && cs.Properties.IsHostedMasterProfile() {
		if cs.GetCloudSpecConfig().CloudName == api.AzureChinaCloud {
			registry = `gcr.azk8s.cn 443`
		} else {
			registry = `mcr.microsoft.com 443`
		}
		outBoundCmd = `retrycmd_if_failure() { r=$1; w=$2; t=$3; shift && shift && shift; for i in $(seq 1 $r); do timeout $t ${@}; [ $? -eq 0  ] && break || if [ $i -eq $r ]; then return 1; else sleep $w; fi; done }; ERR_OUTBOUND_CONN_FAIL=50; retrycmd_if_failure 50 1 3 ` + ncBinary + ` -vz ` + registry + ` || exit $ERR_OUTBOUND_CONN_FAIL;`
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
				"commandToExecute": fmt.Sprintf("[concat('echo $(date),$(hostname); "+outBoundCmd+" for i in $(seq 1 1200); do grep -Fq \"EOF\" /opt/azure/containers/provision.sh && break; if [ $i -eq 1200 ]; then exit 100; else sleep 1; fi; done; ', variables('provisionScriptParametersCommon'),%s,variables('provisionScriptParametersMaster'), ' IS_VHD=%s /usr/bin/nohup /bin/bash -c \"/bin/bash /opt/azure/containers/provision.sh >> /var/log/azure/cluster-provision.log 2>&1\"')]", generateUserAssignedIdentityClientIDParameter(userAssignedIDEnabled), isVHD),
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
	var userAssignedIDEnabled bool
	if cs.Properties.OrchestratorProfile != nil && cs.Properties.OrchestratorProfile.KubernetesConfig != nil {
		userAssignedIDEnabled = cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedIDEnabled()
	} else {
		userAssignedIDEnabled = false
	}

	// TODO The AzureStack constraint has to be relaxed, it should only apply to *disconnected* instances
	if !cs.Properties.FeatureFlags.IsFeatureEnabled("BlockOutboundInternet") && !cs.Properties.IsAzureStackCloud() && cs.Properties.IsHostedMasterProfile() {
		if cs.GetCloudSpecConfig().CloudName == api.AzureChinaCloud {
			registry = `gcr.azk8s.cn 443`
		} else {
			registry = `mcr.microsoft.com 443`
		}
		outBoundCmd = `retrycmd_if_failure() { r=$1; w=$2; t=$3; shift && shift && shift; for i in $(seq 1 $r); do timeout $t ${@}; [ $? -eq 0  ] && break || if [ $i -eq $r ]; then return 1; else sleep $w; fi; done }; ERR_OUTBOUND_CONN_FAIL=50; retrycmd_if_failure 50 1 3 ` + ncBinary + ` -vz ` + registry + ` || exit $ERR_OUTBOUND_CONN_FAIL;`
	}

	runInBackground := ""

	if cs.Properties.FeatureFlags.IsFeatureEnabled("CSERunInBackground") {
		runInBackground = " &"
	}

	nVidiaEnabled := strconv.FormatBool(common.IsNvidiaEnabledSKU(profile.VMSize))
	sgxEnabled := strconv.FormatBool(common.IsSgxEnabledSKU(profile.VMSize))
	auditDEnabled := strconv.FormatBool(to.Bool(profile.AuditDEnabled))
	isVHD := strconv.FormatBool(profile.IsVHDDistro())

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
			"commandToExecute": "[concat('echo %DATE%,%TIME%,%COMPUTERNAME% && powershell.exe -ExecutionPolicy Unrestricted -command \"', '$arguments = ', variables('singleQuote'),'-MasterIP ',variables('kubernetesAPIServerIP'),' -KubeDnsServiceIp ',parameters('kubeDnsServiceIp'),' -MasterFQDNPrefix ',variables('masterFqdnPrefix'),' -Location ',variables('location'),' -TargetEnvironment ',parameters('targetEnvironment'),' -AgentKey ',parameters('clientPrivateKey'),' -AADClientId ',variables('servicePrincipalClientId'),' -AADClientSecret ',variables('singleQuote'),variables('singleQuote'),base64(variables('servicePrincipalClientSecret')),variables('singleQuote'),variables('singleQuote'),' -NetworkAPIVersion ',variables('apiVersionNetwork'),' ',variables('singleQuote'), ' ; ', variables('windowsCustomScriptSuffix'), '\" > %SYSTEMDRIVE%\\AzureData\\CustomDataSetupScript.log 2>&1 ; exit $LASTEXITCODE')]",
		}
	} else {
		vmExtension.Publisher = to.StringPtr("Microsoft.Azure.Extensions")
		vmExtension.VirtualMachineExtensionProperties.Type = to.StringPtr("CustomScript")
		vmExtension.TypeHandlerVersion = to.StringPtr("2.0")
		commandExec := fmt.Sprintf("[concat('echo $(date),$(hostname); %s for i in $(seq 1 1200); do grep -Fq \"EOF\" /opt/azure/containers/provision.sh && break; if [ $i -eq 1200 ]; then exit 100; else sleep 1; fi; done; ', variables('provisionScriptParametersCommon'),%s,' IS_VHD=%s GPU_NODE=%s SGX_NODE=%s AUDITD_ENABLED=%s /usr/bin/nohup /bin/bash -c \"/bin/bash /opt/azure/containers/provision.sh >> /var/log/azure/cluster-provision.log 2>&1%s\"')]", outBoundCmd, generateUserAssignedIdentityClientIDParameter(userAssignedIDEnabled), isVHD, nVidiaEnabled, sgxEnabled, auditDEnabled, runInBackground)
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

// CreateCustomExtensions returns a list of DeploymentARM objects for the custom extensions to be deployed
func CreateCustomExtensions(properties *api.Properties) []DeploymentARM {
	var extensionsARM []DeploymentARM

	if properties.MasterProfile != nil {
		// The first extension needs to depend on the master cse created for all nodes
		// Each proceeding extension needs to depend on the previous one to avoid ARM conflicts in the Compute RP
		nextDependsOn := "[concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')), '/extensions/cse-master-', copyIndex(variables('masterOffset')))]"

		for _, extensionProfile := range properties.ExtensionProfiles {
			masterOptedForExtension, singleOrAll := validateProfileOptedForExtension(extensionProfile.Name, properties.MasterProfile.Extensions)
			if masterOptedForExtension {
				data, e := getMasterLinkedTemplateText(properties.OrchestratorProfile.OrchestratorType, extensionProfile, singleOrAll)
				if e != nil {
					fmt.Println(e.Error())
				}
				var ext DeploymentARM
				e = json.Unmarshal([]byte(data), &ext)
				if e != nil {
					fmt.Println(e.Error())
				}
				ext.DependsOn = []string{nextDependsOn}
				nextDependsOn = *ext.Name
				extensionsARM = append(extensionsARM, ext)
			}
		}
	}

	for _, agentPoolProfile := range properties.AgentPoolProfiles {
		// The first extension needs to depend on the agent cse created for all nodes
		// Each proceeding extension needs to depend on the previous one to avoid ARM conflicts in the Compute RP
		nextDependsOn := fmt.Sprintf("[concat('Microsoft.Compute/virtualMachines/', variables('%[1]sVMNamePrefix'), copyIndex(variables('%[1]sOffset')), '/extensions/cse-agent-', copyIndex(variables('%[1]sOffset')))]", agentPoolProfile.Name)

		for _, extensionProfile := range properties.ExtensionProfiles {
			poolOptedForExtension, singleOrAll := validateProfileOptedForExtension(extensionProfile.Name, agentPoolProfile.Extensions)
			if poolOptedForExtension {
				data, e := getAgentPoolLinkedTemplateText(agentPoolProfile, properties.OrchestratorProfile.OrchestratorType, extensionProfile, singleOrAll)
				if e != nil {
					fmt.Println(e.Error())
				}
				var ext DeploymentARM
				e = json.Unmarshal([]byte(data), &ext)
				if e != nil {
					fmt.Println(e.Error())
				}
				ext.DependsOn = []string{nextDependsOn}
				nextDependsOn = *ext.Name
				extensionsARM = append(extensionsARM, ext)
			}
		}
	}

	return extensionsARM
}
