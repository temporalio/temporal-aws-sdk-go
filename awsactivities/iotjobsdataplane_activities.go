// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotjobsdataplane"
	"github.com/aws/aws-sdk-go/service/iotjobsdataplane/iotjobsdataplaneiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type IoTJobsDataPlaneActivities struct {
	client iotjobsdataplaneiface.IoTJobsDataPlaneAPI

	sessionFactory SessionFactory
}

func NewIoTJobsDataPlaneActivities(sess *session.Session, config ...*aws.Config) *IoTJobsDataPlaneActivities {
	client := iotjobsdataplane.New(sess, config...)
	return &IoTJobsDataPlaneActivities{client: client}
}

func NewIoTJobsDataPlaneActivitiesWithSessionFactory(sessionFactory SessionFactory) *IoTJobsDataPlaneActivities {
	return &IoTJobsDataPlaneActivities{sessionFactory: sessionFactory}
}

func (a *IoTJobsDataPlaneActivities) getClient(ctx context.Context) (iotjobsdataplaneiface.IoTJobsDataPlaneAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return iotjobsdataplane.New(sess), nil
}

func (a *IoTJobsDataPlaneActivities) DescribeJobExecution(ctx context.Context, input *iotjobsdataplane.DescribeJobExecutionInput) (*iotjobsdataplane.DescribeJobExecutionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeJobExecutionWithContext(ctx, input)
}

func (a *IoTJobsDataPlaneActivities) GetPendingJobExecutions(ctx context.Context, input *iotjobsdataplane.GetPendingJobExecutionsInput) (*iotjobsdataplane.GetPendingJobExecutionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetPendingJobExecutionsWithContext(ctx, input)
}

func (a *IoTJobsDataPlaneActivities) StartNextPendingJobExecution(ctx context.Context, input *iotjobsdataplane.StartNextPendingJobExecutionInput) (*iotjobsdataplane.StartNextPendingJobExecutionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartNextPendingJobExecutionWithContext(ctx, input)
}

func (a *IoTJobsDataPlaneActivities) UpdateJobExecution(ctx context.Context, input *iotjobsdataplane.UpdateJobExecutionInput) (*iotjobsdataplane.UpdateJobExecutionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateJobExecutionWithContext(ctx, input)
}
