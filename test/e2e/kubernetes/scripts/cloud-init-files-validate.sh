#!/bin/bash

CLOUD_INIT_FILES="
/etc/issue
/etc/issue.net
"
for CLOUD_INIT_FILE in ${CLOUD_INIT_FILES}; do
    ls -la $CLOUD_INIT_FILE || exit 2
    [ -s $CLOUD_INIT_FILE ] || exit 1
done

# verify that no files under /var/log have read access to everyone
sudo find /var/log -type f -perm '/o+r' | (! grep ^) || exit 1