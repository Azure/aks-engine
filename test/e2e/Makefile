GO              ?= go
BINDIR          := bin
GINKGO          := ginkgo
AIKey           := c92d8284-b550-4b06-b7ba-e80fd7178faa
LDFLAGS         := -X github.com/Azure/$(PROJECT)/pkg/telemetry.AKSEngineAppInsightsKey=$(AIKey)

all: build

.PHONY: build
build: clean runner-build ginkgo-build

.PHONY: runner-build
runner-build:
	$(GO) build -tags=test -o ./$(BINDIR)/e2e-runner -ldflags '$(LDFLAGS)' .

.PHONY: ginkgo-build
ginkgo-build:
	$(GINKGO) build -tags=test -ldflags '$(LDFLAGS)' ./kubernetes

.PHONY: clean
clean:
	rm -f ./kubernetes/kubernetes.test
	rm -f ./$(BINDIR)/e2e-runner