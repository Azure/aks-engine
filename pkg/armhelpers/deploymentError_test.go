package armhelpers

import (
	"testing"

	. "github.com/Azure/aks-engine/pkg/test"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func TestUpgradeCluster(t *testing.T) {
	RunSpecsWithReporters(t, "templatedeployment", "Server Suite")
}

var _ = Describe("Template deployment tests", func() {

	It("Should return InternalOperationError error code", func() {
		mockClient := &MockAKSEngineClient{}
		mockClient.FailDeployTemplate = true
		logger := log.NewEntry(log.New())

		err := DeployTemplateSync(mockClient, logger, "rg1", "agentvm", map[string]interface{}{}, map[string]interface{}{})
		Expect(err).NotTo(BeNil())
		deplErr, ok := err.(*DeploymentError)
		Expect(ok).To(BeTrue())
		Expect(deplErr.TopError).NotTo(BeNil())
		Expect(deplErr.TopError.Error()).To(Equal("DeployTemplate failed"))
		Expect(deplErr.ProvisioningState).To(Equal(""))
		Expect(deplErr.StatusCode).To(Equal(0))
		Expect(len(deplErr.OperationsLists)).To(Equal(0))
	})

	It("Should return QuotaExceeded error code, specified in details", func() {
		mockClient := &MockAKSEngineClient{}
		mockClient.FailDeployTemplateQuota = true
		logger := log.NewEntry(log.New())

		err := DeployTemplateSync(mockClient, logger, "rg1", "agentvm", map[string]interface{}{}, map[string]interface{}{})
		Expect(err).NotTo(BeNil())
		deplErr, ok := err.(*DeploymentError)
		Expect(ok).To(BeTrue())
		Expect(deplErr.TopError).NotTo(BeNil())
		Expect(deplErr.ProvisioningState).To(Equal(""))
		Expect(deplErr.StatusCode).To(Equal(400))
		Expect(string(deplErr.Response)).To(ContainSubstring("\"code\":\"QuotaExceeded\""))
		Expect(len(deplErr.OperationsLists)).To(Equal(0))
	})

	It("Should return Conflict error code, specified in details", func() {
		mockClient := &MockAKSEngineClient{}
		mockClient.FailDeployTemplateConflict = true
		logger := log.NewEntry(log.New())

		err := DeployTemplateSync(mockClient, logger, "rg1", "agentvm", map[string]interface{}{}, map[string]interface{}{})
		Expect(err).NotTo(BeNil())
		deplErr, ok := err.(*DeploymentError)
		Expect(ok).To(BeTrue())
		Expect(deplErr.TopError).NotTo(BeNil())
		Expect(deplErr.ProvisioningState).To(Equal(""))
		Expect(deplErr.StatusCode).To(Equal(200))
		Expect(string(deplErr.Response)).To(ContainSubstring("\"code\":\"Conflict\""))
		Expect(len(deplErr.OperationsLists)).To(Equal(0))
	})

	It("Should return deployment error with Operations Lists", func() {
		mockClient := &MockAKSEngineClient{}
		mockClient.FailDeployTemplateWithProperties = true
		logger := log.NewEntry(log.New())

		err := DeployTemplateSync(mockClient, logger, "rg1", "agentvm", map[string]interface{}{}, map[string]interface{}{})
		Expect(err).NotTo(BeNil())
		deplErr, ok := err.(*DeploymentError)
		Expect(ok).To(BeTrue())
		Expect(deplErr.TopError).NotTo(BeNil())
		Expect(deplErr.ProvisioningState).To(Equal("Failed"))
		Expect(deplErr.StatusCode).To(Equal(200))
		Expect(string(deplErr.Response)).To(ContainSubstring("\"code\":\"Conflict\""))
		Expect(len(deplErr.OperationsLists)).To(Equal(2))
	})

	It("Should return nil on success", func() {
		mockClient := &MockAKSEngineClient{}
		logger := log.NewEntry(log.New())
		err := DeployTemplateSync(mockClient, logger, "rg1", "agentvm", map[string]interface{}{}, map[string]interface{}{})
		Expect(err).To(BeNil())
	})
})

func TestDeploymentError_Error(t *testing.T) {
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
