GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_SHA    = $(shell git rev-parse --short HEAD)
GIT_TAG    = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null || echo "canary")
GIT_DIRTY  = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")

LDFLAGS += -X github.com/Azure/aks-engine/cmd.BuildSHA=${GIT_SHA}
LDFLAGS += -X github.com/Azure/aks-engine/cmd.GitTreeState=${GIT_DIRTY}
DOCKER_VERSION ?= git-${GIT_SHA}

ifneq ($(GIT_TAG),)
	LDFLAGS += -X github.com/Azure/aks-engine/cmd.BuildTag=${GIT_TAG}
endif

info:
	 @echo "Version:           ${VERSION}"
	 @echo "Git Tag:           ${GIT_TAG}"
	 @echo "Git Commit:        ${GIT_COMMIT}"
	 @echo "Git Tree State:    ${GIT_DIRTY}"
	 @echo "Docker Version:    ${DOCKER_VERSION}"
