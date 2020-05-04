// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"testing"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/Azure/aks-engine/pkg/api"
)

func TestNewGenerateCmd(t *testing.T) {
	t.Parallel()

	command := newGenerateCmd()
	if command.Use != generateName || command.Short != generateShortDescription || command.Long != generateLongDescription {
		t.Fatalf("generate command should have use %s equal %s, short %s equal %s and long %s equal to %s", command.Use, generateName, command.Short, generateShortDescription, command.Long, generateLongDescription)
	}

	expectedFlags := []string{"api-model", "output-directory", "ca-certificate-path", "ca-private-key-path", "set", "no-pretty-print", "parameters-only", "client-id", "client-secret"}
	for _, f := range expectedFlags {
		if command.Flags().Lookup(f) == nil {
			t.Fatalf("generate command should have flag %s", f)
		}
	}

	command.SetArgs([]string{})
	if err := command.Execute(); err == nil {
		t.Fatalf("expected an error when calling generate with no arguments")
	}
}

func TestGenerateCmdValidate(t *testing.T) {
	t.Parallel()

	g := &generateCmd{}
	r := &cobra.Command{}

	// validate cmd with 1 arg
	err := g.validate(r, []string{"../pkg/engine/testdata/simple/kubernetes.json"})
	if err != nil {
		t.Fatalf("unexpected error validating 1 arg: %s", err.Error())
	}

	g = &generateCmd{}

	// validate cmd with 0 args
	err = g.validate(r, []string{})
	if err == nil {
		t.Fatalf("expected error validating 0 args")
	}

	g = &generateCmd{}

	// validate cmd with more than 1 arg
	err = g.validate(r, []string{"../pkg/engine/testdata/simple/kubernetes.json", "arg1"})
	if err == nil {
		t.Fatalf("expected error validating multiple args")
	}

}

func TestGenerateCmdMergeAPIModel(t *testing.T) {

	cases := []struct {
		test func(*testing.T)
		name string
	}{
		{
			name: "NoSetFlagDefined",
			test: func(t *testing.T) {
				g := new(generateCmd)
				g.apimodelPath = "../pkg/engine/testdata/simple/kubernetes.json"
				err := g.mergeAPIModel()
				if err != nil {
					t.Fatalf("unexpected error calling mergeAPIModel with no --set flag defined: %s", err.Error())
				}
			},
		},
		{
			name: "OneFlagSet",
			test: func(t *testing.T) {
				g := new(generateCmd)
				g.apimodelPath = "../pkg/engine/testdata/simple/kubernetes.json"
				g.set = []string{"masterProfile.count=3,linuxProfile.adminUsername=testuser"}
				err := g.mergeAPIModel()
				if err != nil {
					t.Fatalf("unexpected error calling mergeAPIModel with one --set flag: %s", err.Error())
				}
			},
		},
		{
			name: "TwoFlagsSet",
			test: func(t *testing.T) {
				g := new(generateCmd)
				g.apimodelPath = "../pkg/engine/testdata/simple/kubernetes.json"
				g.set = []string{"masterProfile.count=3,linuxProfile.adminUsername=testuser"}
				err := g.mergeAPIModel()
				if err != nil {
					t.Fatalf("unexpected error calling mergeAPIModel with one --set flag: %s", err.Error())
				}
			},
		},
		{
			name: "OverrideArrayFlagSet",
			test: func(t *testing.T) {
				g := new(generateCmd)
				g.apimodelPath = "../pkg/engine/testdata/simple/kubernetes.json"
				g.set = []string{"agentPoolProfiles[0].count=1"}
				err := g.mergeAPIModel()
				if err != nil {
					t.Fatalf("unexpected error calling mergeAPIModel with one --set flag to override an array property: %s", err.Error())
				}
			},
		},
		{
			name: "SshKeyContains==FlagSet",
			test: func(t *testing.T) {
				g := new(generateCmd)
				g.apimodelPath = "../pkg/engine/testdata/simple/kubernetes.json"
				g.set = []string{"linuxProfile.ssh.publicKeys[0].keyData=\"ssh-rsa AAAAB3NO8b9== azureuser@cluster.local\",servicePrincipalProfile.clientId=\"123a4321-c6eb-4b61-9d6f-7db123e14a7a\",servicePrincipalProfile.secret=\"=#msRock5!t=\""}
				err := g.mergeAPIModel()
				if err != nil {
					t.Fatalf("unexpected error calling mergeAPIModel with one --set flag to override an array property: %s", err.Error())
				}
			},
		},
		{
			name: "SimpleQuoteContainingFlagSet",
			test: func(t *testing.T) {
				g := new(generateCmd)
				g.apimodelPath = "../pkg/engine/testdata/simple/kubernetes.json"
				g.set = []string{"servicePrincipalProfile.secret='=MsR0ck5!t='"}
				err := g.mergeAPIModel()
				if err != nil {
					t.Fatalf("unexpected error calling mergeAPIModel with one --set flag to override an array property: %s", err.Error())
				}
			},
		},
	}

	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			c.test(t)
		})
	}
}

func TestGenerateCmdMLoadAPIModel(t *testing.T) {
	g := &generateCmd{}
	r := &cobra.Command{}

	g.apimodelPath = "../pkg/engine/testdata/simple/kubernetes.json"
	g.set = []string{"agentPoolProfiles[0].count=1"}

	err := g.validate(r, []string{"../pkg/engine/testdata/simple/kubernetes.json"})
	if err != nil {
		t.Fatalf("unexpected error validating api model: %s", err.Error())
	}
	err = g.mergeAPIModel()
	if err != nil {
		t.Fatalf("unexpected error merging api model: %s", err.Error())
	}
	err = g.loadAPIModel()
	if err != nil {
		t.Fatalf("unexpected error loading api model: %s", err.Error())
	}
}

func TestAPIModelWithoutServicePrincipalProfileAndClientIdAndSecretInGenerateCmd(t *testing.T) {
	t.Parallel()

	apiloader := &api.Apiloader{
		Translator: nil,
	}

	apimodel := getAPIModelWithoutServicePrincipalProfile(false)

	cs, ver, err := apiloader.DeserializeContainerService([]byte(apimodel), false, false, nil)
	if err != nil {
		t.Fatalf("unexpected error deserializing the example apimodel: %s", err)
	}
	cs.Properties.LinuxProfile.SSH.PublicKeys[0].KeyData = "ssh test"
	outfile, del := makeTmpFile(t, "_test_output")
	defer del()

	clientID, _ := uuid.Parse("e810b868-afab-412d-98cc-ce7db5cc840b")
	clientSecret := "Test Client secret"
	generateCmd := &generateCmd{
		apimodelPath:     "./this/is/unused.json",
		outputDirectory:  outfile,
		ClientID:         clientID,
		ClientSecret:     clientSecret,
		containerService: cs,
		apiVersion:       ver,
	}
	err = generateCmd.autofillApimodel()
	if err != nil {
		t.Fatalf("unexpected error autofilling the example apimodel: %s", err)
	}

	if generateCmd.containerService.Properties.ServicePrincipalProfile == nil || generateCmd.containerService.Properties.ServicePrincipalProfile.ClientID == "" || generateCmd.containerService.Properties.ServicePrincipalProfile.Secret == "" {
		t.Fatalf("expected service principal profile to be populated from deployment command arguments")
	}

	if generateCmd.containerService.Properties.ServicePrincipalProfile.ClientID != clientID.String() {
		t.Fatalf("expected service principal profile client id to be %s but got %s", clientID.String(), generateCmd.containerService.Properties.ServicePrincipalProfile.ClientID)
	}

	if generateCmd.containerService.Properties.ServicePrincipalProfile.Secret != clientSecret {
		t.Fatalf("expected service principal profile client secret to be %s but got %s", clientSecret, generateCmd.containerService.Properties.ServicePrincipalProfile.Secret)
	}

	err = generateCmd.validateAPIModelAsVLabs()
	if err != nil {
		t.Fatalf("unexpected error validateAPIModelAsVLabs the example apimodel: %s", err)
	}
}

func TestAPIModelWithoutServicePrincipalProfileAndWithoutClientIdAndSecretInGenerateCmd(t *testing.T) {
	t.Parallel()

	apiloader := &api.Apiloader{
		Translator: nil,
	}

	apimodel := getAPIModelWithoutServicePrincipalProfile(false)
	cs, ver, err := apiloader.DeserializeContainerService([]byte(apimodel), false, false, nil)
	if err != nil {
		t.Fatalf("unexpected error deserializing the example apimodel: %s", err)
	}
	cs.Properties.LinuxProfile.SSH.PublicKeys[0].KeyData = "ssh test"
	outfile, del := makeTmpFile(t, "_test_output")
	defer del()

	generateCmd := &generateCmd{
		apimodelPath:     "./this/is/unused.json",
		outputDirectory:  outfile,
		containerService: cs,
		apiVersion:       ver,
	}
	err = generateCmd.autofillApimodel()
	if err != nil {
		t.Fatalf("unexpected error autofilling the example apimodel: %s", err)
	}

	if generateCmd.containerService.Properties.ServicePrincipalProfile != nil {
		t.Fatalf("expected service principal profile to be nil for unmanaged identity, where client id and secret are not supplied in api model and deployment command")
	}

	err = generateCmd.validateAPIModelAsVLabs()
	expectedErr := errors.New("ServicePrincipalProfile must be specified with Orchestrator Kubernetes")

	if err != nil && err.Error() != expectedErr.Error() {
		t.Fatalf("expected validate generate command to return error %s, but instead got %s", expectedErr.Error(), err.Error())
	}
}

func TestAPIModelWithManagedIdentityWithoutServicePrincipalProfileAndClientIdAndSecretInGenerateCmd(t *testing.T) {
	t.Parallel()

	apiloader := &api.Apiloader{
		Translator: nil,
	}

	apimodel := getAPIModelWithoutServicePrincipalProfile(true)

	cs, ver, err := apiloader.DeserializeContainerService([]byte(apimodel), false, false, nil)
	if err != nil {
		t.Fatalf("unexpected error deserializing the example apimodel: %s", err)
	}
	cs.Properties.LinuxProfile.SSH.PublicKeys[0].KeyData = "ssh test"
	clientID, _ := uuid.Parse("e810b868-afab-412d-98cc-ce7db5cc840b")
	clientSecret := "Test Client secret"
	outfile, del := makeTmpFile(t, "_test_output")
	defer del()

	generateCmd := &generateCmd{
		apimodelPath:     "./this/is/unused.json",
		outputDirectory:  outfile,
		ClientID:         clientID,
		ClientSecret:     clientSecret,
		containerService: cs,
		apiVersion:       ver,
	}
	err = generateCmd.autofillApimodel()
	if err != nil {
		t.Fatalf("unexpected error autofilling the example apimodel: %s", err)
	}

	if generateCmd.containerService.Properties.ServicePrincipalProfile != nil {
		t.Fatalf("expected service principal profile to be nil for managed identity")
	}

	err = generateCmd.validateAPIModelAsVLabs()
	if err != nil {
		t.Fatalf("unexpected error validateAPIModelAsVLabs the example apimodel: %s", err)
	}
}

func TestExampleAPIModels(t *testing.T) {
	defaultSet := []string{"masterProfile.dnsPrefix=my-cluster,linuxProfile.ssh.publicKeys[0].keyData=\"ssh-rsa AAAAB3NO8b9== azureuser@cluster.local\",servicePrincipalProfile.clientId=\"123a4321-c6eb-4b61-9d6f-7db123e14a7a\",servicePrincipalProfile.secret=\"=#msRock5!t=\""}
	tests := []struct {
		name         string
		apiModelPath string
		setArgs      []string
	}{
		{
			name:         "default cluster configuration",
			apiModelPath: "../examples/kubernetes.json",
			setArgs:      defaultSet,
		},
		{
			name:         "AAD pod identity",
			apiModelPath: "../examples/addons/aad-pod-identity/kubernetes-aad-pod-identity.json",
			setArgs:      defaultSet,
		},
		{
			name:         "ACI connector",
			apiModelPath: "../examples/addons/aci-connector/kubernetes-aci-connector.json",
			setArgs:      defaultSet,
		},
		{
			name:         "App gateway ingress",
			apiModelPath: "../examples/addons/appgw-ingress/kubernetes-appgw-ingress.json",
			setArgs:      defaultSet,
		},
		{
			name:         "cluster-autoscaler",
			apiModelPath: "../examples/addons/cluster-autoscaler/kubernetes-cluster-autoscaler.json",
			setArgs:      defaultSet,
		},
		{
			name:         "container-monitoring",
			apiModelPath: "../examples/addons/container-monitoring/kubernetes-container-monitoring.json",
			setArgs:      defaultSet,
		},
		{
			name:         "custom PSP",
			apiModelPath: "../examples/addons/custom-manifests/kubernetes-custom-psp.json",
			setArgs:      defaultSet,
		},
		{
			name:         "keyvault flexvol",
			apiModelPath: "../examples/addons/keyvault-flexvolume/kubernetes-keyvault-flexvolume.json",
			setArgs:      defaultSet,
		},
		{
			name:         "nvidia",
			apiModelPath: "../examples/addons/nvidia-device-plugin/nvidia-device-plugin.json",
			setArgs:      defaultSet,
		},
		{
			name:         "azure-policy",
			apiModelPath: "../examples/addons/azure-policy/azure-policy.json",
			setArgs:      defaultSet,
		},
		{
			name:         "node problem detector",
			apiModelPath: "../examples/addons/node-problem-detector/node-problem-detector.json",
			setArgs:      defaultSet,
		},
		{
			name:         "cosmos etcd",
			apiModelPath: "../examples/cosmos-etcd/kubernetes-3-masters-cosmos.json",
			setArgs:      defaultSet,
		},
		{
			name:         "custom files pod node selector",
			apiModelPath: "../examples/customfiles/kubernetes-customfiles-podnodeselector.json",
			setArgs:      []string{"aadProfile.clientAppID=e810b868-afab-412d-98cc-ce7db5cc840b,aadProfile.serverAppID=f810b868-afab-412d-98cc-ce7db5cc840b,masterProfile.dnsPrefix=my-cluster,linuxProfile.ssh.publicKeys[0].keyData=\"ssh-rsa AAAAB3NO8b9== azureuser@cluster.local\",servicePrincipalProfile.clientId=\"123a4321-c6eb-4b61-9d6f-7db123e14a7a\",servicePrincipalProfile.secret=\"=#msRock5!t=\""},
		},
		{
			name:         "ephemeral disks Standard_D2s_v3",
			apiModelPath: "../examples/disks-ephemeral/ephemeral-disks.json",
			setArgs:      defaultSet,
		},
		{
			name:         "ephemeral disks Standard_D2_v2",
			apiModelPath: "../examples/disks-ephemeral/kubernetes-vmas.json",
			setArgs:      defaultSet,
		},
		{
			name:         "ephemeral disks Standard_D2_v2",
			apiModelPath: "../examples/disks-ephemeral/kubernetes-vmas.json",
			setArgs:      defaultSet,
		},
		{
			name:         "managed disk pre-attached",
			apiModelPath: "../examples/disks-managed/kubernetes-preAttachedDisks-vmas.json",
			setArgs:      defaultSet,
		},
		{
			name:         "managed disk",
			apiModelPath: "../examples/disks-managed/kubernetes-vmas.json",
			setArgs:      defaultSet,
		},
		{
			name:         "storage account on master",
			apiModelPath: "../examples/disks-storageaccount/kubernetes-master-sa.json",
			setArgs:      defaultSet,
		},
		{
			name:         "storage account",
			apiModelPath: "../examples/disks-storageaccount/kubernetes.json",
			setArgs:      defaultSet,
		},
		{
			name:         "dualstack ipv6",
			apiModelPath: "../examples/dualstack/kubernetes.json",
			setArgs:      defaultSet,
		},
		{
			name:         "extensions",
			apiModelPath: "../examples/extensions/kubernetes.json",
			setArgs:      defaultSet,
		},
		{
			name:         "extensions oms",
			apiModelPath: "../examples/extensions/kubernetes.oms.json",
			setArgs:      defaultSet,
		},
		{
			name:         "extensions preprovision",
			apiModelPath: "../examples/extensions/kubernetes.preprovision.json",
			setArgs:      defaultSet,
		},
		{
			name:         "extensions prometheus grafana",
			apiModelPath: "../examples/extensions/prometheus-grafana-k8s.json",
			setArgs:      defaultSet,
		},
		{
			name:         "feature gates",
			apiModelPath: "../examples/feature-gates/kubernetes-featuresgates.json",
			setArgs:      defaultSet,
		},
		{
			name:         "ipvs",
			apiModelPath: "../examples/ipvs/kubernetes-msi.json",
			setArgs:      defaultSet,
		},
		{
			name:         "keyvault params",
			apiModelPath: "../examples/keyvault-params/kubernetes.json",
			setArgs:      []string{"aadProfile.clientAppID=e810b868-afab-412d-98cc-ce7db5cc840b,aadProfile.serverAppID=f810b868-afab-412d-98cc-ce7db5cc840b,masterProfile.dnsPrefix=my-cluster,linuxProfile.ssh.publicKeys[0].keyData=\"ssh-rsa AAAAB3NO8b9== azureuser@cluster.local\",servicePrincipalProfile.clientId=\"123a4321-c6eb-4b61-9d6f-7db123e14a7a\""},
		},
		{
			name:         "keyvault certs",
			apiModelPath: "../examples/keyvaultcerts/kubernetes.json",
			setArgs:      []string{"linuxProfile.secrets[0].sourceVault.id=my-id,masterProfile.dnsPrefix=my-cluster,linuxProfile.ssh.publicKeys[0].keyData=\"ssh-rsa AAAAB3NO8b9== azureuser@cluster.local\",servicePrincipalProfile.clientId=\"123a4321-c6eb-4b61-9d6f-7db123e14a7a\",servicePrincipalProfile.secret=\"=#msRock5!t=\""},
		},
		{
			name:         "accelerated networking",
			apiModelPath: "../examples/kubernetes-config/kubernetes-accelerated-network.json",
			setArgs:      defaultSet,
		},
		{
			name:         "useCloudControllerManager",
			apiModelPath: "../examples/kubernetes-config/kubernetes-cloud-controller-manager.json",
			setArgs:      defaultSet,
		},
		{
			name:         "clusterSubnet",
			apiModelPath: "../examples/kubernetes-config/kubernetes-clustersubnet.json",
			setArgs:      defaultSet,
		},
		{
			name:         "enableDataEncryptionAtRest",
			apiModelPath: "../examples/kubernetes-config/kubernetes-data-encryption-at-rest.json",
			setArgs:      defaultSet,
		},
		{
			name:         "dockerBridgeSubnet",
			apiModelPath: "../examples/kubernetes-config/kubernetes-dockerbridgesubnet.json",
			setArgs:      defaultSet,
		},
		{
			name:         "etcdDiskSizeGB",
			apiModelPath: "../examples/kubernetes-config/kubernetes-etcd-storage-size.json",
			setArgs:      defaultSet,
		},
		{
			name:         "gc thresholds",
			apiModelPath: "../examples/kubernetes-config/kubernetes-gc.json",
			setArgs:      defaultSet,
		},
		{
			name:         "enableEncryptionWithExternalKms",
			apiModelPath: "../examples/kubernetes-config/kubernetes-keyvault-encryption.json",
			setArgs:      []string{"masterProfile.dnsPrefix=my-cluster,linuxProfile.ssh.publicKeys[0].keyData=\"ssh-rsa AAAAB3NO8b9== azureuser@cluster.local\",servicePrincipalProfile.clientId=\"123a4321-c6eb-4b61-9d6f-7db123e14a7a\",servicePrincipalProfile.objectId=\"223a4321-c6eb-4b61-9d6f-7db123e14a7a\",servicePrincipalProfile.secret=\"=#msRock5!t=\""},
		},
		{
			name:         "max pods",
			apiModelPath: "../examples/kubernetes-config/kubernetes-maxpods.json",
			setArgs:      defaultSet,
		},
		{
			name:         "dashboard disabled",
			apiModelPath: "../examples/kubernetes-config/kubernetes-no-dashboard.json",
			setArgs:      defaultSet,
		},
		{
			name:         "private cluster single master",
			apiModelPath: "../examples/kubernetes-config/kubernetes-private-cluster-single-master.json",
			setArgs:      []string{"orchestratorProfile.kubernetesConfig.privateCluster.jumpboxProfile.publicKey=\"ssh-rsa AAAAB3NO8b9== azureuser@cluster.local\",masterProfile.dnsPrefix=my-cluster,linuxProfile.ssh.publicKeys[0].keyData=\"ssh-rsa AAAAB3NO8b9== azureuser@cluster.local\",servicePrincipalProfile.clientId=\"123a4321-c6eb-4b61-9d6f-7db123e14a7a\",servicePrincipalProfile.secret=\"=#msRock5!t=\""},
		},
		{
			name:         "private cluster",
			apiModelPath: "../examples/kubernetes-config/kubernetes-private-cluster.json",
			setArgs:      []string{"orchestratorProfile.kubernetesConfig.privateCluster.jumpboxProfile.publicKey=\"ssh-rsa AAAAB3NO8b9== azureuser@cluster.local\",masterProfile.dnsPrefix=my-cluster,linuxProfile.ssh.publicKeys[0].keyData=\"ssh-rsa AAAAB3NO8b9== azureuser@cluster.local\",servicePrincipalProfile.clientId=\"123a4321-c6eb-4b61-9d6f-7db123e14a7a\",servicePrincipalProfile.secret=\"=#msRock5!t=\""},
		},
		{
			name:         "rescheduler addon",
			apiModelPath: "../examples/kubernetes-config/kubernetes-rescheduler.json",
			setArgs:      defaultSet,
		},
		{
			name:         "standard LB",
			apiModelPath: "../examples/kubernetes-config/kubernetes-standardlb.json",
			setArgs:      defaultSet,
		},
		{
			name:         "gpu",
			apiModelPath: "../examples/kubernetes-gpu/kubernetes.json",
			setArgs:      defaultSet,
		},
		{
			name:         "labels",
			apiModelPath: "../examples/kubernetes-labels/kubernetes.json",
			setArgs:      defaultSet,
		},
		{
			name:         "msi user-assigned vmas",
			apiModelPath: "../examples/kubernetes-msi-userassigned/kube-vma.json",
			setArgs:      defaultSet,
		},
		{
			name:         "msi user-assigned vmss",
			apiModelPath: "../examples/kubernetes-msi-userassigned/kube-vmss.json",
			setArgs:      defaultSet,
		},
		{
			name:         "1.14 example",
			apiModelPath: "../examples/kubernetes-releases/kubernetes1.14.json",
			setArgs:      defaultSet,
		},
		{
			name:         "1.15 example",
			apiModelPath: "../examples/kubernetes-releases/kubernetes1.15.json",
			setArgs:      defaultSet,
		},
		{
			name:         "1.16 example",
			apiModelPath: "../examples/kubernetes-releases/kubernetes1.16.json",
			setArgs:      defaultSet,
		},
		{
			name:         "1.17 example",
			apiModelPath: "../examples/kubernetes-releases/kubernetes1.17.json",
			setArgs:      defaultSet,
		},
		{
			name:         "1.18 example",
			apiModelPath: "../examples/kubernetes-releases/kubernetes1.18.json",
			setArgs:      defaultSet,
		},
		{
			name:         "1.19 example",
			apiModelPath: "../examples/kubernetes-releases/kubernetes1.19.json",
			setArgs:      defaultSet,
		},
		{
			name:         "vmss",
			apiModelPath: "../examples/kubernetes-vmss/kubernetes.json",
			setArgs:      defaultSet,
		},
		{
			name:         "vmss spot",
			apiModelPath: "../examples/kubernetes-vmss-spot/kubernetes.json",
			setArgs:      defaultSet,
		},
		{
			name:         "vmss master",
			apiModelPath: "../examples/kubernetes-vmss-master/kubernetes.json",
			setArgs:      defaultSet,
		},
		{
			name:         "vmss master custom vnet",
			apiModelPath: "../examples/kubernetes-vmss-master/customvnet.json",
			setArgs:      defaultSet,
		},
		{
			name:         "vmss master windows",
			apiModelPath: "../examples/kubernetes-vmss-master/windows.json",
			setArgs:      defaultSet,
		},
		{
			name:         "large cluster",
			apiModelPath: "../examples/largeclusters/kubernetes.json",
			setArgs:      defaultSet,
		},
		{
			name:         "msi",
			apiModelPath: "../examples/managed-identity/kubernetes-msi.json",
			setArgs:      defaultSet,
		},
		{
			name:         "multiple masters - 3",
			apiModelPath: "../examples/multiple-masters/kubernetes-3-masters.json",
			setArgs:      defaultSet,
		},
		{
			name:         "multiple masters - 5",
			apiModelPath: "../examples/multiple-masters/kubernetes-5-masters.json",
			setArgs:      defaultSet,
		},
		{
			name:         "multiple node pools",
			apiModelPath: "../examples/multiple-nodepools/multipool.json",
			setArgs:      defaultSet,
		},
		{
			name:         "Azure CNI",
			apiModelPath: "../examples/networkplugin/kubernetes-azure.json",
			setArgs:      defaultSet,
		},
		{
			name:         "Azure CNI with calico",
			apiModelPath: "../examples/networkpolicy/kubernetes-calico-azure.json",
			setArgs:      defaultSet,
		},
		{
			name:         "kubenet with calico",
			apiModelPath: "../examples/networkpolicy/kubernetes-calico-kubenet.json",
			setArgs:      defaultSet,
		},
		{
			name:         "cilium network policy",
			apiModelPath: "../examples/networkpolicy/kubernetes-cilium.json",
			setArgs:      defaultSet,
		},
		{
			name:         "antrea network policy",
			apiModelPath: "../examples/networkpolicy/kubernetes-antrea.json",
			setArgs:      defaultSet,
		},
		{
			name:         "istio",
			apiModelPath: "../examples/service-mesh/istio.json",
			setArgs:      defaultSet,
		},
		{
			name:         "ubuntu 16.04",
			apiModelPath: "../examples/ubuntu-1604/kubernetes.json",
			setArgs:      defaultSet,
		},
		{
			name:         "ubuntu 18.04",
			apiModelPath: "../examples/ubuntu-1804/kubernetes.json",
			setArgs:      defaultSet,
		},
		{
			name:         "kubenet custom vnet",
			apiModelPath: "../examples/vnet/kubernetesvnet.json",
			setArgs:      defaultSet,
		},
		{
			name:         "Azure CNI custom vnet",
			apiModelPath: "../examples/vnet/kubernetesvnet-azure-cni.json",
			setArgs:      defaultSet,
		},
		{
			name:         "master vmss custom vnet",
			apiModelPath: "../examples/vnet/kubernetes-master-vmss.json",
			setArgs:      defaultSet,
		},
		{
			name:         "custom node DNS custom vnet",
			apiModelPath: "../examples/vnet/kubernetesvnet-customnodesdns.json",
			setArgs:      defaultSet,
		},
		{
			name:         "custom search domain custom vnet",
			apiModelPath: "../examples/vnet/kubernetesvnet-customsearchdomain.json",
			setArgs:      defaultSet,
		},
		{
			name:         "windows custom image",
			apiModelPath: "../examples/windows/kubernetes-custom-image.json",
			setArgs:      defaultSet,
		},
		{
			name:         "windows Standard_D2",
			apiModelPath: "../examples/windows/kubernetes-D2.json",
			setArgs:      defaultSet,
		},
		{
			name:         "windows hybrid",
			apiModelPath: "../examples/windows/kubernetes-hybrid.json",
			setArgs:      defaultSet,
		},
		{
			name:         "windows hyperv",
			apiModelPath: "../examples/windows/kubernetes-hyperv.json",
			setArgs:      defaultSet,
		},
		{
			name:         "windows managed disk",
			apiModelPath: "../examples/windows/kubernetes-manageddisks.json",
			setArgs:      defaultSet,
		},
		{
			name:         "windows master uses storage account",
			apiModelPath: "../examples/windows/kubernetes-master-sa.json",
			setArgs:      defaultSet,
		},
		{
			name:         "windows master uses storage account",
			apiModelPath: "../examples/windows/kubernetes-master-sa.json",
			setArgs:      defaultSet,
		},
		{
			name:         "windows storage account",
			apiModelPath: "../examples/windows/kubernetes-sadisks.json",
			setArgs:      defaultSet,
		},
		{
			name:         "windows winrm extension",
			apiModelPath: "../examples/windows/kubernetes-wincni.json",
			setArgs:      defaultSet,
		},
		{
			name:         "windows 1903",
			apiModelPath: "../examples/windows/kubernetes-windows-1903.json",
			setArgs:      defaultSet,
		},
		{
			name:         "windows automatic updates",
			apiModelPath: "../examples/windows/kubernetes-windows-automatic-update.json",
			setArgs:      defaultSet,
		},
		{
			name:         "windows windowsDockerVersion",
			apiModelPath: "../examples/windows/kubernetes-windows-docker-version.json",
			setArgs:      defaultSet,
		},
		{
			name:         "windows version",
			apiModelPath: "../examples/windows/kubernetes-windows-version.json",
			setArgs:      defaultSet,
		},
		{
			name:         "windows recommended kubernetes config",
			apiModelPath: "../examples/windows/kubernetes.json",
			setArgs:      defaultSet,
		},
		{
			name:         "custom image",
			apiModelPath: "../examples/custom-image.json",
			setArgs:      defaultSet,
		},
		{
			name:         "custom shared image",
			apiModelPath: "../examples/custom-shared-image.json",
			setArgs:      defaultSet,
		},
		{
			name:         "containerd",
			apiModelPath: "../examples/kubernetes-containerd.json",
			setArgs:      defaultSet,
		},
		{
			name:         "Standard_D2",
			apiModelPath: "../examples/kubernetes-D2.json",
			setArgs:      defaultSet,
		},
		{
			name:         "ubuntu distros",
			apiModelPath: "../examples/kubernetes-non-vhd-distros.json",
			setArgs:      defaultSet,
		},
		{
			name:         "docker tmp dir",
			apiModelPath: "../examples/kubernetes-config/kubernetes-docker-tmpdir.json",
			setArgs:      defaultSet,
		},
		{
			name:         "containerd tmp dir",
			apiModelPath: "../examples/kubernetes-config/kubernetes-containerd-tmpdir.json",
			setArgs:      defaultSet,
		},
		{
			name:         "e2e gpu",
			apiModelPath: "../examples/e2e-tests/kubernetes/gpu-enabled/definition.json",
			setArgs:      defaultSet,
		},
		{
			name:         "e2e addons disabled",
			apiModelPath: "../examples/e2e-tests/kubernetes/kubernetes-config/addons-disabled.json",
			setArgs:      defaultSet,
		},
		{
			name:         "e2e addons enabled",
			apiModelPath: "../examples/e2e-tests/kubernetes/kubernetes-config/addons-enabled.json",
			setArgs:      defaultSet,
		},
		{
			name:         "e2e kubenet",
			apiModelPath: "../examples/e2e-tests/kubernetes/kubernetes-config/network-plugin-kubenet.json",
			setArgs:      defaultSet,
		},
		{
			name:         "e2e rbac disabled",
			apiModelPath: "../examples/e2e-tests/kubernetes/kubernetes-config/rbac-disabled.json",
			setArgs:      defaultSet,
		},
		{
			name:         "e2e 50 nodes",
			apiModelPath: "../examples/e2e-tests/kubernetes/node-count/50-nodes/definition.json",
			setArgs:      defaultSet,
		},
		{
			name:         "e2e full configuration",
			apiModelPath: "../examples/e2e-tests/kubernetes/release/default/definition.json",
			setArgs:      defaultSet,
		},
		{
			name:         "e2e windows",
			apiModelPath: "../examples/e2e-tests/kubernetes/windows/definition.json",
			setArgs:      defaultSet,
		},
		{
			name:         "e2e hybrid",
			apiModelPath: "../examples/e2e-tests/kubernetes/windows/hybrid/definition.json",
			setArgs:      defaultSet,
		},
		{
			name:         "e2e zones",
			apiModelPath: "../examples/e2e-tests/kubernetes/zones/definition.json",
			setArgs:      defaultSet,
		},
		{
			name:         "e2e user-assigned identity vmas",
			apiModelPath: "../examples/e2e-tests/userassignedidentity/vmas/kubernetes-vmas.json",
			setArgs:      defaultSet,
		},
		{
			name:         "e2e user-assigned identity vmas multi-master",
			apiModelPath: "../examples/e2e-tests/userassignedidentity/vmas/kubernetes-vmas-multimaster.json",
			setArgs:      defaultSet,
		},
		{
			name:         "e2e user-assigned identity vmss",
			apiModelPath: "../examples/e2e-tests/userassignedidentity/vmss/kubernetes-vmss.json",
			setArgs:      defaultSet,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			dir, del := makeTmpDir(t)
			defer del()
			g := &generateCmd{
				apimodelPath:    test.apiModelPath,
				outputDirectory: dir,
			}
			g.set = test.setArgs
			r := &cobra.Command{}

			if err := g.validate(r, []string{}); err != nil {
				t.Fatalf("unexpected error validating api model: %s", err.Error())
			}
			if err := g.mergeAPIModel(); err != nil {
				t.Fatalf("unexpected error merging api model: %s", err.Error())
			}
			if err := g.loadAPIModel(); err != nil {
				t.Fatalf("unexpected error loading api model: %s", err.Error())
			}

			if err := g.validateAPIModelAsVLabs(); err != nil {
				t.Fatalf("unexpected error validateAPIModelAsVLabs the example apimodel: %s", err)
			}
		})
	}
}
