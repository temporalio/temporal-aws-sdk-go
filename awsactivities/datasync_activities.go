
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/datasync"
	"github.com/aws/aws-sdk-go/service/datasync/datasynciface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type DataSyncActivities struct {
    client datasynciface.DataSyncAPI
}

func NewDataSyncActivities(session *session.Session, config ...*aws.Config) *DataSyncActivities {
    client := datasync.New(session, config...)
    return &DataSyncActivities{client: client}
}

func (a *DataSyncActivities) CancelTaskExecution(ctx context.Context, input *datasync.CancelTaskExecutionInput) (*datasync.CancelTaskExecutionOutput, error) {
    return a.client.CancelTaskExecutionWithContext(ctx, input)
}

func (a *DataSyncActivities) CreateAgent(ctx context.Context, input *datasync.CreateAgentInput) (*datasync.CreateAgentOutput, error) {
    return a.client.CreateAgentWithContext(ctx, input)
}

func (a *DataSyncActivities) CreateLocationEfs(ctx context.Context, input *datasync.CreateLocationEfsInput) (*datasync.CreateLocationEfsOutput, error) {
    return a.client.CreateLocationEfsWithContext(ctx, input)
}

func (a *DataSyncActivities) CreateLocationFsxWindows(ctx context.Context, input *datasync.CreateLocationFsxWindowsInput) (*datasync.CreateLocationFsxWindowsOutput, error) {
    return a.client.CreateLocationFsxWindowsWithContext(ctx, input)
}

func (a *DataSyncActivities) CreateLocationNfs(ctx context.Context, input *datasync.CreateLocationNfsInput) (*datasync.CreateLocationNfsOutput, error) {
    return a.client.CreateLocationNfsWithContext(ctx, input)
}

func (a *DataSyncActivities) CreateLocationObjectStorage(ctx context.Context, input *datasync.CreateLocationObjectStorageInput) (*datasync.CreateLocationObjectStorageOutput, error) {
    return a.client.CreateLocationObjectStorageWithContext(ctx, input)
}

func (a *DataSyncActivities) CreateLocationS3(ctx context.Context, input *datasync.CreateLocationS3Input) (*datasync.CreateLocationS3Output, error) {
    return a.client.CreateLocationS3WithContext(ctx, input)
}

func (a *DataSyncActivities) CreateLocationSmb(ctx context.Context, input *datasync.CreateLocationSmbInput) (*datasync.CreateLocationSmbOutput, error) {
    return a.client.CreateLocationSmbWithContext(ctx, input)
}

func (a *DataSyncActivities) CreateTask(ctx context.Context, input *datasync.CreateTaskInput) (*datasync.CreateTaskOutput, error) {
    return a.client.CreateTaskWithContext(ctx, input)
}

func (a *DataSyncActivities) DeleteAgent(ctx context.Context, input *datasync.DeleteAgentInput) (*datasync.DeleteAgentOutput, error) {
    return a.client.DeleteAgentWithContext(ctx, input)
}

func (a *DataSyncActivities) DeleteLocation(ctx context.Context, input *datasync.DeleteLocationInput) (*datasync.DeleteLocationOutput, error) {
    return a.client.DeleteLocationWithContext(ctx, input)
}

func (a *DataSyncActivities) DeleteTask(ctx context.Context, input *datasync.DeleteTaskInput) (*datasync.DeleteTaskOutput, error) {
    return a.client.DeleteTaskWithContext(ctx, input)
}

func (a *DataSyncActivities) DescribeAgent(ctx context.Context, input *datasync.DescribeAgentInput) (*datasync.DescribeAgentOutput, error) {
    return a.client.DescribeAgentWithContext(ctx, input)
}

func (a *DataSyncActivities) DescribeLocationEfs(ctx context.Context, input *datasync.DescribeLocationEfsInput) (*datasync.DescribeLocationEfsOutput, error) {
    return a.client.DescribeLocationEfsWithContext(ctx, input)
}

func (a *DataSyncActivities) DescribeLocationFsxWindows(ctx context.Context, input *datasync.DescribeLocationFsxWindowsInput) (*datasync.DescribeLocationFsxWindowsOutput, error) {
    return a.client.DescribeLocationFsxWindowsWithContext(ctx, input)
}

func (a *DataSyncActivities) DescribeLocationNfs(ctx context.Context, input *datasync.DescribeLocationNfsInput) (*datasync.DescribeLocationNfsOutput, error) {
    return a.client.DescribeLocationNfsWithContext(ctx, input)
}

func (a *DataSyncActivities) DescribeLocationObjectStorage(ctx context.Context, input *datasync.DescribeLocationObjectStorageInput) (*datasync.DescribeLocationObjectStorageOutput, error) {
    return a.client.DescribeLocationObjectStorageWithContext(ctx, input)
}

func (a *DataSyncActivities) DescribeLocationS3(ctx context.Context, input *datasync.DescribeLocationS3Input) (*datasync.DescribeLocationS3Output, error) {
    return a.client.DescribeLocationS3WithContext(ctx, input)
}

func (a *DataSyncActivities) DescribeLocationSmb(ctx context.Context, input *datasync.DescribeLocationSmbInput) (*datasync.DescribeLocationSmbOutput, error) {
    return a.client.DescribeLocationSmbWithContext(ctx, input)
}

func (a *DataSyncActivities) DescribeTask(ctx context.Context, input *datasync.DescribeTaskInput) (*datasync.DescribeTaskOutput, error) {
    return a.client.DescribeTaskWithContext(ctx, input)
}

func (a *DataSyncActivities) DescribeTaskExecution(ctx context.Context, input *datasync.DescribeTaskExecutionInput) (*datasync.DescribeTaskExecutionOutput, error) {
    return a.client.DescribeTaskExecutionWithContext(ctx, input)
}

func (a *DataSyncActivities) ListAgents(ctx context.Context, input *datasync.ListAgentsInput) (*datasync.ListAgentsOutput, error) {
    return a.client.ListAgentsWithContext(ctx, input)
}

func (a *DataSyncActivities) ListLocations(ctx context.Context, input *datasync.ListLocationsInput) (*datasync.ListLocationsOutput, error) {
    return a.client.ListLocationsWithContext(ctx, input)
}

func (a *DataSyncActivities) ListTagsForResource(ctx context.Context, input *datasync.ListTagsForResourceInput) (*datasync.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *DataSyncActivities) ListTaskExecutions(ctx context.Context, input *datasync.ListTaskExecutionsInput) (*datasync.ListTaskExecutionsOutput, error) {
    return a.client.ListTaskExecutionsWithContext(ctx, input)
}

func (a *DataSyncActivities) ListTasks(ctx context.Context, input *datasync.ListTasksInput) (*datasync.ListTasksOutput, error) {
    return a.client.ListTasksWithContext(ctx, input)
}

func (a *DataSyncActivities) StartTaskExecution(ctx context.Context, input *datasync.StartTaskExecutionInput) (*datasync.StartTaskExecutionOutput, error) {
    return a.client.StartTaskExecutionWithContext(ctx, input)
}

func (a *DataSyncActivities) TagResource(ctx context.Context, input *datasync.TagResourceInput) (*datasync.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *DataSyncActivities) UntagResource(ctx context.Context, input *datasync.UntagResourceInput) (*datasync.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *DataSyncActivities) UpdateAgent(ctx context.Context, input *datasync.UpdateAgentInput) (*datasync.UpdateAgentOutput, error) {
    return a.client.UpdateAgentWithContext(ctx, input)
}

func (a *DataSyncActivities) UpdateTask(ctx context.Context, input *datasync.UpdateTaskInput) (*datasync.UpdateTaskOutput, error) {
    return a.client.UpdateTaskWithContext(ctx, input)
}
