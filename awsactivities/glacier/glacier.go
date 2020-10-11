// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package glacier

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glacier"
	"github.com/aws/aws-sdk-go/service/glacier/glacieriface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client glacieriface.GlacierAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := glacier.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (glacieriface.GlacierAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return glacier.New(sess), nil
}

func (a *Activities) AbortMultipartUpload(ctx context.Context, input *glacier.AbortMultipartUploadInput) (*glacier.AbortMultipartUploadOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AbortMultipartUploadWithContext(ctx, input)
}

func (a *Activities) AbortVaultLock(ctx context.Context, input *glacier.AbortVaultLockInput) (*glacier.AbortVaultLockOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AbortVaultLockWithContext(ctx, input)
}

func (a *Activities) AddTagsToVault(ctx context.Context, input *glacier.AddTagsToVaultInput) (*glacier.AddTagsToVaultOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddTagsToVaultWithContext(ctx, input)
}

func (a *Activities) CompleteMultipartUpload(ctx context.Context, input *glacier.CompleteMultipartUploadInput) (*glacier.ArchiveCreationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CompleteMultipartUploadWithContext(ctx, input)
}

func (a *Activities) CompleteVaultLock(ctx context.Context, input *glacier.CompleteVaultLockInput) (*glacier.CompleteVaultLockOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CompleteVaultLockWithContext(ctx, input)
}

func (a *Activities) CreateVault(ctx context.Context, input *glacier.CreateVaultInput) (*glacier.CreateVaultOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateVaultWithContext(ctx, input)
}

func (a *Activities) DeleteArchive(ctx context.Context, input *glacier.DeleteArchiveInput) (*glacier.DeleteArchiveOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteArchiveWithContext(ctx, input)
}

func (a *Activities) DeleteVault(ctx context.Context, input *glacier.DeleteVaultInput) (*glacier.DeleteVaultOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteVaultWithContext(ctx, input)
}

func (a *Activities) DeleteVaultAccessPolicy(ctx context.Context, input *glacier.DeleteVaultAccessPolicyInput) (*glacier.DeleteVaultAccessPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteVaultAccessPolicyWithContext(ctx, input)
}

func (a *Activities) DeleteVaultNotifications(ctx context.Context, input *glacier.DeleteVaultNotificationsInput) (*glacier.DeleteVaultNotificationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteVaultNotificationsWithContext(ctx, input)
}

func (a *Activities) DescribeJob(ctx context.Context, input *glacier.DescribeJobInput) (*glacier.JobDescription, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeJobWithContext(ctx, input)
}

func (a *Activities) DescribeVault(ctx context.Context, input *glacier.DescribeVaultInput) (*glacier.DescribeVaultOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeVaultWithContext(ctx, input)
}

func (a *Activities) GetDataRetrievalPolicy(ctx context.Context, input *glacier.GetDataRetrievalPolicyInput) (*glacier.GetDataRetrievalPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDataRetrievalPolicyWithContext(ctx, input)
}

func (a *Activities) GetJobOutput(ctx context.Context, input *glacier.GetJobOutputInput) (*glacier.GetJobOutputOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetJobOutputWithContext(ctx, input)
}

func (a *Activities) GetVaultAccessPolicy(ctx context.Context, input *glacier.GetVaultAccessPolicyInput) (*glacier.GetVaultAccessPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetVaultAccessPolicyWithContext(ctx, input)
}

func (a *Activities) GetVaultLock(ctx context.Context, input *glacier.GetVaultLockInput) (*glacier.GetVaultLockOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetVaultLockWithContext(ctx, input)
}

func (a *Activities) GetVaultNotifications(ctx context.Context, input *glacier.GetVaultNotificationsInput) (*glacier.GetVaultNotificationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetVaultNotificationsWithContext(ctx, input)
}

func (a *Activities) InitiateJob(ctx context.Context, input *glacier.InitiateJobInput) (*glacier.InitiateJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.InitiateJobWithContext(ctx, input)
}

func (a *Activities) InitiateMultipartUpload(ctx context.Context, input *glacier.InitiateMultipartUploadInput) (*glacier.InitiateMultipartUploadOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.InitiateMultipartUploadWithContext(ctx, input)
}

func (a *Activities) InitiateVaultLock(ctx context.Context, input *glacier.InitiateVaultLockInput) (*glacier.InitiateVaultLockOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.InitiateVaultLockWithContext(ctx, input)
}

func (a *Activities) ListJobs(ctx context.Context, input *glacier.ListJobsInput) (*glacier.ListJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListJobsWithContext(ctx, input)
}

func (a *Activities) ListMultipartUploads(ctx context.Context, input *glacier.ListMultipartUploadsInput) (*glacier.ListMultipartUploadsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListMultipartUploadsWithContext(ctx, input)
}

func (a *Activities) ListParts(ctx context.Context, input *glacier.ListPartsInput) (*glacier.ListPartsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPartsWithContext(ctx, input)
}

func (a *Activities) ListProvisionedCapacity(ctx context.Context, input *glacier.ListProvisionedCapacityInput) (*glacier.ListProvisionedCapacityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListProvisionedCapacityWithContext(ctx, input)
}

func (a *Activities) ListTagsForVault(ctx context.Context, input *glacier.ListTagsForVaultInput) (*glacier.ListTagsForVaultOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForVaultWithContext(ctx, input)
}

func (a *Activities) ListVaults(ctx context.Context, input *glacier.ListVaultsInput) (*glacier.ListVaultsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListVaultsWithContext(ctx, input)
}

func (a *Activities) PurchaseProvisionedCapacity(ctx context.Context, input *glacier.PurchaseProvisionedCapacityInput) (*glacier.PurchaseProvisionedCapacityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PurchaseProvisionedCapacityWithContext(ctx, input)
}

func (a *Activities) RemoveTagsFromVault(ctx context.Context, input *glacier.RemoveTagsFromVaultInput) (*glacier.RemoveTagsFromVaultOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RemoveTagsFromVaultWithContext(ctx, input)
}

func (a *Activities) SetDataRetrievalPolicy(ctx context.Context, input *glacier.SetDataRetrievalPolicyInput) (*glacier.SetDataRetrievalPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SetDataRetrievalPolicyWithContext(ctx, input)
}

func (a *Activities) SetVaultAccessPolicy(ctx context.Context, input *glacier.SetVaultAccessPolicyInput) (*glacier.SetVaultAccessPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SetVaultAccessPolicyWithContext(ctx, input)
}

func (a *Activities) SetVaultNotifications(ctx context.Context, input *glacier.SetVaultNotificationsInput) (*glacier.SetVaultNotificationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SetVaultNotificationsWithContext(ctx, input)
}

func (a *Activities) UploadArchive(ctx context.Context, input *glacier.UploadArchiveInput) (*glacier.ArchiveCreationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UploadArchiveWithContext(ctx, input)
}

func (a *Activities) UploadMultipartPart(ctx context.Context, input *glacier.UploadMultipartPartInput) (*glacier.UploadMultipartPartOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UploadMultipartPartWithContext(ctx, input)
}

func (a *Activities) WaitUntilVaultExists(ctx context.Context, input *glacier.DescribeVaultInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilVaultExistsWithContext(ctx, input, options...)
	})
}

func (a *Activities) WaitUntilVaultNotExists(ctx context.Context, input *glacier.DescribeVaultInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilVaultNotExistsWithContext(ctx, input, options...)
	})
}