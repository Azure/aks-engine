// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package armhelpers

import (
	"context"

	"github.com/Azure/aks-engine/pkg/api"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type requiredImage struct {
	image     api.AzureOSImageConfig
	errorData error
}

//ValidateRequiredImages checks both master ang agent profiles and ensures these images are available in the environment
func ValidateRequiredImages(ctx context.Context, location string, p *api.Properties, client AKSEngineClient) error {
	missingVMImages := make(map[string]requiredImage)
	imageInformation := getDistroVMImageInformation(string(p.MasterProfile.Distro))

	if imageFetcher, ok := client.(VMImageFetcher); ok {
		_, err := imageFetcher.GetVirtualMachineImage(ctx, location, imageInformation.ImagePublisher, imageInformation.ImageOffer, imageInformation.ImageSku, imageInformation.ImageVersion)

		if err != nil {
			var distro = string(api.Ubuntu)
			if p.MasterProfile.Distro != "" {
				distro = string(p.MasterProfile.Distro)
			}

			missingVMImages[distro] = requiredImage{
				image:     imageInformation,
				errorData: err,
			}
		}

		for _, profile := range p.AgentPoolProfiles {
			var imgInfo api.AzureOSImageConfig
			var distro string
			if profile.OSType == api.Windows {
				imgInfo = api.WindowsServer2019OSImageConfig
				distro = string(api.Windows)
			} else {
				distro = string(api.Ubuntu)
				if profile.Distro != "" {
					distro = string(profile.Distro)
				}

				imgInfo = getDistroVMImageInformation(distro)
			}

			if _, err := imageFetcher.GetVirtualMachineImage(ctx, location, imgInfo.ImagePublisher, imgInfo.ImageOffer, imgInfo.ImageSku, imgInfo.ImageVersion); err != nil {
				missingVMImages[distro] = requiredImage{
					image:     imgInfo,
					errorData: err,
				}
			}
		}

		if len(missingVMImages) > 0 {
			for _, value := range missingVMImages {
				imageDetails := value.image
				log.Infof("Offer: %s, Sku: %s, Version: %s", imageDetails.ImageOffer, imageDetails.ImageSku, imageDetails.ImageVersion)
				log.Errorf("Error: %+v", value.errorData)
			}
			return errors.New("Some VM images are missing from your environment. See above for details")
		}
	}

	return nil

}

func getDistroVMImageInformation(d string) api.AzureOSImageConfig {
	distro := api.Distro(d)
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
