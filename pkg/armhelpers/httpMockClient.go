// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/armhelpers/testserver"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/Azure/go-autorest/autorest/azure"
)

const (
	subscriptionID                             = "cc6b141e-6afc-4786-9bf6-e3b9a5601460"
	tenantID                                   = "19590a3f-b1af-4e6b-8f63-f917cbf40711"
	resourceGroup                              = "TestResourceGroup"
	computeAPIVersion                          = "2019-07-01"
	diskAPIVersion                             = "2019-07-01"
	networkAPIVersion                          = "2018-08-01"
	deploymentAPIVersion                       = "2018-05-01"
	resourceGroupAPIVersion                    = "2018-05-01"
	logAnalyticsAPIVersion                     = "2015-11-01-preview"
	subscriptionsAPIVersion                    = "2016-06-01"
	deploymentName                             = "testDeplomentName"
	deploymentStatus                           = "08586474508192185203"
	virtualMachineScaleSetName                 = "vmscalesetName"
	virtualMachineAvailabilitySetName          = "vmavailabilitysetName"
	virtualMachineName                         = "testVirtualMachineName"
	logAnalyticsDefaultResourceGroupEUS        = "DefaultResourceGroup-EUS"
	logAnalyticsDefaultWorkspaceNameEUS        = "DefaultWorkspace-cc6b141e-6afc-4786-9bf6-e3b9a5601460-EUS"
	logAnalyticsDefaultResourceGroupWEU        = "DefaultResourceGroup-WEU"
	logAnalyticsDefaultWorkspaceNameWEU        = "DefaultWorkspace-cc6b141e-6afc-4786-9bf6-e3b9a5601460-WEU"
	logAnalyticsDefaultResourceGroupMC         = "DefaultResourceGroup-EAST2"
	logAnalyticsDefaultWorkspaceNameMC         = "DefaultWorkspace-cc6b141e-6afc-4786-9bf6-e3b9a5601460-EAST2"
	logAnalyticsWorkspaceName                  = "testLogAnalyticsWorkspace"
	logAnalyticsSolutionName                   = "ContainerInsights(testLogAnalyticsWorkspace)"
	virtualNicName                             = "testVirtualNicName"
	virutalDiskName                            = "testVirtualdickName"
	location                                   = "local"
	operationID                                = "7184adda-13fc-4d49-b941-fbbc3b08ed64"
	publisher                                  = "DefaultPublisher"
	sku                                        = "DefaultSku"
	offer                                      = "DefaultOffer"
	version                                    = "DefaultVersion"
	filePathTokenResponse                      = "httpMockClientData/tokenResponse.json"
	filePathListLocations                      = "httpMockClientData/listLocations.json"
	filePathListVirtualMachineScaleSets        = "httpMockClientData/listVirtualMachineScaleSets.json"
	filePathListVirtualMachineScaleSetVMs      = "httpMockClientData/listVirtualMachineScaleSetVMs.json"
	filePathListVirtualMachines                = "httpMockClientData/listVirtualMachines.json"
	filePathGetVirtualMachine                  = "httpMockClientData/getVirtualMachine.json"
	fileDeployVirtualMachine                   = "httpMockClientData/deployVMResponse.json"
	fileDeployVirtualMachineError              = "httpMockClientData/deploymentVMError.json"
	filePathGetAvailabilitySet                 = "httpMockClientData/getAvailabilitySet.json"
	filePathGetLogAnalyticsWorkspace           = "httpMockClientData/getLogAnalyticsWorkspace.json"
	filePathGetLogAnalyticsWorkspaceSharedKeys = "httpMockClientData/getLogAnalyticsWorkspaceSharedKeys.json"
	filePathListWorkspacesByResourceGroup      = "httpMockClientData/getListWorkspacesByResourceGroup.json"
	filePathCreateOrUpdateWorkspace            = "httpMockClientData/createOrUpdateWorkspace.json"
	filePathListWorkspacesByResourceGroupInMC  = "httpMockClientData/getListWorkspacesByResourceGroup.json"
	filePathCreateOrUpdateWorkspaceInMC        = "httpMockClientData/createOrUpdateWorkspace.json"
	filePathGetVirtualMachineImage             = "httpMockClientData/getVirtualMachineImage.json"
	filePathListVirtualMachineImages           = "httpMockClientData/listVirtualMachineImages.json"
)

//HTTPMockClient is an wrapper of httpmock
type HTTPMockClient struct {
	SubscriptionID                             string
	TenantID                                   string
	ResourceGroup                              string
	ResourceGroupAPIVersion                    string
	ComputeAPIVersion                          string
	DiskAPIVersion                             string
	NetworkAPIVersion                          string
	DeploymentAPIVersion                       string
	LogAnalyticsAPIVersion                     string
	SubscriptionsAPIVersion                    string
	DeploymentName                             string
	DeploymentStatus                           string
	VirtualMachineScaleSetName                 string
	VirtualMachineName                         string
	LogAnalyticsDefaultResourceGroupEUS        string
	LogAnalyticsDefaultWorkspaceNameEUS        string
	LogAnalyticsDefaultResourceGroupWEU        string
	LogAnalyticsDefaultWorkspaceNameWEU        string
	LogAnalyticsDefaultResourceGroupMC         string
	LogAnalyticsDefaultWorkspaceNameMC         string
	LogAnalyticsWorkspaceName                  string
	LogAnalyticsSolutionName                   string
	VirtualNicName                             string
	VirutalDiskName                            string
	Location                                   string
	OperationID                                string
	TokenResponse                              string
	Publisher                                  string
	Sku                                        string
	Offer                                      string
	Version                                    string
	ResponseListLocations                      string
	ResponseListVirtualMachineScaleSets        string
	ResponseListVirtualMachineScaleSetVMs      string
	ResponseListVirtualMachines                string
	ResponseGetVirtualMachine                  string
	ResponseDeployVirtualMachine               string
	ResponseDeployVirtualMachineError          string
	ResponseGetAvailabilitySet                 string
	ResponseGetLogAnalyticsWorkspace           string
	ResponseGetLogAnalyticsWorkspaceSharedKeys string
	ResponseListWorkspacesByResourceGroup      string
	ResponseCreateOrUpdateWorkspace            string
	ResponseListWorkspacesByResourceGroupInMC  string
	ResponseCreateOrUpdateWorkspaceInMC        string
	ResponseGetVirtualMachineImage             string
	ResponseListVirtualMachineImages           string
	mux                                        *http.ServeMux
	server                                     *testserver.TestServer
}

//VirtualMachineScaleSetListValues is an wrapper of virtual machine scale set list response values
type VirtualMachineScaleSetListValues struct {
	Value []compute.VirtualMachineScaleSet
}

//VirtualMachineScaleSetVMValues is an wrapper of virtual machine scale set VM response values
type VirtualMachineScaleSetVMValues struct {
	Value []compute.VirtualMachineScaleSetVM
}

//VirtualMachineVMValues is an wrapper of virtual machine VM response values
type VirtualMachineVMValues struct {
	Value []compute.VirtualMachine
}

//NewHTTPMockClient creates HTTPMockClient with default values
func NewHTTPMockClient() (HTTPMockClient, error) {

	client := HTTPMockClient{
		SubscriptionID:                      subscriptionID,
		TenantID:                            tenantID,
		ResourceGroup:                       resourceGroup,
		ResourceGroupAPIVersion:             resourceGroupAPIVersion,
		ComputeAPIVersion:                   computeAPIVersion,
		DiskAPIVersion:                      diskAPIVersion,
		LogAnalyticsAPIVersion:              logAnalyticsAPIVersion,
		NetworkAPIVersion:                   networkAPIVersion,
		DeploymentAPIVersion:                deploymentAPIVersion,
		SubscriptionsAPIVersion:             subscriptionsAPIVersion,
		DeploymentName:                      deploymentName,
		DeploymentStatus:                    deploymentStatus,
		VirtualMachineScaleSetName:          virtualMachineScaleSetName,
		VirtualMachineName:                  virtualMachineName,
		LogAnalyticsWorkspaceName:           logAnalyticsWorkspaceName,
		LogAnalyticsDefaultResourceGroupEUS: logAnalyticsDefaultResourceGroupEUS,
		LogAnalyticsDefaultWorkspaceNameEUS: logAnalyticsDefaultWorkspaceNameEUS,
		LogAnalyticsDefaultResourceGroupWEU: logAnalyticsDefaultResourceGroupWEU,
		LogAnalyticsDefaultWorkspaceNameWEU: logAnalyticsDefaultWorkspaceNameWEU,
		LogAnalyticsDefaultResourceGroupMC:  logAnalyticsDefaultResourceGroupMC,
		LogAnalyticsDefaultWorkspaceNameMC:  logAnalyticsDefaultWorkspaceNameMC,
		LogAnalyticsSolutionName:            logAnalyticsSolutionName,
		VirtualNicName:                      virtualNicName,
		VirutalDiskName:                     virutalDiskName,
		Location:                            location,
		OperationID:                         operationID,
		Publisher:                           publisher,
		Offer:                               offer,
		Sku:                                 sku,
		Version:                             version,
		mux:                                 http.NewServeMux(),
	}
	var err error
	client.TokenResponse, err = readFromFile(filePathTokenResponse)
	if err != nil {
		return client, err
	}
	client.ResponseListVirtualMachineScaleSets, err = readFromFile(filePathListVirtualMachineScaleSets)
	if err != nil {
		return client, err
	}
	client.ResponseListVirtualMachineScaleSetVMs, err = readFromFile(filePathListVirtualMachineScaleSetVMs)
	if err != nil {
		return client, err
	}
	client.ResponseListVirtualMachines, err = readFromFile(filePathListVirtualMachines)
	if err != nil {
		return client, err
	}
	client.ResponseGetVirtualMachine, err = readFromFile(filePathGetVirtualMachine)
	if err != nil {
		return client, err
	}
	client.ResponseDeployVirtualMachine, err = readFromFile(fileDeployVirtualMachine)
	if err != nil {
		return client, err
	}
	client.ResponseDeployVirtualMachineError, err = readFromFile(fileDeployVirtualMachineError)
	if err != nil {
		return client, err
	}
	client.ResponseGetAvailabilitySet, err = readFromFile(filePathGetAvailabilitySet)
	if err != nil {
		return client, err
	}
	client.ResponseGetLogAnalyticsWorkspace, err = readFromFile(filePathGetLogAnalyticsWorkspace)
	if err != nil {
		return client, err
	}
	client.ResponseGetLogAnalyticsWorkspaceSharedKeys, err = readFromFile(filePathGetLogAnalyticsWorkspaceSharedKeys)
	if err != nil {
		return client, err
	}
	client.ResponseListWorkspacesByResourceGroup, err = readFromFile(filePathListWorkspacesByResourceGroup)
	if err != nil {
		return client, err
	}
	client.ResponseCreateOrUpdateWorkspace, err = readFromFile(filePathCreateOrUpdateWorkspace)
	if err != nil {
		return client, err
	}

	client.ResponseListWorkspacesByResourceGroupInMC, err = readFromFile(filePathListWorkspacesByResourceGroupInMC)
	if err != nil {
		return client, err
	}
	client.ResponseCreateOrUpdateWorkspaceInMC, err = readFromFile(filePathCreateOrUpdateWorkspaceInMC)
	if err != nil {
		return client, err
	}

	client.ResponseGetVirtualMachineImage, err = readFromFile(filePathGetVirtualMachineImage)
	if err != nil {
		return client, err
	}

	client.ResponseListVirtualMachineImages, err = readFromFile(filePathListVirtualMachineImages)
	if err != nil {
		return client, err
	}

	client.ResponseListLocations, err = readFromFile(filePathListLocations)
	if err != nil {
		return client, err
	}

	return client, nil
}

// Activate starts the mock environment and should only be called
// after all required endpoints have been registered.
func (mc *HTTPMockClient) Activate() error {
	server, err := testserver.CreateAndStart(0, mc.mux)
	if err != nil {
		return err
	}

	mc.server = server
	return nil
}

//DeactivateAndReset shuts down the mock environment and removes any registered mocks
func (mc *HTTPMockClient) DeactivateAndReset() {
	if mc.server != nil {
		mc.server.Stop()
	}

	mc.mux = http.NewServeMux()
	mc.server = nil
}

//GetEnvironment return azure.Environment for Azure Stack
func (mc HTTPMockClient) GetEnvironment() azure.Environment {
	env, _ := azure.EnvironmentFromName("AZUREPUBLICCLOUD")
	env.Name = "AzurePublicCloud"

	if mc.server != nil {
		mockURI := fmt.Sprintf("http://localhost:%d/", mc.server.Port)
		env.ActiveDirectoryEndpoint = mockURI
		env.ResourceManagerEndpoint = mockURI
	}

	return env
}

// RegisterLogin registers the mock response for login
func (mc HTTPMockClient) RegisterLogin() {
	mc.mux.HandleFunc(fmt.Sprintf("/subscriptions/%s", mc.SubscriptionID), func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != "2016-06-01" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.Header().Add("Www-Authenticate", fmt.Sprintf(`Bearer authorization_uri="https://login.windows.net/%s", error="invalid_token", error_description="The authentication failed because of missing 'Authorization' header."`, mc.TenantID))
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = fmt.Fprint(w, `{"error":{"code":"AuthenticationFailed","message":"Authentication failed. The 'Authorization' header is missing."}}`)
		}
	})

	mc.mux.HandleFunc(fmt.Sprintf("/%s/oauth2/token", mc.TenantID), func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != "1.0" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, _ = fmt.Fprint(w, mc.TokenResponse)
		}
	})
}

// RegisterListVirtualMachineScaleSets registers the mock response for ListVirtualMachineScaleSets
func (mc HTTPMockClient) RegisterListVirtualMachineScaleSets() {
	pattern := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachineScaleSets", mc.SubscriptionID, mc.ResourceGroup)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.ComputeAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, _ = fmt.Fprint(w, mc.ResponseListVirtualMachineScaleSets)
		}
	})
}

// RegisterListVirtualMachineScaleSetVMs registers the mock response for ListVirtualMachineScaleSetVMs
func (mc HTTPMockClient) RegisterListVirtualMachineScaleSetVMs() {
	pattern := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachineScaleSets/%s/virtualMachines", mc.SubscriptionID, mc.ResourceGroup, mc.VirtualMachineScaleSetName)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.ComputeAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, _ = fmt.Fprint(w, mc.ResponseListVirtualMachineScaleSetVMs)
		}
	})
}

// RegisterListVirtualMachines registers the mock response for ListVirtualMachines
func (mc HTTPMockClient) RegisterListVirtualMachines() {
	pattern := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachines", mc.SubscriptionID, mc.ResourceGroup)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.ComputeAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, _ = fmt.Fprint(w, mc.ResponseListVirtualMachines)
		}
	})
}

// RegisterListLocations registers the mock response for ListVirtualMachines
func (mc HTTPMockClient) RegisterListLocations() {
	pattern := fmt.Sprintf("/subscriptions/%s/locations", mc.SubscriptionID)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.SubscriptionsAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, _ = fmt.Fprint(w, mc.ResponseListLocations)
		}
	})
}

// RegisterGetAvailabilitySet registers the mock response for GetAvailabilitySet.
func (mc HTTPMockClient) RegisterGetAvailabilitySet() {
	pattern := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/availabilitySets/vmavailabilitysetName", mc.SubscriptionID, mc.ResourceGroup)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.ComputeAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, _ = fmt.Fprint(w, mc.ResponseGetAvailabilitySet)
		}
	})
}

// RegisterGetAvailabilitySetFaultDomainCount registers a mock response for GetAvailabilitySet.
func (mc HTTPMockClient) RegisterGetAvailabilitySetFaultDomainCount() {
	pattern := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/availabilitySets/id1", mc.SubscriptionID, mc.ResourceGroup)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.ComputeAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, _ = fmt.Fprint(w, mc.ResponseGetAvailabilitySet)
		}
	})
}

// RegisterVirtualMachineEndpoint registers mock responses for the Microsoft.Compute/virtualMachines endpoint
func (mc *HTTPMockClient) RegisterVirtualMachineEndpoint() {
	pattern := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachines/%s", mc.SubscriptionID, mc.ResourceGroup, mc.VirtualMachineName)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.ComputeAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			switch r.Method {
			case http.MethodGet:
				_, _ = fmt.Fprint(w, mc.ResponseGetVirtualMachine)
			case http.MethodDelete:
				w.Header().Add("Azure-Asyncoperation", fmt.Sprintf("http://localhost:%d/subscriptions/%s/providers/Microsoft.Compute/locations/%s/operations/%s?api-version=%s", mc.server.Port, mc.SubscriptionID, mc.Location, mc.OperationID, mc.ComputeAPIVersion))
				w.Header().Add("Location", fmt.Sprintf("http://localhost:%d/subscriptions/%s/providers/Microsoft.Compute/locations/%s/operations/%s?monitor=true&api-version=%s", mc.server.Port, mc.SubscriptionID, mc.Location, mc.OperationID, mc.ComputeAPIVersion))
				w.Header().Add("Azure-Asyncnotification", "Enabled")
				w.Header().Add("Content-Length", "0")
				w.WriteHeader(http.StatusAccepted)
			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		}
	})
}

// RegisterDeleteOperation registers mock responses for checking the status of a delete operation
func (mc HTTPMockClient) RegisterDeleteOperation() {
	pattern := fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Compute/locations/%s/operations/%s", mc.SubscriptionID, mc.Location, mc.OperationID)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.ComputeAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else if r.URL.Query().Get("monitor") == "true" {
			w.WriteHeader(http.StatusOK)
		} else {
			_, _ = fmt.Fprintf(w, `
			{
			  "startTime": "2019-03-30T00:23:10.9206154+00:00",
			  "endTime": "2019-03-30T00:23:51.8424926+00:00",
			  "status": "Succeeded",
			  "name": "%s"
			}`, mc.OperationID)
		}
	})
}

// RegisterDeployTemplate registers the mock response for DeployTemplate
func (mc *HTTPMockClient) RegisterDeployTemplate() {
	pattern := fmt.Sprintf("/subscriptions/%s/resourcegroups/%s/providers/Microsoft.Resources/deployments/%s", mc.SubscriptionID, mc.ResourceGroup, mc.DeploymentName)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.DeploymentAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			switch r.Method {
			case http.MethodPut:
				w.Header().Add("Azure-Asyncoperation", fmt.Sprintf("http://localhost:%d/subscriptions/%s/resourcegroups/%s/providers/Microsoft.Resources/deployments/%s/operationStatuses/%s?api-version=%s", mc.server.Port, subscriptionID, resourceGroup, deploymentName, deploymentStatus, deploymentAPIVersion))
				w.WriteHeader(http.StatusCreated)
				_, _ = fmt.Fprint(w, mc.ResponseDeployVirtualMachine)
			case http.MethodGet:
				_, _ = fmt.Fprint(w, mc.ResponseDeployVirtualMachine)
			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		}
	})
}

// RegisterDeployOperationSuccess registers the mock response for a successful deployment
func (mc HTTPMockClient) RegisterDeployOperationSuccess() {
	pattern := fmt.Sprintf("/subscriptions/%s/resourcegroups/%s/providers/Microsoft.Resources/deployments/%s/operationStatuses/%s", mc.SubscriptionID, mc.ResourceGroup, mc.DeploymentName, mc.DeploymentStatus)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.DeploymentAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, _ = fmt.Fprint(w, `{"status": "Succeeded"}`)
		}
	})
}

// RegisterDeployOperationSuccess registers the mock response for a failed deployment
func (mc HTTPMockClient) RegisterDeployOperationFailure() {
	pattern := fmt.Sprintf("/subscriptions/%s/resourcegroups/%s/providers/Microsoft.Resources/deployments/%s/operationStatuses/%s", mc.SubscriptionID, mc.ResourceGroup, mc.DeploymentName, mc.DeploymentStatus)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.DeploymentAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, _ = fmt.Fprint(w, mc.ResponseDeployVirtualMachineError)
		}
	})
}

// RegisterDeleteNetworkInterface registers the mock response for DeleteNetworkInterface
func (mc *HTTPMockClient) RegisterDeleteNetworkInterface() {
	pattern := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/networkInterfaces/%s", mc.SubscriptionID, mc.ResourceGroup, mc.VirtualNicName)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.NetworkAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.Header().Add("Azure-Asyncoperation", fmt.Sprintf("http://localhost:%d/subscriptions/%s/providers/Microsoft.Network/locations/%s/operations/%s?api-version=%s", mc.server.Port, mc.SubscriptionID, mc.Location, mc.OperationID, mc.NetworkAPIVersion))
			w.Header().Add("Location", fmt.Sprintf("http://localhost:%d/subscriptions/%s/providers/Microsoft.Network/locations/%s/operationResults/%s?api-version=%s", mc.server.Port, mc.SubscriptionID, mc.Location, mc.OperationID, mc.NetworkAPIVersion))
			w.Header().Add("Azure-Asyncnotification", "Enabled")
			w.Header().Add("Content-Length", "0")
			w.WriteHeader(http.StatusAccepted)
		}
	})

	pattern = fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Network/locations/%s/operations/%s", mc.SubscriptionID, mc.Location, mc.OperationID)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.NetworkAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, _ = fmt.Fprint(w, `{"status": "Succeeded"}`)
		}
	})

	pattern = fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Network/locations/%s/operationResults/%s", mc.SubscriptionID, mc.Location, mc.OperationID)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.NetworkAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	})
}

// RegisterDeleteManagedDisk registers the mock response for DeleteManagedDisk
func (mc *HTTPMockClient) RegisterDeleteManagedDisk() {
	pattern := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/disks/%s", mc.SubscriptionID, mc.ResourceGroup, mc.VirutalDiskName)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.DiskAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.Header().Add("Azure-Asyncoperation", fmt.Sprintf("http://localhost:%d/subscriptions/%s/providers/Microsoft.Compute/locations/%s/DiskOperations/%s?api-version=%s", mc.server.Port, mc.SubscriptionID, mc.Location, mc.OperationID, mc.ComputeAPIVersion))
			w.Header().Add("Location", fmt.Sprintf("http://localhost:%d/subscriptions/%s/providers/Microsoft.Compute/locations/%s/DiskOperations/%s?monitor=true&api-version=%s", mc.server.Port, mc.SubscriptionID, mc.Location, mc.OperationID, mc.ComputeAPIVersion))
			w.Header().Add("Azure-Asyncnotification", "Enabled")
			w.Header().Add("Content-Length", "0")
			w.WriteHeader(http.StatusAccepted)
		}
	})

	pattern = fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Compute/locations/%s/DiskOperations/%s", mc.SubscriptionID, mc.Location, mc.OperationID)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.ComputeAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else if r.URL.Query().Get("monitor") == "true" {
			w.WriteHeader(http.StatusOK)
		} else {
			_, _ = fmt.Fprintf(w, `
			{
			  "startTime": "2019-03-30T00:23:10.9206154+00:00",
			  "endTime": "2019-03-30T00:23:51.8424926+00:00",
			  "status": "Succeeded",
			  "name": "%s"
			}`, mc.OperationID)
		}
	})
}

// RegisterGetLogAnalyticsWorkspaceInfo registers the mock response for GetLogAnalyticsWorkspaceInfo.
func (mc HTTPMockClient) RegisterGetLogAnalyticsWorkspaceInfo() {
	pattern := fmt.Sprintf("/subscriptions/%s/resourcegroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s", mc.SubscriptionID, mc.ResourceGroup, mc.LogAnalyticsWorkspaceName)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.LogAnalyticsAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, _ = fmt.Fprint(w, mc.ResponseGetLogAnalyticsWorkspace)
		}
	})

	pattern = fmt.Sprintf("/subscriptions/%s/resourcegroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/sharedKeys", mc.SubscriptionID, mc.ResourceGroup, mc.LogAnalyticsWorkspaceName)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.LogAnalyticsAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, _ = fmt.Fprint(w, mc.ResponseGetLogAnalyticsWorkspaceSharedKeys)
		}
	})
}

// RegisterEnsureDefaultLogAnalyticsWorkspace registers the mock response for EnsureDefaultLogAnalyticsWorkspace.
func (mc HTTPMockClient) RegisterEnsureDefaultLogAnalyticsWorkspaceUseExisting() {
	pattern := fmt.Sprintf("/subscriptions/%s/resourcegroups/%s", mc.SubscriptionID, mc.LogAnalyticsDefaultResourceGroupEUS)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.ResourceGroupAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	})

	pattern = fmt.Sprintf("/subscriptions/%s/resourcegroups/%s/providers/Microsoft.OperationalInsights/workspaces", mc.SubscriptionID, mc.LogAnalyticsDefaultResourceGroupEUS)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.LogAnalyticsAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, _ = fmt.Fprint(w, mc.ResponseListWorkspacesByResourceGroup)
		}
	})

	pattern = fmt.Sprintf("/subscriptions/%s/resourcegroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s", mc.SubscriptionID, mc.LogAnalyticsDefaultResourceGroupEUS, mc.LogAnalyticsDefaultWorkspaceNameEUS)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.LogAnalyticsAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, _ = fmt.Fprint(w, mc.ResponseCreateOrUpdateWorkspace)
		}
	})
}

// RegisterEnsureDefaultLogAnalyticsWorkspace registers the mock response for EnsureDefaultLogAnalyticsWorkspace.
func (mc HTTPMockClient) RegisterEnsureDefaultLogAnalyticsWorkspaceCreateNew() {
	pattern := fmt.Sprintf("/subscriptions/%s/resourcegroups/%s", mc.SubscriptionID, mc.LogAnalyticsDefaultResourceGroupWEU)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.ResourceGroupAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	})

	pattern = fmt.Sprintf("/subscriptions/%s/resourcegroups/%s/providers/Microsoft.OperationalInsights/workspaces", mc.SubscriptionID, mc.LogAnalyticsDefaultResourceGroupWEU)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.LogAnalyticsAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, _ = fmt.Fprint(w, mc.ResponseListWorkspacesByResourceGroup)
		}
	})

	pattern = fmt.Sprintf("/subscriptions/%s/resourcegroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s", mc.SubscriptionID, mc.LogAnalyticsDefaultResourceGroupWEU, mc.LogAnalyticsDefaultWorkspaceNameWEU)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.LogAnalyticsAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, _ = fmt.Fprint(w, mc.ResponseCreateOrUpdateWorkspace)
		}
	})

}

// RegisterEnsureDefaultLogAnalyticsWorkspace registers the mock response for EnsureDefaultLogAnalyticsWorkspace.
func (mc HTTPMockClient) RegisterEnsureDefaultLogAnalyticsWorkspaceCreateNewInMC() {
	pattern := fmt.Sprintf("/subscriptions/%s/resourcegroups/%s", mc.SubscriptionID, mc.LogAnalyticsDefaultResourceGroupMC)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.ResourceGroupAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	})

	pattern = fmt.Sprintf("/subscriptions/%s/resourcegroups/%s/providers/Microsoft.OperationalInsights/workspaces", mc.SubscriptionID, mc.LogAnalyticsDefaultResourceGroupMC)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.LogAnalyticsAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, _ = fmt.Fprint(w, mc.ResponseListWorkspacesByResourceGroupInMC)
		}
	})

	pattern = fmt.Sprintf("/subscriptions/%s/resourcegroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s", mc.SubscriptionID, mc.LogAnalyticsDefaultResourceGroupMC, mc.LogAnalyticsDefaultWorkspaceNameMC)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.LogAnalyticsAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, _ = fmt.Fprint(w, mc.ResponseCreateOrUpdateWorkspaceInMC)
		}
	})

}

// RegisterVMImageFetcherInterface registers the mock response for VMImageFetcherInterface methods.
func (mc *HTTPMockClient) RegisterVMImageFetcherInterface() {
	pattern := fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Compute/locations/%s/publishers/%s/artifacttypes/vmimage/offers/%s/skus/%s/versions/%s", mc.SubscriptionID, mc.Location, mc.Publisher, mc.Offer, mc.Sku, mc.Version)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.ComputeAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, _ = fmt.Fprint(w, mc.ResponseGetVirtualMachineImage)
		}
	})

	pattern = fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Compute/locations/%s/publishers/%s/artifacttypes/vmimage/offers/%s/skus/%s/versions/%s", mc.SubscriptionID, mc.Location, api.Ubuntu1604OSImageConfig.ImagePublisher, api.Ubuntu1604OSImageConfig.ImageOffer, api.Ubuntu1604OSImageConfig.ImageSku, api.Ubuntu1604OSImageConfig.ImageVersion)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.ComputeAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, _ = fmt.Fprint(w, mc.ResponseGetVirtualMachineImage)
		}
	})

	pattern = fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Compute/locations/%s/publishers/%s/artifacttypes/vmimage/offers/%s/skus/%s/versions/%s", mc.SubscriptionID, mc.Location, api.WindowsServer2019OSImageConfig.ImagePublisher, api.WindowsServer2019OSImageConfig.ImageOffer, api.WindowsServer2019OSImageConfig.ImageSku, api.WindowsServer2019OSImageConfig.ImageVersion)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.ComputeAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, _ = fmt.Fprint(w, mc.ResponseGetVirtualMachineImage)
		}
	})

	pattern = fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Compute/locations/%s/publishers/%s/artifacttypes/vmimage/offers/%s/skus/%s/versions/%s", mc.SubscriptionID, mc.Location, api.AKSUbuntu1604OSImageConfig.ImagePublisher, api.AKSUbuntu1604OSImageConfig.ImageOffer, api.AKSUbuntu1604OSImageConfig.ImageSku, api.AKSUbuntu1604OSImageConfig.ImageVersion)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.ComputeAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, _ = fmt.Fprint(w, mc.ResponseGetVirtualMachineImage)
		}
	})

	pattern = fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Compute/locations/%s/publishers/%s/artifacttypes/vmimage/offers/%s/skus/%s/versions", mc.SubscriptionID, mc.Location, mc.Publisher, mc.Offer, mc.Sku)
	mc.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("api-version") != mc.ComputeAPIVersion {
			w.WriteHeader(http.StatusNotFound)
		} else {
			_, _ = fmt.Fprint(w, mc.ResponseListVirtualMachineImages)
		}
	})
}

func readFromFile(filePath string) (string, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("Fail to read file %q , err -  %q", filePath, err)
	}

	return string(bytes), nil
}
