// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

func GetKubernetesB64Provision() string {
	return getBase64CustomScript(kubernetesCustomScript)
}

func GetKubernetesB64ProvisionSource() string {
	return getBase64CustomScript(kubernetesProvisionSourceScript)
}

func GetKubernetesB64HealthMonitorScript() string {
	return getBase64CustomScript(kubernetesHealthMonitorScript)
}

func GetKubernetesB64Installs() string {
	return getBase64CustomScript(kubernetesInstalls)
}

func GetKubernetesB64Configs() string {
	return getBase64CustomScript(kubernetesConfigurations)
}

func GetKubernetesB64Mountetcd() string {
	return getBase64CustomScript(kubernetesMountetcd)
}

func GetKubernetesB64CustomSearchDomainsScript() string {
	return getBase64CustomScript(kubernetesCustomSearchDomainsScript)
}

func GetKubernetesB64GenerateProxyCerts() string {
	return getBase64CustomScript(kubernetesMasterGenerateProxyCertsScript)
}

func GetB64sshdConfig() string {
	return getBase64CustomScript(sshdConfig)
}

func GetB64systemConf() string {
	return getBase64CustomScript(systemConf)
}
