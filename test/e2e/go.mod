module github.com/Azure/aks-engine/test/e2e

go 1.14

require (
	github.com/Azure/aks-engine v0.43.0
	github.com/Azure/azure-sdk-for-go v41.0.0+incompatible
	github.com/Azure/go-autorest/autorest/to v0.3.0
	github.com/influxdata/influxdb v1.7.9
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/onsi/ginkgo v1.10.3
	github.com/onsi/gomega v1.4.3
	github.com/pkg/errors v0.8.1
	golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413
)

replace github.com/Azure/aks-engine v0.43.0 => ../..
