// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package eventbridgestub

import (
	"github.com/aws/aws-sdk-go/service/eventbridge"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type ActivateEventSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ActivateEventSourceFuture) Get(ctx workflow.Context) (*eventbridge.ActivateEventSourceOutput, error) {
	var output eventbridge.ActivateEventSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateEventBusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateEventBusFuture) Get(ctx workflow.Context) (*eventbridge.CreateEventBusOutput, error) {
	var output eventbridge.CreateEventBusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreatePartnerEventSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreatePartnerEventSourceFuture) Get(ctx workflow.Context) (*eventbridge.CreatePartnerEventSourceOutput, error) {
	var output eventbridge.CreatePartnerEventSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeactivateEventSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeactivateEventSourceFuture) Get(ctx workflow.Context) (*eventbridge.DeactivateEventSourceOutput, error) {
	var output eventbridge.DeactivateEventSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteEventBusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteEventBusFuture) Get(ctx workflow.Context) (*eventbridge.DeleteEventBusOutput, error) {
	var output eventbridge.DeleteEventBusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeletePartnerEventSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeletePartnerEventSourceFuture) Get(ctx workflow.Context) (*eventbridge.DeletePartnerEventSourceOutput, error) {
	var output eventbridge.DeletePartnerEventSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteRuleFuture) Get(ctx workflow.Context) (*eventbridge.DeleteRuleOutput, error) {
	var output eventbridge.DeleteRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeEventBusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeEventBusFuture) Get(ctx workflow.Context) (*eventbridge.DescribeEventBusOutput, error) {
	var output eventbridge.DescribeEventBusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeEventSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeEventSourceFuture) Get(ctx workflow.Context) (*eventbridge.DescribeEventSourceOutput, error) {
	var output eventbridge.DescribeEventSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribePartnerEventSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribePartnerEventSourceFuture) Get(ctx workflow.Context) (*eventbridge.DescribePartnerEventSourceOutput, error) {
	var output eventbridge.DescribePartnerEventSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeRuleFuture) Get(ctx workflow.Context) (*eventbridge.DescribeRuleOutput, error) {
	var output eventbridge.DescribeRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DisableRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DisableRuleFuture) Get(ctx workflow.Context) (*eventbridge.DisableRuleOutput, error) {
	var output eventbridge.DisableRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EnableRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EnableRuleFuture) Get(ctx workflow.Context) (*eventbridge.EnableRuleOutput, error) {
	var output eventbridge.EnableRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListEventBusesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListEventBusesFuture) Get(ctx workflow.Context) (*eventbridge.ListEventBusesOutput, error) {
	var output eventbridge.ListEventBusesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListEventSourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListEventSourcesFuture) Get(ctx workflow.Context) (*eventbridge.ListEventSourcesOutput, error) {
	var output eventbridge.ListEventSourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListPartnerEventSourceAccountsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListPartnerEventSourceAccountsFuture) Get(ctx workflow.Context) (*eventbridge.ListPartnerEventSourceAccountsOutput, error) {
	var output eventbridge.ListPartnerEventSourceAccountsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListPartnerEventSourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListPartnerEventSourcesFuture) Get(ctx workflow.Context) (*eventbridge.ListPartnerEventSourcesOutput, error) {
	var output eventbridge.ListPartnerEventSourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListRuleNamesByTargetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListRuleNamesByTargetFuture) Get(ctx workflow.Context) (*eventbridge.ListRuleNamesByTargetOutput, error) {
	var output eventbridge.ListRuleNamesByTargetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListRulesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListRulesFuture) Get(ctx workflow.Context) (*eventbridge.ListRulesOutput, error) {
	var output eventbridge.ListRulesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*eventbridge.ListTagsForResourceOutput, error) {
	var output eventbridge.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTargetsByRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTargetsByRuleFuture) Get(ctx workflow.Context) (*eventbridge.ListTargetsByRuleOutput, error) {
	var output eventbridge.ListTargetsByRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutEventsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutEventsFuture) Get(ctx workflow.Context) (*eventbridge.PutEventsOutput, error) {
	var output eventbridge.PutEventsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutPartnerEventsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutPartnerEventsFuture) Get(ctx workflow.Context) (*eventbridge.PutPartnerEventsOutput, error) {
	var output eventbridge.PutPartnerEventsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutPermissionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutPermissionFuture) Get(ctx workflow.Context) (*eventbridge.PutPermissionOutput, error) {
	var output eventbridge.PutPermissionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutRuleFuture) Get(ctx workflow.Context) (*eventbridge.PutRuleOutput, error) {
	var output eventbridge.PutRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutTargetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutTargetsFuture) Get(ctx workflow.Context) (*eventbridge.PutTargetsOutput, error) {
	var output eventbridge.PutTargetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RemovePermissionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RemovePermissionFuture) Get(ctx workflow.Context) (*eventbridge.RemovePermissionOutput, error) {
	var output eventbridge.RemovePermissionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RemoveTargetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RemoveTargetsFuture) Get(ctx workflow.Context) (*eventbridge.RemoveTargetsOutput, error) {
	var output eventbridge.RemoveTargetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*eventbridge.TagResourceOutput, error) {
	var output eventbridge.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TestEventPatternFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TestEventPatternFuture) Get(ctx workflow.Context) (*eventbridge.TestEventPatternOutput, error) {
	var output eventbridge.TestEventPatternOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*eventbridge.UntagResourceOutput, error) {
	var output eventbridge.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) ActivateEventSource(ctx workflow.Context, input *eventbridge.ActivateEventSourceInput) (*eventbridge.ActivateEventSourceOutput, error) {
	var output eventbridge.ActivateEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.ActivateEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ActivateEventSourceAsync(ctx workflow.Context, input *eventbridge.ActivateEventSourceInput) *ActivateEventSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.ActivateEventSource", input)
	return &ActivateEventSourceFuture{Future: future}
}

func (a *stub) CreateEventBus(ctx workflow.Context, input *eventbridge.CreateEventBusInput) (*eventbridge.CreateEventBusOutput, error) {
	var output eventbridge.CreateEventBusOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.CreateEventBus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateEventBusAsync(ctx workflow.Context, input *eventbridge.CreateEventBusInput) *CreateEventBusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.CreateEventBus", input)
	return &CreateEventBusFuture{Future: future}
}

func (a *stub) CreatePartnerEventSource(ctx workflow.Context, input *eventbridge.CreatePartnerEventSourceInput) (*eventbridge.CreatePartnerEventSourceOutput, error) {
	var output eventbridge.CreatePartnerEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.CreatePartnerEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreatePartnerEventSourceAsync(ctx workflow.Context, input *eventbridge.CreatePartnerEventSourceInput) *CreatePartnerEventSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.CreatePartnerEventSource", input)
	return &CreatePartnerEventSourceFuture{Future: future}
}

func (a *stub) DeactivateEventSource(ctx workflow.Context, input *eventbridge.DeactivateEventSourceInput) (*eventbridge.DeactivateEventSourceOutput, error) {
	var output eventbridge.DeactivateEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.DeactivateEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeactivateEventSourceAsync(ctx workflow.Context, input *eventbridge.DeactivateEventSourceInput) *DeactivateEventSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.DeactivateEventSource", input)
	return &DeactivateEventSourceFuture{Future: future}
}

func (a *stub) DeleteEventBus(ctx workflow.Context, input *eventbridge.DeleteEventBusInput) (*eventbridge.DeleteEventBusOutput, error) {
	var output eventbridge.DeleteEventBusOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.DeleteEventBus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteEventBusAsync(ctx workflow.Context, input *eventbridge.DeleteEventBusInput) *DeleteEventBusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.DeleteEventBus", input)
	return &DeleteEventBusFuture{Future: future}
}

func (a *stub) DeletePartnerEventSource(ctx workflow.Context, input *eventbridge.DeletePartnerEventSourceInput) (*eventbridge.DeletePartnerEventSourceOutput, error) {
	var output eventbridge.DeletePartnerEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.DeletePartnerEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeletePartnerEventSourceAsync(ctx workflow.Context, input *eventbridge.DeletePartnerEventSourceInput) *DeletePartnerEventSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.DeletePartnerEventSource", input)
	return &DeletePartnerEventSourceFuture{Future: future}
}

func (a *stub) DeleteRule(ctx workflow.Context, input *eventbridge.DeleteRuleInput) (*eventbridge.DeleteRuleOutput, error) {
	var output eventbridge.DeleteRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.DeleteRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteRuleAsync(ctx workflow.Context, input *eventbridge.DeleteRuleInput) *DeleteRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.DeleteRule", input)
	return &DeleteRuleFuture{Future: future}
}

func (a *stub) DescribeEventBus(ctx workflow.Context, input *eventbridge.DescribeEventBusInput) (*eventbridge.DescribeEventBusOutput, error) {
	var output eventbridge.DescribeEventBusOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.DescribeEventBus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEventBusAsync(ctx workflow.Context, input *eventbridge.DescribeEventBusInput) *DescribeEventBusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.DescribeEventBus", input)
	return &DescribeEventBusFuture{Future: future}
}

func (a *stub) DescribeEventSource(ctx workflow.Context, input *eventbridge.DescribeEventSourceInput) (*eventbridge.DescribeEventSourceOutput, error) {
	var output eventbridge.DescribeEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.DescribeEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEventSourceAsync(ctx workflow.Context, input *eventbridge.DescribeEventSourceInput) *DescribeEventSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.DescribeEventSource", input)
	return &DescribeEventSourceFuture{Future: future}
}

func (a *stub) DescribePartnerEventSource(ctx workflow.Context, input *eventbridge.DescribePartnerEventSourceInput) (*eventbridge.DescribePartnerEventSourceOutput, error) {
	var output eventbridge.DescribePartnerEventSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.DescribePartnerEventSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribePartnerEventSourceAsync(ctx workflow.Context, input *eventbridge.DescribePartnerEventSourceInput) *DescribePartnerEventSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.DescribePartnerEventSource", input)
	return &DescribePartnerEventSourceFuture{Future: future}
}

func (a *stub) DescribeRule(ctx workflow.Context, input *eventbridge.DescribeRuleInput) (*eventbridge.DescribeRuleOutput, error) {
	var output eventbridge.DescribeRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.DescribeRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeRuleAsync(ctx workflow.Context, input *eventbridge.DescribeRuleInput) *DescribeRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.DescribeRule", input)
	return &DescribeRuleFuture{Future: future}
}

func (a *stub) DisableRule(ctx workflow.Context, input *eventbridge.DisableRuleInput) (*eventbridge.DisableRuleOutput, error) {
	var output eventbridge.DisableRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.DisableRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisableRuleAsync(ctx workflow.Context, input *eventbridge.DisableRuleInput) *DisableRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.DisableRule", input)
	return &DisableRuleFuture{Future: future}
}

func (a *stub) EnableRule(ctx workflow.Context, input *eventbridge.EnableRuleInput) (*eventbridge.EnableRuleOutput, error) {
	var output eventbridge.EnableRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.EnableRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) EnableRuleAsync(ctx workflow.Context, input *eventbridge.EnableRuleInput) *EnableRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.EnableRule", input)
	return &EnableRuleFuture{Future: future}
}

func (a *stub) ListEventBuses(ctx workflow.Context, input *eventbridge.ListEventBusesInput) (*eventbridge.ListEventBusesOutput, error) {
	var output eventbridge.ListEventBusesOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.ListEventBuses", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListEventBusesAsync(ctx workflow.Context, input *eventbridge.ListEventBusesInput) *ListEventBusesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.ListEventBuses", input)
	return &ListEventBusesFuture{Future: future}
}

func (a *stub) ListEventSources(ctx workflow.Context, input *eventbridge.ListEventSourcesInput) (*eventbridge.ListEventSourcesOutput, error) {
	var output eventbridge.ListEventSourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.ListEventSources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListEventSourcesAsync(ctx workflow.Context, input *eventbridge.ListEventSourcesInput) *ListEventSourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.ListEventSources", input)
	return &ListEventSourcesFuture{Future: future}
}

func (a *stub) ListPartnerEventSourceAccounts(ctx workflow.Context, input *eventbridge.ListPartnerEventSourceAccountsInput) (*eventbridge.ListPartnerEventSourceAccountsOutput, error) {
	var output eventbridge.ListPartnerEventSourceAccountsOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.ListPartnerEventSourceAccounts", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPartnerEventSourceAccountsAsync(ctx workflow.Context, input *eventbridge.ListPartnerEventSourceAccountsInput) *ListPartnerEventSourceAccountsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.ListPartnerEventSourceAccounts", input)
	return &ListPartnerEventSourceAccountsFuture{Future: future}
}

func (a *stub) ListPartnerEventSources(ctx workflow.Context, input *eventbridge.ListPartnerEventSourcesInput) (*eventbridge.ListPartnerEventSourcesOutput, error) {
	var output eventbridge.ListPartnerEventSourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.ListPartnerEventSources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPartnerEventSourcesAsync(ctx workflow.Context, input *eventbridge.ListPartnerEventSourcesInput) *ListPartnerEventSourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.ListPartnerEventSources", input)
	return &ListPartnerEventSourcesFuture{Future: future}
}

func (a *stub) ListRuleNamesByTarget(ctx workflow.Context, input *eventbridge.ListRuleNamesByTargetInput) (*eventbridge.ListRuleNamesByTargetOutput, error) {
	var output eventbridge.ListRuleNamesByTargetOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.ListRuleNamesByTarget", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListRuleNamesByTargetAsync(ctx workflow.Context, input *eventbridge.ListRuleNamesByTargetInput) *ListRuleNamesByTargetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.ListRuleNamesByTarget", input)
	return &ListRuleNamesByTargetFuture{Future: future}
}

func (a *stub) ListRules(ctx workflow.Context, input *eventbridge.ListRulesInput) (*eventbridge.ListRulesOutput, error) {
	var output eventbridge.ListRulesOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.ListRules", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListRulesAsync(ctx workflow.Context, input *eventbridge.ListRulesInput) *ListRulesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.ListRules", input)
	return &ListRulesFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *eventbridge.ListTagsForResourceInput) (*eventbridge.ListTagsForResourceOutput, error) {
	var output eventbridge.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *eventbridge.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) ListTargetsByRule(ctx workflow.Context, input *eventbridge.ListTargetsByRuleInput) (*eventbridge.ListTargetsByRuleOutput, error) {
	var output eventbridge.ListTargetsByRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.ListTargetsByRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTargetsByRuleAsync(ctx workflow.Context, input *eventbridge.ListTargetsByRuleInput) *ListTargetsByRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.ListTargetsByRule", input)
	return &ListTargetsByRuleFuture{Future: future}
}

func (a *stub) PutEvents(ctx workflow.Context, input *eventbridge.PutEventsInput) (*eventbridge.PutEventsOutput, error) {
	var output eventbridge.PutEventsOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.PutEvents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutEventsAsync(ctx workflow.Context, input *eventbridge.PutEventsInput) *PutEventsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.PutEvents", input)
	return &PutEventsFuture{Future: future}
}

func (a *stub) PutPartnerEvents(ctx workflow.Context, input *eventbridge.PutPartnerEventsInput) (*eventbridge.PutPartnerEventsOutput, error) {
	var output eventbridge.PutPartnerEventsOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.PutPartnerEvents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutPartnerEventsAsync(ctx workflow.Context, input *eventbridge.PutPartnerEventsInput) *PutPartnerEventsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.PutPartnerEvents", input)
	return &PutPartnerEventsFuture{Future: future}
}

func (a *stub) PutPermission(ctx workflow.Context, input *eventbridge.PutPermissionInput) (*eventbridge.PutPermissionOutput, error) {
	var output eventbridge.PutPermissionOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.PutPermission", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutPermissionAsync(ctx workflow.Context, input *eventbridge.PutPermissionInput) *PutPermissionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.PutPermission", input)
	return &PutPermissionFuture{Future: future}
}

func (a *stub) PutRule(ctx workflow.Context, input *eventbridge.PutRuleInput) (*eventbridge.PutRuleOutput, error) {
	var output eventbridge.PutRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.PutRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutRuleAsync(ctx workflow.Context, input *eventbridge.PutRuleInput) *PutRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.PutRule", input)
	return &PutRuleFuture{Future: future}
}

func (a *stub) PutTargets(ctx workflow.Context, input *eventbridge.PutTargetsInput) (*eventbridge.PutTargetsOutput, error) {
	var output eventbridge.PutTargetsOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.PutTargets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutTargetsAsync(ctx workflow.Context, input *eventbridge.PutTargetsInput) *PutTargetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.PutTargets", input)
	return &PutTargetsFuture{Future: future}
}

func (a *stub) RemovePermission(ctx workflow.Context, input *eventbridge.RemovePermissionInput) (*eventbridge.RemovePermissionOutput, error) {
	var output eventbridge.RemovePermissionOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.RemovePermission", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemovePermissionAsync(ctx workflow.Context, input *eventbridge.RemovePermissionInput) *RemovePermissionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.RemovePermission", input)
	return &RemovePermissionFuture{Future: future}
}

func (a *stub) RemoveTargets(ctx workflow.Context, input *eventbridge.RemoveTargetsInput) (*eventbridge.RemoveTargetsOutput, error) {
	var output eventbridge.RemoveTargetsOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.RemoveTargets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveTargetsAsync(ctx workflow.Context, input *eventbridge.RemoveTargetsInput) *RemoveTargetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.RemoveTargets", input)
	return &RemoveTargetsFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *eventbridge.TagResourceInput) (*eventbridge.TagResourceOutput, error) {
	var output eventbridge.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *eventbridge.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) TestEventPattern(ctx workflow.Context, input *eventbridge.TestEventPatternInput) (*eventbridge.TestEventPatternOutput, error) {
	var output eventbridge.TestEventPatternOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.TestEventPattern", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TestEventPatternAsync(ctx workflow.Context, input *eventbridge.TestEventPatternInput) *TestEventPatternFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.TestEventPattern", input)
	return &TestEventPatternFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *eventbridge.UntagResourceInput) (*eventbridge.UntagResourceOutput, error) {
	var output eventbridge.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.eventbridge.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *eventbridge.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.eventbridge.UntagResource", input)
	return &UntagResourceFuture{Future: future}
}
