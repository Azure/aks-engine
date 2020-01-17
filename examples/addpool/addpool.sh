#! /bin/bash
set -x

group=$1
subscription=$2
location=$3

az group create --name $group --location $3 --subscription $2

# Create role
az ad sp create-for-rbac --role="Contributor" --scopes="/subscriptions/$subscription/resourceGroups/$group" >~/.kube/$group-sp.json
# shellcheck disable=SC2002,SC2155
export appId=$(cat ~/.kube/$group-sp.json | jq -r .appId)
# shellcheck disable=SC2002,SC2155
export password=$(cat ~/.kube/$group-sp.json | jq -r .password)

# give the sp time to propagate
sleep 180

# Deploy Cluster
aks-engine deploy --subscription-id $subscription --resource-group $group --location $location us --api-model apimodel.json --dns-prefix $group --client-id $appId --client-secret $password --set servicePrincipalProfile.clientId=$appId --set servicePrincipalProfile.secret=$password

cp $(pwd)/_output/$group/kubeconfig/kubeconfig.eastus.json ~/.kube/$group.json

sleep 180

aks-engine addpool --subscription-id $subscription --resource-group $group --location $location us --api-model _output/$group/apimodel.json --agent-pool agentpool.json --client-id $appId --client-secret $password

az vmss list -g $group --subscription $subscription -o table

KUBECONFIG=~/.kube/$group.json kubectl get nodes

#cleanup
#az vmss delete -n $(az vmss list -g $group --subscription $subscription -o json | jq .[1].name) -g $group --subscription $subscription --no-wait

#az group delete -g $group --subscription $subscription --no-wait
