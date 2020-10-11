// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go/service/cloudhsmv2/cloudhsmv2iface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type CloudHSMV2Activities struct {
	client cloudhsmv2iface.CloudHSMV2API

	sessionFactory SessionFactory
}

func NewCloudHSMV2Activities(sess *session.Session, config ...*aws.Config) *CloudHSMV2Activities {
	client := cloudhsmv2.New(sess, config...)
	return &CloudHSMV2Activities{client: client}
}

func NewCloudHSMV2ActivitiesWithSessionFactory(sessionFactory SessionFactory) *CloudHSMV2Activities {
	return &CloudHSMV2Activities{sessionFactory: sessionFactory}
}

func (a *CloudHSMV2Activities) getClient(ctx context.Context) (cloudhsmv2iface.CloudHSMV2API, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return cloudhsmv2.New(sess), nil
}

func (a *CloudHSMV2Activities) CopyBackupToRegion(ctx context.Context, input *cloudhsmv2.CopyBackupToRegionInput) (*cloudhsmv2.CopyBackupToRegionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CopyBackupToRegionWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) CreateCluster(ctx context.Context, input *cloudhsmv2.CreateClusterInput) (*cloudhsmv2.CreateClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateClusterWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) CreateHsm(ctx context.Context, input *cloudhsmv2.CreateHsmInput) (*cloudhsmv2.CreateHsmOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateHsmWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) DeleteBackup(ctx context.Context, input *cloudhsmv2.DeleteBackupInput) (*cloudhsmv2.DeleteBackupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBackupWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) DeleteCluster(ctx context.Context, input *cloudhsmv2.DeleteClusterInput) (*cloudhsmv2.DeleteClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteClusterWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) DeleteHsm(ctx context.Context, input *cloudhsmv2.DeleteHsmInput) (*cloudhsmv2.DeleteHsmOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteHsmWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) DescribeBackups(ctx context.Context, input *cloudhsmv2.DescribeBackupsInput) (*cloudhsmv2.DescribeBackupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeBackupsWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) DescribeClusters(ctx context.Context, input *cloudhsmv2.DescribeClustersInput) (*cloudhsmv2.DescribeClustersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeClustersWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) InitializeCluster(ctx context.Context, input *cloudhsmv2.InitializeClusterInput) (*cloudhsmv2.InitializeClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.InitializeClusterWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) ListTags(ctx context.Context, input *cloudhsmv2.ListTagsInput) (*cloudhsmv2.ListTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) RestoreBackup(ctx context.Context, input *cloudhsmv2.RestoreBackupInput) (*cloudhsmv2.RestoreBackupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RestoreBackupWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) TagResource(ctx context.Context, input *cloudhsmv2.TagResourceInput) (*cloudhsmv2.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *CloudHSMV2Activities) UntagResource(ctx context.Context, input *cloudhsmv2.UntagResourceInput) (*cloudhsmv2.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}
