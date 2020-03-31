// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/engine"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/aks-engine/pkg/i18n"
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
	apiserverURI           string
	linuxSSHPrivateKeyPath string
	linuxScriptPath        string
	outputDirectory        string
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
			if err := glc.validateArgs(cmd, args); err != nil {
				return errors.Wrap(err, "validating get-logs args")
			}
			if err := glc.loadAPIModel(); err != nil {
				return errors.Wrap(err, "loading API model")
			}
			return glc.run(cmd, args)
		},
	}
	command.Flags().StringVarP(&glc.location, "location", "l", "", "Azure location where the cluster is deployed (required)")
	command.Flags().StringVarP(&glc.apiModelPath, "api-model", "m", "", "path to the generated apimodel.json file (required)")
	command.Flags().StringVar(&glc.apiserverURI, "apiserver", "", "apiserver endpoint (required)")
	command.Flags().StringVar(&glc.linuxSSHPrivateKeyPath, "linux-ssh-private-key", "", "path to a valid private ssh key to access the cluster's Linux nodes (required)")
	command.Flags().StringVar(&glc.linuxScriptPath, "linux-script", "", "path to the log collection script to execute on the cluster's Linux nodes")
	command.Flags().StringVarP(&glc.outputDirectory, "output-directory", "o", "", "collected logs destination directory, derived from --api-model if missing")
	command.MarkFlagRequired("location")
	command.MarkFlagRequired("api-model")
	command.MarkFlagRequired("apiserver")
	command.MarkFlagRequired("linux-ssh-private-key")
	return command
}

func (glc *getLogsCmd) validateArgs(cmd *cobra.Command, args []string) (err error) {
	if glc.locale, err = i18n.LoadTranslations(); err != nil {
		return errors.Wrap(err, "loading translation files")
	}
	glc.location = helpers.NormalizeAzureRegion(glc.location)
	if glc.location == "" {
		return errors.New("--location must be specified")
	}
	if _, err := os.Stat(glc.apiModelPath); os.IsNotExist(err) {
		return errors.Errorf("specified --api-model does not exist (%s)", glc.apiModelPath)
	}
	if _, err := os.Stat(glc.linuxSSHPrivateKeyPath); glc.linuxSSHPrivateKeyPath != "" && os.IsNotExist(err) {
		return errors.Errorf("specified --linux-ssh-private-key does not exist (%s)", glc.linuxSSHPrivateKeyPath)
	}
	if _, err := os.Stat(glc.linuxScriptPath); glc.linuxScriptPath != "" && os.IsNotExist(err) {
		return errors.Errorf("specified --linux-script does not exist (%s)", glc.linuxScriptPath)
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
	if glc.cs.Location == "" {
		glc.cs.Location = glc.location
	} else if glc.cs.Location != glc.location {
		return errors.New("--location flag does not match api-model location")
	}
	auth, err := helpers.PublicKeyAuth(glc.linuxSSHPrivateKeyPath)
	if err != nil {
		return errors.Wrap(err, "creating linux ssh config")
	}
	glc.linuxSSHConfig = helpers.SSHClientConfig(glc.cs.Properties.LinuxProfile.AdminUsername, auth)
	if glc.cs.Properties.WindowsProfile != nil {
		auth := ssh.Password(glc.cs.Properties.WindowsProfile.AdminPassword)
		glc.windowsSSHConfig = helpers.SSHClientConfig(glc.cs.Properties.WindowsProfile.AdminUsername, auth)
	}
	var client *armhelpers.AzureClient
	glc.armClient = client
	return nil
}

func (glc *getLogsCmd) run(cmd *cobra.Command, args []string) (err error) {
	if err = glc.getClusterNodes(); err != nil {
		return errors.Wrap(err, "listing cluster nodes")
	}
	// TODO run in parallel
	// for _, n := range glc.masterNodes {
	// 	log.Infof("Processing node: %s\n", n.Name)
	// 	out, err := glc.collectLinuxLogs(n.Name, glc.linuxSSHConfig)
	// 	if err != nil {
	// 		log.Warnf("Remote command output: %s", out)
	// 		log.Warnf("Error: %s", err)
	// 	}
	// }
	// for _, n := range glc.linuxNodes {
	// 	log.Infof("Processing node: %s\n", n.Name)
	// 	out, err := glc.collectLinuxLogs(n.Name, glc.linuxSSHConfig)
	// 	if err != nil {
	// 		log.Warnf("Remote command output: %s", out)
	// 		log.Warnf("Error: %s", err)
	// 	}
	// }
	for _, n := range glc.windowsNodes {
		log.Infof("Processing node: %s\n", n.Name)
		out, err := glc.collectWindowsLogs(n.Name, glc.windowsSSHConfig)
		if err != nil {
			log.Warnf("Remote command output: %s", out)
			log.Warnf("Error: %s", err)
		}
	}
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
		return errors.Wrap(err, "listing cluster nodes")
	}
	for _, node := range nodeList.Items {
		if strings.EqualFold(node.Status.NodeInfo.OperatingSystem, "linux") {
			if strings.HasPrefix(node.Name, "k8s-master") {
				glc.masterNodes = append(glc.masterNodes, node)
			} else {
				glc.linuxNodes = append(glc.linuxNodes, node)
			}
		}
		if strings.EqualFold(node.Status.NodeInfo.OperatingSystem, "windows") {
			glc.windowsNodes = append(glc.windowsNodes, node)
		}
	}
	return nil
}

func (glc *getLogsCmd) collectLinuxLogs(hostname string, config *ssh.ClientConfig) (string, error) {
	// TODO always 22?
	jumpboxPort := "22"
	client, err := helpers.SSHClient(glc.apiserverURI, jumpboxPort, hostname, glc.linuxSSHConfig, config)
	if err != nil {
		return "", errors.Wrap(err, "creating SSH client")
	}
	defer client.Close()

	var stdout string
	if glc.linuxScriptPath != "" {
		stdout, err = glc.uploadScript(hostname, client)
		if err != nil {
			return stdout, err
		}
	}
	stdout, err = glc.execCollectLogs(hostname, client)
	if err != nil {
		return stdout, err
	}
	stdout, err = glc.downloadLogs(hostname, client)
	if err != nil {
		return stdout, err
	}
	return "", nil
}

func (glc *getLogsCmd) collectWindowsLogs(hostname string, config *ssh.ClientConfig) (string, error) {
	jumpboxPort := "22"
	client, err := helpers.SSHClient(glc.apiserverURI, jumpboxPort, hostname, glc.linuxSSHConfig, config)
	if err != nil {
		return "", errors.Wrap(err, "creating SSH client")
	}
	defer client.Close()

	stdout, err := glc.execCollectWindowsLogs(hostname, client)
	if err != nil {
		return stdout, err
	}
	stdout, err = glc.downloadWindowsLogs(hostname, client)
	if err != nil {
		return stdout, err
	}
	return "", nil
}

func (glc *getLogsCmd) uploadScript(hostname string, client *ssh.Client) (string, error) {
	scriptContent, err := ioutil.ReadFile(glc.linuxScriptPath)
	if err != nil {
		return "", errors.Wrap(err, "reading log collection script content")
	}
	log.Debugf("Uploading log collection script (%s)\n", glc.linuxScriptPath)
	session, err := client.NewSession()
	if err != nil {
		return "", errors.Wrap(err, "creating SSH session")
	}
	defer session.Close()
	session.Stdin = bytes.NewReader(scriptContent)
	if co, err := session.CombinedOutput("bash -c \"cat /dev/stdin > /tmp/collect-logs.sh\""); err != nil {
		return fmt.Sprintf("%s -> %s", hostname, string(co)), errors.Wrap(err, "uploading log collection script")
	}
	return "", nil
}

func (glc *getLogsCmd) execCollectLogs(hostname string, client *ssh.Client) (string, error) {
	log.Debug("Collecting logs\n")
	session, err := client.NewSession()
	if err != nil {
		return "", errors.Wrap(err, "creating SSH session")
	}
	defer session.Close()

	var script, cmd string
	if glc.linuxScriptPath != "" {
		script = "/tmp/collect-logs.sh"
		cmd = fmt.Sprintf("sudo chmod +x %s; %s; rm %s", script, script, script)
	} else {
		cmd = "sudo /opt/azure/containers/collect-logs.sh"
	}

	if co, err := session.CombinedOutput(fmt.Sprintf("bash -c \"%s\"", cmd)); err != nil {
		return fmt.Sprintf("%s -> %s", hostname, string(co)), errors.Wrap(err, "collecting logs on remote host")
	}
	return "", nil
}

func (glc *getLogsCmd) execCollectWindowsLogs(hostname string, client *ssh.Client) (string, error) {
	log.Debug("Collecting logs\n")
	session, err := client.NewSession()
	if err != nil {
		return "", errors.Wrap(err, "creating SSH session")
	}
	defer session.Close()

	cmd := "powershell -command \"c:\\k\\debug\\collect-windows-logs.ps1 | Where-Object { $_.extension -eq '.zip' } | Copy-Item -Destination $env:temp\\$env:computername.zip\""
	if co, err := session.CombinedOutput(cmd); err != nil {
		return fmt.Sprintf("%s -> %s", hostname, string(co)), errors.Wrap(err, "collecting logs on remote host")
	}
	return "", nil
}

func (glc *getLogsCmd) downloadLogs(hostname string, client *ssh.Client) (string, error) {
	log.Debug("Downloading logs\n")
	session, err := client.NewSession()
	if err != nil {
		return "", errors.Wrap(err, "creating SSH session")
	}
	defer session.Close()

	localFileName := fmt.Sprintf("%s.zip", hostname)
	localFilePath := path.Join(glc.outputDirectory, localFileName)
	file, err := os.OpenFile(localFilePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return "", errors.Wrap(err, "opening destination file")
	}
	defer file.Close()

	stdout, err := session.StdoutPipe()
	if err != nil {
		return "", errors.Wrap(err, "opening ssh session stdout pipe")
	}

	if err := session.Start("bash -c \"cat /tmp/logs.zip > /dev/stdout\""); err != nil {
		return fmt.Sprintf("%s -> %s", hostname, session.Stderr), errors.Wrap(err, "downloading logs from remote host")
	}
	_, err = io.Copy(file, io.TeeReader(stdout, &DownloadProgressWriter{}))
	if err != nil {
		return "", errors.Wrap(err, "downloading logs")
	}

	fmt.Println()
	return "", nil
}

func (glc *getLogsCmd) downloadWindowsLogs(hostname string, client *ssh.Client) (string, error) {
	log.Debug("Downloading logs\n")
	session, err := client.NewSession()
	if err != nil {
		return "", errors.Wrap(err, "creating SSH session")
	}
	defer session.Close()

	localFileName := fmt.Sprintf("%s.zip", hostname)
	localFilePath := path.Join(glc.outputDirectory, localFileName)
	file, err := os.OpenFile(localFilePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return "", errors.Wrap(err, "opening destination file")
	}
	defer file.Close()

	stdout, err := session.StdoutPipe()
	if err != nil {
		return "", errors.Wrap(err, "opening ssh session stdout pipe")
	}

	cmd := "type %TEMP%"
	if err := session.Start(fmt.Sprintf("%s\\%s.zip", cmd, hostname)); err != nil {
		return fmt.Sprintf("%s -> %s", hostname, session.Stderr), errors.Wrap(err, "downloading logs from remote host")
	}
	_, err = io.Copy(file, io.TeeReader(stdout, &DownloadProgressWriter{}))
	if err != nil {
		return "", errors.Wrap(err, "downloading logs")
	}

	fmt.Println()
	return "", nil
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
