// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/spf13/cobra"
)

const (
	rotateCertsName             = "rotate-certs"
	rotateCertsShortDescription = "Rotate certificates on an existing Kubernetes cluster"
	rotateCertsLongDescription  = "Rotate certificates on an existing Kubernetes cluster"
)

type rotateCertsCmd struct {
	// user input
	resourceGroupName  string
	sshFilepath        string
	kubeconfigFilepath string
	masterFQDN         string
	location           string
	apimodelFilepath   string

	// derived
}

func newRotateCertsCmd() *cobra.Command {
	rcc := getRotateCertsCmd{}

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
	f.StringVarP(&rcc.resourceGroupName, "apimodel-filepath", "g", "", "the resource group where the cluster is deployed (required)")

	return command
}

func (rcc *rotateCertsCmd) run(cmd *cobra.Command, args []string) error {
	// Use the exisiting certificate generation code to generate new certificates.

	apiServerPair, clientPair, kubeConfigPair, etcdServerPair, etcdClientPair, etcdPeerPairs, err := helpers.CreatePki(masterExtraFQDNs, ips, DefaultKubernetesClusterDomain, caPair, p.MasterProfile.Count)

	// Rotate etcd CA and certificates in all of the master nodes.

	// From the first master node, rotate apiserver certificates in the nodes.

	// Update the kubeconfig and rotate the kubelet certificates.

	return nil
}
