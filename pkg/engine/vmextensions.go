package engine

import (
	"fmt"
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
	"strconv"
)

func createManagedIdentityExtension(cs api.ContainerService) VirtualMachineExtensionARM {
	dependentVM := "[concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex())]"
	dependentRoleAssignment := ""
	userAssignedIDEnabled := cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity &&
		cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedID != ""

	if userAssignedIDEnabled {
		dependentRoleAssignment = "[concat('Microsoft.Authorization/roleAssignments/',guid(concat(variables('userAssignedID'), 'roleAssignment', resourceGroup().id)))]"
	} else {
		dependentRoleAssignment = "[concat('Microsoft.Authorization/roleAssignments/', guid(concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex(), 'vmidentity')))]"
	}

	return VirtualMachineExtensionARM{
		ARMResource: ARMResource{
			ApiVersion: "[variables('apiVersionCompute')]",
			Copy: map[string]string{
				"count": "[variables('masterCount')]",
				"name":  "vmLoopNode",
			},
			DependsOn: []string{
				dependentVM,
				dependentRoleAssignment,
			},
		},
		VirtualMachineExtension: compute.VirtualMachineExtension{
			Location: to.StringPtr("[resourceGroup().location]"),
			Name:     to.StringPtr("[concat(variables('masterVMNamePrefix'), copyIndex(), '/ManagedIdentityExtension')]"),
			VirtualMachineExtensionProperties: &compute.VirtualMachineExtensionProperties{
				Publisher:               to.StringPtr("Microsoft.ManagedIdentity"),
				Type:                    to.StringPtr("ManagedIdentityExtensionForLinux"),
				TypeHandlerVersion:      to.StringPtr("1.0"),
				AutoUpgradeMinorVersion: to.BoolPtr(true),
				Settings: &map[string]interface{}{
					"port": "50343",
				},
				ProtectedSettings: &map[string]interface{}{},
			},
			Type: to.StringPtr("Microsoft.Compute/virtualMachines/extensions"),
		},
	}
}

func CreateAKSBillingExtension(cs *api.ContainerService) VirtualMachineExtensionARM {
	location := "[variables('location')]"
	name := "[concat(variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')), '/computeAksLinuxBilling')]"
	useManagedIdentity := cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity
	dependentVM := ""
	if useManagedIdentity {
		dependentVM = "[concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')), '/extensions/ManagedIdentityExtension')]"
	} else {
		dependentVM = "[concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')))]"
	}

	return VirtualMachineExtensionARM{
		ARMResource: ARMResource{
			ApiVersion: "[variables('apiVersionCompute')]",
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
	if !cs.Properties.FeatureFlags.IsFeatureEnabled("BlockOutboundInternet") {
		if cs.GetCloudSpecConfig().CloudName == api.AzureChinaCloud {
			registry = `gcr.azk8s.cn 80`
		} else {
			registry = `k8s.gcr.io 443 && retrycmd_if_failure 50 1 3 nc -vz gcr.io 443 && retrycmd_if_failure 50 1 3 nc -vz docker.io 443`
		}
		outBoundCmd = `ERR_OUTBOUND_CONN_FAIL=50; retrycmd_if_failure 50 1 3 nc -vz ` + registry + ` || exit $ERR_OUTBOUND_CONN_FAIL;`
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
			ApiVersion: "[variables('apiVersionCompute')]",
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
	if !cs.Properties.FeatureFlags.IsFeatureEnabled("BlockOutboundInternet") {
		if cs.GetCloudSpecConfig().CloudName == api.AzureChinaCloud {
			registry = `gcr.azk8s.cn 80`
		} else {
			registry = `k8s.gcr.io 443 && retrycmd_if_failure 50 1 3 nc -vz gcr.io 443 && retrycmd_if_failure 50 1 3 nc -vz docker.io 443`
		}
		outBoundCmd = `ERR_OUTBOUND_CONN_FAIL=50; retrycmd_if_failure 50 1 3 nc -vz ` + registry + ` || exit $ERR_OUTBOUND_CONN_FAIL;`
	}

	runInBackground := ""

	if cs.Properties.FeatureFlags.IsFeatureEnabled("CSERunInBackground") {
		runInBackground = "&"
	}

	nVidiaEnabled := strconv.FormatBool(common.IsNvidiaEnabledSKU(profile.VMSize))

	vmExtension := compute.VirtualMachineExtension{
		Location:                          to.StringPtr(location),
		Name:                              to.StringPtr(name),
		VirtualMachineExtensionProperties: &compute.VirtualMachineExtensionProperties{},
		Type:                              to.StringPtr("Microsoft.Compute/virtualMachines/extensions"),
	}

	if profile.IsWindows() {
		vmExtension.Publisher = to.StringPtr("Microsoft.Compute")
		vmExtension.VirtualMachineExtensionProperties.Type = to.StringPtr("CustomScriptExtension")
		vmExtension.TypeHandlerVersion = to.StringPtr("1.8")
		vmExtension.ProtectedSettings = &map[string]interface{}{
			"commandToExecute": "[concat('powershell.exe -ExecutionPolicy Unrestricted -command \"', '$arguments = ', variables('singleQuote'),'-MasterIP ',variables('kubernetesAPIServerIP'),' -KubeDnsServiceIp ',parameters('kubeDnsServiceIp'),' -MasterFQDNPrefix ',variables('masterFqdnPrefix'),' -Location ',variables('location'),' -AgentKey ',parameters('clientPrivateKey'),' -AADClientId ',variables('servicePrincipalClientId'),' -AADClientSecret ',variables('servicePrincipalClientSecret'),variables('singleQuote'), ' ; ', variables('windowsCustomScriptSuffix'), '\" > %SYSTEMDRIVE%\\AzureData\\CustomDataSetupScript.log 2>&1')]",
		}
	} else {
		vmExtension.Publisher = to.StringPtr("Microsoft.Azure.Extensions")
		vmExtension.VirtualMachineExtensionProperties.Type = to.StringPtr("CustomScript")
		vmExtension.TypeHandlerVersion = to.StringPtr("2.0")
		commandExec := fmt.Sprintf("[concat('retrycmd_if_failure() { r=$1; w=$2; t=$3; shift && shift && shift; for i in $(seq 1 $r); do timeout $t ${@}; [ $? -eq 0  ] && break || if [ $i -eq $r ]; then return 1; else sleep $w; fi; done }; %s for i in $(seq 1 1200); do if [ -f /opt/azure/containers/provision.sh ]; then break; fi; if [ $i -eq 1200 ]; then exit 100; else sleep 1; fi; done; ', variables('provisionScriptParametersCommon'),' GPU_NODE=%s /usr/bin/nohup /bin/bash -c \"/bin/bash /opt/azure/containers/provision.sh >> /var/log/azure/cluster-provision.log 2>&1 %s\"')]", outBoundCmd, nVidiaEnabled, runInBackground)
		vmExtension.ProtectedSettings = &map[string]interface{}{
			"commandToExecute": commandExec,
		}
	}

	dependency := fmt.Sprintf("[concat('Microsoft.Compute/virtualMachines/', variables('%[1]sVMNamePrefix'), copyIndex(variables('%[1]sOffset')))]", profile.Name)

	return VirtualMachineExtensionARM{
		ARMResource: ARMResource{
			ApiVersion: "[variables('apiVersionCompute')]",
			Copy: map[string]string{
				"count": "[sub(variables('masterCount'), variables('masterOffset'))]",
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
		Tags: map[string]*string{},
	}

	if profile.IsWindows() {
		vmExtension.VirtualMachineExtensionProperties.Type = to.StringPtr("Compute.AKS-Engine.Windows.Billing")
	} else {
		if cs.Properties.IsHostedMasterProfile() {
			vmExtension.VirtualMachineExtensionProperties.Type = to.StringPtr("Compute.AKS.Linux.Billing")
		} else {
			vmExtension.VirtualMachineExtensionProperties.Type = to.StringPtr("Compute.AKS-Engine.Linux.Billing")
		}
	}

	return VirtualMachineExtensionARM{
		ARMResource: ARMResource{
			ApiVersion: "[variables('apiVersionCompute')]",
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
