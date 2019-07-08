// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package azurestack

import (
	"context"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestDeployTemplate(t *testing.T) {
	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}

	mc.RegisterLogin()
	mc.RegisterDeployTemplate()
	mc.RegisterDeployOperationSuccess()

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

	_, err = azureClient.DeployTemplate(context.Background(), resourceGroup, deploymentName, map[string]interface{}{}, map[string]interface{}{})
	if err != nil {
		t.Error(err)
	}
}

func TestDeployTemplateSync(t *testing.T) {
	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}

	mc.RegisterLogin()
	mc.RegisterDeployTemplate()
	mc.RegisterDeployOperationFailure()

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

	logger := log.NewEntry(log.New())
	err = DeployTemplateSync(azureClient, logger, resourceGroup, deploymentName, map[string]interface{}{}, map[string]interface{}{})
	if err == nil {
		t.Error("err should not be nil")
	}
}
