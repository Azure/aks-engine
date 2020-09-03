//+build test
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package runner

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/test/e2e/azure"
	"github.com/Azure/aks-engine/test/e2e/config"
	"github.com/Azure/aks-engine/test/e2e/engine"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/node"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	"github.com/Azure/aks-engine/test/e2e/metrics"
	"github.com/Azure/aks-engine/test/e2e/remote"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// CLIProvisioner holds the configuration needed to provision a clusters
type CLIProvisioner struct {
	ClusterDefinition string `envconfig:"CLUSTER_DEFINITION" required:"true" default:"examples/kubernetes.json"` // ClusterDefinition is the path on disk to the json template these are normally located in examples/
	ProvisionRetries  int    `envconfig:"PROVISION_RETRIES" default:"0"`
	CreateVNET        bool   `envconfig:"CREATE_VNET" default:"false"`
	MasterVMSS        bool   `envconfig:"MASTER_VMSS" default:"false"`
	Config            *config.Config
	Account           *azure.Account
	Point             *metrics.Point
	ResourceGroups    []string
	Engine            *engine.Engine
	Masters           []azure.VM
	Agents            []azure.VM
}

// BuildCLIProvisioner will return a ProvisionerConfig object which is used to run a provision
func BuildCLIProvisioner(cfg *config.Config, acct *azure.Account, pt *metrics.Point) (*CLIProvisioner, error) {
	p := new(CLIProvisioner)
	if err := envconfig.Process("provisioner", p); err != nil {
		return nil, err
	}
	p.Config = cfg
	p.Account = acct
	p.Point = pt
	return p, nil
}

// Run will provision a cluster using the azure cli
func (cli *CLIProvisioner) Run() error {
	rgs := make([]string, 0)
	for i := 0; i <= cli.ProvisionRetries; i++ {
		cli.Point.SetProvisionStart()
		err := cli.provision()
		rgs = append(rgs, cli.Config.Name)
		cli.ResourceGroups = rgs
		if err != nil {
			if i < cli.ProvisionRetries {
				cli.Point.RecordProvisionError()
			} else if i == cli.ProvisionRetries {
				cli.Point.RecordProvisionError()
				return errors.Errorf("Exceeded provision retry count: %s", err.Error())
			}
		} else {
			cli.Point.RecordProvisionSuccess()
			if cli.Config.SkipTest {
				return nil
			} else {
				cli.Point.SetNodeWaitStart()
				err := cli.waitForNodes()
				cli.Point.RecordNodeWait(err)
				return err
			}
		}
	}
	return errors.New("Unable to run provisioner")
}

func createSaveSSH(keyPath string, createPrivateKey bool) (string, error) {
	cmd := exec.Command("ssh-keygen", "-f", keyPath, "-q", "-N", "", "-b", "2048", "-t", "rsa")
	if !createPrivateKey {
		cmd = exec.Command("ssh-keygen", "-y", "-f", keyPath)
	}

	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.Wrapf(err, "Error while trying to generate ssh key\nOutput:%s", out)
	}
	if !createPrivateKey {
		err2 := ioutil.WriteFile(keyPath+".pub", out, 0644)
		if err2 != nil {
			return "", errors.Wrapf(err2, "Error while trying to write public ssh key")
		}
	}

	os.Chmod(keyPath, 0600)
	publicSSHKeyBytes, err := ioutil.ReadFile(keyPath + ".pub")
	if err != nil {
		return "", errors.Wrap(err, "Error while trying to read public ssh key")
	}
	return string(publicSSHKeyBytes), nil
}

func (cli *CLIProvisioner) provision() error {
	cli.Config.Name = cli.generateName()
	if cli.Config.SoakClusterName != "" {
		cli.Config.Name = cli.Config.SoakClusterName
	}
	os.Setenv("NAME", cli.Config.Name)

	outputPath := filepath.Join(cli.Config.CurrentWorkingDir, "_output")
	if !cli.Config.UseDeployCommand {
		privateKeyName := cli.Config.Name + "-ssh"
		os.Mkdir(outputPath, 0755)
		privateKeyPath := filepath.Join(outputPath, privateKeyName)
		createPrivateKey := true
		if cli.Config.PrivateSSHKeyPath != "" {
			err := os.Rename(cli.Config.PrivateSSHKeyPath, privateKeyPath)
			if err != nil {
				return errors.Wrapf(err, "Error while trying to move private ssh key")
			}
			createPrivateKey = false
		}
		publicSSHKey, err2 := createSaveSSH(privateKeyPath, createPrivateKey)
		if err2 != nil {
			return errors.Wrap(err2, "Error while generating ssh keys")
		}
		os.Setenv("PUBLIC_SSH_KEY", publicSSHKey)
	}

	os.Setenv("DNS_PREFIX", cli.Config.Name)

	err := cli.Account.CreateGroupWithRetry(cli.Config.Name, cli.Config.Location, 30*time.Second, cli.Config.Timeout)
	if err != nil {
		return errors.Wrap(err, "Error while trying to create resource group")
	}
	cli.Account.ResourceGroup = azure.ResourceGroup{
		Name:     cli.Config.Name,
		Location: cli.Config.Location,
		Tags: map[string]string{
			"now": fmt.Sprintf("now=%v", time.Now().Unix()),
		},
	}
	err = cli.Account.ShowGroupWithRetry(cli.Account.ResourceGroup.Name, 10*time.Second, cli.Config.Timeout)
	if err != nil {
		return errors.Wrap(err, "Unable to successfully get the resource group using the az CLI")
	}

	subnetID := ""
	vnetName := fmt.Sprintf("%sCustomVnet", cli.Config.Name)
	masterSubnetName := fmt.Sprintf("%sCustomSubnetMaster", cli.Config.Name)
	masterSubnetID := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/virtualNetworks/%s/subnets/%s", cli.Account.SubscriptionID, cli.Account.ResourceGroup.Name, vnetName, masterSubnetName)
	agentSubnetID := ""
	agentSubnetIDs := []string{}
	subnets := []string{}
	config, err := engine.ParseConfig(cli.Config.CurrentWorkingDir, cli.Config.ClusterDefinition, cli.Config.Name)
	if err != nil {
		log.Printf("Error while trying to build Engine Configuration:%s\n", err)
	}
	cs, err := engine.ParseInput(config.ClusterDefinitionPath)
	if err != nil {
		return err
	}

	if cli.CreateVNET {
		if cli.MasterVMSS {
			agentSubnetName := fmt.Sprintf("%sCustomSubnetAgent", cli.Config.Name)
			err = cli.Account.CreateVnet(vnetName, "10.239.0.0/16")
			if err != nil {
				return errors.Errorf("Error trying to create vnet:%s", err.Error())
			}
			err = cli.Account.CreateSubnet(vnetName, masterSubnetName, "10.239.0.0/17")
			if err != nil {
				return errors.Errorf("Error trying to create subnet:%s", err.Error())
			}
			subnets = append(subnets, masterSubnetName)
			err = cli.Account.CreateSubnet(vnetName, agentSubnetName, "10.239.128.0/17")
			if err != nil {
				return errors.Errorf("Error trying to create subnet in subnet:%s", err.Error())
			}
			subnets = append(subnets, agentSubnetName)
			agentSubnetID = fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/virtualNetworks/%s/subnets/%s", cli.Account.SubscriptionID, cli.Account.ResourceGroup.Name, vnetName, agentSubnetName)

		} else {
			err = cli.Account.CreateVnet(vnetName, "10.239.0.0/16")
			if err != nil {
				return errors.Errorf("Error trying to create vnet:%s", err.Error())
			}
			err = cli.Account.CreateSubnet(vnetName, masterSubnetName, "10.239.255.0/24")
			if err != nil {
				return errors.Errorf("Error trying to create subnet:%s", err.Error())
			}
			subnets = append(subnets, masterSubnetName)
			for i, pool := range cs.ContainerService.Properties.AgentPoolProfiles {
				subnetName := fmt.Sprintf("%sCustomSubnet", pool.Name)
				err = cli.Account.CreateSubnet(vnetName, subnetName, fmt.Sprintf("10.239.%d.0/20", i*16))
				if err != nil {
					return errors.Errorf("Error trying to create subnet:%s", err.Error())
				}
				subnets = append(subnets, subnetName)
				subnetID = fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/virtualNetworks/%s/subnets/%s", cli.Account.SubscriptionID, cli.Account.ResourceGroup.Name, vnetName, subnetName)
				agentSubnetIDs = append(agentSubnetIDs, subnetID)
			}
		}
	}

	// Lets modify our template and call aks-engine generate on it
	var eng *engine.Engine

	if cli.CreateVNET && cli.MasterVMSS {
		eng, err = engine.Build(cli.Config, masterSubnetID, []string{agentSubnetID}, true)
	} else {
		eng, err = engine.Build(cli.Config, masterSubnetID, agentSubnetIDs, false)
	}

	if err != nil {
		return errors.Wrap(err, "Error while trying to build cluster definition")
	}
	cli.Engine = eng

	cli.EnsureArcResourceGroup()

	err = cli.Engine.Write()
	if err != nil {
		return errors.Wrap(err, "Error while trying to write Engine Template to disk:%s")
	}

	err = cli.generateAndDeploy()
	if err != nil {
		return errors.Wrap(err, "Error in generateAndDeploy:%s")
	}

	if cs.Properties.OrchestratorProfile != nil && cs.Properties.OrchestratorProfile.KubernetesConfig != nil {
		if cli.CreateVNET && cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin == "kubenet" {
			routeTable, err := cli.Account.GetRGRouteTable(10 * time.Minute)
			if err != nil {
				return errors.Errorf("Error trying to get route table in VNET: %s", err.Error())
			}
			err = cli.Account.AddSubnetsToRouteTable(*routeTable.ID, vnetName, subnets)
			if err != nil {
				return errors.Errorf("Error trying to add subnets to route table %s in VNET: %s", *routeTable.ID, err.Error())
			}
		}
	}

	if cli.Config.IsKubernetes() && !cli.Config.SkipTest {
		// Store the hosts for future introspection
		hosts, err := cli.Account.GetHosts(cli.Config.Name)
		if err != nil {
			return errors.Wrap(err, "GetHosts:%s")
		}
		var masters, agents []azure.VM
		for _, host := range hosts {
			if strings.Contains(host.Name, "master") {
				masters = append(masters, host)
			} else if strings.Contains(host.Name, "agent") {
				agents = append(agents, host)
			}
		}
		cli.Masters = masters
		cli.Agents = agents
	}

	return nil
}

func (cli *CLIProvisioner) generateAndDeploy() error {
	if cli.Config.UseDeployCommand {
		fmt.Printf("Provisionning with the Deploy Command\n")
		err := cli.Engine.Deploy(cli.Config.Location)
		if err != nil {
			return errors.Wrap(err, "Error while trying to deploy aks-engine template")
		}
	} else {
		err := cli.Engine.Generate()
		if err != nil {
			return errors.Wrap(err, "Error while trying to generate aks-engine template")
		}
	}

	c, err := config.ParseConfig()
	if err != nil {
		return errors.Wrap(err, "unable to parse base config")
	}
	engCfg, err := engine.ParseConfig(cli.Config.CurrentWorkingDir, c.ClusterDefinition, c.Name)
	if err != nil {
		return errors.Wrap(err, "unable to parse config")
	}
	validate := true
	isUpdate := false
	csGenerated, err := engine.ParseOutput(engCfg.GeneratedDefinitionPath+"/apimodel.json", validate, isUpdate)
	if err != nil {
		return errors.Wrap(err, "unable to parse output")
	}
	cli.Engine.ExpandedDefinition = csGenerated

	// Kubernetes deployments should have a kubeconfig available
	// at this point.
	if cli.Config.IsKubernetes() && !cli.IsPrivate() {
		cli.Config.SetKubeConfig()
	}

	//if we use Generate, then we need to call CreateDeployment
	if !cli.Config.UseDeployCommand {
		err = cli.Account.CreateDeployment(cli.Config.Name, cli.Engine)
		if err != nil {
			return errors.Wrap(err, "Error while trying to create deployment")
		}
	}
	return err
}

// GenerateName will generate a new name if one has not been set
func (cli *CLIProvisioner) generateName() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	suffix := r.Intn(99999)
	prefix := fmt.Sprintf("%s-%s", cli.Config.Orchestrator, cli.Config.Location)
	return fmt.Sprintf("%s-%v", prefix, suffix)
}

func (cli *CLIProvisioner) waitForNodes() error {
	if cli.Config.IsKubernetes() {
		if !cli.IsPrivate() {
			log.Println("Waiting on nodes to go into ready state...")
			var expectedReadyNodes int
			if !cli.Engine.ExpandedDefinition.Properties.HasNonRegularPriorityScaleset() && !cli.Config.RebootControlPlaneNodes {
				expectedReadyNodes = cli.Engine.NodeCount()
				log.Printf("Checking for %d Ready nodes\n", expectedReadyNodes)
			} else {
				expectedReadyNodes = -1
			}
			ready := node.WaitOnReady(expectedReadyNodes, 10*time.Second, cli.Config.Timeout)
			cmd := exec.Command("k", "get", "nodes", "-o", "wide")
			out, _ := cmd.CombinedOutput()
			log.Printf("%s\n", out)
			if !ready {
				return errors.New("Error: Not all nodes in a healthy state")
			}
			nodes, err := node.GetWithRetry(1*time.Second, cli.Config.Timeout)
			if err != nil {
				return errors.Wrap(err, "Unable to get the list of nodes")
			}
			if !cli.Engine.ExpandedDefinition.Properties.HasNonRegularPriorityScaleset() {
				for _, n := range nodes {
					exp, err := regexp.Compile(common.LegacyControlPlaneVMPrefix)
					if err != nil {
						return err
					}
					if !exp.MatchString(n.Metadata.Name) {
						cmd := exec.Command("k", "label", "node", n.Metadata.Name, "foo=bar")
						util.PrintCommand(cmd)
						out, err := cmd.CombinedOutput()
						log.Printf("%s\n", out)
						if err != nil {
							return errors.Wrapf(err, "Unable to assign label to node %s", n.Metadata.Name)
						}
						cmd = exec.Command("k", "annotate", "node", n.Metadata.Name, "foo=bar")
						util.PrintCommand(cmd)
						out, err = cmd.CombinedOutput()
						log.Printf("%s\n", out)
						if err != nil {
							return errors.Wrapf(err, "Unable to add node annotation to node %s", n.Metadata.Name)
						}
					}
				}
			}
		} else {
			log.Println("This cluster is private")
			if cli.Engine.ClusterDefinition.Properties.OrchestratorProfile.KubernetesConfig.PrivateCluster.JumpboxProfile == nil {
				// TODO: add "bring your own jumpbox to e2e"
				return errors.New("Error: cannot test a private cluster without provisioning a jumpbox")
			}
			log.Printf("Testing a %s private cluster...", cli.Config.Orchestrator)
			// TODO: create SSH connection and get nodes and k8s version
		}
	}

	return nil
}

// FetchProvisioningMetrics gets provisioning files from all hosts in a cluster
func (cli *CLIProvisioner) FetchProvisioningMetrics(path string, cfg *config.Config, acct *azure.Account) error {
	agentFiles := []string{"/var/log/azure/cluster-provision.log", "/var/log/cloud-init.log",
		"/var/log/cloud-init-output.log", "/var/log/syslog", "/var/log/azure/custom-script/handler.log",
		"/opt/m", "/opt/azure/containers/kubelet.sh", "/opt/azure/containers/provision.sh",
		"/var/log/azure/kubelet-status.log", "/var/log/azure/docker-status.log", "/var/log/azure/systemd-journald-status.log"}
	masterFiles := agentFiles
	masterFiles = append(masterFiles, "/opt/azure/containers/setup-etcd.sh", "/opt/azure/containers/setup-etcd.log")
	hostname := fmt.Sprintf("%s.%s.cloudapp.azure.com", cli.Config.Name, cli.Config.Location)
	cmd := exec.Command("ssh-agent", "-s")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrapf(err, "Error while trying to start ssh agent \nOutput:%s", out)
	}
	authSock := strings.Split(strings.Split(string(out), "=")[1], ";")
	os.Setenv("SSH_AUTH_SOCK", authSock[0])
	var conn *remote.Connection
	conn, err = remote.NewConnectionWithRetry(hostname, "22", cli.Engine.ClusterDefinition.Properties.LinuxProfile.AdminUsername, cli.Config.GetSSHKeyPath(), 3*time.Second, cli.Config.Timeout)
	if err != nil {
		return err
	}
	for _, master := range cli.Masters {
		for _, fp := range masterFiles {
			err = conn.CopyFrom(master.Name, fp)
			if err != nil {
				log.Printf("Error reading file from path (%s):%s", path, err)
			}
		}
	}

	for _, agent := range cli.Agents {
		for _, fp := range agentFiles {
			err = conn.CopyFrom(agent.Name, fp)
			if err != nil {
				log.Printf("Error reading file from path (%s):%s", path, err)
			}
		}
	}
	connectString := fmt.Sprintf("%s@%s:/tmp/k8s-*", conn.User, hostname)
	logsPath := filepath.Join(cfg.CurrentWorkingDir, "_logs", hostname)
	cmd = exec.Command("scp", "-i", conn.PrivateKeyPath, "-o", "ConnectTimeout=30", "-o", "StrictHostKeyChecking=no", connectString, logsPath)
	util.PrintCommand(cmd)
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error output:%s\n", out)
		return err
	}

	return nil
}

// IsPrivate will return true if the cluster has no public IPs
func (cli *CLIProvisioner) IsPrivate() bool {
	return cli.Config.IsKubernetes() &&
		cli.Engine.ExpandedDefinition.Properties.OrchestratorProfile.KubernetesConfig.PrivateCluster != nil &&
		to.Bool(cli.Engine.ExpandedDefinition.Properties.OrchestratorProfile.KubernetesConfig.PrivateCluster.Enabled)
}

// FetchActivityLog gets the activity log for the all resource groups used in the provisioner.
func (cli *CLIProvisioner) FetchActivityLog(acct *azure.Account, logPath string) error {
	for _, rg := range cli.ResourceGroups {
		log, err := acct.FetchActivityLog(rg)
		if err != nil {
			return errors.Wrapf(err, "cannot fetch activity log for resource group %s", rg)
		}
		path := filepath.Join(logPath, fmt.Sprintf("activity-log-%s", rg))
		if err := ioutil.WriteFile(path, []byte(log), 0644); err != nil {
			return errors.Wrap(err, "cannot write activity log in file")
		}
	}
	return nil
}

// EnsureArcResourceGroup creates the resource group for the connected cluster resource
// Once Arc is supported in all regions, we should delete this method and reuse the cluster resource group
// https://docs.microsoft.com/en-us/azure/azure-arc/kubernetes/overview#supported-regions
func (cli *CLIProvisioner) EnsureArcResourceGroup() error {
	for _, addon := range cli.Engine.ClusterDefinition.Properties.OrchestratorProfile.KubernetesConfig.Addons {
		if addon.Name == common.AzureArcOnboardingAddonName && to.Bool(addon.Enabled) {
			if err := cli.Account.CreateGroupWithRetry(addon.Config["resourceGroup"], addon.Config["location"], 30*time.Second, cli.Config.Timeout); err != nil {
				return errors.Wrapf(err, "Error while trying to create Azure Arc resource group: %s", addon.Config["resourceGroup"])
			}
		}
	}
	return nil
}
