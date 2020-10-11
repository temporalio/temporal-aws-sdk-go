// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package autoscalingplansstub

import (
	"github.com/aws/aws-sdk-go/service/autoscalingplans"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	CreateScalingPlan(ctx workflow.Context, input *autoscalingplans.CreateScalingPlanInput) (*autoscalingplans.CreateScalingPlanOutput, error)
	CreateScalingPlanAsync(ctx workflow.Context, input *autoscalingplans.CreateScalingPlanInput) *AutoScalingPlansCreateScalingPlanFuture

	DeleteScalingPlan(ctx workflow.Context, input *autoscalingplans.DeleteScalingPlanInput) (*autoscalingplans.DeleteScalingPlanOutput, error)
	DeleteScalingPlanAsync(ctx workflow.Context, input *autoscalingplans.DeleteScalingPlanInput) *AutoScalingPlansDeleteScalingPlanFuture

	DescribeScalingPlanResources(ctx workflow.Context, input *autoscalingplans.DescribeScalingPlanResourcesInput) (*autoscalingplans.DescribeScalingPlanResourcesOutput, error)
	DescribeScalingPlanResourcesAsync(ctx workflow.Context, input *autoscalingplans.DescribeScalingPlanResourcesInput) *AutoScalingPlansDescribeScalingPlanResourcesFuture

	DescribeScalingPlans(ctx workflow.Context, input *autoscalingplans.DescribeScalingPlansInput) (*autoscalingplans.DescribeScalingPlansOutput, error)
	DescribeScalingPlansAsync(ctx workflow.Context, input *autoscalingplans.DescribeScalingPlansInput) *AutoScalingPlansDescribeScalingPlansFuture

	GetScalingPlanResourceForecastData(ctx workflow.Context, input *autoscalingplans.GetScalingPlanResourceForecastDataInput) (*autoscalingplans.GetScalingPlanResourceForecastDataOutput, error)
	GetScalingPlanResourceForecastDataAsync(ctx workflow.Context, input *autoscalingplans.GetScalingPlanResourceForecastDataInput) *AutoScalingPlansGetScalingPlanResourceForecastDataFuture

	UpdateScalingPlan(ctx workflow.Context, input *autoscalingplans.UpdateScalingPlanInput) (*autoscalingplans.UpdateScalingPlanOutput, error)
	UpdateScalingPlanAsync(ctx workflow.Context, input *autoscalingplans.UpdateScalingPlanInput) *AutoScalingPlansUpdateScalingPlanFuture
}

func NewClient() Client {
	return &stub{}
}
