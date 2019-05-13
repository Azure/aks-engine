#!/bin/bash

retrycmd_if_failure() {
    retries=$1; wait_sleep=$2; timeout=$3; shift && shift && shift
    for i in $(seq 1 $retries); do
        timeout $timeout ${@}
        [ $? -eq 0  ] && break || \
        if [ $i -eq $retries ]; then
            echo Executed \"$@\" $i times;
            return 1
        else
            sleep $wait_sleep
        fi
    done
    echo Executed \"$@\" $i times;
}

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
retrycmd_if_failure 60 10 5 "STATUS=$(systemctl status systemd-timesyncd) && echo $STATUS | grep 'Active: active'" || exit 1
retrycmd_if_failure 120 10 5 "STATUS=$(systemctl status systemd-timesyncd) && echo $STATUS | grep 'Status: \"Synchronized to time server'\"" || exit 1

