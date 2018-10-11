package armhelpers

import (
	"context"
	"fmt"
	"testing"

	"github.com/Azure/go-autorest/autorest"

	. "github.com/Azure/aks-engine/pkg/test"
	. "github.com/onsi/gomega"

	"github.com/Azure/go-autorest/autorest/azure"
	. "github.com/onsi/ginkgo"
)

func TestAzureClient(t *testing.T) {
	RunSpecsWithReporters(t, "AzureClient Tests", "Server Suite")
}

var _ = Describe("AzureClient Aux token tests", func() {
	It("Should set Aux token", func() {
		env, err := azure.EnvironmentFromName("AZUREPUBLICCLOUD")
		Expect(err).To(BeNil())

		token := "eyJ0eXAiOiJKV1QiL"
		azureClient, _ := NewAzureClientWithClientSecretExternalTenant(env, "subID", "d1a3-4ea4", "clientID", "secret")
		Expect(err).To(BeNil())
		azureClient.AddAuxiliaryTokens([]string{token})
		request, err := azureClient.deploymentsClient.GetPreparer(context.Background(), "testRG", "testDeployment")
		Expect(err).To(BeNil())
		Expect(request).To(Not(BeNil()))
		request, err = autorest.Prepare(request, azureClient.deploymentsClient.WithInspection())
		Expect(err).To(BeNil())
		Expect(request.Header.Get("x-ms-authorization-auxiliary")).To(Equal(fmt.Sprintf("Bearer %s", token)))
	})
})
