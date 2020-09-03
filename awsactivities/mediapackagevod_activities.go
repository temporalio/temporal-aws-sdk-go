package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/mediapackagevod"
	"github.com/aws/aws-sdk-go/service/mediapackagevod/mediapackagevodiface"
)

type MediaPackageVodActivities struct {
	client mediapackagevodiface.MediaPackageVodAPI
}

func NewMediaPackageVodActivities(client mediapackagevodiface.MediaPackageVodAPI) *MediaPackageVodActivities {
    return &MediaPackageVodActivities{client: client}
}

func (a *MediaPackageVodActivities) CreateAsset(input *mediapackagevod.CreateAssetInput) (*mediapackagevod.CreateAssetOutput, error) {
    return a.client.CreateAsset(input)
}

func (a *MediaPackageVodActivities) CreatePackagingConfiguration(input *mediapackagevod.CreatePackagingConfigurationInput) (*mediapackagevod.CreatePackagingConfigurationOutput, error) {
    return a.client.CreatePackagingConfiguration(input)
}

func (a *MediaPackageVodActivities) CreatePackagingGroup(input *mediapackagevod.CreatePackagingGroupInput) (*mediapackagevod.CreatePackagingGroupOutput, error) {
    return a.client.CreatePackagingGroup(input)
}

func (a *MediaPackageVodActivities) DeleteAsset(input *mediapackagevod.DeleteAssetInput) (*mediapackagevod.DeleteAssetOutput, error) {
    return a.client.DeleteAsset(input)
}

func (a *MediaPackageVodActivities) DeletePackagingConfiguration(input *mediapackagevod.DeletePackagingConfigurationInput) (*mediapackagevod.DeletePackagingConfigurationOutput, error) {
    return a.client.DeletePackagingConfiguration(input)
}

func (a *MediaPackageVodActivities) DeletePackagingGroup(input *mediapackagevod.DeletePackagingGroupInput) (*mediapackagevod.DeletePackagingGroupOutput, error) {
    return a.client.DeletePackagingGroup(input)
}

func (a *MediaPackageVodActivities) DescribeAsset(input *mediapackagevod.DescribeAssetInput) (*mediapackagevod.DescribeAssetOutput, error) {
    return a.client.DescribeAsset(input)
}

func (a *MediaPackageVodActivities) DescribePackagingConfiguration(input *mediapackagevod.DescribePackagingConfigurationInput) (*mediapackagevod.DescribePackagingConfigurationOutput, error) {
    return a.client.DescribePackagingConfiguration(input)
}

func (a *MediaPackageVodActivities) DescribePackagingGroup(input *mediapackagevod.DescribePackagingGroupInput) (*mediapackagevod.DescribePackagingGroupOutput, error) {
    return a.client.DescribePackagingGroup(input)
}

func (a *MediaPackageVodActivities) ListAssets(input *mediapackagevod.ListAssetsInput) (*mediapackagevod.ListAssetsOutput, error) {
    return a.client.ListAssets(input)
}

func (a *MediaPackageVodActivities) ListPackagingConfigurations(input *mediapackagevod.ListPackagingConfigurationsInput) (*mediapackagevod.ListPackagingConfigurationsOutput, error) {
    return a.client.ListPackagingConfigurations(input)
}

func (a *MediaPackageVodActivities) ListPackagingGroups(input *mediapackagevod.ListPackagingGroupsInput) (*mediapackagevod.ListPackagingGroupsOutput, error) {
    return a.client.ListPackagingGroups(input)
}

func (a *MediaPackageVodActivities) ListTagsForResource(input *mediapackagevod.ListTagsForResourceInput) (*mediapackagevod.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *MediaPackageVodActivities) TagResource(input *mediapackagevod.TagResourceInput) (*mediapackagevod.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *MediaPackageVodActivities) UntagResource(input *mediapackagevod.UntagResourceInput) (*mediapackagevod.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *MediaPackageVodActivities) UpdatePackagingGroup(input *mediapackagevod.UpdatePackagingGroupInput) (*mediapackagevod.UpdatePackagingGroupOutput, error) {
    return a.client.UpdatePackagingGroup(input)
}
