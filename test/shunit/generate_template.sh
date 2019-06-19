#!/usr/bin/env bash

source "${HOME}/test/common.sh"

function shunittest_generate_template() {
  set -eux -o pipefail

  export OUTPUT="${HOME}/_output/${INSTANCE_NAME}"

  generate_template
}
