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
files=$(find . -type f -name "*.sh" -not -path './vendor/*' -not -path './hack/tools/*' -not -path './test/e2e/vendor/*' -not -path "*dcos*" -not -path "*swarm*")

IGNORED="
SC1127
SC1009
SC1054
SC1056
SC1072
SC1073
SC1083
SC1090
SC1091
SC2004
SC2015
SC2034
SC2046
SC2053
SC2068
SC2086
SC2128
SC2145
SC2154
SC2206
"

shellcheck $(printf -- "-e %s " $IGNORED) $files
