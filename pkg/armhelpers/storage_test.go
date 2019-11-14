// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

func TestAzureStorageClient_CreateContainer(t *testing.T) {
	cases := []struct {
		name                 string
		storageClientFactory func() *MockStorageClient
		errMatcher           types.GomegaMatcher
		resMatcher           types.GomegaMatcher
	}{
		{
			name: "ShouldPassIfCreated",
			storageClientFactory: func() *MockStorageClient {
				return &MockStorageClient{
					FailCreateContainer: false,
				}
			},
			errMatcher: BeNil(),
			resMatcher: BeTrue(),
		},
		{
			name: "ShouldReturnErrorWhenCreationFails",
			storageClientFactory: func() *MockStorageClient {
				return &MockStorageClient{
					FailCreateContainer: true,
				}
			},
			errMatcher: Not(BeNil()),
			resMatcher: BeFalse(),
		},
	}

	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			client := c.storageClientFactory()
			res, err := client.CreateContainer("fakeContainerName", nil)
			g := NewGomegaWithT(t)
			g.Expect(err).To(c.errMatcher)
			g.Expect(res).To(c.resMatcher)
		})
	}
}

func TestAzureStorageClient_SaveBlockBlob(t *testing.T) {
	cases := []struct {
		name                 string
		storageClientFactory func() *MockStorageClient
		errMatcher           types.GomegaMatcher
	}{
		{
			name: "ShouldPassIfCreated",
			storageClientFactory: func() *MockStorageClient {
				return &MockStorageClient{
					FailSaveBlockBlob: false,
				}
			},
			errMatcher: BeNil(),
		},
		{
			name: "ShouldReturnErrorWhenCreationFails",
			storageClientFactory: func() *MockStorageClient {
				return &MockStorageClient{
					FailSaveBlockBlob: true,
				}
			},
			errMatcher: Not(BeNil()),
		},
	}

	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			client := c.storageClientFactory()
			err := client.SaveBlockBlob("fakeContainerName", "fakeBlobName", []byte("entity"), nil)
			g := NewGomegaWithT(t)
			g.Expect(err).To(c.errMatcher)
		})
	}
}
