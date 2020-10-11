// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package mediapackagevodstub

import (
	"github.com/aws/aws-sdk-go/service/mediapackagevod"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	CreateAsset(ctx workflow.Context, input *mediapackagevod.CreateAssetInput) (*mediapackagevod.CreateAssetOutput, error)
	CreateAssetAsync(ctx workflow.Context, input *mediapackagevod.CreateAssetInput) *MediaPackageVodCreateAssetFuture

	CreatePackagingConfiguration(ctx workflow.Context, input *mediapackagevod.CreatePackagingConfigurationInput) (*mediapackagevod.CreatePackagingConfigurationOutput, error)
	CreatePackagingConfigurationAsync(ctx workflow.Context, input *mediapackagevod.CreatePackagingConfigurationInput) *MediaPackageVodCreatePackagingConfigurationFuture

	CreatePackagingGroup(ctx workflow.Context, input *mediapackagevod.CreatePackagingGroupInput) (*mediapackagevod.CreatePackagingGroupOutput, error)
	CreatePackagingGroupAsync(ctx workflow.Context, input *mediapackagevod.CreatePackagingGroupInput) *MediaPackageVodCreatePackagingGroupFuture

	DeleteAsset(ctx workflow.Context, input *mediapackagevod.DeleteAssetInput) (*mediapackagevod.DeleteAssetOutput, error)
	DeleteAssetAsync(ctx workflow.Context, input *mediapackagevod.DeleteAssetInput) *MediaPackageVodDeleteAssetFuture

	DeletePackagingConfiguration(ctx workflow.Context, input *mediapackagevod.DeletePackagingConfigurationInput) (*mediapackagevod.DeletePackagingConfigurationOutput, error)
	DeletePackagingConfigurationAsync(ctx workflow.Context, input *mediapackagevod.DeletePackagingConfigurationInput) *MediaPackageVodDeletePackagingConfigurationFuture

	DeletePackagingGroup(ctx workflow.Context, input *mediapackagevod.DeletePackagingGroupInput) (*mediapackagevod.DeletePackagingGroupOutput, error)
	DeletePackagingGroupAsync(ctx workflow.Context, input *mediapackagevod.DeletePackagingGroupInput) *MediaPackageVodDeletePackagingGroupFuture

	DescribeAsset(ctx workflow.Context, input *mediapackagevod.DescribeAssetInput) (*mediapackagevod.DescribeAssetOutput, error)
	DescribeAssetAsync(ctx workflow.Context, input *mediapackagevod.DescribeAssetInput) *MediaPackageVodDescribeAssetFuture

	DescribePackagingConfiguration(ctx workflow.Context, input *mediapackagevod.DescribePackagingConfigurationInput) (*mediapackagevod.DescribePackagingConfigurationOutput, error)
	DescribePackagingConfigurationAsync(ctx workflow.Context, input *mediapackagevod.DescribePackagingConfigurationInput) *MediaPackageVodDescribePackagingConfigurationFuture

	DescribePackagingGroup(ctx workflow.Context, input *mediapackagevod.DescribePackagingGroupInput) (*mediapackagevod.DescribePackagingGroupOutput, error)
	DescribePackagingGroupAsync(ctx workflow.Context, input *mediapackagevod.DescribePackagingGroupInput) *MediaPackageVodDescribePackagingGroupFuture

	ListAssets(ctx workflow.Context, input *mediapackagevod.ListAssetsInput) (*mediapackagevod.ListAssetsOutput, error)
	ListAssetsAsync(ctx workflow.Context, input *mediapackagevod.ListAssetsInput) *MediaPackageVodListAssetsFuture

	ListPackagingConfigurations(ctx workflow.Context, input *mediapackagevod.ListPackagingConfigurationsInput) (*mediapackagevod.ListPackagingConfigurationsOutput, error)
	ListPackagingConfigurationsAsync(ctx workflow.Context, input *mediapackagevod.ListPackagingConfigurationsInput) *MediaPackageVodListPackagingConfigurationsFuture

	ListPackagingGroups(ctx workflow.Context, input *mediapackagevod.ListPackagingGroupsInput) (*mediapackagevod.ListPackagingGroupsOutput, error)
	ListPackagingGroupsAsync(ctx workflow.Context, input *mediapackagevod.ListPackagingGroupsInput) *MediaPackageVodListPackagingGroupsFuture

	ListTagsForResource(ctx workflow.Context, input *mediapackagevod.ListTagsForResourceInput) (*mediapackagevod.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *mediapackagevod.ListTagsForResourceInput) *MediaPackageVodListTagsForResourceFuture

	TagResource(ctx workflow.Context, input *mediapackagevod.TagResourceInput) (*mediapackagevod.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *mediapackagevod.TagResourceInput) *MediaPackageVodTagResourceFuture

	UntagResource(ctx workflow.Context, input *mediapackagevod.UntagResourceInput) (*mediapackagevod.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *mediapackagevod.UntagResourceInput) *MediaPackageVodUntagResourceFuture

	UpdatePackagingGroup(ctx workflow.Context, input *mediapackagevod.UpdatePackagingGroupInput) (*mediapackagevod.UpdatePackagingGroupOutput, error)
	UpdatePackagingGroupAsync(ctx workflow.Context, input *mediapackagevod.UpdatePackagingGroupInput) *MediaPackageVodUpdatePackagingGroupFuture
}

func NewClient() Client {
	return &stub{}
}
