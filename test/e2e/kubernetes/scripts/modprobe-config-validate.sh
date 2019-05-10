#!/bin/bash

set -x
modprobe -n -v tipc | grep -E 'install /bin/true' || exit 1
[ -s $(lsmod | grep tipc) ] || exit 1
modprobe -n -v dccp | grep -E 'install /bin/true' || exit 1
[ -s $(lsmod | grep dccp) ] || exit 1
modprobe -n -v sctp | grep -E 'install /bin/true' || exit 1
[ -s $(lsmod | grep sctp) ] || exit 1
modprobe -n -v rds | grep -E 'install /bin/true' || exit 1
[ -s $(lsmod | grep rds) ] || exit 1
modprobe -n -v cramfs | grep -E 'install /bin/true' || exit 1
[ -s $(lsmod | grep cramfs) ] || exit 1
modprobe -n -v freevxfs | grep -E 'install /bin/true' || exit 1
[ -s $(lsmod | grep freevxfs) ] || exit 1
modprobe -n -v jffs2 | grep -E 'install /bin/true' || exit 1
[ -s $(lsmod | grep jffs2) ] || exit 1
modprobe -n -v hfs | grep -E 'install /bin/true' || exit 1
[ -s $(lsmod | grep hfs) ] || exit 1
modprobe -n -v hfsplus | grep -E 'install /bin/true' || exit 1
[ -s $(lsmod | grep hfsplus) ] || exit 1
[ -s $(lsmod | grep squashfs) ] || exit 1
[ -s $(lsmod | grep vfat) ] || exit 1
