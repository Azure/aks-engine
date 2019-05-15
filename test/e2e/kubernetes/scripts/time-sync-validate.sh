#!/bin/bash

RELEASE=$(grep ^DISTRIB_RELEASE= /etc/lsb-release | tr -d 'DISTRIB_RELEASE="' | awk '{print toupper($0)}')
# verify that timesyncd configuration is healthy
set -x
if [[ $RELEASE == "16.04" ]]; then
  sudo timedatectl status | grep 'Network time on: yes' || exit 1
  sudo timedatectl status | grep 'NTP synchronized: yes' || exit 1
elif [[ $RELEASE == "18.04" ]]; then
  sudo timedatectl status | grep 'systemd-timesyncd.service active: yes' || exit 1
  sudo timedatectl status | grep 'System clock synchronized: yes' || exit 1

fi
sudo timedatectl status | grep 'RTC in local TZ: no' || exit 1
sudo systemctl status systemd-timesyncd | grep 'Active: active' || exit 1
sudo systemctl status systemd-timesyncd | grep 'Status: "Synchronized to time server' || exit 1
