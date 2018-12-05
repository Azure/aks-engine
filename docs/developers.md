# Developers Guide

This guide explains how to set up your environment for developing on
aks-engine.

## Prerequisites

- [Go](https://golang.org/dl) 1.11 or later
- Golang [dep](https://github.com/golang/dep) 0.5.0 or later
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) 1.7 or later
- An Azure account (needed for deploying VMs and Azure infrastructure)
- Git

## Contribution Guidelines

We welcome contributions. This project has set up some guidelines in
order to ensure that (a) code quality remains high, (b) the project
remains consistent, and (c) contributions follow the open source legal
requirements. Our intent is not to burden contributors, but to build
elegant and high-quality open source code so that our users will benefit.

Make sure you have read and understood the main CONTRIBUTING guide:

https://github.com/Azure/aks-engine/blob/master/CONTRIBUTING.md

### Structure of the Code

The code for the aks-engine project is organized as follows:

- The individual programs are located in `cmd/`. Code inside of `cmd/`
  is not designed for library re-use.
- Shared libraries are stored in `pkg/`.
- The `tests/` directory contains a number of utility scripts. Most of these
  are used by the CI/CD pipeline.
- The `docs/` folder is used for documentation and examples.

Go dependencies are managed with
[Golang Dep](https://github.com/golang/dep) and stored in the
`vendor/` directory.

### Git Conventions

We use Git for our version control system. The `master` branch is the
home of the current development candidate. Releases are tagged.

We accept changes to the code via GitHub Pull Requests (PRs). One
workflow for doing this is as follows:

1. Use `go get` to clone the aks-engine repository: `go get github.com/Azure/aks-engine`
2. Fork that repository into your GitHub account
3. Add your repository as a remote for `$GOPATH/github.com/Azure/aks-engine`
4. Create a new working branch (`git checkout -b feat/my-feature`) and
   do your work on that branch.
5. When you are ready for us to review, push your branch to GitHub, and
   then open a new pull request with us.

### Third Party Dependencies

Third party dependencies reside locally inside the repository under the `vendor/` directory. We use [dep](https://github.com/golang/dep) to enforce our dependency graph, declared in [Gopkg.toml](https://github.com/Azure/aks-engine/blob/master/CONTRIBUTING.md) in the project root.

If you wish to introduce a new third party dependency into `aks-engine`, please file an [issue](https://github.com/Azure/aks-engine/issues), and include the canonical VCS path (e.g., `github.com/Azure/azure-sdk-for-go`) along with either the desired release string expression to depend on (e.g., `~8.1.0`), or the commit hash to pin to a static commit (e.g., `4cdb38c072b86bf795d2c81de50784d9fdd6eb77`). A project maintainer will then own the effort to update the codebase with that dependency, including relevant updates to `Gopkg.toml` and `vendor/`.

As a rule we want to distinguish dependency update PRs from feature/bug PRs; we may ask that feature/bug PRs which include updates to `vendor/` and/or contain any other dependency-related overhead to be triaged into separate PRs that can be managed independently, pre-requisite dependency changes in one, and features/bugs in another. The objective of enforcing these distinctions is to help focus the PR review process, and to make manageable the difficult task of rationalizing a multitude of parallel PRs in flight, many of which which may carry hard-to-reconcile dependency side-effects when aggressively updated with a fresh dependency graph as part of the PR payload.

### Go Conventions

We follow the Go coding style standards very closely. Typically, running
`go fmt` will make your code beautiful for you.

We also typically follow the conventions recommended by `go lint` and
`gometalinter`. Run `make test-style` to test the style conformance.

Read more:

- Effective Go [introduces formatting](https://golang.org/doc/effective_go.html#formatting).
- The Go Wiki has a great article on [formatting](https://github.com/golang/go/wiki/CodeReviewComments).

### Unit Tests

Unit tests may be run locally via `make test`.

### End-to-end Tests

End-to-end tests for the DCOS and Kubernetes orchestrators may be run
via `make test-{dcos,kubernetes}`.  The test process can optionally
deploy and tear down a cluster as part of the test (this is enabled by default).
You'll need access to an Azure subscription, as well as at least the following
environment variables to be set:

* `CLIENT_ID`: "name" field (a URL) from an Azure service principal
* `CLIENT_SECRET`: "password" field from an Azure service principal
* `SUBSCRIPTION_ID`: Azure subscription UUID
* `TENANT_ID`: Azure tenant UUID

### Debugging

For aks-engine code debugging you can use [Delve](https://github.com/derekparker/delve) debugger.

#### CLI

Run command:
```
dlv debug github.com/Azure/aks-engine -- generate ~/Documents/azure/kubernetes.json
```

Test individual package and individual test:
```
dlv test github.com/Azure/aks-engine/pkg/engine
dlv test github.com/Azure/aks-engine/pkg/engine -- -test.run ^TestNetworkPolicyDefaults$
```

#### Visual Code Studio

More on VSC integration with delve can be found [here](https://github.com/Microsoft/vscode-go/wiki/Debugging-Go-code-using-VS-Code)

If delve is installed and configured, you can use native VS functionality to debug code or individual tests (`debug test`)

Example launch.json file:

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch",
      "type": "go",
      "request": "launch",
      "mode": "debug",
      "program": "${workspaceRoot}",
      "env": {},
      "args": ["generate", "${workspaceRoot}/examples/kubernetes.json"],
      "showLog": true
    }
  ]
}
```
