// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package kinesisanalyticsv2stub

import (
	"github.com/aws/aws-sdk-go/service/kinesisanalyticsv2"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

type Client interface {
	AddApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionOutput, error)
	AddApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionInput) *KinesisAnalyticsV2AddApplicationCloudWatchLoggingOptionFuture

	AddApplicationInput(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationInputInput) (*kinesisanalyticsv2.AddApplicationInputOutput, error)
	AddApplicationInputAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationInputInput) *KinesisAnalyticsV2AddApplicationInputFuture

	AddApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationInputProcessingConfigurationInput) (*kinesisanalyticsv2.AddApplicationInputProcessingConfigurationOutput, error)
	AddApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationInputProcessingConfigurationInput) *KinesisAnalyticsV2AddApplicationInputProcessingConfigurationFuture

	AddApplicationOutput(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationOutputInput) (*kinesisanalyticsv2.AddApplicationOutputOutput, error)
	AddApplicationOutputAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationOutputInput) *KinesisAnalyticsV2AddApplicationOutputFuture

	AddApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationReferenceDataSourceInput) (*kinesisanalyticsv2.AddApplicationReferenceDataSourceOutput, error)
	AddApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationReferenceDataSourceInput) *KinesisAnalyticsV2AddApplicationReferenceDataSourceFuture

	AddApplicationVpcConfiguration(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationVpcConfigurationInput) (*kinesisanalyticsv2.AddApplicationVpcConfigurationOutput, error)
	AddApplicationVpcConfigurationAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationVpcConfigurationInput) *KinesisAnalyticsV2AddApplicationVpcConfigurationFuture

	CreateApplication(ctx workflow.Context, input *kinesisanalyticsv2.CreateApplicationInput) (*kinesisanalyticsv2.CreateApplicationOutput, error)
	CreateApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.CreateApplicationInput) *KinesisAnalyticsV2CreateApplicationFuture

	CreateApplicationSnapshot(ctx workflow.Context, input *kinesisanalyticsv2.CreateApplicationSnapshotInput) (*kinesisanalyticsv2.CreateApplicationSnapshotOutput, error)
	CreateApplicationSnapshotAsync(ctx workflow.Context, input *kinesisanalyticsv2.CreateApplicationSnapshotInput) *KinesisAnalyticsV2CreateApplicationSnapshotFuture

	DeleteApplication(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationInput) (*kinesisanalyticsv2.DeleteApplicationOutput, error)
	DeleteApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationInput) *KinesisAnalyticsV2DeleteApplicationFuture

	DeleteApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionInput) (*kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionOutput, error)
	DeleteApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionInput) *KinesisAnalyticsV2DeleteApplicationCloudWatchLoggingOptionFuture

	DeleteApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationInput) (*kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationOutput, error)
	DeleteApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationInput) *KinesisAnalyticsV2DeleteApplicationInputProcessingConfigurationFuture

	DeleteApplicationOutput(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationOutputInput) (*kinesisanalyticsv2.DeleteApplicationOutputOutput, error)
	DeleteApplicationOutputAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationOutputInput) *KinesisAnalyticsV2DeleteApplicationOutputFuture

	DeleteApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationReferenceDataSourceInput) (*kinesisanalyticsv2.DeleteApplicationReferenceDataSourceOutput, error)
	DeleteApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationReferenceDataSourceInput) *KinesisAnalyticsV2DeleteApplicationReferenceDataSourceFuture

	DeleteApplicationSnapshot(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationSnapshotInput) (*kinesisanalyticsv2.DeleteApplicationSnapshotOutput, error)
	DeleteApplicationSnapshotAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationSnapshotInput) *KinesisAnalyticsV2DeleteApplicationSnapshotFuture

	DeleteApplicationVpcConfiguration(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationVpcConfigurationInput) (*kinesisanalyticsv2.DeleteApplicationVpcConfigurationOutput, error)
	DeleteApplicationVpcConfigurationAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationVpcConfigurationInput) *KinesisAnalyticsV2DeleteApplicationVpcConfigurationFuture

	DescribeApplication(ctx workflow.Context, input *kinesisanalyticsv2.DescribeApplicationInput) (*kinesisanalyticsv2.DescribeApplicationOutput, error)
	DescribeApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.DescribeApplicationInput) *KinesisAnalyticsV2DescribeApplicationFuture

	DescribeApplicationSnapshot(ctx workflow.Context, input *kinesisanalyticsv2.DescribeApplicationSnapshotInput) (*kinesisanalyticsv2.DescribeApplicationSnapshotOutput, error)
	DescribeApplicationSnapshotAsync(ctx workflow.Context, input *kinesisanalyticsv2.DescribeApplicationSnapshotInput) *KinesisAnalyticsV2DescribeApplicationSnapshotFuture

	DiscoverInputSchema(ctx workflow.Context, input *kinesisanalyticsv2.DiscoverInputSchemaInput) (*kinesisanalyticsv2.DiscoverInputSchemaOutput, error)
	DiscoverInputSchemaAsync(ctx workflow.Context, input *kinesisanalyticsv2.DiscoverInputSchemaInput) *KinesisAnalyticsV2DiscoverInputSchemaFuture

	ListApplicationSnapshots(ctx workflow.Context, input *kinesisanalyticsv2.ListApplicationSnapshotsInput) (*kinesisanalyticsv2.ListApplicationSnapshotsOutput, error)
	ListApplicationSnapshotsAsync(ctx workflow.Context, input *kinesisanalyticsv2.ListApplicationSnapshotsInput) *KinesisAnalyticsV2ListApplicationSnapshotsFuture

	ListApplications(ctx workflow.Context, input *kinesisanalyticsv2.ListApplicationsInput) (*kinesisanalyticsv2.ListApplicationsOutput, error)
	ListApplicationsAsync(ctx workflow.Context, input *kinesisanalyticsv2.ListApplicationsInput) *KinesisAnalyticsV2ListApplicationsFuture

	ListTagsForResource(ctx workflow.Context, input *kinesisanalyticsv2.ListTagsForResourceInput) (*kinesisanalyticsv2.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.ListTagsForResourceInput) *KinesisAnalyticsV2ListTagsForResourceFuture

	StartApplication(ctx workflow.Context, input *kinesisanalyticsv2.StartApplicationInput) (*kinesisanalyticsv2.StartApplicationOutput, error)
	StartApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.StartApplicationInput) *KinesisAnalyticsV2StartApplicationFuture

	StopApplication(ctx workflow.Context, input *kinesisanalyticsv2.StopApplicationInput) (*kinesisanalyticsv2.StopApplicationOutput, error)
	StopApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.StopApplicationInput) *KinesisAnalyticsV2StopApplicationFuture

	TagResource(ctx workflow.Context, input *kinesisanalyticsv2.TagResourceInput) (*kinesisanalyticsv2.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.TagResourceInput) *KinesisAnalyticsV2TagResourceFuture

	UntagResource(ctx workflow.Context, input *kinesisanalyticsv2.UntagResourceInput) (*kinesisanalyticsv2.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.UntagResourceInput) *KinesisAnalyticsV2UntagResourceFuture

	UpdateApplication(ctx workflow.Context, input *kinesisanalyticsv2.UpdateApplicationInput) (*kinesisanalyticsv2.UpdateApplicationOutput, error)
	UpdateApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.UpdateApplicationInput) *KinesisAnalyticsV2UpdateApplicationFuture
}

func NewClient() Client {
	return &stub{}
}

