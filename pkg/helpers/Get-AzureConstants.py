#!/usr/bin/python

# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT license.

import json
import subprocess
from time import gmtime, strftime

time = strftime("%Y-%m-%d %H:%M:%S", gmtime())

def getAllSizes():
    locations = json.loads(subprocess.check_output(['az', 'account', 'list-locations']).decode('utf-8'))
    sizeMap = {}

    for location in locations:
        # No registered resource provider found for location 'francesouth' and API version '2017-12-01' for type 'locations/vmSizes'.
        # No registered resource provider found for location 'australiacentral' and API version '2017-12-01' for type 'locations/vmSizes'.
        # No registered resource provider found for location 'australiacentral2' and API version '2017-12-01' for type 'locations/vmSizes'.
        if location['name'] != "francesouth" and location['name'] != "australiacentral" and location['name'] != "australiacentral2":
            sizes = json.loads(subprocess.check_output(['az', 'vm', 'list-sizes', '-l', location['name']]).decode('utf-8'))
            for size in sizes:
                if not size['name'] in sizeMap and not size['name'].split('_')[0] == 'Basic':
                    sizeMap[size['name']] = size

    return sizeMap

min_cores_k8s = 1

def getMasterAgentMap(sizeMap):
    agentMap = {}

    for key in sizeMap.keys():
        size = sizeMap[key]
        if size['numberOfCores'] >= min_cores_k8s:
            agentMap[size['name']] = size

    return agentMap

def getLocations():
    locations = json.loads(subprocess.check_output(['az', 'account', 'list-locations']).decode('utf-8'))

    locationList = [l['name'] for l in locations]

    #hard code Azure China Cloud location
    locationList.append('chinanorth')
    locationList.append('chinaeast')
    locationList.append('chinanorth2')
    locationList.append('chinaeast2')
    # Adding two Canary locations
    locationList.append('centraluseuap')
    locationList.append('eastus2euap')

    locationList = sorted(locationList)
    return locationList

def getStorageAccountType(sizeName):
    capability = sizeName.split('_')[1]
    if 'S' in capability or 's' in capability:
        return "Premium_LRS"
    else:
        return "Standard_LRS"

def getFileContents(masterAgentMap, kubernetesAgentMap, sizeMap, locations):
    text = r"""package helpers

// AUTOGENERATED FILE """

    text += r"""

// GetAzureLocations provides all azure regions in prod.
// Related powershell to refresh this list:
//   Get-AzureRmLocation | Select-Object -Property Location
func GetAzureLocations() []string{
return []string{
"""
    for location in locations:
        text += '        "' + location + '",' + '\n'

    text += r"""        "chinaeast",
	"chinanorth",
	"chinanorth2",
	"chinaeast2",
	"germanycentral",
	"germanynortheast",
	"usgovvirginia",
	"usgoviowa",
	"usgovarizona",
	"usgovtexas",
    "francecentral",
}
}

// GetMasterAgentAllowedSizes returns the agent allowed sizes
func GetMasterAgentAllowedSizes() string {
        return `      "allowedValues": [
"""
    masterAgentMapKeys = sorted(masterAgentMap.keys())
    for key in masterAgentMapKeys[:-1]:
        text += '        "' + key + '",\n'
    text += '        "' + masterAgentMapKeys[-1] + '"\n'
    text += r"""    ],
`
}

// GetKubernetesAgentAllowedSizes returns the allowed sizes for Kubernetes agent
func GetKubernetesAgentAllowedSizes() string {
        return `      "allowedValues": [
"""
    kubernetesAgentMapKeys = sorted(kubernetesAgentMap.keys())
    for key in kubernetesAgentMapKeys[:-1]:
        text += '        "' + key + '",\n'
    text += '        "' + kubernetesAgentMapKeys[-1] + '"\n'
    text += r"""    ],
`
}

// GetSizeMap returns the size / storage map
func GetSizeMap() string {
    return `    "vmSizesMap": {
"""
    mergedMap = {}
    for key in kubernetesAgentMapKeys:
        size = kubernetesAgentMap[key]
        if not key in mergedMap:
            mergedMap[size['name']] = size

    mergedMapKeys = sorted(mergedMap.keys())
    for key in mergedMapKeys[:-1]:
        size = mergedMap[key]
        text += '    "' + size['name'] + '": {\n'
        storageAccountType = getStorageAccountType(size['name'])
        text += '      "storageAccountType": "' + storageAccountType + '"\n    },\n'

    key = mergedMapKeys[-1]
    size = mergedMap[key]
    text += '    "' + size['name'] + '": {\n'
    storageAccountType = getStorageAccountType(size['name'])
    text += '      "storageAccountType": "' + storageAccountType + '"\n    }\n'

    text += r"""   }
`
}

// GetClassicAllowedSizes returns the classic allowed sizes
func GetClassicAllowedSizes() string {
        return `      "allowedValues": [
"""
    sizeMapKeys = sorted(sizeMap.keys())
    for key in sizeMapKeys[:-1]:
        text += '        "' + sizeMap[key]['name'] + '",\n'
    key = sizeMapKeys[-1]
    text += '        "' + sizeMap[key]['name'] + '"\n'

    text += r"""    ],
`
}

// GetClassicSizeMap returns the size / storage map
func GetClassicSizeMap() string {
    return `    "vmSizesMap": {
"""
    sizeMapKeys = sorted(sizeMap.keys())
    for key in sizeMapKeys[:-1]:
        text += '        "' + sizeMap[key]['name'] + '": {\n'
        storageAccountType = getStorageAccountType(size['name'])
        text += '      "storageAccountType": "' + storageAccountType + '"\n    },\n'
    key = sizeMapKeys[-1]
    text += '        "' + sizeMap[key]['name'] + '": {\n'
    storageAccountType = getStorageAccountType(size['name'])
    text += '      "storageAccountType": "' + storageAccountType + '"\n    }\n'

    text += r"""   }
`
}"""
    return text


def main():
    outfile = 'pkg/helpers/azureconst.go'
    allSizes = getAllSizes()
    masterAgentMap = getMasterAgentMap(allSizes)
    kubernetesAgentMap = allSizes
    locations = getLocations()
    text = getFileContents(masterAgentMap, kubernetesAgentMap, allSizes, locations)

    with open(outfile, 'w') as f:
        f.write(text)

    subprocess.check_call(['gofmt', '-w', outfile])

if __name__ == '__main__':
    main()
