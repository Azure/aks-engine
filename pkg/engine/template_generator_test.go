// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"text/template"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/api/common"
	"github.com/Azure/aks-engine/pkg/telemetry"
	"github.com/google/go-cmp/cmp"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/pkg/errors"
)

func TestGenerateTemplateV2(t *testing.T) {
	tg, _ := InitializeTemplateGenerator(Context{})

	apiModelStr := getAPIModelString()
	cs := &api.ContainerService{}

	err := json.Unmarshal([]byte(apiModelStr), &cs)

	if err != nil {
		t.Errorf("unexpected error while unmarshalling the apiModel JSON: %s", err.Error())
	}

	_, _, err = tg.GenerateTemplateV2(cs, DefaultGeneratorCode, TestAKSEngineVersion)

	if err != nil {
		t.Errorf("unexpected error while running getParameterDescMap: %s", err.Error())
	}
}

func TestGetTemplateFuncMap(t *testing.T) {
	funcmap, err := getFuncMap(getAPIModelString())
	if err != nil {
		t.Fatalf("Error generating function map: %v", err)
	}
	cases := []string{
		"IsCustomCloudProfile",
		"GetCustomCloudRootCertificates",
		"GetCustomCloudSourcesList",
		"IsAzureStackCloud",
		"IsMultiMasterCluster",
		"IsMasterVirtualMachineScaleSets",
		"IsHostedMaster",
		"IsIPMasqAgentEnabled",
		"IsKubernetesVersionGe",
		"IsKubernetesVersionLt",
		"GetMasterKubernetesLabels",
		"GetMasterKubernetesLabelsDeprecated",
		"GetAgentKubernetesLabels",
		"GetAgentKubernetesLabelsDeprecated",
		"GetKubeletConfigKeyVals",
		"GetKubeletConfigKeyValsPsh",
		"GetK8sRuntimeConfigKeyVals",
		"HasPrivateRegistry",
		"IsSwarmMode",
		"IsKubernetes",
		"IsPublic",
		"IsAzureCNI",
		"HasCosmosEtcd",
		"GetCosmosEndPointUri",
		"IsPrivateCluster",
		"ProvisionJumpbox",
		"UseManagedIdentity",
		"GetVNETSubnetDependencies",
		"GetLBRules",
		"GetProbes",
		"GetSecurityRules",
		"GetUniqueNameSuffix",
		"GetVNETAddressPrefixes",
		"GetVNETSubnets",
		"GetDataDisks",
		"HasBootstrap",
		"GetMasterAllowedSizes",
		"GetDefaultVNETCIDR",
		"GetKubernetesAllowedVMSKUs",
		"GetSizeMap",
		"WriteLinkedTemplatesForExtensions",
		"GetSshPublicKeysPowerShell",
		"GetWindowsMasterSubnetARMParam",
		"GetKubernetesMasterPreprovisionYaml",
		"GetKubernetesAgentPreprovisionYaml",
		"GetMasterSwarmCustomData",
		"GetAgentSwarmCustomData",
		"GetSwarmAgentPreprovisionExtensionCommands",
		"GetLocation",
		"GetWinAgentSwarmCustomData",
		"GetWinAgentSwarmModeCustomData",
		"GetKubernetesWindowsAgentFunctions",
		"GetMasterSwarmModeCustomData",
		"GetAgentSwarmModeCustomData",
		"GetPodInfraContainerSpec",
		"IsKubenet",
		"WrapAsVariable",
		"CloudInitData",
		"WrapAsParameter",
		"WrapAsVerbatim",
		"HasVMASAgentPool",
		"AnyAgentIsLinux",
		"IsNSeriesSKU",
		"HasAvailabilityZones",
		"GetBase64EncodedEnvironmentJSON",
		"GetIdentitySystem",
		"NeedsContainerd",
		// TODO validate that the remaining func strings in getTemplateFuncMap are thinly wrapped and unit tested
	}

	for _, c := range cases {
		f, ok := funcmap[c]
		v := reflect.ValueOf(f)
		if !ok {
			t.Fatalf("Didn't find expected funcmap key %s.", c)
		}
		switch c {
		case "GetDefaultVNETCIDR":
			rargs := make([]reflect.Value, 0)
			ret := v.Call(rargs)
			if ret[0].Interface() != DefaultVNETCIDR {
				t.Fatalf("Got unexpected default VNET CIDR")
			}
		case "IsMultiMasterCluster":
			rargs := make([]reflect.Value, 0)
			ret := v.Call(rargs)
			if ret[0].Interface() != false {
				t.Fatalf("Got unexpected IsMultiMasterCluster response")
			}
		case "GetKubernetesMasterPreprovisionYaml":
			rargs := make([]reflect.Value, 0)
			ret := v.Call(rargs)
			if ret[0].Interface() != "" {
				t.Fatalf("Got unexpected GetKubernetesMasterPreprovisionYaml response")
			}
		case "NeedsContainerd":
			rargs := make([]reflect.Value, 0)
			ret := v.Call(rargs)
			if ret[0].Interface() != false {
				t.Fatalf("Got unexpected NeedsContainerd response")
			}
		}
	}
}

func TestGetContainerServiceFuncMap(t *testing.T) {
	azureStackCloudSpec := api.AzureEnvironmentSpecConfig{
		CloudName: api.AzureStackCloud,
		KubernetesSpecConfig: api.KubernetesSpecConfig{
			KubernetesImageBase: "azurestack/",
		},
		EndpointConfig: api.AzureEndpointConfig{
			ResourceManagerVMDNSSuffix: "ResourceManagerVMDNSSuffix",
		},
	}
	api.AzureCloudSpecEnvMap[api.AzureStackCloud] = azureStackCloudSpec
	var errorCodeStrings []string
	var errorCodes []int
	for k, v := range cseErrorCodes {
		errorCodeStrings = append(errorCodeStrings, k)
		errorCodes = append(errorCodes, v)
	}
	errorCodeStrings = append(errorCodeStrings, "ERR_HOLD_MY_BEER")
	errorCodes = append(errorCodes, -1)
	cases := []struct {
		name                                  string
		cs                                    *api.ContainerService
		expectedHasCustomSearchDomain         bool
		expectedGetSearchDomainName           string
		expectedGetSearchDomainRealmUser      string
		expectedGetSearchDomainRealmPassword  string
		expectedHasCustomNodesDNS             bool
		expectedGetHyperkubeImageReference    string
		expectedGetTargetEnvironment          string
		expectedIsNSeriesSKU                  bool
		expectedIsDockerContainerRuntime      bool
		expectedHasPrivateAzureRegistryServer bool
		expectedGetPrivateAzureRegistryServer string
		expectedGetSysctlDConfigKeyVals       string
		expectedGetCSEErrorCodeVals           []int
		expectedHasVHDDistroNodes             bool
		expectedIsVHDDistroForAllNodes        bool
		expectedHasClusterInitComponent       bool
		expectedIsVirtualMachineScaleSets     bool
		expectedUseManagedIdentity            bool
		expectedHasKubeReservedCgroup         bool
		expectedGetKubeReservedCgroup         string
	}{
		{
			name: "1.15 release",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime:        api.Docker,
							KubernetesImageBaseType: common.KubernetesImageBaseTypeGCR,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
			},
			expectedHasCustomSearchDomain:        false,
			expectedGetSearchDomainName:          "",
			expectedGetSearchDomainRealmUser:     "",
			expectedGetSearchDomainRealmPassword: "",
			expectedHasCustomNodesDNS:            false,
			expectedGetHyperkubeImageReference:   "hyperkube-amd64:v1.15.4",
			expectedGetTargetEnvironment:         "AzurePublicCloud",
			expectedIsNSeriesSKU:                 false,
			expectedIsDockerContainerRuntime:     true,
			expectedGetSysctlDConfigKeyVals:      "",
			expectedGetCSEErrorCodeVals:          []int{-1},
			expectedIsVirtualMachineScaleSets:    true,
		},
		{
			name: "1.16 release",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.16.1",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime:        api.Docker,
							KubernetesImageBaseType: common.KubernetesImageBaseTypeGCR,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
			},
			expectedHasCustomSearchDomain:        false,
			expectedGetSearchDomainName:          "",
			expectedGetSearchDomainRealmUser:     "",
			expectedGetSearchDomainRealmPassword: "",
			expectedHasCustomNodesDNS:            false,
			expectedGetHyperkubeImageReference:   "hyperkube-amd64:v1.16.1",
			expectedGetTargetEnvironment:         "AzurePublicCloud",
			expectedIsNSeriesSKU:                 false,
			expectedIsDockerContainerRuntime:     true,
			expectedGetSysctlDConfigKeyVals:      "",
			expectedIsVirtualMachineScaleSets:    true,
		},
		{
			name: "1.17 release",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.17.0-beta.1",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime:        api.Docker,
							KubernetesImageBaseType: common.KubernetesImageBaseTypeGCR,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
			},
			expectedHasCustomSearchDomain:        false,
			expectedGetSearchDomainName:          "",
			expectedGetSearchDomainRealmUser:     "",
			expectedGetSearchDomainRealmPassword: "",
			expectedHasCustomNodesDNS:            false,
			expectedGetHyperkubeImageReference:   "",
			expectedGetTargetEnvironment:         "AzurePublicCloud",
			expectedIsNSeriesSKU:                 false,
			expectedIsDockerContainerRuntime:     true,
			expectedGetSysctlDConfigKeyVals:      "",
			expectedIsVirtualMachineScaleSets:    true,
		},
		{
			name: "1.17 release w/ VHD distro",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.17.0-beta.1",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime:        api.Docker,
							KubernetesImageBaseType: common.KubernetesImageBaseTypeGCR,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
							Distro:              api.AKSUbuntu1604,
						},
					},
				},
			},
			expectedHasCustomSearchDomain:        false,
			expectedGetSearchDomainName:          "",
			expectedGetSearchDomainRealmUser:     "",
			expectedGetSearchDomainRealmPassword: "",
			expectedHasCustomNodesDNS:            false,
			expectedGetHyperkubeImageReference:   "",
			expectedGetTargetEnvironment:         "AzurePublicCloud",
			expectedIsNSeriesSKU:                 false,
			expectedIsDockerContainerRuntime:     true,
			expectedGetSysctlDConfigKeyVals:      "",
			expectedHasVHDDistroNodes:            true,
			expectedIsVHDDistroForAllNodes:       true,
			expectedIsVirtualMachineScaleSets:    true,
		},
		{
			name: "custom search domain",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime:        api.Docker,
							KubernetesImageBaseType: common.KubernetesImageBaseTypeGCR,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
					LinuxProfile: &api.LinuxProfile{
						CustomSearchDomain: &api.CustomSearchDomain{
							Name:          "foo",
							RealmUser:     "bar",
							RealmPassword: "baz",
						},
					},
				},
			},
			expectedHasCustomSearchDomain:        true,
			expectedGetSearchDomainName:          "foo",
			expectedGetSearchDomainRealmUser:     "bar",
			expectedGetSearchDomainRealmPassword: "baz",
			expectedHasCustomNodesDNS:            false,
			expectedGetHyperkubeImageReference:   "hyperkube-amd64:v1.15.4",
			expectedGetTargetEnvironment:         "AzurePublicCloud",
			expectedIsNSeriesSKU:                 false,
			expectedIsDockerContainerRuntime:     true,
			expectedGetSysctlDConfigKeyVals:      "",
			expectedIsVirtualMachineScaleSets:    true,
		},
		{
			name: "custom nodes DNS",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime:        api.Docker,
							KubernetesImageBaseType: common.KubernetesImageBaseTypeGCR,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
					LinuxProfile: &api.LinuxProfile{
						CustomNodesDNS: &api.CustomNodesDNS{
							DNSServer: "foo",
						},
					},
				},
			},
			expectedHasCustomSearchDomain:        false,
			expectedGetSearchDomainName:          "",
			expectedGetSearchDomainRealmUser:     "",
			expectedGetSearchDomainRealmPassword: "",
			expectedHasCustomNodesDNS:            true,
			expectedGetHyperkubeImageReference:   "hyperkube-amd64:v1.15.4",
			expectedGetTargetEnvironment:         "AzurePublicCloud",
			expectedIsNSeriesSKU:                 false,
			expectedIsDockerContainerRuntime:     true,
			expectedGetSysctlDConfigKeyVals:      "",
			expectedIsVirtualMachineScaleSets:    true,
		},
		{
			name: "1.17 release with custom kube images",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.17.0",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime:                 api.Docker,
							CustomKubeAPIServerImage:         "example.azurecr.io/kube-apiserver-amd64:tag",
							CustomKubeControllerManagerImage: "example.azurecr.io/kube-controller-manager-amd64:tag",
							CustomKubeSchedulerImage:         "example.azurecr.io/kube-scheduler-amd64:tag",
							KubernetesImageBaseType:          common.KubernetesImageBaseTypeGCR,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
			},
			expectedHasCustomSearchDomain:        false,
			expectedGetSearchDomainName:          "",
			expectedGetSearchDomainRealmUser:     "",
			expectedGetSearchDomainRealmPassword: "",
			expectedHasCustomNodesDNS:            false,
			expectedGetHyperkubeImageReference:   "",
			expectedGetTargetEnvironment:         "AzurePublicCloud",
			expectedIsNSeriesSKU:                 false,
			expectedIsDockerContainerRuntime:     true,
			expectedGetSysctlDConfigKeyVals:      "",
			expectedIsVirtualMachineScaleSets:    true,
		},
		{
			name: "china cloud",
			cs: &api.ContainerService{
				Location: "chinaeast",
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime:        api.Docker,
							KubernetesImageBaseType: common.KubernetesImageBaseTypeGCR,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
			},
			expectedHasCustomSearchDomain:        false,
			expectedGetSearchDomainName:          "",
			expectedGetSearchDomainRealmUser:     "",
			expectedGetSearchDomainRealmPassword: "",
			expectedHasCustomNodesDNS:            false,
			expectedGetHyperkubeImageReference:   "hyperkube-amd64:v1.15.4",
			expectedGetTargetEnvironment:         "AzureChinaCloud",
			expectedIsNSeriesSKU:                 false,
			expectedIsDockerContainerRuntime:     true,
			expectedGetSysctlDConfigKeyVals:      "",
			expectedIsVirtualMachineScaleSets:    true,
		},
		{
			name: "german cloud",
			cs: &api.ContainerService{
				Location: "germanynortheast",
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime:        api.Docker,
							KubernetesImageBaseType: common.KubernetesImageBaseTypeGCR,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
			},
			expectedHasCustomSearchDomain:        false,
			expectedGetSearchDomainName:          "",
			expectedGetSearchDomainRealmUser:     "",
			expectedGetSearchDomainRealmPassword: "",
			expectedHasCustomNodesDNS:            false,
			expectedGetHyperkubeImageReference:   "hyperkube-amd64:v1.15.4",
			expectedGetTargetEnvironment:         "AzureGermanCloud",
			expectedIsNSeriesSKU:                 false,
			expectedIsDockerContainerRuntime:     true,
			expectedGetSysctlDConfigKeyVals:      "",
			expectedIsVirtualMachineScaleSets:    true,
		},
		{
			name: "usgov cloud",
			cs: &api.ContainerService{
				Location: "usgoveast2",
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime:        api.Docker,
							KubernetesImageBaseType: common.KubernetesImageBaseTypeGCR,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
			},
			expectedHasCustomSearchDomain:        false,
			expectedGetSearchDomainName:          "",
			expectedGetSearchDomainRealmUser:     "",
			expectedGetSearchDomainRealmPassword: "",
			expectedHasCustomNodesDNS:            false,
			expectedGetHyperkubeImageReference:   "hyperkube-amd64:v1.15.4",
			expectedGetTargetEnvironment:         "AzureUSGovernmentCloud",
			expectedIsNSeriesSKU:                 false,
			expectedIsDockerContainerRuntime:     true,
			expectedGetSysctlDConfigKeyVals:      "",
			expectedIsVirtualMachineScaleSets:    true,
		},
		{
			name: "Azure Stack",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					CustomCloudProfile: &api.CustomCloudProfile{
						Environment: &azure.Environment{
							Name: api.AzureStackCloud,
						},
					},
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime:        api.Docker,
							KubernetesImageBaseType: common.KubernetesImageBaseTypeGCR,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
			},
			expectedHasCustomSearchDomain:        false,
			expectedGetSearchDomainName:          "",
			expectedGetSearchDomainRealmUser:     "",
			expectedGetSearchDomainRealmPassword: "",
			expectedHasCustomNodesDNS:            false,
			expectedGetHyperkubeImageReference:   "hyperkube-amd64:v1.15.4-azs",
			expectedGetTargetEnvironment:         "AzureStackCloud",
			expectedIsNSeriesSKU:                 false,
			expectedIsDockerContainerRuntime:     true,
			expectedGetSysctlDConfigKeyVals:      "",
			expectedIsVirtualMachineScaleSets:    true,
		},
		{
			name: "N series SKU",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime:        api.Docker,
							KubernetesImageBaseType: common.KubernetesImageBaseTypeGCR,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
							VMSize:              "Standard_NC6",
						},
					},
				},
			},
			expectedHasCustomSearchDomain:        false,
			expectedGetSearchDomainName:          "",
			expectedGetSearchDomainRealmUser:     "",
			expectedGetSearchDomainRealmPassword: "",
			expectedHasCustomNodesDNS:            false,
			expectedGetHyperkubeImageReference:   "hyperkube-amd64:v1.15.4",
			expectedGetTargetEnvironment:         "AzurePublicCloud",
			expectedIsNSeriesSKU:                 true,
			expectedIsDockerContainerRuntime:     true,
			expectedGetSysctlDConfigKeyVals:      "",
			expectedIsVirtualMachineScaleSets:    true,
		},
		{
			name: "AvailabilitySet pool",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime:        api.Docker,
							KubernetesImageBaseType: common.KubernetesImageBaseTypeGCR,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.AvailabilitySet,
						},
					},
				},
			},
			expectedHasCustomSearchDomain:        false,
			expectedGetSearchDomainName:          "",
			expectedGetSearchDomainRealmUser:     "",
			expectedGetSearchDomainRealmPassword: "",
			expectedHasCustomNodesDNS:            false,
			expectedGetHyperkubeImageReference:   "hyperkube-amd64:v1.15.4",
			expectedGetTargetEnvironment:         "AzurePublicCloud",
			expectedIsDockerContainerRuntime:     true,
			expectedGetSysctlDConfigKeyVals:      "",
			expectedIsVirtualMachineScaleSets:    false,
		},
		{
			name: "managed identity enabled",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime:        api.Docker,
							KubernetesImageBaseType: common.KubernetesImageBaseTypeGCR,
							UseManagedIdentity:      true,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
			},
			expectedHasCustomSearchDomain:        false,
			expectedGetSearchDomainName:          "",
			expectedGetSearchDomainRealmUser:     "",
			expectedGetSearchDomainRealmPassword: "",
			expectedHasCustomNodesDNS:            false,
			expectedGetHyperkubeImageReference:   "hyperkube-amd64:v1.15.4",
			expectedGetTargetEnvironment:         "AzurePublicCloud",
			expectedIsDockerContainerRuntime:     true,
			expectedGetSysctlDConfigKeyVals:      "",
			expectedIsVirtualMachineScaleSets:    true,
			expectedUseManagedIdentity:           true,
		},
		{
			name: "PrivateAzureRegistryServer",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.15.4",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime:           api.Docker,
							PrivateAzureRegistryServer: "my-server",
							KubernetesImageBaseType:    common.KubernetesImageBaseTypeGCR,
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
			},
			expectedHasCustomSearchDomain:         false,
			expectedGetSearchDomainName:           "",
			expectedGetSearchDomainRealmUser:      "",
			expectedGetSearchDomainRealmPassword:  "",
			expectedHasCustomNodesDNS:             false,
			expectedGetHyperkubeImageReference:    "hyperkube-amd64:v1.15.4",
			expectedGetTargetEnvironment:          "AzurePublicCloud",
			expectedIsNSeriesSKU:                  false,
			expectedIsDockerContainerRuntime:      true,
			expectedHasPrivateAzureRegistryServer: true,
			expectedGetPrivateAzureRegistryServer: "my-server",
			expectedGetSysctlDConfigKeyVals:       "",
			expectedIsVirtualMachineScaleSets:     true,
		},
		{
			name: "cluster-init config",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.17.0-beta.1",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime:        api.Docker,
							KubernetesImageBaseType: common.KubernetesImageBaseTypeGCR,
							Components: []api.KubernetesComponent{
								{
									Name:    common.ClusterInitComponentName,
									Enabled: to.BoolPtr(true),
								},
							},
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
						},
					},
				},
			},
			expectedHasCustomSearchDomain:        false,
			expectedGetSearchDomainName:          "",
			expectedGetSearchDomainRealmUser:     "",
			expectedGetSearchDomainRealmPassword: "",
			expectedHasCustomNodesDNS:            false,
			expectedGetHyperkubeImageReference:   "",
			expectedGetTargetEnvironment:         "AzurePublicCloud",
			expectedIsNSeriesSKU:                 false,
			expectedIsDockerContainerRuntime:     true,
			expectedGetSysctlDConfigKeyVals:      "",
			expectedHasClusterInitComponent:      true,
			expectedIsVirtualMachineScaleSets:    true,
		},
		{
			name: "sysctl config",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.16.1",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime:        api.Docker,
							KubernetesImageBaseType: common.KubernetesImageBaseTypeGCR,
						},
					},
					MasterProfile: &api.MasterProfile{
						SysctlDConfig: map[string]string{
							"net.ipv4.tcp_retries2":             "8",
							"net.core.somaxconn":                "16384",
							"net.ipv4.tcp_max_syn_backlog":      "16384",
							"net.core.message_cost":             "40",
							"net.core.message_burst":            "80",
							"net.ipv4.neigh.default.gc_thresh1": "4096",
							"net.ipv4.neigh.default.gc_thresh2": "8192",
							"net.ipv4.neigh.default.gc_thresh3": "16384",
							"net.ipv4.tcp_keepalive_time":       "7200",
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
							SysctlDConfig: map[string]string{
								"net.ipv4.tcp_retries2":             "8",
								"net.core.somaxconn":                "16384",
								"net.ipv4.tcp_max_syn_backlog":      "16384",
								"net.core.message_cost":             "40",
								"net.core.message_burst":            "80",
								"net.ipv4.neigh.default.gc_thresh1": "4096",
								"net.ipv4.neigh.default.gc_thresh2": "8192",
								"net.ipv4.neigh.default.gc_thresh3": "16384",
								"net.ipv4.tcp_keepalive_time":       "7200",
							},
						},
					},
				},
			},
			expectedHasCustomSearchDomain:        false,
			expectedGetSearchDomainName:          "",
			expectedGetSearchDomainRealmUser:     "",
			expectedGetSearchDomainRealmPassword: "",
			expectedHasCustomNodesDNS:            false,
			expectedGetHyperkubeImageReference:   "hyperkube-amd64:v1.16.1",
			expectedGetTargetEnvironment:         "AzurePublicCloud",
			expectedIsNSeriesSKU:                 false,
			expectedIsDockerContainerRuntime:     true,
			expectedIsVirtualMachineScaleSets:    true,
			expectedGetSysctlDConfigKeyVals: `net.core.message_burst = 80
    net.core.message_cost = 40
    net.core.somaxconn = 16384
    net.ipv4.neigh.default.gc_thresh1 = 4096
    net.ipv4.neigh.default.gc_thresh2 = 8192
    net.ipv4.neigh.default.gc_thresh3 = 16384
    net.ipv4.tcp_keepalive_time = 7200
    net.ipv4.tcp_max_syn_backlog = 16384
    net.ipv4.tcp_retries2 = 8`,
		},
		{
			name: "1.17 release w/ kube-reserved cgroup",
			cs: &api.ContainerService{
				Properties: &api.Properties{
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    api.Kubernetes,
						OrchestratorVersion: "1.17.0-beta.1",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime:        api.Docker,
							KubernetesImageBaseType: common.KubernetesImageBaseTypeGCR,
							KubeReservedCgroup:      "kubereserved",
						},
					},
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:                "pool1",
							Count:               1,
							AvailabilityProfile: api.VirtualMachineScaleSets,
							Distro:              api.AKSUbuntu1604,
						},
					},
				},
			},
			expectedHasCustomSearchDomain:        false,
			expectedGetSearchDomainName:          "",
			expectedGetSearchDomainRealmUser:     "",
			expectedGetSearchDomainRealmPassword: "",
			expectedHasCustomNodesDNS:            false,
			expectedGetHyperkubeImageReference:   "",
			expectedGetTargetEnvironment:         "AzurePublicCloud",
			expectedIsNSeriesSKU:                 false,
			expectedIsDockerContainerRuntime:     true,
			expectedGetSysctlDConfigKeyVals:      "",
			expectedHasVHDDistroNodes:            true,
			expectedIsVHDDistroForAllNodes:       true,
			expectedIsVirtualMachineScaleSets:    true,
			expectedHasKubeReservedCgroup:        true,
			expectedGetKubeReservedCgroup:        "kubereserved",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			funcMap := getContainerServiceFuncMap(c.cs)
			v := reflect.ValueOf(funcMap["GetSearchDomainName"])
			ret := v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedGetSearchDomainName {
				t.Errorf("expected funcMap invocation of GetSearchDomainName to return %s, instead got %s", c.expectedGetSearchDomainName, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["GetSearchDomainRealmUser"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedGetSearchDomainRealmUser {
				t.Errorf("expected funcMap invocation of GetSearchDomainRealmUser to return %s, instead got %s", c.expectedGetSearchDomainRealmUser, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["GetSearchDomainRealmPassword"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedGetSearchDomainRealmPassword {
				t.Errorf("expected funcMap invocation of GetSearchDomainRealmPassword to return %s, instead got %s", c.expectedGetSearchDomainRealmPassword, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["GetHyperkubeImageReference"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedGetHyperkubeImageReference {
				t.Errorf("expected funcMap invocation of GetHyperkubeImageReference to return %s, instead got %s", c.expectedGetHyperkubeImageReference, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["HasCustomNodesDNS"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedHasCustomNodesDNS {
				t.Errorf("expected funcMap invocation of HasCustomNodesDNS to return %t, instead got %t", c.expectedHasCustomNodesDNS, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["GetTargetEnvironment"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedGetTargetEnvironment {
				t.Errorf("expected funcMap invocation of GetTargetEnvironment to return %s, instead got %s", c.expectedGetTargetEnvironment, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["GetCustomCloudConfigCSEScriptFilepath"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != customCloudConfigCSEScriptFilepath {
				t.Errorf("expected funcMap invocation of GetCustomCloudConfigCSEScriptFilepath to return %s, instead got %s", customCloudConfigCSEScriptFilepath, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["GetCSEHelpersScriptFilepath"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != cseHelpersScriptFilepath {
				t.Errorf("expected funcMap invocation of GetCSEHelpersScriptFilepath to return %s, instead got %s", cseHelpersScriptFilepath, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["GetCSEInstallScriptFilepath"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != cseInstallScriptFilepath {
				t.Errorf("expected funcMap invocation of GetCSEInstallScriptFilepath to return %s, instead got %s", cseInstallScriptFilepath, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["GetCSEConfigScriptFilepath"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != cseConfigScriptFilepath {
				t.Errorf("expected funcMap invocation of GetCSEConfigScriptFilepath to return %s, instead got %s", cseConfigScriptFilepath, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["GetCustomSearchDomainsCSEScriptFilepath"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != customSearchDomainsCSEScriptFilepath {
				t.Errorf("expected funcMap invocation of GetCustomSearchDomainsCSEScriptFilepath to return %s, instead got %s", customSearchDomainsCSEScriptFilepath, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["GetDHCPv6ServiceCSEScriptFilepath"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != dhcpV6ServiceCSEScriptFilepath {
				t.Errorf("expected funcMap invocation of GetDHCPv6ServiceCSEScriptFilepath to return %s, instead got %s", dhcpV6ServiceCSEScriptFilepath, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["GetDHCPv6ConfigCSEScriptFilepath"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != dhcpV6ConfigCSEScriptFilepath {
				t.Errorf("expected funcMap invocation of GetDHCPv6ConfigCSEScriptFilepath to return %s, instead got %s", dhcpV6ConfigCSEScriptFilepath, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["HasVHDDistroNodes"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedHasVHDDistroNodes {
				t.Errorf("expected funcMap invocation of HasVHDDistroNodes to return %t, instead got %t", c.expectedHasVHDDistroNodes, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["IsVHDDistroForAllNodes"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedIsVHDDistroForAllNodes {
				t.Errorf("expected funcMap invocation of IsVHDDistroForAllNodes to return %t, instead got %t", c.expectedIsVHDDistroForAllNodes, ret[0].Interface())
			}
			if len(c.cs.Properties.AgentPoolProfiles) > 0 {
				v = reflect.ValueOf(funcMap["IsNSeriesSKU"])
				ret = v.Call([]reflect.Value{reflect.ValueOf(c.cs.Properties.AgentPoolProfiles[0].VMSize)})
				if ret[0].Interface() != c.expectedIsNSeriesSKU {
					t.Errorf("expected funcMap invocation of IsNSeriesSKU to return %t, instead got %t", c.expectedIsNSeriesSKU, ret[0].Interface())
				}
			}
			v = reflect.ValueOf(funcMap["IsDockerContainerRuntime"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedIsDockerContainerRuntime {
				t.Errorf("expected funcMap invocation of IsDockerContainerRuntime to return %t, instead got %t", c.expectedIsDockerContainerRuntime, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["HasPrivateAzureRegistryServer"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedHasPrivateAzureRegistryServer {
				t.Errorf("expected funcMap invocation of HasPrivateAzureRegistryServer to return %t, instead got %t", c.expectedHasPrivateAzureRegistryServer, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["GetPrivateAzureRegistryServer"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedGetPrivateAzureRegistryServer {
				t.Errorf("expected funcMap invocation of GetPrivateAzureRegistryServer to return %s, instead got %s", c.expectedGetPrivateAzureRegistryServer, ret[0].Interface())
			}
			if c.cs.Properties.MasterProfile != nil {
				v = reflect.ValueOf(funcMap["GetSysctlDConfigKeyVals"])
				ret = v.Call([]reflect.Value{reflect.ValueOf(c.cs.Properties.MasterProfile.SysctlDConfig)})
				if ret[0].Interface() != c.expectedGetSysctlDConfigKeyVals {
					t.Errorf("expected funcMap invocation of expectedGetSysctlDConfigKeyVals to return %s, instead got %s", c.expectedGetSysctlDConfigKeyVals, ret[0].Interface())
				}
			}
			for _, pool := range c.cs.Properties.AgentPoolProfiles {
				v = reflect.ValueOf(funcMap["GetSysctlDConfigKeyVals"])
				ret = v.Call([]reflect.Value{reflect.ValueOf(pool.SysctlDConfig)})
				if ret[0].Interface() != c.expectedGetSysctlDConfigKeyVals {
					t.Errorf("expected funcMap invocation of expectedGetSysctlDConfigKeyVals to return %s, instead got %s", c.expectedGetSysctlDConfigKeyVals, ret[0].Interface())
				}
			}
			for i, errorCodeString := range errorCodeStrings {
				v = reflect.ValueOf(funcMap["GetCSEErrorCode"])
				ret = v.Call([]reflect.Value{reflect.ValueOf(errorCodeString)})
				if ret[0].Interface() != errorCodes[i] {
					t.Errorf("expected funcMap invocation of GetCSEErrorCode to return %d, instead got %d", errorCodes[i], ret[0].Interface())
				}
			}
			v = reflect.ValueOf(funcMap["HasClusterInitComponent"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedHasClusterInitComponent {
				t.Errorf("expected funcMap invocation of HasClusterInitComponent to return %t, instead got %t", c.expectedHasClusterInitComponent, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["UseManagedIdentity"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedUseManagedIdentity {
				t.Errorf("expected funcMap invocation of UseManagedIdentity to return %t, instead got %t", c.expectedUseManagedIdentity, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["HasKubeReservedCgroup"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedHasKubeReservedCgroup {
				t.Errorf("expected funcMap invocation of HasKubeReservedCgroup to return %t, instead got %t", c.expectedIsDockerContainerRuntime, ret[0].Interface())
			}
			v = reflect.ValueOf(funcMap["GetKubeReservedCgroup"])
			ret = v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != c.expectedGetKubeReservedCgroup {
				t.Errorf("expected funcMap invocation of GetKubeReservedCgroup to return %t, instead got %t", c.expectedIsDockerContainerRuntime, ret[0].Interface())
			}
			if len(c.cs.Properties.AgentPoolProfiles) > 0 {
				v = reflect.ValueOf(funcMap["IsVirtualMachineScaleSets"])
				ret = v.Call([]reflect.Value{reflect.ValueOf(c.cs.Properties.AgentPoolProfiles[0])})
				if ret[0].Interface() != c.expectedIsVirtualMachineScaleSets {
					t.Errorf("expected funcMap invocation of IsVirtualMachineScaleSets to return %t, instead got %t", c.expectedIsVirtualMachineScaleSets, ret[0].Interface())
				}
			}
		})
	}
}

func TestTemplateGenerator_FunctionMap(t *testing.T) {
	testCases := []struct {
		Name           string
		FuncName       string
		MutateFunc     func(cs api.ContainerService) api.ContainerService
		ExpectedResult interface{}
	}{
		{
			Name:     "IsKubenet_IsTrueWhenNetworkPluginIsKubenet",
			FuncName: "IsKubenet",
			MutateFunc: func(cs api.ContainerService) api.ContainerService {
				cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = NetworkPluginKubenet
				return cs
			},
			ExpectedResult: true,
		},
		{
			Name:           "IsKubenet_IsFalseWhenNetworkPluginIsNotKubenet",
			FuncName:       "IsKubenet",
			ExpectedResult: false,
		},
		{
			Name:     "GetPodInfraContainerSpec",
			FuncName: "GetPodInfraContainerSpec",
			MutateFunc: func(cs api.ContainerService) api.ContainerService {
				cs.Properties.OrchestratorProfile.KubernetesConfig.MCRKubernetesImageBase = "foo/"
				cs.Properties.OrchestratorProfile.OrchestratorVersion = "1.16.0"
				return cs
			},
			ExpectedResult: "foo/k8s/core/pause:1.2.0",
		},
		{
			Name:     "HasCiliumNetworkPolicy - cilium",
			FuncName: "HasCiliumNetworkPolicy",
			MutateFunc: func(cs api.ContainerService) api.ContainerService {
				cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy = NetworkPolicyCilium
				return cs
			},
			ExpectedResult: true,
		},
		{
			Name:     "HasCiliumNetworkPlugin - cilium",
			FuncName: "HasCiliumNetworkPlugin",
			MutateFunc: func(cs api.ContainerService) api.ContainerService {
				cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = NetworkPluginCilium
				return cs
			},
			ExpectedResult: true,
		},
		{
			Name:     "HasCiliumNetworkPlugin - azure",
			FuncName: "HasCiliumNetworkPlugin",
			MutateFunc: func(cs api.ContainerService) api.ContainerService {
				cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = NetworkPluginAzure
				return cs
			},
			ExpectedResult: false,
		},
		{
			Name:     "HasAntreaNetworkPolicy - antrea",
			FuncName: "HasAntreaNetworkPolicy",
			MutateFunc: func(cs api.ContainerService) api.ContainerService {
				cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy = NetworkPluginAntrea
				return cs
			},
			ExpectedResult: true,
		},
		{
			Name:     "HasAntreaNetworkPolicy - azure",
			FuncName: "HasAntreaNetworkPolicy",
			MutateFunc: func(cs api.ContainerService) api.ContainerService {
				cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy = NetworkPolicyAzure
				return cs
			},
			ExpectedResult: false,
		},
		{
			Name:     "HasFlannelNetworkPlugin - flannel",
			FuncName: "HasFlannelNetworkPlugin",
			MutateFunc: func(cs api.ContainerService) api.ContainerService {
				cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = NetworkPluginFlannel
				return cs
			},
			ExpectedResult: true,
		},
		{
			Name:     "HasFlannelNetworkPlugin - azure",
			FuncName: "HasFlannelNetworkPlugin",
			MutateFunc: func(cs api.ContainerService) api.ContainerService {
				cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy = NetworkPolicyAzure
				return cs
			},
			ExpectedResult: false,
		},
		{
			Name:     "HasTelemetryEnabled",
			FuncName: "HasTelemetryEnabled",
			MutateFunc: func(cs api.ContainerService) api.ContainerService {
				cs.Properties.FeatureFlags = &api.FeatureFlags{
					EnableTelemetry: true,
				}
				return cs
			},
			ExpectedResult: true,
		},
		{
			Name:     "GetEmptyApplicationInsightsTelemetryKeys",
			FuncName: "GetApplicationInsightsTelemetryKeys",
			MutateFunc: func(cs api.ContainerService) api.ContainerService {
				cs.Properties.TelemetryProfile = nil
				return cs
			},
			ExpectedResult: telemetry.AKSEngineAppInsightsKey,
		},
		{
			Name:     "GetApplicationInsightsTelemetryKeysWithUserSuppliedKey",
			FuncName: "GetApplicationInsightsTelemetryKeys",
			MutateFunc: func(cs api.ContainerService) api.ContainerService {
				cs.Properties.TelemetryProfile = &api.TelemetryProfile{
					ApplicationInsightsKey: "my_telemetry_key",
				}
				return cs
			},
			ExpectedResult: fmt.Sprintf("%s,%s", telemetry.AKSEngineAppInsightsKey, "my_telemetry_key"),
		},
		{
			Name:     "GetLinuxDefaultTelemetryTags",
			FuncName: "GetLinuxDefaultTelemetryTags",
			MutateFunc: func(cs api.ContainerService) api.ContainerService {
				cs.Properties.OrchestratorProfile.KubernetesConfig.ContainerRuntime = "containerRuntime"
				cs.Properties.OrchestratorProfile.KubernetesConfig.ContainerdVersion = "1.2.4"
				cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkMode = "networkMode"
				cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = "networkPlugin"
				cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy = "networkPolicy"
				cs.Properties.OrchestratorProfile.OrchestratorVersion = "1.2.4"
				return cs
			},
			ExpectedResult: "cri=containerRuntime,cri_version=1.2.4,k8s_version=1.2.4,network_mode=networkMode,network_plugin=networkPlugin,network_policy=networkPolicy,os_type=linux",
		},
		{
			Name:     "GetLinuxDefaultTelemetryTagsWithEmpties",
			FuncName: "GetLinuxDefaultTelemetryTags",
			MutateFunc: func(cs api.ContainerService) api.ContainerService {
				cs.Properties.OrchestratorProfile.KubernetesConfig.ContainerRuntime = "containerRuntime"
				cs.Properties.OrchestratorProfile.KubernetesConfig.ContainerdVersion = "1.2.4"
				cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkMode = ""
				cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = "networkPlugin"
				cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy = ""
				cs.Properties.OrchestratorProfile.OrchestratorVersion = "1.2.4"
				return cs
			},
			ExpectedResult: "cri=containerRuntime,cri_version=1.2.4,k8s_version=1.2.4,network_plugin=networkPlugin,os_type=linux",
		},
	}

	for _, c := range testCases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()

			cs := api.ContainerService{}
			if err := json.Unmarshal([]byte(getAPIModelString()), &cs); err != nil {
				t.Fatalf("unexpected error unmashalling model string: %v", err)
			}

			if c.MutateFunc != nil {
				cs = c.MutateFunc(cs)
			}

			bits, err := json.Marshal(&cs)
			if err != nil {
				t.Fatalf("error marshalling model to string: %v", err)
			}

			funcmap, err := getFuncMap(string(bits))
			if err != nil {
				t.Fatalf("error generating function map: %v", err)
			}

			f, ok := funcmap[c.FuncName]
			if !ok {
				t.Fatalf("didn't find expected funcmap key %s.", c.FuncName)
			}

			v := reflect.ValueOf(f)
			rargs := make([]reflect.Value, 0)
			if ret := v.Call(rargs); ret[0].Interface() != c.ExpectedResult {
				t.Fatalf("expected %v, but got %v", c.ExpectedResult, ret[0].Interface())
			}
		})
	}

}

func TestGetIdentitySystem(t *testing.T) {
	for _, test := range []struct {
		desc           string
		apiModel       string
		expectedResult string
	}{
		{
			desc:           "should return adfs when identitySystem is set to adfs",
			apiModel:       `{"properties":{"customCloudProfile": {"identitySystem": "adfs"}}}`,
			expectedResult: "adfs",
		},
		{
			desc:           "should return azure_ad when identitySystem is set to azure_ad",
			apiModel:       `{"properties":{"customCloudProfile": {"identitySystem": "azure_ad"}}}`,
			expectedResult: "azure_ad",
		},
		{
			desc:           "should return azure_ad when not azure stack",
			apiModel:       getAPIModelString(),
			expectedResult: "azure_ad",
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			funcmap, err := getFuncMap(test.apiModel)
			if err != nil {
				t.Fatalf("Error generating function map: %v", err)
			}

			v := reflect.ValueOf(funcmap["GetIdentitySystem"])
			ret := v.Call(make([]reflect.Value, 0))
			if ret[0].Interface() != test.expectedResult {
				t.Fatalf("IsIdentitySystemADFS returned incorrect value")
			}
		})
	}
}

func TestGetDockerConfig(t *testing.T) {
	cases := []struct {
		name          string
		vmSKU         string
		runtimeConfig map[string]string
	}{
		{
			name:          "default",
			vmSKU:         "Standard_D8s_v3",
			runtimeConfig: nil,
		},
		{
			name:          "gpu",
			vmSKU:         "Standard_NC12s_v2",
			runtimeConfig: nil,
		},
		{
			name:  "reroot",
			vmSKU: "Standard_D8s_v3",
			runtimeConfig: map[string]string{
				"dataDir": "/mnt/docker",
			},
		},
		{
			name:  "all",
			vmSKU: "Standard_NC12s_v2",
			runtimeConfig: map[string]string{
				"dataDir": "/mnt/docker",
			},
		},
	}
	expectedOutputs := common.GetDockerConfigTestCases()

	for _, c := range cases {
		t.Run(fmt.Sprintf("docker %s", c.name), func(t *testing.T) {
			cs := &api.ContainerService{
				Properties: &api.Properties{
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:   "agentpool",
							VMSize: c.vmSKU,
							Count:  1,
						},
					},
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: "1.12.2",
						KubernetesConfig: &api.KubernetesConfig{
							ContainerRuntime:       api.Docker,
							ContainerRuntimeConfig: c.runtimeConfig,
						},
					},
				},
			}
			isGPU := common.IsNvidiaEnabledSKU(cs.Properties.AgentPoolProfiles[0].VMSize)
			got, err := getDockerConfig(cs, isGPU)
			if err != nil {
				t.Fatalf("Error gegetting containerd config: %v", err)

			}
			want := expectedOutputs[c.name]
			diff := cmp.Diff(want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestGetContainerdConfig(t *testing.T) {
	cases := []struct {
		name          string
		vmSKU         string
		runtimeConfig map[string]string
		kubenet       bool
	}{
		{
			name:          "default",
			vmSKU:         "Standard_D8s_v3",
			runtimeConfig: nil,
		},
		{
			name:          "kubenet",
			vmSKU:         "Standard_NC12s_v2", // No change for containerd
			runtimeConfig: nil,
			kubenet:       true,
		},
		{
			name:  "reroot",
			vmSKU: "Standard_D8s_v3",
			runtimeConfig: map[string]string{
				"dataDir": "/mnt/containerd",
			},
		},
		{
			name:  "all",
			vmSKU: "Standard_NC12s_v2",
			runtimeConfig: map[string]string{
				"dataDir": "/mnt/containerd",
			},
			kubenet: true,
		},
	}
	expectedOutputs := common.GetContainerdConfigTestCases()

	for _, c := range cases {
		t.Run(fmt.Sprintf("containerd %s", c.name), func(t *testing.T) {
			cs := &api.ContainerService{
				Properties: &api.Properties{
					AgentPoolProfiles: []*api.AgentPoolProfile{
						{
							Name:   "agentpool",
							VMSize: c.vmSKU,
							Count:  1,
						},
					},
					OrchestratorProfile: &api.OrchestratorProfile{
						OrchestratorType:    Kubernetes,
						OrchestratorVersion: "1.16.0",
						KubernetesConfig: &api.KubernetesConfig{
							MCRKubernetesImageBase: "foo/",
							ContainerRuntime:       api.Containerd,
							ContainerRuntimeConfig: c.runtimeConfig,
						},
					},
				},
			}
			if c.kubenet {
				cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPlugin = NetworkPluginKubenet
			}

			got, err := getContainerdConfig(cs)
			if err != nil {
				t.Fatalf("Error gegetting containerd config: %v", err)

			}
			want := expectedOutputs[c.name]
			diff := cmp.Diff(want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestGetBase64EncodedEnvironmentJSON(t *testing.T) {
	apiModelString := `{"properties":{"customCloudProfile":{"environment":{"name":"AzureStackCloud","managementPortalURL":"https://portal.local.azurestack.external/","publishSettingsURL":"'single quotes'","serviceManagementEndpoint":"https://management.azurestack.onmicrosoft.com/00000000-0000-0000-0000-0000000000","resourceManagerEndpoint":"https://management.local.azurestack.external/","activeDirectoryEndpoint":"https://login.microsoftonline.com/","galleryEndpoint":"https://galleryartifacts.hosting.local.azurestack.external/galleryartifacts/","keyVaultEndpoint":"","graphEndpoint":"https://graph.windows.net/","serviceBusEndpoint":"","batchManagementEndpoint":"","storageEndpointSuffix":"local.azurestack.external","sqlDatabaseDNSSuffix":"","trafficManagerDNSSuffix":"","keyVaultDNSSuffix":"vault.local.azurestack.external","serviceBusEndpointSuffix":"","serviceManagementVMDNSSuffix":"","resourceManagerVMDNSSuffix":"cloudapp.azurestack.external","containerRegistryDNSSuffix":"","tokenAudience":""}}}}`
	funcmap, err := getFuncMap(apiModelString)
	if err != nil {
		t.Fatalf("Error generating function map: %v", err)
	}
	v := reflect.ValueOf(funcmap["GetBase64EncodedEnvironmentJSON"])
	ret := v.Call(make([]reflect.Value, 0))
	encodedEnvironmentJSON := ret[0].String()
	correctlyEncodedString := `eyJuYW1lIjoiQXp1cmVTdGFja0Nsb3VkIiwibWFuYWdlbWVudFBvcnRhbFVSTCI6Imh0dHBzOi8vcG9ydGFsLmxvY2FsLmF6dXJlc3RhY2suZXh0ZXJuYWwvIiwicHVibGlzaFNldHRpbmdzVVJMIjoiJ3NpbmdsZSBxdW90ZXMnIiwic2VydmljZU1hbmFnZW1lbnRFbmRwb2ludCI6Imh0dHBzOi8vbWFuYWdlbWVudC5henVyZXN0YWNrLm9ubWljcm9zb2Z0LmNvbS8wMDAwMDAwMC0wMDAwLTAwMDAtMDAwMC0wMDAwMDAwMDAwIiwicmVzb3VyY2VNYW5hZ2VyRW5kcG9pbnQiOiJodHRwczovL21hbmFnZW1lbnQubG9jYWwuYXp1cmVzdGFjay5leHRlcm5hbC8iLCJhY3RpdmVEaXJlY3RvcnlFbmRwb2ludCI6Imh0dHBzOi8vbG9naW4ubWljcm9zb2Z0b25saW5lLmNvbS8iLCJnYWxsZXJ5RW5kcG9pbnQiOiJodHRwczovL2dhbGxlcnlhcnRpZmFjdHMuaG9zdGluZy5sb2NhbC5henVyZXN0YWNrLmV4dGVybmFsL2dhbGxlcnlhcnRpZmFjdHMvIiwia2V5VmF1bHRFbmRwb2ludCI6IiIsImdyYXBoRW5kcG9pbnQiOiJodHRwczovL2dyYXBoLndpbmRvd3MubmV0LyIsInNlcnZpY2VCdXNFbmRwb2ludCI6IiIsImJhdGNoTWFuYWdlbWVudEVuZHBvaW50IjoiIiwic3RvcmFnZUVuZHBvaW50U3VmZml4IjoibG9jYWwuYXp1cmVzdGFjay5leHRlcm5hbCIsInNxbERhdGFiYXNlRE5TU3VmZml4IjoiIiwidHJhZmZpY01hbmFnZXJETlNTdWZmaXgiOiIiLCJrZXlWYXVsdEROU1N1ZmZpeCI6InZhdWx0LmxvY2FsLmF6dXJlc3RhY2suZXh0ZXJuYWwiLCJzZXJ2aWNlQnVzRW5kcG9pbnRTdWZmaXgiOiIiLCJzZXJ2aWNlTWFuYWdlbWVudFZNRE5TU3VmZml4IjoiIiwicmVzb3VyY2VNYW5hZ2VyVk1ETlNTdWZmaXgiOiJjbG91ZGFwcC5henVyZXN0YWNrLmV4dGVybmFsIiwiY29udGFpbmVyUmVnaXN0cnlETlNTdWZmaXgiOiIiLCJjb3Ntb3NEQkROU1N1ZmZpeCI6IiIsInRva2VuQXVkaWVuY2UiOiIiLCJyZXNvdXJjZUlkZW50aWZpZXJzIjp7ImdyYXBoIjoiIiwia2V5VmF1bHQiOiIiLCJkYXRhbGFrZSI6IiIsImJhdGNoIjoiIiwib3BlcmF0aW9uYWxJbnNpZ2h0cyI6IiIsInN0b3JhZ2UiOiIifX0=`

	if encodedEnvironmentJSON != correctlyEncodedString {
		t.Fatalf("Function GetEnvironmentJSON() failed to produce correct environment: expected: %s, actual: %s", correctlyEncodedString, encodedEnvironmentJSON)
	}
}

func getFuncMap(apiModelStr string) (template.FuncMap, error) {
	tg, err := InitializeTemplateGenerator(Context{})
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error while instantiating TemplateGenerator")
	}
	cs := &api.ContainerService{}
	err = json.Unmarshal([]byte(apiModelStr), &cs)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error while unmarshalling the apiModel JSON")
	}
	return tg.getTemplateFuncMap(cs), nil
}

func getAPIModelString() string {
	return `{"id":"","location":"","name":"","tags":{},"type":"","properties":{"ClusterID":"","orchestratorProfile":{"orchestratorType":"Kubernetes","orchestratorVersion":"1.10.12","kubernetesConfig":{"kubernetesImageBase":"k8s.gcr.io/","clusterSubnet":"10.240.0.0/12","networkPlugin":"azure","containerRuntime":"docker","dockerBridgeSubnet":"172.17.0.1/16","dnsServiceIP":"10.0.0.10","serviceCidr":"10.0.0.0/16","useInstanceMetadata":true,"enableRbac":true,"enableSecureKubelet":true,"enableAggregatedAPIs":true,"privateCluster":{"enabled":false},"gchighthreshold":85,"gclowthreshold":80,"etcdVersion":"3.2.26","etcdDiskSizeGB":"256","addons":[{"name":"heapster","enabled":true,"containers":[{"name":"heapster","image":"k8s.gcr.io/heapster-amd64:v1.5.4"},{"name":"heapster-nanny","image":"k8s.gcr.io/addon-resizer:1.8.5"}]},{"name":"tiller","enabled":true,"containers":[{"name":"tiller","image":"gcr.io/kubernetes-helm/tiller:v2.11.0","cpuRequests":"50m","memoryRequests":"150Mi","cpuLimits":"50m","memoryLimits":"150Mi"}],"config":{"max-history":"0"}},{"name":"aci-connector","enabled":false,"containers":[{"name":"aci-connector","image":"microsoft/virtual-kubelet:latest","cpuRequests":"50m","memoryRequests":"150Mi","cpuLimits":"50m","memoryLimits":"150Mi"}],"config":{"nodeName":"aci-connector","os":"Linux","region":"westus","taint":"azure.com/aci"}},{"name":"cluster-autoscaler","enabled":false,"containers":[{"name":"cluster-autoscaler","image":"k8s.gcr.io/cluster-autoscaler:v1.2.2","cpuRequests":"100m","memoryRequests":"300Mi","cpuLimits":"100m","memoryLimits":"300Mi"}],"config":{"max-nodes":"5","min-nodes":"1","scan-interval":"10s"}},{"name":"blobfuse-flexvolume","enabled":true,"containers":[{"name":"blobfuse-flexvolume","image":"mcr.microsoft.com/k8s/flexvolume/blobfuse-flexvolume:1.0.7","cpuRequests":"50m","memoryRequests":"10Mi","cpuLimits":"50m","memoryLimits":"10Mi"}]},{"name":"smb-flexvolume","enabled":false,"containers":[{"name":"smb-flexvolume","image":"mcr.microsoft.com/k8s/flexvolume/smb-flexvolume:1.0.2","cpuRequests":"50m","memoryRequests":"10Mi","cpuLimits":"50m","memoryLimits":"10Mi"}]},{"name":"keyvault-flexvolume","enabled":true,"containers":[{"name":"keyvault-flexvolume","image":"mcr.microsoft.com/k8s/flexvolume/keyvault-flexvolume:v0.0.5","cpuRequests":"50m","memoryRequests":"10Mi","cpuLimits":"50m","memoryLimits":"10Mi"}]},{"name":"kubernetes-dashboard","enabled":true,"containers":[{"name":"kubernetes-dashboard","image":"k8s.gcr.io/kubernetes-dashboard-amd64:v1.10.1","cpuRequests":"300m","memoryRequests":"150Mi","cpuLimits":"300m","memoryLimits":"150Mi"}]},{"name":"rescheduler","enabled":false,"containers":[{"name":"rescheduler","image":"k8s.gcr.io/rescheduler:v0.3.1","cpuRequests":"10m","memoryRequests":"100Mi","cpuLimits":"10m","memoryLimits":"100Mi"}]},{"name":"metrics-server","enabled":true,"containers":[{"name":"metrics-server","image":"k8s.gcr.io/metrics-server-amd64:v0.2.1"}]},{"name":"nvidia-device-plugin","enabled":false,"containers":[{"name":"nvidia-device-plugin","image":"nvidia/k8s-device-plugin:1.10","cpuRequests":"50m","memoryRequests":"10Mi","cpuLimits":"50m","memoryLimits":"10Mi"}]},{"name":"container-monitoring","enabled":false,"containers":[{"name":"omsagent","image":"microsoft/oms:ciprod11292018","cpuRequests":"50m","memoryRequests":"200Mi","cpuLimits":"150m","memoryLimits":"750Mi"}],"config":{"dockerProviderVersion":"2.0.0-3","omsAgentVersion":"1.6.0-42"}},{"name":"azure-cni-networkmonitor","enabled":true,"containers":[{"name":"azure-cni-networkmonitor","image":"mcr.microsoft.com/containernetworking/networkmonitor:v0.0.4"}]},{"name":"azure-npm-daemonset","enabled":false,"containers":[{"name":"azure-npm-daemonset"}]},{"name":"ip-masq-agent","enabled":true,"containers":[{"name":"ip-masq-agent","image":"k8s.gcr.io/ip-masq-agent-amd64:v2.3.0","cpuRequests":"50m","memoryRequests":"50Mi","cpuLimits":"50m","memoryLimits":"250Mi"}],"config":{"non-masq-cni-cidr":"168.63.129.16/32","non-masquerade-cidr":"10.0.0.0/8"}}],"kubeletConfig":{"--address":"0.0.0.0","--allow-privileged":"true","--anonymous-auth":"false","--authorization-mode":"Webhook","--azure-container-registry-config":"/etc/kubernetes/azure.json","--cadvisor-port":"0","--cgroups-per-qos":"true","--client-ca-file":"/etc/kubernetes/certs/ca.crt","--cloud-config":"/etc/kubernetes/azure.json","--cloud-provider":"azure","--cluster-dns":"10.0.0.10","--cluster-domain":"cluster.local","--enforce-node-allocatable":"pods","--event-qps":"0","--eviction-hard":"memory.available\u003c100Mi,nodefs.available\u003c10%,nodefs.inodesFree\u003c5%","--feature-gates":"PodPriority=true","--image-gc-high-threshold":"85","--image-gc-low-threshold":"80","--image-pull-progress-deadline":"30m","--keep-terminated-pod-volumes":"false","--kubeconfig":"/var/lib/kubelet/kubeconfig","--max-pods":"30","--network-plugin":"cni","--node-status-update-frequency":"10s","--non-masquerade-cidr":"0.0.0.0/0","--pod-infra-container-image":"k8s.gcr.io/pause-amd64:3.1","--pod-manifest-path":"/etc/kubernetes/manifests","--pod-max-pids":"-1"},"controllerManagerConfig":{"--allocate-node-cidrs":"false","--cloud-config":"/etc/kubernetes/azure.json","--cloud-provider":"azure","--cluster-cidr":"10.240.0.0/12","--cluster-name":"blueorange","--cluster-signing-cert-file":"/etc/kubernetes/certs/ca.crt","--cluster-signing-key-file":"/etc/kubernetes/certs/ca.key","--configure-cloud-routes":"false","--controllers":"*,bootstrapsigner,tokencleaner","--feature-gates":"LocalStorageCapacityIsolation=true,ServiceNodeExclusion=true","--kubeconfig":"/var/lib/kubelet/kubeconfig","--leader-elect":"true","--node-monitor-grace-period":"40s","--pod-eviction-timeout":"5m0s","--profiling":"false","--root-ca-file":"/etc/kubernetes/certs/ca.crt","--route-reconciliation-period":"10s","--service-account-private-key-file":"/etc/kubernetes/certs/apiserver.key","--terminated-pod-gc-threshold":"5000","--use-service-account-credentials":"true","--v":"2"},"cloudControllerManagerConfig":{"--allocate-node-cidrs":"false","--cloud-config":"/etc/kubernetes/azure.json","--cloud-provider":"azure","--cluster-cidr":"10.240.0.0/12","--cluster-name":"blueorange","--configure-cloud-routes":"false","--kubeconfig":"/var/lib/kubelet/kubeconfig","--leader-elect":"true","--route-reconciliation-period":"10s","--v":"2"},"apiServerConfig":{"--advertise-address":"\u003cadvertiseAddr\u003e","--allow-privileged":"true","--anonymous-auth":"false","--audit-log-maxage":"30","--audit-log-maxbackup":"10","--audit-log-maxsize":"100","--audit-log-path":"/var/log/kubeaudit/audit.log","--audit-policy-file":"/etc/kubernetes/addons/audit-policy.yaml","--authorization-mode":"Node,RBAC","--bind-address":"0.0.0.0","--client-ca-file":"/etc/kubernetes/certs/ca.crt","--cloud-config":"/etc/kubernetes/azure.json","--cloud-provider":"azure","--enable-admission-plugins":"NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota,ExtendedResourceToleration","--enable-bootstrap-token-auth":"true","--etcd-cafile":"/etc/kubernetes/certs/ca.crt","--etcd-certfile":"/etc/kubernetes/certs/etcdclient.crt","--etcd-keyfile":"/etc/kubernetes/certs/etcdclient.key","--etcd-servers":"https://\u003cetcdEndPointUri\u003e:2379","--insecure-port":"8080","--kubelet-client-certificate":"/etc/kubernetes/certs/client.crt","--kubelet-client-key":"/etc/kubernetes/certs/client.key","--profiling":"false","--proxy-client-cert-file":"/etc/kubernetes/certs/proxy.crt","--proxy-client-key-file":"/etc/kubernetes/certs/proxy.key","--repair-malformed-updates":"false","--requestheader-allowed-names":"","--requestheader-client-ca-file":"/etc/kubernetes/certs/proxy-ca.crt","--requestheader-extra-headers-prefix":"X-Remote-Extra-","--requestheader-group-headers":"X-Remote-Group","--requestheader-username-headers":"X-Remote-User","--secure-port":"443","--service-account-key-file":"/etc/kubernetes/certs/apiserver.key","--service-account-lookup":"true","--service-cluster-ip-range":"10.0.0.0/16","--storage-backend":"etcd3","--tls-cert-file":"/etc/kubernetes/certs/apiserver.crt","--tls-private-key-file":"/etc/kubernetes/certs/apiserver.key","--v":"4"},"schedulerConfig":{"--kubeconfig":"/var/lib/kubelet/kubeconfig","--leader-elect":"true","--profiling":"false","--v":"2"},"cloudProviderBackoff":true,"cloudProviderBackoffRetries":6,"cloudProviderBackoffJitter":1,"cloudProviderBackoffDuration":5,"cloudProviderBackoffExponent":1.5,"cloudProviderRateLimit":true,"cloudProviderRateLimitQPS":3,"cloudProviderRateLimitBucket":10,"loadBalancerSku":"Basic","azureCNIVersion":"v1.0.15","maximumLoadBalancerRuleCount":250}},"masterProfile":{"count":1,"dnsPrefix":"blueorange","subjectAltNames":null,"vmSize":"Standard_D2_v2","firstConsecutiveStaticIP":"10.255.255.5","subnet":"10.240.0.0/12","ipAddressCount":31,"storageProfile":"ManagedDisks","HTTPSourceAddressPrefix":"*","oauthEnabled":false,"preProvisionExtension":null,"extensions":[],"distro":"aks-ubuntu-16.04","kubernetesConfig":{"kubeletConfig":{"--address":"0.0.0.0","--allow-privileged":"true","--anonymous-auth":"false","--authorization-mode":"Webhook","--azure-container-registry-config":"/etc/kubernetes/azure.json","--cadvisor-port":"0","--cgroups-per-qos":"true","--client-ca-file":"/etc/kubernetes/certs/ca.crt","--cloud-config":"/etc/kubernetes/azure.json","--cloud-provider":"azure","--cluster-dns":"10.0.0.10","--cluster-domain":"cluster.local","--enforce-node-allocatable":"pods","--event-qps":"0","--eviction-hard":"memory.available\u003c100Mi,nodefs.available\u003c10%,nodefs.inodesFree\u003c5%","--feature-gates":"PodPriority=true","--image-gc-high-threshold":"85","--image-gc-low-threshold":"80","--image-pull-progress-deadline":"30m","--keep-terminated-pod-volumes":"false","--kubeconfig":"/var/lib/kubelet/kubeconfig","--max-pods":"30","--network-plugin":"cni","--node-status-update-frequency":"10s","--non-masquerade-cidr":"0.0.0.0/0","--pod-infra-container-image":"k8s.gcr.io/pause-amd64:3.1","--pod-manifest-path":"/etc/kubernetes/manifests","--pod-max-pids":"-1"}},"availabilityProfile":"AvailabilitySet","cosmosEtcd":false},"agentPoolProfiles":[{"name":"agentpool1","count":2,"vmSize":"Standard_D2_v2","osType":"Linux","availabilityProfile":"VirtualMachineScaleSets","storageProfile":"ManagedDisks","subnet":"10.240.0.0/12","ipAddressCount":31,"distro":"aks-ubuntu-16.04","acceleratedNetworkingEnabled":true,"acceleratedNetworkingEnabledWindows":false,"preProvisionExtension":null,"extensions":[],"kubernetesConfig":{"kubeletConfig":{"--address":"0.0.0.0","--allow-privileged":"true","--anonymous-auth":"false","--authorization-mode":"Webhook","--azure-container-registry-config":"/etc/kubernetes/azure.json","--cadvisor-port":"0","--cgroups-per-qos":"true","--client-ca-file":"/etc/kubernetes/certs/ca.crt","--cloud-config":"/etc/kubernetes/azure.json","--cloud-provider":"azure","--cluster-dns":"10.0.0.10","--cluster-domain":"cluster.local","--enforce-node-allocatable":"pods","--event-qps":"0","--eviction-hard":"memory.available\u003c100Mi,nodefs.available\u003c10%,nodefs.inodesFree\u003c5%","--feature-gates":"PodPriority=true","--image-gc-high-threshold":"85","--image-gc-low-threshold":"80","--image-pull-progress-deadline":"30m","--keep-terminated-pod-volumes":"false","--kubeconfig":"/var/lib/kubelet/kubeconfig","--max-pods":"30","--network-plugin":"cni","--node-status-update-frequency":"10s","--non-masquerade-cidr":"0.0.0.0/0","--pod-infra-container-image":"k8s.gcr.io/pause-amd64:3.1","--pod-manifest-path":"/etc/kubernetes/manifests","--pod-max-pids":"-1"}},"singlePlacementGroup":true}],"linuxProfile":{"adminUsername":"azureuser","ssh":{"publicKeys":[{"keyData":"ssh-rsa AAAAB3NzaC1P+QQRd1hvsNa6/qEW+L7RDYXOF7CkxMogZFolV809sadaslkjlNMJ4p8eSs/MTl8iUqOvGaa0PXth1bdltj7WQYEjMuGYfMOFCIQBy4ZwHHx0tuLf0N0WKGn+dZ6XZJvK+SBV8FOWVekTOVn1oR51/a4RR0E6elwIThghREixQ0QRPM6oR2gNEp9GPbzqEEjxErYHks88AbnGNS0Myr9e3Jm4Xto811q0XFg3WHE3XFmURkIU7kvOmBF7MW/vltxub4MfQezOptLx0+FGF6Gc7TScIciHXqaSQ5qKumeAUp/1Bk7i6nZRPij37GjoMR34oPxcrM9BJHbw37w6Nj0RCbMUOLPHJ3mSvzcuzO2bdDWpV/prtkBcDvhTvX5ABJIZ4e1Q0cuJAnwau2TyO/iUCatvXBGkBKMTNnxAZJkB+GIDW+aqwq/grYN6Oqi8+u+6sSpWbKlET2tq7rv0sPwMeMfwNwuLVpcKHXL8Ho+YsEB2ocVXKbTY0X5sYny3QOsyRzrooM6kIAGjiA1WVQ== cyz.abc@def.com"}]}},"extensionProfiles":[],"servicePrincipalProfile":{"clientId":"f877da1d-2247-4dd8-9932-7c4a9fc45e7f","secret":"PwEV@QG7/PYt\"re9"},"certificateProfile":{"caCertificate":"-----BEGIN CERTIFICATE-----\nMIIEyDCCArCgAwIBAgIRAMzjrJNUV/NNz9QhxTxQnQ4wDQYJKoZIhvcNAQELBQAw\nDTELMAkGA1UEAxMCY2EwHhcNMTkwMTEyMDQ0MjA3WhcNMjEwMTExMDQ0MjA3WjAN\nMQswCQYDVQQDEwJjYTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBAMtr\n1WzyfwQn5ehM2ig8RHfNjrNLyqVKLBD3+UU8ZCK9T7OlbDUTxg5LG1shRtSGE+um\nQnemcJ415n+pRPdiGn1BMoSGmsDONNZYRhUgeAWpR+Zo4gHBtKNZHELoaJRx3I+K\nhZrr4HnxclHvesvD5/v5J/PRc1e+7m1wGTzIohwbsSanYqcQxH2F1gENNUQscQKV\nOl1s1q6yrescouCQE/h9cqZ2TMdFPshGE2nn3HsI+hVX19j/xZUFtTple1ulfw+7\nduR0fdMPz0fPLhwTIoE9pkDXCpf41lQR0F39qEktRbLVkcxhUZtpU3KOWvzze5iL\nCQErb6/d5ArkQO1rRVextSWkZJrMx8GRdqsiBZ4gKhRdVJ0+nZOpdfr81aoNnkVR\n31lWbbS3l+yTj0PTQ4Hs+QW8rG5k8SDoi6k9Ii5qPvJE55EJ0Drw6Udjr0NrwTF5\nLqLLrQ87uE7xZZ0l1zp8xH3DU77/jxk6V47rhMG9+UySIppeVpwRHqazyg5oQ1O6\n+HVrbT/N1bHzrP2LHH/+nsf8FrBOgi1TO0PDR7xRggICXgl6vmU62W+6zapyKhX0\n9I8U87mMjIB3h0g7RhJRNIHZHVxKHhLdE7/882XdBF+DW8/Qi/KmcRmAziLFDRaN\nb1xCNhKWRflg6nUswl6IPk+osRBVfCoSXi7QnL8/AgMBAAGjIzAhMA4GA1UdDwEB\n/wQEAwICpDAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4ICAQA2uRac\nELAMvc/bfMqpHzj7i4rEACdGFHJnuN7PE/MCLoZM44wqEzhujsNLwboXMOiBo7dd\nx7JNQcff0m3ynZkZDBSVksWY1AOcXZgyX/Uyewu3faIz+wdxe+cEq7kxGBwCB4VY\njJdDu0dNsf3K0+UvQNyOKw2S1p0nSoFSeCG63bXwq5O7f1+B01TRJr8cf6zxUxWt\nvf6q35xKx/fd3FYKoDHCZFJhHXDmV52u3oYHs/7GtGuZGsGemijLjSY5VMU2+Fnw\n730I25f/VJjpOzV01lTSl0yajz6JS+nVV4HsXd4fmGgB4lDnIUTyKuVZZq12Lyzt\nfu3Hy0bkjebeI72IumfRH5UpQSUgEFkh3Ul5Pao/SkCqUswm2Ew+FDNMUQBU8669\n4bmCFFxpk+eTPdHGHzFftH4TqI22Ct5ruw5EP40X1c9pVtZyToK7a/uAYqy5LLM0\noAcjjqKhZdrSmv3XYlxqCtGYxRssING2+GUklnR1Ax9ko7uuJE2dpFikW3lydUwB\nAu8Oa22czHeQ2NT2aefrzR2z4KjZfvjy95rukTmR2m/onJb8xiIXbz7fRCmZehMp\nxffttqCImYqifCDkht+9E+AlM/s4zZVmQF93Fd+NGod4fH4Bb1PenukdGMQEo/1m\ndAoDCDxdIDHNk4uHDRoDp+46h+uPxvRX2b9qJw==\n-----END CERTIFICATE-----\n","caPrivateKey":"-----BEGIN RSA PRIVATE KEY-----\n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\n-----END RSA PRIVATE KEY-----\n","apiServerCertificate":"-----BEGIN CERTIFICATE-----\nMIIOHzCCDAegAwIBAgIQAeMlHla/wGPSzCs5AYk8kTANBgkqhkiG9w0BAQsFADAN\nMQswCQYDVQQDEwJjYTAeFw0xOTAxMTIwNDQyMDhaFw0yMTAxMTEwNDQyMDhaMBQx\nEjAQBgNVBAMTCWFwaXNlcnZlcjCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoC\nggIBAJ1bkKxUBVD+XwW4J/T0abvi3ReTcewqS4urdI39zZbgjgVuoXdN+1d5EWQD\nRvoHdKS4Is/CYNyrmMbsF3QilPsn1E9iNQC7CoPLPslXnsr6bWPNtGa3D5899Ogx\n055uyJ49IE1xIMJnCkdtxr6CP4lQoPG8EzBMVxuyjCkxD9JGVfv/c7yXdu89I/FA\n6NtZOaLA0u/PDUdvHUdYShgmcliEAA0Ttt+/J+7F5wVZbJMpvEJLwNNKfKRe6XER\npHuIu9xw28Dxs9hTSN3rEeoMp5/CPooB32D1ve0ESSK5qJewmWIm+Q+W0uFjjsZr\nlzRFtiqyaXgi/b12xxMkLjh3QkQpkmXoiJiJARWfXllJ+v/bpjgjLJ3wt0d3cGry\nJ/XCpXGSoxf0ooaNYUHR0d66EyjQzxm7van1dLFfX3oryoqeBCem406bssvzb8w9\nzEQ5adJ0kraSxEgemry2/wGjHJv6EKvl5fS912YOkTlZrpipi4Yw/aR4qG0wdYCR\nk/49wITJ59Btk7jbLoQNoUOLjwTF8ObgBsF+odRiuRxfvd1FT6MLkDmAWm86XlFB\n2NI4e0RB9GyqMZV9basMJVlRY8RaQxxdOard/DzVW14kw1wg6SQ2rQrHoJ9pzmQ7\nYNijrNsCfAaslPtNomZvni4o3cNDQ9MVKT4d1tIRlclkAeiTAgMBAAGjgglyMIIJ\nbjAOBgNVHQ8BAf8EBAMCBaAwEwYDVR0lBAwwCgYIKwYBBQUHAwEwDAYDVR0TAQH/\nBAIwADCCCTcGA1UdEQSCCS4wggkqgi5ibHVlb3JhbmdlLmF1c3RyYWxpYWNlbnRy\nYWwuY2xvdWRhcHAuYXp1cmUuY29tgi9ibHVlb3JhbmdlLmF1c3RyYWxpYWNlbnRy\nYWwyLmNsb3VkYXBwLmF6dXJlLmNvbYIrYmx1ZW9yYW5nZS5hdXN0cmFsaWFlYXN0\nLmNsb3VkYXBwLmF6dXJlLmNvbYIwYmx1ZW9yYW5nZS5hdXN0cmFsaWFzb3V0aGVh\nc3QuY2xvdWRhcHAuYXp1cmUuY29tgilibHVlb3JhbmdlLmJyYXppbHNvdXRoLmNs\nb3VkYXBwLmF6dXJlLmNvbYIrYmx1ZW9yYW5nZS5jYW5hZGFjZW50cmFsLmNsb3Vk\nYXBwLmF6dXJlLmNvbYIoYmx1ZW9yYW5nZS5jYW5hZGFlYXN0LmNsb3VkYXBwLmF6\ndXJlLmNvbYIqYmx1ZW9yYW5nZS5jZW50cmFsaW5kaWEuY2xvdWRhcHAuYXp1cmUu\nY29tgidibHVlb3JhbmdlLmNlbnRyYWx1cy5jbG91ZGFwcC5henVyZS5jb22CK2Js\ndWVvcmFuZ2UuY2VudHJhbHVzZXVhcC5jbG91ZGFwcC5henVyZS5jb22CLmJsdWVv\ncmFuZ2UuY2hpbmFlYXN0LmNsb3VkYXBwLmNoaW5hY2xvdWRhcGkuY26CL2JsdWVv\ncmFuZ2UuY2hpbmFlYXN0Mi5jbG91ZGFwcC5jaGluYWNsb3VkYXBpLmNugi9ibHVl\nb3JhbmdlLmNoaW5hbm9ydGguY2xvdWRhcHAuY2hpbmFjbG91ZGFwaS5jboIwYmx1\nZW9yYW5nZS5jaGluYW5vcnRoMi5jbG91ZGFwcC5jaGluYWNsb3VkYXBpLmNugiZi\nbHVlb3JhbmdlLmVhc3Rhc2lhLmNsb3VkYXBwLmF6dXJlLmNvbYIkYmx1ZW9yYW5n\nZS5lYXN0dXMuY2xvdWRhcHAuYXp1cmUuY29tgiVibHVlb3JhbmdlLmVhc3R1czIu\nY2xvdWRhcHAuYXp1cmUuY29tgilibHVlb3JhbmdlLmVhc3R1czJldWFwLmNsb3Vk\nYXBwLmF6dXJlLmNvbYIrYmx1ZW9yYW5nZS5mcmFuY2VjZW50cmFsLmNsb3VkYXBw\nLmF6dXJlLmNvbYIpYmx1ZW9yYW5nZS5mcmFuY2Vzb3V0aC5jbG91ZGFwcC5henVy\nZS5jb22CJ2JsdWVvcmFuZ2UuamFwYW5lYXN0LmNsb3VkYXBwLmF6dXJlLmNvbYIn\nYmx1ZW9yYW5nZS5qYXBhbndlc3QuY2xvdWRhcHAuYXp1cmUuY29tgipibHVlb3Jh\nbmdlLmtvcmVhY2VudHJhbC5jbG91ZGFwcC5henVyZS5jb22CKGJsdWVvcmFuZ2Uu\na29yZWFzb3V0aC5jbG91ZGFwcC5henVyZS5jb22CLGJsdWVvcmFuZ2Uubm9ydGhj\nZW50cmFsdXMuY2xvdWRhcHAuYXp1cmUuY29tgilibHVlb3JhbmdlLm5vcnRoZXVy\nb3BlLmNsb3VkYXBwLmF6dXJlLmNvbYIsYmx1ZW9yYW5nZS5zb3V0aGNlbnRyYWx1\ncy5jbG91ZGFwcC5henVyZS5jb22CK2JsdWVvcmFuZ2Uuc291dGhlYXN0YXNpYS5j\nbG91ZGFwcC5henVyZS5jb22CKGJsdWVvcmFuZ2Uuc291dGhpbmRpYS5jbG91ZGFw\ncC5henVyZS5jb22CJWJsdWVvcmFuZ2UudWtzb3V0aC5jbG91ZGFwcC5henVyZS5j\nb22CJGJsdWVvcmFuZ2UudWt3ZXN0LmNsb3VkYXBwLmF6dXJlLmNvbYIrYmx1ZW9y\nYW5nZS53ZXN0Y2VudHJhbHVzLmNsb3VkYXBwLmF6dXJlLmNvbYIoYmx1ZW9yYW5n\nZS53ZXN0ZXVyb3BlLmNsb3VkYXBwLmF6dXJlLmNvbYInYmx1ZW9yYW5nZS53ZXN0\naW5kaWEuY2xvdWRhcHAuYXp1cmUuY29tgiRibHVlb3JhbmdlLndlc3R1cy5jbG91\nZGFwcC5henVyZS5jb22CJWJsdWVvcmFuZ2Uud2VzdHVzMi5jbG91ZGFwcC5henVy\nZS5jb22CLmJsdWVvcmFuZ2UuY2hpbmFlYXN0LmNsb3VkYXBwLmNoaW5hY2xvdWRh\ncGkuY26CL2JsdWVvcmFuZ2UuY2hpbmFub3J0aC5jbG91ZGFwcC5jaGluYWNsb3Vk\nYXBpLmNugjBibHVlb3JhbmdlLmNoaW5hbm9ydGgyLmNsb3VkYXBwLmNoaW5hY2xv\ndWRhcGkuY26CL2JsdWVvcmFuZ2UuY2hpbmFlYXN0Mi5jbG91ZGFwcC5jaGluYWNs\nb3VkYXBpLmNugjRibHVlb3JhbmdlLmdlcm1hbnljZW50cmFsLmNsb3VkYXBwLm1p\nY3Jvc29mdGF6dXJlLmRlgjZibHVlb3JhbmdlLmdlcm1hbnlub3J0aGVhc3QuY2xv\ndWRhcHAubWljcm9zb2Z0YXp1cmUuZGWCM2JsdWVvcmFuZ2UudXNnb3Z2aXJnaW5p\nYS5jbG91ZGFwcC51c2dvdmNsb3VkYXBpLm5ldIIvYmx1ZW9yYW5nZS51c2dvdmlv\nd2EuY2xvdWRhcHAudXNnb3ZjbG91ZGFwaS5uZXSCMmJsdWVvcmFuZ2UudXNnb3Zh\ncml6b25hLmNsb3VkYXBwLnVzZ292Y2xvdWRhcGkubmV0gjBibHVlb3JhbmdlLnVz\nZ292dGV4YXMuY2xvdWRhcHAudXNnb3ZjbG91ZGFwaS5uZXSCK2JsdWVvcmFuZ2Uu\nZnJhbmNlY2VudHJhbC5jbG91ZGFwcC5henVyZS5jb22CCWxvY2FsaG9zdIIKa3Vi\nZXJuZXRlc4ISa3ViZXJuZXRlcy5kZWZhdWx0ghZrdWJlcm5ldGVzLmRlZmF1bHQu\nc3ZjgiRrdWJlcm5ldGVzLmRlZmF1bHQuc3ZjLmNsdXN0ZXIubG9jYWyCFmt1YmVy\nbmV0ZXMua3ViZS1zeXN0ZW2CGmt1YmVybmV0ZXMua3ViZS1zeXN0ZW0uc3Zjgihr\ndWJlcm5ldGVzLmt1YmUtc3lzdGVtLnN2Yy5jbHVzdGVyLmxvY2FshwQK//8FhwR/\nAAABhwQK//8PhwQKAAABMA0GCSqGSIb3DQEBCwUAA4ICAQCU6ZdOixoG5CZGf/y1\nJ6/dqPxHeVNJAztTSyMThhZcHf8vHn0t+nQi8Js+d9VwgCUyQhIlu1rwOLJGO/3U\ngTzokX59TU5LU+R83wEr9+rsRLFd26sQ7JsrK1uOnXPWUn4J6z2nKLaJIyJU0WYO\n3h4tUwgvBEEA4VzpdAoBbBDTRaoNx1tOGtkgBYCt75PeDe1cXVJJ7kiIl6VxqSgq\nxPvC1ip8ehjdODqJJnC+TRsuZXXTcsja/Y08yg20RHkP3ffiJyorHKpaXCcsgLfr\nXZ/LzFyvNmxq2WnZVRd7Hgaw57cpLpwlruGjRGnVrMk4UP3SO4Fpx6CsKAo//Rtv\nvzHjHWm20XgVADTaAS1A+/l+OT0zGRVpJwm3KwqNPLAWsRJ8d66j8FXaVNGfoYW/\new5dQu39A0bh77sfOgURDHyRIFYhE1czbEBqUmTQW2I0W7daOO5YzQYJoof/fCHW\nRTbnCbRfn2ik127zp34MF+iqdoskRnwLbxTsQ39wYb9q/RviYyCuny/MCkiuigsN\nZGS5NyDon+X0z6ffXJr3ROfQY5pFFw2eYIjWnQjFft7HwLZ6kEsjswqr1mdBP8g6\niwKjKlTiABZc0TzcCOzVlN8fiMAmZ50X93oXjN0nmH9xU58GVkuclyQQqDQYjoCn\nJMZZDn3ped6qrM4kWQThf23QfA==\n-----END CERTIFICATE-----\n","apiServerPrivateKey":"-----BEGIN RSA PRIVATE KEY-----\n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\n-----END RSA PRIVATE KEY-----\n","clientCertificate":"-----BEGIN CERTIFICATE-----\nMIIE9zCCAt+gAwIBAgIRAIqai24txlj3U0yJ69vkqSAwDQYJKoZIhvcNAQELBQAw\nDTELMAkGA1UEAxMCY2EwHhcNMTkwMTEyMDQ0MjA4WhcNMjEwMTExMDQ0MjA4WjAq\nMRcwFQYDVQQKEw5zeXN0ZW06bWFzdGVyczEPMA0GA1UEAxMGY2xpZW50MIICIjAN\nBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAm4Ki+PYfvuWui+o94+kSeCca8MY/\n7LIsyplZeoDo9UTCVICbaxbyXbsummOIQulxzYrJMKFf1mMKbSLSWdkDDUpvEgey\nBdQ/ywtN4GARG7FiRdX0TrNlp/YkgX9BqQNb+7S4fYIV66/4oQXdxGPO0O/Mvt3A\nphdzh5eG6ho7h+vszkVo0WF9xUS/8FU9vg0yB4kxSicjzVs3MyZaM6iryvKa/ORM\nnGBOO3j/Q1tEZcb6E8rOEzRl59jIVHd0lVmWZJxGHUHGpkvnYy2UmUAxK9ahoaG2\n6gZM38W4BGXWEIj1y6vPMrRXkYO5Rz0O06j76IorBKP1B3Y9mhiqRd/0ZNKvmFtl\nXv2k5uxm/pBgDceMBH810wwTXRVlD6CXyQyvjxy/obTxJy+FjBukEXz3KS2UgucP\nHVR6TDKXSo6nIv78SQgQ+0eX+0mFp6E1ULnGHxsYPddVirQb7l2B0MfDijS8mAM0\naSZZ1Qxln5pW+VZgNyHSUIE8oN+AjO7kEp9JrcUN0Vmw+65hdkfVsAqSwdxh0rRO\nFxo30yjEXXFYdjFjJe1gSV5A83MIDt+TMJio8Yt7rqPXgUOOEoAiwRFYIvK0OeMA\n0Wxu6RnFkJmb3UgZbMEHQooUM85uoH0T1mfVXfoGY4FsmblO3dKbYNXNCoCPweHa\n3zmrQtrMUh6zzUECAwEAAaM1MDMwDgYDVR0PAQH/BAQDAgWgMBMGA1UdJQQMMAoG\nCCsGAQUFBwMCMAwGA1UdEwEB/wQCMAAwDQYJKoZIhvcNAQELBQADggIBACAtSjyc\nSpymXUYI4slHWbQrSsdSRY0pqjBOZJILmG0IUxAoP1c+fKpWiOlhlclHkucjAAtw\nnMxoPBnUfAG43PONPX7P2P4T9fyEeT+XRyfV+1gTbbJ0ZmkuAb/GDsTP6UBChpou\nV3UYlZViknC0rhWK/EPo5G4I96JtWtBXzq4/87hTbJaxPV2D7AKKoi5MkkNae56T\nbelkDbQ9+hrdVPqJojfre1eVqTPRDKNSh8L2ZxE4rf6fWvNU9tGebHx8smMcsiwd\nmXpEOTsvdklAFgmWRjBPR+c2GPXrsCbsDzUnWTV4gLx3cbDAtjNXO5aOMYDempsH\ns0z5V8YchNXGVx0KN/6aPGtmcEVr5TKZ3UlL+r1NJn77faqAoNyxnRovUTdlgP32\nVjCFsap8eqXojhDIO18t4n09arWwz1adDTiHvCVsrXLRKt57TCFb3retE5yKYg7D\nnPIloygaOtydwWhTEcIRZBeEruih+MwmkBplQrogIXsQe1R6N5mcKbbfbV+0RjjL\nHXC728me34zHdKF1A3StgW/mjtP/GLHW/fFFJlstiLA877an+nnGnFgSSrVZXTuB\nb5JcPu7cHgfXEkiWh8gthgr5lrxQqQmx8mQeAmtGILGk6sP1an8RCLdQy/OaG4Ri\ndxmzG0Gqb90EhWS/YFuS8y9DV6L4xRstW+UD\n-----END CERTIFICATE-----\n","clientPrivateKey":"-----BEGIN RSA PRIVATE KEY-----\n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\n-----END RSA PRIVATE KEY-----\n","kubeConfigCertificate":"-----BEGIN CERTIFICATE-----\nMIIE9jCCAt6gAwIBAgIQc+LWmkUIaOEl76Qxxw+JFTANBgkqhkiG9w0BAQsFADAN\nMQswCQYDVQQDEwJjYTAeFw0xOTAxMTIwNDQyMDhaFw0yMTAxMTEwNDQyMDhaMCox\nFzAVBgNVBAoTDnN5c3RlbTptYXN0ZXJzMQ8wDQYDVQQDEwZjbGllbnQwggIiMA0G\nCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQC5EUQc3PV091G2NycTe8Mw04KiG/PS\n1HhtZlkZ/RfNr4tt4Nv/J1Axv9VXcovY5L7dzeGBVes+1WPnWnhpj7LvpbVGgngQ\nkVNMXbkdfxwEZSFExXtfmIhK1CBfjfGNw/qIyasIc7F1DB0qBbmUic+UjYeuF6x8\n7dQi8Dag5EYuHhQUFjUJB7DOaf0Xgk6eJcbEyx2TZhVpNvSQDYl+vBYNCqzc++BR\nTVVrIh/+tTKFVuZrC8Quxk30ob/Wlzh5MCta533p45iqQpV30ESeXJOQG8hd+2eW\nklhJxZtOkTF+pe2BBRi5mOx2VWYc2LEQW1Ps4xPRa/d/As543hNYP2urimTDcgF2\nYZ5PFqDmADzgcqf/sA/LpoR2Pllfe909npUEr1GbaoLHEGRlxGoLt6HhGoF27onL\n9/qsYsM8gkUqLE6V7oZVQDCExVXxQ4VAs+7As9Bgj255oBQaGd2FaM6bdG+d6fhV\n/UkcaHdW/DmeCYn6YE9BQvkH4DNLJ8fybJ8juYsRGJqmIJdWtbiHYOUEQVKiIs/z\nBP8W7OXPP+tfHnC1Kbxy05XdFbklFLSc/vnA0pYv2av8YrivyryNw/rzFesLXBai\njA6TxnQ4rb3+I9m+9TGBXQgvpTVbSICL7m7boWbNGTmplnDX5pRijOzRwdMTAOwS\nVDv5f1Bz6uXcNwIDAQABozUwMzAOBgNVHQ8BAf8EBAMCBaAwEwYDVR0lBAwwCgYI\nKwYBBQUHAwIwDAYDVR0TAQH/BAIwADANBgkqhkiG9w0BAQsFAAOCAgEAZXmWiQw9\n7ebyXhCFJ/98BqAQb3R+j5KIKgTDpGfjZca/bJQY9YB/rCVKD0Uw+gLMWGfNdIIo\npRQA5JZUCdCzWuyxUsrjknnGr4Gqeto77+IL77E50IHFkqvo7Jr3YEI/BTjyTMMw\n1HyYFRsPFw8vKvwgEhAR7HD3IR+jbCHBxHNGCmpABS8qH+UXiFu+FsL34SpLnDNz\nb525t2SjFt6XeELmy830jbv1N88DsJ4kYwZSCRG+cKDioWja7J+YMYmZI0rQCpXP\naBgNG+HCUhVu2+vXWb4EuJ2ECAuyA5RISQxaGWykqzEGcU8Lh9xUuvqk6X4QC1ns\n3zgu0pm4sZUqvW1VQUqTwK9o/cZTgSCDVAJSNvpuyOyKBURq9qjeZyszNK+5Oeag\n2TPtAQV/JSM0ewfO+MkeJvzKC8pJ4RRB+b130MbG11WcOy0oUUrdFXbCfdSgHqjr\nU6Qcd/uQOa0rWmlMahdBRRZ4eyDk1HE8XTVP8IdMxbm+lVlwAprWc/jcB7Mo0UpB\nJuPspDQlUnKzWtKj1K4QOcFbycEtuCcuHceTQIvrZNLQI72805LukWSv9UWvw5Hp\nXPGdvqFenssfhoskwu6dw5TGz+qUC8PrOA4bo0ct5sNwr666EVsscwymzj3ZT8Sw\nk72IiT4AKX3uGOcY5WnaDb0AEVmla99SwiM=\n-----END CERTIFICATE-----\n","kubeConfigPrivateKey":"-----BEGIN RSA PRIVATE KEY-----\n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\n-----END RSA PRIVATE KEY-----\n","etcdServerCertificate":"-----BEGIN CERTIFICATE-----\nMIIFDzCCAvegAwIBAgIRAOpZPgJ7zTw5pY7WLJyFtbMwDQYJKoZIhvcNAQELBQAw\nDTELMAkGA1UEAxMCY2EwHhcNMTkwMTEyMDQ0MjA4WhcNMjEwMTExMDQ0MjA4WjAV\nMRMwEQYDVQQDEwpldGNkc2VydmVyMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIIC\nCgKCAgEAtrAM1TMgj511HTirGzsOvG97emUYVuC2mJcP3D6a9Q3usfFSWksGtdhB\ne2qdzNFVO5FxIRLa8IFP3erMDPaayiUHkS6jODTD2V0Kl4Yp7hp9k2y7NeN3SKMJ\nMImOHI6M+pYmXYAUJdpA/8EJAdE9AXQpqT7DNbuURsBlTsOdJ1DGY13nHpkn7vAZ\nuiG5g3Yn6kekWSx2L2LLNwJRD6YRS0xTEKfBEZHPjHmmQL2aa8gPr/NhjGrtwF5T\nQ5A6yMAewS0gci0wI25GXXBbPjSk4coj+8qCGogVGowEe+izmLtx00lD3jZkMwfI\nroQxg+QDxAJK13/3jYkd3nLUN4Zs3hXGKNfZC2NyCEqZkxUlpQ4+Mp1XUwjFVyG1\nB6wvUfS6Eg2lgZaOlrT0QYyNpH323G4qudXBGda13tFTpD+pz4rGpXoAzdeCyRf5\n+QcbuHeOaOdloI/0kHs127dkTkSYA/yt42NlkTMQmvUFXX1l7gyrrrYf6gj69N7l\nRa86YWeTamOp69pJweCvTdGj8aORbBQ2deodMSn5PWtGXyRmwlbaeggD1FoDOVzy\npGkhrZ+y0CYDd74XloJkEHR3n0oDBPVbdZYmOAljmiQbvB7NAoH51nZmaCUTeQO2\nVz+Orgy8MO8NKTaVfGXFOCC5qMEot89h/3sr2Au/68qjs5P8HaECAwEAAaNiMGAw\nDgYDVR0PAQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAM\nBgNVHRMBAf8EAjAAMCEGA1UdEQQaMBiHBAr//wWHBH8AAAGHBAr//w+HBAoAAAEw\nDQYJKoZIhvcNAQELBQADggIBAHohauuLwgvh/V1moffE4SubrLGtjsZxz1wg9Vsh\noPA0QVRoAMxjhFOUy/Ljn32QDS4uJHK2/plACcsQ0S1qiR7zhu3sGtsBFwzkbNjK\nNCPh7as4OUDxNHQN4RcVCjJJUv1kme9SDcXwC1Uorc2jnZqYjzfKxmmYC5pOvFmv\ny/sB48f7lfhVo+4s+KAEhJGqx5W8OIms7v8Vq4lGFRwKTYihjTtEZOxJybwQjIGv\nkVfhxpWONPc4LJQtzsroWMiz+fTJUnqnjbr0sb5zufjbSZI8KZCTsqDnYke61DO3\nVyR7Y7HUs9Ts0YKsw7/EqpmIZ50QeZVmCLL6JiFUhCNI1FlqC3FYDsH/MJ4E5TPg\nvKgS6Jgz83I+d7o7dT5btHeZMErSrh0xnuBs9qyDXbNX551gZCexz7SGqmfxZx7m\nuZiGeR4ik2q7p4Rsmrs64oj8bcAHYS9PQH8NVCvcy62nSWH0HKiC3wwCJmHyMRff\n6SrHwBBKVXQTpJCzTJrdZH7Uw9tIWcxu1ruxnn4SlyXYEdn0VL6he5ZRp+NIo7MZ\nuirMYzCI/3QXkTeiO92s4cFir5pfYGtkTPDvRAWAxVkqQ8tdq0jl30Co01I7bBT4\nRhsHqXrD+IyZPezUp15IltaD2U0GOTXQDG5blm9hQUJMXIODFxfrivucHG19Rt1P\nlOh6\n-----END CERTIFICATE-----\n","etcdServerPrivateKey":"-----BEGIN RSA PRIVATE KEY-----\nMIIJKQIBAAKCAgEAtrAM1TMgj511HTirGzsOvG97emUYVuC2mJcP3D6a9Q3usfFS\nWksGtdhBe2qdzNFVO5FxIRLa8IFP3erMDPaayiUHkS6jODTD2V0Kl4Yp7hp9k2y7\nNeN3SKMJMImOHI6M+pYmXYAUJdpA/8EJAdE9AXQpqT7DNbuURsBlTsOdJ1DGY13n\nHpkn7vAZuiG5g3Yn6kekWSx2L2LLNwJRD6YRS0xTEKfBEZHPjHmmQL2aa8gPr/Nh\njGrtwF5TQ5A6yMAewS0gci0wI25GXXBbPjSk4coj+8qCGogVGowEe+izmLtx00lD\n3jZkMwfIroQxg+QDxAJK13/3jYkd3nLUN4Zs3hXGKNfZC2NyCEqZkxUlpQ4+Mp1X\nUwjFVyG1B6wvUfS6Eg2lgZaOlrT0QYyNpH323G4qudXBGda13tFTpD+pz4rGpXoA\nzdeCyRf5+QcbuHeOaOdloI/0kHs127dkTkSYA/yt42NlkTMQmvUFXX1l7gyrrrYf\n6gj69N7lRa86YWeTamOp69pJweCvTdGj8aORbBQ2deodMSn5PWtGXyRmwlbaeggD\n1FoDOVzypGkhrZ+y0CYDd74XloJkEHR3n0oDBPVbdZYmOAljmiQbvB7NAoH51nZm\naCUTeQO2Vz+Orgy8MO8NKTaVfGXFOCC5qMEot89h/3sr2Au/68qjs5P8HaECAwEA\nAQKCAgA3h+IOuGDQZstfm4cfWt9K1hRRiwNP/TRjw59Vkk4l6RtCSZl/ysh6ZAbb\njffzdzoSRk5+AC4+5v+w9BscYaWBhqn7LpL8lcVmgAqlLmn2b2T2eBmb8s2ibbRw\nZY+mDIq77QIyb6kwLFyPoUysmb6Sf43eXS6XWbJjoz3oKDvP5JS5RaToPyPNVHxt\ngKzUUgkmBKrnVEWEd+JPkUu0lwUwvz2MlYFxZsIQ8DVh/oA+/OwPzso7FZG5ZLKo\nmeHUfdmbXK09J5E3Y+DNrEZ/7R6lZ31ynwbXK8BGdoMyavSUm32o/N89X4krndUZ\nfyNR9PBUF2JKiSJlimVi5cKuMhVLFRPYf/U9osAzbV9fB82ENSAY1mH9KkuqQFBS\nmnHrj6VwH6XO2jtxJc5roQ+PjJkSI6IT+rtI6aDeOFSn4C+qwsLb8WBBfQkrrUNs\ngmp4cpXs/zhIwrxLNL6JU36CIvVRZhwO0rW6JVTcdrAq+Oedda9vK1OOrRhkedYx\nJ85V1YmQU1KyQbJuuo87vraZjQ6SroVJoCPC7v1vZiqbHD/85OXYqJRiw+TdHBIx\nKdkdII5I3fUOoaAG6SRHQihh4ahvyvEIkPkhmq47en1Cl/J0LrmT58ZCgYHnDZPg\n7qSF/K10K6S1mkx3zaM2yheepPiwkejRcSU5SfDL/zinzlnoAQKCAQEA1MhSSMze\nsMT7ZJvDGQ+c4+1FP3vczKxJNPzmvjTmksYriNVlO3IUWpZZroNvFZdB4dCgHttZ\nYDsaqGCSDbgHqzLtAfN2obbLJOuSrHn8fGL4VIahgl2kQt7CgNIGVWvcgmJXyYx/\n9c4X3+6G53AVCg0z9f6oMU59wS+5eXfpHS24BjfYwnZ75jv66wfs98lE4BM3JjPh\nFqBr0zku6SBnBwgnwHX840FEullkO9zvHoV+ZalGCxSUg5sAXWs06uPWDeNyBxSP\nSAgy8dSaQRWygcXj/tiTArHAV6ktJs9saI3Ao2R4h+O8pxe+hKY7ns3zE6NGpUmo\nrp+zAaKpEIme4QKCAQEA28rxEoo1DRIA8JiOXdJpO1BS3l364Gw3rVT5JYnFBbJ3\npC65eAJJ7A+6nRb2KsWT7xTg3XjYn5RAUkq/yQLSUMV7hxJ0/pTdrPQ9FFaflInp\nOkY6Ca6lX5OwXQkuVF3Kx5lyjs1S13gGdiSUfk6xc6SSl3yJPupzNNuPcsthTV+e\nvk0vebAVHq1hMQlf9FKi5BqBOVLkyas2ekAjW1/w0atCw17eIT2XvOMtmBflM9X6\n6mnxDRmXsQx5CV33R5aFvqgynkZ7M6df8HWp8KuakksfLuGAELTo9YUlO31cyttF\nPobqStE7aAZVIqGFQqTsW4AsP6ZLffITVnuIvxQWwQKCAQBmg1RRSpKHK3/KStjF\nvCXkEa3pFe2clex3INiyDp5/XAYhWF37M8zmj8UJNL1R85uEBZT/CMKYgCEpuczj\n2yOmsVRKOlePNZfNY8k4W1TvQGdPocUyH5dBuDyQ/56ZjOuhzWtp6MzFSdNqyWop\nGnCodQ/xlgzGJIClkC6VfNXMfvhH1qQRnC+5DnCmxxy9jDQomFlbiGcLFaKdEtGs\n8zVvx9gO+0ko24sXFHYb+Sci10G8DL94GyQp+4VDfKD0AWxhKJEJJDibhYe1xq3C\naYekygg8AW9iPuzhztm39vrNrG2AyqjfHzMGdYQOOGBE9AgGEAW7IC+qYbpGMW2u\nT/9BAoIBAQCuNzdMtcYFPS8HhjTag0ysGigFYELaHo0nVAJEUeacdHDG785NJKtJ\ndNI7cBbhokh3Knpusjoqi01MrTlFcHoaUd5vGx8nKAJp9BJyC5NkHsiCS2X3hLpK\nfvs9U2iosMtp2ORn8rHmXpnojWgykdewTVNwxeLXvuhgjmQu3qyast9WAkJOV1b2\nQQGX59FqDRJWcffZ4y27+H0u+6XK6Mout2wtBgZWHMcojn32X4JqywKfSigqdQ26\n5aMO3K+Dy/jpPdMZh85iDBpMtEdmn+7eZLMr6swi8fLxoX74n46ssI87V59gfGZ1\nNwfGcnb5c1Zx1K++J+cQxo+zbinfI2JBAoIBAQCK4cOvTdagQIz0NT532O2qhYHs\npKCZbdFxBJM38kPnFn/N01qO8ns5qVgkGYVJqa4+qGpjgapG+Clh7CidCf/zJIzh\nzExvq5jNufuH1K1bf7tb4c6huNVaJSrhE61KzhpKg466SPbm2R0H0mosqqdjDw2B\nUSE/kl0NmSFP5QBpWTseUVv+tBX911AcqDFd04dmuQcxN5PfQ2ifY7rHgkU46HnG\nL+RMyH5TPx3jnMrL3ADSyuclcvc1t79Q2q5k7VswWUsrb/lgesV8krRn3zLvKGZ3\nJtqWZgfoRh2uTuv00zW4lhtZPnLtSrRTT9yLAiK9ntlh6Z4/Iec76a8DKBPO\n-----END RSA PRIVATE KEY-----\n","etcdClientCertificate":"-----BEGIN CERTIFICATE-----\nMIIFDjCCAvagAwIBAgIQdj5P9Jo2HWSN43M4FZoTAjANBgkqhkiG9w0BAQsFADAN\nMQswCQYDVQQDEwJjYTAeFw0xOTAxMTIwNDQyMDhaFw0yMTAxMTEwNDQyMDhaMBUx\nEzARBgNVBAMTCmV0Y2RjbGllbnQwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIK\nAoICAQDPxjN6Gr46Iz6bptko52BEuG8w6oNt2JKh8lZhn7R0WWWM9DBnZsTvJooj\nhB2JSzepQRUSg6Ka29aEeolY4C9Ncq8Wo6zjkjnC9iR9wcMMuGgvGiArWv4LWeE9\nITPPv5rofzZZ07wh8Yr0quMrovqMFJtTKt+o5s+RzA0Y9cAZa4Lt9o0DeZxfsQ6f\nGt8PReF7tcOVnpL0CpHBjKuW7U+WBkVG+WX19BSXp47t6HTlziJY7aHBmcg+QfKH\nJPwDphcbAyr8ITrlemOIxTD2/K+Qby45FPf/jwgyYkJVVMpVIf6updH5u+Zvhw/5\nXY8TjG0W92vpm4DKPEx0S3cPzlS6wKVWKMvxAULzjuxjswLYZDeptB78MDc4yjc7\nNXZcGLPQNlw5usULEkzvCrWdWfv5zsq5WLzSSxhc0HU+VmlmDTe7sA8GBVmAHEbN\nj9KpndKTuGtEX0ic3vt/q4SXbd2rbHYce3vlc6a6B/mlKvUHgZmM7gZLTSolk4bk\nMBpG1EvlaexewH0p87uqvIpXPY17dQGu2b8rPF8iZOX9WDh4fROaoP7Itvu2//NX\nqKxqiuh11osynC6n845N1/RnBnHK3DXC3D9eeCkbOl0gYTaaNPfLB18YXcEqIkxy\n9EcQhchtveILt+IGyARacqs6m4o67I98g6PT18Ymq7jejf2fjwIDAQABo2IwYDAO\nBgNVHQ8BAf8EBAMCBaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMAwG\nA1UdEwEB/wQCMAAwIQYDVR0RBBowGIcECv//BYcEfwAAAYcECv//D4cECgAAATAN\nBgkqhkiG9w0BAQsFAAOCAgEAavSP7vBn1LpMZ/9LV7jFcXF+cZEbgYLDjmNDODeG\n6kEjfGB+ANdVqaG5vdD3X4hFAltmBrb9ZTMCL3fyazGNoFWq78w4BdeVyzZLELhg\nVxzfTofENY0OckwtbwnP08mbpzyKS5BpFZ+jUZQAW6Ok+1M+H6OcfhJh/OE8X1PK\nIIALSdb6tYaND437ogbDInyvkppiaPQMEw+l4v5c8kWojSPrtLirD5ME4iokEzra\naAjNZe6QWpoUoi5RktZWtY9Le3G5FpnWZIZkNHVYQ87JKiCbZL/H4DuWoqeXE5qG\nD1CrLogH34NY6XfCYdA3JGf8YMlG3fqspbY/XmxcWsJimxavl6ZwdRugO+Simncz\ngGRxLfJaNZYnU1s4If9+QJRHWTIniwv19iKaz+k7L/pfGblE2NzqhCSDmdDYukbW\n1eEbQvlwwP9C+W17qeYcUfzAsgVF0OjfHs8qODJKyeUQbfQKHq6A3Usvkbp6kmSI\n9Yq4rPQer2mzfZisR5SV+AuWS8oyYNtGMFVm6yUCJaCtnSts2R0hEHsNVEhO76+0\nUpEj5a16DAPRJ793d1fgPzAfm4FPnPikw2gnM+mucswa87zDQjFqXTC+JoOWsgrL\nwa9ZZTpkm/VzMKj7Y5A8h/JaXLAYMLI6h5gUGvdo4ErA20BI3yQCSqKub4/Zb4Zx\n1Ww=\n-----END CERTIFICATE-----\n","etcdClientPrivateKey":"-----BEGIN RSA PRIVATE KEY-----\n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\n-----END RSA PRIVATE KEY-----\n","etcdPeerCertificates":["-----BEGIN CERTIFICATE-----\nMIIFDDCCAvSgAwIBAgIQRlEOER3iv5t2FjcySjSuXjANBgkqhkiG9w0BAQsFADAN\nMQswCQYDVQQDEwJjYTAeFw0xOTAxMTIwNDQyMDhaFw0yMTAxMTEwNDQyMDhaMBMx\nETAPBgNVBAMTCGV0Y2RwZWVyMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKC\nAgEAtEpCDv6nHyc6llYsjeWLb6Ouc5+prMHnqsXJwn6PKOA/oCBT4RCHxjaSkXXw\nSIhD61sWd5gyPAJNo73Od+5afLWCQJmpWiH4nWrKFyIeWmRztrsIfmh+sJDsk9HK\nl2eODTkEOlUy2xbNWoglECabsLik75hHlw5uYOvxwcn5ioLyPgyNb7TtC1iGSu5q\nVsBNejnG/bFye7X+dEJPXyvru1JLmymX2wCcR4AoM7lfSH67rm2GySzcY0G3YSf2\nBXMXsF4oMFF8UsEyrNekDNuwNIG9fSn1DucWkcy7uOXlGBUs6C6jP9VvzKnIVz1w\nscDB1PTO5I/HYqPVAPJ0YajvesXuKbRqa8BCngyRTcxtcnKwRvjH6twHM3z+I85v\nXrJFoy044vp8ScH+ES9wE50EwO4uGXC+I5caXr+r3CZn603aD0p8zoqlDqhI3IEC\nMbEA+fbt9Q6dZoVLtm8HbsQW6orWr/o8mjUDpb2jkJLwi+fIXF+pB61qhAsujApy\nzF/OtEOPLCZzMQJBmyrAoyCp1CGyX25B+Fi+4schv1hB1YUo12wnO6346hvtwoP7\nImpfR/8vMyrvqFJYknNW5b9cwbv6OqKGH7B/aGlYoYShYjmwRX578ZVtVl9uX8Eu\nmdQqidpMz4ojosbjST0ldtKNvayik6q0Tfo2bfp+355bqPcCAwEAAaNiMGAwDgYD\nVR0PAQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNV\nHRMBAf8EAjAAMCEGA1UdEQQaMBiHBAr//wWHBH8AAAGHBAr//w+HBAoAAAEwDQYJ\nKoZIhvcNAQELBQADggIBAK2LZEV3DQUVbi2AP9DFGKErOlDTMU1gupSZ2Jzrh7hc\nlkR3aIdCF0jYcdqFjs65fDt87RqOt0uMQrYApupnKXp8GIRPjp17gPlPbzYywVRf\nhkU3WtHtntiMynNYIpJgXJzeSnlFXhUiBtAEy/FYmcExzBtEx9g2jT382LnLKeuV\ns/k+szdBPF2x5Bg7ewFCwVAsKASPjnTOC7YrJMef50Wq+sfqyyKVo20OPIuo3VFc\nvx7rGSKtaYy1KSL3y2Ctn5HiZTDTvuNDmSXhHJyIrhksrN4V03bpypq/H1jVmPY9\nP8zGwmF+fAYfmZhBNuHOmBhCH3JaAgun3y6iPXRhBhLkZD05mJ7EzKBvzDpxq2ha\nxP7Pz3KYuwl1DBD70QWVZqtjE562OcVRStGtFHVdPqP5KwrO7lEw6oIRiSEmc1Lx\nQUdVs5hrYXZXguL9TU+Zin/YnXCZwjidhObKX9uV272UFsHv8z8Ybiz7HqDKKSd3\nGBHP+Ehy/rr6SBt5fI3LTUk0aQv/gn0Y7egdezfgu8j0uve1WEc3ppl4pyyTradF\nlWebp0sQE8M+ASYyNWtpbTt86ZHirkKCg0fOlJEY/ndGGMaAUgnlDFx1/DKILURb\nq7LVqeL/TIeuKYv9rUN0/TAqkA4G2bL5pZmXTBUlgatlBcQoP659/fjKJMohLbRo\n-----END CERTIFICATE-----\n"],"etcdPeerPrivateKeys":["-----BEGIN RSA PRIVATE KEY-----\n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\n-----END RSA PRIVATE KEY-----\n"]}}}`
}

func TestGenerateUserAssignedIdentityClientIDParameter(t *testing.T) {
	testCases := []struct {
		name                   string
		isUserAssignedIdentity bool
		expected               string
	}{
		{
			name:                   "enabled",
			isUserAssignedIdentity: true,
			expected:               "' USER_ASSIGNED_IDENTITY_ID=',reference(variables('userAssignedIDReference'), variables('apiVersionManagedIdentity')).clientId, ' '",
		},
		{
			name:                   "disabled",
			isUserAssignedIdentity: false,
			expected:               "' USER_ASSIGNED_IDENTITY_ID=',' '",
		},
	}

	for _, c := range testCases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if ret := generateUserAssignedIdentityClientIDParameter(c.isUserAssignedIdentity); ret != c.expected {
				t.Fatalf("generateUserAssignedIdentityClientIDParameter(%t) returned %s, expected %s", c.isUserAssignedIdentity, ret, c.expected)
			}
		})
	}
}
