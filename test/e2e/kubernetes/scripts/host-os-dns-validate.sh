#!/bin/bash

set -x
sudo ifconfig -a -v || exit 1
cat /etc/resolv.conf || exit 1
set +x
# validate that all node vms can resolve DNS to all other node vms, including themselves
# also validate external DNS lookups
# configurable retries
success="no"
retries=${LOOKUP_RETRIES}
HOSTS="${NODE_HOSTNAMES} www.bing.com google.com"
for i in $(seq 1 $retries); do
  for host in $HOSTS; do
    set -x
    addrs=$(dig +short +search +answer ${host})
    if [ -z "$addrs" ]; then
      success="no"
      break
    fi
    set +x
    success="yes"
  done
  if [[ "${success}" == "yes" ]]; then
    break
  fi
  if [ $i -eq $retries ]; then
    exit 1
  else
    sleep 10
  fi
done

success="no"
HOSTS="www.bing.com google.com"
for i in $(seq 1 $retries); do
  for host in $HOSTS; do
    set -x
    addrs=$(dig +short +search +answer ${host} @8.8.8.8)
    if [ -z "$addrs" ]; then
      success="no"
      break
    fi
    set +x
    success="yes"
  done
  if [[ "${success}" == "yes" ]]; then
    break
  fi
  if [ $i -eq $retries ]; then
    exit 1
  else
    sleep 10
  fi
done
