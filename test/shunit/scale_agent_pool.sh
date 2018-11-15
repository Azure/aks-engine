#!/usr/bin/env bash

# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT license.

source "${HOME}/test/common.sh"

function shunittest_scale_agent_pool {
  set -eux -o pipefail

  export OUTPUT="${HOME}/_output/${INSTANCE_NAME}"

  scale_agent_pool
}
