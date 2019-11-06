GO              ?= go
BINDIR			:= bin
GINKGO			:= ginkgo

all: build

.PHONY: build
build: clean runner-build ginkgo-build

.PHONY: runner-build
runner-build:
	$(GO) build -tags=test -o ./$(BINDIR)/e2e-runner .

.PHONY: ginkgo-build
ginkgo-build:
	$(GINKGO) build -tags=test ./kubernetes

.PHONY: clean
clean:
	rm -f ./kubernetes/kubernetes.test
	rm -f ./$(BINDIR)/e2e-runner
