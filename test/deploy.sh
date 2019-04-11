#!/usr/bin/env bash

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
  DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"

###############################################################################

set -e
set -u
set -o pipefail

ROOT="${DIR}/.."

# Usage:
#
# Manual user usage (Specific name):
#   export INSTANCE_NAME=official-jenkins-infra
#   ./scripts/deploy.sh ./examples/kubernetes.json
#
# Manual user usage (Lots of rapid fire):
# In this mode, the user can repeat the same deploy
# command blindly and get new clusters each time.
#   unset INSTANCE_NAME
#   vim ./test/user.env (add your stuff)
#   ./scripts/deploy.sh ./examples.kubernetes.json
#   sleep 1
#   ./scripts/deploy.sh ./examples.kubernetes.json
#
# Prow:
#   export PULL_NUMBER=...
#   export VALIDATE=<script path>
#   export CLUSTER_DEFIITION=examples/kubernetes.json
#   ./scripts/deploy.sh

# Load any user set environment
if [[ -f "${ROOT}/test/user.env" ]]; then
	source "${ROOT}/test/user.env"
fi


# Ensure Cluster Definition
if [[ -z "${CLUSTER_DEFINITION:-}" ]]; then
	if [[ -z "${1:-}" ]]; then echo "You must specify a parameterized apimodel.json clusterdefinition"; exit -1; fi
	CLUSTER_DEFINITION="${1}"
fi

# Set Instance Name for PR or random run
if [[ -n "${PULL_NUMBER:-}" ]]; then
	export INSTANCE_NAME="${JOB_NAME}-${PULL_NUMBER}-$(printf "%x" $(date '+%s'))"
	# if we're running a pull request, assume we want to cleanup unless the user specified otherwise
	if [[ -z "${CLEANUP:-}" ]]; then
		export CLEANUP="y"
	fi
else
	export INSTANCE_NAME_DEFAULT="${INSTANCE_NAME_PREFIX}-$(printf "%x" $(date '+%s'))"
	export INSTANCE_NAME="${INSTANCE_NAME:-${INSTANCE_NAME_DEFAULT}}"
fi

# Let the example json.env file set any env vars it may need ahead of time
# (For example, the `managed-identity/kubernetes.json.env` sets env vars for a
# custom MSI-compatible build of Kubernetes, as well as the SP cred values.)
ENV_FILE="${CLUSTER_DEFINITION}.env"
if [ -e "${ENV_FILE}" ]; then
  source "${ENV_FILE}"
fi


# Set extra parameters
export OUTPUT="${ROOT}/_output/${INSTANCE_NAME}"
export RESOURCE_GROUP="${INSTANCE_NAME}"
export DEPLOYMENT_NAME="${INSTANCE_NAME}"

source "${ROOT}/test/common.sh"

# Set custom dir so we don't clobber global 'az' config
export AZURE_CONFIG_DIR="$(mktemp -d)"
trap 'rm -rf ${AZURE_CONFIG_DIR}' EXIT

make -C "${ROOT}" ci
generate_template
set_azure_account
trap cleanup EXIT
deploy_template

if [[ -z "${VALIDATE:-}" ]]; then
	exit 0
fi

export SSH_KEY="${OUTPUT}/id_rsa"
export KUBECONFIG="${OUTPUT}/kubeconfig/kubeconfig.${LOCATION}.json"

"${ROOT}/${VALIDATE}"

echo "post-test..."
