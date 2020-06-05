# Running Tests

If you are an AKS Engine developer, running local E2E tests to validate changes can greatly increase iterative velocity.

As mentioned briefly in the [developer guide](developer-guide.md), a `make` target is maintained to provide convenient shell invocation of the E2E test runner across for generic, configurable usage:

```sh
$ ORCHESTRATOR_RELEASE=1.18 CLUSTER_DEFINITION=examples/kubernetes.json SUBSCRIPTION_ID=$TEST_AZURE_SUB_ID CLIENT_ID=$TEST_AZURE_SP_ID CLIENT_SECRET=$TEST_AZURE_SP_PW TENANT_ID=$TEST_AZURE_TENANT_ID LOCATION=$AZURE_REGION CLEANUP_ON_EXIT=false make test-kubernetes
```

The above, simple example describes an E2E test invocation against a base cluster configuration defined by the API model at `examples/kubernetes.json`, overriding any specific Kubernetes version therein to validate the most recent, supported v1.18 release; using Azure service principal authentication defined in the various `$TEST_AZURE_`* environment variables; deployed to the region defined by the environment variable `$AZURE_REGION`; and finally, we tell the E2E test runner not to delete the cluster resources (i.e., the resource group) following the completion of the tests.

Example output from such an invocation is [here](e2e-output-example.log). If your test run succeeded, you'll see this in your console stdout at the conclusion of the test run:

```sh
Ran 40 of 54 Specs in 1077.006 seconds
SUCCESS! -- 40 Passed | 0 Failed | 0 Pending | 14 Skipped
PASS
```

If any test failures occurred, the output will report which tests failed; we'll discuss how you can target just those tests during development below.

## E2E Test Runner Configuration

The E2E test runner is designed to be flexible across a wide range of cluster configurations. Below is the full set of configurable environment variables:

| Name       | Required | Description                                                   |
| ---------- | -------- | ------------------------------------------------------------- |
| `SKIP_TEST` | no      | Don't run any E2E tests, just use the E2E test runner to create a new cluster. E.g., `SKIP_TEST=false` |
| `NAME` | no      | Allows you to re-run E2E tests on an existing cluster. Assumes the cluster has been created via a prior E2E test run, and that its generated artifacts still exist in the relative `_output/` directory. The value of `NAME` should be equal to the resource group created by the E2E test runner, and that value will also map to a directory under `_output/`. E.g., a value of `kubernetes-westus2-13811` will map to a resource group in the configured subscription, using the configured service principal credentials, and a directory under `_output/kubernetes-westus2-13811/` will exist with all cluster configuration artifacts. |
| `LOCATION` | yes      | The Azure region to build your cluster in. E.g., `LOCATION=westus2`. Required if `REGIONS` is empty. Not required if `NAME` is provided, i.e., if you are re-testing an existing E2E-created cluster. |
| `REGIONS` | no      | When you want to deploy to a randomly selected region from a known-working set of regions. E.g., `REGIONS=westus2,westeurope,canadacentral`. Required if `LOCATION` is empty; not required if `NAME` is provided. |
| `CLUSTER_DEFINITION` | yes      | The API model to use as cluster configuration input for creating a new cluster. E.g., `CLUSTER_DEFINITION=examples/kubernetes.json`. Not required if `NAME` is provided. |
| `CLEANUP_ON_EXIT` | no      | Delete cluster after running E2E. E.g., `CLEANUP_ON_EXIT=true`. Default is false. |
| `CLEANUP_IF_FAIL` | no      | Delete cluster only if E2E failed. E.g., `CLEANUP_IF_FAIL=false`. Default is false. |
| `STABILITY_ITERATIONS` | no      | How many basic functional cluster tests to run in rapid succession as a part of E2E validation. This is useful for simulation continual usage of basic cluster reconciliation functionality (schedule/delete a pod, resolve a DNS lookup, etc). E.g., `STABILITY_ITERATIONS=100`. Default is 3. |
| `TIMEOUT` | no      | How much timeout tolerance for tests? Decrease timeout tolerance to do performance-type tests, increase to allow for more operational variability and possibly reduce flakes. E.g., `TIMEOUT=10m`. Default is 20m. |
| `LB_TIMEOUT` | no      | How much timeout tolerance for Load Balancer tests? Decrease timeout tolerance to do performance-type tests, increase to allow for more operational variability and possibly reduce flakes. E.g., `LB_TIMEOUT=5m`. Default is 20m. |
| `GINKGO_FAIL_FAST` | no      | Stop the suite right after the first failure? E.g., `GINKGO_FAIL_FAST=false`. Default is true. |
| `GINKGO_FOCUS` | no      | Regular expression string to pass to test runner to run only a subset of tests that match the regular expression. E.g., `GINKGO_FOCUS="should be able to produce working LoadBalancers"`. Only works if `GINKGO_FAIL_FAST` is set to true. |
| `GINKGO_SKIP` | no      | Regular expression string to pass to test runner to skip the subset of tests that match the regular expression. E.g., `GINKGO_SKIP="should be able to attach azure file"`. Only works if `GINKGO_FAIL_FAST` is set to true. |
| `DEBUG_AFTERSUITE` | no      | Print out Kubernetes resources and logs after E2E suite. This is especially useful if you have to delete your cluster after running the test and you wish to debug at a later time. E.g., `DEBUG_AFTERSUITE=true`. Default is false. |
