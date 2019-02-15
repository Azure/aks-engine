// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"github.com/spf13/cobra"
)

const (
	orchestratorsName             = "orchestrators"
	orchestratorsShortDescription = "Display info about supported orchestrators"
	orchestratorsLongDescription  = "Display supported versions and upgrade versions for each orchestrator"
)

func newOrchestratorsCmd() *cobra.Command {
	gvc := getVersionsCmd{}

	command := &cobra.Command{
		Use:   orchestratorsName,
		Short: orchestratorsShortDescription,
		Long:  orchestratorsLongDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			return gvc.run(cmd, args)
		},
		Hidden: true,
	}

	f := command.Flags()
	f.StringVar(&gvc.orchestrator, "orchestrator", "", "orchestrator name (optional) ")
	f.StringVar(&gvc.version, "version", "", "orchestrator version (optional)")
	f.BoolVar(&gvc.windows, "windows", false, "orchestrator platform (optional, applies to Kubernetes only)")
	gvc.output = "json" // output is always JSON

	return command
}
