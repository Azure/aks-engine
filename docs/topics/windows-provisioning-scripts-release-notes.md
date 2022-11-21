## v0.0.17

- [Feature: Add hnsremediator for hns crash in Windows nodes #4975](https://github.com/Azure/aks-engine/pull/4975)

## v0.0.16

- [feat: Updating script used to run kubelet to not set docker-specific kubelet args when using containerd #4817](https://github.com/Azure/aks-engine/pull/4817)

## v0.0.15

- [feat: Allow setting Kubeproxy parameters by :ClusterConfiguration.Kubernetes.Kubeproxy.ConfigArgs](https://github.com/Azure/aks-engine/pull/4640)

## v0.0.14

- [fix: chore: Add windowssecuretls.ps1 #4567](https://github.com/Azure/aks-engine/pull/4567)

## v0.0.13

- [fix: Force restart kubeproxy to avoid getting stuck #4375](https://github.com/Azure/aks-engine/pull/4375)

## v0.0.12

- [feat: Use NSSM for containerd and collect containerd logs #4219](https://github.com/Azure/aks-engine/pull/4219)

## v0.0.11

- [feat: allow creation of dualstack Windows clusters #4176](https://github.com/Azure/aks-engine/pull/4176)

## v0.0.10

- [chore: Add windowsnodelabelsync.ps1 #4163](https://github.com/Azure/aks-engine/pull/4163)

## v0.0.9

- [fix: fix network cleanup code on windows for contianerd nodes #4154](https://github.com/Azure/aks-engine/pull/4154)

## v0.0.8

- [feat: Updating kubelet/kube-proxy to run with as high priority processes #4073](https://github.com/Azure/aks-engine/pull/4073)
- [feat: Add WinDSR support #4104](https://github.com/Azure/aks-engine/pull/4104)
- [fix: fix that kubeproxy cannot parse featuregates #4145](https://github.com/Azure/aks-engine/pull/4145)

## v0.0.4

- [fix: Re-ordering HNS policy removal due to 10c change in behavior and consolidating logic in Windows cleanup scripts #4002](https://github.com/Azure/aks-engine/pull/4002)
