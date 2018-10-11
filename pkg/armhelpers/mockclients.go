package armhelpers

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-04-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/azure-sdk-for-go/services/preview/msi/mgmt/2015-08-31-preview/msi"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	azStorage "github.com/Azure/azure-sdk-for-go/storage"
	"github.com/Azure/go-autorest/autorest"
	log "github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"
)

//MockAKSEngineClient is an implementation of AKSEngineClient where all requests error out
type MockAKSEngineClient struct {
	FailDeployTemplate                    bool
	FailDeployTemplateQuota               bool
	FailDeployTemplateConflict            bool
	FailDeployTemplateWithProperties      bool
	FailEnsureResourceGroup               bool
	FailListVirtualMachines               bool
	FailListVirtualMachinesTags           bool
	FailListVirtualMachineScaleSets       bool
	FailGetVirtualMachine                 bool
	FailDeleteVirtualMachine              bool
	FailDeleteVirtualMachineScaleSetVM    bool
	FailSetVirtualMachineScaleSetCapacity bool
	FailListVirtualMachineScaleSetVMs     bool
	FailGetStorageClient                  bool
	FailDeleteNetworkInterface            bool
	FailGetKubernetesClient               bool
	FailListProviders                     bool
	ShouldSupportVMIdentity               bool
	FailDeleteRoleAssignment              bool
	MockKubernetesClient                  *MockKubernetesClient
}

//MockStorageClient mock implementation of StorageClient
type MockStorageClient struct {
	FailCreateContainer bool
	FailSaveBlockBlob   bool
}

//MockKubernetesClient mock implementation of KubernetesClient
type MockKubernetesClient struct {
	FailListPods          bool
	FailGetNode           bool
	UpdateNodeFunc        func(*v1.Node) (*v1.Node, error)
	FailUpdateNode        bool
	FailDeleteNode        bool
	FailSupportEviction   bool
	FailDeletePod         bool
	FailEvictPod          bool
	FailWaitForDelete     bool
	ShouldSupportEviction bool
	PodsList              *v1.PodList
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

//ListPods returns all Pods running on the passed in node
func (mkc *MockKubernetesClient) ListPods(node *v1.Node) (*v1.PodList, error) {
	if mkc.FailListPods {
		return nil, errors.New("ListPods failed")
	}
	if mkc.PodsList != nil {
		return mkc.PodsList, nil
	}
	return &v1.PodList{}, nil
}

//GetNode returns details about node with passed in name
func (mkc *MockKubernetesClient) GetNode(name string) (*v1.Node, error) {
	if mkc.FailGetNode {
		return nil, errors.New("GetNode failed")
	}
	node := &v1.Node{}
	node.Status.Conditions = append(node.Status.Conditions, v1.NodeCondition{Type: v1.NodeReady, Status: v1.ConditionTrue})
	node.Status.NodeInfo.KubeletVersion = "1.7.9"
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

//EnsureResourceGroup mock
func (mc *MockAKSEngineClient) EnsureResourceGroup(ctx context.Context, resourceGroup, location string, managedBy *string) (*resources.Group, error) {
	if mc.FailEnsureResourceGroup {
		return nil, errors.New("EnsureResourceGroup failed")
	}

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

	vm1Name := "k8s-agentpool1-12345678-0"

	creationSourceString := "creationSource"
	orchestratorString := "orchestrator"
	resourceNameSuffixString := "resourceNameSuffix"
	poolnameString := "poolName"

	creationSource := "acsengine-k8s-agentpool1-12345678-0"
	orchestrator := "Kubernetes:1.7.9"
	resourceNameSuffix := "12345678"
	poolname := "agentpool1"

	tags := map[string]*string{
		creationSourceString:     &creationSource,
		orchestratorString:       &orchestrator,
		resourceNameSuffixString: &resourceNameSuffix,
		poolnameString:           &poolname,
	}
	if mc.FailListVirtualMachinesTags {
		tags = nil
	}

	vm1 := compute.VirtualMachine{
		Name: &vm1Name,
		Tags: tags,
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

	vmr := compute.VirtualMachineListResult{}
	vmr.Value = &[]compute.VirtualMachine{vm1}

	return &MockVirtualMachineListResultPage{
		Fn: func(lastResults compute.VirtualMachineListResult) (compute.VirtualMachineListResult, error) {
			return compute.VirtualMachineListResult{}, nil
		},
		Vmlr: vmr,
	}, nil
}

//ListVirtualMachineScaleSets mock
func (mc *MockAKSEngineClient) ListVirtualMachineScaleSets(ctx context.Context, resourceGroup string) (compute.VirtualMachineScaleSetListResultPage, error) {
	if mc.FailListVirtualMachineScaleSets {
		return compute.VirtualMachineScaleSetListResultPage{}, errors.New("ListVirtualMachines failed")
	}

	return compute.VirtualMachineScaleSetListResultPage{}, nil
}

//GetVirtualMachine mock
func (mc *MockAKSEngineClient) GetVirtualMachine(ctx context.Context, resourceGroup, name string) (compute.VirtualMachine, error) {
	if mc.FailGetVirtualMachine {
		return compute.VirtualMachine{}, errors.New("GetVirtualMachine failed")
	}

	vm1Name := "k8s-agentpool1-12345678-0"

	creationSourceString := "creationSource"
	orchestratorString := "orchestrator"
	resourceNameSuffixString := "resourceNameSuffix"
	poolnameString := "poolName"

	creationSource := "acsengine-k8s-agentpool1-12345678-0"
	orchestrator := "Kubernetes:1.7.9"
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
	}, nil
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
func (mc *MockAKSEngineClient) ListVirtualMachineScaleSetVMs(ctx context.Context, resourceGroup, virtualMachineScaleSet string) (compute.VirtualMachineScaleSetVMListResultPage, error) {
	if mc.FailDeleteVirtualMachineScaleSetVM {
		return compute.VirtualMachineScaleSetVMListResultPage{}, errors.New("DeleteVirtualMachineScaleSetVM failed")
	}

	return compute.VirtualMachineScaleSetVMListResultPage{}, nil
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
		AppID: helpers.PointerToString("app-id"),
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
func (mc *MockAKSEngineClient) ListManagedDisksByResourceGroup(ctx context.Context, resourceGroupName string) (result compute.DiskListPage, err error) {
	return compute.DiskListPage{}, nil
}

//GetKubernetesClient mock
func (mc *MockAKSEngineClient) GetKubernetesClient(masterURL, kubeConfig string, interval, timeout time.Duration) (KubernetesClient, error) {
	if mc.FailGetKubernetesClient {
		return nil, errors.New("GetKubernetesClient failed")
	}

	if mc.MockKubernetesClient == nil {
		mc.MockKubernetesClient = &MockKubernetesClient{}
	}
	return mc.MockKubernetesClient, nil
}

// ListProviders mock
func (mc *MockAKSEngineClient) ListProviders(ctx context.Context) (resources.ProviderListResultPage, error) {
	if mc.FailListProviders {
		return resources.ProviderListResultPage{}, errors.New("ListProviders failed")
	}

	return resources.ProviderListResultPage{}, nil
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
