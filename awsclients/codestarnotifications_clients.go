package awsclients

import (
	"github.com/aws/aws-sdk-go/service/codestarnotifications"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type CodeStarNotificationsClient interface {
       CreateNotificationRule(ctx workflow.Context, input *codestarnotifications.CreateNotificationRuleInput) (*codestarnotifications.CreateNotificationRuleOutput, error)
       CreateNotificationRuleAsync(ctx workflow.Context, input *codestarnotifications.CreateNotificationRuleInput) *CodestarnotificationsCreateNotificationRuleResult

       DeleteNotificationRule(ctx workflow.Context, input *codestarnotifications.DeleteNotificationRuleInput) (*codestarnotifications.DeleteNotificationRuleOutput, error)
       DeleteNotificationRuleAsync(ctx workflow.Context, input *codestarnotifications.DeleteNotificationRuleInput) *CodestarnotificationsDeleteNotificationRuleResult

       DeleteTarget(ctx workflow.Context, input *codestarnotifications.DeleteTargetInput) (*codestarnotifications.DeleteTargetOutput, error)
       DeleteTargetAsync(ctx workflow.Context, input *codestarnotifications.DeleteTargetInput) *CodestarnotificationsDeleteTargetResult

       DescribeNotificationRule(ctx workflow.Context, input *codestarnotifications.DescribeNotificationRuleInput) (*codestarnotifications.DescribeNotificationRuleOutput, error)
       DescribeNotificationRuleAsync(ctx workflow.Context, input *codestarnotifications.DescribeNotificationRuleInput) *CodestarnotificationsDescribeNotificationRuleResult

       ListEventTypes(ctx workflow.Context, input *codestarnotifications.ListEventTypesInput) (*codestarnotifications.ListEventTypesOutput, error)
       ListEventTypesAsync(ctx workflow.Context, input *codestarnotifications.ListEventTypesInput) *CodestarnotificationsListEventTypesResult

       ListNotificationRules(ctx workflow.Context, input *codestarnotifications.ListNotificationRulesInput) (*codestarnotifications.ListNotificationRulesOutput, error)
       ListNotificationRulesAsync(ctx workflow.Context, input *codestarnotifications.ListNotificationRulesInput) *CodestarnotificationsListNotificationRulesResult

       ListTagsForResource(ctx workflow.Context, input *codestarnotifications.ListTagsForResourceInput) (*codestarnotifications.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *codestarnotifications.ListTagsForResourceInput) *CodestarnotificationsListTagsForResourceResult

       ListTargets(ctx workflow.Context, input *codestarnotifications.ListTargetsInput) (*codestarnotifications.ListTargetsOutput, error)
       ListTargetsAsync(ctx workflow.Context, input *codestarnotifications.ListTargetsInput) *CodestarnotificationsListTargetsResult

       Subscribe(ctx workflow.Context, input *codestarnotifications.SubscribeInput) (*codestarnotifications.SubscribeOutput, error)
       SubscribeAsync(ctx workflow.Context, input *codestarnotifications.SubscribeInput) *CodestarnotificationsSubscribeResult

       TagResource(ctx workflow.Context, input *codestarnotifications.TagResourceInput) (*codestarnotifications.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *codestarnotifications.TagResourceInput) *CodestarnotificationsTagResourceResult

       Unsubscribe(ctx workflow.Context, input *codestarnotifications.UnsubscribeInput) (*codestarnotifications.UnsubscribeOutput, error)
       UnsubscribeAsync(ctx workflow.Context, input *codestarnotifications.UnsubscribeInput) *CodestarnotificationsUnsubscribeResult

       UntagResource(ctx workflow.Context, input *codestarnotifications.UntagResourceInput) (*codestarnotifications.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *codestarnotifications.UntagResourceInput) *CodestarnotificationsUntagResourceResult

       UpdateNotificationRule(ctx workflow.Context, input *codestarnotifications.UpdateNotificationRuleInput) (*codestarnotifications.UpdateNotificationRuleOutput, error)
       UpdateNotificationRuleAsync(ctx workflow.Context, input *codestarnotifications.UpdateNotificationRuleInput) *CodestarnotificationsUpdateNotificationRuleResult
}

type CodestarnotificationsCreateNotificationRuleResult struct {
	Result workflow.Future
}

func (r *CodestarnotificationsCreateNotificationRuleResult) Get(ctx workflow.Context) (*codestarnotifications.CreateNotificationRuleOutput, error) {
    var output codestarnotifications.CreateNotificationRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarnotificationsDeleteNotificationRuleResult struct {
	Result workflow.Future
}

func (r *CodestarnotificationsDeleteNotificationRuleResult) Get(ctx workflow.Context) (*codestarnotifications.DeleteNotificationRuleOutput, error) {
    var output codestarnotifications.DeleteNotificationRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarnotificationsDeleteTargetResult struct {
	Result workflow.Future
}

func (r *CodestarnotificationsDeleteTargetResult) Get(ctx workflow.Context) (*codestarnotifications.DeleteTargetOutput, error) {
    var output codestarnotifications.DeleteTargetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarnotificationsDescribeNotificationRuleResult struct {
	Result workflow.Future
}

func (r *CodestarnotificationsDescribeNotificationRuleResult) Get(ctx workflow.Context) (*codestarnotifications.DescribeNotificationRuleOutput, error) {
    var output codestarnotifications.DescribeNotificationRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarnotificationsListEventTypesResult struct {
	Result workflow.Future
}

func (r *CodestarnotificationsListEventTypesResult) Get(ctx workflow.Context) (*codestarnotifications.ListEventTypesOutput, error) {
    var output codestarnotifications.ListEventTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarnotificationsListNotificationRulesResult struct {
	Result workflow.Future
}

func (r *CodestarnotificationsListNotificationRulesResult) Get(ctx workflow.Context) (*codestarnotifications.ListNotificationRulesOutput, error) {
    var output codestarnotifications.ListNotificationRulesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarnotificationsListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *CodestarnotificationsListTagsForResourceResult) Get(ctx workflow.Context) (*codestarnotifications.ListTagsForResourceOutput, error) {
    var output codestarnotifications.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarnotificationsListTargetsResult struct {
	Result workflow.Future
}

func (r *CodestarnotificationsListTargetsResult) Get(ctx workflow.Context) (*codestarnotifications.ListTargetsOutput, error) {
    var output codestarnotifications.ListTargetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarnotificationsSubscribeResult struct {
	Result workflow.Future
}

func (r *CodestarnotificationsSubscribeResult) Get(ctx workflow.Context) (*codestarnotifications.SubscribeOutput, error) {
    var output codestarnotifications.SubscribeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarnotificationsTagResourceResult struct {
	Result workflow.Future
}

func (r *CodestarnotificationsTagResourceResult) Get(ctx workflow.Context) (*codestarnotifications.TagResourceOutput, error) {
    var output codestarnotifications.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarnotificationsUnsubscribeResult struct {
	Result workflow.Future
}

func (r *CodestarnotificationsUnsubscribeResult) Get(ctx workflow.Context) (*codestarnotifications.UnsubscribeOutput, error) {
    var output codestarnotifications.UnsubscribeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarnotificationsUntagResourceResult struct {
	Result workflow.Future
}

func (r *CodestarnotificationsUntagResourceResult) Get(ctx workflow.Context) (*codestarnotifications.UntagResourceOutput, error) {
    var output codestarnotifications.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarnotificationsUpdateNotificationRuleResult struct {
	Result workflow.Future
}

func (r *CodestarnotificationsUpdateNotificationRuleResult) Get(ctx workflow.Context) (*codestarnotifications.UpdateNotificationRuleOutput, error) {
    var output codestarnotifications.UpdateNotificationRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeStarNotificationsStub struct {
    activities awsactivities.CodeStarNotificationsActivities
}

func NewCodeStarNotificationsStub() CodeStarNotificationsClient {
    return &CodeStarNotificationsStub{}
}

func (a *CodeStarNotificationsStub) CreateNotificationRule(ctx workflow.Context, input *codestarnotifications.CreateNotificationRuleInput) (*codestarnotifications.CreateNotificationRuleOutput, error) {
    var output codestarnotifications.CreateNotificationRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateNotificationRule, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarNotificationsStub) CreateNotificationRuleAsync(ctx workflow.Context, input *codestarnotifications.CreateNotificationRuleInput) *CodestarnotificationsCreateNotificationRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateNotificationRule, input)
    return &CodestarnotificationsCreateNotificationRuleResult{Result: future}
}

func (a *CodeStarNotificationsStub) DeleteNotificationRule(ctx workflow.Context, input *codestarnotifications.DeleteNotificationRuleInput) (*codestarnotifications.DeleteNotificationRuleOutput, error) {
    var output codestarnotifications.DeleteNotificationRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteNotificationRule, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarNotificationsStub) DeleteNotificationRuleAsync(ctx workflow.Context, input *codestarnotifications.DeleteNotificationRuleInput) *CodestarnotificationsDeleteNotificationRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteNotificationRule, input)
    return &CodestarnotificationsDeleteNotificationRuleResult{Result: future}
}

func (a *CodeStarNotificationsStub) DeleteTarget(ctx workflow.Context, input *codestarnotifications.DeleteTargetInput) (*codestarnotifications.DeleteTargetOutput, error) {
    var output codestarnotifications.DeleteTargetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTarget, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarNotificationsStub) DeleteTargetAsync(ctx workflow.Context, input *codestarnotifications.DeleteTargetInput) *CodestarnotificationsDeleteTargetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTarget, input)
    return &CodestarnotificationsDeleteTargetResult{Result: future}
}

func (a *CodeStarNotificationsStub) DescribeNotificationRule(ctx workflow.Context, input *codestarnotifications.DescribeNotificationRuleInput) (*codestarnotifications.DescribeNotificationRuleOutput, error) {
    var output codestarnotifications.DescribeNotificationRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeNotificationRule, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarNotificationsStub) DescribeNotificationRuleAsync(ctx workflow.Context, input *codestarnotifications.DescribeNotificationRuleInput) *CodestarnotificationsDescribeNotificationRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeNotificationRule, input)
    return &CodestarnotificationsDescribeNotificationRuleResult{Result: future}
}

func (a *CodeStarNotificationsStub) ListEventTypes(ctx workflow.Context, input *codestarnotifications.ListEventTypesInput) (*codestarnotifications.ListEventTypesOutput, error) {
    var output codestarnotifications.ListEventTypesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListEventTypes, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarNotificationsStub) ListEventTypesAsync(ctx workflow.Context, input *codestarnotifications.ListEventTypesInput) *CodestarnotificationsListEventTypesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListEventTypes, input)
    return &CodestarnotificationsListEventTypesResult{Result: future}
}

func (a *CodeStarNotificationsStub) ListNotificationRules(ctx workflow.Context, input *codestarnotifications.ListNotificationRulesInput) (*codestarnotifications.ListNotificationRulesOutput, error) {
    var output codestarnotifications.ListNotificationRulesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListNotificationRules, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarNotificationsStub) ListNotificationRulesAsync(ctx workflow.Context, input *codestarnotifications.ListNotificationRulesInput) *CodestarnotificationsListNotificationRulesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListNotificationRules, input)
    return &CodestarnotificationsListNotificationRulesResult{Result: future}
}

func (a *CodeStarNotificationsStub) ListTagsForResource(ctx workflow.Context, input *codestarnotifications.ListTagsForResourceInput) (*codestarnotifications.ListTagsForResourceOutput, error) {
    var output codestarnotifications.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarNotificationsStub) ListTagsForResourceAsync(ctx workflow.Context, input *codestarnotifications.ListTagsForResourceInput) *CodestarnotificationsListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &CodestarnotificationsListTagsForResourceResult{Result: future}
}

func (a *CodeStarNotificationsStub) ListTargets(ctx workflow.Context, input *codestarnotifications.ListTargetsInput) (*codestarnotifications.ListTargetsOutput, error) {
    var output codestarnotifications.ListTargetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTargets, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarNotificationsStub) ListTargetsAsync(ctx workflow.Context, input *codestarnotifications.ListTargetsInput) *CodestarnotificationsListTargetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTargets, input)
    return &CodestarnotificationsListTargetsResult{Result: future}
}

func (a *CodeStarNotificationsStub) Subscribe(ctx workflow.Context, input *codestarnotifications.SubscribeInput) (*codestarnotifications.SubscribeOutput, error) {
    var output codestarnotifications.SubscribeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.Subscribe, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarNotificationsStub) SubscribeAsync(ctx workflow.Context, input *codestarnotifications.SubscribeInput) *CodestarnotificationsSubscribeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.Subscribe, input)
    return &CodestarnotificationsSubscribeResult{Result: future}
}

func (a *CodeStarNotificationsStub) TagResource(ctx workflow.Context, input *codestarnotifications.TagResourceInput) (*codestarnotifications.TagResourceOutput, error) {
    var output codestarnotifications.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarNotificationsStub) TagResourceAsync(ctx workflow.Context, input *codestarnotifications.TagResourceInput) *CodestarnotificationsTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &CodestarnotificationsTagResourceResult{Result: future}
}

func (a *CodeStarNotificationsStub) Unsubscribe(ctx workflow.Context, input *codestarnotifications.UnsubscribeInput) (*codestarnotifications.UnsubscribeOutput, error) {
    var output codestarnotifications.UnsubscribeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.Unsubscribe, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarNotificationsStub) UnsubscribeAsync(ctx workflow.Context, input *codestarnotifications.UnsubscribeInput) *CodestarnotificationsUnsubscribeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.Unsubscribe, input)
    return &CodestarnotificationsUnsubscribeResult{Result: future}
}

func (a *CodeStarNotificationsStub) UntagResource(ctx workflow.Context, input *codestarnotifications.UntagResourceInput) (*codestarnotifications.UntagResourceOutput, error) {
    var output codestarnotifications.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarNotificationsStub) UntagResourceAsync(ctx workflow.Context, input *codestarnotifications.UntagResourceInput) *CodestarnotificationsUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &CodestarnotificationsUntagResourceResult{Result: future}
}

func (a *CodeStarNotificationsStub) UpdateNotificationRule(ctx workflow.Context, input *codestarnotifications.UpdateNotificationRuleInput) (*codestarnotifications.UpdateNotificationRuleOutput, error) {
    var output codestarnotifications.UpdateNotificationRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateNotificationRule, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarNotificationsStub) UpdateNotificationRuleAsync(ctx workflow.Context, input *codestarnotifications.UpdateNotificationRuleInput) *CodestarnotificationsUpdateNotificationRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateNotificationRule, input)
    return &CodestarnotificationsUpdateNotificationRuleResult{Result: future}
}
