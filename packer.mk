build-packer:
	@packer build -var-file=packer/settings.json packer/vhd-image-builder.json

build-packer-windows:
	@packer build -var-file=packer/settings.json packer/windows-vhd-builder.json

init-packer:
	@./packer/init-variables.sh

az-login:
	az login --service-principal -u ${CLIENT_ID} -p ${CLIENT_SECRET} --tenant ${TENANT_ID}

run-packer: az-login
	@packer version && set -o pipefail && ($(MAKE) init-packer | tee packer-output) && ($(MAKE) build-packer | tee -a packer-output)

run-packer-windows: az-login
	@packer version && set -o pipefail && ($(MAKE) init-packer | tee packer-output) && ($(MAKE) build-packer-windows | tee -a packer-output)

az-copy: az-login
	azcopy-preview copy "${OS_DISK_SAS}" "${CLASSIC_BLOB}${CLASSIC_SAS_TOKEN}"

delete-sa: az-login
	az storage account delete -n ${SA_NAME} -g ${AZURE_RESOURCE_GROUP_NAME} --yes

generate-sas: az-login
	az storage container generate-sas --name vhds --permissions lr --connection-string "${CLASSIC_SA_CONNECTION_STRING}" --start ${START_DATE} --expiry ${EXPIRY_DATE} | tr -d '"' | tee -a vhd-sas && cat vhd-sas

windows-vhd-publishing-info: az-login
	@./packer/generate-windows-vhd-publishing-info.sh
