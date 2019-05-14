# AKS Engine Technical and Feature Debt

The purpose of this document is to outline conspicuous engineering and feature debt as a result of the historical evolution of ACS Engine --> AKS Engine.

## An Ideal Kubernetes on Azure IaaS Toolchain

We will assert that a fully implemented Kubernetes on IaaS solution would include:

- Configurable (both in terms of IaaS and Kubernetes functionality), declarative cluster specification
- Single source of truth for cluster state representation
- Clean, consistent CLI UX
- Clean, consistent programming SDK
- Rapid (comparable to vm deployment SLA) cluster deployment implementation
- Per-component cluster lifecycle configuration management
- In-place Kubernetes version upgrades
- Automatic security updates

## AKS Engine debt

- Incomplete cluster state representation
- High-friction CLI UX (non-trivial amount of command line args required)
- No SDK-like interface(s) for 3rd party code re-use
- Cluster configuration is managed monolithically, all cluster configuration delivered via ARM template
- No in-place Kubernetes version upgrade functionality
