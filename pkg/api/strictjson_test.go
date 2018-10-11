package api

import (
	"reflect"
	"strings"
	"testing"

	"github.com/Azure/aks-engine/pkg/api/v20160930"
	"github.com/Azure/aks-engine/pkg/api/v20170131"
	"github.com/Azure/aks-engine/pkg/api/v20170701"
	"github.com/Azure/aks-engine/pkg/api/vlabs"
)

type SubTestProfile struct {
	SP1 bool             `json:"sp1"`
	SP2 bool             `json:"sp2"`
	SP3 []SubTestProfile `json:"sp3,omitempty"`
}

type TestProfile struct {
	Field1 int               `json:"f1"`
	Field2 SubTestProfile    `json:"f2,omitempty"`
	Field3 []SubTestProfile  `json:"f3,omitempty"`
	Field4 *SubTestProfile   `json:"f4"`
	Field5 []*SubTestProfile `json:"f5,omitempty"`
}

func TestIfAllJSONKeysAreExpectedThenCheckPasses(t *testing.T) {
	json := `
	{
		"f1": 1,
		"f2": {
			"sp1": true,
			"sp2": false
		}
	}
	`
	e := checkJSONKeys([]byte(json), reflect.TypeOf(TestProfile{}))
	if e != nil {
		t.Errorf("All JSON keys were expected but the check still failed: %v", e)
	}
}

func TestIsCaseInsensitive(t *testing.T) {
	json := `
	{
		"F1": 1,
		"f2": {
			"sP1": true,
			"Sp2": false
		}
	}
	`
	e := checkJSONKeys([]byte(json), reflect.TypeOf(TestProfile{}))
	if e != nil {
		t.Errorf("All JSON keys were expected (allowing for case) but the check still failed: %v", e)
	}
}

func TestCheckFailsOnUnexpectedJSONKeyAtTopLevel(t *testing.T) {
	json := `
	{
		"f2": {
			"sp1": true,
			"sp2": false
		},
		"fx": "uh-oh"
	}
	`
	e := checkJSONKeys([]byte(json), reflect.TypeOf(TestProfile{}))
	if e == nil {
		t.Fatal("Unexpected JSON key was not detected")
	}
	if !strings.Contains(e.Error(), "fx") {
		t.Errorf("Error message did not name unexpected JSON key 'fx': was %v", e)
	}
}

func TestCheckFailsOnUnexpectedJSONKeyAtSubLevel(t *testing.T) {
	json := `
	{
		"f2": {
			"sp1": true,
			"spx": false
		}
	}
	`
	e := checkJSONKeys([]byte(json), reflect.TypeOf(TestProfile{}))
	if e == nil {
		t.Fatal("Unexpected JSON key was not detected")
	}
	if !strings.Contains(e.Error(), "spx") {
		t.Errorf("Error message did not name unexpected JSON key 'spx': was %v", e) // TODO: f2.spx would be better
	}
}

func TestCheckFailsOnUnexpectedJSONKeyInArray(t *testing.T) {
	json := `
	{
		"f2": {
			"sp1": true,
			"sp3": [
				{
					"sp1": false,
					"sp2": true
				},
				{
					"sp2": false,
					"spz": "unexpected"
				}
			]
		},
		"f3": []
	}
	`
	e := checkJSONKeys([]byte(json), reflect.TypeOf(TestProfile{}))
	if e == nil {
		t.Fatal("Unexpected JSON key was not detected")
	}
	if !strings.Contains(e.Error(), "spz") {
		t.Errorf("Error message did not name unexpected JSON key 'spz': was %v", e) // TODO: f2[1].spz might be better
	}
}

func TestCheckFailsOnUnexpectedJSONKeyInArrayAtSubLevel(t *testing.T) {
	json := `
	{
		"f2": {
			"sp1": true
		},
		"f3": [
			{
				"sp1": true,
				"spy": "unexpected"
			},
			{
				"sp2": false,
				"spz": "unexpected"
			}
		]
	}
	`
	e := checkJSONKeys([]byte(json), reflect.TypeOf(TestProfile{}))
	if e == nil {
		t.Fatal("Unexpected JSON key was not detected")
	}
	if !strings.Contains(e.Error(), "spy") {
		t.Errorf("Error message did not name unexpected JSON key 'spy': was %v", e) // TODO: f3[0].spy might be better
	}
}

func TestCheckFailsOnUnexpectedJSONKeyAtSubLevelViaPointer(t *testing.T) {
	json := `
	{
		"f4": {
			"sp1": true,
			"spx": false
		}
	}
	`
	e := checkJSONKeys([]byte(json), reflect.TypeOf(TestProfile{}))
	if e == nil {
		t.Fatal("Unexpected JSON key was not detected")
	}
	if !strings.Contains(e.Error(), "spx") {
		t.Errorf("Error message did not name unexpected JSON key 'spx': was %v", e) // TODO: f4.spx would be better
	}
}

func TestCheckFailsOnUnexpectedJSONKeyInArrayAtSubLevelViaPointer(t *testing.T) {
	json := `
	{
		"f2": {
			"sp1": true
		},
		"f5": [
			{
				"sp1": true,
				"spy": "unexpected"
			},
			{
				"sp2": false,
				"spz": "unexpected"
			}
		]
	}
	`
	e := checkJSONKeys([]byte(json), reflect.TypeOf(TestProfile{}))
	if e == nil {
		t.Fatal("Unexpected JSON key was not detected")
	}
	if !strings.Contains(e.Error(), "spy") {
		t.Errorf("Error message did not name unexpected JSON key 'spy': was %v", e) // TODO: f3[0].spy might be better
	}
}

const jsonWithTypo = `
{
	"apiVersion": "ignored",
	"properties": {
	  "orchestratorProfile": {
		"orchestratorType": "Kubernetes"
	  },
	  "masterProfile": {
		"count": 1,
		"dnsprefix": "masterdns1",
		"vmsize": "Standard_D2_v2",
		"ventSubnetID": "/this/attribute/was/mistyped"
	  },
	  "agentPoolProfiles": [],
	  "linuxProfile": {
		"adminUsername": "azureuser",
		"ssh": {
		  "publicKeys": [
			{
			  "keyData": "ssh-rsa PUBLICKEY azureuser@linuxvm"
			}
		  ]
		}
	  },
	  "servicePrincipalProfile": {
		"clientId": "ServicePrincipalClientID",
		"secret": "myServicePrincipalClientSecret"
	  }
	}
}
`

func TestStrictJSONValidationIsNotAppliedToApiVersions20170701AndEarlier(t *testing.T) {
	// These API versions existed before we added strict JSON validation:
	// we cannot apply it retrospectively because it could break existing
	// customer apimodels.
	preStrictVersions := []string{v20160930.APIVersion, v20170131.APIVersion, v20170701.APIVersion /*, v20170930.APIVersion */}
	a := &Apiloader{
		Translator: nil,
	}

	for _, version := range preStrictVersions {
		_, e := a.LoadContainerService([]byte(jsonWithTypo), version, true, false, nil)
		if e != nil {
			t.Errorf("Expected mistyped 'ventSubnetID' key to be overlooked in version '%s' but it wasn't: error was %v", version, e)
		}
	}
}

func TestStrictJSONValidationIsAppliedToVersionsAbove20170701(t *testing.T) {
	strictVersions := []string{vlabs.APIVersion}
	a := &Apiloader{
		Translator: nil,
	}
	for _, version := range strictVersions {
		_, e := a.LoadContainerService([]byte(jsonWithTypo), version, true, false, nil)
		if e == nil {
			t.Error("Expected mistyped 'ventSubnetID' key to be detected but it wasn't")
		} else {
			if !strings.Contains(e.Error(), "ventSubnetID") {
				t.Errorf("Expected error on 'ventSubnetID' but error was %v", e)
			}
		}
	}
}
