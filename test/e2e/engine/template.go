//+build test
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Azure/go-autorest/autorest/to"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/api/vlabs"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/Azure/aks-engine/test/e2e/config"
)

// Config represents the configuration values of a template stored as env vars
type Config struct {
	ClientID                       string `envconfig:"CLIENT_ID" required:"true"`
	ClientSecret                   string `envconfig:"CLIENT_SECRET" required:"true"`
	ClientObjectID                 string `envconfig:"CLIENT_OBJECTID" default:""`
	LogAnalyticsWorkspaceKey       string `envconfig:"LOG_ANALYTICS_WORKSPACE_KEY" default:""`
	MasterDNSPrefix                string `envconfig:"DNS_PREFIX" default:""`
	AgentDNSPrefix                 string `envconfig:"DNS_PREFIX" default:""`
	MSIUserAssignedID              string `envconfig:"MSI_USER_ASSIGNED_ID" default:""`
	PublicSSHKey                   string `envconfig:"PUBLIC_SSH_KEY" default:""`
	WindowsAdminPasssword          string `envconfig:"WINDOWS_ADMIN_PASSWORD" default:""`
	WindowsNodeImageGallery        string `envconfig:"WINDOWS_NODE_IMAGE_GALLERY" default:""`
	WindowsNodeImageName           string `envconfig:"WINDOWS_NODE_IMAGE_NAME" default:""`
	WindowsNodeImageResourceGroup  string `envconfig:"WINDOWS_NODE_IMAGE_RESOURCE_GROUP" default:""`
	WindowsNodeImageSubscriptionID string `envconfig:"WINDOWS_NODE_IMAGE_SUBSCRIPTION_ID" default:""`
	WindowsNodeImageVersion        string `envconfig:"WINDOWS_NODE_IMAGE_VERSION" default:""`
	WindowsNodeVhdURL              string `envconfig:"WINDOWS_NODE_VHD_URL" default:""`
	LinuxNodeImageGallery          string `envconfig:"LINUX_NODE_IMAGE_GALLERY" default:""`
	LinuxNodeImageName             string `envconfig:"LINUX_NODE_IMAGE_NAME" default:""`
	LinuxNodeImageResourceGroup    string `envconfig:"LINUX_NODE_IMAGE_RESOURCE_GROUP" default:""`
	LinuxNodeImageSubscriptionID   string `envconfig:"LINUX_NODE_IMAGE_SUBSCRIPTION_ID" default:""`
	LinuxNodeImageVersion          string `envconfig:"LINUX_NODE_IMAGE_VERSION" default:""`
	OSDiskSizeGB                   string `envconfig:"OS_DISK_SIZE_GB" default:""`
	ContainerRuntime               string `envconfig:"CONTAINER_RUNTIME" default:""`
	OrchestratorRelease            string `envconfig:"ORCHESTRATOR_RELEASE" default:""`
	OrchestratorVersion            string `envconfig:"ORCHESTRATOR_VERSION" default:""`
	OutputDirectory                string `envconfig:"OUTPUT_DIR" default:"_output"`
	CreateVNET                     bool   `envconfig:"CREATE_VNET" default:"false"`
	EnableKMSEncryption            bool   `envconfig:"ENABLE_KMS_ENCRYPTION" default:"false"`
	Distro                         string `envconfig:"DISTRO" default:""`
	SubscriptionID                 string `envconfig:"SUBSCRIPTION_ID" required:"true"`
	InfraResourceGroup             string `envconfig:"INFRA_RESOURCE_GROUP" default:""`
	Location                       string `envconfig:"LOCATION" default:""`
	TenantID                       string `envconfig:"TENANT_ID" required:"true"`
	ImageName                      string `envconfig:"IMAGE_NAME" default:""`
	ImageResourceGroup             string `envconfig:"IMAGE_RESOURCE_GROUP" default:""`
	DebugCrashingPods              bool   `envconfig:"DEBUG_CRASHING_PODS" default:"false"`
	CustomHyperKubeImage           string `envconfig:"CUSTOM_HYPERKUBE_IMAGE" default:""`
	EnableTelemetry                bool   `envconfig:"ENABLE_TELEMETRY" default:"true"`
	KubernetesImageBase            string `envconfig:"KUBERNETES_IMAGE_BASE" default:""`
	KubernetesImageBaseType        string `envconfig:"KUBERNETES_IMAGE_BASE_TYPE" default:""`

	ClusterDefinitionPath     string // The original template we want to use to build the cluster from.
	ClusterDefinitionTemplate string // This is the template after we splice in the environment variables
	GeneratedDefinitionPath   string // Holds the contents of running aks-engine generate
	OutputPath                string // This is the root output path
	DefinitionName            string // Unique cluster name
	GeneratedTemplatePath     string // azuredeploy.json path
	GeneratedParametersPath   string // azuredeploy.parameters.json path
}

// Engine holds necessary information to interact with aks-engine cli
type Engine struct {
	Config             *Config
	ClusterDefinition  *api.VlabsARMContainerService // Holds the parsed ClusterDefinition
	ExpandedDefinition *api.ContainerService         // Holds the expanded ClusterDefinition
}

// ParseConfig will return a new engine config struct taking values from env vars
func ParseConfig(cwd, clusterDefinition, name string) (*Config, error) {
	c := new(Config)
	if err := envconfig.Process("config", c); err != nil {
		return nil, err
	}

	clusterDefinitionTemplate := fmt.Sprintf("%s/%s.json", c.OutputDirectory, name)
	generatedDefinitionPath := fmt.Sprintf("%s/%s", c.OutputDirectory, name)
	c.DefinitionName = name
	c.ClusterDefinitionPath = filepath.Join(cwd, clusterDefinition)
	c.ClusterDefinitionTemplate = filepath.Join(cwd, clusterDefinitionTemplate)
	c.OutputPath = filepath.Join(cwd, c.OutputDirectory)
	c.GeneratedDefinitionPath = filepath.Join(cwd, generatedDefinitionPath)
	c.GeneratedTemplatePath = filepath.Join(cwd, generatedDefinitionPath, "azuredeploy.json")
	c.GeneratedParametersPath = filepath.Join(cwd, generatedDefinitionPath, "azuredeploy.parameters.json")
	return c, nil
}

// Build takes a template path and will inject values based on provided environment variables
// it will then serialize the structs back into json and save it to outputPath
func Build(cfg *config.Config, masterSubnetID string, agentSubnetIDs []string, isVMSS bool) (*Engine, error) {
	config, err := ParseConfig(cfg.CurrentWorkingDir, cfg.ClusterDefinition, cfg.Name)
	if err != nil {
		log.Printf("Error while trying to build Engine Configuration:%s\n", err)
	}

	cs, err := ParseInput(config.ClusterDefinitionPath)
	if err != nil {
		return nil, err
	}
	if cs.Location == "" {
		cs.Location = config.Location
	}
	prop := cs.ContainerService.Properties
	var hasWindows bool
	if prop.HasWindows() {
		hasWindows = true
	}

	if config.ClientID != "" && config.ClientSecret != "" {
		if !prop.IsAzureStackCloud() {
			prop.ServicePrincipalProfile = &vlabs.ServicePrincipalProfile{
				ClientID: config.ClientID,
				Secret:   config.ClientSecret,
			}
		}
	}

	if config.MasterDNSPrefix != "" {
		prop.MasterProfile.DNSPrefix = config.MasterDNSPrefix
	}

	if !cfg.IsKubernetes() && config.AgentDNSPrefix != "" {
		for idx, pool := range prop.AgentPoolProfiles {
			pool.DNSPrefix = fmt.Sprintf("%v-%v", config.AgentDNSPrefix, idx)
		}
	}

	if prop.OrchestratorProfile.KubernetesConfig == nil {
		prop.OrchestratorProfile.KubernetesConfig = &vlabs.KubernetesConfig{}
	}

	if prop.LinuxProfile != nil {
		if config.PublicSSHKey != "" {
			prop.LinuxProfile.SSH.PublicKeys[0].KeyData = config.PublicSSHKey
			if prop.OrchestratorProfile.KubernetesConfig != nil && prop.OrchestratorProfile.KubernetesConfig.PrivateCluster != nil && prop.OrchestratorProfile.KubernetesConfig.PrivateCluster.JumpboxProfile != nil {
				prop.OrchestratorProfile.KubernetesConfig.PrivateCluster.JumpboxProfile.PublicKey = config.PublicSSHKey
			}
		}
	}

	if config.KubernetesImageBase != "" {
		prop.OrchestratorProfile.KubernetesConfig.KubernetesImageBase = config.KubernetesImageBase
	}
	if config.KubernetesImageBaseType != "" {
		prop.OrchestratorProfile.KubernetesConfig.KubernetesImageBaseType = config.KubernetesImageBaseType
	}

	if config.WindowsAdminPasssword != "" {
		prop.WindowsProfile.AdminPassword = config.WindowsAdminPasssword
	}

	if config.WindowsNodeVhdURL != "" {
		prop.WindowsProfile.WindowsImageSourceURL = config.WindowsNodeVhdURL
		log.Printf("Windows nodes will use image at %s for test pass", config.WindowsNodeVhdURL)
	} else if config.WindowsNodeImageName != "" && config.WindowsNodeImageResourceGroup != "" {
		prop.WindowsProfile.ImageRef = &vlabs.ImageReference{
			Name:          config.WindowsNodeImageName,
			ResourceGroup: config.WindowsNodeImageResourceGroup,
		}

		if config.WindowsNodeImageGallery != "" && config.WindowsNodeImageSubscriptionID != "" && config.WindowsNodeImageVersion != "" {
			prop.WindowsProfile.ImageRef.Gallery = config.WindowsNodeImageGallery
			prop.WindowsProfile.ImageRef.SubscriptionID = config.WindowsNodeImageSubscriptionID
			prop.WindowsProfile.ImageRef.Version = config.WindowsNodeImageVersion
		}
		log.Printf("Windows nodes will use image reference name:%s, rg:%s, sub:%s, gallery:%s, version:%s for test pass", config.WindowsNodeImageName, config.WindowsNodeImageResourceGroup, config.WindowsNodeImageSubscriptionID, config.WindowsNodeImageGallery, config.WindowsNodeImageVersion)
	}

	if config.LinuxNodeImageName != "" && config.LinuxNodeImageResourceGroup != "" {
		prop.MasterProfile.ImageRef = &vlabs.ImageReference{
			Name:          config.LinuxNodeImageName,
			ResourceGroup: config.LinuxNodeImageResourceGroup,
		}
		prop.MasterProfile.ImageRef.Gallery = config.LinuxNodeImageGallery
		prop.MasterProfile.ImageRef.SubscriptionID = config.LinuxNodeImageSubscriptionID
		prop.MasterProfile.ImageRef.Version = config.LinuxNodeImageVersion
		if len(prop.AgentPoolProfiles) == 1 {
			prop.AgentPoolProfiles[0].ImageRef = &vlabs.ImageReference{
				Name:          config.LinuxNodeImageName,
				ResourceGroup: config.LinuxNodeImageResourceGroup,
			}
			if config.LinuxNodeImageGallery != "" && config.LinuxNodeImageSubscriptionID != "" && config.LinuxNodeImageVersion != "" {
				prop.AgentPoolProfiles[0].ImageRef.Gallery = config.LinuxNodeImageGallery
				prop.AgentPoolProfiles[0].ImageRef.SubscriptionID = config.LinuxNodeImageSubscriptionID
				prop.AgentPoolProfiles[0].ImageRef.Version = config.LinuxNodeImageVersion
			}
		}
	}

	if config.OSDiskSizeGB != "" {
		if osDiskSizeGB, err := strconv.Atoi(config.OSDiskSizeGB); err == nil {
			prop.MasterProfile.OSDiskSizeGB = osDiskSizeGB
			for _, pool := range prop.AgentPoolProfiles {
				pool.OSDiskSizeGB = osDiskSizeGB
			}
		}
	}

	if config.ContainerRuntime == "containerd" &&
		prop.OrchestratorProfile.KubernetesConfig.WindowsContainerdURL == "" &&
		prop.OrchestratorProfile.KubernetesConfig.WindowsSdnPluginURL == "" {
		prop.OrchestratorProfile.KubernetesConfig.WindowsContainerdURL = "https://aksenginee2etestimages.blob.core.windows.net/test-content/windows-cri-containerd.zip"
		prop.OrchestratorProfile.KubernetesConfig.WindowsSdnPluginURL = "https://aksenginee2etestimages.blob.core.windows.net/test-content/windows-cni-containerd.zip"
	}

	if config.ContainerRuntime != "" {
		prop.OrchestratorProfile.KubernetesConfig.ContainerRuntime = config.ContainerRuntime
	}

	// If the parsed api model input has no expressed version opinion, we check if ENV does have an opinion
	if prop.OrchestratorProfile.OrchestratorRelease == "" &&
		prop.OrchestratorProfile.OrchestratorVersion == "" {
		// First, prefer the release string if ENV declares it
		if config.OrchestratorRelease != "" {
			prop.OrchestratorProfile.OrchestratorRelease = config.OrchestratorRelease
			// Or, choose the version string if ENV declares it
		} else if config.OrchestratorVersion != "" {
			prop.OrchestratorProfile.OrchestratorVersion = config.OrchestratorVersion
			// If ENV similarly has no version opinion, we will rely upon the aks-engine default
		} else {
			prop.OrchestratorProfile.OrchestratorVersion = common.GetDefaultKubernetesVersion(hasWindows)
		}
	}

	if config.CreateVNET {
		if isVMSS {
			prop.MasterProfile.VnetSubnetID = masterSubnetID
			prop.MasterProfile.AgentVnetSubnetID = agentSubnetIDs[0]
			for _, p := range prop.AgentPoolProfiles {
				p.VnetSubnetID = agentSubnetIDs[0]
			}
		} else {
			prop.MasterProfile.VnetSubnetID = masterSubnetID
			for i, p := range prop.AgentPoolProfiles {
				p.VnetSubnetID = agentSubnetIDs[i]
			}
		}
	}

	if config.EnableKMSEncryption && config.ClientObjectID != "" {
		if prop.OrchestratorProfile.KubernetesConfig == nil {
			prop.OrchestratorProfile.KubernetesConfig = &vlabs.KubernetesConfig{}
		}
		prop.OrchestratorProfile.KubernetesConfig.EnableEncryptionWithExternalKms = &config.EnableKMSEncryption
		prop.ServicePrincipalProfile.ObjectID = config.ClientObjectID
	}

	var version string
	if prop.OrchestratorProfile.OrchestratorRelease != "" {
		version = prop.OrchestratorProfile.OrchestratorRelease + ".0"
	} else if prop.OrchestratorProfile.OrchestratorVersion != "" {
		version = prop.OrchestratorProfile.OrchestratorVersion
	}
	if common.IsKubernetesVersionGe(version, "1.12.0") {
		if prop.OrchestratorProfile.KubernetesConfig == nil {
			prop.OrchestratorProfile.KubernetesConfig = &vlabs.KubernetesConfig{}
		}
		prop.OrchestratorProfile.KubernetesConfig.ControllerManagerConfig = map[string]string{
			"--horizontal-pod-autoscaler-downscale-stabilization":   "30s",
			"--horizontal-pod-autoscaler-cpu-initialization-period": "30s",
		}
	}

	if config.LogAnalyticsWorkspaceKey != "" && len(prop.OrchestratorProfile.KubernetesConfig.Addons) > 0 {
		for _, addOn := range prop.OrchestratorProfile.KubernetesConfig.Addons {
			if addOn.Name == "container-monitoring" {
				if addOn.Config == nil {
					addOn.Config = make(map[string]string)
				}
				addOn.Config["workspaceKey"] = config.LogAnalyticsWorkspaceKey
				break
			}
		}
	}

	if config.CustomHyperKubeImage != "" {
		prop.OrchestratorProfile.KubernetesConfig.CustomHyperkubeImage = config.CustomHyperKubeImage
	}

	if config.EnableTelemetry == true {
		if prop.FeatureFlags == nil {
			prop.FeatureFlags = new(vlabs.FeatureFlags)
		}
		prop.FeatureFlags.EnableTelemetry = true
	}

	for _, pool := range prop.AgentPoolProfiles {
		if pool.DiskEncryptionSetID != "" {
			str := strings.Replace(pool.DiskEncryptionSetID, "SUB_ID", config.SubscriptionID, 1)
			str = strings.Replace(str, "RESOURCE_GROUP", config.InfraResourceGroup, 1)
			pool.DiskEncryptionSetID = str
		}
		if pool.ProximityPlacementGroupID != "" {
			str := strings.Replace(pool.ProximityPlacementGroupID, "SUB_ID", config.SubscriptionID, 1)
			str = strings.Replace(str, "RESOURCE_GROUP", config.InfraResourceGroup, 1)
			pool.ProximityPlacementGroupID = str
		}
	}

	if config.Distro != "" {
		prop.MasterProfile.Distro = vlabs.Distro(config.Distro)
		for _, pool := range prop.AgentPoolProfiles {
			if !pool.IsWindows() {
				pool.Distro = vlabs.Distro(config.Distro)
			}
		}
	}

	if config.MSIUserAssignedID != "" {
		prop.OrchestratorProfile.KubernetesConfig.UseManagedIdentity = true
		prop.OrchestratorProfile.KubernetesConfig.UserAssignedID = config.MSIUserAssignedID
	}

	return &Engine{
		Config:            config,
		ClusterDefinition: cs,
	}, nil
}

// NodeCount returns the number of nodes that should be provisioned for a given cluster definition
func (e *Engine) NodeCount() int {
	expectedCount := e.ExpandedDefinition.Properties.MasterProfile.Count
	for _, pool := range e.ExpandedDefinition.Properties.AgentPoolProfiles {
		expectedCount += pool.Count
	}
	return expectedCount
}

// AnyAgentIsLinux will return true if there is at least 1 linux agent pool
func (e *Engine) AnyAgentIsLinux() bool {
	for _, ap := range e.ExpandedDefinition.Properties.AgentPoolProfiles {
		if ap.OSType == "" || ap.OSType == "Linux" {
			return true
		}
	}
	return false
}

// HasWindowsAgents will return true is there is at least 1 windows agent pool
func (e *Engine) HasWindowsAgents() bool {
	return e.ExpandedDefinition.Properties.HasWindows()
}

// WindowsTestImages holds the Windows container image names used in this test pass
type WindowsTestImages struct {
	IIS        string
	ServerCore string
}

// GetWindowsTestImages will return the right list of container images for the Windows version used
func (e *Engine) GetWindowsTestImages() (*WindowsTestImages, error) {
	if !e.HasWindowsAgents() {
		return nil, errors.New("Can't guess a Windows version without Windows nodes in the cluster")
	}

	windowsSku := e.ExpandedDefinition.Properties.WindowsProfile.GetWindowsSku()
	// tip: curl -L https://mcr.microsoft.com/v2/windows/servercore/tags/list
	//      curl -L https://mcr.microsoft.com/v2/windows/servercore/iis/tags/list
	switch {
	case strings.Contains(windowsSku, "1903"):
		return &WindowsTestImages{IIS: "mcr.microsoft.com/windows/servercore/iis:windowsservercore-1903",
			ServerCore: "mcr.microsoft.com/windows/servercore:1903"}, nil
	case strings.Contains(windowsSku, "1809"), strings.Contains(windowsSku, "2019"):
		return &WindowsTestImages{IIS: "mcr.microsoft.com/windows/servercore/iis:20191112-windowsservercore-ltsc2019",
			ServerCore: "mcr.microsoft.com/windows/servercore:ltsc2019"}, nil
	case strings.Contains(windowsSku, "1803"):
		return &WindowsTestImages{IIS: "mcr.microsoft.com/windows/servercore/iis:windowsservercore-1803",
			ServerCore: "mcr.microsoft.com/windows/servercore:1803"}, nil
	case strings.Contains(windowsSku, "1709"):
		return nil, errors.New("Windows Server version 1709 is out of support")
	}
	return nil, errors.New("Unknown Windows version. GetWindowsSku() = " + windowsSku)
}

// HasAddon will return true if an addon is enabled
func (e *Engine) HasAddon(name string) (bool, api.KubernetesAddon) {
	for _, addon := range e.ExpandedDefinition.Properties.OrchestratorProfile.KubernetesConfig.Addons {
		if addon.Name == name {
			return to.Bool(addon.Enabled), addon
		}
	}
	return false, api.KubernetesAddon{}
}

// HasNetworkPolicy will return true if the specified network policy is enabled
func (e *Engine) HasNetworkPolicy(name string) bool {
	return strings.Contains(e.ExpandedDefinition.Properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy, name)
}

// Write will write the cluster definition to disk
func (e *Engine) Write() error {
	json, err := helpers.JSONMarshal(e.ClusterDefinition, false)
	if err != nil {
		log.Printf("Error while trying to serialize Container Service object to json:%s\n%+v\n", err, e.ClusterDefinition)
		return err
	}
	err = ioutil.WriteFile(e.Config.ClusterDefinitionTemplate, json, 0777)
	if err != nil {
		log.Printf("Error while trying to write container service definition to file (%s):%s\n%s\n", e.Config.ClusterDefinitionTemplate, err, string(json))
	}

	return nil
}

// ParseInput takes a template path and will parse that into a api.VlabsARMContainerService
func ParseInput(path string) (*api.VlabsARMContainerService, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Error while trying to read cluster definition at (%s):%s\n", path, err)
		return nil, err
	}
	cs := api.VlabsARMContainerService{}
	if err = json.Unmarshal(contents, &cs); err != nil {
		log.Printf("Error while trying to unmarshal container service json:%s\n%s\n", err, string(contents))
		return nil, err
	}
	return &cs, nil
}

// ParseOutput takes the generated api model and will parse that into a api.ContainerService
func ParseOutput(path string, validate, isUpdate bool) (*api.ContainerService, error) {
	locale, err := i18n.LoadTranslations()
	if err != nil {
		return nil, errors.Errorf(fmt.Sprintf("error loading translation files: %s", err.Error()))
	}
	apiloader := &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: locale,
		},
	}
	containerService, _, err := apiloader.LoadContainerServiceFromFile(path, validate, isUpdate, nil)
	if err != nil {
		return nil, err
	}
	return containerService, nil
}
