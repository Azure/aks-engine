#!/bin/bash

set -e

# some tests set NEW_AGENT_NODE_COUNT in .env files
ENV_FILE="${CLUSTER_DEFINITION}.env"
if [ -e "${ENV_FILE}" ]; then
  source "${ENV_FILE}"
fi

[[ ! -z "${NEW_AGENT_NODE_COUNT:-}" ]] || (echo "Must specify NEW_AGENT_NODE_COUNT" && exit 1)

OUTPUT="_output/${INSTANCE_NAME}"

# allow nodes to run for a while before scaling
sleep 180

./bin/aks-engine scale \
  --subscription-id ${SUBSCRIPTION_ID} \
  --deployment-dir ${OUTPUT} \
  --location ${LOCATION} \
  --resource-group ${RESOURCE_GROUP} \
  --master-FQDN "${INSTANCE_NAME}.${LOCATION}.cloudapp.azure.com" \
  --node-pool "agentpool1" \
  --new-node-count ${NEW_AGENT_NODE_COUNT} \
  --auth-method client_secret \
  --client-id ${SERVICE_PRINCIPAL_CLIENT_ID} \
  --client-secret ${SERVICE_PRINCIPAL_CLIENT_SECRET}
