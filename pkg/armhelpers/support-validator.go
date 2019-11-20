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

	var missingVMImages map[string]requiredImage

	imageInformaition := getDistroVMImageInformation(p.MasterProfile.Distro)

	vmImage, err := client.GetVirtualMachineImage(ctx, location, imageInformaition.ImagePublisher, imageInformaition.ImageOffer, imageInformaition.ImageSku, imageInformaition.ImageVersion)

	if err != nil {
		if p.MasterProfile.Distro == "" {
			missingVMImages[string(api.Ubuntu)] = requiredImage{
				image:     imageInformaition,
				errorData: err,
			}

		} else {
			missingVMImages[string(p.MasterProfile.Distro)] = requiredImage{
				image:     imageInformaition,
				errorData: err,
			}
		}
	}

	log.Infof("Image %s", *vmImage.Name)

	for _, profile := range p.AgentPoolProfiles {

		if profile.OSType == api.Windows {
			imageInformaition := api.WindowsServer2019OSImageConfig

			vmImage, err := client.GetVirtualMachineImage(ctx, location, imageInformaition.ImagePublisher, imageInformaition.ImageOffer, imageInformaition.ImageSku, imageInformaition.ImageVersion)

			if err != nil {
				missingVMImages[string(api.Windows)] = requiredImage{
					image:     imageInformaition,
					errorData: err,
				}
			}

			log.Infof("Found Image %s", *vmImage.Name)
		} else {
			imageInformaition := getDistroVMImageInformation(profile.Distro)

			vmImage, err := client.GetVirtualMachineImage(ctx, location, imageInformaition.ImagePublisher, imageInformaition.ImageOffer, imageInformaition.ImageSku, imageInformaition.ImageVersion)

			if err != nil {
				if profile.Distro == "" {
					missingVMImages[string(api.Ubuntu)] = requiredImage{
						image:     imageInformaition,
						errorData: err,
					}
				} else {
					missingVMImages[string(profile.Distro)] = requiredImage{
						image:     imageInformaition,
						errorData: err,
					}
				}
			}

			log.Infof("Found Image %s", *vmImage.Name)
		}
	}

	if len(missingVMImages) > 0 {
		for _, value := range missingVMImages {
			imageDetails := value.image
			log.Infof("Offer: %s, Sku: %s, Version: %s", imageDetails.ImageOffer, imageDetails.ImageSku, imageDetails.ImageVersion)
			log.Errorf("Error: ", value.errorData)
		}
		return errors.New("Some VM images are missing from your environment. See above for details")
	}

	return nil

}

func getDistroVMImageInformation(d api.Distro) api.AzureOSImageConfig {
	switch d {
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
