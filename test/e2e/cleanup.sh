#!/bin/bash

####################################################

if [ -z "$SERVICE_PRINCIPAL_CLIENT_ID" ]; then
    echo "must provide a SERVICE_PRINCIPAL_CLIENT_ID env var"
    exit 1;
fi

if [ -z "$SERVICE_PRINCIPAL_CLIENT_SECRET" ]; then
    echo "must provide a SERVICE_PRINCIPAL_CLIENT_SECRET env var"
    exit 1;
fi

if [ -z "$TENANT_ID" ]; then
    echo "must provide a TENANT_ID env var"
    exit 1;
fi

if [ -z "$SUBSCRIPTION_ID_TO_CLEANUP" ]; then
    echo "must provide a SUBSCRIPTION_ID_TO_CLEANUP env var"
    exit 1;
fi

if [ -z "$EXPIRATION_IN_HOURS" ]; then
    EXPIRATION_IN_HOURS=2
fi

az login --service-principal \
		--username "${SERVICE_PRINCIPAL_CLIENT_ID}" \
		--password "${SERVICE_PRINCIPAL_CLIENT_SECRET}" \
		--tenant "${TENANT_ID}" &>/dev/null

# set to the sub id we want to cleanup
az account set -s $SUBSCRIPTION_ID_TO_CLEANUP

# convert to seconds so we can compare it against the "tags.now" property in the resource group metadata
(( expirationInSecs = ${EXPIRATION_IN_HOURS} * 60 * 60 ))
# deadline = the "date +%s" representation of the oldest age we're willing to keep
(( deadline=$(date +%s)-${expirationInSecs%.*} ))
# find resource groups created before our deadline
echo "Looking for resource groups created over ${EXPIRATION_IN_HOURS} hours ago..."
if [ -z "$RESOURCE_GROUP_SUBSTRING" ]; then
  for resourceGroup in $(az group list | jq --arg dl $deadline '.[] | select(.name | contains("acse-test") | not) | select(.tags.now < $dl).name' | tr -d '\"' || ""); do
    for deployment in $(az deployment group list -g $resourceGroup | jq '.[] | .name' | tr -d '\"' || ""); do
      echo "Will delete deployment ${deployment} from resource group ${resourceGroup}..."
      az deployment group delete -n $deployment -g $resourceGroup || echo "unable to delete deployment ${deployment}, will continue..."
    done
    echo "Will delete resource group ${resourceGroup}..."
    # delete old resource groups
    az group delete -y -n $resourceGroup --no-wait >> delete.log || echo "unable to delete resource group ${resourceGroup}, will continue..."
  done
else
  for resourceGroup in $(az group list | jq --arg dl $deadline --arg rg $RESOURCE_GROUP_SUBSTRING '.[] | select(.name | contains("acse-test") | not) | select(.name | contains($rg)) | select(.tags.now < $dl).name' | tr -d '\"' || ""); do
    for deployment in $(az deployment group list -g $resourceGroup | jq '.[] | .name' | tr -d '\"' || ""); do
      echo "Will delete deployment ${deployment} from resource group ${resourceGroup}..."
      az deployment group delete -n $deployment -g $resourceGroup || echo "unable to delete deployment ${deployment}, will continue..."
    done
    echo "Will delete resource group ${resourceGroup}..."
    # delete old resource groups
    az group delete -y -n $resourceGroup --no-wait >> delete.log || echo "unable to delete resource group ${resourceGroup}, will continue..."
  done
fi
