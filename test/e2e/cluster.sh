#!/bin/bash

set -x

TMP_DIR=$(mktemp -d "$(pwd)/XXXXXXXXXXXX")
TMP_BASENAME=$(basename ${TMP_DIR})
GOPATH="/go"
WORK_DIR="/aks-engine"
MASTER_VM_UPGRADE_SKU="${MASTER_VM_UPGRADE_SKU:-Standard_D4_v3}"
AZURE_ENV="${AZURE_ENV:-AzurePublicCloud}"
mkdir -p _output || exit 1

# Assumes we're running from the git root of aks-engine
if [ "${BUILD_AKS_ENGINE}" = "true" ]; then
  docker run --rm \
  -v $(pwd):${WORK_DIR} \
  -w ${WORK_DIR} \
  "${DEV_IMAGE}" make build-binary || exit 1
fi

cat > ${TMP_DIR}/apimodel-input.json <<END
${API_MODEL_INPUT}
END

# Jenkinsfile will yield a null $ADD_NODE_POOL_INPUT if not set in the test job config
if [ "$ADD_NODE_POOL_INPUT" == "null" ]; then
  ADD_NODE_POOL_INPUT=""
fi
if [ "$LB_TEST_TIMEOUT" == "" ]; then
  LB_TEST_TIMEOUT="${E2E_TEST_TIMEOUT}"
fi

if [ -n "$ADD_NODE_POOL_INPUT" ]; then
  cat > ${TMP_DIR}/addpool-input.json <<END
${ADD_NODE_POOL_INPUT}
END
fi

echo "Running E2E tests against a cluster built with the following API model:"
cat ${TMP_DIR}/apimodel-input.json

CLEANUP_AFTER_DEPLOYMENT=${CLEANUP_ON_EXIT}
if [ "${UPGRADE_CLUSTER}" = "true" ] || [ "${SCALE_CLUSTER}" = "true" ] || [ -n "$ADD_NODE_POOL_INPUT" ]; then
  CLEANUP_AFTER_DEPLOYMENT="false"
fi

if [ -n "${GINKGO_SKIP}" ]; then
  if [ -n "${GINKGO_SKIP_AFTER_SCALE_DOWN}" ]; then
    SKIP_AFTER_SCALE_DOWN="${GINKGO_SKIP}|${GINKGO_SKIP_AFTER_SCALE_DOWN}"
  else
    SKIP_AFTER_SCALE_DOWN="${GINKGO_SKIP}"
  fi
  if [ -n "${GINKGO_SKIP_AFTER_SCALE_UP}" ]; then
    SKIP_AFTER_SCALE_UP="${GINKGO_SKIP}|${GINKGO_SKIP_AFTER_SCALE_UP}"
  else
    SKIP_AFTER_SCALE_UP="${GINKGO_SKIP}"
  fi
  if [ "${SCALE_CLUSTER}" = "true" ]; then
    SKIP_AFTER_UPGRADE="${GINKGO_SKIP}|${SKIP_AFTER_SCALE_DOWN}"
  else
    SKIP_AFTER_UPGRADE="${GINKGO_SKIP}"
  fi
else
  SKIP_AFTER_SCALE_DOWN="${GINKGO_SKIP_AFTER_SCALE_DOWN}"
  SKIP_AFTER_SCALE_UP="${GINKGO_SKIP_AFTER_SCALE_UP}"
  if [ "${SCALE_CLUSTER}" = "true" ]; then
    SKIP_AFTER_UPGRADE="${SKIP_AFTER_SCALE_DOWN}"
  else
    SKIP_AFTER_UPGRADE=""
  fi
fi

docker run --rm \
-v $(pwd):${WORK_DIR} \
-v /etc/ssl/certs:/etc/ssl/certs \
-w ${WORK_DIR} \
-e CLUSTER_DEFINITION=${TMP_BASENAME}/apimodel-input.json \
-e CLIENT_ID="${AZURE_CLIENT_ID}" \
-e CLIENT_SECRET="${AZURE_CLIENT_SECRET}" \
-e CLIENT_OBJECTID="${CLIENT_OBJECTID}" \
-e TENANT_ID="${AZURE_TENANT_ID}" \
-e SUBSCRIPTION_ID="${AZURE_SUBSCRIPTION_ID}" \
-e INFRA_RESOURCE_GROUP="${INFRA_RESOURCE_GROUP}" \
-e ORCHESTRATOR=kubernetes \
-e ORCHESTRATOR_RELEASE="${ORCHESTRATOR_RELEASE}" \
-e CREATE_VNET="${CREATE_VNET}" \
-e TIMEOUT="${E2E_TEST_TIMEOUT}" \
-e LB_TIMEOUT="${LB_TEST_TIMEOUT}" \
-e KUBERNETES_IMAGE_BASE=$KUBERNETES_IMAGE_BASE \
-e KUBERNETES_IMAGE_BASE_TYPE=$KUBERNETES_IMAGE_BASE_TYPE \
-e CLEANUP_ON_EXIT=${CLEANUP_AFTER_DEPLOYMENT} \
-e SKIP_LOGS_COLLECTION="${SKIP_LOGS_COLLECTION}" \
-e REGIONS="${REGION_OPTIONS}" \
-e WINDOWS_NODE_IMAGE_GALLERY="${WINDOWS_NODE_IMAGE_GALLERY}" \
-e WINDOWS_NODE_IMAGE_NAME="${WINDOWS_NODE_IMAGE_NAME}" \
-e WINDOWS_NODE_IMAGE_RESOURCE_GROUP="${WINDOWS_NODE_IMAGE_RESOURCE_GROUP}" \
-e WINDOWS_NODE_IMAGE_SUBSCRIPTION_ID="${WINDOWS_NODE_IMAGE_SUBSCRIPTION_ID}" \
-e WINDOWS_NODE_IMAGE_VERSION="${WINDOWS_NODE_IMAGE_VERSION}" \
-e WINDOWS_NODE_VHD_URL="${WINDOWS_NODE_VHD_URL}" \
-e LINUX_NODE_IMAGE_GALLERY=$LINUX_NODE_IMAGE_GALLERY \
-e LINUX_NODE_IMAGE_NAME=$LINUX_NODE_IMAGE_NAME \
-e LINUX_NODE_IMAGE_RESOURCE_GROUP=$LINUX_NODE_IMAGE_RESOURCE_GROUP \
-e LINUX_NODE_IMAGE_SUBSCRIPTION_ID=$LINUX_NODE_IMAGE_SUBSCRIPTION_ID \
-e LINUX_NODE_IMAGE_VERSION=$LINUX_NODE_IMAGE_VERSION \
-e OS_DISK_SIZE_GB=$OS_DISK_SIZE_GB \
-e DISTRO=$DISTRO \
-e CONTAINER_RUNTIME=$CONTAINER_RUNTIME \
-e LOG_ANALYTICS_WORKSPACE_KEY="${LOG_ANALYTICS_WORKSPACE_KEY}" \
-e CUSTOM_HYPERKUBE_IMAGE="${CUSTOM_HYPERKUBE_IMAGE}" \
-e IS_JENKINS="${IS_JENKINS}" \
-e SKIP_TEST="${SKIP_TESTS}" \
-e GINKGO_FAIL_FAST=true \
-e GINKGO_FOCUS="${GINKGO_FOCUS}" \
-e GINKGO_SKIP="${GINKGO_SKIP}" \
-e API_PROFILE="${API_PROFILE}" \
-e CUSTOM_CLOUD_NAME="${ENVIRONMENT_NAME}" \
-e IDENTITY_SYSTEM="${IDENTITY_SYSTEM}" \
-e AUTHENTICATION_METHOD="${AUTHENTICATION_METHOD}" \
-e LOCATION="${LOCATION}" \
-e CUSTOM_CLOUD_CLIENT_ID="${CUSTOM_CLOUD_CLIENT_ID}" \
-e CUSTOM_CLOUD_SECRET="${CUSTOM_CLOUD_SECRET}" \
-e PORTAL_ENDPOINT="${PORTAL_ENDPOINT}" \
-e SERVICE_MANAGEMENT_ENDPOINT="${SERVICE_MANAGEMENT_ENDPOINT}" \
-e RESOURCE_MANAGER_ENDPOINT="${RESOURCE_MANAGER_ENDPOINT}" \
-e STORAGE_ENDPOINT_SUFFIX="${STORAGE_ENDPOINT_SUFFIX}" \
-e KEY_VAULT_DNS_SUFFIX="${KEY_VAULT_DNS_SUFFIX}" \
-e ACTIVE_DIRECTORY_ENDPOINT="${ACTIVE_DIRECTORY_ENDPOINT}" \
-e GALLERY_ENDPOINT="${GALLERY_ENDPOINT}" \
-e GRAPH_ENDPOINT="${GRAPH_ENDPOINT}" \
-e SERVICE_MANAGEMENT_VM_DNS_SUFFIX="${SERVICE_MANAGEMENT_VM_DNS_SUFFIX}" \
-e RESOURCE_MANAGER_VM_DNS_SUFFIX="${RESOURCE_MANAGER_VM_DNS_SUFFIX}" \
"${DEV_IMAGE}" make test-kubernetes || exit 1

if [ "${UPGRADE_CLUSTER}" = "true" ] || [ "${SCALE_CLUSTER}" = "true" ] || [ -n "$ADD_NODE_POOL_INPUT" ] || [ "${GET_CLUSTER_LOGS}" = "true" ]; then
  # shellcheck disable=SC2012
  RESOURCE_GROUP=$(ls -dt1 _output/* | head -n 1 | cut -d/ -f2)
  docker run --rm \
    -v $(pwd):${WORK_DIR} \
    -w ${WORK_DIR} \
    -e RESOURCE_GROUP=$RESOURCE_GROUP \
    ${DEV_IMAGE} \
    /bin/bash -c "chmod -R 777 _output/$RESOURCE_GROUP _output/$RESOURCE_GROUP/apimodel.json" || exit 1
  # shellcheck disable=SC2012
  REGION=$(ls -dt1 _output/* | head -n 1 | cut -d/ -f2 | cut -d- -f2)
  API_SERVER="$RESOURCE_GROUP.$REGION.cloudapp.azure.com"
  if [ "${AZURE_ENV}" = "AzureStackCloud" ]; then
    API_SERVER="$RESOURCE_GROUP.$REGION.$RESOURCE_MANAGER_VM_DNS_SUFFIX"
  fi
  
  if [ "${GET_CLUSTER_LOGS}" = "true" ]; then
      docker run --rm \
      -v $(pwd):${WORK_DIR} \
      -w ${WORK_DIR} \
      -e RESOURCE_GROUP=$RESOURCE_GROUP \
      -e REGION=$REGION \
      ${DEV_IMAGE} \
      ./bin/aks-engine get-logs \
      --api-model _output/$RESOURCE_GROUP/apimodel.json \
      --location $REGION \
      --ssh-host $API_SERVER \
      --linux-ssh-private-key _output/$RESOURCE_GROUP-ssh \
      --linux-script ./scripts/collect-logs.sh
      # TODO remove --linux-script once collect-logs.sh is part of the VHD
  fi

  if [ $(( RANDOM % 4 )) -eq 3 ]; then
    echo Removing bookkeeping tags from VMs in resource group $RESOURCE_GROUP ...
    az login --username ${AZURE_CLIENT_ID} --password ${AZURE_CLIENT_SECRET} --tenant ${AZURE_TENANT_ID} --service-principal > /dev/null
    for vm_type in vm vmss; do
      for vm in $(az $vm_type list -g $RESOURCE_GROUP --subscription ${AZURE_SUBSCRIPTION_ID} --query '[].name' -o table | tail -n +3); do
        az $vm_type update -n $vm -g $RESOURCE_GROUP --subscription ${AZURE_SUBSCRIPTION_ID} --set tags={} > /dev/null
      done
    done
  fi
  git reset --hard
  git remote rm $UPGRADE_FORK
  git remote add $UPGRADE_FORK https://github.com/$UPGRADE_FORK/aks-engine.git
  git fetch --prune $UPGRADE_FORK
  git branch -D $UPGRADE_FORK/$UPGRADE_BRANCH
  git checkout -b $UPGRADE_FORK/$UPGRADE_BRANCH --track $UPGRADE_FORK/$UPGRADE_BRANCH
  git pull
  git log -1
  docker run --rm \
    -v $(pwd):${WORK_DIR} \
    -w ${WORK_DIR} \
    "${DEV_IMAGE}" make build-binary > /dev/null 2>&1 || exit 1
else
  exit 0
fi

if [ -n "$ADD_NODE_POOL_INPUT" ]; then
  docker run --rm \
    -v $(pwd):${WORK_DIR} \
    -v /etc/ssl/certs:/etc/ssl/certs \
    -w ${WORK_DIR} \
    -e RESOURCE_GROUP=$RESOURCE_GROUP \
    -e REGION=$REGION \
    ${DEV_IMAGE} \
    ./bin/aks-engine addpool \
    --azure-env ${AZURE_ENV} \
    --subscription-id ${AZURE_SUBSCRIPTION_ID} \
    --api-model _output/$RESOURCE_GROUP/apimodel.json \
    --node-pool ${TMP_BASENAME}/addpool-input.json \
    --location $REGION \
    --resource-group $RESOURCE_GROUP \
    --auth-method client_secret \
    --client-id ${AZURE_CLIENT_ID} \
    --client-secret ${AZURE_CLIENT_SECRET} || exit 1

  CLEANUP_AFTER_ADD_NODE_POOL=${CLEANUP_ON_EXIT}
  if [ "${UPGRADE_CLUSTER}" = "true" ] || [ "${SCALE_CLUSTER}" = "true" ]; then
    CLEANUP_AFTER_ADD_NODE_POOL="false"
  fi

  docker run --rm \
    -v $(pwd):${WORK_DIR} \
    -v /etc/ssl/certs:/etc/ssl/certs \
    -w ${WORK_DIR} \
    -e CLIENT_ID=${AZURE_CLIENT_ID} \
    -e CLIENT_SECRET=${AZURE_CLIENT_SECRET} \
    -e CLIENT_OBJECTID=${CLIENT_OBJECTID} \
    -e TENANT_ID=${AZURE_TENANT_ID} \
    -e SUBSCRIPTION_ID=${AZURE_SUBSCRIPTION_ID} \
    -e INFRA_RESOURCE_GROUP="${INFRA_RESOURCE_GROUP}" \
    -e ORCHESTRATOR=kubernetes \
    -e NAME=$RESOURCE_GROUP \
    -e TIMEOUT=${E2E_TEST_TIMEOUT} \
    -e LB_TIMEOUT=${LB_TEST_TIMEOUT} \
    -e KUBERNETES_IMAGE_BASE=$KUBERNETES_IMAGE_BASE \
    -e KUBERNETES_IMAGE_BASE_TYPE=$KUBERNETES_IMAGE_BASE_TYPE \
    -e CLEANUP_ON_EXIT=${CLEANUP_AFTER_ADD_NODE_POOL} \
    -e REGIONS=$REGION \
    -e IS_JENKINS=${IS_JENKINS} \
    -e SKIP_LOGS_COLLECTION=true \
    -e GINKGO_FAIL_FAST=true \
    -e GINKGO_SKIP="${SKIP_AFTER_SCALE_DOWN}" \
    -e GINKGO_FOCUS="${GINKGO_FOCUS}" \
    -e SKIP_TEST=${SKIP_TESTS_AFTER_ADD_POOL} \
    -e ADD_NODE_POOL_INPUT=${ADD_NODE_POOL_INPUT} \
    -e API_PROFILE="${API_PROFILE}" \
    -e CUSTOM_CLOUD_NAME="${ENVIRONMENT_NAME}" \
    -e IDENTITY_SYSTEM="${IDENTITY_SYSTEM}" \
    -e AUTHENTICATION_METHOD="${AUTHENTICATION_METHOD}" \
    -e LOCATION="${LOCATION}" \
    -e CUSTOM_CLOUD_CLIENT_ID="${CUSTOM_CLOUD_CLIENT_ID}" \
    -e CUSTOM_CLOUD_SECRET="${CUSTOM_CLOUD_SECRET}" \
    -e PORTAL_ENDPOINT="${PORTAL_ENDPOINT}" \
    -e SERVICE_MANAGEMENT_ENDPOINT="${SERVICE_MANAGEMENT_ENDPOINT}" \
    -e RESOURCE_MANAGER_ENDPOINT="${RESOURCE_MANAGER_ENDPOINT}" \
    -e STORAGE_ENDPOINT_SUFFIX="${STORAGE_ENDPOINT_SUFFIX}" \
    -e KEY_VAULT_DNS_SUFFIX="${KEY_VAULT_DNS_SUFFIX}" \
    -e ACTIVE_DIRECTORY_ENDPOINT="${ACTIVE_DIRECTORY_ENDPOINT}" \
    -e GALLERY_ENDPOINT="${GALLERY_ENDPOINT}" \
    -e GRAPH_ENDPOINT="${GRAPH_ENDPOINT}" \
    -e SERVICE_MANAGEMENT_VM_DNS_SUFFIX="${SERVICE_MANAGEMENT_VM_DNS_SUFFIX}" \
    -e RESOURCE_MANAGER_VM_DNS_SUFFIX="${RESOURCE_MANAGER_VM_DNS_SUFFIX}" \
    ${DEV_IMAGE} make test-kubernetes || exit 1
fi

if [ "${SCALE_CLUSTER}" = "true" ]; then
  for nodepool in $(jq -r  '.properties.agentPoolProfiles[].name' < _output/$RESOURCE_GROUP/apimodel.json); do
    docker run --rm \
      -v $(pwd):${WORK_DIR} \
      -v /etc/ssl/certs:/etc/ssl/certs \
      -w ${WORK_DIR} \
      -e RESOURCE_GROUP=$RESOURCE_GROUP \
      -e REGION=$REGION \
      ${DEV_IMAGE} \
      ./bin/aks-engine scale \
      --azure-env ${AZURE_ENV} \
      --subscription-id ${AZURE_SUBSCRIPTION_ID} \
      --api-model _output/$RESOURCE_GROUP/apimodel.json \
      --location $REGION \
      --resource-group $RESOURCE_GROUP \
      --apiserver $API_SERVER \
      --node-pool $nodepool \
      --new-node-count 1 \
      --auth-method client_secret \
      --client-id ${AZURE_CLIENT_ID} \
      --client-secret ${AZURE_CLIENT_SECRET} || exit 1
  done

  docker run --rm \
    -v $(pwd):${WORK_DIR} \
    -v /etc/ssl/certs:/etc/ssl/certs \
    -w ${WORK_DIR} \
    -e CLIENT_ID=${AZURE_CLIENT_ID} \
    -e CLIENT_SECRET=${AZURE_CLIENT_SECRET} \
    -e CLIENT_OBJECTID=${CLIENT_OBJECTID} \
    -e TENANT_ID=${AZURE_TENANT_ID} \
    -e SUBSCRIPTION_ID=${AZURE_SUBSCRIPTION_ID} \
    -e INFRA_RESOURCE_GROUP="${INFRA_RESOURCE_GROUP}" \
    -e ORCHESTRATOR=kubernetes \
    -e NAME=$RESOURCE_GROUP \
    -e TIMEOUT=${E2E_TEST_TIMEOUT} \
    -e LB_TIMEOUT=${LB_TEST_TIMEOUT} \
    -e KUBERNETES_IMAGE_BASE=$KUBERNETES_IMAGE_BASE \
    -e KUBERNETES_IMAGE_BASE_TYPE=$KUBERNETES_IMAGE_BASE_TYPE \
    -e CLEANUP_ON_EXIT=false \
    -e REGIONS=$REGION \
    -e IS_JENKINS=${IS_JENKINS} \
    -e SKIP_LOGS_COLLECTION=true \
    -e GINKGO_FAIL_FAST=true \
    -e GINKGO_SKIP="${SKIP_AFTER_SCALE_DOWN}" \
    -e GINKGO_FOCUS="${GINKGO_FOCUS}" \
    -e SKIP_TEST=${SKIP_TESTS_AFTER_SCALE_DOWN} \
    -e ADD_NODE_POOL_INPUT=${ADD_NODE_POOL_INPUT} \
    -e API_PROFILE="${API_PROFILE}" \
    -e CUSTOM_CLOUD_NAME="${ENVIRONMENT_NAME}" \
    -e IDENTITY_SYSTEM="${IDENTITY_SYSTEM}" \
    -e AUTHENTICATION_METHOD="${AUTHENTICATION_METHOD}" \
    -e LOCATION="${LOCATION}" \
    -e CUSTOM_CLOUD_CLIENT_ID="${CUSTOM_CLOUD_CLIENT_ID}" \
    -e CUSTOM_CLOUD_SECRET="${CUSTOM_CLOUD_SECRET}" \
    -e PORTAL_ENDPOINT="${PORTAL_ENDPOINT}" \
    -e SERVICE_MANAGEMENT_ENDPOINT="${SERVICE_MANAGEMENT_ENDPOINT}" \
    -e RESOURCE_MANAGER_ENDPOINT="${RESOURCE_MANAGER_ENDPOINT}" \
    -e STORAGE_ENDPOINT_SUFFIX="${STORAGE_ENDPOINT_SUFFIX}" \
    -e KEY_VAULT_DNS_SUFFIX="${KEY_VAULT_DNS_SUFFIX}" \
    -e ACTIVE_DIRECTORY_ENDPOINT="${ACTIVE_DIRECTORY_ENDPOINT}" \
    -e GALLERY_ENDPOINT="${GALLERY_ENDPOINT}" \
    -e GRAPH_ENDPOINT="${GRAPH_ENDPOINT}" \
    -e SERVICE_MANAGEMENT_VM_DNS_SUFFIX="${SERVICE_MANAGEMENT_VM_DNS_SUFFIX}" \
    -e RESOURCE_MANAGER_VM_DNS_SUFFIX="${RESOURCE_MANAGER_VM_DNS_SUFFIX}" \
    ${DEV_IMAGE} make test-kubernetes || exit 1
fi

if [ "${UPGRADE_CLUSTER}" = "true" ]; then
  # modify the master VM SKU to simulate vertical vm scaling via upgrade
  docker run --rm \
      -v $(pwd):${WORK_DIR} \
      -w ${WORK_DIR} \
      -e RESOURCE_GROUP=$RESOURCE_GROUP \
      -e MASTER_VM_UPGRADE_SKU=$MASTER_VM_UPGRADE_SKU \
      ${DEV_IMAGE} \
      /bin/bash -c "jq --arg sku \"$MASTER_VM_UPGRADE_SKU\" '. | .properties.masterProfile.vmSize = \$sku' < _output/$RESOURCE_GROUP/apimodel.json > _output/$RESOURCE_GROUP/apimodel-upgrade.json" || exit 1
  docker run --rm \
      -v $(pwd):${WORK_DIR} \
      -w ${WORK_DIR} \
      -e RESOURCE_GROUP=$RESOURCE_GROUP \
      ${DEV_IMAGE} \
      /bin/bash -c "mv _output/$RESOURCE_GROUP/apimodel-upgrade.json _output/$RESOURCE_GROUP/apimodel.json" || exit 1
  for ver_target in $UPGRADE_VERSIONS; do
    docker run --rm \
      -v $(pwd):${WORK_DIR} \
      -v /etc/ssl/certs:/etc/ssl/certs \
      -w ${WORK_DIR} \
      -e RESOURCE_GROUP=$RESOURCE_GROUP \
      -e REGION=$REGION \
      ${DEV_IMAGE} \
      ./bin/aks-engine upgrade --force \
      --azure-env ${AZURE_ENV} \
      --subscription-id ${AZURE_SUBSCRIPTION_ID} \
      --api-model _output/$RESOURCE_GROUP/apimodel.json \
      --location $REGION \
      --resource-group $RESOURCE_GROUP \
      --upgrade-version $ver_target \
      --vm-timeout 20 \
      --auth-method client_secret \
      --client-id ${AZURE_CLIENT_ID} \
      --client-secret ${AZURE_CLIENT_SECRET} || exit 1

    docker run --rm \
      -v $(pwd):${WORK_DIR} \
      -v /etc/ssl/certs:/etc/ssl/certs \
      -w ${WORK_DIR} \
      -e CLIENT_ID=${AZURE_CLIENT_ID} \
      -e CLIENT_SECRET=${AZURE_CLIENT_SECRET} \
      -e CLIENT_OBJECTID=${CLIENT_OBJECTID} \
      -e TENANT_ID=${AZURE_TENANT_ID} \
      -e SUBSCRIPTION_ID=${AZURE_SUBSCRIPTION_ID} \
      -e INFRA_RESOURCE_GROUP="${INFRA_RESOURCE_GROUP}" \
      -e ORCHESTRATOR=kubernetes \
      -e NAME=$RESOURCE_GROUP \
      -e TIMEOUT=${E2E_TEST_TIMEOUT} \
      -e LB_TIMEOUT=${LB_TEST_TIMEOUT} \
      -e KUBERNETES_IMAGE_BASE=$KUBERNETES_IMAGE_BASE \
      -e KUBERNETES_IMAGE_BASE_TYPE=$KUBERNETES_IMAGE_BASE_TYPE \
      -e CLEANUP_ON_EXIT=false \
      -e REGIONS=$REGION \
      -e IS_JENKINS=${IS_JENKINS} \
      -e SKIP_LOGS_COLLECTION=${SKIP_LOGS_COLLECTION} \
      -e GINKGO_FAIL_FAST=true \
      -e GINKGO_SKIP="${SKIP_AFTER_UPGRADE}" \
      -e GINKGO_FOCUS="${GINKGO_FOCUS}" \
      -e SKIP_TEST=${SKIP_TESTS_AFTER_UPGRADE} \
      -e ADD_NODE_POOL_INPUT=${ADD_NODE_POOL_INPUT} \
      -e API_PROFILE="${API_PROFILE}" \
      -e CUSTOM_CLOUD_NAME="${ENVIRONMENT_NAME}" \
      -e IDENTITY_SYSTEM="${IDENTITY_SYSTEM}" \
      -e AUTHENTICATION_METHOD="${AUTHENTICATION_METHOD}" \
      -e LOCATION="${LOCATION}" \
      -e CUSTOM_CLOUD_CLIENT_ID="${CUSTOM_CLOUD_CLIENT_ID}" \
      -e CUSTOM_CLOUD_SECRET="${CUSTOM_CLOUD_SECRET}" \
      -e PORTAL_ENDPOINT="${PORTAL_ENDPOINT}" \
      -e SERVICE_MANAGEMENT_ENDPOINT="${SERVICE_MANAGEMENT_ENDPOINT}" \
      -e RESOURCE_MANAGER_ENDPOINT="${RESOURCE_MANAGER_ENDPOINT}" \
      -e STORAGE_ENDPOINT_SUFFIX="${STORAGE_ENDPOINT_SUFFIX}" \
      -e KEY_VAULT_DNS_SUFFIX="${KEY_VAULT_DNS_SUFFIX}" \
      -e ACTIVE_DIRECTORY_ENDPOINT="${ACTIVE_DIRECTORY_ENDPOINT}" \
      -e GALLERY_ENDPOINT="${GALLERY_ENDPOINT}" \
      -e GRAPH_ENDPOINT="${GRAPH_ENDPOINT}" \
      -e SERVICE_MANAGEMENT_VM_DNS_SUFFIX="${SERVICE_MANAGEMENT_VM_DNS_SUFFIX}" \
      -e RESOURCE_MANAGER_VM_DNS_SUFFIX="${RESOURCE_MANAGER_VM_DNS_SUFFIX}" \
      ${DEV_IMAGE} make test-kubernetes || exit 1
  done
fi

if [ "${SCALE_CLUSTER}" = "true" ]; then
  for nodepool in $(jq -r '.properties.agentPoolProfiles[].name' < _output/$RESOURCE_GROUP/apimodel.json); do
    docker run --rm \
    -v $(pwd):${WORK_DIR} \
    -v /etc/ssl/certs:/etc/ssl/certs \
    -w ${WORK_DIR} \
    -e RESOURCE_GROUP=$RESOURCE_GROUP \
    -e REGION=$REGION \
    ${DEV_IMAGE} \
    ./bin/aks-engine scale \
    --azure-env ${AZURE_ENV} \
    --subscription-id ${AZURE_SUBSCRIPTION_ID} \
    --api-model _output/$RESOURCE_GROUP/apimodel.json \
    --location $REGION \
    --resource-group $RESOURCE_GROUP \
    --apiserver $API_SERVER \
    --node-pool $nodepool \
    --new-node-count $NODE_COUNT \
    --auth-method client_secret \
    --client-id ${AZURE_CLIENT_ID} \
    --client-secret ${AZURE_CLIENT_SECRET} || exit 1
  done

  docker run --rm \
    -v $(pwd):${WORK_DIR} \
    -v /etc/ssl/certs:/etc/ssl/certs \
    -w ${WORK_DIR} \
    -e CLIENT_ID=${AZURE_CLIENT_ID} \
    -e CLIENT_SECRET=${AZURE_CLIENT_SECRET} \
    -e CLIENT_OBJECTID=${CLIENT_OBJECTID} \
    -e TENANT_ID=${AZURE_TENANT_ID} \
    -e SUBSCRIPTION_ID=${AZURE_SUBSCRIPTION_ID} \
    -e INFRA_RESOURCE_GROUP="${INFRA_RESOURCE_GROUP}" \
    -e ORCHESTRATOR=kubernetes \
    -e NAME=$RESOURCE_GROUP \
    -e TIMEOUT=${E2E_TEST_TIMEOUT} \
    -e LB_TIMEOUT=${LB_TEST_TIMEOUT} \
    -e KUBERNETES_IMAGE_BASE=$KUBERNETES_IMAGE_BASE \
    -e KUBERNETES_IMAGE_BASE_TYPE=$KUBERNETES_IMAGE_BASE_TYPE \
    -e CLEANUP_ON_EXIT=${CLEANUP_ON_EXIT} \
    -e REGIONS=$REGION \
    -e IS_JENKINS=${IS_JENKINS} \
    -e SKIP_LOGS_COLLECTION=${SKIP_LOGS_COLLECTION} \
    -e GINKGO_FAIL_FAST=true \
    -e GINKGO_SKIP="${SKIP_AFTER_SCALE_UP}" \
    -e GINKGO_FOCUS="${GINKGO_FOCUS}" \
    -e SKIP_TEST=${SKIP_TESTS_AFTER_SCALE_UP} \
    -e ADD_NODE_POOL_INPUT=${ADD_NODE_POOL_INPUT} \
    -e API_PROFILE="${API_PROFILE}" \
    -e CUSTOM_CLOUD_NAME="${ENVIRONMENT_NAME}" \
    -e IDENTITY_SYSTEM="${IDENTITY_SYSTEM}" \
    -e AUTHENTICATION_METHOD="${AUTHENTICATION_METHOD}" \
    -e LOCATION="${LOCATION}" \
    -e CUSTOM_CLOUD_CLIENT_ID="${CUSTOM_CLOUD_CLIENT_ID}" \
    -e CUSTOM_CLOUD_SECRET="${CUSTOM_CLOUD_SECRET}" \
    -e PORTAL_ENDPOINT="${PORTAL_ENDPOINT}" \
    -e SERVICE_MANAGEMENT_ENDPOINT="${SERVICE_MANAGEMENT_ENDPOINT}" \
    -e RESOURCE_MANAGER_ENDPOINT="${RESOURCE_MANAGER_ENDPOINT}" \
    -e STORAGE_ENDPOINT_SUFFIX="${STORAGE_ENDPOINT_SUFFIX}" \
    -e KEY_VAULT_DNS_SUFFIX="${KEY_VAULT_DNS_SUFFIX}" \
    -e ACTIVE_DIRECTORY_ENDPOINT="${ACTIVE_DIRECTORY_ENDPOINT}" \
    -e GALLERY_ENDPOINT="${GALLERY_ENDPOINT}" \
    -e GRAPH_ENDPOINT="${GRAPH_ENDPOINT}" \
    -e SERVICE_MANAGEMENT_VM_DNS_SUFFIX="${SERVICE_MANAGEMENT_VM_DNS_SUFFIX}" \
    -e RESOURCE_MANAGER_VM_DNS_SUFFIX="${RESOURCE_MANAGER_VM_DNS_SUFFIX}" \
    ${DEV_IMAGE} make test-kubernetes || exit 1
fi