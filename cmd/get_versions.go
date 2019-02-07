// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"github.com/spf13/cobra"
)

const (
	getVersionsName             = "get-versions"
	getVersionsShortDescription = "Display info about supported Kubernetes versions"
	getVersionsLongDescription  = "Display supported Kubernetes versions and upgrade versions"
)

func newGetVersionsCmd() *cobra.Command {
	oc := orchestratorsCmd{}

	command := &cobra.Command{
		Use:   getVersionsName,
		Short: getVersionsShortDescription,
		Long:  getVersionsLongDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			return oc.run(cmd, args)
		},
	}

	f := command.Flags()
	oc.orchestrator = "Kubernetes" // orchestrator is always Kubernetes
	f.StringVar(&oc.version, "version", "", "Kubernetes version (optional)")
	f.BoolVar(&oc.windows, "windows", false, "Kubernetes cluster with Windows nodes (optional)")

	return command
}
