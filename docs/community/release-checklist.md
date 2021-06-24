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

## AKS Engine Releases As Needed

AKS Engine releases new versions when the team of maintainers determine it is needed. This usually
amounts to one or more releases each month.

Minor versions—for example, v0.**65**.0—are created from the master branch whenever
important features or changes have been merged and CI testing shows it to be stable over time.

Patch versions—for example, v0.64.**1**—are based on the previous release and created on demand
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

We will use **v0.65.0** as an example herein. You should replace this with the new version you're releasing.

### Prepare a Release Branch

First ensure that all the commits to be included in the release are ready in your local repository.

For a major or minor release, create a branch from master. For a patch, create a branch from the previous release tag and use `git cherry-pick` to apply specific commits. Ensure that the CHANGELOG file that corresponds to this release is present in the release branch.

```sh
$ git checkout master
Switched to branch 'master'
Your branch is up to date with 'origin/master'.
$ git fetch upstream
$ git merge upstream/master
Already up to date.
$ export RELEASE=v0.65.0 && git checkout -b release-$RELEASE && git push upstream release-$RELEASE
Switched to a new branch 'release-v0.65.0'
Total 0 (delta 0), reused 0 (delta 0), pack-reused 0
remote:
remote: Create a pull request for 'release-v0.65.0' on GitHub by visiting:
remote:      https://github.com/Azure/aks-engine/pull/new/release-v0.65.0
remote:
To https://github.com/Azure/aks-engine.git
 * [new branch]          release-v0.65.0 -> release-v0.65.0
```

### Generate Release Notes

First, create a branch on your fork that you will use to push a release notes PR for approval by the maintainers:

```sh
$ git checkout -b CHANGELOG-$RELEASE && git push origin CHANGELOG-$RELEASE
Switched to a new branch 'CHANGELOG-v0.65.0'
Total 0 (delta 0), reused 0 (delta 0), pack-reused 0
remote:
remote: Create a pull request for 'CHANGELOG-v0.65.0' on GitHub by visiting:
remote:      https://github.com/jackfrancis/aks-engine/pull/new/CHANGELOG-v0.65.0
remote:
To https://github.com/jackfrancis/aks-engine.git
 * [new branch]          CHANGELOG-v0.65.0 -> CHANGELOG-v0.65.0
```

Now, create a temporary local tag that correlates with the release version to allow `git-chglog` to create human-readable release notes:

```sh
$ export RELEASE=v0.65.0 && git tag $RELEASE && git-chglog --tag-filter-pattern 'v\d+\.\d+\.\d+$' --output releases/CHANGELOG-$RELEASE.md $RELEASE && git tag -d $RELEASE
```

You may now review the generated release notes, and possibly add some manual curation if appropriate. After the CHANGELOG looks good, push the changes to your fork + branch:

```sh
$ git add releases/CHANGELOG-$RELEASE.md
$ git commit -m "release: $RELEASE CHANGELOG"
[CHANGELOG-v0.65.0 f8d86fbdb] release: v0.65.0 CHANGELOG
 1 file changed, 31 insertions(+)
 create mode 100644 releases/CHANGELOG-v0.65.0.md
$ git push --set-upstream origin CHANGELOG-$RELEASE
Enumerating objects: 6, done.
Counting objects: 100% (6/6), done.
Delta compression using up to 8 threads
Compressing objects: 100% (4/4), done.
Writing objects: 100% (4/4), 1000 bytes | 111.00 KiB/s, done.
Total 4 (delta 2), reused 0 (delta 0), pack-reused 0
remote: Resolving deltas: 100% (2/2), completed with 2 local objects.
To https://github.com/jackfrancis/aks-engine.git
   e6ca055c4..f8d86fbdb  CHANGELOG-v0.65.0 -> CHANGELOG-v0.65.0
Branch 'CHANGELOG-v0.65.0' set up to track remote branch 'CHANGELOG-v0.65.0' from 'origin'.
```

And now create a PR and get a couple maintainers to review and lgtm.

### Add CHANGELOG to release branch

After the CHANGELOG PR merges, you now want to add it to your release branch to enable the automated release CI to progress.

```
$ git checkout release-$RELEASE
Switched to branch 'release-v0.65.0'
$ git cherry-pick <commit SHA of CHANGELOG commit in master branch>
```

Tag the release commit and push it to GitHub:

```
$ git tag $RELEASE && git push upstream $RELEASE
```

### Automated Release CI

When you push a new tag that matches the pattern `v*.*.*`, a GitHub Actions job will run automatically and create a new release from that tag, build and publish release artifacts, and populate the release body with the CHANGELOG created earlier for this release. Before actually publishing the release a series of release-gating E2E scenarios will run. It will take 2-3 hours for the entire process to complete.

### Update Package Managers

Finally, let's make the new aks-engine release easy to install.

#### The `brew` package manager

Create a pull request to add the new release to [brew][] through our [homebrew tap repository][brew-tap]. Update the macOS sha256 checksum with this command:

```
$ shasum -a 256  _dist/aks-engine-$RELEASE-darwin-amd64.tar.gz
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
