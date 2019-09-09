#!/usr/bin/env bash
#
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT license.

exit_code=0

echo "==> Running go mod check <=="

go mod tidy

git diff --exit-code --quiet go.mod go.sum
error_code=$?

if [ $error_code -ne 0 ]; then
  echo "The dependency state is out of sync. Please commit changes to go.mod, go.sum."
else
  echo "go mod ok."
fi

exit $error_code
