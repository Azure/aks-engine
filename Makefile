TARGETS           = darwin/amd64 linux/amd64 windows/amd64
DIST_DIRS         = find * -type d -exec

.NOTPARALLEL:

.PHONY: bootstrap build test test_fmt validate-copyright-headers fmt lint ci

ifdef DEBUG
GOFLAGS   := -gcflags="-N -l" -mod=vendor
else
GOFLAGS   := -mod=vendor
endif

# go option
GO              ?= go
TAGS            :=
LDFLAGS         :=
BINDIR          := $(CURDIR)/bin
PROJECT         := aks-engine
VERSION         ?= $(shell git rev-parse HEAD)
VERSION_SHORT   ?= $(shell git rev-parse --short HEAD)
GITTAG          := $(shell git describe --exact-match --tags $(shell git log -n1 --pretty='%h') 2> /dev/null)
GOBIN           ?= $(shell $(GO) env GOPATH)/bin
TOOLSBIN        := $(CURDIR)/hack/tools/bin
AIKey           ?= c92d8284-b550-4b06-b7ba-e80fd7178faa
ifeq ($(GITTAG),)
GITTAG := $(VERSION_SHORT)
endif

DEV_ENV_IMAGE := quay.io/deis/go-dev:v1.27.0
DEV_ENV_WORK_DIR := /aks-engine
DEV_ENV_OPTS := --rm -v $(GOPATH)/pkg/mod:/go/pkg/mod -v $(CURDIR):$(DEV_ENV_WORK_DIR) -w $(DEV_ENV_WORK_DIR) $(DEV_ENV_VARS)
DEV_ENV_CMD := docker run $(DEV_ENV_OPTS) $(DEV_ENV_IMAGE)
DEV_ENV_CMD_IT := docker run -it $(DEV_ENV_OPTS) $(DEV_ENV_IMAGE)
DEV_CMD_RUN := docker run $(DEV_ENV_OPTS)
ifdef DEBUG
LDFLAGS := -X main.version=$(VERSION)
else
LDFLAGS := -s -X main.version=$(VERSION) -X github.com/Azure/$(PROJECT)/pkg/telemetry.AKSEngineAppInsightsKey=$(AIKey)
endif
BINARY_DEST_DIR ?= bin

ifeq ($(OS),Windows_NT)
	EXTENSION = .exe
	SHELL     = cmd.exe
	CHECK     = where.exe
else
	EXTENSION =
	SHELL     = bash
	CHECK     = which
endif

# Active module mode, as we use go modules to manage dependencies
export GO111MODULE=on

# Add the tools bin to the front of the path
export PATH := $(TOOLSBIN):$(PATH)

all: build

.PHONY: dev
dev:
	$(DEV_ENV_CMD_IT) bash

.PHONY: validate-dependencies
validate-dependencies: bootstrap
	@./scripts/validate-dependencies.sh

.PHONY: validate-copyright-headers
validate-copyright-headers:
	@./scripts/validate-copyright-header.sh

.PHONY: validate-go
validate-go:
	@./scripts/validate-go.sh

.PHONY: validate-shell
validate-shell:
	@./scripts/validate-shell.sh

.PHONY: generate
generate: bootstrap
	@echo "$$(go-bindata --version)"
	go generate $(GOFLAGS) -v ./... > /dev/null 2>&1

.PHONY: generate-azure-constants
generate-azure-constants:
	aks-engine get-locations -o code --client-id=$(AZURE_CLIENT_ID) --client-secret=$(AZURE_CLIENT_SECRET) --subscription-id=$(AZURE_SUBSCRIPTION_ID) \
	  > pkg/helpers/azure_locations.go
	aks-engine get-skus -o code --client-id=$(AZURE_CLIENT_ID) --client-secret=$(AZURE_CLIENT_SECRET) --subscription-id=$(AZURE_SUBSCRIPTION_ID) \
	  > pkg/helpers/azure_skus_const.go

.PHONY: build
build: generate go-build

.PHONY: go-build
go-build:
	$(GO) build $(GOFLAGS) -ldflags '$(LDFLAGS)' -o $(BINDIR)/$(PROJECT)$(EXTENSION) $(REPO_PATH)

.PHONY: tidy
tidy:
	$(GO) mod tidy

.PHONY: vendor
vendor: tidy
	$(GO) mod vendor

build-binary: generate
	go build $(GOFLAGS) -v -ldflags "$(LDFLAGS)" -o $(BINARY_DEST_DIR)/aks-engine .

# usage: make clean build-cross dist VERSION=v0.4.0
.PHONY: build-cross
build-cross: build
build-cross: LDFLAGS += -extldflags "-static"
build-cross:
	CGO_ENABLED=0 gox -output="_dist/aks-engine-$(GITTAG)-{{.OS}}-{{.Arch}}/{{.Dir}}" -osarch='$(TARGETS)' $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)'

.PHONY: build-azs-windows-k8s
build-azs-windows-k8s:
	./scripts/build-windows-k8s.sh -v $(K8S_VERSION) -p $(PATCH_VERSION) -a $(BUILD_AZURE_STACK)

.PHONY: dist
dist: build-cross compress-binaries
	( \
		cd _dist && \
		$(DIST_DIRS) cp ../LICENSE {} \; && \
		$(DIST_DIRS) cp ../README.md {} \; && \
		$(DIST_DIRS) tar -zcf {}.tar.gz {} \; && \
		$(DIST_DIRS) zip -r {}.zip {} \; \
	)

.PHONY: compress-binaries
compress-binaries:
	@which upx || (echo "Please install the upx executable packer tool. See https://upx.github.io/" && exit 1)
	find _dist -type f \( -name "aks-engine" -o -name "aks-engine.exe" \) -exec upx -9 {} +

.PHONY: checksum
checksum:
	for f in _dist/*.{gz,zip} ; do \
		shasum -a 256 "$${f}"  | awk '{print $$1}' > "$${f}.sha256" ; \
	done

.PHONY: build-container
build-container:
	docker build --no-cache --build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		--build-arg AKSENGINE_VERSION="$(VERSION)" -t microsoft/aks-engine:$(VERSION) \
		--file ./releases/Dockerfile.linux ./releases || \
	echo 'This target works only for published releases. For example, "VERSION=0.32.0 make build-container".'

.PHONY: clean
clean: tools-clean
	@rm -rf $(BINDIR) ./_dist ./pkg/helpers/unit_tests

GIT_BASEDIR    = $(shell git rev-parse --show-toplevel 2>/dev/null)
ifneq ($(GIT_BASEDIR),)
	LDFLAGS += -X github.com/Azure/aks-engine/pkg/test.JUnitOutDir=$(GIT_BASEDIR)/test/junit
endif

ginkgoBuild: generate
	make -C ./test/e2e ginkgo-build

test: generate ginkgoBuild
	ginkgo -mod=vendor -skipPackage test/e2e -failFast -r -v -tags=fast -ldflags '$(LDFLAGS)' .

.PHONY: test-style
test-style: validate-go validate-shell validate-copyright-headers

.PHONY: ensure-generated
ensure-generated:
	@echo "==> Checking generated files <=="
	@scripts/ensure-generated.sh

.PHONY: test-e2e
test-e2e:
	@test/e2e.sh

HAS_GIT := $(shell $(CHECK) git)

.PHONY: bootstrap
bootstrap: tools-install
ifndef HAS_GIT
	$(error You must install Git)
endif

.PHONY: tools-reload
tools-reload:
	make -C hack/tools reload

.PHONY: tools-install
tools-install:
	make -C hack/tools/

.PHONY: tools-clean
tools-clean:
	make -C hack/tools/ clean

ci: bootstrap test-style build test lint
	./scripts/coverage.sh --coveralls

.PHONY: coverage
coverage:
	LDFLAGS="$(LDFLAGS)" ./scripts/ginkgo.coverage.sh --codecov

include versioning.mk
include test.mk
include packer.mk
