package v20170930

import (
	"strings"

	"github.com/pkg/errors"
)

// Validate implements APIObject
func (o *OrchestratorVersionProfile) Validate() error {
	switch {
	case strings.EqualFold(o.OrchestratorType, Kubernetes):
		o.OrchestratorType = Kubernetes
	default:
		return errors.Errorf("Unsupported orchestrator '%s'", o.OrchestratorType)
	}
	return nil
}

// ValidateForUpgrade validates upgrade input data
func (o *OrchestratorProfile) ValidateForUpgrade() error {
	switch o.OrchestratorType {
	case Kubernetes:
		switch o.OrchestratorVersion {
		case "1.6.13":
		case "1.7.14":
		default:
			return errors.Errorf("Upgrade to Kubernetes %s is not supported", o.OrchestratorVersion)
		}
	}
	return nil
}
