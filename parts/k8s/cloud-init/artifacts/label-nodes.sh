#!/usr/bin/env bash

# This script looks for nodes that don't have the required kubernetes.io/role label,
# and then labels them appropriately
#
# Because nodes can't self-label their role in this way,
# we need control plane nodes to do this continually
# to ensure that new nodes get the role labels quickly

set -euo pipefail

KUBECONFIG="$(find /home/*/.kube/config)"
KUBECTL="kubectl --kubeconfig=${KUBECONFIG}"

MASTER_SELECTOR="kubernetes.azure.com/role!=agent,kubernetes.io/role!=agent"
MASTER_LABELS="kubernetes.azure.com/role=master kubernetes.io/role=master node-role.kubernetes.io/master="
AGENT_SELECTOR="kubernetes.azure.com/role!=master,kubernetes.io/role!=master"
AGENT_LABELS="kubernetes.azure.com/role=agent kubernetes.io/role=agent node-role.kubernetes.io/agent="

${KUBECTL} label nodes --overwrite -l $MASTER_SELECTOR $MASTER_LABELS
${KUBECTL} label nodes --overwrite -l $AGENT_SELECTOR $AGENT_LABELS
#EOF
