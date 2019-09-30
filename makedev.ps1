$REPO_PATH = "github.com/Azure/aks-engine"
$DEV_ENV_IMAGE = "quay.io/deis/go-dev:v1.23.3"
$DEV_ENV_WORK_DIR = "/go/src/$REPO_PATH"

# Ensure docker is configured for linux containers
$json = (docker version --format "{{json .}}" | ConvertFrom-Json)
if ($json.Server.Os -ne "linux")
{
    Write-Error "Please switch Docker use to Linux containers on Windows"
    exit 1
}

docker.exe run -it --rm -w $DEV_ENV_WORK_DIR -v `"$($PWD)`":$DEV_ENV_WORK_DIR $DEV_ENV_IMAGE bash
