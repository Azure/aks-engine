# Releases

aks-engine uses a [continuous delivery][] approach for creating releases. Every merged commit that passes
testing results in a deliverable that can be given a [semantic version][] tag and shipped.

## Release as Needed

The master `git` branch of a project should always work. Only changes considered ready to be
released publicly are merged.

aks-engine depends on components that release new versions as often as needed. Fixing
a high priority bug requires the project maintainer to create a new patch release.
Merging a backward-compatible feature implies a minor release.

By releasing often, each component release becomes a safe and routine event. This makes it faster
and easier for users to obtain specific fixes. Continuous delivery also reduces the work
necessary to release a product such as aks-engine, which depends on several external projects.

"Components" applies not just to ACS projects, but also to development and release
tools, Kubernetes orchestrator versions, to Docker base images, and to other Azure
projects that do [semantic version][] releases.

## aks-engine Releases Each Month

aks-engine has a regular, public release cadence. From v0.1.0 onward, new aks-engine feature
releases arrive on the first Thursday of each month. Patch releases are created at any time,
as needed. GitHub milestones are used to communicate the content and timing of major and minor
releases, and longer-term planning is visible at [the Roadmap](roadmap.md).

aks-engine release timing is not linked to specific features. If a feature is merged before the
release date, it is included in the next release.

See "[How to Release aks-engine](#how-to-release-aks-engine)" for more detail.

## Semantic Versioning

aks-engine releases comply with [semantic versioning][semantic version], with the "public API" broadly
defined as:

- REST, gRPC, or other API that is network-accessible
- Library or framework API intended for public use
- "Pluggable" socket-level protocols users can redirect
- CLI commands and output formats
- Integration with Azure public APIs such as ARM

In general, changes to anything a user might reasonably link to, customize, or integrate with should
be backward-compatible, or else require a major release. aks-engine users can be confident that upgrading
to a patch or to a minor release will not break anything.

## How to Release aks-engine

This section leads a maintainer through creating an aks-engine release.

### Step 1: Assemble Master Changelog
A change log is a file which contains a curated, chronologically ordered list of changes
for each version of aks-engine, which helps users and contributors see what notable changes
have been made between each version of the project.

The CHANGELOG should be driven by release milestones defined on Github, which track specific deliverables and
work-in-progress.

### Step 2: Manual Testing

Now it's time to go above and beyond current CI tests. Create a testing matrix spreadsheet (copying
from the previous document is a good start) and sign up testers to cover all permutations.

Testers should pay special attention to the overall user experience, make sure upgrading from
earlier versions is smooth, and cover various storage configurations and Kubernetes versions and
infrastructure providers.

When showstopper-level bugs are found, the process is as follows:

1. Create an issue that describes the bug.
1. Create an PR that fixes the bug.
  - PRs should always include tests (unit or e2e as appropriate) to add
 automated coverage for the bug.
1. Once the PR passes and is reviewed, merge it and update the CHANGELOG


### Step 3: Tag and Create a Release

TBD


### Step 4: Close GitHub Milestones

TBD

### Step 5: Let Everyone Know

Let the rest of the team know they can start blogging and tweeting about the new aks-engine release.
Post a message to the #company channel on Slack. Include a link to the released chart and to the
master CHANGELOG:

```
@here aks-engine 0.1.0 is here!
Master CHANGELOG: https://github.com/Azure/aks-engine/CHANGELOG.md
```

You're done with the release. Nice job!

[continuous delivery]: https://en.wikipedia.org/wiki/Continuous_delivery
[semantic version]: http://semver.org
