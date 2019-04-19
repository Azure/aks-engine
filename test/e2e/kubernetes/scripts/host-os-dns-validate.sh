#!/bin/bash

set -x
ifconfig -a -v || exit 1
cat /etc/resolv.conf || exit 1
set +x
# validate that all node vms can resolve DNS to all other node vms, including themselves
# also validate external DNS lookups
# retry failures for up to 10 mins
success="no"
retries=1
HOSTS="${NODE_HOSTNAMES} www.bing.com google.com"
for i in $(seq 1 $retries); do
  for host in $HOSTS; do
    set -x
    [ -s $(dig +short +search +answer ${host}) ]
      if [ $? -eq 0  ]; then
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
    [ -s $(dig +short +search +answer ${host} @8.8.8.8) ]
      if [ $? -eq 0  ]; then
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