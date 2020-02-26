// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package azurestack

import (
	"context"
	"testing"
)

func TestResourceSkusInterface(t *testing.T) {
	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}
	mc.RegisterLogin()
	mc.RegisterListResourceSkus()

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

	page, err := azureClient.ListResourceSkus(context.Background(), "")
	if err != nil {
		t.Error(err)
	}
	if page == nil || len(page.Values()) == 0 {
		t.Fatalf("expected skus not to be empty")
	}
}
