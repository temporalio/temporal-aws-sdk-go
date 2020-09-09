package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/efs"
	"github.com/aws/aws-sdk-go/service/efs/efsiface"
)

type EFSActivities struct {
	client efsiface.EFSAPI
}

func NewEFSActivities(session *session.Session, config ...*aws.Config) *EFSActivities {
	client := efs.New(session, config...)
	return &EFSActivities{client: client}
}

func (a *EFSActivities) CreateAccessPoint(input *efs.CreateAccessPointInput) (*efs.CreateAccessPointOutput, error) {
	return a.client.CreateAccessPoint(input)
}

func (a *EFSActivities) CreateFileSystem(input *efs.CreateFileSystemInput) (*efs.FileSystemDescription, error) {
	return a.client.CreateFileSystem(input)
}

func (a *EFSActivities) CreateMountTarget(input *efs.CreateMountTargetInput) (*efs.MountTargetDescription, error) {
	return a.client.CreateMountTarget(input)
}

func (a *EFSActivities) CreateTags(input *efs.CreateTagsInput) (*efs.CreateTagsOutput, error) {
	return a.client.CreateTags(input)
}

func (a *EFSActivities) DeleteAccessPoint(input *efs.DeleteAccessPointInput) (*efs.DeleteAccessPointOutput, error) {
	return a.client.DeleteAccessPoint(input)
}

func (a *EFSActivities) DeleteFileSystem(input *efs.DeleteFileSystemInput) (*efs.DeleteFileSystemOutput, error) {
	return a.client.DeleteFileSystem(input)
}

func (a *EFSActivities) DeleteFileSystemPolicy(input *efs.DeleteFileSystemPolicyInput) (*efs.DeleteFileSystemPolicyOutput, error) {
	return a.client.DeleteFileSystemPolicy(input)
}

func (a *EFSActivities) DeleteMountTarget(input *efs.DeleteMountTargetInput) (*efs.DeleteMountTargetOutput, error) {
	return a.client.DeleteMountTarget(input)
}

func (a *EFSActivities) DeleteTags(input *efs.DeleteTagsInput) (*efs.DeleteTagsOutput, error) {
	return a.client.DeleteTags(input)
}

func (a *EFSActivities) DescribeAccessPoints(input *efs.DescribeAccessPointsInput) (*efs.DescribeAccessPointsOutput, error) {
	return a.client.DescribeAccessPoints(input)
}

func (a *EFSActivities) DescribeBackupPolicy(input *efs.DescribeBackupPolicyInput) (*efs.DescribeBackupPolicyOutput, error) {
	return a.client.DescribeBackupPolicy(input)
}

func (a *EFSActivities) DescribeFileSystemPolicy(input *efs.DescribeFileSystemPolicyInput) (*efs.DescribeFileSystemPolicyOutput, error) {
	return a.client.DescribeFileSystemPolicy(input)
}

func (a *EFSActivities) DescribeFileSystems(input *efs.DescribeFileSystemsInput) (*efs.DescribeFileSystemsOutput, error) {
	return a.client.DescribeFileSystems(input)
}

func (a *EFSActivities) DescribeLifecycleConfiguration(input *efs.DescribeLifecycleConfigurationInput) (*efs.DescribeLifecycleConfigurationOutput, error) {
	return a.client.DescribeLifecycleConfiguration(input)
}

func (a *EFSActivities) DescribeMountTargetSecurityGroups(input *efs.DescribeMountTargetSecurityGroupsInput) (*efs.DescribeMountTargetSecurityGroupsOutput, error) {
	return a.client.DescribeMountTargetSecurityGroups(input)
}

func (a *EFSActivities) DescribeMountTargets(input *efs.DescribeMountTargetsInput) (*efs.DescribeMountTargetsOutput, error) {
	return a.client.DescribeMountTargets(input)
}

func (a *EFSActivities) DescribeTags(input *efs.DescribeTagsInput) (*efs.DescribeTagsOutput, error) {
	return a.client.DescribeTags(input)
}

func (a *EFSActivities) ListTagsForResource(input *efs.ListTagsForResourceInput) (*efs.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResource(input)
}

func (a *EFSActivities) ModifyMountTargetSecurityGroups(input *efs.ModifyMountTargetSecurityGroupsInput) (*efs.ModifyMountTargetSecurityGroupsOutput, error) {
	return a.client.ModifyMountTargetSecurityGroups(input)
}

func (a *EFSActivities) PutBackupPolicy(input *efs.PutBackupPolicyInput) (*efs.PutBackupPolicyOutput, error) {
	return a.client.PutBackupPolicy(input)
}

func (a *EFSActivities) PutFileSystemPolicy(input *efs.PutFileSystemPolicyInput) (*efs.PutFileSystemPolicyOutput, error) {
	return a.client.PutFileSystemPolicy(input)
}

func (a *EFSActivities) PutLifecycleConfiguration(input *efs.PutLifecycleConfigurationInput) (*efs.PutLifecycleConfigurationOutput, error) {
	return a.client.PutLifecycleConfiguration(input)
}

func (a *EFSActivities) TagResource(input *efs.TagResourceInput) (*efs.TagResourceOutput, error) {
	return a.client.TagResource(input)
}

func (a *EFSActivities) UntagResource(input *efs.UntagResourceInput) (*efs.UntagResourceOutput, error) {
	return a.client.UntagResource(input)
}

func (a *EFSActivities) UpdateFileSystem(input *efs.UpdateFileSystemInput) (*efs.UpdateFileSystemOutput, error) {
	return a.client.UpdateFileSystem(input)
}
