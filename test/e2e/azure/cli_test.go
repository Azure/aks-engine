//+build test
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package azure

import (
	"fmt"
	"testing"
	"time"
)

func TestIsClusterExpired(t *testing.T) {
	cases := []struct {
		rg             ResourceGroup
		a              Account
		duration       string
		expectedResult bool
	}{
		{
			rg: ResourceGroup{
				Name:     "testRG",
				Location: "test",
				Tags: map[string]string{
					"now": "799786800",
				},
			},
			a: Account{
				User:           new(User),
				TenantID:       "1234",
				SubscriptionID: "1234",
				ResourceGroup:  ResourceGroup{},
				Deployment:     Deployment{},
			},
			duration:       "1h",
			expectedResult: true,
		},
		{
			rg: ResourceGroup{
				Name:     "testRG",
				Location: "test",
				Tags: map[string]string{
					"now": fmt.Sprintf("%v", time.Now().Unix()),
				},
			},
			a: Account{
				User:           new(User),
				TenantID:       "1234",
				SubscriptionID: "1234",
				ResourceGroup:  ResourceGroup{},
				Deployment:     Deployment{},
			},
			duration:       "300h",
			expectedResult: false,
		},
		{
			rg: ResourceGroup{
				Name:     "thisRGDoesNotExist",
				Location: "test",
				Tags:     map[string]string{},
			},
			a: Account{
				User:           new(User),
				TenantID:       "1234",
				SubscriptionID: "1234",
				ResourceGroup:  ResourceGroup{},
				Deployment:     Deployment{},
			},
			duration:       "1s",
			expectedResult: true,
		},
	}

	for _, c := range cases {
		c.a.ResourceGroup = c.rg
		d, err := time.ParseDuration(c.duration)
		if err != nil {
			t.Fatalf("unexpected error parsing duration: %s", err)
		}
		result := c.a.IsClusterExpired(d)
		if c.expectedResult != result {
			t.Fatalf("IsClusterExpired returned unexpected result: expected %t but got %t", c.expectedResult, result)
		}
	}
}
