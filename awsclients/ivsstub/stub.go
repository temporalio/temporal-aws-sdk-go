// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package ivsstub

import (
	"github.com/aws/aws-sdk-go/service/ivs"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type IVSBatchGetChannelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IVSBatchGetChannelFuture) Get(ctx workflow.Context) (*ivs.BatchGetChannelOutput, error) {
	var output ivs.BatchGetChannelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IVSBatchGetStreamKeyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IVSBatchGetStreamKeyFuture) Get(ctx workflow.Context) (*ivs.BatchGetStreamKeyOutput, error) {
	var output ivs.BatchGetStreamKeyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IVSCreateChannelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IVSCreateChannelFuture) Get(ctx workflow.Context) (*ivs.CreateChannelOutput, error) {
	var output ivs.CreateChannelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IVSCreateStreamKeyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IVSCreateStreamKeyFuture) Get(ctx workflow.Context) (*ivs.CreateStreamKeyOutput, error) {
	var output ivs.CreateStreamKeyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IVSDeleteChannelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IVSDeleteChannelFuture) Get(ctx workflow.Context) (*ivs.DeleteChannelOutput, error) {
	var output ivs.DeleteChannelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IVSDeletePlaybackKeyPairFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IVSDeletePlaybackKeyPairFuture) Get(ctx workflow.Context) (*ivs.DeletePlaybackKeyPairOutput, error) {
	var output ivs.DeletePlaybackKeyPairOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IVSDeleteStreamKeyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IVSDeleteStreamKeyFuture) Get(ctx workflow.Context) (*ivs.DeleteStreamKeyOutput, error) {
	var output ivs.DeleteStreamKeyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IVSGetChannelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IVSGetChannelFuture) Get(ctx workflow.Context) (*ivs.GetChannelOutput, error) {
	var output ivs.GetChannelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IVSGetPlaybackKeyPairFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IVSGetPlaybackKeyPairFuture) Get(ctx workflow.Context) (*ivs.GetPlaybackKeyPairOutput, error) {
	var output ivs.GetPlaybackKeyPairOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IVSGetStreamFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IVSGetStreamFuture) Get(ctx workflow.Context) (*ivs.GetStreamOutput, error) {
	var output ivs.GetStreamOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IVSGetStreamKeyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IVSGetStreamKeyFuture) Get(ctx workflow.Context) (*ivs.GetStreamKeyOutput, error) {
	var output ivs.GetStreamKeyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IVSImportPlaybackKeyPairFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IVSImportPlaybackKeyPairFuture) Get(ctx workflow.Context) (*ivs.ImportPlaybackKeyPairOutput, error) {
	var output ivs.ImportPlaybackKeyPairOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IVSListChannelsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IVSListChannelsFuture) Get(ctx workflow.Context) (*ivs.ListChannelsOutput, error) {
	var output ivs.ListChannelsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IVSListPlaybackKeyPairsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IVSListPlaybackKeyPairsFuture) Get(ctx workflow.Context) (*ivs.ListPlaybackKeyPairsOutput, error) {
	var output ivs.ListPlaybackKeyPairsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IVSListStreamKeysFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IVSListStreamKeysFuture) Get(ctx workflow.Context) (*ivs.ListStreamKeysOutput, error) {
	var output ivs.ListStreamKeysOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IVSListStreamsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IVSListStreamsFuture) Get(ctx workflow.Context) (*ivs.ListStreamsOutput, error) {
	var output ivs.ListStreamsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IVSListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IVSListTagsForResourceFuture) Get(ctx workflow.Context) (*ivs.ListTagsForResourceOutput, error) {
	var output ivs.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IVSPutMetadataFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IVSPutMetadataFuture) Get(ctx workflow.Context) (*ivs.PutMetadataOutput, error) {
	var output ivs.PutMetadataOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IVSStopStreamFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IVSStopStreamFuture) Get(ctx workflow.Context) (*ivs.StopStreamOutput, error) {
	var output ivs.StopStreamOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IVSTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IVSTagResourceFuture) Get(ctx workflow.Context) (*ivs.TagResourceOutput, error) {
	var output ivs.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IVSUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IVSUntagResourceFuture) Get(ctx workflow.Context) (*ivs.UntagResourceOutput, error) {
	var output ivs.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IVSUpdateChannelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IVSUpdateChannelFuture) Get(ctx workflow.Context) (*ivs.UpdateChannelOutput, error) {
	var output ivs.UpdateChannelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchGetChannel(ctx workflow.Context, input *ivs.BatchGetChannelInput) (*ivs.BatchGetChannelOutput, error) {
	var output ivs.BatchGetChannelOutput
	err := workflow.ExecuteActivity(ctx, "aws.ivs.BatchGetChannel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchGetChannelAsync(ctx workflow.Context, input *ivs.BatchGetChannelInput) *IVSBatchGetChannelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ivs.BatchGetChannel", input)
	return &IVSBatchGetChannelFuture{Future: future}
}

func (a *stub) BatchGetStreamKey(ctx workflow.Context, input *ivs.BatchGetStreamKeyInput) (*ivs.BatchGetStreamKeyOutput, error) {
	var output ivs.BatchGetStreamKeyOutput
	err := workflow.ExecuteActivity(ctx, "aws.ivs.BatchGetStreamKey", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchGetStreamKeyAsync(ctx workflow.Context, input *ivs.BatchGetStreamKeyInput) *IVSBatchGetStreamKeyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ivs.BatchGetStreamKey", input)
	return &IVSBatchGetStreamKeyFuture{Future: future}
}

func (a *stub) CreateChannel(ctx workflow.Context, input *ivs.CreateChannelInput) (*ivs.CreateChannelOutput, error) {
	var output ivs.CreateChannelOutput
	err := workflow.ExecuteActivity(ctx, "aws.ivs.CreateChannel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateChannelAsync(ctx workflow.Context, input *ivs.CreateChannelInput) *IVSCreateChannelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ivs.CreateChannel", input)
	return &IVSCreateChannelFuture{Future: future}
}

func (a *stub) CreateStreamKey(ctx workflow.Context, input *ivs.CreateStreamKeyInput) (*ivs.CreateStreamKeyOutput, error) {
	var output ivs.CreateStreamKeyOutput
	err := workflow.ExecuteActivity(ctx, "aws.ivs.CreateStreamKey", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateStreamKeyAsync(ctx workflow.Context, input *ivs.CreateStreamKeyInput) *IVSCreateStreamKeyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ivs.CreateStreamKey", input)
	return &IVSCreateStreamKeyFuture{Future: future}
}

func (a *stub) DeleteChannel(ctx workflow.Context, input *ivs.DeleteChannelInput) (*ivs.DeleteChannelOutput, error) {
	var output ivs.DeleteChannelOutput
	err := workflow.ExecuteActivity(ctx, "aws.ivs.DeleteChannel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteChannelAsync(ctx workflow.Context, input *ivs.DeleteChannelInput) *IVSDeleteChannelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ivs.DeleteChannel", input)
	return &IVSDeleteChannelFuture{Future: future}
}

func (a *stub) DeletePlaybackKeyPair(ctx workflow.Context, input *ivs.DeletePlaybackKeyPairInput) (*ivs.DeletePlaybackKeyPairOutput, error) {
	var output ivs.DeletePlaybackKeyPairOutput
	err := workflow.ExecuteActivity(ctx, "aws.ivs.DeletePlaybackKeyPair", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeletePlaybackKeyPairAsync(ctx workflow.Context, input *ivs.DeletePlaybackKeyPairInput) *IVSDeletePlaybackKeyPairFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ivs.DeletePlaybackKeyPair", input)
	return &IVSDeletePlaybackKeyPairFuture{Future: future}
}

func (a *stub) DeleteStreamKey(ctx workflow.Context, input *ivs.DeleteStreamKeyInput) (*ivs.DeleteStreamKeyOutput, error) {
	var output ivs.DeleteStreamKeyOutput
	err := workflow.ExecuteActivity(ctx, "aws.ivs.DeleteStreamKey", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteStreamKeyAsync(ctx workflow.Context, input *ivs.DeleteStreamKeyInput) *IVSDeleteStreamKeyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ivs.DeleteStreamKey", input)
	return &IVSDeleteStreamKeyFuture{Future: future}
}

func (a *stub) GetChannel(ctx workflow.Context, input *ivs.GetChannelInput) (*ivs.GetChannelOutput, error) {
	var output ivs.GetChannelOutput
	err := workflow.ExecuteActivity(ctx, "aws.ivs.GetChannel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetChannelAsync(ctx workflow.Context, input *ivs.GetChannelInput) *IVSGetChannelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ivs.GetChannel", input)
	return &IVSGetChannelFuture{Future: future}
}

func (a *stub) GetPlaybackKeyPair(ctx workflow.Context, input *ivs.GetPlaybackKeyPairInput) (*ivs.GetPlaybackKeyPairOutput, error) {
	var output ivs.GetPlaybackKeyPairOutput
	err := workflow.ExecuteActivity(ctx, "aws.ivs.GetPlaybackKeyPair", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetPlaybackKeyPairAsync(ctx workflow.Context, input *ivs.GetPlaybackKeyPairInput) *IVSGetPlaybackKeyPairFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ivs.GetPlaybackKeyPair", input)
	return &IVSGetPlaybackKeyPairFuture{Future: future}
}

func (a *stub) GetStream(ctx workflow.Context, input *ivs.GetStreamInput) (*ivs.GetStreamOutput, error) {
	var output ivs.GetStreamOutput
	err := workflow.ExecuteActivity(ctx, "aws.ivs.GetStream", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetStreamAsync(ctx workflow.Context, input *ivs.GetStreamInput) *IVSGetStreamFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ivs.GetStream", input)
	return &IVSGetStreamFuture{Future: future}
}

func (a *stub) GetStreamKey(ctx workflow.Context, input *ivs.GetStreamKeyInput) (*ivs.GetStreamKeyOutput, error) {
	var output ivs.GetStreamKeyOutput
	err := workflow.ExecuteActivity(ctx, "aws.ivs.GetStreamKey", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetStreamKeyAsync(ctx workflow.Context, input *ivs.GetStreamKeyInput) *IVSGetStreamKeyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ivs.GetStreamKey", input)
	return &IVSGetStreamKeyFuture{Future: future}
}

func (a *stub) ImportPlaybackKeyPair(ctx workflow.Context, input *ivs.ImportPlaybackKeyPairInput) (*ivs.ImportPlaybackKeyPairOutput, error) {
	var output ivs.ImportPlaybackKeyPairOutput
	err := workflow.ExecuteActivity(ctx, "aws.ivs.ImportPlaybackKeyPair", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ImportPlaybackKeyPairAsync(ctx workflow.Context, input *ivs.ImportPlaybackKeyPairInput) *IVSImportPlaybackKeyPairFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ivs.ImportPlaybackKeyPair", input)
	return &IVSImportPlaybackKeyPairFuture{Future: future}
}

func (a *stub) ListChannels(ctx workflow.Context, input *ivs.ListChannelsInput) (*ivs.ListChannelsOutput, error) {
	var output ivs.ListChannelsOutput
	err := workflow.ExecuteActivity(ctx, "aws.ivs.ListChannels", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListChannelsAsync(ctx workflow.Context, input *ivs.ListChannelsInput) *IVSListChannelsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ivs.ListChannels", input)
	return &IVSListChannelsFuture{Future: future}
}

func (a *stub) ListPlaybackKeyPairs(ctx workflow.Context, input *ivs.ListPlaybackKeyPairsInput) (*ivs.ListPlaybackKeyPairsOutput, error) {
	var output ivs.ListPlaybackKeyPairsOutput
	err := workflow.ExecuteActivity(ctx, "aws.ivs.ListPlaybackKeyPairs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPlaybackKeyPairsAsync(ctx workflow.Context, input *ivs.ListPlaybackKeyPairsInput) *IVSListPlaybackKeyPairsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ivs.ListPlaybackKeyPairs", input)
	return &IVSListPlaybackKeyPairsFuture{Future: future}
}

func (a *stub) ListStreamKeys(ctx workflow.Context, input *ivs.ListStreamKeysInput) (*ivs.ListStreamKeysOutput, error) {
	var output ivs.ListStreamKeysOutput
	err := workflow.ExecuteActivity(ctx, "aws.ivs.ListStreamKeys", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListStreamKeysAsync(ctx workflow.Context, input *ivs.ListStreamKeysInput) *IVSListStreamKeysFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ivs.ListStreamKeys", input)
	return &IVSListStreamKeysFuture{Future: future}
}

func (a *stub) ListStreams(ctx workflow.Context, input *ivs.ListStreamsInput) (*ivs.ListStreamsOutput, error) {
	var output ivs.ListStreamsOutput
	err := workflow.ExecuteActivity(ctx, "aws.ivs.ListStreams", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListStreamsAsync(ctx workflow.Context, input *ivs.ListStreamsInput) *IVSListStreamsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ivs.ListStreams", input)
	return &IVSListStreamsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *ivs.ListTagsForResourceInput) (*ivs.ListTagsForResourceOutput, error) {
	var output ivs.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.ivs.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *ivs.ListTagsForResourceInput) *IVSListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ivs.ListTagsForResource", input)
	return &IVSListTagsForResourceFuture{Future: future}
}

func (a *stub) PutMetadata(ctx workflow.Context, input *ivs.PutMetadataInput) (*ivs.PutMetadataOutput, error) {
	var output ivs.PutMetadataOutput
	err := workflow.ExecuteActivity(ctx, "aws.ivs.PutMetadata", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutMetadataAsync(ctx workflow.Context, input *ivs.PutMetadataInput) *IVSPutMetadataFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ivs.PutMetadata", input)
	return &IVSPutMetadataFuture{Future: future}
}

func (a *stub) StopStream(ctx workflow.Context, input *ivs.StopStreamInput) (*ivs.StopStreamOutput, error) {
	var output ivs.StopStreamOutput
	err := workflow.ExecuteActivity(ctx, "aws.ivs.StopStream", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopStreamAsync(ctx workflow.Context, input *ivs.StopStreamInput) *IVSStopStreamFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ivs.StopStream", input)
	return &IVSStopStreamFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *ivs.TagResourceInput) (*ivs.TagResourceOutput, error) {
	var output ivs.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.ivs.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *ivs.TagResourceInput) *IVSTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ivs.TagResource", input)
	return &IVSTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *ivs.UntagResourceInput) (*ivs.UntagResourceOutput, error) {
	var output ivs.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.ivs.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *ivs.UntagResourceInput) *IVSUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ivs.UntagResource", input)
	return &IVSUntagResourceFuture{Future: future}
}

func (a *stub) UpdateChannel(ctx workflow.Context, input *ivs.UpdateChannelInput) (*ivs.UpdateChannelOutput, error) {
	var output ivs.UpdateChannelOutput
	err := workflow.ExecuteActivity(ctx, "aws.ivs.UpdateChannel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateChannelAsync(ctx workflow.Context, input *ivs.UpdateChannelInput) *IVSUpdateChannelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ivs.UpdateChannel", input)
	return &IVSUpdateChannelFuture{Future: future}
}
