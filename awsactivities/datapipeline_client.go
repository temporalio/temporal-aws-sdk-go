package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/datapipeline"
	"go.temporal.io/sdk/workflow"
)

type DataPipelineClient interface {
    ActivatePipeline(ctx workflow.Context, input *datapipeline.ActivatePipelineInput) (*datapipeline.ActivatePipelineOutput, error)
    ActivatePipelineAsync(ctx workflow.Context, input *datapipeline.ActivatePipelineInput) *DatapipelineActivatePipelineResult

    AddTags(ctx workflow.Context, input *datapipeline.AddTagsInput) (*datapipeline.AddTagsOutput, error)
    AddTagsAsync(ctx workflow.Context, input *datapipeline.AddTagsInput) *DatapipelineAddTagsResult

    CreatePipeline(ctx workflow.Context, input *datapipeline.CreatePipelineInput) (*datapipeline.CreatePipelineOutput, error)
    CreatePipelineAsync(ctx workflow.Context, input *datapipeline.CreatePipelineInput) *DatapipelineCreatePipelineResult

    DeactivatePipeline(ctx workflow.Context, input *datapipeline.DeactivatePipelineInput) (*datapipeline.DeactivatePipelineOutput, error)
    DeactivatePipelineAsync(ctx workflow.Context, input *datapipeline.DeactivatePipelineInput) *DatapipelineDeactivatePipelineResult

    DeletePipeline(ctx workflow.Context, input *datapipeline.DeletePipelineInput) (*datapipeline.DeletePipelineOutput, error)
    DeletePipelineAsync(ctx workflow.Context, input *datapipeline.DeletePipelineInput) *DatapipelineDeletePipelineResult

    DescribeObjects(ctx workflow.Context, input *datapipeline.DescribeObjectsInput) (*datapipeline.DescribeObjectsOutput, error)
    DescribeObjectsAsync(ctx workflow.Context, input *datapipeline.DescribeObjectsInput) *DatapipelineDescribeObjectsResult

    DescribePipelines(ctx workflow.Context, input *datapipeline.DescribePipelinesInput) (*datapipeline.DescribePipelinesOutput, error)
    DescribePipelinesAsync(ctx workflow.Context, input *datapipeline.DescribePipelinesInput) *DatapipelineDescribePipelinesResult

    EvaluateExpression(ctx workflow.Context, input *datapipeline.EvaluateExpressionInput) (*datapipeline.EvaluateExpressionOutput, error)
    EvaluateExpressionAsync(ctx workflow.Context, input *datapipeline.EvaluateExpressionInput) *DatapipelineEvaluateExpressionResult

    GetPipelineDefinition(ctx workflow.Context, input *datapipeline.GetPipelineDefinitionInput) (*datapipeline.GetPipelineDefinitionOutput, error)
    GetPipelineDefinitionAsync(ctx workflow.Context, input *datapipeline.GetPipelineDefinitionInput) *DatapipelineGetPipelineDefinitionResult

    ListPipelines(ctx workflow.Context, input *datapipeline.ListPipelinesInput) (*datapipeline.ListPipelinesOutput, error)
    ListPipelinesAsync(ctx workflow.Context, input *datapipeline.ListPipelinesInput) *DatapipelineListPipelinesResult

    PollForTask(ctx workflow.Context, input *datapipeline.PollForTaskInput) (*datapipeline.PollForTaskOutput, error)
    PollForTaskAsync(ctx workflow.Context, input *datapipeline.PollForTaskInput) *DatapipelinePollForTaskResult

    PutPipelineDefinition(ctx workflow.Context, input *datapipeline.PutPipelineDefinitionInput) (*datapipeline.PutPipelineDefinitionOutput, error)
    PutPipelineDefinitionAsync(ctx workflow.Context, input *datapipeline.PutPipelineDefinitionInput) *DatapipelinePutPipelineDefinitionResult

    QueryObjects(ctx workflow.Context, input *datapipeline.QueryObjectsInput) (*datapipeline.QueryObjectsOutput, error)
    QueryObjectsAsync(ctx workflow.Context, input *datapipeline.QueryObjectsInput) *DatapipelineQueryObjectsResult

    RemoveTags(ctx workflow.Context, input *datapipeline.RemoveTagsInput) (*datapipeline.RemoveTagsOutput, error)
    RemoveTagsAsync(ctx workflow.Context, input *datapipeline.RemoveTagsInput) *DatapipelineRemoveTagsResult

    ReportTaskProgress(ctx workflow.Context, input *datapipeline.ReportTaskProgressInput) (*datapipeline.ReportTaskProgressOutput, error)
    ReportTaskProgressAsync(ctx workflow.Context, input *datapipeline.ReportTaskProgressInput) *DatapipelineReportTaskProgressResult

    ReportTaskRunnerHeartbeat(ctx workflow.Context, input *datapipeline.ReportTaskRunnerHeartbeatInput) (*datapipeline.ReportTaskRunnerHeartbeatOutput, error)
    ReportTaskRunnerHeartbeatAsync(ctx workflow.Context, input *datapipeline.ReportTaskRunnerHeartbeatInput) *DatapipelineReportTaskRunnerHeartbeatResult

    SetStatus(ctx workflow.Context, input *datapipeline.SetStatusInput) (*datapipeline.SetStatusOutput, error)
    SetStatusAsync(ctx workflow.Context, input *datapipeline.SetStatusInput) *DatapipelineSetStatusResult

    SetTaskStatus(ctx workflow.Context, input *datapipeline.SetTaskStatusInput) (*datapipeline.SetTaskStatusOutput, error)
    SetTaskStatusAsync(ctx workflow.Context, input *datapipeline.SetTaskStatusInput) *DatapipelineSetTaskStatusResult

    ValidatePipelineDefinition(ctx workflow.Context, input *datapipeline.ValidatePipelineDefinitionInput) (*datapipeline.ValidatePipelineDefinitionOutput, error)
    ValidatePipelineDefinitionAsync(ctx workflow.Context, input *datapipeline.ValidatePipelineDefinitionInput) *DatapipelineValidatePipelineDefinitionResult
}
type DatapipelineActivatePipelineResult struct {
	Result workflow.Future
}

func (r *DatapipelineActivatePipelineResult) Get(ctx workflow.Context) (*datapipeline.ActivatePipelineOutput, error) {
    var output datapipeline.ActivatePipelineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatapipelineAddTagsResult struct {
	Result workflow.Future
}

func (r *DatapipelineAddTagsResult) Get(ctx workflow.Context) (*datapipeline.AddTagsOutput, error) {
    var output datapipeline.AddTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatapipelineCreatePipelineResult struct {
	Result workflow.Future
}

func (r *DatapipelineCreatePipelineResult) Get(ctx workflow.Context) (*datapipeline.CreatePipelineOutput, error) {
    var output datapipeline.CreatePipelineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatapipelineDeactivatePipelineResult struct {
	Result workflow.Future
}

func (r *DatapipelineDeactivatePipelineResult) Get(ctx workflow.Context) (*datapipeline.DeactivatePipelineOutput, error) {
    var output datapipeline.DeactivatePipelineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatapipelineDeletePipelineResult struct {
	Result workflow.Future
}

func (r *DatapipelineDeletePipelineResult) Get(ctx workflow.Context) (*datapipeline.DeletePipelineOutput, error) {
    var output datapipeline.DeletePipelineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatapipelineDescribeObjectsResult struct {
	Result workflow.Future
}

func (r *DatapipelineDescribeObjectsResult) Get(ctx workflow.Context) (*datapipeline.DescribeObjectsOutput, error) {
    var output datapipeline.DescribeObjectsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatapipelineDescribePipelinesResult struct {
	Result workflow.Future
}

func (r *DatapipelineDescribePipelinesResult) Get(ctx workflow.Context) (*datapipeline.DescribePipelinesOutput, error) {
    var output datapipeline.DescribePipelinesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatapipelineEvaluateExpressionResult struct {
	Result workflow.Future
}

func (r *DatapipelineEvaluateExpressionResult) Get(ctx workflow.Context) (*datapipeline.EvaluateExpressionOutput, error) {
    var output datapipeline.EvaluateExpressionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatapipelineGetPipelineDefinitionResult struct {
	Result workflow.Future
}

func (r *DatapipelineGetPipelineDefinitionResult) Get(ctx workflow.Context) (*datapipeline.GetPipelineDefinitionOutput, error) {
    var output datapipeline.GetPipelineDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatapipelineListPipelinesResult struct {
	Result workflow.Future
}

func (r *DatapipelineListPipelinesResult) Get(ctx workflow.Context) (*datapipeline.ListPipelinesOutput, error) {
    var output datapipeline.ListPipelinesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatapipelinePollForTaskResult struct {
	Result workflow.Future
}

func (r *DatapipelinePollForTaskResult) Get(ctx workflow.Context) (*datapipeline.PollForTaskOutput, error) {
    var output datapipeline.PollForTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatapipelinePutPipelineDefinitionResult struct {
	Result workflow.Future
}

func (r *DatapipelinePutPipelineDefinitionResult) Get(ctx workflow.Context) (*datapipeline.PutPipelineDefinitionOutput, error) {
    var output datapipeline.PutPipelineDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatapipelineQueryObjectsResult struct {
	Result workflow.Future
}

func (r *DatapipelineQueryObjectsResult) Get(ctx workflow.Context) (*datapipeline.QueryObjectsOutput, error) {
    var output datapipeline.QueryObjectsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatapipelineRemoveTagsResult struct {
	Result workflow.Future
}

func (r *DatapipelineRemoveTagsResult) Get(ctx workflow.Context) (*datapipeline.RemoveTagsOutput, error) {
    var output datapipeline.RemoveTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatapipelineReportTaskProgressResult struct {
	Result workflow.Future
}

func (r *DatapipelineReportTaskProgressResult) Get(ctx workflow.Context) (*datapipeline.ReportTaskProgressOutput, error) {
    var output datapipeline.ReportTaskProgressOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatapipelineReportTaskRunnerHeartbeatResult struct {
	Result workflow.Future
}

func (r *DatapipelineReportTaskRunnerHeartbeatResult) Get(ctx workflow.Context) (*datapipeline.ReportTaskRunnerHeartbeatOutput, error) {
    var output datapipeline.ReportTaskRunnerHeartbeatOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatapipelineSetStatusResult struct {
	Result workflow.Future
}

func (r *DatapipelineSetStatusResult) Get(ctx workflow.Context) (*datapipeline.SetStatusOutput, error) {
    var output datapipeline.SetStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatapipelineSetTaskStatusResult struct {
	Result workflow.Future
}

func (r *DatapipelineSetTaskStatusResult) Get(ctx workflow.Context) (*datapipeline.SetTaskStatusOutput, error) {
    var output datapipeline.SetTaskStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatapipelineValidatePipelineDefinitionResult struct {
	Result workflow.Future
}

func (r *DatapipelineValidatePipelineDefinitionResult) Get(ctx workflow.Context) (*datapipeline.ValidatePipelineDefinitionOutput, error) {
    var output datapipeline.ValidatePipelineDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type DataPipelineStub struct {
    activities DataPipelineClient
}

func NewDataPipelineStub() DataPipelineClient {
    return &DataPipelineStub{}
}

func (a *DataPipelineStub) ActivatePipeline(ctx workflow.Context, input *datapipeline.ActivatePipelineInput) (*datapipeline.ActivatePipelineOutput, error) {
    var output datapipeline.ActivatePipelineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ActivatePipeline, input).Get(ctx, &output)
    return &output, err
}

func (a *DataPipelineStub) AddTags(ctx workflow.Context, input *datapipeline.AddTagsInput) (*datapipeline.AddTagsOutput, error) {
    var output datapipeline.AddTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddTags, input).Get(ctx, &output)
    return &output, err
}

func (a *DataPipelineStub) CreatePipeline(ctx workflow.Context, input *datapipeline.CreatePipelineInput) (*datapipeline.CreatePipelineOutput, error) {
    var output datapipeline.CreatePipelineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePipeline, input).Get(ctx, &output)
    return &output, err
}

func (a *DataPipelineStub) DeactivatePipeline(ctx workflow.Context, input *datapipeline.DeactivatePipelineInput) (*datapipeline.DeactivatePipelineOutput, error) {
    var output datapipeline.DeactivatePipelineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeactivatePipeline, input).Get(ctx, &output)
    return &output, err
}

func (a *DataPipelineStub) DeletePipeline(ctx workflow.Context, input *datapipeline.DeletePipelineInput) (*datapipeline.DeletePipelineOutput, error) {
    var output datapipeline.DeletePipelineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePipeline, input).Get(ctx, &output)
    return &output, err
}

func (a *DataPipelineStub) DescribeObjects(ctx workflow.Context, input *datapipeline.DescribeObjectsInput) (*datapipeline.DescribeObjectsOutput, error) {
    var output datapipeline.DescribeObjectsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeObjects, input).Get(ctx, &output)
    return &output, err
}

func (a *DataPipelineStub) DescribePipelines(ctx workflow.Context, input *datapipeline.DescribePipelinesInput) (*datapipeline.DescribePipelinesOutput, error) {
    var output datapipeline.DescribePipelinesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePipelines, input).Get(ctx, &output)
    return &output, err
}

func (a *DataPipelineStub) EvaluateExpression(ctx workflow.Context, input *datapipeline.EvaluateExpressionInput) (*datapipeline.EvaluateExpressionOutput, error) {
    var output datapipeline.EvaluateExpressionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EvaluateExpression, input).Get(ctx, &output)
    return &output, err
}

func (a *DataPipelineStub) GetPipelineDefinition(ctx workflow.Context, input *datapipeline.GetPipelineDefinitionInput) (*datapipeline.GetPipelineDefinitionOutput, error) {
    var output datapipeline.GetPipelineDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPipelineDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *DataPipelineStub) ListPipelines(ctx workflow.Context, input *datapipeline.ListPipelinesInput) (*datapipeline.ListPipelinesOutput, error) {
    var output datapipeline.ListPipelinesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPipelines, input).Get(ctx, &output)
    return &output, err
}

func (a *DataPipelineStub) PollForTask(ctx workflow.Context, input *datapipeline.PollForTaskInput) (*datapipeline.PollForTaskOutput, error) {
    var output datapipeline.PollForTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PollForTask, input).Get(ctx, &output)
    return &output, err
}

func (a *DataPipelineStub) PutPipelineDefinition(ctx workflow.Context, input *datapipeline.PutPipelineDefinitionInput) (*datapipeline.PutPipelineDefinitionOutput, error) {
    var output datapipeline.PutPipelineDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutPipelineDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *DataPipelineStub) QueryObjects(ctx workflow.Context, input *datapipeline.QueryObjectsInput) (*datapipeline.QueryObjectsOutput, error) {
    var output datapipeline.QueryObjectsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.QueryObjects, input).Get(ctx, &output)
    return &output, err
}

func (a *DataPipelineStub) RemoveTags(ctx workflow.Context, input *datapipeline.RemoveTagsInput) (*datapipeline.RemoveTagsOutput, error) {
    var output datapipeline.RemoveTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveTags, input).Get(ctx, &output)
    return &output, err
}

func (a *DataPipelineStub) ReportTaskProgress(ctx workflow.Context, input *datapipeline.ReportTaskProgressInput) (*datapipeline.ReportTaskProgressOutput, error) {
    var output datapipeline.ReportTaskProgressOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ReportTaskProgress, input).Get(ctx, &output)
    return &output, err
}

func (a *DataPipelineStub) ReportTaskRunnerHeartbeat(ctx workflow.Context, input *datapipeline.ReportTaskRunnerHeartbeatInput) (*datapipeline.ReportTaskRunnerHeartbeatOutput, error) {
    var output datapipeline.ReportTaskRunnerHeartbeatOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ReportTaskRunnerHeartbeat, input).Get(ctx, &output)
    return &output, err
}

func (a *DataPipelineStub) SetStatus(ctx workflow.Context, input *datapipeline.SetStatusInput) (*datapipeline.SetStatusOutput, error) {
    var output datapipeline.SetStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *DataPipelineStub) SetTaskStatus(ctx workflow.Context, input *datapipeline.SetTaskStatusInput) (*datapipeline.SetTaskStatusOutput, error) {
    var output datapipeline.SetTaskStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetTaskStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *DataPipelineStub) ValidatePipelineDefinition(ctx workflow.Context, input *datapipeline.ValidatePipelineDefinitionInput) (*datapipeline.ValidatePipelineDefinitionOutput, error) {
    var output datapipeline.ValidatePipelineDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ValidatePipelineDefinition, input).Get(ctx, &output)
    return &output, err
}
