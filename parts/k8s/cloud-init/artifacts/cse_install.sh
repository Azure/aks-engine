#!/bin/bash

CNI_CONFIG_DIR="/etc/cni/net.d"
CNI_BIN_DIR="/opt/cni/bin"
CNI_DOWNLOADS_DIR="/opt/cni/downloads"
CONTAINERD_DOWNLOADS_DIR="/opt/containerd/downloads"
K8S_DOWNLOADS_DIR="/opt/kubernetes/downloads"
APMZ_DOWNLOADS_DIR="/opt/apmz/downloads"
BPFTRACE_DOWNLOADS_DIR="/opt/bpftrace/downloads"
UBUNTU_RELEASE=$(lsb_release -r -s)
UBUNTU_CODENAME=$(lsb_release -c -s)

disableTimeSyncd() {
  systemctl_stop 20 5 10 systemd-timesyncd || exit 3
  retrycmd 120 5 25 systemctl disable systemd-timesyncd || exit 3
}
installEtcd() {
  CURRENT_VERSION=$(etcd --version | grep "etcd Version" | cut -d ":" -f 2 | tr -d '[:space:]')
  if [[ $CURRENT_VERSION != "${ETCD_VERSION}" ]]; then
    CLI_TOOL=$1
    local path="/usr/bin"
    CONTAINER_IMAGE=${ETCD_DOWNLOAD_URL}etcd:v${ETCD_VERSION}
    pullContainerImage $CLI_TOOL ${CONTAINER_IMAGE}
    removeEtcd
    if [[ $CLI_TOOL == "docker" ]]; then
      mkdir -p "$path"
      docker run --rm --entrypoint cat ${CONTAINER_IMAGE} /usr/local/bin/etcd >"$path/etcd"
      docker run --rm --entrypoint cat ${CONTAINER_IMAGE} /usr/local/bin/etcdctl >"$path/etcdctl"
    else
      tmpdir=/root/etcd${RANDOM}
      img unpack -o ${tmpdir} ${CONTAINER_IMAGE}
      mv ${tmpdir}/usr/local/bin/etcd ${tmpdir}/usr/local/bin/etcdctl ${path}
      rm -rf ${tmpdir}
    fi
    chmod a+x "$path/etcd" "$path/etcdctl"
  fi
}
installDeps() {
  packages="apache2-utils apt-transport-https blobfuse ca-certificates cifs-utils conntrack cracklib-runtime dbus dkms ebtables ethtool fuse gcc git htop iftop init-system-helpers iotop iproute2 ipset iptables jq libpam-pwquality libpwquality-tools linux-headers-$(uname -r) make mount nfs-common pigz socat sysstat traceroute util-linux xz-utils zip"
  if [[ ${OS} == "${UBUNTU_OS_NAME}" ]]; then
    retrycmd_no_stats 120 5 25 curl -fsSL https://packages.microsoft.com/config/ubuntu/${UBUNTU_RELEASE}/packages-microsoft-prod.deb >/tmp/packages-microsoft-prod.deb || exit 42
    retrycmd 60 5 10 dpkg -i /tmp/packages-microsoft-prod.deb || exit 43
    aptmarkWALinuxAgent hold
    packages+=" cgroup-lite ceph-common glusterfs-client"
    if [[ $UBUNTU_RELEASE == "18.04" ]]; then
      disableTimeSyncd
      packages+=" ntp ntpstat"
    fi
  elif [[ $OS == $DEBIAN_OS_NAME ]]; then
    packages+=" gpg cgroup-bin"
  fi

  apt_get_update || exit 99
  apt_get_dist_upgrade || exit 101

  for apt_package in ${packages}; do
    if ! apt_get_install 30 1 600 $apt_package; then
      journalctl --no-pager -u $apt_package
      exit 9
    fi
  done
  if [[ ${AUDITD_ENABLED} == true ]]; then
    if ! apt_get_install 30 1 600 auditd; then
      journalctl --no-pager -u auditd
      exit 9
    fi
  fi
}
downloadGPUDrivers() {
  mkdir -p $GPU_DEST/tmp
  retrycmd_no_stats 120 5 25 curl -fsSL https://nvidia.github.io/nvidia-docker/gpgkey >$GPU_DEST/tmp/aptnvidia.gpg || exit 85
  wait_for_apt_locks
  retrycmd 120 5 25 apt-key add $GPU_DEST/tmp/aptnvidia.gpg || exit 85
  wait_for_apt_locks
  retrycmd_no_stats 120 5 25 curl -fsSL https://nvidia.github.io/nvidia-docker/ubuntu${UBUNTU_RELEASE}/nvidia-docker.list >$GPU_DEST/tmp/nvidia-docker.list || exit 85
  wait_for_apt_locks
  retrycmd_no_stats 120 5 25 cat $GPU_DEST/tmp/nvidia-docker.list >/etc/apt/sources.list.d/nvidia-docker.list || exit 85
  apt_get_update
  retrycmd 30 5 60 curl -fLS https://us.download.nvidia.com/tesla/$GPU_DV/NVIDIA-Linux-x86_64-${GPU_DV}.run -o ${GPU_DEST}/nvidia-drivers-${GPU_DV} || exit 85
  tmpDir=$GPU_DEST/tmp
  if ! (
    set -e -o pipefail
    cd "${tmpDir}"
    retrycmd 30 5 3600 apt-get download nvidia-docker2="${NVIDIA_DOCKER_VERSION}+${NVIDIA_DOCKER_SUFFIX}" || exit 85
  ); then
    exit 85
  fi
}
installMoby() {
  removeContainerd
  CURRENT_VERSION=$(dockerd --version | grep "Docker version" | cut -d "," -f 1 | cut -d " " -f 3 | cut -d "+" -f 1)
  if [[ $CURRENT_VERSION != "${MOBY_VERSION}" ]]; then
    removeMoby
    retrycmd_no_stats 120 5 25 curl https://packages.microsoft.com/config/ubuntu/${UBUNTU_RELEASE}/prod.list >/tmp/microsoft-prod.list || exit 25
    retrycmd 10 5 10 cp /tmp/microsoft-prod.list /etc/apt/sources.list.d/ || exit 25
    retrycmd_no_stats 120 5 25 curl https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor >/tmp/microsoft.gpg || exit 26
    retrycmd 10 5 10 cp /tmp/microsoft.gpg /etc/apt/trusted.gpg.d/ || exit 26
    apt_get_update || exit 99
    MOBY_CLI=${MOBY_VERSION}
    if [[ ${MOBY_CLI} == "3.0.4" ]]; then
      MOBY_CLI="3.0.3"
    fi
    apt_get_install 20 30 120 moby-engine=${MOBY_VERSION}* moby-cli=${MOBY_CLI}* --allow-downgrades || exit 27
  fi
}
installBcc() {
  IOVISOR_KEY_TMP=/tmp/iovisor-release.key
  IOVISOR_URL=https://repo.iovisor.org/GPG-KEY
  retrycmd_no_stats 120 5 25 curl -fsSL $IOVISOR_URL >$IOVISOR_KEY_TMP || exit 166
  wait_for_apt_locks
  retrycmd 30 5 30 apt-key add $IOVISOR_KEY_TMP || exit 167
  echo "deb https://repo.iovisor.org/apt/${UBUNTU_CODENAME} ${UBUNTU_CODENAME} main" >/etc/apt/sources.list.d/iovisor.list
  apt_get_update || exit 99
  apt_get_install 120 5 25 bcc-tools libbcc-examples linux-headers-$(uname -r) || exit 168
}
downloadCNI() {
  mkdir -p $CNI_DOWNLOADS_DIR
  CNI_TGZ_TMP=${CNI_PLUGINS_URL##*/}
  retrycmd_get_tarball 120 5 "$CNI_DOWNLOADS_DIR/${CNI_TGZ_TMP}" ${CNI_PLUGINS_URL} || exit 41
}
downloadAzureCNI() {
  mkdir -p $CNI_DOWNLOADS_DIR
  CNI_TGZ_TMP=${VNET_CNI_PLUGINS_URL##*/}
  retrycmd_get_tarball 120 5 "$CNI_DOWNLOADS_DIR/${CNI_TGZ_TMP}" ${VNET_CNI_PLUGINS_URL} || exit 41
}
ensureAPMZ() {
  local version=$1
  local apmz_url="https://upstreamartifacts.azureedge.net/apmz/$version/binaries/apmz_linux_amd64.tar.gz" apmz_filepath="/usr/local/bin/apmz"
  if [[ -f $apmz_filepath ]]; then
    installed_version=$($apmz_filepath version)
    if [[ $version == "$installed_version" ]]; then
      return
    fi
  fi
  install_dir="$APMZ_DOWNLOADS_DIR/$version"
  download_path="$install_dir/apmz.gz"
  mkdir -p "$install_dir"
  retrycmd_get_tarball 120 5 "$download_path" "${apmz_url}"
  tar -xvf "$download_path" -C "$install_dir"
  bin_path="$install_dir/apmz_linux_amd64"
  chmod +x "$bin_path"
  ln -Ffs "$bin_path" "$apmz_filepath"
}
installBpftrace() {
  local version="v0.9.4"
  local bpftrace_bin="bpftrace"
  local bpftrace_tools="bpftrace-tools.tar"
  local bpftrace_url="https://upstreamartifacts.azureedge.net/$bpftrace_bin/$version"
  local bpftrace_filepath="/usr/local/bin/$bpftrace_bin"
  local tools_filepath="/usr/local/share/$bpftrace_bin"
  if [[ -f $bpftrace_filepath ]]; then
    installed_version="$($bpftrace_bin -V | cut -d' ' -f2)"
    if [[ $version == "$installed_version" ]]; then
      return
    fi
    rm "$bpftrace_filepath"
    if [[ -d $tools_filepath ]]; then
      rm -r "$tools_filepath"
    fi
  fi
  mkdir -p "$tools_filepath"
  install_dir="$BPFTRACE_DOWNLOADS_DIR/$version"
  mkdir -p "$install_dir"
  download_path="$install_dir/$bpftrace_tools"
  retrycmd 30 5 60 curl -fSL -o "$bpftrace_filepath" "$bpftrace_url/$bpftrace_bin" || exit 169
  retrycmd 30 5 60 curl -fSL -o "$download_path" "$bpftrace_url/$bpftrace_tools" || exit 170
  tar -xvf "$download_path" -C "$tools_filepath"
  chmod +x "$bpftrace_filepath"
  chmod -R +x "$tools_filepath/tools"
}
installImg() {
  img_filepath=/usr/local/bin/img
  retrycmd_get_executable 120 5 $img_filepath "https://upstreamartifacts.azureedge.net/img/img-linux-amd64-v0.5.6" ls || exit 33
}
extractHyperkube() {
  CLI_TOOL=$1
  path="/home/hyperkube-downloads/${KUBERNETES_VERSION}"
  pullContainerImage $CLI_TOOL ${HYPERKUBE_URL}
  if [[ $CLI_TOOL == "docker" ]]; then
    mkdir -p "$path"
    if docker run --rm --entrypoint "" -v $path:$path ${HYPERKUBE_URL} /bin/bash -c "cp /usr/local/bin/{kubelet,kubectl} $path"; then
      mv "$path/kubelet" "/usr/local/bin/kubelet-${KUBERNETES_VERSION}"
      mv "$path/kubectl" "/usr/local/bin/kubectl-${KUBERNETES_VERSION}"
      return
    else
      docker run --rm -v $path:$path ${HYPERKUBE_URL} /bin/bash -c "cp /hyperkube $path"
    fi
  else
    img unpack -o "$path" ${HYPERKUBE_URL}
  fi

  cp "$path/hyperkube" "/usr/local/bin/kubelet-${KUBERNETES_VERSION}"
  mv "$path/hyperkube" "/usr/local/bin/kubectl-${KUBERNETES_VERSION}"
}
extractKubeBinaries() {
  KUBE_BINARY_URL=${KUBE_BINARY_URL:-"https://kubernetesartifacts.azureedge.net/kubernetes/v${KUBERNETES_VERSION}/binaries/kubernetes-node-linux-amd64.tar.gz"}
  K8S_TGZ_TMP=${KUBE_BINARY_URL##*/}
  mkdir -p "${K8S_DOWNLOADS_DIR}"
  retrycmd_get_tarball 120 5 "$K8S_DOWNLOADS_DIR/${K8S_TGZ_TMP}" ${KUBE_BINARY_URL} || exit 31
  tar --transform="s|.*|&-${KUBERNETES_VERSION}|" --show-transformed-names -xzvf "$K8S_DOWNLOADS_DIR/${K8S_TGZ_TMP}" \
    --strip-components=3 -C /usr/local/bin kubernetes/node/bin/kubelet kubernetes/node/bin/kubectl
  rm -f "$K8S_DOWNLOADS_DIR/${K8S_TGZ_TMP}"
}
pullContainerImage() {
  CLI_TOOL=$1
  DOCKER_IMAGE_URL=$2
  retrycmd 60 1 1200 $CLI_TOOL pull $DOCKER_IMAGE_URL || exit 35
}
overrideNetworkConfig() {
  CONFIG_FILEPATH="/etc/cloud/cloud.cfg.d/80_azure_net_config.cfg"
  touch ${CONFIG_FILEPATH}
  cat <<EOF >>${CONFIG_FILEPATH}
datasource:
    Azure:
        apply_network_config: false
EOF
}
#EOF
