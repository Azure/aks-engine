
<a name="v0.69.0"></a>
# [v0.69.0] - 2022-03-04
### Bug Fixes 🐞
- ensure correct agent pool selection in scale operations ([#4851](https://github.com/Azure/aks-engine/issues/4851))
- Update calico-3.3.1-cleanup-after-upgrade.yaml ([#4848](https://github.com/Azure/aks-engine/issues/4848))
- rotate-certs recreates all service account tokens ([#4845](https://github.com/Azure/aks-engine/issues/4845))
- update APIVersion for Kubernetes resources ([#4835](https://github.com/Azure/aks-engine/issues/4835))
- Removing kubelet flag --non-masquerade-cidr in v1.24+ clustesr ([#4829](https://github.com/Azure/aks-engine/issues/4829))
- resumeClusterAutoscaler nil dereference in rotate-certs ([#4812](https://github.com/Azure/aks-engine/issues/4812))

### Documentation 📘
- Update quickstart.md ([#4531](https://github.com/Azure/aks-engine/issues/4531))
- add linuxProfile section to Azure Stack Hub doc page ([#4833](https://github.com/Azure/aks-engine/issues/4833))
- add missing links ([#4832](https://github.com/Azure/aks-engine/issues/4832))
- Adding current status updates ([#4828](https://github.com/Azure/aks-engine/issues/4828))
- Update AKS-Engine project status to deprecated ([#4821](https://github.com/Azure/aks-engine/issues/4821))
- Azure Stack doc update for v0.67.3 ([#4803](https://github.com/Azure/aks-engine/issues/4803))

### Features 🌈
- Adding Windows VHDs with Feb 2022 security updates ([#4852](https://github.com/Azure/aks-engine/issues/4852))
- add support for Kubernetes v1.23.4 ([#4839](https://github.com/Azure/aks-engine/issues/4839))
- add support for Kubernetes v1.21.10 ([#4840](https://github.com/Azure/aks-engine/issues/4840))
- add support for Kubernetes v1.22.7 ([#4838](https://github.com/Azure/aks-engine/issues/4838))
- add support for Kubernetes v1.20.15 ([#4837](https://github.com/Azure/aks-engine/issues/4837))
- add support for Kubernetes v1.24.0-alpha.3 ([#4834](https://github.com/Azure/aks-engine/issues/4834))
- add support for Kubernetes v1.24.0-alpha.2 ([#4825](https://github.com/Azure/aks-engine/issues/4825))
- add support for Kubernetes v1.23.3 ([#4820](https://github.com/Azure/aks-engine/issues/4820))
- add support for Kubernetes v1.22.6 ([#4816](https://github.com/Azure/aks-engine/issues/4816))
- add support for Kubernetes v1.23.2 ([#4815](https://github.com/Azure/aks-engine/issues/4815))
- add support for Kubernetes v1.21.9 ([#4814](https://github.com/Azure/aks-engine/issues/4814))
- Add 1.24 support ([#4813](https://github.com/Azure/aks-engine/issues/4813))

### Maintenance 🔧
- force external cloud provider for Kubernetes v1.21+ on Azure Stack Hub ([#4849](https://github.com/Azure/aks-engine/issues/4849))
- update Linux VHD to 2022.02.24 ([#4850](https://github.com/Azure/aks-engine/issues/4850))
- delete old CNCF conformance API models ([#4847](https://github.com/Azure/aks-engine/issues/4847))
- add support for Kubernetes v1.21.10 and v1.22.7 on Azure Stack Hub ([#4846](https://github.com/Azure/aks-engine/issues/4846))
- Using feb 2022 base image for Windows VHDs ([#4844](https://github.com/Azure/aks-engine/issues/4844))
- only label nodes with kubernetes.io/role if they need it ([#4827](https://github.com/Azure/aks-engine/issues/4827))
- making windows provisioning scripts v0.0.16 the default ([#4818](https://github.com/Azure/aks-engine/issues/4818))
- rev cloud-provider-azure vers to 1.23.1, 1.1.4, 1.0.8, 0.7.11 ([#4807](https://github.com/Azure/aks-engine/issues/4807))

### Testing 💚
- update e2e stability tests ([#4842](https://github.com/Azure/aks-engine/issues/4842))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
[Unreleased]: https://github.com/Azure/aks-engine/compare/v0.69.0...HEAD
[v0.69.0]: https://github.com/Azure/aks-engine/compare/v0.68.0...v0.69.0
