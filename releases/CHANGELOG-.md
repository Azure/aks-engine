
<a name="v0.64.0"></a>
# [v0.64.0] - 2021-06-11
### Bug Fixes 🐞
- Graceful coredns shutdown ([#4443](https://github.com/Azure/aks-engine/issues/4443))
- Run CSI controller driver on master node ([#4432](https://github.com/Azure/aks-engine/issues/4432))
- nvidia + containerd errata ([#4433](https://github.com/Azure/aks-engine/issues/4433))

### Continuous Integration 💜
- don't run vmss-prototype tests for release ([#4458](https://github.com/Azure/aks-engine/issues/4458))
- setup SSH client in release CI ([#4457](https://github.com/Azure/aks-engine/issues/4457))
- don't install ginkgo ([#4456](https://github.com/Azure/aks-engine/issues/4456))
- use go 1.16 to build release artifacts ([#4454](https://github.com/Azure/aks-engine/issues/4454))
- v0.64.0 changelog ([#4453](https://github.com/Azure/aks-engine/issues/4453))
- automated release using github actions ([#4451](https://github.com/Azure/aks-engine/issues/4451))

### Documentation 📘
- remove nvidia-device-plugin addon containerd language ([#4448](https://github.com/Azure/aks-engine/issues/4448))

### Maintenance 🔧
- update 18.04 VHD to 2021.06.08 ([#4449](https://github.com/Azure/aks-engine/issues/4449))
- updates error msg ([#4446](https://github.com/Azure/aks-engine/issues/4446))
- freshen go mod dependencies ([#4440](https://github.com/Azure/aks-engine/issues/4440))
- update metrics-server to v0.5.0 ([#4442](https://github.com/Azure/aks-engine/issues/4442))
- rev 18.04-LTS VHD to 2021.05.24 ([#4438](https://github.com/Azure/aks-engine/issues/4438))
- Add Dockerfile for convenient building ([#4431](https://github.com/Azure/aks-engine/issues/4431))

### Testing 💚
- only validate MTU if Eth0MTU is present ([#4434](https://github.com/Azure/aks-engine/issues/4434))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.63.0"></a>
# [v0.63.0] - 2021-05-27
### Bug Fixes 🐞
- Run CSI controller driver on master node ([#4432](https://github.com/Azure/aks-engine/issues/4432))
- nvidia + containerd errata ([#4433](https://github.com/Azure/aks-engine/issues/4433))
- distinguish ctrd and smalldisk windows image in upgrade scenario ([#4422](https://github.com/Azure/aks-engine/issues/4422))
- Force restart kubeproxy to avoid stucking ([#4375](https://github.com/Azure/aks-engine/issues/4375))

### Continuous Integration 💜
- workaround image publishing errors ([#4413](https://github.com/Azure/aks-engine/issues/4413))
- interpret STABILITY_ITERATIONS as int ([#4390](https://github.com/Azure/aks-engine/issues/4390))
- inherit STABILITY_ITERATIONS from env ([#4387](https://github.com/Azure/aks-engine/issues/4387))
- treat failures in cleanup tasks for pr-windows-signed-scripts.yaml as warmings not errors ([#4384](https://github.com/Azure/aks-engine/issues/4384))
- delete test/e2e/test_cluster_configs/windows/vhd_url.json ([#4383](https://github.com/Azure/aks-engine/issues/4383))
- clear out unnecessary env vars ([#4381](https://github.com/Azure/aks-engine/issues/4381))

### Documentation 📘
- gpu does not require 16.04 ([#4386](https://github.com/Azure/aks-engine/issues/4386))

### Features 🌈
- enable working, validated exec livenessProbes ([#4325](https://github.com/Azure/aks-engine/issues/4325))
- add support for Kubernetes v1.18.19 ([#4410](https://github.com/Azure/aks-engine/issues/4410))
- add support for Kubernetes v1.19.11 ([#4411](https://github.com/Azure/aks-engine/issues/4411))
- add support for Kubernetes v1.20.7 ([#4408](https://github.com/Azure/aks-engine/issues/4408))
- add support for Kubernetes v1.21.1 ([#4407](https://github.com/Azure/aks-engine/issues/4407))
- Add support for custom runc deb URL ([#4403](https://github.com/Azure/aks-engine/issues/4403))
- get-logs from VMs unable to join the cluster ([#4399](https://github.com/Azure/aks-engine/issues/4399))
- add Azure Stack support for k8s 1.18.18, 1.19.10, and 1.20.6 ([#4396](https://github.com/Azure/aks-engine/issues/4396))
- add support for Kubernetes v1.22.0-alpha.1 ([#4395](https://github.com/Azure/aks-engine/issues/4395))
- updates for container monitoring addon omsagent agent March 2021 release ([#4388](https://github.com/Azure/aks-engine/issues/4388))
- add support for Kubernetes v1.18.18 ([#4378](https://github.com/Azure/aks-engine/issues/4378))
- add support for Kubernetes v1.19.10 ([#4377](https://github.com/Azure/aks-engine/issues/4377))
- add support for Kubernetes v1.20.6 ([#4373](https://github.com/Azure/aks-engine/issues/4373))

### Maintenance 🔧
- rev 18.04-LTS VHD to 2021.05.24 ([#4438](https://github.com/Azure/aks-engine/issues/4438))
- Add Dockerfile for convenient building ([#4431](https://github.com/Azure/aks-engine/issues/4431))
- update 18.04 VHD to 2021.05.19 ([#4430](https://github.com/Azure/aks-engine/issues/4430))
- Add May images ([#4429](https://github.com/Azure/aks-engine/issues/4429))
- bump metrics-server to v0.4.4 ([#4424](https://github.com/Azure/aks-engine/issues/4424))
- Ubuntu 16.04-LTS is EOL ([#4418](https://github.com/Azure/aks-engine/issues/4418))
- download nvidia drivers during VHD ([#4421](https://github.com/Azure/aks-engine/issues/4421))
- update Linux VHD to 2021.05.13 ([#4419](https://github.com/Azure/aks-engine/issues/4419))
- bump kube-addon-manager to v9.1.5 ([#4412](https://github.com/Azure/aks-engine/issues/4412))
- Update nvidia runtime ([#4402](https://github.com/Azure/aks-engine/issues/4402))
- May 2021 Windows patch update ([#4404](https://github.com/Azure/aks-engine/issues/4404))
- update CRI builds ([#4394](https://github.com/Azure/aks-engine/issues/4394))
- Update Windows provisioning scripts to 0.0.13 ([#4379](https://github.com/Azure/aks-engine/issues/4379))
- update Go toolchain to v1.16.3 ([#4385](https://github.com/Azure/aks-engine/issues/4385))
- Upgrade nvidia drivers to 450.80.02 ([#4376](https://github.com/Azure/aks-engine/issues/4376))

### Testing 💚
- only validate MTU if Eth0MTU is present ([#4434](https://github.com/Azure/aks-engine/issues/4434))
- use kured 1.7.0 to validate kamino ([#4427](https://github.com/Azure/aks-engine/issues/4427))
- cluster.sh fixes for rotate-certs ([#4382](https://github.com/Azure/aks-engine/issues/4382))
- make e2e singleCommandTimeout configurable ([#4372](https://github.com/Azure/aks-engine/issues/4372))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.62.1"></a>
# [v0.62.1] - 2021-05-05
### Bug Fixes 🐞
- install runc rc92

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.62.0"></a>
# [v0.62.0] - 2021-04-19
### Bug Fixes 🐞
- Fix installer for custom package URL's. ([#4364](https://github.com/Azure/aks-engine/issues/4364))
- Checking return code of curl.exe on Windows ([#4327](https://github.com/Azure/aks-engine/issues/4327))
- specify correct pause image in containerd config for Windows 20h2 ([#4344](https://github.com/Azure/aks-engine/issues/4344))
- disable docker when containerD is specified ([#4282](https://github.com/Azure/aks-engine/issues/4282))
- kube-proxy image on Azure Stack deployments ([#4329](https://github.com/Azure/aks-engine/issues/4329))
- prepend https:// in service-account-issuer flag ([#4262](https://github.com/Azure/aks-engine/issues/4262))
- remove duplicate labels for cilium-operator ([#4260](https://github.com/Azure/aks-engine/issues/4260))
- rotate-certs on custom clouds ([#4248](https://github.com/Azure/aks-engine/issues/4248))
- add retry when downloading files in windows CSE ([#4232](https://github.com/Azure/aks-engine/issues/4232))
- use curl.exe and add retries when downloading artifacts to Windows VHD ([#4206](https://github.com/Azure/aks-engine/issues/4206))
- update nodecustomdata to enable portmapper for kubenet+containerd ([#4191](https://github.com/Azure/aks-engine/issues/4191))
- fix network cleanup code on windows for contianerd nodes ([#4154](https://github.com/Azure/aks-engine/issues/4154))
- update error code for vhd not found ([#4150](https://github.com/Azure/aks-engine/issues/4150))
- fix that kubeproxy cannot parse featuregates ([#4145](https://github.com/Azure/aks-engine/issues/4145))
- Fixing URl for azure-cni v1.20_hotfix package in Windows VHD ([#4136](https://github.com/Azure/aks-engine/issues/4136))
- ensure we wait for the right CRI service name ([#4115](https://github.com/Azure/aks-engine/issues/4115))
- honor user-defined feature gates before applying defaults ([#4113](https://github.com/Azure/aks-engine/issues/4113))
- typo in upgrade log message ([#4109](https://github.com/Azure/aks-engine/issues/4109))
- filter out KeyVault resources during upgrade ([#4072](https://github.com/Azure/aks-engine/issues/4072))
- make system role assignments work w/ upgrade ([#4078](https://github.com/Azure/aks-engine/issues/4078))
- cleanup VMSSName + addpool ([#4067](https://github.com/Azure/aks-engine/issues/4067))
- enforce Azure CNI config via jq instead of sed ([#4060](https://github.com/Azure/aks-engine/issues/4060))
- persist VMSS name in API model ([#4051](https://github.com/Azure/aks-engine/issues/4051))
- Windows and VMSS version validation for azs ([#4030](https://github.com/Azure/aks-engine/issues/4030))
- always ensure etcd on control plane vms ([#4045](https://github.com/Azure/aks-engine/issues/4045))
- fixing an issue where windows cannot start kubelet to get podCIDR on Windows nodes when using kubenet ([#4039](https://github.com/Azure/aks-engine/issues/4039))
- Adding csi-proxy logs to prevent access denied errors ([#4029](https://github.com/Azure/aks-engine/issues/4029))
- upgrade broken for VMAS + managed identities ([#4021](https://github.com/Azure/aks-engine/issues/4021))
- enable dhcpv6 only for ubuntu 16.04 ([#4017](https://github.com/Azure/aks-engine/issues/4017))
- Re-ordering HNS policy removal due to 10c change in behavior and consolidating logic in Windows cleanup scripts  ([#4002](https://github.com/Azure/aks-engine/issues/4002))
- Mount /usr/local/share/ca-certificates folder into kube-controller-manager ([#4001](https://github.com/Azure/aks-engine/issues/4001))
- cancel Pod Eviction jobs when timeout ([#3970](https://github.com/Azure/aks-engine/issues/3970))
- ensure flatcar configs use transparent Azure CNI ([#3972](https://github.com/Azure/aks-engine/issues/3972))
- page not done logic bug ([#3971](https://github.com/Azure/aks-engine/issues/3971))
- resource identifier not passed to custom cloud profile ([#3953](https://github.com/Azure/aks-engine/issues/3953))
- get-logs command collects control plane logs even if apiserver is down ([#3939](https://github.com/Azure/aks-engine/issues/3939))
- use InternalIP for metrics-server DNS resolution ([#3929](https://github.com/Azure/aks-engine/issues/3929))
- VMSS control plane + availability zones ([#3917](https://github.com/Azure/aks-engine/issues/3917))
- Azure Stack upgrade operations clear old MCR ImageBase ([#3922](https://github.com/Azure/aks-engine/issues/3922))
- remove NSG rules for AzureStack ([#3914](https://github.com/Azure/aks-engine/issues/3914))
- fix jumpbox custom data ([#3896](https://github.com/Azure/aks-engine/issues/3896))
- don't warn about CRI update on cluster create ([#3890](https://github.com/Azure/aks-engine/issues/3890))
- ensure addon hostNetwork ports don't conflict ([#3894](https://github.com/Azure/aks-engine/issues/3894))
- de-dup vnet roleAssignment deployments ([#3857](https://github.com/Azure/aks-engine/issues/3857))
- add #EOF sentinel chars to custom search file ([#3862](https://github.com/Azure/aks-engine/issues/3862))
- E2E scale scenario broken if update is false ([#3852](https://github.com/Azure/aks-engine/issues/3852))
- Azure Stack CNI network interfaces file creation fix ([#3792](https://github.com/Azure/aks-engine/issues/3792))

### Code Refactoring 💎
- remove storage type from SKU lookup table ([#4316](https://github.com/Azure/aks-engine/issues/4316))
- remove unsupported orchestrators ([#3965](https://github.com/Azure/aks-engine/issues/3965))
- move ssh-related funcs to its own package ([#3990](https://github.com/Azure/aks-engine/issues/3990))
- delete Azure Stack's KubernetesClient implementation ([#3957](https://github.com/Azure/aks-engine/issues/3957))

### Continuous Integration 💜
- Fix flaky docker rotate e2e test ([#4353](https://github.com/Azure/aks-engine/issues/4353))
- make sure sku is exact match when deciding to create or reuse sku for publishing Windows VHDs ([#4292](https://github.com/Azure/aks-engine/issues/4292))
- skip dns liveness test after upgrade ([#4237](https://github.com/Azure/aks-engine/issues/4237))
- refactor Windows VHD build pipeline updates ([#4105](https://github.com/Azure/aks-engine/issues/4105))
- update dev image to deis/docker-go-dev v1.29.0 ([#4195](https://github.com/Azure/aks-engine/issues/4195))
- remove KMS test from everything cluster config ([#4194](https://github.com/Azure/aks-engine/issues/4194))
- add manual dispatch trigger to nightly build ([#4123](https://github.com/Azure/aks-engine/issues/4123))
- use 3 replicas for no-vnet E2E jobs ([#4061](https://github.com/Azure/aks-engine/issues/4061))
- Updating pr-windows-signed-scripts.yaml to allow for overriding cluster definition ([#4015](https://github.com/Azure/aks-engine/issues/4015))
- remove node-problem-detector addon from PR E2E ([#4010](https://github.com/Azure/aks-engine/issues/4010))
- Adding CI pipeline to verify windows provisioning script content BEFORE it gets check in ([#4008](https://github.com/Azure/aks-engine/issues/4008))
- Windows VHD pipeline tests ([#3977](https://github.com/Azure/aks-engine/issues/3977))
- reduce addons area for containerd config tests ([#3923](https://github.com/Azure/aks-engine/issues/3923))
- adding github actions to create nightly builds ([#3904](https://github.com/Azure/aks-engine/issues/3904))
- E2E fixes to allow more containerd tests ([#3903](https://github.com/Azure/aks-engine/issues/3903))
- add containerd PR E2E tests ([#3892](https://github.com/Azure/aks-engine/issues/3892))
- update OWNERS ([#3853](https://github.com/Azure/aks-engine/issues/3853))

### Documentation 📘
- avoid --control-plane-only upgrade on air-gapped clouds ([#4334](https://github.com/Azure/aks-engine/issues/4334))
- add CSI driver instructions for Azure Stack ([#4326](https://github.com/Azure/aks-engine/issues/4326))
- add v0.60.1 to Azure Stack topic page ([#4324](https://github.com/Azure/aks-engine/issues/4324))
- update get-verions section of Azure Stack documentation ([#4308](https://github.com/Azure/aks-engine/issues/4308))
- Update AAD documentation for Azure Stack and UI changes to Azure ([#4226](https://github.com/Azure/aks-engine/issues/4226))
- refine message about aks-engine project status ([#4312](https://github.com/Azure/aks-engine/issues/4312))
- remove obsolete walkthrough doc ([#4287](https://github.com/Azure/aks-engine/issues/4287))
- add dual-stack iptables api model ([#4185](https://github.com/Azure/aks-engine/issues/4185))
- aad not supported on AZs ([#4182](https://github.com/Azure/aks-engine/issues/4182))
- fix broken link in AAD docs ([#4180](https://github.com/Azure/aks-engine/issues/4180))
- support FAQ ([#4148](https://github.com/Azure/aks-engine/issues/4148))
- remove warning about Windows + custom VNET ([#4147](https://github.com/Azure/aks-engine/issues/4147))
- Clarify README for Azure Stack Hub ([#4128](https://github.com/Azure/aks-engine/issues/4128))
- Added references to capz ([#4124](https://github.com/Azure/aks-engine/issues/4124))
- clarify msi as default, and requirements ([#4122](https://github.com/Azure/aks-engine/issues/4122))
- update dualstack docs to 1.20 ([#4101](https://github.com/Azure/aks-engine/issues/4101))
- fixed broken link in Azure Stack topic page ([#4056](https://github.com/Azure/aks-engine/issues/4056))
- remove double quotes(") at the aks-deploy sample ([#3967](https://github.com/Azure/aks-engine/issues/3967))
- adding instructions for how how to build the Windows VHD for di… ([#3911](https://github.com/Azure/aks-engine/issues/3911))
- confirm flannel + docker is not supported ([#3886](https://github.com/Azure/aks-engine/issues/3886))
- add notes about not upgrading LB config ([#3884](https://github.com/Azure/aks-engine/issues/3884))
- clarify that update is for node pools only ([#3877](https://github.com/Azure/aks-engine/issues/3877))
- CLI operations ([#3837](https://github.com/Azure/aks-engine/issues/3837))
- mark rotate-certs command as experimental ([#3869](https://github.com/Azure/aks-engine/issues/3869))
- add v0.55.4 to Azure Stack topic page ([#3846](https://github.com/Azure/aks-engine/issues/3846))
- remove mention about library for AKS ([#3835](https://github.com/Azure/aks-engine/issues/3835))
- add required attribution reminder to PR template ([#3826](https://github.com/Azure/aks-engine/issues/3826))

### Features 🌈
- deprecate Azure CNI networkmonitor daemonset ([#4363](https://github.com/Azure/aks-engine/issues/4363))
- Windows containerd VHD support ([#4137](https://github.com/Azure/aks-engine/issues/4137))
- don't install nvidia drivers if nvidia-device-plugin is disabled ([#4358](https://github.com/Azure/aks-engine/issues/4358))
- add support for Kubernetes v1.21.0 ([#4362](https://github.com/Azure/aks-engine/issues/4362))
- configurable Linux eth0 MTU ([#4352](https://github.com/Azure/aks-engine/issues/4352))
- add support for Kubernetes v1.21.0-rc.0 ([#4348](https://github.com/Azure/aks-engine/issues/4348))
- add support for Kubernetes v1.20.5 ([#4339](https://github.com/Azure/aks-engine/issues/4339))
- Use NSSM for containerd and collect containerd logs ([#4219](https://github.com/Azure/aks-engine/issues/4219))
- add support for Kubernetes v1.18.17 ([#4340](https://github.com/Azure/aks-engine/issues/4340))
- add support for Kubernetes v1.19.9 ([#4338](https://github.com/Azure/aks-engine/issues/4338))
- update pause image to v3.4.1 ([#4333](https://github.com/Azure/aks-engine/issues/4333))
- add support for Kubernetes 1.21.0-beta.1 ([#4321](https://github.com/Azure/aks-engine/issues/4321))
- updates for container monitoring addon omsagent agent Feb 2021 release ([#4328](https://github.com/Azure/aks-engine/issues/4328))
- Add ACL rules for some network endpoints on Windows ([#3833](https://github.com/Azure/aks-engine/issues/3833))
- allow creation of dualstack Windows clusters ([#4176](https://github.com/Azure/aks-engine/issues/4176))
- add support for Kubernetes 1.21.0-beta.0 ([#4300](https://github.com/Azure/aks-engine/issues/4300))
- add support for Kubernetes v1.18.16 ([#4284](https://github.com/Azure/aks-engine/issues/4284))
- add support for Kubernetes v1.20.4 ([#4285](https://github.com/Azure/aks-engine/issues/4285))
- add support for Kubernetes v1.19.8 ([#4283](https://github.com/Azure/aks-engine/issues/4283))
- add support for Kubernetes 1.21.0-alpha.3 ([#4258](https://github.com/Azure/aks-engine/issues/4258))
- rotate-certs fails faster if invalid ssh params ([#4252](https://github.com/Azure/aks-engine/issues/4252))
- deprecate flannel addon ([#4238](https://github.com/Azure/aks-engine/issues/4238))
- Support configurable `tags` and `enableMultipleStandardLoadBalancers` ([#4048](https://github.com/Azure/aks-engine/issues/4048))
- run unattended upgrades by default ([#4231](https://github.com/Azure/aks-engine/issues/4231))
- run accelerated unattended-upgrade at node creation time ([#4217](https://github.com/Azure/aks-engine/issues/4217))
- reworked rotate-certs command ([#4214](https://github.com/Azure/aks-engine/issues/4214))
- add support for Kubernetes v1.16.15 and v1.18.15 on Azure Stack ([#4187](https://github.com/Azure/aks-engine/issues/4187))
- add support for Kubernetes v1.17.17 on Azure Stack ([#4188](https://github.com/Azure/aks-engine/issues/4188))
- add support for Kubernetes 1.20.2 ([#4192](https://github.com/Azure/aks-engine/issues/4192))
- create kms key as part of cluster bootstrap ([#4170](https://github.com/Azure/aks-engine/issues/4170))
- add support for Kubernetes v1.21.0-alpha.1 ([#4178](https://github.com/Azure/aks-engine/issues/4178))
- add support for Kubernetes v1.18.15 ([#4166](https://github.com/Azure/aks-engine/issues/4166))
- add support for Kubernetes v1.17.17 ([#4167](https://github.com/Azure/aks-engine/issues/4167))
- add support for Kubernetes v1.19.7 ([#4165](https://github.com/Azure/aks-engine/issues/4165))
- setting a default containerd package for Windows ([#4149](https://github.com/Azure/aks-engine/issues/4149))
- add support for Kubernetes v1.18.14 ([#4139](https://github.com/Azure/aks-engine/issues/4139))
- add support for Kubernetes v1.19.6 ([#4140](https://github.com/Azure/aks-engine/issues/4140))
- add support for Kubernetes v1.17.16 ([#4138](https://github.com/Azure/aks-engine/issues/4138))
- add support for Kubernetes v1.20.1 ([#4141](https://github.com/Azure/aks-engine/issues/4141))
- Add WinDSR support ([#4104](https://github.com/Azure/aks-engine/issues/4104))
- Updating kubelet/kube-proxy to run with as high priority processes ([#4073](https://github.com/Azure/aks-engine/issues/4073))
- Add support for specifying linux moby url ([#4120](https://github.com/Azure/aks-engine/issues/4120))
- Updating Windows VHDs to include Dec 2012 patches ([#4118](https://github.com/Azure/aks-engine/issues/4118))
- add support for Kubernetes v1.18.13 ([#4111](https://github.com/Azure/aks-engine/issues/4111))
- add support for Kubernetes v1.19.5 ([#4110](https://github.com/Azure/aks-engine/issues/4110))
- add support for Kubernetes v1.20.0 ([#4102](https://github.com/Azure/aks-engine/issues/4102))
- disable livenessProbe timeout enforcement ([#4085](https://github.com/Azure/aks-engine/issues/4085))
- add support for Kubernetes v1.20.0-rc.0 ([#4076](https://github.com/Azure/aks-engine/issues/4076))
- Enable chrony and host-based time sync by default on Ubuntu 18.04 ([#4011](https://github.com/Azure/aks-engine/issues/4011))
- add support for Kubernetes v1.19.4 ([#4042](https://github.com/Azure/aks-engine/issues/4042))
- add support for Kubernetes 1.18.12 ([#4043](https://github.com/Azure/aks-engine/issues/4043))
- add support for Kubernetes v1.20.0-beta.1 ([#3999](https://github.com/Azure/aks-engine/issues/3999))
- set apiserver tokenRequest flags in 1.20+ ([#3989](https://github.com/Azure/aks-engine/issues/3989))
- add support for Kubernetes v1.20.0-beta.0 ([#3991](https://github.com/Azure/aks-engine/issues/3991))
- add support for Kubernetes v1.20.0-alpha.3 ([#3934](https://github.com/Azure/aks-engine/issues/3934))
- upload collected logs to storage account container ([#3944](https://github.com/Azure/aks-engine/issues/3944))
- add support for Kubernetes v1.17.13 ([#3954](https://github.com/Azure/aks-engine/issues/3954))
- add support for Kubernetes v1.18.10 on Azure Stack ([#3950](https://github.com/Azure/aks-engine/issues/3950))
- add support for Kubernetes v1.18.10 ([#3948](https://github.com/Azure/aks-engine/issues/3948))
- updates for container monitoring addon omsagent agent September 2020 release ([#3942](https://github.com/Azure/aks-engine/issues/3942))
- add support for Kubernetes v1.19.3 ([#3937](https://github.com/Azure/aks-engine/issues/3937))
- custom Windows log collection script ([#3940](https://github.com/Azure/aks-engine/issues/3940))
- enable system-assigned identity by default ([#3856](https://github.com/Azure/aks-engine/issues/3856))
- Target new windows VHD with new K8s binaries (1.19.2, 1.18.9, etc) ([#3905](https://github.com/Azure/aks-engine/issues/3905))
- add ScaleCPULimitsToSandbox for hyperv runtimeclasses ([#3889](https://github.com/Azure/aks-engine/issues/3889))
- allow custom containerd package for Linux nodes ([#3878](https://github.com/Azure/aks-engine/issues/3878))
- Updating Windows VHD build files to support building for multiple OS versions ([#3847](https://github.com/Azure/aks-engine/issues/3847))
- add support for Kubernetes v1.18.9 ([#3841](https://github.com/Azure/aks-engine/issues/3841))
- add support for Kubernetes v1.19.2 ([#3842](https://github.com/Azure/aks-engine/issues/3842))
- add support for Kubernetes v1.17.12 ([#3840](https://github.com/Azure/aks-engine/issues/3840))
- update VMSS node pools ([#3830](https://github.com/Azure/aks-engine/issues/3830))

### Maintenance 🔧
- use rbac.authorization.k8s.io/v1 for metrics-server ([#4367](https://github.com/Azure/aks-engine/issues/4367))
- update Linux VHDs to 2021.04.13 ([#4366](https://github.com/Azure/aks-engine/issues/4366))
- Updating default windows VHDs ([#4368](https://github.com/Azure/aks-engine/issues/4368))
- Adding April 2021 security updates to Windows VHD ([#4365](https://github.com/Azure/aks-engine/issues/4365))
- update Go toolchain to v1.15.11 ([#4355](https://github.com/Azure/aks-engine/issues/4355))
- bump win provisioning scripts to 0.0.12 ([#4345](https://github.com/Azure/aks-engine/issues/4345))
- Update default Windows VHDs to include 3B patches ([#4337](https://github.com/Azure/aks-engine/issues/4337))
- Adding mcr.microsoft.com/oss/kubernetes/pause:3.4.1 to Windows VHD ([#4331](https://github.com/Azure/aks-engine/issues/4331))
- hsn fix for 2c+ ([#4315](https://github.com/Azure/aks-engine/issues/4315))
- March 2021 patch updates for Windows ([#4322](https://github.com/Azure/aks-engine/issues/4322))
- update azurecni to 1.2.7 ([#4319](https://github.com/Azure/aks-engine/issues/4319))
- update windows provisioing scripts to v0.0.11 ([#4314](https://github.com/Azure/aks-engine/issues/4314))
- Updating Windows VHDs for v0.61.0 release ([#4317](https://github.com/Azure/aks-engine/issues/4317))
- update Go toolchain to v1.15.8 ([#4307](https://github.com/Azure/aks-engine/issues/4307))
- Add Feb security patch ([#4298](https://github.com/Azure/aks-engine/issues/4298))
- update logrus to 1.8.0 ([#4310](https://github.com/Azure/aks-engine/issues/4310))
- rev Linux VHDs to 2021.02.22 ([#4296](https://github.com/Azure/aks-engine/issues/4296))
- don't run unattended upgrade if using "no outbound" test feature ([#4297](https://github.com/Azure/aks-engine/issues/4297))
- update node-driver-registrar and liveness-probe images ([#4278](https://github.com/Azure/aks-engine/issues/4278))
- deprecate kube-dashboard addon ([#4268](https://github.com/Azure/aks-engine/issues/4268))
- deprecate aci-connector addon ([#4276](https://github.com/Azure/aks-engine/issues/4276))
- remove addon-resizer from VHD ([#4269](https://github.com/Azure/aks-engine/issues/4269))
- deprecate rescheduler addon ([#4275](https://github.com/Azure/aks-engine/issues/4275))
- keep two versions of Azure CNI in VHD ([#4266](https://github.com/Azure/aks-engine/issues/4266))
- update pause image to 1.4.1 ([#4273](https://github.com/Azure/aks-engine/issues/4273))
- update CNI plugins to v0.9.1 ([#4267](https://github.com/Azure/aks-engine/issues/4267))
- update kube-addon-manager to v9.1.3 ([#4271](https://github.com/Azure/aks-engine/issues/4271))
- support only latest, tested versions of Kubernetes ([#4265](https://github.com/Azure/aks-engine/issues/4265))
- update cluster-autoscaler to 1.20.0 ([#4264](https://github.com/Azure/aks-engine/issues/4264))
- ensure containerd has image cache on VHDs ([#4249](https://github.com/Azure/aks-engine/issues/4249))
- generated code ([#4261](https://github.com/Azure/aks-engine/issues/4261))
- don't include auditd rules in Linux VHDs ([#4253](https://github.com/Azure/aks-engine/issues/4253))
- add T4 GPU as Nvidia GPUs ([#4259](https://github.com/Azure/aks-engine/issues/4259))
- deprecate support for creating new 1.16 clusters ([#4256](https://github.com/Azure/aks-engine/issues/4256))
- don't include gcr-sourced images in Linux VHDs ([#4255](https://github.com/Azure/aks-engine/issues/4255))
- prune non-default images from VHD config ([#4250](https://github.com/Azure/aks-engine/issues/4250))
- fix func arity mismatch in validate_test.go ([#4251](https://github.com/Azure/aks-engine/issues/4251))
- upgrade NPM to v1.2.2_hotfix ([#4225](https://github.com/Azure/aks-engine/issues/4225))
- update windows dockeree version to 19.03.14 ([#4229](https://github.com/Azure/aks-engine/issues/4229))
- add new Azure VM SKUs, brazilsoutheast, westus3 regions ([#4224](https://github.com/Azure/aks-engine/issues/4224))
- block container traffic to 168.63.129.16 ([#4212](https://github.com/Azure/aks-engine/issues/4212))
- rev Linux VHDs to 2021.01.28 ([#4223](https://github.com/Azure/aks-engine/issues/4223))
- update aks-engine VHD Windows VHD ([#4218](https://github.com/Azure/aks-engine/issues/4218))
- update csi-secrets-store to v0.0.19 and akv provider to 0.0.12 ([#4203](https://github.com/Azure/aks-engine/issues/4203))
- update policy addon deployment ([#4201](https://github.com/Azure/aks-engine/issues/4201))
- update adal to v0.9.10 ([#4200](https://github.com/Azure/aks-engine/issues/4200))
- updating windows provisioing scripts to v0.0.10 ([#4184](https://github.com/Azure/aks-engine/issues/4184))
- Update Azure CNI to v1.2.2 ([#4183](https://github.com/Azure/aks-engine/issues/4183))
- Add windowsnodelabelsync.ps1 ([#4163](https://github.com/Azure/aks-engine/issues/4163))
- install Jan 2021 security updates in Windows VHD ([#4168](https://github.com/Azure/aks-engine/issues/4168))
- bump kms keyvault to v0.0.10 ([#4169](https://github.com/Azure/aks-engine/issues/4169))
- Use signed scripts package 0.0.9 on Windows nodes ([#4162](https://github.com/Azure/aks-engine/issues/4162))
- release Windows VHD for v0.59.0 aks-engine release ([#4158](https://github.com/Azure/aks-engine/issues/4158))
- rev Linux VHDs to 2021.01.08 ([#4159](https://github.com/Azure/aks-engine/issues/4159))
- simplify upgrade templates ([#4135](https://github.com/Azure/aks-engine/issues/4135))
- reinforce upgrade warnings/errors ([#4074](https://github.com/Azure/aks-engine/issues/4074))
- Use signed scripts v0.0.8 for Windows deployments ([#4134](https://github.com/Azure/aks-engine/issues/4134))
- simplify scale templates ([#4131](https://github.com/Azure/aks-engine/issues/4131))
- update Azure CNI to v1.2.0_hotfix ([#4129](https://github.com/Azure/aks-engine/issues/4129))
- update vendor deps ([#4125](https://github.com/Azure/aks-engine/issues/4125))
- rev metrics-server to v0.4.1 for all k8s versions ([#4127](https://github.com/Azure/aks-engine/issues/4127))
- Update moby/containerd versions ([#4119](https://github.com/Azure/aks-engine/issues/4119))
- update csi-secrets-store to v0.0.18 and akv provider to 0.0.11 ([#4126](https://github.com/Azure/aks-engine/issues/4126))
- updating azure-npm to 1.2.1 version ([#4094](https://github.com/Azure/aks-engine/issues/4094))
- Installing Dec 2020 cumulative updates for Windows VHDs ([#4103](https://github.com/Azure/aks-engine/issues/4103))
- update adal to v0.9.6 ([#4093](https://github.com/Azure/aks-engine/issues/4093))
- limit number of upgrade retries if new CP nodes bootstrap fails ([#4068](https://github.com/Azure/aks-engine/issues/4068))
- faster rolling updates for daemonset addons ([#4090](https://github.com/Azure/aks-engine/issues/4090))
- update csi-secrets-store addon manifest and images (v0.0.10) ([#4084](https://github.com/Azure/aks-engine/issues/4084))
- update Windows image ([#4086](https://github.com/Azure/aks-engine/issues/4086))
- rev Linux VHDs to 2020.12.02 ([#4081](https://github.com/Azure/aks-engine/issues/4081))
- k8s v1.18 conformance model for Azure Stack ([#4070](https://github.com/Azure/aks-engine/issues/4070))
- validate VHD availability before upgrade/scale on Azure Stack Hub ([#4062](https://github.com/Azure/aks-engine/issues/4062))
- mark Kubernetes 1.17.14 as disabled ([#4044](https://github.com/Azure/aks-engine/issues/4044))
- cleanup dead code ([#4053](https://github.com/Azure/aks-engine/issues/4053))
- only warn about master stuff during create ([#4057](https://github.com/Azure/aks-engine/issues/4057))
- Upgrade CNI to v1.2.0 ([#4058](https://github.com/Azure/aks-engine/issues/4058))
- Updating default csi-proxy version to v0.2.2 ([#4047](https://github.com/Azure/aks-engine/issues/4047))
- format some json in addpool.md ([#4049](https://github.com/Azure/aks-engine/issues/4049))
- remove deprecated AKS code paths ([#4040](https://github.com/Azure/aks-engine/issues/4040))
- deprecate orchestratorType ([#4038](https://github.com/Azure/aks-engine/issues/4038))
- remove deprecated localizations ([#4036](https://github.com/Azure/aks-engine/issues/4036))
- Use Windows November(11b) updates as default vhd ([#4033](https://github.com/Azure/aks-engine/issues/4033))
- Windows November Patches ([#4023](https://github.com/Azure/aks-engine/issues/4023))
- Add 10c image as default ([#4020](https://github.com/Azure/aks-engine/issues/4020))
- rev Linux VHDs to 2020.10.30 ([#4016](https://github.com/Azure/aks-engine/issues/4016))
- Install Windows Server 2019 10C updates in Windows VHD ([#3956](https://github.com/Azure/aks-engine/issues/3956))
- update powershell signed scripts ([#4012](https://github.com/Azure/aks-engine/issues/4012))
- Adding csi-proxy-v0.2.2 to Windows VHD ([#4004](https://github.com/Azure/aks-engine/issues/4004))
- update cluster-autoscalers to latest patches ([#4000](https://github.com/Azure/aks-engine/issues/4000))
- set go sdk log level for Azure Stack clusters ([#3993](https://github.com/Azure/aks-engine/issues/3993))
- update cluster-autoscaler to v1.19.1 ([#3996](https://github.com/Azure/aks-engine/issues/3996))
- update go toolchain to v1.15.3 ([#3879](https://github.com/Azure/aks-engine/issues/3879))
- use transparent mode for Azure CNI ([#3958](https://github.com/Azure/aks-engine/issues/3958))
- gofmt ([#3982](https://github.com/Azure/aks-engine/issues/3982))
- longer default cordonDrainTimeout for Azure Stack Cloud ([#3969](https://github.com/Azure/aks-engine/issues/3969))
- kubelet systemd job depends on CRI service ([#3943](https://github.com/Azure/aks-engine/issues/3943))
- Upgrade CNI to v1.1.8 ([#3907](https://github.com/Azure/aks-engine/issues/3907))
- remove support for Kubernetes 1.15 ([#3751](https://github.com/Azure/aks-engine/issues/3751))
- bump calico to 3.8.9 to get latest patches ([#3924](https://github.com/Azure/aks-engine/issues/3924))
- set the csi-secrets-store to have a priority class ([#3909](https://github.com/Azure/aks-engine/issues/3909))
- rev Linux VHDs to 2020.10.06 ([#3906](https://github.com/Azure/aks-engine/issues/3906))
- enable VHD re-use in no outbound scenarios ([#3897](https://github.com/Azure/aks-engine/issues/3897))
- update csi-secrets-store addon manifest and images (v0.0.9) ([#3891](https://github.com/Azure/aks-engine/issues/3891))
- create azure.json via CSE ([#3876](https://github.com/Azure/aks-engine/issues/3876))
- distribute apiserver.crt to control plane nodes only ([#3860](https://github.com/Azure/aks-engine/issues/3860))
- update azure cni to 1.1.7 ([#3864](https://github.com/Azure/aks-engine/issues/3864))
- update Dashboard addon to v2.0.4 ([#3855](https://github.com/Azure/aks-engine/issues/3855))
- check VHD media name length before being published to Marketplace ([#3799](https://github.com/Azure/aks-engine/issues/3799))
- remove no-op 1.15 version checks in templates ([#3851](https://github.com/Azure/aks-engine/issues/3851))
- updating Windows VHD with new cached artifacts ([#3843](https://github.com/Azure/aks-engine/issues/3843))
- adding 1.19.1 bits to Windows VHD ([#3834](https://github.com/Azure/aks-engine/issues/3834))
- rev Linux VHDs to 2020.09.14 ([#3827](https://github.com/Azure/aks-engine/issues/3827))
- update signed powershell script package to include azure CNI fixes ([#3829](https://github.com/Azure/aks-engine/issues/3829))

### Revert Change ◀️
- "chore: targeting sept updates for Windows 2019 VHD ([#3801](https://github.com/Azure/aks-engine/issues/3801))" ([#3836](https://github.com/Azure/aks-engine/issues/3836))

### Testing 💚
- add containerd no outbound test scenario ([#4371](https://github.com/Azure/aks-engine/issues/4371))
- validate no outbound for 1.20 and 1.21 ([#4370](https://github.com/Azure/aks-engine/issues/4370))
- more resilient node annotations E2E test ([#4361](https://github.com/Azure/aks-engine/issues/4361))
- fix calculation of prior k8s minor version in mocks and upgrade test ([#4359](https://github.com/Azure/aks-engine/issues/4359))
- enable convenient custom kube binary ([#4335](https://github.com/Azure/aks-engine/issues/4335))
- enable parallel E2E tests, update ginkgo to 1.5 ([#4290](https://github.com/Azure/aks-engine/issues/4290))
- use AZURE_CORE_ONLY_SHOW_ERRORS when running E2E ([#4288](https://github.com/Azure/aks-engine/issues/4288))
- expose custom image options as env vars ([#4244](https://github.com/Azure/aks-engine/issues/4244))
- don't test availability sets + 1.21 ([#4239](https://github.com/Azure/aks-engine/issues/4239))
- E2E resilience ([#4235](https://github.com/Azure/aks-engine/issues/4235))
- go routine errata ([#4233](https://github.com/Azure/aks-engine/issues/4233))
- tolerate long SIG image publication times for vmss-prototype test ([#4228](https://github.com/Azure/aks-engine/issues/4228))
- incorporate kured + auto mode into kamino vmss-prototype tests ([#4221](https://github.com/Azure/aks-engine/issues/4221))
- add timestamps to E2E pod logs output ([#4216](https://github.com/Azure/aks-engine/issues/4216))
- enable kamino vmss-prototype dry run tests ([#4215](https://github.com/Azure/aks-engine/issues/4215))
- increase timeout tolerance for vmss-prototype SIG publishing ([#4211](https://github.com/Azure/aks-engine/issues/4211))
- add more stdout to kamino vmss-prototype E2E tests ([#4210](https://github.com/Azure/aks-engine/issues/4210))
- print helm command for vmss-prototype test ([#4207](https://github.com/Azure/aks-engine/issues/4207))
- correct E2E message during vmss-prototype test ([#4205](https://github.com/Azure/aks-engine/issues/4205))
- kamino E2E errata ([#4204](https://github.com/Azure/aks-engine/issues/4204))
- remove dangling fmt.Println ([#4202](https://github.com/Azure/aks-engine/issues/4202))
- enable testing custom kamino vmss prototype images ([#4198](https://github.com/Azure/aks-engine/issues/4198))
- updated kamino vmss-prototype integration E2E ([#4153](https://github.com/Azure/aks-engine/issues/4153))
- ensure hostPort routing ([#4186](https://github.com/Azure/aks-engine/issues/4186))
- enable configurable node prototype tests ([#4100](https://github.com/Azure/aks-engine/issues/4100))
- enable GINKGO_SKIP_AFTER_UPGRADE param in e2e test ([#4089](https://github.com/Azure/aks-engine/issues/4089))
- more resilient azure-arc-onboarding E2E ([#4069](https://github.com/Azure/aks-engine/issues/4069))
- use 3 control plane VMs for base test config ([#4075](https://github.com/Azure/aks-engine/issues/4075))
- dns-liveness livenessProbe tweaks ([#4080](https://github.com/Azure/aks-engine/issues/4080))
- disable useManagedIdentity default value on Azure Stack ([#4063](https://github.com/Azure/aks-engine/issues/4063))
- skip rotate docker log tests if omsagent ([#4059](https://github.com/Azure/aks-engine/issues/4059))
- only test flannel + containerd thru k8s 1.19 ([#4041](https://github.com/Azure/aks-engine/issues/4041))
- re-engage cluster-autoscaler tests if no ssh ([#4034](https://github.com/Azure/aks-engine/issues/4034))
- add dns-loop (actually curl) stress test ([#4024](https://github.com/Azure/aks-engine/issues/4024))
- configurable stability iterations timeout ([#4022](https://github.com/Azure/aks-engine/issues/4022))
- fix two UTs to work with "cli" auth ([#3995](https://github.com/Azure/aks-engine/issues/3995))
- ensure azure-kms-provider E2E coverage ([#3985](https://github.com/Azure/aks-engine/issues/3985))
- "useManagedIdentity": false for VMSS masters ([#3981](https://github.com/Azure/aks-engine/issues/3981))
- Retry label when cluster first comes online ([#3976](https://github.com/Azure/aks-engine/issues/3976))
- Jenkinsfile bool vars ([#3980](https://github.com/Azure/aks-engine/issues/3980))
- Windows Metrics failure when High CPU  ([#3962](https://github.com/Azure/aks-engine/issues/3962))
- validate azure arc in everything config ([#3961](https://github.com/Azure/aks-engine/issues/3961))
- CCM + azurefile + MSI test scenarios ([#3932](https://github.com/Azure/aks-engine/issues/3932))
- skip tests for flatcar cluster config ([#3921](https://github.com/Azure/aks-engine/issues/3921))
- Add azure.json path for custom cloud k8s config & Update stability timeout for Azure CNI network policy ([#3895](https://github.com/Azure/aks-engine/issues/3895))
- E2E fix scale test indexing out of bounds ([#3873](https://github.com/Azure/aks-engine/issues/3873))
- not nil defense broke logic ([#3881](https://github.com/Azure/aks-engine/issues/3881))
- fixes for akse e2e tests ([#3874](https://github.com/Azure/aks-engine/issues/3874))
- E2E optionally validates rotate-certs ([#3866](https://github.com/Azure/aks-engine/issues/3866))
- add more E2E timeout tolerance ([#3850](https://github.com/Azure/aks-engine/issues/3850))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.56.3"></a>
# [v0.56.3] - 2021-04-12
### Bug Fixes 🐞
- Azure Stack CNI network interfaces file creation fix ([#3792](https://github.com/Azure/aks-engine/issues/3792))

### Features 🌈
- don't install nvidia drivers if nvidia-device-plugin is disabled ([#4358](https://github.com/Azure/aks-engine/issues/4358))

### Maintenance 🔧
- add T4 GPU as Nvidia GPUs ([#4259](https://github.com/Azure/aks-engine/issues/4259))
- add new Azure VM SKUs, brazilsoutheast, westus3 regions ([#4224](https://github.com/Azure/aks-engine/issues/4224))
- updating Windows VHD with new cached artifacts ([#3843](https://github.com/Azure/aks-engine/issues/3843))
- adding 1.19.1 bits to Windows VHD ([#3834](https://github.com/Azure/aks-engine/issues/3834))
- rev Linux VHDs to 2020.09.14 ([#3827](https://github.com/Azure/aks-engine/issues/3827))
- update signed powershell script package to include azure CNI fixes ([#3829](https://github.com/Azure/aks-engine/issues/3829))

### Revert Change ◀️
- "chore: targeting sept updates for Windows 2019 VHD ([#3801](https://github.com/Azure/aks-engine/issues/3801))" ([#3836](https://github.com/Azure/aks-engine/issues/3836))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.61.0"></a>
# [v0.61.0] - 2021-03-08
### Bug Fixes 🐞
- remove duplicate labels for cilium-operator ([#4260](https://github.com/Azure/aks-engine/issues/4260))
- rotate-certs on custom clouds ([#4248](https://github.com/Azure/aks-engine/issues/4248))
- add retry when downloading files in windows CSE ([#4232](https://github.com/Azure/aks-engine/issues/4232))
- use curl.exe and add retries when downloading artifacts to Windows VHD ([#4206](https://github.com/Azure/aks-engine/issues/4206))
- update nodecustomdata to enable portmapper for kubenet+containerd ([#4191](https://github.com/Azure/aks-engine/issues/4191))
- fix network cleanup code on windows for contianerd nodes ([#4154](https://github.com/Azure/aks-engine/issues/4154))
- update error code for vhd not found ([#4150](https://github.com/Azure/aks-engine/issues/4150))
- fix that kubeproxy cannot parse featuregates ([#4145](https://github.com/Azure/aks-engine/issues/4145))
- Fixing URl for azure-cni v1.20_hotfix package in Windows VHD ([#4136](https://github.com/Azure/aks-engine/issues/4136))
- ensure we wait for the right CRI service name ([#4115](https://github.com/Azure/aks-engine/issues/4115))
- honor user-defined feature gates before applying defaults ([#4113](https://github.com/Azure/aks-engine/issues/4113))
- typo in upgrade log message ([#4109](https://github.com/Azure/aks-engine/issues/4109))
- filter out KeyVault resources during upgrade ([#4072](https://github.com/Azure/aks-engine/issues/4072))
- make system role assignments work w/ upgrade ([#4078](https://github.com/Azure/aks-engine/issues/4078))
- cleanup VMSSName + addpool ([#4067](https://github.com/Azure/aks-engine/issues/4067))
- enforce Azure CNI config via jq instead of sed ([#4060](https://github.com/Azure/aks-engine/issues/4060))
- persist VMSS name in API model ([#4051](https://github.com/Azure/aks-engine/issues/4051))
- Windows and VMSS version validation for azs ([#4030](https://github.com/Azure/aks-engine/issues/4030))
- always ensure etcd on control plane vms ([#4045](https://github.com/Azure/aks-engine/issues/4045))
- fixing an issue where windows cannot start kubelet to get podCIDR on Windows nodes when using kubenet ([#4039](https://github.com/Azure/aks-engine/issues/4039))
- Adding csi-proxy logs to prevent access denied errors ([#4029](https://github.com/Azure/aks-engine/issues/4029))
- upgrade broken for VMAS + managed identities ([#4021](https://github.com/Azure/aks-engine/issues/4021))
- enable dhcpv6 only for ubuntu 16.04 ([#4017](https://github.com/Azure/aks-engine/issues/4017))
- Re-ordering HNS policy removal due to 10c change in behavior and consolidating logic in Windows cleanup scripts  ([#4002](https://github.com/Azure/aks-engine/issues/4002))
- Mount /usr/local/share/ca-certificates folder into kube-controller-manager ([#4001](https://github.com/Azure/aks-engine/issues/4001))
- cancel Pod Eviction jobs when timeout ([#3970](https://github.com/Azure/aks-engine/issues/3970))
- ensure flatcar configs use transparent Azure CNI ([#3972](https://github.com/Azure/aks-engine/issues/3972))
- page not done logic bug ([#3971](https://github.com/Azure/aks-engine/issues/3971))
- resource identifier not passed to custom cloud profile ([#3953](https://github.com/Azure/aks-engine/issues/3953))
- get-logs command collects control plane logs even if apiserver is down ([#3939](https://github.com/Azure/aks-engine/issues/3939))
- use InternalIP for metrics-server DNS resolution ([#3929](https://github.com/Azure/aks-engine/issues/3929))
- VMSS control plane + availability zones ([#3917](https://github.com/Azure/aks-engine/issues/3917))
- Azure Stack upgrade operations clear old MCR ImageBase ([#3922](https://github.com/Azure/aks-engine/issues/3922))
- remove NSG rules for AzureStack ([#3914](https://github.com/Azure/aks-engine/issues/3914))
- fix jumpbox custom data ([#3896](https://github.com/Azure/aks-engine/issues/3896))
- don't warn about CRI update on cluster create ([#3890](https://github.com/Azure/aks-engine/issues/3890))
- ensure addon hostNetwork ports don't conflict ([#3894](https://github.com/Azure/aks-engine/issues/3894))
- de-dup vnet roleAssignment deployments ([#3857](https://github.com/Azure/aks-engine/issues/3857))
- add #EOF sentinel chars to custom search file ([#3862](https://github.com/Azure/aks-engine/issues/3862))
- E2E scale scenario broken if update is false ([#3852](https://github.com/Azure/aks-engine/issues/3852))
- Azure Stack CNI network interfaces file creation fix ([#3792](https://github.com/Azure/aks-engine/issues/3792))

### Code Refactoring 💎
- remove storage type from SKU lookup table ([#4316](https://github.com/Azure/aks-engine/issues/4316))
- remove unsupported orchestrators ([#3965](https://github.com/Azure/aks-engine/issues/3965))
- move ssh-related funcs to its own package ([#3990](https://github.com/Azure/aks-engine/issues/3990))
- delete Azure Stack's KubernetesClient implementation ([#3957](https://github.com/Azure/aks-engine/issues/3957))

### Continuous Integration 💜
- make sure sku is exact match when deciding to create or reuse sku for publishing Windows VHDs ([#4292](https://github.com/Azure/aks-engine/issues/4292))
- skip dns liveness test after upgrade ([#4237](https://github.com/Azure/aks-engine/issues/4237))
- refactor Windows VHD build pipeline updates ([#4105](https://github.com/Azure/aks-engine/issues/4105))
- update dev image to deis/docker-go-dev v1.29.0 ([#4195](https://github.com/Azure/aks-engine/issues/4195))
- remove KMS test from everything cluster config ([#4194](https://github.com/Azure/aks-engine/issues/4194))
- add manual dispatch trigger to nightly build ([#4123](https://github.com/Azure/aks-engine/issues/4123))
- use 3 replicas for no-vnet E2E jobs ([#4061](https://github.com/Azure/aks-engine/issues/4061))
- Updating pr-windows-signed-scripts.yaml to allow for overriding cluster definition ([#4015](https://github.com/Azure/aks-engine/issues/4015))
- remove node-problem-detector addon from PR E2E ([#4010](https://github.com/Azure/aks-engine/issues/4010))
- Adding CI pipeline to verify windows provisioning script content BEFORE it gets check in ([#4008](https://github.com/Azure/aks-engine/issues/4008))
- Windows VHD pipeline tests ([#3977](https://github.com/Azure/aks-engine/issues/3977))
- reduce addons area for containerd config tests ([#3923](https://github.com/Azure/aks-engine/issues/3923))
- adding github actions to create nightly builds ([#3904](https://github.com/Azure/aks-engine/issues/3904))
- E2E fixes to allow more containerd tests ([#3903](https://github.com/Azure/aks-engine/issues/3903))
- add containerd PR E2E tests ([#3892](https://github.com/Azure/aks-engine/issues/3892))
- update OWNERS ([#3853](https://github.com/Azure/aks-engine/issues/3853))

### Documentation 📘
- Update AAD documentation for Azure Stack and UI changes to Azure ([#4226](https://github.com/Azure/aks-engine/issues/4226))
- refine message about aks-engine project status ([#4312](https://github.com/Azure/aks-engine/issues/4312))
- remove obsolete walkthrough doc ([#4287](https://github.com/Azure/aks-engine/issues/4287))
- add dual-stack iptables api model ([#4185](https://github.com/Azure/aks-engine/issues/4185))
- aad not supported on AZs ([#4182](https://github.com/Azure/aks-engine/issues/4182))
- fix broken link in AAD docs ([#4180](https://github.com/Azure/aks-engine/issues/4180))
- support FAQ ([#4148](https://github.com/Azure/aks-engine/issues/4148))
- remove warning about Windows + custom VNET ([#4147](https://github.com/Azure/aks-engine/issues/4147))
- Clarify README for Azure Stack Hub ([#4128](https://github.com/Azure/aks-engine/issues/4128))
- Added references to capz ([#4124](https://github.com/Azure/aks-engine/issues/4124))
- clarify msi as default, and requirements ([#4122](https://github.com/Azure/aks-engine/issues/4122))
- update dualstack docs to 1.20 ([#4101](https://github.com/Azure/aks-engine/issues/4101))
- fixed broken link in Azure Stack topic page ([#4056](https://github.com/Azure/aks-engine/issues/4056))
- remove double quotes(") at the aks-deploy sample ([#3967](https://github.com/Azure/aks-engine/issues/3967))
- adding instructions for how how to build the Windows VHD for di… ([#3911](https://github.com/Azure/aks-engine/issues/3911))
- confirm flannel + docker is not supported ([#3886](https://github.com/Azure/aks-engine/issues/3886))
- add notes about not upgrading LB config ([#3884](https://github.com/Azure/aks-engine/issues/3884))
- clarify that update is for node pools only ([#3877](https://github.com/Azure/aks-engine/issues/3877))
- CLI operations ([#3837](https://github.com/Azure/aks-engine/issues/3837))
- mark rotate-certs command as experimental ([#3869](https://github.com/Azure/aks-engine/issues/3869))
- add v0.55.4 to Azure Stack topic page ([#3846](https://github.com/Azure/aks-engine/issues/3846))
- remove mention about library for AKS ([#3835](https://github.com/Azure/aks-engine/issues/3835))
- add required attribution reminder to PR template ([#3826](https://github.com/Azure/aks-engine/issues/3826))

### Features 🌈
- allow creation of dualstack Windows clusters ([#4176](https://github.com/Azure/aks-engine/issues/4176))
- add support for Kubernetes 1.21.0-beta.0 ([#4300](https://github.com/Azure/aks-engine/issues/4300))
- add support for Kubernetes v1.18.16 ([#4284](https://github.com/Azure/aks-engine/issues/4284))
- add support for Kubernetes v1.20.4 ([#4285](https://github.com/Azure/aks-engine/issues/4285))
- add support for Kubernetes v1.19.8 ([#4283](https://github.com/Azure/aks-engine/issues/4283))
- add support for Kubernetes 1.21.0-alpha.3 ([#4258](https://github.com/Azure/aks-engine/issues/4258))
- rotate-certs fails faster if invalid ssh params ([#4252](https://github.com/Azure/aks-engine/issues/4252))
- deprecate flannel addon ([#4238](https://github.com/Azure/aks-engine/issues/4238))
- Support configurable `tags` and `enableMultipleStandardLoadBalancers` ([#4048](https://github.com/Azure/aks-engine/issues/4048))
- run unattended upgrades by default ([#4231](https://github.com/Azure/aks-engine/issues/4231))
- run accelerated unattended-upgrade at node creation time ([#4217](https://github.com/Azure/aks-engine/issues/4217))
- reworked rotate-certs command ([#4214](https://github.com/Azure/aks-engine/issues/4214))
- add support for Kubernetes v1.16.15 and v1.18.15 on Azure Stack ([#4187](https://github.com/Azure/aks-engine/issues/4187))
- add support for Kubernetes v1.17.17 on Azure Stack ([#4188](https://github.com/Azure/aks-engine/issues/4188))
- add support for Kubernetes 1.20.2 ([#4192](https://github.com/Azure/aks-engine/issues/4192))
- create kms key as part of cluster bootstrap ([#4170](https://github.com/Azure/aks-engine/issues/4170))
- add support for Kubernetes v1.21.0-alpha.1 ([#4178](https://github.com/Azure/aks-engine/issues/4178))
- add support for Kubernetes v1.18.15 ([#4166](https://github.com/Azure/aks-engine/issues/4166))
- add support for Kubernetes v1.17.17 ([#4167](https://github.com/Azure/aks-engine/issues/4167))
- add support for Kubernetes v1.19.7 ([#4165](https://github.com/Azure/aks-engine/issues/4165))
- setting a default containerd package for Windows ([#4149](https://github.com/Azure/aks-engine/issues/4149))
- add support for Kubernetes v1.18.14 ([#4139](https://github.com/Azure/aks-engine/issues/4139))
- add support for Kubernetes v1.19.6 ([#4140](https://github.com/Azure/aks-engine/issues/4140))
- add support for Kubernetes v1.17.16 ([#4138](https://github.com/Azure/aks-engine/issues/4138))
- add support for Kubernetes v1.20.1 ([#4141](https://github.com/Azure/aks-engine/issues/4141))
- Add WinDSR support ([#4104](https://github.com/Azure/aks-engine/issues/4104))
- Updating kubelet/kube-proxy to run with as high priority processes ([#4073](https://github.com/Azure/aks-engine/issues/4073))
- Add support for specifying linux moby url ([#4120](https://github.com/Azure/aks-engine/issues/4120))
- Updating Windows VHDs to include Dec 2012 patches ([#4118](https://github.com/Azure/aks-engine/issues/4118))
- add support for Kubernetes v1.18.13 ([#4111](https://github.com/Azure/aks-engine/issues/4111))
- add support for Kubernetes v1.19.5 ([#4110](https://github.com/Azure/aks-engine/issues/4110))
- add support for Kubernetes v1.20.0 ([#4102](https://github.com/Azure/aks-engine/issues/4102))
- disable livenessProbe timeout enforcement ([#4085](https://github.com/Azure/aks-engine/issues/4085))
- add support for Kubernetes v1.20.0-rc.0 ([#4076](https://github.com/Azure/aks-engine/issues/4076))
- Enable chrony and host-based time sync by default on Ubuntu 18.04 ([#4011](https://github.com/Azure/aks-engine/issues/4011))
- add support for Kubernetes v1.19.4 ([#4042](https://github.com/Azure/aks-engine/issues/4042))
- add support for Kubernetes 1.18.12 ([#4043](https://github.com/Azure/aks-engine/issues/4043))
- add support for Kubernetes v1.20.0-beta.1 ([#3999](https://github.com/Azure/aks-engine/issues/3999))
- set apiserver tokenRequest flags in 1.20+ ([#3989](https://github.com/Azure/aks-engine/issues/3989))
- add support for Kubernetes v1.20.0-beta.0 ([#3991](https://github.com/Azure/aks-engine/issues/3991))
- add support for Kubernetes v1.20.0-alpha.3 ([#3934](https://github.com/Azure/aks-engine/issues/3934))
- upload collected logs to storage account container ([#3944](https://github.com/Azure/aks-engine/issues/3944))
- add support for Kubernetes v1.17.13 ([#3954](https://github.com/Azure/aks-engine/issues/3954))
- add support for Kubernetes v1.18.10 on Azure Stack ([#3950](https://github.com/Azure/aks-engine/issues/3950))
- add support for Kubernetes v1.18.10 ([#3948](https://github.com/Azure/aks-engine/issues/3948))
- updates for container monitoring addon omsagent agent September 2020 release ([#3942](https://github.com/Azure/aks-engine/issues/3942))
- add support for Kubernetes v1.19.3 ([#3937](https://github.com/Azure/aks-engine/issues/3937))
- custom Windows log collection script ([#3940](https://github.com/Azure/aks-engine/issues/3940))
- enable system-assigned identity by default ([#3856](https://github.com/Azure/aks-engine/issues/3856))
- Target new windows VHD with new K8s binaries (1.19.2, 1.18.9, etc) ([#3905](https://github.com/Azure/aks-engine/issues/3905))
- add ScaleCPULimitsToSandbox for hyperv runtimeclasses ([#3889](https://github.com/Azure/aks-engine/issues/3889))
- allow custom containerd package for Linux nodes ([#3878](https://github.com/Azure/aks-engine/issues/3878))
- Updating Windows VHD build files to support building for multiple OS versions ([#3847](https://github.com/Azure/aks-engine/issues/3847))
- add support for Kubernetes v1.18.9 ([#3841](https://github.com/Azure/aks-engine/issues/3841))
- add support for Kubernetes v1.19.2 ([#3842](https://github.com/Azure/aks-engine/issues/3842))
- add support for Kubernetes v1.17.12 ([#3840](https://github.com/Azure/aks-engine/issues/3840))
- update VMSS node pools ([#3830](https://github.com/Azure/aks-engine/issues/3830))

### Maintenance 🔧
- Updating Windows VHDs for v0.61.0 release ([#4317](https://github.com/Azure/aks-engine/issues/4317))
- update Go toolchain to v1.15.8 ([#4307](https://github.com/Azure/aks-engine/issues/4307))
- Add Feb security patch ([#4298](https://github.com/Azure/aks-engine/issues/4298))
- update logrus to 1.8.0 ([#4310](https://github.com/Azure/aks-engine/issues/4310))
- rev Linux VHDs to 2021.02.22 ([#4296](https://github.com/Azure/aks-engine/issues/4296))
- don't run unattended upgrade if using "no outbound" test feature ([#4297](https://github.com/Azure/aks-engine/issues/4297))
- update node-driver-registrar and liveness-probe images ([#4278](https://github.com/Azure/aks-engine/issues/4278))
- deprecate kube-dashboard addon ([#4268](https://github.com/Azure/aks-engine/issues/4268))
- deprecate aci-connector addon ([#4276](https://github.com/Azure/aks-engine/issues/4276))
- remove addon-resizer from VHD ([#4269](https://github.com/Azure/aks-engine/issues/4269))
- deprecate rescheduler addon ([#4275](https://github.com/Azure/aks-engine/issues/4275))
- keep two versions of Azure CNI in VHD ([#4266](https://github.com/Azure/aks-engine/issues/4266))
- update pause image to 1.4.1 ([#4273](https://github.com/Azure/aks-engine/issues/4273))
- update CNI plugins to v0.9.1 ([#4267](https://github.com/Azure/aks-engine/issues/4267))
- update kube-addon-manager to v9.1.3 ([#4271](https://github.com/Azure/aks-engine/issues/4271))
- support only latest, tested versions of Kubernetes ([#4265](https://github.com/Azure/aks-engine/issues/4265))
- update cluster-autoscaler to 1.20.0 ([#4264](https://github.com/Azure/aks-engine/issues/4264))
- ensure containerd has image cache on VHDs ([#4249](https://github.com/Azure/aks-engine/issues/4249))
- generated code ([#4261](https://github.com/Azure/aks-engine/issues/4261))
- don't include auditd rules in Linux VHDs ([#4253](https://github.com/Azure/aks-engine/issues/4253))
- add T4 GPU as Nvidia GPUs ([#4259](https://github.com/Azure/aks-engine/issues/4259))
- deprecate support for creating new 1.16 clusters ([#4256](https://github.com/Azure/aks-engine/issues/4256))
- don't include gcr-sourced images in Linux VHDs ([#4255](https://github.com/Azure/aks-engine/issues/4255))
- prune non-default images from VHD config ([#4250](https://github.com/Azure/aks-engine/issues/4250))
- fix func arity mismatch in validate_test.go ([#4251](https://github.com/Azure/aks-engine/issues/4251))
- upgrade NPM to v1.2.2_hotfix ([#4225](https://github.com/Azure/aks-engine/issues/4225))
- update windows dockeree version to 19.03.14 ([#4229](https://github.com/Azure/aks-engine/issues/4229))
- add new Azure VM SKUs, brazilsoutheast, westus3 regions ([#4224](https://github.com/Azure/aks-engine/issues/4224))
- block container traffic to 168.63.129.16 ([#4212](https://github.com/Azure/aks-engine/issues/4212))
- rev Linux VHDs to 2021.01.28 ([#4223](https://github.com/Azure/aks-engine/issues/4223))
- update aks-engine VHD Windows VHD ([#4218](https://github.com/Azure/aks-engine/issues/4218))
- update csi-secrets-store to v0.0.19 and akv provider to 0.0.12 ([#4203](https://github.com/Azure/aks-engine/issues/4203))
- update policy addon deployment ([#4201](https://github.com/Azure/aks-engine/issues/4201))
- update adal to v0.9.10 ([#4200](https://github.com/Azure/aks-engine/issues/4200))
- updating windows provisioing scripts to v0.0.10 ([#4184](https://github.com/Azure/aks-engine/issues/4184))
- Update Azure CNI to v1.2.2 ([#4183](https://github.com/Azure/aks-engine/issues/4183))
- Add windowsnodelabelsync.ps1 ([#4163](https://github.com/Azure/aks-engine/issues/4163))
- install Jan 2021 security updates in Windows VHD ([#4168](https://github.com/Azure/aks-engine/issues/4168))
- bump kms keyvault to v0.0.10 ([#4169](https://github.com/Azure/aks-engine/issues/4169))
- Use signed scripts package 0.0.9 on Windows nodes ([#4162](https://github.com/Azure/aks-engine/issues/4162))
- release Windows VHD for v0.59.0 aks-engine release ([#4158](https://github.com/Azure/aks-engine/issues/4158))
- rev Linux VHDs to 2021.01.08 ([#4159](https://github.com/Azure/aks-engine/issues/4159))
- simplify upgrade templates ([#4135](https://github.com/Azure/aks-engine/issues/4135))
- reinforce upgrade warnings/errors ([#4074](https://github.com/Azure/aks-engine/issues/4074))
- Use signed scripts v0.0.8 for Windows deployments ([#4134](https://github.com/Azure/aks-engine/issues/4134))
- simplify scale templates ([#4131](https://github.com/Azure/aks-engine/issues/4131))
- update Azure CNI to v1.2.0_hotfix ([#4129](https://github.com/Azure/aks-engine/issues/4129))
- update vendor deps ([#4125](https://github.com/Azure/aks-engine/issues/4125))
- rev metrics-server to v0.4.1 for all k8s versions ([#4127](https://github.com/Azure/aks-engine/issues/4127))
- Update moby/containerd versions ([#4119](https://github.com/Azure/aks-engine/issues/4119))
- update csi-secrets-store to v0.0.18 and akv provider to 0.0.11 ([#4126](https://github.com/Azure/aks-engine/issues/4126))
- updating azure-npm to 1.2.1 version ([#4094](https://github.com/Azure/aks-engine/issues/4094))
- Installing Dec 2020 cumulative updates for Windows VHDs ([#4103](https://github.com/Azure/aks-engine/issues/4103))
- update adal to v0.9.6 ([#4093](https://github.com/Azure/aks-engine/issues/4093))
- limit number of upgrade retries if new CP nodes bootstrap fails ([#4068](https://github.com/Azure/aks-engine/issues/4068))
- faster rolling updates for daemonset addons ([#4090](https://github.com/Azure/aks-engine/issues/4090))
- update csi-secrets-store addon manifest and images (v0.0.10) ([#4084](https://github.com/Azure/aks-engine/issues/4084))
- update Windows image ([#4086](https://github.com/Azure/aks-engine/issues/4086))
- rev Linux VHDs to 2020.12.02 ([#4081](https://github.com/Azure/aks-engine/issues/4081))
- k8s v1.18 conformance model for Azure Stack ([#4070](https://github.com/Azure/aks-engine/issues/4070))
- validate VHD availability before upgrade/scale on Azure Stack Hub ([#4062](https://github.com/Azure/aks-engine/issues/4062))
- mark Kubernetes 1.17.14 as disabled ([#4044](https://github.com/Azure/aks-engine/issues/4044))
- cleanup dead code ([#4053](https://github.com/Azure/aks-engine/issues/4053))
- only warn about master stuff during create ([#4057](https://github.com/Azure/aks-engine/issues/4057))
- Upgrade CNI to v1.2.0 ([#4058](https://github.com/Azure/aks-engine/issues/4058))
- Updating default csi-proxy version to v0.2.2 ([#4047](https://github.com/Azure/aks-engine/issues/4047))
- format some json in addpool.md ([#4049](https://github.com/Azure/aks-engine/issues/4049))
- remove deprecated AKS code paths ([#4040](https://github.com/Azure/aks-engine/issues/4040))
- deprecate orchestratorType ([#4038](https://github.com/Azure/aks-engine/issues/4038))
- remove deprecated localizations ([#4036](https://github.com/Azure/aks-engine/issues/4036))
- Use Windows November(11b) updates as default vhd ([#4033](https://github.com/Azure/aks-engine/issues/4033))
- Windows November Patches ([#4023](https://github.com/Azure/aks-engine/issues/4023))
- Add 10c image as default ([#4020](https://github.com/Azure/aks-engine/issues/4020))
- rev Linux VHDs to 2020.10.30 ([#4016](https://github.com/Azure/aks-engine/issues/4016))
- Install Windows Server 2019 10C updates in Windows VHD ([#3956](https://github.com/Azure/aks-engine/issues/3956))
- update powershell signed scripts ([#4012](https://github.com/Azure/aks-engine/issues/4012))
- Adding csi-proxy-v0.2.2 to Windows VHD ([#4004](https://github.com/Azure/aks-engine/issues/4004))
- update cluster-autoscalers to latest patches ([#4000](https://github.com/Azure/aks-engine/issues/4000))
- set go sdk log level for Azure Stack clusters ([#3993](https://github.com/Azure/aks-engine/issues/3993))
- update cluster-autoscaler to v1.19.1 ([#3996](https://github.com/Azure/aks-engine/issues/3996))
- update go toolchain to v1.15.3 ([#3879](https://github.com/Azure/aks-engine/issues/3879))
- use transparent mode for Azure CNI ([#3958](https://github.com/Azure/aks-engine/issues/3958))
- gofmt ([#3982](https://github.com/Azure/aks-engine/issues/3982))
- longer default cordonDrainTimeout for Azure Stack Cloud ([#3969](https://github.com/Azure/aks-engine/issues/3969))
- kubelet systemd job depends on CRI service ([#3943](https://github.com/Azure/aks-engine/issues/3943))
- Upgrade CNI to v1.1.8 ([#3907](https://github.com/Azure/aks-engine/issues/3907))
- remove support for Kubernetes 1.15 ([#3751](https://github.com/Azure/aks-engine/issues/3751))
- bump calico to 3.8.9 to get latest patches ([#3924](https://github.com/Azure/aks-engine/issues/3924))
- set the csi-secrets-store to have a priority class ([#3909](https://github.com/Azure/aks-engine/issues/3909))
- rev Linux VHDs to 2020.10.06 ([#3906](https://github.com/Azure/aks-engine/issues/3906))
- enable VHD re-use in no outbound scenarios ([#3897](https://github.com/Azure/aks-engine/issues/3897))
- update csi-secrets-store addon manifest and images (v0.0.9) ([#3891](https://github.com/Azure/aks-engine/issues/3891))
- create azure.json via CSE ([#3876](https://github.com/Azure/aks-engine/issues/3876))
- distribute apiserver.crt to control plane nodes only ([#3860](https://github.com/Azure/aks-engine/issues/3860))
- update azure cni to 1.1.7 ([#3864](https://github.com/Azure/aks-engine/issues/3864))
- update Dashboard addon to v2.0.4 ([#3855](https://github.com/Azure/aks-engine/issues/3855))
- check VHD media name length before being published to Marketplace ([#3799](https://github.com/Azure/aks-engine/issues/3799))
- remove no-op 1.15 version checks in templates ([#3851](https://github.com/Azure/aks-engine/issues/3851))
- updating Windows VHD with new cached artifacts ([#3843](https://github.com/Azure/aks-engine/issues/3843))
- adding 1.19.1 bits to Windows VHD ([#3834](https://github.com/Azure/aks-engine/issues/3834))
- rev Linux VHDs to 2020.09.14 ([#3827](https://github.com/Azure/aks-engine/issues/3827))
- update signed powershell script package to include azure CNI fixes ([#3829](https://github.com/Azure/aks-engine/issues/3829))

### Revert Change ◀️
- "chore: targeting sept updates for Windows 2019 VHD ([#3801](https://github.com/Azure/aks-engine/issues/3801))" ([#3836](https://github.com/Azure/aks-engine/issues/3836))

### Testing 💚
- enable parallel E2E tests, update ginkgo to 1.5 ([#4290](https://github.com/Azure/aks-engine/issues/4290))
- use AZURE_CORE_ONLY_SHOW_ERRORS when running E2E ([#4288](https://github.com/Azure/aks-engine/issues/4288))
- expose custom image options as env vars ([#4244](https://github.com/Azure/aks-engine/issues/4244))
- don't test availability sets + 1.21 ([#4239](https://github.com/Azure/aks-engine/issues/4239))
- E2E resilience ([#4235](https://github.com/Azure/aks-engine/issues/4235))
- go routine errata ([#4233](https://github.com/Azure/aks-engine/issues/4233))
- tolerate long SIG image publication times for vmss-prototype test ([#4228](https://github.com/Azure/aks-engine/issues/4228))
- incorporate kured + auto mode into kamino vmss-prototype tests ([#4221](https://github.com/Azure/aks-engine/issues/4221))
- add timestamps to E2E pod logs output ([#4216](https://github.com/Azure/aks-engine/issues/4216))
- enable kamino vmss-prototype dry run tests ([#4215](https://github.com/Azure/aks-engine/issues/4215))
- increase timeout tolerance for vmss-prototype SIG publishing ([#4211](https://github.com/Azure/aks-engine/issues/4211))
- add more stdout to kamino vmss-prototype E2E tests ([#4210](https://github.com/Azure/aks-engine/issues/4210))
- print helm command for vmss-prototype test ([#4207](https://github.com/Azure/aks-engine/issues/4207))
- correct E2E message during vmss-prototype test ([#4205](https://github.com/Azure/aks-engine/issues/4205))
- kamino E2E errata ([#4204](https://github.com/Azure/aks-engine/issues/4204))
- remove dangling fmt.Println ([#4202](https://github.com/Azure/aks-engine/issues/4202))
- enable testing custom kamino vmss prototype images ([#4198](https://github.com/Azure/aks-engine/issues/4198))
- updated kamino vmss-prototype integration E2E ([#4153](https://github.com/Azure/aks-engine/issues/4153))
- ensure hostPort routing ([#4186](https://github.com/Azure/aks-engine/issues/4186))
- enable configurable node prototype tests ([#4100](https://github.com/Azure/aks-engine/issues/4100))
- enable GINKGO_SKIP_AFTER_UPGRADE param in e2e test ([#4089](https://github.com/Azure/aks-engine/issues/4089))
- more resilient azure-arc-onboarding E2E ([#4069](https://github.com/Azure/aks-engine/issues/4069))
- use 3 control plane VMs for base test config ([#4075](https://github.com/Azure/aks-engine/issues/4075))
- dns-liveness livenessProbe tweaks ([#4080](https://github.com/Azure/aks-engine/issues/4080))
- disable useManagedIdentity default value on Azure Stack ([#4063](https://github.com/Azure/aks-engine/issues/4063))
- skip rotate docker log tests if omsagent ([#4059](https://github.com/Azure/aks-engine/issues/4059))
- only test flannel + containerd thru k8s 1.19 ([#4041](https://github.com/Azure/aks-engine/issues/4041))
- re-engage cluster-autoscaler tests if no ssh ([#4034](https://github.com/Azure/aks-engine/issues/4034))
- add dns-loop (actually curl) stress test ([#4024](https://github.com/Azure/aks-engine/issues/4024))
- configurable stability iterations timeout ([#4022](https://github.com/Azure/aks-engine/issues/4022))
- fix two UTs to work with "cli" auth ([#3995](https://github.com/Azure/aks-engine/issues/3995))
- ensure azure-kms-provider E2E coverage ([#3985](https://github.com/Azure/aks-engine/issues/3985))
- "useManagedIdentity": false for VMSS masters ([#3981](https://github.com/Azure/aks-engine/issues/3981))
- Retry label when cluster first comes online ([#3976](https://github.com/Azure/aks-engine/issues/3976))
- Jenkinsfile bool vars ([#3980](https://github.com/Azure/aks-engine/issues/3980))
- Windows Metrics failure when High CPU  ([#3962](https://github.com/Azure/aks-engine/issues/3962))
- validate azure arc in everything config ([#3961](https://github.com/Azure/aks-engine/issues/3961))
- CCM + azurefile + MSI test scenarios ([#3932](https://github.com/Azure/aks-engine/issues/3932))
- skip tests for flatcar cluster config ([#3921](https://github.com/Azure/aks-engine/issues/3921))
- Add azure.json path for custom cloud k8s config & Update stability timeout for Azure CNI network policy ([#3895](https://github.com/Azure/aks-engine/issues/3895))
- E2E fix scale test indexing out of bounds ([#3873](https://github.com/Azure/aks-engine/issues/3873))
- not nil defense broke logic ([#3881](https://github.com/Azure/aks-engine/issues/3881))
- fixes for akse e2e tests ([#3874](https://github.com/Azure/aks-engine/issues/3874))
- E2E optionally validates rotate-certs ([#3866](https://github.com/Azure/aks-engine/issues/3866))
- add more E2E timeout tolerance ([#3850](https://github.com/Azure/aks-engine/issues/3850))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.56.2"></a>
# [v0.56.2] - 2021-02-11
### Bug Fixes 🐞
- Azure Stack CNI network interfaces file creation fix ([#3792](https://github.com/Azure/aks-engine/issues/3792))

### Maintenance 🔧
- add T4 GPU as Nvidia GPUs ([#4259](https://github.com/Azure/aks-engine/issues/4259))
- add new Azure VM SKUs, brazilsoutheast, westus3 regions ([#4224](https://github.com/Azure/aks-engine/issues/4224))
- updating Windows VHD with new cached artifacts ([#3843](https://github.com/Azure/aks-engine/issues/3843))
- adding 1.19.1 bits to Windows VHD ([#3834](https://github.com/Azure/aks-engine/issues/3834))
- rev Linux VHDs to 2020.09.14 ([#3827](https://github.com/Azure/aks-engine/issues/3827))
- update signed powershell script package to include azure CNI fixes ([#3829](https://github.com/Azure/aks-engine/issues/3829))

### Revert Change ◀️
- "chore: targeting sept updates for Windows 2019 VHD ([#3801](https://github.com/Azure/aks-engine/issues/3801))" ([#3836](https://github.com/Azure/aks-engine/issues/3836))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.60.1"></a>
# [v0.60.1] - 2021-02-11
### Bug Fixes 🐞
- rotate-certs on custom clouds ([#4248](https://github.com/Azure/aks-engine/issues/4248))
- use curl.exe and add retries when downloading artifacts to Windows VHD ([#4206](https://github.com/Azure/aks-engine/issues/4206))
- update nodecustomdata to enable portmapper for kubenet+containerd ([#4191](https://github.com/Azure/aks-engine/issues/4191))
- fix network cleanup code on windows for contianerd nodes ([#4154](https://github.com/Azure/aks-engine/issues/4154))
- update error code for vhd not found ([#4150](https://github.com/Azure/aks-engine/issues/4150))
- fix that kubeproxy cannot parse featuregates ([#4145](https://github.com/Azure/aks-engine/issues/4145))
- Fixing URl for azure-cni v1.20_hotfix package in Windows VHD ([#4136](https://github.com/Azure/aks-engine/issues/4136))
- ensure we wait for the right CRI service name ([#4115](https://github.com/Azure/aks-engine/issues/4115))
- honor user-defined feature gates before applying defaults ([#4113](https://github.com/Azure/aks-engine/issues/4113))
- typo in upgrade log message ([#4109](https://github.com/Azure/aks-engine/issues/4109))
- filter out KeyVault resources during upgrade ([#4072](https://github.com/Azure/aks-engine/issues/4072))
- make system role assignments work w/ upgrade ([#4078](https://github.com/Azure/aks-engine/issues/4078))
- cleanup VMSSName + addpool ([#4067](https://github.com/Azure/aks-engine/issues/4067))
- enforce Azure CNI config via jq instead of sed ([#4060](https://github.com/Azure/aks-engine/issues/4060))
- persist VMSS name in API model ([#4051](https://github.com/Azure/aks-engine/issues/4051))
- Windows and VMSS version validation for azs ([#4030](https://github.com/Azure/aks-engine/issues/4030))
- always ensure etcd on control plane vms ([#4045](https://github.com/Azure/aks-engine/issues/4045))
- fixing an issue where windows cannot start kubelet to get podCIDR on Windows nodes when using kubenet ([#4039](https://github.com/Azure/aks-engine/issues/4039))
- Adding csi-proxy logs to prevent access denied errors ([#4029](https://github.com/Azure/aks-engine/issues/4029))
- upgrade broken for VMAS + managed identities ([#4021](https://github.com/Azure/aks-engine/issues/4021))
- enable dhcpv6 only for ubuntu 16.04 ([#4017](https://github.com/Azure/aks-engine/issues/4017))
- Re-ordering HNS policy removal due to 10c change in behavior and consolidating logic in Windows cleanup scripts  ([#4002](https://github.com/Azure/aks-engine/issues/4002))
- Mount /usr/local/share/ca-certificates folder into kube-controller-manager ([#4001](https://github.com/Azure/aks-engine/issues/4001))
- cancel Pod Eviction jobs when timeout ([#3970](https://github.com/Azure/aks-engine/issues/3970))
- ensure flatcar configs use transparent Azure CNI ([#3972](https://github.com/Azure/aks-engine/issues/3972))
- page not done logic bug ([#3971](https://github.com/Azure/aks-engine/issues/3971))
- resource identifier not passed to custom cloud profile ([#3953](https://github.com/Azure/aks-engine/issues/3953))
- get-logs command collects control plane logs even if apiserver is down ([#3939](https://github.com/Azure/aks-engine/issues/3939))
- use InternalIP for metrics-server DNS resolution ([#3929](https://github.com/Azure/aks-engine/issues/3929))
- VMSS control plane + availability zones ([#3917](https://github.com/Azure/aks-engine/issues/3917))
- Azure Stack upgrade operations clear old MCR ImageBase ([#3922](https://github.com/Azure/aks-engine/issues/3922))
- remove NSG rules for AzureStack ([#3914](https://github.com/Azure/aks-engine/issues/3914))
- fix jumpbox custom data ([#3896](https://github.com/Azure/aks-engine/issues/3896))
- don't warn about CRI update on cluster create ([#3890](https://github.com/Azure/aks-engine/issues/3890))
- ensure addon hostNetwork ports don't conflict ([#3894](https://github.com/Azure/aks-engine/issues/3894))
- de-dup vnet roleAssignment deployments ([#3857](https://github.com/Azure/aks-engine/issues/3857))
- add #EOF sentinel chars to custom search file ([#3862](https://github.com/Azure/aks-engine/issues/3862))
- E2E scale scenario broken if update is false ([#3852](https://github.com/Azure/aks-engine/issues/3852))
- Azure Stack CNI network interfaces file creation fix ([#3792](https://github.com/Azure/aks-engine/issues/3792))

### Code Refactoring 💎
- remove unsupported orchestrators ([#3965](https://github.com/Azure/aks-engine/issues/3965))
- move ssh-related funcs to its own package ([#3990](https://github.com/Azure/aks-engine/issues/3990))
- delete Azure Stack's KubernetesClient implementation ([#3957](https://github.com/Azure/aks-engine/issues/3957))

### Continuous Integration 💜
- update dev image to deis/docker-go-dev v1.29.0 ([#4195](https://github.com/Azure/aks-engine/issues/4195))
- remove KMS test from everything cluster config ([#4194](https://github.com/Azure/aks-engine/issues/4194))
- add manual dispatch trigger to nightly build ([#4123](https://github.com/Azure/aks-engine/issues/4123))
- use 3 replicas for no-vnet E2E jobs ([#4061](https://github.com/Azure/aks-engine/issues/4061))
- Updating pr-windows-signed-scripts.yaml to allow for overriding cluster definition ([#4015](https://github.com/Azure/aks-engine/issues/4015))
- remove node-problem-detector addon from PR E2E ([#4010](https://github.com/Azure/aks-engine/issues/4010))
- Adding CI pipeline to verify windows provisioning script content BEFORE it gets check in ([#4008](https://github.com/Azure/aks-engine/issues/4008))
- Windows VHD pipeline tests ([#3977](https://github.com/Azure/aks-engine/issues/3977))
- reduce addons area for containerd config tests ([#3923](https://github.com/Azure/aks-engine/issues/3923))
- adding github actions to create nightly builds ([#3904](https://github.com/Azure/aks-engine/issues/3904))
- E2E fixes to allow more containerd tests ([#3903](https://github.com/Azure/aks-engine/issues/3903))
- add containerd PR E2E tests ([#3892](https://github.com/Azure/aks-engine/issues/3892))
- update OWNERS ([#3853](https://github.com/Azure/aks-engine/issues/3853))

### Documentation 📘
- add dual-stack iptables api model ([#4185](https://github.com/Azure/aks-engine/issues/4185))
- aad not supported on AZs ([#4182](https://github.com/Azure/aks-engine/issues/4182))
- fix broken link in AAD docs ([#4180](https://github.com/Azure/aks-engine/issues/4180))
- support FAQ ([#4148](https://github.com/Azure/aks-engine/issues/4148))
- remove warning about Windows + custom VNET ([#4147](https://github.com/Azure/aks-engine/issues/4147))
- Clarify README for Azure Stack Hub ([#4128](https://github.com/Azure/aks-engine/issues/4128))
- Added references to capz ([#4124](https://github.com/Azure/aks-engine/issues/4124))
- clarify msi as default, and requirements ([#4122](https://github.com/Azure/aks-engine/issues/4122))
- update dualstack docs to 1.20 ([#4101](https://github.com/Azure/aks-engine/issues/4101))
- fixed broken link in Azure Stack topic page ([#4056](https://github.com/Azure/aks-engine/issues/4056))
- remove double quotes(") at the aks-deploy sample ([#3967](https://github.com/Azure/aks-engine/issues/3967))
- adding instructions for how how to build the Windows VHD for di… ([#3911](https://github.com/Azure/aks-engine/issues/3911))
- confirm flannel + docker is not supported ([#3886](https://github.com/Azure/aks-engine/issues/3886))
- add notes about not upgrading LB config ([#3884](https://github.com/Azure/aks-engine/issues/3884))
- clarify that update is for node pools only ([#3877](https://github.com/Azure/aks-engine/issues/3877))
- CLI operations ([#3837](https://github.com/Azure/aks-engine/issues/3837))
- mark rotate-certs command as experimental ([#3869](https://github.com/Azure/aks-engine/issues/3869))
- add v0.55.4 to Azure Stack topic page ([#3846](https://github.com/Azure/aks-engine/issues/3846))
- remove mention about library for AKS ([#3835](https://github.com/Azure/aks-engine/issues/3835))
- add required attribution reminder to PR template ([#3826](https://github.com/Azure/aks-engine/issues/3826))

### Features 🌈
- rotate-certs fails faster if invalid ssh params ([#4252](https://github.com/Azure/aks-engine/issues/4252))
- run accelerated unattended-upgrade at node creation time ([#4217](https://github.com/Azure/aks-engine/issues/4217))
- reworked rotate-certs command ([#4214](https://github.com/Azure/aks-engine/issues/4214))
- add support for Kubernetes v1.16.15 and v1.18.15 on Azure Stack ([#4187](https://github.com/Azure/aks-engine/issues/4187))
- add support for Kubernetes v1.17.17 on Azure Stack ([#4188](https://github.com/Azure/aks-engine/issues/4188))
- add support for Kubernetes 1.20.2 ([#4192](https://github.com/Azure/aks-engine/issues/4192))
- create kms key as part of cluster bootstrap ([#4170](https://github.com/Azure/aks-engine/issues/4170))
- add support for Kubernetes v1.21.0-alpha.1 ([#4178](https://github.com/Azure/aks-engine/issues/4178))
- add support for Kubernetes v1.18.15 ([#4166](https://github.com/Azure/aks-engine/issues/4166))
- add support for Kubernetes v1.17.17 ([#4167](https://github.com/Azure/aks-engine/issues/4167))
- add support for Kubernetes v1.19.7 ([#4165](https://github.com/Azure/aks-engine/issues/4165))
- setting a default containerd package for Windows ([#4149](https://github.com/Azure/aks-engine/issues/4149))
- add support for Kubernetes v1.18.14 ([#4139](https://github.com/Azure/aks-engine/issues/4139))
- add support for Kubernetes v1.19.6 ([#4140](https://github.com/Azure/aks-engine/issues/4140))
- add support for Kubernetes v1.17.16 ([#4138](https://github.com/Azure/aks-engine/issues/4138))
- add support for Kubernetes v1.20.1 ([#4141](https://github.com/Azure/aks-engine/issues/4141))
- Add WinDSR support ([#4104](https://github.com/Azure/aks-engine/issues/4104))
- Updating kubelet/kube-proxy to run with as high priority processes ([#4073](https://github.com/Azure/aks-engine/issues/4073))
- Add support for specifying linux moby url ([#4120](https://github.com/Azure/aks-engine/issues/4120))
- Updating Windows VHDs to include Dec 2012 patches ([#4118](https://github.com/Azure/aks-engine/issues/4118))
- add support for Kubernetes v1.18.13 ([#4111](https://github.com/Azure/aks-engine/issues/4111))
- add support for Kubernetes v1.19.5 ([#4110](https://github.com/Azure/aks-engine/issues/4110))
- add support for Kubernetes v1.20.0 ([#4102](https://github.com/Azure/aks-engine/issues/4102))
- disable livenessProbe timeout enforcement ([#4085](https://github.com/Azure/aks-engine/issues/4085))
- add support for Kubernetes v1.20.0-rc.0 ([#4076](https://github.com/Azure/aks-engine/issues/4076))
- Enable chrony and host-based time sync by default on Ubuntu 18.04 ([#4011](https://github.com/Azure/aks-engine/issues/4011))
- add support for Kubernetes v1.19.4 ([#4042](https://github.com/Azure/aks-engine/issues/4042))
- add support for Kubernetes 1.18.12 ([#4043](https://github.com/Azure/aks-engine/issues/4043))
- add support for Kubernetes v1.20.0-beta.1 ([#3999](https://github.com/Azure/aks-engine/issues/3999))
- set apiserver tokenRequest flags in 1.20+ ([#3989](https://github.com/Azure/aks-engine/issues/3989))
- add support for Kubernetes v1.20.0-beta.0 ([#3991](https://github.com/Azure/aks-engine/issues/3991))
- add support for Kubernetes v1.20.0-alpha.3 ([#3934](https://github.com/Azure/aks-engine/issues/3934))
- upload collected logs to storage account container ([#3944](https://github.com/Azure/aks-engine/issues/3944))
- add support for Kubernetes v1.17.13 ([#3954](https://github.com/Azure/aks-engine/issues/3954))
- add support for Kubernetes v1.18.10 on Azure Stack ([#3950](https://github.com/Azure/aks-engine/issues/3950))
- add support for Kubernetes v1.18.10 ([#3948](https://github.com/Azure/aks-engine/issues/3948))
- updates for container monitoring addon omsagent agent September 2020 release ([#3942](https://github.com/Azure/aks-engine/issues/3942))
- add support for Kubernetes v1.19.3 ([#3937](https://github.com/Azure/aks-engine/issues/3937))
- custom Windows log collection script ([#3940](https://github.com/Azure/aks-engine/issues/3940))
- enable system-assigned identity by default ([#3856](https://github.com/Azure/aks-engine/issues/3856))
- Target new windows VHD with new K8s binaries (1.19.2, 1.18.9, etc) ([#3905](https://github.com/Azure/aks-engine/issues/3905))
- add ScaleCPULimitsToSandbox for hyperv runtimeclasses ([#3889](https://github.com/Azure/aks-engine/issues/3889))
- allow custom containerd package for Linux nodes ([#3878](https://github.com/Azure/aks-engine/issues/3878))
- Updating Windows VHD build files to support building for multiple OS versions ([#3847](https://github.com/Azure/aks-engine/issues/3847))
- add support for Kubernetes v1.18.9 ([#3841](https://github.com/Azure/aks-engine/issues/3841))
- add support for Kubernetes v1.19.2 ([#3842](https://github.com/Azure/aks-engine/issues/3842))
- add support for Kubernetes v1.17.12 ([#3840](https://github.com/Azure/aks-engine/issues/3840))
- update VMSS node pools ([#3830](https://github.com/Azure/aks-engine/issues/3830))

### Maintenance 🔧
- add T4 GPU as Nvidia GPUs ([#4259](https://github.com/Azure/aks-engine/issues/4259))
- add new Azure VM SKUs, brazilsoutheast, westus3 regions ([#4224](https://github.com/Azure/aks-engine/issues/4224))
- block container traffic to 168.63.129.16 ([#4212](https://github.com/Azure/aks-engine/issues/4212))
- rev Linux VHDs to 2021.01.28 ([#4223](https://github.com/Azure/aks-engine/issues/4223))
- update aks-engine VHD Windows VHD ([#4218](https://github.com/Azure/aks-engine/issues/4218))
- update csi-secrets-store to v0.0.19 and akv provider to 0.0.12 ([#4203](https://github.com/Azure/aks-engine/issues/4203))
- update policy addon deployment ([#4201](https://github.com/Azure/aks-engine/issues/4201))
- update adal to v0.9.10 ([#4200](https://github.com/Azure/aks-engine/issues/4200))
- updating windows provisioing scripts to v0.0.10 ([#4184](https://github.com/Azure/aks-engine/issues/4184))
- Update Azure CNI to v1.2.2 ([#4183](https://github.com/Azure/aks-engine/issues/4183))
- Add windowsnodelabelsync.ps1 ([#4163](https://github.com/Azure/aks-engine/issues/4163))
- install Jan 2021 security updates in Windows VHD ([#4168](https://github.com/Azure/aks-engine/issues/4168))
- bump kms keyvault to v0.0.10 ([#4169](https://github.com/Azure/aks-engine/issues/4169))
- Use signed scripts package 0.0.9 on Windows nodes ([#4162](https://github.com/Azure/aks-engine/issues/4162))
- release Windows VHD for v0.59.0 aks-engine release ([#4158](https://github.com/Azure/aks-engine/issues/4158))
- rev Linux VHDs to 2021.01.08 ([#4159](https://github.com/Azure/aks-engine/issues/4159))
- simplify upgrade templates ([#4135](https://github.com/Azure/aks-engine/issues/4135))
- reinforce upgrade warnings/errors ([#4074](https://github.com/Azure/aks-engine/issues/4074))
- Use signed scripts v0.0.8 for Windows deployments ([#4134](https://github.com/Azure/aks-engine/issues/4134))
- simplify scale templates ([#4131](https://github.com/Azure/aks-engine/issues/4131))
- update Azure CNI to v1.2.0_hotfix ([#4129](https://github.com/Azure/aks-engine/issues/4129))
- update vendor deps ([#4125](https://github.com/Azure/aks-engine/issues/4125))
- rev metrics-server to v0.4.1 for all k8s versions ([#4127](https://github.com/Azure/aks-engine/issues/4127))
- Update moby/containerd versions ([#4119](https://github.com/Azure/aks-engine/issues/4119))
- update csi-secrets-store to v0.0.18 and akv provider to 0.0.11 ([#4126](https://github.com/Azure/aks-engine/issues/4126))
- updating azure-npm to 1.2.1 version ([#4094](https://github.com/Azure/aks-engine/issues/4094))
- Installing Dec 2020 cumulative updates for Windows VHDs ([#4103](https://github.com/Azure/aks-engine/issues/4103))
- update adal to v0.9.6 ([#4093](https://github.com/Azure/aks-engine/issues/4093))
- limit number of upgrade retries if new CP nodes bootstrap fails ([#4068](https://github.com/Azure/aks-engine/issues/4068))
- faster rolling updates for daemonset addons ([#4090](https://github.com/Azure/aks-engine/issues/4090))
- update csi-secrets-store addon manifest and images (v0.0.10) ([#4084](https://github.com/Azure/aks-engine/issues/4084))
- update Windows image ([#4086](https://github.com/Azure/aks-engine/issues/4086))
- rev Linux VHDs to 2020.12.02 ([#4081](https://github.com/Azure/aks-engine/issues/4081))
- k8s v1.18 conformance model for Azure Stack ([#4070](https://github.com/Azure/aks-engine/issues/4070))
- validate VHD availability before upgrade/scale on Azure Stack Hub ([#4062](https://github.com/Azure/aks-engine/issues/4062))
- mark Kubernetes 1.17.14 as disabled ([#4044](https://github.com/Azure/aks-engine/issues/4044))
- cleanup dead code ([#4053](https://github.com/Azure/aks-engine/issues/4053))
- only warn about master stuff during create ([#4057](https://github.com/Azure/aks-engine/issues/4057))
- Upgrade CNI to v1.2.0 ([#4058](https://github.com/Azure/aks-engine/issues/4058))
- Updating default csi-proxy version to v0.2.2 ([#4047](https://github.com/Azure/aks-engine/issues/4047))
- format some json in addpool.md ([#4049](https://github.com/Azure/aks-engine/issues/4049))
- remove deprecated AKS code paths ([#4040](https://github.com/Azure/aks-engine/issues/4040))
- deprecate orchestratorType ([#4038](https://github.com/Azure/aks-engine/issues/4038))
- remove deprecated localizations ([#4036](https://github.com/Azure/aks-engine/issues/4036))
- Use Windows November(11b) updates as default vhd ([#4033](https://github.com/Azure/aks-engine/issues/4033))
- Windows November Patches ([#4023](https://github.com/Azure/aks-engine/issues/4023))
- Add 10c image as default ([#4020](https://github.com/Azure/aks-engine/issues/4020))
- rev Linux VHDs to 2020.10.30 ([#4016](https://github.com/Azure/aks-engine/issues/4016))
- Install Windows Server 2019 10C updates in Windows VHD ([#3956](https://github.com/Azure/aks-engine/issues/3956))
- update powershell signed scripts ([#4012](https://github.com/Azure/aks-engine/issues/4012))
- Adding csi-proxy-v0.2.2 to Windows VHD ([#4004](https://github.com/Azure/aks-engine/issues/4004))
- update cluster-autoscalers to latest patches ([#4000](https://github.com/Azure/aks-engine/issues/4000))
- set go sdk log level for Azure Stack clusters ([#3993](https://github.com/Azure/aks-engine/issues/3993))
- update cluster-autoscaler to v1.19.1 ([#3996](https://github.com/Azure/aks-engine/issues/3996))
- update go toolchain to v1.15.3 ([#3879](https://github.com/Azure/aks-engine/issues/3879))
- use transparent mode for Azure CNI ([#3958](https://github.com/Azure/aks-engine/issues/3958))
- gofmt ([#3982](https://github.com/Azure/aks-engine/issues/3982))
- longer default cordonDrainTimeout for Azure Stack Cloud ([#3969](https://github.com/Azure/aks-engine/issues/3969))
- kubelet systemd job depends on CRI service ([#3943](https://github.com/Azure/aks-engine/issues/3943))
- Upgrade CNI to v1.1.8 ([#3907](https://github.com/Azure/aks-engine/issues/3907))
- remove support for Kubernetes 1.15 ([#3751](https://github.com/Azure/aks-engine/issues/3751))
- bump calico to 3.8.9 to get latest patches ([#3924](https://github.com/Azure/aks-engine/issues/3924))
- set the csi-secrets-store to have a priority class ([#3909](https://github.com/Azure/aks-engine/issues/3909))
- rev Linux VHDs to 2020.10.06 ([#3906](https://github.com/Azure/aks-engine/issues/3906))
- enable VHD re-use in no outbound scenarios ([#3897](https://github.com/Azure/aks-engine/issues/3897))
- update csi-secrets-store addon manifest and images (v0.0.9) ([#3891](https://github.com/Azure/aks-engine/issues/3891))
- create azure.json via CSE ([#3876](https://github.com/Azure/aks-engine/issues/3876))
- distribute apiserver.crt to control plane nodes only ([#3860](https://github.com/Azure/aks-engine/issues/3860))
- update azure cni to 1.1.7 ([#3864](https://github.com/Azure/aks-engine/issues/3864))
- update Dashboard addon to v2.0.4 ([#3855](https://github.com/Azure/aks-engine/issues/3855))
- check VHD media name length before being published to Marketplace ([#3799](https://github.com/Azure/aks-engine/issues/3799))
- remove no-op 1.15 version checks in templates ([#3851](https://github.com/Azure/aks-engine/issues/3851))
- updating Windows VHD with new cached artifacts ([#3843](https://github.com/Azure/aks-engine/issues/3843))
- adding 1.19.1 bits to Windows VHD ([#3834](https://github.com/Azure/aks-engine/issues/3834))
- rev Linux VHDs to 2020.09.14 ([#3827](https://github.com/Azure/aks-engine/issues/3827))
- update signed powershell script package to include azure CNI fixes ([#3829](https://github.com/Azure/aks-engine/issues/3829))

### Revert Change ◀️
- "chore: targeting sept updates for Windows 2019 VHD ([#3801](https://github.com/Azure/aks-engine/issues/3801))" ([#3836](https://github.com/Azure/aks-engine/issues/3836))

### Testing 💚
- incorporate kured + auto mode into kamino vmss-prototype tests ([#4221](https://github.com/Azure/aks-engine/issues/4221))
- add timestamps to E2E pod logs output ([#4216](https://github.com/Azure/aks-engine/issues/4216))
- enable kamino vmss-prototype dry run tests ([#4215](https://github.com/Azure/aks-engine/issues/4215))
- increase timeout tolerance for vmss-prototype SIG publishing ([#4211](https://github.com/Azure/aks-engine/issues/4211))
- add more stdout to kamino vmss-prototype E2E tests ([#4210](https://github.com/Azure/aks-engine/issues/4210))
- print helm command for vmss-prototype test ([#4207](https://github.com/Azure/aks-engine/issues/4207))
- correct E2E message during vmss-prototype test ([#4205](https://github.com/Azure/aks-engine/issues/4205))
- kamino E2E errata ([#4204](https://github.com/Azure/aks-engine/issues/4204))
- remove dangling fmt.Println ([#4202](https://github.com/Azure/aks-engine/issues/4202))
- enable testing custom kamino vmss prototype images ([#4198](https://github.com/Azure/aks-engine/issues/4198))
- updated kamino vmss-prototype integration E2E ([#4153](https://github.com/Azure/aks-engine/issues/4153))
- ensure hostPort routing ([#4186](https://github.com/Azure/aks-engine/issues/4186))
- enable configurable node prototype tests ([#4100](https://github.com/Azure/aks-engine/issues/4100))
- enable GINKGO_SKIP_AFTER_UPGRADE param in e2e test ([#4089](https://github.com/Azure/aks-engine/issues/4089))
- more resilient azure-arc-onboarding E2E ([#4069](https://github.com/Azure/aks-engine/issues/4069))
- use 3 control plane VMs for base test config ([#4075](https://github.com/Azure/aks-engine/issues/4075))
- dns-liveness livenessProbe tweaks ([#4080](https://github.com/Azure/aks-engine/issues/4080))
- disable useManagedIdentity default value on Azure Stack ([#4063](https://github.com/Azure/aks-engine/issues/4063))
- skip rotate docker log tests if omsagent ([#4059](https://github.com/Azure/aks-engine/issues/4059))
- only test flannel + containerd thru k8s 1.19 ([#4041](https://github.com/Azure/aks-engine/issues/4041))
- re-engage cluster-autoscaler tests if no ssh ([#4034](https://github.com/Azure/aks-engine/issues/4034))
- add dns-loop (actually curl) stress test ([#4024](https://github.com/Azure/aks-engine/issues/4024))
- configurable stability iterations timeout ([#4022](https://github.com/Azure/aks-engine/issues/4022))
- fix two UTs to work with "cli" auth ([#3995](https://github.com/Azure/aks-engine/issues/3995))
- ensure azure-kms-provider E2E coverage ([#3985](https://github.com/Azure/aks-engine/issues/3985))
- "useManagedIdentity": false for VMSS masters ([#3981](https://github.com/Azure/aks-engine/issues/3981))
- Retry label when cluster first comes online ([#3976](https://github.com/Azure/aks-engine/issues/3976))
- Jenkinsfile bool vars ([#3980](https://github.com/Azure/aks-engine/issues/3980))
- Windows Metrics failure when High CPU  ([#3962](https://github.com/Azure/aks-engine/issues/3962))
- validate azure arc in everything config ([#3961](https://github.com/Azure/aks-engine/issues/3961))
- CCM + azurefile + MSI test scenarios ([#3932](https://github.com/Azure/aks-engine/issues/3932))
- skip tests for flatcar cluster config ([#3921](https://github.com/Azure/aks-engine/issues/3921))
- Add azure.json path for custom cloud k8s config & Update stability timeout for Azure CNI network policy ([#3895](https://github.com/Azure/aks-engine/issues/3895))
- E2E fix scale test indexing out of bounds ([#3873](https://github.com/Azure/aks-engine/issues/3873))
- not nil defense broke logic ([#3881](https://github.com/Azure/aks-engine/issues/3881))
- fixes for akse e2e tests ([#3874](https://github.com/Azure/aks-engine/issues/3874))
- E2E optionally validates rotate-certs ([#3866](https://github.com/Azure/aks-engine/issues/3866))
- add more E2E timeout tolerance ([#3850](https://github.com/Azure/aks-engine/issues/3850))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.56.1"></a>
# [v0.56.1] - 2021-02-03
### Bug Fixes 🐞
- Azure Stack CNI network interfaces file creation fix ([#3792](https://github.com/Azure/aks-engine/issues/3792))

### Maintenance 🔧
- add new Azure VM SKUs, brazilsoutheast, westus3 regions ([#4224](https://github.com/Azure/aks-engine/issues/4224))
- updating Windows VHD with new cached artifacts ([#3843](https://github.com/Azure/aks-engine/issues/3843))
- adding 1.19.1 bits to Windows VHD ([#3834](https://github.com/Azure/aks-engine/issues/3834))
- rev Linux VHDs to 2020.09.14 ([#3827](https://github.com/Azure/aks-engine/issues/3827))
- update signed powershell script package to include azure CNI fixes ([#3829](https://github.com/Azure/aks-engine/issues/3829))

### Revert Change ◀️
- "chore: targeting sept updates for Windows 2019 VHD ([#3801](https://github.com/Azure/aks-engine/issues/3801))" ([#3836](https://github.com/Azure/aks-engine/issues/3836))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.60.0"></a>
# [v0.60.0] - 2021-02-03
### Bug Fixes 🐞
- use curl.exe and add retries when downloading artifacts to Windows VHD ([#4206](https://github.com/Azure/aks-engine/issues/4206))
- update nodecustomdata to enable portmapper for kubenet+containerd ([#4191](https://github.com/Azure/aks-engine/issues/4191))
- fix network cleanup code on windows for contianerd nodes ([#4154](https://github.com/Azure/aks-engine/issues/4154))

### Continuous Integration 💜
- update dev image to deis/docker-go-dev v1.29.0 ([#4195](https://github.com/Azure/aks-engine/issues/4195))
- remove KMS test from everything cluster config ([#4194](https://github.com/Azure/aks-engine/issues/4194))

### Documentation 📘
- add dual-stack iptables api model ([#4185](https://github.com/Azure/aks-engine/issues/4185))
- aad not supported on AZs ([#4182](https://github.com/Azure/aks-engine/issues/4182))
- fix broken link in AAD docs ([#4180](https://github.com/Azure/aks-engine/issues/4180))

### Features 🌈
- run accelerated unattended-upgrade at node creation time ([#4217](https://github.com/Azure/aks-engine/issues/4217))
- reworked rotate-certs command ([#4214](https://github.com/Azure/aks-engine/issues/4214))
- add support for Kubernetes v1.16.15 and v1.18.15 on Azure Stack ([#4187](https://github.com/Azure/aks-engine/issues/4187))
- add support for Kubernetes v1.17.17 on Azure Stack ([#4188](https://github.com/Azure/aks-engine/issues/4188))
- add support for Kubernetes 1.20.2 ([#4192](https://github.com/Azure/aks-engine/issues/4192))
- create kms key as part of cluster bootstrap ([#4170](https://github.com/Azure/aks-engine/issues/4170))
- add support for Kubernetes v1.21.0-alpha.1 ([#4178](https://github.com/Azure/aks-engine/issues/4178))
- add support for Kubernetes v1.18.15 ([#4166](https://github.com/Azure/aks-engine/issues/4166))
- add support for Kubernetes v1.17.17 ([#4167](https://github.com/Azure/aks-engine/issues/4167))
- add support for Kubernetes v1.19.7 ([#4165](https://github.com/Azure/aks-engine/issues/4165))

### Maintenance 🔧
- add new Azure VM SKUs, brazilsoutheast, westus3 regions ([#4224](https://github.com/Azure/aks-engine/issues/4224))
- block container traffic to 168.63.129.16 ([#4212](https://github.com/Azure/aks-engine/issues/4212))
- rev Linux VHDs to 2021.01.28 ([#4223](https://github.com/Azure/aks-engine/issues/4223))
- update aks-engine VHD Windows VHD ([#4218](https://github.com/Azure/aks-engine/issues/4218))
- update csi-secrets-store to v0.0.19 and akv provider to 0.0.12 ([#4203](https://github.com/Azure/aks-engine/issues/4203))
- update policy addon deployment ([#4201](https://github.com/Azure/aks-engine/issues/4201))
- update adal to v0.9.10 ([#4200](https://github.com/Azure/aks-engine/issues/4200))
- updating windows provisioing scripts to v0.0.10 ([#4184](https://github.com/Azure/aks-engine/issues/4184))
- Update Azure CNI to v1.2.2 ([#4183](https://github.com/Azure/aks-engine/issues/4183))
- Add windowsnodelabelsync.ps1 ([#4163](https://github.com/Azure/aks-engine/issues/4163))
- install Jan 2021 security updates in Windows VHD ([#4168](https://github.com/Azure/aks-engine/issues/4168))
- bump kms keyvault to v0.0.10 ([#4169](https://github.com/Azure/aks-engine/issues/4169))
- Use signed scripts package 0.0.9 on Windows nodes ([#4162](https://github.com/Azure/aks-engine/issues/4162))

### Testing 💚
- incorporate kured + auto mode into kamino vmss-prototype tests ([#4221](https://github.com/Azure/aks-engine/issues/4221))
- add timestamps to E2E pod logs output ([#4216](https://github.com/Azure/aks-engine/issues/4216))
- enable kamino vmss-prototype dry run tests ([#4215](https://github.com/Azure/aks-engine/issues/4215))
- increase timeout tolerance for vmss-prototype SIG publishing ([#4211](https://github.com/Azure/aks-engine/issues/4211))
- add more stdout to kamino vmss-prototype E2E tests ([#4210](https://github.com/Azure/aks-engine/issues/4210))
- print helm command for vmss-prototype test ([#4207](https://github.com/Azure/aks-engine/issues/4207))
- correct E2E message during vmss-prototype test ([#4205](https://github.com/Azure/aks-engine/issues/4205))
- kamino E2E errata ([#4204](https://github.com/Azure/aks-engine/issues/4204))
- remove dangling fmt.Println ([#4202](https://github.com/Azure/aks-engine/issues/4202))
- enable testing custom kamino vmss prototype images ([#4198](https://github.com/Azure/aks-engine/issues/4198))
- updated kamino vmss-prototype integration E2E ([#4153](https://github.com/Azure/aks-engine/issues/4153))
- ensure hostPort routing ([#4186](https://github.com/Azure/aks-engine/issues/4186))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.59.0"></a>
# [v0.59.0] - 2021-01-13
### Bug Fixes 🐞
- update error code for vhd not found ([#4150](https://github.com/Azure/aks-engine/issues/4150))
- fix that kubeproxy cannot parse featuregates ([#4145](https://github.com/Azure/aks-engine/issues/4145))
- Fixing URl for azure-cni v1.20_hotfix package in Windows VHD ([#4136](https://github.com/Azure/aks-engine/issues/4136))
- ensure we wait for the right CRI service name ([#4115](https://github.com/Azure/aks-engine/issues/4115))
- honor user-defined feature gates before applying defaults ([#4113](https://github.com/Azure/aks-engine/issues/4113))
- typo in upgrade log message ([#4109](https://github.com/Azure/aks-engine/issues/4109))

### Continuous Integration 💜
- add manual dispatch trigger to nightly build ([#4123](https://github.com/Azure/aks-engine/issues/4123))

### Documentation 📘
- support FAQ ([#4148](https://github.com/Azure/aks-engine/issues/4148))
- remove warning about Windows + custom VNET ([#4147](https://github.com/Azure/aks-engine/issues/4147))
- Clarify README for Azure Stack Hub ([#4128](https://github.com/Azure/aks-engine/issues/4128))
- Added references to capz ([#4124](https://github.com/Azure/aks-engine/issues/4124))
- clarify msi as default, and requirements ([#4122](https://github.com/Azure/aks-engine/issues/4122))
- update dualstack docs to 1.20 ([#4101](https://github.com/Azure/aks-engine/issues/4101))

### Features 🌈
- setting a default containerd package for Windows ([#4149](https://github.com/Azure/aks-engine/issues/4149))
- add support for Kubernetes v1.18.14 ([#4139](https://github.com/Azure/aks-engine/issues/4139))
- add support for Kubernetes v1.19.6 ([#4140](https://github.com/Azure/aks-engine/issues/4140))
- add support for Kubernetes v1.17.16 ([#4138](https://github.com/Azure/aks-engine/issues/4138))
- add support for Kubernetes v1.20.1 ([#4141](https://github.com/Azure/aks-engine/issues/4141))
- Add WinDSR support ([#4104](https://github.com/Azure/aks-engine/issues/4104))
- Updating kubelet/kube-proxy to run with as high priority processes ([#4073](https://github.com/Azure/aks-engine/issues/4073))
- Add support for specifying linux moby url ([#4120](https://github.com/Azure/aks-engine/issues/4120))
- Updating Windows VHDs to include Dec 2012 patches ([#4118](https://github.com/Azure/aks-engine/issues/4118))
- add support for Kubernetes v1.18.13 ([#4111](https://github.com/Azure/aks-engine/issues/4111))
- add support for Kubernetes v1.19.5 ([#4110](https://github.com/Azure/aks-engine/issues/4110))
- add support for Kubernetes v1.20.0 ([#4102](https://github.com/Azure/aks-engine/issues/4102))

### Maintenance 🔧
- release Windows VHD for v0.59.0 aks-engine release ([#4158](https://github.com/Azure/aks-engine/issues/4158))
- rev Linux VHDs to 2021.01.08 ([#4159](https://github.com/Azure/aks-engine/issues/4159))
- simplify upgrade templates ([#4135](https://github.com/Azure/aks-engine/issues/4135))
- reinforce upgrade warnings/errors ([#4074](https://github.com/Azure/aks-engine/issues/4074))
- Use signed scripts v0.0.8 for Windows deployments ([#4134](https://github.com/Azure/aks-engine/issues/4134))
- simplify scale templates ([#4131](https://github.com/Azure/aks-engine/issues/4131))
- update Azure CNI to v1.2.0_hotfix ([#4129](https://github.com/Azure/aks-engine/issues/4129))
- update vendor deps ([#4125](https://github.com/Azure/aks-engine/issues/4125))
- rev metrics-server to v0.4.1 for all k8s versions ([#4127](https://github.com/Azure/aks-engine/issues/4127))
- Update moby/containerd versions ([#4119](https://github.com/Azure/aks-engine/issues/4119))
- update csi-secrets-store to v0.0.18 and akv provider to 0.0.11 ([#4126](https://github.com/Azure/aks-engine/issues/4126))
- updating azure-npm to 1.2.1 version ([#4094](https://github.com/Azure/aks-engine/issues/4094))
- Installing Dec 2020 cumulative updates for Windows VHDs ([#4103](https://github.com/Azure/aks-engine/issues/4103))
- update adal to v0.9.6 ([#4093](https://github.com/Azure/aks-engine/issues/4093))
- limit number of upgrade retries if new CP nodes bootstrap fails ([#4068](https://github.com/Azure/aks-engine/issues/4068))
- faster rolling updates for daemonset addons ([#4090](https://github.com/Azure/aks-engine/issues/4090))
- update csi-secrets-store addon manifest and images (v0.0.10) ([#4084](https://github.com/Azure/aks-engine/issues/4084))

### Testing 💚
- enable configurable node prototype tests ([#4100](https://github.com/Azure/aks-engine/issues/4100))
- enable GINKGO_SKIP_AFTER_UPGRADE param in e2e test ([#4089](https://github.com/Azure/aks-engine/issues/4089))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.58.0"></a>
# [v0.58.0] - 2020-12-03
### Bug Fixes 🐞
- filter out KeyVault resources during upgrade ([#4072](https://github.com/Azure/aks-engine/issues/4072))
- make system role assignments work w/ upgrade ([#4078](https://github.com/Azure/aks-engine/issues/4078))
- cleanup VMSSName + addpool ([#4067](https://github.com/Azure/aks-engine/issues/4067))
- enforce Azure CNI config via jq instead of sed ([#4060](https://github.com/Azure/aks-engine/issues/4060))
- persist VMSS name in API model ([#4051](https://github.com/Azure/aks-engine/issues/4051))
- Windows and VMSS version validation for azs ([#4030](https://github.com/Azure/aks-engine/issues/4030))
- always ensure etcd on control plane vms ([#4045](https://github.com/Azure/aks-engine/issues/4045))
- fixing an issue where windows cannot start kubelet to get podCIDR on Windows nodes when using kubenet ([#4039](https://github.com/Azure/aks-engine/issues/4039))
- Adding csi-proxy logs to prevent access denied errors ([#4029](https://github.com/Azure/aks-engine/issues/4029))
- upgrade broken for VMAS + managed identities ([#4021](https://github.com/Azure/aks-engine/issues/4021))
- enable dhcpv6 only for ubuntu 16.04 ([#4017](https://github.com/Azure/aks-engine/issues/4017))
- Re-ordering HNS policy removal due to 10c change in behavior and consolidating logic in Windows cleanup scripts  ([#4002](https://github.com/Azure/aks-engine/issues/4002))
- Mount /usr/local/share/ca-certificates folder into kube-controller-manager ([#4001](https://github.com/Azure/aks-engine/issues/4001))
- cancel Pod Eviction jobs when timeout ([#3970](https://github.com/Azure/aks-engine/issues/3970))
- ensure flatcar configs use transparent Azure CNI ([#3972](https://github.com/Azure/aks-engine/issues/3972))
- page not done logic bug ([#3971](https://github.com/Azure/aks-engine/issues/3971))
- resource identifier not passed to custom cloud profile ([#3953](https://github.com/Azure/aks-engine/issues/3953))
- get-logs command collects control plane logs even if apiserver is down ([#3939](https://github.com/Azure/aks-engine/issues/3939))
- use InternalIP for metrics-server DNS resolution ([#3929](https://github.com/Azure/aks-engine/issues/3929))
- VMSS control plane + availability zones ([#3917](https://github.com/Azure/aks-engine/issues/3917))
- Azure Stack upgrade operations clear old MCR ImageBase ([#3922](https://github.com/Azure/aks-engine/issues/3922))
- remove NSG rules for AzureStack ([#3914](https://github.com/Azure/aks-engine/issues/3914))

### Code Refactoring 💎
- remove unsupported orchestrators ([#3965](https://github.com/Azure/aks-engine/issues/3965))
- move ssh-related funcs to its own package ([#3990](https://github.com/Azure/aks-engine/issues/3990))
- delete Azure Stack's KubernetesClient implementation ([#3957](https://github.com/Azure/aks-engine/issues/3957))

### Continuous Integration 💜
- use 3 replicas for no-vnet E2E jobs ([#4061](https://github.com/Azure/aks-engine/issues/4061))
- Updating pr-windows-signed-scripts.yaml to allow for overriding cluster definition ([#4015](https://github.com/Azure/aks-engine/issues/4015))
- remove node-problem-detector addon from PR E2E ([#4010](https://github.com/Azure/aks-engine/issues/4010))
- Adding CI pipeline to verify windows provisioning script content BEFORE it gets check in ([#4008](https://github.com/Azure/aks-engine/issues/4008))
- Windows VHD pipeline tests ([#3977](https://github.com/Azure/aks-engine/issues/3977))
- reduce addons area for containerd config tests ([#3923](https://github.com/Azure/aks-engine/issues/3923))
- adding github actions to create nightly builds ([#3904](https://github.com/Azure/aks-engine/issues/3904))
- E2E fixes to allow more containerd tests ([#3903](https://github.com/Azure/aks-engine/issues/3903))

### Documentation 📘
- fixed broken link in Azure Stack topic page ([#4056](https://github.com/Azure/aks-engine/issues/4056))
- remove double quotes(") at the aks-deploy sample ([#3967](https://github.com/Azure/aks-engine/issues/3967))
- adding instructions for how how to build the Windows VHD for di… ([#3911](https://github.com/Azure/aks-engine/issues/3911))

### Features 🌈
- disable livenessProbe timeout enforcement ([#4085](https://github.com/Azure/aks-engine/issues/4085))
- add support for Kubernetes v1.20.0-rc.0 ([#4076](https://github.com/Azure/aks-engine/issues/4076))
- Enable chrony and host-based time sync by default on Ubuntu 18.04 ([#4011](https://github.com/Azure/aks-engine/issues/4011))
- add support for Kubernetes v1.19.4 ([#4042](https://github.com/Azure/aks-engine/issues/4042))
- add support for Kubernetes 1.18.12 ([#4043](https://github.com/Azure/aks-engine/issues/4043))
- add support for Kubernetes v1.20.0-beta.1 ([#3999](https://github.com/Azure/aks-engine/issues/3999))
- set apiserver tokenRequest flags in 1.20+ ([#3989](https://github.com/Azure/aks-engine/issues/3989))
- add support for Kubernetes v1.20.0-beta.0 ([#3991](https://github.com/Azure/aks-engine/issues/3991))
- add support for Kubernetes v1.20.0-alpha.3 ([#3934](https://github.com/Azure/aks-engine/issues/3934))
- upload collected logs to storage account container ([#3944](https://github.com/Azure/aks-engine/issues/3944))
- add support for Kubernetes v1.17.13 ([#3954](https://github.com/Azure/aks-engine/issues/3954))
- add support for Kubernetes v1.18.10 on Azure Stack ([#3950](https://github.com/Azure/aks-engine/issues/3950))
- add support for Kubernetes v1.18.10 ([#3948](https://github.com/Azure/aks-engine/issues/3948))
- updates for container monitoring addon omsagent agent September 2020 release ([#3942](https://github.com/Azure/aks-engine/issues/3942))
- add support for Kubernetes v1.19.3 ([#3937](https://github.com/Azure/aks-engine/issues/3937))
- custom Windows log collection script ([#3940](https://github.com/Azure/aks-engine/issues/3940))
- enable system-assigned identity by default ([#3856](https://github.com/Azure/aks-engine/issues/3856))
- Target new windows VHD with new K8s binaries (1.19.2, 1.18.9, etc) ([#3905](https://github.com/Azure/aks-engine/issues/3905))

### Maintenance 🔧
- update Windows image ([#4086](https://github.com/Azure/aks-engine/issues/4086))
- rev Linux VHDs to 2020.12.02 ([#4081](https://github.com/Azure/aks-engine/issues/4081))
- k8s v1.18 conformance model for Azure Stack ([#4070](https://github.com/Azure/aks-engine/issues/4070))
- validate VHD availability before upgrade/scale on Azure Stack Hub ([#4062](https://github.com/Azure/aks-engine/issues/4062))
- mark Kubernetes 1.17.14 as disabled ([#4044](https://github.com/Azure/aks-engine/issues/4044))
- cleanup dead code ([#4053](https://github.com/Azure/aks-engine/issues/4053))
- only warn about master stuff during create ([#4057](https://github.com/Azure/aks-engine/issues/4057))
- Upgrade CNI to v1.2.0 ([#4058](https://github.com/Azure/aks-engine/issues/4058))
- Updating default csi-proxy version to v0.2.2 ([#4047](https://github.com/Azure/aks-engine/issues/4047))
- format some json in addpool.md ([#4049](https://github.com/Azure/aks-engine/issues/4049))
- remove deprecated AKS code paths ([#4040](https://github.com/Azure/aks-engine/issues/4040))
- deprecate orchestratorType ([#4038](https://github.com/Azure/aks-engine/issues/4038))
- remove deprecated localizations ([#4036](https://github.com/Azure/aks-engine/issues/4036))
- Use Windows November(11b) updates as default vhd ([#4033](https://github.com/Azure/aks-engine/issues/4033))
- Windows November Patches ([#4023](https://github.com/Azure/aks-engine/issues/4023))
- Add 10c image as default ([#4020](https://github.com/Azure/aks-engine/issues/4020))
- rev Linux VHDs to 2020.10.30 ([#4016](https://github.com/Azure/aks-engine/issues/4016))
- Install Windows Server 2019 10C updates in Windows VHD ([#3956](https://github.com/Azure/aks-engine/issues/3956))
- update powershell signed scripts ([#4012](https://github.com/Azure/aks-engine/issues/4012))
- Adding csi-proxy-v0.2.2 to Windows VHD ([#4004](https://github.com/Azure/aks-engine/issues/4004))
- update cluster-autoscalers to latest patches ([#4000](https://github.com/Azure/aks-engine/issues/4000))
- set go sdk log level for Azure Stack clusters ([#3993](https://github.com/Azure/aks-engine/issues/3993))
- update cluster-autoscaler to v1.19.1 ([#3996](https://github.com/Azure/aks-engine/issues/3996))
- update go toolchain to v1.15.3 ([#3879](https://github.com/Azure/aks-engine/issues/3879))
- use transparent mode for Azure CNI ([#3958](https://github.com/Azure/aks-engine/issues/3958))
- gofmt ([#3982](https://github.com/Azure/aks-engine/issues/3982))
- longer default cordonDrainTimeout for Azure Stack Cloud ([#3969](https://github.com/Azure/aks-engine/issues/3969))
- kubelet systemd job depends on CRI service ([#3943](https://github.com/Azure/aks-engine/issues/3943))
- Upgrade CNI to v1.1.8 ([#3907](https://github.com/Azure/aks-engine/issues/3907))
- remove support for Kubernetes 1.15 ([#3751](https://github.com/Azure/aks-engine/issues/3751))
- bump calico to 3.8.9 to get latest patches ([#3924](https://github.com/Azure/aks-engine/issues/3924))
- set the csi-secrets-store to have a priority class ([#3909](https://github.com/Azure/aks-engine/issues/3909))
- rev Linux VHDs to 2020.10.06 ([#3906](https://github.com/Azure/aks-engine/issues/3906))

### Testing 💚
- more resilient azure-arc-onboarding E2E ([#4069](https://github.com/Azure/aks-engine/issues/4069))
- use 3 control plane VMs for base test config ([#4075](https://github.com/Azure/aks-engine/issues/4075))
- dns-liveness livenessProbe tweaks ([#4080](https://github.com/Azure/aks-engine/issues/4080))
- disable useManagedIdentity default value on Azure Stack ([#4063](https://github.com/Azure/aks-engine/issues/4063))
- skip rotate docker log tests if omsagent ([#4059](https://github.com/Azure/aks-engine/issues/4059))
- only test flannel + containerd thru k8s 1.19 ([#4041](https://github.com/Azure/aks-engine/issues/4041))
- re-engage cluster-autoscaler tests if no ssh ([#4034](https://github.com/Azure/aks-engine/issues/4034))
- add dns-loop (actually curl) stress test ([#4024](https://github.com/Azure/aks-engine/issues/4024))
- configurable stability iterations timeout ([#4022](https://github.com/Azure/aks-engine/issues/4022))
- fix two UTs to work with "cli" auth ([#3995](https://github.com/Azure/aks-engine/issues/3995))
- ensure azure-kms-provider E2E coverage ([#3985](https://github.com/Azure/aks-engine/issues/3985))
- "useManagedIdentity": false for VMSS masters ([#3981](https://github.com/Azure/aks-engine/issues/3981))
- Retry label when cluster first comes online ([#3976](https://github.com/Azure/aks-engine/issues/3976))
- Jenkinsfile bool vars ([#3980](https://github.com/Azure/aks-engine/issues/3980))
- Windows Metrics failure when High CPU  ([#3962](https://github.com/Azure/aks-engine/issues/3962))
- validate azure arc in everything config ([#3961](https://github.com/Azure/aks-engine/issues/3961))
- CCM + azurefile + MSI test scenarios ([#3932](https://github.com/Azure/aks-engine/issues/3932))
- skip tests for flatcar cluster config ([#3921](https://github.com/Azure/aks-engine/issues/3921))
- Add azure.json path for custom cloud k8s config & Update stability timeout for Azure CNI network policy ([#3895](https://github.com/Azure/aks-engine/issues/3895))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.57.0"></a>
# [v0.57.0] - 2020-10-08
### Bug Fixes 🐞
- fix jumpbox custom data ([#3896](https://github.com/Azure/aks-engine/issues/3896))
- don't warn about CRI update on cluster create ([#3890](https://github.com/Azure/aks-engine/issues/3890))
- ensure addon hostNetwork ports don't conflict ([#3894](https://github.com/Azure/aks-engine/issues/3894))
- de-dup vnet roleAssignment deployments ([#3857](https://github.com/Azure/aks-engine/issues/3857))
- add #EOF sentinel chars to custom search file ([#3862](https://github.com/Azure/aks-engine/issues/3862))
- E2E scale scenario broken if update is false ([#3852](https://github.com/Azure/aks-engine/issues/3852))
- Azure Stack CNI network interfaces file creation fix ([#3792](https://github.com/Azure/aks-engine/issues/3792))

### Continuous Integration 💜
- add containerd PR E2E tests ([#3892](https://github.com/Azure/aks-engine/issues/3892))
- update OWNERS ([#3853](https://github.com/Azure/aks-engine/issues/3853))

### Documentation 📘
- confirm flannel + docker is not supported ([#3886](https://github.com/Azure/aks-engine/issues/3886))
- add notes about not upgrading LB config ([#3884](https://github.com/Azure/aks-engine/issues/3884))
- clarify that update is for node pools only ([#3877](https://github.com/Azure/aks-engine/issues/3877))
- CLI operations ([#3837](https://github.com/Azure/aks-engine/issues/3837))
- mark rotate-certs command as experimental ([#3869](https://github.com/Azure/aks-engine/issues/3869))
- add v0.55.4 to Azure Stack topic page ([#3846](https://github.com/Azure/aks-engine/issues/3846))
- remove mention about library for AKS ([#3835](https://github.com/Azure/aks-engine/issues/3835))
- add required attribution reminder to PR template ([#3826](https://github.com/Azure/aks-engine/issues/3826))

### Features 🌈
- Target new windows VHD with new K8s binaries (1.19.2, 1.18.9, etc) ([#3905](https://github.com/Azure/aks-engine/issues/3905))
- add ScaleCPULimitsToSandbox for hyperv runtimeclasses ([#3889](https://github.com/Azure/aks-engine/issues/3889))
- allow custom containerd package for Linux nodes ([#3878](https://github.com/Azure/aks-engine/issues/3878))
- Updating Windows VHD build files to support building for multiple OS versions ([#3847](https://github.com/Azure/aks-engine/issues/3847))
- add support for Kubernetes v1.18.9 ([#3841](https://github.com/Azure/aks-engine/issues/3841))
- add support for Kubernetes v1.19.2 ([#3842](https://github.com/Azure/aks-engine/issues/3842))
- add support for Kubernetes v1.17.12 ([#3840](https://github.com/Azure/aks-engine/issues/3840))
- update VMSS node pools ([#3830](https://github.com/Azure/aks-engine/issues/3830))

### Maintenance 🔧
- rev Linux VHDs to 2020.10.06 ([#3906](https://github.com/Azure/aks-engine/issues/3906))
- enable VHD re-use in no outbound scenarios ([#3897](https://github.com/Azure/aks-engine/issues/3897))
- update csi-secrets-store addon manifest and images (v0.0.9) ([#3891](https://github.com/Azure/aks-engine/issues/3891))
- create azure.json via CSE ([#3876](https://github.com/Azure/aks-engine/issues/3876))
- distribute apiserver.crt to control plane nodes only ([#3860](https://github.com/Azure/aks-engine/issues/3860))
- update azure cni to 1.1.7 ([#3864](https://github.com/Azure/aks-engine/issues/3864))
- update Dashboard addon to v2.0.4 ([#3855](https://github.com/Azure/aks-engine/issues/3855))
- check VHD media name length before being published to Marketplace ([#3799](https://github.com/Azure/aks-engine/issues/3799))
- remove no-op 1.15 version checks in templates ([#3851](https://github.com/Azure/aks-engine/issues/3851))
- updating Windows VHD with new cached artifacts ([#3843](https://github.com/Azure/aks-engine/issues/3843))
- adding 1.19.1 bits to Windows VHD ([#3834](https://github.com/Azure/aks-engine/issues/3834))
- rev Linux VHDs to 2020.09.14 ([#3827](https://github.com/Azure/aks-engine/issues/3827))
- update signed powershell script package to include azure CNI fixes ([#3829](https://github.com/Azure/aks-engine/issues/3829))

### Revert Change ◀️
- "chore: targeting sept updates for Windows 2019 VHD ([#3801](https://github.com/Azure/aks-engine/issues/3801))" ([#3836](https://github.com/Azure/aks-engine/issues/3836))

### Testing 💚
- Add azure.json path for custom cloud k8s config & Update stability timeout for Azure CNI network policy ([#3895](https://github.com/Azure/aks-engine/issues/3895))
- E2E fix scale test indexing out of bounds ([#3873](https://github.com/Azure/aks-engine/issues/3873))
- not nil defense broke logic ([#3881](https://github.com/Azure/aks-engine/issues/3881))
- fixes for akse e2e tests ([#3874](https://github.com/Azure/aks-engine/issues/3874))
- E2E optionally validates rotate-certs ([#3866](https://github.com/Azure/aks-engine/issues/3866))
- add more E2E timeout tolerance ([#3850](https://github.com/Azure/aks-engine/issues/3850))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.56.0"></a>
# [v0.56.0] - 2020-09-17
### Bug Fixes 🐞
- Azure Stack CNI network interfaces file creation fix ([#3792](https://github.com/Azure/aks-engine/issues/3792))
- clean up stale CNI data regardless of HNS state ([#3822](https://github.com/Azure/aks-engine/issues/3822))
- working systemd monitor jobs ([#3788](https://github.com/Azure/aks-engine/issues/3788))
- set correct ssh private key for e2e get-logs ([#3824](https://github.com/Azure/aks-engine/issues/3824))
- check if windows profile exists in upgrade ([#3820](https://github.com/Azure/aks-engine/issues/3820))
- update GPU driver to 450.51.06 ([#3814](https://github.com/Azure/aks-engine/issues/3814)) ([#3821](https://github.com/Azure/aks-engine/issues/3821))
- more resilient etcd systemd ([#3809](https://github.com/Azure/aks-engine/issues/3809))
- updating pub to v0.2.6 ([#3802](https://github.com/Azure/aks-engine/issues/3802))
- Enable Windows VHD version update in aks-engine upgrade scenario ([#3774](https://github.com/Azure/aks-engine/issues/3774))
- mount /etc/ssl/certs for apiserver ([#3800](https://github.com/Azure/aks-engine/issues/3800))
- mount certs directory in apiserver pod ([#3782](https://github.com/Azure/aks-engine/issues/3782))
- update the flag due to using higher version of csi livenessprobe ([#3770](https://github.com/Azure/aks-engine/issues/3770))
- availabilityZones value in template to read from parameter ([#3767](https://github.com/Azure/aks-engine/issues/3767))
- validate k8s control plane before addons ([#3729](https://github.com/Azure/aks-engine/issues/3729))
- ensure cleanup scripts can login w/ sp ([#3765](https://github.com/Azure/aks-engine/issues/3765))
- Restart=always docker systemd service ([#3758](https://github.com/Azure/aks-engine/issues/3758))
- use cached v1.15-azs kube-proxy on Azure Stack ([#3757](https://github.com/Azure/aks-engine/issues/3757))
- fix issue where installing a different version of containerd vs what was pre-installed on VHD failed silently ([#3743](https://github.com/Azure/aks-engine/issues/3743))
- add missing 1.16.14 kube-proxy image to the vhd dependencies ([#3736](https://github.com/Azure/aks-engine/issues/3736))

### Continuous Integration 💜
- use 1.4.0 containerd release for windows in CI ([#3742](https://github.com/Azure/aks-engine/issues/3742))

### Documentation 📘
- clarify project support policy in SUPPORT.md ([#3813](https://github.com/Azure/aks-engine/issues/3813))
- add doc and example api model for IPv6 ([#3777](https://github.com/Azure/aks-engine/issues/3777))

### Features 🌈
- azure kms provider as static pod ([#3667](https://github.com/Azure/aks-engine/issues/3667))
- add support for Kubernetes 1.19.1 ([#3816](https://github.com/Azure/aks-engine/issues/3816))
- secrets store addon can pull from custom env ([#3787](https://github.com/Azure/aks-engine/issues/3787))
- add support for Kubernetes v1.16.15 ([#3780](https://github.com/Azure/aks-engine/issues/3780))
- add support for Kubernetes v1.19.0 ([#3754](https://github.com/Azure/aks-engine/issues/3754))
- variable upgrade timeout based on num nodes ([#3752](https://github.com/Azure/aks-engine/issues/3752))
- collect hyperv logs ([#3737](https://github.com/Azure/aks-engine/issues/3737))

### Maintenance 🔧
- updating Windows VHD with new cached artifacts ([#3843](https://github.com/Azure/aks-engine/issues/3843))
- adding 1.19.1 bits to Windows VHD ([#3834](https://github.com/Azure/aks-engine/issues/3834))
- rev Linux VHDs to 2020.09.14 ([#3827](https://github.com/Azure/aks-engine/issues/3827))
- update signed powershell script package to include azure CNI fixes ([#3829](https://github.com/Azure/aks-engine/issues/3829))
- pin virtual-kubelet to version 1.2.1.2 ([#3815](https://github.com/Azure/aks-engine/issues/3815))
- reduce customData payload ([#3793](https://github.com/Azure/aks-engine/issues/3793))
- targeting sept updates for Windows 2019 VHD ([#3801](https://github.com/Azure/aks-engine/issues/3801))
- remove apiserver /etc/kubernetes/certs mount ([#3808](https://github.com/Azure/aks-engine/issues/3808))
- update Linux VHDs to 2020.09.08 ([#3811](https://github.com/Azure/aks-engine/issues/3811))
- get calico and v-k from mcr.microsoft.com ([#3803](https://github.com/Azure/aks-engine/issues/3803))
- add keyVaultEndpoint to Azure Stack environment ([#3790](https://github.com/Azure/aks-engine/issues/3790))
- update node-problem-detector addon to v0.8.4 ([#3779](https://github.com/Azure/aks-engine/issues/3779))
- simplify addons config for calico and flannel ([#3773](https://github.com/Azure/aks-engine/issues/3773))
- update CNI plugins to v0.8.7 ([#3771](https://github.com/Azure/aks-engine/issues/3771))
- Update NPM to latest version 1.1.7 ([#3740](https://github.com/Azure/aks-engine/issues/3740))
- always install moby-runc ([#3763](https://github.com/Azure/aks-engine/issues/3763))
- specify containerd version to install with moby-engine package ([#3723](https://github.com/Azure/aks-engine/issues/3723))
- rev Linux VHDs to 2020.08.24 ([#3750](https://github.com/Azure/aks-engine/issues/3750))
- add new Azure VM SKUs ([#3744](https://github.com/Azure/aks-engine/issues/3744))
- update windows default VHD for August ([#3730](https://github.com/Azure/aks-engine/issues/3730))
- update csi-secrets-store addon manifest and images ([#3728](https://github.com/Azure/aks-engine/issues/3728))

### Revert Change ◀️
- "chore: targeting sept updates for Windows 2019 VHD ([#3801](https://github.com/Azure/aks-engine/issues/3801))" ([#3836](https://github.com/Azure/aks-engine/issues/3836))

### Testing 💚
- add private key input to e2e suite + keep all junit result files ([#3747](https://github.com/Azure/aks-engine/issues/3747))
- remove tiller addon from PR E2E cluster configuration ([#3794](https://github.com/Azure/aks-engine/issues/3794))
- don't try to kill docker if using containerd ([#3764](https://github.com/Azure/aks-engine/issues/3764))
- add REBOOT_CONTROL_PLANE_NODES E2E config ([#3745](https://github.com/Azure/aks-engine/issues/3745))
- validate kubelet and docker systemd during E2E ([#3759](https://github.com/Azure/aks-engine/issues/3759))
- disable azure-arc-onboarding addon in everything cluster config ([#3756](https://github.com/Azure/aks-engine/issues/3756))
- update kubernetes e2e to use GINKGO_FAIL_FAST parameter value ([#3660](https://github.com/Azure/aks-engine/issues/3660))
- fix ginkgo failFast ([#3738](https://github.com/Azure/aks-engine/issues/3738))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.55.4"></a>
# [v0.55.4] - 2020-09-15
### Bug Fixes 🐞
- update Azure Stack's Linux VHD to 2020.09.14 ([#3828](https://github.com/Azure/aks-engine/issues/3828))

### Maintenance 🔧
- validate Azure Stack's Linux VHD is in PIR ([#3831](https://github.com/Azure/aks-engine/issues/3831))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.55.3"></a>
# [v0.55.3] - 2020-09-08
### Bug Fixes 🐞
- mount /etc/ssl/certs for apiserver ([#3800](https://github.com/Azure/aks-engine/issues/3800))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.55.2"></a>
# [v0.55.2] - 2020-09-03
### Bug Fixes 🐞
- mount certs directory in apiserver pod ([#3782](https://github.com/Azure/aks-engine/issues/3782))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.55.1"></a>
# [v0.55.1] - 2020-09-02
### Bug Fixes 🐞
- update the flag due to using higher version of csi livenessprobe ([#3770](https://github.com/Azure/aks-engine/issues/3770))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.55.0"></a>
# [v0.55.0] - 2020-08-27
### Bug Fixes 🐞
- use cached v1.15-azs kube-proxy on Azure Stack ([#3757](https://github.com/Azure/aks-engine/issues/3757))
- fix issue where installing a different version of containerd vs what was pre-installed on VHD failed silently ([#3743](https://github.com/Azure/aks-engine/issues/3743))
- add missing 1.16.14 kube-proxy image to the vhd dependencies ([#3736](https://github.com/Azure/aks-engine/issues/3736))
- do not restart addon-manager if a reboot is required ([#3721](https://github.com/Azure/aks-engine/issues/3721))
- add rbac rules for nodes/status in cloud-node-manager.yaml ([#3699](https://github.com/Azure/aks-engine/issues/3699))
- use new Windows VHD image version ([#3701](https://github.com/Azure/aks-engine/issues/3701))
- don't wait for pod-security-policy spec if disabled ([#3673](https://github.com/Azure/aks-engine/issues/3673))
- metrics-server upgrade from v1.15 to v1.16 ([#3691](https://github.com/Azure/aks-engine/issues/3691))
- validate "basic" LB and fix related unit tests ([#3690](https://github.com/Azure/aks-engine/issues/3690))
- only remove CRI if re-installing ([#3654](https://github.com/Azure/aks-engine/issues/3654))
- update windows docker version to fix log rotation error ([#3641](https://github.com/Azure/aks-engine/issues/3641))
- reinforce MCR migration during upgrade for older clusters ([#3625](https://github.com/Azure/aks-engine/issues/3625))
- updating hotfix packages used for windows inlcuding CVE and metric ([#3599](https://github.com/Azure/aks-engine/issues/3599))
- use correct VMSS name comparison during Windows scale ([#3590](https://github.com/Azure/aks-engine/issues/3590))
- always extract kubelet and kubectl from a custom URL if defined ([#3574](https://github.com/Azure/aks-engine/issues/3574))
- use mcr.azk8s.cn for Azure CNI in China cloud ([#3587](https://github.com/Azure/aks-engine/issues/3587))
- update akse upgrade and scale in the e2e test to include identity ([#3582](https://github.com/Azure/aks-engine/issues/3582))
- validate signed powershell packages in Azure and China Cloud ([#3579](https://github.com/Azure/aks-engine/issues/3579))
- kube-proxy upgrade from v1.15 to v1.16 ([#3570](https://github.com/Azure/aks-engine/issues/3570))
- change addons image pull policy to IfNotPresent ([#3562](https://github.com/Azure/aks-engine/issues/3562))
- remove v1.17 support on Azure Stack ([#3544](https://github.com/Azure/aks-engine/issues/3544))
- downgrade blobfuse to 1.1.1 since there is breaking change on 1.2.3 ([#3540](https://github.com/Azure/aks-engine/issues/3540))
- always use /var/run/reboot-required ([#3536](https://github.com/Azure/aks-engine/issues/3536))
- updating hotfix packages used for windows to include hyperkube a… ([#3539](https://github.com/Azure/aks-engine/issues/3539))
- use kube-proxy image instead of hyperkube for 1.16 ([#3532](https://github.com/Azure/aks-engine/issues/3532))
- add windows hotfix overrides for 1.18.2, 1.16.10, 1.15.12 for AKS ([#3528](https://github.com/Azure/aks-engine/issues/3528))
- add validation for enableVMSSNodePublicIP + Basic LB ([#3524](https://github.com/Azure/aks-engine/issues/3524))
- ensure cron systemd job restarts ([#3523](https://github.com/Azure/aks-engine/issues/3523))
- use hotfix Windows builds for 1.18.4, 1.17.7, and 1.16.11 ([#3510](https://github.com/Azure/aks-engine/issues/3510))
- addons bootstrapping errata ([#3481](https://github.com/Azure/aks-engine/issues/3481))
- scale + AvailabilitySet errata ([#3490](https://github.com/Azure/aks-engine/issues/3490))
- make security context configs more restrictive ([#3454](https://github.com/Azure/aks-engine/issues/3454))
- validate that --upgrade-version is semver-compatible during upgrade ([#3442](https://github.com/Azure/aks-engine/issues/3442))
- Use mirror in China for AzureChinaCloud ([#3433](https://github.com/Azure/aks-engine/issues/3433))
- VolumePluginDir issue on Windows ([#3426](https://github.com/Azure/aks-engine/issues/3426))
- disable apiserver insecure-port ([#3422](https://github.com/Azure/aks-engine/issues/3422))
- only schedule mic on Linux nodes ([#3420](https://github.com/Azure/aks-engine/issues/3420))
- Enable e2e to install ginko if required ([#3330](https://github.com/Azure/aks-engine/issues/3330))
- bypass AzurePublicCloud assumption in e2e suite ([#3371](https://github.com/Azure/aks-engine/issues/3371))
- use nvidia drivers version 418.126.02 for 18.04 ([#3366](https://github.com/Azure/aks-engine/issues/3366))
- use right nodeSelector for cluster-autoscaler addon ([#3350](https://github.com/Azure/aks-engine/issues/3350))
- use ctr to remove containerd containers ([#3144](https://github.com/Azure/aks-engine/issues/3144))
- ensure pod-security-policy is the first addon loaded ([#3313](https://github.com/Azure/aks-engine/issues/3313))
- prevent panic in Windows scale if no tags ([#3308](https://github.com/Azure/aks-engine/issues/3308))
- ensure ERR_K8S_API_SERVER_DNS_LOOKUP_FAIL is returned when DNS check fails ([#3303](https://github.com/Azure/aks-engine/issues/3303))
- addon-manager definition should be under "labels" ([#3302](https://github.com/Azure/aks-engine/issues/3302))
- whitespace trimming in cloud-provider spec is broken ([#3301](https://github.com/Azure/aks-engine/issues/3301))
- increase apiserver connection timeout for private apiserver ([#3293](https://github.com/Azure/aks-engine/issues/3293))
- fix issue provisioning windows nodes with kubenet ([#3300](https://github.com/Azure/aks-engine/issues/3300))
- ensure /mnt is ext4 ([#3297](https://github.com/Azure/aks-engine/issues/3297))
- resilient etcd data disk mounting ([#3259](https://github.com/Azure/aks-engine/issues/3259))
- USER_ASSIGNED_IDENTITY_ID is empty in azure.json ([#3254](https://github.com/Azure/aks-engine/issues/3254))
- update Windows image validate on Azure Stack Hub ([#3260](https://github.com/Azure/aks-engine/issues/3260))
- don't dry-run cluster-init spec ([#3247](https://github.com/Azure/aks-engine/issues/3247))
- custom cloud Azure CNI config for Windows ([#3228](https://github.com/Azure/aks-engine/issues/3228))
- Add Windows node pool to containerd tests ([#3238](https://github.com/Azure/aks-engine/issues/3238))
- USER_ASSIGNED_IDENTITY_ID is empty in azure.json ([#3232](https://github.com/Azure/aks-engine/issues/3232))
- update Azure Stack Windows binaries component name ([#3231](https://github.com/Azure/aks-engine/issues/3231))
- don't validate Windows version in update scenario ([#3226](https://github.com/Azure/aks-engine/issues/3226))
- set Windows defender process exclustion for containerd ([#3215](https://github.com/Azure/aks-engine/issues/3215))
- aad-pod-identity taints foo works in k8s < 1.16 ([#3199](https://github.com/Azure/aks-engine/issues/3199))
- update dcap checksum file ([#3172](https://github.com/Azure/aks-engine/issues/3172))
- added resource limits for networkmonitor daemonset ([#2624](https://github.com/Azure/aks-engine/issues/2624))
- don't hardcode pause image for containerD on Windows ([#3158](https://github.com/Azure/aks-engine/issues/3158))
- Expose error details in Windows CSE ([#3159](https://github.com/Azure/aks-engine/issues/3159))
- correct destination filename for scheduled maintenance addon ([#3142](https://github.com/Azure/aks-engine/issues/3142))
- Correct wrong file name in the logs ([#3138](https://github.com/Azure/aks-engine/issues/3138))
- don't hardcode csi enableproxy in kubeclusterconfig.json ([#3127](https://github.com/Azure/aks-engine/issues/3127))

### Build 🏭
- fetch k8s node binaries from kubernetesartifacts storage ([#3151](https://github.com/Azure/aks-engine/issues/3151))

### Code Refactoring 💎
- consolidate tiller addon spec ([#3399](https://github.com/Azure/aks-engine/issues/3399))
- consolidate smb-flexvolume addon spec ([#3397](https://github.com/Azure/aks-engine/issues/3397))
- consolidate kube-proxy addon spec ([#3391](https://github.com/Azure/aks-engine/issues/3391))
- consolidate kube-dns addon spec ([#3390](https://github.com/Azure/aks-engine/issues/3390))
- consolidate keyvault-flexvolume addon spec ([#3389](https://github.com/Azure/aks-engine/issues/3389))
- consolidate flannel addon spec ([#3385](https://github.com/Azure/aks-engine/issues/3385))
- drop "beta" from kubernetes.io/os label ([#3369](https://github.com/Azure/aks-engine/issues/3369))
- consolidate aci-connector addon spec ([#3358](https://github.com/Azure/aks-engine/issues/3358))
- consolidate blobfuse-flexvolume addon spec ([#3359](https://github.com/Azure/aks-engine/issues/3359))
- consolidate calico addon spec ([#3361](https://github.com/Azure/aks-engine/issues/3361))
- consolidate cloud-node-manager addon spec ([#3363](https://github.com/Azure/aks-engine/issues/3363))
- consolidate aad-pod-identity addon spec ([#3356](https://github.com/Azure/aks-engine/issues/3356))
- consolidate audit-policy "addon" spec ([#3355](https://github.com/Azure/aks-engine/issues/3355))
- consolidate ip-masq-agent addon spec ([#3347](https://github.com/Azure/aks-engine/issues/3347))
- consolidate kube-rescheduler addon spec ([#3323](https://github.com/Azure/aks-engine/issues/3323))
- consolidate pod-security-policy addon spec ([#3284](https://github.com/Azure/aks-engine/issues/3284))
- consolidate cilium addons specs ([#3289](https://github.com/Azure/aks-engine/issues/3289))
- consolidate azure-network-policy addon spec ([#3283](https://github.com/Azure/aks-engine/issues/3283))
- consolidate azure-cloud-provider addon spec ([#3277](https://github.com/Azure/aks-engine/issues/3277))
- move etcd data disk mounting to cloud init  ([#3273](https://github.com/Azure/aks-engine/issues/3273))
- consolidate nvidia-device-plugin addon spec ([#3235](https://github.com/Azure/aks-engine/issues/3235))
- consolidate metrics-server addon spec ([#3211](https://github.com/Azure/aks-engine/issues/3211))
- use newer "az deployment group" CLI syntax ([#3213](https://github.com/Azure/aks-engine/issues/3213))
- consolidate cluster-autoscaler addon spec ([#3198](https://github.com/Azure/aks-engine/issues/3198))
- consolidate azure-cni-networkmonitor addon spec ([#3189](https://github.com/Azure/aks-engine/issues/3189))
- use simpler file names for addons maintenance ([#3187](https://github.com/Azure/aks-engine/issues/3187))
- simplify golangci-lint invocation ([#3182](https://github.com/Azure/aks-engine/issues/3182))
- simplify kubelet systemd service ([#3167](https://github.com/Azure/aks-engine/issues/3167))
- enable generic --register-with-taints and --register-node ([#3155](https://github.com/Azure/aks-engine/issues/3155))

### Continuous Integration 💜
- use 1.4.0 containerd release for windows in CI ([#3742](https://github.com/Azure/aks-engine/issues/3742))
- clarify apt packages installed during VHD CI ([#3674](https://github.com/Azure/aks-engine/issues/3674))
- updating pub to v0.2.3 ([#3398](https://github.com/Azure/aks-engine/issues/3398))
- add notice file for windows vhd ([#3345](https://github.com/Azure/aks-engine/issues/3345))
- use ginkgo 1.12, run go mod vendor during make coverage ([#3343](https://github.com/Azure/aks-engine/issues/3343))
- enable multiple static resource group exclusions in cleanup script ([#3340](https://github.com/Azure/aks-engine/issues/3340))
- moving windows vhd/image related test configs to westus2 ([#3320](https://github.com/Azure/aks-engine/issues/3320))
- passing container runtime var when executing new-windows-sku.sh in dev container ([#3242](https://github.com/Azure/aks-engine/issues/3242))
- ensure sku title/summary/longsummary is unique when creating new marketplace skus ([#3239](https://github.com/Azure/aks-engine/issues/3239))
- enable staticcheck linter ([#3191](https://github.com/Azure/aks-engine/issues/3191))
- disable artifact publishing ([#3163](https://github.com/Azure/aks-engine/issues/3163))

### Documentation 📘
- update limitations in dualstack docs ([#3727](https://github.com/Azure/aks-engine/issues/3727))
- adding James to OWNERS ([#3706](https://github.com/Azure/aks-engine/issues/3706))
- fix misleading sysctldConfig examples ([#3682](https://github.com/Azure/aks-engine/issues/3682))
- update the SGX DCAP version mentioned in the sgx instructions ([#3684](https://github.com/Azure/aks-engine/issues/3684))
- update stale Istio page and example ([#3639](https://github.com/Azure/aks-engine/issues/3639))
- fix incorrect calico addons section heading ([#3565](https://github.com/Azure/aks-engine/issues/3565))
- update throttling doc ([#3545](https://github.com/Azure/aks-engine/issues/3545))
- add Azure API throttling doc ([#3534](https://github.com/Azure/aks-engine/issues/3534))
- update azure-stack documentation ([#3479](https://github.com/Azure/aks-engine/issues/3479))
- supported versions for Azure Stacks using VHD 2020.05.13 ([#3457](https://github.com/Azure/aks-engine/issues/3457))
- remove references to "whitelist" ([#3428](https://github.com/Azure/aks-engine/issues/3428))
- update sgx doc and sgx-test tag ([#3349](https://github.com/Azure/aks-engine/issues/3349))
- aks-ubuntu-18.04 is default distro ([#3317](https://github.com/Azure/aks-engine/issues/3317))
- fix capitalization ([#3264](https://github.com/Azure/aks-engine/issues/3264))
- AKS-engine (on Stack) considerations around the use of proxy servers ([#3241](https://github.com/Azure/aks-engine/issues/3241))
- remove out of date Calico upgrade note ([#3243](https://github.com/Azure/aks-engine/issues/3243))
- update quickstart.md to add scoop install guide ([#3218](https://github.com/Azure/aks-engine/issues/3218))
- update sgx plugin docs ([#3230](https://github.com/Azure/aks-engine/issues/3230))
- rewrite monitoring topic ([#3205](https://github.com/Azure/aks-engine/issues/3205))
- add missing link to upgrade doc ([#3221](https://github.com/Azure/aks-engine/issues/3221))
- update sgx install instructions to dcv2 vms and newer resource name ([#3214](https://github.com/Azure/aks-engine/issues/3214))
- refresh documentation store ([#3177](https://github.com/Azure/aks-engine/issues/3177))
- remove obsolete api versions documentation ([#3175](https://github.com/Azure/aks-engine/issues/3175))
- document AKS Engine E2E ([#3131](https://github.com/Azure/aks-engine/issues/3131))
- include v0.48.0 in Azure Stack docs ([#3153](https://github.com/Azure/aks-engine/issues/3153))

### Features 🌈
- variable upgrade timeout based on num nodes ([#3752](https://github.com/Azure/aks-engine/issues/3752))
- collect hyperv logs ([#3737](https://github.com/Azure/aks-engine/issues/3737))
- Patch v1.15.11, v1.15.12, v1.16.10, v1.16.13, v1.17.7, v1.17.9, v1.18.4 v1.18.6 ([#3725](https://github.com/Azure/aks-engine/issues/3725))
- add support for K8s v1.17.11 on Azure Stack ([#3702](https://github.com/Azure/aks-engine/issues/3702))
- add support for K8s v1.16.14 on Azure Stack ([#3704](https://github.com/Azure/aks-engine/issues/3704))
- add support for Kubernetes v1.17.11 ([#3696](https://github.com/Azure/aks-engine/issues/3696))
- add support for Kubernetes v1.18.8 ([#3697](https://github.com/Azure/aks-engine/issues/3697))
- configurable microsoft apt repository ([#3698](https://github.com/Azure/aks-engine/issues/3698))
- add support for Kubernetes v1.16.14 ([#3695](https://github.com/Azure/aks-engine/issues/3695))
- cluster upgrade operations upgrade pause image ([#3689](https://github.com/Azure/aks-engine/issues/3689))
- add support for Kubernetes 1.19.0-rc.4 ([#3666](https://github.com/Azure/aks-engine/issues/3666))
- azure arc addon ([#3634](https://github.com/Azure/aks-engine/issues/3634))
- add support for Kubernetes 1.19.0-rc.3 ([#3653](https://github.com/Azure/aks-engine/issues/3653))
- add support for Kubernetes 1.16.13 ([#3602](https://github.com/Azure/aks-engine/issues/3602))
- Support configurable Windows PauseImage ([#3594](https://github.com/Azure/aks-engine/issues/3594))
- add support for Kubernetes 1.17.9 ([#3603](https://github.com/Azure/aks-engine/issues/3603))
- add support for Kubernetes 1.18.6 ([#3604](https://github.com/Azure/aks-engine/issues/3604))
- etcdStorageLimitGB for configurable etcd storage limit ([#3583](https://github.com/Azure/aks-engine/issues/3583))
- add hosts config agent support for Windows nodes ([#3572](https://github.com/Azure/aks-engine/issues/3572))
- Consume signed powershell scripts ([#3441](https://github.com/Azure/aks-engine/issues/3441))
- enable /etc/hosts config agent for AKS private cluster BYO DNS ([#3556](https://github.com/Azure/aks-engine/issues/3556))
- add support for Kubernetes v1.16.12 ([#3551](https://github.com/Azure/aks-engine/issues/3551))
- add support for Kubernetes v1.17.8 ([#3552](https://github.com/Azure/aks-engine/issues/3552))
- add support for Kubernetes v1.18.5 ([#3553](https://github.com/Azure/aks-engine/issues/3553))
- add support for Kubernetes 1.18.4 ([#3500](https://github.com/Azure/aks-engine/issues/3500))
- add support for Kubernetes 1.17.7 ([#3499](https://github.com/Azure/aks-engine/issues/3499))
- add support for Kubernetes 1.16.11 ([#3496](https://github.com/Azure/aks-engine/issues/3496))
- enable configurable usageReportingEnabled in calico addon ([#3493](https://github.com/Azure/aks-engine/issues/3493))
- ensure all AKS required ports on api-server are exposed ([#3488](https://github.com/Azure/aks-engine/issues/3488))
- enable windows daemonset in container monitoring addon ([#3466](https://github.com/Azure/aks-engine/issues/3466))
- add support for Kubernetes v1.19.0-beta.2 ([#3471](https://github.com/Azure/aks-engine/issues/3471))
- Dualstack support for Windows containers ([#3415](https://github.com/Azure/aks-engine/issues/3415))
- flatcar nodes ([#3380](https://github.com/Azure/aks-engine/issues/3380))
- enable configurable apiserver --anonymous-auth ([#3430](https://github.com/Azure/aks-engine/issues/3430))
- Update container monitoring addon for may omsagent release  ([#3403](https://github.com/Azure/aks-engine/issues/3403))
- add support for Kubernetes 1.19.0-beta.1 ([#3408](https://github.com/Azure/aks-engine/issues/3408))
- add "get-versions --azure-env" flag to list custom clouds supported versions ([#3394](https://github.com/Azure/aks-engine/issues/3394))
- configurable calico logging verbosity ([#3396](https://github.com/Azure/aks-engine/issues/3396))
- deprecate heapster addon ([#3387](https://github.com/Azure/aks-engine/issues/3387))
- add support for Kubernetes 1.18.3 ([#3309](https://github.com/Azure/aks-engine/issues/3309))
- add support for Kubernetes 1.16.10 ([#3312](https://github.com/Azure/aks-engine/issues/3312))
- add support for Kubernetes 1.17.6 ([#3311](https://github.com/Azure/aks-engine/issues/3311))
- add support for Kubernetes 1.19.0-beta.0 ([#3299](https://github.com/Azure/aks-engine/issues/3299))
- Updating Windows VHDs to include May patches ([#3263](https://github.com/Azure/aks-engine/issues/3263))
- enable alternate kube-reserved cgroups ([#3201](https://github.com/Azure/aks-engine/issues/3201))
- add support for Kubernetes 1.15.12 ([#3212](https://github.com/Azure/aks-engine/issues/3212))
- add support for Kubernetes 1.19.0-alpha.3 ([#3197](https://github.com/Azure/aks-engine/issues/3197))
- add warning message for empty Location string ([#3174](https://github.com/Azure/aks-engine/issues/3174))
- block imds from VMSS nodes if aad-pod-identity + msi ([#3136](https://github.com/Azure/aks-engine/issues/3136))
- Create Windows containerd VHDs ([#3162](https://github.com/Azure/aks-engine/issues/3162))
- configurable disk caching ([#2863](https://github.com/Azure/aks-engine/issues/2863))
- Kubernetes Dashboard addon v2.0.0 ([#3140](https://github.com/Azure/aks-engine/issues/3140))

### Maintenance 🔧
- rev Linux VHDs to 2020.08.24 ([#3750](https://github.com/Azure/aks-engine/issues/3750))
- add new Azure VM SKUs ([#3744](https://github.com/Azure/aks-engine/issues/3744))
- update windows default VHD for August ([#3730](https://github.com/Azure/aks-engine/issues/3730))
- update csi-secrets-store addon manifest and images ([#3728](https://github.com/Azure/aks-engine/issues/3728))
- upgrade metrics-server to v0.3.7 ([#3669](https://github.com/Azure/aks-engine/issues/3669))
- force delete addon manager when addon manager pods get stuck on terminating state ([#3685](https://github.com/Azure/aks-engine/issues/3685))
- Hyperv and upstream Containerd package support ([#3688](https://github.com/Azure/aks-engine/issues/3688))
- update Windows VHD to include 8B patches ([#3692](https://github.com/Azure/aks-engine/issues/3692))
- Update  docker version to fix log rotation issue ([#3693](https://github.com/Azure/aks-engine/issues/3693))
- update cluster-autoscaler to latest patch versions ([#3683](https://github.com/Azure/aks-engine/issues/3683))
- update cluster-autoscaler to latest patch versions ([#3650](https://github.com/Azure/aks-engine/issues/3650))
- rev Linux VHDs to 2020.07.24 ([#3649](https://github.com/Azure/aks-engine/issues/3649))
- update Azure CNI version to 1.1.6 ([#3644](https://github.com/Azure/aks-engine/issues/3644))
- bump cloud-controller-manager and cloud-node-manager to v0.5.1 ([#3636](https://github.com/Azure/aks-engine/issues/3636))
- update Windows default VHD for July ([#3619](https://github.com/Azure/aks-engine/issues/3619))
- add support for k8s v1.17.9 on Azure Stack ([#3612](https://github.com/Azure/aks-engine/issues/3612))
- add support for K8s v1.16.13 on Azure Stack ([#3613](https://github.com/Azure/aks-engine/issues/3613))
- update CoreDNS to v1.7.0 ([#3608](https://github.com/Azure/aks-engine/issues/3608))
- update Dashboard addon to v2.0.3 ([#3606](https://github.com/Azure/aks-engine/issues/3606))
- July 2020 windows patches ([#3598](https://github.com/Azure/aks-engine/issues/3598))
- update go toolchain to 1.14.4 ([#3596](https://github.com/Azure/aks-engine/issues/3596))
- Use moby 19.03.x packages ([#3549](https://github.com/Azure/aks-engine/issues/3549))
- upgrades force azure-cnms update strategy ([#3571](https://github.com/Azure/aks-engine/issues/3571))
- rev Linux VHDs to 2020.06.25 ([#3558](https://github.com/Azure/aks-engine/issues/3558))
- updating windows VHD used by aks-engine to include June k8s packages ([#3550](https://github.com/Azure/aks-engine/issues/3550))
- get-logs collects audit logs ([#3537](https://github.com/Azure/aks-engine/issues/3537))
- add support for K8s 1.16.11 on Azure Stack ([#3535](https://github.com/Azure/aks-engine/issues/3535))
- Adding 1.15.12, 1.16.10, 1.18.2 hotfix packages for windows ([#3525](https://github.com/Azure/aks-engine/issues/3525))
- bump moby to 3.0.13 ([#3404](https://github.com/Azure/aks-engine/issues/3404))
- improve vmss validation now that vmss is default ([#3526](https://github.com/Azure/aks-engine/issues/3526))
- rev Linux VHDs to 2020.06.22 ([#3527](https://github.com/Azure/aks-engine/issues/3527))
- increase default cloud provider's rate limits on Azure Stack ([#3515](https://github.com/Azure/aks-engine/issues/3515))
- Updates the pause image to 1.4.0 ([#3516](https://github.com/Azure/aks-engine/issues/3516))
- add 1.16.11, 1.17.7, 1.18.4 hotfix packages to Windows VHD ([#3506](https://github.com/Azure/aks-engine/issues/3506))
- add support for K8s 1.17.7 on Azure Stack ([#3509](https://github.com/Azure/aks-engine/issues/3509))
- Support to change LicenseType for Windows ([#3485](https://github.com/Azure/aks-engine/issues/3485))
- triage "k8s-master" references for eventual removal ([#3474](https://github.com/Azure/aks-engine/issues/3474))
- add csi windows related images ([#3448](https://github.com/Azure/aks-engine/issues/3448))
- update metrics-server to v0.3.6 ([#3463](https://github.com/Azure/aks-engine/issues/3463))
- rev azure-sdk-for-go to v43.0.0 ([#3467](https://github.com/Azure/aks-engine/issues/3467))
- Add June Windows Patches ([#3468](https://github.com/Azure/aks-engine/issues/3468))
- update pause image to v1.3.2 ([#3461](https://github.com/Azure/aks-engine/issues/3461))
- update Azure NPM to v1.1.4 ([#3452](https://github.com/Azure/aks-engine/issues/3452))
- June 2020 windows patches ([#3439](https://github.com/Azure/aks-engine/issues/3439))
- standardize nodeSelector declarations in addons ([#3419](https://github.com/Azure/aks-engine/issues/3419))
- bump csi-secrets-store image versions ([#3424](https://github.com/Azure/aks-engine/issues/3424))
- update aad-pod-identity manifest and version ([#3423](https://github.com/Azure/aks-engine/issues/3423))
- update addon-manager to v9.1.1 ([#3409](https://github.com/Azure/aks-engine/issues/3409))
- use local kubeconfig if no outbound ([#3410](https://github.com/Azure/aks-engine/issues/3410))
- standardize nodeSelectors ([#3406](https://github.com/Azure/aks-engine/issues/3406))
- rev Windows VHD to 17763.1217.200603 ([#3407](https://github.com/Azure/aks-engine/issues/3407))
- update node-problem-detector to v0.8.2 ([#3365](https://github.com/Azure/aks-engine/issues/3365))
- Add EnableAHUB in WindowsProfile ([#3322](https://github.com/Azure/aks-engine/issues/3322))
- rev AKS Engine VHDs to 2020.06.02 ([#3395](https://github.com/Azure/aks-engine/issues/3395))
- remove unused, deprecated blobfuse spec ([#3384](https://github.com/Azure/aks-engine/issues/3384))
- Bump moby version to 3.0.12 ([#3376](https://github.com/Azure/aks-engine/issues/3376))
- add support for K8s 1.16.10 & 1.17.6 on Azure Stack ([#3377](https://github.com/Azure/aks-engine/issues/3377))
- rev Linux VHDs to 2020.05.29 ([#3378](https://github.com/Azure/aks-engine/issues/3378))
- rev pause image to 1.3.1 ([#3370](https://github.com/Azure/aks-engine/issues/3370))
- get-logs collects vhd-install.complete ([#3372](https://github.com/Azure/aks-engine/issues/3372))
- change MTU only if not Azure CNI on Azure Stack ([#3367](https://github.com/Azure/aks-engine/issues/3367))
- upgrade cni to v1.1.3 ([#3353](https://github.com/Azure/aks-engine/issues/3353))
- remove old labels.yaml file ([#3357](https://github.com/Azure/aks-engine/issues/3357))
- update CNI binary to 0.8.6 ([#3332](https://github.com/Azure/aks-engine/issues/3332))
- k8s v1.17 conformance model for Azure Stack ([#3338](https://github.com/Azure/aks-engine/issues/3338))
- update kube-dashboard addon to v2.0.1 ([#3327](https://github.com/Azure/aks-engine/issues/3327))
- rev coredns to 1.6.9 ([#3328](https://github.com/Azure/aks-engine/issues/3328))
- rev default etcd version to 3.3.22 ([#3325](https://github.com/Azure/aks-engine/issues/3325))
- remove support for Kubernetes 1.14 ([#3310](https://github.com/Azure/aks-engine/issues/3310))
- make 1.18 the default Kubernetes version ([#3298](https://github.com/Azure/aks-engine/issues/3298))
- disable 18.04 + N series VM node pool ([#3275](https://github.com/Azure/aks-engine/issues/3275))
- don't force create etcd_disk filesystem ([#3276](https://github.com/Azure/aks-engine/issues/3276))
- update Windows pause image ([#3210](https://github.com/Azure/aks-engine/issues/3210))
- rev Linux VHDs to 2020.05.13 ([#3265](https://github.com/Azure/aks-engine/issues/3265))
- remove defunct build-windows-k8s.sh script ([#3248](https://github.com/Azure/aks-engine/issues/3248))
- update zip name for mooncake mirror ([#3237](https://github.com/Azure/aks-engine/issues/3237))
- update Windows VHD to use May 2020 updates ([#3236](https://github.com/Azure/aks-engine/issues/3236))
- updating cni to v1.1.2 ([#3192](https://github.com/Azure/aks-engine/issues/3192))
- update nvidia device plugin ([#3225](https://github.com/Azure/aks-engine/issues/3225))
- bump csi-secrets-store to v0.0.10, keyvault provider to 0.0.5 ([#3233](https://github.com/Azure/aks-engine/issues/3233))
- pull cached Azure Stack Windows binaries from shared SA ([#3223](https://github.com/Azure/aks-engine/issues/3223))
- Updating Azure NPM to v1.1.2 ([#3178](https://github.com/Azure/aks-engine/issues/3178))
- Update Policy addon versions ([#3208](https://github.com/Azure/aks-engine/issues/3208))
- get-logs command on Azure Stack ([#3216](https://github.com/Azure/aks-engine/issues/3216))
- Updating Antrea version to 0.6.0 ([#3204](https://github.com/Azure/aks-engine/issues/3204))
- update Windows VHD build to use mcr.microsoft.com/oss/kubernets/pause:1.3.1 ([#3209](https://github.com/Azure/aks-engine/issues/3209))
- append Azure Stack suffix to custom components images ([#3195](https://github.com/Azure/aks-engine/issues/3195))
- add new M-series SKUs ([#3190](https://github.com/Azure/aks-engine/issues/3190))
- mutex is not necessary now that rand uses local object ([#3048](https://github.com/Azure/aks-engine/issues/3048))
- check all error returns in main code ([#3184](https://github.com/Azure/aks-engine/issues/3184))
- update go-dev image to v1.27.0 ([#3181](https://github.com/Azure/aks-engine/issues/3181))
- check all error returns in tests ([#3183](https://github.com/Azure/aks-engine/issues/3183))
- aad-pod-identity taint ([#3143](https://github.com/Azure/aks-engine/issues/3143))
- add support for Kubernetes 1.17.4 & 1.17.5 on Azure Stack ([#3161](https://github.com/Azure/aks-engine/issues/3161))
- remove //+build !test pragma ([#3169](https://github.com/Azure/aks-engine/issues/3169))
- standardize to {{- for go template expressions ([#3166](https://github.com/Azure/aks-engine/issues/3166))
- remove KUBELET_IMAGE kubelet systemd env var ([#3165](https://github.com/Azure/aks-engine/issues/3165))
- remove deprecated KUBELET_REGISTER_SCHEDULABLE ([#3164](https://github.com/Azure/aks-engine/issues/3164))
- add support for Kubernetes 1.15.11, 1.16.8 & 1.16.9 on Azure Stack ([#3157](https://github.com/Azure/aks-engine/issues/3157))
- update cluster-autoscaler version for k8s 1.19 ([#3133](https://github.com/Azure/aks-engine/issues/3133))

### Testing 💚
- disable azure-arc-onboarding addon in everything cluster config ([#3756](https://github.com/Azure/aks-engine/issues/3756))
- update kubernetes e2e to use GINKGO_FAIL_FAST parameter value ([#3660](https://github.com/Azure/aks-engine/issues/3660))
- fix ginkgo failFast ([#3738](https://github.com/Azure/aks-engine/issues/3738))
- don't run kubelet restart test if ssh is blocked ([#3638](https://github.com/Azure/aks-engine/issues/3638))
- move E2E global vars out of "not block ssh port" conditional block ([#3635](https://github.com/Azure/aks-engine/issues/3635))
- run pod schedule tests first ([#3627](https://github.com/Azure/aks-engine/issues/3627))
- set default IDENTITY_SYSTEM ([#3591](https://github.com/Azure/aks-engine/issues/3591))
- don't return from GetAllRunningByLabelWithRetry until we have pods ([#3588](https://github.com/Azure/aks-engine/issues/3588))
- testing the recommended flow of sgx-device-plugin ([#3560](https://github.com/Azure/aks-engine/issues/3560))
- ensure all resource creation attempts are retried ([#3566](https://github.com/Azure/aks-engine/issues/3566))
- enable daemonset E2E test conveniences ([#3564](https://github.com/Azure/aks-engine/issues/3564))
- enable CustomKubeProxyImage override ([#3538](https://github.com/Azure/aks-engine/issues/3538))
- fix UT ([#3548](https://github.com/Azure/aks-engine/issues/3548))
- Add Windows SAC 2004 e2e ([#3518](https://github.com/Azure/aks-engine/issues/3518))
- pointing Windows containerd binaries to nightly builds for e2e test passes ([#3511](https://github.com/Azure/aks-engine/issues/3511))
- fix validation for windows + containerd ([#3508](https://github.com/Azure/aks-engine/issues/3508))
- fix validation for windows + containerd ([#3508](https://github.com/Azure/aks-engine/issues/3508))
- skip stability tests for "everything" E2E cluster config ([#3498](https://github.com/Azure/aks-engine/issues/3498))
- use regular timeout for functional DNS test ([#3487](https://github.com/Azure/aks-engine/issues/3487))
- retry pod scheduling attempts ([#3486](https://github.com/Azure/aks-engine/issues/3486))
- use full toleration for testing scheduling to control plane ([#3483](https://github.com/Azure/aks-engine/issues/3483))
- basic pods Running E2E tests for all addons ([#3482](https://github.com/Azure/aks-engine/issues/3482))
- assign default stability iterations in cluster.sh ([#3475](https://github.com/Azure/aks-engine/issues/3475))
- don't perform stability tests on long-running cluster configs ([#3464](https://github.com/Azure/aks-engine/issues/3464))
- test Azure NetworkPolicy against all Kubernetes versions ([#3455](https://github.com/Azure/aks-engine/issues/3455))
- generic async top nodes interface for E2E ([#3444](https://github.com/Azure/aks-engine/issues/3444))
- Disable failing kubnet tests ([#3436](https://github.com/Azure/aks-engine/issues/3436))
- improve coredns autoscaler E2E validation ([#3438](https://github.com/Azure/aks-engine/issues/3438))
- e2e tests download of larger files ([#3412](https://github.com/Azure/aks-engine/issues/3412))
- print podsecuritypolicy resources during E2E ([#3386](https://github.com/Azure/aks-engine/issues/3386))
- Add test support for windows server 1909 ([#3379](https://github.com/Azure/aks-engine/issues/3379))
- remove Kubernetes version restrictions from sgx tests ([#3374](https://github.com/Azure/aks-engine/issues/3374))
- run everything E2E config against all cluster versions ([#3348](https://github.com/Azure/aks-engine/issues/3348))
- skip azurefile test if using containerd ([#3316](https://github.com/Azure/aks-engine/issues/3316))
- use common timeout for repeat tests ([#3314](https://github.com/Azure/aks-engine/issues/3314))
- add stability test tolerance for calico cluster test configs ([#3287](https://github.com/Azure/aks-engine/issues/3287))
- az deployment group create in E2E runner ([#3249](https://github.com/Azure/aks-engine/issues/3249))
- rationalize "everything" E2E config for coredns autoscaler test ([#3206](https://github.com/Azure/aks-engine/issues/3206))
- reduce node count in E2E "everything" config ([#3200](https://github.com/Azure/aks-engine/issues/3200))
- fix e2e for custom clouds using ADFS ([#3150](https://github.com/Azure/aks-engine/issues/3150))
- rationalize E2E env defaults ([#3132](https://github.com/Azure/aks-engine/issues/3132))
- faster base config test ([#3130](https://github.com/Azure/aks-engine/issues/3130))
- e2e on custom clouds ([#3117](https://github.com/Azure/aks-engine/issues/3117))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.50.3"></a>
# [v0.50.3] - 2020-08-17
### Bug Fixes 🐞
- update v0.50 to new 17763.1158.200422 ([#3710](https://github.com/Azure/aks-engine/issues/3710))
- resilient etcd data disk mounting ([#3259](https://github.com/Azure/aks-engine/issues/3259))
- don't dry-run cluster-init spec ([#3247](https://github.com/Azure/aks-engine/issues/3247))
- custom cloud Azure CNI config for Windows ([#3228](https://github.com/Azure/aks-engine/issues/3228))
- don't validate Windows version in update scenario ([#3226](https://github.com/Azure/aks-engine/issues/3226))
- don't hardcode csi enableproxy in kubeclusterconfig.json ([#3127](https://github.com/Azure/aks-engine/issues/3127))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.51.1"></a>
# [v0.51.1] - 2020-08-17
### Bug Fixes 🐞
- update v0.51 to new 17763.1217.200514 ([#3709](https://github.com/Azure/aks-engine/issues/3709))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.52.1"></a>
# [v0.52.1] - 2020-08-17
### Bug Fixes 🐞
- update to latest image for this 0.52 release ([#3708](https://github.com/Azure/aks-engine/issues/3708))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.53.1"></a>
# [v0.53.1] - 2020-08-15
### Bug Fixes 🐞
- update to new 17763.1339.200718 ([#3707](https://github.com/Azure/aks-engine/issues/3707))

### Maintenance 🔧
- update default sku ([#3617](https://github.com/Azure/aks-engine/issues/3617))
- Cherrypick July patches to v0.53.0  ([#3610](https://github.com/Azure/aks-engine/issues/3610))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.54.1"></a>
# [v0.54.1] - 2020-08-14
### Bug Fixes 🐞
- use new Windows VHD image version ([#3701](https://github.com/Azure/aks-engine/issues/3701))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.54.0"></a>
# [v0.54.0] - 2020-07-30
### Bug Fixes 🐞
- only remove CRI if re-installing ([#3654](https://github.com/Azure/aks-engine/issues/3654))
- reinforce MCR migration during upgrade for older clusters ([#3625](https://github.com/Azure/aks-engine/issues/3625))
- updating hotfix packages used for windows inlcuding CVE and metric ([#3599](https://github.com/Azure/aks-engine/issues/3599))
- use correct VMSS name comparison during Windows scale ([#3590](https://github.com/Azure/aks-engine/issues/3590))
- always extract kubelet and kubectl from a custom URL if defined ([#3574](https://github.com/Azure/aks-engine/issues/3574))
- use mcr.azk8s.cn for Azure CNI in China cloud ([#3587](https://github.com/Azure/aks-engine/issues/3587))
- update akse upgrade and scale in the e2e test to include identity ([#3582](https://github.com/Azure/aks-engine/issues/3582))
- validate signed powershell packages in Azure and China Cloud ([#3579](https://github.com/Azure/aks-engine/issues/3579))
- kube-proxy upgrade from v1.15 to v1.16 ([#3570](https://github.com/Azure/aks-engine/issues/3570))
- change addons image pull policy to IfNotPresent ([#3562](https://github.com/Azure/aks-engine/issues/3562))

### Documentation 📘
- fix incorrect calico addons section heading ([#3565](https://github.com/Azure/aks-engine/issues/3565))

### Features 🌈
- add support for Kubernetes 1.16.13 ([#3602](https://github.com/Azure/aks-engine/issues/3602))
- Support configurable Windows PauseImage ([#3594](https://github.com/Azure/aks-engine/issues/3594))
- add support for Kubernetes 1.17.9 ([#3603](https://github.com/Azure/aks-engine/issues/3603))
- add support for Kubernetes 1.18.6 ([#3604](https://github.com/Azure/aks-engine/issues/3604))
- etcdStorageLimitGB for configurable etcd storage limit ([#3583](https://github.com/Azure/aks-engine/issues/3583))
- add hosts config agent support for Windows nodes ([#3572](https://github.com/Azure/aks-engine/issues/3572))
- Consume signed powershell scripts ([#3441](https://github.com/Azure/aks-engine/issues/3441))
- enable /etc/hosts config agent for AKS private cluster BYO DNS ([#3556](https://github.com/Azure/aks-engine/issues/3556))
- add support for Kubernetes v1.16.12 ([#3551](https://github.com/Azure/aks-engine/issues/3551))
- add support for Kubernetes v1.17.8 ([#3552](https://github.com/Azure/aks-engine/issues/3552))
- add support for Kubernetes v1.18.5 ([#3553](https://github.com/Azure/aks-engine/issues/3553))

### Maintenance 🔧
- rev Linux VHDs to 2020.07.24 ([#3649](https://github.com/Azure/aks-engine/issues/3649))
- bump cloud-controller-manager and cloud-node-manager to v0.5.1 ([#3636](https://github.com/Azure/aks-engine/issues/3636))
- update Windows default VHD for July ([#3619](https://github.com/Azure/aks-engine/issues/3619))
- add support for k8s v1.17.9 on Azure Stack ([#3612](https://github.com/Azure/aks-engine/issues/3612))
- add support for K8s v1.16.13 on Azure Stack ([#3613](https://github.com/Azure/aks-engine/issues/3613))
- update CoreDNS to v1.7.0 ([#3608](https://github.com/Azure/aks-engine/issues/3608))
- update Dashboard addon to v2.0.3 ([#3606](https://github.com/Azure/aks-engine/issues/3606))
- July 2020 windows patches ([#3598](https://github.com/Azure/aks-engine/issues/3598))
- update go toolchain to 1.14.4 ([#3596](https://github.com/Azure/aks-engine/issues/3596))
- Use moby 19.03.x packages ([#3549](https://github.com/Azure/aks-engine/issues/3549))
- upgrades force azure-cnms update strategy ([#3571](https://github.com/Azure/aks-engine/issues/3571))

### Testing 💚
- move E2E global vars out of "not block ssh port" conditional block ([#3635](https://github.com/Azure/aks-engine/issues/3635))
- run pod schedule tests first ([#3627](https://github.com/Azure/aks-engine/issues/3627))
- set default IDENTITY_SYSTEM ([#3591](https://github.com/Azure/aks-engine/issues/3591))
- don't return from GetAllRunningByLabelWithRetry until we have pods ([#3588](https://github.com/Azure/aks-engine/issues/3588))
- testing the recommended flow of sgx-device-plugin ([#3560](https://github.com/Azure/aks-engine/issues/3560))
- ensure all resource creation attempts are retried ([#3566](https://github.com/Azure/aks-engine/issues/3566))
- enable daemonset E2E test conveniences ([#3564](https://github.com/Azure/aks-engine/issues/3564))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.53.0"></a>
# [v0.53.0] - 2020-06-29
### Bug Fixes 🐞
- remove v1.17 support on Azure Stack ([#3544](https://github.com/Azure/aks-engine/issues/3544))
- downgrade blobfuse to 1.1.1 since there is breaking change on 1.2.3 ([#3540](https://github.com/Azure/aks-engine/issues/3540))
- always use /var/run/reboot-required ([#3536](https://github.com/Azure/aks-engine/issues/3536))
- updating hotfix packages used for windows to include hyperkube a… ([#3539](https://github.com/Azure/aks-engine/issues/3539))
- use kube-proxy image instead of hyperkube for 1.16 ([#3532](https://github.com/Azure/aks-engine/issues/3532))
- add windows hotfix overrides for 1.18.2, 1.16.10, 1.15.12 for AKS ([#3528](https://github.com/Azure/aks-engine/issues/3528))
- add validation for enableVMSSNodePublicIP + Basic LB ([#3524](https://github.com/Azure/aks-engine/issues/3524))
- ensure cron systemd job restarts ([#3523](https://github.com/Azure/aks-engine/issues/3523))
- use hotfix Windows builds for 1.18.4, 1.17.7, and 1.16.11 ([#3510](https://github.com/Azure/aks-engine/issues/3510))
- addons bootstrapping errata ([#3481](https://github.com/Azure/aks-engine/issues/3481))
- scale + AvailabilitySet errata ([#3490](https://github.com/Azure/aks-engine/issues/3490))
- make security context configs more restrictive ([#3454](https://github.com/Azure/aks-engine/issues/3454))
- validate that --upgrade-version is semver-compatible during upgrade ([#3442](https://github.com/Azure/aks-engine/issues/3442))
- Use mirror in China for AzureChinaCloud ([#3433](https://github.com/Azure/aks-engine/issues/3433))
- VolumePluginDir issue on Windows ([#3426](https://github.com/Azure/aks-engine/issues/3426))
- disable apiserver insecure-port ([#3422](https://github.com/Azure/aks-engine/issues/3422))
- only schedule mic on Linux nodes ([#3420](https://github.com/Azure/aks-engine/issues/3420))
- Enable e2e to install ginko if required ([#3330](https://github.com/Azure/aks-engine/issues/3330))

### Documentation 📘
- update throttling doc ([#3545](https://github.com/Azure/aks-engine/issues/3545))
- add Azure API throttling doc ([#3534](https://github.com/Azure/aks-engine/issues/3534))
- update azure-stack documentation ([#3479](https://github.com/Azure/aks-engine/issues/3479))
- supported versions for Azure Stacks using VHD 2020.05.13 ([#3457](https://github.com/Azure/aks-engine/issues/3457))
- remove references to "whitelist" ([#3428](https://github.com/Azure/aks-engine/issues/3428))

### Features 🌈
- add support for Kubernetes 1.18.4 ([#3500](https://github.com/Azure/aks-engine/issues/3500))
- add support for Kubernetes 1.17.7 ([#3499](https://github.com/Azure/aks-engine/issues/3499))
- add support for Kubernetes 1.16.11 ([#3496](https://github.com/Azure/aks-engine/issues/3496))
- enable configurable usageReportingEnabled in calico addon ([#3493](https://github.com/Azure/aks-engine/issues/3493))
- ensure all AKS required ports on api-server are exposed ([#3488](https://github.com/Azure/aks-engine/issues/3488))
- enable windows daemonset in container monitoring addon ([#3466](https://github.com/Azure/aks-engine/issues/3466))
- add support for Kubernetes v1.19.0-beta.2 ([#3471](https://github.com/Azure/aks-engine/issues/3471))
- Dualstack support for Windows containers ([#3415](https://github.com/Azure/aks-engine/issues/3415))
- flatcar nodes ([#3380](https://github.com/Azure/aks-engine/issues/3380))
- enable configurable apiserver --anonymous-auth ([#3430](https://github.com/Azure/aks-engine/issues/3430))
- Update container monitoring addon for may omsagent release  ([#3403](https://github.com/Azure/aks-engine/issues/3403))
- add support for Kubernetes 1.19.0-beta.1 ([#3408](https://github.com/Azure/aks-engine/issues/3408))

### Maintenance 🔧
- rev Linux VHDs to 2020.06.25 ([#3558](https://github.com/Azure/aks-engine/issues/3558))
- updating windows VHD used by aks-engine to include June k8s packages ([#3550](https://github.com/Azure/aks-engine/issues/3550))
- get-logs collects audit logs ([#3537](https://github.com/Azure/aks-engine/issues/3537))
- add support for K8s 1.16.11 on Azure Stack ([#3535](https://github.com/Azure/aks-engine/issues/3535))
- Adding 1.15.12, 1.16.10, 1.18.2 hotfix packages for windows ([#3525](https://github.com/Azure/aks-engine/issues/3525))
- bump moby to 3.0.13 ([#3404](https://github.com/Azure/aks-engine/issues/3404))
- improve vmss validation now that vmss is default ([#3526](https://github.com/Azure/aks-engine/issues/3526))
- rev Linux VHDs to 2020.06.22 ([#3527](https://github.com/Azure/aks-engine/issues/3527))
- increase default cloud provider's rate limits on Azure Stack ([#3515](https://github.com/Azure/aks-engine/issues/3515))
- Updates the pause image to 1.4.0 ([#3516](https://github.com/Azure/aks-engine/issues/3516))
- add 1.16.11, 1.17.7, 1.18.4 hotfix packages to Windows VHD ([#3506](https://github.com/Azure/aks-engine/issues/3506))
- add support for K8s 1.17.7 on Azure Stack ([#3509](https://github.com/Azure/aks-engine/issues/3509))
- Support to change LicenseType for Windows ([#3485](https://github.com/Azure/aks-engine/issues/3485))
- triage "k8s-master" references for eventual removal ([#3474](https://github.com/Azure/aks-engine/issues/3474))
- add csi windows related images ([#3448](https://github.com/Azure/aks-engine/issues/3448))
- update metrics-server to v0.3.6 ([#3463](https://github.com/Azure/aks-engine/issues/3463))
- rev azure-sdk-for-go to v43.0.0 ([#3467](https://github.com/Azure/aks-engine/issues/3467))
- Add June Windows Patches ([#3468](https://github.com/Azure/aks-engine/issues/3468))
- update pause image to v1.3.2 ([#3461](https://github.com/Azure/aks-engine/issues/3461))
- update Azure NPM to v1.1.4 ([#3452](https://github.com/Azure/aks-engine/issues/3452))
- June 2020 windows patches ([#3439](https://github.com/Azure/aks-engine/issues/3439))
- standardize nodeSelector declarations in addons ([#3419](https://github.com/Azure/aks-engine/issues/3419))
- bump csi-secrets-store image versions ([#3424](https://github.com/Azure/aks-engine/issues/3424))
- update aad-pod-identity manifest and version ([#3423](https://github.com/Azure/aks-engine/issues/3423))
- update addon-manager to v9.1.1 ([#3409](https://github.com/Azure/aks-engine/issues/3409))

### Testing 💚
- enable CustomKubeProxyImage override ([#3538](https://github.com/Azure/aks-engine/issues/3538))
- fix UT ([#3548](https://github.com/Azure/aks-engine/issues/3548))
- Add Windows SAC 2004 e2e ([#3518](https://github.com/Azure/aks-engine/issues/3518))
- pointing Windows containerd binaries to nightly builds for e2e test passes ([#3511](https://github.com/Azure/aks-engine/issues/3511))
- fix validation for windows + containerd ([#3508](https://github.com/Azure/aks-engine/issues/3508))
- fix validation for windows + containerd ([#3508](https://github.com/Azure/aks-engine/issues/3508))
- skip stability tests for "everything" E2E cluster config ([#3498](https://github.com/Azure/aks-engine/issues/3498))
- use regular timeout for functional DNS test ([#3487](https://github.com/Azure/aks-engine/issues/3487))
- retry pod scheduling attempts ([#3486](https://github.com/Azure/aks-engine/issues/3486))
- use full toleration for testing scheduling to control plane ([#3483](https://github.com/Azure/aks-engine/issues/3483))
- basic pods Running E2E tests for all addons ([#3482](https://github.com/Azure/aks-engine/issues/3482))
- assign default stability iterations in cluster.sh ([#3475](https://github.com/Azure/aks-engine/issues/3475))
- don't perform stability tests on long-running cluster configs ([#3464](https://github.com/Azure/aks-engine/issues/3464))
- test Azure NetworkPolicy against all Kubernetes versions ([#3455](https://github.com/Azure/aks-engine/issues/3455))
- generic async top nodes interface for E2E ([#3444](https://github.com/Azure/aks-engine/issues/3444))
- Disable failing kubnet tests ([#3436](https://github.com/Azure/aks-engine/issues/3436))
- improve coredns autoscaler E2E validation ([#3438](https://github.com/Azure/aks-engine/issues/3438))
- e2e tests download of larger files ([#3412](https://github.com/Azure/aks-engine/issues/3412))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.52.0"></a>
# [v0.52.0] - 2020-06-04
### Bug Fixes 🐞
- bypass AzurePublicCloud assumption in e2e suite ([#3371](https://github.com/Azure/aks-engine/issues/3371))
- use nvidia drivers version 418.126.02 for 18.04 ([#3366](https://github.com/Azure/aks-engine/issues/3366))
- use right nodeSelector for cluster-autoscaler addon ([#3350](https://github.com/Azure/aks-engine/issues/3350))
- use ctr to remove containerd containers ([#3144](https://github.com/Azure/aks-engine/issues/3144))

### Code Refactoring 💎
- consolidate tiller addon spec ([#3399](https://github.com/Azure/aks-engine/issues/3399))
- consolidate smb-flexvolume addon spec ([#3397](https://github.com/Azure/aks-engine/issues/3397))
- consolidate kube-proxy addon spec ([#3391](https://github.com/Azure/aks-engine/issues/3391))
- consolidate kube-dns addon spec ([#3390](https://github.com/Azure/aks-engine/issues/3390))
- consolidate keyvault-flexvolume addon spec ([#3389](https://github.com/Azure/aks-engine/issues/3389))
- consolidate flannel addon spec ([#3385](https://github.com/Azure/aks-engine/issues/3385))
- drop "beta" from kubernetes.io/os label ([#3369](https://github.com/Azure/aks-engine/issues/3369))
- consolidate aci-connector addon spec ([#3358](https://github.com/Azure/aks-engine/issues/3358))
- consolidate blobfuse-flexvolume addon spec ([#3359](https://github.com/Azure/aks-engine/issues/3359))
- consolidate calico addon spec ([#3361](https://github.com/Azure/aks-engine/issues/3361))
- consolidate cloud-node-manager addon spec ([#3363](https://github.com/Azure/aks-engine/issues/3363))
- consolidate aad-pod-identity addon spec ([#3356](https://github.com/Azure/aks-engine/issues/3356))
- consolidate audit-policy "addon" spec ([#3355](https://github.com/Azure/aks-engine/issues/3355))
- consolidate ip-masq-agent addon spec ([#3347](https://github.com/Azure/aks-engine/issues/3347))
- consolidate kube-rescheduler addon spec ([#3323](https://github.com/Azure/aks-engine/issues/3323))

### Continuous Integration 💜
- updating pub to v0.2.3 ([#3398](https://github.com/Azure/aks-engine/issues/3398))
- add notice file for windows vhd ([#3345](https://github.com/Azure/aks-engine/issues/3345))
- use ginkgo 1.12, run go mod vendor during make coverage ([#3343](https://github.com/Azure/aks-engine/issues/3343))
- enable multiple static resource group exclusions in cleanup script ([#3340](https://github.com/Azure/aks-engine/issues/3340))
- moving windows vhd/image related test configs to westus2 ([#3320](https://github.com/Azure/aks-engine/issues/3320))

### Documentation 📘
- update sgx doc and sgx-test tag ([#3349](https://github.com/Azure/aks-engine/issues/3349))

### Features 🌈
- add "get-versions --azure-env" flag to list custom clouds supported versions ([#3394](https://github.com/Azure/aks-engine/issues/3394))
- configurable calico logging verbosity ([#3396](https://github.com/Azure/aks-engine/issues/3396))
- deprecate heapster addon ([#3387](https://github.com/Azure/aks-engine/issues/3387))
- add support for Kubernetes 1.18.3 ([#3309](https://github.com/Azure/aks-engine/issues/3309))
- add support for Kubernetes 1.16.10 ([#3312](https://github.com/Azure/aks-engine/issues/3312))
- add support for Kubernetes 1.17.6 ([#3311](https://github.com/Azure/aks-engine/issues/3311))
- add support for Kubernetes 1.19.0-beta.0 ([#3299](https://github.com/Azure/aks-engine/issues/3299))

### Maintenance 🔧
- use local kubeconfig if no outbound ([#3410](https://github.com/Azure/aks-engine/issues/3410))
- standardize nodeSelectors ([#3406](https://github.com/Azure/aks-engine/issues/3406))
- rev Windows VHD to 17763.1217.200603 ([#3407](https://github.com/Azure/aks-engine/issues/3407))
- update node-problem-detector to v0.8.2 ([#3365](https://github.com/Azure/aks-engine/issues/3365))
- Add EnableAHUB in WindowsProfile ([#3322](https://github.com/Azure/aks-engine/issues/3322))
- rev AKS Engine VHDs to 2020.06.02 ([#3395](https://github.com/Azure/aks-engine/issues/3395))
- remove unused, deprecated blobfuse spec ([#3384](https://github.com/Azure/aks-engine/issues/3384))
- Bump moby version to 3.0.12 ([#3376](https://github.com/Azure/aks-engine/issues/3376))
- add support for K8s 1.16.10 & 1.17.6 on Azure Stack ([#3377](https://github.com/Azure/aks-engine/issues/3377))
- rev Linux VHDs to 2020.05.29 ([#3378](https://github.com/Azure/aks-engine/issues/3378))
- rev pause image to 1.3.1 ([#3370](https://github.com/Azure/aks-engine/issues/3370))
- get-logs collects vhd-install.complete ([#3372](https://github.com/Azure/aks-engine/issues/3372))
- change MTU only if not Azure CNI on Azure Stack ([#3367](https://github.com/Azure/aks-engine/issues/3367))
- upgrade cni to v1.1.3 ([#3353](https://github.com/Azure/aks-engine/issues/3353))
- remove old labels.yaml file ([#3357](https://github.com/Azure/aks-engine/issues/3357))
- update CNI binary to 0.8.6 ([#3332](https://github.com/Azure/aks-engine/issues/3332))
- k8s v1.17 conformance model for Azure Stack ([#3338](https://github.com/Azure/aks-engine/issues/3338))
- update kube-dashboard addon to v2.0.1 ([#3327](https://github.com/Azure/aks-engine/issues/3327))
- rev coredns to 1.6.9 ([#3328](https://github.com/Azure/aks-engine/issues/3328))
- rev default etcd version to 3.3.22 ([#3325](https://github.com/Azure/aks-engine/issues/3325))

### Testing 💚
- print podsecuritypolicy resources during E2E ([#3386](https://github.com/Azure/aks-engine/issues/3386))
- Add test support for windows server 1909 ([#3379](https://github.com/Azure/aks-engine/issues/3379))
- remove Kubernetes version restrictions from sgx tests ([#3374](https://github.com/Azure/aks-engine/issues/3374))
- run everything E2E config against all cluster versions ([#3348](https://github.com/Azure/aks-engine/issues/3348))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.51.0"></a>
# [v0.51.0] - 2020-05-21
### Bug Fixes 🐞
- ensure pod-security-policy is the first addon loaded ([#3313](https://github.com/Azure/aks-engine/issues/3313))
- prevent panic in Windows scale if no tags ([#3308](https://github.com/Azure/aks-engine/issues/3308))
- ensure ERR_K8S_API_SERVER_DNS_LOOKUP_FAIL is returned when DNS check fails ([#3303](https://github.com/Azure/aks-engine/issues/3303))
- addon-manager definition should be under "labels" ([#3302](https://github.com/Azure/aks-engine/issues/3302))
- whitespace trimming in cloud-provider spec is broken ([#3301](https://github.com/Azure/aks-engine/issues/3301))
- increase apiserver connection timeout for private apiserver ([#3293](https://github.com/Azure/aks-engine/issues/3293))
- fix issue provisioning windows nodes with kubenet ([#3300](https://github.com/Azure/aks-engine/issues/3300))
- ensure /mnt is ext4 ([#3297](https://github.com/Azure/aks-engine/issues/3297))
- resilient etcd data disk mounting ([#3259](https://github.com/Azure/aks-engine/issues/3259))
- USER_ASSIGNED_IDENTITY_ID is empty in azure.json ([#3254](https://github.com/Azure/aks-engine/issues/3254))
- update Windows image validate on Azure Stack Hub ([#3260](https://github.com/Azure/aks-engine/issues/3260))
- don't dry-run cluster-init spec ([#3247](https://github.com/Azure/aks-engine/issues/3247))
- custom cloud Azure CNI config for Windows ([#3228](https://github.com/Azure/aks-engine/issues/3228))
- Add Windows node pool to containerd tests ([#3238](https://github.com/Azure/aks-engine/issues/3238))
- USER_ASSIGNED_IDENTITY_ID is empty in azure.json ([#3232](https://github.com/Azure/aks-engine/issues/3232))
- update Azure Stack Windows binaries component name ([#3231](https://github.com/Azure/aks-engine/issues/3231))
- don't validate Windows version in update scenario ([#3226](https://github.com/Azure/aks-engine/issues/3226))
- set Windows defender process exclustion for containerd ([#3215](https://github.com/Azure/aks-engine/issues/3215))
- aad-pod-identity taints foo works in k8s < 1.16 ([#3199](https://github.com/Azure/aks-engine/issues/3199))
- update dcap checksum file ([#3172](https://github.com/Azure/aks-engine/issues/3172))
- added resource limits for networkmonitor daemonset ([#2624](https://github.com/Azure/aks-engine/issues/2624))
- don't hardcode pause image for containerD on Windows ([#3158](https://github.com/Azure/aks-engine/issues/3158))
- Expose error details in Windows CSE ([#3159](https://github.com/Azure/aks-engine/issues/3159))
- correct destination filename for scheduled maintenance addon ([#3142](https://github.com/Azure/aks-engine/issues/3142))
- Correct wrong file name in the logs ([#3138](https://github.com/Azure/aks-engine/issues/3138))
- don't hardcode csi enableproxy in kubeclusterconfig.json ([#3127](https://github.com/Azure/aks-engine/issues/3127))

### Build 🏭
- fetch k8s node binaries from kubernetesartifacts storage ([#3151](https://github.com/Azure/aks-engine/issues/3151))

### Code Refactoring 💎
- consolidate pod-security-policy addon spec ([#3284](https://github.com/Azure/aks-engine/issues/3284))
- consolidate cilium addons specs ([#3289](https://github.com/Azure/aks-engine/issues/3289))
- consolidate azure-network-policy addon spec ([#3283](https://github.com/Azure/aks-engine/issues/3283))
- consolidate azure-cloud-provider addon spec ([#3277](https://github.com/Azure/aks-engine/issues/3277))
- move etcd data disk mounting to cloud init  ([#3273](https://github.com/Azure/aks-engine/issues/3273))
- consolidate nvidia-device-plugin addon spec ([#3235](https://github.com/Azure/aks-engine/issues/3235))
- consolidate metrics-server addon spec ([#3211](https://github.com/Azure/aks-engine/issues/3211))
- use newer "az deployment group" CLI syntax ([#3213](https://github.com/Azure/aks-engine/issues/3213))
- consolidate cluster-autoscaler addon spec ([#3198](https://github.com/Azure/aks-engine/issues/3198))
- consolidate azure-cni-networkmonitor addon spec ([#3189](https://github.com/Azure/aks-engine/issues/3189))
- use simpler file names for addons maintenance ([#3187](https://github.com/Azure/aks-engine/issues/3187))
- simplify golangci-lint invocation ([#3182](https://github.com/Azure/aks-engine/issues/3182))
- simplify kubelet systemd service ([#3167](https://github.com/Azure/aks-engine/issues/3167))
- enable generic --register-with-taints and --register-node ([#3155](https://github.com/Azure/aks-engine/issues/3155))

### Continuous Integration 💜
- passing container runtime var when executing new-windows-sku.sh in dev container ([#3242](https://github.com/Azure/aks-engine/issues/3242))
- ensure sku title/summary/longsummary is unique when creating new marketplace skus ([#3239](https://github.com/Azure/aks-engine/issues/3239))
- enable staticcheck linter ([#3191](https://github.com/Azure/aks-engine/issues/3191))
- disable artifact publishing ([#3163](https://github.com/Azure/aks-engine/issues/3163))

### Documentation 📘
- aks-ubuntu-18.04 is default distro ([#3317](https://github.com/Azure/aks-engine/issues/3317))
- fix capitalization ([#3264](https://github.com/Azure/aks-engine/issues/3264))
- AKS-engine (on Stack) considerations around the use of proxy servers ([#3241](https://github.com/Azure/aks-engine/issues/3241))
- remove out of date Calico upgrade note ([#3243](https://github.com/Azure/aks-engine/issues/3243))
- update quickstart.md to add scoop install guide ([#3218](https://github.com/Azure/aks-engine/issues/3218))
- update sgx plugin docs ([#3230](https://github.com/Azure/aks-engine/issues/3230))
- rewrite monitoring topic ([#3205](https://github.com/Azure/aks-engine/issues/3205))
- add missing link to upgrade doc ([#3221](https://github.com/Azure/aks-engine/issues/3221))
- update sgx install instructions to dcv2 vms and newer resource name ([#3214](https://github.com/Azure/aks-engine/issues/3214))
- refresh documentation store ([#3177](https://github.com/Azure/aks-engine/issues/3177))
- remove obsolete api versions documentation ([#3175](https://github.com/Azure/aks-engine/issues/3175))
- document AKS Engine E2E ([#3131](https://github.com/Azure/aks-engine/issues/3131))
- include v0.48.0 in Azure Stack docs ([#3153](https://github.com/Azure/aks-engine/issues/3153))

### Features 🌈
- Updating Windows VHDs to include May patches ([#3263](https://github.com/Azure/aks-engine/issues/3263))
- enable alternate kube-reserved cgroups ([#3201](https://github.com/Azure/aks-engine/issues/3201))
- add support for Kubernetes 1.15.12 ([#3212](https://github.com/Azure/aks-engine/issues/3212))
- add support for Kubernetes 1.19.0-alpha.3 ([#3197](https://github.com/Azure/aks-engine/issues/3197))
- add warning message for empty Location string ([#3174](https://github.com/Azure/aks-engine/issues/3174))
- block imds from VMSS nodes if aad-pod-identity + msi ([#3136](https://github.com/Azure/aks-engine/issues/3136))
- Create Windows containerd VHDs ([#3162](https://github.com/Azure/aks-engine/issues/3162))
- configurable disk caching ([#2863](https://github.com/Azure/aks-engine/issues/2863))
- Kubernetes Dashboard addon v2.0.0 ([#3140](https://github.com/Azure/aks-engine/issues/3140))

### Maintenance 🔧
- remove support for Kubernetes 1.14 ([#3310](https://github.com/Azure/aks-engine/issues/3310))
- make 1.18 the default Kubernetes version ([#3298](https://github.com/Azure/aks-engine/issues/3298))
- disable 18.04 + N series VM node pool ([#3275](https://github.com/Azure/aks-engine/issues/3275))
- don't force create etcd_disk filesystem ([#3276](https://github.com/Azure/aks-engine/issues/3276))
- update Windows pause image ([#3210](https://github.com/Azure/aks-engine/issues/3210))
- rev Linux VHDs to 2020.05.13 ([#3265](https://github.com/Azure/aks-engine/issues/3265))
- remove defunct build-windows-k8s.sh script ([#3248](https://github.com/Azure/aks-engine/issues/3248))
- update zip name for mooncake mirror ([#3237](https://github.com/Azure/aks-engine/issues/3237))
- update Windows VHD to use May 2020 updates ([#3236](https://github.com/Azure/aks-engine/issues/3236))
- updating cni to v1.1.2 ([#3192](https://github.com/Azure/aks-engine/issues/3192))
- update nvidia device plugin ([#3225](https://github.com/Azure/aks-engine/issues/3225))
- bump csi-secrets-store to v0.0.10, keyvault provider to 0.0.5 ([#3233](https://github.com/Azure/aks-engine/issues/3233))
- pull cached Azure Stack Windows binaries from shared SA ([#3223](https://github.com/Azure/aks-engine/issues/3223))
- Updating Azure NPM to v1.1.2 ([#3178](https://github.com/Azure/aks-engine/issues/3178))
- Update Policy addon versions ([#3208](https://github.com/Azure/aks-engine/issues/3208))
- get-logs command on Azure Stack ([#3216](https://github.com/Azure/aks-engine/issues/3216))
- Updating Antrea version to 0.6.0 ([#3204](https://github.com/Azure/aks-engine/issues/3204))
- update Windows VHD build to use mcr.microsoft.com/oss/kubernets/pause:1.3.1 ([#3209](https://github.com/Azure/aks-engine/issues/3209))
- append Azure Stack suffix to custom components images ([#3195](https://github.com/Azure/aks-engine/issues/3195))
- add new M-series SKUs ([#3190](https://github.com/Azure/aks-engine/issues/3190))
- mutex is not necessary now that rand uses local object ([#3048](https://github.com/Azure/aks-engine/issues/3048))
- check all error returns in main code ([#3184](https://github.com/Azure/aks-engine/issues/3184))
- update go-dev image to v1.27.0 ([#3181](https://github.com/Azure/aks-engine/issues/3181))
- check all error returns in tests ([#3183](https://github.com/Azure/aks-engine/issues/3183))
- aad-pod-identity taint ([#3143](https://github.com/Azure/aks-engine/issues/3143))
- add support for Kubernetes 1.17.4 & 1.17.5 on Azure Stack ([#3161](https://github.com/Azure/aks-engine/issues/3161))
- remove //+build !test pragma ([#3169](https://github.com/Azure/aks-engine/issues/3169))
- standardize to {{- for go template expressions ([#3166](https://github.com/Azure/aks-engine/issues/3166))
- remove KUBELET_IMAGE kubelet systemd env var ([#3165](https://github.com/Azure/aks-engine/issues/3165))
- remove deprecated KUBELET_REGISTER_SCHEDULABLE ([#3164](https://github.com/Azure/aks-engine/issues/3164))
- add support for Kubernetes 1.15.11, 1.16.8 & 1.16.9 on Azure Stack ([#3157](https://github.com/Azure/aks-engine/issues/3157))
- update cluster-autoscaler version for k8s 1.19 ([#3133](https://github.com/Azure/aks-engine/issues/3133))

### Testing 💚
- skip azurefile test if using containerd ([#3316](https://github.com/Azure/aks-engine/issues/3316))
- use common timeout for repeat tests ([#3314](https://github.com/Azure/aks-engine/issues/3314))
- add stability test tolerance for calico cluster test configs ([#3287](https://github.com/Azure/aks-engine/issues/3287))
- az deployment group create in E2E runner ([#3249](https://github.com/Azure/aks-engine/issues/3249))
- rationalize "everything" E2E config for coredns autoscaler test ([#3206](https://github.com/Azure/aks-engine/issues/3206))
- reduce node count in E2E "everything" config ([#3200](https://github.com/Azure/aks-engine/issues/3200))
- fix e2e for custom clouds using ADFS ([#3150](https://github.com/Azure/aks-engine/issues/3150))
- rationalize E2E env defaults ([#3132](https://github.com/Azure/aks-engine/issues/3132))
- faster base config test ([#3130](https://github.com/Azure/aks-engine/issues/3130))
- e2e on custom clouds ([#3117](https://github.com/Azure/aks-engine/issues/3117))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.50.2"></a>
# [v0.50.2] - 2020-05-15
### Bug Fixes 🐞
- resilient etcd data disk mounting ([#3259](https://github.com/Azure/aks-engine/issues/3259))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.50.1"></a>
# [v0.50.1] - 2020-05-13
### Bug Fixes 🐞
- don't dry-run cluster-init spec ([#3247](https://github.com/Azure/aks-engine/issues/3247))
- custom cloud Azure CNI config for Windows ([#3228](https://github.com/Azure/aks-engine/issues/3228))
- don't validate Windows version in update scenario ([#3226](https://github.com/Azure/aks-engine/issues/3226))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.50.0"></a>
# [v0.50.0] - 2020-04-22
### Bug Fixes 🐞
- don't hardcode csi enableproxy in kubeclusterconfig.json ([#3127](https://github.com/Azure/aks-engine/issues/3127))
- idempotent systemd + disk mount for etcd ([#3126](https://github.com/Azure/aks-engine/issues/3126))
- use full userAssignedID reference in ARM template ([#3116](https://github.com/Azure/aks-engine/issues/3116))
- add nodeSelector for coredns-autoscaler ([#3109](https://github.com/Azure/aks-engine/issues/3109))
- typos in JSON annotations caught by staticcheck ([#3104](https://github.com/Azure/aks-engine/issues/3104))
- Fixed bug in condition check for dualstack ([#3110](https://github.com/Azure/aks-engine/issues/3110))
- update cluster-proportional-autoscaler path at MCR ([#3091](https://github.com/Azure/aks-engine/issues/3091))
- Update CSI Secrets Store Provider for Azure to 0.0.4 ([#3086](https://github.com/Azure/aks-engine/issues/3086))
- don't skip Windows VMSS node pools during upgrade ([#3083](https://github.com/Azure/aks-engine/issues/3083))
- retry more aggressively when install the gpu deb package. ([#3065](https://github.com/Azure/aks-engine/issues/3065))

### Code Refactoring 💎
- Create signable PS scripts ([#3015](https://github.com/Azure/aks-engine/issues/3015))
- skip sgx driver installation if already exists in OS ([#3062](https://github.com/Azure/aks-engine/issues/3062))

### Code Style 🎶
- override linguist's default language detection ([#3100](https://github.com/Azure/aks-engine/issues/3100))

### Documentation 📘
- join the #aks-engine-users Slack channel (not #provider-azure) ([#3103](https://github.com/Azure/aks-engine/issues/3103))
- remove docker package and choco caveat from release docs ([#3099](https://github.com/Azure/aks-engine/issues/3099))
- get-logs proposal ([#2745](https://github.com/Azure/aks-engine/issues/2745))

### Features 🌈
- add support for Kubernetes 1.19.0-alpha.2 ([#3122](https://github.com/Azure/aks-engine/issues/3122))
- Enable Antrea in NetworkPolicyOnly mode with Azure CNI ([#3027](https://github.com/Azure/aks-engine/issues/3027))
- Enabling SSH on windows nodes by default ([#2759](https://github.com/Azure/aks-engine/issues/2759))
- Updating Windows VHDs with 4B patches ([#3115](https://github.com/Azure/aks-engine/issues/3115))
- modify container runtime data dir ([#3072](https://github.com/Azure/aks-engine/issues/3072))
- enable CSI proxy by default when cloud-controller-manager is enabled on Windows cluster ([#3080](https://github.com/Azure/aks-engine/issues/3080))
- Azure CNI dual stack support ([#2862](https://github.com/Azure/aks-engine/issues/2862))
- enable multiple frontend IPs in Standard LB ([#3085](https://github.com/Azure/aks-engine/issues/3085))
- disable dashboard addon by default ([#3093](https://github.com/Azure/aks-engine/issues/3093))
- coredns-autoscaler ([#3067](https://github.com/Azure/aks-engine/issues/3067))
- add support for Kubernetes 1.17.5 ([#3088](https://github.com/Azure/aks-engine/issues/3088))
- add support for Kubernetes 1.18.2 ([#3089](https://github.com/Azure/aks-engine/issues/3089))
- add support for Kubernetes 1.16.9 ([#3087](https://github.com/Azure/aks-engine/issues/3087))
- disabling windows updates by default ([#3073](https://github.com/Azure/aks-engine/issues/3073))
- "aks-engine get-logs" command ([#2987](https://github.com/Azure/aks-engine/issues/2987))
- give metrics-server system-cluster-critical priority ([#3082](https://github.com/Azure/aks-engine/issues/3082))
- Support Non-Azure Stack Custom Clouds with custom endpoints/root certs/sources.list ([#3063](https://github.com/Azure/aks-engine/issues/3063))
- support cloud-node-manager on Windows clusters ([#3044](https://github.com/Azure/aks-engine/issues/3044))
- proximity placement group support ([#3056](https://github.com/Azure/aks-engine/issues/3056))
- generic cluster-init component ([#3023](https://github.com/Azure/aks-engine/issues/3023))

### Maintenance 🔧
- Use prod repos for containerd. ([#3107](https://github.com/Azure/aks-engine/issues/3107))
- update go compiler to 1.14.2 ([#2839](https://github.com/Azure/aks-engine/issues/2839))
- rev Linux VHDs to 2020.04.21 ([#3123](https://github.com/Azure/aks-engine/issues/3123))
- updating Windows VHD to include April patches ([#3101](https://github.com/Azure/aks-engine/issues/3101))
- remove pre-1.14 addons specs ([#3081](https://github.com/Azure/aks-engine/issues/3081))
- use Ubuntu 18.04-LTS as default ([#3070](https://github.com/Azure/aks-engine/issues/3070))
- remove support for Kubernetes 1.13 ([#3059](https://github.com/Azure/aks-engine/issues/3059))

### Testing 💚
- really fix bad distro value in everything cluster config ([#3095](https://github.com/Azure/aks-engine/issues/3095))
- fix bad distro value in everything cluster config ([#3094](https://github.com/Azure/aks-engine/issues/3094))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.49.0"></a>
# [v0.49.0] - 2020-04-11
### Bug Fixes 🐞
- make sure nvidia utility and compute bins can be mounted into co… ([#3029](https://github.com/Azure/aks-engine/issues/3029))
- the os disk size restriction is changed in azure. ([#3043](https://github.com/Azure/aks-engine/issues/3043))
- remove unnecessary apt lock waits ([#3049](https://github.com/Azure/aks-engine/issues/3049))
- In case of api server connection error, still complete script ([#3022](https://github.com/Azure/aks-engine/issues/3022))
- disable unattended upgrades during CSE execution ([#1681](https://github.com/Azure/aks-engine/issues/1681))
- don't download collectlogs.ps1 if it exists ([#3006](https://github.com/Azure/aks-engine/issues/3006))
- add node selector for csi-secrets-store ds ([#3007](https://github.com/Azure/aks-engine/issues/3007))
- check dns status before connection ([#3002](https://github.com/Azure/aks-engine/issues/3002))
- make build with go 1.14 ([#3005](https://github.com/Azure/aks-engine/issues/3005))
- change gatekeeper namespace in azure-policy addon ([#2991](https://github.com/Azure/aks-engine/issues/2991))
- each apiserver process listens on its own IP address ([#2953](https://github.com/Azure/aks-engine/issues/2953))
- use same cluster-autoscaler version in MCR ([#2981](https://github.com/Azure/aks-engine/issues/2981))
- use upstreamartifacts CDN URL for apmz ([#2962](https://github.com/Azure/aks-engine/issues/2962))
- add NVv3 series to NVidia driver install list ([#2959](https://github.com/Azure/aks-engine/issues/2959))
- check network to k8s api for vmss ([#2938](https://github.com/Azure/aks-engine/issues/2938))
- Get WindowsVersion from registry instead of calling Get-ComputerInfo ([#2954](https://github.com/Azure/aks-engine/issues/2954))

### Code Refactoring 💎
- retry nvidia drivers installation ([#3031](https://github.com/Azure/aks-engine/issues/3031))
- cse_install.sh for VHD reqs ([#3016](https://github.com/Azure/aks-engine/issues/3016))
- store CSE exit codes in code ([#3012](https://github.com/Azure/aks-engine/issues/3012))
- Parse sgx dcap version file, enable checksum ([#2942](https://github.com/Azure/aks-engine/issues/2942))

### Code Style 🎶
- format CSE scripts consistently ([#3020](https://github.com/Azure/aks-engine/issues/3020))

### Continuous Integration 💜
- new VHD pipeline definition for Ubuntu + gen2 images ([#2958](https://github.com/Azure/aks-engine/issues/2958))

### Documentation 📘
- update `aks-engine deploy` example output ([#3019](https://github.com/Azure/aks-engine/issues/3019))
- include Basic LB SKU limitations in Azure Stack page ([#2988](https://github.com/Azure/aks-engine/issues/2988))
- update examples location for zh docs ([#2920](https://github.com/Azure/aks-engine/issues/2920))

### Features 🌈
- Updating AKS to use April 2020 Windows VHDs by default ([#3060](https://github.com/Azure/aks-engine/issues/3060))
- default to MCR for Kubernetes images ([#3046](https://github.com/Azure/aks-engine/issues/3046))
- add support for Kubernetes 1.18.1 ([#3045](https://github.com/Azure/aks-engine/issues/3045))
- EncryptionAtHost support ([#3041](https://github.com/Azure/aks-engine/issues/3041))
- allow mixed mode Availability Zone configuration ([#3032](https://github.com/Azure/aks-engine/issues/3032))
- "aks-engine get-skus" command ([#2772](https://github.com/Azure/aks-engine/issues/2772))
- disable unneeded message-of-the-day sections ([#3000](https://github.com/Azure/aks-engine/issues/3000))
- deprecate kata-containers ([#3014](https://github.com/Azure/aks-engine/issues/3014))
- adding kubelet and csi-proxy-server as windows defender excluded processes ([#2967](https://github.com/Azure/aks-engine/issues/2967))
- add ability to use pre-existing user assigned MSI ([#2960](https://github.com/Azure/aks-engine/issues/2960))
- etcd metrics URL ([#2989](https://github.com/Azure/aks-engine/issues/2989))
- add csi-secrets-store addon ([#2936](https://github.com/Azure/aks-engine/issues/2936))
- Windows azure-cni with containerd ([#2864](https://github.com/Azure/aks-engine/issues/2864))
- add support for Kubernetes 1.19.0-alpha.1 ([#2982](https://github.com/Azure/aks-engine/issues/2982))
- add support for Kubernetes 1.18.0 ([#2957](https://github.com/Azure/aks-engine/issues/2957))
- deprecate CoreOS support ([#2945](https://github.com/Azure/aks-engine/issues/2945))

### Maintenance 🔧
- rev Linux VHDs to 2020.04.09 ([#3058](https://github.com/Azure/aks-engine/issues/3058))
- SinglePlacementGroup=false if VMSS + SLB ([#3054](https://github.com/Azure/aks-engine/issues/3054))
- ensure upgraded clusters use MCR images ([#3053](https://github.com/Azure/aks-engine/issues/3053))
- update cluster-autoscaler versions ([#3024](https://github.com/Azure/aks-engine/issues/3024))
- reduce CSE by optimizing var/func names ([#3034](https://github.com/Azure/aks-engine/issues/3034))
- upgrade to azure-sdk v41.0.0 ([#3035](https://github.com/Azure/aks-engine/issues/3035))
- updated network monitor to latest version 0.0.8 ([#3026](https://github.com/Azure/aks-engine/issues/3026))
- use Standard LoadBalancer as default ([#2998](https://github.com/Azure/aks-engine/issues/2998))
- change csi-proxy binary name ([#3003](https://github.com/Azure/aks-engine/issues/3003))
- return distinct CSE err for GPU drivers runtime configuration ([#2941](https://github.com/Azure/aks-engine/issues/2941))
- remove mlocate if auditd disabled ([#2999](https://github.com/Azure/aks-engine/issues/2999))
- switch dev image to go 1.13 version ([#2976](https://github.com/Azure/aks-engine/issues/2976))
- force Basic LB SKU on Azure Stack ([#2974](https://github.com/Azure/aks-engine/issues/2974))
- optimize apt_get_purge CSE func ([#2931](https://github.com/Azure/aks-engine/issues/2931))
- update etcd to 3.3.19 ([#2944](https://github.com/Azure/aks-engine/issues/2944))
- update azure-policy addon versions ([#2903](https://github.com/Azure/aks-engine/issues/2903))

### Testing 💚
- fail fast for cluster.sh test workflows ([#3009](https://github.com/Azure/aks-engine/issues/3009))
- UT for addons ip-masq-agent image override ([#2997](https://github.com/Azure/aks-engine/issues/2997))
- use 1.17 for E2E ([#2980](https://github.com/Azure/aks-engine/issues/2980))
- reduce PR E2E cluster jobs ([#2975](https://github.com/Azure/aks-engine/issues/2975))
- promote --failFast to an e2e flag, defaulting to false ([#2964](https://github.com/Azure/aks-engine/issues/2964))
- don't grab ps logs during E2E operation ([#2966](https://github.com/Azure/aks-engine/issues/2966))
- add egress tests which target pod and namespace. ([#2925](https://github.com/Azure/aks-engine/issues/2925))
- don't test version if customWindowsPackageURL ([#2951](https://github.com/Azure/aks-engine/issues/2951))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.48.0"></a>
# [v0.48.0] - 2020-03-20
### Bug Fixes 🐞
- $ERR_SYSTEMCTL_STOP_FAIL VHD CI errata ([#2934](https://github.com/Azure/aks-engine/issues/2934))
- use ntp instead of systemd-timesyncd in 18.04 ([#2815](https://github.com/Azure/aks-engine/issues/2815))
- use explicit kubeconfig in label-nodes.sh ([#2929](https://github.com/Azure/aks-engine/issues/2929))
- check network from node to k8s api server in CSE ([#2919](https://github.com/Azure/aks-engine/issues/2919))
- parse SGX driver url and compare checksum ([#2914](https://github.com/Azure/aks-engine/issues/2914))
- do not remove VMAS resources from template when adding a VMAS pool ([#2907](https://github.com/Azure/aks-engine/issues/2907))
- scale cannot find VM index for the VMAS+VHD+Linux case ([#2906](https://github.com/Azure/aks-engine/issues/2906))
- honor custom{component}Image fields ([#2886](https://github.com/Azure/aks-engine/issues/2886))
- apply_network_config: false for 18.04-LTS ([#2908](https://github.com/Azure/aks-engine/issues/2908))
- ensure label-nodes systemd service can be enabled ([#2915](https://github.com/Azure/aks-engine/issues/2915))
- add validation for new AKS engine Windows Version ([#2896](https://github.com/Azure/aks-engine/issues/2896))
- fixing nssm logging in windows CSE ([#2890](https://github.com/Azure/aks-engine/issues/2890))
- Windows no outbound fixes ([#2883](https://github.com/Azure/aks-engine/issues/2883))
- Do not remove patched hyperkube image in cleanUpContainerImages ([#2878](https://github.com/Azure/aks-engine/issues/2878))
- make master availabilityProfile an optional field for Azure Stack ([#2866](https://github.com/Azure/aks-engine/issues/2866))
- install cracklib-runtime on ubuntu ([#2871](https://github.com/Azure/aks-engine/issues/2871))
- address containerd errata ([#2865](https://github.com/Azure/aks-engine/issues/2865))
- update EncryptionConfiguration to latest ([#2856](https://github.com/Azure/aks-engine/issues/2856))
- explicitly set kubeconfig when using KUBECTL in CSE ([#2847](https://github.com/Azure/aks-engine/issues/2847))
- update sgx driver download urls ([#2807](https://github.com/Azure/aks-engine/issues/2807))
- update yaml.v2 and golangci-lint for go 1.14 compatibility ([#2824](https://github.com/Azure/aks-engine/issues/2824))
- fix the CNI temp file URL parsing in CSE ([#2817](https://github.com/Azure/aks-engine/issues/2817))
- override components image config during upgrade ([#2814](https://github.com/Azure/aks-engine/issues/2814))
- kube-proxy uses custom hyperkube on Azure Stack ([#2806](https://github.com/Azure/aks-engine/issues/2806))
- use non-reserved exit codes for bcc/bpf install ([#2785](https://github.com/Azure/aks-engine/issues/2785))
- components works with Azure Stack ([#2775](https://github.com/Azure/aks-engine/issues/2775))
- apiserver broken due to incorrect kms cachesize ([#2769](https://github.com/Azure/aks-engine/issues/2769))
- only configAddons if addons enabled ([#2774](https://github.com/Azure/aks-engine/issues/2774))
- support Azure containerd install in Debian ([#2766](https://github.com/Azure/aks-engine/issues/2766))
- Set vnetCidr by the address space of the vnet ([#2725](https://github.com/Azure/aks-engine/issues/2725))
- init Azure clients consistently ([#2739](https://github.com/Azure/aks-engine/issues/2739))
- update apiserver encryption provider config flag ([#2727](https://github.com/Azure/aks-engine/issues/2727))
- Adding skuLongSummary to sku template publishing file to unblock publishing pipeline ([#2729](https://github.com/Azure/aks-engine/issues/2729))
- remove requirement for ldflags or make test ([#2724](https://github.com/Azure/aks-engine/issues/2724))
- Set WindowsSku and ImageVersion by publisher and offer in creating ([#2689](https://github.com/Azure/aks-engine/issues/2689))
- update to use single omsagent yaml for all k8s versions to avoid any manual errors and easy maintainbility ([#2692](https://github.com/Azure/aks-engine/issues/2692))
- raising dynamic port range for windows to better support clusters with 200+ services ([#2688](https://github.com/Azure/aks-engine/issues/2688))
- unknown ostype when getting an index should probably create an error ([#2681](https://github.com/Azure/aks-engine/issues/2681))

### Build 🏭
- Publish junit test results to pipeline ([#2664](https://github.com/Azure/aks-engine/issues/2664))
- fetch k8s Windows .zip from kubernetesartifacts storage ([#2655](https://github.com/Azure/aks-engine/issues/2655))

### Code Refactoring 💎
- add KubernetesImageBaseType ([#2711](https://github.com/Azure/aks-engine/issues/2711))
- user-configurable components ([#2540](https://github.com/Azure/aks-engine/issues/2540))
- put Azure DC/OS SKUs list in a separate source file ([#2706](https://github.com/Azure/aks-engine/issues/2706))
- put Azure locations list in a separate source file ([#2705](https://github.com/Azure/aks-engine/issues/2705))

### Continuous Integration 💜
- stop testing k8s 1.13 ([#2749](https://github.com/Azure/aks-engine/issues/2749))

### Documentation 📘
- ubuntu-18.04-gen2 distro ([#2917](https://github.com/Azure/aks-engine/issues/2917))
- fixed large clusters doc page ([#2826](https://github.com/Azure/aks-engine/issues/2826))
- add CSE example in upgrade do-not-dos ([#2795](https://github.com/Azure/aks-engine/issues/2795))
- adding Kalya to OWNERS ([#2793](https://github.com/Azure/aks-engine/issues/2793))
- add data collection (telemetry) notice to readme ([#2674](https://github.com/Azure/aks-engine/issues/2674))
- proposal for Azure locations + SKUs automation ([#2694](https://github.com/Azure/aks-engine/issues/2694))

### Features 🌈
- installing csi-proxy for windows at node deployment time ([#2930](https://github.com/Azure/aks-engine/issues/2930))
- add UltraSSD support ([#2905](https://github.com/Azure/aks-engine/issues/2905))
- install Windows csi proxy during cluster creation ([#2854](https://github.com/Azure/aks-engine/issues/2854))
- Updating AKS to use March 2020 Windows VHDs by default ([#2911](https://github.com/Azure/aks-engine/issues/2911))
- add support for Kubernetes 1.17.4 ([#2899](https://github.com/Azure/aks-engine/issues/2899))
- add support for Kubernetes 1.15.11 ([#2897](https://github.com/Azure/aks-engine/issues/2897))
- add support for Kubernetes 1.16.8 ([#2898](https://github.com/Azure/aks-engine/issues/2898))
- configurable sysctl.d configuration ([#2880](https://github.com/Azure/aks-engine/issues/2880))
- upgrade control plane only ([#2635](https://github.com/Azure/aks-engine/issues/2635))
- allow iptables mode for dualstack 1.18+ ([#2882](https://github.com/Azure/aks-engine/issues/2882))
- update containermonitoring addon for february 2020 release ([#2850](https://github.com/Azure/aks-engine/issues/2850))
- collect Windows CSE logs during log collection ([#2858](https://github.com/Azure/aks-engine/issues/2858))
- add support for single stack IPv6 ([#2781](https://github.com/Azure/aks-engine/issues/2781))
- Experimental support for Windows+ContainerD ([#1322](https://github.com/Azure/aks-engine/issues/1322))
- add support for Kubernetes 1.15.10, 1.16.6 & 1.16.7 on Azure Stack ([#2834](https://github.com/Azure/aks-engine/issues/2834))
- release Windows VHD with 2C updates ([#2809](https://github.com/Azure/aks-engine/issues/2809))
- Add IsCredentialAutoGenerated for WindowsProfile ([#2804](https://github.com/Azure/aks-engine/issues/2804))
- "aks-engine get-locations" command ([#2771](https://github.com/Azure/aks-engine/issues/2771))
- add support for Kubernetes 1.18.0-beta.1 ([#2791](https://github.com/Azure/aks-engine/issues/2791))
- feb windows updated ([#2786](https://github.com/Azure/aks-engine/issues/2786))
- update pause image to 1.3.0 (includes 1903 and 1909 support) ([#2757](https://github.com/Azure/aks-engine/issues/2757))
- add support for Kubernetes 1.18.0-alpha.5 ([#2748](https://github.com/Azure/aks-engine/issues/2748))
- enable MCR KubernetesImageBaseType ([#2722](https://github.com/Azure/aks-engine/issues/2722))
- add support for Kubernetes 1.18.0-alpha.3 ([#2682](https://github.com/Azure/aks-engine/issues/2682))
- add template/options for using Shared Image Gallery ([#2687](https://github.com/Azure/aks-engine/issues/2687))
- Adding WindowsNodeReset.ps1 script to reset/cleanup state for nodes ([#2457](https://github.com/Azure/aks-engine/issues/2457))
- install bcc tools by default ([#2683](https://github.com/Azure/aks-engine/issues/2683))
- updating windows VHD for Feb k8s versions ([#2731](https://github.com/Azure/aks-engine/issues/2731))
- add support for Kubernetes 1.15.10 ([#2709](https://github.com/Azure/aks-engine/issues/2709))
- add support for Kubernetes 1.16.7 ([#2710](https://github.com/Azure/aks-engine/issues/2710))
- add support for Kubernetes 1.17.3 ([#2707](https://github.com/Azure/aks-engine/issues/2707))
- abort/warn if apimodel contains properties not supported by Azure Stack ([#2717](https://github.com/Azure/aks-engine/issues/2717))
- adding mcr.microsoft.com/oss/kubernetes/pause:1.3.0 to windows VHD ([#2702](https://github.com/Azure/aks-engine/issues/2702))
- read vm size from instance metadata service for windows cse telemetry ([#2663](https://github.com/Azure/aks-engine/issues/2663))
- multi AI telemetry keys ([#2606](https://github.com/Azure/aks-engine/issues/2606))

### Maintenance 🔧
- rev Linux VHDs to 2020.03.19 ([#2939](https://github.com/Azure/aks-engine/issues/2939))
- improve --client-id flag validation message ([#2935](https://github.com/Azure/aks-engine/issues/2935))
- rev default Kubernetes version to 1.15 ([#2932](https://github.com/Azure/aks-engine/issues/2932))
- rev AKS Engine Linux VHDs to 2020.03.16 ([#2918](https://github.com/Azure/aks-engine/issues/2918))
- new Kubernetes versions for Linux VHDs ([#2894](https://github.com/Azure/aks-engine/issues/2894))
- updating windows VHD to include 3B patches + march k8s packages ([#2902](https://github.com/Azure/aks-engine/issues/2902))
- update cluster-autoscaler for k8s 1.18 ([#2901](https://github.com/Azure/aks-engine/issues/2901))
- optimize CSE payload ([#2891](https://github.com/Azure/aks-engine/issues/2891))
- pre-install the version of apmz in CSE ([#2889](https://github.com/Azure/aks-engine/issues/2889))
- Bump moby to 3.0.11 ([#2887](https://github.com/Azure/aks-engine/issues/2887))
- rev Linux VHDs to 2020.03.10 ([#2881](https://github.com/Azure/aks-engine/issues/2881))
- don't re-run provision.sh if already called.  ([#2843](https://github.com/Azure/aks-engine/issues/2843))
- update default to recommended cluster settings on Azure Stack ([#2861](https://github.com/Azure/aks-engine/issues/2861))
- don't require azure.json on node vms ([#2849](https://github.com/Azure/aks-engine/issues/2849))
- update cni-plugins to v0.8.5 ([#2841](https://github.com/Azure/aks-engine/issues/2841))
- default to large cluster settings on Azure Stack ([#2832](https://github.com/Azure/aks-engine/issues/2832))
- update Azure CNI to v1.0.33 ([#2825](https://github.com/Azure/aks-engine/issues/2825))
- force image base to MCR if target cloud is Azure Stack ([#2802](https://github.com/Azure/aks-engine/issues/2802))
- update node-problem-detector to v0.8.1 ([#2808](https://github.com/Azure/aks-engine/issues/2808))
- pre-pull addon images hosted in MCR ([#2800](https://github.com/Azure/aks-engine/issues/2800))
- Update the containerd config ([#2780](https://github.com/Azure/aks-engine/issues/2780))
- Adding 2C patches to windows VHD which addresses some networking issues ([#2796](https://github.com/Azure/aks-engine/issues/2796))
- adding azure-cni v1.0.33 artifacts to VHDs ([#2790](https://github.com/Azure/aks-engine/issues/2790))
- rationalize AKS Engine VHD config ([#2755](https://github.com/Azure/aks-engine/issues/2755))
- update coredns to 1.6.7 ([#2783](https://github.com/Azure/aks-engine/issues/2783))
- update cluster-autoscaler for 1.15 and 1.16 ([#2776](https://github.com/Azure/aks-engine/issues/2776))
- bump keyvault-flexvol to v0.0.16 ([#2760](https://github.com/Azure/aks-engine/issues/2760))
- use MCR URI to validate outbound connectivity ([#2761](https://github.com/Azure/aks-engine/issues/2761))
- more lint! ([#2756](https://github.com/Azure/aks-engine/issues/2756))
- use azure containerd packages ([#2649](https://github.com/Azure/aks-engine/issues/2649))
- latent lint ([#2754](https://github.com/Azure/aks-engine/issues/2754))
- installing Docker EE 19.03.5 by default in Windows VHD ([#2751](https://github.com/Azure/aks-engine/issues/2751))
- update image base for windows vhd ([#2742](https://github.com/Azure/aks-engine/issues/2742))
- cse cleanup ([#2746](https://github.com/Azure/aks-engine/issues/2746))
- apply large ipv3 neigh GC settings to nodes of all sizes ([#2732](https://github.com/Azure/aks-engine/issues/2732))
- update cluster-autoscaler to 1.17.1 ([#2730](https://github.com/Azure/aks-engine/issues/2730))
- rev AKS Engine Linux VHDs to 2020.02.12 ([#2728](https://github.com/Azure/aks-engine/issues/2728))
- use k8s 1.17 for e2e tests in windows vhd pipeline ([#2719](https://github.com/Azure/aks-engine/issues/2719))
- gofmt to avoid lint errors ([#2699](https://github.com/Azure/aks-engine/issues/2699))
- Add ProgressPreference=SilentlyContinue to disable progress bar ([#2693](https://github.com/Azure/aks-engine/issues/2693))

### Revert Change ◀️
- "feat: install Windows csi proxy during cluster creation ([#2854](https://github.com/Azure/aks-engine/issues/2854))"

### Testing 💚
- simplified availabilityset E2E cluster config ([#2926](https://github.com/Azure/aks-engine/issues/2926))
- skip node ready test after scale down ([#2921](https://github.com/Azure/aks-engine/issues/2921))
- better skip test implementation ([#2924](https://github.com/Azure/aks-engine/issues/2924))
- only use 1 node per pool in no outbound test ([#2922](https://github.com/Azure/aks-engine/issues/2922))
- create functions in network policy specific file and let tests call it. ([#2904](https://github.com/Azure/aks-engine/issues/2904))
- add test config to use gen2 images for agent pools ([#2893](https://github.com/Azure/aks-engine/issues/2893))
- working Windows containerd URLs ([#2885](https://github.com/Azure/aks-engine/issues/2885))
- Add default deny egress test for net work policy. ([#2872](https://github.com/Azure/aks-engine/issues/2872))
- don't count nodes, just look for labels ([#2884](https://github.com/Azure/aks-engine/issues/2884))
- run node labels test later ([#2876](https://github.com/Azure/aks-engine/issues/2876))
- remove apiserver from master test ([#2875](https://github.com/Azure/aks-engine/issues/2875))
- improved master/agent pod validation ([#2797](https://github.com/Azure/aks-engine/issues/2797))
- reorder E2E tests ([#2784](https://github.com/Azure/aks-engine/issues/2784))
- get pod network info if dns validation fails ([#2778](https://github.com/Azure/aks-engine/issues/2778))
- add deployment stability tests ([#2767](https://github.com/Azure/aks-engine/issues/2767))
- configurable CONTAINER_RUNTIME via E2E ([#2753](https://github.com/Azure/aks-engine/issues/2753))
- remove non-working SGX E2E cluster config ([#2744](https://github.com/Azure/aks-engine/issues/2744))
- add LB_TEST_TIMEOUT default ([#2703](https://github.com/Azure/aks-engine/issues/2703))
- enable configurable LoadBalancer test timeout ([#2700](https://github.com/Azure/aks-engine/issues/2700))
- curl as a stability test client ([#2697](https://github.com/Azure/aks-engine/issues/2697))
- don't get vms before resource group exists ([#2696](https://github.com/Azure/aks-engine/issues/2696))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.47.0"></a>
# [v0.47.0] - 2020-02-18
### Bug Fixes 🐞
- raising dynamic port range for windows to better support clusters with 200+ services ([#2688](https://github.com/Azure/aks-engine/issues/2688))
- update CSE moby version substring match ([#2662](https://github.com/Azure/aks-engine/issues/2662))
- don't expire adminUser account ([#2661](https://github.com/Azure/aks-engine/issues/2661))
- remove extra param from 'systemctl_restart' invocation ([#2654](https://github.com/Azure/aks-engine/issues/2654))
- persist new pool after addpool operation ([#2653](https://github.com/Azure/aks-engine/issues/2653))
- make Windows scaling work ([#2643](https://github.com/Azure/aks-engine/issues/2643))
- enable upgrade in private cluster + jumpbox scenario ([#2634](https://github.com/Azure/aks-engine/issues/2634))
- default LB AllocatedOutboundPorts to 0 ([#2526](https://github.com/Azure/aks-engine/issues/2526))
- stop using type nesting in billing extensions to address SDK-for-Go breaking change ([#2640](https://github.com/Azure/aks-engine/issues/2640))
- private cluster + multiple masters in backend ILB pool ([#2646](https://github.com/Azure/aks-engine/issues/2646))
- custom kube-controller-manager template should reference hyperkube ([#2639](https://github.com/Azure/aks-engine/issues/2639))
- create role assignment for master VMs ([#2583](https://github.com/Azure/aks-engine/issues/2583))
- Skip storage class e2e test when ephemeral disk is used ([#2604](https://github.com/Azure/aks-engine/issues/2604))
- Moving publish release notes step after e2e test runs so windows vhd pipeline can be re-run ([#2595](https://github.com/Azure/aks-engine/issues/2595))

### Build 🏭
- fetch k8s Windows .zip from kubernetesartifacts storage ([#2655](https://github.com/Azure/aks-engine/issues/2655))
- fetch etcd from MCR container image ([#2644](https://github.com/Azure/aks-engine/issues/2644))
- fetch cni-plugins from kubernetesartifacts storage ([#2617](https://github.com/Azure/aks-engine/issues/2617))
- fetch azure-vnet-cni from kubernetesartifacts storage ([#2593](https://github.com/Azure/aks-engine/issues/2593))

### Code Refactoring 💎
- extract common script parameters into function ([#2651](https://github.com/Azure/aks-engine/issues/2651))
- cleanup duplicate functions and move apiversions to consts ([#2659](https://github.com/Azure/aks-engine/issues/2659))

### Continuous Integration 💜
- change apiVersion of prow deployments from v1beta2 to v1 ([#2669](https://github.com/Azure/aks-engine/issues/2669))

### Documentation 📘
- refer to kubernetesartifacts storage in examples and comments ([#2656](https://github.com/Azure/aks-engine/issues/2656))
- add support note to main readme ([#2623](https://github.com/Azure/aks-engine/issues/2623))
- remove VHD code comment suggesting registry is temporary ([#2602](https://github.com/Azure/aks-engine/issues/2602))

### Features 🌈
- updating windows VHD for Feb k8s versions ([#2731](https://github.com/Azure/aks-engine/issues/2731))
- add support for Kubernetes 1.15.10 ([#2709](https://github.com/Azure/aks-engine/issues/2709))
- add support for Kubernetes 1.16.7 ([#2710](https://github.com/Azure/aks-engine/issues/2710))
- add support for Kubernetes 1.17.3 ([#2707](https://github.com/Azure/aks-engine/issues/2707))
- abort/warn if apimodel contains properties not supported by Azure Stack ([#2717](https://github.com/Azure/aks-engine/issues/2717))
- read vm size from instance metadata service for windows cse telemetry ([#2663](https://github.com/Azure/aks-engine/issues/2663))
- add new VM SKUs ([#2647](https://github.com/Azure/aks-engine/issues/2647))
- updating aks-engine to use jan windows images by default ([#2636](https://github.com/Azure/aks-engine/issues/2636))
- validate VHD availability on azurestack ([#2342](https://github.com/Azure/aks-engine/issues/2342))
- add support for Kubernetes 1.16.6 ([#2588](https://github.com/Azure/aks-engine/issues/2588))
- fetching collect-windows-logs.ps1 during windows CSE ([#2615](https://github.com/Azure/aks-engine/issues/2615))
- add support for Kubernetes 1.15.9 ([#2589](https://github.com/Azure/aks-engine/issues/2589))
- add support for Kubernetes 1.17.2 ([#2612](https://github.com/Azure/aks-engine/issues/2612))
- add support for Kubernetes 1.18.0-alpha.2 ([#2610](https://github.com/Azure/aks-engine/issues/2610))
- Updating aks-windows vhd to to include jan patches ([#2603](https://github.com/Azure/aks-engine/issues/2603))

### Maintenance 🔧
- rev AKS Engine Linux VHDs to 2020.02.12 ([#2728](https://github.com/Azure/aks-engine/issues/2728))
- Add ProgressPreference=SilentlyContinue to disable progress bar ([#2693](https://github.com/Azure/aks-engine/issues/2693))
- don't refer to hyperkube in template if > 1.16 ([#2676](https://github.com/Azure/aks-engine/issues/2676))
- update Azure NPM to v1.0.32 ([#2665](https://github.com/Azure/aks-engine/issues/2665))
- rev Linux VHDs to 2020.01.30 ([#2660](https://github.com/Azure/aks-engine/issues/2660))
- use IsVMSSToBeUpgraded func to validate vmss pools in upgrade ([#2650](https://github.com/Azure/aks-engine/issues/2650))
- Bump moby to 3.0.10 ([#2613](https://github.com/Azure/aks-engine/issues/2613))
- update Azure cloud-provider components to v0.4.1 ([#2627](https://github.com/Azure/aks-engine/issues/2627))
- support v1.15.9 on Azure Stack ([#2628](https://github.com/Azure/aks-engine/issues/2628))
- Updating NPM CPU and Memory requests and limits ([#2530](https://github.com/Azure/aks-engine/issues/2530))
- update go-dev tools image ([#2534](https://github.com/Azure/aks-engine/issues/2534))
- Adding 1.17.1 package to windows vhd ([#2607](https://github.com/Azure/aks-engine/issues/2607))

### Revert Change ◀️
- "chore: update go-dev tools image ([#2534](https://github.com/Azure/aks-engine/issues/2534))" ([#2632](https://github.com/Azure/aks-engine/issues/2632))

### Testing 💚
- replace containerd+kubenet with containerd+azure ([#2678](https://github.com/Azure/aks-engine/issues/2678))
- enable scale/upgrade for base E2E cluster config ([#2670](https://github.com/Azure/aks-engine/issues/2670))
- delete pod if exists in run pod scenario ([#2668](https://github.com/Azure/aks-engine/issues/2668))
- wait for "Ready" nodes, where appropriate ([#2657](https://github.com/Azure/aks-engine/issues/2657))
- zones aren't available everywhere ([#2633](https://github.com/Azure/aks-engine/issues/2633))
- allow cluster-autoscaler more nodes in everything test config ([#2629](https://github.com/Azure/aks-engine/issues/2629))
- modify vm sku during e2e upgrade tests ([#2625](https://github.com/Azure/aks-engine/issues/2625))
- HPA/cluster-autoscaler fixes so we don't scale up too many pods ([#2626](https://github.com/Azure/aks-engine/issues/2626))
- configurable pvc E2E tests ([#2620](https://github.com/Azure/aks-engine/issues/2620))
- don't test for static node count if add node pool context ([#2619](https://github.com/Azure/aks-engine/issues/2619))
- consolidate E2E test cluster configurations ([#2586](https://github.com/Azure/aks-engine/issues/2586))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.46.3"></a>
# [v0.46.3] - 2020-01-30
### Bug Fixes 🐞
- enable upgrade in private cluster + jumpbox scenario ([#2634](https://github.com/Azure/aks-engine/issues/2634))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.46.2"></a>
# [v0.46.2] - 2020-01-29
### Bug Fixes 🐞
- private cluster + multiple masters in backend ILB pool ([#2646](https://github.com/Azure/aks-engine/issues/2646))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.46.1"></a>
# [v0.46.1] - 2020-01-28
### Bug Fixes 🐞
- custom kube-controller-manager template should reference hyperkube ([#2639](https://github.com/Azure/aks-engine/issues/2639))
- create role assignment for master VMs ([#2583](https://github.com/Azure/aks-engine/issues/2583))

### Features 🌈
- updating aks-engine to use jan windows images by default ([#2636](https://github.com/Azure/aks-engine/issues/2636))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.46.0"></a>
# [v0.46.0] - 2020-01-17
### Bug Fixes 🐞
- Fix e2e test failure ([#2594](https://github.com/Azure/aks-engine/issues/2594))
- enable SLB with private clusters ([#2572](https://github.com/Azure/aks-engine/issues/2572))
- don't configure password for "etcd" user ([#2570](https://github.com/Azure/aks-engine/issues/2570))
- update list of SKUs with accelerated networking support ([#2566](https://github.com/Azure/aks-engine/issues/2566))
- k8s component versions errata ([#2565](https://github.com/Azure/aks-engine/issues/2565))
- marketplace-sku.yaml - fix for secret variables ([#2556](https://github.com/Azure/aks-engine/issues/2556))
- Remove PodPriority feature gate ([#2554](https://github.com/Azure/aks-engine/issues/2554))
- promote system addons to system-cluster-critical ([#2533](https://github.com/Azure/aks-engine/issues/2533))
- aci-connector region is ignored ([#2535](https://github.com/Azure/aks-engine/issues/2535))
- configure addons before setting kubelet config ([#2513](https://github.com/Azure/aks-engine/issues/2513))
- apply new master node labels for k8s v1.18+ compatibility ([#2467](https://github.com/Azure/aks-engine/issues/2467))
- Fix some path handling in collect-windows-logs script ([#2488](https://github.com/Azure/aks-engine/issues/2488))
- hard-coding hyper-v generation when using VHD URls as a quick unblock ([#2487](https://github.com/Azure/aks-engine/issues/2487))
- fix ARM dependency issues with vm user-specified extensions on node pools ([#2398](https://github.com/Azure/aks-engine/issues/2398))
- disable accelerated networking for Windows due to instability ([#2453](https://github.com/Azure/aks-engine/issues/2453))
- change kernel settings based on number of cores ([#2375](https://github.com/Azure/aks-engine/issues/2375))
- restore VHD CI spec to match versions support ([#2439](https://github.com/Azure/aks-engine/issues/2439))
- kube-dns and coredns addons UT, back-compat ([#2433](https://github.com/Azure/aks-engine/issues/2433))
- Rotate app insights key b/c resource group was accidently deleted :( ([#2423](https://github.com/Azure/aks-engine/issues/2423))
- default useCloudControllerManager to false for 1.17 ([#2372](https://github.com/Azure/aks-engine/issues/2372))
- install tools bins locally and add to PATH ([#2382](https://github.com/Azure/aks-engine/issues/2382))
- enable CSI addons on upgrade when appropriate ([#2347](https://github.com/Azure/aks-engine/issues/2347))
- correct defaults order for setting IPAddressCount ([#2358](https://github.com/Azure/aks-engine/issues/2358))
- enable cloud-node-manager addon on upgrade when appropriate ([#2345](https://github.com/Azure/aks-engine/issues/2345))
- update Container monitoring add-on with 11012019 agent image changes ([#2318](https://github.com/Azure/aks-engine/issues/2318))
- VHD release notes should output the kernel version at end ([#2338](https://github.com/Azure/aks-engine/issues/2338))
- support cloud-node-manager in k8s 1.16 ([#2330](https://github.com/Azure/aks-engine/issues/2330))
- userAssignedIdentityId in windows azure.json missing quotes ([#2327](https://github.com/Azure/aks-engine/issues/2327))
- lint ([#2322](https://github.com/Azure/aks-engine/issues/2322))
- lint errors ([#2319](https://github.com/Azure/aks-engine/issues/2319))
- remove support for v1.13.12, v1.14.8, v1.15.5, v1.16.2 ([#2314](https://github.com/Azure/aks-engine/issues/2314))
- extract kubelet and kubectl binaries from hyperkube's binary folder when possible ([#2298](https://github.com/Azure/aks-engine/issues/2298))
- add retry/logging/validation around docker command for windows ([#2257](https://github.com/Azure/aks-engine/issues/2257))
- make ensure-generated should call make generate ([#2289](https://github.com/Azure/aks-engine/issues/2289))
- addons "update config" is upgrade-only ([#2282](https://github.com/Azure/aks-engine/issues/2282))
- bumping timeout for getting management ip on windows nodes ([#2284](https://github.com/Azure/aks-engine/issues/2284))
- move e2e runner into a module ([#2262](https://github.com/Azure/aks-engine/issues/2262))
- Add 1.0.28 Azure CNI zip file into windows vhd ([#2268](https://github.com/Azure/aks-engine/issues/2268))

### Build 🏭
- fetch img tool from upstreamartifacts storage ([#2591](https://github.com/Azure/aks-engine/issues/2591))

### Code Refactoring 💎
- standardize to "addons", deprecate "containeraddons" ([#2525](https://github.com/Azure/aks-engine/issues/2525))
- user-configurable flannel and scheduled maintenance addons ([#2517](https://github.com/Azure/aks-engine/issues/2517))
- move StorageClass into azure-cloud-provider addon ([#2497](https://github.com/Azure/aks-engine/issues/2497))
- make audit-policy and azure-cloud-provider addons user-configurable ([#2496](https://github.com/Azure/aks-engine/issues/2496))
- make cilium addon user-configurable ([#2480](https://github.com/Azure/aks-engine/issues/2480))
- make aad addon user-configurable ([#2471](https://github.com/Azure/aks-engine/issues/2471))
- make pod-security-policy addon user-configurable ([#2463](https://github.com/Azure/aks-engine/issues/2463))
- make kube-proxy addon user-configurable ([#2426](https://github.com/Azure/aks-engine/issues/2426))
- make coredns addon user-configurable ([#2416](https://github.com/Azure/aks-engine/issues/2416))
- make kube-dns addon user-configurable ([#2393](https://github.com/Azure/aks-engine/issues/2393))
- move addons consts to pkg/api/common ([#2383](https://github.com/Azure/aks-engine/issues/2383))
- remove unnecessary enabled addon config state representation ([#2381](https://github.com/Azure/aks-engine/issues/2381))
- optimize CSE main script for go template ([#2339](https://github.com/Azure/aks-engine/issues/2339))
- use go template for kubelet systemd config ([#2320](https://github.com/Azure/aks-engine/issues/2320))
- rationalize defaults enforcement order ([#2274](https://github.com/Azure/aks-engine/issues/2274))
- render manifest specs at template generation time ([#2313](https://github.com/Azure/aks-engine/issues/2313))
- install kube- components without hyperkube for k8s 1.17+ ([#2191](https://github.com/Azure/aks-engine/issues/2191))
- use go template comments instead of bash comments ([#2304](https://github.com/Azure/aks-engine/issues/2304))
- process cloud-init files via go template ([#2290](https://github.com/Azure/aks-engine/issues/2290))

### Code Style 🎶
- remove dead code from Windows build script ([#2361](https://github.com/Azure/aks-engine/issues/2361))

### Continuous Integration 💜
- add AvailabilitySet + SLB E2E scenario ([#2578](https://github.com/Azure/aks-engine/issues/2578))
- collect logs during E2E runs ([#2520](https://github.com/Azure/aks-engine/issues/2520))
- fix some Jenkins variables ([#2449](https://github.com/Azure/aks-engine/issues/2449))
- use Jenkins Azure Service Principal instead of hardcoding default client ID/Secret variables in pipelines ([#2438](https://github.com/Azure/aks-engine/issues/2438))
- stop testing k8s 1.11 and 1.12 on PRs ([#2303](https://github.com/Azure/aks-engine/issues/2303))
- ignore generated file in codecov ([#2258](https://github.com/Azure/aks-engine/issues/2258))

### Documentation 📘
- remove mentions of old orchestrators ([#2501](https://github.com/Azure/aks-engine/issues/2501))
- Issue 2464 improve agentpool extensions docs ([#2469](https://github.com/Azure/aks-engine/issues/2469))
- clarify Key Vault as secrets documentation ([#2402](https://github.com/Azure/aks-engine/issues/2402))
- configurable component images design doc ([#2408](https://github.com/Azure/aks-engine/issues/2408))
- remove k8s v1.16.1 from Azure Stack doc
- update kubernetes-developers.md ([#2353](https://github.com/Azure/aks-engine/issues/2353))
- fix bad Linux kernel info on 16.04 vhd 2019.10.24 ([#2341](https://github.com/Azure/aks-engine/issues/2341))
- Update sgx device plugin image name ([#2325](https://github.com/Azure/aks-engine/issues/2325))
- Add SGX Device plugin installation details ([#2321](https://github.com/Azure/aks-engine/issues/2321))
- supported versions for disconnected Azure Stacks using VHD 2019.10.24 ([#2237](https://github.com/Azure/aks-engine/issues/2237))
- Working way to use SGX ([#2291](https://github.com/Azure/aks-engine/issues/2291))

### Features 🌈
- Add Node Pool ([#2557](https://github.com/Azure/aks-engine/issues/2557))
- Upgrade Azure Disk CSI driver to support various new features ([#2541](https://github.com/Azure/aks-engine/issues/2541))
- add cse telemetry with apmz ([#2415](https://github.com/Azure/aks-engine/issues/2415))
- Set the default Windows sku and version when upgrading a VMSS ([#2581](https://github.com/Azure/aks-engine/issues/2581))
- add support for Kubernetes 1.17.1 ([#2579](https://github.com/Azure/aks-engine/issues/2579))
- Cse tracing for windows nodes ([#2400](https://github.com/Azure/aks-engine/issues/2400))
- Win vhd release pipeline ([#2567](https://github.com/Azure/aks-engine/issues/2567))
- Azure Spot VMSS ([#2547](https://github.com/Azure/aks-engine/issues/2547))
- create windows vhd publishing file in build pipeline ([#2529](https://github.com/Azure/aks-engine/issues/2529))
- build pipeline to create new marketplace vhd skus for windows ([#2537](https://github.com/Azure/aks-engine/issues/2537))
- update containermonitoring addon for december release ([#2481](https://github.com/Azure/aks-engine/issues/2481))
- cleaning up old kubelet/kubeproxy logs for Windows nodes ([#2504](https://github.com/Azure/aks-engine/issues/2504))
- add support for Kubernetes 1.18.0-alpha.1 ([#2503](https://github.com/Azure/aks-engine/issues/2503))
- Antrea plugin support in AKS Engine ([#2407](https://github.com/Azure/aks-engine/issues/2407))
- Configuring docker log rotation for Windows nodes ([#2478](https://github.com/Azure/aks-engine/issues/2478))
- Adds kubeconfig option to upgrade cmd ([#2397](https://github.com/Azure/aks-engine/issues/2397))
- Move creation of 'ext' HNS network into cse for Windows ([#2450](https://github.com/Azure/aks-engine/issues/2450))
- BYOK support on os disk ([#2412](https://github.com/Azure/aks-engine/issues/2412))
- enable user-configurable platformUpdateDomainCount ([#2459](https://github.com/Azure/aks-engine/issues/2459))
- add support for Kubernetes 1.16.4 ([#2436](https://github.com/Azure/aks-engine/issues/2436))
- add support for Kubernetes 1.14.10 ([#2434](https://github.com/Azure/aks-engine/issues/2434))
- add support for Kubernetes 1.15.7 ([#2435](https://github.com/Azure/aks-engine/issues/2435))
- caching application insights bins in windows vhd ([#2409](https://github.com/Azure/aks-engine/issues/2409))
- Adding some settings for telemetry collection to api model ([#2417](https://github.com/Azure/aks-engine/issues/2417))
- add support for Kubernetes 1.17.0 ([#2413](https://github.com/Azure/aks-engine/issues/2413))
- add support for Kubernetes 1.17.0-rc.2 ([#2386](https://github.com/Azure/aks-engine/issues/2386))
- Adding a unique id to Windows VHDs for telemtry correlation ([#2388](https://github.com/Azure/aks-engine/issues/2388))
- Updating Windows VHD build def to container Nov patches ([#2387](https://github.com/Azure/aks-engine/issues/2387))
- node-problem-detector addon ([#2371](https://github.com/Azure/aks-engine/issues/2371))
- add support for Kubernetes 1.17.0-rc.1 ([#2356](https://github.com/Azure/aks-engine/issues/2356))
- Support custom k8s components in v1.17 ([#2333](https://github.com/Azure/aks-engine/issues/2333))
- add support for Kubernetes 1.17.0-beta.2 ([#2334](https://github.com/Azure/aks-engine/issues/2334))
- support transparent mode for Azure CNI ([#2259](https://github.com/Azure/aks-engine/issues/2259))
- add support for Kubernetes v1.17.0-beta.1 ([#2158](https://github.com/Azure/aks-engine/issues/2158))
- add DisableOutboundSNAT feature ([#1708](https://github.com/Azure/aks-engine/issues/1708))
- support using existing identity for addons ([#2238](https://github.com/Azure/aks-engine/issues/2238))
- enable multi-node pool cluster-autoscaler addon ([#2138](https://github.com/Azure/aks-engine/issues/2138))
- expose PlatformFaultDomainCount in the apimodel  ([#2207](https://github.com/Azure/aks-engine/issues/2207))

### Maintenance 🔧
- update VHD references to 2020.01.15 ([#2600](https://github.com/Azure/aks-engine/issues/2600))
- move VHD files to vhd directory ([#2590](https://github.com/Azure/aks-engine/issues/2590))
- lint ([#2585](https://github.com/Azure/aks-engine/issues/2585))
- update VHD references to 2020.01.10 ([#2564](https://github.com/Azure/aks-engine/issues/2564))
- bump coredns to v1.6.6 ([#2555](https://github.com/Azure/aks-engine/issues/2555))
- update addon-resizer ([#2527](https://github.com/Azure/aks-engine/issues/2527))
- update Azure NPM to v1.0.31 ([#2521](https://github.com/Azure/aks-engine/issues/2521))
- Targeting dec patches for windows VHD ([#2505](https://github.com/Azure/aks-engine/issues/2505))
- pre-pull k8s v1.15.7-azs ([#2490](https://github.com/Azure/aks-engine/issues/2490))
- update cloud-provider-azure components to v0.4.0 ([#2473](https://github.com/Azure/aks-engine/issues/2473))
- lint ([#2493](https://github.com/Azure/aks-engine/issues/2493))
- upgrade cni-plugins to v0.7.6 ([#2484](https://github.com/Azure/aks-engine/issues/2484))
- use go template comments for generate proxy certs script ([#2336](https://github.com/Azure/aks-engine/issues/2336))
- update azure-npm-daemonset addon to v1.0.30 ([#2472](https://github.com/Azure/aks-engine/issues/2472))
- rev etcd to 3.3.18 ([#2462](https://github.com/Azure/aks-engine/issues/2462))
- update cluster-autoscaler for k8s <= 1.16 ([#2452](https://github.com/Azure/aks-engine/issues/2452))
- rev Linux VHDs to 2019.12.11 ([#2451](https://github.com/Azure/aks-engine/issues/2451))
- releasing Windows VHD with 1.16.4, 1.15.7 k8s packages ([#2447](https://github.com/Azure/aks-engine/issues/2447))
- remove 1.14.10. support ([#2440](https://github.com/Azure/aks-engine/issues/2440))
- update VHDs for Kubernetes patch releases ([#2437](https://github.com/Azure/aks-engine/issues/2437))
- releasing new Windows VHD with Nov 2019 patches/updated bins ([#2424](https://github.com/Azure/aks-engine/issues/2424))
- rev Linux VHDs to 2019.12.09 ([#2427](https://github.com/Azure/aks-engine/issues/2427))
- Add new VM skus (Ev4, ND40, ...) ([#2418](https://github.com/Azure/aks-engine/issues/2418))
- remove support for Kubernetes 1.11 and 1.12 ([#2276](https://github.com/Azure/aks-engine/issues/2276))
- update go toolchain to 1.13.5 ([#2405](https://github.com/Azure/aks-engine/issues/2405))
- update azure sdk to support BYOK ([#2395](https://github.com/Azure/aks-engine/issues/2395))
- update cluster-autoscaler to 1.17.0 ([#2390](https://github.com/Azure/aks-engine/issues/2390))
- update azure sdk for BYOK support ([#2376](https://github.com/Azure/aks-engine/issues/2376))
- clean up unused addons artifacts ([#2380](https://github.com/Azure/aks-engine/issues/2380))
- update CoreDNS to 1.6.5 ([#2357](https://github.com/Azure/aks-engine/issues/2357))
- rev VHD images to 2019.11.18 ([#2340](https://github.com/Azure/aks-engine/issues/2340))
- remove obsolete RP apis ([#2140](https://github.com/Azure/aks-engine/issues/2140))
- add CSI components to VHD ([#2332](https://github.com/Azure/aks-engine/issues/2332))
- use go template comments for addons yaml specs ([#2326](https://github.com/Azure/aks-engine/issues/2326))
- Updating Azure CNI and Azure NPM to v1.0.29 ([#2323](https://github.com/Azure/aks-engine/issues/2323))
- use Azure CCM for k8s 1.16 and later ([#2161](https://github.com/Azure/aks-engine/issues/2161))
- improve build through parallel tests and stable tools ([#2293](https://github.com/Azure/aks-engine/issues/2293))
- Adding more debug files to windows VHD ([#2253](https://github.com/Azure/aks-engine/issues/2253))
- lock the golangci-lint to 1.21.0 and use vendored modules ([#2279](https://github.com/Azure/aks-engine/issues/2279))

### Performance Improvements 🚀
- upgrade AKS tunnel image in VHD ([#2546](https://github.com/Azure/aks-engine/issues/2546))

### Testing 💚
- fix unset ADD_NODE_POOL_INPUT var in cluster.sh test script ([#2599](https://github.com/Azure/aks-engine/issues/2599))
- don't run ssh tests on spot vmss pools ([#2576](https://github.com/Azure/aks-engine/issues/2576))
- use vanilla busybox image for E2E tests ([#2544](https://github.com/Azure/aks-engine/issues/2544))
- use northeurope for byok testing ([#2536](https://github.com/Azure/aks-engine/issues/2536))
- use LOCATION env var for api model in E2E tests ([#2542](https://github.com/Azure/aks-engine/issues/2542))
- don't test non-working >= 1.16 flannel + docker ([#2524](https://github.com/Azure/aks-engine/issues/2524))
- revert change to default kubernetes.json api model example ([#2494](https://github.com/Azure/aks-engine/issues/2494))
- restore Windows node labels spec ([#2466](https://github.com/Azure/aks-engine/issues/2466))
- use uksouth instead of westus2 for AZ tests ([#2468](https://github.com/Azure/aks-engine/issues/2468))
- improve resiliency of addons UT implementations ([#2456](https://github.com/Azure/aks-engine/issues/2456))
- ensure clean local git before switching branches ([#2458](https://github.com/Azure/aks-engine/issues/2458))
- check k8s events to see if node-problem-detector works ([#2444](https://github.com/Azure/aks-engine/issues/2444))
- add node-problem-detector UT ([#2445](https://github.com/Azure/aks-engine/issues/2445))
- stop testing 1.12 ([#2430](https://github.com/Azure/aks-engine/issues/2430))
- add UT for kubernetesContainerAddonSettingsInit ([#2399](https://github.com/Azure/aks-engine/issues/2399))
- ensure cluster-autoscaler tests don't run if min-nodes is > nodes in pool ([#2368](https://github.com/Azure/aks-engine/issues/2368))
- use SLB in E2E cluster configs ([#2360](https://github.com/Azure/aks-engine/issues/2360))
- more resilient kubectl run interfaces ([#2355](https://github.com/Azure/aks-engine/issues/2355))
- resilient az group create during E2E provisioning ([#2359](https://github.com/Azure/aks-engine/issues/2359))
- re-enable E2E test for inotify max_user_watches ([#2351](https://github.com/Azure/aks-engine/issues/2351))
- resilient print pod logs during E2E, print more pod logs ([#2352](https://github.com/Azure/aks-engine/issues/2352))
- add k8s 1.16 and 1.17 to TestExampleAPIModel unit test ([#2354](https://github.com/Azure/aks-engine/issues/2354))
- remove windows-specific tests, use windows node pools ([#2349](https://github.com/Azure/aks-engine/issues/2349))
- add RunWithRetry to E2E, use for curl tests ([#2346](https://github.com/Azure/aks-engine/issues/2346))
- enable "everything" cluster config test for 1.17 ([#2335](https://github.com/Azure/aks-engine/issues/2335))
- ensure manifests aren't change by go template processing ([#2310](https://github.com/Azure/aks-engine/issues/2310))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.43.3"></a>
# [v0.43.3] - 2019-12-19
### Bug Fixes 🐞
- remove support for v1.13.12, v1.14.8, v1.15.5, v1.16.2
- Change Azure CNI default version to match latest windows vhd ([#2271](https://github.com/Azure/aks-engine/issues/2271))

### Maintenance 🔧
- generated code

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.45.0"></a>
# [v0.45.0] - 2019-12-13
### Bug Fixes 🐞
- disable accelerated networking for Windows due to instability ([#2453](https://github.com/Azure/aks-engine/issues/2453))
- change kernel settings based on number of cores ([#2375](https://github.com/Azure/aks-engine/issues/2375))
- restore VHD CI spec to match versions support ([#2439](https://github.com/Azure/aks-engine/issues/2439))
- kube-dns and coredns addons UT, back-compat ([#2433](https://github.com/Azure/aks-engine/issues/2433))
- Rotate app insights key b/c resource group was accidently deleted :( ([#2423](https://github.com/Azure/aks-engine/issues/2423))
- default useCloudControllerManager to false for 1.17 ([#2372](https://github.com/Azure/aks-engine/issues/2372))
- install tools bins locally and add to PATH ([#2382](https://github.com/Azure/aks-engine/issues/2382))
- enable CSI addons on upgrade when appropriate ([#2347](https://github.com/Azure/aks-engine/issues/2347))
- correct defaults order for setting IPAddressCount ([#2358](https://github.com/Azure/aks-engine/issues/2358))
- enable cloud-node-manager addon on upgrade when appropriate ([#2345](https://github.com/Azure/aks-engine/issues/2345))
- update Container monitoring add-on with 11012019 agent image changes ([#2318](https://github.com/Azure/aks-engine/issues/2318))
- VHD release notes should output the kernel version at end ([#2338](https://github.com/Azure/aks-engine/issues/2338))
- support cloud-node-manager in k8s 1.16 ([#2330](https://github.com/Azure/aks-engine/issues/2330))
- userAssignedIdentityId in windows azure.json missing quotes ([#2327](https://github.com/Azure/aks-engine/issues/2327))

### Code Refactoring 💎
- make kube-proxy addon user-configurable ([#2426](https://github.com/Azure/aks-engine/issues/2426))
- make coredns addon user-configurable ([#2416](https://github.com/Azure/aks-engine/issues/2416))
- make kube-dns addon user-configurable ([#2393](https://github.com/Azure/aks-engine/issues/2393))
- move addons consts to pkg/api/common ([#2383](https://github.com/Azure/aks-engine/issues/2383))
- remove unnecessary enabled addon config state representation ([#2381](https://github.com/Azure/aks-engine/issues/2381))
- optimize CSE main script for go template ([#2339](https://github.com/Azure/aks-engine/issues/2339))
- use go template for kubelet systemd config ([#2320](https://github.com/Azure/aks-engine/issues/2320))

### Code Style 🎶
- remove dead code from Windows build script ([#2361](https://github.com/Azure/aks-engine/issues/2361))

### Continuous Integration 💜
- fix some Jenkins variables ([#2449](https://github.com/Azure/aks-engine/issues/2449))
- use Jenkins Azure Service Principal instead of hardcoding default client ID/Secret variables in pipelines ([#2438](https://github.com/Azure/aks-engine/issues/2438))
- stop testing k8s 1.11 and 1.12 on PRs ([#2303](https://github.com/Azure/aks-engine/issues/2303))

### Documentation 📘
- clarify Key Vault as secrets documentation ([#2402](https://github.com/Azure/aks-engine/issues/2402))
- configurable component images design doc ([#2408](https://github.com/Azure/aks-engine/issues/2408))
- remove k8s v1.16.1 from Azure Stack doc
- update kubernetes-developers.md ([#2353](https://github.com/Azure/aks-engine/issues/2353))
- fix bad Linux kernel info on 16.04 vhd 2019.10.24 ([#2341](https://github.com/Azure/aks-engine/issues/2341))

### Features 🌈
- add support for Kubernetes 1.16.4 ([#2436](https://github.com/Azure/aks-engine/issues/2436))
- add support for Kubernetes 1.14.10 ([#2434](https://github.com/Azure/aks-engine/issues/2434))
- add support for Kubernetes 1.15.7 ([#2435](https://github.com/Azure/aks-engine/issues/2435))
- caching application insights bins in windows vhd ([#2409](https://github.com/Azure/aks-engine/issues/2409))
- Adding some settings for telemetry collection to api model ([#2417](https://github.com/Azure/aks-engine/issues/2417))
- add support for Kubernetes 1.17.0 ([#2413](https://github.com/Azure/aks-engine/issues/2413))
- add support for Kubernetes 1.17.0-rc.2 ([#2386](https://github.com/Azure/aks-engine/issues/2386))
- Adding a unique id to Windows VHDs for telemtry correlation ([#2388](https://github.com/Azure/aks-engine/issues/2388))
- Updating Windows VHD build def to container Nov patches ([#2387](https://github.com/Azure/aks-engine/issues/2387))
- node-problem-detector addon ([#2371](https://github.com/Azure/aks-engine/issues/2371))
- add support for Kubernetes 1.17.0-rc.1 ([#2356](https://github.com/Azure/aks-engine/issues/2356))
- Support custom k8s components in v1.17 ([#2333](https://github.com/Azure/aks-engine/issues/2333))
- add support for Kubernetes 1.17.0-beta.2 ([#2334](https://github.com/Azure/aks-engine/issues/2334))

### Maintenance 🔧
- rev Linux VHDs to 2019.12.11 ([#2451](https://github.com/Azure/aks-engine/issues/2451))
- releasing Windows VHD with 1.16.4, 1.15.7 k8s packages ([#2447](https://github.com/Azure/aks-engine/issues/2447))
- remove 1.14.10. support ([#2440](https://github.com/Azure/aks-engine/issues/2440))
- update VHDs for Kubernetes patch releases ([#2437](https://github.com/Azure/aks-engine/issues/2437))
- releasing new Windows VHD with Nov 2019 patches/updated bins ([#2424](https://github.com/Azure/aks-engine/issues/2424))
- rev Linux VHDs to 2019.12.09 ([#2427](https://github.com/Azure/aks-engine/issues/2427))
- Add new VM skus (Ev4, ND40, ...) ([#2418](https://github.com/Azure/aks-engine/issues/2418))
- remove support for Kubernetes 1.11 and 1.12 ([#2276](https://github.com/Azure/aks-engine/issues/2276))
- update go toolchain to 1.13.5 ([#2405](https://github.com/Azure/aks-engine/issues/2405))
- update azure sdk to support BYOK ([#2395](https://github.com/Azure/aks-engine/issues/2395))
- update cluster-autoscaler to 1.17.0 ([#2390](https://github.com/Azure/aks-engine/issues/2390))
- update azure sdk for BYOK support ([#2376](https://github.com/Azure/aks-engine/issues/2376))
- clean up unused addons artifacts ([#2380](https://github.com/Azure/aks-engine/issues/2380))
- update CoreDNS to 1.6.5 ([#2357](https://github.com/Azure/aks-engine/issues/2357))
- rev VHD images to 2019.11.18 ([#2340](https://github.com/Azure/aks-engine/issues/2340))
- remove obsolete RP apis ([#2140](https://github.com/Azure/aks-engine/issues/2140))
- add CSI components to VHD ([#2332](https://github.com/Azure/aks-engine/issues/2332))
- use go template comments for addons yaml specs ([#2326](https://github.com/Azure/aks-engine/issues/2326))
- Updating Azure CNI and Azure NPM to v1.0.29 ([#2323](https://github.com/Azure/aks-engine/issues/2323))

### Testing 💚
- add node-problem-detector UT ([#2445](https://github.com/Azure/aks-engine/issues/2445))
- stop testing 1.12 ([#2430](https://github.com/Azure/aks-engine/issues/2430))
- add UT for kubernetesContainerAddonSettingsInit ([#2399](https://github.com/Azure/aks-engine/issues/2399))
- ensure cluster-autoscaler tests don't run if min-nodes is > nodes in pool ([#2368](https://github.com/Azure/aks-engine/issues/2368))
- use SLB in E2E cluster configs ([#2360](https://github.com/Azure/aks-engine/issues/2360))
- more resilient kubectl run interfaces ([#2355](https://github.com/Azure/aks-engine/issues/2355))
- resilient az group create during E2E provisioning ([#2359](https://github.com/Azure/aks-engine/issues/2359))
- re-enable E2E test for inotify max_user_watches ([#2351](https://github.com/Azure/aks-engine/issues/2351))
- resilient print pod logs during E2E, print more pod logs ([#2352](https://github.com/Azure/aks-engine/issues/2352))
- add k8s 1.16 and 1.17 to TestExampleAPIModel unit test ([#2354](https://github.com/Azure/aks-engine/issues/2354))
- remove windows-specific tests, use windows node pools ([#2349](https://github.com/Azure/aks-engine/issues/2349))
- add RunWithRetry to E2E, use for curl tests ([#2346](https://github.com/Azure/aks-engine/issues/2346))
- enable "everything" cluster config test for 1.17 ([#2335](https://github.com/Azure/aks-engine/issues/2335))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.44.2"></a>
# [v0.44.2] - 2019-11-25
### Bug Fixes 🐞
- enable CSI addons on upgrade when appropriate ([#2347](https://github.com/Azure/aks-engine/issues/2347))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.44.1"></a>
# [v0.44.1] - 2019-11-22
### Bug Fixes 🐞
- correct defaults order for setting IPAddressCount ([#2358](https://github.com/Azure/aks-engine/issues/2358))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.44.0"></a>
# [v0.44.0] - 2019-11-18
### Bug Fixes 🐞
- support cloud-node-manager in k8s 1.16 ([#2330](https://github.com/Azure/aks-engine/issues/2330))
- userAssignedIdentityId in windows azure.json missing quotes ([#2327](https://github.com/Azure/aks-engine/issues/2327))
- lint ([#2322](https://github.com/Azure/aks-engine/issues/2322))
- lint errors ([#2319](https://github.com/Azure/aks-engine/issues/2319))
- remove support for v1.13.12, v1.14.8, v1.15.5, v1.16.2 ([#2314](https://github.com/Azure/aks-engine/issues/2314))
- extract kubelet and kubectl binaries from hyperkube's binary folder when possible ([#2298](https://github.com/Azure/aks-engine/issues/2298))
- add retry/logging/validation around docker command for windows ([#2257](https://github.com/Azure/aks-engine/issues/2257))
- make ensure-generated should call make generate ([#2289](https://github.com/Azure/aks-engine/issues/2289))
- addons "update config" is upgrade-only ([#2282](https://github.com/Azure/aks-engine/issues/2282))
- bumping timeout for getting management ip on windows nodes ([#2284](https://github.com/Azure/aks-engine/issues/2284))
- move e2e runner into a module ([#2262](https://github.com/Azure/aks-engine/issues/2262))
- Add 1.0.28 Azure CNI zip file into windows vhd ([#2268](https://github.com/Azure/aks-engine/issues/2268))

### Code Refactoring 💎
- rationalize defaults enforcement order ([#2274](https://github.com/Azure/aks-engine/issues/2274))
- render manifest specs at template generation time ([#2313](https://github.com/Azure/aks-engine/issues/2313))
- install kube- components without hyperkube for k8s 1.17+ ([#2191](https://github.com/Azure/aks-engine/issues/2191))
- use go template comments instead of bash comments ([#2304](https://github.com/Azure/aks-engine/issues/2304))
- process cloud-init files via go template ([#2290](https://github.com/Azure/aks-engine/issues/2290))

### Continuous Integration 💜
- ignore generated file in codecov ([#2258](https://github.com/Azure/aks-engine/issues/2258))

### Documentation 📘
- Update sgx device plugin image name ([#2325](https://github.com/Azure/aks-engine/issues/2325))
- Add SGX Device plugin installation details ([#2321](https://github.com/Azure/aks-engine/issues/2321))
- supported versions for disconnected Azure Stacks using VHD 2019.10.24 ([#2237](https://github.com/Azure/aks-engine/issues/2237))
- Working way to use SGX ([#2291](https://github.com/Azure/aks-engine/issues/2291))

### Features 🌈
- support transparent mode for Azure CNI ([#2259](https://github.com/Azure/aks-engine/issues/2259))
- add support for Kubernetes v1.17.0-beta.1 ([#2158](https://github.com/Azure/aks-engine/issues/2158))
- add DisableOutboundSNAT feature ([#1708](https://github.com/Azure/aks-engine/issues/1708))
- support using existing identity for addons ([#2238](https://github.com/Azure/aks-engine/issues/2238))
- enable multi-node pool cluster-autoscaler addon ([#2138](https://github.com/Azure/aks-engine/issues/2138))
- expose PlatformFaultDomainCount in the apimodel  ([#2207](https://github.com/Azure/aks-engine/issues/2207))

### Maintenance 🔧
- rev VHD images to 2019.11.18 ([#2340](https://github.com/Azure/aks-engine/issues/2340))
- use Azure CCM for k8s 1.16 and later ([#2161](https://github.com/Azure/aks-engine/issues/2161))
- improve build through parallel tests and stable tools ([#2293](https://github.com/Azure/aks-engine/issues/2293))
- Adding more debug files to windows VHD ([#2253](https://github.com/Azure/aks-engine/issues/2253))
- lock the golangci-lint to 1.21.0 and use vendored modules ([#2279](https://github.com/Azure/aks-engine/issues/2279))

### Testing 💚
- ensure manifests aren't change by go template processing ([#2310](https://github.com/Azure/aks-engine/issues/2310))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.43.2"></a>
# [v0.43.2] - 2019-11-14
### Bug Fixes 🐞
- remove support for v1.13.12, v1.14.8, v1.15.5, v1.16.2
- Change Azure CNI default version to match latest windows vhd ([#2271](https://github.com/Azure/aks-engine/issues/2271))
- only use vendor modules and run outside of GOPATH in e2e tests ([#2256](https://github.com/Azure/aks-engine/issues/2256))
- add go module export ([#2255](https://github.com/Azure/aks-engine/issues/2255))
- retain nsg rules during upgrade ([#2060](https://github.com/Azure/aks-engine/issues/2060))
- Ensured identity system passed from cli is propagated to api model ([#2231](https://github.com/Azure/aks-engine/issues/2231))

### Features 🌈
- Copy linux release notes out of VHD instead of parsing console output ([#2252](https://github.com/Azure/aks-engine/issues/2252))

### Maintenance 🔧
- replace dep by go mod ([#2165](https://github.com/Azure/aks-engine/issues/2165))
- log collection bugfix around some noexistant paths ([#2246](https://github.com/Azure/aks-engine/issues/2246))
- remove k8s 1.10 from VHD ([#2244](https://github.com/Azure/aks-engine/issues/2244))
- Adding a log collection script for Windows ([#2236](https://github.com/Azure/aks-engine/issues/2236))
- update to MS Moby 3.0.8 ([#2239](https://github.com/Azure/aks-engine/issues/2239))
- remove support for Kubernetes 1.10.x ([#2234](https://github.com/Azure/aks-engine/issues/2234))

### Testing 💚
- skip cloudprovider config test in no ssh scenario ([#2241](https://github.com/Azure/aks-engine/issues/2241))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.43.1"></a>
# [v0.43.1] - 2019-11-05
### Bug Fixes 🐞
- Change Azure CNI default version to match latest windows vhd ([#2271](https://github.com/Azure/aks-engine/issues/2271))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.43.0"></a>
# [v0.43.0] - 2019-10-29
### Bug Fixes 🐞
- Resolve incorrect Windows extension params ([#2201](https://github.com/Azure/aks-engine/issues/2201))
- go template and does not short circuit causing nil ptr ([#2217](https://github.com/Azure/aks-engine/issues/2217))
- default Translator if not specified to not panic ([#2211](https://github.com/Azure/aks-engine/issues/2211))
- move CSE endcustomscript log to very end of CSE ([#2194](https://github.com/Azure/aks-engine/issues/2194))
- run all CSE apt-get purge operations in background ([#2196](https://github.com/Azure/aks-engine/issues/2196))
- deploy CSI drivers for Azure Disk & Azure File when cloud-controller-manager is enabled ([#2166](https://github.com/Azure/aks-engine/issues/2166))
- CSE file wait is idempotent ([#2192](https://github.com/Azure/aks-engine/issues/2192))
- don't run apt operations on coreos ([#2189](https://github.com/Azure/aks-engine/issues/2189))
- Wait for management IP after creating ext network. ([#2049](https://github.com/Azure/aks-engine/issues/2049))
- actually use cloudprovider write rate limits ([#2184](https://github.com/Azure/aks-engine/issues/2184))
- ensure core components have system-node-critical priorityClass ([#2174](https://github.com/Azure/aks-engine/issues/2174))
- Update mirror probe port ([#1892](https://github.com/Azure/aks-engine/issues/1892))
- ensure cni conf json is valid when exception list contains a single address ([#2157](https://github.com/Azure/aks-engine/issues/2157))
- set curl's 'cacert' option to download pre-provision scripts on Azure Stack ([#2147](https://github.com/Azure/aks-engine/issues/2147))
- revert 'trust Azure Stack's CA cert before pre-provision extensions' ([#2135](https://github.com/Azure/aks-engine/issues/2135))
- process azure cni exception list for WS2019  ([#2125](https://github.com/Azure/aks-engine/issues/2125))
- don’t assign distro if image ref is provided ([#2055](https://github.com/Azure/aks-engine/issues/2055))
- restore original moby apt prod.list URL ([#2121](https://github.com/Azure/aks-engine/issues/2121))
- auditd improvements ([#2078](https://github.com/Azure/aks-engine/issues/2078))
- container monitoring add-on failure in the generate path ([#2109](https://github.com/Azure/aks-engine/issues/2109))
- add read perms for CSINodes to cluster-autoscaler role ([#2096](https://github.com/Azure/aks-engine/issues/2096))
- remove Azure agent specific artifacts from windows vhd ([#2088](https://github.com/Azure/aks-engine/issues/2088))
- cleaning up features.md ([#2090](https://github.com/Azure/aks-engine/issues/2090))
- Increasing inotify limit to 1048576 from the default 8192 ([#1801](https://github.com/Azure/aks-engine/issues/1801))
- updating windows release notes script to set a file encoding ([#2061](https://github.com/Azure/aks-engine/issues/2061))
- fixes cluster preprovisioning script download on azurestack. ([#2053](https://github.com/Azure/aks-engine/issues/2053))
- verify LCOW in makedev.ps1 ([#2037](https://github.com/Azure/aks-engine/issues/2037))
- Add nil pointer check for data disks ([#2024](https://github.com/Azure/aks-engine/issues/2024))
- lowercase vm name before calling drain node ([#2011](https://github.com/Azure/aks-engine/issues/2011))
- continue upgrading after node drain error ([#2013](https://github.com/Azure/aks-engine/issues/2013))
- cleanup apt artifacts before running apt ([#2006](https://github.com/Azure/aks-engine/issues/2006))
- check the nodes of all agentpool during upgrade ([#1893](https://github.com/Azure/aks-engine/issues/1893))

### Code Refactoring 💎
- Windows node restart cleanup script ([#2148](https://github.com/Azure/aks-engine/issues/2148))

### Continuous Integration 💜
- fix deployment cleanup in cleanup script ([#2132](https://github.com/Azure/aks-engine/issues/2132))
- Win CSE benchmark tests ([#2023](https://github.com/Azure/aks-engine/issues/2023))
- Adding provenance info in Windows VHD release notes ([#1990](https://github.com/Azure/aks-engine/issues/1990))
- consistent whitespace ([#1992](https://github.com/Azure/aks-engine/issues/1992))

### Documentation 📘
- tidying up ([#2126](https://github.com/Azure/aks-engine/issues/2126))
- correcting grammatical mistakes ([#2122](https://github.com/Azure/aks-engine/issues/2122))
- correcting slack channel name in docs ([#2119](https://github.com/Azure/aks-engine/issues/2119))
- Windows vhd doc ([#2070](https://github.com/Azure/aks-engine/issues/2070))
- supported versions for disconnected Azure Stacks using VHD 2019.09.19 ([#1984](https://github.com/Azure/aks-engine/issues/1984))
- changing example to use the aks base image and eliminate entry for k8s version ([#2017](https://github.com/Azure/aks-engine/issues/2017))

### Features 🌈
- Azure policy add-on ([#2173](https://github.com/Azure/aks-engine/issues/2173))
- blang/semver and gofrs/uuid does not match module compatibility ([#2204](https://github.com/Azure/aks-engine/issues/2204))
- enable containermonitoring addon support for Azurestack ([#2153](https://github.com/Azure/aks-engine/issues/2153))
- expose PlatformFaultDomainCount in the apimodel ([#2127](https://github.com/Azure/aks-engine/issues/2127))
- add support for Kubernetes v1.16.2 ([#2163](https://github.com/Azure/aks-engine/issues/2163))
- add support for Kubernetes v1.15.5 ([#2164](https://github.com/Azure/aks-engine/issues/2164))
- add support for Kubernetes 1.13.12 ([#2160](https://github.com/Azure/aks-engine/issues/2160))
- add support for Kubernetes v1.14.8 ([#2159](https://github.com/Azure/aks-engine/issues/2159))
- add linux support for kubenet + containerd ([#2050](https://github.com/Azure/aks-engine/issues/2050))
- enable re-use of VHD ([#2130](https://github.com/Azure/aks-engine/issues/2130))
- add support for Kubernetes 1.17.0-alpha.1 ([#2112](https://github.com/Azure/aks-engine/issues/2112))
- add support for Kubernetes v1.16.1 ([#2092](https://github.com/Azure/aks-engine/issues/2092))
- container monitoring addon supports other Azure clouds ([#2031](https://github.com/Azure/aks-engine/issues/2031))
- dualstack phase2 changes ([#1929](https://github.com/Azure/aks-engine/issues/1929))
- enable azurestack telemetry for scale and upgrade scenarios ([#2001](https://github.com/Azure/aks-engine/issues/2001))
- use aks windows vhds by default ([#2048](https://github.com/Azure/aks-engine/issues/2048))
- Skip installing docker if already installed on windows vhd ([#2052](https://github.com/Azure/aks-engine/issues/2052))
- Skip downloading of pause image if image already present in windows vhd ([#2040](https://github.com/Azure/aks-engine/issues/2040))
- Windows vhd artifact caching ([#1943](https://github.com/Azure/aks-engine/issues/1943))

### Maintenance 🔧
- rev VHD images to 2019.10.24 ([#2222](https://github.com/Azure/aks-engine/issues/2222))
- enable CSE ACR connectivity check only for AKS ([#2210](https://github.com/Azure/aks-engine/issues/2210))
- add Kubernetes patch releases to VHD ([#2200](https://github.com/Azure/aks-engine/issues/2200))
- update go toolchain to 1.12.12 ([#2193](https://github.com/Azure/aks-engine/issues/2193))
- default to ubuntu distro for customHyperkubeImage ([#2188](https://github.com/Azure/aks-engine/issues/2188))
- enable cloudproviderBackoff v2 for k8s >= 1.14 ([#2100](https://github.com/Azure/aks-engine/issues/2100))
- update cluster-autoscaler to latest patch releases ([#2185](https://github.com/Azure/aks-engine/issues/2185))
- document some bug workarounds ([#2103](https://github.com/Azure/aks-engine/issues/2103))
- Updating aks-engine to use oct windows aks vhd ([#2176](https://github.com/Azure/aks-engine/issues/2176))
- rev VHD images to 2019.10.15 ([#2172](https://github.com/Azure/aks-engine/issues/2172))
- Adding collectlogs.ps1 to windows AKS VHD ([#2156](https://github.com/Azure/aks-engine/issues/2156))
- add Kubernetes patch releases to VHD ([#2162](https://github.com/Azure/aks-engine/issues/2162))
- rev AKS VHDs to 2019.10.09 ([#2136](https://github.com/Azure/aks-engine/issues/2136))
- update to MS Moby 3.0.7 ([#2114](https://github.com/Azure/aks-engine/issues/2114))
- rev VHD images to 2019.10.03 ([#2116](https://github.com/Azure/aks-engine/issues/2116))
- move all apt commands to cse script ([#2115](https://github.com/Azure/aks-engine/issues/2115))
- add legal notice to vhd ([#2093](https://github.com/Azure/aks-engine/issues/2093))
- update default IPv6 subnet for dual stack, update docs ([#2095](https://github.com/Azure/aks-engine/issues/2095))
- Windows update to docker 19.03.2 ([#2071](https://github.com/Azure/aks-engine/issues/2071))
- update default tiller version to v2.13.1 ([#2075](https://github.com/Azure/aks-engine/issues/2075))
- update cluster-autoscaler to latest patch releases ([#2038](https://github.com/Azure/aks-engine/issues/2038))
- update go toolchain to 1.12.10 ([#2034](https://github.com/Azure/aks-engine/issues/2034))
- rev VHD to 2019.09.24 ([#2035](https://github.com/Azure/aks-engine/issues/2035))
- RBAC is required for Kubernetes >= 1.15.0 ([#2008](https://github.com/Azure/aks-engine/issues/2008))
- rev VHD images to 2019.09.19 ([#1998](https://github.com/Azure/aks-engine/issues/1998))
- add new k8s releases to VHD CI spec ([#1976](https://github.com/Azure/aks-engine/issues/1976))

### Revert Change ◀️
- Revert "test: retry az commands in E2E tests ([#2089](https://github.com/Azure/aks-engine/issues/2089))" ([#2214](https://github.com/Azure/aks-engine/issues/2214))

### Testing 💚
- remove soak test log statements with connection string ([#2230](https://github.com/Azure/aks-engine/issues/2230))
- add newer k8s versions to components version UT ([#2218](https://github.com/Azure/aks-engine/issues/2218))
- skip tests when block ssh port enabled ([#2213](https://github.com/Azure/aks-engine/issues/2213))
- cluster configs that don't work in 1.17 ([#2198](https://github.com/Azure/aks-engine/issues/2198))
- enable convenient customHyperkubeImage E2E tests ([#2197](https://github.com/Azure/aks-engine/issues/2197))
- add e2e test parameter for skipping all ssh tests ([#2187](https://github.com/Azure/aks-engine/issues/2187))
- build larger subnets for E2E custom VNET scenario ([#2179](https://github.com/Azure/aks-engine/issues/2179))
- make AfterSuite log output optional ([#2139](https://github.com/Azure/aks-engine/issues/2139))
- minimal Windows custom VNET cluster config ([#2144](https://github.com/Azure/aks-engine/issues/2144))
- Add 'latestReleasedVersion' target to Jenkins jobs ([#2133](https://github.com/Azure/aks-engine/issues/2133))
- skip time sync for 18.04-backed sgx test ([#2137](https://github.com/Azure/aks-engine/issues/2137))
- print logs after E2E tests ([#2131](https://github.com/Azure/aks-engine/issues/2131))
- Add api model entry for sgx ([#2105](https://github.com/Azure/aks-engine/issues/2105))
- faster fail during soak test re-build ([#2128](https://github.com/Azure/aks-engine/issues/2128))
- retry "get nodes that match regex" test ([#2120](https://github.com/Azure/aks-engine/issues/2120))
- get Running pods by prefix in most cases ([#2104](https://github.com/Azure/aks-engine/issues/2104))
- retry establishing ssh connection for e2e ([#2097](https://github.com/Azure/aks-engine/issues/2097))
- retry az commands in E2E tests ([#2089](https://github.com/Azure/aks-engine/issues/2089))
- retry “kubectl get config” ([#2085](https://github.com/Azure/aks-engine/issues/2085))
- retry “get pods by prefix substring match” ([#2083](https://github.com/Azure/aks-engine/issues/2083))
- don’t test for iis pod restart ([#2081](https://github.com/Azure/aks-engine/issues/2081))
- disable 1.16 tests for various cluster configs ([#2065](https://github.com/Azure/aks-engine/issues/2065))
- don’t timeout k8s resource deletes in E2E ([#2063](https://github.com/Azure/aks-engine/issues/2063))
- availability zones not available in all regions ([#2047](https://github.com/Azure/aks-engine/issues/2047))
- more resilient “wait for successful pod readiness checks” ([#2015](https://github.com/Azure/aks-engine/issues/2015))
- ensure we don’t dereference nil node List ([#2021](https://github.com/Azure/aks-engine/issues/2021))
- dns-liveness against more than one URL ([#2016](https://github.com/Azure/aks-engine/issues/2016))
- udpate not in vhd config, add ubuntu distro ([#2007](https://github.com/Azure/aks-engine/issues/2007))
- enable stability tests for PR E2E ([#2004](https://github.com/Azure/aks-engine/issues/2004))
- s/Standard_D2_3/Standard_D2_v3 ([#2005](https://github.com/Azure/aks-engine/issues/2005))
- address 4 E2E test flakes ([#1997](https://github.com/Azure/aks-engine/issues/1997))
- Use Standard_D2_v3 and low pri VMSS in E2E test configs ([#2000](https://github.com/Azure/aks-engine/issues/2000))
- retry remote ssh commands ([#1993](https://github.com/Azure/aks-engine/issues/1993))
- skip time sync test for 18.04 ([#1988](https://github.com/Azure/aks-engine/issues/1988))
- because where would we be without standardized orthography ([#1991](https://github.com/Azure/aks-engine/issues/1991))
- retry’able node get ([#1987](https://github.com/Azure/aks-engine/issues/1987))
- rationalize test configs ([#1985](https://github.com/Azure/aks-engine/issues/1985))
- run DNS validation against more than one query ([#1983](https://github.com/Azure/aks-engine/issues/1983))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.41.5"></a>
# [v0.41.5] - 2019-10-22
### Bug Fixes 🐞
- run all CSE apt-get purge operations in background ([#2196](https://github.com/Azure/aks-engine/issues/2196))
- CSE file wait is idempotent ([#2192](https://github.com/Azure/aks-engine/issues/2192))
- cleanup apt artifacts before running apt ([#2006](https://github.com/Azure/aks-engine/issues/2006))
- check the nodes of all agentpool during upgrade ([#1893](https://github.com/Azure/aks-engine/issues/1893))

### Features 🌈
- add support for Kubernetes v1.16.2 ([#2163](https://github.com/Azure/aks-engine/issues/2163))
- add support for Kubernetes v1.15.5 ([#2164](https://github.com/Azure/aks-engine/issues/2164))
- add support for Kubernetes 1.13.12 ([#2160](https://github.com/Azure/aks-engine/issues/2160))
- add support for Kubernetes v1.14.8 ([#2159](https://github.com/Azure/aks-engine/issues/2159))
- Skip downloading of pause image if image already present in windows vhd ([#2040](https://github.com/Azure/aks-engine/issues/2040))
- Windows vhd artifact caching ([#1943](https://github.com/Azure/aks-engine/issues/1943))
- add support for Kubernetes 1.17.0-alpha.1 ([#2112](https://github.com/Azure/aks-engine/issues/2112))
- add support for Kubernetes v1.16.1 ([#2092](https://github.com/Azure/aks-engine/issues/2092))

### Maintenance 🔧
- rev VHD images to 2019.10.15 ([#2172](https://github.com/Azure/aks-engine/issues/2172))
- add Kubernetes patch releases to VHD ([#2162](https://github.com/Azure/aks-engine/issues/2162))
- Windows update to docker 19.03.2 ([#2071](https://github.com/Azure/aks-engine/issues/2071))
- add new k8s releases to VHD CI spec ([#1976](https://github.com/Azure/aks-engine/issues/1976))
- rev VHD to 2019.09.24 ([#2035](https://github.com/Azure/aks-engine/issues/2035))
- rev VHD images to 2019.09.19 ([#1998](https://github.com/Azure/aks-engine/issues/1998))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.42.2"></a>
# [v0.42.2] - 2019-10-22
### Bug Fixes 🐞
- run all CSE apt-get purge operations in background ([#2196](https://github.com/Azure/aks-engine/issues/2196))
- CSE file wait is idempotent ([#2192](https://github.com/Azure/aks-engine/issues/2192))
- don't run apt operations on coreos ([#2189](https://github.com/Azure/aks-engine/issues/2189))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.42.1"></a>
# [v0.42.1] - 2019-10-14
### Bug Fixes 🐞
- ensure cni conf json is valid when exception list contains a single address ([#2157](https://github.com/Azure/aks-engine/issues/2157))
- process azure cni exception list for WS2019  ([#2125](https://github.com/Azure/aks-engine/issues/2125))

### Features 🌈
- add support for Kubernetes v1.16.2 ([#2163](https://github.com/Azure/aks-engine/issues/2163))
- add support for Kubernetes v1.15.5 ([#2164](https://github.com/Azure/aks-engine/issues/2164))
- add support for Kubernetes 1.13.12 ([#2160](https://github.com/Azure/aks-engine/issues/2160))
- add support for Kubernetes v1.14.8 ([#2159](https://github.com/Azure/aks-engine/issues/2159))

### Maintenance 🔧
- Updating aks-engine to use oct windows aks vhd ([#2176](https://github.com/Azure/aks-engine/issues/2176))
- rev VHD images to 2019.10.15 ([#2172](https://github.com/Azure/aks-engine/issues/2172))
- add Kubernetes patch releases to VHD ([#2162](https://github.com/Azure/aks-engine/issues/2162))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.42.0"></a>
# [v0.42.0] - 2019-10-10
### Bug Fixes 🐞
- revert 'trust Azure Stack's CA cert before pre-provision extensions' ([#2135](https://github.com/Azure/aks-engine/issues/2135))
- don’t assign distro if image ref is provided ([#2055](https://github.com/Azure/aks-engine/issues/2055))
- restore original moby apt prod.list URL ([#2121](https://github.com/Azure/aks-engine/issues/2121))
- auditd improvements ([#2078](https://github.com/Azure/aks-engine/issues/2078))
- container monitoring add-on failure in the generate path ([#2109](https://github.com/Azure/aks-engine/issues/2109))
- add read perms for CSINodes to cluster-autoscaler role ([#2096](https://github.com/Azure/aks-engine/issues/2096))
- remove Azure agent specific artifacts from windows vhd ([#2088](https://github.com/Azure/aks-engine/issues/2088))
- cleaning up features.md ([#2090](https://github.com/Azure/aks-engine/issues/2090))
- Increasing inotify limit to 1048576 from the default 8192 ([#1801](https://github.com/Azure/aks-engine/issues/1801))
- updating windows release notes script to set a file encoding ([#2061](https://github.com/Azure/aks-engine/issues/2061))
- fixes cluster preprovisioning script download on azurestack. ([#2053](https://github.com/Azure/aks-engine/issues/2053))
- verify LCOW in makedev.ps1 ([#2037](https://github.com/Azure/aks-engine/issues/2037))
- Add nil pointer check for data disks ([#2024](https://github.com/Azure/aks-engine/issues/2024))
- lowercase vm name before calling drain node ([#2011](https://github.com/Azure/aks-engine/issues/2011))
- continue upgrading after node drain error ([#2013](https://github.com/Azure/aks-engine/issues/2013))
- cleanup apt artifacts before running apt ([#2006](https://github.com/Azure/aks-engine/issues/2006))
- check the nodes of all agentpool during upgrade ([#1893](https://github.com/Azure/aks-engine/issues/1893))
- correct scale message after scale operations ([#1977](https://github.com/Azure/aks-engine/issues/1977))
- delete hyperkube images before starting kubelet ([#1816](https://github.com/Azure/aks-engine/issues/1816))
- not overwrite service principal profile for hosted master ([#1935](https://github.com/Azure/aks-engine/issues/1935))
- update the location of the VHD footprint file on VMs ([#1944](https://github.com/Azure/aks-engine/issues/1944))
- add hostname to CSE command so it shows in ARM deployment errors ([#1921](https://github.com/Azure/aks-engine/issues/1921))
- Remove the blind mkdir when writing outputs ([#1048](https://github.com/Azure/aks-engine/issues/1048)) ([#1889](https://github.com/Azure/aks-engine/issues/1889))
- add #EOF for dhcpv6 service ([#1909](https://github.com/Azure/aks-engine/issues/1909))
- Update rbac settings for cluster autoscaler ([#1582](https://github.com/Azure/aks-engine/issues/1582))
- add deprovision step to packer script ([#1865](https://github.com/Azure/aks-engine/issues/1865))
- set streaming-connection-idle-timeout as 4h ([#1870](https://github.com/Azure/aks-engine/issues/1870))
- improve container monitoring add-on ([#1686](https://github.com/Azure/aks-engine/issues/1686))
- more aks guarding against agentLbID template var ([#1874](https://github.com/Azure/aks-engine/issues/1874))
- guard AKS scenarios against aks-engine SLB template definitions ([#1872](https://github.com/Azure/aks-engine/issues/1872))
- fix customVMTags functionality and add more unit tests ([#1867](https://github.com/Azure/aks-engine/issues/1867))
- ensure master role assignment happens after VM deployment ([#1797](https://github.com/Azure/aks-engine/issues/1797))
- add Standard_F8 to accelerated networking whitelist ([#1858](https://github.com/Azure/aks-engine/issues/1858))
- retain custom nsg rules during upgrade ([#1792](https://github.com/Azure/aks-engine/issues/1792))
- Cannot set boolean property from command line ([#1848](https://github.com/Azure/aks-engine/issues/1848))
- restore working calico networkPolicy vlabs conversion logic ([#1855](https://github.com/Azure/aks-engine/issues/1855))

### Code Refactoring 💎
- configure kube-proxy 1.16 with a ConfigMap ([#1862](https://github.com/Azure/aks-engine/issues/1862))

### Continuous Integration 💜
- fix deployment cleanup in cleanup script ([#2132](https://github.com/Azure/aks-engine/issues/2132))
- Win CSE benchmark tests ([#2023](https://github.com/Azure/aks-engine/issues/2023))
- Adding provenance info in Windows VHD release notes ([#1990](https://github.com/Azure/aks-engine/issues/1990))
- consistent whitespace ([#1992](https://github.com/Azure/aks-engine/issues/1992))
- Running e2e tests as part of windows VHD build ([#1877](https://github.com/Azure/aks-engine/issues/1877))
- enable rg substring match in azure cleanup script ([#1832](https://github.com/Azure/aks-engine/issues/1832))
- add VHD rg cleanup script ([#1837](https://github.com/Azure/aks-engine/issues/1837))

### Documentation 📘
- tidying up ([#2126](https://github.com/Azure/aks-engine/issues/2126))
- correcting grammatical mistakes ([#2122](https://github.com/Azure/aks-engine/issues/2122))
- correcting slack channel name in docs ([#2119](https://github.com/Azure/aks-engine/issues/2119))
- Windows vhd doc ([#2070](https://github.com/Azure/aks-engine/issues/2070))
- supported versions for disconnected Azure Stacks using VHD 2019.09.19 ([#1984](https://github.com/Azure/aks-engine/issues/1984))
- changing example to use the aks base image and eliminate entry for k8s version ([#2017](https://github.com/Azure/aks-engine/issues/2017))
- add 1.16 api model example ([#1963](https://github.com/Azure/aks-engine/issues/1963))
- supported versions for disconnected Azure Stacks using VHD 2019.08.09 & 2019.08.21 ([#1907](https://github.com/Azure/aks-engine/issues/1907))
- fix invalid link ([#1876](https://github.com/Azure/aks-engine/issues/1876))

### Features 🌈
- enable re-use of VHD ([#2130](https://github.com/Azure/aks-engine/issues/2130))
- add support for Kubernetes 1.17.0-alpha.1 ([#2112](https://github.com/Azure/aks-engine/issues/2112))
- add support for Kubernetes v1.16.1 ([#2092](https://github.com/Azure/aks-engine/issues/2092))
- container monitoring addon supports other Azure clouds ([#2031](https://github.com/Azure/aks-engine/issues/2031))
- dualstack phase2 changes ([#1929](https://github.com/Azure/aks-engine/issues/1929))
- enable azurestack telemetry for scale and upgrade scenarios ([#2001](https://github.com/Azure/aks-engine/issues/2001))
- use aks windows vhds by default ([#2048](https://github.com/Azure/aks-engine/issues/2048))
- Skip installing docker if already installed on windows vhd ([#2052](https://github.com/Azure/aks-engine/issues/2052))
- Skip downloading of pause image if image already present in windows vhd ([#2040](https://github.com/Azure/aks-engine/issues/2040))
- Windows vhd artifact caching ([#1943](https://github.com/Azure/aks-engine/issues/1943))
- add support for Kubernetes 1.16.0 ([#1972](https://github.com/Azure/aks-engine/issues/1972))
- add support for Kubernetes 1.13.11 ([#1975](https://github.com/Azure/aks-engine/issues/1975))
- add support for Kubernetes 1.15.4 ([#1974](https://github.com/Azure/aks-engine/issues/1974))
- add support for Kubernetes 1.14.7 ([#1973](https://github.com/Azure/aks-engine/issues/1973))
- upgrade metrics server to v0.3.4 ([#1109](https://github.com/Azure/aks-engine/issues/1109))
- add new Norway regions ([#1939](https://github.com/Azure/aks-engine/issues/1939))
- enable deployment telemetry for generate command on Azure Stack ([#1847](https://github.com/Azure/aks-engine/issues/1847))
- add support for Kubernetes 1.16.0-rc.1 ([#1938](https://github.com/Azure/aks-engine/issues/1938))
- add support for Kubernetes 1.16.0-beta.2 ([#1906](https://github.com/Azure/aks-engine/issues/1906))
- add build and deploy windows zip file support on Azure Stack ([#1888](https://github.com/Azure/aks-engine/issues/1888))
- add new Azure Switzerland regions ([#1904](https://github.com/Azure/aks-engine/issues/1904))
- add germanynorth and germanywestcentral regions ([#1897](https://github.com/Azure/aks-engine/issues/1897))
- add new Azure VM SKUs ([#1896](https://github.com/Azure/aks-engine/issues/1896))
- update default Kubernetes version to 1.13 ([#1850](https://github.com/Azure/aks-engine/issues/1850))
- Windows image references ([#1718](https://github.com/Azure/aks-engine/issues/1718))
- Allow customization of coredns ([#1541](https://github.com/Azure/aks-engine/issues/1541)) ([#1841](https://github.com/Azure/aks-engine/issues/1841))

### Maintenance 🔧
- rev AKS VHDs to 2019.10.09 ([#2136](https://github.com/Azure/aks-engine/issues/2136))
- update to MS Moby 3.0.7 ([#2114](https://github.com/Azure/aks-engine/issues/2114))
- rev VHD images to 2019.10.03 ([#2116](https://github.com/Azure/aks-engine/issues/2116))
- move all apt commands to cse script ([#2115](https://github.com/Azure/aks-engine/issues/2115))
- add legal notice to vhd ([#2093](https://github.com/Azure/aks-engine/issues/2093))
- update default IPv6 subnet for dual stack, update docs ([#2095](https://github.com/Azure/aks-engine/issues/2095))
- Windows update to docker 19.03.2 ([#2071](https://github.com/Azure/aks-engine/issues/2071))
- update default tiller version to v2.13.1 ([#2075](https://github.com/Azure/aks-engine/issues/2075))
- update cluster-autoscaler to latest patch releases ([#2038](https://github.com/Azure/aks-engine/issues/2038))
- update go toolchain to 1.12.10 ([#2034](https://github.com/Azure/aks-engine/issues/2034))
- rev VHD to 2019.09.24 ([#2035](https://github.com/Azure/aks-engine/issues/2035))
- RBAC is required for Kubernetes >= 1.15.0 ([#2008](https://github.com/Azure/aks-engine/issues/2008))
- rev VHD images to 2019.09.19 ([#1998](https://github.com/Azure/aks-engine/issues/1998))
- add new k8s releases to VHD CI spec ([#1976](https://github.com/Azure/aks-engine/issues/1976))
- update ip-masq-agent to v2.5.0 ([#1908](https://github.com/Azure/aks-engine/issues/1908))
- rev VHD images to 2019.09.16 ([#1962](https://github.com/Azure/aks-engine/issues/1962))
- install metrics-server v0.3.4 via VHD ([#1959](https://github.com/Azure/aks-engine/issues/1959))
- rev vhd to 2019.09.13 ([#1955](https://github.com/Azure/aks-engine/issues/1955))
- fix closure capture for UPGRADE_VERSIONS ([#1953](https://github.com/Azure/aks-engine/issues/1953))
- inject all params into env vars for the build ([#1942](https://github.com/Azure/aks-engine/issues/1942))
- add upgrade and scale parallel jobs ([#1941](https://github.com/Azure/aks-engine/issues/1941))
- rev VHD image references to 2019.09.10 ([#1940](https://github.com/Azure/aks-engine/issues/1940))
- rev etcd to v3.3.15 ([#1931](https://github.com/Azure/aks-engine/issues/1931))
- inject log analytics workspace key into test env config ([#1934](https://github.com/Azure/aks-engine/issues/1934))
- Update Windows to version  17763.678.1908092216 ([#1933](https://github.com/Azure/aks-engine/issues/1933))
- bump keyvault-flexvol to v0.0.13 ([#1924](https://github.com/Azure/aks-engine/issues/1924))
- add job exclusion regex to jenkinsfile ([#1922](https://github.com/Azure/aks-engine/issues/1922))
- make the pki key size settable. ([#1891](https://github.com/Azure/aks-engine/issues/1891))
- add windows, monitoring and base test configs ([#1899](https://github.com/Azure/aks-engine/issues/1899))
- remove oms extension docs since they are obsolete ([#1900](https://github.com/Azure/aks-engine/issues/1900))
- disable tiller addon by default ([#1884](https://github.com/Azure/aks-engine/issues/1884))
- update CoreDNS to 1.6.2 ([#1880](https://github.com/Azure/aks-engine/issues/1880))
- create parallel Jenkins pipeline for e2e tests ([#1875](https://github.com/Azure/aks-engine/issues/1875))
- add apache2-utils as an apt dependency ([#1822](https://github.com/Azure/aks-engine/issues/1822))
- remove make target w/ non-existent file reference ([#1861](https://github.com/Azure/aks-engine/issues/1861))
- Targeting August updates for Windows VHD ([#1846](https://github.com/Azure/aks-engine/issues/1846))
- rev VHD image references to 2019.08.21 ([#1849](https://github.com/Azure/aks-engine/issues/1849))
- tolerate lowercase LB SKU vals ([#1838](https://github.com/Azure/aks-engine/issues/1838))
- bump keyvault-flexvol to v0.0.12 ([#1831](https://github.com/Azure/aks-engine/issues/1831))
- update az go sdk to v32.5 and autorest v13 ([#1793](https://github.com/Azure/aks-engine/issues/1793))

### Testing 💚
- skip time sync for 18.04-backed sgx test ([#2137](https://github.com/Azure/aks-engine/issues/2137))
- print logs after E2E tests ([#2131](https://github.com/Azure/aks-engine/issues/2131))
- Add api model entry for sgx ([#2105](https://github.com/Azure/aks-engine/issues/2105))
- faster fail during soak test re-build ([#2128](https://github.com/Azure/aks-engine/issues/2128))
- retry "get nodes that match regex" test ([#2120](https://github.com/Azure/aks-engine/issues/2120))
- get Running pods by prefix in most cases ([#2104](https://github.com/Azure/aks-engine/issues/2104))
- retry establishing ssh connection for e2e ([#2097](https://github.com/Azure/aks-engine/issues/2097))
- retry az commands in E2E tests ([#2089](https://github.com/Azure/aks-engine/issues/2089))
- retry “kubectl get config” ([#2085](https://github.com/Azure/aks-engine/issues/2085))
- retry “get pods by prefix substring match” ([#2083](https://github.com/Azure/aks-engine/issues/2083))
- don’t test for iis pod restart ([#2081](https://github.com/Azure/aks-engine/issues/2081))
- disable 1.16 tests for various cluster configs ([#2065](https://github.com/Azure/aks-engine/issues/2065))
- don’t timeout k8s resource deletes in E2E ([#2063](https://github.com/Azure/aks-engine/issues/2063))
- availability zones not available in all regions ([#2047](https://github.com/Azure/aks-engine/issues/2047))
- more resilient “wait for successful pod readiness checks” ([#2015](https://github.com/Azure/aks-engine/issues/2015))
- ensure we don’t dereference nil node List ([#2021](https://github.com/Azure/aks-engine/issues/2021))
- dns-liveness against more than one URL ([#2016](https://github.com/Azure/aks-engine/issues/2016))
- udpate not in vhd config, add ubuntu distro ([#2007](https://github.com/Azure/aks-engine/issues/2007))
- enable stability tests for PR E2E ([#2004](https://github.com/Azure/aks-engine/issues/2004))
- s/Standard_D2_3/Standard_D2_v3 ([#2005](https://github.com/Azure/aks-engine/issues/2005))
- address 4 E2E test flakes ([#1997](https://github.com/Azure/aks-engine/issues/1997))
- Use Standard_D2_v3 and low pri VMSS in E2E test configs ([#2000](https://github.com/Azure/aks-engine/issues/2000))
- retry remote ssh commands ([#1993](https://github.com/Azure/aks-engine/issues/1993))
- skip time sync test for 18.04 ([#1988](https://github.com/Azure/aks-engine/issues/1988))
- because where would we be without standardized orthography ([#1991](https://github.com/Azure/aks-engine/issues/1991))
- retry’able node get ([#1987](https://github.com/Azure/aks-engine/issues/1987))
- rationalize test configs ([#1985](https://github.com/Azure/aks-engine/issues/1985))
- run DNS validation against more than one query ([#1983](https://github.com/Azure/aks-engine/issues/1983))
- run gpu tests in westus2 ([#1965](https://github.com/Azure/aks-engine/issues/1965))
- fix "run multiple commands in succession" implementation ([#1956](https://github.com/Azure/aks-engine/issues/1956))
- pod WaitOnSucceeded against distinct pod ([#1950](https://github.com/Azure/aks-engine/issues/1950))
- s/southcentralus/uksouth ([#1949](https://github.com/Azure/aks-engine/issues/1949))
- more resilience when k get deployment ([#1948](https://github.com/Azure/aks-engine/issues/1948))
- az storage file download-batch --destination must exist ([#1947](https://github.com/Azure/aks-engine/issues/1947))
- crash docker and validate against common node count ([#1946](https://github.com/Azure/aks-engine/issues/1946))
- E2E: goroutine and err response cleanup ([#1925](https://github.com/Azure/aks-engine/issues/1925))
- printing wrong err ([#1881](https://github.com/Azure/aks-engine/issues/1881))
- retry pod get ([#1879](https://github.com/Azure/aks-engine/issues/1879))
- remove obsolete e2e test scripts ([#1768](https://github.com/Azure/aks-engine/issues/1768))
- improve E2E node readiness tests ([#1859](https://github.com/Azure/aks-engine/issues/1859))
- check DNS before scp’ing ([#1857](https://github.com/Azure/aks-engine/issues/1857))
- e2e accommodations for low-pri vmss configurations ([#1854](https://github.com/Azure/aks-engine/issues/1854))
- improve dashboard e2e tests ([#1853](https://github.com/Azure/aks-engine/issues/1853))
- allow 20 retries for ssh-dependent tests ([#1852](https://github.com/Azure/aks-engine/issues/1852))
- add retries to host OS DNS E2E tests ([#1843](https://github.com/Azure/aks-engine/issues/1843))
- wait for pod readiness in port forward test ([#1845](https://github.com/Azure/aks-engine/issues/1845))
- only check for static n node count if not using low pri VMSS ([#1844](https://github.com/Azure/aks-engine/issues/1844))
- retry all ssh-dependent E2E tests ([#1825](https://github.com/Azure/aks-engine/issues/1825))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.40.2"></a>
# [v0.40.2] - 2019-09-26
### Bug Fixes 🐞
- not overwrite service principal profile for hosted master ([#1935](https://github.com/Azure/aks-engine/issues/1935))
- cleanup apt artifacts before running apt ([#2006](https://github.com/Azure/aks-engine/issues/2006))
- more aks guarding against agentLbID template var ([#1874](https://github.com/Azure/aks-engine/issues/1874))
- guard AKS scenarios against aks-engine SLB template definitions ([#1872](https://github.com/Azure/aks-engine/issues/1872))
- fix customVMTags functionality and add more unit tests ([#1867](https://github.com/Azure/aks-engine/issues/1867))
- ensure master role assignment happens after VM deployment ([#1797](https://github.com/Azure/aks-engine/issues/1797))
- add Standard_F8 to accelerated networking whitelist ([#1858](https://github.com/Azure/aks-engine/issues/1858))
- restore working calico networkPolicy vlabs conversion logic ([#1855](https://github.com/Azure/aks-engine/issues/1855))

### Code Refactoring 💎
- configure kube-proxy 1.16 with a ConfigMap ([#1862](https://github.com/Azure/aks-engine/issues/1862))

### Maintenance 🔧
- update SKU reference to 1604-201909
- rev VHD to 2019.09.25
- rev VHD image references to 2019.08.21 ([#1849](https://github.com/Azure/aks-engine/issues/1849))
- tolerate lowercase LB SKU vals ([#1838](https://github.com/Azure/aks-engine/issues/1838))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.41.4"></a>
# [v0.41.4] - 2019-09-26
### Maintenance 🔧
- rev VHD to 2019.09.24 ([#2035](https://github.com/Azure/aks-engine/issues/2035))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.41.3"></a>
# [v0.41.3] - 2019-09-24
### Bug Fixes 🐞
- cleanup apt artifacts before running apt ([#2006](https://github.com/Azure/aks-engine/issues/2006))
- check the nodes of all agentpool during upgrade ([#1893](https://github.com/Azure/aks-engine/issues/1893))
- correct scale message after scale operations ([#1977](https://github.com/Azure/aks-engine/issues/1977))
- delete hyperkube images before starting kubelet ([#1816](https://github.com/Azure/aks-engine/issues/1816))
- not overwrite service principal profile for hosted master ([#1935](https://github.com/Azure/aks-engine/issues/1935))
- update the location of the VHD footprint file on VMs ([#1944](https://github.com/Azure/aks-engine/issues/1944))
- add hostname to CSE command so it shows in ARM deployment errors ([#1921](https://github.com/Azure/aks-engine/issues/1921))
- Remove the blind mkdir when writing outputs ([#1048](https://github.com/Azure/aks-engine/issues/1048)) ([#1889](https://github.com/Azure/aks-engine/issues/1889))
- add #EOF for dhcpv6 service ([#1909](https://github.com/Azure/aks-engine/issues/1909))
- Update rbac settings for cluster autoscaler ([#1582](https://github.com/Azure/aks-engine/issues/1582))
- add deprovision step to packer script ([#1865](https://github.com/Azure/aks-engine/issues/1865))
- set streaming-connection-idle-timeout as 4h ([#1870](https://github.com/Azure/aks-engine/issues/1870))
- improve container monitoring add-on ([#1686](https://github.com/Azure/aks-engine/issues/1686))
- more aks guarding against agentLbID template var ([#1874](https://github.com/Azure/aks-engine/issues/1874))
- guard AKS scenarios against aks-engine SLB template definitions ([#1872](https://github.com/Azure/aks-engine/issues/1872))
- fix customVMTags functionality and add more unit tests ([#1867](https://github.com/Azure/aks-engine/issues/1867))
- ensure master role assignment happens after VM deployment ([#1797](https://github.com/Azure/aks-engine/issues/1797))
- add Standard_F8 to accelerated networking whitelist ([#1858](https://github.com/Azure/aks-engine/issues/1858))
- retain custom nsg rules during upgrade ([#1792](https://github.com/Azure/aks-engine/issues/1792))
- Cannot set boolean property from command line ([#1848](https://github.com/Azure/aks-engine/issues/1848))
- restore working calico networkPolicy vlabs conversion logic ([#1855](https://github.com/Azure/aks-engine/issues/1855))

### Code Refactoring 💎
- configure kube-proxy 1.16 with a ConfigMap ([#1862](https://github.com/Azure/aks-engine/issues/1862))

### Continuous Integration 💜
- Running e2e tests as part of windows VHD build ([#1877](https://github.com/Azure/aks-engine/issues/1877))
- enable rg substring match in azure cleanup script ([#1832](https://github.com/Azure/aks-engine/issues/1832))
- add VHD rg cleanup script ([#1837](https://github.com/Azure/aks-engine/issues/1837))

### Documentation 📘
- add 1.16 api model example ([#1963](https://github.com/Azure/aks-engine/issues/1963))
- supported versions for disconnected Azure Stacks using VHD 2019.08.09 & 2019.08.21 ([#1907](https://github.com/Azure/aks-engine/issues/1907))
- fix invalid link ([#1876](https://github.com/Azure/aks-engine/issues/1876))

### Features 🌈
- add support for Kubernetes 1.16.0 ([#1972](https://github.com/Azure/aks-engine/issues/1972))
- add support for Kubernetes 1.13.11 ([#1975](https://github.com/Azure/aks-engine/issues/1975))
- add support for Kubernetes 1.15.4 ([#1974](https://github.com/Azure/aks-engine/issues/1974))
- add support for Kubernetes 1.14.7 ([#1973](https://github.com/Azure/aks-engine/issues/1973))
- upgrade metrics server to v0.3.4 ([#1109](https://github.com/Azure/aks-engine/issues/1109))
- add new Norway regions ([#1939](https://github.com/Azure/aks-engine/issues/1939))
- enable deployment telemetry for generate command on Azure Stack ([#1847](https://github.com/Azure/aks-engine/issues/1847))
- add support for Kubernetes 1.16.0-rc.1 ([#1938](https://github.com/Azure/aks-engine/issues/1938))
- add support for Kubernetes 1.16.0-beta.2 ([#1906](https://github.com/Azure/aks-engine/issues/1906))
- add build and deploy windows zip file support on Azure Stack ([#1888](https://github.com/Azure/aks-engine/issues/1888))
- add new Azure Switzerland regions ([#1904](https://github.com/Azure/aks-engine/issues/1904))
- add germanynorth and germanywestcentral regions ([#1897](https://github.com/Azure/aks-engine/issues/1897))
- add new Azure VM SKUs ([#1896](https://github.com/Azure/aks-engine/issues/1896))
- update default Kubernetes version to 1.13 ([#1850](https://github.com/Azure/aks-engine/issues/1850))
- Windows image references ([#1718](https://github.com/Azure/aks-engine/issues/1718))
- Allow customization of coredns ([#1541](https://github.com/Azure/aks-engine/issues/1541)) ([#1841](https://github.com/Azure/aks-engine/issues/1841))

### Maintenance 🔧
- rev VHD images to 2019.09.19 ([#1998](https://github.com/Azure/aks-engine/issues/1998))
- update ip-masq-agent to v2.5.0 ([#1908](https://github.com/Azure/aks-engine/issues/1908))
- rev VHD images to 2019.09.16 ([#1962](https://github.com/Azure/aks-engine/issues/1962))
- install metrics-server v0.3.4 via VHD ([#1959](https://github.com/Azure/aks-engine/issues/1959))
- rev vhd to 2019.09.13 ([#1955](https://github.com/Azure/aks-engine/issues/1955))
- fix closure capture for UPGRADE_VERSIONS ([#1953](https://github.com/Azure/aks-engine/issues/1953))
- inject all params into env vars for the build ([#1942](https://github.com/Azure/aks-engine/issues/1942))
- add upgrade and scale parallel jobs ([#1941](https://github.com/Azure/aks-engine/issues/1941))
- rev VHD image references to 2019.09.10 ([#1940](https://github.com/Azure/aks-engine/issues/1940))
- rev etcd to v3.3.15 ([#1931](https://github.com/Azure/aks-engine/issues/1931))
- inject log analytics workspace key into test env config ([#1934](https://github.com/Azure/aks-engine/issues/1934))
- Update Windows to version  17763.678.1908092216 ([#1933](https://github.com/Azure/aks-engine/issues/1933))
- bump keyvault-flexvol to v0.0.13 ([#1924](https://github.com/Azure/aks-engine/issues/1924))
- add job exclusion regex to jenkinsfile ([#1922](https://github.com/Azure/aks-engine/issues/1922))
- make the pki key size settable. ([#1891](https://github.com/Azure/aks-engine/issues/1891))
- add windows, monitoring and base test configs ([#1899](https://github.com/Azure/aks-engine/issues/1899))
- remove oms extension docs since they are obsolete ([#1900](https://github.com/Azure/aks-engine/issues/1900))
- disable tiller addon by default ([#1884](https://github.com/Azure/aks-engine/issues/1884))
- update CoreDNS to 1.6.2 ([#1880](https://github.com/Azure/aks-engine/issues/1880))
- create parallel Jenkins pipeline for e2e tests ([#1875](https://github.com/Azure/aks-engine/issues/1875))
- add apache2-utils as an apt dependency ([#1822](https://github.com/Azure/aks-engine/issues/1822))
- remove make target w/ non-existent file reference ([#1861](https://github.com/Azure/aks-engine/issues/1861))
- Targeting August updates for Windows VHD ([#1846](https://github.com/Azure/aks-engine/issues/1846))
- rev VHD image references to 2019.08.21 ([#1849](https://github.com/Azure/aks-engine/issues/1849))
- tolerate lowercase LB SKU vals ([#1838](https://github.com/Azure/aks-engine/issues/1838))
- bump keyvault-flexvol to v0.0.12 ([#1831](https://github.com/Azure/aks-engine/issues/1831))
- update az go sdk to v32.5 and autorest v13 ([#1793](https://github.com/Azure/aks-engine/issues/1793))

### Testing 💚
- run gpu tests in westus2 ([#1965](https://github.com/Azure/aks-engine/issues/1965))
- fix "run multiple commands in succession" implementation ([#1956](https://github.com/Azure/aks-engine/issues/1956))
- pod WaitOnSucceeded against distinct pod ([#1950](https://github.com/Azure/aks-engine/issues/1950))
- s/southcentralus/uksouth ([#1949](https://github.com/Azure/aks-engine/issues/1949))
- more resilience when k get deployment ([#1948](https://github.com/Azure/aks-engine/issues/1948))
- az storage file download-batch --destination must exist ([#1947](https://github.com/Azure/aks-engine/issues/1947))
- crash docker and validate against common node count ([#1946](https://github.com/Azure/aks-engine/issues/1946))
- E2E: goroutine and err response cleanup ([#1925](https://github.com/Azure/aks-engine/issues/1925))
- printing wrong err ([#1881](https://github.com/Azure/aks-engine/issues/1881))
- retry pod get ([#1879](https://github.com/Azure/aks-engine/issues/1879))
- remove obsolete e2e test scripts ([#1768](https://github.com/Azure/aks-engine/issues/1768))
- improve E2E node readiness tests ([#1859](https://github.com/Azure/aks-engine/issues/1859))
- check DNS before scp’ing ([#1857](https://github.com/Azure/aks-engine/issues/1857))
- e2e accommodations for low-pri vmss configurations ([#1854](https://github.com/Azure/aks-engine/issues/1854))
- improve dashboard e2e tests ([#1853](https://github.com/Azure/aks-engine/issues/1853))
- allow 20 retries for ssh-dependent tests ([#1852](https://github.com/Azure/aks-engine/issues/1852))
- add retries to host OS DNS E2E tests ([#1843](https://github.com/Azure/aks-engine/issues/1843))
- wait for pod readiness in port forward test ([#1845](https://github.com/Azure/aks-engine/issues/1845))
- only check for static n node count if not using low pri VMSS ([#1844](https://github.com/Azure/aks-engine/issues/1844))
- retry all ssh-dependent E2E tests ([#1825](https://github.com/Azure/aks-engine/issues/1825))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.40.1"></a>
# [v0.40.1] - 2019-09-20
### Bug Fixes 🐞
- cleanup apt artifacts before running apt ([#2006](https://github.com/Azure/aks-engine/issues/2006))
- more aks guarding against agentLbID template var ([#1874](https://github.com/Azure/aks-engine/issues/1874))
- guard AKS scenarios against aks-engine SLB template definitions ([#1872](https://github.com/Azure/aks-engine/issues/1872))
- fix customVMTags functionality and add more unit tests ([#1867](https://github.com/Azure/aks-engine/issues/1867))
- ensure master role assignment happens after VM deployment ([#1797](https://github.com/Azure/aks-engine/issues/1797))
- add Standard_F8 to accelerated networking whitelist ([#1858](https://github.com/Azure/aks-engine/issues/1858))
- restore working calico networkPolicy vlabs conversion logic ([#1855](https://github.com/Azure/aks-engine/issues/1855))

### Code Refactoring 💎
- configure kube-proxy 1.16 with a ConfigMap ([#1862](https://github.com/Azure/aks-engine/issues/1862))

### Maintenance 🔧
- rev VHD image references to 2019.08.21 ([#1849](https://github.com/Azure/aks-engine/issues/1849))
- tolerate lowercase LB SKU vals ([#1838](https://github.com/Azure/aks-engine/issues/1838))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.41.2"></a>
# [v0.41.2] - 2019-09-20
### Bug Fixes 🐞
- cleanup apt artifacts before running apt ([#2006](https://github.com/Azure/aks-engine/issues/2006))
- check the nodes of all agentpool during upgrade ([#1893](https://github.com/Azure/aks-engine/issues/1893))
- correct scale message after scale operations ([#1977](https://github.com/Azure/aks-engine/issues/1977))
- delete hyperkube images before starting kubelet ([#1816](https://github.com/Azure/aks-engine/issues/1816))
- not overwrite service principal profile for hosted master ([#1935](https://github.com/Azure/aks-engine/issues/1935))
- update the location of the VHD footprint file on VMs ([#1944](https://github.com/Azure/aks-engine/issues/1944))
- add hostname to CSE command so it shows in ARM deployment errors ([#1921](https://github.com/Azure/aks-engine/issues/1921))
- Remove the blind mkdir when writing outputs ([#1048](https://github.com/Azure/aks-engine/issues/1048)) ([#1889](https://github.com/Azure/aks-engine/issues/1889))
- add #EOF for dhcpv6 service ([#1909](https://github.com/Azure/aks-engine/issues/1909))
- Update rbac settings for cluster autoscaler ([#1582](https://github.com/Azure/aks-engine/issues/1582))
- add deprovision step to packer script ([#1865](https://github.com/Azure/aks-engine/issues/1865))
- set streaming-connection-idle-timeout as 4h ([#1870](https://github.com/Azure/aks-engine/issues/1870))
- improve container monitoring add-on ([#1686](https://github.com/Azure/aks-engine/issues/1686))
- more aks guarding against agentLbID template var ([#1874](https://github.com/Azure/aks-engine/issues/1874))
- guard AKS scenarios against aks-engine SLB template definitions ([#1872](https://github.com/Azure/aks-engine/issues/1872))
- fix customVMTags functionality and add more unit tests ([#1867](https://github.com/Azure/aks-engine/issues/1867))
- ensure master role assignment happens after VM deployment ([#1797](https://github.com/Azure/aks-engine/issues/1797))
- add Standard_F8 to accelerated networking whitelist ([#1858](https://github.com/Azure/aks-engine/issues/1858))
- retain custom nsg rules during upgrade ([#1792](https://github.com/Azure/aks-engine/issues/1792))
- Cannot set boolean property from command line ([#1848](https://github.com/Azure/aks-engine/issues/1848))
- restore working calico networkPolicy vlabs conversion logic ([#1855](https://github.com/Azure/aks-engine/issues/1855))
- skip creating slb for hostedmaster ([#1835](https://github.com/Azure/aks-engine/issues/1835))
- networkplugin conversion only if networkpolicy is empty ([#1823](https://github.com/Azure/aks-engine/issues/1823))
- use azcopy-preview in VHD pipeline ([#1821](https://github.com/Azure/aks-engine/issues/1821))
- metrics server cluster role ([#1714](https://github.com/Azure/aks-engine/issues/1714))
- kube-proxy addon not critical, doesn't reconcile ([#1814](https://github.com/Azure/aks-engine/issues/1814))
- set kubelet config for containerd in templates ([#1785](https://github.com/Azure/aks-engine/issues/1785))
- race condition which causes concurrent map writes during tests ([#1791](https://github.com/Azure/aks-engine/issues/1791))
- label scheduled-maintenance manifests for addon-manager ([#1755](https://github.com/Azure/aks-engine/issues/1755))
- address all cloud-init file waits ([#1719](https://github.com/Azure/aks-engine/issues/1719))
- re-add Promo SKUs to Azure constants ([#1752](https://github.com/Azure/aks-engine/issues/1752))
- omit NSG rules for Upgrade scenario ([#1705](https://github.com/Azure/aks-engine/issues/1705))
- "dpkg configure -a" interactive config update issue
- Change win-cni to win-bridge binaries and config ([#643](https://github.com/Azure/aks-engine/issues/643))
- document registry container image pull in VHD release notes ([#1699](https://github.com/Azure/aks-engine/issues/1699))
- lock docker distribution version for Azure Stack ([#1688](https://github.com/Azure/aks-engine/issues/1688))
- omit NSG rules created by LoadBalancer during upgrade ([#1646](https://github.com/Azure/aks-engine/issues/1646))
- honour CLEANUP_IF_FAIL in test runner ([#1680](https://github.com/Azure/aks-engine/issues/1680))
- only include OrchestratorProfile validation during data marshalling for create ([#1661](https://github.com/Azure/aks-engine/issues/1661))
- add back-compat support for deprecated distro vals ([#1669](https://github.com/Azure/aks-engine/issues/1669))
- distinct component update messages for upgrade/scale ([#1666](https://github.com/Azure/aks-engine/issues/1666))
- add ipv6 nic config for vmss [IPv6DualStack] ([#1648](https://github.com/Azure/aks-engine/issues/1648))
- clusters with SLB scale and upgrade ([#1622](https://github.com/Azure/aks-engine/issues/1622))
- update container monitoring add-on to use  latest version of the containerized omsagent ([#1637](https://github.com/Azure/aks-engine/issues/1637))
- extend default value -1 for PodMaxPids for master and agentpool profile … ([#1623](https://github.com/Azure/aks-engine/issues/1623))
- update cluster-autoscaler role ([#1643](https://github.com/Azure/aks-engine/issues/1643))
- Remove unnecessary transform on VMSS upgrade ([#1633](https://github.com/Azure/aks-engine/issues/1633))
- CoreOS improvements and fixes ([#1578](https://github.com/Azure/aks-engine/issues/1578))
- Remove pod-critical annotations in 1.16 manifests. ([#1621](https://github.com/Azure/aks-engine/issues/1621))
- old VMSS AKS cluster cannot be upgraded after April 2 ([#1561](https://github.com/Azure/aks-engine/issues/1561))
- remove availability sets in scaling up template ([#1610](https://github.com/Azure/aks-engine/issues/1610))
- use explicit spec.Selector with apps/v1 api manifests ([#1599](https://github.com/Azure/aks-engine/issues/1599))
- improve clear containers clean up ([#1600](https://github.com/Azure/aks-engine/issues/1600))
- do not install clear containers runtime on the VHD ([#1587](https://github.com/Azure/aks-engine/issues/1587))
- kubeconfig incorrectly escape characters for azuredeploy.json ([#1574](https://github.com/Azure/aks-engine/issues/1574))

### Code Refactoring 💎
- configure kube-proxy 1.16 with a ConfigMap ([#1862](https://github.com/Azure/aks-engine/issues/1862))
- simplify scale operation CLI output ([#1345](https://github.com/Azure/aks-engine/issues/1345))

### Code Style 🎶
- read without -r will mangle backslashes ([#1790](https://github.com/Azure/aks-engine/issues/1790))

### Continuous Integration 💜
- Running e2e tests as part of windows VHD build ([#1877](https://github.com/Azure/aks-engine/issues/1877))
- enable rg substring match in azure cleanup script ([#1832](https://github.com/Azure/aks-engine/issues/1832))
- add VHD rg cleanup script ([#1837](https://github.com/Azure/aks-engine/issues/1837))
- fix VHD release notes Azure DevOps log parsing ([#1702](https://github.com/Azure/aks-engine/issues/1702))

### Documentation 📘
- fix typo in examples/addons/cluster-autoscaler ([#1604](https://github.com/Azure/aks-engine/issues/1604))
- supported versions for disconnected Azure Stacks using VHD 2019.08.09 & 2019.08.21 ([#1907](https://github.com/Azure/aks-engine/issues/1907))
- fix invalid link ([#1876](https://github.com/Azure/aks-engine/issues/1876))
- gofish updates are automated now for AKS Engine releases ([#1818](https://github.com/Azure/aks-engine/issues/1818))
- fix the feature table ([#1796](https://github.com/Azure/aks-engine/issues/1796))
- Add devigned as owner ([#1784](https://github.com/Azure/aks-engine/issues/1784))
- increase master node count in Azure Stack template ([#1767](https://github.com/Azure/aks-engine/issues/1767))
- add 1.16 api model example ([#1963](https://github.com/Azure/aks-engine/issues/1963))
- clean up IPv4/v6 dual stack README ([#1690](https://github.com/Azure/aks-engine/issues/1690))
- supported versions for disconnected Azure Stacks using VHD 2019.07.10 ([#1685](https://github.com/Azure/aks-engine/issues/1685))
- Update ipv6 cidr for dualstack ([#1683](https://github.com/Azure/aks-engine/issues/1683))
- use Microsoft's standard code of conduct for GitHub ([#1659](https://github.com/Azure/aks-engine/issues/1659))
- rename vhd release notes files to match sku names ([#1608](https://github.com/Azure/aks-engine/issues/1608))
- **dual-stack:** Add prerequisites ([#1692](https://github.com/Azure/aks-engine/issues/1692))

### Features 🌈
- add support for Kubernetes 1.16.0 ([#1972](https://github.com/Azure/aks-engine/issues/1972))
- add support for Kubernetes 1.13.11 ([#1975](https://github.com/Azure/aks-engine/issues/1975))
- add support for Kubernetes 1.15.4 ([#1974](https://github.com/Azure/aks-engine/issues/1974))
- add support for Kubernetes 1.14.7 ([#1973](https://github.com/Azure/aks-engine/issues/1973))
- upgrade metrics server to v0.3.4 ([#1109](https://github.com/Azure/aks-engine/issues/1109))
- add new Norway regions ([#1939](https://github.com/Azure/aks-engine/issues/1939))
- enable deployment telemetry for generate command on Azure Stack ([#1847](https://github.com/Azure/aks-engine/issues/1847))
- add support for Kubernetes 1.16.0-rc.1 ([#1938](https://github.com/Azure/aks-engine/issues/1938))
- add support for Kubernetes 1.16.0-beta.2 ([#1906](https://github.com/Azure/aks-engine/issues/1906))
- add build and deploy windows zip file support on Azure Stack ([#1888](https://github.com/Azure/aks-engine/issues/1888))
- add new Azure Switzerland regions ([#1904](https://github.com/Azure/aks-engine/issues/1904))
- add germanynorth and germanywestcentral regions ([#1897](https://github.com/Azure/aks-engine/issues/1897))
- add new Azure VM SKUs ([#1896](https://github.com/Azure/aks-engine/issues/1896))
- update default Kubernetes version to 1.13 ([#1850](https://github.com/Azure/aks-engine/issues/1850))
- Windows image references ([#1718](https://github.com/Azure/aks-engine/issues/1718))
- Allow customization of coredns ([#1541](https://github.com/Azure/aks-engine/issues/1541)) ([#1841](https://github.com/Azure/aks-engine/issues/1841))
- add support for Kubernetes 1.16.0-beta.1 ([#1819](https://github.com/Azure/aks-engine/issues/1819))
- update cloud provider rate limit for vmas. ([#1808](https://github.com/Azure/aks-engine/issues/1808))
- autofill client Id and client secret for generate command ([#1766](https://github.com/Azure/aks-engine/issues/1766))
- add support for Kubernetes 1.15.3 ([#1805](https://github.com/Azure/aks-engine/issues/1805))
- add support for Kubernetes 1.14.6 ([#1804](https://github.com/Azure/aks-engine/issues/1804))
- add support for Kubernetes 1.13.10 ([#1803](https://github.com/Azure/aks-engine/issues/1803))
- Windows VHD build pipeline ([#1762](https://github.com/Azure/aks-engine/issues/1762))
- enable configurable cloudprovider ratelimits for write ([#1783](https://github.com/Azure/aks-engine/issues/1783))
- enable AKS VHD image in Fairfax ([#1770](https://github.com/Azure/aks-engine/issues/1770))
- add support for Kubernetes 1.16.0-alpha.3 ([#1745](https://github.com/Azure/aks-engine/issues/1745))
- add support for Kubernetes 1.13.9, 1.14.5, 1.15.2 on Azure Stack ([#1740](https://github.com/Azure/aks-engine/issues/1740))
- add support for Kubernetes 1.14.5 ([#1731](https://github.com/Azure/aks-engine/issues/1731))
- add support for Kubernetes 1.13.9 ([#1732](https://github.com/Azure/aks-engine/issues/1732))
- add support for Kubernetes 1.15.2 ([#1730](https://github.com/Azure/aks-engine/issues/1730))
- enable smart cloudprovider rate limiting ([#1693](https://github.com/Azure/aks-engine/issues/1693))
- Adding ephemeral disk support ([#1651](https://github.com/Azure/aks-engine/issues/1651))
- add support for Kubernetes 1.16.0-alpha.2 ([#1694](https://github.com/Azure/aks-engine/issues/1694))
- add support for Kubernetes 1.16.0-alpha.1 ([#1626](https://github.com/Azure/aks-engine/issues/1626))
- add support for Kubernetes 1.15.1 on Azure Stack. ([#1653](https://github.com/Azure/aks-engine/issues/1653))
- generate network configuration file for Azure CNI on Azure Stack ([#1584](https://github.com/Azure/aks-engine/issues/1584))
- add support for Kubernetes 1.15.1 ([#1642](https://github.com/Azure/aks-engine/issues/1642))
- pre-pull docker distribution for Azure Stack ([#1631](https://github.com/Azure/aks-engine/issues/1631))
- add support for Kubernetes 1.13.8 and 1.14.4 on Azure Stack. ([#1614](https://github.com/Azure/aks-engine/issues/1614))
- re-enable VMSS overprovisioning for agents ([#1601](https://github.com/Azure/aks-engine/issues/1601))
- Support kubectl port-forward with k8s 1.15 on Windows nodes ([#1543](https://github.com/Azure/aks-engine/issues/1543))
- add support for Kubernetes 1.14.4 ([#1596](https://github.com/Azure/aks-engine/issues/1596))
- add support for Kubernetes 1.13.8 ([#1595](https://github.com/Azure/aks-engine/issues/1595))
- change windows profile defaults for Azure Stack ([#1592](https://github.com/Azure/aks-engine/issues/1592))
- add scheduled maintenance addon, by default disabled ([#1575](https://github.com/Azure/aks-engine/issues/1575))

### Maintenance 🔧
- rev VHD images to 2019.09.19 ([#1998](https://github.com/Azure/aks-engine/issues/1998))
- update ip-masq-agent to v2.5.0 ([#1908](https://github.com/Azure/aks-engine/issues/1908))
- rev VHD images to 2019.09.16 ([#1962](https://github.com/Azure/aks-engine/issues/1962))
- install metrics-server v0.3.4 via VHD ([#1959](https://github.com/Azure/aks-engine/issues/1959))
- rev vhd to 2019.09.13 ([#1955](https://github.com/Azure/aks-engine/issues/1955))
- fix closure capture for UPGRADE_VERSIONS ([#1953](https://github.com/Azure/aks-engine/issues/1953))
- inject all params into env vars for the build ([#1942](https://github.com/Azure/aks-engine/issues/1942))
- add upgrade and scale parallel jobs ([#1941](https://github.com/Azure/aks-engine/issues/1941))
- rev VHD image references to 2019.09.10 ([#1940](https://github.com/Azure/aks-engine/issues/1940))
- rev etcd to v3.3.15 ([#1931](https://github.com/Azure/aks-engine/issues/1931))
- inject log analytics workspace key into test env config ([#1934](https://github.com/Azure/aks-engine/issues/1934))
- Update Windows to version  17763.678.1908092216 ([#1933](https://github.com/Azure/aks-engine/issues/1933))
- bump keyvault-flexvol to v0.0.13 ([#1924](https://github.com/Azure/aks-engine/issues/1924))
- add job exclusion regex to jenkinsfile ([#1922](https://github.com/Azure/aks-engine/issues/1922))
- make the pki key size settable. ([#1891](https://github.com/Azure/aks-engine/issues/1891))
- add windows, monitoring and base test configs ([#1899](https://github.com/Azure/aks-engine/issues/1899))
- remove oms extension docs since they are obsolete ([#1900](https://github.com/Azure/aks-engine/issues/1900))
- disable tiller addon by default ([#1884](https://github.com/Azure/aks-engine/issues/1884))
- update CoreDNS to 1.6.2 ([#1880](https://github.com/Azure/aks-engine/issues/1880))
- create parallel Jenkins pipeline for e2e tests ([#1875](https://github.com/Azure/aks-engine/issues/1875))
- add apache2-utils as an apt dependency ([#1822](https://github.com/Azure/aks-engine/issues/1822))
- remove make target w/ non-existent file reference ([#1861](https://github.com/Azure/aks-engine/issues/1861))
- Targeting August updates for Windows VHD ([#1846](https://github.com/Azure/aks-engine/issues/1846))
- rev VHD image references to 2019.08.21 ([#1849](https://github.com/Azure/aks-engine/issues/1849))
- tolerate lowercase LB SKU vals ([#1838](https://github.com/Azure/aks-engine/issues/1838))
- bump keyvault-flexvol to v0.0.12 ([#1831](https://github.com/Azure/aks-engine/issues/1831))
- update az go sdk to v32.5 and autorest v13 ([#1793](https://github.com/Azure/aks-engine/issues/1793))
- update go toolchain to 1.12.9 ([#1789](https://github.com/Azure/aks-engine/issues/1789)) ([#1826](https://github.com/Azure/aks-engine/issues/1826))
- rev VHD references to 2019.08.15 ([#1820](https://github.com/Azure/aks-engine/issues/1820))
- add k8s 1.13.10, 1.14.6, 1.15.3 to VHD script ([#1807](https://github.com/Azure/aks-engine/issues/1807))
- Change Windows default to Windows Server 2019 LTSC, July patch ([#1722](https://github.com/Azure/aks-engine/issues/1722))
- update go toolchain to 1.12.8 ([#1789](https://github.com/Azure/aks-engine/issues/1789))
- use numeric address to avoid DNS reverse lookup ([#1769](https://github.com/Azure/aks-engine/issues/1769))
- rev etcd to v3.3.13 ([#1772](https://github.com/Azure/aks-engine/issues/1772))
- update Azure CNI and NetworkPolicy to v1.0.25 ([#1774](https://github.com/Azure/aks-engine/issues/1774))
- update go-dev image to v1.23.0 ([#1756](https://github.com/Azure/aks-engine/issues/1756))
- rev VHD versions to 2019.08.09 ([#1757](https://github.com/Azure/aks-engine/issues/1757))
- Update to Docker EE 18.09.7 by default ([#1750](https://github.com/Azure/aks-engine/issues/1750))
- update kube-addon-manager to v9.0.2 and v8.9.1 ([#1746](https://github.com/Azure/aks-engine/issues/1746))
- Update calico to v3.8.0 ([#1636](https://github.com/Azure/aks-engine/issues/1636))
- configure docker ExecStartPost config in file ([#1727](https://github.com/Azure/aks-engine/issues/1727))
- fix error message not displayed when binary not found ([#1679](https://github.com/Azure/aks-engine/issues/1679))
- use ACR URI to validate outbound connectivity ([#1698](https://github.com/Azure/aks-engine/issues/1698))
- update VHDs to 2019.07.30 ([#1700](https://github.com/Azure/aks-engine/issues/1700))
- use k8s test infra archive for nssm.exe ([#1697](https://github.com/Azure/aks-engine/issues/1697))
- deliver artifacts via VHD where appropriate ([#1684](https://github.com/Azure/aks-engine/issues/1684))
- rev VHD to 2019.07.29 ([#1691](https://github.com/Azure/aks-engine/issues/1691))
- use HAS_GOLANGCI value if defined ([#1677](https://github.com/Azure/aks-engine/issues/1677))
- update VHD to 2019.07.25 ([#1673](https://github.com/Azure/aks-engine/issues/1673))
- update Azure CNI to v1.0.24 ([#1668](https://github.com/Azure/aks-engine/issues/1668))
- update cluster-autoscaler patch versions ([#1662](https://github.com/Azure/aks-engine/issues/1662))
- update go toolchain to 1.12.7 ([#1658](https://github.com/Azure/aks-engine/issues/1658))
- deprecate clear-containers runtime support ([#1649](https://github.com/Azure/aks-engine/issues/1649))
- update dnsmasq-nanny to v1.15.4 ([#1639](https://github.com/Azure/aks-engine/issues/1639))
- update kube-dns to v1.15.4 ([#1629](https://github.com/Azure/aks-engine/issues/1629))
- Update OWNERS ([#1619](https://github.com/Azure/aks-engine/issues/1619))
- re-enable k8s 1.12.7 ([#1603](https://github.com/Azure/aks-engine/issues/1603))
- update azure-sdk-for-go to v27.3.0 ([#1516](https://github.com/Azure/aks-engine/issues/1516))
- update Azure constants with new VM sizes ([#1591](https://github.com/Azure/aks-engine/issues/1591))

### Revert Change ◀️
- "fix: omit NSG rules for Upgrade scenario ([#1705](https://github.com/Azure/aks-engine/issues/1705))" ([#1754](https://github.com/Azure/aks-engine/issues/1754))

### Testing 💚
- run gpu tests in westus2 ([#1965](https://github.com/Azure/aks-engine/issues/1965))
- fix "run multiple commands in succession" implementation ([#1956](https://github.com/Azure/aks-engine/issues/1956))
- pod WaitOnSucceeded against distinct pod ([#1950](https://github.com/Azure/aks-engine/issues/1950))
- s/southcentralus/uksouth ([#1949](https://github.com/Azure/aks-engine/issues/1949))
- more resilience when k get deployment ([#1948](https://github.com/Azure/aks-engine/issues/1948))
- az storage file download-batch --destination must exist ([#1947](https://github.com/Azure/aks-engine/issues/1947))
- crash docker and validate against common node count ([#1946](https://github.com/Azure/aks-engine/issues/1946))
- E2E: goroutine and err response cleanup ([#1925](https://github.com/Azure/aks-engine/issues/1925))
- printing wrong err ([#1881](https://github.com/Azure/aks-engine/issues/1881))
- retry pod get ([#1879](https://github.com/Azure/aks-engine/issues/1879))
- remove obsolete e2e test scripts ([#1768](https://github.com/Azure/aks-engine/issues/1768))
- improve E2E node readiness tests ([#1859](https://github.com/Azure/aks-engine/issues/1859))
- check DNS before scp’ing ([#1857](https://github.com/Azure/aks-engine/issues/1857))
- e2e accommodations for low-pri vmss configurations ([#1854](https://github.com/Azure/aks-engine/issues/1854))
- improve dashboard e2e tests ([#1853](https://github.com/Azure/aks-engine/issues/1853))
- allow 20 retries for ssh-dependent tests ([#1852](https://github.com/Azure/aks-engine/issues/1852))
- add retries to host OS DNS E2E tests ([#1843](https://github.com/Azure/aks-engine/issues/1843))
- wait for pod readiness in port forward test ([#1845](https://github.com/Azure/aks-engine/issues/1845))
- only check for static n node count if not using low pri VMSS ([#1844](https://github.com/Azure/aks-engine/issues/1844))
- retry all ssh-dependent E2E tests ([#1825](https://github.com/Azure/aks-engine/issues/1825))
- Added unit test for NormalizeMasterResourcesForScaling function ([#1788](https://github.com/Azure/aks-engine/issues/1788))
- improve kubectl port-forward tests ([#1775](https://github.com/Azure/aks-engine/issues/1775))
- make crashing pods test opt-in (debug only) ([#1709](https://github.com/Azure/aks-engine/issues/1709))
- use /22 for agent pool subnet ([#1701](https://github.com/Azure/aks-engine/issues/1701))
- use kill instead of interrupt signal to terminate kubectl port-forward ([#1671](https://github.com/Azure/aks-engine/issues/1671))
- don't check for kubernetes.azure.com/role label ([#1667](https://github.com/Azure/aks-engine/issues/1667))
- replace “vanilla” E2E cluster test with non-VHD cluster test ([#1650](https://github.com/Azure/aks-engine/issues/1650))
- temporarily allow 10 CA restarts ([#1598](https://github.com/Azure/aks-engine/issues/1598))
- remove dependency on httpmock for Azure Stack unit tests ([#1589](https://github.com/Azure/aks-engine/issues/1589))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.38.9"></a>
# [v0.38.9] - 2019-09-20
### Bug Fixes 🐞
- cleanup apt artifacts before running apt ([#2006](https://github.com/Azure/aks-engine/issues/2006))
- re-add Promo SKUs to Azure constants ([#1752](https://github.com/Azure/aks-engine/issues/1752))
- "dpkg configure -a" interactive config update issue
- extend default value -1 for PodMaxPids for master and agentpool profile … ([#1623](https://github.com/Azure/aks-engine/issues/1623))
- Remove unnecessary transform on VMSS upgrade ([#1633](https://github.com/Azure/aks-engine/issues/1633))
- old VMSS AKS cluster cannot be upgraded after April 2 ([#1561](https://github.com/Azure/aks-engine/issues/1561))
- remove availability sets in scaling up template ([#1610](https://github.com/Azure/aks-engine/issues/1610))
- do not install clear containers runtime on the VHD ([#1587](https://github.com/Azure/aks-engine/issues/1587))
- improve clear containers clean up ([#1600](https://github.com/Azure/aks-engine/issues/1600))

### Features 🌈
- add support for Kubernetes 1.15.3 ([#1805](https://github.com/Azure/aks-engine/issues/1805))
- add support for Kubernetes 1.14.6 ([#1804](https://github.com/Azure/aks-engine/issues/1804))
- add support for Kubernetes 1.13.10 ([#1803](https://github.com/Azure/aks-engine/issues/1803))
- add support for Kubernetes 1.13.9, 1.14.5, 1.15.2 on Azure Stack ([#1740](https://github.com/Azure/aks-engine/issues/1740))
- add support for Kubernetes 1.14.5 ([#1731](https://github.com/Azure/aks-engine/issues/1731))
- add support for Kubernetes 1.13.9 ([#1732](https://github.com/Azure/aks-engine/issues/1732))
- add support for Kubernetes 1.15.2 ([#1730](https://github.com/Azure/aks-engine/issues/1730))
- add support for Kubernetes 1.15.1 ([#1642](https://github.com/Azure/aks-engine/issues/1642))
- add support for Kubernetes 1.14.4 ([#1596](https://github.com/Azure/aks-engine/issues/1596))
- add support for Kubernetes 1.13.8 ([#1595](https://github.com/Azure/aks-engine/issues/1595))

### Maintenance 🔧
- update go toolchain to 1.12.8 ([#1789](https://github.com/Azure/aks-engine/issues/1789))
- update Azure CNI and NetworkPolicy to v1.0.25 ([#1774](https://github.com/Azure/aks-engine/issues/1774))
- update Azure CNI to v1.0.24 ([#1668](https://github.com/Azure/aks-engine/issues/1668))
- rev VHD images to 2019.08.09387
- update kube-addon-manager to v9.0.2 and v8.9.1 ([#1746](https://github.com/Azure/aks-engine/issues/1746))
- update VHD to 2019.07.25 ([#1673](https://github.com/Azure/aks-engine/issues/1673))
- update VHD to 2019.07.10
- re-enable k8s 1.12.7 ([#1603](https://github.com/Azure/aks-engine/issues/1603))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.41.1"></a>
# [v0.41.1] - 2019-09-20
### Bug Fixes 🐞
- check the nodes of all agentpool during upgrade ([#1893](https://github.com/Azure/aks-engine/issues/1893))

### Maintenance 🔧
- rev VHD images to 2019.09.19 ([#1998](https://github.com/Azure/aks-engine/issues/1998))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.41.0"></a>
# [v0.41.0] - 2019-09-18
### Bug Fixes 🐞
- correct scale message after scale operations ([#1977](https://github.com/Azure/aks-engine/issues/1977))
- delete hyperkube images before starting kubelet ([#1816](https://github.com/Azure/aks-engine/issues/1816))
- not overwrite service principal profile for hosted master ([#1935](https://github.com/Azure/aks-engine/issues/1935))
- update the location of the VHD footprint file on VMs ([#1944](https://github.com/Azure/aks-engine/issues/1944))
- add hostname to CSE command so it shows in ARM deployment errors ([#1921](https://github.com/Azure/aks-engine/issues/1921))
- Remove the blind mkdir when writing outputs ([#1048](https://github.com/Azure/aks-engine/issues/1048)) ([#1889](https://github.com/Azure/aks-engine/issues/1889))
- add #EOF for dhcpv6 service ([#1909](https://github.com/Azure/aks-engine/issues/1909))
- Update rbac settings for cluster autoscaler ([#1582](https://github.com/Azure/aks-engine/issues/1582))
- add deprovision step to packer script ([#1865](https://github.com/Azure/aks-engine/issues/1865))
- set streaming-connection-idle-timeout as 4h ([#1870](https://github.com/Azure/aks-engine/issues/1870))
- improve container monitoring add-on ([#1686](https://github.com/Azure/aks-engine/issues/1686))
- more aks guarding against agentLbID template var ([#1874](https://github.com/Azure/aks-engine/issues/1874))
- guard AKS scenarios against aks-engine SLB template definitions ([#1872](https://github.com/Azure/aks-engine/issues/1872))
- fix customVMTags functionality and add more unit tests ([#1867](https://github.com/Azure/aks-engine/issues/1867))
- ensure master role assignment happens after VM deployment ([#1797](https://github.com/Azure/aks-engine/issues/1797))
- add Standard_F8 to accelerated networking whitelist ([#1858](https://github.com/Azure/aks-engine/issues/1858))
- retain custom nsg rules during upgrade ([#1792](https://github.com/Azure/aks-engine/issues/1792))
- Cannot set boolean property from command line ([#1848](https://github.com/Azure/aks-engine/issues/1848))
- restore working calico networkPolicy vlabs conversion logic ([#1855](https://github.com/Azure/aks-engine/issues/1855))

### Code Refactoring 💎
- configure kube-proxy 1.16 with a ConfigMap ([#1862](https://github.com/Azure/aks-engine/issues/1862))

### Continuous Integration 💜
- Running e2e tests as part of windows VHD build ([#1877](https://github.com/Azure/aks-engine/issues/1877))
- enable rg substring match in azure cleanup script ([#1832](https://github.com/Azure/aks-engine/issues/1832))
- add VHD rg cleanup script ([#1837](https://github.com/Azure/aks-engine/issues/1837))

### Documentation 📘
- add 1.16 api model example ([#1963](https://github.com/Azure/aks-engine/issues/1963))
- supported versions for disconnected Azure Stacks using VHD 2019.08.09 & 2019.08.21 ([#1907](https://github.com/Azure/aks-engine/issues/1907))
- fix invalid link ([#1876](https://github.com/Azure/aks-engine/issues/1876))

### Features 🌈
- add support for Kubernetes 1.16.0 ([#1972](https://github.com/Azure/aks-engine/issues/1972))
- add support for Kubernetes 1.13.11 ([#1975](https://github.com/Azure/aks-engine/issues/1975))
- add support for Kubernetes 1.15.4 ([#1974](https://github.com/Azure/aks-engine/issues/1974))
- add support for Kubernetes 1.14.7 ([#1973](https://github.com/Azure/aks-engine/issues/1973))
- upgrade metrics server to v0.3.4 ([#1109](https://github.com/Azure/aks-engine/issues/1109))
- add new Norway regions ([#1939](https://github.com/Azure/aks-engine/issues/1939))
- enable deployment telemetry for generate command on Azure Stack ([#1847](https://github.com/Azure/aks-engine/issues/1847))
- add support for Kubernetes 1.16.0-rc.1 ([#1938](https://github.com/Azure/aks-engine/issues/1938))
- add support for Kubernetes 1.16.0-beta.2 ([#1906](https://github.com/Azure/aks-engine/issues/1906))
- add build and deploy windows zip file support on Azure Stack ([#1888](https://github.com/Azure/aks-engine/issues/1888))
- add new Azure Switzerland regions ([#1904](https://github.com/Azure/aks-engine/issues/1904))
- add germanynorth and germanywestcentral regions ([#1897](https://github.com/Azure/aks-engine/issues/1897))
- add new Azure VM SKUs ([#1896](https://github.com/Azure/aks-engine/issues/1896))
- update default Kubernetes version to 1.13 ([#1850](https://github.com/Azure/aks-engine/issues/1850))
- Windows image references ([#1718](https://github.com/Azure/aks-engine/issues/1718))
- Allow customization of coredns ([#1541](https://github.com/Azure/aks-engine/issues/1541)) ([#1841](https://github.com/Azure/aks-engine/issues/1841))

### Maintenance 🔧
- update ip-masq-agent to v2.5.0 ([#1908](https://github.com/Azure/aks-engine/issues/1908))
- rev VHD images to 2019.09.16 ([#1962](https://github.com/Azure/aks-engine/issues/1962))
- install metrics-server v0.3.4 via VHD ([#1959](https://github.com/Azure/aks-engine/issues/1959))
- rev vhd to 2019.09.13 ([#1955](https://github.com/Azure/aks-engine/issues/1955))
- fix closure capture for UPGRADE_VERSIONS ([#1953](https://github.com/Azure/aks-engine/issues/1953))
- inject all params into env vars for the build ([#1942](https://github.com/Azure/aks-engine/issues/1942))
- add upgrade and scale parallel jobs ([#1941](https://github.com/Azure/aks-engine/issues/1941))
- rev VHD image references to 2019.09.10 ([#1940](https://github.com/Azure/aks-engine/issues/1940))
- rev etcd to v3.3.15 ([#1931](https://github.com/Azure/aks-engine/issues/1931))
- inject log analytics workspace key into test env config ([#1934](https://github.com/Azure/aks-engine/issues/1934))
- Update Windows to version  17763.678.1908092216 ([#1933](https://github.com/Azure/aks-engine/issues/1933))
- bump keyvault-flexvol to v0.0.13 ([#1924](https://github.com/Azure/aks-engine/issues/1924))
- add job exclusion regex to jenkinsfile ([#1922](https://github.com/Azure/aks-engine/issues/1922))
- make the pki key size settable. ([#1891](https://github.com/Azure/aks-engine/issues/1891))
- add windows, monitoring and base test configs ([#1899](https://github.com/Azure/aks-engine/issues/1899))
- remove oms extension docs since they are obsolete ([#1900](https://github.com/Azure/aks-engine/issues/1900))
- disable tiller addon by default ([#1884](https://github.com/Azure/aks-engine/issues/1884))
- update CoreDNS to 1.6.2 ([#1880](https://github.com/Azure/aks-engine/issues/1880))
- create parallel Jenkins pipeline for e2e tests ([#1875](https://github.com/Azure/aks-engine/issues/1875))
- add apache2-utils as an apt dependency ([#1822](https://github.com/Azure/aks-engine/issues/1822))
- remove make target w/ non-existent file reference ([#1861](https://github.com/Azure/aks-engine/issues/1861))
- Targeting August updates for Windows VHD ([#1846](https://github.com/Azure/aks-engine/issues/1846))
- rev VHD image references to 2019.08.21 ([#1849](https://github.com/Azure/aks-engine/issues/1849))
- tolerate lowercase LB SKU vals ([#1838](https://github.com/Azure/aks-engine/issues/1838))
- bump keyvault-flexvol to v0.0.12 ([#1831](https://github.com/Azure/aks-engine/issues/1831))
- update az go sdk to v32.5 and autorest v13 ([#1793](https://github.com/Azure/aks-engine/issues/1793))

### Testing 💚
- run gpu tests in westus2 ([#1965](https://github.com/Azure/aks-engine/issues/1965))
- fix "run multiple commands in succession" implementation ([#1956](https://github.com/Azure/aks-engine/issues/1956))
- pod WaitOnSucceeded against distinct pod ([#1950](https://github.com/Azure/aks-engine/issues/1950))
- s/southcentralus/uksouth ([#1949](https://github.com/Azure/aks-engine/issues/1949))
- more resilience when k get deployment ([#1948](https://github.com/Azure/aks-engine/issues/1948))
- az storage file download-batch --destination must exist ([#1947](https://github.com/Azure/aks-engine/issues/1947))
- crash docker and validate against common node count ([#1946](https://github.com/Azure/aks-engine/issues/1946))
- E2E: goroutine and err response cleanup ([#1925](https://github.com/Azure/aks-engine/issues/1925))
- printing wrong err ([#1881](https://github.com/Azure/aks-engine/issues/1881))
- retry pod get ([#1879](https://github.com/Azure/aks-engine/issues/1879))
- remove obsolete e2e test scripts ([#1768](https://github.com/Azure/aks-engine/issues/1768))
- improve E2E node readiness tests ([#1859](https://github.com/Azure/aks-engine/issues/1859))
- check DNS before scp’ing ([#1857](https://github.com/Azure/aks-engine/issues/1857))
- e2e accommodations for low-pri vmss configurations ([#1854](https://github.com/Azure/aks-engine/issues/1854))
- improve dashboard e2e tests ([#1853](https://github.com/Azure/aks-engine/issues/1853))
- allow 20 retries for ssh-dependent tests ([#1852](https://github.com/Azure/aks-engine/issues/1852))
- add retries to host OS DNS E2E tests ([#1843](https://github.com/Azure/aks-engine/issues/1843))
- wait for pod readiness in port forward test ([#1845](https://github.com/Azure/aks-engine/issues/1845))
- only check for static n node count if not using low pri VMSS ([#1844](https://github.com/Azure/aks-engine/issues/1844))
- retry all ssh-dependent E2E tests ([#1825](https://github.com/Azure/aks-engine/issues/1825))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.40.0"></a>
# [v0.40.0] - 2019-08-28
### Bug Fixes 🐞
- more aks guarding against agentLbID template var ([#1874](https://github.com/Azure/aks-engine/issues/1874))
- guard AKS scenarios against aks-engine SLB template definitions ([#1872](https://github.com/Azure/aks-engine/issues/1872))
- fix customVMTags functionality and add more unit tests ([#1867](https://github.com/Azure/aks-engine/issues/1867))
- ensure master role assignment happens after VM deployment ([#1797](https://github.com/Azure/aks-engine/issues/1797))
- add Standard_F8 to accelerated networking whitelist ([#1858](https://github.com/Azure/aks-engine/issues/1858))
- restore working calico networkPolicy vlabs conversion logic ([#1855](https://github.com/Azure/aks-engine/issues/1855))
- skip creating slb for hostedmaster ([#1835](https://github.com/Azure/aks-engine/issues/1835))
- networkplugin conversion only if networkpolicy is empty ([#1823](https://github.com/Azure/aks-engine/issues/1823))
- use azcopy-preview in VHD pipeline ([#1821](https://github.com/Azure/aks-engine/issues/1821))
- metrics server cluster role ([#1714](https://github.com/Azure/aks-engine/issues/1714))
- kube-proxy addon not critical, doesn't reconcile ([#1814](https://github.com/Azure/aks-engine/issues/1814))
- set kubelet config for containerd in templates ([#1785](https://github.com/Azure/aks-engine/issues/1785))
- race condition which causes concurrent map writes during tests ([#1791](https://github.com/Azure/aks-engine/issues/1791))
- label scheduled-maintenance manifests for addon-manager ([#1755](https://github.com/Azure/aks-engine/issues/1755))
- address all cloud-init file waits ([#1719](https://github.com/Azure/aks-engine/issues/1719))
- re-add Promo SKUs to Azure constants ([#1752](https://github.com/Azure/aks-engine/issues/1752))
- omit NSG rules for Upgrade scenario ([#1705](https://github.com/Azure/aks-engine/issues/1705))
- "dpkg configure -a" interactive config update issue
- Change win-cni to win-bridge binaries and config ([#643](https://github.com/Azure/aks-engine/issues/643))
- document registry container image pull in VHD release notes ([#1699](https://github.com/Azure/aks-engine/issues/1699))
- lock docker distribution version for Azure Stack ([#1688](https://github.com/Azure/aks-engine/issues/1688))
- omit NSG rules created by LoadBalancer during upgrade ([#1646](https://github.com/Azure/aks-engine/issues/1646))
- honour CLEANUP_IF_FAIL in test runner ([#1680](https://github.com/Azure/aks-engine/issues/1680))
- only include OrchestratorProfile validation during data marshalling for create ([#1661](https://github.com/Azure/aks-engine/issues/1661))
- add back-compat support for deprecated distro vals ([#1669](https://github.com/Azure/aks-engine/issues/1669))
- distinct component update messages for upgrade/scale ([#1666](https://github.com/Azure/aks-engine/issues/1666))
- add ipv6 nic config for vmss [IPv6DualStack] ([#1648](https://github.com/Azure/aks-engine/issues/1648))
- clusters with SLB scale and upgrade ([#1622](https://github.com/Azure/aks-engine/issues/1622))
- update container monitoring add-on to use  latest version of the containerized omsagent ([#1637](https://github.com/Azure/aks-engine/issues/1637))
- extend default value -1 for PodMaxPids for master and agentpool profile … ([#1623](https://github.com/Azure/aks-engine/issues/1623))
- update cluster-autoscaler role ([#1643](https://github.com/Azure/aks-engine/issues/1643))
- Remove unnecessary transform on VMSS upgrade ([#1633](https://github.com/Azure/aks-engine/issues/1633))
- CoreOS improvements and fixes ([#1578](https://github.com/Azure/aks-engine/issues/1578))
- Remove pod-critical annotations in 1.16 manifests. ([#1621](https://github.com/Azure/aks-engine/issues/1621))
- old VMSS AKS cluster cannot be upgraded after April 2 ([#1561](https://github.com/Azure/aks-engine/issues/1561))
- remove availability sets in scaling up template ([#1610](https://github.com/Azure/aks-engine/issues/1610))
- use explicit spec.Selector with apps/v1 api manifests ([#1599](https://github.com/Azure/aks-engine/issues/1599))
- improve clear containers clean up ([#1600](https://github.com/Azure/aks-engine/issues/1600))
- do not install clear containers runtime on the VHD ([#1587](https://github.com/Azure/aks-engine/issues/1587))
- kubeconfig incorrectly escape characters for azuredeploy.json ([#1574](https://github.com/Azure/aks-engine/issues/1574))

### Code Refactoring 💎
- configure kube-proxy 1.16 with a ConfigMap ([#1862](https://github.com/Azure/aks-engine/issues/1862))
- simplify scale operation CLI output ([#1345](https://github.com/Azure/aks-engine/issues/1345))

### Code Style 🎶
- read without -r will mangle backslashes ([#1790](https://github.com/Azure/aks-engine/issues/1790))

### Continuous Integration 💜
- fix VHD release notes Azure DevOps log parsing ([#1702](https://github.com/Azure/aks-engine/issues/1702))

### Documentation 📘
- gofish updates are automated now for AKS Engine releases ([#1818](https://github.com/Azure/aks-engine/issues/1818))
- fix the feature table ([#1796](https://github.com/Azure/aks-engine/issues/1796))
- Add devigned as owner ([#1784](https://github.com/Azure/aks-engine/issues/1784))
- increase master node count in Azure Stack template ([#1767](https://github.com/Azure/aks-engine/issues/1767))
- fix typo in examples/addons/cluster-autoscaler ([#1604](https://github.com/Azure/aks-engine/issues/1604))
- clean up IPv4/v6 dual stack README ([#1690](https://github.com/Azure/aks-engine/issues/1690))
- supported versions for disconnected Azure Stacks using VHD 2019.07.10 ([#1685](https://github.com/Azure/aks-engine/issues/1685))
- Update ipv6 cidr for dualstack ([#1683](https://github.com/Azure/aks-engine/issues/1683))
- use Microsoft's standard code of conduct for GitHub ([#1659](https://github.com/Azure/aks-engine/issues/1659))
- rename vhd release notes files to match sku names ([#1608](https://github.com/Azure/aks-engine/issues/1608))
- **dual-stack:** Add prerequisites ([#1692](https://github.com/Azure/aks-engine/issues/1692))

### Features 🌈
- add support for Kubernetes 1.16.0-beta.1 ([#1819](https://github.com/Azure/aks-engine/issues/1819))
- update cloud provider rate limit for vmas. ([#1808](https://github.com/Azure/aks-engine/issues/1808))
- autofill client Id and client secret for generate command ([#1766](https://github.com/Azure/aks-engine/issues/1766))
- add support for Kubernetes 1.15.3 ([#1805](https://github.com/Azure/aks-engine/issues/1805))
- add support for Kubernetes 1.14.6 ([#1804](https://github.com/Azure/aks-engine/issues/1804))
- add support for Kubernetes 1.13.10 ([#1803](https://github.com/Azure/aks-engine/issues/1803))
- Windows VHD build pipeline ([#1762](https://github.com/Azure/aks-engine/issues/1762))
- enable configurable cloudprovider ratelimits for write ([#1783](https://github.com/Azure/aks-engine/issues/1783))
- enable AKS VHD image in Fairfax ([#1770](https://github.com/Azure/aks-engine/issues/1770))
- add support for Kubernetes 1.16.0-alpha.3 ([#1745](https://github.com/Azure/aks-engine/issues/1745))
- add support for Kubernetes 1.13.9, 1.14.5, 1.15.2 on Azure Stack ([#1740](https://github.com/Azure/aks-engine/issues/1740))
- add support for Kubernetes 1.14.5 ([#1731](https://github.com/Azure/aks-engine/issues/1731))
- add support for Kubernetes 1.13.9 ([#1732](https://github.com/Azure/aks-engine/issues/1732))
- add support for Kubernetes 1.15.2 ([#1730](https://github.com/Azure/aks-engine/issues/1730))
- enable smart cloudprovider rate limiting ([#1693](https://github.com/Azure/aks-engine/issues/1693))
- Adding ephemeral disk support ([#1651](https://github.com/Azure/aks-engine/issues/1651))
- add support for Kubernetes 1.16.0-alpha.2 ([#1694](https://github.com/Azure/aks-engine/issues/1694))
- add support for Kubernetes 1.16.0-alpha.1 ([#1626](https://github.com/Azure/aks-engine/issues/1626))
- add support for Kubernetes 1.15.1 on Azure Stack. ([#1653](https://github.com/Azure/aks-engine/issues/1653))
- generate network configuration file for Azure CNI on Azure Stack ([#1584](https://github.com/Azure/aks-engine/issues/1584))
- add support for Kubernetes 1.15.1 ([#1642](https://github.com/Azure/aks-engine/issues/1642))
- pre-pull docker distribution for Azure Stack ([#1631](https://github.com/Azure/aks-engine/issues/1631))
- add support for Kubernetes 1.13.8 and 1.14.4 on Azure Stack. ([#1614](https://github.com/Azure/aks-engine/issues/1614))
- re-enable VMSS overprovisioning for agents ([#1601](https://github.com/Azure/aks-engine/issues/1601))
- Support kubectl port-forward with k8s 1.15 on Windows nodes ([#1543](https://github.com/Azure/aks-engine/issues/1543))
- add support for Kubernetes 1.14.4 ([#1596](https://github.com/Azure/aks-engine/issues/1596))
- add support for Kubernetes 1.13.8 ([#1595](https://github.com/Azure/aks-engine/issues/1595))
- change windows profile defaults for Azure Stack ([#1592](https://github.com/Azure/aks-engine/issues/1592))
- add scheduled maintenance addon, by default disabled ([#1575](https://github.com/Azure/aks-engine/issues/1575))

### Maintenance 🔧
- rev VHD image references to 2019.08.21 ([#1849](https://github.com/Azure/aks-engine/issues/1849))
- tolerate lowercase LB SKU vals ([#1838](https://github.com/Azure/aks-engine/issues/1838))
- update go toolchain to 1.12.9 ([#1789](https://github.com/Azure/aks-engine/issues/1789)) ([#1826](https://github.com/Azure/aks-engine/issues/1826))
- rev VHD references to 2019.08.15 ([#1820](https://github.com/Azure/aks-engine/issues/1820))
- add k8s 1.13.10, 1.14.6, 1.15.3 to VHD script ([#1807](https://github.com/Azure/aks-engine/issues/1807))
- Change Windows default to Windows Server 2019 LTSC, July patch ([#1722](https://github.com/Azure/aks-engine/issues/1722))
- update go toolchain to 1.12.8 ([#1789](https://github.com/Azure/aks-engine/issues/1789))
- use numeric address to avoid DNS reverse lookup ([#1769](https://github.com/Azure/aks-engine/issues/1769))
- rev etcd to v3.3.13 ([#1772](https://github.com/Azure/aks-engine/issues/1772))
- update Azure CNI and NetworkPolicy to v1.0.25 ([#1774](https://github.com/Azure/aks-engine/issues/1774))
- update go-dev image to v1.23.0 ([#1756](https://github.com/Azure/aks-engine/issues/1756))
- rev VHD versions to 2019.08.09 ([#1757](https://github.com/Azure/aks-engine/issues/1757))
- Update to Docker EE 18.09.7 by default ([#1750](https://github.com/Azure/aks-engine/issues/1750))
- update kube-addon-manager to v9.0.2 and v8.9.1 ([#1746](https://github.com/Azure/aks-engine/issues/1746))
- Update calico to v3.8.0 ([#1636](https://github.com/Azure/aks-engine/issues/1636))
- configure docker ExecStartPost config in file ([#1727](https://github.com/Azure/aks-engine/issues/1727))
- fix error message not displayed when binary not found ([#1679](https://github.com/Azure/aks-engine/issues/1679))
- use ACR URI to validate outbound connectivity ([#1698](https://github.com/Azure/aks-engine/issues/1698))
- update VHDs to 2019.07.30 ([#1700](https://github.com/Azure/aks-engine/issues/1700))
- use k8s test infra archive for nssm.exe ([#1697](https://github.com/Azure/aks-engine/issues/1697))
- deliver artifacts via VHD where appropriate ([#1684](https://github.com/Azure/aks-engine/issues/1684))
- rev VHD to 2019.07.29 ([#1691](https://github.com/Azure/aks-engine/issues/1691))
- use HAS_GOLANGCI value if defined ([#1677](https://github.com/Azure/aks-engine/issues/1677))
- update VHD to 2019.07.25 ([#1673](https://github.com/Azure/aks-engine/issues/1673))
- update Azure CNI to v1.0.24 ([#1668](https://github.com/Azure/aks-engine/issues/1668))
- update cluster-autoscaler patch versions ([#1662](https://github.com/Azure/aks-engine/issues/1662))
- update go toolchain to 1.12.7 ([#1658](https://github.com/Azure/aks-engine/issues/1658))
- deprecate clear-containers runtime support ([#1649](https://github.com/Azure/aks-engine/issues/1649))
- update dnsmasq-nanny to v1.15.4 ([#1639](https://github.com/Azure/aks-engine/issues/1639))
- update kube-dns to v1.15.4 ([#1629](https://github.com/Azure/aks-engine/issues/1629))
- Update OWNERS ([#1619](https://github.com/Azure/aks-engine/issues/1619))
- re-enable k8s 1.12.7 ([#1603](https://github.com/Azure/aks-engine/issues/1603))
- update azure-sdk-for-go to v27.3.0 ([#1516](https://github.com/Azure/aks-engine/issues/1516))
- update Azure constants with new VM sizes ([#1591](https://github.com/Azure/aks-engine/issues/1591))

### Revert Change ◀️
- "fix: omit NSG rules for Upgrade scenario ([#1705](https://github.com/Azure/aks-engine/issues/1705))" ([#1754](https://github.com/Azure/aks-engine/issues/1754))

### Testing 💚
- Added unit test for NormalizeMasterResourcesForScaling function ([#1788](https://github.com/Azure/aks-engine/issues/1788))
- improve kubectl port-forward tests ([#1775](https://github.com/Azure/aks-engine/issues/1775))
- make crashing pods test opt-in (debug only) ([#1709](https://github.com/Azure/aks-engine/issues/1709))
- use /22 for agent pool subnet ([#1701](https://github.com/Azure/aks-engine/issues/1701))
- use kill instead of interrupt signal to terminate kubectl port-forward ([#1671](https://github.com/Azure/aks-engine/issues/1671))
- don't check for kubernetes.azure.com/role label ([#1667](https://github.com/Azure/aks-engine/issues/1667))
- replace “vanilla” E2E cluster test with non-VHD cluster test ([#1650](https://github.com/Azure/aks-engine/issues/1650))
- temporarily allow 10 CA restarts ([#1598](https://github.com/Azure/aks-engine/issues/1598))
- remove dependency on httpmock for Azure Stack unit tests ([#1589](https://github.com/Azure/aks-engine/issues/1589))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.38.8"></a>
# [v0.38.8] - 2019-08-19
### Bug Fixes 🐞
- re-add Promo SKUs to Azure constants ([#1752](https://github.com/Azure/aks-engine/issues/1752))
- "dpkg configure -a" interactive config update issue
- extend default value -1 for PodMaxPids for master and agentpool profile … ([#1623](https://github.com/Azure/aks-engine/issues/1623))
- Remove unnecessary transform on VMSS upgrade ([#1633](https://github.com/Azure/aks-engine/issues/1633))
- old VMSS AKS cluster cannot be upgraded after April 2 ([#1561](https://github.com/Azure/aks-engine/issues/1561))
- remove availability sets in scaling up template ([#1610](https://github.com/Azure/aks-engine/issues/1610))
- do not install clear containers runtime on the VHD ([#1587](https://github.com/Azure/aks-engine/issues/1587))
- improve clear containers clean up ([#1600](https://github.com/Azure/aks-engine/issues/1600))

### Features 🌈
- add support for Kubernetes 1.15.3 ([#1805](https://github.com/Azure/aks-engine/issues/1805))
- add support for Kubernetes 1.14.6 ([#1804](https://github.com/Azure/aks-engine/issues/1804))
- add support for Kubernetes 1.13.10 ([#1803](https://github.com/Azure/aks-engine/issues/1803))
- add support for Kubernetes 1.13.9, 1.14.5, 1.15.2 on Azure Stack ([#1740](https://github.com/Azure/aks-engine/issues/1740))
- add support for Kubernetes 1.14.5 ([#1731](https://github.com/Azure/aks-engine/issues/1731))
- add support for Kubernetes 1.13.9 ([#1732](https://github.com/Azure/aks-engine/issues/1732))
- add support for Kubernetes 1.15.2 ([#1730](https://github.com/Azure/aks-engine/issues/1730))
- add support for Kubernetes 1.15.1 ([#1642](https://github.com/Azure/aks-engine/issues/1642))
- add support for Kubernetes 1.14.4 ([#1596](https://github.com/Azure/aks-engine/issues/1596))
- add support for Kubernetes 1.13.8 ([#1595](https://github.com/Azure/aks-engine/issues/1595))

### Maintenance 🔧
- update go toolchain to 1.12.8 ([#1789](https://github.com/Azure/aks-engine/issues/1789))
- update Azure CNI and NetworkPolicy to v1.0.25 ([#1774](https://github.com/Azure/aks-engine/issues/1774))
- update Azure CNI to v1.0.24 ([#1668](https://github.com/Azure/aks-engine/issues/1668))
- rev VHD images to 2019.08.09387
- update kube-addon-manager to v9.0.2 and v8.9.1 ([#1746](https://github.com/Azure/aks-engine/issues/1746))
- update VHD to 2019.07.25 ([#1673](https://github.com/Azure/aks-engine/issues/1673))
- update VHD to 2019.07.10
- re-enable k8s 1.12.7 ([#1603](https://github.com/Azure/aks-engine/issues/1603))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.39.2"></a>
# [v0.39.2] - 2019-08-19
### Bug Fixes 🐞
- re-add Promo SKUs to Azure constants ([#1752](https://github.com/Azure/aks-engine/issues/1752))

### Features 🌈
- add support for Kubernetes 1.15.3 ([#1805](https://github.com/Azure/aks-engine/issues/1805))
- add support for Kubernetes 1.14.6 ([#1804](https://github.com/Azure/aks-engine/issues/1804))
- add support for Kubernetes 1.13.10 ([#1803](https://github.com/Azure/aks-engine/issues/1803))

### Maintenance 🔧
- update go toolchain to 1.12.8 ([#1789](https://github.com/Azure/aks-engine/issues/1789))
- update Azure CNI and NetworkPolicy to v1.0.25 ([#1774](https://github.com/Azure/aks-engine/issues/1774))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.39.1"></a>
# [v0.39.1] - 2019-08-10
### Bug Fixes 🐞
- "dpkg configure -a" interactive config update issue
- lock docker distribution version for Azure Stack ([#1688](https://github.com/Azure/aks-engine/issues/1688))
- omit NSG rules created by LoadBalancer during upgrade ([#1646](https://github.com/Azure/aks-engine/issues/1646))
- honour CLEANUP_IF_FAIL in test runner ([#1680](https://github.com/Azure/aks-engine/issues/1680))
- only include OrchestratorProfile validation during data marshalling for create ([#1661](https://github.com/Azure/aks-engine/issues/1661))
- add back-compat support for deprecated distro vals ([#1669](https://github.com/Azure/aks-engine/issues/1669))
- distinct component update messages for upgrade/scale ([#1666](https://github.com/Azure/aks-engine/issues/1666))
- add ipv6 nic config for vmss [IPv6DualStack] ([#1648](https://github.com/Azure/aks-engine/issues/1648))
- clusters with SLB scale and upgrade ([#1622](https://github.com/Azure/aks-engine/issues/1622))
- update container monitoring add-on to use  latest version of the containerized omsagent ([#1637](https://github.com/Azure/aks-engine/issues/1637))
- extend default value -1 for PodMaxPids for master and agentpool profile … ([#1623](https://github.com/Azure/aks-engine/issues/1623))
- update cluster-autoscaler role ([#1643](https://github.com/Azure/aks-engine/issues/1643))
- Remove unnecessary transform on VMSS upgrade ([#1633](https://github.com/Azure/aks-engine/issues/1633))
- CoreOS improvements and fixes ([#1578](https://github.com/Azure/aks-engine/issues/1578))
- Remove pod-critical annotations in 1.16 manifests. ([#1621](https://github.com/Azure/aks-engine/issues/1621))
- old VMSS AKS cluster cannot be upgraded after April 2 ([#1561](https://github.com/Azure/aks-engine/issues/1561))
- remove availability sets in scaling up template ([#1610](https://github.com/Azure/aks-engine/issues/1610))
- use explicit spec.Selector with apps/v1 api manifests ([#1599](https://github.com/Azure/aks-engine/issues/1599))
- improve clear containers clean up ([#1600](https://github.com/Azure/aks-engine/issues/1600))
- do not install clear containers runtime on the VHD ([#1587](https://github.com/Azure/aks-engine/issues/1587))
- kubeconfig incorrectly escape characters for azuredeploy.json ([#1574](https://github.com/Azure/aks-engine/issues/1574))

### Documentation 📘
- clean up IPv4/v6 dual stack README ([#1690](https://github.com/Azure/aks-engine/issues/1690))
- supported versions for disconnected Azure Stacks using VHD 2019.07.10 ([#1685](https://github.com/Azure/aks-engine/issues/1685))
- Update ipv6 cidr for dualstack ([#1683](https://github.com/Azure/aks-engine/issues/1683))
- use Microsoft's standard code of conduct for GitHub ([#1659](https://github.com/Azure/aks-engine/issues/1659))
- rename vhd release notes files to match sku names ([#1608](https://github.com/Azure/aks-engine/issues/1608))
- fix typo in examples/addons/cluster-autoscaler ([#1604](https://github.com/Azure/aks-engine/issues/1604))

### Features 🌈
- add support for Kubernetes 1.13.9, 1.14.5, 1.15.2 on Azure Stack ([#1740](https://github.com/Azure/aks-engine/issues/1740))
- add support for Kubernetes 1.14.5 ([#1731](https://github.com/Azure/aks-engine/issues/1731))
- add support for Kubernetes 1.13.9 ([#1732](https://github.com/Azure/aks-engine/issues/1732))
- add support for Kubernetes 1.15.2 ([#1730](https://github.com/Azure/aks-engine/issues/1730))
- add support for Kubernetes 1.16.0-alpha.1 ([#1626](https://github.com/Azure/aks-engine/issues/1626))
- add support for Kubernetes 1.15.1 on Azure Stack. ([#1653](https://github.com/Azure/aks-engine/issues/1653))
- generate network configuration file for Azure CNI on Azure Stack ([#1584](https://github.com/Azure/aks-engine/issues/1584))
- add support for Kubernetes 1.15.1 ([#1642](https://github.com/Azure/aks-engine/issues/1642))
- pre-pull docker distribution for Azure Stack ([#1631](https://github.com/Azure/aks-engine/issues/1631))
- add support for Kubernetes 1.13.8 and 1.14.4 on Azure Stack. ([#1614](https://github.com/Azure/aks-engine/issues/1614))
- re-enable VMSS overprovisioning for agents ([#1601](https://github.com/Azure/aks-engine/issues/1601))
- Support kubectl port-forward with k8s 1.15 on Windows nodes ([#1543](https://github.com/Azure/aks-engine/issues/1543))
- add support for Kubernetes 1.14.4 ([#1596](https://github.com/Azure/aks-engine/issues/1596))
- add support for Kubernetes 1.13.8 ([#1595](https://github.com/Azure/aks-engine/issues/1595))
- change windows profile defaults for Azure Stack ([#1592](https://github.com/Azure/aks-engine/issues/1592))
- add scheduled maintenance addon, by default disabled ([#1575](https://github.com/Azure/aks-engine/issues/1575))

### Maintenance 🔧
- rev VHD versions to 2019.08.09 ([#1757](https://github.com/Azure/aks-engine/issues/1757))
- update kube-addon-manager to v9.0.2 and v8.9.1 ([#1746](https://github.com/Azure/aks-engine/issues/1746))
- update VHDs to 2019.07.30 ([#1700](https://github.com/Azure/aks-engine/issues/1700))
- deliver artifacts via VHD where appropriate ([#1684](https://github.com/Azure/aks-engine/issues/1684))
- rev VHD to 2019.07.29 ([#1691](https://github.com/Azure/aks-engine/issues/1691))
- use HAS_GOLANGCI value if defined ([#1677](https://github.com/Azure/aks-engine/issues/1677))
- update VHD to 2019.07.25 ([#1673](https://github.com/Azure/aks-engine/issues/1673))
- update Azure CNI to v1.0.24 ([#1668](https://github.com/Azure/aks-engine/issues/1668))
- update cluster-autoscaler patch versions ([#1662](https://github.com/Azure/aks-engine/issues/1662))
- update go toolchain to 1.12.7 ([#1658](https://github.com/Azure/aks-engine/issues/1658))
- deprecate clear-containers runtime support ([#1649](https://github.com/Azure/aks-engine/issues/1649))
- update dnsmasq-nanny to v1.15.4 ([#1639](https://github.com/Azure/aks-engine/issues/1639))
- update kube-dns to v1.15.4 ([#1629](https://github.com/Azure/aks-engine/issues/1629))
- Update OWNERS ([#1619](https://github.com/Azure/aks-engine/issues/1619))
- re-enable k8s 1.12.7 ([#1603](https://github.com/Azure/aks-engine/issues/1603))
- update azure-sdk-for-go to v27.3.0 ([#1516](https://github.com/Azure/aks-engine/issues/1516))
- update Azure constants with new VM sizes ([#1591](https://github.com/Azure/aks-engine/issues/1591))

### Testing 💚
- use kill instead of interrupt signal to terminate kubectl port-forward ([#1671](https://github.com/Azure/aks-engine/issues/1671))
- don't check for kubernetes.azure.com/role label ([#1667](https://github.com/Azure/aks-engine/issues/1667))
- replace “vanilla” E2E cluster test with non-VHD cluster test ([#1650](https://github.com/Azure/aks-engine/issues/1650))
- temporarily allow 10 CA restarts ([#1598](https://github.com/Azure/aks-engine/issues/1598))
- remove dependency on httpmock for Azure Stack unit tests ([#1589](https://github.com/Azure/aks-engine/issues/1589))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.38.7"></a>
# [v0.38.7] - 2019-08-10
### Bug Fixes 🐞
- "dpkg configure -a" interactive config update issue

### Maintenance 🔧
- rev VHD images to 2019.08.09387
- update kube-addon-manager to v9.0.2 and v8.9.1 ([#1746](https://github.com/Azure/aks-engine/issues/1746))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.38.6"></a>
# [v0.38.6] - 2019-08-07
### Bug Fixes 🐞
- extend default value -1 for PodMaxPids for master and agentpool profile … ([#1623](https://github.com/Azure/aks-engine/issues/1623))
- Remove unnecessary transform on VMSS upgrade ([#1633](https://github.com/Azure/aks-engine/issues/1633))
- old VMSS AKS cluster cannot be upgraded after April 2 ([#1561](https://github.com/Azure/aks-engine/issues/1561))
- remove availability sets in scaling up template ([#1610](https://github.com/Azure/aks-engine/issues/1610))
- do not install clear containers runtime on the VHD ([#1587](https://github.com/Azure/aks-engine/issues/1587))
- improve clear containers clean up ([#1600](https://github.com/Azure/aks-engine/issues/1600))

### Features 🌈
- add support for Kubernetes 1.13.9, 1.14.5, 1.15.2 on Azure Stack ([#1740](https://github.com/Azure/aks-engine/issues/1740))
- add support for Kubernetes 1.14.5 ([#1731](https://github.com/Azure/aks-engine/issues/1731))
- add support for Kubernetes 1.13.9 ([#1732](https://github.com/Azure/aks-engine/issues/1732))
- add support for Kubernetes 1.15.2 ([#1730](https://github.com/Azure/aks-engine/issues/1730))
- add support for Kubernetes 1.15.1 ([#1642](https://github.com/Azure/aks-engine/issues/1642))
- add support for Kubernetes 1.14.4 ([#1596](https://github.com/Azure/aks-engine/issues/1596))
- add support for Kubernetes 1.13.8 ([#1595](https://github.com/Azure/aks-engine/issues/1595))

### Maintenance 🔧
- update VHD to 2019.07.25 ([#1673](https://github.com/Azure/aks-engine/issues/1673))
- update VHD to 2019.07.10
- re-enable k8s 1.12.7 ([#1603](https://github.com/Azure/aks-engine/issues/1603))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.39.0"></a>
# [v0.39.0] - 2019-07-31
### Bug Fixes 🐞
- lock docker distribution version for Azure Stack ([#1688](https://github.com/Azure/aks-engine/issues/1688))
- omit NSG rules created by LoadBalancer during upgrade ([#1646](https://github.com/Azure/aks-engine/issues/1646))
- honour CLEANUP_IF_FAIL in test runner ([#1680](https://github.com/Azure/aks-engine/issues/1680))
- only include OrchestratorProfile validation during data marshalling for create ([#1661](https://github.com/Azure/aks-engine/issues/1661))
- add back-compat support for deprecated distro vals ([#1669](https://github.com/Azure/aks-engine/issues/1669))
- distinct component update messages for upgrade/scale ([#1666](https://github.com/Azure/aks-engine/issues/1666))
- add ipv6 nic config for vmss [IPv6DualStack] ([#1648](https://github.com/Azure/aks-engine/issues/1648))
- clusters with SLB scale and upgrade ([#1622](https://github.com/Azure/aks-engine/issues/1622))
- update container monitoring add-on to use  latest version of the containerized omsagent ([#1637](https://github.com/Azure/aks-engine/issues/1637))
- extend default value -1 for PodMaxPids for master and agentpool profile … ([#1623](https://github.com/Azure/aks-engine/issues/1623))
- update cluster-autoscaler role ([#1643](https://github.com/Azure/aks-engine/issues/1643))
- Remove unnecessary transform on VMSS upgrade ([#1633](https://github.com/Azure/aks-engine/issues/1633))
- CoreOS improvements and fixes ([#1578](https://github.com/Azure/aks-engine/issues/1578))
- Remove pod-critical annotations in 1.16 manifests. ([#1621](https://github.com/Azure/aks-engine/issues/1621))
- old VMSS AKS cluster cannot be upgraded after April 2 ([#1561](https://github.com/Azure/aks-engine/issues/1561))
- remove availability sets in scaling up template ([#1610](https://github.com/Azure/aks-engine/issues/1610))
- use explicit spec.Selector with apps/v1 api manifests ([#1599](https://github.com/Azure/aks-engine/issues/1599))
- improve clear containers clean up ([#1600](https://github.com/Azure/aks-engine/issues/1600))
- do not install clear containers runtime on the VHD ([#1587](https://github.com/Azure/aks-engine/issues/1587))
- kubeconfig incorrectly escape characters for azuredeploy.json ([#1574](https://github.com/Azure/aks-engine/issues/1574))

### Documentation 📘
- clean up IPv4/v6 dual stack README ([#1690](https://github.com/Azure/aks-engine/issues/1690))
- supported versions for disconnected Azure Stacks using VHD 2019.07.10 ([#1685](https://github.com/Azure/aks-engine/issues/1685))
- Update ipv6 cidr for dualstack ([#1683](https://github.com/Azure/aks-engine/issues/1683))
- use Microsoft's standard code of conduct for GitHub ([#1659](https://github.com/Azure/aks-engine/issues/1659))
- rename vhd release notes files to match sku names ([#1608](https://github.com/Azure/aks-engine/issues/1608))
- fix typo in examples/addons/cluster-autoscaler ([#1604](https://github.com/Azure/aks-engine/issues/1604))

### Features 🌈
- add support for Kubernetes 1.16.0-alpha.1 ([#1626](https://github.com/Azure/aks-engine/issues/1626))
- add support for Kubernetes 1.15.1 on Azure Stack. ([#1653](https://github.com/Azure/aks-engine/issues/1653))
- generate network configuration file for Azure CNI on Azure Stack ([#1584](https://github.com/Azure/aks-engine/issues/1584))
- add support for Kubernetes 1.15.1 ([#1642](https://github.com/Azure/aks-engine/issues/1642))
- pre-pull docker distribution for Azure Stack ([#1631](https://github.com/Azure/aks-engine/issues/1631))
- add support for Kubernetes 1.13.8 and 1.14.4 on Azure Stack. ([#1614](https://github.com/Azure/aks-engine/issues/1614))
- re-enable VMSS overprovisioning for agents ([#1601](https://github.com/Azure/aks-engine/issues/1601))
- Support kubectl port-forward with k8s 1.15 on Windows nodes ([#1543](https://github.com/Azure/aks-engine/issues/1543))
- add support for Kubernetes 1.14.4 ([#1596](https://github.com/Azure/aks-engine/issues/1596))
- add support for Kubernetes 1.13.8 ([#1595](https://github.com/Azure/aks-engine/issues/1595))
- change windows profile defaults for Azure Stack ([#1592](https://github.com/Azure/aks-engine/issues/1592))
- add scheduled maintenance addon, by default disabled ([#1575](https://github.com/Azure/aks-engine/issues/1575))

### Maintenance 🔧
- update VHDs to 2019.07.30 ([#1700](https://github.com/Azure/aks-engine/issues/1700))
- deliver artifacts via VHD where appropriate ([#1684](https://github.com/Azure/aks-engine/issues/1684))
- rev VHD to 2019.07.29 ([#1691](https://github.com/Azure/aks-engine/issues/1691))
- use HAS_GOLANGCI value if defined ([#1677](https://github.com/Azure/aks-engine/issues/1677))
- update VHD to 2019.07.25 ([#1673](https://github.com/Azure/aks-engine/issues/1673))
- update Azure CNI to v1.0.24 ([#1668](https://github.com/Azure/aks-engine/issues/1668))
- update cluster-autoscaler patch versions ([#1662](https://github.com/Azure/aks-engine/issues/1662))
- update go toolchain to 1.12.7 ([#1658](https://github.com/Azure/aks-engine/issues/1658))
- deprecate clear-containers runtime support ([#1649](https://github.com/Azure/aks-engine/issues/1649))
- update dnsmasq-nanny to v1.15.4 ([#1639](https://github.com/Azure/aks-engine/issues/1639))
- update kube-dns to v1.15.4 ([#1629](https://github.com/Azure/aks-engine/issues/1629))
- Update OWNERS ([#1619](https://github.com/Azure/aks-engine/issues/1619))
- re-enable k8s 1.12.7 ([#1603](https://github.com/Azure/aks-engine/issues/1603))
- update azure-sdk-for-go to v27.3.0 ([#1516](https://github.com/Azure/aks-engine/issues/1516))
- update Azure constants with new VM sizes ([#1591](https://github.com/Azure/aks-engine/issues/1591))

### Testing 💚
- use kill instead of interrupt signal to terminate kubectl port-forward ([#1671](https://github.com/Azure/aks-engine/issues/1671))
- don't check for kubernetes.azure.com/role label ([#1667](https://github.com/Azure/aks-engine/issues/1667))
- replace “vanilla” E2E cluster test with non-VHD cluster test ([#1650](https://github.com/Azure/aks-engine/issues/1650))
- temporarily allow 10 CA restarts ([#1598](https://github.com/Azure/aks-engine/issues/1598))
- remove dependency on httpmock for Azure Stack unit tests ([#1589](https://github.com/Azure/aks-engine/issues/1589))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.38.5"></a>
# [v0.38.5] - 2019-07-26
### Features 🌈
- add support for Kubernetes 1.15.1 ([#1642](https://github.com/Azure/aks-engine/issues/1642))
- add support for Kubernetes 1.14.4 ([#1596](https://github.com/Azure/aks-engine/issues/1596))
- add support for Kubernetes 1.13.8 ([#1595](https://github.com/Azure/aks-engine/issues/1595))

### Maintenance 🔧
- update VHD to 2019.07.25 ([#1673](https://github.com/Azure/aks-engine/issues/1673))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.38.4"></a>
# [v0.38.4] - 2019-07-22
### Bug Fixes 🐞
- extend default value -1 for PodMaxPids for master and agentpool profile … ([#1623](https://github.com/Azure/aks-engine/issues/1623))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.38.3"></a>
# [v0.38.3] - 2019-07-18
### Bug Fixes 🐞
- Remove unnecessary transform on VMSS upgrade ([#1633](https://github.com/Azure/aks-engine/issues/1633))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.38.2"></a>
# [v0.38.2] - 2019-07-11
### Bug Fixes 🐞
- old VMSS AKS cluster cannot be upgraded after April 2 ([#1561](https://github.com/Azure/aks-engine/issues/1561))
- remove availability sets in scaling up template ([#1610](https://github.com/Azure/aks-engine/issues/1610))
- do not install clear containers runtime on the VHD ([#1587](https://github.com/Azure/aks-engine/issues/1587))
- improve clear containers clean up ([#1600](https://github.com/Azure/aks-engine/issues/1600))
- do not override vnetCidr param when empty string ([#1570](https://github.com/Azure/aks-engine/issues/1570))
- use correct path for monitor script in monitor units ([#1555](https://github.com/Azure/aks-engine/issues/1555))
- UAE North and UAE Central regions are missing from certificate host names ([#1554](https://github.com/Azure/aks-engine/issues/1554))
- --anonymous-auth and --client-ca-file in insecure kubelet config ([#1533](https://github.com/Azure/aks-engine/issues/1533))
- wait_for_apt_locks before GPU driver install ([#1538](https://github.com/Azure/aks-engine/issues/1538))
- Ensure etcd starts after disk is mounted ([#1515](https://github.com/Azure/aks-engine/issues/1515)) ([#1535](https://github.com/Azure/aks-engine/issues/1535))
- download aks-engine in Azure Cloud Shell ([#1523](https://github.com/Azure/aks-engine/issues/1523))
- Set encoding:gzip on cilium BPFFS mount unit ([#1514](https://github.com/Azure/aks-engine/issues/1514))
- ensure calico is properly enabled during upgrade ([#1510](https://github.com/Azure/aks-engine/issues/1510))
- --pod-max-pids upgrade to 1.14 scenarios ([#1508](https://github.com/Azure/aks-engine/issues/1508))
- check for nil masterProfile when setting FD count ([#1509](https://github.com/Azure/aks-engine/issues/1509))
- vhd pre-pulls expected flannel images tag ([#1484](https://github.com/Azure/aks-engine/issues/1484))

### Continuous Integration 💜
- swap ensure-generated and test-style tasks ([#1501](https://github.com/Azure/aks-engine/issues/1501))

### Documentation 📘
- change accelerated networking example VMs ([#1569](https://github.com/Azure/aks-engine/issues/1569))
- make build before make clean in release docs ([#1545](https://github.com/Azure/aks-engine/issues/1545))
- add marketplace prerequisites for Azure Stack ([#1499](https://github.com/Azure/aks-engine/issues/1499))
- Link to more details on Windows images ([#1506](https://github.com/Azure/aks-engine/issues/1506))
- list pre-pulled versions for Azure Stack ([#1482](https://github.com/Azure/aks-engine/issues/1482))

### Features 🌈
- support CoreOS-like Linux distributions ([#1572](https://github.com/Azure/aks-engine/issues/1572))
- allow custom os images from external subs via Shared Image Galleries ([#913](https://github.com/Azure/aks-engine/issues/913))
- default AcceleratedNetworkingEnabledWindows to false if target platform is Azure Stack ([#1568](https://github.com/Azure/aks-engine/issues/1568))
- Set EtcdDiskSizeGB max default size to 1023 if target platform is Azure Stack ([#1558](https://github.com/Azure/aks-engine/issues/1558))
- Add checks for Metadata, AcceleratedNetworking and etcdDiskSizeGB on Azure Stack to fail fast. ([#1564](https://github.com/Azure/aks-engine/issues/1564))
- support for Application Gateway Ingress Controller ([#1442](https://github.com/Azure/aks-engine/issues/1442))
- support CoreOS-like Linux distributions ([#1560](https://github.com/Azure/aks-engine/issues/1560))
- set ipam environment to mas in Azure CNI config on Azure Stack ([#1542](https://github.com/Azure/aks-engine/issues/1542))
- Add defaults for useInstanceMetadata and AcceleratedNetworkingEnabled on Azure Stack ([#1549](https://github.com/Azure/aks-engine/issues/1549))
- Set PlatformFaultDomainCount value as 3 for Azure Stack ([#1529](https://github.com/Azure/aks-engine/issues/1529))
- Add vscode containerized development config ([#1526](https://github.com/Azure/aks-engine/issues/1526))
- enable deployment of windows nodes on Azure Stack ([#1456](https://github.com/Azure/aks-engine/issues/1456))
- add support for Kubernetes 1.15.0 ([#1502](https://github.com/Azure/aks-engine/issues/1502))
- Adding Windows Server version 1903 support ([#1464](https://github.com/Azure/aks-engine/issues/1464))
- IPv6dualStack feature flag ([#1424](https://github.com/Azure/aks-engine/issues/1424))

### Maintenance 🔧
- update VHD to 2019.07.10
- re-enable k8s 1.12.7 ([#1603](https://github.com/Azure/aks-engine/issues/1603))
- update VHD to 2019.07.01 ([#1565](https://github.com/Azure/aks-engine/issues/1565))
- replace deprecated extensions/v1beta APIs for 1.16 ([#1527](https://github.com/Azure/aks-engine/issues/1527))
- include https://fqdn:443 for AKS upgrade ([#1544](https://github.com/Azure/aks-engine/issues/1544))
- enable VHD 2019.06.26 for 16.04-LTS ([#1540](https://github.com/Azure/aks-engine/issues/1540))
- Update to MS Moby 3.0.6 ([#1537](https://github.com/Azure/aks-engine/issues/1537))
- Add support for Kubernetes 1.13.7 and 1.14.3 on Azure Stack ([#1525](https://github.com/Azure/aks-engine/issues/1525))
- enable VHD 2019.06.20 ([#1513](https://github.com/Azure/aks-engine/issues/1513))
- update addon-resizer to 1.8.5 for k8s 1.15 ([#1511](https://github.com/Azure/aks-engine/issues/1511))
- Change Windows default image to Jun 2019 ([#1489](https://github.com/Azure/aks-engine/issues/1489))
- update templates_generated.go ([#1500](https://github.com/Azure/aks-engine/issues/1500))
- CoreDNS configmap is not updated during upgrades ([#1493](https://github.com/Azure/aks-engine/issues/1493))
- update go toolchain to 1.12.6 ([#1487](https://github.com/Azure/aks-engine/issues/1487))
- remove support for Kubernetes 1.9.x ([#1486](https://github.com/Azure/aks-engine/issues/1486))
- add k8s 1.14.1 for Azure Stack to VHD script ([#1483](https://github.com/Azure/aks-engine/issues/1483))

### Revert Change ◀️
- Revert "feat: support CoreOS-like Linux distributions ([#1560](https://github.com/Azure/aks-engine/issues/1560))" ([#1571](https://github.com/Azure/aks-engine/issues/1571))

### Security Fix 🛡️
- new VHD images with kernel patch ([#1497](https://github.com/Azure/aks-engine/issues/1497))

### Testing 💚
- improve crashing kube-system pod check ([#1579](https://github.com/Azure/aks-engine/issues/1579))
- update e2e api model definition ([#1547](https://github.com/Azure/aks-engine/issues/1547))
- validate expected nodes version during E2E ([#1524](https://github.com/Azure/aks-engine/issues/1524))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.37.5"></a>
# [v0.37.5] - 2019-07-11
### Bug Fixes 🐞
- remove availability sets in scaling up template ([#1610](https://github.com/Azure/aks-engine/issues/1610))
- do not install clear containers runtime on the VHD ([#1587](https://github.com/Azure/aks-engine/issues/1587))
- improve clear containers clean up ([#1600](https://github.com/Azure/aks-engine/issues/1600))
- --anonymous-auth and --client-ca-file in insecure kubelet config ([#1533](https://github.com/Azure/aks-engine/issues/1533))
- wait_for_apt_locks before GPU driver install ([#1538](https://github.com/Azure/aks-engine/issues/1538))
- Set encoding:gzip on cilium BPFFS mount unit ([#1514](https://github.com/Azure/aks-engine/issues/1514))
- ensure calico is properly enabled during upgrade ([#1510](https://github.com/Azure/aks-engine/issues/1510))
- --pod-max-pids upgrade to 1.14 scenarios ([#1508](https://github.com/Azure/aks-engine/issues/1508))
- check for nil masterProfile when setting FD count ([#1509](https://github.com/Azure/aks-engine/issues/1509))
- vhd pre-pulls expected flannel images tag ([#1484](https://github.com/Azure/aks-engine/issues/1484))

### Documentation 📘
- list pre-pulled versions for Azure Stack ([#1482](https://github.com/Azure/aks-engine/issues/1482))

### Features 🌈
- add support for Kubernetes 1.15.0 ([#1502](https://github.com/Azure/aks-engine/issues/1502))
- IPv6dualStack feature flag ([#1424](https://github.com/Azure/aks-engine/issues/1424))

### Maintenance 🔧
- update VHD to 2019.07.10375
- re-enable k8s 1.12.7 ([#1603](https://github.com/Azure/aks-engine/issues/1603))
- include https://fqdn:443 for AKS upgrade ([#1544](https://github.com/Azure/aks-engine/issues/1544))
- enable VHD 2019.06.26 for 16.04-LTS ([#1540](https://github.com/Azure/aks-engine/issues/1540))
- Add support for Kubernetes 1.13.7 and 1.14.3 on Azure Stack ([#1525](https://github.com/Azure/aks-engine/issues/1525))
- enable VHD 2019.06.20 ([#1513](https://github.com/Azure/aks-engine/issues/1513))
- Change Windows default image to Jun 2019 ([#1489](https://github.com/Azure/aks-engine/issues/1489))
- CoreDNS configmap is not updated during upgrades ([#1493](https://github.com/Azure/aks-engine/issues/1493))
- add k8s 1.14.1 for Azure Stack to VHD script ([#1483](https://github.com/Azure/aks-engine/issues/1483))

### Security Fix 🛡️
- new VHD images with kernel patch ([#1497](https://github.com/Azure/aks-engine/issues/1497))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.38.1"></a>
# [v0.38.1] - 2019-07-11
### Bug Fixes 🐞
- remove availability sets in scaling up template ([#1610](https://github.com/Azure/aks-engine/issues/1610))
- do not install clear containers runtime on the VHD ([#1587](https://github.com/Azure/aks-engine/issues/1587))
- improve clear containers clean up ([#1600](https://github.com/Azure/aks-engine/issues/1600))

### Maintenance 🔧
- update VHD to 2019.07.10
- re-enable k8s 1.12.7 ([#1603](https://github.com/Azure/aks-engine/issues/1603))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.38.0"></a>
# [v0.38.0] - 2019-07-05
### Bug Fixes 🐞
- do not override vnetCidr param when empty string ([#1570](https://github.com/Azure/aks-engine/issues/1570))
- use correct path for monitor script in monitor units ([#1555](https://github.com/Azure/aks-engine/issues/1555))
- UAE North and UAE Central regions are missing from certificate host names ([#1554](https://github.com/Azure/aks-engine/issues/1554))
- --anonymous-auth and --client-ca-file in insecure kubelet config ([#1533](https://github.com/Azure/aks-engine/issues/1533))
- wait_for_apt_locks before GPU driver install ([#1538](https://github.com/Azure/aks-engine/issues/1538))
- Ensure etcd starts after disk is mounted ([#1515](https://github.com/Azure/aks-engine/issues/1515)) ([#1535](https://github.com/Azure/aks-engine/issues/1535))
- download aks-engine in Azure Cloud Shell ([#1523](https://github.com/Azure/aks-engine/issues/1523))
- Set encoding:gzip on cilium BPFFS mount unit ([#1514](https://github.com/Azure/aks-engine/issues/1514))
- ensure calico is properly enabled during upgrade ([#1510](https://github.com/Azure/aks-engine/issues/1510))
- --pod-max-pids upgrade to 1.14 scenarios ([#1508](https://github.com/Azure/aks-engine/issues/1508))
- check for nil masterProfile when setting FD count ([#1509](https://github.com/Azure/aks-engine/issues/1509))
- vhd pre-pulls expected flannel images tag ([#1484](https://github.com/Azure/aks-engine/issues/1484))

### Continuous Integration 💜
- swap ensure-generated and test-style tasks ([#1501](https://github.com/Azure/aks-engine/issues/1501))

### Documentation 📘
- change accelerated networking example VMs ([#1569](https://github.com/Azure/aks-engine/issues/1569))
- make build before make clean in release docs ([#1545](https://github.com/Azure/aks-engine/issues/1545))
- add marketplace prerequisites for Azure Stack ([#1499](https://github.com/Azure/aks-engine/issues/1499))
- Link to more details on Windows images ([#1506](https://github.com/Azure/aks-engine/issues/1506))
- list pre-pulled versions for Azure Stack ([#1482](https://github.com/Azure/aks-engine/issues/1482))

### Features 🌈
- support CoreOS-like Linux distributions ([#1572](https://github.com/Azure/aks-engine/issues/1572))
- allow custom os images from external subs via Shared Image Galleries ([#913](https://github.com/Azure/aks-engine/issues/913))
- default AcceleratedNetworkingEnabledWindows to false if target platform is Azure Stack ([#1568](https://github.com/Azure/aks-engine/issues/1568))
- Set EtcdDiskSizeGB max default size to 1023 if target platform is Azure Stack ([#1558](https://github.com/Azure/aks-engine/issues/1558))
- Add checks for Metadata, AcceleratedNetworking and etcdDiskSizeGB on Azure Stack to fail fast. ([#1564](https://github.com/Azure/aks-engine/issues/1564))
- support for Application Gateway Ingress Controller ([#1442](https://github.com/Azure/aks-engine/issues/1442))
- support CoreOS-like Linux distributions ([#1560](https://github.com/Azure/aks-engine/issues/1560))
- set ipam environment to mas in Azure CNI config on Azure Stack ([#1542](https://github.com/Azure/aks-engine/issues/1542))
- Add defaults for useInstanceMetadata and AcceleratedNetworkingEnabled on Azure Stack ([#1549](https://github.com/Azure/aks-engine/issues/1549))
- Set PlatformFaultDomainCount value as 3 for Azure Stack ([#1529](https://github.com/Azure/aks-engine/issues/1529))
- Add vscode containerized development config ([#1526](https://github.com/Azure/aks-engine/issues/1526))
- enable deployment of windows nodes on Azure Stack ([#1456](https://github.com/Azure/aks-engine/issues/1456))
- add support for Kubernetes 1.15.0 ([#1502](https://github.com/Azure/aks-engine/issues/1502))
- Adding Windows Server version 1903 support ([#1464](https://github.com/Azure/aks-engine/issues/1464))
- IPv6dualStack feature flag ([#1424](https://github.com/Azure/aks-engine/issues/1424))

### Maintenance 🔧
- update VHD to 2019.07.01 ([#1565](https://github.com/Azure/aks-engine/issues/1565))
- replace deprecated extensions/v1beta APIs for 1.16 ([#1527](https://github.com/Azure/aks-engine/issues/1527))
- include https://fqdn:443 for AKS upgrade ([#1544](https://github.com/Azure/aks-engine/issues/1544))
- enable VHD 2019.06.26 for 16.04-LTS ([#1540](https://github.com/Azure/aks-engine/issues/1540))
- Update to MS Moby 3.0.6 ([#1537](https://github.com/Azure/aks-engine/issues/1537))
- Add support for Kubernetes 1.13.7 and 1.14.3 on Azure Stack ([#1525](https://github.com/Azure/aks-engine/issues/1525))
- enable VHD 2019.06.20 ([#1513](https://github.com/Azure/aks-engine/issues/1513))
- update addon-resizer to 1.8.5 for k8s 1.15 ([#1511](https://github.com/Azure/aks-engine/issues/1511))
- Change Windows default image to Jun 2019 ([#1489](https://github.com/Azure/aks-engine/issues/1489))
- update templates_generated.go ([#1500](https://github.com/Azure/aks-engine/issues/1500))
- CoreDNS configmap is not updated during upgrades ([#1493](https://github.com/Azure/aks-engine/issues/1493))
- update go toolchain to 1.12.6 ([#1487](https://github.com/Azure/aks-engine/issues/1487))
- remove support for Kubernetes 1.9.x ([#1486](https://github.com/Azure/aks-engine/issues/1486))
- add k8s 1.14.1 for Azure Stack to VHD script ([#1483](https://github.com/Azure/aks-engine/issues/1483))

### Revert Change ◀️
- Revert "feat: support CoreOS-like Linux distributions ([#1560](https://github.com/Azure/aks-engine/issues/1560))" ([#1571](https://github.com/Azure/aks-engine/issues/1571))

### Security Fix 🛡️
- new VHD images with kernel patch ([#1497](https://github.com/Azure/aks-engine/issues/1497))

### Testing 💚
- improve crashing kube-system pod check ([#1579](https://github.com/Azure/aks-engine/issues/1579))
- update e2e api model definition ([#1547](https://github.com/Azure/aks-engine/issues/1547))
- validate expected nodes version during E2E ([#1524](https://github.com/Azure/aks-engine/issues/1524))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.37.4"></a>
# [v0.37.4] - 2019-06-27
### Bug Fixes 🐞
- --anonymous-auth and --client-ca-file in insecure kubelet config ([#1533](https://github.com/Azure/aks-engine/issues/1533))
- wait_for_apt_locks before GPU driver install ([#1538](https://github.com/Azure/aks-engine/issues/1538))
- Set encoding:gzip on cilium BPFFS mount unit ([#1514](https://github.com/Azure/aks-engine/issues/1514))
- ensure calico is properly enabled during upgrade ([#1510](https://github.com/Azure/aks-engine/issues/1510))
- --pod-max-pids upgrade to 1.14 scenarios ([#1508](https://github.com/Azure/aks-engine/issues/1508))
- check for nil masterProfile when setting FD count ([#1509](https://github.com/Azure/aks-engine/issues/1509))
- vhd pre-pulls expected flannel images tag ([#1484](https://github.com/Azure/aks-engine/issues/1484))
- enableEncryptionWithExternalKms with master VMSS missing objectID ([#1194](https://github.com/Azure/aks-engine/issues/1194))
- start Docker service before the Kubelet service on Windows node ([#1362](https://github.com/Azure/aks-engine/issues/1362))
- remove walinuxagent version pin ([#1473](https://github.com/Azure/aks-engine/issues/1473))
- fix race condition between cloud init walinuxagent hold and CSE ([#1471](https://github.com/Azure/aks-engine/issues/1471))
- fix missing variable in cse ([#1460](https://github.com/Azure/aks-engine/issues/1460))
- reconcile templates_generated.go ([#1447](https://github.com/Azure/aks-engine/issues/1447))
- use reconcile mode in addon specs ([#1401](https://github.com/Azure/aks-engine/issues/1401))
- pause cluster-autoscaler and use VMSS capacity when upgrading ([#1245](https://github.com/Azure/aks-engine/issues/1245))
- prune addon defaults if the addon is disabled ([#1402](https://github.com/Azure/aks-engine/issues/1402))
- disable redundant cosmosdb regions ([#1403](https://github.com/Azure/aks-engine/issues/1403))
- ensure cluster-autoscaler image gets updated during upgrades ([#1385](https://github.com/Azure/aks-engine/issues/1385))
- Update Kata download url ([#1390](https://github.com/Azure/aks-engine/issues/1390))
- remove "allow-privileged" in kubelet 1.15.0 ([#1369](https://github.com/Azure/aks-engine/issues/1369))
- update etcd and containerd versions during upgrade ([#1360](https://github.com/Azure/aks-engine/issues/1360))
- correction of file/path inversion args for custom yaml manifests ([#1367](https://github.com/Azure/aks-engine/issues/1367))
- do not set distro to VHD for US Gov and German cloud ([#1357](https://github.com/Azure/aks-engine/issues/1357))
- never use 3.0.4 for moby-cli and update moby on upgrade ([#1359](https://github.com/Azure/aks-engine/issues/1359))
- decrease default host MTU for Azure Stack ([#1346](https://github.com/Azure/aks-engine/issues/1346))
- revert previous load balancer changes for Azure Stack ([#1347](https://github.com/Azure/aks-engine/issues/1347))
- PSP using GetAddonScript func ([#1290](https://github.com/Azure/aks-engine/issues/1290))
- Base64 encode sp password for windows vmss ([#1327](https://github.com/Azure/aks-engine/issues/1327))
- get vmss node version from K8S API instead of tags for instances on old model ([#1299](https://github.com/Azure/aks-engine/issues/1299))
- enable Windows plus custom VNET ([#1314](https://github.com/Azure/aks-engine/issues/1314))
- USER_ASSIGNED_IDENTITY_ID is empty in azure.json ([#1283](https://github.com/Azure/aks-engine/issues/1283))
- don’t add etcd data disk in cosmos etcd scenario ([#1310](https://github.com/Azure/aks-engine/issues/1310))
- remove grub file permissions enforcement ([#1308](https://github.com/Azure/aks-engine/issues/1308))
- increase upgrade timeout to 180 minutes ([#1300](https://github.com/Azure/aks-engine/issues/1300))
- commit generated go-bindata files with "--no-compress" option ([#1088](https://github.com/Azure/aks-engine/issues/1088))
- Ensure pods scheduled onto new nodes during upgrade respect the original node's labels/taints ([#1044](https://github.com/Azure/aks-engine/issues/1044))
- add validation for non-support of prometheus extension for Windows ([#1259](https://github.com/Azure/aks-engine/issues/1259))
- update short hyperkube commands in manifests for k8s components ([#1279](https://github.com/Azure/aks-engine/issues/1279))
- removing PodSecurityPolicy files from manifests folders ([#1257](https://github.com/Azure/aks-engine/issues/1257))
- aks-engine deploy tutorial errors out for auth method as CLI ([#1263](https://github.com/Azure/aks-engine/issues/1263))
- remove outbound connectivity validation on disconnected Azure Stack stamps ([#1250](https://github.com/Azure/aks-engine/issues/1250))
- default kubelet flags windows/linux reconciliation w/ unit tests ([#1244](https://github.com/Azure/aks-engine/issues/1244))
- Windows Kubelet issues error due to unsupported config ([#1240](https://github.com/Azure/aks-engine/issues/1240))
- private cluster with VMSS masters, jumpbox ([#1226](https://github.com/Azure/aks-engine/issues/1226))
- add Azure China support for Calico DaemonSet ([#1089](https://github.com/Azure/aks-engine/issues/1089))
- **auth:** Explicitly import the AAD library to allow cluster upgrades to succeed. ([#1474](https://github.com/Azure/aks-engine/issues/1474))

### Build 🏭
- validate deps before 'make build', and some cleanup ([#1271](https://github.com/Azure/aks-engine/issues/1271))

### Code Refactoring 💎
- run full validation for aks-engine deploy ([#1429](https://github.com/Azure/aks-engine/issues/1429))
- simplify default addons enabled enforcement ([#1409](https://github.com/Azure/aks-engine/issues/1409))
- triage untestable DCOS code ([#1351](https://github.com/Azure/aks-engine/issues/1351))
- remove _Promo from static SKU lists ([#1216](https://github.com/Azure/aks-engine/issues/1216))
- rename AKS VHD distros ([#1223](https://github.com/Azure/aks-engine/issues/1223))

### Code Style 🎶
- only call validate if validate bool is true in apiloader ([#1430](https://github.com/Azure/aks-engine/issues/1430))
- shell commands can only exit with status 0-255 ([#1336](https://github.com/Azure/aks-engine/issues/1336))
- don't loop over 'find' output ([#1329](https://github.com/Azure/aks-engine/issues/1329))
- check for unused parameters ([#1323](https://github.com/Azure/aks-engine/issues/1323))
- use 'cmd foo' instead of 'cmd $(echo foo)' ([#1321](https://github.com/Azure/aks-engine/issues/1321))
- don't use `tr` to replace words ([#1316](https://github.com/Azure/aks-engine/issues/1316))
- check exit code directly, not indirectly with $? ([#1256](https://github.com/Azure/aks-engine/issues/1256))
- standardize makefile syntax ([#1266](https://github.com/Azure/aks-engine/issues/1266))
- declare and assign separately to avoid masking return values ([#1241](https://github.com/Azure/aks-engine/issues/1241))
- use single quotes in trapping exit codes ([#1242](https://github.com/Azure/aks-engine/issues/1242))
- bash parameter expansion replaces "echo | sed" ([#1234](https://github.com/Azure/aks-engine/issues/1234))
- remove unneeded 'echo' commands ([#1230](https://github.com/Azure/aks-engine/issues/1230))
- sudo doesn't affect redirects ([#1213](https://github.com/Azure/aks-engine/issues/1213))
- replace unneeded 'cat' commands ([#1202](https://github.com/Azure/aks-engine/issues/1202))
- expressions don't expand in single quotes ([#1201](https://github.com/Azure/aks-engine/issues/1201))
- remove quotes that accidentally unquote ([#1200](https://github.com/Azure/aks-engine/issues/1200))
- don't use shell variables in printf format string ([#1197](https://github.com/Azure/aks-engine/issues/1197))
- use 'grep -c' instead of 'grep|wc' ([#1191](https://github.com/Azure/aks-engine/issues/1191))
- fix suspicious unquoted literal strings ([#1187](https://github.com/Azure/aks-engine/issues/1187))
- use builtin 'command -v' instead of nonstandard 'which' ([#1185](https://github.com/Azure/aks-engine/issues/1185))
- use 'grep -E' instead of deprecated 'egrep' ([#1186](https://github.com/Azure/aks-engine/issues/1186))
- use $(...) shell notation instead of legacy backticks ([#1180](https://github.com/Azure/aks-engine/issues/1180))

### Continuous Integration 💜
- permit optional scope field when generating changelog ([#1196](https://github.com/Azure/aks-engine/issues/1196))

### Documentation 📘
- list pre-pulled versions for Azure Stack ([#1482](https://github.com/Azure/aks-engine/issues/1482))
- docs and log messages about upgrading components ([#1396](https://github.com/Azure/aks-engine/issues/1396))
- clean up Azure Stack examples ([#1418](https://github.com/Azure/aks-engine/issues/1418))
- sample extensions workaround for Azure Stack ([#1410](https://github.com/Azure/aks-engine/issues/1410))
- Azure Stack doc page ([#1371](https://github.com/Azure/aks-engine/issues/1371))
- update custom vnet documentation ([#1399](https://github.com/Azure/aks-engine/issues/1399))
- update issue templates ([#1382](https://github.com/Azure/aks-engine/issues/1382))
- update prow docs ([#1364](https://github.com/Azure/aks-engine/issues/1364))
- update quickstart.md ([#1341](https://github.com/Azure/aks-engine/issues/1341))
- update tutorial docs to use kubernetes instead of swarm ([#1337](https://github.com/Azure/aks-engine/issues/1337))
- add customVMTags to clusterdefinitions documentation ([#1332](https://github.com/Azure/aks-engine/issues/1332))
- update group name and corresponding output folder names ([#1280](https://github.com/Azure/aks-engine/issues/1280))
- Clarify terminology around patches and use of windows-patches extension ([#1309](https://github.com/Azure/aks-engine/issues/1309))
- fix typos in launch.json debug example ([#1294](https://github.com/Azure/aks-engine/issues/1294))
- add a more complete VS Code debug configuration ([#1275](https://github.com/Azure/aks-engine/issues/1275))
- Update k8s version references in docs/samples ([#1264](https://github.com/Azure/aks-engine/issues/1264))
- correct "--auth-method" arguments in help ([#1272](https://github.com/Azure/aks-engine/issues/1272))
- improve developer debugging instructions ([#1211](https://github.com/Azure/aks-engine/issues/1211))

### Features 🌈
- add support for Kubernetes 1.15.0 ([#1502](https://github.com/Azure/aks-engine/issues/1502))
- IPv6dualStack feature flag ([#1424](https://github.com/Azure/aks-engine/issues/1424))
- add support for Kubernetes 1.15.0-rc.1  ([#1469](https://github.com/Azure/aks-engine/issues/1469))
- enable PSP w/ privileged ClusterRoleBinding for 1.15 ([#1454](https://github.com/Azure/aks-engine/issues/1454))
- add support for Kubernetes 1.15.0-beta.2 ([#1438](https://github.com/Azure/aks-engine/issues/1438))
- add support for Kubernetes 1.14.3 ([#1439](https://github.com/Azure/aks-engine/issues/1439))
- add support for Kubernetes 1.13.7 ([#1441](https://github.com/Azure/aks-engine/issues/1441))
- use all VMAS fault domains in a region ([#1090](https://github.com/Azure/aks-engine/issues/1090))
- add support for Kubernetes 1.12.8 and 1.13.5 on Azure Stack ([#1419](https://github.com/Azure/aks-engine/issues/1419))
- add support for Kubernetes 1.15.0-beta.1 ([#1394](https://github.com/Azure/aks-engine/issues/1394))
- add support for Kubernetes 1.12.9 ([#1383](https://github.com/Azure/aks-engine/issues/1383))
- added support for direct array assignment using --set ([#709](https://github.com/Azure/aks-engine/issues/709))
- allow "client_secret" auth method with ADFS identity provider ([#1343](https://github.com/Azure/aks-engine/issues/1343))
- add support for Kubernetes 1.14.2 ([#1315](https://github.com/Azure/aks-engine/issues/1315))
- custom tags on VMs and scale sets ([#1277](https://github.com/Azure/aks-engine/issues/1277))
- Make cordon drain timeout configurable with --upgrade ([#1276](https://github.com/Azure/aks-engine/issues/1276))
- add support for Kubernetes 1.13.6 ([#1262](https://github.com/Azure/aks-engine/issues/1262))
- add support for Kubernetes 1.15.0-alpha.3 ([#1247](https://github.com/Azure/aks-engine/issues/1247))
- disable unsupported addons on Azure Stack deployments ([#1233](https://github.com/Azure/aks-engine/issues/1233))
- remove azurefile storage class for Azure Stack ([#1222](https://github.com/Azure/aks-engine/issues/1222))
- add auditd as an ubuntu option ([#1143](https://github.com/Azure/aks-engine/issues/1143))
- add support for Kubernetes 1.11.10 ([#1193](https://github.com/Azure/aks-engine/issues/1193))
- add two US DoD Azure locations ([#1205](https://github.com/Azure/aks-engine/issues/1205))

### Maintenance 🔧
- include https://fqdn:443 for AKS upgrade ([#1544](https://github.com/Azure/aks-engine/issues/1544))
- enable VHD 2019.06.26 for 16.04-LTS ([#1540](https://github.com/Azure/aks-engine/issues/1540))
- Add support for Kubernetes 1.13.7 and 1.14.3 on Azure Stack ([#1525](https://github.com/Azure/aks-engine/issues/1525))
- enable VHD 2019.06.20 ([#1513](https://github.com/Azure/aks-engine/issues/1513))
- Change Windows default image to Jun 2019 ([#1489](https://github.com/Azure/aks-engine/issues/1489))
- CoreDNS configmap is not updated during upgrades ([#1493](https://github.com/Azure/aks-engine/issues/1493))
- add k8s 1.14.1 for Azure Stack to VHD script ([#1483](https://github.com/Azure/aks-engine/issues/1483))
- update VHD to 2019.06.12 ([#1481](https://github.com/Azure/aks-engine/issues/1481))
- add note about apimodel in issue template ([#1478](https://github.com/Azure/aks-engine/issues/1478))
- delete obsolete calico spec ([#1463](https://github.com/Azure/aks-engine/issues/1463))
- reconcile generated files ([#1457](https://github.com/Azure/aks-engine/issues/1457))
- update omsagent addon to use the latest version ([#1156](https://github.com/Azure/aks-engine/issues/1156))
- enable only strong TLS cipher suites for kubelet by default ([#1434](https://github.com/Azure/aks-engine/issues/1434))
- update etcd default to 3.2.26 ([#1451](https://github.com/Azure/aks-engine/issues/1451))
- enable 2019.06.08 VHD ([#1450](https://github.com/Azure/aks-engine/issues/1450))
- move azure stack example VM SKUs from Standard_D2_v2 to Standard_D2_v3 ([#1448](https://github.com/Azure/aks-engine/issues/1448))
- move examples VM SKUs from Standard_D2_v2 to Standard_D2_v3 ([#1436](https://github.com/Azure/aks-engine/issues/1436))
- use CoreDNS 1.5.0 for k8s >= 1.12.0 ([#1411](https://github.com/Azure/aks-engine/issues/1411))
- re-enable k8s v1.12.9 ([#1428](https://github.com/Azure/aks-engine/issues/1428))
- deprecate unused CSE func installDockerEngine ([#1199](https://github.com/Azure/aks-engine/issues/1199))
- ignore junit.xml from e2e ([#1417](https://github.com/Azure/aks-engine/issues/1417))
- add security to changelog commit titles ([#1415](https://github.com/Azure/aks-engine/issues/1415))
- update generated files ([#1405](https://github.com/Azure/aks-engine/issues/1405))
- add support for etcd 3.2.26 ([#1384](https://github.com/Azure/aks-engine/issues/1384))
- ensure all Linux-bound files are LF (not CRLF) ([#1354](https://github.com/Azure/aks-engine/issues/1354))
- bump cluster-autoscaler to latest patch versions ([#1339](https://github.com/Azure/aks-engine/issues/1339))
- update go toolchain to 1.12.5 ([#1305](https://github.com/Azure/aks-engine/issues/1305))
- remove unused go template funcs ([#1331](https://github.com/Azure/aks-engine/issues/1331))
- add example cluster definition for Azure Stack ([#1304](https://github.com/Azure/aks-engine/issues/1304))
- update VHD image to 2019.05.16 ([#1319](https://github.com/Azure/aks-engine/issues/1319))
- Update calico to v3.7.2 ([#1293](https://github.com/Azure/aks-engine/issues/1293))
- update templates_generated.go ([#1317](https://github.com/Azure/aks-engine/issues/1317))
- update templates_generated.go ([#1295](https://github.com/Azure/aks-engine/issues/1295))
- enable VHD 2019.05.08, disable auditd for non-VHD ([#1286](https://github.com/Azure/aks-engine/issues/1286))
- update client-go and k8s.io vendored code ([#1273](https://github.com/Azure/aks-engine/issues/1273))
- rename addons-related types and consts ([#1265](https://github.com/Azure/aks-engine/issues/1265))
- deprecate CIS files cloud-init paving for non-VHD distros ([#1251](https://github.com/Azure/aks-engine/issues/1251))
- remove support for Kubernetes 1.11.8 ([#1243](https://github.com/Azure/aks-engine/issues/1243))
- enable moby 3.0.5, and set 3.0.5 to default ([#1236](https://github.com/Azure/aks-engine/issues/1236))
- update Azure CNI version to v1.0.22 ([#1192](https://github.com/Azure/aks-engine/issues/1192))
- update VHD version to 04.30.2019 ([#1184](https://github.com/Azure/aks-engine/issues/1184))
- onboard Azure CNI images to MCR ([#1153](https://github.com/Azure/aks-engine/issues/1153))
- **CIS:** Remove log file permission enforcement  ([#1412](https://github.com/Azure/aks-engine/issues/1412))
- **CIS:** generate kubelet server cert in aks-e ([#1416](https://github.com/Azure/aks-engine/issues/1416))

### Revert Change ◀️
- remove k8s 1.12.9 support ([#1476](https://github.com/Azure/aks-engine/issues/1476))
- add support for Kubernetes 1.12.9 ([#1420](https://github.com/Azure/aks-engine/issues/1420))
- /etc/default/grub changes aren't reconcilable over time ([#1237](https://github.com/Azure/aks-engine/issues/1237))

### Security Fix 🛡️
- new VHD images with kernel patch ([#1497](https://github.com/Azure/aks-engine/issues/1497))
- remove support for 1.13.6 and 1.14.2 ([#1413](https://github.com/Azure/aks-engine/issues/1413))

### Testing 💚
- HPA scale down config only works if k8s >= 1.12 ([#1480](https://github.com/Azure/aks-engine/issues/1480))
- add subnets to route table in E2E for custom vnet + kubenet ([#1470](https://github.com/Azure/aks-engine/issues/1470))
- parallelize tests part 1 ([#1455](https://github.com/Azure/aks-engine/issues/1455))
- enable DNS liveness test for calico-enabled clusters ([#1459](https://github.com/Azure/aks-engine/issues/1459))
- Switch Windows test passes to VMSS by default to match AKS ([#1452](https://github.com/Azure/aks-engine/issues/1452))
- getTemplateFuncMap unit tests part 5 ([#1449](https://github.com/Azure/aks-engine/issues/1449))
- getTemplateFuncMap unit tests part 4 ([#1388](https://github.com/Azure/aks-engine/issues/1388))
- quicker “defaults addons image” unit test ([#1426](https://github.com/Azure/aks-engine/issues/1426))
- update unit test payload for Azure Stack ([#1408](https://github.com/Azure/aks-engine/issues/1408))
- enable aks-engine e2e test on azurestack ([#1397](https://github.com/Azure/aks-engine/issues/1397))
- refactor addons default image test ([#1407](https://github.com/Azure/aks-engine/issues/1407))
- getTemplateFuncMap unit tests part 3 ([#1344](https://github.com/Azure/aks-engine/issues/1344))
- getTemplateFuncMap unit tests part 2 ([#1338](https://github.com/Azure/aks-engine/issues/1338))
- sorted GetKubernetesLabels method for deterministic unit tests ([#1340](https://github.com/Azure/aks-engine/issues/1340))
- getTemplateFuncMap unit tests part 1 ([#1334](https://github.com/Azure/aks-engine/issues/1334))
- add prototype for getTemplateFuncMap unit tests ([#1333](https://github.com/Azure/aks-engine/issues/1333))
- unit tests for kubernetesManifestSettingsInit ([#1312](https://github.com/Azure/aks-engine/issues/1312))
- add masterSSHPort to auditd E2E SCP ([#1307](https://github.com/Azure/aks-engine/issues/1307))
- fix auditd E2E for master VMSS ([#1297](https://github.com/Azure/aks-engine/issues/1297))
- add storage + Azure Stack addons unit tests ([#1285](https://github.com/Azure/aks-engine/issues/1285))
- reduced name for ilb curl deployment ([#1274](https://github.com/Azure/aks-engine/issues/1274))
- add no outbound CSE unit test ([#1254](https://github.com/Azure/aks-engine/issues/1254))
- test for missing kubelet configs as well ([#1252](https://github.com/Azure/aks-engine/issues/1252))
- idempotent e2e changes for soak test scenarios ([#1214](https://github.com/Azure/aks-engine/issues/1214))
- enable all ssh to master tests for vmss masters ([#1198](https://github.com/Azure/aks-engine/issues/1198))
- hasSSHAbleMaster returns false if master vmss ([#1195](https://github.com/Azure/aks-engine/issues/1195))
- can’t ssh into master vmss master vms ([#1188](https://github.com/Azure/aks-engine/issues/1188))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.35.6"></a>
# [v0.35.6] - 2019-06-27
### Bug Fixes 🐞
- wait_for_apt_locks before GPU driver install ([#1538](https://github.com/Azure/aks-engine/issues/1538))
- ensure /etc/default/grub has no CRLF
- Base64 encode sp password for windows vmss ([#1327](https://github.com/Azure/aks-engine/issues/1327))
- enableEncryptionWithExternalKms with master VMSS missing objectID ([#1194](https://github.com/Azure/aks-engine/issues/1194))
- add Azure China support for Calico DaemonSet ([#1089](https://github.com/Azure/aks-engine/issues/1089))

### Features 🌈
- add two US DoD Azure locations ([#1205](https://github.com/Azure/aks-engine/issues/1205))

### Maintenance 🔧
- add VHD 2019.06.26335 for AKS 16.04-LTS
- rev the SKU as well
- remove temporary CRLF fix
- enable 2019.05.23 VHD versions for v0.35
- fix all CRLF system files
- move to fix_crlf func
- move dos2unix install earlier in flow
- update VHD version to 04.30.2019 ([#1184](https://github.com/Azure/aks-engine/issues/1184))
- update omsagent addon to use the latest version ([#1156](https://github.com/Azure/aks-engine/issues/1156))

### Revert Change ◀️
- /etc/default/grub changes aren't reconcilable over time

### Security Fix 🛡️
- update VHD images with kernel patch

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.37.2"></a>
# [v0.37.2] - 2019-06-21
#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.37.3"></a>
# [v0.37.3] - 2019-06-19
### Bug Fixes 🐞
- Set encoding:gzip on cilium BPFFS mount unit ([#1514](https://github.com/Azure/aks-engine/issues/1514))
- ensure calico is properly enabled during upgrade ([#1510](https://github.com/Azure/aks-engine/issues/1510))
- --pod-max-pids upgrade to 1.14 scenarios ([#1508](https://github.com/Azure/aks-engine/issues/1508))
- check for nil masterProfile when setting FD count ([#1509](https://github.com/Azure/aks-engine/issues/1509))

### Features 🌈
- add support for Kubernetes 1.15.0 ([#1502](https://github.com/Azure/aks-engine/issues/1502))

### Maintenance 🔧
- enable VHD 2019.06.20 ([#1513](https://github.com/Azure/aks-engine/issues/1513))
- Change Windows default image to Jun 2019 ([#1489](https://github.com/Azure/aks-engine/issues/1489))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.37.1"></a>
# [v0.37.1] - 2019-06-19
### Bug Fixes 🐞
- vhd pre-pulls expected flannel images tag ([#1484](https://github.com/Azure/aks-engine/issues/1484))
- enableEncryptionWithExternalKms with master VMSS missing objectID ([#1194](https://github.com/Azure/aks-engine/issues/1194))
- start Docker service before the Kubelet service on Windows node ([#1362](https://github.com/Azure/aks-engine/issues/1362))
- remove walinuxagent version pin ([#1473](https://github.com/Azure/aks-engine/issues/1473))
- fix race condition between cloud init walinuxagent hold and CSE ([#1471](https://github.com/Azure/aks-engine/issues/1471))
- fix missing variable in cse ([#1460](https://github.com/Azure/aks-engine/issues/1460))
- reconcile templates_generated.go ([#1447](https://github.com/Azure/aks-engine/issues/1447))
- use reconcile mode in addon specs ([#1401](https://github.com/Azure/aks-engine/issues/1401))
- pause cluster-autoscaler and use VMSS capacity when upgrading ([#1245](https://github.com/Azure/aks-engine/issues/1245))
- prune addon defaults if the addon is disabled ([#1402](https://github.com/Azure/aks-engine/issues/1402))
- disable redundant cosmosdb regions ([#1403](https://github.com/Azure/aks-engine/issues/1403))
- ensure cluster-autoscaler image gets updated during upgrades ([#1385](https://github.com/Azure/aks-engine/issues/1385))
- Update Kata download url ([#1390](https://github.com/Azure/aks-engine/issues/1390))
- remove "allow-privileged" in kubelet 1.15.0 ([#1369](https://github.com/Azure/aks-engine/issues/1369))
- update etcd and containerd versions during upgrade ([#1360](https://github.com/Azure/aks-engine/issues/1360))
- correction of file/path inversion args for custom yaml manifests ([#1367](https://github.com/Azure/aks-engine/issues/1367))
- do not set distro to VHD for US Gov and German cloud ([#1357](https://github.com/Azure/aks-engine/issues/1357))
- never use 3.0.4 for moby-cli and update moby on upgrade ([#1359](https://github.com/Azure/aks-engine/issues/1359))
- decrease default host MTU for Azure Stack ([#1346](https://github.com/Azure/aks-engine/issues/1346))
- revert previous load balancer changes for Azure Stack ([#1347](https://github.com/Azure/aks-engine/issues/1347))
- PSP using GetAddonScript func ([#1290](https://github.com/Azure/aks-engine/issues/1290))
- Base64 encode sp password for windows vmss ([#1327](https://github.com/Azure/aks-engine/issues/1327))
- get vmss node version from K8S API instead of tags for instances on old model ([#1299](https://github.com/Azure/aks-engine/issues/1299))
- enable Windows plus custom VNET ([#1314](https://github.com/Azure/aks-engine/issues/1314))
- USER_ASSIGNED_IDENTITY_ID is empty in azure.json ([#1283](https://github.com/Azure/aks-engine/issues/1283))
- don’t add etcd data disk in cosmos etcd scenario ([#1310](https://github.com/Azure/aks-engine/issues/1310))
- remove grub file permissions enforcement ([#1308](https://github.com/Azure/aks-engine/issues/1308))
- increase upgrade timeout to 180 minutes ([#1300](https://github.com/Azure/aks-engine/issues/1300))
- commit generated go-bindata files with "--no-compress" option ([#1088](https://github.com/Azure/aks-engine/issues/1088))
- Ensure pods scheduled onto new nodes during upgrade respect the original node's labels/taints ([#1044](https://github.com/Azure/aks-engine/issues/1044))
- add validation for non-support of prometheus extension for Windows ([#1259](https://github.com/Azure/aks-engine/issues/1259))
- update short hyperkube commands in manifests for k8s components ([#1279](https://github.com/Azure/aks-engine/issues/1279))
- removing PodSecurityPolicy files from manifests folders ([#1257](https://github.com/Azure/aks-engine/issues/1257))
- aks-engine deploy tutorial errors out for auth method as CLI ([#1263](https://github.com/Azure/aks-engine/issues/1263))
- remove outbound connectivity validation on disconnected Azure Stack stamps ([#1250](https://github.com/Azure/aks-engine/issues/1250))
- default kubelet flags windows/linux reconciliation w/ unit tests ([#1244](https://github.com/Azure/aks-engine/issues/1244))
- Windows Kubelet issues error due to unsupported config ([#1240](https://github.com/Azure/aks-engine/issues/1240))
- private cluster with VMSS masters, jumpbox ([#1226](https://github.com/Azure/aks-engine/issues/1226))
- add Azure China support for Calico DaemonSet ([#1089](https://github.com/Azure/aks-engine/issues/1089))
- **auth:** Explicitly import the AAD library to allow cluster upgrades to succeed. ([#1474](https://github.com/Azure/aks-engine/issues/1474))

### Build 🏭
- validate deps before 'make build', and some cleanup ([#1271](https://github.com/Azure/aks-engine/issues/1271))

### Code Refactoring 💎
- run full validation for aks-engine deploy ([#1429](https://github.com/Azure/aks-engine/issues/1429))
- simplify default addons enabled enforcement ([#1409](https://github.com/Azure/aks-engine/issues/1409))
- triage untestable DCOS code ([#1351](https://github.com/Azure/aks-engine/issues/1351))
- remove _Promo from static SKU lists ([#1216](https://github.com/Azure/aks-engine/issues/1216))
- rename AKS VHD distros ([#1223](https://github.com/Azure/aks-engine/issues/1223))

### Code Style 🎶
- only call validate if validate bool is true in apiloader ([#1430](https://github.com/Azure/aks-engine/issues/1430))
- shell commands can only exit with status 0-255 ([#1336](https://github.com/Azure/aks-engine/issues/1336))
- don't loop over 'find' output ([#1329](https://github.com/Azure/aks-engine/issues/1329))
- check for unused parameters ([#1323](https://github.com/Azure/aks-engine/issues/1323))
- use 'cmd foo' instead of 'cmd $(echo foo)' ([#1321](https://github.com/Azure/aks-engine/issues/1321))
- don't use `tr` to replace words ([#1316](https://github.com/Azure/aks-engine/issues/1316))
- check exit code directly, not indirectly with $? ([#1256](https://github.com/Azure/aks-engine/issues/1256))
- standardize makefile syntax ([#1266](https://github.com/Azure/aks-engine/issues/1266))
- declare and assign separately to avoid masking return values ([#1241](https://github.com/Azure/aks-engine/issues/1241))
- use single quotes in trapping exit codes ([#1242](https://github.com/Azure/aks-engine/issues/1242))
- bash parameter expansion replaces "echo | sed" ([#1234](https://github.com/Azure/aks-engine/issues/1234))
- remove unneeded 'echo' commands ([#1230](https://github.com/Azure/aks-engine/issues/1230))
- sudo doesn't affect redirects ([#1213](https://github.com/Azure/aks-engine/issues/1213))
- replace unneeded 'cat' commands ([#1202](https://github.com/Azure/aks-engine/issues/1202))
- expressions don't expand in single quotes ([#1201](https://github.com/Azure/aks-engine/issues/1201))
- remove quotes that accidentally unquote ([#1200](https://github.com/Azure/aks-engine/issues/1200))
- don't use shell variables in printf format string ([#1197](https://github.com/Azure/aks-engine/issues/1197))
- use 'grep -c' instead of 'grep|wc' ([#1191](https://github.com/Azure/aks-engine/issues/1191))
- fix suspicious unquoted literal strings ([#1187](https://github.com/Azure/aks-engine/issues/1187))
- use builtin 'command -v' instead of nonstandard 'which' ([#1185](https://github.com/Azure/aks-engine/issues/1185))
- use 'grep -E' instead of deprecated 'egrep' ([#1186](https://github.com/Azure/aks-engine/issues/1186))
- use $(...) shell notation instead of legacy backticks ([#1180](https://github.com/Azure/aks-engine/issues/1180))

### Continuous Integration 💜
- permit optional scope field when generating changelog ([#1196](https://github.com/Azure/aks-engine/issues/1196))

### Documentation 📘
- list pre-pulled versions for Azure Stack ([#1482](https://github.com/Azure/aks-engine/issues/1482))
- docs and log messages about upgrading components ([#1396](https://github.com/Azure/aks-engine/issues/1396))
- clean up Azure Stack examples ([#1418](https://github.com/Azure/aks-engine/issues/1418))
- sample extensions workaround for Azure Stack ([#1410](https://github.com/Azure/aks-engine/issues/1410))
- Azure Stack doc page ([#1371](https://github.com/Azure/aks-engine/issues/1371))
- update custom vnet documentation ([#1399](https://github.com/Azure/aks-engine/issues/1399))
- update issue templates ([#1382](https://github.com/Azure/aks-engine/issues/1382))
- update prow docs ([#1364](https://github.com/Azure/aks-engine/issues/1364))
- update quickstart.md ([#1341](https://github.com/Azure/aks-engine/issues/1341))
- update tutorial docs to use kubernetes instead of swarm ([#1337](https://github.com/Azure/aks-engine/issues/1337))
- add customVMTags to clusterdefinitions documentation ([#1332](https://github.com/Azure/aks-engine/issues/1332))
- update group name and corresponding output folder names ([#1280](https://github.com/Azure/aks-engine/issues/1280))
- Clarify terminology around patches and use of windows-patches extension ([#1309](https://github.com/Azure/aks-engine/issues/1309))
- fix typos in launch.json debug example ([#1294](https://github.com/Azure/aks-engine/issues/1294))
- add a more complete VS Code debug configuration ([#1275](https://github.com/Azure/aks-engine/issues/1275))
- Update k8s version references in docs/samples ([#1264](https://github.com/Azure/aks-engine/issues/1264))
- correct "--auth-method" arguments in help ([#1272](https://github.com/Azure/aks-engine/issues/1272))
- improve developer debugging instructions ([#1211](https://github.com/Azure/aks-engine/issues/1211))

### Features 🌈
- IPv6dualStack feature flag ([#1424](https://github.com/Azure/aks-engine/issues/1424))
- add support for Kubernetes 1.15.0-rc.1  ([#1469](https://github.com/Azure/aks-engine/issues/1469))
- enable PSP w/ privileged ClusterRoleBinding for 1.15 ([#1454](https://github.com/Azure/aks-engine/issues/1454))
- add support for Kubernetes 1.15.0-beta.2 ([#1438](https://github.com/Azure/aks-engine/issues/1438))
- add support for Kubernetes 1.14.3 ([#1439](https://github.com/Azure/aks-engine/issues/1439))
- add support for Kubernetes 1.13.7 ([#1441](https://github.com/Azure/aks-engine/issues/1441))
- use all VMAS fault domains in a region ([#1090](https://github.com/Azure/aks-engine/issues/1090))
- add support for Kubernetes 1.12.8 and 1.13.5 on Azure Stack ([#1419](https://github.com/Azure/aks-engine/issues/1419))
- add support for Kubernetes 1.15.0-beta.1 ([#1394](https://github.com/Azure/aks-engine/issues/1394))
- add support for Kubernetes 1.12.9 ([#1383](https://github.com/Azure/aks-engine/issues/1383))
- added support for direct array assignment using --set ([#709](https://github.com/Azure/aks-engine/issues/709))
- allow "client_secret" auth method with ADFS identity provider ([#1343](https://github.com/Azure/aks-engine/issues/1343))
- add support for Kubernetes 1.14.2 ([#1315](https://github.com/Azure/aks-engine/issues/1315))
- custom tags on VMs and scale sets ([#1277](https://github.com/Azure/aks-engine/issues/1277))
- Make cordon drain timeout configurable with --upgrade ([#1276](https://github.com/Azure/aks-engine/issues/1276))
- add support for Kubernetes 1.13.6 ([#1262](https://github.com/Azure/aks-engine/issues/1262))
- add support for Kubernetes 1.15.0-alpha.3 ([#1247](https://github.com/Azure/aks-engine/issues/1247))
- disable unsupported addons on Azure Stack deployments ([#1233](https://github.com/Azure/aks-engine/issues/1233))
- remove azurefile storage class for Azure Stack ([#1222](https://github.com/Azure/aks-engine/issues/1222))
- add auditd as an ubuntu option ([#1143](https://github.com/Azure/aks-engine/issues/1143))
- add support for Kubernetes 1.11.10 ([#1193](https://github.com/Azure/aks-engine/issues/1193))
- add two US DoD Azure locations ([#1205](https://github.com/Azure/aks-engine/issues/1205))

### Maintenance 🔧
- CoreDNS configmap is not updated during upgrades ([#1493](https://github.com/Azure/aks-engine/issues/1493))
- add k8s 1.14.1 for Azure Stack to VHD script ([#1483](https://github.com/Azure/aks-engine/issues/1483))
- update VHD to 2019.06.12 ([#1481](https://github.com/Azure/aks-engine/issues/1481))
- add note about apimodel in issue template ([#1478](https://github.com/Azure/aks-engine/issues/1478))
- delete obsolete calico spec ([#1463](https://github.com/Azure/aks-engine/issues/1463))
- reconcile generated files ([#1457](https://github.com/Azure/aks-engine/issues/1457))
- update omsagent addon to use the latest version ([#1156](https://github.com/Azure/aks-engine/issues/1156))
- enable only strong TLS cipher suites for kubelet by default ([#1434](https://github.com/Azure/aks-engine/issues/1434))
- update etcd default to 3.2.26 ([#1451](https://github.com/Azure/aks-engine/issues/1451))
- enable 2019.06.08 VHD ([#1450](https://github.com/Azure/aks-engine/issues/1450))
- move azure stack example VM SKUs from Standard_D2_v2 to Standard_D2_v3 ([#1448](https://github.com/Azure/aks-engine/issues/1448))
- move examples VM SKUs from Standard_D2_v2 to Standard_D2_v3 ([#1436](https://github.com/Azure/aks-engine/issues/1436))
- use CoreDNS 1.5.0 for k8s >= 1.12.0 ([#1411](https://github.com/Azure/aks-engine/issues/1411))
- re-enable k8s v1.12.9 ([#1428](https://github.com/Azure/aks-engine/issues/1428))
- deprecate unused CSE func installDockerEngine ([#1199](https://github.com/Azure/aks-engine/issues/1199))
- ignore junit.xml from e2e ([#1417](https://github.com/Azure/aks-engine/issues/1417))
- add security to changelog commit titles ([#1415](https://github.com/Azure/aks-engine/issues/1415))
- update generated files ([#1405](https://github.com/Azure/aks-engine/issues/1405))
- add support for etcd 3.2.26 ([#1384](https://github.com/Azure/aks-engine/issues/1384))
- ensure all Linux-bound files are LF (not CRLF) ([#1354](https://github.com/Azure/aks-engine/issues/1354))
- bump cluster-autoscaler to latest patch versions ([#1339](https://github.com/Azure/aks-engine/issues/1339))
- update go toolchain to 1.12.5 ([#1305](https://github.com/Azure/aks-engine/issues/1305))
- remove unused go template funcs ([#1331](https://github.com/Azure/aks-engine/issues/1331))
- add example cluster definition for Azure Stack ([#1304](https://github.com/Azure/aks-engine/issues/1304))
- update VHD image to 2019.05.16 ([#1319](https://github.com/Azure/aks-engine/issues/1319))
- Update calico to v3.7.2 ([#1293](https://github.com/Azure/aks-engine/issues/1293))
- update templates_generated.go ([#1317](https://github.com/Azure/aks-engine/issues/1317))
- update templates_generated.go ([#1295](https://github.com/Azure/aks-engine/issues/1295))
- enable VHD 2019.05.08, disable auditd for non-VHD ([#1286](https://github.com/Azure/aks-engine/issues/1286))
- update client-go and k8s.io vendored code ([#1273](https://github.com/Azure/aks-engine/issues/1273))
- rename addons-related types and consts ([#1265](https://github.com/Azure/aks-engine/issues/1265))
- deprecate CIS files cloud-init paving for non-VHD distros ([#1251](https://github.com/Azure/aks-engine/issues/1251))
- remove support for Kubernetes 1.11.8 ([#1243](https://github.com/Azure/aks-engine/issues/1243))
- enable moby 3.0.5, and set 3.0.5 to default ([#1236](https://github.com/Azure/aks-engine/issues/1236))
- update Azure CNI version to v1.0.22 ([#1192](https://github.com/Azure/aks-engine/issues/1192))
- update VHD version to 04.30.2019 ([#1184](https://github.com/Azure/aks-engine/issues/1184))
- onboard Azure CNI images to MCR ([#1153](https://github.com/Azure/aks-engine/issues/1153))
- **CIS:** Remove log file permission enforcement  ([#1412](https://github.com/Azure/aks-engine/issues/1412))
- **CIS:** generate kubelet server cert in aks-e ([#1416](https://github.com/Azure/aks-engine/issues/1416))

### Revert Change ◀️
- remove k8s 1.12.9 support ([#1476](https://github.com/Azure/aks-engine/issues/1476))
- add support for Kubernetes 1.12.9 ([#1420](https://github.com/Azure/aks-engine/issues/1420))
- /etc/default/grub changes aren't reconcilable over time ([#1237](https://github.com/Azure/aks-engine/issues/1237))

### Security Fix 🛡️
- new VHD images with kernel patch ([#1497](https://github.com/Azure/aks-engine/issues/1497))
- remove support for 1.13.6 and 1.14.2 ([#1413](https://github.com/Azure/aks-engine/issues/1413))

### Testing 💚
- HPA scale down config only works if k8s >= 1.12 ([#1480](https://github.com/Azure/aks-engine/issues/1480))
- add subnets to route table in E2E for custom vnet + kubenet ([#1470](https://github.com/Azure/aks-engine/issues/1470))
- parallelize tests part 1 ([#1455](https://github.com/Azure/aks-engine/issues/1455))
- enable DNS liveness test for calico-enabled clusters ([#1459](https://github.com/Azure/aks-engine/issues/1459))
- Switch Windows test passes to VMSS by default to match AKS ([#1452](https://github.com/Azure/aks-engine/issues/1452))
- getTemplateFuncMap unit tests part 5 ([#1449](https://github.com/Azure/aks-engine/issues/1449))
- getTemplateFuncMap unit tests part 4 ([#1388](https://github.com/Azure/aks-engine/issues/1388))
- quicker “defaults addons image” unit test ([#1426](https://github.com/Azure/aks-engine/issues/1426))
- update unit test payload for Azure Stack ([#1408](https://github.com/Azure/aks-engine/issues/1408))
- enable aks-engine e2e test on azurestack ([#1397](https://github.com/Azure/aks-engine/issues/1397))
- refactor addons default image test ([#1407](https://github.com/Azure/aks-engine/issues/1407))
- getTemplateFuncMap unit tests part 3 ([#1344](https://github.com/Azure/aks-engine/issues/1344))
- getTemplateFuncMap unit tests part 2 ([#1338](https://github.com/Azure/aks-engine/issues/1338))
- sorted GetKubernetesLabels method for deterministic unit tests ([#1340](https://github.com/Azure/aks-engine/issues/1340))
- getTemplateFuncMap unit tests part 1 ([#1334](https://github.com/Azure/aks-engine/issues/1334))
- add prototype for getTemplateFuncMap unit tests ([#1333](https://github.com/Azure/aks-engine/issues/1333))
- unit tests for kubernetesManifestSettingsInit ([#1312](https://github.com/Azure/aks-engine/issues/1312))
- add masterSSHPort to auditd E2E SCP ([#1307](https://github.com/Azure/aks-engine/issues/1307))
- fix auditd E2E for master VMSS ([#1297](https://github.com/Azure/aks-engine/issues/1297))
- add storage + Azure Stack addons unit tests ([#1285](https://github.com/Azure/aks-engine/issues/1285))
- reduced name for ilb curl deployment ([#1274](https://github.com/Azure/aks-engine/issues/1274))
- add no outbound CSE unit test ([#1254](https://github.com/Azure/aks-engine/issues/1254))
- test for missing kubelet configs as well ([#1252](https://github.com/Azure/aks-engine/issues/1252))
- idempotent e2e changes for soak test scenarios ([#1214](https://github.com/Azure/aks-engine/issues/1214))
- enable all ssh to master tests for vmss masters ([#1198](https://github.com/Azure/aks-engine/issues/1198))
- hasSSHAbleMaster returns false if master vmss ([#1195](https://github.com/Azure/aks-engine/issues/1195))
- can’t ssh into master vmss master vms ([#1188](https://github.com/Azure/aks-engine/issues/1188))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.35.5"></a>
# [v0.35.5] - 2019-06-19
### Bug Fixes 🐞
- ensure /etc/default/grub has no CRLF
- Base64 encode sp password for windows vmss ([#1327](https://github.com/Azure/aks-engine/issues/1327))
- enableEncryptionWithExternalKms with master VMSS missing objectID ([#1194](https://github.com/Azure/aks-engine/issues/1194))
- add Azure China support for Calico DaemonSet ([#1089](https://github.com/Azure/aks-engine/issues/1089))

### Features 🌈
- add two US DoD Azure locations ([#1205](https://github.com/Azure/aks-engine/issues/1205))

### Maintenance 🔧
- rev the SKU as well
- remove temporary CRLF fix
- enable 2019.05.23 VHD versions for v0.35
- fix all CRLF system files
- move to fix_crlf func
- move dos2unix install earlier in flow
- update VHD version to 04.30.2019 ([#1184](https://github.com/Azure/aks-engine/issues/1184))
- update omsagent addon to use the latest version ([#1156](https://github.com/Azure/aks-engine/issues/1156))

### Security Fix 🛡️
- update VHD images with kernel patch

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.37.0"></a>
# [v0.37.0] - 2019-06-14
### Bug Fixes 🐞
- PSP using GetAddonScript func ([#1290](https://github.com/Azure/aks-engine/issues/1290))
- start Docker service before the Kubelet service on Windows node ([#1362](https://github.com/Azure/aks-engine/issues/1362))
- remove walinuxagent version pin ([#1473](https://github.com/Azure/aks-engine/issues/1473))
- fix race condition between cloud init walinuxagent hold and CSE ([#1471](https://github.com/Azure/aks-engine/issues/1471))
- fix missing variable in cse ([#1460](https://github.com/Azure/aks-engine/issues/1460))
- reconcile templates_generated.go ([#1447](https://github.com/Azure/aks-engine/issues/1447))
- use reconcile mode in addon specs ([#1401](https://github.com/Azure/aks-engine/issues/1401))
- pause cluster-autoscaler and use VMSS capacity when upgrading ([#1245](https://github.com/Azure/aks-engine/issues/1245))
- prune addon defaults if the addon is disabled ([#1402](https://github.com/Azure/aks-engine/issues/1402))
- disable redundant cosmosdb regions ([#1403](https://github.com/Azure/aks-engine/issues/1403))
- ensure cluster-autoscaler image gets updated during upgrades ([#1385](https://github.com/Azure/aks-engine/issues/1385))
- Update Kata download url ([#1390](https://github.com/Azure/aks-engine/issues/1390))
- remove "allow-privileged" in kubelet 1.15.0 ([#1369](https://github.com/Azure/aks-engine/issues/1369))
- update etcd and containerd versions during upgrade ([#1360](https://github.com/Azure/aks-engine/issues/1360))
- correction of file/path inversion args for custom yaml manifests ([#1367](https://github.com/Azure/aks-engine/issues/1367))
- do not set distro to VHD for US Gov and German cloud ([#1357](https://github.com/Azure/aks-engine/issues/1357))
- never use 3.0.4 for moby-cli and update moby on upgrade ([#1359](https://github.com/Azure/aks-engine/issues/1359))
- decrease default host MTU for Azure Stack ([#1346](https://github.com/Azure/aks-engine/issues/1346))
- revert previous load balancer changes for Azure Stack ([#1347](https://github.com/Azure/aks-engine/issues/1347))
- **auth:** Explicitly import the AAD library to allow cluster upgrades to succeed. ([#1474](https://github.com/Azure/aks-engine/issues/1474))

### Code Refactoring 💎
- run full validation for aks-engine deploy ([#1429](https://github.com/Azure/aks-engine/issues/1429))
- simplify default addons enabled enforcement ([#1409](https://github.com/Azure/aks-engine/issues/1409))
- triage untestable DCOS code ([#1351](https://github.com/Azure/aks-engine/issues/1351))
- remove _Promo from static SKU lists ([#1216](https://github.com/Azure/aks-engine/issues/1216))
- rename AKS VHD distros ([#1223](https://github.com/Azure/aks-engine/issues/1223))

### Code Style 🎶
- only call validate if validate bool is true in apiloader ([#1430](https://github.com/Azure/aks-engine/issues/1430))
- shell commands can only exit with status 0-255 ([#1336](https://github.com/Azure/aks-engine/issues/1336))
- don't loop over 'find' output ([#1329](https://github.com/Azure/aks-engine/issues/1329))

### Documentation 📘
- docs and log messages about upgrading components ([#1396](https://github.com/Azure/aks-engine/issues/1396))
- clean up Azure Stack examples ([#1418](https://github.com/Azure/aks-engine/issues/1418))
- sample extensions workaround for Azure Stack ([#1410](https://github.com/Azure/aks-engine/issues/1410))
- Azure Stack doc page ([#1371](https://github.com/Azure/aks-engine/issues/1371))
- update custom vnet documentation ([#1399](https://github.com/Azure/aks-engine/issues/1399))
- update issue templates ([#1382](https://github.com/Azure/aks-engine/issues/1382))
- update prow docs ([#1364](https://github.com/Azure/aks-engine/issues/1364))
- update quickstart.md ([#1341](https://github.com/Azure/aks-engine/issues/1341))
- update tutorial docs to use kubernetes instead of swarm ([#1337](https://github.com/Azure/aks-engine/issues/1337))
- add customVMTags to clusterdefinitions documentation ([#1332](https://github.com/Azure/aks-engine/issues/1332))

### Features 🌈
- add support for Kubernetes 1.15.0-rc.1  ([#1469](https://github.com/Azure/aks-engine/issues/1469))
- enable PSP w/ privileged ClusterRoleBinding for 1.15 ([#1454](https://github.com/Azure/aks-engine/issues/1454))
- add support for Kubernetes 1.15.0-beta.2 ([#1438](https://github.com/Azure/aks-engine/issues/1438))
- add support for Kubernetes 1.14.3 ([#1439](https://github.com/Azure/aks-engine/issues/1439))
- add support for Kubernetes 1.13.7 ([#1441](https://github.com/Azure/aks-engine/issues/1441))
- use all VMAS fault domains in a region ([#1090](https://github.com/Azure/aks-engine/issues/1090))
- add support for Kubernetes 1.12.8 and 1.13.5 on Azure Stack ([#1419](https://github.com/Azure/aks-engine/issues/1419))
- add support for Kubernetes 1.15.0-beta.1 ([#1394](https://github.com/Azure/aks-engine/issues/1394))
- add support for Kubernetes 1.12.9 ([#1383](https://github.com/Azure/aks-engine/issues/1383))
- added support for direct array assignment using --set ([#709](https://github.com/Azure/aks-engine/issues/709))
- allow "client_secret" auth method with ADFS identity provider ([#1343](https://github.com/Azure/aks-engine/issues/1343))

### Maintenance 🔧
- update VHD to 2019.06.12 ([#1481](https://github.com/Azure/aks-engine/issues/1481))
- add note about apimodel in issue template ([#1478](https://github.com/Azure/aks-engine/issues/1478))
- delete obsolete calico spec ([#1463](https://github.com/Azure/aks-engine/issues/1463))
- reconcile generated files ([#1457](https://github.com/Azure/aks-engine/issues/1457))
- remove unused go template funcs ([#1331](https://github.com/Azure/aks-engine/issues/1331))
- enable only strong TLS cipher suites for kubelet by default ([#1434](https://github.com/Azure/aks-engine/issues/1434))
- update etcd default to 3.2.26 ([#1451](https://github.com/Azure/aks-engine/issues/1451))
- enable 2019.06.08 VHD ([#1450](https://github.com/Azure/aks-engine/issues/1450))
- move azure stack example VM SKUs from Standard_D2_v2 to Standard_D2_v3 ([#1448](https://github.com/Azure/aks-engine/issues/1448))
- move examples VM SKUs from Standard_D2_v2 to Standard_D2_v3 ([#1436](https://github.com/Azure/aks-engine/issues/1436))
- use CoreDNS 1.5.0 for k8s >= 1.12.0 ([#1411](https://github.com/Azure/aks-engine/issues/1411))
- re-enable k8s v1.12.9 ([#1428](https://github.com/Azure/aks-engine/issues/1428))
- update go toolchain to 1.12.5 ([#1305](https://github.com/Azure/aks-engine/issues/1305))
- ignore junit.xml from e2e ([#1417](https://github.com/Azure/aks-engine/issues/1417))
- add security to changelog commit titles ([#1415](https://github.com/Azure/aks-engine/issues/1415))
- update generated files ([#1405](https://github.com/Azure/aks-engine/issues/1405))
- add support for etcd 3.2.26 ([#1384](https://github.com/Azure/aks-engine/issues/1384))
- ensure all Linux-bound files are LF (not CRLF) ([#1354](https://github.com/Azure/aks-engine/issues/1354))
- bump cluster-autoscaler to latest patch versions ([#1339](https://github.com/Azure/aks-engine/issues/1339))
- add example cluster definition for Azure Stack ([#1304](https://github.com/Azure/aks-engine/issues/1304))
- **CIS:** Remove log file permission enforcement  ([#1412](https://github.com/Azure/aks-engine/issues/1412))
- **CIS:** generate kubelet server cert in aks-e ([#1416](https://github.com/Azure/aks-engine/issues/1416))

### Revert Change ◀️
- remove k8s 1.12.9 support ([#1476](https://github.com/Azure/aks-engine/issues/1476))
- add support for Kubernetes 1.12.9 ([#1420](https://github.com/Azure/aks-engine/issues/1420))

### Security Fix 🛡️
- remove support for 1.13.6 and 1.14.2 ([#1413](https://github.com/Azure/aks-engine/issues/1413))

### Testing 💚
- HPA scale down config only works if k8s >= 1.12 ([#1480](https://github.com/Azure/aks-engine/issues/1480))
- add subnets to route table in E2E for custom vnet + kubenet ([#1470](https://github.com/Azure/aks-engine/issues/1470))
- parallelize tests part 1 ([#1455](https://github.com/Azure/aks-engine/issues/1455))
- enable DNS liveness test for calico-enabled clusters ([#1459](https://github.com/Azure/aks-engine/issues/1459))
- Switch Windows test passes to VMSS by default to match AKS ([#1452](https://github.com/Azure/aks-engine/issues/1452))
- getTemplateFuncMap unit tests part 5 ([#1449](https://github.com/Azure/aks-engine/issues/1449))
- getTemplateFuncMap unit tests part 4 ([#1388](https://github.com/Azure/aks-engine/issues/1388))
- quicker “defaults addons image” unit test ([#1426](https://github.com/Azure/aks-engine/issues/1426))
- update unit test payload for Azure Stack ([#1408](https://github.com/Azure/aks-engine/issues/1408))
- enable aks-engine e2e test on azurestack ([#1397](https://github.com/Azure/aks-engine/issues/1397))
- refactor addons default image test ([#1407](https://github.com/Azure/aks-engine/issues/1407))
- getTemplateFuncMap unit tests part 3 ([#1344](https://github.com/Azure/aks-engine/issues/1344))
- getTemplateFuncMap unit tests part 2 ([#1338](https://github.com/Azure/aks-engine/issues/1338))
- sorted GetKubernetesLabels method for deterministic unit tests ([#1340](https://github.com/Azure/aks-engine/issues/1340))
- getTemplateFuncMap unit tests part 1 ([#1334](https://github.com/Azure/aks-engine/issues/1334))
- add prototype for getTemplateFuncMap unit tests ([#1333](https://github.com/Azure/aks-engine/issues/1333))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.36.5"></a>
# [v0.36.5] - 2019-06-12
### Bug Fixes 🐞
- fix race condition between cloud init walinuxagent hold and CSE ([#1471](https://github.com/Azure/aks-engine/issues/1471))
- fix missing variable in cse ([#1460](https://github.com/Azure/aks-engine/issues/1460))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.36.4"></a>
# [v0.36.4] - 2019-05-30
### Security Fix 🛡️
- remove support for 1.13.6 and 1.14.2 ([#1413](https://github.com/Azure/aks-engine/issues/1413))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.36.3"></a>
# [v0.36.3] - 2019-05-29
### Bug Fixes 🐞
- disable redundant cosmosdb regions ([#1403](https://github.com/Azure/aks-engine/issues/1403))
- remove "allow-privileged" in kubelet 1.15.0 ([#1369](https://github.com/Azure/aks-engine/issues/1369))
- Update Kata download url ([#1390](https://github.com/Azure/aks-engine/issues/1390))
- decrease default host MTU for Azure Stack ([#1346](https://github.com/Azure/aks-engine/issues/1346))
- revert previous load balancer changes for Azure Stack ([#1347](https://github.com/Azure/aks-engine/issues/1347))
- never use 3.0.4 for moby-cli and update moby on upgrade ([#1359](https://github.com/Azure/aks-engine/issues/1359))
- Base64 encode sp password for windows vmss ([#1327](https://github.com/Azure/aks-engine/issues/1327))
- get vmss node version from K8S API instead of tags for instances on old model ([#1299](https://github.com/Azure/aks-engine/issues/1299))
- enable Windows plus custom VNET ([#1314](https://github.com/Azure/aks-engine/issues/1314))
- USER_ASSIGNED_IDENTITY_ID is empty in azure.json ([#1283](https://github.com/Azure/aks-engine/issues/1283))
- don’t add etcd data disk in cosmos etcd scenario ([#1310](https://github.com/Azure/aks-engine/issues/1310))
- remove grub file permissions enforcement ([#1308](https://github.com/Azure/aks-engine/issues/1308))
- increase upgrade timeout to 180 minutes ([#1300](https://github.com/Azure/aks-engine/issues/1300))
- commit generated go-bindata files with "--no-compress" option ([#1088](https://github.com/Azure/aks-engine/issues/1088))
- Ensure pods scheduled onto new nodes during upgrade respect the original node's labels/taints ([#1044](https://github.com/Azure/aks-engine/issues/1044))
- add validation for non-support of prometheus extension for Windows ([#1259](https://github.com/Azure/aks-engine/issues/1259))
- update short hyperkube commands in manifests for k8s components ([#1279](https://github.com/Azure/aks-engine/issues/1279))
- removing PodSecurityPolicy files from manifests folders ([#1257](https://github.com/Azure/aks-engine/issues/1257))
- aks-engine deploy tutorial errors out for auth method as CLI ([#1263](https://github.com/Azure/aks-engine/issues/1263))
- remove outbound connectivity validation on disconnected Azure Stack stamps ([#1250](https://github.com/Azure/aks-engine/issues/1250))
- default kubelet flags windows/linux reconciliation w/ unit tests ([#1244](https://github.com/Azure/aks-engine/issues/1244))
- Windows Kubelet issues error due to unsupported config ([#1240](https://github.com/Azure/aks-engine/issues/1240))
- private cluster with VMSS masters, jumpbox ([#1226](https://github.com/Azure/aks-engine/issues/1226))
- enableEncryptionWithExternalKms with master VMSS missing objectID ([#1194](https://github.com/Azure/aks-engine/issues/1194))
- add Azure China support for Calico DaemonSet ([#1089](https://github.com/Azure/aks-engine/issues/1089))

### Build 🏭
- validate deps before 'make build', and some cleanup ([#1271](https://github.com/Azure/aks-engine/issues/1271))

### Code Style 🎶
- check for unused parameters ([#1323](https://github.com/Azure/aks-engine/issues/1323))
- use 'cmd foo' instead of 'cmd $(echo foo)' ([#1321](https://github.com/Azure/aks-engine/issues/1321))
- don't use `tr` to replace words ([#1316](https://github.com/Azure/aks-engine/issues/1316))
- check exit code directly, not indirectly with $? ([#1256](https://github.com/Azure/aks-engine/issues/1256))
- standardize makefile syntax ([#1266](https://github.com/Azure/aks-engine/issues/1266))
- declare and assign separately to avoid masking return values ([#1241](https://github.com/Azure/aks-engine/issues/1241))
- use single quotes in trapping exit codes ([#1242](https://github.com/Azure/aks-engine/issues/1242))
- bash parameter expansion replaces "echo | sed" ([#1234](https://github.com/Azure/aks-engine/issues/1234))
- remove unneeded 'echo' commands ([#1230](https://github.com/Azure/aks-engine/issues/1230))
- sudo doesn't affect redirects ([#1213](https://github.com/Azure/aks-engine/issues/1213))
- replace unneeded 'cat' commands ([#1202](https://github.com/Azure/aks-engine/issues/1202))
- expressions don't expand in single quotes ([#1201](https://github.com/Azure/aks-engine/issues/1201))
- remove quotes that accidentally unquote ([#1200](https://github.com/Azure/aks-engine/issues/1200))
- don't use shell variables in printf format string ([#1197](https://github.com/Azure/aks-engine/issues/1197))
- use 'grep -c' instead of 'grep|wc' ([#1191](https://github.com/Azure/aks-engine/issues/1191))
- fix suspicious unquoted literal strings ([#1187](https://github.com/Azure/aks-engine/issues/1187))
- use builtin 'command -v' instead of nonstandard 'which' ([#1185](https://github.com/Azure/aks-engine/issues/1185))
- use 'grep -E' instead of deprecated 'egrep' ([#1186](https://github.com/Azure/aks-engine/issues/1186))
- use $(...) shell notation instead of legacy backticks ([#1180](https://github.com/Azure/aks-engine/issues/1180))

### Continuous Integration 💜
- permit optional scope field when generating changelog ([#1196](https://github.com/Azure/aks-engine/issues/1196))

### Documentation 📘
- update group name and corresponding output folder names ([#1280](https://github.com/Azure/aks-engine/issues/1280))
- Clarify terminology around patches and use of windows-patches extension ([#1309](https://github.com/Azure/aks-engine/issues/1309))
- fix typos in launch.json debug example ([#1294](https://github.com/Azure/aks-engine/issues/1294))
- add a more complete VS Code debug configuration ([#1275](https://github.com/Azure/aks-engine/issues/1275))
- Update k8s version references in docs/samples ([#1264](https://github.com/Azure/aks-engine/issues/1264))
- correct "--auth-method" arguments in help ([#1272](https://github.com/Azure/aks-engine/issues/1272))
- improve developer debugging instructions ([#1211](https://github.com/Azure/aks-engine/issues/1211))

### Features 🌈
- add support for Kubernetes 1.15.0-beta.1 ([#1394](https://github.com/Azure/aks-engine/issues/1394))
- allow "client_secret" auth method with ADFS identity provider ([#1343](https://github.com/Azure/aks-engine/issues/1343))
- add support for Kubernetes 1.14.2 ([#1315](https://github.com/Azure/aks-engine/issues/1315))
- custom tags on VMs and scale sets ([#1277](https://github.com/Azure/aks-engine/issues/1277))
- Make cordon drain timeout configurable with --upgrade ([#1276](https://github.com/Azure/aks-engine/issues/1276))
- add support for Kubernetes 1.13.6 ([#1262](https://github.com/Azure/aks-engine/issues/1262))
- add support for Kubernetes 1.15.0-alpha.3 ([#1247](https://github.com/Azure/aks-engine/issues/1247))
- disable unsupported addons on Azure Stack deployments ([#1233](https://github.com/Azure/aks-engine/issues/1233))
- remove azurefile storage class for Azure Stack ([#1222](https://github.com/Azure/aks-engine/issues/1222))
- add auditd as an ubuntu option ([#1143](https://github.com/Azure/aks-engine/issues/1143))
- add support for Kubernetes 1.11.10 ([#1193](https://github.com/Azure/aks-engine/issues/1193))
- add two US DoD Azure locations ([#1205](https://github.com/Azure/aks-engine/issues/1205))

### Maintenance 🔧
- add example cluster definition for Azure Stack ([#1304](https://github.com/Azure/aks-engine/issues/1304))
- update VHD image to 2019.05.16 ([#1319](https://github.com/Azure/aks-engine/issues/1319))
- Update calico to v3.7.2 ([#1293](https://github.com/Azure/aks-engine/issues/1293))
- update templates_generated.go ([#1317](https://github.com/Azure/aks-engine/issues/1317))
- update templates_generated.go ([#1295](https://github.com/Azure/aks-engine/issues/1295))
- enable VHD 2019.05.08, disable auditd for non-VHD ([#1286](https://github.com/Azure/aks-engine/issues/1286))
- update client-go and k8s.io vendored code ([#1273](https://github.com/Azure/aks-engine/issues/1273))
- rename addons-related types and consts ([#1265](https://github.com/Azure/aks-engine/issues/1265))
- deprecate CIS files cloud-init paving for non-VHD distros ([#1251](https://github.com/Azure/aks-engine/issues/1251))
- remove support for Kubernetes 1.11.8 ([#1243](https://github.com/Azure/aks-engine/issues/1243))
- enable moby 3.0.5, and set 3.0.5 to default ([#1236](https://github.com/Azure/aks-engine/issues/1236))
- update Azure CNI version to v1.0.22 ([#1192](https://github.com/Azure/aks-engine/issues/1192))
- update VHD version to 04.30.2019 ([#1184](https://github.com/Azure/aks-engine/issues/1184))
- deprecate unused CSE func installDockerEngine ([#1199](https://github.com/Azure/aks-engine/issues/1199))
- update omsagent addon to use the latest version ([#1156](https://github.com/Azure/aks-engine/issues/1156))
- onboard Azure CNI images to MCR ([#1153](https://github.com/Azure/aks-engine/issues/1153))

### Revert Change ◀️
- /etc/default/grub changes aren't reconcilable over time ([#1237](https://github.com/Azure/aks-engine/issues/1237))

### Testing 💚
- unit tests for kubernetesManifestSettingsInit ([#1312](https://github.com/Azure/aks-engine/issues/1312))
- add masterSSHPort to auditd E2E SCP ([#1307](https://github.com/Azure/aks-engine/issues/1307))
- fix auditd E2E for master VMSS ([#1297](https://github.com/Azure/aks-engine/issues/1297))
- add storage + Azure Stack addons unit tests ([#1285](https://github.com/Azure/aks-engine/issues/1285))
- reduced name for ilb curl deployment ([#1274](https://github.com/Azure/aks-engine/issues/1274))
- add no outbound CSE unit test ([#1254](https://github.com/Azure/aks-engine/issues/1254))
- test for missing kubelet configs as well ([#1252](https://github.com/Azure/aks-engine/issues/1252))
- idempotent e2e changes for soak test scenarios ([#1214](https://github.com/Azure/aks-engine/issues/1214))
- enable all ssh to master tests for vmss masters ([#1198](https://github.com/Azure/aks-engine/issues/1198))
- hasSSHAbleMaster returns false if master vmss ([#1195](https://github.com/Azure/aks-engine/issues/1195))
- can’t ssh into master vmss master vms ([#1188](https://github.com/Azure/aks-engine/issues/1188))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.35.4"></a>
# [v0.35.4] - 2019-05-29
### Bug Fixes 🐞
- ensure /etc/default/grub has no CRLF
- Base64 encode sp password for windows vmss ([#1327](https://github.com/Azure/aks-engine/issues/1327))
- enableEncryptionWithExternalKms with master VMSS missing objectID ([#1194](https://github.com/Azure/aks-engine/issues/1194))
- add Azure China support for Calico DaemonSet ([#1089](https://github.com/Azure/aks-engine/issues/1089))

### Features 🌈
- add two US DoD Azure locations ([#1205](https://github.com/Azure/aks-engine/issues/1205))

### Maintenance 🔧
- rev the SKU as well
- remove temporary CRLF fix
- enable 2019.05.23 VHD versions for v0.35
- fix all CRLF system files
- move to fix_crlf func
- move dos2unix install earlier in flow
- update VHD version to 04.30.2019 ([#1184](https://github.com/Azure/aks-engine/issues/1184))
- update omsagent addon to use the latest version ([#1156](https://github.com/Azure/aks-engine/issues/1156))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.36.2"></a>
# [v0.36.2] - 2019-05-24
### Bug Fixes 🐞
- decrease default host MTU for Azure Stack ([#1346](https://github.com/Azure/aks-engine/issues/1346))
- revert previous load balancer changes for Azure Stack ([#1347](https://github.com/Azure/aks-engine/issues/1347))

### Features 🌈
- allow "client_secret" auth method with ADFS identity provider ([#1343](https://github.com/Azure/aks-engine/issues/1343))

### Maintenance 🔧
- add example cluster definition for Azure Stack ([#1304](https://github.com/Azure/aks-engine/issues/1304))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.36.1"></a>
# [v0.36.1] - 2019-05-23
### Bug Fixes 🐞
- never use 3.0.4 for moby-cli and update moby on upgrade ([#1359](https://github.com/Azure/aks-engine/issues/1359))
- Base64 encode sp password for windows vmss ([#1327](https://github.com/Azure/aks-engine/issues/1327))
- get vmss node version from K8S API instead of tags for instances on old model ([#1299](https://github.com/Azure/aks-engine/issues/1299))
- enable Windows plus custom VNET ([#1314](https://github.com/Azure/aks-engine/issues/1314))
- USER_ASSIGNED_IDENTITY_ID is empty in azure.json ([#1283](https://github.com/Azure/aks-engine/issues/1283))
- don’t add etcd data disk in cosmos etcd scenario ([#1310](https://github.com/Azure/aks-engine/issues/1310))
- remove grub file permissions enforcement ([#1308](https://github.com/Azure/aks-engine/issues/1308))
- increase upgrade timeout to 180 minutes ([#1300](https://github.com/Azure/aks-engine/issues/1300))
- commit generated go-bindata files with "--no-compress" option ([#1088](https://github.com/Azure/aks-engine/issues/1088))
- Ensure pods scheduled onto new nodes during upgrade respect the original node's labels/taints ([#1044](https://github.com/Azure/aks-engine/issues/1044))
- add validation for non-support of prometheus extension for Windows ([#1259](https://github.com/Azure/aks-engine/issues/1259))
- update short hyperkube commands in manifests for k8s components ([#1279](https://github.com/Azure/aks-engine/issues/1279))
- removing PodSecurityPolicy files from manifests folders ([#1257](https://github.com/Azure/aks-engine/issues/1257))
- aks-engine deploy tutorial errors out for auth method as CLI ([#1263](https://github.com/Azure/aks-engine/issues/1263))
- remove outbound connectivity validation on disconnected Azure Stack stamps ([#1250](https://github.com/Azure/aks-engine/issues/1250))
- default kubelet flags windows/linux reconciliation w/ unit tests ([#1244](https://github.com/Azure/aks-engine/issues/1244))
- Windows Kubelet issues error due to unsupported config ([#1240](https://github.com/Azure/aks-engine/issues/1240))
- private cluster with VMSS masters, jumpbox ([#1226](https://github.com/Azure/aks-engine/issues/1226))
- enableEncryptionWithExternalKms with master VMSS missing objectID ([#1194](https://github.com/Azure/aks-engine/issues/1194))
- add Azure China support for Calico DaemonSet ([#1089](https://github.com/Azure/aks-engine/issues/1089))

### Build 🏭
- validate deps before 'make build', and some cleanup ([#1271](https://github.com/Azure/aks-engine/issues/1271))

### Code Style 🎶
- check for unused parameters ([#1323](https://github.com/Azure/aks-engine/issues/1323))
- use 'cmd foo' instead of 'cmd $(echo foo)' ([#1321](https://github.com/Azure/aks-engine/issues/1321))
- don't use `tr` to replace words ([#1316](https://github.com/Azure/aks-engine/issues/1316))
- check exit code directly, not indirectly with $? ([#1256](https://github.com/Azure/aks-engine/issues/1256))
- standardize makefile syntax ([#1266](https://github.com/Azure/aks-engine/issues/1266))
- declare and assign separately to avoid masking return values ([#1241](https://github.com/Azure/aks-engine/issues/1241))
- use single quotes in trapping exit codes ([#1242](https://github.com/Azure/aks-engine/issues/1242))
- bash parameter expansion replaces "echo | sed" ([#1234](https://github.com/Azure/aks-engine/issues/1234))
- remove unneeded 'echo' commands ([#1230](https://github.com/Azure/aks-engine/issues/1230))
- sudo doesn't affect redirects ([#1213](https://github.com/Azure/aks-engine/issues/1213))
- replace unneeded 'cat' commands ([#1202](https://github.com/Azure/aks-engine/issues/1202))
- expressions don't expand in single quotes ([#1201](https://github.com/Azure/aks-engine/issues/1201))
- remove quotes that accidentally unquote ([#1200](https://github.com/Azure/aks-engine/issues/1200))
- don't use shell variables in printf format string ([#1197](https://github.com/Azure/aks-engine/issues/1197))
- use 'grep -c' instead of 'grep|wc' ([#1191](https://github.com/Azure/aks-engine/issues/1191))
- fix suspicious unquoted literal strings ([#1187](https://github.com/Azure/aks-engine/issues/1187))
- use builtin 'command -v' instead of nonstandard 'which' ([#1185](https://github.com/Azure/aks-engine/issues/1185))
- use 'grep -E' instead of deprecated 'egrep' ([#1186](https://github.com/Azure/aks-engine/issues/1186))
- use $(...) shell notation instead of legacy backticks ([#1180](https://github.com/Azure/aks-engine/issues/1180))

### Continuous Integration 💜
- permit optional scope field when generating changelog ([#1196](https://github.com/Azure/aks-engine/issues/1196))

### Documentation 📘
- update group name and corresponding output folder names ([#1280](https://github.com/Azure/aks-engine/issues/1280))
- Clarify terminology around patches and use of windows-patches extension ([#1309](https://github.com/Azure/aks-engine/issues/1309))
- fix typos in launch.json debug example ([#1294](https://github.com/Azure/aks-engine/issues/1294))
- add a more complete VS Code debug configuration ([#1275](https://github.com/Azure/aks-engine/issues/1275))
- Update k8s version references in docs/samples ([#1264](https://github.com/Azure/aks-engine/issues/1264))
- correct "--auth-method" arguments in help ([#1272](https://github.com/Azure/aks-engine/issues/1272))
- improve developer debugging instructions ([#1211](https://github.com/Azure/aks-engine/issues/1211))

### Features 🌈
- add support for Kubernetes 1.14.2 ([#1315](https://github.com/Azure/aks-engine/issues/1315))
- custom tags on VMs and scale sets ([#1277](https://github.com/Azure/aks-engine/issues/1277))
- Make cordon drain timeout configurable with --upgrade ([#1276](https://github.com/Azure/aks-engine/issues/1276))
- add support for Kubernetes 1.13.6 ([#1262](https://github.com/Azure/aks-engine/issues/1262))
- add support for Kubernetes 1.15.0-alpha.3 ([#1247](https://github.com/Azure/aks-engine/issues/1247))
- disable unsupported addons on Azure Stack deployments ([#1233](https://github.com/Azure/aks-engine/issues/1233))
- remove azurefile storage class for Azure Stack ([#1222](https://github.com/Azure/aks-engine/issues/1222))
- add auditd as an ubuntu option ([#1143](https://github.com/Azure/aks-engine/issues/1143))
- add support for Kubernetes 1.11.10 ([#1193](https://github.com/Azure/aks-engine/issues/1193))
- add two US DoD Azure locations ([#1205](https://github.com/Azure/aks-engine/issues/1205))

### Maintenance 🔧
- update VHD image to 2019.05.16 ([#1319](https://github.com/Azure/aks-engine/issues/1319))
- Update calico to v3.7.2 ([#1293](https://github.com/Azure/aks-engine/issues/1293))
- update templates_generated.go ([#1317](https://github.com/Azure/aks-engine/issues/1317))
- update templates_generated.go ([#1295](https://github.com/Azure/aks-engine/issues/1295))
- enable VHD 2019.05.08, disable auditd for non-VHD ([#1286](https://github.com/Azure/aks-engine/issues/1286))
- update client-go and k8s.io vendored code ([#1273](https://github.com/Azure/aks-engine/issues/1273))
- rename addons-related types and consts ([#1265](https://github.com/Azure/aks-engine/issues/1265))
- deprecate CIS files cloud-init paving for non-VHD distros ([#1251](https://github.com/Azure/aks-engine/issues/1251))
- remove support for Kubernetes 1.11.8 ([#1243](https://github.com/Azure/aks-engine/issues/1243))
- enable moby 3.0.5, and set 3.0.5 to default ([#1236](https://github.com/Azure/aks-engine/issues/1236))
- update Azure CNI version to v1.0.22 ([#1192](https://github.com/Azure/aks-engine/issues/1192))
- update VHD version to 04.30.2019 ([#1184](https://github.com/Azure/aks-engine/issues/1184))
- deprecate unused CSE func installDockerEngine ([#1199](https://github.com/Azure/aks-engine/issues/1199))
- update omsagent addon to use the latest version ([#1156](https://github.com/Azure/aks-engine/issues/1156))
- onboard Azure CNI images to MCR ([#1153](https://github.com/Azure/aks-engine/issues/1153))

### Revert Change ◀️
- /etc/default/grub changes aren't reconcilable over time ([#1237](https://github.com/Azure/aks-engine/issues/1237))

### Testing 💚
- unit tests for kubernetesManifestSettingsInit ([#1312](https://github.com/Azure/aks-engine/issues/1312))
- add masterSSHPort to auditd E2E SCP ([#1307](https://github.com/Azure/aks-engine/issues/1307))
- fix auditd E2E for master VMSS ([#1297](https://github.com/Azure/aks-engine/issues/1297))
- add storage + Azure Stack addons unit tests ([#1285](https://github.com/Azure/aks-engine/issues/1285))
- reduced name for ilb curl deployment ([#1274](https://github.com/Azure/aks-engine/issues/1274))
- add no outbound CSE unit test ([#1254](https://github.com/Azure/aks-engine/issues/1254))
- test for missing kubelet configs as well ([#1252](https://github.com/Azure/aks-engine/issues/1252))
- idempotent e2e changes for soak test scenarios ([#1214](https://github.com/Azure/aks-engine/issues/1214))
- enable all ssh to master tests for vmss masters ([#1198](https://github.com/Azure/aks-engine/issues/1198))
- hasSSHAbleMaster returns false if master vmss ([#1195](https://github.com/Azure/aks-engine/issues/1195))
- can’t ssh into master vmss master vms ([#1188](https://github.com/Azure/aks-engine/issues/1188))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.35.3"></a>
# [v0.35.3] - 2019-05-23
### Bug Fixes 🐞
- ensure /etc/default/grub has no CRLF
- Base64 encode sp password for windows vmss ([#1327](https://github.com/Azure/aks-engine/issues/1327))
- enableEncryptionWithExternalKms with master VMSS missing objectID ([#1194](https://github.com/Azure/aks-engine/issues/1194))
- add Azure China support for Calico DaemonSet ([#1089](https://github.com/Azure/aks-engine/issues/1089))

### Features 🌈
- add two US DoD Azure locations ([#1205](https://github.com/Azure/aks-engine/issues/1205))

### Maintenance 🔧
- fix all CRLF system files
- move to fix_crlf func
- move dos2unix install earlier in flow
- update VHD version to 04.30.2019 ([#1184](https://github.com/Azure/aks-engine/issues/1184))
- update omsagent addon to use the latest version ([#1156](https://github.com/Azure/aks-engine/issues/1156))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.36.0"></a>
# [v0.36.0] - 2019-05-20
### Bug Fixes 🐞
- Base64 encode sp password for windows vmss ([#1327](https://github.com/Azure/aks-engine/issues/1327))
- get vmss node version from K8S API instead of tags for instances on old model ([#1299](https://github.com/Azure/aks-engine/issues/1299))
- enable Windows plus custom VNET ([#1314](https://github.com/Azure/aks-engine/issues/1314))
- USER_ASSIGNED_IDENTITY_ID is empty in azure.json ([#1283](https://github.com/Azure/aks-engine/issues/1283))
- don’t add etcd data disk in cosmos etcd scenario ([#1310](https://github.com/Azure/aks-engine/issues/1310))
- remove grub file permissions enforcement ([#1308](https://github.com/Azure/aks-engine/issues/1308))
- increase upgrade timeout to 180 minutes ([#1300](https://github.com/Azure/aks-engine/issues/1300))
- commit generated go-bindata files with "--no-compress" option ([#1088](https://github.com/Azure/aks-engine/issues/1088))
- Ensure pods scheduled onto new nodes during upgrade respect the original node's labels/taints ([#1044](https://github.com/Azure/aks-engine/issues/1044))
- add validation for non-support of prometheus extension for Windows ([#1259](https://github.com/Azure/aks-engine/issues/1259))
- update short hyperkube commands in manifests for k8s components ([#1279](https://github.com/Azure/aks-engine/issues/1279))
- removing PodSecurityPolicy files from manifests folders ([#1257](https://github.com/Azure/aks-engine/issues/1257))
- aks-engine deploy tutorial errors out for auth method as CLI ([#1263](https://github.com/Azure/aks-engine/issues/1263))
- remove outbound connectivity validation on disconnected Azure Stack stamps ([#1250](https://github.com/Azure/aks-engine/issues/1250))
- default kubelet flags windows/linux reconciliation w/ unit tests ([#1244](https://github.com/Azure/aks-engine/issues/1244))
- Windows Kubelet issues error due to unsupported config ([#1240](https://github.com/Azure/aks-engine/issues/1240))
- private cluster with VMSS masters, jumpbox ([#1226](https://github.com/Azure/aks-engine/issues/1226))
- enableEncryptionWithExternalKms with master VMSS missing objectID ([#1194](https://github.com/Azure/aks-engine/issues/1194))
- add Azure China support for Calico DaemonSet ([#1089](https://github.com/Azure/aks-engine/issues/1089))

### Build 🏭
- validate deps before 'make build', and some cleanup ([#1271](https://github.com/Azure/aks-engine/issues/1271))

### Code Style 🎶
- check for unused parameters ([#1323](https://github.com/Azure/aks-engine/issues/1323))
- use 'cmd foo' instead of 'cmd $(echo foo)' ([#1321](https://github.com/Azure/aks-engine/issues/1321))
- don't use `tr` to replace words ([#1316](https://github.com/Azure/aks-engine/issues/1316))
- check exit code directly, not indirectly with $? ([#1256](https://github.com/Azure/aks-engine/issues/1256))
- standardize makefile syntax ([#1266](https://github.com/Azure/aks-engine/issues/1266))
- declare and assign separately to avoid masking return values ([#1241](https://github.com/Azure/aks-engine/issues/1241))
- use single quotes in trapping exit codes ([#1242](https://github.com/Azure/aks-engine/issues/1242))
- bash parameter expansion replaces "echo | sed" ([#1234](https://github.com/Azure/aks-engine/issues/1234))
- remove unneeded 'echo' commands ([#1230](https://github.com/Azure/aks-engine/issues/1230))
- sudo doesn't affect redirects ([#1213](https://github.com/Azure/aks-engine/issues/1213))
- replace unneeded 'cat' commands ([#1202](https://github.com/Azure/aks-engine/issues/1202))
- expressions don't expand in single quotes ([#1201](https://github.com/Azure/aks-engine/issues/1201))
- remove quotes that accidentally unquote ([#1200](https://github.com/Azure/aks-engine/issues/1200))
- don't use shell variables in printf format string ([#1197](https://github.com/Azure/aks-engine/issues/1197))
- use 'grep -c' instead of 'grep|wc' ([#1191](https://github.com/Azure/aks-engine/issues/1191))
- fix suspicious unquoted literal strings ([#1187](https://github.com/Azure/aks-engine/issues/1187))
- use builtin 'command -v' instead of nonstandard 'which' ([#1185](https://github.com/Azure/aks-engine/issues/1185))
- use 'grep -E' instead of deprecated 'egrep' ([#1186](https://github.com/Azure/aks-engine/issues/1186))
- use $(...) shell notation instead of legacy backticks ([#1180](https://github.com/Azure/aks-engine/issues/1180))

### Continuous Integration 💜
- permit optional scope field when generating changelog ([#1196](https://github.com/Azure/aks-engine/issues/1196))

### Documentation 📘
- update group name and corresponding output folder names ([#1280](https://github.com/Azure/aks-engine/issues/1280))
- Clarify terminology around patches and use of windows-patches extension ([#1309](https://github.com/Azure/aks-engine/issues/1309))
- fix typos in launch.json debug example ([#1294](https://github.com/Azure/aks-engine/issues/1294))
- add a more complete VS Code debug configuration ([#1275](https://github.com/Azure/aks-engine/issues/1275))
- Update k8s version references in docs/samples ([#1264](https://github.com/Azure/aks-engine/issues/1264))
- correct "--auth-method" arguments in help ([#1272](https://github.com/Azure/aks-engine/issues/1272))
- improve developer debugging instructions ([#1211](https://github.com/Azure/aks-engine/issues/1211))

### Features 🌈
- add support for Kubernetes 1.14.2 ([#1315](https://github.com/Azure/aks-engine/issues/1315))
- custom tags on VMs and scale sets ([#1277](https://github.com/Azure/aks-engine/issues/1277))
- Make cordon drain timeout configurable with --upgrade ([#1276](https://github.com/Azure/aks-engine/issues/1276))
- add support for Kubernetes 1.13.6 ([#1262](https://github.com/Azure/aks-engine/issues/1262))
- add support for Kubernetes 1.15.0-alpha.3 ([#1247](https://github.com/Azure/aks-engine/issues/1247))
- disable unsupported addons on Azure Stack deployments ([#1233](https://github.com/Azure/aks-engine/issues/1233))
- remove azurefile storage class for Azure Stack ([#1222](https://github.com/Azure/aks-engine/issues/1222))
- add auditd as an ubuntu option ([#1143](https://github.com/Azure/aks-engine/issues/1143))
- add support for Kubernetes 1.11.10 ([#1193](https://github.com/Azure/aks-engine/issues/1193))
- add two US DoD Azure locations ([#1205](https://github.com/Azure/aks-engine/issues/1205))

### Maintenance 🔧
- update VHD image to 2019.05.16 ([#1319](https://github.com/Azure/aks-engine/issues/1319))
- Update calico to v3.7.2 ([#1293](https://github.com/Azure/aks-engine/issues/1293))
- update templates_generated.go ([#1317](https://github.com/Azure/aks-engine/issues/1317))
- update templates_generated.go ([#1295](https://github.com/Azure/aks-engine/issues/1295))
- enable VHD 2019.05.08, disable auditd for non-VHD ([#1286](https://github.com/Azure/aks-engine/issues/1286))
- update client-go and k8s.io vendored code ([#1273](https://github.com/Azure/aks-engine/issues/1273))
- rename addons-related types and consts ([#1265](https://github.com/Azure/aks-engine/issues/1265))
- deprecate CIS files cloud-init paving for non-VHD distros ([#1251](https://github.com/Azure/aks-engine/issues/1251))
- remove support for Kubernetes 1.11.8 ([#1243](https://github.com/Azure/aks-engine/issues/1243))
- enable moby 3.0.5, and set 3.0.5 to default ([#1236](https://github.com/Azure/aks-engine/issues/1236))
- update Azure CNI version to v1.0.22 ([#1192](https://github.com/Azure/aks-engine/issues/1192))
- update VHD version to 04.30.2019 ([#1184](https://github.com/Azure/aks-engine/issues/1184))
- deprecate unused CSE func installDockerEngine ([#1199](https://github.com/Azure/aks-engine/issues/1199))
- update omsagent addon to use the latest version ([#1156](https://github.com/Azure/aks-engine/issues/1156))
- onboard Azure CNI images to MCR ([#1153](https://github.com/Azure/aks-engine/issues/1153))

### Revert Change ◀️
- /etc/default/grub changes aren't reconcilable over time ([#1237](https://github.com/Azure/aks-engine/issues/1237))

### Testing 💚
- unit tests for kubernetesManifestSettingsInit ([#1312](https://github.com/Azure/aks-engine/issues/1312))
- add masterSSHPort to auditd E2E SCP ([#1307](https://github.com/Azure/aks-engine/issues/1307))
- fix auditd E2E for master VMSS ([#1297](https://github.com/Azure/aks-engine/issues/1297))
- add storage + Azure Stack addons unit tests ([#1285](https://github.com/Azure/aks-engine/issues/1285))
- reduced name for ilb curl deployment ([#1274](https://github.com/Azure/aks-engine/issues/1274))
- add no outbound CSE unit test ([#1254](https://github.com/Azure/aks-engine/issues/1254))
- test for missing kubelet configs as well ([#1252](https://github.com/Azure/aks-engine/issues/1252))
- idempotent e2e changes for soak test scenarios ([#1214](https://github.com/Azure/aks-engine/issues/1214))
- enable all ssh to master tests for vmss masters ([#1198](https://github.com/Azure/aks-engine/issues/1198))
- hasSSHAbleMaster returns false if master vmss ([#1195](https://github.com/Azure/aks-engine/issues/1195))
- can’t ssh into master vmss master vms ([#1188](https://github.com/Azure/aks-engine/issues/1188))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.35.2"></a>
# [v0.35.2] - 2019-05-20
### Bug Fixes 🐞
- Base64 encode sp password for windows vmss ([#1327](https://github.com/Azure/aks-engine/issues/1327))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.35.1"></a>
# [v0.35.1] - 2019-05-02
### Bug Fixes 🐞
- enableEncryptionWithExternalKms with master VMSS missing objectID ([#1194](https://github.com/Azure/aks-engine/issues/1194))
- add Azure China support for Calico DaemonSet ([#1089](https://github.com/Azure/aks-engine/issues/1089))

### Features 🌈
- add two US DoD Azure locations ([#1205](https://github.com/Azure/aks-engine/issues/1205))

### Maintenance 🔧
- update VHD version to 04.30.2019 ([#1184](https://github.com/Azure/aks-engine/issues/1184))
- update omsagent addon to use the latest version ([#1156](https://github.com/Azure/aks-engine/issues/1156))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.35.0"></a>
# [v0.35.0] - 2019-04-30
### Bug Fixes 🐞
- kubeconfig enforcement via kubelet.sh ([#1183](https://github.com/Azure/aks-engine/issues/1183))
- delete vm if node is not ready after upgrade ([#1173](https://github.com/Azure/aks-engine/issues/1173))
- [#926](https://github.com/Azure/aks-engine/issues/926) - Base64 encode sp password ([#1174](https://github.com/Azure/aks-engine/issues/1174))
- --rotate-certificates failing in k8s versions prior to 1.11 ([#1175](https://github.com/Azure/aks-engine/issues/1175))
- use cordon/drain timeout for pod eviction timeout during upgrade ([#1157](https://github.com/Azure/aks-engine/issues/1157))
- add UDP Standard LB rule to enable outbound access ([#1155](https://github.com/Azure/aks-engine/issues/1155))
- clean generated and unit test helper files ([#1151](https://github.com/Azure/aks-engine/issues/1151))
- implement extensions with template refactor ([#1133](https://github.com/Azure/aks-engine/issues/1133))
- replace \t with nothing ([#1136](https://github.com/Azure/aks-engine/issues/1136))
- set correct master FQDN for Azure Stack ([#1132](https://github.com/Azure/aks-engine/issues/1132))
- enable VHD enforcement of /etc/pam.d/su ([#1130](https://github.com/Azure/aks-engine/issues/1130))
- /etc/default/grub VHD enforcement ([#1131](https://github.com/Azure/aks-engine/issues/1131))
- restores cosmosDB etcd functionality after template generation refactor ([#1120](https://github.com/Azure/aks-engine/issues/1120))
- remove deprecated sshd configs for 18.04 ([#1118](https://github.com/Azure/aks-engine/issues/1118))
- use the actual pool index in the VM prefix for windows scale operations ([#1098](https://github.com/Azure/aks-engine/issues/1098))
- make conntrack more liberal on packets ([#1073](https://github.com/Azure/aks-engine/issues/1073))
- no longer need 1.12 kubelet start workaround ([#1069](https://github.com/Azure/aks-engine/issues/1069))
- error handling in CIS bash script ([#1062](https://github.com/Azure/aks-engine/issues/1062))
- Windows + VMSS external svc routing broken ([#1028](https://github.com/Azure/aks-engine/issues/1028))
- fix parameter name for ScaleSetEvictionPolicy ([#1059](https://github.com/Azure/aks-engine/issues/1059))
- handle multiple masters for Azure Stack ([#1053](https://github.com/Azure/aks-engine/issues/1053))
- only run installDeps in non-VHD scenarios ([#1043](https://github.com/Azure/aks-engine/issues/1043))
- fix generate command for Azure Stack ([#1038](https://github.com/Azure/aks-engine/issues/1038))
- GetKubernetesB64ConfigsCustomCloud rebase errata ([#1006](https://github.com/Azure/aks-engine/issues/1006))

### Code Refactoring 💎
- remove engine pkg custom script unnecessary functions ([#1036](https://github.com/Azure/aks-engine/issues/1036))
- re-organize cloud-init files and implementation ([#996](https://github.com/Azure/aks-engine/issues/996))
- return error instead of invoking log.Fatal ([#997](https://github.com/Azure/aks-engine/issues/997))
- move isPrivateCluster to types.go ([#998](https://github.com/Azure/aks-engine/issues/998))

### Code Style 🎶
- use { cmd1; cmd2 } >> file instead of individual redirects ([#1167](https://github.com/Azure/aks-engine/issues/1167))
- use #!, not just #, for the shebang ([#1170](https://github.com/Azure/aks-engine/issues/1170))
- the shebang must be on the first line ([#1171](https://github.com/Azure/aks-engine/issues/1171))
- use "-n" instead of "! -z" in shell scripts ([#1166](https://github.com/Azure/aks-engine/issues/1166))
- add bash shebangs to scripts ([#1159](https://github.com/Azure/aks-engine/issues/1159))
- remove literal carriage returns and enable parser errors ([#1163](https://github.com/Azure/aks-engine/issues/1163))
- use 'true' instead of '[ 1 ]' in shell loops ([#1154](https://github.com/Azure/aks-engine/issues/1154))
- remove literal carriage returns and lint all shell scripts ([#1148](https://github.com/Azure/aks-engine/issues/1148))
- comment out unused constant ([#1147](https://github.com/Azure/aks-engine/issues/1147))
- remove unneeded nil check ([#1110](https://github.com/Azure/aks-engine/issues/1110))

### Continuous Integration 💜
- exclude a test mock file from coverage report ([#1176](https://github.com/Azure/aks-engine/issues/1176))
- return error if Go linter can't compile source code ([#1150](https://github.com/Azure/aks-engine/issues/1150))
- exclude some test files from coverage report ([#1160](https://github.com/Azure/aks-engine/issues/1160))
- Add alert on VHD space greater than 75% ([#1097](https://github.com/Azure/aks-engine/issues/1097))

### Documentation 📘
- Add guides for calico cleanup after upgrading to v3.5 ([#1137](https://github.com/Azure/aks-engine/issues/1137))
- fix IPs per vnet limit documented ([#1124](https://github.com/Azure/aks-engine/issues/1124))
- add makedev.ps1 to dev guide ([#1122](https://github.com/Azure/aks-engine/issues/1122))
- clarify help for --api-model arguments ([#1071](https://github.com/Azure/aks-engine/issues/1071))
- add documentation for ACC agent pools running Ubuntu 18.04 ([#1003](https://github.com/Azure/aks-engine/issues/1003))
- fix a command ([#1000](https://github.com/Azure/aks-engine/issues/1000))

### Features 🌈
- add support for Kubernetes v1.15.0-alpha.2 ([#1178](https://github.com/Azure/aks-engine/issues/1178))
- revive CoreOS support ([#892](https://github.com/Azure/aks-engine/issues/892))
- add support for Kubernetes v1.15.0-alpha.1  ([#1140](https://github.com/Azure/aks-engine/issues/1140))
- pre-pull Pause from Azure Stack's docker repo ([#1144](https://github.com/Azure/aks-engine/issues/1144))
- static pods use custom hyperkube on az stack ([#1142](https://github.com/Azure/aks-engine/issues/1142))
- Support existing load balancer backend address pool for agent nodes ([#1145](https://github.com/Azure/aks-engine/issues/1145))
- add support for Kubernetes 1.12.8 ([#1138](https://github.com/Azure/aks-engine/issues/1138))
- grant user assigned identity 'Reader' role for hosted masters ([#1076](https://github.com/Azure/aks-engine/issues/1076))
- enable WindowsProfile in defaults enforcement code flow ([#1103](https://github.com/Azure/aks-engine/issues/1103))
- pre-pull Azure Stack's custom Hyperkube ([#1040](https://github.com/Azure/aks-engine/issues/1040))
- support VMSS agent nodes with public IP. ([#1087](https://github.com/Azure/aks-engine/issues/1087))
- TLS certificate rotation ([#678](https://github.com/Azure/aks-engine/issues/678))
- enable calico 3.5 for AKS ([#995](https://github.com/Azure/aks-engine/issues/995))
- **aks:** expose unversioned orchestrator version profile to better enable aks preview kubernetes flags ([#1135](https://github.com/Azure/aks-engine/issues/1135))

### Maintenance 🔧
- add two missing /var/log files for 18.04 in CIS script ([#1113](https://github.com/Azure/aks-engine/issues/1113))
- enable 2019.04.24 VHD images ([#1165](https://github.com/Azure/aks-engine/issues/1165))
- enable 2019.04.08 VHD versions ([#989](https://github.com/Azure/aks-engine/issues/989))
- add --feature-gates tests ([#1005](https://github.com/Azure/aks-engine/issues/1005))
- run apt-get dist-upgrade during VHD or full install scenarios ([#1007](https://github.com/Azure/aks-engine/issues/1007))
- update Azure VM sizes ([#1101](https://github.com/Azure/aks-engine/issues/1101))
- stop delivering unused kubelet systemd timer script via cloud-init ([#1141](https://github.com/Azure/aks-engine/issues/1141))
- backport tests + cloud-init usage for CIS changes ([#1057](https://github.com/Azure/aks-engine/issues/1057))
- remove comments from cilium addons spec ([#1042](https://github.com/Azure/aks-engine/issues/1042))
- reduce cloud-init var overhead for VHD scenarios ([#1127](https://github.com/Azure/aks-engine/issues/1127))
- clean up containerd and cc-runtime when unused ([#1129](https://github.com/Azure/aks-engine/issues/1129))
- disable --pod-max-pids by default ([#1126](https://github.com/Azure/aks-engine/issues/1126))
- always ensure apt runs non-interactively ([#1102](https://github.com/Azure/aks-engine/issues/1102))
- add newlines to login banner message ([#1114](https://github.com/Azure/aks-engine/issues/1114))
- add 2019.04.24 VHD release notes ([#1182](https://github.com/Azure/aks-engine/issues/1182))
- add CIS script interface ([#972](https://github.com/Azure/aks-engine/issues/972))
- **CIS:** set umask to 027 for ubuntu ([#1128](https://github.com/Azure/aks-engine/issues/1128))
- **CIS:** kernel module hardening for non-essential filesystem types ([#1105](https://github.com/Azure/aks-engine/issues/1105))
- **CIS:** ensure su is restricted ([#1112](https://github.com/Azure/aks-engine/issues/1112))
- **CIS:** Ensure password creation requirements are configured ([#1035](https://github.com/Azure/aks-engine/issues/1035))
- **CIS:** Ensure logging is configured ([#1081](https://github.com/Azure/aks-engine/issues/1081))
- **CIS:** ensure /etc/ssh/sshd_config is configured ([#1030](https://github.com/Azure/aks-engine/issues/1030))
- **CIS:** add apt package validation, ensure postfix is not present ([#1063](https://github.com/Azure/aks-engine/issues/1063))
- **CIS:** enforce CIS modprobe recommendations ([#1061](https://github.com/Azure/aks-engine/issues/1061))
- **CIS:** Ensure remote login warning banner is configured properly ([#1037](https://github.com/Azure/aks-engine/issues/1037))
- **CIS:** harden grub.cfg file permissions ([#1106](https://github.com/Azure/aks-engine/issues/1106))
- **CIS:** pam.d password enforcement ([#1116](https://github.com/Azure/aks-engine/issues/1116))
- **CIS:** CIS network configuration enforcement ([#1039](https://github.com/Azure/aks-engine/issues/1039))
- **CIS:** Ensure permissions on all log files are configured ([#1031](https://github.com/Azure/aks-engine/issues/1031))
- **CIS:** assign root pw ([#1013](https://github.com/Azure/aks-engine/issues/1013))
- **CIS:** add protect-kernel-defaults ([#999](https://github.com/Azure/aks-engine/issues/999))
- **CIS:** ensure local login warning banner is configured properly ([#1024](https://github.com/Azure/aks-engine/issues/1024))
- **CIS:** add streaming-connection-idle-timeout ([#977](https://github.com/Azure/aks-engine/issues/977))
- **CIS:** password expiration, cron file mode enforcement ([#1162](https://github.com/Azure/aks-engine/issues/1162))
- **CIS:** add rotate kubelet certs flag ([#1052](https://github.com/Azure/aks-engine/issues/1052))
- **CIS:** grub configuration changes to accommodate CIS ([#1111](https://github.com/Azure/aks-engine/issues/1111))

### Revert Change ◀️
- JoinControllers system.conf configuration ([#1095](https://github.com/Azure/aks-engine/issues/1095))
- "fix: make conntrack more liberal on packets"

### Testing 💚
- use master branch as root URL in extensions unit test ([#1161](https://github.com/Azure/aks-engine/issues/1161))
- fallback google.com check for Windows outbound test ([#1117](https://github.com/Azure/aks-engine/issues/1117))
- only test Ready nodes for DNS ([#1121](https://github.com/Azure/aks-engine/issues/1121))
- check net.ipv4.tcp_retries2 kernel parameter ([#1094](https://github.com/Azure/aks-engine/issues/1094))
- disable coreos scenario as PR E2E gate ([#1107](https://github.com/Azure/aks-engine/issues/1107))
- single host OS DNS test script ([#1083](https://github.com/Azure/aks-engine/issues/1083))
- add ubuntu time sync E2E validation ([#1080](https://github.com/Azure/aks-engine/issues/1080))
- add nodes.GetReady() method for E2E tests ([#1082](https://github.com/Azure/aks-engine/issues/1082))
- ensure node is ready before running validations ([#1045](https://github.com/Azure/aks-engine/issues/1045))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.34.2"></a>
# [v0.34.2] - 2019-04-23
#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.34.3"></a>
# [v0.34.3] - 2019-04-23
### Bug Fixes 🐞
- set correct master FQDN for Azure Stack ([#1132](https://github.com/Azure/aks-engine/issues/1132))
- restores cosmosDB etcd functionality after template generation refactor ([#1120](https://github.com/Azure/aks-engine/issues/1120))
- use the actual pool index in the VM prefix for windows scale operations ([#1098](https://github.com/Azure/aks-engine/issues/1098))
- Windows + VMSS external svc routing broken ([#1028](https://github.com/Azure/aks-engine/issues/1028))
- fix parameter name for ScaleSetEvictionPolicy ([#1059](https://github.com/Azure/aks-engine/issues/1059))
- handle multiple masters for Azure Stack ([#1053](https://github.com/Azure/aks-engine/issues/1053))
- fix generate command for Azure Stack ([#1038](https://github.com/Azure/aks-engine/issues/1038))
- handle uppercase letters in VM names in upgrade ([#951](https://github.com/Azure/aks-engine/issues/951))
- mount tmp emptyDir to stop CoreDNS restarts ([#949](https://github.com/Azure/aks-engine/issues/949))
- fix disable rbac + enable aggregated API upgrade bug ([#946](https://github.com/Azure/aks-engine/issues/946))
- zones values for VMAS masters is an array of a single string ([#943](https://github.com/Azure/aks-engine/issues/943))
- increase cordon/drain timeout to 20 mins during upgrade ([#938](https://github.com/Azure/aks-engine/issues/938))
- [#926](https://github.com/Azure/aks-engine/issues/926) Escape servicePrincipalClientSecret in Windows custom script block ([#927](https://github.com/Azure/aks-engine/issues/927))
- include vnet CIDR in windows node CNI outbound NAT exception list ([#901](https://github.com/Azure/aks-engine/issues/901))
- AKS VMSS windows agent pool upgrade ([#900](https://github.com/Azure/aks-engine/issues/900))
- ip-masq-agent container image has a different repo base in VHD ([#911](https://github.com/Azure/aks-engine/issues/911))
- typo in get-akse.sh script name ([#914](https://github.com/Azure/aks-engine/issues/914))
- remove packer user ([#909](https://github.com/Azure/aks-engine/issues/909))
- use --dns-prefix for output directory when running deploy command ([#880](https://github.com/Azure/aks-engine/issues/880))
- user contributor role for Master VMAS role assignments ([#886](https://github.com/Azure/aks-engine/issues/886))
- address regressions from v2 template refactor ([#882](https://github.com/Azure/aks-engine/issues/882))
- use our own hpa-example image ([#885](https://github.com/Azure/aks-engine/issues/885))

### Code Refactoring 💎
- combine all cloud-init vars into one ([#960](https://github.com/Azure/aks-engine/issues/960))
- avoid shell globbing in azure-const.sh ([#955](https://github.com/Azure/aks-engine/issues/955))
- put each version on its own line in VHD script ([#893](https://github.com/Azure/aks-engine/issues/893))

### Code Style 🎶
- use shellcheck to validate some scripts  ([#957](https://github.com/Azure/aks-engine/issues/957))
- remove extra pullContainerImage command ([#954](https://github.com/Azure/aks-engine/issues/954))

### Continuous Integration 💜
- add timestamp to artifact names ([#941](https://github.com/Azure/aks-engine/issues/941))
- remove CircleCI end-to-end test support ([#898](https://github.com/Azure/aks-engine/issues/898))

### Documentation 📘
- use "make info" as preflight release check ([#969](https://github.com/Azure/aks-engine/issues/969))
- re-add windows binaries doc link ([#942](https://github.com/Azure/aks-engine/issues/942))
- fix reference to large cluster api model ([#936](https://github.com/Azure/aks-engine/issues/936))
- generate all shasum strings with a one-liner ([#925](https://github.com/Azure/aks-engine/issues/925))
- rename local download script name to match remote ([#917](https://github.com/Azure/aks-engine/issues/917))
- cleaned up examples ([#915](https://github.com/Azure/aks-engine/issues/915))

### Features 🌈
- pre-pull Azure Stack's custom Hyperkube ([#1040](https://github.com/Azure/aks-engine/issues/1040))
- support VMSS agent nodes with public IP. ([#1087](https://github.com/Azure/aks-engine/issues/1087))
- enable calico 3.5 for AKS ([#995](https://github.com/Azure/aks-engine/issues/995))
- retrieve endpoints from metadata endpoint for Azure Stack ([#947](https://github.com/Azure/aks-engine/issues/947))
- add support for Kubernetes 1.14.1 ([#958](https://github.com/Azure/aks-engine/issues/958))
- enable nvidia drivers on moby-backed VMs, deprecate docker-engine entirely ([#897](https://github.com/Azure/aks-engine/issues/897))
- default to k8s 1.14 for Windows and 1.12 for Linux ([#918](https://github.com/Azure/aks-engine/issues/918))
- remove dependency on deployment-dir for scale/upgrade ([#890](https://github.com/Azure/aks-engine/issues/890))
- support deploy command on Azure Stack ([#755](https://github.com/Azure/aks-engine/issues/755))

### Maintenance 🔧
- disable --pod-max-pids by default ([#1126](https://github.com/Azure/aks-engine/issues/1126))
- remove systemConf cruft
- remove comments from cilium addons spec ([#1042](https://github.com/Azure/aks-engine/issues/1042))
- enable 2019.04.08 VHD versions ([#989](https://github.com/Azure/aks-engine/issues/989))
- nvidia driver version update to 418.40.04 ([#920](https://github.com/Azure/aks-engine/issues/920))
- update go toolchain to 1.12.2 ([#959](https://github.com/Azure/aks-engine/issues/959))
- only use new calico CSE config for AKS Engine ([#952](https://github.com/Azure/aks-engine/issues/952))
- enable 2019.04.03 VHD versions ([#948](https://github.com/Azure/aks-engine/issues/948))
- update Azure CNI version to v1.0.18 ([#928](https://github.com/Azure/aks-engine/issues/928))
- Update calico to 3.5 and allow calico to work with azure CNI ([#454](https://github.com/Azure/aks-engine/issues/454))
- ensure AKS gets proper Windows billing extension ([#924](https://github.com/Azure/aks-engine/issues/924))
- use V1beta Policy client, not the deprecated default ([#916](https://github.com/Azure/aks-engine/issues/916))
- update VHD image to 2019.03.27 ([#903](https://github.com/Azure/aks-engine/issues/903))
- remove tariq1890 from the OWNERS list ([#888](https://github.com/Azure/aks-engine/issues/888))
- remove Kubernetes 1.10 from e2e testing ([#895](https://github.com/Azure/aks-engine/issues/895))
- update CNI plugin to v0.7.5 ([#894](https://github.com/Azure/aks-engine/issues/894))
- enable 2019.03.25 VHD image ([#884](https://github.com/Azure/aks-engine/issues/884))

### Performance Improvements 🚀
- Network Monitor v0.0.6 integration ([#912](https://github.com/Azure/aks-engine/issues/912))
- improve CSE script ([#870](https://github.com/Azure/aks-engine/issues/870))

### Revert Change ◀️
- JoinControllers system.conf configuration ([#1095](https://github.com/Azure/aks-engine/issues/1095))

### Testing 💚
- add ubuntu 18.04 job to pr-e2e ([#966](https://github.com/Azure/aks-engine/issues/966))
- unit test for azure stack AzureClient ([#934](https://github.com/Azure/aks-engine/issues/934))
- use managed-standard storage class for pvc test ([#945](https://github.com/Azure/aks-engine/issues/945))
- remove hard-coded k8s version from armvars tests ([#939](https://github.com/Azure/aks-engine/issues/939))
- skip dns-liveness restarts check for k8s 1.14 ([#931](https://github.com/Azure/aks-engine/issues/931))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.33.6"></a>
# [v0.33.6] - 2019-04-18
### Bug Fixes 🐞
- [#926](https://github.com/Azure/aks-engine/issues/926) Escape servicePrincipalClientSecret in Windows custom script block ([#927](https://github.com/Azure/aks-engine/issues/927))
- include vnet CIDR in windows node CNI outbound NAT exception list ([#901](https://github.com/Azure/aks-engine/issues/901))
- AKS VMSS windows agent pool upgrade ([#900](https://github.com/Azure/aks-engine/issues/900))
- user contributor role for Master VMAS role assignments ([#886](https://github.com/Azure/aks-engine/issues/886))
- address regressions from v2 template refactor ([#882](https://github.com/Azure/aks-engine/issues/882))
- use our own hpa-example image ([#885](https://github.com/Azure/aks-engine/issues/885))

### Features 🌈
- enable calico 3.5 for AKS ([#995](https://github.com/Azure/aks-engine/issues/995))

### Maintenance 🔧
- only use new calico CSE config for AKS Engine ([#952](https://github.com/Azure/aks-engine/issues/952))
- Update calico to 3.5 and allow calico to work with azure CNI ([#454](https://github.com/Azure/aks-engine/issues/454))
- ensure AKS gets proper Windows billing extension ([#924](https://github.com/Azure/aks-engine/issues/924))
- update VHD image to 2019.03.27 ([#903](https://github.com/Azure/aks-engine/issues/903))
- update CNI plugin to v0.7.5 ([#894](https://github.com/Azure/aks-engine/issues/894))
- enable 2019.03.25 VHD image ([#884](https://github.com/Azure/aks-engine/issues/884))

### Revert Change ◀️
- JoinControllers system.conf configuration ([#1095](https://github.com/Azure/aks-engine/issues/1095))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.34.1"></a>
# [v0.34.1] - 2019-04-15
### Bug Fixes 🐞
- handle multiple masters for Azure Stack ([#1053](https://github.com/Azure/aks-engine/issues/1053))
- fix generate command for Azure Stack ([#1038](https://github.com/Azure/aks-engine/issues/1038))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.34.0"></a>
# [v0.34.0] - 2019-04-09
### Bug Fixes 🐞
- handle uppercase letters in VM names in upgrade ([#951](https://github.com/Azure/aks-engine/issues/951))
- mount tmp emptyDir to stop CoreDNS restarts ([#949](https://github.com/Azure/aks-engine/issues/949))
- fix disable rbac + enable aggregated API upgrade bug ([#946](https://github.com/Azure/aks-engine/issues/946))
- zones values for VMAS masters is an array of a single string ([#943](https://github.com/Azure/aks-engine/issues/943))
- increase cordon/drain timeout to 20 mins during upgrade ([#938](https://github.com/Azure/aks-engine/issues/938))
- [#926](https://github.com/Azure/aks-engine/issues/926) Escape servicePrincipalClientSecret in Windows custom script block ([#927](https://github.com/Azure/aks-engine/issues/927))
- include vnet CIDR in windows node CNI outbound NAT exception list ([#901](https://github.com/Azure/aks-engine/issues/901))
- AKS VMSS windows agent pool upgrade ([#900](https://github.com/Azure/aks-engine/issues/900))
- ip-masq-agent container image has a different repo base in VHD ([#911](https://github.com/Azure/aks-engine/issues/911))
- typo in get-akse.sh script name ([#914](https://github.com/Azure/aks-engine/issues/914))
- remove packer user ([#909](https://github.com/Azure/aks-engine/issues/909))
- use --dns-prefix for output directory when running deploy command ([#880](https://github.com/Azure/aks-engine/issues/880))
- user contributor role for Master VMAS role assignments ([#886](https://github.com/Azure/aks-engine/issues/886))
- address regressions from v2 template refactor ([#882](https://github.com/Azure/aks-engine/issues/882))
- use our own hpa-example image ([#885](https://github.com/Azure/aks-engine/issues/885))

### Code Refactoring 💎
- combine all cloud-init vars into one ([#960](https://github.com/Azure/aks-engine/issues/960))
- avoid shell globbing in azure-const.sh ([#955](https://github.com/Azure/aks-engine/issues/955))
- put each version on its own line in VHD script ([#893](https://github.com/Azure/aks-engine/issues/893))

### Code Style 🎶
- use shellcheck to validate some scripts  ([#957](https://github.com/Azure/aks-engine/issues/957))
- remove extra pullContainerImage command ([#954](https://github.com/Azure/aks-engine/issues/954))

### Continuous Integration 💜
- add timestamp to artifact names ([#941](https://github.com/Azure/aks-engine/issues/941))
- remove CircleCI end-to-end test support ([#898](https://github.com/Azure/aks-engine/issues/898))

### Documentation 📘
- use "make info" as preflight release check ([#969](https://github.com/Azure/aks-engine/issues/969))
- re-add windows binaries doc link ([#942](https://github.com/Azure/aks-engine/issues/942))
- fix reference to large cluster api model ([#936](https://github.com/Azure/aks-engine/issues/936))
- generate all shasum strings with a one-liner ([#925](https://github.com/Azure/aks-engine/issues/925))
- rename local download script name to match remote ([#917](https://github.com/Azure/aks-engine/issues/917))
- cleaned up examples ([#915](https://github.com/Azure/aks-engine/issues/915))

### Features 🌈
- enable calico 3.5 for AKS ([#995](https://github.com/Azure/aks-engine/issues/995))
- retrieve endpoints from metadata endpoint for Azure Stack ([#947](https://github.com/Azure/aks-engine/issues/947))
- add support for Kubernetes 1.14.1 ([#958](https://github.com/Azure/aks-engine/issues/958))
- enable nvidia drivers on moby-backed VMs, deprecate docker-engine entirely ([#897](https://github.com/Azure/aks-engine/issues/897))
- default to k8s 1.14 for Windows and 1.12 for Linux ([#918](https://github.com/Azure/aks-engine/issues/918))
- remove dependency on deployment-dir for scale/upgrade ([#890](https://github.com/Azure/aks-engine/issues/890))
- support deploy command on Azure Stack ([#755](https://github.com/Azure/aks-engine/issues/755))

### Maintenance 🔧
- enable 2019.04.08 VHD versions ([#989](https://github.com/Azure/aks-engine/issues/989))
- nvidia driver version update to 418.40.04 ([#920](https://github.com/Azure/aks-engine/issues/920))
- update go toolchain to 1.12.2 ([#959](https://github.com/Azure/aks-engine/issues/959))
- only use new calico CSE config for AKS Engine ([#952](https://github.com/Azure/aks-engine/issues/952))
- enable 2019.04.03 VHD versions ([#948](https://github.com/Azure/aks-engine/issues/948))
- update Azure CNI version to v1.0.18 ([#928](https://github.com/Azure/aks-engine/issues/928))
- Update calico to 3.5 and allow calico to work with azure CNI ([#454](https://github.com/Azure/aks-engine/issues/454))
- ensure AKS gets proper Windows billing extension ([#924](https://github.com/Azure/aks-engine/issues/924))
- use V1beta Policy client, not the deprecated default ([#916](https://github.com/Azure/aks-engine/issues/916))
- update VHD image to 2019.03.27 ([#903](https://github.com/Azure/aks-engine/issues/903))
- remove tariq1890 from the OWNERS list ([#888](https://github.com/Azure/aks-engine/issues/888))
- remove Kubernetes 1.10 from e2e testing ([#895](https://github.com/Azure/aks-engine/issues/895))
- update CNI plugin to v0.7.5 ([#894](https://github.com/Azure/aks-engine/issues/894))
- enable 2019.03.25 VHD image ([#884](https://github.com/Azure/aks-engine/issues/884))

### Performance Improvements 🚀
- Network Monitor v0.0.6 integration ([#912](https://github.com/Azure/aks-engine/issues/912))
- improve CSE script ([#870](https://github.com/Azure/aks-engine/issues/870))

### Testing 💚
- add ubuntu 18.04 job to pr-e2e ([#966](https://github.com/Azure/aks-engine/issues/966))
- unit test for azure stack AzureClient ([#934](https://github.com/Azure/aks-engine/issues/934))
- use managed-standard storage class for pvc test ([#945](https://github.com/Azure/aks-engine/issues/945))
- remove hard-coded k8s version from armvars tests ([#939](https://github.com/Azure/aks-engine/issues/939))
- skip dns-liveness restarts check for k8s 1.14 ([#931](https://github.com/Azure/aks-engine/issues/931))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.33.5"></a>
# [v0.33.5] - 2019-04-05
### Maintenance 🔧
- only use new calico CSE config for AKS Engine ([#952](https://github.com/Azure/aks-engine/issues/952))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.33.4"></a>
# [v0.33.4] - 2019-04-01
### Bug Fixes 🐞
- [#926](https://github.com/Azure/aks-engine/issues/926) Escape servicePrincipalClientSecret in Windows custom script block ([#927](https://github.com/Azure/aks-engine/issues/927))
- include vnet CIDR in windows node CNI outbound NAT exception list ([#901](https://github.com/Azure/aks-engine/issues/901))
- AKS VMSS windows agent pool upgrade ([#900](https://github.com/Azure/aks-engine/issues/900))

### Maintenance 🔧
- Update calico to 3.5 and allow calico to work with azure CNI ([#454](https://github.com/Azure/aks-engine/issues/454))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.33.3"></a>
# [v0.33.3] - 2019-04-01
### Maintenance 🔧
- ensure AKS gets proper Windows billing extension ([#924](https://github.com/Azure/aks-engine/issues/924))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.33.2"></a>
# [v0.33.2] - 2019-03-28
### Maintenance 🔧
- update VHD image to 2019.03.27 ([#903](https://github.com/Azure/aks-engine/issues/903))
- update CNI plugin to v0.7.5 ([#894](https://github.com/Azure/aks-engine/issues/894))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.33.1"></a>
# [v0.33.1] - 2019-03-26
### Bug Fixes 🐞
- user contributor role for Master VMAS role assignments ([#886](https://github.com/Azure/aks-engine/issues/886))
- address regressions from v2 template refactor ([#882](https://github.com/Azure/aks-engine/issues/882))
- use our own hpa-example image ([#885](https://github.com/Azure/aks-engine/issues/885))
- restore back-compat for clusters that have an older --audit-policy-file ([#830](https://github.com/Azure/aks-engine/issues/830))
- add missing right paren in ARM expression ([#871](https://github.com/Azure/aks-engine/issues/871))
- fix sshd logging to auth.log for OMS ([#862](https://github.com/Azure/aks-engine/issues/862))
- change blobfuse-flexvol-installer imagePullPolicy to IfNotPresent ([#860](https://github.com/Azure/aks-engine/issues/860))
- increase timeout tolerance for nvidia drivers installation ([#857](https://github.com/Azure/aks-engine/issues/857))

### Code Style 🎶
- use constants consistently in pkg/api ([#865](https://github.com/Azure/aks-engine/issues/865))
- remove unused constants ([#863](https://github.com/Azure/aks-engine/issues/863))
- fix some linting issues in api/validate ([#859](https://github.com/Azure/aks-engine/issues/859))

### Documentation 📘
- how to create a new aks-engine release ([#829](https://github.com/Azure/aks-engine/issues/829))

### Features 🌈
- add support for Kubernetes 1.14.0 ([#876](https://github.com/Azure/aks-engine/issues/876))
- add support for Kubernetes 1.11.9 ([#873](https://github.com/Azure/aks-engine/issues/873))
- add support for Kubernetes 1.13.5 ([#874](https://github.com/Azure/aks-engine/issues/874))
- add support for Kubernetes 1.12.7 ([#872](https://github.com/Azure/aks-engine/issues/872))
- AKS-Engine Install Script for Linux  ([#832](https://github.com/Azure/aks-engine/issues/832))

### Maintenance 🔧
- enable 2019.03.25 VHD image ([#884](https://github.com/Azure/aks-engine/issues/884))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.33.0"></a>
# [v0.33.0] - 2019-03-21
### Bug Fixes 🐞
- perform more nil guard checks against LinuxProfile ([#828](https://github.com/Azure/aks-engine/issues/828))
- add nil checks before accessing the service principal profile ([#824](https://github.com/Azure/aks-engine/issues/824))
- add extra nil guard checks for LinuxProfiles ([#822](https://github.com/Azure/aks-engine/issues/822))
- remove shell output from VHD release notes ([#812](https://github.com/Azure/aks-engine/issues/812))
- add more kubernetesConfig nil checks to account for more resilience. ([#816](https://github.com/Azure/aks-engine/issues/816))
- add nil checks for more resilience in template refactor ([#814](https://github.com/Azure/aks-engine/issues/814))
- rehabilitate constant generator script  ([#785](https://github.com/Azure/aks-engine/issues/785))
- add missing parameter ComputerNamePrefix for VMSS agents using Windows ([#800](https://github.com/Azure/aks-engine/issues/800))
- Bug fixes for the HostedMasterProfile Scenario ([#796](https://github.com/Azure/aks-engine/issues/796))
- move /var/log/ during VHD build to a vhd specific directory ([#777](https://github.com/Azure/aks-engine/issues/777))
- fix bugs in jumpbox custom data generations and add more unit tests to engine package ([#782](https://github.com/Azure/aks-engine/issues/782))
- mark walinux for hold in cloud-init ([#778](https://github.com/Azure/aks-engine/issues/778))
- actually apt-mark hold walinuxagent during all CSE runs ([#771](https://github.com/Azure/aks-engine/issues/771))
- add changes to allow for scaling of VMSS nodes in no_template flows ([#769](https://github.com/Azure/aks-engine/issues/769))
- removed unnecessary elb artifacts when PrivateCluster is enabled ([#747](https://github.com/Azure/aks-engine/issues/747))
- fix the typo in the scale cmd logs ([#757](https://github.com/Azure/aks-engine/issues/757))
- correct typo is err message in scale when unmarshaling ARM params ([#745](https://github.com/Azure/aks-engine/issues/745))
- add err!=nil check when loading an ARM template during scale ops ([#744](https://github.com/Azure/aks-engine/issues/744))
- bug fix for master VMSS scenario in no_template flows ([#742](https://github.com/Azure/aks-engine/issues/742))
- always configure the container runtime for GPU in an N series context ([#675](https://github.com/Azure/aks-engine/issues/675))
- commands return errors rather than exit ([#693](https://github.com/Azure/aks-engine/issues/693))
- add the system role assignment object as the resource instead of the function pointer ([#735](https://github.com/Azure/aks-engine/issues/735))
- Allow masters with custom images to use data disks for etcd ([#710](https://github.com/Azure/aks-engine/issues/710))
- add err check while JSON marshalling the ARM template ([#734](https://github.com/Azure/aks-engine/issues/734))
- use the correct go sdk import for system assigned role assignments ([#732](https://github.com/Azure/aks-engine/issues/732))
- add err != nil checks for json Unmarshal in scale and upgrade ops ([#705](https://github.com/Azure/aks-engine/issues/705))
- primaryScaleSetName and primaryAvailabilitySetName var logic is incorrect ([#708](https://github.com/Azure/aks-engine/issues/708))
- implement cosmosEtcd no_template ([#703](https://github.com/Azure/aks-engine/issues/703))
- remove obsolete allocateNodeCidrs ARM variable ([#699](https://github.com/Azure/aks-engine/issues/699))
- agentpools have their own “is custom vnet” logic ([#701](https://github.com/Azure/aks-engine/issues/701))
- missing variable declarations for masters-only scenario ([#692](https://github.com/Azure/aks-engine/issues/692))
- update apiVersionCompute to 2018-10-01 in armvariables ([#690](https://github.com/Azure/aks-engine/issues/690))
- 2 Azure Stack vars are obligatory ([#688](https://github.com/Azure/aks-engine/issues/688))
- master vmss errors ([#684](https://github.com/Azure/aks-engine/issues/684))

### Build 🏭
- compress aks-engine binaries ([#700](https://github.com/Azure/aks-engine/issues/700))
- add "build-container" target ([#698](https://github.com/Azure/aks-engine/issues/698))

### Code Refactoring 💎
- aks-engine no_template implementation ([#324](https://github.com/Azure/aks-engine/issues/324))

### Code Style 🎶
- use errors.Errorf instead of fmt.Errorf as it's more idiomatic ([#831](https://github.com/Azure/aks-engine/issues/831))
- fix go_vet issues in the code ([#825](https://github.com/Azure/aks-engine/issues/825))
- refactorings suggested by the gocritic linter ([#748](https://github.com/Azure/aks-engine/issues/748))
- add godoc to the armtype structs used in the template refactor ([#772](https://github.com/Azure/aks-engine/issues/772))
- add some golint style simplications to validate.go ([#759](https://github.com/Azure/aks-engine/issues/759))
- enable unnecessary type conversion linter ([#740](https://github.com/Azure/aks-engine/issues/740))
- enable vet shadow and remove the shadowed declaration go anti-pattern ([#715](https://github.com/Azure/aks-engine/issues/715))

### Continuous Integration 💜
- run dep check before linting in the pr-e2e pipeline ([#733](https://github.com/Azure/aks-engine/issues/733))

### Documentation 📘
- Fix language on upgrade.md ([#810](https://github.com/Azure/aks-engine/issues/810))
- updating instructions on how to use Azure Cosmos DB etcd API ([#669](https://github.com/Azure/aks-engine/issues/669))
- add how to install aks-engine with Chocolatey ([#660](https://github.com/Azure/aks-engine/issues/660))
- update docs for NVIDIA driver version ([#752](https://github.com/Azure/aks-engine/issues/752))

### Features 🌈
- enable new template generation implementation ([#682](https://github.com/Azure/aks-engine/issues/682))
- update Azure-NPM to support new network policy specifications ([#638](https://github.com/Azure/aks-engine/issues/638))
- add support for Ubuntu 18.04-LTS ([#223](https://github.com/Azure/aks-engine/issues/223))
- only allow “strong” TLS cipher suites by default ([#681](https://github.com/Azure/aks-engine/issues/681))
- windows agentpool support for AKS ([#714](https://github.com/Azure/aks-engine/issues/714))
- add selective agentpool upgrade support - VMSS ([#672](https://github.com/Azure/aks-engine/issues/672))

### Maintenance 🔧
- enable Kubernetes v1.14.0-rc.1 ([#793](https://github.com/Azure/aks-engine/issues/793))
- update AKS distro version to 2019.03.15 ([#794](https://github.com/Azure/aks-engine/issues/794))
- update go toolchain to 1.12.1 ([#792](https://github.com/Azure/aks-engine/issues/792))
- Include billing extensions for Azure US Government Cloud ([#780](https://github.com/Azure/aks-engine/issues/780))
- enable Kubernetes v1.14.0-beta.2 ([#756](https://github.com/Azure/aks-engine/issues/756))
- update cluster-autoscaler versions ([#749](https://github.com/Azure/aks-engine/issues/749))
- enable no_template with Azure Stack ([#677](https://github.com/Azure/aks-engine/issues/677))

### Revert Change ◀️
- remove systemd/xenial-proposed install ([#763](https://github.com/Azure/aks-engine/issues/763))
- update cluster-autoscaler versions ([#749](https://github.com/Azure/aks-engine/issues/749))

### Testing 💚
- fix NetworkPolicy tests for pre-1.11 cluster configurations ([#821](https://github.com/Azure/aks-engine/issues/821))
- add test case for assignDefaultAddonVals to deal with nil “Enabled” ([#815](https://github.com/Azure/aks-engine/issues/815))
- wait the regular, configurable time for dns-liveness pod readiness ([#826](https://github.com/Azure/aks-engine/issues/826))
- enforce --image-pull-policy=IfNotPresent for windows deployments in e2e tests ([#811](https://github.com/Azure/aks-engine/issues/811))
- add tests for pkg/engine/engine.go ([#801](https://github.com/Azure/aks-engine/issues/801))
- add unit tests for engine.go ([#787](https://github.com/Azure/aks-engine/issues/787))
- add more unit tests for no_template refactor ([#779](https://github.com/Azure/aks-engine/issues/779))
- more apiloader unit tests ([#751](https://github.com/Azure/aks-engine/issues/751))
- add node label and annotation validation ([#628](https://github.com/Azure/aks-engine/issues/628))
- add unit tests for pkg/api/LoadDefaultContainerServiceProperties() ([#743](https://github.com/Azure/aks-engine/issues/743))
- don't hardcode base 64 encoded script values part 2 ([#739](https://github.com/Azure/aks-engine/issues/739))
- describe pod object in error scenario ([#691](https://github.com/Azure/aks-engine/issues/691))
- don't hardcode base 64 encoded script values in template tests ([#702](https://github.com/Azure/aks-engine/issues/702))
- remove obsolete allocateNodeCidrs var ([#706](https://github.com/Azure/aks-engine/issues/706))
- fix template refactor unit tests ([#687](https://github.com/Azure/aks-engine/issues/687))
- check error returned when checking outbound connection of php-apache pod ([#670](https://github.com/Azure/aks-engine/issues/670))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.32.3"></a>
# [v0.32.3] - 2019-03-15
### Bug Fixes 🐞
- mark walinux for hold in cloud-init ([#778](https://github.com/Azure/aks-engine/issues/778))
- actually apt-mark hold walinuxagent during all CSE runs ([#771](https://github.com/Azure/aks-engine/issues/771))
- always configure the container runtime for GPU in an N series context ([#675](https://github.com/Azure/aks-engine/issues/675))

### Revert Change ◀️
- remove systemd/xenial-proposed install ([#763](https://github.com/Azure/aks-engine/issues/763))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.32.2"></a>
# [v0.32.2] - 2019-03-15
### Bug Fixes 🐞
- mark walinux for hold in cloud-init ([#778](https://github.com/Azure/aks-engine/issues/778))
- always configure the container runtime for GPU in an N series context ([#675](https://github.com/Azure/aks-engine/issues/675))

### Revert Change ◀️
- remove systemd/xenial-proposed install ([#763](https://github.com/Azure/aks-engine/issues/763))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.32.1"></a>
# [v0.32.1] - 2019-03-12
### Features 🌈
- windows agentpool support for AKS ([#714](https://github.com/Azure/aks-engine/issues/714))
- add selective agentpool upgrade support - VMSS ([#672](https://github.com/Azure/aks-engine/issues/672))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.32.0"></a>
# [v0.32.0] - 2019-03-06
### Bug Fixes 🐞
- enable systemd proposed implementation for non-VHD ([#668](https://github.com/Azure/aks-engine/issues/668))
- always assign target version on upgrade ([#650](https://github.com/Azure/aks-engine/issues/650))
- Remove retaining custom properties logic for master nodes ([#649](https://github.com/Azure/aks-engine/issues/649))
- address govet and gosimple issues ([#646](https://github.com/Azure/aks-engine/issues/646))
- simplify comparison to satisfy linter ([#629](https://github.com/Azure/aks-engine/issues/629))
- Explicitly use overlay2 storage driver while creation of vhd, by default docker uses aufs ([#574](https://github.com/Azure/aks-engine/issues/574))
- CONTAINERD_VERSION unbound variable in VHD build script ([#583](https://github.com/Azure/aks-engine/issues/583))
- remove extra \n in log messages ([#581](https://github.com/Azure/aks-engine/issues/581))
- allow for gzip and b64 encoding for raw Addon data by decoding the input beforehand ([#565](https://github.com/Azure/aks-engine/issues/565))
- silence `k` script ([#571](https://github.com/Azure/aks-engine/issues/571))
- Updating Windows default Docker version from 18.09.0 to 18.09.2. ([#550](https://github.com/Azure/aks-engine/issues/550))
- move the definition of inboundNatRules inside the loadbalancer specification ([#538](https://github.com/Azure/aks-engine/issues/538))
- Allow to create only masters so that agents can be created at a later… ([#529](https://github.com/Azure/aks-engine/issues/529))
- add comma separation to bash variables in provisionScriptParametersCommon ([#539](https://github.com/Azure/aks-engine/issues/539))
- added defaultValue to accommodate missing privateAzureRegistryServer param ([#526](https://github.com/Azure/aks-engine/issues/526))
- do not create route table when using flannel ([#486](https://github.com/Azure/aks-engine/issues/486))
- do not create ELB for private clusters ([#487](https://github.com/Azure/aks-engine/issues/487))
- set nodes to panic and reboot on oom ([#503](https://github.com/Azure/aks-engine/issues/503))
- upgrade blobfuse version for bug fix ([#493](https://github.com/Azure/aks-engine/issues/493))
- remove the use of deprecated Next() method ([#500](https://github.com/Azure/aks-engine/issues/500))

### Continuous Integration 💜
- update build status badge to ci pipeline ([#656](https://github.com/Azure/aks-engine/issues/656))
- generate release notes for component versions installed in VHD build ([#450](https://github.com/Azure/aks-engine/issues/450))
- add E2E pipeline status to README ([#641](https://github.com/Azure/aks-engine/issues/641))
- DRY_RUN variable is a string not a bool ([#601](https://github.com/Azure/aks-engine/issues/601))
- add storage account cleanup to VHD pipeline ([#589](https://github.com/Azure/aks-engine/issues/589))
- ignore test files in code coverage ([#559](https://github.com/Azure/aks-engine/issues/559))

### Documentation 📘
- update E2E pipeline process in developer docs ([#642](https://github.com/Azure/aks-engine/issues/642))
- add more package documentation for aks-engine's godoc ([#644](https://github.com/Azure/aks-engine/issues/644))
- upgrade --force ([#639](https://github.com/Azure/aks-engine/issues/639))
- add how to install aks-engine with Homebrew ([#562](https://github.com/Azure/aks-engine/issues/562))
- updated acs to aks ([#524](https://github.com/Azure/aks-engine/issues/524))
- clarify that `make ensure-generated` is a PR requirement ([#513](https://github.com/Azure/aks-engine/issues/513))
- add a "code of conduct" file where GitHub can see it ([#510](https://github.com/Azure/aks-engine/issues/510))
- adding godoc package comments ([#502](https://github.com/Azure/aks-engine/issues/502))

### Features 🌈
- Copy custom annotations, labels, taints from old to new nodes during upgrade ([#570](https://github.com/Azure/aks-engine/issues/570))
- default to Kubernetes 1.11 ([#625](https://github.com/Azure/aks-engine/issues/625))
- Azure [#402](https://github.com/Azure/aks-engine/issues/402) - Disable windows auto update ([#599](https://github.com/Azure/aks-engine/issues/599))
- add force parameter to the upgrade command ([#525](https://github.com/Azure/aks-engine/issues/525))
- Multiple SSH keys for Linux & Windows nodes ([#592](https://github.com/Azure/aks-engine/issues/592))
- Multiple SSH keys for Linux & Windows nodes ([#582](https://github.com/Azure/aks-engine/issues/582))
- Support ARM endpoint with non-trusted certificate ([#553](https://github.com/Azure/aks-engine/issues/553))
- enable 1.14.0-alpha.2 support ([#435](https://github.com/Azure/aks-engine/issues/435))
- enable vmss overprovision, do not run extensions on overprovisioned VMs ([#367](https://github.com/Azure/aks-engine/issues/367))
- add human-readable output option to "get-versions" command ([#527](https://github.com/Azure/aks-engine/issues/527))
- support ADFS for azurestack ([#531](https://github.com/Azure/aks-engine/issues/531))
- enable configurable --enforce-node-allocatable ([#535](https://github.com/Azure/aks-engine/issues/535))
- Add Cilium 1.4 support and improve deployment process ([#508](https://github.com/Azure/aks-engine/issues/508))
- add support for private registry ([#523](https://github.com/Azure/aks-engine/issues/523))
- support dynamic region for azurestack  ([#505](https://github.com/Azure/aks-engine/issues/505))
- add "get-versions" command to replace "orchestrators" ([#448](https://github.com/Azure/aks-engine/issues/448))

### Maintenance 🔧
- update VHD to 2019.03.05 ([#654](https://github.com/Azure/aks-engine/issues/654))
- add go report card badge to README ([#658](https://github.com/Azure/aks-engine/issues/658))
- add systemd/xenial-proposed to list of installed packages ([#655](https://github.com/Azure/aks-engine/issues/655))
- install -proposed systemd package ([#630](https://github.com/Azure/aks-engine/issues/630))
- update hybrid api model to "enableAutomaticUpdates": false ([#623](https://github.com/Azure/aks-engine/issues/623))
- add serbrech to OWNERS ([#640](https://github.com/Azure/aks-engine/issues/640))
- rev the nvidia driver version for the newer driver ([#634](https://github.com/Azure/aks-engine/issues/634))
- remove some invalid acs-engine references in aks-engine ([#636](https://github.com/Azure/aks-engine/issues/636))
- update the deis docker go-dev image to v1.19.1 ([#631](https://github.com/Azure/aks-engine/issues/631))
- update VHD reference to 2019.02.28 ([#621](https://github.com/Azure/aks-engine/issues/621))
- add support for k8s version 1.11.8 ([#615](https://github.com/Azure/aks-engine/issues/615))
- enable 1.11.8 in VHD, deprecate 1.11.6 ([#614](https://github.com/Azure/aks-engine/issues/614))
- add support for k8s version 1.13.4 ([#611](https://github.com/Azure/aks-engine/issues/611))
- add system.conf with static JoinControllers configuration ([#605](https://github.com/Azure/aks-engine/issues/605))
- update go-dev image ([#606](https://github.com/Azure/aks-engine/issues/606))
- modify eviction to memory.available.750Mi ([#604](https://github.com/Azure/aks-engine/issues/604))
- add priorityClassName: system-node-critical to kube-system,… ([#555](https://github.com/Azure/aks-engine/issues/555))
- update VHD reference to 2019.02.26 ([#600](https://github.com/Azure/aks-engine/issues/600))
- change vmss casing across code base ([#602](https://github.com/Azure/aks-engine/issues/602))
- enable Kubernetes v1.14.0-beta.1 ([#588](https://github.com/Azure/aks-engine/issues/588))
- enable Kubernetes v1.12.6 ([#587](https://github.com/Azure/aks-engine/issues/587))
- update the compute apiversion to 2018-10-01 to support new vmss features ([#578](https://github.com/Azure/aks-engine/issues/578))
- disable ip-masq-agent when using cilium networkPlugin ([#551](https://github.com/Azure/aks-engine/issues/551))
- apiserverCertificate should be apiServerCertificate ([#564](https://github.com/Azure/aks-engine/issues/564))
- add administration tools ([#554](https://github.com/Azure/aks-engine/issues/554))
- ignore artifacts that may be present from prior commits ([#556](https://github.com/Azure/aks-engine/issues/556))
- fix HasZonesForAllAgentPools() for masters-only scenarios ([#540](https://github.com/Azure/aks-engine/issues/540))
- install specific containerd version, make version configurable ([#516](https://github.com/Azure/aks-engine/issues/516))
- update AKS VHD images to 2019.02.13 ([#514](https://github.com/Azure/aks-engine/issues/514))
- deprecate kubernetes versions 1.8 and 1.7 ([#136](https://github.com/Azure/aks-engine/issues/136))

### Revert Change ◀️
- commit generated code files ([#546](https://github.com/Azure/aks-engine/issues/546))

### Testing 💚
- increase e2e pod delete timeout tolerance to 5 mins ([#662](https://github.com/Azure/aks-engine/issues/662))
- pass in skip_test env variable down to the test container ([#663](https://github.com/Azure/aks-engine/issues/663))
- add unit tests for azureconst.go ([#627](https://github.com/Azure/aks-engine/issues/627))
- enable timeouts for running E2E commands ([#563](https://github.com/Azure/aks-engine/issues/563))
- add VMSS client interfaces and mocks to make upgrade testable ([#590](https://github.com/Azure/aks-engine/issues/590))
- wait longer when incrementing error count waiting for pod Ready ([#560](https://github.com/Azure/aks-engine/issues/560))
- use "k" script to fetch the right kubectl just-in-time ([#558](https://github.com/Azure/aks-engine/issues/558))
- wait longer between kubectl top nodes retries ([#543](https://github.com/Azure/aks-engine/issues/543))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.31.2"></a>
# [v0.31.2] - 2019-03-01
#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.31.3"></a>
# [v0.31.3] - 2019-02-25
### Bug Fixes 🐞
- allow for gzip and b64 encoding for raw Addon data by decoding the input beforehand ([#565](https://github.com/Azure/aks-engine/issues/565))

### Maintenance 🔧
- update VHD reference to 2019.02.28 ([#621](https://github.com/Azure/aks-engine/issues/621))
- add support for k8s version 1.11.8 ([#615](https://github.com/Azure/aks-engine/issues/615))
- enable 1.11.8 in VHD, deprecate 1.11.6 ([#614](https://github.com/Azure/aks-engine/issues/614))
- add support for k8s version 1.13.4 ([#611](https://github.com/Azure/aks-engine/issues/611))
- enable Kubernetes v1.12.6 ([#587](https://github.com/Azure/aks-engine/issues/587))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.31.1"></a>
# [v0.31.1] - 2019-02-15
### Bug Fixes 🐞
- added defaultValue to accommodate missing privateAzureRegistryServer param ([#526](https://github.com/Azure/aks-engine/issues/526))

### Features 🌈
- add support for private registry ([#523](https://github.com/Azure/aks-engine/issues/523))

### Maintenance 🔧
- install specific containerd version, make version configurable ([#516](https://github.com/Azure/aks-engine/issues/516))
- update AKS VHD images to 2019.02.13 ([#514](https://github.com/Azure/aks-engine/issues/514))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.31.0"></a>
# [v0.31.0] - 2019-02-13
### Bug Fixes 🐞
- remove repair-malformed-updates deprecated flag ([#492](https://github.com/Azure/aks-engine/issues/492))
- fix unmanaged disk class issue ([#473](https://github.com/Azure/aks-engine/issues/473))
- Fail deployCmd when orchestrator version is invalid ([#355](https://github.com/Azure/aks-engine/issues/355))
- bump version for keyvault-flexvolume ([#465](https://github.com/Azure/aks-engine/issues/465))
- update kms container image ([#464](https://github.com/Azure/aks-engine/issues/464))
- Initialize data disks on Windows nodes ([#462](https://github.com/Azure/aks-engine/issues/462))

### Build 🏭
- commit generated code files ([#449](https://github.com/Azure/aks-engine/issues/449))

### Code Refactoring 💎
- swap satori's uuid package to gofrs ([#466](https://github.com/Azure/aks-engine/issues/466))

### Continuous Integration 💜
- use diff instead of git diff for the CI validation checks of the generated files ([#491](https://github.com/Azure/aks-engine/issues/491))
- fix failing unit test in master ([#485](https://github.com/Azure/aks-engine/issues/485))

### Features 🌈
- add support for Kubernetes 1.10.13 ([#497](https://github.com/Azure/aks-engine/issues/497))
- Support dynamic AzureEnvironmentSpecConfig ([#386](https://github.com/Azure/aks-engine/issues/386))
- update VHD image to 2019.02.12 ([#482](https://github.com/Azure/aks-engine/issues/482))
- updated network monitor version to newer release v0.0.5 ([#460](https://github.com/Azure/aks-engine/issues/460))

### Maintenance 🔧
- update Moby version to 3.0.4 ([#480](https://github.com/Azure/aks-engine/issues/480))

### Performance Improvements 🚀
- set caching mode as ReadOnly by default ([#476](https://github.com/Azure/aks-engine/issues/476))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.30.1"></a>
# [v0.30.1] - 2019-02-11
### Features 🌈
- update VHD image to 2019.02.12

### Maintenance 🔧
- update Moby version to 3.0.4

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.30.0"></a>
# [v0.30.0] - 2019-02-07
### Bug Fixes 🐞
- containerd downloads every time ([#458](https://github.com/Azure/aks-engine/issues/458))
- prevent race condition with docker/moby install ([#447](https://github.com/Azure/aks-engine/issues/447))
- reduce etcd download retries to avoid timeouts ([#437](https://github.com/Azure/aks-engine/issues/437))
- rename acsengineVersion tag to aksEngineVersion ([#264](https://github.com/Azure/aks-engine/issues/264))
- add toleration to run flannel on not-yet-ready nodes. ([#313](https://github.com/Azure/aks-engine/issues/313))
- ServiceCidr + DNSServiceIP validation error message ([#360](https://github.com/Azure/aks-engine/issues/360))
- the number of probes is of type int32 and not string ([#358](https://github.com/Azure/aks-engine/issues/358))
- Add ILB for vmss masters ([#319](https://github.com/Azure/aks-engine/issues/319))
- use correct OS Type in description text ([#323](https://github.com/Azure/aks-engine/issues/323))
- revert PodDisruptionBudget definitions ([#300](https://github.com/Azure/aks-engine/issues/300))
- set heapster Deployment to EnsureExists to prevent heapster-nann… ([#295](https://github.com/Azure/aks-engine/issues/295))
- replace some leftover acsengine by aksengine ([#268](https://github.com/Azure/aks-engine/issues/268))
- add missing container images in VHD ([#284](https://github.com/Azure/aks-engine/issues/284))
- address goimports checks ([#204](https://github.com/Azure/aks-engine/issues/204))
- addon manager pod discruption budgets ([#266](https://github.com/Azure/aks-engine/issues/266))
- Give cluster-autoscaler permission to list replicasets.apps ([#263](https://github.com/Azure/aks-engine/issues/263))
- **cilium:** Allow the daemonset to deploy regardless of node status. ([#405](https://github.com/Azure/aks-engine/issues/405))

### Code Refactoring 💎
- Remove agentpool index usage from function which gets agent VM name ([#428](https://github.com/Azure/aks-engine/issues/428))
- upgrade code refactor ([#287](https://github.com/Azure/aks-engine/issues/287))

### Continuous Integration 💜
- always collect logs in pipeline ([#415](https://github.com/Azure/aks-engine/issues/415))

### Documentation 📘
- new quickstart guide ([#301](https://github.com/Azure/aks-engine/issues/301))
- minor rewording/clarification of upgrade documentation ([#425](https://github.com/Azure/aks-engine/issues/425))
- SGX support with Kubernetes ([#438](https://github.com/Azure/aks-engine/issues/438))
- refer to the project as "AKS Engine" consistently ([#436](https://github.com/Azure/aks-engine/issues/436))
- upgrade + cluster-autoscaler notes ([#381](https://github.com/Azure/aks-engine/issues/381))
- fix all broken links ([#395](https://github.com/Azure/aks-engine/issues/395))
- deploy.md broken links fix ([#393](https://github.com/Azure/aks-engine/issues/393))
- don’t use upgrade if IaaS has been changed ([#387](https://github.com/Azure/aks-engine/issues/387))
- update documentation for container monitoring addon ([#368](https://github.com/Azure/aks-engine/issues/368))
- Update kubernetes-developers.md ([#370](https://github.com/Azure/aks-engine/issues/370))
- fix formatting of bash command in upgrade doc
- update cosmos link ([#375](https://github.com/Azure/aks-engine/issues/375))
- rewrite upgrade documentation ([#356](https://github.com/Azure/aks-engine/issues/356))
- fix a broken link and rename ACS to AKS ([#348](https://github.com/Azure/aks-engine/issues/348))
- fix broken link in clusterdefinitions.md ([#338](https://github.com/Azure/aks-engine/issues/338))
- update default value of smb-flexvolume in clusterdefinitions.md ([#334](https://github.com/Azure/aks-engine/issues/334))
- fix broken links ([#332](https://github.com/Azure/aks-engine/issues/332))
- Fix docs link to howto/README.md ([#309](https://github.com/Azure/aks-engine/issues/309))
- add note on AKS-Engine and AKS relationship ([#293](https://github.com/Azure/aks-engine/issues/293))
- restructure the documentation ([#253](https://github.com/Azure/aks-engine/issues/253))
- fix link to Azure Active Directory integration ([#267](https://github.com/Azure/aks-engine/issues/267))

### Features 🌈
- update VHD image to 2019.02.06 ([#456](https://github.com/Azure/aks-engine/issues/456))
- Add ability to enable windows ssh with open ssh install on node ([#433](https://github.com/Azure/aks-engine/issues/433))
- update Azure CNI version to 1.0.17 ([#416](https://github.com/Azure/aks-engine/issues/416))
- make Moby version configurable ([#407](https://github.com/Azure/aks-engine/issues/407))
- add the aksEngineVersion tag for VMSS agents ([#413](https://github.com/Azure/aks-engine/issues/413))
- Add CustomCloudProfile for Azure Stack ([#297](https://github.com/Azure/aks-engine/issues/297))
- set publicIPAllocation as static for master node ([#291](https://github.com/Azure/aks-engine/issues/291))
- enable kube-proxy ipvs
- add SGX driver installation on C-series VMs ([#318](https://github.com/Azure/aks-engine/issues/318))
- pre-pull flannel images ([#299](https://github.com/Azure/aks-engine/issues/299))
- update VHD version to 2019.01.11 ([#308](https://github.com/Azure/aks-engine/issues/308))
- use the Azure CLI as an auth method ([#238](https://github.com/Azure/aks-engine/issues/238))
- update heapster to 1.5.4 ([#257](https://github.com/Azure/aks-engine/issues/257))
- add localhost and 127.0.0.1 to certificates ([#243](https://github.com/Azure/aks-engine/issues/243))
- update img binary to v0.5.6 ([#259](https://github.com/Azure/aks-engine/issues/259))
- update kube-dns to v1.15.0 ([#258](https://github.com/Azure/aks-engine/issues/258))
- add new etcd versions, 3.2.25 is default ([#254](https://github.com/Azure/aks-engine/issues/254))
- update addon-resizer to v1.8.4 ([#256](https://github.com/Azure/aks-engine/issues/256))

### Maintenance 🔧
- add agentpool resource(used for agentpool level operations) and UT ([#439](https://github.com/Azure/aks-engine/issues/439))
- per-agent pool k8s versions and provisioning state ([#423](https://github.com/Azure/aks-engine/issues/423))
- add support for k8s version 1.13.3 ([#430](https://github.com/Azure/aks-engine/issues/430))
- increase auto-generated cert expiration to 30 years ([#396](https://github.com/Azure/aks-engine/issues/396))
- add commit message categories to PR template ([#421](https://github.com/Azure/aks-engine/issues/421))
- add support for k8s version 1.11.7 ([#379](https://github.com/Azure/aks-engine/issues/379))
- add support for Kubernetes 1.14.0-alpha.1 ([#330](https://github.com/Azure/aks-engine/issues/330))
- add support for Kubernetes 1.12.5 ([#331](https://github.com/Azure/aks-engine/issues/331))
- update Azure CNI to v1.0.16 ([#290](https://github.com/Azure/aks-engine/issues/290))
- update git-chglog configuration for releases ([#314](https://github.com/Azure/aks-engine/issues/314))
- add config for welcome bot ([#296](https://github.com/Azure/aks-engine/issues/296))
- add support for Kubernetes 1.13.2 ([#292](https://github.com/Azure/aks-engine/issues/292))
- update go-dev image for newer golangci-lint ([#252](https://github.com/Azure/aks-engine/issues/252))

### Revert Change ◀️
- JoinControllers system.conf override ([#410](https://github.com/Azure/aks-engine/issues/410))

### Testing 💚
- add k8s components versions test 1.13 ([#383](https://github.com/Azure/aks-engine/issues/383))
- print pod logs when errors occur ([#274](https://github.com/Azure/aks-engine/issues/274))
- print all pods at beginning of E2E test run ([#279](https://github.com/Azure/aks-engine/issues/279))
- enable multiple subnets per agent pool ([#262](https://github.com/Azure/aks-engine/issues/262))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.29.1"></a>
# [v0.29.1] - 2019-01-14
### Features 🌈
- update VHD version to 2019.01.11 ([#308](https://github.com/Azure/aks-engine/issues/308))

### Maintenance 🔧
- update git-chglog configuration for releases ([#314](https://github.com/Azure/aks-engine/issues/314))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.29.0"></a>
# [v0.29.0] - 2019-01-11
### Bug Fixes 🐞
- revert PodDisruptionBudget definitions ([#300](https://github.com/Azure/aks-engine/issues/300))
- set heapster Deployment to EnsureExists to prevent heapster-nann… ([#295](https://github.com/Azure/aks-engine/issues/295))
- replace some leftover acsengine by aksengine ([#268](https://github.com/Azure/aks-engine/issues/268))
- add missing container images in VHD ([#284](https://github.com/Azure/aks-engine/issues/284))
- address goimports checks ([#204](https://github.com/Azure/aks-engine/issues/204))
- addon manager pod discruption budgets ([#266](https://github.com/Azure/aks-engine/issues/266))
- rename acsengineVersion tag to aksEngineVersion ([#264](https://github.com/Azure/aks-engine/issues/264))
- Give cluster-autoscaler permission to list replicasets.apps ([#263](https://github.com/Azure/aks-engine/issues/263))
- Do not check agent & master pool size when using availability zones ([#206](https://github.com/Azure/aks-engine/issues/206))
- remove VM upgrade version check for AKS ([#236](https://github.com/Azure/aks-engine/issues/236))
- change default auth method for deploy,scale and upgrade ([#231](https://github.com/Azure/aks-engine/issues/231))
- use kube-system namespace instead of flex, kv ([#212](https://github.com/Azure/aks-engine/issues/212))
- add codecov token env var to Azure pipeline ([#230](https://github.com/Azure/aks-engine/issues/230))
- update MSI ([#185](https://github.com/Azure/aks-engine/issues/185))

### Continuous Integration 💜
- check that last commit message is conformant ([#237](https://github.com/Azure/aks-engine/issues/237))

### Documentation 📘
- add note on AKS-Engine and AKS relationship ([#293](https://github.com/Azure/aks-engine/issues/293))
- restructure the documentation ([#253](https://github.com/Azure/aks-engine/issues/253))
- fix link to Azure Active Directory integration ([#267](https://github.com/Azure/aks-engine/issues/267))
- fix Kubernetes release notes broken link ([#260](https://github.com/Azure/aks-engine/issues/260))
- update some docs with the right information ([#217](https://github.com/Azure/aks-engine/issues/217))

### Features 🌈
- use the Azure CLI as an auth method ([#238](https://github.com/Azure/aks-engine/issues/238))
- add localhost and 127.0.0.1 to certificates ([#243](https://github.com/Azure/aks-engine/issues/243))
- Add anti affinity to coredns deployment ([#249](https://github.com/Azure/aks-engine/issues/249))
- add stale app yaml ([#244](https://github.com/Azure/aks-engine/issues/244))
- add pod disruption budgets ([#218](https://github.com/Azure/aks-engine/issues/218))
- support `make build` on Windows ([#219](https://github.com/Azure/aks-engine/issues/219))
- Azure DevOps PR E2E pipeline ([#135](https://github.com/Azure/aks-engine/issues/135))

### Maintenance 🔧
- add support for Kubernetes 1.13.2 ([#292](https://github.com/Azure/aks-engine/issues/292))
- update go-dev image for newer golangci-lint ([#252](https://github.com/Azure/aks-engine/issues/252))
- remove commit check from Azure DevOps pipeline ([#250](https://github.com/Azure/aks-engine/issues/250))
- remove circleci status badge from README ([#251](https://github.com/Azure/aks-engine/issues/251))
- remove stale scripts ([#224](https://github.com/Azure/aks-engine/issues/224))
- update kubernetes-dashboard to v1.10.1 ([#232](https://github.com/Azure/aks-engine/issues/232))
- remove Kubernetes 1.6 heapster version from VHD script ([#233](https://github.com/Azure/aks-engine/issues/233))
- standardize commits to automate CHANGELOG ([#196](https://github.com/Azure/aks-engine/issues/196))
- Update go-dev tools image for go 1.11.4 ([#205](https://github.com/Azure/aks-engine/issues/205))

### Testing 💚
- enable multiple subnets per agent pool ([#262](https://github.com/Azure/aks-engine/issues/262))
- NetworkPolicy E2E tests cleanup for soak ([#246](https://github.com/Azure/aks-engine/issues/246))
- pre-delete ILB stuffs before running ([#228](https://github.com/Azure/aks-engine/issues/228))
- only check top nodes if RBAC is enabled ([#221](https://github.com/Azure/aks-engine/issues/221))
- more tolerance for kube-system pod startup in E2E ([#220](https://github.com/Azure/aks-engine/issues/220))

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.28.1"></a>
# [v0.28.1] - 2018-12-20
#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.28.0"></a>
# [v0.28.0] - 2018-12-19
#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
<a name="v0.27.0"></a>
# v0.27.0 - 2018-12-14
### Bug Fixes 🐞
- test/deploy.sh is usable standalone
- deleting a VM in failed provisioning state ([#1824](https://github.com/Azure/aks-engine/issues/1824))
- adding log-errors config file to acs-engine-test ([#1346](https://github.com/Azure/aks-engine/issues/1346))
- **Get-AzureConstants.py:** remove the date string from the header * add test to circle config * add make target to run test in test.mk * add script that executes the test
- **Makefile:** Call bootstrap before we run generate ([#1276](https://github.com/Azure/aks-engine/issues/1276))
- **Makefile:** run generate before test target ([#1264](https://github.com/Azure/aks-engine/issues/1264))
- **addons:** Adds missing cluster subnet substitution for flannel ([#3356](https://github.com/Azure/aks-engine/issues/3356))
- **ci:** use regions available in new CI subs ([#1157](https://github.com/Azure/aks-engine/issues/1157))
- **cni-networking:** Enables `br_netfilter` kernel module ([#3484](https://github.com/Azure/aks-engine/issues/3484))
- **deploy:** fix SP cred handling in deploy ([#1052](https://github.com/Azure/aks-engine/issues/1052))
- **deploy:** Change SP creation logging to debug ([#1121](https://github.com/Azure/aks-engine/issues/1121))
- **dockerfile:** pin version more correctly
- **docs:** use release-1.5 for influxdb/grafana install ([#1781](https://github.com/Azure/aks-engine/issues/1781))
- **docs:** small cleanups
- **e2e:** run bootstrap/build in e2e.sh ([#1158](https://github.com/Azure/aks-engine/issues/1158))
- **e2e:** remove keyvault tests from feature validation (for now) ([#1155](https://github.com/Azure/aks-engine/issues/1155))
- **gitignore:** ignore generated test report file ([#1229](https://github.com/Azure/aks-engine/issues/1229))
- **glide:** Pin the pflag, go-homedir, and cobra deps to a specific sha
- **kubernetes_test:** Wait for dashboard to be ready before performing test
- **metrics.go:** add env var called IS_JENKINS ([#1746](https://github.com/Azure/aks-engine/issues/1746))
- **oms:** pull oms directly from docker ([#3294](https://github.com/Azure/aks-engine/issues/3294))
- **tiller:** allow tiller resources to be edited by users ([#1319](https://github.com/Azure/aks-engine/issues/1319))
- **version:** don't include timestamp in version ([#1228](https://github.com/Azure/aks-engine/issues/1228))

### Code Style 🎶
- **test:** remove uneeded code

### Documentation 📘
- add initial contributing guide ([#188](https://github.com/Azure/aks-engine/issues/188))
- **README:** encourage getting latest Azure CLI 2.0 release ([#1042](https://github.com/Azure/aks-engine/issues/1042))
- **github:** use labeler to manage project labels ([#955](https://github.com/Azure/aks-engine/issues/955))
- **github:** include acs-engine version in issue template ([#943](https://github.com/Azure/aks-engine/issues/943))
- **github:** clarify issue template language
- **large-clusters:** add docs for building large k8s clusters ([#1001](https://github.com/Azure/aks-engine/issues/1001))
- **release:** address review feedback

### Features 🌈
- ***:** Bumps client-go to v7.0.0 ([#2954](https://github.com/Azure/aks-engine/issues/2954))
- **autodeploy:** test field population; test MSI+(skip)creds scenario
- **ci:** add dev guide, style checks
- **custom hyperkube:** managed-identity uses customHyperkubeImage
- **custom hyperkube:** support customHyperkubeImage in kubernetesConfig
- **deploy:** deploy auto-populates SP, SSH, DNSPrefix ([#925](https://github.com/Azure/aks-engine/issues/925))
- **ext:** Add choco extension with param support ([#2707](https://github.com/Azure/aks-engine/issues/2707))
- **k8s:** add support for 1.7.0 ([#909](https://github.com/Azure/aks-engine/issues/909))
- **kube-proxy:** schedule kube-proxy as critical pod ([#914](https://github.com/Azure/aks-engine/issues/914))
- **kubernetes:** add support for custom docker0 net ([#753](https://github.com/Azure/aks-engine/issues/753))
- **kubernetes_test:** Add outbound internet checks for both agent types ([#1502](https://github.com/Azure/aks-engine/issues/1502))
- **managed identity:** make SPP optional when using MSI
- **metrics:** Add the ability to collect performance metrics from test runs ([#1614](https://github.com/Azure/aks-engine/issues/1614))
- **msi:** add global tmpl fn for checking msi status
- **msi:** add useManagedIdentity to apimodel
- **msi:** template changes: windows agents
- **msi:** template changes: init + masters
- **msi:** test MSI regression as part of PR job
- **msi:** add example apimodel + jenkins config
- **msi:** enable test to source env file earlier
- **msi:** template changes: linux agents
- **msi:** mount MSI settings into req'd containers
- **outputs:** add additional values to output block
- **outputs:** emit subnetName used by agent pools
- **perf:** Invoke-WebRequest much slower then browser download ([#4294](https://github.com/Azure/aks-engine/issues/4294))
- **promote to failure:** Add the ability to track failures across multiple jobs ([#1403](https://github.com/Azure/aks-engine/issues/1403))
- **runner.go:** port bash ginkgo runner to go ([#1365](https://github.com/Azure/aks-engine/issues/1365))
- **v20170831:** agent pools for v20170831 must use ManagedDisks ([#1529](https://github.com/Azure/aks-engine/issues/1529))
- **vsts-ci.yaml:** Scaffolding to support VHD image building in vsts ([#3591](https://github.com/Azure/aks-engine/issues/3591))

### Maintenance 🔧
- **Makefile:** don't run 'make style' on 'make test' ([#1199](https://github.com/Azure/aks-engine/issues/1199))
- **README:** Remove managed disks private preview notice ([#2416](https://github.com/Azure/aks-engine/issues/2416))
- **ci:** run full suite of k8s tests on prs and master commits ([#1252](https://github.com/Azure/aks-engine/issues/1252))
- **circleci:** BUILD_NUMBER should be BUILD_NUM ([#1169](https://github.com/Azure/aks-engine/issues/1169))
- **coverage:** add `make coverage` command ([#1080](https://github.com/Azure/aks-engine/issues/1080))
- **dcos-e2e:** run dcos pr-e2e in southcentralus ([#1267](https://github.com/Azure/aks-engine/issues/1267))
- **deadcode:** remove dead code from deploy cmd
- **deps:** upgrade go-autorest to v10.15.4 ([#3843](https://github.com/Azure/aks-engine/issues/3843))
- **docs:** clarify large cluster docs ([#1409](https://github.com/Azure/aks-engine/issues/1409))
- **engine:** test case should be consistent, helpful
- **glide:** Remove prometheus/common dep
- **mocks:** we should not be sleeping in mocks
- **panics:** print stack trace when recovering
- **scripts:** remove unused rebase.sh ([#1232](https://github.com/Azure/aks-engine/issues/1232))
- **test infra:** fix CUSTOM_HYPERKUBE_SPEC handling
- **vendor:** add JiangtianLi/gettext to glide, hardcode vers, bump Azure vers ([#1146](https://github.com/Azure/aks-engine/issues/1146))

### Testing 💚
- k8s: check node healthy count at the end
- k8s: only test ACR in regions where it is available
- k8s: test acr integration
- fix AZURE_CONFIG_DIR
- **outputs:** test outputs for kubernetes

#### Please report any issues here: https://github.com/Azure/aks-engine/issues/new
[Unreleased]: https://github.com/Azure/aks-engine/compare/v0.64.0...HEAD
[v0.64.0]: https://github.com/Azure/aks-engine/compare/v0.63.0...v0.64.0
[v0.63.0]: https://github.com/Azure/aks-engine/compare/v0.62.1...v0.63.0
[v0.62.1]: https://github.com/Azure/aks-engine/compare/v0.62.0...v0.62.1
[v0.62.0]: https://github.com/Azure/aks-engine/compare/v0.56.3...v0.62.0
[v0.56.3]: https://github.com/Azure/aks-engine/compare/v0.61.0...v0.56.3
[v0.61.0]: https://github.com/Azure/aks-engine/compare/v0.56.2...v0.61.0
[v0.56.2]: https://github.com/Azure/aks-engine/compare/v0.60.1...v0.56.2
[v0.60.1]: https://github.com/Azure/aks-engine/compare/v0.56.1...v0.60.1
[v0.56.1]: https://github.com/Azure/aks-engine/compare/v0.60.0...v0.56.1
[v0.60.0]: https://github.com/Azure/aks-engine/compare/v0.59.0...v0.60.0
[v0.59.0]: https://github.com/Azure/aks-engine/compare/v0.58.0...v0.59.0
[v0.58.0]: https://github.com/Azure/aks-engine/compare/v0.57.0...v0.58.0
[v0.57.0]: https://github.com/Azure/aks-engine/compare/v0.56.0...v0.57.0
[v0.56.0]: https://github.com/Azure/aks-engine/compare/v0.55.4...v0.56.0
[v0.55.4]: https://github.com/Azure/aks-engine/compare/v0.55.3...v0.55.4
[v0.55.3]: https://github.com/Azure/aks-engine/compare/v0.55.2...v0.55.3
[v0.55.2]: https://github.com/Azure/aks-engine/compare/v0.55.1...v0.55.2
[v0.55.1]: https://github.com/Azure/aks-engine/compare/v0.55.0...v0.55.1
[v0.55.0]: https://github.com/Azure/aks-engine/compare/v0.50.3...v0.55.0
[v0.50.3]: https://github.com/Azure/aks-engine/compare/v0.51.1...v0.50.3
[v0.51.1]: https://github.com/Azure/aks-engine/compare/v0.52.1...v0.51.1
[v0.52.1]: https://github.com/Azure/aks-engine/compare/v0.53.1...v0.52.1
[v0.53.1]: https://github.com/Azure/aks-engine/compare/v0.54.1...v0.53.1
[v0.54.1]: https://github.com/Azure/aks-engine/compare/v0.54.0...v0.54.1
[v0.54.0]: https://github.com/Azure/aks-engine/compare/v0.53.0...v0.54.0
[v0.53.0]: https://github.com/Azure/aks-engine/compare/v0.52.0...v0.53.0
[v0.52.0]: https://github.com/Azure/aks-engine/compare/v0.51.0...v0.52.0
[v0.51.0]: https://github.com/Azure/aks-engine/compare/v0.50.2...v0.51.0
[v0.50.2]: https://github.com/Azure/aks-engine/compare/v0.50.1...v0.50.2
[v0.50.1]: https://github.com/Azure/aks-engine/compare/v0.50.0...v0.50.1
[v0.50.0]: https://github.com/Azure/aks-engine/compare/v0.49.0...v0.50.0
[v0.49.0]: https://github.com/Azure/aks-engine/compare/v0.48.0...v0.49.0
[v0.48.0]: https://github.com/Azure/aks-engine/compare/v0.47.0...v0.48.0
[v0.47.0]: https://github.com/Azure/aks-engine/compare/v0.46.3...v0.47.0
[v0.46.3]: https://github.com/Azure/aks-engine/compare/v0.46.2...v0.46.3
[v0.46.2]: https://github.com/Azure/aks-engine/compare/v0.46.1...v0.46.2
[v0.46.1]: https://github.com/Azure/aks-engine/compare/v0.46.0...v0.46.1
[v0.46.0]: https://github.com/Azure/aks-engine/compare/v0.43.3...v0.46.0
[v0.43.3]: https://github.com/Azure/aks-engine/compare/v0.45.0...v0.43.3
[v0.45.0]: https://github.com/Azure/aks-engine/compare/v0.44.2...v0.45.0
[v0.44.2]: https://github.com/Azure/aks-engine/compare/v0.44.1...v0.44.2
[v0.44.1]: https://github.com/Azure/aks-engine/compare/v0.44.0...v0.44.1
[v0.44.0]: https://github.com/Azure/aks-engine/compare/v0.43.2...v0.44.0
[v0.43.2]: https://github.com/Azure/aks-engine/compare/v0.43.1...v0.43.2
[v0.43.1]: https://github.com/Azure/aks-engine/compare/v0.43.0...v0.43.1
[v0.43.0]: https://github.com/Azure/aks-engine/compare/v0.41.5...v0.43.0
[v0.41.5]: https://github.com/Azure/aks-engine/compare/v0.42.2...v0.41.5
[v0.42.2]: https://github.com/Azure/aks-engine/compare/v0.42.1...v0.42.2
[v0.42.1]: https://github.com/Azure/aks-engine/compare/v0.42.0...v0.42.1
[v0.42.0]: https://github.com/Azure/aks-engine/compare/v0.40.2...v0.42.0
[v0.40.2]: https://github.com/Azure/aks-engine/compare/v0.41.4...v0.40.2
[v0.41.4]: https://github.com/Azure/aks-engine/compare/v0.41.3...v0.41.4
[v0.41.3]: https://github.com/Azure/aks-engine/compare/v0.40.1...v0.41.3
[v0.40.1]: https://github.com/Azure/aks-engine/compare/v0.41.2...v0.40.1
[v0.41.2]: https://github.com/Azure/aks-engine/compare/v0.38.9...v0.41.2
[v0.38.9]: https://github.com/Azure/aks-engine/compare/v0.41.1...v0.38.9
[v0.41.1]: https://github.com/Azure/aks-engine/compare/v0.41.0...v0.41.1
[v0.41.0]: https://github.com/Azure/aks-engine/compare/v0.40.0...v0.41.0
[v0.40.0]: https://github.com/Azure/aks-engine/compare/v0.38.8...v0.40.0
[v0.38.8]: https://github.com/Azure/aks-engine/compare/v0.39.2...v0.38.8
[v0.39.2]: https://github.com/Azure/aks-engine/compare/v0.39.1...v0.39.2
[v0.39.1]: https://github.com/Azure/aks-engine/compare/v0.38.7...v0.39.1
[v0.38.7]: https://github.com/Azure/aks-engine/compare/v0.38.6...v0.38.7
[v0.38.6]: https://github.com/Azure/aks-engine/compare/v0.39.0...v0.38.6
[v0.39.0]: https://github.com/Azure/aks-engine/compare/v0.38.5...v0.39.0
[v0.38.5]: https://github.com/Azure/aks-engine/compare/v0.38.4...v0.38.5
[v0.38.4]: https://github.com/Azure/aks-engine/compare/v0.38.3...v0.38.4
[v0.38.3]: https://github.com/Azure/aks-engine/compare/v0.38.2...v0.38.3
[v0.38.2]: https://github.com/Azure/aks-engine/compare/v0.37.5...v0.38.2
[v0.37.5]: https://github.com/Azure/aks-engine/compare/v0.38.1...v0.37.5
[v0.38.1]: https://github.com/Azure/aks-engine/compare/v0.38.0...v0.38.1
[v0.38.0]: https://github.com/Azure/aks-engine/compare/v0.37.4...v0.38.0
[v0.37.4]: https://github.com/Azure/aks-engine/compare/v0.35.6...v0.37.4
[v0.35.6]: https://github.com/Azure/aks-engine/compare/v0.37.2...v0.35.6
[v0.37.2]: https://github.com/Azure/aks-engine/compare/v0.37.3...v0.37.2
[v0.37.3]: https://github.com/Azure/aks-engine/compare/v0.37.1...v0.37.3
[v0.37.1]: https://github.com/Azure/aks-engine/compare/v0.35.5...v0.37.1
[v0.35.5]: https://github.com/Azure/aks-engine/compare/v0.37.0...v0.35.5
[v0.37.0]: https://github.com/Azure/aks-engine/compare/v0.36.5...v0.37.0
[v0.36.5]: https://github.com/Azure/aks-engine/compare/v0.36.4...v0.36.5
[v0.36.4]: https://github.com/Azure/aks-engine/compare/v0.36.3...v0.36.4
[v0.36.3]: https://github.com/Azure/aks-engine/compare/v0.35.4...v0.36.3
[v0.35.4]: https://github.com/Azure/aks-engine/compare/v0.36.2...v0.35.4
[v0.36.2]: https://github.com/Azure/aks-engine/compare/v0.36.1...v0.36.2
[v0.36.1]: https://github.com/Azure/aks-engine/compare/v0.35.3...v0.36.1
[v0.35.3]: https://github.com/Azure/aks-engine/compare/v0.36.0...v0.35.3
[v0.36.0]: https://github.com/Azure/aks-engine/compare/v0.35.2...v0.36.0
[v0.35.2]: https://github.com/Azure/aks-engine/compare/v0.35.1...v0.35.2
[v0.35.1]: https://github.com/Azure/aks-engine/compare/v0.35.0...v0.35.1
[v0.35.0]: https://github.com/Azure/aks-engine/compare/v0.34.2...v0.35.0
[v0.34.2]: https://github.com/Azure/aks-engine/compare/v0.34.3...v0.34.2
[v0.34.3]: https://github.com/Azure/aks-engine/compare/v0.33.6...v0.34.3
[v0.33.6]: https://github.com/Azure/aks-engine/compare/v0.34.1...v0.33.6
[v0.34.1]: https://github.com/Azure/aks-engine/compare/v0.34.0...v0.34.1
[v0.34.0]: https://github.com/Azure/aks-engine/compare/v0.33.5...v0.34.0
[v0.33.5]: https://github.com/Azure/aks-engine/compare/v0.33.4...v0.33.5
[v0.33.4]: https://github.com/Azure/aks-engine/compare/v0.33.3...v0.33.4
[v0.33.3]: https://github.com/Azure/aks-engine/compare/v0.33.2...v0.33.3
[v0.33.2]: https://github.com/Azure/aks-engine/compare/v0.33.1...v0.33.2
[v0.33.1]: https://github.com/Azure/aks-engine/compare/v0.33.0...v0.33.1
[v0.33.0]: https://github.com/Azure/aks-engine/compare/v0.32.3...v0.33.0
[v0.32.3]: https://github.com/Azure/aks-engine/compare/v0.32.2...v0.32.3
[v0.32.2]: https://github.com/Azure/aks-engine/compare/v0.32.1...v0.32.2
[v0.32.1]: https://github.com/Azure/aks-engine/compare/v0.32.0...v0.32.1
[v0.32.0]: https://github.com/Azure/aks-engine/compare/v0.31.2...v0.32.0
[v0.31.2]: https://github.com/Azure/aks-engine/compare/v0.31.3...v0.31.2
[v0.31.3]: https://github.com/Azure/aks-engine/compare/v0.31.1...v0.31.3
[v0.31.1]: https://github.com/Azure/aks-engine/compare/v0.31.0...v0.31.1
[v0.31.0]: https://github.com/Azure/aks-engine/compare/v0.30.1...v0.31.0
[v0.30.1]: https://github.com/Azure/aks-engine/compare/v0.30.0...v0.30.1
[v0.30.0]: https://github.com/Azure/aks-engine/compare/v0.29.1...v0.30.0
[v0.29.1]: https://github.com/Azure/aks-engine/compare/v0.29.0...v0.29.1
[v0.29.0]: https://github.com/Azure/aks-engine/compare/v0.28.1...v0.29.0
[v0.28.1]: https://github.com/Azure/aks-engine/compare/v0.28.0...v0.28.1
[v0.28.0]: https://github.com/Azure/aks-engine/compare/v0.27.0...v0.28.0
