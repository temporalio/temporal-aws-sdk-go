package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/marketplacemetering"
	"github.com/aws/aws-sdk-go/service/marketplacemetering/marketplacemeteringiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type MarketplaceMeteringActivities struct {
	client marketplacemeteringiface.MarketplaceMeteringAPI
}

func NewMarketplaceMeteringActivities(session *session.Session, config ...*aws.Config) *MarketplaceMeteringActivities {
	client := marketplacemetering.New(session, config...)
	return &MarketplaceMeteringActivities{client: client}
}

func (a *MarketplaceMeteringActivities) BatchMeterUsage(ctx context.Context, input *marketplacemetering.BatchMeterUsageInput) (*marketplacemetering.BatchMeterUsageOutput, error) {
	return a.client.BatchMeterUsageWithContext(ctx, input)
}

func (a *MarketplaceMeteringActivities) MeterUsage(ctx context.Context, input *marketplacemetering.MeterUsageInput) (*marketplacemetering.MeterUsageOutput, error) {
	return a.client.MeterUsageWithContext(ctx, input)
}

func (a *MarketplaceMeteringActivities) RegisterUsage(ctx context.Context, input *marketplacemetering.RegisterUsageInput) (*marketplacemetering.RegisterUsageOutput, error) {
	return a.client.RegisterUsageWithContext(ctx, input)
}

func (a *MarketplaceMeteringActivities) ResolveCustomer(ctx context.Context, input *marketplacemetering.ResolveCustomerInput) (*marketplacemetering.ResolveCustomerOutput, error) {
	return a.client.ResolveCustomerWithContext(ctx, input)
}
