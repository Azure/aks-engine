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
		"name":       "[variables('cosmosAccountName')]",
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
			},
			"name":                             "[variables('cosmosAccountName')]",
			"primaryClientCertificatePemBytes": "[variables('cosmosDBCertb64')]",
		},
		"tags": map[string]string{
			"defaultExperience": "Etcd",
		},
		"type": "Microsoft.DocumentDB/databaseAccounts",
	}

	if diff := cmp.Diff(expected, db); diff != "" {
		t.Errorf("unexpected error while comparing CosmosEtcd ARM resources: %s", diff)
	}

}
