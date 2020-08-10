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
// ../../parts/k8s/addons/1.15/calico.yaml
// ../../parts/k8s/addons/aad-default-admin-group-rbac.yaml
// ../../parts/k8s/addons/aad-pod-identity.yaml
// ../../parts/k8s/addons/aci-connector.yaml
// ../../parts/k8s/addons/antrea.yaml
// ../../parts/k8s/addons/arc-onboarding.yaml
// ../../parts/k8s/addons/audit-policy.yaml
// ../../parts/k8s/addons/azure-cloud-provider.yaml
// ../../parts/k8s/addons/azure-cni-networkmonitor.yaml
// ../../parts/k8s/addons/azure-network-policy.yaml
// ../../parts/k8s/addons/azure-policy-deployment.yaml
// ../../parts/k8s/addons/azuredisk-csi-driver-deployment.yaml
// ../../parts/k8s/addons/azurefile-csi-driver-deployment.yaml
// ../../parts/k8s/addons/blobfuse-flexvolume.yaml
// ../../parts/k8s/addons/calico.yaml
// ../../parts/k8s/addons/cilium.yaml
// ../../parts/k8s/addons/cloud-node-manager.yaml
// ../../parts/k8s/addons/cluster-autoscaler.yaml
// ../../parts/k8s/addons/container-monitoring.yaml
// ../../parts/k8s/addons/coredns.yaml
// ../../parts/k8s/addons/flannel.yaml
// ../../parts/k8s/addons/ip-masq-agent.yaml
// ../../parts/k8s/addons/keyvault-flexvolume.yaml
// ../../parts/k8s/addons/kube-dns.yaml
// ../../parts/k8s/addons/kube-proxy.yaml
// ../../parts/k8s/addons/kube-rescheduler.yaml
// ../../parts/k8s/addons/kubernetes-dashboard.yaml
// ../../parts/k8s/addons/metrics-server.yaml
// ../../parts/k8s/addons/node-problem-detector.yaml
// ../../parts/k8s/addons/nvidia-device-plugin.yaml
// ../../parts/k8s/addons/pod-security-policy.yaml
// ../../parts/k8s/addons/scheduled-maintenance-deployment.yaml
// ../../parts/k8s/addons/secrets-store-csi-driver.yaml
// ../../parts/k8s/addons/smb-flexvolume.yaml
// ../../parts/k8s/addons/tiller.yaml
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
// ../../parts/k8s/cloud-init/artifacts/untaint-nodes.service
// ../../parts/k8s/cloud-init/artifacts/untaint-nodes.sh
// ../../parts/k8s/cloud-init/jumpboxcustomdata.yml
// ../../parts/k8s/cloud-init/masternodecustomdata.yml
// ../../parts/k8s/cloud-init/nodecustomdata.yml
// ../../parts/k8s/kubeconfig.json
// ../../parts/k8s/kubernetesparams.t
// ../../parts/k8s/kuberneteswindowsfunctions.ps1
// ../../parts/k8s/kuberneteswindowssetup.ps1
// ../../parts/k8s/manifests/kubernetesmaster-cloud-controller-manager.yaml
// ../../parts/k8s/manifests/kubernetesmaster-kube-addon-manager.yaml
// ../../parts/k8s/manifests/kubernetesmaster-kube-apiserver.yaml
// ../../parts/k8s/manifests/kubernetesmaster-kube-controller-manager.yaml
// ../../parts/k8s/manifests/kubernetesmaster-kube-scheduler.yaml
// ../../parts/k8s/windowsazurecnifunc.ps1
// ../../parts/k8s/windowsazurecnifunc.tests.ps1
// ../../parts/k8s/windowscnifunc.ps1
// ../../parts/k8s/windowsconfigfunc.ps1
// ../../parts/k8s/windowscontainerdfunc.ps1
// ../../parts/k8s/windowscsiproxyfunc.ps1
// ../../parts/k8s/windowshostsconfigagentfunc.ps1
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
{{if .IsSpotScaleSet}}
    "{{.Name}}ScaleSetPriority": {
      "allowedValues":[
        "Spot",
        "Regular",
        ""
      ],
      "defaultValue": "{{.ScaleSetPriority}}",
      "metadata": {
        "description": "The priority for the VM Scale Set. This value can be Spot or Regular."
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
        "description": "The Eviction Policy for a Spot VM Scale Set."
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

var _k8sAddons115CalicoYaml = []byte(`{{- /* Source: calico/templates/calico-config.yaml
This ConfigMap is used to configure a self-hosted Calico installation. */}}
kind: ConfigMap
apiVersion: v1
metadata:
  name: calico-config
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "EnsureExists"
data:
  {{- /* You must set a non-zero value for Typha replicas below. */}}
  typha_service_name: "calico-typha"
  {{- /* The CNI network configuration to install on each node.  The special
  values in this config will be automatically populated. */}}
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
{{- /* Source: calico/templates/kdd-crds.yaml */}}
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: felixconfigurations.crd.projectcalico.org
  labels:
    addonmanager.kubernetes.io/mode: "EnsureExists"
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
    addonmanager.kubernetes.io/mode: "EnsureExists"
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
    addonmanager.kubernetes.io/mode: "EnsureExists"
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
    addonmanager.kubernetes.io/mode: "EnsureExists"
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
    addonmanager.kubernetes.io/mode: "EnsureExists"
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
    addonmanager.kubernetes.io/mode: "EnsureExists"
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
    addonmanager.kubernetes.io/mode: "EnsureExists"
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
    addonmanager.kubernetes.io/mode: "EnsureExists"
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
    addonmanager.kubernetes.io/mode: "EnsureExists"
spec:
  scope: Namespaced
  group: crd.projectcalico.org
  version: v1
  names:
    kind: NetworkSet
    plural: networksets
    singular: networkset
---
{{- /* Source: calico/templates/rbac.yaml
Include a clusterrole for the calico-node DaemonSet,
and bind it to the calico-node serviceaccount. */}}
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: calico-node
  labels:
    addonmanager.kubernetes.io/mode: "EnsureExists"
rules:
{{- /* The CNI plugin needs to get pods, nodes, and namespaces. */}}
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
  {{- /* Used to discover service IPs for advertisement. */}}
  - watch
  - list
  {{- /* Used to discover Typhas. */}}
  - get
- apiGroups: [""]
  resources:
  - nodes/status
  verbs:
  {{- /* Needed for clearing NodeNetworkUnavailable flag. */}}
  - patch
  {{- /* Calico stores some configuration information in node annotations. */}}
  - update
{{- /* Watch for changes to Kubernetes NetworkPolicies. */}}
- apiGroups: ["networking.k8s.io"]
  resources:
  - networkpolicies
  verbs:
  - watch
  - list
{{- /* Used by Calico for policy information. */}}
- apiGroups: [""]
  resources:
  - pods
  - namespaces
  - serviceaccounts
  verbs:
  - list
  - watch
{{- /* The CNI plugin patches pods/status. */}}
- apiGroups: [""]
  resources:
  - pods/status
  verbs:
  - patch
{{- /* Calico monitors various CRDs for config. */}}
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
{{- /* Calico must create and update some CRDs on startup. */}}
- apiGroups: ["crd.projectcalico.org"]
  resources:
  - ippools
  - felixconfigurations
  - clusterinformations
  verbs:
  - create
  - update
{{- /* Calico stores some configuration information on the node. */}}
- apiGroups: [""]
  resources:
  - nodes
  verbs:
  - get
  - list
  - watch
{{- /* These permissions are only requried for upgrade from v2.6, and can
be removed after upgrade or on fresh installations. */}}
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
    addonmanager.kubernetes.io/mode: "EnsureExists"
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: calico-node
subjects:
- kind: ServiceAccount
  name: calico-node
  namespace: kube-system
---
{{- /* Source: calico/templates/calico-typha.yaml
This manifest creates a Service, which will be backed by Calico's Typha daemon.
Typha sits in between Felix and the API server, reducing Calico's load on the API server. */}}
apiVersion: v1
kind: Service
metadata:
  name: calico-typha
  namespace: kube-system
  labels:
    k8s-app: calico-typha
    addonmanager.kubernetes.io/mode: "EnsureExists"
spec:
  ports:
  - port: 5473
    protocol: TCP
    targetPort: calico-typha
    name: calico-typha
  selector:
    k8s-app: calico-typha
---
{{- /* This manifest creates a Deployment of Typha to back the above service. */}}
apiVersion: apps/v1
kind: Deployment
metadata:
  name: calico-typha
  namespace: kube-system
  labels:
    k8s-app: calico-typha
    addonmanager.kubernetes.io/mode: "EnsureExists"
spec:
  {{- /* Number of Typha replicas.  To enable Typha, set this to a non-zero value *and* set the
  typha_service_name variable in the calico-config ConfigMap above.
  We recommend using Typha if you have more than 50 nodes.  Above 100 nodes it is essential
  (when using the Kubernetes datastore).  Use one replica for every 100-200 nodes.  In
  production, we recommend running at least 3 replicas to reduce the impact of rolling upgrade. */}}
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
        {{- /* This, along with the CriticalAddonsOnly toleration below, marks the pod as a critical
        add-on, ensuring it gets priority scheduling and that its resources are reserved
        if it ever gets evicted. */}}
        scheduler.alpha.kubernetes.io/critical-pod: ''
        cluster-autoscaler.kubernetes.io/safe-to-evict: 'true'
    spec:
      nodeSelector:
        kubernetes.io/os: linux
      hostNetwork: true
      tolerations:
      {{- /* Mark the pod as a critical add-on for rescheduling. */}}
      - key: CriticalAddonsOnly
        operator: Exists
      {{- /* Since Calico can't network a pod until Typha is up, we need to run Typha itself
      as a host-networked pod. */}}
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
        {{- /* Enable "info" logging by default.  Can be set to "debug" to increase verbosity. */}}
        - name: TYPHA_LOGSEVERITYSCREEN
          value: "info"
        {{- /* Disable logging to file and syslog since those don't make sense in Kubernetes. */}}
        - name: TYPHA_LOGFILEPATH
          value: "none"
        - name: TYPHA_LOGSEVERITYSYS
          value: "none"
        {{- /* Monitor the Kubernetes API to find the number of running instances and rebalance
        connections. */}}
        - name: TYPHA_CONNECTIONREBALANCINGMODE
          value: "kubernetes"
        - name: TYPHA_DATASTORETYPE
          value: "kubernetes"
        - name: TYPHA_HEALTHENABLED
          value: "true"
        {{- /* Configure route aggregation based on pod CIDR. */}}
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
{{- /* Source: calico/templates/calico-node.yaml
This manifest installs the calico-node container, as well
as the CNI plugins and network config on
each master and worker node in a Kubernetes cluster. */}}
kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: calico-node
  namespace: kube-system
  labels:
    k8s-app: calico-node
    addonmanager.kubernetes.io/mode: "EnsureExists"
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
        {{- /* This, along with the CriticalAddonsOnly toleration below,
        marks the pod as a critical add-on, ensuring it gets
        priority scheduling and that its resources are reserved
        if it ever gets evicted. */}}
        scheduler.alpha.kubernetes.io/critical-pod: ''
    spec:
      nodeSelector:
        kubernetes.io/os: linux
      hostNetwork: true
      tolerations:
      {{- /* Make sure calico-node gets scheduled on all nodes. */}}
      - effect: NoSchedule
        operator: Exists
      {{- /* Mark the pod as a critical add-on for rescheduling. */}}
      - key: CriticalAddonsOnly
        operator: Exists
      - effect: NoExecute
        operator: Exists
      serviceAccountName: calico-node
      {{- /* Minimize downtime during a rolling upgrade or deletion; tell Kubernetes to do a "force deletion":
      https://kubernetes.io/docs/concepts/workloads/pods/pod/#termination-of-pods. */}}
      terminationGracePeriodSeconds: 0
      priorityClassName: system-node-critical
      initContainers:
      {{- /* Start of install-cni initContainer
      This container installs the CNI binaries
      and CNI network config file on each node. */}}
      - name: install-cni
        image: {{ContainerImage "calico-cni"}}
        command: ["/install-cni.sh"]
        env:
        {{- /* Name of the CNI config file to create. */}}
        - name: CNI_CONF_NAME
          value: "10-calico.conflist"
        {{- /* The CNI network config to install on each node. */}}
        - name: CNI_NETWORK_CONFIG
          valueFrom:
            configMapKeyRef:
              name: calico-config
              key: cni_network_config
        {{- /* Set the hostname based on the k8s node name. */}}
        - name: KUBERNETES_NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        {{- /* Prevents the container from sleeping forever. */}}
        - name: SLEEP
          value: "false"
        volumeMounts:
        - mountPath: /host/opt/cni/bin
          name: cni-bin-dir
        - mountPath: /host/etc/cni/net.d
          name: cni-net-dir
      {{- /* End of install-cni initContainer
      Adds a Flex Volume Driver that creates a per-pod Unix Domain Socket to allow Dikastes
      to communicate with Felix over the Policy Sync API. */}}
      - name: flexvol-driver
        image: {{ContainerImage "calico-pod2daemon"}}
        volumeMounts:
        - name: flexvol-driver-host
          mountPath: /host/driver
      containers:
      {{- /* Runs calico-node container on each Kubernetes node.  This
      container programs network policy and routes on each
      host. */}}
      - name: calico-node
        image: {{ContainerImage "calico-node"}}
        env:
        {{- /* Use Kubernetes API as the backing datastore. */}}
        - name: DATASTORE_TYPE
          value: "kubernetes"
        {{- /* Configure route aggregation based on pod CIDR. */}}
        - name: USE_POD_CIDR
          value: "true"
        {{- /* Typha support: controlled by the ConfigMap. */}}
        - name: FELIX_TYPHAK8SSERVICENAME
          valueFrom:
            configMapKeyRef:
              name: calico-config
              key: typha_service_name
        {{- /* Wait for the datastore. */}}
        - name: WAIT_FOR_DATASTORE
          value: "true"
        {{- /* Set based on the k8s node name. */}}
        - name: NODENAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        {{- /* Don't enable BGP. */}}
        - name: CALICO_NETWORKING_BACKEND
          value: "none"
        {{- /* Cluster type to identify the deployment type */}}
        - name: CLUSTER_TYPE
          value: "k8s"
        {{- /* The default IPv4 pool to create on startup if none exists. Pod IPs will be
        chosen from this range. Changing this value after installation will have
        no effect. This should fall within ` + "`" + `--cluster-cidr` + "`" + `. */}}
        - name: CALICO_IPV4POOL_CIDR
          value: "<kubeClusterCidr>"
        {{- /* Disable file logging so ` + "`" + `kubectl logs` + "`" + ` works. */}}
        - name: CALICO_DISABLE_FILE_LOGGING
          value: "true"
        {{- /* Set Felix endpoint to host default action to ACCEPT. */}}
        - name: FELIX_DEFAULTENDPOINTTOHOSTACTION
          value: "ACCEPT"
        {{- /* Disable IPv6 on Kubernetes. */}}
        - name: FELIX_IPV6SUPPORT
          value: "false"
        {{- /* Set Felix logging to "info" */}}
        - name: FELIX_LOGSEVERITYSCREEN
          value: {{ContainerConfig "logSeverityScreen"}}
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
      {{- /* Used by calico-node. */}}
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
      {{- /* Used to install CNI. */}}
      - name: cni-bin-dir
        hostPath:
          path: /opt/cni/bin
      - name: cni-net-dir
        hostPath:
          path: /etc/cni/net.d
      {{- /* Used to create per-pod Unix Domain Sockets */}}
      - name: policysync
        hostPath:
          type: DirectoryOrCreate
          path: /var/run/nodeagent
      {{- /* Used to install Flex Volume Driver */}}
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
    addonmanager.kubernetes.io/mode: "EnsureExists"
---
{{- /* Typha Horizontal Autoscaler ConfigMap */}}
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
{{- /* Typha Horizontal Autoscaler Deployment */}}
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: calico-typha-horizontal-autoscaler
  namespace: kube-system
  labels:
    k8s-app: calico-typha-autoscaler
    addonmanager.kubernetes.io/mode: "EnsureExists"
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
{{- /* Typha Horizontal Autoscaler Cluster Role */}}
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: typha-cpha
  labels:
    addonmanager.kubernetes.io/mode: "EnsureExists"
rules:
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["list"]

---
{{- /* Typha Horizontal Autoscaler Cluster Role Binding */}}
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: typha-cpha
  labels:
    addonmanager.kubernetes.io/mode: "EnsureExists"
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: typha-cpha
subjects:
- kind: ServiceAccount
  name: typha-cpha
  namespace: kube-system
---
{{- /* Typha Horizontal Autoscaler Role */}}
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: typha-cpha
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "EnsureExists"
rules:
- apiGroups: [""]
  resources: ["configmaps"]
  verbs: ["get"]
- apiGroups: ["extensions"]
  resources: ["deployments/scale"]
  verbs: ["get", "update"]

---
{{- /* Typha Horizontal Autoscaler Role Binding */}}
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: typha-cpha
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "EnsureExists"
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: typha-cpha
subjects:
- kind: ServiceAccount
  name: typha-cpha
  namespace: kube-system
---
{{- /* Typha Horizontal Autoscaler Service Account */}}
apiVersion: v1
kind: ServiceAccount
metadata:
  name: typha-cpha
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "EnsureExists"
`)

func k8sAddons115CalicoYamlBytes() ([]byte, error) {
	return _k8sAddons115CalicoYaml, nil
}

func k8sAddons115CalicoYaml() (*asset, error) {
	bytes, err := k8sAddons115CalicoYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/1.15/calico.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsAadDefaultAdminGroupRbacYaml = []byte(`kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: aad-default-admin-group
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
subjects:
- kind: Group
  name: {{ContainerConfig "adminGroupID"}}
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: ClusterRole
  name: cluster-admin
  apiGroup: rbac.authorization.k8s.io
`)

func k8sAddonsAadDefaultAdminGroupRbacYamlBytes() ([]byte, error) {
	return _k8sAddonsAadDefaultAdminGroupRbacYaml, nil
}

func k8sAddonsAadDefaultAdminGroupRbacYaml() (*asset, error) {
	bytes, err := k8sAddonsAadDefaultAdminGroupRbacYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/aad-default-admin-group-rbac.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsAadPodIdentityYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: aad-pod-id-nmi-service-account
  namespace: kube-system
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
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: azurepodidentityexceptions.aadpodidentity.k8s.io
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  group: aadpodidentity.k8s.io
  version: v1
  names:
    kind: AzurePodIdentityException
    singular: azurepodidentityexception
    plural: azurepodidentityexceptions
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
- apiGroups: ["apiextensions.k8s.io"]
  resources: ["customresourcedefinitions"]
  verbs: ["get", "list"]
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get", "list", "watch"]
- apiGroups: [""]
  resources: ["secrets"]
  verbs: ["get"]
- apiGroups: ["aadpodidentity.k8s.io"]
  resources: ["azureidentitybindings", "azureidentities", "azurepodidentityexceptions"]
  verbs: ["get", "list", "watch"]
- apiGroups: ["aadpodidentity.k8s.io"]
  resources: ["azureassignedidentities"]
  verbs: ["get", "list", "watch"]
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: aad-pod-id-nmi-binding
  labels:
    k8s-app: aad-pod-id-nmi-binding
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
subjects:
- kind: ServiceAccount
  name: aad-pod-id-nmi-service-account
  namespace: kube-system
roleRef:
  kind: ClusterRole
  name: aad-pod-id-nmi-role
  apiGroup: rbac.authorization.k8s.io
---
apiVersion: {{if IsKubernetesVersionGe "1.16.0"}}apps/v1{{else}}extensions/v1beta1{{end}}
kind: DaemonSet
metadata:
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
    component: nmi
    tier: node
    k8s-app: aad-pod-id
  name: nmi
  namespace: kube-system
spec:
  updateStrategy:
    type: RollingUpdate
{{- if IsKubernetesVersionGe "1.16.0"}}
  selector:
    matchLabels:
      component: nmi
      tier: node
{{- end}}
  template:
    metadata:
      labels:
        component: nmi
        tier: node
{{- if IsKubernetesVersionGe "1.17.0"}}
      annotations:
        cluster-autoscaler.kubernetes.io/daemonset-pod: "true"
{{- end}}
    spec:
      priorityClassName: system-cluster-critical
      serviceAccountName: aad-pod-id-nmi-service-account
      hostNetwork: true
      volumes:
      - hostPath:
          path: /run/xtables.lock
          type: FileOrCreate
        name: iptableslock
      containers:
      - name: nmi
        image: {{ContainerImage "nmi"}}
        imagePullPolicy: IfNotPresent
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
        resources:
          requests:
            cpu: {{ContainerCPUReqs "nmi"}}
            memory: {{ContainerMemReqs "nmi"}}
          limits:
            cpu: {{ContainerCPULimits "nmi"}}
            memory: {{ContainerMemLimits "nmi"}}
        securityContext:
          privileged: true
          capabilities:
            add:
            - NET_ADMIN
        volumeMounts:
        - mountPath: /run/xtables.lock
          name: iptableslock
        livenessProbe:
          httpGet:
            path: /healthz
            port: 8080
          initialDelaySeconds: 10
          periodSeconds: 5
      nodeSelector:
        kubernetes.io/os: linux
      tolerations:
      - key: {{GetAADPodIdentityTaintKey}}
        operator: Equal
        value: "true"
        effect: NoSchedule
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: aad-pod-id-mic-service-account
  namespace: kube-system
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
  resources: ["pods", "nodes"]
  verbs: [ "list", "watch" ]
- apiGroups: [""]
  resources: ["events"]
  verbs: ["create", "patch"]
- apiGroups: [""]
  resources: ["configmaps"]
  verbs: ["get", "create", "update"]
- apiGroups: [""]
  resources: ["endpoints"]
  verbs: ["create", "get","update"]
- apiGroups: ["aadpodidentity.k8s.io"]
  resources: ["azureidentitybindings", "azureidentities"]
  verbs: ["get", "list", "watch", "post", "update"]
- apiGroups: ["aadpodidentity.k8s.io"]
  resources: ["azurepodidentityexceptions"]
  verbs: ["list", "update"]
- apiGroups: ["aadpodidentity.k8s.io"]
  resources: ["azureassignedidentities"]
  verbs: ["*"]
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: aad-pod-id-mic-binding
  labels:
    k8s-app: aad-pod-id-mic-binding
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
subjects:
- kind: ServiceAccount
  name: aad-pod-id-mic-service-account
  namespace: kube-system
roleRef:
  kind: ClusterRole
  name: aad-pod-id-mic-role
  apiGroup: rbac.authorization.k8s.io
---
apiVersion: {{if IsKubernetesVersionGe "1.16.0"}}apps/v1{{else}}extensions/v1beta1{{end}}
kind: Deployment
metadata:
  labels:
    component: mic
    k8s-app: aad-pod-id
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: mic
  namespace: kube-system
spec:
  replicas: 2
{{- if IsKubernetesVersionGe "1.16.0"}}
  selector:
    matchLabels:
      component: mic
{{- end}}
  template:
    metadata:
      labels:
        component: mic
        app: mic
    spec:
      serviceAccountName: aad-pod-id-mic-service-account
      containers:
      - name: mic
        image: {{ContainerImage "mic"}}
        imagePullPolicy: IfNotPresent
        args:
          - "--cloudconfig=/etc/kubernetes/azure.json"
          - "--logtostderr"
        env:
        - name: MIC_POD_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace                
        resources:
          requests:
            cpu: {{ContainerCPUReqs "mic"}}
            memory: {{ContainerMemReqs "mic"}}
          limits:
            cpu: {{ContainerCPULimits "mic"}}
            memory: {{ContainerMemLimits "mic"}}
        volumeMounts:
        - name: k8s-azure-file
          mountPath: /etc/kubernetes/azure.json
          readOnly: true
        livenessProbe:
          httpGet:
            path: /healthz
            port: 8080
          initialDelaySeconds: 10
          periodSeconds: 5
      volumes:
      - name: k8s-azure-file
        hostPath:
          path: /etc/kubernetes/azure.json
      nodeSelector:
        kubernetes.io/os: linux
`)

func k8sAddonsAadPodIdentityYamlBytes() ([]byte, error) {
	return _k8sAddonsAadPodIdentityYaml, nil
}

func k8sAddonsAadPodIdentityYaml() (*asset, error) {
	bytes, err := k8sAddonsAadPodIdentityYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/aad-pod-identity.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsAciConnectorYaml = []byte(`apiVersion: v1
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
apiVersion: {{if IsKubernetesVersionGe "1.16.0"}}apps/v1{{else}}extensions/v1beta1{{end}}
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
{{- if IsKubernetesVersionGe "1.16.0"}}
  selector:
    matchLabels:
      app: aci-connector
{{- end}}
  template:
    metadata:
      labels:
        app: aci-connector
    spec:
      serviceAccountName: aci-connector
      nodeSelector:
        kubernetes.io/os: linux
      containers:
      - name: aci-connector
        image: {{ContainerImage "aci-connector"}}
        imagePullPolicy: IfNotPresent
        env:
        - name: KUBELET_PORT
          value: "10250"
        - name: AZURE_AUTH_LOCATION
          value: /etc/virtual-kubelet/credentials.json
        - name: ACI_RESOURCE_GROUP
          value: <rgName>
        - name: ACI_REGION
          value: {{ContainerConfig "region"}}
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

func k8sAddonsAciConnectorYamlBytes() ([]byte, error) {
	return _k8sAddonsAciConnectorYaml, nil
}

func k8sAddonsAciConnectorYaml() (*asset, error) {
	bytes, err := k8sAddonsAciConnectorYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/aci-connector.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsAntreaYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: antreaagentinfos.clusterinformation.antrea.tanzu.vmware.com
spec:
  group: clusterinformation.antrea.tanzu.vmware.com
  names:
    kind: AntreaAgentInfo
    plural: antreaagentinfos
    shortNames:
    - aai
    singular: antreaagentinfo
  scope: Cluster
  versions:
  - name: v1beta1
    served: true
    storage: true
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: antreacontrollerinfos.clusterinformation.antrea.tanzu.vmware.com
spec:
  group: clusterinformation.antrea.tanzu.vmware.com
  names:
    kind: AntreaControllerInfo
    plural: antreacontrollerinfos
    shortNames:
    - aci
    singular: antreacontrollerinfo
  scope: Cluster
  versions:
  - name: v1beta1
    served: true
    storage: true
---
apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: antctl
  namespace: kube-system
---
apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: antrea-agent
  namespace: kube-system
---
apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: antrea-controller
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: antctl
rules:
- apiGroups:
  - networking.antrea.tanzu.vmware.com
  resources:
  - networkpolicies
  - appliedtogroups
  - addressgroups
  verbs:
  - get
  - list
- apiGroups:
  - system.antrea.tanzu.vmware.com
  resources:
  - controllerinfos
  - agentinfos
  verbs:
  - get
- nonResourceURLs:
  - /agentinfo
  - /addressgroups
  - /appliedtogroups
  - /networkpolicies
  - /ovsflows
  - /podinterfaces
  verbs:
  - get
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: antrea-agent
rules:
- apiGroups:
  - ""
  resources:
  - nodes
  - pods
  verbs:
  - get
  - watch
  - list
- apiGroups:
  - clusterinformation.antrea.tanzu.vmware.com
  resources:
  - antreaagentinfos
  verbs:
  - get
  - create
  - update
  - delete
- apiGroups:
  - networking.antrea.tanzu.vmware.com
  resources:
  - networkpolicies
  - appliedtogroups
  - addressgroups
  verbs:
  - get
  - watch
  - list
- apiGroups:
  - authentication.k8s.io
  resources:
  - tokenreviews
  verbs:
  - create
- apiGroups:
  - authorization.k8s.io
  resources:
  - subjectaccessreviews
  verbs:
  - create
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: antrea-controller
rules:
- apiGroups:
  - ""
  resources:
  - nodes
  - pods
  - namespaces
  verbs:
  - get
  - watch
  - list
- apiGroups:
  - networking.k8s.io
  resources:
  - networkpolicies
  verbs:
  - get
  - watch
  - list
- apiGroups:
  - clusterinformation.antrea.tanzu.vmware.com
  resources:
  - antreacontrollerinfos
  verbs:
  - get
  - create
  - update
  - delete
- apiGroups:
  - clusterinformation.antrea.tanzu.vmware.com
  resources:
  - antreaagentinfos
  verbs:
  - list
  - delete
- apiGroups:
  - authentication.k8s.io
  resources:
  - tokenreviews
  verbs:
  - create
- apiGroups:
  - authorization.k8s.io
  resources:
  - subjectaccessreviews
  verbs:
  - create
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: antrea-agent-authentication-reader
  namespace: kube-system
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: extension-apiserver-authentication-reader
subjects:
- kind: ServiceAccount
  name: antrea-agent
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: antrea-controller-authentication-reader
  namespace: kube-system
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: extension-apiserver-authentication-reader
subjects:
- kind: ServiceAccount
  name: antrea-controller
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: antctl
  namespace: kube-system
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: antctl
subjects:
- kind: ServiceAccount
  name: antctl
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: antrea-agent
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: antrea-agent
subjects:
- kind: ServiceAccount
  name: antrea-agent
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: antrea-controller
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: antrea-controller
subjects:
- kind: ServiceAccount
  name: antrea-controller
  namespace: kube-system
---
apiVersion: v1
data:
  antrea-agent.conf: |
    # Name of the OpenVSwitch bridge antrea-agent will create and use.
    # Make sure it doesn't conflict with your existing OpenVSwitch bridges.
    #ovsBridge: br-int

    # Datapath type to use for the OpenVSwitch bridge created by Antrea. Supported values are:
    # - system
    # - netdev
    # 'system' is the default value and corresponds to the kernel datapath. Use 'netdev' to run
    # OVS in userspace mode. Userspace mode requires the tun device driver to be available.
    #ovsDatapathType: system

    # Name of the interface antrea-agent will create and use for host <--> pod communication.
    # Make sure it doesn't conflict with your existing interfaces.
    #hostGateway: gw0

    # Encapsulation mode for communication between Pods across Nodes, supported values:
    # - vxlan (default)
    # - geneve
    # - gre
    # - stt
    #tunnelType: vxlan

    # Default MTU to use for the host gateway interface and the network interface of each Pod. If
    # omitted, antrea-agent will default this value to 1450 to accommodate for tunnel encapsulate
    # overhead.
    #defaultMTU: 1450

    # Whether or not to enable IPsec encryption of tunnel traffic. IPsec encryption is only supported
    # for the GRE tunnel type.
    #enableIPSecTunnel: false
    # CIDR Range for services in cluster. It's required to support egress network policy, should
    # be set to the same value as the one specified by --service-cluster-ip-range for kube-apiserver.
    serviceCIDR: {{ContainerConfig "serviceCidr"}}

    # Determines how traffic is encapsulated. It has the following options
    # encap(default): Inter-node Pod traffic is always encapsulated and Pod to outbound traffic is masqueraded.
    # noEncap: Inter-node Pod traffic is not encapsulated, but Pod to outbound traffic is masqueraded.
    #          Underlying network must be capable of supporting Pod traffic across IP subnet.
    # hybrid: noEncap if worker Nodes on same subnet, otherwise encap.
    # networkPolicyOnly: Antrea enforces NetworkPolicy only, and utilizes CNI chaining and delegates Pod IPAM and connectivity to primary CNI.
    #
    trafficEncapMode: {{ContainerConfig "trafficEncapMode"}}

    # The port for the antrea-agent APIServer to serve on.
    # Note that if it's set to another value, the ` + "`" + `containerPort` + "`" + ` of the ` + "`" + `api` + "`" + ` port of the
    # ` + "`" + `antrea-agent` + "`" + ` container must be set to the same value.
    #apiPort: 10350

    # Enable metrics exposure via Prometheus. Initializes Prometheus metrics listener.
    #enablePrometheusMetrics: false
  antrea-cni.conflist: |
    {
        "cniVersion":"0.3.0",
        "name": "antrea",
        "plugins": [
            {
                "type": "antrea",
                "ipam": {
                    "type": "host-local"
                }
            },
            {
                "type": "portmap",
                "capabilities": {"portMappings": true}
            }
        ]
    }
  antrea-controller.conf: |
    # The port for the antrea-controller APIServer to serve on.
    # Note that if it's set to another value, the ` + "`" + `containerPort` + "`" + ` of the ` + "`" + `api` + "`" + ` port of the
    # ` + "`" + `antrea-controller` + "`" + ` container must be set to the same value.
    #apiPort: 10349

    # Enable metrics exposure via Prometheus. Initializes Prometheus metrics listener.
    #enablePrometheusMetrics: false
kind: ConfigMap
metadata:
  annotations: {}
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "EnsureExists"
  name: antrea-config-m8cb9g82tf
  namespace: kube-system
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: antrea
  namespace: kube-system
spec:
  ports:
  - port: 443
    protocol: TCP
    targetPort: api
  selector:
    app: antrea
    component: antrea-controller
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: antrea
    component: antrea-controller
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: antrea-controller
  namespace: kube-system
spec:
  replicas: 1
  selector:
    matchLabels:
      app: antrea
      component: antrea-controller
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: antrea
        component: antrea-controller
    spec:
      containers:
      - args:
        - --config
        - /etc/antrea/antrea-controller.conf
        - --logtostderr=false
        - --log_dir
        - /var/log/antrea
        - --alsologtostderr
        command:
        - antrea-controller
        env:
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        image: {{ContainerImage "antrea-controller"}}
        name: antrea-controller
        ports:
        - containerPort: 10349
          name: api
          protocol: TCP
        readinessProbe:
          failureThreshold: 5
          httpGet:
            host: 127.0.0.1
            path: /healthz
            port: api
            scheme: HTTPS
          initialDelaySeconds: 5
          periodSeconds: 10
          timeoutSeconds: 5
        resources:
          requests:
            cpu: {{ContainerCPUReqs "antrea-controller"}}
        volumeMounts:
        - mountPath: /etc/antrea/antrea-controller.conf
          name: antrea-config
          readOnly: true
          subPath: antrea-controller.conf
        - mountPath: /var/log/antrea
          name: host-var-log-antrea
      hostNetwork: true
      nodeSelector:
        kubernetes.io/os: linux
      priorityClassName: system-cluster-critical
      serviceAccountName: antrea-controller
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      - effect: NoSchedule
        key: node-role.kubernetes.io/master
      volumes:
      - configMap:
          name: antrea-config-m8cb9g82tf
        name: antrea-config
      - hostPath:
          path: /var/log/antrea
          type: DirectoryOrCreate
        name: host-var-log-antrea
---
apiVersion: apiregistration.k8s.io/v1
kind: APIService
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: v1beta1.networking.antrea.tanzu.vmware.com
spec:
  group: networking.antrea.tanzu.vmware.com
  groupPriorityMinimum: 100
  insecureSkipTLSVerify: true
  service:
    name: antrea
    namespace: kube-system
  version: v1beta1
  versionPriority: 100
---
apiVersion: apiregistration.k8s.io/v1
kind: APIService
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: v1beta1.system.antrea.tanzu.vmware.com
spec:
  group: system.antrea.tanzu.vmware.com
  groupPriorityMinimum: 100
  insecureSkipTLSVerify: true
  service:
    name: antrea
    namespace: kube-system
  version: v1beta1
  versionPriority: 100
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  labels:
    app: antrea
    component: antrea-agent
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: antrea-agent
  namespace: kube-system
spec:
  selector:
    matchLabels:
      app: antrea
      component: antrea-agent
  template:
    metadata:
      labels:
        app: antrea
        component: antrea-agent
    spec:
      containers:
      - args:
        - --config
        - /etc/antrea/antrea-agent.conf
        - --logtostderr=false
        - --log_dir
        - /var/log/antrea
        - --alsologtostderr
        command:
        - antrea-agent
        env:
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        image: {{ContainerImage "antrea-agent"}}
        livenessProbe:
          exec:
            command:
            - /bin/sh
            - -c
            - container_liveness_probe agent
          failureThreshold: 5
          initialDelaySeconds: 5
          periodSeconds: 10
          timeoutSeconds: 5
        name: antrea-agent
        ports:
        - containerPort: 10350
          name: api
          protocol: TCP
        readinessProbe:
          failureThreshold: 5
          httpGet:
            host: 127.0.0.1
            path: /healthz
            port: api
            scheme: HTTPS
          initialDelaySeconds: 5
          periodSeconds: 10
          timeoutSeconds: 5
        resources:
          requests:
            cpu: {{ContainerCPUReqs "antrea-agent"}}
        securityContext:
          privileged: true
        volumeMounts:
        - mountPath: /etc/antrea/antrea-agent.conf
          name: antrea-config
          readOnly: true
          subPath: antrea-agent.conf
        - mountPath: /var/run/antrea
          name: host-var-run-antrea
        - mountPath: /var/run/openvswitch
          name: host-var-run-antrea
          subPath: openvswitch
        - mountPath: /var/lib/cni
          name: host-var-run-antrea
          subPath: cni
        - mountPath: /var/log/antrea
          name: host-var-log-antrea
        - mountPath: /host/proc
          name: host-proc
          readOnly: true
        - mountPath: /host/var/run/netns
          mountPropagation: HostToContainer
          name: host-var-run-netns
          readOnly: true
        - mountPath: /run/xtables.lock
          name: xtables-lock
      - command:
        - start_ovs
        image: {{ContainerImage "antrea-ovs"}}
        livenessProbe:
          exec:
            command:
            - /bin/sh
            - -c
            - timeout 5 container_liveness_probe ovs
          initialDelaySeconds: 5
          periodSeconds: 5
        name: antrea-ovs
        resources:
          requests:
            cpu: {{ContainerCPUReqs "antrea-ovs"}}
        securityContext:
          capabilities:
            add:
            - SYS_NICE
            - NET_ADMIN
            - SYS_ADMIN
            - IPC_LOCK
        volumeMounts:
        - mountPath: /var/run/openvswitch
          name: host-var-run-antrea
          subPath: openvswitch
        - mountPath: /var/log/openvswitch
          name: host-var-log-antrea
          subPath: openvswitch
      hostNetwork: true
      initContainers:
      - command:
        - install_cni
        image: {{ContainerImage "install-cni"}}
        name: install-cni
        command: [{{ContainerConfig "installCniCmd"}}]
        resources:
          requests:
            cpu: {{ContainerCPUReqs "install-cni"}}
        securityContext:
          capabilities:
            add:
            - SYS_MODULE
        volumeMounts:
        - mountPath: /etc/antrea/antrea-cni.conflist
          name: antrea-config
          readOnly: true
          subPath: antrea-cni.conflist
        - mountPath: /host/etc/cni/net.d
          name: host-cni-conf
        - mountPath: /host/opt/cni/bin
          name: host-cni-bin
        - mountPath: /lib/modules
          name: host-lib-modules
          readOnly: true
        - mountPath: /sbin/depmod
          name: host-depmod
          readOnly: true
      nodeSelector:
        kubernetes.io/os: linux
      priorityClassName: system-node-critical
      serviceAccountName: antrea-agent
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      - effect: NoSchedule
        operator: Exists
      volumes:
      - configMap:
          name: antrea-config-m8cb9g82tf
        name: antrea-config
      - hostPath:
          path: /etc/cni/net.d
        name: host-cni-conf
      - hostPath:
          path: /opt/cni/bin
        name: host-cni-bin
      - hostPath:
          path: /proc
        name: host-proc
      - hostPath:
          path: /var/run/netns
        name: host-var-run-netns
      - hostPath:
          path: /var/run/antrea
          type: DirectoryOrCreate
        name: host-var-run-antrea
      - hostPath:
          path: /var/log/antrea
          type: DirectoryOrCreate
        name: host-var-log-antrea
      - hostPath:
          path: /lib/modules
        name: host-lib-modules
      - hostPath:
          path: /sbin/depmod
        name: host-depmod
      - hostPath:
          path: /run/xtables.lock
          type: FileOrCreate
        name: xtables-lock
  updateStrategy:
    type: RollingUpdate
`)

func k8sAddonsAntreaYamlBytes() ([]byte, error) {
	return _k8sAddonsAntreaYaml, nil
}

func k8sAddonsAntreaYaml() (*asset, error) {
	bytes, err := k8sAddonsAntreaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/antrea.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsArcOnboardingYaml = []byte(`---
apiVersion: v1
kind: Namespace
metadata:
  name: azure-arc-onboarding
  labels:
    addonmanager.kubernetes.io/mode: "EnsureExists"
---
apiVersion: v1
kind: Secret
metadata:
  name: azure-arc-onboarding
  namespace: azure-arc-onboarding
  labels:
    addonmanager.kubernetes.io/mode: "EnsureExists"
data:
  TENANT_ID: {{ContainerConfigBase64 "tenantID"}}
  SUBSCRIPTION_ID: {{ContainerConfigBase64 "subscriptionID"}}
  RESOURCE_GROUP: {{ContainerConfigBase64 "resourceGroup"}}
  CONNECTED_CLUSTER: {{ContainerConfigBase64 "clusterName"}}
  LOCATION: {{ContainerConfigBase64 "location"}}
  CLIENT_ID: {{ContainerConfigBase64 "clientID"}}
  CLIENT_SECRET: {{ContainerConfigBase64 "clientSecret"}}
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: azure-arc-onboarding
  namespace: azure-arc-onboarding
  labels:
    addonmanager.kubernetes.io/mode: "EnsureExists"
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: azure-arc-onboarding
  labels:
    addonmanager.kubernetes.io/mode: "EnsureExists"
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
  - kind: ServiceAccount
    name: azure-arc-onboarding
    namespace: azure-arc-onboarding
---
apiVersion: batch/v1
kind: Job
metadata:
  name: azure-arc-onboarding
  namespace: azure-arc-onboarding
  labels:
    addonmanager.kubernetes.io/mode: "EnsureExists"
spec:
  template:
    spec:
      serviceAccountName: azure-arc-onboarding
      nodeSelector:
        kubernetes.io/arch: amd64
        kubernetes.io/os: linux
      containers:
      - name: azure-arc-onboarding
        image: {{ContainerImage "azure-arc-onboarding"}}
        env:
        - name: TENANT_ID
          valueFrom:
            secretKeyRef:
              name: azure-arc-onboarding
              key: TENANT_ID
        - name: SUBSCRIPTION_ID
          valueFrom:
            secretKeyRef:
              name: azure-arc-onboarding
              key: SUBSCRIPTION_ID
        - name: RESOURCE_GROUP
          valueFrom:
            secretKeyRef:
              name: azure-arc-onboarding
              key: RESOURCE_GROUP
        - name: CONNECTED_CLUSTER
          valueFrom:
            secretKeyRef:
              name: azure-arc-onboarding
              key: CONNECTED_CLUSTER
        - name: LOCATION
          valueFrom:
            secretKeyRef:
              name: azure-arc-onboarding
              key: LOCATION
        - name: CLIENT_ID
          valueFrom:
            secretKeyRef:
              name: azure-arc-onboarding
              key: CLIENT_ID
        - name: CLIENT_SECRET
          valueFrom:
            secretKeyRef:
              name: azure-arc-onboarding
              key: CLIENT_SECRET
      restartPolicy: Never
  backoffLimit: 4
`)

func k8sAddonsArcOnboardingYamlBytes() ([]byte, error) {
	return _k8sAddonsArcOnboardingYaml, nil
}

func k8sAddonsArcOnboardingYaml() (*asset, error) {
	bytes, err := k8sAddonsArcOnboardingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/arc-onboarding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsAuditPolicyYaml = []byte(`apiVersion: audit.k8s.io/v1{{ if not (IsKubernetesVersionGe "1.16.0")}}beta1{{end}}
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

func k8sAddonsAuditPolicyYamlBytes() ([]byte, error) {
	return _k8sAddonsAuditPolicyYaml, nil
}

func k8sAddonsAuditPolicyYaml() (*asset, error) {
	bytes, err := k8sAddonsAuditPolicyYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/audit-policy.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsAzureCloudProviderYaml = []byte(`---
apiVersion: rbac.authorization.k8s.io/v1{{- if not (IsKubernetesVersionGe "1.16.0")}}beta1{{end}}
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
apiVersion: rbac.authorization.k8s.io/v1{{- if not (IsKubernetesVersionGe "1.16.0")}}beta1{{end}}
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
apiVersion: rbac.authorization.k8s.io/v1{{- if not (IsKubernetesVersionGe "1.16.0")}}beta1{{end}}
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
apiVersion: rbac.authorization.k8s.io/v1{{- if not (IsKubernetesVersionGe "1.16.0")}}beta1{{end}}
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
{{- if IsKubernetesVersionGe "1.15.0"}}
---
apiVersion: rbac.authorization.k8s.io/v1{{- if not (IsKubernetesVersionGe "1.16.0")}}beta1{{end}}
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
apiVersion: rbac.authorization.k8s.io/v1{{- if not (IsKubernetesVersionGe "1.16.0")}}beta1{{end}}
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
{{- end}}
{{- if UsesCloudControllerManager}}
---
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: default
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
  annotations:
    storageclass.beta.kubernetes.io/is-default-class: "true"
provisioner: disk.csi.azure.com
parameters:
  skuName: Standard_LRS
  kind: managed
  cachingMode: ReadOnly
reclaimPolicy: Delete
  {{- if IsKubernetesVersionGe "1.15.0"}}
allowVolumeExpansion: true
  {{- end}}
  {{- if HasAvailabilityZones}}
volumeBindingMode: WaitForFirstConsumer
allowedTopologies:
- matchLabelExpressions:
  - key: topology.disk.csi.azure.com/zone
    values: {{GetZones}}
  {{else}}
volumeBindingMode: Immediate
  {{- end}}
---
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: managed-premium
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
provisioner: disk.csi.azure.com
parameters:
  skuName: Premium_LRS
  kind: managed
  cachingMode: ReadOnly
reclaimPolicy: Delete
  {{- if IsKubernetesVersionGe "1.15.0"}}
allowVolumeExpansion: true
  {{- end}}
  {{- if HasAvailabilityZones}}
volumeBindingMode: WaitForFirstConsumer
allowedTopologies:
- matchLabelExpressions:
  - key: topology.disk.csi.azure.com/zone
    values: {{GetZones}}
  {{else}}
volumeBindingMode: Immediate
  {{- end}}
---
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: managed-standard
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
provisioner: disk.csi.azure.com
parameters:
  skuName: Standard_LRS
  kind: managed
  cachingMode: ReadOnly
reclaimPolicy: Delete
  {{- if IsKubernetesVersionGe "1.15.0"}}
allowVolumeExpansion: true
  {{- end}}
  {{- if HasAvailabilityZones}}
volumeBindingMode: WaitForFirstConsumer
allowedTopologies:
- matchLabelExpressions:
  - key: topology.disk.csi.azure.com/zone
    values: {{GetZones}}
  {{else}}
volumeBindingMode: Immediate
  {{- end}}
---
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: azurefile
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
provisioner: file.csi.azure.com
parameters:
  skuName: Standard_LRS
reclaimPolicy: Delete
  {{- if IsKubernetesVersionGe "1.15.0"}}
allowVolumeExpansion: true
  {{- end}}
volumeBindingMode: Immediate
{{else}}
  {{- if NeedsStorageAccountStorageClasses}}
---
apiVersion: storage.k8s.io/v1beta1
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
    {{- if not IsAzureStackCloud}}
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
    {{- end}}
  {{- end}}
  {{- if NeedsManagedDiskStorageClasses}}
---
apiVersion: storage.k8s.io/v1beta1
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
    {{- if not IsAzureStackCloud}}
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
    {{- end}}
  {{- end}}
{{- end}}
`)

func k8sAddonsAzureCloudProviderYamlBytes() ([]byte, error) {
	return _k8sAddonsAzureCloudProviderYaml, nil
}

func k8sAddonsAzureCloudProviderYaml() (*asset, error) {
	bytes, err := k8sAddonsAzureCloudProviderYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/azure-cloud-provider.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsAzureCniNetworkmonitorYaml = []byte(`{{- if IsKubernetesVersionGe "1.16.0"}}
apiVersion: apps/v1
{{else}}
apiVersion: extensions/v1beta1
{{- end}}
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
  updateStrategy:
    type: RollingUpdate
  template:
    metadata:
      labels:
        k8s-app: azure-cnms
{{- if not (IsKubernetesVersionGe "1.16.0")}}
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ''
{{- end}}
{{- if IsKubernetesVersionGe "1.17.0"}}
      annotations:
        cluster-autoscaler.kubernetes.io/daemonset-pod: "true"
{{- end}}
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
        kubernetes.io/os: linux
      containers:
        - name: azure-cnms
          image: {{ContainerImage "azure-cni-networkmonitor"}}
          imagePullPolicy: IfNotPresent
          securityContext:
            privileged: true
          resources:
            requests:
              cpu: {{ContainerCPUReqs "azure-cni-networkmonitor"}}
              memory: {{ContainerMemReqs "azure-cni-networkmonitor"}}
            limits:
              cpu: {{ContainerCPULimits "azure-cni-networkmonitor"}}
              memory: {{ContainerMemLimits "azure-cni-networkmonitor"}}
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

func k8sAddonsAzureCniNetworkmonitorYamlBytes() ([]byte, error) {
	return _k8sAddonsAzureCniNetworkmonitorYaml, nil
}

func k8sAddonsAzureCniNetworkmonitorYaml() (*asset, error) {
	bytes, err := k8sAddonsAzureCniNetworkmonitorYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/azure-cni-networkmonitor.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsAzureNetworkPolicyYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: azure-npm
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: {{GetMode}}
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: azure-npm
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: {{GetMode}}
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
    addonmanager.kubernetes.io/mode: {{GetMode}}
subjects:
  - kind: ServiceAccount
    name: azure-npm
    namespace: kube-system
roleRef:
  kind: ClusterRole
  name: azure-npm
  apiGroup: rbac.authorization.k8s.io
---
apiVersion: {{if IsKubernetesVersionGe "1.16.0"}}apps/v1{{else}}extensions/v1beta1{{end}}
kind: DaemonSet
metadata:
  name: azure-npm
  namespace: kube-system
  labels:
    app: azure-npm
    addonmanager.kubernetes.io/mode: {{GetMode}}
spec:
  selector:
    matchLabels:
      k8s-app: azure-npm
  updateStrategy:
    type: RollingUpdate
  template:
    metadata:
      labels:
        k8s-app: azure-npm
{{- if IsKubernetesVersionGe "1.16.0"}}
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ''
  {{- if IsKubernetesVersionGe "1.17.0"}}
        cluster-autoscaler.kubernetes.io/daemonset-pod: "true"
  {{- end}}
{{- end}}
    spec:
      priorityClassName: system-node-critical
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      nodeSelector:
        kubernetes.io/os: linux
      containers:
        - name: azure-npm
          image: {{ContainerImage "azure-npm-daemonset"}}
          securityContext:
            privileged: true
          resources:
            requests:
              cpu: {{ContainerCPUReqs "azure-npm-daemonset"}}
              memory: {{ContainerMemReqs "azure-npm-daemonset"}}
            limits:
              cpu: {{ContainerCPULimits "azure-npm-daemonset"}}
              memory: {{ContainerMemLimits "azure-npm-daemonset"}}
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
          - name: protocols
            mountPath: /etc/protocols
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
      - name: protocols
        hostPath:
          path: /etc/protocols
          type: File
      serviceAccountName: azure-npm
`)

func k8sAddonsAzureNetworkPolicyYamlBytes() ([]byte, error) {
	return _k8sAddonsAzureNetworkPolicyYaml, nil
}

func k8sAddonsAzureNetworkPolicyYaml() (*asset, error) {
	bytes, err := k8sAddonsAzureNetworkPolicyYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/azure-network-policy.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsAzurePolicyDeploymentYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  labels:
    admission.gatekeeper.sh/ignore: no-self-managing
    control-plane: controller-manager
    gatekeeper.sh/system: "yes"
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: gatekeeper-system
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.2.4
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: configs.config.gatekeeper.sh
spec:
  group: config.gatekeeper.sh
  names:
    kind: Config
    listKind: ConfigList
    plural: configs
    singular: config
  scope: Namespaced
  validation:
    openAPIV3Schema:
      description: Config is the Schema for the configs API
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          description: ConfigSpec defines the desired state of Config
          properties:
            sync:
              description: Configuration for syncing k8s objects
              properties:
                syncOnly:
                  description: If non-empty, only entries on this list will be replicated
                    into OPA
                  items:
                    properties:
                      group:
                        type: string
                      kind:
                        type: string
                      version:
                        type: string
                    type: object
                  type: array
              type: object
            validation:
              description: Configuration for validation
              properties:
                traces:
                  description: List of requests to trace. Both "user" and "kinds"
                    must be specified
                  items:
                    properties:
                      dump:
                        description: Also dump the state of OPA with the trace. Set
                          to ` + "`" + `All` + "`" + ` to dump everything.
                        type: string
                      kind:
                        description: Only trace requests of the following GroupVersionKind
                        properties:
                          group:
                            type: string
                          kind:
                            type: string
                          version:
                            type: string
                        type: object
                      user:
                        description: Only trace requests from the specified user
                        type: string
                    type: object
                  type: array
              type: object
          type: object
        status:
          description: ConfigStatus defines the observed state of Config
          type: object
      type: object
  version: v1alpha1
  versions:
  - name: v1alpha1
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
---
apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    gatekeeper.sh/system: "yes"
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: gatekeeper-admin
  namespace: gatekeeper-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: gatekeeper-manager-role
  namespace: gatekeeper-system
rules:
- apiGroups:
  - ""
  resources:
  - secrets
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: gatekeeper-manager-role
rules:
- apiGroups:
  - '*'
  resources:
  - '*'
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - apiextensions.k8s.io
  resources:
  - customresourcedefinitions
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - config.gatekeeper.sh
  resources:
  - configs
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - config.gatekeeper.sh
  resources:
  - configs/status
  verbs:
  - get
  - patch
  - update
- apiGroups:
  - constraints.gatekeeper.sh
  resources:
  - '*'
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - templates.gatekeeper.sh
  resources:
  - constrainttemplates
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - templates.gatekeeper.sh
  resources:
  - constrainttemplates/status
  verbs:
  - get
  - patch
  - update
- apiGroups:
  - admissionregistration.k8s.io
  resourceNames:
  - gatekeeper-validating-webhook-configuration
  resources:
  - validatingwebhookconfigurations
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  labels:
    gatekeeper.sh/system: "yes"
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: gatekeeper-manager-rolebinding
  namespace: gatekeeper-system
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: gatekeeper-manager-role
subjects:
- kind: ServiceAccount
  name: gatekeeper-admin
  namespace: gatekeeper-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    gatekeeper.sh/system: "yes"
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: gatekeeper-manager-rolebinding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: gatekeeper-manager-role
subjects:
- kind: ServiceAccount
  name: gatekeeper-admin
  namespace: gatekeeper-system
---
apiVersion: v1
kind: Secret
metadata:
  labels:
    gatekeeper.sh/system: "yes"
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
  name: gatekeeper-webhook-server-cert
  namespace: gatekeeper-system
---
apiVersion: v1
kind: Service
metadata:
  labels:
    gatekeeper.sh/system: "yes"
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: gatekeeper-webhook-service
  namespace: gatekeeper-system
spec:
  ports:
  - port: 443
    targetPort: 8443
  selector:
    control-plane: controller-manager
    gatekeeper.sh/system: "yes"
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    control-plane: controller-manager
    gatekeeper.sh/system: "yes"
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: gatekeeper-controller-manager
  namespace: gatekeeper-system
spec:
  replicas: 1
  selector:
    matchLabels:
      control-plane: controller-manager
      gatekeeper.sh/system: "yes"
  template:
    metadata:
      annotations:
        container.seccomp.security.alpha.kubernetes.io/manager: runtime/default
      labels:
        control-plane: controller-manager
        gatekeeper.sh/system: "yes"
    spec:
      containers:
      - args:
        - --port=8443
        - --logtostderr
        - --exempt-namespace=gatekeeper-system
        - --log-denies
        command:
        - /manager
        env:
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        image: {{ContainerImage "gatekeeper"}}
        resources:
          requests:
            cpu: {{ContainerCPUReqs "gatekeeper"}}
            memory: {{ContainerMemReqs "gatekeeper"}}
          limits:
            cpu: {{ContainerCPULimits "gatekeeper"}}
            memory: {{ContainerMemLimits "gatekeeper"}}
        imagePullPolicy: IfNotPresent
        livenessProbe:
          httpGet:
            path: /healthz
            port: 9090
        name: manager
        ports:
        - containerPort: 8443
          name: webhook-server
          protocol: TCP
        - containerPort: 8888
          name: metrics
          protocol: TCP
        - containerPort: 9090
          name: healthz
          protocol: TCP
        readinessProbe:
          httpGet:
            path: /readyz
            port: 9090
        securityContext:
          allowPrivilegeEscalation: false
          capabilities:
            drop:
            - all
          runAsGroup: 999
          runAsNonRoot: true
          runAsUser: 1000
        volumeMounts:
        - mountPath: /certs
          name: cert
          readOnly: true
      serviceAccountName: gatekeeper-admin
      terminationGracePeriodSeconds: 60
      volumes:
      - name: cert
        secret:
          defaultMode: 420
          secretName: gatekeeper-webhook-server-cert
      nodeSelector:
        kubernetes.io/os: linux
---
apiVersion: admissionregistration.k8s.io/v1beta1
kind: ValidatingWebhookConfiguration
metadata:
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
  name: gatekeeper-validating-webhook-configuration
webhooks:
- clientConfig:
    caBundle: Cg==
    service:
      name: gatekeeper-webhook-service
      namespace: gatekeeper-system
      path: /v1/admit
  failurePolicy: Ignore
  name: validation.gatekeeper.sh
  namespaceSelector:
    matchExpressions:
    - key: control-plane
      operator: DoesNotExist
    - key: admission.gatekeeper.sh/ignore
      operator: DoesNotExist
  rules:
  - apiGroups:
    - '*'
    apiVersions:
    - '*'
    operations:
    - CREATE
    - UPDATE
    resources:
    - '*'
  sideEffects: None
  timeoutSeconds: 5
- clientConfig:
    caBundle: Cg==
    service:
      name: gatekeeper-webhook-service
      namespace: gatekeeper-system
      path: /v1/admitlabel
  failurePolicy: Fail
  name: check-ignore-label.gatekeeper.sh
  rules:
  - apiGroups:
    - ""
    apiVersions:
    - '*'
    operations:
    - CREATE
    - UPDATE
    resources:
    - namespaces
  sideEffects: None
  timeoutSeconds: 5
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  creationTimestamp: null
  labels:
    controller-tools.k8s.io: "1.0"
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: constrainttemplates.templates.gatekeeper.sh
spec:
  group: templates.gatekeeper.sh
  names:
    kind: ConstraintTemplate
    plural: constrainttemplates
  scope: Cluster
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          properties:
            crd:
              properties:
                spec:
                  properties:
                    names:
                      properties:
                        kind:
                          type: string
                        shortNames:
                          items:
                            type: string
                          type: array
                      type: object
                    validation:
                      type: object
                  type: object
              type: object
            targets:
              items:
                properties:
                  libs:
                    items:
                      type: string
                    type: array
                  rego:
                    type: string
                  target:
                    type: string
                type: object
              type: array
          type: object
        status:
          properties:
            byPod:
              items:
                properties:
                  errors:
                    items:
                      properties:
                        code:
                          type: string
                        location:
                          type: string
                        message:
                          type: string
                      required:
                      - code
                      - message
                      type: object
                    type: array
                  id:
                    description: a unique identifier for the pod that wrote the status
                    type: string
                  observedGeneration:
                    format: int64
                    type: integer
                type: object
              type: array
            created:
              type: boolean
          type: object
  version: v1beta1
  versions:
  - name: v1beta1
    served: true
    storage: true
  - name: v1alpha1
    served: true
    storage: false
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: azure-policy
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: policy-agent
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups: ["constraints.gatekeeper.sh"]
  resources: ["*"]
  verbs: ["create", "delete", "update", "list", "get"]
- apiGroups: ["templates.gatekeeper.sh"]
  resources: ["constrainttemplates"]
  verbs: ["create", "delete", "update", "list", "get"]
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: policy-agent
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
subjects:
- kind: ServiceAccount
  name: azure-policy
  namespace: kube-system
roleRef:
  kind: ClusterRole
  name: policy-agent
  apiGroup: rbac.authorization.k8s.io
---
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: policy-pod-agent
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get"]
---
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: policy-pod-agent
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
subjects:
- kind: ServiceAccount
  name: azure-policy
  namespace: kube-system
roleRef:
  kind: Role
  name: policy-pod-agent
  apiGroup: rbac.authorization.k8s.io
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: azure-policy
    aadpodidbinding: policy-identity
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: azure-policy
  namespace: kube-system
spec:
  replicas: 1
  selector:
    matchLabels:
      app: azure-policy
  template:
    metadata:
      labels:
        app: azure-policy
      name: azure-policy
    spec:
      serviceAccountName: azure-policy
      containers:
      - name: azure-policy
        image: {{ContainerImage "azure-policy"}}
        resources:
          requests:
            cpu: {{ContainerCPUReqs "azure-policy"}}
            memory: {{ContainerMemReqs "azure-policy"}}
          limits:
            cpu: {{ContainerCPULimits "azure-policy"}}
            memory: {{ContainerMemLimits "azure-policy"}}
        imagePullPolicy: IfNotPresent
        env:
        - name: K8S_POLICY_PREFIX
          value: azurepolicy
        - name: RESOURCE_ID
          value: <resourceId>
        - name: RESOURCE_TYPE
          value: AKS Engine
        - name: ACS_CREDENTIAL_LOCATION
          value: /etc/acs/azure.json
        - name: DATAPLANE_ENDPOINT
          value: https://gov-prod-policy-data.trafficmanager.net
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: CURRENT_IMAGE
          value: {{ContainerImage "azure-policy"}}
        volumeMounts:
        - name: acs-credential
          mountPath: "/etc/acs/azure.json"
      volumes:
      - hostPath:
          path: /etc/kubernetes/azure.json
          type: File
        name: acs-credential
      nodeSelector:
        kubernetes.io/os: linux
`)

func k8sAddonsAzurePolicyDeploymentYamlBytes() ([]byte, error) {
	return _k8sAddonsAzurePolicyDeploymentYaml, nil
}

func k8sAddonsAzurePolicyDeploymentYaml() (*asset, error) {
	bytes, err := k8sAddonsAzurePolicyDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/azure-policy-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsAzurediskCsiDriverDeploymentYaml = []byte(`---
# Source: azuredisk-csi-driver/templates/serviceaccount-csi-azuredisk-controller.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: csi-azuredisk-controller-sa
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
---
# Source: azuredisk-csi-driver/templates/serviceaccount-csi-azuredisk-node.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: csi-azuredisk-node-sa
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
---
# Source: azuredisk-csi-driver/templates/crd-csi-node-info.yaml
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  creationTimestamp: null
  name: csinodeinfos.csi.storage.k8s.io
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  group: csi.storage.k8s.io
  names:
    kind: CSINodeInfo
    plural: csinodeinfos
  scope: Cluster
  validation:
    openAPIV3Schema:
      properties:
        csiDrivers:
          description: List of CSI drivers running on the node and their properties.
          items:
            properties:
              driver:
                description: The CSI driver that this object refers to.
                type: string
              nodeID:
                description: The node from the driver point of view.
                type: string
              topologyKeys:
                description: List of keys supported by the driver.
                items:
                  type: string
                type: array
          type: array
  version: v1alpha1
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
---
# Source: azuredisk-csi-driver/templates/rbac-csi-azuredisk-controller.yaml
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: azuredisk-external-provisioner-role
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch", "create", "delete"]
  - apiGroups: [""]
    resources: ["persistentvolumeclaims"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["storageclasses"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["events"]
    verbs: ["get", "list", "watch", "create", "update", "patch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["csinodes"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["nodes"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["coordination.k8s.io"]
    resources: ["leases"]
    verbs: ["get", "list", "watch", "create", "update", "patch"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshots"]
    verbs: ["get", "list"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshotcontents"]
    verbs: ["get", "list"]
---
# Source: azuredisk-csi-driver/templates/rbac-csi-azuredisk-controller.yaml
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: azuredisk-external-attacher-role
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: [""]
    resources: ["nodes"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["csi.storage.k8s.io"]
    resources: ["csinodeinfos"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["volumeattachments"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: ["coordination.k8s.io"]
    resources: ["leases"]
    verbs: ["get", "list", "watch", "create", "update", "patch"]
---
# Source: azuredisk-csi-driver/templates/rbac-csi-azuredisk-controller.yaml
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: azuredisk-external-snapshotter-role
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["persistentvolumeclaims"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["storageclasses"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["events"]
    verbs: ["list", "watch", "create", "update", "patch"]
  - apiGroups: [""]
    resources: ["secrets"]
    verbs: ["get", "list"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshotclasses"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshotcontents"]
    verbs: ["create", "get", "list", "watch", "update", "delete"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshots"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: ["apiextensions.k8s.io"]
    resources: ["customresourcedefinitions"]
    verbs: ["create", "list", "watch", "delete"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshotcontents/status"]
    verbs: ["update"]
  - apiGroups: ["coordination.k8s.io"]
    resources: ["leases"]
    verbs: ["get", "watch", "list", "delete", "update", "create"]
---
# Source: azuredisk-csi-driver/templates/rbac-csi-azuredisk-controller.yaml
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: azuredisk-external-resizer-role
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch", "update", "patch"]
  - apiGroups: [""]
    resources: ["persistentvolumeclaims"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["persistentvolumeclaims/status"]
    verbs: ["update", "patch"]
  - apiGroups: [""]
    resources: ["events"]
    verbs: ["list", "watch", "create", "update", "patch"]
  - apiGroups: ["coordination.k8s.io"]
    resources: ["leases"]
    verbs: ["get", "list", "watch", "create", "update", "patch"]
---
# Source: azuredisk-csi-driver/templates/rbac-csi-azuredisk-secret.yaml
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  namespace: kube-system
  name: csi-azuredisk-secret-role
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
  - apiGroups: [""]
    resources: ["secrets"]
    verbs: ["get", "list"]
---
# Source: azuredisk-csi-driver/templates/rbac-csi-azuredisk-controller.yaml
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: azuredisk-csi-provisioner-binding
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
subjects:
  - kind: ServiceAccount
    name: csi-azuredisk-controller-sa
    namespace: kube-system
roleRef:
  kind: ClusterRole
  name: azuredisk-external-provisioner-role
  apiGroup: rbac.authorization.k8s.io
---
# Source: azuredisk-csi-driver/templates/rbac-csi-azuredisk-controller.yaml
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: azuredisk-csi-attacher-binding
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
subjects:
  - kind: ServiceAccount
    name: csi-azuredisk-controller-sa
    namespace: kube-system
roleRef:
  kind: ClusterRole
  name: azuredisk-external-attacher-role
  apiGroup: rbac.authorization.k8s.io
---
# Source: azuredisk-csi-driver/templates/rbac-csi-azuredisk-controller.yaml
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: azuredisk-csi-snapshotter-binding
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
subjects:
  - kind: ServiceAccount
    name: csi-azuredisk-controller-sa
    namespace: kube-system
roleRef:
  kind: ClusterRole
  name: azuredisk-external-snapshotter-role
  apiGroup: rbac.authorization.k8s.io
---
# Source: azuredisk-csi-driver/templates/rbac-csi-azuredisk-controller.yaml
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: azuredisk-csi-resizer-role
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
subjects:
  - kind: ServiceAccount
    name: csi-azuredisk-controller-sa
    namespace: kube-system
roleRef:
  kind: ClusterRole
  name: azuredisk-external-resizer-role
  apiGroup: rbac.authorization.k8s.io
---
# Source: azuredisk-csi-driver/templates/rbac-csi-azuredisk-secret.yaml
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: csi-azuredisk-secret-binding
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
subjects:
  - kind: ServiceAccount
    name: csi-azuredisk-controller-sa
    namespace: kube-system
  - kind: ServiceAccount
    name: csi-azuredisk-node-sa
    namespace: kube-system
roleRef:
  kind: ClusterRole
  name: csi-azuredisk-secret-role
  apiGroup: rbac.authorization.k8s.io
{{if and (IsKubernetesVersionGe "1.18.0") HasWindows}}
---
# Source: azuredisk-csi-driver/templates/csi-azuredisk-node-windows.yaml
kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: csi-azuredisk-node-windows
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  selector:
    matchLabels:
      app: csi-azuredisk-node-windows
  template:
    metadata:
      labels:
        app: csi-azuredisk-node-windows
    spec:
      serviceAccountName: csi-azuredisk-node-sa
      nodeSelector:
        kubernetes.io/os: windows
      priorityClassName: system-node-critical
      tolerations:
        - operator: "Exists"
      containers:
        - name: liveness-probe
          volumeMounts:
            - mountPath: C:\csi
              name: plugin-dir
          image: {{ContainerImage "livenessprobe-windows"}}
          args:
            - "--csi-address=$(CSI_ENDPOINT)"
            - "--probe-timeout=3s"
            - "--health-port=29603"
            - "--v=5"
          env:
            - name: CSI_ENDPOINT
              value: unix://C:\\csi\\csi.sock
          resources:
            limits:
              cpu: {{ContainerCPULimits "livenessprobe-windows"}}
              memory: {{ContainerMemLimits "livenessprobe-windows"}}
            requests:
              cpu: {{ContainerCPUReqs "livenessprobe-windows"}}
              memory: {{ContainerMemReqs "livenessprobe-windows"}}
        - name: node-driver-registrar
          image: {{ContainerImage "csi-node-driver-registrar-windows"}}
          args:
            - "--v=5"
            - "--csi-address=$(CSI_ENDPOINT)"
            - "--kubelet-registration-path=C:\\var\\lib\\kubelet\\plugins\\disk.csi.azure.com\\csi.sock"
          env:
            - name: CSI_ENDPOINT
              value: unix://C:\\csi\\csi.sock
            - name: KUBE_NODE_NAME
              valueFrom:
                fieldRef:
                  fieldPath: spec.nodeName
          volumeMounts:
            - name: kubelet-dir
              mountPath: "C:\\var\\lib\\kubelet"
            - name: plugin-dir
              mountPath: C:\csi
            - name: registration-dir
              mountPath: C:\registration
          resources:
            limits:
              cpu: {{ContainerCPULimits "csi-node-driver-registrar-windows"}}
              memory: {{ContainerMemLimits "csi-node-driver-registrar-windows"}}
            requests:
              cpu: {{ContainerCPUReqs "csi-node-driver-registrar-windows"}}
              memory: {{ContainerMemReqs "csi-node-driver-registrar-windows"}}
        - name: azuredisk
          image: {{ContainerImage "azuredisk-csi"}}
          args:
            - "--v=5"
            - "--endpoint=$(CSI_ENDPOINT)"
            - "--nodeid=$(KUBE_NODE_NAME)"
            - "--metrics-address=0.0.0.0:29605"
          ports:
            - containerPort: 29603
              name: healthz
              protocol: TCP
            - containerPort: 29605
              name: metrics
              protocol: TCP
          livenessProbe:
            failureThreshold: 5
            httpGet:
              path: /healthz
              port: healthz
            initialDelaySeconds: 30
            timeoutSeconds: 10
            periodSeconds: 30
          env:
            - name: AZURE_CREDENTIAL_FILE
              value: "C:\\k\\azure.json"
            - name: CSI_ENDPOINT
              value: unix://C:\\csi\\csi.sock
            - name: KUBE_NODE_NAME
              valueFrom:
                fieldRef:
                  apiVersion: v1
                  fieldPath: spec.nodeName
          securityContext:
            privileged: true
          volumeMounts:
            - name: kubelet-dir
              mountPath: "C:\\var\\lib\\kubelet"
            - name: plugin-dir
              mountPath: C:\csi
            - name: azure-config
              mountPath: C:\k
            - name: csi-proxy-fs-pipe
              mountPath: \\.\pipe\csi-proxy-filesystem-v1alpha1
            - name: csi-proxy-disk-pipe
              mountPath: \\.\pipe\csi-proxy-disk-v1alpha1
            - name: csi-proxy-volume-pipe
              mountPath: \\.\pipe\csi-proxy-volume-v1alpha1
          resources:
            limits:
              cpu: {{ContainerCPULimits "azuredisk-csi"}}
              memory: {{ContainerMemLimits "azuredisk-csi"}}
            requests:
              cpu: {{ContainerCPUReqs "azuredisk-csi"}}
              memory: {{ContainerMemReqs "azuredisk-csi"}}
      volumes:
        - name: csi-proxy-fs-pipe
          hostPath:
            path: \\.\pipe\csi-proxy-filesystem-v1alpha1
            type: ""
        - name: csi-proxy-disk-pipe
          hostPath:
            path: \\.\pipe\csi-proxy-disk-v1alpha1
            type: ""
        - name: csi-proxy-volume-pipe
          hostPath:
            path: \\.\pipe\csi-proxy-volume-v1alpha1
            type: ""
        - name: registration-dir
          hostPath:
            path: C:\var\lib\kubelet\plugins_registry\
            type: Directory
        - name: kubelet-dir
          hostPath:
            path: C:\var\lib\kubelet\
            type: Directory
        - name: plugin-dir
          hostPath:
            path: C:\var\lib\kubelet\plugins\disk.csi.azure.com\
            type: DirectoryOrCreate
        - name: azure-config
          hostPath:
            path: C:\k
            type: Directory
{{end}}
{{if HasLinux}}
---
# Source: azuredisk-csi-driver/templates/csi-azuredisk-node.yaml
kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: csi-azuredisk-node
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  selector:
    matchLabels:
      app: csi-azuredisk-node
  template:
    metadata:
      labels:
        app: csi-azuredisk-node
    spec:
      hostNetwork: true
      serviceAccountName: csi-azuredisk-node-sa
      nodeSelector:
        kubernetes.io/os: linux
      priorityClassName: system-node-critical
      tolerations:
        - operator: "Exists"
      containers:
        - name: liveness-probe
          volumeMounts:
            - mountPath: /csi
              name: socket-dir
          image: {{ContainerImage "livenessprobe"}}
          args:
            - --csi-address=/csi/csi.sock
            - --connection-timeout=3s
            - --health-port=29603
            - --v=5
          resources:
            limits:
              cpu: {{ContainerCPULimits "livenessprobe"}}
              memory: {{ContainerMemLimits "livenessprobe"}}
            requests:
              cpu: {{ContainerCPUReqs "livenessprobe"}}
              memory: {{ContainerMemReqs "livenessprobe"}}
        - name: node-driver-registrar
          image: {{ContainerImage "csi-node-driver-registrar"}}
          args:
            - --csi-address=$(ADDRESS)
            - --kubelet-registration-path=$(DRIVER_REG_SOCK_PATH)
            - --v=5
          lifecycle:
            preStop:
              exec:
                command: ["/bin/sh", "-c", "rm -rf /registration/disk.csi.azure.com-reg.sock /csi/csi.sock"]
          env:
            - name: ADDRESS
              value: /csi/csi.sock
            - name: DRIVER_REG_SOCK_PATH
              value: /var/lib/kubelet/plugins/disk.csi.azure.com/csi.sock
          volumeMounts:
            - name: socket-dir
              mountPath: /csi
            - name: registration-dir
              mountPath: /registration
          resources:
            limits:
              cpu: {{ContainerCPULimits "csi-node-driver-registrar"}}
              memory: {{ContainerMemLimits "csi-node-driver-registrar"}}
            requests:
              cpu: {{ContainerCPUReqs "csi-node-driver-registrar"}}
              memory: {{ContainerMemReqs "csi-node-driver-registrar"}}
        - name: azuredisk
          image: {{ContainerImage "azuredisk-csi"}}
          args:
            - "--v=5"
            - "--endpoint=$(CSI_ENDPOINT)"
            - "--nodeid=$(KUBE_NODE_NAME)"
            - "--metrics-address=0.0.0.0:29605"
          ports:
            - containerPort: 29603
              name: healthz
              protocol: TCP
            - containerPort: 29605
              name: metrics
              protocol: TCP
          livenessProbe:
            failureThreshold: 5
            httpGet:
              path: /healthz
              port: healthz
            initialDelaySeconds: 30
            timeoutSeconds: 10
            periodSeconds: 30
          env:
            - name: AZURE_CREDENTIAL_FILE
              value: "/etc/kubernetes/azure.json"
            - name: CSI_ENDPOINT
              value: unix:///csi/csi.sock
            - name: KUBE_NODE_NAME
              valueFrom:
                fieldRef:
                  apiVersion: v1
                  fieldPath: spec.nodeName
          securityContext:
            privileged: true
          volumeMounts:
            - mountPath: /csi
              name: socket-dir
            - mountPath: /var/lib/kubelet/
              mountPropagation: Bidirectional
              name: mountpoint-dir
            - mountPath: /etc/kubernetes/
              name: azure-cred
            - mountPath: /var/lib/waagent/ManagedIdentity-Settings
              readOnly: true
              name: msi
            - mountPath: /dev
              name: device-dir
            - mountPath: /sys/bus/scsi/devices
              name: sys-devices-dir
            - mountPath: /sys/class/scsi_host/
              name: scsi-host-dir
          resources:
            limits:
              cpu: {{ContainerCPULimits "azuredisk-csi"}}
              memory: {{ContainerMemLimits "azuredisk-csi"}}
            requests:
              cpu: {{ContainerCPUReqs "azuredisk-csi"}}
              memory: {{ContainerMemReqs "azuredisk-csi"}}
      volumes:
        - hostPath:
            path: /var/lib/kubelet/plugins/disk.csi.azure.com
            type: DirectoryOrCreate
          name: socket-dir
        - hostPath:
            path: /var/lib/kubelet/
            type: DirectoryOrCreate
          name: mountpoint-dir
        - hostPath:
            path: /var/lib/kubelet/plugins_registry/
            type: DirectoryOrCreate
          name: registration-dir
        - hostPath:
            path: /etc/kubernetes/
            type: Directory
          name: azure-cred
        - hostPath:
            path: /var/lib/waagent/ManagedIdentity-Settings
          name: msi
        - hostPath:
            path: /dev
            type: Directory
          name: device-dir
        - hostPath:
            path: /sys/bus/scsi/devices
            type: Directory
          name: sys-devices-dir
        - hostPath:
            path: /sys/class/scsi_host/
            type: Directory
          name: scsi-host-dir
{{end}}
---
# Source: azuredisk-csi-driver/templates/csi-azuredisk-controller.yaml
kind: Deployment
apiVersion: apps/v1
metadata:
  name: csi-azuredisk-controller
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  replicas: {{CSIControllerReplicas}}
  selector:
    matchLabels:
      app: csi-azuredisk-controller
  template:
    metadata:
      labels:
        app: csi-azuredisk-controller
    spec:
      hostNetwork: true
      serviceAccountName: csi-azuredisk-controller-sa
      nodeSelector:
        kubernetes.io/os: linux
      priorityClassName: system-cluster-critical
      tolerations:
        - key: "node-role.kubernetes.io/master"
          operator: "Equal"
          value: "true"
          effect: "NoSchedule"
      containers:
        - name: csi-provisioner
          image: {{ContainerImage "csi-provisioner"}}
          args:
            - "--provisioner=disk.csi.azure.com"
            - "--feature-gates=Topology=true"
            - "--csi-address=$(ADDRESS)"
            - "--connection-timeout=15s"
            - "--v=5"
            - "--timeout=120s"
            - "--enable-leader-election"
            - "--leader-election-type=leases"
          env:
            - name: ADDRESS
              value: /csi/csi.sock
          volumeMounts:
            - mountPath: /csi
              name: socket-dir
          resources:
            limits:
              cpu: {{ContainerCPULimits "csi-provisioner"}}
              memory: {{ContainerMemLimits "csi-provisioner"}}
            requests:
              cpu: {{ContainerCPUReqs "csi-provisioner"}}
              memory: {{ContainerMemReqs "csi-provisioner"}}
        - name: csi-attacher
          image: {{ContainerImage "csi-attacher"}}
          args:
            - "-v=5"
            - "-csi-address=$(ADDRESS)"
            - "-timeout=120s"
            - "-leader-election"
            - "-leader-election-type=leases"
          env:
            - name: ADDRESS
              value: /csi/csi.sock
          volumeMounts:
          - mountPath: /csi
            name: socket-dir
          resources:
            limits:
              cpu: {{ContainerCPULimits "csi-attacher"}}
              memory: {{ContainerMemLimits "csi-attacher"}}
            requests:
              cpu: {{ContainerCPUReqs "csi-attacher"}}
              memory: {{ContainerMemReqs "csi-attacher"}}
        {{if ShouldEnableCSISnapshotFeature "azuredisk-csi-driver"}}
        - name: csi-snapshotter
          image: {{ContainerImage "csi-snapshotter"}}
          args:
            - "-csi-address=$(ADDRESS)"
            - "-leader-election"
            - "-v=5"
          env:
            - name: ADDRESS
              value: /csi/csi.sock
          volumeMounts:
            - name: socket-dir
              mountPath: /csi
          resources:
            limits:
              cpu: {{ContainerCPULimits "csi-snapshotter"}}
              memory: {{ContainerMemLimits "csi-snapshotter"}}
            requests:
              cpu: {{ContainerCPUReqs "csi-snapshotter"}}
              memory: {{ContainerMemReqs "csi-snapshotter"}}
        {{end}}
        {{if IsKubernetesVersionGe "1.16.0"}}
        - name: csi-resizer
          image: {{ContainerImage "csi-resizer"}}
          args:
            - "-csi-address=$(ADDRESS)"
            - "-v=5"
            - "-leader-election"
          env:
            - name: ADDRESS
              value: /csi/csi.sock
          volumeMounts:
            - name: socket-dir
              mountPath: /csi
          resources:
            limits:
              cpu: {{ContainerCPULimits "csi-resizer"}}
              memory: {{ContainerMemLimits "csi-resizer"}}
            requests:
              cpu: {{ContainerCPUReqs "csi-resizer"}}
              memory: {{ContainerMemReqs "csi-resizer"}}
        {{end}}
        - name: liveness-probe
          image: {{ContainerImage "livenessprobe"}}
          args:
            - --csi-address=/csi/csi.sock
            - --connection-timeout=3s
            - --health-port=29602
            - --v=5
          volumeMounts:
            - name: socket-dir
              mountPath: /csi
          resources:
            limits:
              cpu: {{ContainerCPULimits "livenessprobe"}}
              memory: {{ContainerMemLimits "livenessprobe"}}
            requests:
              cpu: {{ContainerCPUReqs "livenessprobe"}}
              memory: {{ContainerMemReqs "livenessprobe"}}
        - name: azuredisk
          image: {{ContainerImage "azuredisk-csi"}}
          args:
            - "--v=5"
            - "--endpoint=$(CSI_ENDPOINT)"
            - "--nodeid=$(KUBE_NODE_NAME)"
          ports:
            - containerPort: 29602
              name: healthz
              protocol: TCP
            - containerPort: 29604
              name: metrics
              protocol: TCP
          livenessProbe:
            failureThreshold: 5
            httpGet:
              path: /healthz
              port: healthz
            initialDelaySeconds: 30
            timeoutSeconds: 10
            periodSeconds: 30
          env:
            - name: AZURE_CREDENTIAL_FILE
              value: "/etc/kubernetes/azure.json"
            - name: CSI_ENDPOINT
              value: unix:///csi/csi.sock
          volumeMounts:
            - mountPath: /csi
              name: socket-dir
            - mountPath: /etc/kubernetes/
              name: azure-cred
            - mountPath: /var/lib/waagent/ManagedIdentity-Settings
              readOnly: true
              name: msi
          resources:
            limits:
              cpu: {{ContainerCPULimits "azuredisk-csi"}}
              memory: {{ContainerMemLimits "azuredisk-csi"}}
            requests:
              cpu: {{ContainerCPUReqs "azuredisk-csi"}}
              memory: {{ContainerMemReqs "azuredisk-csi"}}
      volumes:
        - name: socket-dir
          emptyDir: {}
        - name: azure-cred
          hostPath:
            path: /etc/kubernetes/
            type: Directory
        - name: msi
          hostPath:
            path: /var/lib/waagent/ManagedIdentity-Settings
{{if ShouldEnableCSISnapshotFeature "azuredisk-csi-driver"}}
---
# Source: azuredisk-csi-driver/templates/csi-snapshot-controller.yaml
kind: Deployment
apiVersion: apps/v1
metadata:
  name: csi-snapshot-controller
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  replicas: 1
  selector:
    matchLabels:
      app: csi-snapshot-controller
  template:
    metadata:
      labels:
        app: csi-snapshot-controller
    spec:
      serviceAccountName: csi-snapshot-controller-sa
      nodeSelector:
        kubernetes.io/os: linux
      priorityClassName: system-cluster-critical
      tolerations:
        - key: "node-role.kubernetes.io/master"
          operator: "Equal"
          value: "true"
          effect: "NoSchedule"
      containers:
        - name: csi-snapshot-controller
          image: {{ContainerImage "csi-snapshot-controller"}}
          args:
            - "--v=5"
            - "-leader-election"
          resources:
            limits:
              cpu: {{ContainerCPULimits "csi-snapshot-controller"}}
              memory: {{ContainerMemLimits "csi-snapshot-controller"}}
            requests:
              cpu: {{ContainerCPUReqs "csi-snapshot-controller"}}
              memory: {{ContainerMemReqs "csi-snapshot-controller"}}
---
# Source: azuredisk-csi-driver/templates/rbac-csi-snapshot-controller.yaml
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: csi-snapshot-controller-role
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["persistentvolumeclaims"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["storageclasses"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["events"]
    verbs: ["list", "watch", "create", "update", "patch"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshotclasses"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshotcontents"]
    verbs: ["create", "get", "list", "watch", "update", "delete"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshots"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshots/status"]
    verbs: ["update"]
---
# Source: azuredisk-csi-driver/templates/rbac-csi-snapshot-controller.yaml
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  namespace: kube-system
  name: csi-snapshot-controller-leaderelection-role
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
  - apiGroups: ["coordination.k8s.io"]
    resources: ["leases"]
    verbs: ["get", "watch", "list", "delete", "update", "create"]
---
# Source: azuredisk-csi-driver/templates/rbac-csi-snapshot-controller.yaml
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: csi-snapshot-controller-binding
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
subjects:
  - kind: ServiceAccount
    name: csi-snapshot-controller-sa
    namespace: kube-system
roleRef:
  kind: ClusterRole
  name: csi-snapshot-controller-role
  apiGroup: rbac.authorization.k8s.io
---
# Source: azuredisk-csi-driver/templates/rbac-csi-snapshot-controller.yaml
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: csi-snapshot-controller-leaderelection-binding
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
subjects:
  - kind: ServiceAccount
    name: csi-snapshot-controller-sa
    namespace: kube-system
roleRef:
  kind: ClusterRole
  name: csi-snapshot-controller-leaderelection-role
  apiGroup: rbac.authorization.k8s.io
---
# Source: azuredisk-csi-driver/templates/crd-csi-snapshot.yaml
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
    api-approved.kubernetes.io: "https://github.com/kubernetes-csi/external-snapshotter/pull/139"
  creationTimestamp: null
  name: volumesnapshots.snapshot.storage.k8s.io
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  group: snapshot.storage.k8s.io
  names:
    kind: VolumeSnapshot
    listKind: VolumeSnapshotList
    plural: volumesnapshots
    singular: volumesnapshot
  scope: Namespaced
  subresources:
    status: {}
  preserveUnknownFields: false
  validation:
    openAPIV3Schema:
      description: VolumeSnapshot is a user's request for either creating a point-in-time
        snapshot of a persistent volume, or binding to a pre-existing snapshot.
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
          type: string
        spec:
          description: 'spec defines the desired characteristics of a snapshot requested
            by a user. More info: https://kubernetes.io/docs/concepts/storage/volume-snapshots#volumesnapshots
            Required.'
          properties:
            source:
              description: source specifies where a snapshot will be created from.
                This field is immutable after creation. Required.
              properties:
                persistentVolumeClaimName:
                  description: persistentVolumeClaimName specifies the name of the
                    PersistentVolumeClaim object in the same namespace as the VolumeSnapshot
                    object where the snapshot should be dynamically taken from. This
                    field is immutable.
                  type: string
                volumeSnapshotContentName:
                  description: volumeSnapshotContentName specifies the name of a pre-existing
                    VolumeSnapshotContent object. This field is immutable.
                  type: string
              type: object
            volumeSnapshotClassName:
              description: 'volumeSnapshotClassName is the name of the VolumeSnapshotClass
                requested by the VolumeSnapshot. If not specified, the default snapshot
                class will be used if one exists. If not specified, and there is no
                default snapshot class, dynamic snapshot creation will fail. Empty
                string is not allowed for this field. TODO(xiangqian): a webhook validation
                on empty string. More info: https://kubernetes.io/docs/concepts/storage/volume-snapshot-classes'
              type: string
          required:
            - source
          type: object
        status:
          description: 'status represents the current information of a snapshot. NOTE:
            status can be modified by sources other than system controllers, and must
            not be depended upon for accuracy. Controllers should only use information
            from the VolumeSnapshotContent object after verifying that the binding
            is accurate and complete.'
          properties:
            boundVolumeSnapshotContentName:
              description: 'boundVolumeSnapshotContentName represents the name of
                the VolumeSnapshotContent object to which the VolumeSnapshot object
                is bound. If not specified, it indicates that the VolumeSnapshot object
                has not been successfully bound to a VolumeSnapshotContent object
                yet. NOTE: Specified boundVolumeSnapshotContentName alone does not
                mean binding       is valid. Controllers MUST always verify bidirectional
                binding between       VolumeSnapshot and VolumeSnapshotContent to
                avoid possible security issues.'
              type: string
            creationTime:
              description: creationTime is the timestamp when the point-in-time snapshot
                is taken by the underlying storage system. In dynamic snapshot creation
                case, this field will be filled in with the "creation_time" value
                returned from CSI "CreateSnapshotRequest" gRPC call. For a pre-existing
                snapshot, this field will be filled with the "creation_time" value
                returned from the CSI "ListSnapshots" gRPC call if the driver supports
                it. If not specified, it indicates that the creation time of the snapshot
                is unknown.
              format: date-time
              type: string
            error:
              description: error is the last observed error during snapshot creation,
                if any. This field could be helpful to upper level controllers(i.e.,
                application controller) to decide whether they should continue on
                waiting for the snapshot to be created based on the type of error
                reported.
              properties:
                message:
                  description: 'message is a string detailing the encountered error
                    during snapshot creation if specified. NOTE: message may be logged,
                    and it should not contain sensitive information.'
                  type: string
                time:
                  description: time is the timestamp when the error was encountered.
                  format: date-time
                  type: string
              type: object
            readyToUse:
              description: readyToUse indicates if a snapshot is ready to be used
                to restore a volume. In dynamic snapshot creation case, this field
                will be filled in with the "ready_to_use" value returned from CSI
                "CreateSnapshotRequest" gRPC call. For a pre-existing snapshot, this
                field will be filled with the "ready_to_use" value returned from the
                CSI "ListSnapshots" gRPC call if the driver supports it, otherwise,
                this field will be set to "True". If not specified, it means the readiness
                of a snapshot is unknown.
              type: boolean
            restoreSize:
              description: restoreSize represents the complete size of the snapshot
                in bytes. In dynamic snapshot creation case, this field will be filled
                in with the "size_bytes" value returned from CSI "CreateSnapshotRequest"
                gRPC call. For a pre-existing snapshot, this field will be filled
                with the "size_bytes" value returned from the CSI "ListSnapshots"
                gRPC call if the driver supports it. When restoring a volume from
                this snapshot, the size of the volume MUST NOT be smaller than the
                restoreSize if it is specified, otherwise the restoration will fail.
                If not specified, it indicates that the size is unknown.
              type: string
          type: object
      required:
        - spec
      type: object
  version: v1beta1
  versions:
    - name: v1beta1
      served: true
      storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
---
# Source: azuredisk-csi-driver/templates/crd-csi-snapshot.yaml
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
    api-approved.kubernetes.io: "https://github.com/kubernetes-csi/external-snapshotter/pull/139"
  creationTimestamp: null
  name: volumesnapshotclasses.snapshot.storage.k8s.io
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  group: snapshot.storage.k8s.io
  names:
    kind: VolumeSnapshotClass
    listKind: VolumeSnapshotClassList
    plural: volumesnapshotclasses
    singular: volumesnapshotclass
  scope: Cluster
  preserveUnknownFields: false
  validation:
    openAPIV3Schema:
      description: VolumeSnapshotClass specifies parameters that a underlying storage
        system uses when creating a volume snapshot. A specific VolumeSnapshotClass
        is used by specifying its name in a VolumeSnapshot object. VolumeSnapshotClasses
        are non-namespaced
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
          type: string
        deletionPolicy:
          description: deletionPolicy determines whether a VolumeSnapshotContent created
            through the VolumeSnapshotClass should be deleted when its bound VolumeSnapshot
            is deleted. Supported values are "Retain" and "Delete". "Retain" means
            that the VolumeSnapshotContent and its physical snapshot on underlying
            storage system are kept. "Delete" means that the VolumeSnapshotContent
            and its physical snapshot on underlying storage system are deleted. Required.
          enum:
            - Delete
            - Retain
          type: string
        driver:
          description: driver is the name of the storage driver that handles this
            VolumeSnapshotClass. Required.
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
          type: string
        parameters:
          additionalProperties:
            type: string
          description: parameters is a key-value map with storage driver specific
            parameters for creating snapshots. These values are opaque to Kubernetes.
          type: object
      required:
        - deletionPolicy
        - driver
      type: object
  version: v1beta1
  versions:
    - name: v1beta1
      served: true
      storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
---
# Source: azuredisk-csi-driver/templates/crd-csi-snapshot.yaml
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
    api-approved.kubernetes.io: "https://github.com/kubernetes-csi/external-snapshotter/pull/139"
  creationTimestamp: null
  name: volumesnapshotcontents.snapshot.storage.k8s.io
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  group: snapshot.storage.k8s.io
  names:
    kind: VolumeSnapshotContent
    listKind: VolumeSnapshotContentList
    plural: volumesnapshotcontents
    singular: volumesnapshotcontent
  scope: Cluster
  subresources:
    status: {}
  preserveUnknownFields: false
  validation:
    openAPIV3Schema:
      description: VolumeSnapshotContent represents the actual "on-disk" snapshot
        object in the underlying storage system
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
          type: string
        spec:
          description: spec defines properties of a VolumeSnapshotContent created
            by the underlying storage system. Required.
          properties:
            deletionPolicy:
              description: deletionPolicy determines whether this VolumeSnapshotContent
                and its physical snapshot on the underlying storage system should
                be deleted when its bound VolumeSnapshot is deleted. Supported values
                are "Retain" and "Delete". "Retain" means that the VolumeSnapshotContent
                and its physical snapshot on underlying storage system are kept. "Delete"
                means that the VolumeSnapshotContent and its physical snapshot on
                underlying storage system are deleted. In dynamic snapshot creation
                case, this field will be filled in with the "DeletionPolicy" field
                defined in the VolumeSnapshotClass the VolumeSnapshot refers to. For
                pre-existing snapshots, users MUST specify this field when creating
                the VolumeSnapshotContent object. Required.
              enum:
                - Delete
                - Retain
              type: string
            driver:
              description: driver is the name of the CSI driver used to create the
                physical snapshot on the underlying storage system. This MUST be the
                same as the name returned by the CSI GetPluginName() call for that
                driver. Required.
              type: string
            source:
              description: source specifies from where a snapshot will be created.
                This field is immutable after creation. Required.
              properties:
                snapshotHandle:
                  description: snapshotHandle specifies the CSI "snapshot_id" of a
                    pre-existing snapshot on the underlying storage system. This field
                    is immutable.
                  type: string
                volumeHandle:
                  description: volumeHandle specifies the CSI "volume_id" of the volume
                    from which a snapshot should be dynamically taken from. This field
                    is immutable.
                  type: string
              type: object
            volumeSnapshotClassName:
              description: name of the VolumeSnapshotClass to which this snapshot
                belongs.
              type: string
            volumeSnapshotRef:
              description: volumeSnapshotRef specifies the VolumeSnapshot object to
                which this VolumeSnapshotContent object is bound. VolumeSnapshot.Spec.VolumeSnapshotContentName
                field must reference to this VolumeSnapshotContent's name for the
                bidirectional binding to be valid. For a pre-existing VolumeSnapshotContent
                object, name and namespace of the VolumeSnapshot object MUST be provided
                for binding to happen. This field is immutable after creation. Required.
              properties:
                apiVersion:
                  description: API version of the referent.
                  type: string
                fieldPath:
                  description: 'If referring to a piece of an object instead of an
                    entire object, this string should contain a valid JSON/Go field
                    access statement, such as desiredState.manifest.containers[2].
                    For example, if the object reference is to a container within
                    a pod, this would take on a value like: "spec.containers{name}"
                    (where "name" refers to the name of the container that triggered
                    the event) or if no container name is specified "spec.containers[2]"
                    (container with index 2 in this pod). This syntax is chosen only
                    to have some well-defined way of referencing a part of an object.
                    TODO: this design is not final and this field is subject to change
                    in the future.'
                  type: string
                kind:
                  description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
                  type: string
                name:
                  description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                  type: string
                namespace:
                  description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                  type: string
                resourceVersion:
                  description: 'Specific resourceVersion to which this reference is
                    made, if any. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency'
                  type: string
                uid:
                  description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                  type: string
              type: object
          required:
            - deletionPolicy
            - driver
            - source
            - volumeSnapshotRef
          type: object
        status:
          description: status represents the current information of a snapshot.
          properties:
            creationTime:
              description: creationTime is the timestamp when the point-in-time snapshot
                is taken by the underlying storage system. In dynamic snapshot creation
                case, this field will be filled in with the "creation_time" value
                returned from CSI "CreateSnapshotRequest" gRPC call. For a pre-existing
                snapshot, this field will be filled with the "creation_time" value
                returned from the CSI "ListSnapshots" gRPC call if the driver supports
                it. If not specified, it indicates the creation time is unknown. The
                format of this field is a Unix nanoseconds time encoded as an int64.
                On Unix, the command ` + "`" + `date +%s%N` + "`" + ` returns the current time in nanoseconds
                since 1970-01-01 00:00:00 UTC.
              format: int64
              type: integer
            error:
              description: error is the latest observed error during snapshot creation,
                if any.
              properties:
                message:
                  description: 'message is a string detailing the encountered error
                    during snapshot creation if specified. NOTE: message may be logged,
                    and it should not contain sensitive information.'
                  type: string
                time:
                  description: time is the timestamp when the error was encountered.
                  format: date-time
                  type: string
              type: object
            readyToUse:
              description: readyToUse indicates if a snapshot is ready to be used
                to restore a volume. In dynamic snapshot creation case, this field
                will be filled in with the "ready_to_use" value returned from CSI
                "CreateSnapshotRequest" gRPC call. For a pre-existing snapshot, this
                field will be filled with the "ready_to_use" value returned from the
                CSI "ListSnapshots" gRPC call if the driver supports it, otherwise,
                this field will be set to "True". If not specified, it means the readiness
                of a snapshot is unknown.
              type: boolean
            restoreSize:
              description: restoreSize represents the complete size of the snapshot
                in bytes. In dynamic snapshot creation case, this field will be filled
                in with the "size_bytes" value returned from CSI "CreateSnapshotRequest"
                gRPC call. For a pre-existing snapshot, this field will be filled
                with the "size_bytes" value returned from the CSI "ListSnapshots"
                gRPC call if the driver supports it. When restoring a volume from
                this snapshot, the size of the volume MUST NOT be smaller than the
                restoreSize if it is specified, otherwise the restoration will fail.
                If not specified, it indicates that the size is unknown.
              format: int64
              minimum: 0
              type: integer
            snapshotHandle:
              description: snapshotHandle is the CSI "snapshot_id" of a snapshot on
                the underlying storage system. If not specified, it indicates that
                dynamic snapshot creation has either failed or it is still in progress.
              type: string
          type: object
      required:
        - spec
      type: object
  version: v1beta1
  versions:
    - name: v1beta1
      served: true
      storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
---
# Source: azuredisk-csi-driver/templates/serviceaccount-csi-snapshot-controller.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: csi-snapshot-controller-sa
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
{{end}}
---
# Source: azuredisk-csi-driver/templates/csi-azuredisk-driver.yaml
apiVersion: storage.k8s.io/v1beta1
kind: CSIDriver
metadata:
  name: disk.csi.azure.com
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  attachRequired: true
  podInfoOnMount: true
`)

func k8sAddonsAzurediskCsiDriverDeploymentYamlBytes() ([]byte, error) {
	return _k8sAddonsAzurediskCsiDriverDeploymentYaml, nil
}

func k8sAddonsAzurediskCsiDriverDeploymentYaml() (*asset, error) {
	bytes, err := k8sAddonsAzurediskCsiDriverDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/azuredisk-csi-driver-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsAzurefileCsiDriverDeploymentYaml = []byte(`---
# Source: azurefile-csi-driver/templates/serviceaccount-csi-azurefile-controller.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: csi-azurefile-controller-sa
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
---
# Source: azurefile-csi-driver/templates/serviceaccount-csi-azurefile-controller.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: csi-azurefile-node-sa
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
---
# Source: azurefile-csi-driver/templates/crd-csi-node-info.yaml
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  creationTimestamp: null
  name: csinodeinfos.csi.storage.k8s.io
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  group: csi.storage.k8s.io
  names:
    kind: CSINodeInfo
    plural: csinodeinfos
  scope: Cluster
  validation:
    openAPIV3Schema:
      properties:
        csiDrivers:
          description: List of CSI drivers running on the node and their properties.
          items:
            properties:
              driver:
                description: The CSI driver that this object refers to.
                type: string
              nodeID:
                description: The node from the driver point of view.
                type: string
              topologyKeys:
                description: List of keys supported by the driver.
                items:
                  type: string
                type: array
          type: array
  version: v1alpha1
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
---
# Source: azurefile-csi-driver/templates/rbac-csi-azurefile-controller.yaml
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: azurefile-external-provisioner-role
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch", "create", "delete"]
  - apiGroups: [""]
    resources: ["persistentvolumeclaims"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["storageclasses"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["events"]
    verbs: ["get", "list", "watch", "create", "update", "patch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["csinodes"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["nodes"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["coordination.k8s.io"]
    resources: ["leases"]
    verbs: ["get", "list", "watch", "create", "update", "patch"]
---
# Source: azurefile-csi-driver/templates/rbac-csi-azurefile-controller.yaml
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: azurefile-external-attacher-role
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: [""]
    resources: ["nodes"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["csi.storage.k8s.io"]
    resources: ["csinodeinfos"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["volumeattachments"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: ["coordination.k8s.io"]
    resources: ["leases"]
    verbs: ["get", "list", "watch", "create", "update", "patch"]
---
# Source: azurefile-csi-driver/templates/rbac-csi-azurefile-controller.yaml
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: azurefile-cluster-driver-registrar-role
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
  - apiGroups: ["apiextensions.k8s.io"]
    resources: ["customresourcedefinitions"]
    verbs: ["create", "list", "watch", "delete"]
  - apiGroups: ["csi.storage.k8s.io"]
    resources: ["csidrivers"]
    verbs: ["create", "delete"]
  - apiGroups: ["coordination.k8s.io"]
    resources: ["leases"]
    verbs: ["get", "list", "watch", "create", "update", "patch"]
---
# Source: azurefile-csi-driver/templates/rbac-csi-azurefile-controller.yaml
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: azurefile-external-snapshotter-role
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["persistentvolumeclaims"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["storageclasses"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["events"]
    verbs: ["list", "watch", "create", "update", "patch"]
  - apiGroups: [""]
    resources: ["secrets"]
    verbs: ["get", "list"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshotclasses"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshotcontents"]
    verbs: ["create", "get", "list", "watch", "update", "delete"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshots"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: ["apiextensions.k8s.io"]
    resources: ["customresourcedefinitions"]
    verbs: ["create", "list", "watch", "delete"]
  - apiGroups: ["coordination.k8s.io"]
    resources: ["leases"]
    verbs: ["get", "list", "watch", "create", "update", "patch"]
---
# Source: azurefile-csi-driver/templates/rbac-csi-azurefile-controller.yaml
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: azurefile-external-resizer-role
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch", "update", "patch"]
  - apiGroups: [""]
    resources: ["persistentvolumeclaims"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["persistentvolumeclaims/status"]
    verbs: ["update", "patch"]
  - apiGroups: [""]
    resources: ["events"]
    verbs: ["list", "watch", "create", "update", "patch"]
  - apiGroups: ["coordination.k8s.io"]
    resources: ["leases"]
    verbs: ["get", "list", "watch", "create", "update", "patch"]
---
# Source: azurefile-csi-driver/templates/rbac-csi-azurefile-controller.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: azure-cloud-provider-secret-getter
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
  - apiGroups: [""]
    resources: ["secrets"]
    resourceNames: ["azure-cloud-provider"]
    verbs: ["get"]
---
# Source: azurefile-csi-driver/templates/rbac-csi-azurefile-controller.yaml
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: azurefile-csi-provisioner-binding
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
subjects:
  - kind: ServiceAccount
    name: csi-azurefile-controller-sa
    namespace: kube-system
roleRef:
  kind: ClusterRole
  name: azurefile-external-provisioner-role
  apiGroup: rbac.authorization.k8s.io
---
# Source: azurefile-csi-driver/templates/rbac-csi-azurefile-controller.yaml
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: azurefile-csi-attacher-binding
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
subjects:
  - kind: ServiceAccount
    name: csi-azurefile-controller-sa
    namespace: kube-system
roleRef:
  kind: ClusterRole
  name: azurefile-external-attacher-role
  apiGroup: rbac.authorization.k8s.io
---
# Source: azurefile-csi-driver/templates/rbac-csi-azurefile-controller.yaml
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: azurefile-csi-driver-registrar-binding
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
subjects:
  - kind: ServiceAccount
    name: csi-azurefile-controller-sa
    namespace: kube-system
roleRef:
  kind: ClusterRole
  name: azurefile-cluster-driver-registrar-role
  apiGroup: rbac.authorization.k8s.io
---
# Source: azurefile-csi-driver/templates/rbac-csi-azurefile-controller.yaml
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: azurefile-csi-snapshotter-binding
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
subjects:
  - kind: ServiceAccount
    name: csi-azurefile-controller-sa
    namespace: kube-system
roleRef:
  kind: ClusterRole
  name: azurefile-external-snapshotter-role
  apiGroup: rbac.authorization.k8s.io
---
# Source: azurefile-csi-driver/templates/rbac-csi-azurefile-controller.yaml
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: azurefile-csi-resizer-role
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
subjects:
  - kind: ServiceAccount
    name: csi-azurefile-controller-sa
    namespace: kube-system
roleRef:
  kind: ClusterRole
  name: azurefile-external-resizer-role
  apiGroup: rbac.authorization.k8s.io
---
# Source: azurefile-csi-driver/templates/rbac-csi-azurefile-controller.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: azure-cloud-provider-secret-getter-controller
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
subjects:
  - kind: ServiceAccount
    name: csi-azurefile-controller-sa
    namespace: kube-system
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: azure-cloud-provider-secret-getter
---
# Source: azurefile-csi-driver/templates/rbac-csi-azurefile-controller.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: azure-cloud-provider-secret-getter-node
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
subjects:
  - kind: ServiceAccount
    name: csi-azurefile-node-sa
    namespace: kube-system
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: azure-cloud-provider-secret-getter
{{if and (IsKubernetesVersionGe "1.18.0") HasWindows}}
---
# Source: azurefile-csi-driver/templates/csi-azurefile-node-windows.yaml
kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: csi-azurefile-node-windows
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  selector:
    matchLabels:
      app: csi-azurefile-node-windows
  template:
    metadata:
      labels:
        app: csi-azurefile-node-windows
    spec:
      nodeSelector:
        kubernetes.io/os: windows
      priorityClassName: system-node-critical
      tolerations:
        - operator: "Exists"
      containers:
        - name: liveness-probe
          volumeMounts:
            - mountPath: C:\csi
              name: plugin-dir
          image: {{ContainerImage "livenessprobe-windows"}}
          args:
            - "--csi-address=$(CSI_ENDPOINT)"
            - "--probe-timeout=3s"
            - "--health-port=29613"
            - "--v=5"
          env:
            - name: CSI_ENDPOINT
              value: unix://C:\\csi\\csi.sock
          resources:
            limits:
              cpu: {{ContainerCPULimits "livenessprobe-windows"}}
              memory: {{ContainerMemLimits "livenessprobe-windows"}}
            requests:
              cpu: {{ContainerCPUReqs "livenessprobe-windows"}}
              memory: {{ContainerMemReqs "livenessprobe-windows"}}
        - name: node-driver-registrar
          image: {{ContainerImage "csi-node-driver-registrar-windows"}}
          args:
            - "--csi-address=$(CSI_ENDPOINT)"
            - "--kubelet-registration-path=C:\\var\\lib\\kubelet\\plugins\\file.csi.azure.com\\csi.sock"
            - "--v=5"
          env:
            - name: CSI_ENDPOINT
              value: unix://C:\\csi\\csi.sock
            - name: KUBE_NODE_NAME
              valueFrom:
                fieldRef:
                  fieldPath: spec.nodeName
          volumeMounts:
            - name: kubelet-dir
              mountPath: "C:\\var\\lib\\kubelet"
            - name: plugin-dir
              mountPath: C:\csi
            - name: registration-dir
              mountPath: C:\registration
          resources:
            limits:
              cpu: {{ContainerCPULimits "csi-node-driver-registrar-windows"}}
              memory: {{ContainerMemLimits "csi-node-driver-registrar-windows"}}
            requests:
              cpu: {{ContainerCPUReqs "csi-node-driver-registrar-windows"}}
              memory: {{ContainerMemReqs "csi-node-driver-registrar-windows"}}
        - name: azurefile
          image: {{ContainerImage "azurefile-csi"}}
          args:
            - "--v=5"
            - "--endpoint=$(CSI_ENDPOINT)"
            - "--nodeid=$(KUBE_NODE_NAME)"
          ports:
            - containerPort: 29613
              name: healthz
              protocol: TCP
          livenessProbe:
            failureThreshold: 5
            httpGet:
              path: /healthz
              port: healthz
            initialDelaySeconds: 30
            timeoutSeconds: 10
            periodSeconds: 30
          env:
            - name: AZURE_CREDENTIAL_FILE
              value: "C:\\k\\azure.json"
            - name: CSI_ENDPOINT
              value: unix://C:\\csi\\csi.sock
            - name: KUBE_NODE_NAME
              valueFrom:
                fieldRef:
                  apiVersion: v1
                  fieldPath: spec.nodeName
          imagePullPolicy:
          securityContext:
            privileged: true
          volumeMounts:
            - name: kubelet-dir
              mountPath: "C:\\var\\lib\\kubelet"
            - name: plugin-dir
              mountPath: C:\csi
            - name: azure-config
              mountPath: C:\k
            - name: csi-proxy-fs-pipe
              mountPath: \\.\pipe\csi-proxy-filesystem-v1alpha1
            - name: csi-proxy-smb-pipe
              mountPath: \\.\pipe\csi-proxy-smb-v1alpha1
          resources:
            limits:
              cpu: {{ContainerCPULimits "azurefile-csi"}}
              memory: {{ContainerMemLimits "azurefile-csi"}}
            requests:
              cpu: {{ContainerCPUReqs "azurefile-csi"}}
              memory: {{ContainerMemReqs "azurefile-csi"}}
      volumes:
        - name: csi-proxy-fs-pipe
          hostPath:
            path: \\.\pipe\csi-proxy-filesystem-v1alpha1
            type: ""
        - name: csi-proxy-smb-pipe
          hostPath:
            path: \\.\pipe\csi-proxy-smb-v1alpha1
            type: ""
        - name: registration-dir
          hostPath:
            path: C:\var\lib\kubelet\plugins_registry\
            type: Directory
        - name: kubelet-dir
          hostPath:
            path: C:\var\lib\kubelet\
            type: Directory
        - name: plugin-dir
          hostPath:
            path: C:\var\lib\kubelet\plugins\file.csi.azure.com\
            type: DirectoryOrCreate
        - name: azure-config
          hostPath:
            path: C:\k
            type: Directory
{{end}}
{{if HasLinux}}
---
# Source: azurefile-csi-driver/templates/csi-azurefile-node.yaml
kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: csi-azurefile-node
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  selector:
    matchLabels:
      app: csi-azurefile-node
  template:
    metadata:
      labels:
        app: csi-azurefile-node
    spec:
      hostNetwork: true
      serviceAccountName: csi-azurefile-node-sa
      nodeSelector:
        kubernetes.io/os: linux
      priorityClassName: system-node-critical
      tolerations:
        - operator: "Exists"
      containers:
        - name: liveness-probe
          volumeMounts:
            - mountPath: /csi
              name: socket-dir
          image: {{ContainerImage "livenessprobe"}}
          args:
            - --csi-address=/csi/csi.sock
            - --connection-timeout=3s
            - --health-port=29613
            - --v=5
          resources:
            limits:
              cpu: {{ContainerCPULimits "livenessprobe"}}
              memory: {{ContainerMemLimits "livenessprobe"}}
            requests:
              cpu: {{ContainerCPUReqs "livenessprobe"}}
              memory: {{ContainerMemReqs "livenessprobe"}}
        - name: node-driver-registrar
          image: {{ContainerImage "csi-node-driver-registrar"}}
          args:
            - --csi-address=$(ADDRESS)
            - --kubelet-registration-path=$(DRIVER_REG_SOCK_PATH)
            - --v=5
          lifecycle:
            preStop:
              exec:
                command: ["/bin/sh", "-c", "rm -rf /registration/file.csi.azure.com-reg.sock /csi/csi.sock"]
          env:
            - name: ADDRESS
              value: /csi/csi.sock
            - name: DRIVER_REG_SOCK_PATH
              value: /var/lib/kubelet/plugins/file.csi.azure.com/csi.sock
          volumeMounts:
            - name: socket-dir
              mountPath: /csi
            - name: registration-dir
              mountPath: /registration
          resources:
            limits:
              cpu: {{ContainerCPULimits "csi-node-driver-registrar"}}
              memory: {{ContainerMemLimits "csi-node-driver-registrar"}}
            requests:
              cpu: {{ContainerCPUReqs "csi-node-driver-registrar"}}
              memory: {{ContainerMemReqs "csi-node-driver-registrar"}}
        - name: azurefile
          image: {{ContainerImage "azurefile-csi"}}
          args:
            - "--v=5"
            - "--endpoint=$(CSI_ENDPOINT)"
            - "--nodeid=$(KUBE_NODE_NAME)"
          ports:
            - containerPort: 29613
              name: healthz
              protocol: TCP
          livenessProbe:
            failureThreshold: 5
            httpGet:
              path: /healthz
              port: healthz
            initialDelaySeconds: 30
            timeoutSeconds: 10
            periodSeconds: 30
          env:
            - name: AZURE_CREDENTIAL_FILE
              value: "/etc/kubernetes/azure.json"
            - name: CSI_ENDPOINT
              value: unix:///csi/csi.sock
            - name: KUBE_NODE_NAME
              valueFrom:
                fieldRef:
                  apiVersion: v1
                  fieldPath: spec.nodeName
          securityContext:
            privileged: true
          volumeMounts:
            - mountPath: /csi
              name: socket-dir
            - mountPath: /var/lib/kubelet/
              mountPropagation: Bidirectional
              name: mountpoint-dir
            - mountPath: /etc/kubernetes/
              name: azure-cred
            - mountPath: /var/lib/waagent/ManagedIdentity-Settings
              readOnly: true
              name: msi
            - mountPath: /dev
              name: device-dir
          resources:
            limits:
              cpu: {{ContainerCPULimits "azurefile-csi"}}
              memory: {{ContainerMemLimits "azurefile-csi"}}
            requests:
              cpu: {{ContainerCPUReqs "azurefile-csi"}}
              memory: {{ContainerMemReqs "azurefile-csi"}}
      volumes:
        - hostPath:
            path: /var/lib/kubelet/plugins/file.csi.azure.com
            type: DirectoryOrCreate
          name: socket-dir
        - hostPath:
            path: /var/lib/kubelet/
            type: DirectoryOrCreate
          name: mountpoint-dir
        - hostPath:
            path: /var/lib/kubelet/plugins_registry/
            type: DirectoryOrCreate
          name: registration-dir
        - hostPath:
            path: /etc/kubernetes/
            type: Directory
          name: azure-cred
        - hostPath:
            path: /var/lib/waagent/ManagedIdentity-Settings
          name: msi
        - hostPath:
            path: /dev
            type: Directory
          name: device-dir
{{end}}
---
# Source: azurefile-csi-driver/templates/csi-azurefile-controller.yaml
kind: Deployment
apiVersion: apps/v1
metadata:
  name: csi-azurefile-controller
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  replicas: {{CSIControllerReplicas}}
  selector:
    matchLabels:
      app: csi-azurefile-controller
  template:
    metadata:
      labels:
        app: csi-azurefile-controller
    spec:
      hostNetwork: true  # only required for MSI enabled cluster
      serviceAccountName: csi-azurefile-controller-sa
      nodeSelector:
        kubernetes.io/os: linux
      priorityClassName: system-cluster-critical
      tolerations:
        - key: "node-role.kubernetes.io/master"
          operator: "Equal"
          value: "true"
          effect: "NoSchedule"
      containers:
        - name: csi-provisioner
          image: {{ContainerImage "csi-provisioner"}}
          args:
            - "-v=5"
            - "--provisioner=file.csi.azure.com"
            - "--csi-address=$(ADDRESS)"
            - "--connection-timeout=15s"
            - "--enable-leader-election"
            - "--leader-election-type=leases"
          env:
            - name: ADDRESS
              value: /csi/csi.sock
          volumeMounts:
            - mountPath: /csi
              name: socket-dir
          resources:
            limits:
              cpu: {{ContainerCPULimits "csi-provisioner"}}
              memory: {{ContainerMemLimits "csi-provisioner"}}
            requests:
              cpu: {{ContainerCPUReqs "csi-provisioner"}}
              memory: {{ContainerMemReqs "csi-provisioner"}}
        - name: csi-attacher
          image: {{ContainerImage "csi-attacher"}}
          args:
            - "-v=5"
            - "-csi-address=$(ADDRESS)"
            - "-timeout=120s"
            - "-leader-election"
            - "-leader-election-type=leases"
          env:
            - name: ADDRESS
              value: /csi/csi.sock
          volumeMounts:
          - mountPath: /csi
            name: socket-dir
          resources:
            limits:
              cpu: {{ContainerCPULimits "csi-attacher"}}
              memory: {{ContainerMemLimits "csi-attacher"}}
            requests:
              cpu: {{ContainerCPUReqs "csi-attacher"}}
              memory: {{ContainerMemReqs "csi-attacher"}}
        {{if ShouldEnableCSISnapshotFeature "azurefile-csi-driver"}}
        - name: csi-snapshotter
          image: {{ContainerImage "csi-snapshotter"}}
          args:
            - "-v=5"
            - "-csi-address=$(ADDRESS)"
            - "-leader-election"
          env:
            - name: ADDRESS
              value: /csi/csi.sock
          volumeMounts:
            - name: socket-dir
              mountPath: /csi
          resources:
            limits:
              cpu: {{ContainerCPULimits "csi-snapshotter"}}
              memory: {{ContainerMemLimits "csi-snapshotter"}}
            requests:
              cpu: {{ContainerCPUReqs "csi-snapshotter"}}
              memory: {{ContainerMemReqs "csi-snapshotter"}}
        {{end}}
        {{if IsKubernetesVersionGe "1.16.0"}}
        - name: csi-resizer
          image: {{ContainerImage "csi-resizer"}}
          args:
            - "-csi-address=$(ADDRESS)"
            - "-v=5"
            - "-leader-election"
          env:
            - name: ADDRESS
              value: /csi/csi.sock
          volumeMounts:
            - name: socket-dir
              mountPath: /csi
          resources:
            limits:
              cpu: {{ContainerCPULimits "csi-resizer"}}
              memory: {{ContainerMemLimits "csi-resizer"}}
            requests:
              cpu: {{ContainerCPUReqs "csi-resizer"}}
              memory: {{ContainerMemReqs "csi-resizer"}}
        {{end}}
        - name: liveness-probe
          image: {{ContainerImage "livenessprobe"}}
          args:
            - --csi-address=/csi/csi.sock
            - --connection-timeout=3s
            - --health-port=29612
            - --v=5
          volumeMounts:
            - name: socket-dir
              mountPath: /csi
          resources:
            limits:
              cpu: {{ContainerCPULimits "livenessprobe"}}
              memory: {{ContainerMemLimits "livenessprobe"}}
            requests:
              cpu: {{ContainerCPUReqs "livenessprobe"}}
              memory: {{ContainerMemReqs "livenessprobe"}}
        - name: azurefile
          image: {{ContainerImage "azurefile-csi"}}
          args:
            - "--v=5"
            - "--endpoint=$(CSI_ENDPOINT)"
            - "--nodeid=$(KUBE_NODE_NAME)"
          ports:
            - containerPort: 29612
              name: healthz
              protocol: TCP
            - containerPort: 29614
              name: metrics
              protocol: TCP
          livenessProbe:
            failureThreshold: 5
            httpGet:
              path: /healthz
              port: healthz
            initialDelaySeconds: 30
            timeoutSeconds: 10
            periodSeconds: 30
          env:
            - name: AZURE_CREDENTIAL_FILE
              value: "/etc/kubernetes/azure.json"
            - name: CSI_ENDPOINT
              value: unix:///csi/csi.sock
          volumeMounts:
            - mountPath: /csi
              name: socket-dir
            - mountPath: /etc/kubernetes/
              name: azure-cred
            - mountPath: /var/lib/waagent/ManagedIdentity-Settings
              readOnly: true
              name: msi
          resources:
            limits:
              cpu: {{ContainerCPULimits "azurefile-csi"}}
              memory: {{ContainerMemLimits "azurefile-csi"}}
            requests:
              cpu: {{ContainerCPUReqs "azurefile-csi"}}
              memory: {{ContainerMemReqs "azurefile-csi"}}
      volumes:
        - name: socket-dir
          emptyDir: {}
        - name: azure-cred
          hostPath:
            path: /etc/kubernetes/
            type: Directory
        - name: msi
          hostPath:
            path: /var/lib/waagent/ManagedIdentity-Settings
---
# Source: azurefile-csi-driver/templates/csi-azurefile-driver.yaml
apiVersion: storage.k8s.io/v1beta1
kind: CSIDriver
metadata:
  name: file.csi.azure.com
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  attachRequired: true
  podInfoOnMount: true
`)

func k8sAddonsAzurefileCsiDriverDeploymentYamlBytes() ([]byte, error) {
	return _k8sAddonsAzurefileCsiDriverDeploymentYaml, nil
}

func k8sAddonsAzurefileCsiDriverDeploymentYaml() (*asset, error) {
	bytes, err := k8sAddonsAzurefileCsiDriverDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/azurefile-csi-driver-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsBlobfuseFlexvolumeYaml = []byte(`apiVersion: {{if IsKubernetesVersionGe "1.16.0"}}apps/v1{{else}}extensions/v1beta1{{end}}
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
{{- if IsKubernetesVersionGe "1.17.0"}}
      annotations:
        cluster-autoscaler.kubernetes.io/daemonset-pod: "true"
{{- end}}
    spec:
      priorityClassName: system-cluster-critical
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: agentpool
                operator: NotIn
                values:
                - flatcar
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
        kubernetes.io/os: linux
`)

func k8sAddonsBlobfuseFlexvolumeYamlBytes() ([]byte, error) {
	return _k8sAddonsBlobfuseFlexvolumeYaml, nil
}

func k8sAddonsBlobfuseFlexvolumeYaml() (*asset, error) {
	bytes, err := k8sAddonsBlobfuseFlexvolumeYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/blobfuse-flexvolume.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsCalicoYaml = []byte(`{{- /* Source: calico/templates/calico-config.yaml
This ConfigMap is used to configure a self-hosted Calico installation. */}}
kind: ConfigMap
apiVersion: v1
metadata:
  name: calico-config
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: "EnsureExists"
data:
  {{- /* You must set a non-zero value for Typha replicas below. */}}
  typha_service_name: "calico-typha"
  {{- /* The CNI network configuration to install on each node.  The special
  values in this config will be automatically populated. */}}
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
{{- /* Source: calico/templates/kdd-crds.yaml */}}
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: felixconfigurations.crd.projectcalico.org
  labels:
    addonmanager.kubernetes.io/mode: "EnsureExists"
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
    addonmanager.kubernetes.io/mode: "EnsureExists"
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
    addonmanager.kubernetes.io/mode: "EnsureExists"
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
    addonmanager.kubernetes.io/mode: "EnsureExists"
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
    addonmanager.kubernetes.io/mode: "EnsureExists"
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
    addonmanager.kubernetes.io/mode: "EnsureExists"
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
    addonmanager.kubernetes.io/mode: "EnsureExists"
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
    addonmanager.kubernetes.io/mode: "EnsureExists"
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
    addonmanager.kubernetes.io/mode: "EnsureExists"
spec:
  scope: Namespaced
  group: crd.projectcalico.org
  version: v1
  names:
    kind: NetworkSet
    plural: networksets
    singular: networkset
---
{{- /* Source: calico/templates/rbac.yaml
Include a clusterrole for the calico-node DaemonSet,
and bind it to the calico-node serviceaccount. */}}
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: calico-node
  labels:
    addonmanager.kubernetes.io/mode: "EnsureExists"
rules:
{{- /* The CNI plugin needs to get pods, nodes, and namespaces. */}}
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
  {{- /* Used to discover service IPs for advertisement. */}}
  - watch
  - list
  {{- /* Used to discover Typhas. */}}
  - get
- apiGroups: [""]
  resources:
  - nodes/status
  verbs:
  {{- /* Needed for clearing NodeNetworkUnavailable flag. */}}
  - patch
  {{- /* Calico stores some configuration information in node annotations. */}}
  - update
{{- /* Watch for changes to Kubernetes NetworkPolicies. */}}
- apiGroups: ["networking.k8s.io"]
  resources:
  - networkpolicies
  verbs:
  - watch
  - list
{{- /* Used by Calico for policy information. */}}
- apiGroups: [""]
  resources:
  - pods
  - namespaces
  - serviceaccounts
  verbs:
  - list
  - watch
{{- /* The CNI plugin patches pods/status. */}}
- apiGroups: [""]
  resources:
  - pods/status
  verbs:
  - patch
{{- /* Calico monitors various CRDs for config. */}}
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
{{- /* Calico must create and update some CRDs on startup. */}}
- apiGroups: ["crd.projectcalico.org"]
  resources:
  - ippools
  - felixconfigurations
  - clusterinformations
  verbs:
  - create
  - update
{{- /* Calico stores some configuration information on the node. */}}
- apiGroups: [""]
  resources:
  - nodes
  verbs:
  - get
  - list
  - watch
{{- /* These permissions are only requried for upgrade from v2.6, and can
be removed after upgrade or on fresh installations. */}}
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
    addonmanager.kubernetes.io/mode: "EnsureExists"
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: calico-node
subjects:
- kind: ServiceAccount
  name: calico-node
  namespace: kube-system

---
{{- /* Source: calico/templates/calico-typha.yaml
This manifest creates a Service, which will be backed by Calico's Typha daemon.
Typha sits in between Felix and the API server, reducing Calico's load on the API server. */}}
apiVersion: v1
kind: Service
metadata:
  name: calico-typha
  namespace: kube-system
  labels:
    k8s-app: calico-typha
    addonmanager.kubernetes.io/mode: "EnsureExists"
spec:
  ports:
  - port: 5473
    protocol: TCP
    targetPort: calico-typha
    name: calico-typha
  selector:
    k8s-app: calico-typha
---
{{- /* This manifest creates a Deployment of Typha to back the above service. */}}
apiVersion: apps/v1
kind: Deployment
metadata:
  name: calico-typha
  namespace: kube-system
  labels:
    k8s-app: calico-typha
    addonmanager.kubernetes.io/mode: "EnsureExists"
spec:
  {{- /* Number of Typha replicas.  To enable Typha, set this to a non-zero value *and* set the
  typha_service_name variable in the calico-config ConfigMap above.
  We recommend using Typha if you have more than 50 nodes.  Above 100 nodes it is essential
  (when using the Kubernetes datastore).  Use one replica for every 100-200 nodes.  In
  production, we recommend running at least 3 replicas to reduce the impact of rolling upgrade. */}}
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
        kubernetes.io/os: linux
      hostNetwork: true
      tolerations:
      {{- /* Mark the pod as a critical add-on for rescheduling. */}}
      - key: CriticalAddonsOnly
        operator: Exists
      {{- /* Since Calico can't network a pod until Typha is up, we need to run Typha itself
      as a host-networked pod. */}}
      serviceAccountName: calico-node
      containers:
      - image: {{ContainerImage "calico-typha"}}
        name: calico-typha
        ports:
        - containerPort: 5473
          name: calico-typha
          protocol: TCP
        env:
        {{- /* Enable "info" logging by default.  Can be set to "debug" to increase verbosity. */}}
        - name: TYPHA_LOGSEVERITYSCREEN
          value: "info"
        {{- /* Disable logging to file and syslog since those don't make sense in Kubernetes. */}}
        - name: TYPHA_LOGFILEPATH
          value: "none"
        - name: TYPHA_LOGSEVERITYSYS
          value: "none"
        {{- /* Monitor the Kubernetes API to find the number of running instances and rebalance
        connections. */}}
        - name: TYPHA_CONNECTIONREBALANCINGMODE
          value: "kubernetes"
        - name: TYPHA_DATASTORETYPE
          value: "kubernetes"
        - name: TYPHA_HEALTHENABLED
          value: "true"
        {{- /* Configure route aggregation based on pod CIDR. */}}
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
{{- /* Source: calico/templates/calico-node.yaml
This manifest installs the calico-node container, as well
as the CNI plugins and network config on
each master and worker node in a Kubernetes cluster. */}}
kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: calico-node
  namespace: kube-system
  labels:
    k8s-app: calico-node
    addonmanager.kubernetes.io/mode: "EnsureExists"
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
{{- if IsKubernetesVersionGe "1.17.0"}}
      annotations:
        cluster-autoscaler.kubernetes.io/daemonset-pod: "true"
{{- end}}
    spec:
      priorityClassName: system-cluster-critical
      nodeSelector:
        kubernetes.io/os: linux
      hostNetwork: true
      tolerations:
      {{- /* Make sure calico-node gets scheduled on all nodes. */}}
      - effect: NoSchedule
        operator: Exists
      {{- /* Mark the pod as a critical add-on for rescheduling. */}}
      - key: CriticalAddonsOnly
        operator: Exists
      - effect: NoExecute
        operator: Exists
      serviceAccountName: calico-node
      {{- /* Minimize downtime during a rolling upgrade or deletion; tell Kubernetes to do a "force deletion":
      https://kubernetes.io/docs/concepts/workloads/pods/pod/#termination-of-pods. */}}
      terminationGracePeriodSeconds: 0
      initContainers:
      {{- /* This container installs the CNI binaries
      and CNI network config file on each node. */}}
      - name: install-cni
        image: {{ContainerImage "calico-cni"}}
        command: ["/install-cni.sh"]
        env:
        {{- /* Name of the CNI config file to create. */}}
        - name: CNI_CONF_NAME
          value: "10-calico.conflist"
        {{- /* The CNI network config to install on each node. */}}
        - name: CNI_NETWORK_CONFIG
          valueFrom:
            configMapKeyRef:
              name: calico-config
              key: cni_network_config
        {{- /* Set the hostname based on the k8s node name. */}}
        - name: KUBERNETES_NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        {{- /* Prevents the container from sleeping forever. */}}
        - name: SLEEP
          value: "false"
        volumeMounts:
        - mountPath: /host/opt/cni/bin
          name: cni-bin-dir
        - mountPath: /host/etc/cni/net.d
          name: cni-net-dir
      containers:
      {{- /* Runs calico-node container on each Kubernetes node.  This
      container programs network policy and routes on each
      host. */}}
      - name: calico-node
        image: {{ContainerImage "calico-node"}}
        env:
        {{- /* Use Kubernetes API as the backing datastore. */}}
        - name: DATASTORE_TYPE
          value: "kubernetes"
        {{- /* Configure route aggregation based on pod CIDR. */}}
        - name: USE_POD_CIDR
          value: "true"
        {{- /* Typha support: controlled by the ConfigMap. */}}
        - name: FELIX_TYPHAK8SSERVICENAME
          valueFrom:
            configMapKeyRef:
              name: calico-config
              key: typha_service_name
        {{- /* Wait for the datastore. */}}
        - name: WAIT_FOR_DATASTORE
          value: "true"
        {{- /* Set based on the k8s node name. */}}
        - name: NODENAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        {{- /* Don't enable BGP. */}}
        - name: CALICO_NETWORKING_BACKEND
          value: "none"
        {{- /* Cluster type to identify the deployment type */}}
        - name: CLUSTER_TYPE
          value: "k8s"
        {{- /* The default IPv4 pool to create on startup if none exists. Pod IPs will be
        chosen from this range. Changing this value after installation will have
        no effect. This should fall within ` + "`" + `--cluster-cidr` + "`" + `. */}}
        - name: CALICO_IPV4POOL_CIDR
          value: "<kubeClusterCidr>"
        {{- /* Disable file logging so ` + "`" + `kubectl logs` + "`" + ` works. */}}
        - name: CALICO_DISABLE_FILE_LOGGING
          value: "true"
        {{- /* Set Felix endpoint to host default action to ACCEPT. */}}
        - name: FELIX_DEFAULTENDPOINTTOHOSTACTION
          value: "ACCEPT"
        {{- /* Disable IPv6 on Kubernetes. */}}
        - name: FELIX_IPV6SUPPORT
          value: "false"
        {{- /* Set Felix logging to "info" */}}
        - name: FELIX_LOGSEVERITYSCREEN
          value: {{ContainerConfig "logSeverityScreen"}}
        - name: FELIX_HEALTHENABLED
          value: "true"
        - name: CALICO_IPV4POOL_IPIP
          value: "off"
        - name: FELIX_INTERFACEPREFIX
          value: "azv"
        - name: FELIX_USAGEREPORTINGENABLED
          value: "{{ContainerConfig "usageReportingEnabled"}}"
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
      {{- /* Used by calico-node. */}}
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
      {{- /* Used to install CNI. */}}
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
    addonmanager.kubernetes.io/mode: "EnsureExists"
---
{{- /* Typha Horizontal Autoscaler ConfigMap */}}
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
{{- /* Typha Horizontal Autoscaler Deployment */}}
apiVersion: apps/v1
kind: Deployment
metadata:
  name: calico-typha-horizontal-autoscaler
  namespace: kube-system
  labels:
    k8s-app: calico-typha-autoscaler
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: "EnsureExists"
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
{{- /* Typha Horizontal Autoscaler Cluster Role */}}
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: typha-cpha
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: "EnsureExists"
rules:
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["list"]

---
{{- /* Typha Horizontal Autoscaler Cluster Role Binding */}}
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: typha-cpha
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: "EnsureExists"
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: typha-cpha
subjects:
- kind: ServiceAccount
  name: typha-cpha
  namespace: kube-system
---
{{- /* Typha Horizontal Autoscaler Role */}}
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: typha-cpha
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: "EnsureExists"
rules:
- apiGroups: [""]
  resources: ["configmaps"]
  verbs: ["get"]
- apiGroups: ["extensions"]
  resources: ["deployments/scale"]
  verbs: ["get", "update"]

---
{{- /* Typha Horizontal Autoscaler Role Binding */}}
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: typha-cpha
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: "EnsureExists"
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: typha-cpha
subjects:
- kind: ServiceAccount
  name: typha-cpha
  namespace: kube-system
---
{{- /* Typha Horizontal Autoscaler Service Account */}}
apiVersion: v1
kind: ServiceAccount
metadata:
  name: typha-cpha
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
`)

func k8sAddonsCalicoYamlBytes() ([]byte, error) {
	return _k8sAddonsCalicoYaml, nil
}

func k8sAddonsCalicoYaml() (*asset, error) {
	bytes, err := k8sAddonsCalicoYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/calico.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsCiliumYaml = []byte(`---
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
        cluster-autoscaler.kubernetes.io/daemonset-pod: "true"
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
        image: {{ContainerImage "cilium-agent"}}
        imagePullPolicy: IfNotPresent
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
        image: {{ContainerImage "clean-cilium-state"}}
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
        image: {{ContainerImage "cilium-operator"}}
        imagePullPolicy: IfNotPresent
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
        image: {{ContainerImage "cilium-etcd-operator"}}
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

func k8sAddonsCiliumYamlBytes() ([]byte, error) {
	return _k8sAddonsCiliumYaml, nil
}

func k8sAddonsCiliumYaml() (*asset, error) {
	bytes, err := k8sAddonsCiliumYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/cilium.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsCloudNodeManagerYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    k8s-app: cloud-node-manager
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: cloud-node-manager
  namespace: kube-system
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: cloud-node-manager
  labels:
    k8s-app: cloud-node-manager
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["watch","list","get","update", "patch"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: cloud-node-manager
  labels:
    k8s-app: cloud-node-manager
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cloud-node-manager
subjects:
- kind: ServiceAccount
  name: cloud-node-manager
  namespace: kube-system
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: cloud-node-manager
  namespace: kube-system
  labels:
    component: cloud-node-manager
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  selector:
    matchLabels:
      k8s-app: cloud-node-manager
  template:
    metadata:
      labels:
        k8s-app: cloud-node-manager
      annotations:
        cluster-autoscaler.kubernetes.io/daemonset-pod: "true"
    spec:
      priorityClassName: system-node-critical
      serviceAccountName: cloud-node-manager
      hostNetwork: true {{/* required to fetch correct hostname */}}
      nodeSelector:
        kubernetes.io/os: linux
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
      - name: cloud-node-manager
        image: {{ContainerImage "cloud-node-manager"}}
        imagePullPolicy: IfNotPresent
        command:
        - cloud-node-manager
        - --node-name=$(NODE_NAME)
        env:
        - name: NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        resources:
          requests:
            cpu: 50m
            memory: 50Mi
          limits:
            cpu: 2000m
            memory: 512Mi
{{- if and HasWindows (IsKubernetesVersionGe "1.18.0")}}
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: cloud-node-manager-windows
  namespace: kube-system
  labels:
    component: cloud-node-manager
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  selector:
    matchLabels:
      k8s-app: cloud-node-manager-windows
  template:
    metadata:
      labels:
        k8s-app: cloud-node-manager-windows
      annotations:
        cluster-autoscaler.kubernetes.io/daemonset-pod: "true"
    spec:
      priorityClassName: system-node-critical
      serviceAccountName: cloud-node-manager
      nodeSelector:
        kubernetes.io/os: windows
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
      - name: cloud-node-manager
        image: {{ContainerImage "cloud-node-manager"}}
        imagePullPolicy: IfNotPresent
        command:
        - /cloud-node-manager.exe
        - --node-name=$(NODE_NAME)
        env:
        - name: NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        resources:
          requests:
            cpu: 50m
            memory: 50Mi
          limits:
            cpu: 2000m
            memory: 512Mi
{{end}}
`)

func k8sAddonsCloudNodeManagerYamlBytes() ([]byte, error) {
	return _k8sAddonsCloudNodeManagerYaml, nil
}

func k8sAddonsCloudNodeManagerYaml() (*asset, error) {
	bytes, err := k8sAddonsCloudNodeManagerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/cloud-node-manager.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsClusterAutoscalerYaml = []byte(`---
apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    k8s-addon: cluster-autoscaler.addons.k8s.io
    k8s-app: cluster-autoscaler
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: {{GetMode}}
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
    addonmanager.kubernetes.io/mode: {{GetMode}}
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
{{- if not (IsKubernetesVersionGe "1.17.0")}}
- apiGroups: ["batch"]
  resources: ["jobs"]
  verbs: ["watch","list"]
{{- end}}
- apiGroups: ["storage.k8s.io"]
  resources: ["csinodes", "storageclasses"]
  verbs: ["get", "list", "watch"]
- apiGroups: ["batch"]
  resources: ["jobs", "cronjobs"]
  verbs: ["watch", "list", "get"]
{{- if IsKubernetesVersionGe "1.17.0"}}
- apiGroups: ["coordination.k8s.io"]
  resources: ["leases"]
  verbs: ["create"]
- apiGroups: ["coordination.k8s.io"]
  resourceNames: ["cluster-autoscaler"]
  resources: ["leases"]
  verbs: ["get", "update"]
{{- end}}
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
    addonmanager.kubernetes.io/mode: {{GetMode}}
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
    addonmanager.kubernetes.io/mode: {{GetMode}}
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
    addonmanager.kubernetes.io/mode: {{GetMode}}
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
  VMType: {{GetBase64EncodedVMType}}
kind: Secret
metadata:
  name: cluster-autoscaler-azure
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: {{GetMode}}
---
{{- if IsKubernetesVersionGe "1.16.0"}}
apiVersion: apps/v1
{{else}}
apiVersion: extensions/v1beta1
{{- end}}
kind: Deployment
metadata:
  labels:
    app: cluster-autoscaler
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: {{GetMode}}
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
      priorityClassName: system-node-critical{{GetHostNetwork}}
      serviceAccountName: cluster-autoscaler
      tolerations:
      - effect: NoSchedule
        operator: "Equal"
        value: "true"
        key: node-role.kubernetes.io/master
      nodeSelector:
        kubernetes.{{if IsKubernetesVersionGe "1.16.0"}}azure.com{{else}}io{{end}}/role: master
        kubernetes.io/os: linux
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
        - --logtostderr=true
        - --cloud-provider=azure
        - --skip-nodes-with-local-storage=false
        - --scan-interval={{ContainerConfig "scan-interval"}}
        - --expendable-pods-priority-cutoff={{ContainerConfig "expendable-pods-priority-cutoff"}}
        - --ignore-daemonsets-utilization={{ContainerConfig "ignore-daemonsets-utilization"}}
        - --ignore-mirror-pods-utilization={{ContainerConfig "ignore-mirror-pods-utilization"}}
        - --max-autoprovisioned-node-group-count={{ContainerConfig "max-autoprovisioned-node-group-count"}}
        - --max-empty-bulk-delete={{ContainerConfig "max-empty-bulk-delete"}}
        - --max-failing-time={{ContainerConfig "max-failing-time"}}
        - --max-graceful-termination-sec={{ContainerConfig "max-graceful-termination-sec"}}
        - --max-inactivity={{ContainerConfig "max-inactivity"}}
        - --max-node-provision-time={{ContainerConfig "max-node-provision-time"}}
        - --max-nodes-total={{ContainerConfig "max-nodes-total"}}
        - --max-total-unready-percentage={{ContainerConfig "max-total-unready-percentage"}}
        - --memory-total={{ContainerConfig "memory-total"}}
        - --min-replica-count={{ContainerConfig "min-replica-count"}}
        - --namespace=kube-system
        - --new-pod-scale-up-delay={{ContainerConfig "new-pod-scale-up-delay"}}
        - --node-autoprovisioning-enabled={{ContainerConfig "node-autoprovisioning-enabled"}}
        - --ok-total-unready-count={{ContainerConfig "ok-total-unready-count"}}
        - --scale-down-candidates-pool-min-count={{ContainerConfig "scale-down-candidates-pool-min-count"}}
        - --scale-down-candidates-pool-ratio={{ContainerConfig "scale-down-candidates-pool-ratio"}}
        - --scale-down-delay-after-add={{ContainerConfig "scale-down-delay-after-add"}}
        - --scale-down-delay-after-delete={{ContainerConfig "scale-down-delay-after-delete"}}
        - --scale-down-delay-after-failure={{ContainerConfig "scale-down-delay-after-failure"}}
        - --scale-down-enabled={{ContainerConfig "scale-down-enabled"}}
        - --scale-down-non-empty-candidates-count={{ContainerConfig "scale-down-non-empty-candidates-count"}}
        - --scale-down-unneeded-time={{ContainerConfig "scale-down-unneeded-time"}}
        - --scale-down-unready-time={{ContainerConfig "scale-down-unready-time"}}
        - --scale-down-utilization-threshold={{ContainerConfig "scale-down-utilization-threshold"}}
        - --skip-nodes-with-local-storage={{ContainerConfig "skip-nodes-with-local-storage"}}
        - --skip-nodes-with-system-pods={{ContainerConfig "skip-nodes-with-system-pods"}}
        - --stderrthreshold={{ContainerConfig "stderrthreshold"}}
        - --unremovable-node-recheck-timeout={{ContainerConfig "unremovable-node-recheck-timeout"}}
        - --v={{ContainerConfig "v"}}
        - --write-status-configmap={{ContainerConfig "write-status-configmap"}}
        - --balance-similar-node-groups={{ContainerConfig "balance-similar-node-groups"}}
{{GetClusterAutoscalerNodesConfig}}
        env:
        - name: ARM_CLOUD
          value: "{{GetCloud}}"
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
          value: "{{UseManagedIdentity}}"
        volumeMounts:
        - mountPath: /etc/ssl/certs/ca-certificates.crt
          name: ssl-certs
          readOnly: true{{GetVolumeMounts}}
{{- if IsKubernetesVersionGe "1.17.0"}}
      dnsPolicy: ClusterFirst
{{- end}}
      restartPolicy: Always
      volumes:
      - hostPath:
          path: /etc/ssl/certs/ca-certificates.crt
          type: ""
        name: ssl-certs{{GetVolumes}}
#EOF
`)

func k8sAddonsClusterAutoscalerYamlBytes() ([]byte, error) {
	return _k8sAddonsClusterAutoscalerYaml, nil
}

func k8sAddonsClusterAutoscalerYaml() (*asset, error) {
	bytes, err := k8sAddonsClusterAutoscalerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/cluster-autoscaler.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsContainerMonitoringYaml = []byte(`apiVersion: v1
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
  DOMAIN: "{{ContainerConfig "workspaceDomain"}}"
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
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: omsagent-reader
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
rules:
  - apiGroups: [""]
    resources: ["pods", "events", "nodes", "nodes/stats", "nodes/metrics", "nodes/spec", "nodes/proxy", "namespaces", "services"]
    verbs: ["list", "get", "watch"]
  - apiGroups: ["extensions", "apps"]
    resources: ["replicasets"]
    verbs: ["list"]
  - apiGroups: ["azmon.container.insights"]
    resources: ["healthstates"]
    verbs: ["get", "create", "patch"]
  - nonResourceURLs: ["/metrics"]
    verbs: ["get"]
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
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
    #fluent forward plugin
    <source>
     type forward
     port "#{ENV['HEALTHMODEL_REPLICASET_SERVICE_SERVICE_PORT']}"
     bind 0.0.0.0
     chunk_size_limit 4m
    </source>

     #Kubernetes pod inventory
    <source>
     type kubepodinventory
     tag oms.containerinsights.KubePodInventory
     run_interval 60
     log_level debug
     custom_metrics_azure_regions eastus,southcentralus,westcentralus,westus2,southeastasia,northeurope,westeurope,southafricanorth,centralus,northcentralus,eastus2,koreacentral,eastasia,centralindia,uksouth,canadacentral,francecentral,japaneast,australiaeast
    </source>

     #Kubernetes events
    <source>
     type kubeevents
     tag oms.containerinsights.KubeEvents
     run_interval 60
     log_level debug
    </source>

    #Kubernetes Nodes
    <source>
     type kubenodeinventory
     tag oms.containerinsights.KubeNodeInventory
     run_interval 60
     log_level debug
    </source>

    #Kubernetes health
    <source>
     type kubehealth
     tag kubehealth.ReplicaSet
     run_interval 60
     log_level debug
    </source>

    #cadvisor perf- Windows nodes
    <source>
     type wincadvisorperf
     tag oms.api.wincadvisorperf
     run_interval 60
     log_level debug
    </source>

    <filter mdm.kubenodeinventory**>
     type filter_inventory2mdm
     custom_metrics_azure_regions eastus,southcentralus,westcentralus,westus2,southeastasia,northeurope,westEurope,southafricanorth,centralus,northcentralus,eastus2,koreacentral,eastasia,centralindia,uksouth,canadacentral,francecentral,japaneast,australiaeast
     log_level info
    </filter>

    # custom_metrics_mdm filter plugin for perf data from windows nodes
    <filter mdm.cadvisorperf**>
     type filter_cadvisor2mdm
     custom_metrics_azure_regions eastus,southcentralus,westcentralus,westus2,southeastasia,northeurope,westEurope,southafricanorth,centralus,northcentralus,eastus2,koreacentral,eastasia,centralindia,uksouth,canadacentral,francecentral,japaneast,australiaeast
     metrics_to_collect cpuUsageNanoCores,memoryWorkingSetBytes
     log_level info
    </filter>

    #health model aggregation filter
    <filter kubehealth**>
     type filter_health_model_builder
    </filter>

    <match oms.containerinsights.KubePodInventory**>
     type out_oms
     log_level debug
     num_threads 5
     buffer_chunk_limit 4m
     buffer_type file
     buffer_path %STATE_DIR_WS%/out_oms_kubepods*.buffer
     buffer_queue_limit 20
     buffer_queue_full_action drop_oldest_chunk
     flush_interval 20s
     retry_limit 10
     retry_wait 5s
     max_retry_wait 5m
    </match>

    <match oms.containerinsights.KubeEvents**>
     type out_oms
     log_level debug
     num_threads 5
     buffer_chunk_limit 4m
     buffer_type file
     buffer_path %STATE_DIR_WS%/out_oms_kubeevents*.buffer
     buffer_queue_limit 20
     buffer_queue_full_action drop_oldest_chunk
     flush_interval 20s
     retry_limit 10
     retry_wait 5s
     max_retry_wait 5m
    </match>

    <match oms.containerinsights.KubeServices**>
     type out_oms
     log_level debug
     num_threads 2
     buffer_chunk_limit 4m
     buffer_type file
     buffer_path %STATE_DIR_WS%/out_oms_kubeservices*.buffer
     buffer_queue_limit 20
     buffer_queue_full_action drop_oldest_chunk
     flush_interval 20s
     retry_limit 10
     retry_wait 5s
     max_retry_wait 5m
    </match>

    <match oms.containerinsights.KubeNodeInventory**>
     type out_oms
     log_level debug
     num_threads 5
     buffer_chunk_limit 4m
     buffer_type file
     buffer_path %STATE_DIR_WS%/state/out_oms_kubenodes*.buffer
     buffer_queue_limit 20
     buffer_queue_full_action drop_oldest_chunk
     flush_interval 20s
     retry_limit 10
     retry_wait 5s
     max_retry_wait 5m
    </match>

    <match oms.containerinsights.ContainerNodeInventory**>
     type out_oms
     log_level debug
     num_threads 3
     buffer_chunk_limit 4m
     buffer_type file
     buffer_path %STATE_DIR_WS%/out_oms_containernodeinventory*.buffer
     buffer_queue_limit 20
     flush_interval 20s
     retry_limit 10
     retry_wait 5s
     max_retry_wait 5m
    </match>

    <match oms.api.KubePerf**>
     type out_oms
     log_level debug
     num_threads 5
     buffer_chunk_limit 4m
     buffer_type file
     buffer_path %STATE_DIR_WS%/out_oms_kubeperf*.buffer
     buffer_queue_limit 20
     buffer_queue_full_action drop_oldest_chunk
     flush_interval 20s
     retry_limit 10
     retry_wait 5s
     max_retry_wait 5m
    </match>

    <match mdm.kubepodinventory** mdm.kubenodeinventory** >
     type out_mdm
     log_level debug
     num_threads 5
     buffer_chunk_limit 4m
     buffer_type file
     buffer_path %STATE_DIR_WS%/out_mdm_*.buffer
     buffer_queue_limit 20
     buffer_queue_full_action drop_oldest_chunk
     flush_interval 20s
     retry_limit 10
     retry_wait 5s
     max_retry_wait 5m
     retry_mdm_post_wait_minutes 30
    </match>

    <match oms.api.wincadvisorperf**>
     type out_oms
     log_level debug
     num_threads 5
     buffer_chunk_limit 4m
     buffer_type file
     buffer_path %STATE_DIR_WS%/out_oms_api_wincadvisorperf*.buffer
     buffer_queue_limit 20
     buffer_queue_full_action drop_oldest_chunk
     flush_interval 20s
     retry_limit 10
     retry_wait 5s
     max_retry_wait 5m
    </match>

    <match mdm.cadvisorperf**>
     type out_mdm
     log_level debug
     num_threads 5
     buffer_chunk_limit 4m
     buffer_type file
     buffer_path %STATE_DIR_WS%/out_mdm_cdvisorperf*.buffer
     buffer_queue_limit 20
     buffer_queue_full_action drop_oldest_chunk
     flush_interval 20s
     retry_limit 10
     retry_wait 5s
     max_retry_wait 5m
     retry_mdm_post_wait_minutes 30
    </match>

    <match kubehealth.Signals**>
     type out_oms
     log_level debug
     num_threads 5
     buffer_chunk_limit 4m
     buffer_type file
     buffer_path %STATE_DIR_WS%/out_oms_kubehealth*.buffer
     buffer_queue_limit 20
     buffer_queue_full_action drop_oldest_chunk
     flush_interval 20s
     retry_limit 10
     retry_wait 5s
     max_retry_wait 5m
    </match>

    <match oms.api.InsightsMetrics**>
     type out_oms
     log_level debug
     num_threads 5
     buffer_chunk_limit 4m
     buffer_type file
     buffer_path %STATE_DIR_WS%/out_oms_insightsmetrics*.buffer
     buffer_queue_limit 20
     buffer_queue_full_action drop_oldest_chunk
     flush_interval 20s
     retry_limit 10
     retry_wait 5s
     max_retry_wait 5m
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
              value: "DaemonSet"
            - name: ISTEST
              value: "true"
            # Update this with the user assigned msi client id for omsagent addon (if exists)
            - name: USER_ASSIGNED_IDENTITY_CLIENT_ID
              value: ""
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
              readOnly: true
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
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - labelSelector:
              matchExpressions:
                - key: kubernetes.io/os
                  operator: In
                  values:
                  - linux
                - key: type
                  operator: NotIn
                  values:
                  - virtual-kubelet
      tolerations:
        - operator: "Exists"
          effect: "NoSchedule"
        - operator: "Exists"
          effect: "NoExecute"
        - operator: "Exists"
          effect: "PreferNoSchedule"
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
            # Update this with the user assigned msi client id for omsagent addon (if exists)
            - name: USER_ASSIGNED_IDENTITY_CLIENT_ID
              value: ""
          securityContext:
            privileged: true
          ports:
            - containerPort: 25225
              protocol: TCP
            - containerPort: 25224
              protocol: UDP
            - containerPort: 25227
              protocol: TCP
              name: in-rs-tcp
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
            - mountPath: /etc/config
              name: omsagent-rs-config
            - mountPath: /etc/config/settings
              name: settings-vol-config
              readOnly: true
          livenessProbe:
            exec:
              command:
                - /bin/bash
                - -c
                - /opt/livenessprobe.sh
            initialDelaySeconds: 60
            periodSeconds: 60
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - labelSelector:
              matchExpressions:
                - key: kubernetes.io/os
                  operator: In
                  values:
                  - linux
                - key: type
                  operator: NotIn
                  values:
                  - virtual-kubelet
      tolerations:
        - operator: "Exists"
          effect: "NoSchedule"
        - operator: "Exists"
          effect: "NoExecute"
        - operator: "Exists"
          effect: "PreferNoSchedule"
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
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: omsagent-win
  namespace: kube-system
  labels:
    component: oms-agent-win
    tier: node-win
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  updateStrategy:
    type: RollingUpdate
  selector:
    matchLabels:
      component: oms-agent-win
      tier: node-win
  template:
    metadata:
      labels:
        component: oms-agent-win
        tier: node-win
      annotations:
        agentVersion: {{ContainerConfig "omsAgentVersion"}}
        dockerProviderVersion: {{ContainerConfig "dockerProviderVersion"}}
        schema-versions:  {{ContainerConfig "schema-versions"}}
    spec:
      serviceAccountName: omsagent
      containers:
        - name: omsagent-win
          image: {{ContainerImage "omsagent-win"}}
          imagePullPolicy: IfNotPresent
          resources:
            limits:
              cpu: {{ContainerCPULimits "omsagent-win"}}
              memory: {{ContainerMemLimits "omsagent-win"}}
            requests:
              cpu: {{ContainerCPUReqs "omsagent-win"}}
              memory: {{ContainerMemReqs "omsagent-win"}}
          env:
            - name: ACS_RESOURCE_NAME
              value: {{ContainerConfig "clusterName"}}
            - name: CONTROLLER_TYPE
              value: "DaemonSet"
            - name: HOSTNAME
              valueFrom:
                fieldRef:
                  fieldPath: spec.nodeName
          volumeMounts:
            - mountPath: C:\ProgramData\docker\containers
              name: docker-windows-containers
              readOnly: true
            - mountPath: C:\var #Read + Write access on this for position file
              name: docker-windows-kuberenetes-container-logs
            - mountPath: C:\etc\config\settings
              name: settings-vol-config
              readOnly: true
            - mountPath: C:\etc\omsagent-secret
              name: omsagent-secret
              readOnly: true
          livenessProbe:
            exec:
              command:
                - cmd
                - /c
                - C:\opt\omsagentwindows\scripts\cmd\livenessProbe.cmd
            periodSeconds: 60
            initialDelaySeconds: 180
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: kubernetes.io/os
                operator: In
                values:
                - windows
      tolerations:
        - operator: "Exists"
          effect: "NoSchedule"
        - operator: "Exists"
          effect: "NoExecute"
        - operator: "Exists"
          effect: "PreferNoSchedule"
      volumes:
        - name: docker-windows-kuberenetes-container-logs
          hostPath:
            path: C:\var
        - name: docker-windows-containers
          hostPath:
            path: C:\ProgramData\docker\containers
        - name: settings-vol-config
          configMap:
            name: container-azm-ms-agentconfig
            optional: true
        - name: omsagent-secret
          secret:
            secretName: omsagent-secret
---
kind: Service
apiVersion: v1
metadata:
  name: healthmodel-replicaset-service
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  selector:
    rsName: "omsagent-rs"
  ports:
    - protocol: TCP
      port: 25227
      targetPort: in-rs-tcp
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: healthstates.azmon.container.insights
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  group: azmon.container.insights
  version: v1
  scope: Namespaced
  names:
    plural: healthstates
    kind: HealthState
`)

func k8sAddonsContainerMonitoringYamlBytes() ([]byte, error) {
	return _k8sAddonsContainerMonitoringYaml, nil
}

func k8sAddonsContainerMonitoringYaml() (*asset, error) {
	bytes, err := k8sAddonsContainerMonitoringYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/container-monitoring.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
      addonmanager.kubernetes.io/mode: EnsureExists
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    kubernetes.io/bootstrapping: rbac-defaults
    addonmanager.kubernetes.io/mode: EnsureExists
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
    addonmanager.kubernetes.io/mode: EnsureExists
data:
  Corefile: |
    import conf.d/Corefile*
    .:53 {
        errors
        health
        ready
        kubernetes {{ContainerConfig "domain"}} in-addr.arpa ip6.arpa {
            pods insecure
            fallthrough in-addr.arpa ip6.arpa
        }
        prometheus :9153
        forward . /etc/resolv.conf
        cache 30
        loop
        reload
        loadbalance
        import custom/*.override
    }
    import custom/*.server
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
    {{- /*
    See https://github.com/coredns/coredns/tree/master/plugin/azure for information
    about the Azure DNS plugin. */}}
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
    addonmanager.kubernetes.io/mode: EnsureExists
spec:
  {{- /* replicas: not specified here:
  1. In order to make Addon Manager do not reconcile this replicas parameter.
  2. Default is 1.
  3. Will be tuned in real time if DNS horizontal auto-scaling is turned on. */}}
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
        kubernetes.io/os: linux
        {{- if ContainerConfig "use-host-network"}}
        kubernetes.{{if IsKubernetesVersionGe "1.16.0"}}azure.com{{else}}io{{end}}/role: agent
        {{end}}
      containers:
      - name: coredns
        image: {{ContainerImage "coredns"}}
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
      {{- if ContainerConfig "use-host-network"}}
      hostNetwork: {{ContainerConfig "use-host-network"}}
      {{end}}
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
    addonmanager.kubernetes.io/mode: EnsureExists
spec:
  selector:
    k8s-app: kube-dns
  clusterIP: {{ContainerConfig "clusterIP"}}
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
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: coredns-autoscaler
  namespace: kube-system
  labels:
    k8s-addon: coredns.addons.k8s.io
    addonmanager.kubernetes.io/mode: EnsureExists
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    k8s-addon: coredns.addons.k8s.io
    addonmanager.kubernetes.io/mode: EnsureExists
  name: coredns-autoscaler
rules:
  - apiGroups: [""]
    resources: ["nodes"]
    verbs: ["list","watch"]
  - apiGroups: [""]
    resources: ["replicationcontrollers/scale"]
    verbs: ["get", "update"]
  - apiGroups: ["extensions", "apps"]
    resources: ["deployments/scale", "replicasets/scale"]
    verbs: ["get", "update"]
  - apiGroups: [""]
    resources: ["configmaps"]
    verbs: ["get", "create"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    k8s-addon: coredns.addons.k8s.io
    addonmanager.kubernetes.io/mode: EnsureExists
  name: coredns-autoscaler
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: coredns-autoscaler
subjects:
- kind: ServiceAccount
  name: coredns-autoscaler
  namespace: kube-system
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: coredns-autoscaler
  namespace: kube-system
  labels:
    k8s-app: coredns-autoscaler
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: EnsureExists
spec:
  selector:
    matchLabels:
      k8s-app: coredns-autoscaler
  template:
    metadata:
      labels:
        k8s-app: coredns-autoscaler
    spec:
      priorityClassName: system-cluster-critical
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
        kubernetes.io/os: linux
      containers:
      - name: autoscaler
        image: {{ContainerImage "coredns-autoscaler"}}
        resources:
          requests:
            cpu: 20m
            memory: 10Mi
        command:
        - /cluster-proportional-autoscaler
        - --namespace=kube-system
        - --configmap=coredns-autoscaler
        - --target=Deployment/coredns
        - --default-params={"linear":{"coresPerReplica":{{ContainerConfig "cores-per-replica"}},"nodesPerReplica":{{ContainerConfig "nodes-per-replica"}},"min":{{ContainerConfig "min-replicas"}}}}
        - --logtostderr=true
        - --v=2
      serviceAccount: coredns-autoscaler
      serviceAccountName: coredns-autoscaler
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

var _k8sAddonsFlannelYaml = []byte(`{{- /* This file was pulled from:
https://github.com/coreos/flannel (HEAD at time of pull was 4973e02e539378) */}}
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
apiVersion: {{if IsKubernetesVersionGe "1.16.0"}}apps/v1{{else}}extensions/v1beta1{{end}}
kind: DaemonSet
metadata:
  name: kube-flannel-ds
  namespace: kube-system
  labels:
    tier: node
    app: flannel
    addonmanager.kubernetes.io/mode: Reconcile
spec:
{{- if IsKubernetesVersionGe "1.16.0"}}
  selector:
    matchLabels:
      tier: node
      app: flannel
{{- end}}
  template:
    metadata:
      labels:
        tier: node
        app: flannel
{{- if not (IsKubernetesVersionGe "1.16.0")}}
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ''
{{- end}}
{{- if IsKubernetesVersionGe "1.17.0"}}
      annotations:
        cluster-autoscaler.kubernetes.io/daemonset-pod: "true"
{{- end}}
    spec:
      hostNetwork: true
      nodeSelector:
        beta.kubernetes.io/arch: amd64
        kubernetes.io/os: linux
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
        image: {{ContainerImage "kube-flannel"}}
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
        image: {{ContainerImage "install-cni"}}
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
{{- /* This file was pulled from:
https://github.com/coreos/flannel (HEAD at time of pull was 4973e02e539378) */}}
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

func k8sAddonsFlannelYamlBytes() ([]byte, error) {
	return _k8sAddonsFlannelYaml, nil
}

func k8sAddonsFlannelYaml() (*asset, error) {
	bytes, err := k8sAddonsFlannelYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/flannel.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsIpMasqAgentYaml = []byte(`apiVersion: {{if IsKubernetesVersionGe "1.16.0"}}apps/v1{{else}}extensions/v1beta1{{end}}
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
{{- if IsKubernetesVersionGe "1.16.0"}}
  selector:
    matchLabels:
      k8s-app: azure-ip-masq-agent
      tier: node
{{- end}}
  template:
    metadata:
      labels:
        k8s-app: azure-ip-masq-agent
        tier: node
{{- if IsKubernetesVersionGe "1.17.0"}}
      annotations:
        cluster-autoscaler.kubernetes.io/daemonset-pod: "true"
{{- end}}
    spec:
      priorityClassName: system-node-critical
      hostNetwork: true
      nodeSelector:
        kubernetes.io/os: linux
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
{{- if IsKubernetesVersionGe "1.16.0"}}
        args:
          - --enable-ipv6={{ContainerConfig "enable-ipv6"}}
{{- end}}
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
{{- if IsKubernetesVersionGe "1.16.0"}}
  {{- if ContainerConfig "secondary-non-masquerade-cidr"}}
      - {{ContainerConfig "secondary-non-masquerade-cidr"}}
  {{end -}}
{{- end}}
{{- if ContainerConfig "non-masq-cni-cidr"}}
      - {{ContainerConfig "non-masq-cni-cidr"}}
    masqLinkLocal: true
{{else}}
    masqLinkLocal: false
{{- end}}
    resyncInterval: 60s
`)

func k8sAddonsIpMasqAgentYamlBytes() ([]byte, error) {
	return _k8sAddonsIpMasqAgentYaml, nil
}

func k8sAddonsIpMasqAgentYaml() (*asset, error) {
	bytes, err := k8sAddonsIpMasqAgentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/ip-masq-agent.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsKeyvaultFlexvolumeYaml = []byte(`apiVersion: {{if IsKubernetesVersionGe "1.16.0"}}apps/v1{{else}}extensions/v1beta1{{end}}
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
{{- if IsKubernetesVersionGe "1.16.0"}}
  selector:
    matchLabels:
      app: keyvault-flexvolume
{{- end}}
  template:
    metadata:
      labels:
        app: keyvault-flexvolume
        kubernetes.io/cluster-service: "true"
        addonmanager.kubernetes.io/mode: Reconcile
{{- if IsKubernetesVersionGe "1.17.0"}}
      annotations:
        cluster-autoscaler.kubernetes.io/daemonset-pod: "true"
{{- end}}
    spec:
      priorityClassName: system-cluster-critical
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: agentpool
                operator: NotIn
                values:
                - flatcar
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
        kubernetes.io/os: linux
`)

func k8sAddonsKeyvaultFlexvolumeYamlBytes() ([]byte, error) {
	return _k8sAddonsKeyvaultFlexvolumeYaml, nil
}

func k8sAddonsKeyvaultFlexvolumeYaml() (*asset, error) {
	bytes, err := k8sAddonsKeyvaultFlexvolumeYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/keyvault-flexvolume.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsKubeDnsYaml = []byte(`{{- /* Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Should keep target in cluster/addons/dns-horizontal-autoscaler/dns-horizontal-autoscaler.yaml
in sync with this file. */}}
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
  clusterIP: {{ContainerConfig "clusterIP"}}
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
apiVersion: {{if IsKubernetesVersionGe "1.16.0"}}apps/v1{{else}}extensions/v1beta1{{end}}
kind: Deployment
metadata:
  name: kube-dns
  namespace: kube-system
  labels:
    k8s-app: kube-dns
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  {{- /* replicas: not specified here:
  1. In order to make Addon Manager do not reconcile this replicas parameter.
  2. Default is 1.
  3. Will be tuned in real time if DNS horizontal auto-scaling is turned on. */}}
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
{{- if not (IsKubernetesVersionGe "1.16.0")}}
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ''
        seccomp.security.alpha.kubernetes.io/pod: 'docker/default'
{{- end}}
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
        image: {{ContainerImage "kubedns"}}
        imagePullPolicy: IfNotPresent
        resources:
          {{- /* TODO: Set memory limits when we've profiled the container for large
          clusters, then set request = limit to keep this container in
          guaranteed class. Currently, this container falls into the
          "burstable" category so the kubelet doesn't backoff from restarting it. */}}
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
        - --domain={{ContainerConfig "domain"}}.
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
        image: {{ContainerImage "dnsmasq"}}
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
        image: {{ContainerImage "sidecar"}}
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
        - --probe=kubedns,127.0.0.1:10053,kubernetes.default.svc.{{ContainerConfig "domain"}},5,SRV
        - --probe=dnsmasq,127.0.0.1:53,kubernetes.default.svc.{{ContainerConfig "domain"}},5,SRV
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
        kubernetes.io/os: linux
`)

func k8sAddonsKubeDnsYamlBytes() ([]byte, error) {
	return _k8sAddonsKubeDnsYaml, nil
}

func k8sAddonsKubeDnsYaml() (*asset, error) {
	bytes, err := k8sAddonsKubeDnsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/kube-dns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsKubeProxyYaml = []byte(`{{if IsKubernetesVersionGe "1.16.0" -}}
apiVersion: v1
kind: ConfigMap
data:
  config.yaml: |
    apiVersion: kubeproxy.config.k8s.io/v1alpha1
    kind: KubeProxyConfiguration
    clientConnection:
      kubeconfig: /var/lib/kubelet/kubeconfig
    clusterCIDR: "{{ContainerConfig "cluster-cidr"}}"
    mode: "{{ContainerConfig "proxy-mode"}}"
  {{- if IsKubernetesVersionGe "1.18.0"}}
    {{- if ContainerConfig "bind-address"}}
    bindAddress: "{{ContainerConfig "bind-address"}}"
    {{- end}}
    {{- if ContainerConfig "healthz-bind-address"}}
    healthzBindAddress: "{{ContainerConfig "healthz-bind-address"}}"
    {{- end}}
    {{- if ContainerConfig "metrics-bind-address"}}
    metricsBindAddress: "{{ContainerConfig "metrics-bind-address"}}"
    {{- end}}
  {{- end}}
    featureGates:
      {{ContainerConfig "featureGates"}}
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
{{- end}}
apiVersion: {{if IsKubernetesVersionGe "1.16.0"}}apps/v1{{else}}extensions/v1beta1{{end}}
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
{{- if IsKubernetesVersionGe "1.16.0"}}
      component: kube-proxy
      tier: node
{{- end}}
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
{{- if IsKubernetesVersionGe "1.17.0"}}
        cluster-autoscaler.kubernetes.io/daemonset-pod: "true"
{{- end}}
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
{{- if not (IsKubernetesVersionGe "1.16.0")}}
        - /hyperkube
{{- end}}
        - kube-proxy
{{- if not (IsKubernetesVersionGe "1.16.0")}}
        - --kubeconfig=/var/lib/kubelet/kubeconfig
        - --cluster-cidr={{ContainerConfig "cluster-cidr"}}
        - --feature-gates=ExperimentalCriticalPodAnnotation=true
        - --proxy-mode={{ContainerConfig "proxy-mode"}}
{{else}}
        - --config=/var/lib/kube-proxy/config.yaml
{{- end}}
        image: {{ContainerImage "kube-proxy"}}
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
{{- if IsKubernetesVersionGe "1.16.0"}}
        - mountPath: /var/lib/kube-proxy/config.yaml
          subPath: config.yaml
          name: kube-proxy-config-volume
          readOnly: true
{{- end}}
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
{{- if IsKubernetesVersionGe "1.16.0"}}
      - configMap:
          name: kube-proxy-config
        name: kube-proxy-config-volume
{{- end}}
      nodeSelector:
        kubernetes.io/os: linux
`)

func k8sAddonsKubeProxyYamlBytes() ([]byte, error) {
	return _k8sAddonsKubeProxyYaml, nil
}

func k8sAddonsKubeProxyYaml() (*asset, error) {
	bytes, err := k8sAddonsKubeProxyYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/kube-proxy.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsKubeReschedulerYaml = []byte(`apiVersion: {{if IsKubernetesVersionGe "1.16.0"}}apps/v1{{else}}extensions/v1beta1{{end}}
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
{{- if not (IsKubernetesVersionGe "1.16.0")}}
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ''
{{- end}}
    spec:
{{- if IsKubernetesVersionGe "1.16.0"}}
      priorityClassName: system-node-critical
{{- end}}
      nodeSelector:
        kubernetes.io/os: linux
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

func k8sAddonsKubeReschedulerYamlBytes() ([]byte, error) {
	return _k8sAddonsKubeReschedulerYaml, nil
}

func k8sAddonsKubeReschedulerYaml() (*asset, error) {
	bytes, err := k8sAddonsKubeReschedulerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/kube-rescheduler.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsKubernetesDashboardYaml = []byte(`
apiVersion: v1
kind: Namespace
metadata:
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
  name: kubernetes-dashboard
---
apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    k8s-app: kubernetes-dashboard
    addonmanager.kubernetes.io/mode: Reconcile
  name: kubernetes-dashboard
  namespace: kubernetes-dashboard
---
kind: Service
apiVersion: v1
metadata:
  labels:
    k8s-app: kubernetes-dashboard
    addonmanager.kubernetes.io/mode: Reconcile
  name: kubernetes-dashboard
  namespace: kubernetes-dashboard
spec:
  ports:
    - port: 443
      targetPort: 8443
  selector:
    k8s-app: kubernetes-dashboard
---
apiVersion: v1
kind: Secret
metadata:
  labels:
    k8s-app: kubernetes-dashboard
    addonmanager.kubernetes.io/mode: EnsureExists
  name: kubernetes-dashboard-certs
  namespace: kubernetes-dashboard
type: Opaque
---
apiVersion: v1
kind: Secret
metadata:
  labels:
    k8s-app: kubernetes-dashboard
    addonmanager.kubernetes.io/mode: EnsureExists
  name: kubernetes-dashboard-csrf
  namespace: kubernetes-dashboard
type: Opaque
data:
  csrf: ""
---
apiVersion: v1
kind: Secret
metadata:
  labels:
    k8s-app: kubernetes-dashboard
    addonmanager.kubernetes.io/mode: EnsureExists
  name: kubernetes-dashboard-key-holder
  namespace: kubernetes-dashboard
type: Opaque
---
kind: ConfigMap
apiVersion: v1
metadata:
  labels:
    k8s-app: kubernetes-dashboard
    addonmanager.kubernetes.io/mode: EnsureExists
  name: kubernetes-dashboard-settings
  namespace: kubernetes-dashboard
---
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  labels:
    k8s-app: kubernetes-dashboard
    addonmanager.kubernetes.io/mode: Reconcile
  name: kubernetes-dashboard
  namespace: kubernetes-dashboard
rules:
  {{- /* Allow Dashboard to get, update and delete Dashboard exclusive secrets. */}}
  - apiGroups: [""]
    resources: ["secrets"]
    resourceNames: ["kubernetes-dashboard-key-holder", "kubernetes-dashboard-certs", "kubernetes-dashboard-csrf"]
    verbs: ["get", "update", "delete"]
    {{- /* Allow Dashboard to get and update 'kubernetes-dashboard-settings' config map. */}}
  - apiGroups: [""]
    resources: ["configmaps"]
    resourceNames: ["kubernetes-dashboard-settings"]
    verbs: ["get", "update"]
    {{- /* Allow Dashboard to get metrics. */}}
  - apiGroups: [""]
    resources: ["services"]
    resourceNames: ["heapster", "dashboard-metrics-scraper"]
    verbs: ["proxy"]
  - apiGroups: [""]
    resources: ["services/proxy"]
    resourceNames: ["heapster", "http:heapster:", "https:heapster:", "dashboard-metrics-scraper", "http:dashboard-metrics-scraper"]
    verbs: ["get"]
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  labels:
    k8s-app: kubernetes-dashboard
    addonmanager.kubernetes.io/mode: Reconcile
  name: kubernetes-dashboard
rules:
  {{- /* Allow Metrics Scraper to get metrics from the Metrics server */}}
  - apiGroups: ["metrics.k8s.io"]
    resources: ["pods", "nodes"]
    verbs: ["get", "list", "watch"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  labels:
    k8s-app: kubernetes-dashboard
    addonmanager.kubernetes.io/mode: Reconcile
  name: kubernetes-dashboard
  namespace: kubernetes-dashboard
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: kubernetes-dashboard
subjects:
  - kind: ServiceAccount
    name: kubernetes-dashboard
    namespace: kubernetes-dashboard
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
  name: kubernetes-dashboard
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: kubernetes-dashboard
subjects:
  - kind: ServiceAccount
    name: kubernetes-dashboard
    namespace: kubernetes-dashboard
---
kind: Deployment
apiVersion: apps/v1
metadata:
  labels:
    k8s-app: kubernetes-dashboard
    addonmanager.kubernetes.io/mode: Reconcile
  name: kubernetes-dashboard
  namespace: kubernetes-dashboard
spec:
  replicas: 1
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      k8s-app: kubernetes-dashboard
  template:
    metadata:
      labels:
        k8s-app: kubernetes-dashboard
    spec:
      containers:
        - name: kubernetes-dashboard
          image: {{ContainerImage "kubernetes-dashboard"}}
          imagePullPolicy: IfNotPresent
          ports:
            - containerPort: 8443
              protocol: TCP
          args:
            - --auto-generate-certificates
            - --namespace=kubernetes-dashboard
            {{- /* Uncomment the following line to manually specify Kubernetes API server Host */}}
            {{- /* If not specified, Dashboard will attempt to auto discover the API server and connect */}}
            {{- /* to it. Uncomment only if the default does not work. */}}
            {{- /* - --apiserver-host=http://my-address:port */}}
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
              {{- /* Create on-disk volume to store exec logs */}}
            - mountPath: /tmp
              name: tmp-volume
          livenessProbe:
            httpGet:
              scheme: HTTPS
              path: /
              port: 8443
            initialDelaySeconds: 30
            timeoutSeconds: 30
          securityContext:
            allowPrivilegeEscalation: false
            readOnlyRootFilesystem: true
            runAsUser: 1001
            runAsGroup: 2001
      volumes:
        - name: kubernetes-dashboard-certs
          secret:
            secretName: kubernetes-dashboard-certs
        - name: tmp-volume
          emptyDir: {}
      serviceAccountName: kubernetes-dashboard
      nodeSelector:
        kubernetes.io/os: linux
      {{/* Comment the following tolerations if Dashboard must not be deployed on master */}}
      tolerations:
        - key: node-role.kubernetes.io/master
          effect: NoSchedule
---
kind: Service
apiVersion: v1
metadata:
  labels:
    k8s-app: dashboard-metrics-scraper
    addonmanager.kubernetes.io/mode: Reconcile
  name: dashboard-metrics-scraper
  namespace: kubernetes-dashboard
spec:
  ports:
    - port: 8000
      targetPort: 8000
  selector:
    k8s-app: dashboard-metrics-scraper
---
kind: Deployment
apiVersion: apps/v1
metadata:
  labels:
    k8s-app: dashboard-metrics-scraper
    addonmanager.kubernetes.io/mode: Reconcile
  name: dashboard-metrics-scraper
  namespace: kubernetes-dashboard
spec:
  replicas: 1
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      k8s-app: dashboard-metrics-scraper
  template:
    metadata:
      labels:
        k8s-app: dashboard-metrics-scraper
      annotations:
        seccomp.security.alpha.kubernetes.io/pod: 'runtime/default'
    spec:
      containers:
        - name: dashboard-metrics-scraper
          image: {{ContainerImage "kubernetes-dashboard-metrics-scraper"}}
          ports:
            - containerPort: 8000
              protocol: TCP
          livenessProbe:
            httpGet:
              scheme: HTTP
              path: /
              port: 8000
            initialDelaySeconds: 30
            timeoutSeconds: 30
          resources:
            requests:
              cpu: {{ContainerCPUReqs "kubernetes-dashboard-metrics-scraper"}}
              memory: {{ContainerMemReqs "kubernetes-dashboard-metrics-scraper"}}
            limits:
              cpu: {{ContainerCPULimits "kubernetes-dashboard-metrics-scraper"}}
              memory: {{ContainerMemLimits "kubernetes-dashboard-metrics-scraper"}}
          volumeMounts:
          - mountPath: /tmp
            name: tmp-volume
          securityContext:
            allowPrivilegeEscalation: false
            readOnlyRootFilesystem: true
            runAsUser: 1001
            runAsGroup: 2001
      serviceAccountName: kubernetes-dashboard
      nodeSelector:
        kubernetes.io/os: linux
        {{- /* Comment the following tolerations if Dashboard must not be deployed on master */}}
      tolerations:
        - key: node-role.kubernetes.io/master
          effect: NoSchedule
      volumes:
        - name: tmp-volume
          emptyDir: {}
`)

func k8sAddonsKubernetesDashboardYamlBytes() ([]byte, error) {
	return _k8sAddonsKubernetesDashboardYaml, nil
}

func k8sAddonsKubernetesDashboardYaml() (*asset, error) {
	bytes, err := k8sAddonsKubernetesDashboardYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/kubernetes-dashboard.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsMetricsServerYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: metrics-server
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: {{GetMode}}
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: system:metrics-server
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: {{GetMode}}
rules:
- apiGroups:
  - ""
  resources:
  - pods
  - nodes
{{- if IsKubernetesVersionGe "1.16.0"}}
  - nodes/stats
{{- end}}
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
    addonmanager.kubernetes.io/mode: {{GetMode}}
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
    addonmanager.kubernetes.io/mode: {{GetMode}}
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
    addonmanager.kubernetes.io/mode: {{GetMode}}
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
    addonmanager.kubernetes.io/mode: {{GetMode}}
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
{{- if IsKubernetesVersionGe "1.16.0"}}
apiVersion: apps/v1
{{else}}
apiVersion: extensions/v1beta1
{{- end}}
kind: Deployment
metadata:
  name: metrics-server
  namespace: kube-system
  labels:
    k8s-app: metrics-server
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: {{GetMode}}
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
      priorityClassName: system-cluster-critical
      containers:
      - name: metrics-server
        image: {{ContainerImage "metrics-server"}}
        imagePullPolicy: IfNotPresent
        command:
        - /metrics-server
{{- if IsKubernetesVersionGe "1.16.0"}}
        - --kubelet-insecure-tls
{{else}}
        - --source=kubernetes.summary_api:''
{{- end}}
      nodeSelector:
        kubernetes.io/os: linux
---
apiVersion: apiregistration.k8s.io/v1beta1
kind: APIService
metadata:
  name: v1beta1.metrics.k8s.io
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: {{GetMode}}
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

func k8sAddonsMetricsServerYamlBytes() ([]byte, error) {
	return _k8sAddonsMetricsServerYaml, nil
}

func k8sAddonsMetricsServerYaml() (*asset, error) {
	bytes, err := k8sAddonsMetricsServerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/metrics-server.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsNodeProblemDetectorYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: node-problem-detector
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: node-problem-detector-binding
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:node-problem-detector
subjects:
- kind: ServiceAccount
  name: node-problem-detector
  namespace: kube-system
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: node-problem-detector-{{ContainerConfig "versionLabel"}}
  namespace: kube-system
  labels:
    k8s-app: node-problem-detector
    version: {{ContainerConfig "versionLabel"}}
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  selector:
    matchLabels:
      k8s-app: node-problem-detector
      version: {{ContainerConfig "versionLabel"}}
  template:
    metadata:
      labels:
        k8s-app: node-problem-detector
        version: {{ContainerConfig "versionLabel"}}
        kubernetes.io/cluster-service: "true"
      annotations:
        cluster-autoscaler.kubernetes.io/daemonset-pod: "true"
    spec:
      containers:
      - name: node-problem-detector
        image: {{ContainerImage "node-problem-detector"}}
        command:
        - "/bin/sh"
        - "-c"
        - "exec /node-problem-detector --logtostderr --prometheus-address=0.0.0.0 --config.system-log-monitor={{ContainerConfig "systemLogMonitor"}} --config.custom-plugin-monitor={{ContainerConfig "customPluginMonitor"}} --config.system-stats-monitor={{ContainerConfig "systemStatsMonitor"}} >>/var/log/node-problem-detector.log 2>&1"
        securityContext:
          privileged: true
        resources:
          limits:
            cpu: {{ContainerCPULimits "node-problem-detector"}}
            memory: {{ContainerMemLimits "node-problem-detector"}}
          requests:
            cpu: {{ContainerCPUReqs "node-problem-detector"}}
            memory: {{ContainerMemReqs "node-problem-detector"}}
        env:
        - name: NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        volumeMounts:
        - name: log
          mountPath: /var/log
        - name: localtime
          mountPath: /etc/localtime
          readOnly: true
        - name: kmsg
          mountPath: /dev/kmsg
          readOnly: true
      volumes:
      - name: log
        hostPath:
          path: /var/log/
      - name: localtime
        hostPath:
          path: /etc/localtime
          type: "FileOrCreate"
      - name: kmsg
        hostPath:
          path: /dev/kmsg
          type: "CharDevice"
      nodeSelector:
        kubernetes.io/os: linux
      serviceAccountName: node-problem-detector
      tolerations:
      - operator: "Exists"
        effect: "NoExecute"
      - key: "CriticalAddonsOnly"
        operator: "Exists"
---
apiVersion: v1
kind: Service
metadata:
  labels:
    k8s-app: node-problem-detector
    addonmanager.kubernetes.io/mode: Reconcile
  name: node-problem-detector
  namespace: kube-system
spec:
  clusterIP: None
  ports:
  - name: exporter
    port: 20257
    protocol: TCP
  selector:
    k8s-app: node-problem-detector
  type: ClusterIP
`)

func k8sAddonsNodeProblemDetectorYamlBytes() ([]byte, error) {
	return _k8sAddonsNodeProblemDetectorYaml, nil
}

func k8sAddonsNodeProblemDetectorYaml() (*asset, error) {
	bytes, err := k8sAddonsNodeProblemDetectorYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/node-problem-detector.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsNvidiaDevicePluginYaml = []byte(`apiVersion: apps/v1
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
{{- if not (IsKubernetesVersionGe "1.16.0")}}
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ""
{{- end}}
{{- if IsKubernetesVersionGe "1.17.0"}}
      annotations:
        cluster-autoscaler.kubernetes.io/daemonset-pod: "true"
{{- end}}
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
        kubernetes.io/os: linux
        accelerator: nvidia
`)

func k8sAddonsNvidiaDevicePluginYamlBytes() ([]byte, error) {
	return _k8sAddonsNvidiaDevicePluginYaml, nil
}

func k8sAddonsNvidiaDevicePluginYaml() (*asset, error) {
	bytes, err := k8sAddonsNvidiaDevicePluginYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/nvidia-device-plugin.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsPodSecurityPolicyYaml = []byte(`apiVersion: {{if IsKubernetesVersionGe "1.16.0"}}policy{{else}}extensions{{end}}/v1beta1
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
apiVersion: {{if IsKubernetesVersionGe "1.16.0"}}policy{{else}}extensions{{end}}/v1beta1
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
      {{- /* Forbid adding the root group. */}}
      - min: 1
        max: 65535
  fsGroup:
    rule: MustRunAs
    ranges:
      {{- /* Forbid adding the root group. */}}
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
kind: {{if IsKubernetesVersionGe "1.15.0"}}ClusterRoleBinding{{else}}RoleBinding{{end}}
metadata:
  name: default:privileged
{{- if not (IsKubernetesVersionGe "1.15.0")}}
  namespace: kube-system
{{end}}
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: psp:privileged
subjects:
- kind: Group
  name: {{if IsKubernetesVersionGe "1.15.0"}}system:authenticated{{else}}system:masters{{end}}
  apiGroup: rbac.authorization.k8s.io
{{- if not (IsKubernetesVersionGe "1.15.0")}}
- kind: Group
  name: system:serviceaccounts:kube-system
  apiGroup: rbac.authorization.k8s.io
{{end}}
- kind: Group
  name: system:nodes
  apiGroup: rbac.authorization.k8s.io
#EOF
`)

func k8sAddonsPodSecurityPolicyYamlBytes() ([]byte, error) {
	return _k8sAddonsPodSecurityPolicyYaml, nil
}

func k8sAddonsPodSecurityPolicyYaml() (*asset, error) {
	bytes, err := k8sAddonsPodSecurityPolicyYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/pod-security-policy.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsScheduledMaintenanceDeploymentYaml = []byte(`apiVersion: v1
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
        image: {{ContainerImage "kube-rbac-proxy"}}
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
        image: {{ContainerImage "manager"}}
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
        image: {{ContainerImage "manager"}}
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

func k8sAddonsScheduledMaintenanceDeploymentYamlBytes() ([]byte, error) {
	return _k8sAddonsScheduledMaintenanceDeploymentYaml, nil
}

func k8sAddonsScheduledMaintenanceDeploymentYaml() (*asset, error) {
	bytes, err := k8sAddonsScheduledMaintenanceDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/scheduled-maintenance-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsSecretsStoreCsiDriverYaml = []byte(`apiVersion: storage.k8s.io/v1beta1
kind: CSIDriver
metadata:
  name: secrets-store.csi.k8s.io
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  podInfoOnMount: true
  attachRequired: false
  volumeLifecycleModes:
  - Ephemeral
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: secrets-store-csi-driver
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: secretproviderclasses-rolebinding
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: secretproviderclasses-role
subjects:
- kind: ServiceAccount
  name: secrets-store-csi-driver
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: secretproviderclasses-role
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups:
  - secrets-store.csi.x-k8s.io
  resources:
  - secretproviderclasses
  verbs:
  - get
  - list
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: secretproviderclasses.secrets-store.csi.x-k8s.io
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  group: secrets-store.csi.x-k8s.io
  names:
    kind: SecretProviderClass
    listKind: SecretProviderClassList
    plural: secretproviderclasses
    singular: secretproviderclass
  scope: ""
  validation:
    openAPIV3Schema:
      description: SecretProviderClass is the Schema for the secretproviderclasses
        API
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          description: SecretProviderClassSpec defines the desired state of SecretProviderClass
          properties:
            parameters:
              additionalProperties:
                type: string
              description: Configuration for specific provider
              type: object
            provider:
              description: Configuration for provider name
              type: string
          type: object
        status:
          description: SecretProviderClassStatus defines the observed state of SecretProviderClass
          type: object
      type: object
  version: v1alpha1
  versions:
  - name: v1alpha1
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
---
kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: csi-secrets-store
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  selector:
    matchLabels:
      app: csi-secrets-store
  template:
    metadata:
      labels:
        app: csi-secrets-store
    spec:
      serviceAccountName: secrets-store-csi-driver
      hostNetwork: true
      containers:
        - name: node-driver-registrar
          image: {{ContainerImage "csi-node-driver-registrar"}}
          args:
            - --v=5
            - --csi-address=/csi/csi.sock
            - --kubelet-registration-path=/var/lib/kubelet/plugins/csi-secrets-store/csi.sock
          lifecycle:
            preStop:
              exec:
                command:
                  [
                    "/bin/sh",
                    "-c",
                    "rm -rf /registration/secrets-store.csi.k8s.io-reg.sock",
                  ]
          env:
            - name: KUBE_NODE_NAME
              valueFrom:
                fieldRef:
                  apiVersion: v1
                  fieldPath: spec.nodeName
          imagePullPolicy: IfNotPresent
          volumeMounts:
            - name: plugin-dir
              mountPath: /csi
            - name: registration-dir
              mountPath: /registration
          resources:
            limits:
              cpu: {{ContainerCPULimits "csi-node-driver-registrar"}}
              memory: {{ContainerMemLimits "csi-node-driver-registrar"}}
            requests:
              cpu: {{ContainerCPUReqs "csi-node-driver-registrar"}}
              memory: {{ContainerMemReqs "csi-node-driver-registrar"}}
        - name: secrets-store
          image: {{ContainerImage "secrets-store"}}
          args:
            - "--debug=true"
            - "--endpoint=$(CSI_ENDPOINT)"
            - "--nodeid=$(KUBE_NODE_NAME)"
            - "--provider-volume=/etc/kubernetes/secrets-store-csi-providers"
          env:
            - name: CSI_ENDPOINT
              value: unix:///csi/csi.sock
            - name: KUBE_NODE_NAME
              valueFrom:
                fieldRef:
                  apiVersion: v1
                  fieldPath: spec.nodeName
          imagePullPolicy: IfNotPresent
          securityContext:
            privileged: true
          ports:
            - containerPort: 9808
              name: healthz
              protocol: TCP
          livenessProbe:
              failureThreshold: 5
              httpGet:
                path: /healthz
                port: healthz
              initialDelaySeconds: 30
              timeoutSeconds: 10
              periodSeconds: 15
          volumeMounts:
            - name: plugin-dir
              mountPath: /csi
            - name: mountpoint-dir
              mountPath: /var/lib/kubelet/pods
              mountPropagation: Bidirectional
            - name: providers-dir
              mountPath: /etc/kubernetes/secrets-store-csi-providers
          resources:
            limits:
              cpu: {{ContainerCPULimits "secrets-store"}}
              memory: {{ContainerMemLimits "secrets-store"}}
            requests:
              cpu: {{ContainerCPUReqs "secrets-store"}}
              memory: {{ContainerMemReqs "secrets-store"}}
        - name: liveness-probe
          image: {{ContainerImage "livenessprobe"}}
          imagePullPolicy: IfNotPresent
          args:
          - --csi-address=/csi/csi.sock
          - --probe-timeout=3s
          - --health-port=9808
          volumeMounts:
            - name: plugin-dir
              mountPath: /csi
          resources:
            limits:
              cpu: {{ContainerCPULimits "livenessprobe"}}
              memory: {{ContainerMemLimits "livenessprobe"}}
            requests:
              cpu: {{ContainerCPUReqs "livenessprobe"}}
              memory: {{ContainerMemReqs "livenessprobe"}}
      volumes:
        - name: mountpoint-dir
          hostPath:
            path: /var/lib/kubelet/pods
            type: DirectoryOrCreate
        - name: registration-dir
          hostPath:
            path: /var/lib/kubelet/plugins_registry/
            type: Directory
        - name: plugin-dir
          hostPath:
            path: /var/lib/kubelet/plugins/csi-secrets-store/
            type: DirectoryOrCreate
        - name: providers-dir
          hostPath:
            path: /etc/kubernetes/secrets-store-csi-providers
            type: DirectoryOrCreate
      nodeSelector:
        kubernetes.io/os: linux
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: csi-secrets-store-provider-azure
  namespace: kube-system
  labels:
    app: csi-secrets-store-provider-azure
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  updateStrategy:
    type: RollingUpdate
  selector:
    matchLabels:
      app: csi-secrets-store-provider-azure
  template:
    metadata:
      labels:
        app: csi-secrets-store-provider-azure
    spec:
      tolerations:
      containers:
        - name: provider-azure-installer
          image: {{ContainerImage "provider-azure-installer"}}
          imagePullPolicy: IfNotPresent
          env:
            - name: TARGET_DIR
              value: "/etc/kubernetes/secrets-store-csi-providers"
          volumeMounts:
            - mountPath: "/etc/kubernetes/secrets-store-csi-providers"
              name: providervol
          resources:
            limits:
              cpu: {{ContainerCPULimits "provider-azure-installer"}}
              memory: {{ContainerMemLimits "provider-azure-installer"}}
            requests:
              cpu: {{ContainerCPUReqs "provider-azure-installer"}}
              memory: {{ContainerMemReqs "provider-azure-installer"}}
      volumes:
        - name: providervol
          hostPath:
            path: "/etc/kubernetes/secrets-store-csi-providers"
      nodeSelector:
        kubernetes.io/os: linux
`)

func k8sAddonsSecretsStoreCsiDriverYamlBytes() ([]byte, error) {
	return _k8sAddonsSecretsStoreCsiDriverYaml, nil
}

func k8sAddonsSecretsStoreCsiDriverYaml() (*asset, error) {
	bytes, err := k8sAddonsSecretsStoreCsiDriverYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/secrets-store-csi-driver.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsSmbFlexvolumeYaml = []byte(`apiVersion: {{if IsKubernetesVersionGe "1.16.0"}}apps/v1{{else}}extensions/v1beta1{{end}}
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
{{- if IsKubernetesVersionGe "1.17.0"}}
      annotations:
        cluster-autoscaler.kubernetes.io/daemonset-pod: "true"
{{- end}}
    spec:
      containers:
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: agentpool
                operator: NotIn
                values:
                - flatcar
      - name: smb-flexvol-installer
        image: {{ContainerImage "smb-flexvolume"}}
        imagePullPolicy: IfNotPresent
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
        kubernetes.io/os: linux
`)

func k8sAddonsSmbFlexvolumeYamlBytes() ([]byte, error) {
	return _k8sAddonsSmbFlexvolumeYaml, nil
}

func k8sAddonsSmbFlexvolumeYaml() (*asset, error) {
	bytes, err := k8sAddonsSmbFlexvolumeYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/smb-flexvolume.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAddonsTillerYaml = []byte(`apiVersion: v1
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
apiVersion: {{if IsKubernetesVersionGe "1.16.0"}}apps/v1{{else}}extensions/v1beta1{{end}}
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
{{- if IsKubernetesVersionGe "1.16.0"}}
  selector:
    matchLabels:
      app: helm
      name: tiller
{{- end}}
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
        kubernetes.io/os: linux
`)

func k8sAddonsTillerYamlBytes() ([]byte, error) {
	return _k8sAddonsTillerYaml, nil
}

func k8sAddonsTillerYaml() (*asset, error) {
	bytes, err := k8sAddonsTillerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/addons/tiller.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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

    echo 'root:'$HASH | /usr/sbin/chpasswd -e || exit 112
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
    mkdir -p ${DIR} || exit 112
    touch ${FILEPATH} || exit 112
    chmod 640 ${FILEPATH} || exit 112
  done
  find /var/log -type f -perm '/o+r' -exec chmod 'g-wx,o-rwx' {} \;
  chmod 600 /etc/passwd- || exit 112
  chmod 600 /etc/shadow- || exit 112
  chmod 600 /etc/group- || exit 112
  chmod 644 /etc/default/grub || exit 112
  for filepath in /etc/crontab /etc/cron.hourly /etc/cron.daily /etc/cron.weekly /etc/cron.monthly /etc/cron.d; do
    chmod 0600 $filepath || exit 112
  done
}

setPWExpiration() {
  sed -i "s|PASS_MAX_DAYS||g" /etc/login.defs || exit 115
  grep 'PASS_MAX_DAYS' /etc/login.defs && exit 115
  sed -i "s|PASS_MIN_DAYS||g" /etc/login.defs || exit 115
  grep 'PASS_MIN_DAYS' /etc/login.defs && exit 115
  sed -i "s|INACTIVE=||g" /etc/default/useradd || exit 115
  grep 'INACTIVE=' /etc/default/useradd && exit 115
  echo 'PASS_MAX_DAYS 90' >>/etc/login.defs || exit 115
  grep 'PASS_MAX_DAYS 90' /etc/login.defs || exit 115
  echo 'PASS_MIN_DAYS 7' >>/etc/login.defs || exit 115
  grep 'PASS_MIN_DAYS 7' /etc/login.defs || exit 115
  echo 'INACTIVE=30' >>/etc/default/useradd || exit 115
  grep 'INACTIVE=30' /etc/default/useradd || exit 115
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
PRIVATE_IP=$(hostname -I | cut -d' ' -f1)
ETCD_PEER_URL="https://${PRIVATE_IP}:2380"
ETCD_CLIENT_URL="https://${PRIVATE_IP}:2379"
KUBECTL="/usr/local/bin/kubectl --kubeconfig=/home/$ADMINUSER/.kube/config"
ADDONS_DIR=/etc/kubernetes/addons
POD_SECURITY_POLICY_SPEC=$ADDONS_DIR/pod-security-policy.yaml
ADDON_MANAGER_SPEC=/etc/kubernetes/manifests/kube-addon-manager.yaml

systemctlEnableAndStart() {
  systemctl_restart 100 5 30 $1
  RESTART_STATUS=$?
  systemctl status $1 --no-pager -l >/var/log/azure/$1-status.log
  if [ $RESTART_STATUS -ne 0 ]; then
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
  APISERVER_PRIVATE_KEY_PATH="/etc/kubernetes/certs/apiserver.key"
  touch "${APISERVER_PRIVATE_KEY_PATH}"
  CA_PRIVATE_KEY_PATH="/etc/kubernetes/certs/ca.key"
  touch "${CA_PRIVATE_KEY_PATH}"
  ETCD_SERVER_PRIVATE_KEY_PATH="/etc/kubernetes/certs/etcdserver.key"
  touch "${ETCD_SERVER_PRIVATE_KEY_PATH}"
  if [[ -z ${COSMOS_URI} ]]; then
    chown etcd:etcd "${ETCD_SERVER_PRIVATE_KEY_PATH}"
  fi
  ETCD_CLIENT_PRIVATE_KEY_PATH="/etc/kubernetes/certs/etcdclient.key"
  touch "${ETCD_CLIENT_PRIVATE_KEY_PATH}"
  ETCD_PEER_PRIVATE_KEY_PATH="/etc/kubernetes/certs/etcdpeer${NODE_INDEX}.key"
  touch "${ETCD_PEER_PRIVATE_KEY_PATH}"
  if [[ -z ${COSMOS_URI} ]]; then
    chown etcd:etcd "${ETCD_PEER_PRIVATE_KEY_PATH}"
  fi
  chmod 0600 "${APISERVER_PRIVATE_KEY_PATH}" "${CA_PRIVATE_KEY_PATH}" "${ETCD_SERVER_PRIVATE_KEY_PATH}" "${ETCD_CLIENT_PRIVATE_KEY_PATH}" "${ETCD_PEER_PRIVATE_KEY_PATH}"
  chown root:root "${APISERVER_PRIVATE_KEY_PATH}" "${CA_PRIVATE_KEY_PATH}" "${ETCD_CLIENT_PRIVATE_KEY_PATH}"
  ETCD_SERVER_CERTIFICATE_PATH="/etc/kubernetes/certs/etcdserver.crt"
  touch "${ETCD_SERVER_CERTIFICATE_PATH}"
  ETCD_CLIENT_CERTIFICATE_PATH="/etc/kubernetes/certs/etcdclient.crt"
  touch "${ETCD_CLIENT_CERTIFICATE_PATH}"
  ETCD_PEER_CERTIFICATE_PATH="/etc/kubernetes/certs/etcdpeer${NODE_INDEX}.crt"
  touch "${ETCD_PEER_CERTIFICATE_PATH}"
  chmod 0644 "${ETCD_SERVER_CERTIFICATE_PATH}" "${ETCD_CLIENT_CERTIFICATE_PATH}" "${ETCD_PEER_CERTIFICATE_PATH}"
  chown root:root "${ETCD_SERVER_CERTIFICATE_PATH}" "${ETCD_CLIENT_CERTIFICATE_PATH}" "${ETCD_PEER_CERTIFICATE_PATH}"

  set +x
  echo "${APISERVER_PRIVATE_KEY}" | base64 --decode >"${APISERVER_PRIVATE_KEY_PATH}"
  echo "${CA_PRIVATE_KEY}" | base64 --decode >"${CA_PRIVATE_KEY_PATH}"
  echo "${ETCD_SERVER_PRIVATE_KEY}" | base64 --decode >"${ETCD_SERVER_PRIVATE_KEY_PATH}"
  echo "${ETCD_CLIENT_PRIVATE_KEY}" | base64 --decode >"${ETCD_CLIENT_PRIVATE_KEY_PATH}"
  echo "${ETCD_PEER_KEY}" | base64 --decode >"${ETCD_PEER_PRIVATE_KEY_PATH}"
  echo "${ETCD_SERVER_CERTIFICATE}" | base64 --decode >"${ETCD_SERVER_CERTIFICATE_PATH}"
  echo "${ETCD_CLIENT_CERTIFICATE}" | base64 --decode >"${ETCD_CLIENT_CERTIFICATE_PATH}"
  echo "${ETCD_PEER_CERT}" | base64 --decode >"${ETCD_PEER_CERTIFICATE_PATH}"
}
configureEtcd() {
  set -x

  ETCD_SETUP_FILE=/opt/azure/containers/setup-etcd.sh
  wait_for_file 1200 1 $ETCD_SETUP_FILE || exit {{GetCSEErrorCode "ERR_ETCD_CONFIG_FAIL"}}
  $ETCD_SETUP_FILE >/opt/azure/containers/setup-etcd.log 2>&1
  RET=$?
  if [ $RET -ne 0 ]; then
    exit $RET
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
  retrycmd 120 5 25 sudo -E etcdctl member update $MEMBER ${ETCD_PEER_URL} || exit {{GetCSEErrorCode "ERR_ETCD_CONFIG_FAIL"}}
}
ensureNTP() {
  systemctlEnableAndStart ntp || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
}
configPrivateClusterHosts() {
  systemctlEnableAndStart reconcile-private-hosts || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
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
  local CRON_SERVICE=/lib/systemd/system/cron.service
  if [[ -f ${CRON_SERVICE} ]]; then
    if ! grep -q 'Restart=' ${CRON_SERVICE}; then
      sed -i 's/\[Service\]/[Service]\nRestart=always/' ${CRON_SERVICE}
      systemctlEnableAndStart cron
    fi
  fi
}
generateAggregatedAPICerts() {
  AGGREGATED_API_CERTS_SETUP_FILE=/etc/kubernetes/generate-proxy-certs.sh
  wait_for_file 1200 1 $AGGREGATED_API_CERTS_SETUP_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
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
  APISERVER_PUBLIC_KEY_PATH="/etc/kubernetes/certs/apiserver.crt"
  touch "${APISERVER_PUBLIC_KEY_PATH}"
  chmod 0600 "${KUBELET_PRIVATE_KEY_PATH}"
  chmod 0644 "${APISERVER_PUBLIC_KEY_PATH}"
  chown root:root "${KUBELET_PRIVATE_KEY_PATH}" "${APISERVER_PUBLIC_KEY_PATH}"

  set +x
  echo "${KUBELET_PRIVATE_KEY}" | base64 --decode >"${KUBELET_PRIVATE_KEY_PATH}"
  echo "${APISERVER_PUBLIC_KEY}" | base64 --decode >"${APISERVER_PUBLIC_KEY_PATH}"
  configureKubeletServerCert
  AZURE_JSON_PATH="/etc/kubernetes/azure.json"
  if [[ -n ${MASTER_NODE} ]]; then
    if [[ ${ENABLE_AGGREGATED_APIS} == True ]]; then
      generateAggregatedAPICerts
    fi
  else
    {{- /* If we are a node vm then we only proceed w/ local azure.json configuration if cloud-init has pre-paved that file */}}
    wait_for_file 1 1 $AZURE_JSON_PATH || return
  fi

  {{/* Perform the required JSON escaping */}}
  SERVICE_PRINCIPAL_CLIENT_SECRET=${SERVICE_PRINCIPAL_CLIENT_SECRET//\\/\\\\}
  SERVICE_PRINCIPAL_CLIENT_SECRET=${SERVICE_PRINCIPAL_CLIENT_SECRET//\"/\\\"}
  cat <<EOF >"${AZURE_JSON_PATH}"
{
    "cloud":"{{GetTargetEnvironment}}",
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
    "providerKeyVersion": ""
}
EOF
  set -x
  if [[ ${CLOUDPROVIDER_BACKOFF_MODE} == "v2" ]]; then
    sed -i "/cloudProviderBackoffExponent/d" $AZURE_JSON_PATH
    sed -i "/cloudProviderBackoffJitter/d" $AZURE_JSON_PATH
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
  if [[ "${NETWORK_PLUGIN}" == "azure" ]]; then
    mv $CNI_BIN_DIR/10-azure.conflist $CNI_CONFIG_DIR/
    chmod 600 $CNI_CONFIG_DIR/10-azure.conflist
    if [[ "${IS_IPV6_DUALSTACK_FEATURE_ENABLED}" == "true" ]]; then
      echo $(cat "$CNI_CONFIG_DIR/10-azure.conflist" | jq '.plugins[0].ipv6Mode="ipv6nat"') > "$CNI_CONFIG_DIR/10-azure.conflist"
    fi
    if [[ {{GetKubeProxyMode}} == "ipvs" ]]; then
      serviceCidrs={{GetServiceCidr}}
      echo $(cat "$CNI_CONFIG_DIR/10-azure.conflist" | jq  --arg serviceCidrs $serviceCidrs '.plugins[0]+={serviceCidrs: $serviceCidrs}') > /etc/cni/net.d/10-azure.conflist
    fi
    if [[ "${NETWORK_POLICY}" == "calico" ]]; then
      sed -i 's#"mode":"bridge"#"mode":"transparent"#g' $CNI_CONFIG_DIR/10-azure.conflist
    elif [[ "${NETWORK_POLICY}" == "antrea" ]]; then
      sed -i 's#"mode":"bridge"#"mode":"transparent"#g' $CNI_CONFIG_DIR/10-azure.conflist
    elif [[ "${NETWORK_POLICY}" == "" || "${NETWORK_POLICY}" == "none" ]] && [[ "${NETWORK_MODE}" == "transparent" ]]; then
      sed -i 's#"mode":"bridge"#"mode":"transparent"#g' $CNI_CONFIG_DIR/10-azure.conflist
    fi
    /sbin/ebtables -t nat --list
  fi
}
{{- if NeedsContainerd}}
installContainerd() {
  CURRENT_VERSION=$(containerd -version | cut -d " " -f 3 | sed 's|v||')
  if [[ $CURRENT_VERSION != "${CONTAINERD_VERSION}" ]]; then
    os_lower=$(echo ${OS} | tr '[:upper:]' '[:lower:]')
    if [[ ${OS} == "${UBUNTU_OS_NAME}" ]]; then
      url_path="${os_lower}/${UBUNTU_RELEASE}/multiarch/prod"
    elif [[ ${OS} == "${DEBIAN_OS_NAME}" ]]; then
      url_path="${os_lower}/${UBUNTU_RELEASE}/prod"
    else
      exit 25
    fi
    removeMoby
    removeContainerd
    retrycmd_no_stats 120 5 25 curl https://packages.microsoft.com/config/ubuntu/${UBUNTU_RELEASE}/prod.list >/tmp/microsoft-prod.list || exit 25
    retrycmd 10 5 10 cp /tmp/microsoft-prod.list /etc/apt/sources.list.d/ || exit 25
    retrycmd_no_stats 120 5 25 curl https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor >/tmp/microsoft.gpg || exit 26
    retrycmd 10 5 10 cp /tmp/microsoft.gpg /etc/apt/trusted.gpg.d/ || exit 26
    apt_get_update || exit 99
    apt_get_install 20 30 120 moby-containerd=${CONTAINERD_VERSION}* --allow-downgrades || exit 27
  fi
}
ensureContainerd() {
  wait_for_file 1200 1 /etc/systemd/system/containerd.service.d/exec_start.conf || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  wait_for_file 1200 1 /etc/containerd/config.toml || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  {{- if HasKubeReservedCgroup}}
  wait_for_file 1200 1 /etc/systemd/system/containerd.service.d/kubereserved-slice.conf|| exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  {{- end}}
  systemctlEnableAndStart containerd || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
}
{{end}}
{{- if IsDockerContainerRuntime}}
ensureDocker() {
  DOCKER_SERVICE_EXEC_START_FILE=/etc/systemd/system/docker.service.d/exec_start.conf
  wait_for_file 1200 1 $DOCKER_SERVICE_EXEC_START_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  usermod -aG docker ${ADMINUSER}
  DOCKER_MOUNT_FLAGS_SYSTEMD_FILE=/etc/systemd/system/docker.service.d/clear_mount_propagation_flags.conf
  {{- if HasKubeReservedCgroup}}
  DOCKER_SLICE_FILE=/etc/systemd/system/docker.service.d/kubereserved-slice.conf
  wait_for_file 1200 1 $DOCKER_SLICE_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  {{- end}}
  DOCKER_JSON_FILE=/etc/docker/daemon.json
  for i in $(seq 1 1200); do
    if [ -s $DOCKER_JSON_FILE ]; then
      jq '.' <$DOCKER_JSON_FILE && break
    fi
    if [ $i -eq 1200 ]; then
      exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
    else
      sleep 1
    fi
  done
  systemctlEnableAndStart docker || exit {{GetCSEErrorCode "ERR_DOCKER_START_FAIL"}}
  {{/* Delay start of docker-monitor for 30 mins after booting */}}
  DOCKER_MONITOR_SYSTEMD_TIMER_FILE=/etc/systemd/system/docker-monitor.timer
  wait_for_file 1200 1 $DOCKER_MONITOR_SYSTEMD_TIMER_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  DOCKER_MONITOR_SYSTEMD_FILE=/etc/systemd/system/docker-monitor.service
  wait_for_file 1200 1 $DOCKER_MONITOR_SYSTEMD_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  systemctlEnableAndStart docker-monitor.timer || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
}
{{end}}
{{- if EnableEncryptionWithExternalKms}}
ensureKMS() {
  systemctlEnableAndStart kms || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
}
{{end}}
{{- if IsIPv6Enabled}}
ensureDHCPv6() {
  wait_for_file 3600 1 {{GetDHCPv6ServiceCSEScriptFilepath}} || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  wait_for_file 3600 1 {{GetDHCPv6ConfigCSEScriptFilepath}} || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  systemctlEnableAndStart dhcpv6 || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
  retrycmd 120 5 25 modprobe ip6_tables || exit {{GetCSEErrorCode "ERR_MODPROBE_FAIL"}}
}
{{end}}
ensureKubelet() {
  wait_for_file 1200 1 /etc/sysctl.d/11-aks-engine.conf || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  sysctl_reload 10 5 120 || exit {{GetCSEErrorCode "ERR_SYSCTL_RELOAD"}}
  KUBELET_DEFAULT_FILE=/etc/default/kubelet
  wait_for_file 1200 1 $KUBELET_DEFAULT_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  KUBECONFIG_FILE=/var/lib/kubelet/kubeconfig
  wait_for_file 1200 1 $KUBECONFIG_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  KUBELET_RUNTIME_CONFIG_SCRIPT_FILE=/opt/azure/containers/kubelet.sh
  wait_for_file 1200 1 $KUBELET_RUNTIME_CONFIG_SCRIPT_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  {{- if HasKubeReservedCgroup}}
  KUBERESERVED_SLICE_FILE=/etc/systemd/system/{{- GetKubeReservedCgroup -}}.slice
  wait_for_file 1200 1 $KUBERESERVED_SLICE_FILE || exit {{GetCSEErrorCode "ERR_KUBERESERVED_SLICE_SETUP_FAIL"}}
  KUBELET_SLICE_FILE=/etc/systemd/system/kubelet.service.d/kubereserved-slice.conf
  wait_for_file 1200 1 $KUBELET_SLICE_FILE || exit {{GetCSEErrorCode "ERR_KUBELET_SLICE_SETUP_FAIL"}}
  {{- end}}
  systemctlEnableAndStart kubelet || exit {{GetCSEErrorCode "ERR_KUBELET_START_FAIL"}}
}

ensureAddons() {
  retrycmd 120 5 30 $KUBECTL get pods -l app=kube-addon-manager -n kube-system || exit {{GetCSEErrorCode "ERR_ADDONS_START_FAIL"}}
{{- if not HasCustomPodSecurityPolicy}}
  retrycmd 120 5 30 $KUBECTL get podsecuritypolicy privileged restricted || exit {{GetCSEErrorCode "ERR_ADDONS_START_FAIL"}}
  rm -Rf ${ADDONS_DIR}/init
{{- end}}
  wait_for_file 1200 1 $ADDON_MANAGER_SPEC || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  sed -i "s|${ADDONS_DIR}/init|${ADDONS_DIR}|g" $ADDON_MANAGER_SPEC || exit {{GetCSEErrorCode "ERR_ADDONS_START_FAIL"}}
  {{/* Force re-load all addons because we have changed the source location for addon specs */}}
  retrycmd 120 5 30 ${KUBECTL} delete pods -l app=kube-addon-manager -n kube-system --force --grace-period 0 || exit {{GetCSEErrorCode "ERR_ADDONS_START_FAIL"}}
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
ensureLabelNodes() {
  LABEL_NODES_SCRIPT_FILE=/opt/azure/containers/label-nodes.sh
  wait_for_file 1200 1 $LABEL_NODES_SCRIPT_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  LABEL_NODES_SYSTEMD_FILE=/etc/systemd/system/label-nodes.service
  wait_for_file 1200 1 $LABEL_NODES_SYSTEMD_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
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
  binPath=/usr/local/bin
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
  if [ -f /var/run/reboot-required ] || [ "$NO_OUTBOUND" = "true" ]; then
    return
  fi
  retrycmd 120 5 25 $KUBECTL 2>/dev/null cluster-info || exit {{GetCSEErrorCode "ERR_K8S_RUNNING_TIMEOUT"}}
}
{{- if IsAzurePolicyAddonEnabled}}
ensureLabelExclusionForAzurePolicyAddon() {
  retrycmd 120 5 25 $KUBECTL label ns kube-system control-plane=controller-manager --overwrite 2>/dev/null || exit {{GetCSEErrorCode "ERR_K8S_RUNNING_TIMEOUT"}}
}
{{end}}
ensureEtcd() {
  retrycmd 120 5 25 curl --cacert /etc/kubernetes/certs/ca.crt --cert /etc/kubernetes/certs/etcdclient.crt --key /etc/kubernetes/certs/etcdclient.key ${ETCD_CLIENT_URL}/v2/machines || exit {{GetCSEErrorCode "ERR_ETCD_RUNNING_TIMEOUT"}}
}
createKubeManifestDir() {
  KUBEMANIFESTDIR=/etc/kubernetes/manifests
  mkdir -p $KUBEMANIFESTDIR
}
writeKubeConfig() {
  local DIR=/home/$ADMINUSER/.kube
  local FILE=$DIR/config
{{- if HasBlockOutboundInternet}}
  local SERVER=https://localhost
{{else}}
  local SERVER=$KUBECONFIG_SERVER
{{- end}}
  mkdir -p $DIR
  touch $FILE
  chown $ADMINUSER:$ADMINUSER $DIR $FILE
  chmod 700 $DIR
  chmod 600 $FILE
  set +x
  echo "
---
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: \"$CA_CERTIFICATE\"
    server: $SERVER
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
" >$FILE
  set -x
}
{{- if IsClusterAutoscalerAddonEnabled}}
configClusterAutoscalerAddon() {
  CLUSTER_AUTOSCALER_ADDON_FILE=$ADDONS_DIR/cluster-autoscaler.yaml
  wait_for_file 1200 1 $CLUSTER_AUTOSCALER_ADDON_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  sed -i "s|<clientID>|$(echo $SERVICE_PRINCIPAL_CLIENT_ID | base64)|g" $CLUSTER_AUTOSCALER_ADDON_FILE
  sed -i "s|<clientSec>|$(echo $SERVICE_PRINCIPAL_CLIENT_SECRET | base64)|g" $CLUSTER_AUTOSCALER_ADDON_FILE
  sed -i "s|<subID>|$(echo $SUBSCRIPTION_ID | base64)|g" $CLUSTER_AUTOSCALER_ADDON_FILE
  sed -i "s|<tenantID>|$(echo $TENANT_ID | base64)|g" $CLUSTER_AUTOSCALER_ADDON_FILE
  sed -i "s|<rg>|$(echo $RESOURCE_GROUP | base64)|g" $CLUSTER_AUTOSCALER_ADDON_FILE
}
{{end}}
{{- if IsACIConnectorAddonEnabled}}
configACIConnectorAddon() {
  ACI_CONNECTOR_CREDENTIALS=$(printf '{"clientId": "%s", "clientSecret": "%s", "tenantId": "%s", "subscriptionId": "%s", "activeDirectoryEndpointUrl": "https://login.microsoftonline.com","resourceManagerEndpointUrl": "https://management.azure.com/", "activeDirectoryGraphResourceId": "https://graph.windows.net/", "sqlManagementEndpointUrl": "https://management.core.windows.net:8443/", "galleryEndpointUrl": "https://gallery.azure.com/", "managementEndpointUrl": "https://management.core.windows.net/"}' "$SERVICE_PRINCIPAL_CLIENT_ID" "$SERVICE_PRINCIPAL_CLIENT_SECRET" "$TENANT_ID" "$SUBSCRIPTION_ID" | base64 -w 0)

  openssl req -newkey rsa:4096 -new -nodes -x509 -days 3650 -keyout /etc/kubernetes/certs/aci-connector-key.pem -out /etc/kubernetes/certs/aci-connector-cert.pem -subj "/C=US/ST=CA/L=virtualkubelet/O=virtualkubelet/OU=virtualkubelet/CN=virtualkubelet"
  ACI_CONNECTOR_KEY=$(base64 /etc/kubernetes/certs/aci-connector-key.pem -w0)
  ACI_CONNECTOR_CERT=$(base64 /etc/kubernetes/certs/aci-connector-cert.pem -w0)

  ACI_CONNECTOR_ADDON_FILE=$ADDONS_DIR/aci-connector-deployment.yaml
  wait_for_file 1200 1 $ACI_CONNECTOR_ADDON_FILE || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  sed -i "s|<creds>|$ACI_CONNECTOR_CREDENTIALS|g" $ACI_CONNECTOR_ADDON_FILE
  sed -i "s|<rgName>|$RESOURCE_GROUP|g" $ACI_CONNECTOR_ADDON_FILE
  sed -i "s|<cert>|$ACI_CONNECTOR_CERT|g" $ACI_CONNECTOR_ADDON_FILE
  sed -i "s|<key>|$ACI_CONNECTOR_KEY|g" $ACI_CONNECTOR_ADDON_FILE
}
{{end}}
{{- if IsAzurePolicyAddonEnabled}}
configAzurePolicyAddon() {
  AZURE_POLICY_ADDON_FILE=$ADDONS_DIR/azure-policy-deployment.yaml
  sed -i "s|<resourceId>|/subscriptions/$SUBSCRIPTION_ID/resourceGroups/$RESOURCE_GROUP|g" $AZURE_POLICY_ADDON_FILE
}
{{end}}
configAddons() {
  {{if IsClusterAutoscalerAddonEnabled}}
  if [[ ${CLUSTER_AUTOSCALER_ADDON} == true ]]; then
    configClusterAutoscalerAddon
  fi
  {{end}}
  {{if IsACIConnectorAddonEnabled}}
  if [[ ${ACI_CONNECTOR_ADDON} == True ]]; then
    configACIConnectorAddon
  fi
  {{end}}
  {{if IsAzurePolicyAddonEnabled}}
  configAzurePolicyAddon
  {{end}}
  {{- if not HasCustomPodSecurityPolicy}}
  wait_for_file 1200 1 $POD_SECURITY_POLICY_SPEC || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  mkdir -p $ADDONS_DIR/init && cp $POD_SECURITY_POLICY_SPEC $ADDONS_DIR/init/ || exit {{GetCSEErrorCode "ERR_ADDONS_START_FAIL"}}
  {{- end}}
}
{{- if HasNSeriesSKU}}
{{- /* installNvidiaDrivers is idempotent, it will uninstall itself if it is already installed, and then install anew */}}
installNvidiaDrivers() {
  NVIDIA_DKMS_DIR="/var/lib/dkms/nvidia/${GPU_DV}"
  KERNEL_NAME=$(uname -r)
  if [ -d $NVIDIA_DKMS_DIR ]; then
    dkms remove -m nvidia -v $GPU_DV -k $KERNEL_NAME
  fi
  local log_file="/var/log/nvidia-installer-$(date +%s).log"
  sh $GPU_DEST/nvidia-drivers-$GPU_DV -s -k=$KERNEL_NAME --log-file-name=$log_file -a --no-drm --dkms --utility-prefix="${GPU_DEST}" --opengl-prefix="${GPU_DEST}"
}
configGPUDrivers() {
  {{/* only install the runtime since nvidia-docker2 has a hard dep on docker CE packages. */}}
  {{/* we will manually install nvidia-docker2 */}}
  rmmod nouveau
  echo blacklist nouveau >>/etc/modprobe.d/blacklist.conf
  retrycmd_no_stats 120 5 25 update-initramfs -u || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_CONFIG"}}
  wait_for_apt_locks
  {{/* if the unattened upgrade is turned on, and it may takes 10 min to finish the installation, and we use the 1 second just to try to get the lock more aggressively */}}
  retrycmd 600 1 3600 apt-get -o Dpkg::Options::="--force-confold" install -y nvidia-container-runtime="${NVIDIA_CONTAINER_RUNTIME_VERSION}+${NVIDIA_DOCKER_SUFFIX}" || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_CONFIG"}}
  tmpDir=$GPU_DEST/tmp
  (
    set -e -o pipefail
    cd "${tmpDir}"
    wait_for_apt_locks
    dpkg-deb -R ./nvidia-docker2*.deb "${tmpDir}/pkg" || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_CONFIG"}}
    cp -r ${tmpDir}/pkg/usr/* /usr/ || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_CONFIG"}}
  )
  rm -rf $GPU_DEST/tmp
  retrycmd 120 5 25 pkill -SIGHUP dockerd || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_CONFIG"}}
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

  local packages="make gcc dkms"
  wait_for_apt_locks
  retrycmd 30 5 3600 apt-get -y install "$packages" || exit 90

  local oe_dir=/opt/azure/containers/oe
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
  docker rmi $(docker images --format '{{OpenBraces}}.Repository{{CloseBraces}}:{{OpenBraces}}.Tag{{CloseBraces}}' | grep -vE "${KUBERNETES_VERSION}$|${KUBERNETES_VERSION}-|${KUBERNETES_VERSION}_" | grep 'hyperkube') &
  docker rmi $(docker images --format '{{OpenBraces}}.Repository{{CloseBraces}}:{{OpenBraces}}.Tag{{CloseBraces}}' | grep -vE "${KUBERNETES_VERSION}$|${KUBERNETES_VERSION}-|${KUBERNETES_VERSION}_" | grep 'cloud-controller-manager') &
  docker rmi $(docker images --format '{{OpenBraces}}.Repository{{CloseBraces}}:{{OpenBraces}}.Tag{{CloseBraces}}' | grep -vE "${ETCD_VERSION}$|${ETCD_VERSION}-|${ETCD_VERSION}_" | grep 'etcd') &
  if [ "$IS_HOSTED_MASTER" = "false" ]; then
    docker rmi $(docker images --format '{{OpenBraces}}.Repository{{CloseBraces}}:{{OpenBraces}}.Tag{{CloseBraces}}' | grep 'hcp-tunnel-front') &
    docker rmi $(docker images --format '{{OpenBraces}}.Repository{{CloseBraces}}:{{OpenBraces}}.Tag{{CloseBraces}}' | grep 'kube-svc-redirect') &
    docker rmi $(docker images --format '{{OpenBraces}}.Repository{{CloseBraces}}:{{OpenBraces}}.Tag{{CloseBraces}}' | grep 'nginx') &
  fi

  docker rmi registry:2.7.1 &
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

{{- if IsCustomCloudProfile}}
  {{- if not IsAzureStackCloud}}
ensureCustomCloudRootCertificates() {
    CUSTOM_CLOUD_ROOT_CERTIFICATES="{{GetCustomCloudRootCertificates}}"
    KUBE_CONTROLLER_MANAGER_FILE=/etc/kubernetes/manifests/kube-controller-manager.yaml

    if [ ! -z $CUSTOM_CLOUD_ROOT_CERTIFICATES ]; then
        # Replace placeholder for ssl binding
        if [ -f $KUBE_CONTROLLER_MANAGER_FILE ]; then
            sed -i "s|<volumessl>|- name: ssl\n      hostPath:\n        path: \\/etc\\/ssl\\/certs|g" $KUBE_CONTROLLER_MANAGER_FILE
            sed -i "s|<volumeMountssl>|- name: ssl\n          mountPath: \\/etc\\/ssl\\/certs\n          readOnly: true|g" $KUBE_CONTROLLER_MANAGER_FILE
        fi

        local i=1
        for cert in $(echo $CUSTOM_CLOUD_ROOT_CERTIFICATES | tr ',' '\n')
        do
            echo $cert | base64 -d > "/usr/local/share/ca-certificates/customCloudRootCertificate$i.crt"
            ((i++))
        done

        update-ca-certificates
    else
        if [ -f $KUBE_CONTROLLER_MANAGER_FILE ]; then
            # remove the placeholder for ssl binding
            sed -i "/<volumessl>/d" $KUBE_CONTROLLER_MANAGER_FILE
            sed -i "/<volumeMountssl>/d" $KUBE_CONTROLLER_MANAGER_FILE
        fi
    fi
}

ensureCustomCloudSourcesList() {
    CUSTOM_CLOUD_SOURCES_LIST="{{GetCustomCloudSourcesList}}"

    if [ ! -z $CUSTOM_CLOUD_SOURCES_LIST ]; then
        # Just in case, let's take a back up before we overwrite
        cp /etc/apt/sources.list /etc/apt/sources.list.backup
        echo $CUSTOM_CLOUD_SOURCES_LIST | base64 -d > /etc/apt/sources.list
    fi
}
  {{end}}

configureK8sCustomCloud() {
  {{- if IsAzureStackCloud}}
  export -f ensureAzureStackCertificates
  retrycmd 60 10 30 bash -c ensureAzureStackCertificates
  set +x
  # When AUTHENTICATION_METHOD is client_certificate, the certificate is stored into key valut,
  # And SERVICE_PRINCIPAL_CLIENT_SECRET will be the following json payload with based64 encode
  #{
  #    "data": "$pfxAsBase64EncodedString",
  #    "dataType" :"pfx",
  #    "password": "$password"
  #}
  if [[ ${AUTHENTICATION_METHOD,,} == "client_certificate" ]]; then
    SERVICE_PRINCIPAL_CLIENT_SECRET_DECODED=$(echo ${SERVICE_PRINCIPAL_CLIENT_SECRET} | base64 --decode)
    SERVICE_PRINCIPAL_CLIENT_SECRET_CERT=$(echo $SERVICE_PRINCIPAL_CLIENT_SECRET_DECODED | jq .data)
    SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD=$(echo $SERVICE_PRINCIPAL_CLIENT_SECRET_DECODED | jq .password)

    # trim the starting and ending "
    SERVICE_PRINCIPAL_CLIENT_SECRET_CERT=${SERVICE_PRINCIPAL_CLIENT_SECRET_CERT#'"'}
    SERVICE_PRINCIPAL_CLIENT_SECRET_CERT=${SERVICE_PRINCIPAL_CLIENT_SECRET_CERT%'"'}

    SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD=${SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD#'"'}
    SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD=${SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD%'"'}

    KUBERNETES_FILE_DIR=$(dirname "${AZURE_JSON_PATH}")
    K8S_CLIENT_CERT_PATH="${KUBERNETES_FILE_DIR}/k8s_auth_certificate.pfx"
    echo $SERVICE_PRINCIPAL_CLIENT_SECRET_CERT | base64 --decode >$K8S_CLIENT_CERT_PATH
    # shellcheck disable=SC2002,SC2005
    echo $(cat "${AZURE_JSON_PATH}" |
      jq --arg K8S_CLIENT_CERT_PATH ${K8S_CLIENT_CERT_PATH} '. + {aadClientCertPath:($K8S_CLIENT_CERT_PATH)}' |
      jq --arg SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD ${SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD} '. + {aadClientCertPassword:($SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD)}' |
      jq 'del(.aadClientSecret)') >${AZURE_JSON_PATH}
  fi

  if [[ ${IDENTITY_SYSTEM,,} == "adfs" ]]; then
    # update the tenent id for ADFS environment.
    # shellcheck disable=SC2002,SC2005
    echo $(cat "${AZURE_JSON_PATH}" | jq '.tenantId = "adfs"') >${AZURE_JSON_PATH}
  fi
  set -x

  {{- if not IsAzureCNI}}
  # Decrease eth0 MTU to mitigate Azure Stack's NRP issue
  echo "iface eth0 inet dhcp" | sudo tee -a /etc/network/interfaces
  echo "    post-up /sbin/ifconfig eth0 mtu 1350" | sudo tee -a /etc/network/interfaces
  ifconfig eth0 mtu 1350
  {{end}}

  {{else}}
  ensureCustomCloudRootCertificates
  ensureCustomCloudSourcesList
  {{end}}
}
{{end}}

{{- if IsAzureStackCloud}}
ensureAzureStackCertificates() {
  AZURESTACK_ENVIRONMENT_JSON_PATH="/etc/kubernetes/azurestackcloud.json"
  AZURESTACK_RESOURCE_MANAGER_ENDPOINT=$(jq .resourceManagerEndpoint $AZURESTACK_ENVIRONMENT_JSON_PATH | tr -d '"')
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

  # ensureAzureStackCertificates will be retried if the exit code is not 0
  curl $AZURESTACK_RESOURCE_METADATA_ENDPOINT
  exit $?
}

configureAzureStackInterfaces() {
  set +x

  NETWORK_INTERFACES_FILE="/etc/kubernetes/network_interfaces.json"
  AZURE_CNI_CONFIG_FILE="/etc/kubernetes/interfaces.json"
  AZURESTACK_ENVIRONMENT_JSON_PATH="/etc/kubernetes/azurestackcloud.json"
  SERVICE_MANAGEMENT_ENDPOINT=$(jq -r '.serviceManagementEndpoint' ${AZURESTACK_ENVIRONMENT_JSON_PATH})
  ACTIVE_DIRECTORY_ENDPOINT=$(jq -r '.activeDirectoryEndpoint' ${AZURESTACK_ENVIRONMENT_JSON_PATH})
  RESOURCE_MANAGER_ENDPOINT=$(jq -r '.resourceManagerEndpoint' ${AZURESTACK_ENVIRONMENT_JSON_PATH})

  if [[ ${IDENTITY_SYSTEM,,} == "adfs" ]]; then
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

  if [[ -z $TOKEN ]]; then
    echo "Error generating token for Azure Resource Manager"
    exit 120
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
    "${RESOURCE_MANAGER_ENDPOINT}subscriptions/$SUBSCRIPTION_ID/resourceGroups/$RESOURCE_GROUP/providers/Microsoft.Network/networkInterfaces?api-version=$NETWORK_API_VERSION" >${NETWORK_INTERFACES_FILE}

  if [[ ! -s ${NETWORK_INTERFACES_FILE} ]]; then
    echo "Error fetching network interface configuration for node"
    exit 121
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
      "${RESOURCE_MANAGER_ENDPOINT}${SUBNET_ID:1}?api-version=$NETWORK_API_VERSION" |
      jq '.properties.addressPrefix' -r)

    if [[ -z $SUBNET_PREFIX ]]; then
      echo "Error fetching the subnet address prefix for a subnet ID"
      exit 122
    fi

    # shellcheck disable=SC2001
    AZURE_CNI_CONFIG=$(echo ${AZURE_CNI_CONFIG} | sed "s|$SUBNET_ID|$SUBNET_PREFIX|g")
  done

  echo ${AZURE_CNI_CONFIG} >${AZURE_CNI_CONFIG_FILE}

  chmod 0444 ${AZURE_CNI_CONFIG_FILE}

  set -x
}
{{end}}
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

OS=$(sort -r /etc/*-release | gawk 'match($0, /^(ID=(.*))$/, a) { print toupper(a[2] a[3]); exit }')
UBUNTU_OS_NAME="UBUNTU"
RHEL_OS_NAME="RHEL"
FLATCAR_OS_NAME="FLATCAR"
DEBIAN_OS_NAME="DEBIAN"
if ! echo "${UBUNTU_OS_NAME} ${RHEL_OS_NAME} ${FLATCAR_OS_NAME} ${DEBIAN_OS_NAME}" | grep -q "${OS}"; then
  OS=$(sort -r /etc/*-release | gawk 'match($0, /^(ID_LIKE=(.*))$/, a) { print toupper(a[2] a[3]); exit }')
fi
if [[ ${OS} == "${UBUNTU_OS_NAME}" ]]; then
  UBUNTU_RELEASE=$(lsb_release -r -s)
fi
DOCKER=/usr/bin/docker
if [[ $UBUNTU_RELEASE == "18.04" ]]; then
  export GPU_DV=418.126.02
else
  export GPU_DV=418.40.04
fi
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
  retrycmd 120 5 25 apt-mark $1 walinuxagent ||
    if [[ $1 == "hold" ]]; then
      exit 7
    elif [[ $1 == "unhold" ]]; then
      exit 8
    fi
}

retrycmd() {
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
retrycmd_no_stats() {
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
  packages="apache2-utils apt-transport-https blobfuse=1.1.1 ca-certificates cifs-utils conntrack cracklib-runtime dbus dkms ebtables ethtool fuse gcc git htop iftop init-system-helpers iotop iproute2 ipset iptables jq libpam-pwquality libpwquality-tools linux-headers-$(uname -r) make mount nfs-common pigz socat sysstat traceroute util-linux xz-utils zip"
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
removeMoby() {
  apt_get_purge moby-engine moby-cli || exit 27
}
removeContainerd() {
  apt_get_purge moby-containerd || exit 27
}
installMoby() {
  CURRENT_VERSION=$(dockerd --version | grep "Docker version" | cut -d "," -f 1 | cut -d " " -f 3 | cut -d "+" -f 1)
  if [[ $CURRENT_VERSION != "${MOBY_VERSION}" ]]; then
    removeContainerd
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
  if [[ $OS == $FLATCAR_OS_NAME ]]; then
    apmz_filepath="/opt/bin/apmz"
    export PATH="${PATH}:/opt/bin"
  fi
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
  if [[ $OS == $FLATCAR_OS_NAME ]]; then
    bpftrace_filepath="/opt/bin/$bpftrace_bin"
    tools_filepath="/opt/share/$bpftrace_bin"
    export PATH="${PATH}:/opt/bin"
  fi
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
  hyperkubePath="/home/hyperkube-downloads/${KUBERNETES_VERSION}"
  targetpath="/usr/local/bin"
  if [[ $OS == $FLATCAR_OS_NAME ]]; then
    targetpath="/opt/bin"
  fi
  pullContainerImage $CLI_TOOL ${HYPERKUBE_URL}
  if [[ $CLI_TOOL == "docker" ]]; then
    mkdir -p "$hyperkubePath"
    if docker run --rm --entrypoint "" -v $path:$path ${HYPERKUBE_URL} /bin/bash -c "cp $targetpath/{kubelet,kubectl} $hyperkubePath"; then
      mv "${hyperkubePath}/kubelet" "${targetpath}/kubelet-${KUBERNETES_VERSION}"
      mv "${hyperkubePath}/kubectl" "${targetpath}/kubectl-${KUBERNETES_VERSION}"
      return
    else
      docker run --rm -v $hyperkubePath:$hyperkubePath ${HYPERKUBE_URL} /bin/bash -c "cp /hyperkube $hyperkubePath"
    fi
  else
    img unpack -o "$hyperkubePath" ${HYPERKUBE_URL}
  fi

  cp "${hyperkubePath}/hyperkube" "${targetpath}/kubelet-${KUBERNETES_VERSION}"
  mv "${hyperkubePath}/hyperkube" "${targetpath}/kubectl-${KUBERNETES_VERSION}"
  if [[ $OS == $FLATCAR_OS_NAME ]]; then
    chmod a+x ${targetpath}/kubelet-${KUBERNETES_VERSION} ${targetpath}/kubectl-${KUBERNETES_VERSION}
  fi
}
extractKubeBinaries() {
  KUBE_BINARY_URL=${KUBE_BINARY_URL:-"https://kubernetesartifacts.azureedge.net/kubernetes/v${KUBERNETES_VERSION}/binaries/kubernetes-node-linux-amd64.tar.gz"}
  K8S_TGZ_TMP=${KUBE_BINARY_URL##*/}
  mkdir -p "${K8S_DOWNLOADS_DIR}"
  retrycmd_get_tarball 120 5 "$K8S_DOWNLOADS_DIR/${K8S_TGZ_TMP}" ${KUBE_BINARY_URL} || exit 31
  path=/usr/local/bin
  if [[ $OS == $FLATCAR_OS_NAME ]]; then
    path=/opt/bin
  fi
  tar --transform="s|.*|&-${KUBERNETES_VERSION}|" --show-transformed-names -xzvf "$K8S_DOWNLOADS_DIR/${K8S_TGZ_TMP}" \
    --strip-components=3 -C ${path} kubernetes/node/bin/kubelet kubernetes/node/bin/kubectl
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
ERR_FILE_WATCH_TIMEOUT=6 {{/* Timeout waiting for a file */}}

set -x
if [ -f /opt/azure/containers/provision.complete ]; then
  echo "Already ran to success exiting..."
  exit 0
fi

echo $(date),$(hostname), startcustomscript >>/opt/m

for i in $(seq 1 3600); do
  if [ -s {{GetCSEHelpersScriptFilepath}} ]; then
    grep -Fq '#HELPERSEOF' {{GetCSEHelpersScriptFilepath}} && break
  fi
  if [ $i -eq 3600 ]; then
    exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  else
    sleep 1
  fi
done
sed -i "/#HELPERSEOF/d" {{GetCSEHelpersScriptFilepath}}
source {{GetCSEHelpersScriptFilepath}}
configure_prerequisites

wait_for_file 3600 1 {{GetCSEInstallScriptFilepath}} || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
source {{GetCSEInstallScriptFilepath}}

ensureAPMZ "v0.5.1"
{{- if HasTelemetryEnabled }}
eval "$(apmz bash -n "cse" -t "{{GetLinuxDefaultTelemetryTags}}" --api-keys "{{GetApplicationInsightsTelemetryKeys}}")"
{{else}}
eval "$(apmz bash -d)"
{{end}}

wait_for_file 3600 1 {{GetCSEConfigScriptFilepath}} || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
source {{GetCSEConfigScriptFilepath}}

{{- if IsCustomCloudProfile}}
wait_for_file 3600 1 {{GetCustomCloudConfigCSEScriptFilepath}} || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
source {{GetCustomCloudConfigCSEScriptFilepath }}
{{end}}

set +x
ETCD_PEER_CERT=$(echo ${ETCD_PEER_CERTIFICATES} | cut -d'[' -f 2 | cut -d']' -f 1 | cut -d',' -f $((NODE_INDEX + 1)))
ETCD_PEER_KEY=$(echo ${ETCD_PEER_PRIVATE_KEYS} | cut -d'[' -f 2 | cut -d']' -f 1 | cut -d',' -f $((NODE_INDEX + 1)))
set -x

time_metric "ConfigureAdminUser" configureAdminUser

{{- if HasVHDDistroNodes}}
  {{- if not NeedsContainerd}}
time_metric "CleanupContainerd" cleanUpContainerd
  {{end}}
  {{- if HasNSeriesSKU}}
if [[ ${GPU_NODE} != "true" ]]; then
  time_metric "CleanupGPUDrivers" cleanUpGPUDrivers
fi
  {{else}}
time_metric "CleanupGPUDrivers" cleanUpGPUDrivers
  {{end}}
{{end}}

{{- if HasVHDDistroNodes}}
VHD_LOGS_FILEPATH=/opt/azure/vhd-install.complete
if [ -f $VHD_LOGS_FILEPATH ]; then
  echo "detected golden image pre-install"
  time_metric "CleanUpContainerImages" cleanUpContainerImages
  FULL_INSTALL_REQUIRED=false
else
  if [[ ${IS_VHD} == true ]]; then
    echo "Using VHD distro but file $VHD_LOGS_FILEPATH not found"
    exit {{GetCSEErrorCode "ERR_VHD_FILE_NOT_FOUND"}}
  fi
  FULL_INSTALL_REQUIRED=true
fi
{{else}}
FULL_INSTALL_REQUIRED=true
{{end}}

{{- if not IsVHDDistroForAllNodes}}
if [[ $OS == $UBUNTU_OS_NAME || $OS == $DEBIAN_OS_NAME ]] && [ "$FULL_INSTALL_REQUIRED" = "true" ]; then
  time_metric "InstallDeps" installDeps
  if [[ ${UBUNTU_RELEASE} == "18.04" ]]; then
    overrideNetworkConfig
  fi
  if [[ $OS == $UBUNTU_OS_NAME ]]; then
    time_metric "InstallBcc" installBcc
  fi
  {{- if not IsDockerContainerRuntime}}
  time_metric "InstallImg" installImg
  {{end}}
else
  echo "Golden image; skipping dependencies installation"
fi
{{end}}

if [[ ${UBUNTU_RELEASE} == "18.04" ]]; then
  if apt list --installed | grep 'ntp'; then
    time_metric "EnsureNTP" ensureNTP
  fi
fi

if [[ $OS == $UBUNTU_OS_NAME ]]; then
  time_metric "EnsureAuditD" ensureAuditD
fi

{{- if not IsVHDDistroForAllNodes}}
if [[ $FULL_INSTALL_REQUIRED == "true" ]]; then
  time_metric "InstallBpftrace" installBpftrace
fi
{{end}}

if [[ $OS != $FLATCAR_OS_NAME ]]; then
{{- if NeedsContainerd}}
time_metric "InstallContainerd" installContainerd
{{else}}
time_metric "installMoby" installMoby
{{end}}
fi

if [[ -n ${MASTER_NODE} ]] && [[ -z ${COSMOS_URI} ]]; then
  {{- if IsDockerContainerRuntime}}
  CLI_TOOL="docker"
  {{else}}
  CLI_TOOL="img"
  {{end}}
  time_metric "InstallEtcd" installEtcd $CLI_TOOL
fi

{{/* this will capture the amount of time to install of the network plugin during cse */}}
time_metric "InstallNetworkPlugin" installNetworkPlugin

{{- if HasNSeriesSKU}}
if [[ ${GPU_NODE} == true ]]; then
  if $FULL_INSTALL_REQUIRED; then
    time_metric "DownloadGPUDrivers" downloadGPUDrivers
  fi
  time_metric "EnsureGPUDrivers" ensureGPUDrivers
fi
{{end}}

{{- if and IsDockerContainerRuntime HasPrivateAzureRegistryServer}}
docker login -u $SERVICE_PRINCIPAL_CLIENT_ID -p $SERVICE_PRINCIPAL_CLIENT_SECRET {{GetPrivateAzureRegistryServer}}
{{end}}

time_metric "InstallKubeletAndKubectl" installKubeletAndKubectl

if [[ $OS != $FLATCAR_OS_NAME ]]; then
    time_metric "EnsureRPC" ensureRPC
    time_metric "EnsureCron" ensureCron
fi

time_metric "CreateKubeManifestDir" createKubeManifestDir

{{- if HasDCSeriesSKU}}
if [[ ${SGX_NODE} == true && ! -e "/dev/sgx" ]]; then
  time_metric "InstallSGXDrivers" installSGXDrivers
fi
{{end}}

{{/* create etcd user if we are configured for etcd */}}
if [[ -n ${MASTER_NODE} ]] && [[ -z ${COSMOS_URI} ]]; then
  time_metric "ConfigureEtcdUser" configureEtcdUser
fi

if [[ -n ${MASTER_NODE} ]]; then
  {{/* this step configures all certs */}}
  {{/* both configs etcd/cosmos */}}
  time_metric "ConfigureSecrets" configureSecrets
fi

{{/* configure etcd if we are configured for etcd */}}
if [[ -n ${MASTER_NODE} ]] && [[ -z ${COSMOS_URI} ]]; then
  time_metric "ConfigureEtcd" configureEtcd
else
  time_metric "RemoveEtcd" removeEtcd
fi

{{- if HasCustomSearchDomain}}
wait_for_file 3600 1 {{GetCustomSearchDomainsCSEScriptFilepath}} || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
{{GetCustomSearchDomainsCSEScriptFilepath}} >/opt/azure/containers/setup-custom-search-domain.log 2>&1 || exit {{GetCSEErrorCode "ERR_CUSTOM_SEARCH_DOMAINS_FAIL"}}
{{end}}

{{- if IsDockerContainerRuntime}}
time_metric "EnsureDocker" ensureDocker
{{end}}

time_metric "ConfigureK8s" configureK8s

{{- if IsCustomCloudProfile}}
time_metric "ConfigureK8sCustomCloud" configureK8sCustomCloud
{{- if and IsAzureStackCloud IsAzureCNI}}
time_metric "ConfigureAzureStackInterfaces" configureAzureStackInterfaces
{{end}}
{{end}}

time_metric "ConfigureCNI" configureCNI

if [[ -n ${MASTER_NODE} ]]; then
  time_metric "ConfigAddons" configAddons
  time_metric "WriteKubeConfig" writeKubeConfig
fi

{{- if NeedsContainerd}}
time_metric "EnsureContainerd" ensureContainerd
{{end}}

{{- if and IsHostedMaster EnableHostsConfigAgent}}
time_metric "ConfigPrivateClusterHosts" configPrivateClusterHosts
{{end}}

{{- if EnableEncryptionWithExternalKms}}
if [[ -n ${MASTER_NODE} && ${KMS_PROVIDER_VAULT_NAME} != "" ]]; then
  time_metric "EnsureKMS" ensureKMS
fi
{{end}}

{{/* configure and enable dhcpv6 for ipv6 features */}}
{{- if IsIPv6Enabled}}
time_metric "EnsureDHCPv6" ensureDHCPv6
{{end}}

time_metric "EnsureKubelet" ensureKubelet
if [[ -n ${MASTER_NODE} ]]; then
{{if IsAzurePolicyAddonEnabled}}
  time_metric "EnsureLabelExclusionForAzurePolicyAddon" ensureLabelExclusionForAzurePolicyAddon
{{end}}
  time_metric "EnsureAddons" ensureAddons
fi
time_metric "EnsureJournal" ensureJournal

if [[ -n ${MASTER_NODE} ]]; then
  if version_gte ${KUBERNETES_VERSION} 1.16; then
    time_metric "EnsureLabelNodes" ensureLabelNodes
  fi
{{- if IsAADPodIdentityAddonEnabled}}
  time_metric "EnsureTaints" ensureTaints
{{end}}
  if [[ -z ${COSMOS_URI} ]]; then
    if ! { [ "$FULL_INSTALL_REQUIRED" = "true" ] && [ ${UBUNTU_RELEASE} == "18.04" ]; }; then
      time_metric "EnsureEtcd" ensureEtcd
    fi
  fi
  time_metric "EnsureK8sControlPlane" ensureK8sControlPlane
  {{- if HasClusterInitComponent}}
  if [[ $NODE_INDEX == 0 ]]; then
    retrycmd 120 5 30 $KUBECTL apply -f /opt/azure/containers/cluster-init.yaml || exit {{GetCSEErrorCode "ERR_CLUSTER_INIT_FAIL"}}
  fi
  {{end}}
fi

{{- if not IsVHDDistroForAllNodes}}
if $FULL_INSTALL_REQUIRED; then
  if [[ $OS == $UBUNTU_OS_NAME ]]; then
    {{/* mitigation for bug https://bugs.launchpad.net/ubuntu/+source/linux/+bug/1676635 */}}
    echo 2dd1ce17-079e-403c-b352-a1921ee207ee >/sys/bus/vmbus/drivers/hv_util/unbind
    sed -i "13i\echo 2dd1ce17-079e-403c-b352-a1921ee207ee > /sys/bus/vmbus/drivers/hv_util/unbind\n" /etc/rc.local
  fi
fi
{{end}}

{{- /* re-enable unattended upgrades */}}
rm -f /etc/apt/apt.conf.d/99periodic

{{- if not IsAzureStackCloud}}
if [[ $OS == $UBUNTU_OS_NAME ]]; then
  time_metric "PurgeApt" apt_get_purge apache2-utils &
fi
{{end}}

VALIDATION_ERR=0

{{- if IsHostedMaster }}
API_SERVER_DNS_RETRIES=20
if [[ $API_SERVER_NAME == *.privatelink.* ]]; then
  API_SERVER_DNS_RETRIES=200
fi
RES=$(retrycmd ${API_SERVER_DNS_RETRIES} 1 3 nslookup ${API_SERVER_NAME})
STS=$?
if [[ $STS != 0 ]]; then
    if [[ $RES == *"168.63.129.16"*  ]]; then
        VALIDATION_ERR={{GetCSEErrorCode "ERR_K8S_API_SERVER_AZURE_DNS_LOOKUP_FAIL"}}
    else
        VALIDATION_ERR={{GetCSEErrorCode "ERR_K8S_API_SERVER_DNS_LOOKUP_FAIL"}}
    fi
else
    API_SERVER_CONN_RETRIES=50
    if [[ $API_SERVER_NAME == *.privatelink.* ]]; then
        API_SERVER_CONN_RETRIES=100
    fi
    retrycmd ${API_SERVER_CONN_RETRIES} 1 3 nc -vz ${API_SERVER_NAME} 443 &&
    retrycmd ${API_SERVER_CONN_RETRIES} 1 3 nc -vz ${API_SERVER_NAME} 9000 &&
    retrycmd ${API_SERVER_CONN_RETRIES} 1 3 nc -uvz ${API_SERVER_NAME} 1194 ||
    VALIDATION_ERR={{GetCSEErrorCode "ERR_K8S_API_SERVER_CONN_FAIL"}}
fi

{{end}}

if [ -f /var/run/reboot-required ]; then
  trace_info "RebootRequired" "reboot=true"
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

echo "Custom script finished successfully"
echo $(date),$(hostname), endcustomscript >>/opt/m
mkdir -p /opt/azure/containers && touch /opt/azure/containers/provision.complete
ps auxfww >/opt/azure/provision-ps.log &

exit $VALIDATION_ERR

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
ExecStart={{GetDHCPv6ConfigCSEScriptFilepath}}

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

read -r -d '' NETWORK_CONFIGURATION <<EOC || true
iface eth0 inet6 auto
    up sleep 5
    up dhclient -1 -6 -cf /etc/dhcp/dhclient6.conf -lf /var/lib/dhcp/dhclient6.eth0.leases -v eth0 || true
EOC

add_if_not_exists() {
  grep -qxF "${1}" "${2}" || echo "${1}" >>"${2}"
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

source {{GetCSEHelpersScriptFilepath}}

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

if [[ -z ${COSMOS_URI} ]]; then
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
  ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} get $ETCD_REQUESTHEADER_CLIENT_CA --print-value-only >$K8S_PROXY_CA_CRT_FILEPATH
  ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} get $ETCD_PROXY_KEY --print-value-only >$K8S_PROXY_KEY_FILEPATH
  ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} get $ETCD_PROXY_CERT --print-value-only >$K8S_PROXY_CRT_FILEPATH
  {{- /* Remove whitespace padding at beginning of 1st line */}}
  sed -i '1s/\s//' $K8S_PROXY_CA_CRT_FILEPATH $K8S_PROXY_CRT_FILEPATH $K8S_PROXY_KEY_FILEPATH
  chmod 600 $K8S_PROXY_KEY_FILEPATH
}

write_certs_to_disk_with_retry() {
  for i in $(seq 1 12); do
    write_certs_to_disk && break || sleep 5
  done
}
is_etcd_healthy() {
  for i in $(seq 1 100); do
    ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} endpoint health && break || sleep 5
  done
}
is_etcd_healthy
{{- /* lock file to enable "only 1 master generates certs" */}}
rm -f "${PROXY_CERT_LOCK_FILE}"
mkfifo "${PROXY_CERT_LOCK_FILE}"

ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} lock ${PROXY_CERTS_LOCK_NAME} >"${PROXY_CERT_LOCK_FILE}" &

pid=$!
if read -r lockthis <"${PROXY_CERT_LOCK_FILE}"; then
  if [[ "" == "$(ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} get $ETCD_REQUESTHEADER_CLIENT_CA --print-value-only)" ]]; then
    ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} put $ETCD_REQUESTHEADER_CLIENT_CA " $(cat ${PROXY_CRT})" >/dev/null 2>&1
  fi
  if [[ "" == "$(ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} get $ETCD_PROXY_KEY --print-value-only)" ]]; then
    ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} put $ETCD_PROXY_KEY " $(cat ${PROXY_CLIENT_KEY})" >/dev/null 2>&1
  fi
  if [[ "" == "$(ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} get $ETCD_PROXY_CERT --print-value-only)" ]]; then
    ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} put $ETCD_PROXY_CERT " $(cat ${PROXY_CLIENT_CRT})" >/dev/null 2>&1
  fi
fi
kill $pid
wait $pid
rm -f "${PROXY_CERT_LOCK_FILE}"

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
  if [[ ${CONTAINER_RUNTIME:-docker} != "docker" ]]; then
    healthcheck_command="${crictl} pods"
  fi

  until timeout 60 ${healthcheck_command} >/dev/null; do
    if ((attempt == max_attempts)); then
      echo "Max attempt ${max_attempts} reached! Proceeding to monitor container runtime healthiness."
      break
    fi
    echo "$attempt initial attempt \"${healthcheck_command}\"! Trying again in $attempt seconds..."
    sleep "$((2 ** attempt++))"
  done
  while true; do
    if ! timeout 60 ${healthcheck_command} >/dev/null; then
      echo "Container runtime ${container_runtime_name} failed!"
      if [[ $container_runtime_name == "docker" ]]; then
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

if [[ $# -ne 1 ]]; then
  echo "Usage: health-monitor.sh <container-runtime/kubelet>"
  exit 1
fi

KUBE_HOME="/usr/local/bin"
KUBE_ENV="/etc/default/kube-env"
if [[ -e ${KUBE_ENV} ]]; then
  source "${KUBE_ENV}"
fi

SLEEP_SECONDS=10
component=$1
echo "Start kubernetes health monitoring for ${component}"

if [[ ${component} == "container-runtime" ]]; then
  container_runtime_monitoring
elif [[ ${component} == "kubelet" ]]; then
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
{{if EnableEncryptionWithExternalKms}}
Requires=kms.service
{{end}}

[Service]
Restart=always
EnvironmentFile=/etc/default/kubelet
SuccessExitStatus=143
ExecStartPre=/bin/bash /opt/azure/containers/kubelet.sh
ExecStart=/usr/local/bin/kubelet \
        --node-labels="${KUBELET_NODE_LABELS}" \
        $KUBELET_CONFIG
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
[Install]
WantedBy=multi-user.target
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

KUBECONFIG="$(find /home/*/.kube/config)"
KUBECTL="kubectl --kubeconfig=${KUBECONFIG}"

MASTER_SELECTOR="kubernetes.azure.com/role!=agent,kubernetes.io/role!=agent"
MASTER_LABELS="kubernetes.azure.com/role=master kubernetes.io/role=master node-role.kubernetes.io/master="
AGENT_SELECTOR="kubernetes.azure.com/role!=master,kubernetes.io/role!=master"
AGENT_LABELS="kubernetes.azure.com/role=agent kubernetes.io/role=agent node-role.kubernetes.io/agent="

${KUBECTL} label nodes --overwrite -l $MASTER_SELECTOR $MASTER_LABELS
${KUBECTL} label nodes --overwrite -l $AGENT_SELECTOR $AGENT_LABELS
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
source {{GetCSEHelpersScriptFilepath}}

echo "  dns-search {{GetSearchDomainName}}" | tee -a /etc/network/interfaces.d/50-cloud-init.cfg
systemctl_restart 20 5 10 networking
wait_for_apt_locks
retrycmd 10 5 120 apt-get -y install realmd sssd sssd-tools samba-common samba samba-common python2.7 samba-libs packagekit
wait_for_apt_locks
echo "{{GetSearchDomainRealmPassword}}" | realm join -U {{GetSearchDomainRealmUser}}@$(echo "{{GetSearchDomainName}}" | tr /a-z/ /A-Z/) $(echo "{{GetSearchDomainName}}" | tr /a-z/ /A-Z/)
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
# https://github.com/Azure/AKS/issues/772
fs.inotify.max_user_watches = 1048576
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

var _k8sCloudInitArtifactsUntaintNodesService = []byte(`[Unit]
Description=Untaint nodes when pre-scheduling conditions are fulfilled
After=kubelet.service
[Service]
Restart=always
RestartSec=60
ExecStart=/bin/bash /opt/azure/containers/untaint-nodes.sh
[Install]
WantedBy=multi-user.target
#EOF
`)

func k8sCloudInitArtifactsUntaintNodesServiceBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsUntaintNodesService, nil
}

func k8sCloudInitArtifactsUntaintNodesService() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsUntaintNodesServiceBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/untaint-nodes.service", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsUntaintNodesSh = []byte(`#!/usr/bin/env bash

KUBECONFIG="$(find /home/*/.kube/config)"
KUBECTL="kubectl --kubeconfig=${KUBECONFIG}"
AAD_POD_ID_TAINT_KEY={{GetAADPodIdentityTaintKey}}

if ! ${KUBECTL} get daemonsets -n kube-system -o json | jq -e -r '.items[] | select(.metadata.name == "nmi")' > /dev/null; then
  for node in $(${KUBECTL} get nodes -o json | jq -e -r '.items[] | .metadata.name'); do
    ${KUBECTL} taint nodes $node $AAD_POD_ID_TAINT_KEY:NoSchedule- 2>&1 | grep -v 'not found';
  done
  exit 0
fi
for pod in $(${KUBECTL} get pods -n kube-system -o json | jq -r '.items[] | select(.status.phase == "Running") | .metadata.name'); do
  if [[ "$pod" =~ ^nmi ]]; then
    ${KUBECTL} taint nodes $(${KUBECTL} get pod ${pod} -n kube-system -o json | jq -r '.spec.nodeName') $AAD_POD_ID_TAINT_KEY:NoSchedule- 2>&1 | grep -v 'not found';
  fi;
done
exit 0
#EOF
`)

func k8sCloudInitArtifactsUntaintNodesShBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsUntaintNodesSh, nil
}

func k8sCloudInitArtifactsUntaintNodesSh() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsUntaintNodesShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/untaint-nodes.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitJumpboxcustomdataYml = []byte(`#cloud-config

write_files:

- path: {{GetCSEHelpersScriptFilepath}}
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
- . {{GetCSEHelpersScriptFilepath}}
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

write_files:
- path: {{GetCSEHelpersScriptFilepath}}
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

- path: {{GetCSEInstallScriptFilepath}}
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "provisionInstalls"}}

- path: {{GetCSEConfigScriptFilepath}}
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "provisionConfigs"}}

{{- if not .MasterProfile.IsVHDDistro}}
- path: /opt/azure/containers/provision_cis.sh
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "provisionCIS"}}
{{end}}

{{- if not .MasterProfile.IsVHDDistro}}
  {{- if .MasterProfile.IsAuditDEnabled}}
- path: /etc/audit/rules.d/CIS.rules
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "auditdRules"}}
  {{end}}
{{end}}

{{- if .MasterProfile.IsUbuntu1804}}
  {{- if not .MasterProfile.IsVHDDistro}}
- path: /var/run/reboot-required
  permissions: "0644"
  owner: root
  content: |

  {{end}}
{{end}}

{{- if IsCustomCloudProfile}}
- path: {{GetCustomCloudConfigCSEScriptFilepath}}
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{WrapAsVariable "provisionConfigsCustomCloud"}}
{{end}}

{{- if HasKubeReservedCgroup}}
- path: /etc/systemd/system/{{- GetKubeReservedCgroup -}}.slice
  permissions: "0644"
  owner: root
  content: |
    [Unit]
    Description=Limited resources slice for Kubernetes services
    Documentation=man:systemd.special(7)
    DefaultDependencies=no
    Before=slices.target
    Requires=-.slice
    After=-.slice
    #EOF

- path: /etc/systemd/system/kubelet.service.d/kubereserved-slice.conf
  permissions: "0644"
  owner: root
  content: |
    [Service]
    Slice={{- GetKubeReservedCgroup -}}.slice
    #EOF

  {{if NeedsContainerd}}
- path: /etc/systemd/system/containerd.service.d/kubereserved-slice.conf
  permissions: "0644"
  owner: root
  content: |
    [Service]
    Slice={{- GetKubeReservedCgroup -}}.slice
    #EOF
  {{else}}
- path: /etc/systemd/system/docker.service.d/kubereserved-slice.conf
  permissions: "0644"
  owner: root
  content: |
    [Service]
    Slice={{- GetKubeReservedCgroup -}}.slice
    #EOF
  {{end}}
{{end}}

- path: /etc/systemd/system/kubelet.service
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "kubeletSystemdService"}}

{{- if not .MasterProfile.IsVHDDistro}}
- path: /usr/local/bin/health-monitor.sh
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

{{if IsAADPodIdentityAddonEnabled}}
- path: /opt/azure/containers/untaint-nodes.sh
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "untaintNodesScript"}}

- path: /etc/systemd/system/untaint-nodes.service
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "untaintNodesSystemdService"}}
{{end}}

- path: /etc/apt/apt.conf.d/99periodic
  permissions: "0644"
  owner: root
  content: |
    APT::Periodic::Update-Package-Lists "0";
    APT::Periodic::Download-Upgradeable-Packages "0";
    APT::Periodic::AutocleanInterval "0";
    APT::Periodic::Unattended-Upgrade "0";

{{- if IsIPv6Enabled}}
- path: {{GetDHCPv6ServiceCSEScriptFilepath}}
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "dhcpv6SystemdService"}}

- path: {{GetDHCPv6ConfigCSEScriptFilepath}}
  permissions: "0544"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "dhcpv6ConfigurationScript"}}
{{end}}

{{- if .OrchestratorProfile.KubernetesConfig.RequiresDocker}}
    {{- if not .MasterProfile.IsVHDDistro}}
- path: /etc/systemd/system/docker.service.d/clear_mount_propagation_flags.conf
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "dockerClearMountPropagationFlags"}}
    {{end}}

- path: /etc/systemd/system/docker.service.d/exec_start.conf
  permissions: "0644"
  owner: root
  content: |
    [Service]
    ExecStart=
    ExecStart=/usr/bin/dockerd -H fd:// --storage-driver=overlay2 --bip={{WrapAsParameter "dockerBridgeCidr"}}
    ExecStartPost=/sbin/iptables -P FORWARD ACCEPT
    #EOF

- path: /etc/docker/daemon.json
  permissions: "0644"
  owner: root
  content: |
{{IndentString (GetDockerConfig false) 4}}
{{end}}

{{- if HasCiliumNetworkPlugin}}
- path: /etc/systemd/system/sys-fs-bpf.mount
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "systemdBPFMount"}}
{{end}}

- path: /etc/sysctl.d/11-aks-engine.conf
  permissions: "0644"
  owner: root
  content: |
    {{GetSysctlDConfigKeyVals .MasterProfile.SysctlDConfig}}
    #EOF

{{- if NeedsContainerd}}
- path: /etc/systemd/system/containerd.service.d/exec_start.conf
  permissions: "0644"
  owner: root
  content: |
    [Service]
    ExecStartPre=/sbin/iptables -P FORWARD ACCEPT
    #EOF

- path: /etc/containerd/config.toml
  permissions: "0644"
  owner: root
  content: |
{{IndentString GetContainerdConfig 4}}
    #EOF

  {{- if IsKubenet}}
- path: /etc/containerd/kubenet_template.conf
  permissions: "0644"
  owner: root
  content: |
      {
          "cniVersion": "0.3.1",
          "name": "kubenet",
          "plugins": [{
            "type": "bridge",
            "bridge": "cbr0",
            "mtu": 1500,
            "addIf": "eth0",
            "isGateway": true,
            "ipMasq": false,
            "hairpinMode": false,
            "ipam": {
                "type": "host-local",
                "subnet": "{{` + "`" + `{{.PodCIDR}}` + "`" + `}}",
                "routes": [{ "dst": "0.0.0.0/0" }]
            }
          }]
      }
    {{end}}
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

{{- if EnableAggregatedAPIs}}
- path: /etc/kubernetes/generate-proxy-certs.sh
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "generateProxyCertsScript"}}
{{end}}

{{- if HasCustomSearchDomain}}
- path: {{GetCustomSearchDomainsCSEScriptFilepath}}
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "customSearchDomainsScript"}}
{{end}}

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

{{- if EnableDataEncryptionAtRest}}
- path: /etc/kubernetes/encryption-config.yaml
  permissions: "0600"
  owner: root
  content: |
    kind: EncryptionConfiguration
    apiVersion: apiserver.config.k8s.io/v1
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

{{- if EnableEncryptionWithExternalKms}}
- path: /etc/kubernetes/encryption-config.yaml
  permissions: "0444"
  owner: root
  content: |
    kind: EncryptionConfiguration
    apiVersion: apiserver.config.k8s.io/v1
    resources:
      - resources:
        - secrets
        providers:
        - kms:
            name: azurekmsprovider
            endpoint: unix:///opt/azurekms.socket
            cachesize: 1000
        - identity: {}
{{end}}

MASTER_MANIFESTS_CONFIG_PLACEHOLDER

MASTER_CUSTOM_FILES_PLACEHOLDER

MASTER_CONTAINER_ADDONS_PLACEHOLDER

{{- if or (IsDashboardAddonEnabled) (IsAzurePolicyAddonEnabled)}}
- path: /etc/kubernetes/addons/init/namespaces.yaml
  permissions: "0644"
  owner: root
  content: |
  {{- if IsDashboardAddonEnabled}}
    apiVersion: v1
    kind: Namespace
    metadata:
      name: kubernetes-dashboard
      labels:
        addonmanager.kubernetes.io/mode: EnsureExists
    ---
  {{- end}}
  {{- if IsAzurePolicyAddonEnabled}}
    apiVersion: v1
    kind: Namespace
    metadata:
      name: gatekeeper-system
      labels:
        addonmanager.kubernetes.io/mode: EnsureExists
  {{- end}}
{{- end}}

- path: /etc/default/kubelet
  permissions: "0644"
  owner: root
  content: |
    KUBELET_CONFIG={{GetKubeletConfigKeyVals .MasterProfile.KubernetesConfig}}
{{- if IsKubernetesVersionGe "1.16.0"}}
    KUBELET_NODE_LABELS={{GetMasterKubernetesLabels "',variables('labelResourceGroup'),'"}}
{{else}}
    KUBELET_NODE_LABELS={{GetMasterKubernetesLabelsDeprecated "',variables('labelResourceGroup'),'"}}
{{end}}
{{- if IsCustomCloudProfile }}
    AZURE_ENVIRONMENT_FILEPATH=/etc/kubernetes/azurestackcloud.json
{{end}}
    #EOF

- path: /opt/azure/containers/kubelet.sh
  permissions: "0755"
  owner: root
  content: |
    #!/bin/bash
    set -e
    MOUNT_DIR=/var/lib/kubelet
    mkdir -p $MOUNT_DIR /var/lib/cni
    if ! [[ $(findmnt -rno SOURCE,TARGET ${MOUNT_DIR}) ]]; then
      mount --bind $MOUNT_DIR $MOUNT_DIR
    fi
    mount --make-shared $MOUNT_DIR
    PRIVATE_IP=$(hostname -i | cut -d" " -f1)
{{- if IsMasterVirtualMachineScaleSets}}
    PRIVATE_IP=$(hostname -i | cut -d" " -f1)
    sed -i "s|<SERVERIP>|https://$PRIVATE_IP:443|g" "/var/lib/kubelet/kubeconfig"
{{end}}
{{- if gt .MasterProfile.Count 1}}
    {{- /* Redirect ILB (4443) traffic to port 443 (ELB) in the prerouting chain */}}
    iptables -t nat -A PREROUTING -p tcp --dport 4443 -j REDIRECT --to-port 443
{{end}}
    sed -i "s|<advertiseAddr>|$PRIVATE_IP|g" /etc/kubernetes/manifests/kube-apiserver.yaml
{{- if EnableDataEncryptionAtRest }}
    sed -i "s|<etcdEncryptionSecret>|\"{{WrapAsParameter "etcdEncryptionKey"}}\"|g" /etc/kubernetes/encryption-config.yaml
{{end}}
{{- if eq .OrchestratorProfile.KubernetesConfig.NetworkPolicy "calico"}}
    sed -i "s|<kubeClusterCidr>|{{WrapAsParameter "kubeClusterCidr"}}|g" /etc/kubernetes/addons/calico.yaml
    {{- if eq .OrchestratorProfile.KubernetesConfig.NetworkPlugin "azure"}}
    sed -i "/Start of install-cni initContainer/,/End of install-cni initContainer/d" /etc/kubernetes/addons/calico.yaml
    {{else}}
    sed -i "s|<calicoIPAMConfig>|{\"type\": \"host-local\", \"subnet\": \"usePodCidr\"}|g" /etc/kubernetes/addons/calico.yaml
    sed -i "s|azv|cali|g" /etc/kubernetes/addons/calico.yaml
    {{end}}
{{end}}
{{- if eq .OrchestratorProfile.KubernetesConfig.NetworkPlugin "flannel"}}
    sed -i "s|<kubeClusterCidr>|{{WrapAsParameter "kubeClusterCidr"}}|g" /etc/kubernetes/addons/flannel.yaml
{{end}}
    #EOF

{{- if not HasCosmosEtcd  }}
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
    if [[ ! -s /etc/environment ]]; then
        {{- /* /etc/environment is empty, which will break subsequent sed commands
               Append a blank line... */}}
        echo "" >> /etc/environment
    fi
  {{- if IsMasterVirtualMachineScaleSets}}
    MASTER_VM_NAME=$(hostname)
    MASTER_VM_NAME_BASE=$(hostname | sed "s/.$//")
    MASTER_FIRSTADDR={{WrapAsParameter "firstConsecutiveStaticIP"}}
    MASTER_INDEX=$(hostname | tail -c 2)
    PRIVATE_IP=$(hostname -i | cut -d" " -f1)
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
    /bin/echo DAEMON_ARGS=--name $MASTER_VM_NAME --peer-client-cert-auth --peer-trusted-ca-file={{WrapAsVariable "etcdCaFilepath"}} --peer-cert-file=/etc/kubernetes/certs/etcdpeer$MASTER_INDEX.crt --peer-key-file=/etc/kubernetes/certs/etcdpeer$MASTER_INDEX.key --initial-advertise-peer-urls "https://$PRIVATE_IP:$ETCD_SERVER_PORT" --listen-peer-urls "https://$PRIVATE_IP:$ETCD_SERVER_PORT" --client-cert-auth --trusted-ca-file={{WrapAsVariable "etcdCaFilepath"}} --cert-file={{WrapAsVariable "etcdServerCertFilepath"}} --key-file={{WrapAsVariable "etcdServerKeyFilepath"}} --advertise-client-urls "https://$PRIVATE_IP:$ETCD_CLIENT_PORT" --listen-client-urls "https://$PRIVATE_IP:$ETCD_CLIENT_PORT,https://127.0.0.1:$ETCD_CLIENT_PORT" --initial-cluster-token "k8s-etcd-cluster" --initial-cluster $MASTER_URLS --data-dir "/var/lib/etcddisk" --initial-cluster-state "new" --listen-metrics-urls "http://$PRIVATE_IP:2480" --quota-backend-bytes={{GetEtcdStorageLimitGB}} | tee -a /etc/default/etcd
  {{else}}
    sudo sed -i "1iETCDCTL_ENDPOINTS=https://127.0.0.1:2379" /etc/environment
    sudo sed -i "1iETCDCTL_CA_FILE={{WrapAsVariable "etcdCaFilepath"}}" /etc/environment
    sudo sed -i "1iETCDCTL_KEY_FILE={{WrapAsVariable "etcdClientKeyFilepath"}}" /etc/environment
    sudo sed -i "1iETCDCTL_CERT_FILE={{WrapAsVariable "etcdClientCertFilepath"}}" /etc/environment
    sudo sed -i "/^DAEMON_ARGS=/d" /etc/default/etcd
    /bin/echo DAEMON_ARGS=--name "{{WrapAsVerbatim "variables('masterVMNames')[copyIndex(variables('masterOffset'))]"}}" --peer-client-cert-auth --peer-trusted-ca-file={{WrapAsVariable "etcdCaFilepath"}} --peer-cert-file={{WrapAsVerbatim "variables('etcdPeerCertFilepath')[copyIndex(variables('masterOffset'))]"}} --peer-key-file={{WrapAsVerbatim "variables('etcdPeerKeyFilepath')[copyIndex(variables('masterOffset'))]"}} --initial-advertise-peer-urls "{{WrapAsVerbatim "variables('masterEtcdPeerURLs')[copyIndex(variables('masterOffset'))]"}}" --listen-peer-urls "{{WrapAsVerbatim "variables('masterEtcdPeerURLs')[copyIndex(variables('masterOffset'))]"}}" --client-cert-auth --trusted-ca-file={{WrapAsVariable "etcdCaFilepath"}} --cert-file={{WrapAsVariable "etcdServerCertFilepath"}} --key-file={{WrapAsVariable "etcdServerKeyFilepath"}} --advertise-client-urls "{{WrapAsVerbatim "variables('masterEtcdClientURLs')[copyIndex(variables('masterOffset'))]"}}" --listen-client-urls "{{WrapAsVerbatim "concat(variables('masterEtcdClientURLs')[copyIndex(variables('masterOffset'))], ',https://127.0.0.1:', variables('masterEtcdClientPort'))"}}" --initial-cluster-token "k8s-etcd-cluster" --initial-cluster {{WrapAsVerbatim "variables('masterEtcdClusterStates')[div(variables('masterCount'), 2)]"}} --data-dir "/var/lib/etcddisk" --initial-cluster-state "new" --listen-metrics-urls "{{WrapAsVerbatim "variables('masterEtcdMetricURLs')[copyIndex(variables('masterOffset'))]"}}" --quota-backend-bytes={{GetEtcdStorageLimitGB}} | tee -a /etc/default/etcd
  {{end}}
{{end}}
    #EOF

{{- if IsCustomCloudProfile}}
- path: "/etc/kubernetes/azurestackcloud.json"
  permissions: "0600"
  owner: "root"
  content: |
    {{WrapAsVariable "environmentJSON"}}
{{end}}

disk_setup:
  /dev/disk/azure/scsi1/lun0:
    table_type: gpt
    layout: true
    overwrite: false
fs_setup:
  - label: etcd_disk
    filesystem: ext4
    device: /dev/disk/azure/scsi1/lun0
    extra_opts:
      - -E
      - lazy_itable_init=1,lazy_journal_init=1
{{- /* ephemeral (/mnt) filesystem is explicitly configured, see: */}}
{{- /* https://bugs.launchpad.net/cloud-init/+bug/1879552 */}}
  - label: ephemeral0
    filesystem: ext4
    device: ephemeral0.1
    replace_fs: ntfs
mounts:
  - - LABEL=etcd_disk
    - /var/lib/etcddisk
runcmd:
- set -x
- . {{GetCSEHelpersScriptFilepath}}
- aptmarkWALinuxAgent hold{{GetKubernetesMasterPreprovisionYaml}}
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
{{- if .RequiresCloudproviderConfig}}
- path: /etc/kubernetes/azure.json
  permissions: "0600"
  owner: root
  content: |
    #EOF
{{end}}

- path: {{GetCSEHelpersScriptFilepath}}
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

- path: {{GetCSEInstallScriptFilepath}}
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "provisionInstalls"}}

- path: {{GetCSEConfigScriptFilepath}}
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "provisionConfigs"}}

{{- if not .IsVHDDistro}}
- path: /opt/azure/containers/provision_cis.sh
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "provisionCIS"}}
{{end}}

{{- if not .IsVHDDistro}}
  {{- if .IsAuditDEnabled}}
- path: /etc/audit/rules.d/CIS.rules
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "auditdRules"}}
  {{end}}
{{end}}

{{- if .IsUbuntu1804}}
  {{- if not .IsVHDDistro}}
- path: /var/run/reboot-required
  permissions: "0644"
  owner: root
  content: |

  {{end}}
{{end}}

{{- if IsCustomCloudProfile}}
- path: {{GetCustomCloudConfigCSEScriptFilepath}}
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{WrapAsVariable "provisionConfigsCustomCloud"}}
{{end}}

{{- if HasKubeReservedCgroup}}
- path: /etc/systemd/system/{{- GetKubeReservedCgroup -}}.slice
  permissions: "0644"
  owner: root
  content: |
    [Unit]
    Description=Limited resources slice for Kubernetes services
    Documentation=man:systemd.special(7)
    DefaultDependencies=no
    Before=slices.target
    Requires=-.slice
    After=-.slice
    #EOF

- path: /etc/systemd/system/kubelet.service.d/kubereserved-slice.conf
  permissions: "0644"
  owner: root
  content: |
    [Service]
    Slice={{- GetKubeReservedCgroup -}}.slice
    #EOF

  {{if NeedsContainerd}}
- path: /etc/systemd/system/containerd.service.d/kubereserved-slice.conf
  permissions: "0644"
  owner: root
  content: |
    [Service]
    Slice={{- GetKubeReservedCgroup -}}.slice
    #EOF
  {{else}}
- path: /etc/systemd/system/docker.service.d/kubereserved-slice.conf
  permissions: "0644"
  owner: root
  content: |
    [Service]
    Slice={{- GetKubeReservedCgroup -}}.slice
    #EOF
  {{end}}
{{end}}

- path: /etc/systemd/system/kubelet.service
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "kubeletSystemdService"}}

{{- if not .IsVHDDistro}}
    {{- if .IsFlatcar}}
- path: /opt/bin/health-monitor.sh
    {{else}}
- path: /usr/local/bin/health-monitor.sh
    {{- end}}
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

- path: /etc/apt/apt.conf.d/99periodic
  permissions: "0644"
  owner: root
  content: |
    APT::Periodic::Update-Package-Lists "0";
    APT::Periodic::Download-Upgradeable-Packages "0";
    APT::Periodic::AutocleanInterval "0";
    APT::Periodic::Unattended-Upgrade "0";

{{- if IsIPv6Enabled}}
- path: {{GetDHCPv6ServiceCSEScriptFilepath}}
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "dhcpv6SystemdService"}}

- path: {{GetDHCPv6ConfigCSEScriptFilepath}}
  permissions: "0544"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "dhcpv6ConfigurationScript"}}
{{end}}

{{- if .KubernetesConfig.RequiresDocker}}
    {{- if not .IsFlatcar}}
        {{- if not .IsVHDDistro}}
- path: /etc/systemd/system/docker.service.d/clear_mount_propagation_flags.conf
  permissions: "0644"
  encoding: gzip
  owner: "root"
  content: !!binary |
    {{CloudInitData "dockerClearMountPropagationFlags"}}
        {{- end}}
    {{- end}}

- path: /etc/systemd/system/docker.service.d/exec_start.conf
  permissions: "0644"
  owner: root
  content: |
    [Service]
    ExecStart=
    {{- if .IsFlatcar}}
    ExecStart=/usr/bin/env PATH=${TORCX_BINDIR}:${PATH} ${TORCX_BINDIR}/dockerd --host=fd:// --containerd=/var/run/docker/libcontainerd/docker-containerd.sock --storage-driver=overlay2 --bip={{WrapAsParameter "dockerBridgeCidr"}} $DOCKER_SELINUX $DOCKER_OPTS $DOCKER_CGROUPS $DOCKER_OPT_BIP $DOCKER_OPT_MTU $DOCKER_OPT_IPMASQ
    {{else}}
    ExecStart=/usr/bin/dockerd -H fd:// --storage-driver=overlay2 --bip={{WrapAsParameter "dockerBridgeCidr"}}
    {{- end}}
    ExecStartPost=/sbin/iptables -P FORWARD ACCEPT
    #EOF

- path: /etc/docker/daemon.json
  permissions: "0644"
  owner: root
  content: |
{{IndentString (GetDockerConfig (IsNSeriesSKU .VMSize)) 4}}
{{end}}

{{- if HasCiliumNetworkPlugin}}
- path: /etc/systemd/system/sys-fs-bpf.mount
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "systemdBPFMount"}}
{{end}}

- path: /etc/sysctl.d/11-aks-engine.conf
  permissions: "0644"
  owner: root
  content: |
    {{GetSysctlDConfigKeyVals .SysctlDConfig}}
    #EOF

{{- if NeedsContainerd}}
- path: /etc/systemd/system/containerd.service.d/exec_start.conf
  permissions: "0644"
  owner: root
  content: |
    [Service]
    ExecStartPre=/sbin/iptables -P FORWARD ACCEPT
    #EOF

- path: /etc/containerd/config.toml
  permissions: "0644"
  owner: root
  content: |
{{IndentString GetContainerdConfig 4}}
    #EOF

  {{if IsKubenet }}
- path: /etc/containerd/kubenet_template.conf
  permissions: "0644"
  owner: root
  content: |
      {
          "cniVersion": "0.3.1",
          "name": "kubenet",
          "plugins": [{
            "type": "bridge",
            "bridge": "cbr0",
            "mtu": 1500,
            "addIf": "eth0",
            "isGateway": true,
            "ipMasq": false,
            "hairpinMode": false,
            "ipam": {
                "type": "host-local",
                "subnet": "{{` + "`" + `{{.PodCIDR}}` + "`" + `}}",
                "routes": [{ "dst": "0.0.0.0/0" }]
            }
          }]
      }
  {{end}}
{{end}}

{{- if IsNSeriesSKU .VMSize}}
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

{{- if HasCustomSearchDomain}}
- path: {{GetCustomSearchDomainsCSEScriptFilepath}}
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "customSearchDomainsScript"}}
{{end}}

{{if and IsHostedMaster EnableHostsConfigAgent}}
- path: /opt/azure/containers/reconcilePrivateHosts.sh
  permissions: "0744"
  owner: root
  content: |
    #!/usr/bin/env bash
    set -o nounset
    set -o pipefail

    SLEEP_SECONDS=15
    clusterFQDN={{WrapAsVariable "kubernetesAPIServerIP"}}
    if [[ $clusterFQDN != *.privatelink.* ]]; then
      echo "skip reconcile hosts for $clusterFQDN since it's not AKS private cluster"
      exit 0
    fi
    echo "clusterFQDN: $clusterFQDN"

    function get-apiserver-ip-from-tags() {
      tags=$(curl -sSL -H "Metadata: true" "http://169.254.169.254/metadata/instance/compute/tags?api-version=2019-03-11&format=text")
      if [ "$?" == "0" ]; then
        IFS=";" read -ra tagList <<< "$tags"
        for i in "${tagList[@]}"; do
          tagKey=$(cut -d":" -f1 <<<$i)
          tagValue=$(cut -d":" -f2 <<<$i)
          if [ "$tagKey" == "aksAPIServerIPAddress" ]; then
            echo -n "$tagValue"
            return
          fi
        done
      fi

      echo -n ""
    }

    while true; do
      clusterIP=$(get-apiserver-ip-from-tags)
      if [ -z $clusterIP ]; then
        sleep "${SLEEP_SECONDS}"
        continue
      fi

      if grep "$clusterIP $clusterFQDN" /etc/hosts; then
        echo "$clusterFQDN has already been set to $clusterIP"
      else
        sudo sed -i "/$clusterFQDN/d" /etc/hosts
        sudo sed -i "\$a$clusterIP $clusterFQDN" /etc/hosts
        echo "Updated $clusterFQDN to $clusterIP"
      fi
      sleep "${SLEEP_SECONDS}"
    done

- path: /etc/systemd/system/reconcile-private-hosts.service
  permissions: "0644"
  owner: root
  content: |
    [Unit]
    Description=Reconcile /etc/hosts file for private cluster
    [Service]
    Type=simple
    Restart=on-failure
    ExecStart=/bin/bash /opt/azure/containers/reconcilePrivateHosts.sh
    [Install]
    WantedBy=multi-user.target
{{end}}

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
    KUBELET_CONFIG={{GetKubeletConfigKeyVals .KubernetesConfig }}
{{- if IsKubernetesVersionGe "1.16.0"}}
    KUBELET_NODE_LABELS={{GetAgentKubernetesLabels . "',variables('labelResourceGroup'),'"}}
{{else}}
    KUBELET_NODE_LABELS={{GetAgentKubernetesLabelsDeprecated . "',variables('labelResourceGroup'),'"}}
{{end}}
{{- if IsCustomCloudProfile }}
    AZURE_ENVIRONMENT_FILEPATH=/etc/kubernetes/azurestackcloud.json
{{end}}
    #EOF

- path: /opt/azure/containers/kubelet.sh
  permissions: "0755"
  owner: root
  content: |
    #!/bin/bash
    MOUNT_DIR=/var/lib/kubelet
    mkdir -p $MOUNT_DIR /var/lib/cni
    if ! [[ $(findmnt -rno SOURCE,TARGET ${MOUNT_DIR}) ]]; then
      mount --bind $MOUNT_DIR $MOUNT_DIR
    fi
    mount --make-shared $MOUNT_DIR
{{- if and (IsVirtualMachineScaleSets .) IsAADPodIdentityAddonEnabled UseManagedIdentity}}
    {{- /* Disable TCP access to IMDS endpoint, aad-pod-identity nmi component will provide a complementary iptables rule to re-route this traffic */}}
    iptables -A OUTPUT -s 127.0.0.1/32 -d 169.254.169.254/32 -p tcp -m tcp --dport 80 -j DROP
{{end}}
{{- if not IsIPMasqAgentEnabled}}
    {{if IsAzureCNI}}
    iptables -t nat -A POSTROUTING -m iprange ! --dst-range 168.63.129.16 -m addrtype ! --dst-type local ! -d {{WrapAsParameter "vnetCidr"}} -j MASQUERADE
    {{end}}
{{end}}
    #EOF

{{- if IsCustomCloudProfile}}
- path: "/etc/kubernetes/azurestackcloud.json"
  permissions: "0600"
  owner: "root"
  content: |
    {{WrapAsVariable "environmentJSON"}}
{{end}}

{{- if .IsFlatcar}}
- path: "/etc/kubernetes/manifests/.keep"

  {{- if .KubernetesConfig.RequiresDocker}}
groups:
  - docker: [{{WrapAsParameter "linuxAdminUsername"}}]
  {{end}}

coreos:
  units:
    - name: kubelet.service
      enable: true
      drop-ins:
        - name: "10-flatcar.conf"
          content: |
            [Unit]
            Requires=rpc-statd.service
            ConditionPathExists=
            ConditionPathExists=/opt/bin/kubelet
            [Service]
            ExecStart=
            ExecStart=/opt/bin/kubelet \
              --enable-server \
              --node-labels="${KUBELET_NODE_LABELS}" \
              --v=2 \
              --volume-plugin-dir=/etc/kubernetes/volumeplugins \
              $KUBELET_CONFIG $KUBELET_OPTS \
              $KUBELET_REGISTER_NODE $KUBELET_REGISTER_WITH_TAINTS
    - name: kubelet-monitor.service
      enable: true
      drop-ins:
        - name: "10-flatcar.conf"
          content: |
            [Service]
            ExecStart=
            ExecStart=/opt/bin/health-monitor.sh kubelet
    - name: docker-monitor.service
      enable: true
      drop-ins:
        - name: "10-flatcar.conf"
          content: |
            [Service]
            ExecStart=
            ExecStart=/opt/bin/health-monitor.sh container-runtime
    - name: rpcbind.service
      enable: true
{{else}}
runcmd:
- set -x
- . {{GetCSEHelpersScriptFilepath}}
- aptmarkWALinuxAgent hold{{GetKubernetesAgentPreprovisionYaml .}}
{{- end}}
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

var _k8sKubernetesparamsT = []byte(`{{if IsHostedMaster}}
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
    "kubeBinaryURL": {
      "defaultValue": "",
      "metadata": {
        "description": "The package tarball URL to extract kubelet and kubectl binaries from."
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
    "cloudproviderConfig": {
      "type": "object",
      "defaultValue": {
        "cloudProviderBackoff": false,
        "cloudProviderBackoffMode": "v1",
        "cloudProviderBackoffRetries": 10,
        "cloudProviderBackoffJitter": "0",
        "cloudProviderBackoffDuration": 0,
        "cloudProviderBackoffExponent": "0",
        "cloudProviderRateLimit": false,
        "cloudProviderRateLimitQPS": "0",
        "cloudProviderRateLimitQPSWrite": "0",
        "cloudProviderRateLimitBucket": 0,
        "cloudProviderRateLimitBucketWrite": 0,
        "cloudProviderDisableOutboundSNAT": false
      }
    },
    "mobyVersion": {
      "defaultValue": "19.03.12",
      "metadata": {
        "description": "The Azure Moby build version"
      },
      "allowedValues": [
         "3.0.1",
         "3.0.2",
         "3.0.3",
         "3.0.4",
         "3.0.5",
         "3.0.6",
         "3.0.7",
         "3.0.8",
         "3.0.10",
         "3.0.11",
         "3.0.12",
         "3.0.13",
         "19.03.11",
         "19.03.12"
       ],
      "type": "string"
    },
    "containerdVersion": {
      "defaultValue": "1.3.2",
      "metadata": {
        "description": "The Azure Moby build version"
      },
      "allowedValues": [
         "1.3.2"
       ],
      "type": "string"
    },
    "networkPolicy": {
      "defaultValue": "{{.OrchestratorProfile.KubernetesConfig.NetworkPolicy}}",
      "metadata": {
        "description": "The network policy enforcement to use (calico|cilium|antrea); 'none' and 'azure' here for backwards compatibility"
      },
      "allowedValues": [
        "",
        "none",
        "azure",
        "calico",
        "cilium",
        "antrea"
      ],
      "type": "string"
    },
    "networkPlugin": {
      "defaultValue": "{{.OrchestratorProfile.KubernetesConfig.NetworkPlugin}}",
      "metadata": {
        "description": "The network plugin to use for Kubernetes (kubenet|azure|flannel|cilium|antrea)"
      },
      "allowedValues": [
        "kubenet",
        "azure",
        "flannel",
        "cilium",
        "antrea"
      ],
      "type": "string"
    },
    "networkMode": {
      "defaultValue": "{{.OrchestratorProfile.KubernetesConfig.NetworkMode}}",
      "metadata": {
        "description": "The network mode to use for CNI (transparent|bridge)"
      },
      "allowedValues": [
        "",
        "transparent",
        "bridge"
      ],
      "type": "string"
    },
    "containerRuntime": {
      "defaultValue": "{{.OrchestratorProfile.KubernetesConfig.ContainerRuntime}}",
      "metadata": {
        "description": "The container runtime to use (docker|containerd)"
      },
      "allowedValues": [
        "docker",
        "containerd"
      ],
      "type": "string"
    },
    "containerdDownloadURLBase": {
      "defaultValue": "https://storage.googleapis.com/cri-containerd-release/",
      "type": "string"
    },
    "cniPluginsURL": {
      "defaultValue": "https://kubernetesartifacts.azureedge.net/cni-plugins/v0.7.6/binaries/cni-plugins-amd64-v0.7.6.tgz",
      "type": "string"
    },
    "vnetCniLinuxPluginsURL": {
      "defaultValue": "https://kubernetesartifacts.azureedge.net/azure-cni/v1.0.30/binaries/azure-vnet-cni-linux-amd64-v1.0.30.tgz",
      "type": "string"
    },
    "vnetCniWindowsPluginsURL": {
      "defaultValue": "https://kubernetesartifacts.azureedge.net/azure-cni/v1.0.30/binaries/azure-vnet-cni-windows-amd64-v1.0.30.zip",
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
{{ else if and UseManagedIdentity IsHostedMaster}}
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
{{if HasCustomNodesDNS}}
    ,"dnsServer": {
      "defaultValue": "",
      "metadata": {
        "description": "DNS Server IP"
      },
      "type": "string"
    }
{{end}}

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

var _k8sKuberneteswindowsfunctionsPs1 = []byte(`# This filter removes null characters (\0) which are captured in nssm.exe output when logged through powershell
filter RemoveNulls { $_ -replace '\0', '' }

filter Timestamp { "$(Get-Date -Format o): $_" }

function Write-Log($message) {
    $msg = $message | Timestamp
    Write-Output $msg
}

function DownloadFileOverHttp {
    Param(
        [Parameter(Mandatory = $true)][string]
        $Url,
        [Parameter(Mandatory = $true)][string]
        $DestinationPath
    )

    # First check to see if a file with the same name is already cached on the VHD
    $fileName = [IO.Path]::GetFileName($Url)

    $search = @()
    if (Test-Path $global:CacheDir) {
        $search = [IO.Directory]::GetFiles($global:CacheDir, $fileName, [IO.SearchOption]::AllDirectories)
    }

    if ($search.Count -ne 0) {
        Write-Log "Using cached version of $fileName - Copying file from $($search[0]) to $DestinationPath"
        Copy-Item -Path $search[0] -Destination $DestinationPath -Force
    }
    else {
        $secureProtocols = @()
        $insecureProtocols = @([System.Net.SecurityProtocolType]::SystemDefault, [System.Net.SecurityProtocolType]::Ssl3)

        foreach ($protocol in [System.Enum]::GetValues([System.Net.SecurityProtocolType])) {
            if ($insecureProtocols -notcontains $protocol) {
                $secureProtocols += $protocol
            }
        }
        [System.Net.ServicePointManager]::SecurityProtocol = $secureProtocols

        $oldProgressPreference = $ProgressPreference
        $ProgressPreference = 'SilentlyContinue'

        $downloadTimer = [System.Diagnostics.Stopwatch]::StartNew()
        Invoke-WebRequest $Url -UseBasicParsing -OutFile $DestinationPath -Verbose
        $downloadTimer.Stop()

        if ($global:AppInsightsClient -ne $null) {
            $event = New-Object "Microsoft.ApplicationInsights.DataContracts.EventTelemetry"
            $event.Name = "FileDownload"
            $event.Properties["FileName"] = $fileName
            $event.Metrics["DurationMs"] = $downloadTimer.ElapsedMilliseconds
            $global:AppInsightsClient.TrackEvent($event)
        }

        $ProgressPreference = $oldProgressPreference
        Write-Log "Downloaded file to $DestinationPath"
    }
}

function Get-ProvisioningScripts {
    Write-Log "Getting provisioning scripts"
    DownloadFileOverHttp -Url $global:ProvisioningScriptsPackageUrl -DestinationPath 'c:\k\provisioningscripts.zip'
    Expand-Archive -Path 'c:\k\provisioningscripts.zip' -DestinationPath 'c:\k' -Force
    Remove-Item -Path 'c:\k\provisioningscripts.zip' -Force
}

function Get-WindowsVersion {
    $systemInfo = Get-ItemProperty -Path "HKLM:SOFTWARE\Microsoft\Windows NT\CurrentVersion"
    return "$($systemInfo.CurrentBuildNumber).$($systemInfo.UBR)"
}

function Get-CniVersion {
    switch ($global:NetworkPlugin) {
        "azure" {
            if ($global:VNetCNIPluginsURL -match "(v[0-9` + "`" + `.]+).(zip|tar)") {
                return $matches[1]
            }
            else {
                return ""
            }
            break;
        }
        default {
            return ""
        }
    }
}

function Get-InstanceMetadataServiceTelemetry {
    $keys = @{ }

    try {
        # Write-Log "Querying instance metadata service..."
        # Note: 2019-04-30 is latest api available in all clouds
        $metadata = Invoke-RestMethod -Headers @{"Metadata" = "true" } -URI "http://169.254.169.254/metadata/instance?api-version=2019-04-30" -Method get
        # Write-Log ($metadata | ConvertTo-Json)

        $keys.Add("vm_size", $metadata.compute.vmSize)
    }
    catch {
        Write-Log "Error querying instance metadata service."
    }

    return $keys
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

function Retry-Command {
    Param(
        [Parameter(Mandatory = $true)][ValidateNotNullOrEmpty()][string]
        $Command,
        [Parameter(Mandatory = $true)][ValidateNotNullOrEmpty()][hashtable]
        $Args,
        [Parameter(Mandatory = $true)][ValidateNotNullOrEmpty()][int]
        $Retries,
        [Parameter(Mandatory = $true)][ValidateNotNullOrEmpty()][int]
        $RetryDelaySeconds
    )

    for ($i = 0; $i -lt $Retries; $i++) {
        try {
            return & $Command @Args
        }
        catch {
            Start-Sleep $RetryDelaySeconds
        }
    }
}

function Invoke-Executable {
    Param(
        [string]
        $Executable,
        [string[]]
        $ArgList,
        [int[]]
        $AllowedExitCodes = @(0),
        [int]
        $Retries = 1,
        [int]
        $RetryDelaySeconds = 1
    )

    for ($i = 0; $i -lt $Retries; $i++) {
        Write-Log "Running $Executable $ArgList ..."
        & $Executable $ArgList
        if ($LASTEXITCODE -notin $AllowedExitCodes) {
            Write-Log "$Executable returned unsuccessfully with exit code $LASTEXITCODE"
            Start-Sleep -Seconds $RetryDelaySeconds
            continue
        }
        else {
            Write-Log "$Executable returned successfully"
            return
        }
    }

    throw "Exhausted retries for $Executable $ArgList"
}

function Get-LogCollectionScripts {
    Write-Log "Getting various log collect scripts and depencencies"
    mkdir 'c:\k\debug'
    DownloadFileOverHttp -Url 'https://github.com/Azure/aks-engine/raw/master/scripts/collect-windows-logs.ps1' -DestinationPath 'c:\k\debug\collect-windows-logs.ps1'
    DownloadFileOverHttp -Url 'https://github.com/microsoft/SDN/raw/master/Kubernetes/windows/debug/collectlogs.ps1' -DestinationPath 'c:\k\debug\collectlogs.ps1'
    DownloadFileOverHttp -Url 'https://github.com/microsoft/SDN/raw/master/Kubernetes/windows/debug/dumpVfpPolicies.ps1' -DestinationPath 'c:\k\debug\dumpVfpPolicies.ps1'
    DownloadFileOverHttp -Url 'https://github.com/microsoft/SDN/raw/master/Kubernetes/windows/debug/portReservationTest.ps1' -DestinationPath 'c:\k\debug\portReservationTest.ps1'
    DownloadFileOverHttp -Url 'https://github.com/microsoft/SDN/raw/master/Kubernetes/windows/debug/starthnstrace.cmd' -DestinationPath 'c:\k\debug\starthnstrace.cmd'
    DownloadFileOverHttp -Url 'https://github.com/microsoft/SDN/raw/master/Kubernetes/windows/debug/startpacketcapture.cmd' -DestinationPath 'c:\k\debug\startpacketcapture.cmd'
    DownloadFileOverHttp -Url 'https://github.com/microsoft/SDN/raw/master/Kubernetes/windows/debug/stoppacketcapture.cmd' -DestinationPath 'c:\k\debug\stoppacketcapture.cmd'
    DownloadFileOverHttp -Url 'https://github.com/Microsoft/SDN/raw/master/Kubernetes/windows/debug/VFP.psm1' -DestinationPath 'c:\k\debug\VFP.psm1'
    DownloadFileOverHttp -Url 'https://github.com/microsoft/SDN/raw/master/Kubernetes/windows/helper.psm1' -DestinationPath 'c:\k\debug\helper.psm1'
    DownloadFileOverHttp -Url 'https://github.com/microsoft/SDN/raw/master/Kubernetes/windows/hns.psm1' -DestinationPath 'c:\k\debug\hns.psm1'
}

function Register-LogsCleanupScriptTask {
    Write-Log "Creating a scheduled task to run windowslogscleanup.ps1"

    $action = New-ScheduledTaskAction -Execute "powershell.exe" -Argument "-File ` + "`" + `"c:\k\windowslogscleanup.ps1` + "`" + `""
    $principal = New-ScheduledTaskPrincipal -UserId SYSTEM -LogonType ServiceAccount -RunLevel Highest
    $trigger = New-JobTrigger -Daily -At "00:00" -DaysInterval 1
    $definition = New-ScheduledTask -Action $action -Principal $principal -Trigger $trigger -Description "log-cleanup-task"
    Register-ScheduledTask -TaskName "log-cleanup-task" -InputObject $definition
}

function Register-NodeResetScriptTask {
    Write-Log "Creating a startup task to run windowsnodereset.ps1"

    $action = New-ScheduledTaskAction -Execute "powershell.exe" -Argument "-File ` + "`" + `"c:\k\windowsnodereset.ps1` + "`" + `""
    $principal = New-ScheduledTaskPrincipal -UserId SYSTEM -LogonType ServiceAccount -RunLevel Highest
    $trigger = New-JobTrigger -AtStartup -RandomDelay 00:00:05
    $definition = New-ScheduledTask -Action $action -Principal $principal -Trigger $trigger -Description "k8s-restart-job"
    Register-ScheduledTask -TaskName "k8s-restart-job" -InputObject $definition
}

# TODO ksubrmnn parameterize this fully
function Write-KubeClusterConfig {
    param(
        [Parameter(Mandatory = $true)][string]
        $MasterIP,
        [Parameter(Mandatory = $true)][string]
        $KubeDnsServiceIp
    )

    $Global:ClusterConfiguration = [PSCustomObject]@{ }

    $Global:ClusterConfiguration | Add-Member -MemberType NoteProperty -Name Cri -Value @{
        Name   = $global:ContainerRuntime;
        Images = @{
            # e.g. "mcr.microsoft.com/oss/kubernetes/pause:1.4.0"
            "Pause" = $global:WindowsPauseImageURL
        }
    }

    $Global:ClusterConfiguration | Add-Member -MemberType NoteProperty -Name Cni -Value @{
        Name   = $global:NetworkPlugin;
        Plugin = @{
            Name = "bridge";
        };
    }

    $Global:ClusterConfiguration | Add-Member -MemberType NoteProperty -Name Csi -Value @{
        EnableProxy = $global:EnableCsiProxy
    }

    $Global:ClusterConfiguration | Add-Member -MemberType NoteProperty -Name Kubernetes -Value @{
        Source       = @{
            Release = $global:KubeBinariesVersion;
        };
        ControlPlane = @{
            IpAddress    = $MasterIP;
            Username     = "azureuser"
            MasterSubnet = $global:MasterSubnet
        };
        Network      = @{
            ServiceCidr = $global:KubeServiceCIDR;
            ClusterCidr = $global:KubeClusterCIDR;
            DnsIp       = $KubeDnsServiceIp
        };
        Kubelet      = @{
            NodeLabels = $global:KubeletNodeLabels;
            ConfigArgs = $global:KubeletConfigArgs
        };
    }

    $Global:ClusterConfiguration | Add-Member -MemberType NoteProperty -Name Install -Value @{
        Destination = "c:\k";
    }

    $Global:ClusterConfiguration | ConvertTo-Json -Depth 10 | Out-File -FilePath $global:KubeClusterConfigPath
}

function Assert-FileExists {
    Param(
        [Parameter(Mandatory = $true, Position = 0)][string]
        $Filename
    )

    if (-Not (Test-Path $Filename)) {
        throw "$Filename does not exist"
    }
}

function Update-DefenderPreferences {
    Add-MpPreference -ExclusionProcess "c:\k\kubelet.exe"

    if ($global:EnableCsiProxy) {
        Add-MpPreference -ExclusionProcess "c:\k\csi-proxy-server.exe"
    }

    if ($global:ContainerRuntime -eq 'containerd') {
        Add-MpPreference -ExclusionProcess "c:\program files\containerd\containerd.exe"
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
    $TargetEnvironment,

    [string]
    $UserAssignedClientID
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
$global:ContainerdUrl = "{{WrapAsParameter "windowsContainerdURL"}}"
$global:ContainerdSdnPluginUrl = "{{WrapAsParameter "windowsSdnPluginURL"}}"

## Docker Version
$global:DockerVersion = "{{WrapAsParameter "windowsDockerVersion"}}"

## ContainerD Usage
$global:ContainerRuntime = "{{WrapAsParameter "containerRuntime"}}"

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
{{if IsKubernetesVersionGe "1.16.0"}}
$global:KubeletNodeLabels = "{{GetAgentKubernetesLabels . "',variables('labelResourceGroup'),'"}}"
{{else}}
$global:KubeletNodeLabels = "{{GetAgentKubernetesLabelsDeprecated . "',variables('labelResourceGroup'),'"}}"
{{end}}
$global:KubeletConfigArgs = @( {{GetKubeletConfigKeyValsPsh .KubernetesConfig }} )

$global:UseManagedIdentityExtension = "{{WrapAsVariable "useManagedIdentityExtension"}}"
$global:UseInstanceMetadata = "{{WrapAsVariable "useInstanceMetadata"}}"

$global:LoadBalancerSku = "{{WrapAsVariable "loadBalancerSku"}}"
$global:ExcludeMasterFromStandardLB = "{{WrapAsVariable "excludeMasterFromStandardLB"}}"


# Windows defaults, not changed by aks-engine
$global:CacheDir = "c:\akse-cache"
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
$global:IsDualStackEnabled = {{if IsIPv6DualStackFeatureEnabled}}$true{{else}}$false{{end}}

# Telemetry settings
$global:EnableTelemetry = "{{WrapAsVariable "enableTelemetry" }}";
$global:TelemetryKey = "{{WrapAsVariable "applicationInsightsKey" }}";

# CSI Proxy settings
$global:EnableCsiProxy = [System.Convert]::ToBoolean("{{WrapAsVariable "windowsEnableCSIProxy" }}");
$global:CsiProxyUrl = "{{WrapAsVariable "windowsCSIProxyURL" }}";

# Hosts Config Agent settings
$global:EnableHostsConfigAgent = [System.Convert]::ToBoolean("{{WrapAsVariable "enableHostsConfigAgent" }}");

$global:ProvisioningScriptsPackageUrl = "{{WrapAsVariable "windowsProvisioningScriptsPackageURL" }}";

# PauseImage
$global:WindowsPauseImageURL = "{{WrapAsVariable "windowsPauseImageURL" }}";
$global:AlwaysPullWindowsPauseImage = [System.Convert]::ToBoolean("{{WrapAsVariable "alwaysPullWindowsPauseImage" }}");

# Base64 representation of ZIP archive
$zippedFiles = "{{ GetKubernetesWindowsAgentFunctions }}"

# Extract ZIP from script
[io.file]::WriteAllBytes("scripts.zip", [System.Convert]::FromBase64String($zippedFiles))
Expand-Archive scripts.zip -DestinationPath "C:\\AzureData\\"

# Dot-source scripts with functions that are called in this script
. c:\AzureData\k8s\kuberneteswindowsfunctions.ps1
. c:\AzureData\k8s\windowsconfigfunc.ps1
. c:\AzureData\k8s\windowskubeletfunc.ps1
. c:\AzureData\k8s\windowscnifunc.ps1
. c:\AzureData\k8s\windowsazurecnifunc.ps1
. c:\AzureData\k8s\windowscsiproxyfunc.ps1
. c:\AzureData\k8s\windowsinstallopensshfunc.ps1
. c:\AzureData\k8s\windowscontainerdfunc.ps1
. c:\AzureData\k8s\windowshostsconfigagentfunc.ps1

$useContainerD = ($global:ContainerRuntime -eq "containerd")
$global:KubeClusterConfigPath = "c:\k\kubeclusterconfig.json"

try
{
    # Set to false for debugging.  This will output the start script to
    # c:\AzureData\CustomDataSetupScript.log, and then you can RDP
    # to the windows machine, and run the script manually to watch
    # the output.
    if ($true) {
        Write-Log "Provisioning $global:DockerServiceName... with IP $MasterIP"

        $global:globalTimer = [System.Diagnostics.Stopwatch]::StartNew()

        $configAppInsightsClientTimer = [System.Diagnostics.Stopwatch]::StartNew()
        # Get app insights binaries and set up app insights client
        mkdir c:\k\appinsights
        DownloadFileOverHttp -Url "https://globalcdn.nuget.org/packages/microsoft.applicationinsights.2.11.0.nupkg" -DestinationPath "c:\k\appinsights\microsoft.applicationinsights.2.11.0.zip"
        Expand-Archive -Path "c:\k\appinsights\microsoft.applicationinsights.2.11.0.zip" -DestinationPath "c:\k\appinsights"
        $appInsightsDll = "c:\k\appinsights\lib\net46\Microsoft.ApplicationInsights.dll"
        [Reflection.Assembly]::LoadFile($appInsightsDll)
        $conf = New-Object "Microsoft.ApplicationInsights.Extensibility.TelemetryConfiguration"
        $conf.DisableTelemetry = -not $global:enableTelemetry
        $conf.InstrumentationKey = $global:TelemetryKey
        $global:AppInsightsClient = New-Object "Microsoft.ApplicationInsights.TelemetryClient"($conf)

        $global:AppInsightsClient.Context.Properties["correlation_id"] = New-Guid
        $global:AppInsightsClient.Context.Properties["cri"] = $global:ContainerRuntime
        # TODO: Update once containerd versioning story is decided
        $global:AppInsightsClient.Context.Properties["cri_version"] = if ($global:ContainerRuntime -eq "docker") { $global:DockerVersion } else { "" }
        $global:AppInsightsClient.Context.Properties["k8s_version"] = $global:KubeBinariesVersion
        $global:AppInsightsClient.Context.Properties["lb_sku"] = $global:LoadBalancerSku
        $global:AppInsightsClient.Context.Properties["location"] = $Location
        $global:AppInsightsClient.Context.Properties["os_type"] = "windows"
        $global:AppInsightsClient.Context.Properties["os_version"] = Get-WindowsVersion
        $global:AppInsightsClient.Context.Properties["network_plugin"] = $global:NetworkPlugin
        $global:AppInsightsClient.Context.Properties["network_plugin_version"] = Get-CniVersion
        $global:AppInsightsClient.Context.Properties["network_mode"] = $global:NetworkMode
        $global:AppInsightsClient.Context.Properties["subscription_id"] = $global:SubscriptionId

        $vhdId = ""
        if (Test-Path "c:\vhd-id.txt") {
            $vhdId = Get-Content "c:\vhd-id.txt"
        }
        $global:AppInsightsClient.Context.Properties["vhd_id"] = $vhdId

        $imdsProperties = Get-InstanceMetadataServiceTelemetry
        foreach ($key in $imdsProperties.keys) {
            $global:AppInsightsClient.Context.Properties[$key] = $imdsProperties[$key]
        }

        $configAppInsightsClientTimer.Stop()
        $global:AppInsightsClient.TrackMetric("Config-AppInsightsClient", $configAppInsightsClientTimer.Elapsed.TotalSeconds)

        # Install OpenSSH if SSH enabled
        $sshEnabled = [System.Convert]::ToBoolean("{{ WindowsSSHEnabled }}")

        if ( $sshEnabled ) {
            Write-Log "Install OpenSSH"
            $installOpenSSHTimer = [System.Diagnostics.Stopwatch]::StartNew()
            Install-OpenSSH -SSHKeys $SSHKeys
            $installOpenSSHTimer.Stop()
            $global:AppInsightsClient.TrackMetric("Install-OpenSSH", $installOpenSSHTimer.Elapsed.TotalSeconds)
        }

        Write-Log "Apply telemetry data setting"
        Set-TelemetrySetting -WindowsTelemetryGUID $global:WindowsTelemetryGUID

        Write-Log "Resize os drive if possible"
        $resizeTimer = [System.Diagnostics.Stopwatch]::StartNew()
        Resize-OSDrive
        $resizeTimer.Stop()
        $global:AppInsightsClient.TrackMetric("Resize-OSDrive", $resizeTimer.Elapsed.TotalSeconds)

        Write-Log "Initialize data disks"
        Initialize-DataDisks

        Write-Log "Create required data directories as needed"
        Initialize-DataDirectories

        New-Item -ItemType Directory -Path "c:\k" -Force | Out-Null
        Get-ProvisioningScripts

        Write-KubeClusterConfig -MasterIP $MasterIP -KubeDnsServiceIp $KubeDnsServiceIp

        if ($useContainerD) {
            Write-Log "Installing ContainerD"
            $containerdTimer = [System.Diagnostics.Stopwatch]::StartNew()
            $cniBinPath = $global:AzureCNIBinDir
            $cniConfigPath = $global:AzureCNIConfDir
            if ($global:NetworkPlugin -eq "kubenet") {
                $cniBinPath = $global:CNIPath
                $cniConfigPath = $global:CNIConfigPath
            }
            Install-Containerd -ContainerdUrl $global:ContainerdUrl -CNIBinDir $cniBinPath -CNIConfDir $cniConfigPath
            $containerdTimer.Stop()
            $global:AppInsightsClient.TrackMetric("Install-ContainerD", $containerdTimer.Elapsed.TotalSeconds)
            # TODO: disable/uninstall Docker later
        } else {
            Write-Log "Install docker"
            $dockerTimer = [System.Diagnostics.Stopwatch]::StartNew()
            Install-Docker -DockerVersion $global:DockerVersion
            Set-DockerLogFileOptions
            $dockerTimer.Stop()
            $global:AppInsightsClient.TrackMetric("Install-Docker", $dockerTimer.Elapsed.TotalSeconds)
        }

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
            -UserAssignedClientID $UserAssignedClientID ` + "`" + `
            -UseInstanceMetadata $global:UseInstanceMetadata ` + "`" + `
            -LoadBalancerSku $global:LoadBalancerSku ` + "`" + `
            -ExcludeMasterFromStandardLB $global:ExcludeMasterFromStandardLB ` + "`" + `
            -TargetEnvironment $TargetEnvironment

        {{if IsCustomCloudProfile}}
        $azureStackConfigFile = [io.path]::Combine($global:KubeDir, "azurestackcloud.json")
        $envJSON = "{{ GetBase64EncodedEnvironmentJSON }}"
        [io.file]::WriteAllBytes($azureStackConfigFile, [System.Convert]::FromBase64String($envJSON))
        {{end}}

        Write-Log "Write ca root"
        Write-CACert -CACertificate $global:CACertificate ` + "`" + `
            -KubeDir $global:KubeDir

        if ($global:EnableCsiProxy) {
            New-CsiProxyService -CsiProxyPackageUrl $global:CsiProxyUrl -KubeDir $global:KubeDir
        }

        Write-Log "Write kube config"
        Write-KubeConfig -CACertificate $global:CACertificate ` + "`" + `
            -KubeDir $global:KubeDir ` + "`" + `
            -MasterFQDNPrefix $MasterFQDNPrefix ` + "`" + `
            -MasterIP $MasterIP ` + "`" + `
            -AgentKey $AgentKey ` + "`" + `
            -AgentCertificate $global:AgentCertificate

        if ($global:EnableHostsConfigAgent) {
             Write-Log "Starting hosts config agent"
             New-HostsConfigService
         }

        Write-Log "Create the Pause Container kubletwin/pause"
        $infraContainerTimer = [System.Diagnostics.Stopwatch]::StartNew()
        New-InfraContainer -KubeDir $global:KubeDir -ContainerRuntime $global:ContainerRuntime
        $infraContainerTimer.Stop()
        $global:AppInsightsClient.TrackMetric("New-InfraContainer", $infraContainerTimer.Elapsed.TotalSeconds)

        if (-not (Test-ContainerImageExists -Image "kubletwin/pause" -ContainerRuntime $global:ContainerRuntime)) {
            Write-Log "Could not find container with name kubletwin/pause"
            if ($useContainerD) {
                $o = ctr -n k8s.io image list
                Write-Log $o
            } else {
                $o = docker image list
                Write-Log $o
            }
            throw "kubletwin/pause container does not exist!"
        }

        Write-Log "Configuring networking with NetworkPlugin:$global:NetworkPlugin"

        # Configure network policy.
        Get-HnsPsm1 -HNSModule $global:HNSModule
        Import-Module $global:HNSModule

        if ($global:NetworkPlugin -eq "azure") {
            Write-Log "Installing Azure VNet plugins"
            Install-VnetPlugins -AzureCNIConfDir $global:AzureCNIConfDir ` + "`" + `
                -AzureCNIBinDir $global:AzureCNIBinDir ` + "`" + `
                -VNetCNIPluginsURL $global:VNetCNIPluginsURL

            Set-AzureCNIConfig -AzureCNIConfDir $global:AzureCNIConfDir ` + "`" + `
                -KubeDnsSearchPath $global:KubeDnsSearchPath ` + "`" + `
                -KubeClusterCIDR $global:KubeClusterCIDR ` + "`" + `
                -MasterSubnet $global:MasterSubnet ` + "`" + `
                -KubeServiceCIDR $global:KubeServiceCIDR ` + "`" + `
                -VNetCIDR $global:VNetCIDR ` + "`" + `
                {{- /* Azure Stack has discrete Azure CNI config requirements */}}
                -IsAzureStack {{if IsAzureStackCloud}}$true{{else}}$false{{end}} ` + "`" + `
                -IsDualStackEnabled $global:IsDualStackEnabled

            if ($TargetEnvironment -ieq "AzureStackCloud") {
                GenerateAzureStackCNIConfig ` + "`" + `
                    -TenantId $global:TenantId ` + "`" + `
                    -SubscriptionId $global:SubscriptionId ` + "`" + `
                    -ResourceGroup $global:ResourceGroup ` + "`" + `
                    -AADClientId $AADClientId ` + "`" + `
                    -KubeDir $global:KubeDir ` + "`" + `
                    -AADClientSecret $([System.Text.Encoding]::ASCII.GetString([System.Convert]::FromBase64String($AADClientSecret))) ` + "`" + `
                    -NetworkAPIVersion $NetworkAPIVersion ` + "`" + `
                    -AzureEnvironmentFilePath $([io.path]::Combine($global:KubeDir, "azurestackcloud.json")) ` + "`" + `
                    -IdentitySystem "{{ GetIdentitySystem }}"
            }
        }
        elseif ($global:NetworkPlugin -eq "kubenet") {
            Write-Log "Fetching additional files needed for kubenet"
            if ($useContainerD) {
                # TODO: CNI may need to move to c:\program files\containerd\cni\bin with ContainerD
                Install-SdnBridge -Url $global:ContainerdSdnPluginUrl -CNIPath $global:CNIPath
            } else {
                Update-WinCNI -CNIPath $global:CNIPath
            }
        }

        New-ExternalHnsNetwork -IsDualStackEnabled $global:IsDualStackEnabled

        Install-KubernetesServices ` + "`" + `
            -KubeDir $global:KubeDir

        Get-LogCollectionScripts

        Write-Log "Disable Internet Explorer compat mode and set homepage"
        Set-Explorer

        Write-Log "Adjust pagefile size"
        Adjust-PageFileSize

        Write-Log "Start preProvisioning script"
        PREPROVISION_EXTENSION

        Write-Log "Update service failure actions"
        Update-ServiceFailureActions -ContainerRuntime $global:ContainerRuntime

        Adjust-DynamicPortRange
        Register-LogsCleanupScriptTask
        Register-NodeResetScriptTask
        Update-DefenderPreferences

        if (Test-Path $CacheDir)
        {
            Write-Log "Removing aks-engine bits cache directory"
            Remove-Item $CacheDir -Recurse -Force
        }

        $global:globalTimer.Stop()
        $global:AppInsightsClient.TrackMetric("TotalDuration", $global:globalTimer.Elapsed.TotalSeconds)
        $global:AppInsightsClient.Flush()

        Write-Log "Setup Complete, reboot computer"
        Restart-Computer
    }
    else
    {
        # keep for debugging purposes
        Write-Log ".\CustomDataSetupScript.ps1 -MasterIP $MasterIP -KubeDnsServiceIp $KubeDnsServiceIp -MasterFQDNPrefix $MasterFQDNPrefix -Location $Location -AgentKey $AgentKey -AADClientId $AADClientId -AADClientSecret $AADClientSecret -NetworkAPIVersion $NetworkAPIVersion -TargetEnvironment $TargetEnvironment"
    }
}
catch
{
    $exceptionTelemtry = New-Object "Microsoft.ApplicationInsights.DataContracts.ExceptionTelemetry"
    $exceptionTelemtry.Exception = $_.Exception
    $global:AppInsightsClient.TrackException($exceptionTelemtry)
    $global:AppInsightsClient.Flush()

    Write-Error $_
    throw $_
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
  priorityClassName: system-node-critical
  hostNetwork: true
  containers:
    - name: cloud-controller-manager
      image: {{ContainerImage "cloud-controller-manager"}}
      imagePullPolicy: IfNotPresent
      command: [{{ContainerConfig "command"}}]
      args: [{{GetCloudControllerManagerArgs}}]
      resources:
        requests:
          cpu: 100m
          memory: 128Mi
        limits:
          cpu: 4
          memory: 2Gi
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
  labels:
    app: kube-addon-manager
spec:
  priorityClassName: system-node-critical
  hostNetwork: true
  containers:
  - name: kube-addon-manager
    image: {{ContainerImage "kube-addon-manager"}}
    imagePullPolicy: IfNotPresent
    env:
    - name: KUBECONFIG
      value: "/var/lib/kubelet/kubeconfig"
    - name: ADDON_PATH
      value: "/etc/kubernetes/addons/init"
    resources:
      requests:
        cpu: 5m
        memory: 50Mi
    volumeMounts:
    - name: addons
      mountPath: /etc/kubernetes/addons/init
      readOnly: true
    - name: msi
      mountPath: /var/lib/waagent/ManagedIdentity-Settings
      readOnly: true
    - name: var-lib-kubelet
      mountPath: /var/lib/kubelet
      readOnly: true
    - name: etc-kubernetes
      mountPath: /etc/kubernetes
      readOnly: true
  volumes:
  - name: addons
    hostPath:
      path: /etc/kubernetes/addons/init
  - name: msi
    hostPath:
      path: /var/lib/waagent/ManagedIdentity-Settings
  - name: var-lib-kubelet
    hostPath:
      path: /var/lib/kubelet
  - name: etc-kubernetes
    hostPath:
      path: /etc/kubernetes
#EOF
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
  priorityClassName: system-node-critical
  hostNetwork: true
  containers:
    - name: kube-apiserver
      image: {{ContainerImage "kube-apiserver"}}
      imagePullPolicy: IfNotPresent
      command: [{{ContainerConfig "command"}}]
      args: [{{GetAPIServerArgs}}]
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

var _k8sManifestsKubernetesmasterKubeControllerManagerYaml = []byte(`apiVersion: v1
kind: Pod
metadata:
  name: kube-controller-manager
  namespace: kube-system
  labels:
    tier: control-plane
    component: kube-controller-manager
spec:
  priorityClassName: system-node-critical
  hostNetwork: true
  containers:
    - name: kube-controller-manager
      image: {{ContainerImage "kube-controller-manager"}}
      imagePullPolicy: IfNotPresent
      command: [{{ContainerConfig "command"}}]
      args: [{{GetControllerManagerArgs}}]
{{- if IsCustomCloudProfile}}
      env:
      - name: AZURE_ENVIRONMENT_FILEPATH
        value: "/etc/kubernetes/azurestackcloud.json"
{{end}}
      volumeMounts:
        - name: etc-kubernetes
          mountPath: /etc/kubernetes
{{- if IsKubernetesVersionGe "1.17.0"}}
        - name: etc-ssl
          mountPath: /etc/ssl
          readOnly: true
{{end}}
        - name: var-lib-kubelet
          mountPath: /var/lib/kubelet
        - name: msi
          mountPath: /var/lib/waagent/ManagedIdentity-Settings
          readOnly: true
{{- if IsCustomCloudProfile}}
        <volumeMountssl>
{{end}}
  volumes:
    - name: etc-kubernetes
      hostPath:
        path: /etc/kubernetes
{{- if IsKubernetesVersionGe "1.17.0"}}
    - name: etc-ssl
      hostPath:
        path: /etc/ssl
{{end}}
    - name: var-lib-kubelet
      hostPath:
        path: /var/lib/kubelet
    - name: msi
      hostPath:
        path: /var/lib/waagent/ManagedIdentity-Settings
{{- if IsCustomCloudProfile}}
    <volumessl>
{{end}}
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
  priorityClassName: system-node-critical
  hostNetwork: true
  containers:
    - name: kube-scheduler
      image: {{ContainerImage "kube-scheduler"}}
      imagePullPolicy: IfNotPresent
      command: [{{ContainerConfig "command"}}]
      args: [{{GetSchedulerArgs}}]
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
        [Parameter(Mandatory=$true)][bool]
        $IsAzureStack,
        [Parameter(Mandatory=$true)][bool]
        $IsDualStackEnabled
    )
    # Fill in DNS information for kubernetes.
    if ($IsDualStackEnabled){
        $subnetToPass = $KubeClusterCIDR -split ","
        $exceptionAddresses = @($subnetToPass[0], $MasterSubnet, $VNetCIDR)
    }
    else {
        $exceptionAddresses = @($KubeClusterCIDR, $MasterSubnet, $VNetCIDR)
    }

    $fileName  = [Io.path]::Combine("$AzureCNIConfDir", "10-azure.conflist")
    $configJson = Get-Content $fileName | ConvertFrom-Json
    $configJson.plugins.dns.Nameservers[0] = $KubeDnsServiceIp
    $configJson.plugins.dns.Search[0] = $KubeDnsSearchPath

    $osBuildNumber = (get-wmiobject win32_operatingsystem).BuildNumber
    if ($osBuildNumber -le 17763){
        # In WS2019 and below rules in the exception list are generated by dropping the prefix lenght and removing duplicate rules.
        # If multiple execptions are specified with different ranges we should only include the broadest range for each address.
        # This issue has been addressed in 19h1+ builds

        $processedExceptions = GetBroadestRangesForEachAddress $exceptionAddresses
        Write-Host "Filtering CNI config exception list values to work around WS2019 issue processing rules. Original exception list: $exceptionAddresses, processed exception list: $processedExceptions"
        $configJson.plugins.AdditionalArgs[0].Value.ExceptionList = $processedExceptions
    }
    else {
        $configJson.plugins.AdditionalArgs[0].Value.ExceptionList = $exceptionAddresses
    }

    if ($IsDualStackEnabled){
        $configJson.plugins[0]|Add-Member -Name "ipv6Mode" -Value "ipv6nat" -MemberType NoteProperty
        $serviceCidr = $KubeServiceCIDR -split ","
        $configJson.plugins[0].AdditionalArgs[1].Value.DestinationPrefix = $serviceCidr[0]
        $valueObj = [PSCustomObject]@{
            Type = 'ROUTE'
            DestinationPrefix = $serviceCidr[1]
            NeedEncap = $True
        }

        $jsonContent = [PSCustomObject]@{
            Name = 'EndpointPolicy'
            Value = $valueObj
        }
        $configJson.plugins[0].AdditionalArgs += $jsonContent
    }
    else {
        $configJson.plugins[0].AdditionalArgs[1].Value.DestinationPrefix = $KubeServiceCIDR
    }

    if ($IsAzureStack) {
        Add-Member -InputObject $configJson.plugins[0].ipam -MemberType NoteProperty -Name "environment" -Value "mas"
    }

    $configJson | ConvertTo-Json -depth 20 | Out-File -encoding ASCII -filepath $fileName
}

function GetBroadestRangesForEachAddress{
    param([string[]] $values)

    # Create a map of range values to IP addresses
    $map = @{}

    foreach ($value in $Values) {
        if ($value -match '([0-9\.]+)\/([0-9]+)') {
            if (!$map.contains($matches[1])) {
                $map.Add($matches[1], @())
            }

            $map[$matches[1]] += [int]$matches[2]
        }
    }

    # For each IP address select the range with the lagest scope (smallest value)
    $returnValues = @()
    foreach ($ip in $map.Keys) {
        $range = $map[$ip] | Sort-Object | Select-Object -First 1

        $returnValues += $ip + "/" + $range
    }

    # prefix $returnValues with common to ensure single values get returned as an array otherwise invalid json may be generated
    return ,$returnValues
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
        [Parameter(Mandatory=$true)][ValidateNotNullOrEmpty()][string] $IdentitySystem,
        [Parameter(Mandatory=$true)][ValidateNotNullOrEmpty()][string] $KubeDir

    )

    $networkInterfacesFile = "$KubeDir\network-interfaces.json"
    $azureCNIConfigFile = "$KubeDir\interfaces.json"
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

function New-ExternalHnsNetwork
{
    param (
        [Parameter(Mandatory=$true)][bool]
        $IsDualStackEnabled
    )

    Write-Log "Creating new HNS network ` + "`" + `"ext` + "`" + `""
    $externalNetwork = "ext"
    $na = @(Get-NetAdapter -Physical)

    if ($na.Count -eq 0) {
        throw "Failed to find any physical network adapters"
    }

    # If there is more than one adapter, use the first adapter.
    $managementIP = (Get-NetIPAddress -ifIndex $na[0].ifIndex -AddressFamily IPv4).IPAddress
    $adapterName = $na[0].Name
    Write-Log "Using adapter $adapterName with IP address $managementIP"
    $mgmtIPAfterNetworkCreate

    $stopWatch = New-Object System.Diagnostics.Stopwatch
    $stopWatch.Start()

    # Fixme : use a smallest range possible, that will not collide with any pod space
    if ($IsDualStackEnabled) {
        New-HNSNetwork -Type $global:NetworkMode -AddressPrefix @("192.168.255.0/30","192:168:255::0/127") -Gateway @("192.168.255.1","192:168:255::1") -AdapterName $adapterName -Name $externalNetwork -Verbose
    }
    else {
        New-HNSNetwork -Type $global:NetworkMode -AddressPrefix "192.168.255.0/30" -Gateway "192.168.255.1" -AdapterName $adapterName -Name $externalNetwork -Verbose
    }
    # Wait for the switch to be created and the ip address to be assigned.
    for ($i = 0; $i -lt 60; $i++) {
        $mgmtIPAfterNetworkCreate = Get-NetIPAddress $managementIP -ErrorAction SilentlyContinue
        if ($mgmtIPAfterNetworkCreate) {
            break
        }
        Start-Sleep -Milliseconds 500
    }

    $stopWatch.Stop()
    if (-not $mgmtIPAfterNetworkCreate) {
        throw "Failed to find $managementIP after creating $externalNetwork network"
    }
    Write-Log "It took $($StopWatch.Elapsed.Seconds) seconds to create the $externalNetwork network."
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

var _k8sWindowsazurecnifuncTestsPs1 = []byte(`. $PSScriptRoot\windowsazurecnifunc.ps1

Describe 'GetBroadestRangesForEachAddress' {

    It "Values '<Values>' should return '<Expected>'" -TestCases @(
        @{ Values = @('10.240.0.0/12', '10.0.0.0/8'); Expected = @('10.0.0.0/8', '10.240.0.0/12')}
        @{ Values = @('10.0.0.0/8', '10.0.0.0/16'); Expected = @('10.0.0.0/8')}
        @{ Values = @('10.0.0.0/16', '10.240.0.0/12', '10.0.0.0/8' ); Expected = @('10.0.0.0/8', '10.240.0.0/12')}
        @{ Values = @(); Expected = @()}
        @{ Values = @('foobar'); Expected = @()}
    ){
        param ($Values, $Expected)

        $actual = GetBroadestRangesForEachAddress -values $Values
        $actual | Should -Be $Expected
    }
}`)

func k8sWindowsazurecnifuncTestsPs1Bytes() ([]byte, error) {
	return _k8sWindowsazurecnifuncTestsPs1, nil
}

func k8sWindowsazurecnifuncTestsPs1() (*asset, error) {
	bytes, err := k8sWindowsazurecnifuncTestsPs1Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/windowsazurecnifunc.tests.ps1", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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

function Install-SdnBridge
{
    Param(
        [Parameter(Mandatory=$true)][string]
        $Url,
        [Parameter(Mandatory=$true)][string]
        $CNIPath
    )

    $cnizip = [Io.path]::Combine($CNIPath, "cni.zip")
    DownloadFileOverHttp -Url $Url -DestinationPath $cnizip
    Expand-Archive -path $cnizip -DestinationPath $CNIPath
    del $cnizip
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

    # DOCKER_API_VERSION needs to be set for Docker versions older than 18.09.0 EE
    # due to https://github.com/kubernetes/kubernetes/issues/69996
    # this issue was fixed by https://github.com/kubernetes/kubernetes/issues/69996#issuecomment-438499024
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
        $installDocker = $true
        $dockerService = Get-Service | ? Name -like 'docker'
        if ($dockerService.Count -eq 0) {
            Write-Log "Docker is not installed. Install docker version($DockerVersion)."
        }
        else {
            $dockerServerVersion = docker version --format '{{.Server.Version}}'
            Write-Log "Docker service is installed with docker version($dockerServerVersion)."
            if ($dockerServerVersion -eq $DockerVersion) {
                $installDocker = $false
                Write-Log "Same version docker installed will skip installing docker version($dockerServerVersion)."
            }
            else {
                Write-Log "Same version docker is not installed. Will install docker version($DockerVersion)."
            }
        }

        if ($installDocker) {
            Find-Package -Name Docker -ProviderName DockerMsftProvider -RequiredVersion $DockerVersion -ErrorAction Stop
            Write-Log "Found version $DockerVersion. Installing..."
            Install-Package -Name Docker -ProviderName DockerMsftProvider -Update -Force -RequiredVersion $DockerVersion
            net start docker
            Write-Log "Installed version $DockerVersion"
        }
    } catch {
        Write-Log "Error while installing package: $_.Exception.Message"
        $currentDockerVersion = (Get-Package -Name Docker -ProviderName DockerMsftProvider).Version
        Write-Log "Not able to install docker version. Using default version $currentDockerVersion"
    }
}

function Set-DockerLogFileOptions {
    Write-Log "Updating log file options in docker config"
    $dockerConfigPath = "C:\ProgramData\docker\config\daemon.json"

    if (-not (Test-Path $dockerConfigPath)) {
        "{}" | Out-File $dockerConfigPath
    }

    $dockerConfig = Get-Content $dockerConfigPath | ConvertFrom-Json
    $dockerConfig | Add-Member -Name "log-driver" -Value "json-file" -MemberType NoteProperty
    $logOpts = @{ "max-size" = "50m"; "max-file" = "5" }
    $dockerConfig | Add-Member -Name "log-opts" -Value $logOpts -MemberType NoteProperty
    $dockerConfig = $dockerConfig | ConvertTo-Json -Depth 10

    Write-Log "New docker config:"
    Write-Log $dockerConfig

    # daemon.json MUST be encoded as UTF8-no-BOM!
    Remove-Item $dockerConfigPath
    $fileEncoding = New-Object System.Text.UTF8Encoding $false
    [IO.File]::WriteAllLInes($dockerConfigPath, $dockerConfig, $fileEncoding)

    Restart-Service docker
}

# Pagefile adjustments
function Adjust-PageFileSize()
{
    wmic pagefileset set InitialSize=8096,MaximumSize=8096
}

function Adjust-DynamicPortRange()
{
    # Kube-proxy reserves 63 ports per service which limits clusters with Windows nodes
    # to ~225 services if default port reservations are used.
    # https://docs.microsoft.com/en-us/virtualization/windowscontainers/kubernetes/common-problems#load-balancers-are-plumbed-inconsistently-across-the-cluster-nodes
    # Kube-proxy load balancing should be set to DSR mode when it releases with future versions of the OS

    Invoke-Executable -Executable "netsh.exe" -ArgList @("int", "ipv4", "set", "dynamicportrange", "tcp", "16385", "49151")
}

# TODO: should this be in this PR?
# Service start actions. These should be split up later and included in each install step
function Update-ServiceFailureActions
{
    Param(
        [Parameter(Mandatory = $true)][string]
        $ContainerRuntime
    )
    sc.exe failure "kubelet" actions= restart/60000/restart/60000/restart/60000 reset= 900
    sc.exe failure "kubeproxy" actions= restart/60000/restart/60000/restart/60000 reset= 900
    sc.exe failure $ContainerRuntime actions= restart/60000/restart/60000/restart/60000 reset= 900
}

function Add-SystemPathEntry
{
    Param(
        [Parameter(Mandatory = $true)][string]
        $Directory
    )
    # update the path variable if it doesn't have the needed paths
    $path = [Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::Machine)
    $updated = $false
    if(-not ($path -match $Directory.Replace("\","\\")+"(;|$)"))
    {
        $path += ";"+$Directory
        $updated = $true
    }
    if($updated)
    {
        Write-Output "Updating path, added $Directory"
        [Environment]::SetEnvironmentVariable("Path", $path, [EnvironmentVariableTarget]::Machine)
        $env:Path = [System.Environment]::GetEnvironmentVariable("Path","Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path","User")
    }
}`)

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

var _k8sWindowscontainerdfuncPs1 = []byte(`# this is $global to persist across all functions since this is dot-sourced
$global:ContainerdInstallLocation = "$Env:ProgramFiles\containerd"

function RegisterContainerDService {
  Assert-FileExists (Join-Path $global:ContainerdInstallLocation containerd.exe)

  Write-Host "Registering containerd as a service"
  $cdbinary = Join-Path $global:ContainerdInstallLocation containerd.exe
  $svc = Get-Service -Name containerd -ErrorAction SilentlyContinue
  if ($null -ne $svc) {
    & $cdbinary --unregister-service
  }
  & $cdbinary --register-service
  $svc = Get-Service -Name "containerd" -ErrorAction SilentlyContinue
  if ($null -eq $svc) {
    throw "containerd.exe did not get installed as a service correctly."
  }
  $svc | Start-Service
  if ($svc.Status -ne "Running") {
    throw "containerd service is not running"
  }
}


function Install-Containerd {
  Param(
    [Parameter(Mandatory = $true)][string]
    $ContainerdUrl,
    [Parameter(Mandatory = $true)][string]
    $CNIBinDir,
    [Parameter(Mandatory = $true)][string]
    $CNIConfDir
  )

  $svc = Get-Service -Name containerd -ErrorAction SilentlyContinue
  if ($null -ne $svc) {
    Write-Log "Stoping containerd service"
    $svc | Stop-Service
  }

  # TODO: check if containerd is already installed and is the same version before this.
  $zipfile = [Io.path]::Combine($ENV:TEMP, "containerd.zip")
  DownloadFileOverHttp -Url $ContainerdUrl -DestinationPath $zipfile
  Expand-Archive -path $zipfile -DestinationPath $global:ContainerdInstallLocation -Force
  del $zipfile

  Add-SystemPathEntry $global:ContainerdInstallLocation

  # TODO: remove if the node comes up without this code
  # $configDir = [Io.Path]::Combine($ENV:ProgramData, "containerd")
  # if (-Not (Test-Path $configDir)) {
  #     mkdir $configDir
  # }

  # TODO: call containerd.exe dump config, then modify instead of starting with hardcoded
  $configFile = [Io.Path]::Combine($global:ContainerdInstallLocation, "config.toml")

  $clusterConfig = ConvertFrom-Json ((Get-Content $global:KubeClusterConfigPath -ErrorAction Stop) | Out-String)
  $pauseImage = $clusterConfig.Cri.Images.Pause

  @"
version = 2
root = "C:\\ProgramData\\containerd\\root"
state = "C:\\ProgramData\\containerd\\state"
plugin_dir = ""
disabled_plugins = []
required_plugins = []
oom_score = 0

[grpc]
  address = "\\\\.\\pipe\\containerd-containerd"
  tcp_address = ""
  tcp_tls_cert = ""
  tcp_tls_key = ""
  uid = 0
  gid = 0
  max_recv_message_size = 16777216
  max_send_message_size = 16777216

[ttrpc]
  address = ""
  uid = 0
  gid = 0

[debug]
  address = ""
  uid = 0
  gid = 0
  level = ""

[metrics]
  address = ""
  grpc_histogram = false

[cgroup]
  path = ""

[timeouts]
  "io.containerd.timeout.shim.cleanup" = "5s"
  "io.containerd.timeout.shim.load" = "5s"
  "io.containerd.timeout.shim.shutdown" = "3s"
  "io.containerd.timeout.task.state" = "2s"

[plugins]
  [plugins."io.containerd.gc.v1.scheduler"]
    pause_threshold = 0.02
    deletion_threshold = 0
    mutation_threshold = 100
    schedule_delay = "0s"
    startup_delay = "100ms"
  [plugins."io.containerd.grpc.v1.cri"]
    disable_tcp_service = true
    stream_server_address = "127.0.0.1"
    stream_server_port = "0"
    stream_idle_timeout = "4h0m0s"
    enable_selinux = false
    sandbox_image = "$pauseImage"
    stats_collect_period = 10
    systemd_cgroup = false
    enable_tls_streaming = false
    max_container_log_line_size = 16384
    disable_cgroup = false
    disable_apparmor = false
    restrict_oom_score_adj = false
    max_concurrent_downloads = 3
    disable_proc_mount = false
    [plugins."io.containerd.grpc.v1.cri".containerd]
      snapshotter = "windows"
      default_runtime_name = "runhcs-wcow-process"
      no_pivot = false
      [plugins."io.containerd.grpc.v1.cri".containerd.default_runtime]
        runtime_type = ""
        runtime_engine = ""
        runtime_root = ""
        privileged_without_host_devices = false
      [plugins."io.containerd.grpc.v1.cri".containerd.untrusted_workload_runtime]
        runtime_type = ""
        runtime_engine = ""
        runtime_root = ""
        privileged_without_host_devices = false
      [plugins."io.containerd.grpc.v1.cri".containerd.runtimes]
        [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runhcs-wcow-process]
          runtime_type = "io.containerd.runhcs.v1"
          runtime_engine = ""
          runtime_root = ""
          privileged_without_host_devices = false
    [plugins."io.containerd.grpc.v1.cri".cni]
      bin_dir = "$(($CNIBinDir).Replace("\","//"))"
      conf_dir = "$(($CNIConfDir).Replace("\","//"))"
      max_conf_num = 1
      conf_template = ""
    [plugins."io.containerd.grpc.v1.cri".registry]
      [plugins."io.containerd.grpc.v1.cri".registry.mirrors]
        [plugins."io.containerd.grpc.v1.cri".registry.mirrors."docker.io"]
          endpoint = ["https://registry-1.docker.io"]
    [plugins."io.containerd.grpc.v1.cri".x509_key_pair_streaming]
      tls_cert_file = ""
      tls_key_file = ""
  [plugins."io.containerd.metadata.v1.bolt"]
    content_sharing_policy = "shared"
  [plugins."io.containerd.runtime.v2.task"]
    platforms = ["windows/amd64", "linux/amd64"]
  [plugins."io.containerd.service.v1.diff-service"]
    default = ["windows", "windows-lcow"]
"@ | Out-File -Encoding ascii $configFile

  RegisterContainerDService
}`)

func k8sWindowscontainerdfuncPs1Bytes() ([]byte, error) {
	return _k8sWindowscontainerdfuncPs1, nil
}

func k8sWindowscontainerdfuncPs1() (*asset, error) {
	bytes, err := k8sWindowscontainerdfuncPs1Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/windowscontainerdfunc.ps1", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sWindowscsiproxyfuncPs1 = []byte(`function New-CsiProxyService {
    Param(
        [Parameter(Mandatory = $true)][string]
        $CsiProxyPackageUrl,
        [Parameter(Mandatory = $true)][string]
        $KubeDir
    )

    $tempdir = New-TemporaryDirectory
    $binaryPackage = "$tempdir\csiproxy.tar"

    DownloadFileOverHttp -Url $CsiProxyPackageUrl -DestinationPath $binaryPackage

    tar -xzf $binaryPackage -C $tempdir
    cp "$tempdir\bin\csi-proxy.exe" "$KubeDir\csi-proxy.exe"

    del $tempdir -Recurse

    & "$KubeDir\nssm.exe" install csi-proxy "$KubeDir\csi-proxy.exe" | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy AppDirectory "$KubeDir" | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy AppRestartDelay 5000 | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy Description csi-proxy | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy Start SERVICE_DEMAND_START | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy ObjectName LocalSystem | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy Type SERVICE_WIN32_OWN_PROCESS | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy AppThrottle 1500 | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy AppStdout "$KubeDir\csi-proxy.log" | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy AppStderr "$KubeDir\csi-proxy.err.log" | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy AppStdoutCreationDisposition 4 | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy AppStderrCreationDisposition 4 | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy AppRotateFiles 1 | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy AppRotateOnline 1 | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy AppRotateSeconds 86400 | RemoveNulls
    & "$KubeDir\nssm.exe" set csi-proxy AppRotateBytes 10485760 | RemoveNulls
}`)

func k8sWindowscsiproxyfuncPs1Bytes() ([]byte, error) {
	return _k8sWindowscsiproxyfuncPs1, nil
}

func k8sWindowscsiproxyfuncPs1() (*asset, error) {
	bytes, err := k8sWindowscsiproxyfuncPs1Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/windowscsiproxyfunc.ps1", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sWindowshostsconfigagentfuncPs1 = []byte(`function New-HostsConfigService {
    $HostsConfigParameters = [io.path]::Combine($KubeDir, "hostsconfigagent.ps1")

    & "$KubeDir\nssm.exe" install hosts-config-agent C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe | RemoveNulls
    & "$KubeDir\nssm.exe" set hosts-config-agent AppDirectory "$KubeDir" | RemoveNulls
    & "$KubeDir\nssm.exe" set hosts-config-agent AppParameters $HostsConfigParameters | RemoveNulls
    & "$KubeDir\nssm.exe" set hosts-config-agent AppRestartDelay 5000 | RemoveNulls
    & "$KubeDir\nssm.exe" set hosts-config-agent Description hosts-config-agent | RemoveNulls
    & "$KubeDir\nssm.exe" set hosts-config-agent Start SERVICE_DEMAND_START | RemoveNulls
    & "$KubeDir\nssm.exe" set hosts-config-agent ObjectName LocalSystem | RemoveNulls
    & "$KubeDir\nssm.exe" set hosts-config-agent Type SERVICE_WIN32_OWN_PROCESS | RemoveNulls
    & "$KubeDir\nssm.exe" set hosts-config-agent AppThrottle 1500 | RemoveNulls
    & "$KubeDir\nssm.exe" set hosts-config-agent AppStdout "$KubeDir\hosts-config-agent.log" | RemoveNulls
    & "$KubeDir\nssm.exe" set hosts-config-agent AppStderr "$KubeDir\hosts-config-agent.err.log" | RemoveNulls
    & "$KubeDir\nssm.exe" set hosts-config-agent AppStdoutCreationDisposition 4 | RemoveNulls
    & "$KubeDir\nssm.exe" set hosts-config-agent AppStderrCreationDisposition 4 | RemoveNulls
    & "$KubeDir\nssm.exe" set hosts-config-agent AppRotateFiles 1 | RemoveNulls
    & "$KubeDir\nssm.exe" set hosts-config-agent AppRotateOnline 1 | RemoveNulls
    & "$KubeDir\nssm.exe" set hosts-config-agent AppRotateSeconds 86400 | RemoveNulls
    & "$KubeDir\nssm.exe" set hosts-config-agent AppRotateBytes 10485760 | RemoveNulls
}`)

func k8sWindowshostsconfigagentfuncPs1Bytes() ([]byte, error) {
	return _k8sWindowshostsconfigagentfuncPs1, nil
}

func k8sWindowshostsconfigagentfuncPs1() (*asset, error) {
	bytes, err := k8sWindowshostsconfigagentfuncPs1Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/windowshostsconfigagentfunc.ps1", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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

    $sshdService = Get-Service | ? Name -like 'sshd'
    if ($sshdService.Count -eq 0)
    {
        Write-Log "Installing OpenSSH"
        $isAvailable = Get-WindowsCapability -Online | ? Name -like 'OpenSSH*'

        if (!$isAvailable) {
            throw "OpenSSH is not available on this machine"
        }

        Add-WindowsCapability -Online -Name OpenSSH.Server~~~~0.0.1.0
    }
    else
    {
        Write-Log "OpenSSH Server service detected - skipping online install..."
    }

    Start-Service sshd

    if (!(Test-Path "$adminpath")) {
        Write-Log "Created new file and text content added"
        New-Item -path $adminpath -name $adminfile -type "file" -value ""
    }

    Write-Log "$adminpath found."
    Write-Log "Adding keys to: $adminpath\$adminfile ..."
    $SSHKeys | foreach-object {
        Add-Content $adminpath\$adminfile $_
    }

    Write-Log "Setting required permissions..."
    icacls $adminpath\$adminfile /remove "NT AUTHORITY\Authenticated Users"
    icacls $adminpath\$adminfile /inheritance:r
    icacls $adminpath\$adminfile /grant SYSTEM:` + "`" + `(F` + "`" + `)
    icacls $adminpath\$adminfile /grant BUILTIN\Administrators:` + "`" + `(F` + "`" + `)

    Write-Log "Restarting sshd service..."
    Restart-Service sshd
    # OPTIONAL but recommended:
    Set-Service -Name sshd -StartupType 'Automatic'

    # Confirm the Firewall rule is configured. It should be created automatically by setup. 
    $firewall = Get-NetFirewallRule -Name *ssh*

    if (!$firewall) {
        throw "OpenSSH is firewall is not configured properly"
    }
    Write-Log "OpenSSH installed and configured successfully"
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
        $TargetEnvironment,
        [Parameter(Mandatory = $false)][bool]
        $UseContainerD = $false
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
    "userAssignedIdentityID": "$UserAssignedClientID",
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
Test-ContainerImageExists {
    Param(
        [Parameter(Mandatory = $true)][string]
        $Image,
        [Parameter(Mandatory = $false)][string]
        $Tag,
        [Parameter(Mandatory = $false)][string]
        $ContainerRuntime = "docker"
    )

    $target = $Image
    if ($Tag) {
        $target += ":$Tag"
    }

    if ($ContainerRuntime -eq "docker") {
        $images = docker image list $target --format "{{json .}}"
        return $images.Count -gt 0
    }
    else {
        return ( (ctr.exe -n k8s.io images list) | Select-String $target) -ne $Null
    }
}

function
Build-PauseContainer {
    Param(
        [Parameter(Mandatory = $true)][string]
        $WindowsBase,
        $DestinationTag,
        [Parameter(Mandatory = $false)][string]
        $ContainerRuntime = "docker"
    )
    # Future work: This needs to build wincat - see https://github.com/Azure/aks-engine/issues/1461
    # Otherwise, delete this code and require a prebuilt pause image (or override with one from an Azure Container Registry instance)
    # ContainerD can't build, so doing the builds outside of node deployment is probably the right long-term solution.
    "FROM $($WindowsBase)" | Out-File -encoding ascii -FilePath Dockerfile
    "CMD cmd /c ping -t localhost" | Out-File -encoding ascii -FilePath Dockerfile -Append
    if ($ContainerRuntime -eq "docker") {
        Invoke-Executable -Executable "docker" -ArgList @("build", "-t", "$DestinationTag", ".")
    }
    else {
        throw "Cannot build pause container without Docker"
    }
}

function
New-InfraContainer {
    Param(
        [Parameter(Mandatory = $true)][string]
        $KubeDir,
        $DestinationTag = "kubletwin/pause",
        [Parameter(Mandatory = $false)][string]
        $ContainerRuntime = "docker"
    )
    cd $KubeDir
    $windowsVersion = (Get-ItemProperty "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion").ReleaseId

    # Reference for these tags: curl -L https://mcr.microsoft.com/v2/k8s/core/pause/tags/list
    # Then docker run --rm mplatform/manifest-tool inspect mcr.microsoft.com/k8s/core/pause:<tag>

    $clusterConfig = ConvertFrom-Json ((Get-Content $global:KubeClusterConfigPath -ErrorAction Stop) | Out-String)
    $defaultPauseImage = $clusterConfig.Cri.Images.Pause

    $pauseImageVersions = @("1809", "1903", "1909", "2004")

    if ($pauseImageVersions -icontains $windowsVersion) {
        if ($ContainerRuntime -eq "docker") {
            if (-not (Test-ContainerImageExists -Image $defaultPauseImage -ContainerRuntime $ContainerRuntime) -or $global:AlwaysPullWindowsPauseImage) {
                Invoke-Executable -Executable "docker" -ArgList @("pull", "$defaultPauseImage") -Retries 5 -RetryDelaySeconds 30
            }
            Invoke-Executable -Executable "docker" -ArgList @("tag", "$defaultPauseImage", "$DestinationTag")
        }
        else {
            # containerd
            if (-not (Test-ContainerImageExists -Image $defaultPauseImage -ContainerRuntime $ContainerRuntime) -or $global:AlwaysPullWindowsPauseImage) {
                Invoke-Executable -Executable "ctr" -ArgList @("-n", "k8s.io", "image", "pull", "$defaultPauseImage") -Retries 5 -RetryDelaySeconds 30
            }
            Invoke-Executable -Executable "ctr" -ArgList @("-n", "k8s.io", "image", "tag", "$defaultPauseImage", "$DestinationTag")
        }
    }
    else {
        Build-PauseContainer -WindowsBase "mcr.microsoft.com/nanoserver-insider" -DestinationTag $DestinationTag -ContainerRuntime $ContainerRuntime
    }
}

function
Test-ContainerImageExists {
    Param(
        [Parameter(Mandatory = $true)][string]
        $Image,
        [Parameter(Mandatory = $false)][string]
        $Tag,
        [Parameter(Mandatory = $false)][string]
        $ContainerRuntime = "docker"
    )

    $target = $Image
    if ($Tag) {
        $target += ":$Tag"
    }

    if ($ContainerRuntime -eq "docker") {
        $images = docker image list $target --format "{{json .}}"
        return $images.Count -gt 0
    }
    else {
        return ( (ctr.exe -n k8s.io images list) | Select-String $target) -ne $Null
    }
}

# TODO: Deprecate this and replace with methods that get individual components instead of zip containing everything
# This expects the ZIP file created by Azure Pipelines.
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

    $kubeletDependOnServices = "docker"
    if ($global:EnableCsiProxy) {
        $kubeletDependOnServices += " csi-proxy"
    }
    if ($global:EnableHostsConfigAgent) {
        $kubeletDependOnServices += " hosts-config-agent"
    }

    # setup kubelet
    & "$KubeDir\nssm.exe" install Kubelet C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppDirectory $KubeDir | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppParameters $KubeletStartFile | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet DisplayName Kubelet | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppRestartDelay 5000 | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet Description Kubelet | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet Start SERVICE_DEMAND_START | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet ObjectName LocalSystem | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet Type SERVICE_WIN32_OWN_PROCESS | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppThrottle 1500 | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppStdout C:\k\kubelet.log | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppStderr C:\k\kubelet.err.log | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppStdoutCreationDisposition 4 | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppStderrCreationDisposition 4 | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppRotateFiles 1 | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppRotateOnline 1 | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppRotateSeconds 86400 | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubelet AppRotateBytes 10485760 | RemoveNulls
    # Do not use & when calling DependOnService since 'docker csi-proxy'
    # is parsed as a single string instead of two separate strings
    Invoke-Expression "$KubeDir\nssm.exe set Kubelet DependOnService $kubeletDependOnServices | RemoveNulls"

    # setup kubeproxy
    & "$KubeDir\nssm.exe" install Kubeproxy C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy AppDirectory $KubeDir | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy AppParameters $KubeProxyStartFile | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy DisplayName Kubeproxy | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy DependOnService Kubelet | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy Description Kubeproxy | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy Start SERVICE_DEMAND_START | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy ObjectName LocalSystem | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy Type SERVICE_WIN32_OWN_PROCESS | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy AppThrottle 1500 | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy AppStdout C:\k\kubeproxy.log | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy AppStderr C:\k\kubeproxy.err.log | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy AppRotateFiles 1 | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy AppRotateOnline 1 | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy AppRotateSeconds 86400 | RemoveNulls
    & "$KubeDir\nssm.exe" set Kubeproxy AppRotateBytes 10485760 | RemoveNulls
}

# Renamed from Write-KubernetesStartFiles
function
Install-KubernetesServices {
    param(
        [Parameter(Mandatory = $true)][string]
        $KubeDir
    )

    # TODO ksbrmnn fix callers to this function

    $KubeletStartFile = [io.path]::Combine($KubeDir, "kubeletstart.ps1")
    $KubeProxyStartFile = [io.path]::Combine($KubeDir, "kubeproxystart.ps1")

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
{{if HasVMASAgentPool}}
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
        "description": "The download url for kubernetes windows binaries package"
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
    "windowsContainerdURL": {
      "metadata": {
        "description": "TODO: containerd - these binaries are not available yet"
      },
      "type": "string"
    },
    "windowsSdnPluginURL": {
      "metadata": {
        "description": "TODO: containerd - these binaries are not available yet"
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
	"agentoutputs.t":                                                     agentoutputsT,
	"agentparams.t":                                                      agentparamsT,
	"dcos/bstrap/bootstrapcustomdata.yml":                                dcosBstrapBootstrapcustomdataYml,
	"dcos/bstrap/bootstrapparams.t":                                      dcosBstrapBootstrapparamsT,
	"dcos/bstrap/bootstrapprovision.sh":                                  dcosBstrapBootstrapprovisionSh,
	"dcos/bstrap/bootstrapresources.t":                                   dcosBstrapBootstrapresourcesT,
	"dcos/bstrap/bootstrapvars.t":                                        dcosBstrapBootstrapvarsT,
	"dcos/bstrap/dcos1.11.0.customdata.t":                                dcosBstrapDcos1110CustomdataT,
	"dcos/bstrap/dcos1.11.2.customdata.t":                                dcosBstrapDcos1112CustomdataT,
	"dcos/bstrap/dcosbase.t":                                             dcosBstrapDcosbaseT,
	"dcos/bstrap/dcosmasterresources.t":                                  dcosBstrapDcosmasterresourcesT,
	"dcos/bstrap/dcosmastervars.t":                                       dcosBstrapDcosmastervarsT,
	"dcos/bstrap/dcosprovision.sh":                                       dcosBstrapDcosprovisionSh,
	"dcos/dcosWindowsAgentResourcesVmas.t":                               dcosDcoswindowsagentresourcesvmasT,
	"dcos/dcosWindowsAgentResourcesVmss.t":                               dcosDcoswindowsagentresourcesvmssT,
	"dcos/dcosWindowsProvision.ps1":                                      dcosDcoswindowsprovisionPs1,
	"dcos/dcosagentresourcesvmas.t":                                      dcosDcosagentresourcesvmasT,
	"dcos/dcosagentresourcesvmss.t":                                      dcosDcosagentresourcesvmssT,
	"dcos/dcosagentvars.t":                                               dcosDcosagentvarsT,
	"dcos/dcosbase.t":                                                    dcosDcosbaseT,
	"dcos/dcoscustomdata110.t":                                           dcosDcoscustomdata110T,
	"dcos/dcoscustomdata184.t":                                           dcosDcoscustomdata184T,
	"dcos/dcoscustomdata187.t":                                           dcosDcoscustomdata187T,
	"dcos/dcoscustomdata188.t":                                           dcosDcoscustomdata188T,
	"dcos/dcoscustomdata190.t":                                           dcosDcoscustomdata190T,
	"dcos/dcoscustomdata198.t":                                           dcosDcoscustomdata198T,
	"dcos/dcosmasterresources.t":                                         dcosDcosmasterresourcesT,
	"dcos/dcosmastervars.t":                                              dcosDcosmastervarsT,
	"dcos/dcosparams.t":                                                  dcosDcosparamsT,
	"dcos/dcosprovision.sh":                                              dcosDcosprovisionSh,
	"dcos/dcosprovisionsource.sh":                                        dcosDcosprovisionsourceSh,
	"iaasoutputs.t":                                                      iaasoutputsT,
	"k8s/addons/1.15/calico.yaml":                                        k8sAddons115CalicoYaml,
	"k8s/addons/aad-default-admin-group-rbac.yaml":                       k8sAddonsAadDefaultAdminGroupRbacYaml,
	"k8s/addons/aad-pod-identity.yaml":                                   k8sAddonsAadPodIdentityYaml,
	"k8s/addons/aci-connector.yaml":                                      k8sAddonsAciConnectorYaml,
	"k8s/addons/antrea.yaml":                                             k8sAddonsAntreaYaml,
	"k8s/addons/arc-onboarding.yaml":                                     k8sAddonsArcOnboardingYaml,
	"k8s/addons/audit-policy.yaml":                                       k8sAddonsAuditPolicyYaml,
	"k8s/addons/azure-cloud-provider.yaml":                               k8sAddonsAzureCloudProviderYaml,
	"k8s/addons/azure-cni-networkmonitor.yaml":                           k8sAddonsAzureCniNetworkmonitorYaml,
	"k8s/addons/azure-network-policy.yaml":                               k8sAddonsAzureNetworkPolicyYaml,
	"k8s/addons/azure-policy-deployment.yaml":                            k8sAddonsAzurePolicyDeploymentYaml,
	"k8s/addons/azuredisk-csi-driver-deployment.yaml":                    k8sAddonsAzurediskCsiDriverDeploymentYaml,
	"k8s/addons/azurefile-csi-driver-deployment.yaml":                    k8sAddonsAzurefileCsiDriverDeploymentYaml,
	"k8s/addons/blobfuse-flexvolume.yaml":                                k8sAddonsBlobfuseFlexvolumeYaml,
	"k8s/addons/calico.yaml":                                             k8sAddonsCalicoYaml,
	"k8s/addons/cilium.yaml":                                             k8sAddonsCiliumYaml,
	"k8s/addons/cloud-node-manager.yaml":                                 k8sAddonsCloudNodeManagerYaml,
	"k8s/addons/cluster-autoscaler.yaml":                                 k8sAddonsClusterAutoscalerYaml,
	"k8s/addons/container-monitoring.yaml":                               k8sAddonsContainerMonitoringYaml,
	"k8s/addons/coredns.yaml":                                            k8sAddonsCorednsYaml,
	"k8s/addons/flannel.yaml":                                            k8sAddonsFlannelYaml,
	"k8s/addons/ip-masq-agent.yaml":                                      k8sAddonsIpMasqAgentYaml,
	"k8s/addons/keyvault-flexvolume.yaml":                                k8sAddonsKeyvaultFlexvolumeYaml,
	"k8s/addons/kube-dns.yaml":                                           k8sAddonsKubeDnsYaml,
	"k8s/addons/kube-proxy.yaml":                                         k8sAddonsKubeProxyYaml,
	"k8s/addons/kube-rescheduler.yaml":                                   k8sAddonsKubeReschedulerYaml,
	"k8s/addons/kubernetes-dashboard.yaml":                               k8sAddonsKubernetesDashboardYaml,
	"k8s/addons/metrics-server.yaml":                                     k8sAddonsMetricsServerYaml,
	"k8s/addons/node-problem-detector.yaml":                              k8sAddonsNodeProblemDetectorYaml,
	"k8s/addons/nvidia-device-plugin.yaml":                               k8sAddonsNvidiaDevicePluginYaml,
	"k8s/addons/pod-security-policy.yaml":                                k8sAddonsPodSecurityPolicyYaml,
	"k8s/addons/scheduled-maintenance-deployment.yaml":                   k8sAddonsScheduledMaintenanceDeploymentYaml,
	"k8s/addons/secrets-store-csi-driver.yaml":                           k8sAddonsSecretsStoreCsiDriverYaml,
	"k8s/addons/smb-flexvolume.yaml":                                     k8sAddonsSmbFlexvolumeYaml,
	"k8s/addons/tiller.yaml":                                             k8sAddonsTillerYaml,
	"k8s/armparameters.t":                                                k8sArmparametersT,
	"k8s/cloud-init/artifacts/apt-preferences":                           k8sCloudInitArtifactsAptPreferences,
	"k8s/cloud-init/artifacts/auditd-rules":                              k8sCloudInitArtifactsAuditdRules,
	"k8s/cloud-init/artifacts/cis.sh":                                    k8sCloudInitArtifactsCisSh,
	"k8s/cloud-init/artifacts/cse_config.sh":                             k8sCloudInitArtifactsCse_configSh,
	"k8s/cloud-init/artifacts/cse_customcloud.sh":                        k8sCloudInitArtifactsCse_customcloudSh,
	"k8s/cloud-init/artifacts/cse_helpers.sh":                            k8sCloudInitArtifactsCse_helpersSh,
	"k8s/cloud-init/artifacts/cse_install.sh":                            k8sCloudInitArtifactsCse_installSh,
	"k8s/cloud-init/artifacts/cse_main.sh":                               k8sCloudInitArtifactsCse_mainSh,
	"k8s/cloud-init/artifacts/default-grub":                              k8sCloudInitArtifactsDefaultGrub,
	"k8s/cloud-init/artifacts/dhcpv6.service":                            k8sCloudInitArtifactsDhcpv6Service,
	"k8s/cloud-init/artifacts/docker-monitor.service":                    k8sCloudInitArtifactsDockerMonitorService,
	"k8s/cloud-init/artifacts/docker-monitor.timer":                      k8sCloudInitArtifactsDockerMonitorTimer,
	"k8s/cloud-init/artifacts/docker_clear_mount_propagation_flags.conf": k8sCloudInitArtifactsDocker_clear_mount_propagation_flagsConf,
	"k8s/cloud-init/artifacts/enable-dhcpv6.sh":                          k8sCloudInitArtifactsEnableDhcpv6Sh,
	"k8s/cloud-init/artifacts/etc-issue":                                 k8sCloudInitArtifactsEtcIssue,
	"k8s/cloud-init/artifacts/etc-issue.net":                             k8sCloudInitArtifactsEtcIssueNet,
	"k8s/cloud-init/artifacts/etcd.service":                              k8sCloudInitArtifactsEtcdService,
	"k8s/cloud-init/artifacts/generateproxycerts.sh":                     k8sCloudInitArtifactsGenerateproxycertsSh,
	"k8s/cloud-init/artifacts/health-monitor.sh":                         k8sCloudInitArtifactsHealthMonitorSh,
	"k8s/cloud-init/artifacts/kms.service":                               k8sCloudInitArtifactsKmsService,
	"k8s/cloud-init/artifacts/kubelet-monitor.service":                   k8sCloudInitArtifactsKubeletMonitorService,
	"k8s/cloud-init/artifacts/kubelet-monitor.timer":                     k8sCloudInitArtifactsKubeletMonitorTimer,
	"k8s/cloud-init/artifacts/kubelet.service":                           k8sCloudInitArtifactsKubeletService,
	"k8s/cloud-init/artifacts/label-nodes.service":                       k8sCloudInitArtifactsLabelNodesService,
	"k8s/cloud-init/artifacts/label-nodes.sh":                            k8sCloudInitArtifactsLabelNodesSh,
	"k8s/cloud-init/artifacts/modprobe-CIS.conf":                         k8sCloudInitArtifactsModprobeCisConf,
	"k8s/cloud-init/artifacts/pam-d-common-auth":                         k8sCloudInitArtifactsPamDCommonAuth,
	"k8s/cloud-init/artifacts/pam-d-common-password":                     k8sCloudInitArtifactsPamDCommonPassword,
	"k8s/cloud-init/artifacts/pam-d-su":                                  k8sCloudInitArtifactsPamDSu,
	"k8s/cloud-init/artifacts/profile-d-cis.sh":                          k8sCloudInitArtifactsProfileDCisSh,
	"k8s/cloud-init/artifacts/pwquality-CIS.conf":                        k8sCloudInitArtifactsPwqualityCisConf,
	"k8s/cloud-init/artifacts/rsyslog-d-60-CIS.conf":                     k8sCloudInitArtifactsRsyslogD60CisConf,
	"k8s/cloud-init/artifacts/setup-custom-search-domains.sh":            k8sCloudInitArtifactsSetupCustomSearchDomainsSh,
	"k8s/cloud-init/artifacts/sshd_config":                               k8sCloudInitArtifactsSshd_config,
	"k8s/cloud-init/artifacts/sshd_config_1604":                          k8sCloudInitArtifactsSshd_config_1604,
	"k8s/cloud-init/artifacts/sys-fs-bpf.mount":                          k8sCloudInitArtifactsSysFsBpfMount,
	"k8s/cloud-init/artifacts/sysctl-d-60-CIS.conf":                      k8sCloudInitArtifactsSysctlD60CisConf,
	"k8s/cloud-init/artifacts/untaint-nodes.service":                     k8sCloudInitArtifactsUntaintNodesService,
	"k8s/cloud-init/artifacts/untaint-nodes.sh":                          k8sCloudInitArtifactsUntaintNodesSh,
	"k8s/cloud-init/jumpboxcustomdata.yml":                               k8sCloudInitJumpboxcustomdataYml,
	"k8s/cloud-init/masternodecustomdata.yml":                            k8sCloudInitMasternodecustomdataYml,
	"k8s/cloud-init/nodecustomdata.yml":                                  k8sCloudInitNodecustomdataYml,
	"k8s/kubeconfig.json":                                                k8sKubeconfigJson,
	"k8s/kubernetesparams.t":                                             k8sKubernetesparamsT,
	"k8s/kuberneteswindowsfunctions.ps1":                                 k8sKuberneteswindowsfunctionsPs1,
	"k8s/kuberneteswindowssetup.ps1":                                     k8sKuberneteswindowssetupPs1,
	"k8s/manifests/kubernetesmaster-cloud-controller-manager.yaml":       k8sManifestsKubernetesmasterCloudControllerManagerYaml,
	"k8s/manifests/kubernetesmaster-kube-addon-manager.yaml":             k8sManifestsKubernetesmasterKubeAddonManagerYaml,
	"k8s/manifests/kubernetesmaster-kube-apiserver.yaml":                 k8sManifestsKubernetesmasterKubeApiserverYaml,
	"k8s/manifests/kubernetesmaster-kube-controller-manager.yaml":        k8sManifestsKubernetesmasterKubeControllerManagerYaml,
	"k8s/manifests/kubernetesmaster-kube-scheduler.yaml":                 k8sManifestsKubernetesmasterKubeSchedulerYaml,
	"k8s/windowsazurecnifunc.ps1":                                        k8sWindowsazurecnifuncPs1,
	"k8s/windowsazurecnifunc.tests.ps1":                                  k8sWindowsazurecnifuncTestsPs1,
	"k8s/windowscnifunc.ps1":                                             k8sWindowscnifuncPs1,
	"k8s/windowsconfigfunc.ps1":                                          k8sWindowsconfigfuncPs1,
	"k8s/windowscontainerdfunc.ps1":                                      k8sWindowscontainerdfuncPs1,
	"k8s/windowscsiproxyfunc.ps1":                                        k8sWindowscsiproxyfuncPs1,
	"k8s/windowshostsconfigagentfunc.ps1":                                k8sWindowshostsconfigagentfuncPs1,
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
			"1.15": {nil, map[string]*bintree{
				"calico.yaml": {k8sAddons115CalicoYaml, map[string]*bintree{}},
			}},
			"aad-default-admin-group-rbac.yaml":     {k8sAddonsAadDefaultAdminGroupRbacYaml, map[string]*bintree{}},
			"aad-pod-identity.yaml":                 {k8sAddonsAadPodIdentityYaml, map[string]*bintree{}},
			"aci-connector.yaml":                    {k8sAddonsAciConnectorYaml, map[string]*bintree{}},
			"antrea.yaml":                           {k8sAddonsAntreaYaml, map[string]*bintree{}},
			"arc-onboarding.yaml":                   {k8sAddonsArcOnboardingYaml, map[string]*bintree{}},
			"audit-policy.yaml":                     {k8sAddonsAuditPolicyYaml, map[string]*bintree{}},
			"azure-cloud-provider.yaml":             {k8sAddonsAzureCloudProviderYaml, map[string]*bintree{}},
			"azure-cni-networkmonitor.yaml":         {k8sAddonsAzureCniNetworkmonitorYaml, map[string]*bintree{}},
			"azure-network-policy.yaml":             {k8sAddonsAzureNetworkPolicyYaml, map[string]*bintree{}},
			"azure-policy-deployment.yaml":          {k8sAddonsAzurePolicyDeploymentYaml, map[string]*bintree{}},
			"azuredisk-csi-driver-deployment.yaml":  {k8sAddonsAzurediskCsiDriverDeploymentYaml, map[string]*bintree{}},
			"azurefile-csi-driver-deployment.yaml":  {k8sAddonsAzurefileCsiDriverDeploymentYaml, map[string]*bintree{}},
			"blobfuse-flexvolume.yaml":              {k8sAddonsBlobfuseFlexvolumeYaml, map[string]*bintree{}},
			"calico.yaml":                           {k8sAddonsCalicoYaml, map[string]*bintree{}},
			"cilium.yaml":                           {k8sAddonsCiliumYaml, map[string]*bintree{}},
			"cloud-node-manager.yaml":               {k8sAddonsCloudNodeManagerYaml, map[string]*bintree{}},
			"cluster-autoscaler.yaml":               {k8sAddonsClusterAutoscalerYaml, map[string]*bintree{}},
			"container-monitoring.yaml":             {k8sAddonsContainerMonitoringYaml, map[string]*bintree{}},
			"coredns.yaml":                          {k8sAddonsCorednsYaml, map[string]*bintree{}},
			"flannel.yaml":                          {k8sAddonsFlannelYaml, map[string]*bintree{}},
			"ip-masq-agent.yaml":                    {k8sAddonsIpMasqAgentYaml, map[string]*bintree{}},
			"keyvault-flexvolume.yaml":              {k8sAddonsKeyvaultFlexvolumeYaml, map[string]*bintree{}},
			"kube-dns.yaml":                         {k8sAddonsKubeDnsYaml, map[string]*bintree{}},
			"kube-proxy.yaml":                       {k8sAddonsKubeProxyYaml, map[string]*bintree{}},
			"kube-rescheduler.yaml":                 {k8sAddonsKubeReschedulerYaml, map[string]*bintree{}},
			"kubernetes-dashboard.yaml":             {k8sAddonsKubernetesDashboardYaml, map[string]*bintree{}},
			"metrics-server.yaml":                   {k8sAddonsMetricsServerYaml, map[string]*bintree{}},
			"node-problem-detector.yaml":            {k8sAddonsNodeProblemDetectorYaml, map[string]*bintree{}},
			"nvidia-device-plugin.yaml":             {k8sAddonsNvidiaDevicePluginYaml, map[string]*bintree{}},
			"pod-security-policy.yaml":              {k8sAddonsPodSecurityPolicyYaml, map[string]*bintree{}},
			"scheduled-maintenance-deployment.yaml": {k8sAddonsScheduledMaintenanceDeploymentYaml, map[string]*bintree{}},
			"secrets-store-csi-driver.yaml":         {k8sAddonsSecretsStoreCsiDriverYaml, map[string]*bintree{}},
			"smb-flexvolume.yaml":                   {k8sAddonsSmbFlexvolumeYaml, map[string]*bintree{}},
			"tiller.yaml":                           {k8sAddonsTillerYaml, map[string]*bintree{}},
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
				"untaint-nodes.service":                     {k8sCloudInitArtifactsUntaintNodesService, map[string]*bintree{}},
				"untaint-nodes.sh":                          {k8sCloudInitArtifactsUntaintNodesSh, map[string]*bintree{}},
			}},
			"jumpboxcustomdata.yml":    {k8sCloudInitJumpboxcustomdataYml, map[string]*bintree{}},
			"masternodecustomdata.yml": {k8sCloudInitMasternodecustomdataYml, map[string]*bintree{}},
			"nodecustomdata.yml":       {k8sCloudInitNodecustomdataYml, map[string]*bintree{}},
		}},
		"kubeconfig.json":                {k8sKubeconfigJson, map[string]*bintree{}},
		"kubernetesparams.t":             {k8sKubernetesparamsT, map[string]*bintree{}},
		"kuberneteswindowsfunctions.ps1": {k8sKuberneteswindowsfunctionsPs1, map[string]*bintree{}},
		"kuberneteswindowssetup.ps1":     {k8sKuberneteswindowssetupPs1, map[string]*bintree{}},
		"manifests": {nil, map[string]*bintree{
			"kubernetesmaster-cloud-controller-manager.yaml": {k8sManifestsKubernetesmasterCloudControllerManagerYaml, map[string]*bintree{}},
			"kubernetesmaster-kube-addon-manager.yaml":       {k8sManifestsKubernetesmasterKubeAddonManagerYaml, map[string]*bintree{}},
			"kubernetesmaster-kube-apiserver.yaml":           {k8sManifestsKubernetesmasterKubeApiserverYaml, map[string]*bintree{}},
			"kubernetesmaster-kube-controller-manager.yaml":  {k8sManifestsKubernetesmasterKubeControllerManagerYaml, map[string]*bintree{}},
			"kubernetesmaster-kube-scheduler.yaml":           {k8sManifestsKubernetesmasterKubeSchedulerYaml, map[string]*bintree{}},
		}},
		"windowsazurecnifunc.ps1":         {k8sWindowsazurecnifuncPs1, map[string]*bintree{}},
		"windowsazurecnifunc.tests.ps1":   {k8sWindowsazurecnifuncTestsPs1, map[string]*bintree{}},
		"windowscnifunc.ps1":              {k8sWindowscnifuncPs1, map[string]*bintree{}},
		"windowsconfigfunc.ps1":           {k8sWindowsconfigfuncPs1, map[string]*bintree{}},
		"windowscontainerdfunc.ps1":       {k8sWindowscontainerdfuncPs1, map[string]*bintree{}},
		"windowscsiproxyfunc.ps1":         {k8sWindowscsiproxyfuncPs1, map[string]*bintree{}},
		"windowshostsconfigagentfunc.ps1": {k8sWindowshostsconfigagentfuncPs1, map[string]*bintree{}},
		"windowsinstallopensshfunc.ps1":   {k8sWindowsinstallopensshfuncPs1, map[string]*bintree{}},
		"windowskubeletfunc.ps1":          {k8sWindowskubeletfuncPs1, map[string]*bintree{}},
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
