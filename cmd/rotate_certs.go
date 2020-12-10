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

	ops "github.com/Azure/aks-engine/cmd/rotatecerts"
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
	rotateCertsShortDescription = "Rotate certificates on an existing AKS Engine-created Kubernetes cluster"
	rotateCertsLongDescription  = "Rotate CA, etcd, kubelet, kubeconfig and apiserver certificates in a cluster built with AKS Engine. Rotating certificates can break component connectivity and leave the cluster in an unrecoverable state. Before performing any of these instructions on a live cluster, it is preferrable to backup your cluster state and migrate critical workloads to another cluster."
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
	rotateCertsDefaultTimeout  = 15 * time.Minute
)

type nodeMap = map[string]*ssh.RemoteHost
type fileMap = map[string]*ssh.RemoteFile

type rotateCertsCmd struct {
	// user input
	location               string
	apiModelPath           string
	newCertsPath           string
	sshHostURI             string
	linuxSSHPrivateKeyPath string
	outputDirectory        string
	resumeAfterError       bool
	// computed
	backupDirectory   string
	apiVersion        string
	cs                *api.ContainerService
	loader            *api.Apiloader
	newCertsProfile   *api.CertificateProfile
	kubeClient        *kubernetes.CompositeClientSet
	saTokenNamespaces []string
	nodes             nodeMap
	generateCerts     bool
	linuxAuthConfig   *ssh.AuthConfig
	windowsAuthConfig *ssh.AuthConfig
	jumpbox           *ssh.JumpBox
}

func newRotateCertsCmd() *cobra.Command {
	rcc := rotateCertsCmd{
		generateCerts: true,
	}
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
	command.Flags().BoolVarP(&rcc.resumeAfterError, "resume", "", false, "resume a previous execution that did not complete successfully")
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
		rcc.generateCerts = false
		if _, err = os.Stat(rcc.newCertsPath); os.IsNotExist(err) {
			return errors.Errorf("specified --certificate-profile does not exist (%s)", rcc.newCertsPath)
		}
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
	if len(fs) > 0 && !rcc.resumeAfterError {
		return errors.Errorf("output directory %s is not empty, this indicates that a previous rotate-certs execution did not complete successfully, set --resume to continue the process", rcc.outputDirectory)
	}
	if len(fs) == 0 && rcc.resumeAfterError {
		return errors.Errorf("output directory %s is empty, --resume should be unset", rcc.outputDirectory)
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
	if rcc.cs.Properties.IsCustomCloudProfile() {
		if err = writeCustomCloudProfile(rcc.cs); err != nil {
			return errors.Wrap(err, "error writing custom cloud profile")
		}
		if err = rcc.cs.Properties.SetCustomCloudSpec(api.AzureCustomCloudSpecParams{IsUpgrade: false, IsScale: true}); err != nil {
			return errors.Wrap(err, "error parsing the api model")
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
	rcc.saTokenNamespaces = rcc.getNamespacesWithSATokensToRotate()
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
	if err = rcc.backupCerts(); err != nil {
		return errors.Wrap(err, "backing up current state")
	}
	if err = rcc.updateCertificateProfile(); err != nil {
		return errors.Wrap(err, "updating certificate profile")
	}
	rcc.kubeClient, err = rcc.getKubeClient()
	if err != nil {
		return errors.Wrap(err, "creating Kubernetes client")
	}
	resumeClusterAutoscaler, err := ops.PauseClusterAutoscaler(rcc.kubeClient)
	defer func() {
		if e := resumeClusterAutoscaler(); e != nil {
			log.Warn(e)
		}
	}()
	if err != nil {
		return err
	}

	rcc.nodes, err = rcc.getClusterNodes()
	if err != nil {
		return errors.Wrap(err, "listing cluster nodes")
	}
	if err = rcc.waitForNodesReady(rcc.nodes); err != nil {
		return err
	}
	if err = rcc.waitForControlPlaneReadiness(); err != nil {
		return err
	}
	if err = rcc.rotateMasterCerts(); err != nil {
		return errors.Wrap(err, "rotating certificates")
	}

	rcc.nodes, err = rcc.getClusterNodes()
	if err != nil {
		return errors.Wrap(err, "listing cluster nodes")
	}
	if err = rcc.rotateAgentCerts(); err != nil {
		return errors.Wrap(err, "rotating certificates")
	}
	if err := rcc.updateAPIModel(); err != nil {
		return errors.Wrap(err, "updating apimodel")
	}
	if err := rcc.waitForNodesReady(rcc.nodes); err != nil {
		return err
	}
	if err := rcc.waitForControlPlaneReadiness(); err != nil {
		return err
	}

	log.Infoln("Certificate rotation completed")
	return nil
}

func (rcc *rotateCertsCmd) backupCerts() error {
	log.Infof("Backing up artifacts to directory %s", rcc.backupDirectory)
	if err := writeArtifacts(rcc.backupDirectory, rcc.cs, rcc.apiVersion, rcc.loader.Translator); err != nil {
		return errors.Wrap(err, "writing artifacts")
	}
	return nil
}

func (rcc *rotateCertsCmd) updateCertificateProfile() error {
	if rcc.resumeAfterError {
		if err := rcc.loadCertificateProfile(rcc.outputDirectory); err != nil {
			return errors.Wrap(err, "loading artifacts")
		}
	} else {
		if rcc.generateCerts {
			if err := rcc.generateTLSArtifacts(); err != nil {
				return errors.Wrap(err, "generating artifacts")
			}
		} else {
			rcc.cs.Properties.CertificateProfile = rcc.newCertsProfile
		}
		log.Infof("Writing artifacts to output directory %s", rcc.outputDirectory)
		if err := writeArtifacts(rcc.outputDirectory, rcc.cs, rcc.apiVersion, rcc.loader.Translator); err != nil {
			return errors.Wrap(err, "writing artifacts")
		}
	}
	return nil
}

func (rcc *rotateCertsCmd) loadCertificateProfile(dir string) error {
	log.Infof("Loading certificates from output directory: %s", dir)
	cs, _, err := rcc.loader.LoadContainerServiceFromFile(path.Join(dir, "apimodel.json"), false, false, nil)
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

// getClusterNodes returns all cluster nodes
func (rcc *rotateCertsCmd) getClusterNodes() (nodeMap, error) {
	// make sure we always include control plane nodes
	nodes := make(nodeMap)
	for _, master := range rcc.cs.Properties.GetMasterVMNameList() {
		nodes[master] = &ssh.RemoteHost{
			URI:             master,
			Port:            22,
			OperatingSystem: api.Linux,
			AuthConfig:      rcc.linuxAuthConfig,
			Jumpbox:         rcc.jumpbox,
		}
	}
	nodeList, err := rcc.kubeClient.ListNodes()
	if err != nil {
		return nil, err
	}
	for _, nli := range nodeList.Items {
		node := &ssh.RemoteHost{
			URI:     nli.Name,
			Port:    22,
			Jumpbox: rcc.jumpbox,
		}
		switch api.OSType(strings.Title(nli.Status.NodeInfo.OperatingSystem)) {
		case api.Linux:
			node.OperatingSystem = api.Linux
			node.AuthConfig = rcc.linuxAuthConfig
		case api.Windows:
			node.OperatingSystem = api.Windows
			node.AuthConfig = rcc.windowsAuthConfig
		default:
			return nil, errors.Errorf("listing nodes, could not determine operating system of node %s", nli.Name)
		}
		nodes[node.URI] = node
	}
	return nodes, nil
}

// distributeCerts copies the new set of certificates to the cluster nodes.
func (rcc *rotateCertsCmd) distributeCerts(nodes nodeMap) error {
	log.Info("Distributing certificates")
	upload := func(files fileMap, node *ssh.RemoteHost) error {
		for _, file := range files {
			co, err := ssh.CopyToRemote(node, file)
			if err != nil {
				log.Debugf("Remote command output: %s", co)
				return errors.Wrap(err, "uploading certificate")
			}
		}
		return nil
	}
	areCertsDistributedScript := func(node *ssh.RemoteHost) string {
		switch node.OperatingSystem {
		case api.Linux:
			return "bash -euxo pipefail -c \"[ -d /etc/kubernetes/certs.bak ] && exit 25 || exit 0\""
		case api.Windows:
			filePath := "$env:temp\\ca.crt"
			return fmt.Sprintf("powershell -noprofile -command \"if (Test-Path %s) { exit 25 } else { exit 0 }\"", filePath)
		default:
			return ""
		}
	}
	masterCerts, linuxCerts, windowsCerts, e := getFilesToDistribute(rcc.cs, "/etc/kubernetes/rotate-certs/certs")
	if e != nil {
		return errors.Wrap(e, "collectiong files to distribute")
	}
	for _, node := range nodes {
		skip, err := execStepsSequence(allNodes, node, execRemoteFunc(areCertsDistributedScript(node)))
		if err != nil {
			if skip {
				continue
			}
			return errors.Wrapf(err, "checking certs on remote host %s", node.URI)
		}
		log.Debugf("Uploading certificates to node %s", node.URI)
		if isMaster(node) {
			err = upload(masterCerts, node)
		} else if isLinuxAgent(node) {
			err = upload(linuxCerts, node)
		} else if isWindowsAgent(node) {
			err = upload(windowsCerts, node)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (rcc *rotateCertsCmd) rotateMasterCerts() (err error) {
	nodes := make(nodeMap)
	for k, v := range rcc.nodes {
		if isMaster(v) {
			nodes[k] = v
		}
	}
	rcc.nodes = nodes
	if err = rcc.cleanupRemote(nodes, false); err != nil {
		return errors.Wrap(err, "deleting temporary artifacts from cluster nodes")
	}
	if err = rcc.distributeCerts(nodes); err != nil {
		return errors.Wrap(err, "distributing certificates")
	}
	log.Info("Rotating control plane certificates")
	if err := rcc.backupRemote(nodes); err != nil {
		return err
	}
	if err := rcc.rotateEtcd(nodes); err != nil {
		return err
	}
	if err := rcc.rotateSATokens(nodes); err != nil {
		return err
	}
	if err := rcc.rotateMasterKubelet(nodes); err != nil {
		return err
	}
	if err := rcc.cleanupRemote(nodes, true); err != nil {
		return err
	}
	return nil
}

func (rcc *rotateCertsCmd) rotateAgentCerts() (err error) {
	nodes := make(nodeMap)
	for k, v := range rcc.nodes {
		if isAgent(v) {
			nodes[k] = v
		}
	}
	rcc.nodes = nodes
	if err = rcc.cleanupRemote(nodes, false); err != nil {
		return errors.Wrap(err, "deleting temporary artifacts from cluster nodes")
	}
	if err = rcc.distributeCerts(nodes); err != nil {
		return errors.Wrap(err, "distributing certificates")
	}
	log.Info("Rotating agent certificates")
	if err := rcc.backupRemote(nodes); err != nil {
		return err
	}
	if err := rcc.rotateAgentKubelet(nodes); err != nil {
		return err
	}
	if err := rcc.cleanupRemote(nodes, true); err != nil {
		return err
	}
	return nil
}

func (rcc *rotateCertsCmd) backupRemote(nodes nodeMap) (err error) {
	log.Info("Backing up node certificates")
	step := "backup"
	for _, node := range nodes {
		skipped, err := execStepsSequence(isLinux, node, execRemoteFunc(remoteBashScript(step)))
		if !skipped && err != nil {
			return errors.Wrapf(err, "executing %s function on remote host %s", step, node.URI)
		}
		skipped, err = execStepsSequence(isWindowsAgent, node, execRemoteFunc(remotePowershellScript("Backup")))
		if !skipped && err != nil {
			return errors.Wrapf(err, "executing %s function on remote host %s", step, node.URI)
		}
	}
	return nil
}

func (rcc *rotateCertsCmd) rotateEtcd(nodes nodeMap) error {
	log.Infoln("Rotating etcd PKI")
	log.Debugln("Replacing etcd ca with a ca bundle")
	waitForReadiness, restartAfter := make(nodeMap), time.Now()
	for _, node := range nodes {
		step := "etcd_cabundle"
		skipped, err := execStepsSequence(isMaster, node, execRemoteFunc(remoteBashScript(step)))
		if !skipped && err != nil {
			return errors.Wrapf(err, "executing %s function on remote host %s", step, node.URI)
		}
		if shouldWaitForReadiness(skipped, err) {
			waitForReadiness[node.URI] = node
		}
	}
	// No need to wait if all nodes were skipped
	if len(waitForReadiness) > 0 {
		if err := rcc.waitForControlPlaneRestart(waitForReadiness, restartAfter, kubeAPIServer); err != nil {
			return err
		}
		if err := rcc.waitForControlPlaneReadiness(); err != nil {
			return err
		}
	}

	log.Debugln("Rotating etcd non-ca certs")
	for _, node := range nodes {
		step := "etcd_certs"
		skipped, err := execStepsSequence(isMaster, node, execRemoteFunc(remoteBashScript(step)))
		if !skipped && err != nil {
			return errors.Wrapf(err, "executing %s function on remote host %s", step, node.URI)
		}
	}

	log.Debugln("Rotating etcd ca certs")
	waitForReadiness, restartAfter = make(nodeMap), time.Now()
	for _, node := range nodes {
		step := "etcd_ca"
		skipped, err := execStepsSequence(isMaster, node, execRemoteFunc(remoteBashScript(step)), deleteMirrorPodFunc(rcc.kubeClient, kubeSchedulerLabels))
		if !skipped && err != nil {
			return errors.Wrapf(err, "executing %s function on remote host %s", step, node.URI)
		}
		if shouldWaitForReadiness(skipped, err) {
			waitForReadiness[node.URI] = node
		}
	}
	// No need to wait if all nodes were skipped
	if len(waitForReadiness) > 0 {
		if err := rcc.waitForControlPlaneRestart(waitForReadiness, restartAfter, kubeAPIServer); err != nil {
			return err
		}
		if err := rcc.waitForControlPlaneReadiness(); err != nil {
			return err
		}
	}
	return nil
}

func (rcc *rotateCertsCmd) rotateSATokens(nodes nodeMap) error {
	log.Infoln("Rotating service account tokens")
	log.Debugln("Rotating service account signer")
	waitForReadiness, restartAfter := make(nodeMap), time.Now()
	for _, node := range nodes {
		step := "sa_token_signer"
		skipped, err := execStepsSequence(isMaster, node, execRemoteFunc(remoteBashScript("sa_token_signer")))
		if !skipped && err != nil {
			return errors.Wrapf(err, "executing %s function on remote host %s", step, node.URI)
		}
		if shouldWaitForReadiness(skipped, err) {
			waitForReadiness[node.URI] = node
		}
	}
	// No need to wait if all nodes were skipped
	if len(waitForReadiness) > 0 {
		if err := rcc.waitForControlPlaneRestart(nodes, restartAfter, kubeAPIServer, kubeControllerManager); err != nil {
			return err
		}
		if err := rcc.waitForControlPlaneReadiness(); err != nil {
			return err
		}
	}
	log.Debugln("Recreating service account tokens")
	if err := ops.RotateServiceAccountTokens(rcc.kubeClient, rcc.saTokenNamespaces); err != nil {
		return err
	}
	if err := rcc.waitForKubeSystemReadiness(); err != nil {
		return err
	}
	return nil
}

func (rcc *rotateCertsCmd) rotateMasterKubelet(nodes nodeMap) error {
	log.Info("Rotating control plane apiserver-kubelet PKI")
	step := "apiserver_kubelet"
	waitForReadiness, restartAfter := make(nodeMap), time.Now()
	for _, node := range nodes {
		skipped, err := execStepsSequence(isMaster, node, execRemoteFunc(remoteBashScript(step)))
		if !skipped && err != nil {
			return errors.Wrapf(err, "executing %s function on remote host %s", step, node.URI)
		}
		if shouldWaitForReadiness(skipped, err) {
			waitForReadiness[node.URI] = node
		}
	}
	// No need to wait if all nodes were skipped
	if len(waitForReadiness) > 0 {
		if err := rcc.waitForControlPlaneRestart(nodes, restartAfter, kubeAPIServer, kubeControllerManager); err != nil {
			return err
		}
	}
	log.Info("Restarting kube-proxy on control plane nodes")
	for _, node := range nodes {
		skipped, err := execStepsSequence(isMaster, node, deletePodFunc(rcc.kubeClient, kubeProxyLabels), deleteMirrorPodFunc(rcc.kubeClient, kubeSchedulerLabels))
		if !skipped && err != nil {
			return errors.Wrapf(err, "executing %s function on remote host %s", step, node.URI)
		}
	}
	return nil
}

func (rcc *rotateCertsCmd) rotateAgentKubelet(nodes nodeMap) error {
	log.Info("Rotating agents apiserver-kubelet PKI and restarting kube-proxy")
	step := "apiserver_kubelet"
	for _, node := range nodes {
		skipped, err := execStepsSequence(isLinuxAgent, node, execRemoteFunc(remoteBashScript(step)), deletePodFunc(rcc.kubeClient, kubeProxyLabels))
		if !skipped && err != nil {
			return errors.Wrapf(err, "executing %s function on remote host %s", step, node.URI)
		}
		skipped, err = execStepsSequence(isWindowsAgent, node, execRemoteFunc(remotePowershellScript("Start-CertRotation")))
		if !skipped && err != nil {
			return errors.Wrapf(err, "executing %s function on remote host %s", step, node.URI)
		}
	}
	return nil
}

func (rcc *rotateCertsCmd) cleanupRemote(nodes nodeMap, force bool) (err error) {
	if !rcc.resumeAfterError || force {
		log.Infoln("Deleting temporary artifacts from cluster nodes")
		step := "cleanup"
		for _, node := range nodes {
			skipped, err := execStepsSequence(isLinux, node, execRemoteFunc(remoteBashScript(step)))
			if !skipped && err != nil {
				return errors.Wrapf(err, "executing %s function on remote host %s", step, node.URI)
			}
			skipped, err = execStepsSequence(isWindowsAgent, node, execRemoteFunc(remotePowershellScript("Clean")))
			if !skipped && err != nil {
				return errors.Wrapf(err, "executing %s function on remote host %s", step, node.URI)
			}
		}
	}
	return nil
}

func (rcc *rotateCertsCmd) updateAPIModel() error {
	log.Infof("Generating new artifacts")
	if err := writeArtifacts(filepath.Dir(rcc.apiModelPath), rcc.cs, rcc.apiVersion, rcc.loader.Translator); err != nil {
		return errors.Wrap(err, "writing artifacts")
	}
	if err := os.RemoveAll(rcc.outputDirectory); err != nil {
		return errors.Wrap(err, "deleting output directory")
	}
	return nil
}

func execStepsSequence(cond nodeCondition, node *ssh.RemoteHost, steps ...func(node *ssh.RemoteHost) error) (skipped bool, err error) {
	if !cond(node) {
		// node condition not met, do not execute
		return true, nil
	}
	for _, step := range steps {
		if err := step(node); err != nil {
			if skippedByNode(err) {
				log.Debugf("Node %s skipped", node.URI)
				return true, err
			}
			// remote command failure, should be handled
			return false, err
		}
	}
	// all remote commands succeeded
	return false, nil
}

func execRemoteFunc(script string) func(node *ssh.RemoteHost) error {
	return func(node *ssh.RemoteHost) error {
		out, err := ssh.ExecuteRemote(node, script)
		if err != nil && !skippedByNode(err) {
			log.Debugf("Remote command output: %s", out)
		}
		return err
	}
}

func deletePodFunc(client *kubernetes.CompositeClientSet, labels string) func(node *ssh.RemoteHost) error {
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

func deleteMirrorPodFunc(client *kubernetes.CompositeClientSet, labels string) func(node *ssh.RemoteHost) error {
	return func(node *ssh.RemoteHost) error {
		if err := ops.RestartContainer(client, node, labels); err != nil {
			return errors.Wrapf(err, "deleting container with labels %s from node %s", labels, node.URI)
		}
		return nil
	}
}

// waitForControlPlaneReadiness checks that the control plane components are in a healthy state before we move to the next step.
func (rcc *rotateCertsCmd) waitForControlPlaneReadiness() error {
	log.Info("Checking health of control plane components")
	pods := make([]string, 0)
	for _, n := range rcc.cs.Properties.GetMasterVMNameList() {
		for _, c := range []string{kubeAddonManager, kubeAPIServer, kubeControllerManager, kubeScheduler} {
			pods = append(pods, fmt.Sprintf("%s-%s", c, n))
		}
	}
	if err := ops.WaitForReady(rcc.kubeClient, metav1.NamespaceSystem, pods, rotateCertsDefaultInterval, rotateCertsDefaultTimeout, rcc.nodes); err != nil {
		return errors.Wrap(err, "waiting for kube-system containers to reach the Ready state within the timeout period")
	}
	return nil
}

// waitForKubeSystemReadiness checks that all kube-system pods are in a healthy state before we move to the next step.
func (rcc *rotateCertsCmd) waitForKubeSystemReadiness() error {
	log.Info("Checking health of all kube-system pods")
	timeout := time.Duration(len(rcc.nodes)) * time.Duration(float64(time.Minute)*1.25)
	if rotateCertsDefaultTimeout > timeout {
		timeout = rotateCertsDefaultTimeout
	}
	if err := ops.WaitForAllInNamespaceReady(rcc.kubeClient, metav1.NamespaceSystem, rotateCertsDefaultInterval, timeout, rcc.nodes); err != nil {
		return errors.Wrap(err, "waiting for kube-system containers to reach the Ready state within the timeout period")
	}
	return nil
}

// waitForControlPlaneRestart checks that each of the passed pods (1) restarted after restart time (2) and is healthy after the restart.
//
// Updating the yaml manifest of the control plane components triggers a pod restart (sometimes after a few seconds).
// This method makes sure that the restart actually happened before further interaction with these pods.
func (rcc *rotateCertsCmd) waitForControlPlaneRestart(nodes nodeMap, restartAfter time.Time, components ...string) error {
	log.Infof("Waiting for control plane components restart: %s", components)
	pods := make([]string, 0)
	for k := range nodes {
		for _, c := range components {
			pods = append(pods, fmt.Sprintf("%s-%s", c, k))
		}
	}
	if err := ops.WaitForRestart(rcc.kubeClient, metav1.NamespaceSystem, pods, restartAfter, rotateCertsDefaultInterval, rotateCertsDefaultTimeout, nodes); err != nil {
		return errors.Wrap(err, "waiting for control plane components restart")
	}
	return nil
}

func (rcc *rotateCertsCmd) waitForNodesReady(nodes nodeMap) error {
	log.Info("Waiting for cluster nodes readiness")
	if err := ops.WaitForNodesReady(rcc.kubeClient, nodes, rotateCertsDefaultInterval, rotateCertsDefaultTimeout); err != nil {
		return errors.Wrap(err, "waiting for cluster nodes readiness")
	}
	return nil
}

func (rcc *rotateCertsCmd) getKubeClient() (*kubernetes.CompositeClientSet, error) {
	configPathSuffix := path.Join("kubeconfig", fmt.Sprintf("kubeconfig.%s.json", rcc.location))

	oldConfigPath := path.Join(rcc.backupDirectory, configPathSuffix)
	oldConfig, err := ioutil.ReadFile(oldConfigPath)
	if err != nil {
		return nil, errors.Wrapf(err, "reading %s", oldConfigPath)
	}
	oldCAClient, err := kubernetes.NewClient("", string(oldConfig), rotateCertsDefaultInterval, rotateCertsDefaultTimeout)
	if err != nil {
		return nil, errors.Wrapf(err, "creating client from %s", oldConfigPath)
	}

	newConfigPath := path.Join(rcc.outputDirectory, configPathSuffix)
	newConfig, err := ioutil.ReadFile(newConfigPath)
	if err != nil {
		return nil, errors.Wrapf(err, "reading %s", newConfigPath)
	}
	newCAClient, err := kubernetes.NewClient("", string(newConfig), rotateCertsDefaultInterval, rotateCertsDefaultTimeout)
	if err != nil {
		return nil, errors.Wrapf(err, "creating client from %s", newConfigPath)
	}

	return kubernetes.NewCompositeClient(oldCAClient, newCAClient, rotateCertsDefaultInterval, rotateCertsDefaultTimeout), nil
}

func getFilesToDistribute(cs *api.ContainerService, dir string) (fileMap, fileMap, fileMap, error) {
	p := cs.Properties.CertificateProfile

	kubeconfig, err := getKubeConfig(cs, dir)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "generating new kubeconfig")
	}
	linuxScript, err := getLinuxScript()
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "loading rotate-certs.sh")
	}
	windowsScript, err := getWindowsScript()
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "loading rotate-certs.ps1")
	}

	masterFiles := fileMap{
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
		"kubeconfig":        kubeconfig,
		"script":            linuxScript,
	}
	for i := 0; i < cs.Properties.MasterProfile.Count; i++ {
		crt := fmt.Sprintf("etcdpeer%d.crt", i)
		masterFiles[crt] = ssh.NewRemoteFile(path.Join(dir, crt), crtPermissions, etcdUserGroup, []byte(p.EtcdPeerCertificates[i]))
		key := fmt.Sprintf("etcdpeer%d.key", i)
		masterFiles[key] = ssh.NewRemoteFile(path.Join(dir, key), keyPermissions, etcdUserGroup, []byte(p.EtcdPeerPrivateKeys[i]))
	}
	linuxFiles := fileMap{
		"ca.crt":     masterFiles["ca.crt"],
		"client.crt": masterFiles["client.crt"],
		"client.key": masterFiles["client.key"],
		"script":     linuxScript,
	}
	windowsFiles := fileMap{
		"ca.crt":     ssh.NewRemoteFile(fmt.Sprintf("$env:temp\\%s", "ca.crt"), "", "", []byte(p.CaCertificate)),
		"client.crt": ssh.NewRemoteFile(fmt.Sprintf("$env:temp\\%s", "client.crt"), "", "", []byte(p.ClientCertificate)),
		"client.key": ssh.NewRemoteFile(fmt.Sprintf("$env:temp\\%s", "client.key"), "", "", []byte(p.ClientPrivateKey)),
		"script":     windowsScript,
	}
	return masterFiles, linuxFiles, windowsFiles, nil
}

func getKubeConfig(cs *api.ContainerService, dir string) (*ssh.RemoteFile, error) {
	adminUsername := fmt.Sprintf("%s:%s", cs.Properties.LinuxProfile.AdminUsername, cs.Properties.LinuxProfile.AdminUsername)
	kubeconfig, err := engine.GenerateKubeConfig(cs.Properties, cs.Location)
	if err != nil {
		return nil, err
	}
	return ssh.NewRemoteFile(path.Join(dir, "kubeconfig"), configPermissions, adminUsername, []byte(kubeconfig)), nil
}

func getLinuxScript() (*ssh.RemoteFile, error) {
	c, err := engine.Asset("k8s/rotate-certs.sh")
	if err != nil {
		return nil, err
	}
	return ssh.NewRemoteFile("/etc/kubernetes/rotate-certs/rotate-certs.sh", "744", rootUserGroup, c), nil
}

func getWindowsScript() (*ssh.RemoteFile, error) {
	c, err := engine.Asset("k8s/rotate-certs.ps1")
	if err != nil {
		return nil, err
	}
	return ssh.NewRemoteFile("$env:temp\\rotate-certs.ps1", "", "", c), nil
}

func remoteBashScript(step string) string {
	return fmt.Sprintf("bash -euxo pipefail -c \"if [ -f /etc/kubernetes/rotate-certs/rotate-certs.sh ]; then sudo /etc/kubernetes/rotate-certs/rotate-certs.sh %s |& sudo tee -a /var/log/azure/rotate-certs.log; fi\"", step)
}

func remotePowershellScript(step string) string {
	filePath := "$env:temp\\rotate-certs.ps1"
	return fmt.Sprintf("powershell -noprofile -command \"cd c:\\k\\; Import-Module %s; iex %s | Out-File -Append -Encoding utf8 rotate-certs.log\"", filePath, step)
}

type nodeCondition func(*ssh.RemoteHost) bool

func isMaster(node *ssh.RemoteHost) bool {
	return strings.HasPrefix(node.URI, common.LegacyControlPlaneVMPrefix)
}
func isAgent(node *ssh.RemoteHost) bool        { return !isMaster(node) }
func isLinux(node *ssh.RemoteHost) bool        { return node.OperatingSystem == api.Linux }
func isWindowsAgent(node *ssh.RemoteHost) bool { return node.OperatingSystem == api.Windows }
func isLinuxAgent(node *ssh.RemoteHost) bool   { return isLinux(node) && !isMaster(node) }
func allNodes(node *ssh.RemoteHost) bool       { return true }

// skippedByNode returns true if the remote script execution was skipped.
func skippedByNode(err error) bool {
	// remote scripts return SKIP_EXIT_CODE if the steps was completed by a previous rotate-certs execution
	if err == nil {
		return false
	}
	skipExitCode := 25
	switch err := err.(type) {
	case *sshx.ExitError:
		return err.Waitmsg.ExitStatus() == skipExitCode
	default:
		cause := errors.Cause(err)
		if err != cause {
			return skippedByNode(cause)
		}
		return false
	}
}

func shouldWaitForReadiness(skipped bool, err error) bool {
	return !skipped && err == nil
}

func (rcc *rotateCertsCmd) getNamespacesWithSATokensToRotate() []string {
	// TODO parametize addons namespace so hard-coding their names is not required.
	// TODO maybe add an extra cli param so user can add extra namespaces
	namespaces := []string{metav1.NamespaceSystem}
	if rcc.cs.Properties.OrchestratorProfile.KubernetesConfig.IsAddonEnabled(common.DashboardAddonName) {
		namespaces = append(namespaces, "kubernetes-dashboard")
	}
	if rcc.cs.Properties.OrchestratorProfile.KubernetesConfig.IsAddonEnabled(common.AzureArcOnboardingAddonName) {
		namespaces = append(namespaces, "azure-arc")
	}
	if rcc.cs.Properties.OrchestratorProfile.KubernetesConfig.IsAddonEnabled(common.AzurePolicyAddonName) {
		namespaces = append(namespaces, "gatekeeper-system")
	}
	if rcc.cs.Properties.OrchestratorProfile.KubernetesConfig.IsAddonEnabled(common.ScheduledMaintenanceAddonName) {
		namespaces = append(namespaces, "drainsafe-system")
	}
	return namespaces
}
