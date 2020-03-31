# CSI Proxy for Windows

CSI Proxy for Windows enables support for CSI Plugins on Kubernetes nodes running Windows.

More info can be found in the following places:

- <https://github.com/kubernetes/enhancements/blob/master/keps/sig-windows/20190714-windows-csi-support.md>
- <https://github.com/kubernetes-csi/csi-proxy>

## Requirements

- CSI Proxy for Windows requires Kubernetes version 1.18.0 or greater.

## Usage

### Enable in aks-engine

Add the following fields to `windowsProfile`:

```json
"windowsProfile": {
    ...
    "enableCSIProxy": true,
    "csiProxyURL": "<Path to a package containing Windows csi proxy binaries>"
    ...
}
```

For testing purposes the following csi-proxy binary may be used:

- https://kubernetesartifacts.azureedge.net/csi-proxy/master/binaries/csi-proxy.tar.gz
