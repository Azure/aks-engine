
<a name="v0.66.0"></a>
# [v0.66.0] - 2021-07-29
### Features 🌈
- add support for Kubernetes v1.19.13 ([#4586](https://github.com/Azure/aks-engine/issues/4586))
- add support for Kubernetes v1.20.9 ([#4587](https://github.com/Azure/aks-engine/issues/4587))
- add support for Kubernetes v1.21.3 ([#4588](https://github.com/Azure/aks-engine/issues/4588))
- add support for Kubernetes v1.22.0-beta.2 ([#4584](https://github.com/Azure/aks-engine/issues/4584))
- deprecate Kubernetes 1.17 support ([#4553](https://github.com/Azure/aks-engine/issues/4553))
- Support enabling secure TLS protocols on Windows nodes ([#4568](https://github.com/Azure/aks-engine/issues/4568))
- Enable new clusters w/ 1000 count node pools ([#4569](https://github.com/Azure/aks-engine/issues/4569))
- flexvol and csi addons disabled by default ([#4546](https://github.com/Azure/aks-engine/issues/4546))

### Bug Fixes 🐞
- don't start kubelet if reboot is imminent ([#4555](https://github.com/Azure/aks-engine/issues/4555))

### Maintenance 🔧
- update Ubuntu 18.04-LTS VHD to 2021.07.26 ([#4602](https://github.com/Azure/aks-engine/issues/4602))
- update containerd to 1.4.6 ([#4561](https://github.com/Azure/aks-engine/issues/4561))
- Windows July 2021 VHDs ([#4591](https://github.com/Azure/aks-engine/issues/4591))
- update 18.04-LTS VHD to 2021.07.16 ([#4589](https://github.com/Azure/aks-engine/issues/4589))
- update Go toolchain to 1.16.6 ([#4583](https://github.com/Azure/aks-engine/issues/4583))
- WS 2019 July 2021 patches ([#4579](https://github.com/Azure/aks-engine/issues/4579))
- Add windowssecuretls.ps1 ([#4567](https://github.com/Azure/aks-engine/issues/4567))
- update moby/docker to 20.10.7, and runc to 1.0.0 for Linux (#4562)
- upgrade cluster-autoscaler addon to use v1.21.0 ([#4554](https://github.com/Azure/aks-engine/issues/4554))
- Update Windows docker version to 20.10.5 ([#4556](https://github.com/Azure/aks-engine/issues/4556))
- Update Azure CNI to 1.4.0 ([#4500](https://github.com/Azure/aks-engine/issues/4500))
- use storage.k8s.io/v1 apiVersion for 1.22+ ([#4550](https://github.com/Azure/aks-engine/issues/4550))
- upgrade etcd to 3.3.25 ([#4552](https://github.com/Azure/aks-engine/issues/4552))
- Update Azure constants ([#4537](https://github.com/Azure/aks-engine/issues/4537))
- remove hack/tools go module ([#4534](https://github.com/Azure/aks-engine/issues/4534))

### Continuous Integration 💜
- deprecate prow configurations ([#4547](https://github.com/Azure/aks-engine/issues/4547))
- remove codecov integration ([#4544](https://github.com/Azure/aks-engine/issues/4544))
- standardize go mod vendor ([#4539](https://github.com/Azure/aks-engine/issues/4539))
- use go 1.16 everywhere ([#4528](https://github.com/Azure/aks-engine/issues/4528))
- update hack/tools automatically ([#4532](https://github.com/Azure/aks-engine/issues/4532))
- remove goveralls references ([#4526](https://github.com/Azure/aks-engine/issues/4526))
- automate generate Azure consts ([#4514](https://github.com/Azure/aks-engine/issues/4514))
- don't test unmaintained aad-pod-identity addon ([#4548](https://github.com/Azure/aks-engine/issues/4548))

### Testing 💚
- remove manually curated azure const UT ([#4535](https://github.com/Azure/aks-engine/issues/4535))
- don't include azure-policy addon in E2E cluster configs ([#4559](https://github.com/Azure/aks-engine/issues/4559))

### Documentation 📘
- Document automated release process ([#4530](https://github.com/Azure/aks-engine/issues/4530))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
[Unreleased]: https://github.com/Azure/aks-engine/compare/v0.66.0...HEAD
[v0.66.0]: https://github.com/Azure/aks-engine/compare/v0.65.1...v0.66.0
