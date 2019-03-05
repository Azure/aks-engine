// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2015-04-08/documentdb"
	"github.com/Azure/go-autorest/autorest/to"
)

func createCosmosDBAccount() DocumentDBAccountARM {
	armResource := ARMResource{
		APIVersion: "[variables('apiVersionCosmos')]",
	}

	documentDB := documentdb.DatabaseAccountCreateUpdateParameters{
		Type:     to.StringPtr("Microsoft.DocumentDB/databaseAccounts"),
		Kind:     documentdb.GlobalDocumentDB,
		Location: to.StringPtr("[resourceGroup().location]"),
		Name:     to.StringPtr("[variables('cosmosAccountName')]"),
		DatabaseAccountCreateUpdateProperties: &documentdb.DatabaseAccountCreateUpdateProperties{
			Capabilities: &[]documentdb.Capability{
				{
					Name: to.StringPtr("EnableEtcd"),
				},
			},
			ConsistencyPolicy: &documentdb.ConsistencyPolicy{
				DefaultConsistencyLevel: documentdb.BoundedStaleness,
				MaxIntervalInSeconds:    to.Int32Ptr(5),
				MaxStalenessPrefix:      to.Int64Ptr(100),
			},
			DatabaseAccountOfferType: to.StringPtr(string(documentdb.Standard)),
			Locations: &[]documentdb.Location{
				{
					FailoverPriority: to.Int32Ptr(0),
					LocationName:     to.StringPtr("[resourceGroup().location]"),
				},
				{
					FailoverPriority: to.Int32Ptr(1),
					LocationName:     to.StringPtr("[resourceGroup().location]"),
				},
			},
		},
		Tags: map[string]*string{
			"defaultExperience": to.StringPtr("Etcd"),
		},
		//TODO: Need to do something about the primaryClientCertificatePemBytes
	}

	return DocumentDBAccountARM{
		ARMResource:                           armResource,
		DatabaseAccountCreateUpdateParameters: documentDB,
	}
}
