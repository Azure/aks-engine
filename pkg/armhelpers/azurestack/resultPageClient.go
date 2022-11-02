// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package azurestack

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2017-03-30/compute"
	azcompute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-12-01/compute"
)

// VirtualMachineListResultPageClient Virtual Machine List Result Page Client
type VirtualMachineListResultPageClient struct {
	vmlrp compute.VirtualMachineListResultPage
	err   error
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *VirtualMachineListResultPageClient) NextWithContext(ctx context.Context) (err error) {
	return page.vmlrp.NextWithContext(ctx)
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *VirtualMachineListResultPageClient) Next() error {
	return page.vmlrp.Next()
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page VirtualMachineListResultPageClient) NotDone() bool {
	return page.vmlrp.NotDone()
}

// Response returns the raw server response from the last page request.
func (page VirtualMachineListResultPageClient) Response() azcompute.VirtualMachineListResult {
	r := azcompute.VirtualMachineListResult{}
	err := DeepCopy(&r, page.vmlrp.Response())
	if err != nil {
		page.err = fmt.Errorf("fail to get virtual machine list result, %s", err) //nolint:staticcheck
	}
	return r
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page VirtualMachineListResultPageClient) Values() []azcompute.VirtualMachine {
	r := []azcompute.VirtualMachine{}
	err := DeepCopy(&r, page.vmlrp.Values())
	if err != nil {
		page.err = fmt.Errorf("fail to get virtual machine list, %s", err) //nolint:staticcheck
	}
	return r
}

// VirtualMachineScaleSetListResultPageClient Virtual Machine Scale Set List Result Page Client
type VirtualMachineScaleSetListResultPageClient struct {
	vmsslrp compute.VirtualMachineScaleSetListResultPage
	err     error
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *VirtualMachineScaleSetListResultPageClient) NextWithContext(ctx context.Context) (err error) {
	return page.vmsslrp.NextWithContext(ctx)
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *VirtualMachineScaleSetListResultPageClient) Next() error {
	return page.vmsslrp.Next()
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page VirtualMachineScaleSetListResultPageClient) NotDone() bool {
	return page.vmsslrp.NotDone()
}

// Response returns the raw server response from the last page request.
func (page VirtualMachineScaleSetListResultPageClient) Response() azcompute.VirtualMachineScaleSetListResult {
	r := azcompute.VirtualMachineScaleSetListResult{}
	err := DeepCopy(&r, page.vmsslrp.Response())
	if err != nil {
		page.err = fmt.Errorf("fail to get virtual machine scale set list result, %s", err) //nolint:staticcheck
	}
	return r
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page VirtualMachineScaleSetListResultPageClient) Values() []azcompute.VirtualMachineScaleSet {
	r := []azcompute.VirtualMachineScaleSet{}
	err := DeepCopy(&r, page.vmsslrp.Values())
	if err != nil {
		page.err = fmt.Errorf("fail to get virtual machine scale set list, %s", err) //nolint:staticcheck
	}
	return r
}

// VirtualMachineScaleSetVMListResultPageClient Virtual Machine Scale Set VM List Result Page Client
type VirtualMachineScaleSetVMListResultPageClient struct {
	vmssvlrp compute.VirtualMachineScaleSetVMListResultPage
	err      error
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *VirtualMachineScaleSetVMListResultPageClient) NextWithContext(ctx context.Context) (err error) {
	return page.vmssvlrp.NextWithContext(ctx)
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *VirtualMachineScaleSetVMListResultPageClient) Next() error {
	return page.vmssvlrp.Next()
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page VirtualMachineScaleSetVMListResultPageClient) NotDone() bool {
	return page.vmssvlrp.NotDone()
}

// Response returns the raw server response from the last page request.
func (page VirtualMachineScaleSetVMListResultPageClient) Response() azcompute.VirtualMachineScaleSetVMListResult {
	r := azcompute.VirtualMachineScaleSetVMListResult{}
	err := DeepCopy(&r, page.vmssvlrp.Response())
	if err != nil {
		page.err = fmt.Errorf("fail to get virtual machine scale set VM list result, %s", err) //nolint:staticcheck
	}
	return r
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page VirtualMachineScaleSetVMListResultPageClient) Values() []azcompute.VirtualMachineScaleSetVM {
	r := []azcompute.VirtualMachineScaleSetVM{}
	err := DeepCopy(&r, page.vmssvlrp.Values())
	if err != nil {
		page.err = fmt.Errorf("fail to get virtual machine scale set VM list, %s", err) //nolint:staticcheck
	}
	return r
}

// DiskListPageClient contains a page of Disk values.
type DiskListPageClient struct {
	dlp compute.DiskListPage
	err error
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *DiskListPageClient) NextWithContext(ctx context.Context) (err error) {
	return page.dlp.NextWithContext(ctx)
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *DiskListPageClient) Next() error {
	return page.dlp.Next()
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page DiskListPageClient) NotDone() bool {
	return page.dlp.NotDone()
}

// Response returns the raw server response from the last page request.
func (page DiskListPageClient) Response() azcompute.DiskList {
	l := azcompute.DiskList{}
	err := DeepCopy(&l, page.dlp.Response())
	if err != nil {
		page.err = fmt.Errorf("fail to get disk list result, %s", err) //nolint:staticcheck
	}
	return l
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page DiskListPageClient) Values() []azcompute.Disk {
	l := []azcompute.Disk{}
	err := DeepCopy(&l, page.dlp.Values())
	if err != nil {
		page.err = fmt.Errorf("fail to get disk list, %s", err) //nolint:staticcheck
	}
	return l
}
