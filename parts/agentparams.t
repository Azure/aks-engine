    "{{.Name}}Count": {
      "defaultValue": {{.Count}},
      "metadata": {
        "description": "The number of vms in agent pool {{.Name}}"
      },
      "type": "int"
    },
{{if .IsAvailabilitySets}}
    "{{.Name}}Offset": {
      "defaultValue": 0,
      "metadata": {
        "description": "offset to a particular vm within a VMAS agent pool"
      },
      "type": "int"
    },
{{end}}
    {{if .IsLowPriorityScaleSet}}
    "{{.Name}}ScaleSetPriority": {
      "allowedValues":[
        "Low",
        "Regular",
        ""
      ],
      "defaultValue": "{{.ScaleSetPriority}}",
      "metadata": {
        "description": "The priority for the VM Scale Set. This value can be Low or Regular."
      },
      "type": "string"
    },
    "{{.Name}}ScaleSetEvictionPolicy": {
      "allowedValues":[
        "Delete",
        "Deallocate",
        ""
      ],
      "defaultValue": "{{.ScaleSetEvictionPolicy}}",
      "metadata": {
        "description": "The Eviction Policy for a Low-priority VM Scale Set."
      },
      "type": "string"
    },
    {{end}}
{{if .IsSpotScaleSet}}
    "{{.Name}}ScaleSetPriority": {
      "allowedValues":[
        "Spot",
        "Regular",
        ""
      ],
      "defaultValue": "{{.ScaleSetPriority}}",
      "metadata": {
        "description": "The priority for the VM Scale Set. This value can be Spot or Regular."
      },
      "type": "string"
    },
    "{{.Name}}ScaleSetEvictionPolicy": {
      "allowedValues":[
        "Delete",
        "Deallocate",
        ""
      ],
      "defaultValue": "{{.ScaleSetEvictionPolicy}}",
      "metadata": {
        "description": "The Eviction Policy for a Spot VM Scale Set."
      },
      "type": "string"
    },
{{end}}
    "{{.Name}}VMSize": {
      {{GetKubernetesAllowedVMSKUs}}
      "defaultValue": "{{.VMSize}}",
      "metadata": {
        "description": "The size of the Virtual Machine."
      },
      "type": "string"
    },
{{if HasAvailabilityZones .}}
    "{{.Name}}AvailabilityZones": {
      "metadata": {
        "description": "Agent availability zones"
      },
      "type": "array"
    },
{{end}}
    "{{.Name}}osImageName": {
      "defaultValue": "",
      "metadata": {
        "description": "Name of a {{.OSType}} OS image. Needs to be used in conjuction with osImageResourceGroup."
      },
      "type": "string"
    },
    "{{.Name}}osImageResourceGroup": {
      "defaultValue": "",
      "metadata": {
        "description": "Resource group of a {{.OSType}} OS image. Needs to be used in conjuction with osImageName."
      },
      "type": "string"
    },
    "{{.Name}}osImageOffer": {
      "defaultValue": "UbuntuServer",
      "metadata": {
        "description": "{{.OSType}} OS image type."
      },
      "type": "string"
    },
    "{{.Name}}osImagePublisher": {
      "defaultValue": "Canonical",
      "metadata": {
        "description": "OS image publisher."
      },
      "type": "string"
    },
    "{{.Name}}osImageSKU": {
      "defaultValue": "16.04-LTS",
      "metadata": {
        "description": "OS image SKU."
      },
      "type": "string"
    },
    "{{.Name}}osImageVersion": {
      "defaultValue": "latest",
      "metadata": {
        "description": "OS image version."
      },
      "type": "string"
    },
{{if .IsCustomVNET}}
    "{{.Name}}VnetSubnetID": {
      "metadata": {
        "description": "Sets the vnet subnet of agent pool '{{.Name}}'."
      },
      "type": "string"
    }
{{else}}
    "{{.Name}}Subnet": {
      "defaultValue": "{{.Subnet}}",
      "metadata": {
        "description": "Sets the subnet of agent pool '{{.Name}}'."
      },
      "type": "string"
    }
{{end}}
{{if IsPublic .Ports}}
  ,"{{.Name}}EndpointDNSNamePrefix": {
      "metadata": {
        "description": "Sets the Domain name label for the agent pool IP Address.  The concatenation of the domain name label and the regional DNS zone make up the fully qualified domain name associated with the public IP address."
      },
      "type": "string"
    }
{{end}}
{{if HasPrivateRegistry}}
  ,"registry": {
      "metadata": {
        "description": "Private Container Registry"
      },
      "type": "string"
    },
  "registryKey": {
      "metadata": {
        "description": "base64 encoded key to the Private Container Registry"
      },
      "type": "string"
    }
  {{end}}
