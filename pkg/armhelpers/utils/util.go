// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package utils

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-04-01/compute"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	// TODO: merge with the RP code
	k8sLinuxVMNamingFormat         = "^[0-9a-zA-Z]{3}-([0-9a-fA-F]{8})-(.+)-([0-9]+)$"
	k8sLinuxVMAgentPoolNameIndex   = 2
	k8sLinuxVMAgentClusterIDIndex  = 1
	k8sLinuxVMAgentIndexArrayIndex = 3

	// here there are 2 capture groups
	//  the first is the agent pool name, which can contain -s
	//  the second group is the Cluster ID for the cluster
	vmssNamingFormat       = "^[0-9a-zA-Z]+-([0-9a-fA-F]{8})-(.+)-vmss$"
	vmssAgentPoolNameIndex = 2
	vmssClusterIDIndex     = 1

	k8sWindowsOldVMNamingFormat = "^([a-fA-F0-9]{5})([0-9a-zA-Z]{3})([9])([a-zA-Z0-9]{3,5})$"
	k8sWindowsVMNamingFormat    = "^([a-fA-F0-9]{4})([0-9a-zA-Z]{3})([0-9]{3,8})$"
)

var vmnameLinuxRegexp *regexp.Regexp
var vmssnameRegexp *regexp.Regexp
var vmnameWindowsRegexp *regexp.Regexp
var oldvmnameWindowsRegexp *regexp.Regexp

func init() {
	vmnameLinuxRegexp = regexp.MustCompile(k8sLinuxVMNamingFormat)
	vmnameWindowsRegexp = regexp.MustCompile(k8sWindowsVMNamingFormat)
	oldvmnameWindowsRegexp = regexp.MustCompile(k8sWindowsOldVMNamingFormat)

	vmssnameRegexp = regexp.MustCompile(vmssNamingFormat)
}

// ResourceName returns the last segment (the resource name) for the specified resource identifier.
func ResourceName(ID string) (string, error) {
	parts := strings.Split(ID, "/")
	name := parts[len(parts)-1]
	if len(name) == 0 {
		return "", errors.Errorf("resource name was missing from identifier")
	}

	return name, nil
}

// SplitBlobURI returns a decomposed blob URI parts: accountName, containerName, blobName.
func SplitBlobURI(URI string) (string, string, string, error) {
	uri, err := url.Parse(URI)
	if err != nil {
		return "", "", "", err
	}

	accountName := strings.Split(uri.Host, ".")[0]
	urlParts := strings.Split(uri.Path, "/")

	containerName := urlParts[1]
	blobPath := strings.Join(urlParts[2:], "/")

	return accountName, containerName, blobPath, nil
}

// K8sLinuxVMNameParts returns parts of Linux VM name e.g: k8s-11290731-agentpool1-0
func K8sLinuxVMNameParts(props *api.Properties, vmName string) (poolIdentifier, nameSuffix string, agentIndex int, err error) {
	pattern := "^" + props.FormatResourceName("(?P<pool>.*)", "", "") + ".*-(?P<index>\\d+)$"
	re := regexp.MustCompile(pattern)
	vmNameParts := re.FindStringSubmatch(vmName)

	if len(vmNameParts) != 3 {
		return "", "", -1, errors.Wrap(err, "Error parsing VM Name")
	}

	result := make(map[string]string)

	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = vmNameParts[i]
		}
	}

	vmIndex, err := strconv.Atoi(result["index"])

	if err != nil {
		return "", "", -1, errors.Wrap(err, "Error parsing VM Name")
	}

	return result["pool"], props.ClusterID, vmIndex, nil
}

// VmssNameParts returns parts of Linux VM name e.g: k8s-11129073-agentpool1-0
func VmssNameParts(props *api.Properties, vmssName string) (poolIdentifier, nameSuffix string, err error) {
	pattern := "^" + props.FormatResourceName("(?P<pool>.*)", "vmss", "") + "$"
	re := regexp.MustCompile(pattern)
	vmssNameParts := re.FindStringSubmatch(vmssName)

	if len(vmssNameParts) != 2 {
		return "", "", errors.Wrap(err, "Error parsing VMSS Name")
	}

	result := make(map[string]string)

	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = vmssNameParts[i]
		}
	}

	return result["pool"], props.ClusterID, nil
}

// WindowsVMNameParts returns parts of Windows VM name
func WindowsVMNameParts(props *api.Properties, vmName string) (poolPrefix string, orch string, poolIndex int, agentIndex int, err error) {
	var poolInfo string
	vmNameParts := oldvmnameWindowsRegexp.FindStringSubmatch(vmName)
	if len(vmNameParts) != 5 {
		vmNameParts = vmnameWindowsRegexp.FindStringSubmatch(vmName)
		if len(vmNameParts) != 4 {
			return "", "", -1, -1, errors.New("resource name was missing from identifier")
		}
		poolInfo = vmNameParts[3]
	} else {
		poolInfo = vmNameParts[4]
	}

	poolPrefix = vmNameParts[1]
	orch = vmNameParts[2]

	poolIndex, err = strconv.Atoi(poolInfo[:2])
	if err != nil {
		return "", "", -1, -1, errors.Wrap(err, "Error parsing VM Name")
	}
	agentIndex, err = strconv.Atoi(poolInfo[2:])
	if err != nil {
		return "", "", -1, -1, errors.Wrap(err, "Error parsing VM Name")
	}

	return poolPrefix, orch, poolIndex, agentIndex, nil
}

// GetVMNameIndex return VM index of a node in the Kubernetes cluster
func GetVMNameIndex(props *api.Properties, osType compute.OperatingSystemTypes, vmName string) (int, error) {
	var agentIndex int
	var err error
	if osType == compute.Linux {
		_, _, agentIndex, err = K8sLinuxVMNameParts(props, vmName)
		if err != nil {
			log.Errorln(err)
			return 0, err
		}
	} else if osType == compute.Windows {
		_, _, _, agentIndex, err = WindowsVMNameParts(props, vmName)
		if err != nil {
			log.Errorln(err)
			return 0, err
		}
	}

	return agentIndex, nil
}

// GetK8sVMName reconstructs the VM name
func GetK8sVMName(p *api.Properties, agentPoolProfile *api.AgentPoolProfile, agentIndex int) (string, error) {

	vmPrefix := p.GetAgentVMPrefix(agentPoolProfile)
	if vmPrefix != "" {
		return vmPrefix + strconv.Itoa(agentIndex), nil
	}

	return "", errors.Errorf("Failed to reconstruct VM Name")
}
