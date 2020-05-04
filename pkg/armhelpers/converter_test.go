// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.
package armhelpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/cmplx"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2017-03-30/compute"
	azcompute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-12-01/compute"
	"github.com/google/go-cmp/cmp"
)

func TestDeepCopyManagedDisk(t *testing.T) {
	diskJSONAzure := "./converterTestData/disksAZ.json"
	diskJSONAzureStack := "./converterTestData/disksAZS.json"
	diskAzure := []azcompute.Disk{}
	err := unmarshalFromFile(diskJSONAzure, &diskAzure)
	if err != nil {
		t.Error(err)
	}

	diskAzureStack := []compute.Disk{}
	err = unmarshalFromFile(diskJSONAzureStack, &diskAzureStack)
	if err != nil {
		t.Error(err)
	}

	diskAzureConvert := []azcompute.Disk{}
	if err = DeepCopy(&diskAzureConvert, diskAzureStack); err != nil {
		t.Error(err)
	}

	for _, expect := range diskAzure {
		match := false
		for _, actual := range diskAzureConvert {
			if *actual.ID == *expect.ID {
				match = true
				if *actual.Name != *expect.Name {
					t.Errorf("Fail to compare, Name expect : %q -- got : %q", *expect.Name, *actual.Name)
				}
				if *actual.Location != *expect.Location {
					t.Errorf("Fail to compare, Location expect : %q -- got : %q", *expect.Location, *actual.Location)
				}
				if *actual.ManagedBy != *expect.ManagedBy {
					t.Errorf("Fail to compare, ManagedBy expect : %q -- got : %q", *expect.ManagedBy, *actual.ManagedBy)
				}
				if diff := cmp.Diff(actual.Tags, expect.Tags); diff != "" {
					t.Errorf("Fail to compare, Tags %q", diff)
				}
			}
		}
		if !match {
			t.Errorf("Fail to compare, expect : %q ", *expect.ID)
		}
	}
}

func TestDeepCopyVirtualMachine(t *testing.T) {
	vmsJSONAzure := "./converterTestData/vmsAZ.json"
	vmsJSONAzureStack := "./converterTestData/vmsAZS.json"
	vmsAzure := []azcompute.VirtualMachine{}
	err := unmarshalFromFile(vmsJSONAzure, &vmsAzure)
	if err != nil {
		t.Error(err)
	}

	vmsAzureStack := []compute.VirtualMachine{}
	err = unmarshalFromFile(vmsJSONAzureStack, &vmsAzureStack)
	if err != nil {
		t.Error(err)
	}

	vmsAzureConvert := []azcompute.VirtualMachine{}
	if err = DeepCopy(&vmsAzureConvert, vmsAzureStack); err != nil {
		t.Error(err)
	}

	for _, expect := range vmsAzure {
		match := false
		for _, actual := range vmsAzureConvert {
			if *actual.ID == *expect.ID {
				match = true
				if *actual.Name != *expect.Name {
					t.Errorf("Fail to compare, Name expect : %q -- got : %q", *expect.Name, *actual.Name)
				}
				if diff := cmp.Diff(actual.Tags, expect.Tags); diff != "" {
					t.Errorf("Fail to compare, Tags %q", diff)
				}
				if diff := cmp.Diff(actual.StorageProfile, expect.StorageProfile); diff != "" {
					t.Errorf("Fail to compare, StorageProfile %q", diff)
				}
			}
		}
		if !match {
			t.Errorf("Fail to compare, expect : %q ", *expect.ID)
		}
	}
}

func TestDeepCopyVMScaleSet(t *testing.T) {
	vmssJSONAzure := "./converterTestData/vmssAZ.json"
	vmssJSONAzureStack := "./converterTestData/vmssAZS.json"
	vmssAzure := []azcompute.VirtualMachineScaleSet{}
	err := unmarshalFromFile(vmssJSONAzure, &vmssAzure)
	if err != nil {
		t.Error(err)
	}

	vmssAzureStack := []compute.VirtualMachineScaleSet{}
	err = unmarshalFromFile(vmssJSONAzureStack, &vmssAzureStack)
	if err != nil {
		t.Error(err)
	}

	vmssAzureConvert := []azcompute.VirtualMachineScaleSet{}
	if err = DeepCopy(&vmssAzureConvert, vmssAzureStack); err != nil {
		t.Error(err)
	}

	for _, expect := range vmssAzure {
		match := false
		for _, actual := range vmssAzureConvert {
			if *actual.ID == *expect.ID {
				match = true
				if *actual.Name != *expect.Name {
					t.Errorf("Fail to compare, Name expect : %q -- got : %q", *expect.Name, *actual.Name)
				}
				if *actual.Location != *expect.Location {
					t.Errorf("Fail to compare, Name expect : %q -- got : %q", *expect.Name, *actual.Name)
				}
				if diff := cmp.Diff(actual.Sku, expect.Sku); diff != "" {
					t.Errorf("Fail to compare, Sku %q", diff)
				}
				if diff := cmp.Diff(actual.Tags, expect.Tags); diff != "" {
					t.Errorf("Fail to compare, Tags %q", diff)
				}
			}
		}
		if !match {
			t.Errorf("Fail to compare, expect : %q ", *expect.ID)
		}
	}
}

func TestDeepCopyVMScaleSetVM(t *testing.T) {
	vmssvmJSONAzure := "./converterTestData/vmssvmAZ.json"
	vmssvmJSONAzureStack := "./converterTestData/vmssvmAZS.json"
	vmssvmAzure := []azcompute.VirtualMachineScaleSetVM{}
	err := unmarshalFromFile(vmssvmJSONAzure, &vmssvmAzure)
	if err != nil {
		t.Error(err)
	}

	vmssvmAzureStack := []compute.VirtualMachineScaleSetVM{}
	err = unmarshalFromFile(vmssvmJSONAzureStack, &vmssvmAzureStack)
	if err != nil {
		t.Error(err)
	}

	vmssvmAzureConvert := []azcompute.VirtualMachineScaleSetVM{}
	if err = DeepCopy(&vmssvmAzureConvert, vmssvmAzureStack); err != nil {
		t.Error(err)
	}

	for _, expect := range vmssvmAzure {
		match := false
		for _, actual := range vmssvmAzureConvert {
			if *actual.ID == *expect.ID {
				match = true
				if *actual.Name != *expect.Name {
					t.Errorf("Fail to compare, Name expect : %q -- got : %q", *expect.Name, *actual.Name)
				}
				if *actual.Location != *expect.Location {
					t.Errorf("Fail to compare, Name expect : %q -- got : %q", *expect.Name, *actual.Name)
				}
				if diff := cmp.Diff(actual.StorageProfile, expect.StorageProfile); diff != "" {
					t.Errorf("Fail to compare, StorageProfile %q", diff)
				}
				if diff := cmp.Diff(actual.Sku, expect.Sku); diff != "" {
					t.Errorf("Fail to compare, Sku %q", diff)
				}
				if diff := cmp.Diff(actual.Tags, expect.Tags); diff != "" {
					t.Errorf("Fail to compare, Tags %q", diff)
				}
			}
		}
		if !match {
			t.Errorf("Fail to compare, expect : %q ", *expect.ID)
		}
	}
}

type Sample struct {
	Uint        uint
	Uint8       uint8
	Uint16      uint16
	Unit32      uint32
	Unit64      uint64
	Float32     float32
	Float64     float64
	Complex64   complex64
	Complex128  complex128
	StringArray [5]string
}

func TestDeepCopySampleStruct(t *testing.T) {
	s := Sample{
		Uint:        10,
		Uint8:       16,
		Uint16:      256,
		Unit32:      65536,
		Unit64:      4294967296,
		Float32:     2.147483646,
		Float64:     2.1474922337203685477580783646,
		Complex64:   complex(5, 7),
		Complex128:  cmplx.Sqrt(-5 + 12i),
		StringArray: [5]string{"one", "two", "three", "four", "five"},
	}
	sConvert := Sample{}
	if err := DeepCopy(&sConvert, s); err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(s, sConvert); diff != "" {
		t.Errorf("Fail to compare, Sample %q", diff)
	}
}

func unmarshalFromFile(filePath string, v interface{}) error {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Fail to read file %q , err -  %q", filePath, err)
	}
	err = json.Unmarshal(bytes, v)
	if err != nil {
		return fmt.Errorf("Fail to unmarshal file %q , err -  %q", filePath, err)
	}
	return nil
}

func unmarshalFromString(jsonString string, v interface{}) error {
	err := json.Unmarshal([]byte(jsonString), v)
	if err != nil {
		return fmt.Errorf("Fail to unmarshal, err -  %q", err)
	}
	return nil
}
