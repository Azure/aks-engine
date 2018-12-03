#/bin/bash

for f in $(find . -name "*.err"); do
        len=${#f}
	mv ${f} ${f::len-4};
done
