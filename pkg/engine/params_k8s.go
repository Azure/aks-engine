// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"strconv"
	"strings"

	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
)

func assignKubernetesParameters(properties *api.Properties, parametersMap paramsMap,
	cloudSpecConfig api.AzureEnvironmentSpecConfig, generatorCode string) {
	addValue(parametersMap, "generatorCode", generatorCode)

	orchestratorProfile := properties.OrchestratorProfile

	if orchestratorProfile.IsKubernetes() {

		k8sVersion := orchestratorProfile.OrchestratorVersion
		k8sComponents := api.K8sComponentsByVersionMap[k8sVersion]
		kubernetesConfig := orchestratorProfile.KubernetesConfig
		kubernetesImageBase := kubernetesConfig.KubernetesImageBase
		mcrKubernetesImageBase := kubernetesConfig.MCRKubernetesImageBase
		hyperkubeImageBase := kubernetesConfig.KubernetesImageBase

		if properties.IsAzureStackCloud() {
			kubernetesImageBase = cloudSpecConfig.KubernetesSpecConfig.KubernetesImageBase
		}

		if kubernetesConfig != nil {
			if to.Bool(kubernetesConfig.UseCloudControllerManager) {
				kubernetesCcmSpec := kubernetesImageBase + k8sComponents["ccm"]
				if kubernetesConfig.CustomCcmImage != "" {
					kubernetesCcmSpec = kubernetesConfig.CustomCcmImage
				}

				addValue(parametersMap, "kubernetesCcmImageSpec", kubernetesCcmSpec)
			}

			kubernetesHyperkubeSpec := hyperkubeImageBase + k8sComponents["hyperkube"]
			if properties.IsAzureStackCloud() {
				kubernetesHyperkubeSpec = kubernetesHyperkubeSpec + AzureStackSuffix
			}
			if kubernetesConfig.CustomHyperkubeImage != "" {
				kubernetesHyperkubeSpec = kubernetesConfig.CustomHyperkubeImage
			}
			addValue(parametersMap, "kubernetesHyperkubeSpec", kubernetesHyperkubeSpec)

			addValue(parametersMap, "kubeDNSServiceIP", kubernetesConfig.DNSServiceIP)
			if kubernetesConfig.PrivateAzureRegistryServer != "" {
				addValue(parametersMap, "privateAzureRegistryServer", kubernetesConfig.PrivateAzureRegistryServer)
			}
			addValue(parametersMap, "kubernetesAddonManagerSpec", kubernetesImageBase+k8sComponents["addonmanager"])
			if orchestratorProfile.NeedsExecHealthz() {
				addValue(parametersMap, "kubernetesExecHealthzSpec", kubernetesImageBase+k8sComponents["exechealthz"])
			}
			addValue(parametersMap, "kubernetesDNSSidecarSpec", kubernetesImageBase+k8sComponents["k8s-dns-sidecar"])
			if kubernetesConfig.IsAADPodIdentityEnabled() {
				aadPodIdentityAddon := kubernetesConfig.GetAddonByName(AADPodIdentityAddonName)
				aadIndex := aadPodIdentityAddon.GetAddonContainersIndexByName(AADPodIdentityAddonName)
				if aadIndex > -1 {
					addValue(parametersMap, "kubernetesAADPodIdentityEnabled", to.Bool(aadPodIdentityAddon.Enabled))
				}
			}
			if kubernetesConfig.IsAddonEnabled(api.ACIConnectorAddonName) {
				addValue(parametersMap, "kubernetesACIConnectorEnabled", true)
			} else {
				addValue(parametersMap, "kubernetesACIConnectorEnabled", false)
			}
			if kubernetesConfig.IsAddonEnabled(api.ClusterAutoscalerAddonName) {
				clusterAutoscalerAddon := kubernetesConfig.GetAddonByName(ClusterAutoscalerAddonName)
				clusterAutoScalerIndex := clusterAutoscalerAddon.GetAddonContainersIndexByName(ClusterAutoscalerAddonName)
				if clusterAutoScalerIndex > -1 {
					addValue(parametersMap, "kubernetesClusterAutoscalerAzureCloud", cloudSpecConfig.CloudName)
					addValue(parametersMap, "kubernetesClusterAutoscalerEnabled", true)
					addValue(parametersMap, "kubernetesClusterAutoscalerUseManagedIdentity", strings.ToLower(strconv.FormatBool(kubernetesConfig.UseManagedIdentity)))
				}
			} else {
				addValue(parametersMap, "kubernetesClusterAutoscalerEnabled", false)
			}
			if common.IsKubernetesVersionGe(k8sVersion, "1.12.0") {
				addValue(parametersMap, "kubernetesCoreDNSSpec", kubernetesImageBase+k8sComponents["coredns"])
			} else {
				addValue(parametersMap, "kubernetesKubeDNSSpec", kubernetesImageBase+k8sComponents["kube-dns"])
				addValue(parametersMap, "kubernetesDNSMasqSpec", kubernetesImageBase+k8sComponents["dnsmasq"])
			}
			addValue(parametersMap, "kubernetesPodInfraContainerSpec", mcrKubernetesImageBase+k8sComponents["pause"])
			addValue(parametersMap, "cloudproviderConfig", api.CloudProviderConfig{
				CloudProviderBackoff:              kubernetesConfig.CloudProviderBackoff,
				CloudProviderBackoffRetries:       kubernetesConfig.CloudProviderBackoffRetries,
				CloudProviderBackoffJitter:        strconv.FormatFloat(kubernetesConfig.CloudProviderBackoffJitter, 'f', -1, 64),
				CloudProviderBackoffDuration:      kubernetesConfig.CloudProviderBackoffDuration,
				CloudProviderBackoffExponent:      strconv.FormatFloat(kubernetesConfig.CloudProviderBackoffExponent, 'f', -1, 64),
				CloudProviderRateLimit:            kubernetesConfig.CloudProviderRateLimit,
				CloudProviderRateLimitQPS:         strconv.FormatFloat(kubernetesConfig.CloudProviderRateLimitQPS, 'f', -1, 64),
				CloudProviderRateLimitQPSWrite:    strconv.FormatFloat(kubernetesConfig.CloudProviderRateLimitQPSWrite, 'f', -1, 64),
				CloudProviderRateLimitBucket:      kubernetesConfig.CloudProviderRateLimitBucket,
				CloudProviderRateLimitBucketWrite: kubernetesConfig.CloudProviderRateLimitBucketWrite,
			})
			addValue(parametersMap, "kubeClusterCidr", kubernetesConfig.ClusterSubnet)
			addValue(parametersMap, "kubernetesKubeletClusterDomain", kubernetesConfig.KubeletConfig["--cluster-domain"])
			addValue(parametersMap, "dockerBridgeCidr", kubernetesConfig.DockerBridgeSubnet)
			addValue(parametersMap, "networkPolicy", kubernetesConfig.NetworkPolicy)
			addValue(parametersMap, "networkPlugin", kubernetesConfig.NetworkPlugin)
			addValue(parametersMap, "containerRuntime", kubernetesConfig.ContainerRuntime)
			addValue(parametersMap, "containerdDownloadURLBase", cloudSpecConfig.KubernetesSpecConfig.ContainerdDownloadURLBase)
			addValue(parametersMap, "cniPluginsURL", cloudSpecConfig.KubernetesSpecConfig.CNIPluginsDownloadURL)
			addValue(parametersMap, "vnetCniLinuxPluginsURL", kubernetesConfig.GetAzureCNIURLLinux(cloudSpecConfig))
			addValue(parametersMap, "vnetCniWindowsPluginsURL", kubernetesConfig.GetAzureCNIURLWindows(cloudSpecConfig))
			addValue(parametersMap, "gchighthreshold", kubernetesConfig.GCHighThreshold)
			addValue(parametersMap, "gclowthreshold", kubernetesConfig.GCLowThreshold)
			addValue(parametersMap, "etcdDownloadURLBase", cloudSpecConfig.KubernetesSpecConfig.EtcdDownloadURLBase)
			addValue(parametersMap, "etcdVersion", kubernetesConfig.EtcdVersion)
			addValue(parametersMap, "etcdDiskSizeGB", kubernetesConfig.EtcdDiskSizeGB)
			addValue(parametersMap, "etcdEncryptionKey", kubernetesConfig.EtcdEncryptionKey)
			if kubernetesConfig.PrivateJumpboxProvision() {
				addValue(parametersMap, "jumpboxVMName", kubernetesConfig.PrivateCluster.JumpboxProfile.Name)
				addValue(parametersMap, "jumpboxVMSize", kubernetesConfig.PrivateCluster.JumpboxProfile.VMSize)
				addValue(parametersMap, "jumpboxUsername", kubernetesConfig.PrivateCluster.JumpboxProfile.Username)
				addValue(parametersMap, "jumpboxOSDiskSizeGB", kubernetesConfig.PrivateCluster.JumpboxProfile.OSDiskSizeGB)
				addValue(parametersMap, "jumpboxPublicKey", kubernetesConfig.PrivateCluster.JumpboxProfile.PublicKey)
				addValue(parametersMap, "jumpboxStorageProfile", kubernetesConfig.PrivateCluster.JumpboxProfile.StorageProfile)
			}

			addValue(parametersMap, "enableAggregatedAPIs", kubernetesConfig.EnableAggregatedAPIs)

			if properties.HasWindows() {
				// Kubernetes packages as zip file as created by scripts/build-windows-k8s.sh
				// will be removed in future release as if gets phased out (https://github.com/Azure/aks-engine/issues/3851)
				kubeBinariesSASURL := kubernetesConfig.CustomWindowsPackageURL
				if kubeBinariesSASURL == "" {
					if properties.IsAzureStackCloud() {
						kubeBinariesSASURL = cloudSpecConfig.KubernetesSpecConfig.KubeBinariesSASURLBase + AzureStackPrefix + k8sComponents["windowszip"]
					} else {
						kubeBinariesSASURL = cloudSpecConfig.KubernetesSpecConfig.KubeBinariesSASURLBase + k8sComponents["windowszip"]
					}
				}
				addValue(parametersMap, "kubeBinariesSASURL", kubeBinariesSASURL)

				// Kubernetes node binaries as packaged by upstream kubernetes
				// example at https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG-1.11.md#node-binaries-1
				addValue(parametersMap, "windowsKubeBinariesURL", kubernetesConfig.WindowsNodeBinariesURL)
				addValue(parametersMap, "kubeServiceCidr", kubernetesConfig.ServiceCIDR)
				addValue(parametersMap, "kubeBinariesVersion", k8sVersion)
				addValue(parametersMap, "windowsTelemetryGUID", cloudSpecConfig.KubernetesSpecConfig.WindowsTelemetryGUID)
			}
		}

		if kubernetesConfig == nil ||
			!kubernetesConfig.UseManagedIdentity ||
			properties.IsHostedMasterProfile() {
			servicePrincipalProfile := properties.ServicePrincipalProfile

			if servicePrincipalProfile != nil {
				addValue(parametersMap, "servicePrincipalClientId", servicePrincipalProfile.ClientID)
				keyVaultSecretRef := servicePrincipalProfile.KeyvaultSecretRef
				if keyVaultSecretRef != nil {
					addKeyvaultReference(parametersMap, "servicePrincipalClientSecret",
						keyVaultSecretRef.VaultID,
						keyVaultSecretRef.SecretName,
						keyVaultSecretRef.SecretVersion)
				} else {
					addValue(parametersMap, "servicePrincipalClientSecret", servicePrincipalProfile.Secret)
				}

				if kubernetesConfig != nil && to.Bool(kubernetesConfig.EnableEncryptionWithExternalKms) {
					if kubernetesConfig.KeyVaultSku != "" {
						addValue(parametersMap, "clusterKeyVaultSku", kubernetesConfig.KeyVaultSku)
					}
					if !kubernetesConfig.UseManagedIdentity && servicePrincipalProfile.ObjectID != "" {
						addValue(parametersMap, "servicePrincipalObjectId", servicePrincipalProfile.ObjectID)
					}
				}
			}
		}

		addValue(parametersMap, "orchestratorName", properties.K8sOrchestratorName())

		/**
		 The following parameters could be either a plain text, or referenced to a secret in a keyvault:
		 - apiServerCertificate
		 - apiServerPrivateKey
		 - caCertificate
		 - clientCertificate
		 - clientPrivateKey
		 - kubeConfigCertificate
		 - kubeConfigPrivateKey
		 - servicePrincipalClientSecret
		 - etcdClientCertificate
		 - etcdClientPrivateKey
		 - etcdServerCertificate
		 - etcdServerPrivateKey
		 - etcdPeerCertificates
		 - etcdPeerPrivateKeys

		 To refer to a keyvault secret, the value of the parameter in the api model file should be formatted as:

		 "<PARAMETER>": "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>/secrets/<NAME>[/<VERSION>]"
		 where:
		   <SUB_ID> is the subscription ID of the keyvault
		   <RG_NAME> is the resource group of the keyvault
		   <KV_NAME> is the name of the keyvault
		   <NAME> is the name of the secret.
		   <VERSION> (optional) is the version of the secret (default: the latest version)

		 This will generate a reference block in the parameters file:

		 "reference": {
		   "keyVault": {
		     "id": "/subscriptions/<SUB_ID>/resourceGroups/<RG_NAME>/providers/Microsoft.KeyVault/vaults/<KV_NAME>"
		   },
		   "secretName": "<NAME>"
		   "secretVersion": "<VERSION>"
		}
		**/

		certificateProfile := properties.CertificateProfile
		if certificateProfile != nil {
			addSecret(parametersMap, "apiServerCertificate", certificateProfile.APIServerCertificate, true)
			addSecret(parametersMap, "apiServerPrivateKey", certificateProfile.APIServerPrivateKey, true)
			addSecret(parametersMap, "caCertificate", certificateProfile.CaCertificate, true)
			addSecret(parametersMap, "caPrivateKey", certificateProfile.CaPrivateKey, true)
			addSecret(parametersMap, "clientCertificate", certificateProfile.ClientCertificate, true)
			addSecret(parametersMap, "clientPrivateKey", certificateProfile.ClientPrivateKey, true)
			addSecret(parametersMap, "kubeConfigCertificate", certificateProfile.KubeConfigCertificate, true)
			addSecret(parametersMap, "kubeConfigPrivateKey", certificateProfile.KubeConfigPrivateKey, true)
			if properties.MasterProfile != nil {
				addSecret(parametersMap, "etcdServerCertificate", certificateProfile.EtcdServerCertificate, true)
				addSecret(parametersMap, "etcdServerPrivateKey", certificateProfile.EtcdServerPrivateKey, true)
				addSecret(parametersMap, "etcdClientCertificate", certificateProfile.EtcdClientCertificate, true)
				addSecret(parametersMap, "etcdClientPrivateKey", certificateProfile.EtcdClientPrivateKey, true)
				for i, pc := range certificateProfile.EtcdPeerCertificates {
					addSecret(parametersMap, "etcdPeerCertificate"+strconv.Itoa(i), pc, true)
				}
				for i, pk := range certificateProfile.EtcdPeerPrivateKeys {
					addSecret(parametersMap, "etcdPeerPrivateKey"+strconv.Itoa(i), pk, true)
				}
			}
		}

		if properties.HostedMasterProfile != nil && properties.HostedMasterProfile.FQDN != "" {
			addValue(parametersMap, "kubernetesEndpoint", properties.HostedMasterProfile.FQDN)
		}

		if properties.OrchestratorProfile.KubernetesConfig.MobyVersion != "" {
			addValue(parametersMap, "mobyVersion", properties.OrchestratorProfile.KubernetesConfig.MobyVersion)
		}

		if properties.OrchestratorProfile.KubernetesConfig.ContainerdVersion != "" {
			addValue(parametersMap, "containerdVersion", properties.OrchestratorProfile.KubernetesConfig.ContainerdVersion)
		}

		if properties.AADProfile != nil {
			addValue(parametersMap, "aadTenantId", properties.AADProfile.TenantID)
			if properties.AADProfile.AdminGroupID != "" {
				addValue(parametersMap, "aadAdminGroupId", properties.AADProfile.AdminGroupID)
			}
		}

		if kubernetesConfig != nil && kubernetesConfig.IsAddonEnabled(AppGwIngressAddonName) {
			addValue(parametersMap, "appGwSku", kubernetesConfig.GetAddonByName(AppGwIngressAddonName).Config["appgw-sku"])
			addValue(parametersMap, "appGwSubnet", kubernetesConfig.GetAddonByName(AppGwIngressAddonName).Config["appgw-subnet"])
		}
	}
}
