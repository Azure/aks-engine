
<a name="v0.75.0"></a>

This release includes a new Ubuntu 20.04-LTS VHD `distro` to use in either your control plane and/or worker node pools. E.g.: 

```json
{
...
        "masterProfile": {
        ...
            "distro": "aks-ubuntu-20.04",
        ...
        },
        "agentPoolProfiles": [
            {
            ...
                "distro": "aks-ubuntu-20.04",
            ...
            },
...
}

```

# [v0.75.0] - 2023-01-13
### Features 🌈
- add support for Kubernetes v1.23.15 ([#5028](https://github.com/Azure/aks-engine/issues/5028))
- add support for Kubernetes v1.24.9 ([#5029](https://github.com/Azure/aks-engine/issues/5029))
- add support for Kubernetes v1.22.17 ([#5027](https://github.com/Azure/aks-engine/issues/5027))
- Add Ubuntu 20.04-LTS VHD ([#5041](https://github.com/Azure/aks-engine/issues/5041))

### Maintenance 🔧
- updating default windows images - Jan 2023 ([#5038](https://github.com/Azure/aks-engine/issues/5038))
- Updating Windows VHD packer job to use Jan 2023 patches ([#5037](https://github.com/Azure/aks-engine/issues/5037))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
[Unreleased]: https://github.com/Azure/aks-engine/compare/v0.75.0...HEAD
[v0.75.0]: https://github.com/Azure/aks-engine/compare/v0.74.0...v0.75.0
