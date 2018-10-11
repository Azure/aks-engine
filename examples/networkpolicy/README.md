# Microsoft Azure Kubernetes Service Engine - Network Policy

There are 2 different Network Policy options :

- Calico
- Cilium

## Calico

The kubernetes-calico deployment template enables Calico networking and policies for the AKS Engine cluster via `"networkPolicy": "calico"` being present inside the `kubernetesConfig`.

```json
  "properties": {
    "orchestratorProfile": {
      "orchestratorType": "Kubernetes",
      "kubernetesConfig": {
        "networkPolicy": "calico"
      }
```

This template will deploy the [v3.1 release](https://docs.projectcalico.org/v3.1/releases/) of [Kubernetes Datastore Install](https://docs.projectcalico.org/v3.1/getting-started/kubernetes/installation/other) version of calico with the "Calico for policy" with user-supplied networking which supports kubernetes ingress policies.

> Note: The Typha service and deployment is installed on the cluster, but effectively disabled using the default settings of deployment replicas set to 0 and Typha service name not configured.  Typha is recommended to be enabled when scaling to 50+ nodes on the cluster to reduce the load on the Kubernetes API server.  If this functionality is desired to be configurable via the API model, please file an issue on Github requesting this feature be added.  Otherwise, this can be manually changed via modifying and applying changes with the `/etc/kubernetes/addons/calico-daemonset.yaml` file on every master node in the cluster.

If deploying on a K8s 1.8 or later cluster, then egress policies are also supported!

To understand how to deploy this template, please read the baseline [Kubernetes](../../docs/kubernetes.md) document, and use the example **kubernetes-calico.json** file in this folder as an api model reference.

### Post installation

Once the template has been successfully deployed, following the [simple policy tutorial](https://docs.projectcalico.org/v3.1/getting-started/kubernetes/tutorials/simple-policy) or the [advanced policy tutorial](https://docs.projectcalico.org/v3.1/getting-started/kubernetes/tutorials/advanced-policy) will help to understand calico networking.

> Note: `ping` (ICMP) traffic is blocked on the cluster by default.  Wherever `ping` is used in any tutorial substitute testing access with something like `wget -q --timeout=5 google.com -O -` instead.

### Update guidance for clusters deployed by aks-engine releases prior to 0.17.0
Clusters deployed with calico networkPolicy enabled prior to `0.17.0` had calico `2.6.3` deployed, and a daemonset with an `updateStrategy` of `Ondelete`.

aks-engine releases starting with 0.17.0 now produce an addon manifest for calico in `/etc/kubernetes/addons/calico-daemonset.yaml` contaning calico 3.1.x, and an `updateStrategy` of `RollingUpdate`. Due to breaking changes introduced by calico 3, one must first migrate through calico `2.6.5` or a later 2.6.x release in order to migrate to calico 3.1.x. as described in the [calico kubernetes upgrade documentation](https://docs.projectcalico.org/v3.1/getting-started/kubernetes/upgrade/). The aks-engine manifest for calico uses the [kubernetes API datastore, policy-only setup](https://docs.projectcalico.org/v3.1/getting-started/kubernetes/upgrade/upgrade#upgrading-an-installation-that-uses-the-kubernetes-api-datastore).

1. To update to `2.6.5+` in preparation of an upgrade to 3.1.x as specified, edit `/etc/kubernetes/addons/calico-daemonset.yaml` on a master node, replacing `calico/node:v3.1.1` with `calico/node:v2.6.10` and `calico/cni:v3.1.1` with `calico/cni:v2.0.6`. Run `kubectl apply -f /etc/kubernetes/addons/calico-daemonset.yaml`.

    Wait until all the pods in the daemonset get rotated and come up up-to-date, healthy and ready:

    `YYYY-MM-DD HH:MM:SS.FFF [INFO][n] health.go 150: Overall health summary=&health.HealthReport{Live:true, Ready:true}`

2. To complete the upgrade to 3.1.x, edit `/etc/kubernetes/addons/calico-daemonset.yaml` on the master node again, replacing `calico/node:v2.6.10` with `calico/node:v3.1.1` and `calico/cni:v2.0.6` with `calico/cni:v3.1.1`. Run `kubectl apply -f /etc/kubernetes/addons/calico-daemonset.yaml`.

    Propagate this updated manifest to all master nodes in the cluster.

3. Confirm that all the pods in the daemonset get rotated and come up healthy after finishing tasks logged by migrate.go:

    ```
    YYYY-MM-DD HH:MM:SS.FFF [INFO][n] startup.go 1044: Running migration
    YYYY-MM-DD HH:MM:SS.FFF [INFO][n] migrate.go 842: Querying current v1 snapshot and converting to v3
    [xxx]
    YYYY-MM-DD HH:MM:SS.FFF [INFO][n] migrate.go 851: continue by upgrading your calico/node versions to Calico v3.1.x
    YYYY-MM-DD HH:MM:SS.FFF [INFO][n] startup.go 1048: Migration successful
    ```

If you have any customized calico resource manifests, you must also follow the [conversion guide](https://docs.projectcalico.org/v3.0/getting-started/kubernetes/upgrade/convert) for these.

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
