// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package ivsstub

import (
	"github.com/aws/aws-sdk-go/service/ivs"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	BatchGetChannel(ctx workflow.Context, input *ivs.BatchGetChannelInput) (*ivs.BatchGetChannelOutput, error)
	BatchGetChannelAsync(ctx workflow.Context, input *ivs.BatchGetChannelInput) *BatchGetChannelFuture

	BatchGetStreamKey(ctx workflow.Context, input *ivs.BatchGetStreamKeyInput) (*ivs.BatchGetStreamKeyOutput, error)
	BatchGetStreamKeyAsync(ctx workflow.Context, input *ivs.BatchGetStreamKeyInput) *BatchGetStreamKeyFuture

	CreateChannel(ctx workflow.Context, input *ivs.CreateChannelInput) (*ivs.CreateChannelOutput, error)
	CreateChannelAsync(ctx workflow.Context, input *ivs.CreateChannelInput) *CreateChannelFuture

	CreateStreamKey(ctx workflow.Context, input *ivs.CreateStreamKeyInput) (*ivs.CreateStreamKeyOutput, error)
	CreateStreamKeyAsync(ctx workflow.Context, input *ivs.CreateStreamKeyInput) *CreateStreamKeyFuture

	DeleteChannel(ctx workflow.Context, input *ivs.DeleteChannelInput) (*ivs.DeleteChannelOutput, error)
	DeleteChannelAsync(ctx workflow.Context, input *ivs.DeleteChannelInput) *DeleteChannelFuture

	DeletePlaybackKeyPair(ctx workflow.Context, input *ivs.DeletePlaybackKeyPairInput) (*ivs.DeletePlaybackKeyPairOutput, error)
	DeletePlaybackKeyPairAsync(ctx workflow.Context, input *ivs.DeletePlaybackKeyPairInput) *DeletePlaybackKeyPairFuture

	DeleteStreamKey(ctx workflow.Context, input *ivs.DeleteStreamKeyInput) (*ivs.DeleteStreamKeyOutput, error)
	DeleteStreamKeyAsync(ctx workflow.Context, input *ivs.DeleteStreamKeyInput) *DeleteStreamKeyFuture

	GetChannel(ctx workflow.Context, input *ivs.GetChannelInput) (*ivs.GetChannelOutput, error)
	GetChannelAsync(ctx workflow.Context, input *ivs.GetChannelInput) *GetChannelFuture

	GetPlaybackKeyPair(ctx workflow.Context, input *ivs.GetPlaybackKeyPairInput) (*ivs.GetPlaybackKeyPairOutput, error)
	GetPlaybackKeyPairAsync(ctx workflow.Context, input *ivs.GetPlaybackKeyPairInput) *GetPlaybackKeyPairFuture

	GetStream(ctx workflow.Context, input *ivs.GetStreamInput) (*ivs.GetStreamOutput, error)
	GetStreamAsync(ctx workflow.Context, input *ivs.GetStreamInput) *GetStreamFuture

	GetStreamKey(ctx workflow.Context, input *ivs.GetStreamKeyInput) (*ivs.GetStreamKeyOutput, error)
	GetStreamKeyAsync(ctx workflow.Context, input *ivs.GetStreamKeyInput) *GetStreamKeyFuture

	ImportPlaybackKeyPair(ctx workflow.Context, input *ivs.ImportPlaybackKeyPairInput) (*ivs.ImportPlaybackKeyPairOutput, error)
	ImportPlaybackKeyPairAsync(ctx workflow.Context, input *ivs.ImportPlaybackKeyPairInput) *ImportPlaybackKeyPairFuture

	ListChannels(ctx workflow.Context, input *ivs.ListChannelsInput) (*ivs.ListChannelsOutput, error)
	ListChannelsAsync(ctx workflow.Context, input *ivs.ListChannelsInput) *ListChannelsFuture

	ListPlaybackKeyPairs(ctx workflow.Context, input *ivs.ListPlaybackKeyPairsInput) (*ivs.ListPlaybackKeyPairsOutput, error)
	ListPlaybackKeyPairsAsync(ctx workflow.Context, input *ivs.ListPlaybackKeyPairsInput) *ListPlaybackKeyPairsFuture

	ListStreamKeys(ctx workflow.Context, input *ivs.ListStreamKeysInput) (*ivs.ListStreamKeysOutput, error)
	ListStreamKeysAsync(ctx workflow.Context, input *ivs.ListStreamKeysInput) *ListStreamKeysFuture

	ListStreams(ctx workflow.Context, input *ivs.ListStreamsInput) (*ivs.ListStreamsOutput, error)
	ListStreamsAsync(ctx workflow.Context, input *ivs.ListStreamsInput) *ListStreamsFuture

	ListTagsForResource(ctx workflow.Context, input *ivs.ListTagsForResourceInput) (*ivs.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *ivs.ListTagsForResourceInput) *ListTagsForResourceFuture

	PutMetadata(ctx workflow.Context, input *ivs.PutMetadataInput) (*ivs.PutMetadataOutput, error)
	PutMetadataAsync(ctx workflow.Context, input *ivs.PutMetadataInput) *PutMetadataFuture

	StopStream(ctx workflow.Context, input *ivs.StopStreamInput) (*ivs.StopStreamOutput, error)
	StopStreamAsync(ctx workflow.Context, input *ivs.StopStreamInput) *StopStreamFuture

	TagResource(ctx workflow.Context, input *ivs.TagResourceInput) (*ivs.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *ivs.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *ivs.UntagResourceInput) (*ivs.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *ivs.UntagResourceInput) *UntagResourceFuture

	UpdateChannel(ctx workflow.Context, input *ivs.UpdateChannelInput) (*ivs.UpdateChannelOutput, error)
	UpdateChannelAsync(ctx workflow.Context, input *ivs.UpdateChannelInput) *UpdateChannelFuture
}

func NewClient() Client {
	return &stub{}
}
