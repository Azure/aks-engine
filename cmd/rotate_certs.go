// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"os"
	"os/exec"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/engine"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/leonelquinteros/gotext"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const (
	rotateCertsName             = "rotate-certs"
	rotateCertsShortDescription = "Rotate certificates on an existing Kubernetes cluster"
	rotateCertsLongDescription  = "Rotate certificates on an existing Kubernetes cluster"
)

type rotateCertsCmd struct {
	// user input
	resourceGroupName string
	sshFilepath       string
	kubeconfigPath    string
	masterFQDN        string
	location          string
	apiModelPath      string

	// derived
	containerService *api.ContainerService
	apiVersion       string
	locale           *gotext.Locale
}

func newRotateCertsCmd() *cobra.Command {
	rcc := rotateCertsCmd{}

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
	f.StringVarP(&rcc.kubeconfigPath, "kubeconfig", "", "", "the filepath of the cluster's kubeconfig (required)")
	f.StringVarP(&rcc.apiModelPath, "apimodel", "", "", "the filepath of the cluster's apimodel (defaults to _output)")
	f.StringVarP(&rcc.sshFilepath, "ssh", "", "", "the filepath of a valid private ssh key to access the cluster's nodes (defaults to _output)")

	return command
}

func (rcc *rotateCertsCmd) run(cmd *cobra.Command, args []string) error {

	var err error

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

	// reset the certificateProfile and use the exisiting certificate generation code to generate new certificates.

	rcc.containerService.Properties.CertificateProfile = &api.CertificateProfile{}
	certsGenerated, _, err := rcc.containerService.SetDefaultCerts()
	if !certsGenerated || err != nil {
		return errors.Wrap(err, "error generating new certificates")
	}

	rcc.rotateEtcd()

	rcc.rotateApiserver()

	// Update the kubeconfig and rotate the kubelet certificates.
	kubeConfig, err := engine.GenerateKubeConfig(rcc.containerService.Properties, rcc.location)
	if err != nil {
		return errors.Wrap(err, "error generating kubeconfig")
	}

	return nil

	// TODO: save kubeconfig

	// TODO: save apimodel certificateProfile?
}

// Rotate etcd CA and certificates in all of the master nodes.
func (rcc *rotateCertsCmd) rotateEtcd() {
	masterCount := rcc.containerService.Properties.MasterProfile.Count
	// TODO for every master node
	caPrivateKeyPath := "/etc/kubernetes/certs/ca.key"
	caCertificatePath := "/etc/kubernetes/certs/ca.crt"
	etcdServerPrivateKeyPath := "/etc/kubernetes/certs/etcdserver.key"
	etcdServerCertificatePath := "/etc/kubernetes/certs/etcdserver.crt"
	etcdClientPrivateKeyPath := "/etc/kubernetes/certs/etcdclient.key"
	etcdClientCertificatePath := "/etc/kubernetes/certs/etcdclient.crt"
	etcdPeerPrivateKeyPath := "/etc/kubernetes/certs/etcdpeer" + "" + ".key"
	etcdPeerCertificatePath := "/etc/kubernetes/certs/etcdpeer" + "" + ".crt"

}

// From the first master node, rotate apiserver certificates in the nodes.
func (rcc *rotateCertsCmd) rotateApiserver() {
	apiServerPrivateKeyPath := "/etc/kubernetes/certs/apiserver.key"
	apiServerCertificatePath := "/etc/kubernetes/certs/apiserver.crt"
}

func (rcc *rotateCertsCmd) rotateKubelet() {
	kubeletCertificatePath := "/etc/kubernetes/certs/client.crt"
	kubeletPrivateKeyPath := "/etc/kubernetes/certs/client.key"
}

func executeCommandOnNode(port, privateKeyPath, master, command string) {
	cmd := exec.Command("ssh", "-i", privateKeyPath, "-p", port, "-o", "ConnectTimeout=10", "-o", "StrictHostKeyChecking=no", "-o", "UserKnownHostsFile=/dev/null", master, command)
}
