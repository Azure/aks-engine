# Build Docker image

**Bash**
```bash
$ export VERSION=0.32.0
$ docker build --no-cache --build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` --build-arg AKSENGINE_VERSION="$VERSION" -t microsoft/aks-engine:$VERSION --file ./Dockerfile.linux .
```
**PowerShell**
```powershell
PS> $VERSION="0.32.0"
PS> docker build --no-cache --build-arg BUILD_DATE=$(Get-Date((Get-Date).ToUniversalTime()) -UFormat "%Y-%m-%dT%H:%M:%SZ") --build-arg AKSENGINE_VERSION="$VERSION" -t microsoft/aks-engine:$VERSION --file .\Dockerfile.linux .
```

# Inspect Docker image metadata

**Bash**
```bash
$ docker image inspect microsoft/aks-engine:0.32.0 --format "{{json .Config.Labels}}" | jq
{
  "maintainer": "Microsoft",
  "org.label-schema.build-date": "2017-10-25T04:35:06Z",
  "org.label-schema.description": "The Azure Kubernetes Engine (aks-engine) generates ARM (Azure Resource Manager) templates for Kubernetes clusters on Microsoft Azure with your choice of DCOS, Kubernetes, or Swarm orchestrators.",
  "org.label-schema.docker.cmd": "docker run --rm microsoft/aks-engine:0.32.0",
  "org.label-schema.license": "MIT",
  "org.label-schema.name": "Azure Kubernetes Engine (aks-engine)",
  "org.label-schema.schema-version": "1.0",
  "org.label-schema.url": "https://github.com/Azure/aks-engine",
  "org.label-schema.usage": "https://github.com/Azure/aks-engine/blob/master/docs/aksengine.md",
  "org.label-schema.vcs-url": "https://github.com/Azure/aks-engine.git",
  "org.label-schema.vendor": "Microsoft",
  "org.label-schema.version": "0.32.0"
}
```

**PowerShell**
```powershell
PS> docker image inspect microsoft/aks-engine:0.32.0 --format "{{json .Config.Labels}}" | ConvertFrom-Json | ConvertTo-Json
{
    "maintainer":  "Microsoft",
    "org.label-schema.build-date":  "2017-10-25T04:35:06Z",
    "org.label-schema.description":  "The Azure Kubernetes Engine (aks-engine) generates ARM (Azure Resource Manager) templates for Kubernetes clusters on Microsoft Azure.",
    "org.label-schema.docker.cmd":  "docker run --rm microsoft/aks-engine:0.32.0",
    "org.label-schema.license":  "MIT",
    "org.label-schema.name":  "Azure Kubernetes Engine (aks-engine)",
    "org.label-schema.schema-version":  "1.0",
    "org.label-schema.url":  "https://github.com/Azure/aks-engine",
    "org.label-schema.usage":  "https://github.com/Azure/aks-engine/blob/master/docs/aksengine.md",
    "org.label-schema.vcs-url":  "https://github.com/Azure/aks-engine.git",
    "org.label-schema.vendor":  "Microsoft",
    "org.label-schema.version":  "0.32.0"
}
```

# Run Docker image

```
$ docker run --rm microsoft/aks-engine:0.32.0
```
