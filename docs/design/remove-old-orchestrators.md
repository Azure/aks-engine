# Removing Old Orchestrators

This document proposes a staged plan to remove the legacy DCOS, Mesos, Swarm, and SwarmMode
(DockerCE) code from AKS Engine.

## Background

### acs-engine

The AKS Engine project was initially "[acs-engine][]." It was created as a library for the Azure
Container Service when Azure engineers rewrote that resource provider in Go. The needs of this new
"ACSRPv2" effort influenced acs-engine's design and determined what it would support.

When acs-engine began on GitHub ([early 2016][]), it and ACS supported three popular container
orchestrators: DCOS, Mesos, and Swarm. Later SwarmMode and Kubernetes were added.

### ACS Retirement

It wasn't long before the market decided to push all its chips over to Kubernetes. Microsoft also
doubled down by launching the AKS service and announcing an [end-of-life date for ACS][].

acs-engine followed suit by renaming iself to AKS Engine. (We hoped this name change would clarify
things by no longer name-checking the deprecated ACS service, but the overlap with the new, popular
AKS service has been a persistent source of confusion.)

Most mentions of orchestrators other than Kubernetes were removed from the documentation, and some
related code packages were marked as deprecated. However, to support our customers, code for the old
orchestrators remains in AKS Engine until ACS is shut down on January 31, 2020.

### Lessons Already Learned

An existing pull request ([refactor: remove unused orchestrators][]) pointed out a few things that
prompted this design document:

- the scope of this change is much too large to review in one pull request
- localization files have the same old smell
- unit test coverage drops somewhat

## Goals

1. Simplify codebase by removing non-Kubernetes code
2. Stage the removal work into small PRs
3. (Maintain unit test code coverage levels)
4. (Remove stale localization files)

Steps 1 and 2 are self-explanatory. Steps 3 and 4 are hopeful goals that may or may not be met.

## Proposal

1. Remove old orchestrator references in examples and other docs
2. Remove DockerCE/SwarmMode references in code
3. Remove Swarm references in code
4. Remove Mesos references in code
5. Remove DCOS references in code
6. Small refactorings implied by previous changes
7. Remove localization files, scripts, and make targets

The first step ("Remove old orchestrator references in examples and other docs") does not depend on
the ACS service end-of-life date and can be implemented any time.

Steps 2-5 simply break up the removal work into logical units that are humanly reviewable. With each
of these PRs, we will examine unit test coverage to see if it's reasonable to backfill tests to
preserve overall coverage levels, and include such tests with the PR.

Step six indicates that some minor refactorings that were intentionally deferred during previous
steps can be introduced. For example, the naive changes that make steps 2-5 easier to review result
in some empty `switch` or effectively no-op `if` statements. This is the chance to remove such
`if orchestrator == Kubernetes` cruft.

Step seven is a stretch goal, but if it proves straightforward to remove the localization files,
scripts, and related makefile targets, let's do so. Localization was a requirement driven by the ACS
project, and AKS Engine's maintainers consider removal preferable to being perpetually stale.


[acs-engine]: https://github.com/Azure/acs-engine
[early 2016]: https://github.com/Azure/acs-engine/commit/d9bb7f4bbfd1b2932f1d0319e07429eee6127dbf
[end-of-life date for ACS]: https://azure.microsoft.com/en-us/updates/azure-container-service-will-retire-on-january-31-2020/
[refactor: remove unused orchestrators]: https://github.com/Azure/aks-engine/pull/2460
