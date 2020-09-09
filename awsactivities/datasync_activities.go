
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/datasync"
	"github.com/aws/aws-sdk-go/service/datasync/datasynciface"
)

type DataSyncActivities struct {
    client datasynciface.DataSyncAPI
}

func NewDataSyncActivities(session *session.Session, config ...*aws.Config) *DataSyncActivities {
    client := datasync.New(session, config...)
    return &DataSyncActivities{client: client}
}

func (a *DataSyncActivities) CancelTaskExecution(input *datasync.CancelTaskExecutionInput) (*datasync.CancelTaskExecutionOutput, error) {
    return a.client.CancelTaskExecution(input)
}

func (a *DataSyncActivities) CreateAgent(input *datasync.CreateAgentInput) (*datasync.CreateAgentOutput, error) {
    return a.client.CreateAgent(input)
}

func (a *DataSyncActivities) CreateLocationEfs(input *datasync.CreateLocationEfsInput) (*datasync.CreateLocationEfsOutput, error) {
    return a.client.CreateLocationEfs(input)
}

func (a *DataSyncActivities) CreateLocationFsxWindows(input *datasync.CreateLocationFsxWindowsInput) (*datasync.CreateLocationFsxWindowsOutput, error) {
    return a.client.CreateLocationFsxWindows(input)
}

func (a *DataSyncActivities) CreateLocationNfs(input *datasync.CreateLocationNfsInput) (*datasync.CreateLocationNfsOutput, error) {
    return a.client.CreateLocationNfs(input)
}

func (a *DataSyncActivities) CreateLocationObjectStorage(input *datasync.CreateLocationObjectStorageInput) (*datasync.CreateLocationObjectStorageOutput, error) {
    return a.client.CreateLocationObjectStorage(input)
}

func (a *DataSyncActivities) CreateLocationS3(input *datasync.CreateLocationS3Input) (*datasync.CreateLocationS3Output, error) {
    return a.client.CreateLocationS3(input)
}

func (a *DataSyncActivities) CreateLocationSmb(input *datasync.CreateLocationSmbInput) (*datasync.CreateLocationSmbOutput, error) {
    return a.client.CreateLocationSmb(input)
}

func (a *DataSyncActivities) CreateTask(input *datasync.CreateTaskInput) (*datasync.CreateTaskOutput, error) {
    return a.client.CreateTask(input)
}

func (a *DataSyncActivities) DeleteAgent(input *datasync.DeleteAgentInput) (*datasync.DeleteAgentOutput, error) {
    return a.client.DeleteAgent(input)
}

func (a *DataSyncActivities) DeleteLocation(input *datasync.DeleteLocationInput) (*datasync.DeleteLocationOutput, error) {
    return a.client.DeleteLocation(input)
}

func (a *DataSyncActivities) DeleteTask(input *datasync.DeleteTaskInput) (*datasync.DeleteTaskOutput, error) {
    return a.client.DeleteTask(input)
}

func (a *DataSyncActivities) DescribeAgent(input *datasync.DescribeAgentInput) (*datasync.DescribeAgentOutput, error) {
    return a.client.DescribeAgent(input)
}

func (a *DataSyncActivities) DescribeLocationEfs(input *datasync.DescribeLocationEfsInput) (*datasync.DescribeLocationEfsOutput, error) {
    return a.client.DescribeLocationEfs(input)
}

func (a *DataSyncActivities) DescribeLocationFsxWindows(input *datasync.DescribeLocationFsxWindowsInput) (*datasync.DescribeLocationFsxWindowsOutput, error) {
    return a.client.DescribeLocationFsxWindows(input)
}

func (a *DataSyncActivities) DescribeLocationNfs(input *datasync.DescribeLocationNfsInput) (*datasync.DescribeLocationNfsOutput, error) {
    return a.client.DescribeLocationNfs(input)
}

func (a *DataSyncActivities) DescribeLocationObjectStorage(input *datasync.DescribeLocationObjectStorageInput) (*datasync.DescribeLocationObjectStorageOutput, error) {
    return a.client.DescribeLocationObjectStorage(input)
}

func (a *DataSyncActivities) DescribeLocationS3(input *datasync.DescribeLocationS3Input) (*datasync.DescribeLocationS3Output, error) {
    return a.client.DescribeLocationS3(input)
}

func (a *DataSyncActivities) DescribeLocationSmb(input *datasync.DescribeLocationSmbInput) (*datasync.DescribeLocationSmbOutput, error) {
    return a.client.DescribeLocationSmb(input)
}

func (a *DataSyncActivities) DescribeTask(input *datasync.DescribeTaskInput) (*datasync.DescribeTaskOutput, error) {
    return a.client.DescribeTask(input)
}

func (a *DataSyncActivities) DescribeTaskExecution(input *datasync.DescribeTaskExecutionInput) (*datasync.DescribeTaskExecutionOutput, error) {
    return a.client.DescribeTaskExecution(input)
}

func (a *DataSyncActivities) ListAgents(input *datasync.ListAgentsInput) (*datasync.ListAgentsOutput, error) {
    return a.client.ListAgents(input)
}

func (a *DataSyncActivities) ListLocations(input *datasync.ListLocationsInput) (*datasync.ListLocationsOutput, error) {
    return a.client.ListLocations(input)
}

func (a *DataSyncActivities) ListTagsForResource(input *datasync.ListTagsForResourceInput) (*datasync.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *DataSyncActivities) ListTaskExecutions(input *datasync.ListTaskExecutionsInput) (*datasync.ListTaskExecutionsOutput, error) {
    return a.client.ListTaskExecutions(input)
}

func (a *DataSyncActivities) ListTasks(input *datasync.ListTasksInput) (*datasync.ListTasksOutput, error) {
    return a.client.ListTasks(input)
}

func (a *DataSyncActivities) StartTaskExecution(input *datasync.StartTaskExecutionInput) (*datasync.StartTaskExecutionOutput, error) {
    return a.client.StartTaskExecution(input)
}

func (a *DataSyncActivities) TagResource(input *datasync.TagResourceInput) (*datasync.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *DataSyncActivities) UntagResource(input *datasync.UntagResourceInput) (*datasync.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *DataSyncActivities) UpdateAgent(input *datasync.UpdateAgentInput) (*datasync.UpdateAgentOutput, error) {
    return a.client.UpdateAgent(input)
}

func (a *DataSyncActivities) UpdateTask(input *datasync.UpdateTaskInput) (*datasync.UpdateTaskOutput, error) {
    return a.client.UpdateTask(input)
}
