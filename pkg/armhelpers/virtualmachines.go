package armhelpers

import (
	"fmt"
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/engine"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
	"strconv"
)

func createVirtualMachine(cs *api.ContainerService) VirtualMachineARM {
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
		ApiVersion: "[variables('apiVersionCompute')]",
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
			"acsengineVersion":   to.StringPtr("[parameters('acsengineVersion')]"),
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

	t, err := engine.InitializeTemplateGenerator(engine.Context{})

	customDataStr := t.GetMasterCustomDataString(cs, engine.KubernetesMasterCustomDataYaml, cs.Properties)
	customDataStr = fmt.Sprintf("[base64(concat('%s'))]", customDataStr)
	osProfile.CustomData = to.StringPtr(customDataStr)

	if err != nil {
		panic(err)
	}

	if cs.Properties.LinuxProfile.HasSecrets() {
		osProfile.Secrets = &[]compute.VaultSecretGroup{
			//TODO: Need to address secrets case
		}
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
		ApiVersion: "[variables('apiVersionCompute')]",
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
			// TODO: had to override int32 value here
			DiskSizeGB: to.Int32Ptr(30),
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

	t, err := engine.InitializeTemplateGenerator(engine.Context{})

	if err != nil {
		panic(err)
	}

	customDataStr := t.GetKubernetesJumpboxCustomDataString(cs, cs.Properties)

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

func createVirtualMachineVMSS(cs *api.ContainerService) VirtualMachineScaleSetARM {
	isCustomVnet := cs.Properties.MasterProfile.IsCustomVNET()
	hasAvailabilityZones := cs.Properties.MasterProfile.HasAvailabilityZones()
	useManagedIdentity := cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity
	userAssignedIDEnabled := cs.Properties.OrchestratorProfile.KubernetesConfig.UseManagedIdentity &&
		cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedID != ""
	isAzureCNI := cs.Properties.OrchestratorProfile.IsAzureCNI()

	var dependencies []string

	if isCustomVnet {
		dependencies = append(dependencies, "[variables('nsgID')]")
	} else {
		dependencies = append(dependencies, "[variables('vnetID')]")
	}

	if to.Bool(cs.Properties.MasterProfile.CosmosEtcd) {
		dependencies = append(dependencies, "[resourceId('Microsoft.DocumentDB/databaseAccounts/', variables('cosmosAccountName'))]")
	}

	dependencies = append(dependencies, "[variables('masterLbID')]")

	armResource := ARMResource{
		ApiVersion: "[variables('apiVersionCompute')]",
		DependsOn:  dependencies,
	}

	virtualMachine := compute.VirtualMachineScaleSet{
		Location: to.StringPtr("[variables('location')]"),
		Name:     to.StringPtr("[concat(variables('masterVMNamePrefix'), 'vmss')]"),
		Tags: map[string]*string{
			"creationSource":     to.StringPtr("[concat(parameters('generatorCode'), '-', variables('masterVMNamePrefix'), 'vmss']"),
			"resourceNameSuffix": to.StringPtr("[parameters('nameSuffix')]"),
			"orchestrator":       to.StringPtr("[variables('orchestratorNameVersionTag')]"),
			"acsengineVersion":   to.StringPtr("[parameters('acsengineVersion')]"),
			"poolName":           to.StringPtr("master"),
		},
		Type: to.StringPtr("Microsoft.Compute/virtualMachineScaleSets"),
	}

	if hasAvailabilityZones {
		virtualMachine.Zones = &[]string{
			"[parameters('availabilityZones')]",
		}
	}

	if useManagedIdentity && userAssignedIDEnabled {
		identity := &compute.VirtualMachineScaleSetIdentity{}
		identity.Type = compute.ResourceIdentityTypeUserAssigned
		identity.UserAssignedIdentities = map[string]*compute.VirtualMachineScaleSetIdentityUserAssignedIdentitiesValue{
			"[variables('userAssignedIDReference')]": {},
		}
		virtualMachine.Identity = identity
	}

	virtualMachine.Sku = &compute.Sku{
		Tier:     to.StringPtr("Standard"),
		Capacity: to.Int64Ptr(int64(cs.Properties.MasterProfile.Count)),
		Name:     to.StringPtr("[parameters('masterVMSize')]"),
	}

	vmProperties := &compute.VirtualMachineScaleSetProperties{}

	vmProperties.SinglePlacementGroup = cs.Properties.MasterProfile.SinglePlacementGroup
	vmProperties.Overprovision = to.BoolPtr(false)
	vmProperties.UpgradePolicy = &compute.UpgradePolicy{
		Mode: compute.Manual,
	}

	netintconfig := compute.VirtualMachineScaleSetNetworkConfiguration{
		Name: to.StringPtr("[concat(variables('masterVMNamePrefix'), 'netintconfig')]"),
		VirtualMachineScaleSetNetworkConfigurationProperties: &compute.VirtualMachineScaleSetNetworkConfigurationProperties{
			Primary: to.BoolPtr(true),
		},
	}

	if isCustomVnet {
		netintconfig.NetworkSecurityGroup = &compute.SubResource{
			ID: to.StringPtr("[variables('nsgID')]"),
		}
	}

	var ipConfigurations []compute.VirtualMachineScaleSetIPConfiguration

	for i := 1; i <= cs.Properties.MasterProfile.Count; i++ {
		ipConfig := compute.VirtualMachineScaleSetIPConfiguration{
			Name: to.StringPtr(fmt.Sprintf("ipconfig%d", i)),
		}

		ipConfigProps := compute.VirtualMachineScaleSetIPConfigurationProperties{
			Subnet: &compute.APIEntityReference{
				ID: to.StringPtr("[variables('vnetSubnetIDMaster')]"),
			},
		}
		if i == 1 {
			ipConfigProps.Primary = to.BoolPtr(true)
			ipConfigProps.LoadBalancerBackendAddressPools = &[]compute.SubResource{
				{
					ID: to.StringPtr("[concat(variables('masterLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"),
				},
			}
			ipConfigProps.LoadBalancerInboundNatPools = &[]compute.SubResource{
				{
					ID: to.StringPtr("[concat(variables('masterLbID'),'/inboundNatPools/SSH-', variables('masterVMNamePrefix'), 'natpools')]"),
				},
			}
		} else {
			ipConfigProps.Primary = to.BoolPtr(false)
		}
		ipConfigurations = append(ipConfigurations, ipConfig)
	}
	netintconfig.IPConfigurations = &ipConfigurations

	if cs.Properties.LinuxProfile.HasCustomNodesDNS() {
		netintconfig.DNSSettings = &compute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
			DNSServers: &[]string{
				"[parameters('dnsServer')]",
			},
		}
	}

	if isAzureCNI {
		netintconfig.EnableIPForwarding = to.BoolPtr(true)
	}

	networkProfile := compute.VirtualMachineScaleSetNetworkProfile{
		NetworkInterfaceConfigurations: &[]compute.VirtualMachineScaleSetNetworkConfiguration{
			netintconfig,
		},
	}

	osProfile := compute.VirtualMachineScaleSetOSProfile{
		AdminUsername:      to.StringPtr("[parameters('linuxAdminUsername')]"),
		ComputerNamePrefix: to.StringPtr("[concat(variables('masterVMNamePrefix'), 'vmss')]"),
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

	t, err := engine.InitializeTemplateGenerator(engine.Context{})

	customDataStr := t.GetMasterCustomDataString(cs, engine.KubernetesMasterCustomDataYaml, cs.Properties)
	customDataStr = fmt.Sprintf("[base64(concat('%s'))]", customDataStr)
	osProfile.CustomData = to.StringPtr(customDataStr)

	if err != nil {
		panic(err)
	}

	if cs.Properties.LinuxProfile.HasSecrets() {
		osProfile.Secrets = &[]compute.VaultSecretGroup{
			//TODO: Need to address secrets case
		}
	}

	storageProfile := compute.VirtualMachineScaleSetStorageProfile{}
	imageRef := cs.Properties.MasterProfile.ImageRef
	useMasterCustomImage := imageRef != nil && len(imageRef.Name) > 0 && len(imageRef.ResourceGroup) > 0
	if !useMasterCustomImage {
		etcdSizeGB, _ := strconv.Atoi(cs.Properties.OrchestratorProfile.KubernetesConfig.EtcdDiskSizeGB)
		dataDisk := compute.VirtualMachineScaleSetDataDisk{
			CreateOption: compute.DiskCreateOptionTypesEmpty,
			DiskSizeGB:   to.Int32Ptr(int32(etcdSizeGB)),
			Lun:          to.Int32Ptr(0),
		}
		storageProfile.DataDisks = &[]compute.VirtualMachineScaleSetDataDisk{
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

	osDisk := &compute.VirtualMachineScaleSetOSDisk{
		Caching:      compute.CachingTypesReadWrite,
		CreateOption: compute.DiskCreateOptionTypesFromImage,
	}

	if cs.Properties.MasterProfile.OSDiskSizeGB > 0 {
		osDisk.DiskSizeGB = to.Int32Ptr(int32(cs.Properties.MasterProfile.OSDiskSizeGB))
	}

	storageProfile.OsDisk = osDisk
	storageProfile.ImageReference = imgReference

	var extensions []compute.VirtualMachineScaleSetExtension

	if useManagedIdentity {
		managedIdentityExtension := compute.VirtualMachineScaleSetExtension{
			Name: to.StringPtr("[concat(variables('masterVMNamePrefix'), 'vmss-ManagedIdentityExtension')]"),
			VirtualMachineScaleSetExtensionProperties: &compute.VirtualMachineScaleSetExtensionProperties{
				Publisher:               to.StringPtr("Microsoft.ManagedIdentity"),
				Type:                    to.StringPtr("ManagedIdentityExtensionForLinux"),
				TypeHandlerVersion:      to.StringPtr("1.0"),
				AutoUpgradeMinorVersion: to.BoolPtr(true),
				Settings: map[string]interface{}{
					"port": 50343,
				},
				ProtectedSettings: map[string]interface{}{},
			},
		}
		extensions = append(extensions, managedIdentityExtension)
	}

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

	extensions = append(extensions, vmssCSE)

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
		extensions = append(extensions, aksBillingExtension)
	}

	extensionProfile := compute.VirtualMachineScaleSetExtensionProfile{
		Extensions: &extensions,
	}

	vmProperties.VirtualMachineProfile = &compute.VirtualMachineScaleSetVMProfile{
		NetworkProfile:   &networkProfile,
		OsProfile:        &osProfile,
		StorageProfile:   &storageProfile,
		ExtensionProfile: &extensionProfile,
	}

	return VirtualMachineScaleSetARM{
		ARMResource:            armResource,
		VirtualMachineScaleSet: virtualMachine,
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
	}

	if hasDisks {
		dataDiskDep := fmt.Sprintf("[concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(add(div(copyIndex(variables('%[1]sOffset')),variables('maxVMsPerStorageAccount')),variables('%[1]sStorageAccountOffset')),variables('dataStorageAccountPrefixSeed')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(add(div(copyIndex(variables('%[1]sOffset')),variables('maxVMsPerStorageAccount')),variables('%[1]sStorageAccountOffset')),variables('dataStorageAccountPrefixSeed')),variables('storageAccountPrefixesCount'))],variables('%[1]sDataAccountName'))]", profile.Name)
		dependencies = append(dependencies, dataDiskDep)
	}

	dependencies = append(dependencies, fmt.Sprintf("[concat('Microsoft.Network/networkInterfaces/', variables('%[1]sVMNamePrefix'), 'nic-', copyIndex(variables('%[1]sOffset')))]", profile.Name))

	dependencies = append(dependencies, fmt.Sprintf("[concat('Microsoft.Compute/availabilitySets/', variables('%[1]sAvailabilitySet'))]", profile.Name))

	tags := map[string]string{
		"creationSource":   fmt.Sprintf("[concat(parameters('generatorCode'), '-', variables('%[1]sVMNamePrefix'), copyIndex(variables('%[1]sOffset')))]", profile.Name),
		"orchestrator":     "[variables('orchestratorNameVersionTag')]",
		"acsengineVersion": "[parameters('acsengineVersion')]",
		"poolName":         profile.Name,
	}

	if profile.IsWindows() {
		tags["resourceNameSuffix"] = "[variables('winResourceNamePrefix')]"
	} else {
		tags["resourceNameSuffix"] = "[variables('nameSuffix')]"
	}

	armResource := ARMResource{
		ApiVersion: "[variables('apiVersionCompute')]",
		DependsOn:  dependencies,
		Copy: map[string]string{
			"count": fmt.Sprintf("[sub(variables('%[1]sCount'), variables('%[1]sOffset'))]", profile.Name),
		},
	}

	virtualMachine := compute.VirtualMachine{
		Location: to.StringPtr("[variables('location')]"),
		Name:     to.StringPtr(fmt.Sprintf("[concat(variables('%[1]sVMNamePrefix'), copyIndex(variables('%[1]sOffset')))]", profile.Name)),
		Type:     to.StringPtr("Microsoft.Compute/virtualMachines"),
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

	t, err := engine.InitializeTemplateGenerator(engine.Context{})

	if !profile.IsWindows() {
		osProfile.AdminUsername = to.StringPtr("[parameters('linuxAdminUsername')]")
		osProfile.LinuxConfiguration = &compute.LinuxConfiguration{
			DisablePasswordAuthentication: to.BoolPtr(true),
			SSH: &compute.SSHConfiguration{
				PublicKeys: &[]compute.SSHPublicKey{
					{
						KeyData: to.StringPtr("[parameters('sshRSAPublicKey')]"),
						Path:    to.StringPtr("[variables('sshKeyPath')]"),
					},
				},
			},
		}
		if err != nil {
			panic(err)
		}

		agentCustomData := fmt.Sprintf("[base64(concat('%s'))]", t.GetKubernetesAgentCustomDataString(cs, profile))
		osProfile.CustomData = to.StringPtr(agentCustomData)

		if cs.Properties.LinuxProfile.HasSecrets() {
			//osProfile.Secrets = &[]compute.VaultSecretGroup {
			//	"[variables('linuxProfileSecrets')]"
			//}
		}
	} else {
		osProfile.AdminUsername = to.StringPtr("[parameters('windowsAdminUsername')]")
		osProfile.AdminPassword = to.StringPtr("[parameters('windowsAdminPassword')]")
		agentCustomData := fmt.Sprintf("[base64(concat('%s'))]", t.GetKubernetesWindowsAgentCustomDataString(cs, profile))
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
			storageProfile.DataDisks = getDataDisks(profile)
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
			storageProfile.DataDisks = getDataDisks(profile)
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

func getDataDisks(profile *api.AgentPoolProfile) *[]compute.DataDisk {
	var dataDisks []compute.DataDisk
	for i, diskSize := range profile.DiskSizesGB {
		dataDisk := compute.DataDisk{
			DiskSizeGB:   to.Int32Ptr(int32(diskSize)),
			Lun:          to.Int32Ptr(int32(i)),
			CreateOption: compute.DiskCreateOptionTypesEmpty,
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
