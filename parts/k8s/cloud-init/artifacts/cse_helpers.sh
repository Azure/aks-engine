#!/bin/bash

OS=$(sort -r /etc/*-release | gawk 'match($0, /^(ID=(.*))$/, a) { print toupper(a[2] a[3]); exit }')
UBUNTU_OS_NAME="UBUNTU"
RHEL_OS_NAME="RHEL"
DEBIAN_OS_NAME="DEBIAN"
if ! echo "${UBUNTU_OS_NAME} ${RHEL_OS_NAME} ${DEBIAN_OS_NAME}" | grep -q "${OS}"; then
  OS=$(sort -r /etc/*-release | gawk 'match($0, /^(ID_LIKE=(.*))$/, a) { print toupper(a[2] a[3]); exit }')
fi
KUBECTL=/usr/local/bin/kubectl
DOCKER=/usr/bin/docker
export GPU_DV=418.40.04
export GPU_DEST=/usr/local/nvidia
NVIDIA_DOCKER_VERSION=2.0.3
DOCKER_VERSION=1.13.1-1
NVIDIA_CONTAINER_RUNTIME_VERSION=2.0.0
NVIDIA_DOCKER_SUFFIX=docker18.09.2-1

configure_prerequisites() {
  ip_forward_path=/proc/sys/net/ipv4/ip_forward
  ip_forward_setting="net.ipv4.ip_forward=0"
  sysctl_conf=/etc/sysctl.conf
  if ! grep -qE "^1$" ${ip_forward_path}; then
    echo 1 >${ip_forward_path}
  fi
  if grep -qE "${ip_forward_setting}" ${sysctl_conf}; then
    sed -i '/^net.ipv4.ip_forward=0$/d' ${sysctl_conf}
  fi
}

aptmarkWALinuxAgent() {
  wait_for_apt_locks
  retrycmd_if_failure 120 5 25 apt-mark $1 walinuxagent ||
    if [[ $1 == "hold" ]]; then
      exit 7
    elif [[ $1 == "unhold" ]]; then
      exit 8
    fi
}

retrycmd_if_failure() {
  retries=$1; wait_sleep=$2; timeout=$3; shift && shift && shift
  for i in $(seq 1 $retries); do
    timeout $timeout ${@} && break ||
      if [ $i -eq $retries ]; then
        echo Executed \"$@\" $i times
        return 1
      else
        sleep $wait_sleep
      fi
  done
  echo Executed \"$@\" $i times
}
retrycmd_if_failure_no_stats() {
  retries=$1; wait_sleep=$2; timeout=$3; shift && shift && shift
  for i in $(seq 1 $retries); do
    timeout $timeout ${@} && break ||
      if [ $i -eq $retries ]; then
        return 1
      else
        sleep $wait_sleep
      fi
  done
}
retrycmd_get_tarball() {
  tar_retries=$1; wait_sleep=$2; tarball=$3; url=$4
  echo "${tar_retries} retries"
  for i in $(seq 1 $tar_retries); do
    tar -tzf $tarball && break ||
      if [ $i -eq $tar_retries ]; then
        return 1
      else
        timeout 60 curl -fsSL $url -o $tarball
        sleep $wait_sleep
      fi
  done
}
retrycmd_get_executable() {
  retries=$1; wait_sleep=$2; filepath=$3; url=$4; validation_args=$5
  echo "${retries} retries"
  for i in $(seq 1 $retries); do
    $filepath $validation_args && break ||
      if [ $i -eq $retries ]; then
        return 1
      else
        timeout 30 curl -fsSL $url -o $filepath
        chmod +x $filepath
        sleep $wait_sleep
      fi
  done
}
wait_for_file() {
  retries=$1; wait_sleep=$2; filepath=$3
  paved=/opt/azure/cloud-init-files.paved
  grep -Fq "${filepath}" $paved && return 0
  for i in $(seq 1 $retries); do
    grep -Fq '#EOF' $filepath && break
    if [ $i -eq $retries ]; then
      return 1
    else
      sleep $wait_sleep
    fi
  done
  sed -i "/#EOF/d" $filepath
  echo $filepath >>$paved
}
wait_for_apt_locks() {
  while fuser /var/lib/dpkg/lock /var/lib/apt/lists/lock /var/cache/apt/archives/lock >/dev/null 2>&1; do
    echo 'Waiting for release of apt locks'
    sleep 3
  done
}
apt_get_update() {
  retries=10
  apt_update_output=/tmp/apt-get-update.out
  for i in $(seq 1 $retries); do
    wait_for_apt_locks
    export DEBIAN_FRONTEND=noninteractive
    dpkg --configure -a --force-confdef
    apt-get -f -y install
    ! (apt-get update 2>&1 | tee $apt_update_output | grep -E "^([WE]:.*)|([eE]rr.*)$") &&
      cat $apt_update_output && break ||
      cat $apt_update_output
    if [ $i -eq $retries ]; then
      return 1
    else sleep 5
    fi
  done
  echo Executed apt-get update $i times
}
apt_get_install() {
  retries=$1; wait_sleep=$2; timeout=$3; shift && shift && shift
  for i in $(seq 1 $retries); do
    wait_for_apt_locks
    export DEBIAN_FRONTEND=noninteractive
    dpkg --configure -a --force-confdef
    apt-get install -o Dpkg::Options::="--force-confold" --no-install-recommends -y ${@} && break ||
      if [ $i -eq $retries ]; then
        return 1
      else
        sleep $wait_sleep
        apt_get_update
      fi
  done
  echo Executed apt-get install --no-install-recommends -y \"$@\" $i times
}
apt_get_purge() {
  retries=20; wait_sleep=30; timeout=120
  for package in $@; do
    if apt list --installed | grep $package; then
      for i in $(seq 1 $retries); do
        wait_for_apt_locks
        export DEBIAN_FRONTEND=noninteractive
        dpkg --configure -a --force-confdef
        apt-get purge -o Dpkg::Options::="--force-confold" -y $package && break ||
          if [ $i -eq $retries ]; then
            return 1
          else
            sleep $wait_sleep
          fi
      done
    fi
  done
  echo Executed apt-get purge -y \"$package\" $i times
}
apt_get_dist_upgrade() {
  retries=10
  apt_dist_upgrade_output=/tmp/apt-get-dist-upgrade.out
  for i in $(seq 1 $retries); do
    wait_for_apt_locks
    export DEBIAN_FRONTEND=noninteractive
    dpkg --configure -a --force-confdef
    apt-get -f -y install
    apt-mark showhold
    ! (apt-get dist-upgrade -y 2>&1 | tee $apt_dist_upgrade_output | grep -E "^([WE]:.*)|([eE]rr.*)$") && \
    cat $apt_dist_upgrade_output && break || \
    cat $apt_dist_upgrade_output
    if [ $i -eq $retries ]; then
      return 1
    else sleep 5
    fi
  done
  echo Executed apt-get dist-upgrade $i times
}
systemctl_restart() {
  retries=$1; wait_sleep=$2; timeout=$3 svcname=$4
  for i in $(seq 1 $retries); do
    timeout $timeout systemctl daemon-reload
    timeout $timeout systemctl restart $svcname && break ||
      if [ $i -eq $retries ]; then
        return 1
      else
        sleep $wait_sleep
      fi
  done
}
systemctl_stop() {
  retries=$1; wait_sleep=$2; timeout=$3 svcname=$4
  for i in $(seq 1 $retries); do
    timeout $timeout systemctl daemon-reload
    timeout $timeout systemctl stop $svcname && break ||
      if [ $i -eq $retries ]; then
        return 1
      else
        sleep $wait_sleep
      fi
  done
}
sysctl_reload() {
  retries=$1; wait_sleep=$2; timeout=$3
  for i in $(seq 1 $retries); do
    timeout $timeout sysctl --system && break ||
      if [ $i -eq $retries ]; then
        return 1
      else
        sleep $wait_sleep
      fi
  done
}
version_gte() {
  test "$(printf '%s\n' "$@" | sort -rV | head -n 1)" == "$1"
}

#HELPERSEOF
