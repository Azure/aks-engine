// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/engine"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/leonelquinteros/gotext"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"k8s.io/api/core/v1"
)

const (
	rotateCertsName             = "rotate-certs"
	rotateCertsShortDescription = "Rotate certificates on an existing Kubernetes cluster"
	rotateCertsLongDescription  = "Rotate certificates on an existing Kubernetes cluster"
	kubeSystemNamespace         = "kube-system"
)

type rotateCertsCmd struct {
	authProvider

	// user input
	resourceGroupName string
	sshFilepath       string
	masterFQDN        string
	location          string
	apiModelPath      string

	// derived
	containerService *api.ContainerService
	apiVersion       string
	locale           *gotext.Locale
	client           armhelpers.AKSEngineClient

	masterNodes []v1.Node
	agentNodes  []v1.Node
	sshConfig   *ssh.ClientConfig
}

func newRotateCertsCmd() *cobra.Command {
	rcc := rotateCertsCmd{
		authProvider: &authArgs{},
	}

	command := &cobra.Command{
		Use:   rotateCertsName,
		Short: rotateCertsShortDescription,
		Long:  rotateCertsLongDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			return rcc.run(cmd, args)
		},
	}

	f := command.Flags()
	f.StringVarP(&rcc.location, "location", "l", "", "location the cluster is deployed in (required)")
	f.StringVarP(&rcc.resourceGroupName, "resource-group", "g", "", "the resource group where the cluster is deployed (required)")
	f.StringVarP(&rcc.apiModelPath, "apimodel", "", "", "the filepath of the cluster's apimodel (defaults to _output)")
	f.StringVarP(&rcc.sshFilepath, "ssh", "", "", "the filepath of a valid private ssh key to access the cluster's nodes (defaults to _output)")
	f.StringVar(&rcc.masterFQDN, "master-FQDN", "", "FQDN for the master load balancer")
	addAuthFlags(rcc.getAuthArgs(), f)

	return command
}

func (rcc *rotateCertsCmd) run(cmd *cobra.Command, args []string) error {

	log.Printf("Start rotating certs")

	var err error

	if err = rcc.getAuthArgs().validateAuthArgs(); err != nil {
		return errors.Wrap(err, "failed to get validate auth args")
	}

	if rcc.client, err = rcc.getAuthArgs().getClient(); err != nil {
		return errors.Wrap(err, "failed to get client")
	}

	ctx, cancel := context.WithTimeout(context.Background(), armhelpers.DefaultARMOperationTimeout)
	defer cancel()
	_, err = rcc.client.EnsureResourceGroup(ctx, rcc.resourceGroupName, rcc.location, nil)
	if err != nil {
		return errors.Wrap(err, "error ensuring resource group")
	}

	// load the cluster configuration.
	if _, err = os.Stat(rcc.apiModelPath); os.IsNotExist(err) {
		return errors.Errorf("specified api model does not exist (%s)", rcc.apiModelPath)
	}

	rcc.locale, err = i18n.LoadTranslations()
	if err != nil {
		return errors.Wrap(err, "error loading translation files")
	}

	log.Printf("Load container service")

	apiloader := &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: rcc.locale,
		},
	}
	rcc.containerService, rcc.apiVersion, err = apiloader.LoadContainerServiceFromFile(rcc.apiModelPath, true, true, nil)
	if err != nil {
		return errors.Wrap(err, "error parsing the api model")
	}

	log.Printf("Get cluster nodes")

	err = rcc.getClusterNodes()
	if err != nil {
		return errors.Wrap(err, "error listing cluster nodes")
	}

	log.Printf("Generate new certs")

	// reset the certificateProfile and use the exisiting certificate generation code to generate new certificates.
	rcc.containerService.Properties.CertificateProfile = &api.CertificateProfile{}
	certsGenerated, _, err := rcc.containerService.SetDefaultCerts()
	if !certsGenerated || err != nil {
		return errors.Wrap(err, "error generating new certificates")
	}

	if _, err = os.Stat(rcc.sshFilepath); os.IsNotExist(err) {
		return errors.Errorf("specified ssh filepath does not exist (%s)", rcc.sshFilepath)
	}
	rcc.setSSHConfig()

	log.Printf("Rotate apiserver")

	err = rcc.rotateApiserver()
	if err != nil {
		return errors.Wrap(err, "error rotating apiserver")
	}

	log.Printf("Rotate kubelet")

	err = rcc.rotateKubelet()
	if err != nil {
		return errors.Wrap(err, "error rotating kubelet")
	}

	log.Printf("Rotate etcd")

	err = rcc.rotateEtcd()
	if err != nil {
		return errors.Wrap(err, "error rotating etcd cluster")
	}

	// TODO proxy certs?

	// TODO: save kubeconfig in files
	log.Printf("Update kubeconfig in cluster")
	err = rcc.updateKubeconfig()
	if err != nil {
		return errors.Wrap(err, "error updating kubeconfig")
	}

	log.Printf("Rebooting all nodes... This might take a few minutes")
	err = rcc.rebootAllNodes(ctx)
	if err != nil {
		return errors.Wrap(err, "error rebooting the nodes")
	}

	log.Printf("Delete Service Accoutns")
	err = rcc.deleteServiceAccounts()
	if err != nil {
		return errors.Wrap(err, "error deleting service accounts")
	}
	// TODO: save apimodel certificateProfile/cert files as output

	log.Printf("Delete all pods")
	err = rcc.deleteAllPods()
	if err != nil {
		return errors.Wrap(err, "error deleting all the pods")
	}

	log.Printf("Successfully rotated etcd and cluster certificates.")

	return nil
}

func (rcc *rotateCertsCmd) getClusterNodes() error {
	kubeClient, err := rcc.getKubeClient()
	if err != nil {
		return errors.Wrap(err, "failed to get Kubernetes Client")
	}
	nodeList, err := kubeClient.ListNodes()
	if err != nil {
		return errors.Wrap(err, "failed to get cluster nodes")
	}
	for _, node := range nodeList.Items {
		if strings.Contains(node.Name, "master") {
			rcc.masterNodes = append(rcc.masterNodes, node)
		} else {
			rcc.agentNodes = append(rcc.agentNodes, node)
		}
	}
	return nil
}

func (rcc *rotateCertsCmd) rebootAllNodes(ctx context.Context) error {
	vmListPage, err := rcc.client.ListVirtualMachines(ctx, rcc.resourceGroupName)
	if err != nil {
		return errors.Wrap(err, "failed to list Virtual Machines in resource group"+rcc.resourceGroupName)
	}
	vmssListPage, err := rcc.client.ListVirtualMachineScaleSets(ctx, rcc.resourceGroupName)
	if err != nil {
		return errors.Wrap(err, "failed to list Virtual Machine Scale Sets in resource group"+rcc.resourceGroupName)
	}
	for _, vm := range vmListPage.Values() {
		rcc.client.RestartVirtualMachine(ctx, rcc.resourceGroupName, *vm.Name)
	}
	for _, vmss := range vmssListPage.Values() {
		rcc.client.RestartVirtualMachineScaleSets(ctx, rcc.resourceGroupName, *vmss.Name, nil)
	}
	return nil
}

func (rcc *rotateCertsCmd) deleteAllPods() error {
	kubeClient, err := rcc.getKubeClient()
	if err != nil {
		return errors.Wrap(err, "failed to get Kubernetes Client")
	}
	pods, err := kubeClient.ListAllPods()
	if err != nil {
		return errors.Wrap(err, "failed to get pods")
	}
	for _, pod := range pods.Items {
		log.Printf("Deleting pod %s", pod.Name)
		err = kubeClient.DeletePod(&pod)
		if err != nil {
			return errors.Wrap(err, "failed to delete pod "+pod.Name)
		}
	}
	return nil
}

func (rcc *rotateCertsCmd) deleteServiceAccounts() error {
	kubeClient, err := rcc.getKubeClient()
	if err != nil {
		return errors.Wrap(err, "failed to get Kubernetes Client")
	}
	saList, err := kubeClient.ListServiceAccounts(kubeSystemNamespace)
	if err != nil {
		return errors.Wrap(err, "failed to get cluster service accounts in namespace "+kubeSystemNamespace)
	}
	for _, sa := range saList.Items {
		switch sa.Name {
		case "kube-dns", "kubernetes-dashboard", "metrics-server":
			log.Printf("Deleting service account %s", sa.Name)
			err = kubeClient.DeleteServiceAccount(&sa)
			if err != nil {
				return errors.Wrap(err, "failed to delete service account "+sa.Name)
			}
		}
	}
	return nil
}

func (rcc *rotateCertsCmd) updateKubeconfig() error {
	kubeconfig, err := engine.GenerateKubeConfig(rcc.containerService.Properties, rcc.location)
	if err != nil {
		return errors.Wrap(err, "error generating kubeconfig")
	}

	for _, host := range rcc.masterNodes {
		cmd := "sudo bash -c \"cat > ~/.kube/config << EOL \n" + strings.Replace(kubeconfig, "\"", "\\\"", -1) + "EOL\""
		out, err := executeCmd(cmd, rcc.masterFQDN, host.Name, "22", rcc.sshConfig)
		if err != nil {
			log.Printf("Command %s output: %s\n", cmd, out)
			return errors.Wrap(err, "failed replacing kubeconfig file")
		}
	}
	return nil
}

func (rcc *rotateCertsCmd) getKubeClient() (armhelpers.KubernetesClient, error) {
	kubeconfig, err := engine.GenerateKubeConfig(rcc.containerService.Properties, rcc.location)
	if err != nil {
		return nil, errors.Wrap(err, "error generating kubeconfig")
	}
	var kubeClient armhelpers.KubernetesClient
	if rcc.client != nil {
		kubeClient, err = rcc.client.GetKubernetesClient("", kubeconfig, time.Second*1, time.Duration(60)*time.Minute)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get a Kubernetes client")
		}
	}
	return kubeClient, nil
}

// Rotate etcd CA and certificates in all of the master nodes.
func (rcc *rotateCertsCmd) rotateEtcd() error {
	caPrivateKeyCmd := "sudo bash -c \"cat > /etc/kubernetes/certs/ca.key << EOL \n" + rcc.containerService.Properties.CertificateProfile.CaPrivateKey + "EOL\""
	caCertificateCmd := "sudo bash -c \"cat > /etc/kubernetes/certs/ca.crt << EOL \n" + rcc.containerService.Properties.CertificateProfile.CaCertificate + "EOL\""
	etcdServerPrivateKeyCmd := "sudo bash -c \"cat > /etc/kubernetes/certs/etcdserver.key << EOL \n" + rcc.containerService.Properties.CertificateProfile.EtcdServerPrivateKey + "EOL\""
	etcdServerCertificateCmd := "sudo bash -c \"cat > /etc/kubernetes/certs/etcdserver.crt << EOL \n" + rcc.containerService.Properties.CertificateProfile.EtcdServerCertificate + "EOL\""
	etcdClientPrivateKeyCmd := "sudo bash -c \"cat > /etc/kubernetes/certs/etcdclient.key << EOL \n" + rcc.containerService.Properties.CertificateProfile.EtcdClientPrivateKey + "EOL\""
	etcdClientCertificateCmd := "sudo bash -c \"cat > /etc/kubernetes/certs/etcdclient.crt << EOL \n" + rcc.containerService.Properties.CertificateProfile.EtcdClientCertificate + "EOL\""

	for i, host := range rcc.masterNodes {
		log.Printf("Ranging over node: %s\n", host.Name)
		etcdPeerPrivateKeyCmd := "sudo bash -c \"cat > /etc/kubernetes/certs/etcdpeer" + strconv.Itoa(i) + ".key << EOL \n" + rcc.containerService.Properties.CertificateProfile.EtcdPeerPrivateKeys[i] + "EOL\""
		etcdPeerCertificateCmd := "sudo bash -c \"cat > /etc/kubernetes/certs/etcdpeer" + strconv.Itoa(i) + ".crt << EOL \n" + rcc.containerService.Properties.CertificateProfile.EtcdPeerCertificates[i] + "EOL\""

		for _, cmd := range []string{caPrivateKeyCmd, caCertificateCmd} {
			out, err := executeCmd(cmd, rcc.masterFQDN, host.Name, "22", rcc.sshConfig)
			if err != nil {
				log.Printf("Command %s output: %s\n", cmd, out)
				return errors.Wrap(err, "failed replacing certificate file")
			}
		}

		out, err := executeCmd("sudo systemctl restart etcd", rcc.masterFQDN, host.Name, "22", rcc.sshConfig)
		if err != nil {
			log.Printf("Command `sudo systemctl restart etcd` output: %s\n", out)
			return errors.Wrap(err, "failed to restart etcd")
		}

		for _, cmd := range []string{etcdServerPrivateKeyCmd, etcdServerCertificateCmd, etcdClientPrivateKeyCmd, etcdClientCertificateCmd, etcdPeerPrivateKeyCmd, etcdPeerCertificateCmd} {
			out, err := executeCmd(cmd, rcc.masterFQDN, host.Name, "22", rcc.sshConfig)
			if err != nil {
				log.Printf("Command %s output: %s\n", cmd, out)
				return errors.Wrap(err, "failed replacing certificate file")
			}
		}

		out, err = executeCmd("sudo systemctl restart etcd", rcc.masterFQDN, host.Name, "22", rcc.sshConfig)
		if err != nil {
			log.Printf("Command `sudo systemctl restart etcd` output: %s\n", out)
			return errors.Wrap(err, "failed to restart etcd")
		}
	}

	return nil
}

// From the first master node, rotate apiserver certificates in the nodes.
func (rcc *rotateCertsCmd) rotateApiserver() error {
	caCertificateCmd := "sudo bash -c \"cat > /etc/kubernetes/certs/ca.crt << EOL \n" + rcc.containerService.Properties.CertificateProfile.CaCertificate + "EOL\""
	apiServerPrivateKeyCmd := "sudo bash -c \"cat > /etc/kubernetes/certs/apiserver.key << EOL \n" + rcc.containerService.Properties.CertificateProfile.APIServerPrivateKey + "EOL\""
	apiServerCertificateCmd := "sudo bash -c \"cat > /etc/kubernetes/certs/apiserver.crt << EOL \n" + rcc.containerService.Properties.CertificateProfile.APIServerCertificate + "EOL\""

	for _, host := range rcc.masterNodes {
		log.Printf("Ranging over node: %s\n", host.Name)
		for _, cmd := range []string{apiServerPrivateKeyCmd, apiServerCertificateCmd} {
			out, err := executeCmd(cmd, rcc.masterFQDN, host.Name, "22", rcc.sshConfig)
			if err != nil {
				log.Printf("Command %s output: %s\n", cmd, out)
				return errors.Wrap(err, "failed replacing certificate file")
			}
		}
	}

	for _, host := range rcc.agentNodes {
		log.Printf("Ranging over node: %s\n", host.Name)
		for _, cmd := range []string{caCertificateCmd, apiServerCertificateCmd} {
			out, err := executeCmd(cmd, rcc.masterFQDN, host.Name, "22", rcc.sshConfig)
			if err != nil {
				log.Printf("Command %s output: %s\n", cmd, out)
				return errors.Wrap(err, "failed replacing certificate file")
			}
		}
	}
	return nil
}

func (rcc *rotateCertsCmd) rotateKubelet() error {
	clientCertificateCmd := "sudo bash -c \"cat > /etc/kubernetes/certs/client.crt << EOL \n" + rcc.containerService.Properties.CertificateProfile.ClientCertificate + "EOL\""
	clientPrivateKeyCmd := "sudo bash -c \"cat > /etc/kubernetes/certs/client.key << EOL \n" + rcc.containerService.Properties.CertificateProfile.ClientPrivateKey + "EOL\""

	for _, host := range append(rcc.masterNodes, rcc.agentNodes...) {
		log.Printf("Ranging over node: %s\n", host.Name)
		for _, cmd := range []string{clientCertificateCmd, clientPrivateKeyCmd} {
			out, err := executeCmd(cmd, rcc.masterFQDN, host.Name, "22", rcc.sshConfig)
			if err != nil {
				log.Printf("Command %s output: %s\n", cmd, out)
				return errors.Wrap(err, "failed replacing certificate file")
			}
		}
	}
	return nil
}

func (rcc *rotateCertsCmd) setSSHConfig() {
	rcc.sshConfig = &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		User:            "azureuser",
		Auth: []ssh.AuthMethod{
			publicKeyFile(rcc.sshFilepath),
		},
	}
}

func publicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}
	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}

func executeCmd(command, masterFQDN, hostname string, port string, config *ssh.ClientConfig) (string, error) {
	// Dial connection to the master via public load balancer
	lbClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", masterFQDN, port), config)
	if err != nil {
		return "", errors.Wrap(err, "error Dialing LB")
	}

	// Dial a connection to the agent host, from the master
	conn, err := lbClient.Dial("tcp", fmt.Sprintf("%s:%s", hostname, port))
	if err != nil {
		return "", errors.Wrap(err, "error Dialing host")
	}

	ncc, chans, reqs, err := ssh.NewClientConn(conn, hostname, config)
	if err != nil {
		return "", errors.Wrap(err, "error starting new client connection to host")
	}

	sClient := ssh.NewClient(ncc, chans, reqs)

	session, err := sClient.NewSession()

	if err != nil {
		return "", errors.Wrap(err, "error opening SSH session")
	}
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf

	err = session.Run(command)
	if err != nil {
		return fmt.Sprintf("%s -> %s", hostname, stdoutBuf.String()), errors.Wrap(err, "error running command")
	}

	return fmt.Sprintf("%s -> %s", hostname, stdoutBuf.String()), nil
}
