
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics"
	"github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics/marketplacecommerceanalyticsiface"
)

type MarketplaceCommerceAnalyticsActivities struct {
	client marketplacecommerceanalyticsiface.MarketplaceCommerceAnalyticsAPI
}

func NewMarketplaceCommerceAnalyticsActivities(client marketplacecommerceanalyticsiface.MarketplaceCommerceAnalyticsAPI) *MarketplaceCommerceAnalyticsActivities {
    return &MarketplaceCommerceAnalyticsActivities{client: client}
}

func (a *MarketplaceCommerceAnalyticsActivities) GenerateDataSet(input *marketplacecommerceanalytics.GenerateDataSetInput) (*marketplacecommerceanalytics.GenerateDataSetOutput, error) {
    return a.client.GenerateDataSet(input)
}

func (a *MarketplaceCommerceAnalyticsActivities) StartSupportDataExport(input *marketplacecommerceanalytics.StartSupportDataExportInput) (*marketplacecommerceanalytics.StartSupportDataExportOutput, error) {
    return a.client.StartSupportDataExport(input)
}
