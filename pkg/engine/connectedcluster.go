// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/Azure/aks-engine/pkg/api"
)

func connectedClusterOnboardingIntoReaders(p *api.Properties) ([]CustomFileReader, error) {
	if p.ConnectedClusterProfile == nil {
		return []CustomFileReader{}, nil
	}

	versions := strings.Split(p.OrchestratorProfile.OrchestratorVersion, ".")
	componentFile := getCustomDataFilePath(connectedClusterOnboardingSourceFilename, "k8s/cloud-init", versions[0]+"."+versions[1])
	componentFileBytes, err := Asset(componentFile)
	if err != nil {
		return []CustomFileReader{}, err
	}

	templ := template.New("component resolver template").Funcs(getConnectedClusterOnboardingFuncMap(p.ConnectedClusterProfile))
	_, err = templ.Parse(string(componentFileBytes))
	if err != nil {
		return []CustomFileReader{}, err
	}

	var buffer bytes.Buffer
	_ = templ.Execute(&buffer, p.ConnectedClusterProfile)
	customFileReader := CustomFileReader{
		Source: bytes.NewReader(buffer.Bytes()),
		Dest:   fmt.Sprintf("/opt/azure/containers/init/%s", connectedClusterOnboardingDestinationFilename),
	}
	return []CustomFileReader{customFileReader}, nil
}
