package awsclients

import (
	"github.com/aws/aws-sdk-go/service/marketplaceentitlementservice"
	"go.temporal.io/sdk/workflow"
)

type MarketplaceEntitlementServiceClient interface {
       GetEntitlements(ctx workflow.Context, input *marketplaceentitlementservice.GetEntitlementsInput) (*marketplaceentitlementservice.GetEntitlementsOutput, error)
       GetEntitlementsAsync(ctx workflow.Context, input *marketplaceentitlementservice.GetEntitlementsInput) *MarketplaceentitlementserviceGetEntitlementsResult
}

type MarketplaceEntitlementServiceStub struct {
}

func NewMarketplaceEntitlementServiceStub() MarketplaceEntitlementServiceClient {
    return &MarketplaceEntitlementServiceStub{}
}

type MarketplaceentitlementserviceGetEntitlementsResult struct {
	Result workflow.Future
}

func (r *MarketplaceentitlementserviceGetEntitlementsResult) Get(ctx workflow.Context) (*marketplaceentitlementservice.GetEntitlementsOutput, error) {
    var output marketplaceentitlementservice.GetEntitlementsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *MarketplaceEntitlementServiceStub) GetEntitlements(ctx workflow.Context, input *marketplaceentitlementservice.GetEntitlementsInput) (*marketplaceentitlementservice.GetEntitlementsOutput, error) {
    var output marketplaceentitlementservice.GetEntitlementsOutput
    err := workflow.ExecuteActivity(ctx, "MarketplaceEntitlementService.GetEntitlements", input).Get(ctx, &output)
    return &output, err
}

func (a *MarketplaceEntitlementServiceStub) GetEntitlementsAsync(ctx workflow.Context, input *marketplaceentitlementservice.GetEntitlementsInput) *MarketplaceentitlementserviceGetEntitlementsResult {
    future := workflow.ExecuteActivity(ctx, "MarketplaceEntitlementService.GetEntitlements", input)
    return &MarketplaceentitlementserviceGetEntitlementsResult{Result: future}
}
