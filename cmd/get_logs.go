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
	uploadLogs             bool
	storageAccountName     string
	storageAccountKey      string
	storageContainerURL    string
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
	command.Flags().StringVar(&glc.linuxScriptPath, "linux-script", "", "path to the log collection script to execute on the cluster's Linux nodes")
	command.Flags().StringVar(&glc.windowsScriptPath, "windows-script", "", "path to the log collection script to execute on the cluster's Windows nodes")
	command.Flags().StringVarP(&glc.outputDirectory, "output-directory", "o", "", "collected logs destination directory, derived from --api-model if missing")
	command.Flags().BoolVarP(&glc.controlPlaneOnly, "control-plane-only", "", false, "get logs from control plane VMs only")
	command.Flags().BoolVarP(&glc.uploadLogs, "upload-logs", "", false, "upload logs to a Storage Container on Azure or custom cloud")
	command.Flags().StringVar(&glc.storageAccountName, "storage-account-name", "", "storage account name of the storage account that container exist (required if upload-logs is set)")
	command.Flags().StringVar(&glc.storageAccountKey, "storage-account-key", "", "storage account key of the storage account that container exist (required if upload-logs is set)")
	command.Flags().StringVar(&glc.storageContainerURL, "storage-container-url", "", "URL of the storage container that the logs will be uploaded, will create a default container if no URL provided")
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
	if _, err := os.Stat(glc.linuxScriptPath); os.IsNotExist(err) {
		return errors.Errorf("specified --linux-script does not exist (%s)", glc.linuxScriptPath)
	}
	if _, err := os.Stat(glc.windowsScriptPath); os.IsNotExist(err) {
		return errors.Errorf("specified --windows-script does not exist (%s)", glc.windowsScriptPath)
	}
	if glc.outputDirectory == "" {
		glc.outputDirectory = path.Join(filepath.Dir(glc.apiModelPath), "_logs")
		if err := os.MkdirAll(glc.outputDirectory, 0755); err != nil {
			return errors.Errorf("error creating output directory (%s)", glc.outputDirectory)
		}
	}
	if glc.uploadLogs {
		if glc.storageAccountName == "" {
			return errors.New("--storage-account-name must be specified when --upload is set")
		} else if glc.storageAccountKey == "" {
			return errors.New("--storage-account-key must be specified when --upload is set")
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
	if glc.uploadLogs && glc.storageContainerURL == "" {
		log.Infof("No storage container URL provided, will create a defaul storage container 'kuberneteslogs'")
		err = glc.CreateDefaultStorageContainer()
		if err != nil {
			log.Warnf("Failed to create default storage container 'kuberneteslogs', will not upload logs to storage container")
			log.Warnf("Error: %s", err)
			glc.uploadLogs = false
		}
	} else {
		log.Infof("Will upload logs to storage container URL: %s", glc.storageContainerURL)
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
		log.Warnf("unable to list nodes from api server, will only collect logs from control panel VMs")
		glc.controlPlaneOnly = true
		for nodeIndex := 0; nodeIndex < glc.cs.Properties.MasterProfile.Count; nodeIndex++ {
			var controlPanelNode v1.Node
			controlPanelNode.Name = fmt.Sprint(common.LegacyControlPlaneVMPrefix, "-", glc.cs.Properties.GetClusterID(), "-", nodeIndex)
			controlPanelNode.Status.NodeInfo.OperatingSystem = "linux"
			glc.masterNodes = append(glc.masterNodes, controlPanelNode)
		}
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
				log.Warnf("skipping node %s, SSH not enabled", node.Name)
			}
		} else {
			log.Warnf("skipping node %s, could not determine operating system", node.Name)
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

	if glc.uploadLogs {
		log.Debugf("Will upload logs to storage container, URL: (%s)\n", glc.storageContainerURL)
		err = glc.uploadLogsToStorageContainer(node)
		if err != nil {
			return "", errors.Wrap(err, "uploading logs to storage container")
		}
	}

	return "", nil
}

func (glc *getLogsCmd) uploadScript(node v1.Node, client *ssh.Client) (string, error) {
	if (glc.windowsScriptPath == "") && (glc.linuxScriptPath == "") {
		return "", nil
	}

	if isLinuxNode(node) && (glc.linuxScriptPath != "") {
		linuxScriptContent, err := ioutil.ReadFile(glc.linuxScriptPath)
		if err != nil {
			return "", errors.Wrap(err, "reading Linux log collection script content")
		}

		log.Debugf("Uploading Linux log collection script (%s)\n", glc.linuxScriptPath)
		session, err := client.NewSession()
		if err != nil {
			return "", errors.Wrap(err, "creating SSH session")
		}
		defer session.Close()

		session.Stdin = bytes.NewReader(linuxScriptContent)
		if co, err := session.CombinedOutput("bash -c \"cat /dev/stdin > /tmp/collect-logs.sh\""); err != nil {
			return fmt.Sprintf("%s -> %s", node.Name, string(co)), errors.Wrap(err, "uploading Linux log collection script")
		}
	}

	if isWindowsNode(node) && (glc.windowsScriptPath != "") {
		windowsScriptContent, err := ioutil.ReadFile(glc.windowsScriptPath)
		if err != nil {
			return "", errors.Wrap(err, "reading Windows log collection script content")
		}

		log.Debugf("Uploading Windows log collection script (%s)\n", glc.windowsScriptPath)
		session, err := client.NewSession()
		if err != nil {
			return "", errors.Wrap(err, "creating SSH session")
		}
		defer session.Close()

		session.Stdin = bytes.NewReader(windowsScriptContent)
		if co, err := session.CombinedOutput("powershell -noprofile -command \"$Input > $env:temp\\collect-windows-logs.ps1\""); err != nil {
			return fmt.Sprintf("%s -> %s", node.Name, string(co)), errors.Wrap(err, "uploading Windows log collection script")
		}
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
			script = "(gi $env:temp).fullname + '\\collect-windows-logs.ps1'"
		} else {
			script = "c:\\k\\debug\\collect-windows-logs.ps1"
		}
		cmd = fmt.Sprintf("powershell -command \"%s | Where-Object { $_.extension -eq '.zip' } | Copy-Item -Destination $env:temp\\$env:computername.zip\"", script)
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

func (glc *getLogsCmd) CreateDefaultStorageContainer() error {
	//default container name is "kuberneteslogs"
	u, _ := url.Parse(fmt.Sprintf("https://%s.blob.core.windows.net/kuberneteslogs", glc.storageAccountName))
	credential, err := azblob.NewSharedKeyCredential(glc.storageAccountName, glc.storageAccountKey)
	if err != nil {
		return errors.Wrap(err, "getting credential from storage account name and key")
	}

	containerURL := azblob.NewContainerURL(*u, azblob.NewPipeline(credential, azblob.PipelineOptions{}))

	_, err = containerURL.Create(context.Background(), azblob.Metadata{}, azblob.PublicAccessContainer)
	if err != nil {
		return errors.Wrap(err, "creating storage container 'kuberneteslogs'")
	}

	glc.storageContainerURL = fmt.Sprintf("https://%s.blob.core.windows.net/kuberneteslogs", glc.storageAccountName)

	return nil
}

func (glc *getLogsCmd) uploadLogsToStorageContainer(node v1.Node) error {
	log.Infof("Uploading logs for %s", node.Name)
	logFileName := fmt.Sprintf("%s.zip", node.Name)
	logFilePath := path.Join(glc.outputDirectory, logFileName)
	logFile, err := os.Open(logFilePath)
	if err != nil {
		return errors.Wrap(err, "opening zipped log file")
	}

	u, _ := url.Parse(fmt.Sprintf("%s/%s", glc.storageContainerURL, logFileName))
	credential, err := azblob.NewSharedKeyCredential(glc.storageAccountName, glc.storageAccountKey)
	if err != nil {
		return errors.Wrap(err, "getting credential from storage account name and key")
	}
	blockBlobURL := azblob.NewBlockBlobURL(*u, azblob.NewPipeline(credential, azblob.PipelineOptions{}))

	_, err = azblob.UploadFileToBlockBlob(context.Background(), logFile, blockBlobURL, azblob.UploadToBlockBlobOptions{})
	if err != nil {
		return errors.Wrap(err, "uploading log file to storage container blob")
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
