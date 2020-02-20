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
    "kubeProxySpec": {
      "metadata": {
        "description": "The container spec for kube-proxy."
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
      "defaultValue": "3.0.10",
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
         "3.0.10"
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
