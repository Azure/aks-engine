{{if not .IsRHEL}}
    "{{.Name}}RunCmd": "[concat('runcmd:\n {{GetSwarmAgentPreprovisionExtensionCommands .}} \n-  [ /bin/bash, /opt/azure/containers/install-cluster.sh ]\n\n')]",
    "{{.Name}}RunCmdFile": "[concat(' -  content: |\n        #!/bin/bash\n        ','sudo mkdir -p /var/log/azure\n        ',variables('agentCustomScript'),'\n    path: /opt/azure/containers/install-cluster.sh\n    permissions: \"0744\"\n')]",
{{end}}
    "{{.Name}}OSImageOffer": "[variables('osImageOffer')]",
    "{{.Name}}OSImagePublisher": "[variables('osImagePublisher')]",
    "{{.Name}}OSImageSKU": "[variables('osImageSKU')]",
    "{{.Name}}OSImageVersion": "[variables('osImageVersion')]",
    "{{.Name}}Count": "[parameters('{{.Name}}Count')]",
    "{{.Name}}VMNamePrefix": "[concat(variables('orchestratorName'), '-{{.Name}}-', variables('nameSuffix'))]",
    "{{.Name}}VMSize": "[parameters('{{.Name}}VMSize')]",
    "{{.Name}}VMSizeTier": "[split(parameters('{{.Name}}VMSize'),'_')[0]]",
{{if .IsAvailabilitySets}}
    {{if .IsStorageAccount}}
    "{{.Name}}StorageAccountsCount": "[add(div(variables('{{.Name}}Count'), variables('maxVMsPerStorageAccount')), mod(add(mod(variables('{{.Name}}Count'), variables('maxVMsPerStorageAccount')),2), add(mod(variables('{{.Name}}Count'), variables('maxVMsPerStorageAccount')),1)))]",
    "{{.Name}}StorageAccountOffset": "[mul(variables('maxStorageAccountsPerAgent'),variables('{{.Name}}Index'))]",
    {{end}}
    "{{.Name}}AvailabilitySet": "[concat('{{.Name}}-availabilitySet-', variables('nameSuffix'))]",
    "{{.Name}}Offset": "[parameters('{{.Name}}Offset')]",
{{else}}
    {{if .IsStorageAccount}}
    "{{.Name}}StorageAccountsCount": "[variables('maxStorageAccountsPerAgent')]",
    {{end}}
{{end}}
{{if .IsCustomVNET}}
    "{{.Name}}VnetSubnetID": "[parameters('{{.Name}}VnetSubnetID')]",
{{else}}
    "{{.Name}}Subnet": "[parameters('{{.Name}}Subnet')]",
    "{{.Name}}SubnetName": "[concat(variables('orchestratorName'), '-{{.Name}}subnet')]",
    "{{.Name}}VnetSubnetID": "[concat(variables('vnetID'),'/subnets/',variables('{{.Name}}SubnetName'))]",
{{end}}
{{if IsPublic .Ports}}
    "{{.Name}}EndpointDNSNamePrefix": "[tolower(parameters('{{.Name}}EndpointDNSNamePrefix'))]",
    "{{.Name}}IPAddressName": "[concat(variables('orchestratorName'), '-agent-ip-', variables('{{.Name}}EndpointDNSNamePrefix'), '-', variables('nameSuffix'))]",
    "{{.Name}}LbBackendPoolName": "[concat(variables('orchestratorName'), '-{{.Name}}-', variables('nameSuffix'))]",
    "{{.Name}}LbID": "[resourceId('Microsoft.Network/loadBalancers',variables('{{.Name}}LbName'))]",
    "{{.Name}}LbIPConfigID": "[concat(variables('{{.Name}}LbID'),'/frontendIPConfigurations/', variables('{{.Name}}LbIPConfigName'))]",
    "{{.Name}}LbIPConfigName": "[concat(variables('orchestratorName'), '-{{.Name}}-', variables('nameSuffix'))]",
    "{{.Name}}LbName": "[concat(variables('orchestratorName'), '-{{.Name}}-', variables('nameSuffix'))]",
     {{if .IsWindows}}
        "{{.Name}}WindowsRDPNatRangeStart": 3389,
        "{{.Name}}WindowsRDPEndRangeStop": "[add(variables('{{.Name}}WindowsRDPNatRangeStart'), add(variables('{{.Name}}Count'),variables('{{.Name}}Count')))]",
    {{end}}
 {{end}}
