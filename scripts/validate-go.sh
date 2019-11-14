#!/usr/bin/env bash

# Copyright 2016 The Kubernetes Authors All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
set -euo pipefail

echo
echo "==> Running Go linter <=="
golangci-lint --version
if [ -f /.dockerenv ]; then
    echo "Running inside container";
fi

# exit 1 if golangci-lint output contains error text, to work
# around https://github.com/golangci/golangci-lint/issues/276
exec 5>&1
if ! OUTPUT=$(golangci-lint run --modules-download-mode=vendor 2>&1 | tee >(cat - >&5)) || grep 'skipped due to error' <<<$OUTPUT; then
  exit 1
fi
