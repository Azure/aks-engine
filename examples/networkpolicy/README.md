# AKS Engine - Network Policy

There are 3 different Network Policy options :

- Calico
- Cilium
- Antrea

## Calico

Before enabling Calico policy for the AKS Engine cluster, you must first decide which network plugin to use for the cluster: Azure or Kubenet.
The difference between Azure and Kubenet lies in the way IP addresses get allocated; Azure uses Azure-native IPs while Kubenet does static IP assignment based on the node's pod CIDR.
If you're not sure which one to go with, we recommend going with Azure.

The kubernetes-calico deployment template enables Calico networking and policies for the AKS Engine cluster via `"networkPolicy": "calico"` being present inside the `kubernetesConfig`.

```json
  "properties": {
    "orchestratorProfile": {
      "orchestratorType": "Kubernetes",
      "kubernetesConfig": {
        "networkPolicy": "calico",
        "networkPlugin": "azure|kubenet"
      }
```

This template will deploy the [Kubernetes Datastore backed version of Calico](https://docs.projectcalico.org/v3.3/getting-started/kubernetes/installation/other) with user-supplied networking which supports kubernetes ingress policies.

If deploying on a K8s 1.8 or later cluster, then egress policies are also supported!

To understand how to deploy this template, please read the baseline [Kubernetes](../../docs/tutorials/quickstart.md#deploy) document, and use the appropriate **kubernetes-calico-[azure|kubenet].json** example file in this folder as an API model reference.

### Post installation

Once the template has been successfully deployed, following the [simple policy tutorial](https://docs.projectcalico.org/v3.1/getting-started/kubernetes/tutorials/simple-policy) or the [advanced policy tutorial](https://docs.projectcalico.org/v3.1/getting-started/kubernetes/tutorials/advanced-policy) will help to understand calico networking.

> Note: `ping` (ICMP) traffic is blocked on the cluster by default.  Wherever `ping` is used in any tutorial substitute testing access with something like `wget -q --timeout=5 google.com -O -` instead.

## Cilium

The kubernetes-cilium deployment template enables Cilium networking and policies for the AKS Engine cluster via `"networkPolicy": "cilium"` or `"networkPlugin": "cilium"` being present inside the `kubernetesConfig`.

```json
  "properties": {
    "orchestratorProfile": {
      "orchestratorType": "Kubernetes",
      "kubernetesConfig": {
        "networkPolicy": "cilium"
      }
```

> Note:  To execute the `cilium` command that is running inside of the pods, you will need remove the `DenyEscalatingExec` when specifying the Admission Control Values.  If running Kubernetes with the `orchestratorRelease` newer than 1.9 use `--enable-admission-plugins` instead of `--admission-control` as illustrated below:

```json
{
  "apiVersion": "vlabs",
  "properties": {
    "orchestratorProfile": {
      "orchestratorType": "Kubernetes",
      "orchestratorRelease": "1.10",
      "kubernetesConfig": {
        "networkPlugin": "cilium",
        "networkPolicy": "cilium",
        "apiServerConfig": {
           "--enable-admission-plugins": "NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota,AlwaysPullImages"
        },
```

### Post installation

Once the template has been successfully deployed, following the [deploy the demo application](http://cilium.readthedocs.io/en/latest/gettingstarted/minikube/#step-2-deploy-the-demo-application) tutorial will provide a good foundation for how to do L3/4 policy as well as more advanced Layer 7 inspection and routing. If you have [Istio](https://istio.io) you can try this [tutorial](http://cilium.readthedocs.io/en/latest/gettingstarted/istio/) where cilium is used to side by side with Istio to enforce security policies in a Kubernetes deployment.

For the latest documentation on Cilium (including BPF and XDP reference guides), please refer to [this](http://cilium.readthedocs.io/en/latest/)


## Antrea

The kubernetes-antrea deployment template enables Antrea networking and policies for the AKS Engine cluster via `"networkPolicy": "antrea"` and `"networkPlugin": "antrea"` being present inside the `kubernetesConfig`.


```json
  "properties": {
    "orchestratorProfile": {
      "orchestratorType": "Kubernetes",
      "kubernetesConfig": {
        "networkPolicy": "antrea",
        "networkPlugin": "antrea"
      }
```

In Kubernetes 1.18+, Antrea can be enabled on Windows nodes too. Antrea for Windows requires Open Source OVS for networking and enforcing network policies. For the initial implementation in aks-engine, Open Source OVS is not signed. See [this conversation](https://github.com/Azure/aks-engine/pull/3597/files#r458363545) to learn more about the requirements. TestSigning mode is enabled for Windows node to ensure OVS datapath can forward packets correctly.
> Note: TestSigning mode is not recommended for production.

Antrea also supports `NetworkPolicyOnly` mode with Azure CNI. In this mode, Antrea will enforce Network Policies using OVS and Azure CNI will take care of Networking. The kubernetes-antrea deployment template enables Azure Networking and Antrea Network Policies for the AKS Engine via `"networkPolicy": "antrea"` and optional `"networkPlugin": "azure"` being present inside the `kubernetesConfig`. For more details regarding Antrea NetworkPolicyOnly mode, please refer to [this](https://github.com/vmware-tanzu/antrea/blob/master/docs/policy-only.md).


```json
  "apiVersion": "vlabs",
  "properties": {
    "orchestratorProfile": {
      "orchestratorType": "Kubernetes",
      "kubernetesConfig": {
        "networkPolicy": "antrea"
      }
```

### Post installation

For the latest documentation on Antrea, please refer to [this](https://github.com/vmware-tanzu/antrea).
