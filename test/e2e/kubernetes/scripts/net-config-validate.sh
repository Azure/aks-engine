#!/bin/bash
IPV4_SEND_REDIRECTS_VALUE=0
IPV4_ACCEPT_SOURCE_ROUTE_VALUE=0
IPV4_ACCEPT_REDIRECTS_VALUE=0
IPV4_SECURE_REDIRECTS_VALUE=0
IPV4_LOG_MARTIANS_VALUE=1
IPV4_TCP_RETRIES2_VALUE=8
IPV6_ACCEPT_RA_VALUE=0
IPV6_ACCEPT_REDIRECTS_VALUE=0

set -x
cat /proc/sys/net/ipv4/conf/all/send_redirects | grep $IPV4_SEND_REDIRECTS_VALUE || exit 1
cat /proc/sys/net/ipv4/conf/default/send_redirects | grep $IPV4_SEND_REDIRECTS_VALUE || exit 1
cat /proc/sys/net/ipv4/conf/all/accept_source_route | grep $IPV4_ACCEPT_SOURCE_ROUTE_VALUE || exit 1
cat /proc/sys/net/ipv4/conf/default/accept_source_route | grep $IPV4_ACCEPT_SOURCE_ROUTE_VALUE || exit 1
cat /proc/sys/net/ipv4/conf/all/accept_redirects | grep $IPV4_ACCEPT_REDIRECTS_VALUE || exit 1
cat /proc/sys/net/ipv4/conf/default/accept_redirects | grep $IPV4_ACCEPT_REDIRECTS_VALUE || exit 1
cat /proc/sys/net/ipv4/conf/all/secure_redirects | grep $IPV4_SECURE_REDIRECTS_VALUE || exit 1
cat /proc/sys/net/ipv4/conf/default/secure_redirects | grep $IPV4_SECURE_REDIRECTS_VALUE || exit 1
cat /proc/sys/net/ipv4/conf/all/log_martians | grep $IPV4_LOG_MARTIANS_VALUE || exit 1
cat /proc/sys/net/ipv4/conf/default/log_martians | grep $IPV4_LOG_MARTIANS_VALUE || exit 1
cat /proc/sys/net/ipv6/conf/all/accept_ra | grep $IPV6_ACCEPT_RA_VALUE || exit 1
cat /proc/sys/net/ipv6/conf/default/accept_ra | grep $IPV6_ACCEPT_RA_VALUE || exit 1
cat /proc/sys/net/ipv6/conf/all/accept_redirects | grep $IPV6_ACCEPT_REDIRECTS_VALUE || exit 1
cat /proc/sys/net/ipv6/conf/default/accept_redirects | grep $IPV6_ACCEPT_REDIRECTS_VALUE || exit 1

# validate net config workaround from kubelet.service
cat /proc/sys/net/ipv4/tcp_retries2 | grep $IPV4_TCP_RETRIES2_VALUE || exit 1
