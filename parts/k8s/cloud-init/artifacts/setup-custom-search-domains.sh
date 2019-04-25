#!/bin/bash
set -x
source /opt/azure/containers/provision_source.sh

sudo echo "  dns-search <searchDomainName>" >> /etc/network/interfaces.d/50-cloud-init.cfg
systemctl_restart 20 5 10 restart networking
wait_for_apt_locks
retrycmd_if_failure 10 5 120 apt-get -y install realmd sssd sssd-tools samba-common samba samba-common python2.7 samba-libs packagekit
wait_for_apt_locks
echo "<searchDomainRealmPassword>" | realm join -U <searchDomainRealmUser>@`echo "<searchDomainName>" | tr /a-z/ /A-Z/` `echo "<searchDomainName>" | tr /a-z/ /A-Z/`