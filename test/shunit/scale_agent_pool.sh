#!/usr/bin/env bash

source "${HOME}/test/common.sh"

function shunittest_scale_agent_pool() {
  set -eux -o pipefail

  export OUTPUT="${HOME}/_output/${INSTANCE_NAME}"

  scale_agent_pool
}
