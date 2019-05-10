#!/usr/bin/env bash

function shunittest_validate_deployment {
  set -eux -o pipefail

  export OUTPUT="${HOME}/_output/${INSTANCE_NAME}"
  export SSH_KEY="${OUTPUT}/id_rsa"
  if [[ "${ORCHESTRATOR}" == "kubernetes" ]]; then
    export KUBECONFIG="${OUTPUT}/kubeconfig/kubeconfig.${LOCATION}.json"
    nodes=$(${HOME}/test/step.sh get_node_count)
    IFS=':' read -a narr <<< "${nodes}"
    export EXPECTED_NODE_COUNT=${narr[0]}
    export EXPECTED_LINUX_AGENTS=${narr[1]}
    export EXPECTED_WINDOWS_AGENTS=${narr[1]}
    EXPECTED_ORCHESTRATOR_VERSION=$(${HOME}/test/step.sh get_orchestrator_release)
    export EXPECTED_ORCHESTRATOR_VERSION
  fi

  script="${HOME}/test/cluster-tests/${ORCHESTRATOR}/test.sh"

  if [ -x "${script}" ]; then
    "${script}"
  else
    echo "${script}: not an executable or no such file"
    exit 1
  fi
}
