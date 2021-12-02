// Code generated for package engine by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../parts/agentoutputs.t
// ../../parts/agentparams.t
// ../../parts/iaasoutputs.t
// ../../parts/k8s/addons/aad-default-admin-group-rbac.yaml
// ../../parts/k8s/addons/aad-pod-identity.yaml
// ../../parts/k8s/addons/antrea.yaml
// ../../parts/k8s/addons/arc-onboarding.yaml
// ../../parts/k8s/addons/audit-policy.yaml
// ../../parts/k8s/addons/azure-cloud-provider.yaml
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
// ../../parts/k8s/cloud-init/artifacts/apiserver-monitor.service
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
// ../../parts/k8s/cloud-init/artifacts/docker_clear_mount_propagation_flags.conf
// ../../parts/k8s/cloud-init/artifacts/enable-dhcpv6.sh
// ../../parts/k8s/cloud-init/artifacts/etc-issue
// ../../parts/k8s/cloud-init/artifacts/etc-issue.net
// ../../parts/k8s/cloud-init/artifacts/etcd-monitor.service
// ../../parts/k8s/cloud-init/artifacts/etcd.service
// ../../parts/k8s/cloud-init/artifacts/generateproxycerts.sh
// ../../parts/k8s/cloud-init/artifacts/health-monitor.sh
// ../../parts/k8s/cloud-init/artifacts/kms-keyvault-key.service
// ../../parts/k8s/cloud-init/artifacts/kms-keyvault-key.sh
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
// ../../parts/k8s/containerdtemplate.toml
// ../../parts/k8s/kubeconfig.json
// ../../parts/k8s/kubernetesparams.t
// ../../parts/k8s/kuberneteswindowsfunctions.ps1
// ../../parts/k8s/kuberneteswindowssetup.ps1
// ../../parts/k8s/manifests/kubernetesmaster-azure-kubernetes-kms.yaml
// ../../parts/k8s/manifests/kubernetesmaster-cloud-controller-manager.yaml
// ../../parts/k8s/manifests/kubernetesmaster-kube-addon-manager.yaml
// ../../parts/k8s/manifests/kubernetesmaster-kube-apiserver.yaml
// ../../parts/k8s/manifests/kubernetesmaster-kube-controller-manager.yaml
// ../../parts/k8s/manifests/kubernetesmaster-kube-scheduler.yaml
// ../../parts/k8s/rotate-certs.ps1
// ../../parts/k8s/rotate-certs.sh
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

var _agentoutputsT = []byte(`{{if and .IsAvailabilitySets .IsStorageAccount}}
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
{{end}}
`)

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
  namespace: kube-system
spec:
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 50%
  selector:
    matchLabels:
      component: nmi
      tier: node
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
          - "--http-probe-port={{ContainerConfig "probePort"}}"
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
            port: {{ContainerConfig "probePort"}}
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
apiVersion: apps/v1
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
  selector:
    matchLabels:
      component: mic
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

var _k8sAddonsAntreaYaml = []byte(`apiVersion: apiextensions.k8s.io/v1
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
    - laai
    singular: antreaagentinfo
  scope: Cluster
  versions:
  - name: v1beta1
    schema:
      openAPIV3Schema:
        type: object
        x-kubernetes-preserve-unknown-fields: true
    served: true
    storage: true
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: antreaagentinfos.crd.antrea.io
spec:
  group: crd.antrea.io
  names:
    kind: AntreaAgentInfo
    plural: antreaagentinfos
    shortNames:
    - aai
    singular: antreaagentinfo
  scope: Cluster
  versions:
  - additionalPrinterColumns:
    - description: Health status of this Agent
      jsonPath: .agentConditions[?(@.type=='AgentHealthy')].status
      name: Healthy
      type: string
    - description: Last time the Healthy Condition was updated
      jsonPath: .agentConditions[?(@.type=='AgentHealthy')].lastHeartbeatTime
      name: Last Heartbeat
      type: date
    - description: Version of this Agent
      jsonPath: .version
      name: Version
      priority: 1
      type: string
    - description: Node on which this Agent is running
      jsonPath: .nodeRef.name
      name: Node
      priority: 1
      type: string
    - description: Number of local Pods managed by this Agent
      jsonPath: .localPodNum
      name: Num Pods
      priority: 2
      type: integer
    - description: Subnets used by this Agent for Pod IPAM
      jsonPath: .nodeSubnets
      name: Subnets
      priority: 2
      type: string
    name: v1beta1
    schema:
      openAPIV3Schema:
        type: object
        x-kubernetes-preserve-unknown-fields: true
    served: true
    storage: true
---
apiVersion: apiextensions.k8s.io/v1
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
    - laci
    singular: antreacontrollerinfo
  scope: Cluster
  versions:
  - name: v1beta1
    schema:
      openAPIV3Schema:
        type: object
        x-kubernetes-preserve-unknown-fields: true
    served: true
    storage: true
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: antreacontrollerinfos.crd.antrea.io
spec:
  group: crd.antrea.io
  names:
    kind: AntreaControllerInfo
    plural: antreacontrollerinfos
    shortNames:
    - aci
    singular: antreacontrollerinfo
  scope: Cluster
  versions:
  - additionalPrinterColumns:
    - description: Health status of the Controller
      jsonPath: .controllerConditions[?(@.type=='ControllerHealthy')].status
      name: Healthy
      type: string
    - description: Last time the Healthy Condition was updated
      jsonPath: .controllerConditions[?(@.type=='ControllerHealthy')].lastHeartbeatTime
      name: Last Heartbeat
      type: date
    - description: Version of the Controller
      jsonPath: .version
      name: Version
      priority: 1
      type: string
    - description: Number of Agents connected to the Controller
      jsonPath: .connectedAgentNum
      name: Connected Agents
      priority: 1
      type: integer
    - description: Node on which the Controller is running
      jsonPath: .nodeRef.name
      name: Node
      priority: 1
      type: string
    - description: Number of Network Policies computed by Controller
      jsonPath: .networkPolicyControllerInfo.networkPolicyNum
      name: Num Network Policies
      priority: 2
      type: integer
    name: v1beta1
    schema:
      openAPIV3Schema:
        type: object
        x-kubernetes-preserve-unknown-fields: true
    served: true
    storage: true
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: clustergroups.core.antrea.tanzu.vmware.com
spec:
  group: core.antrea.tanzu.vmware.com
  names:
    kind: ClusterGroup
    plural: clustergroups
    shortNames:
    - lcg
    singular: group
  scope: Cluster
  versions:
  - name: v1alpha2
    schema:
      openAPIV3Schema:
        properties:
          spec:
            properties:
              childGroups:
                items:
                  type: string
                type: array
              externalEntitySelector:
                properties:
                  matchExpressions:
                    items:
                      properties:
                        key:
                          type: string
                        operator:
                          enum:
                          - In
                          - NotIn
                          - Exists
                          - DoesNotExist
                          type: string
                        values:
                          items:
                            type: string
                          type: array
                      type: object
                    type: array
                  matchLabels:
                    x-kubernetes-preserve-unknown-fields: true
                type: object
              ipBlock:
                properties:
                  cidr:
                    format: cidr
                    type: string
                type: object
              ipBlocks:
                items:
                  properties:
                    cidr:
                      format: cidr
                      type: string
                  type: object
                type: array
              namespaceSelector:
                properties:
                  matchExpressions:
                    items:
                      properties:
                        key:
                          type: string
                        operator:
                          enum:
                          - In
                          - NotIn
                          - Exists
                          - DoesNotExist
                          type: string
                        values:
                          items:
                            type: string
                          type: array
                      type: object
                    type: array
                  matchLabels:
                    x-kubernetes-preserve-unknown-fields: true
                type: object
              podSelector:
                properties:
                  matchExpressions:
                    items:
                      properties:
                        key:
                          type: string
                        operator:
                          enum:
                          - In
                          - NotIn
                          - Exists
                          - DoesNotExist
                          type: string
                        values:
                          items:
                            type: string
                          type: array
                      type: object
                    type: array
                  matchLabels:
                    x-kubernetes-preserve-unknown-fields: true
                type: object
              serviceReference:
                properties:
                  name:
                    type: string
                  namespace:
                    type: string
                type: object
            type: object
          status:
            properties:
              conditions:
                items:
                  properties:
                    lastTransitionTime:
                      type: string
                    status:
                      type: string
                    type:
                      type: string
                  type: object
                type: array
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: clustergroups.crd.antrea.io
spec:
  conversion:
    strategy: Webhook
    webhook:
      clientConfig:
        service:
          name: antrea
          namespace: kube-system
          path: /convert/clustergroup
      conversionReviewVersions:
      - v1
      - v1beta1
  group: crd.antrea.io
  names:
    kind: ClusterGroup
    plural: clustergroups
    shortNames:
    - cg
    singular: clustergroup
  scope: Cluster
  versions:
  - name: v1alpha2
    schema:
      openAPIV3Schema:
        properties:
          spec:
            properties:
              childGroups:
                items:
                  type: string
                type: array
              externalEntitySelector:
                properties:
                  matchExpressions:
                    items:
                      properties:
                        key:
                          type: string
                        operator:
                          enum:
                          - In
                          - NotIn
                          - Exists
                          - DoesNotExist
                          type: string
                        values:
                          items:
                            type: string
                          type: array
                      type: object
                    type: array
                  matchLabels:
                    x-kubernetes-preserve-unknown-fields: true
                type: object
              ipBlock:
                properties:
                  cidr:
                    format: cidr
                    type: string
                type: object
              ipBlocks:
                items:
                  properties:
                    cidr:
                      format: cidr
                      type: string
                  type: object
                type: array
              namespaceSelector:
                properties:
                  matchExpressions:
                    items:
                      properties:
                        key:
                          type: string
                        operator:
                          enum:
                          - In
                          - NotIn
                          - Exists
                          - DoesNotExist
                          type: string
                        values:
                          items:
                            type: string
                          type: array
                      type: object
                    type: array
                  matchLabels:
                    x-kubernetes-preserve-unknown-fields: true
                type: object
              podSelector:
                properties:
                  matchExpressions:
                    items:
                      properties:
                        key:
                          type: string
                        operator:
                          enum:
                          - In
                          - NotIn
                          - Exists
                          - DoesNotExist
                          type: string
                        values:
                          items:
                            type: string
                          type: array
                      type: object
                    type: array
                  matchLabels:
                    x-kubernetes-preserve-unknown-fields: true
                type: object
              serviceReference:
                properties:
                  name:
                    type: string
                  namespace:
                    type: string
                type: object
            type: object
          status:
            properties:
              conditions:
                items:
                  properties:
                    lastTransitionTime:
                      type: string
                    status:
                      type: string
                    type:
                      type: string
                  type: object
                type: array
            type: object
        type: object
    served: true
    storage: false
  - name: v1alpha3
    schema:
      openAPIV3Schema:
        properties:
          spec:
            properties:
              childGroups:
                items:
                  type: string
                type: array
              externalEntitySelector:
                properties:
                  matchExpressions:
                    items:
                      properties:
                        key:
                          type: string
                        operator:
                          enum:
                          - In
                          - NotIn
                          - Exists
                          - DoesNotExist
                          type: string
                        values:
                          items:
                            type: string
                          type: array
                      type: object
                    type: array
                  matchLabels:
                    x-kubernetes-preserve-unknown-fields: true
                type: object
              ipBlocks:
                items:
                  properties:
                    cidr:
                      format: cidr
                      type: string
                  type: object
                type: array
              namespaceSelector:
                properties:
                  matchExpressions:
                    items:
                      properties:
                        key:
                          type: string
                        operator:
                          enum:
                          - In
                          - NotIn
                          - Exists
                          - DoesNotExist
                          type: string
                        values:
                          items:
                            type: string
                          type: array
                      type: object
                    type: array
                  matchLabels:
                    x-kubernetes-preserve-unknown-fields: true
                type: object
              podSelector:
                properties:
                  matchExpressions:
                    items:
                      properties:
                        key:
                          type: string
                        operator:
                          enum:
                          - In
                          - NotIn
                          - Exists
                          - DoesNotExist
                          type: string
                        values:
                          items:
                            type: string
                          type: array
                      type: object
                    type: array
                  matchLabels:
                    x-kubernetes-preserve-unknown-fields: true
                type: object
              serviceReference:
                properties:
                  name:
                    type: string
                  namespace:
                    type: string
                type: object
            type: object
          status:
            properties:
              conditions:
                items:
                  properties:
                    lastTransitionTime:
                      type: string
                    status:
                      type: string
                    type:
                      type: string
                  type: object
                type: array
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: clusternetworkpolicies.crd.antrea.io
spec:
  group: crd.antrea.io
  names:
    kind: ClusterNetworkPolicy
    plural: clusternetworkpolicies
    shortNames:
    - acnp
    singular: clusternetworkpolicy
  scope: Cluster
  versions:
  - additionalPrinterColumns:
    - description: The Tier to which this ClusterNetworkPolicy belongs to.
      jsonPath: .spec.tier
      name: Tier
      type: string
    - description: The Priority of this ClusterNetworkPolicy relative to other policies.
      format: float
      jsonPath: .spec.priority
      name: Priority
      type: number
    - description: The total number of Nodes that should realize the NetworkPolicy.
      format: int32
      jsonPath: .status.desiredNodesRealized
      name: Desired Nodes
      type: number
    - description: The number of Nodes that have realized the NetworkPolicy.
      format: int32
      jsonPath: .status.currentNodesRealized
      name: Current Nodes
      type: number
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1alpha1
    schema:
      openAPIV3Schema:
        properties:
          spec:
            properties:
              appliedTo:
                items:
                  properties:
                    group:
                      type: string
                    namespaceSelector:
                      properties:
                        matchExpressions:
                          items:
                            properties:
                              key:
                                type: string
                              operator:
                                enum:
                                - In
                                - NotIn
                                - Exists
                                - DoesNotExist
                                type: string
                              values:
                                items:
                                  type: string
                                type: array
                            type: object
                          type: array
                        matchLabels:
                          x-kubernetes-preserve-unknown-fields: true
                      type: object
                    podSelector:
                      properties:
                        matchExpressions:
                          items:
                            properties:
                              key:
                                type: string
                              operator:
                                enum:
                                - In
                                - NotIn
                                - Exists
                                - DoesNotExist
                                type: string
                              values:
                                items:
                                  type: string
                                type: array
                            type: object
                          type: array
                        matchLabels:
                          x-kubernetes-preserve-unknown-fields: true
                      type: object
                  type: object
                type: array
              egress:
                items:
                  properties:
                    action:
                      enum:
                      - Allow
                      - Drop
                      - Reject
                      type: string
                    appliedTo:
                      items:
                        properties:
                          group:
                            type: string
                          namespaceSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                          podSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                        type: object
                      type: array
                    enableLogging:
                      type: boolean
                    name:
                      type: string
                    ports:
                      items:
                        properties:
                          endPort:
                            type: integer
                          port:
                            x-kubernetes-int-or-string: true
                          protocol:
                            type: string
                        type: object
                      type: array
                    to:
                      items:
                        properties:
                          fqdn:
                            type: string
                          group:
                            type: string
                          ipBlock:
                            properties:
                              cidr:
                                format: cidr
                                type: string
                            type: object
                          namespaceSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                          namespaces:
                            properties:
                              match:
                                type: string
                            type: object
                          podSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                        type: object
                      type: array
                  required:
                  - action
                  type: object
                type: array
              ingress:
                items:
                  properties:
                    action:
                      enum:
                      - Allow
                      - Drop
                      - Reject
                      type: string
                    appliedTo:
                      items:
                        properties:
                          group:
                            type: string
                          namespaceSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                          podSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                        type: object
                      type: array
                    enableLogging:
                      type: boolean
                    from:
                      items:
                        properties:
                          group:
                            type: string
                          ipBlock:
                            properties:
                              cidr:
                                format: cidr
                                type: string
                            type: object
                          namespaceSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                          namespaces:
                            properties:
                              match:
                                type: string
                            type: object
                          podSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                        type: object
                      type: array
                    name:
                      type: string
                    ports:
                      items:
                        properties:
                          endPort:
                            type: integer
                          port:
                            x-kubernetes-int-or-string: true
                          protocol:
                            type: string
                        type: object
                      type: array
                  required:
                  - action
                  type: object
                type: array
              priority:
                format: float
                maximum: 10000
                minimum: 1
                type: number
              tier:
                type: string
            required:
            - priority
            type: object
          status:
            properties:
              currentNodesRealized:
                type: integer
              desiredNodesRealized:
                type: integer
              observedGeneration:
                type: integer
              phase:
                type: string
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: clusternetworkpolicies.security.antrea.tanzu.vmware.com
spec:
  group: security.antrea.tanzu.vmware.com
  names:
    kind: ClusterNetworkPolicy
    plural: clusternetworkpolicies
    shortNames:
    - lacnp
    singular: clusternetworkpolicy
  scope: Cluster
  versions:
  - additionalPrinterColumns:
    - description: The Tier to which this ClusterNetworkPolicy belongs to.
      jsonPath: .spec.tier
      name: Tier
      type: string
    - description: The Priority of this ClusterNetworkPolicy relative to other policies.
      format: float
      jsonPath: .spec.priority
      name: Priority
      type: number
    - description: The total number of Nodes that should realize the NetworkPolicy.
      format: int32
      jsonPath: .status.desiredNodesRealized
      name: Desired Nodes
      type: number
    - description: The number of Nodes that have realized the NetworkPolicy.
      format: int32
      jsonPath: .status.currentNodesRealized
      name: Current Nodes
      type: number
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1alpha1
    schema:
      openAPIV3Schema:
        properties:
          spec:
            properties:
              appliedTo:
                items:
                  properties:
                    group:
                      type: string
                    namespaceSelector:
                      properties:
                        matchExpressions:
                          items:
                            properties:
                              key:
                                type: string
                              operator:
                                enum:
                                - In
                                - NotIn
                                - Exists
                                - DoesNotExist
                                type: string
                              values:
                                items:
                                  type: string
                                type: array
                            type: object
                          type: array
                        matchLabels:
                          x-kubernetes-preserve-unknown-fields: true
                      type: object
                    podSelector:
                      properties:
                        matchExpressions:
                          items:
                            properties:
                              key:
                                type: string
                              operator:
                                enum:
                                - In
                                - NotIn
                                - Exists
                                - DoesNotExist
                                type: string
                              values:
                                items:
                                  type: string
                                type: array
                            type: object
                          type: array
                        matchLabels:
                          x-kubernetes-preserve-unknown-fields: true
                      type: object
                  type: object
                type: array
              egress:
                items:
                  properties:
                    action:
                      enum:
                      - Allow
                      - Drop
                      - Reject
                      type: string
                    appliedTo:
                      items:
                        properties:
                          group:
                            type: string
                          namespaceSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                          podSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                        type: object
                      type: array
                    enableLogging:
                      type: boolean
                    name:
                      type: string
                    ports:
                      items:
                        properties:
                          endPort:
                            type: integer
                          port:
                            x-kubernetes-int-or-string: true
                          protocol:
                            type: string
                        type: object
                      type: array
                    to:
                      items:
                        properties:
                          group:
                            type: string
                          ipBlock:
                            properties:
                              cidr:
                                format: cidr
                                type: string
                            type: object
                          namespaceSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                          namespaces:
                            properties:
                              match:
                                type: string
                            type: object
                          podSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                        type: object
                      type: array
                  required:
                  - action
                  type: object
                type: array
              ingress:
                items:
                  properties:
                    action:
                      enum:
                      - Allow
                      - Drop
                      - Reject
                      type: string
                    appliedTo:
                      items:
                        properties:
                          group:
                            type: string
                          namespaceSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                          podSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                        type: object
                      type: array
                    enableLogging:
                      type: boolean
                    from:
                      items:
                        properties:
                          group:
                            type: string
                          ipBlock:
                            properties:
                              cidr:
                                format: cidr
                                type: string
                            type: object
                          namespaceSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                          namespaces:
                            properties:
                              match:
                                type: string
                            type: object
                          podSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                        type: object
                      type: array
                    name:
                      type: string
                    ports:
                      items:
                        properties:
                          endPort:
                            type: integer
                          port:
                            x-kubernetes-int-or-string: true
                          protocol:
                            type: string
                        type: object
                      type: array
                  required:
                  - action
                  type: object
                type: array
              priority:
                format: float
                maximum: 10000
                minimum: 1
                type: number
              tier:
                type: string
            required:
            - priority
            type: object
          status:
            properties:
              currentNodesRealized:
                type: integer
              desiredNodesRealized:
                type: integer
              observedGeneration:
                type: integer
              phase:
                type: string
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: egresses.crd.antrea.io
spec:
  group: crd.antrea.io
  names:
    kind: Egress
    plural: egresses
    shortNames:
    - eg
    singular: egress
  scope: Cluster
  versions:
  - additionalPrinterColumns:
    - description: Specifies the SNAT IP address for the selected workloads.
      jsonPath: .spec.egressIP
      name: EgressIP
      type: string
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    - description: The Owner Node of egress IP
      jsonPath: .status.egressNode
      name: Node
      type: string
    name: v1alpha2
    schema:
      openAPIV3Schema:
        properties:
          spec:
            anyOf:
            - required:
              - egressIP
            - required:
              - externalIPPool
            properties:
              appliedTo:
                properties:
                  namespaceSelector:
                    properties:
                      matchExpressions:
                        items:
                          properties:
                            key:
                              type: string
                            operator:
                              enum:
                              - In
                              - NotIn
                              - Exists
                              - DoesNotExist
                              type: string
                            values:
                              items:
                                type: string
                              type: array
                          type: object
                        type: array
                      matchLabels:
                        x-kubernetes-preserve-unknown-fields: true
                    type: object
                  podSelector:
                    properties:
                      matchExpressions:
                        items:
                          properties:
                            key:
                              type: string
                            operator:
                              enum:
                              - In
                              - NotIn
                              - Exists
                              - DoesNotExist
                              type: string
                            values:
                              items:
                                type: string
                              type: array
                          type: object
                        type: array
                      matchLabels:
                        x-kubernetes-preserve-unknown-fields: true
                    type: object
                type: object
              egressIP:
                oneOf:
                - format: ipv4
                - format: ipv6
                type: string
              externalIPPool:
                type: string
            required:
            - appliedTo
            type: object
          status:
            properties:
              egressNode:
                type: string
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: externalentities.core.antrea.tanzu.vmware.com
spec:
  group: core.antrea.tanzu.vmware.com
  names:
    kind: ExternalEntity
    plural: externalentities
    shortNames:
    - lee
    singular: externalentity
  scope: Namespaced
  versions:
  - name: v1alpha2
    schema:
      openAPIV3Schema:
        properties:
          spec:
            properties:
              endpoints:
                items:
                  properties:
                    ip:
                      oneOf:
                      - format: ipv4
                      - format: ipv6
                      type: string
                    name:
                      type: string
                  type: object
                type: array
              externalNode:
                type: string
              ports:
                items:
                  properties:
                    name:
                      type: string
                    port:
                      x-kubernetes-int-or-string: true
                    protocol:
                      type: string
                  type: object
                type: array
            type: object
        type: object
    served: true
    storage: true
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        type: object
    served: false
    storage: false
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: externalentities.crd.antrea.io
spec:
  group: crd.antrea.io
  names:
    kind: ExternalEntity
    plural: externalentities
    shortNames:
    - ee
    singular: externalentity
  scope: Namespaced
  versions:
  - name: v1alpha2
    schema:
      openAPIV3Schema:
        properties:
          spec:
            properties:
              endpoints:
                items:
                  properties:
                    ip:
                      oneOf:
                      - format: ipv4
                      - format: ipv6
                      type: string
                    name:
                      type: string
                  type: object
                type: array
              externalNode:
                type: string
              ports:
                items:
                  properties:
                    name:
                      type: string
                    port:
                      x-kubernetes-int-or-string: true
                    protocol:
                      type: string
                  type: object
                type: array
            type: object
        type: object
    served: true
    storage: true
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        type: object
    served: false
    storage: false
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: externalippools.crd.antrea.io
spec:
  group: crd.antrea.io
  names:
    kind: ExternalIPPool
    plural: externalippools
    shortNames:
    - eip
    singular: externalippool
  scope: Cluster
  versions:
  - additionalPrinterColumns:
    - description: The number of total IPs
      jsonPath: .status.usage.total
      name: Total
      type: integer
    - description: The number of allocated IPs
      jsonPath: .status.usage.used
      name: Used
      type: integer
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1alpha2
    schema:
      openAPIV3Schema:
        properties:
          spec:
            properties:
              ipRanges:
                items:
                  oneOf:
                  - required:
                    - cidr
                  - required:
                    - start
                    - end
                  properties:
                    cidr:
                      format: cidr
                      type: string
                    end:
                      oneOf:
                      - format: ipv4
                      - format: ipv6
                      type: string
                    start:
                      oneOf:
                      - format: ipv4
                      - format: ipv6
                      type: string
                  type: object
                type: array
              nodeSelector:
                properties:
                  matchExpressions:
                    items:
                      properties:
                        key:
                          type: string
                        operator:
                          enum:
                          - In
                          - NotIn
                          - Exists
                          - DoesNotExist
                          type: string
                        values:
                          items:
                            type: string
                          type: array
                      type: object
                    type: array
                  matchLabels:
                    x-kubernetes-preserve-unknown-fields: true
                type: object
            required:
            - ipRanges
            - nodeSelector
            type: object
          status:
            properties:
              usage:
                properties:
                  total:
                    type: integer
                  used:
                    type: integer
                type: object
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: networkpolicies.crd.antrea.io
spec:
  group: crd.antrea.io
  names:
    kind: NetworkPolicy
    plural: networkpolicies
    shortNames:
    - anp
    singular: networkpolicy
  scope: Namespaced
  versions:
  - additionalPrinterColumns:
    - description: The Tier to which this Antrea NetworkPolicy belongs to.
      jsonPath: .spec.tier
      name: Tier
      type: string
    - description: The Priority of this Antrea NetworkPolicy relative to other policies.
      format: float
      jsonPath: .spec.priority
      name: Priority
      type: number
    - description: The total number of Nodes that should realize the NetworkPolicy.
      format: int32
      jsonPath: .status.desiredNodesRealized
      name: Desired Nodes
      type: number
    - description: The number of Nodes that have realized the NetworkPolicy.
      format: int32
      jsonPath: .status.currentNodesRealized
      name: Current Nodes
      type: number
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1alpha1
    schema:
      openAPIV3Schema:
        properties:
          spec:
            properties:
              appliedTo:
                items:
                  properties:
                    podSelector:
                      properties:
                        matchExpressions:
                          items:
                            properties:
                              key:
                                type: string
                              operator:
                                enum:
                                - In
                                - NotIn
                                - Exists
                                - DoesNotExist
                                type: string
                              values:
                                items:
                                  type: string
                                type: array
                            type: object
                          type: array
                        matchLabels:
                          x-kubernetes-preserve-unknown-fields: true
                      type: object
                  type: object
                type: array
              egress:
                items:
                  properties:
                    action:
                      enum:
                      - Allow
                      - Drop
                      - Reject
                      type: string
                    appliedTo:
                      items:
                        properties:
                          podSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                        type: object
                      type: array
                    enableLogging:
                      type: boolean
                    name:
                      type: string
                    ports:
                      items:
                        properties:
                          endPort:
                            type: integer
                          port:
                            x-kubernetes-int-or-string: true
                          protocol:
                            type: string
                        type: object
                      type: array
                    to:
                      items:
                        properties:
                          externalEntitySelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                          fqdn:
                            type: string
                          ipBlock:
                            properties:
                              cidr:
                                format: cidr
                                type: string
                            type: object
                          namespaceSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                          podSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                        type: object
                      type: array
                  required:
                  - action
                  type: object
                type: array
              ingress:
                items:
                  properties:
                    action:
                      enum:
                      - Allow
                      - Drop
                      - Reject
                      type: string
                    appliedTo:
                      items:
                        properties:
                          podSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                        type: object
                      type: array
                    enableLogging:
                      type: boolean
                    from:
                      items:
                        properties:
                          externalEntitySelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                          ipBlock:
                            properties:
                              cidr:
                                format: cidr
                                type: string
                            type: object
                          namespaceSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                          podSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                        type: object
                      type: array
                    name:
                      type: string
                    ports:
                      items:
                        properties:
                          endPort:
                            type: integer
                          port:
                            x-kubernetes-int-or-string: true
                          protocol:
                            type: string
                        type: object
                      type: array
                  required:
                  - action
                  type: object
                type: array
              priority:
                format: float
                maximum: 10000
                minimum: 1
                type: number
              tier:
                type: string
            required:
            - priority
            type: object
          status:
            properties:
              currentNodesRealized:
                type: integer
              desiredNodesRealized:
                type: integer
              observedGeneration:
                type: integer
              phase:
                type: string
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: networkpolicies.security.antrea.tanzu.vmware.com
spec:
  group: security.antrea.tanzu.vmware.com
  names:
    kind: NetworkPolicy
    plural: networkpolicies
    shortNames:
    - lanp
    singular: networkpolicy
  scope: Namespaced
  versions:
  - additionalPrinterColumns:
    - description: The Tier to which this Antrea NetworkPolicy belongs to.
      jsonPath: .spec.tier
      name: Tier
      type: string
    - description: The Priority of this Antrea NetworkPolicy relative to other policies.
      format: float
      jsonPath: .spec.priority
      name: Priority
      type: number
    - description: The total number of Nodes that should realize the NetworkPolicy.
      format: int32
      jsonPath: .status.desiredNodesRealized
      name: Desired Nodes
      type: number
    - description: The number of Nodes that have realized the NetworkPolicy.
      format: int32
      jsonPath: .status.currentNodesRealized
      name: Current Nodes
      type: number
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1alpha1
    schema:
      openAPIV3Schema:
        properties:
          spec:
            properties:
              appliedTo:
                items:
                  properties:
                    podSelector:
                      properties:
                        matchExpressions:
                          items:
                            properties:
                              key:
                                type: string
                              operator:
                                enum:
                                - In
                                - NotIn
                                - Exists
                                - DoesNotExist
                                type: string
                              values:
                                items:
                                  type: string
                                type: array
                            type: object
                          type: array
                        matchLabels:
                          x-kubernetes-preserve-unknown-fields: true
                      type: object
                  type: object
                type: array
              egress:
                items:
                  properties:
                    action:
                      enum:
                      - Allow
                      - Drop
                      - Reject
                      type: string
                    appliedTo:
                      items:
                        properties:
                          podSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                        type: object
                      type: array
                    enableLogging:
                      type: boolean
                    name:
                      type: string
                    ports:
                      items:
                        properties:
                          endPort:
                            type: integer
                          port:
                            x-kubernetes-int-or-string: true
                          protocol:
                            type: string
                        type: object
                      type: array
                    to:
                      items:
                        properties:
                          externalEntitySelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                          ipBlock:
                            properties:
                              cidr:
                                format: cidr
                                type: string
                            type: object
                          namespaceSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                          podSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                        type: object
                      type: array
                  required:
                  - action
                  type: object
                type: array
              ingress:
                items:
                  properties:
                    action:
                      enum:
                      - Allow
                      - Drop
                      - Reject
                      type: string
                    appliedTo:
                      items:
                        properties:
                          podSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                        type: object
                      type: array
                    enableLogging:
                      type: boolean
                    from:
                      items:
                        properties:
                          externalEntitySelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                          ipBlock:
                            properties:
                              cidr:
                                format: cidr
                                type: string
                            type: object
                          namespaceSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                          podSelector:
                            properties:
                              matchExpressions:
                                items:
                                  properties:
                                    key:
                                      type: string
                                    operator:
                                      enum:
                                      - In
                                      - NotIn
                                      - Exists
                                      - DoesNotExist
                                      type: string
                                    values:
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                type: array
                              matchLabels:
                                x-kubernetes-preserve-unknown-fields: true
                            type: object
                        type: object
                      type: array
                    name:
                      type: string
                    ports:
                      items:
                        properties:
                          endPort:
                            type: integer
                          port:
                            x-kubernetes-int-or-string: true
                          protocol:
                            type: string
                        type: object
                      type: array
                  required:
                  - action
                  type: object
                type: array
              priority:
                format: float
                maximum: 10000
                minimum: 1
                type: number
              tier:
                type: string
            required:
            - priority
            type: object
          status:
            properties:
              currentNodesRealized:
                type: integer
              desiredNodesRealized:
                type: integer
              observedGeneration:
                type: integer
              phase:
                type: string
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: tiers.crd.antrea.io
spec:
  group: crd.antrea.io
  names:
    kind: Tier
    plural: tiers
    shortNames:
    - tr
    singular: tier
  scope: Cluster
  versions:
  - additionalPrinterColumns:
    - description: The Priority of this Tier relative to other Tiers.
      jsonPath: .spec.priority
      name: Priority
      type: integer
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1alpha1
    schema:
      openAPIV3Schema:
        properties:
          spec:
            properties:
              description:
                type: string
              priority:
                maximum: 255
                minimum: 0
                type: integer
            required:
            - priority
            type: object
        type: object
    served: true
    storage: true
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: tiers.security.antrea.tanzu.vmware.com
spec:
  group: security.antrea.tanzu.vmware.com
  names:
    kind: Tier
    plural: tiers
    shortNames:
    - ltr
    singular: tier
  scope: Cluster
  versions:
  - additionalPrinterColumns:
    - description: The Priority of this Tier relative to other Tiers.
      jsonPath: .spec.priority
      name: Priority
      type: integer
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1alpha1
    schema:
      openAPIV3Schema:
        properties:
          spec:
            properties:
              description:
                type: string
              priority:
                maximum: 255
                minimum: 0
                type: integer
            required:
            - priority
            type: object
        type: object
    served: true
    storage: true
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: traceflows.crd.antrea.io
spec:
  group: crd.antrea.io
  names:
    kind: Traceflow
    plural: traceflows
    shortNames:
    - tf
    singular: traceflow
  scope: Cluster
  versions:
  - additionalPrinterColumns:
    - description: The phase of the Traceflow.
      jsonPath: .status.phase
      name: Phase
      type: string
    - description: The name of the source Pod.
      jsonPath: .spec.source.pod
      name: Source-Pod
      priority: 10
      type: string
    - description: The name of the destination Pod.
      jsonPath: .spec.destination.pod
      name: Destination-Pod
      priority: 10
      type: string
    - description: The IP address of the destination.
      jsonPath: .spec.destination.ip
      name: Destination-IP
      priority: 10
      type: string
    - description: Trace live traffic.
      jsonPath: .spec.liveTraffic
      name: Live-Traffic
      priority: 10
      type: boolean
    - description: Capture only the dropped packet.
      jsonPath: .spec.droppedOnly
      name: Dropped-Only
      priority: 10
      type: boolean
    - description: Timeout in seconds.
      jsonPath: .spec.timeout
      name: Timeout
      priority: 10
      type: integer
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1alpha1
    schema:
      openAPIV3Schema:
        properties:
          spec:
            properties:
              destination:
                properties:
                  ip:
                    oneOf:
                    - format: ipv4
                    - format: ipv6
                    type: string
                  namespace:
                    type: string
                  pod:
                    type: string
                  service:
                    type: string
                type: object
              droppedOnly:
                type: boolean
              liveTraffic:
                type: boolean
              packet:
                properties:
                  ipHeader:
                    properties:
                      flags:
                        type: integer
                      protocol:
                        type: integer
                      srcIP:
                        oneOf:
                        - format: ipv4
                        - format: ipv6
                        type: string
                      ttl:
                        type: integer
                    type: object
                  ipv6Header:
                    properties:
                      hopLimit:
                        type: integer
                      nextHeader:
                        type: integer
                      srcIP:
                        format: ipv6
                        type: string
                    type: object
                  transportHeader:
                    properties:
                      icmp:
                        properties:
                          id:
                            type: integer
                          sequence:
                            type: integer
                        type: object
                      tcp:
                        properties:
                          dstPort:
                            type: integer
                          flags:
                            type: integer
                          srcPort:
                            type: integer
                        type: object
                      udp:
                        properties:
                          dstPort:
                            type: integer
                          srcPort:
                            type: integer
                        type: object
                    type: object
                type: object
              source:
                properties:
                  ip:
                    oneOf:
                    - format: ipv4
                    - format: ipv6
                    type: string
                  namespace:
                    type: string
                  pod:
                    type: string
                type: object
              timeout:
                type: integer
            type: object
          status:
            properties:
              capturedPacket:
                properties:
                  dstIP:
                    type: string
                  ipHeader:
                    properties:
                      flags:
                        type: integer
                      protocol:
                        type: integer
                      ttl:
                        type: integer
                    type: object
                  ipv6Header:
                    properties:
                      hopLimit:
                        type: integer
                      nextHeader:
                        type: integer
                    type: object
                  length:
                    type: integer
                  srcIP:
                    type: string
                  transportHeader:
                    properties:
                      icmp:
                        properties:
                          id:
                            type: integer
                          sequence:
                            type: integer
                        type: object
                      tcp:
                        properties:
                          dstPort:
                            type: integer
                          flags:
                            type: integer
                          srcPort:
                            type: integer
                        type: object
                      udp:
                        properties:
                          dstPort:
                            type: integer
                          srcPort:
                            type: integer
                        type: object
                    type: object
                type: object
              dataplaneTag:
                type: integer
              phase:
                type: string
              reason:
                type: string
              results:
                items:
                  properties:
                    node:
                      type: string
                    observations:
                      items:
                        properties:
                          action:
                            type: string
                          component:
                            type: string
                          componentInfo:
                            type: string
                          dstMAC:
                            type: string
                          networkPolicy:
                            type: string
                          pod:
                            type: string
                          translatedDstIP:
                            type: string
                          translatedSrcIP:
                            type: string
                          ttl:
                            type: integer
                          tunnelDstIP:
                            type: string
                        type: object
                      type: array
                    role:
                      type: string
                    timestamp:
                      type: integer
                  type: object
                type: array
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: traceflows.ops.antrea.tanzu.vmware.com
spec:
  group: ops.antrea.tanzu.vmware.com
  names:
    kind: Traceflow
    plural: traceflows
    shortNames:
    - ltf
    singular: traceflow
  scope: Cluster
  versions:
  - additionalPrinterColumns:
    - description: The phase of the Traceflow.
      jsonPath: .status.phase
      name: Phase
      type: string
    - description: The name of the source Pod.
      jsonPath: .spec.source.pod
      name: Source-Pod
      priority: 10
      type: string
    - description: The name of the destination Pod.
      jsonPath: .spec.destination.pod
      name: Destination-Pod
      priority: 10
      type: string
    - description: The IP address of the destination.
      jsonPath: .spec.destination.ip
      name: Destination-IP
      priority: 10
      type: string
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1alpha1
    schema:
      openAPIV3Schema:
        properties:
          spec:
            properties:
              destination:
                properties:
                  ip:
                    oneOf:
                    - format: ipv4
                    - format: ipv6
                    type: string
                  namespace:
                    type: string
                  pod:
                    type: string
                  service:
                    type: string
                type: object
              packet:
                properties:
                  ipHeader:
                    properties:
                      flags:
                        type: integer
                      protocol:
                        type: integer
                      srcIP:
                        oneOf:
                        - format: ipv4
                        - format: ipv6
                        type: string
                      ttl:
                        type: integer
                    type: object
                  ipv6Header:
                    properties:
                      hopLimit:
                        type: integer
                      nextHeader:
                        type: integer
                      srcIP:
                        format: ipv6
                        type: string
                    type: object
                  transportHeader:
                    properties:
                      icmp:
                        properties:
                          id:
                            type: integer
                          sequence:
                            type: integer
                        type: object
                      tcp:
                        properties:
                          dstPort:
                            type: integer
                          flags:
                            type: integer
                          srcPort:
                            type: integer
                        type: object
                      udp:
                        properties:
                          dstPort:
                            type: integer
                          srcPort:
                            type: integer
                        type: object
                    type: object
                type: object
              source:
                properties:
                  namespace:
                    type: string
                  pod:
                    type: string
                required:
                - pod
                - namespace
                type: object
            required:
            - source
            type: object
          status:
            properties:
              dataplaneTag:
                type: integer
              phase:
                type: string
              reason:
                type: string
              results:
                items:
                  properties:
                    node:
                      type: string
                    observations:
                      items:
                        properties:
                          action:
                            type: string
                          component:
                            type: string
                          componentInfo:
                            type: string
                          dstMAC:
                            type: string
                          networkPolicy:
                            type: string
                          pod:
                            type: string
                          translatedDstIP:
                            type: string
                          translatedSrcIP:
                            type: string
                          ttl:
                            type: integer
                          tunnelDstIP:
                            type: string
                        type: object
                      type: array
                    role:
                      type: string
                    timestamp:
                      type: integer
                  type: object
                type: array
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
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
    rbac.authorization.k8s.io/aggregate-to-admin: "true"
    rbac.authorization.k8s.io/aggregate-to-edit: "true"
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: aggregate-antrea-clustergroups-edit
rules:
- apiGroups:
  - core.antrea.tanzu.vmware.com
  resources:
  - clustergroups
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - crd.antrea.io
  resources:
  - clustergroups
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
  labels:
    app: antrea
    rbac.authorization.k8s.io/aggregate-to-view: "true"
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: aggregate-antrea-clustergroups-view
rules:
- apiGroups:
  - core.antrea.tanzu.vmware.com
  resources:
  - clustergroups
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - crd.antrea.io
  resources:
  - clustergroups
  verbs:
  - get
  - list
  - watch
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    app: antrea
    rbac.authorization.k8s.io/aggregate-to-admin: "true"
    rbac.authorization.k8s.io/aggregate-to-edit: "true"
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: aggregate-antrea-policies-edit
rules:
- apiGroups:
  - security.antrea.tanzu.vmware.com
  resources:
  - clusternetworkpolicies
  - networkpolicies
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - crd.antrea.io
  resources:
  - clusternetworkpolicies
  - networkpolicies
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
  labels:
    app: antrea
    rbac.authorization.k8s.io/aggregate-to-view: "true"
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: aggregate-antrea-policies-view
rules:
- apiGroups:
  - security.antrea.tanzu.vmware.com
  resources:
  - clusternetworkpolicies
  - networkpolicies
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - crd.antrea.io
  resources:
  - clusternetworkpolicies
  - networkpolicies
  verbs:
  - get
  - list
  - watch
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    app: antrea
    rbac.authorization.k8s.io/aggregate-to-admin: "true"
    rbac.authorization.k8s.io/aggregate-to-edit: "true"
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: aggregate-traceflows-edit
rules:
- apiGroups:
  - ops.antrea.tanzu.vmware.com
  resources:
  - traceflows
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - crd.antrea.io
  resources:
  - traceflows
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
  labels:
    app: antrea
    rbac.authorization.k8s.io/aggregate-to-view: "true"
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: aggregate-traceflows-view
rules:
- apiGroups:
  - ops.antrea.tanzu.vmware.com
  resources:
  - traceflows
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - crd.antrea.io
  resources:
  - traceflows
  verbs:
  - get
  - list
  - watch
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
  - controlplane.antrea.tanzu.vmware.com
  - controlplane.antrea.io
  resources:
  - networkpolicies
  - appliedtogroups
  - addressgroups
  verbs:
  - get
  - list
- apiGroups:
  - stats.antrea.tanzu.vmware.com
  - stats.antrea.io
  resources:
  - networkpolicystats
  - antreaclusternetworkpolicystats
  - antreanetworkpolicystats
  verbs:
  - get
  - list
- apiGroups:
  - system.antrea.tanzu.vmware.com
  - system.antrea.io
  resources:
  - controllerinfos
  - agentinfos
  verbs:
  - get
- apiGroups:
  - system.antrea.tanzu.vmware.com
  - system.antrea.io
  resources:
  - supportbundles
  verbs:
  - get
  - post
- apiGroups:
  - system.antrea.tanzu.vmware.com
  - system.antrea.io
  resources:
  - supportbundles/download
  verbs:
  - get
- nonResourceURLs:
  - /agentinfo
  - /addressgroups
  - /appliedtogroups
  - /loglevel
  - /networkpolicies
  - /ovsflows
  - /ovstracing
  - /podinterfaces
  - /featuregates
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
  verbs:
  - get
  - watch
  - list
  - patch
- apiGroups:
  - ""
  resources:
  - pods
  verbs:
  - get
  - watch
  - list
  - patch
- apiGroups:
  - ""
  resources:
  - endpoints
  - services
  verbs:
  - get
  - watch
  - list
- apiGroups:
  - discovery.k8s.io
  resources:
  - endpointslices
  verbs:
  - get
  - watch
  - list
- apiGroups:
  - clusterinformation.antrea.tanzu.vmware.com
  - crd.antrea.io
  resources:
  - antreaagentinfos
  verbs:
  - get
  - create
  - update
  - delete
- apiGroups:
  - controlplane.antrea.tanzu.vmware.com
  - controlplane.antrea.io
  resources:
  - networkpolicies
  - appliedtogroups
  - addressgroups
  verbs:
  - get
  - watch
  - list
- apiGroups:
  - controlplane.antrea.io
  resources:
  - egressgroups
  verbs:
  - get
  - watch
  - list
- apiGroups:
  - controlplane.antrea.tanzu.vmware.com
  - controlplane.antrea.io
  resources:
  - nodestatssummaries
  verbs:
  - create
- apiGroups:
  - controlplane.antrea.tanzu.vmware.com
  - controlplane.antrea.io
  resources:
  - networkpolicies/status
  verbs:
  - create
  - get
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
- apiGroups:
  - ""
  resourceNames:
  - extension-apiserver-authentication
  resources:
  - configmaps
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resourceNames:
  - antrea-ca
  resources:
  - configmaps
  verbs:
  - get
  - watch
  - list
- apiGroups:
  - ops.antrea.tanzu.vmware.com
  - crd.antrea.io
  resources:
  - traceflows
  - traceflows/status
  verbs:
  - get
  - watch
  - list
  - update
  - patch
  - create
  - delete
- apiGroups:
  - crd.antrea.io
  resources:
  - egresses
  verbs:
  - get
  - watch
  - list
- apiGroups:
  - crd.antrea.io
  resources:
  - egresses/status
  verbs:
  - update
- apiGroups:
  - crd.antrea.io
  resources:
  - externalippools
  verbs:
  - get
  - watch
  - list
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: antrea-cluster-identity-reader
rules:
- apiGroups:
  - ""
  resourceNames:
  - antrea-cluster-identity
  resources:
  - configmaps
  verbs:
  - get
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
  - services
  - configmaps
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
- apiGroups:
  - apiextensions.k8s.io
  resources:
  - customresourcedefinitions
  verbs:
  - get
  - update
- apiGroups:
  - ""
  resourceNames:
  - extension-apiserver-authentication
  resources:
  - configmaps
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resourceNames:
  - antrea-ca
  - antrea-cluster-identity
  resources:
  - configmaps
  verbs:
  - get
  - update
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - create
- apiGroups:
  - apiregistration.k8s.io
  resourceNames:
  - v1alpha1.stats.antrea.tanzu.vmware.com
  - v1beta1.system.antrea.tanzu.vmware.com
  - v1beta2.controlplane.antrea.tanzu.vmware.com
  - v1alpha1.stats.antrea.io
  - v1beta1.system.antrea.io
  - v1beta2.controlplane.antrea.io
  resources:
  - apiservices
  verbs:
  - get
  - update
- apiGroups:
  - apiregistration.k8s.io
  resourceNames:
  - v1beta1.networking.antrea.tanzu.vmware.com
  - v1beta1.controlplane.antrea.tanzu.vmware.com
  resources:
  - apiservices
  verbs:
  - delete
- apiGroups:
  - admissionregistration.k8s.io
  resourceNames:
  - crdmutator.antrea.tanzu.vmware.com
  - crdvalidator.antrea.tanzu.vmware.com
  - labelsmutator.antrea.io
  - crdmutator.antrea.io
  - crdvalidator.antrea.io
  resources:
  - mutatingwebhookconfigurations
  - validatingwebhookconfigurations
  verbs:
  - get
  - update
- apiGroups:
  - crd.antrea.io
  resources:
  - antreacontrollerinfos
  verbs:
  - get
  - create
  - update
  - delete
- apiGroups:
  - crd.antrea.io
  resources:
  - antreaagentinfos
  verbs:
  - list
  - delete
- apiGroups:
  - crd.antrea.io
  resources:
  - clusternetworkpolicies
  - networkpolicies
  verbs:
  - get
  - watch
  - list
  - update
  - patch
  - create
  - delete
- apiGroups:
  - crd.antrea.io
  resources:
  - clusternetworkpolicies/status
  - networkpolicies/status
  verbs:
  - update
- apiGroups:
  - crd.antrea.io
  resources:
  - tiers
  verbs:
  - get
  - watch
  - list
  - update
  - patch
  - create
  - delete
- apiGroups:
  - crd.antrea.io
  resources:
  - traceflows
  - traceflows/status
  verbs:
  - get
  - watch
  - list
  - update
  - patch
  - create
  - delete
- apiGroups:
  - crd.antrea.io
  resources:
  - externalentities
  - clustergroups
  verbs:
  - get
  - watch
  - list
  - update
  - patch
  - create
  - delete
- apiGroups:
  - crd.antrea.io
  resources:
  - clustergroups/status
  verbs:
  - update
- apiGroups:
  - crd.antrea.io
  resources:
  - egresses
  verbs:
  - get
  - watch
  - list
  - update
  - patch
- apiGroups:
  - crd.antrea.io
  resources:
  - externalippools
  verbs:
  - get
  - watch
  - list
- apiGroups:
  - crd.antrea.io
  resources:
  - externalippools/status
  verbs:
  - update
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
  - security.antrea.tanzu.vmware.com
  resources:
  - clusternetworkpolicies
  - networkpolicies
  verbs:
  - get
  - watch
  - list
  - update
  - patch
  - create
  - delete
- apiGroups:
  - security.antrea.tanzu.vmware.com
  resources:
  - clusternetworkpolicies/status
  - networkpolicies/status
  verbs:
  - update
- apiGroups:
  - security.antrea.tanzu.vmware.com
  resources:
  - tiers
  verbs:
  - get
  - watch
  - list
  - update
  - patch
  - create
  - delete
- apiGroups:
  - ops.antrea.tanzu.vmware.com
  resources:
  - traceflows
  - traceflows/status
  verbs:
  - get
  - watch
  - list
  - update
  - patch
  - create
  - delete
- apiGroups:
  - core.antrea.tanzu.vmware.com
  resources:
  - externalentities
  - clustergroups
  verbs:
  - get
  - watch
  - list
  - update
  - patch
  - create
  - delete
- apiGroups:
  - core.antrea.tanzu.vmware.com
  resources:
  - clustergroups/status
  verbs:
  - update
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
apiVersion: rbac.authorization.k8s.io/v1
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
apiVersion: rbac.authorization.k8s.io/v1
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
    # FeatureGates is a map of feature names to bools that enable or disable experimental features.
    featureGates:
    # Enable AntreaProxy which provides ServiceLB for in-cluster Services in antrea-agent.
    # It should be enabled on Windows, otherwise NetworkPolicy will not take effect on
    # Service traffic.
      AntreaProxy: true

    # Enable EndpointSlice support in AntreaProxy. Don't enable this feature unless that EndpointSlice
    # API version v1beta1 is supported and set as enabled in Kubernetes. If AntreaProxy is not enabled,
    # this flag will not take effect.
    #  EndpointSlice: false

    # Enable traceflow which provides packet tracing feature to diagnose network issue.
    #  Traceflow: true

    # Enable NodePortLocal feature to make the pods reachable externally through NodePort
    #  NodePortLocal: false

    # Enable Antrea ClusterNetworkPolicy feature to complement K8s NetworkPolicy for cluster admins
    # to define security policies which apply to the entire cluster, and Antrea NetworkPolicy
    # feature that supports priorities, rule actions and externalEntities in the future.
    #  AntreaPolicy: true

    # Enable flowexporter which exports polled conntrack connections as IPFIX flow records from each
    # agent to a configured collector.
    #  FlowExporter: false

    # Enable collecting and exposing NetworkPolicy statistics.
    #  NetworkPolicyStats: true

    # Enable controlling SNAT IPs of Pod egress traffic.
    #  Egress: false

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
    #hostGateway: antrea-gw0

    # Determines how traffic is encapsulated. It has the following options:
    # encap(default):    Inter-node Pod traffic is always encapsulated and Pod to external network
    #                    traffic is SNAT'd.
    # noEncap:           Inter-node Pod traffic is not encapsulated; Pod to external network traffic is
    #                    SNAT'd if noSNAT is not set to true. Underlying network must be capable of
    #                    supporting Pod traffic across IP subnets.
    # hybrid:            noEncap if source and destination Nodes are on the same subnet, otherwise encap.
    # networkPolicyOnly: Antrea enforces NetworkPolicy only, and utilizes CNI chaining and delegates Pod
    #                    IPAM and connectivity to the primary CNI.
    #
    trafficEncapMode: {{ContainerConfig "trafficEncapMode"}}

    # Whether or not to SNAT (using the Node IP) the egress traffic from a Pod to the external network.
    # This option is for the noEncap traffic mode only, and the default value is false. In the noEncap
    # mode, if the cluster's Pod CIDR is reachable from the external network, then the Pod traffic to
    # the external network needs not be SNAT'd. In the networkPolicyOnly mode, antrea-agent never
    # performs SNAT and this option will be ignored; for other modes it must be set to false.
    #noSNAT: false

    # Tunnel protocols used for encapsulating traffic across Nodes. If WireGuard is enabled in trafficEncryptionMode,
    # this option will not take effect. Supported values:
    # - geneve (default)
    # - vxlan
    # - gre
    # - stt
    #tunnelType: geneve

    # Determines how tunnel traffic is encrypted. Currently encryption only works with encap mode.
    # It has the following options:
    # - none (default):  Inter-node Pod traffic will not be encrypted.
    # - ipsec:           Enable IPSec (ESP) encryption for Pod traffic across Nodes. Antrea uses
    #                    Preshared Key (PSK) for IKE authentication. When IPSec tunnel is enabled,
    #                    the PSK value must be passed to Antrea Agent through an environment
    #                    variable: ANTREA_IPSEC_PSK.
    # - wireGuard:       Enable WireGuard for tunnel traffic encryption.
    #trafficEncryptionMode: none

    # Default MTU to use for the host gateway interface and the network interface of each Pod.
    # If omitted, antrea-agent will discover the MTU of the Node's primary interface and
    # also adjust MTU to accommodate for tunnel encapsulation overhead (if applicable).
    #defaultMTU: 0

    # wireGuard specifies WireGuard related configurations.
    wireGuard:
    #  The port for WireGuard to receive traffic.
    #  port: 51820

    # ClusterIP CIDR range for Services. It's required when AntreaProxy is not enabled, and should be
    # set to the same value as the one specified by --service-cluster-ip-range for kube-apiserver. When
    # AntreaProxy is enabled, this parameter is not needed and will be ignored if provided.
    serviceCIDR: {{ContainerConfig "serviceCidr"}}

    # ClusterIP CIDR range for IPv6 Services. It's required when using kube-proxy to provide IPv6 Service in a Dual-Stack
    # cluster or an IPv6 only cluster. The value should be the same as the configuration for kube-apiserver specified by
    # --service-cluster-ip-range. When AntreaProxy is enabled, this parameter is not needed.
    # No default value for this field.
    #serviceCIDRv6:

    # The port for the antrea-agent APIServer to serve on.
    # Note that if it's set to another value, the ` + "`" + `containerPort` + "`" + ` of the ` + "`" + `api` + "`" + ` port of the
    # ` + "`" + `antrea-agent` + "`" + ` container must be set to the same value.
    #apiPort: 10350

    # Enable metrics exposure via Prometheus. Initializes Prometheus metrics listener.
    #enablePrometheusMetrics: true

    # Provide the IPFIX collector address as a string with format <HOST>:[<PORT>][:<PROTO>].
    # HOST can either be the DNS name or the IP of the Flow Collector. For example,
    # "flow-aggregator.flow-aggregator.svc" can be provided as DNS name to connect
    # to the Antrea Flow Aggregator service. If IP, it can be either IPv4 or IPv6.
    # However, IPv6 address should be wrapped with [].
    # If PORT is empty, we default to 4739, the standard IPFIX port.
    # If no PROTO is given, we consider "tls" as default. We support "tls", "tcp" and
    # "udp" protocols. "tls" is used for securing communication between flow exporter and
    # flow aggregator.
    #flowCollectorAddr: "flow-aggregator.flow-aggregator.svc:4739:tls"

    # Provide flow poll interval as a duration string. This determines how often the
    # flow exporter dumps connections from the conntrack module. Flow poll interval
    # should be greater than or equal to 1s (one second).
    # Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
    #flowPollInterval: "5s"

    # Provide the active flow export timeout, which is the timeout after which a flow
    # record is sent to the collector for active flows. Thus, for flows with a continuous
    # stream of packets, a flow record will be exported to the collector once the elapsed
    # time since the last export event is equal to the value of this timeout.
    # Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
    #activeFlowExportTimeout: "30s"

    # Provide the idle flow export timeout, which is the timeout after which a flow
    # record is sent to the collector for idle flows. A flow is considered idle if no
    # packet matching this flow has been observed since the last export event.
    # Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
    #idleFlowExportTimeout: "15s"

    # Provide the port range used by NodePortLocal. When the NodePortLocal feature is enabled, a port from that range will be assigned
    # whenever a Pod's container defines a specific port to be exposed (each container can define a list of ports as pod.spec.containers[].ports),
    # and all Node traffic directed to that port will be forwarded to the Pod.
    #nplPortRange: 61000-62000

    # Provide the address of Kubernetes apiserver, to override any value provided in kubeconfig or InClusterConfig.
    # Defaults to "". It must be a host string, a host:port pair, or a URL to the base of the apiserver.
    #kubeAPIServerOverride: ""

    # Comma-separated list of Cipher Suites. If omitted, the default Go Cipher Suites will be used.
    # https://golang.org/pkg/crypto/tls/#pkg-constants
    # Note that TLS1.3 Cipher Suites cannot be added to the list. But the apiserver will always
    # prefer TLS1.3 Cipher Suites whenever possible.
    #tlsCipherSuites:

    # TLS min version from: VersionTLS10, VersionTLS11, VersionTLS12, VersionTLS13.
    #tlsMinVersion:

    # The name of the interface on Node which is used for tunneling or routing the traffic across Nodes.
    # If there are multiple IP addresses configured on the interface, the first one is used.
    # The interface configured with Node IP is used if this parameter is not set.
    #transportInterface:
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
            },
            {
                "type": "bandwidth",
                "capabilities": {"bandwidth": true}
            }
        ]
    }
  antrea-controller.conf: |
    # FeatureGates is a map of feature names to bools that enable or disable experimental features.
    featureGates:
    # Enable traceflow which provides packet tracing feature to diagnose network issue.
    #  Traceflow: true

    # Enable Antrea ClusterNetworkPolicy feature to complement K8s NetworkPolicy for cluster admins
    # to define security policies which apply to the entire cluster, and Antrea NetworkPolicy
    # feature that supports priorities, rule actions and externalEntities in the future.
    #  AntreaPolicy: true

    # Enable collecting and exposing NetworkPolicy statistics.
    #  NetworkPolicyStats: true

    # Enable controlling SNAT IPs of Pod egress traffic.
    #  Egress: false

    # The port for the antrea-controller APIServer to serve on.
    # Note that if it's set to another value, the ` + "`" + `containerPort` + "`" + ` of the ` + "`" + `api` + "`" + ` port of the
    # ` + "`" + `antrea-controller` + "`" + ` container must be set to the same value.
    #apiPort: 10349

    # Enable metrics exposure via Prometheus. Initializes Prometheus metrics listener.
    #enablePrometheusMetrics: true

    # Indicates whether to use auto-generated self-signed TLS certificate.
    # If false, A Secret named "antrea-controller-tls" must be provided with the following keys:
    #   ca.crt: <CA certificate>
    #   tls.crt: <TLS certificate>
    #   tls.key: <TLS private key>
    # And the Secret must be mounted to directory "/var/run/antrea/antrea-controller-tls" of the
    # antrea-controller container.
    #selfSignedCert: true

    # Comma-separated list of Cipher Suites. If omitted, the default Go Cipher Suites will be used.
    # https://golang.org/pkg/crypto/tls/#pkg-constants
    # Note that TLS1.3 Cipher Suites cannot be added to the list. But the apiserver will always
    # prefer TLS1.3 Cipher Suites whenever possible.
    #tlsCipherSuites:

    # TLS min version from: VersionTLS10, VersionTLS11, VersionTLS12, VersionTLS13.
    #tlsMinVersion:

    # If Antrea is upgraded from version <= v0.13 and legacy CRDs are used, this option should be
    # enabled, otherwise the CRDs created with the legacy API groups will not take any effect and
    # work as expected. When the mirroring is enabled, if a legacy CRD is created with legacy API
    # groups, mirroring-controller will create a new CRD with the Spec and Labels from the legacy
    # CRD. Afterwards, the modification of Spec and Label in legacy CRD will be synchronized to new
    # CRD automatically. In addition, the modification of Status in new CRD will also be synchronized
    # to legacy CRD automatically. If a legacy CRD is deleted, the corresponding new CRD will be deleted.
    # Note that: to decouple a new CRD from the corresponding legacy CRD, the legacy CRD should be
    # annotated with "crd.antrea.io/stop-mirror". Afterwards, updates to the legacy CRDs will no
    # longer be reflected in the new CRD, and all CRUD operations should be done through the new
    # API groups. After adding the annotation, legacy CRDs can be deleted safely without impacting
    # new CRDs.
    #legacyCRDMirroring: true
kind: ConfigMap
metadata:
  annotations: {}
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "EnsureExists"
  name: antrea-config-mc8h75hbgg
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
        - --log_dir=/var/log/antrea
        - --alsologtostderr
        - --log_file_max_size=100
        - --log_file_max_num=4
        - --v=0
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
        - name: SERVICEACCOUNT_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.serviceAccountName
        - name: ANTREA_CONFIG_MAP_NAME
          value: antrea-config-mc8h75hbgg
        image: {{ContainerImage "antrea-controller"}}
        livenessProbe:
          failureThreshold: 5
          httpGet:
            host: localhost
            path: /livez
            port: api
            scheme: HTTPS
          periodSeconds: 10
          timeoutSeconds: 5
        name: antrea-controller
        ports:
        - containerPort: 10349
          name: api
          protocol: TCP
        readinessProbe:
          failureThreshold: 5
          httpGet:
            host: localhost
            path: /readyz
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
        - mountPath: /var/run/antrea/antrea-controller-tls
          name: antrea-controller-tls
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
          name: antrea-config-mc8h75hbgg
        name: antrea-config
      - name: antrea-controller-tls
        secret:
          defaultMode: 256
          optional: true
          secretName: antrea-controller-tls
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
  name: v1alpha1.stats.antrea.io
spec:
  group: stats.antrea.io
  groupPriorityMinimum: 100
  service:
    name: antrea
    namespace: kube-system
  version: v1alpha1
  versionPriority: 100
---
apiVersion: apiregistration.k8s.io/v1
kind: APIService
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: v1alpha1.stats.antrea.tanzu.vmware.com
spec:
  group: stats.antrea.tanzu.vmware.com
  groupPriorityMinimum: 100
  service:
    name: antrea
    namespace: kube-system
  version: v1alpha1
  versionPriority: 100
---
apiVersion: apiregistration.k8s.io/v1
kind: APIService
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: v1beta1.system.antrea.io
spec:
  group: system.antrea.io
  groupPriorityMinimum: 100
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
  name: v1beta2.controlplane.antrea.io
spec:
  group: controlplane.antrea.io
  groupPriorityMinimum: 100
  service:
    name: antrea
    namespace: kube-system
  version: v1beta2
  versionPriority: 100
---
apiVersion: apiregistration.k8s.io/v1
kind: APIService
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: v1beta2.controlplane.antrea.tanzu.vmware.com
spec:
  group: controlplane.antrea.tanzu.vmware.com
  groupPriorityMinimum: 100
  service:
    name: antrea
    namespace: kube-system
  version: v1beta2
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
      annotations:
        kubectl.kubernetes.io/default-container: antrea-agent
      labels:
        app: antrea
        component: antrea-agent
    spec:
      containers:
      - args:
        - --config
        - /etc/antrea/antrea-agent.conf
        - --logtostderr=false
        - --log_dir=/var/log/antrea
        - --alsologtostderr
        - --log_file_max_size=100
        - --log_file_max_num=4
        - --v=0
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
          failureThreshold: 8
          httpGet:
            host: localhost
            path: /readyz
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
      - args:
        - --log_file_max_size=100
        - --log_file_max_num=4
        command:
        - start_ovs
        image: {{ContainerImage "antrea-ovs"}}
        livenessProbe:
          exec:
            command:
            - /bin/sh
            - -c
            - timeout 10 container_liveness_probe ovs
          failureThreshold: 5
          initialDelaySeconds: 5
          periodSeconds: 10
          timeoutSeconds: 10
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
      dnsPolicy: ClusterFirstWithHostNet
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
        - mountPath: /var/run/antrea
          name: host-var-run-antrea
      nodeSelector:
        kubernetes.io/os: linux
      priorityClassName: system-node-critical
      serviceAccountName: antrea-agent
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      - effect: NoSchedule
        operator: Exists
      - effect: NoExecute
        operator: Exists
      volumes:
      - configMap:
          name: antrea-config-mc8h75hbgg
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
          path: /run/xtables.lock
          type: FileOrCreate
        name: xtables-lock
  updateStrategy:
    type: RollingUpdate
---
apiVersion: admissionregistration.k8s.io/v1
kind: MutatingWebhookConfiguration
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: crdmutator.antrea.io
webhooks:
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: antrea
      namespace: kube-system
      path: /mutate/acnp
  name: acnpmutator.antrea.io
  rules:
  - apiGroups:
    - crd.antrea.io
    apiVersions:
    - v1alpha1
    operations:
    - CREATE
    - UPDATE
    resources:
    - clusternetworkpolicies
    scope: Cluster
  sideEffects: None
  timeoutSeconds: 5
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: antrea
      namespace: kube-system
      path: /mutate/anp
  name: anpmutator.antrea.io
  rules:
  - apiGroups:
    - crd.antrea.io
    apiVersions:
    - v1alpha1
    operations:
    - CREATE
    - UPDATE
    resources:
    - networkpolicies
    scope: Namespaced
  sideEffects: None
  timeoutSeconds: 5
---
apiVersion: admissionregistration.k8s.io/v1
kind: MutatingWebhookConfiguration
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: crdmutator.antrea.tanzu.vmware.com
webhooks:
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: antrea
      namespace: kube-system
      path: /mutate/acnp
  name: acnpmutator.antrea.tanzu.vmware.com
  rules:
  - apiGroups:
    - security.antrea.tanzu.vmware.com
    apiVersions:
    - v1alpha1
    operations:
    - CREATE
    - UPDATE
    resources:
    - clusternetworkpolicies
    scope: Cluster
  sideEffects: None
  timeoutSeconds: 5
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: antrea
      namespace: kube-system
      path: /mutate/anp
  name: anpmutator.antrea.tanzu.vmware.com
  rules:
  - apiGroups:
    - security.antrea.tanzu.vmware.com
    apiVersions:
    - v1alpha1
    operations:
    - CREATE
    - UPDATE
    resources:
    - networkpolicies
    scope: Namespaced
  sideEffects: None
  timeoutSeconds: 5
---
apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingWebhookConfiguration
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: crdvalidator.antrea.io
webhooks:
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: antrea
      namespace: kube-system
      path: /validate/tier
  name: tiervalidator.antrea.io
  rules:
  - apiGroups:
    - crd.antrea.io
    apiVersions:
    - v1alpha1
    operations:
    - CREATE
    - UPDATE
    - DELETE
    resources:
    - tiers
    scope: Cluster
  sideEffects: None
  timeoutSeconds: 5
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: antrea
      namespace: kube-system
      path: /validate/acnp
  name: acnpvalidator.antrea.io
  rules:
  - apiGroups:
    - crd.antrea.io
    apiVersions:
    - v1alpha1
    operations:
    - CREATE
    - UPDATE
    resources:
    - clusternetworkpolicies
    scope: Cluster
  sideEffects: None
  timeoutSeconds: 5
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: antrea
      namespace: kube-system
      path: /validate/anp
  name: anpvalidator.antrea.io
  rules:
  - apiGroups:
    - crd.antrea.io
    apiVersions:
    - v1alpha1
    operations:
    - CREATE
    - UPDATE
    resources:
    - networkpolicies
    scope: Namespaced
  sideEffects: None
  timeoutSeconds: 5
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: antrea
      namespace: kube-system
      path: /validate/clustergroup
  name: clustergroupvalidator.antrea.io
  rules:
  - apiGroups:
    - crd.antrea.io
    apiVersions:
    - v1alpha3
    - v1alpha2
    operations:
    - CREATE
    - UPDATE
    resources:
    - clustergroups
    scope: Cluster
  sideEffects: None
  timeoutSeconds: 5
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: antrea
      namespace: kube-system
      path: /validate/externalippool
  name: externalippoolvalidator.antrea.io
  rules:
  - apiGroups:
    - crd.antrea.io
    apiVersions:
    - v1alpha2
    operations:
    - UPDATE
    resources:
    - externalippools
    scope: Cluster
  sideEffects: None
  timeoutSeconds: 5
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: antrea
      namespace: kube-system
      path: /validate/egress
  name: egressvalidator.antrea.io
  rules:
  - apiGroups:
    - crd.antrea.io
    apiVersions:
    - v1alpha2
    operations:
    - CREATE
    - UPDATE
    resources:
    - egresses
    scope: Cluster
  sideEffects: None
  timeoutSeconds: 5
---
apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingWebhookConfiguration
metadata:
  labels:
    app: antrea
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: crdvalidator.antrea.tanzu.vmware.com
webhooks:
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: antrea
      namespace: kube-system
      path: /validate/tier
  name: tiervalidator.antrea.tanzu.vmware.com
  rules:
  - apiGroups:
    - security.antrea.tanzu.vmware.com
    apiVersions:
    - v1alpha1
    operations:
    - CREATE
    - UPDATE
    - DELETE
    resources:
    - tiers
    scope: Cluster
  sideEffects: None
  timeoutSeconds: 5
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: antrea
      namespace: kube-system
      path: /validate/acnp
  name: acnpvalidator.antrea.tanzu.vmware.com
  rules:
  - apiGroups:
    - security.antrea.tanzu.vmware.com
    apiVersions:
    - v1alpha1
    operations:
    - CREATE
    - UPDATE
    resources:
    - clusternetworkpolicies
    scope: Cluster
  sideEffects: None
  timeoutSeconds: 5
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: antrea
      namespace: kube-system
      path: /validate/anp
  name: anpvalidator.antrea.tanzu.vmware.com
  rules:
  - apiGroups:
    - security.antrea.tanzu.vmware.com
    apiVersions:
    - v1alpha1
    operations:
    - CREATE
    - UPDATE
    resources:
    - networkpolicies
    scope: Namespaced
  sideEffects: None
  timeoutSeconds: 5
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: antrea
      namespace: kube-system
      path: /validate/clustergroup
  name: clustergroupvalidator.antrea.tanzu.vmware.com
  rules:
  - apiGroups:
    - core.antrea.tanzu.vmware.com
    apiVersions:
    - v1alpha2
    operations:
    - CREATE
    - UPDATE
    - DELETE
    resources:
    - clustergroups
    scope: Cluster
  sideEffects: None
  timeoutSeconds: 5
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

var _k8sAddonsAuditPolicyYaml = []byte(`apiVersion: audit.k8s.io/v1
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
apiVersion: rbac.authorization.k8s.io/v1
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
allowVolumeExpansion: true
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
allowVolumeExpansion: true
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
allowVolumeExpansion: true
  {{- if HasAvailabilityZones}}
volumeBindingMode: WaitForFirstConsumer
allowedTopologies:
- matchLabelExpressions:
  - key: topology.disk.csi.azure.com/zone
    values: {{GetZones}}
  {{else}}
volumeBindingMode: Immediate
  {{- end}}
  {{- if not IsAzureStackCloud}}
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
allowVolumeExpansion: true
volumeBindingMode: Immediate
  {{- end}}
{{else}}
  {{- if NeedsStorageAccountStorageClasses}}
---
apiVersion: storage.k8s.io/v1{{- if not (IsKubernetesVersionGe "1.22.0-alpha.2")}}beta1{{end}}
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
apiVersion: storage.k8s.io/v1{{- if not (IsKubernetesVersionGe "1.22.0-alpha.2")}}beta1{{end}}
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
apiVersion: storage.k8s.io/v1{{- if not (IsKubernetesVersionGe "1.22.0-alpha.2")}}beta1{{end}}
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
apiVersion: storage.k8s.io/v1{{- if not (IsKubernetesVersionGe "1.22.0-alpha.2")}}beta1{{end}}
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
apiVersion: storage.k8s.io/v1{{- if not (IsKubernetesVersionGe "1.22.0-alpha.2")}}beta1{{end}}
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
apiVersion: storage.k8s.io/v1{{- if not (IsKubernetesVersionGe "1.22.0-alpha.2")}}beta1{{end}}
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
apiVersion: apps/v1
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
    rollingUpdate:
      maxUnavailable: 50%
  template:
    metadata:
      labels:
        k8s-app: azure-npm
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ''
{{- if IsKubernetesVersionGe "1.17.0"}}
        cluster-autoscaler.kubernetes.io/daemonset-pod: "true"
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
    controller-gen.kubebuilder.io/version: v0.3.0
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
          description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          description: ConfigSpec defines the desired state of Config
          properties:
            match:
              description: Configuration for namespace exclusion
              items:
                properties:
                  excludedNamespaces:
                    items:
                      type: string
                    type: array
                  processes:
                    items:
                      type: string
                    type: array
                type: object
              type: array
            readiness:
              description: Configuration for readiness tracker
              properties:
                statsEnabled:
                  type: boolean
              type: object
            sync:
              description: Configuration for syncing k8s objects
              properties:
                syncOnly:
                  description: If non-empty, only entries on this list will be replicated into OPA
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
                  description: List of requests to trace. Both "user" and "kinds" must be specified
                  items:
                    properties:
                      dump:
                        description: Also dump the state of OPA with the trace. Set to ` + "`" + `All` + "`" + ` to dump everything.
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
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.3.0
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: constraintpodstatuses.status.gatekeeper.sh
spec:
  group: status.gatekeeper.sh
  names:
    kind: ConstraintPodStatus
    listKind: ConstraintPodStatusList
    plural: constraintpodstatuses
    singular: constraintpodstatus
  scope: Namespaced
  validation:
    openAPIV3Schema:
      description: ConstraintPodStatus is the Schema for the constraintpodstatuses API
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        status:
          description: ConstraintPodStatusStatus defines the observed state of ConstraintPodStatus
          properties:
            constraintUID:
              description: Storing the constraint UID allows us to detect drift, such as when a constraint has been recreated after its CRD was deleted out from under it, interrupting the watch
              type: string
            enforced:
              type: boolean
            errors:
              items:
                description: Error represents a single error caught while adding a constraint to OPA
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
              type: string
            observedGeneration:
              format: int64
              type: integer
            operations:
              items:
                type: string
              type: array
          type: object
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
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.3.0
  creationTimestamp: null
  labels:
    gatekeeper.sh/system: "yes"
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: constrainttemplatepodstatuses.status.gatekeeper.sh
spec:
  group: status.gatekeeper.sh
  names:
    kind: ConstraintTemplatePodStatus
    listKind: ConstraintTemplatePodStatusList
    plural: constrainttemplatepodstatuses
    singular: constrainttemplatepodstatus
  scope: Namespaced
  validation:
    openAPIV3Schema:
      description: ConstraintTemplatePodStatus is the Schema for the constrainttemplatepodstatuses API
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        status:
          description: ConstraintTemplatePodStatusStatus defines the observed state of ConstraintTemplatePodStatus
          properties:
            errors:
              items:
                description: CreateCRDError represents a single error caught during parsing, compiling, etc.
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
              description: 'Important: Run "make" to regenerate code after modifying this file'
              type: string
            observedGeneration:
              format: int64
              type: integer
            operations:
              items:
                type: string
              type: array
            templateUID:
              description: UID is a type that holds unique ID values, including UUIDs.  Because we don't ONLY use UUIDs, this is an alias to string.  Being a type captures intent and helps make sure that UIDs and names do not get conflated.
              type: string
          type: object
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
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  creationTimestamp: null
  labels:
    controller-tools.k8s.io: "1.0"
    gatekeeper.sh/system: "yes"
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
          description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
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
  labels:
    gatekeeper.sh/system: "yes"
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: gatekeeper-admin
  namespace: gatekeeper-system
---
apiVersion: policy/v1beta1
kind: PodSecurityPolicy
metadata:
  annotations:
    seccomp.security.alpha.kubernetes.io/allowedProfileNames: '*'
  labels:
    gatekeeper.sh/system: "yes"
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: gatekeeper-admin
spec:
  allowPrivilegeEscalation: false
  fsGroup:
    ranges:
    - max: 65535
      min: 1
    rule: MustRunAs
  requiredDropCapabilities:
  - ALL
  runAsUser:
    rule: MustRunAsNonRoot
  seLinux:
    rule: RunAsAny
  supplementalGroups:
    ranges:
    - max: 65535
      min: 1
    rule: MustRunAs
  volumes:
  - configMap
  - projected
  - secret
  - downwardAPI
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
  - events
  verbs:
  - create
  - patch
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
  - policy
  resourceNames:
  - gatekeeper-admin
  resources:
  - podsecuritypolicies
  verbs:
  - use
- apiGroups:
  - status.gatekeeper.sh
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
  - constrainttemplates/finalizers
  verbs:
  - delete
  - get
  - patch
  - update
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
    gatekeeper.sh/operation: webhook
    gatekeeper.sh/system: "yes"
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    control-plane: audit-controller
    gatekeeper.sh/operation: audit
    gatekeeper.sh/system: "yes"
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: gatekeeper-audit
  namespace: gatekeeper-system
spec:
  replicas: 1
  selector:
    matchLabels:
      control-plane: audit-controller
      gatekeeper.sh/operation: audit
      gatekeeper.sh/system: "yes"
  template:
    metadata:
      annotations:
        container.seccomp.security.alpha.kubernetes.io/manager: runtime/default
      labels:
        control-plane: audit-controller
        gatekeeper.sh/operation: audit
        gatekeeper.sh/system: "yes"
    spec:
      containers:
      - args:
        - --operation=audit
        - --operation=status
        - --logtostderr
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
      nodeSelector:
        kubernetes.io/os: linux
      serviceAccountName: gatekeeper-admin
      terminationGracePeriodSeconds: 60
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    control-plane: controller-manager
    gatekeeper.sh/operation: webhook
    gatekeeper.sh/system: "yes"
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
  name: gatekeeper-controller-manager
  namespace: gatekeeper-system
spec:
  replicas: 3
  selector:
    matchLabels:
      control-plane: controller-manager
      gatekeeper.sh/operation: webhook
      gatekeeper.sh/system: "yes"
  template:
    metadata:
      annotations:
        container.seccomp.security.alpha.kubernetes.io/manager: runtime/default
      labels:
        control-plane: controller-manager
        gatekeeper.sh/operation: webhook
        gatekeeper.sh/system: "yes"
    spec:
      affinity:
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - podAffinityTerm:
              labelSelector:
                matchExpressions:
                - key: gatekeeper.sh/operation
                  operator: In
                  values:
                  - webhook
              topologyKey: kubernetes.io/hostname
            weight: 100
      containers:
      - args:
        - --port=8443
        - --logtostderr
        - --exempt-namespace=gatekeeper-system
        - --operation=webhook
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
      nodeSelector:
        kubernetes.io/os: linux
      serviceAccountName: gatekeeper-admin
      terminationGracePeriodSeconds: 60
      volumes:
      - name: cert
        secret:
          defaultMode: 420
          secretName: gatekeeper-webhook-server-cert
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
  timeoutSeconds: 3
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
  timeoutSeconds: 3
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
  resources: ["constrainttemplates", "constrainttemplates/finalizers"]
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
      nodeSelector:
        kubernetes.io/os: linux
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
        - name: DATAPLANE_ENDPOINT
          value: https://gov-prod-policy-data.trafficmanager.net
        - name: FULL_SCAN_EXCLUSION_LIST
          value: "kube-system,gatekeeper-system"
        - name: WEBHOOK_EXCLUSION_LIST
          value: "kube-system,gatekeeper-system"
        - name: ACS_CREDENTIAL_LOCATION
          value: /etc/acs/azure.json
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
        livenessProbe:
          httpGet:
            path: /healthz
            port: 9090
          initialDelaySeconds: 5
        readinessProbe:
          httpGet:
            path: /readyz
            port: 9090
          initialDelaySeconds: 5
        ports:
        - containerPort: 9090
          name: healthz
          protocol: TCP
      volumes:
      - hostPath:
          path: /etc/kubernetes/azure.json
          type: File
        name: acs-credential
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
    verbs: ["get", "list", "watch", "update", "patch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["volumeattachments/status"]
    verbs: ["get", "list", "watch", "update", "patch"]
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
    resources: ["volumesnapshotcontents/status"]
    verbs: ["update"]
  - apiGroups: ["apiextensions.k8s.io"]
    resources: ["customresourcedefinitions"]
    verbs: ["create", "list", "watch", "delete"]
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
  - apiGroups: [""]
    resources: ["pods"]
    verbs: ["get", "list", "watch"]
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
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 50%
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
        - key: "node.kubernetes.io/os"
          operator: "Exists"
          effect: "NoSchedule"
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
            - "--v=2"
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
            - "--v=2"
            - "--csi-address=$(CSI_ENDPOINT)"
            - "--kubelet-registration-path=$(DRIVER_REG_SOCK_PATH)"
          livenessProbe:
            exec:
              command:
                - /csi-node-driver-registrar.exe
                - --kubelet-registration-path=$(DRIVER_REG_SOCK_PATH)
                - --mode=kubelet-registration-probe
            initialDelaySeconds: 60
            timeoutSeconds: 30
          env:
            - name: CSI_ENDPOINT
              value: unix://C:\\csi\\csi.sock
            - name: DRIVER_REG_SOCK_PATH
              value: C:\\var\\lib\\kubelet\\plugins\\disk.csi.azure.com\\csi.sock
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
            - "--user-agent-suffix=aks-engine"
          ports:
            - containerPort: 29603
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
          volumeMounts:
            - name: kubelet-dir
              mountPath: "C:\\var\\lib\\kubelet"
            - name: plugin-dir
              mountPath: C:\csi
            - name: azure-config
              mountPath: C:\k
            - name: csi-proxy-fs-pipe-v1
              mountPath: \\.\pipe\csi-proxy-filesystem-v1
            - name: csi-proxy-disk-pipe-v1
              mountPath: \\.\pipe\csi-proxy-disk-v1
            - name: csi-proxy-volume-pipe-v1
              mountPath: \\.\pipe\csi-proxy-volume-v1
            # these paths are still included for compatibility, they're used
            # only if the node has still the beta version of the CSI proxy
            - name: csi-proxy-fs-pipe-v1beta1
              mountPath: \\.\pipe\csi-proxy-filesystem-v1beta1
            - name: csi-proxy-disk-pipe-v1beta2
              mountPath: \\.\pipe\csi-proxy-disk-v1beta2
            - name: csi-proxy-volume-pipe-v1beta2
              mountPath: \\.\pipe\csi-proxy-volume-v1beta2
          resources:
            limits:
              cpu: {{ContainerCPULimits "azuredisk-csi"}}
              memory: {{ContainerMemLimits "azuredisk-csi"}}
            requests:
              cpu: {{ContainerCPUReqs "azuredisk-csi"}}
              memory: {{ContainerMemReqs "azuredisk-csi"}}
      volumes:
        - name: csi-proxy-fs-pipe-v1
          hostPath:
            path: \\.\pipe\csi-proxy-filesystem-v1
        - name: csi-proxy-disk-pipe-v1
          hostPath:
            path: \\.\pipe\csi-proxy-disk-v1
        - name: csi-proxy-volume-pipe-v1
          hostPath:
            path: \\.\pipe\csi-proxy-volume-v1
        # these paths are still included for compatibility, they're used
        # only if the node has still the beta version of the CSI proxy
        - name: csi-proxy-fs-pipe-v1beta1
          hostPath:
            path: \\.\pipe\csi-proxy-filesystem-v1beta1
        - name: csi-proxy-disk-pipe-v1beta2
          hostPath:
            path: \\.\pipe\csi-proxy-disk-v1beta2
        - name: csi-proxy-volume-pipe-v1beta2
          hostPath:
            path: \\.\pipe\csi-proxy-volume-v1beta2
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
            type: DirectoryOrCreate
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
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 50%
  selector:
    matchLabels:
      app: csi-azuredisk-node
  template:
    metadata:
      labels:
        app: csi-azuredisk-node
    spec:
      hostNetwork: true
      dnsPolicy: Default
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
            - --probe-timeout=3s
            - --health-port=29603
            - --v=2
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
            - --v=2
          livenessProbe:
            exec:
              command:
                - /csi-node-driver-registrar
                - --kubelet-registration-path=$(DRIVER_REG_SOCK_PATH)
                - --mode=kubelet-registration-probe
            initialDelaySeconds: 30
            timeoutSeconds: 15
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
            - "--enable-perf-optimization=true"
            - "--user-agent-suffix=aks-engine"
          ports:
            - containerPort: 29603
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
            type: DirectoryOrCreate
          name: azure-cred
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
        kubernetes.io/role: master
      priorityClassName: system-cluster-critical
      tolerations:
        - key: "node-role.kubernetes.io/master"
          operator: "Exists"
          effect: "NoSchedule"
        - key: "node-role.kubernetes.io/controlplane"
          operator: "Exists"
          effect: "NoSchedule"
      containers:
        - name: csi-provisioner
          image: {{ContainerImage "csi-provisioner"}}
          args:
            - "--feature-gates=Topology=true"
            - "--csi-address=$(ADDRESS)"
            - "--v=2"
            - "--timeout=15s"
            - "--leader-election"
            - "--worker-threads=40"
            - "--extra-create-metadata=true"
            - "--strict-topology=true"
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
            - "-v=2"
            - "-csi-address=$(ADDRESS)"
            - "-timeout=600s"
            - "-leader-election"
            - "-worker-threads=500"
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
            - "-v=2"
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
        - name: csi-resizer
          image: {{ContainerImage "csi-resizer"}}
          args:
            - "-csi-address=$(ADDRESS)"
            - "-v=2"
            - "-leader-election"
            - '-handle-volume-inuse-error=false'
            - '-timeout=60s'
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
        - name: liveness-probe
          image: {{ContainerImage "livenessprobe"}}
          args:
            - --csi-address=/csi/csi.sock
            - --probe-timeout=3s
            - --health-port=29602
            - --v=2
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
            - "--metrics-address=0.0.0.0:29604"
            - "--user-agent-suffix=aks-engine"
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
            type: DirectoryOrCreate
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
        kubernetes.io/role: master
      priorityClassName: system-cluster-critical
      tolerations:
        - key: "node-role.kubernetes.io/master"
          operator: "Equal"
          value: "true"
          effect: "NoSchedule"
        - key: "node-role.kubernetes.io/controlplane"
          operator: "Equal"
          value: "true"
          effect: "NoSchedule"
      containers:
        - name: csi-snapshot-controller
          image: {{ContainerImage "csi-snapshot-controller"}}
          args:
            - "--v=2"
            - "--leader-election=false"
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
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.4.0
    api-approved.kubernetes.io: "https://github.com/kubernetes-csi/external-snapshotter/pull/419"
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
  creationTimestamp: null
  name: volumesnapshots.snapshot.storage.k8s.io
spec:
  group: snapshot.storage.k8s.io
  names:
    kind: VolumeSnapshot
    listKind: VolumeSnapshotList
    plural: volumesnapshots
    singular: volumesnapshot
  scope: Namespaced
  versions:
  - additionalPrinterColumns:
    - description: Indicates if the snapshot is ready to be used to restore a volume.
      jsonPath: .status.readyToUse
      name: ReadyToUse
      type: boolean
    - description: If a new snapshot needs to be created, this contains the name of the source PVC from which this snapshot was (or will be) created.
      jsonPath: .spec.source.persistentVolumeClaimName
      name: SourcePVC
      type: string
    - description: If a snapshot already exists, this contains the name of the existing VolumeSnapshotContent object representing the existing snapshot.
      jsonPath: .spec.source.volumeSnapshotContentName
      name: SourceSnapshotContent
      type: string
    - description: Represents the minimum size of volume required to rehydrate from this snapshot.
      jsonPath: .status.restoreSize
      name: RestoreSize
      type: string
    - description: The name of the VolumeSnapshotClass requested by the VolumeSnapshot.
      jsonPath: .spec.volumeSnapshotClassName
      name: SnapshotClass
      type: string
    - description: Name of the VolumeSnapshotContent object to which the VolumeSnapshot object intends to bind to. Please note that verification of binding actually requires checking both VolumeSnapshot and VolumeSnapshotContent to ensure both are pointing at each other. Binding MUST be verified prior to usage of this object.
      jsonPath: .status.boundVolumeSnapshotContentName
      name: SnapshotContent
      type: string
    - description: Timestamp when the point-in-time snapshot was taken by the underlying storage system.
      jsonPath: .status.creationTime
      name: CreationTime
      type: date
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1
    schema:
      openAPIV3Schema:
        description: VolumeSnapshot is a user's request for either creating a point-in-time snapshot of a persistent volume, or binding to a pre-existing snapshot.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          spec:
            description: 'spec defines the desired characteristics of a snapshot requested by a user. More info: https://kubernetes.io/docs/concepts/storage/volume-snapshots#volumesnapshots Required.'
            properties:
              source:
                description: source specifies where a snapshot will be created from. This field is immutable after creation. Required.
                properties:
                  persistentVolumeClaimName:
                    description: persistentVolumeClaimName specifies the name of the PersistentVolumeClaim object representing the volume from which a snapshot should be created. This PVC is assumed to be in the same namespace as the VolumeSnapshot object. This field should be set if the snapshot does not exists, and needs to be created. This field is immutable.
                    type: string
                  volumeSnapshotContentName:
                    description: volumeSnapshotContentName specifies the name of a pre-existing VolumeSnapshotContent object representing an existing volume snapshot. This field should be set if the snapshot already exists and only needs a representation in Kubernetes. This field is immutable.
                    type: string
                type: object
                oneOf:
                - required: ["persistentVolumeClaimName"]
                - required: ["volumeSnapshotContentName"]
              volumeSnapshotClassName:
                description: 'VolumeSnapshotClassName is the name of the VolumeSnapshotClass requested by the VolumeSnapshot. VolumeSnapshotClassName may be left nil to indicate that the default SnapshotClass should be used. A given cluster may have multiple default Volume SnapshotClasses: one default per CSI Driver. If a VolumeSnapshot does not specify a SnapshotClass, VolumeSnapshotSource will be checked to figure out what the associated CSI Driver is, and the default VolumeSnapshotClass associated with that CSI Driver will be used. If more than one VolumeSnapshotClass exist for a given CSI Driver and more than one have been marked as default, CreateSnapshot will fail and generate an event. Empty string is not allowed for this field.'
                type: string
            required:
            - source
            type: object
          status:
            description: status represents the current information of a snapshot. Consumers must verify binding between VolumeSnapshot and VolumeSnapshotContent objects is successful (by validating that both VolumeSnapshot and VolumeSnapshotContent point at each other) before using this object.
            properties:
              boundVolumeSnapshotContentName:
                description: 'boundVolumeSnapshotContentName is the name of the VolumeSnapshotContent object to which this VolumeSnapshot object intends to bind to. If not specified, it indicates that the VolumeSnapshot object has not been successfully bound to a VolumeSnapshotContent object yet. NOTE: To avoid possible security issues, consumers must verify binding between VolumeSnapshot and VolumeSnapshotContent objects is successful (by validating that both VolumeSnapshot and VolumeSnapshotContent point at each other) before using this object.'
                type: string
              creationTime:
                description: creationTime is the timestamp when the point-in-time snapshot is taken by the underlying storage system. In dynamic snapshot creation case, this field will be filled in by the snapshot controller with the "creation_time" value returned from CSI "CreateSnapshot" gRPC call. For a pre-existing snapshot, this field will be filled with the "creation_time" value returned from the CSI "ListSnapshots" gRPC call if the driver supports it. If not specified, it may indicate that the creation time of the snapshot is unknown.
                format: date-time
                type: string
              error:
                description: error is the last observed error during snapshot creation, if any. This field could be helpful to upper level controllers(i.e., application controller) to decide whether they should continue on waiting for the snapshot to be created based on the type of error reported. The snapshot controller will keep retrying when an error occurrs during the snapshot creation. Upon success, this error field will be cleared.
                properties:
                  message:
                    description: 'message is a string detailing the encountered error during snapshot creation if specified. NOTE: message may be logged, and it should not contain sensitive information.'
                    type: string
                  time:
                    description: time is the timestamp when the error was encountered.
                    format: date-time
                    type: string
                type: object
              readyToUse:
                description: readyToUse indicates if the snapshot is ready to be used to restore a volume. In dynamic snapshot creation case, this field will be filled in by the snapshot controller with the "ready_to_use" value returned from CSI "CreateSnapshot" gRPC call. For a pre-existing snapshot, this field will be filled with the "ready_to_use" value returned from the CSI "ListSnapshots" gRPC call if the driver supports it, otherwise, this field will be set to "True". If not specified, it means the readiness of a snapshot is unknown.
                type: boolean
              restoreSize:
                type: string
                description: restoreSize represents the minimum size of volume required to create a volume from this snapshot. In dynamic snapshot creation case, this field will be filled in by the snapshot controller with the "size_bytes" value returned from CSI "CreateSnapshot" gRPC call. For a pre-existing snapshot, this field will be filled with the "size_bytes" value returned from the CSI "ListSnapshots" gRPC call if the driver supports it. When restoring a volume from this snapshot, the size of the volume MUST NOT be smaller than the restoreSize if it is specified, otherwise the restoration will fail. If not specified, it indicates that the size is unknown.
                pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                x-kubernetes-int-or-string: true
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
  - additionalPrinterColumns:
    - description: Indicates if the snapshot is ready to be used to restore a volume.
      jsonPath: .status.readyToUse
      name: ReadyToUse
      type: boolean
    - description: If a new snapshot needs to be created, this contains the name of the source PVC from which this snapshot was (or will be) created.
      jsonPath: .spec.source.persistentVolumeClaimName
      name: SourcePVC
      type: string
    - description: If a snapshot already exists, this contains the name of the existing VolumeSnapshotContent object representing the existing snapshot.
      jsonPath: .spec.source.volumeSnapshotContentName
      name: SourceSnapshotContent
      type: string
    - description: Represents the minimum size of volume required to rehydrate from this snapshot.
      jsonPath: .status.restoreSize
      name: RestoreSize
      type: string
    - description: The name of the VolumeSnapshotClass requested by the VolumeSnapshot.
      jsonPath: .spec.volumeSnapshotClassName
      name: SnapshotClass
      type: string
    - description: Name of the VolumeSnapshotContent object to which the VolumeSnapshot object intends to bind to. Please note that verification of binding actually requires checking both VolumeSnapshot and VolumeSnapshotContent to ensure both are pointing at each other. Binding MUST be verified prior to usage of this object.
      jsonPath: .status.boundVolumeSnapshotContentName
      name: SnapshotContent
      type: string
    - description: Timestamp when the point-in-time snapshot was taken by the underlying storage system.
      jsonPath: .status.creationTime
      name: CreationTime
      type: date
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1beta1
    # This indicates the v1beta1 version of the custom resource is deprecated.
    # API requests to this version receive a warning in the server response.
    deprecated: true
    # This overrides the default warning returned to clients making v1beta1 API requests.
    deprecationWarning: "snapshot.storage.k8s.io/v1beta1 VolumeSnapshot is deprecated; use snapshot.storage.k8s.io/v1 VolumeSnapshot"
    schema:
      openAPIV3Schema:
        description: VolumeSnapshot is a user's request for either creating a point-in-time snapshot of a persistent volume, or binding to a pre-existing snapshot.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          spec:
            description: 'spec defines the desired characteristics of a snapshot requested by a user. More info: https://kubernetes.io/docs/concepts/storage/volume-snapshots#volumesnapshots Required.'
            properties:
              source:
                description: source specifies where a snapshot will be created from. This field is immutable after creation. Required.
                properties:
                  persistentVolumeClaimName:
                    description: persistentVolumeClaimName specifies the name of the PersistentVolumeClaim object representing the volume from which a snapshot should be created. This PVC is assumed to be in the same namespace as the VolumeSnapshot object. This field should be set if the snapshot does not exists, and needs to be created. This field is immutable.
                    type: string
                  volumeSnapshotContentName:
                    description: volumeSnapshotContentName specifies the name of a pre-existing VolumeSnapshotContent object representing an existing volume snapshot. This field should be set if the snapshot already exists and only needs a representation in Kubernetes. This field is immutable.
                    type: string
                type: object
              volumeSnapshotClassName:
                description: 'VolumeSnapshotClassName is the name of the VolumeSnapshotClass requested by the VolumeSnapshot. VolumeSnapshotClassName may be left nil to indicate that the default SnapshotClass should be used. A given cluster may have multiple default Volume SnapshotClasses: one default per CSI Driver. If a VolumeSnapshot does not specify a SnapshotClass, VolumeSnapshotSource will be checked to figure out what the associated CSI Driver is, and the default VolumeSnapshotClass associated with that CSI Driver will be used. If more than one VolumeSnapshotClass exist for a given CSI Driver and more than one have been marked as default, CreateSnapshot will fail and generate an event. Empty string is not allowed for this field.'
                type: string
            required:
            - source
            type: object
          status:
            description: status represents the current information of a snapshot. Consumers must verify binding between VolumeSnapshot and VolumeSnapshotContent objects is successful (by validating that both VolumeSnapshot and VolumeSnapshotContent point at each other) before using this object.
            properties:
              boundVolumeSnapshotContentName:
                description: 'boundVolumeSnapshotContentName is the name of the VolumeSnapshotContent object to which this VolumeSnapshot object intends to bind to. If not specified, it indicates that the VolumeSnapshot object has not been successfully bound to a VolumeSnapshotContent object yet. NOTE: To avoid possible security issues, consumers must verify binding between VolumeSnapshot and VolumeSnapshotContent objects is successful (by validating that both VolumeSnapshot and VolumeSnapshotContent point at each other) before using this object.'
                type: string
              creationTime:
                description: creationTime is the timestamp when the point-in-time snapshot is taken by the underlying storage system. In dynamic snapshot creation case, this field will be filled in by the snapshot controller with the "creation_time" value returned from CSI "CreateSnapshot" gRPC call. For a pre-existing snapshot, this field will be filled with the "creation_time" value returned from the CSI "ListSnapshots" gRPC call if the driver supports it. If not specified, it may indicate that the creation time of the snapshot is unknown.
                format: date-time
                type: string
              error:
                description: error is the last observed error during snapshot creation, if any. This field could be helpful to upper level controllers(i.e., application controller) to decide whether they should continue on waiting for the snapshot to be created based on the type of error reported. The snapshot controller will keep retrying when an error occurrs during the snapshot creation. Upon success, this error field will be cleared.
                properties:
                  message:
                    description: 'message is a string detailing the encountered error during snapshot creation if specified. NOTE: message may be logged, and it should not contain sensitive information.'
                    type: string
                  time:
                    description: time is the timestamp when the error was encountered.
                    format: date-time
                    type: string
                type: object
              readyToUse:
                description: readyToUse indicates if the snapshot is ready to be used to restore a volume. In dynamic snapshot creation case, this field will be filled in by the snapshot controller with the "ready_to_use" value returned from CSI "CreateSnapshot" gRPC call. For a pre-existing snapshot, this field will be filled with the "ready_to_use" value returned from the CSI "ListSnapshots" gRPC call if the driver supports it, otherwise, this field will be set to "True". If not specified, it means the readiness of a snapshot is unknown.
                type: boolean
              restoreSize:
                type: string
                description: restoreSize represents the minimum size of volume required to create a volume from this snapshot. In dynamic snapshot creation case, this field will be filled in by the snapshot controller with the "size_bytes" value returned from CSI "CreateSnapshot" gRPC call. For a pre-existing snapshot, this field will be filled with the "size_bytes" value returned from the CSI "ListSnapshots" gRPC call if the driver supports it. When restoring a volume from this snapshot, the size of the volume MUST NOT be smaller than the restoreSize if it is specified, otherwise the restoration will fail. If not specified, it indicates that the size is unknown.
                pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                x-kubernetes-int-or-string: true
            type: object
        required:
        - spec
        type: object
    served: true
    storage: false
    subresources:
      status: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
---
# Source: azuredisk-csi-driver/templates/crd-csi-snapshot.yaml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.4.0
    api-approved.kubernetes.io: "https://github.com/kubernetes-csi/external-snapshotter/pull/419"
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
  creationTimestamp: null
  name: volumesnapshotclasses.snapshot.storage.k8s.io
spec:
  group: snapshot.storage.k8s.io
  names:
    kind: VolumeSnapshotClass
    listKind: VolumeSnapshotClassList
    plural: volumesnapshotclasses
    singular: volumesnapshotclass
  scope: Cluster
  versions:
  - additionalPrinterColumns:
    - jsonPath: .driver
      name: Driver
      type: string
    - description: Determines whether a VolumeSnapshotContent created through the VolumeSnapshotClass should be deleted when its bound VolumeSnapshot is deleted.
      jsonPath: .deletionPolicy
      name: DeletionPolicy
      type: string
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1
    schema:
      openAPIV3Schema:
        description: VolumeSnapshotClass specifies parameters that a underlying storage system uses when creating a volume snapshot. A specific VolumeSnapshotClass is used by specifying its name in a VolumeSnapshot object. VolumeSnapshotClasses are non-namespaced
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          deletionPolicy:
            description: deletionPolicy determines whether a VolumeSnapshotContent created through the VolumeSnapshotClass should be deleted when its bound VolumeSnapshot is deleted. Supported values are "Retain" and "Delete". "Retain" means that the VolumeSnapshotContent and its physical snapshot on underlying storage system are kept. "Delete" means that the VolumeSnapshotContent and its physical snapshot on underlying storage system are deleted. Required.
            enum:
            - Delete
            - Retain
            type: string
          driver:
            description: driver is the name of the storage driver that handles this VolumeSnapshotClass. Required.
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          parameters:
            additionalProperties:
              type: string
            description: parameters is a key-value map with storage driver specific parameters for creating snapshots. These values are opaque to Kubernetes.
            type: object
        required:
        - deletionPolicy
        - driver
        type: object
    served: true
    storage: true
    subresources: {}
  - additionalPrinterColumns:
    - jsonPath: .driver
      name: Driver
      type: string
    - description: Determines whether a VolumeSnapshotContent created through the VolumeSnapshotClass should be deleted when its bound VolumeSnapshot is deleted.
      jsonPath: .deletionPolicy
      name: DeletionPolicy
      type: string
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1beta1
    # This indicates the v1beta1 version of the custom resource is deprecated.
    # API requests to this version receive a warning in the server response.
    deprecated: true
    # This overrides the default warning returned to clients making v1beta1 API requests.
    deprecationWarning: "snapshot.storage.k8s.io/v1beta1 VolumeSnapshotClass is deprecated; use snapshot.storage.k8s.io/v1 VolumeSnapshotClass"
    schema:
      openAPIV3Schema:
        description: VolumeSnapshotClass specifies parameters that a underlying storage system uses when creating a volume snapshot. A specific VolumeSnapshotClass is used by specifying its name in a VolumeSnapshot object. VolumeSnapshotClasses are non-namespaced
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          deletionPolicy:
            description: deletionPolicy determines whether a VolumeSnapshotContent created through the VolumeSnapshotClass should be deleted when its bound VolumeSnapshot is deleted. Supported values are "Retain" and "Delete". "Retain" means that the VolumeSnapshotContent and its physical snapshot on underlying storage system are kept. "Delete" means that the VolumeSnapshotContent and its physical snapshot on underlying storage system are deleted. Required.
            enum:
            - Delete
            - Retain
            type: string
          driver:
            description: driver is the name of the storage driver that handles this VolumeSnapshotClass. Required.
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          parameters:
            additionalProperties:
              type: string
            description: parameters is a key-value map with storage driver specific parameters for creating snapshots. These values are opaque to Kubernetes.
            type: object
        required:
        - deletionPolicy
        - driver
        type: object
    served: true
    storage: false
    subresources: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
---
# Source: azuredisk-csi-driver/templates/crd-csi-snapshot.yaml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.4.0
    api-approved.kubernetes.io: "https://github.com/kubernetes-csi/external-snapshotter/pull/419"
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
  creationTimestamp: null
  name: volumesnapshotcontents.snapshot.storage.k8s.io
spec:
  group: snapshot.storage.k8s.io
  names:
    kind: VolumeSnapshotContent
    listKind: VolumeSnapshotContentList
    plural: volumesnapshotcontents
    singular: volumesnapshotcontent
  scope: Cluster
  versions:
  - additionalPrinterColumns:
    - description: Indicates if the snapshot is ready to be used to restore a volume.
      jsonPath: .status.readyToUse
      name: ReadyToUse
      type: boolean
    - description: Represents the complete size of the snapshot in bytes
      jsonPath: .status.restoreSize
      name: RestoreSize
      type: integer
    - description: Determines whether this VolumeSnapshotContent and its physical snapshot on the underlying storage system should be deleted when its bound VolumeSnapshot is deleted.
      jsonPath: .spec.deletionPolicy
      name: DeletionPolicy
      type: string
    - description: Name of the CSI driver used to create the physical snapshot on the underlying storage system.
      jsonPath: .spec.driver
      name: Driver
      type: string
    - description: Name of the VolumeSnapshotClass to which this snapshot belongs.
      jsonPath: .spec.volumeSnapshotClassName
      name: VolumeSnapshotClass
      type: string
    - description: Name of the VolumeSnapshot object to which this VolumeSnapshotContent object is bound.
      jsonPath: .spec.volumeSnapshotRef.name
      name: VolumeSnapshot
      type: string
    - description: Namespace of the VolumeSnapshot object to which this VolumeSnapshotContent object is bound.
      jsonPath: .spec.volumeSnapshotRef.namespace
      name: VolumeSnapshotNamespace
      type: string
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1
    schema:
      openAPIV3Schema:
        description: VolumeSnapshotContent represents the actual "on-disk" snapshot object in the underlying storage system
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          spec:
            description: spec defines properties of a VolumeSnapshotContent created by the underlying storage system. Required.
            properties:
              deletionPolicy:
                description: deletionPolicy determines whether this VolumeSnapshotContent and its physical snapshot on the underlying storage system should be deleted when its bound VolumeSnapshot is deleted. Supported values are "Retain" and "Delete". "Retain" means that the VolumeSnapshotContent and its physical snapshot on underlying storage system are kept. "Delete" means that the VolumeSnapshotContent and its physical snapshot on underlying storage system are deleted. For dynamically provisioned snapshots, this field will automatically be filled in by the CSI snapshotter sidecar with the "DeletionPolicy" field defined in the corresponding VolumeSnapshotClass. For pre-existing snapshots, users MUST specify this field when creating the  VolumeSnapshotContent object. Required.
                enum:
                - Delete
                - Retain
                type: string
              driver:
                description: driver is the name of the CSI driver used to create the physical snapshot on the underlying storage system. This MUST be the same as the name returned by the CSI GetPluginName() call for that driver. Required.
                type: string
              source:
                description: source specifies whether the snapshot is (or should be) dynamically provisioned or already exists, and just requires a Kubernetes object representation. This field is immutable after creation. Required.
                properties:
                  snapshotHandle:
                    description: snapshotHandle specifies the CSI "snapshot_id" of a pre-existing snapshot on the underlying storage system for which a Kubernetes object representation was (or should be) created. This field is immutable.
                    type: string
                  volumeHandle:
                    description: volumeHandle specifies the CSI "volume_id" of the volume from which a snapshot should be dynamically taken from. This field is immutable.
                    type: string
                type: object
                oneOf:
                - required: ["snapshotHandle"]
                - required: ["volumeHandle"]
              volumeSnapshotClassName:
                description: name of the VolumeSnapshotClass from which this snapshot was (or will be) created. Note that after provisioning, the VolumeSnapshotClass may be deleted or recreated with different set of values, and as such, should not be referenced post-snapshot creation.
                type: string
              volumeSnapshotRef:
                description: volumeSnapshotRef specifies the VolumeSnapshot object to which this VolumeSnapshotContent object is bound. VolumeSnapshot.Spec.VolumeSnapshotContentName field must reference to this VolumeSnapshotContent's name for the bidirectional binding to be valid. For a pre-existing VolumeSnapshotContent object, name and namespace of the VolumeSnapshot object MUST be provided for binding to happen. This field is immutable after creation. Required.
                properties:
                  apiVersion:
                    description: API version of the referent.
                    type: string
                  fieldPath:
                    description: 'If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: "spec.containers{name}" (where "name" refers to the name of the container that triggered the event) or if no container name is specified "spec.containers[2]" (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.'
                    type: string
                  kind:
                    description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                    type: string
                  name:
                    description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                    type: string
                  namespace:
                    description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                    type: string
                  resourceVersion:
                    description: 'Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
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
                description: creationTime is the timestamp when the point-in-time snapshot is taken by the underlying storage system. In dynamic snapshot creation case, this field will be filled in by the CSI snapshotter sidecar with the "creation_time" value returned from CSI "CreateSnapshot" gRPC call. For a pre-existing snapshot, this field will be filled with the "creation_time" value returned from the CSI "ListSnapshots" gRPC call if the driver supports it. If not specified, it indicates the creation time is unknown. The format of this field is a Unix nanoseconds time encoded as an int64. On Unix, the command ` + "`" + `date +%s%N` + "`" + ` returns the current time in nanoseconds since 1970-01-01 00:00:00 UTC.
                format: int64
                type: integer
              error:
                description: error is the last observed error during snapshot creation, if any. Upon success after retry, this error field will be cleared.
                properties:
                  message:
                    description: 'message is a string detailing the encountered error during snapshot creation if specified. NOTE: message may be logged, and it should not contain sensitive information.'
                    type: string
                  time:
                    description: time is the timestamp when the error was encountered.
                    format: date-time
                    type: string
                type: object
              readyToUse:
                description: readyToUse indicates if a snapshot is ready to be used to restore a volume. In dynamic snapshot creation case, this field will be filled in by the CSI snapshotter sidecar with the "ready_to_use" value returned from CSI "CreateSnapshot" gRPC call. For a pre-existing snapshot, this field will be filled with the "ready_to_use" value returned from the CSI "ListSnapshots" gRPC call if the driver supports it, otherwise, this field will be set to "True". If not specified, it means the readiness of a snapshot is unknown.
                type: boolean
              restoreSize:
                description: restoreSize represents the complete size of the snapshot in bytes. In dynamic snapshot creation case, this field will be filled in by the CSI snapshotter sidecar with the "size_bytes" value returned from CSI "CreateSnapshot" gRPC call. For a pre-existing snapshot, this field will be filled with the "size_bytes" value returned from the CSI "ListSnapshots" gRPC call if the driver supports it. When restoring a volume from this snapshot, the size of the volume MUST NOT be smaller than the restoreSize if it is specified, otherwise the restoration will fail. If not specified, it indicates that the size is unknown.
                format: int64
                minimum: 0
                type: integer
              snapshotHandle:
                description: snapshotHandle is the CSI "snapshot_id" of a snapshot on the underlying storage system. If not specified, it indicates that dynamic snapshot creation has either failed or it is still in progress.
                type: string
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
  - additionalPrinterColumns:
    - description: Indicates if the snapshot is ready to be used to restore a volume.
      jsonPath: .status.readyToUse
      name: ReadyToUse
      type: boolean
    - description: Represents the complete size of the snapshot in bytes
      jsonPath: .status.restoreSize
      name: RestoreSize
      type: integer
    - description: Determines whether this VolumeSnapshotContent and its physical snapshot on the underlying storage system should be deleted when its bound VolumeSnapshot is deleted.
      jsonPath: .spec.deletionPolicy
      name: DeletionPolicy
      type: string
    - description: Name of the CSI driver used to create the physical snapshot on the underlying storage system.
      jsonPath: .spec.driver
      name: Driver
      type: string
    - description: Name of the VolumeSnapshotClass to which this snapshot belongs.
      jsonPath: .spec.volumeSnapshotClassName
      name: VolumeSnapshotClass
      type: string
    - description: Name of the VolumeSnapshot object to which this VolumeSnapshotContent object is bound.
      jsonPath: .spec.volumeSnapshotRef.name
      name: VolumeSnapshot
      type: string
    - description: Namespace of the VolumeSnapshot object to which this VolumeSnapshotContent object is bound.
      jsonPath: .spec.volumeSnapshotRef.namespace
      name: VolumeSnapshotNamespace
      type: string
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1beta1
    # This indicates the v1beta1 version of the custom resource is deprecated.
    # API requests to this version receive a warning in the server response.
    deprecated: true
    # This overrides the default warning returned to clients making v1beta1 API requests.
    deprecationWarning: "snapshot.storage.k8s.io/v1beta1 VolumeSnapshotContent is deprecated; use snapshot.storage.k8s.io/v1 VolumeSnapshotContent"
    schema:
      openAPIV3Schema:
        description: VolumeSnapshotContent represents the actual "on-disk" snapshot object in the underlying storage system
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          spec:
            description: spec defines properties of a VolumeSnapshotContent created by the underlying storage system. Required.
            properties:
              deletionPolicy:
                description: deletionPolicy determines whether this VolumeSnapshotContent and its physical snapshot on the underlying storage system should be deleted when its bound VolumeSnapshot is deleted. Supported values are "Retain" and "Delete". "Retain" means that the VolumeSnapshotContent and its physical snapshot on underlying storage system are kept. "Delete" means that the VolumeSnapshotContent and its physical snapshot on underlying storage system are deleted. For dynamically provisioned snapshots, this field will automatically be filled in by the CSI snapshotter sidecar with the "DeletionPolicy" field defined in the corresponding VolumeSnapshotClass. For pre-existing snapshots, users MUST specify this field when creating the  VolumeSnapshotContent object. Required.
                enum:
                - Delete
                - Retain
                type: string
              driver:
                description: driver is the name of the CSI driver used to create the physical snapshot on the underlying storage system. This MUST be the same as the name returned by the CSI GetPluginName() call for that driver. Required.
                type: string
              source:
                description: source specifies whether the snapshot is (or should be) dynamically provisioned or already exists, and just requires a Kubernetes object representation. This field is immutable after creation. Required.
                properties:
                  snapshotHandle:
                    description: snapshotHandle specifies the CSI "snapshot_id" of a pre-existing snapshot on the underlying storage system for which a Kubernetes object representation was (or should be) created. This field is immutable.
                    type: string
                  volumeHandle:
                    description: volumeHandle specifies the CSI "volume_id" of the volume from which a snapshot should be dynamically taken from. This field is immutable.
                    type: string
                type: object
              volumeSnapshotClassName:
                description: name of the VolumeSnapshotClass from which this snapshot was (or will be) created. Note that after provisioning, the VolumeSnapshotClass may be deleted or recreated with different set of values, and as such, should not be referenced post-snapshot creation.
                type: string
              volumeSnapshotRef:
                description: volumeSnapshotRef specifies the VolumeSnapshot object to which this VolumeSnapshotContent object is bound. VolumeSnapshot.Spec.VolumeSnapshotContentName field must reference to this VolumeSnapshotContent's name for the bidirectional binding to be valid. For a pre-existing VolumeSnapshotContent object, name and namespace of the VolumeSnapshot object MUST be provided for binding to happen. This field is immutable after creation. Required.
                properties:
                  apiVersion:
                    description: API version of the referent.
                    type: string
                  fieldPath:
                    description: 'If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: "spec.containers{name}" (where "name" refers to the name of the container that triggered the event) or if no container name is specified "spec.containers[2]" (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.'
                    type: string
                  kind:
                    description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                    type: string
                  name:
                    description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                    type: string
                  namespace:
                    description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                    type: string
                  resourceVersion:
                    description: 'Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
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
                description: creationTime is the timestamp when the point-in-time snapshot is taken by the underlying storage system. In dynamic snapshot creation case, this field will be filled in by the CSI snapshotter sidecar with the "creation_time" value returned from CSI "CreateSnapshot" gRPC call. For a pre-existing snapshot, this field will be filled with the "creation_time" value returned from the CSI "ListSnapshots" gRPC call if the driver supports it. If not specified, it indicates the creation time is unknown. The format of this field is a Unix nanoseconds time encoded as an int64. On Unix, the command ` + "`" + `date +%s%N` + "`" + ` returns the current time in nanoseconds since 1970-01-01 00:00:00 UTC.
                format: int64
                type: integer
              error:
                description: error is the last observed error during snapshot creation, if any. Upon success after retry, this error field will be cleared.
                properties:
                  message:
                    description: 'message is a string detailing the encountered error during snapshot creation if specified. NOTE: message may be logged, and it should not contain sensitive information.'
                    type: string
                  time:
                    description: time is the timestamp when the error was encountered.
                    format: date-time
                    type: string
                type: object
              readyToUse:
                description: readyToUse indicates if a snapshot is ready to be used to restore a volume. In dynamic snapshot creation case, this field will be filled in by the CSI snapshotter sidecar with the "ready_to_use" value returned from CSI "CreateSnapshot" gRPC call. For a pre-existing snapshot, this field will be filled with the "ready_to_use" value returned from the CSI "ListSnapshots" gRPC call if the driver supports it, otherwise, this field will be set to "True". If not specified, it means the readiness of a snapshot is unknown.
                type: boolean
              restoreSize:
                description: restoreSize represents the complete size of the snapshot in bytes. In dynamic snapshot creation case, this field will be filled in by the CSI snapshotter sidecar with the "size_bytes" value returned from CSI "CreateSnapshot" gRPC call. For a pre-existing snapshot, this field will be filled with the "size_bytes" value returned from the CSI "ListSnapshots" gRPC call if the driver supports it. When restoring a volume from this snapshot, the size of the volume MUST NOT be smaller than the restoreSize if it is specified, otherwise the restoration will fail. If not specified, it indicates that the size is unknown.
                format: int64
                minimum: 0
                type: integer
              snapshotHandle:
                description: snapshotHandle is the CSI "snapshot_id" of a snapshot on the underlying storage system. If not specified, it indicates that dynamic snapshot creation has either failed or it is still in progress.
                type: string
            type: object
        required:
        - spec
        type: object
    served: true
    storage: false
    subresources:
      status: {}
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
apiVersion: storage.k8s.io/v1
kind: CSIDriver
metadata:
  name: disk.csi.azure.com
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  attachRequired: true
  podInfoOnMount: false
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
# Source: azurefile-csi-driver/templates/serviceaccount-csi-azurefile-node.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: csi-azurefile-node-sa
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
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
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshots"]
    verbs: ["get", "list"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshotcontents"]
    verbs: ["get", "list"]
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
    verbs: ["get", "list", "watch", "update", "patch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["volumeattachments/status"]
    verbs: ["get", "list", "watch", "update", "patch"]
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
  - apiGroups: ["apiextensions.k8s.io"]
    resources: ["customresourcedefinitions"]
    verbs: ["create", "list", "watch", "delete"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshotcontents/status"]
    verbs: ["update"]
  - apiGroups: ["coordination.k8s.io"]
    resources: ["leases"]
    verbs: ["get", "watch", "list", "delete", "update", "create", "patch"]
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
---
# Source: azurefile-csi-driver/templates/rbac-csi-azurefile-controller.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: csi-azurefile-controller-secret-role
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
  - apiGroups: [""]
    resources: ["secrets"]
    verbs: ["get", "list", "create"]
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
  name: csi-azurefile-controller-secret-binding
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
  name: csi-azurefile-controller-secret-role
---
# Source: azurefile-csi-driver/templates/rbac-csi-azurefile-node.yaml
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: csi-azurefile-node-secret-role
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
  - apiGroups: [""]
    resources: ["secrets"]
    verbs: ["get", "list"]
---
# Source: azurefile-csi-driver/templates/rbac-csi-azurefile-controller.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: csi-azurefile-node-secret-binding
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
  name: csi-azurefile-node-secret-role
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
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 50%
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
        - key: "node.kubernetes.io/os"
          operator: "Exists"
          effect: "NoSchedule"
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
            - "--v=2"
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
            - --v=2
            - --csi-address=$(CSI_ENDPOINT)
            - --kubelet-registration-path=$(DRIVER_REG_SOCK_PATH)
          livenessProbe:
            exec:
              command:
                - /csi-node-driver-registrar.exe
                - --kubelet-registration-path=$(DRIVER_REG_SOCK_PATH)
                - --mode=kubelet-registration-probe
            initialDelaySeconds: 60
            timeoutSeconds: 30
          env:
            - name: CSI_ENDPOINT
              value: unix://C:\\csi\\csi.sock
            - name: DRIVER_REG_SOCK_PATH
              value: C:\\var\\lib\\kubelet\\plugins\\file.csi.azure.com\\csi.sock
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
            - --v=5
            - --endpoint=$(CSI_ENDPOINT)
            - --nodeid=$(KUBE_NODE_NAME)
            - --kubeconfig=C:\\k\\config
            - --metrics-address=0.0.0.0:29615
            - --user-agent-suffix=aks-engine
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
          volumeMounts:
            - name: kubelet-dir
              mountPath: "C:\\var\\lib\\kubelet"
            - name: plugin-dir
              mountPath: C:\csi
            - name: azure-config
              mountPath: C:\k
            - name: csi-proxy-fs-pipe-v1
              mountPath: \\.\pipe\csi-proxy-filesystem-v1
            - name: csi-proxy-smb-pipe-v1
              mountPath: \\.\pipe\csi-proxy-smb-v1
            # these paths are still included for compatibility, they're used
            # only if the node has still the beta version of the CSI proxy
            - name: csi-proxy-fs-pipe-v1beta1
              mountPath: \\.\pipe\csi-proxy-filesystem-v1beta1
            - name: csi-proxy-smb-pipe-v1beta1
              mountPath: \\.\pipe\csi-proxy-smb-v1beta1
          resources:
            limits:
              cpu: {{ContainerCPULimits "azurefile-csi"}}
              memory: {{ContainerMemLimits "azurefile-csi"}}
            requests:
              cpu: {{ContainerCPUReqs "azurefile-csi"}}
              memory: {{ContainerMemReqs "azurefile-csi"}}
      volumes:
        - name: csi-proxy-fs-pipe-v1
          hostPath:
            path: \\.\pipe\csi-proxy-filesystem-v1
        - name: csi-proxy-smb-pipe-v1
          hostPath:
            path: \\.\pipe\csi-proxy-smb-v1
        # these paths are still included for compatibility, they're used
        # only if the node has still the beta version of the CSI proxy
        - name: csi-proxy-fs-pipe-v1beta1
          hostPath:
            path: \\.\pipe\csi-proxy-filesystem-v1beta1
        - name: csi-proxy-smb-pipe-v1beta1
          hostPath:
            path: \\.\pipe\csi-proxy-smb-v1beta1
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
            type: DirectoryOrCreate
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
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 50%
  selector:
    matchLabels:
      app: csi-azurefile-node
  template:
    metadata:
      labels:
        app: csi-azurefile-node
    spec:
      hostNetwork: true
      dnsPolicy: Default
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
            - --probe-timeout=3s
            - --health-port=29613
            - --v=2
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
            - --v=2
          livenessProbe:
            exec:
              command:
                - /csi-node-driver-registrar
                - --kubelet-registration-path=$(DRIVER_REG_SOCK_PATH)
                - --mode=kubelet-registration-probe
            initialDelaySeconds: 30
            timeoutSeconds: 15
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
            - "--metrics-address=0.0.0.0:29615"
            - "--user-agent-suffix=aks-engine"
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
            type: DirectoryOrCreate
          name: azure-cred
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
        kubernetes.io/role: master
      priorityClassName: system-cluster-critical
      tolerations:
        - key: "node-role.kubernetes.io/master"
          operator: "Exists"
          effect: "NoSchedule"
        - key: "node-role.kubernetes.io/controlplane"
          operator: "Exists"
          effect: "NoSchedule"
      containers:
        - name: csi-provisioner
          image: {{ContainerImage "csi-provisioner"}}
          args:
            - "-v=2"
            - "--csi-address=$(ADDRESS)"
            - "--leader-election"
            - "--timeout=300s"
            - "--extra-create-metadata=true"
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
            - "-v=2"
            - "-csi-address=$(ADDRESS)"
            - "-timeout=120s"
            - "-leader-election"
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
            - "-v=2"
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
        - name: csi-resizer
          image: {{ContainerImage "csi-resizer"}}
          args:
            - "-csi-address=$(ADDRESS)"
            - "-v=2"
            - "-leader-election"
            - '-handle-volume-inuse-error=false'
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
        - name: liveness-probe
          image: {{ContainerImage "livenessprobe"}}
          args:
            - --csi-address=/csi/csi.sock
            - --probe-timeout=3s
            - --health-port=29612
            - --v=2
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
            - "--metrics-address=0.0.0.0:29614"
            - "--user-agent-suffix=aks-engine"
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
            type: DirectoryOrCreate
---
# Source: azurefile-csi-driver/templates/csi-azurefile-driver.yaml
apiVersion: storage.k8s.io/v1
kind: CSIDriver
metadata:
  name: file.csi.azure.com
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  attachRequired: false
  podInfoOnMount: true
  volumeLifecycleModes:
    - Persistent
    - Ephemeral
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

var _k8sAddonsBlobfuseFlexvolumeYaml = []byte(`apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: blobfuse-flexvol-installer
  namespace: kube-system
  labels:
    k8s-app: blobfuse
    kubernetes.io/cluster-service: "true"
spec:
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 50%
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
{{- if not IsAzureCNI}}
          "ipam": {
              "type": "host-local",
              "subnet": "usePodCidr"
          },
{{- end}}
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
          value: "{{if IsAzureCNI}}azv{{else}}cali{{end}}"
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
      maxUnavailable: 50%
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
          value: "{{GetClusterSubnet}}"
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
          value: "{{if IsAzureCNI}}azv{{else}}cali{{end}}"
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
      maxUnavailable: 50%
    type: RollingUpdate
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    io.cilium/app: operator
    name: cilium-operator
    addonmanager.kubernetes.io/mode: "Reconcile"
  name: cilium-operator
  namespace: kube-system
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
- apiGroups: [""]
  resources: ["nodes/status"]
  verbs: ["patch"]
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
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 50%
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
        {{- if IsAzureStackCloud}}
        - --use-instance-metadata=false
        - --cloud-config=/etc/kubernetes/azure.json
        - --kubeconfig=/var/lib/kubelet/kubeconfig
        {{end}}
        env:
        - name: NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        {{- if IsAzureStackCloud}}
        - name: AZURE_ENVIRONMENT_FILEPATH
          value: /etc/kubernetes/azurestackcloud.json
        - name: AZURE_GO_SDK_LOG_LEVEL
          value: INFO
        {{end}}
        resources:
          requests:
            cpu: 50m
            memory: 50Mi
          limits:
            cpu: 2000m
            memory: 512Mi
        {{- if IsAzureStackCloud}}
        volumeMounts:
        - name: etc-kubernetes
          mountPath: /etc/kubernetes
          readOnly: true
        - name: etc-ssl
          mountPath: /etc/ssl
          readOnly: true
        - name: path-kubeconfig
          mountPath: /var/lib/kubelet/kubeconfig
          readOnly: true
      volumes:
        - name: etc-kubernetes
          hostPath:
            path: /etc/kubernetes
        - name: etc-ssl
          hostPath:
            path: /etc/ssl
        - name: path-kubeconfig
          hostPath:
            path: /var/lib/kubelet/kubeconfig
            type: FileOrCreate
        {{end}}
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
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 50%
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
        - --kubeconfig=C:\k\config
        {{- if IsAzureStackCloud}}
        - --use-instance-metadata=false
        - --cloud-config=C:\k\azure.json
        lifecycle:
          postStart:
            exec:
              command:
              - C:\k\addazsroot.bat
        {{end}}
        env:
        - name: NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        {{- if IsAzureStackCloud}}
        - name: AZURE_ENVIRONMENT_FILEPATH
          value: C:\k\azurestackcloud.json
        - name: AZURE_GO_SDK_LOG_LEVEL
          value: INFO
        {{end}}
        resources:
          requests:
            cpu: 50m
            memory: 50Mi
          limits:
            cpu: 2000m
            memory: 512Mi
        volumeMounts:
        - name: azure-config
          mountPath: C:\k
      volumes:
        - name: azure-config
          hostPath:
            path: C:\k
            type: Directory
{{end}}`)

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
apiVersion: apps/v1
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
        kubernetes.azure.com/role: master
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
    resources: ["pods", "events", "nodes", "nodes/stats", "nodes/metrics", "nodes/spec", "nodes/proxy", "namespaces", "services", "persistentvolumes"]
    verbs: ["list", "get", "watch"]
  - apiGroups: ["apps", "extensions", "autoscaling"]
    resources: ["replicasets", "deployments", "horizontalpodautoscalers"]
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
     </source>

     #Kubernetes Persistent Volume inventory
     <source>
      type kubepvinventory
      tag oms.containerinsights.KubePVInventory
      run_interval 60
      log_level debug
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

     #Kubernetes object state - deployments
     <source>
      type kubestatedeployments
      tag oms.containerinsights.KubeStateDeployments
      run_interval 60
      log_level debug
     </source>

     #Kubernetes object state - HPA
     <source>
      type kubestatehpa
      tag oms.containerinsights.KubeStateHpa
      run_interval 60
      log_level debug
     </source>

     <filter mdm.kubenodeinventory**>
      type filter_inventory2mdm
      log_level info
     </filter>

     #custom_metrics_mdm filter plugin for perf data from windows nodes
     <filter mdm.cadvisorperf**>
      type filter_cadvisor2mdm
      metrics_to_collect cpuUsageNanoCores,memoryWorkingSetBytes,pvUsedBytes
      log_level info
     </filter>

     #health model aggregation filter
     <filter kubehealth**>
      type filter_health_model_builder
     </filter>

     <match oms.containerinsights.KubePodInventory**>
      type out_oms
      log_level debug
      num_threads 2
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

     <match oms.containerinsights.KubePVInventory**>
      type out_oms
      log_level debug
      num_threads 5
      buffer_chunk_limit 4m
      buffer_type file
      buffer_path %STATE_DIR_WS%/state/out_oms_kubepv*.buffer
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
      num_threads 2
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
      num_threads 2
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
      num_threads 2
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
      dnsConfig:
        options:
          - name: ndots
            value: "3"
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
            timeoutSeconds: 15
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
              readOnly: true
            - mountPath: /mnt/docker
              name: containerlog-path-2
              readOnly: true
            - mountPath: /mnt/containers
              name: containerlog-path-3
              readOnly: true
            - mountPath: /etc/kubernetes/host
              name: azure-json-path
            - mountPath: /etc/omsagent-secret
              name: omsagent-secret
              readOnly: true
            - mountPath: /etc/config/settings
              name: settings-vol-config
              readOnly: true
            - mountPath: /etc/config/settings/adx
              name: omsagent-adx-secret
              readOnly: true
        - name: omsagent-prometheus
          image: {{ContainerImage "omsagent"}}
          imagePullPolicy: IfNotPresent
          resources:
            limits:
              cpu: {{ContainerCPULimits "omsagent-prometheus"}}
              memory: {{ContainerMemLimits "omsagent-prometheus"}}
            requests:
              cpu: {{ContainerCPUReqs "omsagent-prometheus"}}
              memory: {{ContainerMemReqs "omsagent-prometheus"}}
          env:
            # azure devops pipeline uses AKS_RESOURCE_ID and AKS_REGION hence ensure to uncomment these
            - name: ACS_RESOURCE_NAME
              value: {{ContainerConfig "clusterName"}}
            - name: CONTROLLER_TYPE
              value: "DaemonSet"
            - name: ISTEST
              value: "true"
            - name: CONTAINER_TYPE
              value: "PrometheusSidecar"
            - name: CONTROLLER_TYPE
              value: "DaemonSet"
            - name: NODE_IP
              valueFrom:
                fieldRef:
                  fieldPath: status.hostIP
            - name: USER_ASSIGNED_IDENTITY_CLIENT_ID
              value: ""
            - name: USING_AAD_MSI_AUTH
              value: "false"
          securityContext:
            privileged: true
          volumeMounts:
            - mountPath: /etc/kubernetes/host
              name: azure-json-path
            - mountPath: /etc/omsagent-secret
              name: omsagent-secret
              readOnly: true
            - mountPath: /etc/config/settings
              name: settings-vol-config
              readOnly: true
            - mountPath: /etc/config/osm-settings
              name: osm-settings-vol-config
              readOnly: true
          livenessProbe:
            exec:
              command:
                - /bin/bash
                - -c
                - /opt/livenessprobe.sh
            initialDelaySeconds: 60
            periodSeconds: 60
            timeoutSeconds: 15
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
        - name: containerlog-path-2
          hostPath:
            path: /mnt/docker
        - name: containerlog-path-3
          hostPath:
            path: /mnt/containers
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
        - name: omsagent-adx-secret
          secret:
            secretName: omsagent-adx-secret
            optional: true
        - name: osm-settings-vol-config
          configMap:
            name: container-azm-ms-osmconfig
            optional: true
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 50%
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
              cpu: {{ContainerCPULimits "omsagent-rs"}}
              memory: {{ContainerMemLimits "omsagent-rs"}}
            requests:
              cpu: {{ContainerCPUReqs "omsagent-rs"}}
              memory: {{ContainerMemReqs "omsagent-rs"}}
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
            - name: SIDECAR_SCRAPING_ENABLED
              value: "true"
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
            - mountPath: /etc/config/settings/adx
              name: omsagent-adx-secret
              readOnly: true
            - mountPath: /etc/config/osm-settings
              name: osm-settings-vol-config
              readOnly: true
          livenessProbe:
            exec:
              command:
                - /bin/bash
                - -c
                - /opt/livenessprobe.sh
            initialDelaySeconds: 60
            periodSeconds: 60
            timeoutSeconds: 15
      affinity:
        nodeAffinity:
          # affinity to schedule on to ephemeral os node if its available
          preferredDuringSchedulingIgnoredDuringExecution:
          - weight: 1
            preference:
              matchExpressions:
              - key: storageprofile
                operator: NotIn
                values:
                - managed
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
        - name: omsagent-adx-secret
          secret:
            secretName: omsagent-adx-secret
            optional: true
        - name: osm-settings-vol-config
          configMap:
            name: container-azm-ms-osmconfig
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
    rollingUpdate:
      maxUnavailable: 50%
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
      dnsConfig:
        options:
          - name: ndots
            value: "3"
      containers:
        - name: omsagent-win
          image: {{ContainerImage "omsagent-win"}}
          imagePullPolicy: IfNotPresent
          resources:
            limits:
              cpu: {{ContainerCPULimits "omsagent-win"}}
              memory: {{ContainerMemLimits "omsagent-win"}}
          env:
            - name: ACS_RESOURCE_NAME
              value: {{ContainerConfig "clusterName"}}
            - name: CONTROLLER_TYPE
              value: "DaemonSet"
            - name: HOSTNAME
              valueFrom:
                 fieldRef:
                   fieldPath: spec.nodeName
            - name: PODNAME
              valueFrom:
                 fieldRef:
                   fieldPath: metadata.name
            - name: NODE_IP
              valueFrom:
                 fieldRef:
                   fieldPath: status.hostIP
            - name: SIDECAR_SCRAPING_ENABLED
              value: "true"
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
            - mountPath: C:\etc\config\adx
              name: omsagent-adx-secret
              readOnly: true
          livenessProbe:
            exec:
              command:
                - cmd
                - /c
                - C:\opt\omsagentwindows\scripts\cmd\livenessprobe.exe
                - fluent-bit.exe
                - fluentdwinaks
                - "C:\\etc\\omsagentwindows\\filesystemwatcher.txt"
                - "C:\\etc\\omsagentwindows\\renewcertificate.txt"
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
        - name: omsagent-adx-secret
          secret:
            secretName: omsagent-adx-secret
            optional: true
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
# this is for versions >=1.19, for versions <1.19 we continue to use v1beta1
{{- if IsKubernetesVersionGe "1.19.0"}}
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: healthstates.azmon.container.insights
  namespace: kube-system
spec:
  group: azmon.container.insights
  versions:
  - name: v1
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        type: object
        properties:
          state:
            type: string
  scope: Namespaced
  names:
    plural: healthstates
    kind: HealthState
{{- else }}
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
{{end}}
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
        health {
          # this should be > readiness probe failure time
          lameduck 35s
        }
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
        kubernetes.azure.com/role: agent
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
          periodSeconds: 10
          timeoutSeconds: 1
          successThreshold: 1
          failureThreshold: 3
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
      "Network": "{{GetClusterSubnet}}",
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
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 50%
  selector:
    matchLabels:
      tier: node
      app: flannel
  template:
    metadata:
      labels:
        tier: node
        app: flannel
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

var _k8sAddonsIpMasqAgentYaml = []byte(`apiVersion: apps/v1
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
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 50%
  selector:
    matchLabels:
      k8s-app: azure-ip-masq-agent
      tier: node
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
        args:
          - --enable-ipv6={{ContainerConfig "enable-ipv6"}}
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
{{- if ContainerConfig "secondary-non-masquerade-cidr"}}
      - {{ContainerConfig "secondary-non-masquerade-cidr"}}
{{end -}}
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

var _k8sAddonsKeyvaultFlexvolumeYaml = []byte(`apiVersion: apps/v1
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
    rollingUpdate:
      maxUnavailable: 50%
  selector:
    matchLabels:
      app: keyvault-flexvolume
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

var _k8sAddonsKubeProxyYaml = []byte(`apiVersion: v1
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
  selector:
    matchLabels:
      k8s-app: kube-proxy
      component: kube-proxy
      tier: node
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
        - kube-proxy
        - --config=/var/lib/kube-proxy/config.yaml
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

var _k8sAddonsKubernetesDashboardYaml = []byte(`{{- /* Note: dashboard addon is deprecated */}}
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
  labels:
    k8s-app: metrics-server
    addonmanager.kubernetes.io/mode: {{GetMode}}
  name: metrics-server
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    k8s-app: metrics-server
    addonmanager.kubernetes.io/mode: {{GetMode}}
    rbac.authorization.k8s.io/aggregate-to-admin: "true"
    rbac.authorization.k8s.io/aggregate-to-edit: "true"
    rbac.authorization.k8s.io/aggregate-to-view: "true"
  name: system:aggregated-metrics-reader
rules:
- apiGroups:
  - metrics.k8s.io
  resources:
  - pods
  - nodes
  verbs:
  - get
  - list
  - watch
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    k8s-app: metrics-server
    addonmanager.kubernetes.io/mode: {{GetMode}}
  name: system:metrics-server
rules:
- apiGroups:
  - ""
  resources:
  - pods
  - nodes
  - nodes/stats
  - namespaces
  - configmaps
  verbs:
  - get
  - list
  - watch
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  labels:
    k8s-app: metrics-server
    addonmanager.kubernetes.io/mode: {{GetMode}}
  name: metrics-server-auth-reader
  namespace: kube-system
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: extension-apiserver-authentication-reader
subjects:
- kind: ServiceAccount
  name: metrics-server
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    k8s-app: metrics-server
    addonmanager.kubernetes.io/mode: {{GetMode}}
  name: metrics-server:system:auth-delegator
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:auth-delegator
subjects:
- kind: ServiceAccount
  name: metrics-server
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    k8s-app: metrics-server
    addonmanager.kubernetes.io/mode: {{GetMode}}
  name: system:metrics-server
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:metrics-server
subjects:
- kind: ServiceAccount
  name: metrics-server
  namespace: kube-system
---
apiVersion: v1
kind: Service
metadata:
  labels:
    k8s-app: metrics-server
    addonmanager.kubernetes.io/mode: {{GetMode}}
  name: metrics-server
  namespace: kube-system
spec:
  ports:
  - name: https
    port: 443
    protocol: TCP
    targetPort: https
  selector:
    k8s-app: metrics-server
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    k8s-app: metrics-server
    addonmanager.kubernetes.io/mode: {{GetMode}}
  name: metrics-server
  namespace: kube-system
spec:
  selector:
    matchLabels:
      k8s-app: metrics-server
  strategy:
    rollingUpdate:
      maxUnavailable: 0
  template:
    metadata:
      labels:
        k8s-app: metrics-server
    spec:
      containers:
      - args:
        - --cert-dir=/tmp
        - --secure-port=4443
        - --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname
        - --kubelet-use-node-status-port
        - --metric-resolution=15s
        - --kubelet-insecure-tls
        image: {{ContainerImage "metrics-server"}}
        imagePullPolicy: IfNotPresent
        livenessProbe:
          failureThreshold: 3
          httpGet:
            path: /livez
            port: https
            scheme: HTTPS
          periodSeconds: 10
        name: metrics-server
        ports:
        - containerPort: 4443
          name: https
          protocol: TCP
        readinessProbe:
          failureThreshold: 3
          httpGet:
            path: /readyz
            port: https
            scheme: HTTPS
          initialDelaySeconds: 20
          periodSeconds: 10
        resources:
          requests:
            cpu: 100m
            memory: 200Mi
        securityContext:
          readOnlyRootFilesystem: true
          runAsNonRoot: true
          runAsUser: 1000
        volumeMounts:
        - mountPath: /tmp
          name: tmp-dir
      nodeSelector:
        kubernetes.io/os: linux
      priorityClassName: system-cluster-critical
      serviceAccountName: metrics-server
      volumes:
      - emptyDir: {}
        name: tmp-dir
---
apiVersion: apiregistration.k8s.io/v1{{- if not (IsKubernetesVersionGe "1.19.0")}}beta1{{end}}
kind: APIService
metadata:
  labels:
    k8s-app: metrics-server
    addonmanager.kubernetes.io/mode: {{GetMode}}
  name: v1beta1.metrics.k8s.io
spec:
  group: metrics.k8s.io
  groupPriorityMinimum: 100
  insecureSkipTLSVerify: true
  service:
    name: metrics-server
    namespace: kube-system
  version: v1beta1
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
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 50%
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
    rollingUpdate:
      maxUnavailable: 50%
  template:
    metadata:
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

var _k8sAddonsPodSecurityPolicyYaml = []byte(`apiVersion: policy/v1beta1
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
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 50%
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
kind: ClusterRole
metadata:
  name: secretproviderclasses-role
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups:
  - ""
  resources:
  - events
  verbs:
  - create
  - patch
- apiGroups:
  - ""
  resources:
  - pods
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - secrets-store.csi.x-k8s.io
  resources:
  - secretproviderclasses
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - secrets-store.csi.x-k8s.io
  resources:
  - secretproviderclasspodstatuses
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - secrets-store.csi.x-k8s.io
  resources:
  - secretproviderclasspodstatuses/status
  verbs:
  - get
  - patch
  - update
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
  name: secretprovidersyncing-role
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
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
kind: ClusterRoleBinding
metadata:
  name: secretprovidersyncing-rolebinding
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: secretprovidersyncing-role
subjects:
- kind: ServiceAccount
  name: secrets-store-csi-driver
  namespace: kube-system
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.4.0
  creationTimestamp: null
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
  scope: Namespaced
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: SecretProviderClass is the Schema for the secretproviderclasses API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
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
              secretObjects:
                items:
                  description: SecretObject defines the desired state of synced K8s secret objects
                  properties:
                    data:
                      items:
                        description: SecretObjectData defines the desired state of synced K8s secret object data
                        properties:
                          key:
                            description: data field to populate
                            type: string
                          objectName:
                            description: name of the object to sync
                            type: string
                        type: object
                      type: array
                    labels:
                      additionalProperties:
                        type: string
                      description: labels of K8s secret object
                      type: object
                    secretName:
                      description: name of the K8s secret object
                      type: string
                    type:
                      description: type of K8s secret object
                      type: string
                  type: object
                type: array
            type: object
          status:
            description: SecretProviderClassStatus defines the observed state of SecretProviderClass
            properties:
              byPod:
                items:
                  description: ByPodStatus defines the state of SecretProviderClass as seen by an individual controller
                  properties:
                    id:
                      description: id of the pod that wrote the status
                      type: string
                    namespace:
                      description: namespace of the pod that wrote the status
                      type: string
                  type: object
                type: array
            type: object
        type: object
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.4.0
  creationTimestamp: null
  name: secretproviderclasspodstatuses.secrets-store.csi.x-k8s.io
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  group: secrets-store.csi.x-k8s.io
  names:
    kind: SecretProviderClassPodStatus
    listKind: SecretProviderClassPodStatusList
    plural: secretproviderclasspodstatuses
    singular: secretproviderclasspodstatus
  scope: Namespaced
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: SecretProviderClassPodStatus is the Schema for the secretproviderclassespodstatus API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          status:
            description: SecretProviderClassPodStatusStatus defines the observed state of SecretProviderClassPodStatus
            properties:
              mounted:
                type: boolean
              objects:
                items:
                  description: SecretProviderClassObject defines the object fetched from external secrets store
                  properties:
                    id:
                      type: string
                    version:
                      type: string
                  type: object
                type: array
              podName:
                type: string
              secretProviderClassName:
                type: string
              targetPath:
                type: string
            type: object
        type: object
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
---
{{- /* A priority class for the daemonset such that they are not */}}
{{- /* frozen out of a node due to the node filling up with "normal" */}}
{{- /* pods before the daemonset controller can get the daemonset */}}
{{- /* pods to be scheduled. */}}
apiVersion: scheduling.k8s.io/v1
kind: PriorityClass
metadata:
  name: csi-secrets-store
  labels:
    addonmanager.kubernetes.io/mode: EnsureExists
value: 1000
globalDefault: false
description: "This is the daemonset priority class for csi-secrets-store"
---
kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: csi-secrets-store
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 50%
  selector:
    matchLabels:
      app: csi-secrets-store
  template:
    metadata:
      labels:
        app: csi-secrets-store
    spec:
      priorityClassName: csi-secrets-store
      serviceAccountName: secrets-store-csi-driver
      hostNetwork: true
      containers:
        - name: node-driver-registrar
          image: {{ContainerImage "csi-node-driver-registrar"}}
          args:
            - --v=5
            - --csi-address=/csi/csi.sock
            - --kubelet-registration-path=/var/lib/kubelet/plugins/csi-secrets-store/csi.sock
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
            - "--endpoint=$(CSI_ENDPOINT)"
            - "--nodeid=$(KUBE_NODE_NAME)"
            - "--provider-volume=/etc/kubernetes/secrets-store-csi-providers"
            - "--grpc-supported-providers=azure"
            - "--metrics-addr=:{{ContainerConfig "metricsPort"}}"
            - "--enable-secret-rotation={{ContainerConfig "enableSecretRotation"}}"
            - "--rotation-poll-interval={{ContainerConfig "rotationPollInterval"}}"
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
{{- if IsCustomCloudProfile}}
            - name: custom-environment
              mountPath: /etc/kubernetes/azurestackcloud.json
              readOnly: true
{{end}}
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
          - --http-endpoint=0.0.0.0:9808
          - -v=2
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
{{- if IsCustomCloudProfile}}
        - name: custom-environment
          hostPath:
            path: /etc/kubernetes/azurestackcloud.json
            type: FileOrCreate
{{end}}
      nodeSelector:
        kubernetes.io/os: linux
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: csi-secrets-store-provider-azure
  namespace: kube-system
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
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
    rollingUpdate:
      maxUnavailable: 50%
  selector:
    matchLabels:
      app: csi-secrets-store-provider-azure
  template:
    metadata:
      labels:
        app: csi-secrets-store-provider-azure
    spec:
      priorityClassName: csi-secrets-store
      serviceAccountName: csi-secrets-store-provider-azure
      hostNetwork: true
      containers:
        - name: provider-azure-installer
          image: {{ContainerImage "provider-azure-installer"}}
          imagePullPolicy: IfNotPresent
          args:
            - --endpoint=unix:///etc/kubernetes/secrets-store-csi-providers/azure.sock
          lifecycle:
            preStop:
              exec:
                command:
                  - "rm /etc/kubernetes/secrets-store-csi-providers/azure.sock"
          volumeMounts:
            - mountPath: "/etc/kubernetes/secrets-store-csi-providers"
              name: providervol
            - name: mountpoint-dir
              mountPath: /var/lib/kubelet/pods
              mountPropagation: HostToContainer
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
        - name: mountpoint-dir
          hostPath:
            path: /var/lib/kubelet/pods
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

var _k8sAddonsSmbFlexvolumeYaml = []byte(`apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: smb-flexvol-installer
  namespace: kube-system
  labels:
    k8s-app: smb
    kubernetes.io/cluster-service: "true"
spec:
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 50%
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

var _k8sCloudInitArtifactsApiserverMonitorService = []byte(`[Unit]
Description=a script that checks apiserver health and restarts if needed
After=kubelet.service
[Service]
Restart=always
RestartSec=10
RemainAfterExit=yes
ExecStart=/usr/local/bin/health-monitor.sh apiserver
[Install]
WantedBy=multi-user.target
#EOF
`)

func k8sCloudInitArtifactsApiserverMonitorServiceBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsApiserverMonitorService, nil
}

func k8sCloudInitArtifactsApiserverMonitorService() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsApiserverMonitorServiceBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/apiserver-monitor.service", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
KUBECTL="/usr/local/bin/kubectl --kubeconfig=/home/$ADMINUSER/.kube/config"
ADDONS_DIR=/etc/kubernetes/addons
POD_SECURITY_POLICY_SPEC=$ADDONS_DIR/pod-security-policy.yaml
ADDON_MANAGER_SPEC=/etc/kubernetes/manifests/kube-addon-manager.yaml
GET_KUBELET_LOGS="journalctl -u kubelet --no-pager"

systemctlEnableAndStart() {
  local ret
  systemctl_restart 100 5 30 $1
  ret=$?
  systemctl status $1 --no-pager -l >/var/log/azure/$1-status.log
  if [ $ret -ne 0 ]; then
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
  local apiserver_key="/etc/kubernetes/certs/apiserver.key" ca_key="/etc/kubernetes/certs/ca.key" etcdserver_key="/etc/kubernetes/certs/etcdserver.key"
  touch "${apiserver_key}"
  touch "${ca_key}"
  touch "${etcdserver_key}"
  if [[ -z ${COSMOS_URI} ]]; then
    chown etcd:etcd "${etcdserver_key}"
  fi
  local etcdclient_key="/etc/kubernetes/certs/etcdclient.key" etcdpeer_key="/etc/kubernetes/certs/etcdpeer${NODE_INDEX}.key"
  touch "${etcdclient_key}"
  touch "${etcdpeer_key}"
  if [[ -z ${COSMOS_URI} ]]; then
    chown etcd:etcd "${etcdpeer_key}"
  fi
  chmod 0600 "${apiserver_key}" "${ca_key}" "${etcdserver_key}" "${etcdclient_key}" "${etcdpeer_key}"
  chown root:root "${apiserver_key}" "${ca_key}" "${etcdclient_key}"
  local etcdserver_crt="/etc/kubernetes/certs/etcdserver.crt" etcdclient_crt="/etc/kubernetes/certs/etcdclient.crt" etcdpeer_crt="/etc/kubernetes/certs/etcdpeer${NODE_INDEX}.crt"
  touch "${etcdserver_crt}"
  touch "${etcdclient_crt}"
  touch "${etcdpeer_crt}"
  chmod 0644 "${etcdserver_crt}" "${etcdclient_crt}" "${etcdpeer_crt}"
  chown root:root "${etcdserver_crt}" "${etcdclient_crt}" "${etcdpeer_crt}"

  set +x
  echo "${APISERVER_PRIVATE_KEY}" | base64 --decode >"${apiserver_key}"
  echo "${CA_PRIVATE_KEY}" | base64 --decode >"${ca_key}"
  echo "${ETCD_SERVER_PRIVATE_KEY}" | base64 --decode >"${etcdserver_key}"
  echo "${ETCD_CLIENT_PRIVATE_KEY}" | base64 --decode >"${etcdclient_key}"
  echo "${ETCD_PEER_KEY}" | base64 --decode >"${etcdpeer_key}"
  echo "${ETCD_SERVER_CERTIFICATE}" | base64 --decode >"${etcdserver_crt}"
  echo "${ETCD_CLIENT_CERTIFICATE}" | base64 --decode >"${etcdclient_crt}"
  echo "${ETCD_PEER_CERT}" | base64 --decode >"${etcdpeer_crt}"
}
configureEtcd() {
  set -x

  local ret f=/opt/azure/containers/setup-etcd.sh etcd_peer_url="https://${PRIVATE_IP}:2380"
  wait_for_file 1200 1 $f || exit {{GetCSEErrorCode "ERR_ETCD_CONFIG_FAIL"}}
  $f >/opt/azure/containers/setup-etcd.log 2>&1
  ret=$?
  if [ $ret -ne 0 ]; then
    exit $ret
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
  retrycmd 120 5 25 sudo -E etcdctl member update $MEMBER ${etcd_peer_url} || exit {{GetCSEErrorCode "ERR_ETCD_CONFIG_FAIL"}}
}
configureChrony() {
  sed -i "s/makestep.*/makestep 1.0 -1/g" /etc/chrony/chrony.conf
  echo "refclock PHC /dev/ptp0 poll 3 dpoll -2 offset 0" >> /etc/chrony/chrony.conf
}
ensureChrony() {
  systemctlEnableAndStart chrony || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
}
disable1804SystemdResolved() {
  {{/* Ignoring systemd-resolved query service but using its resolv.conf file */}}
  {{/* This is the simplest approach to workaround resolved issues without completely uninstall it */}}
  [ -f /run/systemd/resolve/resolv.conf ] && sudo ln -sf /run/systemd/resolve/resolv.conf /etc/resolv.conf
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
  local s=/lib/systemd/system/cron.service
  if [[ -f ${s} ]]; then
    if ! grep -q 'Restart=' ${s}; then
      sed -i 's/\[Service\]/[Service]\nRestart=always/' ${s}
      systemctlEnableAndStart cron
    fi
  fi
}
generateAggregatedAPICerts() {
  local f=/etc/kubernetes/generate-proxy-certs.sh
  wait_for_file 1200 1 $f || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  $f
}
configureKubeletServerCert() {
  local kubeletserver_key="/etc/kubernetes/certs/kubeletserver.key" kubeletserver_crt="/etc/kubernetes/certs/kubeletserver.crt"

  openssl genrsa -out $kubeletserver_key 2048
  openssl req -new -x509 -days 7300 -key $kubeletserver_key -out $kubeletserver_crt -subj "/CN=${NODE_NAME}"
}
configureK8s() {
  local client_key="/etc/kubernetes/certs/client.key" apiserver_crt="/etc/kubernetes/certs/apiserver.crt" azure_json="/etc/kubernetes/azure.json"
  touch "${client_key}"
  chmod 0600 "${client_key}"
  chown root:root "${client_key}"
  if [[ -n ${MASTER_NODE} ]]; then
    touch "${apiserver_crt}"
    chmod 0644 "${apiserver_crt}"
    chown root:root "${apiserver_crt}"
  fi
  set +x
  echo "${KUBELET_PRIVATE_KEY}" | base64 --decode >"${client_key}"
  configureKubeletServerCert
  if [[ -n ${MASTER_NODE} ]]; then
    echo "${APISERVER_PUBLIC_KEY}" | base64 --decode >"${apiserver_crt}"
    if [[ ${ENABLE_AGGREGATED_APIS} == True ]]; then
      generateAggregatedAPICerts
    fi
  else
    {{- /* If we are a node that does not need azure.json (cloud-init tells us), then return immediately */}}
    wait_for_file 1 1 /opt/azure/needs_azure.json || return
  fi

  touch $azure_json
  chmod 0600 $azure_json
  chown root:root $azure_json
  {{/* Perform the required JSON escaping */}}
  local sp_secret=${SERVICE_PRINCIPAL_CLIENT_SECRET//\\/\\\\}
  sp_secret=${SERVICE_PRINCIPAL_CLIENT_SECRET//\"/\\\"}
  cat <<EOF >"${azure_json}"
{
    "cloud":"{{GetTargetEnvironment}}",
    "tenantId": "${TENANT_ID}",
    "subscriptionId": "${SUBSCRIPTION_ID}",
    "aadClientId": "${SERVICE_PRINCIPAL_CLIENT_ID}",
    "aadClientSecret": "${sp_secret}",
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
    "providerKeyVersion": "",
    "enableMultipleStandardLoadBalancers": ${ENABLE_MULTIPLE_STANDARD_LOAD_BALANCERS},
    "tags": "${TAGS}"
}
EOF
  set -x
  if [[ ${CLOUDPROVIDER_BACKOFF_MODE} == "v2" ]]; then
    sed -i "/cloudProviderBackoffExponent/d" $azure_json
    sed -i "/cloudProviderBackoffJitter/d" $azure_json
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
  local tmpDir=$(mktemp -d "$(pwd)/XXX")
  if [[ "${NETWORK_PLUGIN}" == "azure" ]]; then
    mv $CNI_BIN_DIR/10-azure.conflist $CNI_CONFIG_DIR/
    chmod 600 $CNI_CONFIG_DIR/10-azure.conflist
{{- if IsIPv6DualStackFeatureEnabled}}
    jq '.plugins[0].ipv6Mode="ipv6nat"' "$CNI_CONFIG_DIR/10-azure.conflist" > $tmpDir/tmp
    mv $tmpDir/tmp $CNI_CONFIG_DIR/10-azure.conflist
{{- end}}
    if [[ {{GetKubeProxyMode}} == "ipvs" ]]; then
      serviceCidrs={{GetServiceCidr}}
      jq --arg serviceCidrs $serviceCidrs '.plugins[0]+={serviceCidrs: $serviceCidrs}' "$CNI_CONFIG_DIR/10-azure.conflist" > $tmpDir/tmp
      mv $tmpDir/tmp $CNI_CONFIG_DIR/10-azure.conflist
    fi
    if [[ "${NETWORK_MODE}" == "bridge" ]]; then
      jq '.plugins[0].mode="bridge"' "$CNI_CONFIG_DIR/10-azure.conflist" > $tmpDir/tmp
      jq '.plugins[0].bridge="azure0"' "$tmpDir/tmp" > $tmpDir/tmp2
      mv $tmpDir/tmp2 $CNI_CONFIG_DIR/10-azure.conflist
    else
      jq '.plugins[0].mode="transparent"' "$CNI_CONFIG_DIR/10-azure.conflist" > $tmpDir/tmp
      mv $tmpDir/tmp $CNI_CONFIG_DIR/10-azure.conflist
    fi
    /sbin/ebtables -t nat --list
  fi
  rm -Rf $tmpDir
}
enableCRISystemdMonitor() {
  wait_for_file 1200 1 /etc/systemd/system/docker-monitor.service || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  systemctlEnableAndStart docker-monitor || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
}
{{- if NeedsContainerd}}
installContainerd() {
  removeMoby
  local v
  v=$(containerd -version | cut -d " " -f 3 | sed 's|v||')
  if [[ $v != "${CONTAINERD_VERSION}"* ]]; then
    os_lower=$(echo ${OS} | tr '[:upper:]' '[:lower:]')
    if [[ ${OS} == "${UBUNTU_OS_NAME}" ]]; then
      url_path="${os_lower}/${UBUNTU_RELEASE}/multiarch/prod"
    elif [[ ${OS} == "${DEBIAN_OS_NAME}" ]]; then
      url_path="${os_lower}/${UBUNTU_RELEASE}/prod"
    else
      exit 25
    fi
    removeContainerd
    apt_get_update || exit 99
    apt_get_install 20 30 120 moby-runc moby-containerd=${CONTAINERD_VERSION}* --allow-downgrades || exit 27
  fi
}
ensureContainerd() {
  wait_for_file 1200 1 /etc/systemd/system/containerd.service.d/exec_start.conf || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  wait_for_file 1200 1 /etc/containerd/config.toml || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  {{- if HasKubeReservedCgroup}}
  wait_for_file 1200 1 /etc/systemd/system/containerd.service.d/kubereserved-slice.conf|| exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  {{- end}}
  systemctlEnableAndStart containerd || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
  enableCRISystemdMonitor
}
{{end}}
{{- if IsDockerContainerRuntime}}
ensureDocker() {
  wait_for_file 1200 1 /etc/systemd/system/docker.service.d/exec_start.conf || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  usermod -aG docker ${ADMINUSER}
  if [[ $OS != $FLATCAR_OS_NAME ]]; then
    wait_for_file 1200 1 /etc/systemd/system/docker.service.d/clear_mount_propagation_flags.conf || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  fi
  {{- if HasKubeReservedCgroup}}
  wait_for_file 1200 1 /etc/systemd/system/docker.service.d/kubereserved-slice.conf || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  {{- end}}
  local daemon_json=/etc/docker/daemon.json
  for i in $(seq 1 1200); do
    if [ -s $daemon_json ]; then
      jq '.' <$daemon_json && break
    fi
    if [ $i -eq 1200 ]; then
      exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
    else
      sleep 1
    fi
  done
  systemctlEnableAndStart docker || exit {{GetCSEErrorCode "ERR_DOCKER_START_FAIL"}}
  enableCRISystemdMonitor
}
{{end}}
{{- if IsIPv6Enabled}}
ensureDHCPv6() {
  if [[ ${UBUNTU_RELEASE} == "16.04" ]]; then
    wait_for_file 3600 1 {{GetDHCPv6ServiceCSEScriptFilepath}} || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
    wait_for_file 3600 1 {{GetDHCPv6ConfigCSEScriptFilepath}} || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
    systemctlEnableAndStart dhcpv6 || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
    retrycmd 120 5 25 modprobe ip6_tables || exit {{GetCSEErrorCode "ERR_MODPROBE_FAIL"}}
  fi
}
{{end}}
{{- if EnableEncryptionWithExternalKms}}
ensureKMSKeyvaultKey() {
    wait_for_file 3600 1 {{GetKMSKeyvaultKeyServiceCSEScriptFilepath}} || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
    wait_for_file 3600 1 {{GetKMSKeyvaultKeyCSEScriptFilepath}} || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
    systemctlEnableAndStart kms-keyvault-key || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
}
{{end}}
ensureKubelet() {
  wait_for_file 1200 1 /etc/sysctl.d/11-aks-engine.conf || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  sysctl_reload 10 5 120 || exit {{GetCSEErrorCode "ERR_SYSCTL_RELOAD"}}
  wait_for_file 1200 1 /etc/default/kubelet || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  wait_for_file 1200 1 /var/lib/kubelet/kubeconfig || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  if [[ -n ${MASTER_NODE} ]]; then
{{- if IsMasterVirtualMachineScaleSets}}
    sed -i "s|<SERVERIP>|https://$PRIVATE_IP:443|g" "/var/lib/kubelet/kubeconfig" || exit {{GetCSEErrorCode "ERR_KUBELET_START_FAIL"}}
{{- end}}
    local f=/etc/kubernetes/manifests/kube-apiserver.yaml
    wait_for_file 1200 1 $f || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
    sed -i "s|<advertiseAddr>|$PRIVATE_IP|g" $f
  fi
  wait_for_file 1200 1 /opt/azure/containers/kubelet.sh || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
{{- if HasKubeReservedCgroup}}
  wait_for_file 1200 1 /etc/systemd/system/{{- GetKubeReservedCgroup -}}.slice || exit {{GetCSEErrorCode "ERR_KUBERESERVED_SLICE_SETUP_FAIL"}}
  wait_for_file 1200 1 /etc/systemd/system/kubelet.service.d/kubereserved-slice.conf || exit {{GetCSEErrorCode "ERR_KUBELET_SLICE_SETUP_FAIL"}}
{{- end}}
  if [[ -n ${MASTER_NODE} ]]; then
    systemctlEnableAndStart kubelet || exit {{GetCSEErrorCode "ERR_KUBELET_START_FAIL"}}
  else
{{- if not RunUnattendedUpgradesOnBootstrap}}
    systemctlEnableAndStart kubelet || exit {{GetCSEErrorCode "ERR_KUBELET_START_FAIL"}}
{{else}}
    systemctl_enable 100 5 30 kubelet || exit {{GetCSEErrorCode "ERR_KUBELET_START_FAIL"}}
{{- end}}
  fi
{{- if HasKubeletHealthZPort}}
  wait_for_file 1200 1 /etc/systemd/system/kubelet-monitor.service || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  if [[ -n ${MASTER_NODE} ]]; then
    systemctlEnableAndStart kubelet-monitor || exit {{GetCSEErrorCode "ERR_KUBELET_START_FAIL"}}
  else
  {{- if not RunUnattendedUpgradesOnBootstrap}}
    systemctlEnableAndStart kubelet-monitor || exit {{GetCSEErrorCode "ERR_KUBELET_START_FAIL"}}
  {{else}}
    systemctl_enable 100 5 30 kubelet-monitor || exit {{GetCSEErrorCode "ERR_KUBELET_START_FAIL"}}
  {{- end}}
  fi
{{- end}}
if [[ -n ${MASTER_NODE} ]]; then
  wait_for_file 1200 1 /etc/systemd/system/apiserver-monitor.service || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  systemctlEnableAndStart apiserver-monitor || exit {{GetCSEErrorCode "ERR_KUBELET_START_FAIL"}}
fi
}

ensureKubeAddonManager() {
  {{/* Wait 5 mins for kube-addon-manager to become Ready */}}
  if ! retrycmd 60 5 30 ${KUBECTL} wait --for=condition=Ready --timeout=5s -l app=kube-addon-manager po -n kube-system; then
    {{/* Restart kubelet if kube-addon-manager is not Ready after 5 mins */}}
    systemctl_restart 3 5 30 kubelet
    {{/* Wait 5 more mins for kube-addon-manager to become Ready, and then return failure if not */}}
    retrycmd 60 5 30 ${KUBECTL} wait --for=condition=Ready --timeout=5s -l app=kube-addon-manager po -n kube-system || exit_cse {{GetCSEErrorCode "ERR_ADDONS_START_FAIL"}} $GET_KUBELET_LOGS
  fi
}

ensureAddons() {
{{- if IsDashboardAddonEnabled}} {{/* Note: dashboard addon is deprecated */}}
  retrycmd 120 5 30 $KUBECTL get namespace kubernetes-dashboard || exit_cse {{GetCSEErrorCode "ERR_ADDONS_START_FAIL"}} $GET_KUBELET_LOGS
{{- end}}
{{- if IsAzurePolicyAddonEnabled}}
  retrycmd 120 5 30 $KUBECTL get namespace gatekeeper-system || exit_cse {{GetCSEErrorCode "ERR_ADDONS_START_FAIL"}} $GET_KUBELET_LOGS
{{- end}}
{{- if not HasCustomPodSecurityPolicy}}
  retrycmd 120 5 30 $KUBECTL get podsecuritypolicy privileged restricted || exit_cse {{GetCSEErrorCode "ERR_ADDONS_START_FAIL"}} $GET_KUBELET_LOGS
{{- end}}
  replaceAddonsInit
  rm -Rf ${ADDONS_DIR}/init
  ensureKubeAddonManager
  {{/* Manually delete any kube-addon-manager pods that point to the init directory */}}
  for initPod in $(${KUBECTL} get pod -l app=kube-addon-manager -n kube-system -o json | jq -r '.items[] | select(.spec.containers[0].env[] | select(.value=="/etc/kubernetes/addons/init")) | select(.status.phase=="Running") .metadata.name'); do
    retrycmd 120 5 30 ${KUBECTL} delete pod $initPod -n kube-system --force --grace-period 0 || exit_cse {{GetCSEErrorCode "ERR_ADDONS_START_FAIL"}} $GET_KUBELET_LOGS
  done
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
replaceAddonsInit() {
  wait_for_file 1200 1 $ADDON_MANAGER_SPEC || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  cp $ADDON_MANAGER_SPEC /tmp/kube-addon-manager.yaml || exit {{GetCSEErrorCode "ERR_ADDONS_START_FAIL"}}
  sed -i "s|${ADDONS_DIR}/init|${ADDONS_DIR}|g" /tmp/kube-addon-manager.yaml || exit {{GetCSEErrorCode "ERR_ADDONS_START_FAIL"}}
  mv /tmp/kube-addon-manager.yaml $ADDON_MANAGER_SPEC || exit {{GetCSEErrorCode "ERR_ADDONS_START_FAIL"}}
}
ensureLabelNodes() {
  wait_for_file 1200 1 /opt/azure/containers/label-nodes.sh || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  wait_for_file 1200 1 /etc/systemd/system/label-nodes.service || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
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
  local binPath=/usr/local/bin
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
  if [ "$NO_OUTBOUND" = "true" ]; then
    return
  fi
  retrycmd 120 5 25 $KUBECTL 2>/dev/null cluster-info || exit_cse {{GetCSEErrorCode "ERR_K8S_RUNNING_TIMEOUT"}} $GET_KUBELET_LOGS
}
{{- if IsAzurePolicyAddonEnabled}}
ensureLabelExclusionForAzurePolicyAddon() {
  retrycmd 120 5 25 $KUBECTL label ns kube-system control-plane=controller-manager --overwrite 2>/dev/null || exit_cse {{GetCSEErrorCode "ERR_K8S_RUNNING_TIMEOUT"}} $GET_KUBELET_LOGS
}
{{end}}
ensureEtcd() {
  local etcd_client_url="https://${PRIVATE_IP}:2379"
  retrycmd 120 5 25 curl --cacert /etc/kubernetes/certs/ca.crt --cert /etc/kubernetes/certs/etcdclient.crt --key /etc/kubernetes/certs/etcdclient.key ${etcd_client_url}/v2/machines || exit {{GetCSEErrorCode "ERR_ETCD_RUNNING_TIMEOUT"}}
  wait_for_file 1200 1 /etc/systemd/system/etcd-monitor.service || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  systemctlEnableAndStart etcd-monitor || exit {{GetCSEErrorCode "ERR_SYSTEMCTL_START_FAIL"}}
}
createKubeManifestDir() {
  mkdir -p /etc/kubernetes/manifests
}
writeKubeConfig() {
  local d=/home/$ADMINUSER/.kube
  local f=$d/config
{{- if HasBlockOutboundInternet}}
  local server=https://localhost
{{else}}
  local server=$KUBECONFIG_SERVER
{{- end}}
  mkdir -p $d
  touch $f
  chown $ADMINUSER:$ADMINUSER $d $f
  chmod 700 $d
  chmod 600 $f
  set +x
  echo "
---
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: \"$CA_CERTIFICATE\"
    server: $server
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
" >$f
  set -x
}
{{- if IsClusterAutoscalerAddonEnabled}}
configClusterAutoscalerAddon() {
  local f=$ADDONS_DIR/cluster-autoscaler.yaml
  wait_for_file 1200 1 $f || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  sed -i "s|<clientID>|$(echo $SERVICE_PRINCIPAL_CLIENT_ID | base64)|g" $f
  sed -i "s|<clientSec>|$(echo $SERVICE_PRINCIPAL_CLIENT_SECRET | base64)|g" $f
  sed -i "s|<subID>|$(echo $SUBSCRIPTION_ID | base64)|g" $f
  sed -i "s|<tenantID>|$(echo $TENANT_ID | base64)|g" $f
  sed -i "s|<rg>|$(echo $RESOURCE_GROUP | base64)|g" $f
}
{{end}}
{{- if IsAzurePolicyAddonEnabled}}
configAzurePolicyAddon() {
  sed -i "s|<resourceId>|/subscriptions/$SUBSCRIPTION_ID/resourceGroups/$RESOURCE_GROUP|g" $ADDONS_DIR/azure-policy-deployment.yaml
}
{{end}}

configAddons() {
  {{if IsClusterAutoscalerAddonEnabled}}
  if [[ ${CLUSTER_AUTOSCALER_ADDON} == true ]]; then
    configClusterAutoscalerAddon
  fi
  {{end}}
  {{if IsAzurePolicyAddonEnabled}}
  configAzurePolicyAddon
  {{end}}
  {{- if and (not HasCustomPodSecurityPolicy) IsPodSecurityPolicyAddonEnabled}}
  wait_for_file 1200 1 $POD_SECURITY_POLICY_SPEC || exit {{GetCSEErrorCode "ERR_FILE_WATCH_TIMEOUT"}}
  mkdir -p $ADDONS_DIR/init && cp $POD_SECURITY_POLICY_SPEC $ADDONS_DIR/init/ || exit {{GetCSEErrorCode "ERR_ADDONS_START_FAIL"}}
  {{- end}}
}
{{- if and HasNSeriesSKU IsNvidiaDevicePluginAddonEnabled}}
{{- /* installNvidiaDrivers is idempotent, it will uninstall itself if it is already installed, and then install anew */}}
installNvidiaDrivers() {
  local d="/var/lib/dkms/nvidia/${GPU_DV}" k log_file="/var/log/nvidia-installer-$(date +%s).log"
  k=$(uname -r)
  if [ -d $d ]; then
    dkms remove -m nvidia -v $GPU_DV -k $k
  fi
  sh $GPU_DEST/nvidia-drivers-$GPU_DV -s -k=$k --log-file-name=$log_file -a --no-drm --dkms --utility-prefix="${GPU_DEST}" --opengl-prefix="${GPU_DEST}"
}
configGPUDrivers() {
  rmmod nouveau
  echo blacklist nouveau >>/etc/modprobe.d/blacklist.conf
  retrycmd_no_stats 120 5 25 update-initramfs -u || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_CONFIG"}}
  wait_for_apt_locks
  dpkg -i $(ls ${APT_CACHE_DIR}libnvidia-container1*) || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_CONFIG"}}
  dpkg -i $(ls ${APT_CACHE_DIR}libnvidia-container-tools*) || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_CONFIG"}}
  dpkg -i $(ls ${APT_CACHE_DIR}nvidia-container-toolkit*) || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_CONFIG"}}
  dpkg -i $(ls ${APT_CACHE_DIR}nvidia-container-runtime*) || exit {{GetCSEErrorCode "ERR_GPU_DRIVERS_CONFIG"}}
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

  local packages="make gcc dkms" oe_dir="/opt/azure/containers/oe"
  wait_for_apt_locks
  retrycmd 30 5 3600 apt-get -y install "$packages" || exit 90
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
  {{- if NeedsContainerd}}
  docker rmi -f $(docker images -a -q) &
  {{else}}
  docker rmi $(docker images --format '{{OpenBraces}}.Repository{{CloseBraces}}:{{OpenBraces}}.Tag{{CloseBraces}}' | grep -vE "${ETCD_VERSION}$|${ETCD_VERSION}-|${ETCD_VERSION}_" | grep 'etcd') &
  docker rmi $(docker images --format '{{OpenBraces}}.Repository{{CloseBraces}}:{{OpenBraces}}.Tag{{CloseBraces}}' | grep -vE "${KUBERNETES_VERSION}$|${KUBERNETES_VERSION}-|${KUBERNETES_VERSION}_" | grep 'kube-proxy') &
  docker rmi $(docker images --format '{{OpenBraces}}.Repository{{CloseBraces}}:{{OpenBraces}}.Tag{{CloseBraces}}' | grep -vE "${KUBERNETES_VERSION}$|${KUBERNETES_VERSION}-|${KUBERNETES_VERSION}_" | grep 'kube-controller-manager') &
  docker rmi $(docker images --format '{{OpenBraces}}.Repository{{CloseBraces}}:{{OpenBraces}}.Tag{{CloseBraces}}' | grep -vE "${KUBERNETES_VERSION}$|${KUBERNETES_VERSION}-|${KUBERNETES_VERSION}_" | grep 'kube-apiserver') &
  docker rmi $(docker images --format '{{OpenBraces}}.Repository{{CloseBraces}}:{{OpenBraces}}.Tag{{CloseBraces}}' | grep -vE "${KUBERNETES_VERSION}$|${KUBERNETES_VERSION}-|${KUBERNETES_VERSION}_" | grep 'kube-scheduler') &
  docker rmi registry:2.7.1 &
  ctr -n=k8s.io image rm $(ctr -n=k8s.io images ls -q) &
  {{- end}}
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
exit_cse() {
  local exit_code=$1
  shift
  $@ >> {{GetLinuxCSELogPath}} &
  exit $exit_code
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
  local azure_json_path="/etc/kubernetes/azure.json"
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

    KUBERNETES_FILE_DIR=$(dirname "${azure_json_path}")
    K8S_CLIENT_CERT_PATH="${KUBERNETES_FILE_DIR}/k8s_auth_certificate.pfx"
    echo $SERVICE_PRINCIPAL_CLIENT_SECRET_CERT | base64 --decode >$K8S_CLIENT_CERT_PATH
    # shellcheck disable=SC2002,SC2005
    echo $(cat "${azure_json_path}" |
      jq --arg K8S_CLIENT_CERT_PATH ${K8S_CLIENT_CERT_PATH} '. + {aadClientCertPath:($K8S_CLIENT_CERT_PATH)}' |
      jq --arg SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD ${SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD} '. + {aadClientCertPassword:($SERVICE_PRINCIPAL_CLIENT_SECRET_PASSWORD)}' |
      jq 'del(.aadClientSecret)') >${azure_json_path}
  fi

  if [[ ${IDENTITY_SYSTEM,,} == "adfs" ]]; then
    # update the tenent id for ADFS environment.
    # shellcheck disable=SC2002,SC2005
    echo $(cat "${azure_json_path}" | jq '.tenantId = "adfs"') >${azure_json_path}
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

  SDN_INTERFACES=$(jq ".value | map(select(.properties != null) | select(.properties.macAddress != null) | select(.properties.macAddress | inside(\"${local_interfaces[*]}\"))) | map(select((.properties.ipConfigurations | length) > 0))" ${NETWORK_INTERFACES_FILE})

  if [[ -z $SDN_INTERFACES ]]; then
      echo "Error extracting the SDN interfaces from the network interfaces file"
      exit 123
  fi

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
FLATCAR_OS_NAME="FLATCAR"
DEBIAN_OS_NAME="DEBIAN"
if ! echo "${UBUNTU_OS_NAME} ${FLATCAR_OS_NAME} ${DEBIAN_OS_NAME}" | grep -q "${OS}"; then
  OS=$(sort -r /etc/*-release | gawk 'match($0, /^(ID_LIKE=(.*))$/, a) { print toupper(a[2] a[3]); exit }')
fi
if [[ ${OS} == "${UBUNTU_OS_NAME}" ]]; then
  UBUNTU_RELEASE=$(lsb_release -r -s)
fi
DOCKER=/usr/bin/docker
if [[ $UBUNTU_RELEASE == "18.04" ]]; then
  export GPU_DV=460.32.03
else
  export GPU_DV=418.40.04
fi
export GPU_DEST=/usr/local/nvidia
PRIVATE_IP=$( (ip -br -4 addr show eth0 || ip -br -4 addr show azure0) | grep -Po '\d+\.\d+\.\d+\.\d+')
if ! [[ $(echo -n "$PRIVATE_IP" | grep -c '^') == 1 ]]; then
  PRIVATE_IP=$(hostname -i)
fi
export PRIVATE_IP
APT_CACHE_DIR=/var/cache/apt/archives/

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
apt_get_download() {
  retries=$1; wait_sleep=$2; shift && shift;
  local ret=0
  pushd $APT_CACHE_DIR || return 1
  for i in $(seq 1 $retries); do
    wait_for_apt_locks; apt-get -o Dpkg::Options::=--force-confold download -y "${@}" && break
    if [ $i -eq $retries ]; then ret=1; else sleep $wait_sleep; fi
  done
  popd || return 1
  return $ret
}
dpkg_install() {
  retries=$1; wait_sleep=$2; shift && shift;
  for i in $(seq 1 $retries); do
    wait_for_apt_locks; dpkg -i "${1}" && break; apt-get update --fix-missing; do_apt_get_install -f && break
    if [ $i -eq $retries ]; then return 1; fi; sleep $wait_sleep
  done
}
do_apt_get_install() {
  dpkg --configure -a --force-confdef; DEBIAN_FRONTEND=noninteractive apt-get install -o Dpkg::Options::="--force-confold" --no-install-recommends -y ${@}
}
apt_get_install() {
  retries=$1; wait_sleep=$2; timeout=$3; shift && shift && shift
  for i in $(seq 1 $retries); do
    wait_for_apt_locks
    do_apt_get_install ${@} && break ||
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
unattended_upgrade() {
  retries=10
  for i in $(seq 1 $retries); do
    wait_for_apt_locks
    /usr/bin/unattended-upgrade && break ||
    if [ $i -eq $retries ]; then
      return 1
    else sleep 5
    fi
  done
  echo Executed unattended-upgrade $i times
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
systemctl_enable() {
  retries=$1; wait_sleep=$2; timeout=$3 svcname=$4
  for i in $(seq 1 $retries); do
    timeout $timeout systemctl daemon-reload
    timeout $timeout systemctl enable $svcname && break ||
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
APMZ_DOWNLOADS_DIR="/opt/apmz/downloads"
UBUNTU_RELEASE=$(lsb_release -r -s)
UBUNTU_CODENAME=$(lsb_release -c -s)

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
  apt_get_download 20 30 libnvidia-container1=1.4.0* libnvidia-container-tools=1.4.0* nvidia-container-toolkit=1.5.0* nvidia-container-runtime=3.5.0* || exit 85
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
  if [[ $v != "1.0.0" ]]; then
    apt_get_install 20 30 120 moby-runc=1.0.0* --allow-downgrades || exit 27
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

{{/* delete non-working iovisor definition to ensure apt operations work */}}
rm -Rf /etc/apt/sources.list.d/iovisor.list

set -x
if [ -f /opt/azure/containers/provision.complete ]; then
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

if [[ ${UBUNTU_RELEASE} == "18.04" ]]; then
  disable1804SystemdResolved
fi

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

VHD_LOGS_FILEPATH=/opt/azure/vhd-install.complete
if [ -f $VHD_LOGS_FILEPATH ]; then
  time_metric "CleanUpContainerImages" cleanUpContainerImages
  FULL_INSTALL_REQUIRED=false
else
  if [[ ${IS_VHD} == true ]]; then
    exit {{GetCSEErrorCode "ERR_VHD_FILE_NOT_FOUND"}}
  fi
  FULL_INSTALL_REQUIRED=true
fi

{{- if not IsVHDDistroForAllNodes}}
if [[ $OS == $UBUNTU_OS_NAME || $OS == $DEBIAN_OS_NAME ]] && [ "$FULL_INSTALL_REQUIRED" = "true" ]; then
  time_metric "InstallDeps" installDeps
  if [[ ${UBUNTU_RELEASE} == "18.04" ]]; then
    overrideNetworkConfig
  fi
  {{- if not IsDockerContainerRuntime}}
  time_metric "InstallImg" installImg
  {{end}}
fi
{{end}}

if [[ ${UBUNTU_RELEASE} == "18.04" ]]; then
  if apt list --installed | grep 'chrony'; then
    time_metric "ConfigureChrony" configureChrony
    time_metric "EnsureChrony" ensureChrony
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
time_metric "installRunc" installRunc
{{- if HasLinuxMobyURL}}
  LINUX_MOBY_URL={{GetLinuxMobyURL}}
  if [[ -n "${LINUX_MOBY_URL:-}" ]]; then
    DEB="${LINUX_MOBY_URL##*/}"
    retrycmd_no_stats 120 5 25 curl -fsSL ${LINUX_MOBY_URL} >/tmp/${DEB} || exit {{GetCSEErrorCode "ERR_DEB_DOWNLOAD_TIMEOUT"}}
    dpkg_install 20 30 /tmp/${DEB} || exit {{GetCSEErrorCode "ERR_DEB_PKG_ADD_FAIL"}}
  fi
{{end}}
{{- if HasLinuxContainerdURL}}
  LINUX_CONTAINERD_URL={{GetLinuxContainerdURL}}
  if [[ -n "${LINUX_CONTAINERD_URL:-}" ]]; then
    DEB="${LINUX_CONTAINERD_URL##*/}"
    retrycmd_no_stats 120 5 25 curl -fsSL ${LINUX_CONTAINERD_URL} >/tmp/${DEB} || exit {{GetCSEErrorCode "ERR_DEB_DOWNLOAD_TIMEOUT"}}
    dpkg_install 20 30 /tmp/${DEB} || exit {{GetCSEErrorCode "ERR_DEB_PKG_ADD_FAIL"}}
  fi
{{end}}
{{- if HasLinuxRuncURL}}
  LINUX_RUNC_URL={{GetLinuxRuncURL}}
  if [[ -n "${LINUX_RUNC_URL:-}" ]]; then
    DEB="${LINUX_RUNC_URL##*/}"
    retrycmd_no_stats 120 5 25 curl -fsSL ${LINUX_RUNC_URL} >/tmp/${DEB} || exit {{GetCSEErrorCode "ERR_DEB_DOWNLOAD_TIMEOUT"}}
    dpkg_install 20 30 /tmp/${DEB} || exit {{GetCSEErrorCode "ERR_DEB_PKG_ADD_FAIL"}}
  fi
{{end}}
fi

if [[ -n ${MASTER_NODE} ]] && [[ -z ${COSMOS_URI} ]]; then
  {{- if IsDockerContainerRuntime}}
  cli_tool="docker"
  {{else}}
  cli_tool="img"
  {{end}}
  time_metric "InstallEtcd" installEtcd $cli_tool
fi

{{/* this will capture the amount of time to install of the network plugin during cse */}}
time_metric "InstallNetworkPlugin" installNetworkPlugin

{{- if and HasNSeriesSKU IsNvidiaDevicePluginAddonEnabled}}
if [[ ${GPU_NODE} == true ]]; then
  time_metric "DownloadGPUDrivers" downloadGPUDrivers
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

{{/* configure and enable dhcpv6 for ipv6 features */}}
{{- if IsIPv6Enabled}}
time_metric "EnsureDHCPv6" ensureDHCPv6
{{end}}

{{/* configure and enable kms plugin */}}
{{- if EnableEncryptionWithExternalKms}}
if [[ -n ${MASTER_NODE} ]]; then
  time_metric "EnsureKMSKeyvaultKey" ensureKMSKeyvaultKey
fi
{{end}}

time_metric "EnsureKubelet" ensureKubelet
{{if IsAzurePolicyAddonEnabled}}
if [[ -n ${MASTER_NODE} ]]; then
  time_metric "EnsureLabelExclusionForAzurePolicyAddon" ensureLabelExclusionForAzurePolicyAddon
fi
{{end}}
time_metric "EnsureJournal" ensureJournal

if [[ -n ${MASTER_NODE} ]]; then
  if version_gte ${KUBERNETES_VERSION} 1.16; then
    time_metric "EnsureLabelNodes" ensureLabelNodes
  fi
{{- if IsAADPodIdentityAddonEnabled}}
  time_metric "EnsureTaints" ensureTaints
{{end}}
  if [[ -z ${COSMOS_URI} ]]; then
    time_metric "EnsureEtcd" ensureEtcd
  fi
  time_metric "EnsureK8sControlPlane" ensureK8sControlPlane
  time_metric "EnsureAddons" ensureAddons
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
{{- if EnableUnattendedUpgrades}}
rm -f /etc/apt/apt.conf.d/99periodic
{{- end}}

{{- if not IsAzureStackCloud}}
if [[ $OS == $UBUNTU_OS_NAME ]]; then
  time_metric "PurgeApt" apt_get_purge apache2-utils &
fi
{{end}}

{{- if not HasBlockOutboundInternet}}
    {{- if RunUnattendedUpgradesOnBootstrap}}
apt_get_update && unattended_upgrade
    {{- end}}
{{- end}}

if [ -f /var/run/reboot-required ]; then
  trace_info "RebootRequired" "reboot=true"
  /bin/bash -c "shutdown -r 1 &"
  if [[ $OS == $UBUNTU_OS_NAME ]]; then
    aptmarkWALinuxAgent unhold &
  fi
else
{{- if RunUnattendedUpgradesOnBootstrap}}
  if [[ -z ${MASTER_NODE} ]]; then
    systemctl_restart 100 5 30 kubelet
    systemctl_restart 100 5 30 kubelet-monitor
  fi
{{- end}}
  if [[ $OS == $UBUNTU_OS_NAME ]]; then
    /usr/lib/apt/apt.systemd.daily &
    aptmarkWALinuxAgent unhold &
  fi
fi

echo "CSE finished successfully"
echo $(date),$(hostname), endcustomscript >>/opt/m
mkdir -p /opt/azure/containers && touch /opt/azure/containers/provision.complete
ps auxfww >/opt/azure/provision-ps.log &

exit 0

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
After={{GetContainerRuntime}}.service
[Service]
Restart=always
RestartSec=10
RemainAfterExit=yes
Environment=CONTAINER_RUNTIME={{GetContainerRuntime}}
ExecStart=/usr/local/bin/health-monitor.sh container-runtime
[Install]
WantedBy=multi-user.target
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

touch /etc/dhcp/dhclient6.conf && add_if_not_exists "timeout 10;" ${DHCLIENT6_CONF_FILE} && \
  add_if_not_exists "${NETWORK_CONFIGURATION}" ${CLOUD_INIT_CFG} && \
  sudo ifdown eth0 && sudo ifup eth0
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

var _k8sCloudInitArtifactsEtcdMonitorService = []byte(`[Unit]
Description=a script that checks etcd health and restarts if needed
After=etcd.service
[Service]
Restart=always
RestartSec=10
RemainAfterExit=yes
ExecStart=/usr/local/bin/health-monitor.sh etcd
[Install]
WantedBy=multi-user.target
#EOF
`)

func k8sCloudInitArtifactsEtcdMonitorServiceBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsEtcdMonitorService, nil
}

func k8sCloudInitArtifactsEtcdMonitorService() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsEtcdMonitorServiceBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/etcd-monitor.service", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsEtcdService = []byte(`[Unit]
Description=etcd - highly-available key value store
Documentation=https://github.com/coreos/etcd
Documentation=man:etcd
After=network.target var-lib-etcddisk.mount
Wants=network-online.target var-lib-etcddisk.mount
[Service]
Environment=DAEMON_ARGS=
Environment=ETCD_NAME=%H
Environment=ETCD_DATA_DIR=
EnvironmentFile=-/etc/default/%p
Type=notify
User=etcd
PermissionsStartOnly=true
ExecStartPre=/bin/mountpoint -q /var/lib/etcddisk
ExecStart=/usr/bin/etcd $DAEMON_ARGS
Restart=always
[Install]
WantedBy=multi-user.target
`)

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
openssl req -new -x509 -days 262824 -key $PROXY_CA_KEY -out $PROXY_CRT -subj '/CN=proxyClientCA'
openssl genrsa -out $PROXY_CLIENT_KEY 2048
openssl req -new -key $PROXY_CLIENT_KEY -out $PROXY_CLIENT_CSR -subj '/CN=aggregator/O=system:masters'
openssl x509 -req -days 262824 -in $PROXY_CLIENT_CSR -CA $PROXY_CRT -CAkey $PROXY_CA_KEY -set_serial 02 -out $PROXY_CLIENT_CRT

write_certs_to_disk() {
  ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} get $ETCD_REQUESTHEADER_CLIENT_CA --print-value-only >$K8S_PROXY_CA_CRT_FILEPATH
  ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} get $ETCD_PROXY_KEY --print-value-only >$K8S_PROXY_KEY_FILEPATH
  ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} get $ETCD_PROXY_CERT --print-value-only >$K8S_PROXY_CRT_FILEPATH
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
rm -f "${PROXY_CERT_LOCK_FILE}"
mkfifo "${PROXY_CERT_LOCK_FILE}"

ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} lock ${PROXY_CERTS_LOCK_NAME} >"${PROXY_CERT_LOCK_FILE}" &

pid=$!
if read -r lockthis <"${PROXY_CERT_LOCK_FILE}"; then
  if [[ "${OVERRIDE_PROXY_CERTS}" == "true" ]] || [[ "" == "$(ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} get $ETCD_REQUESTHEADER_CLIENT_CA --print-value-only)" ]]; then
    ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} put $ETCD_REQUESTHEADER_CLIENT_CA " $(cat ${PROXY_CRT})" >/dev/null 2>&1
  fi
  if [[ "${OVERRIDE_PROXY_CERTS}" == "true" ]] || [[ "" == "$(ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} get $ETCD_PROXY_KEY --print-value-only)" ]]; then
    ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} put $ETCD_PROXY_KEY " $(cat ${PROXY_CLIENT_KEY})" >/dev/null 2>&1
  fi
  if [[ "${OVERRIDE_PROXY_CERTS}" == "true" ]] || [[ "" == "$(ETCDCTL_API=3 etcdctl ${ETCDCTL_PARAMS} get $ETCD_PROXY_CERT --print-value-only)" ]]; then
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
source {{GetCSEHelpersScriptFilepath}}

{{- /* This script originated at https://github.com/kubernetes/kubernetes/blob/master/cluster/gce/gci/health-monitor.sh */}}
{{- /* and has been modified for aks-engine. */}}

set -o nounset
set -o pipefail

container_runtime_monitoring() {
  sleep 300 {{/* Wait for 5 minutes for CRI to be functional/stable */}}
  local cri_name="${CONTAINER_RUNTIME:-docker}" cmd="docker ps"
  if [[ ${CONTAINER_RUNTIME} == "containerd" ]]; then
    cmd="ctr -n k8s.io containers ls"
  fi

  while true; do
    if ! timeout 60 ${cmd} >/dev/null; then
      echo "Container runtime ${cri_name} failed!"
      sleep 10 {{/* Wait 10 more seconds, check again, because the systemd job itself may have already restarted things */}}
      if ! timeout 60 ${cmd} >/dev/null; then
        if [[ $cri_name == "docker" ]]; then
          pkill -SIGUSR1 dockerd
        fi
        if [[ $cri_name == "containerd" ]]; then
          pkill -SIGUSR1 containerd
        fi
        systemctl kill --kill-who=main "${cri_name}"
        sleep 60 {{/* Wait a minute to validate that the systemd job restarted itself after we manually killed the process */}}
        if ! systemctl is-active ${cri_name}; then
          systemctl start ${cri_name}
        fi
      fi
    else
      sleep "${SLEEP_TIME}"
    fi
  done
}

kubelet_monitoring() {
  sleep 300 {{/* Wait for 5 minutes for kubelet to be functional/stable */}}
  local max_seconds=10 output=""
  local monitor_cmd="curl -m ${max_seconds} -f -s -S http://127.0.0.1:${HEALTHZPORT}/healthz"
  while true; do
    if ! output=$(${monitor_cmd} 2>&1); then
      echo $output
      echo "Kubelet is unhealthy!"
      sleep 10 {{/* Wait 10 more seconds, check again, because the systemd job itself may have already restarted things */}}
      if ! output=$(${monitor_cmd} 2>&1); then
        systemctl kill kubelet
        sleep 60 {{/* Wait a minute to validate that the systemd job restarted itself after we manually killed the process */}}
        if ! systemctl is-active kubelet; then
          systemctl start kubelet
        fi
      fi
    else
      sleep "${SLEEP_TIME}"
    fi
  done
}

apiserver_monitoring() {
  sleep 300 {{/* Wait for 5 minutes for apiserver to be functional/stable */}}
  local private_ip=$( (ip -br -4 addr show eth0 || ip -br -4 addr show azure0) | grep -Po '\d+\.\d+\.\d+\.\d+')
  local max_seconds=10 output=""
  local monitor_cmd="curl -m ${max_seconds} --cacert /etc/kubernetes/certs/ca.crt --cert /etc/kubernetes/certs/client.crt --key /etc/kubernetes/certs/client.key https://${private_ip}:443/readyz?verbose"
  while true; do
    if ! output=$(${monitor_cmd} 2>&1); then
      echo $output
      echo "apiserver is unhealthy!"
      sleep 10 {{/* Wait 10 more seconds, check again, because the systemd job itself may have already restarted things */}}
      if ! output=$(${monitor_cmd} 2>&1); then
        systemctl kill kubelet
        sleep 60 {{/* Wait a minute to validate that the systemd job restarted itself after we manually killed the process */}}
        if ! systemctl is-active kubelet; then
          systemctl start kubelet
        fi
      fi
    else
      sleep "${SLEEP_TIME}"
    fi
  done
}

etcd_monitoring() {
  sleep 300 {{/* Wait for 5 minutes for etcd to be functional/stable */}}
  local max_seconds=10 output=""
  local endpoint="https://${PRIVATE_IP}:2379"
  local monitor_cmd="curl -s -S -m ${max_seconds} --cacert /etc/kubernetes/certs/ca.crt --cert /etc/kubernetes/certs/etcdclient.crt --key /etc/kubernetes/certs/etcdclient.key ${endpoint}/v2/machines"
  while true; do
    if ! output=$(${monitor_cmd}); then
      echo $output
      echo "etcd is unhealthy!"
      sleep 10 {{/* Wait 10 more seconds, check again, because the systemd job itself may have already restarted things */}}
      if ! output=$(${monitor_cmd}); then
        systemctl kill etcd
        sleep 60 {{/* Wait a minute to validate that the systemd job restarted itself after we manually killed the process */}}
        if ! systemctl is-active etcd; then
          systemctl start etcd
        fi
      fi
    else
      sleep "${SLEEP_TIME}"
    fi
  done
}

if [[ $# -ne 1 ]]; then
  echo "Usage: health-monitor.sh <container-runtime|kubelet|etcd>"
  exit 1
fi

SLEEP_TIME=10
component=$1
echo "Start kubernetes health monitoring for ${component}"

if [[ ${component} == "container-runtime" ]]; then
  container_runtime_monitoring
elif [[ ${component} == "kubelet" ]]; then
  kubelet_monitoring
elif [[ ${component} == "apiserver" ]]; then
  apiserver_monitoring
elif [[ ${component} == "etcd" ]]; then
  etcd_monitoring
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

var _k8sCloudInitArtifactsKmsKeyvaultKeyService = []byte(`[Unit]
Description=setupkmskey
After=network-online.target

[Service]
Type=oneshot
ExecStart={{GetKMSKeyvaultKeyCSEScriptFilepath}}

[Install]
WantedBy=multi-user.target
#EOF
`)

func k8sCloudInitArtifactsKmsKeyvaultKeyServiceBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsKmsKeyvaultKeyService, nil
}

func k8sCloudInitArtifactsKmsKeyvaultKeyService() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsKmsKeyvaultKeyServiceBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/kms-keyvault-key.service", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsKmsKeyvaultKeySh = []byte(`#!/usr/bin/env bash

set +x
set -euo pipefail

AZURE_JSON_PATH="/etc/kubernetes/azure.json"
SERVICE_PRINCIPAL_CLIENT_ID=$(jq -r '.aadClientId' ${AZURE_JSON_PATH})
SERVICE_PRINCIPAL_CLIENT_SECRET=$(jq -r '.aadClientSecret' ${AZURE_JSON_PATH})
TENANT_ID=$(jq -r '.tenantId' ${AZURE_JSON_PATH})
KMS_KEYVAULT_NAME=$(jq -r '.providerVaultName' ${AZURE_JSON_PATH})
KMS_KEY_NAME=$(jq -r '.providerKeyName' ${AZURE_JSON_PATH})
USER_ASSIGNED_IDENTITY_ID=$(jq -r '.userAssignedIdentityID' ${AZURE_JSON_PATH})
PROVIDER_KEY_VERSION=$(jq -r '.providerKeyVersion' ${AZURE_JSON_PATH})
AZURE_CLOUD=$(jq -r '.cloud' ${AZURE_JSON_PATH})

# get the required parameters specific for cloud
if [[ $AZURE_CLOUD == "AzurePublicCloud" ]]; then
    ACTIVE_DIRECTORY_ENDPOINT="https://login.microsoftonline.com/"
    KEYVAULT_DNS_SUFFIX="vault.azure.net"
elif [[ $AZURE_CLOUD == "AzureChinaCloud" ]]; then
    ACTIVE_DIRECTORY_ENDPOINT="https://login.chinacloudapi.cn/"
    KEYVAULT_DNS_SUFFIX="vault.azure.cn"
elif [[ $AZURE_CLOUD == "AzureGermanCloud" ]]; then
    ACTIVE_DIRECTORY_ENDPOINT="https://login.microsoftonline.de/"
    KEYVAULT_DNS_SUFFIX="vault.microsoftazure.de"
elif [[ $AZURE_CLOUD == "AzureUSGovernmentCloud" ]]; then
    ACTIVE_DIRECTORY_ENDPOINT="https://login.microsoftonline.us/"
    KEYVAULT_DNS_SUFFIX="vault.usgovcloudapi.net"
elif [[ $AZURE_CLOUD == "AzureStackCloud" ]]; then
    AZURESTACK_ENVIRONMENT_JSON_PATH="/etc/kubernetes/azurestackcloud.json"
    ACTIVE_DIRECTORY_ENDPOINT=$(jq -r '.activeDirectoryEndpoint' ${AZURESTACK_ENVIRONMENT_JSON_PATH})
    KEYVAULT_DNS_SUFFIX=$(jq -r '.keyVaultDNSSuffix' ${AZURESTACK_ENVIRONMENT_JSON_PATH})
else
    echo "Invalid cloud name"
    exit 120
fi

TOKEN_URL="${ACTIVE_DIRECTORY_ENDPOINT}${TENANT_ID}/oauth2/token"
KEYVAULT_URL="https://${KMS_KEYVAULT_NAME}.${KEYVAULT_DNS_SUFFIX}/keys/${KMS_KEY_NAME}/versions?maxresults=1&api-version=7.1"
KEYVAULT_ENDPOINT="https://${KEYVAULT_DNS_SUFFIX}"
KMS_KUBERNETES_FILE=/etc/kubernetes/manifests/kube-azure-kms.yaml

# provider key version already exists
# this will be the case for BYOK
if [[ -n $PROVIDER_KEY_VERSION ]]; then
    echo "KMS provider key version already exists"
    exit 0
fi

echo "Generating token for Azure Key Vault"
echo "------------------------------------------------------------------------"
echo "Parameters"
echo "------------------------------------------------------------------------"
echo "SERVICE_PRINCIPAL_CLIENT_ID:     ..."
echo "SERVICE_PRINCIPAL_CLIENT_SECRET: ..."
echo "ACTIVE_DIRECTORY_ENDPOINT:       $ACTIVE_DIRECTORY_ENDPOINT"
echo "TENANT_ID:                       $TENANT_ID"
echo "TOKEN_URL:                       $TOKEN_URL"
echo "SCOPE:                           $KEYVAULT_ENDPOINT"
echo "------------------------------------------------------------------------"

if [[ $SERVICE_PRINCIPAL_CLIENT_ID == "msi" ]] && [[ $SERVICE_PRINCIPAL_CLIENT_SECRET == "msi" ]]; then
    if [[ -z $USER_ASSIGNED_IDENTITY_ID ]]; then
        # using system-assigned identity to access keyvault
        TOKEN=$(curl -s --retry 5 --retry-delay 10 --max-time 60 \
            -H Metadata:true \
            "http://169.254.169.254/metadata/identity/oauth2/token?api-version=2018-02-01&resource=$KEYVAULT_ENDPOINT" | jq '.access_token' | xargs)
    else
        # using user-assigned managed identity to access keyvault
        TOKEN=$(curl -s --retry 5 --retry-delay 10 --max-time 60 \
            -H Metadata:true \
            "http://169.254.169.254/metadata/identity/oauth2/token?api-version=2018-02-01&client_id=$USER_ASSIGNED_IDENTITY_ID&resource=$KEYVAULT_ENDPOINT" | jq '.access_token' | xargs)
    fi
else
    # use service principal token to access keyvault
    TOKEN=$(curl -s --retry 5 --retry-delay 10 --max-time 60 -f -X POST \
        -H "Content-Type: application/x-www-form-urlencoded" \
        -d "grant_type=client_credentials" \
        -d "client_id=$SERVICE_PRINCIPAL_CLIENT_ID" \
        --data-urlencode "client_secret=$SERVICE_PRINCIPAL_CLIENT_SECRET" \
        --data-urlencode "resource=$KEYVAULT_ENDPOINT" \
        ${TOKEN_URL} | jq '.access_token' | xargs)
fi


if [[ -z $TOKEN ]]; then
    echo "Error generating token for Azure Keyvault"
    exit 120
fi

# Get the keyID for the kms key created as part of cluster bootstrap
KEY_ID=$(curl -s --retry 5 --retry-delay 10 --max-time 60 -f \
    ${KEYVAULT_URL} -H "Authorization: Bearer ${TOKEN}" | jq '.value[0].kid' | xargs)

if [[ -z "$KEY_ID" || "$KEY_ID" == "null" ]]; then
    echo "Error getting the kms key version"
    exit 120
fi

# KID format: https://<keyvault name>.vault.azure.net/keys/<key name>/<key version>
# Example KID: "https://akv0112master.vault.azure.net/keys/k8s/128a3d9956bc44feb6a0e2c2f35b732c"
KEY_VERSION=${KEY_ID##*/}

# Set the version in azure.json
if [ -f $AZURE_JSON_PATH ]; then
    # once the version is set in azure.json, kms plugin will just default to using the key
    # this will be changed in upcoming kms release to set the version as container args
    tmpDir=$(mktemp -d "$(pwd)/XXX")
    jq --arg KEY_VERSION ${KEY_VERSION} '.providerKeyVersion=($KEY_VERSION)' "$AZURE_JSON_PATH" > $tmpDir/tmp
    mv $tmpDir/tmp "$AZURE_JSON_PATH"
    # set the permissions for azure json
    chmod 0600 "$AZURE_JSON_PATH"
    chown root:root "$AZURE_JSON_PATH"
    rm -Rf $tmpDir
fi

set -x
#EOF
`)

func k8sCloudInitArtifactsKmsKeyvaultKeyShBytes() ([]byte, error) {
	return _k8sCloudInitArtifactsKmsKeyvaultKeySh, nil
}

func k8sCloudInitArtifactsKmsKeyvaultKeySh() (*asset, error) {
	bytes, err := k8sCloudInitArtifactsKmsKeyvaultKeyShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/cloud-init/artifacts/kms-keyvault-key.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sCloudInitArtifactsKubeletMonitorService = []byte(`[Unit]
Description=a script that checks kubelet health and restarts if needed
After=kubelet.service
[Service]
Environment=HEALTHZPORT={{GetKubeletHealthZPort}}
Restart=always
RestartSec=10
RemainAfterExit=yes
ExecStart=/usr/local/bin/health-monitor.sh kubelet
[Install]
WantedBy=multi-user.target
#EOF
`)

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
Requires={{GetContainerRuntime}}.service

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
#EOF
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
- source {{GetCSEHelpersScriptFilepath}}
- retrycmd 10 5 10 curl -LO https://storage.googleapis.com/kubernetes-release/release/v{{.OrchestratorProfile.OrchestratorVersion}}/bin/linux/amd64/kubectl
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

{{- if .MasterProfile.IsAuditDEnabled}}
- path: /etc/audit/rules.d/CIS.rules
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "auditdRules"}}
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

- path: /etc/systemd/system/{{GetContainerRuntime}}.service.d/kubereserved-slice.conf
  permissions: "0644"
  owner: root
  content: |
    [Service]
    Slice={{- GetKubeReservedCgroup -}}.slice
    #EOF
{{- end}}

- path: /usr/local/bin/health-monitor.sh
  permissions: "0544"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "healthMonitorScript"}}

{{- if HasKubeletHealthZPort}}
- path: /etc/systemd/system/kubelet-monitor.service
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "kubeletMonitorSystemdService"}}
{{- end}}

- path: /etc/systemd/system/apiserver-monitor.service
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "apiserverMonitorSystemdService"}}

- path: /etc/systemd/system/etcd-monitor.service
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "etcdMonitorSystemdService"}}

- path: /etc/systemd/system/kubelet.service
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "kubeletSystemdService"}}

{{- /* for historical reasons, we overload the name "docker" here; in fact this monitor service supports both docker and containerd */}}
- path: /etc/systemd/system/docker-monitor.service
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "dockerMonitorSystemdService"}}

{{- if not .MasterProfile.IsVHDDistro}}
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

- path: /etc/apt/preferences
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "aptPreferences"}}
    {{- if EnableAggregatedAPIs}}
- path: /etc/kubernetes/generate-proxy-certs.sh
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "generateProxyCertsScript"}}
    {{end}}
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

{{- if and IsIPv6Enabled .MasterProfile.IsUbuntu1604}}
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
    Restart=always
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
          },
          {
            "type": "portmap",
            "capabilities": {"portMappings": true},
            "snat": false
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

- path: {{GetKMSKeyvaultKeyServiceCSEScriptFilepath}}
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "kmsKeyvaultKeySystemdService"}}

- path: {{GetKMSKeyvaultKeyCSEScriptFilepath}}
  permissions: "0544"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "kmsKeyvaultKeyScript"}}
{{end}}

MASTER_MANIFESTS_CONFIG_PLACEHOLDER

MASTER_CUSTOM_FILES_PLACEHOLDER

MASTER_CONTAINER_ADDONS_PLACEHOLDER

{{- if or (IsDashboardAddonEnabled) (IsAzurePolicyAddonEnabled)}} {{/* Note: dashboard addon is deprecated */}}
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
    KUBELET_NODE_LABELS={{GetMasterKubernetesLabels "',variables('labelResourceGroup'),'"}}
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
{{- if IsAzureCNI}}
    ifconfig eth0 mtu {{GetEth0MTU}} up
{{- end}}
{{- if gt .MasterProfile.Count 1}}
    {{- /* Redirect ILB (4443) traffic to port 443 (ELB) in the prerouting chain */}}
    iptables -t nat -A PREROUTING -p tcp --dport 4443 -j REDIRECT --to-port 443
{{end}}
{{- if EnableDataEncryptionAtRest }}
    sed -i "s|<etcdEncryptionSecret>|\"{{WrapAsParameter "etcdEncryptionKey"}}\"|g" /etc/kubernetes/encryption-config.yaml
{{end}}
    {{- /* Ensure that container traffic can't connect to internal Azure IP endpoint */}}
    iptables -I FORWARD -d 168.63.129.16 -p tcp --dport 80 -j DROP
    {{- /* Ensure that we have access to the upstream DNS via TCP */}}
    {{- /* There is a long-standing bug where TCP is blocked for all ports */}}
    {{- /* thus impacting DNS via TCP - https://github.com/Azure/WALinuxAgent/issues/1673 */}}
    {{- /* This works around that bug and will not have negative impacts if it gets fixed */}}
    iptables -t security -D OUTPUT -d 168.63.129.16/32 -p tcp -m tcp --dport 53 -j ACCEPT || true
    iptables -t security -I OUTPUT 1 -d 168.63.129.16/32 -p tcp -m tcp --dport 53 -j ACCEPT
    {{- /* The two lines are (1) to delete the rule if it is there already and     */}}
    {{- /* then (2) to add the rule at the top (first entry) which specifically    */}}
    {{- /* allows TCP to the magic VM host controlled IP address for DNS (which is */}}
    {{- /* the upstream DNS provider)                                              */}}
    {{- /* Why the delete? Well, if the kubelet gets restarted, we don't want to   */}}
    {{- /* add the rule again. The -I will add the rule again, even if it is there */}}
    {{- /* already. So, we delete the rule first. It also makes sure that the      */}}
    {{- /* rule is added at the top of the chain such that this will also          */}}
    {{- /* effectively move it to the top if it was already there. In other words, */}}
    {{- /* it just makes sure that the rule exists and is at the right spot. If it */}}
    {{- /* does not exist, the delete (-D) will "fail" to delete the non-existing  */}}
    {{- /* rule, which is the end state we want so we don't care if the delete     */}}
    {{- /* failed and thus the "|| true". Then the second step is to insert the    */}}
    {{- /* rule at the top of the chain, thus making sure that no other rule       */}}
    {{- /* blocks access before it can do the ACCEPT                               */}}
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
    source {{GetCSEHelpersScriptFilepath}}
    MASTER_VM_NAME=$(hostname)
    MASTER_VM_NAME_BASE=$(hostname | sed "s/.$//")
    MASTER_FIRSTADDR={{WrapAsParameter "firstConsecutiveStaticIP"}}
    MASTER_INDEX=$(hostname | tail -c 2)
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
- source {{GetCSEHelpersScriptFilepath}}
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
- path: /opt/azure/needs_azure.json
  permissions: "0644"
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

{{- if .IsAuditDEnabled}}
- path: /etc/audit/rules.d/CIS.rules
  permissions: "0744"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "auditdRules"}}
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

- path: /etc/systemd/system/{{GetContainerRuntime}}.service.d/kubereserved-slice.conf
  permissions: "0644"
  owner: root
  content: |
    [Service]
    Slice={{- GetKubeReservedCgroup -}}.slice
    #EOF
{{- end}}

{{- if HasKubeletHealthZPort}}
- path: /etc/systemd/system/kubelet-monitor.service
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "kubeletMonitorSystemdService"}}
{{- end}}

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

- path: /etc/systemd/system/kubelet.service
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "kubeletSystemdService"}}

{{- /* for historical reasons, we overload the name "docker" here; in fact this monitor service supports both docker and containerd */}}
- path: /etc/systemd/system/docker-monitor.service
  permissions: "0644"
  encoding: gzip
  owner: root
  content: !!binary |
    {{CloudInitData "dockerMonitorSystemdService"}}

{{- if not .IsVHDDistro}}
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

{{- if and IsIPv6Enabled .IsUbuntu1604}}
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
    Restart=always
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
  {{- if IsNSeriesSKU .VMSize}}
{{IndentString GetNvidiaContainerdConfig 4}}
  {{else}}
{{IndentString GetContainerdConfig 4}}
  {{- end}}
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
          },
          {
            "type": "portmap",
            "capabilities": {"portMappings": true},
            "snat": false
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
    KUBELET_NODE_LABELS={{GetAgentKubernetesLabels . "',variables('labelResourceGroup'),'"}}
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
{{- if IsAzureCNI}}
    ifconfig eth0 mtu {{GetEth0MTU}} up
{{- end}}
{{- if and (IsVirtualMachineScaleSets .) IsAADPodIdentityAddonEnabled UseManagedIdentity}}
    {{- /* Disable TCP access to IMDS endpoint, aad-pod-identity nmi component will provide a complementary iptables rule to re-route this traffic */}}
    iptables -A OUTPUT -s 127.0.0.1/32 -d 169.254.169.254/32 -p tcp -m tcp --dport 80 -j DROP
{{end}}
{{- if not IsIPMasqAgentEnabled}}
    {{if IsAzureCNI}}
    iptables -t nat -A POSTROUTING -m iprange ! --dst-range 168.63.129.16 -m addrtype ! --dst-type local ! -d {{WrapAsParameter "vnetCidr"}} -j MASQUERADE
    {{end}}
{{end}}
    {{- /* Ensure that container traffic can't connect to internal Azure IP endpoint */}}
    iptables -I FORWARD -d 168.63.129.16 -p tcp --dport 80 -j DROP
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
            After=kubelet.service
            [Service]
            ExecStart=
            ExecStart=/opt/bin/health-monitor.sh kubelet
    - name: docker-monitor.service
      enable: true
      drop-ins:
        - name: "10-flatcar.conf"
          content: |
            After={{GetContainerRuntime}}.service
            [Service]
            ExecStart=
            ExecStart=/opt/bin/health-monitor.sh container-runtime
    - name: rpcbind.service
      enable: true
{{else}}
runcmd:
- set -x
- source {{GetCSEHelpersScriptFilepath}}
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

var _k8sContainerdtemplateToml = []byte(`root = "C:\\ProgramData\\containerd\\root"
state = "C:\\ProgramData\\containerd\\state"

[grpc]
  address = "\\\\.\\pipe\\containerd-containerd"
  max_recv_message_size = 16777216
  max_send_message_size = 16777216

[ttrpc]
  address = ""

[debug]
  address = ""
  level = "info"

[metrics]
  address = ""
  grpc_histogram = false

[cgroup]
  path = ""

[plugins]
  [plugins.cri]
    stream_server_address = "127.0.0.1"
    stream_server_port = "0"
    enable_selinux = false
    sandbox_image = "{{pauseImage}}-windows-{{currentversion}}-amd64"
    stats_collect_period = 10
    systemd_cgroup = false
    enable_tls_streaming = false
    max_container_log_line_size = 16384
    [plugins.cri.containerd]
      snapshotter = "windows"
      discard_unpacked_layers = true
      no_pivot = false
      [plugins.cri.containerd.default_runtime]
        runtime_type = "io.containerd.runhcs.v1"
        [plugins.cri.containerd.default_runtime.options]
          Debug = true
          DebugType = 2
          SandboxImage = "{{pauseImage}}-windows-{{currentversion}}-amd64"
          SandboxPlatform = "windows/amd64"
          SandboxIsolation = {{sandboxIsolation}}
      [plugins.cri.containerd.runtimes]
        [plugins.cri.containerd.runtimes.runhcs-wcow-process]
          runtime_type = "io.containerd.runhcs.v1"
          [plugins.cri.containerd.runtimes.runhcs-wcow-process.options]
            Debug = true
            DebugType = 2
            SandboxImage = "{{pauseImage}}-windows-{{currentversion}}-amd64"
            SandboxPlatform = "windows/amd64"
{{hypervisors}}
    [plugins.cri.cni]
      bin_dir = "{{cnibin}}"
      conf_dir = "{{cniconf}}"
    [plugins.cri.registry]
      [plugins.cri.registry.mirrors]
        [plugins.cri.registry.mirrors."docker.io"]
          endpoint = ["https://registry-1.docker.io"]
  [plugins.diff-service]
    default = ["windows"]
  [plugins.scheduler]
    pause_threshold = 0.02
    deletion_threshold = 0
    mutation_threshold = 100
    schedule_delay = "0s"
    startup_delay = "100ms"`)

func k8sContainerdtemplateTomlBytes() ([]byte, error) {
	return _k8sContainerdtemplateToml, nil
}

func k8sContainerdtemplateToml() (*asset, error) {
	bytes, err := k8sContainerdtemplateTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/containerdtemplate.toml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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

var _k8sKubernetesparamsT = []byte(`    "etcdServerCertificate": {
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
      "defaultValue": "20.10.11",
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
         "19.03.12",
         "19.03.13",
         "19.03.14",
         "20.10.5",
         "20.10.7",
         "20.10.8",
         "20.10.9",
         "20.10.10",
         "20.10.11"
       ],
      "type": "string"
    },
    "containerdVersion": {
      "defaultValue": "1.4.11",
      "metadata": {
        "description": "The Azure containerd build version"
      },
      "allowedValues": [
         "1.3.2",
         "1.3.3",
         "1.3.4",
         "1.3.5",
         "1.3.6",
         "1.3.7",
         "1.3.8",
         "1.3.9",
         "1.4.4",
         "1.4.6",
         "1.4.7",
         "1.4.8",
         "1.4.9",
         "1.4.11"
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
        try {
            $args = @{Uri=$Url; Method="Get"; OutFile=$DestinationPath}
            Retry-Command -Command "Invoke-RestMethod" -Args $args -Retries 5 -RetryDelaySeconds 10
        } catch {
            throw "Fail in downloading $Url. Error: $_"
        }
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

    for ($i = 0; ; ) {
        try {
            return & $Command @Args
        }
        catch {
            $i++
            if ($i -ge $Retries) {
                throw $_
            }
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
            # e.g. "mcr.microsoft.com/oss/kubernetes/pause:3.4.1"
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
        Kubeproxy    = @{
            FeatureGates = $global:KubeproxyFeatureGates
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
$global:DefaultContainerdRuntimeHandler = "{{WrapAsParameter "defaultContainerdRuntimeHandler"}}"
$global:HypervRuntimeHandlers = "{{WrapAsParameter "hypervRuntimeHandlers"}}"

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
$global:KubeletNodeLabels = "{{GetAgentKubernetesLabels . "',variables('labelResourceGroup'),'"}}"
$global:KubeletConfigArgs = @( {{GetKubeletConfigKeyValsPsh .KubernetesConfig }} )

$global:KubeproxyFeatureGates = @( {{GetKubeProxyFeatureGatesPsh}} )

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

# Secure Windows TLS protocols
$global:WindowsSecureTLSEnabled = [System.Convert]::ToBoolean("{{WrapAsVariable "windowsSecureTLSEnabled" }}");

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
        Write-Log ".\CustomDataSetupScript.ps1 -MasterIP $MasterIP -KubeDnsServiceIp $KubeDnsServiceIp -MasterFQDNPrefix $MasterFQDNPrefix -Location $Location -AgentKey $AgentKey -AADClientId $AADClientId -AADClientSecret $AADClientSecret -NetworkAPIVersion $NetworkAPIVersion -TargetEnvironment $TargetEnvironment"
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
        icacls.exe "c:\k" /inheritance:r
        icacls.exe "c:\k" /grant:r SYSTEM:` + "`" + `(OI` + "`" + `)` + "`" + `(CI` + "`" + `)` + "`" + `(F` + "`" + `)
        icacls.exe "c:\k" /grant:r BUILTIN\Administrators:` + "`" + `(OI` + "`" + `)` + "`" + `(CI` + "`" + `)` + "`" + `(F` + "`" + `)
        icacls.exe "c:\k" /grant:r BUILTIN\Users:` + "`" + `(OI` + "`" + `)` + "`" + `(CI` + "`" + `)` + "`" + `(RX` + "`" + `)
        Write-Log "c:\k permissions: "
        icacls.exe "c:\k"

        Get-ProvisioningScripts

        Write-KubeClusterConfig -MasterIP $MasterIP -KubeDnsServiceIp $KubeDnsServiceIp

        Write-Log "Download kubelet binaries and unzip"
        Get-KubePackage -KubeBinariesSASURL $global:KubeBinariesPackageSASURL

        # The custom package has a few files that are nessary for future steps (nssm.exe)
        # this is a temporary work around to get the binaries until we depreciate
        # custom package and nssm.exe as defined in #3851.
        if ($global:WindowsKubeBinariesURL){
            Write-Log "Overwriting kube node binaries from $global:WindowsKubeBinariesURL"
            Get-KubeBinaries -KubeBinariesURL $global:WindowsKubeBinariesURL
        }

        if ($useContainerD) {
            Write-Log "Installing ContainerD"
            $containerdTimer = [System.Diagnostics.Stopwatch]::StartNew()
            $cniBinPath = $global:AzureCNIBinDir
            $cniConfigPath = $global:AzureCNIConfDir
            if ($global:NetworkPlugin -eq "kubenet") {
                $cniBinPath = $global:CNIPath
                $cniConfigPath = $global:CNIConfigPath
            }
            Install-Containerd -ContainerdUrl $global:ContainerdUrl -CNIBinDir $cniBinPath -CNIConfDir $cniConfigPath -KubeDir $global:KubeDir
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
            -KubeDir $global:KubeDir ` + "`" + `
            -ContainerRuntime $global:ContainerRuntime

        Get-LogCollectionScripts

        Write-Log "Disable Internet Explorer compat mode and set homepage"
        Set-Explorer

        # if multple LB policies are included for same endpoint then HNS hangs.
        # this fix forces an error  
        Write-Host "Enable a HNS fix in 2021-2C+"
        Set-ItemProperty -Path "HKLM:\SYSTEM\CurrentControlSet\Services\hns\State" -Name HNSControlFlag -Value 1 -Type DWORD

        if ($global:WindowsSecureTLSEnabled) {
            Write-Host "Enable secure TLS protocols"
            . C:\k\windowssecuretls.ps1
            Enable-SecureTls
        }

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

        {{if IsAzureStackCloud}}
            {{if UseCloudControllerManager}}
            # Export the Azure Stack root cert for use in cloud node manager container setup.
            $azsConfigFile = [io.path]::Combine($global:KubeDir, "azurestackcloud.json")
            if (Test-Path -Path $azsConfigFile) {
                $azsJson = Get-Content -Raw -Path $azsConfigFile | ConvertFrom-Json
                if (-not [string]::IsNullOrEmpty($azsJson.managementPortalURL)) {
                    $azsARMUri = [System.Uri]$azsJson.managementPortalURL
                    $azsRootCert = Get-ChildItem -Path Cert:\LocalMachine\Root | Where-Object {$_.DnsNameList -contains $azsARMUri.Host.Substring($azsARMUri.Host.IndexOf(".")).TrimStart(".")}
                    if ($null -ne $azsRootCert) {
                        $azsRootCertFilePath =  [io.path]::Combine($global:KubeDir, "azsroot.cer")
                        Export-Certificate -Cert $azsRootCert -FilePath $azsRootCertFilePath -Type CERT
                    } else {
                        throw "$azsRootCert is null, cannot export Azure Stack root cert"
                    }
                } else {
                    throw "managementPortalURL is null or empty in $azsConfigFile, cannot get Azure Stack ARM uri"
                }
            } else {
                throw "$azsConfigFile does not exist, cannot export Azure Stack root cert"
            }

            # Copy certoc tool for use in cloud node manager container setup. [Environment]::SystemDirectory
            $certocSourcePath = [io.path]::Combine([Environment]::SystemDirectory, "certoc.exe")
            if (Test-Path -Path $certocSourcePath) {
                Copy-Item -Path $certocSourcePath -Destination $global:KubeDir
            }

            # Create add cert script
            $addRootCertFile = [io.path]::Combine($global:KubeDir, "addazsroot.bat")
            if ($null -ne $azsRootCert) {
                [io.file]::WriteAllText($addRootCertFile, "${global:KubeDir}\certoc.exe -addstore root ${azsRootCertFilePath}")
            } else {
                throw "$azsRootCertFilePath is null, cannot create add cert script"
            }
            {{end}}
        {{end}}

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
}`)

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

var _k8sManifestsKubernetesmasterAzureKubernetesKmsYaml = []byte(`apiVersion: v1
kind: Pod
metadata:
  name: azure-kms-provider
  namespace: kube-system
  labels:
    tier: control-plane
    component: azure-kms-provider
spec:
  priorityClassName: system-node-critical
  hostNetwork: true
  containers:
    - name: azure-kms-provider
      image: {{ContainerImage "azure-kms-provider"}}
      imagePullPolicy: IfNotPresent
      command: [{{ContainerConfig "command"}}]
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
      - name: sock
        mountPath: /opt
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
    - name: sock
      hostPath:
        path: /opt
`)

func k8sManifestsKubernetesmasterAzureKubernetesKmsYamlBytes() ([]byte, error) {
	return _k8sManifestsKubernetesmasterAzureKubernetesKmsYaml, nil
}

func k8sManifestsKubernetesmasterAzureKubernetesKmsYaml() (*asset, error) {
	bytes, err := k8sManifestsKubernetesmasterAzureKubernetesKmsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/manifests/kubernetesmaster-azure-kubernetes-kms.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
      {{- if IsAzureStackCloud}}
      env:
      - name: AZURE_ENVIRONMENT_FILEPATH
        value: /etc/kubernetes/azurestackcloud.json
      - name: AZURE_GO_SDK_LOG_LEVEL
        value: INFO
      {{end}}
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
        - name: etc-ssl-certs
          mountPath: /etc/ssl/certs
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
    - name: etc-ssl-certs
      hostPath:
        path: /etc/ssl/certs
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
#EOF
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
  {{- if IsAzureStackCloud}}
      - name: AZURE_GO_SDK_LOG_LEVEL
        value: INFO
  {{end}}
{{end}}
      volumeMounts:
        - name: etc-kubernetes
          mountPath: /etc/kubernetes
{{- if IsKubernetesVersionGe "1.17.0"}}
        - name: etc-ssl
          mountPath: /etc/ssl
          readOnly: true
        - name: ca-certificates
          mountPath: /usr/local/share/ca-certificates
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
    - name: ca-certificates
      hostPath:
        path: /usr/local/share/ca-certificates
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

var _k8sRotateCertsPs1 = []byte(`<#
.DESCRIPTION
    This script rotates a windows node certificates.
    It assumes that client.key, client.crt and ca.crt will be dropped in $env:temp.
#>

. c:\AzureData\k8s\windowskubeletfunc.ps1
. c:\AzureData\k8s\kuberneteswindowsfunctions.ps1

$global:KubeDir = "c:\k"

$global:AgentKeyPath = [io.path]::Combine($env:temp, "client.key")
$global:AgentCertificatePath = [io.path]::Combine($env:temp, "client.crt")
$global:CACertificatePath = [io.path]::Combine($env:temp, "ca.crt")

function Prereqs {
    Assert-FileExists $global:AgentKeyPath
    Assert-FileExists $global:AgentCertificatePath
    Assert-FileExists $global:CACertificatePath
}

function Backup {
    Copy-Item "c:\k\config" "c:\k\config.bak"
    Copy-Item "c:\k\ca.crt" "c:\k\ca.crt.bak"
}

function Update-CACertificate {
    Write-Log "Write ca root"
    Write-CACert -CACertificate $global:CACertificate -KubeDir $global:KubeDir
}

function Update-KubeConfig {
    Write-Log "Write kube config"
    $ClusterConfiguration = ConvertFrom-Json ((Get-Content "c:\k\kubeclusterconfig.json" -ErrorAction Stop) | out-string) 
    $MasterIP = $ClusterConfiguration.Kubernetes.ControlPlane.IpAddress

    $CloudProviderConfig = ConvertFrom-Json ((Get-Content "c:\k\azure.json" -ErrorAction Stop) | out-string) 
    $MasterFQDNPrefix = $CloudProviderConfig.ResourceGroup

    $AgentKey = [System.Convert]::ToBase64String([System.Text.Encoding]::UTF8.GetBytes((Get-Content -Raw $AgentKeyPath)))
    $AgentCertificate = [System.Convert]::ToBase64String([System.Text.Encoding]::UTF8.GetBytes((Get-Content -Raw $AgentCertificatePath)))

    Write-KubeConfig -CACertificate $global:CACertificate ` + "`" + `
        -KubeDir $global:KubeDir ` + "`" + `
        -MasterFQDNPrefix $MasterFQDNPrefix ` + "`" + `
        -MasterIP $MasterIP ` + "`" + `
        -AgentKey $AgentKey ` + "`" + `
        -AgentCertificate $AgentCertificate
}

function Force-Kubelet-CertRotation {
    Remove-Item "/var/lib/kubelet/pki/kubelet-client-current.pem" -Force -ErrorAction Ignore
    Remove-Item "/var/lib/kubelet/pki/kubelet.crt" -Force -ErrorAction Ignore
    Remove-Item "/var/lib/kubelet/pki/kubelet.key" -Force -ErrorAction Ignore

    try {
        $err = Retry-Command -Command "c:\k\windowsnodereset.ps1" -Args @{Foo="Bar"} -Retries 3 -RetryDelaySeconds 10
    } catch {
        Write-Error "Error reseting Windows node. Error: $_"
        throw $_
    }
}

function Start-CertRotation {
    try
    {
        $global:CACertificate = [System.Convert]::ToBase64String([System.Text.Encoding]::UTF8.GetBytes((Get-Content -Raw $CACertificatePath)))

        Prereqs
        Update-CACertificate
        Update-KubeConfig
        Force-Kubelet-CertRotation
    }
    catch
    {
        Write-Error $_
        throw $_
    }
}

function Clean {
    Remove-Item "c:\k\config.bak" -Force -ErrorAction Ignore
    Remove-Item "c:\k\ca.crt.bak" -Force -ErrorAction Ignore
    Remove-Item $global:AgentKeyPath -Force -ErrorAction Ignore
    Remove-Item $global:AgentCertificatePath -Force -ErrorAction Ignore
    Remove-Item $global:CACertificatePath -Force -ErrorAction Ignore
}
`)

func k8sRotateCertsPs1Bytes() ([]byte, error) {
	return _k8sRotateCertsPs1, nil
}

func k8sRotateCertsPs1() (*asset, error) {
	bytes, err := k8sRotateCertsPs1Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/rotate-certs.ps1", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sRotateCertsSh = []byte(`#!/bin/bash -ex

export WD=/etc/kubernetes/rotate-certs
export NEW_CERTS_DIR=${WD}/certs

# copied from cse_helpers.sh, sourcing that file not always works
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

backup() {
  if [ ! -d /etc/kubernetes/certs.bak ]; then
    cp -rp /etc/kubernetes/certs/ /etc/kubernetes/certs.bak
  fi
}

cp_certs() {
  cp -p ${NEW_CERTS_DIR}/etcdpeer* /etc/kubernetes/certs/
  cp -p ${NEW_CERTS_DIR}/etcdclient* /etc/kubernetes/certs/
  cp -p ${NEW_CERTS_DIR}/etcdserver* /etc/kubernetes/certs/
  cp -p ${NEW_CERTS_DIR}/ca.* /etc/kubernetes/certs/
  cp -p ${NEW_CERTS_DIR}/client.* /etc/kubernetes/certs/
  cp -p ${NEW_CERTS_DIR}/apiserver.* /etc/kubernetes/certs/
  cp -p ${NEW_CERTS_DIR}/kubeconfig ~/.kube/config

  rm -f /var/lib/kubelet/pki/kubelet-client-current.pem
}

cp_proxy() {
  source /etc/environment
  local NODE_INDEX
  NODE_INDEX=$(hostname | tail -c 2)
  if [[ $NODE_INDEX == 0 ]]; then
    export OVERRIDE_PROXY_CERTS="true"
  fi
  /etc/kubernetes/generate-proxy-certs.sh
}

agent_certs() {
  cp -p ${NEW_CERTS_DIR}/ca.* /etc/kubernetes/certs/
  cp -p ${NEW_CERTS_DIR}/client.* /etc/kubernetes/certs/

  rm -f /var/lib/kubelet/pki/kubelet-client-current.pem
  sync
  sleep 5
  systemctl_restart 10 5 10 kubelet
}

cleanup() {
  rm -rf ${WD}
  rm -rf /etc/kubernetes/certs.bak
}

"$@"
`)

func k8sRotateCertsShBytes() ([]byte, error) {
	return _k8sRotateCertsSh, nil
}

func k8sRotateCertsSh() (*asset, error) {
	bytes, err := k8sRotateCertsShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s/rotate-certs.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
    
    $aclRule1 = [PSCustomObject]@{
        Type = 'ACL'
        Protocols = '6'
        Action = 'Block'
        Direction = 'Out'
        RemoteAddresses = '168.63.129.16/32'
        RemotePorts = '80'
        Priority = 200
        RuleType = 'Switch'
    }
    $aclRule2 = [PSCustomObject]@{
        Type = 'ACL'
        Action = 'Allow'
        Direction = 'In'
        Priority = 65500
    }
    $aclRule3 = [PSCustomObject]@{
        Type = 'ACL'
        Action = 'Allow'
        Direction = 'Out'
        Priority = 65500
    }
    $jsonContent = [PSCustomObject]@{
        Name = 'EndpointPolicy'
        Value = $aclRule1
    }
    $configJson.plugins[0].AdditionalArgs += $jsonContent
    $jsonContent = [PSCustomObject]@{
        Name = 'EndpointPolicy'
        Value = $aclRule2
    }
    $configJson.plugins[0].AdditionalArgs += $jsonContent
    $jsonContent = [PSCustomObject]@{
        Name = 'EndpointPolicy'
        Value = $aclRule3
    }
    $configJson.plugins[0].AdditionalArgs += $jsonContent

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

    try {
        $response = Retry-Command -Command "Invoke-RestMethod" -Args @{Uri=$uri; Method="Get"; ContentType="application/json"; Headers=$headers} -Retries 5 -RetryDelaySeconds 10
    } catch {
        throw "Error getting subnet prefix. Error: $_"
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
    try {
        $tokenResponse = Retry-Command -Command "Invoke-RestMethod" -Args $args -Retries 5 -RetryDelaySeconds 10
    } catch {
        throw "Error generating token for Azure Resource Manager. Error: $_"
    }

    $token = $tokenResponse | Select-Object -ExpandProperty access_token

    Write-Log "Fetching network interface configuration for node"

    $interfacesUri = "$($azureEnvironment.resourceManagerEndpoint)subscriptions/$SubscriptionId/resourceGroups/$ResourceGroup/providers/Microsoft.Network/networkInterfaces?api-version=$NetworkAPIVersion"
    $headers = @{Authorization="Bearer $token"}
    $args = @{Uri=$interfacesUri; Method="Get"; ContentType="application/json"; Headers=$headers; OutFile=$networkInterfacesFile}
    try {
        Retry-Command -Command "Invoke-RestMethod" -Args $args -Retries 5 -RetryDelaySeconds 10
    } catch {
        throw "Error fetching network interface configuration for node. Error: $_"
    }

    Write-Log "Generating Azure CNI interface file"

    $localNics = Get-NetAdapter | Select-Object -ExpandProperty MacAddress | ForEach-Object {$_ -replace "-",""}

    $sdnNics = Get-Content $networkInterfacesFile ` + "`" + `
        | ConvertFrom-Json ` + "`" + `
        | Select-Object -ExpandProperty value ` + "`" + `
        | Where-Object { $null -ne $_.properties.macAddress -and $localNics.Contains($_.properties.macAddress) } ` + "`" + `
        | Where-Object { $_.properties.ipConfigurations.Count -gt 0}

    if (!$sdnNics) {
        throw 'Error extracting the SDN interfaces from the network interfaces file'
    }

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
    $nas = @(Get-NetAdapter -Physical)

    if ($nas.Count -eq 0) {
        throw "Failed to find any physical network adapters"
    }

    # If there is more than one adapter, use the first adapter that is assigned an ipaddress.
    foreach($na in $nas)
    {
        $netIP = Get-NetIPAddress -ifIndex $na.ifIndex -AddressFamily IPv4 -ErrorAction SilentlyContinue -ErrorVariable netIPErr
        if ($netIP)
        {
            $managementIP = $netIP.IPAddress
            $adapterName = $na.Name
            break
        }
        else {
            Write-Log "No IPv4 found on the network adapter $($na.Name); trying the next adapter ..."
            if ($netIPErr) {
                Write-Log "error when retrieving IPAddress: $netIPErr"
                $netIPErr.Clear()
            }
        }
    }

    if(-Not $managementIP)
    {
        throw "None of the physical network adapters has an IP address"
    }

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
    #
    # The fix which reduces dynamic port usage for non-DSR load balancers is shiped with KB4550969 in April 2020
    # Update the range to [33000, 65535] to avoid that it conflicts with NodePort range (30000 - 32767)
    # Will remove this after WinDSR is enabled by default

    Invoke-Executable -Executable "netsh.exe" -ArgList @("int", "ipv4", "set", "dynamicportrange", "tcp", "33000", "32536")
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
$global:Containerdbinary = (Join-Path $global:ContainerdInstallLocation containerd.exe)

function RegisterContainerDService {
  Param(
    [Parameter(Mandatory = $true)][string]
    $kubedir
  )

  Assert-FileExists $global:Containerdbinary

  # in the past service was not installed via nssm so remove it in case
  $svc = Get-Service -Name "containerd" -ErrorAction SilentlyContinue
  if ($null -ne $svc) {
    sc.exe delete containerd
  }

  Write-Host "Registering containerd as a service"
  # setup containerd
  & "$KubeDir\nssm.exe" install containerd $global:Containerdbinary | RemoveNulls
  & "$KubeDir\nssm.exe" set containerd AppDirectory $KubeDir | RemoveNulls
  & "$KubeDir\nssm.exe" set containerd DisplayName containerd | RemoveNulls
  & "$KubeDir\nssm.exe" set containerd Description containerd | RemoveNulls
  & "$KubeDir\nssm.exe" set containerd Start SERVICE_DEMAND_START | RemoveNulls
  & "$KubeDir\nssm.exe" set containerd ObjectName LocalSystem | RemoveNulls
  & "$KubeDir\nssm.exe" set containerd Type SERVICE_WIN32_OWN_PROCESS | RemoveNulls
  & "$KubeDir\nssm.exe" set containerd AppThrottle 1500 | RemoveNulls
  & "$KubeDir\nssm.exe" set containerd AppStdout "$KubeDir\containerd.log" | RemoveNulls
  & "$KubeDir\nssm.exe" set containerd AppStderr "$KubeDir\containerd.err.log" | RemoveNulls
  & "$KubeDir\nssm.exe" set containerd AppRotateFiles 1 | RemoveNulls
  & "$KubeDir\nssm.exe" set containerd AppRotateOnline 1 | RemoveNulls
  & "$KubeDir\nssm.exe" set containerd AppRotateSeconds 86400 | RemoveNulls
  & "$KubeDir\nssm.exe" set containerd AppRotateBytes 10485760 | RemoveNulls

  $svc = Get-Service -Name "containerd" -ErrorAction SilentlyContinue
  if ($svc.Status -ne "Running") {
    Start-Service containerd
  }
}

function CreateHypervisorRuntime {
  Param(
    [Parameter(Mandatory = $true)][string]
    $image,
    [Parameter(Mandatory = $true)][string]
    $version,
    [Parameter(Mandatory = $true)][string]
    $buildNumber
  )

  return @"
        [plugins.cri.containerd.runtimes.runhcs-wcow-hypervisor-$buildnumber]
          runtime_type = "io.containerd.runhcs.v1"
          [plugins.cri.containerd.runtimes.runhcs-wcow-hypervisor-$buildnumber.options]
            Debug = true
            DebugType = 2
            SandboxImage = "$image-windows-$version-amd64"
            SandboxPlatform = "windows/amd64"
            SandboxIsolation = 1
            ScaleCPULimitsToSandbox = true
"@
}

function CreateHypervisorRuntimes {
  Param(
    [Parameter(Mandatory = $true)][string[]]
    $builds,
    [Parameter(Mandatory = $true)][string]
    $image
  )
  
  Write-Host "Adding hyperv runtimes $builds"
  $hypervRuntimes = ""
  ForEach ($buildNumber in $builds) {
    $windowsVersion = Select-Windows-Version -buildNumber $buildNumber
    $runtime = createHypervisorRuntime -image $pauseImage -version $windowsVersion -buildNumber $buildNumber
    if ($hypervRuntimes -eq "") {
      $hypervRuntimes = $runtime
    }
    else {
      $hypervRuntimes = $hypervRuntimes + "` + "`" + `r` + "`" + `n" + $runtime
    }
  }

  return $hypervRuntimes
}

function Select-Windows-Version {
  param (
    [Parameter()]
    [string]
    $buildNumber
  )

  switch ($buildNumber) {
    "17763" { return "1809" }
    "18362" { return "1903" }
    "18363" { return "1909" }
    "19041" { return "2004" }
    "19042" { return "20H2" }
    Default { return "" } 
  }
}

function Enable-Logging {
  if ((Test-Path "$global:ContainerdInstallLocation\diag.ps1") -And (Test-Path "$global:ContainerdInstallLocation\ContainerPlatform.wprp")) {
    $logs = Join-path $pwd.drive.Root logs
    Write-Log "Containerd hyperv logging enabled; temp location $logs"
    $diag = Join-Path $global:ContainerdInstallLocation diag.ps1
    mkdir -Force $logs
    # !ContainerPlatformPersistent profile is made to work with long term and boot tracing
    & $diag -Start -ProfilePath "$global:ContainerdInstallLocation\ContainerPlatform.wprp!ContainerPlatformPersistent" -TempPath $logs
  }
}

function Install-Containerd {
  Param(
    [Parameter(Mandatory = $true)][string]
    $ContainerdUrl,
    [Parameter(Mandatory = $true)][string]
    $CNIBinDir,
    [Parameter(Mandatory = $true)][string]
    $CNIConfDir,
    [Parameter(Mandatory = $true)][string]
    $KubeDir
  )

  $svc = Get-Service -Name containerd -ErrorAction SilentlyContinue
  if ($null -ne $svc) {
    Write-Log "Stoping containerd service"
    $svc | Stop-Service
  }

  # TODO: check if containerd is already installed and is the same version before this.
  
  # Extract the package
  if ($ContainerdUrl.endswith(".zip")) {
    $zipfile = [Io.path]::Combine($ENV:TEMP, "containerd.zip")
    DownloadFileOverHttp -Url $ContainerdUrl -DestinationPath $zipfile
    Expand-Archive -path $zipfile -DestinationPath $global:ContainerdInstallLocation -Force
    del $zipfile
  }
  elseif ($ContainerdUrl.endswith(".tar.gz")) {
    # upstream containerd package is a tar 
    $tarfile = [Io.path]::Combine($ENV:TEMP, "containerd.tar.gz")
    DownloadFileOverHttp -Url $ContainerdUrl -DestinationPath $tarfile
    mkdir -Force $global:ContainerdInstallLocation
    tar -xzf $tarfile -C $global:ContainerdInstallLocation
    mv -Force $global:ContainerdInstallLocation\bin\* $global:ContainerdInstallLocation\
    del $tarfile
    del -Recurse -Force $global:ContainerdInstallLocation\bin
  }

  # get configuration options
  Add-SystemPathEntry $global:ContainerdInstallLocation
  $configFile = [Io.Path]::Combine($global:ContainerdInstallLocation, "config.toml")
  $clusterConfig = ConvertFrom-Json ((Get-Content $global:KubeClusterConfigPath -ErrorAction Stop) | Out-String)
  $pauseImage = $clusterConfig.Cri.Images.Pause
  $formatedbin = $(($CNIBinDir).Replace("\", "/"))
  $formatedconf = $(($CNIConfDir).Replace("\", "/"))
  $sandboxIsolation = 0
  $windowsReleaseId = (Get-ItemProperty "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion").ReleaseId
  # Starting with 20H2 tags used to publish contianer images may not match the 'ReleaseId'
  switch ($windowsReleaseId)
  {
    "2009" { $windowsVersion = "20H2"}
    default  { $windowsVersion = $windowsReleaseId}
  }
  $hypervRuntimes = ""
  $hypervHandlers = $global:HypervRuntimeHandlers.split(",", [System.StringSplitOptions]::RemoveEmptyEntries)

  # configure
  if ($global:DefaultContainerdRuntimeHandler -eq "hyperv") {
    Write-Log "default runtime for containerd set to hyperv"
    $sandboxIsolation = 1
  }

  $template = Get-Content -Path "c:\AzureData\k8s\containerdtemplate.toml" 
  if ($sandboxIsolation -eq 0 -And $hypervHandlers.Count -eq 0) {
    # remove the value hypervisor place holder
    $template = $template | Select-String -Pattern 'hypervisors' -NotMatch | Out-String
  }
  else {
    $hypervRuntimes = CreateHypervisorRuntimes -builds @($hypervHandlers) -image $pauseImage
  }

  $template.Replace('{{sandboxIsolation}}', $sandboxIsolation).
  Replace('{{pauseImage}}', $pauseImage).
  Replace('{{hypervisors}}', $hypervRuntimes).
  Replace('{{cnibin}}', $formatedbin).
  Replace('{{cniconf}}', $formatedconf).
  Replace('{{currentversion}}', $windowsVersion) | ` + "`" + `
    Out-File -FilePath "$configFile" -Encoding ascii

  RegisterContainerDService -KubeDir $KubeDir
  Enable-Logging
}
`)

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

    $pauseImageVersions = @("1809", "1903", "1909", "2004", "2009", "20h2")

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
        $KubeProxyStartFile,
        [Parameter(Mandatory = $false)][string]
        $ContainerRuntime = "docker"
    )

    $kubeletDependOnServices = $ContainerRuntime
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
        $KubeDir,
        [Parameter(Mandatory = $false)][string]
        $ContainerRuntime = "docker"
    )

    # TODO ksbrmnn fix callers to this function

    $KubeletStartFile = [io.path]::Combine($KubeDir, "kubeletstart.ps1")
    $KubeProxyStartFile = [io.path]::Combine($KubeDir, "kubeproxystart.ps1")

    New-NSSMService -KubeDir $KubeDir ` + "`" + `
        -KubeletStartFile $KubeletStartFile ` + "`" + `
        -KubeProxyStartFile $KubeProxyStartFile ` + "`" + `
        -ContainerRuntime $ContainerRuntime
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

var _windowsparamsT = []byte(`    "kubeBinariesSASURL": {
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
    },
    "defaultContainerdRuntimeHandler": {
      "defaultValue": "process",
      "metadata": {
        "description": "The containerd handler type (process isolated or hyperv)"
      },
      "type": "string"
    },
    "hypervRuntimeHandlers": {
      "defaultValue": "",
      "metadata": {
        "description": "comma separated list of hyperv values"
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
	"agentoutputs.t": agentoutputsT,
	"agentparams.t":  agentparamsT,
	"iaasoutputs.t":  iaasoutputsT,
	"k8s/addons/aad-default-admin-group-rbac.yaml":                       k8sAddonsAadDefaultAdminGroupRbacYaml,
	"k8s/addons/aad-pod-identity.yaml":                                   k8sAddonsAadPodIdentityYaml,
	"k8s/addons/antrea.yaml":                                             k8sAddonsAntreaYaml,
	"k8s/addons/arc-onboarding.yaml":                                     k8sAddonsArcOnboardingYaml,
	"k8s/addons/audit-policy.yaml":                                       k8sAddonsAuditPolicyYaml,
	"k8s/addons/azure-cloud-provider.yaml":                               k8sAddonsAzureCloudProviderYaml,
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
	"k8s/cloud-init/artifacts/apiserver-monitor.service":                 k8sCloudInitArtifactsApiserverMonitorService,
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
	"k8s/cloud-init/artifacts/docker_clear_mount_propagation_flags.conf": k8sCloudInitArtifactsDocker_clear_mount_propagation_flagsConf,
	"k8s/cloud-init/artifacts/enable-dhcpv6.sh":                          k8sCloudInitArtifactsEnableDhcpv6Sh,
	"k8s/cloud-init/artifacts/etc-issue":                                 k8sCloudInitArtifactsEtcIssue,
	"k8s/cloud-init/artifacts/etc-issue.net":                             k8sCloudInitArtifactsEtcIssueNet,
	"k8s/cloud-init/artifacts/etcd-monitor.service":                      k8sCloudInitArtifactsEtcdMonitorService,
	"k8s/cloud-init/artifacts/etcd.service":                              k8sCloudInitArtifactsEtcdService,
	"k8s/cloud-init/artifacts/generateproxycerts.sh":                     k8sCloudInitArtifactsGenerateproxycertsSh,
	"k8s/cloud-init/artifacts/health-monitor.sh":                         k8sCloudInitArtifactsHealthMonitorSh,
	"k8s/cloud-init/artifacts/kms-keyvault-key.service":                  k8sCloudInitArtifactsKmsKeyvaultKeyService,
	"k8s/cloud-init/artifacts/kms-keyvault-key.sh":                       k8sCloudInitArtifactsKmsKeyvaultKeySh,
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
	"k8s/containerdtemplate.toml":                                        k8sContainerdtemplateToml,
	"k8s/kubeconfig.json":                                                k8sKubeconfigJson,
	"k8s/kubernetesparams.t":                                             k8sKubernetesparamsT,
	"k8s/kuberneteswindowsfunctions.ps1":                                 k8sKuberneteswindowsfunctionsPs1,
	"k8s/kuberneteswindowssetup.ps1":                                     k8sKuberneteswindowssetupPs1,
	"k8s/manifests/kubernetesmaster-azure-kubernetes-kms.yaml":           k8sManifestsKubernetesmasterAzureKubernetesKmsYaml,
	"k8s/manifests/kubernetesmaster-cloud-controller-manager.yaml":       k8sManifestsKubernetesmasterCloudControllerManagerYaml,
	"k8s/manifests/kubernetesmaster-kube-addon-manager.yaml":             k8sManifestsKubernetesmasterKubeAddonManagerYaml,
	"k8s/manifests/kubernetesmaster-kube-apiserver.yaml":                 k8sManifestsKubernetesmasterKubeApiserverYaml,
	"k8s/manifests/kubernetesmaster-kube-controller-manager.yaml":        k8sManifestsKubernetesmasterKubeControllerManagerYaml,
	"k8s/manifests/kubernetesmaster-kube-scheduler.yaml":                 k8sManifestsKubernetesmasterKubeSchedulerYaml,
	"k8s/rotate-certs.ps1":                                               k8sRotateCertsPs1,
	"k8s/rotate-certs.sh":                                                k8sRotateCertsSh,
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
	"iaasoutputs.t":  {iaasoutputsT, map[string]*bintree{}},
	"k8s": {nil, map[string]*bintree{
		"addons": {nil, map[string]*bintree{
			"aad-default-admin-group-rbac.yaml":     {k8sAddonsAadDefaultAdminGroupRbacYaml, map[string]*bintree{}},
			"aad-pod-identity.yaml":                 {k8sAddonsAadPodIdentityYaml, map[string]*bintree{}},
			"antrea.yaml":                           {k8sAddonsAntreaYaml, map[string]*bintree{}},
			"arc-onboarding.yaml":                   {k8sAddonsArcOnboardingYaml, map[string]*bintree{}},
			"audit-policy.yaml":                     {k8sAddonsAuditPolicyYaml, map[string]*bintree{}},
			"azure-cloud-provider.yaml":             {k8sAddonsAzureCloudProviderYaml, map[string]*bintree{}},
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
				"apiserver-monitor.service": {k8sCloudInitArtifactsApiserverMonitorService, map[string]*bintree{}},
				"apt-preferences":           {k8sCloudInitArtifactsAptPreferences, map[string]*bintree{}},
				"auditd-rules":              {k8sCloudInitArtifactsAuditdRules, map[string]*bintree{}},
				"cis.sh":                    {k8sCloudInitArtifactsCisSh, map[string]*bintree{}},
				"cse_config.sh":             {k8sCloudInitArtifactsCse_configSh, map[string]*bintree{}},
				"cse_customcloud.sh":        {k8sCloudInitArtifactsCse_customcloudSh, map[string]*bintree{}},
				"cse_helpers.sh":            {k8sCloudInitArtifactsCse_helpersSh, map[string]*bintree{}},
				"cse_install.sh":            {k8sCloudInitArtifactsCse_installSh, map[string]*bintree{}},
				"cse_main.sh":               {k8sCloudInitArtifactsCse_mainSh, map[string]*bintree{}},
				"default-grub":              {k8sCloudInitArtifactsDefaultGrub, map[string]*bintree{}},
				"dhcpv6.service":            {k8sCloudInitArtifactsDhcpv6Service, map[string]*bintree{}},
				"docker-monitor.service":    {k8sCloudInitArtifactsDockerMonitorService, map[string]*bintree{}},
				"docker_clear_mount_propagation_flags.conf": {k8sCloudInitArtifactsDocker_clear_mount_propagation_flagsConf, map[string]*bintree{}},
				"enable-dhcpv6.sh":                          {k8sCloudInitArtifactsEnableDhcpv6Sh, map[string]*bintree{}},
				"etc-issue":                                 {k8sCloudInitArtifactsEtcIssue, map[string]*bintree{}},
				"etc-issue.net":                             {k8sCloudInitArtifactsEtcIssueNet, map[string]*bintree{}},
				"etcd-monitor.service":                      {k8sCloudInitArtifactsEtcdMonitorService, map[string]*bintree{}},
				"etcd.service":                              {k8sCloudInitArtifactsEtcdService, map[string]*bintree{}},
				"generateproxycerts.sh":                     {k8sCloudInitArtifactsGenerateproxycertsSh, map[string]*bintree{}},
				"health-monitor.sh":                         {k8sCloudInitArtifactsHealthMonitorSh, map[string]*bintree{}},
				"kms-keyvault-key.service":                  {k8sCloudInitArtifactsKmsKeyvaultKeyService, map[string]*bintree{}},
				"kms-keyvault-key.sh":                       {k8sCloudInitArtifactsKmsKeyvaultKeySh, map[string]*bintree{}},
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
		"containerdtemplate.toml":        {k8sContainerdtemplateToml, map[string]*bintree{}},
		"kubeconfig.json":                {k8sKubeconfigJson, map[string]*bintree{}},
		"kubernetesparams.t":             {k8sKubernetesparamsT, map[string]*bintree{}},
		"kuberneteswindowsfunctions.ps1": {k8sKuberneteswindowsfunctionsPs1, map[string]*bintree{}},
		"kuberneteswindowssetup.ps1":     {k8sKuberneteswindowssetupPs1, map[string]*bintree{}},
		"manifests": {nil, map[string]*bintree{
			"kubernetesmaster-azure-kubernetes-kms.yaml":     {k8sManifestsKubernetesmasterAzureKubernetesKmsYaml, map[string]*bintree{}},
			"kubernetesmaster-cloud-controller-manager.yaml": {k8sManifestsKubernetesmasterCloudControllerManagerYaml, map[string]*bintree{}},
			"kubernetesmaster-kube-addon-manager.yaml":       {k8sManifestsKubernetesmasterKubeAddonManagerYaml, map[string]*bintree{}},
			"kubernetesmaster-kube-apiserver.yaml":           {k8sManifestsKubernetesmasterKubeApiserverYaml, map[string]*bintree{}},
			"kubernetesmaster-kube-controller-manager.yaml":  {k8sManifestsKubernetesmasterKubeControllerManagerYaml, map[string]*bintree{}},
			"kubernetesmaster-kube-scheduler.yaml":           {k8sManifestsKubernetesmasterKubeSchedulerYaml, map[string]*bintree{}},
		}},
		"rotate-certs.ps1":                {k8sRotateCertsPs1, map[string]*bintree{}},
		"rotate-certs.sh":                 {k8sRotateCertsSh, map[string]*bintree{}},
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
