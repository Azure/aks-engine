
<a name="v0.73.0"></a>
# [v0.73.0] - 2022-11-08

# Attention!

This release includes a new version of Nvidia's GPU drivers:

- https://www.nvidia.com/download/driverResults.aspx/191975/

We have observed that this updated version no longer works on some older Azure VM SKUs (e.g., `Standard_NC6`). It validates successfully on newer Azure VM SKUs (e.g., `Standard_NV12s_v3`). If you currently use the AKS Engine-provided `nvidia-device-plugin` addon (enabled by default for N-series VM SKUs), then please test this new version of AKS Engine with the new Nvidia drivers in a staging environment before rolling out any updates to existing clusters using this release. If you confirm that you are definitely not able to use these new drivers with your preferred GPU-enabled VM SKU, you may still use AKS Engine  to create clusters using those GPU-enabled SKUs so long as you disable the `nvidia-device-plugin` addon. For example:

```json
{
  "apiVersion": "vlabs",
  "properties": {
    "orchestratorProfile": {
      "kubernetesConfig": {
        "addons": [
          {
            "name": "nvidia-device-plugin",
            "enabled": false
          }
        ]
      }
    },
...
```

### Bug Fixes 🐞
- enableUnattendedUpgrades not honored ([#4987](https://github.com/Azure/aks-engine/issues/4987))

### Continuous Integration 💜
- disable regular const generation ([#5001](https://github.com/Azure/aks-engine/issues/5001))
- update gpu test SKU ([#4998](https://github.com/Azure/aks-engine/issues/4998))
- enable configurable test cleanup ([#4997](https://github.com/Azure/aks-engine/issues/4997))

### Documentation 📘
- correct upgrade steps in Azure Stack topic page ([#4976](https://github.com/Azure/aks-engine/issues/4976))

### Features 🌈
- add support for Kubernetes v1.23.13 ([#4982](https://github.com/Azure/aks-engine/issues/4982))
- add support for Kubernetes v1.24.7 ([#4983](https://github.com/Azure/aks-engine/issues/4983))

### Maintenance 🔧
- use 2022.11.02 Linux VHD ([#4996](https://github.com/Azure/aks-engine/issues/4996))
- enable Kubernetes v1.23.13 on Azure Stack Hub ([#4994](https://github.com/Azure/aks-engine/issues/4994))
- Adding v0.0.17 signed windows provisioning scripts to VHD builds ([#4978](https://github.com/Azure/aks-engine/issues/4978))
- update Kubernetes libraries to v0.24.7 ([#4990](https://github.com/Azure/aks-engine/issues/4990))
- enable latest k8s v1.22 and v1.23 on Azure Stack Hub ([#4988](https://github.com/Azure/aks-engine/issues/4988))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
[Unreleased]: https://github.com/Azure/aks-engine/compare/v0.73.0...HEAD
[v0.73.0]: https://github.com/Azure/aks-engine/compare/v0.72.0...v0.73.0
