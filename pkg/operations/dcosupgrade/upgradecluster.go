// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package dcosupgrade

import (
	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// ClusterTopology contains resources of the cluster the upgrade operation
// is targeting
type ClusterTopology struct {
	DataModel          *api.ContainerService
	Location           string
	ResourceGroup      string
	CurrentDcosVersion string
	NameSuffix         string
	SSHKey             []byte
}

// UpgradeCluster upgrades a cluster with Orchestrator version X.X to version Y.Y.
// Right now upgrades are supported for Kubernetes cluster only.
type UpgradeCluster struct {
	Translator *i18n.Translator
	Logger     *logrus.Entry
	ClusterTopology
	Client armhelpers.AKSEngineClient
}

// UpgradeCluster runs the workflow to upgrade a DCOS cluster.
func (uc *UpgradeCluster) UpgradeCluster(subscriptionID uuid.UUID, resourceGroup, currentDcosVersion string,
	cs *api.ContainerService, nameSuffix string, sshKey []byte) error {
	uc.ClusterTopology = ClusterTopology{}
	uc.ResourceGroup = resourceGroup
	uc.CurrentDcosVersion = currentDcosVersion
	uc.DataModel = cs
	uc.NameSuffix = nameSuffix
	uc.SSHKey = sshKey

	upgradeVersion := uc.DataModel.Properties.OrchestratorProfile.OrchestratorVersion
	uc.Logger.Infof("Upgrading DCOS from %s to %s", uc.CurrentDcosVersion, upgradeVersion)

	if err := uc.runUpgrade(); err != nil {
		return err
	}

	uc.Logger.Infof("Cluster upgraded successfully to DCOS %s", upgradeVersion)
	return nil
}
