#!/bin/bash
set -eo pipefail

fetch_k8s() {
	git clone https://github.com/Azure/kubernetes ${GOPATH}/src/k8s.io/kubernetes || true
	cd $KUBEPATH
	git remote add upstream https://github.com/kubernetes/kubernetes || true
	git fetch upstream
}

set_git_config() {
	git config user.name "AKS CI"
	git config user.email "containers@microsoft.com"
}

create_version_branch() {
	git checkout -b ${AKS_BRANCH_NAME} ${KUBERNETES_TAG_BRANCH} || true
}

version_lt() {
        [ "$1" != $(printf "$1\n$2" | sort -V | head -n 2 | tail -n 1) ]
}

version_ge() {
        [ "$1" == $(printf "$1\n$2" | sort -V | head -n 2 | tail -n 1) ]
}

k8s_16_cherry_pick() {
	# 232fa6e5bc (HEAD -> release-1.6, origin/release-1.6) Fix the delay caused by network setup in POD Infra container
	# 02b1c2b9e2 Use dns policy to determine setting DNS servers on the correct NIC in Windows container
	# 4c2a2d79aa Fix getting host DNS for Windows node when dnsPolicy is Default
	# caa314ccdc Update docker version parsing to allow nonsemantic versions as they have changed how they do their versions
	# f18be40948 Fix the issue in unqualified name where DNS client such as ping or iwr validate name in response and original question. Switch to use miekg's DNS library
	# c862b583c9 Remove DNS server from NAT network adapter inside container
	# f6c27f9375 Merged libCNI-on-Windows changes from CNI release 0.5.0, PRs 359 and 361
	# 4f196c6cac Fix the issue that ping uses the incorrect NIC to resolve name sometimes
	# 2c9fd27449 Workaround for Outbound Internet traffic in Azure Kubernetes
	# 5fa0725025 Use adapter vEthernet (HNSTransparent) on Windows host network to find node IP

	git cherry-pick 5fa0725025^..232fa6e5bc
}

k8s_17_cherry_pick() {
	if version_ge "${version}" "1.7.10"; then
		echo "version 1.7.10 and after..."
		# In 1.7.10, the following commit is not needed and has conflict with 137f4cb16e
		# due to the out-of-order back porting into Azure 1.7. So removing it.
		# cee32e92f7 fix#50150: azure disk mount failure on coreos
		git revert --no-edit cee32e92f7 || true

		if version_ge "${version}" "1.7.12"; then
			echo "version 1.7.12 and after..."
			# In 1.7.12, the following commits are cherry-picked in upstream and has conflict
			# with 137f4cb16e. So removing them.
			git revert --no-edit 593653c384 || true #only for linux
			git revert --no-edit 7305738dd1 || true #add tests only
			git revert --no-edit e01bafcf80 || true #only for linux
			git revert --no-edit afd79db7a6 || true #only for linux
			git revert --no-edit 3a4abca2f7 || true #covered by commit 3aa179744f
			git revert --no-edit 6a2e2f47d3 || true #covered by commit 3aa179744f

			if version_ge "${version}" "1.7.13"; then
				echo "version 1.7.13 and after..."
				# In 1.7.13, the following commit is cherry-picked in upstream and has conflict
				# with 137f4cb16e. So removing it.
				git revert --no-edit 3aa179744f || true #only for linux

				if version_ge "${version}" "1.7.14"; then
					echo "version 1.7.14 and after..."
					if version_ge "${version}" "1.7.15"; then
						echo "version 1.7.15 and after..."
						if version_ge "${version}" "1.7.16"; then
							echo "version 1.7.16 and after..."
							# From 1.7.16, 3e530479ed conflict with reverting 975d0a4bb9. We use 5f9113ce9b in Azure repo instead of 3e530479ed
							git revert --no-edit 3e530479ed || true
							# From 1.7.16, 76017a002b conflict with reverting 975d0a4bb9. We use 3d97a43e41 in Azure repo instead of 76017a002b
							git revert --no-edit 76017a002b || true
						fi

						# From 1.7.15, 0d5df36f05 conflict with reverting 975d0a4bb9. We use 631e363b9d in Azure repo instead of 0d5df36f05
						git revert --no-edit 0d5df36f05 || true
						# From 1.7.15, 4d50500614 conflict with reverting 51584188ee. It is server side code so jsut revert it.
						git revert --no-edit 4d50500614 || true
					fi
					# From 1.7.14, 975d0a4bb9 conflict with 137f4cb16e. We use a36d59ddda in Azure repo instead of 975d0a4bb9
					git revert --no-edit 975d0a4bb9 || true
					# From 1.7.14, 51584188ee and 3a0db21dcb conflict with af3a93b07e. We use f5c45d3def in Azure repo instead of them
					git revert --no-edit 51584188ee || true
					git revert --no-edit 3a0db21dcb || true
					# From 1.7.14, 273411cc90 conflict with e9591ef03e. We use d18812a049 in Azure repo instead of 273411cc90
					git revert --no-edit 273411cc90 || true
					# From 1.7.14, a97f60fbf3 conflict with a36d59ddda. We use 69c56c6037 in Azure repo instead of a97f60fbf3
					git revert --no-edit a97f60fbf3 || true
				fi
			fi
		fi
	fi

	# 32ceaa7918 fix #60625: add remount logic for azure file plugin on Windows
	# ...
	# b8fe713754 Use adapter vEthernet (HNSTransparent) on Windows host network to find node IP
	# From 1.7.13, acbdec96da is not needed since 060111c603 supercedes it
	git cherry-pick --allow-empty --keep-redundant-commits b8fe713754^..e912889a7f
	if version_lt "${version}" "1.7.13"; then
		echo "before version 1.7.13..."
		git cherry-pick --allow-empty --keep-redundant-commits acbdec96da
	fi
	git cherry-pick --allow-empty --keep-redundant-commits 76d7c23f62^..32ceaa7918

	if version_ge "${version}" "1.7.14"; then
		echo "version 1.7.14 and after..."
		# From 1.7.14, 975d0a4bb9 conflict with 137f4cb16e. We use a36d59ddda in Azure repo instead of 975d0a4bb9
		git cherry-pick --allow-empty --keep-redundant-commits a36d59ddda
		# From 1.7.14, 51584188ee and 3a0db21dcb conflict with af3a93b07e. We use f5c45d3def in Azure repo instead of them
		git cherry-pick -X theirs --allow-empty --keep-redundant-commits f5c45d3def # git complains about a conflict and just take theirs
		# From 1.7.14, 273411cc90 conflict with e9591ef03e. We use d18812a049 in Azure repo instead of 273411cc90
		git cherry-pick --allow-empty --keep-redundant-commits d18812a049
		# From 1.7.14, a97f60fbf3 conflict with a36d59ddda. We use 69c56c6037 in Azure repo instead of a97f60fbf3
		git cherry-pick --allow-empty --keep-redundant-commits 69c56c6037

		git cherry-pick --allow-empty --keep-redundant-commits 3e930be6bc

		if version_ge "${version}" "1.7.15"; then
			echo "version 1.7.15 and after..."
			# From 1.7.15, 0d5df36f05 conflict with reverting 975d0a4bb9. We use 631e363b9d in Azure repo instead of 0d5df36f05
			git cherry-pick --allow-empty --keep-redundant-commits 631e363b9d

			if version_ge "${version}" "1.7.16"; then
				echo "version 1.7.16 and after..."
				# From 1.7.16, 76017a002b conflict with reverting 975d0a4bb9. We use 3d97a43e41 in Azure repo instead of 76017a002b
				git cherry-pick --allow-empty --keep-redundant-commits 3d97a43e41
				# From 1.7.16, 3e530479ed conflict with reverting 975d0a4bb9. We use 5f9113ce9b in Azure repo instead of 3e530479ed
				git cherry-pick --allow-empty --keep-redundant-commits 5f9113ce9b
			fi
		fi
	fi
}

k8s_18_cherry_pick() {
	# 4fd355d04a fix #60625: add remount logic for azure file plugin on Windows
	# ...
	# 4647f2f616 merge #52401: add windows implementation of GetMountRefs

	# d75ef50170 merge#54334: fix azure disk mount failure on coreos and some other distros

	# Only cherry pick before 1.8.6 due to merge conflict with 1.8.6 and up: b8594873f4 merge #50673: azure disk fix of SovereignCloud support

	# cb29df51c0 Fixing 'targetport' to service 'port' mapping
	# ...
	# b42981f90b merge#53629: fix azure file mount limit issue on windows due to using drive letter

	# We have to skip 8b3345114e due to merge conflict with 1.8.4 and up, 4647f2f616 has the change.

	# 8d477271f7 update bazel
	# ...
	# 69644018c8 Use adapter vEthernet (HNSTransparent) on Windows host network to find node IP

	if version_ge "${version}" "1.8.9"; then
		echo "version 1.8.9 and after..."
		if version_ge "${version}" "1.8.10"; then
			echo "version 1.8.10 and after..."
			if version_ge "${version}" "1.8.11"; then
				echo "version 1.8.11 and after..."
				# From 1.8.11, 6487f82b13 conflict with 86107edff5. We use 8e081d5da9 in Azure repo instead of 6487f82b13
				git revert --no-edit 6487f82b13 || true
				# From 1.8.11, 2704ef877b conflict with reverting a418057ddd. We skip it since it is test file
				git revert --no-edit 2704ef877b || true
				# From 1.8.11, dc21ebe936 conflict with reverting 63b4f60e43. We use 18589971f7 in Azure repo instead of dc21ebe936
				git revert --no-edit dc21ebe936 || true
				# From 1.8.11, a418057ddd conflict with reverting 63b4f60e43. We use 57038dce90 in Azure repo instead of a418057ddd
				git revert --no-edit a418057ddd || true
				# From 1.8.11, beaca32611 conflict with reverting 63b4f60e43. We use e6d8af84df in Azure repo instead of beaca32611
				git revert --no-edit beaca32611 || true
			fi
			# From 1.8.10, 5936faed37 conflict with reverting 63b4f60e43. We use 1f26b7a083 in Azure repo instead of 5936faed37
			git revert --no-edit 5936faed37 || true
		fi
		# From 1.8.9, 63b4f60e43 conflict with b42981f90b. We use 6a8305e419 in Azure repo instead of 63b4f60e43
		git revert --no-edit 63b4f60e43 || true
		# From 1.8.9, 40d5e0a34f conflict with 6a8305e419. We use b90d61a48c in Azure repo instead of 40d5e0a34f
		git revert --no-edit 40d5e0a34f || true
	fi

	git cherry-pick --allow-empty --keep-redundant-commits 69644018c8^..8d477271f7
	git cherry-pick --allow-empty --keep-redundant-commits b42981f90b^..cb29df51c0
	if version_lt "${version}" "1.8.6"; then
		echo "before version 1.8.6..."
		git cherry-pick --allow-empty --keep-redundant-commits b8594873f4
	fi
	git cherry-pick --allow-empty --keep-redundant-commits d75ef50170
	git cherry-pick --allow-empty --keep-redundant-commits 4647f2f616^..4fd355d04a

	if version_ge "${version}" "1.8.9"; then
		echo "version 1.8.9 and after..."
		# From 1.8.9, 63b4f60e43 conflict with b42981f90b. We use 6a8305e419 in Azure repo instead of 63b4f60e43
		git cherry-pick --allow-empty --keep-redundant-commits 6a8305e419
		# From 1.8.9, 40d5e0a34f conflict with 6a8305e419. We use b90d61a48c in Azure repo instead of 40d5e0a34f
		git cherry-pick --allow-empty --keep-redundant-commits b90d61a48c
		if version_ge "${version}" "1.8.10"; then
			echo "version 1.8.10 and after..."
			# From 1.8.10, 5936faed37 conflict with reverting 63b4f60e43. We use 1f26b7a083 in Azure repo instead of 5936faed37
			git cherry-pick --allow-empty --keep-redundant-commits 1f26b7a083
			if version_ge "${version}" "1.8.11"; then
				echo "version 1.8.11 and after..."
				# From 1.8.11, 6487f82b13 conflict with 86107edff5. We use 8e081d5da9 in Azure repo instead of 6487f82b13
				# From 1.8.11, dc21ebe936 conflict with reverting 63b4f60e43. We use 18589971f7 in Azure repo instead of dc21ebe936
				# From 1.8.11, a418057ddd conflict with reverting 63b4f60e43. We use 57038dce90 in Azure repo instead of a418057ddd
				# From 1.8.11, beaca32611 conflict with reverting 63b4f60e43. We use e6d8af84df in Azure repo instead of beaca32611
				git cherry-pick --allow-empty --keep-redundant-commits 8e081d5da9^..18589971f7
			fi
		fi
	fi
}

apply_aks_cherry_picks() {
	if [ "${KUBERNETES_RELEASE}" == "1.6" ]; then
		k8s_16_cherry_pick
	elif [ "${KUBERNETES_RELEASE}" == "1.7" ]; then
		k8s_17_cherry_pick
	elif [ "${KUBERNETES_RELEASE}" == "1.8" ]; then
		k8s_18_cherry_pick
	elif version_ge "${KUBERNETES_RELEASE}" "1.9"; then
		echo "No need to cherry-pick for version greater than or equal to 1.9!"
	else
		echo "Unable to apply cherry picks for ${KUBERNETES_RELEASE}."
		exit 1
	fi
}

create_dist_dir() {
	mkdir -p ${DIST_DIR}
}

build_kubelet() {
	echo "building kubelet.exe..."
	$KUBEPATH/build/run.sh make WHAT=cmd/kubelet KUBE_BUILD_PLATFORMS=windows/amd64
	cp ${GOPATH}/src/k8s.io/kubernetes/_output/dockerized/bin/windows/amd64/kubelet.exe ${DIST_DIR}
}

build_kubeproxy() {
	echo "building kube-proxy.exe..."
	$KUBEPATH/build/run.sh make WHAT=cmd/kube-proxy KUBE_BUILD_PLATFORMS=windows/amd64
	cp ${GOPATH}/src/k8s.io/kubernetes/_output/dockerized/bin/windows/amd64/kube-proxy.exe ${DIST_DIR}
}

build_kubectl() {
	echo "building kubectl.exe..."
    $KUBEPATH/build/run.sh make WHAT=cmd/kubectl KUBE_BUILD_PLATFORMS=windows/amd64
    cp ${GOPATH}/src/k8s.io/kubernetes/_output/dockerized/bin/windows/amd64/kubectl.exe ${DIST_DIR}
}

download_kubectl() {
	kubectl="https://storage.googleapis.com/kubernetes-release/release/v${version}/bin/windows/amd64/kubectl.exe"
	echo "dowloading ${kubectl} ..."
	curl -L ${kubectl} -o ${DIST_DIR}/kubectl.exe
	chmod 775 ${DIST_DIR}/kubectl.exe
}

get_kube_binaries() {
	if version_lt "${KUBERNETES_RELEASE}" "1.9"; then
		echo "building kubelet/kubeproxy from azure repo..."
		fetch_k8s
		set_git_config
		create_version_branch
		apply_aks_cherry_picks

		# Due to what appears to be a bug in the Kubernetes Windows build system, one
		# has to first build a linux binary to generate _output/bin/deepcopy-gen.
		# Building to Windows w/o doing this will generate an empty deepcopy-gen.
		build/run.sh make WHAT=cmd/kubelet KUBE_BUILD_PLATFORMS=linux/amd64

		build_kubelet
		build_kubeproxy

		echo "downloading kubectl..."
		download_kubectl
	else
		echo "downloading kubelet/kubeproxy/kubectl from upstream..."
		WIN_TAR=kubernetes-node-windows-amd64.tar.gz
		SUB_DIR=kubernetes/node/bin
		curl -L https://storage.googleapis.com/kubernetes-release/release/v${version}/${WIN_TAR} -o ${TOP_DIR}/${WIN_TAR}
		tar -xzvf ${TOP_DIR}/${WIN_TAR} -C ${TOP_DIR}
		cp ${TOP_DIR}/${SUB_DIR}/kubelet.exe ${DIST_DIR}
		cp ${TOP_DIR}/${SUB_DIR}/kube-proxy.exe ${DIST_DIR}
		cp ${TOP_DIR}/${SUB_DIR}/kubectl.exe ${DIST_DIR}
		chmod 775 ${DIST_DIR}/kubectl.exe
	fi
}

build_kube_binaries_for_upstream_e2e() {
		$KUBEPATH/build/run.sh make WHAT=cmd/kubelet KUBE_BUILD_PLATFORMS=linux/amd64

		build_kubelet
		build_kubeproxy
		build_kubectl
}

download_nssm() {
	NSSM_VERSION=2.24
	NSSM_URL=https://nssm.cc/release/nssm-${NSSM_VERSION}.zip
	echo "downloading nssm ..."
	curl ${NSSM_URL} -o /tmp/nssm-${NSSM_VERSION}.zip
	unzip -q -d /tmp /tmp/nssm-${NSSM_VERSION}.zip
	cp /tmp/nssm-${NSSM_VERSION}/win64/nssm.exe ${DIST_DIR}
	chmod 775 ${DIST_DIR}/nssm.exe
	rm -rf /tmp/nssm-${NSSM_VERSION}*
}

download_wincni() {
	mkdir -p ${DIST_DIR}/cni/config
	WINSDN_URL=https://github.com/Microsoft/SDN/raw/master/Kubernetes/windows/
	WINCNI_EXE=cni/wincni.exe
	HNS_PSM1=hns.psm1
	curl -L ${WINSDN_URL}${WINCNI_EXE} -o ${DIST_DIR}/${WINCNI_EXE}
	curl -L ${WINSDN_URL}${HNS_PSM1} -o ${DIST_DIR}/${HNS_PSM1}
}

copy_dockerfile() {
  cp ${AKS_ENGINE_HOME}/windows/Dockerfile ${DIST_DIR}
}

create_zip() {
	ZIP_NAME="${k8s_e2e_upstream_version:-"v${AKS_VERSION}int.zip"}"
	cd ${DIST_DIR}/..
	zip -r ../${ZIP_NAME} k/*
	cd -
}

upload_zip_to_blob_storage() {
	az storage blob upload -f ${TOP_DIR}/../v${AKS_VERSION}int.zip -c ${AZURE_STORAGE_CONTAINER_NAME} -n v${AKS_VERSION}int.zip
}

push_aks_branch() {
	if version_lt "${KUBERNETES_RELEASE}" "1.9"; then
		echo "push to azure repo..."
		cd ${GOPATH}/src/k8s.io/kubernetes
		git push origin ${AKS_BRANCH_NAME}
	else
		echo "no need to push to azure repo"
	fi
}

cleanup_output() {
	rm ${TOP_DIR}/../v${AKS_VERSION}int.zip
	rm -r ${TOP_DIR}
}


AKS_ENGINE_HOME=${GOPATH}/src/github.com/Azure/aks-engine

usage() {
	echo "$0 [-v version] [-p aks_patch_version]"
	echo " -v <version>: version"
	echo " -p <patched version>: aks_patch_version"
	echo " -u <version build for kubernetes upstream e2e tests>: k8s_e2e_upstream_version"
	echo " -z <zip path>: zip_path"
}

while getopts ":v:p:u:z:" opt; do
  case ${opt} in
    v)
      version=${OPTARG}
      ;;
    p)
      aks_patch_version=${OPTARG}
      ;;
	u)
	  k8s_e2e_upstream_version=${OPTARG}
	  ;;
	z)
	  zip_path=${OPTARG}
	  ;;  
    *)
			usage
			exit
      ;;
  esac
done

KUBEPATH=${GOPATH}/src/k8s.io/kubernetes

if [ -z "${k8s_e2e_upstream_version}" ]; then

	if [ -z "${version}" ] || [ -z "${aks_patch_version}" ]; then
		usage
			exit 1
	fi

	if [ -z "${AZURE_STORAGE_CONNECTION_STRING}" ] || [ -z "${AZURE_STORAGE_CONTAINER_NAME}" ]; then
		echo '$AZURE_STORAGE_CONNECTION_STRING and $AZURE_STORAGE_CONTAINER_NAME need to be set for upload to Azure Blob Storage.'
			exit 1
	fi

	KUBERNETES_RELEASE=$(echo $version | cut -d'.' -f1,2)
	KUBERNETES_TAG_BRANCH=v${version}
	AKS_VERSION=${version}-${aks_patch_version}
	AKS_BRANCH_NAME=acs-v${AKS_VERSION}
	TOP_DIR=${AKS_ENGINE_HOME}/_dist/k8s-windows-v${AKS_VERSION}
	DIST_DIR=${TOP_DIR}/k

	create_dist_dir
	get_kube_binaries
	download_nssm
	download_wincni
	copy_dockerfile
	create_zip
	upload_zip_to_blob_storage
	push_aks_branch
	cleanup_output

else
	DIST_DIR=${zip_path}/k
	create_dist_dir
	build_kube_binaries_for_upstream_e2e
	download_nssm
	download_wincni
	create_zip
fi	

