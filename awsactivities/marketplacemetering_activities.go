
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/marketplacemetering"
	"github.com/aws/aws-sdk-go/service/marketplacemetering/marketplacemeteringiface"
)

type MarketplaceMeteringActivities struct {
	client marketplacemeteringiface.MarketplaceMeteringAPI
}

func NewMarketplaceMeteringActivities(session *session.Session, config... *aws.Config) *MarketplaceMeteringActivities {
    client := marketplacemetering.New(session, config...)
    return &MarketplaceMeteringActivities{client: client}
}

func (a *MarketplaceMeteringActivities) BatchMeterUsage(input *marketplacemetering.BatchMeterUsageInput) (*marketplacemetering.BatchMeterUsageOutput, error) {
    return a.client.BatchMeterUsage(input)
}

func (a *MarketplaceMeteringActivities) MeterUsage(input *marketplacemetering.MeterUsageInput) (*marketplacemetering.MeterUsageOutput, error) {
    return a.client.MeterUsage(input)
}

func (a *MarketplaceMeteringActivities) RegisterUsage(input *marketplacemetering.RegisterUsageInput) (*marketplacemetering.RegisterUsageOutput, error) {
    return a.client.RegisterUsage(input)
}

func (a *MarketplaceMeteringActivities) ResolveCustomer(input *marketplacemetering.ResolveCustomerInput) (*marketplacemetering.ResolveCustomerOutput, error) {
    return a.client.ResolveCustomer(input)
}
