package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/fsx"
	"github.com/aws/aws-sdk-go/service/fsx/fsxiface"
)

type FSxActivities struct {
	client fsxiface.FSxAPI
}

func NewFSxActivities(client fsxiface.FSxAPI) *FSxActivities {
    return &FSxActivities{client: client}
}

func (a *FSxActivities) CancelDataRepositoryTask(input *fsx.CancelDataRepositoryTaskInput) (*fsx.CancelDataRepositoryTaskOutput, error) {
    return a.client.CancelDataRepositoryTask(input)
}

func (a *FSxActivities) CreateBackup(input *fsx.CreateBackupInput) (*fsx.CreateBackupOutput, error) {
    return a.client.CreateBackup(input)
}

func (a *FSxActivities) CreateDataRepositoryTask(input *fsx.CreateDataRepositoryTaskInput) (*fsx.CreateDataRepositoryTaskOutput, error) {
    return a.client.CreateDataRepositoryTask(input)
}

func (a *FSxActivities) CreateFileSystem(input *fsx.CreateFileSystemInput) (*fsx.CreateFileSystemOutput, error) {
    return a.client.CreateFileSystem(input)
}

func (a *FSxActivities) CreateFileSystemFromBackup(input *fsx.CreateFileSystemFromBackupInput) (*fsx.CreateFileSystemFromBackupOutput, error) {
    return a.client.CreateFileSystemFromBackup(input)
}

func (a *FSxActivities) DeleteBackup(input *fsx.DeleteBackupInput) (*fsx.DeleteBackupOutput, error) {
    return a.client.DeleteBackup(input)
}

func (a *FSxActivities) DeleteFileSystem(input *fsx.DeleteFileSystemInput) (*fsx.DeleteFileSystemOutput, error) {
    return a.client.DeleteFileSystem(input)
}

func (a *FSxActivities) DescribeBackups(input *fsx.DescribeBackupsInput) (*fsx.DescribeBackupsOutput, error) {
    return a.client.DescribeBackups(input)
}

func (a *FSxActivities) DescribeDataRepositoryTasks(input *fsx.DescribeDataRepositoryTasksInput) (*fsx.DescribeDataRepositoryTasksOutput, error) {
    return a.client.DescribeDataRepositoryTasks(input)
}

func (a *FSxActivities) DescribeFileSystems(input *fsx.DescribeFileSystemsInput) (*fsx.DescribeFileSystemsOutput, error) {
    return a.client.DescribeFileSystems(input)
}

func (a *FSxActivities) ListTagsForResource(input *fsx.ListTagsForResourceInput) (*fsx.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *FSxActivities) TagResource(input *fsx.TagResourceInput) (*fsx.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *FSxActivities) UntagResource(input *fsx.UntagResourceInput) (*fsx.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *FSxActivities) UpdateFileSystem(input *fsx.UpdateFileSystemInput) (*fsx.UpdateFileSystemOutput, error) {
    return a.client.UpdateFileSystem(input)
}
