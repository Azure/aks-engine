#!/bin/bash
ERR_FILE_WATCH_TIMEOUT=6 # Timeout waiting for a file
set -x
echo $(date),$(hostname), startcustomscript>>/opt/m
AZURE_STACK_ENV="azurestackcloud"

script_lib=/opt/azure/containers/provision_source.sh
for i in $(seq 1 3600); do
    if [ -s $script_lib ]; then
        grep -Fq '#HELPERSEOF' $script_lib && break
    fi
    if [ $i -eq 3600 ]; then
        exit $ERR_FILE_WATCH_TIMEOUT
    else
        sleep 1
    fi
done
sed -i "/#HELPERSEOF/d" $script_lib
source $script_lib

install_script=/opt/azure/containers/provision_installs.sh
wait_for_file 3600 1 $install_script || exit $ERR_FILE_WATCH_TIMEOUT
source $install_script

config_script=/opt/azure/containers/provision_configs.sh
wait_for_file 3600 1 $config_script || exit $ERR_FILE_WATCH_TIMEOUT
source $config_script

if [[ "${TARGET_ENVIRONMENT,,}" == "${AZURE_STACK_ENV}"  ]]; then
    config_script_custom_cloud=/opt/azure/containers/provision_configs_custom_cloud.sh
    wait_for_file 3600 1 $config_script_custom_cloud || exit $ERR_FILE_WATCH_TIMEOUT
    source $config_script_custom_cloud
fi

CUSTOM_SEARCH_DOMAIN_SCRIPT=/opt/azure/containers/setup-custom-search-domains.sh

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
else
    REBOOTREQUIRED=false
fi

if [[ "$CONTAINER_RUNTIME" != "kata-containers" ]] && [[ "$CONTAINER_RUNTIME" != "containerd" ]]; then
  cleanUpContainerd
fi
if [[ "${GPU_NODE}" != "true" ]]; then
  cleanUpGPUDrivers
fi

VHD_LOGS_FILEPATH=/opt/azure/vhd-install.complete
if [ -f $VHD_LOGS_FILEPATH ]; then
    echo "detected golden image pre-install"
    cleanUpContainerImages
    FULL_INSTALL_REQUIRED=false
else
    if [[ "${IS_VHD}" = true ]]; then
        echo "Using VHD distro but file $VHD_LOGS_FILEPATH not found"
        exit $ERR_VHD_FILE_NOT_FOUND
    fi
    FULL_INSTALL_REQUIRED=true
fi

if [[ $OS == $UBUNTU_OS_NAME ]] && [ "$FULL_INSTALL_REQUIRED" = "true" ]; then
    installDeps
else
    echo "Golden image; skipping dependencies installation"
fi
ensureAuditD

if [[ -n "${MASTER_NODE}" ]] && [[ -z "${COSMOS_URI}" ]]; then
    installEtcd
fi

if [[ $OS != $COREOS_OS_NAME ]]; then
    installContainerRuntime
fi
installNetworkPlugin
if [[ "$CONTAINER_RUNTIME" == "kata-containers" ]] || [[ "$CONTAINER_RUNTIME" == "containerd" ]]; then
    installContainerd
fi
if [[ "${GPU_NODE}" = true ]]; then
    if $FULL_INSTALL_REQUIRED; then
        installGPUDrivers
    fi
    ensureGPUDrivers
fi
installKubeletAndKubectl
if [[ $OS != $COREOS_OS_NAME ]]; then
    ensureRPC
fi
createKubeManifestDir
if [[ "${SGX_NODE}" = true ]]; then
    installSGXDrivers
fi

# create etcd user if we are configured for etcd
if [[ -n "${MASTER_NODE}" ]] && [[ -z "${COSMOS_URI}" ]]; then
  configureEtcdUser
fi

if [[ -n "${MASTER_NODE}" ]]; then
  # this step configures all certs
  # both configs etcd/cosmos
  configureSecrets
fi
# configure etcd if we are configured for etcd
if [[ -n "${MASTER_NODE}" ]] && [[ -z "${COSMOS_URI}" ]]; then
    configureEtcd
else
    removeEtcd
fi


if [ -f $CUSTOM_SEARCH_DOMAIN_SCRIPT ]; then
    $CUSTOM_SEARCH_DOMAIN_SCRIPT > /opt/azure/containers/setup-custom-search-domain.log 2>&1 || exit $ERR_CUSTOM_SEARCH_DOMAINS_FAIL
fi

if [[ "$CONTAINER_RUNTIME" == "docker" ]]; then
    ensureDocker
elif [[ "$CONTAINER_RUNTIME" == "kata-containers" ]]; then
    if grep -q vmx /proc/cpuinfo; then
        installKataContainersRuntime
    fi
fi

configureK8s

if [[ "${TARGET_ENVIRONMENT,,}" == "${AZURE_STACK_ENV,,}"  ]]; then
    configureK8sCustomCloud
    if [[ "${NETWORK_PLUGIN,,}" = "azure" ]]; then
        configureAzureStackInterfaces
    fi
fi

configureCNI

if [[ -n "${MASTER_NODE}" ]]; then
    configAddons
fi

if [[ "$CONTAINER_RUNTIME" == "kata-containers" ]] || [[ "$CONTAINER_RUNTIME" == "containerd" ]]; then
    ensureContainerd
fi

if [[ -n "${MASTER_NODE}" && "${KMS_PROVIDER_VAULT_NAME}" != "" ]]; then
    ensureKMS
fi

# configure and enable dhcpv6 for dual stack feature
if [ "$IS_IPV6_DUALSTACK_FEATURE_ENABLED" = "true" ]; then
    dhcpv6_systemd_service=/etc/systemd/system/dhcpv6.service
    dhcpv6_configuration_script=/opt/azure/containers/enable-dhcpv6.sh
    wait_for_file 3600 1 $dhcpv6_systemd_service || exit $ERR_FILE_WATCH_TIMEOUT
    wait_for_file 3600 1 $dhcpv6_configuration_script || exit $ERR_FILE_WATCH_TIMEOUT
    ensureDHCPv6

    retrycmd_if_failure 120 5 25 modprobe ip6_tables || exit $ERR_MODPROBE_FAIL
fi

ensureKubelet
ensureJournal

if [[ -n "${MASTER_NODE}" ]]; then
    if version_gte ${KUBERNETES_VERSION} 1.16; then
      ensureLabelNodes
    fi
    writeKubeConfig
    if [[ -z "${COSMOS_URI}" ]]; then
      ensureEtcd
    fi
    ensureK8sControlPlane
fi

if $FULL_INSTALL_REQUIRED; then
    if [[ $OS == $UBUNTU_OS_NAME ]]; then
        # mitigation for bug https://bugs.launchpad.net/ubuntu/+source/linux/+bug/1676635
        echo 2dd1ce17-079e-403c-b352-a1921ee207ee > /sys/bus/vmbus/drivers/hv_util/unbind
        sed -i "13i\echo 2dd1ce17-079e-403c-b352-a1921ee207ee > /sys/bus/vmbus/drivers/hv_util/unbind\n" /etc/rc.local
    fi
fi

echo "Custom script finished successfully"

echo $(date),$(hostname), endcustomscript>>/opt/m
mkdir -p /opt/azure/containers && touch /opt/azure/containers/provision.complete
ps auxfww > /opt/azure/provision-ps.log &

if [[ "${TARGET_ENVIRONMENT,,}" != "${AZURE_STACK_ENV}"  ]]; then
    # TODO: remove once ACR is available on Azure Stack
    apt_get_purge 20 30 120 apache2-utils || exit $ERR_APT_PURGE_FAIL
fi

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
#EOF
