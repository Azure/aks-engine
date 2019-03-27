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
	"golang.org/x/crypto/ssh/agent"
	"k8s.io/api/core/v1"
)

const (
	rotateCertsName             = "rotate-certs"
	rotateCertsShortDescription = "Rotate certificates on an existing Kubernetes cluster"
	rotateCertsLongDescription  = "Rotate certificates on an existing Kubernetes cluster"
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
	masterNodes      []*v1.Node
	agentNodes       []*v1.Node
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

	apiloader := &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: rcc.locale,
		},
	}
	rcc.containerService, rcc.apiVersion, err = apiloader.LoadContainerServiceFromFile(rcc.apiModelPath, true, true, nil)
	if err != nil {
		return errors.Wrap(err, "error parsing the api model")
	}

	err = rcc.getClusterNodes()
	if err != nil {
		return errors.Wrap(err, "error listing cluster nodes")
	}

	// reset the certificateProfile and use the exisiting certificate generation code to generate new certificates.
	rcc.containerService.Properties.CertificateProfile = &api.CertificateProfile{}
	certsGenerated, _, err := rcc.containerService.SetDefaultCerts()
	if !certsGenerated || err != nil {
		return errors.Wrap(err, "error generating new certificates")
	}

	err = rcc.rotateEtcd()
	if err != nil {
		return errors.Wrap(err, "error rotating etcd cluster")
	}

	err = rcc.rotateApiserver()
	if err != nil {
		return errors.Wrap(err, "error rotating apiserver")
	}

	err = rcc.rotateKubelet()
	if err != nil {
		return errors.Wrap(err, "error rotating kubelet")
	}

	// Update the kubeconfig and rotate the kubelet certificates.
	// kubeConfig, err := engine.GenerateKubeConfig(rcc.containerService.Properties, rcc.location)
	// if err != nil {
	// 	return errors.Wrap(err, "error generating kubeconfig")
	// }

	return nil

	// TODO: save kubeconfig

	// TODO: save apimodel certificateProfile?
}

func (rcc *rotateCertsCmd) getClusterNodes() error {
	kubeConfig, err := engine.GenerateKubeConfig(rcc.containerService.Properties, rcc.location)
	if err != nil {
		return errors.Wrap(err, "failed to generate kubeconfig")
	}
	var kubeClient armhelpers.KubernetesClient
	if rcc.client != nil {
		k, err := rcc.client.GetKubernetesClient("", kubeConfig, time.Second*1, time.Duration(60)*time.Minute)
		if err != nil {
			return errors.Wrap(err, "failed to get a Kubernetes client")
		}
		kubeClient = k
	}

	nodeList, err := kubeClient.ListNodes()
	if err != nil {
		return errors.Wrap(err, "failed to list nodes")
	}
	for _, node := range nodeList.Items {
		if strings.Contains(node.Name, "master") {
			rcc.masterNodes = append(rcc.masterNodes, &node)
		} else {
			rcc.agentNodes = append(rcc.agentNodes, &node)
		}
	}
	return nil
}

// Rotate etcd CA and certificates in all of the master nodes.
func (rcc *rotateCertsCmd) rotateEtcd() error {
	caPrivateKeyCmd := "echo " + rcc.containerService.Properties.CertificateProfile.CaPrivateKey + " | base64 --decode > /etc/kubernetes/certs/ca.key"
	caCertificateCmd := "echo " + rcc.containerService.Properties.CertificateProfile.CaCertificate + " | base64 --decode > /etc/kubernetes/certs/ca.crt"
	etcdServerPrivateKeyCmd := "echo " + rcc.containerService.Properties.CertificateProfile.EtcdServerPrivateKey + " | base64 --decode > /etc/kubernetes/certs/etcdserver.key"
	etcdServerCertificateCmd := "echo " + rcc.containerService.Properties.CertificateProfile.EtcdServerCertificate + " | base64 --decode > /etc/kubernetes/certs/etcdserver.crt"
	etcdClientPrivateKeyCmd := "echo " + rcc.containerService.Properties.CertificateProfile.EtcdClientPrivateKey + " | base64 --decode > /etc/kubernetes/certs/etcdclient.key"
	etcdClientCertificateCmd := "echo " + rcc.containerService.Properties.CertificateProfile.EtcdClientCertificate + " | base64 --decode > /etc/kubernetes/certs/etcdclient.crt"

	sshConfig := &ssh.ClientConfig{
		User: "azureuser",
		Auth: []ssh.AuthMethod{
			publicKeyFile(rcc.sshFilepath),
		},
	}

	for i, host := range rcc.masterNodes {
		etcdPeerPrivateKeyCmd := "echo " + rcc.containerService.Properties.CertificateProfile.EtcdPeerPrivateKeys[i] + " | base64 --decode > /etc/kubernetes/certs/etcdpeer" + strconv.Itoa(i) + ".key"
		etcdPeerCertificateCmd := "echo " + rcc.containerService.Properties.CertificateProfile.EtcdPeerCertificates[i] + " | base64 --decode > /etc/kubernetes/certs/etcdpeer" + strconv.Itoa(i) + ".crt"

		for _, cmd := range []string{caPrivateKeyCmd, caCertificateCmd, etcdServerPrivateKeyCmd, etcdServerCertificateCmd, etcdClientPrivateKeyCmd, etcdClientCertificateCmd, etcdPeerPrivateKeyCmd, etcdPeerCertificateCmd} {
			out, err := executeCmd(cmd, rcc.masterFQDN, host.Name, "22", sshConfig)
			if err != nil {
				log.Printf("Command %s output: %s\n", cmd, out)
				return errors.Wrap(err, "failed replacing certificate file")
			}
		}

		out, err := executeCmd("sudo systemctl restart etcd", rcc.masterFQDN, host.Name, "22", sshConfig)
		if err != nil {
			log.Printf("Command `sudo systemctl restart etcd` output: %s\n", out)
			return errors.Wrap(err, "failed to restart etcd")
		}
	}
	return nil
}

// From the first master node, rotate apiserver certificates in the nodes.
func (rcc *rotateCertsCmd) rotateApiserver() error {
	caCertificateCmd := "echo " + rcc.containerService.Properties.CertificateProfile.CaCertificate + " | base64 --decode > /etc/kubernetes/certs/ca.crt"
	apiServerPrivateKeyCmd := "echo " + rcc.containerService.Properties.CertificateProfile.APIServerPrivateKey + " | base64 --decode > /etc/kubernetes/certs/apiserver.key"
	apiServerCertificateCmd := "echo " + rcc.containerService.Properties.CertificateProfile.APIServerCertificate + " | base64 --decode > /etc/kubernetes/certs/apiserver.crt"

	sshConfig := &ssh.ClientConfig{
		User: "azureuser",
		Auth: []ssh.AuthMethod{
			publicKeyFile(rcc.sshFilepath),
		},
	}

	for _, host := range rcc.masterNodes {
		for _, cmd := range []string{apiServerPrivateKeyCmd, apiServerCertificateCmd} {
			out, err := executeCmd(cmd, rcc.masterFQDN, host.Name, "22", sshConfig)
			if err != nil {
				log.Printf("Command %s output: %s\n", cmd, out)
				return errors.Wrap(err, "failed replacing certificate file")
			}
		}
	}

	for _, host := range rcc.agentNodes {
		for _, cmd := range []string{caCertificateCmd, apiServerCertificateCmd} {
			out, err := executeCmd(cmd, rcc.masterFQDN, host.Name, "22", sshConfig)
			if err != nil {
				log.Printf("Command %s output: %s\n", cmd, out)
				return errors.Wrap(err, "failed replacing certificate file")
			}
		}
	}
	return nil
}

func (rcc *rotateCertsCmd) rotateKubelet() error {
	clientCertificateCmd := "echo " + rcc.containerService.Properties.CertificateProfile.ClientCertificate + " | base64 --decode > /etc/kubernetes/certs/client.crt"
	clientPrivateKeyCmd := "echo " + rcc.containerService.Properties.CertificateProfile.ClientPrivateKey + " | base64 --decode > /etc/kubernetes/certs/client.key"

	sshConfig := &ssh.ClientConfig{
		User: "azureuser",
		Auth: []ssh.AuthMethod{
			publicKeyFile(rcc.sshFilepath),
		},
	}

	for _, host := range append(rcc.masterNodes, rcc.agentNodes...) {
		for _, cmd := range []string{clientCertificateCmd, clientPrivateKeyCmd} {
			out, err := executeCmd(cmd, rcc.masterFQDN, host.Name, "22", sshConfig)
			if err != nil {
				log.Printf("Command %s output: %s\n", cmd, out)
				return errors.Wrap(err, "failed replacing certificate file")
			}
		}
	}
	return nil
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
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", masterFQDN, port), config)
	if err != nil {
		return "", errors.Wrap(err, "error starting SSH client connection")
	}
	session, err := conn.NewSession()
	if err != nil {
		return "", errors.Wrap(err, "error opening SSH session")
	}
	defer session.Close()

	err = agent.RequestAgentForwarding(session)
	if err != nil {
		return "", errors.Wrap(err, "error requesting agent forwarding")
	}
	err = agent.ForwardToRemote(conn, hostname)
	if err != nil {
		return "", errors.Wrap(err, "error forwarding to remote")
	}

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run(command)

	return fmt.Sprintf("%s -> %s", hostname, stdoutBuf.String()), nil
}
