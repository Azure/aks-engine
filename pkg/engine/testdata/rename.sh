#!/bin/bash

while IFS= read -r -d '' f; do
  len=${#f}
  mv ${f} ${f::len-4}
done < <(find . -name "*.err" -print0)
