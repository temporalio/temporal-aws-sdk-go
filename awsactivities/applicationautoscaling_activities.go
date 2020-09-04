package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go/service/applicationautoscaling/applicationautoscalingiface"
)

type ApplicationAutoScalingActivities struct {
	client applicationautoscalingiface.ApplicationAutoScalingAPI
}

func NewApplicationAutoScalingActivities(client applicationautoscalingiface.ApplicationAutoScalingAPI) *ApplicationAutoScalingActivities {
    return &ApplicationAutoScalingActivities{client: client}
}


func (a *ApplicationAutoScalingActivities) DeleteScalingPolicy(input *applicationautoscaling.DeleteScalingPolicyInput) (*applicationautoscaling.DeleteScalingPolicyOutput, error) {
    return a.client.DeleteScalingPolicy(input)
}



func (a *ApplicationAutoScalingActivities) DeleteScheduledAction(input *applicationautoscaling.DeleteScheduledActionInput) (*applicationautoscaling.DeleteScheduledActionOutput, error) {
    return a.client.DeleteScheduledAction(input)
}



func (a *ApplicationAutoScalingActivities) DeregisterScalableTarget(input *applicationautoscaling.DeregisterScalableTargetInput) (*applicationautoscaling.DeregisterScalableTargetOutput, error) {
    return a.client.DeregisterScalableTarget(input)
}



func (a *ApplicationAutoScalingActivities) DescribeScalableTargets(input *applicationautoscaling.DescribeScalableTargetsInput) (*applicationautoscaling.DescribeScalableTargetsOutput, error) {
    return a.client.DescribeScalableTargets(input)
}



func (a *ApplicationAutoScalingActivities) DescribeScalingActivities(input *applicationautoscaling.DescribeScalingActivitiesInput) (*applicationautoscaling.DescribeScalingActivitiesOutput, error) {
    return a.client.DescribeScalingActivities(input)
}



func (a *ApplicationAutoScalingActivities) DescribeScalingPolicies(input *applicationautoscaling.DescribeScalingPoliciesInput) (*applicationautoscaling.DescribeScalingPoliciesOutput, error) {
    return a.client.DescribeScalingPolicies(input)
}



func (a *ApplicationAutoScalingActivities) DescribeScheduledActions(input *applicationautoscaling.DescribeScheduledActionsInput) (*applicationautoscaling.DescribeScheduledActionsOutput, error) {
    return a.client.DescribeScheduledActions(input)
}



func (a *ApplicationAutoScalingActivities) PutScalingPolicy(input *applicationautoscaling.PutScalingPolicyInput) (*applicationautoscaling.PutScalingPolicyOutput, error) {
    return a.client.PutScalingPolicy(input)
}



func (a *ApplicationAutoScalingActivities) PutScheduledAction(input *applicationautoscaling.PutScheduledActionInput) (*applicationautoscaling.PutScheduledActionOutput, error) {
    return a.client.PutScheduledAction(input)
}



func (a *ApplicationAutoScalingActivities) RegisterScalableTarget(input *applicationautoscaling.RegisterScalableTargetInput) (*applicationautoscaling.RegisterScalableTargetOutput, error) {
    return a.client.RegisterScalableTarget(input)
}

