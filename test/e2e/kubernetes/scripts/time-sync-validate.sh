#!/bin/bash
set -x
OS=$(sort -r /etc/*-release | gawk 'match($0, /^(ID=(.*))$/, a) { print toupper(a[2] a[3]); exit }')
UBUNTU_OS_NAME="UBUNTU"
FLATCAR_OS_NAME="FLATCAR"
DEBIAN_OS_NAME="DEBIAN"
if ! echo "${UBUNTU_OS_NAME} ${FLATCAR_OS_NAME} ${DEBIAN_OS_NAME}" | grep -q "${OS}"; then
    OS=$(sort -r /etc/*-release | gawk 'match($0, /^(ID_LIKE=(.*))$/, a) { print toupper(a[2] a[3]); exit }')
fi

if [[ $OS == $UBUNTU_OS_NAME ]]; then
  RELEASE=$(grep -oP 'DISTRIB_RELEASE=\K(\S+)' /etc/lsb-release)
  # verify that timesyncd configuration is healthy
  if [[ $RELEASE == "16.04" ]]; then
    sudo timedatectl status | grep 'Network time on: yes' || exit 1
    sudo timedatectl status | grep 'NTP synchronized: yes' || exit 1
  elif [[ $RELEASE == "20.04" || $RELEASE == "18.04" ]]; then
    if apt list --installed | grep 'chrony'; then
      sudo chronyc sources | grep '#* PHC' || exit 1 # Make sure chrony is running and synced with host-based PTP source clock ('#' means local clock and '*' means synced)
    else
      sudo ntpstat | grep 'synchronised to NTP server' || exit 1
    fi
    sudo timedatectl status | grep 'System clock synchronized: yes' || exit 1
  fi
fi
if [[ $OS == $DEBIAN_OS_NAME ]]; then
  sudo timedatectl status | grep 'NTP service: active' || exit 1
  sudo timedatectl status | grep 'System clock synchronized: yes' || exit 1
fi

sudo timedatectl status | grep 'RTC in local TZ: no' || exit 1
if ! { [ $OS = $UBUNTU_OS_NAME ]; }; then
  sudo systemctl status systemd-timesyncd | grep 'Active: active' || exit 1
  sudo systemctl status systemd-timesyncd | grep 'Status: "Synchronized to time server' || exit 1
fi
