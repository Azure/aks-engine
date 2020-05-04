# Releases

AKS Engine uses a [continuous delivery][] approach for creating releases. Every merged commit that passes
testing results in a deliverable that can be given a [semantic version][] tag and shipped.

## Master Is Always Releasable

The master `git` branch of a project should always work. Only changes considered ready to be
released publicly are merged.

AKS Engine depends on components that release new versions as often as needed. Fixing
a high priority bug requires the project maintainer to create a new patch release.
Merging a backward-compatible feature implies a minor release.

By releasing often, each release becomes a safe and routine event. This makes it faster
and easier for users to obtain specific fixes. Continuous delivery also reduces the work
necessary to release a product such as AKS Engine, which depends on several external projects.

"Components" applies not just to AKS projects, but also to development and release
tools, to orchestrator versions, to Docker base images, and to other Azure
projects that do [semantic version][] releases.

## AKS Engine Releases As Needed

AKS Engine releases new versions when the team of maintainers determine it is needed. This usually
amounts to one or more releases each month.

Minor versions—for example, v0.**32**.0—are created from the master branch whenever
important features or changes have been merged and CI testing shows it to be stable over time.

Patch versions—for example, v0.32.**3**—are based on the previous release and created on demand
whenever important bug fixes arrive.

See "[Creating a New Release](#creating-a-new-release)" for more detail.

## Semantic Versioning

Releases of the `aks-engine` binary comply with [semantic versioning][semantic version], with the "public API" broadly
defined as:

- REST, gRPC, or other API that is network-accessible
- Library or framework API intended for public use
- "Pluggable" socket-level protocols users can redirect
- CLI commands and output formats
- Integration with Azure public APIs such as ARM

In general, changes to anything a user might reasonably link to, customize, or integrate with should
be backward-compatible, or else require a major release. `aks-engine` users can be confident that upgrading
to a patch or to a minor release will not break anything.

## Creating a New Release

Let's go through the process of creating a new release of the [aks-engine][] binary.

We will use **v0.32.3** as an example herein. You should replace this with the new version you're releasing.

```
$ export TAG=v0.32.3
```

### Prepare and Tag a Branch

First ensure that all the commits to be included in the release are ready in your local repository.

For a major or minor release, create a branch from master. For a patch, create a branch from the previous release tag and use `git cherry-pick` to apply specific commits.

Tag the release commit and push it to GitHub:

```
$ git tag $TAG && git push upstream $TAG
```

### Generate Release Notes

Use the [`git-chglog`][git-chglog] tool to generate release notes:

```
$ git-chglog $TAG
```

Be sure to proofread the output and verify that the intended commits appear. If a commit made it to master that didn't have a [conventional commit message][conventional-commit], you'll need to add it to the appropriate section by hand.

Save the markdown that it prints so it can be pasted into the GitHub release.

### Generate Download Artifacts

Make sure your repository has no local changes, then build the aks-engine distribution archives:

```
$ make generate info  # check that the git tree state is clean after a build, and that the tag is correct
$ make clean dist
```

When this finishes, the `_dist` directory will be populated with three .zip and three .tar.gz archive files.

### Make a GitHub Release

Now navigate to the aks-engine project on GitHub and start a [new release][new-release]:

![draft_new_release.png](../static/img/draft_new_release.png)

Select the tag we pushed previously, and use that tag as the release title. Then paste the release notes from the previous step into the big text field:

![release_notes.png](../static/img/release_notes.png)

Finally, drag all six archive files we created with `make dist` into the "Attach binaries" field at the bottom of the release web form:

![attach_archives.png](../static/img/attach_archives.png)

Proofread the release notes and satisfy yourself that everything is in order. Click the "Publish release" button when the new aks-engine release is ready for the world.

### Update Package Managers

Finally, let's make the new aks-engine release easy to install.

#### The `brew` package manager

Create a pull request to add the new release to [brew][] through our [homebrew tap repository][brew-tap]. Update the macOS sha256 checksum with this command:

```
$ shasum -a 256  _dist/aks-engine-$TAG-darwin-amd64.tar.gz
```

The PR will look very similar to [this one][brew-pr].

#### The `gofish` package manager

The [gofish][] package manager has automation in place to create an update when AKS Engine creates a release. Check the [fish-food repository][gofish-food] to see that a pull request was created.

#### The `choco` package manager

Adding new versions to [choco][] is automated, but you can check the status of package approval and publishing at the [aks-engine chocolatey page][choco-status].


[aks-engine]: https://github.com/Azure/aks-engine/releases
[brew]: https://brew.sh/
[brew-pr]: https://github.com/Azure/homebrew-aks-engine/pull/34
[brew-tap]: https://github.com/Azure/homebrew-aks-engine/
[choco]: https://chocolatey.org/
[choco-status]: https://chocolatey.org/packages/aks-engine/
[conventional-commit]: https://www.conventionalcommits.org/en/v1.0.0-beta.3/
[git-chglog]: https://github.com/git-chglog/git-chglog
[gofish]: https://github.com/fishworks/gofish
[gofish-food]: https://github.com/fishworks/fish-food/
[gofish-pr]: https://github.com/fishworks/fish-food/pull/141
[new-release]: https://github.com/Azure/aks-engine/releases/new
[continuous delivery]: https://en.wikipedia.org/wiki/Continuous_delivery
[semantic version]: http://semver.org
