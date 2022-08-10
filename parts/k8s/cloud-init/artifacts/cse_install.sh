#!/bin/bash

CNI_CONFIG_DIR="/etc/cni/net.d"
CNI_BIN_DIR="/opt/cni/bin"
CNI_DOWNLOADS_DIR="/opt/cni/downloads"
CONTAINERD_DOWNLOADS_DIR="/opt/containerd/downloads"
APMZ_DOWNLOADS_DIR="/opt/apmz/downloads"
UBUNTU_RELEASE=$(lsb_release -r -s)
UBUNTU_CODENAME=$(lsb_release -c -s)
NVIDIA_PACKAGES="libnvidia-container1 libnvidia-container-tools nvidia-container-toolkit"
NVIDIA_CONTAINER_TOOLKIT_VER=1.6.0
NVIDIA_RUNTIME_VER=3.6.0

disableTimeSyncd() {
  systemctl_stop 20 5 10 systemd-timesyncd || exit 3
  retrycmd 120 5 25 systemctl disable systemd-timesyncd || exit 3
}
installEtcd() {
  local  v
  v=$(etcd --version | grep "etcd Version" | cut -d ":" -f 2 | tr -d '[:space:]')
  if [[ $v != "${ETCD_VERSION}" ]]; then
    local cli_tool=$1 path="/usr/bin" image=${ETCD_DOWNLOAD_URL}etcd:v${ETCD_VERSION}
    pullContainerImage $cli_tool ${image}
    removeEtcd
    if [[ $cli_tool == "docker" ]]; then
      mkdir -p "$path"
      docker run --rm --entrypoint cat ${image} /usr/local/bin/etcd >"$path/etcd"
      docker run --rm --entrypoint cat ${image} /usr/local/bin/etcdctl >"$path/etcdctl"
    else
      tmpdir=/root/etcd${RANDOM}
      img unpack -o ${tmpdir} ${image}
      mv ${tmpdir}/usr/local/bin/etcd ${tmpdir}/usr/local/bin/etcdctl ${path}
      rm -rf ${tmpdir}
    fi
    chmod a+x "$path/etcd" "$path/etcdctl"
  fi
}
installDeps() {
  packages="apache2-utils apt-transport-https blobfuse=1.1.1 ca-certificates cifs-utils conntrack cracklib-runtime dbus dkms ebtables ethtool fuse gcc git htop iftop init-system-helpers iotop iproute2 ipset iptables jq libpam-pwquality libpwquality-tools linux-headers-$(uname -r) make mount nfs-common pigz socat sysstat traceroute util-linux xz-utils zip"
  if [[ ${OS} == "${UBUNTU_OS_NAME}" ]]; then
    retrycmd_no_stats 120 5 25 curl -fsSL ${MS_APT_REPO}/config/ubuntu/${UBUNTU_RELEASE}/packages-microsoft-prod.deb >/tmp/packages-microsoft-prod.deb || exit 42
    retrycmd 60 5 10 dpkg -i /tmp/packages-microsoft-prod.deb || exit 43
    retrycmd_no_stats 120 5 25 curl ${MS_APT_REPO}/config/ubuntu/${UBUNTU_RELEASE}/prod.list >/tmp/microsoft-prod.list || exit 25
    retrycmd 10 5 10 cp /tmp/microsoft-prod.list /etc/apt/sources.list.d/ || exit 25
    retrycmd_no_stats 120 5 25 curl ${MS_APT_REPO}/keys/microsoft.asc | gpg --dearmor >/tmp/microsoft.gpg || exit 26
    retrycmd 10 5 10 cp /tmp/microsoft.gpg /etc/apt/trusted.gpg.d/ || exit 26
    aptmarkWALinuxAgent hold
    packages+=" cgroup-lite ceph-common glusterfs-client"
    if [[ $UBUNTU_RELEASE == "18.04" ]]; then
      disableTimeSyncd
      packages+=" ntp ntpstat chrony"
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
gpuDriversDownloaded() {
  for apt_package in $NVIDIA_PACKAGES; do
    ls ${PERMANENT_CACHE_DIR}${apt_package}* || return 1
  done
  ls ${PERMANENT_CACHE_DIR}nvidia-container-runtime* || return 1
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
  mkdir -p $PERMANENT_CACHE_DIR
  for apt_package in $NVIDIA_PACKAGES; do
    apt_get_download 20 30 "${apt_package}=${NVIDIA_CONTAINER_TOOLKIT_VER}*" || exit 85
    cp -al ${APT_CACHE_DIR}${apt_package}_${NVIDIA_CONTAINER_TOOLKIT_VER}* $PERMANENT_CACHE_DIR || exit 85
  done
  apt_get_download 20 30 nvidia-container-runtime=${NVIDIA_RUNTIME_VER}* || exit 85
  cp -al ${APT_CACHE_DIR}nvidia-container-runtime_${NVIDIA_RUNTIME_VER}* $PERMANENT_CACHE_DIR || exit 85
}
removeMoby() {
  apt_get_purge moby-engine moby-cli || exit 27
}
removeContainerd() {
  apt_get_purge moby-containerd || exit 27
}
mobyPkgVersion() {
  dpkg -s "${1}" | grep "Version:" | awk '{ print $2 }' | cut -d '+' -f 1
}
installRunc() {
  local v
  v=$(runc --version | head -n 1 | cut -d" " -f3)
  if [[ $v != "1.1.2" ]]; then
    apt_get_install 20 30 120 moby-runc=1.1.2* --allow-downgrades || exit 27
  fi
}
installMoby() {
  local install_pkgs="" v cli_ver="${MOBY_VERSION}"
  v="$(mobyPkgVersion moby-containerd)"
  if [ -n "${CONTAINERD_VERSION}" ] && [ ! "${v}" = "${CONTAINERD_VERSION}" ]; then
    install_pkgs+=" moby-containerd=${CONTAINERD_VERSION}*"
    removeMoby
    removeContainerd
  fi
  v="$(mobyPkgVersion moby-engine)"
  if [ ! "${v}" = "${MOBY_VERSION}" ]; then
    install_pkgs+=" moby-engine=${MOBY_VERSION}*"
    if [ "${cli_ver}" = "3.0.4" ]; then
      cli_ver="3.0.3"
    fi
    install_pkgs+=" moby-cli=${cli_ver}*"
    removeMoby
  fi
  if [ -n "${install_pkgs}" ]; then
    apt_get_install 20 30 120 ${install_pkgs} --allow-downgrades || exit 27
  fi
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
  local ver=$1 v
  local d="$APMZ_DOWNLOADS_DIR/$ver"
  local url="https://upstreamartifacts.azureedge.net/apmz/$ver/binaries/apmz_linux_amd64.tar.gz" fp="/usr/local/bin/apmz" dest="$d/apmz.gz" bin_fp="$d/apmz_linux_amd64"
  if [[ $OS == $FLATCAR_OS_NAME ]]; then
    fp="/opt/bin/apmz"
    export PATH="${PATH}:/opt/bin"
  fi
  if [[ -f $fp ]]; then
    v=$($fp version)
    if [[ $ver == "$v" ]]; then
      return
    fi
  fi
  mkdir -p "$d"
  retrycmd_get_tarball 120 5 "$dest" "${url}"
  tar -xvf "$dest" -C "$d"
  chmod +x "$bin_fp"
  ln -Ffs "$bin_fp" "$fp"
}
installBpftrace() {
  local ver="v0.9.4" v bin="bpftrace" tools="bpftrace-tools.tar"
  local url="https://upstreamartifacts.azureedge.net/$bin/$ver"
  local bpftrace_fp="/usr/local/bin/$bin"
  local tools_fp="/usr/local/share/$bin"
  if [[ $OS == $FLATCAR_OS_NAME ]]; then
    bpftrace_fp="/opt/bin/$bin"
    tools_fp="/opt/share/$bin"
    export PATH="${PATH}:/opt/bin"
  fi
  if [[ -f $bpftrace_fp ]]; then
    v="$($bin -V | cut -d' ' -f2)"
    if [[ $ver == "$v" ]]; then
      return
    fi
    rm "$bpftrace_fp"
    if [[ -d $tools_fp ]]; then
      rm -r "$tools_fp"
    fi
  fi
  mkdir -p "$tools_fp"
  install_dir="/opt/bpftrace/downloads/$ver"
  mkdir -p "$install_dir"
  download_path="$install_dir/$tools"
  retrycmd 30 5 60 curl -fSL -o "$bpftrace_fp" "$url/$bin" || exit 169
  retrycmd 30 5 60 curl -fSL -o "$download_path" "$url/$tools" || exit 170
  tar -xvf "$download_path" -C "$tools_fp"
  chmod +x "$bpftrace_fp"
  chmod -R +x "$tools_fp/tools"
}
installImg() {
  img_filepath=/usr/local/bin/img
  retrycmd_get_executable 120 5 $img_filepath "https://upstreamartifacts.azureedge.net/img/img-linux-amd64-v0.5.6" ls || exit 33
}
extractHyperkube() {
  local cli_tool=$1 fp="/home/hyperkube-downloads/${KUBERNETES_VERSION}" dest="/usr/local/bin"
  if [[ $OS == $FLATCAR_OS_NAME ]]; then
    dest="/opt/bin"
  fi
  pullContainerImage $cli_tool ${HYPERKUBE_URL}
  if [[ $cli_tool == "docker" ]]; then
    mkdir -p "$fp"
    if docker run --rm --entrypoint "" -v $path:$path ${HYPERKUBE_URL} /bin/bash -c "cp $dest/{kubelet,kubectl} $fp"; then
      mv "${fp}/kubelet" "${dest}/kubelet-${KUBERNETES_VERSION}"
      mv "${fp}/kubectl" "${dest}/kubectl-${KUBERNETES_VERSION}"
      return
    else
      docker run --rm -v $fp:$fp ${HYPERKUBE_URL} /bin/bash -c "cp /hyperkube $fp"
    fi
  else
    img unpack -o "$fp" ${HYPERKUBE_URL}
  fi

  cp "${fp}/hyperkube" "${dest}/kubelet-${KUBERNETES_VERSION}"
  mv "${fp}/hyperkube" "${dest}/kubectl-${KUBERNETES_VERSION}"
  if [[ $OS == $FLATCAR_OS_NAME ]]; then
    chmod a+x ${dest}/kubelet-${KUBERNETES_VERSION} ${dest}/kubectl-${KUBERNETES_VERSION}
  fi
}
extractKubeBinaries() {
  KUBE_BINARY_URL=${KUBE_BINARY_URL:-"https://kubernetesartifacts.azureedge.net/kubernetes/v${KUBERNETES_VERSION}/binaries/kubernetes-node-linux-amd64.tar.gz"}
  local dest="/opt/kubernetes/downloads" tmpDir=${KUBE_BINARY_URL##*/}
  mkdir -p "${dest}"
  retrycmd_get_tarball 120 5 "$dest/${tmpDir}" ${KUBE_BINARY_URL} || exit 31
  path=/usr/local/bin
  if [[ $OS == $FLATCAR_OS_NAME ]]; then
    path=/opt/bin
  fi
  tar --transform="s|.*|&-${KUBERNETES_VERSION}|" --show-transformed-names -xzvf "$dest/${tmpDir}" \
    --strip-components=3 -C ${path} kubernetes/node/bin/kubelet kubernetes/node/bin/kubectl
  rm -f "$dest/${tmpDir}"
}
pullContainerImage() {
  local cli_tool=$1 url=$2
  retrycmd 60 1 1200 $cli_tool pull $url || exit 35
}
loadContainerImage() {
  docker pull $1 || exit 35
  docker save $1 | ctr -n=k8s.io images import - || exit 35

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
