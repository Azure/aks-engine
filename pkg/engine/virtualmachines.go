// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
)

func CreateVirtualMachine(cs *api.ContainerService) VirtualMachineARM {
	hasAvailabilityZones := cs.Properties.MasterProfile.HasAvailabilityZones()
	isStorageAccount := cs.Properties.MasterProfile.IsStorageAccount()
	useManagedIdentity := cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity
	userAssignedIDEnabled := cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity &&
		cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedID != ""

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

	virtualMachine := compute.VirtualMachine{
		Location: to.StringPtr("[variables('location')]"),
		Name:     to.StringPtr("[concat(variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')))]"),
		Tags: map[string]*string{
			"creationSource":     to.StringPtr("[concat(parameters('generatorCode'), '-', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')))]"),
			"resourceNameSuffix": to.StringPtr("[parameters('nameSuffix')]"),
			"orchestrator":       to.StringPtr("[variables('orchestratorNameVersionTag')]"),
			"aksEngineVersion":   to.StringPtr("[parameters('aksEngineVersion')]"),
			"poolName":           to.StringPtr("master"),
		},
		Type: to.StringPtr("Microsoft.Compute/virtualMachines"),
	}

	if hasAvailabilityZones {
		virtualMachine.Zones = &[]string{
			"split(string(parameters('availabilityZones')[mod(copyIndex(variables('masterOffset')), length(parameters('availabilityZones')))]), ',')",
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

	if len(cs.Properties.LinuxProfile.SSH.PublicKeys) > 1 {
		publicKeyPath := "[variables('sshKeyPath')]"
		var publicKeys []compute.SSHPublicKey
		for _, publicKey := range cs.Properties.LinuxProfile.SSH.PublicKeys {
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

	customDataStr := getCustomDataFromJSON(t.GetMasterCustomDataJSON(cs))
	osProfile.CustomData = to.StringPtr(customDataStr)

	if err != nil {
		panic(err)
	}

	if cs.Properties.LinuxProfile.HasSecrets() {
		vsg := getVaultSecretGroup(cs.Properties.LinuxProfile)
		osProfile.Secrets = &vsg
	}
	vmProperties.OsProfile = osProfile

	storageProfile := &compute.StorageProfile{}
	imageRef := cs.Properties.MasterProfile.ImageRef
	useMasterCustomImage := imageRef != nil && len(imageRef.Name) > 0 && len(imageRef.ResourceGroup) > 0
	if !useMasterCustomImage {
		etcdSizeGB, _ := strconv.Atoi(cs.Properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB)
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
	if useMasterCustomImage {
		imgReference.ID = to.StringPtr("[resourceId(parameters('osImageResourceGroup'), 'Microsoft.Compute/images', parameters('osImageName'))]")
	} else {
		imgReference.Offer = to.StringPtr("[parameters('osImageOffer')]")
		imgReference.Publisher = to.StringPtr("[parameters('osImagePublisher')]")
		imgReference.Sku = to.StringPtr("[parameters('osImageSku')]")
		imgReference.Version = to.StringPtr("[parameters('osImageVersion')]")
	}

	osDisk := &compute.OSDisk{
		Caching:      compute.CachingTypesReadWrite,
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

	jumpBoxIsManagedDisks := cs.Properties.OrchestratorProfile.KubernetesConfig.PrivateJumpboxProvision() && cs.Properties.OrchestratorProfile.KubernetesConfig.PrivateCluster.JumpboxProfile.StorageProfile == api.ManagedDisks

	if jumpBoxIsManagedDisks {
		storageProfile.OsDisk = &compute.OSDisk{
			CreateOption: compute.DiskCreateOptionTypesFromImage,
			DiskSizeGB:   to.Int32Ptr(int32(cs.Properties.OrchestratorProfile.KubernetesConfig.PrivateCluster.JumpboxProfile.OSDiskSizeGB)),
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
			CustomData: to.StringPtr(fmt.Sprintf("[base64(concat('%s'))]", customDataStr)),
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
	useManagedIdentity := cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity
	userAssignedIDEnabled := cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity &&
		cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedID != ""

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

		if len(cs.Properties.LinuxProfile.SSH.PublicKeys) > 1 {
			publicKeyPath := "[variables('sshKeyPath')]"
			publicKeys := []compute.SSHPublicKey{}
			for _, publicKey := range cs.Properties.LinuxProfile.SSH.PublicKeys {
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

		agentCustomData := getCustomDataFromJSON(t.GetKubernetesAgentCustomDataJSON(cs, profile))
		osProfile.CustomData = to.StringPtr(agentCustomData)

		if cs.Properties.LinuxProfile.HasSecrets() {
			vsg := getVaultSecretGroup(cs.Properties.LinuxProfile)
			osProfile.Secrets = &vsg
		}
	} else {
		osProfile.AdminUsername = to.StringPtr("[parameters('windowsAdminUsername')]")
		osProfile.AdminPassword = to.StringPtr("[parameters('windowsAdminPassword')]")
		osProfile.WindowsConfiguration = &compute.WindowsConfiguration{
			EnableAutomaticUpdates: to.BoolPtr(cs.Properties.WindowsProfile.GetEnableWindowsUpdate()),
		}
		agentCustomData := getCustomDataFromJSON(t.GetKubernetesWindowsAgentCustomDataJSON(cs, profile))
		osProfile.CustomData = to.StringPtr(agentCustomData)
	}

	virtualMachine.OsProfile = &osProfile

	storageProfile := compute.StorageProfile{}

	if profile.IsWindows() {
		if cs.Properties.WindowsProfile.HasCustomImage() {
			storageProfile.ImageReference = &compute.ImageReference{
				ID: to.StringPtr(fmt.Sprintf("[resourceId('Microsoft.Compute/images','%sCustomWindowsImage')]", profile.Name)),
			}
		} else {
			storageProfile.ImageReference = &compute.ImageReference{
				Offer:     to.StringPtr("[parameters('agentWindowsOffer')]"),
				Publisher: to.StringPtr("[parameters('agentWindowsPublisher')]"),
				Sku:       to.StringPtr("[parameters('agentWindowsSku')]"),
				Version:   to.StringPtr("[parameters('agentWindowsVersion')]"),
			}
		}

		if profile.HasDisks() {
			storageProfile.DataDisks = getArmDataDisks(profile)
		}

	} else {
		imageRef := profile.ImageRef
		useAgentCustomImage := imageRef != nil && len(imageRef.Name) > 0 && len(imageRef.ResourceGroup) > 0
		if useAgentCustomImage {
			storageProfile.ImageReference = &compute.ImageReference{
				ID: to.StringPtr(fmt.Sprintf("[resourceId(variables('%[1]sosImageResourceGroup'), 'Microsoft.Compute/images', variables('%[1]sosImageName'))]", profile.Name)),
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
		Caching:      compute.CachingTypesReadWrite,
	}

	if profile.IsStorageAccount() {
		osDisk.Name = to.StringPtr(fmt.Sprintf("[concat(variables('%[1]sVMNamePrefix'), copyIndex(variables('%[1]sOffset')),'-osdisk')]", profile.Name))
		osDisk.Vhd = &compute.VirtualHardDisk{
			URI: to.StringPtr(fmt.Sprintf("[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(div(copyIndex(variables('%[1]sOffset')),variables('maxVMsPerStorageAccount')),variables('%[1]sStorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(div(copyIndex(variables('%[1]sOffset')),variables('maxVMsPerStorageAccount')),variables('%[1]sStorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('%[1]sAccountName')),variables('apiVersionStorage')).primaryEndpoints.blob,'osdisk/', variables('%[1]sVMNamePrefix'), copyIndex(variables('%[1]sOffset')), '-osdisk.vhd')]", profile.Name)),
		}
	}

	if profile.OSDiskSizeGB > 0 {
		osDisk.DiskSizeGB = to.Int32Ptr(int32(profile.OSDiskSizeGB))
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
			Caching:      compute.CachingTypesReadOnly,
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
	json.Unmarshal([]byte(jsonStr), &customDataObj)
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
