
<a name="v0.60.2"></a>
# [v0.60.2] - 2021-11-02
### Bug Fixes 🐞
- bypass systemd-resolved but use the resolv.conf it generates ([#4726](https://github.com/Azure/aks-engine/issues/4726))
- rotate-certs on custom clouds ([#4248](https://github.com/Azure/aks-engine/issues/4248))

### Continuous Integration 💜
- remove gpu + containerd validation
- gh actions validate supported versions
- apply patch .github/*
- add examples expected by Create Release Branch

### Features 🌈
- rotate-certs fails faster if invalid ssh params ([#4252](https://github.com/Azure/aks-engine/issues/4252))

### Maintenance 🔧
- add T4 GPU as Nvidia GPUs ([#4259](https://github.com/Azure/aks-engine/issues/4259))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
[Unreleased]: https://github.com/Azure/aks-engine/compare/v0.60.2...HEAD
[v0.60.2]: https://github.com/Azure/aks-engine/compare/v0.67.0...v0.60.2
