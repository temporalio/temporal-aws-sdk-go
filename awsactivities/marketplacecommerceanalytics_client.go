package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics"
	"go.temporal.io/sdk/workflow"
)

type MarketplaceCommerceAnalyticsClient interface {
    GenerateDataSet(ctx workflow.Context, input *marketplacecommerceanalytics.GenerateDataSetInput) (*marketplacecommerceanalytics.GenerateDataSetOutput, error)
    GenerateDataSetAsync(ctx workflow.Context, input *marketplacecommerceanalytics.GenerateDataSetInput) *MarketplacecommerceanalyticsGenerateDataSetResult

    StartSupportDataExport(ctx workflow.Context, input *marketplacecommerceanalytics.StartSupportDataExportInput) (*marketplacecommerceanalytics.StartSupportDataExportOutput, error)
    StartSupportDataExportAsync(ctx workflow.Context, input *marketplacecommerceanalytics.StartSupportDataExportInput) *MarketplacecommerceanalyticsStartSupportDataExportResult
}
type MarketplacecommerceanalyticsGenerateDataSetResult struct {
	Result workflow.Future
}

func (r *MarketplacecommerceanalyticsGenerateDataSetResult) Get(ctx workflow.Context) (*marketplacecommerceanalytics.GenerateDataSetOutput, error) {
    var output marketplacecommerceanalytics.GenerateDataSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MarketplacecommerceanalyticsStartSupportDataExportResult struct {
	Result workflow.Future
}

func (r *MarketplacecommerceanalyticsStartSupportDataExportResult) Get(ctx workflow.Context) (*marketplacecommerceanalytics.StartSupportDataExportOutput, error) {
    var output marketplacecommerceanalytics.StartSupportDataExportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type MarketplaceCommerceAnalyticsStub struct {
    activities MarketplaceCommerceAnalyticsClient
}

func NewMarketplaceCommerceAnalyticsStub() MarketplaceCommerceAnalyticsClient {
    return &MarketplaceCommerceAnalyticsStub{}
}

func (a *MarketplaceCommerceAnalyticsStub) GenerateDataSet(ctx workflow.Context, input *marketplacecommerceanalytics.GenerateDataSetInput) (*marketplacecommerceanalytics.GenerateDataSetOutput, error) {
    var output marketplacecommerceanalytics.GenerateDataSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GenerateDataSet, input).Get(ctx, &output)
    return &output, err
}

func (a *MarketplaceCommerceAnalyticsStub) StartSupportDataExport(ctx workflow.Context, input *marketplacecommerceanalytics.StartSupportDataExportInput) (*marketplacecommerceanalytics.StartSupportDataExportOutput, error) {
    var output marketplacecommerceanalytics.StartSupportDataExportOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartSupportDataExport, input).Get(ctx, &output)
    return &output, err
}
