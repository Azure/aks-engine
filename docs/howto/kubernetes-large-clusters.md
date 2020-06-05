# Large Kubernetes Clusters with AKS Engine

## Background

AKS Engine supports using exponential cloud backoff that is a feature of Kubernetes v1.6.6 and newer. Cloud backoff allows Kubernetes nodes to backoff on HTTP 429 errors that are usually caused by exceeding Azure API limits.

## To Use

Declare your kubernetes cluster API model config as you normally would, with the following requirements:
- We recommend the use of smaller pools (e.g., count of 20) over larger pools (e.g., count of 100); produce your desired total node count with lots of pools, as opposed to as few as possible.
- We also recommend using large vmSize configurations to reduce node counts, where appropriate. Make sure you have a defensible infrastructure justification for more nodes in terms of node count (for example as of kubernetes 1.7 there is a 100 pods per node limit), instead of opting to use more powerful nodes. Doing so reduces cluster complexity, and azure resource administrative overhead. As Kubernetes excels in binpacking pods onto available instances, vertically scaling VM sizes (more CPU/RAM) is a better approach for expanding cluster capacity, if you are not approaching the pod-per-node limit.

## Backoff configuration options

The following configuration parameters are available in the `properties.orchestratorProfile.kubernetesConfig` configuration object in the API model specification:

```json
    "cloudProviderBackoff": {
      "value": "true" // if true, enable backoff
    },
    "cloudProviderBackoffDuration": {
      "value": "5" // how many seconds for initial backoff retry attempt
    },
    "cloudProviderBackoffExponent": {
      "value": "1.5" // exponent for successive backoff retries
    },
    "cloudProviderBackoffJitter": {
      "value": "1" // non-1 values add jitter to retry intervals
    },
    "cloudProviderBackoffRetries": {
      "value": "6" // maximum retry attempts before failure
    },
    "cloudProviderRatelimit": {
      "value": "true" // if true, enforce rate limits for azure API calls
    },
    "cloudProviderRatelimitBucket": {
      "value": "10" // number of requests in queue
    },
    "cloudProviderRatelimitQPS": {
      "value": "3" // rate limit QPS
    },
    "kubeletConfig": {
      "--node-status-update-frequency": "1m" // how often kubelet posts node status to master
    },
    "controllerManagerConfig": {
      "--node-monitor-grace-period": "5m", // duration after which controller manager marks an AWOL node as NotReady
      "--pod-eviction-timeout": "1m", // grace period for deleting pods on failed nodes
      "--route-reconciliation-period": "1m" // how often to reconcile cloudprovider-originating node routes
    }
```
The [examples/largeclusters/kubernetes.json](https://github.com/Azure/aks-engine/blob/master/examples/largeclusters/kubernetes.json) API model example suggests how you might opt into these large cluster features following the guidelines above.
