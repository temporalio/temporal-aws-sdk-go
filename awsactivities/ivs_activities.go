package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/ivs"
	"github.com/aws/aws-sdk-go/service/ivs/ivsiface"
)

type IVSActivities struct {
	client ivsiface.IVSAPI
}

func NewIVSActivities(client ivsiface.IVSAPI) *IVSActivities {
    return &IVSActivities{client: client}
}


func (a *IVSActivities) BatchGetChannel(input *ivs.BatchGetChannelInput) (*ivs.BatchGetChannelOutput, error) {
    return a.client.BatchGetChannel(input)
}



func (a *IVSActivities) BatchGetStreamKey(input *ivs.BatchGetStreamKeyInput) (*ivs.BatchGetStreamKeyOutput, error) {
    return a.client.BatchGetStreamKey(input)
}



func (a *IVSActivities) CreateChannel(input *ivs.CreateChannelInput) (*ivs.CreateChannelOutput, error) {
    return a.client.CreateChannel(input)
}



func (a *IVSActivities) CreateStreamKey(input *ivs.CreateStreamKeyInput) (*ivs.CreateStreamKeyOutput, error) {
    return a.client.CreateStreamKey(input)
}



func (a *IVSActivities) DeleteChannel(input *ivs.DeleteChannelInput) (*ivs.DeleteChannelOutput, error) {
    return a.client.DeleteChannel(input)
}



func (a *IVSActivities) DeletePlaybackKeyPair(input *ivs.DeletePlaybackKeyPairInput) (*ivs.DeletePlaybackKeyPairOutput, error) {
    return a.client.DeletePlaybackKeyPair(input)
}



func (a *IVSActivities) DeleteStreamKey(input *ivs.DeleteStreamKeyInput) (*ivs.DeleteStreamKeyOutput, error) {
    return a.client.DeleteStreamKey(input)
}



func (a *IVSActivities) GetChannel(input *ivs.GetChannelInput) (*ivs.GetChannelOutput, error) {
    return a.client.GetChannel(input)
}



func (a *IVSActivities) GetPlaybackKeyPair(input *ivs.GetPlaybackKeyPairInput) (*ivs.GetPlaybackKeyPairOutput, error) {
    return a.client.GetPlaybackKeyPair(input)
}



func (a *IVSActivities) GetStream(input *ivs.GetStreamInput) (*ivs.GetStreamOutput, error) {
    return a.client.GetStream(input)
}



func (a *IVSActivities) GetStreamKey(input *ivs.GetStreamKeyInput) (*ivs.GetStreamKeyOutput, error) {
    return a.client.GetStreamKey(input)
}



func (a *IVSActivities) ImportPlaybackKeyPair(input *ivs.ImportPlaybackKeyPairInput) (*ivs.ImportPlaybackKeyPairOutput, error) {
    return a.client.ImportPlaybackKeyPair(input)
}



func (a *IVSActivities) ListChannels(input *ivs.ListChannelsInput) (*ivs.ListChannelsOutput, error) {
    return a.client.ListChannels(input)
}



func (a *IVSActivities) ListPlaybackKeyPairs(input *ivs.ListPlaybackKeyPairsInput) (*ivs.ListPlaybackKeyPairsOutput, error) {
    return a.client.ListPlaybackKeyPairs(input)
}



func (a *IVSActivities) ListStreamKeys(input *ivs.ListStreamKeysInput) (*ivs.ListStreamKeysOutput, error) {
    return a.client.ListStreamKeys(input)
}



func (a *IVSActivities) ListStreams(input *ivs.ListStreamsInput) (*ivs.ListStreamsOutput, error) {
    return a.client.ListStreams(input)
}



func (a *IVSActivities) ListTagsForResource(input *ivs.ListTagsForResourceInput) (*ivs.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}



func (a *IVSActivities) PutMetadata(input *ivs.PutMetadataInput) (*ivs.PutMetadataOutput, error) {
    return a.client.PutMetadata(input)
}



func (a *IVSActivities) StopStream(input *ivs.StopStreamInput) (*ivs.StopStreamOutput, error) {
    return a.client.StopStream(input)
}



func (a *IVSActivities) TagResource(input *ivs.TagResourceInput) (*ivs.TagResourceOutput, error) {
    return a.client.TagResource(input)
}



func (a *IVSActivities) UntagResource(input *ivs.UntagResourceInput) (*ivs.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}



func (a *IVSActivities) UpdateChannel(input *ivs.UpdateChannelInput) (*ivs.UpdateChannelOutput, error) {
    return a.client.UpdateChannel(input)
}

