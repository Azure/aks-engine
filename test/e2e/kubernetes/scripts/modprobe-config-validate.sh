#!/bin/bash

set -x
modprobe -n -v tipc | egrep 'install /bin/true' || exit 1
[ -s `lsmod | grep tipc` ] || exit 1
modprobe -n -v dccp | egrep 'install /bin/true' || exit 1
[ -s `lsmod | grep dccp` ] || exit 1
modprobe -n -v sctp | egrep 'install /bin/true' || exit 1
[ -s `lsmod | grep sctp` ] || exit 1
modprobe -n -v rds | egrep 'install /bin/true' || exit 1
[ -s `lsmod | grep rds` ] || exit 1
modprobe -n -v cramfs | egrep 'install /bin/true' || exit 1
[ -s `lsmod | grep cramfs` ] || exit 1
modprobe -n -v freevxfs | egrep 'install /bin/true' || exit 1
[ -s `lsmod | grep freevxfs` ] || exit 1
modprobe -n -v jffs2 | egrep 'install /bin/true' || exit 1
[ -s `lsmod | grep jffs2` ] || exit 1
modprobe -n -v hfs | egrep 'install /bin/true' || exit 1
[ -s `lsmod | grep hfs` ] || exit 1
modprobe -n -v hfsplus | egrep 'install /bin/true' || exit 1
[ -s `lsmod | grep hfsplus` ] || exit 1
[ -s `lsmod | grep squashfs` ] || exit 1
[ -s `lsmod | grep vfat` ] || exit 1