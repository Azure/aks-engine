// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package operations

import (
	"context"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
)

type getNodesResult struct {
	nodes []v1.Node
	err   error
}

// GetNodes is a thin wrapper around the k8s api list nodes interface
// Pass in a pool string to filter only node objects in that AKS Engine-deployed node pool
// Pass in a waitForNumNodes int to wait for an explicit target node count before returning
func GetNodes(az armhelpers.AKSEngineClient, logger *log.Entry, apiserverURL, kubeConfig string, timeout time.Duration, pool string, waitForNumNodes int) ([]v1.Node, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan getNodesResult)
	var mostRecentGetNodesErr error
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- listNodes(az, logger, apiserverURL, kubeConfig, timeout):
				time.Sleep(3 * time.Second)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentGetNodesErr = result.err
			if result.err == nil {
				var ret []v1.Node
				for _, node := range result.nodes {
					if strings.Contains(node.Name, pool) {
						ret = append(ret, node)
					}
				}
				if waitForNumNodes >= 0 {
					if len(ret) == waitForNumNodes {
						return ret, nil
					}
				} else {
					return ret, nil
				}
			}
		case <-ctx.Done():
			return nil, errors.Errorf("GetAllNodes timed out: %s\n", mostRecentGetNodesErr)
		}
	}
}

func listNodes(az armhelpers.AKSEngineClient, logger *log.Entry, apiserverURL, kubeConfig string, timeout time.Duration) getNodesResult {
	logger.Debugf("Instantiating a Kubernetes client object at apiserver %s", apiserverURL)
	client, err := az.GetKubernetesClient(apiserverURL, kubeConfig, interval, timeout)
	if err != nil {
		return getNodesResult{
			err: err,
		}
	}
	logger.Debugf("Listing Nodes at apiserver %s", apiserverURL)
	nodes, err := client.ListNodes()
	if err != nil {
		return getNodesResult{
			err: err,
		}
	}

	return getNodesResult{
		nodes: nodes.Items,
		err:   nil,
	}
}

// PrintNodes outputs nodes to stdout
func PrintNodes(nodes []v1.Node) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 4, ' ', tabwriter.FilterHTML)
	fmt.Fprintln(w, "NODE\tSTATUS\tVERSION\tOS\tKERNEL")
	for _, node := range nodes {
		nodeStatus := "NotReady"
		for _, condition := range node.Status.Conditions {
			if condition.Type == "Ready" && condition.Status == "True" {
				nodeStatus = "Ready"
			}
		}
		fmt.Fprintf(w, "%s\t", node.Name)
		fmt.Fprintf(w, "%s\t", nodeStatus)
		fmt.Fprintf(w, "%s\t", node.Status.NodeInfo.KubeletVersion)
		fmt.Fprintf(w, "%s\t", node.Status.NodeInfo.OSImage)
		fmt.Fprintf(w, "%s\n", node.Status.NodeInfo.KernelVersion)

	}
	w.Flush()
}
