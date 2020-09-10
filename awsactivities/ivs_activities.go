package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ivs"
	"github.com/aws/aws-sdk-go/service/ivs/ivsiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type IVSActivities struct {
	client ivsiface.IVSAPI
}

func NewIVSActivities(session *session.Session, config ...*aws.Config) *IVSActivities {
	client := ivs.New(session, config...)
	return &IVSActivities{client: client}
}

func (a *IVSActivities) BatchGetChannel(ctx context.Context, input *ivs.BatchGetChannelInput) (*ivs.BatchGetChannelOutput, error) {
	return a.client.BatchGetChannelWithContext(ctx, input)
}

func (a *IVSActivities) BatchGetStreamKey(ctx context.Context, input *ivs.BatchGetStreamKeyInput) (*ivs.BatchGetStreamKeyOutput, error) {
	return a.client.BatchGetStreamKeyWithContext(ctx, input)
}

func (a *IVSActivities) CreateChannel(ctx context.Context, input *ivs.CreateChannelInput) (*ivs.CreateChannelOutput, error) {
	return a.client.CreateChannelWithContext(ctx, input)
}

func (a *IVSActivities) CreateStreamKey(ctx context.Context, input *ivs.CreateStreamKeyInput) (*ivs.CreateStreamKeyOutput, error) {
	return a.client.CreateStreamKeyWithContext(ctx, input)
}

func (a *IVSActivities) DeleteChannel(ctx context.Context, input *ivs.DeleteChannelInput) (*ivs.DeleteChannelOutput, error) {
	return a.client.DeleteChannelWithContext(ctx, input)
}

func (a *IVSActivities) DeletePlaybackKeyPair(ctx context.Context, input *ivs.DeletePlaybackKeyPairInput) (*ivs.DeletePlaybackKeyPairOutput, error) {
	return a.client.DeletePlaybackKeyPairWithContext(ctx, input)
}

func (a *IVSActivities) DeleteStreamKey(ctx context.Context, input *ivs.DeleteStreamKeyInput) (*ivs.DeleteStreamKeyOutput, error) {
	return a.client.DeleteStreamKeyWithContext(ctx, input)
}

func (a *IVSActivities) GetChannel(ctx context.Context, input *ivs.GetChannelInput) (*ivs.GetChannelOutput, error) {
	return a.client.GetChannelWithContext(ctx, input)
}

func (a *IVSActivities) GetPlaybackKeyPair(ctx context.Context, input *ivs.GetPlaybackKeyPairInput) (*ivs.GetPlaybackKeyPairOutput, error) {
	return a.client.GetPlaybackKeyPairWithContext(ctx, input)
}

func (a *IVSActivities) GetStream(ctx context.Context, input *ivs.GetStreamInput) (*ivs.GetStreamOutput, error) {
	return a.client.GetStreamWithContext(ctx, input)
}

func (a *IVSActivities) GetStreamKey(ctx context.Context, input *ivs.GetStreamKeyInput) (*ivs.GetStreamKeyOutput, error) {
	return a.client.GetStreamKeyWithContext(ctx, input)
}

func (a *IVSActivities) ImportPlaybackKeyPair(ctx context.Context, input *ivs.ImportPlaybackKeyPairInput) (*ivs.ImportPlaybackKeyPairOutput, error) {
	return a.client.ImportPlaybackKeyPairWithContext(ctx, input)
}

func (a *IVSActivities) ListChannels(ctx context.Context, input *ivs.ListChannelsInput) (*ivs.ListChannelsOutput, error) {
	return a.client.ListChannelsWithContext(ctx, input)
}

func (a *IVSActivities) ListPlaybackKeyPairs(ctx context.Context, input *ivs.ListPlaybackKeyPairsInput) (*ivs.ListPlaybackKeyPairsOutput, error) {
	return a.client.ListPlaybackKeyPairsWithContext(ctx, input)
}

func (a *IVSActivities) ListStreamKeys(ctx context.Context, input *ivs.ListStreamKeysInput) (*ivs.ListStreamKeysOutput, error) {
	return a.client.ListStreamKeysWithContext(ctx, input)
}

func (a *IVSActivities) ListStreams(ctx context.Context, input *ivs.ListStreamsInput) (*ivs.ListStreamsOutput, error) {
	return a.client.ListStreamsWithContext(ctx, input)
}

func (a *IVSActivities) ListTagsForResource(ctx context.Context, input *ivs.ListTagsForResourceInput) (*ivs.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *IVSActivities) PutMetadata(ctx context.Context, input *ivs.PutMetadataInput) (*ivs.PutMetadataOutput, error) {
	return a.client.PutMetadataWithContext(ctx, input)
}

func (a *IVSActivities) StopStream(ctx context.Context, input *ivs.StopStreamInput) (*ivs.StopStreamOutput, error) {
	return a.client.StopStreamWithContext(ctx, input)
}

func (a *IVSActivities) TagResource(ctx context.Context, input *ivs.TagResourceInput) (*ivs.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *IVSActivities) UntagResource(ctx context.Context, input *ivs.UntagResourceInput) (*ivs.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *IVSActivities) UpdateChannel(ctx context.Context, input *ivs.UpdateChannelInput) (*ivs.UpdateChannelOutput, error) {
	return a.client.UpdateChannelWithContext(ctx, input)
}
