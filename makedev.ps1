$DEV_ENV_IMAGE = "quay.io/deis/go-dev:v1.27.0"
$DEV_ENV_WORK_DIR = "/aks-engine"

# Ensure docker is configured for linux containers
$json = (docker version --format "{{json .}}" | ConvertFrom-Json)
if ($json.Server.Os -ne "linux")
{
    Write-Error "Please switch Docker use to Linux containers on Windows"
    exit 1
}

docker.exe run -it --rm -w $DEV_ENV_WORK_DIR -v `"$($GOPATH)/pkg/mod`":/go/pkg/mod -v `"$($PWD)`":$DEV_ENV_WORK_DIR $DEV_ENV_IMAGE bash
