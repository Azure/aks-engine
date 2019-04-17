#!/bin/bash

# verify that timesyncd configuration is healthy
sudo timedatectl status | grep 'Network time on: yes' || exit 1
sudo timedatectl status | grep 'NTP synchronized: yes' || exit 1
sudo timedatectl status | grep 'RTC in local TZ: no' || exit 1
sudo systemctl status systemd-timesyncd | grep 'Active: active' || exit 1
sudo systemctl status systemd-timesyncd | grep 'Status: "Synchronized to time server' || exit 1