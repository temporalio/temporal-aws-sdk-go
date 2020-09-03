package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/marketplaceentitlementservice"
	"go.temporal.io/sdk/workflow"
)

type MarketplaceEntitlementServiceClient interface {
    GetEntitlements(ctx workflow.Context, input *marketplaceentitlementservice.GetEntitlementsInput) (*marketplaceentitlementservice.GetEntitlementsOutput, error)
    GetEntitlementsAsync(ctx workflow.Context, input *marketplaceentitlementservice.GetEntitlementsInput) *MarketplaceentitlementserviceGetEntitlementsResult
}
type MarketplaceentitlementserviceGetEntitlementsResult struct {
	Result workflow.Future
}

func (r *MarketplaceentitlementserviceGetEntitlementsResult) Get(ctx workflow.Context) (*marketplaceentitlementservice.GetEntitlementsOutput, error) {
    var output marketplaceentitlementservice.GetEntitlementsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type MarketplaceEntitlementServiceStub struct {
    activities MarketplaceEntitlementServiceClient
}

func NewMarketplaceEntitlementServiceStub() MarketplaceEntitlementServiceClient {
    return &MarketplaceEntitlementServiceStub{}
}

func (a *MarketplaceEntitlementServiceStub) GetEntitlements(ctx workflow.Context, input *marketplaceentitlementservice.GetEntitlementsInput) (*marketplaceentitlementservice.GetEntitlementsOutput, error) {
    var output marketplaceentitlementservice.GetEntitlementsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetEntitlements, input).Get(ctx, &output)
    return &output, err
}
