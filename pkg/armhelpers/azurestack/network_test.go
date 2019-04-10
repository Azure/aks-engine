// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package azurestack

import (
	"context"
	"testing"
)

func TestDeleteNetworkInterface(t *testing.T) {
	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}
	mc.Activate()
	defer mc.DeactivateAndReset()
	mc.RegisterLogin()
	env := mc.GetEnvironment()
	azureClient, err := NewAzureClientWithClientSecret(env, subscriptionID, "clientID", "secret")
	if err != nil {
		t.Fatalf("can not get client %s", err)
	}

	mc.RegisterDeleteNetworkInterface()
	err = azureClient.DeleteNetworkInterface(context.Background(), resourceGroup, virtualNicName)
	if err != nil {
		t.Error(err)
	}
}
