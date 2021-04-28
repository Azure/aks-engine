// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/Azure/aks-engine/pkg/armhelpers"
	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	developersOnly       = "\n\nThis command is unsupported and intended to be used only by AKS Engine developers."
	skusName             = "get-skus"
	skusShortDescription = "Show Azure VM SKUs that can be used to deploy an AKS Engine cluster"
	skusLongDescription  = skusShortDescription + developersOnly
)

var skusOutputFormatOptions = append(outputFormatOptions, "code")

type SkusCmd struct {
	authProvider

	// user input
	output string

	// derived
	client armhelpers.AKSEngineClient
}

func newGetSkusCmd() *cobra.Command {
	vmc := SkusCmd{
		authProvider: &authArgs{},
	}

	command := &cobra.Command{
		Use:    skusName,
		Short:  skusShortDescription,
		Long:   skusLongDescription,
		Hidden: true,
		RunE:   vmc.run,
	}

	f := command.Flags()
	outputFlagDescription := fmt.Sprintf("Output format. Allowed values: %s",
		strings.Join(skusOutputFormatOptions, ", "))
	f.StringVarP(&vmc.output, "output", "o", "human", outputFlagDescription)
	addAuthFlags(vmc.getAuthArgs(), f)

	return command
}

func (vmc *SkusCmd) run(cmd *cobra.Command, args []string) error {
	log.Debugf("Start listing VM SKUs")

	var err error

	// validate --output flag value before making API call
	outputFlagValid := false
	for _, opt := range skusOutputFormatOptions {
		if vmc.output == opt {
			outputFlagValid = true
			break
		}
	}
	if !outputFlagValid {
		return errors.New(fmt.Sprintf("invalid output format: \"%s\". Allowed values: %s.\n",
			vmc.output, strings.Join(skusOutputFormatOptions, ", ")))
	}

	if err = vmc.getAuthArgs().validateAuthArgs(); err != nil {
		return err
	}

	if vmc.client, err = vmc.authProvider.getClient(); err != nil {
		return errors.Wrap(err, "failed to get client")
	}

	// List Resource SKUs
	ctx, cancel := context.WithTimeout(context.Background(), armhelpers.DefaultARMOperationTimeout)
	defer cancel()
	page, err := vmc.client.ListResourceSkus(ctx, "")
	if err != nil {
		return errors.Wrap(err, "failed to list resource SKUs")
	}
	skus := helpers.VMSkus
	if page != nil {
		for _, r := range page.Values() {
			name := *r.Name
			if !strings.HasPrefix(name, "Standard_") || strings.HasSuffix(name, "_Promo") {
				continue
			}
			// Add to the list if the SKU isn't already present
			found := false
			for _, s := range skus {
				if name == s.Name {
					found = true
					break
				}
			}
			if !found {
				acceleratedNetworking := false
				if r.Capabilities != nil {
					for _, c := range *r.Capabilities {
						if c.Name != nil && *c.Name == "AcceleratedNetworkingEnabled" {
							if c.Value != nil && strings.EqualFold(*c.Value, "True") {
								acceleratedNetworking = true
							}
						}
					}
				}
				skus = append(skus, helpers.VMSku{
					Name:                  name,
					AcceleratedNetworking: acceleratedNetworking,
				})
			}
		}
	}

	// Sort the SKUs by name
	sort.Slice(skus, func(i, j int) bool {
		return skus[i].Name < skus[j].Name
	})

	switch vmc.output {
	case "json":
		data, jsonErr := helpers.JSONMarshalIndent(skus, "", "  ", false)
		if jsonErr != nil {
			return jsonErr
		}
		fmt.Println(string(data))
	case "code":
		b := strings.Builder{}
		b.WriteString(`// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package helpers

// GetKubernetesAllowedVMSKUs provides the allowed sizes for Kubernetes agent VMs.
//
// Code generated for package helpers by aks-engine DO NOT EDIT. (@generated)
//
// To generate this code, run the command:
//   aks-engine get-skus --output=code

type VMSku struct {
	Name                  string
	AcceleratedNetworking bool
}

var VMSkus = []VMSku{
`)
		formatStr := "\t{\n\t\tName:                  \"%s\",\n\t\tAcceleratedNetworking: %t,\n\t},\n"
		for _, s := range skus {
			b.WriteString(fmt.Sprintf(formatStr, s.Name, s.AcceleratedNetworking))
		}
		b.WriteString("}")
		fmt.Println(b.String())
	case "human":
		w := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', tabwriter.FilterHTML)
		fmt.Fprintln(w, "Name\tAccelerated Networking Support")
		for _, sku := range skus {
			fmt.Fprintf(w, "%s\t%t\n", sku.Name, sku.AcceleratedNetworking)
		}
		w.Flush()
	}

	log.Debugf("Done listing VM SKUs")

	return err
}
