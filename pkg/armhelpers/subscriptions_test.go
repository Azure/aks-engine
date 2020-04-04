// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"
	"testing"
)

func TestSubscriptionsInterface(t *testing.T) {
	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}
	mc.RegisterLogin()
	mc.RegisterListLocations()

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

	locations, err := azureClient.ListLocations(context.Background())
	if err != nil {
		t.Error(err)
	}
	if locations == nil || len(*locations) == 0 {
		t.Fatalf("expected locations not to be empty")
	}
}
