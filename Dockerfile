FROM buildpack-deps:xenial

RUN apt-get update \
    && apt-get -y upgrade \
    && apt-get -y install python-pip make build-essential curl openssl vim jq gettext \
    && rm -rf /var/lib/apt/lists/*

ENV GO_VERSION 1.8.3

RUN wget -q https://storage.googleapis.com/golang/go${GO_VERSION}.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && rm go${GO_VERSION}.linux-amd64.tar.gz

RUN curl -fsSL https://get.docker.com/ | sh

ENV KUBECTL_VERSION 1.7.5
RUN curl "https://storage.googleapis.com/kubernetes-release/release/v${KUBECTL_VERSION}/bin/linux/amd64/kubectl" > /usr/local/bin/kubectl \
    && chmod +x /usr/local/bin/kubectl

ENV GOPATH /gopath
ENV PATH "${PATH}:${GOPATH}/bin:/usr/local/go/bin"

RUN git clone https://github.com/akesterson/cmdarg.git /tmp/cmdarg \
    && cd /tmp/cmdarg && make install && rm -rf /tmp/cmdarg
RUN git clone https://github.com/akesterson/shunit.git /tmp/shunit \
    && cd /tmp/shunit && make install && rm -rf /tmp/shunit

WORKDIR /gopath/src/github.com/Azure/aks-engine

# Cache vendor layer
ADD . /gopath/src/github.com/Azure/aks-engine/
RUN make bootstrap

# https://github.com/dotnet/core/blob/master/release-notes/download-archives/2.1.2-sdk-download.md
RUN curl https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > microsoft.gpg \
    && mv microsoft.gpg /etc/apt/trusted.gpg.d/microsoft.gpg \
    && sh -c 'echo "deb [arch=amd64] https://packages.microsoft.com/repos/microsoft-ubuntu-xenial-prod xenial main" > /etc/apt/sources.list.d/dotnetdev.list' \
    && apt-get update \
    && apt-get -y install dotnet-sdk-2.1.2

# See: https://docs.microsoft.com/en-us/cli/azure/install-azure-cli-apt
RUN apt-get update \
    && apt-get install apt-transport-https \
    && echo "deb [arch=amd64] https://packages.microsoft.com/repos/azure-cli/ wheezy main" > /etc/apt/sources.list.d/azure-cli.list \
    && apt-key adv --keyserver packages.microsoft.com --recv-keys 52E16F86FEE04B979B07E28DB02C46DF417A0893 \
    && curl -L https://packages.microsoft.com/keys/microsoft.asc | apt-key add - \
    && apt-get update \
    && apt-get install azure-cli

