#!/bin/bash

OS=$(cat /etc/*-release | grep ^ID= | tr -d 'ID="' | awk '{print toupper($0)}')
UBUNTU_OS_NAME="UBUNTU"

if [[ $OS == $UBUNTU_OS_NAME ]]; then
    exit_code=0

    ENSURE_NOT_INSTALLED="
    postfix
    "

    for PACKAGE in ${ENSURE_NOT_INSTALLED}; do
        if apt list --installed | grep -E "^${PACKAGE}" > /dev/null; then
            echo >&2 "$PACKAGE is installed but shouldn't be"
            exit_code=1
        fi
    done

    ENSURE_INSTALLED="
    apt-transport-https
    blobfuse
    ca-certificates
    ceph-common
    cgroup-lite
    cifs-utils
    conntrack
    ebtables
    ethtool
    ipvsadm
    fuse
    git
    glusterfs-client
    init-system-helpers
    iproute2
    ipset
    iptables
    jq
    mount
    nfs-common
    pigz
    socat
    util-linux
    xz-utils
    zip
    htop
    iotop
    iftop
    sysstat
    "

    for PACKAGE in ${ENSURE_INSTALLED}; do
        if ! apt list --installed | grep -E "^${PACKAGE}" > /dev/null; then
            echo >&2 "$PACKAGE is not installed but should be"
            exit_code=1
        fi
    done

    exit $exit_code
else
    exit 1
fi
