// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

var cseErrorCodes = map[string]int{
	"ERR_SYSTEMCTL_STOP_FAIL":                    3,
	"ERR_SYSTEMCTL_START_FAIL":                   4,
	"ERR_CLOUD_INIT_TIMEOUT":                     5, // Deprecated
	"ERR_FILE_WATCH_TIMEOUT":                     6,
	"ERR_HOLD_WALINUXAGENT":                      7,
	"ERR_RELEASE_HOLD_WALINUXAGENT":              8,
	"ERR_APT_INSTALL_TIMEOUT":                    9,
	"ERR_ETCD_DATA_DIR_NOT_FOUND":                10, // Deprecated
	"ERR_ETCD_RUNNING_TIMEOUT":                   11,
	"ERR_ETCD_DOWNLOAD_TIMEOUT":                  12, // Deprecated
	"ERR_ETCD_VOL_MOUNT_FAIL":                    13,
	"ERR_ETCD_START_TIMEOUT":                     14,
	"ERR_ETCD_CONFIG_FAIL":                       15,
	"ERR_DOCKER_INSTALL_TIMEOUT":                 20, // Deprecated
	"ERR_DOCKER_DOWNLOAD_TIMEOUT":                21, // Deprecated
	"ERR_DOCKER_KEY_DOWNLOAD_TIMEOUT":            22, // Deprecated
	"ERR_DOCKER_APT_KEY_TIMEOUT":                 23, // Deprecated
	"ERR_DOCKER_START_FAIL":                      24,
	"ERR_MOBY_APT_LIST_TIMEOUT":                  25,
	"ERR_MS_GPG_KEY_DOWNLOAD_TIMEOUT":            26,
	"ERR_MOBY_INSTALL_TIMEOUT":                   27,
	"ERR_K8S_RUNNING_TIMEOUT":                    30,
	"ERR_K8S_DOWNLOAD_TIMEOUT":                   31,
	"ERR_KUBECTL_NOT_FOUND":                      32, // Deprecated
	"ERR_IMG_DOWNLOAD_TIMEOUT":                   33,
	"ERR_KUBELET_START_FAIL":                     34,
	"ERR_CONTAINER_IMG_PULL_TIMEOUT":             35,
	"ERR_CNI_DOWNLOAD_TIMEOUT":                   41,
	"ERR_MS_PROD_DEB_DOWNLOAD_TIMEOUT":           42,
	"ERR_MS_PROD_DEB_PKG_ADD_FAIL":               43,
	"ERR_SYSTEMD_INSTALL_FAIL":                   48, // Deprecated
	"ERR_MODPROBE_FAIL":                          49,
	"ERR_OUTBOUND_CONN_FAIL":                     50,
	"ERR_K8S_API_SERVER_CONN_FAIL":               51,
	"ERR_K8S_API_SERVER_DNS_LOOKUP_FAIL":         52,
	"ERR_K8S_API_SERVER_AZURE_DNS_LOOKUP_FAIL":   53,
	"ERR_KATA_KEY_DOWNLOAD_TIMEOUT":              60,
	"ERR_KATA_APT_KEY_TIMEOUT":                   61,
	"ERR_KATA_INSTALL_TIMEOUT":                   62,
	"ERR_CONTAINERD_DOWNLOAD_TIMEOUT":            70, // Deprecated
	"ERR_CUSTOM_SEARCH_DOMAINS_FAIL":             80,
	"ERR_GPU_DRIVERS_START_FAIL":                 84,
	"ERR_GPU_DRIVERS_INSTALL_TIMEOUT":            85,
	"ERR_GPU_DRIVERS_CONFIG":                     86,
	"ERR_SGX_DRIVERS_INSTALL_TIMEOUT":            90,
	"ERR_SGX_DRIVERS_START_FAIL":                 91,
	"ERR_SGX_DRIVERS_NOT_SUPPORTED":              92,
	"ERR_SGX_DRIVERS_CHECKSUM_MISMATCH":          93,
	"ERR_APT_DAILY_TIMEOUT":                      98, // Deprecated
	"ERR_APT_UPDATE_TIMEOUT":                     99,
	"ERR_CSE_PROVISION_SCRIPT_NOT_READY_TIMEOUT": 100,
	"ERR_APT_DIST_UPGRADE_TIMEOUT":               101,
	"ERR_APT_PURGE_FAIL":                         102, // Deprecated
	"ERR_SYSCTL_RELOAD":                          103,
	"ERR_CIS_ASSIGN_ROOT_PW":                     111, // Deprecated
	"ERR_CIS_ASSIGN_FILE_PERMISSION":             112,
	"ERR_PACKER_COPY_FILE":                       113,
	"ERR_CIS_APPLY_PASSWORD_CONFIG":              115,
	"ERR_VHD_FILE_NOT_FOUND":                     124,
	"ERR_VHD_BUILD_ERROR":                        125, // Deprecated
	"ERR_AZURE_STACK_GET_ARM_TOKEN":              120,
	"ERR_AZURE_STACK_GET_NETWORK_CONFIGURATION":  121,
	"ERR_AZURE_STACK_GET_SUBNET_PREFIX":          122,
	"ERR_IOVISOR_KEY_DOWNLOAD_TIMEOUT":           166,
	"ERR_IOVISOR_APT_KEY_TIMEOUT":                167,
	"ERR_BCC_INSTALL_TIMEOUT":                    168,
	"ERR_BPFTRACE_BIN_DOWNLOAD_FAIL":             169,
	"ERR_BPFTRACE_TOOLS_DOWNLOAD_FAIL":           170,
	"ERR_CLUSTER_INIT_FAIL":                      180,
	"ERR_KUBERESERVED_SLICE_SETUP_FAIL":          181,
	"ERR_KUBELET_SLICE_SETUP_FAIL":               182,
	"ERR_CRI_SLICE_SETUP_FAIL":                   183,
}

func GetCSEErrorCode(errorType string) int {
	if code, ok := cseErrorCodes[errorType]; ok {
		return code
	}
	return -1
}
