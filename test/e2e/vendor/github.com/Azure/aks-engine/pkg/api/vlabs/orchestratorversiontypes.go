// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package vlabs

// PoolUpgradeProfile contains pool properties:
//  - orchestrator type and version
//  - pool name (for agent pool)
//  - OS type of the VMs in the pool
//  - list of applicable upgrades
type PoolUpgradeProfile struct {
	OrchestratorProfile
	Name     string                 `json:"name,omitempty"`
	OSType   OSType                 `json:"osType,omitempty"`
	Upgrades []*OrchestratorProfile `json:"upgrades,omitempty"`
}

// UpgradeProfile contains cluster properties:
//  - orchestrator type and version for the cluster
//  - list of pool profiles, constituting the cluster
type UpgradeProfile struct {
	ControlPlaneProfile *PoolUpgradeProfile   `json:"controlPlaneProfile"`
	AgentPoolProfiles   []*PoolUpgradeProfile `json:"agentPoolProfiles"`
}

// OrchestratorVersionProfile contains information of a supported orchestrator version:
//  - orchestrator type and version
//  - whether this orchestrator version is deployed by default if orchestrator release is not specified
//  - list of available upgrades for this orchestrator version
type OrchestratorVersionProfile struct {
	OrchestratorProfile
	Default  bool                   `json:"default,omitempty"`
	Upgrades []*OrchestratorProfile `json:"upgrades,omitempty"`
}

// OrchestratorVersionProfileList contains list of version profiles for supported orchestrators
type OrchestratorVersionProfileList struct {
	Orchestrators []*OrchestratorVersionProfile `json:"orchestrators"`
}
