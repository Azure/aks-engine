// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/aks-engine/pkg/helpers/ssh"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/Azure/aks-engine/pkg/kubernetes"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/leonelquinteros/gotext"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	getLogsName             = "get-logs"
	getLogsShortDescription = "Collect logs and current cluster nodes configuration."
	getLogsLongDescription  = "Collect deployment logs, running daemons/services logs and current nodes configuration."
)

const (
	getLogsLinuxVHDScriptPath      = "/opt/azure/containers/collect-logs.sh"
	getLogsCustomLinuxScriptPath   = "/tmp/collect-logs.sh"
	getLogsWindowsVHDScriptPath    = "c:\\k\\debug\\collect-windows-logs.ps1"
	getLogsCustomWindowsScriptPath = "$env:temp\\collect-windows-logs.ps1"
	getLogsUploadTimeout           = 300 * time.Second
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
	uploadSASURL           string
	nodeNames              []string
	// computed
	cs                  *api.ContainerService
	locale              *gotext.Locale
	linuxAuthConfig     *ssh.AuthConfig
	linuxVHDScript      *ssh.RemoteFile
	linuxCustomScript   *ssh.RemoteFile
	windowsAuthConfig   *ssh.AuthConfig
	windowsVHDScript    *ssh.RemoteFile
	windowsCustomScript *ssh.RemoteFile
	jumpbox             *ssh.JumpBox
}

func newGetLogsCmd() *cobra.Command {
	glc := getLogsCmd{}
	command := &cobra.Command{
		Use:   getLogsName,
		Short: getLogsShortDescription,
		Long:  getLogsLongDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := glc.validateArgs(); err != nil {
				return errors.Wrap(err, "validating get-logs args")
			}
			if err := glc.loadAPIModel(); err != nil {
				return errors.Wrap(err, "loading API model")
			}
			if err := glc.init(); err != nil {
				return errors.Wrap(err, "loading API model")
			}
			cmd.SilenceUsage = true
			return glc.run()
		},
	}
	command.Flags().StringVarP(&glc.location, "location", "l", "", "Azure location where the cluster is deployed (required)")
	command.Flags().StringVarP(&glc.apiModelPath, "api-model", "m", "", "path to the generated apimodel.json file (required)")
	command.Flags().StringVar(&glc.sshHostURI, "ssh-host", "", "FQDN, or IP address, of an SSH listener that can reach all nodes in the cluster (required)")
	command.Flags().StringVar(&glc.linuxSSHPrivateKeyPath, "linux-ssh-private-key", "", "path to a valid private SSH key to access the cluster's Linux nodes (required)")
	command.Flags().StringVar(&glc.linuxScriptPath, "linux-script", "", "path to the log collection script to execute on the cluster's Linux nodes (required if distro is not aks-ubuntu-18.04)")
	command.Flags().StringVar(&glc.windowsScriptPath, "windows-script", "", "path to the log collection script to execute on the cluster's Windows nodes (required if distro is not aks-windows)")
	command.Flags().StringVarP(&glc.outputDirectory, "output-directory", "o", "", "collected logs destination directory, derived from --api-model if missing")
	command.Flags().BoolVarP(&glc.controlPlaneOnly, "control-plane-only", "", false, "get logs from control plane VMs only")
	command.Flags().StringVarP(&glc.uploadSASURL, "upload-sas-url", "", "", "Azure Storage Account SAS URL to upload the collected logs")
	command.Flags().StringSliceVar(&glc.nodeNames, "vm-names", nil, "get logs from the VM name list only (comma-separated names)")
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
	if glc.uploadSASURL != "" {
		exp, err := regexp.Compile(`^/\w+`)
		if err != nil {
			return err
		}
		sasURL, err := url.ParseRequestURI(glc.uploadSASURL)
		if err != nil {
			return errors.Errorf("error parsing upload SAS URL")
		}
		if !exp.MatchString(sasURL.Path) {
			return errors.New("invalid upload SAS URL format, expected 'https://{blob-service-uri}/{container-name}?{sas-token}'")
		}
	}
	if glc.nodeNames != nil && len(glc.nodeNames) == 0 {
		return errors.New("--vm-names cannot be empty")
	}
	if glc.nodeNames != nil && glc.controlPlaneOnly {
		return errors.New("--control-plane-only and --vm-names are mutually exclusive")
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
	return
}

func (glc *getLogsCmd) init() (err error) {
	if glc.linuxScriptPath != "" {
		sc, err := ioutil.ReadFile(glc.linuxScriptPath)
		if err != nil {
			return errors.Wrapf(err, "error reading log collection script %s", glc.linuxScriptPath)
		}
		glc.linuxCustomScript = &ssh.RemoteFile{
			Path: getLogsCustomLinuxScriptPath, Permissions: "744", Owner: "root:root", Content: sc}
	}
	glc.linuxVHDScript = &ssh.RemoteFile{Path: getLogsLinuxVHDScriptPath}
	glc.linuxAuthConfig = &ssh.AuthConfig{
		User:           glc.cs.Properties.LinuxProfile.AdminUsername,
		PrivateKeyPath: glc.linuxSSHPrivateKeyPath,
	}
	if glc.windowsScriptPath != "" {
		sc, err := ioutil.ReadFile(glc.windowsScriptPath)
		if err != nil {
			return errors.Wrapf(err, "error reading log collection script %s", glc.windowsScriptPath)
		}
		glc.windowsCustomScript = &ssh.RemoteFile{
			Path: getLogsCustomWindowsScriptPath, Permissions: "", Owner: "", Content: sc}
	}
	glc.windowsVHDScript = &ssh.RemoteFile{Path: getLogsWindowsVHDScriptPath}
	if glc.cs.Properties.WindowsProfile != nil {
		if glc.cs.Properties.WindowsProfile.GetSSHEnabled() {
			glc.windowsAuthConfig = &ssh.AuthConfig{
				User:     glc.cs.Properties.WindowsProfile.AdminUsername,
				Password: glc.cs.Properties.WindowsProfile.AdminPassword,
			}
		} else {
			log.Warn("Skipping Windows nodes as SSH is not enabled")
		}
	}
	glc.jumpbox = &ssh.JumpBox{
		URI: glc.sshHostURI, Port: 22, OperatingSystem: api.Linux, AuthConfig: glc.linuxAuthConfig}
	return
}

func (glc *getLogsCmd) run() error {
	kubeClient, err := getKubeClient(glc.cs, 10*time.Second, 10*time.Minute)
	if err != nil {
		return errors.Wrap(err, "creating Kubernetes client")
	}
	nodes := getClusterNodes(glc, kubeClient)
	nodeScripts := getClusterNodeScripts(glc, nodes)
	if len(nodeScripts) == 0 {
		log.Info("All nodes skipped")
		return nil
	}
	for node, script := range nodeScripts {
		err = collectLogs(glc, node, script)
		if err != nil {
			return err
		}
	}
	log.Infof("Logs downloaded to %s", glc.outputDirectory)
	if glc.uploadSASURL != "" {
		for node := range nodeScripts {
			err = uploadLogs(node, glc.outputDirectory, glc.uploadSASURL)
			if err != nil {
				log.Warnf("Error uploading %s logs", node.URI)
				log.Debugf("Error: %s", err)
			}
		}
	}
	return err
}

// getClusterNodes returns the target node list
func getClusterNodes(glc *getLogsCmd, kubeClient kubernetes.NodeLister) (nodes []*ssh.RemoteHost) {
	if glc.nodeNames != nil {
		for _, nodeName := range glc.nodeNames {
			if strings.HasPrefix(nodeName, api.DefaultOrchestratorName) {
				log.Infof("Treating node %s as a Linux agent node", nodeName)
				nodes = append(nodes, &ssh.RemoteHost{
					URI: nodeName, Port: 22, OperatingSystem: api.Linux, AuthConfig: glc.linuxAuthConfig, Jumpbox: glc.jumpbox})
			} else {
				log.Infof("Treating node %s as a Windows agent node", nodeName)
				if glc.windowsAuthConfig != nil {
					nodes = append(nodes, &ssh.RemoteHost{
						URI: nodeName, Port: 22, OperatingSystem: api.Windows, AuthConfig: glc.windowsAuthConfig, Jumpbox: glc.jumpbox})
				} else {
					log.Infof("Skipping node %s, WindowsProfile was not provided", nodeName)
				}
			}
		}
		return nodes
	}
	nodeList, err := kubeClient.ListNodes()
	if err != nil {
		log.Warnf("Error retrieving node list from apiserver: %s", err)
		log.Info("Collecting logs from control plane nodes only")
		for i := 0; i < glc.cs.Properties.MasterProfile.Count; i++ {
			name := fmt.Sprintf("%s%d", glc.cs.Properties.GetMasterVMPrefix(), i)
			nodes = append(nodes, &ssh.RemoteHost{
				URI: name, Port: 22, OperatingSystem: api.Linux, AuthConfig: glc.linuxAuthConfig, Jumpbox: glc.jumpbox})
		}
		return nodes
	}
	for _, node := range nodeList.Items {
		if isMasterNode(node.Name, glc.cs.Properties.GetMasterVMPrefix()) || !glc.controlPlaneOnly {
			switch api.OSType(strings.Title(node.Status.NodeInfo.OperatingSystem)) {
			case api.Linux:
				nodes = append(nodes, &ssh.RemoteHost{
					URI: node.Name, Port: 22, OperatingSystem: api.Linux, AuthConfig: glc.linuxAuthConfig, Jumpbox: glc.jumpbox})
			case api.Windows:
				if glc.windowsAuthConfig != nil {
					nodes = append(nodes, &ssh.RemoteHost{
						URI: node.Name, Port: 22, OperatingSystem: api.Windows, AuthConfig: glc.windowsAuthConfig, Jumpbox: glc.jumpbox})
				}
			default:
				log.Infof("Skipping node %s, could not determine operating system", node.Name)
			}
		}
	}
	return nodes
}

// getClusterNodeScripts maps target nodes with a log collection script
func getClusterNodeScripts(glc *getLogsCmd, nodes []*ssh.RemoteHost) map[*ssh.RemoteHost]*ssh.RemoteFile {
	nodeScript := make(map[*ssh.RemoteHost]*ssh.RemoteFile)
	poolHasScript := make(map[string]bool)
	isWindowsSkipped := false
	for _, node := range nodes {
		switch node.OperatingSystem {
		case api.Linux:
			if isMasterNode(node.URI, glc.cs.Properties.GetMasterVMPrefix()) && glc.cs.Properties.MasterProfile.IsVHDDistro() {
				nodeScript[node] = glc.linuxVHDScript
			} else {
				for i, pool := range glc.cs.Properties.AgentPoolProfiles {
					if pool.IsVHDDistro() && glc.cs.Properties.IsAgentPoolMember(node.URI, pool, i) {
						nodeScript[node] = glc.linuxVHDScript
					}
				}
			}
			if glc.linuxCustomScript != nil {
				nodeScript[node] = glc.linuxCustomScript
			}
			_, ok := nodeScript[node]
			poolName := strings.Split(node.URI, "-")[1]
			poolHasScript[poolName] = ok
		case api.Windows:
			if glc.cs.Properties.WindowsProfile != nil && glc.cs.Properties.WindowsProfile.IsVHDDistro() {
				nodeScript[node] = glc.windowsVHDScript
			}
			if glc.windowsCustomScript != nil {
				nodeScript[node] = glc.windowsCustomScript
			}
			if _, ok := nodeScript[node]; !ok {
				isWindowsSkipped = true
			}
		}
	}
	for pool, hasScript := range poolHasScript {
		if !hasScript {
			log.Warnf("Skipping node pool '%s' as flag '--linux-script' is not set and the pool distro is not aks-ubuntu-18.04", pool)
		}
	}
	if isWindowsSkipped {
		log.Warn("Skipping Windows nodes as flag '--windows-script' is not set and the profile distro is not aks-windows")
	}
	return nodeScript
}

// collectLogs uploads the log collection script (if needed), executes the script and downloads the collected logs
func collectLogs(glc *getLogsCmd, node *ssh.RemoteHost, script *ssh.RemoteFile) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	log.Infof("Processing node: %s", node.URI)
	if script.Content != nil {
		stdout, err := ssh.CopyToRemote(ctx, node, script)
		if err != nil {
			return errors.Wrap(err, stdout)
		}
	}
	isAzureStack := glc.cs.Properties.IsAzureStackCloud()
	stdout, err := ssh.ExecuteRemote(ctx, node, collectLogsScript(script, node.OperatingSystem, isAzureStack))
	if err != nil {
		return errors.Wrap(err, stdout)
	}
	src := fileToDownload(node.OperatingSystem, node.URI)
	dst := path.Join(glc.outputDirectory, fmt.Sprintf("%s.zip", node.URI))
	stdout, err = ssh.CopyFromRemote(ctx, node, src, dst)
	if err != nil {
		return errors.Wrap(err, stdout)
	}
	return err
}

// uploadLogs uploads collected logs to an azure storage account
func uploadLogs(node *ssh.RemoteHost, outputDirectory, uploadSASURL string) error {
	log.Infof("Uploading %s logs", node.URI)
	ctx, cancel := context.WithTimeout(context.Background(), getLogsUploadTimeout)
	defer cancel()
	fp := path.Join(outputDirectory, fmt.Sprintf("%s.zip", node.URI))
	f, err := os.Open(fp)
	if err != nil {
		return errors.Wrapf(err, "reading file %s", fp)
	}
	sas, err := url.Parse(uploadSASURL)
	if err != nil {
		return errors.Wrap(err, "parsing upload SAS URL")
	}
	sas.Path = path.Join(sas.Path, fmt.Sprintf("%s.zip", node.URI))
	_, err = uploadToSASURL(ctx, f, sas)
	if err != nil {
		return err
	}
	return nil
}

func uploadToSASURL(ctx context.Context, file *os.File, destination *url.URL) (azblob.CommonResponse, error) {
	p := azblob.NewPipeline(azblob.NewAnonymousCredential(), azblob.PipelineOptions{})
	u := azblob.NewBlobURL(*destination, p).ToBlockBlobURL()
	cr, err := azblob.UploadFileToBlockBlob(ctx, file, u, azblob.UploadToBlockBlobOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "uploading to storage account")
	}
	return cr, nil
}

func collectLogsScript(f *ssh.RemoteFile, os api.OSType, isAzureStack bool) string {
	switch os {
	case api.Linux:
		if isAzureStack {
			return fmt.Sprintf("sudo -E bash -c \"AZURE_ENV=AzureStackCloud %s\"", f.Path)
		}
		return fmt.Sprintf("sudo -E bash -c %s", f.Path)
	case api.Windows:
		return fmt.Sprintf("powershell -command \"iex %s | Where-Object { $_.extension -eq '.zip' } | Copy-Item -Destination $env:temp\\$env:computername.zip\"", f.Path)
	default:
		return ""
	}
}

func fileToDownload(os api.OSType, nodeName string) *ssh.RemoteFile {
	switch os {
	case api.Linux:
		return &ssh.RemoteFile{Path: "/tmp/logs.zip"}
	case api.Windows:
		return &ssh.RemoteFile{Path: fmt.Sprintf("%%TEMP%%\\%s.zip", nodeName)}
	default:
		return nil
	}
}

func isMasterNode(vmName, masterPrefix string) bool {
	return strings.HasPrefix(vmName, masterPrefix)
}
