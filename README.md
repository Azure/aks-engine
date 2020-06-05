# AKS Engine - Units of Kubernetes on Azure!

[![Build Status](https://msazure.visualstudio.com/One/_apis/build/status/Custom/Compute/ContainerService/AKS%20Engine%20CI%20E2E?branchName=master)](https://msazure.visualstudio.com/One/_build/latest?definitionId=50661&branchName=master)
[![Coverage Status](https://codecov.io/gh/Azure/aks-engine/branch/master/graph/badge.svg)](https://codecov.io/gh/Azure/aks-engine)
[![GoDoc](https://godoc.org/github.com/Azure/aks-engine?status.svg)](https://godoc.org/github.com/Azure/aks-engine)
[![Go Report Card](https://goreportcard.com/badge/github.com/Azure/aks-engine)](https://goreportcard.com/report/github.com/Azure/aks-engine)

AKS Engine is the easiest way to provision a self-managed Kubernetes cluster on Azure.

## Overview

AKS Engine provides convenient tooling to quickly bootstrap Kubernetes clusters on Azure. By leveraging [ARM (Azure Resource Manager)][ARM], AKS Engine helps you create, destroy and maintain clusters provisioned with basic IaaS resources in Azure. AKS Engine is also the library used by AKS for performing these operations to provide managed service implementations.

## Getting started

Depending on how new you are to AKS Engine, you can try [a tutorial][tutorials], or just dive straight into the [documentation][docs].

Please see the [FAQ][] for answers about AKS Engine and its progenitor ACS-Engine.

## Join the community

Want to get involved? The [community guide][community] covers everything you need to know about the AKS Engine community and how you can contribute. The [developer guide][developer-guide] will help you onboard as a developer.

## Support

AKS Engine is an open source project that is [**not** covered by the Microsoft Azure support policy](https://support.microsoft.com/en-us/help/2941892/support-for-linux-and-open-source-technology-in-azure). [Please search open issues here](https://github.com/Azure/aks-engine/issues), and if your issue isn't already represented please [open a new one](https://github.com/Azure/aks-engine/issues/new/choose). The AKS Engine project maintainers will respond to the best of their abilities.

## Code of conduct

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/). For more information, see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq) or contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.

## Data Collection
The software may collect information about you and your use of the software and send it to Microsoft. Microsoft may use this information to provide services and improve our products and services. You may turn off the telemetry [as described in the repository][telemetry-config]. There are also some features in the software that may enable you and Microsoft to collect data from users of your applications. If you use these features, you must comply with applicable law, including providing appropriate notices to users of your applications together with a copy of Microsoft's privacy statement. Our privacy statement is located at https://go.microsoft.com/fwlink/?LinkID=824704. You can learn more about data collection and use in the help documentation and our privacy statement. Your use of the software operates as your consent to these practices.

For more information, please see the [telemetry documentation][telemetry].

[ARM]: https://docs.microsoft.com/en-us/azure/azure-resource-manager/resource-group-overview
[community]: docs/community/README.md
[developer-guide]: docs/community/developer-guide.md
[docs]: docs/README.md
[FAQ]: docs/faq.md
[tutorials]: docs/tutorials/README.md
[telemetry]: docs/topics/telemetry.md
[telemetry-config]: docs/topics/telemetry.md#configuration
