{{if and UseManagedIdentity (not UserAssignedIDEnabled)}}
  {
    "apiVersion": "[variables('apiVersionAuthorizationSystem')]",
    "name": "[guid(concat('Microsoft.Compute/virtualMachineScaleSets/', variables('{{.Name}}VMNamePrefix'), 'vmidentity'))]",
    "type": "Microsoft.Authorization/roleAssignments",
    "properties": {
      "roleDefinitionId": "[variables('readerRoleDefinitionId')]",
      "principalId": "[reference(concat('Microsoft.Compute/virtualMachineScaleSets/', variables('{{.Name}}VMNamePrefix')), '2017-03-30', 'Full').identity.principalId]"
    },
    "dependsOn": [
      "[concat('Microsoft.Compute/virtualMachineScaleSets/', variables('{{.Name}}VMNamePrefix'))]"
    ]
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
      "aksEngineVersion" : "[parameters('aksEngineVersion')]",
      "poolName" : "{{.Name}}"
    },
    "location": "[variables('location')]",
    {{ if HasAvailabilityZones .}}
    "zones": "[parameters('{{.Name}}AvailabilityZones')]",
    {{ end }}
    "name": "[variables('{{.Name}}VMNamePrefix')]",
    {{if UseManagedIdentity}}
      {{if UserAssignedIDEnabled}}
    "identity": {
      "type": "userAssigned",
      "userAssignedIdentities": {
        "[variables('userAssignedIDReference')]":{}
      }
    },
      {{else}}
    "identity": {
      "type": "systemAssigned"
    },
      {{end}}
    {{end}}
    "sku": {
      "tier": "Standard",
      "capacity": "[variables('{{.Name}}Count')]",
      "name": "[variables('{{.Name}}VMSize')]"
    },
    "properties": {
      "singlePlacementGroup": {{UseSinglePlacementGroup .}},
      "overprovision": {{IsVMSSOverProvisioningEnabled}},
      {{if IsVMSSOverProvisioningEnabled}}
      "doNotRunExtensionsOnOverprovisionedVMs": true,
      {{end}}
      "upgradePolicy": {
        "mode": "Manual"
      },
      "virtualMachineProfile": {
        {{if .IsLowPriorityScaleSet}}
          "priority": "[variables('{{.Name}}ScaleSetPriority')]",
          "evictionPolicy": "[variables('{{.Name}}ScaleSetEvictionPolicy')]",
        {{end}}
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
          "adminPassword": "[parameters('windowsAdminPassword')]",
          "windowsConfiguration": {
            "enableAutomaticUpdates": {{WindowsAutomaticUpdateEnabled}}
          }
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
              "name": "[concat(variables('{{.Name}}VMNamePrefix'), '-computeAksWindowsBilling')]",
              "properties": {
                "publisher": "Microsoft.AKS",
                "type": "Compute.AKS-Engine.Windows.Billing",
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
