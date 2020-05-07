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

if [ -z "$DRY_RUN" ]; then
    echo "no DRY_RUN env var provided, will default to true"
    DRY_RUN=true
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

# clean up Packer resource groups

# find packer resource groups created before our deadline
echo "Looking for resource groups created over ${EXPIRATION_IN_HOURS} hours ago..."
for resourceGroup in $( az group list --query "[?contains(name, 'packer-Resource-Group')]" | jq --arg dl $deadline '.[] | select(.tags.now < $dl).name' | tr -d '\"' || ""); do
    for deployment in $(az deployment group list -g $resourceGroup | jq '.[] | .name' | tr -d '\"' || ""); do
        echo "Will delete deployment ${deployment} from resource group ${resourceGroup}..."
        if [[ "${DRY_RUN}" = false ]]; then
            az deployment group delete -n $deployment -g $resourceGroup || echo "unable to delete deployment ${deployment}, will continue..."
        else
            echo "skipping because DRY_RUN is set to true"
        fi
    done
    echo "Will delete resource group ${resourceGroup}..."
    # delete old resource groups
    if [[ "${DRY_RUN}" = false ]]; then
        az group delete -y -n $resourceGroup --no-wait >> delete.log || echo "unable to delete resource group ${resourceGroup}, will continue..."
    else
            echo "skipping because DRY_RUN is set to true"
    fi
done

# cleanup storage accounts. We only want to delete storage accounts created by Packer, NOT the classic storage account that is used by the Marketplace to publish the VHDs as images.
if [ -z "$STORAGE_RG" ]; then
    echo "must provide a STORAGE_RG env var"
    exit 1;
fi

echo "Looking for storage accounts in ${STORAGE_RG} created over ${EXPIRATION_IN_HOURS} hours ago..."
for storage_account in $(az storage account list -g $STORAGE_RG | jq --arg dl $deadline '.[] | select(.tags.now < $dl).name' | tr -d '\"' || ""); do
    echo "Will delete storage account ${storage_account} from resource group ${STORAGE_RG}..."
    if [[ "${DRY_RUN}" = false ]]; then
        az storage account delete -y -n $storage_account -g $STORAGE_RG || echo "unable to delete storage account ${storage_account}, will continue..."
    else
            echo "skipping because DRY_RUN is set to true"
    fi
done
