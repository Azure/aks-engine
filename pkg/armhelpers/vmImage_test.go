package armhelpers

import "testing"

func TestVMImageFetcherInterface(t *testing.T) {

	mc, err := NewHTTPMockClient()
	if err != nil {
		t.Fatalf("failed to create HttpMockClient - %s", err)
	}

	mc.RegisterLogin()

}
