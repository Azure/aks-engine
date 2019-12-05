package armhelpers

import (
	"context"
	"testing"
)

func TestVMImageFetcherInterface(t *testing.T) {

	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}

	mc.RegisterLogin()
	mc.RegisterVMImageFetcherInterface()

	err = mc.Activate()
	if err != nil {
		t.Fatalf("failed to activate HttpMockClient - %s", err)
	}
	defer mc.DeactivateAndReset()

	env := mc.GetEnvironment()
	azureClient, err := NewAzureClientWithClientSecret(env, subscriptionID, "clientID", "secret")
	if err != nil {
		t.Fatalf("can not get client %s", err)
	}

	_, err = azureClient.GetVirtualMachineImage(context.Background(), location, publisher, offer, sku, version)
	if err != nil {
		t.Error(err)
	}

	_, err = azureClient.ListVirtualMachineImages(context.Background(), location, publisher, offer, sku, "", 50, "")
	if err != nil {
		t.Error(err)
	}

}
