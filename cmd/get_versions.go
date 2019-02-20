// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/spf13/cobra"
)

const (
	getVersionsName             = "get-versions"
	getVersionsShortDescription = "Display info about supported Kubernetes versions"
	getVersionsLongDescription  = "Display supported Kubernetes versions and upgrade versions"
)

type getVersionsCmd struct {
	// user input
	orchestrator string
	version      string
	windows      bool
	output       string
}

func newGetVersionsCmd() *cobra.Command {
	gvc := getVersionsCmd{}

	command := &cobra.Command{
		Use:   getVersionsName,
		Short: getVersionsShortDescription,
		Long:  getVersionsLongDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			return gvc.run(cmd, args)
		},
	}

	f := command.Flags()
	gvc.orchestrator = "Kubernetes" // orchestrator is always Kubernetes
	f.StringVar(&gvc.version, "version", "", "Kubernetes version (optional)")
	f.BoolVar(&gvc.windows, "windows", false, "Kubernetes cluster with Windows nodes (optional)")
	getVersionsCmdDescription := fmt.Sprintf("Output format. Allowed values: %s",
		strings.Join(outputFormatOptions, ", "))
	f.StringVarP(&gvc.output, "output", "o", "human", getVersionsCmdDescription)

	return command
}

func (gvc *getVersionsCmd) run(cmd *cobra.Command, args []string) error {
	orchs, err := api.GetOrchestratorVersionProfileListVLabs(gvc.orchestrator, gvc.version, gvc.windows)
	if err != nil {
		return err
	}

	switch gvc.output {
	case "json":
		data, err := helpers.JSONMarshalIndent(orchs, "", "  ", false)
		if err != nil {
			return err
		}
		fmt.Println(string(data))
	case "human":
		w := tabwriter.NewWriter(os.Stdout, 0, 4, 1, ' ', tabwriter.FilterHTML)
		fmt.Fprintln(w, "Version\tUpgrades")
		// iterate in reverse so the newest Kubernetes release is listed first
		for i := len(orchs.Orchestrators) - 1; i >= 0; i-- {
			o := orchs.Orchestrators[i]
			fmt.Fprintf(w, "%s\t", o.OrchestratorVersion)
			// collapse the upgrade fields into a comma-separated list
			lenUpgrades := len(o.Upgrades) - 1
			for j := 0; j < len(o.Upgrades); j++ {
				u := o.Upgrades[j]
				fmt.Fprintf(w, "%s", u.OrchestratorVersion)
				if j < lenUpgrades {
					fmt.Fprintf(w, ", ")
				}
			}
			fmt.Fprintln(w)
		}
		w.Flush()
	default:
		return fmt.Errorf("output format \"%s\" is not supported", gvc.output)
	}

	return nil
}
