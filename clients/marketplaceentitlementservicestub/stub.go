// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package marketplaceentitlementservicestub

import (
	"github.com/aws/aws-sdk-go/service/marketplaceentitlementservice"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type GetEntitlementsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetEntitlementsFuture) Get(ctx workflow.Context) (*marketplaceentitlementservice.GetEntitlementsOutput, error) {
	var output marketplaceentitlementservice.GetEntitlementsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) GetEntitlements(ctx workflow.Context, input *marketplaceentitlementservice.GetEntitlementsInput) (*marketplaceentitlementservice.GetEntitlementsOutput, error) {
	var output marketplaceentitlementservice.GetEntitlementsOutput
	err := workflow.ExecuteActivity(ctx, "aws.marketplaceentitlementservice.GetEntitlements", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetEntitlementsAsync(ctx workflow.Context, input *marketplaceentitlementservice.GetEntitlementsInput) *GetEntitlementsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.marketplaceentitlementservice.GetEntitlements", input)
	return &GetEntitlementsFuture{Future: future}
}
