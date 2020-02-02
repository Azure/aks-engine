//+build test
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package node

import (
	"context"
	"encoding/json"
	"log"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	"github.com/pkg/errors"
)

const (
	//ServerVersion is used to parse out the version of the API running
	ServerVersion = `(Server Version:\s)+(.*)`
)

// Node represents the kubernetes Node Resource
type Node struct {
	Status   Status   `json:"status"`
	Metadata Metadata `json:"metadata"`
	Spec     Spec     `json:"spec"`
}

// Metadata contains things like name and created at
type Metadata struct {
	Name        string            `json:"name"`
	CreatedAt   time.Time         `json:"creationTimestamp"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

// Spec contains things like taints
type Spec struct {
	Taints        []Taint `json:"taints"`
	Unschedulable bool    `json:"unschedulable"`
}

// Taint defines a Node Taint
type Taint struct {
	Effect string `json:"effect"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}

// Status parses information from the status key
type Status struct {
	NodeInfo      Info        `json:"nodeInfo"`
	NodeAddresses []Address   `json:"addresses"`
	Conditions    []Condition `json:"conditions"`
}

// Address contains an address and a type
type Address struct {
	Address string `json:"address"`
	Type    string `json:"type"`
}

// Info contains node information like what version the kubelet is running
type Info struct {
	ContainerRuntimeVersion string `json:"containerRuntimeVersion"`
	KubeProxyVersion        string `json:"kubeProxyVersion"`
	KubeletVersion          string `json:"kubeletVersion"`
	OperatingSystem         string `json:"operatingSystem"`
	OSImage                 string `json:"osImage"`
}

// Condition contains various status information
type Condition struct {
	LastHeartbeatTime  time.Time `json:"lastHeartbeatTime"`
	LastTransitionTime time.Time `json:"lastTransitionTime"`
	Message            string    `json:"message"`
	Reason             string    `json:"reason"`
	Status             string    `json:"status"`
	Type               string    `json:"type"`
}

// List is used to parse out Nodes from a list
type List struct {
	Nodes []Node `json:"items"`
}

// GetNodesResult is the result type for GetAllByPrefixAsync
type GetNodesResult struct {
	Nodes []Node
	Err   error
}

// GetNodesAsync wraps Get with a struct response for goroutine + channel usage
func GetNodesAsync() GetNodesResult {
	list, err := Get()
	if list == nil {
		list = &List{
			Nodes: []Node{},
		}
	}
	return GetNodesResult{
		Nodes: list.Nodes,
		Err:   err,
	}
}

// GetReadyNodesAsync wraps Get with a struct response for goroutine + channel usage
func GetReadyNodesAsync() GetNodesResult {
	list, err := GetReady()
	if list == nil {
		list = &List{
			Nodes: []Node{},
		}
	}
	return GetNodesResult{
		Nodes: list.Nodes,
		Err:   err,
	}
}

// GetByRegexAsync wraps GetByRegex with a struct response for goroutine + channel usage
func GetByRegexAsync(regex string) GetNodesResult {
	nodes, err := GetByRegex(regex)
	if nodes == nil {
		nodes = []Node{}
	}
	return GetNodesResult{
		Nodes: nodes,
		Err:   err,
	}
}

// IsReady returns if the node is in a Ready state
func (n *Node) IsReady() bool {
	if n.Spec.Unschedulable {
		return false
	}
	for _, condition := range n.Status.Conditions {
		if condition.Type == "Ready" && condition.Status == "True" {
			return true
		}
	}
	return false
}

// IsLinux checks for a Linux node
func (n *Node) IsLinux() bool {
	return n.Status.NodeInfo.OperatingSystem == "linux"
}

// IsWindows checks for a Windows node
func (n *Node) IsWindows() bool {
	return n.Status.NodeInfo.OperatingSystem == "windows"
}

// IsUbuntu checks for an Ubuntu-backed node
func (n *Node) IsUbuntu() bool {
	if n.IsLinux() {
		return strings.Contains(strings.ToLower(n.Status.NodeInfo.OSImage), "ubuntu")
	}
	return false
}

// HasSubstring determines if a node name matches includes the passed in substring
func (n *Node) HasSubstring(substrings []string) bool {
	for _, substring := range substrings {
		if strings.Contains(strings.ToLower(n.Metadata.Name), substring) {
			return true
		}
	}
	return false
}

// Version returns the version of the kubelet on the node
func (n *Node) Version() string {
	return n.Status.NodeInfo.KubeletVersion
}

// DescribeNodes describes all nodes
func DescribeNodes() {
	list, err := Get()
	if err != nil {
		log.Printf("Unable to get nodes: %s", err)
	}
	if list != nil {
		for _, node := range list.Nodes {
			err := node.Describe()
			if err != nil {
				log.Printf("Unable to describe node %s: %s", node.Metadata.Name, err)
			}
		}
	}
}

// Describe will describe a node resource
func (n *Node) Describe() error {
	var commandTimeout time.Duration
	cmd := exec.Command("k", "describe", "node", n.Metadata.Name)
	out, err := util.RunAndLogCommand(cmd, commandTimeout)
	log.Printf("\n%s\n", string(out))
	return err
}

// AreAllReady returns if all nodes are ready
func AreAllReady() bool {
	list, _ := Get()
	if list != nil {
		for _, node := range list.Nodes {
			if !node.IsReady() {
				return false
			}
		}
	}
	return true
}

// AreNNodesReady returns a bool depending on cluster state
func AreNNodesReady(nodeCount int) bool {
	if nodeCount == -1 {
		return AreAllReady()
	}
	list, _ := Get()
	var ready int
	if list != nil {
		for _, node := range list.Nodes {
			nodeReady := node.IsReady()
			if !nodeReady {
				return false
			}
			ready++
		}
	}
	if ready == nodeCount {
		return true
	}
	return false
}

// AreMinNodesReady returns if the minimum nodes ready count is met
func AreMinNodesReady(nodeCount int) bool {
	if nodeCount == -1 {
		return AreAllReady()
	}
	list, _ := Get()
	var ready int
	if list != nil {
		for _, node := range list.Nodes {
			nodeReady := node.IsReady()
			if !nodeReady {
				return false
			}
			ready++
		}
	}
	if ready >= nodeCount {
		return true
	}
	return false
}

// AreMaxNodesReady returns if nodes ready count is <= a maximum number
func AreMaxNodesReady(nodeCount int) bool {
	list, _ := Get()
	var ready int
	if list != nil {
		for _, node := range list.Nodes {
			nodeReady := node.IsReady()
			if !nodeReady {
				return false
			}
			ready++
		}
	}
	if ready <= nodeCount {
		return true
	}
	return false
}

// WaitOnReady will block until all nodes are in ready state
func WaitOnReady(nodeCount int, sleep, timeout time.Duration) bool {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan bool)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- AreNNodesReady(nodeCount)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case ready := <-ch:
			if ready {
				return ready
			}
		case <-ctx.Done():
			DescribeNodes()
			return false
		}
	}
}

// WaitOnReadyMin will block until the minimum nodes ready count is met
func WaitOnReadyMin(nodeCount int, sleep, timeout time.Duration) bool {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan bool)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- AreMinNodesReady(nodeCount)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case ready := <-ch:
			if ready {
				return ready
			}
		case <-ctx.Done():
			DescribeNodes()
			return false
		}
	}
}

// WaitOnReadyMax will block until nodes ready count is <= a maximum number
func WaitOnReadyMax(nodeCount int, sleep, timeout time.Duration) bool {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan bool)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- AreMaxNodesReady(nodeCount)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case ready := <-ch:
			if ready {
				return ready
			}
		case <-ctx.Done():
			DescribeNodes()
			return false
		}
	}
}

// Get returns the current nodes for a given kubeconfig
func Get() (*List, error) {
	cmd := exec.Command("k", "get", "nodes", "-o", "json")
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error trying to run 'kubectl get nodes':\n - %s", err)
		if len(string(out)) > 0 {
			log.Printf("\n - %s", string(out))
		}
		return nil, err
	}
	nl := List{}
	err = json.Unmarshal(out, &nl)
	if err != nil {
		log.Printf("Error unmarshalling nodes json:%s", err)
	}
	return &nl, nil
}

// GetReadyWithRetry gets nodes, allowing for retries
func GetReadyWithRetry(sleep, timeout time.Duration) ([]Node, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan GetNodesResult)
	var mostRecentGetReadyWithRetryError error
	var nodes []Node
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- GetReadyNodesAsync()
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentGetReadyWithRetryError = result.Err
			nodes = result.Nodes
			if mostRecentGetReadyWithRetryError == nil {
				if len(nodes) > 0 {
					return nodes, nil
				}
			}
		case <-ctx.Done():
			return nil, errors.Errorf("GetReadyWithRetry timed out: %s\n", mostRecentGetReadyWithRetryError)
		}
	}
}

// GetWithRetry gets nodes, allowing for retries
func GetWithRetry(sleep, timeout time.Duration) ([]Node, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan GetNodesResult)
	var mostRecentGetWithRetryError error
	var nodes []Node
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- GetNodesAsync()
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentGetWithRetryError = result.Err
			nodes = result.Nodes
			if mostRecentGetWithRetryError == nil {
				if len(nodes) > 0 {
					return nodes, nil
				}
			}
		case <-ctx.Done():
			return nil, errors.Errorf("GetWithRetry timed out: %s\n", mostRecentGetWithRetryError)
		}
	}
}

// GetByRegexWithRetry gets nodes that match a regular expression, allowing for retries
func GetByRegexWithRetry(regex string, sleep, timeout time.Duration) ([]Node, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan GetNodesResult)
	var mostRecentGetByRegexWithRetryError error
	var nodes []Node
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- GetByRegexAsync(regex)
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentGetByRegexWithRetryError = result.Err
			nodes = result.Nodes
			if mostRecentGetByRegexWithRetryError == nil {
				if len(nodes) > 0 {
					return nodes, nil
				}
			}
		case <-ctx.Done():
			return nil, errors.Errorf("GetByRegexWithRetry timed out: %s\n", mostRecentGetByRegexWithRetryError)
		}
	}
}

// GetReady returns the current nodes for a given kubeconfig
func GetReady() (*List, error) {
	l, err := Get()
	if err != nil {
		return nil, err
	}
	nl := &List{
		[]Node{},
	}
	for _, node := range l.Nodes {
		if node.IsReady() {
			nl.Nodes = append(nl.Nodes, node)
		} else {
			log.Printf("found an unready node!")
		}
	}
	return nl, nil
}

// GetAddressByType will return the Address object for a given Kubernetes node
func (ns *Status) GetAddressByType(t string) *Address {
	for _, a := range ns.NodeAddresses {
		if a.Type == t {
			return &a
		}
	}
	return nil
}

// GetByRegex will return a []Node of all nodes that have a name that match the regular expression
func GetByRegex(regex string) ([]Node, error) {
	list, err := Get()
	if err != nil {
		return nil, err
	}

	nodes := make([]Node, 0)
	for _, n := range list.Nodes {
		exp, err := regexp.Compile(regex)
		if err != nil {
			return nil, err
		}
		if exp.MatchString(n.Metadata.Name) {
			nodes = append(nodes, n)
		}
	}
	return nodes, nil
}

// GetByLabel will return a []Node of all nodes that have a matching label
func GetByLabel(label string) ([]Node, error) {
	list, err := Get()
	if err != nil {
		return nil, err
	}

	nodes := make([]Node, 0)
	for _, n := range list.Nodes {
		if _, ok := n.Metadata.Labels[label]; ok {
			nodes = append(nodes, n)
		}
	}
	return nodes, nil
}

// GetByAnnotations will return a []Node of all nodes that have a matching annotation
func GetByAnnotations(key, value string) ([]Node, error) {
	list, err := Get()
	if err != nil {
		return nil, err
	}

	nodes := make([]Node, 0)
	for _, n := range list.Nodes {
		if n.Metadata.Annotations[key] == value {
			nodes = append(nodes, n)
		}
	}
	return nodes, nil
}

// GetByTaint will return a []Node of all nodes that have a matching taint
func GetByTaint(key, value, effect string) ([]Node, error) {
	list, err := Get()
	if err != nil {
		return nil, err
	}

	nodes := make([]Node, 0)
	for _, n := range list.Nodes {
		for _, t := range n.Spec.Taints {
			if t.Key == key && t.Value == value && t.Effect == effect {
				nodes = append(nodes, n)
			}
		}
	}
	return nodes, nil
}
