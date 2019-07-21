#!/bin/bash
set -x
source /opt/azure/containers/provision_source.sh

echo "  dns-search <searchDomainName>" | tee -a /etc/network/interfaces.d/50-cloud-init.cfg
systemctl_restart 20 5 10 networking
wait_for_apt_locks
retrycmd_if_failure 10 5 120 apt-get -y install realmd sssd sssd-tools samba-common samba samba-common python2.7 samba-libs packagekit
wait_for_apt_locks

function sourcekubeletscript() {
  source /opt/azure/containers/kubelet.sh
}

function updatevariables() {
  echo "<searchDomainRealmPassword>" | realm join -U <searchDomainRealmUser>@$(echo "<searchDomainName>" | tr /a-z/ /A-Z/) $(echo "<searchDomainName>" | tr /a-z/ /A-Z/)
}

function createrealmjoinscript() {
  grep 'realm join' /opt/azure/containers/setup-custom-search-domains.sh > /opt/azure/containers/realmjoin.sh
  chmod +x /opt/azure/containers/realmjoin.sh
  bash -x /opt/azure/containers/realmjoin.sh
}

sourcekubeletscript
createrealmjoinscript
