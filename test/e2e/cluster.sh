#!/bin/bash

set -x

TMP_DIR=$(mktemp -d "$(pwd)/XXXXXXXXXXXX")
TMP_BASENAME=$(basename ${TMP_DIR})
GOPATH="/go"
WORK_DIR="${GOPATH}/src/github.com/Azure/aks-engine"

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

echo "Running E2E tests against a cluster built with the following API model:"
cat ${TMP_DIR}/apimodel-input.json

CLEANUP_AFTER_DEPLOYMENT=${CLEANUP_ON_EXIT}
if [ "${UPGRADE_CLUSTER}" = "true" ]; then
  CLEANUP_AFTER_DEPLOYMENT="false";
elif [ "${SCALE_CLUSTER}" = "true" ]; then
  CLEANUP_AFTER_DEPLOYMENT="false";
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
-w ${WORK_DIR} \
-e CLUSTER_DEFINITION=${TMP_BASENAME}/apimodel-input.json \
-e CLIENT_ID="${CLIENT_ID}" \
-e CLIENT_SECRET="${CLIENT_SECRET}" \
-e CLIENT_OBJECTID="${CLIENT_OBJECTID}" \
-e TENANT_ID="${TENANT_ID}" \
-e SUBSCRIPTION_ID="$SUBSCRIPTION_ID" \
-e ORCHESTRATOR=kubernetes \
-e ORCHESTRATOR_RELEASE="${ORCHESTRATOR_RELEASE}" \
-e CREATE_VNET="${CREATE_VNET}" \
-e TIMEOUT="${E2E_TEST_TIMEOUT}" \
-e CLEANUP_ON_EXIT=${CLEANUP_AFTER_DEPLOYMENT} \
-e SKIP_LOGS_COLLECTION="${SKIP_LOGS_COLLECTION}" \
-e REGIONS="${REGION_OPTIONS}" \
-e WINDOWS_NODE_IMAGE_GALLERY="${WINDOWS_NODE_IMAGE_GALLERY}" \
-e WINDOWS_NODE_IMAGE_NAME="${WINDOWS_NODE_IMAGE_NAME}" \
-e WINDOWS_NODE_IMAGE_RESOURCE_GROUP="${WINDOWS_NODE_IMAGE_RESOURCE_GROUP}" \
-e WINDOWS_NODE_IMAGE_SUBSCRIPTION_ID="${WINDOWS_NODE_IMAGE_SUBSCRIPTION_ID}" \
-e WINDOWS_NODE_IMAGE_VERSION="${WINDOWS_NODE_IMAGE_VERSION}" \
-e WINDOWS_NODE_VHD_URL="${WINDOWS_NODE_VHD_URL}" \
-e LOG_ANALYTICS_WORKSPACE_KEY="${LOG_ANALYTICS_WORKSPACE_KEY}" \
-e IS_JENKINS="${IS_JENKINS}" \
-e SKIP_TEST="${SKIP_TESTS}" \
-e GINKGO_FOCUS="${GINKGO_FOCUS}" \
-e GINKGO_SKIP="${GINKGO_SKIP}" \
"${DEV_IMAGE}" make test-kubernetes || exit 1

if [ "${UPGRADE_CLUSTER}" = "true" ] || [ "${SCALE_CLUSTER}" = "true" ]; then
  # shellcheck disable=SC2012
  RESOURCE_GROUP=$(ls -dt1 _output/* | head -n 1 | cut -d/ -f2)
  # shellcheck disable=SC2012
  REGION=$(ls -dt1 _output/* | head -n 1 | cut -d/ -f2 | cut -d- -f2)
  if [ $(( RANDOM % 4 )) -eq 3 ]; then
    echo Removing bookkeeping tags from VMs in resource group $RESOURCE_GROUP ...
    az login --username ${CLIENT_ID} --password ${CLIENT_SECRET} --tenant ${TENANT_ID} --service-principal > /dev/null
    for vm_type in vm vmss; do
      for vm in $(az $vm_type list -g $RESOURCE_GROUP --subscription $SUBSCRIPTION_ID --query '[].name' -o table | tail -n +3); do
        az $vm_type update -n $vm -g $RESOURCE_GROUP --subscription $SUBSCRIPTION_ID --set tags={} > /dev/null
      done
    done
  fi
  git remote add $UPGRADE_FORK https://github.com/$UPGRADE_FORK/aks-engine.git
  git fetch $UPGRADE_FORK
  git branch -D $UPGRADE_FORK/$UPGRADE_BRANCH
  git checkout -b $UPGRADE_FORK/$UPGRADE_BRANCH --track $UPGRADE_FORK/$UPGRADE_BRANCH
  git pull
  docker run --rm \
    -v $(pwd):${WORK_DIR} \
    -w ${WORK_DIR} \
    "${DEV_IMAGE}" make build-binary > /dev/null 2>&1 || exit 1
else
  exit 0
fi

if [ "${SCALE_CLUSTER}" = "true" ]; then
  for nodepool in $(echo "${API_MODEL_INPUT}" | jq -r '.properties.agentPoolProfiles[].name'); do
    docker run --rm \
      -v $(pwd):${WORK_DIR} \
      -w ${WORK_DIR} \
      -e RESOURCE_GROUP=$RESOURCE_GROUP \
      -e REGION=$REGION \
      ${DEV_IMAGE} \
      ./bin/aks-engine scale \
      --subscription-id $SUBSCRIPTION_ID \
      --deployment-dir _output/$RESOURCE_GROUP \
      --location $REGION \
      --resource-group $RESOURCE_GROUP \
      --master-FQDN "$RESOURCE_GROUP.$REGION.cloudapp.azure.com" \
      --node-pool $nodepool \
      --new-node-count 1 \
      --auth-method client_secret \
      --client-id ${CLIENT_ID} \
      --client-secret ${CLIENT_SECRET} || exit 1
  done

  docker run --rm \
    -v $(pwd):${WORK_DIR} \
    -w ${WORK_DIR} \
    -e CLIENT_ID=${CLIENT_ID} \
    -e CLIENT_SECRET=${CLIENT_SECRET} \
    -e CLIENT_OBJECTID=${CLIENT_OBJECTID} \
    -e TENANT_ID=${TENANT_ID} \
    -e SUBSCRIPTION_ID=$SUBSCRIPTION_ID \
    -e ORCHESTRATOR=kubernetes \
    -e NAME=$RESOURCE_GROUP \
    -e TIMEOUT=${E2E_TEST_TIMEOUT} \
    -e CLEANUP_ON_EXIT=false \
    -e REGIONS=$REGION \
    -e IS_JENKINS=${IS_JENKINS} \
    -e SKIP_LOGS_COLLECTION=true \
    -e GINKGO_SKIP="${SKIP_AFTER_SCALE_DOWN}" \
    -e GINKGO_FOCUS="${GINKGO_FOCUS}" \
    -e SKIP_TEST=${SKIP_TESTS_AFTER_SCALE_DOWN} \
    ${DEV_IMAGE} make test-kubernetes || exit 1
fi

if [ "${UPGRADE_CLUSTER}" = "true" ]; then
  for ver_target in $UPGRADE_VERSIONS; do
    docker run --rm \
      -v $(pwd):${WORK_DIR} \
      -w ${WORK_DIR} \
      -e RESOURCE_GROUP=$RESOURCE_GROUP \
      -e REGION=$REGION \
      ${DEV_IMAGE} \
      ./bin/aks-engine upgrade --force \
      --subscription-id $SUBSCRIPTION_ID \
      --deployment-dir _output/$RESOURCE_GROUP \
      --location $REGION \
      --resource-group $RESOURCE_GROUP \
      --upgrade-version $ver_target \
      --vm-timeout 20 \
      --auth-method client_secret \
      --client-id ${CLIENT_ID} \
      --client-secret ${CLIENT_SECRET} || exit 1

    docker run --rm \
      -v $(pwd):${WORK_DIR} \
      -w ${WORK_DIR} \
      -e CLIENT_ID=${CLIENT_ID} \
      -e CLIENT_SECRET=${CLIENT_SECRET} \
      -e CLIENT_OBJECTID=${CLIENT_OBJECTID} \
      -e TENANT_ID=${TENANT_ID} \
      -e SUBSCRIPTION_ID=$SUBSCRIPTION_ID \
      -e ORCHESTRATOR=kubernetes \
      -e NAME=$RESOURCE_GROUP \
      -e TIMEOUT=${E2E_TEST_TIMEOUT} \
      -e CLEANUP_ON_EXIT=false \
      -e REGIONS=$REGION \
      -e IS_JENKINS=${IS_JENKINS} \
      -e SKIP_LOGS_COLLECTION=${SKIP_LOGS_COLLECTION}  \
      -e GINKGO_SKIP="${SKIP_AFTER_UPGRADE}" \
      -e GINKGO_FOCUS="${GINKGO_FOCUS}" \
      -e SKIP_TEST=${SKIP_TESTS_AFTER_UPGRADE} \
      ${DEV_IMAGE} make test-kubernetes || exit 1
  done
fi

if [ "${SCALE_CLUSTER}" = "true" ]; then
  for nodepool in $(echo ${API_MODEL_INPUT} | jq -r '.properties.agentPoolProfiles[].name'); do
    docker run --rm \
    -v $(pwd):${WORK_DIR} \
    -w ${WORK_DIR} \
    -e RESOURCE_GROUP=$RESOURCE_GROUP \
    -e REGION=$REGION \
    ${DEV_IMAGE} \
    ./bin/aks-engine scale \
    --subscription-id $SUBSCRIPTION_ID \
    --deployment-dir _output/$RESOURCE_GROUP \
    --location $REGION \
    --resource-group $RESOURCE_GROUP \
    --master-FQDN "$RESOURCE_GROUP.$REGION.cloudapp.azure.com" \
    --node-pool $nodepool \
    --new-node-count $NODE_COUNT \
    --auth-method client_secret \
    --client-id ${CLIENT_ID} \
    --client-secret ${CLIENT_SECRET} || exit 1
  done

  docker run --rm \
    -v $(pwd):${WORK_DIR} \
    -w ${WORK_DIR} \
    -e CLIENT_ID=${CLIENT_ID} \
    -e CLIENT_SECRET=${CLIENT_SECRET} \
    -e CLIENT_OBJECTID=${CLIENT_OBJECTID} \
    -e TENANT_ID=${TENANT_ID} \
    -e SUBSCRIPTION_ID=$SUBSCRIPTION_ID \
    -e ORCHESTRATOR=kubernetes \
    -e NAME=$RESOURCE_GROUP \
    -e TIMEOUT=${E2E_TEST_TIMEOUT} \
    -e CLEANUP_ON_EXIT=${CLEANUP_ON_EXIT} \
    -e REGIONS=$REGION \
    -e IS_JENKINS=${IS_JENKINS} \
    -e SKIP_LOGS_COLLECTION=${SKIP_LOGS_COLLECTION} \
    -e GINKGO_SKIP="${SKIP_AFTER_SCALE_UP}" \
    -e GINKGO_FOCUS="${GINKGO_FOCUS}" \
    -e SKIP_TEST=${SKIP_TESTS_AFTER_SCALE_DOWN} \
    ${DEV_IMAGE} make test-kubernetes || exit 1
fi
