#!/usr/bin/env bash

set -x

# shellcheck disable=SC2034
if ! output=$(git remote show "$UPGRADE_FORK") ; then
  git remote add $UPGRADE_FORK https://github.com/$UPGRADE_FORK/aks-engine.git
fi

git fetch $UPGRADE_FORK
git branch -a --list "${UPGRADE_FORK}/release-*" | sort -r | head -$BACK_COMPAT_VERSIONS | sed "s/remotes\/${UPGRADE_FORK}\///" | sed -e 's/^[[:space:]]*//g'
