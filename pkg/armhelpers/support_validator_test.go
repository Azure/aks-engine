// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"
	"testing"

	"github.com/Azure/aks-engine/pkg/api"
)

func TestValidateRequiredImages(t *testing.T) {
	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}

	mc.RegisterLogin()
	mc.RegisterVMImageFetcherInterface()

	err = mc.Activate()
	if err != nil {
		t.Fatalf("failed to activate HttpMockClient - %s", err)
	}
	defer mc.DeactivateAndReset()

	env := mc.GetEnvironment()
	azureClient, err := NewAzureClientWithClientSecret(env, subscriptionID, "clientID", "secret")
	if err != nil {
		t.Fatalf("can not get client %s", err)
	}

	testProperties := api.Properties{}
	masterProfile := api.MasterProfile{}
	masterProfile.Distro = api.AKSUbuntu1604

	profile := api.AgentPoolProfile{
		OSType: api.Linux,
		Distro: api.AKSUbuntu1604,
	}

	winprofile := api.AgentPoolProfile{
		OSType: api.Windows,
	}

	agentProfiles := []*api.AgentPoolProfile{}
	agentProfiles = append(agentProfiles, &profile)
	agentProfiles = append(agentProfiles, &winprofile)

	testProperties.AgentPoolProfiles = agentProfiles
	testProperties.MasterProfile = &masterProfile

	if err := ValidateRequiredImages(context.Background(), location, &testProperties, azureClient); err != nil {
		t.Fatalf("can not validate required images %s", err)
	}
}

func TestValidateRequiredImagesMissingImageCase(t *testing.T) {

	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}

	mc.RegisterLogin()
	mc.RegisterVMImageFetcherInterface()

	err = mc.Activate()
	if err != nil {
		t.Fatalf("failed to activate HttpMockClient - %s", err)
	}
	defer mc.DeactivateAndReset()

	env := mc.GetEnvironment()
	azureClient, err := NewAzureClientWithClientSecret(env, subscriptionID, "clientID", "secret")
	if err != nil {
		t.Fatalf("can not get client %s", err)
	}

	testProperties := api.Properties{}
	masterProfile := api.MasterProfile{}

	masterProfile.Distro = api.AKSUbuntu1804

	profile := api.AgentPoolProfile{
		OSType: api.Linux,
		Distro: api.AKSUbuntu1804,
	}

	agentProfiles := []*api.AgentPoolProfile{}
	agentProfiles = append(agentProfiles, &profile)

	testProperties.AgentPoolProfiles = agentProfiles
	testProperties.MasterProfile = &masterProfile

	if err := ValidateRequiredImages(context.Background(), location, &testProperties, azureClient); err == nil {
		t.Fatal("could not fail fast for missing images")
	}

}
