
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mediapackagevod"
	"github.com/aws/aws-sdk-go/service/mediapackagevod/mediapackagevodiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type MediaPackageVodActivities struct {
    client mediapackagevodiface.MediaPackageVodAPI
}

func NewMediaPackageVodActivities(session *session.Session, config ...*aws.Config) *MediaPackageVodActivities {
    client := mediapackagevod.New(session, config...)
    return &MediaPackageVodActivities{client: client}
}

func (a *MediaPackageVodActivities) CreateAsset(ctx context.Context, input *mediapackagevod.CreateAssetInput) (*mediapackagevod.CreateAssetOutput, error) {
    return a.client.CreateAssetWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) CreatePackagingConfiguration(ctx context.Context, input *mediapackagevod.CreatePackagingConfigurationInput) (*mediapackagevod.CreatePackagingConfigurationOutput, error) {
    return a.client.CreatePackagingConfigurationWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) CreatePackagingGroup(ctx context.Context, input *mediapackagevod.CreatePackagingGroupInput) (*mediapackagevod.CreatePackagingGroupOutput, error) {
    return a.client.CreatePackagingGroupWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) DeleteAsset(ctx context.Context, input *mediapackagevod.DeleteAssetInput) (*mediapackagevod.DeleteAssetOutput, error) {
    return a.client.DeleteAssetWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) DeletePackagingConfiguration(ctx context.Context, input *mediapackagevod.DeletePackagingConfigurationInput) (*mediapackagevod.DeletePackagingConfigurationOutput, error) {
    return a.client.DeletePackagingConfigurationWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) DeletePackagingGroup(ctx context.Context, input *mediapackagevod.DeletePackagingGroupInput) (*mediapackagevod.DeletePackagingGroupOutput, error) {
    return a.client.DeletePackagingGroupWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) DescribeAsset(ctx context.Context, input *mediapackagevod.DescribeAssetInput) (*mediapackagevod.DescribeAssetOutput, error) {
    return a.client.DescribeAssetWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) DescribePackagingConfiguration(ctx context.Context, input *mediapackagevod.DescribePackagingConfigurationInput) (*mediapackagevod.DescribePackagingConfigurationOutput, error) {
    return a.client.DescribePackagingConfigurationWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) DescribePackagingGroup(ctx context.Context, input *mediapackagevod.DescribePackagingGroupInput) (*mediapackagevod.DescribePackagingGroupOutput, error) {
    return a.client.DescribePackagingGroupWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) ListAssets(ctx context.Context, input *mediapackagevod.ListAssetsInput) (*mediapackagevod.ListAssetsOutput, error) {
    return a.client.ListAssetsWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) ListPackagingConfigurations(ctx context.Context, input *mediapackagevod.ListPackagingConfigurationsInput) (*mediapackagevod.ListPackagingConfigurationsOutput, error) {
    return a.client.ListPackagingConfigurationsWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) ListPackagingGroups(ctx context.Context, input *mediapackagevod.ListPackagingGroupsInput) (*mediapackagevod.ListPackagingGroupsOutput, error) {
    return a.client.ListPackagingGroupsWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) ListTagsForResource(ctx context.Context, input *mediapackagevod.ListTagsForResourceInput) (*mediapackagevod.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) TagResource(ctx context.Context, input *mediapackagevod.TagResourceInput) (*mediapackagevod.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) UntagResource(ctx context.Context, input *mediapackagevod.UntagResourceInput) (*mediapackagevod.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) UpdatePackagingGroup(ctx context.Context, input *mediapackagevod.UpdatePackagingGroupInput) (*mediapackagevod.UpdatePackagingGroupOutput, error) {
    return a.client.UpdatePackagingGroupWithContext(ctx, input)
}
