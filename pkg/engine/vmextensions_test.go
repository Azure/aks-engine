package engine

import (
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/google/go-cmp/cmp"
	"testing"
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
				"[concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')), '/extensions/ManagedIdentityExtension')]",
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
