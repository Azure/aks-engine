package armhelpers

import (
	"fmt"
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/engine"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
	"strconv"
)

type VirtualMachineARM struct {
	ARMResource
	compute.VirtualMachine
}

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
