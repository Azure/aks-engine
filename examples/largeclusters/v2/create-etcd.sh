#!/bin/bash

set -eo pipefail

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Expected input
#SUB_ID							subscription-id									optional(default what az is configured to use)
#RG									resource_group 									mandatory
#VNET_NAME					vnet name												mandatory
#SUBNET_NAME				subnet name											mandatory
#BASE_CLUSTER_NAME	basename for cluster resources	optional(default random)
#CLUSTER_SIZE				etcd cluster size								optional(default 3)
#VM_SIZE						etcd vm size										optional(default: Standard_DS2_v2)
#OUTPUT_DIR					location of generated certs			optional(default /tmp/${BASE_CLUSTER_NAME}
#ETCD_VERSION       etcd version										optional, defaults to 3.3.10

# Global Vars
AV_SET_NAME=""
ADMIN_USER=${ADMIN_USER:-azureuser}
SSH_KEY_VALUE=${SSH_KEY_VALUE:-$(cat ~/.ssh/id_rsa.pub)}

CLOUD_INIT_FILE="cloudinit.yaml"

ETCD_DISCOVERY_FILE="etcd.discovery"
ETCD_SERVER_CERT_DIR="etcd.server"
ETCD_PEER_CERT_DIR="etcd.peer"
ETCD_CLIENT_CERT_DIR="etcd.client"

ETCD_ENV="DAEMON_ARGS=--name <<SERVER_NAME>> --peer-client-cert-auth --peer-trusted-ca-file=/etc/kubernetes/certs/peer.ca.crt --peer-cert-file=/etc/kubernetes/certs/peer.crt --peer-key-file=/etc/kubernetes/certs/peer.key --initial-advertise-peer-urls https://0.0.0.0:2380 --client-cert-auth --trusted-ca-file=/etc/kubernetes/certs/client.ca.crt --cert-file=/etc/kubernetes/certs/server.crt --key-file=/etc/kubernetes/certs/server.key --advertise-client-urls https://0.0.0.0:2379 --listen-client-urls https://0.0.0.0:2379 --discovery <<DISCOVERY>> --data-dir /var/lib/etcddisk --initial-cluster-state new"



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
	if [[ -z "${length}" ]] || [[ "0" -eq "${length}" ]]; then
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

	if [[ -z "${ETCD_VERSION}" ]]; then
		ETCD_VERSION="3.3.10"
		wrn "etcd version [ETCD_VERSION] was not set, defaulting to ${ETCD_VERSION}"
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
								 --image ubuntults \
                 --custom-data "${OUTPUT_DIR}/${CLOUD_INIT_FILE}"
		let current=current+1 
	done
}


generate_certs(){
	local server_cert_dir="${OUTPUT_DIR}/${ETCD_SERVER_CERT_DIR}"
	local peer_cert_dir="${OUTPUT_DIR}/${ETCD_PEER_CERT_DIR}"
	local client_cert_dir="${OUTPUT_DIR}/${ETCD_CLIENT_CERT_DIR}"

	mkdir -p "${server_cert_dir}"
	mkdir -p "${peer_cert_dir}"
	mkdir -p "${client_cert_dir}"

	inf "Generating etcd certs for server"
	#CA
	openssl genrsa -out "${server_cert_dir}/server.ca.key"  2048 2>/dev/null
  #CA crt
	openssl req -x509 \
							-new \
							-nodes \
							-key "${server_cert_dir}/server.ca.key" \
							-days $((5 * 365)) \
							-out "${server_cert_dir}/server.ca.crt" \
							-subj "/C=US" 2>/dev/null
	# key
	openssl genrsa \
					-out "${server_cert_dir}/server.key" 4096 2>/dev/null

  # csr
	openssl req -new \
							-key "${server_cert_dir}/server.key" \
							-out "${server_cert_dir}/server.csr" \
							-subj "/C=US" 2>/dev/null

	# sign
	openssl x509 -req \
							-in "${server_cert_dir}/server.csr" \
							-CA "${server_cert_dir}/server.ca.crt" \
							-CAkey "${server_cert_dir}/server.ca.key" \
							-CAcreateserial -out "${server_cert_dir}/server.crt" \
							-days $((5 * 365))  2>/dev/null

	inf "etcd server certs are in ${server_cert_dir}"
	
	inf "Generating etcd certs for peer"
	#CA
	openssl genrsa -out "${peer_cert_dir}/peer.ca.key"  2048 2>/dev/null
  #CA crt
	openssl req -x509 \
							-new \
							-nodes \
							-key "${peer_cert_dir}/peer.ca.key" \
							-days $((5 * 365)) \
							-out "${peer_cert_dir}/peer.ca.crt" \
							-subj "/C=US" 2>/dev/null
	# key
	openssl genrsa \
					-out "${peer_cert_dir}/peer.key" 4096 2>/dev/null

  # csr
	openssl req -new \
							-key "${peer_cert_dir}/peer.key" \
							-out "${peer_cert_dir}/peer.csr" \
							-subj "/C=US" 2>/dev/null

	# sign
	openssl x509 -req \
							-in "${peer_cert_dir}/peer.csr" \
							-CA "${peer_cert_dir}/peer.ca.crt" \
							-CAkey "${peer_cert_dir}/peer.ca.key" \
							-CAcreateserial -out "${peer_cert_dir}/peer.crt" \
							-days $((5 * 365))  2>/dev/null

 	inf "etcd peer certs are in ${peer_cert_dir}"	
 	inf "Generating etcd certs for client"
	#CA
	openssl genrsa -out "${client_cert_dir}/client.ca.key"  2048 2>/dev/null
  #CA crt
	openssl req -x509 \
							-new \
							-nodes \
							-key "${client_cert_dir}/client.ca.key" \
							-days $((5 * 365)) \
							-out "${client_cert_dir}/client.ca.crt" \
							-subj "/C=US" 2>/dev/null
	# key
	openssl genrsa \
					-out "${client_cert_dir}/client.key" 4096 2>/dev/null

  # csr
	openssl req -new \
							-key "${client_cert_dir}/client.key" \
							-out "${client_cert_dir}/client.csr" \
							-subj "/C=US" 2>/dev/null

	# sign
	openssl x509 -req \
							-in "${client_cert_dir}/client.csr" \
							-CA "${client_cert_dir}/client.ca.crt" \
							-CAkey "${client_cert_dir}/client.ca.key" \
							-CAcreateserial -out "${client_cert_dir}/client.crt" \
							-days $((5 * 365))  2>/dev/null

 inf "etcd client certs are in ${client_cert_dir}"	

}

# all offline work
# 1- Get cluster discovery token
# 2- Generates the certs
# 3- Put vars in cloud init 
pre_work(){
	inf "Generating etcd discovery token"
	curl -s "https://discovery.etcd.io/new?size=${CLUSTER_SIZE}" --output "${OUTPUT_DIR}/${ETCD_DISCOVERY_FILE}"
	inf "etcd token generated in ${OUTPUT_DIR}/${ETCD_DISCOVERY_FILE}"

	generate_certs

	# Modify cloud init
	inf "Creating cloudinit file"
	cloudinit_src="${CURRENT_DIR}/${CLOUD_INIT_FILE}"
	cloudinit_dst="${OUTPUT_DIR}/${CLOUD_INIT_FILE}"
	cp "${cloudinit_src}" "${cloudinit_dst}"

	local server_cert_dir="${OUTPUT_DIR}/${ETCD_SERVER_CERT_DIR}"
	local peer_cert_dir="${OUTPUT_DIR}/${ETCD_PEER_CERT_DIR}"
	local client_cert_dir="${OUTPUT_DIR}/${ETCD_CLIENT_CERT_DIR}"
	
	# certs, keys et al
	sed -i "s/<PEER_CA>/$(cat ${peer_cert_dir}/peer.ca.crt | base64 -w 0)/g" ${cloudinit_dst}
	sed -i "s/<PEER_CRT>/$(cat ${peer_cert_dir}/peer.crt | base64 -w 0)/g" ${cloudinit_dst}
	sed -i "s/<PEER_KEY>/$(cat ${peer_cert_dir}/peer.key | base64 -w 0)/g" ${cloudinit_dst}

	sed -i "s/<SERVER_CRT>/$(cat ${server_cert_dir}/server.crt | base64 -w 0)/g" ${cloudinit_dst}
	sed -i "s/<SERVER_KEY>/$(cat ${server_cert_dir}/server.key | base64 -w 0)/g" ${cloudinit_dst}

	sed -i "s/<CLIENT_CA>/$(cat ${client_cert_dir}/client.ca.crt | base64 -w 0)/g" ${cloudinit_dst}

	# etcd env file 
	sed -i "s/<ETCD_ENV>/$(echo -n ${ENV} | base64 -w 0)/g" ${cloudinit_dst}

	# etcd version
	sed -i "s/<ETCD_VERSION>/${ETCD_VERSION}/g" ${cloudinit_dst}

	inf "cloudinit file is in ${cloudinit_dst}"
}

main(){
	inf "Working dir is [$(pwd)]"
	inf "Script dir is [${CURRENT_DIR}]"

	validate_and_default
	pre_work
	create_vms
}


# Start here
main
