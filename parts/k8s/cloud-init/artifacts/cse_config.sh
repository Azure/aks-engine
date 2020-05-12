#!/bin/bash
NODE_INDEX=$(hostname | tail -c 2)
NODE_NAME=$(hostname)
PRIVATE_IP=$(hostname -I | cut -d' ' -f1)
ETCD_PEER_URL="https://${PRIVATE_IP}:2380"
ETCD_CLIENT_URL="https://${PRIVATE_IP}:2379"
KUBECTL="/usr/local/bin/kubectl --kubeconfig=/home/$ADMINUSER/.kube/config"
MOUNT_ETCD_SCRIPT=/opt/azure/containers/mountetcd.sh

systemctlEnableAndStart() {
  systemctl_restart 100 5 30 $1
  RESTART_STATUS=$?
  systemctl status $1 --no-pager -l >/var/log/azure/$1-status.log
  if [ $RESTART_STATUS -ne 0 ]; then
    return 1
  fi
  if ! retrycmd 120 5 25 systemctl enable $1; then
    return 1
  fi
}
systemctlEtcd() {
  for i in $(seq 1 60); do
    timeout 30 systemctl daemon-reload
    timeout 30 systemctl restart etcd && break ||
      if [ $i -eq 60 ]; then
        return 1
      else
        $MOUNT_ETCD_SCRIPT
        sleep 5
      fi
  done
  if ! retrycmd 120 5 25 systemctl enable etcd; then
    return 1
  fi
}
configureAdminUser(){
  chage -E -1 -I -1 -m 0 -M 99999 "${ADMINUSER}"
  chage -l "${ADMINUSER}"
}
configureEtcdUser(){
  useradd -U etcd
  chage -E -1 -I -1 -m 0 -M 99999 etcd
  chage -l etcd
  id etcd
}
configureSecrets(){
  APISERVER_PRIVATE_KEY_PATH="/etc/kubernetes/certs/apiserver.key"
  touch "${APISERVER_PRIVATE_KEY_PATH}"
  CA_PRIVATE_KEY_PATH="/etc/kubernetes/certs/ca.key"
  touch "${CA_PRIVATE_KEY_PATH}"
  ETCD_SERVER_PRIVATE_KEY_PATH="/etc/kubernetes/certs/etcdserver.key"
  touch "${ETCD_SERVER_PRIVATE_KEY_PATH}"
  if [[ -z ${COSMOS_URI} ]]; then
    chown etcd:etcd "${ETCD_SERVER_PRIVATE_KEY_PATH}"
  fi
  ETCD_CLIENT_PRIVATE_KEY_PATH="/etc/kubernetes/certs/etcdclient.key"
  touch "${ETCD_CLIENT_PRIVATE_KEY_PATH}"
  ETCD_PEER_PRIVATE_KEY_PATH="/etc/kubernetes/certs/etcdpeer${NODE_INDEX}.key"
  touch "${ETCD_PEER_PRIVATE_KEY_PATH}"
  if [[ -z ${COSMOS_URI} ]]; then
    chown etcd:etcd "${ETCD_PEER_PRIVATE_KEY_PATH}"
  fi
  chmod 0600 "${APISERVER_PRIVATE_KEY_PATH}" "${CA_PRIVATE_KEY_PATH}" "${ETCD_SERVER_PRIVATE_KEY_PATH}" "${ETCD_CLIENT_PRIVATE_KEY_PATH}" "${ETCD_PEER_PRIVATE_KEY_PATH}"
  chown root:root "${APISERVER_PRIVATE_KEY_PATH}" "${CA_PRIVATE_KEY_PATH}" "${ETCD_CLIENT_PRIVATE_KEY_PATH}"
  ETCD_SERVER_CERTIFICATE_PATH="/etc/kubernetes/certs/etcdserver.crt"
  touch "${ETCD_SERVER_CERTIFICATE_PATH}"
  ETCD_CLIENT_CERTIFICATE_PATH="/etc/kubernetes/certs/etcdclient.crt"
  touch "${ETCD_CLIENT_CERTIFICATE_PATH}"
  ETCD_PEER_CERTIFICATE_PATH="/etc/kubernetes/certs/etcdpeer${NODE_INDEX}.crt"
  touch "${ETCD_PEER_CERTIFICATE_PATH}"
  chmod 0644 "${ETCD_SERVER_CERTIFICATE_PATH}" "${ETCD_CLIENT_CERTIFICATE_PATH}" "${ETCD_PEER_CERTIFICATE_PATH}"
  chown root:root "${ETCD_SERVER_CERTIFICATE_PATH}" "${ETCD_CLIENT_CERTIFICATE_PATH}" "${ETCD_PEER_CERTIFICATE_PATH}"

  set +x
  echo "${APISERVER_PRIVATE_KEY}" | base64 --decode >"${APISERVER_PRIVATE_KEY_PATH}"
  echo "${CA_PRIVATE_KEY}" | base64 --decode >"${CA_PRIVATE_KEY_PATH}"
  echo "${ETCD_SERVER_PRIVATE_KEY}" | base64 --decode >"${ETCD_SERVER_PRIVATE_KEY_PATH}"
  echo "${ETCD_CLIENT_PRIVATE_KEY}" | base64 --decode >"${ETCD_CLIENT_PRIVATE_KEY_PATH}"
  echo "${ETCD_PEER_KEY}" | base64 --decode >"${ETCD_PEER_PRIVATE_KEY_PATH}"
  echo "${ETCD_SERVER_CERTIFICATE}" | base64 --decode >"${ETCD_SERVER_CERTIFICATE_PATH}"
  echo "${ETCD_CLIENT_CERTIFICATE}" | base64 --decode >"${ETCD_CLIENT_CERTIFICATE_PATH}"
  echo "${ETCD_PEER_CERT}" | base64 --decode >"${ETCD_PEER_CERTIFICATE_PATH}"
}
configureEtcd() {
  set -x

  ETCD_SETUP_FILE=/opt/azure/containers/setup-etcd.sh
  wait_for_file 1200 1 $ETCD_SETUP_FILE || exit {{GetCSEErrorCode "ERR_ETCD_CONFIG_FAIL"}}
  $ETCD_SETUP_FILE >/opt/azure/containers/setup-etcd.log 2>&1
  RET=$?
  if [ $RET -ne 0 ]; then
    exit $RET
  fi

  if [[ -z ${ETCDCTL_ENDPOINTS} ]]; then
    {{/* Variables necessary for etcdctl are not present */}}
    {{/* Must pull them from /etc/environment */}}
    for entry in $(cat /etc/environment); do
      export ${entry}
    done
  fi

  wait_for_file 1200 1 $MOUNT_ETCD_SCRIPT || exit {{GetCSEErrorCode "ERR_ETCD_CONFIG_FAIL"}}
  $MOUNT_ETCD_SCRIPT || exit {{GetCSEErrorCode "ERR_ETCD_VOL_MOUNT_FAIL"}}
  systemctlEtcd || exit {{GetCSEErrorCode "ERR_ETCD_START_TIMEOUT"}}
  for i in $(seq 1 600); do
    MEMBER="$(sudo -E etcdctl member list | grep -E ${NODE_NAME} | cut -d':' -f 1)"
    if [ "$MEMBER" != "" ]; then
      break
    else
      sleep 1
    fi
  done
  retrycmd 120 5 25 sudo -E etcdctl member update $MEMBER ${ETCD_PEER_URL} || exit {{GetCSEErrorCode "ERR_ETCD_CONFIG_FAIL"}}
}
ensureNTP() {
  systemctlEnableAndStart ntp || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
}
ensureRPC() {
  systemctlEnableAndStart rpcbind || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
  systemctlEnableAndStart rpc-statd || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
}
ensureAuditD() {
  if [[ ${AUDITD_ENABLED} == true ]]; then
    systemctlEnableAndStart auditd || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
  else
    apt_get_purge auditd mlocate &
  fi
}
generateAggregatedAPICerts() {
  AGGREGATED_API_CERTS_SETUP_FILE=/etc/kubernetes/generate-proxy-certs.sh
  wait_for_file 1200 1 $AGGREGATED_API_CERTS_SETUP_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  $AGGREGATED_API_CERTS_SETUP_FILE
}
configureKubeletServerCert() {
  KUBELET_SERVER_PRIVATE_KEY_PATH="/etc/kubernetes/certs/kubeletserver.key"
  KUBELET_SERVER_CERT_PATH="/etc/kubernetes/certs/kubeletserver.crt"

  openssl genrsa -out $KUBELET_SERVER_PRIVATE_KEY_PATH 2048
  openssl req -new -x509 -days 7300 -key $KUBELET_SERVER_PRIVATE_KEY_PATH -out $KUBELET_SERVER_CERT_PATH -subj "/CN=${NODE_NAME}"
}
configureK8s() {
  KUBELET_PRIVATE_KEY_PATH="/etc/kubernetes/certs/client.key"
  touch "${KUBELET_PRIVATE_KEY_PATH}"
  APISERVER_PUBLIC_KEY_PATH="/etc/kubernetes/certs/apiserver.crt"
  touch "${APISERVER_PUBLIC_KEY_PATH}"
  chmod 0600 "${KUBELET_PRIVATE_KEY_PATH}"
  chmod 0644 "${APISERVER_PUBLIC_KEY_PATH}"
  chown root:root "${KUBELET_PRIVATE_KEY_PATH}" "${APISERVER_PUBLIC_KEY_PATH}"

  set +x
  echo "${KUBELET_PRIVATE_KEY}" | base64 --decode >"${KUBELET_PRIVATE_KEY_PATH}"
  echo "${APISERVER_PUBLIC_KEY}" | base64 --decode >"${APISERVER_PUBLIC_KEY_PATH}"
  configureKubeletServerCert
  AZURE_JSON_PATH="/etc/kubernetes/azure.json"
  if [[ -n ${MASTER_NODE} ]]; then
    if [[ ${ENABLE_AGGREGATED_APIS} == True ]]; then
      generateAggregatedAPICerts
    fi
  else
    {{- /* If we are a node vm then we only proceed w/ local azure.json configuration if cloud-init has pre-paved that file */}}
    wait_for_file 1 1 $AZURE_JSON_PATH || return
  fi

  {{/* Perform the required JSON escaping */}}
  SERVICE_PRINCIPAL_CLIENT_SECRET=${SERVICE_PRINCIPAL_CLIENT_SECRET//\\/\\\\}
  SERVICE_PRINCIPAL_CLIENT_SECRET=${SERVICE_PRINCIPAL_CLIENT_SECRET//\"/\\\"}
  cat <<EOF >"${AZURE_JSON_PATH}"
{
    "cloud":"{{GetTargetEnvironment}}",
    "tenantId": "${TENANT_ID}",
    "subscriptionId": "${SUBSCRIPTION_ID}",
    "aadClientId": "${SERVICE_PRINCIPAL_CLIENT_ID}",
    "aadClientSecret": "${SERVICE_PRINCIPAL_CLIENT_SECRET}",
    "resourceGroup": "${RESOURCE_GROUP}",
    "location": "${LOCATION}",
    "vmType": "${VM_TYPE}",
    "subnetName": "${SUBNET}",
    "securityGroupName": "${NETWORK_SECURITY_GROUP}",
    "vnetName": "${VIRTUAL_NETWORK}",
    "vnetResourceGroup": "${VIRTUAL_NETWORK_RESOURCE_GROUP}",
    "routeTableName": "${ROUTE_TABLE}",
    "primaryAvailabilitySetName": "${PRIMARY_AVAILABILITY_SET}",
    "primaryScaleSetName": "${PRIMARY_SCALE_SET}",
    "cloudProviderBackoffMode": "${CLOUDPROVIDER_BACKOFF_MODE}",
    "cloudProviderBackoff": ${CLOUDPROVIDER_BACKOFF},
    "cloudProviderBackoffRetries": ${CLOUDPROVIDER_BACKOFF_RETRIES},
    "cloudProviderBackoffExponent": ${CLOUDPROVIDER_BACKOFF_EXPONENT},
    "cloudProviderBackoffDuration": ${CLOUDPROVIDER_BACKOFF_DURATION},
    "cloudProviderBackoffJitter": ${CLOUDPROVIDER_BACKOFF_JITTER},
    "cloudProviderRatelimit": ${CLOUDPROVIDER_RATELIMIT},
    "cloudProviderRateLimitQPS": ${CLOUDPROVIDER_RATELIMIT_QPS},
    "cloudProviderRateLimitBucket": ${CLOUDPROVIDER_RATELIMIT_BUCKET},
    "cloudProviderRatelimitQPSWrite": ${CLOUDPROVIDER_RATELIMIT_QPS_WRITE},
    "cloudProviderRatelimitBucketWrite": ${CLOUDPROVIDER_RATELIMIT_BUCKET_WRITE},
    "useManagedIdentityExtension": ${USE_MANAGED_IDENTITY_EXTENSION},
    "userAssignedIdentityID": "${USER_ASSIGNED_IDENTITY_ID}",
    "useInstanceMetadata": ${USE_INSTANCE_METADATA},
    "loadBalancerSku": "${LOAD_BALANCER_SKU}",
    "disableOutboundSNAT": ${LOAD_BALANCER_DISABLE_OUTBOUND_SNAT},
    "excludeMasterFromStandardLB": ${EXCLUDE_MASTER_FROM_STANDARD_LB},
    "providerVaultName": "${KMS_PROVIDER_VAULT_NAME}",
    "maximumLoadBalancerRuleCount": ${MAXIMUM_LOADBALANCER_RULE_COUNT},
    "providerKeyName": "k8s",
    "providerKeyVersion": ""
}
EOF
  set -x
  if [[ ${CLOUDPROVIDER_BACKOFF_MODE} == "v2" ]]; then
    sed -i "/cloudProviderBackoffExponent/d" $AZURE_JSON_PATH
    sed -i "/cloudProviderBackoffJitter/d" $AZURE_JSON_PATH
  fi
}

installNetworkPlugin() {
{{- if IsAzureCNI}}
  installAzureCNI
{{end}}
  installCNI
  rm -rf $CNI_DOWNLOADS_DIR &
}
installCNI() {
  CNI_TGZ_TMP=${CNI_PLUGINS_URL##*/}
  if [[ ! -f "$CNI_DOWNLOADS_DIR/${CNI_TGZ_TMP}" ]]; then
    downloadCNI
  fi
  mkdir -p $CNI_BIN_DIR
  tar -xzf "$CNI_DOWNLOADS_DIR/${CNI_TGZ_TMP}" -C $CNI_BIN_DIR
  chown -R root:root $CNI_BIN_DIR
  chmod -R 755 $CNI_BIN_DIR
}
{{- if IsAzureCNI}}
installAzureCNI() {
  CNI_TGZ_TMP=${VNET_CNI_PLUGINS_URL##*/}
  if [[ ! -f "$CNI_DOWNLOADS_DIR/${CNI_TGZ_TMP}" ]]; then
    downloadAzureCNI
  fi
  mkdir -p $CNI_CONFIG_DIR
  chown -R root:root $CNI_CONFIG_DIR
  chmod 755 $CNI_CONFIG_DIR
  mkdir -p $CNI_BIN_DIR
  tar -xzf "$CNI_DOWNLOADS_DIR/${CNI_TGZ_TMP}" -C $CNI_BIN_DIR
}
{{end}}
configureCNI() {
  {{/* needed for the iptables rules to work on bridges */}}
  retrycmd 120 5 25 modprobe br_netfilter || exit {{GetCSEErrorCode "ERR_MODPROBE_FAIL"}}
  echo -n "br_netfilter" >/etc/modules-load.d/br_netfilter.conf
  configureAzureCNI
  {{if HasCiliumNetworkPlugin}}
  systemctl enable sys-fs-bpf.mount
  systemctl restart sys-fs-bpf.mount
  REBOOTREQUIRED=true
  {{end}}
{{- if IsAzureStackCloud}}
  if [[ ${NETWORK_PLUGIN} == "azure" ]]; then
    {{/* set environment to mas when using Azure CNI on Azure Stack */}}
    {{/* shellcheck disable=SC2002,SC2005 */}}
    echo $(cat "$CNI_CONFIG_DIR/10-azure.conflist" | jq '.plugins[0].ipam.environment = "mas"') >"$CNI_CONFIG_DIR/10-azure.conflist"
  fi
{{end}}
}
configureAzureCNI() {
  if [[ "${NETWORK_PLUGIN}" == "azure" ]]; then
    mv $CNI_BIN_DIR/10-azure.conflist $CNI_CONFIG_DIR/
    chmod 600 $CNI_CONFIG_DIR/10-azure.conflist
    if [[ "${IS_IPV6_DUALSTACK_FEATURE_ENABLED}" == "true" ]]; then
      echo $(cat "$CNI_CONFIG_DIR/10-azure.conflist" | jq '.plugins[0].ipv6Mode="ipv6nat"') > "$CNI_CONFIG_DIR/10-azure.conflist"
    fi
    if [[ {{GetKubeProxyMode}} == "ipvs" ]]; then
      serviceCidrs={{GetServiceCidr}}
      echo $(cat "$CNI_CONFIG_DIR/10-azure.conflist" | jq  --arg serviceCidrs $serviceCidrs '.plugins[0]+={serviceCidrs: $serviceCidrs}') > /etc/cni/net.d/10-azure.conflist
    fi
    if [[ "${NETWORK_POLICY}" == "calico" ]]; then
      sed -i 's#"mode":"bridge"#"mode":"transparent"#g' $CNI_CONFIG_DIR/10-azure.conflist
    elif [[ "${NETWORK_POLICY}" == "antrea" ]]; then
      sed -i 's#"mode":"bridge"#"mode":"transparent"#g' $CNI_CONFIG_DIR/10-azure.conflist
    elif [[ "${NETWORK_POLICY}" == "" || "${NETWORK_POLICY}" == "none" ]] && [[ "${NETWORK_MODE}" == "transparent" ]]; then
      sed -i 's#"mode":"bridge"#"mode":"transparent"#g' $CNI_CONFIG_DIR/10-azure.conflist
    fi
    /sbin/ebtables -t nat --list
  fi
}
{{- if NeedsContainerd}}
installContainerd() {
  removeMoby
  CURRENT_VERSION=$(containerd -version | cut -d " " -f 3 | sed 's|v||')
  if [[ $CURRENT_VERSION != "${CONTAINERD_VERSION}" ]]; then
    os_lower=$(echo ${OS} | tr '[:upper:]' '[:lower:]')
    if [[ ${OS} == "${UBUNTU_OS_NAME}" ]]; then
      url_path="${os_lower}/${UBUNTU_RELEASE}/multiarch/prod"
    elif [[ ${OS} == "${DEBIAN_OS_NAME}" ]]; then
      url_path="${os_lower}/${UBUNTU_RELEASE}/prod"
    else
      exit 25
    fi
    removeContainerd
    retrycmd_no_stats 120 5 25 curl https://packages.microsoft.com/config/ubuntu/${UBUNTU_RELEASE}/prod.list >/tmp/microsoft-prod.list || exit 25
    retrycmd 10 5 10 cp /tmp/microsoft-prod.list /etc/apt/sources.list.d/ || exit 25
    retrycmd_no_stats 120 5 25 curl https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor >/tmp/microsoft.gpg || exit 26
    retrycmd 10 5 10 cp /tmp/microsoft.gpg /etc/apt/trusted.gpg.d/ || exit 26
    apt_get_update || exit 99
    apt_get_install 20 30 120 moby-containerd=${CONTAINERD_VERSION}* --allow-downgrades || exit 27
  fi
}
ensureContainerd() {
  wait_for_file 1200 1 /etc/systemd/system/containerd.service.d/exec_start.conf || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  wait_for_file 1200 1 /etc/containerd/config.toml || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  {{- if HasKubeReservedCgroup}}
  wait_for_file 1200 1 /etc/systemd/system/containerd.service.d/kubereserved-slice.conf|| exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  {{- end}}
  systemctlEnableAndStart containerd || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
}
{{end}}
{{- if IsDockerContainerRuntime}}
ensureDocker() {
  DOCKER_SERVICE_EXEC_START_FILE=/etc/systemd/system/docker.service.d/exec_start.conf
  wait_for_file 1200 1 $DOCKER_SERVICE_EXEC_START_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  usermod -aG docker ${ADMINUSER}
  DOCKER_MOUNT_FLAGS_SYSTEMD_FILE=/etc/systemd/system/docker.service.d/clear_mount_propagation_flags.conf
  {{- if HasKubeReservedCgroup}}
  DOCKER_SLICE_FILE=/etc/systemd/system/docker.service.d/kubereserved-slice.conf
  wait_for_file 1200 1 $DOCKER_SLICE_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  {{- end}}
  DOCKER_JSON_FILE=/etc/docker/daemon.json
  for i in $(seq 1 1200); do
    if [ -s $DOCKER_JSON_FILE ]; then
      jq '.' <$DOCKER_JSON_FILE && break
    fi
    if [ $i -eq 1200 ]; then
      exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
    else
      sleep 1
    fi
  done
  systemctlEnableAndStart docker || exit {{GetCSEErrorCode "ERR_DOCKER_START_FAIL"}}
  {{/* Delay start of docker-monitor for 30 mins after booting */}}
  DOCKER_MONITOR_SYSTEMD_TIMER_FILE=/etc/systemd/system/docker-monitor.timer
  wait_for_file 1200 1 $DOCKER_MONITOR_SYSTEMD_TIMER_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  DOCKER_MONITOR_SYSTEMD_FILE=/etc/systemd/system/docker-monitor.service
  wait_for_file 1200 1 $DOCKER_MONITOR_SYSTEMD_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  systemctlEnableAndStart docker-monitor.timer || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
}
{{end}}
{{- if EnableEncryptionWithExternalKms}}
ensureKMS() {
  systemctlEnableAndStart kms || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
}
{{end}}
{{- if IsIPv6Enabled}}
ensureDHCPv6() {
  wait_for_file 3600 1 {{GetDHCPv6ServiceCSEScriptFilepath}} || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  wait_for_file 3600 1 {{GetDHCPv6ConfigCSEScriptFilepath}} || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  systemctlEnableAndStart dhcpv6 || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
  retrycmd 120 5 25 modprobe ip6_tables || exit {{GetCSEErrorCode "ERR_MODPROBE_FAIL"}}
}
{{end}}
ensureKubelet() {
  wait_for_file 1200 1 /etc/sysctl.d/11-aks-engine.conf || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  sysctl_reload 10 5 120 || exit {{GetCSEErrorCode "ERR_SYSCTL_RELOAD"}}
  KUBELET_DEFAULT_FILE=/etc/default/kubelet
  wait_for_file 1200 1 $KUBELET_DEFAULT_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  KUBECONFIG_FILE=/var/lib/kubelet/kubeconfig
  wait_for_file 1200 1 $KUBECONFIG_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  KUBELET_RUNTIME_CONFIG_SCRIPT_FILE=/opt/azure/containers/kubelet.sh
  wait_for_file 1200 1 $KUBELET_RUNTIME_CONFIG_SCRIPT_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  {{- if HasKubeReservedCgroup}}
  KUBERESERVED_SLICE_FILE=/etc/systemd/system/{{- GetKubeReservedCgroup -}}.slice
  wait_for_file 1200 1 $KUBERESERVED_SLICE_FILE || exit {{GetCSEErrorCode "ERR_KUBERESERVED_SLICE_SETUP_FAIL"}}
  KUBELET_SLICE_FILE=/etc/systemd/system/kubelet.service.d/kubereserved-slice.conf
  wait_for_file 1200 1 $KUBELET_SLICE_FILE || exit {{GetCSEErrorCode "ERR_KUBELET_SLICE_SETUP_FAIL"}}
  {{- end}}
  systemctlEnableAndStart kubelet || exit {{GetCSEErrorCode "ERR_KUBELET_START_FAIL"}}
  {{if HasCiliumNetworkPolicy}}
  while [ ! -f /etc/cni/net.d/05-cilium.conf ]; do
    sleep 3
  done
  {{end}}
  {{if HasAntreaNetworkPolicy}}
  if [[ "${NETWORK_PLUGIN}" = "azure" ]]; then
    while ! $(grep -sq "antrea" $CNI_CONFIG_DIR/10-azure.conflist); do
      sleep 3
    done
  else
    while [ ! -f $CNI_CONFIG_DIR/10-antrea.conflist ]; do
      sleep 3
    done
  fi
  {{end}}
  {{if HasFlannelNetworkPlugin}}
  while [ ! -f /etc/cni/net.d/10-flannel.conf ]; do
    sleep 3
  done
  {{end}}
}
ensureLabelNodes() {
  LABEL_NODES_SCRIPT_FILE=/opt/azure/containers/label-nodes.sh
  wait_for_file 1200 1 $LABEL_NODES_SCRIPT_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  LABEL_NODES_SYSTEMD_FILE=/etc/systemd/system/label-nodes.service
  wait_for_file 1200 1 $LABEL_NODES_SYSTEMD_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  systemctlEnableAndStart label-nodes || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
}
{{- if IsAADPodIdentityAddonEnabled}}
ensureTaints() {
  wait_for_file 1200 1 /opt/azure/containers/untaint-nodes.sh || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  wait_for_file 1200 1 /etc/systemd/system/untaint-nodes.service || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  systemctlEnableAndStart untaint-nodes || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
}
{{end}}
ensureJournal() {
  {
    echo "Storage=persistent"
    echo "SystemMaxUse=1G"
    echo "RuntimeMaxUse=1G"
    echo "ForwardToSyslog=yes"
  } >>/etc/systemd/journald.conf
  systemctlEnableAndStart systemd-journald || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
}
installKubeletAndKubectl() {
  if [[ ! -f "/usr/local/bin/kubectl-${KUBERNETES_VERSION}" ]]; then
    if version_gte ${KUBERNETES_VERSION} 1.17; then
      extractKubeBinaries
    else
      if [[ $CONTAINER_RUNTIME == "docker" ]]; then
        extractHyperkube "docker"
      else
        extractHyperkube "img"
      fi
    fi
  fi
  mv "/usr/local/bin/kubelet-${KUBERNETES_VERSION}" "/usr/local/bin/kubelet"
  mv "/usr/local/bin/kubectl-${KUBERNETES_VERSION}" "/usr/local/bin/kubectl"
  chmod a+x /usr/local/bin/kubelet /usr/local/bin/kubectl
  rm -rf /usr/local/bin/kubelet-* /usr/local/bin/kubectl-* /home/hyperkube-downloads &
}
ensureK8sControlPlane() {
  if $REBOOTREQUIRED || [ "$NO_OUTBOUND" = "true" ]; then
    return
  fi
  retrycmd 120 5 25 $KUBECTL 2>/dev/null cluster-info || exit {{GetCSEErrorCode "ERR_K8S_RUNNING_TIMEOUT"}}
}
{{- if IsAzurePolicyAddonEnabled}}
ensureLabelExclusionForAzurePolicyAddon() {
  GATEKEEPER_NAMESPACE="gatekeeper-system"
  retrycmd 120 5 25 $KUBECTL create ns --save-config $GATEKEEPER_NAMESPACE 2>/dev/null || exit {{GetCSEErrorCode "ERR_K8S_RUNNING_TIMEOUT"}}

  retrycmd 120 5 25 $KUBECTL label ns kube-system control-plane=controller-manager 2>/dev/null || exit {{GetCSEErrorCode "ERR_K8S_RUNNING_TIMEOUT"}}
}
{{end}}
ensureEtcd() {
  retrycmd 120 5 25 curl --cacert /etc/kubernetes/certs/ca.crt --cert /etc/kubernetes/certs/etcdclient.crt --key /etc/kubernetes/certs/etcdclient.key ${ETCD_CLIENT_URL}/v2/machines || exit {{GetCSEErrorCode "ERR_ETCD_RUNNING_TIMEOUT"}}
}
createKubeManifestDir() {
  KUBEMANIFESTDIR=/etc/kubernetes/manifests
  mkdir -p $KUBEMANIFESTDIR
}
writeKubeConfig() {
  local DIR=/home/$ADMINUSER/.kube
  local FILE=$DIR/config
  mkdir -p $DIR
  touch $FILE
  chown $ADMINUSER:$ADMINUSER $DIR $FILE
  chmod 700 $DIR
  chmod 600 $FILE
  set +x
  echo "
---
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: \"$CA_CERTIFICATE\"
    server: $KUBECONFIG_SERVER
  name: \"$MASTER_FQDN\"
contexts:
- context:
    cluster: \"$MASTER_FQDN\"
    user: \"$MASTER_FQDN-admin\"
  name: \"$MASTER_FQDN\"
current-context: \"$MASTER_FQDN\"
kind: Config
users:
- name: \"$MASTER_FQDN-admin\"
  user:
    client-certificate-data: \"$KUBECONFIG_CERTIFICATE\"
    client-key-data: \"$KUBECONFIG_KEY\"
" >$FILE
  set -x
  KUBECTL="$KUBECTL --kubeconfig=$FILE"
}
{{- if IsClusterAutoscalerAddonEnabled}}
configClusterAutoscalerAddon() {
  CLUSTER_AUTOSCALER_ADDON_FILE=/etc/kubernetes/addons/cluster-autoscaler.yaml
  wait_for_file 1200 1 $CLUSTER_AUTOSCALER_ADDON_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  sed -i "s|<clientID>|$(echo $SERVICE_PRINCIPAL_CLIENT_ID | base64)|g" $CLUSTER_AUTOSCALER_ADDON_FILE
  sed -i "s|<clientSec>|$(echo $SERVICE_PRINCIPAL_CLIENT_SECRET | base64)|g" $CLUSTER_AUTOSCALER_ADDON_FILE
  sed -i "s|<subID>|$(echo $SUBSCRIPTION_ID | base64)|g" $CLUSTER_AUTOSCALER_ADDON_FILE
  sed -i "s|<tenantID>|$(echo $TENANT_ID | base64)|g" $CLUSTER_AUTOSCALER_ADDON_FILE
  sed -i "s|<rg>|$(echo $RESOURCE_GROUP | base64)|g" $CLUSTER_AUTOSCALER_ADDON_FILE
}
{{end}}
{{- if IsACIConnectorAddonEnabled}}
configACIConnectorAddon() {
  ACI_CONNECTOR_CREDENTIALS=$(printf '{"clientId": "%s", "clientSecret": "%s", "tenantId": "%s", "subscriptionId": "%s", "activeDirectoryEndpointUrl": "https://login.microsoftonline.com","resourceManagerEndpointUrl": "https://management.azure.com/", "activeDirectoryGraphResourceId": "https://graph.windows.net/", "sqlManagementEndpointUrl": "https://management.core.windows.net:8443/", "galleryEndpointUrl": "https://gallery.azure.com/", "managementEndpointUrl": "https://management.core.windows.net/"}' "$SERVICE_PRINCIPAL_CLIENT_ID" "$SERVICE_PRINCIPAL_CLIENT_SECRET" "$TENANT_ID" "$SUBSCRIPTION_ID" | base64 -w 0)

  openssl req -newkey rsa:4096 -new -nodes -x509 -days 3650 -keyout /etc/kubernetes/certs/aci-connector-key.pem -out /etc/kubernetes/certs/aci-connector-cert.pem -subj "/C=US/ST=CA/L=virtualkubelet/O=virtualkubelet/OU=virtualkubelet/CN=virtualkubelet"
  ACI_CONNECTOR_KEY=$(base64 /etc/kubernetes/certs/aci-connector-key.pem -w0)
  ACI_CONNECTOR_CERT=$(base64 /etc/kubernetes/certs/aci-connector-cert.pem -w0)

  ACI_CONNECTOR_ADDON_FILE=/etc/kubernetes/addons/aci-connector-deployment.yaml
  wait_for_file 1200 1 $ACI_CONNECTOR_ADDON_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  sed -i "s|<creds>|$ACI_CONNECTOR_CREDENTIALS|g" $ACI_CONNECTOR_ADDON_FILE
  sed -i "s|<rgName>|$RESOURCE_GROUP|g" $ACI_CONNECTOR_ADDON_FILE
  sed -i "s|<cert>|$ACI_CONNECTOR_CERT|g" $ACI_CONNECTOR_ADDON_FILE
  sed -i "s|<key>|$ACI_CONNECTOR_KEY|g" $ACI_CONNECTOR_ADDON_FILE
}
{{end}}
{{- if IsAzurePolicyAddonEnabled}}
configAzurePolicyAddon() {
  AZURE_POLICY_ADDON_FILE=/etc/kubernetes/addons/azure-policy-deployment.yaml
  sed -i "s|<resourceId>|/subscriptions/$SUBSCRIPTION_ID/resourceGroups/$RESOURCE_GROUP|g" $AZURE_POLICY_ADDON_FILE
}
{{end}}
{{- if or IsClusterAutoscalerAddonEnabled IsACIConnectorAddonEnabled IsAzurePolicyAddonEnabled}}
configAddons() {
  {{if IsClusterAutoscalerAddonEnabled}}
  if [[ ${CLUSTER_AUTOSCALER_ADDON} == true ]]; then
    configClusterAutoscalerAddon
  fi
  {{end}}
  {{if IsACIConnectorAddonEnabled}}
  if [[ ${ACI_CONNECTOR_ADDON} == True ]]; then
    configACIConnectorAddon
  fi
  {{end}}
  {{if IsAzurePolicyAddonEnabled}}
  configAzurePolicyAddon
  {{end}}
}
{{end}}
{{- if HasNSeriesSKU}}
{{- /* installNvidiaDrivers is idempotent, it will uninstall itself if it is already installed, and then install anew */}}
installNvidiaDrivers() {
  NVIDIA_DKMS_DIR="/var/lib/dkms/nvidia/${GPU_DV}"
  KERNEL_NAME=$(uname -r)
  if [ -d $NVIDIA_DKMS_DIR ]; then
    dkms remove -m nvidia -v $GPU_DV -k $KERNEL_NAME
  fi
  local log_file="/var/log/nvidia-installer-$(date +%s).log"
  sh $GPU_DEST/nvidia-drivers-$GPU_DV -s -k=$KERNEL_NAME --log-file-name=$log_file -a --no-drm --dkms --utility-prefix="${GPU_DEST}" --opengl-prefix="${GPU_DEST}"
}
configGPUDrivers() {
  {{/* only install the runtime since nvidia-docker2 has a hard dep on docker CE packages. */}}
  {{/* we will manually install nvidia-docker2 */}}
  rmmod nouveau
  echo blacklist nouveau >>/etc/modprobe.d/blacklist.conf
  retrycmd_no_stats 120 5 25 update-initramfs -u || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_CONFIG"}}
  wait_for_apt_locks
  {{/* if the unattened upgrade is turned on, and it may takes 10 min to finish the installation, and we use the 1 second just to try to get the lock more aggressively */}}
  retrycmd 600 1 3600 apt-get -o Dpkg::Options::="--force-confold" install -y nvidia-container-runtime="${NVIDIA_CONTAINER_RUNTIME_VERSION}+${NVIDIA_DOCKER_SUFFIX}" || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_CONFIG"}}
  tmpDir=$GPU_DEST/tmp
  (
    set -e -o pipefail
    cd "${tmpDir}"
    wait_for_apt_locks
    dpkg-deb -R ./nvidia-docker2*.deb "${tmpDir}/pkg" || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_CONFIG"}}
    cp -r ${tmpDir}/pkg/usr/* /usr/ || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_CONFIG"}}
  )
  rm -rf $GPU_DEST/tmp
  retrycmd 120 5 25 pkill -SIGHUP dockerd || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_CONFIG"}}
  mkdir -p $GPU_DEST/lib64 $GPU_DEST/overlay-workdir
  retrycmd 120 5 25 mount -t overlay -o lowerdir=/usr/lib/x86_64-linux-gnu,upperdir=${GPU_DEST}/lib64,workdir=${GPU_DEST}/overlay-workdir none /usr/lib/x86_64-linux-gnu || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_CONFIG"}}
  export -f installNvidiaDrivers
  retrycmd 3 1 600 bash -c installNvidiaDrivers || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_START_FAIL"}}
  mv ${GPU_DEST}/bin/* /usr/bin
  echo "${GPU_DEST}/lib64" >/etc/ld.so.conf.d/nvidia.conf
  retrycmd 120 5 25 ldconfig || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_START_FAIL"}}
  umount -l /usr/lib/x86_64-linux-gnu
  retrycmd 120 5 25 nvidia-modprobe -u -c0 || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_START_FAIL"}}
  retrycmd 120 5 25 nvidia-smi || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_START_FAIL"}}
  retrycmd 120 5 25 ldconfig || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_START_FAIL"}}
}
ensureGPUDrivers() {
  configGPUDrivers
  systemctlEnableAndStart nvidia-modprobe || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_START_FAIL"}}
}
{{end}}
{{- if HasDCSeriesSKU}}
installSGXDrivers() {
  [[ $UBUNTU_RELEASE == "18.04" || $UBUNTU_RELEASE == "16.04" ]] || exit 92

  local packages="make gcc dkms"
  wait_for_apt_locks
  retrycmd 30 5 3600 apt-get -y install "$packages" || exit 90

  local oe_dir=/opt/azure/containers/oe
  rm -rf ${oe_dir}
  mkdir -p ${oe_dir}
  pushd ${oe_dir} || exit
  retrycmd 10 10 120 curl -fsSL -O "https://download.01.org/intel-sgx/latest/version.xml" || exit 90
  dcap_version="$(grep dcap version.xml | grep -o -E "[.0-9]+")"
  sgx_driver_folder_url="https://download.01.org/intel-sgx/sgx-dcap/$dcap_version/linux"
  retrycmd 10 10 120 curl -fsSL -O "$sgx_driver_folder_url/SHA256SUM_dcap_$dcap_version.cfg" || exit 90
  matched_line="$(grep "distro/ubuntuServer$UBUNTU_RELEASE/sgx_linux_x64_driver_.*bin" SHA256SUM_dcap_$dcap_version.cfg)"
  read -ra tmp_array <<<"$matched_line"
  sgx_driver_sha256sum_expected="${tmp_array[0]}"
  sgx_driver_remote_path="${tmp_array[1]}"
  sgx_driver_url="${sgx_driver_folder_url}/${sgx_driver_remote_path}"
  sgx_driver=$(basename "$sgx_driver_url")

  retrycmd 10 10 120 curl -fsSL -O "${sgx_driver_url}" || exit 90
  read -ra tmp_array <<<"$(sha256sum ./"$sgx_driver")"
  sgx_driver_sha256sum_real="${tmp_array[0]}"
  [[ $sgx_driver_sha256sum_real == "$sgx_driver_sha256sum_expected" ]] || exit 93

  chmod a+x ./"${sgx_driver}"
  if ! ./"${sgx_driver}"; then
    popd || exit
    exit 91
  fi
  popd || exit
  rm -rf ${oe_dir}
}
{{end}}
{{- if HasVHDDistroNodes}}
cleanUpContainerImages() {
  docker rmi $(docker images --format '{{OpenBraces}}.Repository{{CloseBraces}}:{{OpenBraces}}.Tag{{CloseBraces}}' | grep -vE "${KUBERNETES_VERSION}$|${KUBERNETES_VERSION}-|${KUBERNETES_VERSION}_" | grep 'hyperkube') &
  docker rmi $(docker images --format '{{OpenBraces}}.Repository{{CloseBraces}}:{{OpenBraces}}.Tag{{CloseBraces}}' | grep -vE "${KUBERNETES_VERSION}$|${KUBERNETES_VERSION}-|${KUBERNETES_VERSION}_" | grep 'cloud-controller-manager') &
  docker rmi $(docker images --format '{{OpenBraces}}.Repository{{CloseBraces}}:{{OpenBraces}}.Tag{{CloseBraces}}' | grep -vE "${ETCD_VERSION}$|${ETCD_VERSION}-|${ETCD_VERSION}_" | grep 'etcd') &
  if [ "$IS_HOSTED_MASTER" = "false" ]; then
    docker rmi $(docker images --format '{{OpenBraces}}.Repository{{CloseBraces}}:{{OpenBraces}}.Tag{{CloseBraces}}' | grep 'hcp-tunnel-front') &
    docker rmi $(docker images --format '{{OpenBraces}}.Repository{{CloseBraces}}:{{OpenBraces}}.Tag{{CloseBraces}}' | grep 'kube-svc-redirect') &
    docker rmi $(docker images --format '{{OpenBraces}}.Repository{{CloseBraces}}:{{OpenBraces}}.Tag{{CloseBraces}}' | grep 'nginx') &
  fi

  docker rmi registry:2.7.1 &
}
cleanUpGPUDrivers() {
  rm -Rf $GPU_DEST
  rm -f /etc/apt/sources.list.d/nvidia-docker.list
  apt-key del $(apt-key list | grep NVIDIA -B 1 | head -n 1 | cut -d "/" -f 2 | cut -d " " -f 1)
}
cleanUpContainerd() {
  rm -Rf $CONTAINERD_DOWNLOADS_DIR
}
{{end}}
removeEtcd() {
  rm -rf /usr/bin/etcd
}
removeMoby() {
  apt_get_purge moby-engine moby-cli || exit 27
}
removeContainerd() {
  apt_get_purge moby-containerd || exit 27
}
#EOF
