// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package iotjobsdataplanestub

import (
	"github.com/aws/aws-sdk-go/service/iotjobsdataplane"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	DescribeJobExecution(ctx workflow.Context, input *iotjobsdataplane.DescribeJobExecutionInput) (*iotjobsdataplane.DescribeJobExecutionOutput, error)
	DescribeJobExecutionAsync(ctx workflow.Context, input *iotjobsdataplane.DescribeJobExecutionInput) *IoTJobsDataPlaneDescribeJobExecutionFuture

	GetPendingJobExecutions(ctx workflow.Context, input *iotjobsdataplane.GetPendingJobExecutionsInput) (*iotjobsdataplane.GetPendingJobExecutionsOutput, error)
	GetPendingJobExecutionsAsync(ctx workflow.Context, input *iotjobsdataplane.GetPendingJobExecutionsInput) *IoTJobsDataPlaneGetPendingJobExecutionsFuture

	StartNextPendingJobExecution(ctx workflow.Context, input *iotjobsdataplane.StartNextPendingJobExecutionInput) (*iotjobsdataplane.StartNextPendingJobExecutionOutput, error)
	StartNextPendingJobExecutionAsync(ctx workflow.Context, input *iotjobsdataplane.StartNextPendingJobExecutionInput) *IoTJobsDataPlaneStartNextPendingJobExecutionFuture

	UpdateJobExecution(ctx workflow.Context, input *iotjobsdataplane.UpdateJobExecutionInput) (*iotjobsdataplane.UpdateJobExecutionOutput, error)
	UpdateJobExecutionAsync(ctx workflow.Context, input *iotjobsdataplane.UpdateJobExecutionInput) *IoTJobsDataPlaneUpdateJobExecutionFuture
}

func NewClient() Client {
	return &stub{}
}
