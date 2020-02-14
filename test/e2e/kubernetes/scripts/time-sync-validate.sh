#!/bin/bash
set -x
OS=$(sort -r /etc/*-release | gawk 'match($0, /^(ID_LIKE=(coreos)|ID=(.*))$/, a) { print toupper(a[2] a[3]); exit }')
UBUNTU_OS_NAME="UBUNTU"
RHEL_OS_NAME="RHEL"
COREOS_OS_NAME="COREOS"
DEBIAN_OS_NAME="DEBIAN"
if ! echo "${UBUNTU_OS_NAME} ${RHEL_OS_NAME} ${COREOS_OS_NAME} ${DEBIAN_OS_NAME}" | grep -q "${OS}"; then
    OS=$(sort -r /etc/*-release | gawk 'match($0, /^(ID_LIKE=(.*))$/, a) { print toupper(a[2] a[3]); exit }')
fi

if [[ $OS == $UBUNTU_OS_NAME ]]; then
  RELEASE=$(grep -oP 'DISTRIB_RELEASE=\K(\S+)' /etc/lsb-release)
  # verify that timesyncd configuration is healthy
  if [[ $RELEASE == "16.04" ]]; then
    sudo timedatectl status | grep 'Network time on: yes' || exit 1
    sudo timedatectl status | grep 'NTP synchronized: yes' || exit 1
  elif [[ $RELEASE == "18.04" ]]; then
    sudo timedatectl status | grep 'systemd-timesyncd.service active: yes' || exit 1
    sudo timedatectl status | grep 'System clock synchronized: yes' || exit 1
  fi
fi
if [[ $OS == $DEBIAN_OS_NAME ]]; then
  sudo timedatectl status | grep 'NTP service: active' || exit 1
  sudo timedatectl status | grep 'System clock synchronized: yes' || exit 1
fi

sudo timedatectl status | grep 'RTC in local TZ: no' || exit 1
sudo systemctl status systemd-timesyncd | grep 'Active: active' || exit 1
sudo systemctl status systemd-timesyncd | grep 'Status: "Synchronized to time server' || exit 1
