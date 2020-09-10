package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/marketplaceentitlementservice"
	"github.com/aws/aws-sdk-go/service/marketplaceentitlementservice/marketplaceentitlementserviceiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type MarketplaceEntitlementServiceActivities struct {
	client marketplaceentitlementserviceiface.MarketplaceEntitlementServiceAPI
}

func NewMarketplaceEntitlementServiceActivities(session *session.Session, config ...*aws.Config) *MarketplaceEntitlementServiceActivities {
	client := marketplaceentitlementservice.New(session, config...)
	return &MarketplaceEntitlementServiceActivities{client: client}
}

func (a *MarketplaceEntitlementServiceActivities) GetEntitlements(ctx context.Context, input *marketplaceentitlementservice.GetEntitlementsInput) (*marketplaceentitlementservice.GetEntitlementsOutput, error) {
	return a.client.GetEntitlementsWithContext(ctx, input)
}
