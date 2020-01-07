#!/bin/bash -e

required_env_vars=(
    "CLASSIC_SA_CONNECTION_STRING"
    "STORAGE_ACCT_BLOB_URL"
    "VHD_NAME"
)

for v in "${required_env_vars[@]}"
do
    if [ -z "${!v}" ]; then
        echo "$v was not set!"
        exit 1
    fi
done

start_date=$(date +"%Y-%m-%dT00:00Z" -d "-1 day")
expiry_date=$(date +"%Y-%m-%dT00:00Z" -d "+1 year")
sas_token=$(az storage container generate-sas --name vhds --permissions lr --connection-string ${CLASSIC_SA_CONNECTION_STRING} --start ${start_date} --expiry ${expiry_date} | tr -d '"')
vhd_url="${STORAGE_ACCT_BLOB_URL}/${VHD_NAME}?$sas_token"

# Get verion info from release notes
RELEASE_NOTES_FILE=release-notes.txt

if [ ! -f "$RELEASE_NOTES_FILE" ]; then
    echo "Could not fine release notes file!"
    exit 1
fi

windows_version=$(cat release-notes.txt | grep "OS Version" | cut -d ":" -f 2 | tr -d [:space:])

cat <<EOF > windows-vhd-publishing-info.json
{
    "vhd_url" : "$vhd_url",
    "windows_version" : "$windows_version"
}
EOF

cat windows-vhd-publishing-info.json
