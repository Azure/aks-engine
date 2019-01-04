#!/usr/bin/env bash

set -euo pipefail

echo "==> Checking commit message <=="

git log -1 --pretty=%B | grep -E '^(build|ci|chore|docs|feat|fix|perf|refactor|revert|style|test)(\(.+\))?: .+' \
  || (echo "Commits must follow https://www.conventionalcommits.org/en/v1.0.0-beta.2/#specification with one of these types:
  build, ci, chore, docs, feat, fix, perf, refactor, revert, style, test" && exit 1)
