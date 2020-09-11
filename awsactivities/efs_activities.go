package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/efs"
	"github.com/aws/aws-sdk-go/service/efs/efsiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type EFSActivities struct {
	client efsiface.EFSAPI
}

func NewEFSActivities(session *session.Session, config ...*aws.Config) *EFSActivities {
	client := efs.New(session, config...)
	return &EFSActivities{client: client}
}

func (a *EFSActivities) CreateAccessPoint(ctx context.Context, input *efs.CreateAccessPointInput) (*efs.CreateAccessPointOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateAccessPointWithContext(ctx, input)
}

func (a *EFSActivities) CreateFileSystem(ctx context.Context, input *efs.CreateFileSystemInput) (*efs.FileSystemDescription, error) {
	return a.client.CreateFileSystemWithContext(ctx, input)
}

func (a *EFSActivities) CreateMountTarget(ctx context.Context, input *efs.CreateMountTargetInput) (*efs.MountTargetDescription, error) {
	return a.client.CreateMountTargetWithContext(ctx, input)
}

func (a *EFSActivities) CreateTags(ctx context.Context, input *efs.CreateTagsInput) (*efs.CreateTagsOutput, error) {
	return a.client.CreateTagsWithContext(ctx, input)
}

func (a *EFSActivities) DeleteAccessPoint(ctx context.Context, input *efs.DeleteAccessPointInput) (*efs.DeleteAccessPointOutput, error) {
	return a.client.DeleteAccessPointWithContext(ctx, input)
}

func (a *EFSActivities) DeleteFileSystem(ctx context.Context, input *efs.DeleteFileSystemInput) (*efs.DeleteFileSystemOutput, error) {
	return a.client.DeleteFileSystemWithContext(ctx, input)
}

func (a *EFSActivities) DeleteFileSystemPolicy(ctx context.Context, input *efs.DeleteFileSystemPolicyInput) (*efs.DeleteFileSystemPolicyOutput, error) {
	return a.client.DeleteFileSystemPolicyWithContext(ctx, input)
}

func (a *EFSActivities) DeleteMountTarget(ctx context.Context, input *efs.DeleteMountTargetInput) (*efs.DeleteMountTargetOutput, error) {
	return a.client.DeleteMountTargetWithContext(ctx, input)
}

func (a *EFSActivities) DeleteTags(ctx context.Context, input *efs.DeleteTagsInput) (*efs.DeleteTagsOutput, error) {
	return a.client.DeleteTagsWithContext(ctx, input)
}

func (a *EFSActivities) DescribeAccessPoints(ctx context.Context, input *efs.DescribeAccessPointsInput) (*efs.DescribeAccessPointsOutput, error) {
	return a.client.DescribeAccessPointsWithContext(ctx, input)
}

func (a *EFSActivities) DescribeBackupPolicy(ctx context.Context, input *efs.DescribeBackupPolicyInput) (*efs.DescribeBackupPolicyOutput, error) {
	return a.client.DescribeBackupPolicyWithContext(ctx, input)
}

func (a *EFSActivities) DescribeFileSystemPolicy(ctx context.Context, input *efs.DescribeFileSystemPolicyInput) (*efs.DescribeFileSystemPolicyOutput, error) {
	return a.client.DescribeFileSystemPolicyWithContext(ctx, input)
}

func (a *EFSActivities) DescribeFileSystems(ctx context.Context, input *efs.DescribeFileSystemsInput) (*efs.DescribeFileSystemsOutput, error) {
	return a.client.DescribeFileSystemsWithContext(ctx, input)
}

func (a *EFSActivities) DescribeLifecycleConfiguration(ctx context.Context, input *efs.DescribeLifecycleConfigurationInput) (*efs.DescribeLifecycleConfigurationOutput, error) {
	return a.client.DescribeLifecycleConfigurationWithContext(ctx, input)
}

func (a *EFSActivities) DescribeMountTargetSecurityGroups(ctx context.Context, input *efs.DescribeMountTargetSecurityGroupsInput) (*efs.DescribeMountTargetSecurityGroupsOutput, error) {
	return a.client.DescribeMountTargetSecurityGroupsWithContext(ctx, input)
}

func (a *EFSActivities) DescribeMountTargets(ctx context.Context, input *efs.DescribeMountTargetsInput) (*efs.DescribeMountTargetsOutput, error) {
	return a.client.DescribeMountTargetsWithContext(ctx, input)
}

func (a *EFSActivities) DescribeTags(ctx context.Context, input *efs.DescribeTagsInput) (*efs.DescribeTagsOutput, error) {
	return a.client.DescribeTagsWithContext(ctx, input)
}

func (a *EFSActivities) ListTagsForResource(ctx context.Context, input *efs.ListTagsForResourceInput) (*efs.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *EFSActivities) ModifyMountTargetSecurityGroups(ctx context.Context, input *efs.ModifyMountTargetSecurityGroupsInput) (*efs.ModifyMountTargetSecurityGroupsOutput, error) {
	return a.client.ModifyMountTargetSecurityGroupsWithContext(ctx, input)
}

func (a *EFSActivities) PutBackupPolicy(ctx context.Context, input *efs.PutBackupPolicyInput) (*efs.PutBackupPolicyOutput, error) {
	return a.client.PutBackupPolicyWithContext(ctx, input)
}

func (a *EFSActivities) PutFileSystemPolicy(ctx context.Context, input *efs.PutFileSystemPolicyInput) (*efs.PutFileSystemPolicyOutput, error) {
	return a.client.PutFileSystemPolicyWithContext(ctx, input)
}

func (a *EFSActivities) PutLifecycleConfiguration(ctx context.Context, input *efs.PutLifecycleConfigurationInput) (*efs.PutLifecycleConfigurationOutput, error) {
	return a.client.PutLifecycleConfigurationWithContext(ctx, input)
}

func (a *EFSActivities) TagResource(ctx context.Context, input *efs.TagResourceInput) (*efs.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *EFSActivities) UntagResource(ctx context.Context, input *efs.UntagResourceInput) (*efs.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *EFSActivities) UpdateFileSystem(ctx context.Context, input *efs.UpdateFileSystemInput) (*efs.UpdateFileSystemOutput, error) {
	return a.client.UpdateFileSystemWithContext(ctx, input)
}
