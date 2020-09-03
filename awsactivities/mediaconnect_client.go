package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/mediaconnect"
	"go.temporal.io/sdk/workflow"
)

type MediaConnectClient interface {
    AddFlowOutputs(ctx workflow.Context, input *mediaconnect.AddFlowOutputsInput) (*mediaconnect.AddFlowOutputsOutput, error)
    AddFlowOutputsAsync(ctx workflow.Context, input *mediaconnect.AddFlowOutputsInput) *MediaconnectAddFlowOutputsResult

    AddFlowSources(ctx workflow.Context, input *mediaconnect.AddFlowSourcesInput) (*mediaconnect.AddFlowSourcesOutput, error)
    AddFlowSourcesAsync(ctx workflow.Context, input *mediaconnect.AddFlowSourcesInput) *MediaconnectAddFlowSourcesResult

    AddFlowVpcInterfaces(ctx workflow.Context, input *mediaconnect.AddFlowVpcInterfacesInput) (*mediaconnect.AddFlowVpcInterfacesOutput, error)
    AddFlowVpcInterfacesAsync(ctx workflow.Context, input *mediaconnect.AddFlowVpcInterfacesInput) *MediaconnectAddFlowVpcInterfacesResult

    CreateFlow(ctx workflow.Context, input *mediaconnect.CreateFlowInput) (*mediaconnect.CreateFlowOutput, error)
    CreateFlowAsync(ctx workflow.Context, input *mediaconnect.CreateFlowInput) *MediaconnectCreateFlowResult

    DeleteFlow(ctx workflow.Context, input *mediaconnect.DeleteFlowInput) (*mediaconnect.DeleteFlowOutput, error)
    DeleteFlowAsync(ctx workflow.Context, input *mediaconnect.DeleteFlowInput) *MediaconnectDeleteFlowResult

    DescribeFlow(ctx workflow.Context, input *mediaconnect.DescribeFlowInput) (*mediaconnect.DescribeFlowOutput, error)
    DescribeFlowAsync(ctx workflow.Context, input *mediaconnect.DescribeFlowInput) *MediaconnectDescribeFlowResult

    GrantFlowEntitlements(ctx workflow.Context, input *mediaconnect.GrantFlowEntitlementsInput) (*mediaconnect.GrantFlowEntitlementsOutput, error)
    GrantFlowEntitlementsAsync(ctx workflow.Context, input *mediaconnect.GrantFlowEntitlementsInput) *MediaconnectGrantFlowEntitlementsResult

    ListEntitlements(ctx workflow.Context, input *mediaconnect.ListEntitlementsInput) (*mediaconnect.ListEntitlementsOutput, error)
    ListEntitlementsAsync(ctx workflow.Context, input *mediaconnect.ListEntitlementsInput) *MediaconnectListEntitlementsResult

    ListFlows(ctx workflow.Context, input *mediaconnect.ListFlowsInput) (*mediaconnect.ListFlowsOutput, error)
    ListFlowsAsync(ctx workflow.Context, input *mediaconnect.ListFlowsInput) *MediaconnectListFlowsResult

    ListTagsForResource(ctx workflow.Context, input *mediaconnect.ListTagsForResourceInput) (*mediaconnect.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *mediaconnect.ListTagsForResourceInput) *MediaconnectListTagsForResourceResult

    RemoveFlowOutput(ctx workflow.Context, input *mediaconnect.RemoveFlowOutputInput) (*mediaconnect.RemoveFlowOutputOutput, error)
    RemoveFlowOutputAsync(ctx workflow.Context, input *mediaconnect.RemoveFlowOutputInput) *MediaconnectRemoveFlowOutputResult

    RemoveFlowSource(ctx workflow.Context, input *mediaconnect.RemoveFlowSourceInput) (*mediaconnect.RemoveFlowSourceOutput, error)
    RemoveFlowSourceAsync(ctx workflow.Context, input *mediaconnect.RemoveFlowSourceInput) *MediaconnectRemoveFlowSourceResult

    RemoveFlowVpcInterface(ctx workflow.Context, input *mediaconnect.RemoveFlowVpcInterfaceInput) (*mediaconnect.RemoveFlowVpcInterfaceOutput, error)
    RemoveFlowVpcInterfaceAsync(ctx workflow.Context, input *mediaconnect.RemoveFlowVpcInterfaceInput) *MediaconnectRemoveFlowVpcInterfaceResult

    RevokeFlowEntitlement(ctx workflow.Context, input *mediaconnect.RevokeFlowEntitlementInput) (*mediaconnect.RevokeFlowEntitlementOutput, error)
    RevokeFlowEntitlementAsync(ctx workflow.Context, input *mediaconnect.RevokeFlowEntitlementInput) *MediaconnectRevokeFlowEntitlementResult

    StartFlow(ctx workflow.Context, input *mediaconnect.StartFlowInput) (*mediaconnect.StartFlowOutput, error)
    StartFlowAsync(ctx workflow.Context, input *mediaconnect.StartFlowInput) *MediaconnectStartFlowResult

    StopFlow(ctx workflow.Context, input *mediaconnect.StopFlowInput) (*mediaconnect.StopFlowOutput, error)
    StopFlowAsync(ctx workflow.Context, input *mediaconnect.StopFlowInput) *MediaconnectStopFlowResult

    TagResource(ctx workflow.Context, input *mediaconnect.TagResourceInput) (*mediaconnect.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *mediaconnect.TagResourceInput) *MediaconnectTagResourceResult

    UntagResource(ctx workflow.Context, input *mediaconnect.UntagResourceInput) (*mediaconnect.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *mediaconnect.UntagResourceInput) *MediaconnectUntagResourceResult

    UpdateFlow(ctx workflow.Context, input *mediaconnect.UpdateFlowInput) (*mediaconnect.UpdateFlowOutput, error)
    UpdateFlowAsync(ctx workflow.Context, input *mediaconnect.UpdateFlowInput) *MediaconnectUpdateFlowResult

    UpdateFlowEntitlement(ctx workflow.Context, input *mediaconnect.UpdateFlowEntitlementInput) (*mediaconnect.UpdateFlowEntitlementOutput, error)
    UpdateFlowEntitlementAsync(ctx workflow.Context, input *mediaconnect.UpdateFlowEntitlementInput) *MediaconnectUpdateFlowEntitlementResult

    UpdateFlowOutput(ctx workflow.Context, input *mediaconnect.UpdateFlowOutputInput) (*mediaconnect.UpdateFlowOutputOutput, error)
    UpdateFlowOutputAsync(ctx workflow.Context, input *mediaconnect.UpdateFlowOutputInput) *MediaconnectUpdateFlowOutputResult

    UpdateFlowSource(ctx workflow.Context, input *mediaconnect.UpdateFlowSourceInput) (*mediaconnect.UpdateFlowSourceOutput, error)
    UpdateFlowSourceAsync(ctx workflow.Context, input *mediaconnect.UpdateFlowSourceInput) *MediaconnectUpdateFlowSourceResult
}
type MediaconnectAddFlowOutputsResult struct {
	Result workflow.Future
}

func (r *MediaconnectAddFlowOutputsResult) Get(ctx workflow.Context) (*mediaconnect.AddFlowOutputsOutput, error) {
    var output mediaconnect.AddFlowOutputsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconnectAddFlowSourcesResult struct {
	Result workflow.Future
}

func (r *MediaconnectAddFlowSourcesResult) Get(ctx workflow.Context) (*mediaconnect.AddFlowSourcesOutput, error) {
    var output mediaconnect.AddFlowSourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconnectAddFlowVpcInterfacesResult struct {
	Result workflow.Future
}

func (r *MediaconnectAddFlowVpcInterfacesResult) Get(ctx workflow.Context) (*mediaconnect.AddFlowVpcInterfacesOutput, error) {
    var output mediaconnect.AddFlowVpcInterfacesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconnectCreateFlowResult struct {
	Result workflow.Future
}

func (r *MediaconnectCreateFlowResult) Get(ctx workflow.Context) (*mediaconnect.CreateFlowOutput, error) {
    var output mediaconnect.CreateFlowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconnectDeleteFlowResult struct {
	Result workflow.Future
}

func (r *MediaconnectDeleteFlowResult) Get(ctx workflow.Context) (*mediaconnect.DeleteFlowOutput, error) {
    var output mediaconnect.DeleteFlowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconnectDescribeFlowResult struct {
	Result workflow.Future
}

func (r *MediaconnectDescribeFlowResult) Get(ctx workflow.Context) (*mediaconnect.DescribeFlowOutput, error) {
    var output mediaconnect.DescribeFlowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconnectGrantFlowEntitlementsResult struct {
	Result workflow.Future
}

func (r *MediaconnectGrantFlowEntitlementsResult) Get(ctx workflow.Context) (*mediaconnect.GrantFlowEntitlementsOutput, error) {
    var output mediaconnect.GrantFlowEntitlementsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconnectListEntitlementsResult struct {
	Result workflow.Future
}

func (r *MediaconnectListEntitlementsResult) Get(ctx workflow.Context) (*mediaconnect.ListEntitlementsOutput, error) {
    var output mediaconnect.ListEntitlementsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconnectListFlowsResult struct {
	Result workflow.Future
}

func (r *MediaconnectListFlowsResult) Get(ctx workflow.Context) (*mediaconnect.ListFlowsOutput, error) {
    var output mediaconnect.ListFlowsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconnectListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *MediaconnectListTagsForResourceResult) Get(ctx workflow.Context) (*mediaconnect.ListTagsForResourceOutput, error) {
    var output mediaconnect.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconnectRemoveFlowOutputResult struct {
	Result workflow.Future
}

func (r *MediaconnectRemoveFlowOutputResult) Get(ctx workflow.Context) (*mediaconnect.RemoveFlowOutputOutput, error) {
    var output mediaconnect.RemoveFlowOutputOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconnectRemoveFlowSourceResult struct {
	Result workflow.Future
}

func (r *MediaconnectRemoveFlowSourceResult) Get(ctx workflow.Context) (*mediaconnect.RemoveFlowSourceOutput, error) {
    var output mediaconnect.RemoveFlowSourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconnectRemoveFlowVpcInterfaceResult struct {
	Result workflow.Future
}

func (r *MediaconnectRemoveFlowVpcInterfaceResult) Get(ctx workflow.Context) (*mediaconnect.RemoveFlowVpcInterfaceOutput, error) {
    var output mediaconnect.RemoveFlowVpcInterfaceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconnectRevokeFlowEntitlementResult struct {
	Result workflow.Future
}

func (r *MediaconnectRevokeFlowEntitlementResult) Get(ctx workflow.Context) (*mediaconnect.RevokeFlowEntitlementOutput, error) {
    var output mediaconnect.RevokeFlowEntitlementOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconnectStartFlowResult struct {
	Result workflow.Future
}

func (r *MediaconnectStartFlowResult) Get(ctx workflow.Context) (*mediaconnect.StartFlowOutput, error) {
    var output mediaconnect.StartFlowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconnectStopFlowResult struct {
	Result workflow.Future
}

func (r *MediaconnectStopFlowResult) Get(ctx workflow.Context) (*mediaconnect.StopFlowOutput, error) {
    var output mediaconnect.StopFlowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconnectTagResourceResult struct {
	Result workflow.Future
}

func (r *MediaconnectTagResourceResult) Get(ctx workflow.Context) (*mediaconnect.TagResourceOutput, error) {
    var output mediaconnect.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconnectUntagResourceResult struct {
	Result workflow.Future
}

func (r *MediaconnectUntagResourceResult) Get(ctx workflow.Context) (*mediaconnect.UntagResourceOutput, error) {
    var output mediaconnect.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconnectUpdateFlowResult struct {
	Result workflow.Future
}

func (r *MediaconnectUpdateFlowResult) Get(ctx workflow.Context) (*mediaconnect.UpdateFlowOutput, error) {
    var output mediaconnect.UpdateFlowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconnectUpdateFlowEntitlementResult struct {
	Result workflow.Future
}

func (r *MediaconnectUpdateFlowEntitlementResult) Get(ctx workflow.Context) (*mediaconnect.UpdateFlowEntitlementOutput, error) {
    var output mediaconnect.UpdateFlowEntitlementOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconnectUpdateFlowOutputResult struct {
	Result workflow.Future
}

func (r *MediaconnectUpdateFlowOutputResult) Get(ctx workflow.Context) (*mediaconnect.UpdateFlowOutputOutput, error) {
    var output mediaconnect.UpdateFlowOutputOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaconnectUpdateFlowSourceResult struct {
	Result workflow.Future
}

func (r *MediaconnectUpdateFlowSourceResult) Get(ctx workflow.Context) (*mediaconnect.UpdateFlowSourceOutput, error) {
    var output mediaconnect.UpdateFlowSourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type MediaConnectStub struct {
    activities MediaConnectClient
}

func NewMediaConnectStub() MediaConnectClient {
    return &MediaConnectStub{}
}

func (a *MediaConnectStub) AddFlowOutputs(ctx workflow.Context, input *mediaconnect.AddFlowOutputsInput) (*mediaconnect.AddFlowOutputsOutput, error) {
    var output mediaconnect.AddFlowOutputsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddFlowOutputs, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConnectStub) AddFlowSources(ctx workflow.Context, input *mediaconnect.AddFlowSourcesInput) (*mediaconnect.AddFlowSourcesOutput, error) {
    var output mediaconnect.AddFlowSourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddFlowSources, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConnectStub) AddFlowVpcInterfaces(ctx workflow.Context, input *mediaconnect.AddFlowVpcInterfacesInput) (*mediaconnect.AddFlowVpcInterfacesOutput, error) {
    var output mediaconnect.AddFlowVpcInterfacesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddFlowVpcInterfaces, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConnectStub) CreateFlow(ctx workflow.Context, input *mediaconnect.CreateFlowInput) (*mediaconnect.CreateFlowOutput, error) {
    var output mediaconnect.CreateFlowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFlow, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConnectStub) DeleteFlow(ctx workflow.Context, input *mediaconnect.DeleteFlowInput) (*mediaconnect.DeleteFlowOutput, error) {
    var output mediaconnect.DeleteFlowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFlow, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConnectStub) DescribeFlow(ctx workflow.Context, input *mediaconnect.DescribeFlowInput) (*mediaconnect.DescribeFlowOutput, error) {
    var output mediaconnect.DescribeFlowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFlow, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConnectStub) GrantFlowEntitlements(ctx workflow.Context, input *mediaconnect.GrantFlowEntitlementsInput) (*mediaconnect.GrantFlowEntitlementsOutput, error) {
    var output mediaconnect.GrantFlowEntitlementsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GrantFlowEntitlements, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConnectStub) ListEntitlements(ctx workflow.Context, input *mediaconnect.ListEntitlementsInput) (*mediaconnect.ListEntitlementsOutput, error) {
    var output mediaconnect.ListEntitlementsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListEntitlements, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConnectStub) ListFlows(ctx workflow.Context, input *mediaconnect.ListFlowsInput) (*mediaconnect.ListFlowsOutput, error) {
    var output mediaconnect.ListFlowsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFlows, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConnectStub) ListTagsForResource(ctx workflow.Context, input *mediaconnect.ListTagsForResourceInput) (*mediaconnect.ListTagsForResourceOutput, error) {
    var output mediaconnect.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConnectStub) RemoveFlowOutput(ctx workflow.Context, input *mediaconnect.RemoveFlowOutputInput) (*mediaconnect.RemoveFlowOutputOutput, error) {
    var output mediaconnect.RemoveFlowOutputOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveFlowOutput, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConnectStub) RemoveFlowSource(ctx workflow.Context, input *mediaconnect.RemoveFlowSourceInput) (*mediaconnect.RemoveFlowSourceOutput, error) {
    var output mediaconnect.RemoveFlowSourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveFlowSource, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConnectStub) RemoveFlowVpcInterface(ctx workflow.Context, input *mediaconnect.RemoveFlowVpcInterfaceInput) (*mediaconnect.RemoveFlowVpcInterfaceOutput, error) {
    var output mediaconnect.RemoveFlowVpcInterfaceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveFlowVpcInterface, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConnectStub) RevokeFlowEntitlement(ctx workflow.Context, input *mediaconnect.RevokeFlowEntitlementInput) (*mediaconnect.RevokeFlowEntitlementOutput, error) {
    var output mediaconnect.RevokeFlowEntitlementOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RevokeFlowEntitlement, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConnectStub) StartFlow(ctx workflow.Context, input *mediaconnect.StartFlowInput) (*mediaconnect.StartFlowOutput, error) {
    var output mediaconnect.StartFlowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartFlow, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConnectStub) StopFlow(ctx workflow.Context, input *mediaconnect.StopFlowInput) (*mediaconnect.StopFlowOutput, error) {
    var output mediaconnect.StopFlowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopFlow, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConnectStub) TagResource(ctx workflow.Context, input *mediaconnect.TagResourceInput) (*mediaconnect.TagResourceOutput, error) {
    var output mediaconnect.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConnectStub) UntagResource(ctx workflow.Context, input *mediaconnect.UntagResourceInput) (*mediaconnect.UntagResourceOutput, error) {
    var output mediaconnect.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConnectStub) UpdateFlow(ctx workflow.Context, input *mediaconnect.UpdateFlowInput) (*mediaconnect.UpdateFlowOutput, error) {
    var output mediaconnect.UpdateFlowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFlow, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConnectStub) UpdateFlowEntitlement(ctx workflow.Context, input *mediaconnect.UpdateFlowEntitlementInput) (*mediaconnect.UpdateFlowEntitlementOutput, error) {
    var output mediaconnect.UpdateFlowEntitlementOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFlowEntitlement, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConnectStub) UpdateFlowOutput(ctx workflow.Context, input *mediaconnect.UpdateFlowOutputInput) (*mediaconnect.UpdateFlowOutputOutput, error) {
    var output mediaconnect.UpdateFlowOutputOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFlowOutput, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaConnectStub) UpdateFlowSource(ctx workflow.Context, input *mediaconnect.UpdateFlowSourceInput) (*mediaconnect.UpdateFlowSourceOutput, error) {
    var output mediaconnect.UpdateFlowSourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFlowSource, input).Get(ctx, &output)
    return &output, err
}
