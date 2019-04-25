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
	mc.Activate()
	defer mc.DeactivateAndReset()
	mc.RegisterLogin()
	env := mc.GetEnvironment()
	azureClient, err := NewAzureClientWithClientSecret(env, subscriptionID, "clientID", "secret")
	if err != nil {
		t.Fatalf("can not get client %s", err)
	}

	mc.RegisterDeployTemplate()

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
	mc.Activate()
	defer mc.DeactivateAndReset()
	mc.RegisterLogin()
	env := mc.GetEnvironment()
	azureClient, err := NewAzureClientWithClientSecret(env, subscriptionID, "clientID", "secret")
	if err != nil {
		t.Fatalf("can not get client %s", err)
	}
	mc.RegisterDeployTemplateSync()
	logger := log.NewEntry(log.New())
	err = DeployTemplateSync(azureClient, logger, resourceGroup, deploymentName, map[string]interface{}{}, map[string]interface{}{})
	if err == nil {
		t.Error("err should be be nil")
	}
}
