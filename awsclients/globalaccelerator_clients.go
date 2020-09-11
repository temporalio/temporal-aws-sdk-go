package awsclients

import (
	"github.com/aws/aws-sdk-go/service/globalaccelerator"
	"go.temporal.io/sdk/workflow"
)

type GlobalAcceleratorClient interface {
       AdvertiseByoipCidr(ctx workflow.Context, input *globalaccelerator.AdvertiseByoipCidrInput) (*globalaccelerator.AdvertiseByoipCidrOutput, error)
       AdvertiseByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.AdvertiseByoipCidrInput) *GlobalacceleratorAdvertiseByoipCidrResult

       CreateAccelerator(ctx workflow.Context, input *globalaccelerator.CreateAcceleratorInput) (*globalaccelerator.CreateAcceleratorOutput, error)
       CreateAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.CreateAcceleratorInput) *GlobalacceleratorCreateAcceleratorResult

       CreateEndpointGroup(ctx workflow.Context, input *globalaccelerator.CreateEndpointGroupInput) (*globalaccelerator.CreateEndpointGroupOutput, error)
       CreateEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.CreateEndpointGroupInput) *GlobalacceleratorCreateEndpointGroupResult

       CreateListener(ctx workflow.Context, input *globalaccelerator.CreateListenerInput) (*globalaccelerator.CreateListenerOutput, error)
       CreateListenerAsync(ctx workflow.Context, input *globalaccelerator.CreateListenerInput) *GlobalacceleratorCreateListenerResult

       DeleteAccelerator(ctx workflow.Context, input *globalaccelerator.DeleteAcceleratorInput) (*globalaccelerator.DeleteAcceleratorOutput, error)
       DeleteAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.DeleteAcceleratorInput) *GlobalacceleratorDeleteAcceleratorResult

       DeleteEndpointGroup(ctx workflow.Context, input *globalaccelerator.DeleteEndpointGroupInput) (*globalaccelerator.DeleteEndpointGroupOutput, error)
       DeleteEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.DeleteEndpointGroupInput) *GlobalacceleratorDeleteEndpointGroupResult

       DeleteListener(ctx workflow.Context, input *globalaccelerator.DeleteListenerInput) (*globalaccelerator.DeleteListenerOutput, error)
       DeleteListenerAsync(ctx workflow.Context, input *globalaccelerator.DeleteListenerInput) *GlobalacceleratorDeleteListenerResult

       DeprovisionByoipCidr(ctx workflow.Context, input *globalaccelerator.DeprovisionByoipCidrInput) (*globalaccelerator.DeprovisionByoipCidrOutput, error)
       DeprovisionByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.DeprovisionByoipCidrInput) *GlobalacceleratorDeprovisionByoipCidrResult

       DescribeAccelerator(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorInput) (*globalaccelerator.DescribeAcceleratorOutput, error)
       DescribeAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorInput) *GlobalacceleratorDescribeAcceleratorResult

       DescribeAcceleratorAttributes(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorAttributesInput) (*globalaccelerator.DescribeAcceleratorAttributesOutput, error)
       DescribeAcceleratorAttributesAsync(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorAttributesInput) *GlobalacceleratorDescribeAcceleratorAttributesResult

       DescribeEndpointGroup(ctx workflow.Context, input *globalaccelerator.DescribeEndpointGroupInput) (*globalaccelerator.DescribeEndpointGroupOutput, error)
       DescribeEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.DescribeEndpointGroupInput) *GlobalacceleratorDescribeEndpointGroupResult

       DescribeListener(ctx workflow.Context, input *globalaccelerator.DescribeListenerInput) (*globalaccelerator.DescribeListenerOutput, error)
       DescribeListenerAsync(ctx workflow.Context, input *globalaccelerator.DescribeListenerInput) *GlobalacceleratorDescribeListenerResult

       ListAccelerators(ctx workflow.Context, input *globalaccelerator.ListAcceleratorsInput) (*globalaccelerator.ListAcceleratorsOutput, error)
       ListAcceleratorsAsync(ctx workflow.Context, input *globalaccelerator.ListAcceleratorsInput) *GlobalacceleratorListAcceleratorsResult

       ListByoipCidrs(ctx workflow.Context, input *globalaccelerator.ListByoipCidrsInput) (*globalaccelerator.ListByoipCidrsOutput, error)
       ListByoipCidrsAsync(ctx workflow.Context, input *globalaccelerator.ListByoipCidrsInput) *GlobalacceleratorListByoipCidrsResult

       ListEndpointGroups(ctx workflow.Context, input *globalaccelerator.ListEndpointGroupsInput) (*globalaccelerator.ListEndpointGroupsOutput, error)
       ListEndpointGroupsAsync(ctx workflow.Context, input *globalaccelerator.ListEndpointGroupsInput) *GlobalacceleratorListEndpointGroupsResult

       ListListeners(ctx workflow.Context, input *globalaccelerator.ListListenersInput) (*globalaccelerator.ListListenersOutput, error)
       ListListenersAsync(ctx workflow.Context, input *globalaccelerator.ListListenersInput) *GlobalacceleratorListListenersResult

       ListTagsForResource(ctx workflow.Context, input *globalaccelerator.ListTagsForResourceInput) (*globalaccelerator.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *globalaccelerator.ListTagsForResourceInput) *GlobalacceleratorListTagsForResourceResult

       ProvisionByoipCidr(ctx workflow.Context, input *globalaccelerator.ProvisionByoipCidrInput) (*globalaccelerator.ProvisionByoipCidrOutput, error)
       ProvisionByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.ProvisionByoipCidrInput) *GlobalacceleratorProvisionByoipCidrResult

       TagResource(ctx workflow.Context, input *globalaccelerator.TagResourceInput) (*globalaccelerator.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *globalaccelerator.TagResourceInput) *GlobalacceleratorTagResourceResult

       UntagResource(ctx workflow.Context, input *globalaccelerator.UntagResourceInput) (*globalaccelerator.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *globalaccelerator.UntagResourceInput) *GlobalacceleratorUntagResourceResult

       UpdateAccelerator(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorInput) (*globalaccelerator.UpdateAcceleratorOutput, error)
       UpdateAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorInput) *GlobalacceleratorUpdateAcceleratorResult

       UpdateAcceleratorAttributes(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorAttributesInput) (*globalaccelerator.UpdateAcceleratorAttributesOutput, error)
       UpdateAcceleratorAttributesAsync(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorAttributesInput) *GlobalacceleratorUpdateAcceleratorAttributesResult

       UpdateEndpointGroup(ctx workflow.Context, input *globalaccelerator.UpdateEndpointGroupInput) (*globalaccelerator.UpdateEndpointGroupOutput, error)
       UpdateEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.UpdateEndpointGroupInput) *GlobalacceleratorUpdateEndpointGroupResult

       UpdateListener(ctx workflow.Context, input *globalaccelerator.UpdateListenerInput) (*globalaccelerator.UpdateListenerOutput, error)
       UpdateListenerAsync(ctx workflow.Context, input *globalaccelerator.UpdateListenerInput) *GlobalacceleratorUpdateListenerResult

       WithdrawByoipCidr(ctx workflow.Context, input *globalaccelerator.WithdrawByoipCidrInput) (*globalaccelerator.WithdrawByoipCidrOutput, error)
       WithdrawByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.WithdrawByoipCidrInput) *GlobalacceleratorWithdrawByoipCidrResult
}

type GlobalAcceleratorStub struct {
}

func NewGlobalAcceleratorStub() GlobalAcceleratorClient {
    return &GlobalAcceleratorStub{}
}

type GlobalacceleratorAdvertiseByoipCidrResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorAdvertiseByoipCidrResult) Get(ctx workflow.Context) (*globalaccelerator.AdvertiseByoipCidrOutput, error) {
    var output globalaccelerator.AdvertiseByoipCidrOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorCreateAcceleratorResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorCreateAcceleratorResult) Get(ctx workflow.Context) (*globalaccelerator.CreateAcceleratorOutput, error) {
    var output globalaccelerator.CreateAcceleratorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorCreateEndpointGroupResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorCreateEndpointGroupResult) Get(ctx workflow.Context) (*globalaccelerator.CreateEndpointGroupOutput, error) {
    var output globalaccelerator.CreateEndpointGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorCreateListenerResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorCreateListenerResult) Get(ctx workflow.Context) (*globalaccelerator.CreateListenerOutput, error) {
    var output globalaccelerator.CreateListenerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorDeleteAcceleratorResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorDeleteAcceleratorResult) Get(ctx workflow.Context) (*globalaccelerator.DeleteAcceleratorOutput, error) {
    var output globalaccelerator.DeleteAcceleratorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorDeleteEndpointGroupResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorDeleteEndpointGroupResult) Get(ctx workflow.Context) (*globalaccelerator.DeleteEndpointGroupOutput, error) {
    var output globalaccelerator.DeleteEndpointGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorDeleteListenerResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorDeleteListenerResult) Get(ctx workflow.Context) (*globalaccelerator.DeleteListenerOutput, error) {
    var output globalaccelerator.DeleteListenerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorDeprovisionByoipCidrResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorDeprovisionByoipCidrResult) Get(ctx workflow.Context) (*globalaccelerator.DeprovisionByoipCidrOutput, error) {
    var output globalaccelerator.DeprovisionByoipCidrOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorDescribeAcceleratorResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorDescribeAcceleratorResult) Get(ctx workflow.Context) (*globalaccelerator.DescribeAcceleratorOutput, error) {
    var output globalaccelerator.DescribeAcceleratorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorDescribeAcceleratorAttributesResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorDescribeAcceleratorAttributesResult) Get(ctx workflow.Context) (*globalaccelerator.DescribeAcceleratorAttributesOutput, error) {
    var output globalaccelerator.DescribeAcceleratorAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorDescribeEndpointGroupResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorDescribeEndpointGroupResult) Get(ctx workflow.Context) (*globalaccelerator.DescribeEndpointGroupOutput, error) {
    var output globalaccelerator.DescribeEndpointGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorDescribeListenerResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorDescribeListenerResult) Get(ctx workflow.Context) (*globalaccelerator.DescribeListenerOutput, error) {
    var output globalaccelerator.DescribeListenerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorListAcceleratorsResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorListAcceleratorsResult) Get(ctx workflow.Context) (*globalaccelerator.ListAcceleratorsOutput, error) {
    var output globalaccelerator.ListAcceleratorsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorListByoipCidrsResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorListByoipCidrsResult) Get(ctx workflow.Context) (*globalaccelerator.ListByoipCidrsOutput, error) {
    var output globalaccelerator.ListByoipCidrsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorListEndpointGroupsResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorListEndpointGroupsResult) Get(ctx workflow.Context) (*globalaccelerator.ListEndpointGroupsOutput, error) {
    var output globalaccelerator.ListEndpointGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorListListenersResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorListListenersResult) Get(ctx workflow.Context) (*globalaccelerator.ListListenersOutput, error) {
    var output globalaccelerator.ListListenersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorListTagsForResourceResult) Get(ctx workflow.Context) (*globalaccelerator.ListTagsForResourceOutput, error) {
    var output globalaccelerator.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorProvisionByoipCidrResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorProvisionByoipCidrResult) Get(ctx workflow.Context) (*globalaccelerator.ProvisionByoipCidrOutput, error) {
    var output globalaccelerator.ProvisionByoipCidrOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorTagResourceResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorTagResourceResult) Get(ctx workflow.Context) (*globalaccelerator.TagResourceOutput, error) {
    var output globalaccelerator.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorUntagResourceResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorUntagResourceResult) Get(ctx workflow.Context) (*globalaccelerator.UntagResourceOutput, error) {
    var output globalaccelerator.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorUpdateAcceleratorResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorUpdateAcceleratorResult) Get(ctx workflow.Context) (*globalaccelerator.UpdateAcceleratorOutput, error) {
    var output globalaccelerator.UpdateAcceleratorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorUpdateAcceleratorAttributesResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorUpdateAcceleratorAttributesResult) Get(ctx workflow.Context) (*globalaccelerator.UpdateAcceleratorAttributesOutput, error) {
    var output globalaccelerator.UpdateAcceleratorAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorUpdateEndpointGroupResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorUpdateEndpointGroupResult) Get(ctx workflow.Context) (*globalaccelerator.UpdateEndpointGroupOutput, error) {
    var output globalaccelerator.UpdateEndpointGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorUpdateListenerResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorUpdateListenerResult) Get(ctx workflow.Context) (*globalaccelerator.UpdateListenerOutput, error) {
    var output globalaccelerator.UpdateListenerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type GlobalacceleratorWithdrawByoipCidrResult struct {
	Result workflow.Future
}

func (r *GlobalacceleratorWithdrawByoipCidrResult) Get(ctx workflow.Context) (*globalaccelerator.WithdrawByoipCidrOutput, error) {
    var output globalaccelerator.WithdrawByoipCidrOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) AdvertiseByoipCidr(ctx workflow.Context, input *globalaccelerator.AdvertiseByoipCidrInput) (*globalaccelerator.AdvertiseByoipCidrOutput, error) {
    var output globalaccelerator.AdvertiseByoipCidrOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.AdvertiseByoipCidr", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) AdvertiseByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.AdvertiseByoipCidrInput) *GlobalacceleratorAdvertiseByoipCidrResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.AdvertiseByoipCidr", input)
    return &GlobalacceleratorAdvertiseByoipCidrResult{Result: future}
}

func (a *GlobalAcceleratorStub) CreateAccelerator(ctx workflow.Context, input *globalaccelerator.CreateAcceleratorInput) (*globalaccelerator.CreateAcceleratorOutput, error) {
    var output globalaccelerator.CreateAcceleratorOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.CreateAccelerator", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) CreateAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.CreateAcceleratorInput) *GlobalacceleratorCreateAcceleratorResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.CreateAccelerator", input)
    return &GlobalacceleratorCreateAcceleratorResult{Result: future}
}

func (a *GlobalAcceleratorStub) CreateEndpointGroup(ctx workflow.Context, input *globalaccelerator.CreateEndpointGroupInput) (*globalaccelerator.CreateEndpointGroupOutput, error) {
    var output globalaccelerator.CreateEndpointGroupOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.CreateEndpointGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) CreateEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.CreateEndpointGroupInput) *GlobalacceleratorCreateEndpointGroupResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.CreateEndpointGroup", input)
    return &GlobalacceleratorCreateEndpointGroupResult{Result: future}
}

func (a *GlobalAcceleratorStub) CreateListener(ctx workflow.Context, input *globalaccelerator.CreateListenerInput) (*globalaccelerator.CreateListenerOutput, error) {
    var output globalaccelerator.CreateListenerOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.CreateListener", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) CreateListenerAsync(ctx workflow.Context, input *globalaccelerator.CreateListenerInput) *GlobalacceleratorCreateListenerResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.CreateListener", input)
    return &GlobalacceleratorCreateListenerResult{Result: future}
}

func (a *GlobalAcceleratorStub) DeleteAccelerator(ctx workflow.Context, input *globalaccelerator.DeleteAcceleratorInput) (*globalaccelerator.DeleteAcceleratorOutput, error) {
    var output globalaccelerator.DeleteAcceleratorOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.DeleteAccelerator", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) DeleteAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.DeleteAcceleratorInput) *GlobalacceleratorDeleteAcceleratorResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.DeleteAccelerator", input)
    return &GlobalacceleratorDeleteAcceleratorResult{Result: future}
}

func (a *GlobalAcceleratorStub) DeleteEndpointGroup(ctx workflow.Context, input *globalaccelerator.DeleteEndpointGroupInput) (*globalaccelerator.DeleteEndpointGroupOutput, error) {
    var output globalaccelerator.DeleteEndpointGroupOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.DeleteEndpointGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) DeleteEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.DeleteEndpointGroupInput) *GlobalacceleratorDeleteEndpointGroupResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.DeleteEndpointGroup", input)
    return &GlobalacceleratorDeleteEndpointGroupResult{Result: future}
}

func (a *GlobalAcceleratorStub) DeleteListener(ctx workflow.Context, input *globalaccelerator.DeleteListenerInput) (*globalaccelerator.DeleteListenerOutput, error) {
    var output globalaccelerator.DeleteListenerOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.DeleteListener", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) DeleteListenerAsync(ctx workflow.Context, input *globalaccelerator.DeleteListenerInput) *GlobalacceleratorDeleteListenerResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.DeleteListener", input)
    return &GlobalacceleratorDeleteListenerResult{Result: future}
}

func (a *GlobalAcceleratorStub) DeprovisionByoipCidr(ctx workflow.Context, input *globalaccelerator.DeprovisionByoipCidrInput) (*globalaccelerator.DeprovisionByoipCidrOutput, error) {
    var output globalaccelerator.DeprovisionByoipCidrOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.DeprovisionByoipCidr", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) DeprovisionByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.DeprovisionByoipCidrInput) *GlobalacceleratorDeprovisionByoipCidrResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.DeprovisionByoipCidr", input)
    return &GlobalacceleratorDeprovisionByoipCidrResult{Result: future}
}

func (a *GlobalAcceleratorStub) DescribeAccelerator(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorInput) (*globalaccelerator.DescribeAcceleratorOutput, error) {
    var output globalaccelerator.DescribeAcceleratorOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.DescribeAccelerator", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) DescribeAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorInput) *GlobalacceleratorDescribeAcceleratorResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.DescribeAccelerator", input)
    return &GlobalacceleratorDescribeAcceleratorResult{Result: future}
}

func (a *GlobalAcceleratorStub) DescribeAcceleratorAttributes(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorAttributesInput) (*globalaccelerator.DescribeAcceleratorAttributesOutput, error) {
    var output globalaccelerator.DescribeAcceleratorAttributesOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.DescribeAcceleratorAttributes", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) DescribeAcceleratorAttributesAsync(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorAttributesInput) *GlobalacceleratorDescribeAcceleratorAttributesResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.DescribeAcceleratorAttributes", input)
    return &GlobalacceleratorDescribeAcceleratorAttributesResult{Result: future}
}

func (a *GlobalAcceleratorStub) DescribeEndpointGroup(ctx workflow.Context, input *globalaccelerator.DescribeEndpointGroupInput) (*globalaccelerator.DescribeEndpointGroupOutput, error) {
    var output globalaccelerator.DescribeEndpointGroupOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.DescribeEndpointGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) DescribeEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.DescribeEndpointGroupInput) *GlobalacceleratorDescribeEndpointGroupResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.DescribeEndpointGroup", input)
    return &GlobalacceleratorDescribeEndpointGroupResult{Result: future}
}

func (a *GlobalAcceleratorStub) DescribeListener(ctx workflow.Context, input *globalaccelerator.DescribeListenerInput) (*globalaccelerator.DescribeListenerOutput, error) {
    var output globalaccelerator.DescribeListenerOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.DescribeListener", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) DescribeListenerAsync(ctx workflow.Context, input *globalaccelerator.DescribeListenerInput) *GlobalacceleratorDescribeListenerResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.DescribeListener", input)
    return &GlobalacceleratorDescribeListenerResult{Result: future}
}

func (a *GlobalAcceleratorStub) ListAccelerators(ctx workflow.Context, input *globalaccelerator.ListAcceleratorsInput) (*globalaccelerator.ListAcceleratorsOutput, error) {
    var output globalaccelerator.ListAcceleratorsOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.ListAccelerators", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) ListAcceleratorsAsync(ctx workflow.Context, input *globalaccelerator.ListAcceleratorsInput) *GlobalacceleratorListAcceleratorsResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.ListAccelerators", input)
    return &GlobalacceleratorListAcceleratorsResult{Result: future}
}

func (a *GlobalAcceleratorStub) ListByoipCidrs(ctx workflow.Context, input *globalaccelerator.ListByoipCidrsInput) (*globalaccelerator.ListByoipCidrsOutput, error) {
    var output globalaccelerator.ListByoipCidrsOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.ListByoipCidrs", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) ListByoipCidrsAsync(ctx workflow.Context, input *globalaccelerator.ListByoipCidrsInput) *GlobalacceleratorListByoipCidrsResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.ListByoipCidrs", input)
    return &GlobalacceleratorListByoipCidrsResult{Result: future}
}

func (a *GlobalAcceleratorStub) ListEndpointGroups(ctx workflow.Context, input *globalaccelerator.ListEndpointGroupsInput) (*globalaccelerator.ListEndpointGroupsOutput, error) {
    var output globalaccelerator.ListEndpointGroupsOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.ListEndpointGroups", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) ListEndpointGroupsAsync(ctx workflow.Context, input *globalaccelerator.ListEndpointGroupsInput) *GlobalacceleratorListEndpointGroupsResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.ListEndpointGroups", input)
    return &GlobalacceleratorListEndpointGroupsResult{Result: future}
}

func (a *GlobalAcceleratorStub) ListListeners(ctx workflow.Context, input *globalaccelerator.ListListenersInput) (*globalaccelerator.ListListenersOutput, error) {
    var output globalaccelerator.ListListenersOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.ListListeners", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) ListListenersAsync(ctx workflow.Context, input *globalaccelerator.ListListenersInput) *GlobalacceleratorListListenersResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.ListListeners", input)
    return &GlobalacceleratorListListenersResult{Result: future}
}

func (a *GlobalAcceleratorStub) ListTagsForResource(ctx workflow.Context, input *globalaccelerator.ListTagsForResourceInput) (*globalaccelerator.ListTagsForResourceOutput, error) {
    var output globalaccelerator.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.ListTagsForResource", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) ListTagsForResourceAsync(ctx workflow.Context, input *globalaccelerator.ListTagsForResourceInput) *GlobalacceleratorListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.ListTagsForResource", input)
    return &GlobalacceleratorListTagsForResourceResult{Result: future}
}

func (a *GlobalAcceleratorStub) ProvisionByoipCidr(ctx workflow.Context, input *globalaccelerator.ProvisionByoipCidrInput) (*globalaccelerator.ProvisionByoipCidrOutput, error) {
    var output globalaccelerator.ProvisionByoipCidrOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.ProvisionByoipCidr", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) ProvisionByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.ProvisionByoipCidrInput) *GlobalacceleratorProvisionByoipCidrResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.ProvisionByoipCidr", input)
    return &GlobalacceleratorProvisionByoipCidrResult{Result: future}
}

func (a *GlobalAcceleratorStub) TagResource(ctx workflow.Context, input *globalaccelerator.TagResourceInput) (*globalaccelerator.TagResourceOutput, error) {
    var output globalaccelerator.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.TagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) TagResourceAsync(ctx workflow.Context, input *globalaccelerator.TagResourceInput) *GlobalacceleratorTagResourceResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.TagResource", input)
    return &GlobalacceleratorTagResourceResult{Result: future}
}

func (a *GlobalAcceleratorStub) UntagResource(ctx workflow.Context, input *globalaccelerator.UntagResourceInput) (*globalaccelerator.UntagResourceOutput, error) {
    var output globalaccelerator.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.UntagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) UntagResourceAsync(ctx workflow.Context, input *globalaccelerator.UntagResourceInput) *GlobalacceleratorUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.UntagResource", input)
    return &GlobalacceleratorUntagResourceResult{Result: future}
}

func (a *GlobalAcceleratorStub) UpdateAccelerator(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorInput) (*globalaccelerator.UpdateAcceleratorOutput, error) {
    var output globalaccelerator.UpdateAcceleratorOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.UpdateAccelerator", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) UpdateAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorInput) *GlobalacceleratorUpdateAcceleratorResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.UpdateAccelerator", input)
    return &GlobalacceleratorUpdateAcceleratorResult{Result: future}
}

func (a *GlobalAcceleratorStub) UpdateAcceleratorAttributes(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorAttributesInput) (*globalaccelerator.UpdateAcceleratorAttributesOutput, error) {
    var output globalaccelerator.UpdateAcceleratorAttributesOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.UpdateAcceleratorAttributes", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) UpdateAcceleratorAttributesAsync(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorAttributesInput) *GlobalacceleratorUpdateAcceleratorAttributesResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.UpdateAcceleratorAttributes", input)
    return &GlobalacceleratorUpdateAcceleratorAttributesResult{Result: future}
}

func (a *GlobalAcceleratorStub) UpdateEndpointGroup(ctx workflow.Context, input *globalaccelerator.UpdateEndpointGroupInput) (*globalaccelerator.UpdateEndpointGroupOutput, error) {
    var output globalaccelerator.UpdateEndpointGroupOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.UpdateEndpointGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) UpdateEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.UpdateEndpointGroupInput) *GlobalacceleratorUpdateEndpointGroupResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.UpdateEndpointGroup", input)
    return &GlobalacceleratorUpdateEndpointGroupResult{Result: future}
}

func (a *GlobalAcceleratorStub) UpdateListener(ctx workflow.Context, input *globalaccelerator.UpdateListenerInput) (*globalaccelerator.UpdateListenerOutput, error) {
    var output globalaccelerator.UpdateListenerOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.UpdateListener", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) UpdateListenerAsync(ctx workflow.Context, input *globalaccelerator.UpdateListenerInput) *GlobalacceleratorUpdateListenerResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.UpdateListener", input)
    return &GlobalacceleratorUpdateListenerResult{Result: future}
}

func (a *GlobalAcceleratorStub) WithdrawByoipCidr(ctx workflow.Context, input *globalaccelerator.WithdrawByoipCidrInput) (*globalaccelerator.WithdrawByoipCidrOutput, error) {
    var output globalaccelerator.WithdrawByoipCidrOutput
    err := workflow.ExecuteActivity(ctx, "GlobalAccelerator.WithdrawByoipCidr", input).Get(ctx, &output)
    return &output, err
}

func (a *GlobalAcceleratorStub) WithdrawByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.WithdrawByoipCidrInput) *GlobalacceleratorWithdrawByoipCidrResult {
    future := workflow.ExecuteActivity(ctx, "GlobalAccelerator.WithdrawByoipCidr", input)
    return &GlobalacceleratorWithdrawByoipCidrResult{Result: future}
}
