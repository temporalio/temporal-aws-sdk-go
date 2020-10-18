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
	BatchGetChannelAsync(ctx workflow.Context, input *ivs.BatchGetChannelInput) *IVSBatchGetChannelFuture

	BatchGetStreamKey(ctx workflow.Context, input *ivs.BatchGetStreamKeyInput) (*ivs.BatchGetStreamKeyOutput, error)
	BatchGetStreamKeyAsync(ctx workflow.Context, input *ivs.BatchGetStreamKeyInput) *IVSBatchGetStreamKeyFuture

	CreateChannel(ctx workflow.Context, input *ivs.CreateChannelInput) (*ivs.CreateChannelOutput, error)
	CreateChannelAsync(ctx workflow.Context, input *ivs.CreateChannelInput) *IVSCreateChannelFuture

	CreateStreamKey(ctx workflow.Context, input *ivs.CreateStreamKeyInput) (*ivs.CreateStreamKeyOutput, error)
	CreateStreamKeyAsync(ctx workflow.Context, input *ivs.CreateStreamKeyInput) *IVSCreateStreamKeyFuture

	DeleteChannel(ctx workflow.Context, input *ivs.DeleteChannelInput) (*ivs.DeleteChannelOutput, error)
	DeleteChannelAsync(ctx workflow.Context, input *ivs.DeleteChannelInput) *IVSDeleteChannelFuture

	DeletePlaybackKeyPair(ctx workflow.Context, input *ivs.DeletePlaybackKeyPairInput) (*ivs.DeletePlaybackKeyPairOutput, error)
	DeletePlaybackKeyPairAsync(ctx workflow.Context, input *ivs.DeletePlaybackKeyPairInput) *IVSDeletePlaybackKeyPairFuture

	DeleteStreamKey(ctx workflow.Context, input *ivs.DeleteStreamKeyInput) (*ivs.DeleteStreamKeyOutput, error)
	DeleteStreamKeyAsync(ctx workflow.Context, input *ivs.DeleteStreamKeyInput) *IVSDeleteStreamKeyFuture

	GetChannel(ctx workflow.Context, input *ivs.GetChannelInput) (*ivs.GetChannelOutput, error)
	GetChannelAsync(ctx workflow.Context, input *ivs.GetChannelInput) *IVSGetChannelFuture

	GetPlaybackKeyPair(ctx workflow.Context, input *ivs.GetPlaybackKeyPairInput) (*ivs.GetPlaybackKeyPairOutput, error)
	GetPlaybackKeyPairAsync(ctx workflow.Context, input *ivs.GetPlaybackKeyPairInput) *IVSGetPlaybackKeyPairFuture

	GetStream(ctx workflow.Context, input *ivs.GetStreamInput) (*ivs.GetStreamOutput, error)
	GetStreamAsync(ctx workflow.Context, input *ivs.GetStreamInput) *IVSGetStreamFuture

	GetStreamKey(ctx workflow.Context, input *ivs.GetStreamKeyInput) (*ivs.GetStreamKeyOutput, error)
	GetStreamKeyAsync(ctx workflow.Context, input *ivs.GetStreamKeyInput) *IVSGetStreamKeyFuture

	ImportPlaybackKeyPair(ctx workflow.Context, input *ivs.ImportPlaybackKeyPairInput) (*ivs.ImportPlaybackKeyPairOutput, error)
	ImportPlaybackKeyPairAsync(ctx workflow.Context, input *ivs.ImportPlaybackKeyPairInput) *IVSImportPlaybackKeyPairFuture

	ListChannels(ctx workflow.Context, input *ivs.ListChannelsInput) (*ivs.ListChannelsOutput, error)
	ListChannelsAsync(ctx workflow.Context, input *ivs.ListChannelsInput) *IVSListChannelsFuture

	ListPlaybackKeyPairs(ctx workflow.Context, input *ivs.ListPlaybackKeyPairsInput) (*ivs.ListPlaybackKeyPairsOutput, error)
	ListPlaybackKeyPairsAsync(ctx workflow.Context, input *ivs.ListPlaybackKeyPairsInput) *IVSListPlaybackKeyPairsFuture

	ListStreamKeys(ctx workflow.Context, input *ivs.ListStreamKeysInput) (*ivs.ListStreamKeysOutput, error)
	ListStreamKeysAsync(ctx workflow.Context, input *ivs.ListStreamKeysInput) *IVSListStreamKeysFuture

	ListStreams(ctx workflow.Context, input *ivs.ListStreamsInput) (*ivs.ListStreamsOutput, error)
	ListStreamsAsync(ctx workflow.Context, input *ivs.ListStreamsInput) *IVSListStreamsFuture

	ListTagsForResource(ctx workflow.Context, input *ivs.ListTagsForResourceInput) (*ivs.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *ivs.ListTagsForResourceInput) *IVSListTagsForResourceFuture

	PutMetadata(ctx workflow.Context, input *ivs.PutMetadataInput) (*ivs.PutMetadataOutput, error)
	PutMetadataAsync(ctx workflow.Context, input *ivs.PutMetadataInput) *IVSPutMetadataFuture

	StopStream(ctx workflow.Context, input *ivs.StopStreamInput) (*ivs.StopStreamOutput, error)
	StopStreamAsync(ctx workflow.Context, input *ivs.StopStreamInput) *IVSStopStreamFuture

	TagResource(ctx workflow.Context, input *ivs.TagResourceInput) (*ivs.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *ivs.TagResourceInput) *IVSTagResourceFuture

	UntagResource(ctx workflow.Context, input *ivs.UntagResourceInput) (*ivs.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *ivs.UntagResourceInput) *IVSUntagResourceFuture

	UpdateChannel(ctx workflow.Context, input *ivs.UpdateChannelInput) (*ivs.UpdateChannelOutput, error)
	UpdateChannelAsync(ctx workflow.Context, input *ivs.UpdateChannelInput) *IVSUpdateChannelFuture
}

func NewClient() Client {
	return &stub{}
}
