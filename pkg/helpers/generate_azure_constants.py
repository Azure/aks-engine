#!/usr/bin/env python

import json
import re
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
                if not size['name'] in size_map and not size['name'].split('_')[0] == 'Basic':
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


def sku_sort_key(sku):
    """Split a SKU string into a list of strings and integers, for sorting."""
    return [int(e) if e.isdigit() else e for e in re.findall(r'\d+|\D+', sku)]


def get_accelerated_skus():
    """Return a list of SKUs that support accelerated networking."""
    # Start with some grandfathered SKUs for backwards compatibility.
    skus = [
        "AZAP_Performance_ComputeV17C",
        "SQLGL",
        "SQLGLCore",
        "Standard_D12_v2_ABC",
        "Standard_D13_v2_ABC",
        "Standard_D14_v2_ABC",
        "Standard_D15_v2_ABC",
        "Standard_D32-16s_v3",
        "Standard_D32-8s_v3",
        "Standard_D3_v2_ABC",
        "Standard_D40_v3",
        "Standard_D40s_v3",
        "Standard_D4_v2_ABC",
        "Standard_D5_v2_ABC",
        "Standard_D64-16s_v3",
        "Standard_D64-32s_v3",
        "Standard_E32-16_v3",
        "Standard_F16_ABC",
        "Standard_F4_ABC",
        "Standard_F8_ABC",
        "Standard_L96s_v2",
     ]

    query = r"[? starts_with(name, `Standard`) && !ends_with(name, `Promo`)]"
    results = json.loads(
        subprocess.check_output(
            ["az", "vm", "list-skus", "--all", "-o", "json", "--query", query]
        ).decode("utf-8")
    )
    for r in results:
        sku = r["name"]
        capabilities = r.get('capabilities') or []

        for cap in capabilities:
            name, value = cap["name"], cap["value"]
            # Add SKUs with the `AcceleratedNetworkingEnabled` capability equal to "True".
            if name == "AcceleratedNetworkingEnabled":
                if value in ("True", True):
                    skus.append(sku)
                break
            # If there's no explicit capability, infer from the rules in the documentation.
            elif name == "vCPUs":
                # add D/DSv2 and F/Fs with vCPUs >= 2
                if re.match(r'Standard_(DS\d+?.*v2|F)', sku) is not None and int(value) >= 2:
                    skus.append(sku)
                # add D/Dsv3, E/Esv3, Fsv2, Lsv2, Ms/Mms and Ms/Mmsv2 with vCPUs >= 4
                elif re.match(r'Standard_([D|E]\d+s?.*v3|[F|L]\d+s.*v2|M\d+s?.*(v2)?)', sku) is not None and int(value) >= 4:
                    skus.append(sku)

    return set(skus)


def get_file_contents(dcos_master_map, kubernetes_size_map, locations, skus):
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

// GetKubernetesAllowedVMSKUs returns the allowed sizes for Kubernetes agent
 func GetKubernetesAllowedVMSKUs() string {
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
}

// AcceleratedNetworkingSkus are those Azure VM SKUs that support accelerated networking.
//
// From https://docs.microsoft.com/en-us/azure/virtual-network/create-vm-accelerated-networking-cli:
// Accelerated Networking is supported on most general purpose and compute-optimized instance sizes
// with 2 or more vCPUs. These supported series are: D/DSv2 and F/Fs.
// On instances that support hyperthreading, Accelerated Networking is supported on VM instances
// with 4 or more vCPUs. Supported series are: D/Dsv3, E/Esv3, Fsv2, Lsv2, Ms/Mms and Ms/Mmsv2.
var AcceleratedNetworkingSkus []string = []string{
"""
    for name in sorted(skus):
        text += '\t"{}",\n'.format(name)
    text += "}"

    return text


def main():
    outfile_name = 'pkg/helpers/azureconst.go'
    all_sizes = get_all_sizes()
    dcos_master_map = get_dcos_master_map(all_sizes)
    kubernetes_size_map = all_sizes
    locations = get_locations()
    accelerated_networking_skus = get_accelerated_skus()
    text = get_file_contents(dcos_master_map, kubernetes_size_map, locations, accelerated_networking_skus)

    with open(outfile_name, 'w') as outfile:
        outfile.write(text)

    subprocess.check_call(['gofmt', '-w', outfile_name])


if __name__ == '__main__':
    main()
