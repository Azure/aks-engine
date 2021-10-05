
<a name="v0.67.0"></a>
# [v0.67.0] - 2021-10-05
### Bug Fixes 🐞
- use Invoke-WebRequest instead of curl.exe to download files ([#4646](https://github.com/Azure/aks-engine/issues/4646))
- make Azure Disk & File CSI drivers work on 1.22 ([#4638](https://github.com/Azure/aks-engine/issues/4638))
- upgrade recreates missing control plane VMs ([#4637](https://github.com/Azure/aks-engine/issues/4637))
- Remove 'NT AUTHORITY\AUTHENICATED USERS' write permissions on files in c:\k ([#4633](https://github.com/Azure/aks-engine/issues/4633))

### Continuous Integration 💜
- test large clusters
  - ([#4649](https://github.com/Azure/aks-engine/issues/4649))
  - ([#4652](https://github.com/Azure/aks-engine/issues/4652))

### Features 🌈
- configurable unattended-upgrades ([#4614](https://github.com/Azure/aks-engine/issues/4614))
- add support for Kubernetes v1.21.5 ([#4666](https://github.com/Azure/aks-engine/issues/4666))
- add support for Kubernetes v1.23.0-alpha.2 ([#4670](https://github.com/Azure/aks-engine/issues/4670))
- add support for Kubernetes v1.19.15 ([#4668](https://github.com/Azure/aks-engine/issues/4668))
- add support for Kubernetes v1.22.2 ([#4665](https://github.com/Azure/aks-engine/issues/4665))
- add support for Kubernetes v1.20.11 ([#4667](https://github.com/Azure/aks-engine/issues/4667))
- add support for Kubernetes v1.23.0-alpha.1 ([#4647](https://github.com/Azure/aks-engine/issues/4647))
- add support for Kubernetes v1.19.15 and v1.20.11 on Azure Stack ([#4671](https://github.com/Azure/aks-engine/issues/4671))
- upgrade azure disk csi driver to v1.7.0 ([#4642](https://github.com/Azure/aks-engine/issues/4642))
- Allow setting Kubeproxy parameters by :ClusterConfiguration.Kubernetes.Kubeproxy.ConfigArgs ([#4640](https://github.com/Azure/aks-engine/issues/4640))
- External cloud provider support for Azure Stack Cloud ([#4635](https://github.com/Azure/aks-engine/issues/4635))

### Maintenance 🔧
- update Ubuntu 18.04 VHD to 2021.09.27 ([#4683](https://github.com/Azure/aks-engine/issues/4683))
- New Windows VHD with 9C patches ([#4677](https://github.com/Azure/aks-engine/issues/4677))
- Update default Windows VHDs to Sept 2021 images ([#4673](https://github.com/Azure/aks-engine/issues/4673))
- Update Azure CNI to v.1.4.12 ([#4656](https://github.com/Azure/aks-engine/issues/4656))
- update cluster-autoscaler addon to v1.22.0 ([#4648](https://github.com/Azure/aks-engine/issues/4648))
- Install sept 2021 security patches in Windows VHDs ([#4654](https://github.com/Azure/aks-engine/issues/4654))
- Update Windows provisioning signed package to v0.0.15 ([#4643](https://github.com/Azure/aks-engine/issues/4643))
- Update Azure constants ([#4630](https://github.com/Azure/aks-engine/issues/4630))

### Testing 💚
- stop testing Ubuntu 16.04-LTS ([#4686](https://github.com/Azure/aks-engine/issues/4686))
- large cluster post-scale errors ([#4679](https://github.com/Azure/aks-engine/issues/4679))
- support local kured chart ([#4678](https://github.com/Azure/aks-engine/issues/4678))
- delete failed vmss extensions, wait longer ([#4669](https://github.com/Azure/aks-engine/issues/4669))
- check vmssCSE resource w/ VMSS health check ([#4663](https://github.com/Azure/aks-engine/issues/4663))
- better VMSS recovery in vmss-health-check ([#4662](https://github.com/Azure/aks-engine/issues/4662))
- known-working 5k node test ([#4661](https://github.com/Azure/aks-engine/issues/4661))
- wait 2 hours for new nodes to account for large cluster tests
- detect node count when adding vmss-prototype nodes ([#4657](https://github.com/Azure/aks-engine/issues/4657))
- parallel vmss-prototype jobs when multiple VMSS ([#4655](https://github.com/Azure/aks-engine/issues/4655))
- print vmss-health-check stdout at E2E end ([#4653](https://github.com/Azure/aks-engine/issues/4653))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
[Unreleased]: https://github.com/Azure/aks-engine/compare/v0.67.0...HEAD
[v0.67.0]: https://github.com/Azure/aks-engine/compare/v0.66.1...v0.67.0
