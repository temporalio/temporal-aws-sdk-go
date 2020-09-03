package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/mediapackage"
	"go.temporal.io/sdk/workflow"
)

type MediaPackageClient interface {
    CreateChannel(ctx workflow.Context, input *mediapackage.CreateChannelInput) (*mediapackage.CreateChannelOutput, error)
    CreateChannelAsync(ctx workflow.Context, input *mediapackage.CreateChannelInput) *MediapackageCreateChannelResult

    CreateHarvestJob(ctx workflow.Context, input *mediapackage.CreateHarvestJobInput) (*mediapackage.CreateHarvestJobOutput, error)
    CreateHarvestJobAsync(ctx workflow.Context, input *mediapackage.CreateHarvestJobInput) *MediapackageCreateHarvestJobResult

    CreateOriginEndpoint(ctx workflow.Context, input *mediapackage.CreateOriginEndpointInput) (*mediapackage.CreateOriginEndpointOutput, error)
    CreateOriginEndpointAsync(ctx workflow.Context, input *mediapackage.CreateOriginEndpointInput) *MediapackageCreateOriginEndpointResult

    DeleteChannel(ctx workflow.Context, input *mediapackage.DeleteChannelInput) (*mediapackage.DeleteChannelOutput, error)
    DeleteChannelAsync(ctx workflow.Context, input *mediapackage.DeleteChannelInput) *MediapackageDeleteChannelResult

    DeleteOriginEndpoint(ctx workflow.Context, input *mediapackage.DeleteOriginEndpointInput) (*mediapackage.DeleteOriginEndpointOutput, error)
    DeleteOriginEndpointAsync(ctx workflow.Context, input *mediapackage.DeleteOriginEndpointInput) *MediapackageDeleteOriginEndpointResult

    DescribeChannel(ctx workflow.Context, input *mediapackage.DescribeChannelInput) (*mediapackage.DescribeChannelOutput, error)
    DescribeChannelAsync(ctx workflow.Context, input *mediapackage.DescribeChannelInput) *MediapackageDescribeChannelResult

    DescribeHarvestJob(ctx workflow.Context, input *mediapackage.DescribeHarvestJobInput) (*mediapackage.DescribeHarvestJobOutput, error)
    DescribeHarvestJobAsync(ctx workflow.Context, input *mediapackage.DescribeHarvestJobInput) *MediapackageDescribeHarvestJobResult

    DescribeOriginEndpoint(ctx workflow.Context, input *mediapackage.DescribeOriginEndpointInput) (*mediapackage.DescribeOriginEndpointOutput, error)
    DescribeOriginEndpointAsync(ctx workflow.Context, input *mediapackage.DescribeOriginEndpointInput) *MediapackageDescribeOriginEndpointResult

    ListChannels(ctx workflow.Context, input *mediapackage.ListChannelsInput) (*mediapackage.ListChannelsOutput, error)
    ListChannelsAsync(ctx workflow.Context, input *mediapackage.ListChannelsInput) *MediapackageListChannelsResult

    ListHarvestJobs(ctx workflow.Context, input *mediapackage.ListHarvestJobsInput) (*mediapackage.ListHarvestJobsOutput, error)
    ListHarvestJobsAsync(ctx workflow.Context, input *mediapackage.ListHarvestJobsInput) *MediapackageListHarvestJobsResult

    ListOriginEndpoints(ctx workflow.Context, input *mediapackage.ListOriginEndpointsInput) (*mediapackage.ListOriginEndpointsOutput, error)
    ListOriginEndpointsAsync(ctx workflow.Context, input *mediapackage.ListOriginEndpointsInput) *MediapackageListOriginEndpointsResult

    ListTagsForResource(ctx workflow.Context, input *mediapackage.ListTagsForResourceInput) (*mediapackage.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *mediapackage.ListTagsForResourceInput) *MediapackageListTagsForResourceResult

    RotateChannelCredentials(ctx workflow.Context, input *mediapackage.RotateChannelCredentialsInput) (*mediapackage.RotateChannelCredentialsOutput, error)
    RotateChannelCredentialsAsync(ctx workflow.Context, input *mediapackage.RotateChannelCredentialsInput) *MediapackageRotateChannelCredentialsResult

    RotateIngestEndpointCredentials(ctx workflow.Context, input *mediapackage.RotateIngestEndpointCredentialsInput) (*mediapackage.RotateIngestEndpointCredentialsOutput, error)
    RotateIngestEndpointCredentialsAsync(ctx workflow.Context, input *mediapackage.RotateIngestEndpointCredentialsInput) *MediapackageRotateIngestEndpointCredentialsResult

    TagResource(ctx workflow.Context, input *mediapackage.TagResourceInput) (*mediapackage.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *mediapackage.TagResourceInput) *MediapackageTagResourceResult

    UntagResource(ctx workflow.Context, input *mediapackage.UntagResourceInput) (*mediapackage.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *mediapackage.UntagResourceInput) *MediapackageUntagResourceResult

    UpdateChannel(ctx workflow.Context, input *mediapackage.UpdateChannelInput) (*mediapackage.UpdateChannelOutput, error)
    UpdateChannelAsync(ctx workflow.Context, input *mediapackage.UpdateChannelInput) *MediapackageUpdateChannelResult

    UpdateOriginEndpoint(ctx workflow.Context, input *mediapackage.UpdateOriginEndpointInput) (*mediapackage.UpdateOriginEndpointOutput, error)
    UpdateOriginEndpointAsync(ctx workflow.Context, input *mediapackage.UpdateOriginEndpointInput) *MediapackageUpdateOriginEndpointResult
}
type MediapackageCreateChannelResult struct {
	Result workflow.Future
}

func (r *MediapackageCreateChannelResult) Get(ctx workflow.Context) (*mediapackage.CreateChannelOutput, error) {
    var output mediapackage.CreateChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackageCreateHarvestJobResult struct {
	Result workflow.Future
}

func (r *MediapackageCreateHarvestJobResult) Get(ctx workflow.Context) (*mediapackage.CreateHarvestJobOutput, error) {
    var output mediapackage.CreateHarvestJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackageCreateOriginEndpointResult struct {
	Result workflow.Future
}

func (r *MediapackageCreateOriginEndpointResult) Get(ctx workflow.Context) (*mediapackage.CreateOriginEndpointOutput, error) {
    var output mediapackage.CreateOriginEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackageDeleteChannelResult struct {
	Result workflow.Future
}

func (r *MediapackageDeleteChannelResult) Get(ctx workflow.Context) (*mediapackage.DeleteChannelOutput, error) {
    var output mediapackage.DeleteChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackageDeleteOriginEndpointResult struct {
	Result workflow.Future
}

func (r *MediapackageDeleteOriginEndpointResult) Get(ctx workflow.Context) (*mediapackage.DeleteOriginEndpointOutput, error) {
    var output mediapackage.DeleteOriginEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackageDescribeChannelResult struct {
	Result workflow.Future
}

func (r *MediapackageDescribeChannelResult) Get(ctx workflow.Context) (*mediapackage.DescribeChannelOutput, error) {
    var output mediapackage.DescribeChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackageDescribeHarvestJobResult struct {
	Result workflow.Future
}

func (r *MediapackageDescribeHarvestJobResult) Get(ctx workflow.Context) (*mediapackage.DescribeHarvestJobOutput, error) {
    var output mediapackage.DescribeHarvestJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackageDescribeOriginEndpointResult struct {
	Result workflow.Future
}

func (r *MediapackageDescribeOriginEndpointResult) Get(ctx workflow.Context) (*mediapackage.DescribeOriginEndpointOutput, error) {
    var output mediapackage.DescribeOriginEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackageListChannelsResult struct {
	Result workflow.Future
}

func (r *MediapackageListChannelsResult) Get(ctx workflow.Context) (*mediapackage.ListChannelsOutput, error) {
    var output mediapackage.ListChannelsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackageListHarvestJobsResult struct {
	Result workflow.Future
}

func (r *MediapackageListHarvestJobsResult) Get(ctx workflow.Context) (*mediapackage.ListHarvestJobsOutput, error) {
    var output mediapackage.ListHarvestJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackageListOriginEndpointsResult struct {
	Result workflow.Future
}

func (r *MediapackageListOriginEndpointsResult) Get(ctx workflow.Context) (*mediapackage.ListOriginEndpointsOutput, error) {
    var output mediapackage.ListOriginEndpointsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackageListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *MediapackageListTagsForResourceResult) Get(ctx workflow.Context) (*mediapackage.ListTagsForResourceOutput, error) {
    var output mediapackage.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackageRotateChannelCredentialsResult struct {
	Result workflow.Future
}

func (r *MediapackageRotateChannelCredentialsResult) Get(ctx workflow.Context) (*mediapackage.RotateChannelCredentialsOutput, error) {
    var output mediapackage.RotateChannelCredentialsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackageRotateIngestEndpointCredentialsResult struct {
	Result workflow.Future
}

func (r *MediapackageRotateIngestEndpointCredentialsResult) Get(ctx workflow.Context) (*mediapackage.RotateIngestEndpointCredentialsOutput, error) {
    var output mediapackage.RotateIngestEndpointCredentialsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackageTagResourceResult struct {
	Result workflow.Future
}

func (r *MediapackageTagResourceResult) Get(ctx workflow.Context) (*mediapackage.TagResourceOutput, error) {
    var output mediapackage.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackageUntagResourceResult struct {
	Result workflow.Future
}

func (r *MediapackageUntagResourceResult) Get(ctx workflow.Context) (*mediapackage.UntagResourceOutput, error) {
    var output mediapackage.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackageUpdateChannelResult struct {
	Result workflow.Future
}

func (r *MediapackageUpdateChannelResult) Get(ctx workflow.Context) (*mediapackage.UpdateChannelOutput, error) {
    var output mediapackage.UpdateChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediapackageUpdateOriginEndpointResult struct {
	Result workflow.Future
}

func (r *MediapackageUpdateOriginEndpointResult) Get(ctx workflow.Context) (*mediapackage.UpdateOriginEndpointOutput, error) {
    var output mediapackage.UpdateOriginEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type MediaPackageStub struct {
    activities MediaPackageClient
}

func NewMediaPackageStub() MediaPackageClient {
    return &MediaPackageStub{}
}

func (a *MediaPackageStub) CreateChannel(ctx workflow.Context, input *mediapackage.CreateChannelInput) (*mediapackage.CreateChannelOutput, error) {
    var output mediapackage.CreateChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageStub) CreateHarvestJob(ctx workflow.Context, input *mediapackage.CreateHarvestJobInput) (*mediapackage.CreateHarvestJobOutput, error) {
    var output mediapackage.CreateHarvestJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateHarvestJob, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageStub) CreateOriginEndpoint(ctx workflow.Context, input *mediapackage.CreateOriginEndpointInput) (*mediapackage.CreateOriginEndpointOutput, error) {
    var output mediapackage.CreateOriginEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateOriginEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageStub) DeleteChannel(ctx workflow.Context, input *mediapackage.DeleteChannelInput) (*mediapackage.DeleteChannelOutput, error) {
    var output mediapackage.DeleteChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageStub) DeleteOriginEndpoint(ctx workflow.Context, input *mediapackage.DeleteOriginEndpointInput) (*mediapackage.DeleteOriginEndpointOutput, error) {
    var output mediapackage.DeleteOriginEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteOriginEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageStub) DescribeChannel(ctx workflow.Context, input *mediapackage.DescribeChannelInput) (*mediapackage.DescribeChannelOutput, error) {
    var output mediapackage.DescribeChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageStub) DescribeHarvestJob(ctx workflow.Context, input *mediapackage.DescribeHarvestJobInput) (*mediapackage.DescribeHarvestJobOutput, error) {
    var output mediapackage.DescribeHarvestJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeHarvestJob, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageStub) DescribeOriginEndpoint(ctx workflow.Context, input *mediapackage.DescribeOriginEndpointInput) (*mediapackage.DescribeOriginEndpointOutput, error) {
    var output mediapackage.DescribeOriginEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeOriginEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageStub) ListChannels(ctx workflow.Context, input *mediapackage.ListChannelsInput) (*mediapackage.ListChannelsOutput, error) {
    var output mediapackage.ListChannelsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListChannels, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageStub) ListHarvestJobs(ctx workflow.Context, input *mediapackage.ListHarvestJobsInput) (*mediapackage.ListHarvestJobsOutput, error) {
    var output mediapackage.ListHarvestJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListHarvestJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageStub) ListOriginEndpoints(ctx workflow.Context, input *mediapackage.ListOriginEndpointsInput) (*mediapackage.ListOriginEndpointsOutput, error) {
    var output mediapackage.ListOriginEndpointsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListOriginEndpoints, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageStub) ListTagsForResource(ctx workflow.Context, input *mediapackage.ListTagsForResourceInput) (*mediapackage.ListTagsForResourceOutput, error) {
    var output mediapackage.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageStub) RotateChannelCredentials(ctx workflow.Context, input *mediapackage.RotateChannelCredentialsInput) (*mediapackage.RotateChannelCredentialsOutput, error) {
    var output mediapackage.RotateChannelCredentialsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RotateChannelCredentials, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageStub) RotateIngestEndpointCredentials(ctx workflow.Context, input *mediapackage.RotateIngestEndpointCredentialsInput) (*mediapackage.RotateIngestEndpointCredentialsOutput, error) {
    var output mediapackage.RotateIngestEndpointCredentialsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RotateIngestEndpointCredentials, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageStub) TagResource(ctx workflow.Context, input *mediapackage.TagResourceInput) (*mediapackage.TagResourceOutput, error) {
    var output mediapackage.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageStub) UntagResource(ctx workflow.Context, input *mediapackage.UntagResourceInput) (*mediapackage.UntagResourceOutput, error) {
    var output mediapackage.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageStub) UpdateChannel(ctx workflow.Context, input *mediapackage.UpdateChannelInput) (*mediapackage.UpdateChannelOutput, error) {
    var output mediapackage.UpdateChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaPackageStub) UpdateOriginEndpoint(ctx workflow.Context, input *mediapackage.UpdateOriginEndpointInput) (*mediapackage.UpdateOriginEndpointOutput, error) {
    var output mediapackage.UpdateOriginEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateOriginEndpoint, input).Get(ctx, &output)
    return &output, err
}
