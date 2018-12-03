// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"github.com/Azure/aks-engine/pkg/api/osa/vlabs"
)

// ConvertVLabsOpenShiftClusterToContainerService converts from a
// vlabs.OpenShiftCluster to a ContainerService.
func ConvertVLabsOpenShiftClusterToContainerService(oc *vlabs.OpenShiftCluster) *ContainerService {
	cs := &ContainerService{
		ID:       oc.ID,
		Location: oc.Location,
		Name:     oc.Name,
		Tags:     oc.Tags,
		Type:     oc.Type,
	}

	if oc.Plan != nil {
		cs.Plan = &ResourcePurchasePlan{
			Name:          oc.Plan.Name,
			Product:       oc.Plan.Product,
			PromotionCode: oc.Plan.PromotionCode,
			Publisher:     oc.Plan.Publisher,
		}
	}

	if oc.Properties != nil {
		cs.Properties = &Properties{
			ProvisioningState: ProvisioningState(oc.Properties.ProvisioningState),
			OrchestratorProfile: &OrchestratorProfile{
				OrchestratorVersion: oc.Properties.OpenShiftVersion,
				OpenShiftConfig: &OpenShiftConfig{
					PublicHostname: oc.Properties.PublicHostname,
				},
			},
			MasterProfile: &MasterProfile{
				FQDN: oc.Properties.FQDN,
			},
			ServicePrincipalProfile: &ServicePrincipalProfile{
				ClientID: oc.Properties.ServicePrincipalProfile.ClientID,
				Secret:   oc.Properties.ServicePrincipalProfile.Secret,
			},
		}

		cs.Properties.OrchestratorProfile.OpenShiftConfig.RouterProfiles = make([]OpenShiftRouterProfile, len(oc.Properties.RouterProfiles))
		for i, rp := range oc.Properties.RouterProfiles {
			cs.Properties.OrchestratorProfile.OpenShiftConfig.RouterProfiles[i] = OpenShiftRouterProfile{
				Name:            rp.Name,
				PublicSubdomain: rp.PublicSubdomain,
				FQDN:            rp.FQDN,
			}
		}

		cs.Properties.AgentPoolProfiles = make([]*AgentPoolProfile, len(oc.Properties.AgentPoolProfiles))
		for i, app := range oc.Properties.AgentPoolProfiles {
			cs.Properties.AgentPoolProfiles[i] = &AgentPoolProfile{
				Name:         app.Name,
				Count:        app.Count,
				VMSize:       app.VMSize,
				OSType:       OSType(app.OSType),
				VnetSubnetID: app.VnetSubnetID,
				Role:         AgentPoolProfileRole(app.Role),
			}
		}
	}

	return cs
}
