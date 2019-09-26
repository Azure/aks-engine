// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package azurestack

import (
	"context"
	"fmt"
	"strings"

	oi "github.com/Azure/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights"
	om "github.com/Azure/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement"
	"github.com/Azure/go-autorest/autorest/to"
)

// EnsureDefaultLogAnalyticsWorkspace ensures the default log analytics exists corresponding to specified location
func (az *AzureClient) EnsureDefaultLogAnalyticsWorkspace(ctx context.Context, resourceGroup, location string) (workspaceResourceID string, err error) {
	AzureCloudLocationToOmsRegionCodeMap := map[string]string{
		"australiasoutheast": "ASE",
		"australiaeast":      "EAU",
		"australiacentral":   "CAU",
		"canadacentral":      "CCA",
		"centralindia":       "CIN",
		"centralus":          "CUS",
		"eastasia":           "EA",
		"eastus":             "EUS",
		"eastus2":            "EUS2",
		"eastus2euap":        "EAP",
		"francecentral":      "PAR",
		"japaneast":          "EJP",
		"koreacentral":       "SE",
		"northeurope":        "NEU",
		"southcentralus":     "SCUS",
		"southeastasia":      "SEA",
		"uksouth":            "SUK",
		"usgovvirginia":      "USGV",
		"westcentralus":      "EUS",
		"westeurope":         "WEU",
		"westus":             "WUS",
		"westus2":            "WUS2",
	}

	AzureCloudRegionToOmsRegionMap := map[string]string{
		"australiacentral":   "australiacentral",
		"australiacentral2":  "australiacentral",
		"australiaeast":      "australiaeast",
		"australiasoutheast": "australiasoutheast",
		"brazilsouth":        "southcentralus",
		"canadacentral":      "canadacentral",
		"canadaeast":         "canadacentral",
		"centralus":          "centralus",
		"centralindia":       "centralindia",
		"eastasia":           "eastasia",
		"eastus":             "eastus",
		"eastus2":            "eastus2",
		"francecentral":      "francecentral",
		"francesouth":        "francecentral",
		"japaneast":          "japaneast",
		"japanwest":          "japaneast",
		"koreacentral":       "koreacentral",
		"koreasouth":         "koreacentral",
		"northcentralus":     "eastus",
		"northeurope":        "northeurope",
		"southafricanorth":   "westeurope",
		"southafricawest":    "westeurope",
		"southcentralus":     "southcentralus",
		"southeastasia":      "southeastasia",
		"southindia":         "centralindia",
		"uksouth":            "uksouth",
		"ukwest":             "uksouth",
		"westcentralus":      "eastus",
		"westeurope":         "westeurope",
		"westindia":          "centralindia",
		"westus":             "westus",
		"westus2":            "westus2",
	}

	defaultWorkspaceRegion := "eastus"
	defaultWorkspaceRegionCode := "EUS"

	if region, found := AzureCloudRegionToOmsRegionMap[location]; found {
		defaultWorkspaceRegion = region
	}

	if regionCode, found := AzureCloudLocationToOmsRegionCodeMap[defaultWorkspaceRegion]; found {
		defaultWorkspaceRegionCode = regionCode
	}

	defaultWorkspaceResourceGroup := fmt.Sprintf("DefaultResourceGroup-%s", defaultWorkspaceRegionCode)
	defaultWorkspaceName := fmt.Sprintf("DefaultWorkspace-%s-%s", az.subscriptionID, defaultWorkspaceRegionCode)
	resp, err := az.CheckResourceGroupExistence(ctx, defaultWorkspaceResourceGroup)
	if err != nil {
		return "", err
	}
	if resp.StatusCode == 404 {
		_, err = az.EnsureResourceGroup(ctx, defaultWorkspaceResourceGroup, defaultWorkspaceRegion, nil)
		if err != nil {
			return "", err
		}
	}

	wsList, err := az.workspacesClient.ListByResourceGroup(ctx, defaultWorkspaceResourceGroup)
	if err != nil {
		if wsList.Response.StatusCode != 404 {
			return "", err
		}
	}

	if wsList.Value != nil {
		for _, ws := range *wsList.Value {
			if strings.EqualFold(defaultWorkspaceName, *ws.Name) {
				return *ws.ID, nil
			}
		}
	}

	WorkspaceParameters := oi.Workspace{
		Location: &defaultWorkspaceRegion,
		WorkspaceProperties: &oi.WorkspaceProperties{
			Sku: &oi.Sku{
				Name: oi.Standalone,
			},
		},
	}
	future, err := az.workspacesClient.CreateOrUpdate(ctx, defaultWorkspaceResourceGroup, defaultWorkspaceName, WorkspaceParameters)
	if err != nil {
		return "", err
	}
	err = future.WaitForCompletionRef(ctx, az.workspacesClient.Client)
	if err != nil {
		return "", err
	}
	ws, err := future.Result(az.workspacesClient)
	if err != nil {
		return "", err
	}

	return *ws.ID, nil
}

// GetLogAnalyticsWorkspaceInfo gets the details about the workspace
func (az *AzureClient) GetLogAnalyticsWorkspaceInfo(ctx context.Context, workspaceSubscriptionID, workspaceResourceGroup, workspaceName string) (workspaceID string, workspaceKey string, workspaceLocation string, err error) {
	if !strings.EqualFold(workspaceSubscriptionID, az.subscriptionID) {
		az.workspacesClient = oi.NewWorkspacesClientWithBaseURI(az.environment.ResourceManagerEndpoint, workspaceSubscriptionID)
		az.workspacesClient.Authorizer = az.authorizationClient.Authorizer
	}

	resp, err := az.workspacesClient.Get(ctx, workspaceResourceGroup, workspaceName)
	if err != nil {
		return "", "", "", err
	}

	workspaceID = *resp.WorkspaceProperties.CustomerID
	workspaceLocation = *resp.Location
	result, err := az.workspacesClient.GetSharedKeys(ctx, workspaceResourceGroup, workspaceName)
	if err != nil {
		return "", "", "", err
	}
	return workspaceID, *result.PrimarySharedKey, workspaceLocation, nil
}

// AddContainerInsightsSolution adds container insights solution for the specified log analytics workspace
func (az *AzureClient) AddContainerInsightsSolution(ctx context.Context, workspaceSubscriptionID, workspaceResourceGroup, workspaceName, workspaceLocation string) (result bool, err error) {
	solutionClient := om.NewSolutionsClientWithBaseURI(az.environment.ResourceManagerEndpoint, workspaceSubscriptionID, "Microsoft.OperationalInsights", "workspaces", workspaceName)
	solutionClient.Authorizer = az.workspacesClient.Authorizer

	solutionName := "ContainerInsights(" + workspaceName + ")"
	workspaceResourceID := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s", workspaceSubscriptionID, workspaceResourceGroup, workspaceName)
	status, err := solutionClient.CreateOrUpdate(ctx, workspaceResourceGroup, solutionName, om.Solution{
		Name:     to.StringPtr(solutionName),
		Type:     to.StringPtr("Microsoft.OperationsManagement/solutions"),
		Location: to.StringPtr(workspaceLocation),
		Plan: &om.SolutionPlan{
			Name:          to.StringPtr(solutionName),
			Publisher:     to.StringPtr("Microsoft"),
			PromotionCode: to.StringPtr(""),
			Product:       to.StringPtr("OMSGallery/ContainerInsights"),
		},
		Properties: &om.SolutionProperties{
			WorkspaceResourceID: to.StringPtr(workspaceResourceID),
		},
	})
	if err != nil {
		return false, err
	}

	err = status.Future.WaitForCompletionRef(ctx, solutionClient.Client)
	if err != nil {
		return false, err
	}

	solution, err := status.Result(solutionClient)
	if err != nil {
		return false, err
	}

	return strings.EqualFold(*solution.Name, solutionName), nil
}
