#!/usr/bin/env bash

KUBECONFIG="$(find /home/*/.kube/config)"
KUBECTL="kubectl --kubeconfig=${KUBECONFIG}"

for pod in $(${KUBECTL} get pods -n kube-system -o json | jq -r '.items[] | select(.status.phase == "Running") | .metadata.name'); do
{{if IsAADPodIdentityAddonEnabled}}
  if [[ "$pod" =~ ^nmi ]]; then
    ${KUBECTL} taint nodes $(${KUBECTL} get pod ${pod} -n kube-system -o json | jq -r '.spec.nodeName') node.kubernetes.io/aad-pod-identity-not-ready=true:NoSchedule- 2>&1 | grep -v 'not found';
  fi;
{{end}}
done
#EOF
