// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package azurestack

import (
	"bytes"
	"context"

	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2017-10-01/storage"
	azStorage "github.com/Azure/azure-sdk-for-go/storage"
	"github.com/Azure/go-autorest/autorest/to"
)

// AzureStorageClient implements the StorageClient interface and wraps the Azure storage client.
type AzureStorageClient struct {
	client *azStorage.Client
}

// GetStorageClient returns an authenticated client for the specified account.
func (az *AzureClient) GetStorageClient(ctx context.Context, resourceGroup, accountName string) (armhelpers.AKSStorageClient, error) {
	keys, err := az.getStorageKeys(ctx, resourceGroup, accountName)
	if err != nil {
		return nil, err
	}

	client, err := azStorage.NewBasicClientOnSovereignCloud(accountName, to.String(keys[0].Value), az.environment)
	if err != nil {
		return nil, err
	}

	return &AzureStorageClient{
		client: &client,
	}, nil
}

func (az *AzureClient) getStorageKeys(ctx context.Context, resourceGroup, accountName string) ([]storage.AccountKey, error) {
	storageKeysResult, err := az.storageAccountsClient.ListKeys(ctx, resourceGroup, accountName)
	if err != nil {
		return nil, err
	}

	return *storageKeysResult.Keys, nil
}

// DeleteBlob deletes the specified blob
// TODO(colemick): why doesn't SDK give a way to just delete a blob by URI?
// it's what it ends up doing internally anyway...
func (as *AzureStorageClient) DeleteBlob(vhdContainer, vhdBlob string, options *azStorage.DeleteBlobOptions) error {
	containerRef := getContainerRef(as.client, vhdContainer)
	blobRef := containerRef.GetBlobReference(vhdBlob)

	return blobRef.Delete(options)
}

// CreateContainer creates the CloudBlobContainer if it does not exist
func (as *AzureStorageClient) CreateContainer(containerName string, options *azStorage.CreateContainerOptions) (bool, error) {
	containerRef := getContainerRef(as.client, containerName)
	created, err := containerRef.CreateIfNotExists(options)

	return created, err
}

// SaveBlockBlob initializes a block blob by taking the byte
func (as *AzureStorageClient) SaveBlockBlob(containerName, blobName string, b []byte, options *azStorage.PutBlobOptions) error {
	containerRef := getContainerRef(as.client, containerName)
	blobRef := containerRef.GetBlobReference(blobName)

	return blobRef.CreateBlockBlobFromReader(bytes.NewReader(b), options)
}

func getContainerRef(client *azStorage.Client, containerName string) *azStorage.Container {
	bs := client.GetBlobService()
	return bs.GetContainerReference(containerName)
}
