#!/bin/bash

ENSURE_NOT_INSTALLED="
postfix
"
for PACKAGE in ${ENSURE_NOT_INSTALLED}; do
    apt list --installed | grep $PACKAGE && exit 1
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
    apt list --installed | grep $PACKAGE || exit 1
done