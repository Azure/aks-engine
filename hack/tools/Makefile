GO              ?= go
LOCALBIN		= $(shell pwd)/bin

all: install

.PHONY: install
install: $(LOCALBIN)/go-bindata $(LOCALBIN)/gox $(LOCALBIN)/ginkgo $(LOCALBIN)/golangci-lint $(LOCALBIN)/pub $(LOCALBIN)/mockgen
	@echo > /dev/null

$(LOCALBIN)/go-bindata:
	GOBIN=$(LOCALBIN) $(GO) get github.com/go-bindata/go-bindata/...@v3.1.2

$(LOCALBIN)/gox:
	GOBIN=$(LOCALBIN) $(GO) get github.com/mitchellh/gox/...@v1.0.1

$(LOCALBIN)/ginkgo:
	GOBIN=$(LOCALBIN) $(GO) get github.com/onsi/ginkgo/ginkgo/...@v1.15.0

$(LOCALBIN)/golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(LOCALBIN) v1.39.0

$(LOCALBIN)/pub:
	GOBIN=$(LOCALBIN) $(GO) get github.com/devigned/pub/...@v0.2.6

$(LOCALBIN)/mockgen:
	GOBIN=$(LOCALBIN) $(GO) get github.com/golang/mock/mockgen@v1.2.0

.PHONY: reload
reload: clean install

.PHONY: clean
clean:
	rm -rf $(LOCALBIN)

.PHONY: tidy
tidy:
	$(GO) mod tidy
