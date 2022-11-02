// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/msi/mgmt/2015-08-31-preview/msi"
	"github.com/Azure/go-autorest/autorest/to"
	log "github.com/sirupsen/logrus"
)

// CreateUserAssignedID - Creates a user assigned msi.
func (az *AzureClient) CreateUserAssignedID(location string, resourceGroup string, userAssignedID string) (id *msi.Identity, err error) {
	idCreated, err := az.msiClient.CreateOrUpdate(context.Background(), resourceGroup, userAssignedID, msi.Identity{
		Location: to.StringPtr(location),
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Infof("Created %s in rg %s", userAssignedID, resourceGroup)
	return &idCreated, nil
}
