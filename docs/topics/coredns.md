# CoreDNS customization

The configuration provided by AKS Engine handles most setups by forwarding to
the dns server configured on the nodes.

To customize CoreDNS ([kubernetes docs][Customizing DNS Service]) you can create
a `configmap` that is appended to the builtin configuration. See the example below.

The kubernetes docs on  also has some guidance.

NB:

* The custom configmap must be named `coredns-custom`
* The configmap must contain the item `Corefile`. This is imported by the main
  Corefile.
  * See [CoreDNS configuration docs][] for more info
* The server block cannot be for `.` (the root domain); this is in the main
  `Corefile` and there can not be duplicates.

## Example `coredns-custom` configmap

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: coredns-custom   # Name must be `coredns-custom`
  namespace: kube-system
data:
  Corefile: |
    example.com {
        errors
        cache
        forward . 1.1.1.1 8.8.8.8
    }
```

After applying it you can force coredns to pick up the new config with

```sh
    $ kubectl -n kube-system rollout restart deployment coredns
    deployment.extensions/coredns restarted
```

[Customizing DNS Service]: https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/#configuration-of-stub-domain-and-upstream-nameserver-using-coredns
[CoreDNS configuration docs]: https://coredns.io/manual/toc/#configuration
