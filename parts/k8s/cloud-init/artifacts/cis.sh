#!/bin/bash

assignRootPW() {
  grep '^root:[!*]:' /etc/shadow
  if [ $? -eq '0' ] ; then
    SALT=`openssl rand -base64 5`
    SECRET=`openssl rand -base64 37`
    CMD="import crypt, getpass, pwd; print crypt.crypt('$SECRET', '\$6\$$SALT\$')"
    HASH=`python -c "$CMD"`

    echo 'root:'$HASH | /usr/sbin/chpasswd -e || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
  fi
}

assignFilePermissions() {
    FILES="
    auth.log
    alternatives.log
    cloud-init.log
    cloud-init-output.log
    daemon.log
    dpkg.log
    kern.log
    lastlog
    waagent.log
    syslog
    unattended-upgrades/unattended-upgrades.log
    azure-vnet-ipam.log
    azure-vnet-telemetry.log
    azure-cnimonitor.log
    azure-vnet.log
    kv-driver.log
    blobfuse-driver.log
    blobfuse-flexvol-installer.log
    "
    for FILE in ${FILES}; do
        FILEPATH="/var/log/${FILE}"
        DIR=$(dirname "${FILEPATH}")
        mkdir -p ${DIR} || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
        touch ${FILEPATH} || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
        chmod 640 ${FILEPATH} || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
    done
    find /var/log -type f -perm '/o+r' -exec chmod 'g-wx,o-rwx' {} \;
    chmod 600 /etc/passwd- || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
    chmod 600 /etc/shadow- || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
    chmod 600 /etc/group- || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
    chmod 644 /etc/sysctl.d/60-CIS.conf || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
    chmod 644 /etc/modprobe.d/CIS.conf || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
}

applyCIS() {
  assignRootPW
  assignFilePermissions
}