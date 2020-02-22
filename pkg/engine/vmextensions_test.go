// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
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
				Type:                    to.StringPtr("Compute.AKS.Windows.Billing"),
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

	// userAssignedID is not enabled in above ContainerService definition
	var userAssignedIDEnabled = false

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
					"commandToExecute": `[concat('echo $(date),$(hostname);  for i in $(seq 1 1200); do grep -Fq "EOF" /opt/azure/containers/provision.sh && break; if [ $i -eq 1200 ]; then exit 100; else sleep 1; fi; done; ', variables('provisionScriptParametersCommon'),` + generateUserAssignedIdentityClientIDParameter(userAssignedIDEnabled) + `,variables('provisionScriptParametersMaster'), ' IS_VHD=false /usr/bin/nohup /bin/bash -c "/bin/bash /opt/azure/containers/provision.sh >> /var/log/azure/cluster-provision.log 2>&1"')]`,
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
}

func TestCreateCustomScriptExtensionWithHostedMaster(t *testing.T) {
	cs := &api.ContainerService{
		Location: "westus2",
		Properties: &api.Properties{
			FeatureFlags: &api.FeatureFlags{
				BlockOutboundInternet:    false,
				EnableCSERunInBackground: false,
			},
			HostedMasterProfile: &api.HostedMasterProfile{
				DNSPrefix: "foodns",
			},
		},
	}

	cse := CreateCustomScriptExtension(cs)

	// userAssignedID is not enabled in above ContainerService definition
	var userAssignedIDEnabled = false

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
					"commandToExecute": `[concat('echo $(date),$(hostname); retrycmd_if_failure() { r=$1; w=$2; t=$3; shift && shift && shift; for i in $(seq 1 $r); do timeout $t ${@}; [ $? -eq 0  ] && break || if [ $i -eq $r ]; then return 1; else sleep $w; fi; done }; ERR_OUTBOUND_CONN_FAIL=50; retrycmd_if_failure 50 1 3 nc -vz mcr.microsoft.com 443 || exit $ERR_OUTBOUND_CONN_FAIL; for i in $(seq 1 1200); do grep -Fq "EOF" /opt/azure/containers/provision.sh && break; if [ $i -eq 1200 ]; then exit 100; else sleep 1; fi; done; ', variables('provisionScriptParametersCommon'),` + generateUserAssignedIdentityClientIDParameter(userAssignedIDEnabled) + `,variables('provisionScriptParametersMaster'), ' IS_VHD=false /usr/bin/nohup /bin/bash -c "/bin/bash /opt/azure/containers/provision.sh >> /var/log/azure/cluster-provision.log 2>&1"')]`,
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
		"commandToExecute": `[concat('echo $(date),$(hostname); retrycmd_if_failure() { r=$1; w=$2; t=$3; shift && shift && shift; for i in $(seq 1 $r); do timeout $t ${@}; [ $? -eq 0  ] && break || if [ $i -eq $r ]; then return 1; else sleep $w; fi; done }; ERR_OUTBOUND_CONN_FAIL=50; retrycmd_if_failure 50 1 3 nc -vz gcr.azk8s.cn 443 || exit $ERR_OUTBOUND_CONN_FAIL; for i in $(seq 1 1200); do grep -Fq "EOF" /opt/azure/containers/provision.sh && break; if [ $i -eq 1200 ]; then exit 100; else sleep 1; fi; done; ', variables('provisionScriptParametersCommon'),` + generateUserAssignedIdentityClientIDParameter(userAssignedIDEnabled) + `,variables('provisionScriptParametersMaster'), ' IS_VHD=false /usr/bin/nohup /bin/bash -c "/bin/bash /opt/azure/containers/provision.sh >> /var/log/azure/cluster-provision.log 2>&1"')]`,
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
		Distro: api.AKSUbuntu1604,
	}

	cse := createAgentVMASCustomScriptExtension(cs, profile)

	// userAssignedID is not enabled in above ContainerService definition
	var userAssignedIDEnabled = false

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
					"commandToExecute": `[concat('echo $(date),$(hostname);  for i in $(seq 1 1200); do grep -Fq "EOF" /opt/azure/containers/provision.sh && break; if [ $i -eq 1200 ]; then exit 100; else sleep 1; fi; done; ', variables('provisionScriptParametersCommon'),` + generateUserAssignedIdentityClientIDParameter(userAssignedIDEnabled) + `,' IS_VHD=true GPU_NODE=false SGX_NODE=false AUDITD_ENABLED=false /usr/bin/nohup /bin/bash -c "/bin/bash /opt/azure/containers/provision.sh >> /var/log/azure/cluster-provision.log 2>&1"')]`,
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

	// Test with BlockOutboundInternet=true
	cseValNoOutboundInternetCheck := `[concat('echo $(date),$(hostname);  for i in $(seq 1 1200); do grep -Fq "EOF" /opt/azure/containers/provision.sh && break; if [ $i -eq 1200 ]; then exit 100; else sleep 1; fi; done; ', variables('provisionScriptParametersCommon'),` + generateUserAssignedIdentityClientIDParameter(userAssignedIDEnabled) + `,' IS_VHD=false GPU_NODE=false SGX_NODE=false AUDITD_ENABLED=false /usr/bin/nohup /bin/bash -c "/bin/bash /opt/azure/containers/provision.sh >> /var/log/azure/cluster-provision.log 2>&1"')]`
	cs.Properties.FeatureFlags.BlockOutboundInternet = true
	profile = &api.AgentPoolProfile{
		Name:   "sample",
		OSType: "Linux",
	}
	cse = createAgentVMASCustomScriptExtension(cs, profile)

	expectedCSE.ProtectedSettings = &map[string]interface{}{
		"commandToExecute": cseValNoOutboundInternetCheck,
	}

	diff = cmp.Diff(cse, expectedCSE)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test with Azure Stack on Linux
	cs.Properties.FeatureFlags.BlockOutboundInternet = false
	cs.Properties.CustomCloudProfile = &api.CustomCloudProfile{}
	profile = &api.AgentPoolProfile{
		Name:   "sample",
		OSType: "Linux",
	}
	cse = createAgentVMASCustomScriptExtension(cs, profile)

	expectedCSE.ProtectedSettings = &map[string]interface{}{
		"commandToExecute": cseValNoOutboundInternetCheck,
	}

	diff = cmp.Diff(cse, expectedCSE)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	// Test with EnableRunInBackground and China Location
	cs.Properties.FeatureFlags.BlockOutboundInternet = false
	cs.Properties.CustomCloudProfile = nil
	cs.Properties.OrchestratorProfile = nil
	cs.Properties.FeatureFlags.EnableCSERunInBackground = true
	cs.Location = "chinanorth"
	profile = &api.AgentPoolProfile{
		Name:   "sample",
		OSType: "Linux",
	}
	cse = createAgentVMASCustomScriptExtension(cs, profile)

	expectedCSE.ProtectedSettings = &map[string]interface{}{
		"commandToExecute": `[concat('echo $(date),$(hostname);  for i in $(seq 1 1200); do grep -Fq "EOF" /opt/azure/containers/provision.sh && break; if [ $i -eq 1200 ]; then exit 100; else sleep 1; fi; done; ', variables('provisionScriptParametersCommon'),` + generateUserAssignedIdentityClientIDParameter(userAssignedIDEnabled) + `,' IS_VHD=false GPU_NODE=false SGX_NODE=false AUDITD_ENABLED=false /usr/bin/nohup /bin/bash -c "/bin/bash /opt/azure/containers/provision.sh >> /var/log/azure/cluster-provision.log 2>&1 &"')]`,
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
		"commandToExecute": "[concat('echo %DATE%,%TIME%,%COMPUTERNAME% && powershell.exe -ExecutionPolicy Unrestricted -command \"', '$arguments = ', variables('singleQuote'),'-MasterIP ',variables('kubernetesAPIServerIP'),' -KubeDnsServiceIp ',parameters('kubeDnsServiceIp'),' -MasterFQDNPrefix ',variables('masterFqdnPrefix'),' -Location ',variables('location'),' -TargetEnvironment ',parameters('targetEnvironment'),' -AgentKey ',parameters('clientPrivateKey'),' -AADClientId ',variables('servicePrincipalClientId'),' -AADClientSecret ',variables('singleQuote'),variables('singleQuote'),base64(variables('servicePrincipalClientSecret')),variables('singleQuote'),variables('singleQuote'),' -NetworkAPIVersion ',variables('apiVersionNetwork'),' ',variables('singleQuote'), ' ; ', variables('windowsCustomScriptSuffix'), '\" > %SYSTEMDRIVE%\\AzureData\\CustomDataSetupScript.log 2>&1 ; exit $LASTEXITCODE')]",
	}

	diff = cmp.Diff(cse, expectedCSE)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}
}

func TestCreateCustomExtensions(t *testing.T) {
	properties := &api.Properties{
		OrchestratorProfile: &api.OrchestratorProfile{
			OrchestratorType: Kubernetes,
		},
		ExtensionProfiles: []*api.ExtensionProfile{
			{
				Name:    "winrm",
				Version: "v1",
				RootURL: "https://raw.githubusercontent.com/Azure/aks-engine/master/",
			},
			{
				Name:    "hello-world-k8s",
				Version: "v1",
				RootURL: "https://raw.githubusercontent.com/Azure/aks-engine/master/",
			},
		},
		AgentPoolProfiles: []*api.AgentPoolProfile{
			{
				Name:                "windowspool1",
				OSType:              api.Windows,
				AvailabilityProfile: "AvailabilitySet",
				Extensions: []api.Extension{
					{
						Name: "winrm",
					},
				},
			},
			{
				Name:                "windowspool2",
				OSType:              api.Windows,
				AvailabilityProfile: "AvailabilitySet",
				Extensions: []api.Extension{
					{
						Name: "winrm",
					},
					{
						Name: "hello-world-k8s",
					},
				},
			},
		},
	}

	extensions := CreateCustomExtensions(properties)

	expectedDeployments := []DeploymentARM{
		{
			DeploymentARMResource: DeploymentARMResource{
				APIVersion: "[variables('apiVersionDeployments')]",
				Copy: map[string]string{
					"count": "[sub(variables('windowspool1Count'), variables('windowspool1Offset'))]",
					"name":  "winrmExtensionLoop",
				},
				DependsOn: []string{"[concat('Microsoft.Compute/virtualMachines/', variables('windowspool1VMNamePrefix'), copyIndex(variables('windowspool1Offset')), '/extensions/cse-agent-', copyIndex(variables('windowspool1Offset')))]"},
			},
			DeploymentExtended: resources.DeploymentExtended{
				Name: to.StringPtr("[concat(variables('windowspool1VMNamePrefix'), copyIndex(variables('windowspool1Offset')), 'winrm')]"),
				Properties: &resources.DeploymentPropertiesExtended{
					TemplateLink: &resources.TemplateLink{
						URI:            to.StringPtr("https://raw.githubusercontent.com/Azure/aks-engine/master/extensions/winrm/v1/template.json"),
						ContentVersion: to.StringPtr("1.0.0.0"),
					},
					Parameters: map[string]interface{}{
						"apiVersionDeployments": map[string]interface{}{"value": "[variables('apiVersionDeployments')]"},
						"artifactsLocation":     map[string]interface{}{"value": "https://raw.githubusercontent.com/Azure/aks-engine/master/"},
						"extensionParameters":   map[string]interface{}{"value": "[parameters('winrmParameters')]"},
						"targetVMName":          map[string]interface{}{"value": "[concat(variables('windowspool1VMNamePrefix'), copyIndex(variables('windowspool1Offset')))]"},
						"targetVMType":          map[string]interface{}{"value": "agent"},
						"vmIndex":               map[string]interface{}{"value": "[copyIndex(variables('windowspool1Offset'))]"},
					},
					Mode: resources.DeploymentMode("Incremental"),
				},
				Type: to.StringPtr("Microsoft.Resources/deployments"),
			},
		},
		{
			DeploymentARMResource: DeploymentARMResource{
				APIVersion: "[variables('apiVersionDeployments')]",
				Copy: map[string]string{
					"count": "[sub(variables('windowspool2Count'), variables('windowspool2Offset'))]",
					"name":  "winrmExtensionLoop",
				},
				DependsOn: []string{"[concat('Microsoft.Compute/virtualMachines/', variables('windowspool2VMNamePrefix'), copyIndex(variables('windowspool2Offset')), '/extensions/cse-agent-', copyIndex(variables('windowspool2Offset')))]"},
			},
			DeploymentExtended: resources.DeploymentExtended{
				Name: to.StringPtr("[concat(variables('windowspool2VMNamePrefix'), copyIndex(variables('windowspool2Offset')), 'winrm')]"),
				Properties: &resources.DeploymentPropertiesExtended{
					TemplateLink: &resources.TemplateLink{
						URI:            to.StringPtr("https://raw.githubusercontent.com/Azure/aks-engine/master/extensions/winrm/v1/template.json"),
						ContentVersion: to.StringPtr("1.0.0.0"),
					},
					Parameters: map[string]interface{}{
						"apiVersionDeployments": map[string]interface{}{"value": "[variables('apiVersionDeployments')]"},
						"artifactsLocation":     map[string]interface{}{"value": "https://raw.githubusercontent.com/Azure/aks-engine/master/"},
						"extensionParameters":   map[string]interface{}{"value": "[parameters('winrmParameters')]"},
						"targetVMName":          map[string]interface{}{"value": "[concat(variables('windowspool2VMNamePrefix'), copyIndex(variables('windowspool2Offset')))]"},
						"targetVMType":          map[string]interface{}{"value": "agent"},
						"vmIndex":               map[string]interface{}{"value": "[copyIndex(variables('windowspool2Offset'))]"},
					},
					Mode: resources.DeploymentMode("Incremental"),
				},
				Type: to.StringPtr("Microsoft.Resources/deployments"),
			},
		},
		{
			DeploymentARMResource: DeploymentARMResource{
				APIVersion: "[variables('apiVersionDeployments')]",
				Copy: map[string]string{
					"count": "[sub(variables('windowspool2Count'), variables('windowspool2Offset'))]",
					"name":  "helloWorldExtensionLoop",
				},
				DependsOn: []string{"[concat(variables('windowspool2VMNamePrefix'), copyIndex(variables('windowspool2Offset')), 'winrm')]"},
			},
			DeploymentExtended: resources.DeploymentExtended{
				Name: to.StringPtr("[concat(variables('windowspool2VMNamePrefix'), copyIndex(variables('windowspool2Offset')), 'HelloWorldK8s')]"),
				Properties: &resources.DeploymentPropertiesExtended{
					TemplateLink: &resources.TemplateLink{
						URI:            to.StringPtr("https://raw.githubusercontent.com/Azure/aks-engine/master/extensions/hello-world-k8s/v1/template.json"),
						ContentVersion: to.StringPtr("1.0.0.0"),
					},
					Parameters: map[string]interface{}{
						"apiVersionDeployments": map[string]interface{}{"value": "[variables('apiVersionDeployments')]"},
						"artifactsLocation":     map[string]interface{}{"value": "https://raw.githubusercontent.com/Azure/aks-engine/master/"},
						"extensionParameters":   map[string]interface{}{"value": "[parameters('hello-world-k8sParameters')]"},
						"targetVMName":          map[string]interface{}{"value": "[concat(variables('windowspool2VMNamePrefix'), copyIndex(variables('windowspool2Offset')))]"},
						"targetVMType":          map[string]interface{}{"value": "agent"},
						"vmIndex":               map[string]interface{}{"value": "[copyIndex(variables('windowspool2Offset'))]"},
					},
					Mode: resources.DeploymentMode("Incremental"),
				},
				Type: to.StringPtr("Microsoft.Resources/deployments"),
			},
		},
	}

	diff := cmp.Diff(extensions, expectedDeployments)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}

	properties = &api.Properties{
		OrchestratorProfile: &api.OrchestratorProfile{
			OrchestratorType: Kubernetes,
		},
		ExtensionProfiles: []*api.ExtensionProfile{
			{
				Name:    "hello-world-k8s",
				Version: "v1",
				RootURL: "https://raw.githubusercontent.com/Azure/aks-engine/master/",
			},
		},
		MasterProfile: &api.MasterProfile{
			Count:               3,
			DNSPrefix:           "testcluster",
			AvailabilityProfile: "AvailabilitySet",
			Extensions: []api.Extension{
				{
					Name: "hello-world-k8s",
				},
			},
		},
	}

	extensions = CreateCustomExtensions(properties)

	expectedDeployments = []DeploymentARM{
		{
			DeploymentARMResource: DeploymentARMResource{
				APIVersion: "[variables('apiVersionDeployments')]",
				Copy: map[string]string{
					"count": "[sub(variables('masterCount'), variables('masterOffset'))]",
					"name":  "helloWorldExtensionLoop",
				},
				DependsOn: []string{"[concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')), '/extensions/cse-master-', copyIndex(variables('masterOffset')))]"},
			},
			DeploymentExtended: resources.DeploymentExtended{
				Name: to.StringPtr("[concat(variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')), 'HelloWorldK8s')]"),
				Properties: &resources.DeploymentPropertiesExtended{
					TemplateLink: &resources.TemplateLink{
						URI:            to.StringPtr("https://raw.githubusercontent.com/Azure/aks-engine/master/extensions/hello-world-k8s/v1/template.json"),
						ContentVersion: to.StringPtr("1.0.0.0"),
					},
					Parameters: map[string]interface{}{
						"apiVersionDeployments": map[string]interface{}{"value": "[variables('apiVersionDeployments')]"},
						"artifactsLocation":     map[string]interface{}{"value": "https://raw.githubusercontent.com/Azure/aks-engine/master/"},
						"extensionParameters":   map[string]interface{}{"value": "[parameters('hello-world-k8sParameters')]"},
						"targetVMName":          map[string]interface{}{"value": "[concat(variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')))]"},
						"targetVMType":          map[string]interface{}{"value": "master"},
						"vmIndex":               map[string]interface{}{"value": "[copyIndex(variables('masterOffset'))]"},
					},
					Mode: resources.DeploymentMode("Incremental"),
				},
				Type: to.StringPtr("Microsoft.Resources/deployments"),
			},
		},
	}

	diff = cmp.Diff(extensions, expectedDeployments)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}
}
