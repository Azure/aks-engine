// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"github.com/Azure/aks-engine/pkg/api/vlabs"
	"github.com/Azure/aks-engine/pkg/i18n"
	"github.com/leonelquinteros/gotext"

	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestLoadContainerServiceFromFile(t *testing.T) {
	existingContainerService := &ContainerService{Name: "test",
		Properties: &Properties{OrchestratorProfile: &OrchestratorProfile{OrchestratorType: Kubernetes, OrchestratorVersion: "1.7.16"}}}

	locale := gotext.NewLocale(path.Join("..", "..", "translations"), "en_US")
	if err := i18n.Initialize(locale); err != nil {
		t.Error(err)
	}
	apiloader := &Apiloader{
		Translator: &i18n.Translator{
			Locale: locale,
		},
	}

	_, _, err := apiloader.LoadContainerServiceFromFile("../engine/testdata/simple/kubernetes.json", false, false, existingContainerService)
	if err != nil {
		t.Error(err.Error())
	}

	// Test error scenario
	_, _, err = apiloader.LoadContainerServiceFromFile("../this-file-doesnt-exist.json", true, false, nil)
	if err == nil {
		t.Errorf("expected error passing a non-existent filepath string to apiloader.LoadContainerServiceFromFile(), instead got nil")
	}
}

func TestLoadContainerServiceWithEmptyLocationCustomCloud(t *testing.T) {
	jsonWithoutlocationcustomcloud := `{
		"apiVersion": "vlabs",
		"properties": {
			"orchestratorProfile": {
				"orchestratorType": "Kubernetes",
				"kubernetesConfig": {
					"kubernetesImageBase": "msazurestackqa/",
					"useInstanceMetadata": false,
					"networkPolicy": "none"
				}
			},
			"customCloudProfile": {
				"environment": {
					"name": "AzureStackCloud",
					"managementPortalURL": "",
					"publishSettingsURL": "",
					"serviceManagementEndpoint": "https://management.azurestackci15.onmicrosoft.com/36f71706-54df-4305-9847-5b038a4cf189",
					"resourceManagerEndpoint": "https://management.local.azurestack.external/",
					"activeDirectoryEndpoint": "https://login.windows.net/",
					"galleryEndpoint": "https://portal.local.azurestack.external:30015/",
					"keyVaultEndpoint": "https://vault.azurestackci15.onmicrosoft.com/36f71706-54df-4305-9847-5b038a4cf189",
					"graphEndpoint": "https://graph.windows.net/",
					"storageEndpointSuffix": "local.azurestack.external",
					"sqlDatabaseDNSSuffix": "",
					"trafficManagerDNSSuffix": "",
					"keyVaultDNSSuffix": "vault.local.azurestack.external",
					"serviceBusEndpointSuffix": "",
					"serviceManagementVMDNSSuffix": "cloudapp.net",
					"resourceManagerVMDNSSuffix": "cloudapp.azurestack.external",
					"containerRegistryDNSSuffix": ""
				}
			},
			"masterProfile": {
				"dnsPrefix": "k111006",
				"distro": "ubuntu",
				"osDiskSizeGB": 200,
				"count": 3,
				"vmSize": "Standard_D2_v2"
			},
			"agentPoolProfiles": [
				{
					"name": "linuxpool",
					"osDiskSizeGB": 200,
					"count": 3,
					"vmSize": "Standard_D2_v2",
					"distro": "ubuntu",
					"availabilityProfile": "AvailabilitySet",
					"AcceleratedNetworkingEnabled": false
				}
			],
			"linuxProfile": {
				"adminUsername": "azureuser",
				"ssh": {
					"publicKeys": [
						{
							"keyData": "ssh-rsa PblicKey"
						}
					]
				}
			},
			"servicePrincipalProfile": {
				"clientId": "clientId",
				"secret": "secret"
			}
		}
	}`

	tmpFile, err := ioutil.TempFile("", "containerService-nolocation")
	if err != nil {
		t.Error(err)
	}
	fileName := tmpFile.Name()
	defer os.Remove(fileName)

	err = ioutil.WriteFile(fileName, []byte(jsonWithoutlocationcustomcloud), os.ModeAppend)
	if err != nil {
		t.Error(err)
	}

	apiloader := &Apiloader{}
	_, _, err = apiloader.LoadContainerServiceFromFile(fileName, true, false, nil)
	if err == nil {
		t.Errorf("Expected error for missing loation to be thrown")
	}
	expectedMsg := "missing ContainerService Location"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error with message %s but got %s", expectedMsg, err.Error())
	}

	jsonWithoutlocationpubliccloud := `{
		"apiVersion": "vlabs",
		"properties": {
			"orchestratorProfile": {
				"orchestratorType": "Kubernetes",
				"kubernetesConfig": {
					"kubernetesImageBase": "msazurestackqa/",
					"useInstanceMetadata": false,
					"networkPolicy": "none"
				}
			},
			"masterProfile": {
				"dnsPrefix": "k111006",
				"distro": "ubuntu",
				"osDiskSizeGB": 200,
				"count": 3,
				"vmSize": "Standard_D2_v2"
			},
			"agentPoolProfiles": [
				{
					"name": "linuxpool",
					"osDiskSizeGB": 200,
					"count": 3,
					"vmSize": "Standard_D2_v2",
					"distro": "ubuntu",
					"availabilityProfile": "AvailabilitySet",
					"AcceleratedNetworkingEnabled": false
				}
			],
			"linuxProfile": {
				"adminUsername": "azureuser",
				"ssh": {
					"publicKeys": [
						{
							"keyData": "ssh-rsa PblicKey"
						}
					]
				}
			},
			"servicePrincipalProfile": {
				"clientId": "clientId",
				"secret": "secret"
			}
		}
	}`

	tmpFilewithoutlocationpubliccloud, err := ioutil.TempFile("", "containerService-nolocationpubliccloud")
	if err != nil {
		t.Error(err)
	}
	fileNamewithoutlocationpubliccloud := tmpFilewithoutlocationpubliccloud.Name()
	defer os.Remove(fileNamewithoutlocationpubliccloud)

	err = ioutil.WriteFile(fileNamewithoutlocationpubliccloud, []byte(jsonWithoutlocationpubliccloud), os.ModeAppend)
	if err != nil {
		t.Error(err)
	}

	apiloaderwithoutlocationpubliccloud := &Apiloader{}
	_, _, err = apiloaderwithoutlocationpubliccloud.LoadContainerServiceFromFile(fileNamewithoutlocationpubliccloud, true, false, nil)
	if err != nil {
		t.Errorf("Expected no error for missing loation for public cloud to be thrown")
	}
}

func TestDeserializeContainerService(t *testing.T) {
	apiloader := &Apiloader{
		Translator: nil,
	}

	// Test AKS Engine api model
	_, version, err := apiloader.DeserializeContainerService([]byte(exampleAPIModel), false, false, nil)
	if err != nil {
		t.Errorf("unexpected error deserializing the example apimodel: %s", err)
	}
	if version != vlabs.APIVersion {
		t.Errorf("expected apiVersion %s, instead got: %s", vlabs.APIVersion, version)
	}

	// Test error case
	_, _, err = apiloader.DeserializeContainerService([]byte(`{thisisnotson}`), false, false, nil)
	if err == nil {
		t.Errorf("expected error from malformed api model input")
	}
}

func TestLoadDefaultContainerServiceProperties(t *testing.T) {
	m, p := LoadDefaultContainerServiceProperties()

	if m.APIVersion != defaultAPIVersion {
		t.Errorf("Expected LoadDefaultContainerServiceProperties() to return API version %s, instead got %s", defaultAPIVersion, m.APIVersion)
	}

	if p.OrchestratorProfile.OrchestratorType != defaultOrchestrator {
		t.Errorf("Expected LoadDefaultContainerServiceProperties() to return %s OrchestratorProfile.OrchestratorType, instead got %s", Kubernetes, p.OrchestratorProfile.OrchestratorType)
	}

	if p.MasterProfile.Count != defaultMasterCount {
		t.Errorf("Expected LoadDefaultContainerServiceProperties() to return %d MasterProfile.Count, instead got %d", defaultMasterCount, p.MasterProfile.Count)
	}

	if p.MasterProfile.VMSize != defaultVMSize {
		t.Errorf("Expected LoadDefaultContainerServiceProperties() to return %s MasterProfile.VMSize, instead got %s", defaultVMSize, p.MasterProfile.VMSize)
	}

	if p.MasterProfile.OSDiskSizeGB != defaultOSDiskSizeGB {
		t.Errorf("Expected LoadDefaultContainerServiceProperties() to return %d MasterProfile.OSDiskSizeGB, instead got %d", defaultOSDiskSizeGB, p.MasterProfile.OSDiskSizeGB)
	}

	if len(p.AgentPoolProfiles) != 1 {
		t.Errorf("Expected 1 agent pool, instead got %d", len(p.AgentPoolProfiles))
	}

	if p.AgentPoolProfiles[0].Name != defaultAgentPoolName {
		t.Errorf("Expected LoadDefaultContainerServiceProperties() to return %s AgentPoolProfiles[0].Name, instead got %s", defaultAgentPoolName, p.AgentPoolProfiles[0].Name)
	}

	if p.AgentPoolProfiles[0].Count != defaultAgentCount {
		t.Errorf("Expected LoadDefaultContainerServiceProperties() to return %d AgentPoolProfiles[0].Count, instead got %d", defaultAgentCount, p.AgentPoolProfiles[0].Count)
	}

	if p.AgentPoolProfiles[0].VMSize != defaultVMSize {
		t.Errorf("Expected LoadDefaultContainerServiceProperties() to return %s AgentPoolProfiles[0].VMSize, instead got %s", defaultVMSize, p.AgentPoolProfiles[0].VMSize)
	}

	if p.AgentPoolProfiles[0].OSDiskSizeGB != defaultOSDiskSizeGB {
		t.Errorf("Expected LoadDefaultContainerServiceProperties() to return %d AgentPoolProfiles[0].OSDiskSizeGB, instead got %d", defaultOSDiskSizeGB, p.AgentPoolProfiles[0].OSDiskSizeGB)
	}

	if p.LinuxProfile.AdminUsername != defaultAdminUser {
		t.Errorf("Expected LoadDefaultContainerServiceProperties() to return %s LinuxProfile.AdminAdminUsernameUsername, instead got %s", defaultAdminUser, p.LinuxProfile.AdminUsername)
	}
}

func TestSerializeContainerService(t *testing.T) {
	cs := getDefaultContainerService()
	apiloader := &Apiloader{
		Translator: &i18n.Translator{},
	}

	b, err := apiloader.SerializeContainerService(cs, vlabs.APIVersion)

	if err != nil {
		t.Errorf("unexpected error while trying to Serialize Container Service: %s", err.Error())
	}

	expected := `{
  "apiVersion": "vlabs",
  "id": "sampleID",
  "location": "westus2",
  "name": "sampleCS",
  "plan": {
    "name": "sampleRPP",
    "product": "sampleProduct",
    "promotionCode": "sampleCode",
    "publisher": "samplePublisher"
  },
  "tags": {
    "foo": "bar"
  },
  "type": "sampleType",
  "properties": {
    "orchestratorProfile": {
      "orchestratorType": "Kubernetes",
      "orchestratorRelease": "1.11",
      "orchestratorVersion": "1.11.6",
      "kubernetesConfig": {
        "cloudProviderBackoffMode": ""
      }
    },
    "masterProfile": {
      "count": 1,
      "dnsPrefix": "blueorange",
      "subjectAltNames": [
        "fooSubjectAltName"
      ],
      "vmSize": "Standard_DS1_v1",
      "osDiskSizeGB": 256,
      "vnetSubnetID": "sampleVnetSubnetID",
      "vnetCidr": "10.240.0.0/8",
      "agentVnetSubnetID": "sampleAgentVnetSubnetID",
      "firstConsecutiveStaticIP": "10.240.0.0",
      "ipAddressCount": 5,
      "storageProfile": "StorageAccount",
      "oauthEnabled": false,
      "preProvisionExtension": {
        "name": "sampleExtension",
        "singleOrAll": "single",
        "template": "{{foobar}}"
      },
      "extensions": [
        {
          "name": "sampleExtension",
          "singleOrAll": "single",
          "template": "{{foobar}}"
        }
      ],
      "distro": "ubuntu",
      "kubernetesConfig": {
        "kubernetesImageBase": "quay.io",
        "clusterSubnet": "fooClusterSubnet",
        "dnsServiceIP": "172.0.0.1",
        "serviceCidr": "172.0.0.1/16",
        "networkPolicy": "calico",
        "networkPlugin": "azure-cni",
        "containerRuntime": "docker",
        "maxPods": 3,
        "dockerBridgeSubnet": "sampleDockerSubnet",
        "useManagedIdentity": true,
        "userAssignedID": "fooUserAssigneID",
        "userAssignedClientID": "fooUserAssigneClientID",
        "mobyVersion": "3.0.0",
        "containerdVersion": "1.2.4",
        "customCcmImage": "sampleCCMImage",
        "useCloudControllerManager": true,
        "customWindowsPackageURL": "https://deisartifacts.windows.net",
        "windowsNodeBinariesURL": "https://deisartifacts.windows.net",
        "useInstanceMetadata": true,
        "enableRbac": true,
        "enableSecureKubelet": true,
        "enableAggregatedAPIs": true,
        "privateCluster": {
          "enabled": true,
          "enableHostsConfigAgent": true,
          "jumpboxProfile": {
            "name": "sampleJumpboxProfile",
            "vmSize": "Standard_DS1_v2",
            "osDiskSizeGB": 512,
            "username": "userName",
            "publicKey": "ssh-rsa AAAAB3NzaC1yc2EAAAABJQAAAQEApD8+lRvLtUcyfO8N2Cwq0zY9DG1Un9d+tcmU3HgnAzBr6UR/dDT5M07NV7DN1lmu/0dt6Ay/ItjF9xK//nwVJL3ezEX32yhLKkCKFMB1LcANNzlhT++SB5tlRBx65CTL8z9FORe4UCWVJNafxu3as/BshQSrSaYt3hjSeYuzTpwd4+4xQutzbTXEUBDUr01zEfjjzfUu0HDrg1IFae62hnLm3ajG6b432IIdUhFUmgjZDljUt5bI3OEz5IWPsNOOlVTuo6fqU8lJHClAtAlZEZkyv0VotidC7ZSCfV153rRsEk9IWscwL2PQIQnCw7YyEYEffDeLjBwkH6MIdJ6OgQ== rsa-key-20170510",
            "storageProfile": "StorageAccount"
          }
        },
        "gchighthreshold": 85,
        "gclowthreshold": 80,
        "etcdVersion": "3.0.0",
        "etcdDiskSizeGB": "256",
        "etcdEncryptionKey": "sampleEncruptionKey",
        "enableDataEncryptionAtRest": true,
        "enableEncryptionWithExternalKms": true,
        "enablePodSecurityPolicy": true,
        "addons": [
          {
            "name": "sampleAddon",
            "enabled": true,
            "containers": [
              {
                "name": "sampleK8sContainer",
                "image": "sampleK8sImage",
                "cpuRequests": "10m",
                "memoryRequests": "20Mi"
              }
            ],
            "config": {
              "sampleKey": "sampleVal"
            }
          }
        ],
        "containerRuntimeConfig": {
          "dataDir": "/mnt/docker"
        },
        "kubeletConfig": {
          "barKey": "bazValue"
        },
        "controllerManagerConfig": {
          "sampleCMKey": "sampleCMVal"
        },
        "cloudControllerManagerConfig": {
          "sampleCCMKey": "sampleCCMVal"
        },
        "apiServerConfig": {
          "sampleAPIServerKey": "sampleAPIServerVal"
        },
        "schedulerConfig": {
          "sampleSchedulerKey": "sampleSchedulerVal"
        },
        "podSecurityPolicyConfig": {
          "samplePSPConfigKey": "samplePSPConfigVal"
        },
        "kubeReservedCgroup": "kubesystem.slice",
        "cloudProviderBackoffMode": "",
        "loadBalancerSku": "Basic",
        "excludeMasterFromStandardLB": false,
        "azureCNIVersion": "1.4.9",
        "azureCNIURLLinux": "https://mirror.azk8s.cn/kubernetes/azure-container-networking/linux",
        "azureCNIURLWindows": "https://mirror.azk8s.cn/kubernetes/azure-container-networking/windows",
        "keyVaultSku": "Basic",
        "maximumLoadBalancerRuleCount": 3,
        "kubeProxyMode": "iptables",
        "privateAzureRegistryServer": "sampleRegistryServerURL"
      },
      "imageReference": {
        "name": "FooImageRef",
        "resourceGroup": "FooImageRefResourceGroup"
      },
      "customFiles": [
        {
          "source": "sampleCustomFileSource",
          "dest": "sampleCustomFileDest"
        }
      ],
      "availabilityProfile": "",
      "platformUpdateDomainCount": null
    },
    "agentPoolProfiles": [
      {
        "name": "sampleAgent",
        "count": 2,
        "vmSize": "sampleVM",
        "dnsPrefix": "blueorange",
        "osType": "Linux",
        "availabilityProfile": "",
        "storageProfile": "",
        "ipAddressCount": 29,
        "fqdn": "blueorange.westus2.com",
        "preProvisionExtension": null,
        "extensions": [],
        "platformUpdateDomainCount": null
      },
      {
        "name": "sampleAgent-public",
        "count": 2,
        "vmSize": "sampleVM",
        "dnsPrefix": "blueorange",
        "osType": "Linux",
        "availabilityProfile": "",
        "storageProfile": "",
        "ipAddressCount": 29,
        "imageReference": {
          "name": "testImage",
          "resourceGroup": "testRg",
          "subscriptionId": "testSub",
          "gallery": "testGallery",
          "version": "0.0.1"
        },
        "fqdn": "blueorange.westus2.com",
        "preProvisionExtension": null,
        "extensions": [],
        "platformUpdateDomainCount": null
      }
    ],
    "linuxProfile": {
      "adminUsername": "azureuser",
      "ssh": {
        "publicKeys": [
          {
            "keyData": "ssh-rsa AAAAB3NzaC1yc2EAAAABJQAAAQEApD8+lRvLtUcyfO8N2Cwq0zY9DG1Un9d+tcmU3HgnAzBr6UR/dDT5M07NV7DN1lmu/0dt6Ay/ItjF9xK//nwVJL3ezEX32yhLKkCKFMB1LcANNzlhT++SB5tlRBx65CTL8z9FORe4UCWVJNafxu3as/BshQSrSaYt3hjSeYuzTpwd4+4xQutzbTXEUBDUr01zEfjjzfUu0HDrg1IFae62hnLm3ajG6b432IIdUhFUmgjZDljUt5bI3OEz5IWPsNOOlVTuo6fqU8lJHClAtAlZEZkyv0VotidC7ZSCfV153rRsEk9IWscwL2PQIQnCw7YyEYEffDeLjBwkH6MIdJ6OgQ== rsa-key-20170510"
          }
        ]
      },
      "secrets": [
        {
          "sourceVault": {
            "id": "sampleKeyVaultID"
          },
          "vaultCertificates": [
            {
              "certificateUrl": "FooCertURL",
              "certificateStore": "BarCertStore"
            }
          ]
        }
      ],
      "customSearchDomain": {
        "name": "FooCustomSearchDomain",
        "realmUser": "sampleRealmUser",
        "realmPassword": "sampleRealmPassword"
      },
      "customNodesDNS": {
        "dnsServer": "SampleDNSServer"
      }
    },
    "extensionProfiles": [
      {
        "name": "fooExtension",
        "version": "fooVersion",
        "extensionParameters": "fooExtensionParameters",
        "parametersKeyvaultSecretRef": {
          "vaultID": "fooVaultID",
          "secretName": "fooSecretName",
          "version": "fooSecretVersion"
        },
        "rootURL": "fooRootURL",
        "script": "fooSsript",
        "urlQuery": "fooURL"
      }
    ],
    "windowsProfile": {
      "adminUsername": "sampleAdminUsername",
      "adminPassword": "sampleAdminPassword",
      "WindowsImageSourceUrl": "",
      "WindowsPublisher": "",
      "WindowsOffer": "",
      "WindowsSku": "",
      "windowsDockerVersion": "",
      "windowsPauseImageURL": ""
    },
    "servicePrincipalProfile": {
      "clientId": "fooClientID",
      "secret": "fooSecret",
      "objectId": "fooObjectID",
      "keyvaultSecretRef": {
        "vaultID": "fooVaultID",
        "secretName": "fooSecretName",
        "version": "fooSecretVersion"
      }
    },
    "certificateProfile": {
      "caCertificate": "SampleCACert",
      "caPrivateKey": "SampleCAPrivateKey",
      "apiServerCertificate": "SampleAPIServerCert",
      "apiServerPrivateKey": "SampleAPIServerPrivateKey",
      "clientCertificate": "SampleClientCert",
      "clientPrivateKey": "SampleClientPrivateKey",
      "kubeConfigCertificate": "SampleKubeConfigCert",
      "kubeConfigPrivateKey": "SampleKubeConfigPrivateKey",
      "etcdServerCertificate": "SampleEtcdServerCert",
      "etcdServerPrivateKey": "SampleEtcdServerPrivateKey",
      "etcdClientCertificate": "SampleEtcdClientCert",
      "etcdClientPrivateKey": "SampleEtcdClientPrivateKey"
    },
    "aadProfile": {
      "clientAppID": "SampleClientAppID",
      "serverAppID": "ServerAppID",
      "tenantID": "SampleTenantID",
      "adminGroupID": "SampleAdminGroupID"
    },
    "featureFlags": {
      "enableCSERunInBackground": true
    }
  }
}
`
	if string(b) != expected {
		t.Errorf("expected SerializedCS JSON %s, but got %s", expected, string(b))
	}
}

func TestLoadCertificateProfileFromFile(t *testing.T) {
	locale := gotext.NewLocale(path.Join("..", "..", "translations"), "en_US")
	if err := i18n.Initialize(locale); err != nil {
		t.Error(err)
	}
	apiloader := &Apiloader{
		Translator: &i18n.Translator{
			Locale: locale,
		},
	}

	_, err := apiloader.LoadCertificateProfileFromFile("../engine/profiles/certificate-profile/kubernetes.json")
	if err != nil {
		t.Error(err.Error())
	}

	// Test error scenario
	_, err = apiloader.LoadCertificateProfileFromFile("../this-file-doesnt-exist.json")
	if err == nil {
		t.Errorf("expected error passing a non-existent filepath string to apiloader.LoadCertificateProfileFromFile(), instead got nil")
	}
}
