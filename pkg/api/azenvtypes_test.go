// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	. "github.com/onsi/gomega"

	"testing"
)

func TestURLForAzureChinaCloud(t *testing.T) {
	var azureChinaCloudMirror = "azk8s.cn"
	var publicCloudMCR = "mcr.microsoft.com"
	g := NewGomegaWithT(t)

	g.Expect(AzureChinaCloudSpec.KubernetesSpecConfig.KubernetesImageBase).To(ContainSubstring(azureChinaCloudMirror))
	g.Expect(AzureChinaCloudSpec.KubernetesSpecConfig.NVIDIAImageBase).To(ContainSubstring(publicCloudMCR))
	g.Expect(AzureChinaCloudSpec.KubernetesSpecConfig.AzureCNIImageBase).To(ContainSubstring(azureChinaCloudMirror))
	g.Expect(AzureChinaCloudSpec.KubernetesSpecConfig.CalicoImageBase).To(ContainSubstring(azureChinaCloudMirror))
	g.Expect(AzureChinaCloudSpec.KubernetesSpecConfig.CNIPluginsDownloadURL).To(ContainSubstring(azureChinaCloudMirror))
	g.Expect(AzureChinaCloudSpec.KubernetesSpecConfig.VnetCNILinuxPluginsDownloadURL).To(ContainSubstring(azureChinaCloudMirror))
	g.Expect(AzureChinaCloudSpec.KubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL).To(ContainSubstring(azureChinaCloudMirror))
	g.Expect(AzureChinaCloudSpec.KubernetesSpecConfig.ContainerdDownloadURLBase).To(ContainSubstring(azureChinaCloudMirror))
	g.Expect(AzureChinaCloudSpec.KubernetesSpecConfig.CSIProxyDownloadURL).To(ContainSubstring(azureChinaCloudMirror))
	g.Expect(AzureChinaCloudSpec.KubernetesSpecConfig.VnetCNIWindowsPluginsDownloadURL).To(ContainSubstring(azureChinaCloudMirror))

	// Do not check KubeBinariesSASURLBase. Please see the comments in azenvtypes.go
	g.Expect(AzureChinaCloudSpec.KubernetesSpecConfig.KubeBinariesSASURLBase).To(ContainSubstring(DefaultKubernetesSpecConfig.KubeBinariesSASURLBase))
}
