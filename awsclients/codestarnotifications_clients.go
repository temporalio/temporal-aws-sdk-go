// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/codestarnotifications"
	"go.temporal.io/sdk/workflow"
)

type CodeStarNotificationsClient interface {
	CreateNotificationRule(ctx workflow.Context, input *codestarnotifications.CreateNotificationRuleInput) (*codestarnotifications.CreateNotificationRuleOutput, error)
	CreateNotificationRuleAsync(ctx workflow.Context, input *codestarnotifications.CreateNotificationRuleInput) *CodestarnotificationsCreateNotificationRuleFuture

	DeleteNotificationRule(ctx workflow.Context, input *codestarnotifications.DeleteNotificationRuleInput) (*codestarnotifications.DeleteNotificationRuleOutput, error)
	DeleteNotificationRuleAsync(ctx workflow.Context, input *codestarnotifications.DeleteNotificationRuleInput) *CodestarnotificationsDeleteNotificationRuleFuture

	DeleteTarget(ctx workflow.Context, input *codestarnotifications.DeleteTargetInput) (*codestarnotifications.DeleteTargetOutput, error)
	DeleteTargetAsync(ctx workflow.Context, input *codestarnotifications.DeleteTargetInput) *CodestarnotificationsDeleteTargetFuture

	DescribeNotificationRule(ctx workflow.Context, input *codestarnotifications.DescribeNotificationRuleInput) (*codestarnotifications.DescribeNotificationRuleOutput, error)
	DescribeNotificationRuleAsync(ctx workflow.Context, input *codestarnotifications.DescribeNotificationRuleInput) *CodestarnotificationsDescribeNotificationRuleFuture

	ListEventTypes(ctx workflow.Context, input *codestarnotifications.ListEventTypesInput) (*codestarnotifications.ListEventTypesOutput, error)
	ListEventTypesAsync(ctx workflow.Context, input *codestarnotifications.ListEventTypesInput) *CodestarnotificationsListEventTypesFuture

	ListNotificationRules(ctx workflow.Context, input *codestarnotifications.ListNotificationRulesInput) (*codestarnotifications.ListNotificationRulesOutput, error)
	ListNotificationRulesAsync(ctx workflow.Context, input *codestarnotifications.ListNotificationRulesInput) *CodestarnotificationsListNotificationRulesFuture

	ListTagsForResource(ctx workflow.Context, input *codestarnotifications.ListTagsForResourceInput) (*codestarnotifications.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *codestarnotifications.ListTagsForResourceInput) *CodestarnotificationsListTagsForResourceFuture

	ListTargets(ctx workflow.Context, input *codestarnotifications.ListTargetsInput) (*codestarnotifications.ListTargetsOutput, error)
	ListTargetsAsync(ctx workflow.Context, input *codestarnotifications.ListTargetsInput) *CodestarnotificationsListTargetsFuture

	Subscribe(ctx workflow.Context, input *codestarnotifications.SubscribeInput) (*codestarnotifications.SubscribeOutput, error)
	SubscribeAsync(ctx workflow.Context, input *codestarnotifications.SubscribeInput) *CodestarnotificationsSubscribeFuture

	TagResource(ctx workflow.Context, input *codestarnotifications.TagResourceInput) (*codestarnotifications.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *codestarnotifications.TagResourceInput) *CodestarnotificationsTagResourceFuture

	Unsubscribe(ctx workflow.Context, input *codestarnotifications.UnsubscribeInput) (*codestarnotifications.UnsubscribeOutput, error)
	UnsubscribeAsync(ctx workflow.Context, input *codestarnotifications.UnsubscribeInput) *CodestarnotificationsUnsubscribeFuture

	UntagResource(ctx workflow.Context, input *codestarnotifications.UntagResourceInput) (*codestarnotifications.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *codestarnotifications.UntagResourceInput) *CodestarnotificationsUntagResourceFuture

	UpdateNotificationRule(ctx workflow.Context, input *codestarnotifications.UpdateNotificationRuleInput) (*codestarnotifications.UpdateNotificationRuleOutput, error)
	UpdateNotificationRuleAsync(ctx workflow.Context, input *codestarnotifications.UpdateNotificationRuleInput) *CodestarnotificationsUpdateNotificationRuleFuture
}

type CodeStarNotificationsStub struct{}

func NewCodeStarNotificationsStub() CodeStarNotificationsClient {
	return &CodeStarNotificationsStub{}
}

type CodestarnotificationsCreateNotificationRuleFuture struct {
	Future workflow.Future
}

func (r *CodestarnotificationsCreateNotificationRuleFuture) Get(ctx workflow.Context) (*codestarnotifications.CreateNotificationRuleOutput, error) {
	var output codestarnotifications.CreateNotificationRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodestarnotificationsDeleteNotificationRuleFuture struct {
	Future workflow.Future
}

func (r *CodestarnotificationsDeleteNotificationRuleFuture) Get(ctx workflow.Context) (*codestarnotifications.DeleteNotificationRuleOutput, error) {
	var output codestarnotifications.DeleteNotificationRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodestarnotificationsDeleteTargetFuture struct {
	Future workflow.Future
}

func (r *CodestarnotificationsDeleteTargetFuture) Get(ctx workflow.Context) (*codestarnotifications.DeleteTargetOutput, error) {
	var output codestarnotifications.DeleteTargetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodestarnotificationsDescribeNotificationRuleFuture struct {
	Future workflow.Future
}

func (r *CodestarnotificationsDescribeNotificationRuleFuture) Get(ctx workflow.Context) (*codestarnotifications.DescribeNotificationRuleOutput, error) {
	var output codestarnotifications.DescribeNotificationRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodestarnotificationsListEventTypesFuture struct {
	Future workflow.Future
}

func (r *CodestarnotificationsListEventTypesFuture) Get(ctx workflow.Context) (*codestarnotifications.ListEventTypesOutput, error) {
	var output codestarnotifications.ListEventTypesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodestarnotificationsListNotificationRulesFuture struct {
	Future workflow.Future
}

func (r *CodestarnotificationsListNotificationRulesFuture) Get(ctx workflow.Context) (*codestarnotifications.ListNotificationRulesOutput, error) {
	var output codestarnotifications.ListNotificationRulesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodestarnotificationsListTagsForResourceFuture struct {
	Future workflow.Future
}

func (r *CodestarnotificationsListTagsForResourceFuture) Get(ctx workflow.Context) (*codestarnotifications.ListTagsForResourceOutput, error) {
	var output codestarnotifications.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodestarnotificationsListTargetsFuture struct {
	Future workflow.Future
}

func (r *CodestarnotificationsListTargetsFuture) Get(ctx workflow.Context) (*codestarnotifications.ListTargetsOutput, error) {
	var output codestarnotifications.ListTargetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodestarnotificationsSubscribeFuture struct {
	Future workflow.Future
}

func (r *CodestarnotificationsSubscribeFuture) Get(ctx workflow.Context) (*codestarnotifications.SubscribeOutput, error) {
	var output codestarnotifications.SubscribeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodestarnotificationsTagResourceFuture struct {
	Future workflow.Future
}

func (r *CodestarnotificationsTagResourceFuture) Get(ctx workflow.Context) (*codestarnotifications.TagResourceOutput, error) {
	var output codestarnotifications.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodestarnotificationsUnsubscribeFuture struct {
	Future workflow.Future
}

func (r *CodestarnotificationsUnsubscribeFuture) Get(ctx workflow.Context) (*codestarnotifications.UnsubscribeOutput, error) {
	var output codestarnotifications.UnsubscribeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodestarnotificationsUntagResourceFuture struct {
	Future workflow.Future
}

func (r *CodestarnotificationsUntagResourceFuture) Get(ctx workflow.Context) (*codestarnotifications.UntagResourceOutput, error) {
	var output codestarnotifications.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodestarnotificationsUpdateNotificationRuleFuture struct {
	Future workflow.Future
}

func (r *CodestarnotificationsUpdateNotificationRuleFuture) Get(ctx workflow.Context) (*codestarnotifications.UpdateNotificationRuleOutput, error) {
	var output codestarnotifications.UpdateNotificationRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *CodeStarNotificationsStub) CreateNotificationRule(ctx workflow.Context, input *codestarnotifications.CreateNotificationRuleInput) (*codestarnotifications.CreateNotificationRuleOutput, error) {
	var output codestarnotifications.CreateNotificationRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.CreateNotificationRule", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeStarNotificationsStub) CreateNotificationRuleAsync(ctx workflow.Context, input *codestarnotifications.CreateNotificationRuleInput) *CodestarnotificationsCreateNotificationRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.CreateNotificationRule", input)
	return &CodestarnotificationsCreateNotificationRuleFuture{Future: future}
}

func (a *CodeStarNotificationsStub) DeleteNotificationRule(ctx workflow.Context, input *codestarnotifications.DeleteNotificationRuleInput) (*codestarnotifications.DeleteNotificationRuleOutput, error) {
	var output codestarnotifications.DeleteNotificationRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.DeleteNotificationRule", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeStarNotificationsStub) DeleteNotificationRuleAsync(ctx workflow.Context, input *codestarnotifications.DeleteNotificationRuleInput) *CodestarnotificationsDeleteNotificationRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.DeleteNotificationRule", input)
	return &CodestarnotificationsDeleteNotificationRuleFuture{Future: future}
}

func (a *CodeStarNotificationsStub) DeleteTarget(ctx workflow.Context, input *codestarnotifications.DeleteTargetInput) (*codestarnotifications.DeleteTargetOutput, error) {
	var output codestarnotifications.DeleteTargetOutput
	err := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.DeleteTarget", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeStarNotificationsStub) DeleteTargetAsync(ctx workflow.Context, input *codestarnotifications.DeleteTargetInput) *CodestarnotificationsDeleteTargetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.DeleteTarget", input)
	return &CodestarnotificationsDeleteTargetFuture{Future: future}
}

func (a *CodeStarNotificationsStub) DescribeNotificationRule(ctx workflow.Context, input *codestarnotifications.DescribeNotificationRuleInput) (*codestarnotifications.DescribeNotificationRuleOutput, error) {
	var output codestarnotifications.DescribeNotificationRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.DescribeNotificationRule", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeStarNotificationsStub) DescribeNotificationRuleAsync(ctx workflow.Context, input *codestarnotifications.DescribeNotificationRuleInput) *CodestarnotificationsDescribeNotificationRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.DescribeNotificationRule", input)
	return &CodestarnotificationsDescribeNotificationRuleFuture{Future: future}
}

func (a *CodeStarNotificationsStub) ListEventTypes(ctx workflow.Context, input *codestarnotifications.ListEventTypesInput) (*codestarnotifications.ListEventTypesOutput, error) {
	var output codestarnotifications.ListEventTypesOutput
	err := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.ListEventTypes", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeStarNotificationsStub) ListEventTypesAsync(ctx workflow.Context, input *codestarnotifications.ListEventTypesInput) *CodestarnotificationsListEventTypesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.ListEventTypes", input)
	return &CodestarnotificationsListEventTypesFuture{Future: future}
}

func (a *CodeStarNotificationsStub) ListNotificationRules(ctx workflow.Context, input *codestarnotifications.ListNotificationRulesInput) (*codestarnotifications.ListNotificationRulesOutput, error) {
	var output codestarnotifications.ListNotificationRulesOutput
	err := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.ListNotificationRules", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeStarNotificationsStub) ListNotificationRulesAsync(ctx workflow.Context, input *codestarnotifications.ListNotificationRulesInput) *CodestarnotificationsListNotificationRulesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.ListNotificationRules", input)
	return &CodestarnotificationsListNotificationRulesFuture{Future: future}
}

func (a *CodeStarNotificationsStub) ListTagsForResource(ctx workflow.Context, input *codestarnotifications.ListTagsForResourceInput) (*codestarnotifications.ListTagsForResourceOutput, error) {
	var output codestarnotifications.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeStarNotificationsStub) ListTagsForResourceAsync(ctx workflow.Context, input *codestarnotifications.ListTagsForResourceInput) *CodestarnotificationsListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.ListTagsForResource", input)
	return &CodestarnotificationsListTagsForResourceFuture{Future: future}
}

func (a *CodeStarNotificationsStub) ListTargets(ctx workflow.Context, input *codestarnotifications.ListTargetsInput) (*codestarnotifications.ListTargetsOutput, error) {
	var output codestarnotifications.ListTargetsOutput
	err := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.ListTargets", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeStarNotificationsStub) ListTargetsAsync(ctx workflow.Context, input *codestarnotifications.ListTargetsInput) *CodestarnotificationsListTargetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.ListTargets", input)
	return &CodestarnotificationsListTargetsFuture{Future: future}
}

func (a *CodeStarNotificationsStub) Subscribe(ctx workflow.Context, input *codestarnotifications.SubscribeInput) (*codestarnotifications.SubscribeOutput, error) {
	var output codestarnotifications.SubscribeOutput
	err := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.Subscribe", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeStarNotificationsStub) SubscribeAsync(ctx workflow.Context, input *codestarnotifications.SubscribeInput) *CodestarnotificationsSubscribeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.Subscribe", input)
	return &CodestarnotificationsSubscribeFuture{Future: future}
}

func (a *CodeStarNotificationsStub) TagResource(ctx workflow.Context, input *codestarnotifications.TagResourceInput) (*codestarnotifications.TagResourceOutput, error) {
	var output codestarnotifications.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeStarNotificationsStub) TagResourceAsync(ctx workflow.Context, input *codestarnotifications.TagResourceInput) *CodestarnotificationsTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.TagResource", input)
	return &CodestarnotificationsTagResourceFuture{Future: future}
}

func (a *CodeStarNotificationsStub) Unsubscribe(ctx workflow.Context, input *codestarnotifications.UnsubscribeInput) (*codestarnotifications.UnsubscribeOutput, error) {
	var output codestarnotifications.UnsubscribeOutput
	err := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.Unsubscribe", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeStarNotificationsStub) UnsubscribeAsync(ctx workflow.Context, input *codestarnotifications.UnsubscribeInput) *CodestarnotificationsUnsubscribeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.Unsubscribe", input)
	return &CodestarnotificationsUnsubscribeFuture{Future: future}
}

func (a *CodeStarNotificationsStub) UntagResource(ctx workflow.Context, input *codestarnotifications.UntagResourceInput) (*codestarnotifications.UntagResourceOutput, error) {
	var output codestarnotifications.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeStarNotificationsStub) UntagResourceAsync(ctx workflow.Context, input *codestarnotifications.UntagResourceInput) *CodestarnotificationsUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.UntagResource", input)
	return &CodestarnotificationsUntagResourceFuture{Future: future}
}

func (a *CodeStarNotificationsStub) UpdateNotificationRule(ctx workflow.Context, input *codestarnotifications.UpdateNotificationRuleInput) (*codestarnotifications.UpdateNotificationRuleOutput, error) {
	var output codestarnotifications.UpdateNotificationRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.UpdateNotificationRule", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeStarNotificationsStub) UpdateNotificationRuleAsync(ctx workflow.Context, input *codestarnotifications.UpdateNotificationRuleInput) *CodestarnotificationsUpdateNotificationRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codestarnotifications.UpdateNotificationRule", input)
	return &CodestarnotificationsUpdateNotificationRuleFuture{Future: future}
}
