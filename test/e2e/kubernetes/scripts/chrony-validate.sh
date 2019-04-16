#!/bin/bash

set -x
timedatectl status | egrep 'NTP synchronized: yes' || exit 1
systemctl status chrony || exit 3
chronyc sources -v | egrep 'Number of sources = 1' || exit 1