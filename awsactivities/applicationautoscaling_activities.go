package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go/service/applicationautoscaling/applicationautoscalingiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type ApplicationAutoScalingActivities struct {
	client applicationautoscalingiface.ApplicationAutoScalingAPI
}

func NewApplicationAutoScalingActivities(session *session.Session, config ...*aws.Config) *ApplicationAutoScalingActivities {
	client := applicationautoscaling.New(session, config...)
	return &ApplicationAutoScalingActivities{client: client}
}

func (a *ApplicationAutoScalingActivities) DeleteScalingPolicy(ctx context.Context, input *applicationautoscaling.DeleteScalingPolicyInput) (*applicationautoscaling.DeleteScalingPolicyOutput, error) {
	return a.client.DeleteScalingPolicyWithContext(ctx, input)
}

func (a *ApplicationAutoScalingActivities) DeleteScheduledAction(ctx context.Context, input *applicationautoscaling.DeleteScheduledActionInput) (*applicationautoscaling.DeleteScheduledActionOutput, error) {
	return a.client.DeleteScheduledActionWithContext(ctx, input)
}

func (a *ApplicationAutoScalingActivities) DeregisterScalableTarget(ctx context.Context, input *applicationautoscaling.DeregisterScalableTargetInput) (*applicationautoscaling.DeregisterScalableTargetOutput, error) {
	return a.client.DeregisterScalableTargetWithContext(ctx, input)
}

func (a *ApplicationAutoScalingActivities) DescribeScalableTargets(ctx context.Context, input *applicationautoscaling.DescribeScalableTargetsInput) (*applicationautoscaling.DescribeScalableTargetsOutput, error) {
	return a.client.DescribeScalableTargetsWithContext(ctx, input)
}

func (a *ApplicationAutoScalingActivities) DescribeScalingActivities(ctx context.Context, input *applicationautoscaling.DescribeScalingActivitiesInput) (*applicationautoscaling.DescribeScalingActivitiesOutput, error) {
	return a.client.DescribeScalingActivitiesWithContext(ctx, input)
}

func (a *ApplicationAutoScalingActivities) DescribeScalingPolicies(ctx context.Context, input *applicationautoscaling.DescribeScalingPoliciesInput) (*applicationautoscaling.DescribeScalingPoliciesOutput, error) {
	return a.client.DescribeScalingPoliciesWithContext(ctx, input)
}

func (a *ApplicationAutoScalingActivities) DescribeScheduledActions(ctx context.Context, input *applicationautoscaling.DescribeScheduledActionsInput) (*applicationautoscaling.DescribeScheduledActionsOutput, error) {
	return a.client.DescribeScheduledActionsWithContext(ctx, input)
}

func (a *ApplicationAutoScalingActivities) PutScalingPolicy(ctx context.Context, input *applicationautoscaling.PutScalingPolicyInput) (*applicationautoscaling.PutScalingPolicyOutput, error) {
	return a.client.PutScalingPolicyWithContext(ctx, input)
}

func (a *ApplicationAutoScalingActivities) PutScheduledAction(ctx context.Context, input *applicationautoscaling.PutScheduledActionInput) (*applicationautoscaling.PutScheduledActionOutput, error) {
	return a.client.PutScheduledActionWithContext(ctx, input)
}

func (a *ApplicationAutoScalingActivities) RegisterScalableTarget(ctx context.Context, input *applicationautoscaling.RegisterScalableTargetInput) (*applicationautoscaling.RegisterScalableTargetOutput, error) {
	return a.client.RegisterScalableTargetWithContext(ctx, input)
}
