// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package cloudhsmv2stub

import (
	"github.com/aws/aws-sdk-go/service/cloudhsmv2"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

type Client interface {
	CopyBackupToRegion(ctx workflow.Context, input *cloudhsmv2.CopyBackupToRegionInput) (*cloudhsmv2.CopyBackupToRegionOutput, error)
	CopyBackupToRegionAsync(ctx workflow.Context, input *cloudhsmv2.CopyBackupToRegionInput) *CloudHSMV2CopyBackupToRegionFuture

	CreateCluster(ctx workflow.Context, input *cloudhsmv2.CreateClusterInput) (*cloudhsmv2.CreateClusterOutput, error)
	CreateClusterAsync(ctx workflow.Context, input *cloudhsmv2.CreateClusterInput) *CloudHSMV2CreateClusterFuture

	CreateHsm(ctx workflow.Context, input *cloudhsmv2.CreateHsmInput) (*cloudhsmv2.CreateHsmOutput, error)
	CreateHsmAsync(ctx workflow.Context, input *cloudhsmv2.CreateHsmInput) *CloudHSMV2CreateHsmFuture

	DeleteBackup(ctx workflow.Context, input *cloudhsmv2.DeleteBackupInput) (*cloudhsmv2.DeleteBackupOutput, error)
	DeleteBackupAsync(ctx workflow.Context, input *cloudhsmv2.DeleteBackupInput) *CloudHSMV2DeleteBackupFuture

	DeleteCluster(ctx workflow.Context, input *cloudhsmv2.DeleteClusterInput) (*cloudhsmv2.DeleteClusterOutput, error)
	DeleteClusterAsync(ctx workflow.Context, input *cloudhsmv2.DeleteClusterInput) *CloudHSMV2DeleteClusterFuture

	DeleteHsm(ctx workflow.Context, input *cloudhsmv2.DeleteHsmInput) (*cloudhsmv2.DeleteHsmOutput, error)
	DeleteHsmAsync(ctx workflow.Context, input *cloudhsmv2.DeleteHsmInput) *CloudHSMV2DeleteHsmFuture

	DescribeBackups(ctx workflow.Context, input *cloudhsmv2.DescribeBackupsInput) (*cloudhsmv2.DescribeBackupsOutput, error)
	DescribeBackupsAsync(ctx workflow.Context, input *cloudhsmv2.DescribeBackupsInput) *CloudHSMV2DescribeBackupsFuture

	DescribeClusters(ctx workflow.Context, input *cloudhsmv2.DescribeClustersInput) (*cloudhsmv2.DescribeClustersOutput, error)
	DescribeClustersAsync(ctx workflow.Context, input *cloudhsmv2.DescribeClustersInput) *CloudHSMV2DescribeClustersFuture

	InitializeCluster(ctx workflow.Context, input *cloudhsmv2.InitializeClusterInput) (*cloudhsmv2.InitializeClusterOutput, error)
	InitializeClusterAsync(ctx workflow.Context, input *cloudhsmv2.InitializeClusterInput) *CloudHSMV2InitializeClusterFuture

	ListTags(ctx workflow.Context, input *cloudhsmv2.ListTagsInput) (*cloudhsmv2.ListTagsOutput, error)
	ListTagsAsync(ctx workflow.Context, input *cloudhsmv2.ListTagsInput) *CloudHSMV2ListTagsFuture

	RestoreBackup(ctx workflow.Context, input *cloudhsmv2.RestoreBackupInput) (*cloudhsmv2.RestoreBackupOutput, error)
	RestoreBackupAsync(ctx workflow.Context, input *cloudhsmv2.RestoreBackupInput) *CloudHSMV2RestoreBackupFuture

	TagResource(ctx workflow.Context, input *cloudhsmv2.TagResourceInput) (*cloudhsmv2.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *cloudhsmv2.TagResourceInput) *CloudHSMV2TagResourceFuture

	UntagResource(ctx workflow.Context, input *cloudhsmv2.UntagResourceInput) (*cloudhsmv2.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *cloudhsmv2.UntagResourceInput) *CloudHSMV2UntagResourceFuture
}

func NewClient() Client {
	return &stub{}
}

