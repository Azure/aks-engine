# AKS Engine Development and Release Process

## Problem Statement

Today, AKS Engine releases irregularly against master. This irregularity prevents folks from scheduling work that depends upon changes/features in AKS Engine, and instead requires them to manually, passively monitor release activity, and to manually introspect release notes to validate inclusion of desired changes.

## Proposal

Adhere to a regular release cadence:

1. Publish an expected release date in advance
2. Schedule in advance and publish the set of expected work to be included in a release
3. Reserve a regular time interval for release validation + integration

## Three week release interval

We propose to release "every three weeks":

1. Two weeks of planned development
2. One week of testing + integration

## Schedule and Publish release intent in advance

At the beginning of the release interval, we will create a new milestone with at date ~3 weeks in the future:

- https://github.com/Azure/aks-engine/milestones

From the existing work backlog, we will triage open issues and assign an achievable number of prioritized issues to the upcoming release.

## Two weeks of planned development

During the first two weeks of the release cycle, we will burn down the list of items in the release milestone. At the end of this duration, we will transition into release testing + integration

## Testing + Integration

During the final week of the release cycle, we will tag a release candidate commit, and perform the following tests:

- Cluster create validation against the cluster configurations under `test/e2e/test_cluster_configs/`
- Back-compat validation of a "default" cluster configuration built with the prior version; back-compat validation includes cluster scale and upgrade operations against the pre-existing cluster using the new, release candidate build of aks-engine
- Upgrade and scale tests across the supported set of Kubernetes versions using this cluster configuration:
  - `test/e2e/test_cluster_configs/base.json`
- "No egress access" validation against a default cluster configuration that includes an NSG to block outbound internet access:
  - `test/e2e/test_cluster_configs/no_outbound.json`
    - This test validates that the VHD reference in the release candidate build includes all required Kubernetes componentry (pre-downloaded during VHD creation) for the default cluster configuration.
