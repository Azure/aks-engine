GO              ?= go
GOBIN			?= $(shell $(GO) env GOPATH)/bin

all: install

.PHONY: install
install: $(GOBIN)/go-bindata $(GOBIN)/gox $(GOBIN)/ginkgo $(GOBIN)/golangci-lint

$(GOBIN)/go-bindata:
	$(GO) get github.com/go-bindata/go-bindata/...@v3.1.2

$(GOBIN)/gox:
	$(GO) get github.com/mitchellh/gox/...@v1.0.1

$(GOBIN)/ginkgo:
	$(GO) get github.com/onsi/ginkgo/ginkgo/...@v1.10.1

$(GOBIN)/golangci-lint:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(GOPATH)/bin v1.21.0


.PHONY: reload
reload: clean install

.PHONY: clean
clean:
	rm -f $(GOBIN)/go-bindata $(GOBIN)/gox $(GOBIN)/ginkgo $(GOBIN)/golangci-lint
