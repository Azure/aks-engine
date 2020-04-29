// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/ghodss/yaml"
)

func getDCOSProvisionScript(script string) string {
	// add the provision script
	bp, err := Asset(script)
	if err != nil {
		panic(fmt.Sprintf("BUG: %s", err.Error()))
	}

	provisionScript := string(bp)
	if strings.Contains(provisionScript, "'") {
		panic(fmt.Sprintf("BUG: %s may not contain character '", script))
	}

	return strings.Replace(strings.Replace(provisionScript, "\r\n", "\n", -1), "\n", "\n\n    ", -1)
}

func getDCOSBootstrapCustomData(p *api.Properties) string {
	masterIPList, err := generateConsecutiveIPsList(p.MasterProfile.Count, p.MasterProfile.FirstConsecutiveStaticIP)
	if err != nil {
		return ""
	}
	for i, v := range masterIPList {
		masterIPList[i] = "    - " + v
	}

	str := getSingleLineDCOSCustomData(
		p.OrchestratorProfile.OrchestratorType,
		dcos2BootstrapCustomdata, 0,
		map[string]string{
			"PROVISION_SOURCE_STR":    getDCOSProvisionScript(dcosProvisionSource),
			"PROVISION_STR":           getDCOSProvisionScript(dcos2BootstrapProvision),
			"MASTER_IP_LIST":          strings.Join(masterIPList, "\n"),
			"BOOTSTRAP_IP":            p.OrchestratorProfile.DcosConfig.BootstrapProfile.StaticIP,
			"BOOTSTRAP_OAUTH_ENABLED": strconv.FormatBool(p.OrchestratorProfile.DcosConfig.BootstrapProfile.OAuthEnabled)})

	return fmt.Sprintf("\"customData\": \"[base64(concat('#cloud-config\\n\\n', '%s'))]\",", str)
}

func getDCOSCustomDataPublicIPStr(orchestratorType string, masterCount int) string {
	if orchestratorType == api.DCOS {
		var buf bytes.Buffer
		for i := 0; i < masterCount; i++ {
			buf.WriteString(fmt.Sprintf("reference(variables('masterVMNic')[%d]).ipConfigurations[0].properties.privateIPAddress,", i))
			if i < (masterCount - 1) {
				buf.WriteString(`'\\\", \\\"', `)
			}
		}
		return buf.String()
	}
	return ""
}

// getSingleLineForTemplate returns the file as a single line for embedding in an arm template
func getSingleLineDCOSCustomData(orchestratorType, yamlFilename string, masterCount int, replaceMap map[string]string) string {
	b, err := Asset(yamlFilename)
	if err != nil {
		panic(fmt.Sprintf("BUG getting yaml custom data file: %s", err.Error()))
	}
	yamlStr := string(b)
	for k, v := range replaceMap {
		yamlStr = strings.Replace(yamlStr, k, v, -1)
	}

	// convert to json
	jsonBytes, err4 := yaml.YAMLToJSON([]byte(yamlStr))
	if err4 != nil {
		panic(fmt.Sprintf("BUG: %s", err4.Error()))
	}
	yamlStr = string(jsonBytes)

	// convert to one line
	yamlStr = strings.Replace(yamlStr, "\\", "\\\\", -1)
	yamlStr = strings.Replace(yamlStr, "\r\n", "\\n", -1)
	yamlStr = strings.Replace(yamlStr, "\n", "\\n", -1)
	yamlStr = strings.Replace(yamlStr, "\"", "\\\"", -1)

	// variable replacement
	rVariable, e1 := regexp.Compile("{{{([^}]*)}}}")
	if e1 != nil {
		panic(fmt.Sprintf("BUG: %s", e1.Error()))
	}
	yamlStr = rVariable.ReplaceAllString(yamlStr, "',variables('$1'),'")

	// replace the internal values
	publicIPStr := getDCOSCustomDataPublicIPStr(orchestratorType, masterCount)
	yamlStr = strings.Replace(yamlStr, "DCOSCUSTOMDATAPUBLICIPSTR", publicIPStr, -1)

	return yamlStr
}

func getDCOSCustomDataTemplate(orchestratorType, orchestratorVersion string) string {
	switch orchestratorType {
	case api.DCOS:
		switch orchestratorVersion {
		case common.DCOSVersion1Dot8Dot8:
			return dcosCustomData188
		case common.DCOSVersion1Dot9Dot0:
			return dcosCustomData190
		case common.DCOSVersion1Dot9Dot8:
			return dcosCustomData198
		case common.DCOSVersion1Dot10Dot0:
			return dcosCustomData110
		case common.DCOSVersion1Dot11Dot0:
			return dcos2CustomData1110
		case common.DCOSVersion1Dot11Dot2:
			return dcos2CustomData1112
		}
	default:
		// it is a bug to get here
		panic(fmt.Sprintf("BUG: invalid orchestrator %s", orchestratorType))
	}
	return ""
}

func getDCOSMasterCustomData(cs *api.ContainerService) string {
	masterAttributeContents := getDCOSMasterCustomNodeLabels()
	masterPreprovisionExtension := ""
	if cs.Properties.MasterProfile.PreprovisionExtension != nil {
		masterPreprovisionExtension += "\n"
		masterPreprovisionExtension += makeMasterExtensionScriptCommands(cs)
	}
	var bootstrapIP string
	if cs.Properties.OrchestratorProfile.DcosConfig != nil && cs.Properties.OrchestratorProfile.DcosConfig.BootstrapProfile != nil {
		bootstrapIP = cs.Properties.OrchestratorProfile.DcosConfig.BootstrapProfile.StaticIP
	}

	str := getSingleLineDCOSCustomData(
		cs.Properties.OrchestratorProfile.OrchestratorType,
		getDCOSCustomDataTemplate(cs.Properties.OrchestratorProfile.OrchestratorType, cs.Properties.OrchestratorProfile.OrchestratorVersion),
		cs.Properties.MasterProfile.Count,
		map[string]string{
			"PROVISION_SOURCE_STR":   getDCOSProvisionScript(dcosProvisionSource),
			"PROVISION_STR":          getDCOSMasterProvisionScript(cs.Properties.OrchestratorProfile, bootstrapIP),
			"ATTRIBUTES_STR":         masterAttributeContents,
			"PREPROVISION_EXTENSION": masterPreprovisionExtension,
			"ROLENAME":               "master"})

	return fmt.Sprintf("\"customData\": \"[base64(concat('#cloud-config\\n\\n', '%s'))]\",", str)
}

func getDCOSAgentProvisionScript(profile *api.AgentPoolProfile, orchProfile *api.OrchestratorProfile, bootstrapIP string) string {
	// add the provision script
	scriptname := dcos2Provision
	if orchProfile.DcosConfig == nil || orchProfile.DcosConfig.BootstrapProfile == nil {
		if profile.OSType == api.Windows {
			scriptname = dcosWindowsProvision
		} else {
			scriptname = dcosProvision
		}
	}

	bp, err := Asset(scriptname)
	if err != nil {
		panic(fmt.Sprintf("BUG: %s", err.Error()))
	}

	provisionScript := string(bp)
	if strings.Contains(provisionScript, "'") {
		panic(fmt.Sprintf("BUG: %s may not contain character '", dcosProvision))
	}

	// the embedded roleFileContents
	var roleFileContents string
	if len(profile.Ports) > 0 {
		// public agents
		roleFileContents = "touch /etc/mesosphere/roles/slave_public"
	} else {
		roleFileContents = "touch /etc/mesosphere/roles/slave"
	}
	provisionScript = strings.Replace(provisionScript, "ROLESFILECONTENTS", roleFileContents, -1)
	provisionScript = strings.Replace(provisionScript, "BOOTSTRAP_IP", bootstrapIP, -1)

	var b bytes.Buffer
	b.WriteString(provisionScript)
	b.WriteString("\n")

	if len(orchProfile.DcosConfig.Registry) == 0 {
		b.WriteString("rm /etc/docker.tar.gz\n")
	}

	return strings.Replace(strings.Replace(b.String(), "\r\n", "\n", -1), "\n", "\n\n    ", -1)
}

func getDCOSAgentCustomData(cs *api.ContainerService, profile *api.AgentPoolProfile) string {
	attributeContents := getDCOSAgentCustomNodeLabels(profile)
	agentPreprovisionExtension := ""
	if profile.PreprovisionExtension != nil {
		agentPreprovisionExtension += "\n"
		agentPreprovisionExtension += makeAgentExtensionScriptCommands(cs, profile)
	}
	var agentRoleName, bootstrapIP string
	if len(profile.Ports) > 0 {
		agentRoleName = "slave_public"
	} else {
		agentRoleName = "slave"
	}
	if cs.Properties.OrchestratorProfile.DcosConfig != nil && cs.Properties.OrchestratorProfile.DcosConfig.BootstrapProfile != nil {
		bootstrapIP = cs.Properties.OrchestratorProfile.DcosConfig.BootstrapProfile.StaticIP
	}

	str := getSingleLineDCOSCustomData(
		cs.Properties.OrchestratorProfile.OrchestratorType,
		getDCOSCustomDataTemplate(cs.Properties.OrchestratorProfile.OrchestratorType, cs.Properties.OrchestratorProfile.OrchestratorVersion),
		cs.Properties.MasterProfile.Count,
		map[string]string{
			"PROVISION_SOURCE_STR":   getDCOSProvisionScript(dcosProvisionSource),
			"PROVISION_STR":          getDCOSAgentProvisionScript(profile, cs.Properties.OrchestratorProfile, bootstrapIP),
			"ATTRIBUTES_STR":         attributeContents,
			"PREPROVISION_EXTENSION": agentPreprovisionExtension,
			"ROLENAME":               agentRoleName})

	return fmt.Sprintf("\"customData\": \"[base64(concat('#cloud-config\\n\\n', '%s'))]\",", str)
}

func getDCOSWindowsAgentCustomData(cs *api.ContainerService, profile *api.AgentPoolProfile) string {
	agentPreprovisionExtension := ""
	if profile.PreprovisionExtension != nil {
		agentPreprovisionExtension += "\n"
		agentPreprovisionExtension += makeAgentExtensionScriptCommands(cs, profile)
	}
	b, err := Asset(dcosWindowsProvision)
	if err != nil {
		// this should never happen and this is a bug
		panic(fmt.Sprintf("BUG: %s", err.Error()))
	}
	// translate the parameters
	csStr := string(b)
	csStr = strings.Replace(csStr, "PREPROVISION_EXTENSION", agentPreprovisionExtension, -1)
	csStr = strings.Replace(csStr, "\r\n", "\n", -1)
	str := getBase64EncodedGzippedCustomScriptFromStr(csStr)
	return fmt.Sprintf("\"customData\": \"%s\"", str)
}

// getLinkedTemplatesForExtensions returns the
// Microsoft.Resources/deployments for each extension
func getLinkedTemplatesForExtensions(properties *api.Properties) string {
	var result string

	extensions := properties.ExtensionProfiles
	masterProfileExtensions := properties.MasterProfile.Extensions
	orchestratorType := properties.OrchestratorProfile.OrchestratorType

	for err, extensionProfile := range extensions {
		_ = err

		masterOptedForExtension, singleOrAll := validateProfileOptedForExtension(extensionProfile.Name, masterProfileExtensions)
		if masterOptedForExtension {
			result += ","
			dta, e := getMasterLinkedTemplateText(orchestratorType, extensionProfile, singleOrAll)
			if e != nil {
				fmt.Println(e.Error())
				return ""
			}
			result += dta
		}

		for _, agentPoolProfile := range properties.AgentPoolProfiles {
			poolProfileExtensions := agentPoolProfile.Extensions
			poolOptedForExtension, singleOrAll := validateProfileOptedForExtension(extensionProfile.Name, poolProfileExtensions)
			if poolOptedForExtension {
				result += ","
				dta, e := getAgentPoolLinkedTemplateText(agentPoolProfile, orchestratorType, extensionProfile, singleOrAll)
				if e != nil {
					fmt.Println(e.Error())
					return ""
				}
				result += dta
			}

		}
	}

	return result
}
