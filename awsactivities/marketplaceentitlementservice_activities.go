
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/marketplaceentitlementservice"
	"github.com/aws/aws-sdk-go/service/marketplaceentitlementservice/marketplaceentitlementserviceiface"
)

type MarketplaceEntitlementServiceActivities struct {
	client marketplaceentitlementserviceiface.MarketplaceEntitlementServiceAPI
}

func NewMarketplaceEntitlementServiceActivities(client marketplaceentitlementserviceiface.MarketplaceEntitlementServiceAPI) *MarketplaceEntitlementServiceActivities {
    return &MarketplaceEntitlementServiceActivities{client: client}
}

func (a *MarketplaceEntitlementServiceActivities) GetEntitlements(input *marketplaceentitlementservice.GetEntitlementsInput) (*marketplaceentitlementservice.GetEntitlementsOutput, error) {
    return a.client.GetEntitlements(input)
}
