// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package openshift

import (
	"path"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/Azure/aks-engine/pkg/i18n"
)

func TestExamplesInSync(t *testing.T) {
	baseExampleDir := "../../examples"
	baseExample := "openshift.json"

	tests := []string{}

	locale, err := i18n.LoadTranslations()
	if err != nil {
		t.Fatalf("error loading translations %v", err)
	}

	apiloader := &api.Apiloader{
		Translator: &i18n.Translator{
			Locale: locale,
		},
	}

	baseExampleCS, _, err := apiloader.LoadContainerServiceFromFile(
		path.Join(baseExampleDir, baseExample),
		false, //don't validate
		false, // not update
		nil,
	)
	if err != nil {
		t.Fatalf("error parsing the api model: %s", err.Error())
	}

	for _, test := range tests {
		testCS, _, err := apiloader.LoadContainerServiceFromFile(
			path.Join(baseExampleDir, test),
			false,
			false,
			nil,
		)
		if err != nil {
			t.Errorf("failed parsing %s: %#v", test, err)
			continue
		}

		// todo normalize where necessary (seems easier than reflect right now)

		if !reflect.DeepEqual(baseExampleCS.Properties, testCS.Properties) {
			t.Errorf(spew.Sprintf("Testing %s\nExpected:\n%+v\nGot:\n%+v", test, baseExampleCS.Properties, testCS.Properties))
		}
	}

}
