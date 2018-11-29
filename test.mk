LOCATION ?= westus2
CLUSTER_DEFINITION ?= examples/kubernetes.json
SSH_KEY_NAME ?= id_rsa

TEST_CMD = docker run --rm \
						-v ${CURDIR}:${DEV_ENV_WORK_DIR} \
						-w ${DEV_ENV_WORK_DIR} \
						-e LOCATION=${LOCATION} \
						-e CLIENT_ID=${CLIENT_ID} \
						-e CLIENT_SECRET=${CLIENT_SECRET} \
						-e TENANT_ID=${TENANT_ID} \
						-e SUBSCRIPTION_ID=${SUBSCRIPTION_ID} \
						-e CLUSTER_DEFINITION=${CLUSTER_DEFINITION} \
						-e DNS_PREFIX=${DNS_PREFIX} \
						-e SSH_KEY_NAME=${SSH_KEY_NAME}

test-interactive:
	${TEST_CMD} -it -e TEST=kubernetes ${DEV_ENV_IMAGE} bash

test-functional: test-kubernetes

test-kubernetes-with-container:
	${TEST_CMD} -e ORCHESTRATOR=kubernetes ${DEV_ENV_IMAGE} test/e2e/runner

test-kubernetes:
	@ORCHESTRATOR=kubernetes go run ./test/e2e/runner.go

test-dcos:
	@ORCHESTRATOR=dcos go run ./test/e2e/runner.go

test-azure-constants:
	./scripts/azure-const.sh