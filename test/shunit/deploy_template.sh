#!/usr/bin/env bash

source "${HOME}/test/common.sh"

function shunittest_deploy_template() {
  set -eux -o pipefail

  export OUTPUT="${HOME}/_output/${INSTANCE_NAME}"

  deploy_template
}
