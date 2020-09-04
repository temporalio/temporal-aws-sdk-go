package awsclients

import (
	"github.com/aws/aws-sdk-go/service/appstream"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type AppStreamClient interface {
    AssociateFleet(ctx workflow.Context, input *appstream.AssociateFleetInput) (*appstream.AssociateFleetOutput, error)
    AssociateFleetAsync(ctx workflow.Context, input *appstream.AssociateFleetInput) *AppstreamAssociateFleetResult

    BatchAssociateUserStack(ctx workflow.Context, input *appstream.BatchAssociateUserStackInput) (*appstream.BatchAssociateUserStackOutput, error)
    BatchAssociateUserStackAsync(ctx workflow.Context, input *appstream.BatchAssociateUserStackInput) *AppstreamBatchAssociateUserStackResult

    BatchDisassociateUserStack(ctx workflow.Context, input *appstream.BatchDisassociateUserStackInput) (*appstream.BatchDisassociateUserStackOutput, error)
    BatchDisassociateUserStackAsync(ctx workflow.Context, input *appstream.BatchDisassociateUserStackInput) *AppstreamBatchDisassociateUserStackResult

    CopyImage(ctx workflow.Context, input *appstream.CopyImageInput) (*appstream.CopyImageOutput, error)
    CopyImageAsync(ctx workflow.Context, input *appstream.CopyImageInput) *AppstreamCopyImageResult

    CreateDirectoryConfig(ctx workflow.Context, input *appstream.CreateDirectoryConfigInput) (*appstream.CreateDirectoryConfigOutput, error)
    CreateDirectoryConfigAsync(ctx workflow.Context, input *appstream.CreateDirectoryConfigInput) *AppstreamCreateDirectoryConfigResult

    CreateFleet(ctx workflow.Context, input *appstream.CreateFleetInput) (*appstream.CreateFleetOutput, error)
    CreateFleetAsync(ctx workflow.Context, input *appstream.CreateFleetInput) *AppstreamCreateFleetResult

    CreateImageBuilder(ctx workflow.Context, input *appstream.CreateImageBuilderInput) (*appstream.CreateImageBuilderOutput, error)
    CreateImageBuilderAsync(ctx workflow.Context, input *appstream.CreateImageBuilderInput) *AppstreamCreateImageBuilderResult

    CreateImageBuilderStreamingURL(ctx workflow.Context, input *appstream.CreateImageBuilderStreamingURLInput) (*appstream.CreateImageBuilderStreamingURLOutput, error)
    CreateImageBuilderStreamingURLAsync(ctx workflow.Context, input *appstream.CreateImageBuilderStreamingURLInput) *AppstreamCreateImageBuilderStreamingURLResult

    CreateStack(ctx workflow.Context, input *appstream.CreateStackInput) (*appstream.CreateStackOutput, error)
    CreateStackAsync(ctx workflow.Context, input *appstream.CreateStackInput) *AppstreamCreateStackResult

    CreateStreamingURL(ctx workflow.Context, input *appstream.CreateStreamingURLInput) (*appstream.CreateStreamingURLOutput, error)
    CreateStreamingURLAsync(ctx workflow.Context, input *appstream.CreateStreamingURLInput) *AppstreamCreateStreamingURLResult

    CreateUsageReportSubscription(ctx workflow.Context, input *appstream.CreateUsageReportSubscriptionInput) (*appstream.CreateUsageReportSubscriptionOutput, error)
    CreateUsageReportSubscriptionAsync(ctx workflow.Context, input *appstream.CreateUsageReportSubscriptionInput) *AppstreamCreateUsageReportSubscriptionResult

    CreateUser(ctx workflow.Context, input *appstream.CreateUserInput) (*appstream.CreateUserOutput, error)
    CreateUserAsync(ctx workflow.Context, input *appstream.CreateUserInput) *AppstreamCreateUserResult

    DeleteDirectoryConfig(ctx workflow.Context, input *appstream.DeleteDirectoryConfigInput) (*appstream.DeleteDirectoryConfigOutput, error)
    DeleteDirectoryConfigAsync(ctx workflow.Context, input *appstream.DeleteDirectoryConfigInput) *AppstreamDeleteDirectoryConfigResult

    DeleteFleet(ctx workflow.Context, input *appstream.DeleteFleetInput) (*appstream.DeleteFleetOutput, error)
    DeleteFleetAsync(ctx workflow.Context, input *appstream.DeleteFleetInput) *AppstreamDeleteFleetResult

    DeleteImage(ctx workflow.Context, input *appstream.DeleteImageInput) (*appstream.DeleteImageOutput, error)
    DeleteImageAsync(ctx workflow.Context, input *appstream.DeleteImageInput) *AppstreamDeleteImageResult

    DeleteImageBuilder(ctx workflow.Context, input *appstream.DeleteImageBuilderInput) (*appstream.DeleteImageBuilderOutput, error)
    DeleteImageBuilderAsync(ctx workflow.Context, input *appstream.DeleteImageBuilderInput) *AppstreamDeleteImageBuilderResult

    DeleteImagePermissions(ctx workflow.Context, input *appstream.DeleteImagePermissionsInput) (*appstream.DeleteImagePermissionsOutput, error)
    DeleteImagePermissionsAsync(ctx workflow.Context, input *appstream.DeleteImagePermissionsInput) *AppstreamDeleteImagePermissionsResult

    DeleteStack(ctx workflow.Context, input *appstream.DeleteStackInput) (*appstream.DeleteStackOutput, error)
    DeleteStackAsync(ctx workflow.Context, input *appstream.DeleteStackInput) *AppstreamDeleteStackResult

    DeleteUsageReportSubscription(ctx workflow.Context, input *appstream.DeleteUsageReportSubscriptionInput) (*appstream.DeleteUsageReportSubscriptionOutput, error)
    DeleteUsageReportSubscriptionAsync(ctx workflow.Context, input *appstream.DeleteUsageReportSubscriptionInput) *AppstreamDeleteUsageReportSubscriptionResult

    DeleteUser(ctx workflow.Context, input *appstream.DeleteUserInput) (*appstream.DeleteUserOutput, error)
    DeleteUserAsync(ctx workflow.Context, input *appstream.DeleteUserInput) *AppstreamDeleteUserResult

    DescribeDirectoryConfigs(ctx workflow.Context, input *appstream.DescribeDirectoryConfigsInput) (*appstream.DescribeDirectoryConfigsOutput, error)
    DescribeDirectoryConfigsAsync(ctx workflow.Context, input *appstream.DescribeDirectoryConfigsInput) *AppstreamDescribeDirectoryConfigsResult

    DescribeFleets(ctx workflow.Context, input *appstream.DescribeFleetsInput) (*appstream.DescribeFleetsOutput, error)
    DescribeFleetsAsync(ctx workflow.Context, input *appstream.DescribeFleetsInput) *AppstreamDescribeFleetsResult

    DescribeImageBuilders(ctx workflow.Context, input *appstream.DescribeImageBuildersInput) (*appstream.DescribeImageBuildersOutput, error)
    DescribeImageBuildersAsync(ctx workflow.Context, input *appstream.DescribeImageBuildersInput) *AppstreamDescribeImageBuildersResult

    DescribeImagePermissions(ctx workflow.Context, input *appstream.DescribeImagePermissionsInput) (*appstream.DescribeImagePermissionsOutput, error)
    DescribeImagePermissionsAsync(ctx workflow.Context, input *appstream.DescribeImagePermissionsInput) *AppstreamDescribeImagePermissionsResult

    DescribeImages(ctx workflow.Context, input *appstream.DescribeImagesInput) (*appstream.DescribeImagesOutput, error)
    DescribeImagesAsync(ctx workflow.Context, input *appstream.DescribeImagesInput) *AppstreamDescribeImagesResult

    DescribeSessions(ctx workflow.Context, input *appstream.DescribeSessionsInput) (*appstream.DescribeSessionsOutput, error)
    DescribeSessionsAsync(ctx workflow.Context, input *appstream.DescribeSessionsInput) *AppstreamDescribeSessionsResult

    DescribeStacks(ctx workflow.Context, input *appstream.DescribeStacksInput) (*appstream.DescribeStacksOutput, error)
    DescribeStacksAsync(ctx workflow.Context, input *appstream.DescribeStacksInput) *AppstreamDescribeStacksResult

    DescribeUsageReportSubscriptions(ctx workflow.Context, input *appstream.DescribeUsageReportSubscriptionsInput) (*appstream.DescribeUsageReportSubscriptionsOutput, error)
    DescribeUsageReportSubscriptionsAsync(ctx workflow.Context, input *appstream.DescribeUsageReportSubscriptionsInput) *AppstreamDescribeUsageReportSubscriptionsResult

    DescribeUserStackAssociations(ctx workflow.Context, input *appstream.DescribeUserStackAssociationsInput) (*appstream.DescribeUserStackAssociationsOutput, error)
    DescribeUserStackAssociationsAsync(ctx workflow.Context, input *appstream.DescribeUserStackAssociationsInput) *AppstreamDescribeUserStackAssociationsResult

    DescribeUsers(ctx workflow.Context, input *appstream.DescribeUsersInput) (*appstream.DescribeUsersOutput, error)
    DescribeUsersAsync(ctx workflow.Context, input *appstream.DescribeUsersInput) *AppstreamDescribeUsersResult

    DisableUser(ctx workflow.Context, input *appstream.DisableUserInput) (*appstream.DisableUserOutput, error)
    DisableUserAsync(ctx workflow.Context, input *appstream.DisableUserInput) *AppstreamDisableUserResult

    DisassociateFleet(ctx workflow.Context, input *appstream.DisassociateFleetInput) (*appstream.DisassociateFleetOutput, error)
    DisassociateFleetAsync(ctx workflow.Context, input *appstream.DisassociateFleetInput) *AppstreamDisassociateFleetResult

    EnableUser(ctx workflow.Context, input *appstream.EnableUserInput) (*appstream.EnableUserOutput, error)
    EnableUserAsync(ctx workflow.Context, input *appstream.EnableUserInput) *AppstreamEnableUserResult

    ExpireSession(ctx workflow.Context, input *appstream.ExpireSessionInput) (*appstream.ExpireSessionOutput, error)
    ExpireSessionAsync(ctx workflow.Context, input *appstream.ExpireSessionInput) *AppstreamExpireSessionResult

    ListAssociatedFleets(ctx workflow.Context, input *appstream.ListAssociatedFleetsInput) (*appstream.ListAssociatedFleetsOutput, error)
    ListAssociatedFleetsAsync(ctx workflow.Context, input *appstream.ListAssociatedFleetsInput) *AppstreamListAssociatedFleetsResult

    ListAssociatedStacks(ctx workflow.Context, input *appstream.ListAssociatedStacksInput) (*appstream.ListAssociatedStacksOutput, error)
    ListAssociatedStacksAsync(ctx workflow.Context, input *appstream.ListAssociatedStacksInput) *AppstreamListAssociatedStacksResult

    ListTagsForResource(ctx workflow.Context, input *appstream.ListTagsForResourceInput) (*appstream.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *appstream.ListTagsForResourceInput) *AppstreamListTagsForResourceResult

    StartFleet(ctx workflow.Context, input *appstream.StartFleetInput) (*appstream.StartFleetOutput, error)
    StartFleetAsync(ctx workflow.Context, input *appstream.StartFleetInput) *AppstreamStartFleetResult

    StartImageBuilder(ctx workflow.Context, input *appstream.StartImageBuilderInput) (*appstream.StartImageBuilderOutput, error)
    StartImageBuilderAsync(ctx workflow.Context, input *appstream.StartImageBuilderInput) *AppstreamStartImageBuilderResult

    StopFleet(ctx workflow.Context, input *appstream.StopFleetInput) (*appstream.StopFleetOutput, error)
    StopFleetAsync(ctx workflow.Context, input *appstream.StopFleetInput) *AppstreamStopFleetResult

    StopImageBuilder(ctx workflow.Context, input *appstream.StopImageBuilderInput) (*appstream.StopImageBuilderOutput, error)
    StopImageBuilderAsync(ctx workflow.Context, input *appstream.StopImageBuilderInput) *AppstreamStopImageBuilderResult

    TagResource(ctx workflow.Context, input *appstream.TagResourceInput) (*appstream.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *appstream.TagResourceInput) *AppstreamTagResourceResult

    UntagResource(ctx workflow.Context, input *appstream.UntagResourceInput) (*appstream.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *appstream.UntagResourceInput) *AppstreamUntagResourceResult

    UpdateDirectoryConfig(ctx workflow.Context, input *appstream.UpdateDirectoryConfigInput) (*appstream.UpdateDirectoryConfigOutput, error)
    UpdateDirectoryConfigAsync(ctx workflow.Context, input *appstream.UpdateDirectoryConfigInput) *AppstreamUpdateDirectoryConfigResult

    UpdateFleet(ctx workflow.Context, input *appstream.UpdateFleetInput) (*appstream.UpdateFleetOutput, error)
    UpdateFleetAsync(ctx workflow.Context, input *appstream.UpdateFleetInput) *AppstreamUpdateFleetResult

    UpdateImagePermissions(ctx workflow.Context, input *appstream.UpdateImagePermissionsInput) (*appstream.UpdateImagePermissionsOutput, error)
    UpdateImagePermissionsAsync(ctx workflow.Context, input *appstream.UpdateImagePermissionsInput) *AppstreamUpdateImagePermissionsResult

    UpdateStack(ctx workflow.Context, input *appstream.UpdateStackInput) (*appstream.UpdateStackOutput, error)
    UpdateStackAsync(ctx workflow.Context, input *appstream.UpdateStackInput) *AppstreamUpdateStackResult

    WaitUntilFleetStarted(ctx workflow.Context, input *appstream.DescribeFleetsInput) error
    WaitUntilFleetStopped(ctx workflow.Context, input *appstream.DescribeFleetsInput) error}
type AppstreamAssociateFleetResult struct {
	Result workflow.Future
}

func (r *AppstreamAssociateFleetResult) Get(ctx workflow.Context) (*appstream.AssociateFleetOutput, error) {
    var output appstream.AssociateFleetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamBatchAssociateUserStackResult struct {
	Result workflow.Future
}

func (r *AppstreamBatchAssociateUserStackResult) Get(ctx workflow.Context) (*appstream.BatchAssociateUserStackOutput, error) {
    var output appstream.BatchAssociateUserStackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamBatchDisassociateUserStackResult struct {
	Result workflow.Future
}

func (r *AppstreamBatchDisassociateUserStackResult) Get(ctx workflow.Context) (*appstream.BatchDisassociateUserStackOutput, error) {
    var output appstream.BatchDisassociateUserStackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamCopyImageResult struct {
	Result workflow.Future
}

func (r *AppstreamCopyImageResult) Get(ctx workflow.Context) (*appstream.CopyImageOutput, error) {
    var output appstream.CopyImageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamCreateDirectoryConfigResult struct {
	Result workflow.Future
}

func (r *AppstreamCreateDirectoryConfigResult) Get(ctx workflow.Context) (*appstream.CreateDirectoryConfigOutput, error) {
    var output appstream.CreateDirectoryConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamCreateFleetResult struct {
	Result workflow.Future
}

func (r *AppstreamCreateFleetResult) Get(ctx workflow.Context) (*appstream.CreateFleetOutput, error) {
    var output appstream.CreateFleetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamCreateImageBuilderResult struct {
	Result workflow.Future
}

func (r *AppstreamCreateImageBuilderResult) Get(ctx workflow.Context) (*appstream.CreateImageBuilderOutput, error) {
    var output appstream.CreateImageBuilderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamCreateImageBuilderStreamingURLResult struct {
	Result workflow.Future
}

func (r *AppstreamCreateImageBuilderStreamingURLResult) Get(ctx workflow.Context) (*appstream.CreateImageBuilderStreamingURLOutput, error) {
    var output appstream.CreateImageBuilderStreamingURLOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamCreateStackResult struct {
	Result workflow.Future
}

func (r *AppstreamCreateStackResult) Get(ctx workflow.Context) (*appstream.CreateStackOutput, error) {
    var output appstream.CreateStackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamCreateStreamingURLResult struct {
	Result workflow.Future
}

func (r *AppstreamCreateStreamingURLResult) Get(ctx workflow.Context) (*appstream.CreateStreamingURLOutput, error) {
    var output appstream.CreateStreamingURLOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamCreateUsageReportSubscriptionResult struct {
	Result workflow.Future
}

func (r *AppstreamCreateUsageReportSubscriptionResult) Get(ctx workflow.Context) (*appstream.CreateUsageReportSubscriptionOutput, error) {
    var output appstream.CreateUsageReportSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamCreateUserResult struct {
	Result workflow.Future
}

func (r *AppstreamCreateUserResult) Get(ctx workflow.Context) (*appstream.CreateUserOutput, error) {
    var output appstream.CreateUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamDeleteDirectoryConfigResult struct {
	Result workflow.Future
}

func (r *AppstreamDeleteDirectoryConfigResult) Get(ctx workflow.Context) (*appstream.DeleteDirectoryConfigOutput, error) {
    var output appstream.DeleteDirectoryConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamDeleteFleetResult struct {
	Result workflow.Future
}

func (r *AppstreamDeleteFleetResult) Get(ctx workflow.Context) (*appstream.DeleteFleetOutput, error) {
    var output appstream.DeleteFleetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamDeleteImageResult struct {
	Result workflow.Future
}

func (r *AppstreamDeleteImageResult) Get(ctx workflow.Context) (*appstream.DeleteImageOutput, error) {
    var output appstream.DeleteImageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamDeleteImageBuilderResult struct {
	Result workflow.Future
}

func (r *AppstreamDeleteImageBuilderResult) Get(ctx workflow.Context) (*appstream.DeleteImageBuilderOutput, error) {
    var output appstream.DeleteImageBuilderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamDeleteImagePermissionsResult struct {
	Result workflow.Future
}

func (r *AppstreamDeleteImagePermissionsResult) Get(ctx workflow.Context) (*appstream.DeleteImagePermissionsOutput, error) {
    var output appstream.DeleteImagePermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamDeleteStackResult struct {
	Result workflow.Future
}

func (r *AppstreamDeleteStackResult) Get(ctx workflow.Context) (*appstream.DeleteStackOutput, error) {
    var output appstream.DeleteStackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamDeleteUsageReportSubscriptionResult struct {
	Result workflow.Future
}

func (r *AppstreamDeleteUsageReportSubscriptionResult) Get(ctx workflow.Context) (*appstream.DeleteUsageReportSubscriptionOutput, error) {
    var output appstream.DeleteUsageReportSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamDeleteUserResult struct {
	Result workflow.Future
}

func (r *AppstreamDeleteUserResult) Get(ctx workflow.Context) (*appstream.DeleteUserOutput, error) {
    var output appstream.DeleteUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamDescribeDirectoryConfigsResult struct {
	Result workflow.Future
}

func (r *AppstreamDescribeDirectoryConfigsResult) Get(ctx workflow.Context) (*appstream.DescribeDirectoryConfigsOutput, error) {
    var output appstream.DescribeDirectoryConfigsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamDescribeFleetsResult struct {
	Result workflow.Future
}

func (r *AppstreamDescribeFleetsResult) Get(ctx workflow.Context) (*appstream.DescribeFleetsOutput, error) {
    var output appstream.DescribeFleetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamDescribeImageBuildersResult struct {
	Result workflow.Future
}

func (r *AppstreamDescribeImageBuildersResult) Get(ctx workflow.Context) (*appstream.DescribeImageBuildersOutput, error) {
    var output appstream.DescribeImageBuildersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamDescribeImagePermissionsResult struct {
	Result workflow.Future
}

func (r *AppstreamDescribeImagePermissionsResult) Get(ctx workflow.Context) (*appstream.DescribeImagePermissionsOutput, error) {
    var output appstream.DescribeImagePermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamDescribeImagesResult struct {
	Result workflow.Future
}

func (r *AppstreamDescribeImagesResult) Get(ctx workflow.Context) (*appstream.DescribeImagesOutput, error) {
    var output appstream.DescribeImagesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamDescribeSessionsResult struct {
	Result workflow.Future
}

func (r *AppstreamDescribeSessionsResult) Get(ctx workflow.Context) (*appstream.DescribeSessionsOutput, error) {
    var output appstream.DescribeSessionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamDescribeStacksResult struct {
	Result workflow.Future
}

func (r *AppstreamDescribeStacksResult) Get(ctx workflow.Context) (*appstream.DescribeStacksOutput, error) {
    var output appstream.DescribeStacksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamDescribeUsageReportSubscriptionsResult struct {
	Result workflow.Future
}

func (r *AppstreamDescribeUsageReportSubscriptionsResult) Get(ctx workflow.Context) (*appstream.DescribeUsageReportSubscriptionsOutput, error) {
    var output appstream.DescribeUsageReportSubscriptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamDescribeUserStackAssociationsResult struct {
	Result workflow.Future
}

func (r *AppstreamDescribeUserStackAssociationsResult) Get(ctx workflow.Context) (*appstream.DescribeUserStackAssociationsOutput, error) {
    var output appstream.DescribeUserStackAssociationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamDescribeUsersResult struct {
	Result workflow.Future
}

func (r *AppstreamDescribeUsersResult) Get(ctx workflow.Context) (*appstream.DescribeUsersOutput, error) {
    var output appstream.DescribeUsersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamDisableUserResult struct {
	Result workflow.Future
}

func (r *AppstreamDisableUserResult) Get(ctx workflow.Context) (*appstream.DisableUserOutput, error) {
    var output appstream.DisableUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamDisassociateFleetResult struct {
	Result workflow.Future
}

func (r *AppstreamDisassociateFleetResult) Get(ctx workflow.Context) (*appstream.DisassociateFleetOutput, error) {
    var output appstream.DisassociateFleetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamEnableUserResult struct {
	Result workflow.Future
}

func (r *AppstreamEnableUserResult) Get(ctx workflow.Context) (*appstream.EnableUserOutput, error) {
    var output appstream.EnableUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamExpireSessionResult struct {
	Result workflow.Future
}

func (r *AppstreamExpireSessionResult) Get(ctx workflow.Context) (*appstream.ExpireSessionOutput, error) {
    var output appstream.ExpireSessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamListAssociatedFleetsResult struct {
	Result workflow.Future
}

func (r *AppstreamListAssociatedFleetsResult) Get(ctx workflow.Context) (*appstream.ListAssociatedFleetsOutput, error) {
    var output appstream.ListAssociatedFleetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamListAssociatedStacksResult struct {
	Result workflow.Future
}

func (r *AppstreamListAssociatedStacksResult) Get(ctx workflow.Context) (*appstream.ListAssociatedStacksOutput, error) {
    var output appstream.ListAssociatedStacksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *AppstreamListTagsForResourceResult) Get(ctx workflow.Context) (*appstream.ListTagsForResourceOutput, error) {
    var output appstream.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamStartFleetResult struct {
	Result workflow.Future
}

func (r *AppstreamStartFleetResult) Get(ctx workflow.Context) (*appstream.StartFleetOutput, error) {
    var output appstream.StartFleetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamStartImageBuilderResult struct {
	Result workflow.Future
}

func (r *AppstreamStartImageBuilderResult) Get(ctx workflow.Context) (*appstream.StartImageBuilderOutput, error) {
    var output appstream.StartImageBuilderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamStopFleetResult struct {
	Result workflow.Future
}

func (r *AppstreamStopFleetResult) Get(ctx workflow.Context) (*appstream.StopFleetOutput, error) {
    var output appstream.StopFleetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamStopImageBuilderResult struct {
	Result workflow.Future
}

func (r *AppstreamStopImageBuilderResult) Get(ctx workflow.Context) (*appstream.StopImageBuilderOutput, error) {
    var output appstream.StopImageBuilderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamTagResourceResult struct {
	Result workflow.Future
}

func (r *AppstreamTagResourceResult) Get(ctx workflow.Context) (*appstream.TagResourceOutput, error) {
    var output appstream.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamUntagResourceResult struct {
	Result workflow.Future
}

func (r *AppstreamUntagResourceResult) Get(ctx workflow.Context) (*appstream.UntagResourceOutput, error) {
    var output appstream.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamUpdateDirectoryConfigResult struct {
	Result workflow.Future
}

func (r *AppstreamUpdateDirectoryConfigResult) Get(ctx workflow.Context) (*appstream.UpdateDirectoryConfigOutput, error) {
    var output appstream.UpdateDirectoryConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamUpdateFleetResult struct {
	Result workflow.Future
}

func (r *AppstreamUpdateFleetResult) Get(ctx workflow.Context) (*appstream.UpdateFleetOutput, error) {
    var output appstream.UpdateFleetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamUpdateImagePermissionsResult struct {
	Result workflow.Future
}

func (r *AppstreamUpdateImagePermissionsResult) Get(ctx workflow.Context) (*appstream.UpdateImagePermissionsOutput, error) {
    var output appstream.UpdateImagePermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppstreamUpdateStackResult struct {
	Result workflow.Future
}

func (r *AppstreamUpdateStackResult) Get(ctx workflow.Context) (*appstream.UpdateStackOutput, error) {
    var output appstream.UpdateStackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type AppStreamStub struct {
    activities awsactivities.AppStreamActivities
}

func NewAppStreamStub() AppStreamClient {
    return &AppStreamStub{}
}
func (a *AppStreamStub) AssociateFleet(ctx workflow.Context, input *appstream.AssociateFleetInput) (*appstream.AssociateFleetOutput, error) {
    var output appstream.AssociateFleetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateFleet, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) AssociateFleetAsync(ctx workflow.Context, input *appstream.AssociateFleetInput) *AppstreamAssociateFleetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateFleet, input)
    return &AppstreamAssociateFleetResult{Result: future}
}
func (a *AppStreamStub) BatchAssociateUserStack(ctx workflow.Context, input *appstream.BatchAssociateUserStackInput) (*appstream.BatchAssociateUserStackOutput, error) {
    var output appstream.BatchAssociateUserStackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchAssociateUserStack, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) BatchAssociateUserStackAsync(ctx workflow.Context, input *appstream.BatchAssociateUserStackInput) *AppstreamBatchAssociateUserStackResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchAssociateUserStack, input)
    return &AppstreamBatchAssociateUserStackResult{Result: future}
}
func (a *AppStreamStub) BatchDisassociateUserStack(ctx workflow.Context, input *appstream.BatchDisassociateUserStackInput) (*appstream.BatchDisassociateUserStackOutput, error) {
    var output appstream.BatchDisassociateUserStackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchDisassociateUserStack, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) BatchDisassociateUserStackAsync(ctx workflow.Context, input *appstream.BatchDisassociateUserStackInput) *AppstreamBatchDisassociateUserStackResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchDisassociateUserStack, input)
    return &AppstreamBatchDisassociateUserStackResult{Result: future}
}
func (a *AppStreamStub) CopyImage(ctx workflow.Context, input *appstream.CopyImageInput) (*appstream.CopyImageOutput, error) {
    var output appstream.CopyImageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CopyImage, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) CopyImageAsync(ctx workflow.Context, input *appstream.CopyImageInput) *AppstreamCopyImageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CopyImage, input)
    return &AppstreamCopyImageResult{Result: future}
}
func (a *AppStreamStub) CreateDirectoryConfig(ctx workflow.Context, input *appstream.CreateDirectoryConfigInput) (*appstream.CreateDirectoryConfigOutput, error) {
    var output appstream.CreateDirectoryConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDirectoryConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) CreateDirectoryConfigAsync(ctx workflow.Context, input *appstream.CreateDirectoryConfigInput) *AppstreamCreateDirectoryConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDirectoryConfig, input)
    return &AppstreamCreateDirectoryConfigResult{Result: future}
}
func (a *AppStreamStub) CreateFleet(ctx workflow.Context, input *appstream.CreateFleetInput) (*appstream.CreateFleetOutput, error) {
    var output appstream.CreateFleetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFleet, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) CreateFleetAsync(ctx workflow.Context, input *appstream.CreateFleetInput) *AppstreamCreateFleetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateFleet, input)
    return &AppstreamCreateFleetResult{Result: future}
}
func (a *AppStreamStub) CreateImageBuilder(ctx workflow.Context, input *appstream.CreateImageBuilderInput) (*appstream.CreateImageBuilderOutput, error) {
    var output appstream.CreateImageBuilderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateImageBuilder, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) CreateImageBuilderAsync(ctx workflow.Context, input *appstream.CreateImageBuilderInput) *AppstreamCreateImageBuilderResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateImageBuilder, input)
    return &AppstreamCreateImageBuilderResult{Result: future}
}
func (a *AppStreamStub) CreateImageBuilderStreamingURL(ctx workflow.Context, input *appstream.CreateImageBuilderStreamingURLInput) (*appstream.CreateImageBuilderStreamingURLOutput, error) {
    var output appstream.CreateImageBuilderStreamingURLOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateImageBuilderStreamingURL, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) CreateImageBuilderStreamingURLAsync(ctx workflow.Context, input *appstream.CreateImageBuilderStreamingURLInput) *AppstreamCreateImageBuilderStreamingURLResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateImageBuilderStreamingURL, input)
    return &AppstreamCreateImageBuilderStreamingURLResult{Result: future}
}
func (a *AppStreamStub) CreateStack(ctx workflow.Context, input *appstream.CreateStackInput) (*appstream.CreateStackOutput, error) {
    var output appstream.CreateStackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateStack, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) CreateStackAsync(ctx workflow.Context, input *appstream.CreateStackInput) *AppstreamCreateStackResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateStack, input)
    return &AppstreamCreateStackResult{Result: future}
}
func (a *AppStreamStub) CreateStreamingURL(ctx workflow.Context, input *appstream.CreateStreamingURLInput) (*appstream.CreateStreamingURLOutput, error) {
    var output appstream.CreateStreamingURLOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateStreamingURL, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) CreateStreamingURLAsync(ctx workflow.Context, input *appstream.CreateStreamingURLInput) *AppstreamCreateStreamingURLResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateStreamingURL, input)
    return &AppstreamCreateStreamingURLResult{Result: future}
}
func (a *AppStreamStub) CreateUsageReportSubscription(ctx workflow.Context, input *appstream.CreateUsageReportSubscriptionInput) (*appstream.CreateUsageReportSubscriptionOutput, error) {
    var output appstream.CreateUsageReportSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateUsageReportSubscription, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) CreateUsageReportSubscriptionAsync(ctx workflow.Context, input *appstream.CreateUsageReportSubscriptionInput) *AppstreamCreateUsageReportSubscriptionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateUsageReportSubscription, input)
    return &AppstreamCreateUsageReportSubscriptionResult{Result: future}
}
func (a *AppStreamStub) CreateUser(ctx workflow.Context, input *appstream.CreateUserInput) (*appstream.CreateUserOutput, error) {
    var output appstream.CreateUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateUser, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) CreateUserAsync(ctx workflow.Context, input *appstream.CreateUserInput) *AppstreamCreateUserResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateUser, input)
    return &AppstreamCreateUserResult{Result: future}
}
func (a *AppStreamStub) DeleteDirectoryConfig(ctx workflow.Context, input *appstream.DeleteDirectoryConfigInput) (*appstream.DeleteDirectoryConfigOutput, error) {
    var output appstream.DeleteDirectoryConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDirectoryConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) DeleteDirectoryConfigAsync(ctx workflow.Context, input *appstream.DeleteDirectoryConfigInput) *AppstreamDeleteDirectoryConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDirectoryConfig, input)
    return &AppstreamDeleteDirectoryConfigResult{Result: future}
}
func (a *AppStreamStub) DeleteFleet(ctx workflow.Context, input *appstream.DeleteFleetInput) (*appstream.DeleteFleetOutput, error) {
    var output appstream.DeleteFleetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFleet, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) DeleteFleetAsync(ctx workflow.Context, input *appstream.DeleteFleetInput) *AppstreamDeleteFleetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteFleet, input)
    return &AppstreamDeleteFleetResult{Result: future}
}
func (a *AppStreamStub) DeleteImage(ctx workflow.Context, input *appstream.DeleteImageInput) (*appstream.DeleteImageOutput, error) {
    var output appstream.DeleteImageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteImage, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) DeleteImageAsync(ctx workflow.Context, input *appstream.DeleteImageInput) *AppstreamDeleteImageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteImage, input)
    return &AppstreamDeleteImageResult{Result: future}
}
func (a *AppStreamStub) DeleteImageBuilder(ctx workflow.Context, input *appstream.DeleteImageBuilderInput) (*appstream.DeleteImageBuilderOutput, error) {
    var output appstream.DeleteImageBuilderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteImageBuilder, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) DeleteImageBuilderAsync(ctx workflow.Context, input *appstream.DeleteImageBuilderInput) *AppstreamDeleteImageBuilderResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteImageBuilder, input)
    return &AppstreamDeleteImageBuilderResult{Result: future}
}
func (a *AppStreamStub) DeleteImagePermissions(ctx workflow.Context, input *appstream.DeleteImagePermissionsInput) (*appstream.DeleteImagePermissionsOutput, error) {
    var output appstream.DeleteImagePermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteImagePermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) DeleteImagePermissionsAsync(ctx workflow.Context, input *appstream.DeleteImagePermissionsInput) *AppstreamDeleteImagePermissionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteImagePermissions, input)
    return &AppstreamDeleteImagePermissionsResult{Result: future}
}
func (a *AppStreamStub) DeleteStack(ctx workflow.Context, input *appstream.DeleteStackInput) (*appstream.DeleteStackOutput, error) {
    var output appstream.DeleteStackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteStack, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) DeleteStackAsync(ctx workflow.Context, input *appstream.DeleteStackInput) *AppstreamDeleteStackResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteStack, input)
    return &AppstreamDeleteStackResult{Result: future}
}
func (a *AppStreamStub) DeleteUsageReportSubscription(ctx workflow.Context, input *appstream.DeleteUsageReportSubscriptionInput) (*appstream.DeleteUsageReportSubscriptionOutput, error) {
    var output appstream.DeleteUsageReportSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteUsageReportSubscription, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) DeleteUsageReportSubscriptionAsync(ctx workflow.Context, input *appstream.DeleteUsageReportSubscriptionInput) *AppstreamDeleteUsageReportSubscriptionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteUsageReportSubscription, input)
    return &AppstreamDeleteUsageReportSubscriptionResult{Result: future}
}
func (a *AppStreamStub) DeleteUser(ctx workflow.Context, input *appstream.DeleteUserInput) (*appstream.DeleteUserOutput, error) {
    var output appstream.DeleteUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteUser, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) DeleteUserAsync(ctx workflow.Context, input *appstream.DeleteUserInput) *AppstreamDeleteUserResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteUser, input)
    return &AppstreamDeleteUserResult{Result: future}
}
func (a *AppStreamStub) DescribeDirectoryConfigs(ctx workflow.Context, input *appstream.DescribeDirectoryConfigsInput) (*appstream.DescribeDirectoryConfigsOutput, error) {
    var output appstream.DescribeDirectoryConfigsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDirectoryConfigs, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) DescribeDirectoryConfigsAsync(ctx workflow.Context, input *appstream.DescribeDirectoryConfigsInput) *AppstreamDescribeDirectoryConfigsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDirectoryConfigs, input)
    return &AppstreamDescribeDirectoryConfigsResult{Result: future}
}
func (a *AppStreamStub) DescribeFleets(ctx workflow.Context, input *appstream.DescribeFleetsInput) (*appstream.DescribeFleetsOutput, error) {
    var output appstream.DescribeFleetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFleets, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) DescribeFleetsAsync(ctx workflow.Context, input *appstream.DescribeFleetsInput) *AppstreamDescribeFleetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeFleets, input)
    return &AppstreamDescribeFleetsResult{Result: future}
}
func (a *AppStreamStub) DescribeImageBuilders(ctx workflow.Context, input *appstream.DescribeImageBuildersInput) (*appstream.DescribeImageBuildersOutput, error) {
    var output appstream.DescribeImageBuildersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeImageBuilders, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) DescribeImageBuildersAsync(ctx workflow.Context, input *appstream.DescribeImageBuildersInput) *AppstreamDescribeImageBuildersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeImageBuilders, input)
    return &AppstreamDescribeImageBuildersResult{Result: future}
}
func (a *AppStreamStub) DescribeImagePermissions(ctx workflow.Context, input *appstream.DescribeImagePermissionsInput) (*appstream.DescribeImagePermissionsOutput, error) {
    var output appstream.DescribeImagePermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeImagePermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) DescribeImagePermissionsAsync(ctx workflow.Context, input *appstream.DescribeImagePermissionsInput) *AppstreamDescribeImagePermissionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeImagePermissions, input)
    return &AppstreamDescribeImagePermissionsResult{Result: future}
}
func (a *AppStreamStub) DescribeImages(ctx workflow.Context, input *appstream.DescribeImagesInput) (*appstream.DescribeImagesOutput, error) {
    var output appstream.DescribeImagesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeImages, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) DescribeImagesAsync(ctx workflow.Context, input *appstream.DescribeImagesInput) *AppstreamDescribeImagesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeImages, input)
    return &AppstreamDescribeImagesResult{Result: future}
}
func (a *AppStreamStub) DescribeSessions(ctx workflow.Context, input *appstream.DescribeSessionsInput) (*appstream.DescribeSessionsOutput, error) {
    var output appstream.DescribeSessionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSessions, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) DescribeSessionsAsync(ctx workflow.Context, input *appstream.DescribeSessionsInput) *AppstreamDescribeSessionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSessions, input)
    return &AppstreamDescribeSessionsResult{Result: future}
}
func (a *AppStreamStub) DescribeStacks(ctx workflow.Context, input *appstream.DescribeStacksInput) (*appstream.DescribeStacksOutput, error) {
    var output appstream.DescribeStacksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStacks, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) DescribeStacksAsync(ctx workflow.Context, input *appstream.DescribeStacksInput) *AppstreamDescribeStacksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeStacks, input)
    return &AppstreamDescribeStacksResult{Result: future}
}
func (a *AppStreamStub) DescribeUsageReportSubscriptions(ctx workflow.Context, input *appstream.DescribeUsageReportSubscriptionsInput) (*appstream.DescribeUsageReportSubscriptionsOutput, error) {
    var output appstream.DescribeUsageReportSubscriptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeUsageReportSubscriptions, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) DescribeUsageReportSubscriptionsAsync(ctx workflow.Context, input *appstream.DescribeUsageReportSubscriptionsInput) *AppstreamDescribeUsageReportSubscriptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeUsageReportSubscriptions, input)
    return &AppstreamDescribeUsageReportSubscriptionsResult{Result: future}
}
func (a *AppStreamStub) DescribeUserStackAssociations(ctx workflow.Context, input *appstream.DescribeUserStackAssociationsInput) (*appstream.DescribeUserStackAssociationsOutput, error) {
    var output appstream.DescribeUserStackAssociationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeUserStackAssociations, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) DescribeUserStackAssociationsAsync(ctx workflow.Context, input *appstream.DescribeUserStackAssociationsInput) *AppstreamDescribeUserStackAssociationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeUserStackAssociations, input)
    return &AppstreamDescribeUserStackAssociationsResult{Result: future}
}
func (a *AppStreamStub) DescribeUsers(ctx workflow.Context, input *appstream.DescribeUsersInput) (*appstream.DescribeUsersOutput, error) {
    var output appstream.DescribeUsersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeUsers, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) DescribeUsersAsync(ctx workflow.Context, input *appstream.DescribeUsersInput) *AppstreamDescribeUsersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeUsers, input)
    return &AppstreamDescribeUsersResult{Result: future}
}
func (a *AppStreamStub) DisableUser(ctx workflow.Context, input *appstream.DisableUserInput) (*appstream.DisableUserOutput, error) {
    var output appstream.DisableUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableUser, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) DisableUserAsync(ctx workflow.Context, input *appstream.DisableUserInput) *AppstreamDisableUserResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisableUser, input)
    return &AppstreamDisableUserResult{Result: future}
}
func (a *AppStreamStub) DisassociateFleet(ctx workflow.Context, input *appstream.DisassociateFleetInput) (*appstream.DisassociateFleetOutput, error) {
    var output appstream.DisassociateFleetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateFleet, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) DisassociateFleetAsync(ctx workflow.Context, input *appstream.DisassociateFleetInput) *AppstreamDisassociateFleetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateFleet, input)
    return &AppstreamDisassociateFleetResult{Result: future}
}
func (a *AppStreamStub) EnableUser(ctx workflow.Context, input *appstream.EnableUserInput) (*appstream.EnableUserOutput, error) {
    var output appstream.EnableUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableUser, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) EnableUserAsync(ctx workflow.Context, input *appstream.EnableUserInput) *AppstreamEnableUserResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnableUser, input)
    return &AppstreamEnableUserResult{Result: future}
}
func (a *AppStreamStub) ExpireSession(ctx workflow.Context, input *appstream.ExpireSessionInput) (*appstream.ExpireSessionOutput, error) {
    var output appstream.ExpireSessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ExpireSession, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) ExpireSessionAsync(ctx workflow.Context, input *appstream.ExpireSessionInput) *AppstreamExpireSessionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ExpireSession, input)
    return &AppstreamExpireSessionResult{Result: future}
}
func (a *AppStreamStub) ListAssociatedFleets(ctx workflow.Context, input *appstream.ListAssociatedFleetsInput) (*appstream.ListAssociatedFleetsOutput, error) {
    var output appstream.ListAssociatedFleetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAssociatedFleets, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) ListAssociatedFleetsAsync(ctx workflow.Context, input *appstream.ListAssociatedFleetsInput) *AppstreamListAssociatedFleetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAssociatedFleets, input)
    return &AppstreamListAssociatedFleetsResult{Result: future}
}
func (a *AppStreamStub) ListAssociatedStacks(ctx workflow.Context, input *appstream.ListAssociatedStacksInput) (*appstream.ListAssociatedStacksOutput, error) {
    var output appstream.ListAssociatedStacksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAssociatedStacks, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) ListAssociatedStacksAsync(ctx workflow.Context, input *appstream.ListAssociatedStacksInput) *AppstreamListAssociatedStacksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAssociatedStacks, input)
    return &AppstreamListAssociatedStacksResult{Result: future}
}
func (a *AppStreamStub) ListTagsForResource(ctx workflow.Context, input *appstream.ListTagsForResourceInput) (*appstream.ListTagsForResourceOutput, error) {
    var output appstream.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) ListTagsForResourceAsync(ctx workflow.Context, input *appstream.ListTagsForResourceInput) *AppstreamListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &AppstreamListTagsForResourceResult{Result: future}
}
func (a *AppStreamStub) StartFleet(ctx workflow.Context, input *appstream.StartFleetInput) (*appstream.StartFleetOutput, error) {
    var output appstream.StartFleetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartFleet, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) StartFleetAsync(ctx workflow.Context, input *appstream.StartFleetInput) *AppstreamStartFleetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartFleet, input)
    return &AppstreamStartFleetResult{Result: future}
}
func (a *AppStreamStub) StartImageBuilder(ctx workflow.Context, input *appstream.StartImageBuilderInput) (*appstream.StartImageBuilderOutput, error) {
    var output appstream.StartImageBuilderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartImageBuilder, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) StartImageBuilderAsync(ctx workflow.Context, input *appstream.StartImageBuilderInput) *AppstreamStartImageBuilderResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartImageBuilder, input)
    return &AppstreamStartImageBuilderResult{Result: future}
}
func (a *AppStreamStub) StopFleet(ctx workflow.Context, input *appstream.StopFleetInput) (*appstream.StopFleetOutput, error) {
    var output appstream.StopFleetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopFleet, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) StopFleetAsync(ctx workflow.Context, input *appstream.StopFleetInput) *AppstreamStopFleetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopFleet, input)
    return &AppstreamStopFleetResult{Result: future}
}
func (a *AppStreamStub) StopImageBuilder(ctx workflow.Context, input *appstream.StopImageBuilderInput) (*appstream.StopImageBuilderOutput, error) {
    var output appstream.StopImageBuilderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopImageBuilder, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) StopImageBuilderAsync(ctx workflow.Context, input *appstream.StopImageBuilderInput) *AppstreamStopImageBuilderResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopImageBuilder, input)
    return &AppstreamStopImageBuilderResult{Result: future}
}
func (a *AppStreamStub) TagResource(ctx workflow.Context, input *appstream.TagResourceInput) (*appstream.TagResourceOutput, error) {
    var output appstream.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) TagResourceAsync(ctx workflow.Context, input *appstream.TagResourceInput) *AppstreamTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &AppstreamTagResourceResult{Result: future}
}
func (a *AppStreamStub) UntagResource(ctx workflow.Context, input *appstream.UntagResourceInput) (*appstream.UntagResourceOutput, error) {
    var output appstream.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) UntagResourceAsync(ctx workflow.Context, input *appstream.UntagResourceInput) *AppstreamUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &AppstreamUntagResourceResult{Result: future}
}
func (a *AppStreamStub) UpdateDirectoryConfig(ctx workflow.Context, input *appstream.UpdateDirectoryConfigInput) (*appstream.UpdateDirectoryConfigOutput, error) {
    var output appstream.UpdateDirectoryConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDirectoryConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) UpdateDirectoryConfigAsync(ctx workflow.Context, input *appstream.UpdateDirectoryConfigInput) *AppstreamUpdateDirectoryConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateDirectoryConfig, input)
    return &AppstreamUpdateDirectoryConfigResult{Result: future}
}
func (a *AppStreamStub) UpdateFleet(ctx workflow.Context, input *appstream.UpdateFleetInput) (*appstream.UpdateFleetOutput, error) {
    var output appstream.UpdateFleetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFleet, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) UpdateFleetAsync(ctx workflow.Context, input *appstream.UpdateFleetInput) *AppstreamUpdateFleetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateFleet, input)
    return &AppstreamUpdateFleetResult{Result: future}
}
func (a *AppStreamStub) UpdateImagePermissions(ctx workflow.Context, input *appstream.UpdateImagePermissionsInput) (*appstream.UpdateImagePermissionsOutput, error) {
    var output appstream.UpdateImagePermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateImagePermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) UpdateImagePermissionsAsync(ctx workflow.Context, input *appstream.UpdateImagePermissionsInput) *AppstreamUpdateImagePermissionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateImagePermissions, input)
    return &AppstreamUpdateImagePermissionsResult{Result: future}
}
func (a *AppStreamStub) UpdateStack(ctx workflow.Context, input *appstream.UpdateStackInput) (*appstream.UpdateStackOutput, error) {
    var output appstream.UpdateStackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateStack, input).Get(ctx, &output)
    return &output, err
}

func (a *AppStreamStub) UpdateStackAsync(ctx workflow.Context, input *appstream.UpdateStackInput) *AppstreamUpdateStackResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateStack, input)
    return &AppstreamUpdateStackResult{Result: future}
}

func (a *AppStreamStub) WaitUntilFleetStarted(ctx workflow.Context, input *appstream.DescribeFleetsInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilFleetStarted, input).Get(ctx, nil)
}

func (a *AppStreamStub) WaitUntilFleetStartedAsync(ctx workflow.Context, input *appstream.DescribeFleetsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilFleetStarted, input)
}

func (a *AppStreamStub) WaitUntilFleetStopped(ctx workflow.Context, input *appstream.DescribeFleetsInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilFleetStopped, input).Get(ctx, nil)
}

func (a *AppStreamStub) WaitUntilFleetStoppedAsync(ctx workflow.Context, input *appstream.DescribeFleetsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilFleetStopped, input)
}