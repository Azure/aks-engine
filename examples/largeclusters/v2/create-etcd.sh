#!/bin/bash

set -eo pipefail

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Expected input
#SUB_ID							subscription-id									optional(default what az is configured to use)
#RG									resource_group 									mandatory
#VNET_NAME					vnet name												mandatory
#SUBNET_NAME				subnet name											mandatory
#BASE_CLUSTER_NAME	basename for cluster resourcs		optional(default random)
#CLUSTER_SIZE				etcd cluster size								optional(default 3)
#VM_SIZE						etcd vm size										optional(default: Standard_DS2_v2)
#OUTPUT_DIR					locatio of generated certs			optional(default /tmp/${BASE_CLUSTER_NAME}


# Global Vars
AV_SET_NAME=""
ADMIN_USER=${ADMIN_USER:-azureuser}
SSH_KEY_VALUE=${SSH_KEY_VALUE:-$(cat ~/.ssh/id_rsa.pub)}


# Colors
Color_Off='\033[0m'

Red='\033[0;31m'          # Red
Green='\033[0;32m'        # Green
Yellow='\033[0;33m'       # Yellow


inf(){
	local msg="$1"
	echo -e "${Green}INF:${msg}"	
}

wrn(){
	local msg="$1"
	echo -e "${Yellow}WRN:${msg}"
}

err(){
	local msg="$1"
	echo -e "${Red}ERR:${msg}"
	exit 1
}

#creates random string (starts with letter)
get_random(){
	local length="$1"
	if [[ -z "${length}" ]] || [[ "0" -eq "$lenght" ]]; then
		length="8" #default
	fi
	length=$((length - 1))
	echo -n "r"
	#cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w ${length} # | head -n 1
	cat /dev/urandom | tr -dc 'a-z0-9' | fold -w ${length} | head -n 1 || true
}

# gets current sub id
get_current_sub_id(){
	az account list -o table | grep True | awk '{print $6}'
}
# validate input and default as needed
validate_and_default(){
	inf "Validating.."

	if [[ -z "${SUB_ID}" ]]; then
	  SUB_ID="$(az account list --all -o table | grep True | awk '{print $6}')"
		# no sub id?
		if [[ -z "${SUB_ID}" ]];then
				err "no sub id was provided, and no sub id is set with az cli"
			else
				wrn "Subscription was not set, using currently set ${SUB_ID}"
		fi
		SUB_ID="${SUB_ID}"
	else
		# set it
			az account set -s  ${SUB_ID}
	fi	


	if [[ -z "${RG}" ]]; then
		err "Resource Group[RG] was not set"
	fi

	if [[ -z "${VNET_NAME}" ]]; then
		err "VNET name [VNET_NAME] was not set"
	fi

	if [[ -z "${SUBNET_NAME}" ]];	 then
		err "Subnet name [SUBNET_NAME] was not set"
	fi	

	if [[ -z "${BASE_CLUSTER_NAME}" ]]; then
		BASE_CLUSTER_NAME="$(get_random)"
		wrn "Cluster base name [CLUSTER_BASE_NAME] was not set, randmozing to ${BASE_CLUSTER_NAME}"
	fi
	if [[ -z "${OUTPUT_DIR}" ]]; then
			OUTPUT_DIR="/tmp/${BASE_CLUSTER_NAME}"
			wrn "output directory was not set, defaulting to ${OUTPUT_DIR}"
	fi
	mkdir -p ${OUTPUT_DIR}

	if [[ -z "${CLUSTER_SIZE}" ]]; then
		wrn "cluster size [CLUSTER_SIZE] was not set, defaulting to 3"
		CLUSTER_SIZE=3
	fi

	AV_SET_NAME="${BASE_CLUSTER_NAME}_avset"
	inf "Availability set name: ${AV_SET_NAME}"
}

create_vms(){
	local current=0
	local vmName=""
	inf "Creating avset:${AV_SET_NAME} in RG:${RG}"
	az vm availability-set create --resource-group "${RG}" --name  "${AV_SET_NAME}"

	inf "Creating cluster with size: ${CLUSTER_SIZE}"
	while [  $current -lt ${CLUSTER_SIZE} ]; do
		vmName="${BASE_CLUSTER_NAME}${current}"
		inf "Creating VM:${vmName} RG ${RG} AVSET:${AV_SET_NAME}"
		# We create a VM with large premium disks
		# to get the highest possible I/O throughput 
		# (without using ultraSSD)
		az vm create --resource-group "${RG}" \
						     --name "${vmName}" \
								 --availability-set "${AV_SET_NAME}" \
								 --vnet-name "${VNET_NAME}" \
								 --subnet "${SUBNET_NAME}" \
								 --public-ip-address "" \
								 --nsg "" \
								 --os-disk-size-gb 4095 \
								 --os-disk-caching none \
								 --storage-sku Premium_LRS \
								 --admin-username "${ADMIN_USER}" \
								 --ssh-key-value "${SSH_KEY_VALUE}" \
								 --image ubuntults
		let current=current+1 
	done
}

generate_certs(){
	inf "Generating certs"
}

main(){
	inf "Working dir is [$(pwd)]"
	inf "Script dir is [${CURRENT_DIR}]"

	validate_and_default
	create_vms
}


# Start here
main
