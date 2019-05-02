#!/bin/bash

az network vnet create -g ${RESOURCE_GROUP} -n KubernetesCustomVNET --address-prefixes 10.239.0.0/16 --subnet-name KubernetesSubnet --subnet-prefix 10.239.0.0/16

tempfile="$(mktemp)"
trap 'rm -rf "${tempfile}"' EXIT

jq ".properties.masterProfile.vnetSubnetId = \"/subscriptions/${SUBSCRIPTION_ID}/resourceGroups/${RESOURCE_GROUP}/providers/Microsoft.Network/virtualNetworks/KubernetesCustomVNET/subnets/KubernetesSubnet\"" ${CLUSTER_DEFINITION} > $tempfile && mv $tempfile ${CLUSTER_DEFINITION}

indx=0
for poolname in $(jq -r '.properties.agentPoolProfiles[].name' "${CLUSTER_DEFINITION}"); do
  jq ".properties.agentPoolProfiles[$indx].vnetSubnetId = \"/subscriptions/${SUBSCRIPTION_ID}/resourceGroups/${RESOURCE_GROUP}/providers/Microsoft.Network/virtualNetworks/KubernetesCustomVNET/subnets/KubernetesSubnet\"" ${CLUSTER_DEFINITION} > $tempfile && mv $tempfile ${CLUSTER_DEFINITION}
  indx=$((indx+1))
done
