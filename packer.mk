build-packer:
	@packer build -var-file=packer/settings.json packer/vhd-image-builder.json

init-packer:
	@./packer/init-variables.sh

az-login:
	az login --service-principal -u ${CLIENT_ID} -p ${CLIENT_SECRET} --tenant ${TENANT_ID}

run-packer: az-login
	@packer version && set -o pipefail && ($(MAKE) init-packer | tee packer-output) && ($(MAKE) build-packer | tee -a packer-output)

az-copy: az-login
	azcopy --source "${OS_DISK_SAS}" --destination "${CLASSIC_BLOB}/${VHD_NAME}" --dest-sas "${CLASSIC_SAS_TOKEN}"

delete-sa: az-login
	az storage account delete -n ${SA_NAME} -g ${AZURE_RESOURCE_GROUP_NAME} --yes

generate-sas: az-login
	az storage container generate-sas --name vhds --permissions lr --connection-string "${CLASSIC_SA_CONNECTION_STRING}" --start ${START_DATE} --expiry ${EXPIRY_DATE} | tr -d '"' | tee -a vhd-sas && cat vhd-sas

make vhd-notes:
	awk '/START_OF_NOTES/{y=1;next}y' packer-output > release-notes-raw.txt && awk '/END_OF_NOTES/ {exit} {print}' release-notes-raw.txt | sed -r "s/\x1B\[([0-9]{1,2}(;[0-9]{1,2})?)?[mGK]//g" > release-notes.txt && sed -e s/azure-arm://g -i release-notes.txt && cat release-notes.txt