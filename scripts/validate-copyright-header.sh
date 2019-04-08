#!/usr/bin/env bash
#
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT license.

set -euo pipefail

echo "==> Checking copyright headers <=="

files=$(find . -type f -iname '*.go' ! -path './vendor/*')
licRes=$(for file in $files; do
           awk 'NR<=3' "$file" | grep -Eq "(Copyright|generated|GENERATED)" || echo "$file";
         done)

if [ -n "$licRes" ]; then
        echo "Copyright header check failed:";
        echo "${licRes}";
        exit 1;
fi
