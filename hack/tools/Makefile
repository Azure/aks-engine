GO              ?= go
GOBIN			?= $(shell $(GO) env GOPATH)/bin

all: install

.PHONY: install
install: $(GOBIN)/go-bindata $(GOBIN)/gox $(GOBIN)/ginko $(GOBIN)/golangci-lint

$(GOBIN)/go-bindata:
	$(GO) get github.com/go-bindata/go-bindata/...@v3.1.2

$(GOBIN)/gox:
	$(GO) get github.com/mitchellh/gox

$(GOBIN)/ginko:
	$(GO) get github.com/onsi/ginkgo/ginkgo

$(GOBIN)/golangci-lint:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(GOPATH)/bin
