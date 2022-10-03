GO              ?= go
LOCALBIN		= $(shell pwd)/bin

all: install

.PHONY: install
install: $(LOCALBIN)/go-bindata $(LOCALBIN)/gox $(LOCALBIN)/ginkgo $(LOCALBIN)/golangci-lint $(LOCALBIN)/pub $(LOCALBIN)/mockgen
	@echo > /dev/null

$(LOCALBIN)/go-bindata:
	GOBIN=$(LOCALBIN) $(GO) install github.com/go-bindata/go-bindata/go-bindata@v3.1.2

$(LOCALBIN)/gox:
	GOBIN=$(LOCALBIN) $(GO) install github.com/mitchellh/gox@v1.0.1

$(LOCALBIN)/ginkgo:
	GOBIN=$(LOCALBIN) $(GO) install github.com/onsi/ginkgo/ginkgo@v1.15.0

$(LOCALBIN)/golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(LOCALBIN) v1.41.1

$(LOCALBIN)/pub:
	GOBIN=$(LOCALBIN) $(GO) install github.com/devigned/pub@v0.3.3

$(LOCALBIN)/mockgen:
	GOBIN=$(LOCALBIN) $(GO) install github.com/golang/mock/mockgen@v1.2.0

.PHONY: reload
reload: clean install

.PHONY: clean
clean:
	rm -rf $(LOCALBIN)
