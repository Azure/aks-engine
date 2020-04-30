// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-12-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
)

func CreateMasterVM(cs *api.ContainerService) VirtualMachineARM {
	hasAvailabilityZones := cs.Properties.MasterProfile.HasAvailabilityZones()
	isStorageAccount := cs.Properties.MasterProfile.IsStorageAccount()
	kubernetesConfig := cs.Properties.OrchestratorProfile.KubernetesConfig

	var useManagedIdentity, userAssignedIDEnabled bool
	if kubernetesConfig != nil {
		useManagedIdentity = kubernetesConfig.UseManagedIdentity
		userAssignedIDEnabled = kubernetesConfig.UserAssignedIDEnabled()
	}

	var dependencies []string
	dependentNIC := "[concat('Microsoft.Network/networkInterfaces/', variables('masterVMNamePrefix'), 'nic-', copyIndex(variables('masterOffset')))]"
	dependencies = append(dependencies, dependentNIC)
	if !hasAvailabilityZones {
		dependencies = append(dependencies, "[concat('Microsoft.Compute/availabilitySets/',variables('masterAvailabilitySet'))]")
	}
	if isStorageAccount {
		dependencies = append(dependencies, "[variables('masterStorageAccountName')]")
	}

	armResource := ARMResource{
		APIVersion: "[variables('apiVersionCompute')]",
		Copy: map[string]string{
			"count": "[sub(variables('masterCount'), variables('masterOffset'))]",
			"name":  "vmLoopNode",
		},
		DependsOn: dependencies,
	}

	vmTags := map[string]*string{
		"creationSource":     to.StringPtr("[concat(parameters('generatorCode'), '-', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')))]"),
		"resourceNameSuffix": to.StringPtr("[parameters('nameSuffix')]"),
		"orchestrator":       to.StringPtr("[variables('orchestratorNameVersionTag')]"),
		"aksEngineVersion":   to.StringPtr("[parameters('aksEngineVersion')]"),
		"poolName":           to.StringPtr("master"),
	}

	if kubernetesConfig != nil && kubernetesConfig.IsContainerMonitoringAddonEnabled() {
		addon := kubernetesConfig.GetAddonByName(common.ContainerMonitoringAddonName)
		clusterDNSPrefix := "aks-engine-cluster"
		if cs.Properties.MasterProfile != nil && cs.Properties.MasterProfile.DNSPrefix != "" {
			clusterDNSPrefix = cs.Properties.MasterProfile.DNSPrefix
		}
		vmTags["logAnalyticsWorkspaceResourceId"] = to.StringPtr(addon.Config["logAnalyticsWorkspaceResourceId"])
		vmTags["clusterName"] = to.StringPtr(clusterDNSPrefix)
	}

	virtualMachine := compute.VirtualMachine{
		Location: to.StringPtr("[variables('location')]"),
		Name:     to.StringPtr("[concat(variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')))]"),
		Tags:     vmTags,
		Type:     to.StringPtr("Microsoft.Compute/virtualMachines"),
	}

	addCustomTagsToVM(cs.Properties.MasterProfile.CustomVMTags, &virtualMachine)

	if hasAvailabilityZones {
		virtualMachine.Zones = &[]string{
			"[string(parameters('availabilityZones')[mod(copyIndex(variables('masterOffset')), length(parameters('availabilityZones')))])]",
		}
	}

	if useManagedIdentity {
		identity := &compute.VirtualMachineIdentity{}
		if userAssignedIDEnabled {
			identity.Type = compute.ResourceIdentityTypeUserAssigned
			identity.UserAssignedIdentities = map[string]*compute.VirtualMachineIdentityUserAssignedIdentitiesValue{
				"[variables('userAssignedIDReference')]": {},
			}
		} else {
			identity.Type = compute.ResourceIdentityTypeSystemAssigned
		}
		virtualMachine.Identity = identity
	}

	associateAddonIdentitiesToVM(cs.Properties.AddonProfiles, &virtualMachine)

	vmProperties := &compute.VirtualMachineProperties{}

	if !hasAvailabilityZones {
		vmProperties.AvailabilitySet = &compute.SubResource{
			ID: to.StringPtr("[resourceId('Microsoft.Compute/availabilitySets',variables('masterAvailabilitySet'))]"),
		}
	}

	vmProperties.HardwareProfile = &compute.HardwareProfile{
		VMSize: compute.VirtualMachineSizeTypes(cs.Properties.MasterProfile.VMSize),
	}

	vmProperties.NetworkProfile = &compute.NetworkProfile{
		NetworkInterfaces: &[]compute.NetworkInterfaceReference{
			{
				ID: to.StringPtr("[resourceId('Microsoft.Network/networkInterfaces',concat(variables('masterVMNamePrefix'),'nic-', copyIndex(variables('masterOffset'))))]"),
			},
		},
	}

	osProfile := &compute.OSProfile{
		AdminUsername: to.StringPtr("[parameters('linuxAdminUsername')]"),
		ComputerName:  to.StringPtr("[concat(variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')))]"),
		LinuxConfiguration: &compute.LinuxConfiguration{
			DisablePasswordAuthentication: to.BoolPtr(true),
		},
	}

	linuxProfile := cs.Properties.LinuxProfile
	if linuxProfile != nil && len(linuxProfile.SSH.PublicKeys) > 1 {
		publicKeyPath := "[variables('sshKeyPath')]"
		var publicKeys []compute.SSHPublicKey
		for _, publicKey := range linuxProfile.SSH.PublicKeys {
			publicKeyTrimmed := strings.TrimSpace(publicKey.KeyData)
			publicKeys = append(publicKeys, compute.SSHPublicKey{
				Path:    &publicKeyPath,
				KeyData: &publicKeyTrimmed,
			})
		}
		osProfile.LinuxConfiguration.SSH = &compute.SSHConfiguration{
			PublicKeys: &publicKeys,
		}

	} else {
		osProfile.LinuxConfiguration.SSH = &compute.SSHConfiguration{
			PublicKeys: &[]compute.SSHPublicKey{
				{
					KeyData: to.StringPtr("[parameters('sshRSAPublicKey')]"),
					Path:    to.StringPtr("[variables('sshKeyPath')]"),
				},
			},
		}
	}

	t, err := InitializeTemplateGenerator(Context{})

	customDataStr := getCustomDataFromJSON(t.GetMasterCustomDataJSONObject(cs))
	osProfile.CustomData = to.StringPtr(customDataStr)

	if err != nil {
		panic(err)
	}

	if linuxProfile != nil && linuxProfile.HasSecrets() {
		vsg := getVaultSecretGroup(linuxProfile)
		osProfile.Secrets = &vsg
	}
	vmProperties.OsProfile = osProfile

	storageProfile := &compute.StorageProfile{}
	imageRef := cs.Properties.MasterProfile.ImageRef
	etcdSizeGB, _ := strconv.Atoi(kubernetesConfig.EtcdDiskSizeGB)
	if !cs.Properties.MasterProfile.HasCosmosEtcd() {
		dataDisk := compute.DataDisk{
			CreateOption: compute.DiskCreateOptionTypesEmpty,
			DiskSizeGB:   to.Int32Ptr(int32(etcdSizeGB)),
			Lun:          to.Int32Ptr(0),
			Name:         to.StringPtr("[concat(variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')),'-etcddisk')]"),
		}
		if cs.Properties.MasterProfile.IsStorageAccount() {
			dataDisk.Vhd = &compute.VirtualHardDisk{
				URI: to.StringPtr("[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('masterStorageAccountName')),variables('apiVersionStorage')).primaryEndpoints.blob,'vhds/', variables('masterVMNamePrefix'),copyIndex(variables('masterOffset')),'-etcddisk.vhd')]"),
			}
		}
		storageProfile.DataDisks = &[]compute.DataDisk{
			dataDisk,
		}
	}
	imgReference := &compute.ImageReference{}
	if cs.Properties.MasterProfile.HasImageRef() {
		if cs.Properties.MasterProfile.HasImageGallery() {
			imgReference.ID = to.StringPtr(fmt.Sprintf("[concat('/subscriptions/', '%s', '/resourceGroups/', parameters('osImageResourceGroup'), '/providers/Microsoft.Compute/galleries/', '%s', '/images/', parameters('osImageName'), '/versions/', '%s')]", imageRef.SubscriptionID, imageRef.Gallery, imageRef.Version))
		} else {
			imgReference.ID = to.StringPtr("[resourceId(parameters('osImageResourceGroup'), 'Microsoft.Compute/images', parameters('osImageName'))]")
		}
	} else {
		imgReference.Offer = to.StringPtr("[parameters('osImageOffer')]")
		imgReference.Publisher = to.StringPtr("[parameters('osImagePublisher')]")
		imgReference.Sku = to.StringPtr("[parameters('osImageSku')]")
		imgReference.Version = to.StringPtr("[parameters('osImageVersion')]")
	}

	osDisk := &compute.OSDisk{
		Caching:      compute.CachingTypes(cs.Properties.MasterProfile.OSDiskCachingType),
		CreateOption: compute.DiskCreateOptionTypesFromImage,
	}

	if isStorageAccount {
		osDisk.Name = to.StringPtr("[concat(variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')),'-osdisk')]")
		osDisk.Vhd = &compute.VirtualHardDisk{
			URI: to.StringPtr("[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('masterStorageAccountName')),variables('apiVersionStorage')).primaryEndpoints.blob,'vhds/',variables('masterVMNamePrefix'),copyIndex(variables('masterOffset')),'-osdisk.vhd')]"),
		}
	}

	if cs.Properties.MasterProfile.OSDiskSizeGB > 0 {
		osDisk.DiskSizeGB = to.Int32Ptr(int32(cs.Properties.MasterProfile.OSDiskSizeGB))
	}

	if to.Bool(cs.Properties.MasterProfile.UltraSSDEnabled) {
		vmProperties.AdditionalCapabilities = &compute.AdditionalCapabilities{
			UltraSSDEnabled: to.BoolPtr(true),
		}
	}

	storageProfile.OsDisk = osDisk
	storageProfile.ImageReference = imgReference
	vmProperties.StorageProfile = storageProfile

	virtualMachine.VirtualMachineProperties = vmProperties

	return VirtualMachineARM{
		ARMResource:    armResource,
		VirtualMachine: virtualMachine,
	}
}

func createJumpboxVirtualMachine(cs *api.ContainerService) VirtualMachineARM {
	armResource := ARMResource{
		APIVersion: "[variables('apiVersionCompute')]",
		DependsOn: []string{
			"[concat('Microsoft.Network/networkInterfaces/', variables('jumpboxNetworkInterfaceName'))]",
		},
	}

	kubernetesConfig := cs.Properties.OrchestratorProfile.KubernetesConfig

	vm := compute.VirtualMachine{
		Location: to.StringPtr("[variables('location')]"),
		Name:     to.StringPtr("[parameters('jumpboxVMName')]"),
		Type:     to.StringPtr("Microsoft.Compute/virtualMachines"),
	}

	storageProfile := compute.StorageProfile{
		ImageReference: &compute.ImageReference{
			Publisher: to.StringPtr("Canonical"),
			Offer:     to.StringPtr("UbuntuServer"),
			Sku:       to.StringPtr("16.04-LTS"),
			Version:   to.StringPtr("latest"),
		},
		DataDisks: &[]compute.DataDisk{},
	}

	var jumpBoxIsManagedDisks bool
	if kubernetesConfig != nil && kubernetesConfig.PrivateCluster != nil {
		jumpBoxIsManagedDisks = kubernetesConfig.PrivateJumpboxProvision() && kubernetesConfig.PrivateCluster.JumpboxProfile.StorageProfile == api.ManagedDisks
	}

	if jumpBoxIsManagedDisks {
		storageProfile.OsDisk = &compute.OSDisk{
			CreateOption: compute.DiskCreateOptionTypesFromImage,
			DiskSizeGB:   to.Int32Ptr(int32(kubernetesConfig.PrivateCluster.JumpboxProfile.OSDiskSizeGB)),
			ManagedDisk: &compute.ManagedDiskParameters{
				StorageAccountType: "[variables('vmSizesMap')[parameters('jumpboxVMSize')].storageAccountType]",
			},
		}
	} else {
		storageProfile.OsDisk = &compute.OSDisk{
			CreateOption: compute.DiskCreateOptionTypesFromImage,
			Vhd: &compute.VirtualHardDisk{
				URI: to.StringPtr("[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('jumpboxStorageAccountName')),variables('apiVersionStorage')).primaryEndpoints.blob,'vhds/',parameters('jumpboxVMName'),'jumpboxdisk.vhd')]"),
			},
			Name: to.StringPtr("[variables('jumpboxOSDiskName')]"),
		}
	}

	t, err := InitializeTemplateGenerator(Context{})

	if err != nil {
		panic(err)
	}

	customDataStr := getCustomDataFromJSON(t.GetJumpboxCustomDataJSON(cs))

	vmProperties := compute.VirtualMachineProperties{
		HardwareProfile: &compute.HardwareProfile{
			VMSize: "[parameters('jumpboxVMSize')]",
		},
		OsProfile: &compute.OSProfile{
			ComputerName:  to.StringPtr("[parameters('jumpboxVMName')]"),
			AdminUsername: to.StringPtr("[parameters('jumpboxUsername')]"),
			LinuxConfiguration: &compute.LinuxConfiguration{
				DisablePasswordAuthentication: to.BoolPtr(true),
				SSH: &compute.SSHConfiguration{
					PublicKeys: &[]compute.SSHPublicKey{
						{
							Path:    to.StringPtr("[concat('/home/', parameters('jumpboxUsername'), '/.ssh/authorized_keys')]"),
							KeyData: to.StringPtr("[parameters('jumpboxPublicKey')]"),
						},
					},
				},
			},
			CustomData: to.StringPtr(customDataStr),
		},
		NetworkProfile: &compute.NetworkProfile{
			NetworkInterfaces: &[]compute.NetworkInterfaceReference{
				{
					ID: to.StringPtr("[resourceId('Microsoft.Network/networkInterfaces', variables('jumpboxNetworkInterfaceName'))]"),
				},
			},
		},
		StorageProfile: &storageProfile,
	}

	vm.VirtualMachineProperties = &vmProperties

	return VirtualMachineARM{
		ARMResource:    armResource,
		VirtualMachine: vm,
	}
}

func createAgentAvailabilitySetVM(cs *api.ContainerService, profile *api.AgentPoolProfile) VirtualMachineARM {
	var dependencies []string

	isStorageAccount := profile.IsStorageAccount()
	hasDisks := profile.HasDisks()
	kubernetesConfig := cs.Properties.OrchestratorProfile.KubernetesConfig

	var useManagedIdentity, userAssignedIDEnabled bool

	if kubernetesConfig != nil {
		useManagedIdentity = kubernetesConfig.UseManagedIdentity
		userAssignedIDEnabled = kubernetesConfig.UserAssignedIDEnabled()
	}

	if isStorageAccount {
		storageDep := fmt.Sprintf("[concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(div(copyIndex(variables('%[1]sOffset')),variables('maxVMsPerStorageAccount')),variables('%[1]sStorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(div(copyIndex(variables('%[1]sOffset')),variables('maxVMsPerStorageAccount')),variables('%[1]sStorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('%[1]sAccountName'))]", profile.Name)
		dependencies = append(dependencies, storageDep)
		if hasDisks {
			dataDiskDep := fmt.Sprintf("[concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(add(div(copyIndex(variables('%[1]sOffset')),variables('maxVMsPerStorageAccount')),variables('%[1]sStorageAccountOffset')),variables('dataStorageAccountPrefixSeed')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(add(div(copyIndex(variables('%[1]sOffset')),variables('maxVMsPerStorageAccount')),variables('%[1]sStorageAccountOffset')),variables('dataStorageAccountPrefixSeed')),variables('storageAccountPrefixesCount'))],variables('%[1]sDataAccountName'))]", profile.Name)
			dependencies = append(dependencies, dataDiskDep)
		}
	}

	dependencies = append(dependencies, fmt.Sprintf("[concat('Microsoft.Network/networkInterfaces/', variables('%[1]sVMNamePrefix'), 'nic-', copyIndex(variables('%[1]sOffset')))]", profile.Name))

	dependencies = append(dependencies, fmt.Sprintf("[concat('Microsoft.Compute/availabilitySets/', variables('%[1]sAvailabilitySet'))]", profile.Name))

	if profile.IsWindows() {
		windowsProfile := cs.Properties.WindowsProfile
		// Add dependency for Image resource created by createWindowsImage()
		if windowsProfile.HasCustomImage() {
			dependencies = append(dependencies, fmt.Sprintf("%sCustomWindowsImage", profile.Name))
		}
	}

	tags := map[string]*string{
		"creationSource":   to.StringPtr(fmt.Sprintf("[concat(parameters('generatorCode'), '-', variables('%[1]sVMNamePrefix'), copyIndex(variables('%[1]sOffset')))]", profile.Name)),
		"orchestrator":     to.StringPtr("[variables('orchestratorNameVersionTag')]"),
		"aksEngineVersion": to.StringPtr("[parameters('aksEngineVersion')]"),
		"poolName":         to.StringPtr(profile.Name),
	}

	if profile.IsWindows() {
		tags["resourceNameSuffix"] = to.StringPtr("[variables('winResourceNamePrefix')]")
	} else {
		tags["resourceNameSuffix"] = to.StringPtr("[parameters('nameSuffix')]")
	}

	armResource := ARMResource{
		APIVersion: "[variables('apiVersionCompute')]",
		DependsOn:  dependencies,
		Copy: map[string]string{
			"count": fmt.Sprintf("[sub(variables('%[1]sCount'), variables('%[1]sOffset'))]", profile.Name),
			"name":  "vmLoopNode",
		},
	}

	virtualMachine := compute.VirtualMachine{
		Location: to.StringPtr("[variables('location')]"),
		Name:     to.StringPtr(fmt.Sprintf("[concat(variables('%[1]sVMNamePrefix'), copyIndex(variables('%[1]sOffset')))]", profile.Name)),
		Type:     to.StringPtr("Microsoft.Compute/virtualMachines"),
		VirtualMachineProperties: &compute.VirtualMachineProperties{
			NetworkProfile: &compute.NetworkProfile{
				NetworkInterfaces: &[]compute.NetworkInterfaceReference{
					{
						ID: to.StringPtr(fmt.Sprintf("[resourceId('Microsoft.Network/networkInterfaces',concat(variables('%[1]sVMNamePrefix'), 'nic-', copyIndex(variables('%[1]sOffset'))))]", profile.Name)),
					},
				},
			},
		},
		Tags: tags,
	}

	addCustomTagsToVM(profile.CustomVMTags, &virtualMachine)

	if useManagedIdentity {
		if userAssignedIDEnabled && !profile.IsWindows() {
			virtualMachine.Identity = &compute.VirtualMachineIdentity{
				Type: compute.ResourceIdentityTypeUserAssigned,
				UserAssignedIdentities: map[string]*compute.VirtualMachineIdentityUserAssignedIdentitiesValue{
					"[variables('userAssignedIDReference')]": {},
				},
			}
		} else {
			virtualMachine.Identity = &compute.VirtualMachineIdentity{
				Type: compute.ResourceIdentityTypeSystemAssigned,
			}
		}
	}

	associateAddonIdentitiesToVM(cs.Properties.AddonProfiles, &virtualMachine)

	virtualMachine.AvailabilitySet = &compute.SubResource{
		ID: to.StringPtr(fmt.Sprintf("[resourceId('Microsoft.Compute/availabilitySets',variables('%sAvailabilitySet'))]", profile.Name)),
	}

	vmSize := fmt.Sprintf("[variables('%sVMSize')]", profile.Name)

	virtualMachine.HardwareProfile = &compute.HardwareProfile{
		VMSize: compute.VirtualMachineSizeTypes(vmSize),
	}

	osProfile := compute.OSProfile{
		ComputerName: to.StringPtr(fmt.Sprintf("[concat(variables('%[1]sVMNamePrefix'), copyIndex(variables('%[1]sOffset')))]", profile.Name)),
	}

	t, err := InitializeTemplateGenerator(Context{})

	if !profile.IsWindows() {
		osProfile.AdminUsername = to.StringPtr("[parameters('linuxAdminUsername')]")
		osProfile.LinuxConfiguration = &compute.LinuxConfiguration{
			DisablePasswordAuthentication: to.BoolPtr(true),
		}

		linuxProfile := cs.Properties.LinuxProfile
		if linuxProfile != nil && len(linuxProfile.SSH.PublicKeys) > 1 {
			publicKeyPath := "[variables('sshKeyPath')]"
			publicKeys := []compute.SSHPublicKey{}
			for _, publicKey := range linuxProfile.SSH.PublicKeys {
				publicKeyTrimmed := strings.TrimSpace(publicKey.KeyData)
				publicKeys = append(publicKeys, compute.SSHPublicKey{
					Path:    &publicKeyPath,
					KeyData: &publicKeyTrimmed,
				})
			}
			osProfile.LinuxConfiguration.SSH = &compute.SSHConfiguration{
				PublicKeys: &publicKeys,
			}

		} else {
			osProfile.LinuxConfiguration.SSH = &compute.SSHConfiguration{
				PublicKeys: &[]compute.SSHPublicKey{
					{
						KeyData: to.StringPtr("[parameters('sshRSAPublicKey')]"),
						Path:    to.StringPtr("[variables('sshKeyPath')]"),
					},
				},
			}
		}

		if err != nil {
			panic(err)
		}

		agentCustomData := getCustomDataFromJSON(t.GetKubernetesLinuxNodeCustomDataJSONObject(cs, profile))
		osProfile.CustomData = to.StringPtr(agentCustomData)

		if linuxProfile != nil && linuxProfile.HasSecrets() {
			vsg := getVaultSecretGroup(linuxProfile)
			osProfile.Secrets = &vsg
		}
	} else {
		osProfile.AdminUsername = to.StringPtr("[parameters('windowsAdminUsername')]")
		osProfile.AdminPassword = to.StringPtr("[parameters('windowsAdminPassword')]")
		osProfile.WindowsConfiguration = &compute.WindowsConfiguration{
			EnableAutomaticUpdates: to.BoolPtr(cs.Properties.WindowsProfile.GetEnableWindowsUpdate()),
		}
		agentCustomData := getCustomDataFromJSON(t.GetKubernetesWindowsNodeCustomDataJSONObject(cs, profile))
		osProfile.CustomData = to.StringPtr(agentCustomData)
	}

	virtualMachine.OsProfile = &osProfile

	storageProfile := compute.StorageProfile{}

	if profile.IsWindows() {
		storageProfile.ImageReference = createWindowsImageReference(profile.Name, cs.Properties.WindowsProfile)

		if profile.HasDisks() {
			storageProfile.DataDisks = getArmDataDisks(profile)
		}
	} else {
		imageRef := profile.ImageRef
		if profile.HasImageRef() {
			if profile.HasImageGallery() {
				storageProfile.ImageReference = &compute.ImageReference{
					ID: to.StringPtr(fmt.Sprintf("[concat('/subscriptions/', '%s', '/resourceGroups/', parameters('%sosImageResourceGroup'), '/providers/Microsoft.Compute/galleries/', '%s', '/images/', parameters('%sosImageName'), '/versions/', '%s')]", imageRef.SubscriptionID, profile.Name, imageRef.Gallery, profile.Name, imageRef.Version)),
				}
			} else {
				storageProfile.ImageReference = &compute.ImageReference{
					ID: to.StringPtr(fmt.Sprintf("[resourceId(variables('%[1]sosImageResourceGroup'), 'Microsoft.Compute/images', variables('%[1]sosImageName'))]", profile.Name)),
				}
			}
		} else {
			storageProfile.ImageReference = &compute.ImageReference{
				Offer:     to.StringPtr(fmt.Sprintf("[variables('%sosImageOffer')]", profile.Name)),
				Publisher: to.StringPtr(fmt.Sprintf("[variables('%sosImagePublisher')]", profile.Name)),
				Sku:       to.StringPtr(fmt.Sprintf("[variables('%sosImageSKU')]", profile.Name)),
				Version:   to.StringPtr(fmt.Sprintf("[variables('%sosImageVersion')]", profile.Name)),
			}
			storageProfile.DataDisks = getArmDataDisks(profile)
		}
	}

	osDisk := compute.OSDisk{
		CreateOption: compute.DiskCreateOptionTypesFromImage,
		Caching:      compute.CachingTypes(profile.OSDiskCachingType),
	}

	if profile.IsStorageAccount() {
		osDisk.Name = to.StringPtr(fmt.Sprintf("[concat(variables('%[1]sVMNamePrefix'), copyIndex(variables('%[1]sOffset')),'-osdisk')]", profile.Name))
		osDisk.Vhd = &compute.VirtualHardDisk{
			URI: to.StringPtr(fmt.Sprintf("[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(div(copyIndex(variables('%[1]sOffset')),variables('maxVMsPerStorageAccount')),variables('%[1]sStorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(div(copyIndex(variables('%[1]sOffset')),variables('maxVMsPerStorageAccount')),variables('%[1]sStorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('%[1]sAccountName')),variables('apiVersionStorage')).primaryEndpoints.blob,'osdisk/', variables('%[1]sVMNamePrefix'), copyIndex(variables('%[1]sOffset')), '-osdisk.vhd')]", profile.Name)),
		}
	}

	if profile.IsEphemeral() {
		osDisk.DiffDiskSettings = &compute.DiffDiskSettings{
			Option: compute.Local,
		}
	}

	if profile.OSDiskSizeGB > 0 {
		osDisk.DiskSizeGB = to.Int32Ptr(int32(profile.OSDiskSizeGB))
	}

	if profile.DiskEncryptionSetID != "" {
		osDisk.ManagedDisk = &compute.ManagedDiskParameters{
			DiskEncryptionSet: &compute.DiskEncryptionSetParameters{ID: to.StringPtr(profile.DiskEncryptionSetID)},
		}
	}

	if to.Bool(profile.UltraSSDEnabled) {
		virtualMachine.AdditionalCapabilities = &compute.AdditionalCapabilities{
			UltraSSDEnabled: to.BoolPtr(true),
		}
	}

	storageProfile.OsDisk = &osDisk

	virtualMachine.StorageProfile = &storageProfile

	return VirtualMachineARM{
		ARMResource:    armResource,
		VirtualMachine: virtualMachine,
	}
}

func getArmDataDisks(profile *api.AgentPoolProfile) *[]compute.DataDisk {
	var dataDisks []compute.DataDisk
	for i, diskSize := range profile.DiskSizesGB {
		dataDisk := compute.DataDisk{
			DiskSizeGB:   to.Int32Ptr(int32(diskSize)),
			Lun:          to.Int32Ptr(int32(i)),
			CreateOption: compute.DiskCreateOptionTypesEmpty,
			Caching:      compute.CachingTypes(profile.DataDiskCachingType),
		}
		if profile.StorageProfile == api.StorageAccount {
			dataDisk.Name = to.StringPtr(fmt.Sprintf("[concat(variables('%sVMNamePrefix'), copyIndex(),'-datadisk%d')]", profile.Name, i))
			dataDisk.Vhd = &compute.VirtualHardDisk{
				URI: to.StringPtr(fmt.Sprintf("[concat('http://',variables('storageAccountPrefixes')[mod(add(add(div(copyIndex(),variables('maxVMsPerStorageAccount')),variables('%sStorageAccountOffset')),variables('dataStorageAccountPrefixSeed')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(add(div(copyIndex(),variables('maxVMsPerStorageAccount')),variables('%sStorageAccountOffset')),variables('dataStorageAccountPrefixSeed')),variables('storageAccountPrefixesCount'))],variables('%sDataAccountName'),'.blob.core.windows.net/vhds/',variables('%sVMNamePrefix'),copyIndex(), '--datadisk%d.vhd')]",
					profile.Name, profile.Name, profile.Name, profile.Name, i)),
			}
		}
		dataDisks = append(dataDisks, dataDisk)
	}
	return &dataDisks
}

func getCustomDataFromJSON(jsonStr string) string {
	var customDataObj map[string]string
	err := json.Unmarshal([]byte(jsonStr), &customDataObj)
	if err != nil {
		panic(err)
	}
	return customDataObj["customData"]
}

func getVaultSecretGroup(linuxProfile *api.LinuxProfile) []compute.VaultSecretGroup {
	var vaultSecretGroups []compute.VaultSecretGroup
	if linuxProfile.HasSecrets() {
		for idx, lVault := range linuxProfile.Secrets {
			computeVault := compute.VaultSecretGroup{
				SourceVault: &compute.SubResource{
					ID: to.StringPtr(fmt.Sprintf("[parameters('linuxKeyVaultID%d')]", idx)),
				},
			}
			var vaultCerts []compute.VaultCertificate
			for certIdx := range lVault.VaultCertificates {
				vaultCert := compute.VaultCertificate{
					CertificateURL: to.StringPtr(fmt.Sprintf("[parameters('linuxKeyVaultID%dCertificateURL%d')]", idx, certIdx)),
				}
				vaultCerts = append(vaultCerts, vaultCert)
			}
			computeVault.VaultCertificates = &vaultCerts
			vaultSecretGroups = append(vaultSecretGroups, computeVault)
		}
	}
	return vaultSecretGroups
}

func addCustomTagsToVM(tags map[string]string, vm *compute.VirtualMachine) {
	for key, value := range tags {
		_, found := vm.Tags[key]
		if !found {
			vm.Tags[key] = to.StringPtr(value)
		}
	}
}

func associateAddonIdentitiesToVM(addonProfiles map[string]api.AddonProfile, virtualMachine *compute.VirtualMachine) {
	if virtualMachine == nil {
		return
	}
	for _, addonProfile := range addonProfiles {
		if addonProfile.Enabled && addonProfile.Identity != nil && addonProfile.Identity.ResourceID != "" {
			// We need to associate addon's identity to VM, there're 3 cases:
			// 1. virtualMachine.Identity is nil. In this case, we need to initialize "virtualMachine.Identity" and set its type to UserAssigned.
			// 2. virtualMachine.Identity is not nil, and its type is SystemAssigned. This case will happen in an MSI cluster and the VM uses system
			// assigned identity. In this case, we need to set current `virtualMachine.Identity.Type` to `ResourceIdentityTypeSystemAssignedUserAssigned`.
			// 3. virtualMachine.Identity is not nil, and its type is UserAssigned. This case will happen in an MSI cluster and the VM uses user assigned
			// identity. In this case, no additional step is needed. Just keep `virtualMachine.Identity.Type` unchanged and fill in addon's identity later.
			// Note: virtualMachine.Identity is not nil and its type is None will NEVER happen in current AKS-Engine's implementation.
			if virtualMachine.Identity == nil {
				virtualMachine.Identity = &compute.VirtualMachineIdentity{
					Type:                   compute.ResourceIdentityTypeUserAssigned,
					UserAssignedIdentities: make(map[string]*compute.VirtualMachineIdentityUserAssignedIdentitiesValue),
				}
			} else if virtualMachine.Identity.Type == compute.ResourceIdentityTypeSystemAssigned {
				virtualMachine.Identity.Type = compute.ResourceIdentityTypeSystemAssignedUserAssigned
				virtualMachine.Identity.UserAssignedIdentities = make(map[string]*compute.VirtualMachineIdentityUserAssignedIdentitiesValue)
			} else if virtualMachine.Identity.Type == compute.ResourceIdentityTypeNone {
				// Note: in current AKS-Engine's implementation, we will never enter into this branch. Just handle it here in case implementation
				// changes later.
				virtualMachine.Identity.Type = compute.ResourceIdentityTypeUserAssigned
				virtualMachine.Identity.UserAssignedIdentities = make(map[string]*compute.VirtualMachineIdentityUserAssignedIdentitiesValue)
			}

			virtualMachine.Identity.UserAssignedIdentities[addonProfile.Identity.ResourceID] = &compute.VirtualMachineIdentityUserAssignedIdentitiesValue{}
		}
	}
}
