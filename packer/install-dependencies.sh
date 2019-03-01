#!/bin/bash
source /home/packer/provision_installs.sh
source /home/packer/provision_source.sh

echo "Starting build on " `date` > /var/log/azure/golden-image-install.complete
echo "Using kernel:" >> /var/log/azure/golden-image-install.complete
cat /proc/version | tee -a /var/log/azure/golden-image-install.complete

ETCD_VERSION="3.2.25"
ETCD_DOWNLOAD_URL="https://acs-mirror.azureedge.net/github-coreos"
installEtcd

installDeps

if [[ ${FEATURE_FLAGS} == *"docker-engine"* ]]; then
    DOCKER_ENGINE_REPO="https://apt.dockerproject.org/repo"
    installDockerEngine
    overrideDockerEngineStorageDriver
    installGPUDrivers
else
    MOBY_VERSION="3.0.4"
    installMoby
fi

installClearContainersRuntime

VNET_CNI_VERSIONS="1.0.16 1.0.17"
CNI_PLUGIN_VERSIONS="0.7.1"
CONTAINERD_VERSIONS="1.1.5 1.1.6 1.2.4"

for VNET_CNI_VERSION in $VNET_CNI_VERSIONS; do
    VNET_CNI_PLUGINS_URL="https://acs-mirror.azureedge.net/cni/azure-vnet-cni-linux-amd64-v${VNET_CNI_VERSION}.tgz"
    downloadAzureCNI
done

for CNI_PLUGIN_VERSION in $CNI_PLUGIN_VERSIONS; do
    CNI_PLUGINS_URL="https://acs-mirror.azureedge.net/cni/cni-plugins-amd64-v${CNI_PLUGIN_VERSION}.tgz"
    downloadCNI
done

CONTAINERD_DOWNLOAD_URL_BASE="https://storage.googleapis.com/cri-containerd-release/"
for CONTAINERD_VERSION in ${CONTAINERD_VERSIONS}; do
    downloadContainerd
done

installImg

DASHBOARD_VERSIONS="1.10.1"
for DASHBOARD_VERSION in ${DASHBOARD_VERSIONS}; do
    pullContainerImage "docker" "k8s.gcr.io/kubernetes-dashboard-amd64:v${DASHBOARD_VERSION}"
done

EXECHEALTHZ_VERSIONS="1.2"
for EXECHEALTHZ_VERSION in ${EXECHEALTHZ_VERSIONS}; do
    pullContainerImage "docker" "k8s.gcr.io/exechealthz-amd64:${EXECHEALTHZ_VERSION}"
done

ADDON_RESIZER_VERSIONS="1.8.4 1.8.1 1.7"
for ADDON_RESIZER_VERSION in ${ADDON_RESIZER_VERSIONS}; do
    pullContainerImage "docker" "k8s.gcr.io/addon-resizer:${ADDON_RESIZER_VERSION}"
done

HEAPSTER_VERSIONS="1.5.4 1.5.3 1.5.1"
for HEAPSTER_VERSION in ${HEAPSTER_VERSIONS}; do
    pullContainerImage "docker" "k8s.gcr.io/heapster-amd64:v${HEAPSTER_VERSION}"
done

METRICS_SERVER_VERSIONS="0.2.1"
for METRICS_SERVER_VERSION in ${METRICS_SERVER_VERSIONS}; do
    pullContainerImage "docker" "k8s.gcr.io/metrics-server-amd64:v${METRICS_SERVER_VERSION}"
done

KUBE_DNS_VERSIONS="1.15.0 1.14.13 1.14.5"
for KUBE_DNS_VERSION in ${KUBE_DNS_VERSIONS}; do
    pullContainerImage "docker" "k8s.gcr.io/k8s-dns-kube-dns-amd64:${KUBE_DNS_VERSION}"
done

KUBE_ADDON_MANAGER_VERSIONS="8.9 8.8 8.7 8.6"
for KUBE_ADDON_MANAGER_VERSION in ${KUBE_ADDON_MANAGER_VERSIONS}; do
    pullContainerImage "docker" "k8s.gcr.io/kube-addon-manager-amd64:v${KUBE_ADDON_MANAGER_VERSION}"
done

KUBE_DNS_MASQ_VERSIONS="1.15.0 1.14.10 1.14.8 1.14.5"
for KUBE_DNS_MASQ_VERSION in ${KUBE_DNS_MASQ_VERSIONS}; do
    pullContainerImage "docker" "k8s.gcr.io/k8s-dns-dnsmasq-nanny-amd64:${KUBE_DNS_MASQ_VERSION}"
done

PAUSE_VERSIONS="3.1"
for PAUSE_VERSION in ${PAUSE_VERSIONS}; do
    pullContainerImage "docker" "k8s.gcr.io/pause-amd64:${PAUSE_VERSION}"
done

TILLER_VERSIONS="2.8.1 2.11.0"
for TILLER_VERSION in ${TILLER_VERSIONS}; do
    pullContainerImage "docker" "gcr.io/kubernetes-helm/tiller:v${TILLER_VERSION}"
done

CLUSTER_AUTOSCALER_VERSIONS="1.13.1 1.12.2 1.3.7 1.3.4 1.3.3 1.2.2 1.1.2"
for CLUSTER_AUTOSCALER_VERSION in ${CLUSTER_AUTOSCALER_VERSIONS}; do
    pullContainerImage "docker" "k8s.gcr.io/cluster-autoscaler:v${CLUSTER_AUTOSCALER_VERSION}"
done

K8S_DNS_SIDECAR_VERSIONS="1.14.10 1.14.8 1.14.7"
for K8S_DNS_SIDECAR_VERSION in ${K8S_DNS_SIDECAR_VERSIONS}; do
    pullContainerImage "docker" "k8s.gcr.io/k8s-dns-sidecar-amd64:${K8S_DNS_SIDECAR_VERSION}"
done

CORE_DNS_VERSIONS="1.2.6 1.2.2"
for CORE_DNS_VERSION in ${CORE_DNS_VERSIONS}; do
    pullContainerImage "docker" "k8s.gcr.io/coredns:${CORE_DNS_VERSION}"
done

RESCHEDULER_VERSIONS="0.4.0 0.3.1"
for RESCHEDULER_VERSION in ${RESCHEDULER_VERSIONS}; do
    pullContainerImage "docker" "k8s.gcr.io/rescheduler:v${RESCHEDULER_VERSION}"
done

VIRTUAL_KUBELET_VERSIONS="latest"
for VIRTUAL_KUBELET_VERSION in ${VIRTUAL_KUBELET_VERSIONS}; do
    pullContainerImage "docker" "microsoft/virtual-kubelet:${VIRTUAL_KUBELET_VERSION}"
done

AZURE_CNI_NETWORKMONITOR_VERSIONS="0.0.5"
for AZURE_CNI_NETWORKMONITOR_VERSION in ${AZURE_CNI_NETWORKMONITOR_VERSIONS}; do
    pullContainerImage "docker" "containernetworking/networkmonitor:v${AZURE_CNI_NETWORKMONITOR_VERSION}"
done

NVIDIA_DEVICE_PLUGIN_VERSIONS="1.11 1.10"
for NVIDIA_DEVICE_PLUGIN_VERSION in ${NVIDIA_DEVICE_PLUGIN_VERSIONS}; do
    pullContainerImage "docker" "nvidia/k8s-device-plugin:${NVIDIA_DEVICE_PLUGIN_VERSION}"
done

TUNNELFRONT_VERSIONS="v1.9.2-v4.0.4"
for TUNNELFRONT_VERSION in ${TUNNELFRONT_VERSIONS}; do
    pullContainerImage "docker" "docker.io/deis/hcp-tunnel-front:${TUNNELFRONT_VERSION}"
done

KUBE_SVC_REDIRECT_VERSIONS="1.0.2"
for KUBE_SVC_REDIRECT_VERSION in ${KUBE_SVC_REDIRECT_VERSIONS}; do
    pullContainerImage "docker" "docker.io/deis/kube-svc-redirect:v${KUBE_SVC_REDIRECT_VERSION}"
done

KV_FLEXVOLUME_VERSIONS="0.0.7"
for KV_FLEXVOLUME_VERSION in ${KV_FLEXVOLUME_VERSIONS}; do
    pullContainerImage "docker" "mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v${KV_FLEXVOLUME_VERSION}"
done

BLOBFUSE_FLEXVOLUME_VERSIONS="1.0.8"
for BLOBFUSE_FLEXVOLUME_VERSION in ${BLOBFUSE_FLEXVOLUME_VERSIONS}; do
    pullContainerImage "docker" "mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:${BLOBFUSE_FLEXVOLUME_VERSION}"
done

IP_MASQ_AGENT_VERSIONS="2.0.0"
for IP_MASQ_AGENT_VERSION in ${IP_MASQ_AGENT_VERSIONS}; do
    pullContainerImage "docker" "gcr.io/google-containers/ip-masq-agent-amd64:v${IP_MASQ_AGENT_VERSION}"
done

NGINX_VERSIONS="1.13.12-alpine"
for NGINX_VERSION in ${NGINX_VERSIONS}; do
    pullContainerImage "docker" "nginx:${NGINX_VERSION}"
done

KMS_PLUGIN_VERSIONS="0.0.9"
for KMS_PLUGIN_VERSION in ${KMS_PLUGIN_VERSIONS}; do
    pullContainerImage "docker" "mcr.microsoft.com/k8s/kms/keyvault:v${KMS_PLUGIN_VERSION}"
done

FLANNEL_VERSIONS="0.8.0 0.10.0"
for FLANNEL_VERSION in ${FLANNEL_VERSIONS}; do
    pullContainerImage "docker" "quay.io/coreos/flannel:v${FLANNEL_VERSION}"
done

pullContainerImage "docker" "busybox"

# TODO: fetch supported k8s versions from an aks-engine command instead of hardcoding them here
K8S_VERSIONS="1.9.10 1.9.11 1.10.12 1.10.13 1.11.7 1.11.8 1.12.5 1.12.6 1.13.3 1.13.4"

for KUBERNETES_VERSION in ${K8S_VERSIONS}; do
    HYPERKUBE_URL="k8s.gcr.io/hyperkube-amd64:v${KUBERNETES_VERSION}"
    extractHyperkube "docker"
    pullContainerImage "docker" "k8s.gcr.io/cloud-controller-manager-amd64:v${KUBERNETES_VERSION}"
done

df -h

echo "Install completed successfully on " `date` >> /var/log/azure/golden-image-install.complete
echo "VSTS Build NUMBER: ${BUILD_NUMBER}" >> /var/log/azure/golden-image-install.complete
echo "VSTS Build ID: ${BUILD_ID}" >> /var/log/azure/golden-image-install.complete
echo "Commit: ${COMMIT}" >> /var/log/azure/golden-image-install.complete
echo "Feature flags: ${FEATURE_FLAGS}" >> /var/log/azure/golden-image-install.complete
