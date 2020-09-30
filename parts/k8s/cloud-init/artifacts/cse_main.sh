#!/bin/bash
ERR_FILE_WATCH_TIMEOUT=6 {{/* Timeout waiting for a file */}}

set -x
if [ -f /opt/azure/containers/provision.complete ]; then
  exit 0
fi

echo $(date),$(hostname), startcustomscript >>/opt/m

for i in $(seq 1 3600); do
  if [ -s {{GetCSEHelpersScriptFilepath}} ]; then
    grep -Fq '#HELPERSEOF' {{GetCSEHelpersScriptFilepath}} && break
  fi
  if [ $i -eq 3600 ]; then
    exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  else
    sleep 1
  fi
done
sed -i "/#HELPERSEOF/d" {{GetCSEHelpersScriptFilepath}}
source {{GetCSEHelpersScriptFilepath}}
configure_prerequisites

wait_for_file 3600 1 {{GetCSEInstallScriptFilepath}} || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
source {{GetCSEInstallScriptFilepath}}

ensureAPMZ "v0.5.1"
{{- if HasTelemetryEnabled }}
eval "$(apmz bash -n "cse" -t "{{GetLinuxDefaultTelemetryTags}}" --api-keys "{{GetApplicationInsightsTelemetryKeys}}")"
{{else}}
eval "$(apmz bash -d)"
{{end}}

wait_for_file 3600 1 {{GetCSEConfigScriptFilepath}} || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
source {{GetCSEConfigScriptFilepath}}

{{- if IsCustomCloudProfile}}
wait_for_file 3600 1 {{GetCustomCloudConfigCSEScriptFilepath}} || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
source {{GetCustomCloudConfigCSEScriptFilepath }}
{{end}}

set +x
ETCD_PEER_CERT=$(echo ${ETCD_PEER_CERTIFICATES} | cut -d'[' -f 2 | cut -d']' -f 1 | cut -d',' -f $((NODE_INDEX + 1)))
ETCD_PEER_KEY=$(echo ${ETCD_PEER_PRIVATE_KEYS} | cut -d'[' -f 2 | cut -d']' -f 1 | cut -d',' -f $((NODE_INDEX + 1)))
set -x

time_metric "ConfigureAdminUser" configureAdminUser

{{- if HasVHDDistroNodes}}
  {{- if not NeedsContainerd}}
time_metric "CleanupContainerd" cleanUpContainerd
  {{end}}
  {{- if HasNSeriesSKU}}
if [[ ${GPU_NODE} != "true" ]]; then
  time_metric "CleanupGPUDrivers" cleanUpGPUDrivers
fi
  {{else}}
time_metric "CleanupGPUDrivers" cleanUpGPUDrivers
  {{end}}
{{end}}

{{- if HasVHDDistroNodes}}
VHD_LOGS_FILEPATH=/opt/azure/vhd-install.complete
if [ -f $VHD_LOGS_FILEPATH ]; then
  time_metric "CleanUpContainerImages" cleanUpContainerImages
  FULL_INSTALL_REQUIRED=false
else
  if [[ ${IS_VHD} == true ]]; then
    exit {{GetCSEErrorCode "ERR_VHD_FILE_NOT_FOUND"}}
  fi
  FULL_INSTALL_REQUIRED=true
fi
{{else}}
FULL_INSTALL_REQUIRED=true
{{end}}

{{- if not IsVHDDistroForAllNodes}}
if [[ $OS == $UBUNTU_OS_NAME || $OS == $DEBIAN_OS_NAME ]] && [ "$FULL_INSTALL_REQUIRED" = "true" ]; then
  time_metric "InstallDeps" installDeps
  if [[ ${UBUNTU_RELEASE} == "18.04" ]]; then
    overrideNetworkConfig
  fi
  if [[ $OS == $UBUNTU_OS_NAME ]]; then
    time_metric "InstallBcc" installBcc
  fi
  {{- if not IsDockerContainerRuntime}}
  time_metric "InstallImg" installImg
  {{end}}
fi
{{end}}

if [[ ${UBUNTU_RELEASE} == "18.04" ]]; then
  if apt list --installed | grep 'ntp'; then
    time_metric "EnsureNTP" ensureNTP
  fi
fi

if [[ $OS == $UBUNTU_OS_NAME ]]; then
  time_metric "EnsureAuditD" ensureAuditD
fi

{{- if not IsVHDDistroForAllNodes}}
if [[ $FULL_INSTALL_REQUIRED == "true" ]]; then
  time_metric "InstallBpftrace" installBpftrace
fi
{{end}}

if [[ $OS != $FLATCAR_OS_NAME ]]; then
{{- if NeedsContainerd}}
time_metric "InstallContainerd" installContainerd
{{else}}
time_metric "installMoby" installMoby
{{end}}
{{- if HasLinuxContainerdURL}}
  LINUX_CONTAINERD_URL={{GetLinuxContainerdURL}}
  if [[ -n "${LINUX_CONTAINERD_URL:-}" ]]; then
    DEB="${LINUX_CONTAINERD_URL##*/}"
    retrycmd_no_stats 120 5 25 curl -fsSL ${LINUX_CONTAINERD_URL} >/tmp/${DEB} || exit {{GetCSEErrorCode "ERR_DEB_DOWNLOAD_TIMEOUT"}}
    retrycmd 20 30 120 dpkg -i /tmp/${DEB} || exit {{GetCSEErrorCode "ERR_DEB_PKG_ADD_FAIL"}}
  fi
{{end}}
fi

if [[ -n ${MASTER_NODE} ]] && [[ -z ${COSMOS_URI} ]]; then
  {{- if IsDockerContainerRuntime}}
  cli_tool="docker"
  {{else}}
  cli_tool="img"
  {{end}}
  time_metric "InstallEtcd" installEtcd $cli_tool
fi

{{/* this will capture the amount of time to install of the network plugin during cse */}}
time_metric "InstallNetworkPlugin" installNetworkPlugin

{{- if HasNSeriesSKU}}
if [[ ${GPU_NODE} == true ]]; then
  if $FULL_INSTALL_REQUIRED; then
    time_metric "DownloadGPUDrivers" downloadGPUDrivers
  fi
  time_metric "EnsureGPUDrivers" ensureGPUDrivers
fi
{{end}}

{{- if and IsDockerContainerRuntime HasPrivateAzureRegistryServer}}
docker login -u $SERVICE_PRINCIPAL_CLIENT_ID -p $SERVICE_PRINCIPAL_CLIENT_SECRET {{GetPrivateAzureRegistryServer}}
{{end}}

time_metric "InstallKubeletAndKubectl" installKubeletAndKubectl

if [[ $OS != $FLATCAR_OS_NAME ]]; then
    time_metric "EnsureRPC" ensureRPC
    time_metric "EnsureCron" ensureCron
fi

time_metric "CreateKubeManifestDir" createKubeManifestDir

{{- if HasDCSeriesSKU}}
if [[ ${SGX_NODE} == true && ! -e "/dev/sgx" ]]; then
  time_metric "InstallSGXDrivers" installSGXDrivers
fi
{{end}}

{{/* create etcd user if we are configured for etcd */}}
if [[ -n ${MASTER_NODE} ]] && [[ -z ${COSMOS_URI} ]]; then
  time_metric "ConfigureEtcdUser" configureEtcdUser
fi

if [[ -n ${MASTER_NODE} ]]; then
  {{/* this step configures all certs */}}
  {{/* both configs etcd/cosmos */}}
  time_metric "ConfigureSecrets" configureSecrets
fi

{{/* configure etcd if we are configured for etcd */}}
if [[ -n ${MASTER_NODE} ]] && [[ -z ${COSMOS_URI} ]]; then
  time_metric "ConfigureEtcd" configureEtcd
else
  time_metric "RemoveEtcd" removeEtcd
fi

{{- if HasCustomSearchDomain}}
wait_for_file 3600 1 {{GetCustomSearchDomainsCSEScriptFilepath}} || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
{{GetCustomSearchDomainsCSEScriptFilepath}} >/opt/azure/containers/setup-custom-search-domain.log 2>&1 || exit {{GetCSEErrorCode "ERR_CUSTOM_SEARCH_DOMAINS_FAIL"}}
{{end}}

{{- if IsDockerContainerRuntime}}
time_metric "EnsureDocker" ensureDocker
{{end}}

time_metric "ConfigureK8s" configureK8s

{{- if IsCustomCloudProfile}}
time_metric "ConfigureK8sCustomCloud" configureK8sCustomCloud
{{- if and IsAzureStackCloud IsAzureCNI}}
time_metric "ConfigureAzureStackInterfaces" configureAzureStackInterfaces
{{end}}
{{end}}

time_metric "ConfigureCNI" configureCNI

if [[ -n ${MASTER_NODE} ]]; then
  time_metric "ConfigAddons" configAddons
  time_metric "WriteKubeConfig" writeKubeConfig
fi

{{- if NeedsContainerd}}
time_metric "EnsureContainerd" ensureContainerd
{{end}}

{{- if and IsHostedMaster EnableHostsConfigAgent}}
time_metric "ConfigPrivateClusterHosts" configPrivateClusterHosts
{{end}}

{{/* configure and enable dhcpv6 for ipv6 features */}}
{{- if IsIPv6Enabled}}
time_metric "EnsureDHCPv6" ensureDHCPv6
{{end}}

time_metric "EnsureKubelet" ensureKubelet
{{if IsAzurePolicyAddonEnabled}}
if [[ -n ${MASTER_NODE} ]]; then
  time_metric "EnsureLabelExclusionForAzurePolicyAddon" ensureLabelExclusionForAzurePolicyAddon
fi
{{end}}
time_metric "EnsureJournal" ensureJournal

if [[ -n ${MASTER_NODE} ]]; then
  if version_gte ${KUBERNETES_VERSION} 1.16; then
    time_metric "EnsureLabelNodes" ensureLabelNodes
  fi
{{- if IsAADPodIdentityAddonEnabled}}
  time_metric "EnsureTaints" ensureTaints
{{end}}
  if [[ -z ${COSMOS_URI} ]]; then
    if ! { [ "$FULL_INSTALL_REQUIRED" = "true" ] && [ ${UBUNTU_RELEASE} == "18.04" ]; }; then
      time_metric "EnsureEtcd" ensureEtcd
    fi
  fi
  time_metric "EnsureK8sControlPlane" ensureK8sControlPlane
  if [ -f /var/run/reboot-required ]; then
    time_metric "ReplaceAddonsInit" replaceAddonsInit
  else
    time_metric "EnsureAddons" ensureAddons
  fi
  {{- if HasClusterInitComponent}}
  if [[ $NODE_INDEX == 0 ]]; then
    retrycmd 120 5 30 $KUBECTL apply -f /opt/azure/containers/cluster-init.yaml || exit {{GetCSEErrorCode "ERR_CLUSTER_INIT_FAIL"}}
  fi
  {{end}}
fi

{{- if not IsVHDDistroForAllNodes}}
if $FULL_INSTALL_REQUIRED; then
  if [[ $OS == $UBUNTU_OS_NAME ]]; then
    {{/* mitigation for bug https://bugs.launchpad.net/ubuntu/+source/linux/+bug/1676635 */}}
    echo 2dd1ce17-079e-403c-b352-a1921ee207ee >/sys/bus/vmbus/drivers/hv_util/unbind
    sed -i "13i\echo 2dd1ce17-079e-403c-b352-a1921ee207ee > /sys/bus/vmbus/drivers/hv_util/unbind\n" /etc/rc.local
  fi
fi
{{end}}

{{- /* re-enable unattended upgrades */}}
rm -f /etc/apt/apt.conf.d/99periodic

{{- if not IsAzureStackCloud}}
if [[ $OS == $UBUNTU_OS_NAME ]]; then
  time_metric "PurgeApt" apt_get_purge apache2-utils &
fi
{{end}}

VALIDATION_ERR=0

{{- if IsHostedMaster }}
API_SERVER_DNS_RETRIES=20
if [[ $API_SERVER_NAME == *.privatelink.* ]]; then
  API_SERVER_DNS_RETRIES=200
fi
RES=$(retrycmd ${API_SERVER_DNS_RETRIES} 1 3 nslookup ${API_SERVER_NAME})
STS=$?
if [[ $STS != 0 ]]; then
    if [[ $RES == *"168.63.129.16"*  ]]; then
        VALIDATION_ERR={{GetCSEErrorCode "ERR_K8S_API_SERVER_AZURE_DNS_LOOKUP_FAIL"}}
    else
        VALIDATION_ERR={{GetCSEErrorCode "ERR_K8S_API_SERVER_DNS_LOOKUP_FAIL"}}
    fi
else
    API_SERVER_CONN_RETRIES=50
    if [[ $API_SERVER_NAME == *.privatelink.* ]]; then
        API_SERVER_CONN_RETRIES=100
    fi
    retrycmd ${API_SERVER_CONN_RETRIES} 1 3 nc -vz ${API_SERVER_NAME} 443 &&
    retrycmd ${API_SERVER_CONN_RETRIES} 1 3 nc -vz ${API_SERVER_NAME} 9000 &&
    retrycmd ${API_SERVER_CONN_RETRIES} 1 3 nc -uvz ${API_SERVER_NAME} 1194 ||
    VALIDATION_ERR={{GetCSEErrorCode "ERR_K8S_API_SERVER_CONN_FAIL"}}
fi

{{end}}

if [ -f /var/run/reboot-required ]; then
  trace_info "RebootRequired" "reboot=true"
  /bin/bash -c "shutdown -r 1 &"
  if [[ $OS == $UBUNTU_OS_NAME ]]; then
    aptmarkWALinuxAgent unhold &
  fi
else
  if [[ $OS == $UBUNTU_OS_NAME ]]; then
    /usr/lib/apt/apt.systemd.daily &
    aptmarkWALinuxAgent unhold &
  fi
fi

echo "CSE finished successfully"
echo $(date),$(hostname), endcustomscript >>/opt/m
mkdir -p /opt/azure/containers && touch /opt/azure/containers/provision.complete
ps auxfww >/opt/azure/provision-ps.log &

exit $VALIDATION_ERR

#EOF
