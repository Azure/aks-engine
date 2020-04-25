//+build test
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/vlabs"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	"github.com/kelseyhightower/envconfig"
)

// Config holds global test configuration
type Config struct {
	SkipTest            bool          `envconfig:"SKIP_TEST" default:"false"`
	SkipLogsCollection  bool          `envconfig:"SKIP_LOGS_COLLECTION" default:"true"`
	Orchestrator        string        `envconfig:"ORCHESTRATOR" default:"kubernetes"`
	Name                string        `envconfig:"NAME" default:""`                                                       // Name allows you to set the name of a cluster already created
	Location            string        `envconfig:"LOCATION" default:""`                                                   // Location where you want to create the cluster
	Regions             []string      `envconfig:"REGIONS" default:""`                                                    // A whitelist of availableregions
	ClusterDefinition   string        `envconfig:"CLUSTER_DEFINITION" required:"true" default:"examples/kubernetes.json"` // ClusterDefinition is the path on disk to the json template these are normally located in examples/
	CleanUpOnExit       bool          `envconfig:"CLEANUP_ON_EXIT" default:"false"`                                       // if true the tests will clean up rgs when tests finish
	CleanUpIfFail       bool          `envconfig:"CLEANUP_IF_FAIL" default:"false"`
	RetainSSH           bool          `envconfig:"RETAIN_SSH" default:"true"`
	StabilityIterations int           `envconfig:"STABILITY_ITERATIONS" default:"3"`
	ClusterInitPodName  string        `envconfig:"CLUSTER_INIT_POD_NAME" default:""`
	ClusterInitJobName  string        `envconfig:"CLUSTER_INIT_JOB_NAME" default:""`
	Timeout             time.Duration `envconfig:"TIMEOUT" default:"20m"`
	LBTimeout           time.Duration `envconfig:"LB_TIMEOUT" default:"20m"`
	CurrentWorkingDir   string
	SoakClusterName     string `envconfig:"SOAK_CLUSTER_NAME" default:""`
	ForceDeploy         bool   `envconfig:"FORCE_DEPLOY" default:"false"`
	UseDeployCommand    bool   `envconfig:"USE_DEPLOY_COMMAND" default:"false"`
	GinkgoFocus         string `envconfig:"GINKGO_FOCUS" default:""`
	GinkgoSkip          string `envconfig:"GINKGO_SKIP" default:""`
	GinkgoFailFast      bool   `envconfig:"GINKGO_FAIL_FAST" default:"true"`
	DebugAfterSuite     bool   `envconfig:"DEBUG_AFTERSUITE" default:"false"`
	BlockSSHPort        bool   `envconfig:"BLOCK_SSH" default:"false"`
	AddNodePoolInput    string `envconfig:"ADD_NODE_POOL_INPUT" default:""`
	TestPVC             bool   `envconfig:"TEST_PVC" default:"false"`
}

// CustomCloudConfig holds configurations for custom clould
type CustomCloudConfig struct {
	ServiceManagementEndpoint    string `envconfig:"SERVICE_MANAGEMENT_ENDPOINT" default:""`
	ResourceManagerEndpoint      string `envconfig:"RESOURCE_MANAGER_ENDPOINT" default:""`
	ActiveDirectoryEndpoint      string `envconfig:"ACTIVE_DIRECTORY_ENDPOINT" default:""`
	GalleryEndpoint              string `envconfig:"GALLERY_ENDPOINT" default:""`
	StorageEndpointSuffix        string `envconfig:"STORAGE_ENDPOINT_SUFFIX" default:""`
	KeyVaultDNSSuffix            string `envconfig:"KEY_VAULT_DNS_SUFFIX" default:""`
	GraphEndpoint                string `envconfig:"GRAPH_ENDPOINT" default:""`
	ServiceManagementVMDNSSuffix string `envconfig:"SERVICE_MANAGEMENT_VM_DNS_SUFFIX" default:""`
	ResourceManagerVMDNSSuffix   string `envconfig:"RESOURCE_MANAGER_VM_DNS_SUFFIX" default:""`
	IdentitySystem               string `envconfig:"IDENTITY_SYSTEM" default:""`
	AuthenticationMethod         string `envconfig:"AUTHENTICATION_METHOD" default:""`
	VaultID                      string `envconfig:"VAULT_ID" default:""`
	SecretName                   string `envconfig:"SECRET_NAME" default:""`
	CustomCloudClientID          string `envconfig:"CUSTOM_CLOUD_CLIENT_ID" default:""`
	CustomCloudSecret            string `envconfig:"CUSTOM_CLOUD_SECRET" default:""`
	APIProfile                   string `envconfig:"API_PROFILE" default:""`
	PortalURL                    string `envconfig:"PORTAL_ENDPOINT" default:""`
	TimeoutCommands              bool
	CustomCloudName              string `envconfig:"CUSTOM_CLOUD_NAME"`
}

const (
	kubernetesOrchestrator = "kubernetes"
)

// ParseConfig will parse needed environment variables for running the tests
func ParseConfig() (*Config, error) {
	c := new(Config)
	if err := envconfig.Process("config", c); err != nil {
		return nil, err
	}
	if c.Location == "" {
		c.SetRandomRegion()
	}
	return c, nil
}

// ParseCustomCloudConfig will parse needed environment variables for running the tests
func ParseCustomCloudConfig() (*CustomCloudConfig, error) {
	ccc := new(CustomCloudConfig)
	if err := envconfig.Process("customcloudconfig", ccc); err != nil {
		return nil, err
	}
	return ccc, nil
}

// GetKubeConfig returns the absolute path to the kubeconfig for c.Location
func (c *Config) GetKubeConfig() string {
	var kubeconfigPath string

	if c.IsKubernetes() {
		file := fmt.Sprintf("kubeconfig.%s.json", c.Location)
		kubeconfigPath = filepath.Join(c.CurrentWorkingDir, "_output", c.Name, "kubeconfig", file)
	}
	return kubeconfigPath
}

// IsCustomCloudProfile returns true if the cloud is a custom cloud
func (c *Config) IsCustomCloudProfile() bool {
	// c.ClusterDefinition is only set for new deployments
	// Not for upgrade/scale operations
	return os.Getenv("CUSTOM_CLOUD_NAME") != ""
}

// UpdateCustomCloudClusterDefinition updates the cluster definition from environment variables
func (c *Config) UpdateCustomCloudClusterDefinition(ccc *CustomCloudConfig) error {
	clusterDefinitionFullPath := fmt.Sprintf("%s/%s", c.CurrentWorkingDir, c.ClusterDefinition)
	cs := parseVlabsContainerSerice(clusterDefinitionFullPath)

	cs.Location = c.Location
	cs.Properties.CustomCloudProfile.PortalURL = ccc.PortalURL
	cs.Properties.ServicePrincipalProfile.ClientID = ccc.CustomCloudClientID
	cs.Properties.ServicePrincipalProfile.Secret = ccc.CustomCloudSecret
	cs.Properties.CustomCloudProfile.AuthenticationMethod = ccc.AuthenticationMethod
	cs.Properties.CustomCloudProfile.IdentitySystem = ccc.IdentitySystem

	if ccc.AuthenticationMethod == "client_certificate" {
		cs.Properties.ServicePrincipalProfile.Secret = ""
		cs.Properties.ServicePrincipalProfile.KeyvaultSecretRef = &vlabs.KeyvaultSecretRef{
			VaultID:    ccc.VaultID,
			SecretName: ccc.SecretName,
		}
	}

	csBytes, err := json.Marshal(cs)
	if err != nil {
		return fmt.Errorf("Error fail to marshal containerService object %p", err)
	}
	err = ioutil.WriteFile(clusterDefinitionFullPath, csBytes, 644)
	if err != nil {
		return fmt.Errorf("Error fail to write file object %p", err)
	}
	return nil
}

func parseVlabsContainerSerice(clusterDefinitionFullPath string) api.VlabsARMContainerService {

	bytes, err := ioutil.ReadFile(clusterDefinitionFullPath)
	if err != nil {
		log.Fatalf("Error while trying to read cluster definition at (%s):%s\n", clusterDefinitionFullPath, err)
	}
	cs := api.VlabsARMContainerService{}
	err = json.Unmarshal(bytes, &cs)
	if err != nil {
		log.Fatalf("Fail to unmarshal file %q , err -  %q", clusterDefinitionFullPath, err)
	}
	return cs
}

// SetEnvironment will set the cloud context
func (ccc *CustomCloudConfig) SetEnvironment() error {
	var cmd *exec.Cmd
	var err error

	// Add to python cert store the self-signed root CA generated by Azure Stack's CI
	// as azure-cli complains otherwise
	azsSelfSignedCaPath := "/aks-engine/Certificates.pem"
	if _, err = os.Stat(azsSelfSignedCaPath); err == nil {
		// latest dev_image has an azure-cli version that requires python3
		cmd := exec.Command("/bin/bash", "-c",
			fmt.Sprintf(`VER=$(python3 -V | grep -o [0-9].[0-9]*. | grep -o [0-9].[0-9]*);
		CA=/usr/local/lib/python${VER}/dist-packages/certifi/cacert.pem;
		if [ -f ${CA} ]; then cat %s >> ${CA}; fi;`, azsSelfSignedCaPath))
		if out, err := cmd.CombinedOutput(); err != nil {
			log.Printf("output:%s\n", out)
			return err
		}
	}

	environmentName := fmt.Sprintf("AzureStack%v", time.Now().Unix())
	if ccc.TimeoutCommands {
		cmd = exec.Command("timeout", "60", "az", "cloud", "register",
			"-n", environmentName,
			"--endpoint-resource-manager", ccc.ResourceManagerEndpoint,
			"--suffix-storage-endpoint", ccc.StorageEndpointSuffix,
			"--suffix-keyvault-dns", ccc.KeyVaultDNSSuffix,
			"--endpoint-active-directory-resource-id", ccc.ServiceManagementEndpoint,
			"--endpoint-active-directory", ccc.ActiveDirectoryEndpoint,
			"--endpoint-active-directory-graph-resource-id", ccc.GraphEndpoint)
	} else {
		cmd = exec.Command("az", "cloud", "register",
			"-n", environmentName,
			"--endpoint-resource-manager", ccc.ResourceManagerEndpoint,
			"--suffix-storage-endpoint", ccc.StorageEndpointSuffix,
			"--suffix-keyvault-dns", ccc.KeyVaultDNSSuffix,
			"--endpoint-active-directory-resource-id", ccc.ServiceManagementEndpoint,
			"--endpoint-active-directory", ccc.ActiveDirectoryEndpoint,
			"--endpoint-active-directory-graph-resource-id", ccc.GraphEndpoint)
	}
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Printf("output:%s\n", out)
		return err
	}

	if ccc.TimeoutCommands {
		cmd = exec.Command("timeout", "60", "az", "cloud", "set", "-n", environmentName)
	} else {
		cmd = exec.Command("az", "cloud", "set", "-n", environmentName)
	}
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Printf("output:%s\n", out)
		return err
	}

	if ccc.TimeoutCommands {
		cmd = exec.Command("timeout", "60", "az", "cloud", "update", "--profile", ccc.APIProfile)
	} else {
		cmd = exec.Command("az", "cloud", "update", "--profile", ccc.APIProfile)
	}
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Printf("output:%s\n", out)
		return err
	}

	return nil
}

// SetKubeConfig will set the KUBECONIFG env var
func (c *Config) SetKubeConfig() {
	os.Setenv("KUBECONFIG", c.GetKubeConfig())
	log.Printf("\nKubeconfig:%s\n", c.GetKubeConfig())
}

// GetSSHKeyPath will return the absolute path to the ssh private key
func (c *Config) GetSSHKeyPath() string {
	if c.UseDeployCommand {
		return filepath.Join(c.CurrentWorkingDir, "_output", c.Name, "azureuser_rsa")
	}
	return filepath.Join(c.CurrentWorkingDir, "_output", c.Name+"-ssh")
}

// SetEnvVars will determine if we need to
func (c *Config) SetEnvVars() error {
	envFile := fmt.Sprintf("%s/%s.env", c.CurrentWorkingDir, c.ClusterDefinition)
	if _, err := os.Stat(envFile); err == nil {
		file, err := os.Open(envFile)
		if err != nil {
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			log.Printf("Setting the following:%s\n", line)
			env := strings.Split(line, "=")
			if len(env) > 0 {
				os.Setenv(env[0], env[1])
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

// ReadPublicSSHKey will read the contents of the public ssh key on disk into a string
func (c *Config) ReadPublicSSHKey() (string, error) {
	file := c.GetSSHKeyPath() + ".pub"
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("Error while trying to read public ssh key at (%s):%s\n", file, err)
		return "", err
	}
	return string(contents), nil
}

// SetSSHKeyPermissions will change the ssh file permission to 0600
func (c *Config) SetSSHKeyPermissions() error {
	privateKey := c.GetSSHKeyPath()
	cmd := exec.Command("chmod", "0600", privateKey)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying to change private ssh key permissions at %s: %s\n", privateKey, out)
		return err
	}
	publicKey := c.GetSSHKeyPath() + ".pub"
	cmd = exec.Command("chmod", "0600", publicKey)
	util.PrintCommand(cmd)
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying to change public ssh key permissions at %s: %s\n", publicKey, out)
		return err
	}
	return nil
}

// IsKubernetes will return true if the ORCHESTRATOR env var is set to kubernetes or not set at all
func (c *Config) IsKubernetes() bool {
	return c.Orchestrator == kubernetesOrchestrator
}

// SetRandomRegion sets Location to a random region
func (c *Config) SetRandomRegion() {
	var regions []string
	if c.Regions == nil || len(c.Regions) == 0 {
		regions = []string{"eastus", "uksouth", "southeastasia", "westus2", "westeurope"}
	} else {
		regions = c.Regions
	}
	log.Printf("Picking Random Region from list %s\n", regions)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	c.Location = regions[r.Intn(len(regions))]
	os.Setenv("LOCATION", c.Location)
	log.Printf("Picked Random Region:%s\n", c.Location)
}
