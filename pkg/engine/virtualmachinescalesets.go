// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
)

func CreateMasterVMSS(cs *api.ContainerService) VirtualMachineScaleSetARM {

	masterProfile := cs.Properties.MasterProfile
	orchProfile := cs.Properties.OrchestratorProfile
	k8sConfig := orchProfile.KubernetesConfig
	linuxProfile := cs.Properties.LinuxProfile

	isCustomVnet := masterProfile.IsCustomVNET()
	hasAvailabilityZones := masterProfile.HasAvailabilityZones()
	useManagedIdentity := k8sConfig.UseManagedIdentity
	userAssignedIDEnabled := k8sConfig.UseManagedIdentity &&
		k8sConfig.UserAssignedID != ""
	isAzureCNI := orchProfile.IsAzureCNI()
	masterCount := masterProfile.Count

	var dependencies []string

	if isCustomVnet {
		dependencies = append(dependencies, "[variables('nsgID')]")
	} else {
		dependencies = append(dependencies, "[variables('vnetID')]")
	}

	if masterCount > 1 {
		dependencies = append(dependencies, "[variables('masterInternalLbName')]")
	}

	if to.Bool(masterProfile.CosmosEtcd) {
		dependencies = append(dependencies, "[resourceId('Microsoft.DocumentDB/databaseAccounts/', variables('cosmosAccountName'))]")
	}

	dependencies = append(dependencies, "[variables('masterLbID')]")

	armResource := ARMResource{
		APIVersion: "[variables('apiVersionCompute')]",
		DependsOn:  dependencies,
	}

	virtualMachine := compute.VirtualMachineScaleSet{
		Location: to.StringPtr("[variables('location')]"),
		Name:     to.StringPtr("[concat(variables('masterVMNamePrefix'), 'vmss')]"),
		Tags: map[string]*string{
			"creationSource":     to.StringPtr("[concat(parameters('generatorCode'), '-', variables('masterVMNamePrefix'), 'vmss']"),
			"resourceNameSuffix": to.StringPtr("[parameters('nameSuffix')]"),
			"orchestrator":       to.StringPtr("[variables('orchestratorNameVersionTag')]"),
			"aksEngineVersion":   to.StringPtr("[parameters('aksEngineVersion')]"),
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
		Capacity: to.Int64Ptr(int64(masterProfile.Count)),
		Name:     to.StringPtr("[parameters('masterVMSize')]"),
	}

	vmProperties := &compute.VirtualMachineScaleSetProperties{}

	vmProperties.SinglePlacementGroup = masterProfile.SinglePlacementGroup
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

	for i := 1; i <= masterProfile.Count; i++ {
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
			backendAddressPools := []compute.SubResource{
				{
					ID: to.StringPtr("[concat(variables('masterLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"),
				},
			}
			if masterCount > 1 {
				internalLbBackendAddressPool := compute.SubResource{
					ID: to.StringPtr("[concat(variables('masterInternalLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"),
				}
				backendAddressPools = append(backendAddressPools, internalLbBackendAddressPool)
			}
			ipConfigProps.LoadBalancerBackendAddressPools = &backendAddressPools

			ipConfigProps.LoadBalancerInboundNatPools = &[]compute.SubResource{
				{
					ID: to.StringPtr("[concat(variables('masterLbID'),'/inboundNatPools/SSH-', variables('masterVMNamePrefix'), 'natpools')]"),
				},
			}
		} else {
			ipConfigProps.Primary = to.BoolPtr(false)
		}
		ipConfig.VirtualMachineScaleSetIPConfigurationProperties = &ipConfigProps
		ipConfigurations = append(ipConfigurations, ipConfig)
	}
	netintconfig.IPConfigurations = &ipConfigurations

	if linuxProfile.HasCustomNodesDNS() {
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

	if linuxProfile.HasSecrets() {
		vsg := getVaultSecretGroup(cs.Properties.LinuxProfile)
		osProfile.Secrets = &vsg
	}

	storageProfile := compute.VirtualMachineScaleSetStorageProfile{}
	imageRef := masterProfile.ImageRef
	useMasterCustomImage := imageRef != nil && len(imageRef.Name) > 0 && len(imageRef.ResourceGroup) > 0
	if !useMasterCustomImage {
		etcdSizeGB, _ := strconv.Atoi(k8sConfig.EtcdDiskSizeGB)
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

	if masterProfile.OSDiskSizeGB > 0 {
		osDisk.DiskSizeGB = to.Int32Ptr(int32(masterProfile.OSDiskSizeGB))
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
	if !cs.Properties.FeatureFlags.IsFeatureEnabled("BlockOutboundInternet") {
		if cs.GetCloudSpecConfig().CloudName == api.AzureChinaCloud {
			registry = `gcr.azk8s.cn 80`
		} else {
			registry = `k8s.gcr.io 443 && retrycmd_if_failure 50 1 3 nc -vz gcr.io 443 && retrycmd_if_failure 50 1 3 nc -vz docker.io 443`
		}
		outBoundCmd = `ERR_OUTBOUND_CONN_FAIL=50; retrycmd_if_failure 50 1 3 nc -vz ` + registry + ` || exit $ERR_OUTBOUND_CONN_FAIL;`
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

	if cs.IsAKSBillingEnabled() {
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

	virtualMachine.VirtualMachineScaleSetProperties = vmProperties

	return VirtualMachineScaleSetARM{
		ARMResource:            armResource,
		VirtualMachineScaleSet: virtualMachine,
	}
}

func CreateAgentVMSS(cs *api.ContainerService, profile *api.AgentPoolProfile) VirtualMachineScaleSetARM {
	armResource := ARMResource{
		APIVersion: "[variables('apiVersionCompute')]",
	}
	var dependencies []string

	if profile.IsCustomVNET() {
		dependencies = append(dependencies, "[variables('nsgID')]")
	} else {
		dependencies = append(dependencies, "[variables('vnetID')]")
	}

	orchProfile := cs.Properties.OrchestratorProfile
	k8sConfig := orchProfile.KubernetesConfig
	linuxProfile := cs.Properties.LinuxProfile

	armResource.DependsOn = dependencies

	var resourceNameSuffix *string

	if profile.IsWindows() {
		resourceNameSuffix = to.StringPtr("[variables('winResourceNamePrefix')]")
	} else {
		resourceNameSuffix = to.StringPtr("[parameters('nameSuffix')]")
	}
	tags := map[string]*string{
		"creationSource":     to.StringPtr(fmt.Sprintf("[concat(parameters('generatorCode'), '-', variables('%sVMNamePrefix'))]", profile.Name)),
		"orchestrator":       to.StringPtr("[variables('orchestratorNameVersionTag')]"),
		"aksEngineVersion":   to.StringPtr("[parameters('aksEngineVersion')]"),
		"poolName":           to.StringPtr(profile.Name),
		"resourceNameSuffix": resourceNameSuffix,
	}

	virtualMachineScaleSet := compute.VirtualMachineScaleSet{
		Name:     to.StringPtr(fmt.Sprintf("[variables('%sVMNamePrefix')]", profile.Name)),
		Type:     to.StringPtr("Microsoft.Compute/virtualMachineScaleSets"),
		Location: to.StringPtr("[variables('location')]"),
		Sku: &compute.Sku{
			Tier:     to.StringPtr("Standard"),
			Capacity: to.Int64Ptr(int64(profile.Count)), //"[variables('{{.Name}}Count')]",
			Name:     to.StringPtr(fmt.Sprintf("[variables('%sVMSize')]", profile.Name)),
		},
		Tags: tags,
	}

	if profile.HasAvailabilityZones() {
		virtualMachineScaleSet.Zones = &[]string{
			fmt.Sprintf("[parameters('%sAvailabilityZones')]", profile.Name),
		}
	}

	useManagedIdentity := k8sConfig.UseManagedIdentity
	if useManagedIdentity {
		userAssignedIdentityEnabled := k8sConfig.UserAssignedID != ""
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

	if profile.IsWindows() {
		vmssNICConfig.EnableAcceleratedNetworking = profile.AcceleratedNetworkingEnabledWindows
	}

	if profile.IsCustomVNET() {
		vmssNICConfig.NetworkSecurityGroup = &compute.SubResource{
			ID: to.StringPtr("[variables('nsgID')]"),
		}
	}

	var ipConfigurations []compute.VirtualMachineScaleSetIPConfiguration

	for i := 1; i <= profile.IPAddressCount; i++ {
		ipconfig := compute.VirtualMachineScaleSetIPConfiguration{
			Name: to.StringPtr(fmt.Sprintf("ipconfig%d", i)),
			VirtualMachineScaleSetIPConfigurationProperties: &compute.VirtualMachineScaleSetIPConfigurationProperties{
				Subnet: &compute.APIEntityReference{
					ID: to.StringPtr(fmt.Sprintf("[variables('%sVnetSubnetID')]", profile.Name)),
				},
			},
		}
		if i == 1 {
			ipconfig.Primary = to.BoolPtr(true)
		}
		ipConfigurations = append(ipConfigurations, ipconfig)
	}

	vmssNICConfig.IPConfigurations = &ipConfigurations

	if linuxProfile.HasCustomNodesDNS() && !profile.IsWindows() {
		vmssNICConfig.DNSSettings = &compute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
			DNSServers: &[]string{
				"[parameters('dnsServer')]",
			},
		}
	}

	if !orchProfile.IsAzureCNI() {
		vmssNICConfig.EnableIPForwarding = to.BoolPtr(true)
	}

	vmssNetworkProfile := compute.VirtualMachineScaleSetNetworkProfile{
		NetworkInterfaceConfigurations: &[]compute.VirtualMachineScaleSetNetworkConfiguration{
			vmssNICConfig,
		},
	}

	vmssVMProfile.NetworkProfile = &vmssNetworkProfile

	t, err := InitializeTemplateGenerator(Context{})

	if err != nil {
		panic(err)
	}

	if profile.IsWindows() {

		customDataStr := getCustomDataFromJSON(t.GetKubernetesWindowsAgentCustomDataJSON(cs, profile))
		windowsOsProfile := compute.VirtualMachineScaleSetOSProfile{
			AdminUsername: to.StringPtr("[parameters('windowsAdminUsername')]"),
			AdminPassword: to.StringPtr("[parameters('windowsAdminPassword')]"),
			WindowsConfiguration: &compute.WindowsConfiguration{
				EnableAutomaticUpdates: to.BoolPtr(cs.Properties.WindowsProfile.GetEnableWindowsUpdate()),
			},
			CustomData: to.StringPtr(customDataStr),
		}
		vmssVMProfile.OsProfile = &windowsOsProfile
	} else {
		customDataStr := getCustomDataFromJSON(t.GetKubernetesAgentCustomDataJSON(cs, profile))
		linuxOsProfile := compute.VirtualMachineScaleSetOSProfile{
			AdminUsername:      to.StringPtr("[parameters('linuxAdminUsername')]"),
			ComputerNamePrefix: to.StringPtr(fmt.Sprintf("[variables('%sVMNamePrefix')]", profile.Name)),
			CustomData:         to.StringPtr(customDataStr),
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
			linuxOsProfile.LinuxConfiguration.SSH = &compute.SSHConfiguration{
				PublicKeys: &publicKeys,
			}

		} else {
			linuxOsProfile.LinuxConfiguration.SSH = &compute.SSHConfiguration{
				PublicKeys: &[]compute.SSHPublicKey{
					{
						KeyData: to.StringPtr("[parameters('sshRSAPublicKey')]"),
						Path:    to.StringPtr("[variables('sshKeyPath')]"),
					},
				},
			}
		}

		if linuxProfile.HasSecrets() {
			vsg := getVaultSecretGroup(cs.Properties.LinuxProfile)
			linuxOsProfile.Secrets = &vsg
		}

		vmssVMProfile.OsProfile = &linuxOsProfile
	}

	vmssStorageProfile := compute.VirtualMachineScaleSetStorageProfile{}

	if profile.IsWindows() {
		vmssStorageProfile.ImageReference = &compute.ImageReference{
			Offer:     to.StringPtr("[parameters('agentWindowsOffer')]"),
			Publisher: to.StringPtr("[parameters('agentWindowsPublisher')]"),
			Sku:       to.StringPtr("[parameters('agentWindowsSku')]"),
			Version:   to.StringPtr("[parameters('agentWindowsVersion')]"),
		}
		vmssStorageProfile.DataDisks = getVMSSDataDisks(profile)
	} else {
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

	featureFlags := cs.Properties.FeatureFlags

	if !featureFlags.IsFeatureEnabled("BlockOutboundInternet") {
		if cs.GetCloudSpecConfig().CloudName == api.AzureChinaCloud {
			registry = `gcr.azk8s.cn 80`
		} else {
			registry = `k8s.gcr.io 443 && retrycmd_if_failure 50 1 3 nc -vz gcr.io 443 && retrycmd_if_failure 50 1 3 nc -vz docker.io 443`
		}
		outBoundCmd = `ERR_OUTBOUND_CONN_FAIL=50; retrycmd_if_failure 50 1 3 nc -vz ` + registry + ` || exit $ERR_OUTBOUND_CONN_FAIL;`
	}

	var vmssCSE compute.VirtualMachineScaleSetExtension

	if profile.IsWindows() {
		vmssCSE = compute.VirtualMachineScaleSetExtension{
			Name: to.StringPtr("vmssCSE"),
			VirtualMachineScaleSetExtensionProperties: &compute.VirtualMachineScaleSetExtensionProperties{
				Publisher:               to.StringPtr("Microsoft.Compute"),
				Type:                    to.StringPtr("CustomScriptExtension"),
				TypeHandlerVersion:      to.StringPtr("1.8"),
				AutoUpgradeMinorVersion: to.BoolPtr(true),
				Settings:                map[string]interface{}{},
				ProtectedSettings: map[string]interface{}{
					"commandToExecute": "[concat('powershell.exe -ExecutionPolicy Unrestricted -command \"', '$arguments = ', variables('singleQuote'),'-MasterIP ',variables('kubernetesAPIServerIP'),' -KubeDnsServiceIp ',parameters('kubeDnsServiceIp'),' -MasterFQDNPrefix ',variables('masterFqdnPrefix'),' -Location ',variables('location'),' -AgentKey ',parameters('clientPrivateKey'),' -AADClientId ',variables('servicePrincipalClientId'),' -AADClientSecret ',variables('servicePrincipalClientSecret'),variables('singleQuote'), ' ; ', variables('windowsCustomScriptSuffix'), '\" > %SYSTEMDRIVE%\\AzureData\\CustomDataSetupScript.log 2>&1')]",
				},
			},
		}
	} else {
		runInBackground := ""
		if featureFlags.IsFeatureEnabled("CSERunInBackground") {
			runInBackground = " &"
		}
		nVidiaEnabled := strconv.FormatBool(common.IsNvidiaEnabledSKU(profile.VMSize))
		sgxEnabled := strconv.FormatBool(common.IsSgxEnabledSKU(profile.VMSize))

		commandExec := fmt.Sprintf("[concat('retrycmd_if_failure() { r=$1; w=$2; t=$3; shift && shift && shift; for i in $(seq 1 $r); do timeout $t ${@}; [ $? -eq 0  ] && break || if [ $i -eq $r ]; then return 1; else sleep $w; fi; done }; %s for i in $(seq 1 1200); do if [ -f /opt/azure/containers/provision.sh ]; then break; fi; if [ $i -eq 1200 ]; then exit 100; else sleep 1; fi; done; ', variables('provisionScriptParametersCommon'),' GPU_NODE=%s SGX_NODE=%s /usr/bin/nohup /bin/bash -c \"/bin/bash /opt/azure/containers/provision.sh >> /var/log/azure/cluster-provision.log 2>&1%s\"')]", outBoundCmd, nVidiaEnabled, sgxEnabled, runInBackground)
		vmssCSE = compute.VirtualMachineScaleSetExtension{
			Name: to.StringPtr("vmssCSE"),
			VirtualMachineScaleSetExtensionProperties: &compute.VirtualMachineScaleSetExtensionProperties{
				Publisher:               to.StringPtr("Microsoft.Azure.Extensions"),
				Type:                    to.StringPtr("CustomScript"),
				TypeHandlerVersion:      to.StringPtr("2.0"),
				AutoUpgradeMinorVersion: to.BoolPtr(true),
				Settings:                map[string]interface{}{},
				ProtectedSettings: map[string]interface{}{
					"commandToExecute": commandExec,
				},
			},
		}
	}

	vmssExtensions = append(vmssExtensions, vmssCSE)

	if cs.IsAKSBillingEnabled() {
		aksBillingExtension := compute.VirtualMachineScaleSetExtension{
			Name: to.StringPtr(fmt.Sprintf("[concat(variables('%sVMNamePrefix'), '-computeAksLinuxBilling')]", profile.Name)),
			VirtualMachineScaleSetExtensionProperties: &compute.VirtualMachineScaleSetExtensionProperties{
				Publisher:               to.StringPtr("Microsoft.AKS"),
				Type:                    to.StringPtr("Compute.AKS-Engine.Linux.Billing"),
				TypeHandlerVersion:      to.StringPtr("1.0"),
				AutoUpgradeMinorVersion: to.BoolPtr(true),
				Settings:                map[string]interface{}{},
			},
		}

		if profile.IsWindows() {
			aksBillingExtension.Type = to.StringPtr("Compute.AKS-Engine.Windows.Billing")
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
			Caching:      compute.CachingTypesReadOnly,
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
