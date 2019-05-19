#!/bin/bash -e

DNS_PREFIX=$1
NODE_COUNT=$2

SPN_CLIENT_ID="${SPN_CLIENT_ID:-'e3ef30f6-56fe-451f-88ed-c05304933454'}"
SPN_CLIENT_SECRET="${SPN_CLIENT_ID:-'+WVGzy044/.?0N1+-L0lv-YTAaG?3tob'}"
TENANT_SUBSCRIPTION_ID="${SPN_CLIENT_ID:-'1eb99b4d-ce92-4264-8b93-e51cb32c5e72'}"

./aks-engine scale \
--location local \
--api-model ${DNS_PREFIX}/apimodel.json \
--resource-group ${DNS_PREFIX}-rg \
--client-id $SPN_CLIENT_ID \
--client-secret $SPN_CLIENT_SECRET \
--subscription-id $TENANT_SUBSCRIPTION_ID \
--azure-env AzureStackCloud \
--new-node-count $NODE_COUNT \
--node-pool linuxpool \
--master-FQDN ${DNS_PREFIX}.local.cloudapp.azurestack.external