#!/usr/bin/env bash

set -eu -o pipefail
set -x

docker build --pull -t aks-engine .

docker run -it \
	--privileged \
	--security-opt seccomp:unconfined \
	-v /var/run/docker.sock:/var/run/docker.sock \
	-v `pwd`:/gopath/src/github.com/Azure/aks-engine \
	-v ~/.azure:/root/.azure \
	-w /gopath/src/github.com/Azure/aks-engine \
		aks-engine /bin/bash

chown -R "$(logname):$(id -gn $(logname))" . ~/.azure
