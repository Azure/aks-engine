#!/bin/bash
# Add support for registering host on dns server. Must allow non-secure updates

set -e

DOMAINNAME=$1

echo $(date) " - Starting Script"

cat >/etc/network/if-up.d/register-dns <<EOFDHCP
#!/bin/sh

# only execute on the primary nic
if [ "\${IFACE}" = "eth0" ]
then
    ip=\$(ip address show eth0 | awk '/inet / {print \$2}' | cut -d/ -f1)
    host=\$(hostname -s)
    nsupdatecmds=/var/tmp/nsupdatecmds
    echo "update delete \${host}.${DOMAINNAME} a" > \$nsupdatecmds
    echo "update add \${host}.${DOMAINNAME} 3600 a \${ip}" >> \$nsupdatecmds
    echo "send" >> \$nsupdatecmds

    nsupdate \$nsupdatecmds
fi
EOFDHCP

chmod 755 /etc/network/if-up.d/register-dns

if ! grep -Fq "${DOMAINNAME}" /etc/dhcp/dhclient.conf; then
  echo $(date) " - Adding domain to dhclient.conf"

  echo "supersede domain-name \"${DOMAINNAME}\";" >>/etc/dhcp/dhclient.conf
  echo "prepend domain-search \"${DOMAINNAME}\";" >>/etc/dhcp/dhclient.conf
fi

# service networking restart
echo $(date) " - Restarting network"
sudo ifdown eth0 && sudo ifup eth0
