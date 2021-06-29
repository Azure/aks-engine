// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/google/uuid"
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
	return az.applicationsClient.Create(ctx, applicationCreateParameters)
}

// DeleteGraphApplication deletes an application via the graphrbac client
func (az *AzureClient) DeleteGraphApplication(ctx context.Context, applicationObjectID string) (result autorest.Response, err error) {
	return az.applicationsClient.Delete(ctx, applicationObjectID)
}

// CreateGraphPrincipal creates a service principal via the graphrbac client
func (az *AzureClient) CreateGraphPrincipal(ctx context.Context, servicePrincipalCreateParameters graphrbac.ServicePrincipalCreateParameters) (graphrbac.ServicePrincipal, error) {
	return az.servicePrincipalsClient.Create(ctx, servicePrincipalCreateParameters)
}

// CreateRoleAssignment creates a role assignment via the authorization client
func (az *AzureClient) CreateRoleAssignment(ctx context.Context, scope string, roleAssignmentName string, parameters authorization.RoleAssignmentCreateParameters) (authorization.RoleAssignment, error) {
	return az.authorizationClient.Create(ctx, scope, roleAssignmentName, parameters)
}

// DeleteRoleAssignmentByID deletes a roleAssignment via its unique identifier
func (az *AzureClient) DeleteRoleAssignmentByID(ctx context.Context, roleAssignmentID string) (authorization.RoleAssignment, error) {
	return az.authorizationClient.DeleteByID(ctx, roleAssignmentID)
}

// ListRoleAssignmentsForPrincipal (e.g. a VM) via the scope and the unique identifier of the principal
func (az *AzureClient) ListRoleAssignmentsForPrincipal(ctx context.Context, scope string, principalID string) (RoleAssignmentListResultPage, error) {
	page, err := az.authorizationClient.ListForScope(ctx, scope, fmt.Sprintf("principalId eq '%s'", principalID))
	return &page, err
}

// CreateApp is a simpler method for creating an application
func (az *AzureClient) CreateApp(ctx context.Context, appName, appURL string, replyURLs *[]string, requiredResourceAccess *[]graphrbac.RequiredResourceAccess) (applicationResp graphrbac.Application, servicePrincipalObjectID, servicePrincipalClientSecret string, err error) {
	notBefore := time.Now()
	notAfter := time.Now().Add(10000 * 24 * time.Hour)

	startDate := date.Time{Time: notBefore}
	endDate := date.Time{Time: notAfter}

	servicePrincipalClientSecret = uuid.Must(uuid.NewRandom()).String()

	log.Debugf("ad: creating application with name=%q identifierURL=%q", appName, appURL)
	applicationReq := graphrbac.ApplicationCreateParameters{
		AvailableToOtherTenants: to.BoolPtr(false),
		DisplayName:             to.StringPtr(appName),
		Homepage:                to.StringPtr(appURL),
		IdentifierUris:          to.StringSlicePtr([]string{appURL}),
		ReplyUrls:               replyURLs,
		PasswordCredentials: &[]graphrbac.PasswordCredential{
			{
				KeyID:     to.StringPtr(uuid.Must(uuid.NewRandom()).String()),
				StartDate: &startDate,
				EndDate:   &endDate,
				Value:     to.StringPtr(servicePrincipalClientSecret),
			},
		},
		RequiredResourceAccess: requiredResourceAccess,
	}
	applicationResp, err = az.CreateGraphApplication(ctx, applicationReq)
	if err != nil {
		return applicationResp, "", "", err
	}
	applicationID := to.String(applicationResp.AppID)

	log.Debugf("ad: creating servicePrincipal for applicationID: %q", applicationID)

	servicePrincipalReq := graphrbac.ServicePrincipalCreateParameters{
		AppID:          applicationResp.AppID,
		AccountEnabled: to.BoolPtr(true),
	}
	servicePrincipalResp, err := az.servicePrincipalsClient.Create(ctx, servicePrincipalReq)
	if err != nil {
		return applicationResp, "", "", err
	}

	servicePrincipalObjectID = to.String(servicePrincipalResp.ObjectID)

	return applicationResp, servicePrincipalObjectID, servicePrincipalClientSecret, nil
}

// DeleteApp is a simpler method for deleting an application and the associated spn
func (az *AzureClient) DeleteApp(ctx context.Context, applicationName, applicationObjectID string) (autorest.Response, error) {
	log.Debugf("ad: deleting application with name=%q", applicationName)
	return az.DeleteGraphApplication(ctx, applicationObjectID)
}

// CreateRoleAssignmentSimple is a wrapper around RoleAssignmentsClient.Create
func (az *AzureClient) CreateRoleAssignmentSimple(ctx context.Context, resourceGroup, servicePrincipalObjectID string) error {
	roleAssignmentName := uuid.Must(uuid.NewRandom()).String()

	roleDefinitionID := fmt.Sprintf(AADRoleReferenceTemplate, az.subscriptionID, AADContributorRoleID)
	scope := fmt.Sprintf(AADRoleResourceGroupScopeTemplate, az.subscriptionID, resourceGroup)

	roleAssignmentParameters := authorization.RoleAssignmentCreateParameters{
		Properties: &authorization.RoleAssignmentProperties{
			RoleDefinitionID: to.StringPtr(roleDefinitionID),
			PrincipalID:      to.StringPtr(servicePrincipalObjectID),
		},
	}

	re := regexp.MustCompile(`(?i)status=(\d+)`)
	for {
		_, err := az.CreateRoleAssignment(
			ctx,
			scope,
			roleAssignmentName,
			roleAssignmentParameters,
		)
		if err != nil {
			match := re.FindStringSubmatch(err.Error())
			if match != nil && (match[1] == "403") {
				//insufficient permissions. stop now
				log.Debugf("Failed to create role assignment (will abort now): %q", err)
				return err
			}
			// TODO: Should we handle 409 errors as well here ?
			log.Debugf("Failed to create role assignment (will retry): %q", err)
			time.Sleep(3 * time.Second)
			continue
		}
		break
	}

	return nil
}
