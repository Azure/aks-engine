package v20170131

const (
	// APIVersion is the version of this API
	APIVersion = "2017-01-31"
)

// the orchestrators supported by 2017-01-31
const (
	// Kubernetes is the string constant for the Kubernetes orchestrator type
	Kubernetes string = "Kubernetes"
)

const (
	// Windows string constant for VMs
	Windows OSType = "Windows"
	// Linux string constant for VMs
	Linux OSType = "Linux"
)

// validation values
const (
	// MinAgentCount are the minimum number of agents
	MinAgentCount = 1
	// MaxAgentCount are the maximum number of agents
	MaxAgentCount = 100
)
