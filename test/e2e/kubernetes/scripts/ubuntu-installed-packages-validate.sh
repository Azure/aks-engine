#!/bin/bash
set -x
OS=$(sort -r /etc/*-release | gawk 'match($0, /^(ID_LIKE=(coreos)|ID=(.*))$/, a) { print toupper(a[2] a[3]); exit }')
UBUNTU_OS_NAME="UBUNTU"
RHEL_OS_NAME="RHEL"
COREOS_OS_NAME="COREOS"
DEBIAN_OS_NAME="DEBIAN"
if ! echo "${UBUNTU_OS_NAME} ${RHEL_OS_NAME} ${COREOS_OS_NAME} ${DEBIAN_OS_NAME}" | grep -q "${OS}"; then
    OS=$(sort -r /etc/*-release | gawk 'match($0, /^(ID_LIKE=(.*))$/, a) { print toupper(a[2] a[3]); exit }')
fi

ENSURE_NOT_INSTALLED="
postfix
"
for PACKAGE in ${ENSURE_NOT_INSTALLED}; do
    apt list --installed | grep -E "^${PACKAGE}" && exit 1
done

if [[ $OS == $UBUNTU_OS_NAME ]]; then
    ENSURE_INSTALLED_UBUNTU="
ceph-common
cgroup-lite
glusterfs-client
"
    for PACKAGE in ${ENSURE_INSTALLED_UBUNTU}; do
        apt list --installed | grep -E "^${PACKAGE}" || exit 1
    done
fi

ENSURE_INSTALLED="
apt-transport-https
blobfuse
ca-certificates
cifs-utils
conntrack
dbus
ebtables
ethtool
fuse
git
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
    apt list --installed | grep -E "^${PACKAGE}" || exit 1
done
