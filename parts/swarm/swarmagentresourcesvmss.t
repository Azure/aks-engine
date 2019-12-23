{{if .IsStorageAccount}}
    {
      "apiVersion": "[variables('apiVersionStorage')]",
      "copy": {
        "count": "[variables('{{.Name}}StorageAccountsCount')]",
        "name": "vmLoopNode"
      },
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('storageAccountPrefixes')[mod(add(copyIndex(),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(copyIndex(),variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
      "properties": {
        "accountType": "[variables('vmSizesMap')[variables('{{.Name}}VMSize')].storageAccountType]"
      },
      "type": "Microsoft.Storage/storageAccounts"
    },
{{end}}
{{if IsPublic .Ports}}
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}IPAddressName')]",
      "properties": {
        "dnsSettings": {
          "domainNameLabel": "[variables('{{.Name}}EndpointDNSNamePrefix')]"
        },
        "publicIPAllocationMethod": "Dynamic"
      },
      "type": "Microsoft.Network/publicIPAddresses"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('{{.Name}}IPAddressName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[variables('{{.Name}}LbName')]",
      "properties": {
        "backendAddressPools": [
          {
            "name": "[variables('{{.Name}}LbBackendPoolName')]"
          }
        ],
        "frontendIPConfigurations": [
          {
            "name": "[variables('{{.Name}}LbIPConfigName')]",
            "properties": {
              "publicIPAddress": {
                "id": "[resourceId('Microsoft.Network/publicIPAddresses',variables('{{.Name}}IPAddressName'))]"
              }
            }
          }
        ],
        "inboundNatRules": [],
        "loadBalancingRules": [
          {{(GetLBRules .Name .Ports)}}
        ],
        "probes": [
          {{(GetProbes .Ports)}}
        ]
      },
      "type": "Microsoft.Network/loadBalancers"
    },
{{end}}
    {
{{if .IsManagedDisks}}
      "apiVersion": "[variables('apiVersionStorageManagedDisks')]",
{{else}}
      "apiVersion": "[variables('apiVersionDefault')]",
{{end}}
      "dependsOn": [
        "[concat('Microsoft.Network/publicIPAddresses/', variables('masterPublicIPAddressName'))]"
{{if .IsStorageAccount}}
        ,"[concat('Microsoft.Storage/storageAccounts/', variables('storageAccountPrefixes')[mod(add(0,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(0,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
        "[concat('Microsoft.Storage/storageAccounts/', variables('storageAccountPrefixes')[mod(add(1,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(1,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
		"[concat('Microsoft.Storage/storageAccounts/', variables('storageAccountPrefixes')[mod(add(2,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(2,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
        "[concat('Microsoft.Storage/storageAccounts/', variables('storageAccountPrefixes')[mod(add(3,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(3,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]",
        "[concat('Microsoft.Storage/storageAccounts/', variables('storageAccountPrefixes')[mod(add(4,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(4,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName'))]"
{{end}}
{{if not .IsCustomVNET}}
      ,"[variables('vnetID')]"
{{end}}
{{if IsPublic .Ports}}
       ,"[variables('{{.Name}}LbID')]"
{{end}}
      ],
      "tags":
      {
        "creationSource" : "[concat('acsengine-', variables('{{.Name}}VMNamePrefix'), '-vmss')]"
      },
      "location": "[variables('location')]",
      "name": "[concat(variables('{{.Name}}VMNamePrefix'), '-vmss')]",
      "properties": {
        "upgradePolicy": {
          "mode": "Automatic"
        },
        "virtualMachineProfile": {
          "networkProfile": {
            "networkInterfaceConfigurations": [
              {
                "name": "nic",
                "properties": {
                  "ipConfigurations": [
                    {
                      "name": "nicipconfig",
                      "properties": {
{{if IsPublic .Ports}}
                        "loadBalancerBackendAddressPools": [
                          {
                            "id": "[concat(variables('{{.Name}}LbID'), '/backendAddressPools/', variables('{{.Name}}LbBackendPoolName'))]"
                          }
                        ],
{{end}}
                        "subnet": {
                          "id": "[variables('{{.Name}}VnetSubnetID')]"
                        }
                      }
                    }
                  ],
                  "primary": "true"
                }
              }
            ]
          },
          "osProfile": {
            "adminUsername": "[variables('adminUsername')]",
            "computerNamePrefix": "[variables('{{.Name}}VMNamePrefix')]",
            {{GetAgentSwarmCustomData .}}
            "linuxConfiguration": {
              "disablePasswordAuthentication": true,
              "ssh": {
                "publicKeys": [
                  {
                    "keyData": "[parameters('sshRSAPublicKey')]",
                    "path": "[variables('sshKeyPath')]"
                  }
                ]
              }
            }
            {{if HasLinuxSecrets}}
              ,
              "secrets": "[variables('linuxProfileSecrets')]"
            {{end}}
          },
          "storageProfile": {
            "imageReference": {
              "offer": "[variables('{{.Name}}OSImageOffer')]",
              "publisher": "[variables('{{.Name}}OSImagePublisher')]",
              "sku": "[variables('{{.Name}}OSImageSKU')]",
              "version": "[variables('{{.Name}}OSImageVersion')]"
            },
            {{GetDataDisks .}}
            "osDisk": {
              "caching": "ReadWrite"
              ,"createOption": "FromImage"
{{if .IsStorageAccount}}
              ,"name": "vmssosdisk"
              ,"vhdContainers": [
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(0,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(0,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')), variables('apiVersionStorage') ).primaryEndpoints.blob, 'osdisk')]",
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(1,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(1,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')), variables('apiVersionStorage')).primaryEndpoints.blob, 'osdisk')]",
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(2,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(2,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')), variables('apiVersionStorage')).primaryEndpoints.blob, 'osdisk')]",
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(3,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(3,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')), variables('apiVersionStorage')).primaryEndpoints.blob, 'osdisk')]",
                "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('storageAccountPrefixes')[mod(add(4,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('storageAccountPrefixes')[div(add(4,variables('{{.Name}}StorageAccountOffset')),variables('storageAccountPrefixesCount'))],variables('{{.Name}}AccountName')), variables('apiVersionStorage')).primaryEndpoints.blob, 'osdisk')]"
              ]
{{end}}
{{if ne .OSDiskSizeGB 0}}
            ,"diskSizeGB": {{.OSDiskSizeGB}}
{{end}}
            }
          }
{{if .IsRHEL}}
          ,"extensionProfile": {
            "extensions": [
              {
                "name": "configure{{.Name}}",
                "properties": {
                  "publisher": "Microsoft.Azure.Extensions",
                  "settings": {
                    "commandToExecute": "[variables('agentCustomScript')]",
                    "fileUris": [
                      "[concat('{{ GetConfigurationScriptRootURL }}', variables('configureClusterScriptFile'))]"
                    ]
                  },
                  "type": "CustomScript",
                  "typeHandlerVersion": "2.0"
                }
              }
            ]
          }
{{end}}
        }
      },
      "sku": {
        "capacity": "[variables('{{.Name}}Count')]",
        "name": "[variables('{{.Name}}VMSize')]",
        "tier": "[variables('{{.Name}}VMSizeTier')]"
      },
      "type": "Microsoft.Compute/virtualMachineScaleSets"
    }
