// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func TestDeployTemplateSync_Errors(t *testing.T) {
	cases := []struct {
		name                     string
		mockClientFactory        func() *MockAKSEngineClient
		topErrorMatcher          types.GomegaMatcher
		provisioningStateMatcher types.GomegaMatcher
		statusCodeMatcher        types.GomegaMatcher
		responseMatcher          types.GomegaMatcher
		opsListLenMatcher        types.GomegaMatcher
	}{
		{
			name: "InternalOperationError",
			mockClientFactory: func() *MockAKSEngineClient {
				return &MockAKSEngineClient{
					FailDeployTemplate: true,
				}
			},
			topErrorMatcher:          Equal("DeployTemplate failed"),
			provisioningStateMatcher: Equal(""),
			statusCodeMatcher:        Equal(0),
			responseMatcher:          BeEmpty(),
			opsListLenMatcher:        Equal(0),
		},
		{
			name: "QuotaExceeded",
			mockClientFactory: func() *MockAKSEngineClient {
				return &MockAKSEngineClient{
					FailDeployTemplateQuota: true,
				}
			},
			topErrorMatcher:          ContainSubstring("resources.DeploymentsClient#CreateOrUpdate: Failure responding"),
			provisioningStateMatcher: Equal(""),
			statusCodeMatcher:        Equal(400),
			responseMatcher:          ContainSubstring("\"code\":\"QuotaExceeded\""),
			opsListLenMatcher:        Equal(0),
		},
		{
			name: "Conflict",
			mockClientFactory: func() *MockAKSEngineClient {
				return &MockAKSEngineClient{
					FailDeployTemplateConflict: true,
				}
			},
			topErrorMatcher:          ContainSubstring("At least one resource deployment operation failed."),
			provisioningStateMatcher: Equal(""),
			statusCodeMatcher:        Equal(200),
			responseMatcher:          ContainSubstring("\"code\":\"Conflict\""),
			opsListLenMatcher:        Equal(0),
		},
		{
			name: "DeployErrorWithOperationsLists",
			mockClientFactory: func() *MockAKSEngineClient {
				return &MockAKSEngineClient{
					FailDeployTemplateWithProperties: true,
				}
			},
			topErrorMatcher:          ContainSubstring("At least one resource deployment operation failed."),
			provisioningStateMatcher: Equal("Failed"),
			statusCodeMatcher:        Equal(200),
			responseMatcher:          ContainSubstring("\"code\":\"Conflict\""),
			opsListLenMatcher:        Equal(2),
		},
	}

	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			client := c.mockClientFactory()
			logger := log.NewEntry(log.New())
			err := DeployTemplateSync(client, logger, "rg1", "agentvm", map[string]interface{}{}, map[string]interface{}{})
			g := NewGomegaWithT(t)
			g.Expect(err).NotTo(BeNil())
			deplErr, ok := err.(*DeploymentError)
			g.Expect(ok).To(BeTrue())
			g.Expect(deplErr.TopError.Error()).To(c.topErrorMatcher)
			g.Expect(deplErr.ProvisioningState).To(c.provisioningStateMatcher)
			g.Expect(deplErr.StatusCode).To(c.statusCodeMatcher)
			g.Expect(string(deplErr.Response)).To(c.responseMatcher)
			g.Expect(len(deplErr.OperationsLists)).To(c.opsListLenMatcher)
		})
	}
}

func TestDeployTemplateSync_Success(t *testing.T) {
	t.Parallel()

	mockClient := &MockAKSEngineClient{}
	logger := log.NewEntry(log.New())
	err := DeployTemplateSync(mockClient, logger, "rg1", "agentvm", map[string]interface{}{}, map[string]interface{}{})
	g := NewGomegaWithT(t)
	g.Expect(err).To(BeNil())
}

func TestDeploymentError_Error(t *testing.T) {
	t.Parallel()

	operationsLists := make([]resources.DeploymentOperationsListResult, 0)
	operationsList := resources.DeploymentOperationsListResult{}
	operations := make([]resources.DeploymentOperation, 0)
	id := "1234"
	oID := "342"
	provisioningState := "Failed"
	status := map[string]interface{}{
		"message": "sample status message",
	}
	properties := resources.DeploymentOperationProperties{
		ProvisioningState: &provisioningState,
		StatusMessage:     &status,
	}
	operation1 := resources.DeploymentOperation{
		ID:          &id,
		OperationID: &oID,
		Properties:  &properties,
	}
	operations = append(operations, operation1)
	operationsList.Value = &operations
	operationsLists = append(operationsLists, operationsList)
	deploymentErr := &DeploymentError{
		DeploymentName:    "agentvm",
		ResourceGroup:     "rg1",
		TopError:          errors.New("sample error"),
		ProvisioningState: "Failed",
		Response:          []byte("sample resp"),
		StatusCode:        500,
		OperationsLists:   operationsLists,
	}
	errString := deploymentErr.Error()
	expected := `DeploymentName[agentvm] ResourceGroup[rg1] TopError[sample error] StatusCode[500] Response[sample resp] ProvisioningState[Failed] Operations[{
  "message": "sample status message"
}]`
	if errString != expected {
		t.Errorf("expected error with message %s, but got %s", expected, errString)
	}
}
