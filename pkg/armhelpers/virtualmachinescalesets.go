package armhelpers

import (
	"fmt"
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/engine"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
)

func createAgentVMSS(cs *api.ContainerService, profile *api.AgentPoolProfile) VirtualMachineScaleSetARM {
	armResource := ARMResource{
		ApiVersion: "[variables('apiVersionCompute')]",
	}
	var dependencies []string

	if profile.IsCustomVNET() {
		dependencies = append(dependencies, "[variables('nsgID')]")
	} else {
		dependencies = append(dependencies, "[variables('vnetID')]")
	}

	armResource.DependsOn = dependencies

	virtualMachineScaleSet := compute.VirtualMachineScaleSet{
		Name:     to.StringPtr(fmt.Sprintf("[variables('%sVMNamePrefix')]", profile.Name)),
		Type:     to.StringPtr("Microsoft.Compute/virtualMachineScaleSets"),
		Location: to.StringPtr("[variables('location')]"),
		Sku: &compute.Sku{
			Tier:     to.StringPtr("Standard"),
			Capacity: to.Int64Ptr(int64(profile.Count)), //"[variables('{{.Name}}Count')]",
			Name:     to.StringPtr(fmt.Sprintf("[variables('%sVMSize')]", profile.Name)),
		},
		Tags: map[string]*string{
			"creationSource":     to.StringPtr(fmt.Sprintf("[concat(parameters('generatorCode'), '-', variables('%sVMNamePrefix'))]", profile.Name)),
			"resourceNameSuffix": to.StringPtr("[parameters('nameSuffix')]"),
			"orchestrator":       to.StringPtr("[variables('orchestratorNameVersionTag')]"),
			"poolName":           to.StringPtr(profile.Name),
		},
	}

	if profile.HasAvailabilityZones() {
		virtualMachineScaleSet.Zones = &[]string{
			fmt.Sprintf("[parameters('%sAvailabilityZones')]", profile.Name),
		}
	}

	useManagedIdentity := cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity
	if useManagedIdentity {
		userAssignedIdentityEnabled := cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedID != ""
		if userAssignedIdentityEnabled {
			virtualMachineScaleSet.Identity = &compute.VirtualMachineScaleSetIdentity{
				Type: compute.ResourceIdentityTypeUserAssigned,
				UserAssignedIdentities: map[string]*compute.VirtualMachineScaleSetIdentityUserAssignedIdentitiesValue{
					"[variables('userAssignedIDReference')]": {},
				},
			}
		} else {
			virtualMachineScaleSet.Identity = &compute.VirtualMachineScaleSetIdentity{
				Type: compute.ResourceIdentityTypeSystemAssigned,
			}
		}
	}

	vmssProperties := compute.VirtualMachineScaleSetProperties{
		SinglePlacementGroup: profile.SinglePlacementGroup,
		Overprovision:        to.BoolPtr(false),
		UpgradePolicy: &compute.UpgradePolicy{
			Mode: compute.Manual,
		},
	}

	vmssVMProfile := compute.VirtualMachineScaleSetVMProfile{}

	if profile.IsLowPriorityScaleSet() {
		vmssVMProfile.Priority = compute.VirtualMachinePriorityTypes(fmt.Sprintf("[variables('%sScaleSetPriority')]", profile.Name))
		vmssVMProfile.EvictionPolicy = compute.VirtualMachineEvictionPolicyTypes(fmt.Sprintf("[variables('%sScaleSetEvictionPolicy')]", profile.Name))
	}

	vmssNICConfig := compute.VirtualMachineScaleSetNetworkConfiguration{
		Name: to.StringPtr(fmt.Sprintf("[variables('%sVMNamePrefix')]", profile.Name)),
		VirtualMachineScaleSetNetworkConfigurationProperties: &compute.VirtualMachineScaleSetNetworkConfigurationProperties{
			Primary:                     to.BoolPtr(true),
			EnableAcceleratedNetworking: profile.AcceleratedNetworkingEnabled,
		},
	}

	if profile.IsCustomVNET() {
		vmssNICConfig.NetworkSecurityGroup = &compute.SubResource{
			ID: to.StringPtr("[variables('nsgID')]"),
		}
	}

	var ipConfigurations []compute.VirtualMachineScaleSetIPConfiguration

	for i := 0; i < profile.IPAddressCount; i++ {
		ipconfig := compute.VirtualMachineScaleSetIPConfiguration{
			Name: to.StringPtr("ipconfig" + string(i)),
			VirtualMachineScaleSetIPConfigurationProperties: &compute.VirtualMachineScaleSetIPConfigurationProperties{
				Subnet: &compute.APIEntityReference{
					ID: to.StringPtr(fmt.Sprintf("[variables('%sVnetSubnetID')]", profile.Name)),
				},
			},
		}
		if i == 0 {
			ipconfig.Primary = to.BoolPtr(true)
		}
		ipConfigurations = append(ipConfigurations, ipconfig)
	}

	vmssNICConfig.IPConfigurations = &ipConfigurations

	if cs.Properties.LinuxProfile.HasCustomNodesDNS() {
		vmssNICConfig.DNSSettings = &compute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
			DNSServers: &[]string{
				"[parameters('dnsServer')]",
			},
		}
	}

	if !cs.Properties.OrchestratorProfile.IsAzureCNI() {
		vmssNICConfig.EnableIPForwarding = to.BoolPtr(true)
	}

	vmssNetworkProfile := compute.VirtualMachineScaleSetNetworkProfile{
		NetworkInterfaceConfigurations: &[]compute.VirtualMachineScaleSetNetworkConfiguration{
			vmssNICConfig,
		},
	}

	vmssVMProfile.NetworkProfile = &vmssNetworkProfile

	t, err := engine.InitializeTemplateGenerator(engine.Context{})

	if err != nil {
		panic(err)
	}

	vmssOsProfile := compute.VirtualMachineScaleSetOSProfile{
		AdminUsername:      to.StringPtr("[parameters('linuxAdminUsername')]"),
		ComputerNamePrefix: to.StringPtr(fmt.Sprintf("[variables('%sVMNamePrefix')]", profile.Name)),
		CustomData:         to.StringPtr(t.GetKubernetesAgentCustomDataString(cs, profile)),
		LinuxConfiguration: &compute.LinuxConfiguration{
			DisablePasswordAuthentication: to.BoolPtr(true),
			SSH: &compute.SSHConfiguration{
				PublicKeys: &[]compute.SSHPublicKey{
					{
						KeyData: to.StringPtr("[parameters('sshRSAPublicKey')]"),
						Path:    to.StringPtr("[variables('sshKeyPath')]"),
					},
				},
			},
		},
	}

	//TODO : Need to address secrets
	if cs.Properties.LinuxProfile.HasSecrets() {
		vmssOsProfile.Secrets = &[]compute.VaultSecretGroup{
			//"[variables('linuxProfileSecrets')]",
		}
	}

	vmssVMProfile.OsProfile = &vmssOsProfile

	vmssStorageProfile := compute.VirtualMachineScaleSetStorageProfile{}

	imageRef := profile.ImageRef
	useAgentCustomImage := imageRef != nil && len(imageRef.Name) > 0 && len(imageRef.ResourceGroup) > 0

	if useAgentCustomImage {
		vmssStorageProfile.ImageReference = &compute.ImageReference{
			ID: to.StringPtr(fmt.Sprintf("[resourceId(variables('%[1]sosImageResourceGroup'), 'Microsoft.Compute/images', variables('%[1]sosImageName'))]", profile.Name)),
		}
	} else {
		vmssStorageProfile.ImageReference = &compute.ImageReference{
			Offer:     to.StringPtr(fmt.Sprintf("[variables('%sosImageOffer')]", profile.Name)),
			Publisher: to.StringPtr(fmt.Sprintf("[variables('%sosImagePublisher')]", profile.Name)),
			Sku:       to.StringPtr(fmt.Sprintf("[variables('%sosImageSKU')]", profile.Name)),
			Version:   to.StringPtr(fmt.Sprintf("[variables('%sosImageVersion')]", profile.Name)),
		}
		vmssStorageProfile.DataDisks = getVMSSDataDisks(profile)
	}

	osDisk := compute.VirtualMachineScaleSetOSDisk{
		CreateOption: compute.DiskCreateOptionTypesFromImage,
		Caching:      compute.CachingTypesReadWrite,
	}

	if profile.OSDiskSizeGB > 0 {
		osDisk.DiskSizeGB = to.Int32Ptr(int32(profile.OSDiskSizeGB))
	}

	vmssStorageProfile.OsDisk = &osDisk

	vmssVMProfile.StorageProfile = &vmssStorageProfile

	var vmssExtensions []compute.VirtualMachineScaleSetExtension

	outBoundCmd := ""
	registry := ""
	if !cs.Properties.FeatureFlags.BlockOutboundInternet {
		if cs.GetCloudSpecConfig().CloudName == api.AzureChinaCloud {
			registry = `gcr.azk8s.cn 80`
		} else {
			registry = `k8s.gcr.io 443 && retrycmd_if_failure 50 1 3 nc -vz gcr.io 443 && retrycmd_if_failure 50 1 3 nc -vz docker.io 443`
		}
		outBoundCmd = `ERR_OUTBOUND_CONN_FAIL=50; retrycmd_if_failure 50 1 3 nc -vz ` + registry + `|| exit $ERR_OUTBOUND_CONN_FAIL;`
	}

	vmssCSE := compute.VirtualMachineScaleSetExtension{
		Name: to.StringPtr("[concat(variables('masterVMNamePrefix'), 'vmssCSE')]"),
		VirtualMachineScaleSetExtensionProperties: &compute.VirtualMachineScaleSetExtensionProperties{
			Publisher:               to.StringPtr("Microsoft.Azure.Extensions"),
			Type:                    to.StringPtr("CustomScript"),
			TypeHandlerVersion:      to.StringPtr("2.0"),
			AutoUpgradeMinorVersion: to.BoolPtr(true),
			Settings:                map[string]interface{}{},
			ProtectedSettings: map[string]interface{}{
				"commandToExecute": "[concat('retrycmd_if_failure() { r=$1; w=$2; t=$3; shift && shift && shift; for i in $(seq 1 $r); do timeout $t ${@}; [ $? -eq 0  ] && break || if [ $i -eq $r ]; then return 1; else sleep $w; fi; done }; " + outBoundCmd + " for i in $(seq 1 1200); do if [ -f /opt/azure/containers/provision.sh ]; then break; fi; if [ $i -eq 1200 ]; then exit 100; else sleep 1; fi; done; ', variables('provisionScriptParametersCommon'),' ',variables('provisionScriptParametersMaster'), ' /usr/bin/nohup /bin/bash -c \"/bin/bash /opt/azure/containers/provision.sh >> /var/log/azure/cluster-provision.log 2>&1\"')]",
			},
		},
	}

	vmssExtensions = append(vmssExtensions, vmssCSE)

	if cs.GetCloudSpecConfig().CloudName == api.AzurePublicCloud {
		aksBillingExtension := compute.VirtualMachineScaleSetExtension{
			Name: to.StringPtr("[concat(variables('masterVMNamePrefix'), 'vmss-computeAksLinuxBilling')]"),
			VirtualMachineScaleSetExtensionProperties: &compute.VirtualMachineScaleSetExtensionProperties{
				Publisher:               to.StringPtr("Microsoft.AKS"),
				Type:                    to.StringPtr("Compute.AKS-Engine.Linux.Billing"),
				TypeHandlerVersion:      to.StringPtr("1.0"),
				AutoUpgradeMinorVersion: to.BoolPtr(true),
				Settings:                map[string]interface{}{},
			},
		}
		vmssExtensions = append(vmssExtensions, aksBillingExtension)
	}

	vmssVMProfile.ExtensionProfile = &compute.VirtualMachineScaleSetExtensionProfile{
		Extensions: &vmssExtensions,
	}

	vmssProperties.VirtualMachineProfile = &vmssVMProfile
	virtualMachineScaleSet.VirtualMachineScaleSetProperties = &vmssProperties

	return VirtualMachineScaleSetARM{
		ARMResource:            armResource,
		VirtualMachineScaleSet: virtualMachineScaleSet,
	}
}

func getVMSSDataDisks(profile *api.AgentPoolProfile) *[]compute.VirtualMachineScaleSetDataDisk {
	var dataDisks []compute.VirtualMachineScaleSetDataDisk
	for i, diskSize := range profile.DiskSizesGB {
		dataDisk := compute.VirtualMachineScaleSetDataDisk{
			DiskSizeGB:   to.Int32Ptr(int32(diskSize)),
			Lun:          to.Int32Ptr(int32(i)),
			CreateOption: compute.DiskCreateOptionTypesEmpty,
		}
		if profile.StorageProfile == api.StorageAccount {
			dataDisk.Name = to.StringPtr(fmt.Sprintf("[concat(variables('%sVMNamePrefix'), copyIndex(),'-datadisk%d')]", profile.Name, i))
			//dataDisk.Vhd = &compute.VirtualHardDisk{
			//	URI: to.StringPtr(fmt.Sprintf("[concat('http://',variables('storageAccountPrefixes')[mod(add(add(div(copyIndex(),variables('maxVMsPerStorageAccount')),variables('%sStorageAccountOffset')),variables('dataStorageAccountPrefixSeed')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(add(div(copyIndex(),variables('maxVMsPerStorageAccount')),variables('%sStorageAccountOffset')),variables('dataStorageAccountPrefixSeed')),variables('storageAccountPrefixesCount'))],variables('%sDataAccountName'),'.blob.core.windows.net/vhds/',variables('%sVMNamePrefix'),copyIndex(), '--datadisk%d.vhd')]",
			//		profile.Name, profile.Name, profile.Name, profile.Name, i)),
			//}
		}
		dataDisks = append(dataDisks, dataDisk)
	}
	return &dataDisks
}
