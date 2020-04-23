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
SOMAXCONN=16384
TCP_MAX_SYN_BACKLOG=16384
MESSAGE_COST=40
MESSAGE_BURST=80
IPV4_NEIGH_GC_THRESH1=4096
IPV4_NEIGH_GC_THRESH2=8192
IPV4_NEIGH_GC_THRESH3=16384

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

# validate kubelet.service configs
grep $IPV4_TCP_RETRIES2_VALUE /proc/sys/net/ipv4/tcp_retries2 || exit 1
grep $KERNEL_PANIC_VALUE /proc/sys/kernel/panic || exit 1
grep $KERNEL_PANIC_ON_OOPS_VALUE /proc/sys/kernel/panic_on_oops || exit 1
grep $VM_OVERCOMMIT_MEMORY_VALUE /proc/sys/vm/overcommit_memory || exit 1
grep $INOTIFY_MAX_USER_WATCHES /proc/sys/fs/inotify/max_user_watches || exit 1
grep $SOMAXCONN /proc/sys/net/core/somaxconn || exit 1
grep $TCP_MAX_SYN_BACKLOG /proc/sys/net/ipv4/tcp_max_syn_backlog || exit 1
grep $MESSAGE_COST /proc/sys/net/core/message_cost || exit 1
grep $MESSAGE_BURST /proc/sys/net/core/message_burst || exit 1

grep $IPV4_NEIGH_GC_THRESH1 /proc/sys/net/ipv4/neigh/default/gc_thresh1 || exit 1
grep $IPV4_NEIGH_GC_THRESH2 /proc/sys/net/ipv4/neigh/default/gc_thresh2 || exit 1
grep $IPV4_NEIGH_GC_THRESH3 /proc/sys/net/ipv4/neigh/default/gc_thresh3 || exit 1
