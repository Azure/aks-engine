module github.com/Azure/aks-engine/test/e2e

go 1.16

require (
	github.com/Azure/aks-engine v0.43.0
	github.com/Azure/azure-sdk-for-go v43.0.0+incompatible
	github.com/Azure/go-autorest/autorest v0.11.18
	github.com/Azure/go-autorest/autorest/to v0.3.0
	github.com/influxdata/influxdb v1.7.9
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/onsi/ginkgo v1.15.0
	github.com/onsi/gomega v1.10.1
	github.com/pkg/errors v0.9.1
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292
)

replace github.com/Azure/aks-engine v0.43.0 => ../..
