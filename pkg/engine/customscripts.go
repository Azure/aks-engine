// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

func GetKubernetesB64CSEMain() string {
	return getBase64EncodedGzippedCustomScript(kubernetesCSEMainScript)
}

func GetKubernetesB64CSEHelpers() string {
	return getBase64EncodedGzippedCustomScript(kubernetesCSEHelpersScript)
}

func GetKubernetesB64HealthMonitorScript() string {
	return getBase64EncodedGzippedCustomScript(kubernetesHealthMonitorScript)
}

func GetKubernetesB64KubeletMonitorSystemdTimer() string {
	return getBase64EncodedGzippedCustomScript(kubernetesKubeletMonitorSystemdTimer)
}

func GetKubernetesB64KubeletMonitorSystemdService() string {
	return getBase64EncodedGzippedCustomScript(kubernetesKubeletMonitorSystemdService)
}

func GetKubernetesB64DockerMonitorSystemdTimer() string {
	return getBase64EncodedGzippedCustomScript(kubernetesDockerMonitorSystemdTimer)
}

func GetKubernetesB64DockerMonitorSystemdService() string {
	return getBase64EncodedGzippedCustomScript(kubernetesDockerMonitorSystemdService)
}

func GetKubernetesB64CSEInstall() string {
	return getBase64EncodedGzippedCustomScript(kubernetesCSEInstall)
}

func GetKubernetesB64CIS() string {
	return getBase64EncodedGzippedCustomScript(kubernetesCISScript)
}

func GetKubernetesB64CSEConfig() string {
	return getBase64EncodedGzippedCustomScript(kubernetesCSEConfig)
}

func GetKubernetesB64CSECustomCloud() string {
	return getBase64EncodedGzippedCustomScript(kubernetesCSECustomCloud)
}

func GetKubernetesB64MountEtcd() string {
	return getBase64EncodedGzippedCustomScript(kubernetesMountEtcd)
}

func GetKubernetesB64CustomSearchDomainsScript() string {
	return getBase64EncodedGzippedCustomScript(kubernetesCustomSearchDomainsScript)
}

func GetKubernetesB64GenerateProxyCerts() string {
	return getBase64EncodedGzippedCustomScript(kubernetesMasterGenerateProxyCertsScript)
}

func GetB64sshdConfig() string {
	return getBase64EncodedGzippedCustomScript(sshdConfig)
}

func GetB64systemConf() string {
	return getBase64EncodedGzippedCustomScript(systemConf)
}

func GetKubernetesKubeletSystemdService() string {
	return getBase64EncodedGzippedCustomScript(kubeletSystemdService)
}

func GetKubernetesKMSSystemdService() string {
	return getBase64EncodedGzippedCustomScript(kmsSystemdService)
}

func GetKubernetesB64AptPreferences() string {
	return getBase64EncodedGzippedCustomScript(aptPreferences)
}

func GetKubernetesB64DockerClearMountPropagationFlags() string {
	return getBase64EncodedGzippedCustomScript(dockerClearMountPropagationFlags)
}

func GetKubernetesB64SystemdBPFMount() string {
	return getBase64EncodedGzippedCustomScript(systemdBPFMount)
}

func GetKubernetesB64EtcdSystemdService() string {
	return getBase64EncodedGzippedCustomScript(etcdSystemdService)
}

func GetKubernetesB64EtcIssue() string {
	return getBase64EncodedGzippedCustomScript(etcIssue)
}
