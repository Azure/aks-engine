// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package azurestack

import (
	"context"

	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	// AADContributorRoleID is the role id that exists in every subscription for 'Contributor'
	AADContributorRoleID = "b24988ac-6180-42a0-ab88-20f7382dd24c"
	// AADRoleReferenceTemplate is a template for a roleDefinitionId
	AADRoleReferenceTemplate = "/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/%s"
	// AADRoleResourceGroupScopeTemplate is a template for a roleDefinition scope
	AADRoleResourceGroupScopeTemplate = "/subscriptions/%s/resourceGroups/%s"
)

// CreateGraphApplication creates an application via the graphrbac client
func (az *AzureClient) CreateGraphApplication(ctx context.Context, applicationCreateParameters graphrbac.ApplicationCreateParameters) (graphrbac.Application, error) {
	errorMessage := "error azure stack does not support creating application"
	log.Fatal(errorMessage)
	return graphrbac.Application{}, errors.New(errorMessage)
}

// DeleteGraphApplication deletes an application via the graphrbac client
func (az *AzureClient) DeleteGraphApplication(ctx context.Context, applicationObjectID string) (result autorest.Response, err error) {
	errorMessage := "error azure stack does not support deleting application"
	log.Fatal(errorMessage)
	return autorest.Response{}, errors.New(errorMessage)
}

// CreateGraphPrincipal creates a service principal via the graphrbac client
func (az *AzureClient) CreateGraphPrincipal(ctx context.Context, servicePrincipalCreateParameters graphrbac.ServicePrincipalCreateParameters) (graphrbac.ServicePrincipal, error) {
	errorMessage := "error azure stack does not support creating service principal"
	log.Fatal(errorMessage)
	return graphrbac.ServicePrincipal{}, errors.New(errorMessage)
}

// CreateRoleAssignment creates a role assignment via the authorization client
func (az *AzureClient) CreateRoleAssignment(ctx context.Context, scope string, roleAssignmentName string, parameters authorization.RoleAssignmentCreateParameters) (authorization.RoleAssignment, error) {
	errorMessage := "error azure stack does not support creating role assignement"
	log.Fatal(errorMessage)
	return authorization.RoleAssignment{}, errors.New(errorMessage)
}

// DeleteRoleAssignmentByID deletes a roleAssignment via its unique identifier
func (az *AzureClient) DeleteRoleAssignmentByID(ctx context.Context, roleAssignmentID string) (authorization.RoleAssignment, error) {
	errorMessage := "error azure stack does not support deleting role assignement"
	log.Fatal(errorMessage)
	return authorization.RoleAssignment{}, errors.New(errorMessage)
}

// ListRoleAssignmentsForPrincipal (e.g. a VM) via the scope and the unique identifier of the principal
func (az *AzureClient) ListRoleAssignmentsForPrincipal(ctx context.Context, scope string, principalID string) (armhelpers.RoleAssignmentListResultPage, error) {
	errorMessage := "error azure stack does not support listing role assignement"
	log.Fatal(errorMessage)
	return nil, errors.New(errorMessage)
}

// CreateApp is a simpler method for creating an application
func (az *AzureClient) CreateApp(ctx context.Context, appName, appURL string, replyURLs *[]string, requiredResourceAccess *[]graphrbac.RequiredResourceAccess) (applicationResp graphrbac.Application, servicePrincipalObjectID, servicePrincipalClientSecret string, err error) {
	errorMessage := "error azure stack does not support creating application"
	log.Fatal(errorMessage)
	return graphrbac.Application{}, "", "", errors.New(errorMessage)
}

// DeleteApp is a simpler method for deleting an application and the associated spn
func (az *AzureClient) DeleteApp(ctx context.Context, applicationName, applicationObjectID string) (autorest.Response, error) {
	errorMessage := "error azure stack does not support deleting application"
	log.Fatal(errorMessage)
	return autorest.Response{}, errors.New(errorMessage)
}

// CreateRoleAssignmentSimple is a wrapper around RoleAssignmentsClient.Create
func (az *AzureClient) CreateRoleAssignmentSimple(ctx context.Context, resourceGroup, servicePrincipalObjectID string) error {
	errorMessage := "error azure stack does not support creating role assignment"
	log.Fatal(errorMessage)
	return errors.New(errorMessage)
}
