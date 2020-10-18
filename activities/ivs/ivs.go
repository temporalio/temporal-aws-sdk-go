// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package ivs

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.temporal.io/aws-sdk/internal"
	"github.com/aws/aws-sdk-go/service/ivs"
	"github.com/aws/aws-sdk-go/service/ivs/ivsiface"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client ivsiface.IVSAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := ivs.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (ivsiface.IVSAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return ivs.New(sess), nil
}

func (a *Activities) BatchGetChannel(ctx context.Context, input *ivs.BatchGetChannelInput) (*ivs.BatchGetChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchGetChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchGetStreamKey(ctx context.Context, input *ivs.BatchGetStreamKeyInput) (*ivs.BatchGetStreamKeyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchGetStreamKeyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateChannel(ctx context.Context, input *ivs.CreateChannelInput) (*ivs.CreateChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateStreamKey(ctx context.Context, input *ivs.CreateStreamKeyInput) (*ivs.CreateStreamKeyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateStreamKeyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteChannel(ctx context.Context, input *ivs.DeleteChannelInput) (*ivs.DeleteChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeletePlaybackKeyPair(ctx context.Context, input *ivs.DeletePlaybackKeyPairInput) (*ivs.DeletePlaybackKeyPairOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeletePlaybackKeyPairWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteStreamKey(ctx context.Context, input *ivs.DeleteStreamKeyInput) (*ivs.DeleteStreamKeyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteStreamKeyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetChannel(ctx context.Context, input *ivs.GetChannelInput) (*ivs.GetChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetPlaybackKeyPair(ctx context.Context, input *ivs.GetPlaybackKeyPairInput) (*ivs.GetPlaybackKeyPairOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetPlaybackKeyPairWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetStream(ctx context.Context, input *ivs.GetStreamInput) (*ivs.GetStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetStreamWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetStreamKey(ctx context.Context, input *ivs.GetStreamKeyInput) (*ivs.GetStreamKeyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetStreamKeyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ImportPlaybackKeyPair(ctx context.Context, input *ivs.ImportPlaybackKeyPairInput) (*ivs.ImportPlaybackKeyPairOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ImportPlaybackKeyPairWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListChannels(ctx context.Context, input *ivs.ListChannelsInput) (*ivs.ListChannelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListChannelsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListPlaybackKeyPairs(ctx context.Context, input *ivs.ListPlaybackKeyPairsInput) (*ivs.ListPlaybackKeyPairsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListPlaybackKeyPairsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListStreamKeys(ctx context.Context, input *ivs.ListStreamKeysInput) (*ivs.ListStreamKeysOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListStreamKeysWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListStreams(ctx context.Context, input *ivs.ListStreamsInput) (*ivs.ListStreamsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListStreamsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *ivs.ListTagsForResourceInput) (*ivs.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutMetadata(ctx context.Context, input *ivs.PutMetadataInput) (*ivs.PutMetadataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutMetadataWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopStream(ctx context.Context, input *ivs.StopStreamInput) (*ivs.StopStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopStreamWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *ivs.TagResourceInput) (*ivs.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *ivs.UntagResourceInput) (*ivs.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateChannel(ctx context.Context, input *ivs.UpdateChannelInput) (*ivs.UpdateChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateChannelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
