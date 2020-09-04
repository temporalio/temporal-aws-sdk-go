
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/mediapackage"
	"github.com/aws/aws-sdk-go/service/mediapackage/mediapackageiface"
)

type MediaPackageActivities struct {
	client mediapackageiface.MediaPackageAPI
}

func NewMediaPackageActivities(client mediapackageiface.MediaPackageAPI) *MediaPackageActivities {
    return &MediaPackageActivities{client: client}
}

func (a *MediaPackageActivities) CreateChannel(input *mediapackage.CreateChannelInput) (*mediapackage.CreateChannelOutput, error) {
    return a.client.CreateChannel(input)
}

func (a *MediaPackageActivities) CreateHarvestJob(input *mediapackage.CreateHarvestJobInput) (*mediapackage.CreateHarvestJobOutput, error) {
    return a.client.CreateHarvestJob(input)
}

func (a *MediaPackageActivities) CreateOriginEndpoint(input *mediapackage.CreateOriginEndpointInput) (*mediapackage.CreateOriginEndpointOutput, error) {
    return a.client.CreateOriginEndpoint(input)
}

func (a *MediaPackageActivities) DeleteChannel(input *mediapackage.DeleteChannelInput) (*mediapackage.DeleteChannelOutput, error) {
    return a.client.DeleteChannel(input)
}

func (a *MediaPackageActivities) DeleteOriginEndpoint(input *mediapackage.DeleteOriginEndpointInput) (*mediapackage.DeleteOriginEndpointOutput, error) {
    return a.client.DeleteOriginEndpoint(input)
}

func (a *MediaPackageActivities) DescribeChannel(input *mediapackage.DescribeChannelInput) (*mediapackage.DescribeChannelOutput, error) {
    return a.client.DescribeChannel(input)
}

func (a *MediaPackageActivities) DescribeHarvestJob(input *mediapackage.DescribeHarvestJobInput) (*mediapackage.DescribeHarvestJobOutput, error) {
    return a.client.DescribeHarvestJob(input)
}

func (a *MediaPackageActivities) DescribeOriginEndpoint(input *mediapackage.DescribeOriginEndpointInput) (*mediapackage.DescribeOriginEndpointOutput, error) {
    return a.client.DescribeOriginEndpoint(input)
}

func (a *MediaPackageActivities) ListChannels(input *mediapackage.ListChannelsInput) (*mediapackage.ListChannelsOutput, error) {
    return a.client.ListChannels(input)
}

func (a *MediaPackageActivities) ListHarvestJobs(input *mediapackage.ListHarvestJobsInput) (*mediapackage.ListHarvestJobsOutput, error) {
    return a.client.ListHarvestJobs(input)
}

func (a *MediaPackageActivities) ListOriginEndpoints(input *mediapackage.ListOriginEndpointsInput) (*mediapackage.ListOriginEndpointsOutput, error) {
    return a.client.ListOriginEndpoints(input)
}

func (a *MediaPackageActivities) ListTagsForResource(input *mediapackage.ListTagsForResourceInput) (*mediapackage.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *MediaPackageActivities) RotateChannelCredentials(input *mediapackage.RotateChannelCredentialsInput) (*mediapackage.RotateChannelCredentialsOutput, error) {
    return a.client.RotateChannelCredentials(input)
}

func (a *MediaPackageActivities) RotateIngestEndpointCredentials(input *mediapackage.RotateIngestEndpointCredentialsInput) (*mediapackage.RotateIngestEndpointCredentialsOutput, error) {
    return a.client.RotateIngestEndpointCredentials(input)
}

func (a *MediaPackageActivities) TagResource(input *mediapackage.TagResourceInput) (*mediapackage.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *MediaPackageActivities) UntagResource(input *mediapackage.UntagResourceInput) (*mediapackage.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *MediaPackageActivities) UpdateChannel(input *mediapackage.UpdateChannelInput) (*mediapackage.UpdateChannelOutput, error) {
    return a.client.UpdateChannel(input)
}

func (a *MediaPackageActivities) UpdateOriginEndpoint(input *mediapackage.UpdateOriginEndpointInput) (*mediapackage.UpdateOriginEndpointOutput, error) {
    return a.client.UpdateOriginEndpoint(input)
}
