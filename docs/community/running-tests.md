# Running Tests

If you are an AKS Engine developer, running local E2E tests to validate changes can greatly increase iterative velocity.

As mentioned briefly in the [developer guide](developer-guide.md), a `make` target is maintained to provide convenient shell invocation of the E2E test runner across for generic, configurable usage:

```sh
$ ORCHESTRATOR_RELEASE=1.18 CLUSTER_DEFINITION=examples/kubernetes.json SUBSCRIPTION_ID=$TEST_AZURE_SUB_ID CLIENT_ID=$TEST_AZURE_SP_ID CLIENT_SECRET=$TEST_AZURE_SP_PW TENANT_ID=$TEST_AZURE_TENANT_ID LOCATION=$AZURE_REGION CLEANUP_ON_EXIT=false make test-kubernetes
```

The above, simple example describes an E2E test invocation against a base cluster configuration defined by the api model at `examples/kubernetes.json`, overriding any specific Kubernetes version therein to validate the most recent, supported v1.18 release; using Azure service principal authentication defined in the various `$TEST_AZURE_`* environment variables; deployed to the region defined by the environment variable `$AZURE_REGION`; and finally, we tell the E2E test runner not to delete the cluster resources (i.e., the resource group) following the completion of the tests.

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
| `CLUSTER_DEFINITION` | yes      | The api model to use as cluster configuration input for creating a new cluster. E.g., `CLUSTER_DEFINITION=examples/kubernetes.json`. Not required if `NAME` is provided. |
| `CLEANUP_ON_EXIT` | no      | Delete cluster after running E2E. E.g., `CLEANUP_ON_EXIT=true`. Default is false. |
| `CLEANUP_IF_FAIL` | no      | Delete cluster only if E2E failed. E.g., `CLEANUP_IF_FAIL=false`. Default is false. |
