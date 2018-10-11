#!/usr/bin/env bash

# TODO: does this script need refactoring post-OpenShift, or can it be deleted?
echo "No generated files to check!"
exit 0

####################################################
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
  DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
####################################################

set -x

T="$(mktemp -d)"
trap "rm -rf ${T}" EXIT

cp -a "${DIR}/.." "${T}/"

(cd "${T}/" && go generate ./...)

GENERATED_FILES=(
	"pkg/openshift/certgen/unstable/templates/bindata.go"
	"pkg/openshift/certgen/release39/templates/bindata.go"
)

for file in $GENERATED_FILES; do
	if ! diff  -r "${DIR}/../${file}" "${T}/${file}" 2>&1 ; then
		echo "go generate produced changes that were not already present"
		exit 1
	fi
done

echo "Generated assets have no material difference than what is committed."
