#!/usr/bin/env bash
#
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT license.

set -euo pipefail

echo "==> Checking commit message <=="

msg=$(git log --no-merges -1 --pretty=%B)
echo "$msg" | grep -E '^(build|chore|ci|docs|feat|fix|perf|refactor|revert|style|test)(\(.+\))?: .+' \
  || (echo -e "Commit message \"$msg\" failed validation.\n" \
      "Commits must follow https://www.conventionalcommits.org/en/v1.0.0-beta.2/#specification with one of these types:\n" \
      "  build, ci, chore, docs, feat, fix, perf, refactor, revert, style, test" \
      && exit 1)
