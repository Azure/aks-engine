#!/bin/bash
IPV4_SEND_REDIRECTS_VALUE=0
IPV4_ACCEPT_SOURCE_ROUTE_VALUE=0
IPV4_ACCEPT_REDIRECTS_VALUE=0
IPV4_SECURE_REDIRECTS_VALUE=0
IPV4_LOG_MARTIANS_VALUE=1
IPV4_TCP_RETRIES2_VALUE=8
IPV6_ACCEPT_RA_VALUE=0
IPV6_ACCEPT_REDIRECTS_VALUE=0
KERNEL_PANIC_VALUE=10
KERNEL_PANIC_ON_OOPS_VALUE=1
VM_OVERCOMMIT_MEMORY_VALUE=1
INOTIFY_MAX_USER_WATCHES=1048576

set -x
grep $IPV4_SEND_REDIRECTS_VALUE /proc/sys/net/ipv4/conf/all/send_redirects || exit 1
grep $IPV4_SEND_REDIRECTS_VALUE /proc/sys/net/ipv4/conf/default/send_redirects || exit 1
grep $IPV4_ACCEPT_SOURCE_ROUTE_VALUE /proc/sys/net/ipv4/conf/all/accept_source_route || exit 1
grep $IPV4_ACCEPT_SOURCE_ROUTE_VALUE /proc/sys/net/ipv4/conf/default/accept_source_route || exit 1
grep $IPV4_ACCEPT_REDIRECTS_VALUE /proc/sys/net/ipv4/conf/all/accept_redirects || exit 1
grep $IPV4_ACCEPT_REDIRECTS_VALUE /proc/sys/net/ipv4/conf/default/accept_redirects || exit 1
grep $IPV4_SECURE_REDIRECTS_VALUE /proc/sys/net/ipv4/conf/all/secure_redirects || exit 1
grep $IPV4_SECURE_REDIRECTS_VALUE /proc/sys/net/ipv4/conf/default/secure_redirects || exit 1
grep $IPV4_LOG_MARTIANS_VALUE /proc/sys/net/ipv4/conf/all/log_martians || exit 1
grep $IPV4_LOG_MARTIANS_VALUE /proc/sys/net/ipv4/conf/default/log_martians || exit 1
grep $IPV6_ACCEPT_RA_VALUE /proc/sys/net/ipv6/conf/all/accept_ra || exit 1
grep $IPV6_ACCEPT_RA_VALUE /proc/sys/net/ipv6/conf/default/accept_ra || exit 1
grep $IPV6_ACCEPT_REDIRECTS_VALUE /proc/sys/net/ipv6/conf/all/accept_redirects || exit 1
grep $IPV6_ACCEPT_REDIRECTS_VALUE /proc/sys/net/ipv6/conf/default/accept_redirects || exit 1

# validate net config workaround from kubelet.service
grep $IPV4_TCP_RETRIES2_VALUE /proc/sys/net/ipv4/tcp_retries2 || exit 1
grep $KERNEL_PANIC_VALUE /proc/sys/kernel/panic || exit 1
grep $KERNEL_PANIC_ON_OOPS_VALUE /proc/sys/kernel/panic_on_oops || exit 1
grep $VM_OVERCOMMIT_MEMORY_VALUE /proc/sys/vm/overcommit_memory || exit 1

# TODO (@junaid-ali) Re-enable this test: validate inotify max_user_watches
#grep $INOTIFY_MAX_USER_WATCHES /proc/sys/fs/inotify/max_user_watches || exit 1
