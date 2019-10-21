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

	var useManagedIdentity, userAssignedIDEnabled bool
	if k8sConfig != nil {
		useManagedIdentity = k8sConfig.UseManagedIdentity
		userAssignedIDEnabled = useManagedIdentity && k8sConfig.UserAssignedID != ""
	}
	isAzureCNI := orchProfile.IsAzureCNI()
	masterCount := masterProfile.Count
	isVHD := strconv.FormatBool(masterProfile.IsVHDDistro())

	var dependencies []string

	if isCustomVnet {
		dependencies = append(dependencies, "[variables('nsgID')]")
	} else {
		dependencies = append(dependencies, "[variables('vnetID')]")
	}

	if masterCount > 1 {
		dependencies = append(dependencies, "[variables('masterInternalLbName')]")
	}

	if masterProfile.HasCosmosEtcd() {
		dependencies = append(dependencies, "[resourceId('Microsoft.DocumentDB/databaseAccounts/', variables('cosmosAccountName'))]")
	}

	if !cs.Properties.OrchestratorProfile.IsPrivateCluster() {
		dependencies = append(dependencies, "[variables('masterLbID')]")
	}

	armResource := ARMResource{
		APIVersion: "[variables('apiVersionCompute')]",
		DependsOn:  dependencies,
	}

	vmScaleSetTags := map[string]*string{
		"creationSource":     to.StringPtr("[concat(parameters('generatorCode'), '-', variables('masterVMNamePrefix'), 'vmss')]"),
		"resourceNameSuffix": to.StringPtr("[parameters('nameSuffix')]"),
		"orchestrator":       to.StringPtr("[variables('orchestratorNameVersionTag')]"),
		"aksEngineVersion":   to.StringPtr("[parameters('aksEngineVersion')]"),
		"poolName":           to.StringPtr("master"),
	}

	if k8sConfig != nil && k8sConfig.IsContainerMonitoringAddonEnabled() {
		addon := k8sConfig.GetAddonByName(ContainerMonitoringAddonName)
		clusterDNSPrefix := "aks-engine-cluster"
		if cs.Properties.MasterProfile != nil && cs.Properties.MasterProfile.DNSPrefix != "" {
			clusterDNSPrefix = cs.Properties.MasterProfile.DNSPrefix
		}
		vmScaleSetTags["logAnalyticsWorkspaceResourceId"] = to.StringPtr(addon.Config["logAnalyticsWorkspaceResourceId"])
		vmScaleSetTags["clusterName"] = to.StringPtr(clusterDNSPrefix)
	}

	virtualMachine := compute.VirtualMachineScaleSet{
		Location: to.StringPtr("[variables('location')]"),
		Name:     to.StringPtr("[concat(variables('masterVMNamePrefix'), 'vmss')]"),
		Tags:     vmScaleSetTags,
		Type:     to.StringPtr("Microsoft.Compute/virtualMachineScaleSets"),
	}

	addCustomTagsToVMScaleSets(cs.Properties.MasterProfile.CustomVMTags, &virtualMachine)

	if hasAvailabilityZones {
		virtualMachine.Zones = &masterProfile.AvailabilityZones
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

	for i := 1; i <= masterProfile.IPAddressCount; i++ {
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
			backendAddressPools := []compute.SubResource{}
			if !cs.Properties.OrchestratorProfile.IsPrivateCluster() {
				publicBackendAddressPools := compute.SubResource{
					ID: to.StringPtr("[concat(variables('masterLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"),
				}
				backendAddressPools = append(backendAddressPools, publicBackendAddressPools)
				ipConfigProps.LoadBalancerInboundNatPools = &[]compute.SubResource{
					{
						ID: to.StringPtr("[concat(variables('masterLbID'),'/inboundNatPools/SSH-', variables('masterVMNamePrefix'), 'natpools')]"),
					},
				}
			}
			if masterCount > 1 {
				internalLbBackendAddressPool := compute.SubResource{
					ID: to.StringPtr("[concat(variables('masterInternalLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"),
				}
				backendAddressPools = append(backendAddressPools, internalLbBackendAddressPool)
			}
			ipConfigProps.LoadBalancerBackendAddressPools = &backendAddressPools

		} else {
			ipConfigProps.Primary = to.BoolPtr(false)
		}
		ipConfig.VirtualMachineScaleSetIPConfigurationProperties = &ipConfigProps
		ipConfigurations = append(ipConfigurations, ipConfig)
	}
	netintconfig.IPConfigurations = &ipConfigurations

	if linuxProfile != nil && linuxProfile.HasCustomNodesDNS() {
		netintconfig.DNSSettings = &compute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
			DNSServers: &[]string{
				"[parameters('dnsServer')]",
			},
		}
	}

	if !isAzureCNI && !cs.Properties.IsAzureStackCloud() {
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

	storageProfile := compute.VirtualMachineScaleSetStorageProfile{}
	imageRef := masterProfile.ImageRef
	etcdSizeGB, _ := strconv.Atoi(k8sConfig.EtcdDiskSizeGB)
	dataDisk := compute.VirtualMachineScaleSetDataDisk{
		CreateOption: compute.DiskCreateOptionTypesEmpty,
		DiskSizeGB:   to.Int32Ptr(int32(etcdSizeGB)),
		Lun:          to.Int32Ptr(0),
	}
	storageProfile.DataDisks = &[]compute.VirtualMachineScaleSetDataDisk{
		dataDisk,
	}
	imgReference := &compute.ImageReference{}
	if masterProfile.HasImageRef() {
		if masterProfile.HasImageGallery() {
			imgReference.ID = to.StringPtr(fmt.Sprintf("[concat('/subscriptions/', '%s',  '/resourceGroups/', parameters('osImageResourceGroup'), '/providers/Microsoft.Compute/galleries/', '%s', '/images/', parameters('osImageName'), '/versions/', '%s')]", imageRef.SubscriptionID, imageRef.Gallery, imageRef.Version))
		} else {
			imgReference.ID = to.StringPtr("[resourceId(parameters('osImageResourceGroup'), 'Microsoft.Compute/images', parameters('osImageName'))]")
		}
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
	ncBinary := "nc"
	if cs.Properties.MasterProfile != nil && cs.Properties.MasterProfile.IsCoreOS() {
		ncBinary = "ncat"
	}
	// TODO The AzureStack constraint has to be relaxed, it should only apply to *disconnected* instances
	if !cs.Properties.FeatureFlags.IsFeatureEnabled("BlockOutboundInternet") && !cs.Properties.IsAzureStackCloud() {
		if cs.GetCloudSpecConfig().CloudName == api.AzureChinaCloud {
			registry = `gcr.azk8s.cn 443`
		} else {
			registry = `aksrepos.azurecr.io 443`
		}
		outBoundCmd = `ERR_OUTBOUND_CONN_FAIL=50; retrycmd_if_failure 50 1 3 ` + ncBinary + ` -vz ` + registry + ` || exit $ERR_OUTBOUND_CONN_FAIL;`
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
				"commandToExecute": fmt.Sprintf("[concat('echo $(date),$(hostname); retrycmd_if_failure() { r=$1; w=$2; t=$3; shift && shift && shift; for i in $(seq 1 $r); do timeout $t ${@}; [ $? -eq 0  ] && break || if [ $i -eq $r ]; then return 1; else sleep $w; fi; done }; "+outBoundCmd+" for i in $(seq 1 1200); do grep -Fq \"EOF\" /opt/azure/containers/provision.sh && break; if [ $i -eq 1200 ]; then exit 100; else sleep 1; fi; done; ', variables('provisionScriptParametersCommon'),%s,variables('provisionScriptParametersMaster'), ' IS_VHD=%s /usr/bin/nohup /bin/bash -c \"/bin/bash /opt/azure/containers/provision.sh >> /var/log/azure/cluster-provision.log 2>&1\"')]", generateUserAssignedIdentityClientIDParameter(userAssignedIDEnabled), isVHD),
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
	isHostedMaster := cs.Properties.IsHostedMasterProfile()
	armResource := ARMResource{
		APIVersion: "[variables('apiVersionCompute')]",
	}
	var dependencies []string

	if profile.IsCustomVNET() {
		dependencies = append(dependencies, "[variables('nsgID')]")
	} else {
		dependencies = append(dependencies, "[variables('vnetID')]")
	}

	if !cs.Properties.OrchestratorProfile.IsPrivateCluster() &&
		profile.LoadBalancerBackendAddressPoolIDs == nil &&
		cs.Properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku == api.StandardLoadBalancerSku &&
		!isHostedMaster {
		dependencies = append(dependencies, "[variables('agentLbID')]")
	}

	if profile.IsWindows() {
		windowsProfile := cs.Properties.WindowsProfile
		// Add dependency for Image resource created by createWindowsImage()
		if windowsProfile.HasCustomImage() {
			dependencies = append(dependencies, fmt.Sprintf("%sCustomWindowsImage", profile.Name))
		}
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

	addCustomTagsToVMScaleSets(profile.CustomVMTags, &virtualMachineScaleSet)

	if profile.HasAvailabilityZones() {
		virtualMachineScaleSet.Zones = &profile.AvailabilityZones
	}

	var useManagedIdentity bool
	if k8sConfig != nil {
		useManagedIdentity = k8sConfig.UseManagedIdentity
	}
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
		Overprovision:        profile.VMSSOverProvisioningEnabled,
		UpgradePolicy: &compute.UpgradePolicy{
			Mode: compute.Manual,
		},
	}

	if to.Bool(profile.VMSSOverProvisioningEnabled) {
		vmssProperties.DoNotRunExtensionsOnOverprovisionedVMs = to.BoolPtr(true)
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
		}
		ipConfigProps := compute.VirtualMachineScaleSetIPConfigurationProperties{
			Subnet: &compute.APIEntityReference{
				ID: to.StringPtr(fmt.Sprintf("[variables('%sVnetSubnetID')]", profile.Name)),
			},
		}

		if i == 1 {
			ipConfigProps.Primary = to.BoolPtr(true)

			backendAddressPools := []compute.SubResource{}
			if profile.LoadBalancerBackendAddressPoolIDs != nil {
				for _, lbBackendPoolID := range profile.LoadBalancerBackendAddressPoolIDs {
					backendAddressPools = append(backendAddressPools,
						compute.SubResource{
							ID: to.StringPtr(lbBackendPoolID),
						},
					)
				}
			} else {
				if !cs.Properties.OrchestratorProfile.IsPrivateCluster() &&
					cs.Properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku == api.StandardLoadBalancerSku &&
					!isHostedMaster {
					agentLbBackendAddressPools := compute.SubResource{
						ID: to.StringPtr("[concat(variables('agentLbID'), '/backendAddressPools/', variables('agentLbBackendPoolName'))]"),
					}
					backendAddressPools = append(backendAddressPools, agentLbBackendAddressPools)
				}
			}

			ipConfigProps.LoadBalancerBackendAddressPools = &backendAddressPools
			if cs.Properties.FeatureFlags.IsFeatureEnabled("EnableIPv6DualStack") {
				if cs.Properties.OrchestratorProfile.KubernetesConfig.LoadBalancerSku != StandardLoadBalancerSku {
					defaultIPv4BackendPool := compute.SubResource{
						ID: to.StringPtr("[concat(resourceId('Microsoft.Network/loadBalancers',parameters('masterEndpointDNSNamePrefix')), '/backendAddressPools/', parameters('masterEndpointDNSNamePrefix'))]"),
					}
					backendPools := make([]compute.SubResource, 0)
					if ipConfigProps.LoadBalancerBackendAddressPools != nil {
						backendPools = *ipConfigProps.LoadBalancerBackendAddressPools
					}
					backendPools = append(backendPools, defaultIPv4BackendPool)
					ipConfigProps.LoadBalancerBackendAddressPools = &backendPools
				}
			}

			// Set VMSS node public IP if requested
			if to.Bool(profile.EnableVMSSNodePublicIP) {
				publicIPAddressConfiguration := &compute.VirtualMachineScaleSetPublicIPAddressConfiguration{
					Name: to.StringPtr(fmt.Sprintf("pub%d", i)),
					VirtualMachineScaleSetPublicIPAddressConfigurationProperties: &compute.VirtualMachineScaleSetPublicIPAddressConfigurationProperties{
						IdleTimeoutInMinutes: to.Int32Ptr(30),
					},
				}
				ipConfigProps.PublicIPAddressConfiguration = publicIPAddressConfiguration
			}
		}
		ipconfig.VirtualMachineScaleSetIPConfigurationProperties = &ipConfigProps
		ipConfigurations = append(ipConfigurations, ipconfig)

		if cs.Properties.FeatureFlags.IsFeatureEnabled("EnableIPv6DualStack") {
			ipconfigv6 := compute.VirtualMachineScaleSetIPConfiguration{
				Name: to.StringPtr(fmt.Sprintf("ipconfig%dv6", i)),
				VirtualMachineScaleSetIPConfigurationProperties: &compute.VirtualMachineScaleSetIPConfigurationProperties{
					Subnet: &compute.APIEntityReference{
						ID: to.StringPtr(fmt.Sprintf("[variables('%sVnetSubnetID')]", profile.Name)),
					},
					Primary:                 to.BoolPtr(false),
					PrivateIPAddressVersion: "IPv6",
				},
			}
			ipConfigurations = append(ipConfigurations, ipconfigv6)
		}
	}

	vmssNICConfig.IPConfigurations = &ipConfigurations

	if linuxProfile != nil && linuxProfile.HasCustomNodesDNS() && !profile.IsWindows() {
		vmssNICConfig.DNSSettings = &compute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
			DNSServers: &[]string{
				"[parameters('dnsServer')]",
			},
		}
	}

	if !orchProfile.IsAzureCNI() && !cs.Properties.IsAzureStackCloud() {
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

		customDataStr := getCustomDataFromJSON(t.GetKubernetesWindowsNodeCustomDataJSONObject(cs, profile))
		windowsOsProfile := compute.VirtualMachineScaleSetOSProfile{
			AdminUsername:      to.StringPtr("[parameters('windowsAdminUsername')]"),
			AdminPassword:      to.StringPtr("[parameters('windowsAdminPassword')]"),
			ComputerNamePrefix: to.StringPtr(fmt.Sprintf("[variables('%sVMNamePrefix')]", profile.Name)),
			WindowsConfiguration: &compute.WindowsConfiguration{
				EnableAutomaticUpdates: to.BoolPtr(cs.Properties.WindowsProfile.GetEnableWindowsUpdate()),
			},
			CustomData: to.StringPtr(customDataStr),
		}
		vmssVMProfile.OsProfile = &windowsOsProfile
	} else {
		customDataStr := getCustomDataFromJSON(t.GetKubernetesLinuxNodeCustomDataJSONObject(cs, profile))
		linuxOsProfile := compute.VirtualMachineScaleSetOSProfile{
			AdminUsername:      to.StringPtr("[parameters('linuxAdminUsername')]"),
			ComputerNamePrefix: to.StringPtr(fmt.Sprintf("[variables('%sVMNamePrefix')]", profile.Name)),
			CustomData:         to.StringPtr(customDataStr),
			LinuxConfiguration: &compute.LinuxConfiguration{
				DisablePasswordAuthentication: to.BoolPtr(true),
			},
		}

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

		if linuxProfile != nil && linuxProfile.HasSecrets() {
			vsg := getVaultSecretGroup(linuxProfile)
			linuxOsProfile.Secrets = &vsg
		}

		vmssVMProfile.OsProfile = &linuxOsProfile
	}

	vmssStorageProfile := compute.VirtualMachineScaleSetStorageProfile{}

	if profile.IsWindows() {
		vmssStorageProfile.ImageReference = createWindowsImageReference(profile.Name, cs.Properties.WindowsProfile)
		vmssStorageProfile.DataDisks = getVMSSDataDisks(profile)
	} else {
		if profile.HasImageRef() {
			imageRef := profile.ImageRef
			if profile.HasImageGallery() {
				v := fmt.Sprintf("[concat('/subscriptions/', '%s', '/resourceGroups/', variables('%sosImageResourceGroup'), '/providers/Microsoft.Compute/galleries/', '%s', '/images/', variables('%sosImageName'), '/versions/', '%s')]", imageRef.SubscriptionID, profile.Name, imageRef.Gallery, profile.Name, imageRef.Version)
				vmssStorageProfile.ImageReference = &compute.ImageReference{
					ID: to.StringPtr(v),
				}
			} else {
				vmssStorageProfile.ImageReference = &compute.ImageReference{
					ID: to.StringPtr(fmt.Sprintf("[resourceId(variables('%[1]sosImageResourceGroup'), 'Microsoft.Compute/images', variables('%[1]sosImageName'))]", profile.Name)),
				}
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

	if profile.IsEphemeral() {
		osDisk.Caching = compute.CachingTypesReadOnly
		osDisk.DiffDiskSettings = &compute.DiffDiskSettings{
			Option: compute.Local,
		}
	}

	vmssStorageProfile.OsDisk = &osDisk

	vmssVMProfile.StorageProfile = &vmssStorageProfile

	var vmssExtensions []compute.VirtualMachineScaleSetExtension

	outBoundCmd := ""
	registry := ""
	ncBinary := "nc"
	if profile.IsCoreOS() {
		ncBinary = "ncat"
	}
	featureFlags := cs.Properties.FeatureFlags

	if !featureFlags.IsFeatureEnabled("BlockOutboundInternet") {
		if cs.GetCloudSpecConfig().CloudName == api.AzureChinaCloud {
			registry = `gcr.azk8s.cn 443`
		} else {
			registry = `aksrepos.azurecr.io 443`
		}
		outBoundCmd = `ERR_OUTBOUND_CONN_FAIL=50; retrycmd_if_failure 50 1 3 ` + ncBinary + ` -vz ` + registry + ` || exit $ERR_OUTBOUND_CONN_FAIL;`
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
					"commandToExecute": "[concat('echo %DATE%,%TIME%,%COMPUTERNAME% && powershell.exe -ExecutionPolicy Unrestricted -command \"', '$arguments = ', variables('singleQuote'),'-MasterIP ',variables('kubernetesAPIServerIP'),' -KubeDnsServiceIp ',parameters('kubeDnsServiceIp'),' -MasterFQDNPrefix ',variables('masterFqdnPrefix'),' -Location ',variables('location'),' -TargetEnvironment ',parameters('targetEnvironment'),' -AgentKey ',parameters('clientPrivateKey'),' -AADClientId ',variables('servicePrincipalClientId'),' -AADClientSecret ',variables('singleQuote'),variables('singleQuote'),base64(variables('servicePrincipalClientSecret')),variables('singleQuote'),variables('singleQuote'),' -NetworkAPIVersion ',variables('apiVersionNetwork'),' ',variables('singleQuote'), ' ; ', variables('windowsCustomScriptSuffix'), '\" > %SYSTEMDRIVE%\\AzureData\\CustomDataSetupScript.log 2>&1 ; exit $LASTEXITCODE')]",
				},
			},
		}
	} else {
		runInBackground := ""
		if featureFlags.IsFeatureEnabled("CSERunInBackground") {
			runInBackground = " &"
		}
		var userAssignedIDEnabled bool
		if cs.Properties.OrchestratorProfile != nil && cs.Properties.OrchestratorProfile.KubernetesConfig != nil {
			userAssignedIDEnabled = cs.Properties.OrchestratorProfile.KubernetesConfig.UserAssignedIDEnabled()
		} else {
			userAssignedIDEnabled = false
		}
		nVidiaEnabled := strconv.FormatBool(common.IsNvidiaEnabledSKU(profile.VMSize))
		sgxEnabled := strconv.FormatBool(common.IsSgxEnabledSKU(profile.VMSize))
		auditDEnabled := strconv.FormatBool(to.Bool(profile.AuditDEnabled))
		isVHD := strconv.FormatBool(profile.IsVHDDistro())

		commandExec := fmt.Sprintf("[concat('echo $(date),$(hostname); retrycmd_if_failure() { r=$1; w=$2; t=$3; shift && shift && shift; for i in $(seq 1 $r); do timeout $t ${@}; [ $? -eq 0  ] && break || if [ $i -eq $r ]; then return 1; else sleep $w; fi; done }; %s for i in $(seq 1 1200); do grep -Fq \"EOF\" /opt/azure/containers/provision.sh && break; if [ $i -eq 1200 ]; then exit 100; else sleep 1; fi; done; ', variables('provisionScriptParametersCommon'),%s,' IS_VHD=%s GPU_NODE=%s SGX_NODE=%s AUDITD_ENABLED=%s /usr/bin/nohup /bin/bash -c \"/bin/bash /opt/azure/containers/provision.sh >> /var/log/azure/cluster-provision.log 2>&1%s\"')]", outBoundCmd, generateUserAssignedIdentityClientIDParameter(userAssignedIDEnabled), isVHD, nVidiaEnabled, sgxEnabled, auditDEnabled, runInBackground)
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

		if cs.Properties.IsHostedMasterProfile() {
			if profile.IsWindows() {
				aksBillingExtension.Name = to.StringPtr(fmt.Sprintf("[concat(variables('%sVMNamePrefix'), '-AKSWindowsBilling')]", profile.Name))
				aksBillingExtension.Type = to.StringPtr("Compute.AKS.Windows.Billing")
			} else {
				aksBillingExtension.Name = to.StringPtr(fmt.Sprintf("[concat(variables('%sVMNamePrefix'), '-AKSLinuxBilling')]", profile.Name))
				aksBillingExtension.Type = to.StringPtr("Compute.AKS.Linux.Billing")
			}
		} else {
			if profile.IsWindows() {
				aksBillingExtension.Type = to.StringPtr("Compute.AKS-Engine.Windows.Billing")
			} else {
				aksBillingExtension.Type = to.StringPtr("Compute.AKS-Engine.Linux.Billing")
			}
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
		}
		dataDisks = append(dataDisks, dataDisk)
	}
	return &dataDisks
}

func addCustomTagsToVMScaleSets(tags map[string]string, vm *compute.VirtualMachineScaleSet) {
	for key, value := range tags {
		_, found := vm.Tags[key]
		if !found {
			vm.Tags[key] = to.StringPtr(value)
		}
	}
}
