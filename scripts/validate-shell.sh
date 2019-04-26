#!/usr/bin/env bash
#
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT license.

set -euo pipefail

echo "==> Running shell script validator <=="

# All shell scripts, except those that support deprecated orchestrators or are in vendored code.
files=$(find . -type f -name "*.sh" -not -path './vendor/*' -not -path "*dcos*" -not -path "*swarm*")

IGNORED="
SC1041
SC1042
SC1072
SC1073
SC1090
SC1091
SC1113
SC1128
SC2001
SC2002
SC2004
SC2005
SC2006
SC2015
SC2016
SC2020
SC2024
SC2027
SC2034
SC2044
SC2046
SC2053
SC2059
SC2064
SC2068
SC2086
SC2116
SC2126
SC2128
SC2129
SC2140
SC2145
SC2154
SC2155
SC2162
SC2181
SC2196
SC2206
SC2230
SC2236
SC2242
"

shellcheck $(printf -- "-e %s " $IGNORED) $files
