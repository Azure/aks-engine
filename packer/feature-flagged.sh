#!/bin/bash

# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT license.

function installDockerEngine() {
    DOCKER_REPO="https://apt.dockerproject.org/repo"
    DOCKER_ENGINE_VERSION="1.13.*"
    dockerd --version
    if [ $? -eq 0 ]; then
        echo "dockerd is already installed, skipping download"
    else
        retrycmd_if_failure_no_stats 20 1 5 curl -fsSL https://aptdocker.azureedge.net/gpg > /tmp/aptdocker.gpg || exit $ERR_DOCKER_KEY_DOWNLOAD_TIMEOUT
        retrycmd_if_failure 10 5 10 apt-key add /tmp/aptdocker.gpg || exit $ERR_DOCKER_APT_KEY_TIMEOUT
        echo "deb ${DOCKER_REPO} ubuntu-xenial main" | sudo tee /etc/apt/sources.list.d/docker.list
        printf "Package: docker-engine\nPin: version ${DOCKER_ENGINE_VERSION}\nPin-Priority: 550\n" > /etc/apt/preferences.d/docker.pref
    fi
}