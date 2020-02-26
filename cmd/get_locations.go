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
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-06-01/subscriptions"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	locationsName             = "get-locations"
	locationsShortDescription = "Show Azure locations where an AKS Engine cluster can be deployed"
	locationsLongDescription  = locationsShortDescription + "\n\nThis command is unsupported and intended to be used only by AKS Engine developers."
)

var (
	chinaEast2DisplayName       = "China East 2"
	chinaEast2Name              = "chinaeast2"
	chinaEastDisplayName        = "China East"
	chinaEastName               = "chinaeast"
	chinaNorth2DisplayName      = "China North 2"
	chinaNorth2Name             = "chinanorth2"
	chinaNorthDisplayName       = "China North"
	chinaNorthName              = "chinanorth"
	centralUsEuapDisplayName    = "Central US EUAP (Canary)"
	centralUsEuapName           = "centraluseuap"
	eastUs2EuapDisplayName      = "East US 2 EUAP (Canary)"
	eastUs2EuapName             = "eastus2euap"
	germanyCentralDisplayName   = "Germany Central"
	germanyCentralName          = "germanycentral"
	germanyNortheastDisplayName = "Germany Northeast"
	germanyNortheastName        = "germanynortheast"
	usDodCentralDisplayName     = "US DoD Central"
	usDodCentralName            = "usdodcentral"
	usDodEastDisplayName        = "US Dod East"
	usDodEastName               = "usdodeast"
	usGovArizonaDisplayName     = "US Gov Arizona"
	usGovArizonaName            = "usgovarizona"
	usGovIowaDisplayName        = "US Gov Iowa"
	usGovIowaName               = "usgoviowa"
	usGovTexasDisplayName       = "US Gov Texas"
	usGovTexasName              = "usgovtexas"
	usGovVirginiaDisplayName    = "US Gov Virginia"
	usGovVirginiaName           = "usgovvirginia"
	notAvailable                = "N/A"
)

var locationsOutputFormatOptions = append(outputFormatOptions, "code")

type LocationsCmd struct {
	authProvider

	// user input
	output string

	// derived
	client armhelpers.AKSEngineClient
}

func newGetLocationsCmd() *cobra.Command {
	glc := LocationsCmd{
		authProvider: &authArgs{},
	}

	command := &cobra.Command{
		Use:    locationsName,
		Short:  locationsShortDescription,
		Long:   locationsLongDescription,
		Hidden: true,
		RunE:   glc.run,
	}

	f := command.Flags()
	getVersionsCmdDescription := fmt.Sprintf("Output format. Allowed values: %s",
		strings.Join(locationsOutputFormatOptions, ", "))
	f.StringVarP(&glc.output, "output", "o", "human", getVersionsCmdDescription)
	addAuthFlags(glc.getAuthArgs(), f)

	return command
}

func (glc *LocationsCmd) run(cmd *cobra.Command, args []string) error {
	log.Debugf("Start listing Azure locations")

	var err error

	// validate --output flag value before making API call
	outputFlagValid := false
	for _, opt := range locationsOutputFormatOptions {
		if glc.output == opt {
			outputFlagValid = true
			break
		}
	}
	if !outputFlagValid {
		return errors.New(fmt.Sprintf("invalid output format: \"%s\". Allowed values: %s.\n",
			glc.output, strings.Join(locationsOutputFormatOptions, ", ")))
	}

	if err = glc.getAuthArgs().validateAuthArgs(); err != nil {
		return err
	}

	if glc.client, err = glc.authProvider.getClient(); err != nil {
		return errors.Wrap(err, "failed to get client")
	}

	ctx, cancel := context.WithTimeout(context.Background(), armhelpers.DefaultARMOperationTimeout)
	defer cancel()
	list, err := glc.client.ListLocations(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to list locations")
	}
	locations := *list

	specialLocations := []subscriptions.Location{
		// Azure China Cloud locations
		{
			ID:          &notAvailable,
			Name:        &chinaEastName,
			DisplayName: &chinaEastDisplayName,
			Latitude:    &notAvailable,
			Longitude:   &notAvailable,
		},
		{
			ID:          &notAvailable,
			Name:        &chinaEast2Name,
			DisplayName: &chinaEast2DisplayName,
			Latitude:    &notAvailable,
			Longitude:   &notAvailable,
		},
		{
			ID:          &notAvailable,
			Name:        &chinaNorthName,
			DisplayName: &chinaNorthDisplayName,
			Latitude:    &notAvailable,
			Longitude:   &notAvailable,
		},
		{
			ID:          &notAvailable,
			Name:        &chinaNorth2Name,
			DisplayName: &chinaNorth2DisplayName,
			Latitude:    &notAvailable,
			Longitude:   &notAvailable,
		},
		// Canary locations
		{
			ID:          &notAvailable,
			Name:        &centralUsEuapName,
			DisplayName: &centralUsEuapDisplayName,
			Latitude:    &notAvailable,
			Longitude:   &notAvailable,
		},
		{
			ID:          &notAvailable,
			Name:        &eastUs2EuapName,
			DisplayName: &eastUs2EuapDisplayName,
			Latitude:    &notAvailable,
			Longitude:   &notAvailable,
		},
		// US DoD locations
		{
			ID:          &notAvailable,
			Name:        &usDodCentralName,
			DisplayName: &usDodCentralDisplayName,
			Latitude:    &notAvailable,
			Longitude:   &notAvailable,
		},
		{
			ID:          &notAvailable,
			Name:        &usDodEastName,
			DisplayName: &usDodEastDisplayName,
			Latitude:    &notAvailable,
			Longitude:   &notAvailable,
		},
		// US Gov locations
		{
			ID:          &notAvailable,
			Name:        &usGovArizonaName,
			DisplayName: &usGovArizonaDisplayName,
			Latitude:    &notAvailable,
			Longitude:   &notAvailable,
		},
		{
			ID:          &notAvailable,
			Name:        &usGovIowaName,
			DisplayName: &usGovIowaDisplayName,
			Latitude:    &notAvailable,
			Longitude:   &notAvailable,
		},
		{
			ID:          &notAvailable,
			Name:        &usGovTexasName,
			DisplayName: &usGovTexasDisplayName,
			Latitude:    &notAvailable,
			Longitude:   &notAvailable,
		},
		{
			ID:          &notAvailable,
			Name:        &usGovVirginiaName,
			DisplayName: &usGovVirginiaDisplayName,
			Latitude:    &notAvailable,
			Longitude:   &notAvailable,
		},
		// Germany locations
		{
			ID:          &notAvailable,
			Name:        &germanyCentralName,
			DisplayName: &germanyCentralDisplayName,
			Latitude:    &notAvailable,
			Longitude:   &notAvailable,
		},
		{
			ID:          &notAvailable,
			Name:        &germanyNortheastName,
			DisplayName: &germanyNortheastDisplayName,
			Latitude:    &notAvailable,
			Longitude:   &notAvailable,
		},
	}
	// Add special locations if they aren't already in the list
	for _, s := range specialLocations {
		found := false
		for _, l := range locations {
			if l.Name == s.Name {
				found = true
				break
			}
		}
		if !found {
			locations = append(locations, s)
		}
	}

	// Sort the locations by name
	sort.Slice(locations, func(i, j int) bool {
		return *locations[i].Name < *locations[j].Name
	})

	switch glc.output {
	case "json":
		data, jsonErr := helpers.JSONMarshalIndent(locations, "", "  ", false)
		if jsonErr != nil {
			return jsonErr
		}
		fmt.Println(string(data))
	case "code":
		b := strings.Builder{}
		b.WriteString(`// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package helpers

// GetAzureLocations provides all available Azure cloud locations.
//
// Code generated for package helpers by aks-engine DO NOT EDIT. (@generated)
//
// To generate this code, run the command:
//   aks-engine get-locations --output=code
func GetAzureLocations() []string {
	return []string{
`)
		for _, l := range locations {
			b.WriteString(fmt.Sprintf("\t\t\"%s\",\n", *l.Name))
		}
		b.WriteString("\t}\n}")
		fmt.Println(b.String())
	case "human":
		w := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', tabwriter.FilterHTML)
		fmt.Fprintln(w, "Location\tName\tLatitude\tLongitude")
		for _, location := range locations {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
				*location.Name, *location.DisplayName, *location.Latitude, *location.Longitude)
		}
		w.Flush()
	}

	log.Debugf("Done listing Azure locations")

	return err
}
