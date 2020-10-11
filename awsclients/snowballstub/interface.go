// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package snowballstub

import (
	"github.com/aws/aws-sdk-go/service/snowball"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	CancelCluster(ctx workflow.Context, input *snowball.CancelClusterInput) (*snowball.CancelClusterOutput, error)
	CancelClusterAsync(ctx workflow.Context, input *snowball.CancelClusterInput) *SnowballCancelClusterFuture

	CancelJob(ctx workflow.Context, input *snowball.CancelJobInput) (*snowball.CancelJobOutput, error)
	CancelJobAsync(ctx workflow.Context, input *snowball.CancelJobInput) *SnowballCancelJobFuture

	CreateAddress(ctx workflow.Context, input *snowball.CreateAddressInput) (*snowball.CreateAddressOutput, error)
	CreateAddressAsync(ctx workflow.Context, input *snowball.CreateAddressInput) *SnowballCreateAddressFuture

	CreateCluster(ctx workflow.Context, input *snowball.CreateClusterInput) (*snowball.CreateClusterOutput, error)
	CreateClusterAsync(ctx workflow.Context, input *snowball.CreateClusterInput) *SnowballCreateClusterFuture

	CreateJob(ctx workflow.Context, input *snowball.CreateJobInput) (*snowball.CreateJobOutput, error)
	CreateJobAsync(ctx workflow.Context, input *snowball.CreateJobInput) *SnowballCreateJobFuture

	CreateReturnShippingLabel(ctx workflow.Context, input *snowball.CreateReturnShippingLabelInput) (*snowball.CreateReturnShippingLabelOutput, error)
	CreateReturnShippingLabelAsync(ctx workflow.Context, input *snowball.CreateReturnShippingLabelInput) *SnowballCreateReturnShippingLabelFuture

	DescribeAddress(ctx workflow.Context, input *snowball.DescribeAddressInput) (*snowball.DescribeAddressOutput, error)
	DescribeAddressAsync(ctx workflow.Context, input *snowball.DescribeAddressInput) *SnowballDescribeAddressFuture

	DescribeAddresses(ctx workflow.Context, input *snowball.DescribeAddressesInput) (*snowball.DescribeAddressesOutput, error)
	DescribeAddressesAsync(ctx workflow.Context, input *snowball.DescribeAddressesInput) *SnowballDescribeAddressesFuture

	DescribeCluster(ctx workflow.Context, input *snowball.DescribeClusterInput) (*snowball.DescribeClusterOutput, error)
	DescribeClusterAsync(ctx workflow.Context, input *snowball.DescribeClusterInput) *SnowballDescribeClusterFuture

	DescribeJob(ctx workflow.Context, input *snowball.DescribeJobInput) (*snowball.DescribeJobOutput, error)
	DescribeJobAsync(ctx workflow.Context, input *snowball.DescribeJobInput) *SnowballDescribeJobFuture

	DescribeReturnShippingLabel(ctx workflow.Context, input *snowball.DescribeReturnShippingLabelInput) (*snowball.DescribeReturnShippingLabelOutput, error)
	DescribeReturnShippingLabelAsync(ctx workflow.Context, input *snowball.DescribeReturnShippingLabelInput) *SnowballDescribeReturnShippingLabelFuture

	GetJobManifest(ctx workflow.Context, input *snowball.GetJobManifestInput) (*snowball.GetJobManifestOutput, error)
	GetJobManifestAsync(ctx workflow.Context, input *snowball.GetJobManifestInput) *SnowballGetJobManifestFuture

	GetJobUnlockCode(ctx workflow.Context, input *snowball.GetJobUnlockCodeInput) (*snowball.GetJobUnlockCodeOutput, error)
	GetJobUnlockCodeAsync(ctx workflow.Context, input *snowball.GetJobUnlockCodeInput) *SnowballGetJobUnlockCodeFuture

	GetSnowballUsage(ctx workflow.Context, input *snowball.GetSnowballUsageInput) (*snowball.GetSnowballUsageOutput, error)
	GetSnowballUsageAsync(ctx workflow.Context, input *snowball.GetSnowballUsageInput) *SnowballGetSnowballUsageFuture

	GetSoftwareUpdates(ctx workflow.Context, input *snowball.GetSoftwareUpdatesInput) (*snowball.GetSoftwareUpdatesOutput, error)
	GetSoftwareUpdatesAsync(ctx workflow.Context, input *snowball.GetSoftwareUpdatesInput) *SnowballGetSoftwareUpdatesFuture

	ListClusterJobs(ctx workflow.Context, input *snowball.ListClusterJobsInput) (*snowball.ListClusterJobsOutput, error)
	ListClusterJobsAsync(ctx workflow.Context, input *snowball.ListClusterJobsInput) *SnowballListClusterJobsFuture

	ListClusters(ctx workflow.Context, input *snowball.ListClustersInput) (*snowball.ListClustersOutput, error)
	ListClustersAsync(ctx workflow.Context, input *snowball.ListClustersInput) *SnowballListClustersFuture

	ListCompatibleImages(ctx workflow.Context, input *snowball.ListCompatibleImagesInput) (*snowball.ListCompatibleImagesOutput, error)
	ListCompatibleImagesAsync(ctx workflow.Context, input *snowball.ListCompatibleImagesInput) *SnowballListCompatibleImagesFuture

	ListJobs(ctx workflow.Context, input *snowball.ListJobsInput) (*snowball.ListJobsOutput, error)
	ListJobsAsync(ctx workflow.Context, input *snowball.ListJobsInput) *SnowballListJobsFuture

	UpdateCluster(ctx workflow.Context, input *snowball.UpdateClusterInput) (*snowball.UpdateClusterOutput, error)
	UpdateClusterAsync(ctx workflow.Context, input *snowball.UpdateClusterInput) *SnowballUpdateClusterFuture

	UpdateJob(ctx workflow.Context, input *snowball.UpdateJobInput) (*snowball.UpdateJobOutput, error)
	UpdateJobAsync(ctx workflow.Context, input *snowball.UpdateJobInput) *SnowballUpdateJobFuture

	UpdateJobShipmentState(ctx workflow.Context, input *snowball.UpdateJobShipmentStateInput) (*snowball.UpdateJobShipmentStateOutput, error)
	UpdateJobShipmentStateAsync(ctx workflow.Context, input *snowball.UpdateJobShipmentStateInput) *SnowballUpdateJobShipmentStateFuture
}

func NewClient() Client {
	return &stub{}
}
