# Windows Patching Extension

This extension will install Windows Server patches, including prerelease hotfixes. It's useful for the following cases:

1. Microsoft support has provided a pre-release hotfix for your testing
2. Installing additional Windows update packages (MSU) that are not yet included in the default Windows Server with Containers VM on the Azure Marketplace.

## Configuration

|Name               |Required| Acceptable Value     |
|-------------------|--------|----------------------|
|name               |yes     | windows-patches      |
|version            |yes     | v1                   |
|rootURL            |optional| `https://raw.githubusercontent.com/Azure/aks-engine/master/` or any repo with the same extensions/... directory structure |
|extensionParameters|yes     | comma-delimited list of URIs enclosed with ' such as `'https://privateupdates.domain.ext/Windows10.0-KB999999-x64-InstallForTestingPurposesOnly.exe', 'https://privateupdates.domain.ext/Windows10.0-KB123456-x64-InstallForTestingPurposesOnly.exe'` |

## Example

```json
    ...
    "agentPoolProfiles": [
      {
        "name": "windowspool1",
        "extensions": [
          {
            "name": "windows-patches"
          }
        ]
      }
    ],
    ...
    "extensionProfiles": [
      {
        "name": "windows-patches",
        "version": "v1",
        "rootURL": "https://raw.githubusercontent.com/Azure/aks-engine/master/",
        "extensionParameters": "'https://mypatches.blob.core.windows.net/hotfix3692/Windows10.0-KB999999-x64-InstallForTestingPurposesOnly.exe?sp=r&st=2018-08-17T00:25:01Z&se=2018-09-17T08:25:01Z&spr=https&sv=2017-11-09&sig=0000000000%3D&sr=b', 'http://download.windowsupdate.com/c/msdownload/update/software/secu/2018/08/windows10.0-kb4343909-x64_f931af6d56797388715fe3b0d97569af7aebdae6.msu'"
      }
    ]
    ...
```

## Supported Orchestrators

This has been tested with Kubernetes clusters, and does not depend on any specific version.

## Selecting Patches

### Cumulative Updates

If you would like to include a cumulative update as part of your deployment that isn't in the Windows Server with Containers image, then follow these steps.

1. Browse to [Windows 10 Update History](https://support.microsoft.com/en-us/help/4099479), and follow the link to the right version (1709 or 1803) in the left. This page should be titled "Windows 10 and Windows Server update history" because the links also lead you to Windows Server updates.
2. Next, look for the latest patch in the lower left such as ["August 14, 2018—KB4343909 (OS Build 17134.228)"](https://support.microsoft.com/en-us/help/4343909), and click that link.
3. Scroll down to "How to get this update", and click on the ["Microsoft Update Catalog"](http://catalog.update.microsoft.com/v7/site/Search.aspx?q=KB4343909) link.
4. Find the row for `(year)-(month) Cumulative Update for Windows Server 2016 ((1709 or 1803)) for x64-based Systems (KB######)`, and click the "Download" button.
5. This will pop up a new window with a hyperlink such as `windows10.0-kb4343909-x64_f931af6d56797388715fe3b0d97569af7aebdae6.msu`. Copy that link.
6. Include that link in the `extensionParameters` as shown in the [Example](#Example)

### Supplied by Microsoft support

Once you have downloaded a private hotfix from Microsoft support, it needs to be put in an Azure-accessible location. The easiest way to do this is to create an [Azure Blob Storage](https://docs.microsoft.com/en-us/azure/storage/common/storage-create-storage-account#blob-storage-accounts) account. Once uploaded, you can create a private link with a key to access it that will work with this extension.

1. If you haven't already, install the [Azure CLI](https://docs.microsoft.com/cli/azure/get-started-with-az-cli2), and run `az login` to log in to Azure
2. Copy the sample below for either bash (Linux, Mac, or WSL), or PowerShell (Windows), and modify the variables at the top.


#### Using az cli and bash to upload the patch

```bash
#!/bin/bash
export resource_group=patchgroup
export storage_location=westus2
export storage_account_name=privatepatches
export container_name=hotfix
export blob_name=Windows10.0-KB999999-x64-InstallForTestingPurposesOnly.exe
export file_to_upload=Windows10.0-KB999999-x64-InstallForTestingPurposesOnly.exe

echo "Creating the group..."
az group create --location $storage_location --resource-group $resource_group

echo "Creating the storage account..."
az storage account create --location $storage_location --name $storage_account_name --resource-group $resource_group --sku Standard_LRS

echo "Getting the connection string..."
export AZURE_STORAGE_CONNECTION_STRING="`az storage account show-connection-string --name $storage_account_name --resource-group $resource_group`"

echo "Creating the container..."
az storage container create --name $container_name

echo "Uploading the file..."
az storage blob upload --container-name $container_name --file $file_to_upload --name $blob_name

echo "Getting a read-only SAS token, good for 30 days..."
export EXPIRY=`date +"%Y-%m-%dT%H:%M:%SZ" -d '30 days'`
export TEMPORARY_SAS=`az storage blob generate-sas --container-name $container_name --name $blob_name --permissions r --expiry $EXPIRY`

echo "Getting a URL to the file..."
export ABSURL=`az storage blob url --container-name $container_name --name $blob_name --sas-token $TEMPORARY_SAS`

echo "Full URL including access token:"
echo "$ABSURL?$TEMPORARY_SAS" | sed "s/\"//g"
```

#### Using az cli and PowerShell to upload the patch

```powershell
$resource_group="patchgroup"
$storage_location="westus2"
$storage_account_name="privatepatches"
$container_name="hotfix"
$blob_name="Windows10.0-KB999999-x64-InstallForTestingPurposesOnly.exe"
$file_to_upload="Windows10.0-KB999999-x64-InstallForTestingPurposesOnly.exe"

Write-Host "Creating the group..."
az group create --location $storage_location --resource-group $resource_group

Write-Host "Creating the storage account..."
az storage account create --location $storage_location --name $storage_account_name --resource-group $resource_group --sku Standard_LRS

Write-Host "Getting the connection string..."
$ENV:AZURE_STORAGE_CONNECTION_STRING = az storage account show-connection-string --name $storage_account_name --resource-group $resource_group

Write-Host "Creating the container..."
az storage container create --name $container_name

Write-Host "Uploading the file..."
az storage blob upload --container-name $container_name --file $file_to_upload --name $blob_name

Write-Host "Getting a read-only SAS token, good for 30 days..."
$EXPIRY = ([DateTime]::Now + [timespan]::FromDays(30)).ToUniversalTime().ToString("yyyy-MM-ddTHH:mm:ssZ")
$TEMPORARY_SAS = az storage blob generate-sas --container-name $container_name --name $blob_name --permissions r --expiry $EXPIRY

Write-Host "Getting a URL to the file..."
$ABSURL = az storage blob url --container-name $container_name --name $blob_name --sas-token $TEMPORARY_SAS

Write-Host "Full URL including access token:"
$full_url = "$($ABSURL)?$($TEMPORARY_SAS)".Replace("""","")
$full_url | Write-Host
$full_url | Set-Clipboard
```

The last line of the script will output a URL, and also put it on the Windows clipboard. Copy it into `extensionParameters` as shown in the sample above. Do not share this URL, keep it private.

## Troubleshooting

Extension execution output is logged to files found under the following directory on the target virtual machine.

```powershell
C:\WindowsAzure\Logs\Plugins\Microsoft.Compute.CustomScriptExtension
```

The specified files are downloaded into the following directory on the target virtual machine.

```powershell
C:\Packages\Plugins\Microsoft.Compute.CustomScriptExtension\1.*\Downloads\<n>
```
