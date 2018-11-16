#/bin/bash

# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT license.

for f in $(find . -name "*.err"); do
        len=${#f}
	mv ${f} ${f::len-4};
done
