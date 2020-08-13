 {{if IsKubernetes}}
    "kubeBinariesSASURL": {
      "metadata": {
        "description": "The download url for kubernetes windows binaries package"
      },
      "type": "string"
    },
    "windowsKubeBinariesURL": {
      "metadata": {
        "description": "The download url for kubernetes windows binaries produce by Kubernetes. This contains only the node binaries (example: https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG-1.11.md#node-binaries-1)"
      },
      "type": "string"
    },
    "kubeBinariesVersion": {
      "metadata": {
        "description": "Kubernetes windows binaries version"
      },
      "type": "string"
    },
    "windowsContainerdURL": {
      "metadata": {
        "description": "TODO: containerd - these binaries are not available yet"
      },
      "type": "string"
    },
    "windowsSdnPluginURL": {
      "metadata": {
        "description": "TODO: containerd - these binaries are not available yet"
      },
      "type": "string"
    },
    "kubeServiceCidr": {
      "metadata": {
        "description": "Kubernetes service address space"
      },
      "type": "string"
    },
    "windowsTelemetryGUID": {
      "metadata": {
        "description": "The GUID to set in windows agent to collect telemetry data."
      },
      "type": "string"
    },
 {{end}}
    "windowsAdminUsername": {
      "type": "string",
      "metadata": {
        "description": "User name for the Windows Swarm Agent Virtual Machines (Password Only Supported)."
      }
    },
    "windowsAdminPassword": {
      "type": "securestring",
      "metadata": {
        "description": "Password for the Windows Swarm Agent Virtual Machines."
      }
    },
    "agentWindowsImageName": {
      "defaultValue": "",
      "type": "string",
      "metadata": {
        "description": "Image name when specifying a Windows image reference."
      }
    },
    "agentWindowsImageResourceGroup": {
      "defaultValue": "",
      "type": "string",
      "metadata": {
        "description": "Resource group when specifying a Windows image reference."
      }
    },
    "agentWindowsVersion": {
      "defaultValue": "latest",
      "metadata": {
        "description": "Version of the Windows Server OS image to use for the agent virtual machines."
      },
      "type": "string"
    },
    "agentWindowsSourceUrl": {
      "defaultValue": "",
      "metadata": {
        "description": "The source of the generalized blob which will be used to create a custom windows image for the agent virtual machines."
      },
      "type": "string"
    },
    "agentWindowsPublisher": {
      "defaultValue": "MicrosoftWindowsServer",
      "metadata": {
        "description": "The publisher of windows image for the agent virtual machines."
      },
      "type": "string"
    },
    "agentWindowsOffer": {
      "defaultValue": "WindowsServerSemiAnnual",
      "metadata": {
        "description": "The offer of windows image for the agent virtual machines."
      },
      "type": "string"
    },
    "agentWindowsSku": {
      "defaultValue": "Datacenter-Core-1809-with-Containers-smalldisk",
      "metadata": {
        "description": "The SKU of windows image for the agent virtual machines."
      },
      "type": "string"
    },
    "windowsDockerVersion": {
      "defaultValue": "18.09.2",
      "metadata": {
        "description": "The version of Docker to be installed on Windows Nodes"
      },
      "type": "string"
    },
    "defaultContainerdRuntimeHandler": {
      "defaultValue": "process",
      "metadata": {
        "description": "The containerd handler type (process isolated or hyperv)"
      },
      "type": "string"
    },
    "hypervRuntimeHandlers": {
      "defaultValue": "",
      "metadata": {
        "description": "comma separated list of hyperv values"
      },
      "type": "string"
    }
