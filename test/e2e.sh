#!/bin/bash

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
  DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"

###############################################################################

set -e
set -o pipefail

ROOT="${DIR}/.."

# Check pre-requisites
[[ ! -z "${SERVICE_PRINCIPAL_CLIENT_ID:-}" ]]             || (echo "Must specify SERVICE_PRINCIPAL_CLIENT_ID" && exit -1)
[[ ! -z "${SERVICE_PRINCIPAL_CLIENT_SECRET:-}" ]]         || (echo "Must specify SERVICE_PRINCIPAL_CLIENT_SECRET" && exit -1)
[[ ! -z "${TENANT_ID:-}" ]]                               || (echo "Must specify TENANT_ID" && exit -1)
[[ ! -z "${SUBSCRIPTION_ID:-}" ]]                         || (echo "Must specify SUBSCRIPTION_ID" && exit -1)
[[ ! -z "${CLUSTER_SERVICE_PRINCIPAL_CLIENT_ID:-}" ]]     || (echo "Must specify CLUSTER_SERVICE_PRINCIPAL_CLIENT_ID" && exit -1)
[[ ! -z "${CLUSTER_SERVICE_PRINCIPAL_CLIENT_SECRET:-}" ]] || (echo "Must specify CLUSTER_SERVICE_PRINCIPAL_CLIENT_SECRET" && exit -1)
[[ ! -z "${STAGE_TIMEOUT_MIN:-}" ]]                       || (echo "Must specify STAGE_TIMEOUT_MIN" && exit -1)
[[ ! -z "${TEST_CONFIG:-}" ]]                             || (echo "Must specify TEST_CONFIG" && exit -1)

make bootstrap build

${ROOT}/test/aks-engine-test/aks-engine-test -c ${TEST_CONFIG} -d ${ROOT} -e ${LOGERROR_CONFIG:-${ROOT}/test/aks-engine-test/aks-engine-errors.json} -j ${SA_NAME} -k ${SA_KEY}
