package armhelpers

import (
	"testing"

	. "github.com/Azure/aks-engine/pkg/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAzureStorageClient(t *testing.T) {
	RunSpecsWithReporters(t, "AzureStorageClient", "Server Suite")
}

var _ = Describe("CreateContainer Test", func() {
	It("Should pass if container created", func() {
		client := MockStorageClient{
			FailCreateContainer: false,
		}
		created, err := client.CreateContainer("fakeContainerName", nil)
		Expect(err).To(BeNil())
		Expect(created).To(BeTrue())
	})

	It("Should return error when container creation failed", func() {
		client := MockStorageClient{
			FailCreateContainer: true,
		}
		created, err := client.CreateContainer("fakeContainerName", nil)
		Expect(err).NotTo(BeNil())
		Expect(created).To(BeFalse())
	})
})

var _ = Describe("SaveBlockBlob Test", func() {
	It("Should pass if container created", func() {
		client := MockStorageClient{
			FailSaveBlockBlob: false,
		}
		err := client.SaveBlockBlob("fakeContainerName", "fakeBlobName", []byte("entity"), nil)
		Expect(err).To(BeNil())
	})

	It("Should return error when container creation failed", func() {
		client := MockStorageClient{
			FailSaveBlockBlob: true,
		}
		err := client.SaveBlockBlob("fakeContainerName", "fakeBlobName", []byte("entity"), nil)
		Expect(err).NotTo(BeNil())
	})
})
