// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/google/uuid"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/helpers"
)

func mockExecuteCmd(command, masterFQDN, hostname string, port string, config *ssh.ClientConfig) (string, error) {
	if masterFQDN != "valid" {
		return "error running command", errors.New("executeCmd failed")
	}
	return "success", nil
}

func TestNewRotateCertsCmd(t *testing.T) {
	t.Parallel()

	output := newRotateCertsCmd()
	if output.Use != rotateCertsName || output.Short != rotateCertsShortDescription || output.Long != rotateCertsLongDescription {
		t.Fatalf("rotate-certs command should have use %s equal %s, short %s equal %s and long %s equal to %s", output.Use, rotateCertsName, output.Short, rotateCertsShortDescription, output.Long, rotateCertsLongDescription)
	}

	expectedFlags := []string{"location", "resource-group", "apiserver", "api-model", "ssh"}
	for _, f := range expectedFlags {
		if output.Flags().Lookup(f) == nil {
			t.Fatalf("rotate-certs command should have flag %s", f)
		}
	}
}

func TestRotateCertsCmdRun(t *testing.T) {
	t.Parallel()

	tmpSSHFile, del := makeTmpFile(t, "_test_ssh")
	defer del()

	tmpOutputDir, del := makeTmpDir(t)
	defer del()

	rcc := &rotateCertsCmd{
		client: &armhelpers.MockAKSEngineClient{},
		authProvider: &mockAuthProvider{
			authArgs:      &authArgs{},
			getClientMock: &armhelpers.MockAKSEngineClient{},
		},
		apiModelPath:       "../pkg/engine/testdata/key-vault-certs/kubernetes.json",
		outputDirectory:    tmpOutputDir,
		location:           "westus",
		sshFilepath:        tmpSSHFile,
		sshCommandExecuter: mockExecuteCmd,
		masterFQDN:         "valid",
	}

	r := &cobra.Command{}
	f := r.Flags()
	addAuthFlags(rcc.getAuthArgs(), f)
	fakeRawSubscriptionID := "6dc93fae-9a76-421f-bbe5-cc6460ea81cb"
	fakeSubscriptionID, err := uuid.Parse(fakeRawSubscriptionID)
	fakeClientID := "b829b379-ca1f-4f1d-91a2-0d26b244680d"
	fakeClientSecret := "0se43bie-3zs5-303e-aav5-dcf231vb82ds"
	if err != nil {
		t.Fatalf("Invalid SubscriptionId in Test: %s", err)
	}

	rcc.getAuthArgs().SubscriptionID = fakeSubscriptionID
	rcc.getAuthArgs().rawSubscriptionID = fakeRawSubscriptionID
	rcc.getAuthArgs().rawClientID = fakeClientID
	rcc.getAuthArgs().ClientSecret = fakeClientSecret
	err = rcc.run(r, []string{})
	if err != nil {
		t.Fatalf("Failed to run rotate-certs command: %s", err)
	}
}

func TestGetClusterNodes(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	mockClient := &armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
	mockClient.MockKubernetesClient.FailListNodes = true
	rcc := rotateCertsCmd{
		authProvider:     &authArgs{},
		client:           mockClient,
		containerService: api.CreateMockContainerService("testcluster", "1.10.13", 3, 2, false),
	}
	err := rcc.getClusterNodes()
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("failed to get cluster nodes"))

	mockClient.MockKubernetesClient.FailListNodes = false
	err = rcc.getClusterNodes()
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(len(rcc.masterNodes)).To(Equal(1))
	g.Expect(len(rcc.agentNodes)).To(Equal(1))
}

func TestDeleteAllPods(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	mockClient := &armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
	mockClient.MockKubernetesClient.FailListPods = true
	rcc := rotateCertsCmd{
		authProvider:     &authArgs{},
		client:           mockClient,
		containerService: api.CreateMockContainerService("testcluster", "1.10.13", 3, 2, false),
	}
	err := rcc.deleteAllPods()
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("failed to get pods"))

	mockClient.MockKubernetesClient.FailListPods = false
	mockClient.MockKubernetesClient.FailDeletePod = true
	mockClient.MockKubernetesClient.PodsList = &v1.PodList{
		Items: []v1.Pod{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      common.KubeDNSAddonName,
					Namespace: "kube-system",
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-pod",
					Namespace: "kube-system",
				},
			},
		},
	}
	err = rcc.deleteAllPods()
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("failed to delete pod"))

	mockClient.MockKubernetesClient.FailDeletePod = false
	err = rcc.deleteAllPods()
	g.Expect(err).NotTo(HaveOccurred())
}

func TestRebootAllNodes(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	g := NewGomegaWithT(t)
	mockClient := &armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
	mockClient.FailListVirtualMachines = true
	mockClient.FailListVirtualMachineScaleSets = false
	mockClient.FakeListVirtualMachineScaleSetsResult = func() []compute.VirtualMachineScaleSet {
		scalesetName := "scalesetName"
		sku := compute.Sku{}
		location := "eastus"
		return []compute.VirtualMachineScaleSet{
			{
				Name:     &scalesetName,
				Sku:      &sku,
				Location: &location,
			},
		}
	}
	rcc := rotateCertsCmd{
		authProvider:      &authArgs{},
		client:            mockClient,
		containerService:  api.CreateMockContainerService("testcluster", "1.10.13", 3, 2, false),
		resourceGroupName: "test-rg",
	}
	err := rcc.rebootAllNodes(ctx)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("failed to list Virtual Machines in resource group test-rg"))

	mockClient.FailListVirtualMachines = false
	mockClient.FailListVirtualMachineScaleSets = true
	err = rcc.rebootAllNodes(ctx)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("failed to list Virtual Machine Scale Sets in resource group test-rg"))

	mockClient.FailListVirtualMachines = false
	mockClient.FailListVirtualMachineScaleSets = false
	mockClient.FailRestartVirtualMachine = true
	mockClient.FailRestartVirtualMachineScaleSets = false
	err = rcc.rebootAllNodes(ctx)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("failed to restart Virtual Machine"))

	mockClient.FailRestartVirtualMachine = false
	mockClient.FailRestartVirtualMachineScaleSets = true
	err = rcc.rebootAllNodes(ctx)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("failed to restart Virtual Machine Scale Sets"))

	mockClient.FailRestartVirtualMachine = false
	mockClient.FailRestartVirtualMachineScaleSets = false
	err = rcc.rebootAllNodes(ctx)
	g.Expect(err).NotTo(HaveOccurred())
}

func TestDeleteServiceAccounts(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	mockClient := &armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
	mockClient.MockKubernetesClient.FailListServiceAccounts = true
	mockClient.MockKubernetesClient.ServiceAccountList = &v1.ServiceAccountList{
		Items: []v1.ServiceAccount{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      common.KubeDNSAddonName,
					Namespace: "kube-system",
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-sa",
					Namespace: "kube-system",
				},
			},
		},
	}
	rcc := rotateCertsCmd{
		authProvider:     &authArgs{},
		client:           mockClient,
		containerService: api.CreateMockContainerService("testcluster", "1.10.13", 3, 2, false),
	}
	err := rcc.deleteServiceAccounts()
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("failed to get cluster service accounts in namespace kube-system"))

	mockClient.MockKubernetesClient.FailListServiceAccounts = false
	mockClient.MockKubernetesClient.FailDeleteServiceAccount = true
	err = rcc.deleteServiceAccounts()
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("failed to delete service account kube-dns"))

	mockClient.MockKubernetesClient.FailDeleteServiceAccount = false
	err = rcc.deleteServiceAccounts()
	g.Expect(err).NotTo(HaveOccurred())
}

func TestWriteArtifacts(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	cs := api.CreateMockContainerService("testcluster", "1.13.12", 3, 2, false)
	_, err := cs.SetPropertiesDefaults(api.PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	g.Expect(err).NotTo(HaveOccurred())
	outdir, del := makeTmpDir(t)
	defer del()

	rcc := rotateCertsCmd{
		authProvider:     &authArgs{},
		containerService: cs,
		apiVersion:       "vlabs",
		outputDirectory:  outdir,
	}

	err = rcc.writeArtifacts()
	g.Expect(err).NotTo(HaveOccurred())
}

func TestUpdateKubeconfig(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	cs := api.CreateMockContainerService("testcluster", "1.10.13", 3, 2, false)
	_, err := cs.SetPropertiesDefaults(api.PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	g.Expect(err).NotTo(HaveOccurred())

	rcc := rotateCertsCmd{
		authProvider:       &authArgs{},
		containerService:   cs,
		apiVersion:         "vlabs",
		sshCommandExecuter: mockExecuteCmd,
		masterFQDN:         "valid",
		masterNodes: []v1.Node{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: "k8s-master-1234-0",
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: "k8s-master-1234-2",
				},
			},
		},
	}
	err = rcc.updateKubeconfig()
	g.Expect(err).NotTo(HaveOccurred())

	rcc.masterFQDN = "invalid"
	err = rcc.updateKubeconfig()
	g.Expect(err).To(HaveOccurred())
}

func TestRotateCerts(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	g := NewGomegaWithT(t)
	cs := api.CreateMockContainerService("testcluster", "1.10.13", 3, 2, false)
	_, err := cs.SetPropertiesDefaults(api.PropertiesDefaultsParams{
		IsScale:    false,
		IsUpgrade:  false,
		PkiKeySize: helpers.DefaultPkiKeySize,
	})
	g.Expect(err).NotTo(HaveOccurred())

	mockClient := &armhelpers.MockAKSEngineClient{MockKubernetesClient: &armhelpers.MockKubernetesClient{}}
	rcc := rotateCertsCmd{
		authProvider:       &authArgs{},
		containerService:   cs,
		sshCommandExecuter: mockExecuteCmd,
		masterFQDN:         "valid",
		client:             mockClient,
		masterNodes: []v1.Node{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: "k8s-master-1234-0",
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: "k8s-master-1234-1",
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: "k8s-master-1234-2",
				},
			},
		},
		agentNodes: []v1.Node{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: "k8s-agents-1234-0",
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: "k8s-agents-1234-1",
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: "k8s-agents-1234-2",
				},
			},
		},
	}

	err = rcc.rotateEtcd(ctx)
	g.Expect(err).NotTo(HaveOccurred())

	err = rcc.rotateApiserver()
	g.Expect(err).NotTo(HaveOccurred())

	err = rcc.rotateKubelet()
	g.Expect(err).NotTo(HaveOccurred())

	rcc.masterFQDN = "invalid"
	err = rcc.rotateEtcd(ctx)
	g.Expect(err).To(HaveOccurred())

	err = rcc.rotateApiserver()
	g.Expect(err).To(HaveOccurred())

	err = rcc.rotateKubelet()
	g.Expect(err).To(HaveOccurred())
}

func makeTmpFile(t *testing.T, name string) (string, func()) {
	tmpF, err := ioutil.TempFile(os.TempDir(), name)
	if err != nil {
		t.Fatalf("unable to create file: %s", err.Error())
	}

	return tmpF.Name(), func() {
		defer os.Remove(tmpF.Name())
	}
}

func makeTmpDir(t *testing.T) (string, func()) {
	tmpDir, err := ioutil.TempDir(os.TempDir(), "_tmp_dir")
	if err != nil {
		t.Fatalf("unable to create dir: %s", err.Error())
	}
	return tmpDir, func() {
		defer os.RemoveAll(tmpDir)
	}
}
