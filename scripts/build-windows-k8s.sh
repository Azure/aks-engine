#!/bin/bash
#
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT license.

set -eo pipefail

fetch_azs_k8s() {
	mkdir -p "$KUBEPATH"
	git clone https://github.com/msazurestackworkloads/kubernetes "$KUBEPATH" || true
	cd "$KUBEPATH"
	git fetch --all
}

create_version_branch() {
	git checkout -b "${AKS_BRANCH_NAME}" "${KUBERNETES_TAG_BRANCH}" || true
}

version_lt() {
  # shellcheck disable=SC2046
  [ "$1" != $(printf "%s\n%s" "$1" "$2" | sort -V | head -n 2 | tail -n 1) ]
}

version_ge() {
  # shellcheck disable=SC2046
  [ "$1" == $(printf "%s\n%s" "$1" "$2" | sort -V | head -n 2 | tail -n 1) ]
}

create_dist_dir() {
	mkdir -p "${DIST_DIR}"
}

build_kubelet() {
	echo "building kubelet.exe..."
	"$KUBEPATH"/build/run.sh make WHAT=cmd/kubelet KUBE_BUILD_PLATFORMS=windows/amd64
	cp "$KUBEPATH"/_output/dockerized/bin/windows/amd64/kubelet.exe "${DIST_DIR}"
}

build_kubeproxy() {
	echo "building kube-proxy.exe..."
	"$KUBEPATH"/build/run.sh make WHAT=cmd/kube-proxy KUBE_BUILD_PLATFORMS=windows/amd64
	cp "$KUBEPATH"/_output/dockerized/bin/windows/amd64/kube-proxy.exe "${DIST_DIR}"
}

build_kubectl() {
	echo "building kubectl.exe..."
    "$KUBEPATH"/build/run.sh make WHAT=cmd/kubectl KUBE_BUILD_PLATFORMS=windows/amd64
    cp "$KUBEPATH"/_output/dockerized/bin/windows/amd64/kubectl.exe "${DIST_DIR}"
}

get_kube_binaries() {
	if [ -n "${build_azs}" ]; then
		echo "building kubelet/kubeproxy from Azure Stack repo..."
		fetch_azs_k8s
		create_version_branch
		build_kube_binaries_for_upstream_e2e
	else
		echo "downloading kubelet/kubeproxy/kubectl from upstream..."
		WIN_TAR=kubernetes-node-windows-amd64.tar.gz
		SUB_DIR=kubernetes/node/bin
		curl -L https://storage.googleapis.com/kubernetes-release/release/v"${version}"/${WIN_TAR} -o "${TOP_DIR}"/${WIN_TAR}
		tar -xzvf "${TOP_DIR}"/${WIN_TAR} -C "${TOP_DIR}"
		cp "${TOP_DIR}"/${SUB_DIR}/kubelet.exe "${DIST_DIR}"
		cp "${TOP_DIR}"/${SUB_DIR}/kube-proxy.exe "${DIST_DIR}"
		cp "${TOP_DIR}"/${SUB_DIR}/kubectl.exe "${DIST_DIR}"
		chmod 775 "${DIST_DIR}"/kubectl.exe
	fi
}

build_kube_binaries_for_upstream_e2e() {
		"$KUBEPATH"/build/run.sh make WHAT=cmd/kubelet KUBE_BUILD_PLATFORMS=linux/amd64

		build_kubelet
		build_kubeproxy
		build_kubectl
}

download_nssm() {
	NSSM_VERSION=2.24
	NSSM_URL=https://k8stestinfrabinaries.blob.core.windows.net/nssm-mirror/nssm-${NSSM_VERSION}.zip
	echo "downloading nssm ..."
	curl ${NSSM_URL} -o /tmp/nssm-${NSSM_VERSION}.zip
	unzip -q -d /tmp /tmp/nssm-${NSSM_VERSION}.zip
	cp /tmp/nssm-${NSSM_VERSION}/win64/nssm.exe "${DIST_DIR}"
	chmod 775 "${DIST_DIR}"/nssm.exe
	rm -rf /tmp/nssm-${NSSM_VERSION}*
}

download_wincni() {
	mkdir -p "${DIST_DIR}"/cni/config
	WINSDN_URL=https://github.com/Microsoft/SDN/raw/master/Kubernetes/windows/
	WINBRIDGE_URL=https://github.com/Microsoft/SDN/raw/master/Kubernetes/flannel/l2bridge/
	WINBRIDGE_EXE=cni/win-bridge.exe
	HNS_PSM1=hns.psm1
	curl -L ${WINBRIDGE_URL}${WINBRIDGE_EXE} -o "${DIST_DIR}"/${WINBRIDGE_EXE}
	curl -L ${WINSDN_URL}${HNS_PSM1} -o "${DIST_DIR}"/${HNS_PSM1}
}

create_zip() {
	ZIP_NAME="${k8s_e2e_upstream_version:-${KUBERNETES_WIN_ZIP_FILENAME}}"
	cd "${DIST_DIR}"/..
	zip -r ../"${ZIP_NAME}" k/*
	cd -
}

upload_zip_to_blob_storage() {
	az storage blob upload -f "${TOP_DIR}"/../${KUBERNETES_WIN_ZIP_FILENAME} -c "${AZURE_STORAGE_CONTAINER_NAME}" -n ${KUBERNETES_WIN_ZIP_FILENAME}
}

cleanup_output() {
	rm "${TOP_DIR}"/../${KUBERNETES_WIN_ZIP_FILENAME}
	rm -r "${TOP_DIR}"
}

AKS_ENGINE_HOME="${GOPATH}"/src/github.com/Azure/aks-engine

usage() {
	echo "$0 [-v version] [-p acs_patch_version]"
	echo " -v <version>: version"
	echo " -p <patched version>: acs_patch_version"
	echo " -u <version build for kubernetes upstream e2e tests>: k8s_e2e_upstream_version"
	echo " -a <build Azure Stack specific windows file>: build_azs"
	echo " -z <zip path>: zip_path"
}

while getopts ":v:p:u:z:a:" opt; do
  case ${opt} in
    v)
      version=${OPTARG}
      ;;
    p)
      acs_patch_version=${OPTARG}
      ;;
	u)
	  k8s_e2e_upstream_version=${OPTARG}
	  ;;
	z)
	  zip_path=${OPTARG}
	  ;;
	a)
	  build_azs=${OPTARG}
	  ;;
    *)
	  usage
	  exit
      ;;
  esac
done

if [ -n "${build_azs}" ]; then
	KUBEPATH="${GOPATH}"/azurestack/src/k8s.io/kubernetes
else
	KUBEPATH="${GOPATH}"/src/k8s.io/kubernetes
fi

if [ -z "${k8s_e2e_upstream_version}" ]; then

	if [ -z "${version}" ] || [ -z "${acs_patch_version}" ]; then
		usage
		exit 1
	fi

	if [ -z "${AZURE_STORAGE_CONNECTION_STRING}" ] || [ -z "${AZURE_STORAGE_CONTAINER_NAME}" ]; then
		# shellcheck disable=SC2016
		echo '$AZURE_STORAGE_CONNECTION_STRING and $AZURE_STORAGE_CONTAINER_NAME need to be set for upload to Azure Blob Storage.'
		exit 1
	fi

	KUBERNETES_RELEASE=$(echo "$version" | cut -d'.' -f1,2)
	KUBERNETES_TAG_BRANCH=v${version}
	AKS_VERSION=${version}-${acs_patch_version}

	if [ -n "${build_azs}" ]; then
		AKS_BRANCH_NAME=azs-v${AKS_VERSION}
		TOP_DIR=${AKS_ENGINE_HOME}/_dist/k8s-windows-azs-v${AKS_VERSION}
		KUBERNETES_WIN_ZIP_FILENAME=azs-v"${AKS_VERSION}"int.zip
	else
		AKS_BRANCH_NAME=acs-v${AKS_VERSION}
		TOP_DIR=${AKS_ENGINE_HOME}/_dist/k8s-windows-v${AKS_VERSION}
		KUBERNETES_WIN_ZIP_FILENAME=v"${AKS_VERSION}"int.zip
	fi

	DIST_DIR="${TOP_DIR}"/k

	create_dist_dir
	get_kube_binaries
	download_nssm
	download_wincni
	create_zip
	upload_zip_to_blob_storage
	cleanup_output

else
	DIST_DIR=${zip_path}/k
	create_dist_dir
	build_kube_binaries_for_upstream_e2e
	download_nssm
	download_wincni
	create_zip
fi
