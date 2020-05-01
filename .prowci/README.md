# Prow

Prow is a CI system that offers various features such as rich Github automation,
and running tests in Jenkins or on a Kubernetes cluster. You can read more about
Prow in [upstream docs][0].

## AKS Engine setup

Deploy a new Kubernetes cluster.

Set up an Ingress controller and a mechanism to do TLS. The [Azure docs][1]
explain how to setup Ingress with TLS on top of a Kubernetes cluster in Azure. (make sure you specify `--set rbac.create=true` when creating the ingress controller)

A Github webhook also needs to be setup in the repo that points to `dns-name/hook`.
`dns-name` is the DNS name setup during the DNS configuration of the Ingress controller.
The Github webhook also needs to send `application/json` type of payloads and use a
secret. This secret is going to be used by Prow to decrypt the payload inside Kubernetes.

Another secret that needs to be setup is a Github token from the bot account that is
going to manage PRs and issues. The token needs the `repo` and `read:org` scopes
enabled. The bot account also needs to be added as a collaborator in the repository
it is going to manage.

To automate the installation of Prow, store the webhook secret as `hmac` and the bot
token as `oauth` inside the `.prowci` directory. Then, installing Prow involves
running the following command:
```
make prow
```

## What is installed

`hook` is installed that manages receiving webhooks from Github and reacting
appropriately on Github. `deck` is installed as the Prow frontend. Last, `tide`
is also installed that takes care of merging pull requests that pass all tests
and satisfy a set of label requirements.

[0]: https://github.com/kubernetes/test-infra/tree/master/prow#prow
[1]: https://docs.microsoft.com/en-us/azure/aks/ingress
