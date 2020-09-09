
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/marketplaceentitlementservice"
	"github.com/aws/aws-sdk-go/service/marketplaceentitlementservice/marketplaceentitlementserviceiface"
)

type MarketplaceEntitlementServiceActivities struct {
    client marketplaceentitlementserviceiface.MarketplaceEntitlementServiceAPI
}

func NewMarketplaceEntitlementServiceActivities(session *session.Session, config ...*aws.Config) *MarketplaceEntitlementServiceActivities {
    client := marketplaceentitlementservice.New(session, config...)
    return &MarketplaceEntitlementServiceActivities{client: client}
}

func (a *MarketplaceEntitlementServiceActivities) GetEntitlements(input *marketplaceentitlementservice.GetEntitlementsInput) (*marketplaceentitlementservice.GetEntitlementsOutput, error) {
    return a.client.GetEntitlements(input)
}
