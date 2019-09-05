// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package azurestack

import (
	"context"
	"testing"
)

func TestGetLogAnalyticsWorkspaceInfo(t *testing.T) {
	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}

	mc.RegisterLogin()
	mc.RegisterGetLogAnalyticsWorkspaceInfo()
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
	_, _, _, err = azureClient.GetLogAnalyticsWorkspaceInfo(context.Background(), subscriptionID, resourceGroup, logAnalyticsWorkspaceName)
	if err != nil {
		t.Error(err)
	}
}

func TestEnsureDefaultLogAnalyticsWorkspaceUseExisting(t *testing.T) {
	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}

	mc.RegisterLogin()
	mc.RegisterEnsureDefaultLogAnalyticsWorkspaceUseExisting()

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
	_, err = azureClient.EnsureDefaultLogAnalyticsWorkspace(context.Background(), resourceGroup, location)
	if err != nil {
		t.Error(err)
	}
}

func TestEnsureDefaultLogAnalyticsWorkspaceCreateNew(t *testing.T) {
	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}

	mc.RegisterLogin()
	mc.RegisterEnsureDefaultLogAnalyticsWorkspaceCreateNew()

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
	_, err = azureClient.EnsureDefaultLogAnalyticsWorkspace(context.Background(), resourceGroup, "westeurope")
	if err != nil {
		t.Error(err)
	}
}
