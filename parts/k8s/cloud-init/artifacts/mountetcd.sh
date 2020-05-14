#!/bin/bash
# Mounting is done here instead of etcd because of bug https://bugs.launchpad.net/cloud-init/+bug/1692093
# Once the bug is fixed, replace the below with the cloud init changes replaced in https://github.com/Azure/aks-engine/pull/661.
set -x
ETCDDISK=""
PARTITION=""
udevadm settle
for DISK in $(grep -o -G "sd[a-z]" /proc/partitions | uniq); do
  if ! grep "$DISK"1 /proc/partitions; then
    if [[ -n $PARTITION ]]; then
      exit 1
    fi
    ETCDDISK=/dev/${DISK}
    PARTITION=${ETCDDISK}1
  fi;
done
if [[ -n $ETCDDISK ]]; then
  exit 1
fi
MOUNTPOINT=/var/lib/etcddisk
mkdir -p $MOUNTPOINT
umount $MOUNTPOINT
if ! grep "$MOUNTPOINT" /etc/fstab; then
  echo "LABEL=etcd_disk       $MOUNTPOINT       auto    defaults,nofail       0       2" >>/etc/fstab
fi
if ! ls $PARTITION; then
  /sbin/sgdisk --new 1 $ETCDDISK
  /sbin/mkfs.ext4 $PARTITION -L etcd_disk -F -E lazy_itable_init=1,lazy_journal_init=1
fi
mount $MOUNTPOINT
/bin/chown -R etcd:etcd /var/lib/etcddisk
#EOF
