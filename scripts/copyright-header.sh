#!/bin/bash
#
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT license.

set -euo pipefail

files=$(find .. -type f -iname '*.go' ! -path '../vendor/*')  # do not run on the vendor dir

for i in $files
do
  if ! grep -q Copyright "$i"
  then
    cat copyright.txt "$i" > "$i".new && mv "$i".new "$i"
  fi
done
