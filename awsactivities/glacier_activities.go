
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glacier"
	"github.com/aws/aws-sdk-go/service/glacier/glacieriface"
)

type GlacierActivities struct {
	client glacieriface.GlacierAPI
}

func NewGlacierActivities(session *session.Session, config... *aws.Config) *GlacierActivities {
    client := glacier.New(session, config...)
    return &GlacierActivities{client: client}
}

func (a *GlacierActivities) AbortMultipartUpload(input *glacier.AbortMultipartUploadInput) (*glacier.AbortMultipartUploadOutput, error) {
    return a.client.AbortMultipartUpload(input)
}

func (a *GlacierActivities) AbortVaultLock(input *glacier.AbortVaultLockInput) (*glacier.AbortVaultLockOutput, error) {
    return a.client.AbortVaultLock(input)
}

func (a *GlacierActivities) AddTagsToVault(input *glacier.AddTagsToVaultInput) (*glacier.AddTagsToVaultOutput, error) {
    return a.client.AddTagsToVault(input)
}

func (a *GlacierActivities) CompleteMultipartUpload(input *glacier.CompleteMultipartUploadInput) (*glacier.ArchiveCreationOutput, error) {
    return a.client.CompleteMultipartUpload(input)
}

func (a *GlacierActivities) CompleteVaultLock(input *glacier.CompleteVaultLockInput) (*glacier.CompleteVaultLockOutput, error) {
    return a.client.CompleteVaultLock(input)
}

func (a *GlacierActivities) CreateVault(input *glacier.CreateVaultInput) (*glacier.CreateVaultOutput, error) {
    return a.client.CreateVault(input)
}

func (a *GlacierActivities) DeleteArchive(input *glacier.DeleteArchiveInput) (*glacier.DeleteArchiveOutput, error) {
    return a.client.DeleteArchive(input)
}

func (a *GlacierActivities) DeleteVault(input *glacier.DeleteVaultInput) (*glacier.DeleteVaultOutput, error) {
    return a.client.DeleteVault(input)
}

func (a *GlacierActivities) DeleteVaultAccessPolicy(input *glacier.DeleteVaultAccessPolicyInput) (*glacier.DeleteVaultAccessPolicyOutput, error) {
    return a.client.DeleteVaultAccessPolicy(input)
}

func (a *GlacierActivities) DeleteVaultNotifications(input *glacier.DeleteVaultNotificationsInput) (*glacier.DeleteVaultNotificationsOutput, error) {
    return a.client.DeleteVaultNotifications(input)
}

func (a *GlacierActivities) DescribeJob(input *glacier.DescribeJobInput) (*glacier.JobDescription, error) {
    return a.client.DescribeJob(input)
}

func (a *GlacierActivities) DescribeVault(input *glacier.DescribeVaultInput) (*glacier.DescribeVaultOutput, error) {
    return a.client.DescribeVault(input)
}

func (a *GlacierActivities) GetDataRetrievalPolicy(input *glacier.GetDataRetrievalPolicyInput) (*glacier.GetDataRetrievalPolicyOutput, error) {
    return a.client.GetDataRetrievalPolicy(input)
}

func (a *GlacierActivities) GetJobOutput(input *glacier.GetJobOutputInput) (*glacier.GetJobOutputOutput, error) {
    return a.client.GetJobOutput(input)
}

func (a *GlacierActivities) GetVaultAccessPolicy(input *glacier.GetVaultAccessPolicyInput) (*glacier.GetVaultAccessPolicyOutput, error) {
    return a.client.GetVaultAccessPolicy(input)
}

func (a *GlacierActivities) GetVaultLock(input *glacier.GetVaultLockInput) (*glacier.GetVaultLockOutput, error) {
    return a.client.GetVaultLock(input)
}

func (a *GlacierActivities) GetVaultNotifications(input *glacier.GetVaultNotificationsInput) (*glacier.GetVaultNotificationsOutput, error) {
    return a.client.GetVaultNotifications(input)
}

func (a *GlacierActivities) InitiateJob(input *glacier.InitiateJobInput) (*glacier.InitiateJobOutput, error) {
    return a.client.InitiateJob(input)
}

func (a *GlacierActivities) InitiateMultipartUpload(input *glacier.InitiateMultipartUploadInput) (*glacier.InitiateMultipartUploadOutput, error) {
    return a.client.InitiateMultipartUpload(input)
}

func (a *GlacierActivities) InitiateVaultLock(input *glacier.InitiateVaultLockInput) (*glacier.InitiateVaultLockOutput, error) {
    return a.client.InitiateVaultLock(input)
}

func (a *GlacierActivities) ListJobs(input *glacier.ListJobsInput) (*glacier.ListJobsOutput, error) {
    return a.client.ListJobs(input)
}

func (a *GlacierActivities) ListMultipartUploads(input *glacier.ListMultipartUploadsInput) (*glacier.ListMultipartUploadsOutput, error) {
    return a.client.ListMultipartUploads(input)
}

func (a *GlacierActivities) ListParts(input *glacier.ListPartsInput) (*glacier.ListPartsOutput, error) {
    return a.client.ListParts(input)
}

func (a *GlacierActivities) ListProvisionedCapacity(input *glacier.ListProvisionedCapacityInput) (*glacier.ListProvisionedCapacityOutput, error) {
    return a.client.ListProvisionedCapacity(input)
}

func (a *GlacierActivities) ListTagsForVault(input *glacier.ListTagsForVaultInput) (*glacier.ListTagsForVaultOutput, error) {
    return a.client.ListTagsForVault(input)
}

func (a *GlacierActivities) ListVaults(input *glacier.ListVaultsInput) (*glacier.ListVaultsOutput, error) {
    return a.client.ListVaults(input)
}

func (a *GlacierActivities) PurchaseProvisionedCapacity(input *glacier.PurchaseProvisionedCapacityInput) (*glacier.PurchaseProvisionedCapacityOutput, error) {
    return a.client.PurchaseProvisionedCapacity(input)
}

func (a *GlacierActivities) RemoveTagsFromVault(input *glacier.RemoveTagsFromVaultInput) (*glacier.RemoveTagsFromVaultOutput, error) {
    return a.client.RemoveTagsFromVault(input)
}

func (a *GlacierActivities) SetDataRetrievalPolicy(input *glacier.SetDataRetrievalPolicyInput) (*glacier.SetDataRetrievalPolicyOutput, error) {
    return a.client.SetDataRetrievalPolicy(input)
}

func (a *GlacierActivities) SetVaultAccessPolicy(input *glacier.SetVaultAccessPolicyInput) (*glacier.SetVaultAccessPolicyOutput, error) {
    return a.client.SetVaultAccessPolicy(input)
}

func (a *GlacierActivities) SetVaultNotifications(input *glacier.SetVaultNotificationsInput) (*glacier.SetVaultNotificationsOutput, error) {
    return a.client.SetVaultNotifications(input)
}

func (a *GlacierActivities) UploadArchive(input *glacier.UploadArchiveInput) (*glacier.ArchiveCreationOutput, error) {
    return a.client.UploadArchive(input)
}

func (a *GlacierActivities) UploadMultipartPart(input *glacier.UploadMultipartPartInput) (*glacier.UploadMultipartPartOutput, error) {
    return a.client.UploadMultipartPart(input)
}

func (a *GlacierActivities) WaitUntilVaultExists(input *glacier.DescribeVaultInput) error {
    return a.client.WaitUntilVaultExists(input)
}

func (a *GlacierActivities) WaitUntilVaultNotExists(input *glacier.DescribeVaultInput) error {
    return a.client.WaitUntilVaultNotExists(input)
}
