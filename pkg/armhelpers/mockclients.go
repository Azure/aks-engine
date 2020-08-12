// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-12-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/azure-sdk-for-go/services/preview/msi/mgmt/2015-08-31-preview/msi"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-06-01/subscriptions"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	azStorage "github.com/Azure/azure-sdk-for-go/storage"
	"github.com/Azure/go-autorest/autorest"
	log "github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
)

const (
	defaultK8sVersionForFakeVMs = "Kubernetes:1.17.5"
	//DefaultFakeVMName is the default name assigned to VMs part of FakeListVirtualMachineScaleSetVMsResult and FakeListVirtualMachineResult
	DefaultFakeVMName = "k8s-agentpool1-12345678-0"
)

//MockAKSEngineClient is an implementation of AKSEngineClient where all requests error out
type MockAKSEngineClient struct {
	FailDeployTemplate                      bool
	FailDeployTemplateQuota                 bool
	FailDeployTemplateConflict              bool
	FailDeployTemplateWithProperties        bool
	FailEnsureResourceGroup                 bool
	FailListVirtualMachines                 bool
	FailListVirtualMachinesTags             bool
	FailListVirtualMachineScaleSets         bool
	FailRestartVirtualMachineScaleSets      bool
	FailGetVirtualMachine                   bool
	FailRestartVirtualMachine               bool
	FailDeleteVirtualMachine                bool
	FailDeleteVirtualMachineScaleSetVM      bool
	FailSetVirtualMachineScaleSetCapacity   bool
	FailListVirtualMachineScaleSetVMs       bool
	FailGetStorageClient                    bool
	FailDeleteNetworkInterface              bool
	FailGetKubernetesClient                 bool
	FailListProviders                       bool
	ShouldSupportVMIdentity                 bool
	FailDeleteRoleAssignment                bool
	FailEnsureDefaultLogAnalyticsWorkspace  bool
	FailAddContainerInsightsSolution        bool
	FailGetLogAnalyticsWorkspaceInfo        bool
	MockKubernetesClient                    *MockKubernetesClient
	FakeListVirtualMachineScaleSetsResult   func() []compute.VirtualMachineScaleSet
	FakeListVirtualMachineResult            func() []compute.VirtualMachine
	FakeListVirtualMachineScaleSetVMsResult func() []compute.VirtualMachineScaleSetVM
}

//MockStorageClient mock implementation of StorageClient
type MockStorageClient struct {
	FailCreateContainer bool
	FailSaveBlockBlob   bool
}

//MockKubernetesClient mock implementation of KubernetesClient
type MockKubernetesClient struct {
	FailListPods              bool
	FailListNodes             bool
	FailListServiceAccounts   bool
	FailGetNode               bool
	UpdateNodeFunc            func(*v1.Node) (*v1.Node, error)
	GetNodeFunc               func(name string) (*v1.Node, error)
	FailUpdateNode            bool
	FailDeleteNode            bool
	FailDeleteServiceAccount  bool
	FailSupportEviction       bool
	FailDeletePod             bool
	FailDeleteClusterRole     bool
	FailDeleteDaemonSet       bool
	FailDeleteDeployment      bool
	FailEvictPod              bool
	FailWaitForDelete         bool
	ShouldSupportEviction     bool
	PodsList                  *v1.PodList
	ServiceAccountList        *v1.ServiceAccountList
	FailGetDeploymentCount    int
	FailUpdateDeploymentCount int
}

// MockVirtualMachineListResultPage contains a page of VirtualMachine values.
type MockVirtualMachineListResultPage struct {
	Fn   func(compute.VirtualMachineListResult) (compute.VirtualMachineListResult, error)
	Vmlr compute.VirtualMachineListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *MockVirtualMachineListResultPage) Next() error {
	next, err := page.Fn(page.Vmlr)
	if err != nil {
		return err
	}
	page.Vmlr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page MockVirtualMachineListResultPage) NotDone() bool {
	return !page.Vmlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page MockVirtualMachineListResultPage) Response() compute.VirtualMachineListResult {
	return page.Vmlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page MockVirtualMachineListResultPage) Values() []compute.VirtualMachine {
	if page.Vmlr.IsEmpty() {
		return nil
	}
	return *page.Vmlr.Value
}

// MockVirtualMachineScaleSetListResultPage contains a page of VirtualMachine values.
type MockVirtualMachineScaleSetListResultPage struct {
	Fn     func(compute.VirtualMachineScaleSetListResult) (compute.VirtualMachineScaleSetListResult, error)
	Vmsslr compute.VirtualMachineScaleSetListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *MockVirtualMachineScaleSetListResultPage) Next() error {
	next, err := page.Fn(page.Vmsslr)
	if err != nil {
		return err
	}
	page.Vmsslr = next
	return nil
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned. context is ignored in the mock impl.
func (page *MockVirtualMachineScaleSetListResultPage) NextWithContext(context context.Context) error {
	return page.Next()
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page MockVirtualMachineScaleSetListResultPage) NotDone() bool {
	return !page.Vmsslr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page MockVirtualMachineScaleSetListResultPage) Response() compute.VirtualMachineScaleSetListResult {
	return page.Vmsslr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page MockVirtualMachineScaleSetListResultPage) Values() []compute.VirtualMachineScaleSet {
	if page.Vmsslr.IsEmpty() {
		return nil
	}
	return *page.Vmsslr.Value
}

// MockVirtualMachineScaleSetVMListResultPage contains a page of VMSS VirtualMachine values.
type MockVirtualMachineScaleSetVMListResultPage struct {
	Fn      func(compute.VirtualMachineScaleSetVMListResult) (compute.VirtualMachineScaleSetVMListResult, error)
	Vmssvlr compute.VirtualMachineScaleSetVMListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned. Context is ignored for the mock implementation
func (page *MockVirtualMachineScaleSetVMListResultPage) NextWithContext(ctx context.Context) error {
	return page.Next()
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *MockVirtualMachineScaleSetVMListResultPage) Next() error {
	next, err := page.Fn(page.Vmssvlr)
	if err != nil {
		return err
	}
	page.Vmssvlr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page MockVirtualMachineScaleSetVMListResultPage) NotDone() bool {
	return !page.Vmssvlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page MockVirtualMachineScaleSetVMListResultPage) Response() compute.VirtualMachineScaleSetVMListResult {
	return page.Vmssvlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page MockVirtualMachineScaleSetVMListResultPage) Values() []compute.VirtualMachineScaleSetVM {
	if page.Vmssvlr.IsEmpty() {
		return nil
	}
	return *page.Vmssvlr.Value
}

// MockDeploymentOperationsListResultPage contains a page of DeploymentOperation values.
type MockDeploymentOperationsListResultPage struct {
	Fn   func(resources.DeploymentOperationsListResult) (resources.DeploymentOperationsListResult, error)
	Dolr resources.DeploymentOperationsListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *MockDeploymentOperationsListResultPage) Next() error {
	next, err := page.Fn(page.Dolr)
	if err != nil {
		return err
	}
	page.Dolr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page MockDeploymentOperationsListResultPage) NotDone() bool {
	return !page.Dolr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page MockDeploymentOperationsListResultPage) Response() resources.DeploymentOperationsListResult {
	return page.Dolr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page MockDeploymentOperationsListResultPage) Values() []resources.DeploymentOperation {
	if page.Dolr.IsEmpty() {
		return nil
	}
	return *page.Dolr.Value
}

// MockRoleAssignmentListResultPage contains a page of RoleAssignment values.
type MockRoleAssignmentListResultPage struct {
	Fn   func(authorization.RoleAssignmentListResult) (authorization.RoleAssignmentListResult, error)
	Ralr authorization.RoleAssignmentListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *MockRoleAssignmentListResultPage) Next() error {
	next, err := page.Fn(page.Ralr)
	if err != nil {
		return err
	}
	page.Ralr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page MockRoleAssignmentListResultPage) NotDone() bool {
	return !page.Ralr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page MockRoleAssignmentListResultPage) Response() authorization.RoleAssignmentListResult {
	return page.Ralr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page MockRoleAssignmentListResultPage) Values() []authorization.RoleAssignment {
	if page.Ralr.IsEmpty() {
		return nil
	}
	return *page.Ralr.Value
}

//ListPods returns Pods running on the passed in node
func (mkc *MockKubernetesClient) ListPods(node *v1.Node) (*v1.PodList, error) {
	if mkc.FailListPods {
		return nil, errors.New("ListPods failed")
	}
	if mkc.PodsList != nil {
		return mkc.PodsList, nil
	}
	return &v1.PodList{}, nil
}

//ListAllPods returns all Pods running
func (mkc *MockKubernetesClient) ListAllPods() (*v1.PodList, error) {
	if mkc.FailListPods {
		return nil, errors.New("ListAllPods failed")
	}
	if mkc.PodsList != nil {
		return mkc.PodsList, nil
	}
	return &v1.PodList{}, nil
}

// ListNodes returns a list of Nodes registered in the api server
func (mkc *MockKubernetesClient) ListNodes() (*v1.NodeList, error) {
	if mkc.FailListNodes {
		return nil, errors.New("ListNodes failed")
	}
	node := &v1.Node{}
	node.Name = fmt.Sprintf("%s-1234", common.LegacyControlPlaneVMPrefix)
	node.Status.Conditions = append(node.Status.Conditions, v1.NodeCondition{Type: v1.NodeReady, Status: v1.ConditionTrue})
	node.Status.NodeInfo.KubeletVersion = "1.9.10"
	node2 := &v1.Node{}
	node2.Name = "k8s-agentpool3-1234"
	node2.Status.Conditions = append(node2.Status.Conditions, v1.NodeCondition{Type: v1.NodeOutOfDisk, Status: v1.ConditionTrue})
	node2.Status.NodeInfo.KubeletVersion = "1.9.9"
	nodeList := &v1.NodeList{}
	nodeList.Items = append(nodeList.Items, *node)
	nodeList.Items = append(nodeList.Items, *node2)
	return nodeList, nil
}

// ListServiceAccounts returns a list of Service Accounts in the provided namespace
func (mkc *MockKubernetesClient) ListServiceAccounts(namespace string) (*v1.ServiceAccountList, error) {
	if mkc.FailListServiceAccounts {
		return nil, errors.New("ListServiceAccounts failed")
	}
	if mkc.ServiceAccountList != nil {
		return mkc.ServiceAccountList, nil
	}
	sa := &v1.ServiceAccount{}
	sa.Namespace = namespace
	sa.Name = "service-account-1"
	sa2 := &v1.ServiceAccount{}
	sa2.Namespace = namespace
	sa.Name = "service-account-2"
	saList := &v1.ServiceAccountList{}
	saList.Items = append(saList.Items, *sa)
	saList.Items = append(saList.Items, *sa2)
	return saList, nil
}

//GetNode returns details about node with passed in name
func (mkc *MockKubernetesClient) GetNode(name string) (*v1.Node, error) {
	if mkc.GetNodeFunc != nil {
		return mkc.GetNodeFunc(name)
	}
	if mkc.FailGetNode {
		return nil, errors.New("GetNode failed")
	}
	node := &v1.Node{}
	node.Status.Conditions = append(node.Status.Conditions, v1.NodeCondition{Type: v1.NodeReady, Status: v1.ConditionTrue})
	node.Status.NodeInfo.KubeletVersion = "1.17.5"
	return node, nil
}

//UpdateNode updates the node in the api server with the passed in info
func (mkc *MockKubernetesClient) UpdateNode(node *v1.Node) (*v1.Node, error) {
	if mkc.UpdateNodeFunc != nil {
		return mkc.UpdateNodeFunc(node)
	}
	if mkc.FailUpdateNode {
		return nil, errors.New("UpdateNode failed")
	}
	return node, nil
}

//DeleteNode deregisters node in the api server
func (mkc *MockKubernetesClient) DeleteNode(name string) error {
	if mkc.FailDeleteNode {
		return errors.New("DeleteNode failed")
	}
	return nil
}

// DeleteServiceAccount deletes the provided service account
func (mkc *MockKubernetesClient) DeleteServiceAccount(sa *v1.ServiceAccount) error {
	if mkc.FailDeleteServiceAccount {
		return errors.New("DeleteServiceAccount failed")
	}
	return nil
}

//SupportEviction queries the api server to discover if it supports eviction, and returns supported type if it is supported
func (mkc *MockKubernetesClient) SupportEviction() (string, error) {
	if mkc.FailSupportEviction {
		return "", errors.New("SupportEviction failed")
	}
	if mkc.ShouldSupportEviction {
		return "version", nil
	}
	return "", nil
}

//DeleteDeployment deletes the passed in daemonset
func (mkc *MockKubernetesClient) DeleteClusterRole(role *rbacv1.ClusterRole) error {
	if mkc.FailDeleteClusterRole {
		return errors.New("ClusterRole failed")
	}
	return nil
}

//DeleteDaemonSet deletes the passed in daemonset
func (mkc *MockKubernetesClient) DeleteDaemonSet(pod *appsv1.DaemonSet) error {
	if mkc.FailDeleteDaemonSet {
		return errors.New("DaemonSet failed")
	}
	return nil
}

//DeleteDeployment deletes the passed in daemonset
func (mkc *MockKubernetesClient) DeleteDeployment(pod *appsv1.Deployment) error {
	if mkc.FailDeleteDeployment {
		return errors.New("Deployment failed")
	}
	return nil
}

//DeletePod deletes the passed in pod
func (mkc *MockKubernetesClient) DeletePod(pod *v1.Pod) error {
	if mkc.FailDeletePod {
		return errors.New("DeletePod failed")
	}
	return nil
}

//EvictPod evicts the passed in pod using the passed in api version
func (mkc *MockKubernetesClient) EvictPod(pod *v1.Pod, policyGroupVersion string) error {
	if mkc.FailEvictPod {
		return errors.New("EvictPod failed")
	}
	return nil
}

//WaitForDelete waits until all pods are deleted. Returns all pods not deleted and an error on failure
func (mkc *MockKubernetesClient) WaitForDelete(logger *log.Entry, pods []v1.Pod, usingEviction bool) ([]v1.Pod, error) {
	if mkc.FailWaitForDelete {
		return nil, errors.New("WaitForDelete failed")
	}
	return []v1.Pod{}, nil
}

// DaemonSet returns a given daemonset in a namespace.
func (mkc *MockKubernetesClient) GetDaemonSet(namespace, name string) (*appsv1.DaemonSet, error) {
	return &appsv1.DaemonSet{
		Spec: appsv1.DaemonSetSpec{},
	}, nil
}

// GetDeployment returns a given deployment in a namespace.
func (mkc *MockKubernetesClient) GetDeployment(namespace, name string) (*appsv1.Deployment, error) {
	if mkc.FailGetDeploymentCount > 0 {
		mkc.FailGetDeploymentCount--
		return nil, errors.New("GetDeployment failed")
	}
	var replicas int32 = 1
	return &appsv1.Deployment{
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
		},
	}, nil
}

// UpdateDeployment updates a deployment to match the given specification.
func (mkc *MockKubernetesClient) UpdateDeployment(namespace string, deployment *appsv1.Deployment) (*appsv1.Deployment, error) {
	if mkc.FailUpdateDeploymentCount > 0 {
		mkc.FailUpdateDeploymentCount--
		return nil, errors.New("UpdateDeployment failed")
	}
	return &appsv1.Deployment{}, nil
}

//DeleteBlob mock
func (msc *MockStorageClient) DeleteBlob(container, blob string, options *azStorage.DeleteBlobOptions) error {
	return nil
}

//CreateContainer mock
func (msc *MockStorageClient) CreateContainer(container string, options *azStorage.CreateContainerOptions) (bool, error) {
	if !msc.FailCreateContainer {
		return true, nil
	}
	return false, errors.New("CreateContainer failed")
}

//SaveBlockBlob mock
func (msc *MockStorageClient) SaveBlockBlob(container, blob string, b []byte, options *azStorage.PutBlobOptions) error {
	if !msc.FailSaveBlockBlob {
		return nil
	}
	return errors.New("SaveBlockBlob failed")
}

//AddAcceptLanguages mock
func (mc *MockAKSEngineClient) AddAcceptLanguages(languages []string) {}

// AddAuxiliaryTokens mock
func (mc *MockAKSEngineClient) AddAuxiliaryTokens(tokens []string) {}

//DeployTemplate mock
func (mc *MockAKSEngineClient) DeployTemplate(ctx context.Context, resourceGroup, name string, template, parameters map[string]interface{}) (de resources.DeploymentExtended, err error) {
	switch {
	case mc.FailDeployTemplate:
		return de, errors.New("DeployTemplate failed")

	case mc.FailDeployTemplateQuota:
		errmsg := `resources.DeploymentsClient#CreateOrUpdate: Failure responding to request: StatusCode=400 -- Original Error: autorest/azure: Service returned an error.`
		resp := `{
"error":{
	"code":"InvalidTemplateDeployment",
	"message":"The template deployment is not valid according to the validation procedure. The tracking id is 'b5bd7d6b-fddf-4ec3-a3b0-ce285a48bd31'. See inner errors for details. Please see https://aka.ms/arm-deploy for usage details.",
	"details":[{
		"code":"QuotaExceeded",
		"message":"Operation results in exceeding quota limits of Core. Maximum allowed: 10, Current in use: 10, Additional requested: 2. Please read more about quota increase at http://aka.ms/corequotaincrease."
}]}}`

		return resources.DeploymentExtended{
				Response: autorest.Response{
					Response: &http.Response{
						Status:     "400 Bad Request",
						StatusCode: 400,
						Body:       ioutil.NopCloser(bytes.NewReader([]byte(resp))),
					}}},
			errors.New(errmsg)

	case mc.FailDeployTemplateConflict:
		errmsg := `resources.DeploymentsClient#CreateOrUpdate: Failure sending request: StatusCode=200 -- Original Error: Long running operation terminated with status 'Failed': Code="DeploymentFailed" Message="At least one resource deployment operation failed. Please list deployment operations for details. Please see https://aka.ms/arm-debug for usage details.`
		resp := `{
"status":"Failed",
"error":{
	"code":"DeploymentFailed",
	"message":"At least one resource deployment operation failed. Please list deployment operations for details. Please see https://aka.ms/arm-debug for usage details.",
	"details":[{
		"code":"Conflict",
		"message":"{\r\n  \"error\": {\r\n    \"code\": \"PropertyChangeNotAllowed\",\r\n    \"target\": \"dataDisk.createOption\",\r\n    \"message\": \"Changing property 'dataDisk.createOption' is not allowed.\"\r\n  }\r\n}"
}]}}`
		return resources.DeploymentExtended{
				Response: autorest.Response{
					Response: &http.Response{
						Status:     "200 OK",
						StatusCode: 200,
						Body:       ioutil.NopCloser(bytes.NewReader([]byte(resp))),
					}}},
			errors.New(errmsg)

	case mc.FailDeployTemplateWithProperties:
		errmsg := `resources.DeploymentsClient#CreateOrUpdate: Failure sending request: StatusCode=200 -- Original Error: Long running operation terminated with status 'Failed': Code="DeploymentFailed" Message="At least one resource deployment operation failed. Please list deployment operations for details. Please see https://aka.ms/arm-debug for usage details.`
		resp := `{
"status":"Failed",
"error":{
	"code":"DeploymentFailed",
	"message":"At least one resource deployment operation failed. Please list deployment operations for details. Please see https://aka.ms/arm-debug for usage details.",
	"details":[{
		"code":"Conflict",
		"message":"{\r\n  \"error\": {\r\n    \"code\": \"PropertyChangeNotAllowed\",\r\n    \"target\": \"dataDisk.createOption\",\r\n    \"message\": \"Changing property 'dataDisk.createOption' is not allowed.\"\r\n  }\r\n}"
}]}}`
		provisioningState := "Failed"
		return resources.DeploymentExtended{
				Response: autorest.Response{
					Response: &http.Response{
						Status:     "200 OK",
						StatusCode: 200,
						Body:       ioutil.NopCloser(bytes.NewReader([]byte(resp))),
					}},
				Properties: &resources.DeploymentPropertiesExtended{
					ProvisioningState: &provisioningState,
				}},
			errors.New(errmsg)
	default:
		return de, nil
	}
}

// ListLocations mock
func (mc *MockAKSEngineClient) ListLocations(ctx context.Context) (*[]subscriptions.Location, error) {
	locations := []subscriptions.Location{}
	return &locations, nil
}

//EnsureResourceGroup mock
func (mc *MockAKSEngineClient) EnsureResourceGroup(ctx context.Context, resourceGroup, location string, managedBy *string) (*resources.Group, error) {
	if mc.FailEnsureResourceGroup {
		return nil, errors.New("EnsureResourceGroup failed")
	}

	return nil, nil
}

// ListResourceSkus mock
func (mc *MockAKSEngineClient) ListResourceSkus(ctx context.Context, filter string) (ResourceSkusResultPage, error) {
	return nil, nil
}

//ListVirtualMachines mock
func (mc *MockAKSEngineClient) ListVirtualMachines(ctx context.Context, resourceGroup string) (VirtualMachineListResultPage, error) {
	if mc.FailListVirtualMachines {
		return &MockVirtualMachineListResultPage{
			Vmlr: compute.VirtualMachineListResult{
				Value: &[]compute.VirtualMachine{{}},
			},
		}, errors.New("ListVirtualMachines failed")
	}

	if mc.FakeListVirtualMachineResult == nil {
		mc.FakeListVirtualMachineResult = func() []compute.VirtualMachine {
			machine := mc.MakeFakeVirtualMachine(DefaultFakeVMName, defaultK8sVersionForFakeVMs)
			machine.AvailabilitySet = &compute.SubResource{
				ID: to.StringPtr("MockAvailabilitySet"),
			}
			return []compute.VirtualMachine{machine}
		}
	}
	vms := mc.FakeListVirtualMachineResult()
	vmr := compute.VirtualMachineListResult{}
	vmr.Value = &vms

	return &MockVirtualMachineListResultPage{
		Fn: func(lastResults compute.VirtualMachineListResult) (compute.VirtualMachineListResult, error) {
			return compute.VirtualMachineListResult{}, nil
		},
		Vmlr: vmr,
	}, nil
}

//ListVirtualMachineScaleSets mock
func (mc *MockAKSEngineClient) ListVirtualMachineScaleSets(ctx context.Context, resourceGroup string) (VirtualMachineScaleSetListResultPage, error) {
	if mc.FailListVirtualMachineScaleSets {
		return &MockVirtualMachineScaleSetListResultPage{}, errors.New("ListVirtualMachineScaleSets failed")
	}
	if mc.FakeListVirtualMachineScaleSetsResult == nil {
		//return 0 machines by default
		mc.FakeListVirtualMachineScaleSetsResult = func() []compute.VirtualMachineScaleSet {
			return []compute.VirtualMachineScaleSet{}
		}
	}

	vmsslr := compute.VirtualMachineScaleSetListResult{}
	vmss := mc.FakeListVirtualMachineScaleSetsResult()
	vmsslr.Value = &vmss

	return &MockVirtualMachineScaleSetListResultPage{
		Fn: func(compute.VirtualMachineScaleSetListResult) (compute.VirtualMachineScaleSetListResult, error) {
			return compute.VirtualMachineScaleSetListResult{}, nil
		},
		Vmsslr: vmsslr,
	}, nil
}

// RestartVirtualMachineScaleSets mock
func (mc *MockAKSEngineClient) RestartVirtualMachineScaleSets(ctx context.Context, resourceGroup, name string, instanceIDs *compute.VirtualMachineScaleSetVMInstanceIDs) error {
	if mc.FailRestartVirtualMachineScaleSets {
		return errors.New("RestartVirtualMachineScaleSets failed")
	}
	return nil
}

//GetVirtualMachine mock
func (mc *MockAKSEngineClient) GetVirtualMachine(ctx context.Context, resourceGroup, name string) (compute.VirtualMachine, error) {
	if mc.FailGetVirtualMachine {
		return compute.VirtualMachine{}, errors.New("GetVirtualMachine failed")
	}
	return mc.MakeFakeVirtualMachine(DefaultFakeVMName, defaultK8sVersionForFakeVMs), nil
}

// RestartVirtualMachine mock
func (mc *MockAKSEngineClient) RestartVirtualMachine(ctx context.Context, resourceGroup, name string) error {
	if mc.FailRestartVirtualMachine {
		return errors.New("RestartVirtualMachine failed")
	}
	return nil
}

// MakeFakeVirtualMachineScaleSetVM creates a fake VMSS VM
func (mc *MockAKSEngineClient) MakeFakeVirtualMachineScaleSetVM(orchestratorTag string) compute.VirtualMachineScaleSetVM {
	return mc.MakeFakeVirtualMachineScaleSetVMWithGivenName(orchestratorTag, "computerName")
}

// MakeFakeVirtualMachineScaleSetVM creates a fake VMSS VM with name provided
func (mc *MockAKSEngineClient) MakeFakeVirtualMachineScaleSetVMWithGivenName(orchestratorTag string, computerName string) compute.VirtualMachineScaleSetVM {
	vm1Name := "k8s-agentpool1-12345678-0"

	creationSourceString := "creationSource"
	orchestratorString := "orchestrator"
	resourceNameSuffixString := "resourceNameSuffix"
	poolnameString := "poolName"

	creationSource := "aksengine-k8s-agentpool1-12345678-0"
	orchestrator := orchestratorTag
	resourceNameSuffix := "12345678"
	poolname := "agentpool1"
	instanceID := "someguidthatshouldbeunique"

	tags := map[string]*string{
		creationSourceString:     &creationSource,
		orchestratorString:       &orchestrator,
		resourceNameSuffixString: &resourceNameSuffix,
		poolnameString:           &poolname,
	}

	if mc.FailListVirtualMachinesTags {
		tags = nil
	}

	trueVar := true
	return compute.VirtualMachineScaleSetVM{
		Name:       &vm1Name,
		Tags:       tags,
		InstanceID: &instanceID,
		VirtualMachineScaleSetVMProperties: &compute.VirtualMachineScaleSetVMProperties{
			LatestModelApplied: &trueVar,
			OsProfile:          &compute.OSProfile{ComputerName: &computerName},
			StorageProfile: &compute.StorageProfile{
				OsDisk: &compute.OSDisk{
					Vhd: &compute.VirtualHardDisk{
						URI: &validOSDiskResourceName},
				},
			},
			NetworkProfile: &compute.NetworkProfile{
				NetworkInterfaces: &[]compute.NetworkInterfaceReference{
					{
						ID: &validNicResourceName,
					},
				},
			},
		},
	}
}

//MakeFakeVirtualMachine returns a fake compute.VirtualMachine
func (mc *MockAKSEngineClient) MakeFakeVirtualMachine(vmName string, orchestratorVersion string) compute.VirtualMachine {
	vm1Name := vmName

	creationSourceString := "creationSource"
	orchestratorString := "orchestrator"
	resourceNameSuffixString := "resourceNameSuffix"
	poolnameString := "poolName"

	creationSource := "aksengine-k8s-agentpool1-12345678-0"
	orchestrator := orchestratorVersion
	resourceNameSuffix := "12345678"
	poolname := "agentpool1"

	principalID := "00000000-1111-2222-3333-444444444444"

	tags := map[string]*string{
		creationSourceString:     &creationSource,
		orchestratorString:       &orchestrator,
		resourceNameSuffixString: &resourceNameSuffix,
		poolnameString:           &poolname,
	}

	var vmIdentity *compute.VirtualMachineIdentity
	if mc.ShouldSupportVMIdentity {
		vmIdentity = &compute.VirtualMachineIdentity{PrincipalID: &principalID}
	}

	if mc.FailListVirtualMachinesTags {
		tags = nil
	}

	return compute.VirtualMachine{
		Name:     &vm1Name,
		Tags:     tags,
		Identity: vmIdentity,
		VirtualMachineProperties: &compute.VirtualMachineProperties{
			StorageProfile: &compute.StorageProfile{
				OsDisk: &compute.OSDisk{
					Vhd: &compute.VirtualHardDisk{
						URI: &validOSDiskResourceName},
				},
			},
			NetworkProfile: &compute.NetworkProfile{
				NetworkInterfaces: &[]compute.NetworkInterfaceReference{
					{
						ID: &validNicResourceName,
					},
				},
			},
		},
	}
}

//DeleteVirtualMachine mock
func (mc *MockAKSEngineClient) DeleteVirtualMachine(ctx context.Context, resourceGroup, name string) error {
	if mc.FailDeleteVirtualMachine {
		return errors.New("DeleteVirtualMachine failed")
	}

	return nil
}

//DeleteVirtualMachineScaleSetVM mock
func (mc *MockAKSEngineClient) DeleteVirtualMachineScaleSetVM(ctx context.Context, resourceGroup, virtualMachineScaleSet, instanceID string) error {
	if mc.FailDeleteVirtualMachineScaleSetVM {
		return errors.New("DeleteVirtualMachineScaleSetVM failed")
	}

	return nil
}

//SetVirtualMachineScaleSetCapacity mock
func (mc *MockAKSEngineClient) SetVirtualMachineScaleSetCapacity(ctx context.Context, resourceGroup, virtualMachineScaleSet string, sku compute.Sku, location string) error {
	if mc.FailSetVirtualMachineScaleSetCapacity {
		return errors.New("SetVirtualMachineScaleSetCapacity failed")
	}

	return nil
}

//ListVirtualMachineScaleSetVMs mock
func (mc *MockAKSEngineClient) ListVirtualMachineScaleSetVMs(ctx context.Context, resourceGroup, virtualMachineScaleSet string) (VirtualMachineScaleSetVMListResultPage, error) {
	if mc.FailDeleteVirtualMachineScaleSetVM {
		return &compute.VirtualMachineScaleSetVMListResultPage{}, errors.New("DeleteVirtualMachineScaleSetVM failed")
	}

	if mc.FakeListVirtualMachineScaleSetVMsResult == nil {
		//return 0 machined by default
		mc.FakeListVirtualMachineScaleSetVMsResult = func() []compute.VirtualMachineScaleSetVM {
			return []compute.VirtualMachineScaleSetVM{}
		}
	}

	result := MockVirtualMachineScaleSetVMListResultPage{}
	vms := mc.FakeListVirtualMachineScaleSetVMsResult()
	result.Vmssvlr = compute.VirtualMachineScaleSetVMListResult{Value: &vms}

	return &MockVirtualMachineScaleSetVMListResultPage{
		Fn: func(compute.VirtualMachineScaleSetVMListResult) (compute.VirtualMachineScaleSetVMListResult, error) {
			return compute.VirtualMachineScaleSetVMListResult{}, nil
		},
		Vmssvlr: compute.VirtualMachineScaleSetVMListResult{Value: &vms},
	}, nil
}

// GetAvailabilitySet mock
func (mc *MockAKSEngineClient) GetAvailabilitySet(ctx context.Context, resourceGroup, availabilitySetName string) (compute.AvailabilitySet, error) {
	return compute.AvailabilitySet{}, nil
}

// GetAvailabilitySetFaultDomainCount mock
func (mc *MockAKSEngineClient) GetAvailabilitySetFaultDomainCount(ctx context.Context, resourceGroup string, vmasIDs []string) (int, error) {
	return 3, nil
}

//GetStorageClient mock
func (mc *MockAKSEngineClient) GetStorageClient(ctx context.Context, resourceGroup, accountName string) (AKSStorageClient, error) {
	if mc.FailGetStorageClient {
		return nil, errors.New("GetStorageClient failed")
	}

	return &MockStorageClient{}, nil
}

//DeleteNetworkInterface mock
func (mc *MockAKSEngineClient) DeleteNetworkInterface(ctx context.Context, resourceGroup, nicName string) error {
	if mc.FailDeleteNetworkInterface {
		return errors.New("DeleteNetworkInterface failed")
	}

	return nil
}

var validOSDiskResourceName = "https://00k71r4u927seqiagnt0.blob.core.windows.net/osdisk/k8s-agentpool1-12345678-0-osdisk.vhd"
var validNicResourceName = "/subscriptions/DEC923E3-1EF1-4745-9516-37906D56DEC4/resourceGroups/acsK8sTest/providers/Microsoft.Network/networkInterfaces/k8s-agent-12345678-nic-0"

// Active Directory
// Mocks

// Graph Mocks

// CreateGraphApplication creates an application via the graphrbac client
func (mc *MockAKSEngineClient) CreateGraphApplication(ctx context.Context, applicationCreateParameters graphrbac.ApplicationCreateParameters) (graphrbac.Application, error) {
	return graphrbac.Application{}, nil
}

// CreateGraphPrincipal creates a service principal via the graphrbac client
func (mc *MockAKSEngineClient) CreateGraphPrincipal(ctx context.Context, servicePrincipalCreateParameters graphrbac.ServicePrincipalCreateParameters) (graphrbac.ServicePrincipal, error) {
	return graphrbac.ServicePrincipal{}, nil
}

// CreateApp is a simpler method for creating an application
func (mc *MockAKSEngineClient) CreateApp(ctx context.Context, applicationName, applicationURL string, replyURLs *[]string, requiredResourceAccess *[]graphrbac.RequiredResourceAccess) (result graphrbac.Application, servicePrincipalObjectID, secret string, err error) {
	return graphrbac.Application{
		AppID: to.StringPtr("app-id"),
	}, "client-id", "client-secret", nil
}

// DeleteApp is a simpler method for deleting an application
func (mc *MockAKSEngineClient) DeleteApp(ctx context.Context, appName, applicationObjectID string) (response autorest.Response, err error) {
	return response, nil
}

// User Assigned MSI

//CreateUserAssignedID - Creates a user assigned msi.
func (mc *MockAKSEngineClient) CreateUserAssignedID(location string, resourceGroup string, userAssignedID string) (*msi.Identity, error) {
	return &msi.Identity{}, nil
}

// RBAC Mocks

// CreateRoleAssignment creates a role assignment via the authorization client
func (mc *MockAKSEngineClient) CreateRoleAssignment(ctx context.Context, scope string, roleAssignmentName string, parameters authorization.RoleAssignmentCreateParameters) (authorization.RoleAssignment, error) {
	return authorization.RoleAssignment{}, nil
}

// CreateRoleAssignmentSimple is a wrapper around RoleAssignmentsClient.Create
func (mc *MockAKSEngineClient) CreateRoleAssignmentSimple(ctx context.Context, applicationID, roleID string) error {
	return nil
}

// DeleteManagedDisk is a wrapper around disksClient.Delete
func (mc *MockAKSEngineClient) DeleteManagedDisk(ctx context.Context, resourceGroupName string, diskName string) error {
	return nil
}

// ListManagedDisksByResourceGroup is a wrapper around disksClient.ListManagedDisksByResourceGroup
func (mc *MockAKSEngineClient) ListManagedDisksByResourceGroup(ctx context.Context, resourceGroupName string) (result DiskListPage, err error) {
	return &compute.DiskListPage{}, nil
}

//GetKubernetesClient mock
func (mc *MockAKSEngineClient) GetKubernetesClient(apiserverURL, kubeConfig string, interval, timeout time.Duration) (KubernetesClient, error) {
	if mc.FailGetKubernetesClient {
		return nil, errors.New("GetKubernetesClient failed")
	}

	if mc.MockKubernetesClient == nil {
		mc.MockKubernetesClient = &MockKubernetesClient{}
	}
	return mc.MockKubernetesClient, nil
}

// ListProviders mock
func (mc *MockAKSEngineClient) ListProviders(ctx context.Context) (ProviderListResultPage, error) {
	if mc.FailListProviders {
		return &resources.ProviderListResultPage{}, errors.New("ListProviders failed")
	}

	return &resources.ProviderListResultPage{}, nil
}

// ListDeploymentOperations gets all deployments operations for a deployment.
func (mc *MockAKSEngineClient) ListDeploymentOperations(ctx context.Context, resourceGroupName string, deploymentName string, top *int32) (result DeploymentOperationsListResultPage, err error) {
	resp := `{
	"properties": {
	"provisioningState":"Failed",
	"correlationId":"d5062e45-6e9f-4fd3-a0a0-6b2c56b15757",
	"error":{
	"code":"DeploymentFailed","message":"At least one resource deployment operation failed. Please list deployment operations for details. Please see http://aka.ms/arm-debug for usage details.",
	"details":[{"code":"Conflict","message":"{\r\n  \"error\": {\r\n    \"message\": \"Conflict\",\r\n    \"code\": \"Conflict\"\r\n  }\r\n}"}]
	}
	}
	}`

	provisioningState := "Failed"
	id := "00000000"
	operationID := "d5062e45-6e9f-4fd3-a0a0-6b2c56b15757"
	nextLink := fmt.Sprintf("https://management.azure.com/subscriptions/11111/resourcegroups/%s/deployments/%s/operations?$top=%s&api-version=2018-05-01", resourceGroupName, deploymentName, "5")
	return &MockDeploymentOperationsListResultPage{
		Fn: func(lastResults resources.DeploymentOperationsListResult) (result resources.DeploymentOperationsListResult, err error) {
			if lastResults.NextLink != nil {
				return resources.DeploymentOperationsListResult{
					Response: autorest.Response{
						Response: &http.Response{
							Status:     "200 OK",
							StatusCode: 200,
							Body:       ioutil.NopCloser(bytes.NewReader([]byte(resp))),
						},
					},
					Value: &[]resources.DeploymentOperation{
						{
							ID:          &id,
							OperationID: &operationID,
							Properties: &resources.DeploymentOperationProperties{
								ProvisioningState: &provisioningState,
							},
						},
					},
				}, nil
			}
			return resources.DeploymentOperationsListResult{}, nil
		},
		Dolr: resources.DeploymentOperationsListResult{
			Response: autorest.Response{
				Response: &http.Response{
					Status:     "200 OK",
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(resp))),
				},
			},
			Value: &[]resources.DeploymentOperation{
				{
					ID:          &id,
					OperationID: &operationID,
					Properties: &resources.DeploymentOperationProperties{
						ProvisioningState: &provisioningState,
					},
				},
			},
			NextLink: &nextLink,
		},
	}, nil
}

// ListDeploymentOperationsNextResults retrieves the next set of results, if any.
func (mc *MockAKSEngineClient) ListDeploymentOperationsNextResults(lastResults resources.DeploymentOperationsListResult) (result resources.DeploymentOperationsListResult, err error) {
	return resources.DeploymentOperationsListResult{}, nil
}

// DeleteRoleAssignmentByID deletes a roleAssignment via its unique identifier
func (mc *MockAKSEngineClient) DeleteRoleAssignmentByID(ctx context.Context, roleAssignmentID string) (authorization.RoleAssignment, error) {
	if mc.FailDeleteRoleAssignment {
		return authorization.RoleAssignment{}, errors.New("DeleteRoleAssignmentByID failed")
	}

	return authorization.RoleAssignment{}, nil
}

// ListRoleAssignmentsForPrincipal (e.g. a VM) via the scope and the unique identifier of the principal
func (mc *MockAKSEngineClient) ListRoleAssignmentsForPrincipal(ctx context.Context, scope string, principalID string) (RoleAssignmentListResultPage, error) {
	roleAssignments := []authorization.RoleAssignment{}

	if mc.ShouldSupportVMIdentity {
		var assignmentID = "role-assignment-id"
		var assignment = authorization.RoleAssignment{
			ID: &assignmentID}
		roleAssignments = append(roleAssignments, assignment)
	}

	return &MockRoleAssignmentListResultPage{
		Ralr: authorization.RoleAssignmentListResult{
			Value: &roleAssignments,
		},
	}, nil
}

//EnsureDefaultLogAnalyticsWorkspace mock
func (mc *MockAKSEngineClient) EnsureDefaultLogAnalyticsWorkspace(ctx context.Context, resourceGroup, location string) (workspaceResourceID string, err error) {
	if mc.FailEnsureDefaultLogAnalyticsWorkspace {
		return "", errors.New("EnsureDefaultLogAnalyticsWorkspace failed")
	}

	return "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-workspace-rg/providers/Microsoft.OperationalInsights/workspaces/test-workspace", nil
}

//AddContainerInsightsSolution mock
func (mc *MockAKSEngineClient) AddContainerInsightsSolution(ctx context.Context, workspaceSubscriptionID, workspaceResourceGroup, workspaceName, workspaceLocation string) (result bool, err error) {
	if mc.FailAddContainerInsightsSolution {
		return false, errors.New("AddContainerInsightsSolution failed")
	}

	return true, nil
}

//GetLogAnalyticsWorkspaceInfo mock
func (mc *MockAKSEngineClient) GetLogAnalyticsWorkspaceInfo(ctx context.Context, workspaceSubscriptionID, workspaceResourceGroup, workspaceName string) (workspaceID string, workspaceKey, workspaceLocation string, err error) {
	if mc.FailGetLogAnalyticsWorkspaceInfo {
		return "", "", "", errors.New("GetLogAnalyticsWorkspaceInfo failed")
	}

	return "00000000-0000-0000-0000-000000000000", "4D+vyd5/jScBmsAwZOF/0GOBQ5kuFQc9JVaW+HlnJ58cyePJcwTpks+rVmvgcXGmmyujLDNEVPiT8pB274a9Yg==", "westus", nil
}
