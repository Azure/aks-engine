#!/usr/bin/env bash
#
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT license.

set -euo pipefail

echo "==> Running shell linter <=="
shellcheck --version
if [ -f /.dockerenv ]; then
    echo "Running inside container";
fi

# All shell scripts, except those that support deprecated orchestrators or are in vendored code.
files=$(find . -type f -name "*.sh" -not -path './vendor/*' -not -path "*dcos*" -not -path "*swarm*")

IGNORED="
SC1090
SC1091
SC2001
SC2004
SC2005
SC2015
SC2020
SC2034
SC2044
SC2046
SC2053
SC2064
SC2068
SC2086
SC2116
SC2128
SC2145
SC2154
SC2155
SC2162
SC2181
SC2206
SC2242
"

shellcheck $(printf -- "-e %s " $IGNORED) $files
