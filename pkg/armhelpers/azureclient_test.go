// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"
	"fmt"
	"testing"

	"github.com/Azure/go-autorest/autorest"

	. "github.com/onsi/gomega"

	"github.com/Azure/go-autorest/autorest/azure"
)

func TestAzureClientAuxToken(t *testing.T) {
	t.Parallel()

	env, err := azure.EnvironmentFromName("AZUREPUBLICCLOUD")
	g := NewGomegaWithT(t)
	g.Expect(err).To(BeNil())

	token := "eyJ0eXAiOiJKV1QiL"
	azureClient, err := NewAzureClientWithClientSecretExternalTenant(env, "subID", "d1a3-4ea4", "clientID", "secret")
	g.Expect(err).To(BeNil())
	azureClient.AddAuxiliaryTokens([]string{token})
	request, err := azureClient.deploymentsClient.GetPreparer(context.Background(), "testRG", "testDeployment")
	g.Expect(err).To(BeNil())
	g.Expect(request).To(Not(BeNil()))
	request, err = autorest.Prepare(request, azureClient.deploymentsClient.WithInspection())
	g.Expect(err).To(BeNil())
	g.Expect(request.Header.Get("x-ms-authorization-auxiliary")).To(Equal(fmt.Sprintf("Bearer %s", token)))
}
