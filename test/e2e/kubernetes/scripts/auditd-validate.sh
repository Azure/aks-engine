#!/bin/bash

set -x
systemctl status auditd
RET=$?
if $ENABLED; then
  if [ $RET -eq 0 ]; then
    exit 0
  else
    exit 1
  fi
else
  if [ $RET -ne 0 ]; then
    exit 0
  else
    exit 1
  fi
fi
