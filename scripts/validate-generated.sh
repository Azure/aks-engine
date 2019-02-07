#!/usr/bin/env bash

set -euo pipefail

####################################################
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
  DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
####################################################

export GOFILES="$(go list ./... | grep -v "github.com/Azure/aks-engine/vendor" | sed 's|github.com/Azure/aks-engine|.|g' | grep -v -w '^.$')"

T="$(mktemp -d)"
trap "rm -rf ${T}" EXIT

cp -a "${DIR}/.." "${T}/"

(cd "${T}/" && go generate ${GOFILES})

if ! diff -I '.*bindataFileInfo.*' --exclude=.git --exclude=vendor -r "${DIR}/.." "${T}" 2>&1 ; then
	echo "go generate produced changes that were not already present"
	exit 1
fi

echo "Generated assets have no material differences from what is committed."
