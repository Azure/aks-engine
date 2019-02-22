#!/bin/bash
NODE_INDEX=$(hostname | tail -c 2)
NODE_NAME=$(hostname)
PRIVATE_IP=$(hostname -I | cut -d' ' -f1)
ETCD_PEER_URL="https://${PRIVATE_IP}:2380"
ETCD_CLIENT_URL="https://${PRIVATE_IP}:2379"

systemctlEnableAndStart() {
    systemctl_restart 100 5 30 $1
    RESTART_STATUS=$?
    systemctl status $1 --no-pager -l > /var/log/azure/$1-status.log
    if [ $RESTART_STATUS -ne 0 ]; then
        echo "$1 could not be started"
        return 1
    fi
    retrycmd_if_failure 120 5 25 systemctl enable $1
    if [ $? -ne 0 ]; then
        echo "$1 could not be enabled by systemctl"
        return 1
    fi
}

configureEtcdUser(){
    useradd -U "etcd"
    usermod -p "$(head -c 32 /dev/urandom | base64)" "etcd"
    passwd -u "etcd"
    id "etcd"
}

configureSecrets(){
    APISERVER_PRIVATE_KEY_PATH="/etc/kubernetes/certs/apiserver.key"
    touch "${APISERVER_PRIVATE_KEY_PATH}"
    chmod 0600 "${APISERVER_PRIVATE_KEY_PATH}"
    chown root:root "${APISERVER_PRIVATE_KEY_PATH}"

    CA_PRIVATE_KEY_PATH="/etc/kubernetes/certs/ca.key"
    touch "${CA_PRIVATE_KEY_PATH}"
    chmod 0600 "${CA_PRIVATE_KEY_PATH}"
    chown root:root "${CA_PRIVATE_KEY_PATH}"

    ETCD_SERVER_PRIVATE_KEY_PATH="/etc/kubernetes/certs/etcdserver.key"
    touch "${ETCD_SERVER_PRIVATE_KEY_PATH}"
    chmod 0600 "${ETCD_SERVER_PRIVATE_KEY_PATH}"
    if [[ -z "${COSMOS_URI}" ]]; then
      chown etcd:etcd "${ETCD_SERVER_PRIVATE_KEY_PATH}"
    fi

    ETCD_CLIENT_PRIVATE_KEY_PATH="/etc/kubernetes/certs/etcdclient.key"
    touch "${ETCD_CLIENT_PRIVATE_KEY_PATH}"
    chmod 0600 "${ETCD_CLIENT_PRIVATE_KEY_PATH}"
    chown root:root "${ETCD_CLIENT_PRIVATE_KEY_PATH}"

    ETCD_PEER_PRIVATE_KEY_PATH="/etc/kubernetes/certs/etcdpeer${NODE_INDEX}.key"
    touch "${ETCD_PEER_PRIVATE_KEY_PATH}"
    chmod 0600 "${ETCD_PEER_PRIVATE_KEY_PATH}"
    if [[ -z "${COSMOS_URI}" ]]; then
      chown etcd:etcd "${ETCD_PEER_PRIVATE_KEY_PATH}"
    fi

    ETCD_SERVER_CERTIFICATE_PATH="/etc/kubernetes/certs/etcdserver.crt"
    touch "${ETCD_SERVER_CERTIFICATE_PATH}"
    chmod 0644 "${ETCD_SERVER_CERTIFICATE_PATH}"
    chown root:root "${ETCD_SERVER_CERTIFICATE_PATH}"

    ETCD_CLIENT_CERTIFICATE_PATH="/etc/kubernetes/certs/etcdclient.crt"
    touch "${ETCD_CLIENT_CERTIFICATE_PATH}"
    chmod 0644 "${ETCD_CLIENT_CERTIFICATE_PATH}"
    chown root:root "${ETCD_CLIENT_CERTIFICATE_PATH}"

    ETCD_PEER_CERTIFICATE_PATH="/etc/kubernetes/certs/etcdpeer${NODE_INDEX}.crt"
    touch "${ETCD_PEER_CERTIFICATE_PATH}"
    chmod 0644 "${ETCD_PEER_CERTIFICATE_PATH}"
    chown root:root "${ETCD_PEER_CERTIFICATE_PATH}"

    set +x
    echo "${APISERVER_PRIVATE_KEY}" | base64 --decode > "${APISERVER_PRIVATE_KEY_PATH}"
    echo "${CA_PRIVATE_KEY}" | base64 --decode > "${CA_PRIVATE_KEY_PATH}"
    echo "${ETCD_SERVER_PRIVATE_KEY}" | base64 --decode > "${ETCD_SERVER_PRIVATE_KEY_PATH}"
    echo "${ETCD_CLIENT_PRIVATE_KEY}" | base64 --decode > "${ETCD_CLIENT_PRIVATE_KEY_PATH}"
    echo "${ETCD_PEER_KEY}" | base64 --decode > "${ETCD_PEER_PRIVATE_KEY_PATH}"
    echo "${ETCD_SERVER_CERTIFICATE}" | base64 --decode > "${ETCD_SERVER_CERTIFICATE_PATH}"
    echo "${ETCD_CLIENT_CERTIFICATE}" | base64 --decode > "${ETCD_CLIENT_CERTIFICATE_PATH}"
    echo "${ETCD_PEER_CERT}" | base64 --decode > "${ETCD_PEER_CERTIFICATE_PATH}"
}

configureEtcd() {
    set -x

    ETCD_SETUP_FILE=/opt/azure/containers/setup-etcd.sh
    wait_for_file 1200 1 $ETCD_SETUP_FILE || exit $ERR_ETCD_CONFIG_FAIL
    $ETCD_SETUP_FILE > /opt/azure/containers/setup-etcd.log 2>&1
    RET=$?
    if [ $RET -ne 0 ]; then
        exit $RET
    fi

    MOUNT_ETCD_FILE=/opt/azure/containers/mountetcd.sh
    wait_for_file 1200 1 $MOUNT_ETCD_FILE || exit $ERR_ETCD_CONFIG_FAIL
    $MOUNT_ETCD_FILE || exit $ERR_ETCD_VOL_MOUNT_FAIL
    systemctlEnableAndStart etcd || exit $ERR_ETCD_START_TIMEOUT
    for i in $(seq 1 600); do
        MEMBER="$(sudo etcdctl member list | grep -E ${NODE_NAME} | cut -d':' -f 1)"
        if [ "$MEMBER" != "" ]; then
            break
        else
            sleep 1
        fi
    done
    retrycmd_if_failure 120 5 25 sudo etcdctl member update $MEMBER ${ETCD_PEER_URL} || exit $ERR_ETCD_CONFIG_FAIL
}

ensureRPC() {
    systemctlEnableAndStart rpcbind || exit $ERR_SYSTEMCTL_START_FAIL
    systemctlEnableAndStart rpc-statd || exit $ERR_SYSTEMCTL_START_FAIL
}

runAptDaily() {
    wait_for_apt_locks
    /usr/lib/apt/apt.systemd.daily
}

generateAggregatedAPICerts() {
    AGGREGATED_API_CERTS_SETUP_FILE=/etc/kubernetes/generate-proxy-certs.sh
    wait_for_file 1200 1 $AGGREGATED_API_CERTS_SETUP_FILE || exit $ERR_FILE_WATCH_TIMEOUT
    $AGGREGATED_API_CERTS_SETUP_FILE
}

configureK8s() {
    KUBELET_PRIVATE_KEY_PATH="/etc/kubernetes/certs/client.key"
    touch "${KUBELET_PRIVATE_KEY_PATH}"
    chmod 0600 "${KUBELET_PRIVATE_KEY_PATH}"
    chown root:root "${KUBELET_PRIVATE_KEY_PATH}"

    APISERVER_PUBLIC_KEY_PATH="/etc/kubernetes/certs/apiserver.crt"
    touch "${APISERVER_PUBLIC_KEY_PATH}"
    chmod 0644 "${APISERVER_PUBLIC_KEY_PATH}"
    chown root:root "${APISERVER_PUBLIC_KEY_PATH}"

    AZURE_JSON_PATH="/etc/kubernetes/azure.json"
    touch "${AZURE_JSON_PATH}"
    chmod 0600 "${AZURE_JSON_PATH}"
    chown root:root "${AZURE_JSON_PATH}"

    set +x
    echo "${KUBELET_PRIVATE_KEY}" | base64 --decode > "${KUBELET_PRIVATE_KEY_PATH}"
    echo "${APISERVER_PUBLIC_KEY}" | base64 --decode > "${APISERVER_PUBLIC_KEY_PATH}"
    # Perform the required JSON escaping for special characters " and \
    SERVICE_PRINCIPAL_CLIENT_SECRET=$(echo $SERVICE_PRINCIPAL_CLIENT_SECRET | sed "s|\\\\|\\\\\\\|g")
    SERVICE_PRINCIPAL_CLIENT_SECRET=$(echo $SERVICE_PRINCIPAL_CLIENT_SECRET | sed 's|"|\\"|g')
    cat << EOF > "${AZURE_JSON_PATH}"
{
    "cloud":"${TARGET_ENVIRONMENT}",
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
    "cloudProviderBackoff": ${CLOUDPROVIDER_BACKOFF},
    "cloudProviderBackoffRetries": ${CLOUDPROVIDER_BACKOFF_RETRIES},
    "cloudProviderBackoffExponent": ${CLOUDPROVIDER_BACKOFF_EXPONENT},
    "cloudProviderBackoffDuration": ${CLOUDPROVIDER_BACKOFF_DURATION},
    "cloudProviderBackoffJitter": ${CLOUDPROVIDER_BACKOFF_JITTER},
    "cloudProviderRatelimit": ${CLOUDPROVIDER_RATELIMIT},
    "cloudProviderRateLimitQPS": ${CLOUDPROVIDER_RATELIMIT_QPS},
    "cloudProviderRateLimitBucket": ${CLOUDPROVIDER_RATELIMIT_BUCKET},
    "useManagedIdentityExtension": ${USE_MANAGED_IDENTITY_EXTENSION},
    "userAssignedIdentityID": "${USER_ASSIGNED_IDENTITY_ID}",
    "useInstanceMetadata": ${USE_INSTANCE_METADATA},
    "loadBalancerSku": "${LOAD_BALANCER_SKU}",
    "excludeMasterFromStandardLB": ${EXCLUDE_MASTER_FROM_STANDARD_LB},
    "providerVaultName": "${KMS_PROVIDER_VAULT_NAME}",
    "maximumLoadBalancerRuleCount": ${MAXIMUM_LOADBALANCER_RULE_COUNT},
    "providerKeyName": "k8s",
    "providerKeyVersion": ""
}
EOF
    set -x
    if [[ ! -z "${MASTER_NODE}" ]]; then
        if [[ "${ENABLE_AGGREGATED_APIS}" = True ]]; then
            generateAggregatedAPICerts
        fi
    fi
}

configureCNI() {
    # needed for the iptables rules to work on bridges
    retrycmd_if_failure 120 5 25 modprobe br_netfilter || exit $ERR_MODPROBE_FAIL
    echo -n "br_netfilter" > /etc/modules-load.d/br_netfilter.conf
    if [[ "${NETWORK_PLUGIN}" = "azure" ]]; then
        if [[ "${NETWORK_POLICY}" != "calico" ]]; then
            mv $CNI_BIN_DIR/10-azure.conflist $CNI_CONFIG_DIR/
            chmod 600 $CNI_CONFIG_DIR/10-azure.conflist
        fi
        /sbin/ebtables -t nat --list
    fi
    if [[ "${NETWORK_PLUGIN}" = "cilium" ]]; then
        systemctl enable sys-fs-bpf.mount
        systemctl restart sys-fs-bpf.mount
        REBOOTREQUIRED=true
    fi
}

setKubeletOpts () {
    KUBELET_DEFAULT_FILE=/etc/default/kubelet
    wait_for_file 1200 1 $KUBELET_DEFAULT_FILE || exit $ERR_FILE_WATCH_TIMEOUT
    sed -i "s#^KUBELET_OPTS=.*#KUBELET_OPTS=${1}#" $KUBELET_DEFAULT_FILE
}

ensureCCProxy() {
    cat $CC_SERVICE_IN_TMP | sed 's#@libexecdir@#/usr/libexec#' > /etc/systemd/system/cc-proxy.service
    cat $CC_SOCKET_IN_TMP sed 's#@localstatedir@#/var#' > /etc/systemd/system/cc-proxy.socket
	echo "Enabling and starting Clear Containers proxy service..."
	systemctlEnableAndStart cc-proxy || exit $ERR_SYSTEMCTL_START_FAIL
}

setupContainerd() {
    echo "Configuring cri-containerd..."
    mkdir -p "/etc/containerd"
    CRI_CONTAINERD_CONFIG="/etc/containerd/config.toml"
    echo "subreaper = false" > "$CRI_CONTAINERD_CONFIG"
    echo "oom_score = 0" >> "$CRI_CONTAINERD_CONFIG"
    echo "[plugins.cri]" >> "$CRI_CONTAINERD_CONFIG"
    echo "sandbox_image = \"$POD_INFRA_CONTAINER_SPEC\"" >> "$CRI_CONTAINERD_CONFIG"
    echo "[plugins.cri.containerd.untrusted_workload_runtime]" >> "$CRI_CONTAINERD_CONFIG"
    echo "runtime_type = 'io.containerd.runtime.v1.linux'" >> "$CRI_CONTAINERD_CONFIG"
    if [[ "$CONTAINER_RUNTIME" == "clear-containers" ]]; then
        echo "runtime_engine = '/usr/bin/cc-runtime'" >> "$CRI_CONTAINERD_CONFIG"
    elif [[ "$CONTAINER_RUNTIME" == "kata-containers" ]]; then
        echo "runtime_engine = '/usr/bin/kata-runtime'" >> "$CRI_CONTAINERD_CONFIG"
    else
        echo "runtime_engine = '/usr/local/sbin/runc'" >> "$CRI_CONTAINERD_CONFIG"
    fi
    echo "[plugins.cri.containerd.default_runtime]" >> "$CRI_CONTAINERD_CONFIG"
    echo "runtime_type = 'io.containerd.runtime.v1.linux'" >> "$CRI_CONTAINERD_CONFIG"
    echo "runtime_engine = '/usr/local/sbin/runc'" >> "$CRI_CONTAINERD_CONFIG"
    setKubeletOpts " --container-runtime=remote --runtime-request-timeout=15m --container-runtime-endpoint=unix:///run/containerd/containerd.sock"
}

ensureContainerd() {
    setupContainerd
    echo "Enabling and starting cri-containerd service..."
    systemctlEnableAndStart containerd || exit $ERR_SYSTEMCTL_START_FAIL
}

ensureDocker() {
    DOCKER_SERVICE_EXEC_START_FILE=/etc/systemd/system/docker.service.d/exec_start.conf
    wait_for_file 1200 1 $DOCKER_SERVICE_EXEC_START_FILE || exit $ERR_FILE_WATCH_TIMEOUT
    echo "ExecStartPost=/sbin/iptables -P FORWARD ACCEPT" >> $DOCKER_SERVICE_EXEC_START_FILE
    usermod -aG docker ${ADMINUSER}
    DOCKER_MOUNT_FLAGS_SYSTEMD_FILE=/etc/systemd/system/docker.service.d/clear_mount_propagation_flags.conf
    wait_for_file 1200 1 $DOCKER_MOUNT_FLAGS_SYSTEMD_FILE || exit $ERR_FILE_WATCH_TIMEOUT
    DOCKER_JSON_FILE=/etc/docker/daemon.json
    wait_for_file 1200 1 $DOCKER_JSON_FILE || exit $ERR_FILE_WATCH_TIMEOUT
    systemctlEnableAndStart docker
    # Delay start of docker-monitor for 30 mins after booting
    DOCKER_MONITOR_SYSTEMD_TIMER_FILE=/etc/systemd/system/docker-monitor.timer
    wait_for_file 1200 1 $DOCKER_MONITOR_SYSTEMD_TIMER_FILE || exit $ERR_FILE_WATCH_TIMEOUT
    DOCKER_MONITOR_SYSTEMD_FILE=/etc/systemd/system/docker-monitor.service
    wait_for_file 1200 1 $DOCKER_MONITOR_SYSTEMD_FILE || exit $ERR_FILE_WATCH_TIMEOUT
    systemctlEnableAndStart docker-monitor.timer || exit $ERR_SYSTEMCTL_START_FAIL
}

ensureKMS() {
    systemctlEnableAndStart kms || exit $ERR_SYSTEMCTL_START_FAIL
}

ensureKubelet() {
    KUBELET_DEFAULT_FILE=/etc/default/kubelet
    wait_for_file 1200 1 $KUBELET_DEFAULT_FILE || exit $ERR_FILE_WATCH_TIMEOUT
    KUBECONFIG_FILE=/var/lib/kubelet/kubeconfig
    wait_for_file 1200 1 $KUBECONFIG_FILE || exit $ERR_FILE_WATCH_TIMEOUT
    KUBELET_RUNTIME_CONFIG_SCRIPT_FILE=/opt/azure/containers/kubelet.sh
    wait_for_file 1200 1 $KUBELET_RUNTIME_CONFIG_SCRIPT_FILE || exit $ERR_FILE_WATCH_TIMEOUT
    systemctlEnableAndStart kubelet || exit $ERR_KUBELET_START_FAIL
    # Delay start of kubelet-monitor for 30 mins after booting
    #KUBELET_MONITOR_SYSTEMD_TIMER_FILE=/etc/systemd/system/kubelet-monitor.timer
    #wait_for_file 1200 1 $KUBELET_MONITOR_SYSTEMD_TIMER_FILE || exit $ERR_FILE_WATCH_TIMEOUT
    #KUBELET_MONITOR_SYSTEMD_FILE=/etc/systemd/system/kubelet-monitor.service
    #wait_for_file 1200 1 $KUBELET_MONITOR_SYSTEMD_FILE || exit $ERR_FILE_WATCH_TIMEOUT
    #systemctlEnableAndStart kubelet-monitor.timer || exit $ERR_SYSTEMCTL_START_FAIL
}

ensureJournal(){
    echo "Storage=persistent" >> /etc/systemd/journald.conf
    echo "SystemMaxUse=1G" >> /etc/systemd/journald.conf
    echo "RuntimeMaxUse=1G" >> /etc/systemd/journald.conf
    echo "ForwardToSyslog=no" >> /etc/systemd/journald.conf
    systemctlEnableAndStart systemd-journald || exit $ERR_SYSTEMCTL_START_FAIL
}

ensurePodSecurityPolicy() {
    POD_SECURITY_POLICY_FILE="/etc/kubernetes/manifests/pod-security-policy.yaml"
    if [ -f $POD_SECURITY_POLICY_FILE ]; then
        $KUBECTL create -f $POD_SECURITY_POLICY_FILE
    fi
}

ensureK8sControlPlane() {
    if $REBOOTREQUIRED || [ "$NO_OUTBOUND" = "true" ]; then
        return
    fi
    wait_for_file 3600 1 $KUBECTL || exit $ERR_FILE_WATCH_TIMEOUT
    # workaround for 1.12 bug https://github.com/Azure/aks-engine/issues/3681
    if [[ "${KUBERNETES_VERSION}" = 1.12.* ]]; then
        ensureKubelet
        retrycmd_if_failure 120 5 25 $KUBECTL 2>/dev/null cluster-info || ensureKubelet && retrycmd_if_failure 900 1 20 $KUBECTL 2>/dev/null cluster-info || exit $ERR_K8S_RUNNING_TIMEOUT
    else
        retrycmd_if_failure 120 5 25 $KUBECTL 2>/dev/null cluster-info || exit $ERR_K8S_RUNNING_TIMEOUT
    fi
}

ensureEtcd() {
    retrycmd_if_failure 120 5 25 curl --cacert /etc/kubernetes/certs/ca.crt --cert /etc/kubernetes/certs/etcdclient.crt --key /etc/kubernetes/certs/etcdclient.key ${ETCD_CLIENT_URL}/v2/machines || exit $ERR_ETCD_RUNNING_TIMEOUT
}

createKubeManifestDir() {
    KUBEMANIFESTDIR=/etc/kubernetes/manifests
    mkdir -p $KUBEMANIFESTDIR
}

writeKubeConfig() {
    KUBECONFIGDIR=/home/$ADMINUSER/.kube
    KUBECONFIGFILE=$KUBECONFIGDIR/config
    mkdir -p $KUBECONFIGDIR
    touch $KUBECONFIGFILE
    chown $ADMINUSER:$ADMINUSER $KUBECONFIGDIR
    chown $ADMINUSER:$ADMINUSER $KUBECONFIGFILE
    chmod 700 $KUBECONFIGDIR
    chmod 600 $KUBECONFIGFILE
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
" > $KUBECONFIGFILE
    set -x
}

configClusterAutoscalerAddon() {
    CLUSTER_AUTOSCALER_ADDON_FILE=/etc/kubernetes/addons/cluster-autoscaler-deployment.yaml
    wait_for_file 1200 1 $CLUSTER_AUTOSCALER_ADDON_FILE || exit $ERR_FILE_WATCH_TIMEOUT
    if [[ "${USE_MANAGED_IDENTITY_EXTENSION}" == true ]]; then
        CLUSTER_AUTOSCALER_MSI_VOLUME_MOUNT="- mountPath: /var/lib/waagent/\n\          name: waagent\n\          readOnly: true"
        CLUSTER_AUTOSCALER_MSI_VOLUME="- hostPath:\n\          path: /var/lib/waagent/\n\        name: waagent"
        CLUSTER_AUTOSCALER_MSI_HOST_NETWORK="hostNetwork: true"

        sed -i "s|<volMounts>|${CLUSTER_AUTOSCALER_MSI_VOLUME_MOUNT}|g" $CLUSTER_AUTOSCALER_ADDON_FILE
        sed -i "s|<vols>|${CLUSTER_AUTOSCALER_MSI_VOLUME}|g" $CLUSTER_AUTOSCALER_ADDON_FILE
        sed -i "s|<hostNet>|$(echo "${CLUSTER_AUTOSCALER_MSI_HOST_NETWORK}")|g" $CLUSTER_AUTOSCALER_ADDON_FILE
    elif [[ "${USE_MANAGED_IDENTITY_EXTENSION}" == false ]]; then
        sed -i "s|<volMounts>|""|g" $CLUSTER_AUTOSCALER_ADDON_FILE
        sed -i "s|<vols>|""|g" $CLUSTER_AUTOSCALER_ADDON_FILE
        sed -i "s|<hostNet>|""|g" $CLUSTER_AUTOSCALER_ADDON_FILE
    fi

    sed -i "s|<clientID>|$(echo $SERVICE_PRINCIPAL_CLIENT_ID | base64)|g" $CLUSTER_AUTOSCALER_ADDON_FILE
    sed -i "s|<clientSec>|$(echo $SERVICE_PRINCIPAL_CLIENT_SECRET | base64)|g" $CLUSTER_AUTOSCALER_ADDON_FILE
    sed -i "s|<subID>|$(echo $SUBSCRIPTION_ID | base64)|g" $CLUSTER_AUTOSCALER_ADDON_FILE
    sed -i "s|<tenantID>|$(echo $TENANT_ID | base64)|g" $CLUSTER_AUTOSCALER_ADDON_FILE
    sed -i "s|<rg>|$(echo $RESOURCE_GROUP | base64)|g" $CLUSTER_AUTOSCALER_ADDON_FILE
    sed -i "s|<vmType>|$(echo $VM_TYPE | base64)|g" $CLUSTER_AUTOSCALER_ADDON_FILE
    sed -i "s|<vmssName>|$(echo $PRIMARY_SCALE_SET)|g" $CLUSTER_AUTOSCALER_ADDON_FILE
}

configACIConnectorAddon() {
    ACI_CONNECTOR_CREDENTIALS=$(printf "{\"clientId\": \"$(echo $SERVICE_PRINCIPAL_CLIENT_ID)\", \"clientSecret\": \"$(echo $SERVICE_PRINCIPAL_CLIENT_SECRET)\", \"tenantId\": \"$(echo $TENANT_ID)\", \"subscriptionId\": \"$(echo $SUBSCRIPTION_ID)\", \"activeDirectoryEndpointUrl\": \"https://login.microsoftonline.com\",\"resourceManagerEndpointUrl\": \"https://management.azure.com/\", \"activeDirectoryGraphResourceId\": \"https://graph.windows.net/\", \"sqlManagementEndpointUrl\": \"https://management.core.windows.net:8443/\", \"galleryEndpointUrl\": \"https://gallery.azure.com/\", \"managementEndpointUrl\": \"https://management.core.windows.net/\"}" | base64 -w 0)

    openssl req -newkey rsa:4096 -new -nodes -x509 -days 3650 -keyout /etc/kubernetes/certs/aci-connector-key.pem -out /etc/kubernetes/certs/aci-connector-cert.pem -subj "/C=US/ST=CA/L=virtualkubelet/O=virtualkubelet/OU=virtualkubelet/CN=virtualkubelet"
    ACI_CONNECTOR_KEY=$(base64 /etc/kubernetes/certs/aci-connector-key.pem -w0)
    ACI_CONNECTOR_CERT=$(base64 /etc/kubernetes/certs/aci-connector-cert.pem -w0)

    ACI_CONNECTOR_ADDON_FILE=/etc/kubernetes/addons/aci-connector-deployment.yaml
    wait_for_file 1200 1 $ACI_CONNECTOR_ADDON_FILE || exit $ERR_FILE_WATCH_TIMEOUT
    sed -i "s|<creds>|$ACI_CONNECTOR_CREDENTIALS|g" $ACI_CONNECTOR_ADDON_FILE
    sed -i "s|<rgName>|$(echo $RESOURCE_GROUP)|g" $ACI_CONNECTOR_ADDON_FILE
    sed -i "s|<cert>|$(echo $ACI_CONNECTOR_CERT)|g" $ACI_CONNECTOR_ADDON_FILE
    sed -i "s|<key>|$(echo $ACI_CONNECTOR_KEY)|g" $ACI_CONNECTOR_ADDON_FILE
}

configAddons() {
    if [[ "${CLUSTER_AUTOSCALER_ADDON}" = True ]]; then
        configClusterAutoscalerAddon
    fi

    if [[ "${ACI_CONNECTOR_ADDON}" = True ]]; then
        configACIConnectorAddon
    fi
}

configGPUDrivers() {
    retrycmd_if_failure 30 1 120 sh $GPU_DEST/nvidia-drivers-$GPU_DV --silent --accept-license --no-drm --dkms --utility-prefix="${GPU_DEST}" --opengl-prefix="${GPU_DEST}" || exit $ERR_GPU_DRIVERS_START_FAIL
    echo "${GPU_DEST}/lib64" > /etc/ld.so.conf.d/nvidia.conf
    retrycmd_if_failure 120 5 25 ldconfig || exit $ERR_GPU_DRIVERS_START_FAIL
    umount -l /usr/lib/x86_64-linux-gnu
    retrycmd_if_failure 120 5 25 nvidia-modprobe -u -c0 || exit $ERR_GPU_DRIVERS_START_FAIL
    retrycmd_if_failure 120 5 25 $GPU_DEST/bin/nvidia-smi || exit $ERR_GPU_DRIVERS_START_FAIL
    retrycmd_if_failure 120 5 25 ldconfig || exit $ERR_GPU_DRIVERS_START_FAIL
}

ensureGPUDrivers() {
    configGPUDrivers
    systemctlEnableAndStart nvidia-modprobe || exit $ERR_GPU_DRIVERS_START_FAIL
    retrycmd_if_failure 30 1 120 systemctl restart kubelet || exit $ERR_GPU_DRIVERS_START_FAIL
}
