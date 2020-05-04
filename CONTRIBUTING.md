# Contributing Guidelines

The Microsoft AKS Engine project accepts contributions via GitHub pull requests. This document outlines the process to help get your contribution accepted.

Please see also the [AKS Engine Developer Guide](docs/community/developer-guide.md).

## Contributor License Agreements

We'd love to accept your patches! Before we can take them, we have to jump a
couple of legal hurdles.

The [Microsoft CLA](https://cla.microsoft.com/) must be signed by all contributors. Please fill out either the individual or corporate Contributor License Agreement (CLA). Once you are CLA'ed, we'll be able to accept your pull requests.

***NOTE***: Only original source code from you and other people that have
signed the CLA can be accepted into the repository.


## Support Channels

This is an open source project and as such no formal support is available. However, like all good open source projects we do offer "best effort" support through github issues.

GitHub issues:
- AKS Engine: https://github.com/Azure/aks-engine/issues - file issues and PRs related to AKS Engine
- AKS: https://github.com/Azure/AKS/issues - file issues and PRs related to Azure Kubernetes Service

Before opening a new issue or submitting a new pull request, it's helpful to search the project - it's likely that another user has already reported the issue you're facing, or it's a known issue that we're already aware of.

## Milestones
We use milestones to track progress of releases.

For example, if the current version is `2.2.0` an issue/PR could fall in to one of 2 different active milestones:
`2.2.1`, `2.3.0`.  If an issue pertains to a
specific upcoming bug or minor release, it would go into `2.2.1` or `2.3.0`.

A milestone (and hence release) is considered done when all outstanding issues/PRs have been closed or moved to another milestone.

## Issues
Issues are used as the primary method for tracking anything to do with the AKS Engine project.

### Issue Lifecycle
The issue lifecycle is mainly driven by the core maintainers, but is good information for those
contributing to AKS Engine. All issue types follow the same general lifecycle. Differences are noted below.
1. Issue creation
2. Triage
    - The maintainer in charge of triaging will apply the proper labels for the issue. This
    includes labels for priority, type, and metadata (such as "orchestrator/k8s"). If additional
    levels are needed in the future, we will add them.
    - (If needed) Clean up the title to succinctly and clearly state the issue. Also ensure
    that proposals are prefaced with "Proposal".
    - Add the issue to the correct milestone. If any questions come up, don't worry about
    adding the issue to a milestone until the questions are answered.
    - We attempt to do this process at least once per work day.
3. Discussion
    - "Feature" and "Bug" issues should be connected to the PR that resolves it.
    - Whoever is working on a "Feature" or "Bug" issue (whether a maintainer or someone from
    the community), should either assign the issue to them self or make a comment in the issue
    saying that they are taking it.
    - "Proposal" and "Question" issues should stay open until resolved or if they have not been
    active for more than 30 days. This will help keep the issue queue to a manageable size and
    reduce noise. Should the issue need to stay open, the `keep open` label can be added.
4. Issue closure

## How to Contribute a Patch

1. If you haven't already done so, sign a Contributor License Agreement (see details above).
2. Fork the desired repo, develop and test your code changes.
3. Submit a pull request.
