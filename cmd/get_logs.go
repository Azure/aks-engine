// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/engine"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/leonelquinteros/gotext"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	v1 "k8s.io/api/core/v1"
)

const (
	getLogsName             = "get-logs"
	getLogsShortDescription = "Collect logs and current cluster nodes configuration."
	getLogsLongDescription  = "Collect deployment logs, running daemons/services logs and current nodes configuration."
)

type getLogsCmd struct {
	// user input
	location               string
	apiModelPath           string
	sshHostURI             string
	linuxSSHPrivateKeyPath string
	linuxScriptPath        string
	windowsScriptPath      string
	outputDirectory        string
	controlPlaneOnly       bool
	blobServiceSASURL      string
	storageContainerName   string
	// computed
	cs               *api.ContainerService
	locale           *gotext.Locale
	armClient        armhelpers.AKSEngineClient
	masterNodes      []v1.Node
	linuxNodes       []v1.Node
	linuxSSHConfig   *ssh.ClientConfig
	windowsNodes     []v1.Node
	windowsSSHConfig *ssh.ClientConfig
}

func newGetLogsCmd() *cobra.Command {
	glc := getLogsCmd{}
	command := &cobra.Command{
		Use:   getLogsName,
		Short: getLogsShortDescription,
		Long:  getLogsLongDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := glc.validateArgs(); err != nil {
				_ = cmd.Usage()
				return errors.Wrap(err, "validating get-logs args")
			}
			if err := glc.loadAPIModel(); err != nil {
				return errors.Wrap(err, "loading API model")
			}
			return glc.run()
		},
	}
	command.Flags().StringVarP(&glc.location, "location", "l", "", "Azure location where the cluster is deployed (required)")
	command.Flags().StringVarP(&glc.apiModelPath, "api-model", "m", "", "path to the generated apimodel.json file (required)")
	command.Flags().StringVar(&glc.sshHostURI, "ssh-host", "", "FQDN, or IP address, of an SSH listener that can reach all nodes in the cluster (required)")
	command.Flags().StringVar(&glc.linuxSSHPrivateKeyPath, "linux-ssh-private-key", "", "path to a valid private SSH key to access the cluster's Linux nodes (required)")
	command.Flags().StringVar(&glc.linuxScriptPath, "linux-script", "", "path to the log collection script to execute on the cluster's Linux nodes (required if distro is not aks-ubuntu)")
	command.Flags().StringVar(&glc.windowsScriptPath, "windows-script", "", "path to the log collection script to execute on the cluster's Windows nodes (required if distro is not aks-windows)")
	command.Flags().StringVarP(&glc.outputDirectory, "output-directory", "o", "", "collected logs destination directory, derived from --api-model if missing")
	command.Flags().BoolVarP(&glc.controlPlaneOnly, "control-plane-only", "", false, "get logs from control plane VMs only")
	command.Flags().StringVarP(&glc.blobServiceSASURL, "sas-url", "", "", "blob service SAS URL of the storage account on Azure or custom cloud to upload kubernetes logs")
	command.Flags().StringVar(&glc.storageContainerName, "storage-container-name", "", "name of the storage container that to upload kubernetes logs (required if --sas-url is set and storage container name not included in the SAS URL)")
	_ = command.MarkFlagRequired("location")
	_ = command.MarkFlagRequired("api-model")
	_ = command.MarkFlagRequired("ssh-host")
	_ = command.MarkFlagRequired("linux-ssh-private-key")
	return command
}

func (glc *getLogsCmd) validateArgs() (err error) {
	if glc.locale, err = i18n.LoadTranslations(); err != nil {
		return errors.Wrap(err, "loading translation files")
	}
	glc.location = helpers.NormalizeAzureRegion(glc.location)
	if glc.location == "" {
		return errors.New("--location must be specified")
	}
	if glc.sshHostURI == "" {
		return errors.New("--ssh-host must be specified")
	}
	if glc.apiModelPath == "" {
		return errors.New("--api-model must be specified")
	} else if _, err := os.Stat(glc.apiModelPath); os.IsNotExist(err) {
		return errors.Errorf("specified --api-model does not exist (%s)", glc.apiModelPath)
	}
	if glc.linuxSSHPrivateKeyPath == "" {
		return errors.New("--linux-ssh-private-key must be specified")
	} else if _, err := os.Stat(glc.linuxSSHPrivateKeyPath); os.IsNotExist(err) {
		return errors.Errorf("specified --linux-ssh-private-key does not exist (%s)", glc.linuxSSHPrivateKeyPath)
	}
	if glc.linuxScriptPath != "" {
		if _, err := os.Stat(glc.linuxScriptPath); os.IsNotExist(err) {
			return errors.Errorf("specified --linux-script does not exist (%s)", glc.linuxScriptPath)
		}
	}
	if glc.windowsScriptPath != "" {
		if _, err := os.Stat(glc.windowsScriptPath); os.IsNotExist(err) {
			return errors.Errorf("specified --windows-script does not exist (%s)", glc.windowsScriptPath)
		}
	}
	if glc.outputDirectory == "" {
		glc.outputDirectory = path.Join(filepath.Dir(glc.apiModelPath), "_logs")
		if err := os.MkdirAll(glc.outputDirectory, 0755); err != nil {
			return errors.Errorf("error creating output directory (%s)", glc.outputDirectory)
		}
	}
	return nil
}

func (glc *getLogsCmd) loadAPIModel() (err error) {
	apiloader := &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: glc.locale,
		},
	}
	if glc.cs, _, err = apiloader.LoadContainerServiceFromFile(glc.apiModelPath, false, false, nil); err != nil {
		return errors.Wrap(err, "error parsing api-model")
	}
	if glc.cs.Properties.IsCustomCloudProfile() {
		if err = writeCustomCloudProfile(glc.cs); err != nil {
			return errors.Wrap(err, "error writing custom cloud profile")
		}
		if err = glc.cs.Properties.SetCustomCloudSpec(api.AzureCustomCloudSpecParams{IsUpgrade: false, IsScale: true}); err != nil {
			return errors.Wrap(err, "error parsing the api model")
		}
	}

	if glc.cs.Location == "" {
		glc.cs.Location = glc.location
	} else if glc.cs.Location != glc.location {
		return errors.New("--location flag does not match api-model location")
	}

	lauth, err := helpers.PublicKeyAuth(glc.linuxSSHPrivateKeyPath)
	if err != nil {
		return errors.Wrap(err, "creating linux SSH config")
	}
	glc.linuxSSHConfig = helpers.SSHClientConfig(glc.cs.Properties.LinuxProfile.AdminUsername, lauth)

	if glc.cs.Properties.WindowsProfile != nil && glc.cs.Properties.WindowsProfile.GetSSHEnabled() {
		glc.windowsSSHConfig = helpers.SSHClientConfig(
			glc.cs.Properties.WindowsProfile.AdminUsername,
			ssh.Password(glc.cs.Properties.WindowsProfile.AdminPassword))
	}

	var client *armhelpers.AzureClient
	glc.armClient = client
	return nil
}

func (glc *getLogsCmd) run() (err error) {
	if err = glc.getClusterNodes(); err != nil {
		return errors.Wrap(err, "listing cluster nodes")
	}
	if err = glc.validateLogScript(); err != nil {
		return errors.Wrap(err, "validating log collection scripts for nodes")
	}
	for _, n := range glc.masterNodes {
		log.Infof("Processing master node: %s\n", n.Name)
		out, err := glc.collectLogs(n, glc.linuxSSHConfig)
		if err != nil {
			log.Warnf("Remote command output: %s", out)
			log.Warnf("Error: %s", err)
		}
	}
	if glc.controlPlaneOnly {
		return nil
	}
	for _, n := range glc.linuxNodes {
		log.Infof("Processing Linux node: %s\n", n.Name)
		out, err := glc.collectLogs(n, glc.linuxSSHConfig)
		if err != nil {
			log.Warnf("Remote command output: %s", out)
			log.Warnf("Error: %s", err)
		}
	}
	for _, n := range glc.windowsNodes {
		log.Infof("Processing Windows node: %s\n", n.Name)
		out, err := glc.collectLogs(n, glc.windowsSSHConfig)
		if err != nil {
			log.Warnf("Remote command output: %s", out)
			log.Warnf("Error: %s", err)
		}
	}
	log.Infof("Logs downloaded to %s", glc.outputDirectory)
	return nil
}

func (glc *getLogsCmd) getClusterNodes() error {
	kubeconfig, err := engine.GenerateKubeConfig(glc.cs.Properties, glc.location)
	if err != nil {
		return errors.Wrap(err, "generating kubeconfig")
	}
	kubeClient, err := glc.armClient.GetKubernetesClient("", kubeconfig, time.Second*1, time.Duration(60)*time.Minute)
	if err != nil {
		return errors.Wrap(err, "creating Kubernetes client")
	}
	nodeList, err := kubeClient.ListNodes()
	if err != nil {
		log.Warnf("Error getting node list: %s", err)
		log.Infof("Collecting logs from control plane VMs")
		glc.controlPlaneOnly = true
		glc.masterNodes = computeControlPlaneNodes(glc.cs.Properties.MasterProfile.Count, glc.cs.Properties.GetClusterID())
		return nil
	}
	for _, node := range nodeList.Items {
		if isLinuxNode(node) {
			if strings.HasPrefix(node.Name, common.LegacyControlPlaneVMPrefix) {
				glc.masterNodes = append(glc.masterNodes, node)
			} else {
				glc.linuxNodes = append(glc.linuxNodes, node)
			}
		} else if isWindowsNode(node) {
			if glc.windowsSSHConfig != nil {
				glc.windowsNodes = append(glc.windowsNodes, node)
			} else {
				log.Warnf("Skipping node %s, SSH not enabled", node.Name)
			}
		} else {
			log.Warnf("Skipping node %s, could not determine operating system", node.Name)
		}
	}
	return nil
}

func (glc *getLogsCmd) collectLogs(node v1.Node, config *ssh.ClientConfig) (string, error) {
	jumpboxPort := "22"
	client, err := helpers.SSHClient(glc.sshHostURI, jumpboxPort, node.Name, glc.linuxSSHConfig, config)
	if err != nil {
		return "", errors.Wrap(err, "creating SSH client")
	}
	defer client.Close()

	stdout, err := glc.uploadScript(node, client)
	if err != nil {
		return stdout, err
	}
	stdout, err = glc.executeScript(node, client)
	if err != nil {
		return stdout, err
	}
	stdout, err = glc.downloadLogs(node, client)
	if err != nil {
		return stdout, err
	}
	if glc.blobServiceSASURL != "" {
		log.Debug("Will upload logs to storage account")
		err = glc.uploadLogsToStorageContainer(node)
		if err != nil {
			return "", errors.Wrap(err, "uploading logs to storage container")
		}
	}
	return "", nil
}

func (glc *getLogsCmd) uploadScript(node v1.Node, client *ssh.Client) (string, error) {
	var script, cmd string
	if isLinuxNode(node) && glc.linuxScriptPath != "" {
		script = glc.linuxScriptPath
		cmd = "bash -c \"cat /dev/stdin > /tmp/collect-logs.sh\""
	} else if isWindowsNode(node) && glc.windowsScriptPath != "" {
		script = glc.windowsScriptPath
		cmd = "powershell -noprofile -command \"$Input > $env:temp\\collect-windows-logs.ps1\""
	} else {
		return "", nil
	}

	sc, err := ioutil.ReadFile(script)
	if err != nil {
		return "", errors.Wrapf(err, "reading log collection script %s", script)
	}
	session, err := client.NewSession()
	if err != nil {
		return "", errors.Wrap(err, "creating SSH session")
	}
	defer session.Close()

	log.Debugf("Uploading log collection script (%s)\n", script)
	session.Stdin = bytes.NewReader(sc)
	if co, err := session.CombinedOutput(cmd); err != nil {
		return fmt.Sprintf("%s -> %s", node.Name, string(co)), errors.Wrap(err, "uploading log collection script")
	}
	return "", nil
}

func (glc *getLogsCmd) executeScript(node v1.Node, client *ssh.Client) (string, error) {
	log.Debug("Collecting logs\n")
	session, err := client.NewSession()
	if err != nil {
		return "", errors.Wrap(err, "creating SSH session")
	}
	defer session.Close()

	var script, cmd string
	if isLinuxNode(node) {
		if glc.linuxScriptPath != "" {
			script = "/tmp/collect-logs.sh"
			cmd = fmt.Sprintf("bash -c \"sudo chmod +x %s; export AZURE_ENV=%s; sudo -E %s; rm %s\"", script, glc.getCloudName(), script, script)
		} else {
			script = "/opt/azure/containers/collect-logs.sh"
			cmd = fmt.Sprintf("bash -c \"export AZURE_ENV=%s; sudo -E %s\"", glc.getCloudName(), script)
		}
	} else {
		if glc.windowsScriptPath != "" {
			script = "$env:temp\\collect-windows-logs.ps1"
		} else {
			script = "c:\\k\\debug\\collect-windows-logs.ps1"
		}
		cmd = fmt.Sprintf("powershell -command \"iex %s | Where-Object { $_.extension -eq '.zip' } | Copy-Item -Destination $env:temp\\$env:computername.zip\"", script)
	}

	if co, err := session.CombinedOutput(cmd); err != nil {
		return fmt.Sprintf("%s -> %s", node.Name, string(co)), errors.Wrap(err, "collecting logs on remote host")
	}
	return "", nil
}

func (glc *getLogsCmd) downloadLogs(node v1.Node, client *ssh.Client) (string, error) {
	log.Debug("Downloading logs\n")
	session, err := client.NewSession()
	if err != nil {
		return "", errors.Wrap(err, "creating SSH session")
	}
	defer session.Close()

	localFileName := fmt.Sprintf("%s.zip", node.Name)
	localFilePath := path.Join(glc.outputDirectory, localFileName)
	file, err := os.OpenFile(localFilePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return "", errors.Wrap(err, "opening destination file")
	}
	defer file.Close()

	stdout, err := session.StdoutPipe()
	if err != nil {
		return "", errors.Wrap(err, "opening SSH session stdout pipe")
	}

	var cmd string
	if isLinuxNode(node) {
		cmd = "bash -c \"cat /tmp/logs.zip > /dev/stdout\""
	} else {
		cmd = "type %TEMP%"
		cmd = fmt.Sprintf("%s\\%s.zip", cmd, node.Name)
	}

	if err = session.Start(cmd); err != nil {
		return fmt.Sprintf("%s -> %s", node.Name, session.Stderr), errors.Wrap(err, "downloading logs from remote host")
	}
	_, err = io.Copy(file, io.TeeReader(stdout, &DownloadProgressWriter{}))
	if err != nil {
		return "", errors.Wrap(err, "downloading logs")
	}

	fmt.Println("")
	return "", nil
}

func (glc *getLogsCmd) uploadLogsToStorageContainer(node v1.Node) error {
	log.Infof("Uploading logs for %s", node.Name)
	logFileName := fmt.Sprintf("%s.zip", node.Name)
	logFilePath := path.Join(glc.outputDirectory, logFileName)
	logFile, err := os.Open(logFilePath)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("finding zipped log file %s", logFilePath))
	}

	urls := strings.Split(glc.blobServiceSASURL, "/?")
	if len(urls) != 2 {
		return errors.Wrap(err, fmt.Sprintf("validating blob service SAS URL %s", glc.blobServiceSASURL))
	}
	var fullURL string
	if glc.storageContainerName == "" {
		fullURL = fmt.Sprintf("%s/%s?%s", urls[0], node.Name, urls[1])
	} else {
		fullURL = fmt.Sprintf("%s/%s/%s?%s", urls[0], glc.storageContainerName, node.Name, urls[1])
	}
	u, _ := url.Parse(fullURL)
	blobURL := azblob.NewBlobURL(*u, azblob.NewPipeline(azblob.NewAnonymousCredential(), azblob.PipelineOptions{}))
	blockBlobURL := blobURL.ToBlockBlobURL()

	_, err = azblob.UploadFileToBlockBlob(context.Background(), logFile, blockBlobURL, azblob.UploadToBlockBlobOptions{})
	if err != nil {
		return errors.Wrap(err, "uploading log file to storage container blob")
	}
}

func (glc *getLogsCmd) validateLogScript() error {
	if glc.linuxScriptPath == "" && !glc.cs.Properties.MasterProfile.IsVHDDistro() {
		if glc.controlPlaneOnly {
			return errors.Errorf("No log collection script found for control plane nodes")
		}

		log.Warn("Skipping control plane nodes as flag '--linux-script' is not set and the distro in masterProfiles is not aks-ubuntu VHD")
		glc.masterNodes = nil
	}
	for _, profile := range glc.cs.Properties.AgentPoolProfiles {
		if glc.linuxScriptPath == "" && strings.EqualFold(string(profile.OSType), "Linux") && !profile.IsVHDDistro() {
			log.Warnf("Skipping linux agentpool %s as flag '--linux-script' is not set and the distro in agentPoolProfiles is not aks-ubuntu VHD", profile.Name)
			glc.linuxNodes = filterNodesFromPool(glc.linuxNodes, profile.Name)
		}
		if glc.windowsScriptPath == "" && strings.EqualFold(string(profile.OSType), "Windows") && !glc.cs.Properties.WindowsProfile.IsVHDDistro() {
			log.Warnf("Skipping windows agentpool %s as flag '--windows-script' is not set and the distro in windowsProfiles is not aks-windows VHD", profile.Name)
			glc.windowsNodes = nil
		}
	}
	return nil
}

func isLinuxNode(node v1.Node) bool {
	return strings.EqualFold(node.Status.NodeInfo.OperatingSystem, "linux")
}

func isWindowsNode(node v1.Node) bool {
	return strings.EqualFold(node.Status.NodeInfo.OperatingSystem, "windows")
}

func (glc *getLogsCmd) getCloudName() string {
	if glc.cs.Properties.IsAzureStackCloud() {
		return "AzureStackCloud"
	}
	return ""
}

func filterNodesFromPool(nodeList []v1.Node, agentPoolName string) []v1.Node {
	var linuxNodeList []v1.Node
	for _, node := range nodeList {
		if !strings.EqualFold(strings.Split(node.Name, "-")[1], agentPoolName) {
			linuxNodeList = append(linuxNodeList, node)
		}
	}
	return linuxNodeList
}

func computeControlPlaneNodes(nodesCount int, clusterID string) []v1.Node {
	var nodeList []v1.Node
	for i := 0; i < nodesCount; i++ {
		var node v1.Node
		node.Name = fmt.Sprintf("%s-%s-%d", common.LegacyControlPlaneVMPrefix, clusterID, i)
		node.Status.NodeInfo.OperatingSystem = "linux"
		nodeList = append(nodeList, node)
	}
	return nodeList
}

type DownloadProgressWriter struct {
	Total uint64
}

func (wc *DownloadProgressWriter) Write(p []byte) (int, error) {
	// TODO maybe something like DownloadProgressWriter already exists
	n := len(p)
	wc.Total += uint64(n)
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rDownloading... %d bytes complete", wc.Total)
	return n, nil
}
