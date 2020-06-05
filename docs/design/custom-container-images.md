# AKS Engine configurable component images Design Doc

## Ultimate Problem Statement

The existing AKS Engine Kubernetes component container image configuration surface area presents obstacles in the way of:

1. quickly testing/validating specific container images across the set of Kubernetes components in a working cluster; and
2. using Azure Container Compute Upstream-curated MCR container images instead of Kubernetes SIG-Release-curated k8s.gcr.io container images.

## Proximate Problem Statements

1. At present, AKS Engine offers a hodgepodge of user-configurable options for customizing the container images that a particular Kubernetes component (e.g., `kube-proxy`, `kube-scheduler`, `coredns`) requires. Some of the component images are user-configurable, but require a distinct per-component property to do so (e.g., `kube-proxy`); but maintaining per-component properties doesn’t scale well in terms of code maintenance, and will lead to unwieldy flat interface structures; other component images are not user-configurable at all (e.g., `coredns`); still other component images are user-configurable according to a generic "addons" configuration interface (e.g., `metrics-server`). This heterogeneous user configurable surface area makes code maintenance more expensive, and adds friction between the user and a working expression of their desired configuration.
  - This is documented in this public issue:
    - https://github.com/Azure/aks-engine/issues/2378
2. At present, the "blessed" component configuration image URIs are maintained via a concatenation of two properties:
  - A "base URI" property (`KubernetesImageBase` is the property that has the widest impact across the set of component images)
    - e.g., `"k8s.gcr.io/"`
  - A hardcoded string that represents the right-most concatenation substring of the fully qualified image reference URI
    - e.g., `"kube-proxy:v1.16.1"`

In summary, in order to render `"k8s.gcr.io/kube-proxy:v1.16.1"` as the desired container image reference to derive the kube-proxy runtime, we set the KubernetesImageBase property to `"k8s.gcr.io/"`, and rely upon AKS Engine to append `"kube-proxy:v1.16.1"` by way of its hardcoded authority in the codebase for the particular version of Kubernetes in the cluster configuration (1.16.1 in this example).

In practice, this means that the `KubernetesImageBase` property is effectively a "Kubernetes component image registry mirror base URI" property, and in fact this is exactly how that property is leveraged, to redirect container image references to proximate origin URIs when building clusters in non-public cloud environments (e.g., China Cloud, Azure Stack).

To conclude with a concrete problem statement, it is this: the current accommodations that AKS Engine provides for redirecting Kubernetes component container images to another origin assume a k8s.gcr.io container registry mirror. This presents a problem w/ respect to migrating container image configuration to an entirely different container registry URI reference specification, which is what the MCR container image migration effort effectively does.

# A Proposed Solution

In the "Proximate Problem Statements" above, we observe that one of the three existing patterns described in the "hodgepodge of user-configurable options" is a generic, resilient one that will address (1) in the "Ultimate Problem Statement":

- quickly testing/validating specific container images across the set of Kubernetes components in a working cluster

More specifically, the "addons" interface summarized above will allow for the required container image reference configuration across a large set of the Kubernetes components that either aren’t configurable, or which require non-generic, distinct flat properties. That would look like this in the API model:

```json
{
...
                "addons": [
                    {
                        "name": "coredns",
                        "enabled": true,
                        "containers": {
                            "name": "coredns",
                            "image": "mycustomregistry.com/coredns/dev/coredns:v2.5-dfkjlasdklfjsa"
                        }
                    }
                ]
...
}
```

And etc for the remaining kube-addon-manager-enforced addons that happen to fall under this current state (the ones that ultimately render as pods running as a container):

- `kube-dns`
- `kube-proxy`
- `cilium`
- `flannel`

This refactor phase we’ll call **Make All Addons User-Configurable**.

~

The remaining component container images that we need to expose to user-configurable space are the "core" components that we render on disk at /etc/kubernetes/manifests. Those components are those which are essentially bootstrapped by kubelet as pods (via the `"--pod-manifest-path"` kubelet runtime configuration). These components are:

- `kube-scheduler`
- `kube-controller-manager`
- `cloud-controller-manager`
- `kube-apiserver`
- `kube-addon-manager`

There is already some work to expose the container image references for some of those components, e.g.:

- `CustomKubeSchedulerImage`
- `CustomKubeControllerManagerImage`
- `CustomKubeAPIServerImage`

However, the argument is that a flat structure like the above doesn’t scale across the set of configurable properties (we are focused on configurable container images, but this is simply a single example of a useful configurable vector for these core components) across the set of Kubernetes components. An interface like the above addons interface will work much better, and will be generic, accommodating future Kubernetes "core components" that aren’t yet implemented:

```json

{
...
               "components": [
                    {
                        "name": "kube-apiserver",
                        "enabled": true,
                        "containers": {
                            "name": "kube-apiserver",
                            "image": "mycustomregistry.com/kube-apiserver/dev/kube-apiserver:v1.16.1-dfkjlasdklfjsa"
                        }
                    }
                ]
...
}
```

In summary, we will introduce a new "components" configuration interface (a sibling to the existing "addons" configuration interface) that exposes the required configurable container image vector. This refactor phase we’ll call **Make All Core Components User-Configurable**.

~

Now we have addressed the problem of "how to quickly test and validate specific container images across the set of Kubernetes components in a working cluster", which is a critical requirement for the Azure Container Compute Upstream effort to maintain and curate Kubernetes component container images for AKS and AKS Engine. Next we have to address the problem of "how to re-use existing AKS Engine code to introduce a novel mirror specification (MCR) while maintaining backwards compatibility with existing clusters running images from gcr; and without breaking any existing users who are not able to convert to MCR (or don’t want to), and must rely upon the k8s.gcr.io container registry origin, or a mirror that follows its specification".

As stated above, the main point of friction is that the configuration vector currently available to "redirect" the base URI of the origin for sourcing Kubernetes component images assumes, in practice, a "k8s.gcr.io mirror". The MCR container registry origin that is being bootstrapped by the Azure Container Compute Upstream team right now does not match that assumption, and thus we can’t simply re-use the existing configurable space to "migrate to MCR images" (e.g., we cannot simply change the value of `KubernetesImageBase` to `"mcr.microsoft.com/oss/kubernetes/"`, because "mcr.microsoft.com/oss/kubernetes/" is not a mirror of k8s.gcr.io.

What we can do is add a "mirror type" (or "mirror flavor", if you prefer) configuration context to the existing `KubernetesImageBase` property, allowing us to maintain easy backwards-compatibility (by keeping that property valid), and then adapt the underlying hardcoded "image URI substring" values to be sensitive to that context.

Concretely, we could add a new sibling (of KubernetesImageBase) configuration property:
- `KubernetesImageBaseMirrorType`
  - Type: string
  - Supported values: `"gcr"`, `"mcr"`

The value of that property tells the template generation code flows to generate container image reference URI strings according to one of the known specifications supported by AKS Engine:

- k8s.gcr.io
  - e.g., `"k8s.gcr.io/kube-addon-manager-amd64:v9.0.2"`
- mcr.microsoft.com/oss/kubernetes
  - e.g., `"mcr.microsoft.com/oss/kubernetes/kube-addon-manager:v9.0.2"`

The above solution would support a per-environment migration from the current, known-working k8s.gcr.io mirrors (including the origin) to the newly created MCR mirror specification (including unlocking the creation of new MCR mirrors, e.g., in China Cloud, usgov cloud, etc). This refactor phase we’ll call **Enable MCR as an Additive Kubernetes Container Image Registry Mirror**.

# A Proposed Implementation

The above defines three distinct refactor phases:

- **Make All Addons User-Configurable**
- **Make All Core Components User-Configurable**
- **Enable MCR as an Additive Kubernetes Container Image Registry Mirror**

These phases are not serially interdependent, and can be implemented in any order. Arguably, the most critical gap at present is the inability to easily test and validate the Azure Container Compute Upstream-curated images being built from source across the set of known-working cluster configurations. It would make sense to unblock that effort before cutting over to those images generally.

Thus, we will do these refactor phases first of all:

- **Make All Addons User-Configurable**
- **Make All Core Components User-Configurable**

Again, there is no strict functional dependency between those two phases: we can do them in any order. However, it would make the most sense to refactor the addons first of all, as it is arguably the easiest chunk of work, owing to the fact that it re-uses an existing, well tested interface. Doing that work first will allow us to get practical feedback quickly on these refactor proposals regarding exposing per-component configurable container images.

Compare to the user-configurable core components phase, which will require a new interface (or an evolution of the existing addons interface to make it more generic across both "addons" and "components"). This work will require a bit more investment, and if we do it after the addons work, we will have more confidence that we’re on the right track, and can invest even more in this refactor effort.

Thus:

1. **Make All Addons User-Configurable**
2. **Make All Core Components User-Configurable**

Finally, we will tackle the MCR migration problem by re-using the existing KubernetesImageBase configuration vector, and adding a "mirror type" context to the container image reference URIs that AKS Engine outputs during template generation:

3. **3.	Enable MCR as an Additive Kubernetes Container Image Registry Mirror**

The per-component/addon container image configuration work will allow the Azure Container Upstream Compute team to start building per-image testing and validation tooling while we work on the MCR migration phase, improving the likelihood that when the AKS Engine codebase is ready to seamlessly cut over to MCR from gcr.io, we will do so with 100% of the MCR images tested and validated across the set of known-working cluster configurations.

## Alternative Solutions Considered

1. Refactor/re-design the AKS Engine Kubernetes component container image configuration surface area entirely for optimal per-component composition via user configuration.
  - Downsides:
    - Part of the tech debt in the existing implementation is its tight coupling with other fundamental data model composition implementations, making it practically impossible to "triage" the refactor into exclusively this surface area. In other words, the refactor impact is large and unknown, so in order to do that refactor while managing risk of regression, a non-trivial amount of discovery and subsequent redesign across a wide area would be required.
    - The relative urgency with which we want to unblock the MCR effort precludes the time required to do the a full refactor without significant risk of introducing platform regressions.
2. "Rev a new API". We could produce a new data model with a new api version, and new business logic that acts upon it, but that would entail breaking the "mutable vlabs version" pattern we've been adhering to since the beginning of acs-engine.
  - Downsides
    - Such a transition, given the requirements to maintain backwards compatibility with existing the "vlabs" spec (with respect to cluster upgrade and scale especially against pre-existing clusters built using an aks-engine toolchain) will, similar to point 1 above, require non-trivial investigation and purposeful design.
    - Ditto not practical within the time constraints that this refactor work is meant to unblock.

In conclusion: This effort will need to be iterative and practical: we will have to inherit the existing patterns and evolve them in a way that meets the requirements w/ the minimal amount of code refactor side-effects. The above proposed path forward assumes those constraints.
