package azurestack

import (
	"context"

	"github.com/Azure/aks-engine/pkg/armhelpers"
)

//GetImagesList returns the list of images available in the current environment
func (az *AzureClient) GetImagesList(ctx context.Context) (result armhelpers.ImageListResultPage, err error) {
	page, err := az.imageClient.List(ctx)

	return &ImageListResultPageClient{
		dlp: page,
		err: err,
	}, err
}
