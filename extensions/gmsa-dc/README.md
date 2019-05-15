# gMSA-DC Extension

This extension will create an Active Directory Forest and Domain using one of the agents nodes.  It will also create a Group Managed Service Account (gMSA) 
and create the needed YAML file for creation of the gMSA credential spec resource.  We also enable ssh on the node for easier retrieval of the credspec file.

It is required to make a separate pool of 1 member to house the DC for gMSA testing.

# Configuration

|Name               |Required|Acceptable Value     |
|-------------------|--------|---------------------|
|name               |yes     |gmsa-dc              |
|version            |yes     |v1                   |
|rootURL            |optional|                     |

# Example

```
    ...
    "agentPoolProfiles": [
      {
        "name": "windowspool1",
        "count": 2
		...
      },
	  {
        "name": "windowsgmsa",
		"count": 1
        "extensions": [
          {
            "name": "gmsa-dc"
          }
        ]
      }
    ],
    ...
    "extensionProfiles": [
      {
        "name": "gmsa-dc",
        "version": "v1"
      }
    ]
    ...
```


# Supported Orchestrators

Kubernetes

# Troubleshoot

The three scripts that run use PowerShell transcripts.  They log to the following directory
```sh
C:\gmsa
```

Extension execution output is logged to files found under the following directory on the target virtual machine.

```sh
C:\WindowsAzure\Logs\Plugins\Microsoft.Compute.CustomScriptExtension
```

The specified files are downloaded into the following directory on the target virtual machine.

```sh
C:\Packages\Plugins\Microsoft.Compute.CustomScriptExtension\1.*\Downloads\<n>
```
