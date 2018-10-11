package runner

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/Azure/aks-engine/pkg/helpers"
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
			cli.Point.SetNodeWaitStart()
			err := cli.waitForNodes()
			cli.Point.RecordNodeWait(err)
			return err
		}
	}
	return errors.New("Unable to run provisioner")
}

func createSaveSSH(outputPath string, privateKeyName string) (string, error) {
	os.Mkdir(outputPath, 0755)
	keyPath := filepath.Join(outputPath, privateKeyName)
	cmd := exec.Command("ssh-keygen", "-f", keyPath, "-q", "-N", "", "-b", "2048", "-t", "rsa")

	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.Wrapf(err, "Error while trying to generate ssh key\nOutput:%s", out)
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
		publicSSHKey, err := createSaveSSH(outputPath, cli.Config.Name+"-ssh")
		if err != nil {
			return errors.Wrap(err, "Error while generating ssh keys")
		}
		os.Setenv("PUBLIC_SSH_KEY", publicSSHKey)
	}

	os.Setenv("DNS_PREFIX", cli.Config.Name)

	err := cli.Account.CreateGroup(cli.Config.Name, cli.Config.Location)
	if err != nil {
		return errors.Wrap(err, "Error while trying to create resource group")
	}

	subnetID := ""
	vnetName := fmt.Sprintf("%sCustomVnet", cli.Config.Name)
	subnetName := fmt.Sprintf("%sCustomSubnet", cli.Config.Name)
	masterSubnetID := ""
	agentSubnetID := ""

	if cli.CreateVNET {
		if cli.MasterVMSS {
			masterSubnetName := fmt.Sprintf("%sCustomSubnetMaster", cli.Config.Name)
			agentSubnetName := fmt.Sprintf("%sCustomSubnetAgent", cli.Config.Name)
			err = cli.Account.CreateVnet(vnetName, "10.239.0.0/16", masterSubnetName, "10.239.0.0/17")
			if err != nil {
				return errors.Errorf("Error trying to create vnet:%s", err.Error())
			}

			masterSubnetID = fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/virtualNetworks/%s/subnets/%s", cli.Account.SubscriptionID, cli.Account.ResourceGroup.Name, vnetName, masterSubnetName)

			err = cli.Account.CreateSubnet(vnetName, agentSubnetName, "10.239.128.0/17")
			if err != nil {
				return errors.Errorf("Error trying to create subnet in vnet:%s", err.Error())
			}

			agentSubnetID = fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/virtualNetworks/%s/subnets/%s", cli.Account.SubscriptionID, cli.Account.ResourceGroup.Name, vnetName, agentSubnetName)
		} else {
			err = cli.Account.CreateVnet(vnetName, "10.239.0.0/16", subnetName, "10.239.0.0/16")
			if err != nil {
				return errors.Errorf("Error trying to create vnet:%s", err.Error())
			}
			subnetID = fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/virtualNetworks/%s/subnets/%s", cli.Account.SubscriptionID, cli.Account.ResourceGroup.Name, vnetName, subnetName)
		}
	}

	// Lets modify our template and call aks-engine generate on it
	var eng *engine.Engine

	if cli.CreateVNET && cli.MasterVMSS {
		eng, err = engine.Build(cli.Config, masterSubnetID, agentSubnetID, true)
	} else {
		eng, err = engine.Build(cli.Config, subnetID, subnetID, false)
	}

	if err != nil {
		return errors.Wrap(err, "Error while trying to build cluster definition")
	}
	cli.Engine = eng

	err = cli.Engine.Write()
	if err != nil {
		return errors.Wrap(err, "Error while trying to write Engine Template to disk:%s")
	}

	err = cli.generateAndDeploy()
	if err != nil {
		return errors.Wrap(err, "Error in generateAndDeploy:%s")
	}

	if cli.Config.IsKubernetes() {
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
	csGenerated, err := engine.ParseOutput(engCfg.GeneratedDefinitionPath + "/apimodel.json")
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
			ready := node.WaitOnReady(cli.Engine.NodeCount(), 10*time.Second, cli.Config.Timeout)
			if !ready {
				return errors.New("Error: Not all nodes in a healthy state")
			}
			var version string
			var err error
			if cli.Config.IsKubernetes() {
				version, err = node.Version()
			}
			if err != nil {
				log.Printf("Ready nodes did not return a version: %s", err)
			}
			log.Printf("Testing a %s %s cluster...\n", cli.Config.Orchestrator, version)
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
		"/opt/azure/provision-ps.log", "/var/log/azure/kubelet-status.log",
		"/var/log/azure/docker-status.log", "/var/log/azure/systemd-journald-status.log"}
	masterFiles := agentFiles
	masterFiles = append(masterFiles, "/opt/azure/containers/mountetcd.sh", "/opt/azure/containers/setup-etcd.sh", "/opt/azure/containers/setup-etcd.log")
	hostname := fmt.Sprintf("%s.%s.cloudapp.azure.com", cli.Config.Name, cli.Config.Location)
	cmd := exec.Command("ssh-agent", "-s")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrapf(err, "Error while trying to start ssh agent \nOutput:%s", out)
	}
	authSock := strings.Split(strings.Split(string(out), "=")[1], ";")
	os.Setenv("SSH_AUTH_SOCK", authSock[0])
	conn, err := remote.NewConnection(hostname, "22", cli.Engine.ClusterDefinition.Properties.LinuxProfile.AdminUsername, cli.Config.GetSSHKeyPath())
	if err != nil {
		return err
	}
	for _, master := range cli.Masters {
		for _, fp := range masterFiles {
			err := conn.CopyRemote(master.Name, fp)
			if err != nil {
				log.Printf("Error reading file from path (%s):%s", path, err)
			}
		}
	}

	for _, agent := range cli.Agents {
		for _, fp := range agentFiles {
			err := conn.CopyRemote(agent.Name, fp)
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
		helpers.IsTrueBoolPointer(cli.Engine.ExpandedDefinition.Properties.OrchestratorProfile.KubernetesConfig.PrivateCluster.Enabled)
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
