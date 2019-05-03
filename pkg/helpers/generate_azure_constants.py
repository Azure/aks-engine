#!/usr/bin/env python

import json
import subprocess


MIN_CORES_DCOS = 2
DCOS_MASTERS_EPHEMERAL_DISK_MIN = 16384


def get_all_sizes():
    locations = json.loads(subprocess.check_output(['az', 'account', 'list-locations', '-o', 'json']).decode('utf-8'))
    size_map = {}

    for location in locations:
        try:
            # NOTE: "az vm list-sizes" fails in francesouth, australiacentral, australiacentral2, and southafricawest.
            sizes = json.loads(subprocess.check_output(['az', 'vm', 'list-sizes', '-l', location['name'], '-o', 'json']).decode('utf-8'))
            for size in sizes:
                if not size['name'] in size_map and not size['name'].split('_')[0] == 'Basic' and not size['name'][-6:] == '_Promo':
                    size_map[size['name']] = size
        except subprocess.CalledProcessError:
            continue

    # ensure Azure ML/AI SKUs stay listed
    if 'Standard_PB12s' not in size_map:
        size_map.update({'Standard_PB12s': {'name': 'Standard_PB12s', 'numberOfCores': MIN_CORES_DCOS, 'resourceDiskSizeInMb': DCOS_MASTERS_EPHEMERAL_DISK_MIN}})
    if 'Standard_PB24s' not in size_map:
        size_map.update({'Standard_PB24s': {'name': 'Standard_PB24s', 'numberOfCores': MIN_CORES_DCOS, 'resourceDiskSizeInMb': DCOS_MASTERS_EPHEMERAL_DISK_MIN}})
    if 'Standard_PB6s' not in size_map:
        size_map.update({'Standard_PB6s': {'name': 'Standard_PB6s', 'numberOfCores': MIN_CORES_DCOS, 'resourceDiskSizeInMb': DCOS_MASTERS_EPHEMERAL_DISK_MIN}})

    return size_map


def get_dcos_master_map(size_map):
    master_map = {}

    for key in size_map.keys():
        size = size_map[key]
        if size['numberOfCores'] >= MIN_CORES_DCOS and \
           size['resourceDiskSizeInMb'] >= DCOS_MASTERS_EPHEMERAL_DISK_MIN:
            master_map[size['name']] = size

    return master_map


def get_locations():
    output = json.loads(subprocess.check_output(['az', 'account', 'list-locations', '-o', 'json']).decode('utf-8'))

    locations = [l['name'] for l in output]

    # Hard-code Azure China Cloud locations
    locations.append('chinanorth')
    locations.append('chinaeast')
    locations.append('chinanorth2')
    locations.append('chinaeast2')

    # Add two Canary locations
    locations.append('centraluseuap')
    locations.append('eastus2euap')

    # Add US DoD locations
    locations.append('usdodcentral')
    locations.append('usdodeast')

    locations = sorted(locations)
    return locations


def get_storage_account_type(size_name):
    capability = size_name.split('_')[1]
    if 'S' in capability or 's' in capability:
        return "Premium_LRS"
    return "Standard_LRS"


def get_file_contents(dcos_master_map, kubernetes_size_map, locations):
    text = r"""// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package helpers

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

    text += r"""	"chinaeast",
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

// GetDCOSMasterAllowedSizes returns the master allowed sizes
func GetDCOSMasterAllowedSizes() string {
        return `      "allowedValues": [
"""
    dcos_master_map_keys = sorted(dcos_master_map.keys())
    for key in dcos_master_map_keys[:-1]:
        text += '        "' + key + '",\n'
    text += '        "' + dcos_master_map_keys[-1] + '"\n'

    text += r"""    ],
`
}

// GetKubernetesAllowedSizes returns the allowed sizes for Kubernetes agent
 func GetKubernetesAllowedSizes() string {
        return `      "allowedValues": [
"""
    kubernetes_agent_map_keys = sorted(kubernetes_size_map.keys())
    for key in kubernetes_agent_map_keys[:-1]:
        text += '        "' + key + '",\n'
    text += '        "' + kubernetes_agent_map_keys[-1] + '"\n'
    text += r"""    ],
`
}

// GetSizeMap returns the size / storage map
func GetSizeMap() string {
    return `    "vmSizesMap": {
"""
    merged_map = {}
    for key in kubernetes_agent_map_keys:
        size = kubernetes_size_map[key]
        if key not in merged_map:
            merged_map[size['name']] = size

    merged_map_keys = sorted(merged_map.keys())
    for key in merged_map_keys[:-1]:
        size = merged_map[key]
        text += '    "' + size['name'] + '": {\n'
        storage_account_type = get_storage_account_type(size['name'])
        text += '      "storageAccountType": "' + storage_account_type + '"\n    },\n'

    key = merged_map_keys[-1]
    size = merged_map[key]
    text += '    "' + size['name'] + '": {\n'
    storage_account_type = get_storage_account_type(size['name'])
    text += '      "storageAccountType": "' + storage_account_type + '"\n    }\n'

    text += r"""   }
`
}"""
    return text


def main():
    outfile_name = 'pkg/helpers/azureconst.go'
    all_sizes = get_all_sizes()
    dcos_master_map = get_dcos_master_map(all_sizes)
    kubernetes_size_map = all_sizes
    locations = get_locations()
    text = get_file_contents(dcos_master_map, kubernetes_size_map, locations)

    with open(outfile_name, 'w') as outfile:
        outfile.write(text)

    subprocess.check_call(['gofmt', '-w', outfile_name])


if __name__ == '__main__':
    main()
