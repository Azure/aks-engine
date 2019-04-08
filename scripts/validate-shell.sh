#!/usr/bin/env bash
#
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT license.

set -euo pipefail

echo "==> Running shell script validator <=="

# All shell scripts, except those that support deprecated orchestrators or are in vendored code.
#files=$(find . -type f -name "*.sh" -not -path './vendor/*' -not -path "*dcos*" -not -path "*swarm*")
# TODO: make the below list converge on the above one as we clean up scripts.
files=$(find scripts -type f -name "*.sh")

# shellcheck disable=SC2086
shellcheck $files
