{{if .IsStorageAccount}}
    "{{.Name}}StorageAccountOffset": "[mul(variables('maxStorageAccountsPerAgent'),variables('{{.Name}}Index'))]",
    "{{.Name}}StorageAccountsCount": "[add(div(variables('{{.Name}}Count'), variables('maxVMsPerStorageAccount')), mod(add(mod(variables('{{.Name}}Count'), variables('maxVMsPerStorageAccount')),2), add(mod(variables('{{.Name}}Count'), variables('maxVMsPerStorageAccount')),1)))]",
{{end}}
    "{{.Name}}Count": "[parameters('{{.Name}}Count')]",
    "{{.Name}}VMNamePrefix": "{{GetAgentVMPrefix .}}",
{{if .IsWindows}}
    "winResourceNamePrefix" : "[substring(parameters('nameSuffix'), 0, 5)]",
{{end}}
{{if .IsAvailabilitySets}}
    "{{.Name}}Offset": "[parameters('{{.Name}}Offset')]",
    "{{.Name}}AvailabilitySet": "[concat('{{.Name}}-availabilitySet-', parameters('nameSuffix'))]",
{{else}}
    {{if .IsLowPriorityScaleSet}}
    "{{.Name}}ScaleSetPriority": "[parameters('{{.Name}}ScaleSetPriority')]",
    "{{.Name}}ScaleSetEvictionPolicy": "[parameters('{{.Name}}ScaleSetEvictionPolicy')]",
    {{end}}
{{end}}
    "{{.Name}}VMSize": "[parameters('{{.Name}}VMSize')]",
{{if .IsCustomVNET}}
    "{{.Name}}VnetSubnetID": "[parameters('{{.Name}}VnetSubnetID')]",
    "{{.Name}}SubnetName": "[parameters('{{.Name}}VnetSubnetID')]",
    "{{.Name}}VnetParts": "[split(parameters('{{.Name}}VnetSubnetID'),'/subnets/')]",
{{else}}
    "{{.Name}}VnetSubnetID": "[variables('vnetSubnetID')]",
    "{{.Name}}SubnetName": "[variables('subnetName')]",
{{end}}
    "{{.Name}}osImageOffer": "[parameters('{{.Name}}osImageOffer')]",
    "{{.Name}}osImageSKU": "[parameters('{{.Name}}osImageSKU')]",
    "{{.Name}}osImagePublisher": "[parameters('{{.Name}}osImagePublisher')]",
    "{{.Name}}osImageVersion": "[parameters('{{.Name}}osImageVersion')]",
    "{{.Name}}osImageName": "[parameters('{{.Name}}osImageName')]",
    "{{.Name}}osImageResourceGroup": "[parameters('{{.Name}}osImageResourceGroup')]",
