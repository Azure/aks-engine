#!/usr/bin/env bash

# Applies missing master and agent labels to Kubernetes nodes.
#
# Kubelet 1.16+ rejects the `kubernetes.io/role` and `node-role.kubernetes.io`
# labels in its `--node-labels` argument, but they need to be present for
# backward compatibility.

set -euo pipefail

MASTER_SELECTOR="kubernetes.azure.com/role!=agent,kubernetes.io/role!=agent"
MASTER_LABELS="kubernetes.azure.com/role=master kubernetes.io/role=master node-role.kubernetes.io/master="
AGENT_SELECTOR="kubernetes.azure.com/role!=master,kubernetes.io/role!=master"
AGENT_LABELS="kubernetes.azure.com/role=agent kubernetes.io/role=agent node-role.kubernetes.io/agent="

kubectl label nodes --overwrite -l $MASTER_SELECTOR $MASTER_LABELS
kubectl label nodes --overwrite -l $AGENT_SELECTOR $AGENT_LABELS
#EOF
