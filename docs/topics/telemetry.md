# Telemetry

## Custom Script Extensions

The AKS Engine team has instrumented custom script extensions (CSEs) so collection of telemetry may 
be enabled. This instrumentation is used to collect aggregate, non-identifiable data to help us 
answer gather the following information:

- Timing metrics for various operations performed during CSE execution
- Metrics on failures encountered during CSE execution
- General product usage
  - kubernetes and component version usage
  - Node pool counts and VM Sku data
  - Which versions of VHDs are still in use

This data will be used to monitor the health of cluster deployments (including AKS clusters) that 
are deployed with AKS Engine as well as to help us prioritize future investments/feature work in 
the tool.

## Configuration

Collection of all telemetry is currently **disabled** by default.

Telemetry can be enabled by setting the `enableTelemetry` feature flag to `true` in the 
apimodel.json. When `enableTelemetry` is set to `true`, telemetry will be sent to the AKS Engine 
Application Insights cluster.

Telemetry can be routed to an additional Application Insights instance by specifying 
`telemetryProfile.applicationInsightsKey` with the value of the instrumentation key of your 
Application Insights instance. 

That means you can see and use the same data that is being sent to the AKS Engine team.

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

### Windows Specific Differences (wip)

Currently, Windows CSE will only log telemetry to the AKS Engine Application Insights instance, even
if `applicationInsightsKey` is set. In a future change set, this feature will be enabled.

### Collection Settings and Results

AKS Engine has the concept of a system defined AKS Engine Application Insights key as well as a user
defined key. The table below describes where telemetry is logged based on the configuration
settings.

| enableTelemetry  | applicationInsightsKey set | result |
| ------------- | ------------- | ------------- |
| true | true | both user and AKS Engine telemetry is tracked |
| false  | true  | no telemetry is tracked |
| true | false | only AKS Engine telemetry is tracked |
| false | false | no telemetry is tracked |