#!/usr/bin/env bash

set -euo pipefail

echo "==> Running shell script validator <=="

files=$(find . -type f -iname "*.sh" -not -path './vendor/*' -not -path "*swarm*" -not -path "*dcos*"  -not -iname "*swarm*" -not -iname "*dcos*")

shellcheck $files
