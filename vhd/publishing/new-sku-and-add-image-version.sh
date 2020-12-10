#!/bin/bash -e

required_env_vars=(
    "SKU_PREFIX"
    "SKU_TEMPLATE_FILE"
    "AZURE_TENANT_ID"
    "AZURE_CLIENT_ID"
    "AZURE_CLIENT_SECRET"
    "PUBLISHER"
    "OFFER"
    "CONTAINER_RUNTIME"
)

for v in "${required_env_vars[@]}"
do
    if [ -z "${!v}" ]; then
        echo "$v was not set!"
        exit 1
    fi
done

if [ ! -f "$SKU_TEMPLATE_FILE" ]; then
    echo "Could not find sku template file: ${SKU_TEMPLATE_FILE}!"
    exit 1
fi

(set -x; ls -lhR artifacts )

VHD_INFO="artifacts/vhd/windows-vhd-publishing-info.json"

if [ ! -f "$VHD_INFO" ]; then
    echo "Could not find sku template file: ${VHD_INFO}!"
    exit 1
fi

short_date=$(date +"%y%m")
pretty_date=$(date +"%b %Y")

sku_id="${SKU_PREFIX}-${short_date}"

echo "Checking if offer contains SKU: $sku_id"
# Check if SKU already exists in offer
(set -x; hack/tools/bin/pub skus list -p microsoft-aks -o aks-windows | jq ".[] | .planId" | tr -d '"' | tee skus.txt)
echo ""

if grep -q $sku_id skus.txt; then
    echo "Offer already has SKU"
else
    echo "Creating new SKU"

    < $SKU_TEMPLATE_FILE sed s/{{ID}}/"$sku_id"/ | sed s/{{MONTH-YEAR}}/"$pretty_date/" | sed s/{{CONTAINER_RUNTIME}}/"$CONTAINER_RUNTIME/" > sku.json

    echo "" ; cat sku.json ; echo ""

    (set -x ; hack/tools/bin/pub skus put -p $PUBLISHER -o "$OFFER" -f sku.json ; echo "")
fi

# Get VHD version info for windows-vhd-publishing-info.json produced by previous pipeline stage
vhd_url=$(< $VHD_INFO jq -r ".vhd_url")
vhd_version=$(< $VHD_INFO jq -r ".windows_version")
version_date=$(date +"%y%m%d")
image_version="${vhd_version}.${version_date}"

# media name must be under 63 characters
media_name="aks-windows-${SKU_PREFIX}-${image_version}"
if [ "${#media_name}" -ge 63 ]; then
	echo "$media_name should be under 63 characters"
	exit 1
fi

published_date=$(date +"%m/%d/%Y")

(set -x ; hack/tools/bin/pub versions put corevm -p $PUBLISHER -o $OFFER -s $sku_id --version $image_version --vhd-url $vhd_uri --media-name $media_name --label "AKS Base Image for Windows" --desc "AKS Base Image for Windows" --published-date "$published_date")
