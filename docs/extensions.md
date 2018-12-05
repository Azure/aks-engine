# Microsoft Azure Container Service Engine - Extensions

Extensions in aks-engine provide an easy way for aks-engine users to add pre-packaged functionality into their cluster.  For example, an extension could configure a monitoring solution on an AKS cluster.  The user would not need to know the details of how to install the monitoring solution.  Rather, the user would simply add the extension into the extensionProfiles section of the template.

# extensionProfiles

The extensionProfiles contains the extensions that the cluster will install. The following illustrates a template with a hello-world-dcos extension.

``` javascript
{
  ...
  "extensionProfiles": [
    {
        "name": "hello-world-dcos",
        "version": "v1",
        "extensionParameters": "parameters",
        "rootURL": "http://mytestlocation.com/hello-world-dcos/",
        "script": "hello-world-dcos.sh"
    }
  ]
}
```

|Name|Required|Description|
|---|---|---|
|name|yes|the name of the extension.  This has to exactly match the name of a folder under the extensions folder|
|version|yes|the version of the extension.  This has to exactly match the name of the folder under the extension name folder|
|extensionParameters|optional|extension parameters may be required by extensions.  The format of the parameters is also extension dependant.|
|rootURL|optional|url to the root location of extensions.  The rootURL must have an extensions child folder that follows the extensions convention.  The rootURL is mainly used for testing purposes.|
|script|optional|Used for preprovision scripts this points to the location of the script to run inside of the extension folder.|

# rootURL
You normally would not provide a rootURL.  The extensions are normally loaded from the extensions folder in GitHub.  However, you may specify the rootURL when testing a new extension.  The rootURL must adhere to the extensions conventions.  For example, in order to use an Azure Storage account to test an extension named extension-one, you would do the following:
- Create a storage account.  For the purposes of this example, we will call it 'mystorageaccount'
- Create a blob container called 'extensions'
- Under 'extensions', create a folder called 'extension-one'
- Under 'extension-one', create a folder called 'v1'
- Under 'v1', upload your files (see Required Extension Files)
- Set the rootURL to: 'https://mystorageaccount.blob.core.windows.net/'

# masterProfile
Extensions, in the current implementation run a script on a master node. The extensions array in the masterProfile define that the master pool will have the script run on a single node on it. If you want it to run on all pass in All to singleOrAll

``` javascript
{
  "masterProfile": {
      "count": 3,
      "dnsPrefix": "dnsprefix",
      "vmSize": "Standard_D2_v2",
      "osType": "Linux",
      "firstConsecutiveStaticIP": "10.240.255.5",
      "extensions": [
        {
          "name": "hello-world-k8s",
          "singleOrAll": "single"
        }
     ]
  },
  "extensionProfiles": [
    {
        "name": "hello-world-k8s",
        "version": "v1",
        "extensionParameters": "parameters"
    }
  ]
}
```
Or they can be referenced as a preprovision extension, this will run during cloud init before the cluster is brought up. Usually used for installing antivirus or the like. These will run on all the masters or the nodes in the agent pool it is specified to run on if it is specified. Single or all is a formality at this point.
``` javascript
{
  "masterProfile": {
      "count": 3,
      "dnsPrefix": "dnsprefix",
      "vmSize": "Standard_D2_v2",
      "osType": "Linux",
      "firstConsecutiveStaticIP": "10.240.255.5",
      "preProvisionExtension": {
          "name": "hello-world",
          "singleOrAll": "All"
      }

  },
  "extensionProfiles": [
    {
        "name": "hello-world-k8s",
        "version": "v1",
        "extensionParameters": "parameters",
        "script": "hello.sh"
    }
  ]
}
```
|Name|Required|Description|
|---|---|---|
|name|yes|The name of the extension. This must match the name in the extensionProfiles|

# Required Extension Files

In order to install a post provision extension, there are four required files - supported-orchestrators.json, template.json, template-link.json and EXTENSION-NAME.sh. Following is a description of each file.

In order to install a preprovision extension, there are two required files - supported-orchestrators.json and EXTENSION-NAME.sh. Following is a description of each file.


|File Name|Description|
|-----------------------------|---|
|supported-orchestrators.json |Defines what orchestrators are supported by the extension (Swarm, Dcos, or Kubernetes)|
|template.json               |The ARM template used to deploy the extension|
|template-link.json          |The ARM template snippet which will be injected into azuredeploy.json to call template.json|
|EXTENSION-NAME.sh           |The script file that will execute on the VM itself via Custom Script Extension to perform installation of the extension|

# Creating supported-orchestrators.json

The supported-orchestrators.json file is a simple one line file that contains the list of supported orchestrators for which the extension can be installed into.

``` javascript
["Kubernetes"]
```

# Creating extension template.json

The template.json file is a linked template that will be called by the main cluster deployment template and must adhere to all the rules of a normal ARM template. All the necessary parameters needed from the azuredeploy.json file must be passed into this template and defined appropriately.

Additional variables can be defined for use in creating additional resources. Additional resources can also be created.  The key resource for installing the extension is the custom script extension.

Modify the commandToExecute entry with the necessary command and paramters to install the desired extension. Replace EXTENSION-NAME with the name of the extension. The resource name of the custom script extension has to have the same name as the other custom script on the box as we aren't allowed to have two, this is also why we use a linked deployment so we can have the same resource twice and just make this one depend on the other so that it always runs after the provision extension is done.

The following is an example of the template.json file.

``` javascript
{
   "$schema": "http://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
   "contentVersion": "1.0.0.0",
   "parameters": {
		"apiVersionStorage": {
			"type": "string",
			"minLength": 1,
			"metadata": {
				"description": "Storage API Version"
			}
		},
		"apiVersionCompute": {
			"type": "string",
			"minLength": 1,
			"metadata": {
				"description": "Compute API Version"
			}
		},
		"username": {
			"type": "string",
			"minLength": 1,
			"metadata": {
				"description": "Username for OS"
			}
		},
		"storageAccountBaseName": {
			"type": "string",
			"minLength": 1,
			"metadata": {
				"description": "Base Name of Storage Account"
			}
		},
		"extensionParameters": {
			"type": "securestring",
			"minLength": 1,
			"metadata": {
				"description": "Custom Parameter for Extension"
			}
		}
   },
   "variables": {
		"singleQuote": "'",
		"sampleStorageAccountName": "[concat(uniqueString(concat(parameters('storageAccountBaseName'), 'sample')), 'aa')]"
		"initScriptUrl": "https://raw.githubusercontent.com/Azure/aks-engine/master/extensions/EXTENSION-NAME/v1/EXTENSION-NAME.sh"
   },
   "resources": [
	{
      "apiVersion": "[parameters('apiVersionStorage')]",
      "dependsOn": [],
      "location": "[resourceGroup().location]",
      "name": "[variables('sampleStorageAccountName')]",
      "properties": {
        "accountType": "Standard_LRS"
      },
      "type": "Microsoft.Storage/storageAccounts"
	}, {
      "apiVersion": "[parameters('apiVersionCompute')]",
      "dependsOn": [],
      "location": "[resourceGroup().location]",
      "type": "Microsoft.Compute/virtualMachines/extensions",
	  "name": "CustomExtension",
      "properties": {
        "publisher": "Microsoft.OSTCExtensions",
        "type": "CustomScriptForLinux",
        "typeHandlerVersion": "1.5",
        "autoUpgradeMinorVersion": true,
        "settings": {
			"fileUris": [
			   "[variables('initScriptUrl')]"
			 ]
		},
        "protectedSettings": {
			"commandToExecute": "[concat('/bin/bash -c \"/bin/bash ./EXTENSION-NAME.sh ', variables('singleQuote'), parameters('extensionParameters'), variables('singleQuote'), ' ', variables('singleQuote'), parameters('sampleStorageAccountName'), variables('singleQuote'), ' >> /var/log/azure/sysdig-provision.log 2>&1 &\" &')]"
        }
      }
    }
	],
   "outputs": {  }
}
```

# Creating extension template-link.json

When aks-engine generates the azuredeploy.json file, this JSON snippet will be injected. This code calls the linked template (template.json) defined above.

Any parameters from the main azuredeploy.json file that is needed by template.json must be passed in via the parameters section. The parameter, "extensionParameters" is an optional parameter that is passed in directly by the user in the **extensionProfiles** section as defined in an earlier section. This special parameter can be used to pass in information such as an activation key or access code (as an example). If the extension does not need this capability, this optional parameter can be deleted.

Before this resource is created, all the dependencies must be satisfied first as defined by "dependsOn". The default dependency is that the entire cluster is fully provisioned before the script extension executes. This can be changed to meet your needs.

Replace "**EXTENSION-NAME**" with the name of the extension.

``` javascript
{
    "name": "EXTENSION-NAME",
    "type": "Microsoft.Resources/deployments",
    "apiVersion": "[variables('apiVersionCompute')]",
    "dependsOn": [
        "vmLoopNode"
    ],
    "properties": {
        "mode": "Incremental",
        "templateLink": {
            "uri": "https://raw.githubusercontent.com/Azure/aks-engine/master/extensions/EXTENSION-NAME/v1/template.json",
            "contentVersion": "1.0.0.0"
        },
        "parameters": {
            "apiVersionCompute": {
                "value": "[variables('apiVersionCompute')]"
            },
            "username": {
                "value": "[parameters('linuxAdminUsername')]"
            },
            "storageAccountBaseName": {
                "value": "[variables('storageAccountBaseName')]"
            },
            "extensionParameters": {
                "value": "EXTENSION_PARAMETERS_REPLACE"
            }
        }
    }
}
```

# Creating extension script file

The script file will get executed on the VM to install the extension. Following is an example of a script.sh file. For a preprovision extension the relative path of the file inside the version folder needs to be passed in as the "script" property in the extensions profile

``` bash
#!/bin/bash

# Add comments to explain the components of the script for easier troubleshooting
# Include echo statements so comments are written to output file
# Include necessary error checking

# Local variables

VARIABLE1=$1
VARIABLE2=$2
VARIABLE3=$3

echo $(date) " - Starting Script"

# Step 1 - example of creating config file
echo $(date) " - Creating sample.yaml file using local variables"

cat > sample.yaml <<EOF
line 1 $VARIABLE1
line 2 $VARIABLE2
line 3 $VARIABLE3
EOF

# Step 2 - example of downloading file
echo $(date) " - Downloading file"

curl -sSL http://example.com/sample/install

# Step 3 - example of executing other commands
echo $(date) " - Executing command"

install sample.yaml

echo $(date) " - Script complete"
```

# Current list of extensions
- [hello-world-dcos](../extensions/hello-world-dcos/README.md)
- [hello-world-k8s](../extensions/hello-world-k8s/README.md)

# Known issues
Kubernetes extensions that run after provisioning don't currently work if the VM needs to reboot for security reboots. this is a timing issue. the extension script is started before the vm reboots and it will be cutoff before it finishes but will still report success. I've tried to get the provision script to only finish as reboot happens and I haven't gotten that to work. An extension could work most of the time if it cancelled the restart at the start and checked if a restart was needed and scheduled one at the end of its work
