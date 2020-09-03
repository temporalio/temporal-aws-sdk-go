package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/autoscalingplans"
	"github.com/aws/aws-sdk-go/service/autoscalingplans/autoscalingplansiface"
)

type AutoScalingPlansActivities struct {
	client autoscalingplansiface.AutoScalingPlansAPI
}

func NewAutoScalingPlansActivities(client autoscalingplansiface.AutoScalingPlansAPI) *AutoScalingPlansActivities {
    return &AutoScalingPlansActivities{client: client}
}

func (a *AutoScalingPlansActivities) CreateScalingPlan(input *autoscalingplans.CreateScalingPlanInput) (*autoscalingplans.CreateScalingPlanOutput, error) {
    return a.client.CreateScalingPlan(input)
}

func (a *AutoScalingPlansActivities) DeleteScalingPlan(input *autoscalingplans.DeleteScalingPlanInput) (*autoscalingplans.DeleteScalingPlanOutput, error) {
    return a.client.DeleteScalingPlan(input)
}

func (a *AutoScalingPlansActivities) DescribeScalingPlanResources(input *autoscalingplans.DescribeScalingPlanResourcesInput) (*autoscalingplans.DescribeScalingPlanResourcesOutput, error) {
    return a.client.DescribeScalingPlanResources(input)
}

func (a *AutoScalingPlansActivities) DescribeScalingPlans(input *autoscalingplans.DescribeScalingPlansInput) (*autoscalingplans.DescribeScalingPlansOutput, error) {
    return a.client.DescribeScalingPlans(input)
}

func (a *AutoScalingPlansActivities) GetScalingPlanResourceForecastData(input *autoscalingplans.GetScalingPlanResourceForecastDataInput) (*autoscalingplans.GetScalingPlanResourceForecastDataOutput, error) {
    return a.client.GetScalingPlanResourceForecastData(input)
}

func (a *AutoScalingPlansActivities) UpdateScalingPlan(input *autoscalingplans.UpdateScalingPlanInput) (*autoscalingplans.UpdateScalingPlanOutput, error) {
    return a.client.UpdateScalingPlan(input)
}
