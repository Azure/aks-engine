// Code generated for package engine by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../parts/agentoutputs.t
// ../../parts/agentparams.t
// ../../parts/dcos/bstrap/bootstrapcustomdata.yml
// ../../parts/dcos/bstrap/bootstrapparams.t
// ../../parts/dcos/bstrap/bootstrapprovision.sh
// ../../parts/dcos/bstrap/bootstrapresources.t
// ../../parts/dcos/bstrap/bootstrapvars.t
// ../../parts/dcos/bstrap/dcos1.11.0.customdata.t
// ../../parts/dcos/bstrap/dcos1.11.2.customdata.t
// ../../parts/dcos/bstrap/dcosbase.t
// ../../parts/dcos/bstrap/dcosmasterresources.t
// ../../parts/dcos/bstrap/dcosmastervars.t
// ../../parts/dcos/bstrap/dcosprovision.sh
// ../../parts/dcos/dcosWindowsAgentResourcesVmas.t
// ../../parts/dcos/dcosWindowsAgentResourcesVmss.t
// ../../parts/dcos/dcosWindowsProvision.ps1
// ../../parts/dcos/dcosagentresourcesvmas.t
// ../../parts/dcos/dcosagentresourcesvmss.t
// ../../parts/dcos/dcosagentvars.t
// ../../parts/dcos/dcosbase.t
// ../../parts/dcos/dcoscustomdata110.t
// ../../parts/dcos/dcoscustomdata184.t
// ../../parts/dcos/dcoscustomdata187.t
// ../../parts/dcos/dcoscustomdata188.t
// ../../parts/dcos/dcoscustomdata190.t
// ../../parts/dcos/dcoscustomdata198.t
// ../../parts/dcos/dcosmasterresources.t
// ../../parts/dcos/dcosmastervars.t
// ../../parts/dcos/dcosparams.t
// ../../parts/dcos/dcosprovision.sh
// ../../parts/dcos/dcosprovisionsource.sh
// ../../parts/iaasoutputs.t
// ../../parts/k8s/addons/1.10/kubernetesmasteraddons-kube-dns-deployment.yaml
// ../../parts/k8s/addons/1.15/kubernetesmasteraddons-azure-cloud-provider-deployment.yaml
// ../../parts/k8s/addons/1.15/kubernetesmasteraddons-pod-security-policy.yaml
// ../../parts/k8s/addons/1.16/kubernetesmaster-audit-policy.yaml
// ../../parts/k8s/addons/1.16/kubernetesmasteraddons-azure-cloud-provider-deployment.yaml
// ../../parts/k8s/addons/1.16/kubernetesmasteraddons-cilium-daemonset.yaml
// ../../parts/k8s/addons/1.16/kubernetesmasteraddons-flannel-daemonset.yaml
// ../../parts/k8s/addons/1.16/kubernetesmasteraddons-kube-dns-deployment.yaml
// ../../parts/k8s/addons/1.16/kubernetesmasteraddons-kube-proxy-daemonset.yaml
// ../../parts/k8s/addons/1.16/kubernetesmasteraddons-pod-security-policy.yaml
// ../../parts/k8s/addons/1.6/kubernetesmasteraddons-kubernetes-dashboard-deployment.yaml
// ../../parts/k8s/addons/1.7/kubernetesmasteraddons-kube-dns-deployment.yaml
// ../../parts/k8s/addons/1.7/kubernetesmasteraddons-kubernetes-dashboard-deployment.yaml
// ../../parts/k8s/addons/1.8/kubernetesmasteraddons-kube-dns-deployment.yaml
// ../../parts/k8s/addons/1.8/kubernetesmasteraddons-kubernetes-dashboard-deployment.yaml
// ../../parts/k8s/addons/1.9/kubernetesmasteraddons-kube-dns-deployment.yaml
// ../../parts/k8s/addons/coredns.yaml
// ../../parts/k8s/addons/kubernetesmaster-audit-policy.yaml
// ../../parts/k8s/addons/kubernetesmasteraddons-aad-default-admin-group-rbac.yaml
// ../../parts/k8s/addons/kubernetesmasteraddons-azure-cloud-provider-deployment.yaml
// ../../parts/k8s/addons/kubernetesmasteraddons-cilium-daemonset.yaml
// ../../parts/k8s/addons/kubernetesmasteraddons-elb-svc.yaml
// ../../parts/k8s/addons/kubernetesmasteraddons-flannel-daemonset.yaml
// ../../parts/k8s/addons/kubernetesmasteraddons-kube-dns-deployment.yaml
// ../../parts/k8s/addons/kubernetesmasteraddons-kube-proxy-daemonset.yaml
// ../../parts/k8s/addons/kubernetesmasteraddons-managed-azure-storage-classes-custom.yaml
// ../../parts/k8s/addons/kubernetesmasteraddons-managed-azure-storage-classes.yaml
// ../../parts/k8s/addons/kubernetesmasteraddons-pod-security-policy.yaml
// ../../parts/k8s/addons/kubernetesmasteraddons-scheduled-maintenance-deployment.yaml
// ../../parts/k8s/addons/kubernetesmasteraddons-unmanaged-azure-storage-classes-custom.yaml
// ../../parts/k8s/addons/kubernetesmasteraddons-unmanaged-azure-storage-classes.yaml
// ../../parts/k8s/armparameters.t
// ../../parts/k8s/cloud-init/artifacts/apt-preferences
// ../../parts/k8s/cloud-init/artifacts/auditd-rules
// ../../parts/k8s/cloud-init/artifacts/cis.sh
// ../../parts/k8s/cloud-init/artifacts/cse_config.sh
// ../../parts/k8s/cloud-init/artifacts/cse_customcloud.sh
// ../../parts/k8s/cloud-init/artifacts/cse_helpers.sh
// ../../parts/k8s/cloud-init/artifacts/cse_install.sh
// ../../parts/k8s/cloud-init/artifacts/cse_main.sh
// ../../parts/k8s/cloud-init/artifacts/default-grub
// ../../parts/k8s/cloud-init/artifacts/dhcpv6.service
// ../../parts/k8s/cloud-init/artifacts/docker-monitor.service
// ../../parts/k8s/cloud-init/artifacts/docker-monitor.timer
// ../../parts/k8s/cloud-init/artifacts/docker_clear_mount_propagation_flags.conf
// ../../parts/k8s/cloud-init/artifacts/enable-dhcpv6.sh
// ../../parts/k8s/cloud-init/artifacts/etc-issue
// ../../parts/k8s/cloud-init/artifacts/etc-issue.net
// ../../parts/k8s/cloud-init/artifacts/etcd.service
// ../../parts/k8s/cloud-init/artifacts/generateproxycerts.sh
// ../../parts/k8s/cloud-init/artifacts/health-monitor.sh
// ../../parts/k8s/cloud-init/artifacts/kms.service
// ../../parts/k8s/cloud-init/artifacts/kubelet-monitor.service
// ../../parts/k8s/cloud-init/artifacts/kubelet-monitor.timer
// ../../parts/k8s/cloud-init/artifacts/kubelet.service
// ../../parts/k8s/cloud-init/artifacts/label-nodes.service
// ../../parts/k8s/cloud-init/artifacts/label-nodes.sh
// ../../parts/k8s/cloud-init/artifacts/modprobe-CIS.conf
// ../../parts/k8s/cloud-init/artifacts/mountetcd.sh
// ../../parts/k8s/cloud-init/artifacts/pam-d-common-auth
// ../../parts/k8s/cloud-init/artifacts/pam-d-common-password
// ../../parts/k8s/cloud-init/artifacts/pam-d-su
// ../../parts/k8s/cloud-init/artifacts/profile-d-cis.sh
// ../../parts/k8s/cloud-init/artifacts/pwquality-CIS.conf
// ../../parts/k8s/cloud-init/artifacts/rsyslog-d-60-CIS.conf
// ../../parts/k8s/cloud-init/artifacts/setup-custom-search-domains.sh
// ../../parts/k8s/cloud-init/artifacts/sshd_config
// ../../parts/k8s/cloud-init/artifacts/sshd_config_1604
// ../../parts/k8s/cloud-init/artifacts/sys-fs-bpf.mount
// ../../parts/k8s/cloud-init/artifacts/sysctl-d-60-CIS.conf
// ../../parts/k8s/cloud-init/jumpboxcustomdata.yml
// ../../parts/k8s/cloud-init/masternodecustomdata.yml
// ../../parts/k8s/cloud-init/nodecustomdata.yml
// ../../parts/k8s/containeraddons/1.16/azure-cni-networkmonitor.yaml
// ../../parts/k8s/containeraddons/1.16/ip-masq-agent.yaml
// ../../parts/k8s/containeraddons/1.16/kubernetesmasteraddons-aad-pod-identity-deployment.yaml
// ../../parts/k8s/containeraddons/1.16/kubernetesmasteraddons-aci-connector-deployment.yaml
// ../../parts/k8s/containeraddons/1.16/kubernetesmasteraddons-azure-npm-daemonset.yaml
// ../../parts/k8s/containeraddons/1.16/kubernetesmasteraddons-blobfuse-flexvolume-installer.yaml
// ../../parts/k8s/containeraddons/1.16/kubernetesmasteraddons-calico-daemonset.yaml
// ../../parts/k8s/containeraddons/1.16/kubernetesmasteraddons-cluster-autoscaler-deployment.yaml
// ../../parts/k8s/containeraddons/1.16/kubernetesmasteraddons-heapster-deployment.yaml
// ../../parts/k8s/containeraddons/1.16/kubernetesmasteraddons-keyvault-flexvolume-installer.yaml
// ../../parts/k8s/containeraddons/1.16/kubernetesmasteraddons-kube-rescheduler-deployment.yaml
// ../../parts/k8s/containeraddons/1.16/kubernetesmasteraddons-kubernetes-dashboard-deployment.yaml
// ../../parts/k8s/containeraddons/1.16/kubernetesmasteraddons-metrics-server-deployment.yaml
// ../../parts/k8s/containeraddons/1.16/kubernetesmasteraddons-nvidia-device-plugin-daemonset.yaml
// ../../parts/k8s/containeraddons/1.16/kubernetesmasteraddons-omsagent-daemonset.yaml
// ../../parts/k8s/containeraddons/1.16/kubernetesmasteraddons-smb-flexvolume-installer.yaml
// ../../parts/k8s/containeraddons/1.16/kubernetesmasteraddons-tiller-deployment.yaml
// ../../parts/k8s/containeraddons/1.6/kubernetesmasteraddons-heapster-deployment.yaml
// ../../parts/k8s/containeraddons/1.7/kubernetesmasteraddons-heapster-deployment.yaml
// ../../parts/k8s/containeraddons/1.8/kubernetesmasteraddons-heapster-deployment.yaml
// ../../parts/k8s/containeraddons/azure-cni-networkmonitor.yaml
// ../../parts/k8s/containeraddons/dns-autoscaler.yaml
// ../../parts/k8s/containeraddons/ip-masq-agent.yaml
// ../../parts/k8s/containeraddons/kubernetesmasteraddons-aad-pod-identity-deployment.yaml
// ../../parts/k8s/containeraddons/kubernetesmasteraddons-aci-connector-deployment.yaml
// ../../parts/k8s/containeraddons/kubernetesmasteraddons-azure-npm-daemonset.yaml
// ../../parts/k8s/containeraddons/kubernetesmasteraddons-blobfuse-flexvolume-installer.yaml
// ../../parts/k8s/containeraddons/kubernetesmasteraddons-calico-daemonset.yaml
// ../../parts/k8s/containeraddons/kubernetesmasteraddons-cluster-autoscaler-deployment.yaml
// ../../parts/k8s/containeraddons/kubernetesmasteraddons-heapster-deployment.yaml
// ../../parts/k8s/containeraddons/kubernetesmasteraddons-keyvault-flexvolume-installer.yaml
// ../../parts/k8s/containeraddons/kubernetesmasteraddons-kube-rescheduler-deployment.yaml
// ../../parts/k8s/containeraddons/kubernetesmasteraddons-kubernetes-dashboard-deployment.yaml
// ../../parts/k8s/containeraddons/kubernetesmasteraddons-metrics-server-deployment.yaml
// ../../parts/k8s/containeraddons/kubernetesmasteraddons-nvidia-device-plugin-daemonset.yaml
// ../../parts/k8s/containeraddons/kubernetesmasteraddons-omsagent-daemonset.yaml
// ../../parts/k8s/containeraddons/kubernetesmasteraddons-smb-flexvolume-installer.yaml
// ../../parts/k8s/containeraddons/kubernetesmasteraddons-tiller-deployment.yaml
// ../../parts/k8s/kubeconfig.json
// ../../parts/k8s/kubernetesparams.t
// ../../parts/k8s/kuberneteswindowsfunctions.ps1
// ../../parts/k8s/kuberneteswindowssetup.ps1
// ../../parts/k8s/manifests/kubernetesmaster-cloud-controller-manager.yaml
// ../../parts/k8s/manifests/kubernetesmaster-kube-addon-manager.yaml
// ../../parts/k8s/manifests/kubernetesmaster-kube-apiserver.yaml
// ../../parts/k8s/manifests/kubernetesmaster-kube-controller-manager-custom.yaml
// ../../parts/k8s/manifests/kubernetesmaster-kube-controller-manager.yaml
// ../../parts/k8s/manifests/kubernetesmaster-kube-scheduler.yaml
// ../../parts/k8s/windowsazurecnifunc.ps1
// ../../parts/k8s/windowscnifunc.ps1
// ../../parts/k8s/windowsconfigfunc.ps1
// ../../parts/k8s/windowsinstallopensshfunc.ps1
// ../../parts/k8s/windowskubeletfunc.ps1
// ../../parts/masteroutputs.t
// ../../parts/masterparams.t
// ../../parts/swarm/Install-ContainerHost-And-Join-Swarm.ps1
// ../../parts/swarm/Join-SwarmMode-cluster.ps1
// ../../parts/swarm/configure-swarm-cluster.sh
// ../../parts/swarm/configure-swarmmode-cluster.sh
// ../../parts/swarm/swarmagentresourcesvmas.t
// ../../parts/swarm/swarmagentresourcesvmss.t
// ../../parts/swarm/swarmagentvars.t
// ../../parts/swarm/swarmbase.t
// ../../parts/swarm/swarmmasterresources.t
// ../../parts/swarm/swarmmastervars.t
// ../../parts/swarm/swarmparams.t
// ../../parts/swarm/swarmwinagentresourcesvmas.t
// ../../parts/swarm/swarmwinagentresourcesvmss.t
// ../../parts/windowsparams.t
package engine

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _agentoutputsT = []byte(`{{if IsPublic .Ports}}
  {{ if not IsKubernetes }}
    "{{.Name}}FQDN": {
        "type": "string",
        "value": "[reference(concat('Microsoft.Network/publicIPAddresses/', variables('{{.Name}}IPAddressName'))).dnsSettings.fqdn]"
    },
  {{end}}
{{end}}
{{if and .IsAvailabilitySets .IsStorageAccount}}
  "{{.Name}}StorageAccountOffset": {
      "type": "int",
      "value": "[variables('{{.Name}}StorageAccountOffset')]"
    },
    "{{.Name}}StorageAccountCount": {
      "type": "int",
      "value": "[variables('{{.Name}}StorageAccountsCount')]"
    },
    "{{.Name}}SubnetName": {
      "type": "string",
      "value": "[variables('{{.Name}}SubnetName')]"
    },
{{end}}`)

func agentoutputsTBytes() ([]byte, error) {
	return _agentoutputsT, nil
}

func agentoutputsT() (*asset, error) {
	bytes, err := agentoutputsTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "agentoutputs.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _agentparamsT = []byte(`    "{{.Name}}Count": {
      "defaultValue": {{.Count}},
      "metadata": {
        "description": "The number of vms in agent pool {{.Name}}"
      },
      "type": "int"
    },
{{if .IsAvailabilitySets}}
    "{{.Name}}Offset": {
      "defaultValue": 0,
      "metadata": {
        "description": "offset to a particular vm within a VMAS agent pool"
      },
      "type": "int"
    },
{{end}}
    {{if .IsLowPriorityScaleSet}}
    "{{.Name}}ScaleSetPriority": {
      "allowedValues":[
        "Low",
        "Regular",
        ""
      ],
      "defaultValue": "{{.ScaleSetPriority}}",
      "metadata": {
        "description": "The priority for the VM Scale Set. This value can be Low or Regular."
      },
      "type": "string"
    },
    "{{.Name}}ScaleSetEvictionPolicy": {
      "allowedValues":[
        "Delete",
        "Deallocate",
        ""
      ],
      "defaultValue": "{{.ScaleSetEvictionPolicy}}",
      "metadata": {
        "description": "The Eviction Policy for a Low-priority VM Scale Set."
      },
      "type": "string"
    },
    {{end}}
    "{{.Name}}VMSize": {
      {{GetKubernetesAllowedVMSKUs}}
      "defaultValue": "{{.VMSize}}",
      "metadata": {
        "description": "The size of the Virtual Machine."
      },
      "type": "string"
    },
{{if HasAvailabilityZones .}}
    "{{.Name}}AvailabilityZones": {
      "metadata": {
        "description": "Agent availability zones"
      },
      "type": "array"
    },
{{end}}
    "{{.Name}}osImageName": {
      "defaultValue": "",
      "metadata": {
        "description": "Name of a {{.OSType}} OS image. Needs to be used in conjuction with osImageResourceGroup."
      },
      "type": "string"
    },
    "{{.Name}}osImageResourceGroup": {
      "defaultValue": "",
      "metadata": {
        "description": "Resource group of a {{.OSType}} OS image. Needs to be used in conjuction with osImageName."
      },
      "type": "string"
    },
    "{{.Name}}osImageOffer": {
      "defaultValue": "UbuntuServer",
      "metadata": {
        "description": "{{.OSType}} OS image type."
      },
      "type": "string"
    },
    "{{.Name}}osImagePublisher": {
      "defaultValue": "Canonical",
      "metadata": {
        "description": "OS image publisher."
      },
      "type": "string"
    },
    "{{.Name}}osImageSKU": {
      "defaultValue": "16.04-LTS",
      "metadata": {
        "description": "OS image SKU."
      },
      "type": "string"
    },
    "{{.Name}}osImageVersion": {
      "defaultValue": "latest",
      "metadata": {
        "description": "OS image version."
      },
      "type": "string"
    },
{{if .IsCustomVNET}}
    "{{.Name}}VnetSubnetID": {
      "metadata": {
        "description": "Sets the vnet subnet of agent pool '{{.Name}}'."
      },
      "type": "string"
    }
{{else}}
    "{{.Name}}Subnet": {
      "defaultValue": "{{.Subnet}}",
      "metadata": {
        "description": "Sets the subnet of agent pool '{{.Name}}'."
      },
      "type": "string"
    }
{{end}}
{{if IsPublic .Ports}}
  ,"{{.Name}}EndpointDNSNamePrefix": {
      "metadata": {
        "description": "Sets the Domain name label for the agent pool IP Address.  The concatenation of the domain name label and the regional DNS zone make up the fully qualified domain name associated with the public IP address."
      },
      "type": "string"
    }
{{end}}
{{if HasPrivateRegistry}}
  ,"registry": {
      "metadata": {
        "description": "Private Container Registry"
      },
      "type": "string"
    },
  "registryKey": {
      "metadata": {
        "description": "base64 encoded key to the Private Container Registry"
      },
      "type": "string"
    }
  {{end}}
`)

func agentparamsTBytes() ([]byte, error) {
	return _agentparamsT, nil
}

func agentparamsT() (*asset, error) {
	bytes, err := agentparamsTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "agentparams.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosBstrapBootstrapcustomdataYml = []byte(`bootcmd:
- bash -c "if [ ! -f /var/lib/sdb-gpt ];then echo DCOS-5890;parted -s /dev/sdb mklabel
  gpt;touch /var/lib/sdb-gpt;fi"
disk_setup:
  ephemeral0:
    layout:
    - 50
    - 50
    overwrite: true
    table_type: gpt
fs_setup:
- device: ephemeral0.1
  filesystem: ext4
  overwrite: true
- device: ephemeral0.2
  filesystem: ext4
  overwrite: true
mounts:
- - ephemeral0.1
  - /var/lib/mesos
- - ephemeral0.2
  - /var/lib/docker
runcmd:
    - [ ln, -s, /bin/rm, /usr/bin/rm ]
    - [ ln, -s, /bin/mkdir, /usr/bin/mkdir ]
    - [ ln, -s, /bin/tar, /usr/bin/tar ]
    - [ ln, -s, /bin/ln, /usr/bin/ln ]
    - [ ln, -s, /bin/cp, /usr/bin/cp ]
    - [ ln, -s, /bin/systemctl, /usr/bin/systemctl ]
    - [ ln, -s, /bin/mount, /usr/bin/mount ]
    - [ ln, -s, /bin/bash, /usr/bin/bash ]
    - [ ln, -s, /usr/sbin/useradd, /usr/bin/useradd ]
    - /opt/azure/containers/provision.sh
    - /opt/azure/dcos/init_bootstrap.sh
write_files:
- content: |
    DCOS_ENVIRONMENT={{{targetEnvironment}}}
  owner: root
  path: /opt/azure/dcos/environment
  permissions: '0644'
- content: |
    #!/bin/sh

    curl -H Metadata:true -fsSL "http://169.254.169.254/metadata/instance/network/interface/0/ipv4/ipAddress/0/privateIpAddress?api-version=2017-04-02&format=text"
  owner: root
  path: /opt/azure/dcos/genconf/ip-detect
  permissions: '0755'
- content: |
    bootstrap_url: http://BOOTSTRAP_IP:8086
    cluster_name: azure-dcos
    exhibitor_storage_backend: static
    master_discovery: static
    oauth_enabled: BOOTSTRAP_OAUTH_ENABLED
    ip_detect_public_filename: genconf/ip-detect
    master_list:
MASTER_IP_LIST
    resolvers:
    - 168.63.129.16
    dns_search:
  owner: root
  path: /opt/azure/dcos/genconf/config.yaml
  permissions: '0644'
- content: |
    #!/bin/bash

    source /opt/azure/containers/provision_source.sh

    # update dns_search
    dns=$(grep search /etc/resolv.conf | cut -d " " -f 2)
    sed -i "/dns_search:/c dns_search: $dns" /opt/azure/dcos/genconf/config.yaml
    # install and run bootstrap package
    cd /opt/azure/dcos
    retrycmd_if_failure 10 10 120 curl -fsSL -o dcos_generate_config.sh.sha1sum {{{dcosBootstrapURL}}}.sha1sum
    retry_download 1 1 120 {{{dcosBootstrapURL}}} dcos_generate_config.sh $(cat dcos_generate_config.sh.sha1sum)
    bash dcos_generate_config.sh
    docker run -d -p 8086:80 -v $PWD/genconf/serve:/usr/share/nginx/html:ro nginx
  owner: root
  path: /opt/azure/dcos/init_bootstrap.sh
  permissions: '0755'
- content: 'PROVISION_SOURCE_STR'
  path: /opt/azure/containers/provision_source.sh
  permissions: "0744"
  owner: "root"
- content: 'PROVISION_STR'
  path: /opt/azure/containers/provision.sh
  permissions: "0744"
  owner: "root"
`)

func dcosBstrapBootstrapcustomdataYmlBytes() ([]byte, error) {
	return _dcosBstrapBootstrapcustomdataYml, nil
}

func dcosBstrapBootstrapcustomdataYml() (*asset, error) {
	bytes, err := dcosBstrapBootstrapcustomdataYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/bstrap/bootstrapcustomdata.yml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosBstrapBootstrapparamsT = []byte(`    "linuxAdminUsername": {
      "metadata": {
        "description": "User name for the Linux Virtual Machines (SSH or Password)."
      },
      "type": "string"
    },
    {{range .ExtensionProfiles}}
      "{{.Name}}Parameters": {
        "metadata": {
        "description": "Parameters for the extension"
      },
      "type": "securestring"
      },
    {{end}}
    "bootstrapStaticIP": {
      "metadata": {
        "description": "Sets the static IP of the first bootstrap"
      },
      "type": "string"
    },
    "bootstrapVMSize": {
      {{GetMasterAllowedSizes}}
      "metadata": {
        "description": "The size of the Virtual Machine."
      },
      "type": "string"
    },
    "sshRSAPublicKey": {
      "metadata": {
        "description": "SSH public key used for auth to all Linux machines.  Not Required.  If not set, you must provide a password key."
      },
      "type": "string"
    },
    "nameSuffix": {
      "defaultValue": "{{GetUniqueNameSuffix}}",
      "metadata": {
        "description": "A string hash of the bootstrap DNS name to uniquely identify the cluster."
      },
      "type": "string"
    },
    "osImageName": {
      "defaultValue": "",
      "metadata": {
        "description": "Name of a Linux OS image. Needs to be used in conjuction with osImageResourceGroup."
      },
      "type": "string"
    },
    "osImageResourceGroup": {
      "defaultValue": "",
      "metadata": {
        "description": "Resource group of a Linux OS image. Needs to be used in conjuction with osImageName."
      },
      "type": "string"
    },
    "osImageOffer": {
      "defaultValue": "UbuntuServer",
      "metadata": {
        "description": "Linux OS image type."
      },
      "type": "string"
    },
    "osImagePublisher": {
      "defaultValue": "Canonical",
      "metadata": {
        "description": "OS image publisher."
      },
      "type": "string"
    },
    "osImageSKU": {
      "defaultValue": "16.04-LTS",
      "metadata": {
        "description": "OS image SKU."
      },
      "type": "string"
    },
    "osImageVersion": {
      "defaultValue": "16.04.201804050",
      "metadata": {
        "description": "OS image version."
      },
      "type": "string"
    },
    "fqdnEndpointSuffix":{
      "defaultValue": "cloudapp.azure.com",
      "metadata": {
        "description": "Endpoint of FQDN."
      },
      "type": "string"
    },
    "targetEnvironment": {
      "defaultValue": "AzurePublicCloud",
      "metadata": {
        "description": "The azure deploy environment. Currently support: AzurePublicCloud, AzureChinaCloud"
      },
      "type": "string"
    },
    "location": {
      "defaultValue": "{{GetLocation}}",
      "metadata": {
        "description": "Sets the location for all resources in the cluster"
      },
      "type": "string"
    }
{{if .LinuxProfile.HasSecrets}}
  {{range  $vIndex, $vault := .LinuxProfile.Secrets}}
    ,
    "linuxKeyVaultID{{$vIndex}}": {
      "metadata": {
        "description": "KeyVaultId{{$vIndex}} to install certificates from on linux machines."
      },
      "type": "string"
    }
    {{range $cIndex, $cert := $vault.VaultCertificates}}
      ,
      "linuxKeyVaultID{{$vIndex}}CertificateURL{{$cIndex}}": {
        "metadata": {
          "description": "CertificateURL{{$cIndex}} to install from KeyVaultId{{$vIndex}} on linux machines."
        },
        "type": "string"
      }
    {{end}}
  {{end}}
{{end}}
{{if .HasWindows}}{{if .WindowsProfile.HasSecrets}}
  {{range  $vIndex, $vault := .WindowsProfile.Secrets}}
    ,
    "windowsKeyVaultID{{$vIndex}}": {
      "metadata": {
        "description": "KeyVaultId{{$vIndex}} to install certificates from on windows machines."
      },
      "type": "string"
    }
    {{range $cIndex, $cert := $vault.VaultCertificates}}
      ,
      "windowsKeyVaultID{{$vIndex}}CertificateURL{{$cIndex}}": {
        "metadata": {
          "description": "Url to retrieve Certificate{{$cIndex}} from KeyVaultId{{$vIndex}} to install on windows machines."
        },
        "type": "string"
      },
      "windowsKeyVaultID{{$vIndex}}CertificateStore{{$cIndex}}": {
        "metadata": {
          "description": "CertificateStore to install Certificate{{$cIndex}} from KeyVaultId{{$vIndex}} on windows machines."
        },
        "type": "string"
      }
    {{end}}
  {{end}}
{{end}} {{end}}
`)

func dcosBstrapBootstrapparamsTBytes() ([]byte, error) {
	return _dcosBstrapBootstrapparamsT, nil
}

func dcosBstrapBootstrapparamsT() (*asset, error) {
	bytes, err := dcosBstrapBootstrapparamsTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/bstrap/bootstrapparams.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosBstrapBootstrapprovisionSh = []byte(`#!/bin/bash

source /opt/azure/containers/provision_source.sh
source /opt/azure/dcos/environment

# default dc/os component download address (Azure CDN)
packages=(
  https://dcos-mirror.azureedge.net/pkg/libltdl7_2.4.6-0.1_amd64.deb
  https://dcos-mirror.azureedge.net/pkg/docker-ce_17.09.0~ce-0~ubuntu_amd64.deb
)

# sha1sum checksums for @packages
sha1sums=(
  9a0f9f2769d3dc834737aa7df50aaaea369af98d
  94f6e89be6d45d9988269a237eb27c7d6a844d7f
)

case $DCOS_ENVIRONMENT in
  # because of Chinese GreatWall Firewall, the default packages on Azure CDN is blocked. So the following Chinese local mirror url should be used instead.
  AzureChinaCloud)
    packages=(
      http://acsengine.blob.core.chinacloudapi.cn/dcos/libltdl7_2.4.6-0.1_amd64.deb
      http://mirror.kaiyuanshe.cn/docker-ce/linux/ubuntu/dists/xenial/pool/stable/amd64/docker-ce_17.09.0~ce-0~ubuntu_amd64.deb
    )
    ;;
esac

len=$((${#packages[@]}-1))
for i in $(seq 0 $len); do
  retry_get_install_deb 10 10 120 ${packages[$i]} ${sha1sums[$i]}
    if [ $? -ne 0  ]; then
    exit 1
  fi
done
`)

func dcosBstrapBootstrapprovisionShBytes() ([]byte, error) {
	return _dcosBstrapBootstrapprovisionSh, nil
}

func dcosBstrapBootstrapprovisionSh() (*asset, error) {
	bytes, err := dcosBstrapBootstrapprovisionShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/bstrap/bootstrapprovision.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosBstrapBootstrapresourcesT = []byte(`    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('bootstrapNSGName')]",
      "properties": {
        "securityRules": [
            {
                "properties": {
                    "priority": 200,
                    "access": "Allow",
                    "direction": "Inbound",
                    "destinationPortRange": "22",
                    "sourcePortRange": "*",
                    "destinationAddressPrefix": "*",
                    "protocol": "Tcp",
                    "description": "Allow SSH",
                    "sourceAddressPrefix": "*"
                },
                "name": "ssh"
            },
            {
                "properties": {
                    "priority": 201,
                    "access": "Allow",
                    "direction": "Inbound",
                    "destinationPortRange": "8086",
                    "sourcePortRange": "*",
                    "destinationAddressPrefix": "*",
                    "protocol": "Tcp",
                    "description": "Allow bootstrap service",
                    "sourceAddressPrefix": "*"
                },
                "name": "Port8086"
            }
        ]
      },
      "type": "Microsoft.Network/networkSecurityGroups"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "dependsOn": [
{{if not .MasterProfile.IsCustomVNET}}
        "[variables('vnetID')]",
{{end}}
        "[variables('bootstrapNSGID')]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('bootstrapVMName'), '-nic')]",
      "properties": {
        "ipConfigurations": [
          {
            "name": "ipConfigNode",
            "properties": {
              "privateIPAddress": "[variables('bootstrapStaticIP')]",
              "privateIPAllocationMethod": "Static",
              "subnet": {
                "id": "[variables('masterVnetSubnetID')]"
              }
            }
          }
        ],
        "networkSecurityGroup": {
          "id": "[variables('bootstrapNSGID')]"
        }
      },
      "type": "Microsoft.Network/networkInterfaces"
    },
    {
      "apiVersion": "[variables('apiVersionStorageManagedDisks')]",
      "dependsOn": [
        "[concat('Microsoft.Network/networkInterfaces/', variables('bootstrapVMName'), '-nic')]",
{{if .MasterProfile.IsStorageAccount}}
        "[variables('masterStorageAccountName')]",
{{end}}
        "[variables('masterStorageAccountExhibitorName')]"
      ],
      "tags":
      {
        "creationSource": "[concat('acsengine-', variables('bootstrapVMName'))]",
        "orchestratorName": "dcos",
        "orchestratorVersion": "[variables('orchestratorVersion')]",
        "orchestratorNode": "bootstrap"
      },
      "location": "[variables('location')]",
      "name": "[variables('bootstrapVMName')]",
      "properties": {
        "hardwareProfile": {
          "vmSize": "[variables('bootstrapVMSize')]"
        },
        "networkProfile": {
          "networkInterfaces": [
            {
              "id": "[resourceId('Microsoft.Network/networkInterfaces',concat(variables('bootstrapVMName'), '-nic'))]"
            }
          ]
        },
        "osProfile": {
          "adminUsername": "[variables('adminUsername')]",
          "computername": "[variables('bootstrapVMName')]",
          {{GetDCOSBootstrapCustomData}}
          "linuxConfiguration": {
            "disablePasswordAuthentication": true,
            "ssh": {
                "publicKeys": [
                    {
                        "keyData": "[variables('sshRSAPublicKey')]",
                        "path": "[variables('sshKeyPath')]"
                    }
                ]
            }
          }
          {{if .LinuxProfile.HasSecrets}}
          ,
          "secrets": "[variables('linuxProfileSecrets')]"
          {{end}}
        },
        "storageProfile": {
          "imageReference": {
            "offer": "[variables('osImageOffer')]",
            "publisher": "[variables('osImagePublisher')]",
            "sku": "[variables('osImageSKU')]",
            "version": "[variables('osImageVersion')]"
          },
          "osDisk": {
            "caching": "ReadWrite"
            ,"createOption": "FromImage"
{{if .MasterProfile.IsStorageAccount}}
            ,"name": "[concat(variables('bootstrapVMName'), '-osdisk')]"
            ,"vhd": {
              "uri": "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('masterStorageAccountName')),variables('apiVersionStorage')).primaryEndpoints.blob,'vhds/',variables('bootstrapVMName'),-osdisk.vhd')]"
            }
{{end}}
{{if ne .OrchestratorProfile.DcosConfig.BootstrapProfile.OSDiskSizeGB 0}}
            ,"diskSizeGB": "60"
{{end}}
          }
        }
      },
      "type": "Microsoft.Compute/virtualMachines"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "dependsOn": [
        "[concat('Microsoft.Compute/virtualMachines/', variables('bootstrapVMName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('bootstrapVMName'), '/bootstrapready')]",
      "properties": {
        "autoUpgradeMinorVersion": true,
        "publisher": "Microsoft.OSTCExtensions",
        "settings": {
          "commandToExecute": "[concat('/bin/bash -c \"until curl -f http://', variables('bootstrapStaticIP'), ':8086/dcos_install.sh > /dev/null; do echo waiting for bootstrap node; sleep 15; done; echo bootstrap node up\"')]"
        },
        "type": "CustomScriptForLinux",
        "typeHandlerVersion": "1.4"
      },
      "type": "Microsoft.Compute/virtualMachines/extensions"
    }{{WriteLinkedTemplatesForExtensions}}
`)

func dcosBstrapBootstrapresourcesTBytes() ([]byte, error) {
	return _dcosBstrapBootstrapresourcesT, nil
}

func dcosBstrapBootstrapresourcesT() (*asset, error) {
	bytes, err := dcosBstrapBootstrapresourcesTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/bstrap/bootstrapresources.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosBstrapBootstrapvarsT = []byte(`{{if .OrchestratorProfile.DcosConfig.BootstrapProfile}}
    ,
    "dcosBootstrapURL": "[parameters('dcosBootstrapURL')]",
    "bootstrapVMSize": "[parameters('bootstrapVMSize')]",
    "bootstrapNSGID": "[resourceId('Microsoft.Network/networkSecurityGroups',variables('bootstrapNSGName'))]",
    "bootstrapNSGName": "[concat('bootstrap-nsg-', variables('nameSuffix'))]",
    "bootstrapVMName": "[concat('bootstrap-', variables('nameSuffix'))]",
    "bootstrapStaticIP": "[parameters('bootstrapStaticIP')]"
{{end}}
`)

func dcosBstrapBootstrapvarsTBytes() ([]byte, error) {
	return _dcosBstrapBootstrapvarsT, nil
}

func dcosBstrapBootstrapvarsT() (*asset, error) {
	bytes, err := dcosBstrapBootstrapvarsTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/bstrap/bootstrapvars.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosBstrapDcos1110CustomdataT = []byte(`bootcmd:
- bash -c "if [ ! -f /var/lib/sdb-gpt ];then echo DCOS-5890;parted -s /dev/sdb mklabel
  gpt;touch /var/lib/sdb-gpt;fi"
disk_setup:
  ephemeral0:
    layout:
    - 50
    - 50
    overwrite: true
    table_type: gpt
fs_setup:
- device: ephemeral0.1
  filesystem: ext4
  overwrite: true
- device: ephemeral0.2
  filesystem: ext4
  overwrite: true
mounts:
- - ephemeral0.1
  - /var/lib/mesos
- - ephemeral0.2
  - /var/lib/docker
runcmd: PREPROVISION_EXTENSION
- ln -s /bin/rm /usr/bin/rm
- ln -s /bin/mkdir /usr/bin/mkdir
- ln -s /bin/tar /usr/bin/tar
- ln -s /bin/ln /usr/bin/ln
- ln -s /bin/cp /usr/bin/cp
- ln -s /bin/systemctl /usr/bin/systemctl
- ln -s /bin/mount /usr/bin/mount
- ln -s /bin/bash /usr/bin/bash
- ln -s /usr/sbin/useradd /usr/bin/useradd
- systemctl disable --now resolvconf.service
- systemctl mask --now lxc-net.service
- systemctl disable --now unscd.service
- systemctl stop --now unscd.service
- /opt/azure/containers/provision.sh
- bash /tmp/dcos/dcos_install.sh ROLENAME
- /opt/azure/dcos/postinstall-cond.sh
- bash /opt/azure/dcos/diagnostics_fix.sh
write_files:
- content: |
    [Service]
    Restart=always
    StartLimitInterval=0
    RestartSec=15
    ExecStartPre=-/sbin/ip link del docker0
    ExecStart=
    ExecStart=/usr/bin/dockerd --storage-driver=overlay
  path: /etc/systemd/system/docker.service.d/execstart.conf
  permissions: '0644'
- content: |
    [Unit]
    PartOf=docker.service
    [Socket]
    ListenStream=/var/run/docker.sock
    SocketMode=0660
    SocketUser=root
    SocketGroup=docker
    ListenStream=2375
    BindIPv6Only=both
    [Install]
    WantedBy=sockets.target
  path: /etc/systemd/system/docker.socket
  permissions: '0644'
- content: |
    DCOS_ENVIRONMENT={{{targetEnvironment}}}
  owner: root
  path: /opt/azure/dcos/environment
  permissions: '0644'
- content: |
    #!/bin/bash

    for f in /opt/mesosphere/packages/dcos-config--setup_*/etc/dcos-diagnostics-runner-config.json; do
      if [ -e $f ]; then
        sed -i.bak "99 s/1s/10s/" $f
      fi
    done
  owner: root
  path: /opt/azure/dcos/diagnostics_fix.sh
  permissions: '0744'
- path: /var/lib/dcos/mesos-slave-common
  content: 'ATTRIBUTES_STR'
  permissions: "0644"
  owner: "root"
- content: 'PROVISION_SOURCE_STR'
  path: /opt/azure/containers/provision_source.sh
  permissions: "0744"
  owner: "root"
- content: 'PROVISION_STR'
  path: /opt/azure/containers/provision.sh
  permissions: "0744"
  owner: "root"
- content: |
    #!/bin/bash
    if [ -f /opt/azure/dcos/postinstall.sh ]; then /opt/azure/dcos/postinstall.sh; fi
  path: /opt/azure/dcos/postinstall-cond.sh
  permissions: "0744"
  owner: "root"
`)

func dcosBstrapDcos1110CustomdataTBytes() ([]byte, error) {
	return _dcosBstrapDcos1110CustomdataT, nil
}

func dcosBstrapDcos1110CustomdataT() (*asset, error) {
	bytes, err := dcosBstrapDcos1110CustomdataTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/bstrap/dcos1.11.0.customdata.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosBstrapDcos1112CustomdataT = []byte(`bootcmd:
- bash -c "if [ ! -f /var/lib/sdb-gpt ];then echo DCOS-5890;parted -s /dev/sdb mklabel
  gpt;touch /var/lib/sdb-gpt;fi"
disk_setup:
  ephemeral0:
    layout:
    - 50
    - 50
    overwrite: true
    table_type: gpt
fs_setup:
- device: ephemeral0.1
  filesystem: ext4
  overwrite: true
- device: ephemeral0.2
  filesystem: ext4
  overwrite: true
mounts:
- - ephemeral0.1
  - /var/lib/mesos
- - ephemeral0.2
  - /var/lib/docker
runcmd: PREPROVISION_EXTENSION
- ln -s /bin/rm /usr/bin/rm
- ln -s /bin/mkdir /usr/bin/mkdir
- ln -s /bin/tar /usr/bin/tar
- ln -s /bin/ln /usr/bin/ln
- ln -s /bin/cp /usr/bin/cp
- ln -s /bin/systemctl /usr/bin/systemctl
- ln -s /bin/mount /usr/bin/mount
- ln -s /bin/bash /usr/bin/bash
- ln -s /usr/sbin/useradd /usr/bin/useradd
- systemctl disable --now resolvconf.service
- systemctl mask --now lxc-net.service
- systemctl disable --now unscd.service
- systemctl stop --now unscd.service
- /opt/azure/containers/provision.sh
- bash /tmp/dcos/dcos_install.sh ROLENAME
- /opt/azure/dcos/postinstall-cond.sh
write_files:
- content: |
    [Service]
    Restart=always
    StartLimitInterval=0
    RestartSec=15
    ExecStartPre=-/sbin/ip link del docker0
    ExecStart=
    ExecStart=/usr/bin/dockerd --storage-driver=overlay
  path: /etc/systemd/system/docker.service.d/execstart.conf
  permissions: '0644'
- content: |
    [Unit]
    PartOf=docker.service
    [Socket]
    ListenStream=/var/run/docker.sock
    SocketMode=0660
    SocketUser=root
    SocketGroup=docker
    ListenStream=2375
    BindIPv6Only=both
    [Install]
    WantedBy=sockets.target
  path: /etc/systemd/system/docker.socket
  permissions: '0644'
- content: |
    DCOS_ENVIRONMENT={{{targetEnvironment}}}
  owner: root
  path: /opt/azure/dcos/environment
  permissions: '0644'
- path: /var/lib/dcos/mesos-slave-common
  content: 'ATTRIBUTES_STR'
  permissions: "0644"
  owner: "root"
- content: 'PROVISION_SOURCE_STR'
  path: /opt/azure/containers/provision_source.sh
  permissions: "0744"
  owner: "root"
- content: 'PROVISION_STR'
  path: /opt/azure/containers/provision.sh
  permissions: "0744"
  owner: "root"
- content: |
    #!/bin/bash
    if [ -f /opt/azure/dcos/postinstall.sh ]; then /opt/azure/dcos/postinstall.sh; fi
  path: /opt/azure/dcos/postinstall-cond.sh
  permissions: "0744"
  owner: "root"
`)

func dcosBstrapDcos1112CustomdataTBytes() ([]byte, error) {
	return _dcosBstrapDcos1112CustomdataT, nil
}

func dcosBstrapDcos1112CustomdataT() (*asset, error) {
	bytes, err := dcosBstrapDcos1112CustomdataTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/bstrap/dcos1.11.2.customdata.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosBstrapDcosbaseT = []byte(`{
  "$schema": "https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
  "contentVersion": "1.0.0.0",
  "parameters": {
    {{range .AgentPoolProfiles}}{{template "agentparams.t" .}},{{end}}
    {{if .HasWindows}}
      "dcosBinariesURL": {
        "metadata": {
          "description": "The download url for dcos/mesos windows binaries."
        },
        "type": "string"
      },
      "dcosBinariesVersion": {
        "metadata": {
          "description": "DCOS windows binaries version"
        },
        "type": "string"
      },
      {{template "windowsparams.t"}},
    {{end}}
    {{template "dcos/dcosparams.t" .}}
    {{template "dcos/bstrap/bootstrapparams.t" .}},
    {{template "masterparams.t" .}}
  },
  "variables": {
    {{range $index, $agent := .AgentPoolProfiles}}
        "{{.Name}}Index": {{$index}},
        {{template "dcos/dcosagentvars.t" .}}
        {{if .IsStorageAccount}}
          "{{.Name}}StorageAccountOffset": "[mul(variables('maxStorageAccountsPerAgent'),{{$index}})]",
          "{{.Name}}AccountName": "[concat(variables('storageAccountBaseName'), 'agnt{{$index}}')]",
          {{if .HasDisks}}
            "{{.Name}}DataAccountName": "[concat(variables('storageAccountBaseName'), 'data{{$index}}')]",
          {{end}}
        {{end}}
    {{end}}

    {{template "dcos/bstrap/dcosmastervars.t" .}}
    {{template "dcos/bstrap/bootstrapvars.t" .}}
  },
  "resources": [
    {{range .AgentPoolProfiles}}
      {{if .IsWindows}}
        {{if .IsAvailabilitySets}}
          {{template "dcos/dcosWindowsAgentResourcesVmas.t" .}},
        {{else}}
          {{template "dcos/dcosWindowsAgentResourcesVmss.t" .}},
        {{end}}
      {{else}}
        {{if .IsAvailabilitySets}}
          {{template "dcos/dcosagentresourcesvmas.t" .}},
        {{else}}
          {{template "dcos/dcosagentresourcesvmss.t" .}},
        {{end}}
      {{end}}
    {{end}}
    {{template "dcos/bstrap/bootstrapresources.t" .}},
    {{template "dcos/bstrap/dcosmasterresources.t" .}}
  ],
  "outputs": {
    {{range .AgentPoolProfiles}}{{template "agentoutputs.t" .}}
    {{end}}
    {{template "masteroutputs.t" .}}
  }
}
`)

func dcosBstrapDcosbaseTBytes() ([]byte, error) {
	return _dcosBstrapDcosbaseT, nil
}

func dcosBstrapDcosbaseT() (*asset, error) {
	bytes, err := dcosBstrapDcosbaseTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/bstrap/dcosbase.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosBstrapDcosmasterresourcesT = []byte(`{{if .MasterProfile.IsManagedDisks}}
    {
      "apiVersion": "[variables('apiVersionStorageManagedDisks')]",
      "location": "[variables('location')]",
      "name": "[variables('masterAvailabilitySet')]",
      "properties": {
        "platformFaultDomainCount": 2,
        "platformUpdateDomainCount": 3,
        "managed": "true"
      },
      "type": "Microsoft.Compute/availabilitySets"
    },
{{else if .MasterProfile.IsStorageAccount}}
    {
      "apiVersion": "[variables('apiVersionStorage')]",
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[variables('masterStorageAccountName')]",
      "properties": {
        "accountType": "[variables('vmSizesMap')[variables('masterVMSize')].storageAccountType]"
      },
      "type": "Microsoft.Storage/storageAccounts"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('masterAvailabilitySet')]",
      "properties": {},
      "type": "Microsoft.Compute/availabilitySets"
    },
{{end}}
    {
      "apiVersion": "[variables('apiVersionStorage')]",
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[variables('masterStorageAccountExhibitorName')]",
      "properties": {
        "accountType": "Standard_LRS"
      },
      "type": "Microsoft.Storage/storageAccounts"
    },
{{if not .MasterProfile.IsCustomVNET}}
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "dependsOn": [
          {{GetVNETSubnetDependencies}}
      ],
      "location": "[variables('location')]",
      "name": "[variables('virtualNetworkName')]",
      "properties": {
        "addressSpace": {
          "addressPrefixes": [
            {{GetVNETAddressPrefixes}}
          ]
        },
        "subnets": [
          {{GetVNETSubnets true}}
        ]
      },
      "type": "Microsoft.Network/virtualNetworks"
    },
{{end}}
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('masterPublicIPAddressName')]",
      "properties": {
        "dnsSettings": {
          "domainNameLabel": "[variables('masterEndpointDNSNamePrefix')]"
        },
        "publicIPAllocationMethod": "Dynamic"
      },
      "type": "Microsoft.Network/publicIPAddresses"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[variables('masterLbName')]",
      "properties": {
        "backendAddressPools": [
          {
            "name": "[variables('masterLbBackendPoolName')]"
          }
        ],
        "frontendIPConfigurations": [
          {
            "name": "[variables('masterLbIPConfigName')]",
            "properties": {
              "publicIPAddress": {
                "id": "[resourceId('Microsoft.Network/publicIPAddresses',variables('masterPublicIPAddressName'))]"
              }
            }
          }
        ]
{{if .MasterProfile.OAuthEnabled}}
        ,"loadBalancingRules": [
	        {
            "name": "LBRule443",
            "properties": {
              "frontendIPConfiguration": {
                "id": "[variables('masterLbIPConfigID')]"
              },
              "frontendPort": 443,
              "backendPort": 443,
              "enableFloatingIP": false,
              "idleTimeoutInMinutes": 4,
              "protocol": "Tcp",
              "loadDistribution": "Default",
              "backendAddressPool": {
                "id": "[concat(variables('masterLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"
              },
              "probe": {
                "id": "[concat(variables('masterLbID'),'/probes/dcosMasterProbe')]"
              }
            }
          },
          {
            "name": "LBRule80",
            "properties": {
              "frontendIPConfiguration": {
                "id": "[variables('masterLbIPConfigID')]"
              },
              "frontendPort": 80,
              "backendPort": 80,
              "enableFloatingIP": false,
              "idleTimeoutInMinutes": 4,
              "protocol": "Tcp",
              "loadDistribution": "Default",
              "backendAddressPool": {
                "id": "[concat(variables('masterLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"
              },
              "probe": {
                "id": "[concat(variables('masterLbID'),'/probes/dcosMasterProbe')]"
              }
            }
          }
        ],
        "probes": [
          {
            "name": "dcosMasterProbe",
            "properties": {
              "protocol": "Http",
              "port": 5050,
              "requestPath": "/health",
              "intervalInSeconds": 5,
              "numberOfProbes": 2
            }
          }
        ]
{{end}}
      },
      "type": "Microsoft.Network/loadBalancers"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "copy": {
        "count": "[variables('masterCount')]",
        "name": "masterLbLoopNode"
      },
      "dependsOn": [
        "[variables('masterLbID')]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('masterLbName'), '/', 'SSH-', variables('masterVMNamePrefix'), copyIndex())]",
      "properties": {
        "backendPort": 22,
        "enableFloatingIP": false,
        "frontendIPConfiguration": {
          "id": "[variables('masterLbIPConfigID')]"
        },
        "frontendPort": "[copyIndex(2200)]",
        "protocol": "Tcp"
      },
      "type": "Microsoft.Network/loadBalancers/inboundNatRules"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('masterNSGName')]",
      "properties": {
        "securityRules": [
          {
              "properties": {
                  "priority": 200,
                  "access": "Allow",
                  "direction": "Inbound",
                  "destinationPortRange": "22",
                  "sourcePortRange": "*",
                  "destinationAddressPrefix": "*",
                  "protocol": "Tcp",
                  "description": "Allow SSH",
                  "sourceAddressPrefix": "*"
              },
              "name": "ssh"
          }
        ]
      },
      "type": "Microsoft.Network/networkSecurityGroups"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "copy": {
        "count": "[variables('masterCount')]",
        "name": "nicLoopNode"
      },
      "dependsOn": [
        "[variables('masterNSGID')]",
{{if not .MasterProfile.IsCustomVNET}}
        "[variables('vnetID')]",
{{end}}
        "[variables('masterLbID')]",
        "[concat(variables('masterLbID'),'/inboundNatRules/SSH-',variables('masterVMNamePrefix'),copyIndex())]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('masterVMNamePrefix'), 'nic-', copyIndex())]",
      "properties": {
        "ipConfigurations": [
          {
            "name": "ipConfigNode",
            "properties": {
              "loadBalancerBackendAddressPools": [
                {
                  "id": "[concat(variables('masterLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"
                }
              ],
              "loadBalancerInboundNatRules": [
                {
                  "id": "[concat(variables('masterLbID'),'/inboundNatRules/SSH-',variables('masterVMNamePrefix'),copyIndex())]"
                }
              ],
              "privateIPAddress": "[concat(variables('masterFirstAddrPrefix'), copyIndex(int(variables('masterFirstAddrOctet4'))))]",
              "privateIPAllocationMethod": "Static",
              "subnet": {
                "id": "[variables('masterVnetSubnetID')]"
              }
            }
          }
        ]
        ,"networkSecurityGroup": {
          "id": "[variables('masterNSGID')]"
        }
      },
      "type": "Microsoft.Network/networkInterfaces"
    },
    {
{{if .MasterProfile.IsManagedDisks}}
      "apiVersion": "[variables('apiVersionStorageManagedDisks')]",
{{else}}
      "apiVersion": "[variables('apiVersionDefault')]",
{{end}}
      "copy": {
        "count": "[variables('masterCount')]",
        "name": "vmLoopNode"
      },
      "dependsOn": [
        "[concat('Microsoft.Network/networkInterfaces/', variables('masterVMNamePrefix'), 'nic-', copyIndex())]",
        "[concat('Microsoft.Compute/availabilitySets/',variables('masterAvailabilitySet'))]",
{{if .MasterProfile.IsStorageAccount}}
        "[variables('masterStorageAccountName')]",
{{end}}
        "[variables('masterStorageAccountExhibitorName')]"
       ,"[concat('Microsoft.Compute/virtualMachines/', variables('bootstrapVMName'), '/extensions/bootstrapready')]"
      ],
      "tags":
      {
        "creationSource" : "[concat('acsengine-', variables('masterVMNamePrefix'), copyIndex())]",
        "orchestratorName": "dcos",
        "orchestratorVersion": "[variables('orchestratorVersion')]",
        "orchestratorNode": "master"
      },
      "location": "[variables('location')]",
      "name": "[concat(variables('masterVMNamePrefix'), copyIndex())]",
      "properties": {
        "availabilitySet": {
          "id": "[resourceId('Microsoft.Compute/availabilitySets',variables('masterAvailabilitySet'))]"
        },
        "hardwareProfile": {
          "vmSize": "[variables('masterVMSize')]"
        },
        "networkProfile": {
          "networkInterfaces": [
            {
              "id": "[resourceId('Microsoft.Network/networkInterfaces',concat(variables('masterVMNamePrefix'), 'nic-', copyIndex()))]"
            }
          ]
        },
        "osProfile": {
          "adminUsername": "[variables('adminUsername')]",
          "computername": "[concat(variables('masterVMNamePrefix'), copyIndex())]",
          {{GetDCOSMasterCustomData}}
          "linuxConfiguration": {
            "disablePasswordAuthentication": true,
            "ssh": {
                "publicKeys": [
                    {
                        "keyData": "[variables('sshRSAPublicKey')]",
                        "path": "[variables('sshKeyPath')]"
                    }
                ]
            }
          }
          {{if .LinuxProfile.HasSecrets}}
          ,
          "secrets": "[variables('linuxProfileSecrets')]"
          {{end}}
        },
        "storageProfile": {
          "imageReference": {
            "offer": "[variables('osImageOffer')]",
            "publisher": "[variables('osImagePublisher')]",
            "sku": "[variables('osImageSKU')]",
            "version": "[variables('osImageVersion')]"
          },
          "osDisk": {
            "caching": "ReadWrite"
            ,"createOption": "FromImage"
{{if .MasterProfile.IsStorageAccount}}
            ,"name": "[concat(variables('masterVMNamePrefix'), copyIndex(),'-osdisk')]"
            ,"vhd": {
              "uri": "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('masterStorageAccountName')),variables('apiVersionStorage')).primaryEndpoints.blob,'vhds/',variables('masterVMNamePrefix'),copyIndex(),'-osdisk.vhd')]"
            }
{{end}}
{{if ne .MasterProfile.OSDiskSizeGB 0}}
            ,"diskSizeGB": {{.MasterProfile.OSDiskSizeGB}}
{{end}}
          }
        }
      },
      "type": "Microsoft.Compute/virtualMachines"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "dependsOn": [
        "[concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), sub(variables('masterCount'), 1))]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('masterVMNamePrefix'), sub(variables('masterCount'), 1), '/waitforleader')]",
      "properties": {
        "autoUpgradeMinorVersion": true,
        "publisher": "Microsoft.OSTCExtensions",
        "settings": {
          "commandToExecute": "sh -c 'until ping -c1 leader.mesos;do echo waiting for leader.mesos;sleep 15;done;echo leader.mesos up'"
        },
        "type": "CustomScriptForLinux",
        "typeHandlerVersion": "1.4"
      },
      "type": "Microsoft.Compute/virtualMachines/extensions"
    }{{WriteLinkedTemplatesForExtensions}}
`)

func dcosBstrapDcosmasterresourcesTBytes() ([]byte, error) {
	return _dcosBstrapDcosmasterresourcesT, nil
}

func dcosBstrapDcosmasterresourcesT() (*asset, error) {
	bytes, err := dcosBstrapDcosmasterresourcesTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/bstrap/dcosmasterresources.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosBstrapDcosmastervarsT = []byte(`    "adminUsername": "[parameters('linuxAdminUsername')]",
    "targetEnvironment": "[parameters('targetEnvironment')]",
    "maxVMsPerPool": 100,
    "apiVersionDefault": "2016-03-30",
    "apiVersionLinkDefault": "2015-01-01",
    "singleQuote": "'",
    "doubleSingleQuote": "''",
{{if .LinuxProfile.HasSecrets}}
    "linuxProfileSecrets" :
      [
          {{range  $vIndex, $vault := .LinuxProfile.Secrets}}
            {{if $vIndex}} , {{end}}
              {
                "sourceVault":{
                  "id":"[parameters('linuxKeyVaultID{{$vIndex}}')]"
                },
                "vaultCertificates":[
                {{range $cIndex, $cert := $vault.VaultCertificates}}
                  {{if $cIndex}} , {{end}}
                  {
                    "certificateUrl" :"[parameters('linuxKeyVaultID{{$vIndex}}CertificateURL{{$cIndex}}')]"
                  }
                {{end}}
                ]
              }
        {{end}}
      ],
{{end}}
    "orchestratorVersion": "{{.OrchestratorProfile.OrchestratorVersion}}",
{{if .HasWindows}}
    "windowsAdminUsername": "[parameters('windowsAdminUsername')]",
    "windowsAdminPassword": "[parameters('windowsAdminPassword')]",
    "agentWindowsBackendPort": 3389,
    "agentWindowsPublisher": "[parameters('agentWindowsPublisher')]",
    "agentWindowsOffer": "[parameters('agentWindowsOffer')]",
    "agentWindowsSku": "[parameters('agentWindowsSku')]",
    "agentWindowsVersion": "[parameters('agentWindowsVersion')]",
    "dcosWindowsBootstrapURL" : "[parameters('dcosWindowsBootstrapURL')]",
    "windowsCustomScriptSuffix": " $inputFile = '%SYSTEMDRIVE%\\AzureData\\CustomData.bin' ; $outputFile = '%SYSTEMDRIVE%\\AzureData\\dcosWindowsProvision.ps1' ; $inputStream = New-Object System.IO.FileStream $inputFile, ([IO.FileMode]::Open), ([IO.FileAccess]::Read), ([IO.FileShare]::Read) ; $sr = New-Object System.IO.StreamReader(New-Object System.IO.Compression.GZipStream($inputStream, [System.IO.Compression.CompressionMode]::Decompress)) ; $sr.ReadToEnd() | Out-File($outputFile) ; Invoke-Expression('{0} {1}' -f $outputFile, $arguments) ; ",
    "windowsMasterCustomScriptArguments": "[concat('$arguments = ', variables('singleQuote'),'-MasterCount ', variables('masterCount'), ' -firstMasterIP ', parameters('firstConsecutiveStaticIP'), variables('singleQuote'), ' ; ')]",

    "windowsMasterCustomScript": "[concat('powershell.exe -ExecutionPolicy Unrestricted -command \"', variables('windowsMasterCustomScriptArguments'), variables('windowsCustomScriptSuffix'), '\" > %SYSTEMDRIVE%\\AzureData\\dcosWindowsProvision.log 2>&1')]",
{{end}}
    "masterAvailabilitySet": "[concat(variables('orchestratorName'), '-master-availabilitySet-', variables('nameSuffix'))]",
    "masterCount": {{.MasterProfile.Count}},
    "masterEndpointDNSNamePrefix": "[tolower(parameters('masterEndpointDNSNamePrefix'))]",
    "masterHttpSourceAddressPrefix": "{{.MasterProfile.HTTPSourceAddressPrefix}}",
    "masterLbBackendPoolName": "[concat(variables('orchestratorName'), '-master-pool-', variables('nameSuffix'))]",
    "masterLbID": "[resourceId('Microsoft.Network/loadBalancers',variables('masterLbName'))]",
    "masterLbIPConfigID": "[concat(variables('masterLbID'),'/frontendIPConfigurations/', variables('masterLbIPConfigName'))]",
    "masterLbIPConfigName": "[concat(variables('orchestratorName'), '-master-lbFrontEnd-', variables('nameSuffix'))]",
    "masterLbName": "[concat(variables('orchestratorName'), '-master-lb-', variables('nameSuffix'))]",
    "masterNSGID": "[resourceId('Microsoft.Network/networkSecurityGroups',variables('masterNSGName'))]",
    "masterNSGName": "[concat(variables('orchestratorName'), '-master-nsg-', variables('nameSuffix'))]",
    "masterPublicIPAddressName": "[concat(variables('orchestratorName'), '-master-ip-', variables('masterEndpointDNSNamePrefix'), '-', variables('nameSuffix'))]",
    "apiVersionStorage": "2015-06-15",

    "storageAccountBaseName": "[uniqueString(concat(variables('masterEndpointDNSNamePrefix'),variables('location'),variables('orchestratorName')))]",
    "masterStorageAccountExhibitorName": "[concat(variables('storageAccountBaseName'), 'exhb0')]",
    "storageAccountType": "Standard_LRS",
{{if .HasStorageAccountDisks}}
    "maxVMsPerStorageAccount": 20,
    "maxStorageAccountsPerAgent": "[div(variables('maxVMsPerPool'),variables('maxVMsPerStorageAccount'))]",
    "dataStorageAccountPrefixSeed": 97,
    "storageAccountPrefixes": [ "0", "6", "c", "i", "o", "u", "1", "7", "d", "j", "p", "v", "2", "8", "e", "k", "q", "w", "3", "9", "f", "l", "r", "x", "4", "a", "g", "m", "s", "y", "5", "b", "h", "n", "t", "z" ],
    "storageAccountPrefixesCount": "[length(variables('storageAccountPrefixes'))]",
    {{GetSizeMap}},
{{else}}
    "storageAccountPrefixes": [],
{{end}}
{{if .HasManagedDisks}}
    "apiVersionStorageManagedDisks": "2016-04-30-preview",
{{end}}
{{if .MasterProfile.IsStorageAccount}}
    "masterStorageAccountName": "[concat(variables('storageAccountBaseName'), 'mstr0')]",
{{end}}
{{if .MasterProfile.IsCustomVNET}}
    "masterVnetSubnetID": "[parameters('masterVnetSubnetID')]",
{{else}}
    "masterSubnet": "[parameters('masterSubnet')]",
    "masterSubnetName": "[concat(variables('orchestratorName'), '-masterSubnet')]",
    "vnetID": "[resourceId('Microsoft.Network/virtualNetworks',variables('virtualNetworkName'))]",
    "masterVnetSubnetID": "[concat(variables('vnetID'),'/subnets/',variables('masterSubnetName'))]",
    "virtualNetworkName": "[concat(variables('orchestratorName'), '-vnet-', variables('nameSuffix'))]",
{{end}}
    "masterFirstAddrOctets": "[split(parameters('firstConsecutiveStaticIP'),'.')]",
    "masterFirstAddrOctet4": "[variables('masterFirstAddrOctets')[3]]",
    "masterFirstAddrPrefix": "[concat(variables('masterFirstAddrOctets')[0],'.',variables('masterFirstAddrOctets')[1],'.',variables('masterFirstAddrOctets')[2],'.')]",
    "masterVMNamePrefix": "[concat(variables('orchestratorName'), '-master-', variables('nameSuffix'), '-')]",
    "masterVMNic": [
      "[concat(variables('masterVMNamePrefix'), 'nic-0')]",
      "[concat(variables('masterVMNamePrefix'), 'nic-1')]",
      "[concat(variables('masterVMNamePrefix'), 'nic-2')]",
      "[concat(variables('masterVMNamePrefix'), 'nic-3')]",
      "[concat(variables('masterVMNamePrefix'), 'nic-4')]",
      "[concat(variables('masterVMNamePrefix'), 'nic-5')]",
      "[concat(variables('masterVMNamePrefix'), 'nic-6')]"
    ],
    "masterVMSize": "[parameters('masterVMSize')]",
    "nameSuffix": "[parameters('nameSuffix')]",
    "oauthEnabled": "{{.MasterProfile.OAuthEnabled}}",
    "orchestratorName": "dcos",
    "osImageOffer": "[parameters('osImageOffer')]",
    "osImagePublisher": "[parameters('osImagePublisher')]",
    "osImageSKU": "[parameters('osImageSKU')]",
    "osImageVersion": "[parameters('osImageVersion')]",
    "sshKeyPath": "[concat('/home/', variables('adminUsername'), '/.ssh/authorized_keys')]",
    "sshRSAPublicKey": "[parameters('sshRSAPublicKey')]",
    "locations": [
         "[resourceGroup().location]",
         "[parameters('location')]"
    ],
    "location": "[variables('locations')[mod(add(2,length(parameters('location'))),add(1,length(parameters('location'))))]]",
    "masterSshInboundNatRuleIdPrefix": "[concat(variables('masterLbID'),'/inboundNatRules/SSH-',variables('masterVMNamePrefix'))]",
    "masterLbInboundNatRules": [
            [
                {
                    "id": "[concat(variables('masterSshInboundNatRuleIdPrefix'),'0')]"
                }
            ],
            [
                {
                    "id": "[concat(variables('masterSshInboundNatRuleIdPrefix'),'1')]"
                }
            ],
            [
                {
                    "id": "[concat(variables('masterSshInboundNatRuleIdPrefix'),'2')]"
                }
            ],
            [
                {
                    "id": "[concat(variables('masterSshInboundNatRuleIdPrefix'),'3')]"
                }
            ],
            [
                {
                    "id": "[concat(variables('masterSshInboundNatRuleIdPrefix'),'4')]"
                }
            ]
        ]
`)

func dcosBstrapDcosmastervarsTBytes() ([]byte, error) {
	return _dcosBstrapDcosmastervarsT, nil
}

func dcosBstrapDcosmastervarsT() (*asset, error) {
	bytes, err := dcosBstrapDcosmastervarsTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/bstrap/dcosmastervars.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosBstrapDcosprovisionSh = []byte(`#!/bin/bash

source /opt/azure/containers/provision_source.sh
source /opt/azure/dcos/environment

TMPDIR="/tmp/dcos"
mkdir -p $TMPDIR

# default dc/os component download address (Azure CDN)
packages=(
  https://dcos-mirror.azureedge.net/pkg/libipset3_6.29-1_amd64.deb
  https://dcos-mirror.azureedge.net/pkg/ipset_6.29-1_amd64.deb
  https://dcos-mirror.azureedge.net/pkg/unzip_6.0-20ubuntu1_amd64.deb
  https://dcos-mirror.azureedge.net/pkg/libltdl7_2.4.6-0.1_amd64.deb
  https://dcos-mirror.azureedge.net/pkg/docker-ce_17.09.0~ce-0~ubuntu_amd64.deb
  https://dcos-mirror.azureedge.net/pkg/selinux-utils_2.4-3build2_amd64.deb
)

# sha1sum checksums for @packages
sha1sums=(
  f88d09688291917c8bb65682fea9f5d571ec8d6a
  807dc11f5bfa39bb4b0dc9024fc51bb309905a21
  57ae2bb6ded1fdf91b6d518294134df1ff13fcca
  9a0f9f2769d3dc834737aa7df50aaaea369af98d
  94f6e89be6d45d9988269a237eb27c7d6a844d7f
  77bdb5847060845c0a158f567b1ddd7fa34b7236
)

case $DCOS_ENVIRONMENT in
  # because of Chinese GreatWall Firewall, the default packages on Azure CDN is blocked. So the following Chinese local mirror url should be used instead.
  AzureChinaCloud)
    packages=(
      http://acsengine.blob.core.chinacloudapi.cn/dcos/libipset3_6.29-1_amd64.deb
      http://acsengine.blob.core.chinacloudapi.cn/dcos/ipset_6.29-1_amd64.deb
      http://acsengine.blob.core.chinacloudapi.cn/dcos/unzip_6.0-20ubuntu1_amd64.deb
      http://acsengine.blob.core.chinacloudapi.cn/dcos/libltdl7_2.4.6-0.1_amd64.deb
      http://mirror.kaiyuanshe.cn/docker-ce/linux/ubuntu/dists/xenial/pool/stable/amd64/docker-ce_17.09.0~ce-0~ubuntu_amd64.deb
      http://acsengine.blob.core.chinacloudapi.cn/dcos/selinux-utils_2.4-3build2_amd64.deb
    )
  ;;
esac

len=$((${#packages[@]}-1))
for i in $(seq 0 $len); do
  retry_get_install_deb 10 10 120 ${packages[$i]} ${sha1sums[$i]}
    if [ $? -ne 0  ]; then
    exit 1
  fi
done

retrycmd_if_failure 10 10 120 curl -fsSL -o $TMPDIR/dcos_install.sh http://BOOTSTRAP_IP:8086/dcos_install.sh
if [ $? -ne 0  ]; then
  exit 1
fi
`)

func dcosBstrapDcosprovisionShBytes() ([]byte, error) {
	return _dcosBstrapDcosprovisionSh, nil
}

func dcosBstrapDcosprovisionSh() (*asset, error) {
	bytes, err := dcosBstrapDcosprovisionShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/bstrap/dcosprovision.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosDcoswindowsagentresourcesvmasT = []byte(`    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}NSGName')]",
      "properties": {
        "securityRules": [
            {{GetSecurityRules .Ports}}
        ]
      },
      "type": "Microsoft.Network/networkSecurityGroups"
    },
{{if HasWindowsCustomImage}}
    {"type": "Microsoft.Compute/images",
      "apiVersion": "2017-12-01",
      "name": "{{.Name}}CustomWindowsImage",
      "location": "[variables('location')]",
      "properties": {
        "storageProfile": {
          "osDisk": {
            "osType": "Windows",
            "osState": "Generalized",
            "blobUri": "[parameters('agentWindowsSourceUrl')]",
            "storageAccountType": "Standard_LRS"
          }
        }
      }
    },
{{end}}
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "copy": {
        "count": "[sub(variables('{{.Name}}Count'), variables('{{.Name}}Offset'))]",
        "name": "loop"
      },
      "dependsOn": [
{{if .IsCustomVNET}}
      "[concat('Microsoft.Network/networkSecurityGroups/', variables('{{.Name}}NSGName'))]"
{{else}}
      "[variables('vnetID')]"
{{end}}
{{if IsPublic .Ports}}
	  ,"[variables('{{.Name}}LbID')]"
{{end}}
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('{{.Name}}VMNamePrefix'), 'nic-', copyIndex(variables('{{.Name}}Offset')))]",
      "properties": {
{{if .IsCustomVNET}}
	    "networkSecurityGroup": {
		  "id": "[resourceId('Microsoft.Network/networkSecurityGroups/', variables('{{.Name}}NSGName'))]"
	    },
{{end}}
        "ipConfigurations": [
          {
            "name": "ipConfigNode",
            "properties": {
{{if IsPublic .Ports}}
              "loadBalancerBackendAddressPools": [
		        {
		      	  "id": "[concat('/subscriptions/', subscription().subscriptionId,'/resourceGroups/', resourceGroup().name, '/providers/Microsoft.Network/loadBalancers/', variables('{{.Name}}LbName'), '/backendAddressPools/',variables('{{.Name}}LbBackendPoolName'))]"
		        }
              ],
              "loadBalancerInboundNatPools": [
                {
                  "id": "[concat(variables('{{.Name}}LbID'), '/inboundNatPools/', 'RDP-', variables('{{.Name}}VMNamePrefix'))]"
                }
		      ],
{{end}}
              "privateIPAllocationMethod": "Dynamic",
              "subnet": {
                "id": "[variables('{{.Name}}VnetSubnetID')]"
             }
            }
          }
        ]
      },
      "type": "Microsoft.Network/networkInterfaces"
    },
{{if .IsManagedDisks}}
    {
      "apiVersion": "[variables('apiVersionStorageManagedDisks')]",
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}AvailabilitySet')]",
      "properties": {
        "platformFaultDomainCount": 2,
        "platformUpdateDomainCount": 3,
        "managed": "true"
      },
      "type": "Microsoft.Compute/availabilitySets"
    },
{{else if .IsStorageAccount}}
    {
      "apiVersion": "[variables('apiVersionStorage')]",
      "copy": {
        "count": "[variables('{{.Name}}StorageAccountsCount')]",
        "name": "loop"
      },
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('storageAccountPrefixes')[mod(add(copyIndex(),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(copyIndex(),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
      "properties": {
        "accountType": "[variables('vmSizesMap')[variables('{{.Name}}VMSize')].storageAccountType]"
      },
      "type": "Microsoft.Storage/storageAccounts"
    },
    {{if .HasDisks}}
        {
          "apiVersion": "[variables('apiVersionStorage')]",
          "copy": {
            "count": "[variables('{{.Name}}StorageAccountsCount')]",
            "name": "datadiskLoop"
          },
          "dependsOn": [
            "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
          ],
          "location": "[variables('location')]",
          "name": "[concat(variables('storageAccountPrefixes')[mod(add(copyIndex(variables('dataStorageAccountPrefixSeed')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(copyIndex(variables('dataStorageAccountPrefixSeed')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}DataAccountName'))]",
          "properties": {
            "accountType": "[variables('vmSizesMap')[variables('{{.Name}}VMSize')].storageAccountType]"
          },
          "type": "Microsoft.Storage/storageAccounts"
        },
    {{end}}
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}AvailabilitySet')]",
      "properties": {},
      "type": "Microsoft.Compute/availabilitySets"
    },
{{end}}
{{if IsPublic .Ports}}
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}IPAddressName')]",
      "properties": {
        "dnsSettings": {
          "domainNameLabel": "[variables('{{.Name}}EndpointDNSNamePrefix')]"
        },
        "publicIPAllocationMethod": "Dynamic"
      },
      "type": "Microsoft.Network/publicIPAddresses"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('{{.Name}}IPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}LbName')]",
      "properties": {
        "backendAddressPools": [
          {
            "name": "[variables('{{.Name}}LbBackendPoolName')]"
          }
        ],
        "frontendIPConfigurations": [
          {
            "name": "[variables('{{.Name}}LbIPConfigName')]",
            "properties": {
              "publicIPAddress": {
                "id": "[resourceId('Microsoft.Network/publicIPAddresses',variables('{{.Name}}IPAddressName'))]"
              }
            }
          }
        ],
        "inboundNatPools": [
          {
            "name": "[concat('RDP-', variables('{{.Name}}VMNamePrefix'))]",
            "properties": {
              "frontendIPConfiguration": {
                "id": "[variables('{{.Name}}LbIPConfigID')]"
              },
              "protocol": "Tcp",
              "frontendPortRangeStart": "[variables('{{.Name}}WindowsRDPNatRangeStart')]",
              "frontendPortRangeEnd": "[variables('{{.Name}}WindowsRDPEndRangeStop')]",
              "backendPort": "[variables('agentWindowsBackendPort')]"
            }
          }
        ],
        "loadBalancingRules": [
          {{(GetLBRules .Name .Ports)}}
        ],
        "probes": [
          {{(GetProbes .Ports)}}
        ]
      },
      "type": "Microsoft.Network/loadBalancers"
    },
{{end}}
    {
{{if .IsManagedDisks}}
      "apiVersion": "[variables('apiVersionStorageManagedDisks')]",
{{else}}
      "apiVersion": "[variables('apiVersionDefault')]",
{{end}}
      "copy": {
        "count": "[sub(variables('{{.Name}}Count'), variables('{{.Name}}Offset'))]",
        "name": "vmLoopNode"
      },
      "dependsOn": [
{{if .IsStorageAccount}}
        "[concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
  {{if .HasDisks}}
        "[concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('dataStorageAccountPrefixSeed')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('dataStorageAccountPrefixSeed')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}DataAccountName'))]",
  {{end}}
{{end}}
        "[concat('Microsoft.Network/networkInterfaces/', variables('{{.Name}}VMNamePrefix'), 'nic-', copyIndex(variables('{{.Name}}Offset')))]",
        "[concat('Microsoft.Compute/availabilitySets/', variables('{{.Name}}AvailabilitySet'))]"
      ],
      "tags":
      {
        "creationSource" : "[concat('acsengine-', variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')))]",
        "orchestratorName": "dcos",
        "orchestratorVersion": "[variables('orchestratorVersion')]",
        "orchestratorNode": "agent"
      },
      "location": "[variables('location')]",
      "name": "[concat(variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')))]",
      "properties": {
        "availabilitySet": {
          "id": "[resourceId('Microsoft.Compute/availabilitySets',variables('{{.Name}}AvailabilitySet'))]"
        },
        "hardwareProfile": {
          "vmSize": "[variables('{{.Name}}VMSize')]"
        },
        "networkProfile": {
          "networkInterfaces": [
            {
              "id": "[resourceId('Microsoft.Network/networkInterfaces',concat(variables('{{.Name}}VMNamePrefix'), 'nic-', copyIndex(variables('{{.Name}}Offset'))))]"
            }
          ]
        },
        "osProfile": {
          "computername": "[concat(variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')))]",
          "adminUsername": "[variables('windowsAdminUsername')]",
          "adminPassword": "[variables('windowsAdminPassword')]",
          {{GetDCOSWindowsAgentCustomData .}}

        },
        "storageProfile": {
          {{GetDataDisks .}}
          "imageReference": {
{{if HasWindowsCustomImage}}
            "id": "[resourceId('Microsoft.Compute/images','{{.Name}}CustomWindowsImage')]"
{{else}}
            "offer": "[variables('agentWindowsOffer')]",
            "publisher": "[variables('agentWindowsPublisher')]",
            "sku": "[variables('agentWindowsSKU')]",
            "version": "[variables('agentWindowsVersion')]"
{{end}}
          }
          ,"osDisk": {
            "caching": "ReadOnly"
            ,"createOption": "FromImage"
{{if .IsStorageAccount}}
            ,"name": "[concat(variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')),'-osdisk')]"
            ,"vhd": {
              "uri": "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')),variables('apiVersionStorage')).primaryEndpoints.blob,'osdisk/', variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')), '-osdisk.vhd')]"
            }
{{end}}
{{if ne .OSDiskSizeGB 0}}
            ,"diskSizeGB": {{.OSDiskSizeGB}}
{{end}}
          }
        }
      },
      "type": "Microsoft.Compute/virtualMachines"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "copy": {
        "count": "[sub(variables('{{.Name}}Count'), variables('{{.Name}}Offset'))]",
        "name": "vmLoopNode"
      },
      "dependsOn": [
        "[concat('Microsoft.Compute/virtualMachines/', variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')))]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')), '/cse')]",
      "properties": {
        "publisher": "Microsoft.Compute",
        "type": "CustomScriptExtension",
        "typeHandlerVersion": "1.8",
        "autoUpgradeMinorVersion": true,
        "settings": {
          "commandToExecute": "[variables('{{.Name}}windowsAgentCustomScript')]"
        }
      },
      "type": "Microsoft.Compute/virtualMachines/extensions"
    }
`)

func dcosDcoswindowsagentresourcesvmasTBytes() ([]byte, error) {
	return _dcosDcoswindowsagentresourcesvmasT, nil
}

func dcosDcoswindowsagentresourcesvmasT() (*asset, error) {
	bytes, err := dcosDcoswindowsagentresourcesvmasTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/dcosWindowsAgentResourcesVmas.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosDcoswindowsagentresourcesvmssT = []byte(`    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}NSGName')]",
      "properties": {
        "securityRules": [
            {{GetSecurityRules .Ports}}
        ]
      },
      "type": "Microsoft.Network/networkSecurityGroups"
    },
{{if HasWindowsCustomImage}}
    {"type": "Microsoft.Compute/images",
      "apiVersion": "2017-12-01",
      "name": "{{.Name}}CustomWindowsImage",
      "location": "[variables('location')]",
      "properties": {
        "storageProfile": {
          "osDisk": {
            "osType": "Windows",
            "osState": "Generalized",
            "blobUri": "[parameters('agentWindowsSourceUrl')]",
            "storageAccountType": "Standard_LRS"
          }
        }
      }
    },
{{end}}
{{if .IsStorageAccount}}
    {
      "apiVersion": "[variables('apiVersionStorage')]",
      "copy": {
        "count": "[variables('{{.Name}}StorageAccountsCount')]",
        "name": "loop"
      },
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('storageAccountPrefixes')[mod(add(copyIndex(),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(copyIndex(),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
      "properties": {
        "accountType": "[variables('vmSizesMap')[variables('{{.Name}}VMSize')].storageAccountType]"
      },
      "type": "Microsoft.Storage/storageAccounts"
    },
{{end}}
{{if IsPublic .Ports}}
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}IPAddressName')]",
      "properties": {
        "dnsSettings": {
          "domainNameLabel": "[variables('{{.Name}}EndpointDNSNamePrefix')]"
        },
        "publicIPAllocationMethod": "Dynamic"
      },
      "type": "Microsoft.Network/publicIPAddresses"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('{{.Name}}IPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}LbName')]",
      "properties": {
        "backendAddressPools": [
          {
            "name": "[variables('{{.Name}}LbBackendPoolName')]"
          }
        ],
        "frontendIPConfigurations": [
          {
            "name": "[variables('{{.Name}}LbIPConfigName')]",
            "properties": {
              "publicIPAddress": {
                "id": "[resourceId('Microsoft.Network/publicIPAddresses',variables('{{.Name}}IPAddressName'))]"
              }
            }
          }
        ],
        "inboundNatRules": [],
        "loadBalancingRules": [
            {{(GetLBRules .Name .Ports)}}
        ],
        "probes": [
            {{(GetProbes .Ports)}}
        ]
      },
      "type": "Microsoft.Network/loadBalancers"
    },
{{end}}
    {
{{if .IsManagedDisks}}
      "apiVersion": "[variables('apiVersionStorageManagedDisks')]",
{{else}}
      "apiVersion": "[variables('apiVersionDefault')]",
{{end}}
      "dependsOn": [
{{if .IsCustomVNET}}
      "[concat('Microsoft.Network/networkSecurityGroups/', variables('{{.Name}}NSGName'))]"
{{else}}
      "[variables('vnetID')]"
{{end}}
{{if .IsStorageAccount}}
        ,"[concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(0,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(0,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
        "[concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(1,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(1,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
        "[concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(2,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(2,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
        "[concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(3,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(3,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
        "[concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(4,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(4,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]"
{{end}}
{{if IsPublic .Ports}}
       ,"[concat('Microsoft.Network/loadBalancers/', variables('{{.Name}}LbName'))]"
{{end}}
      ],
      "tags":
      {
        "creationSource" : "[concat('acsengine-', variables('{{.Name}}VMNamePrefix'), '-vmss')]",
        "orchestratorName": "dcos",
        "orchestratorVersion": "[variables('orchestratorVersion')]",
        "orchestratorNode": "agent"
      },
      "location": "[variables('location')]",
      "name": "[concat(variables('{{.Name}}VMNamePrefix'), '-vmss')]",
      "properties": {
        "overprovision": false,
        "upgradePolicy": {
          "mode": "Manual"
        },
        "virtualMachineProfile": {
          "networkProfile": {
            "networkInterfaceConfigurations": [
              {
                "name": "nic",
                "properties": {
{{if .IsCustomVNET}}
                  "networkSecurityGroup": {
                    "id": "[resourceId('Microsoft.Network/networkSecurityGroups/', variables('{{.Name}}NSGName'))]"
                  },
{{end}}
                  "ipConfigurations": [
                    {
                      "name": "nicipconfig",
                      "properties": {
{{if IsPublic .Ports}}
                        "loadBalancerBackendAddressPools": [
                          {
                            "id": "[concat('/subscriptions/', subscription().subscriptionId,'/resourceGroups/', resourceGroup().name, '/providers/Microsoft.Network/loadBalancers/', variables('{{.Name}}LbName'), '/backendAddressPools/',variables('{{.Name}}LbBackendPoolName'))]"
                          }
                        ],
{{end}}
                        "subnet": {
                          "id": "[variables('{{.Name}}VnetSubnetID')]"
                        }
                      }
                    }
                  ],
                  "primary": "true"
                }
              }
            ]
          },
          "osProfile": {
            "computerNamePrefix": "[concat(substring(variables('nameSuffix'), 0, 5), 'acs')]",
            "adminUsername": "[variables('windowsAdminUsername')]",
            "adminPassword": "[variables('windowsAdminPassword')]",
            {{GetDCOSWindowsAgentCustomData .}}
            {{if HasWindowsSecrets}}
              ,
              "secrets": "[variables('windowsProfileSecrets')]"
            {{end}}
          },
          "storageProfile": {
            "imageReference": {
{{if HasWindowsCustomImage}}
              "id": "[resourceId('Microsoft.Compute/images','{{.Name}}CustomWindowsImage')]"
{{else}}
              "publisher": "[variables('agentWindowsPublisher')]",
              "offer": "[variables('agentWindowsOffer')]",
              "sku": "[variables('agentWindowsSku')]",
              "version": "latest"
{{end}}
            },
            {{GetDataDisks .}}
            "osDisk": {
              "caching": "ReadOnly",
              "createOption": "FromImage"
{{if .IsStorageAccount}}
              ,"name": "vmssosdisk"
              ,"vhdContainers": [
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(0,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(0,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')),variables('apiVersionStorage')).primaryEndpoints.blob,'osdisk')]",
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(1,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(1,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')),variables('apiVersionStorage')).primaryEndpoints.blob,'osdisk')]",
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(2,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(2,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')),variables('apiVersionStorage')).primaryEndpoints.blob,'osdisk')]",
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(3,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(3,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')),variables('apiVersionStorage')).primaryEndpoints.blob,'osdisk')]",
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(4,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(4,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')),variables('apiVersionStorage')).primaryEndpoints.blob,'osdisk')]"

              ]
{{end}}
{{if ne .OSDiskSizeGB 0}}
            ,"diskSizeGB": {{.OSDiskSizeGB}}
{{end}}
            }
          },
          "extensionProfile": {
            "extensions": [
              {
                "name": "vmssCustomScriptExtension",
                "properties": {
                  "publisher": "Microsoft.Compute",
                  "type": "CustomScriptExtension",
                  "typeHandlerVersion": "1.8",
                  "autoUpgradeMinorVersion": true,
                  "settings": {
                     "commandToExecute": "[variables('{{.Name}}windowsAgentCustomScript')]"
                  }
                }
              }
            ]
          }
        }
      },
      "sku": {
        "capacity": "[variables('{{.Name}}Count')]",
        "name": "[variables('{{.Name}}VMSize')]",
        "tier": "[variables('{{.Name}}VMSizeTier')]"
      },
      "type": "Microsoft.Compute/virtualMachineScaleSets"
    }
`)

func dcosDcoswindowsagentresourcesvmssTBytes() ([]byte, error) {
	return _dcosDcoswindowsagentresourcesvmssT, nil
}

func dcosDcoswindowsagentresourcesvmssT() (*asset, error) {
	bytes, err := dcosDcoswindowsagentresourcesvmssTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/dcosWindowsAgentResourcesVmss.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosDcoswindowsprovisionPs1 = []byte(`<#
    .SYNOPSIS
        Provisions VM as a DCOS agent.

    .DESCRIPTION
        Provisions VM as a DCOS agent.

     Invoke by:
       
#>

[CmdletBinding(DefaultParameterSetName="Standard")]
param(
    [string]
    [ValidateNotNullOrEmpty()]
    $masterCount,

    [string]
    [ValidateNotNullOrEmpty()]
    $firstMasterIP,
    
    [string]
    [ValidateNotNullOrEmpty()]
    $bootstrapUri,

    [parameter()]
    [ValidateNotNullOrEmpty()]
    $isAgent,

    [parameter()]
    [ValidateNotNullOrEmpty()]
    $subnet,

    [parameter()]
    [AllowNull()]
    $isPublic = $false,

    [string]
    [AllowNull()]
    $customAttrs = "",

    [string]
    [AllowNull()]
    $preprovisionExtensionParams = ""
)




$global:BootstrapInstallDir = "C:\AzureData"

filter Timestamp {"$(Get-Date -Format o): $_"}


function
Write-Log($message)
{
    $msg = $message | Timestamp
    Write-Output $msg
}


function
Expand-ZIPFile($file, $destination)
{
    $shell = new-object -com shell.application
    $zip = $shell.NameSpace($file)
    foreach($item in $zip.items())
    {
        $shell.Namespace($destination).copyhere($item, 0x14)
    }
}


function 
Remove-Directory($dirname)
{

    try {
        #Get-ChildItem $dirname -Recurse | Remove-Item  -force -confirm:$false
        # This doesn't work because of long file names
        # But this does:
        Invoke-Expression ("cmd /C rmdir /s /q "+$dirname)
    }
    catch {
        # If this fails we don't want it to stop

    }
}


function 
Check-Subnet ([string]$cidr, [string]$ip)
{
    try {

        $network, [int]$subnetlen = $cidr.Split('/')
    
        if ($subnetlen -eq 0)
        {
            $subnetlen = 8 # Default in case we get an IP addr, not CIDR
        }
        $a = ([IPAddress] $network)
        [uint32] $unetwork = [uint32]$a.Address
    
        $mask = -bnot ((-bnot [uint32]0) -shl (32 - $subnetlen))
    
        $a = [IPAddress]$ip
        [uint32] $uip = [uint32]$a.Address
    
        return ($unetwork -eq ($mask -band $uip))
    }
    catch {
        return $false
    }
}

#
# Gets the bootstrap script from the blob store and places it in c:\AzureData

function
Get-BootstrapScript($download_uri, $download_dir)
{
    # Get Mesos Binaries
    $scriptfile = "DCOSWindowsAgentSetup.ps1"

    Write-Log "get script $download_uri/$scriptfile and put it $download_dir\$scriptfile"

    Invoke-WebRequest -Uri ($download_uri+"/"+$scriptfile) -OutFile ($download_dir+"\"+$scriptfile)
}


try
{
    # Set to false for debugging.  This will output the start script to
    # c:\AzureData\dcosProvisionScript.log, and then you can RDP 
    # to the windows machine, and run the script manually to watch
    # the output.
    Write-Log "Get the install script"

    Write-Log ("Parameters = isAgent = ["+ $isAgent + "] mastercount = ["+$MasterCount + "] First master ip= [" + $firstMasterIp+ "] boostrap URI = ["+ $bootstrapUri+"] Subnet = ["+ $subnet +"]" + " -customAttrs " + $customAttrs + " -preprovisionExtensionParms = "+ $preprovisionExtensionParams )

    # Get the boostrap script

    Get-BootstrapScript $bootstrapUri $global:BootstrapInstallDir

    # Convert Master count and first IP to a JSON array of IPAddresses
    $ip = ([IPAddress]$firstMasterIp).getAddressBytes()
    [Array]::Reverse($ip)
    $ip = ([IPAddress]($ip -join '.')).Address

    $MasterIP = @([IPAddress]$null)
    
    for ($i = 0; $i -lt $MasterCount; $i++ ) 
    {
       $new_ip = ([IPAddress]$ip).getAddressBytes()
       [Array]::Reverse($new_ip)
       $new_ip = [IPAddress]($new_ip -join '.')
       $MasterIP += $new_ip
      
       $ip++
     
    }
    $master_str  = $MasterIP.IPAddressToString

    # Add the port numbers
    if ($master_str.count -eq 1) {
        $master_str += ":2181"
    }
    else {
        for ($i = 0; $i -lt $master_str.count; $i++) 
        {
            $master_str[$i] += ":2181"
        }
    }
    $master_json = ConvertTo-Json $master_str
    $master_json = $master_json -replace [Environment]::NewLine,""

    $private_ip = ( Get-NetIPAddress | where { $_.AddressFamily -eq "IPv4" } | where { Check-Subnet $subnet $_.IPAddress } )  # We know the subnet we are on. Makes it easier and more robust
    [Environment]::SetEnvironmentVariable("DCOS_AGENT_IP", $private_ip.IPAddress, "Machine")

    if ($isAgent)
    {
        $run_cmd = $global:BootstrapInstallDir+"\DCOSWindowsAgentSetup.ps1 -MasterIP '$master_json' -AgentPrivateIP "+($private_ip.IPAddress) +" -BootstrapUrl '$bootstrapUri' " 
        if ($isPublic) 
        {
            $run_cmd += " -isPublic:` + "`" + `$true "
        }
        if ($customAttrs) 
        {
            $run_cmd += " -customAttrs '$customAttrs'"
        }
        $run_cmd += ">"+$global:BootstrapInstallDir+"\DCOSWindowsAgentSetup.log 2>&1"
        Write-Log "run setup script $run_cmd"
        Invoke-Expression $run_cmd
        Write-Log "setup script completed"
    }
    else # We must be deploying a master
    {
        $run_cmd = $global:BootstrapInstallDir+"\DCOSWindowsMasterSetup.ps1 -MasterIP '$master_json' -MasterPrivateIP $privateIP.IPAddress -BootstrapUrl '$bootstrapUri'"
        Write-Log "run setup script $run_cmd"
        Invoke-Expression $run_cmd
    }

    PREPROVISION_EXTENSION

    Write-Log "Provisioning script succeeded"
}
catch
{
    Write-Log "Provisioning script failed"
    Write-Error $_
    exit 1
}
`)

func dcosDcoswindowsprovisionPs1Bytes() ([]byte, error) {
	return _dcosDcoswindowsprovisionPs1, nil
}

func dcosDcoswindowsprovisionPs1() (*asset, error) {
	bytes, err := dcosDcoswindowsprovisionPs1Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/dcosWindowsProvision.ps1", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosDcosagentresourcesvmasT = []byte(`    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}NSGName')]",
      "properties": {
        "securityRules": [
            {{GetSecurityRules .Ports}}
        ]
      },
      "type": "Microsoft.Network/networkSecurityGroups"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "copy": {
        "count": "[sub(variables('{{.Name}}Count'), variables('{{.Name}}Offset'))]",
        "name": "loop"
      },
      "dependsOn": [
{{if .IsCustomVNET}}
      "[concat('Microsoft.Network/networkSecurityGroups/', variables('{{.Name}}NSGName'))]"
{{else}}
      "[variables('vnetID')]"
{{end}}
{{if IsPublic .Ports}}
	  ,"[variables('{{.Name}}LbID')]"
{{end}}
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('{{.Name}}VMNamePrefix'), 'nic-', copyIndex(variables('{{.Name}}Offset')))]",
      "properties": {
{{if .IsCustomVNET}}
	    "networkSecurityGroup": {
		  "id": "[resourceId('Microsoft.Network/networkSecurityGroups/', variables('{{.Name}}NSGName'))]"
	    },
{{end}}
        "ipConfigurations": [
          {
            "name": "ipConfigNode",
            "properties": {
{{if IsPublic .Ports}}
              "loadBalancerBackendAddressPools": [
		        {
		      	  "id": "[concat('/subscriptions/', subscription().subscriptionId,'/resourceGroups/', resourceGroup().name, '/providers/Microsoft.Network/loadBalancers/', variables('{{.Name}}LbName'), '/backendAddressPools/',variables('{{.Name}}LbBackendPoolName'))]"
		        }
		      ],
{{end}}
              "privateIPAllocationMethod": "Dynamic",
              "subnet": {
                "id": "[variables('{{.Name}}VnetSubnetID')]"
             }
            }
          }
        ]
      },
      "type": "Microsoft.Network/networkInterfaces"
    },
{{if .IsManagedDisks}}
    {
      "apiVersion": "[variables('apiVersionStorageManagedDisks')]",
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}AvailabilitySet')]",
      "properties": {
        "platformFaultDomainCount": 2,
        "platformUpdateDomainCount": 3,
        "managed": "true"
      },
      "type": "Microsoft.Compute/availabilitySets"
    },
{{else if .IsStorageAccount}}
    {
      "apiVersion": "[variables('apiVersionStorage')]",
      "copy": {
        "count": "[variables('{{.Name}}StorageAccountsCount')]",
        "name": "loop"
      },
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('storageAccountPrefixes')[mod(add(copyIndex(),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(copyIndex(),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
      "properties": {
        "accountType": "[variables('vmSizesMap')[variables('{{.Name}}VMSize')].storageAccountType]"
      },
      "type": "Microsoft.Storage/storageAccounts"
    },
    {{if .HasDisks}}
        {
          "apiVersion": "[variables('apiVersionStorage')]",
          "copy": {
            "count": "[variables('{{.Name}}StorageAccountsCount')]",
            "name": "datadiskLoop"
          },
          "dependsOn": [
            "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
          ],
          "location": "[variables('location')]",
          "name": "[concat(variables('storageAccountPrefixes')[mod(add(copyIndex(variables('dataStorageAccountPrefixSeed')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(copyIndex(variables('dataStorageAccountPrefixSeed')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}DataAccountName'))]",
          "properties": {
            "accountType": "[variables('vmSizesMap')[variables('{{.Name}}VMSize')].storageAccountType]"
          },
          "type": "Microsoft.Storage/storageAccounts"
        },
    {{end}}
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}AvailabilitySet')]",
      "properties": {},
      "type": "Microsoft.Compute/availabilitySets"
    },
{{end}}
{{if IsPublic .Ports}}
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}IPAddressName')]",
      "properties": {
        "dnsSettings": {
          "domainNameLabel": "[variables('{{.Name}}EndpointDNSNamePrefix')]"
        },
        "publicIPAllocationMethod": "Dynamic"
      },
      "type": "Microsoft.Network/publicIPAddresses"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('{{.Name}}IPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}LbName')]",
      "properties": {
        "backendAddressPools": [
          {
            "name": "[variables('{{.Name}}LbBackendPoolName')]"
          }
        ],
        "frontendIPConfigurations": [
          {
            "name": "[variables('{{.Name}}LbIPConfigName')]",
            "properties": {
              "publicIPAddress": {
                "id": "[resourceId('Microsoft.Network/publicIPAddresses',variables('{{.Name}}IPAddressName'))]"
              }
            }
          }
        ],
        "inboundNatRules": [],
        "loadBalancingRules": [
            {{(GetLBRules .Name .Ports)}}
        ],
        "probes": [
            {{(GetProbes .Ports)}}
        ]
      },
      "type": "Microsoft.Network/loadBalancers"
    },
{{end}}
    {
{{if .IsManagedDisks}}
      "apiVersion": "[variables('apiVersionStorageManagedDisks')]",
{{else}}
      "apiVersion": "[variables('apiVersionDefault')]",
{{end}}
      "copy": {
        "count": "[sub(variables('{{.Name}}Count'), variables('{{.Name}}Offset'))]",
        "name": "vmLoopNode"
      },
      "dependsOn": [
{{if .IsStorageAccount}}
        "[concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
  {{if .HasDisks}}
        "[concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('dataStorageAccountPrefixSeed')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('dataStorageAccountPrefixSeed')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}DataAccountName'))]",
  {{end}}
{{end}}
        "[concat('Microsoft.Network/networkInterfaces/', variables('{{.Name}}VMNamePrefix'), 'nic-', copyIndex(variables('{{.Name}}Offset')))]",
        "[concat('Microsoft.Compute/availabilitySets/', variables('{{.Name}}AvailabilitySet'))]"
{{if HasBootstrap}}
       ,"[concat('Microsoft.Compute/virtualMachines/', variables('bootstrapVMName'), /extensions/bootstrapready')]"
{{end}}
      ],
      "tags":
      {
        "creationSource" : "[concat('acsengine-', variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')))]",
        "orchestratorName": "dcos",
        "orchestratorVersion": "[variables('orchestratorVersion')]",
        "orchestratorNode": "agent"
      },
      "location": "[variables('location')]",
      "name": "[concat(variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')))]",
      "properties": {
        "availabilitySet": {
          "id": "[resourceId('Microsoft.Compute/availabilitySets',variables('{{.Name}}AvailabilitySet'))]"
        },
        "hardwareProfile": {
          "vmSize": "[variables('{{.Name}}VMSize')]"
        },
        "networkProfile": {
          "networkInterfaces": [
            {
              "id": "[resourceId('Microsoft.Network/networkInterfaces',concat(variables('{{.Name}}VMNamePrefix'), 'nic-', copyIndex(variables('{{.Name}}Offset'))))]"
            }
          ]
        },
        "osProfile": {
          "adminUsername": "[variables('adminUsername')]",
          "computername": "[concat(variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')))]",
          {{GetDCOSAgentCustomData .}}
          "linuxConfiguration": {
              "disablePasswordAuthentication": true,
              "ssh": {
                "publicKeys": [
                  {
                    "keyData": "[parameters('sshRSAPublicKey')]",
                    "path": "[variables('sshKeyPath')]"
                  }
                ]
              }
            }
            {{if HasLinuxSecrets}}
              ,
              "secrets": "[variables('linuxProfileSecrets')]"
            {{end}}
        },
        "storageProfile": {
          {{GetDataDisks .}}
          "imageReference": {
            "offer": "[variables('osImageOffer')]",
            "publisher": "[variables('osImagePublisher')]",
            "sku": "[variables('osImageSKU')]",
            "version": "[variables('osImageVersion')]"
          }
          ,"osDisk": {
            "caching": "ReadOnly"
            ,"createOption": "FromImage"
{{if .IsStorageAccount}}
            ,"name": "[concat(variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')),'-osdisk')]"
            ,"vhd": {
              "uri": "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')),variables('apiVersionStorage')).primaryEndpoints.blob,'osdisk/', variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')), '-osdisk.vhd')]"
            }
{{end}}
{{if ne .OSDiskSizeGB 0}}
            ,"diskSizeGB": {{.OSDiskSizeGB}}
{{end}}
          }

        }
      },
      "type": "Microsoft.Compute/virtualMachines"
    }
`)

func dcosDcosagentresourcesvmasTBytes() ([]byte, error) {
	return _dcosDcosagentresourcesvmasT, nil
}

func dcosDcosagentresourcesvmasT() (*asset, error) {
	bytes, err := dcosDcosagentresourcesvmasTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/dcosagentresourcesvmas.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosDcosagentresourcesvmssT = []byte(`    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}NSGName')]",
      "properties": {
        "securityRules": [
            {{GetSecurityRules .Ports}}
        ]
      },
      "type": "Microsoft.Network/networkSecurityGroups"
    },
{{if .IsStorageAccount}}
    {
      "apiVersion": "[variables('apiVersionStorage')]",
      "copy": {
        "count": "[variables('{{.Name}}StorageAccountsCount')]",
        "name": "loop"
      },
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('storageAccountPrefixes')[mod(add(copyIndex(),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(copyIndex(),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
      "properties": {
        "accountType": "[variables('vmSizesMap')[variables('{{.Name}}VMSize')].storageAccountType]"
      },
      "type": "Microsoft.Storage/storageAccounts"
    },
{{end}}
{{if IsPublic .Ports}}
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}IPAddressName')]",
      "properties": {
        "dnsSettings": {
          "domainNameLabel": "[variables('{{.Name}}EndpointDNSNamePrefix')]"
        },
        "publicIPAllocationMethod": "Dynamic"
      },
      "type": "Microsoft.Network/publicIPAddresses"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('{{.Name}}IPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}LbName')]",
      "properties": {
        "backendAddressPools": [
          {
            "name": "[variables('{{.Name}}LbBackendPoolName')]"
          }
        ],
        "frontendIPConfigurations": [
          {
            "name": "[variables('{{.Name}}LbIPConfigName')]",
            "properties": {
              "publicIPAddress": {
                "id": "[resourceId('Microsoft.Network/publicIPAddresses',variables('{{.Name}}IPAddressName'))]"
              }
            }
          }
        ],
        "inboundNatRules": [],
        "loadBalancingRules": [
            {{(GetLBRules .Name .Ports)}}
        ],
        "probes": [
            {{(GetProbes .Ports)}}
        ]
      },
      "type": "Microsoft.Network/loadBalancers"
    },
{{end}}
    {
{{if .IsManagedDisks}}
      "apiVersion": "[variables('apiVersionStorageManagedDisks')]",
{{else}}
      "apiVersion": "[variables('apiVersionDefault')]",
{{end}}
      "dependsOn": [
{{if .IsCustomVNET}}
      "[concat('Microsoft.Network/networkSecurityGroups/', variables('{{.Name}}NSGName'))]"
{{else}}
      "[variables('vnetID')]"
{{end}}
{{if .IsStorageAccount}}
        ,"[concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(0,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(0,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
        "[concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(1,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(1,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
        "[concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(2,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(2,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
        "[concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(3,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(3,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
        "[concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(4,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(4,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]"
{{end}}
{{if IsPublic .Ports}}
       ,"[concat('Microsoft.Network/loadBalancers/', variables('{{.Name}}LbName'))]"
{{end}}
{{if HasBootstrap}}
       ,"[concat('Microsoft.Compute/virtualMachines/', variables('bootstrapVMName'), '/extensions/bootstrapready')]"
{{end}}
      ],
      "tags":
      {
        "creationSource" : "[concat('acsengine-', variables('{{.Name}}VMNamePrefix'), 'vmss')]",
        "orchestratorName": "dcos",
        "orchestratorVersion": "[variables('orchestratorVersion')]",
        "orchestratorNode": "agent"
      },
      "location": "[variables('location')]",
      "name": "[concat(variables('{{.Name}}VMNamePrefix'), 'vmss')]",
      "properties": {
        "overprovision": false,
        "upgradePolicy": {
          "mode": "Manual"
        },
        "virtualMachineProfile": {
          "networkProfile": {
            "networkInterfaceConfigurations": [
              {
                "name": "nic",
                "properties": {
{{if .IsCustomVNET}}
                  "networkSecurityGroup": {
                    "id": "[resourceId('Microsoft.Network/networkSecurityGroups/', variables('{{.Name}}NSGName'))]"
                  },
{{end}}
                  "ipConfigurations": [
                    {
                      "name": "nicipconfig",
                      "properties": {
{{if IsPublic .Ports}}
                        "loadBalancerBackendAddressPools": [
                          {
                            "id": "[concat('/subscriptions/', subscription().subscriptionId,'/resourceGroups/', resourceGroup().name, '/providers/Microsoft.Network/loadBalancers/', variables('{{.Name}}LbName'), '/backendAddressPools/',variables('{{.Name}}LbBackendPoolName'))]"
                          }
                        ],
{{end}}
                        "subnet": {
                          "id": "[variables('{{.Name}}VnetSubnetID')]"
                        }
                      }
                    }
                  ],
                  "primary": "true"
                }
              }
            ]
          },
          "osProfile": {
            "adminUsername": "[variables('adminUsername')]",
            "computerNamePrefix": "[variables('{{.Name}}VMNamePrefix')]",
            {{GetDCOSAgentCustomData .}}
            "linuxConfiguration": {
              "disablePasswordAuthentication": true,
              "ssh": {
                "publicKeys": [
                  {
                    "keyData": "[parameters('sshRSAPublicKey')]",
                    "path": "[variables('sshKeyPath')]"
                  }
                ]
              }
            }
            {{if HasLinuxSecrets}}
              ,
              "secrets": "[variables('linuxProfileSecrets')]"
            {{end}}
          },
          "storageProfile": {
            "imageReference": {
              "offer": "[variables('osImageOffer')]",
              "publisher": "[variables('osImagePublisher')]",
              "sku": "[variables('osImageSKU')]",
              "version": "[variables('osImageVersion')]"
            },
            {{GetDataDisks .}}
            "osDisk": {
              "caching": "ReadOnly",
              "createOption": "FromImage"
{{if .IsStorageAccount}}
              ,"name": "vmssosdisk"
              ,"vhdContainers": [
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(0,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(0,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')),variables('apiVersionStorage')).primaryEndpoints.blob,'osdisk')]",
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(1,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(1,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')),variables('apiVersionStorage')).primaryEndpoints.blob,'osdisk')]",
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(2,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(2,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')),variables('apiVersionStorage')).primaryEndpoints.blob,'osdisk')]",
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(3,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(3,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')),variables('apiVersionStorage')).primaryEndpoints.blob,'osdisk')]",
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(4,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(4,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')),variables('apiVersionStorage')).primaryEndpoints.blob,'osdisk')]"

              ]
{{end}}
{{if ne .OSDiskSizeGB 0}}
            ,"diskSizeGB": {{.OSDiskSizeGB}}
{{end}}
            }
          }
        }
      },
      "sku": {
        "capacity": "[variables('{{.Name}}Count')]",
        "name": "[variables('{{.Name}}VMSize')]",
        "tier": "[variables('{{.Name}}VMSizeTier')]"
      },
      "type": "Microsoft.Compute/virtualMachineScaleSets"
    }
`)

func dcosDcosagentresourcesvmssTBytes() ([]byte, error) {
	return _dcosDcosagentresourcesvmssT, nil
}

func dcosDcosagentresourcesvmssT() (*asset, error) {
	bytes, err := dcosDcosagentresourcesvmssTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/dcosagentresourcesvmss.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosDcosagentvarsT = []byte(`    "{{.Name}}Count": "[parameters('{{.Name}}Count')]",
    "{{.Name}}NSGID": "[resourceId('Microsoft.Network/networkSecurityGroups',variables('{{.Name}}NSGName'))]",
    "{{.Name}}NSGName": "[concat(variables('orchestratorName'), '-{{.Name}}-nsg-', variables('nameSuffix'))]",
{{if .IsWindows}}

    "winResourceNamePrefix" : "[substring(variables('nameSuffix'), 0, 5)]",
    {{if IsPublic .Ports}}
        "{{.Name}}VMNamePrefix": "[concat('wp', variables('winResourceNamePrefix'), add(900,variables('{{.Name}}Index')))]",
        "{{.Name}}windowsAgentCustomAttributes": "[concat(' -customAttrs ', variables('doubleSingleQuote'), '{{GetDCOSWindowsAgentCustomNodeAttributes . }}', variables('doubleSingleQuote') )]",
        "{{.Name}}windowsAgentCustomScriptArguments": "[concat('$arguments = ', variables('singleQuote'), '-subnet ', variables('{{.Name}}Subnet'), ' -MasterCount ', variables('masterCount'), ' -firstMasterIP ', parameters('firstConsecutiveStaticIP'), ' -bootstrapUri ', '\"', variables('dcosWindowsBootstrapURL'), '\"', ' -isAgent $true -isPublic $true ',  variables('{{.Name}}windowsAgentCustomAttributes'), ' -preprovisionExtensionParams ', variables('doubleSingleQuote'), '{{GetDCOSWindowsAgentPreprovisionParameters .}}', variables('doubleSingleQuote'),  variables('singleQuote'), ' ; ')]",
    {{else}}
        "{{.Name}}VMNamePrefix": "[concat('w', variables('winResourceNamePrefix'), add(900,variables('{{.Name}}Index')))]",
        "{{.Name}}windowsAgentCustomAttributes": "[concat(' -customAttrs ', variables('doubleSingleQuote'), '{{GetDCOSWindowsAgentCustomNodeAttributes . }}', variables('doubleSingleQuote') )]",
        "{{.Name}}windowsAgentCustomScriptArguments": "[concat('$arguments = ', variables('singleQuote'), '-subnet ', variables('{{.Name}}Subnet'), ' -MasterCount ', variables('masterCount'), ' -firstMasterIP ', parameters('firstConsecutiveStaticIP'), ' -bootstrapUri ', '\"', variables('dcosWindowsBootstrapURL'), '\"', ' -isAgent $true -isPublic $false ',  variables('{{.Name}}windowsAgentCustomAttributes'), ' -preprovisionExtensionParams ', variables('doubleSingleQuote'), '{{GetDCOSWindowsAgentPreprovisionParameters .}}', variables('doubleSingleQuote'), variables('singleQuote'), ' ; ')]",
    {{end}}

    "{{.Name}}windowsAgentCustomScript": "[concat('powershell.exe -ExecutionPolicy Unrestricted -command \"', variables('{{.Name}}windowsAgentCustomScriptArguments'), variables('windowsCustomScriptSuffix'), '\" > %SYSTEMDRIVE%\\AzureData\\dcosWindowsProvision.log 2>&1; exit $LASTEXITCODE')]",

{{else}}
    "{{.Name}}VMNamePrefix": "[concat(variables('orchestratorName'), '-{{.Name}}-', variables('nameSuffix'), '-')]",
{{end}}

    "{{.Name}}VMSize": "[parameters('{{.Name}}VMSize')]",
    "{{.Name}}VMSizeTier": "[split(parameters('{{.Name}}VMSize'),'_')[0]]",
{{if .IsAvailabilitySets}}
    {{if .IsStorageAccount}}
    "{{.Name}}StorageAccountsCount": "[add(div(variables('{{.Name}}Count'), variables('maxVMsPerStorageAccount')), mod(add(mod(variables('{{.Name}}Count'), variables('maxVMsPerStorageAccount')),2), add(mod(variables('{{.Name}}Count'), variables('maxVMsPerStorageAccount')),1)))]",
    "{{.Name}}StorageAccountOffset": "[mul(variables('maxStorageAccountsPerAgent'),variables('{{.Name}}Index'))]",
    {{end}}
    "{{.Name}}AvailabilitySet": "[concat('{{.Name}}-availabilitySet-', variables('nameSuffix'))]",
    "{{.Name}}Offset": "[parameters('{{.Name}}Offset')]",
{{else}}
    {{if .IsStorageAccount}}
    "{{.Name}}StorageAccountsCount": "[variables('maxStorageAccountsPerAgent')]",
    {{end}}
{{end}}
{{if .IsCustomVNET}}
    "{{.Name}}VnetSubnetID": "[parameters('{{.Name}}VnetSubnetID')]",
{{else}}
    "{{.Name}}Subnet": "[parameters('{{.Name}}Subnet')]",
    "{{.Name}}SubnetName": "[concat(variables('orchestratorName'), '-{{.Name}}Subnet')]",
    "{{.Name}}VnetSubnetID": "[concat(variables('vnetID'),'/subnets/',variables('{{.Name}}SubnetName'))]",
{{end}}
{{if IsPublic .Ports}}
    "{{.Name}}EndpointDNSNamePrefix": "[tolower(parameters('{{.Name}}EndpointDNSNamePrefix'))]",
    "{{.Name}}IPAddressName": "[concat(variables('orchestratorName'), '-agent-ip-', variables('{{.Name}}EndpointDNSNamePrefix'), '-', variables('nameSuffix'))]",
    "{{.Name}}LbBackendPoolName": "[concat(variables('orchestratorName'), '-{{.Name}}-', variables('nameSuffix'))]",
    "{{.Name}}LbID": "[resourceId('Microsoft.Network/loadBalancers',variables('{{.Name}}LbName'))]",
    "{{.Name}}LbIPConfigID": "[concat(variables('{{.Name}}LbID'),'/frontendIPConfigurations/', variables('{{.Name}}LbIPConfigName'))]",
    "{{.Name}}LbIPConfigName": "[concat(variables('orchestratorName'), '-{{.Name}}-', variables('nameSuffix'))]",
    "{{.Name}}LbName": "[concat(variables('orchestratorName'), '-{{.Name}}-', variables('nameSuffix'))]",
     {{if .IsWindows}}
        "{{.Name}}WindowsRDPNatRangeStart": 3389,
        "{{.Name}}WindowsRDPEndRangeStop": "[add(variables('{{.Name}}WindowsRDPNatRangeStart'), add(variables('{{.Name}}Count'),variables('{{.Name}}Count')))]",

    {{end}}
{{end}}
{{if HasPrivateRegistry}}
    "registry" : "[tolower(parameters('registry'))]",
    "registryKey" : "[parameters('registryKey')]",
{{else}}
    "registry" : "",
    "registryKey" : "",
{{end}}
`)

func dcosDcosagentvarsTBytes() ([]byte, error) {
	return _dcosDcosagentvarsT, nil
}

func dcosDcosagentvarsT() (*asset, error) {
	bytes, err := dcosDcosagentvarsTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/dcosagentvars.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosDcosbaseT = []byte(`{
  "$schema": "https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
  "contentVersion": "1.0.0.0",
  "parameters": {
    {{range .AgentPoolProfiles}}{{template "agentparams.t" .}},{{end}}
    {{if .HasWindows}}
      "dcosBinariesURL": {
        "metadata": {
          "description": "The download url for dcos/mesos windows binaries."
        },
        "type": "string"
      },
      "dcosBinariesVersion": {
        "metadata": {
          "description": "DCOS windows binaries version"
        },
        "type": "string"
      },
      {{template "windowsparams.t"}},
    {{end}}
    {{template "dcos/dcosparams.t" .}}
    {{template "masterparams.t" .}}
  },
  "variables": {
    "dcosRepositoryURL": "[parameters('dcosRepositoryURL')]",
    "dcosClusterPackageListID": "[parameters('dcosClusterPackageListID')]",
    "dcosProviderPackageID": "[parameters('dcosProviderPackageID')]",
    {{range $index, $agent := .AgentPoolProfiles}}
        "{{.Name}}Index": {{$index}},
        {{template "dcos/dcosagentvars.t" .}}
        {{if .IsStorageAccount}}
          "{{.Name}}StorageAccountOffset": "[mul(variables('maxStorageAccountsPerAgent'),{{$index}})]",
          "{{.Name}}AccountName": "[concat(variables('storageAccountBaseName'), 'agnt{{$index}}')]",
          {{if .HasDisks}}
            "{{.Name}}DataAccountName": "[concat(variables('storageAccountBaseName'), 'data{{$index}}')]",
          {{end}}
        {{end}}
    {{end}}

    {{template "dcos/dcosmastervars.t" .}}
  },
  "resources": [
    {{range .AgentPoolProfiles}}
      {{if .IsWindows}}
        {{if .IsAvailabilitySets}}
          {{template "dcos/dcosWindowsAgentResourcesVmas.t" .}},
        {{else}}
          {{template "dcos/dcosWindowsAgentResourcesVmss.t" .}},
        {{end}}
      {{else}}
        {{if .IsAvailabilitySets}}
          {{template "dcos/dcosagentresourcesvmas.t" .}},
        {{else}}
          {{template "dcos/dcosagentresourcesvmss.t" .}},
        {{end}}
      {{end}}
    {{end}}
    {{template "dcos/dcosmasterresources.t" .}}
  ],
  "outputs": {
    {{range .AgentPoolProfiles}}{{template "agentoutputs.t" .}}
    {{end}}
    {{template "masteroutputs.t" .}}
  }
}
`)

func dcosDcosbaseTBytes() ([]byte, error) {
	return _dcosDcosbaseT, nil
}

func dcosDcosbaseT() (*asset, error) {
	bytes, err := dcosDcosbaseTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/dcosbase.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosDcoscustomdata110T = []byte(`bootcmd:
- bash -c "if [ ! -f /var/lib/sdb-gpt ];then echo DCOS-5890;parted -s /dev/sdb mklabel
  gpt;touch /var/lib/sdb-gpt;fi"
- bash -c "if [ ! -f /var/lib/sdc-gpt ];then echo DCOS-5890;parted -s /dev/sdc mklabel
  gpt&&touch /var/lib/sdc-gpt;fi"
- bash -c "if [ ! -f /var/lib/sdd-gpt ];then echo DCOS-5890;parted -s /dev/sdd mklabel
  gpt&&touch /var/lib/sdd-gpt;fi"
- bash -c "if [ ! -f /var/lib/sde-gpt ];then echo DCOS-5890;parted -s /dev/sde mklabel
  gpt&&touch /var/lib/sde-gpt;fi"
- bash -c "if [ ! -f /var/lib/sdf-gpt ];then echo DCOS-5890;parted -s /dev/sdf mklabel
  gpt&&touch /var/lib/sdf-gpt;fi"
- bash -c "mkdir -p /dcos/volume{0,1,2,3}"
disk_setup:
  ephemeral0:
    layout:
    - 45
    - 45
    - 10
    overwrite: true
    table_type: gpt
  /dev/sdc:
    layout: true
    overwrite: true
    table_type: gpt
  /dev/sdd:
    layout: true
    overwrite: true
    table_type: gpt
  /dev/sde:
    layout: true
    overwrite: true
    table_type: gpt
  /dev/sdf:
    layout: true
    overwrite: true
    table_type: gpt
fs_setup:
- device: ephemeral0.1
  filesystem: ext4
  overwrite: true
- device: ephemeral0.2
  filesystem: ext4
  overwrite: true
- device: ephemeral0.3
  filesystem: ext4
  overwrite: true
- device: /dev/sdc1
  filesystem: ext4
  overwrite: true
- device: /dev/sdd1
  filesystem: ext4
  overwrite: true
- device: /dev/sde1
  filesystem: ext4
  overwrite: true
- device: /dev/sdf1
  filesystem: ext4
  overwrite: true
mounts:
- - ephemeral0.1
  - /var/lib/mesos
- - ephemeral0.2
  - /var/lib/docker
- - ephemeral0.3
  - /var/tmp
- - /dev/sdc1
  - /dcos/volume0
- - /dev/sdd1
  - /dcos/volume1
- - /dev/sde1
  - /dcos/volume2
- - /dev/sdf1
  - /dcos/volume3
runcmd: PREPROVISION_EXTENSION
- - ln
  - -s
  - /bin/rm
  - /usr/bin/rm
- - ln
  - -s
  - /bin/mkdir
  - /usr/bin/mkdir
- - ln
  - -s
  - /bin/tar
  - /usr/bin/tar
- - ln
  - -s
  - /bin/ln
  - /usr/bin/ln
- - ln
  - -s
  - /bin/cp
  - /usr/bin/cp
- - ln
  - -s
  - /bin/systemctl
  - /usr/bin/systemctl
- - ln
  - -s
  - /bin/mount
  - /usr/bin/mount
- - ln
  - -s
  - /bin/bash
  - /usr/bin/bash
- - ln
  - -s
  - /usr/sbin/useradd
  - /usr/bin/useradd
- - systemctl
  - disable
  - --now
  - resolvconf.service
- - systemctl
  - mask
  - --now
  - lxc-net.service
- - systemctl
  - disable
  - --now
  - unscd.service
- - systemctl
  - stop
  - --now
  - unscd.service
- /opt/azure/containers/provision.sh
- - cp
  - -p
  - /etc/resolv.conf
  - /tmp/resolv.conf
- - rm
  - -f
  - /etc/resolv.conf
- - cp
  - -p
  - /tmp/resolv.conf
  - /etc/resolv.conf
- - systemctl
  - start
  - dcos-docker-install.service
- - systemctl
  - start
  - dcos-config-writer.service
- - systemctl
  - restart
  - systemd-journald.service
- - systemctl
  - restart
  - docker.service
- - systemctl
  - start
  - dcos-link-env.service
- - systemctl
  - enable
  - dcos-setup.service
- - systemctl
  - --no-block
  - start
  - dcos-setup.service
write_files:
- content: '{{{dcosRepositoryURL}}}
'
  owner: root
  path: /etc/mesosphere/setup-flags/repository-url
  permissions: '0644'
- content: '["adminrouter--1166a3736442e7963a68d1d644bf5f54ca3cb01d", "avro-cpp--9cb0ee14e3cd5bbdb171efcc72a84d16862ea02d",
    "boost-libs--8d515c2f703c666ae1b6c5ccc35cc0f8fa36677f", "bootstrap--c1bc86593e212cf9fe83db2246bacd129a6b3adc",
    "boto--3890cb2817c00b874ba033abe784b5b343caa3c7", "check-time--79e3f6ab99125471e1d94d5f6bc0fea88446831c",
    "cni--7a8572e385c3f5262945c52c8003d1bbb22cf7aa", "cosmos--e84c5bf3259405df90d682536ba445cc4839a324",
    "curl--17866a8ae9305826aa5f357a09db2c1f2b2c2ad0", "dcos-checks--8fd33919e6f163dba1bd13e4c7e4e0523919a719",
    "dcos-cni--12a77c1e9bebd4cbd600524a864c2bd8483330d3", "dcos-config--setup_{{{dcosProviderPackageID}}}",
    "dcos-diagnostics--e3b557b0ec8e98617d0cd0fdf136ef9dded96316", "dcos-history--23de88ddc1a5f9018dd11b279c5be6a768a18de4",
    "dcos-image--df630d8e930d6650ce3d0ade519660142233d862", "dcos-image-deps--81d23d00b1acddb316c9b15fd8499c2b10f6b697",
    "dcos-integration-test--9ec173650d4e73ba494603324e7583d23970e4b8", "dcos-log--d2af4b1a47d3755a51823e95fbc6c366cf0f9269",
    "dcos-metadata--setup_{{{dcosProviderPackageID}}}", "dcos-metrics--2a26c0b50b0b6564f86c48d50aa86f681c9af93c",
    "dcos-oauth--445bb1388670981c6acc667b2529fc32d4c1fbd4", "dcos-signal--4366023212ea49a64c5c9aef1965e5a3133c4b61",
    "dcos-test-utils--1066d896d25f4c1e3f6d9a5e7f9c1c6e8c675bb7", "dcos-ui--cc2e3d26537ea190efacd6f899dd4cc2210d45b7",
    "dnspython--0be432372a3820eafcfa66975943c9536dbe1164", "docker-gc--89f5535aea154dca504f84cd60eac6f61836aef9",
    "dvdcli--ee85411e3cb9f0988ed54b5cc0789172b887f12f", "erlang--d693172f6f033707c7f07ff78fc18ac543d66b41",
    "exhibitor--c3e48bbae19c0ed9c30d7f9396305d1e77130658", "flask--6d0f985ad677e8422c7190cbe207424acd813c3b",
    "java--ce5ff19502fca31eaf4a9af86d50a10a8c212a5b", "libevent--05dc18bc0ab7434b2738318c5ebaa2e61a311f50",
    "libffi--0e5b99b94f296b2a9a1b75e9fe5f74f5446f5e9b", "libsodium--e7056355f1fe160ade83aac0d11352a2bf3844e6",
    "logrotate--877aece1fd506af3b9167b6938c316adfa79d4f5", "marathon--accdc43bafeca02da1be340baba4b55011eadf63",
    "mesos--0677ce2b7d2e8c45091f6481884542f1f765c3d5", "mesos-dns--600da87080b7634f2380594499004a7ff0b34662",
    "mesos-modules--1f5c4860450949db92ed27326c3146526041e681", "metronome--2ec6f56be44ed822e7228cb66c4dae6a78345789",
    "navstar--c66f92f01d837433de3e2b19d221c64d26cc54b1", "ncurses--030fd6b08ed46a7ecce001c36901f5b4ad5d2af5",
    "octarine--4e37c062d2f145f9c2ce01d30dadf72c2aac5c4a", "openssl--44777d19d54a3c33cc19543f2201cb20bf085d98",
    "pkgpanda-api--30cb1e68f92ed5d4b89d57ca526f8a69b44132c8", "pkgpanda-role--612a6734567cc0c7c2ae1d508f03172f4bc7beed",
    "pytest--5e26c8ed9fd2c325672d56fe558299bfbd0f7018", "python--5a4285ff7296548732203950bf73d360ea67f6ab",
    "python-azure-mgmt-resource--26cbe8349f3fe139f7dc8bff7f0cb735382314fc", "python-cryptography--0d83d8afef4a8faddf0d8b713619d9d76e510a9e",
    "python-dateutil--519201adebeba186049ecd79a9f358f614173b10", "python-docopt--0af809c220a922f7f6c58f15beafebaa043477c7",
    "python-gunicorn--2ceb53716237da0736f67f4004682083f6ac68e1", "python-isodate--c9efb5859a0cfb06d82f25220cc5b387914af85d",
    "python-jinja2--601a1443aa4c649ab1da10c2a6d7a4477a263fb3", "python-kazoo--0ff8e6ef528f58c6f36f0a9df6dc27d3871e5c27",
    "python-markupsafe--1388c95920b4eb920c7a753d620a1ad07fc8b64d", "python-passlib--4691268be760073188b555dc436f836c6706b37a",
    "python-pyyaml--d8a775d6e43da5eb239af5cccdf1d3fceeb0335f", "python-requests--db0474fab16019ba29a609a354285f221c1a2859",
    "python-retrying--37dd25bf69bcbefe0c50139085d6bb2e22ccf439", "python-tox--322c468e2a75c5b143cb06af460b5e801ee34342",
    "rexray--da7f17f8a4b772c0bac3f8d289a08abd4ff272b4", "six--93734bac9907087744815f9cb5b6152e9a198fae",
    "spartan--c3d8005b1340bcbc3a00496861745b2d0bb2d697", "strace--9be573456909e3931a890785eb6474af7e0dcce4",
    "teamcity-messages--073793b16cf369e58ebdb6348b93ed14b0e5e59a", "toybox--0c49f879bfe2f99e6f99b397136894fa5096fa0c"]

'
  owner: root
  path: /etc/mesosphere/setup-flags/cluster-packages.json
  permissions: '0644'
- content: |
    [Journal]
    MaxLevelConsole=warning
    RateLimitInterval=1s
    RateLimitBurst=20000
  owner: root
  path: /etc/systemd/journald.conf.d/dcos.conf
  permissions: '0644'
- content: |
    rexray:
      loglevel: info
      modules:
        default-docker:
          disabled: true
      service: vfs
  path: /etc/rexray/config.yml
  permissions: '0644'
- content: |
    [Unit]
    After=network-online.target
    Wants=network-online.target
    [Service]
    Type=oneshot
    Environment=DEBIAN_FRONTEND=noninteractive
    StandardOutput=journal+console
    StandardError=journal+console
    ExecStartPre=/usr/bin/curl -fLsSv --retry 20 -Y 100000 -y 60 -o /var/tmp/d.deb https://mesosphere.blob.core.windows.net/dcos-deps/docker-engine_1.13.1-0-ubuntu-xenial_amd64.deb
    ExecStart=/usr/bin/bash -c "try=1;until dpkg -D3 -i /var/tmp/d.deb || ((try>9));do echo retry $((try++));sleep $((try*try));done;systemctl --now start docker;systemctl restart docker.socket"
  path: /etc/systemd/system/dcos-docker-install.service
  permissions: '0644'
- content: |
    [Service]
    Restart=always
    StartLimitInterval=0
    RestartSec=15
    ExecStartPre=-/sbin/ip link del docker0
    ExecStart=
    ExecStart=/usr/bin/dockerd --storage-driver=overlay
  path: /etc/systemd/system/docker.service.d/execstart.conf
  permissions: '0644'
- content: |
    [Unit]
    PartOf=docker.service
    [Socket]
    ListenStream=/var/run/docker.sock
    SocketMode=0660
    SocketUser=root
    SocketGroup=docker
    ListenStream=2375
    BindIPv6Only=both
    [Install]
    WantedBy=sockets.target
  path: /etc/systemd/system/docker.socket
  permissions: '0644'
- content: |
      [Unit]
      Requires=dcos-setup.service
      After=dcos-setup.service
      [Service]
      Type=oneshot
      EnvironmentFile=/etc/environment
      EnvironmentFile=/opt/mesosphere/environment
      ExecStart=/usr/bin/bash -c "echo $(detect_ip) $(hostname) > /etc/hosts"
  path: /etc/systemd/system/dcos-config-writer.service
  permissions: '0644'
- content: |
    "bound_values":
      "adminrouter_auth_enabled": |-
        {{{oauthEnabled}}}
      "cluster_name": |-
        {{{masterPublicIPAddressName}}}
      "exhibitor_azure_account_key": |-
        ', listKeys(resourceId('Microsoft.Storage/storageAccounts', variables('masterStorageAccountExhibitorName')), '2015-06-15').key1, '
      "exhibitor_azure_account_name": |-
        {{{masterStorageAccountExhibitorName}}}
      "exhibitor_azure_prefix": |-
        {{{masterPublicIPAddressName}}}
      "master_list": |-
        ["', DCOSCUSTOMDATAPUBLICIPSTR'"]
      "oauth_enabled": |-
        {{{oauthEnabled}}}
    "late_bound_package_id": |-
      dcos-provider-{{{dcosProviderPackageID}}}-azure--setup
  owner: root
  path: /etc/mesosphere/setup-flags/late-config.yaml
  permissions: '0644'
- content: |
    [Unit]
    Before=dcos.target
    [Service]
    Type=oneshot
    StandardOutput=journal+console
    StandardError=journal+console
    ExecStartPre=/usr/bin/mkdir -p /etc/profile.d
    ExecStart=/usr/bin/ln -sf /opt/mesosphere/bin/add_dcos_path.sh /etc/profile.d/dcos.sh
  path: /etc/systemd/system/dcos-link-env.service
  permissions: '0644'
- content: |
    [Unit]
    Description=Pkgpanda: Download DC/OS to this host.
    After=network-online.target
    Wants=network-online.target
    ConditionPathExists=!/opt/mesosphere/
    [Service]
    Type=oneshot
    StandardOutput=journal+console
    StandardError=journal+console
    ExecStartPre=/usr/bin/curl --keepalive-time 2 -fLsSv --retry 20 -Y 100000 -y 60 -o //var/tmp/bootstrap.tar.xz {{{dcosBootstrapURL}}}
    ExecStartPre=/usr/bin/mkdir -p /opt/mesosphere
    ExecStart=/usr/bin/tar -axf //var/tmp/bootstrap.tar.xz -C /opt/mesosphere
    ExecStartPost=-/usr/bin/rm -f //var/tmp/bootstrap.tar.xz
  path: /etc/systemd/system/dcos-download.service
  permissions: '0644'
- content: |
    [Unit]
    Description=Pkgpanda: Specialize DC/OS for this host.
    Requires=dcos-download.service
    After=dcos-download.service
    [Service]
    Type=oneshot
    StandardOutput=journal+console
    StandardError=journal+console
    EnvironmentFile=/opt/mesosphere/environment
    ExecStart=/opt/mesosphere/bin/pkgpanda setup --no-block-systemd
    [Install]
    WantedBy=multi-user.target
  path: /etc/systemd/system/dcos-setup.service
  permissions: '0644'
- content: ''
  path: /etc/mesosphere/roles/azure
- path: /var/lib/dcos/mesos-slave-common
  content: 'ATTRIBUTES_STR'
  permissions: "0644"
  owner: "root"
- content: 'PROVISION_STR'
  path: /opt/azure/containers/provision.sh
  permissions: "0744"
  owner: "root"
`)

func dcosDcoscustomdata110TBytes() ([]byte, error) {
	return _dcosDcoscustomdata110T, nil
}

func dcosDcoscustomdata110T() (*asset, error) {
	bytes, err := dcosDcoscustomdata110TBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/dcoscustomdata110.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosDcoscustomdata184T = []byte(`bootcmd:
- bash -c "if [ ! -f /var/lib/sdb-gpt ];then echo DCOS-5890;parted -s /dev/sdb mklabel
  gpt;touch /var/lib/sdb-gpt;fi"
disk_setup:
  ephemeral0:
    layout:
    - 50
    - 50
    overwrite: true
    table_type: gpt
fs_setup:
- device: ephemeral0.1
  filesystem: ext4
  overwrite: true
- device: ephemeral0.2
  filesystem: ext4
  overwrite: true
mounts:
- - ephemeral0.1
  - /var/lib/mesos
- - ephemeral0.2
  - /var/lib/docker
runcmd: PREPROVISION_EXTENSION
- /usr/lib/apt/apt.systemd.daily
- echo 2dd1ce17-079e-403c-b352-a1921ee207ee > /sys/bus/vmbus/drivers/hv_util/unbind # mitigation for bug https://bugs.launchpad.net/ubuntu/+source/linux/+bug/1676635
- sed -i "13i\echo 2dd1ce17-079e-403c-b352-a1921ee207ee > /sys/bus/vmbus/drivers/hv_util/unbind\n" /etc/rc.local # mitigation for bug https://bugs.launchpad.net/ubuntu/+source/linux/+bug/1676635
- - ln
  - -s
  - /bin/rm
  - /usr/bin/rm
- - ln
  - -s
  - /bin/mkdir
  - /usr/bin/mkdir
- - ln
  - -s
  - /bin/tar
  - /usr/bin/tar
- - ln
  - -s
  - /bin/ln
  - /usr/bin/ln
- - ln
  - -s
  - /bin/cp
  - /usr/bin/cp
- - ln
  - -s
  - /bin/systemctl
  - /usr/bin/systemctl
- - ln
  - -s
  - /bin/mount
  - /usr/bin/mount
- - ln
  - -s
  - /bin/bash
  - /usr/bin/bash
- - ln
  - -s
  - /usr/sbin/useradd
  - /usr/bin/useradd
- - systemctl
  - disable
  - --now
  - resolvconf.service
- - systemctl
  - mask
  - --now
  - lxc-net.service
- - tar
  - czf
  - /etc/docker.tar.gz
  - -C
  - /tmp/xtoph
  - .docker
- - rm
  - -rf
  - /tmp/xtoph
- /opt/azure/containers/provision.sh
- - cp
  - -p
  - /etc/resolv.conf
  - /tmp/resolv.conf
- - rm
  - -f
  - /etc/resolv.conf
- - cp
  - -p
  - /tmp/resolv.conf
  - /etc/resolv.conf
- - systemctl
  - start
  - dcos-docker-install.service
- - systemctl
  - start
  - dcos-config-writer.service
- - systemctl
  - restart
  - systemd-journald.service
- - systemctl
  - restart
  - docker.service
- - systemctl
  - start
  - dcos-link-env.service
- - systemctl
  - enable
  - dcos-setup.service
- - systemctl
  - --no-block
  - start
  - dcos-setup.service
write_files:
- content: 'https://dcosio.azureedge.net/dcos/testing

    '
  owner: root
  path: /etc/mesosphere/setup-flags/repository-url
  permissions: '0644'
- content: 'DCOS_ENVIRONMENT={{{targetEnvironment}}}

    '
  owner: root
  path: /etc/mesosphere/setup-flags/dcos-deploy-environment
  permissions: '0644'
- content: 'BOOTSTRAP_ID=5b4aa43610c57ee1d60b4aa0751a1fb75824c083

    '
  owner: root
  path: /etc/mesosphere/setup-flags/bootstrap-id
  permissions: '0644'
- content: '["dcos-config--setup_{{{dcosProviderPackageID}}}", "dcos-metadata--setup_{{{dcosProviderPackageID}}}"]

    '
  owner: root
  path: /etc/mesosphere/setup-flags/cluster-packages.json
  permissions: '0644'
- content: '[Journal]

    MaxLevelConsole=warning

    RateLimitInterval=1s

    RateLimitBurst=20000

    '
  owner: root
  path: /etc/systemd/journald.conf.d/dcos.conf
  permissions: '0644'
- content: "rexray:\n  loglevel: info\n  modules:\n    default-admin:\n      host:\
    \ tcp://127.0.0.1:61003\n    default-docker:\n      disabled: true\n"
  path: /etc/rexray/config.yml
  permissions: '0644'
- content: '[Unit]

    After=network-online.target

    Wants=network-online.target

    [Service]

    Type=oneshot

    Environment=DEBIAN_FRONTEND=noninteractive

    StandardOutput=journal+console

    StandardError=journal+console

    ExecStart=/usr/bin/bash -c "try=1;until dpkg -D3 -i /var/lib/mesos/dl/d.deb || ((try>9));do
    echo retry $((try++));sleep $((try*try));done;systemctl --now start docker;systemctl
    restart docker.socket"

    '
  path: /etc/systemd/system/dcos-docker-install.service
  permissions: '0644'
- content: '[Service]

    Restart=always

    StartLimitInterval=0

    RestartSec=15

    ExecStartPre=-/sbin/ip link del docker0

    ExecStart=

    ExecStart=/usr/bin/docker daemon -H fd:// --storage-driver=overlay

    '
  path: /etc/systemd/system/docker.service.d/execstart.conf
  permissions: '0644'
- content: '[Unit]

    PartOf=docker.service

    [Socket]

    ListenStream=/var/run/docker.sock

    SocketMode=0660

    SocketUser=root

    SocketGroup=docker

    ListenStream=2375

    BindIPv6Only=both

    [Install]

    WantedBy=sockets.target

    '
  path: /etc/systemd/system/docker.socket
  permissions: '0644'
- content: '[Unit]

    Requires=dcos-setup.service

    After=dcos-setup.service

    [Service]

    Type=oneshot

    EnvironmentFile=/etc/environment

    EnvironmentFile=/opt/mesosphere/environment

    ExecStart=/usr/bin/bash -c "echo $(detect_ip) $(hostname) > /etc/hosts"

    '
  path: /etc/systemd/system/dcos-config-writer.service
  permissions: '0644'
- content: 'MESOS_CLUSTER={{{masterPublicIPAddressName}}}

    '
  path: /etc/mesosphere/setup-packages/dcos-provider-azure--setup/etc/mesos-master-provider
- content: 'ADMINROUTER_ACTIVATE_AUTH_MODULE={{{oauthEnabled}}}

    '
  path: /etc/mesosphere/setup-packages/dcos-provider-azure--setup/etc/adminrouter.env
- content: '["'', DCOSCUSTOMDATAPUBLICIPSTR''"]

    '
  path: /etc/mesosphere/setup-packages/dcos-provider-azure--setup/etc/master_list
- content: 'EXHIBITOR_BACKEND=AZURE

    AZURE_CONTAINER=dcos-exhibitor

    AZURE_PREFIX={{{masterPublicIPAddressName}}}

    '
  path: /etc/mesosphere/setup-packages/dcos-provider-azure--setup/etc/exhibitor
- content: 'com.netflix.exhibitor.azure.account-name={{{masterStorageAccountExhibitorName}}}

    com.netflix.exhibitor.azure.account-key='', listKeys(resourceId(''Microsoft.Storage/storageAccounts'',
    variables(''masterStorageAccountExhibitorName'')), ''2015-06-15'').key1,''

    '
  path: /etc/mesosphere/setup-packages/dcos-provider-azure--setup/etc/exhibitor.properties
- content: '{"uiConfiguration":{"plugins":{"banner":{"enabled":false,"backgroundColor":"#1E232F","foregroundColor":"#FFFFFF","headerTitle":null,"headerContent":null,"footerContent":null,"imagePath":null,"dismissible":null},"branding":{"enabled":false},"external-links":
    {"enabled": false},


    "authentication":{"enabled":false},


    "oauth":{"enabled":{{{oauthEnabled}}},"authHost":"https://dcos.auth0.com"},



    "tracking":{"enabled":false}}}}

    '
  path: /etc/mesosphere/setup-packages/dcos-provider-azure--setup/etc/ui-config.json
- content: '{}'
  path: /etc/mesosphere/setup-packages/dcos-provider-azure--setup/pkginfo.json
- content: '[Unit]

    Before=dcos.target

    [Service]

    Type=oneshot

    StandardOutput=journal+console

    StandardError=journal+console

    ExecStartPre=/usr/bin/mkdir -p /etc/profile.d

    ExecStart=/usr/bin/ln -sf /opt/mesosphere/environment.export /etc/profile.d/dcos.sh

    '
  path: /etc/systemd/system/dcos-link-env.service
  permissions: '0644'
- content: '[Unit]

    Description=Pkgpanda: Download DC/OS to this host.

    After=network-online.target

    Wants=network-online.target

    ConditionPathExists=!/opt/mesosphere/

    [Service]

    EnvironmentFile=/etc/mesosphere/setup-flags/bootstrap-id

    Type=oneshot

    StandardOutput=journal+console

    StandardError=journal+console

    ExecStartPre=/usr/bin/curl --keepalive-time 2 -fLsSv --retry 20 -Y 100000 -y 60
    -o /var/lib/mesos/dl/bootstrap.tar.xz {{{dcosBootstrapURL}}}

    ExecStartPre=/usr/bin/mkdir -p /opt/mesosphere

    ExecStart=/usr/bin/tar -axf /var/lib/mesos/dl/bootstrap.tar.xz -C /opt/mesosphere

    ExecStartPost=-/usr/bin/rm -f /var/lib/mesos/dl/bootstrap.tar.xz

    '
  path: /etc/systemd/system/dcos-download.service
  permissions: '0644'
- content: '[Unit]

    Description=Pkgpanda: Specialize DC/OS for this host.

    Requires=dcos-download.service

    After=dcos-download.service

    [Service]

    Type=oneshot

    StandardOutput=journal+console

    StandardError=journal+console

    EnvironmentFile=/opt/mesosphere/environment

    ExecStart=/opt/mesosphere/bin/pkgpanda setup --no-block-systemd

    [Install]

    WantedBy=multi-user.target

    '
  path: /etc/systemd/system/dcos-setup.service
  permissions: '0644'
- path: /var/lib/dcos/mesos-slave-common
  content: 'ATTRIBUTES_STR'
- content: ''
  path: /etc/mesosphere/roles/azure
- content: 'PROVISION_STR'
  path: "/opt/azure/containers/provision.sh"
  permissions: "0744"
  owner: "root"
- content: '{ "auths": { "{{{registry}}}": { "auth" : "{{{registryKey}}}" } } }'
  path: "/tmp/xtoph/.docker/config.json"
  owner: "root"
`)

func dcosDcoscustomdata184TBytes() ([]byte, error) {
	return _dcosDcoscustomdata184T, nil
}

func dcosDcoscustomdata184T() (*asset, error) {
	bytes, err := dcosDcoscustomdata184TBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/dcoscustomdata184.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosDcoscustomdata187T = []byte(`bootcmd:
- bash -c "if [ ! -f /var/lib/sdb-gpt ];then echo DCOS-5890;parted -s /dev/sdb mklabel
  gpt;touch /var/lib/sdb-gpt;fi"
disk_setup:
  ephemeral0:
    layout:
    - 50
    - 50
    overwrite: true
    table_type: gpt
fs_setup:
- device: ephemeral0.1
  filesystem: ext4
  overwrite: true
- device: ephemeral0.2
  filesystem: ext4
  overwrite: true
mounts:
- - ephemeral0.1
  - /var/lib/mesos
- - ephemeral0.2
  - /var/lib/docker
runcmd: PREPROVISION_EXTENSION
- /usr/lib/apt/apt.systemd.daily
- echo 2dd1ce17-079e-403c-b352-a1921ee207ee > /sys/bus/vmbus/drivers/hv_util/unbind # mitigation for bug https://bugs.launchpad.net/ubuntu/+source/linux/+bug/1676635
- sed -i "13i\echo 2dd1ce17-079e-403c-b352-a1921ee207ee > /sys/bus/vmbus/drivers/hv_util/unbind\n" /etc/rc.local # mitigation for bug https://bugs.launchpad.net/ubuntu/+source/linux/+bug/1676635
- - ln
  - -s
  - /bin/rm
  - /usr/bin/rm
- - ln
  - -s
  - /bin/mkdir
  - /usr/bin/mkdir
- - ln
  - -s
  - /bin/tar
  - /usr/bin/tar
- - ln
  - -s
  - /bin/ln
  - /usr/bin/ln
- - ln
  - -s
  - /bin/cp
  - /usr/bin/cp
- - ln
  - -s
  - /bin/systemctl
  - /usr/bin/systemctl
- - ln
  - -s
  - /bin/mount
  - /usr/bin/mount
- - ln
  - -s
  - /bin/bash
  - /usr/bin/bash
- - ln
  - -s
  - /usr/sbin/useradd
  - /usr/bin/useradd
- - systemctl
  - disable
  - --now
  - resolvconf.service
- - systemctl
  - mask
  - --now
  - lxc-net.service
- - tar
  - czf 
  - /etc/docker.tar.gz
  - -C
  - /tmp/xtoph
  - .docker
- - rm 
  - -rf 
  - /tmp/xtoph
- /opt/azure/containers/provision.sh
- - cp
  - -p
  - /etc/resolv.conf
  - /tmp/resolv.conf
- - rm
  - -f
  - /etc/resolv.conf
- - cp
  - -p
  - /tmp/resolv.conf
  - /etc/resolv.conf
- - systemctl
  - start
  - dcos-docker-install.service
- - systemctl
  - start
  - dcos-config-writer.service
- - systemctl
  - restart
  - systemd-journald.service
- - systemctl
  - restart
  - docker.service
- - systemctl
  - start
  - dcos-link-env.service
- - systemctl
  - enable
  - dcos-setup.service
- - systemctl
  - --no-block
  - start
  - dcos-setup.service
write_files:
- content: 'https://dcosio.azureedge.net/dcos/stable

    '
  owner: root
  path: /etc/mesosphere/setup-flags/repository-url
  permissions: '0644'
- content: '["dcos-config--setup_{{{dcosProviderPackageID}}}", "dcos-metadata--setup_{{{dcosProviderPackageID}}}"]

    '
  owner: root
  path: /etc/mesosphere/setup-flags/cluster-packages.json
  permissions: '0644'
- content: 'DCOS_ENVIRONMENT={{{targetEnvironment}}}

    '
  owner: root
  path: /etc/mesosphere/setup-flags/dcos-deploy-environment
  permissions: '0644'
- content: '[Journal]

    MaxLevelConsole=warning

    RateLimitInterval=1s

    RateLimitBurst=20000

    '
  owner: root
  path: /etc/systemd/journald.conf.d/dcos.conf
  permissions: '0644'
- content: "rexray:\n  loglevel: info\n  modules:\n    default-admin:\n      host:\
    \ tcp://127.0.0.1:61003\n    default-docker:\n      disabled: true\n"
  path: /etc/rexray/config.yml
  permissions: '0644'
- content: '[Unit]

    After=network-online.target

    Wants=network-online.target

    [Service]

    Type=oneshot

    Environment=DEBIAN_FRONTEND=noninteractive

    StandardOutput=journal+console

    StandardError=journal+console

    ExecStart=/usr/bin/bash -c "try=1;until dpkg -D3 -i /var/lib/mesos/dl/d.deb || ((try>9));do
    echo retry $((try++));sleep $((try*try));done;systemctl --now start docker;systemctl
    restart docker.socket"

    '
  path: /etc/systemd/system/dcos-docker-install.service
  permissions: '0644'
- content: '[Service]

    Restart=always

    StartLimitInterval=0

    RestartSec=15

    ExecStartPre=-/sbin/ip link del docker0

    ExecStart=

    ExecStart=/usr/bin/docker daemon -H fd:// --storage-driver=overlay

    '
  path: /etc/systemd/system/docker.service.d/execstart.conf
  permissions: '0644'
- content: '[Unit]

    PartOf=docker.service

    [Socket]

    ListenStream=/var/run/docker.sock

    SocketMode=0660

    SocketUser=root

    SocketGroup=docker

    ListenStream=2375

    BindIPv6Only=both

    [Install]

    WantedBy=sockets.target

    '
  path: /etc/systemd/system/docker.socket
  permissions: '0644'
- content: '[Unit]

    Requires=dcos-setup.service

    After=dcos-setup.service

    [Service]

    Type=oneshot

    EnvironmentFile=/etc/environment

    EnvironmentFile=/opt/mesosphere/environment

    ExecStart=/usr/bin/bash -c "echo $(detect_ip) $(hostname) > /etc/hosts"

    '
  path: /etc/systemd/system/dcos-config-writer.service
  permissions: '0644'
- content: 'MESOS_CLUSTER={{{masterPublicIPAddressName}}}

    '
  path: /etc/mesosphere/setup-packages/dcos-provider-azure--setup/etc/mesos-master-provider
- content: 'ADMINROUTER_ACTIVATE_AUTH_MODULE={{{oauthEnabled}}}

    '
  path: /etc/mesosphere/setup-packages/dcos-provider-azure--setup/etc/adminrouter.env
- content: '["'', DCOSCUSTOMDATAPUBLICIPSTR''"]

    '
  path: /etc/mesosphere/setup-packages/dcos-provider-azure--setup/etc/master_list
- content: 'EXHIBITOR_BACKEND=AZURE

    AZURE_CONTAINER=dcos-exhibitor

    AZURE_PREFIX={{{masterPublicIPAddressName}}}

    '
  path: /etc/mesosphere/setup-packages/dcos-provider-azure--setup/etc/exhibitor
- content: 'com.netflix.exhibitor.azure.account-name={{{masterStorageAccountExhibitorName}}}

    com.netflix.exhibitor.azure.account-key='', listKeys(resourceId(''Microsoft.Storage/storageAccounts'',
    variables(''masterStorageAccountExhibitorName'')), ''2015-06-15'').key1,''

    '
  path: /etc/mesosphere/setup-packages/dcos-provider-azure--setup/etc/exhibitor.properties
- content: '{"uiConfiguration":{"plugins":{"banner":{"enabled":false,"backgroundColor":"#1E232F","foregroundColor":"#FFFFFF","headerTitle":null,"headerContent":null,"footerContent":null,"imagePath":null,"dismissible":null},"branding":{"enabled":false},"external-links":
    {"enabled": false},


    "authentication":{"enabled":false},


    "oauth":{"enabled":{{{oauthEnabled}}},"authHost":"https://dcos.auth0.com"},



    "tracking":{"enabled":false}}}}

    '
  path: /etc/mesosphere/setup-packages/dcos-provider-azure--setup/etc/ui-config.json
- content: '{}'
  path: /etc/mesosphere/setup-packages/dcos-provider-azure--setup/pkginfo.json
- content: '[Unit]

    Before=dcos.target

    [Service]

    Type=oneshot

    StandardOutput=journal+console

    StandardError=journal+console

    ExecStartPre=/usr/bin/mkdir -p /etc/profile.d

    ExecStart=/usr/bin/ln -sf /opt/mesosphere/environment.export /etc/profile.d/dcos.sh

    '
  path: /etc/systemd/system/dcos-link-env.service
  permissions: '0644'
- content: '[Unit]

    Description=Pkgpanda: Download DC/OS to this host.

    After=network-online.target

    Wants=network-online.target

    ConditionPathExists=!/opt/mesosphere/

    [Service]

    Type=oneshot

    StandardOutput=journal+console

    StandardError=journal+console

    ExecStartPre=/usr/bin/curl --keepalive-time 2 -fLsSv --retry 20 -Y 100000 -y 60 -o /var/lib/mesos/dl/bootstrap.tar.xz {{{dcosBootstrapURL}}}

    ExecStartPre=/usr/bin/mkdir -p /opt/mesosphere

    ExecStart=/usr/bin/tar -axf /var/lib/mesos/dl/bootstrap.tar.xz -C /opt/mesosphere

    ExecStartPost=-/usr/bin/rm -f /var/lib/mesos/dl/bootstrap.tar.xz

    '
  path: /etc/systemd/system/dcos-download.service
  permissions: '0644'
- content: '[Unit]

    Description=Pkgpanda: Specialize DC/OS for this host.

    Requires=dcos-download.service

    After=dcos-download.service

    [Service]

    Type=oneshot

    StandardOutput=journal+console

    StandardError=journal+console

    EnvironmentFile=/opt/mesosphere/environment

    ExecStart=/opt/mesosphere/bin/pkgpanda setup --no-block-systemd

    [Install]

    WantedBy=multi-user.target

    '
  path: /etc/systemd/system/dcos-setup.service
  permissions: '0644'
- content: ''
  path: /etc/mesosphere/roles/azure
- content: 'PROVISION_STR'
  path: "/opt/azure/containers/provision.sh"
  permissions: "0744"
  owner: "root"
- path: /var/lib/dcos/mesos-slave-common
  content: 'ATTRIBUTES_STR'
  permissions: "0644"
  owner: "root"
- content: '{ "auths": { "{{{registry}}}": { "auth" : "{{{registryKey}}}" } } }'
  path: "/tmp/xtoph/.docker/config.json"
  owner: "root"
`)

func dcosDcoscustomdata187TBytes() ([]byte, error) {
	return _dcosDcoscustomdata187T, nil
}

func dcosDcoscustomdata187T() (*asset, error) {
	bytes, err := dcosDcoscustomdata187TBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/dcoscustomdata187.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosDcoscustomdata188T = []byte(`bootcmd:
- bash -c "if [ ! -f /var/lib/sdb-gpt ];then echo DCOS-5890;parted -s /dev/sdb mklabel
  gpt;touch /var/lib/sdb-gpt;fi"
disk_setup:
  ephemeral0:
    layout:
    - 50
    - 50
    overwrite: true
    table_type: gpt
fs_setup:
- device: ephemeral0.1
  filesystem: ext4
  overwrite: true
- device: ephemeral0.2
  filesystem: ext4
  overwrite: true
mounts:
- - ephemeral0.1
  - /var/lib/mesos
- - ephemeral0.2
  - /var/lib/docker
runcmd: PREPROVISION_EXTENSION
- /usr/lib/apt/apt.systemd.daily
- echo 2dd1ce17-079e-403c-b352-a1921ee207ee > /sys/bus/vmbus/drivers/hv_util/unbind # mitigation for bug https://bugs.launchpad.net/ubuntu/+source/linux/+bug/1676635
- sed -i "13i\echo 2dd1ce17-079e-403c-b352-a1921ee207ee > /sys/bus/vmbus/drivers/hv_util/unbind\n" /etc/rc.local # mitigation for bug https://bugs.launchpad.net/ubuntu/+source/linux/+bug/1676635
- - ln
  - -s
  - /bin/rm
  - /usr/bin/rm
- - ln
  - -s
  - /bin/mkdir
  - /usr/bin/mkdir
- - ln
  - -s
  - /bin/tar
  - /usr/bin/tar
- - ln
  - -s
  - /bin/ln
  - /usr/bin/ln
- - ln
  - -s
  - /bin/cp
  - /usr/bin/cp
- - ln
  - -s
  - /bin/systemctl
  - /usr/bin/systemctl
- - ln
  - -s
  - /bin/mount
  - /usr/bin/mount
- - ln
  - -s
  - /bin/bash
  - /usr/bin/bash
- - ln
  - -s
  - /usr/sbin/useradd
  - /usr/bin/useradd
- - systemctl
  - disable
  - --now
  - resolvconf.service
- - systemctl
  - mask
  - --now
  - lxc-net.service
- - tar
  - czf 
  - /etc/docker.tar.gz
  - -C
  - /tmp/xtoph
  - .docker
- - rm 
  - -rf 
  - /tmp/xtoph
- /opt/azure/containers/provision.sh
- - cp
  - -p
  - /etc/resolv.conf
  - /tmp/resolv.conf
- - rm
  - -f
  - /etc/resolv.conf
- - cp
  - -p
  - /tmp/resolv.conf
  - /etc/resolv.conf
- - systemctl
  - start
  - dcos-docker-install.service
- - systemctl
  - start
  - dcos-config-writer.service
- - systemctl
  - restart
  - systemd-journald.service
- - systemctl
  - restart
  - docker.service
- - systemctl
  - start
  - dcos-link-env.service
- - systemctl
  - enable
  - dcos-setup.service
- - systemctl
  - --no-block
  - start
  - dcos-setup.service
write_files:
- content: 'https://dcosio.azureedge.net/dcos/stable

    '
  owner: root
  path: /etc/mesosphere/setup-flags/repository-url
  permissions: '0644'
- content: 'DCOS_ENVIRONMENT={{{targetEnvironment}}}

    '
  owner: root
  path: /etc/mesosphere/setup-flags/dcos-deploy-environment
  permissions: '0644'
- content: '["dcos-config--setup_{{{dcosProviderPackageID}}}", "dcos-metadata--setup_{{{dcosProviderPackageID}}}"]

    '
  owner: root
  path: /etc/mesosphere/setup-flags/cluster-packages.json
  permissions: '0644'
- content: '[Journal]

    MaxLevelConsole=warning

    RateLimitInterval=1s

    RateLimitBurst=20000

    '
  owner: root
  path: /etc/systemd/journald.conf.d/dcos.conf
  permissions: '0644'
- content: "rexray:\n  loglevel: info\n  modules:\n    default-admin:\n      host:\
    \ tcp://127.0.0.1:61003\n    default-docker:\n      disabled: true\n"
  path: /etc/rexray/config.yml
  permissions: '0644'
- content: '[Unit]

    After=network-online.target

    Wants=network-online.target

    [Service]

    Type=oneshot

    Environment=DEBIAN_FRONTEND=noninteractive

    StandardOutput=journal+console

    StandardError=journal+console

    ExecStart=/usr/bin/bash -c "try=1;until dpkg -D3 -i /var/lib/mesos/dl/d.deb || ((try>9));do
    echo retry $((try++));sleep $((try*try));done;systemctl --now start docker;systemctl
    restart docker.socket"

    '
  path: /etc/systemd/system/dcos-docker-install.service
  permissions: '0644'
- content: '[Service]

    Restart=always

    StartLimitInterval=0

    RestartSec=15

    ExecStartPre=-/sbin/ip link del docker0

    ExecStart=

    ExecStart=/usr/bin/docker daemon -H fd:// --storage-driver=overlay

    '
  path: /etc/systemd/system/docker.service.d/execstart.conf
  permissions: '0644'
- content: '[Unit]

    PartOf=docker.service

    [Socket]

    ListenStream=/var/run/docker.sock

    SocketMode=0660

    SocketUser=root

    SocketGroup=docker

    ListenStream=2375

    BindIPv6Only=both

    [Install]

    WantedBy=sockets.target

    '
  path: /etc/systemd/system/docker.socket
  permissions: '0644'
- content: '[Unit]

    Requires=dcos-setup.service

    After=dcos-setup.service

    [Service]

    Type=oneshot

    EnvironmentFile=/etc/environment

    EnvironmentFile=/opt/mesosphere/environment

    ExecStart=/usr/bin/bash -c "echo $(detect_ip) $(hostname) > /etc/hosts"

    '
  path: /etc/systemd/system/dcos-config-writer.service
  permissions: '0644'
- content: 'MESOS_CLUSTER={{{masterPublicIPAddressName}}}

    '
  path: /etc/mesosphere/setup-packages/dcos-provider-azure--setup/etc/mesos-master-provider
- content: 'ADMINROUTER_ACTIVATE_AUTH_MODULE={{{oauthEnabled}}}

    '
  path: /etc/mesosphere/setup-packages/dcos-provider-azure--setup/etc/adminrouter.env
- content: '["'', DCOSCUSTOMDATAPUBLICIPSTR''"]

    '
  path: /etc/mesosphere/setup-packages/dcos-provider-azure--setup/etc/master_list
- content: 'EXHIBITOR_BACKEND=AZURE

    AZURE_CONTAINER=dcos-exhibitor

    AZURE_PREFIX={{{masterPublicIPAddressName}}}

    '
  path: /etc/mesosphere/setup-packages/dcos-provider-azure--setup/etc/exhibitor
- content: 'com.netflix.exhibitor.azure.account-name={{{masterStorageAccountExhibitorName}}}

    com.netflix.exhibitor.azure.account-key='', listKeys(resourceId(''Microsoft.Storage/storageAccounts'',
    variables(''masterStorageAccountExhibitorName'')), ''2015-06-15'').key1,''

    '
  path: /etc/mesosphere/setup-packages/dcos-provider-azure--setup/etc/exhibitor.properties
- content: '{"uiConfiguration":{"plugins":{"banner":{"enabled":false,"backgroundColor":"#1E232F","foregroundColor":"#FFFFFF","headerTitle":null,"headerContent":null,"footerContent":null,"imagePath":null,"dismissible":null},"branding":{"enabled":false},"external-links":
    {"enabled": false},


    "authentication":{"enabled":false},


    "oauth":{"enabled":{{{oauthEnabled}}},"authHost":"https://dcos.auth0.com"},



    "tracking":{"enabled":false}}}}

    '
  path: /etc/mesosphere/setup-packages/dcos-provider-azure--setup/etc/ui-config.json
- content: '{}'
  path: /etc/mesosphere/setup-packages/dcos-provider-azure--setup/pkginfo.json
- content: '[Unit]

    Before=dcos.target

    [Service]

    Type=oneshot

    StandardOutput=journal+console

    StandardError=journal+console

    ExecStartPre=/usr/bin/mkdir -p /etc/profile.d

    ExecStart=/usr/bin/ln -sf /opt/mesosphere/environment.export /etc/profile.d/dcos.sh

    '
  path: /etc/systemd/system/dcos-link-env.service
  permissions: '0644'
- content: '[Unit]

    Description=Pkgpanda: Download DC/OS to this host.

    After=network-online.target

    Wants=network-online.target

    ConditionPathExists=!/opt/mesosphere/

    [Service]

    Type=oneshot

    StandardOutput=journal+console

    StandardError=journal+console

    ExecStartPre=/usr/bin/curl --keepalive-time 2 -fLsSv --retry 20 -Y 100000 -y 60 -o /var/lib/mesos/dl/bootstrap.tar.xz {{{dcosBootstrapURL}}}

    ExecStartPre=/usr/bin/mkdir -p /opt/mesosphere

    ExecStart=/usr/bin/tar -axf /var/lib/mesos/dl/bootstrap.tar.xz -C /opt/mesosphere

    ExecStartPost=-/usr/bin/rm -f /var/lib/mesos/dl/bootstrap.tar.xz

    '
  path: /etc/systemd/system/dcos-download.service
  permissions: '0644'
- content: '[Unit]

    Description=Pkgpanda: Specialize DC/OS for this host.

    Requires=dcos-download.service

    After=dcos-download.service

    [Service]

    Type=oneshot

    StandardOutput=journal+console

    StandardError=journal+console

    EnvironmentFile=/opt/mesosphere/environment

    ExecStart=/opt/mesosphere/bin/pkgpanda setup --no-block-systemd

    [Install]

    WantedBy=multi-user.target

    '
  path: /etc/systemd/system/dcos-setup.service
  permissions: '0644'
- content: ''
  path: /etc/mesosphere/roles/azure
- content: 'PROVISION_STR'
  path: "/opt/azure/containers/provision.sh"
  permissions: "0744"
  owner: "root"
- path: /var/lib/dcos/mesos-slave-common
  content: 'ATTRIBUTES_STR'
  permissions: "0644"
  owner: "root"
- content: '{ "auths": { "{{{registry}}}": { "auth" : "{{{registryKey}}}" } } }'
  path: "/tmp/xtoph/.docker/config.json"
  owner: "root"
`)

func dcosDcoscustomdata188TBytes() ([]byte, error) {
	return _dcosDcoscustomdata188T, nil
}

func dcosDcoscustomdata188T() (*asset, error) {
	bytes, err := dcosDcoscustomdata188TBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/dcoscustomdata188.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosDcoscustomdata190T = []byte(`bootcmd:
- bash -c "if [ ! -f /var/lib/sdb-gpt ];then echo DCOS-5890;parted -s /dev/sdb mklabel
  gpt;touch /var/lib/sdb-gpt;fi"
disk_setup:
  ephemeral0:
    layout:
    - 45
    - 45
    - 10
    overwrite: true
    table_type: gpt
fs_setup:
- device: ephemeral0.1
  filesystem: ext4
  overwrite: true
- device: ephemeral0.2
  filesystem: ext4
  overwrite: true
- device: ephemeral0.3
  filesystem: ext4
  overwrite: true
mounts:
- - ephemeral0.1
  - /var/lib/mesos
- - ephemeral0.2
  - /var/lib/docker
- - ephemeral0.3
  - /var/tmp
runcmd: PREPROVISION_EXTENSION
- /usr/lib/apt/apt.systemd.daily
- echo 2dd1ce17-079e-403c-b352-a1921ee207ee > /sys/bus/vmbus/drivers/hv_util/unbind # mitigation for bug https://bugs.launchpad.net/ubuntu/+source/linux/+bug/1676635
- sed -i "13i\echo 2dd1ce17-079e-403c-b352-a1921ee207ee > /sys/bus/vmbus/drivers/hv_util/unbind\n" /etc/rc.local # mitigation for bug https://bugs.launchpad.net/ubuntu/+source/linux/+bug/1676635
- - ln
  - -s
  - /bin/rm
  - /usr/bin/rm
- - ln
  - -s
  - /bin/mkdir
  - /usr/bin/mkdir
- - ln
  - -s
  - /bin/tar
  - /usr/bin/tar
- - ln
  - -s
  - /bin/ln
  - /usr/bin/ln
- - ln
  - -s
  - /bin/cp
  - /usr/bin/cp
- - ln
  - -s
  - /bin/systemctl
  - /usr/bin/systemctl
- - ln
  - -s
  - /bin/mount
  - /usr/bin/mount
- - ln
  - -s
  - /bin/bash
  - /usr/bin/bash
- - ln
  - -s
  - /usr/sbin/useradd
  - /usr/bin/useradd
- - systemctl
  - disable
  - --now
  - resolvconf.service
- - systemctl
  - mask
  - --now
  - lxc-net.service
- - systemctl
  - disable
  - --now
  - unscd.service
- - systemctl
  - stop
  - --now
  - unscd.service
- sed -i "s/^Port 22$/Port 22\nPort 2222/1" /etc/ssh/sshd_config
- service ssh restart 
- /opt/azure/containers/setup_ephemeral_disk.sh
- - tar
  - czf 
  - /etc/docker.tar.gz
  - -C
  - /tmp/xtoph
  - .docker
- - rm 
  - -rf 
  - /tmp/xtoph
- /opt/azure/containers/provision.sh
- - cp
  - -p
  - /etc/resolv.conf
  - /tmp/resolv.conf
- - rm
  - -f
  - /etc/resolv.conf
- - cp
  - -p
  - /tmp/resolv.conf
  - /etc/resolv.conf
- - systemctl
  - start
  - dcos-docker-install.service
- - systemctl
  - start
  - dcos-config-writer.service
- - systemctl
  - restart
  - systemd-journald.service
- - systemctl
  - restart
  - docker.service
- - systemctl
  - start
  - dcos-link-env.service
- - systemctl
  - enable
  - dcos-setup.service
- - systemctl
  - --no-block
  - start
  - dcos-setup.service
write_files:
- content: '{{{dcosRepositoryURL}}}

'
  owner: root
  path: /etc/mesosphere/setup-flags/repository-url
  permissions: '0644'
- content: '["3dt--7847ebb24bf6756c3103902971b34c3f09c3afbd", "adminrouter--0493a6fdaed08e1971871818e194aa4607df4f09",
    "avro-cpp--760c214063f6b038b522eaf4b768b905fed56ebc", "boost-libs--2015ccb58fb756f61c02ee6aa05cc1e27459a9ec",
    "bootstrap--59a905ecee27e71168ed44cefda4481fb76b816d", "boto--6344d31eef082c7bd13259b17034ea7b5c34aedf",
    "check-time--be7d0ba757ec87f9965378fee7c76a6ee5ae996d", "cni--e48337da39a8cd379414acfe0da52a9226a10d24",
    "cosmos--20decef90f0623ed253a12ec4cf5c148b18d8249", "curl--fc3486c43f98e63f9b12675f1356e8fe842f26b0",
    "dcos-config--setup_{{{dcosProviderPackageID}}}", "dcos-history--77b0e97d7b25c8bedf8f7da0689cac65b83e3813",
    "dcos-image--bda6a02bcb2eb21c4218453a870cc584f921a800", "dcos-image-deps--83584fd868e5b470f7cf754424a9a75b328e9b68",
    "dcos-integration-test--c28bcb2347799dca43083f55e4c7b28503176f9c", "dcos-log--4d630df863228f38c6333e44670b4c4b20a74832",
    "dcos-metadata--setup_{{{dcosProviderPackageID}}}", "dcos-metrics--23ee2f89c58b1258bc959f1d0dd7debcbb3d79d2",
    "dcos-oauth--0079529da183c0f23a06d2b069721b6fa6cc7b52", "dcos-signal--1bcd3b612cbdc379380dcba17cdf9a3b6652d9dc",
    "dcos-ui--d4afd695796404a5b35950c3daddcae322481ac4", "dnspython--0f833eb9a8abeba3179b43f3a200a8cd42d3795a",
    "docker-gc--59a98ed6446a084bf74e4ff4b8e3479f59ea8528", "dvdcli--5374dd4ffb519f1dcefdec89b2247e3404f2e2e3",
    "erlang--a9ee2530357a3301e53056b36a93420847b339a3", "exhibitor--72d9d8f947e5411eda524d40dde1a58edeb158ed",
    "flask--26d1bcdb2d1c3dcf1d2c03bc0d4f29c86d321b21", "java--cd5e921ce66b0d3303883c06d73a657314044304",
    "libevent--208be855d2be29c9271a7bd6c04723ff79946e02", "libffi--83ce3bd7eda2ef089e57efd2bc16c144d5a1f094",
    "libsodium--9ff915db08c6bba7d6738af5084e782b13c84bf8", "logrotate--7f7bc4416d3ad101d0c5218872858483b516be07",
    "marathon--bfb24f7f90cb3cd52a1cb22a07caafa5013bba21", "mesos--aaedd03eee0d57f5c0d49c74ff1e5721862cad98",
    "mesos-dns--0401501b2b5152d01bfa84ff6d007fdafe414b16", "mesos-modules--311849eaae42696b8a7eefe86b9ab3ebd9bd48f5",
    "metronome--467e4c64f804dbd4cd8572516e111a3f9298c10d", "navstar--1128db0234105a64fb4be52f4453cd6aa895ff30",
    "ncurses--d889894b71aa1a5b311bafef0e85479025b4dacb", "octarine--e86d3312691b12523280d56f6260216729aaa0ad",
    "openssl--b01a32a42e3ccba52b417276e9509a441e1d4a82", "pkgpanda-api--541feb8a8be58bdde8fecf1d2e5bfa0515f5a7d0",
    "pkgpanda-role--f8a749a4a821476ad2ef7e9dd9d12b6a8c4643a4", "pytest--78aee3e58a049cdab0d266af74f77d658b360b4f",
    "python--b7a144a49577a223d37d447c568f51330ee95390", "python-azure-mgmt-resource--03c05550f43b0e7a4455c33fe43b0deb755d87f0",
    "python-cryptography--4184767c68e48801dd394072cb370c610a05029d", "python-dateutil--fdc6ff929f65dd0918cf75a9ad56704683d31781",
    "python-docopt--beba78faa13e5bf4c52393b4b82d81f3c391aa65", "python-gunicorn--a537f95661fb2689c52fe12510eb0d01cb83af60",
    "python-isodate--40d378c688e6badfd16676dd8b51b742bfebc8d5", "python-jinja2--7450f5ae5a822f63f7a58c717207be0456df51ed",
    "python-kazoo--cb7ce13a1068cd82dd84ea0de32b529a760a4bdd", "python-markupsafe--dd46d2a3c58611656a235f96d4adc51b2a7a590e",
    "python-passlib--802ec3605c0b82428fedba60983b1bafaa036bb8", "python-pyyaml--81dd44cc4a24db7cefa7016c6586a131acf279c3",
    "python-requests--1b2cadbd3811cc0c2ee235ce927e13ea1d6af41d", "python-retrying--eb7b8bac133f50492b1e1349cbe77c3e38bd02c3",
    "python-tox--07244f8a939a10353634c952c6d88ec4a3c05736", "rexray--869621bb411c9f2a793ea42cdfeed489e1972aaa",
    "six--f06424b68523c4dfa2a7c3e7475d479f3d361e42", "spartan--9cc57a3d55452b905d90e3201f56913140914ecc",
    "strace--7d01796d64994451c1b2b82d161a335cbe90569b", "teamcity-messages--e623a4d86eb3a8d199cefcc240dd4c5460cb2962",
    "toybox--f235594ab8ea9a2864ee72abe86723d76f92e848"]

'
  owner: root
  path: /etc/mesosphere/setup-flags/cluster-packages.json
  permissions: '0644'
- content: |
    [Journal]
    MaxLevelConsole=warning
    RateLimitInterval=1s
    RateLimitBurst=20000
  owner: root
  path: /etc/systemd/journald.conf.d/dcos.conf
  permissions: '0644'
- content: |
    rexray:
      loglevel: info
      modules:
        default-admin:
          host: tcp://127.0.0.1:61003
        default-docker:
          disabled: true
  path: /etc/rexray/config.yml
  permissions: '0644'
- content: |
    [Unit]
    After=network-online.target
    Wants=network-online.target
    [Service]
    Type=oneshot
    Environment=DEBIAN_FRONTEND=noninteractive
    StandardOutput=journal+console
    StandardError=journal+console
    ExecStartPre=/usr/bin/curl -fLsSv --retry 20 -Y 100000 -y 60 -o /var/tmp/d.deb https://az837203.vo.msecnd.net/dcos-deps/docker-engine_1.13.1-0-ubuntu-xenial_amd64.deb
    ExecStart=/usr/bin/bash -c "try=1;until dpkg -D3 -i /var/tmp/d.deb || ((try>9));do echo retry $((try++));sleep $((try*try));done;systemctl --now start docker;systemctl restart docker.socket"
  path: /etc/systemd/system/dcos-docker-install.service
  permissions: '0644'
- content: |
    [Service]
    Restart=always
    StartLimitInterval=0
    RestartSec=15
    LimitNOFILE=16384
    ExecStartPre=-/sbin/ip link del docker0
    ExecStart=
    ExecStart=/usr/bin/docker daemon -H fd:// --storage-driver=overlay
  path: /etc/systemd/system/docker.service.d/execstart.conf
  permissions: '0644'
- content: |
    [Unit]
    PartOf=docker.service
    [Socket]
    ListenStream=/var/run/docker.sock
    SocketMode=0660
    SocketUser=root
    SocketGroup=docker
    ListenStream=2375
    BindIPv6Only=both
    [Install]
    WantedBy=sockets.target
  path: /etc/systemd/system/docker.socket
  permissions: '0644'
  content: |
      [Unit]
      Requires=dcos-setup.service
      After=dcos-setup.service
      [Service]
      Type=oneshot
      EnvironmentFile=/etc/environment
      EnvironmentFile=/opt/mesosphere/environment
      ExecStart=/usr/bin/bash -c "echo $(detect_ip) $(hostname) > /etc/hosts"
  path: /etc/systemd/system/dcos-config-writer.service
  permissions: '0644'
- content: |
    "bound_values":
      "adminrouter_auth_enabled": |-
        {{{oauthEnabled}}}
      "cluster_name": |-
        {{{masterPublicIPAddressName}}}
      "exhibitor_azure_account_key": |-
        ', listKeys(resourceId('Microsoft.Storage/storageAccounts', variables('masterStorageAccountExhibitorName')), '2015-06-15').key1, '
      "exhibitor_azure_account_name": |-
        {{{masterStorageAccountExhibitorName}}}
      "exhibitor_azure_prefix": |-
        {{{masterPublicIPAddressName}}}
      "master_list": |-
        ["', DCOSCUSTOMDATAPUBLICIPSTR'"]
      "oauth_enabled": |-
        {{{oauthEnabled}}}
    "late_bound_package_id": |-
      dcos-provider-{{{dcosProviderPackageID}}}-azure--setup
  owner: root
  path: /etc/mesosphere/setup-flags/late-config.yaml
  permissions: '0644'
- content: |
    [Unit]
    Before=dcos.target
    [Service]
    Type=oneshot
    StandardOutput=journal+console
    StandardError=journal+console
    ExecStartPre=/usr/bin/mkdir -p /etc/profile.d
    ExecStart=/usr/bin/ln -sf /opt/mesosphere/bin/add_dcos_path.sh /etc/profile.d/dcos.sh
  path: /etc/systemd/system/dcos-link-env.service
  permissions: '0644'
- content: |
    [Unit]
    Description=Pkgpanda: Download DC/OS to this host.
    After=network-online.target
    Wants=network-online.target
    ConditionPathExists=!/opt/mesosphere/
    [Service]
    Type=oneshot
    StandardOutput=journal+console
    StandardError=journal+console
    ExecStartPre=/usr/bin/curl --keepalive-time 2 -fLsSv --retry 20 -Y 100000 -y 60 -o //var/tmp/bootstrap.tar.xz {{{dcosBootstrapURL}}}
    ExecStartPre=/usr/bin/mkdir -p /opt/mesosphere
    ExecStart=/usr/bin/tar -axf //var/tmp/bootstrap.tar.xz -C /opt/mesosphere
    ExecStartPost=-/usr/bin/rm -f //var/tmp/bootstrap.tar.xz
  path: /etc/systemd/system/dcos-download.service
  permissions: '0644'
- content: |
    [Unit]
    Description=Pkgpanda: Specialize DC/OS for this host.
    Requires=dcos-download.service
    After=dcos-download.service
    [Service]
    Type=oneshot
    StandardOutput=journal+console
    StandardError=journal+console
    EnvironmentFile=/opt/mesosphere/environment
    ExecStart=/opt/mesosphere/bin/pkgpanda setup --no-block-systemd
    [Install]
    WantedBy=multi-user.target
  path: /etc/systemd/system/dcos-setup.service
  permissions: '0644'
- content: ''
  path: /etc/mesosphere/roles/azure
- content: 'PROVISION_STR'
  path: "/opt/azure/containers/provision.sh"
  permissions: "0744"
  owner: "root"
- path: /var/lib/dcos/mesos-slave-common
  content: 'ATTRIBUTES_STR'
  permissions: "0644"
  owner: "root"
- content: '{ "auths": { "{{{registry}}}": { "auth" : "{{{registryKey}}}" } } }'
  path: "/tmp/xtoph/.docker/config.json"
  owner: "root"
- content: |
    #!/bin/bash
    # Check the partitions on /dev/sdb created by cloudinit and force a detach and
    # reformat of the parition.  After which, all will be remounted.
    EPHEMERAL_DISK="/dev/sdb"
    PARTITIONS=` + "`" + `fdisk -l $EPHEMERAL_DISK | grep "^$EPHEMERAL_DISK" | cut -d" " -f1 | sed "s~$EPHEMERAL_DISK~~"` + "`" + `
    if [ -n "$PARTITIONS" ]; then
        for f in $PARTITIONS; do
            df -k | grep "/dev/sdb$f"
            if [ $? -eq 0 ]; then
                umount -f /dev/sdb$f
            fi
            mkfs.ext4 /dev/sdb$f
        done
        mount -a
    fi
    # If there is a /var/tmp partition on the ephemeral disk, create a symlink such
    # that the /var/log/mesos and /var/log/journal placed on the ephemeral disk.
    VAR_TMP_PARTITION=` + "`" + `df -P /var/tmp | tail -1 | cut -d" " -f 1` + "`" + `
    echo $VAR_TMP_PARTITION | grep "^$EPHEMERAL_DISK"
    if [ $? -eq 0 ]; then
        # Handle the /var/log/mesos directory
        mkdir -p /var/tmp/log/mesos
        if [ -d "/var/log/mesos" ]; then
            cp -rp /var/log/mesos/* /var/tmp/log/mesos/
            rm -rf /var/log/mesos
        fi
        ln -s /var/tmp/log/mesos /var/log/mesos
        # Handle the /var/log/journal direcotry
        mkdir -p /var/tmp/log/journal
        if [ -d "/var/log/journal" ]; then
            cp -rp /var/log/journal/* /var/tmp/log/journal/
            rm -rf /var/log/journal
        fi
        ln -s /var/tmp/log/journal /var/log/journal
    fi
  path: "/opt/azure/containers/setup_ephemeral_disk.sh"
  permissions: "0744"
  owner: "root"`)

func dcosDcoscustomdata190TBytes() ([]byte, error) {
	return _dcosDcoscustomdata190T, nil
}

func dcosDcoscustomdata190T() (*asset, error) {
	bytes, err := dcosDcoscustomdata190TBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/dcoscustomdata190.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosDcoscustomdata198T = []byte(`bootcmd:
- bash -c "if [ ! -f /var/lib/sdb-gpt ];then echo DCOS-5890;parted -s /dev/sdb mklabel
  gpt;touch /var/lib/sdb-gpt;fi"
disk_setup:
  ephemeral0:
    layout:
    - 45
    - 45
    - 10
    overwrite: true
    table_type: gpt
fs_setup:
- device: ephemeral0.1
  filesystem: ext4
  overwrite: true
- device: ephemeral0.2
  filesystem: ext4
  overwrite: true
- device: ephemeral0.3
  filesystem: ext4
  overwrite: true
mounts:
- - ephemeral0.1
  - /var/lib/mesos
- - ephemeral0.2
  - /var/lib/docker
- - ephemeral0.3
  - /var/tmp
runcmd: PREPROVISION_EXTENSION
- /usr/lib/apt/apt.systemd.daily
- echo 2dd1ce17-079e-403c-b352-a1921ee207ee > /sys/bus/vmbus/drivers/hv_util/unbind # mitigation for bug https://bugs.launchpad.net/ubuntu/+source/linux/+bug/1676635
- sed -i "13i\echo 2dd1ce17-079e-403c-b352-a1921ee207ee > /sys/bus/vmbus/drivers/hv_util/unbind\n" /etc/rc.local # mitigation for bug https://bugs.launchpad.net/ubuntu/+source/linux/+bug/1676635
- - ln
  - -s
  - /bin/rm
  - /usr/bin/rm
- - ln
  - -s
  - /bin/mkdir
  - /usr/bin/mkdir
- - ln
  - -s
  - /bin/tar
  - /usr/bin/tar
- - ln
  - -s
  - /bin/ln
  - /usr/bin/ln
- - ln
  - -s
  - /bin/cp
  - /usr/bin/cp
- - ln
  - -s
  - /bin/systemctl
  - /usr/bin/systemctl
- - ln
  - -s
  - /bin/mount
  - /usr/bin/mount
- - ln
  - -s
  - /bin/bash
  - /usr/bin/bash
- - ln
  - -s
  - /usr/sbin/useradd
  - /usr/bin/useradd
- - systemctl
  - disable
  - --now
  - resolvconf.service
- - systemctl
  - mask
  - --now
  - lxc-net.service
- - systemctl
  - disable
  - --now
  - unscd.service
- - systemctl
  - stop
  - --now
  - unscd.service
- sed -i "s/^Port 22$/Port 22\nPort 2222/1" /etc/ssh/sshd_config
- service ssh restart 
- /opt/azure/containers/setup_ephemeral_disk.sh
- - tar
  - czf 
  - /etc/docker.tar.gz
  - -C
  - /tmp/xtoph
  - .docker
- - rm 
  - -rf 
  - /tmp/xtoph
- /opt/azure/containers/provision.sh
- - cp
  - -p
  - /etc/resolv.conf
  - /tmp/resolv.conf
- - rm
  - -f
  - /etc/resolv.conf
- - cp
  - -p
  - /tmp/resolv.conf
  - /etc/resolv.conf
- - systemctl
  - start
  - dcos-docker-install.service
- - systemctl
  - start
  - dcos-config-writer.service
- - systemctl
  - restart
  - systemd-journald.service
- - systemctl
  - restart
  - docker.service
- - systemctl
  - start
  - dcos-link-env.service
- - systemctl
  - enable
  - dcos-setup.service
- - systemctl
  - --no-block
  - start
  - dcos-setup.service
write_files:
- content: '{{{dcosRepositoryURL}}}

'
  owner: root
  path: /etc/mesosphere/setup-flags/repository-url
  permissions: '0644'
- content: '["3dt--4eb6a10d16421bc87cb6e93ac97746f36aded925", "adminrouter--31f3f6390c8ef79a2774f42390d6340a24d67f08",
    "avro-cpp--6194e9a67928c357c1c1b2bb409536ceef888e04", "boost-libs--2015ccb58fb756f61c02ee6aa05cc1e27459a9ec",
    "bootstrap--d50592de9bf45937df7bcc7008e84a8739239c99", "boto--471853efd730e52e4ed7bfb890587432a576982a",
    "check-time--be7d0ba757ec87f9965378fee7c76a6ee5ae996d", "cni--e48337da39a8cd379414acfe0da52a9226a10d24",
    "cosmos--74e0339c91c278622d9f45b5fb0771872f443140", "curl--e7fd5880e4f94db05692d7e43279d8fe6348cb21",
    "dcos-config--setup_{{{dcosProviderPackageID}}}", "dcos-history--787ce2fd81cb7469590c12951033f0482e879d2a",
    "dcos-image--078703170a2f218447abea4b1be00b7431b340f1", "dcos-image-deps--5512ff49cdbba7f404759a5751a4ab1eae44c677",
    "dcos-integration-test--bad12974ed31ace44432ad9a451c5b5dc3e20e81", "dcos-log--4d630df863228f38c6333e44670b4c4b20a74832",
    "dcos-metadata--setup_{{{dcosProviderPackageID}}}", "dcos-metrics--e65d65e1b65335efdaa6bf7609a671f4288e7af9",
    "dcos-oauth--23d8ca77549c1ac6087c11c9f7e8f8a4fddfc948", "dcos-signal--5633dc8da7e864cb34e3d29ed13e6756c7a6df94",
    "dcos-ui--6f4af319cf4dd9bb8366de22ec37775beaa96747", "dnspython--1118f0ffaa60e6a779d4614f0ed692d215005f0e",
    "docker-gc--9737ec72de5d1edc71175028762f06fe22c8a48c", "dvdcli--5374dd4ffb519f1dcefdec89b2247e3404f2e2e3",
    "erlang--984871e11f69e37aeb76a471d4a4b90e93fdf355", "exhibitor--300da0c612afcf27541dbc681da5de3a6408de7e",
    "flask--2936647fa917d16ee289d34e61fd1afcc49157b5", "java--091eb5a0f3dcbd7762a43e84c3e2d6aac8891111",
    "libevent--468f4ae789f659e452e8356a9d2309c4f41135a8", "libffi--83ce3bd7eda2ef089e57efd2bc16c144d5a1f094",
    "libsodium--9ff915db08c6bba7d6738af5084e782b13c84bf8", "logrotate--7f7bc4416d3ad101d0c5218872858483b516be07",
    "marathon--99d0cbc65da6be31872878174f3a28fa63d0fa34", "mesos--0c992033b8d43e00dc69f0c548c826d573c82642",
    "mesos-dns--ca591a18f9b010999106285fedddd010606c0d06", "mesos-modules--4c176c23a4fd3670d059fec55e2d4c8c7dbf1f6c",
    "metronome--138ec50cd4da05bce74b6cd2c84ae873c2bd67ab", "navstar--fdf7e79fdf210548d183badfde00d60c1a540257",
    "ncurses--d889894b71aa1a5b311bafef0e85479025b4dacb", "octarine--4e37c062d2f145f9c2ce01d30dadf72c2aac5c4a",
    "openssl--ef04a6f76f6e5e194c783bc129fdabad16816aff", "pkgpanda-api--220e45fbd93403f8b4fd7f9c8c3d5178aff6e34b",
    "pkgpanda-role--f8a749a4a821476ad2ef7e9dd9d12b6a8c4643a4", "pytest--63ab7e9520e4da70202b81076880fcdf2c1236cf",
    "python--3c96ab7f21312f4d7d54a9b901cfe6382aa66b8a", "python-azure-mgmt-resource--2313114eec2adcb37ef61082cd2cfdceabf5c21e",
    "python-cryptography--39ee7d59411569700f3343e64c32e9711a83decc", "python-dateutil--d098c1933ca6d754a90734afd366d556cc3107a8",
    "python-docopt--85e7726dbb777584a9f5d4dd7bd58ed8ca5466d8", "python-gunicorn--bd425f55abd9236b5ead7e68a3c40c39b8d75bb7",
    "python-isodate--9a15007db453e141892966ebf50a9175ee0ba08b", "python-jinja2--9fbc35d1405f06f1959c54629ab7d443cef79076",
    "python-kazoo--050358610274815ebacabcdfca874729e53f4e0b", "python-markupsafe--09c65e6cdedd4783137a203cbc1b5a64ef3124eb",
    "python-passlib--27056b95ad1a067b7992402e679c6260e673a554", "python-pyyaml--5be319fd73348558d69a03fb6dcb134e9b7f4c48",
    "python-requests--63e1c3f4f03efc4607a4c20c5492026a9af7a9c7", "python-retrying--692b1a298d22436e25b2d14fc4f980be444adbe7",
    "python-tox--7962137d89dae9eb45dd80b0ea59731fa3f5bbc9", "rexray--f07795e2c10f9a1a27de9d8e67ab171029db2e1d",
    "six--9229b1a9d7d57bc086fa50f73fc9a753d9a4605d", "spartan--3dc1785bf698e65ceb2fecf26b2a439de219269f",
    "strace--7d01796d64994451c1b2b82d161a335cbe90569b", "teamcity-messages--d13bc3f52ed0e30de3a71d86ff8718984b60b65f",
    "toybox--c0e85790eb8aaeefe5037b053c2fcd140ab800a4"]

'
  owner: root
  path: /etc/mesosphere/setup-flags/cluster-packages.json
  permissions: '0644'
- content: |
    [Journal]
    MaxLevelConsole=warning
    RateLimitInterval=1s
    RateLimitBurst=20000
  owner: root
  path: /etc/systemd/journald.conf.d/dcos.conf
  permissions: '0644'
- content: |
    rexray:
      loglevel: info
      modules:
        default-admin:
          host: tcp://127.0.0.1:61003
        default-docker:
          disabled: true
  path: /etc/rexray/config.yml
  permissions: '0644'
- content: |
    [Unit]
    After=network-online.target
    Wants=network-online.target
    [Service]
    Type=oneshot
    Environment=DEBIAN_FRONTEND=noninteractive
    StandardOutput=journal+console
    StandardError=journal+console
    ExecStartPre=/usr/bin/curl -fLsSv --retry 20 -Y 100000 -y 60 -o /var/tmp/d.deb https://az837203.vo.msecnd.net/dcos-deps/docker-engine_1.13.1-0-ubuntu-xenial_amd64.deb
    ExecStart=/usr/bin/bash -c "try=1;until dpkg -D3 -i /var/tmp/d.deb || ((try>9));do echo retry $((try++));sleep $((try*try));done;systemctl --now start docker;systemctl restart docker.socket"
  path: /etc/systemd/system/dcos-docker-install.service
  permissions: '0644'
- content: |
    [Service]
    Restart=always
    StartLimitInterval=0
    RestartSec=15
    LimitNOFILE=16384
    ExecStartPre=-/sbin/ip link del docker0
    ExecStart=
    ExecStart=/usr/bin/docker daemon -H fd:// --storage-driver=overlay
  path: /etc/systemd/system/docker.service.d/execstart.conf
  permissions: '0644'
- content: |
    [Unit]
    PartOf=docker.service
    [Socket]
    ListenStream=/var/run/docker.sock
    SocketMode=0660
    SocketUser=root
    SocketGroup=docker
    ListenStream=2375
    BindIPv6Only=both
    [Install]
    WantedBy=sockets.target
  path: /etc/systemd/system/docker.socket
  permissions: '0644'
  content: |
      [Unit]
      Requires=dcos-setup.service
      After=dcos-setup.service
      [Service]
      Type=oneshot
      EnvironmentFile=/etc/environment
      EnvironmentFile=/opt/mesosphere/environment
      ExecStart=/usr/bin/bash -c "echo $(detect_ip) $(hostname) > /etc/hosts"
  path: /etc/systemd/system/dcos-config-writer.service
  permissions: '0644'
- content: |
    "bound_values":
      "adminrouter_auth_enabled": |-
        {{{oauthEnabled}}}
      "cluster_name": |-
        {{{masterPublicIPAddressName}}}
      "exhibitor_azure_account_key": |-
        ', listKeys(resourceId('Microsoft.Storage/storageAccounts', variables('masterStorageAccountExhibitorName')), '2015-06-15').key1, '
      "exhibitor_azure_account_name": |-
        {{{masterStorageAccountExhibitorName}}}
      "exhibitor_azure_prefix": |-
        {{{masterPublicIPAddressName}}}
      "master_list": |-
        ["', DCOSCUSTOMDATAPUBLICIPSTR'"]
      "oauth_enabled": |-
        {{{oauthEnabled}}}
    "late_bound_package_id": |-
      dcos-provider-{{{dcosProviderPackageID}}}-azure--setup
  owner: root
  path: /etc/mesosphere/setup-flags/late-config.yaml
  permissions: '0644'
- content: |
    [Unit]
    Before=dcos.target
    [Service]
    Type=oneshot
    StandardOutput=journal+console
    StandardError=journal+console
    ExecStartPre=/usr/bin/mkdir -p /etc/profile.d
    ExecStart=/usr/bin/ln -sf /opt/mesosphere/bin/add_dcos_path.sh /etc/profile.d/dcos.sh
  path: /etc/systemd/system/dcos-link-env.service
  permissions: '0644'
- content: |
    [Unit]
    Description=Pkgpanda: Download DC/OS to this host.
    After=network-online.target
    Wants=network-online.target
    ConditionPathExists=!/opt/mesosphere/
    [Service]
    Type=oneshot
    StandardOutput=journal+console
    StandardError=journal+console
    ExecStartPre=/usr/bin/curl --keepalive-time 2 -fLsSv --retry 20 -Y 100000 -y 60 -o //var/tmp/bootstrap.tar.xz {{{dcosBootstrapURL}}}
    ExecStartPre=/usr/bin/mkdir -p /opt/mesosphere
    ExecStart=/usr/bin/tar -axf //var/tmp/bootstrap.tar.xz -C /opt/mesosphere
    ExecStartPost=-/usr/bin/rm -f //var/tmp/bootstrap.tar.xz
  path: /etc/systemd/system/dcos-download.service
  permissions: '0644'
- content: |
    [Unit]
    Description=Pkgpanda: Specialize DC/OS for this host.
    Requires=dcos-download.service
    After=dcos-download.service
    [Service]
    Type=oneshot
    StandardOutput=journal+console
    StandardError=journal+console
    EnvironmentFile=/opt/mesosphere/environment
    ExecStart=/opt/mesosphere/bin/pkgpanda setup --no-block-systemd
    [Install]
    WantedBy=multi-user.target
  path: /etc/systemd/system/dcos-setup.service
  permissions: '0644'
- content: ''
  path: /etc/mesosphere/roles/azure
- content: 'PROVISION_STR'
  path: "/opt/azure/containers/provision.sh"
  permissions: "0744"
  owner: "root"
- path: /var/lib/dcos/mesos-slave-common
  content: 'ATTRIBUTES_STR'
  permissions: "0644"
  owner: "root"
- content: '{ "auths": { "{{{registry}}}": { "auth" : "{{{registryKey}}}" } } }'
  path: "/tmp/xtoph/.docker/config.json"
  owner: "root"
- content: |
    #!/bin/bash
    # Check the partitions on /dev/sdb created by cloudinit and force a detach and
    # reformat of the parition.  After which, all will be remounted.
    EPHEMERAL_DISK="/dev/sdb"
    PARTITIONS=` + "`" + `fdisk -l $EPHEMERAL_DISK | grep "^$EPHEMERAL_DISK" | cut -d" " -f1 | sed "s~$EPHEMERAL_DISK~~"` + "`" + `
    if [ -n "$PARTITIONS" ]; then
        for f in $PARTITIONS; do
            df -k | grep "/dev/sdb$f"
            if [ $? -eq 0 ]; then
                umount -f /dev/sdb$f
            fi
            mkfs.ext4 /dev/sdb$f
        done
        mount -a
    fi
    # If there is a /var/tmp partition on the ephemeral disk, create a symlink such
    # that the /var/log/mesos and /var/log/journal placed on the ephemeral disk.
    VAR_TMP_PARTITION=` + "`" + `df -P /var/tmp | tail -1 | cut -d" " -f 1` + "`" + `
    echo $VAR_TMP_PARTITION | grep "^$EPHEMERAL_DISK"
    if [ $? -eq 0 ]; then
        # Handle the /var/log/mesos directory
        mkdir -p /var/tmp/log/mesos
        if [ -d "/var/log/mesos" ]; then
            cp -rp /var/log/mesos/* /var/tmp/log/mesos/
            rm -rf /var/log/mesos
        fi
        ln -s /var/tmp/log/mesos /var/log/mesos
        # Handle the /var/log/journal direcotry
        mkdir -p /var/tmp/log/journal
        if [ -d "/var/log/journal" ]; then
            cp -rp /var/log/journal/* /var/tmp/log/journal/
            rm -rf /var/log/journal
        fi
        ln -s /var/tmp/log/journal /var/log/journal
    fi
  path: "/opt/azure/containers/setup_ephemeral_disk.sh"
  permissions: "0744"
  owner: "root"
`)

func dcosDcoscustomdata198TBytes() ([]byte, error) {
	return _dcosDcoscustomdata198T, nil
}

func dcosDcoscustomdata198T() (*asset, error) {
	bytes, err := dcosDcoscustomdata198TBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/dcoscustomdata198.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosDcosmasterresourcesT = []byte(`{{if .MasterProfile.IsManagedDisks}}
    {
      "apiVersion": "[variables('apiVersionStorageManagedDisks')]",
      "location": "[variables('location')]",
      "name": "[variables('masterAvailabilitySet')]",
      "properties": {
        "platformFaultDomainCount": 2,
        "platformUpdateDomainCount": 3,
        "managed": "true"
      },
      "type": "Microsoft.Compute/availabilitySets"
    },
{{else if .MasterProfile.IsStorageAccount}}
    {
      "apiVersion": "[variables('apiVersionStorage')]",
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[variables('masterStorageAccountName')]",
      "properties": {
        "accountType": "[variables('vmSizesMap')[variables('masterVMSize')].storageAccountType]"
      },
      "type": "Microsoft.Storage/storageAccounts"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('masterAvailabilitySet')]",
      "properties": {},
      "type": "Microsoft.Compute/availabilitySets"
    },
{{end}}
    {
      "apiVersion": "[variables('apiVersionStorage')]",
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[variables('masterStorageAccountExhibitorName')]",
      "properties": {
        "accountType": "Standard_LRS"
      },
      "type": "Microsoft.Storage/storageAccounts"
    },
{{if not .MasterProfile.IsCustomVNET}}
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "dependsOn": [
          {{GetVNETSubnetDependencies}}
      ],
      "location": "[variables('location')]",
      "name": "[variables('virtualNetworkName')]",
      "properties": {
        "addressSpace": {
          "addressPrefixes": [
            {{GetVNETAddressPrefixes}}
          ]
        },
        "subnets": [
          {{GetVNETSubnets true}}
        ]
      },
      "type": "Microsoft.Network/virtualNetworks"
    },
{{end}}
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('masterPublicIPAddressName')]",
      "properties": {
        "dnsSettings": {
          "domainNameLabel": "[variables('masterEndpointDNSNamePrefix')]"
        },
        "publicIPAllocationMethod": "Dynamic"
      },
      "type": "Microsoft.Network/publicIPAddresses"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[variables('masterLbName')]",
      "properties": {
        "backendAddressPools": [
          {
            "name": "[variables('masterLbBackendPoolName')]"
          }
        ],
        "frontendIPConfigurations": [
          {
            "name": "[variables('masterLbIPConfigName')]",
            "properties": {
              "publicIPAddress": {
                "id": "[resourceId('Microsoft.Network/publicIPAddresses',variables('masterPublicIPAddressName'))]"
              }
            }
          }
        ]
{{if .MasterProfile.OAuthEnabled}}
        ,"loadBalancingRules": [
	        {
            "name": "LBRule443",
            "properties": {
              "frontendIPConfiguration": {
                "id": "[variables('masterLbIPConfigID')]"
              },
              "frontendPort": 443,
              "backendPort": 443,
              "enableFloatingIP": false,
              "idleTimeoutInMinutes": 4,
              "protocol": "Tcp",
              "loadDistribution": "Default",
              "backendAddressPool": {
                "id": "[concat(variables('masterLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"
              },
              "probe": {
                "id": "[concat(variables('masterLbID'),'/probes/dcosMasterProbe')]"
              }
            }
          },
          {
            "name": "LBRule80",
            "properties": {
              "frontendIPConfiguration": {
                "id": "[variables('masterLbIPConfigID')]"
              },
              "frontendPort": 80,
              "backendPort": 80,
              "enableFloatingIP": false,
              "idleTimeoutInMinutes": 4,
              "protocol": "Tcp",
              "loadDistribution": "Default",
              "backendAddressPool": {
                "id": "[concat(variables('masterLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"
              },
              "probe": {
                "id": "[concat(variables('masterLbID'),'/probes/dcosMasterProbe')]"
              }
            }
          }
        ],
        "probes": [
          {
            "name": "dcosMasterProbe",
            "properties": {
              "protocol": "Http",
              "port": 5050,
              "requestPath": "/health",
              "intervalInSeconds": 5,
              "numberOfProbes": 2
            }
          }
        ]
{{end}}
      },
      "type": "Microsoft.Network/loadBalancers"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "copy": {
        "count": "[variables('masterCount')]",
        "name": "masterLbLoopNode"
      },
      "dependsOn": [
        "[variables('masterLbID')]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('masterLbName'), '/', 'SSH-', variables('masterVMNamePrefix'), copyIndex())]",
      "properties": {
        "backendPort": 22,
        "enableFloatingIP": false,
        "frontendIPConfiguration": {
          "id": "[variables('masterLbIPConfigID')]"
        },
        "frontendPort": "[copyIndex(2200)]",
        "protocol": "Tcp"
      },
      "type": "Microsoft.Network/loadBalancers/inboundNatRules"
    },
{{if IsDCOS19}}
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "dependsOn": [
        "[variables('masterLbID')]"
      ],
      "location": "[resourceGroup().location]",

      "name": "[concat(variables('masterLbName'), '/', 'SSHPort22-', variables('masterVMNamePrefix'), '0')]",
      "properties": {
        "backendPort": 2222,
        "enableFloatingIP": false,
        "frontendIPConfiguration": {
          "id": "[variables('masterLbIPConfigID')]"
        },
        "frontendPort": "22",
        "protocol": "Tcp"
      },
      "type": "Microsoft.Network/loadBalancers/inboundNatRules"
    },
{{end}}
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('masterNSGName')]",
      "properties": {
        "securityRules": [
{{if IsDCOS19}}
            {
                "properties": {
                    "priority": 201,
                    "access": "Allow",
                    "direction": "Inbound",
                    "destinationPortRange": "2222",
                    "sourcePortRange": "*",
                    "destinationAddressPrefix": "*",
                    "protocol": "Tcp",
                    "description": "Allow SSH",
                    "sourceAddressPrefix": "*"
                },
                "name": "sshPort22"
            },
{{if .MasterProfile.OAuthEnabled}}
            {
                "name": "http",
                "properties": {
                    "protocol": "Tcp",
                    "sourcePortRange": "*",
                    "destinationPortRange": "80",
                    "sourceAddressPrefix": "[variables('masterHttpSourceAddressPrefix')]",
                    "destinationAddressPrefix": "*",
                    "access": "Allow",
                    "priority": 202,
                    "direction": "Inbound"
                }
            },
            {
                "name": "https",
                "properties": {
                    "protocol": "Tcp",
                    "sourcePortRange": "*",
                    "destinationPortRange": "443",
                    "sourceAddressPrefix": "[variables('masterHttpSourceAddressPrefix')]",
                    "destinationAddressPrefix": "*",
                    "access": "Allow",
                    "priority": 203,
                    "direction": "Inbound"
                }
            },
{{end}}
{{end}}
            {
                "properties": {
                    "priority": 200,
                    "access": "Allow",
                    "direction": "Inbound",
                    "destinationPortRange": "22",
                    "sourcePortRange": "*",
                    "destinationAddressPrefix": "*",
                    "protocol": "Tcp",
                    "description": "Allow SSH",
                    "sourceAddressPrefix": "*"
                },
                "name": "ssh"
            }
        ]
      },
      "type": "Microsoft.Network/networkSecurityGroups"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "copy": {
        "count": "[variables('masterCount')]",
        "name": "nicLoopNode"
      },
      "dependsOn": [
        "[variables('masterNSGID')]",
{{if not .MasterProfile.IsCustomVNET}}
        "[variables('vnetID')]",
{{end}}
        "[variables('masterLbID')]",
{{if IsDCOS19}}
        "[concat(variables('masterLbID'),'/inboundNatRules/SSHPort22-',variables('masterVMNamePrefix'),0)]",
{{end}}
        "[concat(variables('masterLbID'),'/inboundNatRules/SSH-',variables('masterVMNamePrefix'),copyIndex())]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('masterVMNamePrefix'), 'nic-', copyIndex())]",
      "properties": {
        "ipConfigurations": [
          {
            "name": "ipConfigNode",
            "properties": {
              "loadBalancerBackendAddressPools": [
                {
                  "id": "[concat(variables('masterLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"
                }
              ],
{{if IsDCOS19}}
              "loadBalancerInboundNatRules": "[variables('masterLbInboundNatRules')[copyIndex()]]",
{{else}}
              "loadBalancerInboundNatRules": [
                {
                  "id": "[concat(variables('masterLbID'),'/inboundNatRules/SSH-',variables('masterVMNamePrefix'),copyIndex())]"
                }
              ],
{{end}}
              "privateIPAddress": "[concat(variables('masterFirstAddrPrefix'), copyIndex(int(variables('masterFirstAddrOctet4'))))]",
              "privateIPAllocationMethod": "Static",
              "subnet": {
                "id": "[variables('masterVnetSubnetID')]"
              }
            }
          }
        ]
        ,"networkSecurityGroup": {
          "id": "[variables('masterNSGID')]"
        }
      },
      "type": "Microsoft.Network/networkInterfaces"
    },
    {
{{if .MasterProfile.IsManagedDisks}}
      "apiVersion": "[variables('apiVersionStorageManagedDisks')]",
{{else}}
      "apiVersion": "[variables('apiVersionDefault')]",
{{end}}
      "copy": {
        "count": "[variables('masterCount')]",
        "name": "vmLoopNode"
      },
      "dependsOn": [
        "[concat('Microsoft.Network/networkInterfaces/', variables('masterVMNamePrefix'), 'nic-', copyIndex())]",
        "[concat('Microsoft.Compute/availabilitySets/',variables('masterAvailabilitySet'))]",
{{if .MasterProfile.IsStorageAccount}}
        "[variables('masterStorageAccountName')]",
{{end}}
        "[variables('masterStorageAccountExhibitorName')]"
      ],
      "tags":
      {
        "creationSource" : "[concat('acsengine-', variables('masterVMNamePrefix'), copyIndex())]"
      },
      "location": "[variables('location')]",
      "name": "[concat(variables('masterVMNamePrefix'), copyIndex())]",
      "properties": {
        "availabilitySet": {
          "id": "[resourceId('Microsoft.Compute/availabilitySets',variables('masterAvailabilitySet'))]"
        },
        "hardwareProfile": {
          "vmSize": "[variables('masterVMSize')]"
        },
        "networkProfile": {
          "networkInterfaces": [
            {
              "id": "[resourceId('Microsoft.Network/networkInterfaces',concat(variables('masterVMNamePrefix'), 'nic-', copyIndex()))]"
            }
          ]
        },
        "osProfile": {
          "adminUsername": "[variables('adminUsername')]",
          "computername": "[concat(variables('masterVMNamePrefix'), copyIndex())]",
          {{GetDCOSMasterCustomData}}
          "linuxConfiguration": {
            "disablePasswordAuthentication": true,
            "ssh": {
                "publicKeys": [
                    {
                        "keyData": "[variables('sshRSAPublicKey')]",
                        "path": "[variables('sshKeyPath')]"
                    }
                ]
            }
          }
          {{if .LinuxProfile.HasSecrets}}
          ,
          "secrets": "[variables('linuxProfileSecrets')]"
          {{end}}
        },
        "storageProfile": {
          "imageReference": {
            "offer": "[variables('osImageOffer')]",
            "publisher": "[variables('osImagePublisher')]",
            "sku": "[variables('osImageSKU')]",
            "version": "[variables('osImageVersion')]"
          },
          "osDisk": {
            "caching": "ReadWrite"
            ,"createOption": "FromImage"
{{if .MasterProfile.IsStorageAccount}}
            ,"name": "[concat(variables('masterVMNamePrefix'), copyIndex(),'-osdisk')]"
            ,"vhd": {
              "uri": "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('masterStorageAccountName')),variables('apiVersionStorage')).primaryEndpoints.blob,'vhds/',variables('masterVMNamePrefix'),copyIndex(),'-osdisk.vhd')]"
            }
{{end}}
{{if ne .MasterProfile.OSDiskSizeGB 0}}
            ,"diskSizeGB": {{.MasterProfile.OSDiskSizeGB}}
{{end}}
          }
        }
      },
      "type": "Microsoft.Compute/virtualMachines"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "dependsOn": [
        "[concat('Microsoft.Compute/virtualMachines/', variables('masterVMNamePrefix'), sub(variables('masterCount'), 1))]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('masterVMNamePrefix'), sub(variables('masterCount'), 1), '/waitforleader')]",
      "properties": {
        "autoUpgradeMinorVersion": true,
        "publisher": "Microsoft.OSTCExtensions",
        "settings": {
          "commandToExecute": "sh -c 'until ping -c1 leader.mesos;do echo waiting for leader.mesos;sleep 15;done;echo leader.mesos up'"
        },
        "type": "CustomScriptForLinux",
        "typeHandlerVersion": "1.4"
      },
      "type": "Microsoft.Compute/virtualMachines/extensions"
    }{{WriteLinkedTemplatesForExtensions}}
`)

func dcosDcosmasterresourcesTBytes() ([]byte, error) {
	return _dcosDcosmasterresourcesT, nil
}

func dcosDcosmasterresourcesT() (*asset, error) {
	bytes, err := dcosDcosmasterresourcesTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/dcosmasterresources.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosDcosmastervarsT = []byte(`    "adminUsername": "[parameters('linuxAdminUsername')]",
    "targetEnvironment": "[parameters('targetEnvironment')]",
    "maxVMsPerPool": 100,
    "apiVersionDefault": "2016-03-30",
    "apiVersionLinkDefault": "2015-01-01",
    "singleQuote": "'",
    "doubleSingleQuote": "''",
{{if .LinuxProfile.HasSecrets}}
    "linuxProfileSecrets" :
      [
          {{range  $vIndex, $vault := .LinuxProfile.Secrets}}
            {{if $vIndex}} , {{end}}
              {
                "sourceVault":{
                  "id":"[parameters('linuxKeyVaultID{{$vIndex}}')]"
                },
                "vaultCertificates":[
                {{range $cIndex, $cert := $vault.VaultCertificates}}
                  {{if $cIndex}} , {{end}}
                  {
                    "certificateUrl" :"[parameters('linuxKeyVaultID{{$vIndex}}CertificateURL{{$cIndex}}')]"
                  }
                {{end}}
                ]
              }
        {{end}}
      ],
{{end}}
    "orchestratorVersion": "{{.OrchestratorProfile.OrchestratorVersion}}",
{{if .HasWindows}}
    "windowsAdminUsername": "[parameters('windowsAdminUsername')]",
    "windowsAdminPassword": "[parameters('windowsAdminPassword')]",
    "agentWindowsBackendPort": 3389,
    "agentWindowsPublisher": "[parameters('agentWindowsPublisher')]",
    "agentWindowsOffer": "[parameters('agentWindowsOffer')]",
    "agentWindowsSku": "[parameters('agentWindowsSku')]",
    "agentWindowsVersion": "[parameters('agentWindowsVersion')]",
    "dcosWindowsBootstrapURL" : "[parameters('dcosWindowsBootstrapURL')]",
    "windowsCustomScriptSuffix": " $inputFile = '%SYSTEMDRIVE%\\AzureData\\CustomData.bin' ; $outputFile = '%SYSTEMDRIVE%\\AzureData\\dcosWindowsProvision.ps1' ; $inputStream = New-Object System.IO.FileStream $inputFile, ([IO.FileMode]::Open), ([IO.FileAccess]::Read), ([IO.FileShare]::Read) ; $sr = New-Object System.IO.StreamReader(New-Object System.IO.Compression.GZipStream($inputStream, [System.IO.Compression.CompressionMode]::Decompress)) ; $sr.ReadToEnd() | Out-File($outputFile) ; Invoke-Expression('{0} {1}' -f $outputFile, $arguments) ; ",
    "windowsMasterCustomScriptArguments": "[concat('$arguments = ', variables('singleQuote'),'-MasterCount ', variables('masterCount'), ' -firstMasterIP ', parameters('firstConsecutiveStaticIP'), variables('singleQuote'), ' ; ')]",

    "windowsMasterCustomScript": "[concat('powershell.exe -ExecutionPolicy Unrestricted -command \"', variables('windowsMasterCustomScriptArguments'), variables('windowsCustomScriptSuffix'), '\" > %SYSTEMDRIVE%\\AzureData\\dcosWindowsProvision.log 2>&1')]",
{{end}}
    "masterAvailabilitySet": "[concat(variables('orchestratorName'), '-master-availabilitySet-', variables('nameSuffix'))]",
    "masterCount": {{.MasterProfile.Count}},
    "masterEndpointDNSNamePrefix": "[tolower(parameters('masterEndpointDNSNamePrefix'))]",
    "masterHttpSourceAddressPrefix": "{{.MasterProfile.HTTPSourceAddressPrefix}}",
    "masterLbBackendPoolName": "[concat(variables('orchestratorName'), '-master-pool-', variables('nameSuffix'))]",
    "masterLbID": "[resourceId('Microsoft.Network/loadBalancers',variables('masterLbName'))]",
    "masterLbIPConfigID": "[concat(variables('masterLbID'),'/frontendIPConfigurations/', variables('masterLbIPConfigName'))]",
    "masterLbIPConfigName": "[concat(variables('orchestratorName'), '-master-lbFrontEnd-', variables('nameSuffix'))]",
    "masterLbName": "[concat(variables('orchestratorName'), '-master-lb-', variables('nameSuffix'))]",
    "masterNSGID": "[resourceId('Microsoft.Network/networkSecurityGroups',variables('masterNSGName'))]",
    "masterNSGName": "[concat(variables('orchestratorName'), '-master-nsg-', variables('nameSuffix'))]",
    "masterPublicIPAddressName": "[concat(variables('orchestratorName'), '-master-ip-', variables('masterEndpointDNSNamePrefix'), '-', variables('nameSuffix'))]",
    "apiVersionStorage": "2015-06-15",

    "storageAccountBaseName": "[uniqueString(concat(variables('masterEndpointDNSNamePrefix'),variables('location'),variables('orchestratorName')))]",
    "masterStorageAccountExhibitorName": "[concat(variables('storageAccountBaseName'), 'exhb0')]",
    "storageAccountType": "Standard_LRS",
{{if .HasStorageAccountDisks}}
    "maxVMsPerStorageAccount": 20,
    "maxStorageAccountsPerAgent": "[div(variables('maxVMsPerPool'),variables('maxVMsPerStorageAccount'))]",
    "dataStorageAccountPrefixSeed": 97,
    "storageAccountPrefixes": [ "0", "6", "c", "i", "o", "u", "1", "7", "d", "j", "p", "v", "2", "8", "e", "k", "q", "w", "3", "9", "f", "l", "r", "x", "4", "a", "g", "m", "s", "y", "5", "b", "h", "n", "t", "z" ],
    "storageAccountPrefixesCount": "[length(variables('storageAccountPrefixes'))]",
    {{GetSizeMap}},
{{else}}
    "storageAccountPrefixes": [],
{{end}}
{{if .HasManagedDisks}}
    "apiVersionStorageManagedDisks": "2016-04-30-preview",
{{end}}
{{if .MasterProfile.IsStorageAccount}}
    "masterStorageAccountName": "[concat(variables('storageAccountBaseName'), 'mstr0')]",
{{end}}
{{if .MasterProfile.IsCustomVNET}}
    "masterVnetSubnetID": "[parameters('masterVnetSubnetID')]",
{{else}}
    "masterSubnet": "[parameters('masterSubnet')]",
    "masterSubnetName": "[concat(variables('orchestratorName'), '-masterSubnet')]",
    "vnetID": "[resourceId('Microsoft.Network/virtualNetworks',variables('virtualNetworkName'))]",
    "masterVnetSubnetID": "[concat(variables('vnetID'),'/subnets/',variables('masterSubnetName'))]",
    "virtualNetworkName": "[concat(variables('orchestratorName'), '-vnet-', variables('nameSuffix'))]",
{{end}}
    "masterFirstAddrOctets": "[split(parameters('firstConsecutiveStaticIP'),'.')]",
    "masterFirstAddrOctet4": "[variables('masterFirstAddrOctets')[3]]",
    "masterFirstAddrPrefix": "[concat(variables('masterFirstAddrOctets')[0],'.',variables('masterFirstAddrOctets')[1],'.',variables('masterFirstAddrOctets')[2],'.')]",
    "masterVMNamePrefix": "[concat(variables('orchestratorName'), '-master-', variables('nameSuffix'), '-')]",
    "masterVMNic": [
      "[concat(variables('masterVMNamePrefix'), 'nic-0')]",
      "[concat(variables('masterVMNamePrefix'), 'nic-1')]",
      "[concat(variables('masterVMNamePrefix'), 'nic-2')]",
      "[concat(variables('masterVMNamePrefix'), 'nic-3')]",
      "[concat(variables('masterVMNamePrefix'), 'nic-4')]",
      "[concat(variables('masterVMNamePrefix'), 'nic-5')]",
      "[concat(variables('masterVMNamePrefix'), 'nic-6')]"
    ],
    "masterVMSize": "[parameters('masterVMSize')]",
    "nameSuffix": "[parameters('nameSuffix')]",
    "oauthEnabled": "{{.MasterProfile.OAuthEnabled}}",
    "orchestratorName": "dcos",
    "osImageOffer": "[parameters('osImageOffer')]",
    "osImagePublisher": "[parameters('osImagePublisher')]",
    "osImageSKU": "[parameters('osImageSKU')]",
    "osImageVersion": "[parameters('osImageVersion')]",
    "sshKeyPath": "[concat('/home/', variables('adminUsername'), '/.ssh/authorized_keys')]",
    "sshRSAPublicKey": "[parameters('sshRSAPublicKey')]",
    "locations": [
         "[resourceGroup().location]",
         "[parameters('location')]"
    ],
    "location": "[variables('locations')[mod(add(2,length(parameters('location'))),add(1,length(parameters('location'))))]]",
{{if IsDCOS19}}
    "masterSshInboundNatRuleIdPrefix": "[concat(variables('masterLbID'),'/inboundNatRules/SSH-',variables('masterVMNamePrefix'))]",
    "masterSshPort22InboundNatRuleIdPrefix": "[concat(variables('masterLbID'),'/inboundNatRules/SSHPort22-',variables('masterVMNamePrefix'))]",
    "masterLbInboundNatRules": [
            [
                {
                    "id": "[concat(variables('masterSshInboundNatRuleIdPrefix'),'0')]"
                },
                {
                    "id": "[concat(variables('masterSshPort22InboundNatRuleIdPrefix'),'0')]"
                }
            ],
            [
                {
                    "id": "[concat(variables('masterSshInboundNatRuleIdPrefix'),'1')]"
                }
            ],
            [
                {
                    "id": "[concat(variables('masterSshInboundNatRuleIdPrefix'),'2')]"
                }
            ],
            [
                {
                    "id": "[concat(variables('masterSshInboundNatRuleIdPrefix'),'3')]"
                }
            ],
            [
                {
                    "id": "[concat(variables('masterSshInboundNatRuleIdPrefix'),'4')]"
                }
            ]
        ],
{{end}}
    "dcosBootstrapURL": "[parameters('dcosBootstrapURL')]"

`)

func dcosDcosmastervarsTBytes() ([]byte, error) {
	return _dcosDcosmastervarsT, nil
}

func dcosDcosmastervarsT() (*asset, error) {
	bytes, err := dcosDcosmastervarsTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/dcosmastervars.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosDcosparamsT = []byte(`    "dcosBootstrapURL": {
      "defaultValue": "https://dcosio.azureedge.net/dcos/stable/bootstrap/58fd0833ce81b6244fc73bf65b5deb43217b0bd7.bootstrap.tar.xz",
      "metadata": {
        "description": "The default mesosphere bootstrap package."
      },
      "type": "string"
    },
    "dcosWindowsBootstrapURL": {
      "defaultValue": "http://dcos-win.westus.cloudapp.azure.com/dcos-windows/stable/",
      "metadata": {
        "description": "The default mesosphere bootstrap package location for windows."
      },
      "type": "string"
    },
    "dcosRepositoryURL": {
      "defaultValue": "https://dcosio.azureedge.net/dcos/stable",
      "metadata": {
        "description": "The repository URL"
      }, 
      "type": "string"
    },
    "dcosClusterPackageListID": {
      "defaultValue": "77282d8864a5bf36db345b54a0d1de3674a0e937",
      "metadata": {
        "description": "The default cluster package list IDs."
      }, 
      "type": "string"
    },
    "dcosProviderPackageID": {
      "defaultValue": "",
      "metadata": {
        "description": "The guid for provider dcos-provider package."
      }, 
      "type": "string"
    },
`)

func dcosDcosparamsTBytes() ([]byte, error) {
	return _dcosDcosparamsT, nil
}

func dcosDcosparamsT() (*asset, error) {
	bytes, err := dcosDcosparamsTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/dcosparams.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosDcosprovisionSh = []byte(`#!/bin/bash

MESOSDIR=/var/lib/mesos/dl
mkdir $MESOSDIR

# load the env vars
. /etc/mesosphere/setup-flags/dcos-deploy-environment

# default dc/os component download address (Azure CDN)
DOCKER_ENGINE_DOWNLOAD_URL=https://mesosphere.blob.core.windows.net/dcos-deps/docker-engine_1.13.1-0-ubuntu-xenial_amd64.deb
LIBIPSET_DOWNLOAD_URL=https://az837203.vo.msecnd.net/dcos-deps/libipset3_6.29-1_amd64.deb
IPSET_DOWNLOAD_URL=https://az837203.vo.msecnd.net/dcos-deps/ipset_6.29-1_amd64.deb
UNZIP_DOWNLOAD_URL=https://az837203.vo.msecnd.net/dcos-deps/unzip_6.0-20ubuntu1_amd64.deb
LIBLTDL_DOWNLOAD_URL=https://az837203.vo.msecnd.net/dcos-deps/libltdl7_2.4.6-0.1_amd64.deb

case $DCOS_ENVIRONMENT in
    # because of Chinese GreatWall Firewall, the default packages on Azure CDN is blocked. So the following Chinese local mirror url should be used instead.
    AzureChinaCloud)
        DOCKER_ENGINE_DOWNLOAD_URL=http://acsengine.blob.core.chinacloudapi.cn/dcos/docker-engine_1.11.2-0~xenial_amd64.deb
        LIBIPSET_DOWNLOAD_URL=http://acsengine.blob.core.chinacloudapi.cn/dcos/libipset3_6.29-1_amd64.deb
        IPSET_DOWNLOAD_URL=http://acsengine.blob.core.chinacloudapi.cn/dcos/ipset_6.29-1_amd64.deb
        UNZIP_DOWNLOAD_URL=http://acsengine.blob.core.chinacloudapi.cn/dcos/unzip_6.0-20ubuntu1_amd64.deb
        LIBLTDL_DOWNLOAD_URL=http://acsengine.blob.core.chinacloudapi.cn/dcos/libltdl7_2.4.6-0.1_amd64.deb
    ;;
esac

curl -fLsSv --retry 20 -Y 100000 -y 60 -o $MESOSDIR/d.deb $DOCKER_ENGINE_DOWNLOAD_URL &
curl -fLsSv --retry 20 -Y 100000 -y 60 -o $MESOSDIR/1.deb $LIBIPSET_DOWNLOAD_URL &
curl -fLsSv --retry 20 -Y 100000 -y 60 -o $MESOSDIR/2.deb $IPSET_DOWNLOAD_URL &
curl -fLsSv --retry 20 -Y 100000 -y 60 -o $MESOSDIR/3.deb $UNZIP_DOWNLOAD_URL &
curl -fLsSv --retry 20 -Y 100000 -y 60 -o $MESOSDIR/4.deb $LIBLTDL_DOWNLOAD_URL &
wait

for i in {1..300}; do
    dpkg -i $MESOSDIR/{1,2,3,4}.deb
    if [ "$?" = "0" ]
    then
        echo "succeeded"
        break
    fi
    sleep 1
done

ROLESFILECONTENTS

# add Azure update domain and fault domain attributes
ud=$( curl -H Metadata:true "http://169.254.169.254/metadata/instance/compute/platformUpdateDomain?api-version=2017-04-02&format=text" )
fd=$( curl -H Metadata:true "http://169.254.169.254/metadata/instance/compute/platformFaultDomain?api-version=2017-04-02&format=text" )
echo ";azure.faultdomain:$fd;azure.updatedomain:$ud" >> /var/lib/dcos/mesos-slave-common`)

func dcosDcosprovisionShBytes() ([]byte, error) {
	return _dcosDcosprovisionSh, nil
}

func dcosDcosprovisionSh() (*asset, error) {
	bytes, err := dcosDcosprovisionShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/dcosprovision.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dcosDcosprovisionsourceSh = []byte(`function retry_download() {
  retries=$1; wait_sleep=$2; timeout=$3; url=$4; path=$5 checksum=$6
  for i in $(seq 1 $retries); do
    rm -f $path
    timeout $timeout curl -fsSL $url -o $path
    if [ $? -ne 0 ]; then
      echo "retry_download[$i] Error: Failed to execute curl -fsSL $url -o $path"
      sleep $wait_sleep
      continue
    fi
    if [ ! -z "${checksum:-}" ]; then
      actual=$(sha1sum -b $path | cut -f 1 -d " ")
      if [ $? -ne 0 ]; then
        echo "retry_download[$i] Error: Failed to execute sha1sum -b $path (per $url)"
        sleep $wait_sleep
        continue
      fi
      if [ "$checksum" != "$actual" ]; then
        echo "retry_download[$i] Error: sha1sum mismatch for $url"
        sleep $wait_sleep
        continue
      fi
    fi
    return 0
  done
  return 1
}

function retrycmd_if_failure() {
    retries=$1; wait_sleep=$2; timeout=$3; shift && shift && shift
    for i in $(seq 1 $retries); do
        timeout $timeout ${@}
        [ $? -eq 0 ] && break || \
        if [ $i -eq $retries ]; then
            echo "Error: Failed to execute \"$@\" after $i attempts"
            return 1
        else
            sleep $wait_sleep
        fi
    done
    echo Executed \"$@\" $i times;
}

function retry_get_install_deb() {
  retries=$1; wait_sleep=$2; timeout=$3; url=$4; checksum=$5
  deb=$(mktemp)
  trap "rm -f $deb" RETURN
  retry_download $retries $wait_sleep $timeout $url $deb $checksum
  if [ $? -ne 0 ]; then
    echo "Error: Failed to download $url"
    return 1
  fi
  retrycmd_if_failure $retries $wait_sleep $timeout dpkg -i $deb
  if [ $? -ne 0 ]; then
    echo "Error: Failed to install $url"
    return 1
  fi
}
`)

func dcosDcosprovisionsourceShBytes() ([]byte, error) {
	return _dcosDcosprovisionsourceSh, nil
}

func dcosDcosprovisionsourceSh() (*asset, error) {
	bytes, err := dcosDcosprovisionsourceShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dcos/dcosprovisionsource.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _iaasoutputsT = []byte(`    "resourceGroup": {
        "type": "string",
        "value": "[variables('resourceGroup')]"
    },
    "vnetResourceGroup": {
        "type": "string",
        "value": "[variables('virtualNetworkResourceGroupName')]"
    },
    "subnetName": {
        "type": "string",
        "value": "[variables('subnetName')]"
    },
    "securityGroupName": {
        "type": "string",
        "value": "[variables('nsgName')]"
    },
    "virtualNetworkName": {
        "type": "string",
        "value": "[variables('virtualNetworkName')]"
    },
    "routeTableName": {
        "type": "string",
        "value": "[variables('routeTableName')]"
    },
    "primaryAvailabilitySetName": {
        "type": "string",
        "value": "[variables('primaryAvailabilitySetName')]"
    },
    "primaryScaleSetName": {
        "type": "string",
        "value": "[variables('primaryScaleSetName')]"
    }

`)

func iaasoutputsTBytes() ([]byte, error) {
	return _iaasoutputsT, nil
}

func iaasoutputsT() (*asset, error) {
	bytes, err := iaasoutputsTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "iaasoutputs.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddons110KubernetesmasteraddonsKubeDnsDeploymentYaml = []byte(`# Copyright 2016 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Should keep target in cluster/addons/dns-horizontal-autoscaler/dns-horizontal-autoscaler.yaml
# in sync with this file.

apiVersion: v1
kind: Service
metadata:
  name: kube-dns
  namespace: kube-system
  labels:
    k8s-app: kube-dns
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
    kubernetes.io/name: "KubeDNS"
spec:
  selector:
    k8s-app: kube-dns
  clusterIP: <clustIP>
  ports:
  - name: dns
    port: 53
    protocol: UDP
  - name: dns-tcp
    port: 53
    protocol: TCP
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: kube-dns
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: kube-dns
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: EnsureExists
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: kube-dns
  namespace: kube-system
  labels:
    k8s-app: kube-dns
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  # replicas: not specified here:
  # 1. In order to make Addon Manager do not reconcile this replicas parameter.
  # 2. Default is 1.
  # 3. Will be tuned in real time if DNS horizontal auto-scaling is turned on.
  strategy:
    rollingUpdate:
      maxSurge: 10%
      maxUnavailable: 0
  selector:
    matchLabels:
      k8s-app: kube-dns
  template:
    metadata:
      labels:
        k8s-app: kube-dns
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ''
    spec:
      priorityClassName: system-node-critical
      tolerations:
      - key: "CriticalAddonsOnly"
        operator: "Exists"
      volumes:
      - name: kube-dns-config
        configMap:
          name: kube-dns
          optional: true
      containers:
      - name: kubedns
        image: <img>
        imagePullPolicy: IfNotPresent
        resources:
          # TODO: Set memory limits when we've profiled the container for large
          # clusters, then set request = limit to keep this container in
          # guaranteed class. Currently, this container falls into the
          # "burstable" category so the kubelet doesn't backoff from restarting it.
          limits:
            memory: 170Mi
          requests:
            cpu: 100m
            memory: 70Mi
        livenessProbe:
          httpGet:
            path: /healthcheck/kubedns
            port: 10054
            scheme: HTTP
          initialDelaySeconds: 60
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 5
        readinessProbe:
          httpGet:
            path: /readiness
            port: 8081
            scheme: HTTP
          initialDelaySeconds: 3
          timeoutSeconds: 5
        args:
        - --domain=<domain>.
        - --dns-port=10053
        - --config-dir=/kube-dns-config
        - --v=2
        env:
        - name: PROMETHEUS_PORT
          value: "10055"
        ports:
        - containerPort: 10053
          name: dns-local
          protocol: UDP
        - containerPort: 10053
          name: dns-tcp-local
          protocol: TCP
        - containerPort: 10055
          name: metrics
          protocol: TCP
        volumeMounts:
        - name: kube-dns-config
          mountPath: /kube-dns-config
      - name: dnsmasq
        image: <imgMasq>
        imagePullPolicy: IfNotPresent
        livenessProbe:
          httpGet:
            path: /healthcheck/dnsmasq
            port: 10054
            scheme: HTTP
          initialDelaySeconds: 60
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 5
        args:
        - -v=2
        - -logtostderr
        - -configDir=/etc/k8s/dns/dnsmasq-nanny
        - -restartDnsmasq=true
        - --
        - -k
        - --cache-size=1000
        - --no-negcache
        - --log-facility=-
        - --server=/cluster.local/127.0.0.1#10053
        - --server=/in-addr.arpa/127.0.0.1#10053
        - --server=/ip6.arpa/127.0.0.1#10053
        ports:
        - containerPort: 53
          name: dns
          protocol: UDP
        - containerPort: 53
          name: dns-tcp
          protocol: TCP
        resources:
          requests:
            cpu: 150m
            memory: 20Mi
        volumeMounts:
        - name: kube-dns-config
          mountPath: /etc/k8s/dns/dnsmasq-nanny
      - name: sidecar
        image: <imgSidecar>
        imagePullPolicy: IfNotPresent
        livenessProbe:
          httpGet:
            path: /metrics
            port: 10054
            scheme: HTTP
          initialDelaySeconds: 60
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 5
        args:
        - --v=2
        - --logtostderr
        - --probe=kubedns,127.0.0.1:10053,kubernetes.default.svc.<domain>,5,SRV
        - --probe=dnsmasq,127.0.0.1:53,kubernetes.default.svc.<domain>,5,SRV
        ports:
        - containerPort: 10054
          name: metrics
          protocol: TCP
        resources:
          requests:
            memory: 20Mi
            cpu: 10m
      dnsPolicy: Default
      serviceAccountName: kube-dns
      nodeSelector:
        beta.kubernetes.io/os: linux`)

func k8sAddons110KubernetesmasteraddonsKubeDnsDeploymentYamlBytes() ([]byte, error) {
	return _k8sAddons110KubernetesmasteraddonsKubeDnsDeploymentYaml, nil
}

func k8sAddons110KubernetesmasteraddonsKubeDnsDeploymentYaml() (*asset, error) {
	bytes, err := k8sAddons110KubernetesmasteraddonsKubeDnsDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/1.10/kubernetesmasteraddons-kube-dns-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddons115KubernetesmasteraddonsAzureCloudProviderDeploymentYaml = []byte(`---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
  name: system:azure-cloud-provider
rules:
- apiGroups: [""]
  resources: ["events"]
  verbs:
  - create
  - patch
  - update
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
  name: system:azure-cloud-provider
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:azure-cloud-provider
subjects:
- kind: ServiceAccount
  name: azure-cloud-provider
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: system:azure-persistent-volume-binder
  labels:
    kubernetes.io/cluster-service: "true"
rules:
- apiGroups: ['']
  resources: ['secrets']
  verbs:     ['get','create']
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: system:azure-persistent-volume-binder
  labels:
    kubernetes.io/cluster-service: "true"
roleRef:
  kind: ClusterRole
  apiGroup: rbac.authorization.k8s.io
  name: system:azure-persistent-volume-binder
subjects:
- kind: ServiceAccount
  name: persistent-volume-binder
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
  name: system:azure-cloud-provider-secret-getter
rules:
- apiGroups: [""]
  resources: ["secrets"]
  verbs:
  - get
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
  name: system:azure-cloud-provider-secret-getter
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:azure-cloud-provider-secret-getter
subjects:
- kind: ServiceAccount
  name: azure-cloud-provider
  namespace: kube-system
`)

func k8sAddons115KubernetesmasteraddonsAzureCloudProviderDeploymentYamlBytes() ([]byte, error) {
	return _k8sAddons115KubernetesmasteraddonsAzureCloudProviderDeploymentYaml, nil
}

func k8sAddons115KubernetesmasteraddonsAzureCloudProviderDeploymentYaml() (*asset, error) {
	bytes, err := k8sAddons115KubernetesmasteraddonsAzureCloudProviderDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/1.15/kubernetesmasteraddons-azure-cloud-provider-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddons115KubernetesmasteraddonsPodSecurityPolicyYaml = []byte(`apiVersion: extensions/v1beta1
kind: PodSecurityPolicy
metadata:
  name: privileged
  annotations:
    seccomp.security.alpha.kubernetes.io/allowedProfileNames: "*"
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  privileged: true
  allowPrivilegeEscalation: true
  allowedCapabilities:
  - "*"
  volumes:
  - "*"
  hostNetwork: true
  hostPorts:
  - min: 0
    max: 65535
  hostIPC: true
  hostPID: true
  runAsUser:
    rule: RunAsAny
  seLinux:
    rule: RunAsAny
  supplementalGroups:
    rule: RunAsAny
  fsGroup:
    rule: RunAsAny
---
apiVersion: extensions/v1beta1
kind: PodSecurityPolicy
metadata:
  name: restricted
  annotations:
    seccomp.security.alpha.kubernetes.io/allowedProfileNames: docker/default
    apparmor.security.beta.kubernetes.io/allowedProfileNames: runtime/default
    seccomp.security.alpha.kubernetes.io/defaultProfileName:  docker/default
    apparmor.security.beta.kubernetes.io/defaultProfileName:  runtime/default
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  privileged: false
  allowPrivilegeEscalation: false
  requiredDropCapabilities:
    - ALL
  volumes:
    - configMap
    - emptyDir
    - projected
    - secret
    - downwardAPI
    - persistentVolumeClaim
  hostNetwork: false
  hostIPC: false
  hostPID: false
  runAsUser:
    rule: MustRunAsNonRoot
  seLinux:
    rule: RunAsAny
  supplementalGroups:
    rule: MustRunAs
    ranges:
      # Forbid adding the root group.
      - min: 1
        max: 65535
  fsGroup:
    rule: MustRunAs
    ranges:
      # Forbid adding the root group.
      - min: 1
        max: 65535
  readOnlyRootFilesystem: false
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: psp:privileged
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups: ['extensions']
  resources: ['podsecuritypolicies']
  verbs:     ['use']
  resourceNames:
  - privileged
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: psp:restricted
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups: ['extensions']
  resources: ['podsecuritypolicies']
  verbs:     ['use']
  resourceNames:
  - restricted
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: default:restricted
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: psp:restricted
subjects:
- kind: Group
  name: system:authenticated
  apiGroup: rbac.authorization.k8s.io
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: default:privileged
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: psp:privileged
subjects:
- kind: Group
  name: system:authenticated
  apiGroup: rbac.authorization.k8s.io
- kind: Group
  name: system:nodes
  apiGroup: rbac.authorization.k8s.io
`)

func k8sAddons115KubernetesmasteraddonsPodSecurityPolicyYamlBytes() ([]byte, error) {
	return _k8sAddons115KubernetesmasteraddonsPodSecurityPolicyYaml, nil
}

func k8sAddons115KubernetesmasteraddonsPodSecurityPolicyYaml() (*asset, error) {
	bytes, err := k8sAddons115KubernetesmasteraddonsPodSecurityPolicyYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/1.15/kubernetesmasteraddons-pod-security-policy.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddons116KubernetesmasterAuditPolicyYaml = []byte(`apiVersion: audit.k8s.io/v1
kind: Policy
omitStages:
  - RequestReceived
rules:
  - level: RequestResponse
    resources:
    - group: ""
      resources: ["pods"]
  - level: Metadata
    resources:
    - group: ""
      resources: ["pods/log", "pods/status"]
  - level: None
    users: ["system:kube-proxy"]
    verbs: ["watch"]
    resources:
    - group: ""
      resources: ["endpoints", "services"]
  - level: None
    userGroups: ["system:authenticated"]
    nonResourceURLs:
    - /api*
    - /version
  - level: Request
    resources:
    - group: ""
      resources: ["configmaps"]
    namespaces: ["kube-system"]
  - level: Request
    resources:
    - group: ""
      resources: ["secrets"]
  - level: Request
    resources:
    - group: ""
    - group: extensions
  - level: Metadata
    omitStages:
      - RequestReceived
`)

func k8sAddons116KubernetesmasterAuditPolicyYamlBytes() ([]byte, error) {
	return _k8sAddons116KubernetesmasterAuditPolicyYaml, nil
}

func k8sAddons116KubernetesmasterAuditPolicyYaml() (*asset, error) {
	bytes, err := k8sAddons116KubernetesmasterAuditPolicyYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/1.16/kubernetesmaster-audit-policy.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddons116KubernetesmasteraddonsAzureCloudProviderDeploymentYaml = []byte(`---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
  name: system:azure-cloud-provider
rules:
- apiGroups: [""]
  resources: ["events"]
  verbs:
  - create
  - patch
  - update
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
  name: system:azure-cloud-provider
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:azure-cloud-provider
subjects:
- kind: ServiceAccount
  name: azure-cloud-provider
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: system:azure-persistent-volume-binder
  labels:
    kubernetes.io/cluster-service: "true"
rules:
- apiGroups: ['']
  resources: ['secrets']
  verbs:     ['get','create']
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system:azure-persistent-volume-binder
  labels:
    kubernetes.io/cluster-service: "true"
roleRef:
  kind: ClusterRole
  apiGroup: rbac.authorization.k8s.io
  name: system:azure-persistent-volume-binder
subjects:
- kind: ServiceAccount
  name: persistent-volume-binder
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
  name: system:azure-cloud-provider-secret-getter
rules:
- apiGroups: [""]
  resources: ["secrets"]
  verbs:
  - get
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
  name: system:azure-cloud-provider-secret-getter
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:azure-cloud-provider-secret-getter
subjects:
- kind: ServiceAccount
  name: azure-cloud-provider
  namespace: kube-system
`)

func k8sAddons116KubernetesmasteraddonsAzureCloudProviderDeploymentYamlBytes() ([]byte, error) {
	return _k8sAddons116KubernetesmasteraddonsAzureCloudProviderDeploymentYaml, nil
}

func k8sAddons116KubernetesmasteraddonsAzureCloudProviderDeploymentYaml() (*asset, error) {
	bytes, err := k8sAddons116KubernetesmasteraddonsAzureCloudProviderDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/1.16/kubernetesmasteraddons-azure-cloud-provider-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddons116KubernetesmasteraddonsCiliumDaemonsetYaml = []byte(`---
apiVersion: v1
kind: ConfigMap
metadata:
  name: cilium-config
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "EnsureExists"
data:
  etcd-config: |-
    ---
    endpoints:
      - https://cilium-etcd-client.kube-system.svc:2379
    ca-file: '/var/lib/etcd-secrets/etcd-client-ca.crt'
    key-file: '/var/lib/etcd-secrets/etcd-client.key'
    cert-file: '/var/lib/etcd-secrets/etcd-client.crt'
  debug: "false"
  enable-ipv4: "true"
  enable-ipv6: "false"
  clean-cilium-state: "false"
  clean-cilium-bpf-state: "false"
  monitor-aggregation-level: "none"
  ct-global-max-entries-tcp: "524288"
  ct-global-max-entries-other: "262144"
  preallocate-bpf-maps: "false"
  sidecar-istio-proxy-image: "cilium/istio_proxy"
  tunnel: "vxlan"
  cluster-name: default
  flannel-master-device: ""
  flannel-uninstall-on-exit: "false"
  flannel-manage-existing-containers: "false"
  tofqdns-enable-poller: "false"
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  labels:
    k8s-app: cilium
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: cilium
  namespace: kube-system
spec:
  selector:
    matchLabels:
      k8s-app: cilium
      kubernetes.io/cluster-service: "true"
  template:
    metadata:
      annotations:
        prometheus.io/port: "9090"
        prometheus.io/scrape: "true"
        scheduler.alpha.kubernetes.io/tolerations: '[{"key":"dedicated","operator":"Equal","value":"master","effect":"NoSchedule"}]'
      labels:
        k8s-app: cilium
        kubernetes.io/cluster-service: "true"
    spec:
      containers:
      - args:
        - --debug=$(CILIUM_DEBUG)
        - --kvstore=etcd
        - --kvstore-opt=etcd.config=/var/lib/etcd-config/etcd.config
        command:
        - cilium-agent
        env:
        - name: K8S_NODE_NAME
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: spec.nodeName
        - name: CILIUM_K8S_NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        - name: CILIUM_DEBUG
          valueFrom:
            configMapKeyRef:
              key: debug
              name: cilium-config
        - name: CILIUM_ENABLE_IPV4
          valueFrom:
            configMapKeyRef:
              key: enable-ipv4
              name: cilium-config
              optional: true
        - name: CILIUM_ENABLE_IPV6
          valueFrom:
            configMapKeyRef:
              key: enable-ipv6
              name: cilium-config
              optional: true
        - name: CILIUM_PROMETHEUS_SERVE_ADDR
          valueFrom:
            configMapKeyRef:
              key: prometheus-serve-addr
              name: cilium-metrics-config
              optional: true
        - name: CILIUM_LEGACY_HOST_ALLOWS_WORLD
          valueFrom:
            configMapKeyRef:
              key: legacy-host-allows-world
              name: cilium-config
              optional: true
        - name: CILIUM_SIDECAR_ISTIO_PROXY_IMAGE
          valueFrom:
            configMapKeyRef:
              key: sidecar-istio-proxy-image
              name: cilium-config
              optional: true
        - name: CILIUM_TUNNEL
          valueFrom:
            configMapKeyRef:
              key: tunnel
              name: cilium-config
              optional: true
        - name: CILIUM_MONITOR_AGGREGATION_LEVEL
          valueFrom:
            configMapKeyRef:
              key: monitor-aggregation-level
              name: cilium-config
              optional: true
        - name: CILIUM_CLUSTERMESH_CONFIG
          value: /var/lib/cilium/clustermesh/
        - name: CILIUM_CLUSTER_NAME
          valueFrom:
            configMapKeyRef:
              key: cluster-name
              name: cilium-config
              optional: true
        - name: CILIUM_CLUSTER_ID
          valueFrom:
            configMapKeyRef:
              key: cluster-id
              name: cilium-config
              optional: true
        - name: CILIUM_GLOBAL_CT_MAX_TCP
          valueFrom:
            configMapKeyRef:
              key: ct-global-max-entries-tcp
              name: cilium-config
              optional: true
        - name: CILIUM_GLOBAL_CT_MAX_ANY
          valueFrom:
            configMapKeyRef:
              key: ct-global-max-entries-other
              name: cilium-config
              optional: true
        - name: CILIUM_PREALLOCATE_BPF_MAPS
          valueFrom:
            configMapKeyRef:
              key: preallocate-bpf-maps
              name: cilium-config
              optional: true
        - name: CILIUM_FLANNEL_MASTER_DEVICE
          valueFrom:
            configMapKeyRef:
              key: flannel-master-device
              name: cilium-config
              optional: true
        - name: CILIUM_FLANNEL_UNINSTALL_ON_EXIT
          valueFrom:
            configMapKeyRef:
              key: flannel-uninstall-on-exit
              name: cilium-config
              optional: true
        - name: CILIUM_FLANNEL_MANAGE_EXISTING_CONTAINERS
          valueFrom:
            configMapKeyRef:
              key: flannel-manage-existing-containers
              name: cilium-config
              optional: true
        - name: CILIUM_DATAPATH_MODE
          valueFrom:
            configMapKeyRef:
              key: datapath-mode
              name: cilium-config
              optional: true
        - name: CILIUM_IPVLAN_MASTER_DEVICE
          valueFrom:
            configMapKeyRef:
              key: ipvlan-master-device
              name: cilium-config
              optional: true
        - name: CILIUM_INSTALL_IPTABLES_RULES
          valueFrom:
            configMapKeyRef:
              key: install-iptables-rules
              name: cilium-config
              optional: true
        - name: CILIUM_MASQUERADE
          valueFrom:
            configMapKeyRef:
              key: masquerade
              name: cilium-config
              optional: true
        - name: CILIUM_AUTO_DIRECT_NODE_ROUTES
          valueFrom:
            configMapKeyRef:
              key: auto-direct-node-routes
              name: cilium-config
              optional: true
        - name: CILIUM_TOFQDNS_ENABLE_POLLER
          valueFrom:
            configMapKeyRef:
              key: tofqdns-enable-poller
              name: cilium-config
              optional: true
        image: docker.io/cilium/cilium:v1.4
        imagePullPolicy: Always
        lifecycle:
          postStart:
            exec:
              command:
              - /cni-install.sh
          preStop:
            exec:
              command:
              - /cni-uninstall.sh
        livenessProbe:
          exec:
            command:
            - cilium
            - status
          failureThreshold: 10
          initialDelaySeconds: 120
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 1
        name: cilium-agent
        ports:
        - containerPort: 9090
          hostPort: 9090
          name: prometheus
          protocol: TCP
        readinessProbe:
          exec:
            command:
            - cilium
            - status
          failureThreshold: 3
          initialDelaySeconds: 5
          periodSeconds: 5
          successThreshold: 1
          timeoutSeconds: 1
        securityContext:
          capabilities:
            add:
            - NET_ADMIN
          privileged: true
        volumeMounts:
        - mountPath: /sys/fs/bpf
          name: bpf-maps
        - mountPath: /var/run/cilium
          name: cilium-run
        - mountPath: /host/opt/cni/bin
          name: cni-path
        - mountPath: /host/etc/cni/net.d
          name: etc-cni-netd
        - mountPath: /var/run/docker.sock
          name: docker-socket
          readOnly: true
        - mountPath: /var/lib/etcd-config
          name: etcd-config-path
          readOnly: true
        - mountPath: /var/lib/etcd-secrets
          name: etcd-secrets
          readOnly: true
        - mountPath: /var/lib/cilium/clustermesh
          name: clustermesh-secrets
          readOnly: true
      dnsPolicy: ClusterFirstWithHostNet
      hostNetwork: true
      hostPID: false
      initContainers:
      - command:
        - /init-container.sh
        env:
        - name: CLEAN_CILIUM_STATE
          valueFrom:
            configMapKeyRef:
              key: clean-cilium-state
              name: cilium-config
              optional: true
        - name: CLEAN_CILIUM_BPF_STATE
          valueFrom:
            configMapKeyRef:
              key: clean-cilium-bpf-state
              name: cilium-config
              optional: true
        image: docker.io/cilium/cilium-init:2018-10-16
        imagePullPolicy: IfNotPresent
        name: clean-cilium-state
        securityContext:
          capabilities:
            add:
            - NET_ADMIN
          privileged: true
        volumeMounts:
        - mountPath: /sys/fs/bpf
          name: bpf-maps
        - mountPath: /var/run/cilium
          name: cilium-run
      priorityClassName: system-node-critical
      restartPolicy: Always
      serviceAccount: cilium
      serviceAccountName: cilium
      terminationGracePeriodSeconds: 1
      tolerations:
      - operator: Exists
      - effect: NoSchedule
        key: node.kubernetes.io/not-ready
        operator: Exists
      volumes:
      - hostPath:
          path: /var/run/cilium
          type: DirectoryOrCreate
        name: cilium-run
      - hostPath:
          path: /sys/fs/bpf
          type: DirectoryOrCreate
        name: bpf-maps
      - hostPath:
          path: /var/run/docker.sock
          type: Socket
        name: docker-socket
      - hostPath:
          path: /opt/cni/bin
          type: DirectoryOrCreate
        name: cni-path
      - hostPath:
          path: /etc/cni/net.d
          type: DirectoryOrCreate
        name: etc-cni-netd
      - configMap:
          defaultMode: 420
          items:
          - key: etcd-config
            path: etcd.config
          name: cilium-config
        name: etcd-config-path
      - name: etcd-secrets
        secret:
          defaultMode: 420
          optional: true
          secretName: cilium-etcd-secrets
      - name: clustermesh-secrets
        secret:
          defaultMode: 420
          optional: true
          secretName: cilium-clustermesh
  updateStrategy:
    rollingUpdate:
      maxUnavailable: 2
    type: RollingUpdate
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    io.cilium/app: operator
    name: cilium-operator
  name: cilium-operator
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  replicas: 1
  selector:
    matchLabels:
      io.cilium/app: operator
      name: cilium-operator
  strategy:
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 1
    type: RollingUpdate
  template:
    metadata:
      labels:
        io.cilium/app: operator
        name: cilium-operator
    spec:
      containers:
      - args:
        - --debug=$(CILIUM_DEBUG)
        - --kvstore=etcd
        - --kvstore-opt=etcd.config=/var/lib/etcd-config/etcd.config
        command:
        - cilium-operator
        env:
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        - name: K8S_NODE_NAME
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: spec.nodeName
        - name: CILIUM_DEBUG
          valueFrom:
            configMapKeyRef:
              key: debug
              name: cilium-config
              optional: true
        - name: CILIUM_CLUSTER_NAME
          valueFrom:
            configMapKeyRef:
              key: cluster-name
              name: cilium-config
              optional: true
        - name: CILIUM_CLUSTER_ID
          valueFrom:
            configMapKeyRef:
              key: cluster-id
              name: cilium-config
              optional: true
        - name: CILIUM_DISABLE_ENDPOINT_CRD
          valueFrom:
            configMapKeyRef:
              key: disable-endpoint-crd
              name: cilium-config
              optional: true
        - name: AWS_ACCESS_KEY_ID
          valueFrom:
            secretKeyRef:
              key: AWS_ACCESS_KEY_ID
              name: cilium-aws
              optional: true
        - name: AWS_SECRET_ACCESS_KEY
          valueFrom:
            secretKeyRef:
              key: AWS_SECRET_ACCESS_KEY
              name: cilium-aws
              optional: true
        - name: AWS_DEFAULT_REGION
          valueFrom:
            secretKeyRef:
              key: AWS_DEFAULT_REGION
              name: cilium-aws
              optional: true
        image: docker.io/cilium/operator:v1.4
        imagePullPolicy: Always
        name: cilium-operator
        volumeMounts:
        - mountPath: /var/lib/etcd-config
          name: etcd-config-path
          readOnly: true
        - mountPath: /var/lib/etcd-secrets
          name: etcd-secrets
          readOnly: true
      dnsPolicy: ClusterFirst
      priorityClassName: system-node-critical
      restartPolicy: Always
      serviceAccount: cilium-operator
      serviceAccountName: cilium-operator
      volumes:
      - configMap:
          defaultMode: 420
          items:
          - key: etcd-config
            path: etcd.config
          name: cilium-config
        name: etcd-config-path
      - name: etcd-secrets
        secret:
          defaultMode: 420
          optional: true
          secretName: cilium-etcd-secrets
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: cilium-operator
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: cilium-operator
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
rules:
- apiGroups:
  - ""
  resources:
  - pods
  - deployments
  - componentstatuses
  verbs:
  - '*'
- apiGroups:
  - ""
  resources:
  - services
  - endpoints
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - cilium.io
  resources:
  - ciliumnetworkpolicies
  - ciliumnetworkpolicies/status
  - ciliumendpoints
  - ciliumendpoints/status
  verbs:
  - '*'
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: cilium-operator
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cilium-operator
subjects:
- kind: ServiceAccount
  name: cilium-operator
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: cilium-etcd-operator
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
rules:
- apiGroups:
  - etcd.database.coreos.com
  resources:
  - etcdclusters
  verbs:
  - get
  - delete
  - create
- apiGroups:
  - apiextensions.k8s.io
  resources:
  - customresourcedefinitions
  verbs:
  - delete
  - get
  - create
- apiGroups:
  - ""
  resources:
  - deployments
  verbs:
  - delete
  - create
  - get
  - update
- apiGroups:
  - ""
  resources:
  - pods
  verbs:
  - list
  - delete
  - get
- apiGroups:
  - apps
  resources:
  - deployments
  verbs:
  - delete
  - create
  - get
  - update
- apiGroups:
  - ""
  resources:
  - componentstatuses
  verbs:
  - get
- apiGroups:
  - extensions
  resources:
  - deployments
  verbs:
  - delete
  - create
  - get
  - update
- apiGroups:
  - ""
  resources:
  - secrets
  verbs:
  - get
  - create
  - delete
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: cilium-etcd-operator
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cilium-etcd-operator
subjects:
- kind: ServiceAccount
  name: cilium-etcd-operator
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: etcd-operator
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
rules:
- apiGroups:
  - etcd.database.coreos.com
  resources:
  - etcdclusters
  - etcdbackups
  - etcdrestores
  verbs:
  - '*'
- apiGroups:
  - apiextensions.k8s.io
  resources:
  - customresourcedefinitions
  verbs:
  - '*'
- apiGroups:
  - ""
  resources:
  - pods
  - services
  - endpoints
  - persistentvolumeclaims
  - events
  - deployments
  verbs:
  - '*'
- apiGroups:
  - apps
  resources:
  - deployments
  verbs:
  - '*'
- apiGroups:
  - extensions
  resources:
  - deployments
  verbs:
  - create
  - get
  - list
  - patch
  - update
- apiGroups:
  - ""
  resources:
  - secrets
  verbs:
  - get
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: etcd-operator
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: etcd-operator
subjects:
- kind: ServiceAccount
  name: cilium-etcd-sa
  namespace: kube-system
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: cilium-etcd-operator
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: cilium-etcd-sa
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    io.cilium/app: etcd-operator
    name: cilium-etcd-operator
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: cilium-etcd-operator
  namespace: kube-system
spec:
  replicas: 1
  selector:
    matchLabels:
      io.cilium/app: etcd-operator
      name: cilium-etcd-operator
  strategy:
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 1
    type: RollingUpdate
  template:
    metadata:
      labels:
        io.cilium/app: etcd-operator
        name: cilium-etcd-operator
    spec:
      containers:
      - command:
        - /usr/bin/cilium-etcd-operator
        env:
        - name: CILIUM_ETCD_OPERATOR_CLUSTER_DOMAIN
          value: cluster.local
        - name: CILIUM_ETCD_OPERATOR_ETCD_CLUSTER_SIZE
          value: "3"
        - name: CILIUM_ETCD_OPERATOR_NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        - name: CILIUM_ETCD_OPERATOR_POD_NAME
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.name
        - name: CILIUM_ETCD_OPERATOR_POD_UID
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.uid
        image: docker.io/cilium/cilium-etcd-operator:v2.0.5
        imagePullPolicy: IfNotPresent
        name: cilium-etcd-operator
      dnsPolicy: ClusterFirst
      hostNetwork: true
      priorityClassName: system-node-critical
      restartPolicy: Always
      serviceAccount: cilium-etcd-operator
      serviceAccountName: cilium-etcd-operator
      tolerations:
      - operator: Exists
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: cilium
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cilium
subjects:
- kind: ServiceAccount
  name: cilium
  namespace: kube-system
- apiGroup: rbac.authorization.k8s.io
  kind: Group
  name: system:nodes
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: cilium
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
rules:
- apiGroups:
  - networking.k8s.io
  resources:
  - networkpolicies
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - namespaces
  - services
  - nodes
  - endpoints
  - componentstatuses
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - pods
  - nodes
  verbs:
  - get
  - list
  - watch
  - update
- apiGroups:
  - extensions
  resources:
  - ingresses
  verbs:
  - create
  - get
  - list
  - watch
- apiGroups:
  - apiextensions.k8s.io
  resources:
  - customresourcedefinitions
  verbs:
  - create
  - get
  - list
  - watch
  - update
- apiGroups:
  - cilium.io
  resources:
  - ciliumnetworkpolicies
  - ciliumnetworkpolicies/status
  - ciliumendpoints
  - ciliumendpoints/status
  verbs:
  - '*'
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: cilium
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
`)

func k8sAddons116KubernetesmasteraddonsCiliumDaemonsetYamlBytes() ([]byte, error) {
	return _k8sAddons116KubernetesmasteraddonsCiliumDaemonsetYaml, nil
}

func k8sAddons116KubernetesmasteraddonsCiliumDaemonsetYaml() (*asset, error) {
	bytes, err := k8sAddons116KubernetesmasteraddonsCiliumDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/1.16/kubernetesmasteraddons-cilium-daemonset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddons116KubernetesmasteraddonsFlannelDaemonsetYaml = []byte(`# This file was pulled from:
# https://github.com/coreos/flannel (HEAD at time of pull was 4973e02e539378)
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: flannel
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
---
kind: ConfigMap
apiVersion: v1
metadata:
  name: kube-flannel-cfg
  namespace: kube-system
  labels:
    tier: node
    app: flannel
    addonmanager.kubernetes.io/mode: EnsureExists
data:
  cni-conf.json: |
    {
      "name": "cbr0",
      "type": "flannel",
      "delegate": {
        "isDefaultGateway": true
      }
    }
  net-conf.json: |
    {
      "Network": "<kubeClusterCidr>",
      "Backend": {
        "Type": "vxlan"
      }
    }
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: kube-flannel-ds
  namespace: kube-system
  labels:
    tier: node
    app: flannel
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  selector:
    matchLabels:
      tier: node
      app: flannel
  template:
    metadata:
      labels:
        tier: node
        app: flannel
    spec:
      hostNetwork: true
      nodeSelector:
        beta.kubernetes.io/arch: amd64
        beta.kubernetes.io/os: linux
      priorityClassName: system-node-critical
      tolerations:
        - key: node.kubernetes.io/not-ready
          operator: Exists
          effect: NoSchedule
        - key: node-role.kubernetes.io/master
          operator: Equal
          value: "true"
          effect: NoSchedule
        - key: CriticalAddonsOnly
          operator: Exists
      serviceAccountName: flannel
      containers:
      - name: kube-flannel
        image: quay.io/coreos/flannel:v0.8.0-amd64
        command: [ "/opt/bin/flanneld", "--ip-masq", "--kube-subnet-mgr" ]
        securityContext:
          privileged: true
        env:
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        volumeMounts:
        - name: run
          mountPath: /run
        - name: flannel-cfg
          mountPath: /etc/kube-flannel/
      - name: install-cni
        image: quay.io/coreos/flannel:v0.10.0-amd64
        command: [ "/bin/sh", "-c", "set -e -x; cp -f /etc/kube-flannel/cni-conf.json /etc/cni/net.d/10-flannel.conf; while true; do sleep 3600; done" ]
        volumeMounts:
        - name: cni
          mountPath: /etc/cni/net.d
        - name: flannel-cfg
          mountPath: /etc/kube-flannel/
      volumes:
        - name: run
          hostPath:
            path: /run
        - name: cni
          hostPath:
            path: /etc/cni/net.d
        - name: flannel-cfg
          configMap:
            name: kube-flannel-cfg
---
# This file was pulled from:
# https://github.com/coreos/flannel (HEAD at time of pull was 4973e02e539378)
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: flannel
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
  - apiGroups:
      - ""
    resources:
      - pods
    verbs:
      - get
  - apiGroups:
      - ""
    resources:
      - nodes
    verbs:
      - list
      - watch
  - apiGroups:
      - ""
    resources:
      - nodes/status
    verbs:
      - patch
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: flannel
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: flannel
subjects:
- kind: ServiceAccount
  name: flannel
  namespace: kube-system
`)

func k8sAddons116KubernetesmasteraddonsFlannelDaemonsetYamlBytes() ([]byte, error) {
	return _k8sAddons116KubernetesmasteraddonsFlannelDaemonsetYaml, nil
}

func k8sAddons116KubernetesmasteraddonsFlannelDaemonsetYaml() (*asset, error) {
	bytes, err := k8sAddons116KubernetesmasteraddonsFlannelDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/1.16/kubernetesmasteraddons-flannel-daemonset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddons116KubernetesmasteraddonsKubeDnsDeploymentYaml = []byte(`# Copyright 2016 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Should keep target in cluster/addons/dns-horizontal-autoscaler/dns-horizontal-autoscaler.yaml
# in sync with this file.

apiVersion: v1
kind: Service
metadata:
  name: kube-dns
  namespace: kube-system
  labels:
    k8s-app: kube-dns
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
    kubernetes.io/name: "KubeDNS"
spec:
  selector:
    k8s-app: kube-dns
  clusterIP: <clustIP>
  ports:
  - name: dns
    port: 53
    protocol: UDP
  - name: dns-tcp
    port: 53
    protocol: TCP
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: kube-dns
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: kube-dns
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: EnsureExists
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kube-dns
  namespace: kube-system
  labels:
    k8s-app: kube-dns
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  # replicas: not specified here:
  # 1. In order to make Addon Manager do not reconcile this replicas parameter.
  # 2. Default is 1.
  # 3. Will be tuned in real time if DNS horizontal auto-scaling is turned on.
  strategy:
    rollingUpdate:
      maxSurge: 10%
      maxUnavailable: 0
  selector:
    matchLabels:
      k8s-app: kube-dns
  template:
    metadata:
      labels:
        k8s-app: kube-dns
    spec:
      priorityClassName: system-node-critical
      tolerations:
      - key: "CriticalAddonsOnly"
        operator: "Exists"
      volumes:
      - name: kube-dns-config
        configMap:
          name: kube-dns
          optional: true
      containers:
      - name: kubedns
        image: <img>
        imagePullPolicy: IfNotPresent
        resources:
          # TODO: Set memory limits when we've profiled the container for large
          # clusters, then set request = limit to keep this container in
          # guaranteed class. Currently, this container falls into the
          # "burstable" category so the kubelet doesn't backoff from restarting it.
          limits:
            memory: 170Mi
          requests:
            cpu: 100m
            memory: 70Mi
        livenessProbe:
          httpGet:
            path: /healthcheck/kubedns
            port: 10054
            scheme: HTTP
          initialDelaySeconds: 60
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 5
        readinessProbe:
          httpGet:
            path: /readiness
            port: 8081
            scheme: HTTP
          initialDelaySeconds: 3
          timeoutSeconds: 5
        args:
        - --domain=<domain>.
        - --dns-port=10053
        - --config-dir=/kube-dns-config
        - --v=2
        env:
        - name: PROMETHEUS_PORT
          value: "10055"
        ports:
        - containerPort: 10053
          name: dns-local
          protocol: UDP
        - containerPort: 10053
          name: dns-tcp-local
          protocol: TCP
        - containerPort: 10055
          name: metrics
          protocol: TCP
        volumeMounts:
        - name: kube-dns-config
          mountPath: /kube-dns-config
      - name: dnsmasq
        image: <imgMasq>
        imagePullPolicy: IfNotPresent
        livenessProbe:
          httpGet:
            path: /healthcheck/dnsmasq
            port: 10054
            scheme: HTTP
          initialDelaySeconds: 60
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 5
        args:
        - -v=2
        - -logtostderr
        - -configDir=/etc/k8s/dns/dnsmasq-nanny
        - -restartDnsmasq=true
        - --
        - -k
        - --cache-size=1000
        - --no-negcache
        - --log-facility=-
        - --server=/cluster.local/127.0.0.1#10053
        - --server=/in-addr.arpa/127.0.0.1#10053
        - --server=/ip6.arpa/127.0.0.1#10053
        ports:
        - containerPort: 53
          name: dns
          protocol: UDP
        - containerPort: 53
          name: dns-tcp
          protocol: TCP
        resources:
          requests:
            cpu: 150m
            memory: 20Mi
        volumeMounts:
        - name: kube-dns-config
          mountPath: /etc/k8s/dns/dnsmasq-nanny
      - name: sidecar
        image: <imgSidecar>
        imagePullPolicy: IfNotPresent
        livenessProbe:
          httpGet:
            path: /metrics
            port: 10054
            scheme: HTTP
          initialDelaySeconds: 60
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 5
        args:
        - --v=2
        - --logtostderr
        - --probe=kubedns,127.0.0.1:10053,kubernetes.default.svc.<domain>,5,SRV
        - --probe=dnsmasq,127.0.0.1:53,kubernetes.default.svc.<domain>,5,SRV
        ports:
        - containerPort: 10054
          name: metrics
          protocol: TCP
        resources:
          requests:
            memory: 20Mi
            cpu: 10m
      dnsPolicy: Default
      serviceAccountName: kube-dns
      nodeSelector:
        beta.kubernetes.io/os: linux
`)

func k8sAddons116KubernetesmasteraddonsKubeDnsDeploymentYamlBytes() ([]byte, error) {
	return _k8sAddons116KubernetesmasteraddonsKubeDnsDeploymentYaml, nil
}

func k8sAddons116KubernetesmasteraddonsKubeDnsDeploymentYaml() (*asset, error) {
	bytes, err := k8sAddons116KubernetesmasteraddonsKubeDnsDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/1.16/kubernetesmasteraddons-kube-dns-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddons116KubernetesmasteraddonsKubeProxyDaemonsetYaml = []byte(`---
apiVersion: v1
kind: ConfigMap
data:
  config.yaml: |
    apiVersion: kubeproxy.config.k8s.io/v1alpha1
    kind: KubeProxyConfiguration
    clientConnection:
      kubeconfig: /var/lib/kubelet/kubeconfig
    clusterCIDR: "<CIDR>"
    mode: "<kubeProxyMode>"
metadata:
  name: kube-proxy-config
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
    kubernetes.io/cluster-service: "true"
    component: kube-proxy
    tier: node
    k8s-app: kube-proxy
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
    kubernetes.io/cluster-service: "true"
    component: kube-proxy
    tier: node
    k8s-app: kube-proxy
  name: kube-proxy
  namespace: kube-system
spec:
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 50%
  selector:
    matchLabels:
      component: kube-proxy
      tier: node
      k8s-app: kube-proxy
  template:
    metadata:
      labels:
        component: kube-proxy
        tier: node
        k8s-app: kube-proxy
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ''
    spec:
      priorityClassName: system-node-critical
      tolerations:
      - key: node-role.kubernetes.io/master
        operator: Equal
        value: "true"
        effect: NoSchedule
      - operator: "Exists"
        effect: NoExecute
      - operator: "Exists"
        effect: NoSchedule
      - key: CriticalAddonsOnly
        operator: Exists
      containers:
      - command:
        - /hyperkube
        - kube-proxy
        - --config=/var/lib/kube-proxy/config.yaml
        image: <img>
        imagePullPolicy: IfNotPresent
        name: kube-proxy
        resources:
          requests:
            cpu: 100m
        securityContext:
          privileged: true
        volumeMounts:
        - mountPath: /etc/ssl/certs
          name: ssl-certs-host
          readOnly: true
        - mountPath: /etc/kubernetes
          name: etc-kubernetes
          readOnly: true
        - mountPath: /var/lib/kubelet/kubeconfig
          name: kubeconfig
          readOnly: true
        - mountPath: /run/xtables.lock
          name: iptableslock
        - mountPath: /lib/modules/
          name: kernelmodules
          readOnly: true
        - mountPath: /var/lib/kube-proxy/config.yaml
          subPath: config.yaml
          name: kube-proxy-config-volume
          readOnly: true
      hostNetwork: true
      volumes:
      - hostPath:
          path: /usr/share/ca-certificates
        name: ssl-certs-host
      - hostPath:
          path: /var/lib/kubelet/kubeconfig
        name: kubeconfig
      - hostPath:
          path: /etc/kubernetes
        name: etc-kubernetes
      - hostPath:
          path: /run/xtables.lock
        name: iptableslock
      - hostPath:
          path: /lib/modules/
        name: kernelmodules
      - configMap:
          name: kube-proxy-config
        name: kube-proxy-config-volume
      nodeSelector:
        beta.kubernetes.io/os: linux
`)

func k8sAddons116KubernetesmasteraddonsKubeProxyDaemonsetYamlBytes() ([]byte, error) {
	return _k8sAddons116KubernetesmasteraddonsKubeProxyDaemonsetYaml, nil
}

func k8sAddons116KubernetesmasteraddonsKubeProxyDaemonsetYaml() (*asset, error) {
	bytes, err := k8sAddons116KubernetesmasteraddonsKubeProxyDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/1.16/kubernetesmasteraddons-kube-proxy-daemonset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddons116KubernetesmasteraddonsPodSecurityPolicyYaml = []byte(`apiVersion: policy/v1beta1
kind: PodSecurityPolicy
metadata:
  name: privileged
  annotations:
    seccomp.security.alpha.kubernetes.io/allowedProfileNames: "*"
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  privileged: true
  allowPrivilegeEscalation: true
  allowedCapabilities:
  - "*"
  volumes:
  - "*"
  hostNetwork: true
  hostPorts:
  - min: 0
    max: 65535
  hostIPC: true
  hostPID: true
  runAsUser:
    rule: RunAsAny
  seLinux:
    rule: RunAsAny
  supplementalGroups:
    rule: RunAsAny
  fsGroup:
    rule: RunAsAny
---
apiVersion: policy/v1beta1
kind: PodSecurityPolicy
metadata:
  name: restricted
  annotations:
    seccomp.security.alpha.kubernetes.io/allowedProfileNames: docker/default
    apparmor.security.beta.kubernetes.io/allowedProfileNames: runtime/default
    seccomp.security.alpha.kubernetes.io/defaultProfileName:  docker/default
    apparmor.security.beta.kubernetes.io/defaultProfileName:  runtime/default
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  privileged: false
  allowPrivilegeEscalation: false
  requiredDropCapabilities:
    - ALL
  volumes:
    - configMap
    - emptyDir
    - projected
    - secret
    - downwardAPI
    - persistentVolumeClaim
  hostNetwork: false
  hostIPC: false
  hostPID: false
  runAsUser:
    rule: MustRunAsNonRoot
  seLinux:
    rule: RunAsAny
  supplementalGroups:
    rule: MustRunAs
    ranges:
      # Forbid adding the root group.
      - min: 1
        max: 65535
  fsGroup:
    rule: MustRunAs
    ranges:
      # Forbid adding the root group.
      - min: 1
        max: 65535
  readOnlyRootFilesystem: false
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: psp:privileged
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups: ['extensions']
  resources: ['podsecuritypolicies']
  verbs:     ['use']
  resourceNames:
  - privileged
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: psp:restricted
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups: ['extensions']
  resources: ['podsecuritypolicies']
  verbs:     ['use']
  resourceNames:
  - restricted
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: default:restricted
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: psp:restricted
subjects:
- kind: Group
  name: system:authenticated
  apiGroup: rbac.authorization.k8s.io
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: default:privileged
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: psp:privileged
subjects:
- kind: Group
  name: system:authenticated
  apiGroup: rbac.authorization.k8s.io
- kind: Group
  name: system:nodes
  apiGroup: rbac.authorization.k8s.io
`)

func k8sAddons116KubernetesmasteraddonsPodSecurityPolicyYamlBytes() ([]byte, error) {
	return _k8sAddons116KubernetesmasteraddonsPodSecurityPolicyYaml, nil
}

func k8sAddons116KubernetesmasteraddonsPodSecurityPolicyYaml() (*asset, error) {
	bytes, err := k8sAddons116KubernetesmasteraddonsPodSecurityPolicyYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/1.16/kubernetesmasteraddons-pod-security-policy.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddons16KubernetesmasteraddonsKubernetesDashboardDeploymentYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    k8s-app: kubernetes-dashboard
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: kubernetes-dashboard
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: kubernetes-dashboard
  labels:
    k8s-app: kubernetes-dashboard
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: kubernetes-dashboard
  namespace: kube-system
---
apiVersion: v1
kind: Service
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kube-system
spec:
  ports:
  - port: 80
    targetPort: 9090
  selector:
    k8s-app: kubernetes-dashboard
  type: NodePort
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kube-system
spec:
  replicas: 1
  selector:
    matchLabels:
      k8s-app: kubernetes-dashboard
  template:
    metadata:
      labels:
        k8s-app: kubernetes-dashboard
    spec:
      containers:
      - args:
        - --heapster-host=http://heapster.kube-system:80
        image: <img>
        imagePullPolicy: IfNotPresent
        livenessProbe:
          httpGet:
            path: "/"
            port: 9090
          initialDelaySeconds: 30
          timeoutSeconds: 30
        name: kubernetes-dashboard
        ports:
        - containerPort: 9090
          protocol: TCP
        resources:
          requests:
            cpu: <cpuReq>
            memory: <memReq>
          limits:
            cpu: <cpuLim>
            memory: <memLim>
      serviceAccountName: kubernetes-dashboard
      nodeSelector:
        beta.kubernetes.io/os: linux
`)

func k8sAddons16KubernetesmasteraddonsKubernetesDashboardDeploymentYamlBytes() ([]byte, error) {
	return _k8sAddons16KubernetesmasteraddonsKubernetesDashboardDeploymentYaml, nil
}

func k8sAddons16KubernetesmasteraddonsKubernetesDashboardDeploymentYaml() (*asset, error) {
	bytes, err := k8sAddons16KubernetesmasteraddonsKubernetesDashboardDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/1.6/kubernetesmasteraddons-kubernetes-dashboard-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddons17KubernetesmasteraddonsKubeDnsDeploymentYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: kube-dns
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: v1
kind: Service
metadata:
  labels:
    k8s-app: kube-dns
    kubernetes.io/cluster-service: "true"
    kubernetes.io/name: KubeDNS
  name: kube-dns
  namespace: kube-system
spec:
  clusterIP: <clustIP>
  ports:
  - name: dns
    port: 53
    protocol: UDP
  - name: dns-tcp
    port: 53
    protocol: TCP
  selector:
    k8s-app: kube-dns
---
apiVersion: apps/v1beta1
kind: Deployment
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
    k8s-app: kube-dns
    version: v20
  name: kube-dns-v20
  namespace: kube-system
spec:
  replicas: 2
  selector:
    matchLabels:
      k8s-app: kube-dns
      version: v20
  template:
    metadata:
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ""
        prometheus.io/scrape: "true"
        prometheus.io/port: "10055"
      labels:
        k8s-app: kube-dns
        kubernetes.io/cluster-service: "true"
        version: v20
    spec:
      affinity:
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - weight: 100
            podAffinityTerm:
              labelSelector:
                matchExpressions:
                - key: k8s-app
                  operator: In
                  values:
                  - kube-dns
              topologyKey: kubernetes.io/hostname
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      volumes:
      - name: kube-dns-config
        configMap:
          name: kube-dns
          optional: true
      containers:
      - args:
        - "--domain=<domain>."
        - "--dns-port=10053"
        - "--v=2"
        - "--config-dir=/kube-dns-config"
        env:
        - name: PROMETHEUS_PORT
          value: "10055"
        image: <img>
        imagePullPolicy: IfNotPresent
        livenessProbe:
          failureThreshold: 5
          httpGet:
            path: "/healthz-kubedns"
            port: 8080
            scheme: HTTP
          initialDelaySeconds: 60
          successThreshold: 1
          timeoutSeconds: 5
        name: kubedns
        ports:
        - containerPort: 10053
          name: dns-local
          protocol: UDP
        - containerPort: 10053
          name: dns-tcp-local
          protocol: TCP
        - containerPort: 10055
          name: metrics
          protocol: TCP
        readinessProbe:
          httpGet:
            path: "/readiness"
            port: 8081
            scheme: HTTP
          initialDelaySeconds: 30
          timeoutSeconds: 5
        resources:
          limits:
            memory: 170Mi
          requests:
            cpu: 100m
            memory: 70Mi
        volumeMounts:
        - name: kube-dns-config
          mountPath: /kube-dns-config
      - args:
        - "-v=2"
        - "-logtostderr"
        - "-configDir=/kube-dns-config"
        - "-restartDnsmasq=true"
        - "--"
        - "-k"
        - "--cache-size=1000"
        - "--no-resolv"
        - "--server=127.0.0.1#10053"
        - "--server=/in-addr.arpa/127.0.0.1#10053"
        - "--server=/ip6.arpa/127.0.0.1#10053"
        - "--log-facility=-"
        image: <imgMasq>
        imagePullPolicy: IfNotPresent
        name: dnsmasq
        ports:
        - containerPort: 53
          name: dns
          protocol: UDP
        - containerPort: 53
          name: dns-tcp
          protocol: TCP
        volumeMounts:
        - mountPath: /kube-dns-config
          name: kube-dns-config
      - args:
        - "--cmd=for d in $PROBE_DOMAINS; do nslookup $d 127.0.0.1 >/dev/null || exit 1; done"
        - "--url=/healthz-dnsmasq"
        - "--cmd=for d in $PROBE_DOMAINS; do nslookup $d 127.0.0.1:10053 >/dev/null || exit 1; done"
        - "--url=/healthz-kubedns"
        - "--port=8080"
        - "--quiet"
        env:
        - name: PROBE_DOMAINS
          value: bing.com kubernetes.default.svc.<domain>
        image: <imgHealthz>
        imagePullPolicy: IfNotPresent
        livenessProbe:
          failureThreshold: 5
          httpGet:
            path: "/healthz-dnsmasq"
            port: 8080
            scheme: HTTP
          initialDelaySeconds: 60
          successThreshold: 1
          timeoutSeconds: 5
        name: healthz
        ports:
        - containerPort: 8080
          protocol: TCP
        resources:
          limits:
            memory: 50Mi
          requests:
            cpu: 10m
            memory: 50Mi
      dnsPolicy: Default
      serviceAccountName: kube-dns
      nodeSelector:
        beta.kubernetes.io/os: linux
`)

func k8sAddons17KubernetesmasteraddonsKubeDnsDeploymentYamlBytes() ([]byte, error) {
	return _k8sAddons17KubernetesmasteraddonsKubeDnsDeploymentYaml, nil
}

func k8sAddons17KubernetesmasteraddonsKubeDnsDeploymentYaml() (*asset, error) {
	bytes, err := k8sAddons17KubernetesmasteraddonsKubeDnsDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/1.7/kubernetesmasteraddons-kube-dns-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddons17KubernetesmasteraddonsKubernetesDashboardDeploymentYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    k8s-app: kubernetes-dashboard
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: kubernetes-dashboard
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: kubernetes-dashboard
  labels:
    k8s-app: kubernetes-dashboard
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: kubernetes-dashboard
  namespace: kube-system
---
apiVersion: v1
kind: Service
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kube-system
spec:
  ports:
  - port: 80
    targetPort: 9090
  selector:
    k8s-app: kubernetes-dashboard
  type: NodePort
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kube-system
spec:
  replicas: 1
  selector:
    matchLabels:
      k8s-app: kubernetes-dashboard
  template:
    metadata:
      labels:
        k8s-app: kubernetes-dashboard
    spec:
      containers:
      - args:
        - --heapster-host=http://heapster.kube-system:80
        image: <img>
        imagePullPolicy: IfNotPresent
        livenessProbe:
          httpGet:
            path: "/"
            port: 9090
          initialDelaySeconds: 30
          timeoutSeconds: 30
        name: kubernetes-dashboard
        ports:
        - containerPort: 9090
          protocol: TCP
        resources:
          requests:
            cpu: <cpuReq>
            memory: <memReq>
          limits:
            cpu: <cpuLim>
            memory: <memLim>
      serviceAccountName: kubernetes-dashboard
      nodeSelector:
        beta.kubernetes.io/os: linux
`)

func k8sAddons17KubernetesmasteraddonsKubernetesDashboardDeploymentYamlBytes() ([]byte, error) {
	return _k8sAddons17KubernetesmasteraddonsKubernetesDashboardDeploymentYaml, nil
}

func k8sAddons17KubernetesmasteraddonsKubernetesDashboardDeploymentYaml() (*asset, error) {
	bytes, err := k8sAddons17KubernetesmasteraddonsKubernetesDashboardDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/1.7/kubernetesmasteraddons-kubernetes-dashboard-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddons18KubernetesmasteraddonsKubeDnsDeploymentYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: kube-dns
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: v1
kind: Service
metadata:
  labels:
    k8s-app: kube-dns
    kubernetes.io/cluster-service: "true"
    kubernetes.io/name: KubeDNS
  name: kube-dns
  namespace: kube-system
spec:
  clusterIP: <clustIP>
  ports:
  - name: dns
    port: 53
    protocol: UDP
  - name: dns-tcp
    port: 53
    protocol: TCP
  selector:
    k8s-app: kube-dns
---
apiVersion: apps/v1beta1
kind: Deployment
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
    k8s-app: kube-dns
    version: v20
  name: kube-dns-v20
  namespace: kube-system
spec:
  replicas: 2
  selector:
    matchLabels:
      k8s-app: kube-dns
      version: v20
  template:
    metadata:
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ""
        prometheus.io/scrape: "true"
        prometheus.io/port: "10055"
      labels:
        k8s-app: kube-dns
        kubernetes.io/cluster-service: "true"
        version: v20
    spec:
      affinity:
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - weight: 100
            podAffinityTerm:
              labelSelector:
                matchExpressions:
                - key: k8s-app
                  operator: In
                  values:
                  - kube-dns
              topologyKey: kubernetes.io/hostname
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      volumes:
      - name: kube-dns-config
        configMap:
          name: kube-dns
          optional: true
      containers:
      - args:
        - "--domain=<domain>."
        - "--dns-port=10053"
        - "--v=2"
        - "--config-dir=/kube-dns-config"
        env:
        - name: PROMETHEUS_PORT
          value: "10055"
        image: <img>
        imagePullPolicy: IfNotPresent
        livenessProbe:
          failureThreshold: 5
          httpGet:
            path: "/healthz-kubedns"
            port: 8080
            scheme: HTTP
          initialDelaySeconds: 60
          successThreshold: 1
          timeoutSeconds: 5
        name: kubedns
        ports:
        - containerPort: 10053
          name: dns-local
          protocol: UDP
        - containerPort: 10053
          name: dns-tcp-local
          protocol: TCP
        - containerPort: 10055
          name: metrics
          protocol: TCP
        readinessProbe:
          httpGet:
            path: "/readiness"
            port: 8081
            scheme: HTTP
          initialDelaySeconds: 30
          timeoutSeconds: 5
        resources:
          limits:
            memory: 170Mi
          requests:
            cpu: 100m
            memory: 70Mi
        volumeMounts:
        - name: kube-dns-config
          mountPath: /kube-dns-config
      - args:
        - "-v=2"
        - "-logtostderr"
        - "-configDir=/kube-dns-config"
        - "-restartDnsmasq=true"
        - "--"
        - "-k"
        - "--cache-size=1000"
        - "--no-resolv"
        - "--server=127.0.0.1#10053"
        - "--server=/in-addr.arpa/127.0.0.1#10053"
        - "--server=/ip6.arpa/127.0.0.1#10053"
        - "--log-facility=-"
        image: <imgMasq>
        imagePullPolicy: IfNotPresent
        name: dnsmasq
        ports:
        - containerPort: 53
          name: dns
          protocol: UDP
        - containerPort: 53
          name: dns-tcp
          protocol: TCP
        volumeMounts:
        - mountPath: /kube-dns-config
          name: kube-dns-config
      - args:
        - "--cmd=for d in $PROBE_DOMAINS; do nslookup $d 127.0.0.1 >/dev/null || exit 1; done"
        - "--url=/healthz-dnsmasq"
        - "--cmd=for d in $PROBE_DOMAINS; do nslookup $d 127.0.0.1:10053 >/dev/null || exit 1; done"
        - "--url=/healthz-kubedns"
        - "--port=8080"
        - "--quiet"
        env:
        - name: PROBE_DOMAINS
          value: bing.com kubernetes.default.svc.<domain>
        image: <imgHealthz>
        imagePullPolicy: IfNotPresent
        livenessProbe:
          failureThreshold: 5
          httpGet:
            path: "/healthz-dnsmasq"
            port: 8080
            scheme: HTTP
          initialDelaySeconds: 60
          successThreshold: 1
          timeoutSeconds: 5
        name: healthz
        ports:
        - containerPort: 8080
          protocol: TCP
        resources:
          limits:
            memory: 50Mi
          requests:
            cpu: 10m
            memory: 50Mi
      dnsPolicy: Default
      serviceAccountName: kube-dns
      nodeSelector:
        beta.kubernetes.io/os: linux
`)

func k8sAddons18KubernetesmasteraddonsKubeDnsDeploymentYamlBytes() ([]byte, error) {
	return _k8sAddons18KubernetesmasteraddonsKubeDnsDeploymentYaml, nil
}

func k8sAddons18KubernetesmasteraddonsKubeDnsDeploymentYaml() (*asset, error) {
	bytes, err := k8sAddons18KubernetesmasteraddonsKubeDnsDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/1.8/kubernetesmasteraddons-kube-dns-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddons18KubernetesmasteraddonsKubernetesDashboardDeploymentYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    k8s-app: kubernetes-dashboard
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: kubernetes-dashboard
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: kubernetes-dashboard
  labels:
    k8s-app: kubernetes-dashboard
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: kubernetes-dashboard
  namespace: kube-system
---
apiVersion: v1
kind: Service
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kube-system
spec:
  ports:
  - port: 80
    targetPort: 9090
  selector:
    k8s-app: kubernetes-dashboard
  type: NodePort
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kube-system
spec:
  replicas: 1
  selector:
    matchLabels:
      k8s-app: kubernetes-dashboard
  template:
    metadata:
      labels:
        k8s-app: kubernetes-dashboard
    spec:
      containers:
      - args:
        - --heapster-host=http://heapster.kube-system:80
        image: <img>
        imagePullPolicy: IfNotPresent
        livenessProbe:
          httpGet:
            path: "/"
            port: 9090
          initialDelaySeconds: 30
          timeoutSeconds: 30
        name: kubernetes-dashboard
        ports:
        - containerPort: 9090
          protocol: TCP
        resources:
          requests:
            cpu: <cpuReq>
            memory: <memReq>
          limits:
            cpu: <cpuLim>
            memory: <memLim>
      serviceAccountName: kubernetes-dashboard
      nodeSelector:
        beta.kubernetes.io/os: linux
`)

func k8sAddons18KubernetesmasteraddonsKubernetesDashboardDeploymentYamlBytes() ([]byte, error) {
	return _k8sAddons18KubernetesmasteraddonsKubernetesDashboardDeploymentYaml, nil
}

func k8sAddons18KubernetesmasteraddonsKubernetesDashboardDeploymentYaml() (*asset, error) {
	bytes, err := k8sAddons18KubernetesmasteraddonsKubernetesDashboardDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/1.8/kubernetesmasteraddons-kubernetes-dashboard-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddons19KubernetesmasteraddonsKubeDnsDeploymentYaml = []byte(`# Copyright 2016 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Should keep target in cluster/addons/dns-horizontal-autoscaler/dns-horizontal-autoscaler.yaml
# in sync with this file.

apiVersion: v1
kind: Service
metadata:
  name: kube-dns
  namespace: kube-system
  labels:
    k8s-app: kube-dns
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
    kubernetes.io/name: "KubeDNS"
spec:
  selector:
    k8s-app: kube-dns
  clusterIP: <clustIP>
  ports:
  - name: dns
    port: 53
    protocol: UDP
  - name: dns-tcp
    port: 53
    protocol: TCP
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: kube-dns
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: kube-dns
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: EnsureExists
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: kube-dns
  namespace: kube-system
  labels:
    k8s-app: kube-dns
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  # replicas: not specified here:
  # 1. In order to make Addon Manager do not reconcile this replicas parameter.
  # 2. Default is 1.
  # 3. Will be tuned in real time if DNS horizontal auto-scaling is turned on.
  strategy:
    rollingUpdate:
      maxSurge: 10%
      maxUnavailable: 0
  selector:
    matchLabels:
      k8s-app: kube-dns
  template:
    metadata:
      labels:
        k8s-app: kube-dns
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ''
    spec:
      priorityClassName: system-node-critical
      tolerations:
      - key: "CriticalAddonsOnly"
        operator: "Exists"
      volumes:
      - name: kube-dns-config
        configMap:
          name: kube-dns
          optional: true
      containers:
      - name: kubedns
        image: <img>
        imagePullPolicy: IfNotPresent
        resources:
          # TODO: Set memory limits when we've profiled the container for large
          # clusters, then set request = limit to keep this container in
          # guaranteed class. Currently, this container falls into the
          # "burstable" category so the kubelet doesn't backoff from restarting it.
          limits:
            memory: 170Mi
          requests:
            cpu: 100m
            memory: 70Mi
        livenessProbe:
          httpGet:
            path: /healthcheck/kubedns
            port: 10054
            scheme: HTTP
          initialDelaySeconds: 60
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 5
        readinessProbe:
          httpGet:
            path: /readiness
            port: 8081
            scheme: HTTP
          initialDelaySeconds: 3
          timeoutSeconds: 5
        args:
        - --domain=<domain>.
        - --dns-port=10053
        - --config-dir=/kube-dns-config
        - --v=2
        env:
        - name: PROMETHEUS_PORT
          value: "10055"
        ports:
        - containerPort: 10053
          name: dns-local
          protocol: UDP
        - containerPort: 10053
          name: dns-tcp-local
          protocol: TCP
        - containerPort: 10055
          name: metrics
          protocol: TCP
        volumeMounts:
        - name: kube-dns-config
          mountPath: /kube-dns-config
      - name: dnsmasq
        image: <imgMasq>
        imagePullPolicy: IfNotPresent
        livenessProbe:
          httpGet:
            path: /healthcheck/dnsmasq
            port: 10054
            scheme: HTTP
          initialDelaySeconds: 60
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 5
        args:
        - -v=2
        - -logtostderr
        - -configDir=/etc/k8s/dns/dnsmasq-nanny
        - -restartDnsmasq=true
        - --
        - -k
        - --cache-size=1000
        - --no-negcache
        - --log-facility=-
        - --server=/cluster.local/127.0.0.1#10053
        - --server=/in-addr.arpa/127.0.0.1#10053
        - --server=/ip6.arpa/127.0.0.1#10053
        ports:
        - containerPort: 53
          name: dns
          protocol: UDP
        - containerPort: 53
          name: dns-tcp
          protocol: TCP
        resources:
          requests:
            cpu: 150m
            memory: 20Mi
        volumeMounts:
        - name: kube-dns-config
          mountPath: /etc/k8s/dns/dnsmasq-nanny
      - name: sidecar
        image: <imgSidecar>
        imagePullPolicy: IfNotPresent
        livenessProbe:
          httpGet:
            path: /metrics
            port: 10054
            scheme: HTTP
          initialDelaySeconds: 60
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 5
        args:
        - --v=2
        - --logtostderr
        - --probe=kubedns,127.0.0.1:10053,kubernetes.default.svc.<domain>,5,SRV
        - --probe=dnsmasq,127.0.0.1:53,kubernetes.default.svc.<domain>,5,SRV
        ports:
        - containerPort: 10054
          name: metrics
          protocol: TCP
        resources:
          requests:
            memory: 20Mi
            cpu: 10m
      dnsPolicy: Default
      serviceAccountName: kube-dns
      nodeSelector:
        beta.kubernetes.io/os: linux`)

func k8sAddons19KubernetesmasteraddonsKubeDnsDeploymentYamlBytes() ([]byte, error) {
	return _k8sAddons19KubernetesmasteraddonsKubeDnsDeploymentYaml, nil
}

func k8sAddons19KubernetesmasteraddonsKubeDnsDeploymentYaml() (*asset, error) {
	bytes, err := k8sAddons19KubernetesmasteraddonsKubeDnsDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/1.9/kubernetesmasteraddons-kube-dns-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsCorednsYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: coredns
  namespace: kube-system
  labels:
      kubernetes.io/cluster-service: "true"
      addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    kubernetes.io/bootstrapping: rbac-defaults
    addonmanager.kubernetes.io/mode: Reconcile
  name: system:coredns
rules:
- apiGroups:
  - ""
  resources:
  - endpoints
  - services
  - pods
  - namespaces
  verbs:
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - nodes
  verbs:
  - get
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  annotations:
    rbac.authorization.kubernetes.io/autoupdate: "true"
  labels:
    kubernetes.io/bootstrapping: rbac-defaults
    addonmanager.kubernetes.io/mode: EnsureExists
  name: system:coredns
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:coredns
subjects:
- kind: ServiceAccount
  name: coredns
  namespace: kube-system
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: coredns
  namespace: kube-system
  labels:
      addonmanager.kubernetes.io/mode: Reconcile
data:
  Corefile: |
    import conf.d/Corefile*
    .:53 {
        errors
        health
        ready
        kubernetes <domain> in-addr.arpa ip6.arpa {
            pods insecure
            fallthrough in-addr.arpa ip6.arpa
        }
        prometheus :9153
        forward . /etc/resolv.conf
        cache 30
        loop
        reload
        loadbalance
    }
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: coredns-custom
  namespace: kube-system
  labels:
      addonmanager.kubernetes.io/mode: EnsureExists
data:
  Corefile: |
    # Add custom CoreDNS configuration here.
    #
    # See https://github.com/coredns/coredns/tree/master/plugin/azure for information
    # about the Azure DNS plugin.
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: coredns
  namespace: kube-system
  labels:
    k8s-app: kube-dns
    kubernetes.io/name: "CoreDNS"
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  # replicas: not specified here:
  # 1. In order to make Addon Manager do not reconcile this replicas parameter.
  # 2. Default is 1.
  # 3. Will be tuned in real time if DNS horizontal auto-scaling is turned on.
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 1
  selector:
    matchLabels:
      k8s-app: kube-dns
  template:
    metadata:
      labels:
        k8s-app: kube-dns
      annotations:
        seccomp.security.alpha.kubernetes.io/pod: docker/default
    spec:
      priorityClassName: system-cluster-critical
      affinity:
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - podAffinityTerm:
              labelSelector:
                matchExpressions:
                - key: k8s-app
                  operator: In
                  values:
                  - kube-dns
              topologyKey: failure-domain.beta.kubernetes.io/zone
            weight: 10
          - podAffinityTerm:
              labelSelector:
                matchExpressions:
                - key: k8s-app
                  operator: In
                  values:
                  - kube-dns
              topologyKey: kubernetes.io/hostname
            weight: 5
      serviceAccountName: coredns
      tolerations:
        - key: node-role.kubernetes.io/master
          effect: NoSchedule
        - key: CriticalAddonsOnly
          operator: "Exists"
        - operator: "Exists"
          effect: NoExecute
        - operator: "Exists"
          effect: NoSchedule
      nodeSelector:
        beta.kubernetes.io/os: linux
      containers:
      - name: coredns
        image: <img>
        imagePullPolicy: IfNotPresent
        resources:
          limits:
            memory: 170Mi
          requests:
            cpu: 100m
            memory: 70Mi
        args: [ "-conf", "/etc/coredns/Corefile" ]
        volumeMounts:
        - name: config-volume
          mountPath: /etc/coredns
          readOnly: true
        - mountPath: /etc/coredns/conf.d
          name: config-custom
          readOnly: true
        ports:
        - containerPort: 53
          name: dns
          protocol: UDP
        - containerPort: 53
          name: dns-tcp
          protocol: TCP
        - containerPort: 9153
          name: metrics
          protocol: TCP
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
            scheme: HTTP
          initialDelaySeconds: 60
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 5
        readinessProbe:
          httpGet:
            path: /ready
            port: 8181
            scheme: HTTP
        securityContext:
          allowPrivilegeEscalation: false
          capabilities:
            add:
            - NET_BIND_SERVICE
            drop:
            - all
          readOnlyRootFilesystem: true
      dnsPolicy: Default
      volumes:
        - name: config-volume
          configMap:
            name: coredns
            items:
            - key: Corefile
              path: Corefile
        - name: config-custom
          configMap:
            name: coredns-custom
            items:
            - key: Corefile
              path: Corefile
            optional: true

---
apiVersion: v1
kind: Service
metadata:
  name: kube-dns
  namespace: kube-system
  annotations:
    prometheus.io/port: "9153"
    prometheus.io/scrape: "true"
  labels:
    k8s-app: kube-dns
    kubernetes.io/cluster-service: "true"
    kubernetes.io/name: CoreDNS
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  selector:
    k8s-app: kube-dns
  clusterIP: <clustIP>
  ports:
  - name: dns
    port: 53
    protocol: UDP
  - name: dns-tcp
    port: 53
    protocol: TCP
  - name: metrics
    port: 9153
    protocol: TCP
`)

func k8sAddonsCorednsYamlBytes() ([]byte, error) {
	return _k8sAddonsCorednsYaml, nil
}

func k8sAddonsCorednsYaml() (*asset, error) {
	bytes, err := k8sAddonsCorednsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/coredns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsKubernetesmasterAuditPolicyYaml = []byte(`apiVersion: audit.k8s.io/v1beta1 # This is required.
kind: Policy
omitStages:
  - RequestReceived
rules:
  - level: RequestResponse
    resources:
    - group: ""
      resources: ["pods"]
  - level: Metadata
    resources:
    - group: ""
      resources: ["pods/log", "pods/status"]
  - level: None
    users: ["system:kube-proxy"]
    verbs: ["watch"]
    resources:
    - group: ""
      resources: ["endpoints", "services"]
  - level: None
    userGroups: ["system:authenticated"]
    nonResourceURLs:
    - /api*
    - /version
  - level: Request
    resources:
    - group: ""
      resources: ["configmaps"]
    namespaces: ["kube-system"]
  - level: Request
    resources:
    - group: ""
      resources: ["secrets"]
  - level: Request
    resources:
    - group: ""
    - group: extensions
  - level: Metadata
    omitStages:
      - RequestReceived
`)

func k8sAddonsKubernetesmasterAuditPolicyYamlBytes() ([]byte, error) {
	return _k8sAddonsKubernetesmasterAuditPolicyYaml, nil
}

func k8sAddonsKubernetesmasterAuditPolicyYaml() (*asset, error) {
	bytes, err := k8sAddonsKubernetesmasterAuditPolicyYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/kubernetesmaster-audit-policy.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsKubernetesmasteraddonsAadDefaultAdminGroupRbacYaml = []byte(`kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: aad-default-admin-group
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
subjects:
- kind: Group
  name: <gID>
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: ClusterRole
  name: cluster-admin
  apiGroup: rbac.authorization.k8s.io
`)

func k8sAddonsKubernetesmasteraddonsAadDefaultAdminGroupRbacYamlBytes() ([]byte, error) {
	return _k8sAddonsKubernetesmasteraddonsAadDefaultAdminGroupRbacYaml, nil
}

func k8sAddonsKubernetesmasteraddonsAadDefaultAdminGroupRbacYaml() (*asset, error) {
	bytes, err := k8sAddonsKubernetesmasteraddonsAadDefaultAdminGroupRbacYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/kubernetesmasteraddons-aad-default-admin-group-rbac.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsKubernetesmasteraddonsAzureCloudProviderDeploymentYaml = []byte(`---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
  name: system:azure-cloud-provider
rules:
- apiGroups: [""]
  resources: ["events"]
  verbs:
  - create
  - patch
  - update
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
  name: system:azure-cloud-provider
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:azure-cloud-provider
subjects:
- kind: ServiceAccount
  name: azure-cloud-provider
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: system:azure-persistent-volume-binder
  labels:
    kubernetes.io/cluster-service: "true"
rules:
- apiGroups: ['']
  resources: ['secrets']
  verbs:     ['get','create']
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: system:azure-persistent-volume-binder
  labels:
    kubernetes.io/cluster-service: "true"
roleRef:
  kind: ClusterRole
  apiGroup: rbac.authorization.k8s.io
  name: system:azure-persistent-volume-binder
subjects:
- kind: ServiceAccount
  name: persistent-volume-binder
  namespace: kube-system
`)

func k8sAddonsKubernetesmasteraddonsAzureCloudProviderDeploymentYamlBytes() ([]byte, error) {
	return _k8sAddonsKubernetesmasteraddonsAzureCloudProviderDeploymentYaml, nil
}

func k8sAddonsKubernetesmasteraddonsAzureCloudProviderDeploymentYaml() (*asset, error) {
	bytes, err := k8sAddonsKubernetesmasteraddonsAzureCloudProviderDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/kubernetesmasteraddons-azure-cloud-provider-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsKubernetesmasteraddonsCiliumDaemonsetYaml = []byte(`---
apiVersion: v1
kind: ConfigMap
metadata:
  name: cilium-config
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "EnsureExists"
data:
  etcd-config: |-
    ---
    endpoints:
      - https://cilium-etcd-client.kube-system.svc:2379
    ca-file: '/var/lib/etcd-secrets/etcd-client-ca.crt'
    key-file: '/var/lib/etcd-secrets/etcd-client.key'
    cert-file: '/var/lib/etcd-secrets/etcd-client.crt'
  debug: "false"
  enable-ipv4: "true"
  enable-ipv6: "false"
  clean-cilium-state: "false"
  clean-cilium-bpf-state: "false"
  monitor-aggregation-level: "none"
  ct-global-max-entries-tcp: "524288"
  ct-global-max-entries-other: "262144"
  preallocate-bpf-maps: "false"
  sidecar-istio-proxy-image: "cilium/istio_proxy"
  tunnel: "vxlan"
  cluster-name: default
  flannel-master-device: ""
  flannel-uninstall-on-exit: "false"
  flannel-manage-existing-containers: "false"
  tofqdns-enable-poller: "false"
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  labels:
    k8s-app: cilium
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: cilium
  namespace: kube-system
spec:
  selector:
    matchLabels:
      k8s-app: cilium
      kubernetes.io/cluster-service: "true"
  template:
    metadata:
      annotations:
        prometheus.io/port: "9090"
        prometheus.io/scrape: "true"
        scheduler.alpha.kubernetes.io/critical-pod: ""
        scheduler.alpha.kubernetes.io/tolerations: '[{"key":"dedicated","operator":"Equal","value":"master","effect":"NoSchedule"}]'
      labels:
        k8s-app: cilium
        kubernetes.io/cluster-service: "true"
    spec:
      containers:
      - args:
        - --debug=$(CILIUM_DEBUG)
        - --kvstore=etcd
        - --kvstore-opt=etcd.config=/var/lib/etcd-config/etcd.config
        command:
        - cilium-agent
        env:
        - name: K8S_NODE_NAME
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: spec.nodeName
        - name: CILIUM_K8S_NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        - name: CILIUM_DEBUG
          valueFrom:
            configMapKeyRef:
              key: debug
              name: cilium-config
        - name: CILIUM_ENABLE_IPV4
          valueFrom:
            configMapKeyRef:
              key: enable-ipv4
              name: cilium-config
              optional: true
        - name: CILIUM_ENABLE_IPV6
          valueFrom:
            configMapKeyRef:
              key: enable-ipv6
              name: cilium-config
              optional: true
        - name: CILIUM_PROMETHEUS_SERVE_ADDR
          valueFrom:
            configMapKeyRef:
              key: prometheus-serve-addr
              name: cilium-metrics-config
              optional: true
        - name: CILIUM_LEGACY_HOST_ALLOWS_WORLD
          valueFrom:
            configMapKeyRef:
              key: legacy-host-allows-world
              name: cilium-config
              optional: true
        - name: CILIUM_SIDECAR_ISTIO_PROXY_IMAGE
          valueFrom:
            configMapKeyRef:
              key: sidecar-istio-proxy-image
              name: cilium-config
              optional: true
        - name: CILIUM_TUNNEL
          valueFrom:
            configMapKeyRef:
              key: tunnel
              name: cilium-config
              optional: true
        - name: CILIUM_MONITOR_AGGREGATION_LEVEL
          valueFrom:
            configMapKeyRef:
              key: monitor-aggregation-level
              name: cilium-config
              optional: true
        - name: CILIUM_CLUSTERMESH_CONFIG
          value: /var/lib/cilium/clustermesh/
        - name: CILIUM_CLUSTER_NAME
          valueFrom:
            configMapKeyRef:
              key: cluster-name
              name: cilium-config
              optional: true
        - name: CILIUM_CLUSTER_ID
          valueFrom:
            configMapKeyRef:
              key: cluster-id
              name: cilium-config
              optional: true
        - name: CILIUM_GLOBAL_CT_MAX_TCP
          valueFrom:
            configMapKeyRef:
              key: ct-global-max-entries-tcp
              name: cilium-config
              optional: true
        - name: CILIUM_GLOBAL_CT_MAX_ANY
          valueFrom:
            configMapKeyRef:
              key: ct-global-max-entries-other
              name: cilium-config
              optional: true
        - name: CILIUM_PREALLOCATE_BPF_MAPS
          valueFrom:
            configMapKeyRef:
              key: preallocate-bpf-maps
              name: cilium-config
              optional: true
        - name: CILIUM_FLANNEL_MASTER_DEVICE
          valueFrom:
            configMapKeyRef:
              key: flannel-master-device
              name: cilium-config
              optional: true
        - name: CILIUM_FLANNEL_UNINSTALL_ON_EXIT
          valueFrom:
            configMapKeyRef:
              key: flannel-uninstall-on-exit
              name: cilium-config
              optional: true
        - name: CILIUM_FLANNEL_MANAGE_EXISTING_CONTAINERS
          valueFrom:
            configMapKeyRef:
              key: flannel-manage-existing-containers
              name: cilium-config
              optional: true
        - name: CILIUM_DATAPATH_MODE
          valueFrom:
            configMapKeyRef:
              key: datapath-mode
              name: cilium-config
              optional: true
        - name: CILIUM_IPVLAN_MASTER_DEVICE
          valueFrom:
            configMapKeyRef:
              key: ipvlan-master-device
              name: cilium-config
              optional: true
        - name: CILIUM_INSTALL_IPTABLES_RULES
          valueFrom:
            configMapKeyRef:
              key: install-iptables-rules
              name: cilium-config
              optional: true
        - name: CILIUM_MASQUERADE
          valueFrom:
            configMapKeyRef:
              key: masquerade
              name: cilium-config
              optional: true
        - name: CILIUM_AUTO_DIRECT_NODE_ROUTES
          valueFrom:
            configMapKeyRef:
              key: auto-direct-node-routes
              name: cilium-config
              optional: true
        - name: CILIUM_TOFQDNS_ENABLE_POLLER
          valueFrom:
            configMapKeyRef:
              key: tofqdns-enable-poller
              name: cilium-config
              optional: true
        image: docker.io/cilium/cilium:v1.4
        imagePullPolicy: Always
        lifecycle:
          postStart:
            exec:
              command:
              - /cni-install.sh
          preStop:
            exec:
              command:
              - /cni-uninstall.sh
        livenessProbe:
          exec:
            command:
            - cilium
            - status
          failureThreshold: 10
          initialDelaySeconds: 120
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 1
        name: cilium-agent
        ports:
        - containerPort: 9090
          hostPort: 9090
          name: prometheus
          protocol: TCP
        readinessProbe:
          exec:
            command:
            - cilium
            - status
          failureThreshold: 3
          initialDelaySeconds: 5
          periodSeconds: 5
          successThreshold: 1
          timeoutSeconds: 1
        securityContext:
          capabilities:
            add:
            - NET_ADMIN
          privileged: true
        volumeMounts:
        - mountPath: /sys/fs/bpf
          name: bpf-maps
        - mountPath: /var/run/cilium
          name: cilium-run
        - mountPath: /host/opt/cni/bin
          name: cni-path
        - mountPath: /host/etc/cni/net.d
          name: etc-cni-netd
        - mountPath: /var/run/docker.sock
          name: docker-socket
          readOnly: true
        - mountPath: /var/lib/etcd-config
          name: etcd-config-path
          readOnly: true
        - mountPath: /var/lib/etcd-secrets
          name: etcd-secrets
          readOnly: true
        - mountPath: /var/lib/cilium/clustermesh
          name: clustermesh-secrets
          readOnly: true
      dnsPolicy: ClusterFirstWithHostNet
      hostNetwork: true
      hostPID: false
      initContainers:
      - command:
        - /init-container.sh
        env:
        - name: CLEAN_CILIUM_STATE
          valueFrom:
            configMapKeyRef:
              key: clean-cilium-state
              name: cilium-config
              optional: true
        - name: CLEAN_CILIUM_BPF_STATE
          valueFrom:
            configMapKeyRef:
              key: clean-cilium-bpf-state
              name: cilium-config
              optional: true
        image: docker.io/cilium/cilium-init:2018-10-16
        imagePullPolicy: IfNotPresent
        name: clean-cilium-state
        securityContext:
          capabilities:
            add:
            - NET_ADMIN
          privileged: true
        volumeMounts:
        - mountPath: /sys/fs/bpf
          name: bpf-maps
        - mountPath: /var/run/cilium
          name: cilium-run
      priorityClassName: system-node-critical
      restartPolicy: Always
      serviceAccount: cilium
      serviceAccountName: cilium
      terminationGracePeriodSeconds: 1
      tolerations:
      - operator: Exists
      - effect: NoSchedule
        key: node.kubernetes.io/not-ready
        operator: Exists
      volumes:
      - hostPath:
          path: /var/run/cilium
          type: DirectoryOrCreate
        name: cilium-run
      - hostPath:
          path: /sys/fs/bpf
          type: DirectoryOrCreate
        name: bpf-maps
      - hostPath:
          path: /var/run/docker.sock
          type: Socket
        name: docker-socket
      - hostPath:
          path: /opt/cni/bin
          type: DirectoryOrCreate
        name: cni-path
      - hostPath:
          path: /etc/cni/net.d
          type: DirectoryOrCreate
        name: etc-cni-netd
      - configMap:
          defaultMode: 420
          items:
          - key: etcd-config
            path: etcd.config
          name: cilium-config
        name: etcd-config-path
      - name: etcd-secrets
        secret:
          defaultMode: 420
          optional: true
          secretName: cilium-etcd-secrets
      - name: clustermesh-secrets
        secret:
          defaultMode: 420
          optional: true
          secretName: cilium-clustermesh
  updateStrategy:
    rollingUpdate:
      maxUnavailable: 2
    type: RollingUpdate
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    io.cilium/app: operator
    name: cilium-operator
  name: cilium-operator
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  replicas: 1
  selector:
    matchLabels:
      io.cilium/app: operator
      name: cilium-operator
  strategy:
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 1
    type: RollingUpdate
  template:
    metadata:
      labels:
        io.cilium/app: operator
        name: cilium-operator
    spec:
      containers:
      - args:
        - --debug=$(CILIUM_DEBUG)
        - --kvstore=etcd
        - --kvstore-opt=etcd.config=/var/lib/etcd-config/etcd.config
        command:
        - cilium-operator
        env:
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        - name: K8S_NODE_NAME
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: spec.nodeName
        - name: CILIUM_DEBUG
          valueFrom:
            configMapKeyRef:
              key: debug
              name: cilium-config
              optional: true
        - name: CILIUM_CLUSTER_NAME
          valueFrom:
            configMapKeyRef:
              key: cluster-name
              name: cilium-config
              optional: true
        - name: CILIUM_CLUSTER_ID
          valueFrom:
            configMapKeyRef:
              key: cluster-id
              name: cilium-config
              optional: true
        - name: CILIUM_DISABLE_ENDPOINT_CRD
          valueFrom:
            configMapKeyRef:
              key: disable-endpoint-crd
              name: cilium-config
              optional: true
        - name: AWS_ACCESS_KEY_ID
          valueFrom:
            secretKeyRef:
              key: AWS_ACCESS_KEY_ID
              name: cilium-aws
              optional: true
        - name: AWS_SECRET_ACCESS_KEY
          valueFrom:
            secretKeyRef:
              key: AWS_SECRET_ACCESS_KEY
              name: cilium-aws
              optional: true
        - name: AWS_DEFAULT_REGION
          valueFrom:
            secretKeyRef:
              key: AWS_DEFAULT_REGION
              name: cilium-aws
              optional: true
        image: docker.io/cilium/operator:v1.4
        imagePullPolicy: Always
        name: cilium-operator
        volumeMounts:
        - mountPath: /var/lib/etcd-config
          name: etcd-config-path
          readOnly: true
        - mountPath: /var/lib/etcd-secrets
          name: etcd-secrets
          readOnly: true
      dnsPolicy: ClusterFirst
      priorityClassName: system-node-critical
      restartPolicy: Always
      serviceAccount: cilium-operator
      serviceAccountName: cilium-operator
      volumes:
      - configMap:
          defaultMode: 420
          items:
          - key: etcd-config
            path: etcd.config
          name: cilium-config
        name: etcd-config-path
      - name: etcd-secrets
        secret:
          defaultMode: 420
          optional: true
          secretName: cilium-etcd-secrets
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: cilium-operator
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: cilium-operator
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
rules:
- apiGroups:
  - ""
  resources:
  - pods
  - deployments
  - componentstatuses
  verbs:
  - '*'
- apiGroups:
  - ""
  resources:
  - services
  - endpoints
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - cilium.io
  resources:
  - ciliumnetworkpolicies
  - ciliumnetworkpolicies/status
  - ciliumendpoints
  - ciliumendpoints/status
  verbs:
  - '*'
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: cilium-operator
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cilium-operator
subjects:
- kind: ServiceAccount
  name: cilium-operator
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: cilium-etcd-operator
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
rules:
- apiGroups:
  - etcd.database.coreos.com
  resources:
  - etcdclusters
  verbs:
  - get
  - delete
  - create
- apiGroups:
  - apiextensions.k8s.io
  resources:
  - customresourcedefinitions
  verbs:
  - delete
  - get
  - create
- apiGroups:
  - ""
  resources:
  - deployments
  verbs:
  - delete
  - create
  - get
  - update
- apiGroups:
  - ""
  resources:
  - pods
  verbs:
  - list
  - delete
  - get
- apiGroups:
  - apps
  resources:
  - deployments
  verbs:
  - delete
  - create
  - get
  - update
- apiGroups:
  - ""
  resources:
  - componentstatuses
  verbs:
  - get
- apiGroups:
  - extensions
  resources:
  - deployments
  verbs:
  - delete
  - create
  - get
  - update
- apiGroups:
  - ""
  resources:
  - secrets
  verbs:
  - get
  - create
  - delete
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: cilium-etcd-operator
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cilium-etcd-operator
subjects:
- kind: ServiceAccount
  name: cilium-etcd-operator
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: etcd-operator
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
rules:
- apiGroups:
  - etcd.database.coreos.com
  resources:
  - etcdclusters
  - etcdbackups
  - etcdrestores
  verbs:
  - '*'
- apiGroups:
  - apiextensions.k8s.io
  resources:
  - customresourcedefinitions
  verbs:
  - '*'
- apiGroups:
  - ""
  resources:
  - pods
  - services
  - endpoints
  - persistentvolumeclaims
  - events
  - deployments
  verbs:
  - '*'
- apiGroups:
  - apps
  resources:
  - deployments
  verbs:
  - '*'
- apiGroups:
  - extensions
  resources:
  - deployments
  verbs:
  - create
  - get
  - list
  - patch
  - update
- apiGroups:
  - ""
  resources:
  - secrets
  verbs:
  - get
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: etcd-operator
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: etcd-operator
subjects:
- kind: ServiceAccount
  name: cilium-etcd-sa
  namespace: kube-system
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: cilium-etcd-operator
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: cilium-etcd-sa
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    io.cilium/app: etcd-operator
    name: cilium-etcd-operator
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: cilium-etcd-operator
  namespace: kube-system
spec:
  replicas: 1
  selector:
    matchLabels:
      io.cilium/app: etcd-operator
      name: cilium-etcd-operator
  strategy:
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 1
    type: RollingUpdate
  template:
    metadata:
      labels:
        io.cilium/app: etcd-operator
        name: cilium-etcd-operator
    spec:
      containers:
      - command:
        - /usr/bin/cilium-etcd-operator
        env:
        - name: CILIUM_ETCD_OPERATOR_CLUSTER_DOMAIN
          value: cluster.local
        - name: CILIUM_ETCD_OPERATOR_ETCD_CLUSTER_SIZE
          value: "3"
        - name: CILIUM_ETCD_OPERATOR_NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        - name: CILIUM_ETCD_OPERATOR_POD_NAME
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.name
        - name: CILIUM_ETCD_OPERATOR_POD_UID
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.uid
        image: docker.io/cilium/cilium-etcd-operator:v2.0.5
        imagePullPolicy: IfNotPresent
        name: cilium-etcd-operator
      dnsPolicy: ClusterFirst
      hostNetwork: true
      priorityClassName: system-node-critical
      restartPolicy: Always
      serviceAccount: cilium-etcd-operator
      serviceAccountName: cilium-etcd-operator
      tolerations:
      - operator: Exists
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: cilium
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cilium
subjects:
- kind: ServiceAccount
  name: cilium
  namespace: kube-system
- apiGroup: rbac.authorization.k8s.io
  kind: Group
  name: system:nodes
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: cilium
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
rules:
- apiGroups:
  - networking.k8s.io
  resources:
  - networkpolicies
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - namespaces
  - services
  - nodes
  - endpoints
  - componentstatuses
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - pods
  - nodes
  verbs:
  - get
  - list
  - watch
  - update
- apiGroups:
  - extensions
  resources:
  - ingresses
  verbs:
  - create
  - get
  - list
  - watch
- apiGroups:
  - apiextensions.k8s.io
  resources:
  - customresourcedefinitions
  verbs:
  - create
  - get
  - list
  - watch
  - update
- apiGroups:
  - cilium.io
  resources:
  - ciliumnetworkpolicies
  - ciliumnetworkpolicies/status
  - ciliumendpoints
  - ciliumendpoints/status
  verbs:
  - '*'
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: cilium
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"`)

func k8sAddonsKubernetesmasteraddonsCiliumDaemonsetYamlBytes() ([]byte, error) {
	return _k8sAddonsKubernetesmasteraddonsCiliumDaemonsetYaml, nil
}

func k8sAddonsKubernetesmasteraddonsCiliumDaemonsetYaml() (*asset, error) {
	bytes, err := k8sAddonsKubernetesmasteraddonsCiliumDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/kubernetesmasteraddons-cilium-daemonset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsKubernetesmasteraddonsElbSvcYaml = []byte(`---
apiVersion: v1
kind: Service
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
  name: elb
  namespace: kube-system
spec:
  ports:
  - port: 8765
    targetPort: 9376
    protocol: TCP
  selector:
    app: "<svcName>"
  type: LoadBalancer
---
apiVersion: v1
kind: Service
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
  name: elb-udp
  namespace: kube-system
spec:
  ports:
  - port: 8765
    targetPort: 9376
    protocol: UDP
  selector:
    app: "<svcName>"
  type: LoadBalancer`)

func k8sAddonsKubernetesmasteraddonsElbSvcYamlBytes() ([]byte, error) {
	return _k8sAddonsKubernetesmasteraddonsElbSvcYaml, nil
}

func k8sAddonsKubernetesmasteraddonsElbSvcYaml() (*asset, error) {
	bytes, err := k8sAddonsKubernetesmasteraddonsElbSvcYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/kubernetesmasteraddons-elb-svc.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsKubernetesmasteraddonsFlannelDaemonsetYaml = []byte(`# This file was pulled from:
# https://github.com/coreos/flannel (HEAD at time of pull was 4973e02e539378)
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: flannel
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
---
kind: ConfigMap
apiVersion: v1
metadata:
  name: kube-flannel-cfg
  namespace: kube-system
  labels:
    tier: node
    app: flannel
    addonmanager.kubernetes.io/mode: EnsureExists
data:
  cni-conf.json: |
    {
      "name": "cbr0",
      "type": "flannel",
      "delegate": {
        "isDefaultGateway": true
      }
    }
  net-conf.json: |
    {
      "Network": "<kubeClusterCidr>",
      "Backend": {
        "Type": "vxlan"
      }
    }
---
apiVersion: extensions/v1beta1
kind: DaemonSet
metadata:
  name: kube-flannel-ds
  namespace: kube-system
  labels:
    tier: node
    app: flannel
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  template:
    metadata:
      labels:
        tier: node
        app: flannel
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ''
    spec:
      hostNetwork: true
      nodeSelector:
        beta.kubernetes.io/arch: amd64
        beta.kubernetes.io/os: linux
      priorityClassName: system-node-critical
      tolerations:
        - key: node.kubernetes.io/not-ready
          operator: Exists
          effect: NoSchedule
        - key: node-role.kubernetes.io/master
          operator: Equal
          value: "true"
          effect: NoSchedule
        - key: CriticalAddonsOnly
          operator: Exists
      serviceAccountName: flannel
      containers:
      - name: kube-flannel
        image: quay.io/coreos/flannel:v0.8.0-amd64
        command: [ "/opt/bin/flanneld", "--ip-masq", "--kube-subnet-mgr" ]
        securityContext:
          privileged: true
        env:
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        volumeMounts:
        - name: run
          mountPath: /run
        - name: flannel-cfg
          mountPath: /etc/kube-flannel/
      - name: install-cni
        image: quay.io/coreos/flannel:v0.10.0-amd64
        command: [ "/bin/sh", "-c", "set -e -x; cp -f /etc/kube-flannel/cni-conf.json /etc/cni/net.d/10-flannel.conf; while true; do sleep 3600; done" ]
        volumeMounts:
        - name: cni
          mountPath: /etc/cni/net.d
        - name: flannel-cfg
          mountPath: /etc/kube-flannel/
      volumes:
        - name: run
          hostPath:
            path: /run
        - name: cni
          hostPath:
            path: /etc/cni/net.d
        - name: flannel-cfg
          configMap:
            name: kube-flannel-cfg
---
# This file was pulled from:
# https://github.com/coreos/flannel (HEAD at time of pull was 4973e02e539378)
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: flannel
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
  - apiGroups:
      - ""
    resources:
      - pods
    verbs:
      - get
  - apiGroups:
      - ""
    resources:
      - nodes
    verbs:
      - list
      - watch
  - apiGroups:
      - ""
    resources:
      - nodes/status
    verbs:
      - patch
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: flannel
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: flannel
subjects:
- kind: ServiceAccount
  name: flannel
  namespace: kube-system
`)

func k8sAddonsKubernetesmasteraddonsFlannelDaemonsetYamlBytes() ([]byte, error) {
	return _k8sAddonsKubernetesmasteraddonsFlannelDaemonsetYaml, nil
}

func k8sAddonsKubernetesmasteraddonsFlannelDaemonsetYaml() (*asset, error) {
	bytes, err := k8sAddonsKubernetesmasteraddonsFlannelDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/kubernetesmasteraddons-flannel-daemonset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsKubernetesmasteraddonsKubeDnsDeploymentYaml = []byte(`# Copyright 2016 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Should keep target in cluster/addons/dns-horizontal-autoscaler/dns-horizontal-autoscaler.yaml
# in sync with this file.

apiVersion: v1
kind: Service
metadata:
  name: kube-dns
  namespace: kube-system
  labels:
    k8s-app: kube-dns
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
    kubernetes.io/name: KubeDNS
spec:
  selector:
    k8s-app: kube-dns
  clusterIP: <clustIP>
  ports:
  - name: dns
    port: 53
    protocol: UDP
  - name: dns-tcp
    port: 53
    protocol: TCP
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: kube-dns
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: kube-dns
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: EnsureExists
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: kube-dns
  namespace: kube-system
  labels:
    k8s-app: kube-dns
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  # replicas: not specified here:
  # 1. In order to make Addon Manager do not reconcile this replicas parameter.
  # 2. Default is 1.
  # 3. Will be tuned in real time if DNS horizontal auto-scaling is turned on.
  strategy:
    rollingUpdate:
      maxSurge: 10%
      maxUnavailable: 0
  selector:
    matchLabels:
      k8s-app: kube-dns
  template:
    metadata:
      labels:
        k8s-app: kube-dns
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ''
        seccomp.security.alpha.kubernetes.io/pod: 'docker/default'
    spec:
      priorityClassName: system-node-critical
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      volumes:
      - name: kube-dns-config
        configMap:
          name: kube-dns
          optional: true
      containers:
      - name: kubedns
        image: <img>
        imagePullPolicy: IfNotPresent
        resources:
          # TODO: Set memory limits when we've profiled the container for large
          # clusters, then set request = limit to keep this container in
          # guaranteed class. Currently, this container falls into the
          # "burstable" category so the kubelet doesn't backoff from restarting it.
          limits:
            memory: 170Mi
          requests:
            cpu: 100m
            memory: 70Mi
        livenessProbe:
          httpGet:
            path: /healthcheck/kubedns
            port: 10054
            scheme: HTTP
          initialDelaySeconds: 60
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 5
        readinessProbe:
          httpGet:
            path: /readiness
            port: 8081
            scheme: HTTP
          initialDelaySeconds: 3
          timeoutSeconds: 5
        args:
        - --domain=<domain>.
        - --dns-port=10053
        - --config-dir=/kube-dns-config
        - --v=2
        env:
        - name: PROMETHEUS_PORT
          value: "10055"
        ports:
        - containerPort: 10053
          name: dns-local
          protocol: UDP
        - containerPort: 10053
          name: dns-tcp-local
          protocol: TCP
        - containerPort: 10055
          name: metrics
          protocol: TCP
        volumeMounts:
        - name: kube-dns-config
          mountPath: /kube-dns-config
      - name: dnsmasq
        image: <imgMasq>
        imagePullPolicy: IfNotPresent
        livenessProbe:
          httpGet:
            path: /healthcheck/dnsmasq
            port: 10054
            scheme: HTTP
          initialDelaySeconds: 60
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 5
        args:
        - -v=2
        - -logtostderr
        - -configDir=/etc/k8s/dns/dnsmasq-nanny
        - -restartDnsmasq=true
        - --
        - -k
        - --cache-size=1000
        - --no-negcache
        - --log-facility=-
        - --server=/cluster.local/127.0.0.1#10053
        - --server=/in-addr.arpa/127.0.0.1#10053
        - --server=/ip6.arpa/127.0.0.1#10053
        ports:
        - containerPort: 53
          name: dns
          protocol: UDP
        - containerPort: 53
          name: dns-tcp
          protocol: TCP
        resources:
          requests:
            cpu: 150m
            memory: 20Mi
        volumeMounts:
        - name: kube-dns-config
          mountPath: /etc/k8s/dns/dnsmasq-nanny
      - name: sidecar
        image: <imgSidecar>
        imagePullPolicy: IfNotPresent
        livenessProbe:
          httpGet:
            path: /metrics
            port: 10054
            scheme: HTTP
          initialDelaySeconds: 60
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 5
        args:
        - --v=2
        - --logtostderr
        - --probe=kubedns,127.0.0.1:10053,kubernetes.default.svc.<domain>,5,SRV
        - --probe=dnsmasq,127.0.0.1:53,kubernetes.default.svc.<domain>,5,SRV
        ports:
        - containerPort: 10054
          name: metrics
          protocol: TCP
        resources:
          requests:
            memory: 20Mi
            cpu: 10m
      dnsPolicy: Default
      serviceAccountName: kube-dns
      nodeSelector:
        beta.kubernetes.io/os: linux`)

func k8sAddonsKubernetesmasteraddonsKubeDnsDeploymentYamlBytes() ([]byte, error) {
	return _k8sAddonsKubernetesmasteraddonsKubeDnsDeploymentYaml, nil
}

func k8sAddonsKubernetesmasteraddonsKubeDnsDeploymentYaml() (*asset, error) {
	bytes, err := k8sAddonsKubernetesmasteraddonsKubeDnsDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/kubernetesmasteraddons-kube-dns-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsKubernetesmasteraddonsKubeProxyDaemonsetYaml = []byte(`apiVersion: extensions/v1beta1
kind: DaemonSet
metadata:
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
    kubernetes.io/cluster-service: "true"
    component: kube-proxy
    tier: node
    k8s-app: kube-proxy
  name: kube-proxy
  namespace: kube-system
spec:
  selector:
    matchLabels:
      k8s-app: kube-proxy
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 50%
  template:
    metadata:
      labels:
        component: kube-proxy
        tier: node
        k8s-app: kube-proxy
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ''
    spec:
      priorityClassName: system-node-critical
      tolerations:
      - key: node-role.kubernetes.io/master
        operator: Equal
        value: "true"
        effect: NoSchedule
      - operator: "Exists"
        effect: NoExecute
      - operator: "Exists"
        effect: NoSchedule
      - key: CriticalAddonsOnly
        operator: Exists
      containers:
      - command:
        - /hyperkube
        - kube-proxy
        - --kubeconfig=/var/lib/kubelet/kubeconfig
        - --cluster-cidr=<CIDR>
        - --feature-gates=ExperimentalCriticalPodAnnotation=true
        - --proxy-mode=<kubeProxyMode>
        image: <img>
        imagePullPolicy: IfNotPresent
        name: kube-proxy
        resources:
          requests:
            cpu: 100m
        securityContext:
          privileged: true
        volumeMounts:
        - mountPath: /etc/ssl/certs
          name: ssl-certs-host
          readOnly: true
        - mountPath: /etc/kubernetes
          name: etc-kubernetes
          readOnly: true
        - mountPath: /var/lib/kubelet/kubeconfig
          name: kubeconfig
          readOnly: true
        - mountPath: /run/xtables.lock
          name: iptableslock
        - mountPath: /lib/modules/
          name: kernelmodules
          readOnly: true
      hostNetwork: true
      volumes:
      - hostPath:
          path: /usr/share/ca-certificates
        name: ssl-certs-host
      - hostPath:
          path: /var/lib/kubelet/kubeconfig
        name: kubeconfig
      - hostPath:
          path: /etc/kubernetes
        name: etc-kubernetes
      - hostPath:
          path: /run/xtables.lock
        name: iptableslock
      - hostPath:
          path: /lib/modules/
        name: kernelmodules
      nodeSelector:
        beta.kubernetes.io/os: linux
`)

func k8sAddonsKubernetesmasteraddonsKubeProxyDaemonsetYamlBytes() ([]byte, error) {
	return _k8sAddonsKubernetesmasteraddonsKubeProxyDaemonsetYaml, nil
}

func k8sAddonsKubernetesmasteraddonsKubeProxyDaemonsetYaml() (*asset, error) {
	bytes, err := k8sAddonsKubernetesmasteraddonsKubeProxyDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/kubernetesmasteraddons-kube-proxy-daemonset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsKubernetesmasteraddonsManagedAzureStorageClassesCustomYaml = []byte(`apiVersion: storage.k8s.io/v1beta1
kind: StorageClass
metadata:
  name: default
  annotations:
    storageclass.beta.kubernetes.io/is-default-class: "true"
  labels:
    kubernetes.io/cluster-service: "true"
provisioner: kubernetes.io/azure-disk
parameters:
  kind: Managed
  storageaccounttype: Standard_LRS
  cachingmode: ReadOnly
---
apiVersion: storage.k8s.io/v1beta1
kind: StorageClass
metadata:
  name: managed-premium
  annotations:
  labels:
    kubernetes.io/cluster-service: "true"
provisioner: kubernetes.io/azure-disk
parameters:
  kind: Managed
  storageaccounttype: Premium_LRS
  cachingmode: ReadOnly
---
apiVersion: storage.k8s.io/v1beta1
kind: StorageClass
metadata:
  name: managed-standard
  annotations:
  labels:
    kubernetes.io/cluster-service: "true"
provisioner: kubernetes.io/azure-disk
parameters:
  kind: Managed
  storageaccounttype: Standard_LRS
  cachingmode: ReadOnly
`)

func k8sAddonsKubernetesmasteraddonsManagedAzureStorageClassesCustomYamlBytes() ([]byte, error) {
	return _k8sAddonsKubernetesmasteraddonsManagedAzureStorageClassesCustomYaml, nil
}

func k8sAddonsKubernetesmasteraddonsManagedAzureStorageClassesCustomYaml() (*asset, error) {
	bytes, err := k8sAddonsKubernetesmasteraddonsManagedAzureStorageClassesCustomYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/kubernetesmasteraddons-managed-azure-storage-classes-custom.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsKubernetesmasteraddonsManagedAzureStorageClassesYaml = []byte(`apiVersion: storage.k8s.io/v1beta1
kind: StorageClass
metadata:
  name: default
  annotations:
    storageclass.beta.kubernetes.io/is-default-class: "true"
  labels:
    kubernetes.io/cluster-service: "true"
provisioner: kubernetes.io/azure-disk
parameters:
  kind: Managed
  storageaccounttype: Standard_LRS
  cachingmode: ReadOnly
---
apiVersion: storage.k8s.io/v1beta1
kind: StorageClass
metadata:
  name: managed-premium
  annotations:
  labels:
    kubernetes.io/cluster-service: "true"
provisioner: kubernetes.io/azure-disk
parameters:
  kind: Managed
  storageaccounttype: Premium_LRS
  cachingmode: ReadOnly
---
apiVersion: storage.k8s.io/v1beta1
kind: StorageClass
metadata:
  name: managed-standard
  annotations:
  labels:
    kubernetes.io/cluster-service: "true"
provisioner: kubernetes.io/azure-disk
parameters:
  kind: Managed
  storageaccounttype: Standard_LRS
  cachingmode: ReadOnly
---
kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: azurefile
  annotations:
  labels:
    kubernetes.io/cluster-service: "true"
provisioner: kubernetes.io/azure-file
parameters:
  skuName: Standard_LRS
`)

func k8sAddonsKubernetesmasteraddonsManagedAzureStorageClassesYamlBytes() ([]byte, error) {
	return _k8sAddonsKubernetesmasteraddonsManagedAzureStorageClassesYaml, nil
}

func k8sAddonsKubernetesmasteraddonsManagedAzureStorageClassesYaml() (*asset, error) {
	bytes, err := k8sAddonsKubernetesmasteraddonsManagedAzureStorageClassesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/kubernetesmasteraddons-managed-azure-storage-classes.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsKubernetesmasteraddonsPodSecurityPolicyYaml = []byte(`apiVersion: extensions/v1beta1
kind: PodSecurityPolicy
metadata:
  name: privileged
  annotations:
    seccomp.security.alpha.kubernetes.io/allowedProfileNames: "*"
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  privileged: true
  allowPrivilegeEscalation: true
  allowedCapabilities:
  - "*"
  volumes:
  - "*"
  hostNetwork: true
  hostPorts:
  - min: 0
    max: 65535
  hostIPC: true
  hostPID: true
  runAsUser:
    rule: RunAsAny
  seLinux:
    rule: RunAsAny
  supplementalGroups:
    rule: RunAsAny
  fsGroup:
    rule: RunAsAny
---
apiVersion: extensions/v1beta1
kind: PodSecurityPolicy
metadata:
  name: restricted
  annotations:
    seccomp.security.alpha.kubernetes.io/allowedProfileNames: docker/default
    apparmor.security.beta.kubernetes.io/allowedProfileNames: runtime/default
    seccomp.security.alpha.kubernetes.io/defaultProfileName:  docker/default
    apparmor.security.beta.kubernetes.io/defaultProfileName:  runtime/default
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  privileged: false
  allowPrivilegeEscalation: false
  requiredDropCapabilities:
    - ALL
  volumes:
    - configMap
    - emptyDir
    - projected
    - secret
    - downwardAPI
    - persistentVolumeClaim
  hostNetwork: false
  hostIPC: false
  hostPID: false
  runAsUser:
    rule: MustRunAsNonRoot
  seLinux:
    rule: RunAsAny
  supplementalGroups:
    rule: MustRunAs
    ranges:
      # Forbid adding the root group.
      - min: 1
        max: 65535
  fsGroup:
    rule: MustRunAs
    ranges:
      # Forbid adding the root group.
      - min: 1
        max: 65535
  readOnlyRootFilesystem: false
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: psp:privileged
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups: ['extensions']
  resources: ['podsecuritypolicies']
  verbs:     ['use']
  resourceNames:
  - privileged
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: psp:restricted
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups: ['extensions']
  resources: ['podsecuritypolicies']
  verbs:     ['use']
  resourceNames:
  - restricted
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: default:restricted
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: psp:restricted
subjects:
- kind: Group
  name: system:authenticated
  apiGroup: rbac.authorization.k8s.io
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: default:privileged
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: psp:privileged
subjects:
- kind: Group
  name: system:masters
  apiGroup: rbac.authorization.k8s.io
- kind: Group
  name: system:serviceaccounts:kube-system
  apiGroup: rbac.authorization.k8s.io
- kind: Group
  name: system:nodes
  apiGroup: rbac.authorization.k8s.io`)

func k8sAddonsKubernetesmasteraddonsPodSecurityPolicyYamlBytes() ([]byte, error) {
	return _k8sAddonsKubernetesmasteraddonsPodSecurityPolicyYaml, nil
}

func k8sAddonsKubernetesmasteraddonsPodSecurityPolicyYaml() (*asset, error) {
	bytes, err := k8sAddonsKubernetesmasteraddonsPodSecurityPolicyYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/kubernetesmasteraddons-pod-security-policy.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsKubernetesmasteraddonsScheduledMaintenanceDeploymentYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  labels:
    control-plane: controller-manager
    addonmanager.kubernetes.io/mode: Reconcile
  name: drainsafe-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: drainsafe-leader-election-role
  namespace: drainsafe-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - ""
  resources:
  - configmaps/status
  verbs:
  - get
  - update
  - patch
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  creationTimestamp: null
  name: drainsafe-manager-role
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups:
  - apiextensions.k8s.io
  resources:
  - customresourcedefinitions
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - repairman.k8s.io
  resources:
  - maintenancerequests
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - ""
  resources:
  - nodes
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - apps
  resources:
  - daemonsets
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - extensions
  resources:
  - daemonsets
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - ""
  resources:
  - pods
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - ""
  resources:
  - pods/eviction
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - ""
  resources:
  - events
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: drainsafe-proxy-role
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups:
  - authentication.k8s.io
  resources:
  - tokenreviews
  - subjectaccessreviews
  verbs:
  - create
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: drainsafe-leader-election-rolebinding
  namespace: drainsafe-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: drainsafe-leader-election-role
subjects:
- kind: ServiceAccount
  name: default
  namespace: drainsafe-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: drainsafe-manager-rolebinding
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: drainsafe-manager-role
subjects:
- kind: ServiceAccount
  name: default
  namespace: drainsafe-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: drainsafe-proxy-rolebinding
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: drainsafe-proxy-role
subjects:
- kind: ServiceAccount
  name: default
  namespace: drainsafe-system
---
apiVersion: v1
kind: Service
metadata:
  annotations:
    prometheus.io/port: "8443"
    prometheus.io/scheme: https
    prometheus.io/scrape: "true"
  labels:
    control-plane: controller-manager
    addonmanager.kubernetes.io/mode: Reconcile
  name: drainsafe-controller-manager-metrics-service
  namespace: drainsafe-system
spec:
  ports:
  - name: https
    port: 8443
    targetPort: https
  selector:
    control-plane: controller-manager
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    control-plane: controller-manager
    addonmanager.kubernetes.io/mode: Reconcile
  name: drainsafe-controller-manager
  namespace: drainsafe-system
spec:
  replicas: 1
  selector:
    matchLabels:
      control-plane: controller-manager
  template:
    metadata:
      labels:
        control-plane: controller-manager
    spec:
      containers:
      - args:
        - --secure-listen-address=0.0.0.0:8443
        - --upstream=http://127.0.0.1:8080/
        - --logtostderr=true
        - --v=10
        image: gcr.io/kubebuilder/kube-rbac-proxy:v0.4.0
        name: kube-rbac-proxy
        ports:
        - containerPort: 8443
          name: https
      - args:
        - --metrics-addr=127.0.0.1:8080
        command:
        - /manager
        env:
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        image: quay.io/awesomenix/drainsafe-manager:latest
        name: manager
        resources:
          limits:
            cpu: 100m
            memory: 30Mi
          requests:
            cpu: 100m
            memory: 20Mi
      terminationGracePeriodSeconds: 10
      tolerations:
      - effect: NoSchedule
        key: node-role.kubernetes.io/master
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  labels:
    control-plane: controller-manager
    addonmanager.kubernetes.io/mode: Reconcile
  name: drainsafe-controller-scheduledevent-manager
  namespace: drainsafe-system
spec:
  selector:
    matchLabels:
      control-plane: controller-manager
  template:
    metadata:
      labels:
        control-plane: controller-manager
    spec:
      containers:
      - command:
        - /scheduledevent-manager
        env:
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        image: quay.io/awesomenix/drainsafe-manager:latest
        name: manager
        resources:
          limits:
            cpu: 100m
            memory: 30Mi
          requests:
            cpu: 100m
            memory: 20Mi
      terminationGracePeriodSeconds: 10
      tolerations:
      - effect: NoSchedule
        key: node-role.kubernetes.io/master
`)

func k8sAddonsKubernetesmasteraddonsScheduledMaintenanceDeploymentYamlBytes() ([]byte, error) {
	return _k8sAddonsKubernetesmasteraddonsScheduledMaintenanceDeploymentYaml, nil
}

func k8sAddonsKubernetesmasteraddonsScheduledMaintenanceDeploymentYaml() (*asset, error) {
	bytes, err := k8sAddonsKubernetesmasteraddonsScheduledMaintenanceDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/kubernetesmasteraddons-scheduled-maintenance-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsKubernetesmasteraddonsUnmanagedAzureStorageClassesCustomYaml = []byte(`apiVersion: storage.k8s.io/v1beta1
kind: StorageClass
metadata:
  name: default
  annotations:
    storageclass.beta.kubernetes.io/is-default-class: "true"
  labels:
    kubernetes.io/cluster-service: "true"
provisioner: kubernetes.io/azure-disk
parameters:
  cachingmode: ReadOnly
---
apiVersion: storage.k8s.io/v1beta1
kind: StorageClass
metadata:
  name: unmanaged-premium
  annotations:
  labels:
    kubernetes.io/cluster-service: "true"
provisioner: kubernetes.io/azure-disk
parameters:
  kind: shared
  storageaccounttype: Premium_LRS
  cachingmode: ReadOnly
---
apiVersion: storage.k8s.io/v1beta1
kind: StorageClass
metadata:
  name: unmanaged-standard
  annotations:
  labels:
    kubernetes.io/cluster-service: "true"
provisioner: kubernetes.io/azure-disk
parameters:
  kind: shared
  storageaccounttype: Standard_LRS
  cachingmode: ReadOnly
`)

func k8sAddonsKubernetesmasteraddonsUnmanagedAzureStorageClassesCustomYamlBytes() ([]byte, error) {
	return _k8sAddonsKubernetesmasteraddonsUnmanagedAzureStorageClassesCustomYaml, nil
}

func k8sAddonsKubernetesmasteraddonsUnmanagedAzureStorageClassesCustomYaml() (*asset, error) {
	bytes, err := k8sAddonsKubernetesmasteraddonsUnmanagedAzureStorageClassesCustomYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/kubernetesmasteraddons-unmanaged-azure-storage-classes-custom.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsKubernetesmasteraddonsUnmanagedAzureStorageClassesYaml = []byte(`apiVersion: storage.k8s.io/v1beta1
kind: StorageClass
metadata:
  name: default
  annotations:
    storageclass.beta.kubernetes.io/is-default-class: "true"
  labels:
    kubernetes.io/cluster-service: "true"
provisioner: kubernetes.io/azure-disk
parameters:
  cachingmode: ReadOnly
---
apiVersion: storage.k8s.io/v1beta1
kind: StorageClass
metadata:
  name: unmanaged-premium
  annotations:
  labels:
    kubernetes.io/cluster-service: "true"
provisioner: kubernetes.io/azure-disk
parameters:
  kind: shared
  storageaccounttype: Premium_LRS
  cachingmode: ReadOnly
---
apiVersion: storage.k8s.io/v1beta1
kind: StorageClass
metadata:
  name: unmanaged-standard
  annotations:
  labels:
    kubernetes.io/cluster-service: "true"
provisioner: kubernetes.io/azure-disk
parameters:
  kind: shared
  storageaccounttype: Standard_LRS
  cachingmode: ReadOnly
---
kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: azurefile
  annotations:
  labels:
    kubernetes.io/cluster-service: "true"
provisioner: kubernetes.io/azure-file
parameters:
  skuName: Standard_LRS
`)

func k8sAddonsKubernetesmasteraddonsUnmanagedAzureStorageClassesYamlBytes() ([]byte, error) {
	return _k8sAddonsKubernetesmasteraddonsUnmanagedAzureStorageClassesYaml, nil
}

func k8sAddonsKubernetesmasteraddonsUnmanagedAzureStorageClassesYaml() (*asset, error) {
	bytes, err := k8sAddonsKubernetesmasteraddonsUnmanagedAzureStorageClassesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/kubernetesmasteraddons-unmanaged-azure-storage-classes.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sArmparametersT = []byte(`{
  "parameters": {
    {{range .AgentPoolProfiles}}{{template "agentparams.t" .}},{{end}}
    {{if .HasWindows}}
      {{template "windowsparams.t"}},
    {{end}}
    {{template "masterparams.t" .}},
    {{template "k8s/kubernetesparams.t" .}}
  }
}`)

func k8sArmparametersTBytes() ([]byte, error) {
	return _k8sArmparametersT, nil
}

func k8sArmparametersT() (*asset, error) {
	bytes, err := k8sArmparametersTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/armparameters.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsAptPreferences = []byte(``)

func k8sCloudInitArtifactsAptPreferencesBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsAptPreferences, nil
}

func k8sCloudInitArtifactsAptPreferences() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsAptPreferencesBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/apt-preferences", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsAuditdRules = []byte(`# increase kernel audit buffers since we have a lot of rules
-b 8192

# 4.1.4 Ensure events that modify date and time information are collected
-a always,exit -F arch=b64 -S adjtimex -S settimeofday -k time-change
-a always,exit -F arch=b32 -S adjtimex -S settimeofday -S stime -k time-change
-a always,exit -F arch=b64 -S clock_settime -k time-change
-a always,exit -F arch=b32 -S clock_settime -k time-change
-w /etc/localtime -p wa -k time-change

# 4.1.5 Ensure events that modify user/group information are collected
-w /etc/group -p wa -k identity
-w /etc/passwd -p wa -k identity
-w /etc/gshadow -p wa -k identity
-w /etc/shadow -p wa -k identity
-w /etc/security/opasswd -p wa -k identity

# 4.1.6 Ensure events that modify the system's network environment are collected
-a always,exit -F arch=b64 -S sethostname -S setdomainname -k system-locale
-a always,exit -F arch=b32 -S sethostname -S setdomainname -k system-locale
-w /etc/issue -p wa -k system-locale
-w /etc/issue.net -p wa -k system-locale
-w /etc/hosts -p wa -k system-locale
-w /etc/network -p wa -k system-locale
-w /etc/networks -p wa -k system-locale

# 4.1.7 Ensure events that modify the system's Mandatory Access Controls are collected
-w /etc/selinux/ -p wa -k MAC-policy

# 4.1.8 Ensure login and logout events are collected
-w /var/log/faillog -p wa -k logins
-w /var/log/lastlog -p wa -k logins
-w /var/log/tallylog -p wa -k logins

# 4.1.9 Ensure session initiation information is collected
-w /var/run/utmp -p wa -k session
-w /var/log/wtmp -p wa -k session
-w /var/log/btmp -p wa -k session

# 4.1.10 Ensure discretionary access control permission modification events are collected
-a always,exit -F arch=b64 -S chmod -S fchmod -S fchmodat -F auid>=1000 -F auid!=4294967295 -k perm_mod
-a always,exit -F arch=b32 -S chmod -S fchmod -S fchmodat -F auid>=1000 -F auid!=4294967295 -k perm_mod
-a always,exit -F arch=b64 -S chown -S fchown -S fchownat -S lchown -F auid>=1000 -F auid!=4294967295 -k perm_mod
-a always,exit -F arch=b32 -S chown -S fchown -S fchownat -S lchown -F auid>=1000 -F auid!=4294967295 -k perm_mod
-a always,exit -F arch=b64 -S setxattr -S lsetxattr -S fsetxattr -S removexattr -S lremovexattr -S fremovexattr -F auid>=1000 -F auid!=4294967295 -k perm_mod
-a always,exit -F arch=b32 -S setxattr -S lsetxattr -S fsetxattr -S removexattr -S lremovexattr -S fremovexattr -F auid>=1000 -F auid!=4294967295 -k perm_mod

# 4.1.11 Ensure unsuccessful unauthorized file access attempts are collected
-a always,exit -F arch=b64 -S creat -S open -S openat -S truncate -S ftruncate -F exit=-EACCES -F auid>=1000 -F auid!=4294967295 -k access
-a always,exit -F arch=b32 -S creat -S open -S openat -S truncate -S ftruncate -F exit=-EACCES -F auid>=1000 -F auid!=4294967295 -k access
-a always,exit -F arch=b64 -S creat -S open -S openat -S truncate -S ftruncate -F exit=-EPERM -F auid>=1000 -F auid!=4294967295 -k access
-a always,exit -F arch=b32 -S creat -S open -S openat -S truncate -S ftruncate -F exit=-EPERM -F auid>=1000 -F auid!=4294967295 -k access

# 4.1.12 Ensure use of privileged commands is collected
-a always,exit -F path=/usr/lib/dbus-1.0/dbus-daemon-launch-helper -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/usr/lib/openssh/ssh-keysign -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/usr/lib/eject/dmcrypt-get-device -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/usr/bin/sudo -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/usr/bin/wall -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/usr/bin/ssh-agent -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/usr/bin/expiry -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/usr/bin/chfn -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/usr/bin/pkexec -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/usr/bin/screen -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/usr/bin/chsh -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/usr/bin/newgidmap -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/usr/bin/chage -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/usr/bin/crontab -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/usr/bin/at -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/usr/bin/newgrp -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/usr/bin/mlocate -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/usr/bin/gpasswd -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/usr/bin/newuidmap -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/usr/bin/passwd -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/usr/bin/bsd-write -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/bin/umount -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/bin/mount -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/bin/ntfs-3g -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/bin/ping6 -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/bin/su -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/bin/ping -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/bin/fusermount -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/sbin/pam_extrausers_chkpwd -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/sbin/mount.nfs -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged
-a always,exit -F path=/sbin/unix_chkpwd -F perm=x -F auid>=1000 -F auid!=4294967295  -k privileged

# 4.1.13 Ensure successful file system mounts are collected
-a always,exit -F arch=b64 -S mount -F auid>=1000 -F auid!=4294967295 -k mounts
-a always,exit -F arch=b32 -S mount -F auid>=1000 -F auid!=4294967295 -k mounts

# 4.1.14 Ensure file deletion events by users are collected
-a always,exit -F arch=b64 -S unlink -S unlinkat -S rename -S renameat -F auid>=1000 -F auid!=4294967295 -k delete
-a always,exit -F arch=b32 -S unlink -S unlinkat -S rename -S renameat -F auid>=1000 -F auid!=4294967295 -k delete

# 4.1.15 Ensure changes to system administration scope (sudoers) is collected
-w /etc/sudoers -p wa -k scope
-w /etc/sudoers.d -p wa -k scope

# 4.1.16 Ensure system administrator actions (sudolog) are collected
-w /var/log/sudo.log -p wa -k actions

# 4.1.17 Ensure kernel module loading and unloading is collected
-w /sbin/insmod -p x -k modules
-w /sbin/rmmod -p x -k modules
-w /sbin/modprobe -p x -k modules
-a always,exit -F arch=b64 -S init_module -S delete_module -k modules

# 4.1.18 Ensure the audit configuration is immutable
-e 2
`)

func k8sCloudInitArtifactsAuditdRulesBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsAuditdRules, nil
}

func k8sCloudInitArtifactsAuditdRules() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsAuditdRulesBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/auditd-rules", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsCisSh = []byte(`#!/bin/bash

assignRootPW() {
  if grep '^root:[!*]:' /etc/shadow; then
    SALT=$(openssl rand -base64 5)
    SECRET=$(openssl rand -base64 37)
    CMD="import crypt, getpass, pwd; print crypt.crypt('$SECRET', '\$6\$$SALT\$')"
    HASH=$(python -c "$CMD")

    echo 'root:'$HASH | /usr/sbin/chpasswd -e || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
  fi
}

assignFilePermissions() {
    FILES="
    auth.log
    alternatives.log
    cloud-init.log
    cloud-init-output.log
    daemon.log
    dpkg.log
    kern.log
    lastlog
    waagent.log
    syslog
    unattended-upgrades/unattended-upgrades.log
    unattended-upgrades/unattended-upgrades-dpkg.log
    azure-vnet-ipam.log
    azure-vnet-telemetry.log
    azure-cnimonitor.log
    azure-vnet.log
    kv-driver.log
    blobfuse-driver.log
    blobfuse-flexvol-installer.log
    landscape/sysinfo.log
    "
    for FILE in ${FILES}; do
        FILEPATH="/var/log/${FILE}"
        DIR=$(dirname "${FILEPATH}")
        mkdir -p ${DIR} || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
        touch ${FILEPATH} || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
        chmod 640 ${FILEPATH} || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
    done
    find /var/log -type f -perm '/o+r' -exec chmod 'g-wx,o-rwx' {} \;
    chmod 600 /etc/passwd- || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
    chmod 600 /etc/shadow- || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
    chmod 600 /etc/group- || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
    chmod 644 /etc/default/grub || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
    for filepath in /etc/crontab /etc/cron.hourly /etc/cron.daily /etc/cron.weekly /etc/cron.monthly /etc/cron.d; do
      chmod 0600 $filepath || exit $ERR_CIS_ASSIGN_FILE_PERMISSION
    done
}

setPWExpiration() {
  sed -i "s|PASS_MAX_DAYS||g" /etc/login.defs || exit $ERR_CIS_APPLY_PASSWORD_CONFIG
  grep 'PASS_MAX_DAYS' /etc/login.defs && exit $ERR_CIS_APPLY_PASSWORD_CONFIG
  sed -i "s|PASS_MIN_DAYS||g" /etc/login.defs || exit $ERR_CIS_APPLY_PASSWORD_CONFIG
  grep 'PASS_MIN_DAYS' /etc/login.defs && exit $ERR_CIS_APPLY_PASSWORD_CONFIG
  sed -i "s|INACTIVE=||g" /etc/default/useradd || exit $ERR_CIS_APPLY_PASSWORD_CONFIG
  grep 'INACTIVE=' /etc/default/useradd && exit $ERR_CIS_APPLY_PASSWORD_CONFIG
  echo 'PASS_MAX_DAYS 90' >> /etc/login.defs || exit $ERR_CIS_APPLY_PASSWORD_CONFIG
  grep 'PASS_MAX_DAYS 90' /etc/login.defs || exit $ERR_CIS_APPLY_PASSWORD_CONFIG
  echo 'PASS_MIN_DAYS 7' >> /etc/login.defs || exit $ERR_CIS_APPLY_PASSWORD_CONFIG
  grep 'PASS_MIN_DAYS 7' /etc/login.defs || exit $ERR_CIS_APPLY_PASSWORD_CONFIG
  echo 'INACTIVE=30' >> /etc/default/useradd || exit $ERR_CIS_APPLY_PASSWORD_CONFIG
  grep 'INACTIVE=30' /etc/default/useradd || exit $ERR_CIS_APPLY_PASSWORD_CONFIG
}

applyCIS() {
  setPWExpiration
  assignRootPW
  assignFilePermissions
}

applyCIS

#EOF
`)

func k8sCloudInitArtifactsCisShBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsCisSh, nil
}

func k8sCloudInitArtifactsCisSh() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsCisShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/cis.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsCse_configSh = []byte(`#!/bin/bash
NODE_INDEX=$(hostname | tail -c 2)
NODE_NAME=$(hostname)
if [[ $OS == $COREOS_OS_NAME ]]; then
    PRIVATE_IP=$(ip a show eth0 | grep -Po 'inet \K[\d.]+')
else
    PRIVATE_IP=$(hostname -I | cut -d' ' -f1)
fi
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
    if ! retrycmd_if_failure 120 5 25 systemctl enable $1; then
        echo "$1 could not be enabled by systemctl"
        return 1
    fi
}
systemctlDisableAndStop() {
    if ! systemctl_stop 100 5 30 $1; then
        echo "$1 could not be stopped"
        return 1
    fi
    if ! retrycmd_if_failure 120 5 25 systemctl disable $1; then
        echo "$1 could not be disabled by systemctl"
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

ensureAuditD() {
  if [[ "${AUDITD_ENABLED}" == true ]]; then
    systemctlEnableAndStart auditd || exit $ERR_SYSTEMCTL_START_FAIL
  else
    if apt list --installed | grep 'auditd'; then
      systemctlDisableAndStop auditd || exit $ERR_SYSTEMCTL_START_FAIL
    fi
  fi
}

generateAggregatedAPICerts() {
    AGGREGATED_API_CERTS_SETUP_FILE=/etc/kubernetes/generate-proxy-certs.sh
    wait_for_file 1200 1 $AGGREGATED_API_CERTS_SETUP_FILE || exit $ERR_FILE_WATCH_TIMEOUT
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
    # Perform the required JSON escaping for special characters \ and "
    SERVICE_PRINCIPAL_CLIENT_SECRET=${SERVICE_PRINCIPAL_CLIENT_SECRET//\\/\\\\}
    SERVICE_PRINCIPAL_CLIENT_SECRET=${SERVICE_PRINCIPAL_CLIENT_SECRET//\"/\\\"}
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
    if [[ -n "${MASTER_NODE}" ]]; then
        if [[ "${ENABLE_AGGREGATED_APIS}" = True ]]; then
            generateAggregatedAPICerts
        fi
    fi

    configureKubeletServerCert
}

configureCNI() {
    # needed for the iptables rules to work on bridges
    retrycmd_if_failure 120 5 25 modprobe br_netfilter || exit $ERR_MODPROBE_FAIL
    echo -n "br_netfilter" > /etc/modules-load.d/br_netfilter.conf
    configureCNIIPTables
    if [[ "${NETWORK_PLUGIN}" = "cilium" ]]; then
        systemctl enable sys-fs-bpf.mount
        systemctl restart sys-fs-bpf.mount
        REBOOTREQUIRED=true
    fi

    if [[ "${TARGET_ENVIRONMENT,,}" == "${AZURE_STACK_ENV}"  ]] && [[ "${NETWORK_PLUGIN}" = "azure" ]]; then
        # set environment to mas when using Azure CNI on Azure Stack
        # shellcheck disable=SC2002,SC2005
        echo $(cat "$CNI_CONFIG_DIR/10-azure.conflist" | jq '.plugins[0].ipam.environment = "mas"') > "$CNI_CONFIG_DIR/10-azure.conflist"
    fi
}

configureCNIIPTables() {
    if [[ "${NETWORK_PLUGIN}" = "azure" ]]; then
        mv $CNI_BIN_DIR/10-azure.conflist $CNI_CONFIG_DIR/
        chmod 600 $CNI_CONFIG_DIR/10-azure.conflist
        if [[ "${NETWORK_POLICY}" == "calico" ]]; then
          sed -i 's#"mode":"bridge"#"mode":"transparent"#g' $CNI_CONFIG_DIR/10-azure.conflist
        fi
        /sbin/ebtables -t nat --list
    fi
}

setupContainerd() {
    echo "Configuring cri-containerd..."
    mkdir -p "/etc/containerd"
    CRI_CONTAINERD_CONFIG="/etc/containerd/config.toml"
    {
        echo "subreaper = false"
        echo "oom_score = 0"
        echo "[plugins.cri]"
        echo "sandbox_image = \"$POD_INFRA_CONTAINER_SPEC\""
        echo "[plugins.cri.containerd.untrusted_workload_runtime]"
        echo "runtime_type = 'io.containerd.runtime.v1.linux'"
        if [[ "$CONTAINER_RUNTIME" == "kata-containers" ]]; then
            echo "runtime_engine = '/usr/bin/kata-runtime'"
        else
            echo "runtime_engine = '/usr/local/sbin/runc'"
        fi
        echo "[plugins.cri.containerd.default_runtime]"
        echo "runtime_type = 'io.containerd.runtime.v1.linux'"
        echo "runtime_engine = '/usr/local/sbin/runc'"
    } > "$CRI_CONTAINERD_CONFIG"
}

ensureContainerd() {
    setupContainerd
    echo "Enabling and starting cri-containerd service..."
    systemctlEnableAndStart containerd || exit $ERR_SYSTEMCTL_START_FAIL
}

ensureDocker() {
    DOCKER_SERVICE_EXEC_START_FILE=/etc/systemd/system/docker.service.d/exec_start.conf
    wait_for_file 1200 1 $DOCKER_SERVICE_EXEC_START_FILE || exit $ERR_FILE_WATCH_TIMEOUT
    usermod -aG docker ${ADMINUSER}
    DOCKER_MOUNT_FLAGS_SYSTEMD_FILE=/etc/systemd/system/docker.service.d/clear_mount_propagation_flags.conf
    if [[ $OS != $COREOS_OS_NAME ]]; then
        wait_for_file 1200 1 $DOCKER_MOUNT_FLAGS_SYSTEMD_FILE || exit $ERR_FILE_WATCH_TIMEOUT
    fi
    DOCKER_JSON_FILE=/etc/docker/daemon.json
    for i in $(seq 1 1200); do
        if [ -s $DOCKER_JSON_FILE ]; then
            jq '.' < $DOCKER_JSON_FILE && break
        fi
        if [ $i -eq 1200 ]; then
            exit $ERR_FILE_WATCH_TIMEOUT
        else
            sleep 1
        fi
    done
    systemctlEnableAndStart docker || exit $ERR_DOCKER_START_FAIL
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

ensureDHCPv6() {
    systemctlEnableAndStart dhcpv6 || exit $ERR_SYSTEMCTL_START_FAIL
}

ensureKubelet() {
    KUBELET_DEFAULT_FILE=/etc/default/kubelet
    wait_for_file 1200 1 $KUBELET_DEFAULT_FILE || exit $ERR_FILE_WATCH_TIMEOUT
    KUBECONFIG_FILE=/var/lib/kubelet/kubeconfig
    wait_for_file 1200 1 $KUBECONFIG_FILE || exit $ERR_FILE_WATCH_TIMEOUT
    KUBELET_RUNTIME_CONFIG_SCRIPT_FILE=/opt/azure/containers/kubelet.sh
    wait_for_file 1200 1 $KUBELET_RUNTIME_CONFIG_SCRIPT_FILE || exit $ERR_FILE_WATCH_TIMEOUT
    systemctlEnableAndStart kubelet || exit $ERR_KUBELET_START_FAIL
}

ensureLabelNodes() {
    LABEL_NODES_SCRIPT_FILE=/opt/azure/containers/label-nodes.sh
    wait_for_file 1200 1 $LABEL_NODES_SCRIPT_FILE || exit $ERR_FILE_WATCH_TIMEOUT
    LABEL_NODES_SYSTEMD_FILE=/etc/systemd/system/label-nodes.service
    wait_for_file 1200 1 $LABEL_NODES_SYSTEMD_FILE || exit $ERR_FILE_WATCH_TIMEOUT
    systemctlEnableAndStart label-nodes || exit $ERR_SYSTEMCTL_START_FAIL
}

ensureJournal() {
    {
        echo "Storage=persistent"
        echo "SystemMaxUse=1G"
        echo "RuntimeMaxUse=1G"
        echo "ForwardToSyslog=yes"
    } >> /etc/systemd/journald.conf
    systemctlEnableAndStart systemd-journald || exit $ERR_SYSTEMCTL_START_FAIL
}

ensureK8sControlPlane() {
    if $REBOOTREQUIRED || [ "$NO_OUTBOUND" = "true" ]; then
        return
    fi
    retrycmd_if_failure 120 5 25 $KUBECTL 2>/dev/null cluster-info || exit $ERR_K8S_RUNNING_TIMEOUT
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
        sed -i "s|<hostNet>|${CLUSTER_AUTOSCALER_MSI_HOST_NETWORK}|g" $CLUSTER_AUTOSCALER_ADDON_FILE
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
    sed -i "s|<vmssName>|$PRIMARY_SCALE_SET|g" $CLUSTER_AUTOSCALER_ADDON_FILE
}

configACIConnectorAddon() {
    ACI_CONNECTOR_CREDENTIALS=$(printf "{\"clientId\": \"%s\", \"clientSecret\": \"%s\", \"tenantId\": \"%s\", \"subscriptionId\": \"%s\", \"activeDirectoryEndpointUrl\": \"https://login.microsoftonline.com\",\"resourceManagerEndpointUrl\": \"https://management.azure.com/\", \"activeDirectoryGraphResourceId\": \"https://graph.windows.net/\", \"sqlManagementEndpointUrl\": \"https://management.core.windows.net:8443/\", \"galleryEndpointUrl\": \"https://gallery.azure.com/\", \"managementEndpointUrl\": \"https://management.core.windows.net/\"}" "$SERVICE_PRINCIPAL_CLIENT_ID" "$SERVICE_PRINCIPAL_CLIENT_SECRET" "$TENANT_ID" "$SUBSCRIPTION_ID" | base64 -w 0)

    openssl req -newkey rsa:4096 -new -nodes -x509 -days 3650 -keyout /etc/kubernetes/certs/aci-connector-key.pem -out /etc/kubernetes/certs/aci-connector-cert.pem -subj "/C=US/ST=CA/L=virtualkubelet/O=virtualkubelet/OU=virtualkubelet/CN=virtualkubelet"
    ACI_CONNECTOR_KEY=$(base64 /etc/kubernetes/certs/aci-connector-key.pem -w0)
    ACI_CONNECTOR_CERT=$(base64 /etc/kubernetes/certs/aci-connector-cert.pem -w0)

    ACI_CONNECTOR_ADDON_FILE=/etc/kubernetes/addons/aci-connector-deployment.yaml
    wait_for_file 1200 1 $ACI_CONNECTOR_ADDON_FILE || exit $ERR_FILE_WATCH_TIMEOUT
    sed -i "s|<creds>|$ACI_CONNECTOR_CREDENTIALS|g" $ACI_CONNECTOR_ADDON_FILE
    sed -i "s|<rgName>|$RESOURCE_GROUP|g" $ACI_CONNECTOR_ADDON_FILE
    sed -i "s|<cert>|$ACI_CONNECTOR_CERT|g" $ACI_CONNECTOR_ADDON_FILE
    sed -i "s|<key>|$ACI_CONNECTOR_KEY|g" $ACI_CONNECTOR_ADDON_FILE
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
    # only install the runtime since nvidia-docker2 has a hard dep on docker CE packages.
    # we will manually install nvidia-docker2
    rmmod nouveau
    echo blacklist nouveau >> /etc/modprobe.d/blacklist.conf
    retrycmd_if_failure_no_stats 120 5 25 update-initramfs -u || exit $ERR_GPU_DRIVERS_INSTALL_TIMEOUT
    wait_for_apt_locks
    retrycmd_if_failure 30 5 3600 apt-get -o Dpkg::Options::="--force-confold" install -y nvidia-container-runtime="${NVIDIA_CONTAINER_RUNTIME_VERSION}+docker18.09.2-1" || exit $ERR_GPU_DRIVERS_INSTALL_TIMEOUT
    tmpDir=$GPU_DEST/tmp
    (
      set -e -o pipefail
      cd "${tmpDir}"
      wait_for_apt_locks
      dpkg-deb -R ./nvidia-docker2*.deb "${tmpDir}/pkg" || exit $ERR_GPU_DRIVERS_INSTALL_TIMEOUT
      cp -r ${tmpDir}/pkg/usr/* /usr/ || exit $ERR_GPU_DRIVERS_INSTALL_TIMEOUT
    )
    rm -rf $GPU_DEST/tmp
    retrycmd_if_failure 120 5 25 pkill -SIGHUP dockerd || exit $ERR_GPU_DRIVERS_INSTALL_TIMEOUT
    mkdir -p $GPU_DEST/lib64 $GPU_DEST/overlay-workdir
    retrycmd_if_failure 120 5 25 mount -t overlay -o lowerdir=/usr/lib/x86_64-linux-gnu,upperdir=${GPU_DEST}/lib64,workdir=${GPU_DEST}/overlay-workdir none /usr/lib/x86_64-linux-gnu || exit $ERR_GPU_DRIVERS_INSTALL_TIMEOUT
    retrycmd_if_failure 3 1 600 sh $GPU_DEST/nvidia-drivers-$GPU_DV --silent --accept-license --no-drm --dkms --utility-prefix="${GPU_DEST}" --opengl-prefix="${GPU_DEST}" || exit $ERR_GPU_DRIVERS_START_FAIL
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
}
#EOF
`)

func k8sCloudInitArtifactsCse_configShBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsCse_configSh, nil
}

func k8sCloudInitArtifactsCse_configSh() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsCse_configShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/cse_config.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsCse_customcloudSh = []byte(`#!/bin/bash

ensureCertificates() {
    AZURESTACK_ENVIRONMENT_JSON_PATH="/etc/kubernetes/azurestackcloud.json"
    AZURESTACK_RESOURCE_MANAGER_ENDPOINT=$(jq .resourceManagerEndpoint $AZURESTACK_ENVIRONMENT_JSON_PATH | tr -d "\"")
    AZURESTACK_RESOURCE_METADATA_ENDPOINT="$AZURESTACK_RESOURCE_MANAGER_ENDPOINT/metadata/endpoints?api-version=2015-01-01"
    curl $AZURESTACK_RESOURCE_METADATA_ENDPOINT
    CURL_RETURNCODE=$?
    KUBE_CONTROLLER_MANAGER_FILE=/etc/kubernetes/manifests/kube-controller-manager.yaml
    if [ $CURL_RETURNCODE != 0 ]; then
        # Replace placeholder for ssl binding
        if [ -f $KUBE_CONTROLLER_MANAGER_FILE ]; then
            sed -i "s|<volumessl>|- name: ssl\n      hostPath:\n        path: \\/etc\\/ssl\\/certs|g" $KUBE_CONTROLLER_MANAGER_FILE
            sed -i "s|<volumeMountssl>|- name: ssl\n          mountPath: \\/etc\\/ssl\\/certs\n          readOnly: true|g" $KUBE_CONTROLLER_MANAGER_FILE
        fi

        # Copying the AzureStack root certificate to the appropriate store to be updated.
        AZURESTACK_ROOT_CERTIFICATE_SOURCE_PATH="/var/lib/waagent/Certificates.pem"
        AZURESTACK_ROOT_CERTIFICATE__DEST_PATH="/usr/local/share/ca-certificates/azsCertificate.crt"
        cp $AZURESTACK_ROOT_CERTIFICATE_SOURCE_PATH $AZURESTACK_ROOT_CERTIFICATE__DEST_PATH
        update-ca-certificates
    else
        if [ -f $KUBE_CONTROLLER_MANAGER_FILE ]; then
            # the ARM resource manager endpoint binding certificate is trusted, remove the placeholder for ssl binding
            sed -i "/<volumessl>/d" $KUBE_CONTROLLER_MANAGER_FILE
            sed -i "/<volumeMountssl>/d" $KUBE_CONTROLLER_MANAGER_FILE
        fi
    fi

    # ensureCertificates will be retried if the exit code is not 0
    curl $AZURESTACK_RESOURCE_METADATA_ENDPOINT
    exit $?
}

configureK8sCustomCloud() {
    export -f ensureCertificates
    retrycmd_if_failure 60 10 30 bash -c ensureCertificates
    set +x
    # When AUTHENTICATION_METHOD is client_certificate, the certificate is stored into key valut,
    # And SERVICE_PRINCIPAL_CLIENT_SECRET will be the following json payload with based64 encode
    #{
    #    "data": "$pfxAsBase64EncodedString",
    #    "dataType" :"pfx",
    #    "password": "$password"
    #}
    if [[ "${AUTHENTICATION_METHOD,,}" == "client_certificate" ]]; then
        SERVICE_PRINCIPAL_CLIENT_SECRET_DECODED=$(echo ${SERVICE_PRINCIPAL_CLIENT_SECRET} | base64 --decode)
        SERVICE_PRINCIPAL_CLIENT_SECRET_CERT=$(echo $SERVICE_PRINCIPAL_CLIENT_SECRET_DECODED | jq .data)
        SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD=$(echo $SERVICE_PRINCIPAL_CLIENT_SECRET_DECODED | jq .password)

        # trim the starting and ending "
        SERVICE_PRINCIPAL_CLIENT_SECRET_CERT=${SERVICE_PRINCIPAL_CLIENT_SECRET_CERT#"\""}
        SERVICE_PRINCIPAL_CLIENT_SECRET_CERT=${SERVICE_PRINCIPAL_CLIENT_SECRET_CERT%"\""}

        SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD=${SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD#"\""}
        SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD=${SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD%"\""}

        KUBERNETES_FILE_DIR=$(dirname "${AZURE_JSON_PATH}")
        K8S_CLIENT_CERT_PATH="${KUBERNETES_FILE_DIR}/k8s_auth_certificate.pfx"
        echo $SERVICE_PRINCIPAL_CLIENT_SECRET_CERT | base64 --decode > $K8S_CLIENT_CERT_PATH
        # shellcheck disable=SC2002,SC2005
        echo $(cat "${AZURE_JSON_PATH}" | \
            jq --arg K8S_CLIENT_CERT_PATH ${K8S_CLIENT_CERT_PATH} '. + {aadClientCertPath:($K8S_CLIENT_CERT_PATH)}' | \
            jq --arg SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD ${SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD} '. + {aadClientCertPassword:($SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD)}' |\
            jq 'del(.aadClientSecret)') > ${AZURE_JSON_PATH}
    fi

    if [[ "${IDENTITY_SYSTEM,,}" == "adfs"  ]]; then
        # update the tenent id for ADFS environment.
        # shellcheck disable=SC2002,SC2005
        echo $(cat "${AZURE_JSON_PATH}" | jq '.tenantId = "adfs"') > ${AZURE_JSON_PATH}
    fi

    # Decrease eth0 MTU to mitigate Azure Stack's NRP issue
    echo "iface eth0 inet dhcp" | sudo tee -a /etc/network/interfaces
    echo "    post-up /sbin/ifconfig eth0 mtu 1350" | sudo tee -a /etc/network/interfaces

    ifconfig eth0 mtu 1350

    set -x
}

configureAzureStackInterfaces() {
    set +x

    NETWORK_INTERFACES_FILE="/etc/kubernetes/network_interfaces.json"
    AZURE_CNI_CONFIG_FILE="/etc/kubernetes/interfaces.json"
    AZURESTACK_ENVIRONMENT_JSON_PATH="/etc/kubernetes/azurestackcloud.json"
    SERVICE_MANAGEMENT_ENDPOINT=$(jq -r '.serviceManagementEndpoint' ${AZURESTACK_ENVIRONMENT_JSON_PATH})
    ACTIVE_DIRECTORY_ENDPOINT=$(jq -r '.activeDirectoryEndpoint' ${AZURESTACK_ENVIRONMENT_JSON_PATH})
    RESOURCE_MANAGER_ENDPOINT=$(jq -r '.resourceManagerEndpoint' ${AZURESTACK_ENVIRONMENT_JSON_PATH})

    if [[ "${IDENTITY_SYSTEM,,}" == "adfs"  ]]; then
        TOKEN_URL="${ACTIVE_DIRECTORY_ENDPOINT}adfs/oauth2/token"
    else
        TOKEN_URL="${ACTIVE_DIRECTORY_ENDPOINT}${TENANT_ID}/oauth2/token"
    fi

    echo "Generating token for Azure Resource Manager"
    echo "------------------------------------------------------------------------"
    echo "Parameters"
    echo "------------------------------------------------------------------------"
    echo "SERVICE_PRINCIPAL_CLIENT_ID:     ..."
    echo "SERVICE_PRINCIPAL_CLIENT_SECRET: ..."
    echo "SERVICE_MANAGEMENT_ENDPOINT:     $SERVICE_MANAGEMENT_ENDPOINT"
    echo "ACTIVE_DIRECTORY_ENDPOINT:       $ACTIVE_DIRECTORY_ENDPOINT"
    echo "TENANT_ID:                       $TENANT_ID"
    echo "IDENTITY_SYSTEM:                 $IDENTITY_SYSTEM"
    echo "TOKEN_URL:                       $TOKEN_URL"
    echo "------------------------------------------------------------------------"

    TOKEN=$(curl -s --retry 5 --retry-delay 10 --max-time 60 -f -X POST \
        -H "Content-Type: application/x-www-form-urlencoded" \
        -d "grant_type=client_credentials" \
        -d "client_id=$SERVICE_PRINCIPAL_CLIENT_ID" \
        --data-urlencode "client_secret=$SERVICE_PRINCIPAL_CLIENT_SECRET" \
        --data-urlencode "resource=$SERVICE_MANAGEMENT_ENDPOINT" \
        ${TOKEN_URL} | jq '.access_token' | xargs)

    if [[ -z "$TOKEN" ]]; then
        echo "Error generating token for Azure Resource Manager"
        exit ${ERR_AZURE_STACK_GET_ARM_TOKEN}
    fi

    echo "Fetching network interface configuration for node"
    echo "------------------------------------------------------------------------"
    echo "Parameters"
    echo "------------------------------------------------------------------------"
    echo "RESOURCE_MANAGER_ENDPOINT: $RESOURCE_MANAGER_ENDPOINT"
    echo "SUBSCRIPTION_ID:           $SUBSCRIPTION_ID"
    echo "RESOURCE_GROUP:            $RESOURCE_GROUP"
    echo "NETWORK_API_VERSION:       $NETWORK_API_VERSION"
    echo "------------------------------------------------------------------------"

    curl -s --retry 5 --retry-delay 10 --max-time 60 -f -X GET \
        -H "Authorization: Bearer $TOKEN" \
        -H "Content-Type: application/json" \
        "${RESOURCE_MANAGER_ENDPOINT}subscriptions/$SUBSCRIPTION_ID/resourceGroups/$RESOURCE_GROUP/providers/Microsoft.Network/networkInterfaces?api-version=$NETWORK_API_VERSION" > ${NETWORK_INTERFACES_FILE}

    if [[ ! -s ${NETWORK_INTERFACES_FILE} ]]; then
        echo "Error fetching network interface configuration for node"
        exit ${ERR_AZURE_STACK_GET_NETWORK_CONFIGURATION}
    fi

    echo "Generating Azure CNI interface file"

    mapfile -t local_interfaces < <(cat /sys/class/net/*/address | tr -d : | sed 's/.*/\U&/g')

    SDN_INTERFACES=$(jq ".value | map(select(.properties.macAddress | inside(\"${local_interfaces[*]}\"))) | map(select((.properties.ipConfigurations | length) > 0))" ${NETWORK_INTERFACES_FILE})

    AZURE_CNI_CONFIG=$(echo ${SDN_INTERFACES} | jq "{Interfaces: [.[] | {MacAddress: .properties.macAddress, IsPrimary: .properties.primary, IPSubnets: [{Prefix: .properties.ipConfigurations[0].properties.subnet.id, IPAddresses: .properties.ipConfigurations | [.[] | {Address: .properties.privateIPAddress, IsPrimary: .properties.primary}]}]}]}")

    mapfile -t SUBNET_IDS < <(echo ${SDN_INTERFACES} | jq '[.[].properties.ipConfigurations[0].properties.subnet.id] | unique | .[]' -r)

    for SUBNET_ID in "${SUBNET_IDS[@]}"; do
        SUBNET_PREFIX=$(curl -s --retry 5 --retry-delay 10 --max-time 60 -f -X GET \
            -H "Authorization: Bearer $TOKEN" \
            -H "Content-Type: application/json" \
            "${RESOURCE_MANAGER_ENDPOINT}${SUBNET_ID:1}?api-version=$NETWORK_API_VERSION" | \
            jq '.properties.addressPrefix' -r)

        if [[ -z "$SUBNET_PREFIX" ]]; then
            echo "Error fetching the subnet address prefix for a subnet ID"
            exit ${ERR_AZURE_STACK_GET_SUBNET_PREFIX}
        fi

        # shellcheck disable=SC2001
        AZURE_CNI_CONFIG=$(echo ${AZURE_CNI_CONFIG} | sed "s|$SUBNET_ID|$SUBNET_PREFIX|g")
    done

    echo ${AZURE_CNI_CONFIG} > ${AZURE_CNI_CONFIG_FILE}

    chmod 0444 ${AZURE_CNI_CONFIG_FILE}

    set -x
}
#EOF
`)

func k8sCloudInitArtifactsCse_customcloudShBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsCse_customcloudSh, nil
}

func k8sCloudInitArtifactsCse_customcloudSh() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsCse_customcloudShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/cse_customcloud.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsCse_helpersSh = []byte(`#!/bin/bash

#ERR_SYSTEMCTL_ENABLE_FAIL=3 # Service could not be enabled by systemctl -- DEPRECATED
ERR_SYSTEMCTL_START_FAIL=4 # Service could not be started or enabled by systemctl
ERR_CLOUD_INIT_TIMEOUT=5 # Timeout waiting for cloud-init runcmd to complete
ERR_FILE_WATCH_TIMEOUT=6 # Timeout waiting for a file
ERR_HOLD_WALINUXAGENT=7 # Unable to place walinuxagent apt package on hold during install
ERR_RELEASE_HOLD_WALINUXAGENT=8 # Unable to release hold on walinuxagent apt package after install
ERR_APT_INSTALL_TIMEOUT=9 # Timeout installing required apt packages
ERR_ETCD_DATA_DIR_NOT_FOUND=10 # Etcd data dir not found
ERR_ETCD_RUNNING_TIMEOUT=11 # Timeout waiting for etcd to be accessible
ERR_ETCD_DOWNLOAD_TIMEOUT=12 # Timeout waiting for etcd to download
ERR_ETCD_VOL_MOUNT_FAIL=13 # Unable to mount etcd disk volume
ERR_ETCD_START_TIMEOUT=14 # Unable to start etcd runtime
ERR_ETCD_CONFIG_FAIL=15 # Unable to configure etcd cluster
ERR_DOCKER_INSTALL_TIMEOUT=20 # Timeout waiting for docker install
ERR_DOCKER_DOWNLOAD_TIMEOUT=21 # Timout waiting for docker download(s)
ERR_DOCKER_KEY_DOWNLOAD_TIMEOUT=22 # Timeout waiting to download docker repo key
ERR_DOCKER_APT_KEY_TIMEOUT=23 # Timeout waiting for docker apt-key
ERR_DOCKER_START_FAIL=24 # Docker could not be started by systemctl
ERR_MOBY_APT_LIST_TIMEOUT=25 # Timeout waiting for moby apt sources
ERR_MS_GPG_KEY_DOWNLOAD_TIMEOUT=26 # Timeout waiting for MS GPG key download
ERR_MOBY_INSTALL_TIMEOUT=27 # Timeout waiting for moby install
ERR_K8S_RUNNING_TIMEOUT=30 # Timeout waiting for k8s cluster to be healthy
ERR_K8S_DOWNLOAD_TIMEOUT=31 # Timeout waiting for Kubernetes download(s)
ERR_KUBECTL_NOT_FOUND=32 # kubectl client binary not found on local disk
ERR_IMG_DOWNLOAD_TIMEOUT=33 # Timeout waiting for img download
ERR_KUBELET_START_FAIL=34 # kubelet could not be started by systemctl
ERR_CONTAINER_IMG_PULL_TIMEOUT=35 # Timeout trying to pull a container image
ERR_CNI_DOWNLOAD_TIMEOUT=41 # Timeout waiting for CNI download(s)
ERR_MS_PROD_DEB_DOWNLOAD_TIMEOUT=42 # Timeout waiting for https://packages.microsoft.com/config/ubuntu/16.04/packages-microsoft-prod.deb
ERR_MS_PROD_DEB_PKG_ADD_FAIL=43 # Failed to add repo pkg file
#ERR_FLEXVOLUME_DOWNLOAD_TIMEOUT=44 # Failed to add repo pkg file -- DEPRECATED
ERR_SYSTEMD_INSTALL_FAIL=48 # Unable to install required systemd version
ERR_MODPROBE_FAIL=49 # Unable to load a kernel module using modprobe
ERR_OUTBOUND_CONN_FAIL=50 # Unable to establish outbound connection
ERR_KATA_KEY_DOWNLOAD_TIMEOUT=60 # Timeout waiting to download kata repo key
ERR_KATA_APT_KEY_TIMEOUT=61 # Timeout waiting for kata apt-key
ERR_KATA_INSTALL_TIMEOUT=62 # Timeout waiting for kata install
ERR_CONTAINERD_DOWNLOAD_TIMEOUT=70 # Timeout waiting for containerd download(s)
ERR_CUSTOM_SEARCH_DOMAINS_FAIL=80 # Unable to configure custom search domains
ERR_GPU_DRIVERS_START_FAIL=84 # nvidia-modprobe could not be started by systemctl
ERR_GPU_DRIVERS_INSTALL_TIMEOUT=85 # Timeout waiting for GPU drivers install
ERR_SGX_DRIVERS_INSTALL_TIMEOUT=90 # Timeout waiting for SGX prereqs to download
ERR_SGX_DRIVERS_START_FAIL=91 # Failed to execute SGX driver binary
ERR_APT_DAILY_TIMEOUT=98 # Timeout waiting for apt daily updates
ERR_APT_UPDATE_TIMEOUT=99 # Timeout waiting for apt-get update to complete
ERR_CSE_PROVISION_SCRIPT_NOT_READY_TIMEOUT=100 # Timeout waiting for cloud-init to place this (!) script on the vm
ERR_APT_DIST_UPGRADE_TIMEOUT=101 # Timeout waiting for apt-get dist-upgrade to complete
ERR_APT_PURGE_FAIL=102 # Error purging distro packages
ERR_SYSCTL_RELOAD=103 # Error reloading sysctl config
ERR_CIS_ASSIGN_ROOT_PW=111 # Error assigning root password in CIS enforcement
ERR_CIS_ASSIGN_FILE_PERMISSION=112 # Error assigning permission to a file in CIS enforcement
ERR_PACKER_COPY_FILE=113 # Error writing a file to disk during VHD CI
ERR_CIS_APPLY_PASSWORD_CONFIG=115 # Error applying CIS-recommended passwd configuration

ERR_VHD_BUILD_ERROR=125 # Reserved for VHD CI exit conditions

# Azure Stack specific errors
ERR_AZURE_STACK_GET_ARM_TOKEN=120 # Error generating a token to use with Azure Resource Manager
ERR_AZURE_STACK_GET_NETWORK_CONFIGURATION=121 # Error fetching the network configuration for the node
ERR_AZURE_STACK_GET_SUBNET_PREFIX=122 # Error fetching the subnet address prefix for a subnet ID

OS=$(sort -r /etc/*-release | gawk 'match($0, /^(ID_LIKE=(coreos)|ID=(.*))$/, a) { print toupper(a[2] a[3]); exit }')
UBUNTU_OS_NAME="UBUNTU"
RHEL_OS_NAME="RHEL"
COREOS_OS_NAME="COREOS"
KUBECTL=/usr/local/bin/kubectl
DOCKER=/usr/bin/docker
GPU_DV=418.40.04
GPU_DEST=/usr/local/nvidia
NVIDIA_DOCKER_VERSION=2.0.3
DOCKER_VERSION=1.13.1-1
NVIDIA_CONTAINER_RUNTIME_VERSION=2.0.0

aptmarkWALinuxAgent() {
    wait_for_apt_locks
    retrycmd_if_failure 120 5 25 apt-mark $1 walinuxagent || \
    if [[ "$1" == "hold" ]]; then
        exit $ERR_HOLD_WALINUXAGENT
    elif [[ "$1" == "unhold" ]]; then
        exit $ERR_RELEASE_HOLD_WALINUXAGENT
    fi
}

retrycmd_if_failure() {
    retries=$1; wait_sleep=$2; timeout=$3; shift && shift && shift
    for i in $(seq 1 $retries); do
        timeout $timeout ${@} && break || \
        if [ $i -eq $retries ]; then
            echo Executed \"$@\" $i times;
            return 1
        else
            sleep $wait_sleep
        fi
    done
    echo Executed \"$@\" $i times;
}
retrycmd_if_failure_no_stats() {
    retries=$1; wait_sleep=$2; timeout=$3; shift && shift && shift
    for i in $(seq 1 $retries); do
        timeout $timeout ${@} && break || \
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
        tar -tzf $tarball && break || \
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
        $filepath $validation_args && break || \
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
    for i in $(seq 1 $retries); do
        grep -Fq '#EOF' $filepath && break
        if [ $i -eq $retries ]; then
            return 1
        else
            sleep $wait_sleep
        fi
    done
    sed -i "/#EOF/d" $filepath
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
        ! (apt-get update 2>&1 | tee $apt_update_output | grep -E "^([WE]:.*)|([eE]rr.*)$") && \
        cat $apt_update_output && break || \
        cat $apt_update_output
        if [ $i -eq $retries ]; then
            return 1
        else sleep 5
        fi
    done
    echo Executed apt-get update $i times
    wait_for_apt_locks
}
apt_get_install() {
    retries=$1; wait_sleep=$2; timeout=$3; shift && shift && shift
    for i in $(seq 1 $retries); do
        wait_for_apt_locks
        export DEBIAN_FRONTEND=noninteractive
        dpkg --configure -a --force-confdef
        apt-get install -o Dpkg::Options::="--force-confold" --no-install-recommends -y ${@} && break || \
        if [ $i -eq $retries ]; then
            return 1
        else
            sleep $wait_sleep
            apt_get_update
        fi
    done
    echo Executed apt-get install --no-install-recommends -y \"$@\" $i times;
    wait_for_apt_locks
}
apt_get_purge() {
    retries=$1; wait_sleep=$2; timeout=$3; shift && shift && shift
    for i in $(seq 1 $retries); do
        wait_for_apt_locks
        export DEBIAN_FRONTEND=noninteractive
        dpkg --configure -a --force-confdef
        apt-get purge -o Dpkg::Options::="--force-confold" -y ${@} && break || \
        if [ $i -eq $retries ]; then
            return 1
        else
            sleep $wait_sleep
        fi
    done
    echo Executed apt-get purge -y \"$@\" $i times;
    wait_for_apt_locks
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
  wait_for_apt_locks
}
systemctl_restart() {
    retries=$1; wait_sleep=$2; timeout=$3 svcname=$4
    for i in $(seq 1 $retries); do
        timeout $timeout systemctl daemon-reload
        timeout $timeout systemctl restart $svcname && break || \
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
        timeout $timeout systemctl stop $svcname && break || \
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
        timeout $timeout sysctl --system && break || \
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
`)

func k8sCloudInitArtifactsCse_helpersShBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsCse_helpersSh, nil
}

func k8sCloudInitArtifactsCse_helpersSh() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsCse_helpersShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/cse_helpers.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsCse_installSh = []byte(`#!/bin/bash

CC_SERVICE_IN_TMP=/opt/azure/containers/cc-proxy.service.in
CC_SOCKET_IN_TMP=/opt/azure/containers/cc-proxy.socket.in
CNI_CONFIG_DIR="/etc/cni/net.d"
CNI_BIN_DIR="/opt/cni/bin"
CNI_DOWNLOADS_DIR="/opt/cni/downloads"
CONTAINERD_DOWNLOADS_DIR="/opt/containerd/downloads"
UBUNTU_RELEASE=$(lsb_release -r -s)

removeEtcd() {
    if [[ $OS == $COREOS_OS_NAME ]]; then
        rm -rf /opt/bin/etcd
    else
        rm -rf /usr/bin/etcd
    fi
}

removeMoby() {
    apt-get purge -y moby-engine moby-cli
}

installEtcd() {
    CURRENT_VERSION=$(etcd --version | grep "etcd Version" | cut -d ":" -f 2 | tr -d '[:space:]')
    if [[ "$CURRENT_VERSION" == "${ETCD_VERSION}" ]]; then
        echo "etcd version ${ETCD_VERSION} is already installed, skipping download"
    else
        retrycmd_get_tarball 120 5 /tmp/etcd-v${ETCD_VERSION}-linux-amd64.tar.gz ${ETCD_DOWNLOAD_URL}/etcd-v${ETCD_VERSION}-linux-amd64.tar.gz || exit $ERR_ETCD_DOWNLOAD_TIMEOUT
        removeEtcd
        if [[ $OS == $COREOS_OS_NAME ]]; then
            tar -xzvf /tmp/etcd-v${ETCD_VERSION}-linux-amd64.tar.gz -C /opt/bin/ --strip-components=1 || exit $ERR_ETCD_DOWNLOAD_TIMEOUT
        else
            tar -xzvf /tmp/etcd-v${ETCD_VERSION}-linux-amd64.tar.gz -C /usr/bin/ --strip-components=1 || exit $ERR_ETCD_DOWNLOAD_TIMEOUT
        fi
    fi
}

installDeps() {
    retrycmd_if_failure_no_stats 120 5 25 curl -fsSL https://packages.microsoft.com/config/ubuntu/${UBUNTU_RELEASE}/packages-microsoft-prod.deb > /tmp/packages-microsoft-prod.deb || exit $ERR_MS_PROD_DEB_DOWNLOAD_TIMEOUT
    retrycmd_if_failure 60 5 10 dpkg -i /tmp/packages-microsoft-prod.deb || exit $ERR_MS_PROD_DEB_PKG_ADD_FAIL
    aptmarkWALinuxAgent hold
    apt_get_update || exit $ERR_APT_UPDATE_TIMEOUT
    apt_get_dist_upgrade || exit $ERR_APT_DIST_UPGRADE_TIMEOUT
    for apt_package in apache2-utils apt-transport-https auditd blobfuse ca-certificates ceph-common cgroup-lite cifs-utils conntrack cracklib-runtime ebtables ethtool fuse git glusterfs-client htop iftop init-system-helpers iotop iproute2 ipset iptables jq libpam-pwquality libpwquality-tools mount nfs-common pigz socat sysstat util-linux xz-utils zip; do
      if ! apt_get_install 30 1 600 $apt_package; then
        journalctl --no-pager -u $apt_package
        exit $ERR_APT_INSTALL_TIMEOUT
      fi
    done
}

installGPUDrivers() {
    mkdir -p $GPU_DEST/tmp
    retrycmd_if_failure_no_stats 120 5 25 curl -fsSL https://nvidia.github.io/nvidia-docker/gpgkey > $GPU_DEST/tmp/aptnvidia.gpg || exit $ERR_GPU_DRIVERS_INSTALL_TIMEOUT
    wait_for_apt_locks
    retrycmd_if_failure 120 5 25 apt-key add $GPU_DEST/tmp/aptnvidia.gpg || exit $ERR_GPU_DRIVERS_INSTALL_TIMEOUT
    wait_for_apt_locks
    retrycmd_if_failure_no_stats 120 5 25 curl -fsSL https://nvidia.github.io/nvidia-docker/ubuntu${UBUNTU_RELEASE}/nvidia-docker.list > $GPU_DEST/tmp/nvidia-docker.list || exit  $ERR_GPU_DRIVERS_INSTALL_TIMEOUT
    wait_for_apt_locks
    retrycmd_if_failure_no_stats 120 5 25 cat $GPU_DEST/tmp/nvidia-docker.list > /etc/apt/sources.list.d/nvidia-docker.list || exit  $ERR_GPU_DRIVERS_INSTALL_TIMEOUT
    apt_get_update
    retrycmd_if_failure 30 5 3600 apt-get install -y linux-headers-$(uname -r) gcc make dkms || exit $ERR_GPU_DRIVERS_INSTALL_TIMEOUT
    retrycmd_if_failure 30 5 60 curl -fLS https://us.download.nvidia.com/tesla/$GPU_DV/NVIDIA-Linux-x86_64-${GPU_DV}.run -o ${GPU_DEST}/nvidia-drivers-${GPU_DV} || exit $ERR_GPU_DRIVERS_INSTALL_TIMEOUT
    tmpDir=$GPU_DEST/tmp
    if ! (
      set -e -o pipefail
      cd "${tmpDir}"
      retrycmd_if_failure 30 5 3600 apt-get download nvidia-docker2="${NVIDIA_DOCKER_VERSION}+docker18.09.2-1" || exit $ERR_GPU_DRIVERS_INSTALL_TIMEOUT
    ); then
      exit $ERR_GPU_DRIVERS_INSTALL_TIMEOUT
    fi
}

installSGXDrivers() {
    echo "Installing SGX driver"
    local VERSION
    VERSION=$(grep DISTRIB_RELEASE /etc/*-release| cut -f 2 -d "=")
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

    local SGX_DRIVER
    SGX_DRIVER=$(basename $SGX_DRIVER_URL)
    local OE_DIR=/opt/azure/containers/oe
    mkdir -p ${OE_DIR}

    retrycmd_if_failure 120 5 25 curl -fsSL ${SGX_DRIVER_URL} -o ${OE_DIR}/${SGX_DRIVER} || exit $ERR_SGX_DRIVERS_INSTALL_TIMEOUT
    chmod a+x ${OE_DIR}/${SGX_DRIVER}
    ${OE_DIR}/${SGX_DRIVER} || exit $ERR_SGX_DRIVERS_START_FAIL
}

installContainerRuntime() {
    if [[ "$CONTAINER_RUNTIME" == "docker" ]]; then
        installMoby
    fi
}

installMoby() {
    CURRENT_VERSION=$(dockerd --version | grep "Docker version" | cut -d "," -f 1 | cut -d " " -f 3)
    if [[ "$CURRENT_VERSION" == "${MOBY_VERSION}" ]]; then
        echo "dockerd $MOBY_VERSION is already installed, skipping Moby download"
    else
        removeMoby
        retrycmd_if_failure_no_stats 120 5 25 curl https://packages.microsoft.com/config/ubuntu/${UBUNTU_RELEASE}/prod.list > /tmp/microsoft-prod.list || exit $ERR_MOBY_APT_LIST_TIMEOUT
        retrycmd_if_failure 10 5 10 cp /tmp/microsoft-prod.list /etc/apt/sources.list.d/ || exit $ERR_MOBY_APT_LIST_TIMEOUT
        retrycmd_if_failure_no_stats 120 5 25 curl https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > /tmp/microsoft.gpg || exit $ERR_MS_GPG_KEY_DOWNLOAD_TIMEOUT
        retrycmd_if_failure 10 5 10 cp /tmp/microsoft.gpg /etc/apt/trusted.gpg.d/ || exit $ERR_MS_GPG_KEY_DOWNLOAD_TIMEOUT
        apt_get_update || exit $ERR_APT_UPDATE_TIMEOUT
        MOBY_CLI=${MOBY_VERSION}
        if [[ "${MOBY_CLI}" == "3.0.4" ]]; then
            MOBY_CLI="3.0.3"
        fi
        apt_get_install 20 30 120 moby-engine=${MOBY_VERSION} moby-cli=${MOBY_CLI} --allow-downgrades || exit $ERR_MOBY_INSTALL_TIMEOUT
    fi
}

installKataContainersRuntime() {
    # TODO incorporate this into packer CI so that it is pre-baked into the VHD image
    echo "Adding Kata Containers repository key..."
    ARCH=$(arch)
    BRANCH=stable-1.7
    KATA_RELEASE_KEY_TMP=/tmp/kata-containers-release.key
    KATA_URL=http://download.opensuse.org/repositories/home:/katacontainers:/releases:/${ARCH}:/${BRANCH}/xUbuntu_${UBUNTU_RELEASE}/Release.key
    retrycmd_if_failure_no_stats 120 5 25 curl -fsSL $KATA_URL > $KATA_RELEASE_KEY_TMP || exit $ERR_KATA_KEY_DOWNLOAD_TIMEOUT
    wait_for_apt_locks
    retrycmd_if_failure 30 5 30 apt-key add $KATA_RELEASE_KEY_TMP || exit $ERR_KATA_APT_KEY_TIMEOUT
    echo "Adding Kata Containers repository..."
    echo "deb http://download.opensuse.org/repositories/home:/katacontainers:/releases:/${ARCH}:/${BRANCH}/xUbuntu_${UBUNTU_RELEASE}/ /" > /etc/apt/sources.list.d/kata-containers.list
    echo "Installing Kata Containers runtime..."
    apt_get_update || exit $ERR_APT_UPDATE_TIMEOUT
    apt_get_install 120 5 25 kata-runtime || exit $ERR_KATA_INSTALL_TIMEOUT
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
    if [[ -n "${PRIVATE_AZURE_REGISTRY_SERVER:-}" ]]; then
        $CLI_TOOL login -u $SERVICE_PRINCIPAL_CLIENT_ID -p $SERVICE_PRINCIPAL_CLIENT_SECRET $PRIVATE_AZURE_REGISTRY_SERVER
    fi
    retrycmd_if_failure 60 1 1200 $CLI_TOOL pull $DOCKER_IMAGE_URL || exit $ERR_CONTAINER_IMG_PULL_TIMEOUT
}

cleanUpContainerImages() {
    # TODO remove all unused container images at runtime
    docker rmi $(docker images --format '{{.Repository}}:{{.Tag}}' | grep -v "${KUBERNETES_VERSION}$" | grep 'hyperkube') &
    docker rmi $(docker images --format '{{.Repository}}:{{.Tag}}' | grep -v "${KUBERNETES_VERSION}$" | grep 'cloud-controller-manager') &
    if [ "$IS_HOSTED_MASTER" = "false" ]; then
        echo "Cleaning up AKS container images, not an AKS cluster"
        docker rmi $(docker images --format '{{.Repository}}:{{.Tag}}' | grep 'hcp-tunnel-front') &
        docker rmi $(docker images --format '{{.Repository}}:{{.Tag}}' | grep 'kube-svc-redirect') &
        docker rmi $(docker images --format '{{.Repository}}:{{.Tag}}' | grep 'nginx') &
    fi

    # TODO: remove once ACR is available on Azure Stack
    docker rmi registry:2 &
}

cleanUpGPUDrivers() {
    rm -Rf $GPU_DEST
    rm -f /etc/apt/sources.list.d/nvidia-docker.list
}

cleanUpContainerd() {
    rm -Rf $CONTAINERD_DOWNLOADS_DIR
}

overrideNetworkConfig() {
    CONFIG_FILEPATH="/etc/cloud/cloud.cfg.d/80_azure_net_config.cfg"
    touch ${CONFIG_FILEPATH}
    cat << EOF >> ${CONFIG_FILEPATH}
datasource:
    Azure:
        apply_network_config: false
EOF
}
#EOF
`)

func k8sCloudInitArtifactsCse_installShBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsCse_installSh, nil
}

func k8sCloudInitArtifactsCse_installSh() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsCse_installShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/cse_install.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsCse_mainSh = []byte(`#!/bin/bash
ERR_FILE_WATCH_TIMEOUT=6 # Timeout waiting for a file
set -x
echo $(date),$(hostname), startcustomscript>>/opt/m
AZURE_STACK_ENV="azurestackcloud"

script_lib=/opt/azure/containers/provision_source.sh
for i in $(seq 1 3600); do
    if [ -s $script_lib ]; then
        grep -Fq '#HELPERSEOF' $script_lib && break
    fi
    if [ $i -eq 3600 ]; then
        exit $ERR_FILE_WATCH_TIMEOUT
    else
        sleep 1
    fi
done
sed -i "/#HELPERSEOF/d" $script_lib
source $script_lib

install_script=/opt/azure/containers/provision_installs.sh
wait_for_file 3600 1 $install_script || exit $ERR_FILE_WATCH_TIMEOUT
source $install_script

config_script=/opt/azure/containers/provision_configs.sh
wait_for_file 3600 1 $config_script || exit $ERR_FILE_WATCH_TIMEOUT
source $config_script

if [[ "${TARGET_ENVIRONMENT,,}" == "${AZURE_STACK_ENV}"  ]]; then
    config_script_custom_cloud=/opt/azure/containers/provision_configs_custom_cloud.sh
    wait_for_file 3600 1 $config_script_custom_cloud || exit $ERR_FILE_WATCH_TIMEOUT
    source $config_script_custom_cloud
fi

CUSTOM_SEARCH_DOMAIN_SCRIPT=/opt/azure/containers/setup-custom-search-domains.sh

set +x
ETCD_PEER_CERT=$(echo ${ETCD_PEER_CERTIFICATES} | cut -d'[' -f 2 | cut -d']' -f 1 | cut -d',' -f $((${NODE_INDEX}+1)))
ETCD_PEER_KEY=$(echo ${ETCD_PEER_PRIVATE_KEYS} | cut -d'[' -f 2 | cut -d']' -f 1 | cut -d',' -f $((${NODE_INDEX}+1)))
set -x

if [[ $OS == $COREOS_OS_NAME ]]; then
    echo "Changing default kubectl bin location"
    KUBECTL=/opt/kubectl
fi

if [ -f /var/run/reboot-required ]; then
    REBOOTREQUIRED=true
else
    REBOOTREQUIRED=false
fi

if [ -f /var/log.vhd/azure/golden-image-install.complete ]; then
    echo "detected golden image pre-install"
    FULL_INSTALL_REQUIRED=false
    rm -rf /home/packer
    deluser packer
    groupdel packer
else
    FULL_INSTALL_REQUIRED=true
fi

if [[ $OS == $UBUNTU_OS_NAME ]] && [ "$FULL_INSTALL_REQUIRED" = "true" ]; then
    installDeps
else
    echo "Golden image; skipping dependencies installation"
fi
ensureAuditD

if [[ -n "${MASTER_NODE}" ]] && [[ -z "${COSMOS_URI}" ]]; then
    installEtcd
fi

if [[ $OS != $COREOS_OS_NAME ]]; then
    installContainerRuntime
fi
installNetworkPlugin
if [[ "$CONTAINER_RUNTIME" == "kata-containers" ]] || [[ "$CONTAINER_RUNTIME" == "containerd" ]]; then
    installContainerd
else
    cleanUpContainerd
fi
if [[ "${GPU_NODE}" = true ]]; then
    if $FULL_INSTALL_REQUIRED; then
        installGPUDrivers
    fi
    ensureGPUDrivers
else
    cleanUpGPUDrivers
fi
installKubeletAndKubectl
if [[ $OS != $COREOS_OS_NAME ]]; then
    ensureRPC
fi
createKubeManifestDir
if [[ "${SGX_NODE}" = true ]]; then
    installSGXDrivers
fi

# create etcd user if we are configured for etcd
if [[ -n "${MASTER_NODE}" ]] && [[ -z "${COSMOS_URI}" ]]; then
  configureEtcdUser
fi

if [[ -n "${MASTER_NODE}" ]]; then
  # this step configures all certs
  # both configs etcd/cosmos
  configureSecrets
fi
# configure etcd if we are configured for etcd
if [[ -n "${MASTER_NODE}" ]] && [[ -z "${COSMOS_URI}" ]]; then
    configureEtcd
else
    removeEtcd
fi


if [ -f $CUSTOM_SEARCH_DOMAIN_SCRIPT ]; then
    $CUSTOM_SEARCH_DOMAIN_SCRIPT > /opt/azure/containers/setup-custom-search-domain.log 2>&1 || exit $ERR_CUSTOM_SEARCH_DOMAINS_FAIL
fi

if [[ "$CONTAINER_RUNTIME" == "docker" ]]; then
    ensureDocker
elif [[ "$CONTAINER_RUNTIME" == "kata-containers" ]]; then
    if grep -q vmx /proc/cpuinfo; then
        installKataContainersRuntime
    fi
fi

configureK8s

if [[ "${TARGET_ENVIRONMENT,,}" == "${AZURE_STACK_ENV,,}"  ]]; then
    configureK8sCustomCloud
    if [[ "${NETWORK_PLUGIN,,}" = "azure" ]]; then
        configureAzureStackInterfaces
    fi
fi

configureCNI

if [[ -n "${MASTER_NODE}" ]]; then
    configAddons
fi

if [[ "$CONTAINER_RUNTIME" == "kata-containers" ]] || [[ "$CONTAINER_RUNTIME" == "containerd" ]]; then
    ensureContainerd
fi

if [[ -n "${MASTER_NODE}" && "${KMS_PROVIDER_VAULT_NAME}" != "" ]]; then
    ensureKMS
fi

# configure and enable dhcpv6 for dual stack feature
if [ "$IS_IPV6_DUALSTACK_FEATURE_ENABLED" = "true" ]; then
    dhcpv6_systemd_service=/etc/systemd/system/dhcpv6.service
    dhcpv6_configuration_script=/opt/azure/containers/enable-dhcpv6.sh
    wait_for_file 3600 1 $dhcpv6_systemd_service || exit $ERR_FILE_WATCH_TIMEOUT
    wait_for_file 3600 1 $dhcpv6_configuration_script || exit $ERR_FILE_WATCH_TIMEOUT
    ensureDHCPv6

    retrycmd_if_failure 120 5 25 modprobe ip6_tables || exit $ERR_MODPROBE_FAIL
fi

ensureKubelet
ensureJournal

if [[ -n "${MASTER_NODE}" ]]; then
    if version_gte ${KUBERNETES_VERSION} 1.16; then
      ensureLabelNodes
    fi
    writeKubeConfig
    if [[ -z "${COSMOS_URI}" ]]; then
      ensureEtcd
    fi
    ensureK8sControlPlane
fi

if $FULL_INSTALL_REQUIRED; then
    if [[ $OS == $UBUNTU_OS_NAME ]]; then
        # mitigation for bug https://bugs.launchpad.net/ubuntu/+source/linux/+bug/1676635
        echo 2dd1ce17-079e-403c-b352-a1921ee207ee > /sys/bus/vmbus/drivers/hv_util/unbind
        sed -i "13i\echo 2dd1ce17-079e-403c-b352-a1921ee207ee > /sys/bus/vmbus/drivers/hv_util/unbind\n" /etc/rc.local
    fi
fi

echo "Custom script finished successfully"

echo $(date),$(hostname), endcustomscript>>/opt/m
mkdir -p /opt/azure/containers && touch /opt/azure/containers/provision.complete
ps auxfww > /opt/azure/provision-ps.log &

if ! $FULL_INSTALL_REQUIRED; then
  cleanUpContainerImages
fi

if [[ "${TARGET_ENVIRONMENT,,}" != "${AZURE_STACK_ENV}"  ]]; then
    # TODO: remove once ACR is available on Azure Stack
    apt_get_purge 20 30 120 apache2-utils || exit $ERR_APT_PURGE_FAIL
fi

if $REBOOTREQUIRED; then
  echo 'reboot required, rebooting node in 1 minute'
  /bin/bash -c "shutdown -r 1 &"
  if [[ $OS == $UBUNTU_OS_NAME ]]; then
      aptmarkWALinuxAgent unhold &
  fi
else
  if [[ $OS == $UBUNTU_OS_NAME ]]; then
      /usr/lib/apt/apt.systemd.daily &
      aptmarkWALinuxAgent unhold &
  fi
fi
#EOF
`)

func k8sCloudInitArtifactsCse_mainShBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsCse_mainSh, nil
}

func k8sCloudInitArtifactsCse_mainSh() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsCse_mainShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/cse_main.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsDefaultGrub = []byte(`# If you change this file, run 'update-grub' afterwards to update
# /boot/grub/grub.cfg.
# For full documentation of the options in this file, see:
#   info -f grub -n 'Simple configuration'

GRUB_DEFAULT=0
GRUB_HIDDEN_TIMEOUT=0
GRUB_HIDDEN_TIMEOUT_QUIET=true
GRUB_TIMEOUT=0
GRUB_DISTRIBUTOR=` + "`" + `lsb_release -i -s 2> /dev/null || echo Debian` + "`" + `
GRUB_CMDLINE_LINUX_DEFAULT="console=tty1 console=ttyS0 earlyprintk=ttyS0 rootdelay=300"
# 4.1.3 Ensure auditing for processes that start prior to auditd is enabled
GRUB_CMDLINE_LINUX="audit=1 ipv6.disable=1"
`)

func k8sCloudInitArtifactsDefaultGrubBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsDefaultGrub, nil
}

func k8sCloudInitArtifactsDefaultGrub() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsDefaultGrubBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/default-grub", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsDhcpv6Service = []byte(`[Unit]
Description=enabledhcpv6
After=network-online.target

[Service]
Type=oneshot
ExecStart=/opt/azure/containers/enable-dhcpv6.sh

[Install]
WantedBy=multi-user.target
#EOF
`)

func k8sCloudInitArtifactsDhcpv6ServiceBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsDhcpv6Service, nil
}

func k8sCloudInitArtifactsDhcpv6Service() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsDhcpv6ServiceBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/dhcpv6.service", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsDockerMonitorService = []byte(`[Unit]
Description=a script that checks docker health and restarts if needed
After=docker.service
[Service]
Restart=always
RestartSec=10
RemainAfterExit=yes
ExecStart=/usr/local/bin/health-monitor.sh container-runtime
#EOF
`)

func k8sCloudInitArtifactsDockerMonitorServiceBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsDockerMonitorService, nil
}

func k8sCloudInitArtifactsDockerMonitorService() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsDockerMonitorServiceBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/docker-monitor.service", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsDockerMonitorTimer = []byte(`[Unit]
Description=a timer that delays docker-monitor from starting too soon after boot
[Timer]
OnBootSec=30min
[Install]
WantedBy=multi-user.target
#EOF
`)

func k8sCloudInitArtifactsDockerMonitorTimerBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsDockerMonitorTimer, nil
}

func k8sCloudInitArtifactsDockerMonitorTimer() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsDockerMonitorTimerBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/docker-monitor.timer", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsDocker_clear_mount_propagation_flagsConf = []byte(`[Service]
MountFlags=shared
#EOF
`)

func k8sCloudInitArtifactsDocker_clear_mount_propagation_flagsConfBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsDocker_clear_mount_propagation_flagsConf, nil
}

func k8sCloudInitArtifactsDocker_clear_mount_propagation_flagsConf() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsDocker_clear_mount_propagation_flagsConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/docker_clear_mount_propagation_flags.conf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsEnableDhcpv6Sh = []byte(`#!/usr/bin/env bash

set -e
set -o pipefail
set -u

DHCLIENT6_CONF_FILE=/etc/dhcp/dhclient6.conf
CLOUD_INIT_CFG=/etc/network/interfaces.d/50-cloud-init.cfg

read -r -d '' NETWORK_CONFIGURATION << EOC || true
iface eth0 inet6 auto
    up sleep 5
    up dhclient -1 -6 -cf /etc/dhcp/dhclient6.conf -lf /var/lib/dhcp/dhclient6.eth0.leases -v eth0 || true
EOC

add_if_not_exists() {
    grep -qxF "${1}" "${2}" || echo "${1}" >> "${2}"
}

echo "Configuring dhcpv6 ..."

touch /etc/dhcp/dhclient6.conf && add_if_not_exists "timeout 10;" ${DHCLIENT6_CONF_FILE} && \
    add_if_not_exists "${NETWORK_CONFIGURATION}" ${CLOUD_INIT_CFG} && \
    sudo ifdown eth0 && sudo ifup eth0

echo "Configuration complete"
#EOF
`)

func k8sCloudInitArtifactsEnableDhcpv6ShBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsEnableDhcpv6Sh, nil
}

func k8sCloudInitArtifactsEnableDhcpv6Sh() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsEnableDhcpv6ShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/enable-dhcpv6.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsEtcIssue = []byte(`
Authorized uses only. All activity may be monitored and reported.
`)

func k8sCloudInitArtifactsEtcIssueBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsEtcIssue, nil
}

func k8sCloudInitArtifactsEtcIssue() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsEtcIssueBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/etc-issue", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsEtcIssueNet = []byte(`
Authorized uses only. All activity may be monitored and reported.
`)

func k8sCloudInitArtifactsEtcIssueNetBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsEtcIssueNet, nil
}

func k8sCloudInitArtifactsEtcIssueNet() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsEtcIssueNetBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/etc-issue.net", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsEtcdService = []byte(`[Unit]
Description=etcd - highly-available key value store
Documentation=https://github.com/coreos/etcd
Documentation=man:etcd
After=network.target
Wants=network-online.target
RequiresMountsFor=/var/lib/etcddisk
[Service]
Environment=DAEMON_ARGS=
Environment=ETCD_NAME=%H
Environment=ETCD_DATA_DIR=
EnvironmentFile=-/etc/default/%p
Type=notify
User=etcd
PermissionsStartOnly=true
ExecStart=/usr/bin/etcd $DAEMON_ARGS
Restart=always
[Install]
WantedBy=multi-user.target`)

func k8sCloudInitArtifactsEtcdServiceBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsEtcdService, nil
}

func k8sCloudInitArtifactsEtcdService() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsEtcdServiceBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/etcd.service", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsGenerateproxycertsSh = []byte(`#!/bin/bash

source /opt/azure/containers/provision_source.sh

PROXY_CA_KEY="${PROXY_CA_KEY:=/tmp/proxy-client-ca.key}"
PROXY_CRT="${PROXY_CRT:=/tmp/proxy-client-ca.crt}"
PROXY_CLIENT_KEY="${PROXY_CLIENT_KEY:=/tmp/proxy-client.key}"
PROXY_CLIENT_CSR="${PROXY_CLIENT_CSR:=/tmp/proxy-client.csr}"
PROXY_CLIENT_CRT="${PROXY_CLIENT_CRT:=/tmp/proxy-client.crt}"
ETCD_REQUESTHEADER_CLIENT_CA="${ETCD_REQUESTHEADER_CLIENT_CA:=/proxycerts/requestheader-client-ca-file}"
ETCD_PROXY_CERT="${ETCD_PROXY_CERT:=/proxycerts/proxy-client-cert-file}"
ETCD_PROXY_KEY="${ETCD_PROXY_KEY:=/proxycerts/proxy-client-key-file}"
K8S_PROXY_CA_CRT_FILEPATH="${K8S_PROXY_CA_CRT_FILEPATH:=/etc/kubernetes/certs/proxy-ca.crt}"
K8S_PROXY_KEY_FILEPATH="${K8S_PROXY_KEY_FILEPATH:=/etc/kubernetes/certs/proxy.key}"
K8S_PROXY_CRT_FILEPATH="${K8S_PROXY_CRT_FILEPATH:=/etc/kubernetes/certs/proxy.crt}"

PROXY_CERTS_LOCK_NAME="master_proxy_cert_lock"
PROXY_CERT_LOCK_FILE="/tmp/create_cert.fifl"

if [[ -z "${COSMOS_URI}" ]]; then
  ETCDCTL_ENDPOINTS="${ETCDCTL_ENDPOINTS:=https://127.0.0.1:2379}"
  ETCDCTL_CA_FILE="${ETCDCTL_CA_FILE:=/etc/kubernetes/certs/ca.crt}"
  ETCD_CA_PARAM="--cacert=${ETCDCTL_CA_FILE}"
else
  ETCDCTL_ENDPOINTS="${ETCDCTL_ENDPOINTS:=https://${COSMOS_URI}:2379}"
  ETCD_CA_PARAM=""
fi
ETCDCTL_KEY_FILE="${ETCDCTL_KEY_FILE:=/etc/kubernetes/certs/etcdclient.key}"
ETCDCTL_CERT_FILE="${ETCDCTL_CERT_FILE:=/etc/kubernetes/certs/etcdclient.crt}"

ETCDCTL_PARAMS="--command-timeout=30s --cert=${ETCDCTL_CERT_FILE} --key=${ETCDCTL_KEY_FILE} ${ETCD_CA_PARAM} --endpoints=${ETCDCTL_ENDPOINTS}"
RANDFILE=$(mktemp)
export RANDFILE

openssl genrsa -out $PROXY_CA_KEY 2048
openssl req -new -x509 -days 1826 -key $PROXY_CA_KEY -out $PROXY_CRT -subj '/CN=proxyClientCA'
openssl genrsa -out $PROXY_CLIENT_KEY 2048
openssl req -new -key $PROXY_CLIENT_KEY -out $PROXY_CLIENT_CSR -subj '/CN=aggregator/O=system:masters'
openssl x509 -req -days 730 -in $PROXY_CLIENT_CSR -CA $PROXY_CRT -CAkey $PROXY_CA_KEY -set_serial 02 -out $PROXY_CLIENT_CRT

write_certs_to_disk() {
    ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} get $ETCD_REQUESTHEADER_CLIENT_CA --print-value-only > $K8S_PROXY_CA_CRT_FILEPATH
    ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} get $ETCD_PROXY_KEY --print-value-only > $K8S_PROXY_KEY_FILEPATH
    ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} get $ETCD_PROXY_CERT --print-value-only > $K8S_PROXY_CRT_FILEPATH
    # Remove whitespace padding at beginning of 1st line
    sed -i '1s/\s//' $K8S_PROXY_CA_CRT_FILEPATH $K8S_PROXY_CRT_FILEPATH $K8S_PROXY_KEY_FILEPATH
    chmod 600 $K8S_PROXY_KEY_FILEPATH
}

write_certs_to_disk_with_retry() {
    for i in $(seq 1 12); do
        write_certs_to_disk && break || sleep 5
    done
}
is_etcd_healthy(){
    for i in $(seq 1 100); do
        ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} endpoint health && break || sleep 5
    done
}
is_etcd_healthy
# lock file to enable "only 1 master generates certs"
rm -f "${PROXY_CERT_LOCK_FILE}"
mkfifo "${PROXY_CERT_LOCK_FILE}"

echo "$(date) attempting to acquire lock for proxy cert gen"
ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} lock ${PROXY_CERTS_LOCK_NAME}  > "${PROXY_CERT_LOCK_FILE}" &
echo "$(date) lock acquired"

pid=$!
if read -r lockthis < "${PROXY_CERT_LOCK_FILE}"; then
  if [[ "" == "$(ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} get $ETCD_REQUESTHEADER_CLIENT_CA --print-value-only)" ]]; then
    ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} put $ETCD_REQUESTHEADER_CLIENT_CA " $(cat ${PROXY_CRT})" >/dev/null 2>&1;
	else
		echo "found client request header ca, not creating one"
  fi
  if [[ "" == "$(ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} get $ETCD_PROXY_KEY --print-value-only)" ]]; then
    ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} put $ETCD_PROXY_KEY " $(cat ${PROXY_CLIENT_KEY})" >/dev/null 2>&1;
	else
		 echo "found proxy key, not creating one"
  fi
  if [[ "" == "$(ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} get $ETCD_PROXY_CERT --print-value-only)" ]]; then
    ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} put $ETCD_PROXY_CERT " $(cat ${PROXY_CLIENT_CRT})" >/dev/null 2>&1;
	else
		echo "found proxy cert, not creating one"
  fi
fi
kill $pid
wait $pid
rm -f "${PROXY_CERT_LOCK_FILE}"

echo "$(date) cert gen and save/check etcd completed"

write_certs_to_disk_with_retry
#EOF
`)

func k8sCloudInitArtifactsGenerateproxycertsShBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsGenerateproxycertsSh, nil
}

func k8sCloudInitArtifactsGenerateproxycertsSh() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsGenerateproxycertsShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/generateproxycerts.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsHealthMonitorSh = []byte(`#!/usr/bin/env bash

# This script originated at https://github.com/kubernetes/kubernetes/blob/master/cluster/gce/gci/health-monitor.sh
# and has been modified for aks-engine.

set -o nounset
set -o pipefail

container_runtime_monitoring() {
  local -r max_attempts=5
  local attempt=1
  local -r crictl="${KUBE_HOME}/bin/crictl"
  local -r container_runtime_name="${CONTAINER_RUNTIME_NAME:-docker}"
  local healthcheck_command="docker ps"
  if [[ "${CONTAINER_RUNTIME:-docker}" != "docker" ]]; then
    healthcheck_command="${crictl} pods"
  fi

  until timeout 60 ${healthcheck_command} > /dev/null; do
    if (( attempt == max_attempts )); then
      echo "Max attempt ${max_attempts} reached! Proceeding to monitor container runtime healthiness."
      break
    fi
    echo "$attempt initial attempt \"${healthcheck_command}\"! Trying again in $attempt seconds..."
    sleep "$(( 2 ** attempt++ ))"
  done
  while true; do
    if ! timeout 60 ${healthcheck_command} > /dev/null; then
      echo "Container runtime ${container_runtime_name} failed!"
      if [[ "$container_runtime_name" == "docker" ]]; then
          pkill -SIGUSR1 dockerd
      fi
      systemctl kill --kill-who=main "${container_runtime_name}"
      sleep 120
    else
      sleep "${SLEEP_SECONDS}"
    fi
  done
}

kubelet_monitoring() {
  echo "Wait for 2 minutes for kubelet to be functional"
  sleep 120
  local -r max_seconds=10
  local output=""
  while true; do
    if ! output=$(curl -m "${max_seconds}" -f -s -S http://127.0.0.1:10255/healthz 2>&1); then
      echo $output
      echo "Kubelet is unhealthy!"
      systemctl kill kubelet
      sleep 60
    else
      sleep "${SLEEP_SECONDS}"
    fi
  done
}

if [[ "$#" -ne 1 ]]; then
  echo "Usage: health-monitor.sh <container-runtime/kubelet>"
  exit 1
fi

KUBE_HOME="/usr/local/bin"
KUBE_ENV="/etc/default/kube-env"
if [[  -e "${KUBE_ENV}" ]]; then
  source "${KUBE_ENV}"
fi

SLEEP_SECONDS=10
component=$1
echo "Start kubernetes health monitoring for ${component}"

if [[ "${component}" == "container-runtime" ]]; then
  container_runtime_monitoring
elif [[ "${component}" == "kubelet" ]]; then
  kubelet_monitoring
else
  echo "Health monitoring for component ${component} is not supported!"
fi
`)

func k8sCloudInitArtifactsHealthMonitorShBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsHealthMonitorSh, nil
}

func k8sCloudInitArtifactsHealthMonitorSh() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsHealthMonitorShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/health-monitor.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsKmsService = []byte(`[Unit]
Description=azurekms
Requires=docker.service
After=network-online.target

[Service]
Type=simple
Restart=always
TimeoutStartSec=0
ExecStart=/usr/bin/docker run \
  --net=host \
  --volume=/opt:/opt \
  --volume=/etc/kubernetes:/etc/kubernetes \
  --volume=/etc/ssl/certs/ca-certificates.crt:/etc/ssl/certs/ca-certificates.crt \
  --volume=/var/lib/waagent:/var/lib/waagent \
  mcr.microsoft.com/k8s/kms/keyvault:v0.0.9

[Install]
WantedBy=multi-user.target
`)

func k8sCloudInitArtifactsKmsServiceBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsKmsService, nil
}

func k8sCloudInitArtifactsKmsService() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsKmsServiceBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/kms.service", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsKubeletMonitorService = []byte(`[Unit]
Description=a script that checks kubelet health and restarts if needed
After=kubelet.service
[Service]
Restart=always
RestartSec=10
RemainAfterExit=yes
ExecStart=/usr/local/bin/health-monitor.sh kubelet`)

func k8sCloudInitArtifactsKubeletMonitorServiceBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsKubeletMonitorService, nil
}

func k8sCloudInitArtifactsKubeletMonitorService() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsKubeletMonitorServiceBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/kubelet-monitor.service", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsKubeletMonitorTimer = []byte(`[Unit]
Description=a timer that delays kubelet-monitor from starting too soon after boot
[Timer]
OnBootSec=30min
[Install]
WantedBy=multi-user.target`)

func k8sCloudInitArtifactsKubeletMonitorTimerBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsKubeletMonitorTimer, nil
}

func k8sCloudInitArtifactsKubeletMonitorTimer() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsKubeletMonitorTimerBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/kubelet-monitor.timer", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsKubeletService = []byte(`[Unit]
Description=Kubelet
ConditionPathExists=/usr/local/bin/kubelet
# Required

[Service]
Restart=always
EnvironmentFile=/etc/default/kubelet
SuccessExitStatus=143
ExecStartPre=/bin/bash /opt/azure/containers/kubelet.sh
ExecStartPre=/bin/mkdir -p /var/lib/kubelet
ExecStartPre=/bin/mkdir -p /var/lib/cni
ExecStartPre=/bin/bash -c "if [ $(mount | grep \"/var/lib/kubelet\" | wc -l) -le 0 ] ; then /bin/mount --bind /var/lib/kubelet /var/lib/kubelet ; fi"
ExecStartPre=/bin/mount --make-shared /var/lib/kubelet
# This is a partial workaround to this upstream Kubernetes issue:
#  https://github.com/kubernetes/kubernetes/issues/41916#issuecomment-312428731
ExecStartPre=/sbin/sysctl -w net.ipv4.tcp_retries2=8
ExecStartPre=-/sbin/ebtables -t nat --list
ExecStartPre=-/sbin/iptables -t nat --numeric --list
ExecStart=/usr/local/bin/kubelet \
        --enable-server \
        --node-labels="${KUBELET_NODE_LABELS}" \
        --v=2 \
        --volume-plugin-dir=/etc/kubernetes/volumeplugins \
        $KUBELET_CONFIG $KUBELET_OPTS \
        $KUBELET_REGISTER_NODE $KUBELET_REGISTER_WITH_TAINTS

[Install]
WantedBy=multi-user.target
`)

func k8sCloudInitArtifactsKubeletServiceBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsKubeletService, nil
}

func k8sCloudInitArtifactsKubeletService() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsKubeletServiceBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/kubelet.service", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsLabelNodesService = []byte(`[Unit]
Description=Label Kubernetes nodes as masters or agents
After=kubelet.service
[Service]
Restart=always
RestartSec=60
ExecStart=/bin/bash /opt/azure/containers/label-nodes.sh
#EOF
`)

func k8sCloudInitArtifactsLabelNodesServiceBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsLabelNodesService, nil
}

func k8sCloudInitArtifactsLabelNodesService() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsLabelNodesServiceBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/label-nodes.service", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsLabelNodesSh = []byte(`#!/usr/bin/env bash

# Applies missing master and agent labels to Kubernetes nodes.
#
# Kubelet 1.16+ rejects the ` + "`" + `kubernetes.io/role` + "`" + ` and ` + "`" + `node-role.kubernetes.io` + "`" + `
# labels in its ` + "`" + `--node-labels` + "`" + ` argument, but they need to be present for
# backward compatibility.

set -euo pipefail

MASTER_SELECTOR="kubernetes.azure.com/role!=agent,kubernetes.io/role!=agent"
MASTER_LABELS="kubernetes.azure.com/role=master kubernetes.io/role=master node-role.kubernetes.io/master="
AGENT_SELECTOR="kubernetes.azure.com/role!=master,kubernetes.io/role!=master"
AGENT_LABELS="kubernetes.azure.com/role=agent kubernetes.io/role=agent node-role.kubernetes.io/agent="

kubectl label nodes --overwrite -l $MASTER_SELECTOR $MASTER_LABELS
kubectl label nodes --overwrite -l $AGENT_SELECTOR $AGENT_LABELS
#EOF
`)

func k8sCloudInitArtifactsLabelNodesShBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsLabelNodesSh, nil
}

func k8sCloudInitArtifactsLabelNodesSh() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsLabelNodesShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/label-nodes.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsModprobeCisConf = []byte(`# 3.5.1 Ensure DCCP is disabled
install dccp /bin/true
# 3.5.2 Ensure SCTP is disabled
install sctp /bin/true
# 3.5.3 Ensure RDS is disabled
install rds /bin/true
# 3.5.4 Ensure TIPC is disabled
install tipc /bin/true
# 1.1.1.1 Ensure mounting of cramfs filesystems is disabled
install cramfs /bin/true
# 1.1.1.2 Ensure mounting of freevxfs filesystems is disabled
install freevxfs /bin/true
# 1.1.1.3 Ensure mounting of jffs2 filesystems is disabled
install jffs2 /bin/true
# 1.1.1.4 Ensure mounting of hfs filesystems is disabled
install hfs /bin/true
# 1.1.1.5 Ensure mounting of hfsplus filesystems is disabled
install hfsplus /bin/true`)

func k8sCloudInitArtifactsModprobeCisConfBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsModprobeCisConf, nil
}

func k8sCloudInitArtifactsModprobeCisConf() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsModprobeCisConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/modprobe-CIS.conf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsMountetcdSh = []byte(`#!/bin/bash
# Mounting is done here instead of etcd because of bug https://bugs.launchpad.net/cloud-init/+bug/1692093
# Once the bug is fixed, replace the below with the cloud init changes replaced in https://github.com/Azure/aks-engine/pull/661.
set -x
DISK=/dev/sdc
PARTITION=${DISK}1
MOUNTPOINT=/var/lib/etcddisk
udevadm settle
mkdir -p $MOUNTPOINT
if mount | grep $MOUNTPOINT
then
    echo "disk is already mounted"
    exit 0
fi
if ! grep "/dev/sdc1" /etc/fstab
then
    echo "$PARTITION       $MOUNTPOINT       auto    defaults,nofail       0       2" >> /etc/fstab
fi
if ! ls $PARTITION
then
    /sbin/sgdisk --new 1 $DISK
    /sbin/mkfs.ext4 $PARTITION -L etcd_disk -F -E lazy_itable_init=1,lazy_journal_init=1
fi
mount $MOUNTPOINT
/bin/chown -R etcd:etcd /var/lib/etcddisk
#EOF
`)

func k8sCloudInitArtifactsMountetcdShBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsMountetcdSh, nil
}

func k8sCloudInitArtifactsMountetcdSh() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsMountetcdShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/mountetcd.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsPamDCommonAuth = []byte(`#
# /etc/pam.d/common-auth - authentication settings common to all services
#
# This file is included from other service-specific PAM config files,
# and should contain a list of the authentication modules that define
# the central authentication scheme for use on the system
# (e.g., /etc/shadow, LDAP, Kerberos, etc.).  The default is to use the
# traditional Unix authentication mechanisms.
#
# As of pam 1.0.1-6, this file is managed by pam-auth-update by default.
# To take advantage of this, it is recommended that you configure any
# local modules either before or after the default block, and use
# pam-auth-update to manage selection of other modules.  See
# pam-auth-update(8) for details.

# here are the per-package modules (the "Primary" block)
auth	[success=1 default=ignore]	pam_unix.so nullok_secure
# here's the fallback if no module succeeds
auth	requisite			pam_deny.so
# prime the stack with a positive return value if there isn't one already;
# this avoids us returning an error just because nothing sets a success code
# since the modules above will each just jump around
auth	required			pam_permit.so
# and here are more per-package modules (the "Additional" block)
# end of pam-auth-update config

# 5.3.2 Ensure lockout for failed password attempts is configured
auth required pam_tally2.so onerr=fail audit silent deny=5 unlock_time=900
`)

func k8sCloudInitArtifactsPamDCommonAuthBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsPamDCommonAuth, nil
}

func k8sCloudInitArtifactsPamDCommonAuth() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsPamDCommonAuthBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/pam-d-common-auth", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsPamDCommonPassword = []byte(`#
# /etc/pam.d/common-password - password-related modules common to all services
#
# This file is included from other service-specific PAM config files,
# and should contain a list of modules that define the services to be
# used to change user passwords.  The default is pam_unix.

# Explanation of pam_unix options:
#
# The "sha512" option enables salted SHA512 passwords.  Without this option,
# the default is Unix crypt.  Prior releases used the option "md5".
#
# The "obscure" option replaces the old ` + "`" + `OBSCURE_CHECKS_ENAB' option in
# login.defs.
#
# See the pam_unix manpage for other options.

# As of pam 1.0.1-6, this file is managed by pam-auth-update by default.
# To take advantage of this, it is recommended that you configure any
# local modules either before or after the default block, and use
# pam-auth-update to manage selection of other modules.  See
# pam-auth-update(8) for details.

# here are the per-package modules (the "Primary" block)
password	requisite			pam_pwquality.so retry=3
password	[success=1 default=ignore]	pam_unix.so obscure use_authtok try_first_pass sha512
# here's the fallback if no module succeeds
password	requisite			pam_deny.so
# prime the stack with a positive return value if there isn't one already;
# this avoids us returning an error just because nothing sets a success code
# since the modules above will each just jump around
password	required			pam_permit.so
# and here are more per-package modules (the "Additional" block)
# end of pam-auth-update config

# 5.3.3 Ensure password reuse is limited
# 5.3.4 Ensure password hashing algorithm is SHA-512
password	[success=1 default=ignore]	pam_unix.so obscure use_authtok try_first_pass sha512 remember=5
`)

func k8sCloudInitArtifactsPamDCommonPasswordBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsPamDCommonPassword, nil
}

func k8sCloudInitArtifactsPamDCommonPassword() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsPamDCommonPasswordBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/pam-d-common-password", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsPamDSu = []byte(`#
# The PAM configuration file for the Shadow ` + "`" + `su' service
#

# This allows root to su without passwords (normal operation)
auth       sufficient pam_rootok.so

# Uncomment this to force users to be a member of group root
# before they can use ` + "`" + `su'. You can also add "group=foo"
# to the end of this line if you want to use a group other
# than the default "root" (but this may have side effect of
# denying "root" user, unless she's a member of "foo" or explicitly
# permitted earlier by e.g. "sufficient pam_rootok.so").
# (Replaces the ` + "`" + `SU_WHEEL_ONLY' option from login.defs)

# 5.6 Ensure access to the su command is restricted
auth required pam_wheel.so use_uid

# Uncomment this if you want wheel members to be able to
# su without a password.
# auth       sufficient pam_wheel.so trust

# Uncomment this if you want members of a specific group to not
# be allowed to use su at all.
# auth       required   pam_wheel.so deny group=nosu

# Uncomment and edit /etc/security/time.conf if you need to set
# time restrainst on su usage.
# (Replaces the ` + "`" + `PORTTIME_CHECKS_ENAB' option from login.defs
# as well as /etc/porttime)
# account    requisite  pam_time.so

# This module parses environment configuration file(s)
# and also allows you to use an extended config
# file /etc/security/pam_env.conf.
#
# parsing /etc/environment needs "readenv=1"
session       required   pam_env.so readenv=1
# locale variables are also kept into /etc/default/locale in etch
# reading this file *in addition to /etc/environment* does not hurt
session       required   pam_env.so readenv=1 envfile=/etc/default/locale

# Defines the MAIL environment variable
# However, userdel also needs MAIL_DIR and MAIL_FILE variables
# in /etc/login.defs to make sure that removing a user
# also removes the user's mail spool file.
# See comments in /etc/login.defs
#
# "nopen" stands to avoid reporting new mail when su'ing to another user
session    optional   pam_mail.so nopen

# Sets up user limits according to /etc/security/limits.conf
# (Replaces the use of /etc/limits in old login)
session    required   pam_limits.so

# The standard Unix authentication modules, used with
# NIS (man nsswitch) as well as normal /etc/passwd and
# /etc/shadow entries.
@include common-auth
@include common-account
@include common-session
`)

func k8sCloudInitArtifactsPamDSuBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsPamDSu, nil
}

func k8sCloudInitArtifactsPamDSu() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsPamDSuBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/pam-d-su", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsProfileDCisSh = []byte(`#!/bin/bash

# 5.4.4 Ensure default user umask is 027 or more restrictive
umask 027
`)

func k8sCloudInitArtifactsProfileDCisShBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsProfileDCisSh, nil
}

func k8sCloudInitArtifactsProfileDCisSh() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsProfileDCisShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/profile-d-cis.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsPwqualityCisConf = []byte(`# 5.3.1 Ensure password creation requirements are configured (Scored)

minlen=14
dcredit=-1
ucredit=-1
ocredit=-1
lcredit=-1`)

func k8sCloudInitArtifactsPwqualityCisConfBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsPwqualityCisConf, nil
}

func k8sCloudInitArtifactsPwqualityCisConf() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsPwqualityCisConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/pwquality-CIS.conf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsRsyslogD60CisConf = []byte(`# 4.2.1.2 Ensure logging is configured (Not Scored)
*.emerg                            :omusrmsg:*
mail.*                             -/var/log/mail
mail.info                          -/var/log/mail.info
mail.warning                       -/var/log/mail.warn
mail.err                           /var/log/mail.err
news.crit                          -/var/log/news/news.crit
news.err                           -/var/log/news/news.err
news.notice                        -/var/log/news/news.notice
*.=warning;*.=err                  -/var/log/warn
*.crit                             /var/log/warn
*.*;mail.none;news.none            -/var/log/messages
local0,local1.*                    -/var/log/localmessages
local2,local3.*                    -/var/log/localmessages
local4,local5.*                    -/var/log/localmessages
local6,local7.*                    -/var/log/localmessages`)

func k8sCloudInitArtifactsRsyslogD60CisConfBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsRsyslogD60CisConf, nil
}

func k8sCloudInitArtifactsRsyslogD60CisConf() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsRsyslogD60CisConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/rsyslog-d-60-CIS.conf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsSetupCustomSearchDomainsSh = []byte(`#!/bin/bash
set -x
source /opt/azure/containers/provision_source.sh

echo "  dns-search <searchDomainName>" | tee -a /etc/network/interfaces.d/50-cloud-init.cfg
systemctl_restart 20 5 10 restart networking
wait_for_apt_locks
retrycmd_if_failure 10 5 120 apt-get -y install realmd sssd sssd-tools samba-common samba samba-common python2.7 samba-libs packagekit
wait_for_apt_locks
echo "<searchDomainRealmPassword>" | realm join -U <searchDomainRealmUser>@$(echo "<searchDomainName>" | tr /a-z/ /A-Z/) $(echo "<searchDomainName>" | tr /a-z/ /A-Z/)
`)

func k8sCloudInitArtifactsSetupCustomSearchDomainsShBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsSetupCustomSearchDomainsSh, nil
}

func k8sCloudInitArtifactsSetupCustomSearchDomainsSh() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsSetupCustomSearchDomainsShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/setup-custom-search-domains.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsSshd_config = []byte(`# What ports, IPs and protocols we listen for
Port 22
# Use these options to restrict which interfaces/protocols sshd will bind to
#ListenAddress ::
#ListenAddress 0.0.0.0
Protocol 2

# 5.2.11 Ensure only approved MAC algorithms are used
MACs hmac-sha2-512-etm@openssh.com,hmac-sha2-256-etm@openssh.com,umac-128-etm@openssh.com,hmac-sha2-512,hmac-sha2-256,umac-128@openssh.com
KexAlgorithms curve25519-sha256@libssh.org
Ciphers chacha20-poly1305@openssh.com,aes256-gcm@openssh.com,aes128-gcm@openssh.com,aes256-ctr,aes192-ctr,aes128-ctr

# 5.2.12 Ensure SSH Idle Timeout Interval is configured
ClientAliveInterval 120
ClientAliveCountMax 3

# HostKeys for protocol version 2
HostKey /etc/ssh/ssh_host_rsa_key
HostKey /etc/ssh/ssh_host_dsa_key
HostKey /etc/ssh/ssh_host_ecdsa_key
HostKey /etc/ssh/ssh_host_ed25519_key

# Logging
SyslogFacility AUTH
LogLevel INFO

# Authentication:
LoginGraceTime 60

# 5.2.8 Ensure SSH root login is disabled
PermitRootLogin no
# 5.2.10 Ensure SSH PermitUserEnvironment is disabled
PermitUserEnvironment no

StrictModes yes
PubkeyAuthentication yes
#AuthorizedKeysFile	%h/.ssh/authorized_keys

# Don't read the user's ~/.rhosts and ~/.shosts files
IgnoreRhosts yes
# similar for protocol version 2
HostbasedAuthentication no

# To enable empty passwords, change to yes (NOT RECOMMENDED)
PermitEmptyPasswords no

# Change to yes to enable challenge-response passwords (beware issues with
# some PAM modules and threads)
ChallengeResponseAuthentication no

# Change to no to disable tunnelled clear text passwords
PasswordAuthentication no

# 5.2.4 Ensure SSH X11 forwarding is disabled
X11Forwarding no

# 5.2.5 Ensure SSH MaxAuthTries is set to 4 or less
MaxAuthTries 4

X11DisplayOffset 10
PrintMotd no
PrintLastLog yes
TCPKeepAlive yes
#UseLogin no

#MaxStartups 10:30:60
Banner /etc/issue.net

# Allow client to pass locale environment variables
AcceptEnv LANG LC_*

Subsystem sftp /usr/lib/openssh/sftp-server

# Set this to 'yes' to enable PAM authentication, account processing,
# and session processing. If this is enabled, PAM authentication will
# be allowed through the ChallengeResponseAuthentication and
# PasswordAuthentication.  Depending on your PAM configuration,
# PAM authentication via ChallengeResponseAuthentication may bypass
# the setting of "PermitRootLogin without-password".
# If you just want the PAM account and session checks to run without
# PAM authentication, then enable this but set PasswordAuthentication
# and ChallengeResponseAuthentication to 'no'.
UsePAM yes
UseDNS no
GSSAPIAuthentication no
`)

func k8sCloudInitArtifactsSshd_configBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsSshd_config, nil
}

func k8sCloudInitArtifactsSshd_config() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsSshd_configBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/sshd_config", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsSshd_config_1604 = []byte(`# What ports, IPs and protocols we listen for
Port 22
# Use these options to restrict which interfaces/protocols sshd will bind to
#ListenAddress ::
#ListenAddress 0.0.0.0
Protocol 2

# 5.2.11 Ensure only approved MAC algorithms are used
MACs hmac-sha2-512-etm@openssh.com,hmac-sha2-256-etm@openssh.com,umac-128-etm@openssh.com,hmac-sha2-512,hmac-sha2-256,umac-128@openssh.com
KexAlgorithms curve25519-sha256@libssh.org
Ciphers chacha20-poly1305@openssh.com,aes256-gcm@openssh.com,aes128-gcm@openssh.com,aes256-ctr,aes192-ctr,aes128-ctr

# 5.2.12 Ensure SSH Idle Timeout Interval is configured
ClientAliveInterval 120
ClientAliveCountMax 3

# HostKeys for protocol version 2
HostKey /etc/ssh/ssh_host_rsa_key
HostKey /etc/ssh/ssh_host_dsa_key
HostKey /etc/ssh/ssh_host_ecdsa_key
HostKey /etc/ssh/ssh_host_ed25519_key

#Privilege Separation is turned on for security
UsePrivilegeSeparation yes

# Lifetime and size of ephemeral version 1 server key
KeyRegenerationInterval 3600
ServerKeyBits 1024

# Logging
SyslogFacility AUTH
LogLevel INFO

# Authentication:
LoginGraceTime 60

# 5.2.8 Ensure SSH root login is disabled
PermitRootLogin no
# 5.2.10 Ensure SSH PermitUserEnvironment is disabled
PermitUserEnvironment no

StrictModes yes
RSAAuthentication yes
PubkeyAuthentication yes
#AuthorizedKeysFile	%h/.ssh/authorized_keys

# Don't read the user's ~/.rhosts and ~/.shosts files
IgnoreRhosts yes
# For this to work you will also need host keys in /etc/ssh_known_hosts
RhostsRSAAuthentication no
# similar for protocol version 2
HostbasedAuthentication no

# To enable empty passwords, change to yes (NOT RECOMMENDED)
PermitEmptyPasswords no

# Change to yes to enable challenge-response passwords (beware issues with
# some PAM modules and threads)
ChallengeResponseAuthentication no

# Change to no to disable tunnelled clear text passwords
PasswordAuthentication no

# 5.2.4 Ensure SSH X11 forwarding is disabled
X11Forwarding no

# 5.2.5 Ensure SSH MaxAuthTries is set to 4 or less
MaxAuthTries 4

X11DisplayOffset 10
PrintMotd no
PrintLastLog yes
TCPKeepAlive yes
#UseLogin no

#MaxStartups 10:30:60
Banner /etc/issue.net

# Allow client to pass locale environment variables
AcceptEnv LANG LC_*

Subsystem sftp /usr/lib/openssh/sftp-server

# Set this to 'yes' to enable PAM authentication, account processing,
# and session processing. If this is enabled, PAM authentication will
# be allowed through the ChallengeResponseAuthentication and
# PasswordAuthentication.  Depending on your PAM configuration,
# PAM authentication via ChallengeResponseAuthentication may bypass
# the setting of "PermitRootLogin without-password".
# If you just want the PAM account and session checks to run without
# PAM authentication, then enable this but set PasswordAuthentication
# and ChallengeResponseAuthentication to 'no'.
UsePAM yes
UseDNS no
GSSAPIAuthentication no
`)

func k8sCloudInitArtifactsSshd_config_1604Bytes() ([]byte, error) {
	return _k8sCloudInitArtifactsSshd_config_1604, nil
}

func k8sCloudInitArtifactsSshd_config_1604() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsSshd_config_1604Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/sshd_config_1604", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsSysFsBpfMount = []byte(`[Unit]
Description=Cilium BPF mounts
Documentation=http://docs.cilium.io/
DefaultDependencies=no
Before=local-fs.target umount.target
After=swap.target

[Mount]
What=bpffs
Where=/sys/fs/bpf
Type=bpf

[Install]
WantedBy=multi-user.target`)

func k8sCloudInitArtifactsSysFsBpfMountBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsSysFsBpfMount, nil
}

func k8sCloudInitArtifactsSysFsBpfMount() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsSysFsBpfMountBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/sys-fs-bpf.mount", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsSysctlD60CisConf = []byte(`# 3.1.2 Ensure packet redirect sending is disabled
net.ipv4.conf.all.send_redirects = 0
net.ipv4.conf.default.send_redirects = 0
# 3.2.1 Ensure source routed packets are not accepted 
net.ipv4.conf.all.accept_source_route = 0
net.ipv4.conf.default.accept_source_route = 0
# 3.2.2 Ensure ICMP redirects are not accepted
net.ipv4.conf.all.accept_redirects = 0
net.ipv4.conf.default.accept_redirects = 0
# 3.2.3 Ensure secure ICMP redirects are not accepted
net.ipv4.conf.all.secure_redirects = 0
net.ipv4.conf.default.secure_redirects = 0
# 3.2.4 Ensure suspicious packets are logged
net.ipv4.conf.all.log_martians = 1
net.ipv4.conf.default.log_martians = 1
# 3.3.1 Ensure IPv6 router advertisements are not accepted
net.ipv6.conf.all.accept_ra = 0
net.ipv6.conf.default.accept_ra = 0
# 3.3.2 Ensure IPv6 redirects are not accepted
net.ipv6.conf.all.accept_redirects = 0
net.ipv6.conf.default.accept_redirects = 0
# refer to https://github.com/kubernetes/kubernetes/blob/75d45bdfc9eeda15fb550e00da662c12d7d37985/pkg/kubelet/cm/container_manager_linux.go#L359-L397
vm.overcommit_memory = 1
kernel.panic = 10
kernel.panic_on_oops = 1
`)

func k8sCloudInitArtifactsSysctlD60CisConfBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsSysctlD60CisConf, nil
}

func k8sCloudInitArtifactsSysctlD60CisConf() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsSysctlD60CisConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/sysctl-d-60-CIS.conf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitJumpboxcustomdataYml = []byte(`#cloud-config

write_files:

- path: "/opt/azure/containers/provision_source.sh"
  permissions: "0744"
  encoding: gzip
  owner: "root"
  content: !!binary |
    {{CloudInitData "provisionSource"}}

- path: "/home/{{WrapAsParameter "jumpboxUsername"}}/.kube/config"
  permissions: "0644"
  owner: "{{WrapAsParameter "jumpboxUsername"}}"
  content: |
{{WrapAsVariable "kubeconfig"}}

runcmd:
- . /opt/azure/containers/provision_source.sh
- retrycmd_if_failure 10 5 10 curl -LO https://storage.googleapis.com/kubernetes-release/release/v{{.OrchestratorProfile.OrchestratorVersion}}/bin/linux/amd64/kubectl
- chmod +x ./kubectl
- sudo mv ./kubectl /usr/local/bin/kubectl
- chown -R "{{WrapAsParameter "jumpboxUsername"}}" "/home/{{WrapAsParameter "jumpboxUsername"}}"
- chgrp -R "{{WrapAsParameter "jumpboxUsername"}}" "/home/{{WrapAsParameter "jumpboxUsername"}}"
- chown -R root "/home/{{WrapAsParameter "jumpboxUsername"}}/.kube"
- chgrp -R root "/home/{{WrapAsParameter "jumpboxUsername"}}/.kube"
`)

func k8sCloudInitJumpboxcustomdataYmlBytes() ([]byte, error) {
	return _k8sCloudInitJumpboxcustomdataYml, nil
}

func k8sCloudInitJumpboxcustomdataYml() (*asset, error) {
	bytes, err := k8sCloudInitJumpboxcustomdataYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/jumpboxcustomdata.yml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitMasternodecustomdataYml = []byte(`#cloud-config

{{if not .MasterProfile.IsCoreOS}}
packages:
 - jq
 - traceroute
{{end}}

write_files:
- path: /opt/azure/containers/provision_source.sh
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "provisionSource"}}

- path: /opt/azure/containers/provision.sh
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "provisionScript"}}

- path: /opt/azure/containers/provision_installs.sh
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "provisionInstalls"}}

- path: /opt/azure/containers/provision_configs.sh
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "provisionConfigs"}}

{{if not .MasterProfile.IsVHDDistro}}
- path: /opt/azure/containers/provision_cis.sh
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "provisionCIS"}}
{{end}}

{{if IsAzureStackCloud}}
- path: /opt/azure/containers/provision_configs_custom_cloud.sh
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{WrapAsVariable "provisionConfigsCustomCloud"}}
{{end}}

{{if not .MasterProfile.IsVHDDistro}}
    {{if .MasterProfile.IsCoreOS}}
- path: /opt/bin/health-monitor.sh
    {{else}}
- path: /usr/local/bin/health-monitor.sh
    {{end}}
  permissions: "0544"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "healthMonitorScript"}}

- path: /etc/systemd/system/kubelet-monitor.service
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "kubeletMonitorSystemdService"}}

- path: /etc/systemd/system/docker-monitor.timer
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "dockerMonitorSystemdTimer"}}

- path: /etc/systemd/system/docker-monitor.service
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "dockerMonitorSystemdService"}}

- path: /etc/systemd/system/kubelet.service
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "kubeletSystemdService"}}

- path: /opt/azure/containers/label-nodes.sh
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "labelNodesScript"}}

- path: /etc/systemd/system/label-nodes.service
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "labelNodesSystemdService"}}

- path: /etc/systemd/system/kms.service
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "kmsSystemdService"}}

- path: /etc/apt/preferences
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "aptPreferences"}}
{{end}}

{{if IsIPv6DualStackFeatureEnabled}}
- path: /etc/systemd/system/dhcpv6.service
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "dhcpv6SystemdService"}}

- path: /opt/azure/containers/enable-dhcpv6.sh
  permissions: "0544"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "dhcpv6ConfigurationScript"}}
{{end}}

{{if .OrchestratorProfile.KubernetesConfig.RequiresDocker}}
    {{if not .MasterProfile.IsCoreOS}}
        {{if not .MasterProfile.IsVHDDistro}}
- path: /etc/systemd/system/docker.service.d/clear_mount_propagation_flags.conf
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "dockerClearMountPropagationFlags"}}
         {{end}}
    {{end}}

- path: /etc/systemd/system/docker.service.d/exec_start.conf
  permissions: "0644"
  owner: root
  content: |
    [Service]
    ExecStart=
    {{if .MasterProfile.IsCoreOS}}
    ExecStart=/usr/bin/env PATH=${TORCX_BINDIR}:${PATH} ${TORCX_BINDIR}/dockerd --host=fd:// --containerd=/var/run/docker/libcontainerd/docker-containerd.sock --storage-driver=overlay2 --bip={{WrapAsParameter "dockerBridgeCidr"}} $DOCKER_SELINUX $DOCKER_OPTS $DOCKER_CGROUPS $DOCKER_OPT_BIP $DOCKER_OPT_MTU $DOCKER_OPT_IPMASQ
    {{else}}
    ExecStart=/usr/bin/dockerd -H fd:// --storage-driver=overlay2 --bip={{WrapAsParameter "dockerBridgeCidr"}}
    {{end}}
    ExecStartPost=/sbin/iptables -P FORWARD ACCEPT
    #EOF

- path: /etc/docker/daemon.json
  permissions: "0644"
  owner: root
  content: |
    {
      "live-restore": true,
      "log-driver": "json-file",
      "log-opts":  {
         "max-size": "50m",
         "max-file": "5"
      }
    }
{{end}}

{{if eq .OrchestratorProfile.KubernetesConfig.NetworkPlugin "cilium"}}
- path: /etc/systemd/system/sys-fs-bpf.mount
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{WrapAsVariable "systemdBPFMount"}}
{{end}}

- path: /etc/kubernetes/certs/ca.crt
  permissions: "0644"
  encoding: base64
  owner: root
  content: |
    {{WrapAsParameter "caCertificate"}}

- path: /etc/kubernetes/certs/client.crt
  permissions: "0644"
  encoding: "base64"
  owner: "root"
  content: |
    {{WrapAsParameter "clientCertificate"}}

{{if EnableAggregatedAPIs}}
- path: /etc/kubernetes/generate-proxy-certs.sh
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "generateProxyCertsScript"}}
{{end}}

{{if HasLinuxProfile}}{{if HasCustomSearchDomain}}
- path: /opt/azure/containers/setup-custom-search-domains.sh
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "customSearchDomainsScript"}}
{{end}}{{end}}

- path: /var/lib/kubelet/kubeconfig
  permissions: "0644"
  owner: root
  content: |
    apiVersion: v1
    kind: Config
    clusters:
    - name: localcluster
      cluster:
        certificate-authority: /etc/kubernetes/certs/ca.crt
      {{if IsMasterVirtualMachineScaleSets}}
        server: <SERVERIP>
      {{else}}
        server: {{WrapAsVerbatim "concat('https://', variables('masterPrivateIpAddrs')[copyIndex(variables('masterOffset'))], ':443')"}}
      {{end}}
    users:
    - name: client
      user:
        client-certificate: /etc/kubernetes/certs/client.crt
        client-key: /etc/kubernetes/certs/client.key
    contexts:
    - context:
        cluster: localcluster
        user: client
      name: localclustercontext
    current-context: localclustercontext
    #EOF

{{if EnableDataEncryptionAtRest}}
- path: /etc/kubernetes/encryption-config.yaml
  permissions: "0600"
  owner: root
  content: |
    kind: EncryptionConfig
    apiVersion: v1
    resources:
      - resources:
          - secrets
        providers:
          - aescbc:
              keys:
                - name: key1
                  secret: <etcdEncryptionSecret>
          - identity: {}
{{end}}

{{if EnableEncryptionWithExternalKms}}
- path: /etc/kubernetes/encryption-config.yaml
  permissions: "0444"
  owner: root
  content: |
    kind: EncryptionConfig
    apiVersion: v1
    resources:
      - resources:
        - secrets
        providers:
        - kms:
            name: azurekmsprovider
            endpoint: unix:///opt/azurekms.socket
            cachesize: 0
        - identity: {}
{{end}}

MASTER_MANIFESTS_CONFIG_PLACEHOLDER

MASTER_ADDONS_CONFIG_PLACEHOLDER

MASTER_CUSTOM_FILES_PLACEHOLDER

MASTER_CONTAINER_ADDONS_PLACEHOLDER

- path: /etc/default/kubelet
  permissions: "0644"
  owner: root
  content: |
{{if NeedsContainerd}}
    KUBELET_OPTS=--container-runtime=remote --runtime-request-timeout=15m --container-runtime-endpoint=unix:///run/containerd/containerd.sock
{{else}}
    KUBELET_OPTS=
{{end}}
    KUBELET_CONFIG={{GetKubeletConfigKeyVals .MasterProfile.KubernetesConfig}}
    KUBELET_IMAGE={{WrapAsParameter "kubernetesHyperkubeSpec"}}
{{if IsKubernetesVersionGe "1.16.0-alpha.1"}}
    KUBELET_NODE_LABELS={{GetMasterKubernetesLabels "',variables('labelResourceGroup'),'"}}
{{else}}
    KUBELET_NODE_LABELS={{GetMasterKubernetesLabelsDeprecated "',variables('labelResourceGroup'),'"}}
{{end}}
{{if IsAzureStackCloud }}
    AZURE_ENVIRONMENT_FILEPATH=/etc/kubernetes/azurestackcloud.json
{{end}}
{{if IsKubernetesVersionGe "1.6.0"}}
  {{if AnyAgentIsLinux}}
    KUBELET_REGISTER_NODE=--register-node=true
    KUBELET_REGISTER_WITH_TAINTS=--register-with-taints=node-role.kubernetes.io/master=true:NoSchedule
  {{end}}
{{else}}
    KUBELET_REGISTER_SCHEDULABLE={{WrapAsVariable "registerSchedulable"}}
{{end}}
    #EOF

- path: /opt/azure/containers/kubelet.sh
  permissions: "0755"
  owner: root
  content: |
    #!/bin/bash
    set -e
  {{if IsMasterVirtualMachineScaleSets}}
    {{if .MasterProfile.IsCoreOS}}
    PRIVATE_IP=$(hostname -I | cut -d" " -f1)
    {{else}}
    PRIVATE_IP=$(hostname -i | cut -d" " -f1)
    {{end}}
    sed -i "s|<SERVERIP>|https://$PRIVATE_IP:443|g" "/var/lib/kubelet/kubeconfig"
  {{end}}
{{if gt .MasterProfile.Count 1}}
    # Redirect ILB (4443) traffic to port 443 (ELB) in the prerouting chain
    iptables -t nat -A PREROUTING -p tcp --dport 4443 -j REDIRECT --to-port 443
{{end}}

    sed -i "s|<img>|{{WrapAsParameter "kubernetesAddonManagerSpec"}}|g" /etc/kubernetes/manifests/kube-addon-manager.yaml
    for a in "/etc/kubernetes/manifests/kube-apiserver.yaml /etc/kubernetes/manifests/kube-controller-manager.yaml /etc/kubernetes/manifests/kube-scheduler.yaml"; do
      sed -i "s|<img>|{{WrapAsParameter "kubernetesHyperkubeSpec"}}|g" $a
    done
    a=/etc/kubernetes/manifests/kube-apiserver.yaml
    sed -i "s|<args>|{{GetK8sRuntimeConfigKeyVals .OrchestratorProfile.KubernetesConfig.APIServerConfig}}|g" $a
{{ if HasCosmosEtcd  }}
    sed -i "s|<etcdEndPointUri>|{{ GetCosmosEndPointUri }}|g" $a
{{ else }}
    sed -i "s|<etcdEndPointUri>|127.0.0.1|g" $a
{{ end }}
    sed -i "s|<advertiseAddr>|{{WrapAsVariable "kubernetesAPIServerIP"}}|g" $a
    sed -i "s|<args>|{{GetK8sRuntimeConfigKeyVals .OrchestratorProfile.KubernetesConfig.ControllerManagerConfig}}|g" /etc/kubernetes/manifests/kube-controller-manager.yaml
    sed -i "s|<args>|{{GetK8sRuntimeConfigKeyVals .OrchestratorProfile.KubernetesConfig.SchedulerConfig}}|g" /etc/kubernetes/manifests/kube-scheduler.yaml
    {{ if IsIPv6DualStackFeatureEnabled }}
    sed -i "s|<img>|{{WrapAsParameter "kubernetesHyperkubeSpec"}}|g; s|<CIDR>|',first(split(parameters('kubeClusterCidr'),',')),'|g; s|<kubeProxyMode>|{{ .OrchestratorProfile.KubernetesConfig.ProxyMode}}|g" /etc/kubernetes/addons/kube-proxy-daemonset.yaml
    {{ else }}
    sed -i "s|<img>|{{WrapAsParameter "kubernetesHyperkubeSpec"}}|g; s|<CIDR>|{{WrapAsParameter "kubeClusterCidr"}}|g; s|<kubeProxyMode>|{{ .OrchestratorProfile.KubernetesConfig.ProxyMode}}|g" /etc/kubernetes/addons/kube-proxy-daemonset.yaml
    {{ end }}
    KUBEDNS=/etc/kubernetes/addons/kube-dns-deployment.yaml
{{if NeedsKubeDNSWithExecHealthz}}
    sed -i "s|<img>|{{WrapAsParameter "kubernetesKubeDNSSpec"}}|g; s|<imgMasq>|{{WrapAsParameter "kubernetesDNSMasqSpec"}}|g; s|<imgHealthz>|{{WrapAsParameter "kubernetesExecHealthzSpec"}}|g; s|<imgSidecar>|{{WrapAsParameter "kubernetesDNSSidecarSpec"}}|g; s|<domain>|{{WrapAsParameter "kubernetesKubeletClusterDomain"}}|g; s|<clustIP>|{{WrapAsParameter "kubeDNSServiceIP"}}|g" $KUBEDNS
{{else if IsKubernetesVersionGe "1.12.0"}}
    sed -i "s|<img>|{{WrapAsParameter "kubernetesCoreDNSSpec"}}|g; s|<domain>|{{WrapAsParameter "kubernetesKubeletClusterDomain"}}|g; s|<clustIP>|{{WrapAsParameter "kubeDNSServiceIP"}}|g" /etc/kubernetes/addons/coredns.yaml
{{else}}
    sed -i "s|<img>|{{WrapAsParameter "kubernetesKubeDNSSpec"}}|g; s|<imgMasq>|{{WrapAsParameter "kubernetesDNSMasqSpec"}}|g; s|<imgSidecar>|{{WrapAsParameter "kubernetesDNSSidecarSpec"}}|g; s|<domain>|{{WrapAsParameter "kubernetesKubeletClusterDomain"}}|g; s|<clustIP>|{{WrapAsParameter "kubeDNSServiceIP"}}|g" $KUBEDNS
{{end}}

{{if AdminGroupID }}
    sed -i "s|<gID>|{{WrapAsParameter "aadAdminGroupId"}}|g" "/etc/kubernetes/addons/aad-default-admin-group-rbac.yaml"
{{end}}

{{if .OrchestratorProfile.KubernetesConfig.IsClusterAutoscalerEnabled}}
    sed -i "s|<cloud>|{{WrapAsParameter "kubernetesClusterAutoscalerAzureCloud"}}|g; s|<useManagedIdentity>|{{WrapAsParameter "kubernetesClusterAutoscalerUseManagedIdentity"}}|g" /etc/kubernetes/addons/cluster-autoscaler-deployment.yaml
{{end}}

{{if EnableDataEncryptionAtRest }}
    sed -i "s|<etcdEncryptionSecret>|\"{{WrapAsParameter "etcdEncryptionKey"}}\"|g" /etc/kubernetes/encryption-config.yaml
{{end}}

{{if eq .OrchestratorProfile.KubernetesConfig.NetworkPolicy "calico"}}
    sed -i "s|<kubeClusterCidr>|{{WrapAsParameter "kubeClusterCidr"}}|g" /etc/kubernetes/addons/calico-daemonset.yaml
    {{if eq .OrchestratorProfile.KubernetesConfig.NetworkPlugin "azure"}}
    sed -i "/Start of install-cni initContainer/,/End of install-cni initContainer/d" /etc/kubernetes/addons/calico-daemonset.yaml
    {{else}}
    sed -i "s|<calicoIPAMConfig>|{\"type\": \"host-local\", \"subnet\": \"usePodCidr\"}|g" /etc/kubernetes/addons/calico-daemonset.yaml
    sed -i "s|azv|cali|g" /etc/kubernetes/addons/calico-daemonset.yaml
    {{end}}
{{end}}
{{if eq .OrchestratorProfile.KubernetesConfig.NetworkPlugin "flannel"}}
    sed -i "s|<kubeClusterCidr>|{{WrapAsParameter "kubeClusterCidr"}}|g" /etc/kubernetes/addons/flannel-daemonset.yaml
{{end}}
{{if UseCloudControllerManager }}
    sed -i "s|<img>|{{WrapAsParameter "kubernetesCcmImageSpec"}}|g" /etc/kubernetes/manifests/cloud-controller-manager.yaml
    sed -i "s|<config>|{{GetK8sRuntimeConfigKeyVals .OrchestratorProfile.KubernetesConfig.CloudControllerManagerConfig}}|g" /etc/kubernetes/manifests/cloud-controller-manager.yaml
{{end}}
{{if EnableEncryptionWithExternalKms}}
    sed -i "s|# Required|Requires=kms.service|g" /etc/systemd/system/kubelet.service
{{end}}
{{if HasLinuxProfile}}{{if HasCustomSearchDomain}}
    sed -i "s|<searchDomainName>|{{WrapAsParameter "searchDomainName"}}|g; s|<searchDomainRealmUser>|{{WrapAsParameter "searchDomainRealmUser"}}|g; s|<searchDomainRealmPassword>|{{WrapAsParameter "searchDomainRealmPassword"}}|g" /opt/azure/containers/setup-custom-search-domains.sh
{{end}}{{end}}
    #EOF

- path: /opt/azure/containers/mountetcd.sh
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "mountEtcdScript"}}
{{ if not HasCosmosEtcd  }}
- path: /etc/systemd/system/etcd.service
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "etcdSystemdService"}}

- path: /opt/azure/containers/setup-etcd.sh
  permissions: "0744"
  owner: root
  content: |
    #!/bin/bash
    set -x
  {{if IsMasterVirtualMachineScaleSets}}
    MASTER_VM_NAME=$(hostname)
    MASTER_VM_NAME_BASE=$(hostname | sed "s/.$//")
    MASTER_FIRSTADDR={{WrapAsParameter "firstConsecutiveStaticIP"}}
    MASTER_INDEX=$(hostname | tail -c 2)
    {{if .MasterProfile.IsCoreOS}}
    PRIVATE_IP=$(hostname -I | cut -d" " -f1)
    {{else}}
    PRIVATE_IP=$(hostname -i | cut -d" " -f1)
    {{end}}
    MASTER_COUNT={{WrapAsVariable "masterCount"}}
    IPADDRESS_COUNT={{WrapAsVariable "masterIpAddressCount"}}
    echo $IPADDRESS_COUNT
    ETCD_SERVER_PORT={{WrapAsVariable "masterEtcdServerPort"}}
    ETCD_CLIENT_PORT={{WrapAsVariable "masterEtcdClientPort"}}
    MASTER_URLS=""
    index=0
    IFS=. read -r a b c d <<< "$MASTER_FIRSTADDR"
    d=$((a * 256 ** 3 + b * 256 ** 2 + c * 256 + d))
    echo $d
    while [ $index -lt $MASTER_COUNT ]
    do
        echo $index
        x=` + "`" + `expr $d + $IPADDRESS_COUNT \\* $index` + "`" + `
        echo $x
        s=""
        for i in 1 2 3 4; do s="."$((x%256))$s && ((x>>=8)); done;
        s=$(echo $s | tail -c +2)
        MASTER_URLS="$MASTER_URLS$MASTER_VM_NAME_BASE$index=https://$s:$ETCD_SERVER_PORT,"
        index=` + "`" + `expr $index + 1` + "`" + `
    done
    MASTER_URLS=$(echo $MASTER_URLS | sed "s/.$//")
    echo $MASTER_URLS
    sudo sed -i "1iETCDCTL_ENDPOINTS=https://127.0.0.1:$ETCD_CLIENT_PORT" /etc/environment
    sudo sed -i "1iETCDCTL_CA_FILE={{WrapAsVariable "etcdCaFilepath"}}" /etc/environment
    sudo sed -i "1iETCDCTL_KEY_FILE={{WrapAsVariable "etcdClientKeyFilepath"}}" /etc/environment
    sudo sed -i "1iETCDCTL_CERT_FILE={{WrapAsVariable "etcdClientCertFilepath"}}" /etc/environment
    sudo sed -i "/^DAEMON_ARGS=/d" /etc/default/etcd
    /bin/echo DAEMON_ARGS=--name $MASTER_VM_NAME --peer-client-cert-auth --peer-trusted-ca-file={{WrapAsVariable "etcdCaFilepath"}} --peer-cert-file=/etc/kubernetes/certs/etcdpeer$MASTER_INDEX.crt --peer-key-file=/etc/kubernetes/certs/etcdpeer$MASTER_INDEX.key --initial-advertise-peer-urls "https://$PRIVATE_IP:$ETCD_SERVER_PORT" --listen-peer-urls "https://$PRIVATE_IP:$ETCD_SERVER_PORT" --client-cert-auth --trusted-ca-file={{WrapAsVariable "etcdCaFilepath"}} --cert-file={{WrapAsVariable "etcdServerCertFilepath"}} --key-file={{WrapAsVariable "etcdServerKeyFilepath"}} --advertise-client-urls "https://$PRIVATE_IP:$ETCD_CLIENT_PORT" --listen-client-urls "https://$PRIVATE_IP:$ETCD_CLIENT_PORT,https://127.0.0.1:$ETCD_CLIENT_PORT" --initial-cluster-token "k8s-etcd-cluster" --initial-cluster $MASTER_URLS --data-dir "/var/lib/etcddisk" --initial-cluster-state "new" | tee -a /etc/default/etcd
  {{else}}
    sudo sed -i "1iETCDCTL_ENDPOINTS=https://127.0.0.1:2379" /etc/environment
    sudo sed -i "1iETCDCTL_CA_FILE={{WrapAsVariable "etcdCaFilepath"}}" /etc/environment
    sudo sed -i "1iETCDCTL_KEY_FILE={{WrapAsVariable "etcdClientKeyFilepath"}}" /etc/environment
    sudo sed -i "1iETCDCTL_CERT_FILE={{WrapAsVariable "etcdClientCertFilepath"}}" /etc/environment
    sudo sed -i "/^DAEMON_ARGS=/d" /etc/default/etcd
    /bin/echo DAEMON_ARGS=--name "{{WrapAsVerbatim "variables('masterVMNames')[copyIndex(variables('masterOffset'))]"}}" --peer-client-cert-auth --peer-trusted-ca-file={{WrapAsVariable "etcdCaFilepath"}} --peer-cert-file={{WrapAsVerbatim "variables('etcdPeerCertFilepath')[copyIndex(variables('masterOffset'))]"}} --peer-key-file={{WrapAsVerbatim "variables('etcdPeerKeyFilepath')[copyIndex(variables('masterOffset'))]"}} --initial-advertise-peer-urls "{{WrapAsVerbatim "variables('masterEtcdPeerURLs')[copyIndex(variables('masterOffset'))]"}}" --listen-peer-urls "{{WrapAsVerbatim "variables('masterEtcdPeerURLs')[copyIndex(variables('masterOffset'))]"}}" --client-cert-auth --trusted-ca-file={{WrapAsVariable "etcdCaFilepath"}} --cert-file={{WrapAsVariable "etcdServerCertFilepath"}} --key-file={{WrapAsVariable "etcdServerKeyFilepath"}} --advertise-client-urls "{{WrapAsVerbatim "variables('masterEtcdClientURLs')[copyIndex(variables('masterOffset'))]"}}" --listen-client-urls "{{WrapAsVerbatim "concat(variables('masterEtcdClientURLs')[copyIndex(variables('masterOffset'))], ',https://127.0.0.1:', variables('masterEtcdClientPort'))"}}" --initial-cluster-token "k8s-etcd-cluster" --initial-cluster {{WrapAsVerbatim "variables('masterEtcdClusterStates')[div(variables('masterCount'), 2)]"}} --data-dir "/var/lib/etcddisk" --initial-cluster-state "new" | tee -a /etc/default/etcd
  {{end}}
{{end}}
    #EOF

{{if IsAzureStackCloud}}
- path: "/etc/kubernetes/azurestackcloud.json"
  permissions: "0600"
  owner: "root"
  content: |
    {{WrapAsVariable "environmentJSON"}}
{{end}}

{{if .MasterProfile.IsCoreOS}}
- path: /opt/azure/containers/provision-setup.sh
  permissions: "0755"
  owner: root
  content: |
    #!/bin/bash
    source /opt/azure/containers/provision_source.sh
    /opt/azure/containers/mountetcd.sh
    retrycmd_if_failure 5 5 10 curl --retry 5 --retry-delay 10 --retry-max-time 10 --max-time 60 https://127.0.0.1:2379/v2/machines

    {{if EnableAggregatedAPIs}}
    sudo bash /etc/kubernetes/generate-proxy-certs.sh
    {{end}}

    touch /opt/azure/containers/runcmd.complete

- path: "/etc/kubernetes/manifests/.keep"

{{if .OrchestratorProfile.KubernetesConfig.RequiresDocker}}
groups:
  - docker: [{{WrapAsParameter "linuxAdminUsername"}}]
{{end}}

coreos:
  units:
    - name: start-provision-setup.service
      command: "start"
      content: |
        [Unit]
        Description=Start provision setup service

        [Service]
        ExecStart=/opt/azure/containers/provision-setup.sh
    - name: kubelet.service
      enable: true
      drop-ins:
        - name: "10-coreos.conf"
          content: |
            [Unit]
            Requires=rpc-statd.service
            After=etcd.service
            ConditionPathExists=
            ConditionPathExists=/opt/kubelet
            [Service]
            ExecStart=
            ExecStart=/opt/kubelet \
              --enable-server \
              --node-labels="${KUBELET_NODE_LABELS}" \
              --v=2 \
              --volume-plugin-dir=/etc/kubernetes/volumeplugins \
              $KUBELET_CONFIG $KUBELET_OPTS \
              $KUBELET_REGISTER_NODE $KUBELET_REGISTER_WITH_TAINTS
    - name: kubelet-monitor.service
      enable: true
      drop-ins:
        - name: "10-coreos.conf"
          content: |
            [Service]
            ExecStart=
            ExecStart=/opt/bin/health-monitor.sh kubelet
    - name: docker-monitor.service
      enable: true
      drop-ins:
        - name: "10-coreos.conf"
          content: |
            [Service]
            ExecStart=
            ExecStart=/opt/bin/health-monitor.sh container-runtime
    - name: etcd.service
      enable: true
      drop-ins:
        - name: "10-coreos.conf"
          content: |
            [Service]
            ExecStart=
            ExecStart=/opt/bin/etcd $DAEMON_ARGS
    - name: etcd-member.service
      mask: true
{{else}}
runcmd:
- set -x
- . /opt/azure/containers/provision_source.sh
- aptmarkWALinuxAgent hold{{GetKubernetesMasterPreprovisionYaml}}
{{end}}
`)

func k8sCloudInitMasternodecustomdataYmlBytes() ([]byte, error) {
	return _k8sCloudInitMasternodecustomdataYml, nil
}

func k8sCloudInitMasternodecustomdataYml() (*asset, error) {
	bytes, err := k8sCloudInitMasternodecustomdataYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/masternodecustomdata.yml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitNodecustomdataYml = []byte(`#cloud-config

write_files:
- path: /opt/azure/containers/provision_source.sh
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "provisionSource"}}

- path: /opt/azure/containers/provision.sh
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "provisionScript"}}

- path: /opt/azure/containers/provision_installs.sh
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "provisionInstalls"}}

- path: /opt/azure/containers/provision_configs.sh
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "provisionConfigs"}}

{{if not .IsVHDDistro}}
- path: /opt/azure/containers/provision_cis.sh
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "provisionCIS"}}
{{end}}

{{if IsAzureStackCloud}}
- path: /opt/azure/containers/provision_configs_custom_cloud.sh
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{WrapAsVariable "provisionConfigsCustomCloud"}}
{{end}}

{{if not .IsVHDDistro}}
    {{if .IsCoreOS}}
- path: /opt/bin/health-monitor.sh
    {{else}}
- path: /usr/local/bin/health-monitor.sh
    {{end}}
  permissions: "0544"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "healthMonitorScript"}}

- path: /etc/systemd/system/kubelet-monitor.service
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "kubeletMonitorSystemdService"}}

- path: /etc/systemd/system/docker-monitor.timer
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "dockerMonitorSystemdTimer"}}

- path: /etc/systemd/system/docker-monitor.service
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "dockerMonitorSystemdService"}}

- path: /etc/systemd/system/kubelet.service
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "kubeletSystemdService"}}

- path: /etc/systemd/system/kms.service
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "kmsSystemdService"}}

- path: /etc/apt/preferences
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "aptPreferences"}}
{{end}}

{{if IsIPv6DualStackFeatureEnabled}}
- path: /etc/systemd/system/dhcpv6.service
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "dhcpv6SystemdService"}}

- path: /opt/azure/containers/enable-dhcpv6.sh
  permissions: "0544"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "dhcpv6ConfigurationScript"}}
{{end}}

{{if .KubernetesConfig.RequiresDocker}}
    {{if not .IsCoreOS}}
        {{if not .IsVHDDistro}}
- path: /etc/systemd/system/docker.service.d/clear_mount_propagation_flags.conf
  permissions: "0644"
  encoding: gzip
  owner: "root"
  content: !!binary |
    {{CloudInitData "dockerClearMountPropagationFlags"}}
        {{end}}
    {{end}}

- path: /etc/systemd/system/docker.service.d/exec_start.conf
  permissions: "0644"
  owner: root
  content: |
    [Service]
    ExecStart=
    {{if .IsCoreOS}}
    ExecStart=/usr/bin/env PATH=${TORCX_BINDIR}:${PATH} ${TORCX_BINDIR}/dockerd --host=fd:// --containerd=/var/run/docker/libcontainerd/docker-containerd.sock --storage-driver=overlay2 --bip={{WrapAsParameter "dockerBridgeCidr"}} $DOCKER_SELINUX $DOCKER_OPTS $DOCKER_CGROUPS $DOCKER_OPT_BIP $DOCKER_OPT_MTU $DOCKER_OPT_IPMASQ
    {{else}}
    ExecStart=/usr/bin/dockerd -H fd:// --storage-driver=overlay2 --bip={{WrapAsParameter "dockerBridgeCidr"}}
    {{end}}
    ExecStartPost=/sbin/iptables -P FORWARD ACCEPT
    #EOF

- path: /etc/docker/daemon.json
  permissions: "0644"
  owner: root
  content: |
    {
      "live-restore": true,
      "log-driver": "json-file",
      "log-opts":  {
         "max-size": "50m",
         "max-file": "5"
      }{{if IsNSeriesSKU .}}
      ,"default-runtime": "nvidia",
      "runtimes": {
         "nvidia": {
             "path": "/usr/bin/nvidia-container-runtime",
             "runtimeArgs": []
        }
      }{{end}}
    }
{{end}}

{{if HasCiliumNetworkPlugin }}
- path: /etc/systemd/system/sys-fs-bpf.mount
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{WrapAsVariable "systemdBPFMount"}}
{{end}}

{{if IsNSeriesSKU .}}
- path: /etc/systemd/system/nvidia-modprobe.service
  permissions: "0644"
  owner: root
  content: |
    [Unit]
    Description=Installs and loads Nvidia GPU kernel module
    [Service]
    Type=oneshot
    RemainAfterExit=true
    ExecStartPre=/bin/sh -c "dkms autoinstall --verbose"
    ExecStart=/bin/sh -c "nvidia-modprobe -u -c0"
    ExecStartPost=/bin/sh -c "sleep 10 && systemctl restart kubelet"
    [Install]
    WantedBy=multi-user.target
{{end}}

- path: /etc/kubernetes/certs/ca.crt
  permissions: "0644"
  encoding: base64
  owner: root
  content: |
    {{WrapAsParameter "caCertificate"}}

- path: /etc/kubernetes/certs/client.crt
  permissions: "0644"
  encoding: base64
  owner: root
  content: |
    {{WrapAsParameter "clientCertificate"}}

{{if HasLinuxProfile}}{{if HasCustomSearchDomain}}
- path: /opt/azure/containers/setup-custom-search-domains.sh
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "customSearchDomainsScript"}}
{{end}}{{end}}

- path: /var/lib/kubelet/kubeconfig
  permissions: "0644"
  owner: root
  content: |
    apiVersion: v1
    kind: Config
    clusters:
    - name: localcluster
      cluster:
        certificate-authority: /etc/kubernetes/certs/ca.crt
        server: https://{{WrapAsVariable "kubernetesAPIServerIP"}}:443
    users:
    - name: client
      user:
        client-certificate: /etc/kubernetes/certs/client.crt
        client-key: /etc/kubernetes/certs/client.key
    contexts:
    - context:
        cluster: localcluster
        user: client
      name: localclustercontext
    current-context: localclustercontext
    #EOF

- path: /etc/default/kubelet
  permissions: "0644"
  owner: root
  content: |
{{if NeedsContainerd}}
    KUBELET_OPTS=--container-runtime=remote --runtime-request-timeout=15m --container-runtime-endpoint=unix:///run/containerd/containerd.sock
{{else}}
    KUBELET_OPTS=
{{end}}
    KUBELET_CONFIG={{GetKubeletConfigKeyVals .KubernetesConfig }}
    KUBELET_IMAGE={{WrapAsParameter "kubernetesHyperkubeSpec"}}
    KUBELET_REGISTER_SCHEDULABLE=true
{{if IsKubernetesVersionGe "1.16.0-alpha.1"}}
    KUBELET_NODE_LABELS={{GetAgentKubernetesLabels . "',variables('labelResourceGroup'),'"}}
{{else}}
    KUBELET_NODE_LABELS={{GetAgentKubernetesLabelsDeprecated . "',variables('labelResourceGroup'),'"}}
{{end}}
{{if IsAzureStackCloud }}
    AZURE_ENVIRONMENT_FILEPATH=/etc/kubernetes/azurestackcloud.json
{{end}}
    #EOF

- path: /opt/azure/containers/kubelet.sh
  permissions: "0755"
  owner: root
  content: |
    #!/bin/bash
{{if not IsIPMasqAgentEnabled}}
    {{if IsAzureCNI}}
    iptables -t nat -A POSTROUTING -m iprange ! --dst-range 168.63.129.16 -m addrtype ! --dst-type local ! -d {{WrapAsParameter "vnetCidr"}} -j MASQUERADE
    {{end}}
{{end}}
{{if HasLinuxProfile}}{{if HasCustomSearchDomain}}
    sed -i "s|<searchDomainName>|{{WrapAsParameter "searchDomainName"}}|g" "/opt/azure/containers/setup-custom-search-domains.sh"
    sed -i "s|<searchDomainRealmUser>|{{WrapAsParameter "searchDomainRealmUser"}}|g" "/opt/azure/containers/setup-custom-search-domains.sh"
    sed -i "s|<searchDomainRealmPassword>|{{WrapAsParameter "searchDomainRealmPassword"}}|g" "/opt/azure/containers/setup-custom-search-domains.sh"
{{end}}{{end}}
    #EOF

{{if IsAzureStackCloud}}
- path: "/etc/kubernetes/azurestackcloud.json"
  permissions: "0600"
  owner: "root"
  content: |
    {{WrapAsVariable "environmentJSON"}}
{{end}}

{{if .IsCoreOS}}
- path: "/etc/kubernetes/manifests/.keep"

{{if .KubernetesConfig.RequiresDocker}}
groups:
  - docker: [{{WrapAsParameter "linuxAdminUsername"}}]
{{end}}

coreos:
  units:
    - name: kubelet.service
      enable: true
      drop-ins:
        - name: "10-coreos.conf"
          content: |
            [Unit]
            Requires=rpc-statd.service
            ConditionPathExists=
            ConditionPathExists=/opt/kubelet
            [Service]
            ExecStart=
            ExecStart=/opt/kubelet \
              --enable-server \
              --node-labels="${KUBELET_NODE_LABELS}" \
              --v=2 \
              --volume-plugin-dir=/etc/kubernetes/volumeplugins \
              $KUBELET_CONFIG $KUBELET_OPTS \
              $KUBELET_REGISTER_NODE $KUBELET_REGISTER_WITH_TAINTS
    - name: kubelet-monitor.service
      enable: true
      drop-ins:
        - name: "10-coreos.conf"
          content: |
            [Service]
            ExecStart=
            ExecStart=/opt/bin/health-monitor.sh kubelet
    - name: docker-monitor.service
      enable: true
      drop-ins:
        - name: "10-coreos.conf"
          content: |
            [Service]
            ExecStart=
            ExecStart=/opt/bin/health-monitor.sh container-runtime
    - name: rpcbind.service
      enable: true
{{else}}
runcmd:
- set -x
- . /opt/azure/containers/provision_source.sh
- aptmarkWALinuxAgent hold{{GetKubernetesAgentPreprovisionYaml .}}
{{end}}
`)

func k8sCloudInitNodecustomdataYmlBytes() ([]byte, error) {
	return _k8sCloudInitNodecustomdataYml, nil
}

func k8sCloudInitNodecustomdataYml() (*asset, error) {
	bytes, err := k8sCloudInitNodecustomdataYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/nodecustomdata.yml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddons116AzureCniNetworkmonitorYaml = []byte(`apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: azure-cni-networkmonitor
  namespace: kube-system
  labels:
    app: azure-cnms
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  selector:
    matchLabels:
      k8s-app: azure-cnms
  template:
    metadata:
      labels:
        k8s-app: azure-cnms
    spec:
      priorityClassName: system-node-critical
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      - key: node-role.kubernetes.io/master
        operator: Equal
        value: "true"
        effect: NoSchedule
      - operator: "Exists"
        effect: NoExecute
      - operator: "Exists"
        effect: NoSchedule
      nodeSelector:
        beta.kubernetes.io/os: linux
      containers:
        - name: azure-cnms
          image: {{ContainerImage "azure-cni-networkmonitor"}}
          imagePullPolicy: IfNotPresent
          securityContext:
            privileged: true
          env:
            - name: HOSTNAME
              valueFrom:
                fieldRef:
                  apiVersion: v1
                  fieldPath: spec.nodeName
          volumeMounts:
          - name: ebtables-rule-repo
            mountPath: /var/run
          - name: log
            mountPath: /var/log
          - name: telemetry
            mountPath: /opt/cni/bin
      hostNetwork: true
      volumes:
      - name: log
        hostPath:
          path: /var/log
          type: Directory
      - name: ebtables-rule-repo
        hostPath:
          path: /var/run/
          type: Directory
      - name: telemetry
        hostPath:
          path: /opt/cni/bin
          type: Directory
`)

func k8sContaineraddons116AzureCniNetworkmonitorYamlBytes() ([]byte, error) {
	return _k8sContaineraddons116AzureCniNetworkmonitorYaml, nil
}

func k8sContaineraddons116AzureCniNetworkmonitorYaml() (*asset, error) {
	bytes, err := k8sContaineraddons116AzureCniNetworkmonitorYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/1.16/azure-cni-networkmonitor.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddons116IpMasqAgentYaml = []byte(`apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: azure-ip-masq-agent
  namespace: kube-system
  labels:
    component: azure-ip-masq-agent
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
    tier: node
spec:
  selector:
    matchLabels:
      k8s-app: azure-ip-masq-agent
      tier: node
  template:
    metadata:
      labels:
        k8s-app: azure-ip-masq-agent
        tier: node
    spec:
      priorityClassName: system-node-critical
      hostNetwork: true
      nodeSelector:
        beta.kubernetes.io/os: linux
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      - key: node-role.kubernetes.io/master
        operator: Equal
        value: "true"
        effect: NoSchedule
      - operator: "Exists"
        effect: NoExecute
      - operator: "Exists"
        effect: NoSchedule
      containers:
      - name: azure-ip-masq-agent
        image: {{ContainerImage "ip-masq-agent"}}
        imagePullPolicy: IfNotPresent
        securityContext:
          privileged: true
        volumeMounts:
          - name: azure-ip-masq-agent-config-volume
            mountPath: /etc/config
        resources:
          requests:
            cpu: {{ContainerCPUReqs "ip-masq-agent"}}
            memory: {{ContainerMemReqs "ip-masq-agent"}}
          limits:
            cpu: {{ContainerCPULimits "ip-masq-agent"}}
            memory: {{ContainerMemLimits "ip-masq-agent"}}
      volumes:
        - name: azure-ip-masq-agent-config-volume
          configMap:
            name: azure-ip-masq-agent-config
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: azure-ip-masq-agent-config
  namespace: kube-system
  labels:
    component: azure-ip-masq-agent
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
data:
  ip-masq-agent: |-
    nonMasqueradeCIDRs:
      - {{ContainerConfig "non-masquerade-cidr"}}
    {{- if ContainerConfig "non-masq-cni-cidr"}}
      - {{ContainerConfig "non-masq-cni-cidr"}}
    masqLinkLocal: true
    {{else}}
    masqLinkLocal: false
    {{end -}}
    resyncInterval: 60s
`)

func k8sContaineraddons116IpMasqAgentYamlBytes() ([]byte, error) {
	return _k8sContaineraddons116IpMasqAgentYaml, nil
}

func k8sContaineraddons116IpMasqAgentYaml() (*asset, error) {
	bytes, err := k8sContaineraddons116IpMasqAgentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/1.16/ip-masq-agent.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddons116KubernetesmasteraddonsAadPodIdentityDeploymentYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: aad-pod-id-nmi-service-account
  namespace: default
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: azureassignedidentities.aadpodidentity.k8s.io
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  group: aadpodidentity.k8s.io
  version: v1
  names:
    kind: AzureAssignedIdentity
    plural: azureassignedidentities
  scope: Namespaced
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: azureidentitybindings.aadpodidentity.k8s.io
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  group: aadpodidentity.k8s.io
  version: v1
  names:
    kind: AzureIdentityBinding
    plural: azureidentitybindings
  scope: Namespaced
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: azureidentities.aadpodidentity.k8s.io
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  group: aadpodidentity.k8s.io
  version: v1
  names:
    kind: AzureIdentity
    singular: azureidentity
    plural: azureidentities
  scope: Namespaced
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: aad-pod-id-nmi-role
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups: ["*"]
  resources: ["*"]
  verbs: ["get", "list"]
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: aad-pod-id-nmi-binding
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
    k8s-app: aad-pod-id-nmi-binding
subjects:
- kind: ServiceAccount
  name: aad-pod-id-nmi-service-account
  namespace: default
roleRef:
  kind: ClusterRole
  name: aad-pod-id-nmi-role
  apiGroup: rbac.authorization.k8s.io
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
    component: nmi
    tier: node
    k8s-app: aad-pod-id
  name: nmi
  namespace: default
spec:
  selector:
    matchLabels:
      component: nmi
      tier: node
  template:
    metadata:
      labels:
        component: nmi
        tier: node
    spec:
      serviceAccountName: aad-pod-id-nmi-service-account
      hostNetwork: true
      containers:
      - name: nmi
        image: {{ContainerImage "nmi"}}
        imagePullPolicy: IfNotPresent
        resources:
          requests:
            cpu: {{ContainerCPUReqs "nmi"}}
            memory: {{ContainerMemReqs "nmi"}}
          limits:
            cpu: {{ContainerCPULimits "nmi"}}
            memory: {{ContainerMemLimits "nmi"}}
        args:
          - "--host-ip=$(HOST_IP)"
          - "--node=$(NODE_NAME)"
        env:
          - name: HOST_IP
            valueFrom:
              fieldRef:
                fieldPath: status.podIP
          - name: NODE_NAME
            valueFrom:
              fieldRef:
                fieldPath: spec.nodeName
        securityContext:
          privileged: true
          capabilities:
            add:
            - NET_ADMIN
      nodeSelector:
        beta.kubernetes.io/os: linux
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: aad-pod-id-mic-service-account
  namespace: default
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: aad-pod-id-mic-role
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups: ["apiextensions.k8s.io"]
  resources: ["customresourcedefinitions"]
  verbs: ["*"]
- apiGroups: [""]
  resources: ["pods"]
  verbs: [ "list", "watch" ]
- apiGroups: [""]
  resources: ["events"]
  verbs: ["create", "patch"]
- apiGroups: ["aadpodidentity.k8s.io"]
  resources: ["azureidentitybindings", "azureidentities"]
  verbs: ["get", "list", "watch", "post"]
- apiGroups: ["aadpodidentity.k8s.io"]
  resources: ["azureassignedidentities"]
  verbs: ["*"]
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: aad-pod-id-mic-binding
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
    k8s-app: aad-pod-id-mic-binding
subjects:
- kind: ServiceAccount
  name: aad-pod-id-mic-service-account
  namespace: default
roleRef:
  kind: ClusterRole
  name: aad-pod-id-mic-role
  apiGroup: rbac.authorization.k8s.io
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    component: mic
    k8s-app: aad-pod-id
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: mic
  namespace: default
spec:
  selector:
    matchLabels:
      component: mic
  template:
    metadata:
      labels:
        component: mic
    spec:
      serviceAccountName: aad-pod-id-mic-service-account
      containers:
      - name: mic
        image: {{ContainerImage "mic"}}
        imagePullPolicy: IfNotPresent
        resources:
          requests:
            cpu: {{ContainerCPUReqs "mic"}}
            memory: {{ContainerMemReqs "mic"}}
          limits:
            cpu: {{ContainerCPULimits "mic"}}
            memory: {{ContainerMemLimits "mic"}}
        args:
          - --cloudconfig=/etc/kubernetes/azure.json
          - --logtostderr
        volumeMounts:
          - name: k8s-azure-file
            mountPath: /etc/kubernetes/azure.json
            readOnly: true
      volumes:
      - name: k8s-azure-file
        hostPath:
          path: /etc/kubernetes/azure.json
`)

func k8sContaineraddons116KubernetesmasteraddonsAadPodIdentityDeploymentYamlBytes() ([]byte, error) {
	return _k8sContaineraddons116KubernetesmasteraddonsAadPodIdentityDeploymentYaml, nil
}

func k8sContaineraddons116KubernetesmasteraddonsAadPodIdentityDeploymentYaml() (*asset, error) {
	bytes, err := k8sContaineraddons116KubernetesmasteraddonsAadPodIdentityDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/1.16/kubernetesmasteraddons-aad-pod-identity-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddons116KubernetesmasteraddonsAciConnectorDeploymentYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: aci-connector
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: aci-connector
  labels:
    app: aci-connector
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups:
  - ""
  resources:
  - configmaps
  - pods
  - services
  - endpoints
  - events
  - secrets
  - nodes
  - nodes/status
  - pods/status
  verbs:
  - "*"
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: aci-connector
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: aci-connector
subjects:
- kind: ServiceAccount
  name: aci-connector
  namespace: kube-system
---
apiVersion: v1
kind: Secret
metadata:
  name: aci-connector-secret
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
type: Opaque
data:
  credentials.json: <creds>
  cert.pem: <cert>
  key.pem: <key>
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: aci-connector
  namespace: kube-system
  labels:
    app: aci-connector
    name: aci-connector
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  replicas: 1
  selector:
    matchLabels:
      app: aci-connector
  template:
    metadata:
      labels:
        app: aci-connector
    spec:
      serviceAccountName: aci-connector
      nodeSelector:
        beta.kubernetes.io/os: linux
      containers:
      - name: aci-connector
        image: {{ContainerImage "aci-connector"}}
        imagePullPolicy: Always
        env:
        - name: KUBELET_PORT
          value: "10250"
        - name: AZURE_AUTH_LOCATION
          value: /etc/virtual-kubelet/credentials.json
        - name: ACI_RESOURCE_GROUP
          value: <rgName>
        - name: ACI_REGION
          value: <region>
        - name: APISERVER_CERT_LOCATION
          value: /etc/virtual-kubelet/cert.pem
        - name: APISERVER_KEY_LOCATION
          value: /etc/virtual-kubelet/key.pem
        - name: VKUBELET_POD_IP
          valueFrom:
            fieldRef:
              fieldPath: status.podIP
        resources:
          requests:
            cpu: {{ContainerCPUReqs "aci-connector"}}
            memory: {{ContainerMemReqs "aci-connector"}}
          limits:
            cpu: {{ContainerCPULimits "aci-connector"}}
            memory: {{ContainerMemLimits "aci-connector"}}
        volumeMounts:
        - name: credentials
          mountPath: "/etc/virtual-kubelet"
          readOnly: true
        command: ["virtual-kubelet"]
        args: ["--provider", "azure", "--nodename", "{{ContainerConfig "nodeName"}}" , "--os", "{{ContainerConfig "os"}}", "--taint", "{{ContainerConfig "taint"}}"]
      volumes:
      - name: credentials
        secret:
          secretName: aci-connector-secret
#EOF
`)

func k8sContaineraddons116KubernetesmasteraddonsAciConnectorDeploymentYamlBytes() ([]byte, error) {
	return _k8sContaineraddons116KubernetesmasteraddonsAciConnectorDeploymentYaml, nil
}

func k8sContaineraddons116KubernetesmasteraddonsAciConnectorDeploymentYaml() (*asset, error) {
	bytes, err := k8sContaineraddons116KubernetesmasteraddonsAciConnectorDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/1.16/kubernetesmasteraddons-aci-connector-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddons116KubernetesmasteraddonsAzureNpmDaemonsetYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: azure-npm
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: azure-npm
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
  - apiGroups:
    - ""
    resources:
      - pods
      - nodes
      - namespaces
    verbs:
      - get
      - list
      - watch
  - apiGroups:
    - networking.k8s.io
    resources:
      - networkpolicies
    verbs:
      - get
      - list
      - watch
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: azure-npm-binding
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
subjects:
  - kind: ServiceAccount
    name: azure-npm
    namespace: kube-system
roleRef:
  kind: ClusterRole
  name: azure-npm
  apiGroup: rbac.authorization.k8s.io
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: azure-npm
  namespace: kube-system
  labels:
    app: azure-npm
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  selector:
    matchLabels:
      k8s-app: azure-npm
  template:
    metadata:
      labels:
        k8s-app: azure-npm
    spec:
      priorityClassName: system-node-critical
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      nodeSelector:
        beta.kubernetes.io/os: linux
      containers:
        - name: azure-npm
          image: {{ContainerImage "azure-npm-daemonset"}}
          securityContext:
            privileged: true
          env:
            - name: HOSTNAME
              valueFrom:
                fieldRef:
                  apiVersion: v1
                  fieldPath: spec.nodeName
          volumeMounts:
          - name: xtables-lock
            mountPath: /run/xtables.lock
          - name: log
            mountPath: /var/log
      hostNetwork: true
      volumes:
      - name: log
        hostPath:
          path: /var/log
          type: Directory
      - name: xtables-lock
        hostPath:
          path: /run/xtables.lock
          type: File
      serviceAccountName: azure-npm
`)

func k8sContaineraddons116KubernetesmasteraddonsAzureNpmDaemonsetYamlBytes() ([]byte, error) {
	return _k8sContaineraddons116KubernetesmasteraddonsAzureNpmDaemonsetYaml, nil
}

func k8sContaineraddons116KubernetesmasteraddonsAzureNpmDaemonsetYaml() (*asset, error) {
	bytes, err := k8sContaineraddons116KubernetesmasteraddonsAzureNpmDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/1.16/kubernetesmasteraddons-azure-npm-daemonset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddons116KubernetesmasteraddonsBlobfuseFlexvolumeInstallerYaml = []byte(`apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: blobfuse-flexvol-installer
  namespace: kube-system
  labels:
    k8s-app: blobfuse
    kubernetes.io/cluster-service: "true"
spec:
  selector:
    matchLabels:
      name: blobfuse
  template:
    metadata:
      labels:
        name: blobfuse
        kubernetes.io/cluster-service: "true"
    spec:
      containers:
      - name: blobfuse-flexvol-installer
        image: {{ContainerImage "blobfuse-flexvolume"}}
        imagePullPolicy: IfNotPresent
        resources:
          requests:
            cpu: {{ContainerCPUReqs "blobfuse-flexvolume"}}
            memory: {{ContainerMemReqs "blobfuse-flexvolume"}}
          limits:
            cpu: {{ContainerCPULimits "blobfuse-flexvolume"}}
            memory: {{ContainerMemLimits "blobfuse-flexvolume"}}
        volumeMounts:
        - name: volplugins
          mountPath: /etc/kubernetes/volumeplugins/
        - name: varlog
          mountPath: /var/log/
      volumes:
      - name: varlog
        hostPath:
          path: /var/log/
      - name: volplugins
        hostPath:
          path: /etc/kubernetes/volumeplugins/
      nodeSelector:
        beta.kubernetes.io/os: linux
`)

func k8sContaineraddons116KubernetesmasteraddonsBlobfuseFlexvolumeInstallerYamlBytes() ([]byte, error) {
	return _k8sContaineraddons116KubernetesmasteraddonsBlobfuseFlexvolumeInstallerYaml, nil
}

func k8sContaineraddons116KubernetesmasteraddonsBlobfuseFlexvolumeInstallerYaml() (*asset, error) {
	bytes, err := k8sContaineraddons116KubernetesmasteraddonsBlobfuseFlexvolumeInstallerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/1.16/kubernetesmasteraddons-blobfuse-flexvolume-installer.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddons116KubernetesmasteraddonsCalicoDaemonsetYaml = []byte(`# Source: calico/templates/calico-config.yaml
# This ConfigMap is used to configure a self-hosted Calico installation.
kind: ConfigMap
apiVersion: v1
metadata:
  name: calico-config
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "EnsureExists"
data:
  # You must set a non-zero value for Typha replicas below.
  typha_service_name: "calico-typha"

  # The CNI network configuration to install on each node.  The special
  # values in this config will be automatically populated.
  cni_network_config: |-
    {
      "name": "k8s-pod-network",
      "cniVersion": "0.3.0",
      "plugins": [
        {
          "type": "calico",
          "log_level": "info",
          "datastore_type": "kubernetes",
          "nodename": "__KUBERNETES_NODE_NAME__",
          "mtu": 1500,
          "ipam": <calicoIPAMConfig>,
          "policy": {
              "type": "k8s"
          },
          "kubernetes": {
              "kubeconfig": "__KUBECONFIG_FILEPATH__"
          }
        },
        {
          "type": "portmap",
          "snat": true,
          "capabilities": {"portMappings": true}
        }
      ]
    }

---
# Source: calico/templates/kdd-crds.yaml
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: felixconfigurations.crd.projectcalico.org
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: FelixConfiguration
    plural: felixconfigurations
    singular: felixconfiguration
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: bgpconfigurations.crd.projectcalico.org
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: BGPConfiguration
    plural: bgpconfigurations
    singular: bgpconfiguration
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: ippools.crd.projectcalico.org
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: IPPool
    plural: ippools
    singular: ippool
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: hostendpoints.crd.projectcalico.org
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: HostEndpoint
    plural: hostendpoints
    singular: hostendpoint
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: clusterinformations.crd.projectcalico.org
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: ClusterInformation
    plural: clusterinformations
    singular: clusterinformation
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: globalnetworkpolicies.crd.projectcalico.org
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: GlobalNetworkPolicy
    plural: globalnetworkpolicies
    singular: globalnetworkpolicy
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: globalnetworksets.crd.projectcalico.org
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: GlobalNetworkSet
    plural: globalnetworksets
    singular: globalnetworkset
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: networkpolicies.crd.projectcalico.org
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  scope: Namespaced
  group: crd.projectcalico.org
  version: v1
  names:
    kind: NetworkPolicy
    plural: networkpolicies
    singular: networkpolicy
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: networksets.crd.projectcalico.org
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  scope: Namespaced
  group: crd.projectcalico.org
  version: v1
  names:
    kind: NetworkSet
    plural: networksets
    singular: networkset
---
# Source: calico/templates/rbac.yaml

# Include a clusterrole for the calico-node DaemonSet,
# and bind it to the calico-node serviceaccount.
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: calico-node
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
rules:
# The CNI plugin needs to get pods, nodes, and namespaces.
- apiGroups: [""]
  resources:
  - pods
  - nodes
  - namespaces
  verbs:
  - get
- apiGroups: [""]
  resources:
  - endpoints
  - services
  verbs:
  # Used to discover service IPs for advertisement.
  - watch
  - list
  # Used to discover Typhas.
  - get
- apiGroups: [""]
  resources:
  - nodes/status
  verbs:
  # Needed for clearing NodeNetworkUnavailable flag.
  - patch
  # Calico stores some configuration information in node annotations.
  - update
# Watch for changes to Kubernetes NetworkPolicies.
- apiGroups: ["networking.k8s.io"]
  resources:
  - networkpolicies
  verbs:
  - watch
  - list
# Used by Calico for policy information.
- apiGroups: [""]
  resources:
  - pods
  - namespaces
  - serviceaccounts
  verbs:
  - list
  - watch
# The CNI plugin patches pods/status.
- apiGroups: [""]
  resources:
  - pods/status
  verbs:
  - patch
# Calico monitors various CRDs for config.
- apiGroups: ["crd.projectcalico.org"]
  resources:
  - globalfelixconfigs
  - felixconfigurations
  - bgppeers
  - globalbgpconfigs
  - bgpconfigurations
  - ippools
  - ipamblocks
  - globalnetworkpolicies
  - globalnetworksets
  - networkpolicies
  - networksets
  - clusterinformations
  - hostendpoints
  verbs:
  - get
  - list
  - watch
# Calico must create and update some CRDs on startup.
- apiGroups: ["crd.projectcalico.org"]
  resources:
  - ippools
  - felixconfigurations
  - clusterinformations
  verbs:
  - create
  - update
# Calico stores some configuration information on the node.
- apiGroups: [""]
  resources:
  - nodes
  verbs:
  - get
  - list
  - watch
# These permissions are only requried for upgrade from v2.6, and can
# be removed after upgrade or on fresh installations.
- apiGroups: ["crd.projectcalico.org"]
  resources:
  - bgpconfigurations
  - bgppeers
  verbs:
  - create
  - update
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: calico-node
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: "Reconcile"
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: calico-node
subjects:
- kind: ServiceAccount
  name: calico-node
  namespace: kube-system

---
# Source: calico/templates/calico-typha.yaml
# This manifest creates a Service, which will be backed by Calico's Typha daemon.
# Typha sits in between Felix and the API server, reducing Calico's load on the API server.

apiVersion: v1
kind: Service
metadata:
  name: calico-typha
  namespace: kube-system
  labels:
    k8s-app: calico-typha
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  ports:
  - port: 5473
    protocol: TCP
    targetPort: calico-typha
    name: calico-typha
  selector:
    k8s-app: calico-typha
---

# This manifest creates a Deployment of Typha to back the above service.

apiVersion: apps/v1
kind: Deployment
metadata:
  name: calico-typha
  namespace: kube-system
  labels:
    k8s-app: calico-typha
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  # Number of Typha replicas.  To enable Typha, set this to a non-zero value *and* set the
  # typha_service_name variable in the calico-config ConfigMap above.
  #
  # We recommend using Typha if you have more than 50 nodes.  Above 100 nodes it is essential
  # (when using the Kubernetes datastore).  Use one replica for every 100-200 nodes.  In
  # production, we recommend running at least 3 replicas to reduce the impact of rolling upgrade.
  replicas: 1
  revisionHistoryLimit: 2
  selector:
    matchLabels:
      k8s-app: calico-typha
  template:
    metadata:
      labels:
        k8s-app: calico-typha
      annotations:
        cluster-autoscaler.kubernetes.io/safe-to-evict: 'true'
    spec:
      priorityClassName: system-cluster-critical
      nodeSelector:
        beta.kubernetes.io/os: linux
      hostNetwork: true
      tolerations:
      # Mark the pod as a critical add-on for rescheduling.
      - key: CriticalAddonsOnly
        operator: Exists
      # Since Calico can't network a pod until Typha is up, we need to run Typha itself
      # as a host-networked pod.
      serviceAccountName: calico-node
      containers:
      - image: {{ContainerImage "calico-typha"}}
        name: calico-typha
        ports:
        - containerPort: 5473
          name: calico-typha
          protocol: TCP
        env:
        # Enable "info" logging by default.  Can be set to "debug" to increase verbosity.
        - name: TYPHA_LOGSEVERITYSCREEN
          value: "info"
        # Disable logging to file and syslog since those don't make sense in Kubernetes.
        - name: TYPHA_LOGFILEPATH
          value: "none"
        - name: TYPHA_LOGSEVERITYSYS
          value: "none"
        # Monitor the Kubernetes API to find the number of running instances and rebalance
        # connections.
        - name: TYPHA_CONNECTIONREBALANCINGMODE
          value: "kubernetes"
        - name: TYPHA_DATASTORETYPE
          value: "kubernetes"
        - name: TYPHA_HEALTHENABLED
          value: "true"
        # Configure route aggregation based on pod CIDR.
        - name: USE_POD_CIDR
          value: "true"
        - name: FELIX_INTERFACEPREFIX
          value: "azv"
        # Uncomment these lines to enable prometheus metrics.  Since Typha is host-networked,
        # this opens a port on the host, which may need to be secured.
        #- name: TYPHA_PROMETHEUSMETRICSENABLED
        #  value: "true"
        #- name: TYPHA_PROMETHEUSMETRICSPORT
        #  value: "9093"
        livenessProbe:
          httpGet:
            path: /liveness
            port: 9098
            host: localhost
          periodSeconds: 30
          initialDelaySeconds: 30
        readinessProbe:
          httpGet:
            path: /readiness
            port: 9098
            host: localhost
          periodSeconds: 10
---
# Source: calico/templates/calico-node.yaml
# This manifest installs the calico-node container, as well
# as the CNI plugins and network config on
# each master and worker node in a Kubernetes cluster.
kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: calico-node
  namespace: kube-system
  labels:
    k8s-app: calico-node
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  selector:
    matchLabels:
      k8s-app: calico-node
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 1
  template:
    metadata:
      labels:
        k8s-app: calico-node
    spec:
      priorityClassName: system-cluster-critical
      nodeSelector:
        beta.kubernetes.io/os: linux
      hostNetwork: true
      tolerations:
      # Make sure calico-node gets scheduled on all nodes.
      - effect: NoSchedule
        operator: Exists
      # Mark the pod as a critical add-on for rescheduling.
      - key: CriticalAddonsOnly
        operator: Exists
      - effect: NoExecute
        operator: Exists
      serviceAccountName: calico-node
      # Minimize downtime during a rolling upgrade or deletion; tell Kubernetes to do a "force
      # deletion": https://kubernetes.io/docs/concepts/workloads/pods/pod/#termination-of-pods.
      terminationGracePeriodSeconds: 0
      initContainers:
      # This container installs the CNI binaries
      # and CNI network config file on each node.
      - name: install-cni
        image: {{ContainerImage "calico-cni"}}
        command: ["/install-cni.sh"]
        env:
        # Name of the CNI config file to create.
        - name: CNI_CONF_NAME
          value: "10-calico.conflist"
        # The CNI network config to install on each node.
        - name: CNI_NETWORK_CONFIG
          valueFrom:
            configMapKeyRef:
              name: calico-config
              key: cni_network_config
        # Set the hostname based on the k8s node name.
        - name: KUBERNETES_NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        # Prevents the container from sleeping forever.
        - name: SLEEP
          value: "false"
        volumeMounts:
        - mountPath: /host/opt/cni/bin
          name: cni-bin-dir
        - mountPath: /host/etc/cni/net.d
          name: cni-net-dir
      containers:
      # Runs calico-node container on each Kubernetes node.  This
      # container programs network policy and routes on each
      # host.
      - name: calico-node
        image: {{ContainerImage "calico-node"}}
        env:
        # Use Kubernetes API as the backing datastore.
        - name: DATASTORE_TYPE
          value: "kubernetes"
        # Configure route aggregation based on pod CIDR.
        - name: USE_POD_CIDR
          value: "true"
        # Typha support: controlled by the ConfigMap.
        - name: FELIX_TYPHAK8SSERVICENAME
          valueFrom:
            configMapKeyRef:
              name: calico-config
              key: typha_service_name
        # Wait for the datastore.
        - name: WAIT_FOR_DATASTORE
          value: "true"
        # Set based on the k8s node name.
        - name: NODENAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        # Don't enable BGP.
        - name: CALICO_NETWORKING_BACKEND
          value: "none"
        # Cluster type to identify the deployment type
        - name: CLUSTER_TYPE
          value: "k8s"
        # The default IPv4 pool to create on startup if none exists. Pod IPs will be
        # chosen from this range. Changing this value after installation will have
        # no effect. This should fall within ` + "`" + `--cluster-cidr` + "`" + `.
        - name: CALICO_IPV4POOL_CIDR
          value: "<kubeClusterCidr>"
        # Disable file logging so ` + "`" + `kubectl logs` + "`" + ` works.
        - name: CALICO_DISABLE_FILE_LOGGING
          value: "true"
        # Set Felix endpoint to host default action to ACCEPT.
        - name: FELIX_DEFAULTENDPOINTTOHOSTACTION
          value: "ACCEPT"
        # Disable IPv6 on Kubernetes.
        - name: FELIX_IPV6SUPPORT
          value: "false"
        # Set Felix logging to "info"
        - name: FELIX_LOGSEVERITYSCREEN
          value: "info"
        - name: FELIX_HEALTHENABLED
          value: "true"
        - name: CALICO_IPV4POOL_IPIP
          value: "off"
        - name: FELIX_INTERFACEPREFIX
          value: "azv"
        securityContext:
          privileged: true
        resources:
          requests:
            cpu: 250m
        livenessProbe:
          httpGet:
            path: /liveness
            port: 9099
            host: localhost
          periodSeconds: 10
          initialDelaySeconds: 10
          failureThreshold: 6
        readinessProbe:
          exec:
            command:
            - /bin/calico-node
            - -felix-ready
          periodSeconds: 10
        volumeMounts:
        - mountPath: /lib/modules
          name: lib-modules
          readOnly: true
        - mountPath: /run/xtables.lock
          name: xtables-lock
          readOnly: false
        - mountPath: /var/run/calico
          name: var-run-calico
          readOnly: false
        - mountPath: /var/lib/calico
          name: var-lib-calico
          readOnly: false
      volumes:
      # Used by calico-node.
      - name: lib-modules
        hostPath:
          path: /lib/modules
      - name: var-run-calico
        hostPath:
          path: /var/run/calico
      - name: var-lib-calico
        hostPath:
          path: /var/lib/calico
      - name: xtables-lock
        hostPath:
          path: /run/xtables.lock
          type: FileOrCreate
      # Used to install CNI.
      - name: cni-bin-dir
        hostPath:
          path: /opt/cni/bin
      - name: cni-net-dir
        hostPath:
          path: /etc/cni/net.d
---

apiVersion: v1
kind: ServiceAccount
metadata:
  name: calico-node
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
---

# Typha Horizontal Autoscaler ConfigMap
kind: ConfigMap
apiVersion: v1
metadata:
  name: calico-typha-horizontal-autoscaler
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: "EnsureExists"
data:
  ladder: |-
    {
      "coresToReplicas": [],
      "nodesToReplicas":
      [
        [1, 1],
        [10, 2],
        [100, 3],
        [250, 4],
        [500, 5],
        [1000, 6],
        [1500, 7],
        [2000, 8]
      ]
    }

---

# Typha Horizontal Autoscaler Deployment
apiVersion: apps/v1
kind: Deployment
metadata:
  name: calico-typha-horizontal-autoscaler
  namespace: kube-system
  labels:
    k8s-app: calico-typha-autoscaler
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  replicas: 1
  selector:
    matchLabels:
      k8s-app: calico-typha-autoscaler
  template:
    metadata:
      labels:
        k8s-app: calico-typha-autoscaler
    spec:
      priorityClassName: system-cluster-critical
      securityContext:
        supplementalGroups: [65534]
        fsGroup: 65534
      containers:
      - image: {{ContainerImage "calico-cluster-proportional-autoscaler"}}
        name: autoscaler
        command:
        - /cluster-proportional-autoscaler
        - --namespace=kube-system
        - --configmap=calico-typha-horizontal-autoscaler
        - --target=deployment/calico-typha
        - --logtostderr=true
        - --v=2
        resources:
          requests:
            cpu: 10m
          limits:
            cpu: 10m
      serviceAccountName: typha-cpha
---

# Typha Horizontal Autoscaler Cluster Role
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: typha-cpha
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: "Reconcile"
rules:
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["list"]

---

# Typha Horizontal Autoscaler Cluster Role Binding
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: typha-cpha
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: "Reconcile"
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: typha-cpha
subjects:
- kind: ServiceAccount
  name: typha-cpha
  namespace: kube-system
---

# Typha Horizontal Autoscaler Role
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: typha-cpha
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: "Reconcile"
rules:
- apiGroups: [""]
  resources: ["configmaps"]
  verbs: ["get"]
- apiGroups: ["extensions"]
  resources: ["deployments/scale"]
  verbs: ["get", "update"]

---

# Typha Horizontal Autoscaler Role Binding
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: typha-cpha
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: "Reconcile"
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: typha-cpha
subjects:
- kind: ServiceAccount
  name: typha-cpha
  namespace: kube-system
---

# Typha Horizontal Autoscaler Service Account
apiVersion: v1
kind: ServiceAccount
metadata:
  name: typha-cpha
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
`)

func k8sContaineraddons116KubernetesmasteraddonsCalicoDaemonsetYamlBytes() ([]byte, error) {
	return _k8sContaineraddons116KubernetesmasteraddonsCalicoDaemonsetYaml, nil
}

func k8sContaineraddons116KubernetesmasteraddonsCalicoDaemonsetYaml() (*asset, error) {
	bytes, err := k8sContaineraddons116KubernetesmasteraddonsCalicoDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/1.16/kubernetesmasteraddons-calico-daemonset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddons116KubernetesmasteraddonsClusterAutoscalerDeploymentYaml = []byte(`---
apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    k8s-addon: cluster-autoscaler.addons.k8s.io
    k8s-app: cluster-autoscaler
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: cluster-autoscaler
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: cluster-autoscaler
  labels:
    k8s-addon: cluster-autoscaler.addons.k8s.io
    k8s-app: cluster-autoscaler
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups: [""]
  resources: ["events", "endpoints"]
  verbs: ["create", "patch"]
- apiGroups: [""]
  resources: ["pods/eviction"]
  verbs: ["create"]
- apiGroups: [""]
  resources: ["pods/status"]
  verbs: ["update"]
- apiGroups: [""]
  resources: ["endpoints"]
  resourceNames: ["cluster-autoscaler"]
  verbs: ["get", "update"]
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["watch", "list", "get", "update"]
- apiGroups: [""]
  resources:
  - "pods"
  - "services"
  - "replicationcontrollers"
  - "persistentvolumeclaims"
  - "persistentvolumes"
  verbs: ["watch", "list", "get"]
- apiGroups: ["extensions"]
  resources: ["replicasets", "daemonsets"]
  verbs: ["watch", "list", "get"]
- apiGroups: ["policy"]
  resources: ["poddisruptionbudgets"]
  verbs: ["watch", "list"]
- apiGroups: ["apps"]
  resources: ["statefulsets","replicasets","daemonsets"]
  verbs: ["watch","list","get"]
- apiGroups: ["batch"]
  resources: ["jobs"]
  verbs: ["watch","list"]
- apiGroups: ["storage.k8s.io"]
  resources: ["storageclasses"]
  verbs: ["get", "list", "watch"]
- apiGroups: ["batch"]
  resources: ["jobs", "cronjobs"]
  verbs: ["watch", "list", "get"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: cluster-autoscaler
  namespace: kube-system
  labels:
    k8s-addon: cluster-autoscaler.addons.k8s.io
    k8s-app: cluster-autoscaler
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups: [""]
  resources: ["configmaps"]
  verbs: ["create", "list", "watch"]
- apiGroups: [""]
  resources: ["configmaps"]
  resourceNames:
  - "cluster-autoscaler-status"
  - "cluster-autoscaler-priority-expander"
  verbs: ["delete", "get", "update", "watch"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: cluster-autoscaler
  labels:
    k8s-addon: cluster-autoscaler.addons.k8s.io
    k8s-app: cluster-autoscaler
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-autoscaler
subjects:
  - kind: ServiceAccount
    name: cluster-autoscaler
    namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: cluster-autoscaler
  namespace: kube-system
  labels:
    k8s-addon: cluster-autoscaler.addons.k8s.io
    k8s-app: cluster-autoscaler
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: cluster-autoscaler
subjects:
  - kind: ServiceAccount
    name: cluster-autoscaler
    namespace: kube-system
---
apiVersion: v1
data:
  ClientID: <clientID>
  ClientSecret: <clientSec>
  ResourceGroup: <rg>
  SubscriptionID: <subID>
  TenantID: <tenantID>
  VMType: <vmType>
kind: Secret
metadata:
  name: cluster-autoscaler-azure
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: cluster-autoscaler
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: cluster-autoscaler
  namespace: kube-system
spec:
  replicas: 1
  selector:
    matchLabels:
      app: cluster-autoscaler
  template:
    metadata:
      labels:
        app: cluster-autoscaler
    spec:
      priorityClassName: system-node-critical
      <hostNet>
      serviceAccountName: cluster-autoscaler
      tolerations:
      - effect: NoSchedule
        operator: "Equal"
        value: "true"
        key: node-role.kubernetes.io/master
      nodeSelector:
        kubernetes.io/role: master
        beta.kubernetes.io/os: linux
      containers:
      - image: {{ContainerImage "cluster-autoscaler"}}
        imagePullPolicy: IfNotPresent
        name: cluster-autoscaler
        resources:
          limits:
            cpu: {{ContainerCPULimits "cluster-autoscaler"}}
            memory: {{ContainerMemLimits "cluster-autoscaler"}}
          requests:
            cpu: {{ContainerCPUReqs "cluster-autoscaler"}}
            memory: {{ContainerMemReqs "cluster-autoscaler"}}
        command:
        - ./cluster-autoscaler
        - --v=3
        - --logtostderr=true
        - --cloud-provider=azure
        - --skip-nodes-with-local-storage=false
        - --nodes={{ContainerConfig "min-nodes"}}:{{ContainerConfig "max-nodes"}}:<vmssName>
        - --scan-interval={{ContainerConfig "scan-interval"}}
        env:
        - name: ARM_CLOUD
          value: "<cloud>"
        - name: ARM_SUBSCRIPTION_ID
          valueFrom:
            secretKeyRef:
              key: SubscriptionID
              name: cluster-autoscaler-azure
        - name: ARM_RESOURCE_GROUP
          valueFrom:
            secretKeyRef:
              key: ResourceGroup
              name: cluster-autoscaler-azure
        - name: ARM_TENANT_ID
          valueFrom:
            secretKeyRef:
              key: TenantID
              name: cluster-autoscaler-azure
        - name: ARM_CLIENT_ID
          valueFrom:
            secretKeyRef:
              key: ClientID
              name: cluster-autoscaler-azure
        - name: ARM_CLIENT_SECRET
          valueFrom:
            secretKeyRef:
              key: ClientSecret
              name: cluster-autoscaler-azure
        - name: ARM_VM_TYPE
          valueFrom:
            secretKeyRef:
              key: VMType
              name: cluster-autoscaler-azure
        - name: ARM_USE_MANAGED_IDENTITY_EXTENSION
          value: "<useManagedIdentity>"
        volumeMounts:
        - mountPath: /etc/ssl/certs/ca-certificates.crt
          name: ssl-certs
          readOnly: true
        <volMounts>
      restartPolicy: Always
      volumes:
      - hostPath:
          path: /etc/ssl/certs/ca-certificates.crt
          type: ""
        name: ssl-certs
      <vols>
#EOF
`)

func k8sContaineraddons116KubernetesmasteraddonsClusterAutoscalerDeploymentYamlBytes() ([]byte, error) {
	return _k8sContaineraddons116KubernetesmasteraddonsClusterAutoscalerDeploymentYaml, nil
}

func k8sContaineraddons116KubernetesmasteraddonsClusterAutoscalerDeploymentYaml() (*asset, error) {
	bytes, err := k8sContaineraddons116KubernetesmasteraddonsClusterAutoscalerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/1.16/kubernetesmasteraddons-cluster-autoscaler-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddons116KubernetesmasteraddonsHeapsterDeploymentYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: heapster
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: system:heapster-with-nanny
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
rules:
- apiGroups:
  - extensions
  - apps
  resources:
  - deployments
  verbs:
  - get
  - list
  - watch
  - update
  - patch
- apiGroups:
  - ""
  resources:
  - events
  - namespaces
  - nodes
  - pods
  verbs:
  - get
  - list
  - watch
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: system:heapster-with-nanny
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:heapster-with-nanny
subjects:
- kind: ServiceAccount
  name: heapster
  namespace: kube-system
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: heapster-config
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
data:
  NannyConfiguration: |-
    apiVersion: nannyconfig/v1alpha1
    kind: NannyConfiguration
---
apiVersion: v1
kind: Service
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
    kubernetes.io/name: Heapster
  name: heapster
  namespace: kube-system
spec:
  ports:
  - port: 80
    targetPort: 8082
  selector:
    k8s-app: heapster
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: heapster
  namespace: kube-system
  labels:
    k8s-app: heapster
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
spec:
  replicas: 1
  selector:
    matchLabels:
      k8s-app: heapster
  template:
    metadata:
      labels:
        k8s-app: heapster
    spec:
      priorityClassName: system-node-critical
      containers:
        - image: {{ContainerImage "heapster"}}
          imagePullPolicy: IfNotPresent
          name: heapster
          resources:
            requests:
              cpu: {{ContainerCPUReqs "heapster"}}
              memory: {{ContainerMemReqs "heapster"}}
            limits:
              cpu: {{ContainerCPULimits "heapster"}}
              memory: {{ContainerMemLimits "heapster"}}
          livenessProbe:
            httpGet:
              path: /healthz
              port: 8082
              scheme: HTTP
            initialDelaySeconds: 180
            timeoutSeconds: 5
          command:
            - /heapster
            - --source=kubernetes.summary_api:''
        - image: {{ContainerImage "heapster-nanny"}}
          imagePullPolicy: IfNotPresent
          name: heapster-nanny
          resources:
            requests:
              cpu: {{ContainerCPUReqs "heapster-nanny"}}
              memory: {{ContainerMemReqs "heapster-nanny"}}
            limits:
              cpu: {{ContainerCPULimits "heapster-nanny"}}
              memory: {{ContainerMemLimits "heapster-nanny"}}
          env:
            - name: MY_POD_NAME
              valueFrom:
                fieldRef:
                  fieldPath: metadata.name
            - name: MY_POD_NAMESPACE
              valueFrom:
                fieldRef:
                  fieldPath: metadata.namespace
          volumeMounts:
          - name: heapster-config-volume
            mountPath: /etc/config
          command:
            - /pod_nanny
            - --config-dir=/etc/config
            - --cpu=80m
            - --extra-cpu=0.5m
            - --memory=140Mi
            - --extra-memory=4Mi
            - --threshold=5
            - --deployment=heapster
            - --container=heapster
            - --poll-period=300000
            - --estimator=exponential
      volumes:
        - name: heapster-config-volume
          configMap:
            name: heapster-config
      serviceAccountName: heapster
      tolerations:
        - key: CriticalAddonsOnly
          operator: Exists
      nodeSelector:
        beta.kubernetes.io/os: linux
`)

func k8sContaineraddons116KubernetesmasteraddonsHeapsterDeploymentYamlBytes() ([]byte, error) {
	return _k8sContaineraddons116KubernetesmasteraddonsHeapsterDeploymentYaml, nil
}

func k8sContaineraddons116KubernetesmasteraddonsHeapsterDeploymentYaml() (*asset, error) {
	bytes, err := k8sContaineraddons116KubernetesmasteraddonsHeapsterDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/1.16/kubernetesmasteraddons-heapster-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddons116KubernetesmasteraddonsKeyvaultFlexvolumeInstallerYaml = []byte(`apiVersion: apps/v1
kind: DaemonSet
metadata:
  labels:
    app: keyvault-flexvolume
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: keyvault-flexvolume
  namespace: kube-system
spec:
  updateStrategy:
    type: RollingUpdate
  selector:
    matchLabels:
      app: keyvault-flexvolume
  template:
    metadata:
      labels:
        app: keyvault-flexvolume
        kubernetes.io/cluster-service: "true"
        addonmanager.kubernetes.io/mode: Reconcile
    spec:
      tolerations:
      containers:
      - name: keyvault-flexvolume
        image: {{ContainerImage "keyvault-flexvolume"}}
        imagePullPolicy: IfNotPresent
        resources:
          requests:
            cpu: {{ContainerCPUReqs "keyvault-flexvolume"}}
            memory: {{ContainerMemReqs "keyvault-flexvolume"}}
          limits:
            cpu: {{ContainerCPULimits "keyvault-flexvolume"}}
            memory: {{ContainerMemLimits "keyvault-flexvolume"}}
        env:
        - name: TARGET_DIR
          value: /etc/kubernetes/volumeplugins
        volumeMounts:
        - mountPath: /etc/kubernetes/volumeplugins
          name: volplugins
      volumes:
      - hostPath:
          path: /etc/kubernetes/volumeplugins
        name: volplugins
      nodeSelector:
        beta.kubernetes.io/os: linux
`)

func k8sContaineraddons116KubernetesmasteraddonsKeyvaultFlexvolumeInstallerYamlBytes() ([]byte, error) {
	return _k8sContaineraddons116KubernetesmasteraddonsKeyvaultFlexvolumeInstallerYaml, nil
}

func k8sContaineraddons116KubernetesmasteraddonsKeyvaultFlexvolumeInstallerYaml() (*asset, error) {
	bytes, err := k8sContaineraddons116KubernetesmasteraddonsKeyvaultFlexvolumeInstallerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/1.16/kubernetesmasteraddons-keyvault-flexvolume-installer.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddons116KubernetesmasteraddonsKubeReschedulerDeploymentYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  name: rescheduler
  namespace: kube-system
  labels:
    k8s-app: rescheduler
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  replicas: 1
  selector:
    matchLabels:
      k8s-app: rescheduler
  template:
    metadata:
      labels:
        k8s-app: rescheduler
    spec:
      priorityClassName: system-node-critical
      nodeSelector:
        beta.kubernetes.io/os: linux
      containers:
      - image: {{ContainerImage "rescheduler"}}
        imagePullPolicy: IfNotPresent
        name: rescheduler
        resources:
          requests:
            cpu: {{ContainerCPUReqs "rescheduler"}}
            memory: {{ContainerMemReqs "rescheduler"}}
          limits:
            cpu: {{ContainerCPULimits "rescheduler"}}
            memory: {{ContainerMemLimits "rescheduler"}}
        command:
        - sh
        - -c
        - '/rescheduler'
`)

func k8sContaineraddons116KubernetesmasteraddonsKubeReschedulerDeploymentYamlBytes() ([]byte, error) {
	return _k8sContaineraddons116KubernetesmasteraddonsKubeReschedulerDeploymentYaml, nil
}

func k8sContaineraddons116KubernetesmasteraddonsKubeReschedulerDeploymentYaml() (*asset, error) {
	bytes, err := k8sContaineraddons116KubernetesmasteraddonsKubeReschedulerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/1.16/kubernetesmasteraddons-kube-rescheduler-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddons116KubernetesmasteraddonsKubernetesDashboardDeploymentYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    k8s-app: kubernetes-dashboard
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: kubernetes-dashboard
  namespace: kube-system
---
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: kubernetes-dashboard-minimal
  namespace: kube-system
  labels:
    k8s-app: kubernetes-dashboard
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups: [""]
  resources: ["secrets"]
  verbs: ["create"]
- apiGroups: [""]
  resources: ["configmaps"]
  verbs: ["create"]
- apiGroups: [""]
  resources: ["secrets"]
  resourceNames: ["kubernetes-dashboard-key-holder"]
  verbs: ["get", "update", "delete"]
- apiGroups: [""]
  resources: ["configmaps"]
  resourceNames: ["kubernetes-dashboard-settings"]
  verbs: ["get", "update"]
- apiGroups: [""]
  resources: ["services"]
  resourceNames: ["heapster"]
  verbs: ["proxy"]
- apiGroups: [""]
  resources: ["services/proxy"]
  resourceNames: ["heapster", "http:heapster:", "https:heapster:"]
  verbs: ["get"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: kubernetes-dashboard-minimal
  namespace: kube-system
  labels:
    k8s-app: kubernetes-dashboard
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: kubernetes-dashboard-minimal
subjects:
- kind: ServiceAccount
  name: kubernetes-dashboard
  namespace: kube-system
---
apiVersion: v1
kind: Service
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kube-system
spec:
  ports:
  - port: 443
    targetPort: 8443
  selector:
    k8s-app: kubernetes-dashboard
  type: NodePort
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kube-system
spec:
  replicas: 1
  selector:
    matchLabels:
      k8s-app: kubernetes-dashboard
  template:
    metadata:
      labels:
        k8s-app: kubernetes-dashboard
    spec:
      containers:
      - args:
        - --auto-generate-certificates
        - --heapster-host=http://heapster.kube-system:80
        image: {{ContainerImage "kubernetes-dashboard"}}
        imagePullPolicy: IfNotPresent
        livenessProbe:
          httpGet:
            path: "/"
            port: 8443
            scheme: HTTPS
          initialDelaySeconds: 30
          timeoutSeconds: 30
        name: kubernetes-dashboard
        ports:
        - containerPort: 8443
          protocol: TCP
        resources:
          requests:
            cpu: {{ContainerCPUReqs "kubernetes-dashboard"}}
            memory: {{ContainerMemReqs "kubernetes-dashboard"}}
          limits:
            cpu: {{ContainerCPULimits "kubernetes-dashboard"}}
            memory: {{ContainerMemLimits "kubernetes-dashboard"}}
        volumeMounts:
         - name: kubernetes-dashboard-certs
           mountPath: /certs
      volumes:
        - name: kubernetes-dashboard-certs
          emptyDir: {}
      serviceAccountName: kubernetes-dashboard
      nodeSelector:
        beta.kubernetes.io/os: linux
`)

func k8sContaineraddons116KubernetesmasteraddonsKubernetesDashboardDeploymentYamlBytes() ([]byte, error) {
	return _k8sContaineraddons116KubernetesmasteraddonsKubernetesDashboardDeploymentYaml, nil
}

func k8sContaineraddons116KubernetesmasteraddonsKubernetesDashboardDeploymentYaml() (*asset, error) {
	bytes, err := k8sContaineraddons116KubernetesmasteraddonsKubernetesDashboardDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/1.16/kubernetesmasteraddons-kubernetes-dashboard-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddons116KubernetesmasteraddonsMetricsServerDeploymentYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: metrics-server
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: system:metrics-server
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups:
  - ""
  resources:
  - pods
  - nodes
  - nodes/stats
  - namespaces
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - "extensions"
  resources:
  - deployments
  verbs:
  - get
  - list
  - watch
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system:metrics-server
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:metrics-server
subjects:
- kind: ServiceAccount
  name: metrics-server
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: RoleBinding
metadata:
  name: metrics-server-auth-reader
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: extension-apiserver-authentication-reader
subjects:
- kind: ServiceAccount
  name: metrics-server
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: metrics-server:system:auth-delegator
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:auth-delegator
subjects:
- kind: ServiceAccount
  name: metrics-server
  namespace: kube-system
---
apiVersion: v1
kind: Service
metadata:
  name: metrics-server
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
    kubernetes.io/name: "Metrics-server"
    kubernetes.io/cluster-service: "true"
spec:
  selector:
    k8s-app: metrics-server
  ports:
  - port: 443
    protocol: TCP
    targetPort: 443
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: metrics-server
  namespace: kube-system
  labels:
    k8s-app: metrics-server
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  selector:
    matchLabels:
      k8s-app: metrics-server
  template:
    metadata:
      name: metrics-server
      labels:
        k8s-app: metrics-server
    spec:
      serviceAccountName: metrics-server
      containers:
      - name: metrics-server
        image: {{ContainerImage "metrics-server"}}
        imagePullPolicy: IfNotPresent
        command:
        - /metrics-server
        - --source=kubernetes.summary_api:''
      nodeSelector:
        beta.kubernetes.io/os: linux
---
apiVersion: apiregistration.k8s.io/v1beta1
kind: APIService
metadata:
  name: v1beta1.metrics.k8s.io
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  service:
    name: metrics-server
    namespace: kube-system
  group: metrics.k8s.io
  version: v1beta1
  insecureSkipTLSVerify: true
  groupPriorityMinimum: 100
  versionPriority: 100
`)

func k8sContaineraddons116KubernetesmasteraddonsMetricsServerDeploymentYamlBytes() ([]byte, error) {
	return _k8sContaineraddons116KubernetesmasteraddonsMetricsServerDeploymentYaml, nil
}

func k8sContaineraddons116KubernetesmasteraddonsMetricsServerDeploymentYaml() (*asset, error) {
	bytes, err := k8sContaineraddons116KubernetesmasteraddonsMetricsServerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/1.16/kubernetesmasteraddons-metrics-server-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddons116KubernetesmasteraddonsNvidiaDevicePluginDaemonsetYaml = []byte(`apiVersion: apps/v1
kind: DaemonSet
metadata:
  labels:
    k8s-app: nvidia-device-plugin
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: nvidia-device-plugin
  namespace: kube-system
spec:
  selector:
    matchLabels:
      k8s-app: nvidia-device-plugin
  updateStrategy:
    type: RollingUpdate
  template:
    metadata:
      labels:
        k8s-app: nvidia-device-plugin
    spec:
      priorityClassName: system-node-critical
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: accelerator
                operator: In
                values:
                - nvidia
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      - key: nvidia.com/gpu
        effect: NoSchedule
        operator: Equal
        value: "true"
      containers:
      - image: {{ContainerImage "nvidia-device-plugin"}}
        name: nvidia-device-plugin-ctr
        resources:
          requests:
            cpu: {{ContainerCPUReqs "nvidia-device-plugin"}}
            memory: {{ContainerMemReqs "nvidia-device-plugin"}}
          limits:
            cpu: {{ContainerCPULimits "nvidia-device-plugin"}}
            memory: {{ContainerMemLimits "nvidia-device-plugin"}}
        securityContext:
          allowPrivilegeEscalation: false
          capabilities:
            drop: ["ALL"]
        volumeMounts:
          - name: device-plugin
            mountPath: /var/lib/kubelet/device-plugins
      volumes:
        - name: device-plugin
          hostPath:
            path: /var/lib/kubelet/device-plugins
      nodeSelector:
        beta.kubernetes.io/os: linux
        accelerator: nvidia
`)

func k8sContaineraddons116KubernetesmasteraddonsNvidiaDevicePluginDaemonsetYamlBytes() ([]byte, error) {
	return _k8sContaineraddons116KubernetesmasteraddonsNvidiaDevicePluginDaemonsetYaml, nil
}

func k8sContaineraddons116KubernetesmasteraddonsNvidiaDevicePluginDaemonsetYaml() (*asset, error) {
	bytes, err := k8sContaineraddons116KubernetesmasteraddonsNvidiaDevicePluginDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/1.16/kubernetesmasteraddons-nvidia-device-plugin-daemonset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddons116KubernetesmasteraddonsOmsagentDaemonsetYaml = []byte(`apiVersion: v1
kind: Secret
metadata:
  name: omsagent-secret
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
type: Opaque
data:
  WSID: "{{ContainerConfig "workspaceGuid"}}"
  KEY: "{{ContainerConfig "workspaceKey"}}"
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: omsagent
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: omsagent-reader
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups: [""]
  resources: ["pods", "events", "nodes", "namespaces", "services"]
  verbs: ["list", "get", "watch"]
- nonResourceURLs: ["/metrics"]
  verbs: ["get"]
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: omsagentclusterrolebinding
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
subjects:
  - kind: ServiceAccount
    name: omsagent
    namespace: kube-system
roleRef:
  kind: ClusterRole
  name: omsagent-reader
  apiGroup: rbac.authorization.k8s.io
---
kind: ConfigMap
apiVersion: v1
data:
  kube.conf: |-
     # Fluentd config file for OMS Docker - cluster components (kubeAPI)
     #Kubernetes pod inventory
     <source>
      type kubepodinventory
      tag oms.containerinsights.KubePodInventory
      run_interval 60s
      log_level debug
     </source>

     #Kubernetes events
     <source>
      type kubeevents
      tag oms.containerinsights.KubeEvents
      run_interval 60s
      log_level debug
      </source>

     #Kubernetes logs
     <source>
      type kubelogs
      tag oms.api.KubeLogs
      run_interval 60s
     </source>

     #Kubernetes services
     <source>
      type kubeservices
      tag oms.containerinsights.KubeServices
      run_interval 60s
      log_level debug
     </source>

     #Kubernetes Nodes
     <source>
      type kubenodeinventory
      tag oms.containerinsights.KubeNodeInventory
      run_interval 60s
      log_level debug
     </source>

     #Kubernetes perf
     <source>
      type kubeperf
      tag oms.api.KubePerf
      run_interval 60s
      log_level debug
     </source>

     #cadvisor perf- Windows nodes
     <source>
      type wincadvisorperf
      tag oms.api.wincadvisorperf
      run_interval 60s
      log_level debug
     </source>

     <filter mdm.kubepodinventory** mdm.kubenodeinventory**>
      type filter_inventory2mdm
      custom_metrics_azure_regions eastus,southcentralus,westcentralus,westus2,southeastasia,northeurope,westEurope
      log_level info
     </filter>

     # custom_metrics_mdm filter plugin for perf data from windows nodes
     <filter mdm.cadvisorperf**>
      type filter_cadvisor2mdm
      custom_metrics_azure_regions eastus,southcentralus,westcentralus,westus2,southeastasia,northeurope,westEurope
      metrics_to_collect cpuUsageNanoCores,memoryWorkingSetBytes
      log_level info
     </filter>

     <match oms.containerinsights.KubePodInventory**>
      type out_oms
      log_level debug
      num_threads 5
      buffer_chunk_limit 20m
      buffer_type file
      buffer_path %STATE_DIR_WS%/out_oms_kubepods*.buffer
      buffer_queue_limit 20
      buffer_queue_full_action drop_oldest_chunk
      flush_interval 20s
      retry_limit 10
      retry_wait 30s
      max_retry_wait 9m
     </match>

     <match oms.containerinsights.KubeEvents**>
      type out_oms
      log_level debug
      num_threads 5
      buffer_chunk_limit 5m
      buffer_type file
      buffer_path %STATE_DIR_WS%/out_oms_kubeevents*.buffer
      buffer_queue_limit 10
      buffer_queue_full_action drop_oldest_chunk
      flush_interval 20s
      retry_limit 10
      retry_wait 30s
      max_retry_wait 9m
     </match>

     <match oms.api.KubeLogs**>
      type out_oms_api
      log_level debug
      buffer_chunk_limit 10m
      buffer_type file
      buffer_path %STATE_DIR_WS%/out_oms_api_kubernetes_logs*.buffer
      buffer_queue_limit 10
      flush_interval 20s
      retry_limit 10
      retry_wait 30s
     </match>

     <match oms.containerinsights.KubeServices**>
      type out_oms
      log_level debug
      num_threads 5
      buffer_chunk_limit 20m
      buffer_type file
      buffer_path %STATE_DIR_WS%/out_oms_kubeservices*.buffer
      buffer_queue_limit 20
      buffer_queue_full_action drop_oldest_chunk
      flush_interval 20s
      retry_limit 10
      retry_wait 30s
      max_retry_wait 9m
     </match>

     <match oms.containerinsights.KubeNodeInventory**>
      type out_oms
      log_level debug
      num_threads 5
      buffer_chunk_limit 20m
      buffer_type file
      buffer_path %STATE_DIR_WS%/state/out_oms_kubenodes*.buffer
      buffer_queue_limit 20
      buffer_queue_full_action drop_oldest_chunk
      flush_interval 20s
      retry_limit 10
      retry_wait 30s
      max_retry_wait 9m
     </match>

     <match oms.containerinsights.ContainerNodeInventory**>
      type out_oms_api
      log_level debug
      buffer_chunk_limit 20m
      buffer_type file
      buffer_path %STATE_DIR_WS%/out_oms_containernodeinventory*.buffer
      buffer_queue_limit 20
      flush_interval 20s
      retry_limit 10
      retry_wait 15s
      max_retry_wait 9m
     </match>

     <match oms.api.KubePerf**>
      type out_oms
      log_level debug
      num_threads 5
      buffer_chunk_limit 20m
      buffer_type file
      buffer_path %STATE_DIR_WS%/out_oms_kubeperf*.buffer
      buffer_queue_limit 20
      buffer_queue_full_action drop_oldest_chunk
      flush_interval 20s
      retry_limit 10
      retry_wait 30s
      max_retry_wait 9m
     </match>

     <match mdm.kubepodinventory** mdm.kubenodeinventory** >
      type out_mdm
      log_level debug
      num_threads 5
      buffer_chunk_limit 20m
      buffer_type file
      buffer_path %STATE_DIR_WS%/out_mdm_*.buffer
      buffer_queue_limit 20
      buffer_queue_full_action drop_oldest_chunk
      flush_interval 20s
      retry_limit 10
      retry_wait 30s
      max_retry_wait 9m
      retry_mdm_post_wait_minutes 60
     </match>

     <match oms.api.wincadvisorperf**>
      type out_oms
      log_level debug
      num_threads 5
      buffer_chunk_limit 20m
      buffer_type file
      buffer_path %STATE_DIR_WS%/out_oms_api_wincadvisorperf*.buffer
      buffer_queue_limit 20
      buffer_queue_full_action drop_oldest_chunk
      flush_interval 20s
      retry_limit 10
      retry_wait 30s
      max_retry_wait 9m
     </match>

     <match mdm.cadvisorperf**>
      type out_mdm
      log_level debug
      num_threads 5
      buffer_chunk_limit 20m
      buffer_type file
      buffer_path %STATE_DIR_WS%/out_mdm_cdvisorperf*.buffer
      buffer_queue_limit 20
      buffer_queue_full_action drop_oldest_chunk
      flush_interval 20s
      retry_limit 10
      retry_wait 30s
      max_retry_wait 9m
      retry_mdm_post_wait_minutes 60
     </match>
metadata:
  name: omsagent-rs-config
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  labels:
    component: oms-agent
    tier: node
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: omsagent
  namespace: kube-system
spec:
  selector:
    matchLabels:
      component: oms-agent
      tier: node
  template:
    metadata:
      annotations:
        agentVersion: {{ContainerConfig "omsAgentVersion"}}
        dockerProviderVersion: {{ContainerConfig "dockerProviderVersion"}}
        schema-versions:  {{ContainerConfig "schema-versions"}}
      labels:
        component: oms-agent
        tier: node
    spec:
      priorityClassName: system-node-critical
      serviceAccountName: omsagent
      containers:
        - name: omsagent
          image: {{ContainerImage "omsagent"}}
          imagePullPolicy: Always
          resources:
            limits:
              cpu: {{ContainerCPULimits "omsagent"}}
              memory: {{ContainerMemLimits "omsagent"}}
            requests:
              cpu: {{ContainerCPUReqs "omsagent"}}
              memory: {{ContainerMemReqs "omsagent"}}
          env:
            - name: NODE_IP
              valueFrom:
                fieldRef:
                  fieldPath: status.hostIP
            - name: ACS_RESOURCE_NAME
              value: {{ContainerConfig "clusterName"}}            
            - name: CONTROLLER_TYPE
              value: "DaemonSet"
            - name: ISTEST
              value: "true"
          livenessProbe:
            exec:
              command:
                - /bin/bash
                - -c
                - /opt/livenessprobe.sh
            initialDelaySeconds: 60
            periodSeconds: 60
          ports:
            - containerPort: 25225
              protocol: TCP
            - containerPort: 25224
              protocol: UDP
          securityContext:
            privileged: true
          volumeMounts:
            - mountPath: /hostfs
              name: host-root
            - mountPath: /var/run/host
              name: docker-sock
            - mountPath: /var/log
              name: host-log
            - mountPath: /var/lib/docker/containers
              name: containerlog-path
            - mountPath: /etc/kubernetes/host
              name: azure-json-path
            - mountPath: /etc/omsagent-secret
              name: omsagent-secret
              readOnly: true
            - mountPath: /etc/config/settings
              name: settings-vol-config  
      nodeSelector:
        beta.kubernetes.io/os: linux
      tolerations:
        - effect: NoSchedule
          key: node-role.kubernetes.io/master
          operator: Equal
          value: "true"
      volumes:
        - name: host-root
          hostPath:
            path: /
        - name: docker-sock
          hostPath:
            path: /var/run
        - name: container-hostname
          hostPath:
            path: /etc/hostname
        - name: host-log
          hostPath:
            path: /var/log
        - name: containerlog-path
          hostPath:
            path: /var/lib/docker/containers
        - name: azure-json-path
          hostPath:
            path: /etc/kubernetes
        - name: omsagent-secret
          secret:
            secretName: omsagent-secret
        - name: settings-vol-config
          configMap:
            name: container-azm-ms-agentconfig
            optional: true    
  updateStrategy:
    type: RollingUpdate
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: omsagent-rs
  namespace: kube-system
  labels:
    component: oms-agent
    tier: node
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  replicas: 1
  selector:
    matchLabels:
      rsName: omsagent-rs
  strategy:
    type: RollingUpdate
  template:
    metadata:
      labels:
        rsName: omsagent-rs
      annotations:
        agentVersion: {{ContainerConfig "omsAgentVersion"}}
        dockerProviderVersion: {{ContainerConfig "dockerProviderVersion"}}
        schema-versions:  {{ContainerConfig "schema-versions"}}
    spec:
      serviceAccountName: omsagent
      containers:
        - name: omsagent
          image: {{ContainerImage "omsagent"}}
          imagePullPolicy: IfNotPresent
          resources:
            limits:
              cpu: {{ContainerCPULimits "omsagent"}}
              memory: {{ContainerMemLimits "omsagent"}}
            requests:
              cpu: {{ContainerCPUReqs "omsagent"}}
              memory: {{ContainerMemReqs "omsagent"}}
          env:
            - name: NODE_IP
              valueFrom:
                fieldRef:
                  fieldPath: status.hostIP
            - name: ACS_RESOURCE_NAME
              value: {{ContainerConfig "clusterName"}}            
            - name: CONTROLLER_TYPE
              value: "ReplicaSet"
            - name: ISTEST
              value: "true"
          securityContext:
            privileged: true
          ports:
            - containerPort: 25225
              protocol: TCP
            - containerPort: 25224
              protocol: UDP
          volumeMounts:
            - mountPath: /var/run/host
              name: docker-sock
            - mountPath: /var/log
              name: host-log
            - mountPath: /var/lib/docker/containers
              name: containerlog-path
            - mountPath: /etc/kubernetes/host
              name: azure-json-path
            - mountPath: /etc/omsagent-secret
              name: omsagent-secret
              readOnly: true
            - mountPath : /etc/config
              name: omsagent-rs-config
            - mountPath: /etc/config/settings
              name: settings-vol-config              
          livenessProbe:
            exec:
              command:
                - /bin/bash
                - -c
                - /opt/livenessprobe.sh
            initialDelaySeconds: 60
            periodSeconds: 60
      nodeSelector:
        beta.kubernetes.io/os: linux
        kubernetes.io/role: agent
      volumes:
        - name: docker-sock
          hostPath:
            path: /var/run
        - name: container-hostname
          hostPath:
            path: /etc/hostname
        - name: host-log
          hostPath:
            path: /var/log
        - name: containerlog-path
          hostPath:
            path: /var/lib/docker/containers
        - name: azure-json-path
          hostPath:
            path: /etc/kubernetes
        - name: omsagent-secret
          secret:
            secretName: omsagent-secret
        - name: omsagent-rs-config
          configMap:
            name: omsagent-rs-config
        - name: settings-vol-config
          configMap:
            name: container-azm-ms-agentconfig
            optional: true    
`)

func k8sContaineraddons116KubernetesmasteraddonsOmsagentDaemonsetYamlBytes() ([]byte, error) {
	return _k8sContaineraddons116KubernetesmasteraddonsOmsagentDaemonsetYaml, nil
}

func k8sContaineraddons116KubernetesmasteraddonsOmsagentDaemonsetYaml() (*asset, error) {
	bytes, err := k8sContaineraddons116KubernetesmasteraddonsOmsagentDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/1.16/kubernetesmasteraddons-omsagent-daemonset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddons116KubernetesmasteraddonsSmbFlexvolumeInstallerYaml = []byte(`apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: smb-flexvol-installer
  namespace: kube-system
  labels:
    k8s-app: smb
    kubernetes.io/cluster-service: "true"
spec:
  selector:
    matchLabels:
      name: smb
  template:
    metadata:
      labels:
        name: smb
        kubernetes.io/cluster-service: "true"
    spec:
      containers:
      - name: smb-flexvol-installer
        image: {{ContainerImage "smb-flexvolume"}}
        imagePullPolicy: Always
        resources:
          requests:
            cpu: {{ContainerCPUReqs "smb-flexvolume"}}
            memory: {{ContainerMemReqs "smb-flexvolume"}}
          limits:
            cpu: {{ContainerCPULimits "smb-flexvolume"}}
            memory: {{ContainerMemLimits "smb-flexvolume"}}
        volumeMounts:
        - name: volplugins
          mountPath: /etc/kubernetes/volumeplugins/
        - name: varlog
          mountPath: /var/log/
      volumes:
      - name: varlog
        hostPath:
          path: /var/log/
      - name: volplugins
        hostPath:
          path: /etc/kubernetes/volumeplugins/
          type: DirectoryOrCreate
      nodeSelector:
        beta.kubernetes.io/os: linux
`)

func k8sContaineraddons116KubernetesmasteraddonsSmbFlexvolumeInstallerYamlBytes() ([]byte, error) {
	return _k8sContaineraddons116KubernetesmasteraddonsSmbFlexvolumeInstallerYaml, nil
}

func k8sContaineraddons116KubernetesmasteraddonsSmbFlexvolumeInstallerYaml() (*asset, error) {
	bytes, err := k8sContaineraddons116KubernetesmasteraddonsSmbFlexvolumeInstallerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/1.16/kubernetesmasteraddons-smb-flexvolume-installer.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddons116KubernetesmasteraddonsTillerDeploymentYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: tiller
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: tiller
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: tiller
  namespace: kube-system
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: helm
    name: tiller
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: tiller-deploy
  namespace: kube-system
spec:
  ports:
  - name: tiller
    port: 44134
    targetPort: tiller
  selector:
    app: helm
    name: tiller
  type: ClusterIP
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: helm
    name: tiller
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: tiller-deploy
  namespace: kube-system
spec:
  selector:
    matchLabels:
      app: helm
      name: tiller
  template:
    metadata:
      labels:
        app: helm
        name: tiller
    spec:
      serviceAccountName: tiller
      containers:
      - env:
        - name: TILLER_NAMESPACE
          value: kube-system
        - name: TILLER_HISTORY_MAX
          value: "{{ContainerConfig "max-history"}}"
        image: {{ContainerImage "tiller"}}
        imagePullPolicy: IfNotPresent
        livenessProbe:
          httpGet:
            path: /liveness
            port: 44135
          initialDelaySeconds: 1
          timeoutSeconds: 1
        name: tiller
        ports:
        - containerPort: 44134
          name: tiller
        readinessProbe:
          httpGet:
            path: /readiness
            port: 44135
          initialDelaySeconds: 1
          timeoutSeconds: 1
        resources:
          requests:
            cpu: {{ContainerCPUReqs "tiller"}}
            memory: {{ContainerMemReqs "tiller"}}
          limits:
            cpu: {{ContainerCPULimits "tiller"}}
            memory: {{ContainerMemLimits "tiller"}}
      nodeSelector:
        beta.kubernetes.io/os: linux
`)

func k8sContaineraddons116KubernetesmasteraddonsTillerDeploymentYamlBytes() ([]byte, error) {
	return _k8sContaineraddons116KubernetesmasteraddonsTillerDeploymentYaml, nil
}

func k8sContaineraddons116KubernetesmasteraddonsTillerDeploymentYaml() (*asset, error) {
	bytes, err := k8sContaineraddons116KubernetesmasteraddonsTillerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/1.16/kubernetesmasteraddons-tiller-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddons16KubernetesmasteraddonsHeapsterDeploymentYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: heapster
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: system:heapster-with-nanny
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
rules:
- apiGroups:
  - extensions
  - apps
  resources:
  - deployments
  verbs:
  - get
  - list
  - watch
  - update
  - patch
- apiGroups:
  - ""
  resources:
  - events
  - namespaces
  - nodes
  - pods
  verbs:
  - get
  - list
  - watch
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: system:heapster-with-nanny
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:heapster-with-nanny
subjects:
- kind: ServiceAccount
  name: heapster
  namespace: kube-system
---
apiVersion: v1
kind: Service
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
    kubernetes.io/name: Heapster
  name: heapster
  namespace: kube-system
spec:
  ports:
  - port: 80
    targetPort: 8082
  selector:
    k8s-app: heapster
---
kind: Deployment
apiVersion: extensions/v1beta1
metadata:
  labels:
    k8s-app: heapster
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
  namespace: kube-system
  name: heapster
spec:
  replicas: 1
  selector:
    matchLabels:
      k8s-app: heapster
  template:
    metadata:
      labels:
        k8s-app: heapster
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ""
    spec:
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      serviceAccountName: heapster
      containers:
      - image: {{ContainerImage "heapster"}}
        imagePullPolicy: IfNotPresent
        command:
        - "/heapster"
        - "--source=kubernetes.summary_api:\"\""
        name: heapster
        resources:
          requests:
            cpu: 80m
            memory: 140Mi
          limits:
            cpu: 80m
            memory: 140Mi
      - image: {{ContainerImage "heapster-nanny"}}
        imagePullPolicy: IfNotPresent
        command:
        - "/pod_nanny"
        - "--cpu=80m"
        - "--extra-cpu=0.5m"
        - "--memory=140Mi"
        - "--extra-memory=4Mi"
        - "--threshold=5"
        - "--deployment=heapster"
        - "--container=heapster"
        - "--poll-period=300000"
        - "--estimator=exponential"
        name: heapster-nanny
        resources:
          requests:
            cpu: 50m
            memory: 90Mi
          limits:
            cpu: 50m
            memory: 90Mi
        env:
        - valueFrom:
            fieldRef:
              fieldPath: metadata.name
          name: MY_POD_NAME
        - valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
          name: MY_POD_NAMESPACE
      nodeSelector:
        beta.kubernetes.io/os: linux
`)

func k8sContaineraddons16KubernetesmasteraddonsHeapsterDeploymentYamlBytes() ([]byte, error) {
	return _k8sContaineraddons16KubernetesmasteraddonsHeapsterDeploymentYaml, nil
}

func k8sContaineraddons16KubernetesmasteraddonsHeapsterDeploymentYaml() (*asset, error) {
	bytes, err := k8sContaineraddons16KubernetesmasteraddonsHeapsterDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/1.6/kubernetesmasteraddons-heapster-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddons17KubernetesmasteraddonsHeapsterDeploymentYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: heapster
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: system:heapster-with-nanny
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
rules:
- apiGroups:
  - extensions
  - apps
  resources:
  - deployments
  verbs:
  - get
  - list
  - watch
  - update
  - patch
- apiGroups:
  - ""
  resources:
  - events
  - namespaces
  - nodes
  - pods
  verbs:
  - get
  - list
  - watch
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: system:heapster-with-nanny
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:heapster-with-nanny
subjects:
- kind: ServiceAccount
  name: heapster
  namespace: kube-system
---
apiVersion: v1
kind: Service
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
    kubernetes.io/name: Heapster
  name: heapster
  namespace: kube-system
spec:
  ports:
  - port: 80
    targetPort: 8082
  selector:
    k8s-app: heapster
---
kind: Deployment
apiVersion: extensions/v1beta1
metadata:
  labels:
    k8s-app: heapster
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
  namespace: kube-system
  name: heapster
spec:
  replicas: 1
  selector:
    matchLabels:
      k8s-app: heapster
  template:
    metadata:
      labels:
        k8s-app: heapster
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ""
    spec:
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      serviceAccountName: heapster
      containers:
      - image: {{ContainerImage "heapster"}}
        imagePullPolicy: IfNotPresent
        command:
        - "/heapster"
        - "--source=kubernetes.summary_api:\"\""
        name: heapster
        resources:
          requests:
            cpu: 80m
            memory: 140Mi
          limits:
            cpu: 80m
            memory: 140Mi
      - image: {{ContainerImage "heapster-nanny"}}
        imagePullPolicy: IfNotPresent
        command:
        - "/pod_nanny"
        - "--cpu=80m"
        - "--extra-cpu=0.5m"
        - "--memory=140Mi"
        - "--extra-memory=4Mi"
        - "--threshold=5"
        - "--deployment=heapster"
        - "--container=heapster"
        - "--poll-period=300000"
        - "--estimator=exponential"
        name: heapster-nanny
        resources:
          requests:
            cpu: 50m
            memory: 90Mi
          limits:
            cpu: 50m
            memory: 90Mi
        env:
        - valueFrom:
            fieldRef:
              fieldPath: metadata.name
          name: MY_POD_NAME
        - valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
          name: MY_POD_NAMESPACE
      nodeSelector:
        beta.kubernetes.io/os: linux
`)

func k8sContaineraddons17KubernetesmasteraddonsHeapsterDeploymentYamlBytes() ([]byte, error) {
	return _k8sContaineraddons17KubernetesmasteraddonsHeapsterDeploymentYaml, nil
}

func k8sContaineraddons17KubernetesmasteraddonsHeapsterDeploymentYaml() (*asset, error) {
	bytes, err := k8sContaineraddons17KubernetesmasteraddonsHeapsterDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/1.7/kubernetesmasteraddons-heapster-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddons18KubernetesmasteraddonsHeapsterDeploymentYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: heapster
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: system:heapster-with-nanny
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
rules:
- apiGroups:
  - extensions
  - apps
  resources:
  - deployments
  verbs:
  - get
  - list
  - watch
  - update
  - patch
- apiGroups:
  - ""
  resources:
  - events
  - namespaces
  - nodes
  - pods
  verbs:
  - get
  - list
  - watch
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: system:heapster-with-nanny
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:heapster-with-nanny
subjects:
- kind: ServiceAccount
  name: heapster
  namespace: kube-system
---
apiVersion: v1
kind: Service
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
    kubernetes.io/name: Heapster
  name: heapster
  namespace: kube-system
spec:
  ports:
  - port: 80
    targetPort: 8082
  selector:
    k8s-app: heapster
---
kind: Deployment
apiVersion: extensions/v1beta1
metadata:
  labels:
    k8s-app: heapster
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
  namespace: kube-system
  name: heapster
spec:
  replicas: 1
  selector:
    matchLabels:
      k8s-app: heapster
  template:
    metadata:
      labels:
        k8s-app: heapster
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ""
    spec:
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      serviceAccountName: heapster
      containers:
      - image: {{ContainerImage "heapster"}}
        imagePullPolicy: IfNotPresent
        command:
        - "/heapster"
        - "--source=kubernetes.summary_api:\"\""
        name: heapster
        resources:
          requests:
            cpu: 80m
            memory: 140Mi
          limits:
            cpu: 80m
            memory: 140Mi
      - image: {{ContainerImage "heapster-nanny"}}
        imagePullPolicy: IfNotPresent
        command:
        - "/pod_nanny"
        - "--cpu=80m"
        - "--extra-cpu=0.5m"
        - "--memory=140Mi"
        - "--extra-memory=4Mi"
        - "--threshold=5"
        - "--deployment=heapster"
        - "--container=heapster"
        - "--poll-period=300000"
        - "--estimator=exponential"
        name: heapster-nanny
        resources:
          requests:
            cpu: 50m
            memory: 90Mi
          limits:
            cpu: 50m
            memory: 90Mi
        env:
        - valueFrom:
            fieldRef:
              fieldPath: metadata.name
          name: MY_POD_NAME
        - valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
          name: MY_POD_NAMESPACE
      nodeSelector:
        beta.kubernetes.io/os: linux
`)

func k8sContaineraddons18KubernetesmasteraddonsHeapsterDeploymentYamlBytes() ([]byte, error) {
	return _k8sContaineraddons18KubernetesmasteraddonsHeapsterDeploymentYaml, nil
}

func k8sContaineraddons18KubernetesmasteraddonsHeapsterDeploymentYaml() (*asset, error) {
	bytes, err := k8sContaineraddons18KubernetesmasteraddonsHeapsterDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/1.8/kubernetesmasteraddons-heapster-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddonsAzureCniNetworkmonitorYaml = []byte(`apiVersion: extensions/v1beta1
kind: DaemonSet
metadata:
  name: azure-cni-networkmonitor
  namespace: kube-system
  labels:
    app: azure-cnms
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  selector:
    matchLabels:
      k8s-app: azure-cnms
  template:
    metadata:
      labels:
        k8s-app: azure-cnms
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ''
    spec:
      priorityClassName: system-node-critical
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      - key: node-role.kubernetes.io/master
        operator: Equal
        value: "true"
        effect: NoSchedule
      - operator: "Exists"
        effect: NoExecute
      - operator: "Exists"
        effect: NoSchedule
      nodeSelector:
        beta.kubernetes.io/os: linux
      containers:
        - name: azure-cnms
          image: {{ContainerImage "azure-cni-networkmonitor"}}
          imagePullPolicy: IfNotPresent
          securityContext:
            privileged: true
          env:
            - name: HOSTNAME
              valueFrom:
                fieldRef:
                  apiVersion: v1
                  fieldPath: spec.nodeName
          volumeMounts:
          - name: ebtables-rule-repo
            mountPath: /var/run
          - name: log
            mountPath: /var/log
          - name: telemetry
            mountPath: /opt/cni/bin
      hostNetwork: true
      volumes:
      - name: log
        hostPath:
          path: /var/log
          type: Directory
      - name: ebtables-rule-repo
        hostPath:
          path: /var/run/
          type: Directory
      - name: telemetry
        hostPath:
          path: /opt/cni/bin
          type: Directory`)

func k8sContaineraddonsAzureCniNetworkmonitorYamlBytes() ([]byte, error) {
	return _k8sContaineraddonsAzureCniNetworkmonitorYaml, nil
}

func k8sContaineraddonsAzureCniNetworkmonitorYaml() (*asset, error) {
	bytes, err := k8sContaineraddonsAzureCniNetworkmonitorYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/azure-cni-networkmonitor.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddonsDnsAutoscalerYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  name: dns-autoscaler
  namespace: kube-system
  labels:
    k8s-app: dns-autoscaler
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  selector:
    matchLabels:
       k8s-app: dns-autoscaler
  template:
    metadata:
      labels:
        k8s-app: dns-autoscaler
    spec:
      containers:
      - name: autoscaler
        image: {{ContainerImage "dns-autoscaler"}}
        resources:
          requests:
            cpu: {{ContainerCPUReqs "dns-autoscaler"}}
            memory: {{ContainerMemReqs "dns-autoscaler"}}
        command:
          - /cluster-proportional-autoscaler
          - --namespace=kube-system
          - --configmap=dns-autoscaler
          - --target=Deployment/coredns
          # When cluster is using large nodes(with more cores), "coresPerReplica" should dominate.
          # If using small nodes, "nodesPerReplica" should dominate.
          - --default-params={"linear":{"coresPerReplica":256,"nodesPerReplica":16,"min":1}}
          - --logtostderr=true
          - --v=2
`)

func k8sContaineraddonsDnsAutoscalerYamlBytes() ([]byte, error) {
	return _k8sContaineraddonsDnsAutoscalerYaml, nil
}

func k8sContaineraddonsDnsAutoscalerYaml() (*asset, error) {
	bytes, err := k8sContaineraddonsDnsAutoscalerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/dns-autoscaler.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddonsIpMasqAgentYaml = []byte(`apiVersion: extensions/v1beta1
kind: DaemonSet
metadata:
  name: azure-ip-masq-agent
  namespace: kube-system
  labels:
    component: azure-ip-masq-agent
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
    tier: node
spec:
  template:
    metadata:
      labels:
        k8s-app: azure-ip-masq-agent
        tier: node
    spec:
      priorityClassName: system-node-critical
      hostNetwork: true
      nodeSelector:
        beta.kubernetes.io/os: linux
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      - key: node-role.kubernetes.io/master
        operator: Equal
        value: "true"
        effect: NoSchedule
      - operator: "Exists"
        effect: NoExecute
      - operator: "Exists"
        effect: NoSchedule
      containers:
      - name: azure-ip-masq-agent
        image: {{ContainerImage "ip-masq-agent"}}
        imagePullPolicy: IfNotPresent
        securityContext:
          privileged: true
        volumeMounts:
          - name: azure-ip-masq-agent-config-volume
            mountPath: /etc/config
        resources:
          requests:
            cpu: {{ContainerCPUReqs "ip-masq-agent"}}
            memory: {{ContainerMemReqs "ip-masq-agent"}}
          limits:
            cpu: {{ContainerCPULimits "ip-masq-agent"}}
            memory: {{ContainerMemLimits "ip-masq-agent"}}
      volumes:
        - name: azure-ip-masq-agent-config-volume
          configMap:
            name: azure-ip-masq-agent-config
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: azure-ip-masq-agent-config
  namespace: kube-system
  labels:
    component: azure-ip-masq-agent
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
data:
  ip-masq-agent: |-
    nonMasqueradeCIDRs:
      - {{ContainerConfig "non-masquerade-cidr"}}
    {{- if ContainerConfig "non-masq-cni-cidr"}}
      - {{ContainerConfig "non-masq-cni-cidr"}}
    masqLinkLocal: true
    {{else}}
    masqLinkLocal: false
    {{end -}}
    resyncInterval: 60s`)

func k8sContaineraddonsIpMasqAgentYamlBytes() ([]byte, error) {
	return _k8sContaineraddonsIpMasqAgentYaml, nil
}

func k8sContaineraddonsIpMasqAgentYaml() (*asset, error) {
	bytes, err := k8sContaineraddonsIpMasqAgentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/ip-masq-agent.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddonsKubernetesmasteraddonsAadPodIdentityDeploymentYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: aad-pod-id-nmi-service-account
  namespace: default
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: azureassignedidentities.aadpodidentity.k8s.io
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  group: aadpodidentity.k8s.io
  version: v1
  names:
    kind: AzureAssignedIdentity
    plural: azureassignedidentities
  scope: Namespaced
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: azureidentitybindings.aadpodidentity.k8s.io
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  group: aadpodidentity.k8s.io
  version: v1
  names:
    kind: AzureIdentityBinding
    plural: azureidentitybindings
  scope: Namespaced
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: azureidentities.aadpodidentity.k8s.io
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  group: aadpodidentity.k8s.io
  version: v1
  names:
    kind: AzureIdentity
    singular: azureidentity
    plural: azureidentities
  scope: Namespaced
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: aad-pod-id-nmi-role
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups: ["*"]
  resources: ["*"]
  verbs: ["get", "list"]
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: aad-pod-id-nmi-binding
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
    k8s-app: aad-pod-id-nmi-binding
subjects:
- kind: ServiceAccount
  name: aad-pod-id-nmi-service-account
  namespace: default
roleRef:
  kind: ClusterRole
  name: aad-pod-id-nmi-role
  apiGroup: rbac.authorization.k8s.io
---
apiVersion: extensions/v1beta1
kind: DaemonSet
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
    component: nmi
    tier: node
    k8s-app: aad-pod-id
  name: nmi
  namespace: default
spec:
  template:
    metadata:
      labels:
        component: nmi
        tier: node
    spec:
      serviceAccountName: aad-pod-id-nmi-service-account
      hostNetwork: true
      containers:
      - name: nmi
        image: {{ContainerImage "nmi"}}
        imagePullPolicy: IfNotPresent
        resources:
          requests:
            cpu: {{ContainerCPUReqs "nmi"}}
            memory: {{ContainerMemReqs "nmi"}}
          limits:
            cpu: {{ContainerCPULimits "nmi"}}
            memory: {{ContainerMemLimits "nmi"}}
        args:
          - "--host-ip=$(HOST_IP)"
          - "--node=$(NODE_NAME)"
        env:
          - name: HOST_IP
            valueFrom:
              fieldRef:
                fieldPath: status.podIP
          - name: NODE_NAME
            valueFrom:
              fieldRef:
                fieldPath: spec.nodeName
        securityContext:
          privileged: true
          capabilities:
            add:
            - NET_ADMIN
      nodeSelector:
        beta.kubernetes.io/os: linux
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: aad-pod-id-mic-service-account
  namespace: default
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: aad-pod-id-mic-role
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups: ["apiextensions.k8s.io"]
  resources: ["customresourcedefinitions"]
  verbs: ["*"]
- apiGroups: [""]
  resources: ["pods"]
  verbs: [ "list", "watch" ]
- apiGroups: [""]
  resources: ["events"]
  verbs: ["create", "patch"]
- apiGroups: ["aadpodidentity.k8s.io"]
  resources: ["azureidentitybindings", "azureidentities"]
  verbs: ["get", "list", "watch", "post"]
- apiGroups: ["aadpodidentity.k8s.io"]
  resources: ["azureassignedidentities"]
  verbs: ["*"]
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: aad-pod-id-mic-binding
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
    k8s-app: aad-pod-id-mic-binding
subjects:
- kind: ServiceAccount
  name: aad-pod-id-mic-service-account
  namespace: default
roleRef:
  kind: ClusterRole
  name: aad-pod-id-mic-role
  apiGroup: rbac.authorization.k8s.io
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    component: mic
    k8s-app: aad-pod-id
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: mic
  namespace: default
spec:
  template:
    metadata:
      labels:
        component: mic
    spec:
      serviceAccountName: aad-pod-id-mic-service-account
      containers:
      - name: mic
        image: {{ContainerImage "mic"}}
        imagePullPolicy: IfNotPresent
        resources:
          requests:
            cpu: {{ContainerCPUReqs "mic"}}
            memory: {{ContainerMemReqs "mic"}}
          limits:
            cpu: {{ContainerCPULimits "mic"}}
            memory: {{ContainerMemLimits "mic"}}
        args:
          - --cloudconfig=/etc/kubernetes/azure.json
          - --logtostderr
        volumeMounts:
          - name: k8s-azure-file
            mountPath: /etc/kubernetes/azure.json
            readOnly: true
      volumes:
      - name: k8s-azure-file
        hostPath:
          path: /etc/kubernetes/azure.json
`)

func k8sContaineraddonsKubernetesmasteraddonsAadPodIdentityDeploymentYamlBytes() ([]byte, error) {
	return _k8sContaineraddonsKubernetesmasteraddonsAadPodIdentityDeploymentYaml, nil
}

func k8sContaineraddonsKubernetesmasteraddonsAadPodIdentityDeploymentYaml() (*asset, error) {
	bytes, err := k8sContaineraddonsKubernetesmasteraddonsAadPodIdentityDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/kubernetesmasteraddons-aad-pod-identity-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddonsKubernetesmasteraddonsAciConnectorDeploymentYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: aci-connector
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: aci-connector
  labels:
    app: aci-connector
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups:
  - ""
  resources:
  - configmaps
  - pods
  - services
  - endpoints
  - events
  - secrets
  - nodes
  - nodes/status
  - pods/status
  verbs:
  - "*"
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: aci-connector
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: aci-connector
subjects:
- kind: ServiceAccount
  name: aci-connector
  namespace: kube-system
---
apiVersion: v1
kind: Secret
metadata:
  name: aci-connector-secret
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
type: Opaque
data:
  credentials.json: <creds>
  cert.pem: <cert>
  key.pem: <key>
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: aci-connector
  namespace: kube-system
  labels:
    app: aci-connector
    name: aci-connector
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: aci-connector
    spec:
      serviceAccountName: aci-connector
      nodeSelector:
        beta.kubernetes.io/os: linux
      containers:
      - name: aci-connector
        image: {{ContainerImage "aci-connector"}}
        imagePullPolicy: Always
        env:
        - name: KUBELET_PORT
          value: "10250"
        - name: AZURE_AUTH_LOCATION
          value: /etc/virtual-kubelet/credentials.json
        - name: ACI_RESOURCE_GROUP
          value: <rgName>
        - name: ACI_REGION
          value: <region>
        - name: APISERVER_CERT_LOCATION
          value: /etc/virtual-kubelet/cert.pem
        - name: APISERVER_KEY_LOCATION
          value: /etc/virtual-kubelet/key.pem
        - name: VKUBELET_POD_IP
          valueFrom:
            fieldRef:
              fieldPath: status.podIP
        resources:
          requests:
            cpu: {{ContainerCPUReqs "aci-connector"}}
            memory: {{ContainerMemReqs "aci-connector"}}
          limits:
            cpu: {{ContainerCPULimits "aci-connector"}}
            memory: {{ContainerMemLimits "aci-connector"}}
        volumeMounts:
        - name: credentials
          mountPath: "/etc/virtual-kubelet"
          readOnly: true
        command: ["virtual-kubelet"]
        args: ["--provider", "azure", "--nodename", "{{ContainerConfig "nodeName"}}" , "--os", "{{ContainerConfig "os"}}", "--taint", "{{ContainerConfig "taint"}}"]
      volumes:
      - name: credentials
        secret:
          secretName: aci-connector-secret
#EOF
`)

func k8sContaineraddonsKubernetesmasteraddonsAciConnectorDeploymentYamlBytes() ([]byte, error) {
	return _k8sContaineraddonsKubernetesmasteraddonsAciConnectorDeploymentYaml, nil
}

func k8sContaineraddonsKubernetesmasteraddonsAciConnectorDeploymentYaml() (*asset, error) {
	bytes, err := k8sContaineraddonsKubernetesmasteraddonsAciConnectorDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/kubernetesmasteraddons-aci-connector-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddonsKubernetesmasteraddonsAzureNpmDaemonsetYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: azure-npm
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: EnsureExists
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: azure-npm
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: EnsureExists
rules:
  - apiGroups:
    - ""
    resources:
      - pods
      - nodes
      - namespaces
    verbs:
      - get
      - list
      - watch
  - apiGroups:
    - networking.k8s.io
    resources:
      - networkpolicies
    verbs:
      - get
      - list
      - watch
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: azure-npm-binding
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: EnsureExists
subjects:
  - kind: ServiceAccount
    name: azure-npm
    namespace: kube-system
roleRef:
  kind: ClusterRole
  name: azure-npm
  apiGroup: rbac.authorization.k8s.io
---
apiVersion: extensions/v1beta1
kind: DaemonSet
metadata:
  name: azure-npm
  namespace: kube-system
  labels:
    app: azure-npm
    addonmanager.kubernetes.io/mode: EnsureExists
spec:
  selector:
    matchLabels:
      k8s-app: azure-npm
  template:
    metadata:
      labels:
        k8s-app: azure-npm
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ''
    spec:
      priorityClassName: system-node-critical
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      nodeSelector:
        beta.kubernetes.io/os: linux
      containers:
        - name: azure-npm
          image: {{ContainerImage "azure-npm-daemonset"}}
          securityContext:
            privileged: true
          env:
            - name: HOSTNAME
              valueFrom:
                fieldRef:
                  apiVersion: v1
                  fieldPath: spec.nodeName
          volumeMounts:
          - name: xtables-lock
            mountPath: /run/xtables.lock
          - name: log
            mountPath: /var/log
          - name: socket-dir
            mountPath: /var/run
          - name: tmp
            mountPath: /tmp
        - name: azure-vnet-telemetry
          image: {{ContainerImage "azure-vnet-telemetry-daemonset"}}
          volumeMounts:
          - name: socket-dir
            mountPath: /var/run
          - name: tmp
            mountPath: /tmp
      hostNetwork: true
      volumes:
      - name: log
        hostPath:
          path: /var/log
          type: Directory
      - name: xtables-lock
        hostPath:
          path: /run/xtables.lock
          type: File
      - name: tmp
        hostPath:
          path: /tmp
          type: Directory
      - name: socket-dir
        emptyDir: {}
      serviceAccountName: azure-npm
`)

func k8sContaineraddonsKubernetesmasteraddonsAzureNpmDaemonsetYamlBytes() ([]byte, error) {
	return _k8sContaineraddonsKubernetesmasteraddonsAzureNpmDaemonsetYaml, nil
}

func k8sContaineraddonsKubernetesmasteraddonsAzureNpmDaemonsetYaml() (*asset, error) {
	bytes, err := k8sContaineraddonsKubernetesmasteraddonsAzureNpmDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/kubernetesmasteraddons-azure-npm-daemonset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddonsKubernetesmasteraddonsBlobfuseFlexvolumeInstallerYaml = []byte(`apiVersion: extensions/v1beta1
kind: DaemonSet
metadata:
  name: blobfuse-flexvol-installer
  namespace: kube-system
  labels:
    k8s-app: blobfuse
    kubernetes.io/cluster-service: "true"
spec:
  selector:
    matchLabels:
      name: blobfuse
  template:
    metadata:
      labels:
        name: blobfuse
        kubernetes.io/cluster-service: "true"
    spec:
      containers:
      - name: blobfuse-flexvol-installer
        image: {{ContainerImage "blobfuse-flexvolume"}}
        imagePullPolicy: IfNotPresent
        resources:
          requests:
            cpu: {{ContainerCPUReqs "blobfuse-flexvolume"}}
            memory: {{ContainerMemReqs "blobfuse-flexvolume"}}
          limits:
            cpu: {{ContainerCPULimits "blobfuse-flexvolume"}}
            memory: {{ContainerMemLimits "blobfuse-flexvolume"}}
        volumeMounts:
        - name: volplugins
          mountPath: /etc/kubernetes/volumeplugins/
        - name: varlog
          mountPath: /var/log/      
      volumes:
      - name: varlog
        hostPath:
          path: /var/log/              
      - name: volplugins
        hostPath:
          path: /etc/kubernetes/volumeplugins/
      nodeSelector:
        beta.kubernetes.io/os: linux
`)

func k8sContaineraddonsKubernetesmasteraddonsBlobfuseFlexvolumeInstallerYamlBytes() ([]byte, error) {
	return _k8sContaineraddonsKubernetesmasteraddonsBlobfuseFlexvolumeInstallerYaml, nil
}

func k8sContaineraddonsKubernetesmasteraddonsBlobfuseFlexvolumeInstallerYaml() (*asset, error) {
	bytes, err := k8sContaineraddonsKubernetesmasteraddonsBlobfuseFlexvolumeInstallerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/kubernetesmasteraddons-blobfuse-flexvolume-installer.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddonsKubernetesmasteraddonsCalicoDaemonsetYaml = []byte(`# Source: calico/templates/calico-config.yaml
# This ConfigMap is used to configure a self-hosted Calico installation.
kind: ConfigMap
apiVersion: v1
metadata:
  name: calico-config
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "EnsureExists"
data:
  # You must set a non-zero value for Typha replicas below.
  typha_service_name: "calico-typha"

  # The CNI network configuration to install on each node.  The special
  # values in this config will be automatically populated.
  cni_network_config: |-
    {
      "name": "k8s-pod-network",
      "cniVersion": "0.3.1",
      "plugins": [
        {
          "type": "calico",
          "log_level": "info",
          "datastore_type": "kubernetes",
          "nodename": "__KUBERNETES_NODE_NAME__",
          "mtu": 1500,
          "ipam": <calicoIPAMConfig>,
          "policy": {
              "type": "k8s"
          },
          "kubernetes": {
              "kubeconfig": "__KUBECONFIG_FILEPATH__"
          }
        },
        {
          "type": "portmap",
          "snat": true,
          "capabilities": {"portMappings": true}
        }
      ]
    }

---
# Source: calico/templates/kdd-crds.yaml
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: felixconfigurations.crd.projectcalico.org
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: FelixConfiguration
    plural: felixconfigurations
    singular: felixconfiguration
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: bgpconfigurations.crd.projectcalico.org
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: BGPConfiguration
    plural: bgpconfigurations
    singular: bgpconfiguration
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: ippools.crd.projectcalico.org
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: IPPool
    plural: ippools
    singular: ippool
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: hostendpoints.crd.projectcalico.org
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: HostEndpoint
    plural: hostendpoints
    singular: hostendpoint
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: clusterinformations.crd.projectcalico.org
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: ClusterInformation
    plural: clusterinformations
    singular: clusterinformation
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: globalnetworkpolicies.crd.projectcalico.org
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: GlobalNetworkPolicy
    plural: globalnetworkpolicies
    singular: globalnetworkpolicy
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: globalnetworksets.crd.projectcalico.org
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  scope: Cluster
  group: crd.projectcalico.org
  version: v1
  names:
    kind: GlobalNetworkSet
    plural: globalnetworksets
    singular: globalnetworkset
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: networkpolicies.crd.projectcalico.org
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  scope: Namespaced
  group: crd.projectcalico.org
  version: v1
  names:
    kind: NetworkPolicy
    plural: networkpolicies
    singular: networkpolicy
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: networksets.crd.projectcalico.org
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  scope: Namespaced
  group: crd.projectcalico.org
  version: v1
  names:
    kind: NetworkSet
    plural: networksets
    singular: networkset
---
# Source: calico/templates/rbac.yaml

# Include a clusterrole for the calico-node DaemonSet,
# and bind it to the calico-node serviceaccount.
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: calico-node
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
rules:
# The CNI plugin needs to get pods, nodes, and namespaces.
- apiGroups: [""]
  resources:
  - pods
  - nodes
  - namespaces
  verbs:
  - get
- apiGroups: [""]
  resources:
  - endpoints
  - services
  verbs:
  # Used to discover service IPs for advertisement.
  - watch
  - list
  # Used to discover Typhas.
  - get
- apiGroups: [""]
  resources:
  - nodes/status
  verbs:
  # Needed for clearing NodeNetworkUnavailable flag.
  - patch
  # Calico stores some configuration information in node annotations.
  - update
# Watch for changes to Kubernetes NetworkPolicies.
- apiGroups: ["networking.k8s.io"]
  resources:
  - networkpolicies
  verbs:
  - watch
  - list
# Used by Calico for policy information.
- apiGroups: [""]
  resources:
  - pods
  - namespaces
  - serviceaccounts
  verbs:
  - list
  - watch
# The CNI plugin patches pods/status.
- apiGroups: [""]
  resources:
  - pods/status
  verbs:
  - patch
# Calico monitors various CRDs for config.
- apiGroups: ["crd.projectcalico.org"]
  resources:
  - globalfelixconfigs
  - felixconfigurations
  - bgppeers
  - globalbgpconfigs
  - bgpconfigurations
  - ippools
  - ipamblocks
  - globalnetworkpolicies
  - globalnetworksets
  - networkpolicies
  - networksets
  - clusterinformations
  - hostendpoints
  verbs:
  - get
  - list
  - watch
# Calico must create and update some CRDs on startup.
- apiGroups: ["crd.projectcalico.org"]
  resources:
  - ippools
  - felixconfigurations
  - clusterinformations
  verbs:
  - create
  - update
# Calico stores some configuration information on the node.
- apiGroups: [""]
  resources:
  - nodes
  verbs:
  - get
  - list
  - watch
# These permissions are only requried for upgrade from v2.6, and can
# be removed after upgrade or on fresh installations.
- apiGroups: ["crd.projectcalico.org"]
  resources:
  - bgpconfigurations
  - bgppeers
  verbs:
  - create
  - update
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: calico-node
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: calico-node
subjects:
- kind: ServiceAccount
  name: calico-node
  namespace: kube-system
---
# Source: calico/templates/calico-typha.yaml
# This manifest creates a Service, which will be backed by Calico's Typha daemon.
# Typha sits in between Felix and the API server, reducing Calico's load on the API server.

apiVersion: v1
kind: Service
metadata:
  name: calico-typha
  namespace: kube-system
  labels:
    k8s-app: calico-typha
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  ports:
  - port: 5473
    protocol: TCP
    targetPort: calico-typha
    name: calico-typha
  selector:
    k8s-app: calico-typha
---

# This manifest creates a Deployment of Typha to back the above service.

apiVersion: apps/v1
kind: Deployment
metadata:
  name: calico-typha
  namespace: kube-system
  labels:
    k8s-app: calico-typha
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  # Number of Typha replicas.  To enable Typha, set this to a non-zero value *and* set the
  # typha_service_name variable in the calico-config ConfigMap above.
  #
  # We recommend using Typha if you have more than 50 nodes.  Above 100 nodes it is essential
  # (when using the Kubernetes datastore).  Use one replica for every 100-200 nodes.  In
  # production, we recommend running at least 3 replicas to reduce the impact of rolling upgrade.
  replicas: 1
  revisionHistoryLimit: 2
  selector:
    matchLabels:
      k8s-app: calico-typha
  template:
    metadata:
      labels:
        k8s-app: calico-typha
      annotations:
        # This, along with the CriticalAddonsOnly toleration below, marks the pod as a critical
        # add-on, ensuring it gets priority scheduling and that its resources are reserved
        # if it ever gets evicted.
        scheduler.alpha.kubernetes.io/critical-pod: ''
        cluster-autoscaler.kubernetes.io/safe-to-evict: 'true'
    spec:
      nodeSelector:
        beta.kubernetes.io/os: linux
      hostNetwork: true
      tolerations:
      # Mark the pod as a critical add-on for rescheduling.
      - key: CriticalAddonsOnly
        operator: Exists
      # Since Calico can't network a pod until Typha is up, we need to run Typha itself
      # as a host-networked pod.
      serviceAccountName: calico-node
      priorityClassName: system-cluster-critical
      containers:
      - image: {{ContainerImage "calico-typha"}}
        name: calico-typha
        ports:
        - containerPort: 5473
          name: calico-typha
          protocol: TCP
        env:
        # Enable "info" logging by default.  Can be set to "debug" to increase verbosity.
        - name: TYPHA_LOGSEVERITYSCREEN
          value: "info"
        # Disable logging to file and syslog since those don't make sense in Kubernetes.
        - name: TYPHA_LOGFILEPATH
          value: "none"
        - name: TYPHA_LOGSEVERITYSYS
          value: "none"
        # Monitor the Kubernetes API to find the number of running instances and rebalance
        # connections.
        - name: TYPHA_CONNECTIONREBALANCINGMODE
          value: "kubernetes"
        - name: TYPHA_DATASTORETYPE
          value: "kubernetes"
        - name: TYPHA_HEALTHENABLED
          value: "true"
        # Configure route aggregation based on pod CIDR.
        - name: USE_POD_CIDR
          value: "true"
        - name: FELIX_INTERFACEPREFIX
          value: "azv"
        # Uncomment these lines to enable prometheus metrics.  Since Typha is host-networked,
        # this opens a port on the host, which may need to be secured.
        #- name: TYPHA_PROMETHEUSMETRICSENABLED
        #  value: "true"
        #- name: TYPHA_PROMETHEUSMETRICSPORT
        #  value: "9093"
        livenessProbe:
          httpGet:
            path: /liveness
            port: 9098
            host: localhost
          periodSeconds: 30
          initialDelaySeconds: 30
        readinessProbe:
          httpGet:
            path: /readiness
            port: 9098
            host: localhost
          periodSeconds: 10
---
# Source: calico/templates/calico-node.yaml
# This manifest installs the calico-node container, as well
# as the CNI plugins and network config on
# each master and worker node in a Kubernetes cluster.
kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: calico-node
  namespace: kube-system
  labels:
    k8s-app: calico-node
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  selector:
    matchLabels:
      k8s-app: calico-node
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 1
  template:
    metadata:
      labels:
        k8s-app: calico-node
      annotations:
        # This, along with the CriticalAddonsOnly toleration below,
        # marks the pod as a critical add-on, ensuring it gets
        # priority scheduling and that its resources are reserved
        # if it ever gets evicted.
        scheduler.alpha.kubernetes.io/critical-pod: ''
    spec:
      nodeSelector:
        beta.kubernetes.io/os: linux
      hostNetwork: true
      tolerations:
      # Make sure calico-node gets scheduled on all nodes.
      - effect: NoSchedule
        operator: Exists
      # Mark the pod as a critical add-on for rescheduling.
      - key: CriticalAddonsOnly
        operator: Exists
      - effect: NoExecute
        operator: Exists
      serviceAccountName: calico-node
      # Minimize downtime during a rolling upgrade or deletion; tell Kubernetes to do a "force
      # deletion": https://kubernetes.io/docs/concepts/workloads/pods/pod/#termination-of-pods.
      terminationGracePeriodSeconds: 0
      priorityClassName: system-node-critical
      initContainers:
      # Start of install-cni initContainer
      # This container installs the CNI binaries
      # and CNI network config file on each node.
      - name: install-cni
        image: {{ContainerImage "calico-cni"}}
        command: ["/install-cni.sh"]
        env:
        # Name of the CNI config file to create.
        - name: CNI_CONF_NAME
          value: "10-calico.conflist"
        # The CNI network config to install on each node.
        - name: CNI_NETWORK_CONFIG
          valueFrom:
            configMapKeyRef:
              name: calico-config
              key: cni_network_config
        # Set the hostname based on the k8s node name.
        - name: KUBERNETES_NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        # Prevents the container from sleeping forever.
        - name: SLEEP
          value: "false"
        volumeMounts:
        - mountPath: /host/opt/cni/bin
          name: cni-bin-dir
        - mountPath: /host/etc/cni/net.d
          name: cni-net-dir
      # End of install-cni initContainer
      # Adds a Flex Volume Driver that creates a per-pod Unix Domain Socket to allow Dikastes
      # to communicate with Felix over the Policy Sync API.
      - name: flexvol-driver
        image: {{ContainerImage "calico-pod2daemon"}}
        volumeMounts:
        - name: flexvol-driver-host
          mountPath: /host/driver
      containers:
      # Runs calico-node container on each Kubernetes node.  This
      # container programs network policy and routes on each
      # host.
      - name: calico-node
        image: {{ContainerImage "calico-node"}}
        env:
        # Use Kubernetes API as the backing datastore.
        - name: DATASTORE_TYPE
          value: "kubernetes"
        # Configure route aggregation based on pod CIDR.
        - name: USE_POD_CIDR
          value: "true"
        # Typha support: controlled by the ConfigMap.
        - name: FELIX_TYPHAK8SSERVICENAME
          valueFrom:
            configMapKeyRef:
              name: calico-config
              key: typha_service_name
        # Wait for the datastore.
        - name: WAIT_FOR_DATASTORE
          value: "true"
        # Set based on the k8s node name.
        - name: NODENAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        # Don't enable BGP.
        - name: CALICO_NETWORKING_BACKEND
          value: "none"
        # Cluster type to identify the deployment type
        - name: CLUSTER_TYPE
          value: "k8s"
        # The default IPv4 pool to create on startup if none exists. Pod IPs will be
        # chosen from this range. Changing this value after installation will have
        # no effect. This should fall within ` + "`" + `--cluster-cidr` + "`" + `.
        - name: CALICO_IPV4POOL_CIDR
          value: "<kubeClusterCidr>"
        # Disable file logging so ` + "`" + `kubectl logs` + "`" + ` works.
        - name: CALICO_DISABLE_FILE_LOGGING
          value: "true"
        # Set Felix endpoint to host default action to ACCEPT.
        - name: FELIX_DEFAULTENDPOINTTOHOSTACTION
          value: "ACCEPT"
        # Disable IPv6 on Kubernetes.
        - name: FELIX_IPV6SUPPORT
          value: "false"
        # Set Felix logging to "info"
        - name: FELIX_LOGSEVERITYSCREEN
          value: "info"
        - name: FELIX_HEALTHENABLED
          value: "true"
        - name: CALICO_IPV4POOL_IPIP
          value: "off"
        - name: FELIX_INTERFACEPREFIX
          value: "azv"
        securityContext:
          privileged: true
        resources:
          requests:
            cpu: 250m
        livenessProbe:
          httpGet:
            path: /liveness
            port: 9099
            host: localhost
          periodSeconds: 10
          initialDelaySeconds: 10
          failureThreshold: 6
        readinessProbe:
          exec:
            command:
            - /bin/calico-node
            - -felix-ready
          periodSeconds: 10
        volumeMounts:
        - mountPath: /lib/modules
          name: lib-modules
          readOnly: true
        - mountPath: /run/xtables.lock
          name: xtables-lock
          readOnly: false
        - mountPath: /var/run/calico
          name: var-run-calico
          readOnly: false
        - mountPath: /var/lib/calico
          name: var-lib-calico
          readOnly: false
        - name: policysync
          mountPath: /var/run/nodeagent
      volumes:
      # Used by calico-node.
      - name: lib-modules
        hostPath:
          path: /lib/modules
      - name: var-run-calico
        hostPath:
          path: /var/run/calico
      - name: var-lib-calico
        hostPath:
          path: /var/lib/calico
      - name: xtables-lock
        hostPath:
          path: /run/xtables.lock
          type: FileOrCreate
      # Used to install CNI.
      - name: cni-bin-dir
        hostPath:
          path: /opt/cni/bin
      - name: cni-net-dir
        hostPath:
          path: /etc/cni/net.d
      # Used to create per-pod Unix Domain Sockets
      - name: policysync
        hostPath:
          type: DirectoryOrCreate
          path: /var/run/nodeagent
      # Used to install Flex Volume Driver
      - name: flexvol-driver-host
        hostPath:
          type: DirectoryOrCreate
          path: /etc/kubernetes/volumeplugins/nodeagent~uds
---

apiVersion: v1
kind: ServiceAccount
metadata:
  name: calico-node
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
---

# Typha Horizontal Autoscaler ConfigMap
kind: ConfigMap
apiVersion: v1
metadata:
  name: calico-typha-horizontal-autoscaler
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "EnsureExists"
data:
  ladder: |-
    {
      "coresToReplicas": [],
      "nodesToReplicas":
      [
        [1, 1],
        [10, 2],
        [100, 3],
        [250, 4],
        [500, 5],
        [1000, 6],
        [1500, 7],
        [2000, 8]
      ]
    }

---

# Typha Horizontal Autoscaler Deployment
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: calico-typha-horizontal-autoscaler
  namespace: kube-system
  labels:
    k8s-app: calico-typha-autoscaler
    addonmanager.kubernetes.io/mode: "Reconcile"
spec:
  replicas: 1
  template:
    metadata:
      labels:
        k8s-app: calico-typha-autoscaler
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ''
    spec:
      priorityClassName: system-cluster-critical
      securityContext:
        supplementalGroups: [65534]
        fsGroup: 65534
      containers:
      - image: {{ContainerImage "calico-cluster-proportional-autoscaler"}}
        name: autoscaler
        command:
        - /cluster-proportional-autoscaler
        - --namespace=kube-system
        - --configmap=calico-typha-horizontal-autoscaler
        - --target=deployment/calico-typha
        - --logtostderr=true
        - --v=2
        resources:
          requests:
            cpu: 10m
          limits:
            cpu: 10m
      serviceAccountName: typha-cpha
---

# Typha Horizontal Autoscaler Cluster Role
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: typha-cpha
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
rules:
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["list"]

---

# Typha Horizontal Autoscaler Cluster Role Binding
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: typha-cpha
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: typha-cpha
subjects:
- kind: ServiceAccount
  name: typha-cpha
  namespace: kube-system
---

# Typha Horizontal Autoscaler Role
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: typha-cpha
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
rules:
- apiGroups: [""]
  resources: ["configmaps"]
  verbs: ["get"]
- apiGroups: ["extensions"]
  resources: ["deployments/scale"]
  verbs: ["get", "update"]

---

# Typha Horizontal Autoscaler Role Binding
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: typha-cpha
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: typha-cpha
subjects:
- kind: ServiceAccount
  name: typha-cpha
  namespace: kube-system
---

# Typha Horizontal Autoscaler Service Account
apiVersion: v1
kind: ServiceAccount
metadata:
  name: typha-cpha
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "Reconcile"
`)

func k8sContaineraddonsKubernetesmasteraddonsCalicoDaemonsetYamlBytes() ([]byte, error) {
	return _k8sContaineraddonsKubernetesmasteraddonsCalicoDaemonsetYaml, nil
}

func k8sContaineraddonsKubernetesmasteraddonsCalicoDaemonsetYaml() (*asset, error) {
	bytes, err := k8sContaineraddonsKubernetesmasteraddonsCalicoDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/kubernetesmasteraddons-calico-daemonset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddonsKubernetesmasteraddonsClusterAutoscalerDeploymentYaml = []byte(`---
apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    k8s-addon: cluster-autoscaler.addons.k8s.io
    k8s-app: cluster-autoscaler
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: cluster-autoscaler
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: cluster-autoscaler
  labels:
    k8s-addon: cluster-autoscaler.addons.k8s.io
    k8s-app: cluster-autoscaler
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups: [""]
  resources: ["events", "endpoints"]
  verbs: ["create", "patch"]
- apiGroups: [""]
  resources: ["pods/eviction"]
  verbs: ["create"]
- apiGroups: [""]
  resources: ["pods/status"]
  verbs: ["update"]
- apiGroups: [""]
  resources: ["endpoints"]
  resourceNames: ["cluster-autoscaler"]
  verbs: ["get", "update"]
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["watch", "list", "get", "update"]
- apiGroups: [""]
  resources:
  - "pods"
  - "services"
  - "replicationcontrollers"
  - "persistentvolumeclaims"
  - "persistentvolumes"
  verbs: ["watch", "list", "get"]
- apiGroups: ["extensions"]
  resources: ["replicasets", "daemonsets"]
  verbs: ["watch", "list", "get"]
- apiGroups: ["policy"]
  resources: ["poddisruptionbudgets"]
  verbs: ["watch", "list"]
- apiGroups: ["apps"]
  resources: ["statefulsets","replicasets","daemonsets"]
  verbs: ["watch","list","get"]
- apiGroups: ["batch"]
  resources: ["jobs"]
  verbs: ["watch","list"]
- apiGroups: ["storage.k8s.io"]
  resources: ["storageclasses"]
  verbs: ["get", "list", "watch"]
- apiGroups: ["batch"]
  resources: ["jobs", "cronjobs"]
  verbs: ["watch", "list", "get"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: cluster-autoscaler
  namespace: kube-system
  labels:
    k8s-addon: cluster-autoscaler.addons.k8s.io
    k8s-app: cluster-autoscaler
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups: [""]
  resources: ["configmaps"]
  verbs: ["create", "list", "watch"]
- apiGroups: [""]
  resources: ["configmaps"]
  resourceNames:
  - "cluster-autoscaler-status"
  - "cluster-autoscaler-priority-expander"
  verbs: ["delete", "get", "update", "watch"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: cluster-autoscaler
  labels:
    k8s-addon: cluster-autoscaler.addons.k8s.io
    k8s-app: cluster-autoscaler
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-autoscaler
subjects:
  - kind: ServiceAccount
    name: cluster-autoscaler
    namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: cluster-autoscaler
  namespace: kube-system
  labels:
    k8s-addon: cluster-autoscaler.addons.k8s.io
    k8s-app: cluster-autoscaler
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: cluster-autoscaler
subjects:
  - kind: ServiceAccount
    name: cluster-autoscaler
    namespace: kube-system
---
apiVersion: v1
data:
  ClientID: <clientID>
  ClientSecret: <clientSec>
  ResourceGroup: <rg>
  SubscriptionID: <subID>
  TenantID: <tenantID>
  VMType: <vmType>
kind: Secret
metadata:
  name: cluster-autoscaler-azure
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    app: cluster-autoscaler
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: cluster-autoscaler
  namespace: kube-system
spec:
  replicas: 1
  selector:
    matchLabels:
      app: cluster-autoscaler
  template:
    metadata:
      labels:
        app: cluster-autoscaler
    spec:
      priorityClassName: system-node-critical
      <hostNet>
      serviceAccountName: cluster-autoscaler
      tolerations:
      - effect: NoSchedule
        operator: "Equal"
        value: "true"
        key: node-role.kubernetes.io/master
      nodeSelector:
        kubernetes.io/role: master
        beta.kubernetes.io/os: linux
      containers:
      - image: {{ContainerImage "cluster-autoscaler"}}
        imagePullPolicy: IfNotPresent
        name: cluster-autoscaler
        resources:
          limits:
            cpu: {{ContainerCPULimits "cluster-autoscaler"}}
            memory: {{ContainerMemLimits "cluster-autoscaler"}}
          requests:
            cpu: {{ContainerCPUReqs "cluster-autoscaler"}}
            memory: {{ContainerMemReqs "cluster-autoscaler"}}
        command:
        - ./cluster-autoscaler
        - --v=3
        - --logtostderr=true
        - --cloud-provider=azure
        - --skip-nodes-with-local-storage=false
        - --nodes={{ContainerConfig "min-nodes"}}:{{ContainerConfig "max-nodes"}}:<vmssName>
        - --scan-interval={{ContainerConfig "scan-interval"}}
        env:
        - name: ARM_CLOUD
          value: "<cloud>"
        - name: ARM_SUBSCRIPTION_ID
          valueFrom:
            secretKeyRef:
              key: SubscriptionID
              name: cluster-autoscaler-azure
        - name: ARM_RESOURCE_GROUP
          valueFrom:
            secretKeyRef:
              key: ResourceGroup
              name: cluster-autoscaler-azure
        - name: ARM_TENANT_ID
          valueFrom:
            secretKeyRef:
              key: TenantID
              name: cluster-autoscaler-azure
        - name: ARM_CLIENT_ID
          valueFrom:
            secretKeyRef:
              key: ClientID
              name: cluster-autoscaler-azure
        - name: ARM_CLIENT_SECRET
          valueFrom:
            secretKeyRef:
              key: ClientSecret
              name: cluster-autoscaler-azure
        - name: ARM_VM_TYPE
          valueFrom:
            secretKeyRef:
              key: VMType
              name: cluster-autoscaler-azure
        - name: ARM_USE_MANAGED_IDENTITY_EXTENSION
          value: "<useManagedIdentity>"
        volumeMounts:
        - mountPath: /etc/ssl/certs/ca-certificates.crt
          name: ssl-certs
          readOnly: true
        <volMounts>
      restartPolicy: Always
      volumes:
      - hostPath:
          path: /etc/ssl/certs/ca-certificates.crt
          type: ""
        name: ssl-certs
      <vols>
#EOF
`)

func k8sContaineraddonsKubernetesmasteraddonsClusterAutoscalerDeploymentYamlBytes() ([]byte, error) {
	return _k8sContaineraddonsKubernetesmasteraddonsClusterAutoscalerDeploymentYaml, nil
}

func k8sContaineraddonsKubernetesmasteraddonsClusterAutoscalerDeploymentYaml() (*asset, error) {
	bytes, err := k8sContaineraddonsKubernetesmasteraddonsClusterAutoscalerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/kubernetesmasteraddons-cluster-autoscaler-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddonsKubernetesmasteraddonsHeapsterDeploymentYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: heapster
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: system:heapster-with-nanny
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
rules:
- apiGroups:
  - extensions
  - apps
  resources:
  - deployments
  verbs:
  - get
  - list
  - watch
  - update
  - patch
- apiGroups:
  - ""
  resources:
  - events
  - namespaces
  - nodes
  - pods
  verbs:
  - get
  - list
  - watch
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: system:heapster-with-nanny
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:heapster-with-nanny
subjects:
- kind: ServiceAccount
  name: heapster
  namespace: kube-system
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: heapster-config
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
data:
  NannyConfiguration: |-
    apiVersion: nannyconfig/v1alpha1
    kind: NannyConfiguration
---
apiVersion: v1
kind: Service
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
    kubernetes.io/name: Heapster
  name: heapster
  namespace: kube-system
spec:
  ports:
  - port: 80
    targetPort: 8082
  selector:
    k8s-app: heapster
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: heapster
  namespace: kube-system
  labels:
    k8s-app: heapster
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
spec:
  replicas: 1
  selector:
    matchLabels:
      k8s-app: heapster
  template:
    metadata:
      labels:
        k8s-app: heapster
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ''
    spec:
      priorityClassName: system-node-critical
      containers:
        - image: {{ContainerImage "heapster"}}
          imagePullPolicy: IfNotPresent
          name: heapster
          resources:
            requests:
              cpu: {{ContainerCPUReqs "heapster"}}
              memory: {{ContainerMemReqs "heapster"}}
            limits:
              cpu: {{ContainerCPULimits "heapster"}}
              memory: {{ContainerMemLimits "heapster"}}
          livenessProbe:
            httpGet:
              path: /healthz
              port: 8082
              scheme: HTTP
            initialDelaySeconds: 180
            timeoutSeconds: 5
          command:
            - /heapster
            - --source=kubernetes.summary_api:''
        - image: {{ContainerImage "heapster-nanny"}}
          imagePullPolicy: IfNotPresent
          name: heapster-nanny
          resources:
            requests:
              cpu: {{ContainerCPUReqs "heapster-nanny"}}
              memory: {{ContainerMemReqs "heapster-nanny"}}
            limits:
              cpu: {{ContainerCPULimits "heapster-nanny"}}
              memory: {{ContainerMemLimits "heapster-nanny"}}
          env:
            - name: MY_POD_NAME
              valueFrom:
                fieldRef:
                  fieldPath: metadata.name
            - name: MY_POD_NAMESPACE
              valueFrom:
                fieldRef:
                  fieldPath: metadata.namespace
          volumeMounts:
          - name: heapster-config-volume
            mountPath: /etc/config
          command:
            - /pod_nanny
            - --config-dir=/etc/config
            - --cpu=80m
            - --extra-cpu=0.5m
            - --memory=140Mi
            - --extra-memory=4Mi
            - --threshold=5
            - --deployment=heapster
            - --container=heapster
            - --poll-period=300000
            - --estimator=exponential
      volumes:
        - name: heapster-config-volume
          configMap:
            name: heapster-config
      serviceAccountName: heapster
      tolerations:
        - key: CriticalAddonsOnly
          operator: Exists
      nodeSelector:
        beta.kubernetes.io/os: linux
`)

func k8sContaineraddonsKubernetesmasteraddonsHeapsterDeploymentYamlBytes() ([]byte, error) {
	return _k8sContaineraddonsKubernetesmasteraddonsHeapsterDeploymentYaml, nil
}

func k8sContaineraddonsKubernetesmasteraddonsHeapsterDeploymentYaml() (*asset, error) {
	bytes, err := k8sContaineraddonsKubernetesmasteraddonsHeapsterDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/kubernetesmasteraddons-heapster-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddonsKubernetesmasteraddonsKeyvaultFlexvolumeInstallerYaml = []byte(`apiVersion: extensions/v1beta1
kind: DaemonSet
metadata:
  labels:
    app: keyvault-flexvolume
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: keyvault-flexvolume
  namespace: kube-system
spec:
  updateStrategy:
    type: RollingUpdate
  template:
    metadata:
      labels:
        app: keyvault-flexvolume
        kubernetes.io/cluster-service: "true"
        addonmanager.kubernetes.io/mode: Reconcile
    spec:
      tolerations:
      containers:
      - name: keyvault-flexvolume
        image: {{ContainerImage "keyvault-flexvolume"}}
        imagePullPolicy: IfNotPresent
        resources:
          requests:
            cpu: {{ContainerCPUReqs "keyvault-flexvolume"}}
            memory: {{ContainerMemReqs "keyvault-flexvolume"}}
          limits:
            cpu: {{ContainerCPULimits "keyvault-flexvolume"}}
            memory: {{ContainerMemLimits "keyvault-flexvolume"}}
        env:
        - name: TARGET_DIR
          value: /etc/kubernetes/volumeplugins
        volumeMounts:
        - mountPath: /etc/kubernetes/volumeplugins
          name: volplugins
      volumes:
      - hostPath:
          path: /etc/kubernetes/volumeplugins
        name: volplugins
      nodeSelector:
        beta.kubernetes.io/os: linux
`)

func k8sContaineraddonsKubernetesmasteraddonsKeyvaultFlexvolumeInstallerYamlBytes() ([]byte, error) {
	return _k8sContaineraddonsKubernetesmasteraddonsKeyvaultFlexvolumeInstallerYaml, nil
}

func k8sContaineraddonsKubernetesmasteraddonsKeyvaultFlexvolumeInstallerYaml() (*asset, error) {
	bytes, err := k8sContaineraddonsKubernetesmasteraddonsKeyvaultFlexvolumeInstallerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/kubernetesmasteraddons-keyvault-flexvolume-installer.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddonsKubernetesmasteraddonsKubeReschedulerDeploymentYaml = []byte(`apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: rescheduler
  namespace: kube-system
  labels:
    k8s-app: rescheduler
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  replicas: 1
  selector:
    matchLabels:
      k8s-app: rescheduler
  template:
    metadata:
      labels:
        k8s-app: rescheduler
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ''
    spec:
      nodeSelector:
        beta.kubernetes.io/os: linux
      containers:
      - image: {{ContainerImage "rescheduler"}}
        imagePullPolicy: IfNotPresent
        name: rescheduler
        resources:
          requests:
            cpu: {{ContainerCPUReqs "rescheduler"}}
            memory: {{ContainerMemReqs "rescheduler"}}
          limits:
            cpu: {{ContainerCPULimits "rescheduler"}}
            memory: {{ContainerMemLimits "rescheduler"}}
        command:
        - sh
        - -c
        - '/rescheduler'
`)

func k8sContaineraddonsKubernetesmasteraddonsKubeReschedulerDeploymentYamlBytes() ([]byte, error) {
	return _k8sContaineraddonsKubernetesmasteraddonsKubeReschedulerDeploymentYaml, nil
}

func k8sContaineraddonsKubernetesmasteraddonsKubeReschedulerDeploymentYaml() (*asset, error) {
	bytes, err := k8sContaineraddonsKubernetesmasteraddonsKubeReschedulerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/kubernetesmasteraddons-kube-rescheduler-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddonsKubernetesmasteraddonsKubernetesDashboardDeploymentYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    k8s-app: kubernetes-dashboard
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: kubernetes-dashboard
  namespace: kube-system
---
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: kubernetes-dashboard-minimal
  namespace: kube-system
  labels:
    k8s-app: kubernetes-dashboard
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups: [""]
  resources: ["secrets"]
  verbs: ["create"]
- apiGroups: [""]
  resources: ["configmaps"]
  verbs: ["create"]
- apiGroups: [""]
  resources: ["secrets"]
  resourceNames: ["kubernetes-dashboard-key-holder"]
  verbs: ["get", "update", "delete"]
- apiGroups: [""]
  resources: ["configmaps"]
  resourceNames: ["kubernetes-dashboard-settings"]
  verbs: ["get", "update"]
- apiGroups: [""]
  resources: ["services"]
  resourceNames: ["heapster"]
  verbs: ["proxy"]
- apiGroups: [""]
  resources: ["services/proxy"]
  resourceNames: ["heapster", "http:heapster:", "https:heapster:"]
  verbs: ["get"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: kubernetes-dashboard-minimal
  namespace: kube-system
  labels:
    k8s-app: kubernetes-dashboard
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: kubernetes-dashboard-minimal
subjects:
- kind: ServiceAccount
  name: kubernetes-dashboard
  namespace: kube-system
---
apiVersion: v1
kind: Service
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kube-system
spec:
  ports:
  - port: 443
    targetPort: 8443
  selector:
    k8s-app: kubernetes-dashboard
  type: NodePort
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kube-system
spec:
  replicas: 1
  selector:
    matchLabels:
      k8s-app: kubernetes-dashboard
  template:
    metadata:
      labels:
        k8s-app: kubernetes-dashboard
    spec:
      containers:
      - args:
        - --auto-generate-certificates
        - --heapster-host=http://heapster.kube-system:80
        image: {{ContainerImage "kubernetes-dashboard"}}
        imagePullPolicy: IfNotPresent
        livenessProbe:
          httpGet:
            path: "/"
            port: 8443
            scheme: HTTPS
          initialDelaySeconds: 30
          timeoutSeconds: 30
        name: kubernetes-dashboard
        ports:
        - containerPort: 8443
          protocol: TCP
        resources:
          requests:
            cpu: {{ContainerCPUReqs "kubernetes-dashboard"}}
            memory: {{ContainerMemReqs "kubernetes-dashboard"}}
          limits:
            cpu: {{ContainerCPULimits "kubernetes-dashboard"}}
            memory: {{ContainerMemLimits "kubernetes-dashboard"}}
        volumeMounts:
         - name: kubernetes-dashboard-certs
           mountPath: /certs
      volumes:
        - name: kubernetes-dashboard-certs
          emptyDir: {}
      serviceAccountName: kubernetes-dashboard
      nodeSelector:
        beta.kubernetes.io/os: linux
`)

func k8sContaineraddonsKubernetesmasteraddonsKubernetesDashboardDeploymentYamlBytes() ([]byte, error) {
	return _k8sContaineraddonsKubernetesmasteraddonsKubernetesDashboardDeploymentYaml, nil
}

func k8sContaineraddonsKubernetesmasteraddonsKubernetesDashboardDeploymentYaml() (*asset, error) {
	bytes, err := k8sContaineraddonsKubernetesmasteraddonsKubernetesDashboardDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/kubernetesmasteraddons-kubernetes-dashboard-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddonsKubernetesmasteraddonsMetricsServerDeploymentYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: metrics-server
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: system:metrics-server
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups:
  - ""
  resources:
  - pods
  - nodes
  - nodes/stats
  - namespaces
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - "extensions"
  resources:
  - deployments
  verbs:
  - get
  - list
  - watch
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system:metrics-server
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:metrics-server
subjects:
- kind: ServiceAccount
  name: metrics-server
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: RoleBinding
metadata:
  name: metrics-server-auth-reader
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: extension-apiserver-authentication-reader
subjects:
- kind: ServiceAccount
  name: metrics-server
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: metrics-server:system:auth-delegator
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:auth-delegator
subjects:
- kind: ServiceAccount
  name: metrics-server
  namespace: kube-system
---
apiVersion: v1
kind: Service
metadata:
  name: metrics-server
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
    kubernetes.io/name: "Metrics-server"
    kubernetes.io/cluster-service: "true"
spec:
  selector:
    k8s-app: metrics-server
  ports:
  - port: 443
    protocol: TCP
    targetPort: 443
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: metrics-server
  namespace: kube-system
  labels:
    k8s-app: metrics-server
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  selector:
    matchLabels:
      k8s-app: metrics-server
  template:
    metadata:
      name: metrics-server
      labels:
        k8s-app: metrics-server
    spec:
      serviceAccountName: metrics-server
      containers:
      - name: metrics-server
        image: {{ContainerImage "metrics-server"}}
        imagePullPolicy: IfNotPresent
        command:
        - /metrics-server
        - --source=kubernetes.summary_api:''
      nodeSelector:
        beta.kubernetes.io/os: linux
---
apiVersion: apiregistration.k8s.io/v1beta1
kind: APIService
metadata:
  name: v1beta1.metrics.k8s.io
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  service:
    name: metrics-server
    namespace: kube-system
  group: metrics.k8s.io
  version: v1beta1
  insecureSkipTLSVerify: true
  groupPriorityMinimum: 100
  versionPriority: 100
`)

func k8sContaineraddonsKubernetesmasteraddonsMetricsServerDeploymentYamlBytes() ([]byte, error) {
	return _k8sContaineraddonsKubernetesmasteraddonsMetricsServerDeploymentYaml, nil
}

func k8sContaineraddonsKubernetesmasteraddonsMetricsServerDeploymentYaml() (*asset, error) {
	bytes, err := k8sContaineraddonsKubernetesmasteraddonsMetricsServerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/kubernetesmasteraddons-metrics-server-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddonsKubernetesmasteraddonsNvidiaDevicePluginDaemonsetYaml = []byte(`apiVersion: apps/v1
kind: DaemonSet
metadata:
  labels:
    k8s-app: nvidia-device-plugin
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: nvidia-device-plugin
  namespace: kube-system
spec:
  selector:
    matchLabels:
      k8s-app: nvidia-device-plugin
  updateStrategy:
    type: RollingUpdate
  template:
    metadata:
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ""
      labels:
        k8s-app: nvidia-device-plugin
    spec:
      priorityClassName: system-node-critical
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: accelerator
                operator: In
                values:
                - nvidia
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      - key: nvidia.com/gpu
        effect: NoSchedule
        operator: Equal
        value: "true"
      containers:
      - image: {{ContainerImage "nvidia-device-plugin"}}
        name: nvidia-device-plugin-ctr
        resources:
          requests:
            cpu: {{ContainerCPUReqs "nvidia-device-plugin"}}
            memory: {{ContainerMemReqs "nvidia-device-plugin"}}
          limits:
            cpu: {{ContainerCPULimits "nvidia-device-plugin"}}
            memory: {{ContainerMemLimits "nvidia-device-plugin"}}
        securityContext:
          allowPrivilegeEscalation: false
          capabilities:
            drop: ["ALL"]
        volumeMounts:
          - name: device-plugin
            mountPath: /var/lib/kubelet/device-plugins
      volumes:
        - name: device-plugin
          hostPath:
            path: /var/lib/kubelet/device-plugins
      nodeSelector:
        beta.kubernetes.io/os: linux
        accelerator: nvidia
`)

func k8sContaineraddonsKubernetesmasteraddonsNvidiaDevicePluginDaemonsetYamlBytes() ([]byte, error) {
	return _k8sContaineraddonsKubernetesmasteraddonsNvidiaDevicePluginDaemonsetYaml, nil
}

func k8sContaineraddonsKubernetesmasteraddonsNvidiaDevicePluginDaemonsetYaml() (*asset, error) {
	bytes, err := k8sContaineraddonsKubernetesmasteraddonsNvidiaDevicePluginDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/kubernetesmasteraddons-nvidia-device-plugin-daemonset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddonsKubernetesmasteraddonsOmsagentDaemonsetYaml = []byte(`apiVersion: v1
kind: Secret
metadata:
  name: omsagent-secret
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
type: Opaque
data:
  WSID: "{{ContainerConfig "workspaceGuid"}}"
  KEY: "{{ContainerConfig "workspaceKey"}}"
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: omsagent
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: omsagent-reader
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups: [""]
  resources: ["pods", "events", "nodes", "namespaces", "services"]
  verbs: ["list", "get", "watch"]
- nonResourceURLs: ["/metrics"]
  verbs: ["get"]
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: omsagentclusterrolebinding
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
subjects:
  - kind: ServiceAccount
    name: omsagent
    namespace: kube-system
roleRef:
  kind: ClusterRole
  name: omsagent-reader
  apiGroup: rbac.authorization.k8s.io
---
kind: ConfigMap
apiVersion: v1
data:
  kube.conf: |-
     # Fluentd config file for OMS Docker - cluster components (kubeAPI)
     #Kubernetes pod inventory
     <source>
      type kubepodinventory
      tag oms.containerinsights.KubePodInventory
      run_interval 60s
      log_level debug
     </source>

     #Kubernetes events
     <source>
      type kubeevents
      tag oms.containerinsights.KubeEvents
      run_interval 60s
      log_level debug
      </source>

     #Kubernetes logs
     <source>
      type kubelogs
      tag oms.api.KubeLogs
      run_interval 60s
     </source>

     #Kubernetes services
     <source>
      type kubeservices
      tag oms.containerinsights.KubeServices
      run_interval 60s
      log_level debug
     </source>

     #Kubernetes Nodes
     <source>
      type kubenodeinventory
      tag oms.containerinsights.KubeNodeInventory
      run_interval 60s
      log_level debug
     </source>

     #Kubernetes perf
     <source>
      type kubeperf
      tag oms.api.KubePerf
      run_interval 60s
      log_level debug
     </source>

     #cadvisor perf- Windows nodes
     <source>
      type wincadvisorperf
      tag oms.api.wincadvisorperf
      run_interval 60s
      log_level debug
     </source>

     <filter mdm.kubepodinventory** mdm.kubenodeinventory**>
      type filter_inventory2mdm
      custom_metrics_azure_regions eastus,southcentralus,westcentralus,westus2,southeastasia,northeurope,westEurope
      log_level info
     </filter>

     # custom_metrics_mdm filter plugin for perf data from windows nodes
     <filter mdm.cadvisorperf**>
      type filter_cadvisor2mdm
      custom_metrics_azure_regions eastus,southcentralus,westcentralus,westus2,southeastasia,northeurope,westEurope
      metrics_to_collect cpuUsageNanoCores,memoryWorkingSetBytes
      log_level info
     </filter>

     <match oms.containerinsights.KubePodInventory**>
      type out_oms
      log_level debug
      num_threads 5
      buffer_chunk_limit 20m
      buffer_type file
      buffer_path %STATE_DIR_WS%/out_oms_kubepods*.buffer
      buffer_queue_limit 20
      buffer_queue_full_action drop_oldest_chunk
      flush_interval 20s
      retry_limit 10
      retry_wait 30s
      max_retry_wait 9m
     </match>

     <match oms.containerinsights.KubeEvents**>
      type out_oms
      log_level debug
      num_threads 5
      buffer_chunk_limit 5m
      buffer_type file
      buffer_path %STATE_DIR_WS%/out_oms_kubeevents*.buffer
      buffer_queue_limit 10
      buffer_queue_full_action drop_oldest_chunk
      flush_interval 20s
      retry_limit 10
      retry_wait 30s
      max_retry_wait 9m
     </match>

     <match oms.api.KubeLogs**>
      type out_oms_api
      log_level debug
      buffer_chunk_limit 10m
      buffer_type file
      buffer_path %STATE_DIR_WS%/out_oms_api_kubernetes_logs*.buffer
      buffer_queue_limit 10
      flush_interval 20s
      retry_limit 10
      retry_wait 30s
     </match>

     <match oms.containerinsights.KubeServices**>
      type out_oms
      log_level debug
      num_threads 5
      buffer_chunk_limit 20m
      buffer_type file
      buffer_path %STATE_DIR_WS%/out_oms_kubeservices*.buffer
      buffer_queue_limit 20
      buffer_queue_full_action drop_oldest_chunk
      flush_interval 20s
      retry_limit 10
      retry_wait 30s
      max_retry_wait 9m
     </match>

     <match oms.containerinsights.KubeNodeInventory**>
      type out_oms
      log_level debug
      num_threads 5
      buffer_chunk_limit 20m
      buffer_type file
      buffer_path %STATE_DIR_WS%/state/out_oms_kubenodes*.buffer
      buffer_queue_limit 20
      buffer_queue_full_action drop_oldest_chunk
      flush_interval 20s
      retry_limit 10
      retry_wait 30s
      max_retry_wait 9m
     </match>

     <match oms.containerinsights.ContainerNodeInventory**>
      type out_oms_api
      log_level debug
      buffer_chunk_limit 20m
      buffer_type file
      buffer_path %STATE_DIR_WS%/out_oms_containernodeinventory*.buffer
      buffer_queue_limit 20
      flush_interval 20s
      retry_limit 10
      retry_wait 15s
      max_retry_wait 9m
     </match>

     <match oms.api.KubePerf**>
      type out_oms
      log_level debug
      num_threads 5
      buffer_chunk_limit 20m
      buffer_type file
      buffer_path %STATE_DIR_WS%/out_oms_kubeperf*.buffer
      buffer_queue_limit 20
      buffer_queue_full_action drop_oldest_chunk
      flush_interval 20s
      retry_limit 10
      retry_wait 30s
      max_retry_wait 9m
     </match>

     <match mdm.kubepodinventory** mdm.kubenodeinventory** >
      type out_mdm
      log_level debug
      num_threads 5
      buffer_chunk_limit 20m
      buffer_type file
      buffer_path %STATE_DIR_WS%/out_mdm_*.buffer
      buffer_queue_limit 20
      buffer_queue_full_action drop_oldest_chunk
      flush_interval 20s
      retry_limit 10
      retry_wait 30s
      max_retry_wait 9m
      retry_mdm_post_wait_minutes 60
     </match>

     <match oms.api.wincadvisorperf**>
      type out_oms
      log_level debug
      num_threads 5
      buffer_chunk_limit 20m
      buffer_type file
      buffer_path %STATE_DIR_WS%/out_oms_api_wincadvisorperf*.buffer
      buffer_queue_limit 20
      buffer_queue_full_action drop_oldest_chunk
      flush_interval 20s
      retry_limit 10
      retry_wait 30s
      max_retry_wait 9m
     </match>

     <match mdm.cadvisorperf**>
      type out_mdm
      log_level debug
      num_threads 5
      buffer_chunk_limit 20m
      buffer_type file
      buffer_path %STATE_DIR_WS%/out_mdm_cdvisorperf*.buffer
      buffer_queue_limit 20
      buffer_queue_full_action drop_oldest_chunk
      flush_interval 20s
      retry_limit 10
      retry_wait 30s
      max_retry_wait 9m
      retry_mdm_post_wait_minutes 60
     </match>
metadata:
  name: omsagent-rs-config
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: extensions/v1beta1
kind: DaemonSet
metadata:
  labels:
    component: oms-agent
    tier: node
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: omsagent
  namespace: kube-system
spec:
  selector:
    matchLabels:
      component: oms-agent
      tier: node
  template:
    metadata:
      annotations:
        agentVersion: {{ContainerConfig "omsAgentVersion"}}
        dockerProviderVersion: {{ContainerConfig "dockerProviderVersion"}}
        schema-versions:  {{ContainerConfig "schema-versions"}}
      labels:
        component: oms-agent
        tier: node
    spec:
      priorityClassName: system-node-critical
      serviceAccountName: omsagent
      containers:
        - name: omsagent
          image: {{ContainerImage "omsagent"}}
          imagePullPolicy: Always
          resources:
            limits:
              cpu: {{ContainerCPULimits "omsagent"}}
              memory: {{ContainerMemLimits "omsagent"}}
            requests:
              cpu: {{ContainerCPUReqs "omsagent"}}
              memory: {{ContainerMemReqs "omsagent"}}
          env:
            - name: NODE_IP
              valueFrom:
                fieldRef:
                  fieldPath: status.hostIP
            - name: ACS_RESOURCE_NAME
              value: {{ContainerConfig "clusterName"}}            
            - name: CONTROLLER_TYPE
              value: "DaemonSet"
            - name: ISTEST
              value: "true"
          livenessProbe:
            exec:
              command:
                - /bin/bash
                - -c
                - /opt/livenessprobe.sh
            initialDelaySeconds: 60
            periodSeconds: 60
          ports:
            - containerPort: 25225
              protocol: TCP
            - containerPort: 25224
              protocol: UDP
          securityContext:
            privileged: true
          volumeMounts:
            - mountPath: /hostfs
              name: host-root
            - mountPath: /var/run/host
              name: docker-sock
            - mountPath: /var/log
              name: host-log
            - mountPath: /var/lib/docker/containers
              name: containerlog-path
            - mountPath: /etc/kubernetes/host
              name: azure-json-path
            - mountPath: /etc/omsagent-secret
              name: omsagent-secret
              readOnly: true
            - mountPath: /etc/config/settings
              name: settings-vol-config  
      nodeSelector:
        beta.kubernetes.io/os: linux
      tolerations:
        - effect: NoSchedule
          key: node-role.kubernetes.io/master
          operator: Equal
          value: "true"
      volumes:
        - name: host-root
          hostPath:
            path: /
        - name: docker-sock
          hostPath:
            path: /var/run
        - name: container-hostname
          hostPath:
            path: /etc/hostname
        - name: host-log
          hostPath:
            path: /var/log
        - name: containerlog-path
          hostPath:
            path: /var/lib/docker/containers
        - name: azure-json-path
          hostPath:
            path: /etc/kubernetes
        - name: omsagent-secret
          secret:
            secretName: omsagent-secret
        - name: settings-vol-config
          configMap:
            name: container-azm-ms-agentconfig
            optional: true    
  updateStrategy:
    type: RollingUpdate
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: omsagent-rs
  namespace: kube-system
  labels:
    component: oms-agent
    tier: node
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  replicas: 1
  selector:
    matchLabels:
      rsName: omsagent-rs
  strategy:
    type: RollingUpdate
  template:
    metadata:
      labels:
        rsName: omsagent-rs
      annotations:
        agentVersion: {{ContainerConfig "omsAgentVersion"}}
        dockerProviderVersion: {{ContainerConfig "dockerProviderVersion"}}
        schema-versions:  {{ContainerConfig "schema-versions"}}
    spec:
      serviceAccountName: omsagent
      containers:
        - name: omsagent
          image: {{ContainerImage "omsagent"}}
          imagePullPolicy: IfNotPresent
          resources:
            limits:
              cpu: {{ContainerCPULimits "omsagent"}}
              memory: {{ContainerMemLimits "omsagent"}}
            requests:
              cpu: {{ContainerCPUReqs "omsagent"}}
              memory: {{ContainerMemReqs "omsagent"}}
          env:
            - name: NODE_IP
              valueFrom:
                fieldRef:
                  fieldPath: status.hostIP
            - name: ACS_RESOURCE_NAME
              value: {{ContainerConfig "clusterName"}}            
            - name: CONTROLLER_TYPE
              value: "ReplicaSet"
            - name: ISTEST
              value: "true"
          securityContext:
            privileged: true
          ports:
            - containerPort: 25225
              protocol: TCP
            - containerPort: 25224
              protocol: UDP
          volumeMounts:
            - mountPath: /var/run/host
              name: docker-sock
            - mountPath: /var/log
              name: host-log
            - mountPath: /var/lib/docker/containers
              name: containerlog-path
            - mountPath: /etc/kubernetes/host
              name: azure-json-path
            - mountPath: /etc/omsagent-secret
              name: omsagent-secret
              readOnly: true
            - mountPath : /etc/config
              name: omsagent-rs-config
            - mountPath: /etc/config/settings
              name: settings-vol-config              
          livenessProbe:
            exec:
              command:
                - /bin/bash
                - -c
                - /opt/livenessprobe.sh
            initialDelaySeconds: 60
            periodSeconds: 60
      nodeSelector:
        beta.kubernetes.io/os: linux
        kubernetes.io/role: agent
      volumes:
        - name: docker-sock
          hostPath:
            path: /var/run
        - name: container-hostname
          hostPath:
            path: /etc/hostname
        - name: host-log
          hostPath:
            path: /var/log
        - name: containerlog-path
          hostPath:
            path: /var/lib/docker/containers
        - name: azure-json-path
          hostPath:
            path: /etc/kubernetes
        - name: omsagent-secret
          secret:
            secretName: omsagent-secret
        - name: omsagent-rs-config
          configMap:
            name: omsagent-rs-config
        - name: settings-vol-config
          configMap:
            name: container-azm-ms-agentconfig
            optional: true    
`)

func k8sContaineraddonsKubernetesmasteraddonsOmsagentDaemonsetYamlBytes() ([]byte, error) {
	return _k8sContaineraddonsKubernetesmasteraddonsOmsagentDaemonsetYaml, nil
}

func k8sContaineraddonsKubernetesmasteraddonsOmsagentDaemonsetYaml() (*asset, error) {
	bytes, err := k8sContaineraddonsKubernetesmasteraddonsOmsagentDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/kubernetesmasteraddons-omsagent-daemonset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddonsKubernetesmasteraddonsSmbFlexvolumeInstallerYaml = []byte(`apiVersion: extensions/v1beta1
kind: DaemonSet
metadata:
  name: smb-flexvol-installer
  namespace: kube-system
  labels:
    k8s-app: smb
    kubernetes.io/cluster-service: "true"
spec:
  selector:
    matchLabels:
      name: smb
  template:
    metadata:
      labels:
        name: smb
        kubernetes.io/cluster-service: "true"
    spec:
      containers:
      - name: smb-flexvol-installer
        image: {{ContainerImage "smb-flexvolume"}}
        imagePullPolicy: Always
        resources:
          requests:
            cpu: {{ContainerCPUReqs "smb-flexvolume"}}
            memory: {{ContainerMemReqs "smb-flexvolume"}}
          limits:
            cpu: {{ContainerCPULimits "smb-flexvolume"}}
            memory: {{ContainerMemLimits "smb-flexvolume"}}
        volumeMounts:
        - name: volplugins
          mountPath: /etc/kubernetes/volumeplugins/
        - name: varlog
          mountPath: /var/log/      
      volumes:
      - name: varlog
        hostPath:
          path: /var/log/              
      - name: volplugins
        hostPath:
          path: /etc/kubernetes/volumeplugins/
          type: DirectoryOrCreate
      nodeSelector:
        beta.kubernetes.io/os: linux
`)

func k8sContaineraddonsKubernetesmasteraddonsSmbFlexvolumeInstallerYamlBytes() ([]byte, error) {
	return _k8sContaineraddonsKubernetesmasteraddonsSmbFlexvolumeInstallerYaml, nil
}

func k8sContaineraddonsKubernetesmasteraddonsSmbFlexvolumeInstallerYaml() (*asset, error) {
	bytes, err := k8sContaineraddonsKubernetesmasteraddonsSmbFlexvolumeInstallerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/kubernetesmasteraddons-smb-flexvolume-installer.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sContaineraddonsKubernetesmasteraddonsTillerDeploymentYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: tiller
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: tiller
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: tiller
  namespace: kube-system
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: helm
    name: tiller
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: tiller-deploy
  namespace: kube-system
spec:
  ports:
  - name: tiller
    port: 44134
    targetPort: tiller
  selector:
    app: helm
    name: tiller
  type: ClusterIP
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    app: helm
    name: tiller
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: tiller-deploy
  namespace: kube-system
spec:
  template:
    metadata:
      labels:
        app: helm
        name: tiller
    spec:
      serviceAccountName: tiller
      containers:
      - env:
        - name: TILLER_NAMESPACE
          value: kube-system
        - name: TILLER_HISTORY_MAX
          value: "{{ContainerConfig "max-history"}}"
        image: {{ContainerImage "tiller"}}
        imagePullPolicy: IfNotPresent
        livenessProbe:
          httpGet:
            path: /liveness
            port: 44135
          initialDelaySeconds: 1
          timeoutSeconds: 1
        name: tiller
        ports:
        - containerPort: 44134
          name: tiller
        readinessProbe:
          httpGet:
            path: /readiness
            port: 44135
          initialDelaySeconds: 1
          timeoutSeconds: 1
        resources:
          requests:
            cpu: {{ContainerCPUReqs "tiller"}}
            memory: {{ContainerMemReqs "tiller"}}
          limits:
            cpu: {{ContainerCPULimits "tiller"}}
            memory: {{ContainerMemLimits "tiller"}}
      nodeSelector:
        beta.kubernetes.io/os: linux
`)

func k8sContaineraddonsKubernetesmasteraddonsTillerDeploymentYamlBytes() ([]byte, error) {
	return _k8sContaineraddonsKubernetesmasteraddonsTillerDeploymentYaml, nil
}

func k8sContaineraddonsKubernetesmasteraddonsTillerDeploymentYaml() (*asset, error) {
	bytes, err := k8sContaineraddonsKubernetesmasteraddonsTillerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containeraddons/kubernetesmasteraddons-tiller-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sKubeconfigJson = []byte(`    {
        "apiVersion": "v1",
        "clusters": [
            {
                "cluster": {
                    "certificate-authority-data": "{{WrapAsVerbatim "parameters('caCertificate')"}}",
                    "server": "https://{{WrapAsVerbatim "reference(concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))).dnsSettings.fqdn"}}"
                },
                "name": "{{WrapAsVariable "resourceGroup"}}"
            }
        ],
        "contexts": [
            {
                "context": {
                    "cluster": "{{WrapAsVariable "resourceGroup"}}",
                    "user": "{{WrapAsVariable "resourceGroup"}}-admin"
                },
                "name": "{{WrapAsVariable "resourceGroup"}}"
            }
        ],
        "current-context": "{{WrapAsVariable "resourceGroup"}}",
        "kind": "Config",
        "users": [
            {
                "name": "{{WrapAsVariable "resourceGroup"}}-admin",
                "user": {{authInfo}}
            }
        ]
    }
`)

func k8sKubeconfigJsonBytes() ([]byte, error) {
	return _k8sKubeconfigJson, nil
}

func k8sKubeconfigJson() (*asset, error) {
	bytes, err := k8sKubeconfigJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/kubeconfig.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sKubernetesparamsT = []byte(`{{if .HasAadProfile}}
    "aadTenantId": {
      "defaultValue": "",
      "metadata": {
        "description": "The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription."
      },
      "type": "string"
    },
    "aadAdminGroupId": {
      "defaultValue": "",
      "metadata": {
        "description": "The AAD default Admin group Object ID used to create a cluster-admin RBAC role."
      },
      "type": "string"
    },
{{end}}
{{if IsHostedMaster}}
    "kubernetesEndpoint": {
      "metadata": {
        "description": "The Kubernetes API endpoint https://<kubernetesEndpoint>:443"
      },
      "type": "string"
    },
{{else}}
    "etcdServerCertificate": {
      "metadata": {
        "description": "The base 64 server certificate used on the master"
      },
      "type": "string"
    },
    "etcdServerPrivateKey": {
      "metadata": {
        "description": "The base 64 server private key used on the master."
      },
      "type": "securestring"
    },
    "etcdClientCertificate": {
      "metadata": {
        "description": "The base 64 server certificate used on the master"
      },
      "type": "string"
    },
    "etcdClientPrivateKey": {
      "metadata": {
        "description": "The base 64 server private key used on the master."
      },
      "type": "securestring"
    },
    "etcdPeerCertificate0": {
      "metadata": {
        "description": "The base 64 server certificates used on the master"
      },
      "type": "string"
    },
    "etcdPeerPrivateKey0": {
      "metadata": {
        "description": "The base 64 server private keys used on the master."
      },
      "type": "securestring"
    },
    {{if ge .MasterProfile.Count 3}}
      "etcdPeerCertificate1": {
        "metadata": {
          "description": "The base 64 server certificates used on the master"
        },
        "type": "string"
      },
      "etcdPeerCertificate2": {
        "metadata": {
          "description": "The base 64 server certificates used on the master"
        },
        "type": "string"
      },
      "etcdPeerPrivateKey1": {
        "metadata": {
          "description": "The base 64 server private keys used on the master."
        },
        "type": "securestring"
      },
      "etcdPeerPrivateKey2": {
        "metadata": {
          "description": "The base 64 server private keys used on the master."
        },
        "type": "securestring"
      },
      {{if ge .MasterProfile.Count 5}}
        "etcdPeerCertificate3": {
          "metadata": {
            "description": "The base 64 server certificates used on the master"
          },
          "type": "string"
        },
        "etcdPeerCertificate4": {
          "metadata": {
            "description": "The base 64 server certificates used on the master"
          },
          "type": "string"
        },
        "etcdPeerPrivateKey3": {
          "metadata": {
            "description": "The base 64 server private keys used on the master."
          },
          "type": "securestring"
        },
        "etcdPeerPrivateKey4": {
          "metadata": {
            "description": "The base 64 server private keys used on the master."
          },
          "type": "securestring"
        },
      {{end}}
    {{end}}
{{end}}
    "apiServerCertificate": {
      "metadata": {
        "description": "The base 64 server certificate used on the master"
      },
      "type": "string"
    },
    "apiServerPrivateKey": {
      "metadata": {
        "description": "The base 64 server private key used on the master."
      },
      "type": "securestring"
    },
    "caCertificate": {
      "metadata": {
        "description": "The base 64 certificate authority certificate"
      },
      "type": "string"
    },
    "caPrivateKey": {
      "metadata": {
        "description": "The base 64 CA private key used on the master."
      },
      "type": "securestring"
    },
    "clientCertificate": {
      "metadata": {
        "description": "The base 64 client certificate used to communicate with the master"
      },
      "type": "string"
    },
    "clientPrivateKey": {
      "metadata": {
        "description": "The base 64 client private key used to communicate with the master"
      },
      "type": "securestring"
    },
    "kubeConfigCertificate": {
      "metadata": {
        "description": "The base 64 certificate used by cli to communicate with the master"
      },
      "type": "string"
    },
    "kubeConfigPrivateKey": {
      "metadata": {
        "description": "The base 64 private key used by cli to communicate with the master"
      },
      "type": "securestring"
    },
    "generatorCode": {
      "metadata": {
        "description": "The generator code used to identify the generator"
      },
      "type": "string"
    },
    "orchestratorName": {
      "metadata": {
        "description": "The orchestrator name used to identify the orchestrator.  This must be no more than 3 digits in length, otherwise it will exceed Windows Naming"
      },
      "minLength": 3,
      "maxLength": 3,
      "type": "string"
    },
    "dockerBridgeCidr": {
      "metadata": {
        "description": "Docker bridge network IP address and subnet"
      },
      "type": "string"
    },
    "kubeClusterCidr": {
      "metadata": {
        "description": "Kubernetes cluster subnet"
      },
      "type": "string"
    },
    "kubeDNSServiceIP": {
      "metadata": {
        "description": "Kubernetes DNS IP"
      },
      "type": "string"
    },
    "kubernetesKubeletClusterDomain": {
      "metadata": {
        "description": "--cluster-domain Kubelet config"
      },
      "type": "string"
    },
    "kubernetesHyperkubeSpec": {
      "metadata": {
        "description": "The container spec for hyperkube."
      },
      "type": "string"
    },
    "privateAzureRegistryServer": {
      "defaultValue": "",
      "metadata": {
        "description": "The private Azure registry server for hyperkube."
      },
      "type": "string"
    },
    "kubernetesCcmImageSpec": {
      "defaultValue": "",
      "metadata": {
        "description": "The container spec for cloud-controller-manager."
      },
      "type": "string"
    },
    "kubernetesAddonManagerSpec": {
      "metadata": {
        "description": "The container spec for hyperkube."
      },
      "type": "string"
    },
    "enableAggregatedAPIs": {
      "metadata": {
        "description": "Enable aggregated API on master nodes"
      },
      "defaultValue": false,
      "type": "bool"
    },
{{if NeedsKubeDNSWithExecHealthz}}
    "kubernetesExecHealthzSpec": {
      "metadata": {
        "description": "The container spec for exechealthz-amd64."
      },
      "type": "string"
    },
{{end}}
    "kubernetesDNSSidecarSpec": {
      "metadata": {
        "description": "The container spec for k8s-dns-sidecar-amd64."
      },
      "type": "string"
    },
{{if .OrchestratorProfile.KubernetesConfig.IsAADPodIdentityEnabled}}
    "kubernetesAADPodIdentityEnabled": {
      "defaultValue": false,
      "metadata": {
        "description": "AAD Pod Identity status"
      },
      "type": "bool"
    },
{{end}}
    "kubernetesACIConnectorEnabled": {
      "metadata": {
        "description": "ACI Connector Status"
      },
      "type": "bool"
    },
    "kubernetesClusterAutoscalerEnabled": {
      "metadata": {
        "description": "Cluster autoscaler status"
      },
      "type": "bool"
    },
{{if .OrchestratorProfile.KubernetesConfig.IsClusterAutoscalerEnabled}}
    "kubernetesClusterAutoscalerAzureCloud": {
      "metadata": {
        "description": "Name of the Azure cloud for the cluster autoscaler."
      },
      "type": "string"
    },
    "kubernetesClusterAutoscalerUseManagedIdentity": {
      "metadata": {
        "description": "Managed identity for the cluster autoscaler addon"
      },
      "type": "string"
    },
{{end}}
    "kubernetesPodInfraContainerSpec": {
      "metadata": {
        "description": "The container spec for pod infra."
      },
      "type": "string"
    },
    "cloudproviderConfig": {
      "type": "object",
      "defaultValue": {
        "cloudProviderBackoff": true,
        "cloudProviderBackoffRetries": 10,
        "cloudProviderBackoffJitter": "0",
        "cloudProviderBackoffDuration": 0,
        "cloudProviderBackoffExponent": "0",
        "cloudProviderRateLimit": false,
        "cloudProviderRateLimitQPS": "0",
        "cloudProviderRateLimitQPSWrite": "0",
        "cloudProviderRateLimitBucket": 0,
        "cloudProviderRateLimitBucketWrite": 0
      }
    },
{{if IsKubernetesVersionGe "1.12.0"}}
    "kubernetesCoreDNSSpec": {
      "metadata": {
        "description": "The container spec for coredns"
      },
      "type": "string"
    },
{{else}}
    "kubernetesKubeDNSSpec": {
      "metadata": {
        "description": "The container spec for kubedns-amd64."
      },
      "type": "string"
    },
    "kubernetesDNSMasqSpec": {
      "metadata": {
        "description": "The container spec for kube-dnsmasq-amd64."
      },
      "type": "string"
    },
{{end}}
    "mobyVersion": {
      "defaultValue": "3.0.6",
      "metadata": {
        "description": "The Azure Moby build version"
      },
      "allowedValues": [
         "3.0.1",
         "3.0.2",
         "3.0.3",
         "3.0.4",
         "3.0.5",
         "3.0.6"
       ],
      "type": "string"
    },
    "containerdVersion": {
      "defaultValue": "1.1.5",
      "metadata": {
        "description": "The Azure Moby build version"
      },
      "allowedValues": [
         "1.1.5",
         "1.1.6",
         "1.2.4"
       ],
      "type": "string"
    },
    "networkPolicy": {
      "defaultValue": "{{.OrchestratorProfile.KubernetesConfig.NetworkPolicy}}",
      "metadata": {
        "description": "The network policy enforcement to use (calico|cilium); 'none' and 'azure' here for backwards compatibility"
      },
      "allowedValues": [
        "",
        "none",
        "azure",
        "calico",
        "cilium"
      ],
      "type": "string"
    },
    "networkPlugin": {
      "defaultValue": "{{.OrchestratorProfile.KubernetesConfig.NetworkPlugin}}",
      "metadata": {
        "description": "The network plugin to use for Kubernetes (kubenet|azure|flannel|cilium)"
      },
      "allowedValues": [
        "kubenet",
        "azure",
        "flannel",
        "cilium"
      ],
      "type": "string"
    },
    "containerRuntime": {
      "defaultValue": "{{.OrchestratorProfile.KubernetesConfig.ContainerRuntime}}",
      "metadata": {
        "description": "The container runtime to use (docker|kata-containers|containerd)"
      },
      "allowedValues": [
        "docker",
        "kata-containers",
        "containerd"
      ],
      "type": "string"
    },
    "containerdDownloadURLBase": {
      "defaultValue": "https://storage.googleapis.com/cri-containerd-release/",
      "type": "string"
    },
    "cniPluginsURL": {
      "defaultValue": "https://acs-mirror.azureedge.net/cni/cni-plugins-amd64-latest.tgz",
      "type": "string"
    },
    "vnetCniLinuxPluginsURL": {
      "defaultValue": "https://acs-mirror.azureedge.net/cni/azure-vnet-cni-linux-amd64-latest.tgz",
      "type": "string"
    },
    "vnetCniWindowsPluginsURL": {
      "defaultValue": "https://acs-mirror.azureedge.net/cni/azure-vnet-cni-windows-amd64-latest.zip",
      "type": "string"
    },
    "maxPods": {
      "defaultValue": 30,
      "metadata": {
        "description": "This param has been deprecated."
      },
      "type": "int"
    },
    "vnetCidr": {
      "defaultValue": "{{GetDefaultVNETCIDR}}",
      "metadata": {
        "description": "Cluster vnet cidr"
      },
      "type": "string"
    },
    "vnetCidrIPv6": {
      "defaultValue": "{{GetDefaultVNETCIDRIPv6}}",
      "metadata": {
        "description": "Cluster vnet cidr IPv6"
      },
      "type": "string"
    },
    "gcHighThreshold": {
      "defaultValue": 85,
      "metadata": {
        "description": "High Threshold for Image Garbage collection on each node"
      },
      "type": "int"
    },
    "gcLowThreshold": {
      "defaultValue": 80,
      "metadata": {
        "description": "Low Threshold for Image Garbage collection on each node."
      },
      "type": "int"
    },
{{ if not UseManagedIdentity }}
    "servicePrincipalClientId": {
      "metadata": {
        "description": "Client ID (used by cloudprovider)"
      },
      "type": "securestring"
    },
    "servicePrincipalClientSecret": {
      "metadata": {
        "description": "The Service Principal Client Secret."
      },
      "type": "securestring"
    },
{{ end }}
    "masterOffset": {
      "defaultValue": 0,
      "allowedValues": [
        0,
        1,
        2,
        3,
        4
      ],
      "metadata": {
        "description": "The offset into the master pool where to start creating master VMs.  This value can be from 0 to 4, but must be less than masterCount."
      },
      "type": "int"
    },
    "etcdDiskSizeGB": {
      "metadata": {
        "description": "Size in GB to allocate for etcd volume"
      },
      "type": "string"
    },
    "etcdDownloadURLBase": {
      "metadata": {
        "description": "etcd image base URL"
      },
      "type": "string"
    },
    "etcdVersion": {
      "metadata": {
        "description": "etcd version"
      },
      "type": "string"
    },
    "etcdEncryptionKey": {
      "metadata": {
        "description": "Encryption at rest key for etcd"
      },
      "type": "string"
    }
{{if ProvisionJumpbox}}
    ,"jumpboxVMName": {
      "metadata": {
        "description": "jumpbox VM Name"
      },
      "type": "string"
    },
    "jumpboxVMSize": {
      {{GetMasterAllowedSizes}}
      "metadata": {
        "description": "The size of the Virtual Machine. Required"
      },
      "type": "string"
    },
    "jumpboxOSDiskSizeGB": {
      "metadata": {
        "description": "Size in GB to allocate to the private cluster jumpbox VM OS."
      },
      "type": "int"
    },
    "jumpboxPublicKey": {
      "metadata": {
        "description": "SSH public key used for auth to the private cluster jumpbox"
      },
      "type": "string"
    },
    "jumpboxUsername": {
      "metadata": {
        "description": "Username for the private cluster jumpbox"
      },
      "type": "string"
    },
    "jumpboxStorageProfile": {
      "metadata": {
        "description": "Storage Profile for the private cluster jumpbox"
      },
      "type": "string"
    }
{{end}}
{{if HasLinuxProfile}}{{if HasCustomSearchDomain}}
    ,"searchDomainName": {
      "defaultValue": "",
      "metadata": {
        "description": "Custom Search Domain name."
      },
      "type": "string"
    },
    "searchDomainRealmUser": {
      "defaultValue": "",
      "metadata": {
        "description": "Windows server AD user name to join the Linux Machines with active directory and be able to change dns registries."
      },
      "type": "string"
    },
    "searchDomainRealmPassword": {
      "defaultValue": "",
      "metadata": {
        "description": "Windows server AD user password to join the Linux Machines with active directory and be able to change dns registries."
      },
      "type": "securestring"
    }
{{end}}{{end}}
{{if HasLinuxProfile}}{{if HasCustomNodesDNS}}
    ,"dnsServer": {
      "defaultValue": "",
      "metadata": {
        "description": "DNS Server IP"
      },
      "type": "string"
    }
{{end}}{{end}}

{{if EnableEncryptionWithExternalKms}}
   ,
   {{if not UseManagedIdentity}}
   "servicePrincipalObjectId": {
      "metadata": {
        "description": "Object ID (used by cloudprovider)"
      },
      "type": "securestring"
    },
    {{end}}
    "clusterKeyVaultSku": {
       "type": "string",
       "defaultValue": "Standard",
       "allowedValues": [
         "Standard",
         "Premium"
       ],
       "metadata": {
         "description": "SKU for the key vault used by the cluster"
       }
     }
 {{end}}
 {{if IsAzureCNI}}
    ,"AzureCNINetworkMonitorImageURL": {
      "defaultValue": "",
      "metadata": {
        "description": "Azure CNI networkmonitor Image URL"
      },
      "type": "string"
    }
 {{end}}
 {{if .OrchestratorProfile.KubernetesConfig.IsAppGWIngressEnabled}}
    ,"appGwSubnet": {
      "metadata": {
        "description": "Sets the subnet of the Application Gateway"
      },
      "type": "string"
    }
    ,"appGwSku": {
      "metadata": {
        "description": "Sets the subnet of the Application Gateway"
      },
      "type": "string"
    }
 {{end}}
`)

func k8sKubernetesparamsTBytes() ([]byte, error) {
	return _k8sKubernetesparamsT, nil
}

func k8sKubernetesparamsT() (*asset, error) {
	bytes, err := k8sKubernetesparamsTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/kubernetesparams.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sKuberneteswindowsfunctionsPs1 = []byte(`# This is a temporary file to test dot-sourcing functions stored in separate scripts in a zip file

filter Timestamp {"$(Get-Date -Format o): $_"}

function
Write-Log($message)
{
    $msg = $message | Timestamp
    Write-Output $msg
}

function DownloadFileOverHttp
{
    Param(
        [Parameter(Mandatory=$true)][string]
        $Url,
        [Parameter(Mandatory=$true)][string]
        $DestinationPath
    )
    $secureProtocols = @()
    $insecureProtocols = @([System.Net.SecurityProtocolType]::SystemDefault, [System.Net.SecurityProtocolType]::Ssl3)

    foreach ($protocol in [System.Enum]::GetValues([System.Net.SecurityProtocolType]))
    {
        if ($insecureProtocols -notcontains $protocol)
        {
            $secureProtocols += $protocol
        }
    }
    [System.Net.ServicePointManager]::SecurityProtocol = $secureProtocols

    $oldProgressPreference = $ProgressPreference
    $ProgressPreference = 'SilentlyContinue'
    Invoke-WebRequest $Url -UseBasicParsing -OutFile $DestinationPath -Verbose
    $ProgressPreference = $oldProgressPreference
    Write-Log "Downloaded file to $DestinationPath"
}

# https://stackoverflow.com/a/34559554/697126
function New-TemporaryDirectory {
    $parent = [System.IO.Path]::GetTempPath()
    [string] $name = [System.Guid]::NewGuid()
    New-Item -ItemType Directory -Path (Join-Path $parent $name)
}

function Initialize-DataDirectories {
    # Some of the Kubernetes tests that were designed for Linux try to mount /tmp into a pod
    # On Windows, Go translates to c:\tmp. If that path doesn't exist, then some node tests fail

    $requiredPaths = 'c:\tmp'

    $requiredPaths | ForEach-Object {
        if (-Not (Test-Path $_)) {
            New-Item -ItemType Directory -Path $_
        }
    }
}

function Retry-Command
{
    Param(
        [Parameter(Mandatory=$true)][ValidateNotNullOrEmpty()][string]
        $Command,
        [Parameter(Mandatory=$true)][ValidateNotNullOrEmpty()][hashtable]
        $Args,
        [Parameter(Mandatory=$true)][ValidateNotNullOrEmpty()][int]
        $Retries,
        [Parameter(Mandatory=$true)][ValidateNotNullOrEmpty()][int]
        $RetryDelaySeconds
    )

    for ($i = 0; $i -lt $Retries; $i++) {
        try {
            return & $Command @Args
        } catch {
            Start-Sleep $RetryDelaySeconds
        }
    }
}
`)

func k8sKuberneteswindowsfunctionsPs1Bytes() ([]byte, error) {
	return _k8sKuberneteswindowsfunctionsPs1, nil
}

func k8sKuberneteswindowsfunctionsPs1() (*asset, error) {
	bytes, err := k8sKuberneteswindowsfunctionsPs1Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/kuberneteswindowsfunctions.ps1", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sKuberneteswindowssetupPs1 = []byte(`<#
    .SYNOPSIS
        Provisions VM as a Kubernetes agent.

    .DESCRIPTION
        Provisions VM as a Kubernetes agent.

        The parameters passed in are required, and will vary per-deployment.

        Notes on modifying this file:
        - This file extension is PS1, but it is actually used as a template from pkg/engine/template_generator.go
        - All of the lines that have braces in them will be modified. Please do not change them here, change them in the Go sources
        - Single quotes are forbidden, they are reserved to delineate the different members for the ARM template concat() call
#>
[CmdletBinding(DefaultParameterSetName="Standard")]
param(
    [string]
    [ValidateNotNullOrEmpty()]
    $MasterIP,

    [parameter()]
    [ValidateNotNullOrEmpty()]
    $KubeDnsServiceIp,

    [parameter(Mandatory=$true)]
    [ValidateNotNullOrEmpty()]
    $MasterFQDNPrefix,

    [parameter(Mandatory=$true)]
    [ValidateNotNullOrEmpty()]
    $Location,

    [parameter(Mandatory=$true)]
    [ValidateNotNullOrEmpty()]
    $AgentKey,

    [parameter(Mandatory=$true)]
    [ValidateNotNullOrEmpty()]
    $AADClientId,

    [parameter(Mandatory=$true)]
    [ValidateNotNullOrEmpty()]
    $AADClientSecret, # base64

    [parameter(Mandatory=$true)]
    [ValidateNotNullOrEmpty()]
    $NetworkAPIVersion,

    [parameter(Mandatory=$true)]
    [ValidateNotNullOrEmpty()]
    $TargetEnvironment
)



# These globals will not change between nodes in the same cluster, so they are not
# passed as powershell parameters

## SSH public keys to add to authorized_keys
$global:SSHKeys = @( {{ GetSshPublicKeysPowerShell }} )

## Certificates generated by aks-engine
$global:CACertificate = "{{WrapAsParameter "caCertificate"}}"
$global:AgentCertificate = "{{WrapAsParameter "clientCertificate"}}"

## Download sources provided by aks-engine
$global:KubeBinariesPackageSASURL = "{{WrapAsParameter "kubeBinariesSASURL"}}"
$global:WindowsKubeBinariesURL = "{{WrapAsParameter "windowsKubeBinariesURL"}}"
$global:KubeBinariesVersion = "{{WrapAsParameter "kubeBinariesVersion"}}"

## Docker Version
$global:DockerVersion = "{{WrapAsParameter "windowsDockerVersion"}}"

## VM configuration passed by Azure
$global:WindowsTelemetryGUID = "{{WrapAsParameter "windowsTelemetryGUID"}}"
{{if eq GetIdentitySystem "adfs"}}
$global:TenantId = "adfs"
{{else}}
$global:TenantId = "{{WrapAsVariable "tenantID"}}"
{{end}}
$global:SubscriptionId = "{{WrapAsVariable "subscriptionId"}}"
$global:ResourceGroup = "{{WrapAsVariable "resourceGroup"}}"
$global:VmType = "{{WrapAsVariable "vmType"}}"
$global:SubnetName = "{{WrapAsVariable "subnetName"}}"
$global:MasterSubnet = "{{GetWindowsMasterSubnetARMParam}}"
$global:SecurityGroupName = "{{WrapAsVariable "nsgName"}}"
$global:VNetName = "{{WrapAsVariable "virtualNetworkName"}}"
$global:RouteTableName = "{{WrapAsVariable "routeTableName"}}"
$global:PrimaryAvailabilitySetName = "{{WrapAsVariable "primaryAvailabilitySetName"}}"
$global:PrimaryScaleSetName = "{{WrapAsVariable "primaryScaleSetName"}}"

$global:KubeClusterCIDR = "{{WrapAsParameter "kubeClusterCidr"}}"
$global:KubeServiceCIDR = "{{WrapAsParameter "kubeServiceCidr"}}"
$global:VNetCIDR = "{{WrapAsParameter "vnetCidr"}}"
{{if IsKubernetesVersionGe "1.16.0-alpha.1"}}
$global:KubeletNodeLabels = "{{GetAgentKubernetesLabels . "',variables('labelResourceGroup'),'"}}"
{{else}}
$global:KubeletNodeLabels = "{{GetAgentKubernetesLabelsDeprecated . "',variables('labelResourceGroup'),'"}}"
{{end}}
$global:KubeletConfigArgs = @( {{GetKubeletConfigKeyValsPsh .KubernetesConfig }} )

$global:UseManagedIdentityExtension = "{{WrapAsVariable "useManagedIdentityExtension"}}"
$global:UserAssignedClientID = "{{WrapAsVariable "userAssignedClientID"}}"
$global:UseInstanceMetadata = "{{WrapAsVariable "useInstanceMetadata"}}"

$global:LoadBalancerSku = "{{WrapAsVariable "loadBalancerSku"}}"
$global:ExcludeMasterFromStandardLB = "{{WrapAsVariable "excludeMasterFromStandardLB"}}"


# Windows defaults, not changed by aks-engine
$global:KubeDir = "c:\k"
$global:HNSModule = [Io.path]::Combine("$global:KubeDir", "hns.psm1")

$global:KubeDnsSearchPath = "svc.cluster.local"

$global:CNIPath = [Io.path]::Combine("$global:KubeDir", "cni")
$global:NetworkMode = "L2Bridge"
$global:CNIConfig = [Io.path]::Combine($global:CNIPath, "config", "` + "`" + `$global:NetworkMode.conf")
$global:CNIConfigPath = [Io.path]::Combine("$global:CNIPath", "config")


$global:AzureCNIDir = [Io.path]::Combine("$global:KubeDir", "azurecni")
$global:AzureCNIBinDir = [Io.path]::Combine("$global:AzureCNIDir", "bin")
$global:AzureCNIConfDir = [Io.path]::Combine("$global:AzureCNIDir", "netconf")

# Azure cni configuration
# $global:NetworkPolicy = "{{WrapAsParameter "networkPolicy"}}" # BUG: unused
$global:NetworkPlugin = "{{WrapAsParameter "networkPlugin"}}"
$global:VNetCNIPluginsURL = "{{WrapAsParameter "vnetCniWindowsPluginsURL"}}"

# Base64 representation of ZIP archive
$zippedFiles = "{{ GetKubernetesWindowsAgentFunctions }}"

# Extract ZIP from script
[io.file]::WriteAllBytes("scripts.zip", [System.Convert]::FromBase64String($zippedFiles))
Expand-Archive scripts.zip -DestinationPath "C:\\AzureData\\"

# Dot-source contents of zip. This should match the list in template_generator.go GetKubernetesWindowsAgentFunctions
. c:\AzureData\k8s\kuberneteswindowsfunctions.ps1
. c:\AzureData\k8s\windowsconfigfunc.ps1
. c:\AzureData\k8s\windowskubeletfunc.ps1
. c:\AzureData\k8s\windowscnifunc.ps1
. c:\AzureData\k8s\windowsazurecnifunc.ps1
. c:\AzureData\k8s\windowsinstallopensshfunc.ps1

function
Update-ServiceFailureActions()
{
    sc.exe failure "kubelet" actions= restart/60000/restart/60000/restart/60000 reset= 900
    sc.exe failure "kubeproxy" actions= restart/60000/restart/60000/restart/60000 reset= 900
    sc.exe failure "docker" actions= restart/60000/restart/60000/restart/60000 reset= 900
}

try
{
    # Set to false for debugging.  This will output the start script to
    # c:\AzureData\CustomDataSetupScript.log, and then you can RDP
    # to the windows machine, and run the script manually to watch
    # the output.
    if ($true) {
        Write-Log "Provisioning $global:DockerServiceName... with IP $MasterIP"

        Write-Log "Apply telemetry data setting"
        Set-TelemetrySetting -WindowsTelemetryGUID $global:WindowsTelemetryGUID

        Write-Log "Resize os drive if possible"
        Resize-OSDrive

        Write-Log "Initialize data disks"
        Initialize-DataDisks

        Write-Log "Create required data directories as needed"
        Initialize-DataDirectories

        Write-Log "Install docker"
        Install-Docker -DockerVersion $global:DockerVersion

        Write-Log "Download kubelet binaries and unzip"
        Get-KubePackage -KubeBinariesSASURL $global:KubeBinariesPackageSASURL

        # this overwrite the binaries that are download from the custom packge with binaries
        # The custom package has a few files that are nessary for future steps (nssm.exe)
        # this is a temporary work around to get the binaries until we depreciate
        # custom package and nssm.exe as defined in #3851.
        if ($global:WindowsKubeBinariesURL){
            Write-Log "Overwriting kube node binaries from $global:WindowsKubeBinariesURL"
            Get-KubeBinaries -KubeBinariesURL $global:WindowsKubeBinariesURL
        }


        Write-Log "Write Azure cloud provider config"
        Write-AzureConfig ` + "`" + `
            -KubeDir $global:KubeDir ` + "`" + `
            -AADClientId $AADClientId ` + "`" + `
            -AADClientSecret $([System.Text.Encoding]::ASCII.GetString([System.Convert]::FromBase64String($AADClientSecret))) ` + "`" + `
            -TenantId $global:TenantId ` + "`" + `
            -SubscriptionId $global:SubscriptionId ` + "`" + `
            -ResourceGroup $global:ResourceGroup ` + "`" + `
            -Location $Location ` + "`" + `
            -VmType $global:VmType ` + "`" + `
            -SubnetName $global:SubnetName ` + "`" + `
            -SecurityGroupName $global:SecurityGroupName ` + "`" + `
            -VNetName $global:VNetName ` + "`" + `
            -RouteTableName $global:RouteTableName ` + "`" + `
            -PrimaryAvailabilitySetName $global:PrimaryAvailabilitySetName ` + "`" + `
            -PrimaryScaleSetName $global:PrimaryScaleSetName ` + "`" + `
            -UseManagedIdentityExtension $global:UseManagedIdentityExtension ` + "`" + `
            -UserAssignedClientID $global:UserAssignedClientID ` + "`" + `
            -UseInstanceMetadata $global:UseInstanceMetadata ` + "`" + `
            -LoadBalancerSku $global:LoadBalancerSku ` + "`" + `
            -ExcludeMasterFromStandardLB $global:ExcludeMasterFromStandardLB ` + "`" + `
            -TargetEnvironment $TargetEnvironment

        {{if IsAzureStackCloud}}
        $azureStackConfigFile = [io.path]::Combine($global:KubeDir, "azurestackcloud.json")
        $envJSON = "{{ GetBase64EncodedEnvironmentJSON }}"
        [io.file]::WriteAllBytes($azureStackConfigFile, [System.Convert]::FromBase64String($envJSON))
        {{end}}

        Write-Log "Write ca root"
        Write-CACert -CACertificate $global:CACertificate ` + "`" + `
                     -KubeDir $global:KubeDir

        Write-Log "Write kube config"
        Write-KubeConfig -CACertificate $global:CACertificate ` + "`" + `
                         -KubeDir $global:KubeDir ` + "`" + `
                         -MasterFQDNPrefix $MasterFQDNPrefix ` + "`" + `
                         -MasterIP $MasterIP ` + "`" + `
                         -AgentKey $AgentKey ` + "`" + `
                         -AgentCertificate $global:AgentCertificate


        Write-Log "Create the Pause Container kubletwin/pause"
        New-InfraContainer -KubeDir $global:KubeDir

        Write-Log "Configuring networking with NetworkPlugin:$global:NetworkPlugin"

        # Configure network policy.
        if ($global:NetworkPlugin -eq "azure") {
            Install-VnetPlugins -AzureCNIConfDir $global:AzureCNIConfDir ` + "`" + `
                                -AzureCNIBinDir $global:AzureCNIBinDir ` + "`" + `
                                -VNetCNIPluginsURL $global:VNetCNIPluginsURL
            Set-AzureCNIConfig -AzureCNIConfDir $global:AzureCNIConfDir ` + "`" + `
                               -KubeDnsSearchPath $global:KubeDnsSearchPath ` + "`" + `
                               -KubeClusterCIDR $global:KubeClusterCIDR ` + "`" + `
                               -MasterSubnet $global:MasterSubnet ` + "`" + `
                               -KubeServiceCIDR $global:KubeServiceCIDR ` + "`" + `
                               -VNetCIDR $global:VNetCIDR ` + "`" + `
                               -TargetEnvironment $TargetEnvironment

            if ($TargetEnvironment -ieq "AzureStackCloud") {
                GenerateAzureStackCNIConfig ` + "`" + `
                    -TenantId $global:TenantId ` + "`" + `
                    -SubscriptionId $global:SubscriptionId ` + "`" + `
                    -ResourceGroup $global:ResourceGroup ` + "`" + `
                    -AADClientId $AADClientId ` + "`" + `
                    -AADClientSecret $([System.Text.Encoding]::ASCII.GetString([System.Convert]::FromBase64String($AADClientSecret))) ` + "`" + `
                    -NetworkAPIVersion $NetworkAPIVersion ` + "`" + `
                    -AzureEnvironmentFilePath $([io.path]::Combine($global:KubeDir, "azurestackcloud.json")) ` + "`" + `
                    -IdentitySystem "{{ GetIdentitySystem }}"
            }

        } elseif ($global:NetworkPlugin -eq "kubenet") {
            Update-WinCNI -CNIPath $global:CNIPath
            Get-HnsPsm1 -HNSModule $global:HNSModule
        }

        Write-Log "Write kubelet startfile with pod CIDR of $podCIDR"
        Install-KubernetesServices ` + "`" + `
            -KubeletConfigArgs $global:KubeletConfigArgs ` + "`" + `
            -KubeBinariesVersion $global:KubeBinariesVersion ` + "`" + `
            -NetworkPlugin $global:NetworkPlugin ` + "`" + `
            -NetworkMode $global:NetworkMode ` + "`" + `
            -KubeDir $global:KubeDir ` + "`" + `
            -AzureCNIBinDir $global:AzureCNIBinDir ` + "`" + `
            -AzureCNIConfDir $global:AzureCNIConfDir ` + "`" + `
            -CNIPath $global:CNIPath ` + "`" + `
            -CNIConfig $global:CNIConfig ` + "`" + `
            -CNIConfigPath $global:CNIConfigPath ` + "`" + `
            -MasterIP $MasterIP ` + "`" + `
            -KubeDnsServiceIp $KubeDnsServiceIp ` + "`" + `
            -MasterSubnet $global:MasterSubnet ` + "`" + `
            -KubeClusterCIDR $global:KubeClusterCIDR ` + "`" + `
            -KubeServiceCIDR $global:KubeServiceCIDR ` + "`" + `
            -HNSModule $global:HNSModule ` + "`" + `
            -KubeletNodeLabels $global:KubeletNodeLabels

        # Install OpenSSH if SSH enabled
        $sshEnabled = [System.Convert]::ToBoolean("{{ WindowsSSHEnabled }}")

        if ( $sshEnabled ) {
            Install-OpenSSH -SSHKeys $SSHKeys
        }

        Write-Log "Disable Internet Explorer compat mode and set homepage"
        Set-Explorer

        Write-Log "Adjust pagefile size"
        Adjust-PageFileSize

        Write-Log "Start preProvisioning script"
        PREPROVISION_EXTENSION

        Write-Log "Update service failure actions"
        Update-ServiceFailureActions

        Write-Log "Setup Complete, reboot computer"
        Restart-Computer
    }
    else
    {
        # keep for debugging purposes
        Write-Log ".\CustomDataSetupScript.ps1 -MasterIP $MasterIP -KubeDnsServiceIp $KubeDnsServiceIp -MasterFQDNPrefix $MasterFQDNPrefix -Location $Location -AgentKey $AgentKey -AADClientId $AADClientId -AADClientSecret $AADClientSecret"
    }
}
catch
{
    Write-Error $_
}
`)

func k8sKuberneteswindowssetupPs1Bytes() ([]byte, error) {
	return _k8sKuberneteswindowssetupPs1, nil
}

func k8sKuberneteswindowssetupPs1() (*asset, error) {
	bytes, err := k8sKuberneteswindowssetupPs1Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/kuberneteswindowssetup.ps1", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sManifestsKubernetesmasterCloudControllerManagerYaml = []byte(`apiVersion: v1
kind: Pod
metadata:
  name: cloud-controller-manager
  namespace: kube-system
  labels:
    tier: control-plane
    component: cloud-controller-manager
spec:
  hostNetwork: true
  containers:
    - name: cloud-controller-manager
      image: <img>
      imagePullPolicy: IfNotPresent
      command: ["cloud-controller-manager"]
      args: [<config>]
      volumeMounts:
        - name: etc-kubernetes
          mountPath: /etc/kubernetes
        - name: etc-ssl
          mountPath: /etc/ssl
          readOnly: true
        - name: var-lib-kubelet
          mountPath: /var/lib/kubelet
        - name: msi
          mountPath: /var/lib/waagent/ManagedIdentity-Settings
          readOnly: true
  volumes:
    - name: etc-kubernetes
      hostPath:
        path: /etc/kubernetes
    - name: etc-ssl
      hostPath:
        path: /etc/ssl
    - name: var-lib-kubelet
      hostPath:
        path: /var/lib/kubelet
    - name: msi
      hostPath:
        path: /var/lib/waagent/ManagedIdentity-Settings
`)

func k8sManifestsKubernetesmasterCloudControllerManagerYamlBytes() ([]byte, error) {
	return _k8sManifestsKubernetesmasterCloudControllerManagerYaml, nil
}

func k8sManifestsKubernetesmasterCloudControllerManagerYaml() (*asset, error) {
	bytes, err := k8sManifestsKubernetesmasterCloudControllerManagerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/manifests/kubernetesmaster-cloud-controller-manager.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sManifestsKubernetesmasterKubeAddonManagerYaml = []byte(`apiVersion: v1
kind: Pod
metadata:
  name: kube-addon-manager
  namespace: kube-system
  version: v1
spec:
  hostNetwork: true
  containers:
  - name: kube-addon-manager
    image: <img>
    imagePullPolicy: IfNotPresent
    resources:
      requests:
        cpu: 5m
        memory: 50Mi
    volumeMounts:
    - name: addons
      mountPath: /etc/kubernetes/addons
      readOnly: true
    - name: msi
      mountPath: /var/lib/waagent/ManagedIdentity-Settings
      readOnly: true
  volumes:
  - name: addons
    hostPath:
      path: /etc/kubernetes/addons
  - name: msi
    hostPath:
      path: /var/lib/waagent/ManagedIdentity-Settings
`)

func k8sManifestsKubernetesmasterKubeAddonManagerYamlBytes() ([]byte, error) {
	return _k8sManifestsKubernetesmasterKubeAddonManagerYaml, nil
}

func k8sManifestsKubernetesmasterKubeAddonManagerYaml() (*asset, error) {
	bytes, err := k8sManifestsKubernetesmasterKubeAddonManagerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/manifests/kubernetesmaster-kube-addon-manager.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sManifestsKubernetesmasterKubeApiserverYaml = []byte(`apiVersion: v1
kind: Pod
metadata:
  name: kube-apiserver
  namespace: kube-system
  labels:
    tier: control-plane
    component: kube-apiserver
spec:
  hostNetwork: true
  containers:
    - name: kube-apiserver
      image: <img>
      imagePullPolicy: IfNotPresent
      command: ["/hyperkube", "kube-apiserver"]
      args: [<args>]
      volumeMounts:
        - name: etc-kubernetes
          mountPath: /etc/kubernetes
        - name: var-lib-kubelet
          mountPath: /var/lib/kubelet
        - name: msi
          mountPath: /var/lib/waagent/ManagedIdentity-Settings
          readOnly: true
        - name: sock
          mountPath: /opt
        - name: auditlog
          mountPath: /var/log/kubeaudit
  volumes:
    - name: etc-kubernetes
      hostPath:
        path: /etc/kubernetes
    - name: var-lib-kubelet
      hostPath:
        path: /var/lib/kubelet
    - name: msi
      hostPath:
        path: /var/lib/waagent/ManagedIdentity-Settings
    - name: sock
      hostPath:
        path: /opt
    - name: auditlog
      hostPath:
        path: /var/log/kubeaudit
`)

func k8sManifestsKubernetesmasterKubeApiserverYamlBytes() ([]byte, error) {
	return _k8sManifestsKubernetesmasterKubeApiserverYaml, nil
}

func k8sManifestsKubernetesmasterKubeApiserverYaml() (*asset, error) {
	bytes, err := k8sManifestsKubernetesmasterKubeApiserverYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/manifests/kubernetesmaster-kube-apiserver.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sManifestsKubernetesmasterKubeControllerManagerCustomYaml = []byte(`apiVersion: v1
kind: Pod
metadata:
  name: kube-controller-manager
  namespace: kube-system
  labels:
    tier: control-plane
    component: kube-controller-manager
spec:
  hostNetwork: true
  containers:
    - name: kube-controller-manager
      image: <img>
      imagePullPolicy: IfNotPresent
      command: ["/hyperkube", "kube-controller-manager"]
      args: [<args>]
      env:
      - name: AZURE_ENVIRONMENT_FILEPATH
        value: "/etc/kubernetes/azurestackcloud.json"
      volumeMounts:
        - name: etc-kubernetes
          mountPath: /etc/kubernetes
        - name: var-lib-kubelet
          mountPath: /var/lib/kubelet
        - name: msi
          mountPath: /var/lib/waagent/ManagedIdentity-Settings
          readOnly: true
        <volumeMountssl>
  volumes:
    - name: etc-kubernetes
      hostPath:
        path: /etc/kubernetes
    - name: var-lib-kubelet
      hostPath:
        path: /var/lib/kubelet
    - name: msi
      hostPath:
        path: /var/lib/waagent/ManagedIdentity-Settings
    <volumessl>`)

func k8sManifestsKubernetesmasterKubeControllerManagerCustomYamlBytes() ([]byte, error) {
	return _k8sManifestsKubernetesmasterKubeControllerManagerCustomYaml, nil
}

func k8sManifestsKubernetesmasterKubeControllerManagerCustomYaml() (*asset, error) {
	bytes, err := k8sManifestsKubernetesmasterKubeControllerManagerCustomYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/manifests/kubernetesmaster-kube-controller-manager-custom.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sManifestsKubernetesmasterKubeControllerManagerYaml = []byte(`apiVersion: v1
kind: Pod
metadata:
  name: kube-controller-manager
  namespace: kube-system
  labels:
    tier: control-plane
    component: kube-controller-manager
spec:
  hostNetwork: true
  containers:
    - name: kube-controller-manager
      image: <img>
      imagePullPolicy: IfNotPresent
      command: ["/hyperkube", "kube-controller-manager"]
      args: [<args>]
      volumeMounts:
        - name: etc-kubernetes
          mountPath: /etc/kubernetes
        - name: var-lib-kubelet
          mountPath: /var/lib/kubelet
        - name: msi
          mountPath: /var/lib/waagent/ManagedIdentity-Settings
          readOnly: true
  volumes:
    - name: etc-kubernetes
      hostPath:
        path: /etc/kubernetes
    - name: var-lib-kubelet
      hostPath:
        path: /var/lib/kubelet
    - name: msi
      hostPath:
        path: /var/lib/waagent/ManagedIdentity-Settings
`)

func k8sManifestsKubernetesmasterKubeControllerManagerYamlBytes() ([]byte, error) {
	return _k8sManifestsKubernetesmasterKubeControllerManagerYaml, nil
}

func k8sManifestsKubernetesmasterKubeControllerManagerYaml() (*asset, error) {
	bytes, err := k8sManifestsKubernetesmasterKubeControllerManagerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/manifests/kubernetesmaster-kube-controller-manager.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sManifestsKubernetesmasterKubeSchedulerYaml = []byte(`apiVersion: v1
kind: Pod
metadata:
  name: kube-scheduler
  namespace: kube-system
  labels:
    tier: control-plane
    component: kube-scheduler
spec:
  hostNetwork: true
  containers:
    - name: kube-scheduler
      image: <img>
      imagePullPolicy: IfNotPresent
      command: ["/hyperkube", "kube-scheduler"]
      args: [<args>]
      volumeMounts:
        - name: etc-kubernetes
          mountPath: /etc/kubernetes
        - name: var-lib-kubelet
          mountPath: /var/lib/kubelet
        - name: msi
          mountPath: /var/lib/waagent/ManagedIdentity-Settings
          readOnly: true
  volumes:
    - name: etc-kubernetes
      hostPath:
        path: /etc/kubernetes
    - name: var-lib-kubelet
      hostPath:
        path: /var/lib/kubelet
    - name: msi
      hostPath:
        path: /var/lib/waagent/ManagedIdentity-Settings
`)

func k8sManifestsKubernetesmasterKubeSchedulerYamlBytes() ([]byte, error) {
	return _k8sManifestsKubernetesmasterKubeSchedulerYaml, nil
}

func k8sManifestsKubernetesmasterKubeSchedulerYaml() (*asset, error) {
	bytes, err := k8sManifestsKubernetesmasterKubeSchedulerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/manifests/kubernetesmaster-kube-scheduler.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sWindowsazurecnifuncPs1 = []byte(`

# TODO: remove - dead code?
function
Set-VnetPluginMode()
{
    Param(
        [Parameter(Mandatory=$true)][string]
        $AzureCNIConfDir,
        [Parameter(Mandatory=$true)][string]
        $Mode
    )
    # Sets Azure VNET CNI plugin operational mode.
    $fileName  = [Io.path]::Combine("$AzureCNIConfDir", "10-azure.conflist")
    (Get-Content $fileName) | %{$_ -replace "` + "`" + `"mode` + "`" + `":.*", "` + "`" + `"mode` + "`" + `": ` + "`" + `"$Mode` + "`" + `","} | Out-File -encoding ASCII -filepath $fileName
}


function
Install-VnetPlugins
{
    Param(
        [Parameter(Mandatory=$true)][string]
        $AzureCNIConfDir,
        [Parameter(Mandatory=$true)][string]
        $AzureCNIBinDir,
        [Parameter(Mandatory=$true)][string]
        $VNetCNIPluginsURL
    )
    # Create CNI directories.
    mkdir $AzureCNIBinDir
    mkdir $AzureCNIConfDir

    # Download Azure VNET CNI plugins.
    # Mirror from https://github.com/Azure/azure-container-networking/releases
    $zipfile =  [Io.path]::Combine("$AzureCNIDir", "azure-vnet.zip")
    DownloadFileOverHttp -Url $VNetCNIPluginsURL -DestinationPath $zipfile
    Expand-Archive -path $zipfile -DestinationPath $AzureCNIBinDir
    del $zipfile

    # Windows does not need a separate CNI loopback plugin because the Windows
    # kernel automatically creates a loopback interface for each network namespace.
    # Copy CNI network config file and set bridge mode.
    move $AzureCNIBinDir/*.conflist $AzureCNIConfDir
}

# TODO: remove - dead code?
function
Set-AzureNetworkPlugin()
{
    # Azure VNET network policy requires tunnel (hairpin) mode because policy is enforced in the host.
    Set-VnetPluginMode "tunnel"
}

function
Set-AzureCNIConfig
{
    Param(
        [Parameter(Mandatory=$true)][string]
        $AzureCNIConfDir,
        [Parameter(Mandatory=$true)][string]
        $KubeDnsSearchPath,
        [Parameter(Mandatory=$true)][string]
        $KubeClusterCIDR,
        [Parameter(Mandatory=$true)][string]
        $MasterSubnet,
        [Parameter(Mandatory=$true)][string]
        $KubeServiceCIDR,
        [Parameter(Mandatory=$true)][string]
        $VNetCIDR,
        [Parameter(Mandatory=$true)][string]
        $TargetEnvironment
    )
    # Fill in DNS information for kubernetes.
    $fileName  = [Io.path]::Combine("$AzureCNIConfDir", "10-azure.conflist")
    $configJson = Get-Content $fileName | ConvertFrom-Json
    $configJson.plugins.dns.Nameservers[0] = $KubeDnsServiceIp
    $configJson.plugins.dns.Search[0] = $KubeDnsSearchPath
    $configJson.plugins.AdditionalArgs[0].Value.ExceptionList[0] = $KubeClusterCIDR
    $configJson.plugins.AdditionalArgs[0].Value.ExceptionList[1] = $MasterSubnet
    $configJson.plugins.AdditionalArgs[1].Value.DestinationPrefix  = $KubeServiceCIDR
    $configJson.plugins.AdditionalArgs[0].Value.ExceptionList += $VNetCIDR

    if ($TargetEnvironment -ieq "AzureStackCloud") {
        Add-Member -InputObject $configJson.plugins[0].ipam -MemberType NoteProperty -Name "environment" -Value "mas"
    }

    $configJson | ConvertTo-Json -depth 20 | Out-File -encoding ASCII -filepath $fileName
}

function GetSubnetPrefix
{
    Param(
        [Parameter(Mandatory=$true)][ValidateNotNullOrEmpty()][string] $Token,
        [Parameter(Mandatory=$true)][ValidateNotNullOrEmpty()][string] $SubnetId,
        [Parameter(Mandatory=$true)][ValidateNotNullOrEmpty()][string] $ResourceManagerEndpoint,
        [Parameter(Mandatory=$true)][ValidateNotNullOrEmpty()][string] $NetworkAPIVersion
    )

    $uri = "$($ResourceManagerEndpoint)$($SubnetId)?api-version=$NetworkAPIVersion"
    $headers = @{Authorization="Bearer $Token"}

    $response = Retry-Command -Command "Invoke-RestMethod" -Args @{Uri=$uri; Method="Get"; ContentType="application/json"; Headers=$headers} -Retries 5 -RetryDelaySeconds 10

    if(!$response) {
        throw 'Error getting subnet prefix'
    }

    $response.properties.addressPrefix
}

function GenerateAzureStackCNIConfig
{
    Param(
        [Parameter(Mandatory=$true)][ValidateNotNullOrEmpty()][string] $TenantId,
        [Parameter(Mandatory=$true)][ValidateNotNullOrEmpty()][string] $SubscriptionId,
        [Parameter(Mandatory=$true)][ValidateNotNullOrEmpty()][string] $AADClientId,
        [Parameter(Mandatory=$true)][ValidateNotNullOrEmpty()][string] $AADClientSecret,
        [Parameter(Mandatory=$true)][ValidateNotNullOrEmpty()][string] $ResourceGroup,
        [Parameter(Mandatory=$true)][ValidateNotNullOrEmpty()][string] $NetworkAPIVersion,
        [Parameter(Mandatory=$true)][ValidateNotNullOrEmpty()][string] $AzureEnvironmentFilePath,
        [Parameter(Mandatory=$true)][ValidateNotNullOrEmpty()][string] $IdentitySystem
    )

    $networkInterfacesFile = "C:\k\network-interfaces.json"
    $azureCNIConfigFile = "C:\k\interfaces.json"
    $azureEnvironment = Get-Content $AzureEnvironmentFilePath | ConvertFrom-Json

    Write-Log "------------------------------------------------------------------------"
    Write-Log "Parameters"
    Write-Log "------------------------------------------------------------------------"
    Write-Log "TenantId:                  $TenantId"
    Write-Log "SubscriptionId:            $SubscriptionId"
    Write-Log "AADClientId:               ..."
    Write-Log "AADClientSecret:           ..."
    Write-Log "ResourceGroup:             $ResourceGroup"
    Write-Log "NetworkAPIVersion:         $NetworkAPIVersion"
    Write-Log "ServiceManagementEndpoint: $($azureEnvironment.serviceManagementEndpoint)"
    Write-Log "ActiveDirectoryEndpoint:   $($azureEnvironment.activeDirectoryEndpoint)"
    Write-Log "ResourceManagerEndpoint:   $($azureEnvironment.resourceManagerEndpoint)"
    Write-Log "------------------------------------------------------------------------"
    Write-Log "Variables"
    Write-Log "------------------------------------------------------------------------"
    Write-Log "azureCNIConfigFile: $azureCNIConfigFile"
    Write-Log "networkInterfacesFile: $networkInterfacesFile"
    Write-Log "------------------------------------------------------------------------"

    Write-Log "Generating token for Azure Resource Manager"

    $tokenURL = ""
    if($IdentitySystem -ieq "adfs") {
        $tokenURL = "$($azureEnvironment.activeDirectoryEndpoint)adfs/oauth2/token"
    } else {
        $tokenURL = "$($azureEnvironment.activeDirectoryEndpoint)$TenantId/oauth2/token"
    }

    Add-Type -AssemblyName System.Web
    $encodedSecret = [System.Web.HttpUtility]::UrlEncode($AADClientSecret)

    $body = "grant_type=client_credentials&client_id=$AADClientId&client_secret=$encodedSecret&resource=$($azureEnvironment.serviceManagementEndpoint)"
    $args = @{Uri=$tokenURL; Method="Post"; Body=$body; ContentType='application/x-www-form-urlencoded'}
    $tokenResponse = Retry-Command -Command "Invoke-RestMethod" -Args $args -Retries 5 -RetryDelaySeconds 10

    if(!$tokenResponse) {
        throw 'Error generating token for Azure Resource Manager'
    }

    $token = $tokenResponse | Select-Object -ExpandProperty access_token

    Write-Log "Fetching network interface configuration for node"

    $interfacesUri = "$($azureEnvironment.resourceManagerEndpoint)subscriptions/$SubscriptionId/resourceGroups/$ResourceGroup/providers/Microsoft.Network/networkInterfaces?api-version=$NetworkAPIVersion"
    $headers = @{Authorization="Bearer $token"}
    $args = @{Uri=$interfacesUri; Method="Get"; ContentType="application/json"; Headers=$headers; OutFile=$networkInterfacesFile}
    Retry-Command -Command "Invoke-RestMethod" -Args $args -Retries 5 -RetryDelaySeconds 10

    if(!$(Test-Path $networkInterfacesFile)) {
        throw 'Error fetching network interface configuration for node'
    }

    Write-Log "Generating Azure CNI interface file"

    $localNics = Get-NetAdapter | Select-Object -ExpandProperty MacAddress | ForEach-Object {$_ -replace "-",""}

    $sdnNics = Get-Content $networkInterfacesFile ` + "`" + `
        | ConvertFrom-Json ` + "`" + `
        | Select-Object -ExpandProperty value ` + "`" + `
        | Where-Object { $localNics.Contains($_.properties.macAddress) } ` + "`" + `
        | Where-Object { $_.properties.ipConfigurations.Count -gt 0}

    $interfaces = @{
        Interfaces = @( $sdnNics | ForEach-Object { @{
            MacAddress = $_.properties.macAddress
            IsPrimary = $_.properties.primary
            IPSubnets = @(@{
                Prefix = GetSubnetPrefix ` + "`" + `
                            -Token $token ` + "`" + `
                            -SubnetId $_.properties.ipConfigurations[0].properties.subnet.id ` + "`" + `
                            -NetworkAPIVersion $NetworkAPIVersion ` + "`" + `
                            -ResourceManagerEndpoint $($azureEnvironment.resourceManagerEndpoint)
                IPAddresses = $_.properties.ipConfigurations | ForEach-Object { @{
                    Address = $_.properties.privateIPAddress
                    IsPrimary = $_.properties.primary
                }}
            })
        }})
    }

    ConvertTo-Json $interfaces -Depth 6 | Out-File -FilePath $azureCNIConfigFile -Encoding ascii

    Set-ItemProperty -Path $azureCNIConfigFile -Name IsReadOnly -Value $true
}

`)

func k8sWindowsazurecnifuncPs1Bytes() ([]byte, error) {
	return _k8sWindowsazurecnifuncPs1, nil
}

func k8sWindowsazurecnifuncPs1() (*asset, error) {
	bytes, err := k8sWindowsazurecnifuncPs1Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/windowsazurecnifunc.ps1", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sWindowscnifuncPs1 = []byte(`function Get-HnsPsm1
{
    Param(
        [string]
        $HnsUrl = "https://github.com/Microsoft/SDN/raw/master/Kubernetes/windows/hns.psm1",
        [Parameter(Mandatory=$true)][string]
        $HNSModule
    )
    DownloadFileOverHttp -Url $HnsUrl -DestinationPath "$HNSModule"
}

function Update-WinCNI
{
    Param(
        [string]
        $WinCniUrl = "https://github.com/Microsoft/SDN/raw/master/Kubernetes/flannel/l2bridge/cni/win-bridge.exe",
        [Parameter(Mandatory=$true)][string]
        $CNIPath
    )
    $wincni = "win-bridge.exe"
    $wincniFile = [Io.path]::Combine($CNIPath, $wincni)
    DownloadFileOverHttp -Url $WinCniUrl -DestinationPath $wincniFile
}

# TODO: Move the code that creates the wincni configuration file out of windowskubeletfunc.ps1 and put it here`)

func k8sWindowscnifuncPs1Bytes() ([]byte, error) {
	return _k8sWindowscnifuncPs1, nil
}

func k8sWindowscnifuncPs1() (*asset, error) {
	bytes, err := k8sWindowscnifuncPs1Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/windowscnifunc.ps1", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sWindowsconfigfuncPs1 = []byte(`

# Set the service telemetry GUID. This is used with Windows Analytics https://docs.microsoft.com/en-us/sccm/core/clients/manage/monitor-windows-analytics
function Set-TelemetrySetting
{
    Param(
        [Parameter(Mandatory=$true)][string]
        $WindowsTelemetryGUID
    )
    Set-ItemProperty -Path "HKLM:\Software\Microsoft\Windows\CurrentVersion\Policies\DataCollection" -Name "CommercialId" -Value $WindowsTelemetryGUID -Force
}

# Resize the system partition to the max available size. Azure can resize a managed disk, but the VM still needs to extend the partition boundary
function Resize-OSDrive
{
    $osDrive = ((Get-WmiObject Win32_OperatingSystem).SystemDrive).TrimEnd(":")
    $size = (Get-Partition -DriveLetter $osDrive).Size
    $maxSize = (Get-PartitionSupportedSize -DriveLetter $osDrive).SizeMax
    if ($size -lt $maxSize)
    {
        Resize-Partition -DriveLetter $osDrive -Size $maxSize
    }
}

# https://docs.microsoft.com/en-us/powershell/module/storage/new-partition
function Initialize-DataDisks
{
    Get-Disk | Where-Object PartitionStyle -eq 'raw' | Initialize-Disk -PartitionStyle MBR -PassThru | New-Partition -UseMaximumSize -AssignDriveLetter | Format-Volume -FileSystem NTFS -Force
}

# Set the Internet Explorer to use the latest rendering mode on all sites
# https://docs.microsoft.com/en-us/windows-hardware/customize/desktop/unattend/microsoft-windows-ie-internetexplorer-intranetcompatibilitymode
# (This only affects installations with UI)
function Set-Explorer
{
    New-Item -Path HKLM:"\\SOFTWARE\\Policies\\Microsoft\\Internet Explorer"
    New-Item -Path HKLM:"\\SOFTWARE\\Policies\\Microsoft\\Internet Explorer\\BrowserEmulation"
    New-ItemProperty -Path HKLM:"\\SOFTWARE\\Policies\\Microsoft\\Internet Explorer\\BrowserEmulation" -Name IntranetCompatibilityMode -Value 0 -Type DWord
    New-Item -Path HKLM:"\\SOFTWARE\\Policies\\Microsoft\\Internet Explorer\\Main"
    New-ItemProperty -Path HKLM:"\\SOFTWARE\\Policies\\Microsoft\\Internet Explorer\\Main" -Name "Start Page" -Type String -Value http://bing.com
}

function Install-Docker
{
    Param(
        [Parameter(Mandatory=$true)][string]
        $DockerVersion
    )

    # Note: to get a list of all versions, use this snippet
    # $versions = (curl.exe -L "https://go.microsoft.com/fwlink/?LinkID=825636&clcid=0x409" | ConvertFrom-Json).Versions | Get-Member -Type NoteProperty | Select-Object Name
    # Docker version to API version decoder: https://docs.docker.com/develop/sdk/#api-version-matrix

    switch ($DockerVersion.Substring(0,5))
    {
        "17.06" {
            Write-Log "Docker 17.06 found, setting DOCKER_API_VERSION to 1.30"
            [System.Environment]::SetEnvironmentVariable('DOCKER_API_VERSION', '1.30', [System.EnvironmentVariableTarget]::Machine)
        }

        "18.03" {
            Write-Log "Docker 18.03 found, setting DOCKER_API_VERSION to 1.37"
            [System.Environment]::SetEnvironmentVariable('DOCKER_API_VERSION', '1.37', [System.EnvironmentVariableTarget]::Machine)
        }

        default {
            Write-Log "Docker version $DockerVersion found, clearing DOCKER_API_VERSION"
            [System.Environment]::SetEnvironmentVariable('DOCKER_API_VERSION', $null, [System.EnvironmentVariableTarget]::Machine)
        }
    }

    try {
        Find-Package -Name Docker -ProviderName DockerMsftProvider -RequiredVersion $DockerVersion -ErrorAction Stop
        Write-Log "Found version $DockerVersion. Installing..."
        Install-Package -Name Docker -ProviderName DockerMsftProvider -Update -Force -RequiredVersion $DockerVersion
        net start docker
        Write-Log "Installed version $DockerVersion"
    } catch {
        Write-Log "Error while installing package: $_.Exception.Message"
        $currentDockerVersion = (Get-Package -Name Docker -ProviderName DockerMsftProvider).Version
        Write-Log "Not able to install docker version. Using default version $currentDockerVersion"
    }
}

# Pagefile adjustments
function Adjust-PageFileSize()
{
    wmic pagefileset set InitialSize=8096,MaximumSize=8096
}
`)

func k8sWindowsconfigfuncPs1Bytes() ([]byte, error) {
	return _k8sWindowsconfigfuncPs1, nil
}

func k8sWindowsconfigfuncPs1() (*asset, error) {
	bytes, err := k8sWindowsconfigfuncPs1Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/windowsconfigfunc.ps1", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sWindowsinstallopensshfuncPs1 = []byte(`function
Install-OpenSSH {
    Param(
        [Parameter(Mandatory = $true)][string[]] 
        $SSHKeys
    )

    $adminpath = "c:\ProgramData\ssh"
    $adminfile = "administrators_authorized_keys"

    Write-Host "Installing OpenSSH"
    $isAvailable = Get-WindowsCapability -Online | ? Name -like 'OpenSSH*'

    if (!$isAvailable) {
        Write-Error "OpenSSH is not avaliable on this machine"
        exit 1
    }

    Add-WindowsCapability -Online -Name OpenSSH.Server~~~~0.0.1.0

    Start-Service sshd

    if (!(Test-Path "$adminpath")) {
        Write-Host "Created new file and text content added"
        New-Item -path $adminpath -name $adminfile -type "file" -value ""
    }

    Write-Host "$adminpath found."
    Write-Host "Adding keys to: $adminpath\$adminfile ..."
    $SSHKeys | foreach-object {
        Add-Content $adminpath\$adminfile $_
    }

    Write-Host "Setting required permissions..."
    icacls $adminpath\$adminfile /remove "NT AUTHORITY\Authenticated Users"
    icacls $adminpath\$adminfile /inheritance:r
    icacls $adminpath\$adminfile /grant SYSTEM:` + "`" + `(F` + "`" + `)
    icacls $adminpath\$adminfile /grant BUILTIN\Administrators:` + "`" + `(F` + "`" + `)

    Write-Host "Restarting sshd service..."
    Restart-Service sshd
    # OPTIONAL but recommended:
    Set-Service -Name sshd -StartupType 'Automatic'

    # Confirm the Firewall rule is configured. It should be created automatically by setup. 
    $firewall = Get-NetFirewallRule -Name *ssh*

    if (!$firewall) {
        Write-Error "OpenSSH is firewall is not configured properly"
        exit 1
    }
    Write-Host "OpenSSH installed and configured successfully"
}
`)

func k8sWindowsinstallopensshfuncPs1Bytes() ([]byte, error) {
	return _k8sWindowsinstallopensshfuncPs1, nil
}

func k8sWindowsinstallopensshfuncPs1() (*asset, error) {
	bytes, err := k8sWindowsinstallopensshfuncPs1Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/windowsinstallopensshfunc.ps1", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sWindowskubeletfuncPs1 = []byte(`function
Write-AzureConfig {
    Param(

        [Parameter(Mandatory = $true)][string]
        $AADClientId,
        [Parameter(Mandatory = $true)][string]
        $AADClientSecret,
        [Parameter(Mandatory = $true)][string]
        $TenantId,
        [Parameter(Mandatory = $true)][string]
        $SubscriptionId,
        [Parameter(Mandatory = $true)][string]
        $ResourceGroup,
        [Parameter(Mandatory = $true)][string]
        $Location,
        [Parameter(Mandatory = $true)][string]
        $VmType,
        [Parameter(Mandatory = $true)][string]
        $SubnetName,
        [Parameter(Mandatory = $true)][string]
        $SecurityGroupName,
        [Parameter(Mandatory = $true)][string]
        $VNetName,
        [Parameter(Mandatory = $true)][string]
        $RouteTableName,
        [Parameter(Mandatory = $false)][string] # Need one of these configured
        $PrimaryAvailabilitySetName,
        [Parameter(Mandatory = $false)][string] # Need one of these configured
        $PrimaryScaleSetName,
        [Parameter(Mandatory = $true)][string]
        $UseManagedIdentityExtension,
        [string]
        $UserAssignedClientID,
        [Parameter(Mandatory = $true)][string]
        $UseInstanceMetadata,
        [Parameter(Mandatory = $true)][string]
        $LoadBalancerSku,
        [Parameter(Mandatory = $true)][string]
        $ExcludeMasterFromStandardLB,
        [Parameter(Mandatory = $true)][string]
        $KubeDir,
        [Parameter(Mandatory = $true)][string]
        $TargetEnvironment
    )

    if ( -Not $PrimaryAvailabilitySetName -And -Not $PrimaryScaleSetName ) {
        throw "Either PrimaryAvailabilitySetName or PrimaryScaleSetName must be set"
    }

    $azureConfigFile = [io.path]::Combine($KubeDir, "azure.json")

    $azureConfig = @"
{
    "cloud": "$TargetEnvironment",
    "tenantId": "$TenantId",
    "subscriptionId": "$SubscriptionId",
    "aadClientId": "$AADClientId",
    "aadClientSecret": "$AADClientSecret",
    "resourceGroup": "$ResourceGroup",
    "location": "$Location",
    "vmType": "$VmType",
    "subnetName": "$SubnetName",
    "securityGroupName": "$SecurityGroupName",
    "vnetName": "$VNetName",
    "routeTableName": "$RouteTableName",
    "primaryAvailabilitySetName": "$PrimaryAvailabilitySetName",
    "primaryScaleSetName": "$PrimaryScaleSetName",
    "useManagedIdentityExtension": $UseManagedIdentityExtension,
    "userAssignedIdentityID": $UserAssignedClientID,
    "useInstanceMetadata": $UseInstanceMetadata,
    "loadBalancerSku": "$LoadBalancerSku",
    "excludeMasterFromStandardLB": $ExcludeMasterFromStandardLB
}
"@

    $azureConfig | Out-File -encoding ASCII -filepath "$azureConfigFile"
}


function
Write-CACert {
    Param(
        [Parameter(Mandatory = $true)][string]
        $CACertificate,
        [Parameter(Mandatory = $true)][string]
        $KubeDir
    )
    $caFile = [io.path]::Combine($KubeDir, "ca.crt")
    [System.Text.Encoding]::ASCII.GetString([System.Convert]::FromBase64String($CACertificate)) | Out-File -Encoding ascii $caFile
}

function
Write-KubeConfig {
    Param(
        [Parameter(Mandatory = $true)][string]
        $CACertificate,
        [Parameter(Mandatory = $true)][string]
        $MasterFQDNPrefix,
        [Parameter(Mandatory = $true)][string]
        $MasterIP,
        [Parameter(Mandatory = $true)][string]
        $AgentKey,
        [Parameter(Mandatory = $true)][string]
        $AgentCertificate,
        [Parameter(Mandatory = $true)][string]
        $KubeDir
    )
    $kubeConfigFile = [io.path]::Combine($KubeDir, "config")

    $kubeConfig = @"
---
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: "$CACertificate"
    server: https://${MasterIP}:443
  name: "$MasterFQDNPrefix"
contexts:
- context:
    cluster: "$MasterFQDNPrefix"
    user: "$MasterFQDNPrefix-admin"
  name: "$MasterFQDNPrefix"
current-context: "$MasterFQDNPrefix"
kind: Config
users:
- name: "$MasterFQDNPrefix-admin"
  user:
    client-certificate-data: "$AgentCertificate"
    client-key-data: "$AgentKey"
"@

    $kubeConfig | Out-File -encoding ASCII -filepath "$kubeConfigFile"
}

function
Build-PauseContainer {
    Param(
        [Parameter(Mandatory = $true)][string]
        $WindowsBase,
        $DestinationTag
    )
    # Future work: This needs to build wincat - see https://github.com/Azure/aks-engine/issues/1461
    "FROM $($WindowsBase)" | Out-File -encoding ascii -FilePath Dockerfile
    "CMD cmd /c ping -t localhost" | Out-File -encoding ascii -FilePath Dockerfile -Append
    docker build -t $DestinationTag .
}

function
New-InfraContainer {
    Param(
        [Parameter(Mandatory = $true)][string]
        $KubeDir,
        $DestinationTag = "kubletwin/pause"
    )
    cd $KubeDir
    $computerInfo = Get-ComputerInfo

    # Reference for these tags: curl -L https://mcr.microsoft.com/v2/k8s/core/pause/tags/list
    # Then docker run --rm mplatform/manifest-tool inspect mcr.microsoft.com/k8s/core/pause:<tag>

    $defaultPauseImage = "mcr.microsoft.com/k8s/core/pause:1.2.0"

    switch ($computerInfo.WindowsVersion) {
        "1803" { docker pull $defaultPauseImage ; docker tag $defaultPauseImage $DestinationTag }
        "1809" { docker pull $defaultPauseImage ; docker tag $defaultPauseImage $DestinationTag }
        "1903" { Build-PauseContainer -WindowsBase "mcr.microsoft.com/windows/nanoserver:1903" -DestinationTag $DestinationTag}
        default { Build-PauseContainer -WindowsBase "mcr.microsoft.com/nanoserver-insider" -DestinationTag $DestinationTag}
    }
}


# TODO: Deprecate this and replace with methods that get individual components instead of zip containing everything
# This expects the ZIP file to be created by scripts/build-windows-k8s.sh
function
Get-KubePackage {
    Param(
        [Parameter(Mandatory = $true)][string]
        $KubeBinariesSASURL
    )

    $zipfile = "c:\k.zip"
    for ($i = 0; $i -le 10; $i++) {
        DownloadFileOverHttp -Url $KubeBinariesSASURL -DestinationPath $zipfile
        if ($?) {
            break
        }
        else {
            Write-Log $Error[0].Exception.Message
        }
    }
    Expand-Archive -path $zipfile -DestinationPath C:\
}

function
Get-KubeBinaries {
    Param(
        [Parameter(Mandatory = $true)][string]
        $KubeBinariesURL
    )

    $tempdir = New-TemporaryDirectory
    $binaryPackage = "$tempdir\k.tar.gz"
    for ($i = 0; $i -le 10; $i++) {
        DownloadFileOverHttp -Url $KubeBinariesURL -DestinationPath $binaryPackage
        if ($?) {
            break
        }
        else {
            Write-Log $Error[0].Exception.Message
        }
    }

    # using tar to minimize dependencies
    # tar should be avalible on 1803+
    tar -xzf $binaryPackage -C $tempdir

    # copy binaries over to kube folder
    $windowsbinariespath = "c:\k\"
    if (!(Test-path $windowsbinariespath)) {
        mkdir $windowsbinariespath
    }
    cp $tempdir\kubernetes\node\bin\* $windowsbinariespath -Recurse

    #remove temp folder created when unzipping
    del $tempdir -Recurse
}

# TODO: replace KubeletStartFile with a Kubelet config, remove NSSM, and use built-in service integration
function
New-NSSMService {
    Param(
        [string]
        [Parameter(Mandatory = $true)]
        $KubeDir,
        [string]
        [Parameter(Mandatory = $true)]
        $KubeletStartFile,
        [string]
        [Parameter(Mandatory = $true)]
        $KubeProxyStartFile
    )

    # setup kubelet
    & "$KubeDir\nssm.exe" install Kubelet C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe
    & "$KubeDir\nssm.exe" set Kubelet AppDirectory $KubeDir
    & "$KubeDir\nssm.exe" set Kubelet AppParameters $KubeletStartFile
    & "$KubeDir\nssm.exe" set Kubelet DisplayName Kubelet
    & "$KubeDir\nssm.exe" set Kubelet AppRestartDelay 5000
    & "$KubeDir\nssm.exe" set Kubelet DependOnService docker
    & "$KubeDir\nssm.exe" set Kubelet Description Kubelet
    & "$KubeDir\nssm.exe" set Kubelet Start SERVICE_AUTO_START
    & "$KubeDir\nssm.exe" set Kubelet ObjectName LocalSystem
    & "$KubeDir\nssm.exe" set Kubelet Type SERVICE_WIN32_OWN_PROCESS
    & "$KubeDir\nssm.exe" set Kubelet AppThrottle 1500
    & "$KubeDir\nssm.exe" set Kubelet AppStdout C:\k\kubelet.log
    & "$KubeDir\nssm.exe" set Kubelet AppStderr C:\k\kubelet.err.log
    & "$KubeDir\nssm.exe" set Kubelet AppStdoutCreationDisposition 4
    & "$KubeDir\nssm.exe" set Kubelet AppStderrCreationDisposition 4
    & "$KubeDir\nssm.exe" set Kubelet AppRotateFiles 1
    & "$KubeDir\nssm.exe" set Kubelet AppRotateOnline 1
    & "$KubeDir\nssm.exe" set Kubelet AppRotateSeconds 86400
    & "$KubeDir\nssm.exe" set Kubelet AppRotateBytes 1048576

    # setup kubeproxy
    & "$KubeDir\nssm.exe" install Kubeproxy C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe
    & "$KubeDir\nssm.exe" set Kubeproxy AppDirectory $KubeDir
    & "$KubeDir\nssm.exe" set Kubeproxy AppParameters $KubeProxyStartFile
    & "$KubeDir\nssm.exe" set Kubeproxy DisplayName Kubeproxy
    & "$KubeDir\nssm.exe" set Kubeproxy DependOnService Kubelet
    & "$KubeDir\nssm.exe" set Kubeproxy Description Kubeproxy
    & "$KubeDir\nssm.exe" set Kubeproxy Start SERVICE_AUTO_START
    & "$KubeDir\nssm.exe" set Kubeproxy ObjectName LocalSystem
    & "$KubeDir\nssm.exe" set Kubeproxy Type SERVICE_WIN32_OWN_PROCESS
    & "$KubeDir\nssm.exe" set Kubeproxy AppThrottle 1500
    & "$KubeDir\nssm.exe" set Kubeproxy AppStdout C:\k\kubeproxy.log
    & "$KubeDir\nssm.exe" set Kubeproxy AppStderr C:\k\kubeproxy.err.log
    & "$KubeDir\nssm.exe" set Kubeproxy AppRotateFiles 1
    & "$KubeDir\nssm.exe" set Kubeproxy AppRotateOnline 1
    & "$KubeDir\nssm.exe" set Kubeproxy AppRotateSeconds 86400
    & "$KubeDir\nssm.exe" set Kubeproxy AppRotateBytes 1048576
}

# Renamed from Write-KubernetesStartFiles
function
Install-KubernetesServices {
    param(
        [Parameter(Mandatory = $true)][string[]]
        $KubeletConfigArgs,
        [Parameter(Mandatory = $true)][string]
        $KubeBinariesVersion,
        [Parameter(Mandatory = $true)][string]
        $NetworkPlugin,
        [Parameter(Mandatory = $true)][string]
        $NetworkMode,
        [Parameter(Mandatory = $true)][string]
        $KubeDir,
        [Parameter(Mandatory = $true)][string]
        $AzureCNIBinDir,
        [Parameter(Mandatory = $true)][string]
        $AzureCNIConfDir,
        [Parameter(Mandatory = $true)][string]
        $CNIPath,
        [Parameter(Mandatory = $true)][string]
        $CNIConfig,
        [Parameter(Mandatory = $true)][string]
        $CNIConfigPath,
        [Parameter(Mandatory = $true)][string]
        $MasterIP,
        [Parameter(Mandatory = $true)][string]
        $KubeDnsServiceIp,
        [Parameter(Mandatory = $true)][string]
        $MasterSubnet,
        [Parameter(Mandatory = $true)][string]
        $KubeClusterCIDR,
        [Parameter(Mandatory = $true)][string]
        $KubeServiceCIDR,
        [Parameter(Mandatory = $true)][string]
        $HNSModule,
        [Parameter(Mandatory = $true)][string]
        $KubeletNodeLabels
    )

    # Calculate some local paths
    $VolumePluginDir = [Io.path]::Combine($KubeDir, "volumeplugins")
    $KubeletStartFile = [io.path]::Combine($KubeDir, "kubeletstart.ps1")
    $KubeProxyStartFile = [io.path]::Combine($KubeDir, "kubeproxystart.ps1")

    mkdir $VolumePluginDir
    $KubeletArgList = $KubeletConfigArgs # This is the initial list passed in from aks-engine
    $KubeletArgList += "--node-labels=` + "`" + `$global:KubeletNodeLabels"
    # $KubeletArgList += "--hostname-override=` + "`" + `$global:AzureHostname" TODO: remove - dead code?
    $KubeletArgList += "--volume-plugin-dir=` + "`" + `$global:VolumePluginDir"
    # If you are thinking about adding another arg here, you should be considering pkg/engine/defaults-kubelet.go first
    # Only args that need to be calculated or combined with other ones on the Windows agent should be added here.


    # Regex to strip version to Major.Minor.Build format such that the following check does not crash for version like x.y.z-alpha
    [regex]$regex = "^[0-9.]+"
    $KubeBinariesVersionStripped = $regex.Matches($KubeBinariesVersion).Value
    if ([System.Version]$KubeBinariesVersionStripped -lt [System.Version]"1.8.0") {
        # --api-server deprecates from 1.8.0
        $KubeletArgList += "--api-servers=https://` + "`" + `${global:MasterIP}:443"
    }

    # Configure kubelet to use CNI plugins if enabled.
    if ($NetworkPlugin -eq "azure") {
        $KubeletArgList += @("--cni-bin-dir=$AzureCNIBinDir", "--cni-conf-dir=$AzureCNIConfDir")
    }
    elseif ($NetworkPlugin -eq "kubenet") {
        $KubeletArgList += @("--cni-bin-dir=$CNIPath", "--cni-conf-dir=$CNIConfigPath")
        # handle difference in naming between Linux & Windows reference plugin
        $KubeletArgList = $KubeletArgList -replace "kubenet", "cni"
    }
    else {
        throw "Unknown network type $NetworkPlugin, can't configure kubelet"
    }

    # Used in WinCNI version of kubeletstart.ps1
    $KubeletArgListStr = ""
    $KubeletArgList | Foreach-Object {
        # Since generating new code to be written to a file, need to escape quotes again
        if ($KubeletArgListStr.length -gt 0) {
            $KubeletArgListStr = $KubeletArgListStr + ", "
        }
        $KubeletArgListStr = $KubeletArgListStr + "` + "`" + `"" + $_.Replace("` + "`" + `"` + "`" + `"", "` + "`" + `"` + "`" + `"` + "`" + `"` + "`" + `"") + "` + "`" + `""
    }
    $KubeletArgListStr = "@` + "`" + `($KubeletArgListStr` + "`" + `)"

    # Used in Azure-CNI version of kubeletstart.ps1
    $KubeletCommandLine = "$KubeDir\kubelet.exe " + ($KubeletArgList -join " ")

    $kubeStartStr = @"
` + "`" + `$global:MasterIP = "$MasterIP"
` + "`" + `$global:KubeDnsSearchPath = "svc.cluster.local"
` + "`" + `$global:KubeDnsServiceIp = "$KubeDnsServiceIp"
` + "`" + `$global:MasterSubnet = "$MasterSubnet"
` + "`" + `$global:KubeClusterCIDR = "$KubeClusterCIDR"
` + "`" + `$global:KubeServiceCIDR = "$KubeServiceCIDR"
` + "`" + `$global:KubeBinariesVersion = "$KubeBinariesVersion"
` + "`" + `$global:CNIPath = "$CNIPath"
` + "`" + `$global:NetworkMode = "$NetworkMode"
` + "`" + `$global:ExternalNetwork = "ext"
` + "`" + `$global:CNIConfig = "$CNIConfig"
` + "`" + `$global:HNSModule = "$HNSModule"
` + "`" + `$global:VolumePluginDir = "$VolumePluginDir"
` + "`" + `$global:NetworkPlugin="$NetworkPlugin"
` + "`" + `$global:KubeletNodeLabels="$KubeletNodeLabels"

"@

    if ($NetworkPlugin -eq "azure") {
        $KubeNetwork = "azure"
        $kubeStartStr += @"
Write-Host "NetworkPlugin azure, starting kubelet."

# Turn off Firewall to enable pods to talk to service endpoints. (Kubelet should eventually do this)
netsh advfirewall set allprofiles state off
# startup the service

# Find if the primary external switch network exists. If not create one.
# This is done only once in the lifetime of the node
` + "`" + `$hnsNetwork = Get-HnsNetwork | ? Name -EQ ` + "`" + `$global:ExternalNetwork
if (!` + "`" + `$hnsNetwork)
{
    Write-Host "Creating a new hns Network"
    ipmo ` + "`" + `$global:HNSModule
    # Fixme : use a smallest range possible, that will not collide with any pod space
    New-HNSNetwork -Type ` + "`" + `$global:NetworkMode -AddressPrefix "192.168.255.0/30" -Gateway "192.168.255.1" -Name ` + "`" + `$global:ExternalNetwork -Verbose
}

# Find if network created by CNI exists, if yes, remove it
# This is required to keep the network non-persistent behavior
# Going forward, this would be done by HNS automatically during restart of the node

` + "`" + `$hnsNetwork = Get-HnsNetwork | ? Name -EQ $KubeNetwork
if (` + "`" + `$hnsNetwork)
{
    # Cleanup all containers
    docker ps -q | foreach {docker rm ` + "`" + `$_ -f}

    Write-Host "Cleaning up old HNS network found"
    Remove-HnsNetwork ` + "`" + `$hnsNetwork
    # Kill all cni instances & stale data left by cni
    # Cleanup all files related to cni
    taskkill /IM azure-vnet.exe /f
    taskkill /IM azure-vnet-ipam.exe /f
    ` + "`" + `$cnijson = [io.path]::Combine("$KubeDir", "azure-vnet-ipam.json")
    if ((Test-Path ` + "`" + `$cnijson))
    {
        Remove-Item ` + "`" + `$cnijson
    }
    ` + "`" + `$cnilock = [io.path]::Combine("$KubeDir", "azure-vnet-ipam.json.lock")
    if ((Test-Path ` + "`" + `$cnilock))
    {
        Remove-Item ` + "`" + `$cnilock
    }

    ` + "`" + `$cnijson = [io.path]::Combine("$KubeDir", "azure-vnet.json")
    if ((Test-Path ` + "`" + `$cnijson))
    {
        Remove-Item ` + "`" + `$cnijson
    }
    ` + "`" + `$cnilock = [io.path]::Combine("$KubeDir", "azure-vnet.json.lock")
    if ((Test-Path ` + "`" + `$cnilock))
    {
        Remove-Item ` + "`" + `$cnilock
    }
}

# Restart Kubeproxy, which would wait, until the network is created
Restart-Service Kubeproxy

` + "`" + `$env:AZURE_ENVIRONMENT_FILEPATH="c:\k\azurestackcloud.json"

$KubeletCommandLine

"@
    }
    else {
        # using WinCNI. TODO: If WinCNI support is removed, then delete this as dead code later
        $KubeNetwork = "l2bridge"
        $kubeStartStr += @"

function
Get-DefaultGateway(` + "`" + `$CIDR)
{
    return ` + "`" + `$CIDR.substring(0,` + "`" + `$CIDR.lastIndexOf(".")) + ".1"
}

function
Get-PodCIDR()
{
    ` + "`" + `$podCIDR = c:\k\kubectl.exe --kubeconfig=c:\k\config get nodes/` + "`" + `$(` + "`" + `$env:computername.ToLower()) -o custom-columns=podCidr:.spec.podCIDR --no-headers
    return ` + "`" + `$podCIDR
}

function
Test-PodCIDR(` + "`" + `$podCIDR)
{
    return ` + "`" + `$podCIDR.length -gt 0
}

function
Update-CNIConfig(` + "`" + `$podCIDR, ` + "`" + `$masterSubnetGW)
{
    ` + "`" + `$jsonSampleConfig =
"{
    ""cniVersion"": ""0.2.0"",
    ""name"": ""<NetworkMode>"",
    ""type"": ""win-bridge"",
    ""master"": ""Ethernet"",
    ""dns"" : {
        ""Nameservers"" : [ ""<NameServers>"" ],
        ""Search"" : [ ""<Cluster DNS Suffix or Search Path>"" ]
    },
    ""policies"": [
    {
        ""Name"" : ""EndpointPolicy"", ""Value"" : { ""Type"" : ""OutBoundNAT"", ""ExceptionList"": [ ""<ClusterCIDR>"", ""<MgmtSubnet>"" ] }
    },
    {
        ""Name"" : ""EndpointPolicy"", ""Value"" : { ""Type"" : ""ROUTE"", ""DestinationPrefix"": ""<ServiceCIDR>"", ""NeedEncap"" : true }
    }
    ]
}"

    ` + "`" + `$configJson = ConvertFrom-Json ` + "`" + `$jsonSampleConfig
    ` + "`" + `$configJson.name = ` + "`" + `$global:NetworkMode.ToLower()
    ` + "`" + `$configJson.dns.Nameservers[0] = ` + "`" + `$global:KubeDnsServiceIp
    ` + "`" + `$configJson.dns.Search[0] = ` + "`" + `$global:KubeDnsSearchPath

    ` + "`" + `$configJson.policies[0].Value.ExceptionList[0] = ` + "`" + `$global:KubeClusterCIDR
    ` + "`" + `$configJson.policies[0].Value.ExceptionList[1] = ` + "`" + `$global:MasterSubnet
    ` + "`" + `$configJson.policies[1].Value.DestinationPrefix  = ` + "`" + `$global:KubeServiceCIDR

    if (Test-Path ` + "`" + `$global:CNIConfig)
    {
        Clear-Content -Path ` + "`" + `$global:CNIConfig
    }

    Write-Host "Generated CNI Config [` + "`" + `$configJson]"

    Add-Content -Path ` + "`" + `$global:CNIConfig -Value (ConvertTo-Json ` + "`" + `$configJson -Depth 20)
}

try
{
    ` + "`" + `$env:AZURE_ENVIRONMENT_FILEPATH="c:\k\azurestackcloud.json"

    ` + "`" + `$masterSubnetGW = Get-DefaultGateway ` + "`" + `$global:MasterSubnet
    ` + "`" + `$podCIDR=Get-PodCIDR
    ` + "`" + `$podCidrDiscovered=Test-PodCIDR(` + "`" + `$podCIDR)

    # if the podCIDR has not yet been assigned to this node, start the kubelet process to get the podCIDR, and then promptly kill it.
    if (-not ` + "`" + `$podCidrDiscovered)
    {
        ` + "`" + `$argList = $KubeletArgListStr

        ` + "`" + `$process = Start-Process -FilePath c:\k\kubelet.exe -PassThru -ArgumentList ` + "`" + `$argList

        # run kubelet until podCidr is discovered
        Write-Host "waiting to discover pod CIDR"
        while (-not ` + "`" + `$podCidrDiscovered)
        {
            Write-Host "Sleeping for 10s, and then waiting to discover pod CIDR"
            Start-Sleep 10

            ` + "`" + `$podCIDR=Get-PodCIDR
            ` + "`" + `$podCidrDiscovered=Test-PodCIDR(` + "`" + `$podCIDR)
        }

        # stop the kubelet process now that we have our CIDR, discard the process output
        ` + "`" + `$process | Stop-Process | Out-Null
    }

    # Turn off Firewall to enable pods to talk to service endpoints. (Kubelet should eventually do this)
    netsh advfirewall set allprofiles state off

    # startup the service
    ` + "`" + `$hnsNetwork = Get-HnsNetwork | ? Name -EQ ` + "`" + `$global:NetworkMode.ToLower()

    if (` + "`" + `$hnsNetwork)
    {
        # Kubelet has been restarted with existing network.
        # Cleanup all containers
        docker ps -q | foreach {docker rm ` + "`" + `$_ -f}
        # cleanup network
        Write-Host "Cleaning up old HNS network found"
        Remove-HnsNetwork ` + "`" + `$hnsNetwork
        Start-Sleep 10
    }

    Write-Host "Creating a new hns Network"
    ipmo ` + "`" + `$global:HNSModule

    ` + "`" + `$hnsNetwork = New-HNSNetwork -Type ` + "`" + `$global:NetworkMode -AddressPrefix ` + "`" + `$podCIDR -Gateway ` + "`" + `$masterSubnetGW -Name ` + "`" + `$global:NetworkMode.ToLower() -Verbose
    # New network has been created, Kubeproxy service has to be restarted
    Restart-Service Kubeproxy

    Start-Sleep 10
    # Add route to all other POD networks
    Update-CNIConfig ` + "`" + `$podCIDR ` + "`" + `$masterSubnetGW

    $KubeletCommandLine
}
catch
{
    Write-Error ` + "`" + `$_
}

"@
    } # end else using WinCNI.

    # Now that the script is generated, based on what CNI plugin and startup options are needed, write it to disk
    $kubeStartStr | Out-File -encoding ASCII -filepath $KubeletStartFile

    $kubeProxyStartStr = @"
` + "`" + `$env:KUBE_NETWORK = "$KubeNetwork"
` + "`" + `$global:NetworkMode = "$NetworkMode"
` + "`" + `$global:HNSModule = "$HNSModule"
` + "`" + `$hnsNetwork = Get-HnsNetwork | ? Name -EQ $KubeNetwork
while (!` + "`" + `$hnsNetwork)
{
    Write-Host "Waiting for Network [$KubeNetwork] to be created . . ."
    Start-Sleep 10
    ` + "`" + `$hnsNetwork = Get-HnsNetwork | ? Name -EQ $KubeNetwork
}

#
# cleanup the persisted policy lists
#
ipmo ` + "`" + `$global:HNSModule
Get-HnsPolicyList | Remove-HnsPolicyList

$KubeDir\kube-proxy.exe --v=3 --proxy-mode=kernelspace --hostname-override=$env:computername --kubeconfig=$KubeDir\config
"@

    $kubeProxyStartStr | Out-File -encoding ASCII -filepath $KubeProxyStartFile

    New-NSSMService -KubeDir $KubeDir ` + "`" + `
        -KubeletStartFile $KubeletStartFile ` + "`" + `
        -KubeProxyStartFile $KubeProxyStartFile
}
`)

func k8sWindowskubeletfuncPs1Bytes() ([]byte, error) {
	return _k8sWindowskubeletfuncPs1, nil
}

func k8sWindowskubeletfuncPs1() (*asset, error) {
	bytes, err := k8sWindowskubeletfuncPs1Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/windowskubeletfunc.ps1", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masteroutputsT = []byte(`    "masterFQDN": {
      "type": "string",
{{if not IsPrivateCluster}}
      "value": "[reference(concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))).dnsSettings.fqdn]"
{{else}}
      "value": ""
{{end}}
    }
{{if AnyAgentUsesAvailabilitySets}}
    ,
    "agentStorageAccountSuffix": {
      "type": "string",
      "value": "[variables('storageAccountBaseName')]"
    },
    "agentStorageAccountPrefixes": {
      "type": "array",
      "value": "[variables('storageAccountPrefixes')]"
    }
{{end}}`)

func masteroutputsTBytes() ([]byte, error) {
	return _masteroutputsT, nil
}

func masteroutputsT() (*asset, error) {
	bytes, err := masteroutputsTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "masteroutputs.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterparamsT = []byte(`    "linuxAdminUsername": {
      "metadata": {
        "description": "User name for the Linux Virtual Machines (SSH or Password)."
      },
      "type": "string"
    },
    "masterEndpointDNSNamePrefix": {
      "metadata": {
        "description": "Sets the Domain name label for the master IP Address.  The concatenation of the domain name label and the regional DNS zone make up the fully qualified domain name associated with the public IP address."
      },
      "type": "string"
    },
    "aksEngineVersion": {
      "metadata": {
        "description": "Contains details of the aks-engine version which was used to provision the cluster"
      },
      "type": "string"
    },
    {{range .ExtensionProfiles}}
      "{{.Name}}Parameters": {
        "metadata": {
        "description": "Parameters for the extension"
      },
      "type": "securestring"
      },
    {{end}}
{{if not IsHostedMaster }}
  {{if .MasterProfile.IsCustomVNET}}
    "masterVnetSubnetID": {
      "metadata": {
        "description": "Sets the vnet subnet of the master."
      },
      "type": "string"
    },
    {{if .MasterProfile.IsVirtualMachineScaleSets}}
    "agentVnetSubnetID": {
      "metadata": {
        "description": "Sets the vnet subnet of the agent."
      },
      "type": "string"
    },
    {{end}}
    "masterSubnet": {
      "defaultValue": "",
      "metadata": {
        "description": "Sets the subnet of the master node(s)"
      },
      "type": "string"
    },
  {{else}}
    "masterSubnet": {
      "defaultValue": "{{.MasterProfile.Subnet}}",
      "metadata": {
        "description": "Sets the subnet of the master node(s)."
      },
      "type": "string"
    },
    "agentSubnet": {
      "defaultValue": "{{.MasterProfile.AgentSubnet}}",
      "metadata": {
        "description": "Sets the subnet of the agent node(s)."
      },
      "type": "string"
    },
  {{end}}
  "masterSubnetIPv6": {
      "defaultValue": "{{.MasterProfile.SubnetIPv6}}",
      "metadata": {
        "description": "Sets the IPv6 subnet of the master node(s)."
      },
      "type": "string"
    },
  {{if .MasterProfile.HasAvailabilityZones}}
  "availabilityZones": {
    "metadata": {
      "description": "Master availability zones"
    },
    "type": "array"
  },
  {{end}}
{{end}}
{{if IsHostedMaster}}
    "masterSubnet": {
      "defaultValue": "{{.HostedMasterProfile.Subnet}}",
      "metadata": {
        "description": "Sets the subnet for the VMs in the cluster."
      },
      "type": "string"
    },
    "kubernetesEndpoint": {
      "defaultValue": "{{.HostedMasterProfile.FQDN}}",
      "metadata": {
        "description": "Sets the static IP of the first master"
      },
      "type": "string"
    },
{{else}}
    "firstConsecutiveStaticIP": {
      "defaultValue": "{{.MasterProfile.FirstConsecutiveStaticIP}}",
      "metadata": {
        "description": "Sets the static IP of the first master"
      },
      "type": "string"
    },
    "masterVMSize": {
      {{GetMasterAllowedSizes}}
      "metadata": {
        "description": "The size of the Virtual Machine."
      },
      "type": "string"
    },
{{end}}
    "sshRSAPublicKey": {
      "metadata": {
        "description": "SSH public key used for auth to all Linux machines.  Not Required.  If not set, you must provide a password key."
      },
      "type": "string"
    },
    "nameSuffix": {
      "defaultValue": "{{GetUniqueNameSuffix}}",
      "metadata": {
        "description": "A string hash of the master DNS name to uniquely identify the cluster."
      },
      "type": "string"
    },
    "osImageName": {
      "defaultValue": "",
      "metadata": {
        "description": "Name of a Linux OS image. Needs to be used in conjuction with osImageResourceGroup."
      },
      "type": "string"
    },
    "osImageResourceGroup": {
      "defaultValue": "",
      "metadata": {
        "description": "Resource group of a Linux OS image. Needs to be used in conjuction with osImageName."
      },
      "type": "string"
    },
    "osImageOffer": {
      "defaultValue": "UbuntuServer",
      "metadata": {
        "description": "Linux OS image type."
      },
      "type": "string"
    },
    "osImagePublisher": {
      "defaultValue": "Canonical",
      "metadata": {
        "description": "OS image publisher."
      },
      "type": "string"
    },
    "osImageSKU": {
      "defaultValue": "16.04-LTS",
      "metadata": {
        "description": "OS image SKU."
      },
      "type": "string"
    },
    "osImageVersion": {
      "defaultValue": "latest",
      "metadata": {
        "description": "OS image version."
      },
      "type": "string"
    },
    "fqdnEndpointSuffix":{
      "defaultValue": "cloudapp.azure.com",
      "metadata": {
        "description": "Endpoint of FQDN."
      },
      "type": "string"
    },
    "targetEnvironment": {
      "defaultValue": "AzurePublicCloud",
      "metadata": {
        "description": "The azure deploy environment. Currently support: AzurePublicCloud, AzureChinaCloud"
      },
      "type": "string"
    },
    "location": {
      "defaultValue": "{{GetLocation}}",
      "metadata": {
        "description": "Sets the location for all resources in the cluster"
      },
      "type": "string"
    }
{{if .LinuxProfile}}{{if .LinuxProfile.HasSecrets}}
  {{range  $vIndex, $vault := .LinuxProfile.Secrets}}
    ,
    "linuxKeyVaultID{{$vIndex}}": {
      "metadata": {
        "description": "KeyVaultId{{$vIndex}} to install certificates from on linux machines."
      },
      "type": "string"
    }
    {{range $cIndex, $cert := $vault.VaultCertificates}}
      ,
      "linuxKeyVaultID{{$vIndex}}CertificateURL{{$cIndex}}": {
        "metadata": {
          "description": "CertificateURL{{$cIndex}} to install from KeyVaultId{{$vIndex}} on linux machines."
        },
        "type": "string"
      }
    {{end}}
  {{end}}
{{end}}{{end}}
{{if .HasWindows}}{{if .WindowsProfile.HasSecrets}}
  {{range  $vIndex, $vault := .WindowsProfile.Secrets}}
    ,
    "windowsKeyVaultID{{$vIndex}}": {
      "metadata": {
        "description": "KeyVaultId{{$vIndex}} to install certificates from on windows machines."
      },
      "type": "string"
    }
    {{range $cIndex, $cert := $vault.VaultCertificates}}
      ,
      "windowsKeyVaultID{{$vIndex}}CertificateURL{{$cIndex}}": {
        "metadata": {
          "description": "Url to retrieve Certificate{{$cIndex}} from KeyVaultId{{$vIndex}} to install on windows machines."
        },
        "type": "string"
      },
      "windowsKeyVaultID{{$vIndex}}CertificateStore{{$cIndex}}": {
        "metadata": {
          "description": "CertificateStore to install Certificate{{$cIndex}} from KeyVaultId{{$vIndex}} on windows machines."
        },
        "type": "string"
      }
    {{end}}
  {{end}}
{{end}} {{end}}
`)

func masterparamsTBytes() ([]byte, error) {
	return _masterparamsT, nil
}

func masterparamsT() (*asset, error) {
	bytes, err := masterparamsTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "masterparams.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _swarmInstallContainerhostAndJoinSwarmPs1 = []byte(`############################################################
# Script adapted from
# https://raw.githubusercontent.com/Microsoft/Virtualization-Documentation/master/windows-server-container-tools/Install-ContainerHost/Install-ContainerHost.ps1

<#
    .NOTES
        Copyright (c) Microsoft Corporation.  All rights reserved.

        Use of this sample source code is subject to the terms of the Microsoft
        license agreement under which you licensed this sample source code. If
        you did not accept the terms of the license agreement, you are not
        authorized to use this sample source code. For the terms of the license,
        please see the license agreement between you and Microsoft or, if applicable,
        see the LICENSE.RTF on your install media or the root of your tools installation.
        THE SAMPLE SOURCE CODE IS PROVIDED "AS IS", WITH NO WARRANTIES.

    .SYNOPSIS
        Installs the prerequisites for creating Windows containers
        Opens TCP ports (80,443,2375,8080) in Windows Firewall.
        Connects Docker to a swarm master.

    .DESCRIPTION
        Installs the prerequisites for creating Windows containers
        Opens TCP ports (80,443,2375,8080) in Windows Firewall.
        Connects Docker to a swarm master.

    .PARAMETER SwarmMasterIP
        IP Address of Docker Swarm Master

    .EXAMPLE
        .\Install-ContainerHost.ps1 -SwarmMasterIP 192.168.255.5

#>
#Requires -Version 5.0

[CmdletBinding(DefaultParameterSetName="Standard")]
param(
    [string]
    [ValidateNotNullOrEmpty()]
    $SwarmMasterIP = "172.16.0.5"
)

$global:DockerServiceName = "Docker"
$global:HNSServiceName = "hns"

filter Timestamp {"$(Get-Date -Format o): $_"}

function Write-Log($message)
{
    $msg = $message | Timestamp
    Write-Output $msg
}

function
Start-Docker()
{
    Write-Log "Starting $global:DockerServiceName..."
    $startTime = Get-Date
        
    while (-not $dockerReady)
    {
        try
        {
            Start-Service -Name $global:DockerServiceName -ea Stop

            $dockerReady = $true            
        }
        catch
        {
            $timeElapsed = $(Get-Date) - $startTime
            if ($($timeElapsed).TotalMinutes -ge 5)
            {
                Write-Log "Docker Daemon did not start successfully within 5 minutes."
                break
            }

            $errorStr = $_.Exception.Message
            Write-Log "Starting Service failed: $errorStr" 
            Write-Log "sleeping for 10 seconds..."
            Start-Sleep -sec 10
        }
    }
}


function
Stop-Docker()
{
    Write-Log "Stopping $global:DockerServiceName..."
    try
    {
        Stop-Service -Name $global:DockerServiceName -ea Stop   
    }
    catch
    {
        Write-Log "Failed to stop Docker"
    }
}

function
Update-DockerServiceRecoveryPolicy()
{
    $dockerReady = $false
    $startTime = Get-Date
    
    # wait until the service exists
    while (-not $dockerReady)
    {
        if (Get-Service $global:DockerServiceName -ErrorAction SilentlyContinue)
        {
            $dockerReady = $true
        }
        else 
        {
            $timeElapsed = $(Get-Date) - $startTime
            if ($($timeElapsed).TotalMinutes -ge 5)
            {
                Write-Log "Unable to find service $global:DockerServiceName within 5 minutes."
                break
            }
            Write-Log "failed to find $global:DockerServiceName, sleeping for 5 seconds"
            Start-Sleep -sec 5
        }
    }
    
    Write-Log "Updating docker restart policy, to ensure it restarts on error"
    $services = Get-WMIObject win32_service | Where-Object {$_.name -imatch $global:DockerServiceName}
    foreach ($service in $services)
    {
        sc.exe failure $service.name reset= 86400 actions= restart/5000
    }
}

# Open Windows Firewall Ports Needed
function Open-FirewallPorts()
{
    $ports = @(80,443,2375,8080)
    foreach ($port in $ports)
    {
        $netsh = "netsh advfirewall firewall add rule name='Open Port $port' dir=in action=allow protocol=TCP localport=$port"
        Write-Log "enabling port with command $netsh"
        Invoke-Expression -Command:$netsh
    }
}

# Update Docker Config to have cluster-store=consul:// address configured for Swarm cluster.
function Write-DockerDaemonJson()
{
    $dataDir = $env:ProgramData

    # create the target directory
    $targetDir = $dataDir + '\docker\config'
    if(!(Test-Path -Path $targetDir )){
        New-Item -ItemType directory -Path $targetDir
    }

    Write-Log "Delete key file, so that this node is unique to swarm"
    $keyFileName = "$targetDir\key.json"
    Write-Log "Removing $($keyFileName)"
    if (Test-Path $keyFileName) {
      Remove-Item $keyFileName
    }

    $ipAddress = Get-IPAddress

    Write-Log "Advertise $($ipAddress) to consul://$($SwarmMasterIP):8500"
    $OutFile = @"
{
    "hosts": ["tcp://0.0.0.0:2375", "npipe://"],
    "cluster-store": "consul://$($SwarmMasterIP):8500",
    "cluster-advertise": "$($ipAddress):2375"
}
"@

    $OutFile | Out-File -encoding ASCII -filepath "$targetDir\daemon.json"
}

# Get Node IPV4 Address
function Get-IPAddress()
{
    return (Get-NetIPAddress | where {$_.IPAddress -Like '10.*' -and $_.AddressFamily -eq 'IPV4'})[0].IPAddress
}

try
{
    Write-Log "Provisioning $global:DockerServiceName... with Swarm IP $SwarmMasterIP"

    Write-Log "Stop Docker"
    Stop-Docker

    Write-Log "Opening firewall ports"
    Open-FirewallPorts

    Write-Log "Write Docker Configuration"
    Write-DockerDaemonJson

    Write-Log "Update Docker restart policy"
    Update-DockerServiceRecoveryPolicy
    
    Write-Log "Start Docker"
    Start-Docker

    #remove-ItemProperty -Path "HKLM:\SYSTEM\CurrentControlSet\Control\Wininit"  Headless
    #Write-Log "shutdown /r /f /t 60"
    #shutdown /r /f /t 60

    Write-Log "Setup Complete"
}
catch
{
    Write-Error $_
}


`)

func swarmInstallContainerhostAndJoinSwarmPs1Bytes() ([]byte, error) {
	return _swarmInstallContainerhostAndJoinSwarmPs1, nil
}

func swarmInstallContainerhostAndJoinSwarmPs1() (*asset, error) {
	bytes, err := swarmInstallContainerhostAndJoinSwarmPs1Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "swarm/Install-ContainerHost-And-Join-Swarm.ps1", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _swarmJoinSwarmmodeClusterPs1 = []byte(`############################################################
# Script adapted from
# https://raw.githubusercontent.com/Microsoft/Virtualization-Documentation/master/windows-server-container-tools/Install-ContainerHost/Install-ContainerHost.ps1

<#
    .NOTES
        Copyright (c) Microsoft Corporation.  All rights reserved.

        Use of this sample source code is subject to the terms of the Microsoft
        license agreement under which you licensed this sample source code. If
        you did not accept the terms of the license agreement, you are not
        authorized to use this sample source code. For the terms of the license,
        please see the license agreement between you and Microsoft or, if applicable,
        see the LICENSE.RTF on your install media or the root of your tools installation.
        THE SAMPLE SOURCE CODE IS PROVIDED "AS IS", WITH NO WARRANTIES.

    .SYNOPSIS
        Installs the prerequisites for creating Windows containers
        Opens TCP ports (80,443,2375,8080) in Windows Firewall.
        Connects Docker to a Swarm Mode master.

    .DESCRIPTION
        Installs the prerequisites for creating Windows containers
        Opens TCP ports (80,443,2375,8080) in Windows Firewall.
        Connects Docker to a Swarm Mode master.

    .PARAMETER SwarmMasterIP
        IP Address of Docker Swarm Mode Master

    .EXAMPLE
        .\Join-SwarmMode-cluster.ps1 -SwarmMasterIP 192.168.255.5

#>
#Requires -Version 5.0

[CmdletBinding(DefaultParameterSetName="Standard")]
param(
    [string]
    [ValidateNotNullOrEmpty()]
    $SwarmMasterIP = "172.16.0.5"
)

$global:DockerServiceName = "Docker"
$global:DockerBinariesURL = "https://acsengine.blob.core.windows.net/swarmm/docker.zip"
$global:DockerExePath = "C:\Program Files\Docker"
$global:IsNewDockerVersion = $false

filter Timestamp {"$(Get-Date -Format o): $_"}

function Write-Log($message)
{
    $msg = $message | Timestamp
    Write-Output $msg
}

function Start-Docker()
{
    Write-Log "Starting $global:DockerServiceName..."
    $startTime = Get-Date
        
    while (-not $dockerReady)
    {
        try
        {
            Start-Service -Name $global:DockerServiceName -ea Stop

            $dockerReady = $true            
        }
        catch
        {
            $timeElapsed = $(Get-Date) - $startTime
            if ($($timeElapsed).TotalMinutes -ge 5)
            {
                Write-Log "Docker Daemon did not start successfully within 5 minutes."
                break
            }

            $errorStr = $_.Exception.Message
            Write-Log "Starting Service failed: $errorStr" 
            Write-Log "sleeping for 10 seconds..."
            Start-Sleep -sec 10
        }
    }
}

function Stop-Docker()
{
    Write-Log "Stopping $global:DockerServiceName..."
    try
    {
        Stop-Service -Name $global:DockerServiceName -ea Stop   
    }
    catch
    {
        Write-Log "Failed to stop Docker"
    }
}

function Expand-ZIPFile($file, $destination)
{
    $shell = new-object -com shell.application
    $zip = $shell.NameSpace($file)
    foreach($item in $zip.items())
    {
        $shell.Namespace($destination).copyhere($item, 0x14)
    }
}

function Install-DockerBinaries()
{
    if( $global:IsNewDockerVersion)
    {
        Write-Log "Skipping installation of new Docker binaries because latest is already installed."
        return
    }

    $currentRetry = 0;
    $success = $false;

    $zipfile = "c:\swarmm.zip"

    do {
        try
        {
            Write-Log "Downloading and installing Docker binaries...."
            Invoke-WebRequest -Uri $global:DockerBinariesURL -OutFile $zipfile
            $success = $true;
            Write-Log "Successfully downloaded Docker binaries. Number of retries: $currentRetry";
        }
        catch [System.Exception]
        {
            $message = 'Exception occurred while trying to download binaries:' + $_.Exception.ToString();
            Write-Log $message;
            if ($currentRetry -gt 5) {
                $message = "Could not download Docker binaries, aborting install. Error: " + $_.Exception.ToString();
                throw $message;
            } else {
                Write-Log "Sleeping before retry number: $currentRetry to download binaries.";
                Start-Sleep -sec 5;
            }
            $currentRetry = $currentRetry + 1;
        }
    } while (!$success);
      
    Write-Log "Expanding zip file at destination: $global:DockerExePath"
    Expand-ZIPFile -File $zipfile -Destination $global:DockerExePath

    Write-Log "Deleting zip file at: $zipfile"
    Remove-Item $zipfile
}

function Update-DockerServiceRecoveryPolicy()
{
    $dockerReady = $false
    $startTime = Get-Date
    
    # wait until the service exists
    while (-not $dockerReady)
    {
        if (Get-Service $global:DockerServiceName -ErrorAction SilentlyContinue)
        {
            $dockerReady = $true
        }
        else 
        {
            $timeElapsed = $(Get-Date) - $startTime
            if ($($timeElapsed).TotalMinutes -ge 5)
            {
                Write-Log "Unable to find service $global:DockerServiceName within 5 minutes."
                break
            }
            Write-Log "failed to find $global:DockerServiceName, sleeping for 5 seconds"
            Start-Sleep -sec 5
        }
    }
    
    Write-Log "Updating docker restart policy, to ensure it restarts on error"
    $services = Get-WMIObject win32_service | Where-Object {$_.name -imatch $global:DockerServiceName}
    foreach ($service in $services)
    {
        sc.exe failure $service.name reset= 86400 actions= restart/5000
    }
}

# Open Windows Firewall Ports Needed
function Open-FirewallPorts()
{
    $tcpports = @(80,443,2375,8080,2377,7946,4789)
    foreach ($tcpport in $tcpports)
    {
        $netsh = "netsh advfirewall firewall add rule name='Open Port $tcpport' dir=in action=allow protocol=TCP localport=$tcpport"
        Write-Log "enabling port with command $netsh"
        Invoke-Expression -Command:$netsh
    }

    $udpports = @(7946,4789)
    foreach ($udpport in $udpports)
    {
        $netsh = "netsh advfirewall firewall add rule name='Open Port $udpport' dir=in action=allow protocol=UDP localport=$udpport"
        Write-Log "enabling port with command $netsh"
        Invoke-Expression -Command:$netsh
    }
}

# Update Docker Config to have cluster-store=consul:// address configured for Swarm cluster.
function Write-DockerDaemonJson()
{
    $dataDir = $env:ProgramData

    # create the target directory
    $targetDir = $dataDir + '\docker\config'
    if(!(Test-Path -Path $targetDir )){
        New-Item -ItemType directory -Path $targetDir
    }

    Write-Log "Delete key file, so that this node is unique to swarm"
    $keyFileName = "$targetDir\key.json"
    Write-Log "Removing $($keyFileName)"
    if (Test-Path $keyFileName) {
      Remove-Item $keyFileName
    }

    Write-Log "Configure Docker Engine to accept incoming connections on port 2375"
    $OutFile = @"
{
    "hosts": ["tcp://0.0.0.0:2375", "npipe://"]
}
"@

    $OutFile | Out-File -encoding ASCII -filepath "$targetDir\daemon.json"
}

function Join-Swarm()
{
    $currentRetry = 0;
    $success = $false;
    $getTokenCommand = "docker -H $($SwarmMasterIP):2375 swarm join-token -q worker"
    $swarmmodetoken;

    do {
        try
        {
            Write-Log "Executing [$getTokenCommand] command...."
            <#& $swarmmodetoken#>
            $swarmmodetoken = Invoke-Expression -Command:$getTokenCommand
            $success = $true;
            Write-Log "Successfully executed [$getTokenCommand] command. Number of entries: $currentRetry. Token: [$swarmmodetoken]";
        }
        catch [System.Exception]
        {
            $message = 'Exception occurred while trying to execute command [$swarmmodetoken]:' + $_.Exception.ToString();
            Write-Log $message;
            if ($currentRetry -gt 120) {
                $message = "Agent couldn't join Swarm, aborting install. Error: " + $_.Exception.ToString();
                throw $message;
            } else {
                Write-Log "Sleeping before $currentRetry retry of [$getTokenCommand] command";
                Start-Sleep -sec 5;
            }
            $currentRetry = $currentRetry + 1;
        }
    } while (!$success);

    $joinSwarmCommand = "docker swarm join --token $($swarmmodetoken) $($SwarmMasterIP):2377"
    Write-Log "Joining Swarm. Command [$joinSwarmCommand]...."
    Invoke-Expression -Command:$joinSwarmCommand
}

function Confirm-DockerVersion()
{
   $dockerServerVersionCmd = "docker version --format '{{.Server.Version}}'"
   Write-Log "Running command: $dockerServerVersionCmd"
   $dockerServerVersion = Invoke-Expression -Command:$dockerServerVersionCmd

   $dockerClientVersionCmd = "docker version --format '{{.Client.Version}}'"
   Write-Log "Running command: $dockerClientVersionCmd"
   $dockerClientVersion = Invoke-Expression -Command:$dockerClientVersionCmd

   Write-Log "Docker Server version: $dockerServerVersion, Docker Client verison: $dockerClientVersion"
   
   $serverVersionData = $dockerServerVersion.Split(".")
   $isNewServerVersion = $false;
   if(($serverVersionData[0] -ge 1) -and ($serverVersionData[1] -ge 13)){
       $isNewServerVersion = $true;
       Write-Log "Setting isNewServerVersion to $isNewServerVersion"
   }

   $clientVersionData = $dockerClientVersion.Split(".")
   $isNewClientVersion = $false;
   if(($clientVersionData[0] -ge 1) -and ($clientVersionData[1] -ge 13)){
       $isNewClientVersion = $true;
       Write-Log "Setting  isNewClientVersion to $isNewClientVersion"   
   }

   if($isNewServerVersion -and $isNewClientVersion)
   {
       $global:IsNewDockerVersion = $true;
       Write-Log "Setting IsNewDockerVersion to $global:IsNewDockerVersion"
   }
}

try
{
    Write-Log "Provisioning $global:DockerServiceName... with Swarm IP $SwarmMasterIP"

    Write-Log "Checking Docker version"
    Confirm-DockerVersion

    Write-Log "Stop Docker"
    Stop-Docker

    Write-Log "Installing Docker binaries"
    Install-DockerBinaries

    Write-Log "Opening firewall ports"
    Open-FirewallPorts

    Write-Log "Write Docker Configuration"
    Write-DockerDaemonJson

    Write-Log "Update Docker restart policy"
    Update-DockerServiceRecoveryPolicy
    
    Write-Log "Start Docker"
    Start-Docker
    
    Write-Log "Join existing Swarm"
    Join-Swarm

    Write-Log "Setup Complete"
}
catch
{
    Write-Error $_
}`)

func swarmJoinSwarmmodeClusterPs1Bytes() ([]byte, error) {
	return _swarmJoinSwarmmodeClusterPs1, nil
}

func swarmJoinSwarmmodeClusterPs1() (*asset, error) {
	bytes, err := swarmJoinSwarmmodeClusterPs1Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "swarm/Join-SwarmMode-cluster.ps1", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _swarmConfigureSwarmClusterSh = []byte(`#!/bin/bash

set -x

echo "starting swarm cluster configuration"
date
ps ax

#############
# Parameters
#############

SWARM_VERSION=${1}
DOCKER_COMPOSE_VERSION=${2}
MASTERCOUNT=${3}
MASTERPREFIX=${4}
MASTERFIRSTADDR=${5}
AZUREUSER=${6}
POSTINSTALLSCRIPTURI=${7}
BASESUBNET=${8}
DOCKERENGINEDOWNLOADREPO=${9}
DOCKERCOMPOSEDOWNLOADURL=${10}
DOCKER_CE_VERSION=17.03.*
VMNAME=` + "`" + `hostname` + "`" + `
VMNUMBER=` + "`" + `echo $VMNAME | sed 's/.*[^0-9]\([0-9]\+\)*$/\1/'` + "`" + `
VMPREFIX=` + "`" + `echo $VMNAME | sed 's/\(.*[^0-9]\)*[0-9]\+$/\1/'` + "`" + `

echo "Master Count: $MASTERCOUNT"
echo "Master Prefix: $MASTERPREFIX"
echo "Master First Addr: $MASTERFIRSTADDR"
echo "vmname: $VMNAME"
echo "VMNUMBER: $VMNUMBER, VMPREFIX: $VMPREFIX"
echo "BASESUBNET: $BASESUBNET"
echo "AZUREUSER: $AZUREUSER"

###################
# Common Functions
###################

ensureAzureNetwork()
{
  # ensure the network works
  networkHealthy=1
  for i in {1..12}; do
    wget -O/dev/null http://bing.com
    if [ $? -eq 0 ]
    then
      # hostname has been found continue
      networkHealthy=0
      echo "the network is healthy"
      break
    fi
    sleep 10
  done
  if [ $networkHealthy -ne 0 ]
  then
    echo "the network is not healthy, aborting install"
    ifconfig
    ip a
    exit 1
  fi
  # ensure the host ip can resolve
  networkHealthy=1
  for i in {1..120}; do
    hostname -i
    if [ $? -eq 0 ]
    then
      # hostname has been found continue
      networkHealthy=0
      echo "the network is healthy"
      break
    fi
    sleep 1
  done
  # attempt to fix hostname, in case dns is not resolving Azure IPs (but can resolve public ips)
  if [ $networkHealthy -ne 0 ]
  then
    HOSTNAME=` + "`" + `hostname` + "`" + `
    HOSTADDR=` + "`" + `ip address show dev eth0 | grep -Eo 'inet (addr:)?([0-9]*\.){3}[0-9]*' | grep -Eo '([0-9]*\.){3}[0-9]*'` + "`" + `
    echo $HOSTADDR $HOSTNAME >> /etc/hosts
    hostname -i
    if [ $? -eq 0 ]
    then
      # hostname has been found continue
      networkHealthy=0
      echo "the network is healthy by updating /etc/hosts"
    fi
  fi
  if [ $networkHealthy -ne 0 ]
  then
    echo "the network is not healthy, cannot resolve ip address, aborting install"
    ifconfig
    ip a
    exit 2
  fi
}
ensureAzureNetwork
HOSTADDR=` + "`" + `hostname -i` + "`" + `

# apply all Canonical security updates during provisioning
/usr/lib/apt/apt.systemd.daily

ismaster ()
{
  if [ "$MASTERPREFIX" == "$VMPREFIX" ]
  then
    return 0
  else
    return 1
  fi
}
if ismaster ; then
  echo "this node is a master"
fi

isagent()
{
  if ismaster ; then
    return 1
  else
    return 0
  fi
}
if isagent ; then
  echo "this node is an agent"
fi

consulstr()
{
  consulargs=""
  for i in ` + "`" + `seq 0 $((MASTERCOUNT-1))` + "`" + ` ;
  do
    MASTEROCTET=` + "`" + `expr $MASTERFIRSTADDR + $i` + "`" + `
    IPADDR="${BASESUBNET}${MASTEROCTET}"

    if [ "$VMNUMBER" -eq "0" ]
    then
      consulargs="${consulargs}-bootstrap-expect $MASTERCOUNT "
    fi
    if [ "$VMNUMBER" -eq "$i" ]
    then
      consulargs="${consulargs}-advertise $IPADDR "
    else
      consulargs="${consulargs}-retry-join $IPADDR "
    fi
  done
  echo $consulargs
}

consulargs=$(consulstr)
MASTER0IPADDR="${BASESUBNET}${MASTERFIRSTADDR}"

######################
# resolve self in DNS
######################

echo "$HOSTADDR $VMNAME" | sudo tee -a /etc/hosts

################
# Install Docker
################

echo "Installing and configuring docker"

# simple general command retry function
retrycmd_if_failure() { for i in 1 2 3 4 5; do $@; [ $? -eq 0  ] && break || sleep 5; done ; }

installDocker()
{
  for i in {1..10}; do
    apt-get install -y apt-transport-https ca-certificates curl software-properties-common
    curl --max-time 60 -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add - 
    add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
    apt-get update
    apt-get install -y docker-ce=${DOCKER_CE_VERSION}
    if [ $? -eq 0 ]
    then
      # hostname has been found continue
      echo "Docker installed successfully"
      break
    fi
    sleep 10
  done
}
time installDocker
sudo usermod -aG docker $AZUREUSER
if isagent ; then
  # Start Docker and listen on :2375 (no auth, but in vnet)
  echo 'DOCKER_OPTS="-H unix:///var/run/docker.sock -H 0.0.0.0:2375 --cluster-store=consul://'$MASTER0IPADDR:8500 --cluster-advertise=$HOSTADDR:2375'"' | sudo tee -a /etc/default/docker
fi

echo "Installing docker compose"
installDockerCompose()
{
  for i in {1..10}; do
    wget --tries 4 --retry-connrefused --waitretry=15 -qO- $DOCKERCOMPOSEDOWNLOADURL/$DOCKER_COMPOSE_VERSION/docker-compose-` + "`" + `uname -s` + "`" + `-` + "`" + `uname -m` + "`" + ` > /usr/local/bin/docker-compose
    if [ $? -eq 0 ]
    then
      # hostname has been found continue
      echo "docker-compose installed successfully"
      break
    fi
    sleep 10
  done
}
time installDockerCompose
chmod +x /usr/local/bin/docker-compose

sudo service docker restart

ensureDocker()
{
  # ensure that docker is healthy
  dockerHealthy=1
  for i in {1..3}; do
    sudo docker info
    if [ $? -eq 0 ]
    then
      # hostname has been found continue
      dockerHealthy=0
      echo "Docker is healthy"
      sudo docker ps -a
      break
    fi
    sleep 10
  done
  if [ $dockerHealthy -ne 0 ]
  then
    echo "Docker is not healthy"
  fi
}
ensureDocker

##############################################
# configure init rules restart all processes
##############################################

if ismaster ; then
  mkdir -p /data/consul
  echo "consul:
  image: \"progrium/consul\"
  command: -server -node $VMNAME $consulargs
  ports:
    - \"8500:8500\"
    - \"8300:8300\"
    - \"8301:8301\"
    - \"8301:8301/udp\"
    - \"8302:8302\"
    - \"8302:8302/udp\"
    - \"8400:8400\"
  volumes:
    - \"/data/consul:/data\"
  restart: \"always\"
swarm:
  image: \"$SWARM_VERSION\"
  command: manage --replication --advertise $HOSTADDR:2375 --discovery-opt kv.path=docker/nodes consul://$MASTER0IPADDR:8500
  ports:
    - \"2375:2375\"
  links:
    - \"consul\"
  volumes:
    - \"/etc/docker:/etc/docker\"
  restart: \"always\"
" > /opt/azure/containers/docker-compose.yml

  pushd /opt/azure/containers/
  docker-compose up -d
  popd
  echo "completed starting docker swarm on the master"
fi

if ismaster ; then
  echo "Having ssh listen to port 2222 as well as 22"
  sudo sed  -i "s/^Port 22$/Port 22\nPort 2222/1" /etc/ssh/sshd_config
fi

if [ $POSTINSTALLSCRIPTURI != "disabled" ]
then
  echo "downloading, and kicking off post install script"
  /bin/bash -c "wget --tries 20 --retry-connrefused --waitretry=15 -qO- $POSTINSTALLSCRIPTURI | nohup /bin/bash >> /var/log/azure/cluster-bootstrap-postinstall.log 2>&1 &"
fi

echo "processes at end of script"
ps ax
date
echo "completed Swarm cluster configuration"

echo "restart system to install any remaining software"
if isagent ; then
  shutdown -r now
else
  # wait 1 minute to restart master
  /bin/bash -c "shutdown -r 1 &"
fi
`)

func swarmConfigureSwarmClusterShBytes() ([]byte, error) {
	return _swarmConfigureSwarmClusterSh, nil
}

func swarmConfigureSwarmClusterSh() (*asset, error) {
	bytes, err := swarmConfigureSwarmClusterShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "swarm/configure-swarm-cluster.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _swarmConfigureSwarmmodeClusterSh = []byte(`#!/bin/bash

###########################################################
# Configure Swarm Mode One Box
#
# This installs the following components
# - Docker
# - Docker Compose
# - Swarm Mode masters
# - Swarm Mode agents
###########################################################

set -x

echo "starting Swarm Mode cluster configuration"
date
ps ax

#############
# Parameters
#############

DOCKER_CE_VERSION=${1}
DOCKER_COMPOSE_VERSION=${2}
MASTERCOUNT=${3}
MASTERPREFIX=${4}
MASTERFIRSTADDR=${5}
AZUREUSER=${6}
POSTINSTALLSCRIPTURI=${7}
BASESUBNET=${8}
DOCKERENGINEDOWNLOADREPO=${9}
DOCKERCOMPOSEDOWNLOADURL=${10}
VMNAME=` + "`" + `hostname` + "`" + `
VMNUMBER=` + "`" + `echo $VMNAME | sed 's/.*[^0-9]\([0-9]\+\)*$/\1/'` + "`" + `
VMPREFIX=` + "`" + `echo $VMNAME | sed 's/\(.*[^0-9]\)*[0-9]\+$/\1/'` + "`" + `
OS="$(. /etc/os-release; echo $ID)"

echo "Master Count: $MASTERCOUNT"
echo "Master Prefix: $MASTERPREFIX"
echo "Master First Addr: $MASTERFIRSTADDR"
echo "vmname: $VMNAME"
echo "VMNUMBER: $VMNUMBER, VMPREFIX: $VMPREFIX"
echo "BASESUBNET: $BASESUBNET"
echo "AZUREUSER: $AZUREUSER"
echo "OS ID: $OS"

###################
# Common Functions
###################

isUbuntu()
{
  if [ "$OS" == "ubuntu" ]
  then
    return 0
  else
    return 1
  fi
}

isRHEL()
{
  if [ "$OS" == "rhel" ]
  then
    return 0
  else
    return 1
  fi
}

ensureAzureNetwork()
{
  # ensure the network works
  networkHealthy=1
  for i in {1..12}; do
    wget -O/dev/null http://bing.com
    if [ $? -eq 0 ]
    then
      # hostname has been found continue
      networkHealthy=0
      echo "the network is healthy"
      break
    fi
    sleep 10
  done
  if [ $networkHealthy -ne 0 ]
  then
    echo "the network is not healthy, aborting install"
    ifconfig
    ip a
    exit 1
  fi
  # ensure the host ip can resolve
  networkHealthy=1
  for i in {1..120}; do
    hostname -i
    if [ $? -eq 0 ]
    then
      # hostname has been found continue
      networkHealthy=0
      echo "the network is healthy"
      break
    fi
    sleep 1
  done
  # attempt to fix hostname, in case dns is not resolving Azure IPs (but can resolve public ips)
  if [ $networkHealthy -ne 0 ]
  then
    HOSTNAME=` + "`" + `hostname` + "`" + `
    HOSTADDR=` + "`" + `ip address show dev eth0 | grep -Eo 'inet (addr:)?([0-9]*\.){3}[0-9]*' | grep -Eo '([0-9]*\.){3}[0-9]*'` + "`" + `
    echo $HOSTADDR $HOSTNAME >> /etc/hosts
    hostname -i
    if [ $? -eq 0 ]
    then
      # hostname has been found continue
      networkHealthy=0
      echo "the network is healthy by updating /etc/hosts"
    fi
  fi
  if [ $networkHealthy -ne 0 ]
  then
    echo "the network is not healthy, cannot resolve ip address, aborting install"
    ifconfig
    ip a
    exit 2
  fi
}
ensureAzureNetwork
HOSTADDR=` + "`" + `hostname -i` + "`" + `

# apply all Canonical security updates during provisioning
/usr/lib/apt/apt.systemd.daily

ismaster ()
{
  if [ "$MASTERPREFIX" == "$VMPREFIX" ]
  then
    return 0
  else
    return 1
  fi
}
if ismaster ; then
  echo "this node is a master"
fi

isagent()
{
  if ismaster ; then
    return 1
  else
    return 0
  fi
}
if isagent ; then
  echo "this node is an agent"
fi

MASTER0IPADDR="${BASESUBNET}${MASTERFIRSTADDR}"

######################
# resolve self in DNS
######################

if [ -z "$(grep "$HOSTADDR $VMNAME" /etc/hosts)" ]; then
    echo "$HOSTADDR $VMNAME" | sudo tee -a /etc/hosts
fi

################
# Install Docker
################

echo "Installing and configuring Docker"

installDockerUbuntu()
{
  for i in {1..10}; do
    apt-get install -y apt-transport-https ca-certificates curl software-properties-common
    curl --max-time 60 -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add - 
    add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
    apt-get update
    apt-get install -y docker-ce=${DOCKER_CE_VERSION}
    if [ $? -eq 0 ]
    then
      systemctl restart docker
      # hostname has been found continue
      echo "Docker installed successfully"
      break
    fi
    sleep 10
  done
}

installDockerRHEL()
{
  for i in {1..10}; do
    yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
    yum makecache fast
    yum -y install docker-ce
    if [ $? -eq 0 ]
    then
      systemctl enable docker
      systemctl start docker
      echo "Docker installed successfully"
      break
    fi
    sleep 10
  done
}

installDocker()
{
  if isUbuntu ; then
    installDockerUbuntu
  elif isRHEL ; then
    installDockerRHEL
  else
    echo "OS not supported, aborting install"
    exit 5
  fi
}

time installDocker

sudo usermod -aG docker $AZUREUSER

echo "Updating Docker daemon options"

updateDockerDaemonOptions()
{
    sudo mkdir -p /etc/systemd/system/docker.service.d
    # Start Docker and listen on :2375 (no auth, but in vnet) and
    # also have it bind to the unix socket at /var/run/docker.sock
    sudo bash -c 'echo "[Service]
    ExecStart=
    ExecStart=/usr/bin/dockerd -H tcp://0.0.0.0:2375 -H unix:///var/run/docker.sock
  " > /etc/systemd/system/docker.service.d/override.conf'
}
time updateDockerDaemonOptions

echo "Installing Docker Compose"
installDockerCompose()
{
  # sudo -i

  for i in {1..10}; do
    wget --tries 4 --retry-connrefused --waitretry=15 -qO- $DOCKERCOMPOSEDOWNLOADURL/$DOCKER_COMPOSE_VERSION/docker-compose-` + "`" + `uname -s` + "`" + `-` + "`" + `uname -m` + "`" + ` > /usr/local/bin/docker-compose
    if [ $? -eq 0 ]
    then
      # hostname has been found continue
      echo "docker-compose installed successfully"
      break
    fi
    sleep 10
  done
}
time installDockerCompose
chmod +x /usr/local/bin/docker-compose

if ismaster && isRHEL ; then
  echo "Opening Docker ports"
  firewall-cmd --add-port=2375/tcp --permanent
  firewall-cmd --add-port=2377/tcp --permanent
  firewall-cmd --reload
fi

echo "Restarting Docker"
sudo systemctl daemon-reload
sudo service docker restart

ensureDocker()
{
  # ensure that docker is healthy
  dockerHealthy=1
  for i in {1..3}; do
    sudo docker info
    if [ $? -eq 0 ]
    then
      # hostname has been found continue
      dockerHealthy=0
      echo "Docker is healthy"
      sudo docker ps -a
      break
    fi
    sleep 10
  done
  if [ $dockerHealthy -ne 0 ]
  then
    echo "Docker is not healthy"
  fi
}
ensureDocker

##############################################
# configure init rules restart all processes
##############################################

if ismaster ; then
    if [ "$HOSTADDR" = "$MASTER0IPADDR" ]; then
          echo "Creating a new Swarm on first master"
          docker swarm init --advertise-addr $(hostname -i):2377 --listen-addr $(hostname -i):2377
    else
        echo "Secondary master attempting to join an existing Swarm"
        swarmmodetoken=""
        swarmmodetokenAcquired=1
        for i in {1..120}; do
            swarmmodetoken=$(docker -H $MASTER0IPADDR:2375 swarm join-token -q manager)
            if [ $? -eq 0 ]; then
                swarmmodetokenAcquired=0
                break
            fi
            sleep 5
        done
        if [ $swarmmodetokenAcquired -ne 0 ]
        then
            echo "Secondary master couldn't connect to Swarm, aborting install"
            exit 3
        fi
        docker swarm join --token $swarmmodetoken $MASTER0IPADDR:2377
    fi
fi

if ismaster ; then
  echo "Having ssh listen to port 2222 as well as 22"
  sudo sed  -i "s/^Port 22$/Port 22\nPort 2222/1" /etc/ssh/sshd_config
fi

if ismaster ; then
  echo "Setting availability of master node: '$VMNAME' to pause"
  docker node update --availability pause $VMNAME
fi

if isagent ; then
    echo "Agent attempting to join an existing Swarm"
    swarmmodetoken=""
    swarmmodetokenAcquired=1
    for i in {1..120}; do
        swarmmodetoken=$(docker -H $MASTER0IPADDR:2375 swarm join-token -q worker)
        if [ $? -eq 0 ]; then
            swarmmodetokenAcquired=0
            break
        fi
        sleep 5
    done
    if [ $swarmmodetokenAcquired -ne 0 ]
    then
        echo "Agent couldn't join Swarm, aborting install"
        exit 4
    fi
    docker swarm join --token $swarmmodetoken $MASTER0IPADDR:2377
fi

if [ $POSTINSTALLSCRIPTURI != "disabled" ]
then
  echo "downloading, and kicking off post install script"
  /bin/bash -c "wget --tries 20 --retry-connrefused --waitretry=15 -qO- $POSTINSTALLSCRIPTURI | nohup /bin/bash >> /var/log/azure/cluster-bootstrap-postinstall.log 2>&1 &"
fi

# mitigation for bug https://bugs.launchpad.net/ubuntu/+source/linux/+bug/1676635
echo 2dd1ce17-079e-403c-b352-a1921ee207ee > /sys/bus/vmbus/drivers/hv_util/unbind
sed -i "13i\echo 2dd1ce17-079e-403c-b352-a1921ee207ee > /sys/bus/vmbus/drivers/hv_util/unbind\n" /etc/rc.local

echo "processes at end of script"
ps ax
date
echo "completed Swarm Mode cluster configuration"

echo "restart system to install any remaining software"
if isagent ; then
  shutdown -r now
else
  # wait 1 minute to restart master
  /bin/bash -c "shutdown -r 1 &"
fi
`)

func swarmConfigureSwarmmodeClusterShBytes() ([]byte, error) {
	return _swarmConfigureSwarmmodeClusterSh, nil
}

func swarmConfigureSwarmmodeClusterSh() (*asset, error) {
	bytes, err := swarmConfigureSwarmmodeClusterShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "swarm/configure-swarmmode-cluster.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _swarmSwarmagentresourcesvmasT = []byte(`    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "copy": {
        "count": "[variables('{{.Name}}Count')]",
        "name": "loop"
      },
      "dependsOn": [
{{if not .IsCustomVNET}}
      "[variables('vnetID')]"
{{end}}
{{if IsPublic .Ports}}
	  ,"[variables('{{.Name}}LbID')]"
{{end}}
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('{{.Name}}VMNamePrefix'), 'nic-', copyIndex(variables('{{.Name}}Offset')))]",
      "properties": {
        "ipConfigurations": [
          {
            "name": "ipConfigNode",
            "properties": {
{{if IsPublic .Ports}}
              "loadBalancerBackendAddressPools": [
		        {
		      	  "id": "[concat('/subscriptions/', subscription().subscriptionId,'/resourceGroups/', resourceGroup().name, '/providers/Microsoft.Network/loadBalancers/', variables('{{.Name}}LbName'), '/backendAddressPools/',variables('{{.Name}}LbBackendPoolName'))]"
		        }
		      ],
{{end}}
              "privateIPAllocationMethod": "Dynamic",
              "subnet": {
                "id": "[variables('{{.Name}}VnetSubnetID')]"
             }
            }
          }
        ]
      },
      "type": "Microsoft.Network/networkInterfaces"
    },
{{if .IsManagedDisks}}
    {
      "apiVersion": "[variables('apiVersionStorageManagedDisks')]",
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}AvailabilitySet')]",
      "properties": {
        "platformFaultDomainCount": 2,
        "platformUpdateDomainCount": 3,
        "managed": "true"
      },
      "type": "Microsoft.Compute/availabilitySets"
    },
{{else if .IsStorageAccount}}
    {
      "apiVersion": "[variables('apiVersionStorage')]",
      "copy": {
        "count": "[variables('{{.Name}}StorageAccountsCount')]",
        "name": "vmLoopNode"
      },
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('storageAccountPrefixes')[mod(add(copyIndex(),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(copyIndex(),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
      "properties": {
        "accountType": "[variables('vmSizesMap')[variables('{{.Name}}VMSize')].storageAccountType]"
      },
      "type": "Microsoft.Storage/storageAccounts"
    },
    {{if .HasDisks}}
        {
          "apiVersion": "[variables('apiVersionStorage')]",
          "copy": {
            "count": "[variables('{{.Name}}StorageAccountsCount')]",
            "name": "datadiskLoop"
          },
          "dependsOn": [
            "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
          ],
          "location": "[variables('location')]",
          "name": "[concat(variables('storageAccountPrefixes')[mod(add(copyIndex(variables('dataStorageAccountPrefixSeed')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(copyIndex(variables('dataStorageAccountPrefixSeed')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}DataAccountName'))]",
          "properties": {
            "accountType": "[variables('vmSizesMap')[variables('{{.Name}}VMSize')].storageAccountType]"
          },
          "type": "Microsoft.Storage/storageAccounts"
        },
    {{end}}
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}AvailabilitySet')]",
      "properties": {},
      "type": "Microsoft.Compute/availabilitySets"
    },
{{end}}
{{if IsPublic .Ports}}
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}IPAddressName')]",
      "properties": {
        "dnsSettings": {
          "domainNameLabel": "[variables('{{.Name}}EndpointDNSNamePrefix')]"
        },
        "publicIPAllocationMethod": "Dynamic"
      },
      "type": "Microsoft.Network/publicIPAddresses"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('{{.Name}}IPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}LbName')]",
      "properties": {
        "backendAddressPools": [
          {
            "name": "[variables('{{.Name}}LbBackendPoolName')]"
          }
        ],
        "frontendIPConfigurations": [
          {
            "name": "[variables('{{.Name}}LbIPConfigName')]",
            "properties": {
              "publicIPAddress": {
                "id": "[resourceId('Microsoft.Network/publicIPAddresses',variables('{{.Name}}IPAddressName'))]"
              }
            }
          }
        ],
        "inboundNatRules": [],
        "loadBalancingRules": [
          {{(GetLBRules .Name .Ports)}}
        ],
        "probes": [
          {{(GetProbes .Ports)}}
        ]
      },
      "type": "Microsoft.Network/loadBalancers"
    },
{{end}}
    {
{{if .IsManagedDisks}}
    "apiVersion": "[variables('apiVersionStorageManagedDisks')]",
{{else}}
    "apiVersion": "[variables('apiVersionDefault')]",
{{end}}
      "copy": {
        "count": "[variables('{{.Name}}Count')]",
        "name": "vmLoopNode"
      },
      "dependsOn": [
{{if .IsStorageAccount}}
        "[concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
  {{if .HasDisks}}
          "[concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('dataStorageAccountPrefixSeed')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('dataStorageAccountPrefixSeed')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}DataAccountName'))]",
  {{end}}
{{end}}
        "[concat('Microsoft.Network/networkInterfaces/', variables('{{.Name}}VMNamePrefix'), 'nic-', copyIndex(variables('{{.Name}}Offset')))]",
        "[concat('Microsoft.Compute/availabilitySets/', variables('{{.Name}}AvailabilitySet'))]"
      ],
      "tags":
      {
        "creationSource" : "[concat('acsengine-', variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')))]"
      },
      "location": "[variables('location')]",
      "name": "[concat(variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')))]",
      "properties": {
        "availabilitySet": {
          "id": "[resourceId('Microsoft.Compute/availabilitySets',variables('{{.Name}}AvailabilitySet'))]"
        },
        "hardwareProfile": {
          "vmSize": "[variables('{{.Name}}VMSize')]"
        },
        "networkProfile": {
          "networkInterfaces": [
            {
              "id": "[resourceId('Microsoft.Network/networkInterfaces',concat(variables('{{.Name}}VMNamePrefix'), 'nic-', copyIndex(variables('{{.Name}}Offset'))))]"
            }
          ]
        },
        "osProfile": {
          "adminUsername": "[variables('adminUsername')]",
          "computername": "[concat(variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')))]",
{{if IsSwarmMode}}
  {{if not .IsRHEL}}
            {{GetAgentSwarmModeCustomData .}} 
  {{end}}
{{else}}
            {{GetAgentSwarmCustomData .}} 
{{end}}
          "linuxConfiguration": {
              "disablePasswordAuthentication": true,
              "ssh": {
                "publicKeys": [
                  {
                    "keyData": "[parameters('sshRSAPublicKey')]",
                    "path": "[variables('sshKeyPath')]"
                  }
                ]
              }
            }
            {{if HasLinuxSecrets}}
              ,
              "secrets": "[variables('linuxProfileSecrets')]"
            {{end}}
        },
        "storageProfile": {
          {{GetDataDisks .}}
          "imageReference": {
            "offer": "[variables('{{.Name}}OSImageOffer')]",
            "publisher": "[variables('{{.Name}}OSImagePublisher')]",
            "sku": "[variables('{{.Name}}OSImageSKU')]",
            "version": "[variables('{{.Name}}OSImageVersion')]"
          }

          ,"osDisk": {
            "caching": "ReadOnly"
            ,"createOption": "FromImage"
{{if .IsStorageAccount}}
            ,"name": "[concat(variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')),'-osdisk')]"
            ,"vhd": {
              "uri": "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')),variables('apiVersionStorage')).primaryEndpoints.blob,'osdisk/', variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')), '-osdisk.vhd')]"
            }
{{end}}
{{if ne .OSDiskSizeGB 0}}
            ,"diskSizeGB": {{.OSDiskSizeGB}}
{{end}}
          }
        }
      },
      "type": "Microsoft.Compute/virtualMachines"
    }
{{if .IsRHEL}}
    ,{
      "apiVersion": "[variables('apiVersionDefault')]",
      "copy": {
        "count": "[variables('{{.Name}}Count')]",
        "name": "vmLoopNode"
      },
      "dependsOn": [
          "[concat('Microsoft.Compute/virtualMachines/', concat(variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset'))))]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')), '/configureagent')]",
      "properties": {
        "publisher": "Microsoft.Azure.Extensions",
        "settings": {
          "commandToExecute": "[variables('agentCustomScript')]",
          "fileUris": [
            "[concat('{{ GetConfigurationScriptRootURL }}', variables('configureClusterScriptFile'))]"
          ]
        },
        "type": "CustomScript",
        "typeHandlerVersion": "2.0"
      },
      "type": "Microsoft.Compute/virtualMachines/extensions"
    }
{{end}}
`)

func swarmSwarmagentresourcesvmasTBytes() ([]byte, error) {
	return _swarmSwarmagentresourcesvmasT, nil
}

func swarmSwarmagentresourcesvmasT() (*asset, error) {
	bytes, err := swarmSwarmagentresourcesvmasTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "swarm/swarmagentresourcesvmas.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _swarmSwarmagentresourcesvmssT = []byte(`{{if .IsStorageAccount}}
    {
      "apiVersion": "[variables('apiVersionStorage')]",
      "copy": {
        "count": "[variables('{{.Name}}StorageAccountsCount')]",
        "name": "vmLoopNode"
      },
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('storageAccountPrefixes')[mod(add(copyIndex(),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(copyIndex(),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
      "properties": {
        "accountType": "[variables('vmSizesMap')[variables('{{.Name}}VMSize')].storageAccountType]"
      },
      "type": "Microsoft.Storage/storageAccounts"
    },
{{end}}
{{if IsPublic .Ports}}
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}IPAddressName')]",
      "properties": {
        "dnsSettings": {
          "domainNameLabel": "[variables('{{.Name}}EndpointDNSNamePrefix')]"
        },
        "publicIPAllocationMethod": "Dynamic"
      },
      "type": "Microsoft.Network/publicIPAddresses"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('{{.Name}}IPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}LbName')]",
      "properties": {
        "backendAddressPools": [
          {
            "name": "[variables('{{.Name}}LbBackendPoolName')]"
          }
        ],
        "frontendIPConfigurations": [
          {
            "name": "[variables('{{.Name}}LbIPConfigName')]",
            "properties": {
              "publicIPAddress": {
                "id": "[resourceId('Microsoft.Network/publicIPAddresses',variables('{{.Name}}IPAddressName'))]"
              }
            }
          }
        ],
        "inboundNatRules": [],
        "loadBalancingRules": [
          {{(GetLBRules .Name .Ports)}}
        ],
        "probes": [
          {{(GetProbes .Ports)}}
        ]
      },
      "type": "Microsoft.Network/loadBalancers"
    },
{{end}}
    {
{{if .IsManagedDisks}}
      "apiVersion": "[variables('apiVersionStorageManagedDisks')]",
{{else}}
      "apiVersion": "[variables('apiVersionDefault')]",
{{end}}
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
{{if .IsStorageAccount}}
        ,"[concat('Microsoft.Storage/storageAccounts/', variables('storageAccountPrefixes')[mod(add(0,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(0,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
        "[concat('Microsoft.Storage/storageAccounts/', variables('storageAccountPrefixes')[mod(add(1,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(1,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
		"[concat('Microsoft.Storage/storageAccounts/', variables('storageAccountPrefixes')[mod(add(2,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(2,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
        "[concat('Microsoft.Storage/storageAccounts/', variables('storageAccountPrefixes')[mod(add(3,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(3,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
        "[concat('Microsoft.Storage/storageAccounts/', variables('storageAccountPrefixes')[mod(add(4,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(4,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]"
{{end}}
{{if not .IsCustomVNET}}
      ,"[variables('vnetID')]"
{{end}}
{{if IsPublic .Ports}}
       ,"[variables('{{.Name}}LbID')]"
{{end}}
      ],
      "tags":
      {
        "creationSource" : "[concat('acsengine-', variables('{{.Name}}VMNamePrefix'), '-vmss')]"
      },
      "location": "[variables('location')]",
      "name": "[concat(variables('{{.Name}}VMNamePrefix'), '-vmss')]",
      "properties": {
        "upgradePolicy": {
          "mode": "Automatic"
        },
        "virtualMachineProfile": {
          "networkProfile": {
            "networkInterfaceConfigurations": [
              {
                "name": "nic",
                "properties": {
                  "ipConfigurations": [
                    {
                      "name": "nicipconfig",
                      "properties": {
{{if IsPublic .Ports}}
                        "loadBalancerBackendAddressPools": [
                          {
                            "id": "[concat(variables('{{.Name}}LbID'), '/backendAddressPools/', variables('{{.Name}}LbBackendPoolName'))]"
                          }
                        ],
{{end}}
                        "subnet": {
                          "id": "[variables('{{.Name}}VnetSubnetID')]"
                        }
                      }
                    }
                  ],
                  "primary": "true"
                }
              }
            ]
          },
          "osProfile": {
            "adminUsername": "[variables('adminUsername')]",
            "computerNamePrefix": "[variables('{{.Name}}VMNamePrefix')]",
{{if IsSwarmMode}}
  {{if not .IsRHEL}}
            {{GetAgentSwarmModeCustomData .}}
  {{end}}
{{else}}
            {{GetAgentSwarmCustomData .}}
{{end}}
            "linuxConfiguration": {
              "disablePasswordAuthentication": true,
              "ssh": {
                "publicKeys": [
                  {
                    "keyData": "[parameters('sshRSAPublicKey')]",
                    "path": "[variables('sshKeyPath')]"
                  }
                ]
              }
            }
            {{if HasLinuxSecrets}}
              ,
              "secrets": "[variables('linuxProfileSecrets')]"
            {{end}}
          },
          "storageProfile": {
            "imageReference": {
              "offer": "[variables('{{.Name}}OSImageOffer')]",
              "publisher": "[variables('{{.Name}}OSImagePublisher')]",
              "sku": "[variables('{{.Name}}OSImageSKU')]",
              "version": "[variables('{{.Name}}OSImageVersion')]"
            },
            {{GetDataDisks .}}
            "osDisk": {
              "caching": "ReadWrite"
              ,"createOption": "FromImage"
{{if .IsStorageAccount}}
              ,"name": "vmssosdisk"
              ,"vhdContainers": [
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(0,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(0,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')), variables('apiVersionStorage') ).primaryEndpoints.blob, 'osdisk')]",
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(1,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(1,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')), variables('apiVersionStorage')).primaryEndpoints.blob, 'osdisk')]",
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(2,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(2,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')), variables('apiVersionStorage')).primaryEndpoints.blob, 'osdisk')]",
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(3,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(3,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')), variables('apiVersionStorage')).primaryEndpoints.blob, 'osdisk')]",
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(4,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(4,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')), variables('apiVersionStorage')).primaryEndpoints.blob, 'osdisk')]"
              ]
{{end}}
{{if ne .OSDiskSizeGB 0}}
            ,"diskSizeGB": {{.OSDiskSizeGB}}
{{end}}
            }
          }
{{if .IsRHEL}}
          ,"extensionProfile": {
            "extensions": [
              {
                "name": "configure{{.Name}}",
                "properties": {
                  "publisher": "Microsoft.Azure.Extensions",
                  "settings": {
                    "commandToExecute": "[variables('agentCustomScript')]",
                    "fileUris": [
                      "[concat('{{ GetConfigurationScriptRootURL }}', variables('configureClusterScriptFile'))]"
                    ]
                  },
                  "type": "CustomScript",
                  "typeHandlerVersion": "2.0"
                }
              }
            ]
          }
{{end}}
        }
      },
      "sku": {
        "capacity": "[variables('{{.Name}}Count')]",
        "name": "[variables('{{.Name}}VMSize')]",
        "tier": "[variables('{{.Name}}VMSizeTier')]"
      },
      "type": "Microsoft.Compute/virtualMachineScaleSets"
    }
`)

func swarmSwarmagentresourcesvmssTBytes() ([]byte, error) {
	return _swarmSwarmagentresourcesvmssT, nil
}

func swarmSwarmagentresourcesvmssT() (*asset, error) {
	bytes, err := swarmSwarmagentresourcesvmssTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "swarm/swarmagentresourcesvmss.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _swarmSwarmagentvarsT = []byte(`{{if not .IsRHEL}}
    "{{.Name}}RunCmd": "[concat('runcmd:\n {{GetSwarmAgentPreprovisionExtensionCommands .}} \n-  [ /bin/bash, /opt/azure/containers/install-cluster.sh ]\n\n')]",
    "{{.Name}}RunCmdFile": "[concat(' -  content: |\n        #!/bin/bash\n        ','sudo mkdir -p /var/log/azure\n        ',variables('agentCustomScript'),'\n    path: /opt/azure/containers/install-cluster.sh\n    permissions: \"0744\"\n')]",
{{end}}
{{if IsSwarmMode }}
    "{{.Name}}OSImageOffer": {{GetAgentOSImageOffer .}},
    "{{.Name}}OSImagePublisher": {{GetAgentOSImagePublisher .}},
    "{{.Name}}OSImageSKU": {{GetAgentOSImageSKU .}},
    "{{.Name}}OSImageVersion": {{GetAgentOSImageVersion .}},
{{else}}
    "{{.Name}}OSImageOffer": "[variables('osImageOffer')]",
    "{{.Name}}OSImagePublisher": "[variables('osImagePublisher')]",
    "{{.Name}}OSImageSKU": "[variables('osImageSKU')]",
    "{{.Name}}OSImageVersion": "[variables('osImageVersion')]",
{{end}}
    "{{.Name}}Count": "[parameters('{{.Name}}Count')]",
    "{{.Name}}VMNamePrefix": "[concat(variables('orchestratorName'), '-{{.Name}}-', variables('nameSuffix'))]",
    "{{.Name}}VMSize": "[parameters('{{.Name}}VMSize')]",
    "{{.Name}}VMSizeTier": "[split(parameters('{{.Name}}VMSize'),'_')[0]]",
{{if .IsAvailabilitySets}}
    {{if .IsStorageAccount}}
    "{{.Name}}StorageAccountsCount": "[add(div(variables('{{.Name}}Count'), variables('maxVMsPerStorageAccount')), mod(add(mod(variables('{{.Name}}Count'), variables('maxVMsPerStorageAccount')),2), add(mod(variables('{{.Name}}Count'), variables('maxVMsPerStorageAccount')),1)))]",
    "{{.Name}}StorageAccountOffset": "[mul(variables('maxStorageAccountsPerAgent'),variables('{{.Name}}Index'))]",
    {{end}}
    "{{.Name}}AvailabilitySet": "[concat('{{.Name}}-availabilitySet-', variables('nameSuffix'))]",
    "{{.Name}}Offset": "[parameters('{{.Name}}Offset')]",
{{else}}
    {{if .IsStorageAccount}}
    "{{.Name}}StorageAccountsCount": "[variables('maxStorageAccountsPerAgent')]",
    {{end}}
{{end}}
{{if .IsCustomVNET}}
    "{{.Name}}VnetSubnetID": "[parameters('{{.Name}}VnetSubnetID')]",
{{else}}
    "{{.Name}}Subnet": "[parameters('{{.Name}}Subnet')]",
    "{{.Name}}SubnetName": "[concat(variables('orchestratorName'), '-{{.Name}}subnet')]",
    "{{.Name}}VnetSubnetID": "[concat(variables('vnetID'),'/subnets/',variables('{{.Name}}SubnetName'))]",
{{end}}
{{if IsPublic .Ports}}
    "{{.Name}}EndpointDNSNamePrefix": "[tolower(parameters('{{.Name}}EndpointDNSNamePrefix'))]",
    "{{.Name}}IPAddressName": "[concat(variables('orchestratorName'), '-agent-ip-', variables('{{.Name}}EndpointDNSNamePrefix'), '-', variables('nameSuffix'))]",
    "{{.Name}}LbBackendPoolName": "[concat(variables('orchestratorName'), '-{{.Name}}-', variables('nameSuffix'))]",
    "{{.Name}}LbID": "[resourceId('Microsoft.Network/loadBalancers',variables('{{.Name}}LbName'))]",
    "{{.Name}}LbIPConfigID": "[concat(variables('{{.Name}}LbID'),'/frontendIPConfigurations/', variables('{{.Name}}LbIPConfigName'))]",
    "{{.Name}}LbIPConfigName": "[concat(variables('orchestratorName'), '-{{.Name}}-', variables('nameSuffix'))]",
    "{{.Name}}LbName": "[concat(variables('orchestratorName'), '-{{.Name}}-', variables('nameSuffix'))]",
     {{if .IsWindows}}
        "{{.Name}}WindowsRDPNatRangeStart": 3389,
        "{{.Name}}WindowsRDPEndRangeStop": "[add(variables('{{.Name}}WindowsRDPNatRangeStart'), add(variables('{{.Name}}Count'),variables('{{.Name}}Count')))]",
    {{end}}
 {{end}}
`)

func swarmSwarmagentvarsTBytes() ([]byte, error) {
	return _swarmSwarmagentvarsT, nil
}

func swarmSwarmagentvarsT() (*asset, error) {
	bytes, err := swarmSwarmagentvarsTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "swarm/swarmagentvars.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _swarmSwarmbaseT = []byte(`{
  "$schema": "https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
  "contentVersion": "1.0.0.0",
  "parameters": {
    {{range .AgentPoolProfiles}}{{template "agentparams.t" .}},{{end}}
    {{if .HasWindows}}
      {{template "windowsparams.t"}},
    {{end}}
    {{template "masterparams.t" .}}
    {{template "swarm/swarmparams.t" .}}
  },
  "variables": {
    {{range $index, $agent := .AgentPoolProfiles}}
        "{{.Name}}Index": {{$index}},
        {{template "swarm/swarmagentvars.t" .}}
        {{if .IsStorageAccount}}
          "{{.Name}}StorageAccountOffset": "[mul(variables('maxStorageAccountsPerAgent'),{{$index}})]",
          "{{.Name}}AccountName": "[concat(variables('storageAccountBaseName'), 'agnt{{$index}}')]",
          {{if .HasDisks}}
            "{{.Name}}DataAccountName": "[concat(variables('storageAccountBaseName'), 'data{{$index}}')]",
          {{end}}
        {{end}}
    {{end}}

    {{template "swarm/swarmmastervars.t" .}}
  },
  "resources": [
    {{range .AgentPoolProfiles}}
      {{if .IsWindows}}
        {{if .IsAvailabilitySets}}
          {{template "swarm/swarmwinagentresourcesvmas.t" .}},
        {{else}}
          {{template "swarm/swarmwinagentresourcesvmss.t" .}},
        {{end}}
      {{else}}
        {{if .IsAvailabilitySets}}
          {{template "swarm/swarmagentresourcesvmas.t" .}},
        {{else}}
          {{template "swarm/swarmagentresourcesvmss.t" .}},
        {{end}}
      {{end}}
    {{end}}
    {{template "swarm/swarmmasterresources.t" .}}
  ],
  "outputs": {
    {{range .AgentPoolProfiles}}{{template "agentoutputs.t" .}}
    {{end}}
    {{template "masteroutputs.t" .}}
  }
}
`)

func swarmSwarmbaseTBytes() ([]byte, error) {
	return _swarmSwarmbaseT, nil
}

func swarmSwarmbaseT() (*asset, error) {
	bytes, err := swarmSwarmbaseTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "swarm/swarmbase.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _swarmSwarmmasterresourcesT = []byte(`{{if not .MasterProfile.IsCustomVNET}}
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('virtualNetworkName')]",
      "properties": {
        "addressSpace": {
          "addressPrefixes": [
            {{GetVNETAddressPrefixes}}
          ]
        },
        "subnets": [
          {{GetVNETSubnets false}}
        ]
      },
      "type": "Microsoft.Network/virtualNetworks"
    },
{{end}}
{{if .MasterProfile.IsManagedDisks}}
    {
      "apiVersion": "[variables('apiVersionStorageManagedDisks')]",
      "location": "[variables('location')]",
      "name": "[variables('masterAvailabilitySet')]",
      "properties": {
        "platformFaultDomainCount": 2,
        "platformUpdateDomainCount": 3,
        "managed": "true"
      },
      "type": "Microsoft.Compute/availabilitySets"
    },
{{else if .MasterProfile.IsStorageAccount}}
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('masterAvailabilitySet')]",
      "properties": {},
      "type": "Microsoft.Compute/availabilitySets"
    },
    {
      "apiVersion": "[variables('apiVersionStorage')]",
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[variables('masterStorageAccountName')]",
      "properties": {
        "accountType": "[variables('vmSizesMap')[variables('masterVMSize')].storageAccountType]"
      },
      "type": "Microsoft.Storage/storageAccounts"
    },
{{end}}
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('masterPublicIPAddressName')]",
      "properties": {
        "dnsSettings": {
          "domainNameLabel": "[variables('masterEndpointDNSNamePrefix')]"
        },
        "publicIPAllocationMethod": "Dynamic"
      },
      "type": "Microsoft.Network/publicIPAddresses"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[variables('masterLbName')]",
      "properties": {
        "backendAddressPools": [
          {
            "name": "[variables('masterLbBackendPoolName')]"
          }
        ],
        "frontendIPConfigurations": [
          {
            "name": "[variables('masterLbIPConfigName')]",
            "properties": {
              "publicIPAddress": {
                "id": "[resourceId('Microsoft.Network/publicIPAddresses',variables('masterPublicIPAddressName'))]"
              }
            }
          }
        ]
      },
      "type": "Microsoft.Network/loadBalancers"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "copy": {
        "count": "[variables('masterCount')]",
        "name": "masterLbLoopNode"
      },
      "dependsOn": [
        "[variables('masterLbID')]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('masterLbName'), '/', 'SSH-', variables('masterVMNamePrefix'), copyIndex())]",
      "properties": {
        "backendPort": 22,
        "enableFloatingIP": false,
        "frontendIPConfiguration": {
          "id": "[variables('masterLbIPConfigID')]"
        },
        "frontendPort": "[copyIndex(2200)]",
        "protocol": "Tcp"
      },
      "type": "Microsoft.Network/loadBalancers/inboundNatRules"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "dependsOn": [
        "[variables('masterLbID')]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('masterSshPort22InboundNatRuleNamePrefix'), '0')]",
      "properties": {
        "backendPort": 2222,
        "enableFloatingIP": false,
        "frontendIPConfiguration": {
          "id": "[variables('masterLbIPConfigID')]"
        },
        "frontendPort": "22",
        "protocol": "Tcp"
      },
      "type": "Microsoft.Network/loadBalancers/inboundNatRules"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "copy": {
        "count": "[variables('masterCount')]",
        "name": "nicLoopNode"
      },
      "dependsOn": [
{{if not .MasterProfile.IsCustomVNET}}
        "[variables('vnetID')]",
{{end}}
        "[variables('masterLbID')]",
        "[concat(variables('masterSshPort22InboundNatRuleIdPrefix'),'0')]",
        "[concat(variables('masterSshInboundNatRuleIdPrefix'),copyIndex())]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('masterVMNamePrefix'), 'nic-', copyIndex())]",
      "properties": {
        "ipConfigurations": [
          {
            "name": "ipConfigNode",
            "properties": {
              "loadBalancerBackendAddressPools": [
                {
                  "id": "[concat(variables('masterLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"
                }
              ],
              "loadBalancerInboundNatRules": "[variables('masterLbInboundNatRules')[copyIndex()]]",
              "privateIPAddress": "[concat(variables('masterFirstAddrPrefix'), copyIndex(int(variables('masterFirstAddrOctet4'))))]",
              "privateIPAllocationMethod": "Static",
              "subnet": {
                "id": "[variables('masterVnetSubnetID')]"
              }
            }
          }
        ]
      },
      "type": "Microsoft.Network/networkInterfaces"
    },
    {
{{if .MasterProfile.IsManagedDisks}}
    "apiVersion": "[variables('apiVersionStorageManagedDisks')]",
{{else}}
    "apiVersion": "[variables('apiVersionDefault')]",
{{end}}
      "copy": {
        "count": "[variables('masterCount')]",
        "name": "vmLoopNode"
      },
      "dependsOn": [
        "[concat('Microsoft.Network/networkInterfaces/', variables('masterVMNamePrefix'), 'nic-', copyIndex())]",
        "[concat('Microsoft.Compute/availabilitySets/',variables('masterAvailabilitySet'))]"
{{if .MasterProfile.IsStorageAccount}}
        ,"[variables('masterStorageAccountName')]"
{{end}}
      ],
      "tags":
      {
        "creationSource" : "[concat('acsengine-', variables('masterVMNamePrefix'), copyIndex())]"
      },
      "location": "[variables('location')]",
      "name": "[concat(variables('masterVMNamePrefix'), copyIndex())]",
      "properties": {
        "availabilitySet": {
          "id": "[resourceId('Microsoft.Compute/availabilitySets',variables('masterAvailabilitySet'))]"
        },
        "hardwareProfile": {
          "vmSize": "[variables('masterVMSize')]"
        },
        "networkProfile": {
          "networkInterfaces": [
            {
              "id": "[resourceId('Microsoft.Network/networkInterfaces',concat(variables('masterVMNamePrefix'), 'nic-', copyIndex()))]"
            }
          ]
        },
        "osProfile": {
          "adminUsername": "[variables('adminUsername')]",
          "computername": "[concat(variables('masterVMNamePrefix'), copyIndex())]",
          {{if .OrchestratorProfile.IsSwarmMode}}
            {{if not .MasterProfile.IsRHEL}}
              {{GetMasterSwarmModeCustomData}}
            {{end}}
          {{else}}
            {{GetMasterSwarmCustomData}}
          {{end}}
          "linuxConfiguration": {
            "disablePasswordAuthentication": true,
            "ssh": {
                "publicKeys": [
                    {
                        "keyData": "[variables('sshRSAPublicKey')]",
                        "path": "[variables('sshKeyPath')]"
                    }
                ]
            }
          }
          {{if .LinuxProfile.HasSecrets}}
          ,
          "secrets": "[variables('linuxProfileSecrets')]"
          {{end}}
        },
        "storageProfile": {
          "imageReference": {
            {{if .OrchestratorProfile.IsSwarmMode}}
            "offer": "[variables('masterOSImageOffer')]",
            "publisher": "[variables('masterOSImagePublisher')]",
            "sku": "[variables('masterOSImageSKU')]",
            "version": "[variables('masterOSImageVersion')]"
            {{else}}
            "offer": "[variables('osImageOffer')]",
            "publisher": "[variables('osImagePublisher')]",
            "sku": "[variables('osImageSKU')]",
            "version": "[variables('osImageVersion')]"
            {{end}}
          },
          "osDisk": {
            "caching": "ReadWrite"
            ,"createOption": "FromImage"
{{if .MasterProfile.IsStorageAccount}}
            ,"name": "[concat(variables('masterVMNamePrefix'), copyIndex(),'-osdisk')]"
            ,"vhd": {
              "uri": "[concat(reference(concat('Microsoft.Storage/storageAccounts/', variables('masterStorageAccountName')), variables('apiVersionStorage')).primaryEndpoints.blob, 'vhds/', variables('masterVMNamePrefix'), copyIndex(), '-osdisk.vhd')]"
            }
{{end}}
{{if ne .MasterProfile.OSDiskSizeGB 0}}
            ,"diskSizeGB": {{.MasterProfile.OSDiskSizeGB}}
{{end}}
          }
        }
      },
      "type": "Microsoft.Compute/virtualMachines"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "copy": {
        "count": "[variables('masterCount')]",
        "name": "vmLoopNode"
      },
      "dependsOn": [
          "[concat('Microsoft.Compute/virtualMachines/', concat(variables('masterVMNamePrefix'), copyIndex()))]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('masterVMNamePrefix'), copyIndex(), '/configuremaster')]",
      "properties": {
        "publisher": "Microsoft.Azure.Extensions",
        "settings": {
          "commandToExecute": "[variables('masterCustomScript')]",
          "fileUris": [
{{if .MasterProfile.IsRHEL}}
            "[concat('{{ GetConfigurationScriptRootURL }}', variables('configureClusterScriptFile'))]"
{{end}}
          ]
        },
        "type": "CustomScript",
        "typeHandlerVersion": "2.0"
      },
      "type": "Microsoft.Compute/virtualMachines/extensions"
    }
`)

func swarmSwarmmasterresourcesTBytes() ([]byte, error) {
	return _swarmSwarmmasterresourcesT, nil
}

func swarmSwarmmasterresourcesT() (*asset, error) {
	bytes, err := swarmSwarmmasterresourcesTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "swarm/swarmmasterresources.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _swarmSwarmmastervarsT = []byte(`    "adminUsername": "[parameters('linuxAdminUsername')]",
    "maxVMsPerPool": 100,
    "apiVersionDefault": "2016-03-30",
{{if .OrchestratorProfile.IsSwarmMode}}
    "configureClusterScriptFile": "configure-swarmmode-cluster.sh",
{{else}}
    "configureClusterScriptFile": "configure-swarm-cluster.sh",
{{end}}
{{if .MasterProfile.IsRHEL}}
    "agentCustomScript": "[concat('/usr/bin/nohup /bin/bash -c \"/bin/bash ',variables('configureClusterScriptFile'), ' ',variables('clusterInstallParameters'),' >> /var/log/azure/cluster-bootstrap.log 2>&1 &\" &')]",
{{else}}
    "agentCustomScript": "[concat('/usr/bin/nohup /bin/bash -c \"/bin/bash /opt/azure/containers/',variables('configureClusterScriptFile'), ' ',variables('clusterInstallParameters'),' >> /var/log/azure/cluster-bootstrap.log 2>&1 &\" &')]",
{{end}}
    "agentMaxVMs": 100,
    "clusterInstallParameters": "[concat(variables('orchestratorVersion'), ' ',variables('dockerComposeVersion'), ' ',variables('masterCount'), ' ',variables('masterVMNamePrefix'), ' ',variables('masterFirstAddrOctet4'), ' ',variables('adminUsername'),' ',variables('postInstallScriptURI'),' ',variables('masterFirstAddrPrefix'),' ', parameters('dockerEngineDownloadRepo'), ' ', parameters('dockerComposeDownloadURL'))]",
{{if .LinuxProfile.HasSecrets}}
    "linuxProfileSecrets" :
      [
          {{range  $vIndex, $vault := .LinuxProfile.Secrets}}
            {{if $vIndex}} , {{end}}
              {
                "sourceVault":{
                  "id":"[parameters('linuxKeyVaultID{{$vIndex}}')]"
                },
                "vaultCertificates":[
                {{range $cIndex, $cert := $vault.VaultCertificates}}
                  {{if $cIndex}} , {{end}}
                  {
                    "certificateUrl" :"[parameters('linuxKeyVaultID{{$vIndex}}CertificateURL{{$cIndex}}')]"
                  }
                {{end}}
                ]
              }
        {{end}}
      ],
{{end}}
    "masterAvailabilitySet": "[concat(variables('orchestratorName'), '-master-availabilitySet-', variables('nameSuffix'))]",
    "masterCount": {{.MasterProfile.Count}},
{{if .MasterProfile.IsRHEL}}
    "masterCustomScript": "[concat('/bin/bash -c \"/bin/bash ',variables('configureClusterScriptFile'), ' ',variables('clusterInstallParameters'),' >> /var/log/azure/cluster-bootstrap.log 2>&1\"')]",
{{else}}
    "masterCustomScript": "[concat('/bin/bash -c \"/bin/bash /opt/azure/containers/',variables('configureClusterScriptFile'), ' ',variables('clusterInstallParameters'),' >> /var/log/azure/cluster-bootstrap.log 2>&1\"')]",
{{end}}
    "masterEndpointDNSNamePrefix": "[tolower(parameters('masterEndpointDNSNamePrefix'))]",
    "masterLbBackendPoolName": "[concat(variables('orchestratorName'), '-master-pool-', variables('nameSuffix'))]",
    "masterLbID": "[resourceId('Microsoft.Network/loadBalancers',variables('masterLbName'))]",
    "masterLbIPConfigID": "[concat(variables('masterLbID'),'/frontendIPConfigurations/', variables('masterLbIPConfigName'))]",
    "masterLbIPConfigName": "[concat(variables('orchestratorName'), '-master-lbFrontEnd-', variables('nameSuffix'))]",
    "masterLbName": "[concat(variables('orchestratorName'), '-master-lb-', variables('nameSuffix'))]",
    "masterPublicIPAddressName": "[concat(variables('orchestratorName'), '-master-ip-', variables('masterEndpointDNSNamePrefix'), '-', variables('nameSuffix'))]",
{{if .MasterProfile.IsCustomVNET}}
    "masterVnetSubnetID": "[parameters('masterVnetSubnetID')]",
{{else}}
    "masterSubnet": "[parameters('masterSubnet')]",
    "masterSubnetName": "[concat(variables('orchestratorName'), '-masterSubnet')]",
    "vnetID": "[resourceId('Microsoft.Network/virtualNetworks',variables('virtualNetworkName'))]",
    "masterVnetSubnetID": "[concat(variables('vnetID'),'/subnets/',variables('masterSubnetName'))]",
    "virtualNetworkName": "[concat(variables('orchestratorName'), '-vnet-', variables('nameSuffix'))]",
{{end}}
    "masterFirstAddrOctets": "[split(parameters('firstConsecutiveStaticIP'),'.')]",
    "masterFirstAddrOctet4": "[variables('masterFirstAddrOctets')[3]]",
    "masterFirstAddrPrefix": "[concat(variables('masterFirstAddrOctets')[0],'.',variables('masterFirstAddrOctets')[1],'.',variables('masterFirstAddrOctets')[2],'.')]",
    "masterVMNamePrefix": "[concat(variables('orchestratorName'), '-master-', variables('nameSuffix'), '-')]",
    "masterVMSize": "[parameters('masterVMSize')]",
    "nameSuffix": "[parameters('nameSuffix')]",
    "masterSshInboundNatRuleIdPrefix": "[concat(variables('masterLbID'),'/inboundNatRules/SSH-',variables('masterVMNamePrefix'))]",
    "masterSshPort22InboundNatRuleNamePrefix": "[concat(variables('masterLbName'),'/SSHPort22-',variables('masterVMNamePrefix'))]",
    "masterSshPort22InboundNatRuleIdPrefix": "[concat(variables('masterLbID'),'/inboundNatRules/SSHPort22-',variables('masterVMNamePrefix'))]",
     "masterLbInboundNatRules":[
      [
        {
          "id": "[concat(variables('masterSshInboundNatRuleIdPrefix'),'0')]"
        },
        {
          "id": "[concat(variables('masterSshPort22InboundNatRuleIdPrefix'),'0')]"
        }
      ],
      [
        {
          "id": "[concat(variables('masterSshInboundNatRuleIdPrefix'),'1')]"
        }
      ],
      [
        {
          "id": "[concat(variables('masterSshInboundNatRuleIdPrefix'),'2')]"
        }
      ],
      [
        {
          "id": "[concat(variables('masterSshInboundNatRuleIdPrefix'),'3')]"
        }
      ],
      [
        {
          "id": "[concat(variables('masterSshInboundNatRuleIdPrefix'),'4')]"
        }
      ]
    ],
{{if .OrchestratorProfile.IsSwarmMode}}
    "orchestratorName": "swarmm",
    "masterOSImageOffer": {{GetMasterOSImageOffer}},
    "masterOSImagePublisher": {{GetMasterOSImagePublisher}},
    "masterOSImageSKU": {{GetMasterOSImageSKU}},
    "masterOSImageVersion": {{GetMasterOSImageVersion}},
    {{GetSwarmModeVersions}}
{{else}}
    "orchestratorName": "swarm",
    "osImageOffer": "[parameters('osImageOffer')]",
    "osImagePublisher": "[parameters('osImagePublisher')]",
    "osImageSKU": "14.04.5-LTS",
    "osImageVersion": "14.04.201706190",
    {{getSwarmVersions}}
{{end}}
    "locations": [
         "[resourceGroup().location]",
         "[parameters('location')]"
    ],
    "location": "[variables('locations')[mod(add(2,length(parameters('location'))),add(1,length(parameters('location'))))]]",
    "postInstallScriptURI": "disabled",
    "sshKeyPath": "[concat('/home/', variables('adminUsername'), '/.ssh/authorized_keys')]",
{{if .HasStorageAccountDisks}}
    "apiVersionStorage": "2015-06-15",
    "maxVMsPerStorageAccount": 20,
    "maxStorageAccountsPerAgent": "[div(variables('maxVMsPerPool'),variables('maxVMsPerStorageAccount'))]",
    "dataStorageAccountPrefixSeed": 97,
    "storageAccountPrefixes": [ "0", "6", "c", "i", "o", "u", "1", "7", "d", "j", "p", "v", "2", "8", "e", "k", "q", "w", "3", "9", "f", "l", "r", "x", "4", "a", "g", "m", "s", "y", "5", "b", "h", "n", "t", "z" ],
    "storageAccountPrefixesCount": "[length(variables('storageAccountPrefixes'))]",
    "vmsPerStorageAccount": 20,
    "storageAccountBaseName": "[uniqueString(concat(variables('masterEndpointDNSNamePrefix'),variables('location')))]",
    {{GetSizeMap}},
{{else}}
    "storageAccountPrefixes": [],
    "storageAccountBaseName": "",
{{end}}
{{if .HasManagedDisks}}
    "apiVersionStorageManagedDisks": "2016-04-30-preview",
{{end}}
{{if .MasterProfile.IsStorageAccount}}
    "masterStorageAccountName": "[concat(variables('storageAccountBaseName'), '0')]",
{{end}}
    "sshRSAPublicKey": "[parameters('sshRSAPublicKey')]"
{{if .HasWindows}}
    ,"windowsAdminUsername": "[parameters('windowsAdminUsername')]",
    "windowsAdminPassword": "[parameters('windowsAdminPassword')]",
    "agentWindowsPublisher": "[parameters('agentWindowsPublisher')]",
    "agentWindowsOffer": "[parameters('agentWindowsOffer')]",
    "agentWindowsSku": "[parameters('agentWindowsSku')]",
    "agentWindowsVersion": "[parameters('agentWindowsVersion')]",
    "singleQuote": "'",
    "windowsCustomScriptArguments": "[concat('$arguments = ', variables('singleQuote'),'-SwarmMasterIP ', variables('masterFirstAddrPrefix'), variables('masterFirstAddrOctet4'), variables('singleQuote'), ' ; ')]",
    "windowsCustomScriptSuffix": " $inputFile = '%SYSTEMDRIVE%\\AzureData\\CustomData.bin' ; $outputFile = '%SYSTEMDRIVE%\\AzureData\\CustomDataSetupScript.ps1' ; $inputStream = New-Object System.IO.FileStream $inputFile, ([IO.FileMode]::Open), ([IO.FileAccess]::Read), ([IO.FileShare]::Read) ; $sr = New-Object System.IO.StreamReader(New-Object System.IO.Compression.GZipStream($inputStream, [System.IO.Compression.CompressionMode]::Decompress)) ; $sr.ReadToEnd() | Out-File($outputFile) ; Invoke-Expression('{0} {1}' -f $outputFile, $arguments) ; ",
    "windowsCustomScript": "[concat('powershell.exe -ExecutionPolicy Unrestricted -command \"', variables('windowsCustomScriptArguments'), variables('windowsCustomScriptSuffix'), '\" > %SYSTEMDRIVE%\\AzureData\\CustomDataSetupScript.log 2>&1')]",
    "agentWindowsBackendPort": 3389
    {{if .WindowsProfile.HasSecrets}}
    ,
    "windowsProfileSecrets" :
      [
          {{range  $vIndex, $vault := .LinuxProfile.Secrets}}
            {{if $vIndex}} , {{end}}
              {
                "sourceVault":{
                  "id":"[parameters('windowsKeyVaultID{{$vIndex}}')]"
                },
                "vaultCertificates":[
                {{range $cIndex, $cert := $vault.VaultCertificates}}
                  {{if $cIndex}} , {{end}}
                  {
                    "certificateUrl" :"[parameters('windowsKeyVaultID{{$vIndex}}CertificateURL{{$cIndex}}')]",
                    "certificateStore" :"[parameters('windowsKeyVaultID{{$vIndex}}CertificateStore{{$cIndex}}')]"
                  }
                {{end}}
                ]
              }
        {{end}}
      ]
      {{end}}
{{end}}

`)

func swarmSwarmmastervarsTBytes() ([]byte, error) {
	return _swarmSwarmmastervarsT, nil
}

func swarmSwarmmastervarsT() (*asset, error) {
	bytes, err := swarmSwarmmastervarsTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "swarm/swarmmastervars.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _swarmSwarmparamsT = []byte(`    ,
    "dockerEngineDownloadRepo": {
      "defaultValue": "",
      "metadata": {
        "description": "Docker engine download repo."
      },
      "type": "string"
    },
    "dockerComposeDownloadURL": {
      "defaultValue": "",
      "metadata": {
        "description": "Docker compose download URL."
      },
      "type": "string"
    }`)

func swarmSwarmparamsTBytes() ([]byte, error) {
	return _swarmSwarmparamsT, nil
}

func swarmSwarmparamsT() (*asset, error) {
	bytes, err := swarmSwarmparamsTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "swarm/swarmparams.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _swarmSwarmwinagentresourcesvmasT = []byte(`    {
      "apiVersion": "[variables('apiVersionDefault')]", 
      "copy": {
        "count": "[sub(variables('{{.Name}}Count'), variables('{{.Name}}Offset'))]", 
        "name": "loop"
      }, 
      "dependsOn": [
{{if not .IsCustomVNET}}
      "[variables('vnetID')]"
{{end}}
{{if IsPublic .Ports}}
	  ,"[variables('{{.Name}}LbID')]"
{{end}}
      ], 
      "location": "[variables('location')]", 
      "name": "[concat(variables('{{.Name}}VMNamePrefix'), 'nic-', copyIndex(variables('{{.Name}}Offset')))]",
      "properties": {
        "ipConfigurations": [
          {
            "name": "ipConfigNode", 
            "properties": {
{{if IsPublic .Ports}}
              "loadBalancerBackendAddressPools": [
                {
                  "id": "[concat('/subscriptions/', subscription().subscriptionId,'/resourceGroups/', resourceGroup().name, '/providers/Microsoft.Network/loadBalancers/', variables('{{.Name}}LbName'), '/backendAddressPools/',variables('{{.Name}}LbBackendPoolName'))]"
                }
              ],
              "loadBalancerInboundNatPools": [
                {
                  "id": "[concat(variables('{{.Name}}LbID'), '/inboundNatPools/', 'RDP-', variables('{{.Name}}VMNamePrefix'))]"
                }
              ],
{{end}}  
              "privateIPAllocationMethod": "Dynamic", 
              "subnet": {
                "id": "[variables('{{.Name}}VnetSubnetID')]"
             }
            }
          }
        ]
      }, 
      "type": "Microsoft.Network/networkInterfaces"
    },
{{if .IsManagedDisks}}
    {
      "apiVersion": "[variables('apiVersionStorageManagedDisks')]", 
      "location": "[variables('location')]", 
      "name": "[variables('{{.Name}}AvailabilitySet')]", 
      "properties": { 
        "platformFaultDomainCount": 2, 
        "platformUpdateDomainCount": 3,
        "managed": "true"
      },
      "type": "Microsoft.Compute/availabilitySets"
    },
{{else if .IsStorageAccount}}
    {
      "apiVersion": "[variables('apiVersionStorage')]", 
      "copy": {
        "count": "[variables('{{.Name}}StorageAccountsCount')]", 
        "name": "vmLoopNode"
      }, 
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
      ], 
      "location": "[variables('location')]", 
      "name": "[concat(variables('storageAccountPrefixes')[mod(add(copyIndex(),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(copyIndex(),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]", 
      "properties": {
        "accountType": "[variables('vmSizesMap')[variables('{{.Name}}VMSize')].storageAccountType]"
      }, 
      "type": "Microsoft.Storage/storageAccounts"
    },
  {{if .HasDisks}}
      {
        "apiVersion": "[variables('apiVersionStorage')]", 
        "copy": {
          "count": "[variables('{{.Name}}StorageAccountsCount')]", 
          "name": "datadiskLoop"
        }, 
        "dependsOn": [
          "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
        ], 
        "location": "[variables('location')]", 
        "name": "[concat(variables('storageAccountPrefixes')[mod(add(copyIndex(variables('dataStorageAccountPrefixSeed')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(copyIndex(variables('dataStorageAccountPrefixSeed')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}DataAccountName'))]", 
        "properties": {
          "accountType": "[variables('vmSizesMap')[variables('{{.Name}}VMSize')].storageAccountType]"
        }, 
        "type": "Microsoft.Storage/storageAccounts"
      }, 
  {{end}}
    {
      "apiVersion": "[variables('apiVersionDefault')]", 
      "location": "[variables('location')]", 
      "name": "[variables('{{.Name}}AvailabilitySet')]", 
      "properties": {}, 
      "type": "Microsoft.Compute/availabilitySets"
    },
{{end}}
{{if IsPublic .Ports}}
    {
      "apiVersion": "[variables('apiVersionDefault')]", 
      "location": "[variables('location')]", 
      "name": "[variables('{{.Name}}IPAddressName')]", 
      "properties": {
        "dnsSettings": {
          "domainNameLabel": "[variables('{{.Name}}EndpointDNSNamePrefix')]"
        }, 
        "publicIPAllocationMethod": "Dynamic"
      }, 
      "type": "Microsoft.Network/publicIPAddresses"
    }, 
    {
      "apiVersion": "[variables('apiVersionDefault')]", 
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('{{.Name}}IPAddressName'))]"
      ], 
      "location": "[variables('location')]", 
      "name": "[variables('{{.Name}}LbName')]", 
      "properties": {
        "backendAddressPools": [
          {
            "name": "[variables('{{.Name}}LbBackendPoolName')]"
          }
        ], 
        "frontendIPConfigurations": [
          {
            "name": "[variables('{{.Name}}LbIPConfigName')]", 
            "properties": {
              "publicIPAddress": {
                "id": "[resourceId('Microsoft.Network/publicIPAddresses',variables('{{.Name}}IPAddressName'))]"
              }
            }
          }
        ],
        "inboundNatPools": [
          {
            "name": "[concat('RDP-', variables('{{.Name}}VMNamePrefix'))]",
            "properties": {
              "frontendIPConfiguration": {
                "id": "[variables('{{.Name}}LbIPConfigID')]"
              },
              "protocol": "Tcp",
              "frontendPortRangeStart": "[variables('{{.Name}}WindowsRDPNatRangeStart')]",
              "frontendPortRangeEnd": "[variables('{{.Name}}WindowsRDPEndRangeStop')]",
              "backendPort": "[variables('agentWindowsBackendPort')]"
            }
          }
        ], 
        "loadBalancingRules": [
          {{(GetLBRules .Name .Ports)}}
        ], 
        "probes": [
          {{(GetProbes .Ports)}}
        ]
      }, 
      "type": "Microsoft.Network/loadBalancers"
    }, 
{{end}}
    {
{{if .IsManagedDisks}}
    "apiVersion": "[variables('apiVersionStorageManagedDisks')]",
{{else}}
    "apiVersion": "[variables('apiVersionDefault')]",
{{end}}
      "copy": {
        "count": "[sub(variables('{{.Name}}Count'), variables('{{.Name}}Offset'))]", 
        "name": "vmLoopNode"
      }, 
      "dependsOn": [
{{if .IsStorageAccount}}
        "[concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
  {{if .HasDisks}}
          "[concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('dataStorageAccountPrefixSeed')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('dataStorageAccountPrefixSeed')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}DataAccountName'))]",
  {{end}}
{{end}}
        "[concat('Microsoft.Network/networkInterfaces/', variables('{{.Name}}VMNamePrefix'), 'nic-', copyIndex(variables('{{.Name}}Offset')))]", 
        "[concat('Microsoft.Compute/availabilitySets/', variables('{{.Name}}AvailabilitySet'))]"
      ],
      "tags":
      {
        "creationSource" : "[concat('acsengine-', variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')))]"
      },
      "location": "[variables('location')]",  
      "name": "[concat(variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')))]",
      "properties": {
        "availabilitySet": {
          "id": "[resourceId('Microsoft.Compute/availabilitySets',variables('{{.Name}}AvailabilitySet'))]"
        }, 
        "hardwareProfile": {
          "vmSize": "[variables('{{.Name}}VMSize')]"
        }, 
        "networkProfile": {
          "networkInterfaces": [
            {
              "id": "[resourceId('Microsoft.Network/networkInterfaces',concat(variables('{{.Name}}VMNamePrefix'), 'nic-', copyIndex(variables('{{.Name}}Offset'))))]"
            }
          ]
        }, 
        "osProfile": {
          "computername": "[concat(substring(variables('nameSuffix'), 0, 5), 'acs', copyIndex(variables('{{.Name}}Offset')), add(900,variables('{{.Name}}Index')))]",
          "adminUsername": "[variables('windowsAdminUsername')]",
          "adminPassword": "[variables('windowsAdminPassword')]",
          {{if IsSwarmMode}}
            {{GetWinAgentSwarmModeCustomData}}           
          {{else}}
            {{GetWinAgentSwarmCustomData}}
          {{end}}
          {{if HasWindowsSecrets}}
              ,
              "secrets": "[variables('windowsProfileSecrets')]"
          {{end}}
        }, 
        "storageProfile": {
          {{GetDataDisks .}}
          "imageReference": {
            "publisher": "[variables('agentWindowsPublisher')]",
            "offer": "[variables('agentWindowsOffer')]",
            "sku": "[variables('agentWindowsSku')]",
            "version": "[variables('agentWindowsVersion')]"
          }
          ,"osDisk": {
            "caching": "ReadOnly"
            ,"createOption": "FromImage"
{{if .IsStorageAccount}}
            ,"name": "[concat(variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')),'-osdisk')]"
            ,"vhd": {
              "uri": "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(div(copyIndex(variables('{{.Name}}Offset')),variables('maxVMsPerStorageAccount')),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')),variables('apiVersionStorage')).primaryEndpoints.blob,'osdisk/', variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')), '-osdisk.vhd')]"
            }
{{end}}
{{if ne .OSDiskSizeGB 0}}
            ,"diskSizeGB": {{.OSDiskSizeGB}}
{{end}}
          }
        }
      }, 
      "type": "Microsoft.Compute/virtualMachines"
    }, 
    {
      "apiVersion": "[variables('apiVersionDefault')]", 
      "copy": {
        "count": "[sub(variables('{{.Name}}Count'), variables('{{.Name}}Offset'))]", 
        "name": "vmLoopNode"
      }, 
      "dependsOn": [
        "[concat('Microsoft.Compute/virtualMachines/', variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')))]"
      ], 
      "location": "[variables('location')]", 
      "name": "[concat(variables('{{.Name}}VMNamePrefix'), copyIndex(variables('{{.Name}}Offset')), '/cse')]",
      "properties": {
        "publisher": "Microsoft.Compute",
        "type": "CustomScriptExtension",
        "typeHandlerVersion": "1.8",
        "autoUpgradeMinorVersion": true,
        "settings": {
          "commandToExecute": "[variables('windowsCustomScript')]"
        }
      }, 
      "type": "Microsoft.Compute/virtualMachines/extensions"
    }
`)

func swarmSwarmwinagentresourcesvmasTBytes() ([]byte, error) {
	return _swarmSwarmwinagentresourcesvmasT, nil
}

func swarmSwarmwinagentresourcesvmasT() (*asset, error) {
	bytes, err := swarmSwarmwinagentresourcesvmasTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "swarm/swarmwinagentresourcesvmas.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _swarmSwarmwinagentresourcesvmssT = []byte(`{{if .IsStorageAccount}}
    {
      "apiVersion": "[variables('apiVersionStorage')]",
      "copy": {
        "count": "[variables('{{.Name}}StorageAccountsCount')]",
        "name": "vmLoopNode"
      },
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('storageAccountPrefixes')[mod(add(copyIndex(),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(copyIndex(),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
      "properties": {
        "accountType": "[variables('vmSizesMap')[variables('{{.Name}}VMSize')].storageAccountType]"
      },
      "type": "Microsoft.Storage/storageAccounts"
    },
{{end}}
{{if IsPublic .Ports}}
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}IPAddressName')]",
      "properties": {
        "dnsSettings": {
          "domainNameLabel": "[variables('{{.Name}}EndpointDNSNamePrefix')]"
        },
        "publicIPAllocationMethod": "Dynamic"
      },
      "type": "Microsoft.Network/publicIPAddresses"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('{{.Name}}IPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}LbName')]",
      "properties": {
        "backendAddressPools": [
          {
            "name": "[variables('{{.Name}}LbBackendPoolName')]"
          }
        ],
        "frontendIPConfigurations": [
          {
            "name": "[variables('{{.Name}}LbIPConfigName')]",
            "properties": {
              "publicIPAddress": {
                "id": "[resourceId('Microsoft.Network/publicIPAddresses',variables('{{.Name}}IPAddressName'))]"
              }
            }
          }
        ],
        "inboundNatRules": [],
        "loadBalancingRules": [
          {{(GetLBRules .Name .Ports)}}
        ],
        "probes": [
          {{(GetProbes .Ports)}}
        ],
        "inboundNatPools": [
          {
            "name": "[concat('RDP-', variables('{{.Name}}VMNamePrefix'))]",
            "properties": {
              "frontendIPConfiguration": {
                "id": "[variables('{{.Name}}LbIPConfigID')]"
              },
              "protocol": "Tcp",
              "frontendPortRangeStart": "[variables('{{.Name}}WindowsRDPNatRangeStart')]",
              "frontendPortRangeEnd": "[variables('{{.Name}}WindowsRDPEndRangeStop')]",
              "backendPort": "[variables('agentWindowsBackendPort')]"
            }
          }
        ]
      },
      "type": "Microsoft.Network/loadBalancers"
    },
{{end}}
    {
{{if .IsManagedDisks}}
      "apiVersion": "[variables('apiVersionStorageManagedDisks')]",
{{else}}
      "apiVersion": "[variables('apiVersionDefault')]",
{{end}}
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
{{if .IsStorageAccount}}
        ,"[concat('Microsoft.Storage/storageAccounts/', variables('storageAccountPrefixes')[mod(add(0,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(0,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
        "[concat('Microsoft.Storage/storageAccounts/', variables('storageAccountPrefixes')[mod(add(1,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(1,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
		    "[concat('Microsoft.Storage/storageAccounts/', variables('storageAccountPrefixes')[mod(add(2,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(2,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
        "[concat('Microsoft.Storage/storageAccounts/', variables('storageAccountPrefixes')[mod(add(3,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(3,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
        "[concat('Microsoft.Storage/storageAccounts/', variables('storageAccountPrefixes')[mod(add(4,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(4,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]"
{{end}}
{{if not .IsCustomVNET}}
      ,"[variables('vnetID')]"
{{end}}
{{if IsPublic .Ports}}
       ,"[variables('{{.Name}}LbID')]"
{{end}}
      ],
      "tags":
      {
        "creationSource" : "[concat('acsengine-', variables('{{.Name}}VMNamePrefix'), '-vmss')]"
      },
      "location": "[variables('location')]",
      "name": "[concat(variables('{{.Name}}VMNamePrefix'), '-vmss')]",
      "properties": {
        "upgradePolicy": {
          "mode": "Automatic"
        },
        "virtualMachineProfile": {
          "networkProfile": {
            "networkInterfaceConfigurations": [
              {
                "name": "nic",
                "properties": {
                  "ipConfigurations": [
                    {
                      "name": "nicipconfig",
                      "properties": {
{{if IsPublic .Ports}}
                        "loadBalancerBackendAddressPools": [
                          {
                            "id": "[concat(variables('{{.Name}}LbID'), '/backendAddressPools/', variables('{{.Name}}LbBackendPoolName'))]"
                          }
                        ],
                        "loadBalancerInboundNatPools": [
                          {
                            "id": "[concat(variables('{{.Name}}LbID'), '/inboundNatPools/', 'RDP-', variables('{{.Name}}VMNamePrefix'))]"
                          }
                        ],
{{end}}
                        "subnet": {
                          "id": "[variables('{{.Name}}VnetSubnetID')]"
                        }
                      }
                    }
                  ],
                  "primary": "true"
                }
              }
            ]
          },
          "osProfile": {
            "computerNamePrefix": "[concat(substring(variables('nameSuffix'), 0, 5), 'acs')]",
            "adminUsername": "[variables('windowsAdminUsername')]",
            "adminPassword": "[variables('windowsAdminPassword')]",
            {{if IsSwarmMode}}
              {{GetWinAgentSwarmModeCustomData}}
            {{else}}
              {{GetWinAgentSwarmCustomData}}
            {{end}}
            {{if HasWindowsSecrets}}
              ,
              "secrets": "[variables('windowsProfileSecrets')]"
            {{end}}
          },
          "storageProfile": {
            "imageReference": {
              "publisher": "[variables('agentWindowsPublisher')]",
              "offer": "[variables('agentWindowsOffer')]",
              "sku": "[variables('agentWindowsSku')]",
              "version": "[variables('agentWindowsVersion')]"
            },
            "osDisk": {
              "caching": "ReadWrite"
              ,"createOption": "FromImage"
{{if .IsStorageAccount}}
              ,"name": "vmssosdisk"
              ,"vhdContainers": [
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(0,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(0,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')), variables('apiVersionStorage') ).primaryEndpoints.blob, 'osdisk')]",
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(1,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(1,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')), variables('apiVersionStorage')).primaryEndpoints.blob, 'osdisk')]",
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(2,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(2,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')), variables('apiVersionStorage')).primaryEndpoints.blob, 'osdisk')]",
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(3,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(3,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')), variables('apiVersionStorage')).primaryEndpoints.blob, 'osdisk')]",
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(4,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(4,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')), variables('apiVersionStorage')).primaryEndpoints.blob, 'osdisk')]"
              ]
{{end}}
{{if ne .OSDiskSizeGB 0}}
            ,"diskSizeGB": {{.OSDiskSizeGB}}
{{end}}
            }
          },
          "extensionProfile": {
            "extensions": [
              {
                "name": "vmssCustomScriptExtension",
                "properties": {
                  "publisher": "Microsoft.Compute",
                  "type": "CustomScriptExtension",
                  "typeHandlerVersion": "1.8",
                  "autoUpgradeMinorVersion": true,
                  "settings": {
                    "commandToExecute": "[variables('windowsCustomScript')]"
                  }
                }
              }
            ]
          }
        }
      },
      "sku": {
        "capacity": "[variables('{{.Name}}Count')]",
        "name": "[variables('{{.Name}}VMSize')]",
        "tier": "[variables('{{.Name}}VMSizeTier')]"
      },
      "type": "Microsoft.Compute/virtualMachineScaleSets"
    }
`)

func swarmSwarmwinagentresourcesvmssTBytes() ([]byte, error) {
	return _swarmSwarmwinagentresourcesvmssT, nil
}

func swarmSwarmwinagentresourcesvmssT() (*asset, error) {
	bytes, err := swarmSwarmwinagentresourcesvmssTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "swarm/swarmwinagentresourcesvmss.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _windowsparamsT = []byte(` {{if IsKubernetes}}
    "kubeBinariesSASURL": {
      "metadata": {
        "description": "The download url for kubernetes windows binaries package that is created by scripts/build-windows-k8s.sh"
      },
      "type": "string"
    },
    "windowsKubeBinariesURL": {
      "metadata": {
        "description": "The download url for kubernetes windows binaries produce by Kubernetes. This contains only the node binaries (example: https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG-1.11.md#node-binaries-1)"
      },
      "type": "string"
    },
    "kubeBinariesVersion": {
      "metadata": {
        "description": "Kubernetes windows binaries version"
      },
      "type": "string"
    },
    "kubeServiceCidr": {
      "metadata": {
        "description": "Kubernetes service address space"
      },
      "type": "string"
    },
    "windowsTelemetryGUID": {
      "metadata": {
        "description": "The GUID to set in windows agent to collect telemetry data."
      },
      "type": "string"
    },
 {{end}}
    "windowsAdminUsername": {
      "type": "string",
      "metadata": {
        "description": "User name for the Windows Swarm Agent Virtual Machines (Password Only Supported)."
      }
    },
    "windowsAdminPassword": {
      "type": "securestring",
      "metadata": {
        "description": "Password for the Windows Swarm Agent Virtual Machines."
      }
    },
    "agentWindowsImageName": {
      "defaultValue": "",
      "type": "string",
      "metadata": {
        "description": "Image name when specifying a Windows image reference."
      }
    },
    "agentWindowsImageResourceGroup": {
      "defaultValue": "",
      "type": "string",
      "metadata": {
        "description": "Resource group when specifying a Windows image reference."
      }
    },
    "agentWindowsVersion": {
      "defaultValue": "latest",
      "metadata": {
        "description": "Version of the Windows Server OS image to use for the agent virtual machines."
      },
      "type": "string"
    },
    "agentWindowsSourceUrl": {
      "defaultValue": "",
      "metadata": {
        "description": "The source of the generalized blob which will be used to create a custom windows image for the agent virtual machines."
      },
      "type": "string"
    },
    "agentWindowsPublisher": {
      "defaultValue": "MicrosoftWindowsServer",
      "metadata": {
        "description": "The publisher of windows image for the agent virtual machines."
      },
      "type": "string"
    },
    "agentWindowsOffer": {
      "defaultValue": "WindowsServerSemiAnnual",
      "metadata": {
        "description": "The offer of windows image for the agent virtual machines."
      },
      "type": "string"
    },
    "agentWindowsSku": {
      "defaultValue": "Datacenter-Core-1809-with-Containers-smalldisk",
      "metadata": {
        "description": "The SKU of windows image for the agent virtual machines."
      },
      "type": "string"
    },
    "windowsDockerVersion": {
      "defaultValue": "18.09.2",
      "metadata": {
        "description": "The version of Docker to be installed on Windows Nodes"
      },
      "type": "string"
    }
`)

func windowsparamsTBytes() ([]byte, error) {
	return _windowsparamsT, nil
}

func windowsparamsT() (*asset, error) {
	bytes, err := windowsparamsTBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "windowsparams.t", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"agentoutputs.t":                                                  agentoutputsT,
	"agentparams.t":                                                   agentparamsT,
	"dcos/bstrap/bootstrapcustomdata.yml":                             dcosBstrapBootstrapcustomdataYml,
	"dcos/bstrap/bootstrapparams.t":                                   dcosBstrapBootstrapparamsT,
	"dcos/bstrap/bootstrapprovision.sh":                               dcosBstrapBootstrapprovisionSh,
	"dcos/bstrap/bootstrapresources.t":                                dcosBstrapBootstrapresourcesT,
	"dcos/bstrap/bootstrapvars.t":                                     dcosBstrapBootstrapvarsT,
	"dcos/bstrap/dcos1.11.0.customdata.t":                             dcosBstrapDcos1110CustomdataT,
	"dcos/bstrap/dcos1.11.2.customdata.t":                             dcosBstrapDcos1112CustomdataT,
	"dcos/bstrap/dcosbase.t":                                          dcosBstrapDcosbaseT,
	"dcos/bstrap/dcosmasterresources.t":                               dcosBstrapDcosmasterresourcesT,
	"dcos/bstrap/dcosmastervars.t":                                    dcosBstrapDcosmastervarsT,
	"dcos/bstrap/dcosprovision.sh":                                    dcosBstrapDcosprovisionSh,
	"dcos/dcosWindowsAgentResourcesVmas.t":                            dcosDcoswindowsagentresourcesvmasT,
	"dcos/dcosWindowsAgentResourcesVmss.t":                            dcosDcoswindowsagentresourcesvmssT,
	"dcos/dcosWindowsProvision.ps1":                                   dcosDcoswindowsprovisionPs1,
	"dcos/dcosagentresourcesvmas.t":                                   dcosDcosagentresourcesvmasT,
	"dcos/dcosagentresourcesvmss.t":                                   dcosDcosagentresourcesvmssT,
	"dcos/dcosagentvars.t":                                            dcosDcosagentvarsT,
	"dcos/dcosbase.t":                                                 dcosDcosbaseT,
	"dcos/dcoscustomdata110.t":                                        dcosDcoscustomdata110T,
	"dcos/dcoscustomdata184.t":                                        dcosDcoscustomdata184T,
	"dcos/dcoscustomdata187.t":                                        dcosDcoscustomdata187T,
	"dcos/dcoscustomdata188.t":                                        dcosDcoscustomdata188T,
	"dcos/dcoscustomdata190.t":                                        dcosDcoscustomdata190T,
	"dcos/dcoscustomdata198.t":                                        dcosDcoscustomdata198T,
	"dcos/dcosmasterresources.t":                                      dcosDcosmasterresourcesT,
	"dcos/dcosmastervars.t":                                           dcosDcosmastervarsT,
	"dcos/dcosparams.t":                                               dcosDcosparamsT,
	"dcos/dcosprovision.sh":                                           dcosDcosprovisionSh,
	"dcos/dcosprovisionsource.sh":                                     dcosDcosprovisionsourceSh,
	"iaasoutputs.t":                                                   iaasoutputsT,
	"k8s/addons/1.10/kubernetesmasteraddons-kube-dns-deployment.yaml": k8sAddons110KubernetesmasteraddonsKubeDnsDeploymentYaml,
	"k8s/addons/1.15/kubernetesmasteraddons-azure-cloud-provider-deployment.yaml": k8sAddons115KubernetesmasteraddonsAzureCloudProviderDeploymentYaml,
	"k8s/addons/1.15/kubernetesmasteraddons-pod-security-policy.yaml":             k8sAddons115KubernetesmasteraddonsPodSecurityPolicyYaml,
	"k8s/addons/1.16/kubernetesmaster-audit-policy.yaml":                          k8sAddons116KubernetesmasterAuditPolicyYaml,
	"k8s/addons/1.16/kubernetesmasteraddons-azure-cloud-provider-deployment.yaml": k8sAddons116KubernetesmasteraddonsAzureCloudProviderDeploymentYaml,
	"k8s/addons/1.16/kubernetesmasteraddons-cilium-daemonset.yaml":                k8sAddons116KubernetesmasteraddonsCiliumDaemonsetYaml,
	"k8s/addons/1.16/kubernetesmasteraddons-flannel-daemonset.yaml":               k8sAddons116KubernetesmasteraddonsFlannelDaemonsetYaml,
	"k8s/addons/1.16/kubernetesmasteraddons-kube-dns-deployment.yaml":             k8sAddons116KubernetesmasteraddonsKubeDnsDeploymentYaml,
	"k8s/addons/1.16/kubernetesmasteraddons-kube-proxy-daemonset.yaml":            k8sAddons116KubernetesmasteraddonsKubeProxyDaemonsetYaml,
	"k8s/addons/1.16/kubernetesmasteraddons-pod-security-policy.yaml":             k8sAddons116KubernetesmasteraddonsPodSecurityPolicyYaml,
	"k8s/addons/1.6/kubernetesmasteraddons-kubernetes-dashboard-deployment.yaml":  k8sAddons16KubernetesmasteraddonsKubernetesDashboardDeploymentYaml,
	"k8s/addons/1.7/kubernetesmasteraddons-kube-dns-deployment.yaml":              k8sAddons17KubernetesmasteraddonsKubeDnsDeploymentYaml,
	"k8s/addons/1.7/kubernetesmasteraddons-kubernetes-dashboard-deployment.yaml":  k8sAddons17KubernetesmasteraddonsKubernetesDashboardDeploymentYaml,
	"k8s/addons/1.8/kubernetesmasteraddons-kube-dns-deployment.yaml":              k8sAddons18KubernetesmasteraddonsKubeDnsDeploymentYaml,
	"k8s/addons/1.8/kubernetesmasteraddons-kubernetes-dashboard-deployment.yaml":  k8sAddons18KubernetesmasteraddonsKubernetesDashboardDeploymentYaml,
	"k8s/addons/1.9/kubernetesmasteraddons-kube-dns-deployment.yaml":              k8sAddons19KubernetesmasteraddonsKubeDnsDeploymentYaml,
	"k8s/addons/coredns.yaml":                                                       k8sAddonsCorednsYaml,
	"k8s/addons/kubernetesmaster-audit-policy.yaml":                                 k8sAddonsKubernetesmasterAuditPolicyYaml,
	"k8s/addons/kubernetesmasteraddons-aad-default-admin-group-rbac.yaml":           k8sAddonsKubernetesmasteraddonsAadDefaultAdminGroupRbacYaml,
	"k8s/addons/kubernetesmasteraddons-azure-cloud-provider-deployment.yaml":        k8sAddonsKubernetesmasteraddonsAzureCloudProviderDeploymentYaml,
	"k8s/addons/kubernetesmasteraddons-cilium-daemonset.yaml":                       k8sAddonsKubernetesmasteraddonsCiliumDaemonsetYaml,
	"k8s/addons/kubernetesmasteraddons-elb-svc.yaml":                                k8sAddonsKubernetesmasteraddonsElbSvcYaml,
	"k8s/addons/kubernetesmasteraddons-flannel-daemonset.yaml":                      k8sAddonsKubernetesmasteraddonsFlannelDaemonsetYaml,
	"k8s/addons/kubernetesmasteraddons-kube-dns-deployment.yaml":                    k8sAddonsKubernetesmasteraddonsKubeDnsDeploymentYaml,
	"k8s/addons/kubernetesmasteraddons-kube-proxy-daemonset.yaml":                   k8sAddonsKubernetesmasteraddonsKubeProxyDaemonsetYaml,
	"k8s/addons/kubernetesmasteraddons-managed-azure-storage-classes-custom.yaml":   k8sAddonsKubernetesmasteraddonsManagedAzureStorageClassesCustomYaml,
	"k8s/addons/kubernetesmasteraddons-managed-azure-storage-classes.yaml":          k8sAddonsKubernetesmasteraddonsManagedAzureStorageClassesYaml,
	"k8s/addons/kubernetesmasteraddons-pod-security-policy.yaml":                    k8sAddonsKubernetesmasteraddonsPodSecurityPolicyYaml,
	"k8s/addons/kubernetesmasteraddons-scheduled-maintenance-deployment.yaml":       k8sAddonsKubernetesmasteraddonsScheduledMaintenanceDeploymentYaml,
	"k8s/addons/kubernetesmasteraddons-unmanaged-azure-storage-classes-custom.yaml": k8sAddonsKubernetesmasteraddonsUnmanagedAzureStorageClassesCustomYaml,
	"k8s/addons/kubernetesmasteraddons-unmanaged-azure-storage-classes.yaml":        k8sAddonsKubernetesmasteraddonsUnmanagedAzureStorageClassesYaml,
	"k8s/armparameters.t":                                                                  k8sArmparametersT,
	"k8s/cloud-init/artifacts/apt-preferences":                                             k8sCloudInitArtifactsAptPreferences,
	"k8s/cloud-init/artifacts/auditd-rules":                                                k8sCloudInitArtifactsAuditdRules,
	"k8s/cloud-init/artifacts/cis.sh":                                                      k8sCloudInitArtifactsCisSh,
	"k8s/cloud-init/artifacts/cse_config.sh":                                               k8sCloudInitArtifactsCse_configSh,
	"k8s/cloud-init/artifacts/cse_customcloud.sh":                                          k8sCloudInitArtifactsCse_customcloudSh,
	"k8s/cloud-init/artifacts/cse_helpers.sh":                                              k8sCloudInitArtifactsCse_helpersSh,
	"k8s/cloud-init/artifacts/cse_install.sh":                                              k8sCloudInitArtifactsCse_installSh,
	"k8s/cloud-init/artifacts/cse_main.sh":                                                 k8sCloudInitArtifactsCse_mainSh,
	"k8s/cloud-init/artifacts/default-grub":                                                k8sCloudInitArtifactsDefaultGrub,
	"k8s/cloud-init/artifacts/dhcpv6.service":                                              k8sCloudInitArtifactsDhcpv6Service,
	"k8s/cloud-init/artifacts/docker-monitor.service":                                      k8sCloudInitArtifactsDockerMonitorService,
	"k8s/cloud-init/artifacts/docker-monitor.timer":                                        k8sCloudInitArtifactsDockerMonitorTimer,
	"k8s/cloud-init/artifacts/docker_clear_mount_propagation_flags.conf":                   k8sCloudInitArtifactsDocker_clear_mount_propagation_flagsConf,
	"k8s/cloud-init/artifacts/enable-dhcpv6.sh":                                            k8sCloudInitArtifactsEnableDhcpv6Sh,
	"k8s/cloud-init/artifacts/etc-issue":                                                   k8sCloudInitArtifactsEtcIssue,
	"k8s/cloud-init/artifacts/etc-issue.net":                                               k8sCloudInitArtifactsEtcIssueNet,
	"k8s/cloud-init/artifacts/etcd.service":                                                k8sCloudInitArtifactsEtcdService,
	"k8s/cloud-init/artifacts/generateproxycerts.sh":                                       k8sCloudInitArtifactsGenerateproxycertsSh,
	"k8s/cloud-init/artifacts/health-monitor.sh":                                           k8sCloudInitArtifactsHealthMonitorSh,
	"k8s/cloud-init/artifacts/kms.service":                                                 k8sCloudInitArtifactsKmsService,
	"k8s/cloud-init/artifacts/kubelet-monitor.service":                                     k8sCloudInitArtifactsKubeletMonitorService,
	"k8s/cloud-init/artifacts/kubelet-monitor.timer":                                       k8sCloudInitArtifactsKubeletMonitorTimer,
	"k8s/cloud-init/artifacts/kubelet.service":                                             k8sCloudInitArtifactsKubeletService,
	"k8s/cloud-init/artifacts/label-nodes.service":                                         k8sCloudInitArtifactsLabelNodesService,
	"k8s/cloud-init/artifacts/label-nodes.sh":                                              k8sCloudInitArtifactsLabelNodesSh,
	"k8s/cloud-init/artifacts/modprobe-CIS.conf":                                           k8sCloudInitArtifactsModprobeCisConf,
	"k8s/cloud-init/artifacts/mountetcd.sh":                                                k8sCloudInitArtifactsMountetcdSh,
	"k8s/cloud-init/artifacts/pam-d-common-auth":                                           k8sCloudInitArtifactsPamDCommonAuth,
	"k8s/cloud-init/artifacts/pam-d-common-password":                                       k8sCloudInitArtifactsPamDCommonPassword,
	"k8s/cloud-init/artifacts/pam-d-su":                                                    k8sCloudInitArtifactsPamDSu,
	"k8s/cloud-init/artifacts/profile-d-cis.sh":                                            k8sCloudInitArtifactsProfileDCisSh,
	"k8s/cloud-init/artifacts/pwquality-CIS.conf":                                          k8sCloudInitArtifactsPwqualityCisConf,
	"k8s/cloud-init/artifacts/rsyslog-d-60-CIS.conf":                                       k8sCloudInitArtifactsRsyslogD60CisConf,
	"k8s/cloud-init/artifacts/setup-custom-search-domains.sh":                              k8sCloudInitArtifactsSetupCustomSearchDomainsSh,
	"k8s/cloud-init/artifacts/sshd_config":                                                 k8sCloudInitArtifactsSshd_config,
	"k8s/cloud-init/artifacts/sshd_config_1604":                                            k8sCloudInitArtifactsSshd_config_1604,
	"k8s/cloud-init/artifacts/sys-fs-bpf.mount":                                            k8sCloudInitArtifactsSysFsBpfMount,
	"k8s/cloud-init/artifacts/sysctl-d-60-CIS.conf":                                        k8sCloudInitArtifactsSysctlD60CisConf,
	"k8s/cloud-init/jumpboxcustomdata.yml":                                                 k8sCloudInitJumpboxcustomdataYml,
	"k8s/cloud-init/masternodecustomdata.yml":                                              k8sCloudInitMasternodecustomdataYml,
	"k8s/cloud-init/nodecustomdata.yml":                                                    k8sCloudInitNodecustomdataYml,
	"k8s/containeraddons/1.16/azure-cni-networkmonitor.yaml":                               k8sContaineraddons116AzureCniNetworkmonitorYaml,
	"k8s/containeraddons/1.16/ip-masq-agent.yaml":                                          k8sContaineraddons116IpMasqAgentYaml,
	"k8s/containeraddons/1.16/kubernetesmasteraddons-aad-pod-identity-deployment.yaml":     k8sContaineraddons116KubernetesmasteraddonsAadPodIdentityDeploymentYaml,
	"k8s/containeraddons/1.16/kubernetesmasteraddons-aci-connector-deployment.yaml":        k8sContaineraddons116KubernetesmasteraddonsAciConnectorDeploymentYaml,
	"k8s/containeraddons/1.16/kubernetesmasteraddons-azure-npm-daemonset.yaml":             k8sContaineraddons116KubernetesmasteraddonsAzureNpmDaemonsetYaml,
	"k8s/containeraddons/1.16/kubernetesmasteraddons-blobfuse-flexvolume-installer.yaml":   k8sContaineraddons116KubernetesmasteraddonsBlobfuseFlexvolumeInstallerYaml,
	"k8s/containeraddons/1.16/kubernetesmasteraddons-calico-daemonset.yaml":                k8sContaineraddons116KubernetesmasteraddonsCalicoDaemonsetYaml,
	"k8s/containeraddons/1.16/kubernetesmasteraddons-cluster-autoscaler-deployment.yaml":   k8sContaineraddons116KubernetesmasteraddonsClusterAutoscalerDeploymentYaml,
	"k8s/containeraddons/1.16/kubernetesmasteraddons-heapster-deployment.yaml":             k8sContaineraddons116KubernetesmasteraddonsHeapsterDeploymentYaml,
	"k8s/containeraddons/1.16/kubernetesmasteraddons-keyvault-flexvolume-installer.yaml":   k8sContaineraddons116KubernetesmasteraddonsKeyvaultFlexvolumeInstallerYaml,
	"k8s/containeraddons/1.16/kubernetesmasteraddons-kube-rescheduler-deployment.yaml":     k8sContaineraddons116KubernetesmasteraddonsKubeReschedulerDeploymentYaml,
	"k8s/containeraddons/1.16/kubernetesmasteraddons-kubernetes-dashboard-deployment.yaml": k8sContaineraddons116KubernetesmasteraddonsKubernetesDashboardDeploymentYaml,
	"k8s/containeraddons/1.16/kubernetesmasteraddons-metrics-server-deployment.yaml":       k8sContaineraddons116KubernetesmasteraddonsMetricsServerDeploymentYaml,
	"k8s/containeraddons/1.16/kubernetesmasteraddons-nvidia-device-plugin-daemonset.yaml":  k8sContaineraddons116KubernetesmasteraddonsNvidiaDevicePluginDaemonsetYaml,
	"k8s/containeraddons/1.16/kubernetesmasteraddons-omsagent-daemonset.yaml":              k8sContaineraddons116KubernetesmasteraddonsOmsagentDaemonsetYaml,
	"k8s/containeraddons/1.16/kubernetesmasteraddons-smb-flexvolume-installer.yaml":        k8sContaineraddons116KubernetesmasteraddonsSmbFlexvolumeInstallerYaml,
	"k8s/containeraddons/1.16/kubernetesmasteraddons-tiller-deployment.yaml":               k8sContaineraddons116KubernetesmasteraddonsTillerDeploymentYaml,
	"k8s/containeraddons/1.6/kubernetesmasteraddons-heapster-deployment.yaml":              k8sContaineraddons16KubernetesmasteraddonsHeapsterDeploymentYaml,
	"k8s/containeraddons/1.7/kubernetesmasteraddons-heapster-deployment.yaml":              k8sContaineraddons17KubernetesmasteraddonsHeapsterDeploymentYaml,
	"k8s/containeraddons/1.8/kubernetesmasteraddons-heapster-deployment.yaml":              k8sContaineraddons18KubernetesmasteraddonsHeapsterDeploymentYaml,
	"k8s/containeraddons/azure-cni-networkmonitor.yaml":                                    k8sContaineraddonsAzureCniNetworkmonitorYaml,
	"k8s/containeraddons/dns-autoscaler.yaml":                                              k8sContaineraddonsDnsAutoscalerYaml,
	"k8s/containeraddons/ip-masq-agent.yaml":                                               k8sContaineraddonsIpMasqAgentYaml,
	"k8s/containeraddons/kubernetesmasteraddons-aad-pod-identity-deployment.yaml":          k8sContaineraddonsKubernetesmasteraddonsAadPodIdentityDeploymentYaml,
	"k8s/containeraddons/kubernetesmasteraddons-aci-connector-deployment.yaml":             k8sContaineraddonsKubernetesmasteraddonsAciConnectorDeploymentYaml,
	"k8s/containeraddons/kubernetesmasteraddons-azure-npm-daemonset.yaml":                  k8sContaineraddonsKubernetesmasteraddonsAzureNpmDaemonsetYaml,
	"k8s/containeraddons/kubernetesmasteraddons-blobfuse-flexvolume-installer.yaml":        k8sContaineraddonsKubernetesmasteraddonsBlobfuseFlexvolumeInstallerYaml,
	"k8s/containeraddons/kubernetesmasteraddons-calico-daemonset.yaml":                     k8sContaineraddonsKubernetesmasteraddonsCalicoDaemonsetYaml,
	"k8s/containeraddons/kubernetesmasteraddons-cluster-autoscaler-deployment.yaml":        k8sContaineraddonsKubernetesmasteraddonsClusterAutoscalerDeploymentYaml,
	"k8s/containeraddons/kubernetesmasteraddons-heapster-deployment.yaml":                  k8sContaineraddonsKubernetesmasteraddonsHeapsterDeploymentYaml,
	"k8s/containeraddons/kubernetesmasteraddons-keyvault-flexvolume-installer.yaml":        k8sContaineraddonsKubernetesmasteraddonsKeyvaultFlexvolumeInstallerYaml,
	"k8s/containeraddons/kubernetesmasteraddons-kube-rescheduler-deployment.yaml":          k8sContaineraddonsKubernetesmasteraddonsKubeReschedulerDeploymentYaml,
	"k8s/containeraddons/kubernetesmasteraddons-kubernetes-dashboard-deployment.yaml":      k8sContaineraddonsKubernetesmasteraddonsKubernetesDashboardDeploymentYaml,
	"k8s/containeraddons/kubernetesmasteraddons-metrics-server-deployment.yaml":            k8sContaineraddonsKubernetesmasteraddonsMetricsServerDeploymentYaml,
	"k8s/containeraddons/kubernetesmasteraddons-nvidia-device-plugin-daemonset.yaml":       k8sContaineraddonsKubernetesmasteraddonsNvidiaDevicePluginDaemonsetYaml,
	"k8s/containeraddons/kubernetesmasteraddons-omsagent-daemonset.yaml":                   k8sContaineraddonsKubernetesmasteraddonsOmsagentDaemonsetYaml,
	"k8s/containeraddons/kubernetesmasteraddons-smb-flexvolume-installer.yaml":             k8sContaineraddonsKubernetesmasteraddonsSmbFlexvolumeInstallerYaml,
	"k8s/containeraddons/kubernetesmasteraddons-tiller-deployment.yaml":                    k8sContaineraddonsKubernetesmasteraddonsTillerDeploymentYaml,
	"k8s/kubeconfig.json":                                                k8sKubeconfigJson,
	"k8s/kubernetesparams.t":                                             k8sKubernetesparamsT,
	"k8s/kuberneteswindowsfunctions.ps1":                                 k8sKuberneteswindowsfunctionsPs1,
	"k8s/kuberneteswindowssetup.ps1":                                     k8sKuberneteswindowssetupPs1,
	"k8s/manifests/kubernetesmaster-cloud-controller-manager.yaml":       k8sManifestsKubernetesmasterCloudControllerManagerYaml,
	"k8s/manifests/kubernetesmaster-kube-addon-manager.yaml":             k8sManifestsKubernetesmasterKubeAddonManagerYaml,
	"k8s/manifests/kubernetesmaster-kube-apiserver.yaml":                 k8sManifestsKubernetesmasterKubeApiserverYaml,
	"k8s/manifests/kubernetesmaster-kube-controller-manager-custom.yaml": k8sManifestsKubernetesmasterKubeControllerManagerCustomYaml,
	"k8s/manifests/kubernetesmaster-kube-controller-manager.yaml":        k8sManifestsKubernetesmasterKubeControllerManagerYaml,
	"k8s/manifests/kubernetesmaster-kube-scheduler.yaml":                 k8sManifestsKubernetesmasterKubeSchedulerYaml,
	"k8s/windowsazurecnifunc.ps1":                                        k8sWindowsazurecnifuncPs1,
	"k8s/windowscnifunc.ps1":                                             k8sWindowscnifuncPs1,
	"k8s/windowsconfigfunc.ps1":                                          k8sWindowsconfigfuncPs1,
	"k8s/windowsinstallopensshfunc.ps1":                                  k8sWindowsinstallopensshfuncPs1,
	"k8s/windowskubeletfunc.ps1":                                         k8sWindowskubeletfuncPs1,
	"masteroutputs.t":                                                    masteroutputsT,
	"masterparams.t":                                                     masterparamsT,
	"swarm/Install-ContainerHost-And-Join-Swarm.ps1":                     swarmInstallContainerhostAndJoinSwarmPs1,
	"swarm/Join-SwarmMode-cluster.ps1":                                   swarmJoinSwarmmodeClusterPs1,
	"swarm/configure-swarm-cluster.sh":                                   swarmConfigureSwarmClusterSh,
	"swarm/configure-swarmmode-cluster.sh":                               swarmConfigureSwarmmodeClusterSh,
	"swarm/swarmagentresourcesvmas.t":                                    swarmSwarmagentresourcesvmasT,
	"swarm/swarmagentresourcesvmss.t":                                    swarmSwarmagentresourcesvmssT,
	"swarm/swarmagentvars.t":                                             swarmSwarmagentvarsT,
	"swarm/swarmbase.t":                                                  swarmSwarmbaseT,
	"swarm/swarmmasterresources.t":                                       swarmSwarmmasterresourcesT,
	"swarm/swarmmastervars.t":                                            swarmSwarmmastervarsT,
	"swarm/swarmparams.t":                                                swarmSwarmparamsT,
	"swarm/swarmwinagentresourcesvmas.t":                                 swarmSwarmwinagentresourcesvmasT,
	"swarm/swarmwinagentresourcesvmss.t":                                 swarmSwarmwinagentresourcesvmssT,
	"windowsparams.t":                                                    windowsparamsT,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"agentoutputs.t": {agentoutputsT, map[string]*bintree{}},
	"agentparams.t":  {agentparamsT, map[string]*bintree{}},
	"dcos": {nil, map[string]*bintree{
		"bstrap": {nil, map[string]*bintree{
			"bootstrapcustomdata.yml": {dcosBstrapBootstrapcustomdataYml, map[string]*bintree{}},
			"bootstrapparams.t":       {dcosBstrapBootstrapparamsT, map[string]*bintree{}},
			"bootstrapprovision.sh":   {dcosBstrapBootstrapprovisionSh, map[string]*bintree{}},
			"bootstrapresources.t":    {dcosBstrapBootstrapresourcesT, map[string]*bintree{}},
			"bootstrapvars.t":         {dcosBstrapBootstrapvarsT, map[string]*bintree{}},
			"dcos1.11.0.customdata.t": {dcosBstrapDcos1110CustomdataT, map[string]*bintree{}},
			"dcos1.11.2.customdata.t": {dcosBstrapDcos1112CustomdataT, map[string]*bintree{}},
			"dcosbase.t":              {dcosBstrapDcosbaseT, map[string]*bintree{}},
			"dcosmasterresources.t":   {dcosBstrapDcosmasterresourcesT, map[string]*bintree{}},
			"dcosmastervars.t":        {dcosBstrapDcosmastervarsT, map[string]*bintree{}},
			"dcosprovision.sh":        {dcosBstrapDcosprovisionSh, map[string]*bintree{}},
		}},
		"dcosWindowsAgentResourcesVmas.t": {dcosDcoswindowsagentresourcesvmasT, map[string]*bintree{}},
		"dcosWindowsAgentResourcesVmss.t": {dcosDcoswindowsagentresourcesvmssT, map[string]*bintree{}},
		"dcosWindowsProvision.ps1":        {dcosDcoswindowsprovisionPs1, map[string]*bintree{}},
		"dcosagentresourcesvmas.t":        {dcosDcosagentresourcesvmasT, map[string]*bintree{}},
		"dcosagentresourcesvmss.t":        {dcosDcosagentresourcesvmssT, map[string]*bintree{}},
		"dcosagentvars.t":                 {dcosDcosagentvarsT, map[string]*bintree{}},
		"dcosbase.t":                      {dcosDcosbaseT, map[string]*bintree{}},
		"dcoscustomdata110.t":             {dcosDcoscustomdata110T, map[string]*bintree{}},
		"dcoscustomdata184.t":             {dcosDcoscustomdata184T, map[string]*bintree{}},
		"dcoscustomdata187.t":             {dcosDcoscustomdata187T, map[string]*bintree{}},
		"dcoscustomdata188.t":             {dcosDcoscustomdata188T, map[string]*bintree{}},
		"dcoscustomdata190.t":             {dcosDcoscustomdata190T, map[string]*bintree{}},
		"dcoscustomdata198.t":             {dcosDcoscustomdata198T, map[string]*bintree{}},
		"dcosmasterresources.t":           {dcosDcosmasterresourcesT, map[string]*bintree{}},
		"dcosmastervars.t":                {dcosDcosmastervarsT, map[string]*bintree{}},
		"dcosparams.t":                    {dcosDcosparamsT, map[string]*bintree{}},
		"dcosprovision.sh":                {dcosDcosprovisionSh, map[string]*bintree{}},
		"dcosprovisionsource.sh":          {dcosDcosprovisionsourceSh, map[string]*bintree{}},
	}},
	"iaasoutputs.t": {iaasoutputsT, map[string]*bintree{}},
	"k8s": {nil, map[string]*bintree{
		"addons": {nil, map[string]*bintree{
			"1.10": {nil, map[string]*bintree{
				"kubernetesmasteraddons-kube-dns-deployment.yaml": {k8sAddons110KubernetesmasteraddonsKubeDnsDeploymentYaml, map[string]*bintree{}},
			}},
			"1.15": {nil, map[string]*bintree{
				"kubernetesmasteraddons-azure-cloud-provider-deployment.yaml": {k8sAddons115KubernetesmasteraddonsAzureCloudProviderDeploymentYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-pod-security-policy.yaml":             {k8sAddons115KubernetesmasteraddonsPodSecurityPolicyYaml, map[string]*bintree{}},
			}},
			"1.16": {nil, map[string]*bintree{
				"kubernetesmaster-audit-policy.yaml":                          {k8sAddons116KubernetesmasterAuditPolicyYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-azure-cloud-provider-deployment.yaml": {k8sAddons116KubernetesmasteraddonsAzureCloudProviderDeploymentYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-cilium-daemonset.yaml":                {k8sAddons116KubernetesmasteraddonsCiliumDaemonsetYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-flannel-daemonset.yaml":               {k8sAddons116KubernetesmasteraddonsFlannelDaemonsetYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-kube-dns-deployment.yaml":             {k8sAddons116KubernetesmasteraddonsKubeDnsDeploymentYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-kube-proxy-daemonset.yaml":            {k8sAddons116KubernetesmasteraddonsKubeProxyDaemonsetYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-pod-security-policy.yaml":             {k8sAddons116KubernetesmasteraddonsPodSecurityPolicyYaml, map[string]*bintree{}},
			}},
			"1.6": {nil, map[string]*bintree{
				"kubernetesmasteraddons-kubernetes-dashboard-deployment.yaml": {k8sAddons16KubernetesmasteraddonsKubernetesDashboardDeploymentYaml, map[string]*bintree{}},
			}},
			"1.7": {nil, map[string]*bintree{
				"kubernetesmasteraddons-kube-dns-deployment.yaml":             {k8sAddons17KubernetesmasteraddonsKubeDnsDeploymentYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-kubernetes-dashboard-deployment.yaml": {k8sAddons17KubernetesmasteraddonsKubernetesDashboardDeploymentYaml, map[string]*bintree{}},
			}},
			"1.8": {nil, map[string]*bintree{
				"kubernetesmasteraddons-kube-dns-deployment.yaml":             {k8sAddons18KubernetesmasteraddonsKubeDnsDeploymentYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-kubernetes-dashboard-deployment.yaml": {k8sAddons18KubernetesmasteraddonsKubernetesDashboardDeploymentYaml, map[string]*bintree{}},
			}},
			"1.9": {nil, map[string]*bintree{
				"kubernetesmasteraddons-kube-dns-deployment.yaml": {k8sAddons19KubernetesmasteraddonsKubeDnsDeploymentYaml, map[string]*bintree{}},
			}},
			"coredns.yaml":                       {k8sAddonsCorednsYaml, map[string]*bintree{}},
			"kubernetesmaster-audit-policy.yaml": {k8sAddonsKubernetesmasterAuditPolicyYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-aad-default-admin-group-rbac.yaml":           {k8sAddonsKubernetesmasteraddonsAadDefaultAdminGroupRbacYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-azure-cloud-provider-deployment.yaml":        {k8sAddonsKubernetesmasteraddonsAzureCloudProviderDeploymentYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-cilium-daemonset.yaml":                       {k8sAddonsKubernetesmasteraddonsCiliumDaemonsetYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-elb-svc.yaml":                                {k8sAddonsKubernetesmasteraddonsElbSvcYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-flannel-daemonset.yaml":                      {k8sAddonsKubernetesmasteraddonsFlannelDaemonsetYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-kube-dns-deployment.yaml":                    {k8sAddonsKubernetesmasteraddonsKubeDnsDeploymentYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-kube-proxy-daemonset.yaml":                   {k8sAddonsKubernetesmasteraddonsKubeProxyDaemonsetYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-managed-azure-storage-classes-custom.yaml":   {k8sAddonsKubernetesmasteraddonsManagedAzureStorageClassesCustomYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-managed-azure-storage-classes.yaml":          {k8sAddonsKubernetesmasteraddonsManagedAzureStorageClassesYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-pod-security-policy.yaml":                    {k8sAddonsKubernetesmasteraddonsPodSecurityPolicyYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-scheduled-maintenance-deployment.yaml":       {k8sAddonsKubernetesmasteraddonsScheduledMaintenanceDeploymentYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-unmanaged-azure-storage-classes-custom.yaml": {k8sAddonsKubernetesmasteraddonsUnmanagedAzureStorageClassesCustomYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-unmanaged-azure-storage-classes.yaml":        {k8sAddonsKubernetesmasteraddonsUnmanagedAzureStorageClassesYaml, map[string]*bintree{}},
		}},
		"armparameters.t": {k8sArmparametersT, map[string]*bintree{}},
		"cloud-init": {nil, map[string]*bintree{
			"artifacts": {nil, map[string]*bintree{
				"apt-preferences":        {k8sCloudInitArtifactsAptPreferences, map[string]*bintree{}},
				"auditd-rules":           {k8sCloudInitArtifactsAuditdRules, map[string]*bintree{}},
				"cis.sh":                 {k8sCloudInitArtifactsCisSh, map[string]*bintree{}},
				"cse_config.sh":          {k8sCloudInitArtifactsCse_configSh, map[string]*bintree{}},
				"cse_customcloud.sh":     {k8sCloudInitArtifactsCse_customcloudSh, map[string]*bintree{}},
				"cse_helpers.sh":         {k8sCloudInitArtifactsCse_helpersSh, map[string]*bintree{}},
				"cse_install.sh":         {k8sCloudInitArtifactsCse_installSh, map[string]*bintree{}},
				"cse_main.sh":            {k8sCloudInitArtifactsCse_mainSh, map[string]*bintree{}},
				"default-grub":           {k8sCloudInitArtifactsDefaultGrub, map[string]*bintree{}},
				"dhcpv6.service":         {k8sCloudInitArtifactsDhcpv6Service, map[string]*bintree{}},
				"docker-monitor.service": {k8sCloudInitArtifactsDockerMonitorService, map[string]*bintree{}},
				"docker-monitor.timer":   {k8sCloudInitArtifactsDockerMonitorTimer, map[string]*bintree{}},
				"docker_clear_mount_propagation_flags.conf": {k8sCloudInitArtifactsDocker_clear_mount_propagation_flagsConf, map[string]*bintree{}},
				"enable-dhcpv6.sh":                          {k8sCloudInitArtifactsEnableDhcpv6Sh, map[string]*bintree{}},
				"etc-issue":                                 {k8sCloudInitArtifactsEtcIssue, map[string]*bintree{}},
				"etc-issue.net":                             {k8sCloudInitArtifactsEtcIssueNet, map[string]*bintree{}},
				"etcd.service":                              {k8sCloudInitArtifactsEtcdService, map[string]*bintree{}},
				"generateproxycerts.sh":                     {k8sCloudInitArtifactsGenerateproxycertsSh, map[string]*bintree{}},
				"health-monitor.sh":                         {k8sCloudInitArtifactsHealthMonitorSh, map[string]*bintree{}},
				"kms.service":                               {k8sCloudInitArtifactsKmsService, map[string]*bintree{}},
				"kubelet-monitor.service":                   {k8sCloudInitArtifactsKubeletMonitorService, map[string]*bintree{}},
				"kubelet-monitor.timer":                     {k8sCloudInitArtifactsKubeletMonitorTimer, map[string]*bintree{}},
				"kubelet.service":                           {k8sCloudInitArtifactsKubeletService, map[string]*bintree{}},
				"label-nodes.service":                       {k8sCloudInitArtifactsLabelNodesService, map[string]*bintree{}},
				"label-nodes.sh":                            {k8sCloudInitArtifactsLabelNodesSh, map[string]*bintree{}},
				"modprobe-CIS.conf":                         {k8sCloudInitArtifactsModprobeCisConf, map[string]*bintree{}},
				"mountetcd.sh":                              {k8sCloudInitArtifactsMountetcdSh, map[string]*bintree{}},
				"pam-d-common-auth":                         {k8sCloudInitArtifactsPamDCommonAuth, map[string]*bintree{}},
				"pam-d-common-password":                     {k8sCloudInitArtifactsPamDCommonPassword, map[string]*bintree{}},
				"pam-d-su":                                  {k8sCloudInitArtifactsPamDSu, map[string]*bintree{}},
				"profile-d-cis.sh":                          {k8sCloudInitArtifactsProfileDCisSh, map[string]*bintree{}},
				"pwquality-CIS.conf":                        {k8sCloudInitArtifactsPwqualityCisConf, map[string]*bintree{}},
				"rsyslog-d-60-CIS.conf":                     {k8sCloudInitArtifactsRsyslogD60CisConf, map[string]*bintree{}},
				"setup-custom-search-domains.sh":            {k8sCloudInitArtifactsSetupCustomSearchDomainsSh, map[string]*bintree{}},
				"sshd_config":                               {k8sCloudInitArtifactsSshd_config, map[string]*bintree{}},
				"sshd_config_1604":                          {k8sCloudInitArtifactsSshd_config_1604, map[string]*bintree{}},
				"sys-fs-bpf.mount":                          {k8sCloudInitArtifactsSysFsBpfMount, map[string]*bintree{}},
				"sysctl-d-60-CIS.conf":                      {k8sCloudInitArtifactsSysctlD60CisConf, map[string]*bintree{}},
			}},
			"jumpboxcustomdata.yml":    {k8sCloudInitJumpboxcustomdataYml, map[string]*bintree{}},
			"masternodecustomdata.yml": {k8sCloudInitMasternodecustomdataYml, map[string]*bintree{}},
			"nodecustomdata.yml":       {k8sCloudInitNodecustomdataYml, map[string]*bintree{}},
		}},
		"containeraddons": {nil, map[string]*bintree{
			"1.16": {nil, map[string]*bintree{
				"azure-cni-networkmonitor.yaml":                               {k8sContaineraddons116AzureCniNetworkmonitorYaml, map[string]*bintree{}},
				"ip-masq-agent.yaml":                                          {k8sContaineraddons116IpMasqAgentYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-aad-pod-identity-deployment.yaml":     {k8sContaineraddons116KubernetesmasteraddonsAadPodIdentityDeploymentYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-aci-connector-deployment.yaml":        {k8sContaineraddons116KubernetesmasteraddonsAciConnectorDeploymentYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-azure-npm-daemonset.yaml":             {k8sContaineraddons116KubernetesmasteraddonsAzureNpmDaemonsetYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-blobfuse-flexvolume-installer.yaml":   {k8sContaineraddons116KubernetesmasteraddonsBlobfuseFlexvolumeInstallerYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-calico-daemonset.yaml":                {k8sContaineraddons116KubernetesmasteraddonsCalicoDaemonsetYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-cluster-autoscaler-deployment.yaml":   {k8sContaineraddons116KubernetesmasteraddonsClusterAutoscalerDeploymentYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-heapster-deployment.yaml":             {k8sContaineraddons116KubernetesmasteraddonsHeapsterDeploymentYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-keyvault-flexvolume-installer.yaml":   {k8sContaineraddons116KubernetesmasteraddonsKeyvaultFlexvolumeInstallerYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-kube-rescheduler-deployment.yaml":     {k8sContaineraddons116KubernetesmasteraddonsKubeReschedulerDeploymentYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-kubernetes-dashboard-deployment.yaml": {k8sContaineraddons116KubernetesmasteraddonsKubernetesDashboardDeploymentYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-metrics-server-deployment.yaml":       {k8sContaineraddons116KubernetesmasteraddonsMetricsServerDeploymentYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-nvidia-device-plugin-daemonset.yaml":  {k8sContaineraddons116KubernetesmasteraddonsNvidiaDevicePluginDaemonsetYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-omsagent-daemonset.yaml":              {k8sContaineraddons116KubernetesmasteraddonsOmsagentDaemonsetYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-smb-flexvolume-installer.yaml":        {k8sContaineraddons116KubernetesmasteraddonsSmbFlexvolumeInstallerYaml, map[string]*bintree{}},
				"kubernetesmasteraddons-tiller-deployment.yaml":               {k8sContaineraddons116KubernetesmasteraddonsTillerDeploymentYaml, map[string]*bintree{}},
			}},
			"1.6": {nil, map[string]*bintree{
				"kubernetesmasteraddons-heapster-deployment.yaml": {k8sContaineraddons16KubernetesmasteraddonsHeapsterDeploymentYaml, map[string]*bintree{}},
			}},
			"1.7": {nil, map[string]*bintree{
				"kubernetesmasteraddons-heapster-deployment.yaml": {k8sContaineraddons17KubernetesmasteraddonsHeapsterDeploymentYaml, map[string]*bintree{}},
			}},
			"1.8": {nil, map[string]*bintree{
				"kubernetesmasteraddons-heapster-deployment.yaml": {k8sContaineraddons18KubernetesmasteraddonsHeapsterDeploymentYaml, map[string]*bintree{}},
			}},
			"azure-cni-networkmonitor.yaml":                               {k8sContaineraddonsAzureCniNetworkmonitorYaml, map[string]*bintree{}},
			"dns-autoscaler.yaml":                                         {k8sContaineraddonsDnsAutoscalerYaml, map[string]*bintree{}},
			"ip-masq-agent.yaml":                                          {k8sContaineraddonsIpMasqAgentYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-aad-pod-identity-deployment.yaml":     {k8sContaineraddonsKubernetesmasteraddonsAadPodIdentityDeploymentYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-aci-connector-deployment.yaml":        {k8sContaineraddonsKubernetesmasteraddonsAciConnectorDeploymentYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-azure-npm-daemonset.yaml":             {k8sContaineraddonsKubernetesmasteraddonsAzureNpmDaemonsetYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-blobfuse-flexvolume-installer.yaml":   {k8sContaineraddonsKubernetesmasteraddonsBlobfuseFlexvolumeInstallerYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-calico-daemonset.yaml":                {k8sContaineraddonsKubernetesmasteraddonsCalicoDaemonsetYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-cluster-autoscaler-deployment.yaml":   {k8sContaineraddonsKubernetesmasteraddonsClusterAutoscalerDeploymentYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-heapster-deployment.yaml":             {k8sContaineraddonsKubernetesmasteraddonsHeapsterDeploymentYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-keyvault-flexvolume-installer.yaml":   {k8sContaineraddonsKubernetesmasteraddonsKeyvaultFlexvolumeInstallerYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-kube-rescheduler-deployment.yaml":     {k8sContaineraddonsKubernetesmasteraddonsKubeReschedulerDeploymentYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-kubernetes-dashboard-deployment.yaml": {k8sContaineraddonsKubernetesmasteraddonsKubernetesDashboardDeploymentYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-metrics-server-deployment.yaml":       {k8sContaineraddonsKubernetesmasteraddonsMetricsServerDeploymentYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-nvidia-device-plugin-daemonset.yaml":  {k8sContaineraddonsKubernetesmasteraddonsNvidiaDevicePluginDaemonsetYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-omsagent-daemonset.yaml":              {k8sContaineraddonsKubernetesmasteraddonsOmsagentDaemonsetYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-smb-flexvolume-installer.yaml":        {k8sContaineraddonsKubernetesmasteraddonsSmbFlexvolumeInstallerYaml, map[string]*bintree{}},
			"kubernetesmasteraddons-tiller-deployment.yaml":               {k8sContaineraddonsKubernetesmasteraddonsTillerDeploymentYaml, map[string]*bintree{}},
		}},
		"kubeconfig.json":                {k8sKubeconfigJson, map[string]*bintree{}},
		"kubernetesparams.t":             {k8sKubernetesparamsT, map[string]*bintree{}},
		"kuberneteswindowsfunctions.ps1": {k8sKuberneteswindowsfunctionsPs1, map[string]*bintree{}},
		"kuberneteswindowssetup.ps1":     {k8sKuberneteswindowssetupPs1, map[string]*bintree{}},
		"manifests": {nil, map[string]*bintree{
			"kubernetesmaster-cloud-controller-manager.yaml":       {k8sManifestsKubernetesmasterCloudControllerManagerYaml, map[string]*bintree{}},
			"kubernetesmaster-kube-addon-manager.yaml":             {k8sManifestsKubernetesmasterKubeAddonManagerYaml, map[string]*bintree{}},
			"kubernetesmaster-kube-apiserver.yaml":                 {k8sManifestsKubernetesmasterKubeApiserverYaml, map[string]*bintree{}},
			"kubernetesmaster-kube-controller-manager-custom.yaml": {k8sManifestsKubernetesmasterKubeControllerManagerCustomYaml, map[string]*bintree{}},
			"kubernetesmaster-kube-controller-manager.yaml":        {k8sManifestsKubernetesmasterKubeControllerManagerYaml, map[string]*bintree{}},
			"kubernetesmaster-kube-scheduler.yaml":                 {k8sManifestsKubernetesmasterKubeSchedulerYaml, map[string]*bintree{}},
		}},
		"windowsazurecnifunc.ps1":       {k8sWindowsazurecnifuncPs1, map[string]*bintree{}},
		"windowscnifunc.ps1":            {k8sWindowscnifuncPs1, map[string]*bintree{}},
		"windowsconfigfunc.ps1":         {k8sWindowsconfigfuncPs1, map[string]*bintree{}},
		"windowsinstallopensshfunc.ps1": {k8sWindowsinstallopensshfuncPs1, map[string]*bintree{}},
		"windowskubeletfunc.ps1":        {k8sWindowskubeletfuncPs1, map[string]*bintree{}},
	}},
	"masteroutputs.t": {masteroutputsT, map[string]*bintree{}},
	"masterparams.t":  {masterparamsT, map[string]*bintree{}},
	"swarm": {nil, map[string]*bintree{
		"Install-ContainerHost-And-Join-Swarm.ps1": {swarmInstallContainerhostAndJoinSwarmPs1, map[string]*bintree{}},
		"Join-SwarmMode-cluster.ps1":               {swarmJoinSwarmmodeClusterPs1, map[string]*bintree{}},
		"configure-swarm-cluster.sh":               {swarmConfigureSwarmClusterSh, map[string]*bintree{}},
		"configure-swarmmode-cluster.sh":           {swarmConfigureSwarmmodeClusterSh, map[string]*bintree{}},
		"swarmagentresourcesvmas.t":                {swarmSwarmagentresourcesvmasT, map[string]*bintree{}},
		"swarmagentresourcesvmss.t":                {swarmSwarmagentresourcesvmssT, map[string]*bintree{}},
		"swarmagentvars.t":                         {swarmSwarmagentvarsT, map[string]*bintree{}},
		"swarmbase.t":                              {swarmSwarmbaseT, map[string]*bintree{}},
		"swarmmasterresources.t":                   {swarmSwarmmasterresourcesT, map[string]*bintree{}},
		"swarmmastervars.t":                        {swarmSwarmmastervarsT, map[string]*bintree{}},
		"swarmparams.t":                            {swarmSwarmparamsT, map[string]*bintree{}},
		"swarmwinagentresourcesvmas.t":             {swarmSwarmwinagentresourcesvmasT, map[string]*bintree{}},
		"swarmwinagentresourcesvmss.t":             {swarmSwarmwinagentresourcesvmssT, map[string]*bintree{}},
	}},
	"windowsparams.t": {windowsparamsT, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
