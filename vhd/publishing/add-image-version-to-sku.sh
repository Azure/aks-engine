#!/bin/bash

cd /build
echo "PWD: $PWD"

required_env_vars=(
    "AZURE_CLIENT_ID"
    "AZURE_CLIENT_SECRET"
    "AZURE_TENANT_ID"
)

for v in "${required_env_vars[@]}"
do
    if [ -z "${!v}" ]; then
        echo "$v was not set!"
        exit 1
    fi
done

SKU_INFO="sku/sku_info/sku-publishing-info.json"
VHD_INFO="vhd/publishing-info/windows-vhd-publishing-info.json"

required_files=(
    "SKU_INFO"
    "VHD_INFO"
)

for f in "${required_files[@]}"
do
    if [ ! -f "${!f}" ]; then
        echo "could not find file: ${!f}"
        exit 1
    fi
done

echo "Getting pub..."
(set -x ; go get -u github.com/devigned/pub@v0.2.0 > /dev/null 2>&1)

echo "Sku publishing info:"
cat $SKU_INFO
echo

echo "Vhd publishing info:"
cat $VHD_INFO
echo

# generate image version
vhd_version=$(cat $VHD_INFO | jq -r ".windows_version")
version_date=$(date +"%y%m%d")
image_version="${vhd_version}.${version_date}"

# generate media name
sku_prefix=$(cat $SKU_INFO | jq -r ".sku_prefix")
media_name="aks-windows-${sku_prefix}-${image_version}"

# generate published date
published_date=$(date +"%d/%m/%Y")

# get vhd url
vhd_url=$(cat $VHD_INFO | jq -r ".vhd_url")

# create version.json
cat <<EOF > version.json
{
    "$image_version" : {
        "mediaName": "$media_name",
        "showInGui": false,
        "publishedDate": "$published_date",
        "label": "AKS Base Image for Windows",
        "description": "AKS Base Image for Windows",
        "osVHdUrl": "$vhd_url"
    }
}
EOF

echo "Version info:"
cat version.json

publisher=$(cat $SKU_INFO | jq -r ".publisher")
offer=$(cat $SKU_INFO | jq -r ".offer")
sku=$(cat $SKU_INFO | jq -r ".sku_id")

(set -x ; pub versions put corevm -p $publisher -o aks-windows -s $sku --version $image_version --vhd-uri $vhd_url --media-name $media_name --label "AKS Base Image for Windows" --desc "AKS Base Image for Windows" --published-date "$published_date")