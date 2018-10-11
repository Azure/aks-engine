# Microsoft Azure Kubernetes Service Engine - Kubernetes

* Create a Kubernetes Cluster
  * [Linux](kubernetes/deploy.md) - Create your first Linux Kubernetes cluster
  * [Windows](kubernetes/windows.md) - Create your first Windows Kubernetes cluster
* [Kubernetes Next Steps](kubernetes/walkthrough.md) - You have successfully deployed a Kubernetes cluster, now what?
* [Troubleshooting](kubernetes/troubleshooting.md) - Running into issues? Start here to troubleshoot Kubernetes.
* [Features](kubernetes/features.md) - Guide to alpha, beta, and stable functionality in aks-engine.
* [For Kubernetes Developers](kubernetes/k8s-developers.md) - Info for devs working on Kubernetes upstream and wanting to test using aks-engine.

## Known Issues

### Node "NotReady" due to lost TCP connection

Nodes might appear in the "NotReady" state for approx. 15 minutes if master stops receiving updates from agents.
This is a known upstream kubernetes [issue #41916](https://github.com/kubernetes/kubernetes/issues/41916#issuecomment-312428731). This fixing PR is currently under review.

AKS Engine partially mitigates this issue on Linux by detecting dead TCP connections more quickly via **net.ipv4.tcp_retries2=8**.

## Additional Kubernetes Resources

Here are recommended links to learn more about Kubernetes:

1. [Kubernetes Bootcamp](https://kubernetesbootcamp.github.io/kubernetes-bootcamp/index.html) - shows you how to deploy, scale, update, and debug containerized applications.
2. [Kubernetes Userguide](http://kubernetes.io/docs/user-guide/) - provides information on running programs in an existing Kubernetes cluster.
3. [Kubernetes Examples](https://github.com/kubernetes/examples) - provides a number of examples on how to run real applications with Kubernetes.
