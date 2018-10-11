$pwd = (Get-Location).Path

docker build --pull -t aks-engine .
docker run --security-opt seccomp:unconfined -it `
	-v ${pwd}:/gopath/src/github.com/Azure/aks-engine `
	-w /gopath/src/github.com/Azure/aks-engine `
		aks-engine /bin/bash

