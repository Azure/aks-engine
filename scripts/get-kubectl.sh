#!/usr/bin/env bash

# IMPORTANT: This script is meant for use only within CI environments as it assumes the linux amd64 architecture.

echo "==> Downloading kubectl version ${KUBECTL_VERSION} <=="

curl -L https://storage.googleapis.com/kubernetes-release/release/v${KUBECTL_VERSION}/bin/linux/amd64/kubectl -o /usr/local/bin/kubectl
chmod +x /usr/local/bin/kubectl