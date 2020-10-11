// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package datapipelinestub

import (
	"github.com/aws/aws-sdk-go/service/datapipeline"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

type Client interface {
	ActivatePipeline(ctx workflow.Context, input *datapipeline.ActivatePipelineInput) (*datapipeline.ActivatePipelineOutput, error)
	ActivatePipelineAsync(ctx workflow.Context, input *datapipeline.ActivatePipelineInput) *DataPipelineActivatePipelineFuture

	AddTags(ctx workflow.Context, input *datapipeline.AddTagsInput) (*datapipeline.AddTagsOutput, error)
	AddTagsAsync(ctx workflow.Context, input *datapipeline.AddTagsInput) *DataPipelineAddTagsFuture

	CreatePipeline(ctx workflow.Context, input *datapipeline.CreatePipelineInput) (*datapipeline.CreatePipelineOutput, error)
	CreatePipelineAsync(ctx workflow.Context, input *datapipeline.CreatePipelineInput) *DataPipelineCreatePipelineFuture

	DeactivatePipeline(ctx workflow.Context, input *datapipeline.DeactivatePipelineInput) (*datapipeline.DeactivatePipelineOutput, error)
	DeactivatePipelineAsync(ctx workflow.Context, input *datapipeline.DeactivatePipelineInput) *DataPipelineDeactivatePipelineFuture

	DeletePipeline(ctx workflow.Context, input *datapipeline.DeletePipelineInput) (*datapipeline.DeletePipelineOutput, error)
	DeletePipelineAsync(ctx workflow.Context, input *datapipeline.DeletePipelineInput) *DataPipelineDeletePipelineFuture

	DescribeObjects(ctx workflow.Context, input *datapipeline.DescribeObjectsInput) (*datapipeline.DescribeObjectsOutput, error)
	DescribeObjectsAsync(ctx workflow.Context, input *datapipeline.DescribeObjectsInput) *DataPipelineDescribeObjectsFuture

	DescribePipelines(ctx workflow.Context, input *datapipeline.DescribePipelinesInput) (*datapipeline.DescribePipelinesOutput, error)
	DescribePipelinesAsync(ctx workflow.Context, input *datapipeline.DescribePipelinesInput) *DataPipelineDescribePipelinesFuture

	EvaluateExpression(ctx workflow.Context, input *datapipeline.EvaluateExpressionInput) (*datapipeline.EvaluateExpressionOutput, error)
	EvaluateExpressionAsync(ctx workflow.Context, input *datapipeline.EvaluateExpressionInput) *DataPipelineEvaluateExpressionFuture

	GetPipelineDefinition(ctx workflow.Context, input *datapipeline.GetPipelineDefinitionInput) (*datapipeline.GetPipelineDefinitionOutput, error)
	GetPipelineDefinitionAsync(ctx workflow.Context, input *datapipeline.GetPipelineDefinitionInput) *DataPipelineGetPipelineDefinitionFuture

	ListPipelines(ctx workflow.Context, input *datapipeline.ListPipelinesInput) (*datapipeline.ListPipelinesOutput, error)
	ListPipelinesAsync(ctx workflow.Context, input *datapipeline.ListPipelinesInput) *DataPipelineListPipelinesFuture

	PollForTask(ctx workflow.Context, input *datapipeline.PollForTaskInput) (*datapipeline.PollForTaskOutput, error)
	PollForTaskAsync(ctx workflow.Context, input *datapipeline.PollForTaskInput) *DataPipelinePollForTaskFuture

	PutPipelineDefinition(ctx workflow.Context, input *datapipeline.PutPipelineDefinitionInput) (*datapipeline.PutPipelineDefinitionOutput, error)
	PutPipelineDefinitionAsync(ctx workflow.Context, input *datapipeline.PutPipelineDefinitionInput) *DataPipelinePutPipelineDefinitionFuture

	QueryObjects(ctx workflow.Context, input *datapipeline.QueryObjectsInput) (*datapipeline.QueryObjectsOutput, error)
	QueryObjectsAsync(ctx workflow.Context, input *datapipeline.QueryObjectsInput) *DataPipelineQueryObjectsFuture

	RemoveTags(ctx workflow.Context, input *datapipeline.RemoveTagsInput) (*datapipeline.RemoveTagsOutput, error)
	RemoveTagsAsync(ctx workflow.Context, input *datapipeline.RemoveTagsInput) *DataPipelineRemoveTagsFuture

	ReportTaskProgress(ctx workflow.Context, input *datapipeline.ReportTaskProgressInput) (*datapipeline.ReportTaskProgressOutput, error)
	ReportTaskProgressAsync(ctx workflow.Context, input *datapipeline.ReportTaskProgressInput) *DataPipelineReportTaskProgressFuture

	ReportTaskRunnerHeartbeat(ctx workflow.Context, input *datapipeline.ReportTaskRunnerHeartbeatInput) (*datapipeline.ReportTaskRunnerHeartbeatOutput, error)
	ReportTaskRunnerHeartbeatAsync(ctx workflow.Context, input *datapipeline.ReportTaskRunnerHeartbeatInput) *DataPipelineReportTaskRunnerHeartbeatFuture

	SetStatus(ctx workflow.Context, input *datapipeline.SetStatusInput) (*datapipeline.SetStatusOutput, error)
	SetStatusAsync(ctx workflow.Context, input *datapipeline.SetStatusInput) *DataPipelineSetStatusFuture

	SetTaskStatus(ctx workflow.Context, input *datapipeline.SetTaskStatusInput) (*datapipeline.SetTaskStatusOutput, error)
	SetTaskStatusAsync(ctx workflow.Context, input *datapipeline.SetTaskStatusInput) *DataPipelineSetTaskStatusFuture

	ValidatePipelineDefinition(ctx workflow.Context, input *datapipeline.ValidatePipelineDefinitionInput) (*datapipeline.ValidatePipelineDefinitionOutput, error)
	ValidatePipelineDefinitionAsync(ctx workflow.Context, input *datapipeline.ValidatePipelineDefinitionInput) *DataPipelineValidatePipelineDefinitionFuture
}

func NewClient() Client {
	return &stub{}
}

