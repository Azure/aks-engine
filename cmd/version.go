// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/Azure/aks-engine/pkg/helpers"

	"github.com/spf13/cobra"
)

var (
	// BuildTag holds the `git tag` if this is a tagged build/release
	BuildTag = "canary"

	// BuildSHA holds the git commit SHA at `make build` time.
	BuildSHA = "unset"

	// GitTreeState is the state of the git tree, either clean or dirty
	GitTreeState = "unset"

	outputFormatOptions = []string{"human", "json"}
	outputFormat        string
	version             versionInfo
)

const (
	versionName             = "version"
	versionShortDescription = "Print the version of aks-engine"
	versionLongDescription  = versionShortDescription
)

type versionInfo struct {
	GitTag       string
	GitCommit    string
	GitTreeState string
}

func init() {
	version = versionInfo{
		GitTag:       BuildTag,
		GitCommit:    BuildSHA,
		GitTreeState: GitTreeState,
	}
}

func getHumanVersion() string {
	return fmt.Sprintf("Version: %s\nGitCommit: %s\nGitTreeState: %s",
		version.GitTag,
		version.GitCommit,
		version.GitTreeState)
}

func getJSONVersion() string {
	jsonVersion, _ := helpers.JSONMarshalIndent(version, "", "  ", false)
	return string(jsonVersion)
}

func getVersion(outputType string) (string, error) {
	switch outputType {
	case "human":
		return getHumanVersion(), nil
	case "json":
		return getJSONVersion(), nil
	default:
		return "", errors.Errorf(`output format "%s" is not supported`, outputType)
	}
}

func newVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   versionName,
		Short: versionShortDescription,
		Long:  versionLongDescription,

		RunE: func(cmd *cobra.Command, args []string) error {
			output, err := getVersion(outputFormat)
			if err == nil {
				fmt.Println(output)
			}
			return err
		},
	}

	versionCmdDescription := fmt.Sprintf("Output format. Allowed values: %s",
		strings.Join(outputFormatOptions, ", "))

	versionCmd.Flags().StringVarP(&outputFormat, "output", "o", "human", versionCmdDescription)

	return versionCmd
}
