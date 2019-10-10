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

func (cs *ContainerService) setCustomCloudProfileDefaults(params CustomCloudProfileDefaultsParams) error {
	p := cs.Properties
	if p.IsAzureStackCloud() {
		p.CustomCloudProfile.AuthenticationMethod = helpers.EnsureString(p.CustomCloudProfile.AuthenticationMethod, ClientSecretAuthMethod)
		p.CustomCloudProfile.IdentitySystem = helpers.EnsureString(p.CustomCloudProfile.IdentitySystem, AzureADIdentitySystem)
		p.CustomCloudProfile.DependenciesLocation = DependenciesLocation(helpers.EnsureString(string(p.CustomCloudProfile.DependenciesLocation), AzureStackDependenciesLocationPublic))
		err := cs.SetCustomCloudProfileEnvironment()
		if err != nil {
			return fmt.Errorf("Failed to set environment - %s", err)
		}
		err = p.SetAzureStackCloudSpec(AzureStackCloudSpecParams(params))
		if err != nil {
			return fmt.Errorf("Failed to set cloud spec - %s", err)
		}
	}
	return nil
}

// SetCustomCloudProfileEnvironment retrieves the endpoints from Azure Stack metadata endpoint and sets the values for azure.Environment
func (cs *ContainerService) SetCustomCloudProfileEnvironment() error {
	p := cs.Properties
	if p.IsAzureStackCloud() {
		if p.CustomCloudProfile.Environment == nil {
			p.CustomCloudProfile.Environment = &azure.Environment{}
		}

		env := p.CustomCloudProfile.Environment
		if env.Name == "" || env.ResourceManagerEndpoint == "" || env.ServiceManagementEndpoint == "" || env.ActiveDirectoryEndpoint == "" || env.GraphEndpoint == "" || env.ResourceManagerVMDNSSuffix == "" {
			env.Name = AzureStackCloud
			if !strings.HasPrefix(p.CustomCloudProfile.PortalURL, fmt.Sprintf("https://portal.%s.", cs.Location)) {
				return fmt.Errorf("portalURL needs to start with https://portal.%s. ", cs.Location)
			}
			azsFQDNSuffix := strings.Replace(p.CustomCloudProfile.PortalURL, fmt.Sprintf("https://portal.%s.", cs.Location), "", -1)
			azsFQDNSuffix = strings.TrimSuffix(azsFQDNSuffix, "/")
			env.ResourceManagerEndpoint = fmt.Sprintf("https://management.%s.%s/", cs.Location, azsFQDNSuffix)
			metadataURL := fmt.Sprintf("%s/metadata/endpoints?api-version=1.0", strings.TrimSuffix(env.ResourceManagerEndpoint, "/"))

			// Retrieve the metadata
			httpClient := &http.Client{
				Timeout: 30 * time.Second,
			}
			endpointsresp, err := httpClient.Get(metadataURL)
			if err != nil || endpointsresp.StatusCode != 200 {
				return fmt.Errorf("%s . apimodel invalid: failed to retrieve Azure Stack endpoints from %s", err, metadataURL)
			}

			body, err := ioutil.ReadAll(endpointsresp.Body)
			if err != nil {
				return fmt.Errorf("%s . apimodel invalid: failed to read the response from %s", err, metadataURL)
			}

			endpoints := AzureStackMetadataEndpoints{}
			err = json.Unmarshal(body, &endpoints)
			if err != nil {
				return fmt.Errorf("%s . apimodel invalid: failed to parse the response from %s", err, metadataURL)
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
			env.ResourceManagerVMDNSSuffix = fmt.Sprintf("cloudapp.%s", azsFQDNSuffix)
			env.StorageEndpointSuffix = fmt.Sprintf("%s.%s", cs.Location, azsFQDNSuffix)
			env.KeyVaultDNSSuffix = fmt.Sprintf("vault.%s.%s", cs.Location, azsFQDNSuffix)
		}
	}
	return nil
}

// AzureStackCloudSpecParams is the parameters when we set the azure stack cloud spec defaults for ContainerService.
type AzureStackCloudSpecParams struct {
	IsUpgrade bool
	IsScale   bool
}

// SetAzureStackCloudSpec sets the cloud spec for Azure Stack .
func (p *Properties) SetAzureStackCloudSpec(params AzureStackCloudSpecParams) error {
	if p.IsAzureStackCloud() {
		var azureStackCloudSpec AzureEnvironmentSpecConfig
		switch p.CustomCloudProfile.DependenciesLocation {
		case AzureStackDependenciesLocationPublic:
			azureStackCloudSpec = AzureCloudSpecEnvMap[AzurePublicCloud]
		case AzureStackDependenciesLocationChina:
			azureStackCloudSpec = AzureCloudSpecEnvMap[AzureChinaCloud]
		case AzureStackDependenciesLocationGerman:
			azureStackCloudSpec = AzureCloudSpecEnvMap[AzureGermanCloud]
		case AzureStackDependenciesLocationUSGovernment:
			azureStackCloudSpec = AzureCloudSpecEnvMap[AzureUSGovernmentCloud]
		default:
			azureStackCloudSpec = AzureCloudSpecEnvMap[AzurePublicCloud]
		}
		if p.CustomCloudProfile.Environment == nil || p.CustomCloudProfile.Environment.ResourceManagerVMDNSSuffix == "" {
			return errors.New("Failed to set Cloud Spec for Azure Stack due to invalid environment")
		}

		azureStackCloudSpec.EndpointConfig.ResourceManagerVMDNSSuffix = p.CustomCloudProfile.Environment.ResourceManagerVMDNSSuffix
		azureStackCloudSpec.CloudName = AzureStackCloud

		//Sets default values for telemetry PID where none is set
		if p.CustomCloudProfile.AzureEnvironmentSpecConfig == nil {
			switch {
			case params.IsScale:
				azureStackCloudSpec.KubernetesSpecConfig.AzureTelemetryPID = DefaultAzureStackScaleTelemetryPID
			case params.IsUpgrade:
				azureStackCloudSpec.KubernetesSpecConfig.AzureTelemetryPID = DefaultAzureStackUpgradeTelemetryPID
			default:
				azureStackCloudSpec.KubernetesSpecConfig.AzureTelemetryPID = DefaultAzureStackDeployTelemetryPID
			}

		}

		// Use the custom input to overwrite the default values in AzureStackCloudSpec
		if p.CustomCloudProfile.AzureEnvironmentSpecConfig != nil {
			ascc := p.CustomCloudProfile.AzureEnvironmentSpecConfig
			azureStackCloudSpec.CloudName = helpers.EnsureString(ascc.CloudName, azureStackCloudSpec.CloudName)

			// DockerSpecConfig
			asccDockerSpecConfig := ascc.DockerSpecConfig
			azsDockerSpecConfig := azureStackCloudSpec.DockerSpecConfig
			azureStackCloudSpec.DockerSpecConfig.DockerComposeDownloadURL = helpers.EnsureString(asccDockerSpecConfig.DockerComposeDownloadURL, azsDockerSpecConfig.DockerComposeDownloadURL)
			azureStackCloudSpec.DockerSpecConfig.DockerEngineRepo = helpers.EnsureString(asccDockerSpecConfig.DockerEngineRepo, azsDockerSpecConfig.DockerComposeDownloadURL)

			//KubernetesSpecConfig
			asccKubernetesSpecConfig := ascc.KubernetesSpecConfig
			azsKubernetesSpecConfig := azureStackCloudSpec.KubernetesSpecConfig

			azureStackCloudSpec.KubernetesSpecConfig.AzureTelemetryPID = helpers.EnsureString(asccKubernetesSpecConfig.AzureTelemetryPID, DefaultAzureStackDeployTelemetryPID)
			azureStackCloudSpec.KubernetesSpecConfig.ACIConnectorImageBase = helpers.EnsureString(asccKubernetesSpecConfig.ACIConnectorImageBase, azsKubernetesSpecConfig.ACIConnectorImageBase)
			azureStackCloudSpec.KubernetesSpecConfig.AzureCNIImageBase = helpers.EnsureString(asccKubernetesSpecConfig.AzureCNIImageBase, azsKubernetesSpecConfig.AzureCNIImageBase)
			azureStackCloudSpec.KubernetesSpecConfig.CalicoImageBase = helpers.EnsureString(asccKubernetesSpecConfig.CalicoImageBase, azsKubernetesSpecConfig.CalicoImageBase)
			azureStackCloudSpec.KubernetesSpecConfig.CNIPluginsDownloadURL = helpers.EnsureString(asccKubernetesSpecConfig.CNIPluginsDownloadURL, azsKubernetesSpecConfig.CNIPluginsDownloadURL)
			azureStackCloudSpec.KubernetesSpecConfig.ContainerdDownloadURLBase = helpers.EnsureString(asccKubernetesSpecConfig.ContainerdDownloadURLBase, azsKubernetesSpecConfig.ContainerdDownloadURLBase)
			azureStackCloudSpec.KubernetesSpecConfig.EtcdDownloadURLBase = helpers.EnsureString(asccKubernetesSpecConfig.EtcdDownloadURLBase, azsKubernetesSpecConfig.EtcdDownloadURLBase)
			azureStackCloudSpec.KubernetesSpecConfig.KubeBinariesSASURLBase = helpers.EnsureString(asccKubernetesSpecConfig.KubeBinariesSASURLBase, azsKubernetesSpecConfig.KubeBinariesSASURLBase)
			azureStackCloudSpec.KubernetesSpecConfig.KubernetesImageBase = helpers.EnsureString(asccKubernetesSpecConfig.KubernetesImageBase, azsKubernetesSpecConfig.KubernetesImageBase)
			azureStackCloudSpec.KubernetesSpecConfig.MCRKubernetesImageBase = helpers.EnsureString(asccKubernetesSpecConfig.MCRKubernetesImageBase, azsKubernetesSpecConfig.MCRKubernetesImageBase)
			azureStackCloudSpec.KubernetesSpecConfig.NVIDIAImageBase = helpers.EnsureString(asccKubernetesSpecConfig.NVIDIAImageBase, azsKubernetesSpecConfig.NVIDIAImageBase)
			azureStackCloudSpec.KubernetesSpecConfig.TillerImageBase = helpers.EnsureString(asccKubernetesSpecConfig.TillerImageBase, azsKubernetesSpecConfig.TillerImageBase)
			azureStackCloudSpec.KubernetesSpecConfig.VnetCNILinuxPluginsDownloadURL = helpers.EnsureString(asccKubernetesSpecConfig.VnetCNILinuxPluginsDownloadURL, azsKubernetesSpecConfig.VnetCNILinuxPluginsDownloadURL)
			azureStackCloudSpec.KubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL = helpers.EnsureString(asccKubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL, azsKubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL)
			azureStackCloudSpec.KubernetesSpecConfig.WindowsTelemetryGUID = helpers.EnsureString(asccKubernetesSpecConfig.WindowsTelemetryGUID, azsKubernetesSpecConfig.WindowsTelemetryGUID)

			//EndpointConfig
			asccEndpointConfig := ascc.EndpointConfig
			azsEndpointConfig := azureStackCloudSpec.EndpointConfig
			azureStackCloudSpec.EndpointConfig.ResourceManagerVMDNSSuffix = helpers.EnsureString(asccEndpointConfig.ResourceManagerVMDNSSuffix, azsEndpointConfig.ResourceManagerVMDNSSuffix)

			//OSImageConfig
			azureStackCloudSpec.OSImageConfig = make(map[Distro]AzureOSImageConfig)
			for k, v := range ascc.OSImageConfig {
				azureStackCloudSpec.OSImageConfig[k] = v
			}
			p.CustomCloudProfile.AzureEnvironmentSpecConfig = &azureStackCloudSpec
		}
		AzureCloudSpecEnvMap[AzureStackCloud] = azureStackCloudSpec
	}
	return nil
}
