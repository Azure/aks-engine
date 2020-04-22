// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/pkg/errors"
)

// CustomCloudProfileDefaultsParams is the parameters when we set the cloud profile defaults for ContainerService.
type CustomCloudProfileDefaultsParams struct {
	IsUpgrade bool
	IsScale   bool
}

func getAzureStackFQDNSuffix(portalURL, location string) string {
	azsFQDNSuffix := strings.Replace(portalURL, fmt.Sprintf("https://portal.%s.", location), "", -1)
	azsFQDNSuffix = strings.TrimSuffix(azsFQDNSuffix, "/")

	return azsFQDNSuffix
}

func (cs *ContainerService) setCustomCloudProfileDefaults(params CustomCloudProfileDefaultsParams) error {
	p := cs.Properties
	if p.IsCustomCloudProfile() {
		p.CustomCloudProfile.AuthenticationMethod = helpers.EnsureString(p.CustomCloudProfile.AuthenticationMethod, ClientSecretAuthMethod)
		p.CustomCloudProfile.IdentitySystem = helpers.EnsureString(p.CustomCloudProfile.IdentitySystem, AzureADIdentitySystem)
		p.CustomCloudProfile.DependenciesLocation = DependenciesLocation(helpers.EnsureString(string(p.CustomCloudProfile.DependenciesLocation), AzureCustomCloudDependenciesLocationPublic))
		err := cs.SetCustomCloudProfileEnvironment()
		if err != nil {
			return fmt.Errorf("Failed to set environment - %s", err)
		}
		err = p.SetCustomCloudSpec(AzureCustomCloudSpecParams(params))
		if err != nil {
			return fmt.Errorf("Failed to set cloud spec - %s", err)
		}
	}
	return nil
}

// SetCustomCloudProfileEnvironment retrieves the endpoints from metadata endpoint (when required) and sets the values for azure.Environment
func (cs *ContainerService) SetCustomCloudProfileEnvironment() error {
	p := cs.Properties
	if p.IsCustomCloudProfile() {
		if p.CustomCloudProfile.Environment == nil {
			p.CustomCloudProfile.Environment = &azure.Environment{}
		}

		env := p.CustomCloudProfile.Environment
		if env.Name == "" || env.ServiceManagementEndpoint == "" || env.ActiveDirectoryEndpoint == "" || env.GraphEndpoint == "" || env.ResourceManagerVMDNSSuffix == "" {
			if env.Name == "" {
				env.Name = AzureStackCloud
			}

			if p.IsAzureStackCloud() {
				if !strings.HasPrefix(p.CustomCloudProfile.PortalURL, fmt.Sprintf("https://portal.%s.", cs.Location)) {
					return fmt.Errorf("portalURL needs to start with https://portal.%s. ", cs.Location)
				}

				azsFQDNSuffix := getAzureStackFQDNSuffix(p.CustomCloudProfile.PortalURL, cs.Location)
				env.ResourceManagerEndpoint = fmt.Sprintf("https://management.%s.%s/", cs.Location, azsFQDNSuffix)
			} else if env.ResourceManagerEndpoint == "" {
				return fmt.Errorf("Non-AzureStack CustomCloudProfile MUST provide ResourceManagerEndpoint")
			}

			metadataURL := fmt.Sprintf("%s/metadata/endpoints?api-version=1.0", strings.TrimSuffix(env.ResourceManagerEndpoint, "/"))

			// Retrieve the metadata
			httpClient := &http.Client{
				Timeout: 30 * time.Second,
			}
			endpointsresp, err := httpClient.Get(metadataURL)
			if err != nil || endpointsresp.StatusCode != 200 {
				return fmt.Errorf("%s . apimodel invalid: failed to retrieve custom endpoints from metadataURL %s", err, metadataURL)
			}

			body, err := ioutil.ReadAll(endpointsresp.Body)
			if err != nil {
				return fmt.Errorf("%s . apimodel invalid: failed to read the response from metadataURL %s", err, metadataURL)
			}

			endpoints := AzureStackMetadataEndpoints{}
			err = json.Unmarshal(body, &endpoints)
			if err != nil {
				return fmt.Errorf("%s . apimodel invalid: failed to parse the response from metadataURL %s", err, metadataURL)
			}

			if endpoints.GraphEndpoint == "" || endpoints.Authentication == nil || endpoints.Authentication.LoginEndpoint == "" || len(endpoints.Authentication.Audiences) == 0 || endpoints.Authentication.Audiences[0] == "" {
				return fmt.Errorf("%s . apimodel invalid: invalid response from %s", err, metadataURL)
			}

			env.GraphEndpoint = endpoints.GraphEndpoint
			env.ServiceManagementEndpoint = endpoints.Authentication.Audiences[0]
			env.GalleryEndpoint = endpoints.GalleryEndpoint
			env.ActiveDirectoryEndpoint = endpoints.Authentication.LoginEndpoint
			if p.CustomCloudProfile.IdentitySystem == ADFSIdentitySystem {
				env.ActiveDirectoryEndpoint = strings.TrimSuffix(env.ActiveDirectoryEndpoint, "/")
				env.ActiveDirectoryEndpoint = strings.TrimSuffix(env.ActiveDirectoryEndpoint, "adfs")
			}

			env.ManagementPortalURL = endpoints.PortalEndpoint

			if p.IsAzureStackCloud() {
				azsFQDNSuffix := getAzureStackFQDNSuffix(p.CustomCloudProfile.PortalURL, cs.Location)
				env.ResourceManagerVMDNSSuffix = fmt.Sprintf("cloudapp.%s", azsFQDNSuffix)
				env.StorageEndpointSuffix = fmt.Sprintf("%s.%s", cs.Location, azsFQDNSuffix)
				env.KeyVaultDNSSuffix = fmt.Sprintf("vault.%s.%s", cs.Location, azsFQDNSuffix)
			} else if env.ResourceManagerVMDNSSuffix == "" || env.StorageEndpointSuffix == "" || env.KeyVaultDNSSuffix == "" {
				// Non-AzureStack CustomCloud MUST provide suffixes
				return fmt.Errorf("Non-AzureStack CustomCloudProfile MUST provide ResourceManagerVMDNSSuffix, StorageEndpointSuffix, KeyVaultDNSSuffix")
			}
		}
	}

	return nil
}

// AzureCustomCloudSpecParams is the parameters when we set the custom cloud spec defaults for ContainerService.
type AzureCustomCloudSpecParams struct {
	IsUpgrade bool
	IsScale   bool
}

// SetAzureCustomCloudSpec sets the cloud spec for Custom Cloud .
func (p *Properties) SetCustomCloudSpec(params AzureCustomCloudSpecParams) error {
	if p.IsCustomCloudProfile() {
		var azureCustomCloudSpec AzureEnvironmentSpecConfig
		switch p.CustomCloudProfile.DependenciesLocation {
		case AzureCustomCloudDependenciesLocationPublic:
			azureCustomCloudSpec = AzureCloudSpecEnvMap[AzurePublicCloud]
		case AzureCustomCloudDependenciesLocationChina:
			azureCustomCloudSpec = AzureCloudSpecEnvMap[AzureChinaCloud]
		case AzureCustomCloudDependenciesLocationGerman:
			azureCustomCloudSpec = AzureCloudSpecEnvMap[AzureGermanCloud]
		case AzureCustomCloudDependenciesLocationUSGovernment:
			azureCustomCloudSpec = AzureCloudSpecEnvMap[AzureUSGovernmentCloud]
		default:
			azureCustomCloudSpec = AzureCloudSpecEnvMap[AzurePublicCloud]
		}
		if p.CustomCloudProfile.Environment == nil || p.CustomCloudProfile.Environment.ResourceManagerVMDNSSuffix == "" {
			return errors.New("Failed to set Cloud Spec for Azure Stack due to invalid environment")
		}

		azureCustomCloudSpec.EndpointConfig.ResourceManagerVMDNSSuffix = p.CustomCloudProfile.Environment.ResourceManagerVMDNSSuffix

		if p.CustomCloudProfile.Environment.Name == "" || p.IsAzureStackCloud() {
			azureCustomCloudSpec.CloudName = AzureStackCloud
		} else {
			azureCustomCloudSpec.CloudName = p.CustomCloudProfile.Environment.Name
		}

		//Sets default values for telemetry PID where none is set
		if p.CustomCloudProfile.AzureEnvironmentSpecConfig == nil {
			switch {
			case params.IsScale:
				azureCustomCloudSpec.KubernetesSpecConfig.AzureTelemetryPID = DefaultAzureStackScaleTelemetryPID
			case params.IsUpgrade:
				azureCustomCloudSpec.KubernetesSpecConfig.AzureTelemetryPID = DefaultAzureStackUpgradeTelemetryPID
			default:
				azureCustomCloudSpec.KubernetesSpecConfig.AzureTelemetryPID = DefaultAzureStackDeployTelemetryPID
			}

		}

		// Use the custom input to overwrite the default values in azureCustomCloudSpec
		if p.CustomCloudProfile.AzureEnvironmentSpecConfig != nil {
			ascc := p.CustomCloudProfile.AzureEnvironmentSpecConfig
			azureCustomCloudSpec.CloudName = helpers.EnsureString(ascc.CloudName, azureCustomCloudSpec.CloudName)

			// DockerSpecConfig
			asccDockerSpecConfig := ascc.DockerSpecConfig
			azsDockerSpecConfig := azureCustomCloudSpec.DockerSpecConfig
			azureCustomCloudSpec.DockerSpecConfig.DockerComposeDownloadURL = helpers.EnsureString(asccDockerSpecConfig.DockerComposeDownloadURL, azsDockerSpecConfig.DockerComposeDownloadURL)
			azureCustomCloudSpec.DockerSpecConfig.DockerEngineRepo = helpers.EnsureString(asccDockerSpecConfig.DockerEngineRepo, azsDockerSpecConfig.DockerComposeDownloadURL)

			//KubernetesSpecConfig
			asccKubernetesSpecConfig := ascc.KubernetesSpecConfig
			azsKubernetesSpecConfig := azureCustomCloudSpec.KubernetesSpecConfig

			azureCustomCloudSpec.KubernetesSpecConfig.AzureTelemetryPID = helpers.EnsureString(asccKubernetesSpecConfig.AzureTelemetryPID, DefaultAzureStackDeployTelemetryPID)
			azureCustomCloudSpec.KubernetesSpecConfig.ACIConnectorImageBase = helpers.EnsureString(asccKubernetesSpecConfig.ACIConnectorImageBase, azsKubernetesSpecConfig.ACIConnectorImageBase)
			azureCustomCloudSpec.KubernetesSpecConfig.AzureCNIImageBase = helpers.EnsureString(asccKubernetesSpecConfig.AzureCNIImageBase, azsKubernetesSpecConfig.AzureCNIImageBase)
			azureCustomCloudSpec.KubernetesSpecConfig.CalicoImageBase = helpers.EnsureString(asccKubernetesSpecConfig.CalicoImageBase, azsKubernetesSpecConfig.CalicoImageBase)
			azureCustomCloudSpec.KubernetesSpecConfig.CNIPluginsDownloadURL = helpers.EnsureString(asccKubernetesSpecConfig.CNIPluginsDownloadURL, azsKubernetesSpecConfig.CNIPluginsDownloadURL)
			azureCustomCloudSpec.KubernetesSpecConfig.ContainerdDownloadURLBase = helpers.EnsureString(asccKubernetesSpecConfig.ContainerdDownloadURLBase, azsKubernetesSpecConfig.ContainerdDownloadURLBase)
			azureCustomCloudSpec.KubernetesSpecConfig.CSIProxyDownloadURL = helpers.EnsureString(asccKubernetesSpecConfig.CSIProxyDownloadURL, azsKubernetesSpecConfig.CSIProxyDownloadURL)
			azureCustomCloudSpec.KubernetesSpecConfig.EtcdDownloadURLBase = helpers.EnsureString(asccKubernetesSpecConfig.EtcdDownloadURLBase, azsKubernetesSpecConfig.EtcdDownloadURLBase)
			azureCustomCloudSpec.KubernetesSpecConfig.KubeBinariesSASURLBase = helpers.EnsureString(asccKubernetesSpecConfig.KubeBinariesSASURLBase, azsKubernetesSpecConfig.KubeBinariesSASURLBase)
			azureCustomCloudSpec.KubernetesSpecConfig.KubernetesImageBase = helpers.EnsureString(asccKubernetesSpecConfig.KubernetesImageBase, azsKubernetesSpecConfig.KubernetesImageBase)
			azureCustomCloudSpec.KubernetesSpecConfig.MCRKubernetesImageBase = helpers.EnsureString(asccKubernetesSpecConfig.MCRKubernetesImageBase, azsKubernetesSpecConfig.MCRKubernetesImageBase)
			azureCustomCloudSpec.KubernetesSpecConfig.NVIDIAImageBase = helpers.EnsureString(asccKubernetesSpecConfig.NVIDIAImageBase, azsKubernetesSpecConfig.NVIDIAImageBase)
			azureCustomCloudSpec.KubernetesSpecConfig.TillerImageBase = helpers.EnsureString(asccKubernetesSpecConfig.TillerImageBase, azsKubernetesSpecConfig.TillerImageBase)
			azureCustomCloudSpec.KubernetesSpecConfig.VnetCNILinuxPluginsDownloadURL = helpers.EnsureString(asccKubernetesSpecConfig.VnetCNILinuxPluginsDownloadURL, azsKubernetesSpecConfig.VnetCNILinuxPluginsDownloadURL)
			azureCustomCloudSpec.KubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL = helpers.EnsureString(asccKubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL, azsKubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL)
			azureCustomCloudSpec.KubernetesSpecConfig.WindowsTelemetryGUID = helpers.EnsureString(asccKubernetesSpecConfig.WindowsTelemetryGUID, azsKubernetesSpecConfig.WindowsTelemetryGUID)

			//EndpointConfig
			asccEndpointConfig := ascc.EndpointConfig
			azsEndpointConfig := azureCustomCloudSpec.EndpointConfig
			azureCustomCloudSpec.EndpointConfig.ResourceManagerVMDNSSuffix = helpers.EnsureString(asccEndpointConfig.ResourceManagerVMDNSSuffix, azsEndpointConfig.ResourceManagerVMDNSSuffix)

			//OSImageConfig
			azureCustomCloudSpec.OSImageConfig = make(map[Distro]AzureOSImageConfig)
			for k, v := range ascc.OSImageConfig {
				azureCustomCloudSpec.OSImageConfig[k] = v
			}
			p.CustomCloudProfile.AzureEnvironmentSpecConfig = &azureCustomCloudSpec
		}

		// Kubernetes only understand AzureStackCloud environment (AzureCloudSpecEnvMap is only accessed using AzureStackCloud for custom clouds)
		AzureCloudSpecEnvMap[AzureStackCloud] = azureCustomCloudSpec
	}
	return nil
}
