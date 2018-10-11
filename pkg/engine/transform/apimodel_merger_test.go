package transform

import (
	"io/ioutil"
	"testing"

	"github.com/Jeffail/gabs"
	. "github.com/onsi/gomega"
)

func TestAPIModelMergerMapValues(t *testing.T) {
	RegisterTestingT(t)

	m := make(map[string]APIModelValue)
	values := []string{
		"masterProfile.count=5",
		"agentPoolProfiles[0].name=agentpool1",
		"linuxProfile.adminUsername=admin",
		"servicePrincipalProfile.clientId='123a1238-c6eb-4b61-9d6f-7db6f1e14123',servicePrincipalProfile.secret='=!,Test$^='",
	}

	MapValues(m, values)
	Expect(m["masterProfile.count"].intValue).To(BeIdenticalTo(int64(5)))
	Expect(m["agentPoolProfiles[0].name"].arrayValue).To(BeTrue())
	Expect(m["agentPoolProfiles[0].name"].arrayIndex).To(BeIdenticalTo(0))
	Expect(m["agentPoolProfiles[0].name"].arrayProperty).To(BeIdenticalTo("name"))
	Expect(m["agentPoolProfiles[0].name"].arrayName).To(BeIdenticalTo("agentPoolProfiles"))
	Expect(m["agentPoolProfiles[0].name"].stringValue).To(BeIdenticalTo("agentpool1"))
	Expect(m["linuxProfile.adminUsername"].stringValue).To(BeIdenticalTo("admin"))
	Expect(m["servicePrincipalProfile.secret"].stringValue).To(BeIdenticalTo("=!,Test$^="))
	Expect(m["servicePrincipalProfile.clientId"].stringValue).To(BeIdenticalTo("123a1238-c6eb-4b61-9d6f-7db6f1e14123"))
}

func TestMergeValuesWithAPIModel(t *testing.T) {
	RegisterTestingT(t)

	m := make(map[string]APIModelValue)
	values := []string{"masterProfile.count=5", "agentPoolProfiles[0].name=agentpool1", "linuxProfile.adminUsername=admin"}

	MapValues(m, values)
	tmpFile, _ := MergeValuesWithAPIModel("../testdata/simple/kubernetes.json", m)

	jsonFileContent, err := ioutil.ReadFile(tmpFile)
	Expect(err).To(BeNil())

	jsonAPIModel, err := gabs.ParseJSON(jsonFileContent)
	Expect(err).To(BeNil())

	masterProfileCount := jsonAPIModel.Path("properties.masterProfile.count").Data()
	Expect(masterProfileCount).To(BeIdenticalTo(float64(5)))

	adminUsername := jsonAPIModel.Path("properties.linuxProfile.adminUsername").Data()
	Expect(adminUsername).To(BeIdenticalTo("admin"))

	agentPoolProfileName := jsonAPIModel.Path("properties.agentPoolProfiles").Index(0).Path("name").Data().(string)
	Expect(agentPoolProfileName).To(BeIdenticalTo("agentpool1"))
}
