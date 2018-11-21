{{if and UseManagedIdentity (not UserAssignedIDEnabled)}}
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
      "resourceNameSuffix" : "[parameters('nameSuffix')]",
      "orchestrator" : "[variables('orchestratorNameVersionTag')]",
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
      "overprovision": false,
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
                "enableAcceleratedNetworking" : "{{.AcceleratedNetworkingEnabled}}",
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
          "computerNamePrefix": "[variables('{{.Name}}VMNamePrefix')]",
          {{GetKubernetesAgentCustomData .}}
          "linuxConfiguration": {
              "disablePasswordAuthentication": "true",
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
          {{if not (UseAgentCustomImage .)}}
            {{GetDataDisks .}}
          {{end}}
          "imageReference": {
            {{if UseAgentCustomImage .}}
            "id": "[resourceId(variables('{{.Name}}osImageResourceGroup'), 'Microsoft.Compute/images', variables('{{.Name}}osImageName'))]"
            {{else}}
            "offer": "[variables('{{.Name}}osImageOffer')]",
            "publisher": "[variables('{{.Name}}osImagePublisher')]",
            "sku": "[variables('{{.Name}}osImageSKU')]",
            "version": "[variables('{{.Name}}osImageVersion')]"
            {{end}}
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
                "publisher": "Microsoft.Azure.Extensions",
                "type": "CustomScript",
                "typeHandlerVersion": "2.0",
                "autoUpgradeMinorVersion": true,
                "settings": {},
                "protectedSettings": {
                  "commandToExecute": "[concat('retrycmd_if_failure() { r=$1; w=$2; t=$3; shift && shift && shift; for i in $(seq 1 $r); do timeout $t ${@}; [ $? -eq 0  ] && break || if [ $i -eq $r ]; then return 1; else sleep $w; fi; done };{{if not (IsFeatureEnabled "BlockOutboundInternet")}} ERR_OUTBOUND_CONN_FAIL=50; retrycmd_if_failure 150 1 3 nc -vz {{if IsMooncake}}gcr.azk8s.cn 80{{else}}k8s.gcr.io 443 && nc -vz gcr.io 443 && nc -vz docker.io 443{{end}} || exit $ERR_OUTBOUND_CONN_FAIL;{{end}} for i in $(seq 1 1200); do if [ -f /opt/azure/containers/provision.sh ]; then break; fi; if [ $i -eq 1200 ]; then exit 100; else sleep 1; fi; done; ', variables('provisionScriptParametersCommon'),' GPU_NODE={{IsNSeriesSKU .}} /usr/bin/nohup /bin/bash -c \"/bin/bash /opt/azure/containers/provision.sh >> /var/log/azure/cluster-provision.log 2>&1{{if IsFeatureEnabled "CSERunInBackground" }} &{{end}}\"')]"
                }
              }
            }
            {{if UseAksExtension}}
            ,{
              "name": "[concat(variables('{{.Name}}VMNamePrefix'), '-computeAksLinuxBilling')]",
              "location": "[variables('location')]",
              "properties": {
                "publisher": "Microsoft.AKS",
                "type": "Compute.AKS-Engine.Linux.Billing",
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
                "type": "ManagedIdentityExtensionForLinux",
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
