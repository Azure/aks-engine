#!/bin/bash

CC_SERVICE_IN_TMP=/opt/azure/containers/cc-proxy.service.in
CC_SOCKET_IN_TMP=/opt/azure/containers/cc-proxy.socket.in
CNI_CONFIG_DIR="/etc/cni/net.d"
CNI_BIN_DIR="/opt/cni/bin"
CNI_DOWNLOADS_DIR="/opt/cni/downloads"
CONTAINERD_DOWNLOADS_DIR="/opt/containerd/downloads"

removeEtcd() {
    rm -rf /usr/bin/etcd
}

removeMoby() {
    sudo apt-get purge -y moby-engine moby-cli
}

installEtcd() {
    CURRENT_VERSION=$(etcd --version | grep "etcd Version" | cut -d ":" -f 2 | tr -d '[:space:]')
    if [[ "$CURRENT_VERSION" == "${ETCD_VERSION}" ]]; then
        echo "etcd version ${ETCD_VERSION} is already installed, skipping download"
    else
        retrycmd_get_tarball 120 5 /tmp/etcd-v${ETCD_VERSION}-linux-amd64.tar.gz ${ETCD_DOWNLOAD_URL}/etcd-v${ETCD_VERSION}-linux-amd64.tar.gz || exit $ERR_ETCD_DOWNLOAD_TIMEOUT
        removeEtcd
        tar -xzvf /tmp/etcd-v${ETCD_VERSION}-linux-amd64.tar.gz -C /usr/bin/ --strip-components=1 || exit $ERR_ETCD_DOWNLOAD_TIMEOUT
    fi
}

installDeps() {
    retrycmd_if_failure_no_stats 120 5 25 curl -fsSL https://packages.microsoft.com/config/ubuntu/16.04/packages-microsoft-prod.deb > /tmp/packages-microsoft-prod.deb || exit $ERR_MS_PROD_DEB_DOWNLOAD_TIMEOUT
    retrycmd_if_failure 60 5 10 dpkg -i /tmp/packages-microsoft-prod.deb || exit $ERR_MS_PROD_DEB_PKG_ADD_FAIL
    apt_get_update || exit $ERR_APT_UPDATE_TIMEOUT
    apt_get_install 30 1 600 apt-transport-https blobfuse ca-certificates ceph-common cgroup-lite cifs-utils conntrack ebtables ethtool fuse git glusterfs-client init-system-helpers iproute2 ipset iptables jq mount nfs-common pigz socat util-linux xz-utils zip htop iotop iftop sysstat || exit $ERR_APT_INSTALL_TIMEOUT
    retrycmd_if_failure_no_stats 120 5 25 echo "deb http://azure.archive.ubuntu.com/ubuntu/ xenial-proposed restricted main multiverse universe" >> /etc/apt/sources.list || exit $ERR_SYSTEMD_INSTALL_FAIL
    cat /etc/apt/sources.list
    apt_get_update || exit $ERR_APT_UPDATE_TIMEOUT
    apt_get_install 20 30 120 systemd/xenial-proposed || exit $ERR_SYSTEMD_INSTALL_FAIL
}

installGPUDrivers() {
    rmmod nouveau
    echo blacklist nouveau >> /etc/modprobe.d/blacklist.conf
    retrycmd_if_failure_no_stats 120 5 25 update-initramfs -u || exit $ERR_GPU_DRIVERS_INSTALL_TIMEOUT
    mkdir -p $GPU_DEST
    retrycmd_if_failure_no_stats 120 5 25 curl -fsSL https://nvidia.github.io/nvidia-docker/gpgkey > /tmp/aptnvidia.gpg || exit $ERR_GPU_DRIVERS_INSTALL_TIMEOUT
    wait_for_apt_locks
    retrycmd_if_failure 120 5 25 apt-key add /tmp/aptnvidia.gpg || exit $ERR_GPU_DRIVERS_INSTALL_TIMEOUT
    wait_for_apt_locks
    retrycmd_if_failure_no_stats 120 5 25 curl -fsSL https://nvidia.github.io/nvidia-docker/ubuntu16.04/amd64/nvidia-docker.list > /tmp/nvidia-docker.list || exit  $ERR_GPU_DRIVERS_INSTALL_TIMEOUT
    wait_for_apt_locks
    retrycmd_if_failure_no_stats 120 5 25 cat /tmp/nvidia-docker.list > /etc/apt/sources.list.d/nvidia-docker.list
    apt_get_update
    retrycmd_if_failure 30 5 3600 apt-get install -y linux-headers-$(uname -r) gcc make dkms || exit $ERR_GPU_DRIVERS_INSTALL_TIMEOUT
    retrycmd_if_failure 30 5 3600 apt-get -o Dpkg::Options::="--force-confold" install -y nvidia-docker2=${NVIDIA_DOCKER_VERSION}+docker${DOCKER_VERSION} nvidia-container-runtime=${NVIDIA_CONTAINER_RUNTIME_VERSION}+docker${DOCKER_VERSION} || exit $ERR_GPU_DRIVERS_INSTALL_TIMEOUT
    retrycmd_if_failure 120 5 25 pkill -SIGHUP dockerd || exit $ERR_GPU_DRIVERS_INSTALL_TIMEOUT
    retrycmd_if_failure 30 5 60 curl -fLS https://us.download.nvidia.com/tesla/$GPU_DV/NVIDIA-Linux-x86_64-${GPU_DV}.run -o ${GPU_DEST}/nvidia-drivers-${GPU_DV} || exit $ERR_GPU_DRIVERS_INSTALL_TIMEOUT
    mkdir -p $GPU_DEST/lib64 $GPU_DEST/overlay-workdir
    retrycmd_if_failure 120 5 25 mount -t overlay -o lowerdir=/usr/lib/x86_64-linux-gnu,upperdir=${GPU_DEST}/lib64,workdir=${GPU_DEST}/overlay-workdir none /usr/lib/x86_64-linux-gnu || exit $ERR_GPU_DRIVERS_INSTALL_TIMEOUT
}

installSGXDrivers() {
    echo "Installing SGX driver"
    local VERSION=`grep DISTRIB_RELEASE /etc/*-release| cut -f 2 -d "="`
    case $VERSION in
    "18.04")
        SGX_DRIVER_URL="https://download.01.org/intel-sgx/dcap-1.0.1/dcap_installer/ubuntuServer1804/sgx_linux_x64_driver_dcap_4f32b98.bin"
        ;;
    "16.04")
        SGX_DRIVER_URL="https://download.01.org/intel-sgx/dcap-1.0.1/dcap_installer/ubuntuServer1604/sgx_linux_x64_driver_dcap_4f32b98.bin"
        ;;
    "*")
        echo "Version $VERSION is not supported"
        exit 1
        ;;
    esac

    local PACKAGES="make gcc dkms"
    wait_for_apt_locks
    retrycmd_if_failure 30 5 3600 apt-get -y install $PACKAGES  || exit $ERR_SGX_DRIVERS_INSTALL_TIMEOUT

    local SGX_DRIVER=$(basename $SGX_DRIVER_URL)
    local OE_DIR=/opt/azure/containers/oe
    mkdir -p ${OE_DIR}

    retrycmd_if_failure 120 5 25 curl -fsSL ${SGX_DRIVER_URL} -o ${OE_DIR}/${SGX_DRIVER} || exit $ERR_SGX_DRIVERS_INSTALL_TIMEOUT
    chmod a+x ${OE_DIR}/${SGX_DRIVER}
    ${OE_DIR}/${SGX_DRIVER} || exit $ERR_SGX_DRIVERS_START_FAIL
}

installContainerRuntime() {
    if [[ "$CONTAINER_RUNTIME" == "docker" ]]; then
        if [[ "$DOCKER_ENGINE_REPO" != "" ]]; then
            installDockerEngine
        else
            installMoby
        fi
    elif [[ "$CONTAINER_RUNTIME" == "clear-containers" ]]; then
	    # Ensure we can nest virtualization
        if grep -q vmx /proc/cpuinfo; then
            installClearContainersRuntime
        fi
    fi
}

installMoby() {
    CURRENT_VERSION=$(dockerd --version | grep "Docker version" | cut -d "," -f 1 | cut -d " " -f 3)
    if [[ "$CURRENT_VERSION" == "${MOBY_VERSION}" ]]; then
        echo "dockerd $MOBY_VERSION is already installed, skipping Moby download"
    else
        removeMoby
        retrycmd_if_failure_no_stats 120 5 25 curl https://packages.microsoft.com/config/ubuntu/16.04/prod.list > /tmp/microsoft-prod.list || exit $ERR_MOBY_APT_LIST_TIMEOUT
        retrycmd_if_failure 10 5 10 cp /tmp/microsoft-prod.list /etc/apt/sources.list.d/ || exit $ERR_MOBY_APT_LIST_TIMEOUT
        retrycmd_if_failure_no_stats 120 5 25 curl https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > /tmp/microsoft.gpg || exit $ERR_MS_GPG_KEY_DOWNLOAD_TIMEOUT
        retrycmd_if_failure 10 5 10 cp /tmp/microsoft.gpg /etc/apt/trusted.gpg.d/ || exit $ERR_MS_GPG_KEY_DOWNLOAD_TIMEOUT
        apt_get_update || exit $ERR_APT_UPDATE_TIMEOUT
        apt_get_install 20 30 120 moby-engine=${MOBY_VERSION} moby-cli=3.0.3 --allow-downgrades || exit $ERR_MOBY_INSTALL_TIMEOUT  # HACK: revert moby-cli to ${MOBY_VERSION} for next release
    fi
}

installDockerEngine() {
    DOCKER_ENGINE_VERSION="1.13.*"
    dockerd --version
    if [ $? -eq 0 ]; then
        echo "dockerd is already installed, skipping download"
    else
        retrycmd_if_failure_no_stats 20 1 5 curl -fsSL https://aptdocker.azureedge.net/gpg > /tmp/aptdocker.gpg || exit $ERR_DOCKER_KEY_DOWNLOAD_TIMEOUT
        retrycmd_if_failure 10 5 10 apt-key add /tmp/aptdocker.gpg || exit $ERR_DOCKER_APT_KEY_TIMEOUT
        echo "deb ${DOCKER_ENGINE_REPO} ubuntu-xenial main" | sudo tee /etc/apt/sources.list.d/docker.list
        printf "Package: docker-engine\nPin: version ${DOCKER_ENGINE_VERSION}\nPin-Priority: 550\n" > /etc/apt/preferences.d/docker.pref
        apt_get_update || exit $ERR_APT_UPDATE_TIMEOUT
        apt_get_install 20 30 120 docker-engine || exit $ERR_DOCKER_INSTALL_TIMEOUT
    fi
}

overrideDockerEngineStorageDriver() {
    echo "stopping docker engine"
    sudo systemctl stop docker
    cat << EOF > /tmp/daemon.json 
{
    "storage-driver": "overlay2",
    "live-restore": true,
    "log-driver": "json-file",
    "log-opts":  {
        "max-size": "50m",
        "max-file": "5"
    }
}
EOF
    sudo mv /tmp/daemon.json /etc/docker/daemon.json
    echo "cleaning up aufs storage"
    sudo rm -rf /var/lib/docker/aufs
    echo "starting docker engine with overlay2 driver"
    sudo systemctl start docker
}

installKataContainersRuntime() {
    # TODO incorporate this into packer CI so that it is pre-baked into the VHD image
    echo "Adding Kata Containers repository key..."
    KATA_RELEASE_KEY_TMP=/tmp/kata-containers-release.key
    KATA_URL=http://download.opensuse.org/repositories/home:/katacontainers:/release/xUbuntu_16.04/Release.key
    retrycmd_if_failure_no_stats 120 5 25 curl -fsSL $KATA_URL > $KATA_RELEASE_KEY_TMP || exit $ERR_KATA_KEY_DOWNLOAD_TIMEOUT
    wait_for_apt_locks
    retrycmd_if_failure 30 5 30 apt-key add $KATA_RELEASE_KEY_TMP || exit $ERR_KATA_APT_KEY_TIMEOUT
    echo "Adding Kata Containers repository..."
    echo 'deb http://download.opensuse.org/repositories/home:/katacontainers:/release/xUbuntu_16.04/ /' > /etc/apt/sources.list.d/kata-containers.list
    echo "Installing Kata Containers runtime..."
    apt_get_update || exit $ERR_APT_UPDATE_TIMEOUT
    apt_get_install 120 5 25 kata-runtime || exit $ERR_KATA_INSTALL_TIMEOUT
}

installClearContainersRuntime() {
    cc-runtime --version
    if [ $? -eq 0 ]; then
        echo "cc-runtime is already installed, skipping download"
    else
        echo "Adding Clear Containers repository key..."
        CC_RELEASE_KEY_TMP=/tmp/clear-containers-release.key
        CC_URL=https://download.opensuse.org/repositories/home:clearcontainers:clear-containers-3/xUbuntu_16.04/Release.key
        retrycmd_if_failure_no_stats 120 5 25 curl -fsSL $CC_URL > $CC_RELEASE_KEY_TMP || exit $ERR_APT_INSTALL_TIMEOUT
        wait_for_apt_locks
        retrycmd_if_failure 120 5 25 apt-key add $CC_RELEASE_KEY_TMP || exit $ERR_APT_INSTALL_TIMEOUT
        echo "Adding Clear Containers repository..."
        echo 'deb http://download.opensuse.org/repositories/home:/clearcontainers:/clear-containers-3/xUbuntu_16.04/ /' > /etc/apt/sources.list.d/cc-runtime.list
        echo "Installing Clear Containers runtime..."
        apt_get_update || exit $ERR_APT_UPDATE_TIMEOUT
        apt_get_install 120 5 25 cc-runtime
        local repo_uri="https://raw.githubusercontent.com/clearcontainers/proxy/3.0.23"
        retrycmd_if_failure_no_stats 120 5 25 curl -fsSL "${repo_uri}/cc-proxy.service.in" > $CC_SERVICE_IN_TMP
        retrycmd_if_failure_no_stats 120 5 25 curl -fsSL "${repo_uri}/cc-proxy.socket.in" > $CC_SOCKET_IN_TMP
    fi
}

installNetworkPlugin() {
    if [[ "${NETWORK_PLUGIN}" = "azure" ]]; then
        installAzureCNI
    fi
    installCNI
    rm -rf $CNI_DOWNLOADS_DIR &
}

downloadCNI() {
    mkdir -p $CNI_DOWNLOADS_DIR
    CNI_TGZ_TMP=$(echo ${CNI_PLUGINS_URL} | cut -d "/" -f 5)
    retrycmd_get_tarball 120 5 "$CNI_DOWNLOADS_DIR/${CNI_TGZ_TMP}" ${CNI_PLUGINS_URL} || exit $ERR_CNI_DOWNLOAD_TIMEOUT
}

downloadAzureCNI() {
    mkdir -p $CNI_DOWNLOADS_DIR
    CNI_TGZ_TMP=$(echo ${VNET_CNI_PLUGINS_URL} | cut -d "/" -f 5)
    retrycmd_get_tarball 120 5 "$CNI_DOWNLOADS_DIR/${CNI_TGZ_TMP}" ${VNET_CNI_PLUGINS_URL} || exit $ERR_CNI_DOWNLOAD_TIMEOUT
}

downloadContainerd() {
    CONTAINERD_DOWNLOAD_URL="${CONTAINERD_DOWNLOAD_URL_BASE}cri-containerd-${CONTAINERD_VERSION}.linux-amd64.tar.gz"
    mkdir -p $CONTAINERD_DOWNLOADS_DIR
    CONTAINERD_TGZ_TMP=$(echo ${CONTAINERD_DOWNLOAD_URL} | cut -d "/" -f 5)
    retrycmd_get_tarball 120 5 "$CONTAINERD_DOWNLOADS_DIR/${CONTAINERD_TGZ_TMP}" ${CONTAINERD_DOWNLOAD_URL} || exit $ERR_CONTAINERD_DOWNLOAD_TIMEOUT
}

installCNI() {
    CNI_TGZ_TMP=$(echo ${CNI_PLUGINS_URL} | cut -d "/" -f 5)
    if [[ ! -f "$CNI_DOWNLOADS_DIR/${CNI_TGZ_TMP}" ]]; then
        downloadCNI
    fi
    mkdir -p $CNI_BIN_DIR
    tar -xzf "$CNI_DOWNLOADS_DIR/${CNI_TGZ_TMP}" -C $CNI_BIN_DIR
    chown -R root:root $CNI_BIN_DIR
    chmod -R 755 $CNI_BIN_DIR
}

installAzureCNI() {
    CNI_TGZ_TMP=$(echo ${VNET_CNI_PLUGINS_URL} | cut -d "/" -f 5)
    if [[ ! -f "$CNI_DOWNLOADS_DIR/${CNI_TGZ_TMP}" ]]; then
        downloadAzureCNI
    fi
    mkdir -p $CNI_CONFIG_DIR
    chown -R root:root $CNI_CONFIG_DIR
    chmod 755 $CNI_CONFIG_DIR
    mkdir -p $CNI_BIN_DIR
    tar -xzf "$CNI_DOWNLOADS_DIR/${CNI_TGZ_TMP}" -C $CNI_BIN_DIR
}

installContainerd() {
    CURRENT_VERSION=$(containerd -version | cut -d " " -f 3 | sed 's|v||')
    if [[ "$CURRENT_VERSION" == "${CONTAINERD_VERSION}" ]]; then
        echo "containerd is already installed, skipping install"
    else
        CONTAINERD_TGZ_TMP="cri-containerd-${CONTAINERD_VERSION}.linux-amd64.tar.gz"
        rm -Rf /usr/bin/containerd
        rm -Rf /var/lib/docker/containerd
        rm -Rf /run/docker/containerd
        if [[ ! -f "$CONTAINERD_DOWNLOADS_DIR/${CONTAINERD_TGZ_TMP}" ]]; then
            downloadContainerd
        fi
        tar -xzf "$CONTAINERD_DOWNLOADS_DIR/$CONTAINERD_TGZ_TMP" -C /
        sed -i '/\[Service\]/a ExecStartPost=\/sbin\/iptables -P FORWARD ACCEPT -w' /etc/systemd/system/containerd.service
        echo "Successfully installed cri-containerd..."
    fi
    rm -Rf $CONTAINERD_DOWNLOADS_DIR &
}

installImg() {
    img_filepath=/usr/local/bin/img
    retrycmd_get_executable 120 5 $img_filepath "https://acs-mirror.azureedge.net/img/img-linux-amd64-v0.5.6" ls || exit $ERR_IMG_DOWNLOAD_TIMEOUT
}

extractHyperkube() {
    CLI_TOOL=$1
    path="/home/hyperkube-downloads/${KUBERNETES_VERSION}"
    pullContainerImage $CLI_TOOL ${HYPERKUBE_URL}
    if [[ "$CLI_TOOL" == "docker" ]]; then
        mkdir -p "$path"
        docker run --rm -v $path:$path ${HYPERKUBE_URL} /bin/bash -c "cp /hyperkube $path"
    else
        img unpack -o "$path" ${HYPERKUBE_URL}
    fi

    if [[ $OS == $COREOS_OS_NAME ]]; then
        cp "$path/hyperkube" "/opt/kubelet"
        mv "$path/hyperkube" "/opt/kubectl"
        chmod a+x /opt/kubelet /opt/kubectl
    else
        cp "$path/hyperkube" "/usr/local/bin/kubelet-${KUBERNETES_VERSION}"
        mv "$path/hyperkube" "/usr/local/bin/kubectl-${KUBERNETES_VERSION}"
    fi
}

installKubeletAndKubectl() {
    if [[ ! -f "/usr/local/bin/kubectl-${KUBERNETES_VERSION}" ]]; then
        if [[ "$CONTAINER_RUNTIME" == "docker" ]]; then
            extractHyperkube "docker"
        else
            installImg
            extractHyperkube "img"
        fi
    fi
    mv "/usr/local/bin/kubelet-${KUBERNETES_VERSION}" "/usr/local/bin/kubelet"
    mv "/usr/local/bin/kubectl-${KUBERNETES_VERSION}" "/usr/local/bin/kubectl"
    chmod a+x /usr/local/bin/kubelet /usr/local/bin/kubectl
    rm -rf /usr/local/bin/kubelet-* /usr/local/bin/kubectl-* /home/hyperkube-downloads &
}

pullContainerImage() {
    CLI_TOOL=$1
    DOCKER_IMAGE_URL=$2
    if [[ ! -z "${PRIVATE_AZURE_REGISTRY_SERVER:-}" ]]; then
        $CLI_TOOL login -u $SERVICE_PRINCIPAL_CLIENT_ID -p $SERVICE_PRINCIPAL_CLIENT_SECRET $PRIVATE_AZURE_REGISTRY_SERVER
    fi
    retrycmd_if_failure 60 1 1200 $CLI_TOOL pull $DOCKER_IMAGE_URL || exit $ERR_CONTAINER_IMG_PULL_TIMEOUT
}

cleanUpContainerImages() {
    # TODO remove all unused container images at runtime
    docker rmi $(docker images --format '{{.Repository}}:{{.Tag}}' | grep -v ${KUBERNETES_VERSION} | grep 'hyperkube') &
    docker rmi $(docker images --format '{{.Repository}}:{{.Tag}}' | grep -v ${KUBERNETES_VERSION} | grep 'cloud-controller-manager') &
    if [ "$IS_HOSTED_MASTER" = "false" ]; then
        echo "Cleaning up AKS container images, not an AKS cluster"
        docker rmi $(docker images --format '{{.Repository}}:{{.Tag}}' | grep 'hcp-tunnel-front') &
        docker rmi $(docker images --format '{{.Repository}}:{{.Tag}}' | grep 'kube-svc-redirect') &
        docker rmi $(docker images --format '{{.Repository}}:{{.Tag}}' | grep 'nginx') &
    fi
}
