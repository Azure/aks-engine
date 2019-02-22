#!/bin/bash

####################################################
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
  DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
####################################################

# exit on errors
set -e
# exit on unbound variables
set -u
# verbose logging
set -x

source "$DIR/../utils.sh"
source "$DIR/k8s-utils.sh"

ENV_FILE="${CLUSTER_DEFINITION}.env"
if [ -e "${ENV_FILE}" ]; then
  source "${ENV_FILE}"
fi

EXPECTED_NODE_COUNT="${EXPECTED_NODE_COUNT:-4}"
EXPECTED_LINUX_AGENTS="${EXPECTED_LINUX_AGENTS:-3}"
EXPECTED_WINDOWS_AGENTS="${EXPECTED_WINDOWS_AGENTS:-0}"
EXPECTED_DNS="${EXPECTED_DNS:-2}"
EXPECTED_DASHBOARD="${EXPECTED_DASHBOARD:-1}"
EXPECTED_RESCHEDULER="${EXPECTED_RESCHEDULER:-0}"
EXPECTED_ORCHESTRATOR_VERSION="${EXPECTED_ORCHESTRATOR_VERSION:-}"

KUBE_PROXY_COUNT=$((EXPECTED_NODE_COUNT-$EXPECTED_WINDOWS_AGENTS))

# set TEST_ACR to "y" for ACR testing
TEST_ACR="${TEST_ACR:-n}"

namespace="namespace-${RANDOM}"
log "Running test in namespace: ${namespace}"
trap teardown EXIT

function teardown {
  k get all --all-namespaces || echo "teardown error"
  k get nodes || echo "teardown error"
  k get namespaces || echo "teardown error"
  k delete namespaces ${namespace} || echo "teardown error"
}

# TODO: cleanup the loops more
# TODO: the wc|awk business can just be kubectl with an output format and wc -l

###### Deploy ACR
if [[ "${TEST_ACR}" == "y" ]]; then
	ACR_NAME="${INSTANCE_NAME//[-._]/}1"
	ACR_REGISTRY="${ACR_NAME}-microsoft.azurecr.io" # fix this for non-ms tenant users
	if ! az acr show --resource-group "${RESOURCE_GROUP}" --name "${ACR_NAME}" ; then
		az acr create --location "${LOCATION}" --resource-group "${RESOURCE_GROUP}" --name "${ACR_NAME}" &
	fi
fi

###### Check node count
function check_node_count() {
  log "Checking node count"
  count=120
  while (( $count > 0 )); do
    log "  ... counting down $count"
    node_count=$(k get nodes --no-headers | grep -v NotReady | grep Ready | wc | awk '{print $1}')
    if (( ${node_count} == ${EXPECTED_NODE_COUNT} )); then break; fi
    sleep 5; count=$((count-1))
  done
  if (( $node_count != ${EXPECTED_NODE_COUNT} )); then
    log "K8S: gave up waiting for apiserver / node counts"; exit 1
  fi
}

check_node_count

###### Validate Kubernetes version
kubernetes_version=$(k version --short)
DASHBOARD_PORT=80
if [[ ${kubernetes_version} == *"Server Version: v1.9."* ]]; then
  DASHBOARD_PORT=443
fi
log "Setting dashboard port to ${DASHBOARD_PORT}"

log "Checking Kubernetes version. Expected: ${EXPECTED_ORCHESTRATOR_VERSION}"
if [ ! -z "${EXPECTED_ORCHESTRATOR_VERSION}" ]; then
  if [[ ${kubernetes_version} != *"Server Version: v${EXPECTED_ORCHESTRATOR_VERSION}"* ]]; then
    log "K8S: unexpected kubernetes version:\n${kubernetes_version}"; exit 1
  fi
fi

###### Wait for no more container creating
log "Checking containers being created"
count=60
while (( $count > 0 )); do
  log "  ... counting down $count"
  creating_count=$(k get nodes --no-headers | grep 'CreatingContainer' | wc | awk '{print $1}')
  if (( ${creating_count} == 0 )); then break; fi
  sleep 5; count=$((count-1))
done
if (( ${creating_count} != 0 )); then
  log "K8S: gave up waiting for containers"; exit 1
fi

###### Check existence and status of essential pods

# we test other essential pods (kube-dns, kube-proxy) separately
pods="heapster kube-addon-manager kube-apiserver kube-controller-manager kube-scheduler tiller"
if (( ${EXPECTED_RESCHEDULER} != 0 )); then
    pods="$pods rescheduler"
fi
log "Checking $pods"

count=60
while (( $count > 0 )); do
  for pod in $pods; do
    running=$(k get pods --all-namespaces | grep $pod | grep Running | wc -l)
    if (( $running > 0 )); then
      log "... $pod is Running"
      pods=$(echo $pods | sed -e "s/ *$pod */ /")
    fi
  done
  if [ -z "$(echo $pods | tr -d '[:space:]')" ]; then
    break
  fi
  sleep 5; count=$((count-1))
done

if [ ! -z "$(echo $pods | tr -d '[:space:]')" ]; then
  log "K8S: gave up waiting for running pods [$pods]"; exit 1
fi

###### Check for Kube-DNS
log "Checking Kube-DNS"
count=60
while (( $count > 0 )); do
  log "  ... counting down $count"
  running=$(k get pods --namespace=kube-system | grep kube-dns | grep Running | wc | awk '{print $1}')
  if (( ${running} == ${EXPECTED_DNS} )); then break; fi
  sleep 5; count=$((count-1))
done
if (( ${running} != ${EXPECTED_DNS} )); then
  log "K8S: gave up waiting for kube-dns"; exit 1
fi

###### Check for Kube-Dashboard
if (( ${EXPECTED_DASHBOARD} != 0 )); then
  log "Checking Kube-Dashboard"
  count=60
  while (( $count > 0 )); do
    log "  ... counting down $count"
    running=$(k get pods --namespace=kube-system | grep kubernetes-dashboard | grep Running | wc | awk '{print $1}')
    if (( ${running} == ${EXPECTED_DASHBOARD} )); then break; fi
    sleep 5; count=$((count-1))
  done
  if (( ${running} != ${EXPECTED_DASHBOARD} )); then
    log "K8S: gave up waiting for kubernetes-dashboard"; exit 1
  fi
else
  log "Expecting no dashboard"
fi

###### Check for Kube-Proxys
log "Checking Kube-Proxys"
count=60
while (( $count > 0 )); do
  log "  ... counting down $count"
  running=$(k get pods --namespace=kube-system | grep kube-proxy | grep Running | wc | awk '{print $1}')
  if (( ${running} == ${KUBE_PROXY_COUNT} )); then break; fi
  sleep 5; count=$((count-1))
done
if (( ${running} != ${KUBE_PROXY_COUNT} )); then
  log "K8S: gave up waiting for kube-proxy"; exit 1
fi

if ! [ $EXPECTED_WINDOWS_AGENTS -gt 0 ] ; then
  if (( ${EXPECTED_DASHBOARD} != 0 )); then
    # get master public hostname
    master=$(k config view | grep server | cut -f 3- -d "/" | tr -d " ")
    # get dashboard port
    port=$(k get svc --namespace=kube-system | grep dashboard | awk '{print $4}' | sed -n 's/^'${DASHBOARD_PORT}':\(.*\)\/TCP$/\1/p')
    # get internal IPs of the nodes
    ips=$(k get nodes --all-namespaces -o yaml | grep -B 1 InternalIP | grep address | awk '{print $3}')

    for ip in $ips; do
      log "Probing IP address ${ip}"
      count=60
      success="n"
      while (( $count > 0 )); do
        log "  ... counting down $count"
        ret=$(ssh -i "${OUTPUT}/id_rsa" -o ConnectTimeout=30 -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null "azureuser@${master}" "curl --max-time 60 http://${ip}:${port}" || echo "curl_error")
        if [[ ! $ret =~ .*curl_error.* ]]; then
          success="y"
          break
        fi
        if (( $count < 2 )); then
            log $ret
        fi
        sleep 5; count=$((count-1))
      done
      if [[ "${success}" == "n" ]]; then
        log "K8S: gave up verifying proxy"; exit 1
      fi
    done
  fi
fi

if [ $EXPECTED_LINUX_AGENTS -gt 0 ] ; then
  test_linux_deployment
fi

if [ $EXPECTED_WINDOWS_AGENTS -gt 0 ] ; then
  test_windows_deployment
fi

check_node_count
