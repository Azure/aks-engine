#!/bin/bash

assignRootPW() {
  if grep '^root:[!*]:' /etc/shadow; then
    SALT=$(openssl rand -base64 5)
    SECRET=$(openssl rand -base64 37)
    CMD="import crypt, getpass, pwd; print crypt.crypt('$SECRET', '\$6\$$SALT\$')"
    HASH=$(python -c "$CMD")

    echo 'root:'$HASH | /usr/sbin/chpasswd -e || exit 112
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
    mkdir -p ${DIR} || exit 112
    touch ${FILEPATH} || exit 112
    chmod 640 ${FILEPATH} || exit 112
  done
  find /var/log -type f -perm '/o+r' -exec chmod 'g-wx,o-rwx' {} \;
  chmod 600 /etc/passwd- || exit 112
  chmod 600 /etc/shadow- || exit 112
  chmod 600 /etc/group- || exit 112
  chmod 644 /etc/default/grub || exit 112
  for filepath in /etc/crontab /etc/cron.hourly /etc/cron.daily /etc/cron.weekly /etc/cron.monthly /etc/cron.d; do
    chmod 0600 $filepath || exit 112
  done
}

setPWExpiration() {
  sed -i "s|PASS_MAX_DAYS||g" /etc/login.defs || exit 115
  grep 'PASS_MAX_DAYS' /etc/login.defs && exit 115
  sed -i "s|PASS_MIN_DAYS||g" /etc/login.defs || exit 115
  grep 'PASS_MIN_DAYS' /etc/login.defs && exit 115
  sed -i "s|INACTIVE=||g" /etc/default/useradd || exit 115
  grep 'INACTIVE=' /etc/default/useradd && exit 115
  echo 'PASS_MAX_DAYS 90' >>/etc/login.defs || exit 115
  grep 'PASS_MAX_DAYS 90' /etc/login.defs || exit 115
  echo 'PASS_MIN_DAYS 7' >>/etc/login.defs || exit 115
  grep 'PASS_MIN_DAYS 7' /etc/login.defs || exit 115
  echo 'INACTIVE=30' >>/etc/default/useradd || exit 115
  grep 'INACTIVE=30' /etc/default/useradd || exit 115
}

applyCIS() {
  setPWExpiration
  assignRootPW
  assignFilePermissions
}

applyCIS

#EOF
