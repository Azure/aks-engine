// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package azurestack

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2017-03-30/compute"
	"github.com/Azure/go-autorest/autorest/azure"
	"gopkg.in/jarcoal/httpmock.v1"
)

const (
	subscriptionID                        = "cc6b141e-6afc-4786-9bf6-e3b9a5601460"
	tenantID                              = "19590a3f-b1af-4e6b-8f63-f917cbf40711"
	resourceGroup                         = "TestResourceGroup"
	computeAPIVersion                     = "2017-03-30"
	networkAPIVersion                     = "2017-10-01"
	deploymentAPIVersion                  = "2018-05-01"
	deploymentName                        = "testDeplomentName"
	deploymentStatus                      = "08586474508192185203"
	virtualMachineScaleSetName            = "vmscalesetName"
	virtualMachineName                    = "testVirtualMachineName"
	virtualNicName                        = "testVirtualNicName"
	virutalDiskName                       = "testVirtualdickName"
	location                              = "local"
	operationID                           = "7184adda-13fc-4d49-b941-fbbc3b08ed64"
	filePathTokenResponse                 = "httpMockClientData/tokenResponse.json"
	filePathListVirtualMachineScaleSets   = "httpMockClientData/listVirtualMachineScaleSets.json"
	filePathListVirtualMachineScaleSetVMs = "httpMockClientData/listVirtualMachineScaleSetVMs.json"
	filePathListVirtualMachines           = "httpMockClientData/listVirtualMachines.json"
	filePathGetVirtualMachine             = "httpMockClientData/getVirtualMachine.json"
	fileDeployVirtualMachine              = "httpMockClientData/deployVMResponse.json"
	fileDeployVirtualMachineError         = "httpMockClientData/deploymentVMError.json"
)

//HTTPMockClient is an wrapper of httpmock
type HTTPMockClient struct {
	SubscriptionID                        string
	TenantID                              string
	ResourceGroup                         string
	ComputeAPIVersion                     string
	NetworkAPIVersion                     string
	DeploymentAPIVersion                  string
	DeploymentName                        string
	DeploymentStatus                      string
	VirtualMachineScaleSetName            string
	VirtualMachineName                    string
	VirtualNicName                        string
	VirutalDiskName                       string
	Location                              string
	OperationID                           string
	TokenResponse                         string
	ResponseListVirtualMachineScaleSets   string
	ResponseListVirtualMachineScaleSetVMs string
	ResponseListVirtualMachines           string
	ResponseGetVirtualMachine             string
	ResponseDeployVirtualMachine          string
	ResponseDeployVirtualMachineError     string
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
		SubscriptionID:             subscriptionID,
		TenantID:                   tenantID,
		ResourceGroup:              resourceGroup,
		ComputeAPIVersion:          computeAPIVersion,
		NetworkAPIVersion:          networkAPIVersion,
		DeploymentAPIVersion:       deploymentAPIVersion,
		DeploymentName:             deploymentName,
		DeploymentStatus:           deploymentStatus,
		VirtualMachineScaleSetName: virtualMachineScaleSetName,
		VirtualMachineName:         virtualMachineName,
		VirtualNicName:             virtualNicName,
		VirutalDiskName:            virutalDiskName,
		Location:                   location,
		OperationID:                operationID,
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
	return client, nil
}

//Activate starts the mock environment
func (mc HTTPMockClient) Activate() {
	httpmock.Activate()
}

//DeactivateAndReset shuts down the mock environment and removes any registered mocks
func (mc HTTPMockClient) DeactivateAndReset() {
	httpmock.DeactivateAndReset()
}

//GetEnvironment return azure.Environment for Azure Stack
func (mc HTTPMockClient) GetEnvironment() azure.Environment {
	env, _ := azure.EnvironmentFromName("AZUREPUBLICCLOUD")
	env.Name = "AzureStackCloud"
	return env
}

// RegisterLogin registers the mock response for login
func (mc HTTPMockClient) RegisterLogin() {
	httpmock.RegisterResponder("GET", fmt.Sprintf("https://management.azure.com/subscriptions/%s?api-version=2016-06-01", mc.SubscriptionID),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(401, `{"error":{"code":"AuthenticationFailed","message":"Authentication failed. The 'Authorization' header is missing."}}`)
			resp.Header.Add("Www-Authenticate", fmt.Sprintf(`Bearer authorization_uri="https://login.windows.net/%s", error="invalid_token", error_description="The authentication failed because of missing 'Authorization' header."`, mc.TenantID))
			return resp, nil
		},
	)

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/token?api-version=1.0", mc.TenantID),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mc.TokenResponse)
			return resp, nil
		},
	)
}

// RegisterListVirtualMachineScaleSets registers the mock response for ListVirtualMachineScaleSets
func (mc HTTPMockClient) RegisterListVirtualMachineScaleSets() {
	httpmock.RegisterResponder("GET", fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachineScaleSets?api-version=%s", mc.SubscriptionID, mc.ResourceGroup, mc.ComputeAPIVersion),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mc.ResponseListVirtualMachineScaleSets)
			return resp, nil
		},
	)
}

// RegisterListVirtualMachineScaleSetVMs registers the mock response for ListVirtualMachineScaleSetVMs
func (mc HTTPMockClient) RegisterListVirtualMachineScaleSetVMs() {
	httpmock.RegisterResponder("GET", fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachineScaleSets/%s/virtualMachines?api-version=%s", mc.SubscriptionID, mc.ResourceGroup, mc.VirtualMachineScaleSetName, mc.ComputeAPIVersion),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mc.ResponseListVirtualMachineScaleSetVMs)
			return resp, nil
		},
	)
}

// RegisterListVirtualMachines registers the mock response for ListVirtualMachines
func (mc HTTPMockClient) RegisterListVirtualMachines() {
	httpmock.RegisterResponder("GET", fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachines?api-version=%s", mc.SubscriptionID, mc.ResourceGroup, mc.ComputeAPIVersion),

		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mc.ResponseListVirtualMachines)
			return resp, nil
		},
	)
}

// RegisterGetVirtualMachine registers the mock response for GetVirtualMachine
func (mc HTTPMockClient) RegisterGetVirtualMachine() {
	httpmock.RegisterResponder("GET", fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachines/%s?api-version=%s", mc.SubscriptionID, mc.ResourceGroup, mc.VirtualMachineName, mc.ComputeAPIVersion),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mc.ResponseGetVirtualMachine)
			return resp, nil
		},
	)
}

// RegisterDeleteVirtualMachine registers the mock response for DeleteVirtualMachine
func (mc HTTPMockClient) RegisterDeleteVirtualMachine() {
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachines/%s?api-version=%s", mc.SubscriptionID, mc.ResourceGroup, mc.VirtualMachineName, mc.ComputeAPIVersion),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(202, "")
			resp.Header.Add("Azure-Asyncoperation", fmt.Sprintf("https://management.azure.com/subscriptions/%s/providers/Microsoft.Compute/locations/%s/operations/%s?api-version=%s", mc.SubscriptionID, mc.Location, mc.OperationID, mc.ComputeAPIVersion))
			resp.Header.Add("Location", fmt.Sprintf("https://management.azure.com/subscriptions/%s/providers/Microsoft.Compute/locations/%s/operations/%s?monitor=true&api-version=%s", mc.SubscriptionID, mc.Location, mc.OperationID, mc.ComputeAPIVersion))
			resp.Header.Add("Azure-Asyncnotification", "Enabled")
			resp.Header.Add("Content-Length", "0")
			resp.Request = req
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://management.azure.com/subscriptions/%s/providers/Microsoft.Compute/locations/%s/operations/%s?api-version=%s", mc.SubscriptionID, mc.Location, mc.OperationID, mc.ComputeAPIVersion),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fmt.Sprintf(`{
			"startTime": "2019-03-30T00:23:10.9206154+00:00",
			"endTime": "2019-03-30T00:23:51.8424926+00:00",
			"status": "Succeeded",
			"name": "%s"
		  }`, mc.OperationID))
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://management.azure.com/subscriptions/%s/providers/Microsoft.Compute/locations/%s/operations/%s?monitor=true&api-version=%s", mc.SubscriptionID, mc.Location, mc.OperationID, mc.ComputeAPIVersion),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, "")
			return resp, nil
		},
	)
}

// RegisterDeployTemplate registers the mock response for DeployTemplate
func (mc HTTPMockClient) RegisterDeployTemplate() {
	httpmock.RegisterResponder("PUT", fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourcegroups/%s/providers/Microsoft.Resources/deployments/%s?api-version=%s", subscriptionID, resourceGroup, deploymentName, deploymentAPIVersion),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(201, mc.ResponseDeployVirtualMachine)
			resp.Header.Add("Azure-Asyncoperation", fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourcegroups/%s/providers/Microsoft.Resources/deployments/%s/operationStatuses/%s?api-version=%s", subscriptionID, resourceGroup, deploymentName, deploymentStatus, deploymentAPIVersion))
			resp.Request = req
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourcegroups/%s/providers/Microsoft.Resources/deployments/%s?api-version=%s", subscriptionID, resourceGroup, deploymentName, deploymentAPIVersion),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mc.ResponseDeployVirtualMachine)
			resp.Request = req
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourcegroups/%s/providers/Microsoft.Resources/deployments/%s/operationStatuses/%s?api-version=%s", subscriptionID, resourceGroup, deploymentName, deploymentStatus, deploymentAPIVersion),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `{
			"status": "Succeeded"
		  }`)
			resp.Request = req
			return resp, nil
		},
	)
}

// RegisterDeployTemplateSync registers the mock response for DeployTemplate
func (mc HTTPMockClient) RegisterDeployTemplateSync() {

	httpmock.RegisterResponder("PUT", fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourcegroups/%s/providers/Microsoft.Resources/deployments/%s?api-version=%s", subscriptionID, resourceGroup, deploymentName, deploymentAPIVersion),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(201, mc.ResponseDeployVirtualMachine)
			resp.Header.Add("Azure-Asyncoperation", fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourcegroups/%s/providers/Microsoft.Resources/deployments/%s/operationStatuses/%s?api-version=%s", subscriptionID, resourceGroup, deploymentName, deploymentStatus, deploymentAPIVersion))
			resp.Request = req
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourcegroups/%s/providers/Microsoft.Resources/deployments/%s/operationStatuses/%s?api-version=%s", subscriptionID, resourceGroup, deploymentName, deploymentStatus, deploymentAPIVersion),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mc.ResponseDeployVirtualMachineError)
			resp.Request = req
			return resp, nil
		},
	)
}

// RegisterDeleteNetworkInterface registers the mock response for DeleteNetworkInterface
func (mc HTTPMockClient) RegisterDeleteNetworkInterface() {

	httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/networkInterfaces/%s?api-version=%s", mc.SubscriptionID, mc.ResourceGroup, mc.VirtualNicName, mc.NetworkAPIVersion),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(202, "")
			resp.Header.Add("Azure-Asyncoperation", fmt.Sprintf("https://management.azure.com/subscriptions/%s/providers/Microsoft.Network/locations/%s/operations/%s?api-version=%s", mc.SubscriptionID, mc.Location, mc.OperationID, mc.NetworkAPIVersion))
			resp.Header.Add("Location", fmt.Sprintf("https://management.azure.com/subscriptions/%s/providers/Microsoft.Network/locations/%s/operationResults/%s?api-version=%s", mc.SubscriptionID, mc.Location, mc.OperationID, mc.NetworkAPIVersion))
			resp.Header.Add("Azure-Asyncnotification", "Enabled")
			resp.Header.Add("Content-Length", "0")
			resp.Request = req
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://management.azure.com/subscriptions/%s/providers/Microsoft.Network/locations/%s/operations/%s?api-version=%s", mc.SubscriptionID, mc.Location, mc.OperationID, mc.NetworkAPIVersion),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `{
			"status": "Succeeded"
		  }`)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://management.azure.com/subscriptions/%s/providers/Microsoft.Network/locations/%s/operationResults/%s?api-version=%s", mc.SubscriptionID, mc.Location, mc.OperationID, mc.NetworkAPIVersion),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, "")
			return resp, nil
		},
	)
}

// RegisterDeleteManagedDisk registers the mock response for DeleteManagedDisk
func (mc HTTPMockClient) RegisterDeleteManagedDisk() {
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/disks/%s?api-version=%s", mc.SubscriptionID, resourceGroup, virutalDiskName, mc.ComputeAPIVersion),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(202, "")
			resp.Header.Add("Azure-Asyncoperation", fmt.Sprintf("https://management.azure.com/subscriptions/%s/providers/Microsoft.Compute/locations/%s/DiskOperations/%s?api-version=%s", mc.SubscriptionID, mc.Location, mc.OperationID, mc.ComputeAPIVersion))
			resp.Header.Add("Location", fmt.Sprintf("https://management.azure.com/subscriptions/%s/providers/Microsoft.Compute/locations/%s/DiskOperations/%s?monitor=true&api-version=%s", mc.SubscriptionID, mc.Location, mc.OperationID, mc.ComputeAPIVersion))
			resp.Header.Add("Azure-Asyncnotification", "Enabled")
			resp.Header.Add("Content-Length", "0")
			resp.Request = req
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://management.azure.com/subscriptions/%s/providers/Microsoft.Compute/locations/%s/DiskOperations/%s?api-version=%s", mc.SubscriptionID, mc.Location, mc.OperationID, mc.ComputeAPIVersion),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, fmt.Sprintf(`{
				"startTime": "2019-03-30T00:23:10.9206154+00:00",
				"endTime": "2019-03-30T00:23:51.8424926+00:00",
				"status": "Succeeded",
				"name": "%s"
			  }`, mc.OperationID))
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://management.azure.com/subscriptions/%s/providers/Microsoft.Compute/locations/%s/DiskOperations/%s?monitor=true&api-version=%s", mc.SubscriptionID, mc.Location, mc.OperationID, mc.ComputeAPIVersion),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, "")
			return resp, nil
		},
	)
}

func readFromFile(filePath string) (string, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("Fail to read file %q , err -  %q", filePath, err)
	}

	return string(bytes), nil
}
