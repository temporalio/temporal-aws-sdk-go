package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics"
	"github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics/marketplacecommerceanalyticsiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type MarketplaceCommerceAnalyticsActivities struct {
	client marketplacecommerceanalyticsiface.MarketplaceCommerceAnalyticsAPI
}

func NewMarketplaceCommerceAnalyticsActivities(session *session.Session, config ...*aws.Config) *MarketplaceCommerceAnalyticsActivities {
	client := marketplacecommerceanalytics.New(session, config...)
	return &MarketplaceCommerceAnalyticsActivities{client: client}
}

func (a *MarketplaceCommerceAnalyticsActivities) GenerateDataSet(ctx context.Context, input *marketplacecommerceanalytics.GenerateDataSetInput) (*marketplacecommerceanalytics.GenerateDataSetOutput, error) {
	return a.client.GenerateDataSetWithContext(ctx, input)
}

func (a *MarketplaceCommerceAnalyticsActivities) StartSupportDataExport(ctx context.Context, input *marketplacecommerceanalytics.StartSupportDataExportInput) (*marketplacecommerceanalytics.StartSupportDataExportOutput, error) {
	return a.client.StartSupportDataExportWithContext(ctx, input)
}
