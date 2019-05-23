// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package operations

import (
	"time"

	"github.com/Azure/aks-engine/pkg/armhelpers"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
)

const listNodesRetries = 5

// GetNodes is a thin wrapper around the k8s api list nodes interface
func GetNodes(az armhelpers.AKSEngineClient, logger *log.Entry, apiserverURL, kubeConfig string, timeout time.Duration) ([]v1.Node, error) {
	//get client using kubeconfig
	client, err := az.GetKubernetesClient(apiserverURL, kubeConfig, interval, timeout)
	if err != nil {
		return nil, err
	}
	var nodes *v1.NodeList
	for i := 1; i <= listNodesRetries; i++ {
		nodes, err = client.ListNodes()
		if err != nil {
			if i == listNodesRetries {
				return nil, err
			}
			time.Sleep(5 * time.Second)
		}
	}
	return nodes.Items, nil
}
