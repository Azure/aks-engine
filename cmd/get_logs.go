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
	getLogsShortDescription = "Creates a SSH session to each cluster node and collects log files"
	getLogsLongDescription  = "Creates a SSH session to each cluster node and collects log files"
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
	cs           *api.ContainerService
	locale       *gotext.Locale
	armClient    armhelpers.AKSEngineClient
	masterNodes  []*clusterNode
	linuxNodes   []*clusterNode
	windowsNodes []*clusterNode
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
	command.Flags().StringVar(&glc.linuxSSHPrivateKeyPath, "linux-ssh-private-key", "", "path to a valid private ssh key to access the cluster's nodes")
	command.Flags().StringVar(&glc.linuxScriptPath, "linux-script", "", "path to the log collection script to execute on the cluster's linux nodes")
	command.Flags().StringVarP(&glc.outputDirectory, "output-directory", "o", "", "")
	command.MarkFlagRequired("location")
	command.MarkFlagRequired("api-model")
	command.MarkFlagRequired("apiserver")
	return command
}

func (glc *getLogsCmd) validateArgs(cmd *cobra.Command, args []string) (err error) {
	if glc.locale, err = i18n.LoadTranslations(); err != nil {
		return errors.Wrap(err, "loading translation files")
	}
	glc.location = helpers.NormalizeAzureRegion(glc.location)
	if glc.location == "" {
		return errors.New("location must be specified")
	}
	if _, err := os.Stat(glc.apiModelPath); os.IsNotExist(err) {
		return errors.Errorf("specified api-model does not exist (%s)", glc.apiModelPath)
	}
	if _, err := os.Stat(glc.linuxSSHPrivateKeyPath); os.IsNotExist(err) {
		return errors.Errorf("specified ssh-private-key does not exist (%s)", glc.linuxSSHPrivateKeyPath)
	}
	if _, err := os.Stat(glc.linuxScriptPath); glc.linuxScriptPath != "" && os.IsNotExist(err) {
		return errors.Errorf("specified linux-script does not exist (%s)", glc.linuxScriptPath)
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
	var client *armhelpers.AzureClient
	glc.armClient = client
	return nil
}

func (glc *getLogsCmd) run(cmd *cobra.Command, args []string) (err error) {
	if err = glc.getClusterNodes(); err != nil {
		return errors.Wrap(err, "listing cluster nodes")
	}
	// TODO run in parallel
	for _, l := range glc.masterNodes {
		log.Infof("Processing node: %s\n", l.node.Name)
		out, err := glc.collectLinuxLogs(l.node.Name, l.sshConfig)
		if err != nil {
			log.Warnf("Remote command output: %s", out)
			log.Warnf("Error: %s", err)
		}
	}
	for _, l := range glc.linuxNodes {
		log.Infof("Processing node: %s\n", l.node.Name)
		out, err := glc.collectLinuxLogs(l.node.Name, l.sshConfig)
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
	if glc.cs.Properties.LinuxProfile != nil {
		linuxSSHConfig, err := helpers.SSHClientConfig(glc.cs.Properties.LinuxProfile.AdminUsername, glc.linuxSSHPrivateKeyPath)
		if err != nil {
			return errors.Wrap(err, "creating ssh config")
		}
		for _, node := range nodeList.Items {
			if node.Status.NodeInfo.OperatingSystem == "linux" {
				if strings.HasPrefix(node.Name, "k8s-master") {
					glc.masterNodes = append(glc.masterNodes, &clusterNode{
						node:      node,
						sshConfig: linuxSSHConfig,
					})
				} else {
					glc.linuxNodes = append(glc.linuxNodes, &clusterNode{
						node:      node,
						sshConfig: linuxSSHConfig,
					})
				}
			}
		}
	}
	return nil
}

type clusterNode struct {
	node      v1.Node
	sshConfig *ssh.ClientConfig
}

func (glc *getLogsCmd) collectLinuxLogs(hostname string, config *ssh.ClientConfig) (string, error) {
	// TODO always 22?
	jumpboxPort := "22"
	client, err := helpers.SSHClient(glc.apiserverURI, jumpboxPort, hostname, config)
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
	stdout, err = glc.collectLogs(hostname, client)
	if err != nil {
		return stdout, err
	}
	stdout, err = glc.downloadLogs(hostname, client)
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

func (glc *getLogsCmd) collectLogs(hostname string, client *ssh.Client) (string, error) {
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
