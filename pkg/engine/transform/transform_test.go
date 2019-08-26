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
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

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

func TestNormalizeMasterResourcesForScaling(t *testing.T) {
	RegisterTestingT(t)
	logger := logrus.New().WithField("testName", "TestNormalizeMasterResourcesForScaling")
	fileContents, e := ioutil.ReadFile("./transformtestfiles/k8s_template.json")
	Expect(e).To(BeNil())
	expectedFileContents, e := ioutil.ReadFile("./transformtestfiles/master_resources_scale_temaplate.json")
	Expect(e).To(BeNil())
	templateJSON := string(fileContents)
	var template interface{}
	json.Unmarshal([]byte(templateJSON), &template)
	templateMap := template.(map[string]interface{})
	transformer := Transformer{}
	e = transformer.NormalizeMasterResourcesForScaling(logger, templateMap)
	Expect(e).To(BeNil())
	ValidateTemplate(templateMap, expectedFileContents, "TestNormalizeMasterResourcesForScaling")
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

func TestNormalizeForK8sSLBScalingOrUpgrade(t *testing.T) {
	RegisterTestingT(t)
	logger := logrus.New().WithField("testName", "NormalizeForK8sSLBScalingOrUpgrade")
	fileContents, e := ioutil.ReadFile("./transformtestfiles/k8s_slb_template.json")
	Expect(e).To(BeNil())
	expectedFileContents, e := ioutil.ReadFile("./transformtestfiles/k8s_slb_scale_template.json")
	Expect(e).To(BeNil())
	templateJSON := string(fileContents)
	var template interface{}
	json.Unmarshal([]byte(templateJSON), &template)
	templateMap := template.(map[string]interface{})
	transformer := Transformer{}
	e = transformer.NormalizeForK8sSLBScalingOrUpgrade(logger, templateMap)
	Expect(e).To(BeNil())
	ValidateTemplate(templateMap, expectedFileContents, "TestNormalizeForK8sSLBScalingOrUpgrade")
}

func TestNormalizeForK8sSLBScalingOrUpgradeVMSS(t *testing.T) {
	RegisterTestingT(t)
	logger := logrus.New().WithField("testName", "NormalizeForK8sSLBScalingOrUpgrade")
	fileContents, e := ioutil.ReadFile("./transformtestfiles/k8s_slb_vmss_template.json")
	Expect(e).To(BeNil())
	expectedFileContents, e := ioutil.ReadFile("./transformtestfiles/k8s_slb_vmss_scale_template.json")
	Expect(e).To(BeNil())
	templateJSON := string(fileContents)
	var template interface{}
	json.Unmarshal([]byte(templateJSON), &template)
	templateMap := template.(map[string]interface{})
	transformer := Transformer{}
	e = transformer.NormalizeForK8sSLBScalingOrUpgrade(logger, templateMap)
	Expect(e).To(BeNil())
	ValidateTemplate(templateMap, expectedFileContents, "TestNormalizeForK8sSLBScalingOrUpgradeVMSS")
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
