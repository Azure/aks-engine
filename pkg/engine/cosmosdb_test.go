// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreateCosmosDB(t *testing.T) {
	db := createCosmosDBAccount()
	expected := map[string]interface{}{
		"apiVersion": "[variables('apiVersionCosmos')]",
		"name":       "Microsoft.DocumentDB/databaseAccounts",
		"location":   "[resourceGroup().location]",
		"kind":       "GlobalDocumentDB",
		"properties": map[string]interface{}{
			"consistencyPolicy": map[string]interface{}{
				"defaultConsistencyLevel": "BoundedStaleness",
				"maxStalenessPrefix":      100,
				"maxIntervalInSeconds":    5,
			},
			"databaseAccountOfferType": "Standard",
			"capabilities": []map[string]string{
				{
					"name": "EnableEtcd",
				},
			},
			"locations": []map[string]interface{}{
				{
					"locationName":     "[resourceGroup().location]",
					"failoverPriority": 0,
				},
				{
					"locationName":     "[resourceGroup().location]",
					"failoverPriority": 1,
				},
			},
		},
		"tags": map[string]string{
			"defaultExperience": "Etcd",
		},
		"consistencyPolicy": map[string]interface{}{
			"defaultConsistencyLevel": "BoundedStaleness",
			"maxIntervalInSeconds":    5,
			"maxStalenessPrefix":      100,
		},
		"primaryClientCertificatePemBytes": "[variables('cosmosDBCertb64')]",
	}

	if diff := cmp.Diff(expected, db); diff != "" {
		t.Errorf("unexpected error while comparing CosmosEtcd ARM resources: %s", diff)
	}

}
