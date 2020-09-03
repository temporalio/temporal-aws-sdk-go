package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/datapipeline"
	"github.com/aws/aws-sdk-go/service/datapipeline/datapipelineiface"
)

type DataPipelineActivities struct {
	client datapipelineiface.DataPipelineAPI
}

func NewDataPipelineActivities(client datapipelineiface.DataPipelineAPI) *DataPipelineActivities {
    return &DataPipelineActivities{client: client}
}

func (a *DataPipelineActivities) ActivatePipeline(input *datapipeline.ActivatePipelineInput) (*datapipeline.ActivatePipelineOutput, error) {
    return a.client.ActivatePipeline(input)
}

func (a *DataPipelineActivities) AddTags(input *datapipeline.AddTagsInput) (*datapipeline.AddTagsOutput, error) {
    return a.client.AddTags(input)
}

func (a *DataPipelineActivities) CreatePipeline(input *datapipeline.CreatePipelineInput) (*datapipeline.CreatePipelineOutput, error) {
    return a.client.CreatePipeline(input)
}

func (a *DataPipelineActivities) DeactivatePipeline(input *datapipeline.DeactivatePipelineInput) (*datapipeline.DeactivatePipelineOutput, error) {
    return a.client.DeactivatePipeline(input)
}

func (a *DataPipelineActivities) DeletePipeline(input *datapipeline.DeletePipelineInput) (*datapipeline.DeletePipelineOutput, error) {
    return a.client.DeletePipeline(input)
}

func (a *DataPipelineActivities) DescribeObjects(input *datapipeline.DescribeObjectsInput) (*datapipeline.DescribeObjectsOutput, error) {
    return a.client.DescribeObjects(input)
}

func (a *DataPipelineActivities) DescribePipelines(input *datapipeline.DescribePipelinesInput) (*datapipeline.DescribePipelinesOutput, error) {
    return a.client.DescribePipelines(input)
}

func (a *DataPipelineActivities) EvaluateExpression(input *datapipeline.EvaluateExpressionInput) (*datapipeline.EvaluateExpressionOutput, error) {
    return a.client.EvaluateExpression(input)
}

func (a *DataPipelineActivities) GetPipelineDefinition(input *datapipeline.GetPipelineDefinitionInput) (*datapipeline.GetPipelineDefinitionOutput, error) {
    return a.client.GetPipelineDefinition(input)
}

func (a *DataPipelineActivities) ListPipelines(input *datapipeline.ListPipelinesInput) (*datapipeline.ListPipelinesOutput, error) {
    return a.client.ListPipelines(input)
}

func (a *DataPipelineActivities) PollForTask(input *datapipeline.PollForTaskInput) (*datapipeline.PollForTaskOutput, error) {
    return a.client.PollForTask(input)
}

func (a *DataPipelineActivities) PutPipelineDefinition(input *datapipeline.PutPipelineDefinitionInput) (*datapipeline.PutPipelineDefinitionOutput, error) {
    return a.client.PutPipelineDefinition(input)
}

func (a *DataPipelineActivities) QueryObjects(input *datapipeline.QueryObjectsInput) (*datapipeline.QueryObjectsOutput, error) {
    return a.client.QueryObjects(input)
}

func (a *DataPipelineActivities) RemoveTags(input *datapipeline.RemoveTagsInput) (*datapipeline.RemoveTagsOutput, error) {
    return a.client.RemoveTags(input)
}

func (a *DataPipelineActivities) ReportTaskProgress(input *datapipeline.ReportTaskProgressInput) (*datapipeline.ReportTaskProgressOutput, error) {
    return a.client.ReportTaskProgress(input)
}

func (a *DataPipelineActivities) ReportTaskRunnerHeartbeat(input *datapipeline.ReportTaskRunnerHeartbeatInput) (*datapipeline.ReportTaskRunnerHeartbeatOutput, error) {
    return a.client.ReportTaskRunnerHeartbeat(input)
}

func (a *DataPipelineActivities) SetStatus(input *datapipeline.SetStatusInput) (*datapipeline.SetStatusOutput, error) {
    return a.client.SetStatus(input)
}

func (a *DataPipelineActivities) SetTaskStatus(input *datapipeline.SetTaskStatusInput) (*datapipeline.SetTaskStatusOutput, error) {
    return a.client.SetTaskStatus(input)
}

func (a *DataPipelineActivities) ValidatePipelineDefinition(input *datapipeline.ValidatePipelineDefinitionInput) (*datapipeline.ValidatePipelineDefinitionOutput, error) {
    return a.client.ValidatePipelineDefinition(input)
}
