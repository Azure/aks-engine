#!/usr/bin/env bash
#
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT license.

####################################################
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
  DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
ROOT="${DIR}/.."
####################################################

set -x

GENERATED_FILES=(
	"pkg/i18n/translations_generated.go"
	"pkg/engine/templates_generated.go"
)

T="$(mktemp -d)"
trap 'rm -rf "${T}"' EXIT

for file in ${GENERATED_FILES[*]}; do
	cp -a "${file}" "${T}/"
done

make generate

for file in ${GENERATED_FILES[*]}; do
	basefile=$(basename "${file}")
	if ! diff -r "${T}/${basefile}" "${ROOT}/${file}" 2>&1 ; then
		echo "go generate produced changes that were not already present."
		echo "Make sure you include generated assets in your commit."
		exit 1
	fi
done

echo "Generated assets have no material difference than what is committed."
