# AAD integration Walkthrough

This walkthrough is to help you get start with Azure Active Directory(AAD) integration with an AKS Engine-created Kubernetes cluster.

[OpenID Connect](http://openid.net/connect/) is a simple identity layer built on top of the OAuth 2.0 protocol, and it is supported by both AAD and Kubernetes. Here we're going to use OpenID Connect as the communication protocol.

Please also refer to [Azure Active Directory plugin for client authentication](https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/client-go/plugin/pkg/client/auth/azure/README.md) in Kubernetes repo for more details about OpenID Connect and AAD support in upstream.

## Prerequisites
1. An Azure Active Directory tenant, referred to as `AAD Tenant`. You can use the tenant for your Azure subscription;
2. A `Web app / API` type AAD application, referred to as `Server Application`. This application represents the `apiserver`. For groups to work properly, you'll need to edit the `Server Application` Manifest and set `groupMembershipClaims` to either `All` or `SecurityGroup`.
3. A `Native` type AAD application, referred to as `Client Application`. This application is for user login via `kubectl`. You'll need to add delegated permission to `Server Application`, please see [troubleshooting](#loginpageerror) section for detail.

You also need to delegate permission to the application as follows:

1. Go to Azure Portal, navigate to `Azure Active Directory` -> `App registrations`.
2. Select the `Client Application`, Navigate to `Settings` -> `Required permissions`
3. Choose `Add`, select the `Server Application`. You may need to enter the Server Application's name into the search field and search for it.
   In permissions tab, select `Delegated permissions` -> `Access {Server Application}`


## Deployment
Follow the [deployment steps](../tutorials/deploy.md). In step #4, add the following under 'properties' section:
```json
"aadProfile": {
    "serverAppID": "",
    "clientAppID": "",
    "tenantID": ""
}
```

- `serverAppID`   : the `Server Application`'s ID
- `clientAppID`   : the `Client Application`'s ID
- `tenantID`      : the `AAD tenant`'s ID

After template generation, the locally generated kubeconfig file (`_output/<instance>/kubeconfig/kubeconfig.<location>.json`) will have the default user using AAD.
Initially it isn't assoicated with any AAD user yet. To get started, try any kubectl command (like `kubectl get pods`), and you'll be prompted to the device login process. After login, you will be able to operate the cluster using your AAD identity.

It should look something like:
```sh
To sign in, use a web browser to open the page https://aka.ms/devicelogin and enter the code FCVDE87XY to authenticate.
```

### Setting up authorization
You can now authenticate to the Kubernetes cluster, but you need to set up authorization as well.

#### Authentication
With AKS Engine, the cluster is locked down by default.

This means that when you try to use your AAD account you will see something
like:
```sh
Error from server (Forbidden): User "https://sts.windows.net/<tenant-id>#<user-id>" cannot list nodes at the cluster scope. (get nodes)
```

See [enabling cluster-admin](#enabling-cluster-admin) below.

#### Enabling cluster admin

To enable authorization, you need to add a cluster admin role account, and add your user to that account.

The user name would be in this format: `IssuerUrl#ObjectID`.

It should be printed in the error message from the previous kubectl request.

Alternately, you can find the `IssuerUrl` under `issuer` property in this url:

```
https://login.microsoftonline.com/<REPLACE_WITH_TENANTID>/.well-known/openid-configuration
```

Once you have the user name you can add it to the `cluster-admin` role (cluster super-user) as follows:

```sh
CLUSTER=<cluster-name-here>
REGION=<your-azure-region-name, e.g. 'centralus'>

ssh -i _output/${CLUSTER}/azureuser_rsa azureuser@${CLUSTER}.${REGION}.cloudapp.azure.com \
    kubectl create clusterrolebinding aad-default-cluster-admin-binding \
        --clusterrole=cluster-admin \
        --user 'https://sts.windows.net/<tenant-id>/#<user-id>'
```

That should output:
```sh
clusterrolebinding "aad-default-cluster-admin-binding" created
```

At which point you should be able to use any Kubernetes commands to administer the cluster, including adding other AAD identities to particular RBAC roles.

#### Enabling AAD groups

You can also optionally add groups into your admin role

For example, if your `IssuerUrl` is `https://sts.windows.net/e2917176-1632-47a0-ad18-671d485757a3/`, and your Group `ObjectID` is `7d04bcd3-3c48-49ab-a064-c0b7d69896da`, the command would be:

```sh
kubectl create clusterrolebinding aad-default-group-cluster-admin-binding --clusterrole=cluster-admin --group=7d04bcd3-3c48-49ab-a064-c0b7d69896da
```

```json
"aadProfile": {
    "serverAppID": "",
    "clientAppID": "",
    "adminGroupID": "7d04bcd3-3c48-49ab-a064-c0b7d69896da"
}
```
The above config will automatically generate a clusterrolebinding with the cluster-admin clusterrole for the specified Group `ObjectID` on cluster deployment via the "aad" addon. See [addons](clusterdefinitions.md#addons) for more info on addons.

#### Adding another client user:
To add another client user run the following:

```sh
kubectl config set-credentials "user1" --auth-provider=azure \
    --auth-provider-arg=environment=AzurePublicCloud \
    --auth-provider-arg=client-id={ClientAppID} \
    --auth-provider-arg=apiserver-id={ServerAppID} \
    --auth-provider-arg=tenant-id={TenantID}
```

And to test that user's login
```sh
kubectl get pods --user=user1
```

Now you'll be prompted to login again, you can try logging in with another AAD user account.
The login would succeed, but later you can see following message since the server denies access:
```
Error from server (Forbidden): User "https://sts.windows.net/{tenantID}/#{objectID}" cannot list pods in the namespace "default". (get pods)
```

You can then update the cluster's role bindings and RBAC to suit your needs for that user. See the [default role bindings](https://kubernetes.io/docs/admin/authorization/rbac/#default-roles-and-role-bindings) for more details, and
the [general guide to Kubernetes RBAC](https://kubernetes.io/docs/admin/authorization/rbac/).

## Troubleshooting

### LoginPageError
If you failed at the login page, you may see following error message
```
Invalid resource. The client has requested access to a resource which is not listed in the requested permissions in the client's application registration. Client app ID: {UUID} Resource value from request: {UUID}. Resource app ID: {UUID}. List of valid resources from app registration: {UUID}.
```
This could be caused by `Client Application` not being authorized.

1. Go to Azure Portal, navigate to `Azure Active Directory` -> `App registrations`.
2. Select the `Client Application`, Navigate to `Settings` -> `Required permissions`
3. Choose `Add`, select the `Server Application`. In permissions tab, select `Delegated permissions` -> `Access {Server Application}`

### ClientError
If you see the following message returned from the server via `kubectl`
```
Error from server (Forbidden)
```

It is usually caused by an incorrect configuration. You could find more debug information in apiserver log. On a master node, run the following command:
```sh
docker logs -f $(docker ps|grep 'hyperkube apiserver'|cut -d' ' -f1) 2>&1 |grep -a auth
```

You might see s message like this:
```
Unable to authenticate the request due to an error: [invalid bearer token, [crypto/rsa: verification error, oidc: JWT claims invalid: invalid claims, 'aud' claim and 'client_id' do not match, aud=UUID1, client_id=spn:UUID2]]
```
This indicates server and client is using different `Server Application` ID, which usually happens when the configurations are being updated manually.

For other auth issues, you may also find some useful information from the log.
