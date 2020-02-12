#!/bin/bash
set -x
source {{GetCSEHelpersScriptFilepath}}

echo "  dns-search {{GetSearchDomainName}}" | tee -a /etc/network/interfaces.d/50-cloud-init.cfg
systemctl_restart 20 5 10 networking
wait_for_apt_locks
retrycmd_if_failure 10 5 120 apt-get -y install realmd sssd sssd-tools samba-common samba samba-common python2.7 samba-libs packagekit
wait_for_apt_locks
echo "{{GetSearchDomainRealmPassword}}" | realm join -U {{GetSearchDomainRealmUser}}@$(echo "{{GetSearchDomainName}}" | tr /a-z/ /A-Z/) $(echo "{{GetSearchDomainName}}" | tr /a-z/ /A-Z/)
