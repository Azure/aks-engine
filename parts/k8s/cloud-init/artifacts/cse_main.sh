#!/bin/bash
ERR_FILE_WATCH_TIMEOUT=6 {{/* Timeout waiting for a file */}}

set -x
echo $(date),$(hostname), startcustomscript>>/opt/m

for i in $(seq 1 3600); do
    if [ -s {{GetCSEHelpersScriptFilepath}} ]; then
        grep -Fq '#HELPERSEOF' {{GetCSEHelpersScriptFilepath}} && break
    fi
    if [ $i -eq 3600 ]; then
        exit $ERR_FILE_WATCH_TIMEOUT
    else
        sleep 1
    fi
done
sed -i "/#HELPERSEOF/d" {{GetCSEHelpersScriptFilepath}}
source {{GetCSEHelpersScriptFilepath}}
configure_prerequisites

wait_for_file 3600 1 {{GetCSEInstallScriptFilepath}} || exit $ERR_FILE_WATCH_TIMEOUT
source {{GetCSEInstallScriptFilepath}}

ensureAPMZ "v0.5.1"
{{- if HasTelemetryEnabled }}
eval "$(apmz bash -n "cse" -t "{{GetLinuxDefaultTelemetryTags}}" --api-keys "{{GetApplicationInsightsTelemetryKeys}}")"
{{else}}
eval "$(apmz bash -d)"
{{end}}

wait_for_file 3600 1 {{GetCSEConfigScriptFilepath}} || exit $ERR_FILE_WATCH_TIMEOUT
source {{GetCSEConfigScriptFilepath}}

{{- if IsAzureStackCloud}}
wait_for_file 3600 1 {{GetCustomCloudConfigCSEScriptFilepath}} || exit $ERR_FILE_WATCH_TIMEOUT
source {{GetCustomCloudConfigCSEScriptFilepath }}
{{end}}

set +x
ETCD_PEER_CERT=$(echo ${ETCD_PEER_CERTIFICATES} | cut -d'[' -f 2 | cut -d']' -f 1 | cut -d',' -f $((${NODE_INDEX}+1)))
ETCD_PEER_KEY=$(echo ${ETCD_PEER_PRIVATE_KEYS} | cut -d'[' -f 2 | cut -d']' -f 1 | cut -d',' -f $((${NODE_INDEX}+1)))
set -x

if [[ $OS == $COREOS_OS_NAME ]]; then
    echo "Changing default kubectl bin location"
    KUBECTL=/opt/kubectl
fi

if [ -f /var/run/reboot-required ]; then
    REBOOTREQUIRED=true
    trace_info "RebootRequired" "reboot=true"
else
    REBOOTREQUIRED=false
fi

time_metric "ConfigureAdminUser" configureAdminUser

{{- if not NeedsContainerd}}
time_metric "CleanupContainerd" cleanUpContainerd
{{end}}

if [[ "${GPU_NODE}" != "true" ]]; then
    time_metric "CleanupGPUDrivers" cleanUpGPUDrivers
fi

VHD_LOGS_FILEPATH=/opt/azure/vhd-install.complete
if [ -f $VHD_LOGS_FILEPATH ]; then
    echo "detected golden image pre-install"
    time_metric "CleanUpContainerImages" cleanUpContainerImages
    FULL_INSTALL_REQUIRED=false
else
    if [[ "${IS_VHD}" = true ]]; then
        echo "Using VHD distro but file $VHD_LOGS_FILEPATH not found"
        exit $ERR_VHD_FILE_NOT_FOUND
    fi
    FULL_INSTALL_REQUIRED=true
fi

if [[ ( $OS == $UBUNTU_OS_NAME || $OS == $DEBIAN_OS_NAME ) ]] && [ "$FULL_INSTALL_REQUIRED" = "true" ]; then
    time_metric "InstallDeps" installDeps
    if [[ $OS == $UBUNTU_OS_NAME ]]; then
        time_metric "InstallBcc" installBcc
    fi
    {{- if not IsDockerContainerRuntime}}
    time_metric "InstallImg" installImg
    {{end}}
else
    echo "Golden image; skipping dependencies installation"
fi

if [[ $OS == $UBUNTU_OS_NAME ]]; then
    time_metric "EnsureAuditD" ensureAuditD
fi

if [[ "$FULL_INSTALL_REQUIRED" = "true" ]]; then
    time_metric "InstallBpftrace" installBpftrace
fi

{{- if not HasCoreOS}}
time_metric "InstallContainerRuntime" installContainerRuntime
{{end}}

{{- if NeedsContainerd}}
time_metric "InstallContainerd" installContainerd
{{end}}

if [[ -n "${MASTER_NODE}" ]] && [[ -z "${COSMOS_URI}" ]]; then
    {{- if IsDockerContainerRuntime}}
    CLI_TOOL="docker"
    {{else}}
    CLI_TOOL="img"
    {{end}}
    time_metric "InstallEtcd" installEtcd $CLI_TOOL
fi

# this will capture the amount of time to install of the network plugin during cse
time_metric "InstallNetworkPlugin" installNetworkPlugin

{{- if HasNSeriesSKU}}
if [[ "${GPU_NODE}" = true ]]; then
    if $FULL_INSTALL_REQUIRED; then
        time_metric "InstallGPUDrivers" installGPUDrivers
    fi
    time_metric "EnsureGPUDrivers" ensureGPUDrivers
fi
{{end}}

{{- if and IsDockerContainerRuntime HasPrivateAzureRegistryServer}}
docker login -u $SERVICE_PRINCIPAL_CLIENT_ID -p $SERVICE_PRINCIPAL_CLIENT_SECRET {{GetPrivateAzureRegistryServer}}
{{end}}

time_metric "InstallKubeletAndKubectl" installKubeletAndKubectl

if [[ $OS != $COREOS_OS_NAME ]]; then
    time_metric "EnsureRPC" ensureRPC
fi

time_metric "CreateKubeManifestDir" createKubeManifestDir

{{- if HasDCSeriesSKU}}
if [[ "${SGX_NODE}" = true ]]; then
    time_metric "InstallSGXDrivers" installSGXDrivers
fi
{{end}}

{{/* create etcd user if we are configured for etcd */}}
if [[ -n "${MASTER_NODE}" ]] && [[ -z "${COSMOS_URI}" ]]; then
    time_metric "ConfigureEtcdUser" configureEtcdUser
fi

if [[ -n "${MASTER_NODE}" ]]; then
    {{/* this step configures all certs */}}
    {{/* both configs etcd/cosmos */}}
    time_metric "ConfigureSecrets" configureSecrets
fi

{{/* configure etcd if we are configured for etcd */}}
if [[ -n "${MASTER_NODE}" ]] && [[ -z "${COSMOS_URI}" ]]; then
    time_metric "ConfigureEtcd" configureEtcd
else
    time_metric "RemoveEtcd" removeEtcd
fi

{{- if HasCustomSearchDomain}}
wait_for_file 3600 1 {{GetCustomSearchDomainsCSEScriptFilepath}} || exit $ERR_FILE_WATCH_TIMEOUT
{{GetCustomSearchDomainsCSEScriptFilepath}} > /opt/azure/containers/setup-custom-search-domain.log 2>&1 || exit $ERR_CUSTOM_SEARCH_DOMAINS_FAIL
{{end}}

{{- if IsDockerContainerRuntime}}
time_metric "EnsureDocker" ensureDocker
{{else if IsKataContainerRuntime}}
if grep -q vmx /proc/cpuinfo; then
    time_metric "InstallKataContainers" installKataContainersRuntime
fi
{{end}}

time_metric "ConfigureK8s" configureK8s

{{- if IsAzureStackCloud}}
time_metric "ConfigureK8sCustomCloud" configureK8sCustomCloud
    {{- if IsAzureCNI}}
    time_metric "ConfigureAzureStackInterfaces" configureAzureStackInterfaces
    {{end}}
{{end}}

time_metric "ConfigureCNI" configureCNI

{{if or IsClusterAutoscalerAddonEnabled IsACIConnectorAddonEnabled IsAzurePolicyAddonEnabled}}
if [[ -n "${MASTER_NODE}" ]]; then
    time_metric "ConfigAddons" configAddons
fi
{{end}}

{{- if NeedsContainerd}}
time_metric "EnsureContainerd" ensureContainerd
{{end}}

{{- if EnableEncryptionWithExternalKms}}
if [[ -n "${MASTER_NODE}" && "${KMS_PROVIDER_VAULT_NAME}" != "" ]]; then
    time_metric "EnsureKMS" ensureKMS
fi
{{end}}

{{/* configure and enable dhcpv6 for dual stack feature */}}
{{- if IsIPv6DualStackFeatureEnabled}}
time_metric "EnsureDHCPv6" ensureDHCPv6
{{end}}

time_metric "EnsureKubelet" ensureKubelet
time_metric "EnsureJournal" ensureJournal

if [[ -n "${MASTER_NODE}" ]]; then
    if version_gte ${KUBERNETES_VERSION} 1.16; then
        time_metric "EnsureLabelNodes" ensureLabelNodes
    fi
    time_metric "WriteKubeConfig" writeKubeConfig
    if [[ -z "${COSMOS_URI}" ]]; then
        time_metric "EnsureEtcd" ensureEtcd
    fi
    time_metric "EnsureK8sControlPlane" ensureK8sControlPlane
    {{if IsAzurePolicyAddonEnabled}}
    time_metric "EnsureLabelExclusionForAzurePolicyAddon" ensureLabelExclusionForAzurePolicyAddon
    {{end}}
fi

if $FULL_INSTALL_REQUIRED; then
    if [[ $OS == $UBUNTU_OS_NAME ]]; then
        {{/* mitigation for bug https://bugs.launchpad.net/ubuntu/+source/linux/+bug/1676635 */}}
        echo 2dd1ce17-079e-403c-b352-a1921ee207ee > /sys/bus/vmbus/drivers/hv_util/unbind
        sed -i "13i\echo 2dd1ce17-079e-403c-b352-a1921ee207ee > /sys/bus/vmbus/drivers/hv_util/unbind\n" /etc/rc.local
    fi
fi

{{- if not IsAzureStackCloud}}
if [[ $OS == $UBUNTU_OS_NAME ]]; then
    time_metric "PurgeApt" apt_get_purge 20 30 120 apache2-utils &
fi
{{end}}

if $REBOOTREQUIRED; then
    echo 'reboot required, rebooting node in 1 minute'
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

echo "Custom script finished successfully"
echo $(date),$(hostname), endcustomscript>>/opt/m
mkdir -p /opt/azure/containers && touch /opt/azure/containers/provision.complete
ps auxfww > /opt/azure/provision-ps.log &

#EOF
