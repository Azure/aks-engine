// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package azurestack

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/msi/mgmt/2015-08-31-preview/msi"
	"github.com/pkg/errors"
)

// CreateUserAssignedID - Creates a user assigned msi.
func (az *AzureClient) CreateUserAssignedID(location string, resourceGroup string, userAssignedID string) (id *msi.Identity, err error) {
	errorMessage := "error azure stack does not support creating user assigned msi"
	return &msi.Identity{}, errors.New(errorMessage)
}
