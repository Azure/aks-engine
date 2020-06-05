#!/bin/bash

set -x
[ $(findmnt /mnt -no FSTYPE) == "ext4" ] || exit 1
if [[ $MASTER_NODE == "true" ]]; then
  [ $(findmnt /var/lib/etcddisk -no FSTYPE) == "ext4" ] || exit 1
fi
