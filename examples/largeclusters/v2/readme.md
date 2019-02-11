

#Steps 

##Create Resource Group

```
az group create --name <<resource_group_name>> --location <<azure_location>>
```


##Create vnet and etcd subnet

> You only need to do that if you are creating the cluster in a non exitent vnet
```
#address prefix example: 10.0.0.0/16 (256^2 - 4)
#subnet address prefix: 10.0.1.0/24 (256 - 4)
az network vnet create --resource-group <<resource_group_name>> \
                       --name <<vnet_name>> \
                       --address-prefix <<vnet address space>> \
                       --subnet-name <<etcd_subnet_name>>  \
                       --subnet-prefix <<etcd_subnet_address_space>>
```


> Depending on your cluster size you may need to create 2 or more etcd cluster. For exampl an etcd cluster for `events` another for everything else.

## Create etcd Cluster(s)

```
# base cluster name used as a prefix for vm name
# default admin user is azureuser
# default ssh key value is $(cat ~/.ssh/id_rsa.pub)

RG="<<resource_group_name>>" VNET_NAME="<<vnet_name>>" SUBNET_NAME="<<subnet_name>>" BASE_CLUSTER_NAME=<<base_cluster_name>> ADMIN_USER="<<override_admin_username>>" SSH_KEY_VALUE="$(cat <<custom key path>>)" ./examples/largeclusters/v2/create-etcd.sh
```  
