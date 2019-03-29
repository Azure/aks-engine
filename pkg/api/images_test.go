// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"testing"

	"github.com/docker/distribution/reference"
	"github.com/pkg/errors"
)

func TestGetKubernetesComponentImage(t *testing.T) {
	type testCase struct {
		desc        string
		expectErr   error
		expectValue string

		comp   string
		config *KubernetesConfig
	}

	components := map[string]string{
		"foo":     "foo:v1.0.0",
		"invalid": "invalid:",
		"noimage": "",
		"notag":   "notag",
	}

	cases := []testCase{
		{
			desc:      "missing component",
			expectErr: ErrComponentNotFound,
			comp:      "not exist",
			config:    &KubernetesConfig{},
		},
		{
			desc:      "no compnent image",
			expectErr: ErrComponentNotFound,
			comp:      "noimage",
			config:    &KubernetesConfig{},
		},

		{
			desc:      "invalid base image reference format",
			expectErr: reference.ErrReferenceInvalidFormat,
			comp:      "foo",
			config:    &KubernetesConfig{KubernetesImageBase: "garbage:entry"},
		},
		{
			desc:      "invalid component image name",
			comp:      "invalid",
			config:    &KubernetesConfig{},
			expectErr: reference.ErrReferenceInvalidFormat,
		},
		{
			desc:        "no base image reference and no overrides",
			comp:        "foo",
			config:      &KubernetesConfig{},
			expectValue: "docker.io/library/foo:v1.0.0",
		},
		{
			desc:        "with base image reference and no overrides",
			comp:        "foo",
			config:      &KubernetesConfig{KubernetesImageBase: "foo.com"},
			expectValue: "foo.com/foo:v1.0.0",
		},
		{
			desc:        "with empty component tag",
			comp:        "notag",
			config:      &KubernetesConfig{KubernetesImageBase: "foo.com"},
			expectValue: "foo.com/notag:latest",
		},
		{
			desc: "with base image reference and empty override",
			comp: "foo",
			config: &KubernetesConfig{KubernetesImageBase: "foo.com", ImageRepoOverrides: map[string]ImageRepoOverride{
				"foo.com/foo": {},
			}},
			expectValue: "foo.com/foo:v1.0.0",
		},
		{
			desc: "with base image reference and override registry name only",
			comp: "foo",
			config: &KubernetesConfig{KubernetesImageBase: "foo.com", ImageRepoOverrides: map[string]ImageRepoOverride{
				"foo.com/foo": {
					Registry: "bar.com",
				},
			}},
			expectValue: "bar.com/foo:v1.0.0",
		},
		{
			desc: "with base image reference and override image repo only",
			comp: "foo",
			config: &KubernetesConfig{KubernetesImageBase: "foo.com", ImageRepoOverrides: map[string]ImageRepoOverride{
				"foo.com/foo": {
					Repo: "bar/baz",
				},
			}},
			expectValue: "foo.com/bar/baz:v1.0.0",
		},
		{
			desc: "with base image reference and override image tag only",
			comp: "foo",
			config: &KubernetesConfig{KubernetesImageBase: "foo.com", ImageRepoOverrides: map[string]ImageRepoOverride{
				"foo.com/foo": {
					Tags: map[string]string{"v1.0.0": "whatever", "v2.0.0": "another"},
				},
			}},
			expectValue: "foo.com/foo:whatever",
		},
		{
			desc: "with base image reference and overrides for all",
			comp: "foo",
			config: &KubernetesConfig{KubernetesImageBase: "foo.com", ImageRepoOverrides: map[string]ImageRepoOverride{
				"foo.com/foo": {
					Registry: "bar.com",
					Repo:     "bar/baz",
					Tags:     map[string]string{"v1.0.0": "whatever", "v2.0.0": "another"},
				},
			}},
			expectValue: "bar.com/bar/baz:whatever",
		},
		{
			desc: "with base image reference and overrides with invalid repo name",
			comp: "foo",
			config: &KubernetesConfig{KubernetesImageBase: "foo.com", ImageRepoOverrides: map[string]ImageRepoOverride{
				"foo.com/foo": {
					Registry: "bar.com",
					Repo:     "bar:baz",
					Tags:     map[string]string{"v1.0.0": "whatever", "v2.0.0": "another"},
				},
			}},
			expectErr: reference.ErrReferenceInvalidFormat,
		},
		{
			desc: "with base image reference and overrides with invalid tag",
			comp: "foo",
			config: &KubernetesConfig{KubernetesImageBase: "foo.com", ImageRepoOverrides: map[string]ImageRepoOverride{
				"foo.com/foo": {
					Registry: "bar.com",
					Repo:     "bar/baz",
					Tags:     map[string]string{"v1.0.0": "bro:ken", "v2.0.0": "another"},
				},
			}},
			expectErr: reference.ErrTagInvalidFormat,
		},
		{
			desc: "with base image reference and overrides with unmatched tag override",
			comp: "foo",
			config: &KubernetesConfig{KubernetesImageBase: "foo.com", ImageRepoOverrides: map[string]ImageRepoOverride{
				"foo.com/foo": {
					Registry: "bar.com",
					Repo:     "bar/baz",
					Tags:     map[string]string{"unmatched1": "v1.0.0", "unmatched2": "v1.0.0"},
				},
			}},
			expectValue: "bar.com/bar/baz:v1.0.0",
		},
		{
			desc: "with base image reference and overrides with empty tag override",
			comp: "foo",
			config: &KubernetesConfig{KubernetesImageBase: "foo.com", ImageRepoOverrides: map[string]ImageRepoOverride{
				"foo.com/foo": {
					Registry: "bar.com",
					Repo:     "bar/baz",
					Tags:     map[string]string{"v1.0.0": ""},
				},
			}},
			expectErr: reference.ErrTagInvalidFormat,
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			cloudSpecConfig := AzureEnvironmentSpecConfig{}
			val, err := GetKubernetesComponentImage(c.comp, components, c.config, false, cloudSpecConfig)
			if e := errors.Cause(err); e != c.expectErr {
				t.Fatalf("expected error %q, got %q, result value: %s", c.expectErr, e, val)
			}

			if val != c.expectValue {
				t.Fatalf("expected value %q, got %q", c.expectValue, val)
			}
		})
	}
}
