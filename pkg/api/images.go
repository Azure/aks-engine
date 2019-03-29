// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"path"
	"strings"

	"github.com/docker/distribution/reference"
	"github.com/pkg/errors"
)

var (
	// ErrComponentNotFound signifies that the provided component name was not found
	// amonst the list of available components.
	ErrComponentNotFound = errors.New("component not found")
)

// GetKubernetesComponentImage gets the image repo for the kubernetes component passed in.
// This takes into account `KubernetesImageBase` and `ImageRepoOverrides`.
//
// This does not account for things like `CustomHyperkubeImage`, it is expected
// that the caller knows about these fields and handles accordingly.
//
// KubernetesConfig must not be nil, it is a major bug if it is nil, so nils are not checked
func GetKubernetesComponentImage(name string, components map[string]string, c *KubernetesConfig, isAzureStack bool, cloudSpecConfig AzureEnvironmentSpecConfig) (string, error) {
	base := strings.TrimSuffix(c.KubernetesImageBase, "/")
	if isAzureStack && name != "hyperkube" {
		base = strings.TrimSuffix(cloudSpecConfig.KubernetesSpecConfig.KubernetesImageBase, "/")
	}

	image := components[name]
	if image == "" {
		return "", errors.Wrap(ErrComponentNotFound, name)
	}
	ref, err := reference.ParseAnyReference(path.Join(base, image))
	if err != nil {
		return "", errors.Wrapf(err, "error parsing image reference for component: %s", name)
	}

	named, ok := ref.(reference.Named)
	if !ok {
		return "", errors.Errorf("invalid reference format: %s", ref.String())
	}

	// This sets a default tag ("latest") if no tag is provided.
	// This really shouldn't happen, but it makes certain that we always have a
	// tagged reference.
	named = reference.TagNameOnly(named)

	domain := reference.Domain(named)
	repo := reference.Path(named)
	tag := named.(reference.Tagged).Tag()

	noTag := reference.TrimNamed(named)
	if o, ok := c.ImageRepoOverrides[noTag.String()]; ok {
		if o.Registry != "" {
			domain = o.Registry
		}
		if o.Repo != "" {
			repo = o.Repo
		}

		if t, ok := o.Tags[tag]; ok {
			tag = t
		}

		r, err := reference.ParseAnyReference(path.Join(domain, repo))
		if err != nil {
			return "", errors.Wrapf(err, "error parsing reference with overrides for component: %s", name)
		}
		if _, ok := r.(reference.Tagged); ok {
			// repo name seems to contain a tag reference, which is unexpected
			return "", errors.Wrapf(reference.ErrReferenceInvalidFormat, "invalid repo override for %s: %s", name, r)
		}
		named, err = reference.WithTag(r.(reference.Named), tag)
		if err != nil {
			return "", errors.Wrapf(err, "error making tagged reference for component: %s", name)
		}
	}

	return named.String(), nil
}
