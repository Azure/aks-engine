// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package v20170831

const (
	// APIVersion is the version of this API
	APIVersion = "2017-08-31"
)

const (
	// Windows OSType
	Windows OSType = "Windows"
	// Linux OSType
	Linux OSType = "Linux"
)

// validation values
const (
	// MinAgentCount are the minimum number of agents
	MinAgentCount = 1
	// MaxAgentCount are the maximum number of agents
	MaxAgentCount = 100
	// MinDiskSizeGB specifies the minimum attached disk size
	MinDiskSizeGB = 1
	// MaxDiskSizeGB specifies the maximum attached disk size
	MaxDiskSizeGB = 1023
)

// storage profiles
const (
	// ManagedDisks means that the nodes use managed disks for their os and attached volumes
	ManagedDisks = "ManagedDisks"
)
