#!/usr/bin/env bash

set -x

GOPATH="/go"
WORK_DIR="${GOPATH}/src/github.com/Azure/aks-engine"

ORIGINAL_BRANCH=$(git branch | grep \* | cut -d ' ' -f2)

if [[ -n "${FORK}" ]]; then
  # shellcheck disable=SC2034
  if ! output=$(git remote show "$FORK") ; then
    git remote add $FORK https://github.com/$FORK/aks-engine.git
  fi

  git fetch $FORK
  git branch -D $FORK/$BRANCH
  git checkout -b $FORK/$BRANCH --track $FORK/$BRANCH
  git pull
fi

# Assumes we're running from the git root of aks-engine
docker run --rm \
-v $(pwd):${WORK_DIR} \
-w ${WORK_DIR} \
"${DEV_IMAGE}" make build-binary || exit 1

git checkout "${ORIGINAL_BRANCH}"
