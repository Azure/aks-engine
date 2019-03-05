# Developer Guide

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

## Build AKS Engine from Source

### Docker Development Environment

The easiest way to start hacking on AKS Engine is to use a Docker-based environment. If you already have Docker installed then you can get started with a few commands.

```console
$ make dev
```

This make target mounts the `aks-engine` source directory as a volume into the Docker container, which means you can edit your source code in your favorite editor on your machine, while still being able to compile and test inside of the Docker container. This environment mirrors the environment used in the AKS Engine continuous integration (CI) system.

When `make dev` completes, you will be left at a command prompt inside a Docker container.

Run the following commands to pull the latest dependencies and build the `aks-engine` tool.

```sh
# install and download build dependencies
make bootstrap
# build the `aks-engine` binary
make build
```

The build process leaves the compiled `aks-engine` binary in the `bin` directory. Make sure everything completed successfully by running `bin/aks-engine` without any arguments:

```sh
$ ./bin/aks-engine
AKS Engine deploys and manages Kubernetes, Swarm Mode, and DC/OS clusters in Azure

Usage:
  aks-engine [command]

Available Commands:
  deploy        Deploy an Azure Resource Manager template
  generate      Generate an Azure Resource Manager template
  help          Help about any command
  orchestrators Display info about supported orchestrators
  scale         Scale an existing Kubernetes cluster
  upgrade       Upgrade an existing Kubernetes cluster
  version       Print the version of AKS Engine

Flags:
      --debug   enable verbose debug logs
  -h, --help    help for aks-engine

Use "aks-engine [command] --help" for more information about a command.
```

[Here is a reference to the information on Kubernetes cluster upgrade.](https://github.com/Azure/aks-engine/blob/master/examples/k8s-upgrade/README.md)

[Here's a quick demo video showing the dev/build/test cycle with this setup.](https://www.youtube.com/watch?v=lc6UZmqxQMs)

### Building on Windows, OSX, and Linux

Building AKS Engine from source has a few requirements for each of the platforms. Download and install the prerequisites for your platform: Windows, Linux, or Mac:

#### Windows

Setup steps:

- Setup your go workspace. This guide assumes you are using `c:\Users\me\code\go` as your Go workspace:
  1. Type WIN+R to open the run prompt
  1. Type `rundll32 sysdm.cpl,EditEnvironmentVariables` to open the system variables
  1. Add `c:\go\bin` and `c:\Users\me\code\go\bin` to your PATH variables
  1. Click "new" and add new environment variable named `GOPATH` and set the value to `c:\Users\me\code\go`

- Build aks-engine:
  1. Type Windows key-R to open the run prompt
  1. Type `cmd` to open a command prompt
  1. Type `mkdir %GOPATH%` to create your gopath
  1. Type `cd %GOPATH%`
  1. Type `mkdir -p src\github.com\Azure` to create the gopath to aks-engine
  1. Type `cd src\github.com\Azure`
  1. Type `git clone https://github.com/Azure/aks-engine` to download aks-engine from GitHub
  1. Type `cd aks-engine`
  1. Type `make bootstrap` to get the supporting components
  1. Type `make` to build the project
  1. Run `.\bin\aks-engine.exe` to see the command line parameters

#### OS X and Linux

Setup steps:

  1. Open a command prompt to setup your gopath:
  1. `mkdir $HOME/go`
  1. Edit `$HOME/.bash_profile` and add the following lines to setup your go path:
      ```sh
      export GOPATH=$HOME/go
      export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
      ```
  1. `source $HOME/.bash_profile`

Build aks-engine:

  1. Type `mkdir -p $HOME/go/src/github.com/Azure` to create your gopath
  1. Type `cd $_` to switch directories to that same path
  1. Type `git clone https://github.com/Azure/aks-engine` to download aks-engine from GitHub
  1. Type `cd aks-engine` to change to the source directory
  1. Type `make bootstrap` to install supporting components
  1. Type `make` to build the project
  1. Type `./bin/aks-engine` to see the command line parameters

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

End-to-end tests for Kubernetes may be run
via `make test-kubernetes`.  The test process can optionally
deploy and tear down a cluster as part of the test (this is enabled by default).
You'll need access to an Azure subscription, as well as at least the following
environment variables to be set:

* `CLIENT_ID`: "name" field (a URL) from an Azure service principal
* `CLIENT_SECRET`: "password" field from an Azure service principal
* `SUBSCRIPTION_ID`: Azure subscription UUID
* `TENANT_ID`: Azure tenant UUID

The end-to-end tests also require the `k` script from the `scripts/` folder in to
be in your search $PATH. This ensures that testing uses a `kubectl` client that
matches the version of the Kubernetes server.

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

## Test pipeline

AKS Engine employs a Continuous Integration (CI) system that incorporates Azure DevOps, configured to interact with the AKS Engine GitHub project.

The following steps constitute the AKS Engine CI pipeline:

1. Contributor opens a Pull Request (PR) against the AKS Engine project
1. An AKS Engine team member comments on the PR to trigger an Azure DevOps job that
   + applies the changes to the HEAD of the master branch
   + runs unit tests and code coverage reports
   + generates multiple ARM templates for different deployment scenarios
   + simultaneously provisions the clusters based on generated templates in Azure
1. The PR is code reviewed by the members of AKS Engine team
1. Once the PR is approved and the end-to-end job has passed, the PR can now be merged into the master branch
1. Once merged, another job is triggered to verify integrity of the master branch. This job is similar to the PR job.
