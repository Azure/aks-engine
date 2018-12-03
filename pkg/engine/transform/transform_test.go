// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package transform

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/Azure/aks-engine/pkg/i18n"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

func TestNormalizeForVMSSScaling(t *testing.T) {
	RegisterTestingT(t)
	logger := logrus.New().WithField("testName", "TestNormalizeForVMSSScaling")
	fileContents, e := ioutil.ReadFile("./transformtestfiles/dcos_template.json")
	Expect(e).To(BeNil())
	expectedFileContents, e := ioutil.ReadFile("./transformtestfiles/dcos_scale_template.json")
	Expect(e).To(BeNil())
	templateJSON := string(fileContents)
	var template interface{}
	json.Unmarshal([]byte(templateJSON), &template)
	templateMap := template.(map[string]interface{})
	transformer := Transformer{}
	e = transformer.NormalizeForVMSSScaling(logger, templateMap)
	Expect(e).To(BeNil())
	ValidateTemplate(templateMap, expectedFileContents, "TestNormalizeForVMSSScaling")
}

func TestNormalizeForK8sVMASScalingUp(t *testing.T) {
	RegisterTestingT(t)
	logger := logrus.New().WithField("testName", "TestNormalizeForK8sVMASScalingUp")
	fileContents, e := ioutil.ReadFile("./transformtestfiles/k8s_template.json")
	Expect(e).To(BeNil())
	expectedFileContents, e := ioutil.ReadFile("./transformtestfiles/k8s_scale_template.json")
	Expect(e).To(BeNil())
	templateJSON := string(fileContents)
	var template interface{}
	json.Unmarshal([]byte(templateJSON), &template)
	templateMap := template.(map[string]interface{})
	transformer := Transformer{}
	e = transformer.NormalizeForK8sVMASScalingUp(logger, templateMap)
	Expect(e).To(BeNil())
	ValidateTemplate(templateMap, expectedFileContents, "TestNormalizeForK8sVMASScalingUp")
}

func TestNormalizeForK8sVMASScalingUpWithVnet(t *testing.T) {
	RegisterTestingT(t)
	logger := logrus.New().WithField("testName", "TestNormalizeForK8sVMASScalingUp")
	fileContents, e := ioutil.ReadFile("./transformtestfiles/k8s_vnet_template.json")
	Expect(e).To(BeNil())
	expectedFileContents, e := ioutil.ReadFile("./transformtestfiles/k8s_vnet_scale_template.json")
	Expect(e).To(BeNil())
	templateJSON := string(fileContents)
	var template interface{}
	json.Unmarshal([]byte(templateJSON), &template)
	templateMap := template.(map[string]interface{})
	transformer := Transformer{}
	e = transformer.NormalizeForK8sVMASScalingUp(logger, templateMap)
	Expect(e).To(BeNil())
	ValidateTemplate(templateMap, expectedFileContents, "TestNormalizeForK8sVMASScalingUpWithVnet")
}

func TestNormalizeForOpenShiftVMASScalingUp(t *testing.T) {
	RegisterTestingT(t)

	tests := []struct {
		agentPoolName       string
		templateMap         map[string]interface{}
		expectedTemplateMap map[string]interface{}
		expectedErr         bool
	}{
		{
			// a badly constructed input should result in an error, not a panic
			expectedErr: true,
		},
		{
			agentPoolName: "compute",
			templateMap: map[string]interface{}{
				"resources": []interface{}{
					map[string]interface{}{
						"name": "foo",
					},
					map[string]interface{}{
						"name": "barVMNamePrefix",
						"dependsOn": []interface{}{
							"foo",
						},
					},
					map[string]interface{}{
						"name": "computeVMNamePrefix",
						"dependsOn": []interface{}{
							"foo",
							"barVMNamePrefix",
							"computeVMNamePrefix",
						},
					},
				},
				"outputs": map[string]interface{}{
					"junk": "junk",
				},
			},
			expectedTemplateMap: map[string]interface{}{
				"resources": []interface{}{
					map[string]interface{}{
						"name": "computeVMNamePrefix",
						"dependsOn": []interface{}{
							"computeVMNamePrefix",
						},
					},
				},
				"outputs": map[string]interface{}{},
			},
		},
	}

	transformer := &Transformer{}

	for i, test := range tests {
		fmt.Fprintf(GinkgoWriter, "test %d\n", i)
		err := transformer.NormalizeForOpenShiftVMASScalingUp(nil, test.agentPoolName, test.templateMap)
		Expect(err != nil).To(Equal(test.expectedErr))
		Expect(test.templateMap).To(BeEquivalentTo(test.expectedTemplateMap))
	}
}

func TestNormalizeResourcesForK8sMasterUpgrade(t *testing.T) {
	RegisterTestingT(t)
	logger := logrus.New().WithField("testName", "TestNormalizeResourcesForK8sMasterUpgrade")
	fileContents, e := ioutil.ReadFile("./transformtestfiles/k8s_template.json")
	Expect(e).To(BeNil())
	expectedFileContents, e := ioutil.ReadFile("./transformtestfiles/k8s_master_upgrade_template.json")
	Expect(e).To(BeNil())
	templateJSON := string(fileContents)
	var template interface{}
	json.Unmarshal([]byte(templateJSON), &template)
	templateMap := template.(map[string]interface{})
	transformer := &Transformer{
		Translator: &i18n.Translator{
			Locale: nil,
		},
	}
	agentsToKeepMap := make(map[string]bool)
	agentsToKeepMap["agentpool1"] = true
	agentsToKeepMap["agentpool2"] = true
	e = transformer.NormalizeResourcesForK8sMasterUpgrade(logger, templateMap, false, agentsToKeepMap)
	Expect(e).To(BeNil())
	ValidateTemplate(templateMap, expectedFileContents, "TestNormalizeResourcesForK8sMasterUpgrade")
}

func TestNormalizeResourcesForK8sAgentUpgrade(t *testing.T) {
	RegisterTestingT(t)
	logger := logrus.New().WithField("testName", "TestNormalizeResourcesForK8sAgentUpgrade")
	fileContents, e := ioutil.ReadFile("./transformtestfiles/k8s_template.json")
	Expect(e).To(BeNil())
	expectedFileContents, e := ioutil.ReadFile("./transformtestfiles/k8s_agent_upgrade_template.json")
	Expect(e).To(BeNil())
	templateJSON := string(fileContents)
	var template interface{}
	json.Unmarshal([]byte(templateJSON), &template)
	templateMap := template.(map[string]interface{})
	transformer := &Transformer{
		Translator: &i18n.Translator{
			Locale: nil,
		},
	}
	agentsToKeepMap := make(map[string]bool)
	agentsToKeepMap["agentpool1"] = true
	agentsToKeepMap["agentpool2"] = false
	e = transformer.NormalizeResourcesForK8sAgentUpgrade(logger, templateMap, false, agentsToKeepMap)
	Expect(e).To(BeNil())
	ValidateTemplate(templateMap, expectedFileContents, "TestNormalizeResourcesForK8sAgentUpgrade")
}

func ValidateTemplate(templateMap map[string]interface{}, expectedFileContents []byte, testFileName string) {
	output, e := helpers.JSONMarshal(templateMap, false)
	Expect(e).To(BeNil())
	prettyOutput, e := PrettyPrintArmTemplate(string(output))
	Expect(e).To(BeNil())
	prettyExpectedOutput, e := PrettyPrintArmTemplate(string(expectedFileContents))
	Expect(e).To(BeNil())
	if prettyOutput != prettyExpectedOutput {
		ioutil.WriteFile(fmt.Sprintf("./transformtestfiles/%s.failure.json", testFileName), []byte(prettyOutput), 0600)
	}
	Expect(prettyOutput).To(Equal(prettyExpectedOutput))
}
