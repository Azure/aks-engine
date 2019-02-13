#!/usr/bin/env bash

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
	"pkg/engine/templates_generated.go"
	"pkg/i18n/translations_generated.go"
)

for file in $GENERATED_FILES; do
	if ! diff  -r "${DIR}/../${file}" "${T}/${file}" 2>&1 ; then
		echo "go generate produced changes that were not already present"
		exit 1
	fi
done

echo "Generated assets have no material difference than what is committed."
