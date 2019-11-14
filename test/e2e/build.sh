#!/usr/bin/env bash

set -x

GOPATH="/go"
WORK_DIR="/aks-engine"

# Assumes we're running from the git root of aks-engine
docker run --rm \
-v $(pwd):${WORK_DIR} \
-w ${WORK_DIR} \
"${DEV_IMAGE}" make build-binary || exit 1
