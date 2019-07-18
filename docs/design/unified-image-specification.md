
# Unified image specification for node pools

## Motivation

There are inconsistencies in where and how you can specify images in the AKS Engine cluster definitions.
This leads to unnecessary complexity in the aks-engine codebase and potential confusion for AKS Engine users.

Additionally, it is not possible to specify different images for different pools running Windows which breaks some upgrade/test scenarios.

Images for virtual machines in Azure are specified the same way regardless of guest OS. This proposal aims to unify how AKS Engine cluster defintions specify images for Linux and Windows.

## Current behavior

### Linux

Images are specified in the agentPoolProfile one of the following ways. 
Specifing the image in agentPoolProfile allows multiple pools running different images.

* Distro - Specifies a curated list of images for various 'distros' (pulled from marketplace or shared image gallery)

```json
"Distro": "[ubuntu | rhel | coreos | AKSUbuntu1604 | AKSUbuntu1804]"
```

* Image Reference - Specifies an Azure manage image
  * If only name and resourceGroup are specified, image will be pulled from current subscription
  * If subscription, gallery, and version are also specified image will be pulled from a shared gallery

```json
"imageReference": {
    "name": "" ,
    "resourceGroup": "",

    (optional)
    "subscription": "",
    "gallery": "",
    "version": ""
}
```

### Windows

Images are specified in the WindowsProfile one of the following ways.
Specifying images in WindowsProfile only allows a single image to be used for all Windows images in the cluster.

* WindowsImageSourceUrl - Specifies a vhd in an Azure storage blob

```json
"WindowsImageSourceUrl" : ""
```

* Publisher/Offer/Sku/Version - Specifies an image in Azure image marketplace

```json
"windowsPublisher" = "",
"windowsOffer" = "",
"windowsSku" = "",
"imageVersion" = ""
```

## Proposed Behavior

### Linux and Windows

Images are specified as port of agentPoolProfile.

* Distro - Specifies a curated list of images for various 'distros' (pulled from marketplace or shared image gallery)

```json
"Distro": "[ubuntu | rhel | coreos | AKSUbuntu1604 | AKSUbuntu1804 | AKSWindows1809 | AKSWindows1903 | AKSWindows2019]"
```

* Image URL - path to a vhd in azure blob storage

```json
"image" : {
    "url": ""
}
```

* Image Reference - Specifies an Azure manage image
  * If only name and resourceGroup are specified, image will be pulled from current subscription
  * If subscription, gallery, and version are also specified image will be pulled from a gallery

```json
"image": {
    "reference": {
        "name": "" ,
        "resourceGroup": "",

        (optional)
        "subscription": "",
        "gallery": "",
        "version": ""
    }
}
```

* Marketplace Image - Specifies an Azure marketplace image

```json
"image": {
    "marketplace": {
        "publisher": "",
        "offer": "",
        "sku": "",
        "version": ""
    }
}
```

### Maintaining backwards compatability

In order to maintain backwards compatability previously supported settings for specifying images for agent pools.

For Linux agent pools if any of the imageReference.* properties are specified on AgentPoolPorfile they will get mapped to their coresponding image.reference.* properties.

For Windows agent pools if all of windowsPublisher, windowsOffer, and windowsSku are specified on the WindowsProfile then these settings will be used for all windows agent pools (including imageVersion if specified).
If WindowsImageSourceURL is specified then this image will be used for all Windows agent pools.
