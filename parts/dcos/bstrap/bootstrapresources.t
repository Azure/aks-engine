    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "location": "[variables('location')]",
      "name": "[variables('bootstrapNSGName')]",
      "properties": {
        "securityRules": [
            {
                "properties": {
                    "priority": 200,
                    "access": "Allow",
                    "direction": "Inbound",
                    "destinationPortRange": "22",
                    "sourcePortRange": "*",
                    "destinationAddressPrefix": "*",
                    "protocol": "Tcp",
                    "description": "Allow SSH",
                    "sourceAddressPrefix": "*"
                },
                "name": "ssh"
            },
            {
                "properties": {
                    "priority": 201,
                    "access": "Allow",
                    "direction": "Inbound",
                    "destinationPortRange": "8086",
                    "sourcePortRange": "*",
                    "destinationAddressPrefix": "*",
                    "protocol": "Tcp",
                    "description": "Allow bootstrap service",
                    "sourceAddressPrefix": "*"
                },
                "name": "Port8086"
            }
        ]
      },
      "type": "Microsoft.Network/networkSecurityGroups"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "dependsOn": [
{{if not .MasterProfile.IsCustomVNET}}
        "[variables('vnetID')]",
{{end}}
        "[variables('bootstrapNSGID')]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('bootstrapVMName'), '-nic')]",
      "properties": {
        "ipConfigurations": [
          {
            "name": "ipConfigNode",
            "properties": {
              "privateIPAddress": "[variables('bootstrapStaticIP')]",
              "privateIPAllocationMethod": "Static",
              "subnet": {
                "id": "[variables('masterVnetSubnetID')]"
              }
            }
          }
        ],
        "networkSecurityGroup": {
          "id": "[variables('bootstrapNSGID')]"
        }
      },
      "type": "Microsoft.Network/networkInterfaces"
    },
    {
      "apiVersion": "[variables('apiVersionStorageManagedDisks')]",
      "dependsOn": [
        "[concat('Microsoft.Network/networkInterfaces/', variables('bootstrapVMName'), '-nic')]",
{{if .MasterProfile.IsStorageAccount}}
        "[variables('masterStorageAccountName')]",
{{end}}
        "[variables('masterStorageAccountExhibitorName')]"
      ],
      "tags":
      {
        "creationSource": "[concat('acsengine-', variables('bootstrapVMName'))]",
        "orchestratorName": "dcos",
        "orchestratorVersion": "[variables('orchestratorVersion')]",
        "orchestratorNode": "bootstrap"
      },
      "location": "[variables('location')]",
      "name": "[variables('bootstrapVMName')]",
      "properties": {
        "hardwareProfile": {
          "vmSize": "[variables('bootstrapVMSize')]"
        },
        "networkProfile": {
          "networkInterfaces": [
            {
              "id": "[resourceId('Microsoft.Network/networkInterfaces',concat(variables('bootstrapVMName'), '-nic'))]"
            }
          ]
        },
        "osProfile": {
          "adminUsername": "[variables('adminUsername')]",
          "computername": "[variables('bootstrapVMName')]",
          {{GetDCOSBootstrapCustomData}}
          "linuxConfiguration": {
            "disablePasswordAuthentication": true,
            "ssh": {
                "publicKeys": [
                    {
                        "keyData": "[variables('sshRSAPublicKey')]",
                        "path": "[variables('sshKeyPath')]"
                    }
                ]
            }
          }
          {{if .LinuxProfile.HasSecrets}}
          ,
          "secrets": "[variables('linuxProfileSecrets')]"
          {{end}}
        },
        "storageProfile": {
          "imageReference": {
            "offer": "[variables('osImageOffer')]",
            "publisher": "[variables('osImagePublisher')]",
            "sku": "[variables('osImageSKU')]",
            "version": "[variables('osImageVersion')]"
          },
          "osDisk": {
            "caching": "ReadWrite"
            ,"createOption": "FromImage"
{{if .MasterProfile.IsStorageAccount}}
            ,"name": "[concat(variables('bootstrapVMName'), '-osdisk')]"
            ,"vhd": {
              "uri": "[concat(reference(concat('Microsoft.Storage/storageAccounts/',variables('masterStorageAccountName')),variables('apiVersionStorage')).primaryEndpoints.blob,'vhds/',variables('bootstrapVMName'),-osdisk.vhd')]"
            }
{{end}}
{{if ne .OrchestratorProfile.DcosConfig.BootstrapProfile.OSDiskSizeGB 0}}
            ,"diskSizeGB": "60"
{{end}}
          }
        }
      },
      "type": "Microsoft.Compute/virtualMachines"
    },
    {
      "apiVersion": "[variables('apiVersionDefault')]",
      "dependsOn": [
        "[concat('Microsoft.Compute/virtualMachines/', variables('bootstrapVMName'))]"
      ],
      "location": "[variables('location')]",
      "name": "[concat(variables('bootstrapVMName'), '/bootstrapready')]",
      "properties": {
        "autoUpgradeMinorVersion": true,
        "publisher": "Microsoft.OSTCExtensions",
        "settings": {
          "commandToExecute": "[concat('/bin/bash -c \"until curl -f http://', variables('bootstrapStaticIP'), ':8086/dcos_install.sh > /dev/null; do echo waiting for bootstrap node; sleep 15; done; echo bootstrap node up\"')]"
        },
        "type": "CustomScriptForLinux",
        "typeHandlerVersion": "1.4"
      },
      "type": "Microsoft.Compute/virtualMachines/extensions"
    }{{WriteLinkedTemplatesForExtensions}}
