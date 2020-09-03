package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/kinesisvideo"
	"github.com/aws/aws-sdk-go/service/kinesisvideo/kinesisvideoiface"
)

type KinesisVideoActivities struct {
	client kinesisvideoiface.KinesisVideoAPI
}

func NewKinesisVideoActivities(client kinesisvideoiface.KinesisVideoAPI) *KinesisVideoActivities {
    return &KinesisVideoActivities{client: client}
}


func (a *KinesisVideoActivities) CreateSignalingChannel(input *kinesisvideo.CreateSignalingChannelInput) (*kinesisvideo.CreateSignalingChannelOutput, error) {
    return a.client.CreateSignalingChannel(input)
}



func (a *KinesisVideoActivities) CreateStream(input *kinesisvideo.CreateStreamInput) (*kinesisvideo.CreateStreamOutput, error) {
    return a.client.CreateStream(input)
}



func (a *KinesisVideoActivities) DeleteSignalingChannel(input *kinesisvideo.DeleteSignalingChannelInput) (*kinesisvideo.DeleteSignalingChannelOutput, error) {
    return a.client.DeleteSignalingChannel(input)
}



func (a *KinesisVideoActivities) DeleteStream(input *kinesisvideo.DeleteStreamInput) (*kinesisvideo.DeleteStreamOutput, error) {
    return a.client.DeleteStream(input)
}



func (a *KinesisVideoActivities) DescribeSignalingChannel(input *kinesisvideo.DescribeSignalingChannelInput) (*kinesisvideo.DescribeSignalingChannelOutput, error) {
    return a.client.DescribeSignalingChannel(input)
}



func (a *KinesisVideoActivities) DescribeStream(input *kinesisvideo.DescribeStreamInput) (*kinesisvideo.DescribeStreamOutput, error) {
    return a.client.DescribeStream(input)
}



func (a *KinesisVideoActivities) GetDataEndpoint(input *kinesisvideo.GetDataEndpointInput) (*kinesisvideo.GetDataEndpointOutput, error) {
    return a.client.GetDataEndpoint(input)
}



func (a *KinesisVideoActivities) GetSignalingChannelEndpoint(input *kinesisvideo.GetSignalingChannelEndpointInput) (*kinesisvideo.GetSignalingChannelEndpointOutput, error) {
    return a.client.GetSignalingChannelEndpoint(input)
}



func (a *KinesisVideoActivities) ListSignalingChannels(input *kinesisvideo.ListSignalingChannelsInput) (*kinesisvideo.ListSignalingChannelsOutput, error) {
    return a.client.ListSignalingChannels(input)
}



func (a *KinesisVideoActivities) ListStreams(input *kinesisvideo.ListStreamsInput) (*kinesisvideo.ListStreamsOutput, error) {
    return a.client.ListStreams(input)
}



func (a *KinesisVideoActivities) ListTagsForResource(input *kinesisvideo.ListTagsForResourceInput) (*kinesisvideo.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}



func (a *KinesisVideoActivities) ListTagsForStream(input *kinesisvideo.ListTagsForStreamInput) (*kinesisvideo.ListTagsForStreamOutput, error) {
    return a.client.ListTagsForStream(input)
}



func (a *KinesisVideoActivities) TagResource(input *kinesisvideo.TagResourceInput) (*kinesisvideo.TagResourceOutput, error) {
    return a.client.TagResource(input)
}



func (a *KinesisVideoActivities) TagStream(input *kinesisvideo.TagStreamInput) (*kinesisvideo.TagStreamOutput, error) {
    return a.client.TagStream(input)
}



func (a *KinesisVideoActivities) UntagResource(input *kinesisvideo.UntagResourceInput) (*kinesisvideo.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}



func (a *KinesisVideoActivities) UntagStream(input *kinesisvideo.UntagStreamInput) (*kinesisvideo.UntagStreamOutput, error) {
    return a.client.UntagStream(input)
}



func (a *KinesisVideoActivities) UpdateDataRetention(input *kinesisvideo.UpdateDataRetentionInput) (*kinesisvideo.UpdateDataRetentionOutput, error) {
    return a.client.UpdateDataRetention(input)
}



func (a *KinesisVideoActivities) UpdateSignalingChannel(input *kinesisvideo.UpdateSignalingChannelInput) (*kinesisvideo.UpdateSignalingChannelOutput, error) {
    return a.client.UpdateSignalingChannel(input)
}



func (a *KinesisVideoActivities) UpdateStream(input *kinesisvideo.UpdateStreamInput) (*kinesisvideo.UpdateStreamOutput, error) {
    return a.client.UpdateStream(input)
}

