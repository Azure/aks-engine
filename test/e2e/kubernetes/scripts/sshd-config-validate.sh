#!/bin/bash

set -x

UBUNTU_RELEASE=$(lsb_release -r -s)
CONFIGS=("ClientAliveInterval 120"
"ClientAliveCountMax 3"
"MACs hmac-sha2-512-etm@openssh.com,hmac-sha2-256-etm@openssh.com,umac-128-etm@openssh.com,hmac-sha2-512,hmac-sha2-256,umac-128@openssh.com"
"KexAlgorithms curve25519-sha256@libssh.org"
"Ciphers chacha20-poly1305@openssh.com,aes256-gcm@openssh.com,aes128-gcm@openssh.com,aes256-ctr,aes192-ctr,aes128-ctr"
"HostKey /etc/ssh/ssh_host_rsa_key"
"HostKey /etc/ssh/ssh_host_dsa_key"
"HostKey /etc/ssh/ssh_host_ecdsa_key"
"HostKey /etc/ssh/ssh_host_ed25519_key"
"SyslogFacility AUTH"
"LogLevel INFO"
"LoginGraceTime 60"
"PermitRootLogin no"
"PermitUserEnvironment no"
"StrictModes yes"
"PubkeyAuthentication yes"
"IgnoreRhosts yes"
"HostbasedAuthentication no"
"X11Forwarding no"
"MaxAuthTries 4"
"Banner /etc/issue.net"
"AcceptEnv LANG"
"AcceptEnv LC_*"
"Subsystem sftp /usr/lib/openssh/sftp-server"
"UsePAM yes"
"UseDNS no"
"GSSAPIAuthentication no")

for ((i = 0; i < ${#CONFIGS[@]}; i++))
do
    sshd -T | grep -i "${CONFIGS[$i]}" || exit 1
done

CONFIGS=("RSAAuthentication yes"
"UsePrivilegeSeparation yes"
"KeyRegenerationInterval 3600"
"ServerKeyBits 1024"
"RhostsRSAAuthentication no")
    
if [[ ${UBUNTU_RELEASE} == "16.04" ]]; then
    for ((i = 0; i < ${#CONFIGS[@]}; i++))
    do
        sshd -T | grep -i "${CONFIGS[$i]}" || exit 1
    done
fi