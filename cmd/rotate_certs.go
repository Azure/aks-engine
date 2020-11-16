// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/Azure/aks-engine/cmd/rotatecerts"
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/engine"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/aks-engine/pkg/helpers/ssh"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/Azure/aks-engine/pkg/kubernetes"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	sshx "golang.org/x/crypto/ssh"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	rotateCertsName             = "rotate-certs"
	rotateCertsShortDescription = "(experimental) Rotate certificates on an existing AKS Engine-created Kubernetes cluster"
	rotateCertsLongDescription  = "(experimental) Rotate CA, etcd, kubelet, kubeconfig and apiserver certificates in a cluster built with AKS Engine. Rotating certificates can break component connectivity and leave the cluster in an unrecoverable state. Before performing any of these instructions on a live cluster, it is preferrable to backup your cluster state and migrate critical workloads to another cluster."
)

const (
	rootUserGroup     = "root:root"
	etcdUserGroup     = "etcd:etcd"
	keyPermissions    = "600"
	crtPermissions    = "644"
	configPermissions = "600"

	kubeAPIServer         = "kube-apiserver"
	kubeAddonManager      = "kube-addon-manager"
	kubeControllerManager = "kube-controller-manager"
	kubeScheduler         = "kube-scheduler"

	kubeProxyLabels     = "component=kube-proxy,k8s-app=kube-proxy,tier=node"
	kubeSchedulerLabels = "component=kube-scheduler,tier=control-plane"

	rotateCertsDefaultInterval = 10 * time.Second
	rotateCertsDefaultTimeout  = 10 * time.Minute
)

type rotateCertsCmd struct {
	// user input
	location               string
	apiModelPath           string
	newCertsPath           string
	sshHostURI             string
	linuxSSHPrivateKeyPath string
	outputDirectory        string
	continueAfterError     bool
	generateCerts          bool
	abort                  bool
	// computed
	backupDirectory   string
	apiVersion        string
	cs                *api.ContainerService
	loader            *api.Apiloader
	newCertsProfile   *api.CertificateProfile
	kubeClient        rotatecerts.Client
	namespaces        []string
	nodes             []*ssh.RemoteHost
	linuxAuthConfig   *ssh.AuthConfig
	windowsAuthConfig *ssh.AuthConfig
	jumpbox           *ssh.JumpBox
}

func newRotateCertsCmd() *cobra.Command {
	rcc := rotateCertsCmd{}
	command := &cobra.Command{
		Use:   rotateCertsName,
		Short: rotateCertsShortDescription,
		Long:  rotateCertsLongDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := rcc.validateArgs(); err != nil {
				return errors.Wrap(err, "validating rotate-certs args")
			}
			if err := rcc.loadAPIModel(); err != nil {
				return errors.Wrap(err, "loading API model")
			}
			if err := rcc.init(); err != nil {
				return errors.Wrap(err, "loading API model")
			}
			cmd.SilenceUsage = true
			return rcc.run()
		},
	}
	command.Flags().StringVarP(&rcc.location, "location", "l", "", "Azure location where the cluster is deployed")
	command.Flags().StringVarP(&rcc.apiModelPath, "api-model", "m", "", "path to the generated apimodel.json file")
	command.Flags().StringVar(&rcc.sshHostURI, "ssh-host", "", "FQDN, or IP address, of an SSH listener that can reach all nodes in the cluster")
	command.Flags().StringVar(&rcc.linuxSSHPrivateKeyPath, "linux-ssh-private-key", "", "path to a valid private SSH key to access the cluster's Linux nodes")
	_ = command.MarkFlagRequired("location")
	_ = command.MarkFlagRequired("api-model")
	_ = command.MarkFlagRequired("ssh-host")
	_ = command.MarkFlagRequired("linux-ssh-private-key")

	command.Flags().StringVarP(&rcc.newCertsPath, "certificate-profile", "", "", "path to a JSON file containing the new set of certificates")
	command.Flags().BoolVarP(&rcc.generateCerts, "generate-new-certificates", "", false, "generate a new set of certificates, required if a --certificate-profile is not provided")
	command.Flags().BoolVarP(&rcc.continueAfterError, "continue", "", false, "resume a previous execution that did not complete successfully")
	command.Flags().BoolVarP(&rcc.abort, "abort", "", false, "only clean temporary artifacts from the cluster nodes")
	return command
}

func (rcc *rotateCertsCmd) validateArgs() (err error) {
	locale, err := i18n.LoadTranslations()
	if err != nil {
		return errors.Wrap(err, "loading translation files")
	}
	rcc.loader = &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: locale,
		},
	}
	rcc.location = helpers.NormalizeAzureRegion(rcc.location)
	if rcc.location == "" {
		return errors.New("--location must be specified")
	}
	if rcc.sshHostURI == "" {
		return errors.New("--ssh-host must be specified")
	}
	if rcc.linuxSSHPrivateKeyPath == "" {
		return errors.New("--linux-ssh-private-key must be specified")
	} else if _, err = os.Stat(rcc.linuxSSHPrivateKeyPath); os.IsNotExist(err) {
		return errors.Errorf("specified --linux-ssh-private-key does not exist (%s)", rcc.linuxSSHPrivateKeyPath)
	}
	if rcc.apiModelPath == "" {
		return errors.New("--api-model must be specified")
	} else if _, err = os.Stat(rcc.apiModelPath); os.IsNotExist(err) {
		return errors.Errorf("specified --api-model does not exist (%s)", rcc.apiModelPath)
	}

	if rcc.newCertsPath != "" {
		if _, err = os.Stat(rcc.newCertsPath); os.IsNotExist(err) {
			return errors.Errorf("specified --certificate-profile does not exist (%s)", rcc.newCertsPath)
		}
	}
	if rcc.newCertsPath == "" && !rcc.generateCerts {
		return errors.New("either --generate-new-certificates or --certificate-profile should be specified")
	}
	if rcc.newCertsPath != "" && rcc.generateCerts {
		rcc.generateCerts = false
		log.Infof("ignoring --generate-new-certificates as --certificate-profile is specified")
	}
	if rcc.outputDirectory == "" {
		rcc.outputDirectory = path.Join(filepath.Dir(rcc.apiModelPath), "_rotate_certs_output")
		if err = os.MkdirAll(rcc.outputDirectory, 0755); err != nil {
			return errors.Errorf("error creating output directory (%s)", rcc.outputDirectory)
		}
	}
	fs, err := ioutil.ReadDir(rcc.outputDirectory)
	if err != nil {
		return errors.Wrapf(err, "reading output directory %s", rcc.outputDirectory)
	}
	if len(fs) > 0 && !rcc.continueAfterError {
		return errors.Errorf("output directory %s is not empty, this indicates that a previous rotate-certs execution did not complete successfully, set --continue to resume the process", rcc.outputDirectory)
	}
	if len(fs) == 0 && rcc.continueAfterError {
		return errors.Errorf("output directory %s is empty, --continue should be unset", rcc.outputDirectory)
	}
	return nil
}

func (rcc *rotateCertsCmd) loadAPIModel() (err error) {
	if rcc.cs, rcc.apiVersion, err = rcc.loader.LoadContainerServiceFromFile(rcc.apiModelPath, false, false, nil); err != nil {
		return errors.Wrap(err, "error parsing api-model")
	}
	if rcc.newCertsPath != "" {
		// TODO validate certificates metadata
		if rcc.newCertsProfile, err = rcc.loader.LoadCertificateProfileFromFile(rcc.newCertsPath); err != nil {
			return errors.Wrap(err, "error parsing certificate-profile")
		}
	}
	if rcc.cs.Location == "" {
		rcc.cs.Location = rcc.location
	} else if rcc.cs.Location != rcc.location {
		return errors.New("--location flag does not match api-model location")
	}
	if rcc.cs.Properties.WindowsProfile != nil && !rcc.cs.Properties.WindowsProfile.GetSSHEnabled() {
		return errors.New("SSH not enabled on Windows nodes. SSH is required in order to rotate agent nodes certificates")
	}

	return
}

func (rcc *rotateCertsCmd) init() (err error) {
	rcc.namespaces = rcc.getNamespacesWithSATokensToRotate()
	rcc.backupDirectory = path.Join(filepath.Dir(rcc.apiModelPath), "_rotate_certs_backup")

	rcc.linuxAuthConfig = &ssh.AuthConfig{
		User:           rcc.cs.Properties.LinuxProfile.AdminUsername,
		PrivateKeyPath: rcc.linuxSSHPrivateKeyPath,
	}
	if rcc.cs.Properties.WindowsProfile != nil {
		rcc.windowsAuthConfig = &ssh.AuthConfig{
			User:     rcc.cs.Properties.WindowsProfile.AdminUsername,
			Password: rcc.cs.Properties.WindowsProfile.AdminPassword,
		}
	}
	rcc.jumpbox = &ssh.JumpBox{URI: rcc.sshHostURI, Port: 22, OperatingSystem: api.Linux, AuthConfig: rcc.linuxAuthConfig}
	return
}

func (rcc *rotateCertsCmd) run() (err error) {
	err = rcc.backupCerts()
	if err != nil {
		return errors.Wrap(err, "backing up current state")
	}
	err = rcc.updateCertificateProfile()
	if err != nil {
		return errors.Wrap(err, "backing up current state")
	}
	rcc.kubeClient, err = rcc.getKubeClient()
	if err != nil {
		return errors.Wrap(err, "creating Kubernetes client")
	}
	rcc.nodes, err = rcc.getClusterNodes(rcc.kubeClient)
	if err != nil {
		return errors.Wrap(err, "listing cluster nodes")
	}

	if rcc.abort {
		err = rcc.cleanupRemote()
		if err != nil {
			return errors.Wrapf(err, "aborting previous %s execution", rotateCertsName)
		}
		log.Infoln("Abort completed")
		return nil
	}

	// TODO avoid re-uploads
	err = rcc.distributeCerts()
	if err != nil {
		return errors.Wrap(err, "distributing certificates")
	}
	err = rcc.distributeScripts()
	if err != nil {
		return errors.Wrap(err, "distributing scripts")
	}

	// ctx := context.Background()
	// defer ctx.Done()
	// go logDowntime(ctx, rcc)

	err = rcc.rotateCerts()
	if err != nil {
		return errors.Wrap(err, "rotating certificates")
	}
	err = rcc.updateAPIModel()
	if err != nil {
		return errors.Wrap(err, "rotating certificates")
	}

	log.Infoln("Certificate rotation completed")
	return nil
}

func (rcc *rotateCertsCmd) backupCerts() error {
	log.Infof("Backing up artifacts to directory %s", rcc.backupDirectory)
	err := writeArtifacts(rcc.backupDirectory, rcc.cs, rcc.apiVersion, rcc.loader.Translator)
	if err != nil {
		return errors.Wrap(err, "writing artifacts")
	}
	return nil
}

func (rcc *rotateCertsCmd) updateCertificateProfile() error {
	if rcc.continueAfterError {
		err := rcc.loadCertificateProfile(rcc.outputDirectory)
		if err != nil {
			return errors.Wrap(err, "loading artifacts")
		}
	} else {
		if rcc.generateCerts {
			err := rcc.generateTLSArtifacts()
			if err != nil {
				return errors.Wrap(err, "generating artifacts")
			}
		} else if rcc.newCertsPath != "" {
			rcc.cs.Properties.CertificateProfile = rcc.newCertsProfile
		}
		log.Infof("Writing artifacts to output directory %s", rcc.outputDirectory)
		err := writeArtifacts(rcc.outputDirectory, rcc.cs, rcc.apiVersion, rcc.loader.Translator)
		if err != nil {
			return errors.Wrap(err, "writing artifacts")
		}
	}
	return nil
}

func (rcc *rotateCertsCmd) loadCertificateProfile(dir string) error {
	log.Infof("Loading certificates from output directory: %s", dir)
	path := path.Join(dir, "apimodel.json")
	cs, _, err := rcc.loader.LoadContainerServiceFromFile(path, false, false, nil)
	if err != nil {
		return errors.Wrap(err, "error parsing api-model")
	}
	rcc.cs.Properties.CertificateProfile = cs.Properties.CertificateProfile
	return nil
}

func (rcc *rotateCertsCmd) generateTLSArtifacts() error {
	log.Infoln("Generating new certificates")
	rcc.cs.Properties.CertificateProfile = &api.CertificateProfile{}
	if ok, _, err := rcc.cs.SetDefaultCerts(api.DefaultCertParams{PkiKeySize: helpers.DefaultPkiKeySize}); !ok || err != nil {
		return errors.Wrap(err, "generating new certificates")
	}
	return nil
}

// getClusterNodes returns all cluster nodes or an error if a node status is not Ready
func (rcc *rotateCertsCmd) getClusterNodes(kubeClient kubernetes.NodeLister) (nodes []*ssh.RemoteHost, err error) {
	nodeList, err := kubeClient.ListNodes()
	if err != nil {
		return nil, err
	}
	for _, node := range nodeList.Items {
		switch api.OSType(strings.Title(node.Status.NodeInfo.OperatingSystem)) {
		case api.Linux:
			nodes = append(nodes, &ssh.RemoteHost{
				URI: node.Name, Port: 22, OperatingSystem: api.Linux, AuthConfig: rcc.linuxAuthConfig, Jumpbox: rcc.jumpbox})
		case api.Windows:
			nodes = append(nodes, &ssh.RemoteHost{
				URI: node.Name, Port: 22, OperatingSystem: api.Windows, AuthConfig: rcc.windowsAuthConfig, Jumpbox: rcc.jumpbox})
		default:
			return nil, errors.Errorf("Skipping node %s, could not determine operating system", node.Name)
		}
	}
	nodeCount := rcc.cs.Properties.MasterProfile.Count
	for _, pool := range rcc.cs.Properties.AgentPoolProfiles {
		nodeCount += pool.Count
	}
	if len(nodes) != nodeCount {
		return nil, errors.Errorf("Unexpected node count. Got %d, expected %d (from api-model)", nodeCount, len(nodes))
	}
	return nodes, nil
}

// distributeCerts copies the new set of certificates to the cluster nodes.
func (rcc *rotateCertsCmd) distributeCerts() (err error) {
	log.Info("Distributing certificates")
	masterCerts, linuxCerts, windowsCerts := getFilesToDistribute(rcc.cs)
	for _, node := range rcc.nodes {
		log.Debugf("Uploading certificates to node %s", node.URI)
		if isMaster(node) {
			for _, file := range masterCerts {
				co, err := ssh.CopyToRemote(node, file)
				if err != nil {
					log.Debugf("Remote command output: %s", co)
					return errors.Wrap(err, "uploading certificate")
				}
			}
		} else if isLinuxAgent(node) {
			for name := range linuxCerts {
				co, err := ssh.CopyToRemote(node, linuxCerts[name])
				if err != nil {
					log.Debugf("Remote command output: %s", co)
					return errors.Wrap(err, "uploading certificate")
				}
			}
		} else if isWindows(node) {
			for _, file := range windowsCerts {
				co, err := ssh.CopyToRemote(node, file)
				if err != nil {
					log.Debugf("Remote command output: %s", co)
					return errors.Wrap(err, "uploading certificate")
				}
			}
		}
	}
	return
}

// distributeScripts copies the corresponding certificate rotation script to the cluster nodes.
func (rcc *rotateCertsCmd) distributeScripts() (err error) {
	lf, err := engine.Asset("k8s/cloud-init/artifacts/rotate-certs.sh")
	if err != nil {
		return err
	}
	linuxFile := ssh.NewRemoteFile("/tmp/akse/rotate-certs.sh", "744", rootUserGroup, lf)
	wf, err := engine.Asset("k8s/cloud-init/artifacts/rotate-certs.ps1")
	if err != nil {
		return err
	}
	windowsFile := ssh.NewRemoteFile("$env:temp\\rotate-certs.ps1", "", "", wf)

	for _, n := range rcc.nodes {
		if isLinux(n) {
			out, err := ssh.CopyToRemote(n, linuxFile)
			if err != nil {
				log.Debugf("Remote command output: %s", out)
				return errors.Wrap(err, "uploading Linux certiciate rotation script")
			}
		} else {
			out, err := ssh.CopyToRemote(n, windowsFile)
			if err != nil {
				log.Debugf("Remote command output: %s", out)
				return errors.Wrap(err, "uploading Windows certificate rotation script")
			}
		}
	}
	return
}

func (rcc *rotateCertsCmd) rotateCerts() error {
	log.Info("Rotating certificates")
	if err := rotatecerts.WaitForNodesReady(rcc.kubeClient, rcc.nodes, 5, rotateCertsDefaultInterval, rotateCertsDefaultTimeout); err != nil {
		return errors.Wrap(err, "waiting for cluster nodes readiness")
	}
	if err := rcc.waitForControlPlaneReadiness(); err != nil {
		return err
	}
	if err := rcc.backupRemote(); err != nil {
		return err
	}
	if err := rcc.rotateEtcd(); err != nil {
		return err
	}
	if err := rcc.rotateSATokens(); err != nil {
		return err
	}
	if err := rcc.rotateKubelet(); err != nil {
		return err
	}
	if err := rcc.cleanupRemote(); err != nil {
		return err
	}
	return nil
}

func (rcc *rotateCertsCmd) backupRemote() (err error) {
	log.Info("Backing up node certificates")
	for _, node := range rcc.nodes {
		step := "backup"
		skipped, err := execStepsSequence(isLinux, node, execRemoteFunc(remoteBashScript(step)))
		if !skipped && err != nil {
			return errors.Wrapf(err, "executing %s function on remote host %s", step, node.URI)
		}
	}
	return nil
}

func (rcc *rotateCertsCmd) rotateEtcd() error {
	log.Infoln("Rotating etcd PKI")
	var err error
	var skipped bool

	log.Debugln("Replacing etcd ca with a ca bundle")
	waitForRestart, restartAfter := false, time.Now()
	for _, node := range rcc.nodes {
		step := "etcd_cabundle"
		skipped, err = execStepsSequence(isMaster, node, execRemoteFunc(remoteBashScript(step)))
		if !skipped && err != nil {
			return errors.Wrapf(err, "executing %s function on remote host %s", step, node.URI)
		}
		waitForRestart = waitForRestart || (!skipped && err == nil)
	}
	// No need to wait if all nodes were skipped
	if waitForRestart {
		err = rcc.waitForControlPlaneRestart(restartAfter, kubeAPIServer)
		if err != nil {
			return err
		}
		err = rcc.waitForControlPlaneReadiness()
		if err != nil {
			return err
		}
	}

	log.Debugln("Rotating etcd non-ca certs")
	for _, node := range rcc.nodes {
		step := "etcd_certs"
		skipped, err = execStepsSequence(isMaster, node, execRemoteFunc(remoteBashScript(step)))
		if !skipped && err != nil {
			return errors.Wrapf(err, "executing %s function on remote host %s", step, node.URI)
		}
	}

	log.Debugln("Rotating etcd ca certs")
	waitForRestart, restartAfter = false, time.Now()
	for _, node := range rcc.nodes {
		step := "etcd_ca"
		skipped, err = execStepsSequence(isMaster, node, execRemoteFunc(remoteBashScript(step)), deletePodFunc(rcc.kubeClient, kubeSchedulerLabels))
		if !skipped && err != nil {
			return errors.Wrapf(err, "executing %s function on remote host %s", step, node.URI)
		}
		waitForRestart = waitForRestart || (!skipped && err == nil)
	}
	// No need to wait if all nodes were skipped
	if waitForRestart {
		err = rcc.waitForControlPlaneRestart(restartAfter, kubeAPIServer)
		if err != nil {
			return err
		}
		err = rcc.waitForControlPlaneReadiness()
		if err != nil {
			return err
		}
	}
	return nil
}

func (rcc *rotateCertsCmd) rotateSATokens() error {
	log.Infoln("Rotating service account tokens")
	var err error
	var skipped bool

	log.Debugln("Rotating service account signer")
	waitForRestart, restartAfter := false, time.Now()
	for _, node := range rcc.nodes {
		step := "sa_token_signer"
		skipped, err = execStepsSequence(isMaster, node, execRemoteFunc(remoteBashScript("sa_token_signer")))
		if !skipped && err != nil {
			return errors.Wrapf(err, "executing %s function on remote host %s", step, node.URI)
		}
		waitForRestart = waitForRestart || (!skipped && err == nil)
	}
	// No need to wait if all nodes were skipped
	if waitForRestart {
		err = rcc.waitForControlPlaneRestart(restartAfter, kubeAPIServer, kubeControllerManager)
		if err != nil {
			return err
		}
		err = rcc.waitForControlPlaneReadiness()
		if err != nil {
			return err
		}
	}

	log.Debugln("Recreating service account tokens")
	err = rotatecerts.RotateServiceAccountTokens(rcc.kubeClient, rcc.namespaces)
	if err != nil {
		return err
	}
	err = rcc.waitForKubeSystemReadiness()
	if err != nil {
		return err
	}
	return nil
}

func (rcc *rotateCertsCmd) rotateKubelet() error {
	log.Info("Rotating apiserver-kubelet PKI")
	var err error
	var skipped bool
	step := "apiserver_kubelet"

	log.Debugln("Rotating control plane apiserver-kubelet PKI")
	waitForRestart, restartAfter := false, time.Now()
	for _, node := range rcc.nodes {
		skipped, err = execStepsSequence(isMaster, node, execRemoteFunc(remoteBashScript(step)))
		if !skipped && err != nil {
			return errors.Wrapf(err, "executing %s function on remote host %s", step, node.URI)
		}
		waitForRestart = waitForRestart || (!skipped && err == nil)
	}

	// No need to wait if all nodes were skipped
	if waitForRestart {
		err = rcc.waitForControlPlaneRestart(restartAfter, kubeAPIServer, kubeControllerManager)
		if err != nil {
			return err
		}
	}

	log.Debugln("Restarting kube-proxy on control plane nodes")
	for _, node := range rcc.nodes {
		skipped, err = execStepsSequence(isMaster, node, deletePodFunc(rcc.kubeClient, kubeProxyLabels))
		if !skipped && err != nil {
			return errors.Wrapf(err, "executing %s function on remote host %s", step, node.URI)
		}
	}

	log.Debugln("Rotating agents apiserver-kubelet PKI and restarting kube-proxy")
	for _, node := range rcc.nodes {
		skipped, err = execStepsSequence(isLinuxAgent, node, execRemoteFunc(remoteBashScript(step)), deletePodFunc(rcc.kubeClient, kubeProxyLabels))
		if !skipped && err != nil {
			return errors.Wrapf(err, "executing %s function on remote host %s", step, node.URI)
		}
		skipped, err = execStepsSequence(isWindows, node, execRemoteFunc(remotePowershellScript()))
		if !skipped && err != nil {
			return errors.Wrapf(err, "executing %s function on remote host %s", step, node.URI)
		}
	}
	return nil
}

func (rcc *rotateCertsCmd) cleanupRemote() (err error) {
	log.Infoln("Deleting temporary artifacts from cluster nodes")
	step := "cleanup"
	for _, node := range rcc.nodes {
		skipped, err := execStepsSequence(isLinux, node, execRemoteFunc(remoteBashScript(step)))
		if !skipped && err != nil {
			return errors.Wrapf(err, "executing %s function on remote host %s", step, node.URI)
		}
	}
	return nil
}

func (rcc *rotateCertsCmd) updateAPIModel() error {
	log.Infof("Generating new artifacts")
	err := writeArtifacts(filepath.Dir(rcc.apiModelPath), rcc.cs, rcc.apiVersion, rcc.loader.Translator)
	if err != nil {
		return errors.Wrap(err, "writing artifacts")
	}
	err = os.RemoveAll(rcc.outputDirectory)
	if err != nil {
		return errors.Wrap(err, "deleting output directory")
	}
	return nil
}

func execStepsSequence(nodeCondition func(*ssh.RemoteHost) bool, node *ssh.RemoteHost, steps ...func(node *ssh.RemoteHost) error) (skipped bool, err error) {
	if nodeCondition(node) {
		for _, step := range steps {
			if err := step(node); err != nil {
				if ignorable(err) {
					log.Debugf("Node %s skipped", node.URI)
					return true, err
				}
				// remote command failure, should be handled
				return false, err
			}
		}
		// remote command success
		return false, nil
	}
	// node condition not met
	return true, nil
}

func execRemoteFunc(script string) func(node *ssh.RemoteHost) error {
	return func(node *ssh.RemoteHost) error {
		out, err := ssh.ExecuteRemote(node, script)
		if err != nil && !ignorable(err) {
			log.Debugf("Remote command output: %s", out)
		}
		return err
	}
}

func deletePodFunc(client rotatecerts.Client, labels string) func(node *ssh.RemoteHost) error {
	return func(node *ssh.RemoteHost) error {
		err := client.DeletePods(metav1.NamespaceSystem, metav1.ListOptions{
			FieldSelector: fmt.Sprintf("spec.nodeName=%s", node.URI),
			LabelSelector: labels,
		})
		if err != nil {
			return errors.Wrapf(err, "deleting pod with labels %s from node %s", labels, node.URI)
		}
		return nil
	}
}

// waitForControlPlaneReadiness checks that the control plane components are in a healthy state before we move to the next step.
func (rcc *rotateCertsCmd) waitForControlPlaneReadiness() error {
	log.Info("Checking health of control plane components")
	pods := make([]string, 0)
	components := []string{kubeAddonManager, kubeAPIServer, kubeControllerManager, kubeScheduler}
	for _, n := range rcc.nodes {
		if isMaster(n) {
			for _, c := range components {
				pods = append(pods, fmt.Sprintf("%s-%s", c, n.URI))
			}
		}
	}
	if err := rotatecerts.WaitForReady(rcc.kubeClient, metav1.NamespaceSystem, pods, 5, rotateCertsDefaultInterval, rotateCertsDefaultTimeout); err != nil {
		return errors.Wrap(err, "waiting for kube-system containers to reach the Ready state within the timeout period")
	}
	return nil
}

// waitForKubeSystemReadiness checks that all kube-system pods are in a healthy state before we move to the next step.
func (rcc *rotateCertsCmd) waitForKubeSystemReadiness() error {
	log.Info("Checking health of all kube-system pods")
	if err := rotatecerts.WaitForAllInNamespaceReady(rcc.kubeClient, metav1.NamespaceSystem, 5, rotateCertsDefaultInterval, rotateCertsDefaultTimeout); err != nil {
		return errors.Wrap(err, "waiting for kube-system containers to reach the Ready state within the timeout period")
	}
	return nil
}

// waitForControlPlaneRestart checks that each of the passed pods (1) restarted after restart time (2) and is healthy after the restart.
//
// Updating the yaml manifest of the control plane components triggers a pod restart (sometimes after a few seconds).
// This method makes sure that the restart actually happened before further interaction with these pods.
func (rcc *rotateCertsCmd) waitForControlPlaneRestart(restartAfter time.Time, components ...string) error {
	log.Infof("Waiting for control plane components restart: %s", components)
	pods := []string{}
	for _, n := range rcc.nodes {
		if isMaster(n) {
			for _, c := range components {
				pods = append(pods, fmt.Sprintf("%s-%s", c, n.URI))
			}
		}
	}
	if err := rotatecerts.WaitForRestart(rcc.kubeClient, metav1.NamespaceSystem, pods, restartAfter, rotateCertsDefaultInterval, rotateCertsDefaultTimeout); err != nil {
		return errors.Wrap(err, "waiting for control plane components restart")
	}
	return nil
}

func (rcc *rotateCertsCmd) getKubeClient() (rotatecerts.Client, error) {
	config := path.Join("kubeconfig", fmt.Sprintf("kubeconfig.%s.json", rcc.location))

	oldConfig := path.Join(rcc.backupDirectory, config)
	content, err := ioutil.ReadFile(oldConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "reading %s", oldConfig)
	}
	oldCAClient, err := kubernetes.NewClient("", string(content), rotateCertsDefaultInterval, rotateCertsDefaultTimeout)
	if err != nil {
		return nil, err
	}

	newConfig := path.Join(rcc.outputDirectory, config)
	content, err = ioutil.ReadFile(newConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "reading %s", newConfig)
	}
	newCAClient, err := kubernetes.NewClient("", string(content), rotateCertsDefaultInterval, rotateCertsDefaultTimeout)
	if err != nil {
		return nil, err
	}

	return kubernetes.NewCompositeClient(oldCAClient, newCAClient, rotateCertsDefaultInterval, rotateCertsDefaultTimeout), nil
}

func getFilesToDistribute(cs *api.ContainerService) (masterFiles map[string]*ssh.RemoteFile, linuxFiles map[string]*ssh.RemoteFile, windowsFiles map[string]*ssh.RemoteFile) {
	p := cs.Properties.CertificateProfile

	linuxAdmin := fmt.Sprintf("%s:%s", cs.Properties.LinuxProfile.AdminUsername, cs.Properties.LinuxProfile.AdminUsername)
	kubeconfig, err := engine.GenerateKubeConfig(cs.Properties, cs.Location)
	if err != nil {
		return nil, nil, nil
	}

	dir := "/tmp/akse/certs"
	masterFiles = map[string]*ssh.RemoteFile{
		"apiserver.crt":     ssh.NewRemoteFile(path.Join(dir, "apiserver.crt"), crtPermissions, rootUserGroup, []byte(p.APIServerCertificate)),
		"apiserver.key":     ssh.NewRemoteFile(path.Join(dir, "apiserver.key"), keyPermissions, rootUserGroup, []byte(p.APIServerPrivateKey)),
		"ca.crt":            ssh.NewRemoteFile(path.Join(dir, "ca.crt"), crtPermissions, rootUserGroup, []byte(p.CaCertificate)),
		"ca.key":            ssh.NewRemoteFile(path.Join(dir, "ca.key"), keyPermissions, rootUserGroup, []byte(p.CaPrivateKey)),
		"client.crt":        ssh.NewRemoteFile(path.Join(dir, "client.crt"), crtPermissions, rootUserGroup, []byte(p.ClientCertificate)),
		"client.key":        ssh.NewRemoteFile(path.Join(dir, "client.key"), keyPermissions, rootUserGroup, []byte(p.ClientPrivateKey)),
		"etcdclient.crt":    ssh.NewRemoteFile(path.Join(dir, "etcdclient.crt"), crtPermissions, rootUserGroup, []byte(p.EtcdClientCertificate)),
		"etcdclient.key":    ssh.NewRemoteFile(path.Join(dir, "etcdclient.key"), keyPermissions, rootUserGroup, []byte(p.EtcdClientPrivateKey)),
		"etcdserver.crt":    ssh.NewRemoteFile(path.Join(dir, "etcdserver.crt"), crtPermissions, rootUserGroup, []byte(p.EtcdServerCertificate)),
		"etcdserver.key":    ssh.NewRemoteFile(path.Join(dir, "etcdserver.key"), keyPermissions, etcdUserGroup, []byte(p.EtcdServerPrivateKey)),
		"kubectlClient.crt": ssh.NewRemoteFile(path.Join(dir, "kubectlClient.crt"), crtPermissions, rootUserGroup, []byte(p.KubeConfigCertificate)),
		"kubectlClient.key": ssh.NewRemoteFile(path.Join(dir, "kubectlClient.key"), keyPermissions, rootUserGroup, []byte(p.KubeConfigPrivateKey)),
		"kubeconfig":        ssh.NewRemoteFile(path.Join(dir, "kubeconfig"), configPermissions, linuxAdmin, []byte(kubeconfig)),
	}
	for i := 0; i < cs.Properties.MasterProfile.Count; i++ {
		crt := fmt.Sprintf("etcdpeer%d.crt", i)
		masterFiles[crt] = ssh.NewRemoteFile(path.Join(dir, crt), crtPermissions, etcdUserGroup, []byte(p.EtcdPeerCertificates[i]))
		key := fmt.Sprintf("etcdpeer%d.key", i)
		masterFiles[key] = ssh.NewRemoteFile(path.Join(dir, key), keyPermissions, etcdUserGroup, []byte(p.EtcdPeerPrivateKeys[i]))
	}
	linuxFiles = map[string]*ssh.RemoteFile{
		"ca.crt":     masterFiles["ca.crt"],
		"client.crt": masterFiles["client.crt"],
		"client.key": masterFiles["client.key"],
	}
	windowsFiles = map[string]*ssh.RemoteFile{
		"ca.crt":     ssh.NewRemoteFile(fmt.Sprintf("$env:temp\\%s", "ca.crt"), "", "", []byte(p.CaCertificate)),
		"client.crt": ssh.NewRemoteFile(fmt.Sprintf("$env:temp\\%s", "client.crt"), "", "", []byte(p.ClientCertificate)),
		"client.key": ssh.NewRemoteFile(fmt.Sprintf("$env:temp\\%s", "client.key"), "", "", []byte(p.ClientPrivateKey)),
	}
	return
}

func remoteBashScript(step string) string {
	return fmt.Sprintf("bash -euxo pipefail -c \"sudo /tmp/akse/rotate-certs.sh %s |& sudo tee -a /var/log/azure/rotate-certs.log\"", step)
}

func remotePowershellScript() string {
	filePath := "$env:temp\\rotate-certs.ps1"
	return fmt.Sprintf("powershell -noprofile -command \"cd c:\\k\\; iex %s | Out-File -Append -Encoding utf8 rotate-certs.log \"", filePath)
}

func isMaster(node *ssh.RemoteHost) bool {
	return strings.HasPrefix(node.URI, common.LegacyControlPlaneVMPrefix)
}
func isLinux(node *ssh.RemoteHost) bool      { return node.OperatingSystem == api.Linux }
func isLinuxAgent(node *ssh.RemoteHost) bool { return isLinux(node) && !isMaster(node) }
func isWindows(node *ssh.RemoteHost) bool    { return node.OperatingSystem == api.Windows }

// ignoreable returns true if the remote script execution was skipped.
func ignorable(err error) bool {
	// remote scripts return SKIP_EXIT_CODE if the steps was completed by a previous rotate-certs execution
	skipExitCode := 25
	e := errors.Cause(err)
	switch v := e.(type) {
	case *sshx.ExitError:
		return v.Waitmsg.ExitStatus() == skipExitCode
	default:
		return false
	}
}

func (rcc *rotateCertsCmd) getNamespacesWithSATokensToRotate() []string {
	// TODO parametize addons namespace so hard-coding their names is not required.
	// TODO maybe add an extra cli param so user can add extra namespaces
	namespaces := []string{metav1.NamespaceSystem}
	if rcc.cs.Properties.OrchestratorProfile.KubernetesConfig.IsAddonEnabled(common.DashboardAddonName) {
		namespaces = append(namespaces, "kubernetes-dashboard")
	}
	if rcc.cs.Properties.OrchestratorProfile.KubernetesConfig.IsAddonEnabled(common.AzureArcOnboardingAddonName) {
		namespaces = append(namespaces, "azure-arc-onboarding")
	}
	if rcc.cs.Properties.OrchestratorProfile.KubernetesConfig.IsAddonEnabled(common.AzurePolicyAddonName) {
		namespaces = append(namespaces, "gatekeeper-system")
	}
	if rcc.cs.Properties.OrchestratorProfile.KubernetesConfig.IsAddonEnabled(common.ScheduledMaintenanceAddonName) {
		namespaces = append(namespaces, "drainsafe-system")
	}
	return namespaces
}

// // logDowntime is quick and dirty way of checking if the cert rotation process causes downtime.
// func logDowntime(ctx context.Context, rcc *rotateCertsCmd) {
// 	log.Debug("Checking downtime in the background")
// 	config := path.Join("kubeconfig", fmt.Sprintf("kubeconfig.%s.json", rcc.location))

// 	c, err := ioutil.ReadFile(path.Join(filepath.Dir(rcc.apiModelPath), config))
// 	if err != nil {
// 		log.Debug("Error checking downtime: reading old --kubeconfig")
// 		return
// 	}
// 	oldClient, err := kubernetes.NewClient("", string(c), rotateCertsDefaultInterval, rotateCertsDefaultTimeout)
// 	if err != nil {
// 		log.Debug("Error checking downtime: creating old --kubeconfig")
// 		return
// 	}

// 	c, err = ioutil.ReadFile(path.Join(rcc.outputDirectory, config))
// 	if err != nil {
// 		log.Debug("Error checking downtime: reading new --kubeconfig")
// 		return
// 	}
// 	newClient, err := kubernetes.NewClient("", string(c), rotateCertsDefaultInterval, rotateCertsDefaultTimeout)
// 	if err != nil {
// 		log.Debug("Error checking downtime: creating new --kubeconfig")
// 		return
// 	}

// 	tick := time.Tick(5 * time.Second)
// 	for {
// 		select {
// 		case t := <-tick:
// 			_, oldClientError := oldClient.ListNodes()
// 			_, newClientError := newClient.ListNodes()
// 			if oldClientError != nil && newClientError != nil {
// 				log.Debugf("Downtime: %s", t)
// 			}
// 		case <-ctx.Done():
// 			return
// 		}
// 	}
// }
