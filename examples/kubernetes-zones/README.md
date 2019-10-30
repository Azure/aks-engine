# Availability Zones

To protect your cluster from datacenter-level failures, you can enable the Availability Zones feature for your cluster by configuring `"availabilityZones"` for the master profile and all of the agentPool profiles in the cluster definition.

 - This feature only applies to Kubernetes clusters version 1.12+.
 - Supported values are arrays of strings, each representing a supported availability zone in a region for your subscription. For example, `"availabilityZones": ["1","2"]` indicates zone 1 and zone 2 can be used.

    > To get supported zones for a region in your subscription, run `az vm list-skus --location centralus --query "[?name=='Standard_DS2_v2'].[locationInfo, restrictions]" -o table`. You should see values like `'zones': ['2', '3', '1']` appear in the first column. If `NotAvailableForSubscription` appears in the output, then create an Azure support ticket to enable zones for that region.

- To ensure high availability, each profile must define at least two nodes per zone. For example, an agent pool profile with 2 zones must have at least 4 nodes total: `"availabilityZones": ["1","2"],"count": 4`.
- When `"availabilityZones"` is configured, the `"loadBalancerSku"` will default to `Standard` as Standard LoadBalancer is required for availability zones.

Here is an [example of a Kubernetes cluster with Availability Zones support](../e2e-tests/kubernetes/zones/definition.json)

```json
{
    "apiVersion": "vlabs",
    "properties": {
      "orchestratorProfile": {
        "orchestratorType": "Kubernetes",
        "orchestratorRelease": "1.13"
      },
      "masterProfile": {
        "count": 5,
        "dnsPrefix": "",
        "vmSize": "Standard_DS2_v2",
        "availabilityProfile": "VirtualMachineScaleSets",
        "availabilityZones": [
            "1",
            "2"
        ]
      },
      "agentPoolProfiles": [
        {
            "name": "agentpool",
            "count": 4,
            "vmSize": "Standard_DS2_v2",
            "availabilityProfile": "VirtualMachineScaleSets",
            "availabilityZones": [
                "1",
                "2"
            ]
        }
      ],
      "linuxProfile": {
        "adminUsername": "azureuser",
        "ssh": {
          "publicKeys": [
            {
              "keyData": ""
            }
          ]
        }
      },
      "servicePrincipalProfile": {
        "clientId": "",
        "secret": ""
      }
    }
}

```

To validate availability zones are working as expected, run the following commands:

```bash
kubectl get nodes --show-labels | grep failure-domain.beta.kubernetes.io/zone

...,failure-domain.beta.kubernetes.io/zone=eastus2-1, ...
...,failure-domain.beta.kubernetes.io/zone=eastus2-2, ...

```

Each node in the cluster should have `REGION-ZONE` as values for the `failure-domain.beta.kubernetes.io/zone` label.
