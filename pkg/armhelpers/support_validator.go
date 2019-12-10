// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type validationResult struct {
	image     api.AzureOSImageConfig
	errorData error
}

// ValidateRequiredImages checks that the OS images required by both
// master and agent pools are available on the target cloud
func ValidateRequiredImages(ctx context.Context, location string, p *api.Properties, client AKSEngineClient) error {
	if fetcher, ok := client.(VMImageFetcher); ok {
		missingImages := make(map[api.Distro]validationResult)
		for distro, i := range requiredImages(p) {
			if i.ImageVersion == "latest" {
				list, err := fetcher.ListVirtualMachineImages(ctx, location, i.ImagePublisher, i.ImageOffer, i.ImageSku)
				if err != nil || len(*list.Value) == 0 {
					missingImages[distro] = validationResult{
						image:     i,
						errorData: err,
					}
				}
			} else {
				if _, err := fetcher.GetVirtualMachineImage(ctx, location, i.ImagePublisher, i.ImageOffer, i.ImageSku, i.ImageVersion); err != nil {
					missingImages[distro] = validationResult{
						image:     i,
						errorData: err,
					}
				}
			}
		}
		if len(missingImages) == 0 {
			return nil
		}
		return printErrorIfAny(missingImages)
	}
	return errors.New("parameter client is not a VMImageFetcher")
}

func requiredImages(p *api.Properties) map[api.Distro]api.AzureOSImageConfig {
	images := make(map[api.Distro]api.AzureOSImageConfig)
	images[p.MasterProfile.Distro] = toImageConfig(p.MasterProfile.Distro)
	for _, app := range p.AgentPoolProfiles {
		if app.OSType == api.Windows {
			images[app.Distro] = api.WindowsServer2019OSImageConfig
		} else {
			images[app.Distro] = toImageConfig(app.Distro)
		}
	}
	return images
}

func printErrorIfAny(missingImages map[api.Distro]validationResult) error {
	for _, value := range missingImages {
		i := value.image
		log.Errorf("error: %+v", value.errorData)
		log.Errorf("Image Publisher: %s, Offer: %s, SKU: %s, Version: %s", i.ImagePublisher, i.ImageOffer, i.ImageSku, i.ImageVersion)
	}
	return errors.New("some VM images are missing on the target cloud")
}

func toImageConfig(distro api.Distro) api.AzureOSImageConfig {
	if distro == "" {
		return api.Ubuntu1604OSImageConfig
	}
	switch distro {
	case api.Ubuntu:
		return api.Ubuntu1604OSImageConfig
	case api.Ubuntu1804:
		return api.Ubuntu1804OSImageConfig
	case api.RHEL:
		return api.RHELOSImageConfig
	case api.CoreOS:
		return api.CoreOSImageConfig
	case api.AKSUbuntu1604:
		return api.AKSUbuntu1604OSImageConfig
	case api.AKSUbuntu1804:
		return api.AKSUbuntu1804OSImageConfig
	case api.ACC1604:
		return api.ACC1604OSImageConfig
	default:
		return api.Ubuntu1604OSImageConfig
	}
}
