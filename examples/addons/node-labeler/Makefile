VERSION := $(shell git describe --tags --exact-match 2>/dev/null || echo latest)
REGISTRY ?= quay.io/
IMAGE_PREFIX ?= mboersma
IMAGE := ${REGISTRY}${IMAGE_PREFIX}/node-labeler:${VERSION}

all: info build push

build:
	GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags="-s -w" -o node-labeler *.go
	-upx node-labeler
	chmod +x node-labeler
	docker build -t ${IMAGE} .

push: build
	docker push ${IMAGE}

info:
	@echo "Version: ${VERSION}"
	@echo "Image:   ${IMAGE}"

.PHONY: all build info push
