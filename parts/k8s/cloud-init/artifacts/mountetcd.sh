#!/bin/bash
# Mounting is done here instead of etcd because of bug https://bugs.launchpad.net/cloud-init/+bug/1692093
# Once the bug is fixed, replace the below with the cloud init changes replaced in https://github.com/Azure/aks-engine/pull/661.
set -x
DISK=/dev/sdc
PARTITION=${DISK}1
MOUNTPOINT=/var/lib/etcddisk
udevadm settle
mkdir -p $MOUNTPOINT
if mount | grep $MOUNTPOINT; then
  echo "disk is already mounted"
  umount /dev/sdc1
fi
if ! grep "/dev/sdc1" /etc/fstab; then
  echo "$PARTITION       $MOUNTPOINT       auto    defaults,nofail       0       2" >>/etc/fstab
fi
if ! ls $PARTITION; then
  /sbin/sgdisk --new 1 $DISK
  /sbin/mkfs.ext4 $PARTITION -L etcd_disk -F -E lazy_itable_init=1,lazy_journal_init=1
fi
mount $MOUNTPOINT
/bin/chown -R etcd:etcd /var/lib/etcddisk
#EOF
