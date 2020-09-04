package awsclients

import (
	"github.com/aws/aws-sdk-go/service/batch"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type BatchClient interface {
    CancelJob(ctx workflow.Context, input *batch.CancelJobInput) (*batch.CancelJobOutput, error)
    CancelJobAsync(ctx workflow.Context, input *batch.CancelJobInput) *BatchCancelJobResult

    CreateComputeEnvironment(ctx workflow.Context, input *batch.CreateComputeEnvironmentInput) (*batch.CreateComputeEnvironmentOutput, error)
    CreateComputeEnvironmentAsync(ctx workflow.Context, input *batch.CreateComputeEnvironmentInput) *BatchCreateComputeEnvironmentResult

    CreateJobQueue(ctx workflow.Context, input *batch.CreateJobQueueInput) (*batch.CreateJobQueueOutput, error)
    CreateJobQueueAsync(ctx workflow.Context, input *batch.CreateJobQueueInput) *BatchCreateJobQueueResult

    DeleteComputeEnvironment(ctx workflow.Context, input *batch.DeleteComputeEnvironmentInput) (*batch.DeleteComputeEnvironmentOutput, error)
    DeleteComputeEnvironmentAsync(ctx workflow.Context, input *batch.DeleteComputeEnvironmentInput) *BatchDeleteComputeEnvironmentResult

    DeleteJobQueue(ctx workflow.Context, input *batch.DeleteJobQueueInput) (*batch.DeleteJobQueueOutput, error)
    DeleteJobQueueAsync(ctx workflow.Context, input *batch.DeleteJobQueueInput) *BatchDeleteJobQueueResult

    DeregisterJobDefinition(ctx workflow.Context, input *batch.DeregisterJobDefinitionInput) (*batch.DeregisterJobDefinitionOutput, error)
    DeregisterJobDefinitionAsync(ctx workflow.Context, input *batch.DeregisterJobDefinitionInput) *BatchDeregisterJobDefinitionResult

    DescribeComputeEnvironments(ctx workflow.Context, input *batch.DescribeComputeEnvironmentsInput) (*batch.DescribeComputeEnvironmentsOutput, error)
    DescribeComputeEnvironmentsAsync(ctx workflow.Context, input *batch.DescribeComputeEnvironmentsInput) *BatchDescribeComputeEnvironmentsResult

    DescribeJobDefinitions(ctx workflow.Context, input *batch.DescribeJobDefinitionsInput) (*batch.DescribeJobDefinitionsOutput, error)
    DescribeJobDefinitionsAsync(ctx workflow.Context, input *batch.DescribeJobDefinitionsInput) *BatchDescribeJobDefinitionsResult

    DescribeJobQueues(ctx workflow.Context, input *batch.DescribeJobQueuesInput) (*batch.DescribeJobQueuesOutput, error)
    DescribeJobQueuesAsync(ctx workflow.Context, input *batch.DescribeJobQueuesInput) *BatchDescribeJobQueuesResult

    DescribeJobs(ctx workflow.Context, input *batch.DescribeJobsInput) (*batch.DescribeJobsOutput, error)
    DescribeJobsAsync(ctx workflow.Context, input *batch.DescribeJobsInput) *BatchDescribeJobsResult

    ListJobs(ctx workflow.Context, input *batch.ListJobsInput) (*batch.ListJobsOutput, error)
    ListJobsAsync(ctx workflow.Context, input *batch.ListJobsInput) *BatchListJobsResult

    RegisterJobDefinition(ctx workflow.Context, input *batch.RegisterJobDefinitionInput) (*batch.RegisterJobDefinitionOutput, error)
    RegisterJobDefinitionAsync(ctx workflow.Context, input *batch.RegisterJobDefinitionInput) *BatchRegisterJobDefinitionResult

    SubmitJob(ctx workflow.Context, input *batch.SubmitJobInput) (*batch.SubmitJobOutput, error)
    SubmitJobAsync(ctx workflow.Context, input *batch.SubmitJobInput) *BatchSubmitJobResult

    TerminateJob(ctx workflow.Context, input *batch.TerminateJobInput) (*batch.TerminateJobOutput, error)
    TerminateJobAsync(ctx workflow.Context, input *batch.TerminateJobInput) *BatchTerminateJobResult

    UpdateComputeEnvironment(ctx workflow.Context, input *batch.UpdateComputeEnvironmentInput) (*batch.UpdateComputeEnvironmentOutput, error)
    UpdateComputeEnvironmentAsync(ctx workflow.Context, input *batch.UpdateComputeEnvironmentInput) *BatchUpdateComputeEnvironmentResult

    UpdateJobQueue(ctx workflow.Context, input *batch.UpdateJobQueueInput) (*batch.UpdateJobQueueOutput, error)
    UpdateJobQueueAsync(ctx workflow.Context, input *batch.UpdateJobQueueInput) *BatchUpdateJobQueueResult
}

type BatchCancelJobResult struct {
	Result workflow.Future
}

func (r *BatchCancelJobResult) Get(ctx workflow.Context) (*batch.CancelJobOutput, error) {
    var output batch.CancelJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BatchCreateComputeEnvironmentResult struct {
	Result workflow.Future
}

func (r *BatchCreateComputeEnvironmentResult) Get(ctx workflow.Context) (*batch.CreateComputeEnvironmentOutput, error) {
    var output batch.CreateComputeEnvironmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BatchCreateJobQueueResult struct {
	Result workflow.Future
}

func (r *BatchCreateJobQueueResult) Get(ctx workflow.Context) (*batch.CreateJobQueueOutput, error) {
    var output batch.CreateJobQueueOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BatchDeleteComputeEnvironmentResult struct {
	Result workflow.Future
}

func (r *BatchDeleteComputeEnvironmentResult) Get(ctx workflow.Context) (*batch.DeleteComputeEnvironmentOutput, error) {
    var output batch.DeleteComputeEnvironmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BatchDeleteJobQueueResult struct {
	Result workflow.Future
}

func (r *BatchDeleteJobQueueResult) Get(ctx workflow.Context) (*batch.DeleteJobQueueOutput, error) {
    var output batch.DeleteJobQueueOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BatchDeregisterJobDefinitionResult struct {
	Result workflow.Future
}

func (r *BatchDeregisterJobDefinitionResult) Get(ctx workflow.Context) (*batch.DeregisterJobDefinitionOutput, error) {
    var output batch.DeregisterJobDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BatchDescribeComputeEnvironmentsResult struct {
	Result workflow.Future
}

func (r *BatchDescribeComputeEnvironmentsResult) Get(ctx workflow.Context) (*batch.DescribeComputeEnvironmentsOutput, error) {
    var output batch.DescribeComputeEnvironmentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BatchDescribeJobDefinitionsResult struct {
	Result workflow.Future
}

func (r *BatchDescribeJobDefinitionsResult) Get(ctx workflow.Context) (*batch.DescribeJobDefinitionsOutput, error) {
    var output batch.DescribeJobDefinitionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BatchDescribeJobQueuesResult struct {
	Result workflow.Future
}

func (r *BatchDescribeJobQueuesResult) Get(ctx workflow.Context) (*batch.DescribeJobQueuesOutput, error) {
    var output batch.DescribeJobQueuesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BatchDescribeJobsResult struct {
	Result workflow.Future
}

func (r *BatchDescribeJobsResult) Get(ctx workflow.Context) (*batch.DescribeJobsOutput, error) {
    var output batch.DescribeJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BatchListJobsResult struct {
	Result workflow.Future
}

func (r *BatchListJobsResult) Get(ctx workflow.Context) (*batch.ListJobsOutput, error) {
    var output batch.ListJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BatchRegisterJobDefinitionResult struct {
	Result workflow.Future
}

func (r *BatchRegisterJobDefinitionResult) Get(ctx workflow.Context) (*batch.RegisterJobDefinitionOutput, error) {
    var output batch.RegisterJobDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BatchSubmitJobResult struct {
	Result workflow.Future
}

func (r *BatchSubmitJobResult) Get(ctx workflow.Context) (*batch.SubmitJobOutput, error) {
    var output batch.SubmitJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BatchTerminateJobResult struct {
	Result workflow.Future
}

func (r *BatchTerminateJobResult) Get(ctx workflow.Context) (*batch.TerminateJobOutput, error) {
    var output batch.TerminateJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BatchUpdateComputeEnvironmentResult struct {
	Result workflow.Future
}

func (r *BatchUpdateComputeEnvironmentResult) Get(ctx workflow.Context) (*batch.UpdateComputeEnvironmentOutput, error) {
    var output batch.UpdateComputeEnvironmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BatchUpdateJobQueueResult struct {
	Result workflow.Future
}

func (r *BatchUpdateJobQueueResult) Get(ctx workflow.Context) (*batch.UpdateJobQueueOutput, error) {
    var output batch.UpdateJobQueueOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BatchStub struct {
    activities awsactivities.BatchActivities
}

func NewBatchStub() BatchClient {
    return &BatchStub{}
}

func (a *BatchStub) CancelJob(ctx workflow.Context, input *batch.CancelJobInput) (*batch.CancelJobOutput, error) {
    var output batch.CancelJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelJob, input).Get(ctx, &output)
    return &output, err
}

func (a *BatchStub) CancelJobAsync(ctx workflow.Context, input *batch.CancelJobInput) *BatchCancelJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelJob, input)
    return &BatchCancelJobResult{Result: future}
}

func (a *BatchStub) CreateComputeEnvironment(ctx workflow.Context, input *batch.CreateComputeEnvironmentInput) (*batch.CreateComputeEnvironmentOutput, error) {
    var output batch.CreateComputeEnvironmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateComputeEnvironment, input).Get(ctx, &output)
    return &output, err
}

func (a *BatchStub) CreateComputeEnvironmentAsync(ctx workflow.Context, input *batch.CreateComputeEnvironmentInput) *BatchCreateComputeEnvironmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateComputeEnvironment, input)
    return &BatchCreateComputeEnvironmentResult{Result: future}
}

func (a *BatchStub) CreateJobQueue(ctx workflow.Context, input *batch.CreateJobQueueInput) (*batch.CreateJobQueueOutput, error) {
    var output batch.CreateJobQueueOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateJobQueue, input).Get(ctx, &output)
    return &output, err
}

func (a *BatchStub) CreateJobQueueAsync(ctx workflow.Context, input *batch.CreateJobQueueInput) *BatchCreateJobQueueResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateJobQueue, input)
    return &BatchCreateJobQueueResult{Result: future}
}

func (a *BatchStub) DeleteComputeEnvironment(ctx workflow.Context, input *batch.DeleteComputeEnvironmentInput) (*batch.DeleteComputeEnvironmentOutput, error) {
    var output batch.DeleteComputeEnvironmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteComputeEnvironment, input).Get(ctx, &output)
    return &output, err
}

func (a *BatchStub) DeleteComputeEnvironmentAsync(ctx workflow.Context, input *batch.DeleteComputeEnvironmentInput) *BatchDeleteComputeEnvironmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteComputeEnvironment, input)
    return &BatchDeleteComputeEnvironmentResult{Result: future}
}

func (a *BatchStub) DeleteJobQueue(ctx workflow.Context, input *batch.DeleteJobQueueInput) (*batch.DeleteJobQueueOutput, error) {
    var output batch.DeleteJobQueueOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteJobQueue, input).Get(ctx, &output)
    return &output, err
}

func (a *BatchStub) DeleteJobQueueAsync(ctx workflow.Context, input *batch.DeleteJobQueueInput) *BatchDeleteJobQueueResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteJobQueue, input)
    return &BatchDeleteJobQueueResult{Result: future}
}

func (a *BatchStub) DeregisterJobDefinition(ctx workflow.Context, input *batch.DeregisterJobDefinitionInput) (*batch.DeregisterJobDefinitionOutput, error) {
    var output batch.DeregisterJobDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterJobDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *BatchStub) DeregisterJobDefinitionAsync(ctx workflow.Context, input *batch.DeregisterJobDefinitionInput) *BatchDeregisterJobDefinitionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeregisterJobDefinition, input)
    return &BatchDeregisterJobDefinitionResult{Result: future}
}

func (a *BatchStub) DescribeComputeEnvironments(ctx workflow.Context, input *batch.DescribeComputeEnvironmentsInput) (*batch.DescribeComputeEnvironmentsOutput, error) {
    var output batch.DescribeComputeEnvironmentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeComputeEnvironments, input).Get(ctx, &output)
    return &output, err
}

func (a *BatchStub) DescribeComputeEnvironmentsAsync(ctx workflow.Context, input *batch.DescribeComputeEnvironmentsInput) *BatchDescribeComputeEnvironmentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeComputeEnvironments, input)
    return &BatchDescribeComputeEnvironmentsResult{Result: future}
}

func (a *BatchStub) DescribeJobDefinitions(ctx workflow.Context, input *batch.DescribeJobDefinitionsInput) (*batch.DescribeJobDefinitionsOutput, error) {
    var output batch.DescribeJobDefinitionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeJobDefinitions, input).Get(ctx, &output)
    return &output, err
}

func (a *BatchStub) DescribeJobDefinitionsAsync(ctx workflow.Context, input *batch.DescribeJobDefinitionsInput) *BatchDescribeJobDefinitionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeJobDefinitions, input)
    return &BatchDescribeJobDefinitionsResult{Result: future}
}

func (a *BatchStub) DescribeJobQueues(ctx workflow.Context, input *batch.DescribeJobQueuesInput) (*batch.DescribeJobQueuesOutput, error) {
    var output batch.DescribeJobQueuesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeJobQueues, input).Get(ctx, &output)
    return &output, err
}

func (a *BatchStub) DescribeJobQueuesAsync(ctx workflow.Context, input *batch.DescribeJobQueuesInput) *BatchDescribeJobQueuesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeJobQueues, input)
    return &BatchDescribeJobQueuesResult{Result: future}
}

func (a *BatchStub) DescribeJobs(ctx workflow.Context, input *batch.DescribeJobsInput) (*batch.DescribeJobsOutput, error) {
    var output batch.DescribeJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *BatchStub) DescribeJobsAsync(ctx workflow.Context, input *batch.DescribeJobsInput) *BatchDescribeJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeJobs, input)
    return &BatchDescribeJobsResult{Result: future}
}

func (a *BatchStub) ListJobs(ctx workflow.Context, input *batch.ListJobsInput) (*batch.ListJobsOutput, error) {
    var output batch.ListJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *BatchStub) ListJobsAsync(ctx workflow.Context, input *batch.ListJobsInput) *BatchListJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListJobs, input)
    return &BatchListJobsResult{Result: future}
}

func (a *BatchStub) RegisterJobDefinition(ctx workflow.Context, input *batch.RegisterJobDefinitionInput) (*batch.RegisterJobDefinitionOutput, error) {
    var output batch.RegisterJobDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterJobDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *BatchStub) RegisterJobDefinitionAsync(ctx workflow.Context, input *batch.RegisterJobDefinitionInput) *BatchRegisterJobDefinitionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterJobDefinition, input)
    return &BatchRegisterJobDefinitionResult{Result: future}
}

func (a *BatchStub) SubmitJob(ctx workflow.Context, input *batch.SubmitJobInput) (*batch.SubmitJobOutput, error) {
    var output batch.SubmitJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SubmitJob, input).Get(ctx, &output)
    return &output, err
}

func (a *BatchStub) SubmitJobAsync(ctx workflow.Context, input *batch.SubmitJobInput) *BatchSubmitJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SubmitJob, input)
    return &BatchSubmitJobResult{Result: future}
}

func (a *BatchStub) TerminateJob(ctx workflow.Context, input *batch.TerminateJobInput) (*batch.TerminateJobOutput, error) {
    var output batch.TerminateJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TerminateJob, input).Get(ctx, &output)
    return &output, err
}

func (a *BatchStub) TerminateJobAsync(ctx workflow.Context, input *batch.TerminateJobInput) *BatchTerminateJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TerminateJob, input)
    return &BatchTerminateJobResult{Result: future}
}

func (a *BatchStub) UpdateComputeEnvironment(ctx workflow.Context, input *batch.UpdateComputeEnvironmentInput) (*batch.UpdateComputeEnvironmentOutput, error) {
    var output batch.UpdateComputeEnvironmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateComputeEnvironment, input).Get(ctx, &output)
    return &output, err
}

func (a *BatchStub) UpdateComputeEnvironmentAsync(ctx workflow.Context, input *batch.UpdateComputeEnvironmentInput) *BatchUpdateComputeEnvironmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateComputeEnvironment, input)
    return &BatchUpdateComputeEnvironmentResult{Result: future}
}

func (a *BatchStub) UpdateJobQueue(ctx workflow.Context, input *batch.UpdateJobQueueInput) (*batch.UpdateJobQueueOutput, error) {
    var output batch.UpdateJobQueueOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateJobQueue, input).Get(ctx, &output)
    return &output, err
}

func (a *BatchStub) UpdateJobQueueAsync(ctx workflow.Context, input *batch.UpdateJobQueueInput) *BatchUpdateJobQueueResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateJobQueue, input)
    return &BatchUpdateJobQueueResult{Result: future}
}
