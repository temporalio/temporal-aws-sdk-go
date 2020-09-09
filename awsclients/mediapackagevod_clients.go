package awsclients

import (
	"github.com/aws/aws-sdk-go/service/mediapackagevod"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type MediaPackageVodClient interface {
       CreateAsset(ctx workflow.Context, input *mediapackagevod.CreateAssetInput) (*mediapackagevod.CreateAssetOutput, error)
       CreateAssetAsync(ctx workflow.Context, input *mediapackagevod.CreateAssetInput) *MediapackagevodCreateAssetResult

       CreatePackagingConfiguration(ctx workflow.Context, input *mediapackagevod.CreatePackagingConfigurationInput) (*mediapackagevod.CreatePackagingConfigurationOutput, error)
       CreatePackagingConfigurationAsync(ctx workflow.Context, input *mediapackagevod.CreatePackagingConfigurationInput) *MediapackagevodCreatePackagingConfigurationResult

       CreatePackagingGroup(ctx workflow.Context, input *mediapackagevod.CreatePackagingGroupInput) (*mediapackagevod.CreatePackagingGroupOutput, error)
       CreatePackagingGroupAsync(ctx workflow.Context, input *mediapackagevod.CreatePackagingGroupInput) *MediapackagevodCreatePackagingGroupResult

       DeleteAsset(ctx workflow.Context, input *mediapackagevod.DeleteAssetInput) (*mediapackagevod.DeleteAssetOutput, error)
       DeleteAssetAsync(ctx workflow.Context, input *mediapackagevod.DeleteAssetInput) *MediapackagevodDeleteAssetResult

       DeletePackagingConfiguration(ctx workflow.Context, input *mediapackagevod.DeletePackagingConfigurationInput) (*mediapackagevod.DeletePackagingConfigurationOutput, error)
       DeletePackagingConfigurationAsync(ctx workflow.Context, input *mediapackagevod.DeletePackagingConfigurationInput) *MediapackagevodDeletePackagingConfigurationResult

       DeletePackagingGroup(ctx workflow.Context, input *mediapackagevod.DeletePackagingGroupInput) (*mediapackagevod.DeletePackagingGroupOutput, error)
       DeletePackagingGroupAsync(ctx workflow.Context, input *mediapackagevod.DeletePackagingGroupInput) *MediapackagevodDeletePackagingGroupResult

       DescribeAsset(ctx workflow.Context, input *mediapackagevod.DescribeAssetInput) (*mediapackagevod.DescribeAssetOutput, error)
       DescribeAssetAsync(ctx workflow.Context, input *mediapackagevod.DescribeAssetInput) *MediapackagevodDescribeAssetResult

       DescribePackagingConfiguration(ctx workflow.Context, input *mediapackagevod.DescribePackagingConfigurationInput) (*mediapackagevod.DescribePackagingConfigurationOutput, error)
       DescribePackagingConfigurationAsync(ctx workflow.Context, input *mediapackagevod.DescribePackagingConfigurationInput) *MediapackagevodDescribePackagingConfigurationResult

       DescribePackagingGroup(ctx workflow.Context, input *mediapackagevod.DescribePackagingGroupInput) (*mediapackagevod.DescribePackagingGroupOutput, error)
       DescribePackagingGroupAsync(ctx workflow.Context, input *mediapackagevod.DescribePackagingGroupInput) *MediapackagevodDescribePackagingGroupResult

       ListAssets(ctx workflow.Context, input *mediapackagevod.ListAssetsInput) (*mediapackagevod.ListAssetsOutput, error)
       ListAssetsAsync(ctx workflow.Context, input *mediapackagevod.ListAssetsInput) *MediapackagevodListAssetsResult

       ListPackagingConfigurations(ctx workflow.Context, input *mediapackagevod.ListPackagingConfigurationsInput) (*mediapackagevod.ListPackagingConfigurationsOutput, error)
       ListPackagingConfigurationsAsync(ctx workflow.Context, input *mediapackagevod.ListPackagingConfigurationsInput) *MediapackagevodListPackagingConfigurationsResult

       ListPackagingGroups(ctx workflow.Context, input *mediapackagevod.ListPackagingGroupsInput) (*mediapackagevod.ListPackagingGroupsOutput, error)
       ListPackagingGroupsAsync(ctx workflow.Context, input *mediapackagevod.ListPackagingGroupsInput) *MediapackagevodListPackagingGroupsResult

       ListTagsForResource(ctx workflow.Context, input *mediapackagevod.ListTagsForResourceInput) (*mediapackagevod.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *mediapackagevod.ListTagsForResourceInput) *MediapackagevodListTagsForResourceResult

       TagResource(ctx workflow.Context, input *mediapackagevod.TagResourceInput) (*mediapackagevod.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *mediapackagevod.TagResourceInput) *MediapackagevodTagResourceResult

       UntagResource(ctx workflow.Context, input *mediapackagevod.UntagResourceInput) (*mediapackagevod.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *mediapackagevod.UntagResourceInput) *MediapackagevodUntagResourceResult

       UpdatePackagingGroup(ctx workflow.Context, input *mediapackagevod.UpdatePackagingGroupInput) (*mediapackagevod.UpdatePackagingGroupOutput, error)
       UpdatePackagingGroupAsync(ctx workflow.Context, input *mediapackagevod.UpdatePackagingGroupInput) *MediapackagevodUpdatePackagingGroupResult
}

type MediapackagevodCreateAssetResult struct {
	Result workflow.Future
}

func (r *MediapackagevodCreateAssetResult) Get(ctx workflow.Context) (*mediapackagevod.CreateAssetOutput, error) {
    var output mediapackagevod.CreateAssetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackagevodCreatePackagingConfigurationResult struct {
	Result workflow.Future
}

func (r *MediapackagevodCreatePackagingConfigurationResult) Get(ctx workflow.Context) (*mediapackagevod.CreatePackagingConfigurationOutput, error) {
    var output mediapackagevod.CreatePackagingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackagevodCreatePackagingGroupResult struct {
	Result workflow.Future
}

func (r *MediapackagevodCreatePackagingGroupResult) Get(ctx workflow.Context) (*mediapackagevod.CreatePackagingGroupOutput, error) {
    var output mediapackagevod.CreatePackagingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackagevodDeleteAssetResult struct {
	Result workflow.Future
}

func (r *MediapackagevodDeleteAssetResult) Get(ctx workflow.Context) (*mediapackagevod.DeleteAssetOutput, error) {
    var output mediapackagevod.DeleteAssetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackagevodDeletePackagingConfigurationResult struct {
	Result workflow.Future
}

func (r *MediapackagevodDeletePackagingConfigurationResult) Get(ctx workflow.Context) (*mediapackagevod.DeletePackagingConfigurationOutput, error) {
    var output mediapackagevod.DeletePackagingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackagevodDeletePackagingGroupResult struct {
	Result workflow.Future
}

func (r *MediapackagevodDeletePackagingGroupResult) Get(ctx workflow.Context) (*mediapackagevod.DeletePackagingGroupOutput, error) {
    var output mediapackagevod.DeletePackagingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackagevodDescribeAssetResult struct {
	Result workflow.Future
}

func (r *MediapackagevodDescribeAssetResult) Get(ctx workflow.Context) (*mediapackagevod.DescribeAssetOutput, error) {
    var output mediapackagevod.DescribeAssetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackagevodDescribePackagingConfigurationResult struct {
	Result workflow.Future
}

func (r *MediapackagevodDescribePackagingConfigurationResult) Get(ctx workflow.Context) (*mediapackagevod.DescribePackagingConfigurationOutput, error) {
    var output mediapackagevod.DescribePackagingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackagevodDescribePackagingGroupResult struct {
	Result workflow.Future
}

func (r *MediapackagevodDescribePackagingGroupResult) Get(ctx workflow.Context) (*mediapackagevod.DescribePackagingGroupOutput, error) {
    var output mediapackagevod.DescribePackagingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackagevodListAssetsResult struct {
	Result workflow.Future
}

func (r *MediapackagevodListAssetsResult) Get(ctx workflow.Context) (*mediapackagevod.ListAssetsOutput, error) {
    var output mediapackagevod.ListAssetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackagevodListPackagingConfigurationsResult struct {
	Result workflow.Future
}

func (r *MediapackagevodListPackagingConfigurationsResult) Get(ctx workflow.Context) (*mediapackagevod.ListPackagingConfigurationsOutput, error) {
    var output mediapackagevod.ListPackagingConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackagevodListPackagingGroupsResult struct {
	Result workflow.Future
}

func (r *MediapackagevodListPackagingGroupsResult) Get(ctx workflow.Context) (*mediapackagevod.ListPackagingGroupsOutput, error) {
    var output mediapackagevod.ListPackagingGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackagevodListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *MediapackagevodListTagsForResourceResult) Get(ctx workflow.Context) (*mediapackagevod.ListTagsForResourceOutput, error) {
    var output mediapackagevod.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackagevodTagResourceResult struct {
	Result workflow.Future
}

func (r *MediapackagevodTagResourceResult) Get(ctx workflow.Context) (*mediapackagevod.TagResourceOutput, error) {
    var output mediapackagevod.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackagevodUntagResourceResult struct {
	Result workflow.Future
}

func (r *MediapackagevodUntagResourceResult) Get(ctx workflow.Context) (*mediapackagevod.UntagResourceOutput, error) {
    var output mediapackagevod.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackagevodUpdatePackagingGroupResult struct {
	Result workflow.Future
}

func (r *MediapackagevodUpdatePackagingGroupResult) Get(ctx workflow.Context) (*mediapackagevod.UpdatePackagingGroupOutput, error) {
    var output mediapackagevod.UpdatePackagingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediaPackageVodStub struct {
    activities awsactivities.MediaPackageVodActivities
}

func NewMediaPackageVodStub() MediaPackageVodClient {
    return &MediaPackageVodStub{}
}

func (a *MediaPackageVodStub) CreateAsset(ctx workflow.Context, input *mediapackagevod.CreateAssetInput) (*mediapackagevod.CreateAssetOutput, error) {
    var output mediapackagevod.CreateAssetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAsset, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageVodStub) CreateAssetAsync(ctx workflow.Context, input *mediapackagevod.CreateAssetInput) *MediapackagevodCreateAssetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateAsset, input)
    return &MediapackagevodCreateAssetResult{Result: future}
}

func (a *MediaPackageVodStub) CreatePackagingConfiguration(ctx workflow.Context, input *mediapackagevod.CreatePackagingConfigurationInput) (*mediapackagevod.CreatePackagingConfigurationOutput, error) {
    var output mediapackagevod.CreatePackagingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePackagingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageVodStub) CreatePackagingConfigurationAsync(ctx workflow.Context, input *mediapackagevod.CreatePackagingConfigurationInput) *MediapackagevodCreatePackagingConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePackagingConfiguration, input)
    return &MediapackagevodCreatePackagingConfigurationResult{Result: future}
}

func (a *MediaPackageVodStub) CreatePackagingGroup(ctx workflow.Context, input *mediapackagevod.CreatePackagingGroupInput) (*mediapackagevod.CreatePackagingGroupOutput, error) {
    var output mediapackagevod.CreatePackagingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePackagingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageVodStub) CreatePackagingGroupAsync(ctx workflow.Context, input *mediapackagevod.CreatePackagingGroupInput) *MediapackagevodCreatePackagingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePackagingGroup, input)
    return &MediapackagevodCreatePackagingGroupResult{Result: future}
}

func (a *MediaPackageVodStub) DeleteAsset(ctx workflow.Context, input *mediapackagevod.DeleteAssetInput) (*mediapackagevod.DeleteAssetOutput, error) {
    var output mediapackagevod.DeleteAssetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAsset, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageVodStub) DeleteAssetAsync(ctx workflow.Context, input *mediapackagevod.DeleteAssetInput) *MediapackagevodDeleteAssetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteAsset, input)
    return &MediapackagevodDeleteAssetResult{Result: future}
}

func (a *MediaPackageVodStub) DeletePackagingConfiguration(ctx workflow.Context, input *mediapackagevod.DeletePackagingConfigurationInput) (*mediapackagevod.DeletePackagingConfigurationOutput, error) {
    var output mediapackagevod.DeletePackagingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePackagingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageVodStub) DeletePackagingConfigurationAsync(ctx workflow.Context, input *mediapackagevod.DeletePackagingConfigurationInput) *MediapackagevodDeletePackagingConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeletePackagingConfiguration, input)
    return &MediapackagevodDeletePackagingConfigurationResult{Result: future}
}

func (a *MediaPackageVodStub) DeletePackagingGroup(ctx workflow.Context, input *mediapackagevod.DeletePackagingGroupInput) (*mediapackagevod.DeletePackagingGroupOutput, error) {
    var output mediapackagevod.DeletePackagingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePackagingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageVodStub) DeletePackagingGroupAsync(ctx workflow.Context, input *mediapackagevod.DeletePackagingGroupInput) *MediapackagevodDeletePackagingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeletePackagingGroup, input)
    return &MediapackagevodDeletePackagingGroupResult{Result: future}
}

func (a *MediaPackageVodStub) DescribeAsset(ctx workflow.Context, input *mediapackagevod.DescribeAssetInput) (*mediapackagevod.DescribeAssetOutput, error) {
    var output mediapackagevod.DescribeAssetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAsset, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageVodStub) DescribeAssetAsync(ctx workflow.Context, input *mediapackagevod.DescribeAssetInput) *MediapackagevodDescribeAssetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAsset, input)
    return &MediapackagevodDescribeAssetResult{Result: future}
}

func (a *MediaPackageVodStub) DescribePackagingConfiguration(ctx workflow.Context, input *mediapackagevod.DescribePackagingConfigurationInput) (*mediapackagevod.DescribePackagingConfigurationOutput, error) {
    var output mediapackagevod.DescribePackagingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePackagingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageVodStub) DescribePackagingConfigurationAsync(ctx workflow.Context, input *mediapackagevod.DescribePackagingConfigurationInput) *MediapackagevodDescribePackagingConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribePackagingConfiguration, input)
    return &MediapackagevodDescribePackagingConfigurationResult{Result: future}
}

func (a *MediaPackageVodStub) DescribePackagingGroup(ctx workflow.Context, input *mediapackagevod.DescribePackagingGroupInput) (*mediapackagevod.DescribePackagingGroupOutput, error) {
    var output mediapackagevod.DescribePackagingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePackagingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageVodStub) DescribePackagingGroupAsync(ctx workflow.Context, input *mediapackagevod.DescribePackagingGroupInput) *MediapackagevodDescribePackagingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribePackagingGroup, input)
    return &MediapackagevodDescribePackagingGroupResult{Result: future}
}

func (a *MediaPackageVodStub) ListAssets(ctx workflow.Context, input *mediapackagevod.ListAssetsInput) (*mediapackagevod.ListAssetsOutput, error) {
    var output mediapackagevod.ListAssetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAssets, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageVodStub) ListAssetsAsync(ctx workflow.Context, input *mediapackagevod.ListAssetsInput) *MediapackagevodListAssetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAssets, input)
    return &MediapackagevodListAssetsResult{Result: future}
}

func (a *MediaPackageVodStub) ListPackagingConfigurations(ctx workflow.Context, input *mediapackagevod.ListPackagingConfigurationsInput) (*mediapackagevod.ListPackagingConfigurationsOutput, error) {
    var output mediapackagevod.ListPackagingConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPackagingConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageVodStub) ListPackagingConfigurationsAsync(ctx workflow.Context, input *mediapackagevod.ListPackagingConfigurationsInput) *MediapackagevodListPackagingConfigurationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPackagingConfigurations, input)
    return &MediapackagevodListPackagingConfigurationsResult{Result: future}
}

func (a *MediaPackageVodStub) ListPackagingGroups(ctx workflow.Context, input *mediapackagevod.ListPackagingGroupsInput) (*mediapackagevod.ListPackagingGroupsOutput, error) {
    var output mediapackagevod.ListPackagingGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPackagingGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageVodStub) ListPackagingGroupsAsync(ctx workflow.Context, input *mediapackagevod.ListPackagingGroupsInput) *MediapackagevodListPackagingGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPackagingGroups, input)
    return &MediapackagevodListPackagingGroupsResult{Result: future}
}

func (a *MediaPackageVodStub) ListTagsForResource(ctx workflow.Context, input *mediapackagevod.ListTagsForResourceInput) (*mediapackagevod.ListTagsForResourceOutput, error) {
    var output mediapackagevod.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageVodStub) ListTagsForResourceAsync(ctx workflow.Context, input *mediapackagevod.ListTagsForResourceInput) *MediapackagevodListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &MediapackagevodListTagsForResourceResult{Result: future}
}

func (a *MediaPackageVodStub) TagResource(ctx workflow.Context, input *mediapackagevod.TagResourceInput) (*mediapackagevod.TagResourceOutput, error) {
    var output mediapackagevod.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageVodStub) TagResourceAsync(ctx workflow.Context, input *mediapackagevod.TagResourceInput) *MediapackagevodTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &MediapackagevodTagResourceResult{Result: future}
}

func (a *MediaPackageVodStub) UntagResource(ctx workflow.Context, input *mediapackagevod.UntagResourceInput) (*mediapackagevod.UntagResourceOutput, error) {
    var output mediapackagevod.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageVodStub) UntagResourceAsync(ctx workflow.Context, input *mediapackagevod.UntagResourceInput) *MediapackagevodUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &MediapackagevodUntagResourceResult{Result: future}
}

func (a *MediaPackageVodStub) UpdatePackagingGroup(ctx workflow.Context, input *mediapackagevod.UpdatePackagingGroupInput) (*mediapackagevod.UpdatePackagingGroupOutput, error) {
    var output mediapackagevod.UpdatePackagingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdatePackagingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageVodStub) UpdatePackagingGroupAsync(ctx workflow.Context, input *mediapackagevod.UpdatePackagingGroupInput) *MediapackagevodUpdatePackagingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdatePackagingGroup, input)
    return &MediapackagevodUpdatePackagingGroupResult{Result: future}
}
