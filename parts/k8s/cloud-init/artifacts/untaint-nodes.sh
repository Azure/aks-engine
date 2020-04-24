#!/usr/bin/env bash

KUBECONFIG="$(find /home/*/.kube/config)"
KUBECTL="kubectl --kubeconfig=${KUBECONFIG}"
AAD_POD_IDENTITY_TAINT=node.kubernetes.io/aad-pod-identity-not-ready=true:NoSchedule

if ! ${KUBECTL} get daemonsets -n kube-system -o json | jq -e -r '.items[] | select(.metadata.name == "nmi")' > /dev/null; then
  for node in $(${KUBECTL} get nodes -o json | jq -e -r '.items[] | .metadata.name'); do
    ${KUBECTL} taint nodes $node $AAD_POD_IDENTITY_TAINT- 2>&1 | grep -v 'not found';
  done
  exit 0
fi
for pod in $(${KUBECTL} get pods -n kube-system -o json | jq -r '.items[] | select(.status.phase == "Running") | .metadata.name'); do
  if [[ "$pod" =~ ^nmi ]]; then
    ${KUBECTL} taint nodes $(${KUBECTL} get pod ${pod} -n kube-system -o json | jq -r '.spec.nodeName') $AAD_POD_IDENTITY_TAINT- 2>&1 | grep -v 'not found';
  fi;
done
#EOF
