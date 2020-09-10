package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glacier"
	"github.com/aws/aws-sdk-go/service/glacier/glacieriface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type GlacierActivities struct {
	client glacieriface.GlacierAPI
}

func NewGlacierActivities(session *session.Session, config ...*aws.Config) *GlacierActivities {
	client := glacier.New(session, config...)
	return &GlacierActivities{client: client}
}

func (a *GlacierActivities) AbortMultipartUpload(ctx context.Context, input *glacier.AbortMultipartUploadInput) (*glacier.AbortMultipartUploadOutput, error) {
	return a.client.AbortMultipartUploadWithContext(ctx, input)
}

func (a *GlacierActivities) AbortVaultLock(ctx context.Context, input *glacier.AbortVaultLockInput) (*glacier.AbortVaultLockOutput, error) {
	return a.client.AbortVaultLockWithContext(ctx, input)
}

func (a *GlacierActivities) AddTagsToVault(ctx context.Context, input *glacier.AddTagsToVaultInput) (*glacier.AddTagsToVaultOutput, error) {
	return a.client.AddTagsToVaultWithContext(ctx, input)
}

func (a *GlacierActivities) CompleteMultipartUpload(ctx context.Context, input *glacier.CompleteMultipartUploadInput) (*glacier.ArchiveCreationOutput, error) {
	return a.client.CompleteMultipartUploadWithContext(ctx, input)
}

func (a *GlacierActivities) CompleteVaultLock(ctx context.Context, input *glacier.CompleteVaultLockInput) (*glacier.CompleteVaultLockOutput, error) {
	return a.client.CompleteVaultLockWithContext(ctx, input)
}

func (a *GlacierActivities) CreateVault(ctx context.Context, input *glacier.CreateVaultInput) (*glacier.CreateVaultOutput, error) {
	return a.client.CreateVaultWithContext(ctx, input)
}

func (a *GlacierActivities) DeleteArchive(ctx context.Context, input *glacier.DeleteArchiveInput) (*glacier.DeleteArchiveOutput, error) {
	return a.client.DeleteArchiveWithContext(ctx, input)
}

func (a *GlacierActivities) DeleteVault(ctx context.Context, input *glacier.DeleteVaultInput) (*glacier.DeleteVaultOutput, error) {
	return a.client.DeleteVaultWithContext(ctx, input)
}

func (a *GlacierActivities) DeleteVaultAccessPolicy(ctx context.Context, input *glacier.DeleteVaultAccessPolicyInput) (*glacier.DeleteVaultAccessPolicyOutput, error) {
	return a.client.DeleteVaultAccessPolicyWithContext(ctx, input)
}

func (a *GlacierActivities) DeleteVaultNotifications(ctx context.Context, input *glacier.DeleteVaultNotificationsInput) (*glacier.DeleteVaultNotificationsOutput, error) {
	return a.client.DeleteVaultNotificationsWithContext(ctx, input)
}

func (a *GlacierActivities) DescribeJob(ctx context.Context, input *glacier.DescribeJobInput) (*glacier.JobDescription, error) {
	return a.client.DescribeJobWithContext(ctx, input)
}

func (a *GlacierActivities) DescribeVault(ctx context.Context, input *glacier.DescribeVaultInput) (*glacier.DescribeVaultOutput, error) {
	return a.client.DescribeVaultWithContext(ctx, input)
}

func (a *GlacierActivities) GetDataRetrievalPolicy(ctx context.Context, input *glacier.GetDataRetrievalPolicyInput) (*glacier.GetDataRetrievalPolicyOutput, error) {
	return a.client.GetDataRetrievalPolicyWithContext(ctx, input)
}

func (a *GlacierActivities) GetJobOutput(ctx context.Context, input *glacier.GetJobOutputInput) (*glacier.GetJobOutputOutput, error) {
	return a.client.GetJobOutputWithContext(ctx, input)
}

func (a *GlacierActivities) GetVaultAccessPolicy(ctx context.Context, input *glacier.GetVaultAccessPolicyInput) (*glacier.GetVaultAccessPolicyOutput, error) {
	return a.client.GetVaultAccessPolicyWithContext(ctx, input)
}

func (a *GlacierActivities) GetVaultLock(ctx context.Context, input *glacier.GetVaultLockInput) (*glacier.GetVaultLockOutput, error) {
	return a.client.GetVaultLockWithContext(ctx, input)
}

func (a *GlacierActivities) GetVaultNotifications(ctx context.Context, input *glacier.GetVaultNotificationsInput) (*glacier.GetVaultNotificationsOutput, error) {
	return a.client.GetVaultNotificationsWithContext(ctx, input)
}

func (a *GlacierActivities) InitiateJob(ctx context.Context, input *glacier.InitiateJobInput) (*glacier.InitiateJobOutput, error) {
	return a.client.InitiateJobWithContext(ctx, input)
}

func (a *GlacierActivities) InitiateMultipartUpload(ctx context.Context, input *glacier.InitiateMultipartUploadInput) (*glacier.InitiateMultipartUploadOutput, error) {
	return a.client.InitiateMultipartUploadWithContext(ctx, input)
}

func (a *GlacierActivities) InitiateVaultLock(ctx context.Context, input *glacier.InitiateVaultLockInput) (*glacier.InitiateVaultLockOutput, error) {
	return a.client.InitiateVaultLockWithContext(ctx, input)
}

func (a *GlacierActivities) ListJobs(ctx context.Context, input *glacier.ListJobsInput) (*glacier.ListJobsOutput, error) {
	return a.client.ListJobsWithContext(ctx, input)
}

func (a *GlacierActivities) ListMultipartUploads(ctx context.Context, input *glacier.ListMultipartUploadsInput) (*glacier.ListMultipartUploadsOutput, error) {
	return a.client.ListMultipartUploadsWithContext(ctx, input)
}

func (a *GlacierActivities) ListParts(ctx context.Context, input *glacier.ListPartsInput) (*glacier.ListPartsOutput, error) {
	return a.client.ListPartsWithContext(ctx, input)
}

func (a *GlacierActivities) ListProvisionedCapacity(ctx context.Context, input *glacier.ListProvisionedCapacityInput) (*glacier.ListProvisionedCapacityOutput, error) {
	return a.client.ListProvisionedCapacityWithContext(ctx, input)
}

func (a *GlacierActivities) ListTagsForVault(ctx context.Context, input *glacier.ListTagsForVaultInput) (*glacier.ListTagsForVaultOutput, error) {
	return a.client.ListTagsForVaultWithContext(ctx, input)
}

func (a *GlacierActivities) ListVaults(ctx context.Context, input *glacier.ListVaultsInput) (*glacier.ListVaultsOutput, error) {
	return a.client.ListVaultsWithContext(ctx, input)
}

func (a *GlacierActivities) PurchaseProvisionedCapacity(ctx context.Context, input *glacier.PurchaseProvisionedCapacityInput) (*glacier.PurchaseProvisionedCapacityOutput, error) {
	return a.client.PurchaseProvisionedCapacityWithContext(ctx, input)
}

func (a *GlacierActivities) RemoveTagsFromVault(ctx context.Context, input *glacier.RemoveTagsFromVaultInput) (*glacier.RemoveTagsFromVaultOutput, error) {
	return a.client.RemoveTagsFromVaultWithContext(ctx, input)
}

func (a *GlacierActivities) SetDataRetrievalPolicy(ctx context.Context, input *glacier.SetDataRetrievalPolicyInput) (*glacier.SetDataRetrievalPolicyOutput, error) {
	return a.client.SetDataRetrievalPolicyWithContext(ctx, input)
}

func (a *GlacierActivities) SetVaultAccessPolicy(ctx context.Context, input *glacier.SetVaultAccessPolicyInput) (*glacier.SetVaultAccessPolicyOutput, error) {
	return a.client.SetVaultAccessPolicyWithContext(ctx, input)
}

func (a *GlacierActivities) SetVaultNotifications(ctx context.Context, input *glacier.SetVaultNotificationsInput) (*glacier.SetVaultNotificationsOutput, error) {
	return a.client.SetVaultNotificationsWithContext(ctx, input)
}

func (a *GlacierActivities) UploadArchive(ctx context.Context, input *glacier.UploadArchiveInput) (*glacier.ArchiveCreationOutput, error) {
	return a.client.UploadArchiveWithContext(ctx, input)
}

func (a *GlacierActivities) UploadMultipartPart(ctx context.Context, input *glacier.UploadMultipartPartInput) (*glacier.UploadMultipartPartOutput, error) {
	return a.client.UploadMultipartPartWithContext(ctx, input)
}

func (a *GlacierActivities) WaitUntilVaultExists(ctx context.Context, input *glacier.DescribeVaultInput) error {
	return a.client.WaitUntilVaultExistsWithContext(ctx, input)

}

func (a *GlacierActivities) WaitUntilVaultNotExists(ctx context.Context, input *glacier.DescribeVaultInput) error {
	return a.client.WaitUntilVaultNotExistsWithContext(ctx, input)

}
