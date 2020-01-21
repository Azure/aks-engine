# [WIP]

## Motivation

The aks-engine team is working towards collecting telemetry from custom script extensions (CSEs) used to configure Kubernetes nodes.
This data will be used to collect aggregate, non-identifiable data to help us answer gather the following information:

- Timing metrics for various operations performed during CSE execution
- Metrics on failures encountered during CSE execution
- General product usage
  - kubernetes and component version usage
  - Node pool counts and VM Sku data
  - Which versions of VHDs are still in use

This data will be used to monitor the health of cluster deployments (including AKS clusters) that are deployed with aks-engine as well as to help us prioritize future investments/feature work in the tool.

## Configuration

Collection of all telemetry is currently **disabled** by default.

Telemetry can be enabled by setting the `EnableTelemetry` feature flag to `true` in the apimodel.json.

Telemetry can be routed to different application insights instance be specifying `telemetryProfile.applicationInsightsKey` to the instrumentation key of your application insights instance.

``` javascript
{
  "properties": {
    "featureFlags": {
      "enableTelemetry": true
    },
    ...
    "telemetryProfile": {
      "applicationInsightsKey": "7440e089-79e1-4cb4-96b7-d8e9a460eb74"
    }
  }
}
```
