    "adminUsername": "[parameters('linuxAdminUsername')]",
    "maxVMsPerPool": 100,
    "apiVersionDefault": "2016-03-30",
    "configureClusterScriptFile": "configure-swarm-cluster.sh",
{{if .MasterProfile.IsRHEL}}
    "agentCustomScript": "[concat('/usr/bin/nohup /bin/bash -c \"/bin/bash ',variables('configureClusterScriptFile'), ' ',variables('clusterInstallParameters'),' >> /var/log/azure/cluster-bootstrap.log 2>&1 &\" &')]",
{{else}}
    "agentCustomScript": "[concat('/usr/bin/nohup /bin/bash -c \"/bin/bash /opt/azure/containers/',variables('configureClusterScriptFile'), ' ',variables('clusterInstallParameters'),' >> /var/log/azure/cluster-bootstrap.log 2>&1 &\" &')]",
{{end}}
    "agentMaxVMs": 100,
    "clusterInstallParameters": "[concat(variables('orchestratorVersion'), ' ',variables('dockerComposeVersion'), ' ',variables('masterCount'), ' ',variables('masterVMNamePrefix'), ' ',variables('masterFirstAddrOctet4'), ' ',variables('adminUsername'),' ',variables('postInstallScriptURI'),' ',variables('masterFirstAddrPrefix'),' ', parameters('dockerEngineDownloadRepo'), ' ', parameters('dockerComposeDownloadURL'))]",
{{if .LinuxProfile.HasSecrets}}
    "linuxProfileSecrets" :
      [
          {{range  $vIndex, $vault := .LinuxProfile.Secrets}}
            {{if $vIndex}} , {{end}}
              {
                "sourceVault":{
                  "id":"[parameters('linuxKeyVaultID{{$vIndex}}')]"
                },
                "vaultCertificates":[
                {{range $cIndex, $cert := $vault.VaultCertificates}}
                  {{if $cIndex}} , {{end}}
                  {
                    "certificateUrl" :"[parameters('linuxKeyVaultID{{$vIndex}}CertificateURL{{$cIndex}}')]"
                  }
                {{end}}
                ]
              }
        {{end}}
      ],
{{end}}
    "masterAvailabilitySet": "[concat(variables('orchestratorName'), '-master-availabilitySet-', variables('nameSuffix'))]",
    "masterCount": {{.MasterProfile.Count}},
{{if .MasterProfile.IsRHEL}}
    "masterCustomScript": "[concat('/bin/bash -c \"/bin/bash ',variables('configureClusterScriptFile'), ' ',variables('clusterInstallParameters'),' >> /var/log/azure/cluster-bootstrap.log 2>&1\"')]",
{{else}}
    "masterCustomScript": "[concat('/bin/bash -c \"/bin/bash /opt/azure/containers/',variables('configureClusterScriptFile'), ' ',variables('clusterInstallParameters'),' >> /var/log/azure/cluster-bootstrap.log 2>&1\"')]",
{{end}}
    "masterEndpointDNSNamePrefix": "[tolower(parameters('masterEndpointDNSNamePrefix'))]",
    "masterLbBackendPoolName": "[concat(variables('orchestratorName'), '-master-pool-', variables('nameSuffix'))]",
    "masterLbID": "[resourceId('Microsoft.Network/loadBalancers',variables('masterLbName'))]",
    "masterLbIPConfigID": "[concat(variables('masterLbID'),'/frontendIPConfigurations/', variables('masterLbIPConfigName'))]",
    "masterLbIPConfigName": "[concat(variables('orchestratorName'), '-master-lbFrontEnd-', variables('nameSuffix'))]",
    "masterLbName": "[concat(variables('orchestratorName'), '-master-lb-', variables('nameSuffix'))]",
    "masterPublicIPAddressName": "[concat(variables('orchestratorName'), '-master-ip-', variables('masterEndpointDNSNamePrefix'), '-', variables('nameSuffix'))]",
{{if .MasterProfile.IsCustomVNET}}
    "masterVnetSubnetID": "[parameters('masterVnetSubnetID')]",
{{else}}
    "masterSubnet": "[parameters('masterSubnet')]",
    "masterSubnetName": "[concat(variables('orchestratorName'), '-masterSubnet')]",
    "vnetID": "[resourceId('Microsoft.Network/virtualNetworks',variables('virtualNetworkName'))]",
    "masterVnetSubnetID": "[concat(variables('vnetID'),'/subnets/',variables('masterSubnetName'))]",
    "virtualNetworkName": "[concat(variables('orchestratorName'), '-vnet-', variables('nameSuffix'))]",
{{end}}
    "masterFirstAddrOctets": "[split(parameters('firstConsecutiveStaticIP'),'.')]",
    "masterFirstAddrOctet4": "[variables('masterFirstAddrOctets')[3]]",
    "masterFirstAddrPrefix": "[concat(variables('masterFirstAddrOctets')[0],'.',variables('masterFirstAddrOctets')[1],'.',variables('masterFirstAddrOctets')[2],'.')]",
    "masterVMNamePrefix": "[concat(variables('orchestratorName'), '-master-', variables('nameSuffix'), '-')]",
    "masterVMSize": "[parameters('masterVMSize')]",
    "nameSuffix": "[parameters('nameSuffix')]",
    "masterSshInboundNatRuleIdPrefix": "[concat(variables('masterLbID'),'/inboundNatRules/SSH-',variables('masterVMNamePrefix'))]",
    "masterSshPort22InboundNatRuleNamePrefix": "[concat(variables('masterLbName'),'/SSHPort22-',variables('masterVMNamePrefix'))]",
    "masterSshPort22InboundNatRuleIdPrefix": "[concat(variables('masterLbID'),'/inboundNatRules/SSHPort22-',variables('masterVMNamePrefix'))]",
     "masterLbInboundNatRules":[
      [
        {
          "id": "[concat(variables('masterSshInboundNatRuleIdPrefix'),'0')]"
        },
        {
          "id": "[concat(variables('masterSshPort22InboundNatRuleIdPrefix'),'0')]"
        }
      ],
      [
        {
          "id": "[concat(variables('masterSshInboundNatRuleIdPrefix'),'1')]"
        }
      ],
      [
        {
          "id": "[concat(variables('masterSshInboundNatRuleIdPrefix'),'2')]"
        }
      ],
      [
        {
          "id": "[concat(variables('masterSshInboundNatRuleIdPrefix'),'3')]"
        }
      ],
      [
        {
          "id": "[concat(variables('masterSshInboundNatRuleIdPrefix'),'4')]"
        }
      ]
    ],
    "orchestratorName": "swarm",
    "osImageOffer": "[parameters('osImageOffer')]",
    "osImagePublisher": "[parameters('osImagePublisher')]",
    "osImageSKU": "14.04.5-LTS",
    "osImageVersion": "14.04.201706190",
    {{getSwarmVersions}}
    "locations": [
         "[resourceGroup().location]",
         "[parameters('location')]"
    ],
    "location": "[variables('locations')[mod(add(2,length(parameters('location'))),add(1,length(parameters('location'))))]]",
    "postInstallScriptURI": "disabled",
    "sshKeyPath": "[concat('/home/', variables('adminUsername'), '/.ssh/authorized_keys')]",
{{if .HasStorageAccountDisks}}
    "apiVersionStorage": "2015-06-15",
    "maxVMsPerStorageAccount": 20,
    "maxStorageAccountsPerAgent": "[div(variables('maxVMsPerPool'),variables('maxVMsPerStorageAccount'))]",
    "dataStorageAccountPrefixSeed": 97,
    "storageAccountPrefixes": [ "0", "6", "c", "i", "o", "u", "1", "7", "d", "j", "p", "v", "2", "8", "e", "k", "q", "w", "3", "9", "f", "l", "r", "x", "4", "a", "g", "m", "s", "y", "5", "b", "h", "n", "t", "z" ],
    "storageAccountPrefixesCount": "[length(variables('storageAccountPrefixes'))]",
    "vmsPerStorageAccount": 20,
    "storageAccountBaseName": "[uniqueString(concat(variables('masterEndpointDNSNamePrefix'),variables('location')))]",
    {{GetSizeMap}},
{{else}}
    "storageAccountPrefixes": [],
    "storageAccountBaseName": "",
{{end}}
{{if .HasManagedDisks}}
    "apiVersionStorageManagedDisks": "2016-04-30-preview",
{{end}}
{{if .MasterProfile.IsStorageAccount}}
    "masterStorageAccountName": "[concat(variables('storageAccountBaseName'), '0')]",
{{end}}
    "sshRSAPublicKey": "[parameters('sshRSAPublicKey')]"
{{if .HasWindows}}
    ,"windowsAdminUsername": "[parameters('windowsAdminUsername')]",
    "windowsAdminPassword": "[parameters('windowsAdminPassword')]",
    "agentWindowsPublisher": "[parameters('agentWindowsPublisher')]",
    "agentWindowsOffer": "[parameters('agentWindowsOffer')]",
    "agentWindowsSku": "[parameters('agentWindowsSku')]",
    "agentWindowsVersion": "[parameters('agentWindowsVersion')]",
    "singleQuote": "'",
    "windowsCustomScriptArguments": "[concat('$arguments = ', variables('singleQuote'),'-SwarmMasterIP ', variables('masterFirstAddrPrefix'), variables('masterFirstAddrOctet4'), variables('singleQuote'), ' ; ')]",
    "windowsCustomScriptSuffix": " $inputFile = '%SYSTEMDRIVE%\\AzureData\\CustomData.bin' ; $outputFile = '%SYSTEMDRIVE%\\AzureData\\CustomDataSetupScript.ps1' ; $inputStream = New-Object System.IO.FileStream $inputFile, ([IO.FileMode]::Open), ([IO.FileAccess]::Read), ([IO.FileShare]::Read) ; $sr = New-Object System.IO.StreamReader(New-Object System.IO.Compression.GZipStream($inputStream, [System.IO.Compression.CompressionMode]::Decompress)) ; $sr.ReadToEnd() | Out-File($outputFile) ; Invoke-Expression('{0} {1}' -f $outputFile, $arguments) ; ",
    "windowsCustomScript": "[concat('powershell.exe -ExecutionPolicy Unrestricted -command \"', variables('windowsCustomScriptArguments'), variables('windowsCustomScriptSuffix'), '\" > %SYSTEMDRIVE%\\AzureData\\CustomDataSetupScript.log 2>&1')]",
    "agentWindowsBackendPort": 3389
    {{if .WindowsProfile.HasSecrets}}
    ,
    "windowsProfileSecrets" :
      [
          {{range  $vIndex, $vault := .LinuxProfile.Secrets}}
            {{if $vIndex}} , {{end}}
              {
                "sourceVault":{
                  "id":"[parameters('windowsKeyVaultID{{$vIndex}}')]"
                },
                "vaultCertificates":[
                {{range $cIndex, $cert := $vault.VaultCertificates}}
                  {{if $cIndex}} , {{end}}
                  {
                    "certificateUrl" :"[parameters('windowsKeyVaultID{{$vIndex}}CertificateURL{{$cIndex}}')]",
                    "certificateStore" :"[parameters('windowsKeyVaultID{{$vIndex}}CertificateStore{{$cIndex}}')]"
                  }
                {{end}}
                ]
              }
        {{end}}
      ]
      {{end}}
{{end}}
