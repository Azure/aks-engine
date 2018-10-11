package operations

import (
	"testing"

	"github.com/Azure/aks-engine/pkg/armhelpers"
	. "github.com/Azure/aks-engine/pkg/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
)

func TestOperations(t *testing.T) {
	RunSpecsWithReporters(t, "operations", "Server Suite")
}

var _ = Describe("Scale down vms operation tests", func() {
	It("Should return error messages for failing vms", func() {
		mockClient := armhelpers.MockAKSEngineClient{}
		mockClient.FailGetVirtualMachine = true
		errs := ScaleDownVMs(&mockClient, log.NewEntry(log.New()), "sid", "rg", "vm1", "vm2", "vm3", "vm5")
		Expect(errs.Len()).To(Equal(4))
		for e := errs.Front(); e != nil; e = e.Next() {
			output := e.Value.(*VMScalingErrorDetails)
			Expect(output.Name).To(ContainSubstring("vm"))
			Expect(output.Error).To(Not(BeNil()))
		}
	})
	It("Should return nil for errors if all deletes successful", func() {
		mockClient := armhelpers.MockAKSEngineClient{}
		errs := ScaleDownVMs(&mockClient, log.NewEntry(log.New()), "sid", "rg", "k8s-agent-F8EADCCF-0", "k8s-agent-F8EADCCF-3", "k8s-agent-F8EADCCF-2", "k8s-agent-F8EADCCF-4")
		Expect(errs).To(BeNil())
	})
})
