#!/usr/bin/env bash

# Scans master and agent nodes and applies missing labels.
# Kubernetes 1.16+ rejects the kubernetes.io/role and node-role.kubernetes.io
# labels in the "--node-labels" argument, but they need to be present for
# backward compatibility with Azure clusters.

set -euo pipefail

MASTER_LABELS="kubernetes.azure.com/role=master kubernetes.io/role=master node-role.kubernetes.io/master="
AGENT_LABELS="kubernetes.azure.com/role=agent kubernetes.io/role=agent node-role.kubernetes.io/agent="
MASTER_SELECTOR="kubernetes.azure.com/role!=agent,kubernetes.io/role!=agent"
AGENT_SELECTOR="kubernetes.azure.com/role!=master,kubernetes.io/role!=master"

# Find master nodes and label them
for node in $(kubectl get nodes -l $MASTER_SELECTOR -o name); do
  kubectl label --overwrite $node $MASTER_LABELS || echo "Error labeling master nodes"
done

# Find agent nodes and label them
for node in $(kubectl get nodes -l $AGENT_SELECTOR -o name); do
  kubectl label --overwrite $node $AGENT_LABELS || echo "Error labeling agent nodes"
done
