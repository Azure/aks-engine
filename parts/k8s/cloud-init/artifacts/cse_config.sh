#!/bin/bash
NODE_INDEX=$(hostname | tail -c 2)
NODE_NAME=$(hostname)
KUBECTL="/usr/local/bin/kubectl --kubeconfig=/home/$ADMINUSER/.kube/config"
ADDONS_DIR=/etc/kubernetes/addons
POD_SECURITY_POLICY_SPEC=$ADDONS_DIR/pod-security-policy.yaml
ADDON_MANAGER_SPEC=/etc/kubernetes/manifests/kube-addon-manager.yaml
GET_KUBELET_LOGS="journalctl -u kubelet --no-pager"

systemctlEnableAndStart() {
  local ret
  systemctl_restart 100 5 30 $1
  ret=$?
  systemctl status $1 --no-pager -l >/var/log/azure/$1-status.log
  if [ $ret -ne 0 ]; then
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
  local apiserver_key="/etc/kubernetes/certs/apiserver.key" ca_key="/etc/kubernetes/certs/ca.key" etcdserver_key="/etc/kubernetes/certs/etcdserver.key"
  touch "${apiserver_key}"
  touch "${ca_key}"
  touch "${etcdserver_key}"
  if [[ -z ${COSMOS_URI} ]]; then
    chown etcd:etcd "${etcdserver_key}"
  fi
  local etcdclient_key="/etc/kubernetes/certs/etcdclient.key" etcdpeer_key="/etc/kubernetes/certs/etcdpeer${NODE_INDEX}.key"
  touch "${etcdclient_key}"
  touch "${etcdpeer_key}"
  if [[ -z ${COSMOS_URI} ]]; then
    chown etcd:etcd "${etcdpeer_key}"
  fi
  chmod 0600 "${apiserver_key}" "${ca_key}" "${etcdserver_key}" "${etcdclient_key}" "${etcdpeer_key}"
  chown root:root "${apiserver_key}" "${ca_key}" "${etcdclient_key}"
  local etcdserver_crt="/etc/kubernetes/certs/etcdserver.crt" etcdclient_crt="/etc/kubernetes/certs/etcdclient.crt" etcdpeer_crt="/etc/kubernetes/certs/etcdpeer${NODE_INDEX}.crt"
  touch "${etcdserver_crt}"
  touch "${etcdclient_crt}"
  touch "${etcdpeer_crt}"
  chmod 0644 "${etcdserver_crt}" "${etcdclient_crt}" "${etcdpeer_crt}"
  chown root:root "${etcdserver_crt}" "${etcdclient_crt}" "${etcdpeer_crt}"

  set +x
  echo "${APISERVER_PRIVATE_KEY}" | base64 --decode >"${apiserver_key}"
  echo "${CA_PRIVATE_KEY}" | base64 --decode >"${ca_key}"
  echo "${ETCD_SERVER_PRIVATE_KEY}" | base64 --decode >"${etcdserver_key}"
  echo "${ETCD_CLIENT_PRIVATE_KEY}" | base64 --decode >"${etcdclient_key}"
  echo "${ETCD_PEER_KEY}" | base64 --decode >"${etcdpeer_key}"
  echo "${ETCD_SERVER_CERTIFICATE}" | base64 --decode >"${etcdserver_crt}"
  echo "${ETCD_CLIENT_CERTIFICATE}" | base64 --decode >"${etcdclient_crt}"
  echo "${ETCD_PEER_CERT}" | base64 --decode >"${etcdpeer_crt}"
}
configureEtcd() {
  set -x

  local ret f=/opt/azure/containers/setup-etcd.sh etcd_peer_url="https://${PRIVATE_IP}:2380"
  wait_for_file 1200 1 $f || exit {{GetCSEErrorCode "ERR_ETCD_CONFIG_FAIL"}}
  $f >/opt/azure/containers/setup-etcd.log 2>&1
  ret=$?
  if [ $ret -ne 0 ]; then
    exit $ret
  fi

  if [[ -z ${ETCDCTL_ENDPOINTS} ]]; then
    {{/* Variables necessary for etcdctl are not present */}}
    {{/* Must pull them from /etc/environment */}}
    for entry in $(cat /etc/environment); do
      export ${entry}
    done
  fi

  chown -R etcd:etcd /var/lib/etcddisk
  systemctlEtcd || exit {{GetCSEErrorCode "ERR_ETCD_START_TIMEOUT"}}
  for i in $(seq 1 600); do
    MEMBER="$(sudo -E etcdctl member list | grep -E ${NODE_NAME} | cut -d':' -f 1)"
    if [ "$MEMBER" != "" ]; then
      break
    else
      sleep 1
    fi
  done
  retrycmd 120 5 25 sudo -E etcdctl member update $MEMBER ${etcd_peer_url} || exit {{GetCSEErrorCode "ERR_ETCD_CONFIG_FAIL"}}
}
configureChrony() {
  sed -i "s/makestep.*/makestep 1.0 -1/g" /etc/chrony/chrony.conf
  echo "refclock PHC /dev/ptp0 poll 3 dpoll -2 offset 0" >> /etc/chrony/chrony.conf
}
ensureChrony() {
  systemctlEnableAndStart chrony || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
}
disable1804SystemdResolved() {
  {{/* Ignoring systemd-resolved query service but using its resolv.conf file */}}
  {{/* This is the simplest approach to workaround resolved issues without completely uninstall it */}}
  [ -f /run/systemd/resolve/resolv.conf ] && sudo ln -sf /run/systemd/resolve/resolv.conf /etc/resolv.conf
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
ensureCron() {
  local s=/lib/systemd/system/cron.service
  if [[ -f ${s} ]]; then
    if ! grep -q 'Restart=' ${s}; then
      sed -i 's/\[Service\]/[Service]\nRestart=always/' ${s}
      systemctlEnableAndStart cron
    fi
  fi
}
generateAggregatedAPICerts() {
  local f=/etc/kubernetes/generate-proxy-certs.sh
  wait_for_file 1200 1 $f || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  $f
}
configureKubeletServerCert() {
  local kubeletserver_key="/etc/kubernetes/certs/kubeletserver.key" kubeletserver_crt="/etc/kubernetes/certs/kubeletserver.crt"

  openssl genrsa -out $kubeletserver_key 2048
  openssl req -new -x509 -days 7300 -key $kubeletserver_key -out $kubeletserver_crt -subj "/CN=${NODE_NAME}"
}
configureK8s() {
  local client_key="/etc/kubernetes/certs/client.key" apiserver_crt="/etc/kubernetes/certs/apiserver.crt" azure_json="/etc/kubernetes/azure.json"
  touch "${client_key}"
  chmod 0600 "${client_key}"
  chown root:root "${client_key}"
  if [[ -n ${MASTER_NODE} ]]; then
    touch "${apiserver_crt}"
    chmod 0644 "${apiserver_crt}"
    chown root:root "${apiserver_crt}"
  fi
  set +x
  echo "${KUBELET_PRIVATE_KEY}" | base64 --decode >"${client_key}"
  configureKubeletServerCert
  if [[ -n ${MASTER_NODE} ]]; then
    echo "${APISERVER_PUBLIC_KEY}" | base64 --decode >"${apiserver_crt}"
    if [[ ${ENABLE_AGGREGATED_APIS} == True ]]; then
      generateAggregatedAPICerts
    fi
  else
    {{- /* If we are a node that does not need azure.json (cloud-init tells us), then return immediately */}}
    wait_for_file 1 1 /opt/azure/needs_azure.json || return
  fi

  touch $azure_json
  chmod 0600 $azure_json
  chown root:root $azure_json
  {{/* Perform the required JSON escaping */}}
  local sp_secret=${SERVICE_PRINCIPAL_CLIENT_SECRET//\\/\\\\}
  sp_secret=${SERVICE_PRINCIPAL_CLIENT_SECRET//\"/\\\"}
  cat <<EOF >"${azure_json}"
{
    "cloud":"{{GetTargetEnvironment}}",
    "tenantId": "${TENANT_ID}",
    "subscriptionId": "${SUBSCRIPTION_ID}",
    "aadClientId": "${SERVICE_PRINCIPAL_CLIENT_ID}",
    "aadClientSecret": "${sp_secret}",
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
    "providerKeyVersion": "",
    "enableMultipleStandardLoadBalancers": ${ENABLE_MULTIPLE_STANDARD_LOAD_BALANCERS},
    "tags": "${TAGS}"
}
EOF
  set -x
  if [[ ${CLOUDPROVIDER_BACKOFF_MODE} == "v2" ]]; then
    sed -i "/cloudProviderBackoffExponent/d" $azure_json
    sed -i "/cloudProviderBackoffJitter/d" $azure_json
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
  touch /var/run/reboot-required
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
  local tmpDir=$(mktemp -d "$(pwd)/XXX")
  if [[ "${NETWORK_PLUGIN}" == "azure" ]]; then
    mv $CNI_BIN_DIR/10-azure.conflist $CNI_CONFIG_DIR/
    chmod 600 $CNI_CONFIG_DIR/10-azure.conflist
{{- if IsIPv6DualStackFeatureEnabled}}
    jq '.plugins[0].ipv6Mode="ipv6nat"' "$CNI_CONFIG_DIR/10-azure.conflist" > $tmpDir/tmp
    mv $tmpDir/tmp $CNI_CONFIG_DIR/10-azure.conflist
{{- end}}
    if [[ {{GetKubeProxyMode}} == "ipvs" ]]; then
      serviceCidrs={{GetServiceCidr}}
      jq --arg serviceCidrs $serviceCidrs '.plugins[0]+={serviceCidrs: $serviceCidrs}' "$CNI_CONFIG_DIR/10-azure.conflist" > $tmpDir/tmp
      mv $tmpDir/tmp $CNI_CONFIG_DIR/10-azure.conflist
    fi
    if [[ "${NETWORK_MODE}" == "bridge" ]]; then
      jq '.plugins[0].mode="bridge"' "$CNI_CONFIG_DIR/10-azure.conflist" > $tmpDir/tmp
      jq '.plugins[0].bridge="azure0"' "$tmpDir/tmp" > $tmpDir/tmp2
      mv $tmpDir/tmp2 $CNI_CONFIG_DIR/10-azure.conflist
    else
      jq '.plugins[0].mode="transparent"' "$CNI_CONFIG_DIR/10-azure.conflist" > $tmpDir/tmp
      mv $tmpDir/tmp $CNI_CONFIG_DIR/10-azure.conflist
    fi
    /sbin/ebtables -t nat --list
  fi
  rm -Rf $tmpDir
}
enableCRISystemdMonitor() {
  wait_for_file 1200 1 /etc/systemd/system/docker-monitor.service || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  systemctlEnableAndStart docker-monitor || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
}
{{- if NeedsContainerd}}
installContainerd() {
  removeMoby
  local v
  v=$(containerd -version | cut -d " " -f 3 | sed 's|v||')
  if [[ $v != "${CONTAINERD_VERSION}"* ]]; then
    os_lower=$(echo ${OS} | tr '[:upper:]' '[:lower:]')
    if [[ ${OS} == "${UBUNTU_OS_NAME}" ]]; then
      url_path="${os_lower}/${UBUNTU_RELEASE}/multiarch/prod"
    elif [[ ${OS} == "${DEBIAN_OS_NAME}" ]]; then
      url_path="${os_lower}/${UBUNTU_RELEASE}/prod"
    else
      exit 25
    fi
    removeContainerd
    apt_get_update || exit 99
    apt_get_install 20 30 120 moby-runc moby-containerd=${CONTAINERD_VERSION}* --allow-downgrades || exit 27
  fi
}
ensureContainerd() {
  wait_for_file 1200 1 /etc/systemd/system/containerd.service.d/exec_start.conf || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  wait_for_file 1200 1 /etc/containerd/config.toml || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  {{- if HasKubeReservedCgroup}}
  wait_for_file 1200 1 /etc/systemd/system/containerd.service.d/kubereserved-slice.conf|| exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  {{- end}}
  systemctlEnableAndStart containerd || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
  enableCRISystemdMonitor
}
{{end}}
{{- if IsDockerContainerRuntime}}
ensureDocker() {
  wait_for_file 1200 1 /etc/systemd/system/docker.service.d/exec_start.conf || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  usermod -aG docker ${ADMINUSER}
  if [[ $OS != $FLATCAR_OS_NAME ]]; then
    wait_for_file 1200 1 /etc/systemd/system/docker.service.d/clear_mount_propagation_flags.conf || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  fi
  {{- if HasKubeReservedCgroup}}
  wait_for_file 1200 1 /etc/systemd/system/docker.service.d/kubereserved-slice.conf || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  {{- end}}
  local daemon_json=/etc/docker/daemon.json
  for i in $(seq 1 1200); do
    if [ -s $daemon_json ]; then
      jq '.' <$daemon_json && break
    fi
    if [ $i -eq 1200 ]; then
      exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
    else
      sleep 1
    fi
  done
  systemctlEnableAndStart docker || exit {{GetCSEErrorCode "ERR_DOCKER_START_FAIL"}}
  enableCRISystemdMonitor
}
{{end}}
{{- if IsIPv6Enabled}}
ensureDHCPv6() {
  if [[ ${UBUNTU_RELEASE} == "16.04" ]]; then
    wait_for_file 3600 1 {{GetDHCPv6ServiceCSEScriptFilepath}} || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
    wait_for_file 3600 1 {{GetDHCPv6ConfigCSEScriptFilepath}} || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
    systemctlEnableAndStart dhcpv6 || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
    retrycmd 120 5 25 modprobe ip6_tables || exit {{GetCSEErrorCode "ERR_MODPROBE_FAIL"}}
  fi
}
{{end}}
{{- if EnableEncryptionWithExternalKms}}
ensureKMSKeyvaultKey() {
    wait_for_file 3600 1 {{GetKMSKeyvaultKeyServiceCSEScriptFilepath}} || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
    wait_for_file 3600 1 {{GetKMSKeyvaultKeyCSEScriptFilepath}} || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
    systemctlEnableAndStart kms-keyvault-key || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
}
{{end}}
ensureKubelet() {
  wait_for_file 1200 1 /etc/sysctl.d/11-aks-engine.conf || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  sysctl_reload 10 5 120 || exit {{GetCSEErrorCode "ERR_SYSCTL_RELOAD"}}
  wait_for_file 1200 1 /etc/default/kubelet || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  wait_for_file 1200 1 /var/lib/kubelet/kubeconfig || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  if [[ -n ${MASTER_NODE} ]]; then
{{- if IsMasterVirtualMachineScaleSets}}
    sed -i "s|<SERVERIP>|https://$PRIVATE_IP:443|g" "/var/lib/kubelet/kubeconfig" || exit {{GetCSEErrorCode "ERR_KUBELET_START_FAIL"}}
{{- end}}
    local f=/etc/kubernetes/manifests/kube-apiserver.yaml
    wait_for_file 1200 1 $f || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
    sed -i "s|<advertiseAddr>|$PRIVATE_IP|g" $f
  fi
  wait_for_file 1200 1 /opt/azure/containers/kubelet.sh || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
{{- if HasKubeReservedCgroup}}
  wait_for_file 1200 1 /etc/systemd/system/{{- GetKubeReservedCgroup -}}.slice || exit {{GetCSEErrorCode "ERR_KUBERESERVED_SLICE_SETUP_FAIL"}}
  wait_for_file 1200 1 /etc/systemd/system/kubelet.service.d/kubereserved-slice.conf || exit {{GetCSEErrorCode "ERR_KUBELET_SLICE_SETUP_FAIL"}}
{{- end}}
  if [[ -n ${MASTER_NODE} ]]; then
    systemctlEnableAndStart kubelet || exit {{GetCSEErrorCode "ERR_KUBELET_START_FAIL"}}
  else
{{- if not RunUnattendedUpgradesOnBootstrap}}
    systemctlEnableAndStart kubelet || exit {{GetCSEErrorCode "ERR_KUBELET_START_FAIL"}}
{{else}}
    systemctl_enable 100 5 30 kubelet || exit {{GetCSEErrorCode "ERR_KUBELET_START_FAIL"}}
{{- end}}
  fi
{{- if HasKubeletHealthZPort}}
  wait_for_file 1200 1 /etc/systemd/system/kubelet-monitor.service || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  if [[ -n ${MASTER_NODE} ]]; then
    systemctlEnableAndStart kubelet-monitor || exit {{GetCSEErrorCode "ERR_KUBELET_START_FAIL"}}
  else
  {{- if not RunUnattendedUpgradesOnBootstrap}}
    systemctlEnableAndStart kubelet-monitor || exit {{GetCSEErrorCode "ERR_KUBELET_START_FAIL"}}
  {{else}}
    systemctl_enable 100 5 30 kubelet-monitor || exit {{GetCSEErrorCode "ERR_KUBELET_START_FAIL"}}
  {{- end}}
  fi
{{- end}}
}

ensureAddons() {
{{- if IsDashboardAddonEnabled}} {{/* Note: dashboard addon is deprecated */}}
  retrycmd 120 5 30 $KUBECTL get namespace kubernetes-dashboard || exit_cse {{GetCSEErrorCode "ERR_ADDONS_START_FAIL"}} $GET_KUBELET_LOGS
{{- end}}
{{- if IsAzurePolicyAddonEnabled}}
  retrycmd 120 5 30 $KUBECTL get namespace gatekeeper-system || exit_cse {{GetCSEErrorCode "ERR_ADDONS_START_FAIL"}} $GET_KUBELET_LOGS
{{- end}}
{{- if not HasCustomPodSecurityPolicy}}
  retrycmd 120 5 30 $KUBECTL get podsecuritypolicy privileged restricted || exit_cse {{GetCSEErrorCode "ERR_ADDONS_START_FAIL"}} $GET_KUBELET_LOGS
{{- end}}
  rm -Rf ${ADDONS_DIR}/init
  replaceAddonsInit
  {{/* Force re-load all addons because we have changed the source location for addon specs */}}
  while [ ! $(${KUBECTL} get pods -l app=kube-addon-manager -n kube-system --template={{.status.phase}}) == "Running"]; do
    sleep 3
  done
  retrycmd 10 5 30 ${KUBECTL} delete pods -l app=kube-addon-manager -n kube-system --force --grace-period 0 || \
  exit_cse {{GetCSEErrorCode "ERR_ADDONS_START_FAIL"}} $GET_KUBELET_LOGS
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
replaceAddonsInit() {
  wait_for_file 1200 1 $ADDON_MANAGER_SPEC || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  sed -i "s|${ADDONS_DIR}/init|${ADDONS_DIR}|g" $ADDON_MANAGER_SPEC || exit {{GetCSEErrorCode "ERR_ADDONS_START_FAIL"}}
}
ensureLabelNodes() {
  wait_for_file 1200 1 /opt/azure/containers/label-nodes.sh || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  wait_for_file 1200 1 /etc/systemd/system/label-nodes.service || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
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
  local binPath=/usr/local/bin
  if [[ $OS == $FLATCAR_OS_NAME ]]; then
    binPath=/opt/bin
  fi
  if [[ ! -f "${binPath}/kubectl-${KUBERNETES_VERSION}" ]] || [[ -n "${CUSTOM_HYPERKUBE_IMAGE}" ]] || [[ -n "${KUBE_BINARY_URL}" ]]; then
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
  mv "${binPath}/kubelet-${KUBERNETES_VERSION}" "${binPath}/kubelet"
  mv "${binPath}/kubectl-${KUBERNETES_VERSION}" "${binPath}/kubectl"
  chmod a+x ${binPath}/kubelet ${binPath}/kubectl
  rm -rf ${binPath}/kubelet-* ${binPath}/kubectl-* /home/hyperkube-downloads &
}
ensureK8sControlPlane() {
  if [ "$NO_OUTBOUND" = "true" ]; then
    return
  fi
  retrycmd 120 5 25 $KUBECTL 2>/dev/null cluster-info || exit_cse {{GetCSEErrorCode "ERR_K8S_RUNNING_TIMEOUT"}} $GET_KUBELET_LOGS
}
{{- if IsAzurePolicyAddonEnabled}}
ensureLabelExclusionForAzurePolicyAddon() {
  retrycmd 120 5 25 $KUBECTL label ns kube-system control-plane=controller-manager --overwrite 2>/dev/null || exit_cse {{GetCSEErrorCode "ERR_K8S_RUNNING_TIMEOUT"}} $GET_KUBELET_LOGS
}
{{end}}
ensureEtcd() {
  local etcd_client_url="https://${PRIVATE_IP}:2379"
  retrycmd 120 5 25 curl --cacert /etc/kubernetes/certs/ca.crt --cert /etc/kubernetes/certs/etcdclient.crt --key /etc/kubernetes/certs/etcdclient.key ${etcd_client_url}/v2/machines || exit {{GetCSEErrorCode "ERR_ETCD_RUNNING_TIMEOUT"}}
  wait_for_file 1200 1 /etc/systemd/system/etcd-monitor.service || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  systemctlEnableAndStart etcd-monitor || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
}
createKubeManifestDir() {
  mkdir -p /etc/kubernetes/manifests
}
writeKubeConfig() {
  local d=/home/$ADMINUSER/.kube
  local f=$d/config
{{- if HasBlockOutboundInternet}}
  local server=https://localhost
{{else}}
  local server=$KUBECONFIG_SERVER
{{- end}}
  mkdir -p $d
  touch $f
  chown $ADMINUSER:$ADMINUSER $d $f
  chmod 700 $d
  chmod 600 $f
  set +x
  echo "
---
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: \"$CA_CERTIFICATE\"
    server: $server
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
" >$f
  set -x
}
{{- if IsClusterAutoscalerAddonEnabled}}
configClusterAutoscalerAddon() {
  local f=$ADDONS_DIR/cluster-autoscaler.yaml
  wait_for_file 1200 1 $f || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  sed -i "s|<clientID>|$(echo $SERVICE_PRINCIPAL_CLIENT_ID | base64)|g" $f
  sed -i "s|<clientSec>|$(echo $SERVICE_PRINCIPAL_CLIENT_SECRET | base64)|g" $f
  sed -i "s|<subID>|$(echo $SUBSCRIPTION_ID | base64)|g" $f
  sed -i "s|<tenantID>|$(echo $TENANT_ID | base64)|g" $f
  sed -i "s|<rg>|$(echo $RESOURCE_GROUP | base64)|g" $f
}
{{end}}
{{- if IsAzurePolicyAddonEnabled}}
configAzurePolicyAddon() {
  sed -i "s|<resourceId>|/subscriptions/$SUBSCRIPTION_ID/resourceGroups/$RESOURCE_GROUP|g" $ADDONS_DIR/azure-policy-deployment.yaml
}
{{end}}

configAddons() {
  {{if IsClusterAutoscalerAddonEnabled}}
  if [[ ${CLUSTER_AUTOSCALER_ADDON} == true ]]; then
    configClusterAutoscalerAddon
  fi
  {{end}}
  {{if IsAzurePolicyAddonEnabled}}
  configAzurePolicyAddon
  {{end}}
  {{- if and (not HasCustomPodSecurityPolicy) IsPodSecurityPolicyAddonEnabled}}
  wait_for_file 1200 1 $POD_SECURITY_POLICY_SPEC || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  mkdir -p $ADDONS_DIR/init && cp $POD_SECURITY_POLICY_SPEC $ADDONS_DIR/init/ || exit {{GetCSEErrorCode "ERR_ADDONS_START_FAIL"}}
  {{- end}}
}
{{- if and HasNSeriesSKU IsNvidiaDevicePluginAddonEnabled}}
{{- /* installNvidiaDrivers is idempotent, it will uninstall itself if it is already installed, and then install anew */}}
installNvidiaDrivers() {
  local d="/var/lib/dkms/nvidia/${GPU_DV}" k log_file="/var/log/nvidia-installer-$(date +%s).log"
  k=$(uname -r)
  if [ -d $d ]; then
    dkms remove -m nvidia -v $GPU_DV -k $k
  fi
  sh $GPU_DEST/nvidia-drivers-$GPU_DV -s -k=$k --log-file-name=$log_file -a --no-drm --dkms --utility-prefix="${GPU_DEST}" --opengl-prefix="${GPU_DEST}"
}
configGPUDrivers() {
  rmmod nouveau
  echo blacklist nouveau >>/etc/modprobe.d/blacklist.conf
  retrycmd_no_stats 120 5 25 update-initramfs -u || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_CONFIG"}}
  wait_for_apt_locks
  dpkg -i $(ls ${APT_CACHE_DIR}libnvidia-container1*) || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_CONFIG"}}
  dpkg -i $(ls ${APT_CACHE_DIR}libnvidia-container-tools*) || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_CONFIG"}}
  dpkg -i $(ls ${APT_CACHE_DIR}nvidia-container-toolkit*) || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_CONFIG"}}
  dpkg -i $(ls ${APT_CACHE_DIR}nvidia-container-runtime*) || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_CONFIG"}}
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

  local packages="make gcc dkms" oe_dir="/opt/azure/containers/oe"
  wait_for_apt_locks
  retrycmd 30 5 3600 apt-get -y install "$packages" || exit 90
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
  {{- if NeedsContainerd}}
  docker rmi -f $(docker images -a -q) &
  {{else}}
  docker rmi $(docker images --format '{{OpenBraces}}.Repository{{CloseBraces}}:{{OpenBraces}}.Tag{{CloseBraces}}' | grep -vE "${ETCD_VERSION}$|${ETCD_VERSION}-|${ETCD_VERSION}_" | grep 'etcd') &
  docker rmi $(docker images --format '{{OpenBraces}}.Repository{{CloseBraces}}:{{OpenBraces}}.Tag{{CloseBraces}}' | grep -vE "${KUBERNETES_VERSION}$|${KUBERNETES_VERSION}-|${KUBERNETES_VERSION}_" | grep 'kube-proxy') &
  docker rmi $(docker images --format '{{OpenBraces}}.Repository{{CloseBraces}}:{{OpenBraces}}.Tag{{CloseBraces}}' | grep -vE "${KUBERNETES_VERSION}$|${KUBERNETES_VERSION}-|${KUBERNETES_VERSION}_" | grep 'kube-controller-manager') &
  docker rmi $(docker images --format '{{OpenBraces}}.Repository{{CloseBraces}}:{{OpenBraces}}.Tag{{CloseBraces}}' | grep -vE "${KUBERNETES_VERSION}$|${KUBERNETES_VERSION}-|${KUBERNETES_VERSION}_" | grep 'kube-apiserver') &
  docker rmi $(docker images --format '{{OpenBraces}}.Repository{{CloseBraces}}:{{OpenBraces}}.Tag{{CloseBraces}}' | grep -vE "${KUBERNETES_VERSION}$|${KUBERNETES_VERSION}-|${KUBERNETES_VERSION}_" | grep 'kube-scheduler') &
  docker rmi registry:2.7.1 &
  ctr -n=k8s.io image rm $(ctr -n=k8s.io images ls -q) &
  {{- end}}
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
exit_cse() {
  local exit_code=$1
  shift
  $@ >> {{GetLinuxCSELogPath}} &
  exit $exit_code
}
#EOF
