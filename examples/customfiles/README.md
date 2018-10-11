# Microsoft Azure Kubernetes Service Engine - Provisioning of master node custom files

## Overview

AKS Engine enables you to provision custom files to your master nodes. This can be used to put whichever files you want on your master nodes to whichever path you want (and have permission to). For example, the use case is when you want additional configurations to native kubernetes features, such as in
the [given example](../examples/customfiles/kubernetes-customfiles-podnodeselector.yaml)

## Examples
[Admission control with pod node selector](../examples/customfiles/kubernetes-customfiles-podnodeselector.yaml) provisions two local files to defined paths on our master nodes. These files define admission control for the apiserver. They could look like:

`admission-control.yaml`
```
kind: AdmissionConfiguration
apiVersion: apiserver.k8s.io/v1alpha1
plugins:
- name: PodNodeSelector
  path: /etc/kubernetes/podnodeselector.yaml
```

`podnodeselector.yaml`
```
podNodeSelectorPluginConfig:
 clusterDefaultNodeSelector: "agentpool=defaultpool"
```

These two need to be provisioned to your master nodes in order for the api server to be able to use them. As seen in the example, the `apiServerConfig` inside the `kubernetesConfig` has been defined as:

```
"apiServerConfig": {
    "--admission-control-config-file":  "/etc/kubernetes/admissioncontrol.yaml"
}
```

This way, the files are provisioned to `/etc/kubernetes` on our master nodes and the apiserver boots up with those provisioned files defining the admission control.