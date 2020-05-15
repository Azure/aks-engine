#!/bin/bash
# Mounting is done here instead of etcd because of bug https://bugs.launchpad.net/cloud-init/+bug/1692093
# Once the bug is fixed, replace the below with the cloud init changes replaced in https://github.com/Azure/aks-engine/pull/661.
set -x
MOUNTPOINT=/var/lib/etcddisk
LABEL=etcd_disk
ETCDDISK=""
PARTITION=""
ETCDDISK_PERSIST=/opt/azure/containers/etcd_disk_initial
udevadm settle
if [ ! -f $ETCDDISK_PERSIST ]; then
  for DISK in $(grep -o -G "sd[a-z]" /proc/partitions | uniq); do
    if ! grep "$DISK"1 /proc/partitions; then
      if [[ -n $ETCDDISK ]]; then
        exit 1
      fi
      ETCDDISK=/dev/${DISK}
      echo "${ETCDDISK}" > /opt/azure/containers/etcd_disk_initial
    fi;
  done
  if [[ -z $ETCDDISK ]]; then
    exit 1
  fi
else
  ETCDDISK=$(cat ${ETCDDISK_PERSIST})
fi
PARTITION=${ETCDDISK}1
if ! ls $PARTITION; then
  /sbin/sgdisk --new 1 $ETCDDISK
fi
if ! blkid -L $LABEL; then
  /sbin/mkfs.ext4 $PARTITION -L $LABEL -F -E lazy_itable_init=1,lazy_journal_init=1
fi
mkdir -p $MOUNTPOINT
if ! grep "$MOUNTPOINT" /etc/fstab; then
  echo "LABEL=${LABEL}       $MOUNTPOINT       auto    defaults,nofail       0       2" >>/etc/fstab
fi
umount $MOUNTPOINT
mount $MOUNTPOINT
/bin/chown -R etcd:etcd /var/lib/etcddisk
#EOF
