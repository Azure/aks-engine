{{if IsPublic .Ports}}
  {{ if not IsKubernetes }}
    "{{.Name}}FQDN": {
        "type": "string",
        "value": "[reference(concat('Microsoft.Network/publicIPAddresses/', variables('{{.Name}}IPAddressName'))).dnsSettings.fqdn]"
    },
  {{end}}
{{end}}
{{if and .IsAvailabilitySets .IsStorageAccount}}
  "{{.Name}}StorageAccountOffset": {
      "type": "int",
      "value": "[variables('{{.Name}}StorageAccountOffset')]"
    },
    "{{.Name}}StorageAccountCount": {
      "type": "int",
      "value": "[variables('{{.Name}}StorageAccountsCount')]"
    },
    "{{.Name}}SubnetName": {
      "type": "string",
      "value": "[variables('{{.Name}}SubnetName')]"
    },
{{end}}