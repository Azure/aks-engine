{{if UseManagedIdentity}}
  {
    "apiVersion": "[variables('apiVersionAuthorization')]",
    "name": "[guid(concat('Microsoft.Compute/virtualMachineScaleSets/', variables('{{.Name}}VMNamePrefix'), 'vmidentity'))]",
    "type": "Microsoft.Authorization/roleAssignments",
    "properties": {
      "roleDefinitionId": "[variables('readerRoleDefinitionId')]",
      "principalId": "[reference(concat('Microsoft.Compute/virtualMachineScaleSets/', variables('{{.Name}}VMNamePrefix')), '2017-03-30', 'Full').identity.principalId]"
    }
  },
{{end}}
  {
    "apiVersion": "[variables('apiVersionCompute')]",
    "dependsOn": [
    {{if .IsCustomVNET}}
      "[variables('nsgID')]"
    {{else}}
      "[variables('vnetID')]"
    {{end}}
    ],
    "tags":
    {
      "creationSource" : "[concat(parameters('generatorCode'), '-', variables('{{.Name}}VMNamePrefix'))]",
      "resourceNameSuffix" : "[variables('winResourceNamePrefix')]",
      "orchestrator" : "[variables('orchestratorNameVersionTag')]",
      "poolName" : "{{.Name}}"
    },
    "location": "[variables('location')]",
    {{ if HasAvailabilityZones .}}
    "zones": "[parameters('{{.Name}}AvailabilityZones')]",
    {{ end }}
    "name": "[variables('{{.Name}}VMNamePrefix')]",
    {{if UseManagedIdentity}}
    "identity": {
      "type": "systemAssigned"
    },
    {{end}}
    "sku": {
      "tier": "Standard",
      "capacity": "[variables('{{.Name}}Count')]",
      "name": "[variables('{{.Name}}VMSize')]"
    },
    "properties": {
      "singlePlacementGroup": {{UseSinglePlacementGroup .}},
      "overprovision": false,
      "upgradePolicy": {
        "mode": "Manual"
      },
      "virtualMachineProfile": {
        "networkProfile": {
          "networkInterfaceConfigurations": [
            {
              "name": "[variables('{{.Name}}VMNamePrefix')]",
              "properties": {
                "primary": true,
                "enableAcceleratedNetworking" : "{{.AcceleratedNetworkingEnabledWindows}}",
                {{if .IsCustomVNET}}
                "networkSecurityGroup": {
                  "id": "[variables('nsgID')]"
                },
                {{end}}
                "ipConfigurations": [
                  {{range $seq := loop 1 .IPAddressCount}}
                  {
                    "name": "ipconfig{{$seq}}",
                    "properties": {
                      {{if eq $seq 1}}
                      "primary": true,
                      {{end}}
                      "subnet": {
                        "id": "[variables('{{$.Name}}VnetSubnetID')]"
                      }
                    }
                  }
                  {{if lt $seq $.IPAddressCount}},{{end}}
                  {{end}}
                ]
                {{if not IsAzureCNI}}
                ,"enableIPForwarding": true
                {{end}}
              }
            }
          ]
        },
        "osProfile": {
          "computerNamePrefix": "[variables('{{.Name}}VMNamePrefix')]",
          {{GetKubernetesWindowsAgentCustomData .}}
          "adminUsername": "[parameters('windowsAdminUsername')]",
          "adminPassword": "[parameters('windowsAdminPassword')]"
        },
        "storageProfile": {
          {{GetDataDisks .}}
          "imageReference": {
            "offer": "[parameters('agentWindowsOffer')]",
            "publisher": "[parameters('agentWindowsPublisher')]",
            "sku": "[parameters('agentWindowsSku')]",
            "version": "[parameters('agentWindowsVersion')]"
          },
          "osDisk": {
            "createOption": "FromImage",
            "caching": "ReadWrite"
          {{if ne .OSDiskSizeGB 0}}
            ,"diskSizeGB": {{.OSDiskSizeGB}}
          {{end}}
          }
        },
        "extensionProfile": {
          "extensions": [
            {
              "name": "vmssCSE",
              "properties": {
                "publisher": "Microsoft.Compute",
                "type": "CustomScriptExtension",
                "typeHandlerVersion": "1.8",
                "autoUpgradeMinorVersion": true,
                "settings": {},
                "protectedSettings": {
                    "commandToExecute": "[concat('powershell.exe -ExecutionPolicy Unrestricted -command \"', '$arguments = ', variables('singleQuote'),'-MasterIP ',variables('kubernetesAPIServerIP'),' -KubeDnsServiceIp ',parameters('kubeDnsServiceIp'),' -MasterFQDNPrefix ',variables('masterFqdnPrefix'),' -Location ',variables('location'),' -AgentKey ',parameters('clientPrivateKey'),' -AADClientId ',variables('servicePrincipalClientId'),' -AADClientSecret ',variables('servicePrincipalClientSecret'),variables('singleQuote'), ' ; ', variables('windowsCustomScriptSuffix'), '\" > %SYSTEMDRIVE%\\AzureData\\CustomDataSetupScript.log 2>&1')]"
                }
              }
            }
            {{if UseAksExtension}}
            ,{
              "name": "[concat(variables('{{.Name}}VMNamePrefix'), '-computeAksLinuxBilling')]",
              "location": "[variables('location')]",
              "properties": {
                "publisher": "Microsoft.AKS",
                "type": "Compute.AKS-Engine.Windows.Billing",
                "typeHandlerVersion": "1.0",
                "autoUpgradeMinorVersion": true,
                "settings": {}
              }
            }
            {{end}}
            {{if UseManagedIdentity}}
            ,{
              "name": "managedIdentityExtension",
              "properties": {
                "publisher": "Microsoft.ManagedIdentity",
                "type": "ManagedIdentityExtensionForWindows",
                "typeHandlerVersion": "1.0",
                "autoUpgradeMinorVersion": true,
                "settings": {
                  "port": 50343
                },
                "protectedSettings": {}
              }
            }
            {{end}}
          ]
        }
      }
    },
    "type": "Microsoft.Compute/virtualMachineScaleSets"
  }