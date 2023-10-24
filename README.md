# AKS Engine - Deprecated tool for self-managed Kubernetes on Azure

## Project status

This project is deprecated for Azure public cloud customers. Please use [Azure Kubernetes Service (AKS)](https://azure.microsoft.com/en-us/services/kubernetes-service/#overview) for managed Kubernetes or [Cluster API Provider Azure](https://github.com/kubernetes-sigs/cluster-api-provider-azure) for self-managed Kubernetes. There are no further releases planned; Kubernetes 1.24 was the final version to receive updates.

For use on the [Azure Stack Hub product](https://docs.microsoft.com/en-us/azure-stack/user/azure-stack-kubernetes-aks-engine-overview) this project is fully supported and will continue to be supported by the Hub team throughout the lifespan of Azure Stack Hub. Development is already moved to a new Azure Stack Hub specific repository ([Azure/aks-engine-azurestack](https://github.com/Azure/aks-engine-azurestack)). This new repository is where new [releases](https://github.com/Azure/aks-engine-azurestack/releases) for Azure Stack Hub clouds, starting at v0.75.3, will be published and where [issues](https://github.com/Azure/aks-engine-azurestack/issues/new) concerning Azure Stack Hub should be created.


## Support

Please see our [support policy][support-policy].

## Code of conduct

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/). For more information, see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq) or contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.

## Data Collection
The software may collect information about you and your use of the software and send it to Microsoft. Microsoft may use this information to provide services and improve our products and services. You may turn off the telemetry [as described in the repository][telemetry-config]. There are also some features in the software that may enable you and Microsoft to collect data from users of your applications. If you use these features, you must comply with applicable law, including providing appropriate notices to users of your applications together with a copy of Microsoft's privacy statement. Our privacy statement is located at https://go.microsoft.com/fwlink/?LinkID=824704. You can learn more about data collection and use in the help documentation and our privacy statement. Your use of the software operates as your consent to these practices.

For more information, please see the [telemetry documentation][telemetry].

[support-policy]: SUPPORT.md
[telemetry]: docs/topics/telemetry.md
[telemetry-config]: docs/topics/telemetry.md#configuration
