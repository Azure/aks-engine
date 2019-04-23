#!/bin/bash

copyPackerFiles() {
  SYSCTL_CONFIG_SRC=/home/packer/sysctl-d-60-CIS.conf
  SYSCTL_CONFIG_DEST=/etc/sysctl.d/60-CIS.conf
  RSYSLOG_CONFIG_SRC=/home/packer/rsyslog-d-60-CIS.conf
  RSYSLOG_CONFIG_DEST=/etc/rsyslog.d/60-CIS.conf
  ETC_ISSUE_CONFIG_SRC=/home/packer/etc-issue
  ETC_ISSUE_CONFIG_DEST=/etc/issue
  ETC_ISSUE_NET_CONFIG_SRC=/home/packer/etc-issue.net
  ETC_ISSUE_NET_CONFIG_DEST=/etc/issue.net
  SSHD_CONFIG_SRC=/home/packer/sshd_config
  SSHD_CONFIG_DEST=/etc/ssh/sshd_config
  MODPROBE_CIS_SRC=/home/packer/modprobe-CIS.conf
  MODPROBE_CIS_DEST=/etc/modprobe.d/CIS.conf
  PWQUALITY_CONF_SRC=/home/packer/pwquality-CIS.conf
  PWQUALITY_CONF_DEST=/etc/security/pwquality.conf
  DEFAULT_GRUB_SRC=/home/packer/default-grub
  DEFAULT_GRUB_DEST=/etc/default/grub
  if [[ ${UBUNTU_RELEASE} == "16.04" ]]; then
    SSHD_CONFIG_SRC=/home/packer/sshd_config_1604
  fi
  DIR=$(dirname "$SYSCTL_CONFIG_DEST") && mkdir -p ${DIR} && cp $SYSCTL_CONFIG_SRC $SYSCTL_CONFIG_DEST || exit $ERR_CIS_COPY_FILE
  DIR=$(dirname "$RSYSLOG_CONFIG_DEST") && mkdir -p ${DIR} && cp $RSYSLOG_CONFIG_SRC $RSYSLOG_CONFIG_DEST || exit $ERR_CIS_COPY_FILE
  DIR=$(dirname "$ETC_ISSUE_CONFIG_DEST") && mkdir -p ${DIR} && cp $ETC_ISSUE_CONFIG_SRC $ETC_ISSUE_CONFIG_DEST || exit $ERR_CIS_COPY_FILE
  DIR=$(dirname "$ETC_ISSUE_NET_CONFIG_DEST") && mkdir -p ${DIR} && cp $ETC_ISSUE_NET_CONFIG_SRC $ETC_ISSUE_NET_CONFIG_DEST || exit $ERR_CIS_COPY_FILE
  DIR=$(dirname "$SSHD_CONFIG_DEST") && mkdir -p ${DIR} && cp $SSHD_CONFIG_SRC $SSHD_CONFIG_DEST || exit $ERR_CIS_COPY_FILE
  DIR=$(dirname "$MODPROBE_CIS_DEST") && mkdir -p ${DIR} && cp $MODPROBE_CIS_SRC $MODPROBE_CIS_DEST || exit $ERR_CIS_COPY_FILE
  DIR=$(dirname "$PWQUALITY_CONF_DEST") && mkdir -p ${DIR} && cp $PWQUALITY_CONF_SRC $PWQUALITY_CONF_DEST || exit $ERR_CIS_COPY_FILE
  DIR=$(dirname "$PWQUALITY_CONF_DEST") && mkdir -p ${DIR} && cp $DEFAULT_GRUB_SRC $DEFAULT_GRUB_DEST || exit $ERR_CIS_COPY_FILE
}

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
    unattended-upgrades/unattended-upgrades-dpkg.log
    azure-vnet-ipam.log
    azure-vnet-telemetry.log
    azure-cnimonitor.log
    azure-vnet.log
    kv-driver.log
    blobfuse-driver.log
    blobfuse-flexvol-installer.log
    landscape/sysinfo.log
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
    chmod 644 /etc/rsyslog.d/60-CIS.conf || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
    chmod 644 /etc/modprobe.d/CIS.conf || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
    chmod 644 /etc/ssh/sshd_config || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
    chmod 600 /etc/security/pwquality.conf || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
    chmod 400 /boot/grub/grub.cfg || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
    chmod 644 /etc/default/grub || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
}

applyCIS() {
  assignRootPW
  assignFilePermissions
}