#!/bin/bash

# Description:
# This script was created to help simplify the migration process for applications which use persistent volumes. 
# It can be used on deployment and statefulsets. 
# The script will migrate all of the deployment/statefulset's associated PersistentVolumeClaims and PersistentVolumes 
# from "kubernetes.io/azure-disk" provisioner to "disk.csi.azure.com" provisioner. 
# Then it will recreate the deployment/statefulset to reference the newly migrated PersistentVolumeClaims and PersistentVolumes.

# How to use:
# ./migratepv.sh -t ResourceType -r ResourceName -n Namespace
# ./migratepv.sh -t StatefulSet -r web -n default
# ./migratepv.sh -t Deployment -r depa -n dev

# Input: ResourceType, ResourceName, Namespace
while getopts t:r:n: flag
do
    case "${flag}" in
        t) ResourceType=${OPTARG};;
        r) ResourceName=${OPTARG};;
        n) Namespace=${OPTARG};;
    esac
done
echo "ResourceType: $ResourceType";
echo "ResourceName: $ResourceName";
echo "Namespace: $Namespace";

# Check input is valid
if [[ $ResourceType != "StatefulSet" && $ResourceType != "Deployment" ]]; then
    echo "The ResourceType $ResourceType cannot be migrated. The accepted resource types are StatefulSet and Deployment."
    exit 1
fi

if [[ -z $ResourceType || -z $ResourceName || -z $Namespace || -z $(kubectl get $ResourceType $ResourceName -n $Namespace) ]]; then
    echo "The given $ResourceType $ResourceName could not be found in $Namespace."
    exit 1
fi
echo "Migrating $ResourceType $ResourceName in namespace $Namespace"

# Save the resource information
kubectl get $ResourceType $ResourceName -n $Namespace -o json > backup-$ResourceType-$ResourceName-$Namespace.json
echo "Saved current resource information to backup-$ResourceType-$ResourceName-$Namespace.json"

# Check that resource's each referenced storage classe exists in current storage classes and has provisioner AzureDisk CSI Driver
declare -a StorageClassList=()
if [[ $ResourceType == "Deployment" ]]; then 
    pvcList=$(jq -r '.spec.template.spec.volumes[].persistentVolumeClaim.claimName' backup-$ResourceType-$ResourceName-$Namespace.json)
    for PVC in ${pvcList[@]}
    do
        StorageClassName=$(kubectl get pvc $PVC -n $Namespace -o json | jq -r '.spec.storageClassName')
        StorageClassList+=($StorageClassName)
    done
elif [[ $ResourceType == "StatefulSet"  ]]; then
    StorageClassList=$(jq -r '.spec.volumeClaimTemplates[].spec.storageClassName' backup-$ResourceType-$ResourceName-$Namespace.json)
fi

for StorageClassName in ${StorageClassList[@]}
do
    if [[ -z $(kubectl get storageclass $StorageClassName) ]]; then
        echo "Storage class could not be found which matches storageClassName $StorageClassName in $ResourceType $ResourceName. Cannot migrate pvc + pvs. Ending script."
        exit 1
    fi
    if [[ "$(kubectl get storageclass $StorageClassName -o json | jq -r '.provisioner')" != "disk.csi.azure.com" ]]; then 
        echo "Storage class with name $StorageClassName is currently not provisioned with AzureDisk CSI Driver." 
        echo "Please first follow instructions here to install AzureDisk CSI Driver storage classes: https://github.com/Azure/aks-engine/blob/master/docs/topics/azure-stack.md#upgrade-from-kubernetes-v120-to-v121-on-azure-stack-hub"
        exit 1
    fi
done
echo "Verified AzureDisk CSI Driver storage classes exist. Good to proceed."

# Find all pvcs associated with this resource
declare -a pvcList=()
if [[ $ResourceType == "Deployment" ]]; then
    pvcList=$(jq -r '.spec.template.spec.volumes[].persistentVolumeClaim.claimName' backup-$ResourceType-$ResourceName-$Namespace.json)
elif [[ $ResourceType == "StatefulSet"  ]]; then
    # If statefulset, find pvcs which match volumeclaimtemplatename-statefulsetname-*
    VolumeClaimTemplateList=$(jq -r '.spec.volumeClaimTemplates[].metadata.name' backup-$ResourceType-$ResourceName-$Namespace.json)
    StatefulSetName=$(jq -r '.metadata.name' backup-$ResourceType-$ResourceName-$Namespace.json)
    for volumeClaimTemplateName in $VolumeClaimTemplateList
    do
        pvcs=$(kubectl get pvc -n $Namespace -o name | grep "$volumeClaimTemplateName-$StatefulSetName-.*" | sed -e "s/^persistentvolumeclaim\///")
        pvcList+=(${pvcs[@]})
    done
fi

if [ ${#pvcList[@]} -eq 0 ]; then
    echo "No pvcs are associated with this statefulset, no need to migrate. Ending script."
    exit 0
else
    echo "Beginning migration of each pvc + pv pair for $ResourceType $ResourceName."
fi

# Delete the resource. Pods must be deleted for pvc + pv to be migrated.
kubectl delete $ResourceType $ResourceName -n $Namespace

for PVC in ${pvcList[@]}
do
    # for each pvc, get its pv
    PV=$(kubectl get pvc $PVC -n $Namespace -o json | jq -r '.spec.volumeName')

    # step 0: save the current pvc and pv information
    kubectl get pvc $PVC -n $Namespace -o json > backup-pvc-$PVC.json
    echo "Saved current pvc information to backup-pv-$PVC.json"
    kubectl get pv $PV -o json > backup-pv-$PV.json
    echo "Saved current pv information to backup-pv-$PV.json"
    # note down the DiskURI and FSType. They will be used later in step 3.
    PV_DiskURI=$(jq -r '.spec.azureDisk.diskURI' backup-pv-$PV.json)
    PV_FSType=$(jq -r '.spec.azureDisk.fSType' backup-pv-$PV.json)

    # step 1: retain the original physical AzureDisk
    kubectl patch pv $PV  -p '{"spec":{"persistentVolumeReclaimPolicy":"Retain"}}'

    # step 2: delete the original pvc + pv reference
    kubectl delete pvc $PVC -n $Namespace
    kubectl delete pv $PV

    # step 3: create a new pv which references the original physical AzureDisk
    cp backup-pv-$PV.json new-pv-$PV.json
    # remove fields not needed for migration
    newPV="$(jq 'del(.metadata.annotations, .metadata.creationTimestamp, .metadata.managedFields, .metadata.resourceVersion, .metadata.uid, .status, .spec.claimRef.resourceVersion, .spec.claimRef.uid, .spec.azureDisk)' new-pv-$PV.json)"
    echo "$newPV" > new-pv-$PV.json
    # add .spec.csi with references to the original physical AzureDisk
    driver="disk.csi.azure.com"
    if [[ $PV_FSType == "null" ]]; then
        newPV="$(jq ".spec += {"csi": {"driver": \"$driver\", "volumeHandle": \"$PV_DiskURI\"}}" new-pv-$PV.json)"
    else
        newPV="$(jq ".spec += {"csi": {"driver": \"$driver\", "volumeHandle": \"$PV_DiskURI\", "fsType": \"$PV_FSType\"}}" new-pv-$PV.json)"
    fi
    echo "$newPV" > new-pv-$PV.json

    # apply new pv 
    kubectl apply -f new-pv-$PV.json --validate=false

    # check new pv is healthy
    if [[ "$(kubectl get pv $PV -o json | jq -r '.status.phase')" == "Available" ]]; then 
        echo "PV $PV has been successfully migrated to out of tree AzureDisk CSI Driver provisioner."
    else
        echo "PV $PV has failed to migrate to out of tree. Exiting script." 
        exit 1
    fi

    # step 4: if deployment, create a new pvc. (statefulset will create the pvc itself)
    if [[ $ResourceType == "Deployment" ]]; then
        cp backup-pvc-$PVC.json new-pvc-$PVC.json
        # remove fields not needed for migration
        newPVC="$(jq 'del(.metadata.annotations, .metadata.creationTimestamp, .metadata.resourceVersion, .metadata.uid, .status)' new-pvc-$PVC.json)"
        echo "$newPVC" > new-pvc-$PVC.json

        # apply new pvc
        kubectl apply -f new-pvc-$PVC.json --validate=false

        # check new pvc exists
        if [[ "$(kubectl get pvc $PVC -n $Namespace -o json | jq -r '.status.phase')" ]]; then 
            echo "PVC $PVC has been successfully migrated to out of tree AzureDisk CSI Driver provisioner."
        else
            echo "PVC $PVC has failed to migrate to out of tree. Exiting script." 
            exit 1
        fi
    fi
done

# Recreate the resource
cp backup-$ResourceType-$ResourceName-$Namespace.json new-$ResourceType-$ResourceName-$Namespace.json
# remove fields not needed for recreation
if [[ $ResourceType == "Deployment" ]]; then 
    newResource="$(jq 'del(.metadata.annotations, .metadata.creationTimestamp, .metadata.generation, .metadata.resourceVersion, .metadata.uid, .status, .spec.template.metadata.creationTimestamp)' new-$ResourceType-$ResourceName-$Namespace.json)"
elif [[ $ResourceType == "StatefulSet"  ]]; then
    newResource="$(jq 'del(.metadata.annotations, .metadata.creationTimestamp, .metadata.generation, .metadata.managedFields, .metadata.resourceVersion, .metadata.uid, .status, .spec.template.metadata.creationTimestamp, .spec.volumeClaimTemplates[].metadata.creationTimestamp, .spec.volumeClaimTemplates[].status)' new-$ResourceType-$ResourceName-$Namespace.json)"
fi
echo "$newResource" > new-$ResourceType-$ResourceName-$Namespace.json

# apply new resource
kubectl apply -f new-$ResourceType-$ResourceName-$Namespace.json --validate=false

# check new resource exists
if [[ $(kubectl get $ResourceType $ResourceName -n $Namespace) ]]; then 
    echo "All pvs in $ResourceType $ResourceName have been successfully migrated to out of tree AzureDisk CSI Driver provisioner. $ResourceType $ResourceName has been redeployed, please monitor its progress."
else
    echo "$ResourceType $ResourceName failed to redeploy. Exiting script." 
    exit 1
fi

echo "Migration of pvc + pv has completed for $ResourceType $ResourceName."