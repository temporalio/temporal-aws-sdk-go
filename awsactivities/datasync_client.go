package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/datasync"
	"go.temporal.io/sdk/workflow"
)

type DataSyncClient interface {
    CancelTaskExecution(ctx workflow.Context, input *datasync.CancelTaskExecutionInput) (*datasync.CancelTaskExecutionOutput, error)
    CancelTaskExecutionAsync(ctx workflow.Context, input *datasync.CancelTaskExecutionInput) *DatasyncCancelTaskExecutionResult

    CreateAgent(ctx workflow.Context, input *datasync.CreateAgentInput) (*datasync.CreateAgentOutput, error)
    CreateAgentAsync(ctx workflow.Context, input *datasync.CreateAgentInput) *DatasyncCreateAgentResult

    CreateLocationEfs(ctx workflow.Context, input *datasync.CreateLocationEfsInput) (*datasync.CreateLocationEfsOutput, error)
    CreateLocationEfsAsync(ctx workflow.Context, input *datasync.CreateLocationEfsInput) *DatasyncCreateLocationEfsResult

    CreateLocationFsxWindows(ctx workflow.Context, input *datasync.CreateLocationFsxWindowsInput) (*datasync.CreateLocationFsxWindowsOutput, error)
    CreateLocationFsxWindowsAsync(ctx workflow.Context, input *datasync.CreateLocationFsxWindowsInput) *DatasyncCreateLocationFsxWindowsResult

    CreateLocationNfs(ctx workflow.Context, input *datasync.CreateLocationNfsInput) (*datasync.CreateLocationNfsOutput, error)
    CreateLocationNfsAsync(ctx workflow.Context, input *datasync.CreateLocationNfsInput) *DatasyncCreateLocationNfsResult

    CreateLocationObjectStorage(ctx workflow.Context, input *datasync.CreateLocationObjectStorageInput) (*datasync.CreateLocationObjectStorageOutput, error)
    CreateLocationObjectStorageAsync(ctx workflow.Context, input *datasync.CreateLocationObjectStorageInput) *DatasyncCreateLocationObjectStorageResult

    CreateLocationS3(ctx workflow.Context, input *datasync.CreateLocationS3Input) (*datasync.CreateLocationS3Output, error)
    CreateLocationS3Async(ctx workflow.Context, input *datasync.CreateLocationS3Input) *DatasyncCreateLocationS3Result

    CreateLocationSmb(ctx workflow.Context, input *datasync.CreateLocationSmbInput) (*datasync.CreateLocationSmbOutput, error)
    CreateLocationSmbAsync(ctx workflow.Context, input *datasync.CreateLocationSmbInput) *DatasyncCreateLocationSmbResult

    CreateTask(ctx workflow.Context, input *datasync.CreateTaskInput) (*datasync.CreateTaskOutput, error)
    CreateTaskAsync(ctx workflow.Context, input *datasync.CreateTaskInput) *DatasyncCreateTaskResult

    DeleteAgent(ctx workflow.Context, input *datasync.DeleteAgentInput) (*datasync.DeleteAgentOutput, error)
    DeleteAgentAsync(ctx workflow.Context, input *datasync.DeleteAgentInput) *DatasyncDeleteAgentResult

    DeleteLocation(ctx workflow.Context, input *datasync.DeleteLocationInput) (*datasync.DeleteLocationOutput, error)
    DeleteLocationAsync(ctx workflow.Context, input *datasync.DeleteLocationInput) *DatasyncDeleteLocationResult

    DeleteTask(ctx workflow.Context, input *datasync.DeleteTaskInput) (*datasync.DeleteTaskOutput, error)
    DeleteTaskAsync(ctx workflow.Context, input *datasync.DeleteTaskInput) *DatasyncDeleteTaskResult

    DescribeAgent(ctx workflow.Context, input *datasync.DescribeAgentInput) (*datasync.DescribeAgentOutput, error)
    DescribeAgentAsync(ctx workflow.Context, input *datasync.DescribeAgentInput) *DatasyncDescribeAgentResult

    DescribeLocationEfs(ctx workflow.Context, input *datasync.DescribeLocationEfsInput) (*datasync.DescribeLocationEfsOutput, error)
    DescribeLocationEfsAsync(ctx workflow.Context, input *datasync.DescribeLocationEfsInput) *DatasyncDescribeLocationEfsResult

    DescribeLocationFsxWindows(ctx workflow.Context, input *datasync.DescribeLocationFsxWindowsInput) (*datasync.DescribeLocationFsxWindowsOutput, error)
    DescribeLocationFsxWindowsAsync(ctx workflow.Context, input *datasync.DescribeLocationFsxWindowsInput) *DatasyncDescribeLocationFsxWindowsResult

    DescribeLocationNfs(ctx workflow.Context, input *datasync.DescribeLocationNfsInput) (*datasync.DescribeLocationNfsOutput, error)
    DescribeLocationNfsAsync(ctx workflow.Context, input *datasync.DescribeLocationNfsInput) *DatasyncDescribeLocationNfsResult

    DescribeLocationObjectStorage(ctx workflow.Context, input *datasync.DescribeLocationObjectStorageInput) (*datasync.DescribeLocationObjectStorageOutput, error)
    DescribeLocationObjectStorageAsync(ctx workflow.Context, input *datasync.DescribeLocationObjectStorageInput) *DatasyncDescribeLocationObjectStorageResult

    DescribeLocationS3(ctx workflow.Context, input *datasync.DescribeLocationS3Input) (*datasync.DescribeLocationS3Output, error)
    DescribeLocationS3Async(ctx workflow.Context, input *datasync.DescribeLocationS3Input) *DatasyncDescribeLocationS3Result

    DescribeLocationSmb(ctx workflow.Context, input *datasync.DescribeLocationSmbInput) (*datasync.DescribeLocationSmbOutput, error)
    DescribeLocationSmbAsync(ctx workflow.Context, input *datasync.DescribeLocationSmbInput) *DatasyncDescribeLocationSmbResult

    DescribeTask(ctx workflow.Context, input *datasync.DescribeTaskInput) (*datasync.DescribeTaskOutput, error)
    DescribeTaskAsync(ctx workflow.Context, input *datasync.DescribeTaskInput) *DatasyncDescribeTaskResult

    DescribeTaskExecution(ctx workflow.Context, input *datasync.DescribeTaskExecutionInput) (*datasync.DescribeTaskExecutionOutput, error)
    DescribeTaskExecutionAsync(ctx workflow.Context, input *datasync.DescribeTaskExecutionInput) *DatasyncDescribeTaskExecutionResult

    ListAgents(ctx workflow.Context, input *datasync.ListAgentsInput) (*datasync.ListAgentsOutput, error)
    ListAgentsAsync(ctx workflow.Context, input *datasync.ListAgentsInput) *DatasyncListAgentsResult

    ListLocations(ctx workflow.Context, input *datasync.ListLocationsInput) (*datasync.ListLocationsOutput, error)
    ListLocationsAsync(ctx workflow.Context, input *datasync.ListLocationsInput) *DatasyncListLocationsResult

    ListTagsForResource(ctx workflow.Context, input *datasync.ListTagsForResourceInput) (*datasync.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *datasync.ListTagsForResourceInput) *DatasyncListTagsForResourceResult

    ListTaskExecutions(ctx workflow.Context, input *datasync.ListTaskExecutionsInput) (*datasync.ListTaskExecutionsOutput, error)
    ListTaskExecutionsAsync(ctx workflow.Context, input *datasync.ListTaskExecutionsInput) *DatasyncListTaskExecutionsResult

    ListTasks(ctx workflow.Context, input *datasync.ListTasksInput) (*datasync.ListTasksOutput, error)
    ListTasksAsync(ctx workflow.Context, input *datasync.ListTasksInput) *DatasyncListTasksResult

    StartTaskExecution(ctx workflow.Context, input *datasync.StartTaskExecutionInput) (*datasync.StartTaskExecutionOutput, error)
    StartTaskExecutionAsync(ctx workflow.Context, input *datasync.StartTaskExecutionInput) *DatasyncStartTaskExecutionResult

    TagResource(ctx workflow.Context, input *datasync.TagResourceInput) (*datasync.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *datasync.TagResourceInput) *DatasyncTagResourceResult

    UntagResource(ctx workflow.Context, input *datasync.UntagResourceInput) (*datasync.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *datasync.UntagResourceInput) *DatasyncUntagResourceResult

    UpdateAgent(ctx workflow.Context, input *datasync.UpdateAgentInput) (*datasync.UpdateAgentOutput, error)
    UpdateAgentAsync(ctx workflow.Context, input *datasync.UpdateAgentInput) *DatasyncUpdateAgentResult

    UpdateTask(ctx workflow.Context, input *datasync.UpdateTaskInput) (*datasync.UpdateTaskOutput, error)
    UpdateTaskAsync(ctx workflow.Context, input *datasync.UpdateTaskInput) *DatasyncUpdateTaskResult
}
type DatasyncCancelTaskExecutionResult struct {
	Result workflow.Future
}

func (r *DatasyncCancelTaskExecutionResult) Get(ctx workflow.Context) (*datasync.CancelTaskExecutionOutput, error) {
    var output datasync.CancelTaskExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncCreateAgentResult struct {
	Result workflow.Future
}

func (r *DatasyncCreateAgentResult) Get(ctx workflow.Context) (*datasync.CreateAgentOutput, error) {
    var output datasync.CreateAgentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncCreateLocationEfsResult struct {
	Result workflow.Future
}

func (r *DatasyncCreateLocationEfsResult) Get(ctx workflow.Context) (*datasync.CreateLocationEfsOutput, error) {
    var output datasync.CreateLocationEfsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncCreateLocationFsxWindowsResult struct {
	Result workflow.Future
}

func (r *DatasyncCreateLocationFsxWindowsResult) Get(ctx workflow.Context) (*datasync.CreateLocationFsxWindowsOutput, error) {
    var output datasync.CreateLocationFsxWindowsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncCreateLocationNfsResult struct {
	Result workflow.Future
}

func (r *DatasyncCreateLocationNfsResult) Get(ctx workflow.Context) (*datasync.CreateLocationNfsOutput, error) {
    var output datasync.CreateLocationNfsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncCreateLocationObjectStorageResult struct {
	Result workflow.Future
}

func (r *DatasyncCreateLocationObjectStorageResult) Get(ctx workflow.Context) (*datasync.CreateLocationObjectStorageOutput, error) {
    var output datasync.CreateLocationObjectStorageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncCreateLocationS3Result struct {
	Result workflow.Future
}

func (r *DatasyncCreateLocationS3Result) Get(ctx workflow.Context) (*datasync.CreateLocationS3Output, error) {
    var output datasync.CreateLocationS3Output
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncCreateLocationSmbResult struct {
	Result workflow.Future
}

func (r *DatasyncCreateLocationSmbResult) Get(ctx workflow.Context) (*datasync.CreateLocationSmbOutput, error) {
    var output datasync.CreateLocationSmbOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncCreateTaskResult struct {
	Result workflow.Future
}

func (r *DatasyncCreateTaskResult) Get(ctx workflow.Context) (*datasync.CreateTaskOutput, error) {
    var output datasync.CreateTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncDeleteAgentResult struct {
	Result workflow.Future
}

func (r *DatasyncDeleteAgentResult) Get(ctx workflow.Context) (*datasync.DeleteAgentOutput, error) {
    var output datasync.DeleteAgentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncDeleteLocationResult struct {
	Result workflow.Future
}

func (r *DatasyncDeleteLocationResult) Get(ctx workflow.Context) (*datasync.DeleteLocationOutput, error) {
    var output datasync.DeleteLocationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncDeleteTaskResult struct {
	Result workflow.Future
}

func (r *DatasyncDeleteTaskResult) Get(ctx workflow.Context) (*datasync.DeleteTaskOutput, error) {
    var output datasync.DeleteTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncDescribeAgentResult struct {
	Result workflow.Future
}

func (r *DatasyncDescribeAgentResult) Get(ctx workflow.Context) (*datasync.DescribeAgentOutput, error) {
    var output datasync.DescribeAgentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncDescribeLocationEfsResult struct {
	Result workflow.Future
}

func (r *DatasyncDescribeLocationEfsResult) Get(ctx workflow.Context) (*datasync.DescribeLocationEfsOutput, error) {
    var output datasync.DescribeLocationEfsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncDescribeLocationFsxWindowsResult struct {
	Result workflow.Future
}

func (r *DatasyncDescribeLocationFsxWindowsResult) Get(ctx workflow.Context) (*datasync.DescribeLocationFsxWindowsOutput, error) {
    var output datasync.DescribeLocationFsxWindowsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncDescribeLocationNfsResult struct {
	Result workflow.Future
}

func (r *DatasyncDescribeLocationNfsResult) Get(ctx workflow.Context) (*datasync.DescribeLocationNfsOutput, error) {
    var output datasync.DescribeLocationNfsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncDescribeLocationObjectStorageResult struct {
	Result workflow.Future
}

func (r *DatasyncDescribeLocationObjectStorageResult) Get(ctx workflow.Context) (*datasync.DescribeLocationObjectStorageOutput, error) {
    var output datasync.DescribeLocationObjectStorageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncDescribeLocationS3Result struct {
	Result workflow.Future
}

func (r *DatasyncDescribeLocationS3Result) Get(ctx workflow.Context) (*datasync.DescribeLocationS3Output, error) {
    var output datasync.DescribeLocationS3Output
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncDescribeLocationSmbResult struct {
	Result workflow.Future
}

func (r *DatasyncDescribeLocationSmbResult) Get(ctx workflow.Context) (*datasync.DescribeLocationSmbOutput, error) {
    var output datasync.DescribeLocationSmbOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncDescribeTaskResult struct {
	Result workflow.Future
}

func (r *DatasyncDescribeTaskResult) Get(ctx workflow.Context) (*datasync.DescribeTaskOutput, error) {
    var output datasync.DescribeTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncDescribeTaskExecutionResult struct {
	Result workflow.Future
}

func (r *DatasyncDescribeTaskExecutionResult) Get(ctx workflow.Context) (*datasync.DescribeTaskExecutionOutput, error) {
    var output datasync.DescribeTaskExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncListAgentsResult struct {
	Result workflow.Future
}

func (r *DatasyncListAgentsResult) Get(ctx workflow.Context) (*datasync.ListAgentsOutput, error) {
    var output datasync.ListAgentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncListLocationsResult struct {
	Result workflow.Future
}

func (r *DatasyncListLocationsResult) Get(ctx workflow.Context) (*datasync.ListLocationsOutput, error) {
    var output datasync.ListLocationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *DatasyncListTagsForResourceResult) Get(ctx workflow.Context) (*datasync.ListTagsForResourceOutput, error) {
    var output datasync.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncListTaskExecutionsResult struct {
	Result workflow.Future
}

func (r *DatasyncListTaskExecutionsResult) Get(ctx workflow.Context) (*datasync.ListTaskExecutionsOutput, error) {
    var output datasync.ListTaskExecutionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncListTasksResult struct {
	Result workflow.Future
}

func (r *DatasyncListTasksResult) Get(ctx workflow.Context) (*datasync.ListTasksOutput, error) {
    var output datasync.ListTasksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncStartTaskExecutionResult struct {
	Result workflow.Future
}

func (r *DatasyncStartTaskExecutionResult) Get(ctx workflow.Context) (*datasync.StartTaskExecutionOutput, error) {
    var output datasync.StartTaskExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncTagResourceResult struct {
	Result workflow.Future
}

func (r *DatasyncTagResourceResult) Get(ctx workflow.Context) (*datasync.TagResourceOutput, error) {
    var output datasync.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncUntagResourceResult struct {
	Result workflow.Future
}

func (r *DatasyncUntagResourceResult) Get(ctx workflow.Context) (*datasync.UntagResourceOutput, error) {
    var output datasync.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncUpdateAgentResult struct {
	Result workflow.Future
}

func (r *DatasyncUpdateAgentResult) Get(ctx workflow.Context) (*datasync.UpdateAgentOutput, error) {
    var output datasync.UpdateAgentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatasyncUpdateTaskResult struct {
	Result workflow.Future
}

func (r *DatasyncUpdateTaskResult) Get(ctx workflow.Context) (*datasync.UpdateTaskOutput, error) {
    var output datasync.UpdateTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type DataSyncStub struct {
    activities DataSyncClient
}

func NewDataSyncStub() DataSyncClient {
    return &DataSyncStub{}
}

func (a *DataSyncStub) CancelTaskExecution(ctx workflow.Context, input *datasync.CancelTaskExecutionInput) (*datasync.CancelTaskExecutionOutput, error) {
    var output datasync.CancelTaskExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelTaskExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) CreateAgent(ctx workflow.Context, input *datasync.CreateAgentInput) (*datasync.CreateAgentOutput, error) {
    var output datasync.CreateAgentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAgent, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) CreateLocationEfs(ctx workflow.Context, input *datasync.CreateLocationEfsInput) (*datasync.CreateLocationEfsOutput, error) {
    var output datasync.CreateLocationEfsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLocationEfs, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) CreateLocationFsxWindows(ctx workflow.Context, input *datasync.CreateLocationFsxWindowsInput) (*datasync.CreateLocationFsxWindowsOutput, error) {
    var output datasync.CreateLocationFsxWindowsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLocationFsxWindows, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) CreateLocationNfs(ctx workflow.Context, input *datasync.CreateLocationNfsInput) (*datasync.CreateLocationNfsOutput, error) {
    var output datasync.CreateLocationNfsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLocationNfs, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) CreateLocationObjectStorage(ctx workflow.Context, input *datasync.CreateLocationObjectStorageInput) (*datasync.CreateLocationObjectStorageOutput, error) {
    var output datasync.CreateLocationObjectStorageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLocationObjectStorage, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) CreateLocationS3(ctx workflow.Context, input *datasync.CreateLocationS3Input) (*datasync.CreateLocationS3Output, error) {
    var output datasync.CreateLocationS3Output
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLocationS3, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) CreateLocationSmb(ctx workflow.Context, input *datasync.CreateLocationSmbInput) (*datasync.CreateLocationSmbOutput, error) {
    var output datasync.CreateLocationSmbOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLocationSmb, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) CreateTask(ctx workflow.Context, input *datasync.CreateTaskInput) (*datasync.CreateTaskOutput, error) {
    var output datasync.CreateTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTask, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) DeleteAgent(ctx workflow.Context, input *datasync.DeleteAgentInput) (*datasync.DeleteAgentOutput, error) {
    var output datasync.DeleteAgentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAgent, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) DeleteLocation(ctx workflow.Context, input *datasync.DeleteLocationInput) (*datasync.DeleteLocationOutput, error) {
    var output datasync.DeleteLocationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteLocation, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) DeleteTask(ctx workflow.Context, input *datasync.DeleteTaskInput) (*datasync.DeleteTaskOutput, error) {
    var output datasync.DeleteTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTask, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) DescribeAgent(ctx workflow.Context, input *datasync.DescribeAgentInput) (*datasync.DescribeAgentOutput, error) {
    var output datasync.DescribeAgentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAgent, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) DescribeLocationEfs(ctx workflow.Context, input *datasync.DescribeLocationEfsInput) (*datasync.DescribeLocationEfsOutput, error) {
    var output datasync.DescribeLocationEfsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLocationEfs, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) DescribeLocationFsxWindows(ctx workflow.Context, input *datasync.DescribeLocationFsxWindowsInput) (*datasync.DescribeLocationFsxWindowsOutput, error) {
    var output datasync.DescribeLocationFsxWindowsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLocationFsxWindows, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) DescribeLocationNfs(ctx workflow.Context, input *datasync.DescribeLocationNfsInput) (*datasync.DescribeLocationNfsOutput, error) {
    var output datasync.DescribeLocationNfsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLocationNfs, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) DescribeLocationObjectStorage(ctx workflow.Context, input *datasync.DescribeLocationObjectStorageInput) (*datasync.DescribeLocationObjectStorageOutput, error) {
    var output datasync.DescribeLocationObjectStorageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLocationObjectStorage, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) DescribeLocationS3(ctx workflow.Context, input *datasync.DescribeLocationS3Input) (*datasync.DescribeLocationS3Output, error) {
    var output datasync.DescribeLocationS3Output
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLocationS3, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) DescribeLocationSmb(ctx workflow.Context, input *datasync.DescribeLocationSmbInput) (*datasync.DescribeLocationSmbOutput, error) {
    var output datasync.DescribeLocationSmbOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLocationSmb, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) DescribeTask(ctx workflow.Context, input *datasync.DescribeTaskInput) (*datasync.DescribeTaskOutput, error) {
    var output datasync.DescribeTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTask, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) DescribeTaskExecution(ctx workflow.Context, input *datasync.DescribeTaskExecutionInput) (*datasync.DescribeTaskExecutionOutput, error) {
    var output datasync.DescribeTaskExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTaskExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) ListAgents(ctx workflow.Context, input *datasync.ListAgentsInput) (*datasync.ListAgentsOutput, error) {
    var output datasync.ListAgentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAgents, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) ListLocations(ctx workflow.Context, input *datasync.ListLocationsInput) (*datasync.ListLocationsOutput, error) {
    var output datasync.ListLocationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListLocations, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) ListTagsForResource(ctx workflow.Context, input *datasync.ListTagsForResourceInput) (*datasync.ListTagsForResourceOutput, error) {
    var output datasync.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) ListTaskExecutions(ctx workflow.Context, input *datasync.ListTaskExecutionsInput) (*datasync.ListTaskExecutionsOutput, error) {
    var output datasync.ListTaskExecutionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTaskExecutions, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) ListTasks(ctx workflow.Context, input *datasync.ListTasksInput) (*datasync.ListTasksOutput, error) {
    var output datasync.ListTasksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTasks, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) StartTaskExecution(ctx workflow.Context, input *datasync.StartTaskExecutionInput) (*datasync.StartTaskExecutionOutput, error) {
    var output datasync.StartTaskExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartTaskExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) TagResource(ctx workflow.Context, input *datasync.TagResourceInput) (*datasync.TagResourceOutput, error) {
    var output datasync.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) UntagResource(ctx workflow.Context, input *datasync.UntagResourceInput) (*datasync.UntagResourceOutput, error) {
    var output datasync.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) UpdateAgent(ctx workflow.Context, input *datasync.UpdateAgentInput) (*datasync.UpdateAgentOutput, error) {
    var output datasync.UpdateAgentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAgent, input).Get(ctx, &output)
    return &output, err
}

func (a *DataSyncStub) UpdateTask(ctx workflow.Context, input *datasync.UpdateTaskInput) (*datasync.UpdateTaskOutput, error) {
    var output datasync.UpdateTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTask, input).Get(ctx, &output)
    return &output, err
}
