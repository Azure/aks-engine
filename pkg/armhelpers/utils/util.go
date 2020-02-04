// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

// Package utils provides helper methods to assist with ARM operations.
package utils

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	// TODO: merge with the RP code
	k8sLinuxVMNamingFormat         = "^[0-9a-zA-Z]{3}-(.+)-([0-9a-fA-F]{8})-{0,2}([0-9]+)$"
	k8sLinuxVMAgentPoolNameIndex   = 1
	k8sLinuxVMAgentClusterIDIndex  = 2
	k8sLinuxVMAgentIndexArrayIndex = 3

	// here there are 2 capture groups
	//  the first is the agent pool name, which can contain -s
	//  the second group is the Cluster ID for the cluster
	vmssNamingFormat       = "^[0-9a-zA-Z]+-(.+)-([0-9a-fA-F]{8})-vmss$"
	vmssAgentPoolNameIndex = 1
	vmssClusterIDIndex     = 2

	k8sWindowsOldVMNamingFormat = "^([a-fA-F0-9]{5})([0-9a-zA-Z]{3})([9])([a-zA-Z0-9]{3,5})$"
	k8sWindowsVMNamingFormat    = "^([a-fA-F0-9]{4})([0-9a-zA-Z]{3})([0-9]{3,8})$"

	windowsVmssNamingFormatV2       = "aks([0-9a-zA-Z]+)$"
	windowsVmssAgentPoolNameIndexV2 = 1
)

var vmnameLinuxRegexp *regexp.Regexp
var vmssnameRegexp *regexp.Regexp
var vmnameWindowsRegexp *regexp.Regexp
var oldvmnameWindowsRegexp *regexp.Regexp
var windowsVmssRegexV2 *regexp.Regexp

func init() {
	vmnameLinuxRegexp = regexp.MustCompile(k8sLinuxVMNamingFormat)
	vmnameWindowsRegexp = regexp.MustCompile(k8sWindowsVMNamingFormat)
	oldvmnameWindowsRegexp = regexp.MustCompile(k8sWindowsOldVMNamingFormat)

	vmssnameRegexp = regexp.MustCompile(vmssNamingFormat)
	windowsVmssRegexV2 = regexp.MustCompile(windowsVmssNamingFormatV2)
}

// ResourceName returns the last segment (the resource name) for the specified resource identifier.
func ResourceName(id string) (string, error) {
	parts := strings.Split(id, "/")
	name := parts[len(parts)-1]
	if len(name) == 0 {
		return "", errors.Errorf("resource name was missing from identifier")
	}

	return name, nil
}

// SplitBlobURI returns a decomposed blob URI parts: accountName, containerName, blobName.
func SplitBlobURI(uri string) (string, string, string, error) {
	parsed, err := url.Parse(uri)
	if err != nil {
		return "", "", "", err
	}

	accountName := strings.Split(parsed.Host, ".")[0]
	urlParts := strings.Split(parsed.Path, "/")

	containerName := urlParts[1]
	blobPath := strings.Join(urlParts[2:], "/")

	return accountName, containerName, blobPath, nil
}

// K8sLinuxVMNameParts returns parts of Linux VM name e.g: k8s-agentpool1-11290731-0
func K8sLinuxVMNameParts(vmName string) (poolIdentifier, nameSuffix string, agentIndex int, err error) {
	vmNameParts := vmnameLinuxRegexp.FindStringSubmatch(vmName)
	if len(vmNameParts) != 4 {
		return "", "", -1, errors.Errorf("resource name was missing from identifier")
	}

	vmNum, err := strconv.Atoi(vmNameParts[k8sLinuxVMAgentIndexArrayIndex])

	if err != nil {
		return "", "", -1, errors.Wrap(err, "Error parsing VM Name")
	}

	return vmNameParts[k8sLinuxVMAgentPoolNameIndex], vmNameParts[k8sLinuxVMAgentClusterIDIndex], vmNum, nil
}

// VmssNameParts returns parts of Linux VM name e.g: k8s-agentpool1-11290731-0
func VmssNameParts(vmssName string) (poolIdentifier, nameSuffix string, err error) {
	vmssNameParts := vmssnameRegexp.FindStringSubmatch(vmssName)
	if len(vmssNameParts) != 3 {
		return "", "", errors.New("resource name was missing from identifier")
	}

	return vmssNameParts[vmssAgentPoolNameIndex], vmssNameParts[vmssClusterIDIndex], nil
}

// WindowsVmssNameParts returns the agent pool name from the VMSS name.
// For windows, the VMSS name is 'aks' concatenated with agentpool name
func WindowsVmssNameParts(vmssName string) (poolIdentifier string, err error) {
	vmssNameParts := windowsVmssRegexV2.FindStringSubmatch(vmssName)
	if len(vmssNameParts) != 2 {
		return "", errors.New("resource name was missing from identifier")
	}

	return vmssNameParts[windowsVmssAgentPoolNameIndexV2], nil
}

// WindowsVMNameParts returns parts of Windows VM name
func WindowsVMNameParts(vmName string) (poolPrefix string, orch string, poolIndex int, agentIndex int, err error) {
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
func GetVMNameIndex(osType compute.OperatingSystemTypes, vmName string) (int, error) {
	var agentIndex int
	var err error
	if osType == compute.Linux {
		_, _, agentIndex, err = K8sLinuxVMNameParts(vmName)
		if err != nil {
			log.Errorln(err)
			return 0, err
		}
		return agentIndex, nil
	}
	if osType == compute.Windows {
		_, _, _, agentIndex, err = WindowsVMNameParts(vmName)
		if err != nil {
			log.Errorln(err)
			return 0, err
		}
		return agentIndex, nil
	}

	return 0, errors.Errorf("Unknown os type. Can't determine vm index")
}

// GetK8sVMName reconstructs the VM name
func GetK8sVMName(p *api.Properties, agentPoolProfile *api.AgentPoolProfile, agentIndex int) (string, error) {
	vmPrefix := p.GetAgentVMPrefix(agentPoolProfile, p.GetAgentPoolIndexByName(agentPoolProfile.Name))
	if vmPrefix != "" {
		return vmPrefix + strconv.Itoa(agentIndex), nil
	}

	return "", errors.Errorf("Failed to reconstruct VM Name")
}
