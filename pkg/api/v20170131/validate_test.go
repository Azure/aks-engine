// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package v20170131

import "testing"

func Test_ServicePrincipalProfile_ValidateSecret(t *testing.T) {

	t.Run("ServicePrincipalProfile is nil should fail", func(t *testing.T) {
		t.Parallel()
		p := getK8sDefaultProperties()
		p.ServicePrincipalProfile = nil

		if err := p.Validate(); err == nil {
			t.Errorf("should error %v", err)
		}
	})

	t.Run("ServicePrincipalProfile with secret should pass", func(t *testing.T) {
		t.Parallel()
		p := getK8sDefaultProperties()

		if err := p.Validate(); err != nil {
			t.Errorf("should not error %v", err)
		}
	})

	t.Run("ServicePrincipalProfile with missing secret should pass", func(t *testing.T) {
		t.Parallel()
		p := getK8sDefaultProperties()
		p.ServicePrincipalProfile.Secret = ""

		if err := p.Validate(); err == nil {
			t.Error("error should have occurred")
		}
	})

}

func getK8sDefaultProperties() *Properties {
	return &Properties{
		OrchestratorProfile: &OrchestratorProfile{
			OrchestratorType: Kubernetes,
		},
		MasterProfile: &MasterProfile{
			Count:     1,
			DNSPrefix: "foo",
		},
		AgentPoolProfiles: []*AgentPoolProfile{
			{
				Name:   "agentpool",
				VMSize: "Standard_D2_v2",
				Count:  1,
			},
		},
		LinuxProfile: &LinuxProfile{
			AdminUsername: "azureuser",
			SSH: struct {
				PublicKeys []PublicKey `json:"publicKeys"`
			}{
				PublicKeys: []PublicKey{{
					KeyData: "publickeydata",
				}},
			},
		},
		ServicePrincipalProfile: &ServicePrincipalProfile{
			ClientID: "clientID",
			Secret:   "clientSecret",
		},
	}
}
