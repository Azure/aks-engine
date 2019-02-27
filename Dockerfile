FROM golang:1.11.5-alpine3.9

LABEL name="aks-toolbox"

ENV HELM_VERSION=2.12.3 \
    KUBECTL_VERSION=1.10.12 \
    GOLANGCI_LINT_VERSION=1.12.5 \
    GOPATH=/go \
    PATH=$PATH:/go/bin

RUN apk add --no-cache -u bash curl dep gcc git jq make musl-dev \
    && go get -u -v \
        github.com/hashicorp/packer \
        github.com/jteeuwen/go-bindata/... \
        github.com/onsi/ginkgo/ginkgo \
    && curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $GOPATH/bin v$GOLANGCI_LINT_VERSION \
    && mkdir -p $GOPATH/src/k8s.io/helm \
    && curl -sSL https://github.com/helm/helm/archive/v$HELM_VERSION.tar.gz | tar -vxz -C $GOPATH/src/k8s.io/helm --strip=1 \
    && cd $GOPATH/src/k8s.io/helm && make bootstrap build \
    && curl -sSL https://storage.googleapis.com/kubernetes-release/release/$KUBECTL_VERSION/bin/linux/amd64/kubectl -o /usr/local/bin/kubectl \
    && chmod +x /usr/local/bin/kubectl \
    && cp ./bin/* /usr/local/bin \
    && rm -rf $GOPATH/pkg/* $GOPATH/src/* /root/.cache /root/.glide /tmp/*

COPY scripts/k /usr/local/bin
