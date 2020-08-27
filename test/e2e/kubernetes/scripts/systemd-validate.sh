#!/bin/bash

set -x
sudo kill $(pgrep dockerd)
sleep 10
if ! sudo systemctl is-active docker; then
  exit 1
fi
sudo kill $(pgrep kubelet)
sleep 10
if ! sudo systemctl is-active kubelet; then
  exit 1
fi
