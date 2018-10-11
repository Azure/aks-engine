package v20170930

// OSType represents OS types of agents
type OSType string

// OrchestratorProfile contains orchestrator properties:
//  - type: kubernetes, etc.
//  - release: major and minor version numbers
//  - version: major, minor, and patch version numbers
type OrchestratorProfile struct {
	OrchestratorType    string `json:"orchestratorType,omitempty"`
	OrchestratorVersion string `json:"orchestratorVersion"`
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

// OrchestratorVersionProfileListProperties contains properties of OrchestratorVersionProfileList
type OrchestratorVersionProfileListProperties struct {
	Orchestrators []*OrchestratorVersionProfile `json:"orchestrators"`
}

// OrchestratorVersionProfileList contains list of version profiles for supported orchestrators
type OrchestratorVersionProfileList struct {
	ID         string                                   `json:"id,omitempty"`
	Name       string                                   `json:"name,omitempty"`
	Type       string                                   `json:"type,omitempty"`
	Properties OrchestratorVersionProfileListProperties `json:"properties"`
}
