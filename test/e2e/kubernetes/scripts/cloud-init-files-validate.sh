#!/bin/bash
source /etc/profile.d/CIS.sh

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
for filepath in /etc/crontab /etc/cron.hourly /etc/cron.daily /etc/cron.weekly /etc/cron.monthly /etc/cron.d; do
  stat $filepath | grep 'Access: (0600' || exit 1
done
# validate grub configuration
sudo grep "^\s*linux" /boot/grub/grub.cfg | grep 'audit=1' || exit 1
# validate su configuration
sudo grep "auth required pam_wheel.so use_uid" /etc/pam.d/su
# validate password lockout config
sudo grep "auth required pam_tally2.so onerr=fail audit silent deny=5 unlock_time=900" /etc/pam.d/common-auth
# validate password change uniqueness config
sudo grep "pam_unix.so obscure use_authtok try_first_pass sha512 remember=5" /etc/pam.d/common-password
# validate umask configuration
touch test-umask
stat test-umask | grep 'Access: (0640' || exit 1
rm -f test-umask
