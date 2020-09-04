
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/pricing"
	"github.com/aws/aws-sdk-go/service/pricing/pricingiface"
)

type PricingActivities struct {
	client pricingiface.PricingAPI
}

func NewPricingActivities(client pricingiface.PricingAPI) *PricingActivities {
    return &PricingActivities{client: client}
}

func (a *PricingActivities) DescribeServices(input *pricing.DescribeServicesInput) (*pricing.DescribeServicesOutput, error) {
    return a.client.DescribeServices(input)
}

func (a *PricingActivities) GetAttributeValues(input *pricing.GetAttributeValuesInput) (*pricing.GetAttributeValuesOutput, error) {
    return a.client.GetAttributeValues(input)
}

func (a *PricingActivities) GetProducts(input *pricing.GetProductsInput) (*pricing.GetProductsOutput, error) {
    return a.client.GetProducts(input)
}
