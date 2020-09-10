package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/datapipeline"
	"github.com/aws/aws-sdk-go/service/datapipeline/datapipelineiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type DataPipelineActivities struct {
	client datapipelineiface.DataPipelineAPI
}

func NewDataPipelineActivities(session *session.Session, config ...*aws.Config) *DataPipelineActivities {
	client := datapipeline.New(session, config...)
	return &DataPipelineActivities{client: client}
}

func (a *DataPipelineActivities) ActivatePipeline(ctx context.Context, input *datapipeline.ActivatePipelineInput) (*datapipeline.ActivatePipelineOutput, error) {
	return a.client.ActivatePipelineWithContext(ctx, input)
}

func (a *DataPipelineActivities) AddTags(ctx context.Context, input *datapipeline.AddTagsInput) (*datapipeline.AddTagsOutput, error) {
	return a.client.AddTagsWithContext(ctx, input)
}

func (a *DataPipelineActivities) CreatePipeline(ctx context.Context, input *datapipeline.CreatePipelineInput) (*datapipeline.CreatePipelineOutput, error) {
	return a.client.CreatePipelineWithContext(ctx, input)
}

func (a *DataPipelineActivities) DeactivatePipeline(ctx context.Context, input *datapipeline.DeactivatePipelineInput) (*datapipeline.DeactivatePipelineOutput, error) {
	return a.client.DeactivatePipelineWithContext(ctx, input)
}

func (a *DataPipelineActivities) DeletePipeline(ctx context.Context, input *datapipeline.DeletePipelineInput) (*datapipeline.DeletePipelineOutput, error) {
	return a.client.DeletePipelineWithContext(ctx, input)
}

func (a *DataPipelineActivities) DescribeObjects(ctx context.Context, input *datapipeline.DescribeObjectsInput) (*datapipeline.DescribeObjectsOutput, error) {
	return a.client.DescribeObjectsWithContext(ctx, input)
}

func (a *DataPipelineActivities) DescribePipelines(ctx context.Context, input *datapipeline.DescribePipelinesInput) (*datapipeline.DescribePipelinesOutput, error) {
	return a.client.DescribePipelinesWithContext(ctx, input)
}

func (a *DataPipelineActivities) EvaluateExpression(ctx context.Context, input *datapipeline.EvaluateExpressionInput) (*datapipeline.EvaluateExpressionOutput, error) {
	return a.client.EvaluateExpressionWithContext(ctx, input)
}

func (a *DataPipelineActivities) GetPipelineDefinition(ctx context.Context, input *datapipeline.GetPipelineDefinitionInput) (*datapipeline.GetPipelineDefinitionOutput, error) {
	return a.client.GetPipelineDefinitionWithContext(ctx, input)
}

func (a *DataPipelineActivities) ListPipelines(ctx context.Context, input *datapipeline.ListPipelinesInput) (*datapipeline.ListPipelinesOutput, error) {
	return a.client.ListPipelinesWithContext(ctx, input)
}

func (a *DataPipelineActivities) PollForTask(ctx context.Context, input *datapipeline.PollForTaskInput) (*datapipeline.PollForTaskOutput, error) {
	return a.client.PollForTaskWithContext(ctx, input)
}

func (a *DataPipelineActivities) PutPipelineDefinition(ctx context.Context, input *datapipeline.PutPipelineDefinitionInput) (*datapipeline.PutPipelineDefinitionOutput, error) {
	return a.client.PutPipelineDefinitionWithContext(ctx, input)
}

func (a *DataPipelineActivities) QueryObjects(ctx context.Context, input *datapipeline.QueryObjectsInput) (*datapipeline.QueryObjectsOutput, error) {
	return a.client.QueryObjectsWithContext(ctx, input)
}

func (a *DataPipelineActivities) RemoveTags(ctx context.Context, input *datapipeline.RemoveTagsInput) (*datapipeline.RemoveTagsOutput, error) {
	return a.client.RemoveTagsWithContext(ctx, input)
}

func (a *DataPipelineActivities) ReportTaskProgress(ctx context.Context, input *datapipeline.ReportTaskProgressInput) (*datapipeline.ReportTaskProgressOutput, error) {
	return a.client.ReportTaskProgressWithContext(ctx, input)
}

func (a *DataPipelineActivities) ReportTaskRunnerHeartbeat(ctx context.Context, input *datapipeline.ReportTaskRunnerHeartbeatInput) (*datapipeline.ReportTaskRunnerHeartbeatOutput, error) {
	return a.client.ReportTaskRunnerHeartbeatWithContext(ctx, input)
}

func (a *DataPipelineActivities) SetStatus(ctx context.Context, input *datapipeline.SetStatusInput) (*datapipeline.SetStatusOutput, error) {
	return a.client.SetStatusWithContext(ctx, input)
}

func (a *DataPipelineActivities) SetTaskStatus(ctx context.Context, input *datapipeline.SetTaskStatusInput) (*datapipeline.SetTaskStatusOutput, error) {
	return a.client.SetTaskStatusWithContext(ctx, input)
}

func (a *DataPipelineActivities) ValidatePipelineDefinition(ctx context.Context, input *datapipeline.ValidatePipelineDefinitionInput) (*datapipeline.ValidatePipelineDefinitionOutput, error) {
	return a.client.ValidatePipelineDefinitionWithContext(ctx, input)
}
