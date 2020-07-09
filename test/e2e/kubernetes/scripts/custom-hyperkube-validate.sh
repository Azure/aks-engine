#!/bin/bash

set -x

kubelet_version=$(docker run ${CUSTOM_HYPERKUBE_IMAGE} kubelet --version)
[[ $(kubelet --version) == "${kubelet_version}" ]] || exit 1

kubectl_version=$(docker run ${CUSTOM_HYPERKUBE_IMAGE} kubectl version --client)
[[ $(kubectl version --client) == "${kubectl_version}" ]] || exit 1
