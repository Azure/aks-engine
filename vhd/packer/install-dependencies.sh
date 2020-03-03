#!/bin/bash
source /home/packer/provision_installs.sh
source /home/packer/provision_source.sh
source /home/packer/packer_source.sh

VHD_LOGS_FILEPATH=/opt/azure/vhd-install.complete

echo "Starting build on " $(date) > ${VHD_LOGS_FILEPATH}

copyPackerFiles

echo ""
echo "Components downloaded in this VHD build (some of the below components might get deleted during cluster provisioning if they are not needed):" >> ${VHD_LOGS_FILEPATH}

AUDITD_ENABLED=true
installDeps
cat << EOF >> ${VHD_LOGS_FILEPATH}
  - apache2-utils
  - apt-transport-https
  - auditd
  - blobfuse
  - ca-certificates
  - ceph-common
  - cgroup-lite
  - cifs-utils
  - conntrack
  - cracklib-runtime
  - dbus
  - ebtables
  - ethtool
  - fuse
  - git
  - glusterfs-client
  - init-system-helpers
  - iproute2
  - ipset
  - iptables
  - jq
  - libpam-pwquality
  - libpwquality-tools
  - mount
  - nfs-common
  - pigz socat
  - traceroute
  - util-linux
  - xz-utils
  - zip
EOF

if [[ ${UBUNTU_RELEASE} == "18.04" ]]; then
  overrideNetworkConfig
fi

apmz_version="v0.5.0"
ensureAPMZ "${apmz_version}"
echo "  - apmz $apmz_version" >> ${VHD_LOGS_FILEPATH}

installBpftrace
echo "  - bpftrace" >> ${VHD_LOGS_FILEPATH}

MOBY_VERSION="3.0.10"
installMoby
echo "  - moby v${MOBY_VERSION}" >> ${VHD_LOGS_FILEPATH}
installGPUDrivers
echo "  - nvidia-docker2 nvidia-container-runtime" >> ${VHD_LOGS_FILEPATH}

ETCD_VERSION="3.3.18"
ETCD_DOWNLOAD_URL="mcr.microsoft.com/oss/etcd-io/"
installEtcd "docker"
echo "  - etcd v${ETCD_VERSION}" >> ${VHD_LOGS_FILEPATH}

installBcc
cat << EOF >> ${VHD_LOGS_FILEPATH}
  - bcc-tools
  - libbcc-examples
EOF

VNET_CNI_VERSIONS="
1.0.33
1.0.30
1.0.29
"
for VNET_CNI_VERSION in $VNET_CNI_VERSIONS; do
    VNET_CNI_PLUGINS_URL="https://kubernetesartifacts.azureedge.net/azure-cni/v${VNET_CNI_VERSION}/binaries/azure-vnet-cni-linux-amd64-v${VNET_CNI_VERSION}.tgz"
    downloadAzureCNI
    echo "  - Azure CNI version ${VNET_CNI_VERSION}" >> ${VHD_LOGS_FILEPATH}
done

CNI_PLUGIN_VERSIONS="
0.7.6
0.7.5
0.7.1
"
for CNI_PLUGIN_VERSION in $CNI_PLUGIN_VERSIONS; do
    CNI_PLUGINS_URL="https://kubernetesartifacts.azureedge.net/cni-plugins/v${CNI_PLUGIN_VERSION}/binaries/cni-plugins-amd64-v${CNI_PLUGIN_VERSION}.tgz"
    downloadCNI
    echo "  - CNI plugin version ${CNI_PLUGIN_VERSION}" >> ${VHD_LOGS_FILEPATH}
done

installImg
echo "  - img" >> ${VHD_LOGS_FILEPATH}

echo "Docker images pre-pulled:" >> ${VHD_LOGS_FILEPATH}

DASHBOARD_VERSIONS="1.10.1"
for DASHBOARD_VERSION in ${DASHBOARD_VERSIONS}; do
    CONTAINER_IMAGE="k8s.gcr.io/kubernetes-dashboard-amd64:v${DASHBOARD_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
    CONTAINER_IMAGE="mcr.microsoft.com/oss/kubernetes/kubernetes-dashboard:v${DASHBOARD_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

ADDON_RESIZER_VERSIONS="
1.8.7
1.8.4
"
for ADDON_RESIZER_VERSION in ${ADDON_RESIZER_VERSIONS}; do
    CONTAINER_IMAGE="k8s.gcr.io/addon-resizer:${ADDON_RESIZER_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
    CONTAINER_IMAGE="mcr.microsoft.com/oss/kubernetes/autoscaler/addon-resizer:${ADDON_RESIZER_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

METRICS_SERVER_VERSIONS="
0.3.5
0.3.4
0.2.1
"
for METRICS_SERVER_VERSION in ${METRICS_SERVER_VERSIONS}; do
    CONTAINER_IMAGE="k8s.gcr.io/metrics-server-amd64:v${METRICS_SERVER_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
    CONTAINER_IMAGE="mcr.microsoft.com/oss/kubernetes/metrics-server:v${METRICS_SERVER_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

KUBE_DNS_VERSIONS="
1.15.4
"
for KUBE_DNS_VERSION in ${KUBE_DNS_VERSIONS}; do
    CONTAINER_IMAGE="k8s.gcr.io/k8s-dns-kube-dns-amd64:${KUBE_DNS_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
    CONTAINER_IMAGE="mcr.microsoft.com/oss/kubernetes/k8s-dns-kube-dns:${KUBE_DNS_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

KUBE_ADDON_MANAGER_VERSIONS="
9.0.2
8.9.1
"
for KUBE_ADDON_MANAGER_VERSION in ${KUBE_ADDON_MANAGER_VERSIONS}; do
    CONTAINER_IMAGE="k8s.gcr.io/kube-addon-manager-amd64:v${KUBE_ADDON_MANAGER_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
    CONTAINER_IMAGE="mcr.microsoft.com/oss/kubernetes/kube-addon-manager:v${KUBE_ADDON_MANAGER_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

KUBE_DNS_MASQ_VERSIONS="
1.15.4
"
for KUBE_DNS_MASQ_VERSION in ${KUBE_DNS_MASQ_VERSIONS}; do
    CONTAINER_IMAGE="k8s.gcr.io/k8s-dns-dnsmasq-nanny-amd64:${KUBE_DNS_MASQ_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
    CONTAINER_IMAGE="mcr.microsoft.com/oss/kubernetes/k8s-dns-dnsmasq-nanny:${KUBE_DNS_MASQ_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

MCR_PAUSE_VERSIONS="1.2.0"
for PAUSE_VERSION in ${MCR_PAUSE_VERSIONS}; do
    # Pull the arch independent MCR pause image which is built for Linux and Windows
    CONTAINER_IMAGE="mcr.microsoft.com/k8s/core/pause:${PAUSE_VERSION}"
    pullContainerImage "docker" "${CONTAINER_IMAGE}"
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

GCR_PAUSE_VERSIONS="3.1"
for PAUSE_VERSION in ${GCR_PAUSE_VERSIONS}; do
    CONTAINER_IMAGE="k8s.gcr.io/pause-amd64:${PAUSE_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

TILLER_VERSIONS="
2.13.1
"
for TILLER_VERSION in ${TILLER_VERSIONS}; do
    CONTAINER_IMAGE="gcr.io/kubernetes-helm/tiller:v${TILLER_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
    CONTAINER_IMAGE="mcr.microsoft.com/oss/kubernetes/tiller:v${TILLER_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

CLUSTER_AUTOSCALER_VERSIONS="
1.17.1
1.16.4
1.15.5
1.14.7
1.13.9
"
for CLUSTER_AUTOSCALER_VERSION in ${CLUSTER_AUTOSCALER_VERSIONS}; do
    CONTAINER_IMAGE="k8s.gcr.io/cluster-autoscaler:v${CLUSTER_AUTOSCALER_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

K8S_DNS_SIDECAR_VERSIONS="
1.14.10
"
for K8S_DNS_SIDECAR_VERSION in ${K8S_DNS_SIDECAR_VERSIONS}; do
    CONTAINER_IMAGE="k8s.gcr.io/k8s-dns-sidecar-amd64:${K8S_DNS_SIDECAR_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
    CONTAINER_IMAGE="mcr.microsoft.com/oss/kubernetes/k8s-dns-sidecar:${K8S_DNS_SIDECAR_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

CORE_DNS_VERSIONS="
1.6.7
"
for CORE_DNS_VERSION in ${CORE_DNS_VERSIONS}; do
    CONTAINER_IMAGE="k8s.gcr.io/coredns:${CORE_DNS_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
    CONTAINER_IMAGE="mcr.microsoft.com/oss/kubernetes/coredns:${CORE_DNS_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

RESCHEDULER_VERSIONS="
0.4.0
"
for RESCHEDULER_VERSION in ${RESCHEDULER_VERSIONS}; do
    CONTAINER_IMAGE="k8s.gcr.io/rescheduler:v${RESCHEDULER_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
    CONTAINER_IMAGE="mcr.microsoft.com/oss/kubernetes/rescheduler:v${RESCHEDULER_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

VIRTUAL_KUBELET_VERSIONS="latest"
for VIRTUAL_KUBELET_VERSION in ${VIRTUAL_KUBELET_VERSIONS}; do
    CONTAINER_IMAGE="microsoft/virtual-kubelet:${VIRTUAL_KUBELET_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

AZURE_CNIIMAGEBASE="mcr.microsoft.com/containernetworking"
AZURE_CNI_NETWORKMONITOR_VERSIONS="
0.0.6
"
for AZURE_CNI_NETWORKMONITOR_VERSION in ${AZURE_CNI_NETWORKMONITOR_VERSIONS}; do
    CONTAINER_IMAGE="${AZURE_CNIIMAGEBASE}/networkmonitor:v${AZURE_CNI_NETWORKMONITOR_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

AZURE_NPM_VERSIONS="
1.0.32
"
for AZURE_NPM_VERSION in ${AZURE_NPM_VERSIONS}; do
    CONTAINER_IMAGE="${AZURE_CNIIMAGEBASE}/azure-npm:v${AZURE_NPM_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

NVIDIA_DEVICE_PLUGIN_VERSIONS="
1.11
"
for NVIDIA_DEVICE_PLUGIN_VERSION in ${NVIDIA_DEVICE_PLUGIN_VERSIONS}; do
    CONTAINER_IMAGE="nvidia/k8s-device-plugin:${NVIDIA_DEVICE_PLUGIN_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

KV_FLEXVOLUME_VERSIONS="0.0.16"
for KV_FLEXVOLUME_VERSION in ${KV_FLEXVOLUME_VERSIONS}; do
    CONTAINER_IMAGE="mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v${KV_FLEXVOLUME_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

BLOBFUSE_FLEXVOLUME_VERSIONS="1.0.8"
for BLOBFUSE_FLEXVOLUME_VERSION in ${BLOBFUSE_FLEXVOLUME_VERSIONS}; do
    CONTAINER_IMAGE="mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:${BLOBFUSE_FLEXVOLUME_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

IP_MASQ_AGENT_VERSIONS="
2.5.0
"
for IP_MASQ_AGENT_VERSION in ${IP_MASQ_AGENT_VERSIONS}; do
    CONTAINER_IMAGE="k8s.gcr.io/ip-masq-agent-amd64:v${IP_MASQ_AGENT_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
    CONTAINER_IMAGE="mcr.microsoft.com/oss/kubernetes/ip-masq-agent:v${IP_MASQ_AGENT_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

KMS_PLUGIN_VERSIONS="0.0.9"
for KMS_PLUGIN_VERSION in ${KMS_PLUGIN_VERSIONS}; do
    CONTAINER_IMAGE="mcr.microsoft.com/k8s/kms/keyvault:v${KMS_PLUGIN_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

FLANNEL_VERSIONS="
0.10.0
0.8.0
"
for FLANNEL_VERSION in ${FLANNEL_VERSIONS}; do
    CONTAINER_IMAGE="quay.io/coreos/flannel:v${FLANNEL_VERSION}-amd64"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

pullContainerImage "docker" "busybox"
echo "  - busybox" >> ${VHD_LOGS_FILEPATH}

K8S_VERSIONS="
1.18.0-beta.1
1.17.3
1.17.2
1.16.7
1.16.6
1.15.10
1.15.9
1.15.9-azs
1.14.8
1.14.7
1.14.7-azs
1.13.12
1.13.11
"
for KUBERNETES_VERSION in ${K8S_VERSIONS}; do
  if (( $(echo ${KUBERNETES_VERSION} | cut -d"." -f2) < 17 )); then
    if [[ $KUBERNETES_VERSION == *"azs"* ]]; then
      HYPERKUBE_URL="mcr.microsoft.com/oss/kubernetes/hyperkube:v${KUBERNETES_VERSION}"
    else
      HYPERKUBE_URL="k8s.gcr.io/hyperkube-amd64:v${KUBERNETES_VERSION}"
    fi
    extractHyperkube "docker"
    echo "  - ${HYPERKUBE_URL}" >> ${VHD_LOGS_FILEPATH}
  else
    for component in kube-apiserver kube-controller-manager kube-proxy kube-scheduler; do
      CONTAINER_IMAGE="k8s.gcr.io/${component}:v${KUBERNETES_VERSION}"
      pullContainerImage "docker" ${CONTAINER_IMAGE}
      echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
    done
    KUBE_BINARY_URL="https://dl.k8s.io/v${KUBERNETES_VERSION}/kubernetes-node-linux-amd64.tar.gz"
    extractKubeBinaries
  fi
  if (( $(echo ${KUBERNETES_VERSION} | cut -d"." -f2) < 16 )) && [[ $KUBERNETES_VERSION != *"azs"* ]]; then
    CONTAINER_IMAGE="k8s.gcr.io/cloud-controller-manager-amd64:v${KUBERNETES_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
  fi
done

# Starting with 1.16 we pull a common cloud-controller-manager image
CLOUD_MANAGER_VERSIONS="
0.4.1
"
for CLOUD_MANAGER_VERSION in ${CLOUD_MANAGER_VERSIONS}; do
  for COMPONENT in azure-cloud-controller-manager azure-cloud-node-manager; do
    CONTAINER_IMAGE="mcr.microsoft.com/oss/kubernetes/${COMPONENT}:v${CLOUD_MANAGER_VERSION}"
    pullContainerImage "docker" ${CONTAINER_IMAGE}
    echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
  done
done

AZUREDISK_CSI_VERSIONS="
0.5.0
"
for AZUREDISK_CSI_VERSION in ${AZUREDISK_CSI_VERSIONS}; do
  CONTAINER_IMAGE="mcr.microsoft.com/k8s/csi/azuredisk-csi:v${AZUREDISK_CSI_VERSION}"
  pullContainerImage "docker" ${CONTAINER_IMAGE}
  echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

AZUREFILE_CSI_VERSIONS="
0.3.0
"
for AZUREFILE_CSI_VERSION in ${AZUREFILE_CSI_VERSIONS}; do
  CONTAINER_IMAGE="mcr.microsoft.com/k8s/csi/azurefile-csi:v${AZUREFILE_CSI_VERSION}"
  pullContainerImage "docker" ${CONTAINER_IMAGE}
  echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

CSI_ATTACHER_VERSIONS="
1.2.0
"
for CSI_ATTACHER_VERSION in ${CSI_ATTACHER_VERSIONS}; do
  CONTAINER_IMAGE="mcr.microsoft.com/oss/kubernetes-csi/csi-attacher:v${CSI_ATTACHER_VERSION}"
  pullContainerImage "docker" ${CONTAINER_IMAGE}
  echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

CSI_CLUSTER_DRIVER_REGISTRAR_VERSIONS="
1.0.1
"
for CSI_CLUSTER_DRIVER_REGISTRAR_VERSION in ${CSI_CLUSTER_DRIVER_REGISTRAR_VERSIONS}; do
  CONTAINER_IMAGE="mcr.microsoft.com/oss/kubernetes-csi/csi-cluster-driver-registrar:v${CSI_CLUSTER_DRIVER_REGISTRAR_VERSION}"
  pullContainerImage "docker" ${CONTAINER_IMAGE}
  echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

CSI_NODE_DRIVER_REGISTRAR_VERSIONS="
1.1.0
"
for CSI_NODE_DRIVER_REGISTRAR_VERSION in ${CSI_NODE_DRIVER_REGISTRAR_VERSIONS}; do
  CONTAINER_IMAGE="mcr.microsoft.com/oss/kubernetes-csi/csi-node-driver-registrar:v${CSI_NODE_DRIVER_REGISTRAR_VERSION}"
  pullContainerImage "docker" ${CONTAINER_IMAGE}
  echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

CSI_PROVISIONER_VERSIONS="
1.4.0
"
for CSI_PROVISIONER_VERSION in ${CSI_PROVISIONER_VERSIONS}; do
  CONTAINER_IMAGE="mcr.microsoft.com/oss/kubernetes-csi/csi-provisioner:v${CSI_PROVISIONER_VERSION}"
  pullContainerImage "docker" ${CONTAINER_IMAGE}
  echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

LIVENESSPROBE_VERSIONS="
1.1.0
"
for LIVENESSPROBE_VERSION in ${LIVENESSPROBE_VERSIONS}; do
  CONTAINER_IMAGE="mcr.microsoft.com/oss/kubernetes-csi/livenessprobe:v${LIVENESSPROBE_VERSION}"
  pullContainerImage "docker" ${CONTAINER_IMAGE}
  echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

CSI_RESIZER_VERSIONS="
0.3.0
"
for CSI_RESIZER_VERSION in ${CSI_RESIZER_VERSIONS}; do
  CONTAINER_IMAGE="mcr.microsoft.com/oss/kubernetes-csi/csi-resizer:v${CSI_RESIZER_VERSION}"
  pullContainerImage "docker" ${CONTAINER_IMAGE}
  echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

CSI_SNAPSHOTTER_VERSIONS="
1.1.0
"
for CSI_SNAPSHOTTER_VERSION in ${CSI_SNAPSHOTTER_VERSIONS}; do
  CONTAINER_IMAGE="mcr.microsoft.com/oss/kubernetes-csi/csi-snapshotter:v${CSI_SNAPSHOTTER_VERSION}"
  pullContainerImage "docker" ${CONTAINER_IMAGE}
  echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

NODE_PROBLEM_DETECTOR_VERSIONS="
0.8.1
"
for NODE_PROBLEM_DETECTOR_VERSION in ${NODE_PROBLEM_DETECTOR_VERSIONS}; do
  CONTAINER_IMAGE="k8s.gcr.io/node-problem-detector:v${NODE_PROBLEM_DETECTOR_VERSION}"
  pullContainerImage "docker" ${CONTAINER_IMAGE}
  echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}
done

# This is to accommodate air-gapped environments, e.g., Azure Stack
CONTAINER_IMAGE="registry:2.7.1"
pullContainerImage "docker" ${CONTAINER_IMAGE}
echo "  - ${CONTAINER_IMAGE}" >> ${VHD_LOGS_FILEPATH}

df -h

# warn at 75% space taken
[ -s $(df -P | grep '/dev/sda1' | awk '0+$5 >= 75 {print}') ] || echo "WARNING: 75% of /dev/sda1 is used" >> ${VHD_LOGS_FILEPATH}
# error at 90% space taken
[ -s $(df -P | grep '/dev/sda1' | awk '0+$5 >= 90 {print}') ] || exit 1

echo "Using kernel:" >> ${VHD_LOGS_FILEPATH}
tee -a ${VHD_LOGS_FILEPATH} < /proc/version
{
  echo "Install completed successfully on " $(date)
  echo "VSTS Build NUMBER: ${BUILD_NUMBER}"
  echo "VSTS Build ID: ${BUILD_ID}"
  echo "Commit: ${COMMIT}"
  echo "Feature flags: ${FEATURE_FLAGS}"
} >> ${VHD_LOGS_FILEPATH}
