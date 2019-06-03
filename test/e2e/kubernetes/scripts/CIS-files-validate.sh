#!/bin/bash
source /etc/profile.d/CIS.sh

CIS_FILES="
/etc/issue
/etc/issue.net
/etc/sysctl.d/60-CIS.conf
/etc/rsyslog.d/60-CIS.conf
/etc/security/pwquality.conf
/etc/pam.d/su
"
for CIS_FILE in ${CIS_FILES}; do
    ls -la $CIS_FILE || exit 2
    [ -s $CIS_FILE ] || exit 1
done

# specific file mode validations
for filepath in /etc/crontab /etc/cron.hourly /etc/cron.daily /etc/cron.weekly /etc/cron.monthly /etc/cron.d; do
  stat $filepath | grep 'Access: (0600' || exit 1
done
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
