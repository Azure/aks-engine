# syntax=docker/dockerfile:1.2

ARG GO_VERSION=1.19
ARG GO_IMAGE=golang:${GO_VERSION}


FROM --platform=$BUILDPLATFORM ${GO_IMAGE} AS build
COPY . /go/src/github.com/azure/aks-engine
WORKDIR /go/src/github.com/azure/aks-engine
ARG TARGETARCH
ARG TARGETOS
ARG TARGETVARIANT
SHELL ["/bin/bash", "-exc"]
ENV GOCACHE=/go/.buildcache
# We need to first generate everything w/o setting GOOS/GOARCH/GOARM
# Otherwise it will try to build supporting binaries (for the build itself) for the $TARGETPLATFORM, which we don't want.
# Those helper bins need to be for the $BUILDPLATFORM.
RUN \
	--mount=type=cache,target=/go/pkg/mod \
	--mount=type=cache,target=/go/.buildcache \
	make clean generate; \
	export GOOS="$TARGETOS"; \
	export GOARCH="$TARGETARCH"; \
	if [ -n TARGETVARIANT ] && [ "$TARGETARCH" = "arm" ]; then \
		export GOARM="${TARGETVARIANT//v}"; \
	fi; \
	make build


# Alaways last so a simple `docker build` with no `--target` will produce the binary
FROM scratch AS binary
COPY --from=build /go/src/github.com/azure/aks-engine/bin/* /
