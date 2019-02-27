{{if HasCosmosEtcd }}
   {
     "apiVersion": "[variables('apiVersionCosmos')]",
     "type": "Microsoft.DocumentDB/databaseAccounts",
     "kind": "GlobalDocumentDB",
     "location": "[resourceGroup().location]",
     "name": "[variables('cosmosAccountName')]",
     "properties": {
       "capabilities": [
         {
           "name": "EnableEtcd"
         }
        ],
        "consistencyPolicy": {
          "defaultConsistencyLevel": "BoundedStaleness",
          "maxIntervalInSeconds": 5,
          "maxStalenessPrefix": 100
        },
        "databaseAccountOfferType": "Standard",
        "locations": [
          {
           "failoverPriority": 0,
           "locationName": "[resourceGroup().location]"
          },
          {
           "failoverPriority": 1,
           "locationName": "[resourceGroup().location]"
          }
        ],
        "name": "[variables('cosmosAccountName')]",
        "primaryClientCertificatePemBytes": "[variables('cosmosDBCertb64')]"
       },
     "tags": {
       "defaultExperience": "Etcd"
      }
     },
{{end}}
{{if EnableEncryptionWithExternalKms}}
  {
    "type": "Microsoft.Storage/storageAccounts",
    "name": "[variables('clusterKeyVaultName')]",
    "apiVersion": "[variables('apiVersionStorage')]",
    "location": "[variables('location')]",
    "sku": {
      "name": "Standard_LRS"
    }
  },
  {
    "type": "Microsoft.KeyVault/vaults",
    "name": "[variables('clusterKeyVaultName')]",
    "apiVersion": "[variables('apiVersionKeyVault')]",
    "location": "[variables('location')]",
    {{if UseManagedIdentity}}
    "dependsOn":
    [
      "[concat('Microsoft.Compute/virtualMachineScaleSets/', variables('masterVMNamePrefix'), 'vmss')]"
      {{if UserAssignedIDEnabled}}
      ,"[variables('userAssignedIDReference')]"
      {{end}}
    ],
    {{end}}
    "properties": {
      "enabledForDeployment": "false",
      "enabledForDiskEncryption": "false",
      "enabledForTemplateDeployment": "false",
      "tenantId": "[variables('tenantID')]",
    "accessPolicies": 
      [
        {
          "tenantId": "[variables('tenantID')]",
          {{if UseManagedIdentity}}
          {{if UserAssignedIDEnabled}}
          "objectId": "[reference(variables('userAssignedIDReference'), variables('apiVersionManagedIdentity')).principalId]",
          {{end}}
          {{else}}
          "objectId": "[parameters('servicePrincipalObjectId')]",
          {{end}}
          "permissions": {
            "keys": ["create", "encrypt", "decrypt", "get", "list"]
          }
        }
      ],
      "sku": {
        "name": "[parameters('clusterKeyVaultSku')]",
        "family": "A"
      }
    }
  },
{{end}}
{
  "apiVersion": "[variables('apiVersionNetwork')]",
  "location": "[variables('location')]",
  "name": "[variables('nsgName')]",
  "properties": {
    "securityRules": [
    {{if .HasWindows}}
      {
        "name": "allow_rdp",
        "properties": {
          "access": "Allow",
          "description": "Allow RDP traffic to master",
          "destinationAddressPrefix": "*",
          "destinationPortRange": "3389-3389",
          "direction": "Inbound",
          "priority": 102,
          "protocol": "Tcp",
          "sourceAddressPrefix": "*",
          "sourcePortRange": "*"
        }
      },
    {{end}}
      {
        "name": "allow_ssh",
        "properties": {
          "access": "Allow",
          "description": "Allow SSH traffic to master",
          "destinationAddressPrefix": "*",
          "destinationPortRange": "22-22",
          "direction": "Inbound",
          "priority": 101,
          "protocol": "Tcp",
          "sourceAddressPrefix": "*",
          "sourcePortRange": "*"
        }
      },
      {
        "name": "allow_kube_tls",
        "properties": {
          "access": "Allow",
          "description": "Allow kube-apiserver (tls) traffic to master",
          "destinationAddressPrefix": "*",
          "destinationPortRange":"443-443",
          "direction": "Inbound",
          "priority": 100,
          "protocol": "Tcp",
          "sourceAddressPrefix": "*",
          "sourcePortRange": "*"
        }
      }
    {{if IsAzureStackCloud}}
      ,{
        "name": "allow_vnet_inbound",
        "properties": {
          "access": "Allow",
          "description": "Allow traffic to specific addresses.",
          "destinationAddressPrefix": "10.0.0.0/8",
          "destinationPortRange": "*",
          "direction": "Inbound",
          "priority": 4095,
          "protocol": "*",
          "sourceAddressPrefix": "10.0.0.0/8",
          "sourcePortRange": "*"
        }
      },
      {
        "name": "allow_vnet_outbound",
        "properties": {
          "access": "Allow",
          "description": "Allow traffic to specific addresses.",
          "destinationAddressPrefix": "10.0.0.0/8",
          "destinationPortRange": "*",
          "direction": "Outbound",
          "priority": 4095,
          "protocol": "*",
          "sourceAddressPrefix": "10.0.0.0/8",
          "sourcePortRange": "*"
        }
      }
    {{end}}
    {{if IsFeatureEnabled "BlockOutboundInternet"}}
      ,{
        "name": "allow_vnet",
        "properties": {
          "access": "Allow",
          "description": "Allow outbound internet to vnet",
          "destinationAddressPrefix": "[parameters('masterSubnet')]",
          "destinationPortRange": "*",
          "direction": "Outbound",
          "priority": 110,
          "protocol": "*",
          "sourceAddressPrefix": "VirtualNetwork",
          "sourcePortRange": "*"
        }
      },
      {
        "name": "block_outbound",
        "properties": {
          "access": "Deny",
          "description": "Block outbound internet from master",
          "destinationAddressPrefix": "*",
          "destinationPortRange": "*",
          "direction": "Outbound",
          "priority": 120,
          "protocol": "*",
          "sourceAddressPrefix": "*",
          "sourcePortRange": "*"
        }
      }
    {{end}}
    ]
  },
  "type": "Microsoft.Network/networkSecurityGroups"
},
{{if RequireRouteTable}}
{
  "apiVersion": "[variables('apiVersionNetwork')]",
  "location": "[variables('location')]",
  "name": "[variables('routeTableName')]",
  "type": "Microsoft.Network/routeTables"
},
{{end}}
{{if not .MasterProfile.IsCustomVNET}}
{
  "apiVersion": "[variables('apiVersionNetwork')]",
  "dependsOn": [
    {{if RequireRouteTable}}
    "[concat('Microsoft.Network/routeTables/', variables('routeTableName'))]",
    {{end}}
    "[concat('Microsoft.Network/networkSecurityGroups/', variables('nsgName'))]"
    ],
    "location": "[variables('location')]",
    "name": "[variables('virtualNetworkName')]",
    "properties": {
      "addressSpace": {
        "addressPrefixes": [
          "[parameters('vnetCidr')]"
        ]
      },
      "subnets": [
      {
        "name": "subnetmaster",
        "properties": {
          "addressPrefix": "[parameters('masterSubnet')]"
          ,"networkSecurityGroup": {
            "id": "[variables('nsgID')]"
          }
          {{if RequireRouteTable}}
          ,"routeTable": {
            "id": "[variables('routeTableID')]"
          }
          {{end}}
        }
      },
      {  
        "name":"subnetagent",
        "properties":{  
            "addressPrefix": "[parameters('agentSubnet')]",
            "networkSecurityGroup": {
            "id": "[variables('nsgID')]"
          }
          {{if RequireRouteTable}}
          ,"routeTable": {
            "id": "[variables('routeTableID')]"
          }
          {{end}}
        }
      }
      ]
    },
    "type": "Microsoft.Network/virtualNetworks"
},
{{end}}
{
  "apiVersion": "[variables('apiVersionNetwork')]",
  "location": "[variables('location')]",
  "name": "[variables('masterPublicIPAddressName')]",
  "properties": {
    "dnsSettings": {
      "domainNameLabel": "[variables('masterFqdnPrefix')]"
    },
    "publicIPAllocationMethod": "Static"
  },
  "sku": {
      "name": "[variables('loadBalancerSku')]"
  },
  "type": "Microsoft.Network/publicIPAddresses"
},
{
    "type": "Microsoft.Network/loadBalancers",
    "name": "[variables('masterLbName')]",
    "location": "[variables('location')]",
    "apiVersion": "[variables('apiVersionNetwork')]",
    "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
    ],
    "properties": {
        "frontendIPConfigurations": [
            {
                "name": "[variables('masterLbIPConfigName')]",
                "properties": {
                    "publicIPAddress": {
                        "id": "[resourceId('Microsoft.Network/publicIpAddresses', variables('masterPublicIPAddressName'))]"
                    }
                }
            }
        ],
        "backendAddressPools": [
            {
                "name": "[variables('masterLbBackendPoolName')]"
            }
        ],
        "probes": [
          {
              "name": "tcpHTTPSProbe",
              "properties": {
                  "protocol": "Tcp",
                  "port": 443,
                  "intervalInSeconds": 5,
                  "numberOfProbes": 2
              }
          }
        ],
        "inboundNatPools": [
          {
              "name": "[concat('SSH-', variables('masterVMNamePrefix'), 'natpools')]",
              "properties": {
                  "frontendIPConfiguration": {
                      "id": "[variables('masterLbIPConfigID')]"
                  },
                  "protocol": "Tcp",
                  "backendPort": "22",
                  "frontendPortRangeStart": "50001",
                  "frontendPortRangeEnd": "50119",
                  "enableFloatingIP": false
              }
          }
        ],
        "loadBalancingRules": [
          {
            "name": "LBRuleHTTPS",
            "properties": {
                "frontendIPConfiguration": {
                    "id": "[variables('masterLbIPConfigID')]"
                },
                "backendAddressPool": {
                    "id": "[concat(variables('masterLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"
                },
                "protocol": "Tcp",
                "frontendPort": 443,
                "backendPort": 443,
                "enableFloatingIP": false,
                "idleTimeoutInMinutes": 5,
                "loadDistribution": "Default",
                "probe": {
                    "id": "[concat(variables('masterLbID'),'/probes/tcpHTTPSProbe')]"
                }
            }
          }
        ]
    },
    "sku": {
        "name": "[variables('loadBalancerSku')]"
    }
},
{{if gt .MasterProfile.Count 1}}
    {
      "apiVersion": "[variables('apiVersionNetwork')]",
      "dependsOn": [
{{if .MasterProfile.IsCustomVNET}}
        "[variables('nsgID')]"
{{else}}
        "[variables('vnetID')]"
{{end}}
      ],
      "location": "[variables('location')]",
      "name": "[variables('masterInternalLbName')]",
      "properties": {
        "backendAddressPools": [
          {
            "name": "[variables('masterLbBackendPoolName')]"
          }
        ],
        "frontendIPConfigurations": [
          {
            "name": "[variables('masterInternalLbIPConfigName')]",
            "properties": {
              "privateIPAddress": "[variables('kubernetesAPIServerIP')]",
              "privateIPAllocationMethod": "Static",
              "subnet": {
                "id": "[variables('vnetSubnetIDMaster')]"
              }
            }
          }
        ],
        "loadBalancingRules": [
          {
            "name": "InternalLBRuleHTTPS",
            "properties": {
              "backendAddressPool": {
                "id": "[concat(variables('masterInternalLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"
              },
              "backendPort": 4443,
              "enableFloatingIP": false,
              "frontendIPConfiguration": {
                "id": "[variables('masterInternalLbIPConfigID')]"
              },
              "frontendPort": 443,
              "idleTimeoutInMinutes": 5,
              "protocol": "Tcp",
              "probe": {
                "id": "[concat(variables('masterInternalLbID'),'/probes/tcpHTTPSProbe')]"
              }
            }
          }
        ],
        "probes": [
          {
            "name": "tcpHTTPSProbe",
            "properties": {
              "intervalInSeconds": 5,
              "numberOfProbes": 2,
              "port": 4443,
              "protocol": "Tcp"
            }
          }
        ]
      },
      "sku": {
        "name": "[variables('loadBalancerSku')]"
      },
      "type": "Microsoft.Network/loadBalancers"
    },
{{end}}
{
    "apiVersion": "[variables('apiVersionCompute')]",
    "dependsOn": [
    {{if .MasterProfile.IsCustomVNET}}
      "[variables('nsgID')]"
    {{else}}
      "[variables('vnetID')]"
    {{end}}
    {{if gt .MasterProfile.Count 1}}
        ,"[variables('masterInternalLbName')]"
    {{end}}
    {{ if HasCosmosEtcd }}
      ,"[resourceId('Microsoft.DocumentDB/databaseAccounts/', variables('cosmosAccountName'))]"
    {{ end }}
      ,"[variables('masterLbID')]"
    ],
    "tags":
    {
      "creationSource": "[concat(parameters('generatorCode'), '-', variables('masterVMNamePrefix'), 'vmss')]",
      "resourceNameSuffix": "[parameters('nameSuffix')]",
      "orchestrator": "[variables('orchestratorNameVersionTag')]",
      "aksEngineVersion" : "[parameters('aksEngineVersion')]",
      "poolName": "master"
    },
    "location": "[variables('location')]",
    {{ if .MasterProfile.HasAvailabilityZones}}
    "zones": "[parameters('availabilityZones')]",
    {{ end }}
    "name": "[concat(variables('masterVMNamePrefix'), 'vmss')]",
    {{if UseManagedIdentity}}
    {{if UserAssignedIDEnabled}}
    "identity": {
      "type": "userAssigned",
      "userAssignedIdentities": {
        "[variables('userAssignedIDReference')]":{}
      }
    },
    {{end}}
    {{end}}
    "sku": {
      "tier": "Standard",
      "capacity": "[variables('masterCount')]",
      "name": "[parameters('masterVMSize')]"
    },
    "properties": {
      "singlePlacementGroup": {{ .MasterProfile.SinglePlacementGroup}},
      "overprovision": false,
      "upgradePolicy": {
        "mode": "Manual"
      },
      "virtualMachineProfile": {
        "networkProfile": {
          "networkInterfaceConfigurations": [
            {
              "name": "[concat(variables('masterVMNamePrefix'), 'netintconfig')]",
              "properties": {
                "primary": true,
                {{if .MasterProfile.IsCustomVNET}}
                "networkSecurityGroup": {
                  "id": "[variables('nsgID')]"
                },
                {{end}}
                "ipConfigurations": [
                  {{range $seq := loop 1 .MasterProfile.IPAddressCount}}
                  {
                    "name": "ipconfig{{$seq}}",
                    "properties": {
                      {{if eq $seq 1}}
                      "loadBalancerBackendAddressPools": [
                        {
                          "id": "[concat(variables('masterLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"
                        }
                        {{if gt $.MasterProfile.Count 1}}
                          ,{
                            "id": "[concat(variables('masterInternalLbID'), '/backendAddressPools/', variables('masterLbBackendPoolName'))]"
                          }
                        {{end}}
                      ],
                      "loadBalancerInboundNatPools": [
                        {
                          "id": "[concat(variables('masterLbID'),'/inboundNatPools/SSH-', variables('masterVMNamePrefix'), 'natpools')]"
                        }
                      ],
                      "primary": true,
                      {{else}}
                      "primary": false,
                      {{end}}
                      "subnet": {
                        "id": "[variables('vnetSubnetIDMaster')]"
                      }
                    }
                  }
                  {{if lt $seq $.MasterProfile.IPAddressCount}},{{end}}
                  {{end}}
                ]
                {{if HasCustomNodesDNS}}
                 ,"dnsSettings": {
                    "dnsServers": [
                        "[parameters('dnsServer')]"
                    ]
                }
                {{end}}
                {{if not IsAzureCNI}}
                ,"enableIPForwarding": true
                {{end}}
              }
            }
          ]
        },
        "osProfile": {
          "adminUsername": "[parameters('linuxAdminUsername')]",
          "computerNamePrefix": "[concat(variables('masterVMNamePrefix'), 'vmss')]",
          {{GetKubernetesMasterCustomData .}}
          "linuxConfiguration": {
              "disablePasswordAuthentication": true,
              {{if HasMultipleSshKeys }}
              "ssh": {{ GetSshPublicKeys }}
              {{ else }}
              "ssh": {
                "publicKeys": [
                  {
                    "keyData": "[parameters('sshRSAPublicKey')]",
                    "path": "[variables('sshKeyPath')]"
                  }
                ]
              }
              {{ end }}
            }
            {{if .LinuxProfile.HasSecrets}}
              ,
              "secrets": "[variables('linuxProfileSecrets')]"
            {{end}}
        },
        "storageProfile": {
          {{if not UseMasterCustomImage }}
          "dataDisks": [
            {
              "createOption": "Empty",
              "diskSizeGB": "[parameters('etcdDiskSizeGB')]",
              "lun": 0
            }
          ],
          {{end}}
          "imageReference": {
            {{if UseMasterCustomImage}}
            "id": "[resourceId(parameters('osImageResourceGroup'), 'Microsoft.Compute/images', parameters('osImageName'))]"
            {{else}}
            "offer": "[parameters('osImageOffer')]",
            "publisher": "[parameters('osImagePublisher')]",
            "sku": "[parameters('osImageSku')]",
            "version": "[parameters('osImageVersion')]"
            {{end}}
          },
          "osDisk": {
            "caching": "ReadWrite"
            ,"createOption": "FromImage"
            {{if ne .MasterProfile.OSDiskSizeGB 0}}
            ,"diskSizeGB": {{.MasterProfile.OSDiskSizeGB}}
            {{end}}
          }
        },
        "extensionProfile": {
          "extensions": [
            {
              "name": "[concat(variables('masterVMNamePrefix'), 'vmssCSE')]",
              "properties": {
                "publisher": "Microsoft.Azure.Extensions",
                "type": "CustomScript",
                "typeHandlerVersion": "2.0",
                "autoUpgradeMinorVersion": true,
                "settings": {},
                "protectedSettings": {
                  "commandToExecute": "[concat('retrycmd_if_failure() { r=$1; w=$2; t=$3; shift && shift && shift; for i in $(seq 1 $r); do timeout $t ${@}; [ $? -eq 0  ] && break || if [ $i -eq $r ]; then return 1; else sleep $w; fi; done };{{if not (IsFeatureEnabled "BlockOutboundInternet")}} ERR_OUTBOUND_CONN_FAIL=50; retrycmd_if_failure 50 1 3 nc -vz {{if IsMooncake}}gcr.azk8s.cn 80{{else}}k8s.gcr.io 443 && retrycmd_if_failure 50 1 3 nc -vz gcr.io 443 && retrycmd_if_failure 50 1 3 nc -vz docker.io 443{{end}} || exit $ERR_OUTBOUND_CONN_FAIL;{{end}} for i in $(seq 1 1200); do if [ -f /opt/azure/containers/provision.sh ]; then break; fi; if [ $i -eq 1200 ]; then exit 100; else sleep 1; fi; done; ', variables('provisionScriptParametersCommon'),' ',variables('provisionScriptParametersMaster'), ' /usr/bin/nohup /bin/bash -c \"/bin/bash /opt/azure/containers/provision.sh >> /var/log/azure/cluster-provision.log 2>&1\"')]"
                }
              }
            }
            {{if UseAksExtension}}
            ,{
              "name": "[concat(variables('masterVMNamePrefix'), 'vmss-computeAksLinuxBilling')]",
              "properties": {
                "publisher": "Microsoft.AKS",
                "type": "Compute.AKS-Engine.Linux.Billing",
                "typeHandlerVersion": "1.0",
                "autoUpgradeMinorVersion": true,
                "settings": {}
              }
            }
            {{end}}
          ]
        }
      }
    },
    "type": "Microsoft.Compute/virtualMachineScaleSets"
  }
