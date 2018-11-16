#!/bin/bash -eux

# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT license.

## Cleanup packer SSH key and machine ID generated for this boot
rm -f /root/.ssh/authorized_keys
rm -f /home/packer/.ssh/authorized_keys
rm -f /etc/machine-id
touch /etc/machine-id