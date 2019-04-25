#!/bin/bash

set -e

# some tests set EXPECTED_ORCHESTRATOR_VERSION in .env files
ENV_FILE="${CLUSTER_DEFINITION}.env"
if [ -e "${ENV_FILE}" ]; then
  source "${ENV_FILE}"
fi

[[ ! -z "${EXPECTED_ORCHESTRATOR_VERSION:-}" ]] || (echo "Must specify EXPECTED_ORCHESTRATOR_VERSION" && exit 1)

APIMODEL="_output/${INSTANCE_NAME}/apimodel.json"

./bin/aks-engine upgrade \
  --subscription-id ${SUBSCRIPTION_ID} \
  --api-model ${APIMODEL} \
  --location ${LOCATION} \
  --resource-group ${RESOURCE_GROUP} \
  --upgrade-version ${EXPECTED_ORCHESTRATOR_VERSION} \
  --auth-method client_secret \
  --client-id ${SERVICE_PRINCIPAL_CLIENT_ID} \
  --client-secret ${SERVICE_PRINCIPAL_CLIENT_SECRET}
