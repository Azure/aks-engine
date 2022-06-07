// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package main

import (
	"os"

	"github.com/Azure/aks-engine/cmd"
	colorable "github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(colorable.NewColorableStderr())
	msg := `
aks-engine is deprecated for Azure public cloud customers. Learn more at:
    https://github.com/Azure/aks-engine#project-status

Please consider using Azure Kubernetes Service (AKS) for managed Kubernetes:
    https://azure.microsoft.com/services/kubernetes-service/
or Cluster API Provider Azure (CAPZ) for self-managed Kubernetes:
    https://capz.sigs.k8s.io/
`
	log.Warningf("\u001b[33m%s\u001b[0m", msg)
	log.SetOutput(colorable.NewColorableStdout())
	if err := cmd.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
