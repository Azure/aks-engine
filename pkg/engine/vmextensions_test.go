// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/google/go-cmp/cmp"
)

func TestCreateAKSBillingExtension(t *testing.T) {

	// Test with UseManagedIdentity as true.
	cs := &api.ContainerService{
		Properties: &api.Properties{
			OrchestratorProfile: &api.OrchestratorProfile{
				KubernetesConfig: &api.KubernetesConfig{
					UseManagedIdentity: true,
				},
			},
		},
	}

	vmExtension := CreateAKSBillingExtension(cs)

	expected := VirtualMachineExtensionARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionCompute')]",
			Copy: map[string]string{
				"count": "[sub(variables('masterCount'), variables('masterOffset'))]",
				"name":  "vmLoopNode",
			},
			DependsOn: []string{
				"[concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')))]",
			},
		},

		VirtualMachineExtension: compute.VirtualMachineExtension{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[concat(variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')), '/computeAksLinuxBilling')]"),
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

	diff := cmp.Diff(vmExtension, expected)
	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}

	// Test with UseManagedIdentity set to false.
	cs = &api.ContainerService{
		Properties: &api.Properties{
			OrchestratorProfile: &api.OrchestratorProfile{
				KubernetesConfig: &api.KubernetesConfig{
					UseManagedIdentity: false,
				},
			},
		},
	}

	expected.DependsOn = []string{
		"[concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')))]",
	}

	vmExtension = CreateAKSBillingExtension(cs)

	diff = cmp.Diff(vmExtension, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}
}

func TestCreateAgentVMASAKSBillingExtension(t *testing.T) {
	// Test with HostedMasterProfile Windows.
	cs := &api.ContainerService{
		Properties: &api.Properties{
			OrchestratorProfile: &api.OrchestratorProfile{
				KubernetesConfig: &api.KubernetesConfig{
					UseManagedIdentity: true,
				},
			},
			HostedMasterProfile: &api.HostedMasterProfile{
				DNSPrefix: "foodns",
			},
		},
	}

	profile := &api.AgentPoolProfile{
		Name:   "sample",
		OSType: "Windows",
	}

	vmExtension := CreateAgentVMASAKSBillingExtension(cs, profile)

	expected := VirtualMachineExtensionARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionCompute')]",
			DependsOn: []string{
				"[concat('Microsoft.Compute/virtualMachines/', variables('sampleVMNamePrefix'), copyIndex(variables('sampleOffset')))]",
			},
			Copy: map[string]string{
				"count": "[sub(variables('sampleCount'), variables('sampleOffset'))]",
				"name":  "vmLoopNode",
			},
		},

		VirtualMachineExtension: compute.VirtualMachineExtension{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[concat(variables('sampleVMNamePrefix'), copyIndex(variables('sampleOffset')), '/computeAksLinuxBilling')]"),
			VirtualMachineExtensionProperties: &compute.VirtualMachineExtensionProperties{
				Publisher:               to.StringPtr("Microsoft.AKS"),
				TypeHandlerVersion:      to.StringPtr("1.0"),
				AutoUpgradeMinorVersion: to.BoolPtr(true),
				Settings:                &map[string]interface{}{},
				Type:                    to.StringPtr("Compute.AKS-Engine.Windows.Billing"),
			},
			Type: to.StringPtr("Microsoft.Compute/virtualMachines/extensions"),
		},
	}

	diff := cmp.Diff(vmExtension, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}

	// Test with MasterProfile with windows.
	cs = &api.ContainerService{
		Properties: &api.Properties{
			OrchestratorProfile: &api.OrchestratorProfile{
				KubernetesConfig: &api.KubernetesConfig{
					UseManagedIdentity: true,
				},
			},
			MasterProfile: &api.MasterProfile{
				DNSPrefix: "foodns",
			},
		},
	}

	vmExtension = CreateAgentVMASAKSBillingExtension(cs, profile)
	expected.VirtualMachineExtensionProperties.Type = to.StringPtr("Compute.AKS-Engine.Windows.Billing")
	diff = cmp.Diff(vmExtension, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}

	// Test with MasterProfile with linux.
	cs = &api.ContainerService{
		Properties: &api.Properties{
			OrchestratorProfile: &api.OrchestratorProfile{
				KubernetesConfig: &api.KubernetesConfig{
					UseManagedIdentity: true,
				},
			},
			MasterProfile: &api.MasterProfile{
				DNSPrefix: "foodns",
			},
		},
	}

	profile = &api.AgentPoolProfile{
		Name:   "sample",
		OSType: "Linux",
	}

	vmExtension = CreateAgentVMASAKSBillingExtension(cs, profile)
	expected.VirtualMachineExtensionProperties.Type = to.StringPtr("Compute.AKS-Engine.Linux.Billing")
	diff = cmp.Diff(vmExtension, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}

	// Test with HostedMasterProfile with linux.
	cs = &api.ContainerService{
		Properties: &api.Properties{
			OrchestratorProfile: &api.OrchestratorProfile{
				KubernetesConfig: &api.KubernetesConfig{
					UseManagedIdentity: true,
				},
			},
			HostedMasterProfile: &api.HostedMasterProfile{
				DNSPrefix: "foodns",
			},
		},
	}

	profile = &api.AgentPoolProfile{
		Name:   "sample",
		OSType: "Linux",
	}

	vmExtension = CreateAgentVMASAKSBillingExtension(cs, profile)
	expected.VirtualMachineExtensionProperties.Type = to.StringPtr("Compute.AKS.Linux.Billing")
	diff = cmp.Diff(vmExtension, expected)

	if diff != "" {
		t.Errorf("unexpected diff while comparing: %s", diff)
	}
}

func TestCreateCustomScriptExtension(t *testing.T) {
	cs := &api.ContainerService{
		Location: "westus2",
		Properties: &api.Properties{
			FeatureFlags: &api.FeatureFlags{
				BlockOutboundInternet:    false,
				EnableCSERunInBackground: false,
			},
		},
	}

	cse := CreateCustomScriptExtension(cs)

	expectedCSE := VirtualMachineExtensionARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionCompute')]",
			Copy: map[string]string{
				"count": "[sub(variables('masterCount'), variables('masterOffset'))]",
				"name":  "vmLoopNode",
			},
			DependsOn: []string{"[concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')))]"},
		},
		VirtualMachineExtension: compute.VirtualMachineExtension{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[concat(variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')),'/cse', '-master-', copyIndex(variables('masterOffset')))]"),
			VirtualMachineExtensionProperties: &compute.VirtualMachineExtensionProperties{
				Publisher:               to.StringPtr("Microsoft.Azure.Extensions"),
				Type:                    to.StringPtr("CustomScript"),
				TypeHandlerVersion:      to.StringPtr("2.0"),
				AutoUpgradeMinorVersion: to.BoolPtr(true),
				Settings:                &map[string]interface{}{},
				ProtectedSettings: &map[string]interface{}{
					"commandToExecute": `[concat('retrycmd_if_failure() { r=$1; w=$2; t=$3; shift && shift && shift; for i in $(seq 1 $r); do timeout $t ${@}; [ $? -eq 0  ] && break || if [ $i -eq $r ]; then return 1; else sleep $w; fi; done }; ERR_OUTBOUND_CONN_FAIL=50; retrycmd_if_failure 50 1 3 nc -vz k8s.gcr.io 443 && retrycmd_if_failure 50 1 3 nc -vz gcr.io 443 && retrycmd_if_failure 50 1 3 nc -vz docker.io 443 || exit $ERR_OUTBOUND_CONN_FAIL; for i in $(seq 1 1200); do if [ -f /opt/azure/containers/provision.sh ]; then break; fi; if [ $i -eq 1200 ]; then exit 100; else sleep 1; fi; done; ', variables('provisionScriptParametersCommon'),' ',variables('provisionScriptParametersMaster'), ' /usr/bin/nohup /bin/bash -c "/bin/bash /opt/azure/containers/provision.sh >> /var/log/azure/cluster-provision.log 2>&1"')]`,
				},
			},
			Type: to.StringPtr("Microsoft.Compute/virtualMachines/extensions"),
			Tags: map[string]*string{},
		},
	}

	diff := cmp.Diff(cse, expectedCSE)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test with AzureChinaCloud location
	cs.Location = "chinaeast"
	cse = CreateCustomScriptExtension(cs)

	expectedCSE.ProtectedSettings = &map[string]interface{}{
		"commandToExecute": `[concat('retrycmd_if_failure() { r=$1; w=$2; t=$3; shift && shift && shift; for i in $(seq 1 $r); do timeout $t ${@}; [ $? -eq 0  ] && break || if [ $i -eq $r ]; then return 1; else sleep $w; fi; done }; ERR_OUTBOUND_CONN_FAIL=50; retrycmd_if_failure 50 1 3 nc -vz gcr.azk8s.cn 80 || exit $ERR_OUTBOUND_CONN_FAIL; for i in $(seq 1 1200); do if [ -f /opt/azure/containers/provision.sh ]; then break; fi; if [ $i -eq 1200 ]; then exit 100; else sleep 1; fi; done; ', variables('provisionScriptParametersCommon'),' ',variables('provisionScriptParametersMaster'), ' /usr/bin/nohup /bin/bash -c "/bin/bash /opt/azure/containers/provision.sh >> /var/log/azure/cluster-provision.log 2>&1"')]`,
	}

	diff = cmp.Diff(cse, expectedCSE)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}
}

func TestCreateAgentVMASCustomScriptExtension(t *testing.T) {
	cs := &api.ContainerService{
		Location: "westus2",
		Properties: &api.Properties{
			FeatureFlags: &api.FeatureFlags{
				BlockOutboundInternet:    false,
				EnableCSERunInBackground: false,
			},
		},
	}

	profile := &api.AgentPoolProfile{
		Name:   "sample",
		OSType: "Linux",
	}

	cse := createAgentVMASCustomScriptExtension(cs, profile)

	expectedCSE := VirtualMachineExtensionARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionCompute')]",
			Copy: map[string]string{
				"count": "[sub(variables('sampleCount'), variables('sampleOffset'))]",
				"name":  "vmLoopNode",
			},
			DependsOn: []string{"[concat('Microsoft.Compute/virtualMachines/', variables('sampleVMNamePrefix'), copyIndex(variables('sampleOffset')))]"},
		},
		VirtualMachineExtension: compute.VirtualMachineExtension{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[concat(variables('sampleVMNamePrefix'), copyIndex(variables('sampleOffset')),'/cse', '-agent-', copyIndex(variables('sampleOffset')))]"),
			VirtualMachineExtensionProperties: &compute.VirtualMachineExtensionProperties{
				Publisher:               to.StringPtr("Microsoft.Azure.Extensions"),
				Type:                    to.StringPtr("CustomScript"),
				TypeHandlerVersion:      to.StringPtr("2.0"),
				AutoUpgradeMinorVersion: to.BoolPtr(true),
				Settings:                &map[string]interface{}{},
				ProtectedSettings: &map[string]interface{}{
					"commandToExecute": `[concat('retrycmd_if_failure() { r=$1; w=$2; t=$3; shift && shift && shift; for i in $(seq 1 $r); do timeout $t ${@}; [ $? -eq 0  ] && break || if [ $i -eq $r ]; then return 1; else sleep $w; fi; done }; ERR_OUTBOUND_CONN_FAIL=50; retrycmd_if_failure 50 1 3 nc -vz k8s.gcr.io 443 && retrycmd_if_failure 50 1 3 nc -vz gcr.io 443 && retrycmd_if_failure 50 1 3 nc -vz docker.io 443 || exit $ERR_OUTBOUND_CONN_FAIL; for i in $(seq 1 1200); do if [ -f /opt/azure/containers/provision.sh ]; then break; fi; if [ $i -eq 1200 ]; then exit 100; else sleep 1; fi; done; ', variables('provisionScriptParametersCommon'),' GPU_NODE=false SGX_NODE=false /usr/bin/nohup /bin/bash -c "/bin/bash /opt/azure/containers/provision.sh >> /var/log/azure/cluster-provision.log 2>&1 "')]`,
				},
			},
			Type: to.StringPtr("Microsoft.Compute/virtualMachines/extensions"),
			Tags: nil,
		},
	}

	diff := cmp.Diff(cse, expectedCSE)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test with EnableRunInBackground and China Location
	cs.Properties.FeatureFlags.EnableCSERunInBackground = true
	cs.Location = "chinanorth"
	profile = &api.AgentPoolProfile{
		Name:   "sample",
		OSType: "Linux",
	}
	cse = createAgentVMASCustomScriptExtension(cs, profile)

	expectedCSE.ProtectedSettings = &map[string]interface{}{
		"commandToExecute": `[concat('retrycmd_if_failure() { r=$1; w=$2; t=$3; shift && shift && shift; for i in $(seq 1 $r); do timeout $t ${@}; [ $? -eq 0  ] && break || if [ $i -eq $r ]; then return 1; else sleep $w; fi; done }; ERR_OUTBOUND_CONN_FAIL=50; retrycmd_if_failure 50 1 3 nc -vz gcr.azk8s.cn 80 || exit $ERR_OUTBOUND_CONN_FAIL; for i in $(seq 1 1200); do if [ -f /opt/azure/containers/provision.sh ]; then break; fi; if [ $i -eq 1200 ]; then exit 100; else sleep 1; fi; done; ', variables('provisionScriptParametersCommon'),' GPU_NODE=false SGX_NODE=false /usr/bin/nohup /bin/bash -c "/bin/bash /opt/azure/containers/provision.sh >> /var/log/azure/cluster-provision.log 2>&1  &"')]`,
	}

	diff = cmp.Diff(cse, expectedCSE)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test with Windows agent profile

	profile = &api.AgentPoolProfile{
		Name:   "sample",
		OSType: "Windows",
	}

	cse = createAgentVMASCustomScriptExtension(cs, profile)

	expectedCSE.Publisher = to.StringPtr("Microsoft.Compute")
	expectedCSE.VirtualMachineExtensionProperties.Type = to.StringPtr("CustomScriptExtension")
	expectedCSE.TypeHandlerVersion = to.StringPtr("1.8")
	expectedCSE.ProtectedSettings = &map[string]interface{}{
		"commandToExecute": "[concat('powershell.exe -ExecutionPolicy Unrestricted -command \"', '$arguments = ', variables('singleQuote'),'-MasterIP ',variables('kubernetesAPIServerIP'),' -KubeDnsServiceIp ',parameters('kubeDnsServiceIp'),' -MasterFQDNPrefix ',variables('masterFqdnPrefix'),' -Location ',variables('location'),' -AgentKey ',parameters('clientPrivateKey'),' -AADClientId ',variables('servicePrincipalClientId'),' -AADClientSecret ',variables('servicePrincipalClientSecret'),variables('singleQuote'), ' ; ', variables('windowsCustomScriptSuffix'), '\" > %SYSTEMDRIVE%\\AzureData\\CustomDataSetupScript.log 2>&1')]",
	}

	diff = cmp.Diff(cse, expectedCSE)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}
}
