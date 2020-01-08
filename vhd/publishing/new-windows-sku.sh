#!/bin/bash -e

required_env_vars=(
    "SKU_PREFIX"
    "SKU_TEMPLATE_FILE"
    "AZURE_TENANT_ID"
    "AZURE_CLIENT_ID"
    "AZURE_CLIENT_SECRET"
    "PUBLISHER"
    "OFFER"
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

short_date=$(date +"%y%m")
pretty_date=$(date +"%b %Y")

sku_id="${SKU_PREFIX}-${short_date}"

cat $SKU_TEMPLATE_FILE | sed s/{{ID}}/"$sku_id"/ | sed s/{{MONTH-YEAR}}/"$pretty_date/" > sku.json
cat sku.json

echo "Creating new SKU"
(set -x ; hack/tools/bin/pub skus put -p $PUBLISHER -o "$OFFER" -f sku.json ; echo "")

echo "Wrting publishing info"
cat <<EOF > sku-publishing-info.json
{
    "publisher" : "$PUBLISHER",
    "offer" : "$OFFER",
    "sku_id" : "$sku_id",
    "sku_prefix" : "$SKU_PREFIX"
}
EOF

cat sku-publishing-info.json
