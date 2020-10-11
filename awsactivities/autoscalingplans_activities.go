// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscalingplans"
	"github.com/aws/aws-sdk-go/service/autoscalingplans/autoscalingplansiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type AutoScalingPlansActivities struct {
	client autoscalingplansiface.AutoScalingPlansAPI

	sessionFactory SessionFactory
}

func NewAutoScalingPlansActivities(sess *session.Session, config ...*aws.Config) *AutoScalingPlansActivities {
	client := autoscalingplans.New(sess, config...)
	return &AutoScalingPlansActivities{client: client}
}

func NewAutoScalingPlansActivitiesWithSessionFactory(sessionFactory SessionFactory) *AutoScalingPlansActivities {
	return &AutoScalingPlansActivities{sessionFactory: sessionFactory}
}

func (a *AutoScalingPlansActivities) getClient(ctx context.Context) (autoscalingplansiface.AutoScalingPlansAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return autoscalingplans.New(sess), nil
}

func (a *AutoScalingPlansActivities) CreateScalingPlan(ctx context.Context, input *autoscalingplans.CreateScalingPlanInput) (*autoscalingplans.CreateScalingPlanOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateScalingPlanWithContext(ctx, input)
}

func (a *AutoScalingPlansActivities) DeleteScalingPlan(ctx context.Context, input *autoscalingplans.DeleteScalingPlanInput) (*autoscalingplans.DeleteScalingPlanOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteScalingPlanWithContext(ctx, input)
}

func (a *AutoScalingPlansActivities) DescribeScalingPlanResources(ctx context.Context, input *autoscalingplans.DescribeScalingPlanResourcesInput) (*autoscalingplans.DescribeScalingPlanResourcesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeScalingPlanResourcesWithContext(ctx, input)
}

func (a *AutoScalingPlansActivities) DescribeScalingPlans(ctx context.Context, input *autoscalingplans.DescribeScalingPlansInput) (*autoscalingplans.DescribeScalingPlansOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeScalingPlansWithContext(ctx, input)
}

func (a *AutoScalingPlansActivities) GetScalingPlanResourceForecastData(ctx context.Context, input *autoscalingplans.GetScalingPlanResourceForecastDataInput) (*autoscalingplans.GetScalingPlanResourceForecastDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetScalingPlanResourceForecastDataWithContext(ctx, input)
}

func (a *AutoScalingPlansActivities) UpdateScalingPlan(ctx context.Context, input *autoscalingplans.UpdateScalingPlanInput) (*autoscalingplans.UpdateScalingPlanOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateScalingPlanWithContext(ctx, input)
}
