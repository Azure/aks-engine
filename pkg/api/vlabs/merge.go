// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package vlabs

import (
	"github.com/imdario/mergo"
)

// Merge existing containerService attribute into cs
func (cs *ContainerService) Merge(ecs *ContainerService) error {
	if ecs.Properties.ServicePrincipalProfile != nil {
		if cs.Properties.ServicePrincipalProfile == nil {
			cs.Properties.ServicePrincipalProfile = &ServicePrincipalProfile{}
		}
		if err := mergo.Merge(cs.Properties.ServicePrincipalProfile,
			*ecs.Properties.ServicePrincipalProfile); err != nil {
			return err
		}
	}
	if ecs.Properties.WindowsProfile != nil {
		if cs.Properties.WindowsProfile == nil {
			cs.Properties.WindowsProfile = &WindowsProfile{}
		}
		if err := mergo.Merge(cs.Properties.WindowsProfile,
			*ecs.Properties.WindowsProfile); err != nil {
			return err
		}
	}
	return nil
}
