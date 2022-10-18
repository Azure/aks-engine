#!/bin/bash

KIND= # allowed values "deployment" or "statefulset"
NS= # namespace name of the deployment/statefulset, ex: "default"
NAME= # resource name of the deployment/statefulset
PVC_LIST= # PVCs to migrate, ex: "pvc1-name pvc2-name"
TMP=$(mktemp -d)

mkdir -p ${TMP}/${KIND} ${TMP}/pvc ${TMP}/pv

# Back kind/name, remove ID
ID_PROPS=".metadata.annotations, .metadata.creationTimestamp, .metadata.generation, .metadata.resourceVersion, .metadata.uid, .metadata.managedFields, .status, .spec.template.metadata.creationTimestamp"
if [[ "${KIND,,}" == "statefulset" ]]; then 
    ID_PROPS="${ID_PROPS}, .spec.volumeClaimTemplates[].metadata.creationTimestamp, .spec.volumeClaimTemplates[].status"
fi

kubectl get ${KIND}/${NAME} -n ${NS} -o json | \
    jq "del(${ID_PROPS})" \
    > ${TMP}/${KIND}/${NS}-${NAME}.json

# Delete kind/name
kubectl delete -n ${NS} ${KIND}/${NAME}

for pvc in ${PVC_LIST}
do
    # Get PV data from PVC
    PV_NAME=$(kubectl get pvc ${pvc} -o jsonpath='{.spec.volumeName}')
    PV_DISK_URI=$(kubectl get pv ${PV_NAME} -o jsonpath='{.spec.azureDisk.diskURI}')
    PV_FS_TYPE=$(kubectl get pv ${PV_NAME} -o jsonpath='{.spec.azureDisk.fSType}')

    # Back PVC, remove ID
    kubectl get pvc ${pvc} -n ${NS} -o json | \
        jq 'del(.metadata.annotations, .metadata.creationTimestamp, .metadata.resourceVersion, .metadata.uid, .status)' \
        > ${TMP}/pvc/${NS}-${pvc}.json

    # Back PV, remove ID, remove spec.azureDisk, add spec.csi
    cat << EOF > ${TMP}/csi-pv.spec
{
    "csi": {
        "driver": "disk.csi.azure.com",
        "volumeHandle": "${PV_DISK_URI}",
        "fsType": "${PV_FS_TYPE}"
    }
}
EOF

    kubectl get pv ${PV_NAME} -n ${NS} -o json | \
        jq 'del(.metadata.annotations, .metadata.creationTimestamp, .metadata.managedFields, .metadata.resourceVersion, .metadata.uid, .status, .spec.claimRef.resourceVersion, .spec.claimRef.uid, .spec.azureDisk)' | \
        jq ".spec += $(cat ${TMP}/csi-pv.spec)" \
        > ${TMP}/pv/${NS}-${PV_NAME}.json

    # Retain Azure disk on PV deletion
    kubectl patch pv ${PV_NAME} -p '{"spec":{"persistentVolumeReclaimPolicy":"Retain"}}'

    # Delete PVC+PV pair
    kubectl delete -n ${NS} pvc/${pvc}
    kubectl delete pv/${PV_NAME}
done

# Recreate resources
kubectl apply -f ${TMP}/pv
if [[ "${KIND,,}" == "deployment" ]]; then 
    kubectl apply -f ${TMP}/pvc; 
fi
kubectl apply -f ${TMP}/${KIND}