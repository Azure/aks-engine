#!/bin/bash

set -x
timedatectl status | egrep 'NTP synchronized: yes' || exit 1
systemctl status chrony || exit 3
[ -s $(chronyc sources -v | grep 'Number of sources =' | awk '0+$6 <3 {print}') ] || exit 1