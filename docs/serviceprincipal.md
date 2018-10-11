# Microsoft Azure Kubernetes Service Engine

## Service Principals

### Overview

Service Accounts in Azure are tied to Active Directory Service Principals. You can read more about
Service Principals and AD Applications: ["Application and service principal objects in Azure Active Directory"](https://azure.microsoft.com/en-us/documentation/articles/active-directory-application-objects/).

Kubernetes uses a Service Principal to talk to Azure APIs to dynamically manage
resources such as
[User Defined Routes](https://azure.microsoft.com/en-us/documentation/articles/virtual-networks-udr-overview/)
and [L4 Load Balancers](https://azure.microsoft.com/en-us/documentation/articles/load-balancer-overview/).

### Creating a Service Principal


There are several ways to create a Service Principal in Azure Active Directory:

* **With the [Azure CLI](https://github.com/Azure/azure-cli)**

  * Subscription level scope
     ```shell
     az login
     az account set --subscription="${SUBSCRIPTION_ID}"
     az ad sp create-for-rbac --role="Contributor" --scopes="/subscriptions/${SUBSCRIPTION_ID}"
     ```
  * Resource group level scope
     ```shell
     az login
     az account set --subscription="${SUBSCRIPTION_ID}"
     az ad sp create-for-rbac --role="Contributor" --scopes="/subscriptions/${SUBSCRIPTION_ID}/resourceGroups/${RESOURCE_GROUP_NAME}"
     ```

   This will output your `appId`, `password`, `name`, and `tenant`.  The `name` or `appId` may be used for the `servicePrincipalProfile.clientId` and the `password` is used for `servicePrincipalProfile.secret`.

   Confirm your service principal by opening a new shell and run the following commands substituting in `name`, `password`, and `tenant`:

   ```shell
   az login --service-principal -u NAME -p PASSWORD --tenant TENANT
   az vm list-sizes --location westus
   ```

* **With [PowerShell](https://github.com/Azure/azure-powershell)**

   Instructions: ["Use Azure PowerShell to create a service principal to access resources"](https://azure.microsoft.com/en-us/documentation/articles/resource-group-authenticate-service-principal/)

   To get you started quickly, the following are simplified instructions for creating a single-tenant AD application and a service principal with password authentication. Please read the full instructions above for proper RBAC setup of your application. Display name and URI are a friendly arbitrary name and address for your application.

   ```powershell
   PS> Login-AzureRmAccount -SubscriptionId $subscriptionId
   PS> $app = New-AzureRmADApplication -DisplayName $name -IdentifierUris $uri -Password $passwd
   PS> New-AzureRmADServicePrincipal -ApplicationId $app.ApplicationId
   PS> New-AzureRmRoleAssignment -RoleDefinitionName Contributor -ServicePrincipalName $app.ApplicationId
   ```

   The first command outputs your `tenantId`, used below. The `$app.ApplicationId` is used for the `servicePrincipalProfile.clientId` and the `$passwd` is used for `servicePrincipalProfile.secret`.

   Confirm your service principal by opening a new PowerShell session and running the following commands. Enter `$app.ApplicationId` for username.

   ```powershell
   PS> $creds = Get-Credential
   PS> Login-AzureRmAccount -ServicePrincipal -TenantId $tenantId -Credential $creds
   PS> Get-AzureRmVMSize -Location westus
   ```

* **With the [Portal](https://portal.azure.com)**

   Instructions: ["Use portal to create Active Directory application and service principal that can access resources"](https://azure.microsoft.com/en-us/documentation/articles/resource-group-create-service-principal-portal/)
