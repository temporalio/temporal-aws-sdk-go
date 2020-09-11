package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/pricing"
	"github.com/aws/aws-sdk-go/service/pricing/pricingiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type PricingActivities struct {
	client pricingiface.PricingAPI
}

func NewPricingActivities(session *session.Session, config ...*aws.Config) *PricingActivities {
	client := pricing.New(session, config...)
	return &PricingActivities{client: client}
}

func (a *PricingActivities) DescribeServices(ctx context.Context, input *pricing.DescribeServicesInput) (*pricing.DescribeServicesOutput, error) {
	return a.client.DescribeServicesWithContext(ctx, input)
}

func (a *PricingActivities) GetAttributeValues(ctx context.Context, input *pricing.GetAttributeValuesInput) (*pricing.GetAttributeValuesOutput, error) {
	return a.client.GetAttributeValuesWithContext(ctx, input)
}

func (a *PricingActivities) GetProducts(ctx context.Context, input *pricing.GetProductsInput) (*pricing.GetProductsOutput, error) {
	return a.client.GetProductsWithContext(ctx, input)
}
