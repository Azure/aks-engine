# How to build VHDs

## Windows VHD

To build the Windows VHD

1. Enter our dev container

    ```bash
    make dev
    ```

1. Set the following environment varibles

    |   |   |
    |---|---|
    | CLIENT_ID | Azure service princicle ID |
    | CLIENT_SECRET | Azure service principle secret |
    | TENANT_ID | Azure tennant ID |
    |  AZURE_RESOURCE_GROUP_NAME | Resource group name to use / create |
    | AZURE_LOCATION | Azure region to use |
    | AZURE_VM_SIZE | VM size packer will use (Standard_D4s_v3 is recommended) |
    | CONTAINER_RUNTIME | Container runtime VHD will be generated for (Docker or Containerd) |
    | WINDOWS_SERVER_VERSION | Windows Server OS version vHD will be generated for (2019 or 2004) |

1. Run packer

   ```bash
   make run-packer-windows
   ```

Once complete, packer will output a URI to the VHD it created.
