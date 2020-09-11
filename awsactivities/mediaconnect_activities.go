package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mediaconnect"
	"github.com/aws/aws-sdk-go/service/mediaconnect/mediaconnectiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type MediaConnectActivities struct {
	client mediaconnectiface.MediaConnectAPI
}

func NewMediaConnectActivities(session *session.Session, config ...*aws.Config) *MediaConnectActivities {
	client := mediaconnect.New(session, config...)
	return &MediaConnectActivities{client: client}
}

func (a *MediaConnectActivities) AddFlowOutputs(ctx context.Context, input *mediaconnect.AddFlowOutputsInput) (*mediaconnect.AddFlowOutputsOutput, error) {
	return a.client.AddFlowOutputsWithContext(ctx, input)
}

func (a *MediaConnectActivities) AddFlowSources(ctx context.Context, input *mediaconnect.AddFlowSourcesInput) (*mediaconnect.AddFlowSourcesOutput, error) {
	return a.client.AddFlowSourcesWithContext(ctx, input)
}

func (a *MediaConnectActivities) AddFlowVpcInterfaces(ctx context.Context, input *mediaconnect.AddFlowVpcInterfacesInput) (*mediaconnect.AddFlowVpcInterfacesOutput, error) {
	return a.client.AddFlowVpcInterfacesWithContext(ctx, input)
}

func (a *MediaConnectActivities) CreateFlow(ctx context.Context, input *mediaconnect.CreateFlowInput) (*mediaconnect.CreateFlowOutput, error) {
	return a.client.CreateFlowWithContext(ctx, input)
}

func (a *MediaConnectActivities) DeleteFlow(ctx context.Context, input *mediaconnect.DeleteFlowInput) (*mediaconnect.DeleteFlowOutput, error) {
	return a.client.DeleteFlowWithContext(ctx, input)
}

func (a *MediaConnectActivities) DescribeFlow(ctx context.Context, input *mediaconnect.DescribeFlowInput) (*mediaconnect.DescribeFlowOutput, error) {
	return a.client.DescribeFlowWithContext(ctx, input)
}

func (a *MediaConnectActivities) GrantFlowEntitlements(ctx context.Context, input *mediaconnect.GrantFlowEntitlementsInput) (*mediaconnect.GrantFlowEntitlementsOutput, error) {
	return a.client.GrantFlowEntitlementsWithContext(ctx, input)
}

func (a *MediaConnectActivities) ListEntitlements(ctx context.Context, input *mediaconnect.ListEntitlementsInput) (*mediaconnect.ListEntitlementsOutput, error) {
	return a.client.ListEntitlementsWithContext(ctx, input)
}

func (a *MediaConnectActivities) ListFlows(ctx context.Context, input *mediaconnect.ListFlowsInput) (*mediaconnect.ListFlowsOutput, error) {
	return a.client.ListFlowsWithContext(ctx, input)
}

func (a *MediaConnectActivities) ListTagsForResource(ctx context.Context, input *mediaconnect.ListTagsForResourceInput) (*mediaconnect.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *MediaConnectActivities) RemoveFlowOutput(ctx context.Context, input *mediaconnect.RemoveFlowOutputInput) (*mediaconnect.RemoveFlowOutputOutput, error) {
	return a.client.RemoveFlowOutputWithContext(ctx, input)
}

func (a *MediaConnectActivities) RemoveFlowSource(ctx context.Context, input *mediaconnect.RemoveFlowSourceInput) (*mediaconnect.RemoveFlowSourceOutput, error) {
	return a.client.RemoveFlowSourceWithContext(ctx, input)
}

func (a *MediaConnectActivities) RemoveFlowVpcInterface(ctx context.Context, input *mediaconnect.RemoveFlowVpcInterfaceInput) (*mediaconnect.RemoveFlowVpcInterfaceOutput, error) {
	return a.client.RemoveFlowVpcInterfaceWithContext(ctx, input)
}

func (a *MediaConnectActivities) RevokeFlowEntitlement(ctx context.Context, input *mediaconnect.RevokeFlowEntitlementInput) (*mediaconnect.RevokeFlowEntitlementOutput, error) {
	return a.client.RevokeFlowEntitlementWithContext(ctx, input)
}

func (a *MediaConnectActivities) StartFlow(ctx context.Context, input *mediaconnect.StartFlowInput) (*mediaconnect.StartFlowOutput, error) {
	return a.client.StartFlowWithContext(ctx, input)
}

func (a *MediaConnectActivities) StopFlow(ctx context.Context, input *mediaconnect.StopFlowInput) (*mediaconnect.StopFlowOutput, error) {
	return a.client.StopFlowWithContext(ctx, input)
}

func (a *MediaConnectActivities) TagResource(ctx context.Context, input *mediaconnect.TagResourceInput) (*mediaconnect.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *MediaConnectActivities) UntagResource(ctx context.Context, input *mediaconnect.UntagResourceInput) (*mediaconnect.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *MediaConnectActivities) UpdateFlow(ctx context.Context, input *mediaconnect.UpdateFlowInput) (*mediaconnect.UpdateFlowOutput, error) {
	return a.client.UpdateFlowWithContext(ctx, input)
}

func (a *MediaConnectActivities) UpdateFlowEntitlement(ctx context.Context, input *mediaconnect.UpdateFlowEntitlementInput) (*mediaconnect.UpdateFlowEntitlementOutput, error) {
	return a.client.UpdateFlowEntitlementWithContext(ctx, input)
}

func (a *MediaConnectActivities) UpdateFlowOutput(ctx context.Context, input *mediaconnect.UpdateFlowOutputInput) (*mediaconnect.UpdateFlowOutputOutput, error) {
	return a.client.UpdateFlowOutputWithContext(ctx, input)
}

func (a *MediaConnectActivities) UpdateFlowSource(ctx context.Context, input *mediaconnect.UpdateFlowSourceInput) (*mediaconnect.UpdateFlowSourceOutput, error) {
	return a.client.UpdateFlowSourceWithContext(ctx, input)
}
