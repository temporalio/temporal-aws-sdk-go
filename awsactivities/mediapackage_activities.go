package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mediapackage"
	"github.com/aws/aws-sdk-go/service/mediapackage/mediapackageiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type MediaPackageActivities struct {
	client mediapackageiface.MediaPackageAPI
}

func NewMediaPackageActivities(session *session.Session, config ...*aws.Config) *MediaPackageActivities {
	client := mediapackage.New(session, config...)
	return &MediaPackageActivities{client: client}
}

func (a *MediaPackageActivities) CreateChannel(ctx context.Context, input *mediapackage.CreateChannelInput) (*mediapackage.CreateChannelOutput, error) {
	return a.client.CreateChannelWithContext(ctx, input)
}

func (a *MediaPackageActivities) CreateHarvestJob(ctx context.Context, input *mediapackage.CreateHarvestJobInput) (*mediapackage.CreateHarvestJobOutput, error) {
	return a.client.CreateHarvestJobWithContext(ctx, input)
}

func (a *MediaPackageActivities) CreateOriginEndpoint(ctx context.Context, input *mediapackage.CreateOriginEndpointInput) (*mediapackage.CreateOriginEndpointOutput, error) {
	return a.client.CreateOriginEndpointWithContext(ctx, input)
}

func (a *MediaPackageActivities) DeleteChannel(ctx context.Context, input *mediapackage.DeleteChannelInput) (*mediapackage.DeleteChannelOutput, error) {
	return a.client.DeleteChannelWithContext(ctx, input)
}

func (a *MediaPackageActivities) DeleteOriginEndpoint(ctx context.Context, input *mediapackage.DeleteOriginEndpointInput) (*mediapackage.DeleteOriginEndpointOutput, error) {
	return a.client.DeleteOriginEndpointWithContext(ctx, input)
}

func (a *MediaPackageActivities) DescribeChannel(ctx context.Context, input *mediapackage.DescribeChannelInput) (*mediapackage.DescribeChannelOutput, error) {
	return a.client.DescribeChannelWithContext(ctx, input)
}

func (a *MediaPackageActivities) DescribeHarvestJob(ctx context.Context, input *mediapackage.DescribeHarvestJobInput) (*mediapackage.DescribeHarvestJobOutput, error) {
	return a.client.DescribeHarvestJobWithContext(ctx, input)
}

func (a *MediaPackageActivities) DescribeOriginEndpoint(ctx context.Context, input *mediapackage.DescribeOriginEndpointInput) (*mediapackage.DescribeOriginEndpointOutput, error) {
	return a.client.DescribeOriginEndpointWithContext(ctx, input)
}

func (a *MediaPackageActivities) ListChannels(ctx context.Context, input *mediapackage.ListChannelsInput) (*mediapackage.ListChannelsOutput, error) {
	return a.client.ListChannelsWithContext(ctx, input)
}

func (a *MediaPackageActivities) ListHarvestJobs(ctx context.Context, input *mediapackage.ListHarvestJobsInput) (*mediapackage.ListHarvestJobsOutput, error) {
	return a.client.ListHarvestJobsWithContext(ctx, input)
}

func (a *MediaPackageActivities) ListOriginEndpoints(ctx context.Context, input *mediapackage.ListOriginEndpointsInput) (*mediapackage.ListOriginEndpointsOutput, error) {
	return a.client.ListOriginEndpointsWithContext(ctx, input)
}

func (a *MediaPackageActivities) ListTagsForResource(ctx context.Context, input *mediapackage.ListTagsForResourceInput) (*mediapackage.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *MediaPackageActivities) RotateChannelCredentials(ctx context.Context, input *mediapackage.RotateChannelCredentialsInput) (*mediapackage.RotateChannelCredentialsOutput, error) {
	return a.client.RotateChannelCredentialsWithContext(ctx, input)
}

func (a *MediaPackageActivities) RotateIngestEndpointCredentials(ctx context.Context, input *mediapackage.RotateIngestEndpointCredentialsInput) (*mediapackage.RotateIngestEndpointCredentialsOutput, error) {
	return a.client.RotateIngestEndpointCredentialsWithContext(ctx, input)
}

func (a *MediaPackageActivities) TagResource(ctx context.Context, input *mediapackage.TagResourceInput) (*mediapackage.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *MediaPackageActivities) UntagResource(ctx context.Context, input *mediapackage.UntagResourceInput) (*mediapackage.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *MediaPackageActivities) UpdateChannel(ctx context.Context, input *mediapackage.UpdateChannelInput) (*mediapackage.UpdateChannelOutput, error) {
	return a.client.UpdateChannelWithContext(ctx, input)
}

func (a *MediaPackageActivities) UpdateOriginEndpoint(ctx context.Context, input *mediapackage.UpdateOriginEndpointInput) (*mediapackage.UpdateOriginEndpointOutput, error) {
	return a.client.UpdateOriginEndpointWithContext(ctx, input)
}
