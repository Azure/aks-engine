$REPO_PATH = "github.com/Azure/aks-engine"
$DEV_ENV_IMAGE = "quay.io/deis/go-dev:v1.17.2"
$DEV_ENV_WORK_DIR = "/go/src/$REPO_PATH"

docker.exe run -it --rm -w $DEV_ENV_WORK_DIR -v `"$($PWD)`":$DEV_ENV_WORK_DIR $DEV_ENV_IMAGE bash
