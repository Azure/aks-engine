// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

//
//import (
//	"fmt"
//
//	"github.com/Azure/aks-engine/pkg/api"
//	"github.com/Azure/go-autorest/autorest/to"
//)
//
//type ARMParamMetadata struct {
//	Description *string
//}
//
//type ARMParam struct {
//	DefaultValue  *string
//	Metadata      *ARMParamMetadata
//	Type          *string
//	AllowedValues *[]string
//	MaxLength     *int
//	MinLength     *int
//}
//
//func toARMParameterObj(armParam ARMParam) interface{} {
//	paramMap := map[string]interface{}{}
//	if armParam.DefaultValue != nil {
//		paramMap["defaultValue"] = armParam.DefaultValue
//	}
//
//	if armParam.Metadata != nil {
//		metadataMap := map[string]string{}
//		if armParam.Metadata.Description != nil {
//			metadataMap["description"] = *armParam.Metadata.Description
//		}
//		paramMap["metadata"] = metadataMap
//	}
//
//	if armParam.Type != nil {
//		paramMap["type"] = armParam.Type
//	}
//
//	if armParam.AllowedValues != nil {
//		allowedValues := &armParam.AllowedValues
//		paramMap["allowedValues"] = allowedValues
//	}
//
//	if armParam.MinLength != nil {
//		paramMap["minLength"] = armParam.MinLength
//	}
//
//	if armParam.MinLength != nil {
//		paramMap["maxLength"] = armParam.MinLength
//	}
//
//	return paramMap
//}
//
//func getK8sParams(cs *api.ContainerService) map[string]interface{} {
//
//	paramsMap := map[string]interface{}{}
//
//	p := cs.Properties
//
//	isHostedMaster := p.IsHostedMasterProfile()
//
//	if p.HasAadProfile() {
//		aadTenantID := ARMParam{
//			DefaultValue: to.StringPtr(""),
//			Metadata: &ARMParamMetadata{
//				Description: to.StringPtr("The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription."),
//			},
//			Type: to.StringPtr("string"),
//		}
//
//		paramsMap["aadTenantId"] = toARMParameterObj(aadTenantID)
//
//		aadAdminGroupID := ARMParam{
//			DefaultValue: to.StringPtr(""),
//			Metadata: &ARMParamMetadata{
//				Description: to.StringPtr("The AAD default Admin group Object ID used to create a cluster-admin RBAC role."),
//			},
//			Type: to.StringPtr("string"),
//		}
//
//		paramsMap["aadAdminGroupId"] = toARMParameterObj(aadAdminGroupID)
//	}
//
//	if isHostedMaster {
//		kubernetesEndpoint := ARMParam{
//			Metadata: &ARMParamMetadata{
//				Description: to.StringPtr("The Kubernetes API endpoint https://<kubernetesEndpoint>:443"),
//			},
//			Type: to.StringPtr("string"),
//		}
//		paramsMap["kubernetesEndpoint"] = toARMParameterObj(kubernetesEndpoint)
//	} else {
//		etcdServerCertificate := ARMParam{
//			Metadata: &ARMParamMetadata{
//				Description: to.StringPtr("The base 64 server certificate used on the master"),
//			},
//			Type: to.StringPtr("string"),
//		}
//		paramsMap["etcdServerCertificate"] = toARMParameterObj(etcdServerCertificate)
//
//		etcdServerPrivateKey := ARMParam{
//			Metadata: &ARMParamMetadata{
//				Description: to.StringPtr("The base 64 server private key used on the master."),
//			},
//			Type: to.StringPtr("securestring"),
//		}
//		paramsMap["etcdServerPrivateKey"] = toARMParameterObj(etcdServerPrivateKey)
//
//		etcdClientCertificate := ARMParam{
//			Metadata: &ARMParamMetadata{
//				Description: to.StringPtr("The base 64 server certificate used on the master"),
//			},
//			Type: to.StringPtr("string"),
//		}
//
//		paramsMap["etcdClientCertificate"] = toARMParameterObj(etcdClientCertificate)
//
//		etcdClientPrivateKey := ARMParam{
//			Metadata: &ARMParamMetadata{
//				Description: to.StringPtr("The base 64 server private key used on the master."),
//			},
//			Type: to.StringPtr("securestring"),
//		}
//
//		paramsMap["etcdClientPrivateKey"] = toARMParameterObj(etcdClientPrivateKey)
//
//		masterCount := p.MasterProfile.Count
//
//		if masterCount == 1 || masterCount == 3 || masterCount == 5 {
//			for i := 0; i < masterCount; i++ {
//				paramsMap[fmt.Sprintf("etcdPeerCertificate%d", i)] = toARMParameterObj(ARMParam{
//					Metadata: &ARMParamMetadata{
//						Description: to.StringPtr("The base 64 server certificates used on the master"),
//					},
//					Type: to.StringPtr("string"),
//				})
//				paramsMap[fmt.Sprintf("etcdPeerPrivateKey%d", i)] = toARMParameterObj(ARMParam{
//					Metadata: &ARMParamMetadata{
//						Description: to.StringPtr("The base 64 server private keys used on the master."),
//					},
//					Type: to.StringPtr("securestring"),
//				})
//			}
//		}
//	}
//
//	apiServerCertificate := ARMParam{
//		Metadata: &ARMParamMetadata{
//			Description: to.StringPtr("The base 64 server certificate used on the master"),
//		},
//		Type: to.StringPtr("string"),
//	}
//	paramsMap["apiServerCertificate"] = toARMParameterObj(apiServerCertificate)
//
//	apiServerPrivateKey := ARMParam{
//		Metadata: &ARMParamMetadata{
//			Description: to.StringPtr("The base 64 server private key used on the master."),
//		},
//		Type: to.StringPtr("securestring"),
//	}
//	paramsMap["apiServerPrivateKey"] = toARMParameterObj(apiServerPrivateKey)
//
//	caCertificate := ARMParam{
//		Metadata: &ARMParamMetadata{
//			Description: to.StringPtr("The base 64 certificate authority certificate"),
//		},
//		Type: to.StringPtr("string"),
//	}
//	paramsMap["caCertificate"] = toARMParameterObj(caCertificate)
//
//	caPrivateKey := ARMParam{
//		Metadata: &ARMParamMetadata{
//			Description: to.StringPtr("The base 64 CA private key used on the master."),
//		},
//		Type: to.StringPtr("securestring"),
//	}
//
//	paramsMap["caPrivateKey"] = toARMParameterObj(caPrivateKey)
//
//	clientCertificate := ARMParam{
//		Metadata: &ARMParamMetadata{
//			Description: to.StringPtr("The base 64 client certificate used to communicate with the master"),
//		},
//		Type: to.StringPtr("string"),
//	}
//
//	paramsMap["clientCertificate"] = toARMParameterObj(clientCertificate)
//
//	clientPrivateKey := ARMParam{
//		Metadata: &ARMParamMetadata{
//			Description: to.StringPtr("The base 64 client private key used to communicate with the master"),
//		},
//		Type: to.StringPtr("securestring"),
//	}
//
//	paramsMap["clientPrivateKey"] = toARMParameterObj(clientPrivateKey)
//
//	kubeConfigCertificate := ARMParam{
//		Metadata: &ARMParamMetadata{
//			Description: to.StringPtr("The base 64 certificate used by cli to communicate with the master"),
//		},
//		Type: to.StringPtr("string"),
//	}
//
//	paramsMap["kubeConfigCertificate"] = toARMParameterObj(kubeConfigCertificate)
//
//	kubeConfigPrivateKey := ARMParam{
//		Metadata: &ARMParamMetadata{
//			Description: to.StringPtr("The base 64 private key used by cli to communicate with the master"),
//		},
//		Type: to.StringPtr("securestring"),
//	}
//	paramsMap["kubeConfigPrivateKey"] = toARMParameterObj(kubeConfigPrivateKey)
//
//	generatorCode := ARMParam{
//		Metadata: &ARMParamMetadata{
//			Description: to.StringPtr("The generator code used to identify the generator"),
//		},
//		Type: to.StringPtr("string"),
//	}
//
//	paramsMap["generatorCode"] = toARMParameterObj(generatorCode)
//
//	orchestratorName := ARMParam{
//		Metadata: &ARMParamMetadata{
//			Description: to.StringPtr("The orchestrator name used to identify the orchestrator.  This must be no more than 3 digits in length, otherwise it will exceed Windows Naming"),
//		},
//		Type:      to.StringPtr("string"),
//		MinLength: to.IntPtr(3),
//		MaxLength: to.IntPtr(3),
//	}
//
//	paramsMap["orchestratorName"] = toARMParameterObj(orchestratorName)
//
//	dockerBridgeCidr := ARMParam{
//		Metadata: &ARMParamMetadata{
//			Description: to.StringPtr("Docker bridge network IP address and subnet"),
//		},
//		Type: to.StringPtr("string"),
//	}
//
//	paramsMap["dockerBridgeCidr"] = toARMParameterObj(dockerBridgeCidr)
//
//	kubeClusterCidr := ARMParam{
//		Metadata: &ARMParamMetadata{
//			Description: to.StringPtr("Kubernetes cluster subnet"),
//		},
//		Type: to.StringPtr("string"),
//	}
//
//	paramsMap["kubeClusterCidr"] = toARMParameterObj(kubeClusterCidr)
//
//	kubeDNSServiceIP := ARMParam{
//		Metadata: &ARMParamMetadata{
//			Description: to.StringPtr("Kubernetes DNS IP"),
//		},
//		Type: to.StringPtr("string"),
//	}
//
//	paramsMap["kubeDNSServiceIP"] = toARMParameterObj(kubeDNSServiceIP)
//
//	kubernetesKubeletClusterDomain := ARMParam{
//		Metadata: &ARMParamMetadata{
//			Description: to.StringPtr("--cluster-domain Kubelet config"),
//		},
//		Type: to.StringPtr("string"),
//	}
//
//	paramsMap["kubernetesKubeletClusterDomain"] = toARMParameterObj(kubernetesKubeletClusterDomain)
//
//	kubernetesHyperkubeSpec := ARMParam{
//		Metadata: &ARMParamMetadata{
//			Description: to.StringPtr("The container spec for hyperkube."),
//		},
//		Type: to.StringPtr("string"),
//	}
//
//	paramsMap["kubernetesHyperkubeSpec"] = toARMParameterObj(kubernetesHyperkubeSpec)
//
//	kubernetesCcmImageSpec := ARMParam{
//		DefaultValue: to.StringPtr(""),
//		Metadata: &ARMParamMetadata{
//			Description: to.StringPtr("The container spec for cloud-controller-manager."),
//		},
//		Type: to.StringPtr("string"),
//	}
//
//	paramsMap["kubernetesCcmImageSpec"] = toARMParameterObj(kubernetesCcmImageSpec)
//
//	kubernetesAddonManagerSpec := ARMParam{
//		Metadata: &ARMParamMetadata{
//			Description: to.StringPtr("The container spec for hyperkube."),
//		},
//		Type: to.StringPtr("string"),
//	}
//
//	paramsMap["kubernetesAddonManagerSpec"] = toARMParameterObj(kubernetesAddonManagerSpec)
//
//	enableAggregatedAPIs := ARMParam{
//		Metadata: &ARMParamMetadata{
//			Description: to.StringPtr("Enable aggregated API on master nodes"),
//		},
//		DefaultValue: to.StringPtr("false"),
//		Type:         to.StringPtr("bool"),
//	}
//
//	paramsMap["enableAggregatedAPIs"] = toARMParameterObj(enableAggregatedAPIs)
//
//	return paramsMap
//}
