// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package globalacceleratorstub

import (
	"github.com/aws/aws-sdk-go/service/globalaccelerator"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type AdvertiseByoipCidrFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AdvertiseByoipCidrFuture) Get(ctx workflow.Context) (*globalaccelerator.AdvertiseByoipCidrOutput, error) {
	var output globalaccelerator.AdvertiseByoipCidrOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateAcceleratorFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateAcceleratorFuture) Get(ctx workflow.Context) (*globalaccelerator.CreateAcceleratorOutput, error) {
	var output globalaccelerator.CreateAcceleratorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateEndpointGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateEndpointGroupFuture) Get(ctx workflow.Context) (*globalaccelerator.CreateEndpointGroupOutput, error) {
	var output globalaccelerator.CreateEndpointGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateListenerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateListenerFuture) Get(ctx workflow.Context) (*globalaccelerator.CreateListenerOutput, error) {
	var output globalaccelerator.CreateListenerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteAcceleratorFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteAcceleratorFuture) Get(ctx workflow.Context) (*globalaccelerator.DeleteAcceleratorOutput, error) {
	var output globalaccelerator.DeleteAcceleratorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteEndpointGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteEndpointGroupFuture) Get(ctx workflow.Context) (*globalaccelerator.DeleteEndpointGroupOutput, error) {
	var output globalaccelerator.DeleteEndpointGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteListenerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteListenerFuture) Get(ctx workflow.Context) (*globalaccelerator.DeleteListenerOutput, error) {
	var output globalaccelerator.DeleteListenerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeprovisionByoipCidrFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeprovisionByoipCidrFuture) Get(ctx workflow.Context) (*globalaccelerator.DeprovisionByoipCidrOutput, error) {
	var output globalaccelerator.DeprovisionByoipCidrOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeAcceleratorFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeAcceleratorFuture) Get(ctx workflow.Context) (*globalaccelerator.DescribeAcceleratorOutput, error) {
	var output globalaccelerator.DescribeAcceleratorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeAcceleratorAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeAcceleratorAttributesFuture) Get(ctx workflow.Context) (*globalaccelerator.DescribeAcceleratorAttributesOutput, error) {
	var output globalaccelerator.DescribeAcceleratorAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeEndpointGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeEndpointGroupFuture) Get(ctx workflow.Context) (*globalaccelerator.DescribeEndpointGroupOutput, error) {
	var output globalaccelerator.DescribeEndpointGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeListenerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeListenerFuture) Get(ctx workflow.Context) (*globalaccelerator.DescribeListenerOutput, error) {
	var output globalaccelerator.DescribeListenerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListAcceleratorsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListAcceleratorsFuture) Get(ctx workflow.Context) (*globalaccelerator.ListAcceleratorsOutput, error) {
	var output globalaccelerator.ListAcceleratorsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListByoipCidrsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListByoipCidrsFuture) Get(ctx workflow.Context) (*globalaccelerator.ListByoipCidrsOutput, error) {
	var output globalaccelerator.ListByoipCidrsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListEndpointGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListEndpointGroupsFuture) Get(ctx workflow.Context) (*globalaccelerator.ListEndpointGroupsOutput, error) {
	var output globalaccelerator.ListEndpointGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListListenersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListListenersFuture) Get(ctx workflow.Context) (*globalaccelerator.ListListenersOutput, error) {
	var output globalaccelerator.ListListenersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*globalaccelerator.ListTagsForResourceOutput, error) {
	var output globalaccelerator.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ProvisionByoipCidrFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ProvisionByoipCidrFuture) Get(ctx workflow.Context) (*globalaccelerator.ProvisionByoipCidrOutput, error) {
	var output globalaccelerator.ProvisionByoipCidrOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*globalaccelerator.TagResourceOutput, error) {
	var output globalaccelerator.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*globalaccelerator.UntagResourceOutput, error) {
	var output globalaccelerator.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateAcceleratorFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateAcceleratorFuture) Get(ctx workflow.Context) (*globalaccelerator.UpdateAcceleratorOutput, error) {
	var output globalaccelerator.UpdateAcceleratorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateAcceleratorAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateAcceleratorAttributesFuture) Get(ctx workflow.Context) (*globalaccelerator.UpdateAcceleratorAttributesOutput, error) {
	var output globalaccelerator.UpdateAcceleratorAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateEndpointGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateEndpointGroupFuture) Get(ctx workflow.Context) (*globalaccelerator.UpdateEndpointGroupOutput, error) {
	var output globalaccelerator.UpdateEndpointGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateListenerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateListenerFuture) Get(ctx workflow.Context) (*globalaccelerator.UpdateListenerOutput, error) {
	var output globalaccelerator.UpdateListenerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WithdrawByoipCidrFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WithdrawByoipCidrFuture) Get(ctx workflow.Context) (*globalaccelerator.WithdrawByoipCidrOutput, error) {
	var output globalaccelerator.WithdrawByoipCidrOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AdvertiseByoipCidr(ctx workflow.Context, input *globalaccelerator.AdvertiseByoipCidrInput) (*globalaccelerator.AdvertiseByoipCidrOutput, error) {
	var output globalaccelerator.AdvertiseByoipCidrOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.AdvertiseByoipCidr", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AdvertiseByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.AdvertiseByoipCidrInput) *AdvertiseByoipCidrFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.AdvertiseByoipCidr", input)
	return &AdvertiseByoipCidrFuture{Future: future}
}

func (a *stub) CreateAccelerator(ctx workflow.Context, input *globalaccelerator.CreateAcceleratorInput) (*globalaccelerator.CreateAcceleratorOutput, error) {
	var output globalaccelerator.CreateAcceleratorOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.CreateAccelerator", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.CreateAcceleratorInput) *CreateAcceleratorFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.CreateAccelerator", input)
	return &CreateAcceleratorFuture{Future: future}
}

func (a *stub) CreateEndpointGroup(ctx workflow.Context, input *globalaccelerator.CreateEndpointGroupInput) (*globalaccelerator.CreateEndpointGroupOutput, error) {
	var output globalaccelerator.CreateEndpointGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.CreateEndpointGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.CreateEndpointGroupInput) *CreateEndpointGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.CreateEndpointGroup", input)
	return &CreateEndpointGroupFuture{Future: future}
}

func (a *stub) CreateListener(ctx workflow.Context, input *globalaccelerator.CreateListenerInput) (*globalaccelerator.CreateListenerOutput, error) {
	var output globalaccelerator.CreateListenerOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.CreateListener", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateListenerAsync(ctx workflow.Context, input *globalaccelerator.CreateListenerInput) *CreateListenerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.CreateListener", input)
	return &CreateListenerFuture{Future: future}
}

func (a *stub) DeleteAccelerator(ctx workflow.Context, input *globalaccelerator.DeleteAcceleratorInput) (*globalaccelerator.DeleteAcceleratorOutput, error) {
	var output globalaccelerator.DeleteAcceleratorOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DeleteAccelerator", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.DeleteAcceleratorInput) *DeleteAcceleratorFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DeleteAccelerator", input)
	return &DeleteAcceleratorFuture{Future: future}
}

func (a *stub) DeleteEndpointGroup(ctx workflow.Context, input *globalaccelerator.DeleteEndpointGroupInput) (*globalaccelerator.DeleteEndpointGroupOutput, error) {
	var output globalaccelerator.DeleteEndpointGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DeleteEndpointGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.DeleteEndpointGroupInput) *DeleteEndpointGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DeleteEndpointGroup", input)
	return &DeleteEndpointGroupFuture{Future: future}
}

func (a *stub) DeleteListener(ctx workflow.Context, input *globalaccelerator.DeleteListenerInput) (*globalaccelerator.DeleteListenerOutput, error) {
	var output globalaccelerator.DeleteListenerOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DeleteListener", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteListenerAsync(ctx workflow.Context, input *globalaccelerator.DeleteListenerInput) *DeleteListenerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DeleteListener", input)
	return &DeleteListenerFuture{Future: future}
}

func (a *stub) DeprovisionByoipCidr(ctx workflow.Context, input *globalaccelerator.DeprovisionByoipCidrInput) (*globalaccelerator.DeprovisionByoipCidrOutput, error) {
	var output globalaccelerator.DeprovisionByoipCidrOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DeprovisionByoipCidr", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeprovisionByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.DeprovisionByoipCidrInput) *DeprovisionByoipCidrFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DeprovisionByoipCidr", input)
	return &DeprovisionByoipCidrFuture{Future: future}
}

func (a *stub) DescribeAccelerator(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorInput) (*globalaccelerator.DescribeAcceleratorOutput, error) {
	var output globalaccelerator.DescribeAcceleratorOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DescribeAccelerator", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorInput) *DescribeAcceleratorFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DescribeAccelerator", input)
	return &DescribeAcceleratorFuture{Future: future}
}

func (a *stub) DescribeAcceleratorAttributes(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorAttributesInput) (*globalaccelerator.DescribeAcceleratorAttributesOutput, error) {
	var output globalaccelerator.DescribeAcceleratorAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DescribeAcceleratorAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAcceleratorAttributesAsync(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorAttributesInput) *DescribeAcceleratorAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DescribeAcceleratorAttributes", input)
	return &DescribeAcceleratorAttributesFuture{Future: future}
}

func (a *stub) DescribeEndpointGroup(ctx workflow.Context, input *globalaccelerator.DescribeEndpointGroupInput) (*globalaccelerator.DescribeEndpointGroupOutput, error) {
	var output globalaccelerator.DescribeEndpointGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DescribeEndpointGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.DescribeEndpointGroupInput) *DescribeEndpointGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DescribeEndpointGroup", input)
	return &DescribeEndpointGroupFuture{Future: future}
}

func (a *stub) DescribeListener(ctx workflow.Context, input *globalaccelerator.DescribeListenerInput) (*globalaccelerator.DescribeListenerOutput, error) {
	var output globalaccelerator.DescribeListenerOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DescribeListener", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeListenerAsync(ctx workflow.Context, input *globalaccelerator.DescribeListenerInput) *DescribeListenerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DescribeListener", input)
	return &DescribeListenerFuture{Future: future}
}

func (a *stub) ListAccelerators(ctx workflow.Context, input *globalaccelerator.ListAcceleratorsInput) (*globalaccelerator.ListAcceleratorsOutput, error) {
	var output globalaccelerator.ListAcceleratorsOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ListAccelerators", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAcceleratorsAsync(ctx workflow.Context, input *globalaccelerator.ListAcceleratorsInput) *ListAcceleratorsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ListAccelerators", input)
	return &ListAcceleratorsFuture{Future: future}
}

func (a *stub) ListByoipCidrs(ctx workflow.Context, input *globalaccelerator.ListByoipCidrsInput) (*globalaccelerator.ListByoipCidrsOutput, error) {
	var output globalaccelerator.ListByoipCidrsOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ListByoipCidrs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListByoipCidrsAsync(ctx workflow.Context, input *globalaccelerator.ListByoipCidrsInput) *ListByoipCidrsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ListByoipCidrs", input)
	return &ListByoipCidrsFuture{Future: future}
}

func (a *stub) ListEndpointGroups(ctx workflow.Context, input *globalaccelerator.ListEndpointGroupsInput) (*globalaccelerator.ListEndpointGroupsOutput, error) {
	var output globalaccelerator.ListEndpointGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ListEndpointGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListEndpointGroupsAsync(ctx workflow.Context, input *globalaccelerator.ListEndpointGroupsInput) *ListEndpointGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ListEndpointGroups", input)
	return &ListEndpointGroupsFuture{Future: future}
}

func (a *stub) ListListeners(ctx workflow.Context, input *globalaccelerator.ListListenersInput) (*globalaccelerator.ListListenersOutput, error) {
	var output globalaccelerator.ListListenersOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ListListeners", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListListenersAsync(ctx workflow.Context, input *globalaccelerator.ListListenersInput) *ListListenersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ListListeners", input)
	return &ListListenersFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *globalaccelerator.ListTagsForResourceInput) (*globalaccelerator.ListTagsForResourceOutput, error) {
	var output globalaccelerator.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *globalaccelerator.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) ProvisionByoipCidr(ctx workflow.Context, input *globalaccelerator.ProvisionByoipCidrInput) (*globalaccelerator.ProvisionByoipCidrOutput, error) {
	var output globalaccelerator.ProvisionByoipCidrOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ProvisionByoipCidr", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ProvisionByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.ProvisionByoipCidrInput) *ProvisionByoipCidrFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ProvisionByoipCidr", input)
	return &ProvisionByoipCidrFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *globalaccelerator.TagResourceInput) (*globalaccelerator.TagResourceOutput, error) {
	var output globalaccelerator.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *globalaccelerator.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *globalaccelerator.UntagResourceInput) (*globalaccelerator.UntagResourceOutput, error) {
	var output globalaccelerator.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *globalaccelerator.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.UntagResource", input)
	return &UntagResourceFuture{Future: future}
}

func (a *stub) UpdateAccelerator(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorInput) (*globalaccelerator.UpdateAcceleratorOutput, error) {
	var output globalaccelerator.UpdateAcceleratorOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.UpdateAccelerator", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorInput) *UpdateAcceleratorFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.UpdateAccelerator", input)
	return &UpdateAcceleratorFuture{Future: future}
}

func (a *stub) UpdateAcceleratorAttributes(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorAttributesInput) (*globalaccelerator.UpdateAcceleratorAttributesOutput, error) {
	var output globalaccelerator.UpdateAcceleratorAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.UpdateAcceleratorAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateAcceleratorAttributesAsync(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorAttributesInput) *UpdateAcceleratorAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.UpdateAcceleratorAttributes", input)
	return &UpdateAcceleratorAttributesFuture{Future: future}
}

func (a *stub) UpdateEndpointGroup(ctx workflow.Context, input *globalaccelerator.UpdateEndpointGroupInput) (*globalaccelerator.UpdateEndpointGroupOutput, error) {
	var output globalaccelerator.UpdateEndpointGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.UpdateEndpointGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.UpdateEndpointGroupInput) *UpdateEndpointGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.UpdateEndpointGroup", input)
	return &UpdateEndpointGroupFuture{Future: future}
}

func (a *stub) UpdateListener(ctx workflow.Context, input *globalaccelerator.UpdateListenerInput) (*globalaccelerator.UpdateListenerOutput, error) {
	var output globalaccelerator.UpdateListenerOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.UpdateListener", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateListenerAsync(ctx workflow.Context, input *globalaccelerator.UpdateListenerInput) *UpdateListenerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.UpdateListener", input)
	return &UpdateListenerFuture{Future: future}
}

func (a *stub) WithdrawByoipCidr(ctx workflow.Context, input *globalaccelerator.WithdrawByoipCidrInput) (*globalaccelerator.WithdrawByoipCidrOutput, error) {
	var output globalaccelerator.WithdrawByoipCidrOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.WithdrawByoipCidr", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) WithdrawByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.WithdrawByoipCidrInput) *WithdrawByoipCidrFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.WithdrawByoipCidr", input)
	return &WithdrawByoipCidrFuture{Future: future}
}
