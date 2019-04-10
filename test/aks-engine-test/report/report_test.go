// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package report

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestReportParse(t *testing.T) {
	jobName := "TestJob"
	buildNum := 001
	nDeploys := 4
	fileName := "../aks-engine-errors.json"
	dummy := New(jobName, buildNum, nDeploys, fileName)

	txt := "Error loading command module"
	step := "step"
	testName := "dummyTest"
	d := "westus"
	_ = dummy.Process(txt, step, testName, d)

	testReport := "TestReport.json"
	if err := dummy.CreateTestReport(testReport); err != nil {
		t.Fatal(err)
	}

	raw, err := ioutil.ReadFile(testReport)
	if err != nil {
		t.Fatal(err)
	}

	h := &Manager{}
	if err = json.Unmarshal(raw, &h); err != nil {
		t.Fatalf("Error unmarshalling JSON: %s", err)
	}

	if len(h.LogErrors.LogErrors) != 0 {
		t.Fatalf("Expected LogErrors to be empty, instead it is of size %d", len(h.LogErrors.LogErrors))
	}
}
