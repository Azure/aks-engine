
<a name="v0.68.0"></a>
# [v0.68.0] - 2022-01-11
### Bug Fixes 🐞
- Azure CNI IP address allocation for win nodes ([#4791](https://github.com/Azure/aks-engine/issues/4791))
- remove api-server insecure flag for 1.20+ ([#4782](https://github.com/Azure/aks-engine/issues/4782))
- Update Windows dynamic port range ([#4767](https://github.com/Azure/aks-engine/issues/4767))

### Continuous Integration 💜
- removing flaky log rotation test for Windows nodes ([#4800](https://github.com/Azure/aks-engine/issues/4800))
- extend sas token expiration for publishing logos ([#4801](https://github.com/Azure/aks-engine/issues/4801))

### Documentation 📘
- clarify CSI support note on air-gapped ASH ([#4774](https://github.com/Azure/aks-engine/issues/4774))

### Features 🌈
- only download nvidia drivers if needed ([#4797](https://github.com/Azure/aks-engine/issues/4797))
- add support for Kubernetes v1.22.5 ([#4789](https://github.com/Azure/aks-engine/issues/4789))
- add support for Kubernetes v1.20.14 ([#4785](https://github.com/Azure/aks-engine/issues/4785))
- add support for Kubernetes v1.23.1 ([#4790](https://github.com/Azure/aks-engine/issues/4790))
- add support for Kubernetes v1.21.8 ([#4786](https://github.com/Azure/aks-engine/issues/4786))

### Maintenance 🔧
- update Linux VHD to 2022.01.10 ([#4804](https://github.com/Azure/aks-engine/issues/4804))
- update Windows VHDs to 17763.2300.220111 ([#4802](https://github.com/Azure/aks-engine/issues/4802))
- update runc to 1.0.2 ([#4798](https://github.com/Azure/aks-engine/issues/4798))
- update nvidia-container-runtime to 3.6.0 ([#4796](https://github.com/Azure/aks-engine/issues/4796))
- Update Azure constants
  - [#4794](https://github.com/Azure/aks-engine/issues/4794)
  - [#4783](https://github.com/Azure/aks-engine/issues/4783)
- update out-of-tree cloudprovider releases ([#4784](https://github.com/Azure/aks-engine/issues/4784))
- use windows VHDs with Nov 2021 security patches ([#4777](https://github.com/Azure/aks-engine/issues/4777))
- check if custom login endpoint is reachable ([#4780](https://github.com/Azure/aks-engine/issues/4780))
- use Nov 2021 as Windows Server 2019 VHD base ([#4772](https://github.com/Azure/aks-engine/issues/4772))
- update CoreDNS to v1.8.6 ([#4771](https://github.com/Azure/aks-engine/issues/4771))
- update cni to v1.4.16 ([#4769](https://github.com/Azure/aks-engine/issues/4769))
- update containerd versions ([#4764](https://github.com/Azure/aks-engine/issues/4764))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
[Unreleased]: https://github.com/Azure/aks-engine/compare/v0.68.0...HEAD
[v0.68.0]: https://github.com/Azure/aks-engine/compare/v0.67.3...v0.68.0
