#!/bin/bash

CLOUD_INIT_FILES="
/etc/issue
/etc/issue.net
/etc/sysctl.d/60-CIS.conf
/etc/rsyslog.d/60-CIS.conf
/etc/security/pwquality.conf
/etc/default/grub
/etc/pam.d/su
"
for CLOUD_INIT_FILE in ${CLOUD_INIT_FILES}; do
    ls -la $CLOUD_INIT_FILE || exit 2
    [ -s $CLOUD_INIT_FILE ] || exit 1
done

# verify that no files under /var/log have read access to everyone
sudo find /var/log -type f -perm '/o+r' | (! grep ^) || exit 1
# specific file mode validations
stat /boot/grub/grub.cfg | grep 'Access: (0400' || exit 1
# validate grub configuration
sudo grep "^\s*linux" /boot/grub/grub.cfg | grep 'audit=1' || exit 1
# validate su configuration
sudo grep "auth required pam_wheel.so use_uid" /etc/pam.d/su