package awsclients

import (
	"github.com/aws/aws-sdk-go/service/kinesisvideo"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type KinesisVideoClient interface {
	CreateSignalingChannel(ctx workflow.Context, input *kinesisvideo.CreateSignalingChannelInput) (*kinesisvideo.CreateSignalingChannelOutput, error)
	CreateSignalingChannelAsync(ctx workflow.Context, input *kinesisvideo.CreateSignalingChannelInput) *KinesisvideoCreateSignalingChannelResult

	CreateStream(ctx workflow.Context, input *kinesisvideo.CreateStreamInput) (*kinesisvideo.CreateStreamOutput, error)
	CreateStreamAsync(ctx workflow.Context, input *kinesisvideo.CreateStreamInput) *KinesisvideoCreateStreamResult

	DeleteSignalingChannel(ctx workflow.Context, input *kinesisvideo.DeleteSignalingChannelInput) (*kinesisvideo.DeleteSignalingChannelOutput, error)
	DeleteSignalingChannelAsync(ctx workflow.Context, input *kinesisvideo.DeleteSignalingChannelInput) *KinesisvideoDeleteSignalingChannelResult

	DeleteStream(ctx workflow.Context, input *kinesisvideo.DeleteStreamInput) (*kinesisvideo.DeleteStreamOutput, error)
	DeleteStreamAsync(ctx workflow.Context, input *kinesisvideo.DeleteStreamInput) *KinesisvideoDeleteStreamResult

	DescribeSignalingChannel(ctx workflow.Context, input *kinesisvideo.DescribeSignalingChannelInput) (*kinesisvideo.DescribeSignalingChannelOutput, error)
	DescribeSignalingChannelAsync(ctx workflow.Context, input *kinesisvideo.DescribeSignalingChannelInput) *KinesisvideoDescribeSignalingChannelResult

	DescribeStream(ctx workflow.Context, input *kinesisvideo.DescribeStreamInput) (*kinesisvideo.DescribeStreamOutput, error)
	DescribeStreamAsync(ctx workflow.Context, input *kinesisvideo.DescribeStreamInput) *KinesisvideoDescribeStreamResult

	GetDataEndpoint(ctx workflow.Context, input *kinesisvideo.GetDataEndpointInput) (*kinesisvideo.GetDataEndpointOutput, error)
	GetDataEndpointAsync(ctx workflow.Context, input *kinesisvideo.GetDataEndpointInput) *KinesisvideoGetDataEndpointResult

	GetSignalingChannelEndpoint(ctx workflow.Context, input *kinesisvideo.GetSignalingChannelEndpointInput) (*kinesisvideo.GetSignalingChannelEndpointOutput, error)
	GetSignalingChannelEndpointAsync(ctx workflow.Context, input *kinesisvideo.GetSignalingChannelEndpointInput) *KinesisvideoGetSignalingChannelEndpointResult

	ListSignalingChannels(ctx workflow.Context, input *kinesisvideo.ListSignalingChannelsInput) (*kinesisvideo.ListSignalingChannelsOutput, error)
	ListSignalingChannelsAsync(ctx workflow.Context, input *kinesisvideo.ListSignalingChannelsInput) *KinesisvideoListSignalingChannelsResult

	ListStreams(ctx workflow.Context, input *kinesisvideo.ListStreamsInput) (*kinesisvideo.ListStreamsOutput, error)
	ListStreamsAsync(ctx workflow.Context, input *kinesisvideo.ListStreamsInput) *KinesisvideoListStreamsResult

	ListTagsForResource(ctx workflow.Context, input *kinesisvideo.ListTagsForResourceInput) (*kinesisvideo.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *kinesisvideo.ListTagsForResourceInput) *KinesisvideoListTagsForResourceResult

	ListTagsForStream(ctx workflow.Context, input *kinesisvideo.ListTagsForStreamInput) (*kinesisvideo.ListTagsForStreamOutput, error)
	ListTagsForStreamAsync(ctx workflow.Context, input *kinesisvideo.ListTagsForStreamInput) *KinesisvideoListTagsForStreamResult

	TagResource(ctx workflow.Context, input *kinesisvideo.TagResourceInput) (*kinesisvideo.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *kinesisvideo.TagResourceInput) *KinesisvideoTagResourceResult

	TagStream(ctx workflow.Context, input *kinesisvideo.TagStreamInput) (*kinesisvideo.TagStreamOutput, error)
	TagStreamAsync(ctx workflow.Context, input *kinesisvideo.TagStreamInput) *KinesisvideoTagStreamResult

	UntagResource(ctx workflow.Context, input *kinesisvideo.UntagResourceInput) (*kinesisvideo.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *kinesisvideo.UntagResourceInput) *KinesisvideoUntagResourceResult

	UntagStream(ctx workflow.Context, input *kinesisvideo.UntagStreamInput) (*kinesisvideo.UntagStreamOutput, error)
	UntagStreamAsync(ctx workflow.Context, input *kinesisvideo.UntagStreamInput) *KinesisvideoUntagStreamResult

	UpdateDataRetention(ctx workflow.Context, input *kinesisvideo.UpdateDataRetentionInput) (*kinesisvideo.UpdateDataRetentionOutput, error)
	UpdateDataRetentionAsync(ctx workflow.Context, input *kinesisvideo.UpdateDataRetentionInput) *KinesisvideoUpdateDataRetentionResult

	UpdateSignalingChannel(ctx workflow.Context, input *kinesisvideo.UpdateSignalingChannelInput) (*kinesisvideo.UpdateSignalingChannelOutput, error)
	UpdateSignalingChannelAsync(ctx workflow.Context, input *kinesisvideo.UpdateSignalingChannelInput) *KinesisvideoUpdateSignalingChannelResult

	UpdateStream(ctx workflow.Context, input *kinesisvideo.UpdateStreamInput) (*kinesisvideo.UpdateStreamOutput, error)
	UpdateStreamAsync(ctx workflow.Context, input *kinesisvideo.UpdateStreamInput) *KinesisvideoUpdateStreamResult
}

type KinesisvideoCreateSignalingChannelResult struct {
	Result workflow.Future
}

func (r *KinesisvideoCreateSignalingChannelResult) Get(ctx workflow.Context) (*kinesisvideo.CreateSignalingChannelOutput, error) {
	var output kinesisvideo.CreateSignalingChannelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisvideoCreateStreamResult struct {
	Result workflow.Future
}

func (r *KinesisvideoCreateStreamResult) Get(ctx workflow.Context) (*kinesisvideo.CreateStreamOutput, error) {
	var output kinesisvideo.CreateStreamOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisvideoDeleteSignalingChannelResult struct {
	Result workflow.Future
}

func (r *KinesisvideoDeleteSignalingChannelResult) Get(ctx workflow.Context) (*kinesisvideo.DeleteSignalingChannelOutput, error) {
	var output kinesisvideo.DeleteSignalingChannelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisvideoDeleteStreamResult struct {
	Result workflow.Future
}

func (r *KinesisvideoDeleteStreamResult) Get(ctx workflow.Context) (*kinesisvideo.DeleteStreamOutput, error) {
	var output kinesisvideo.DeleteStreamOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisvideoDescribeSignalingChannelResult struct {
	Result workflow.Future
}

func (r *KinesisvideoDescribeSignalingChannelResult) Get(ctx workflow.Context) (*kinesisvideo.DescribeSignalingChannelOutput, error) {
	var output kinesisvideo.DescribeSignalingChannelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisvideoDescribeStreamResult struct {
	Result workflow.Future
}

func (r *KinesisvideoDescribeStreamResult) Get(ctx workflow.Context) (*kinesisvideo.DescribeStreamOutput, error) {
	var output kinesisvideo.DescribeStreamOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisvideoGetDataEndpointResult struct {
	Result workflow.Future
}

func (r *KinesisvideoGetDataEndpointResult) Get(ctx workflow.Context) (*kinesisvideo.GetDataEndpointOutput, error) {
	var output kinesisvideo.GetDataEndpointOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisvideoGetSignalingChannelEndpointResult struct {
	Result workflow.Future
}

func (r *KinesisvideoGetSignalingChannelEndpointResult) Get(ctx workflow.Context) (*kinesisvideo.GetSignalingChannelEndpointOutput, error) {
	var output kinesisvideo.GetSignalingChannelEndpointOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisvideoListSignalingChannelsResult struct {
	Result workflow.Future
}

func (r *KinesisvideoListSignalingChannelsResult) Get(ctx workflow.Context) (*kinesisvideo.ListSignalingChannelsOutput, error) {
	var output kinesisvideo.ListSignalingChannelsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisvideoListStreamsResult struct {
	Result workflow.Future
}

func (r *KinesisvideoListStreamsResult) Get(ctx workflow.Context) (*kinesisvideo.ListStreamsOutput, error) {
	var output kinesisvideo.ListStreamsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisvideoListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *KinesisvideoListTagsForResourceResult) Get(ctx workflow.Context) (*kinesisvideo.ListTagsForResourceOutput, error) {
	var output kinesisvideo.ListTagsForResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisvideoListTagsForStreamResult struct {
	Result workflow.Future
}

func (r *KinesisvideoListTagsForStreamResult) Get(ctx workflow.Context) (*kinesisvideo.ListTagsForStreamOutput, error) {
	var output kinesisvideo.ListTagsForStreamOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisvideoTagResourceResult struct {
	Result workflow.Future
}

func (r *KinesisvideoTagResourceResult) Get(ctx workflow.Context) (*kinesisvideo.TagResourceOutput, error) {
	var output kinesisvideo.TagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisvideoTagStreamResult struct {
	Result workflow.Future
}

func (r *KinesisvideoTagStreamResult) Get(ctx workflow.Context) (*kinesisvideo.TagStreamOutput, error) {
	var output kinesisvideo.TagStreamOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisvideoUntagResourceResult struct {
	Result workflow.Future
}

func (r *KinesisvideoUntagResourceResult) Get(ctx workflow.Context) (*kinesisvideo.UntagResourceOutput, error) {
	var output kinesisvideo.UntagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisvideoUntagStreamResult struct {
	Result workflow.Future
}

func (r *KinesisvideoUntagStreamResult) Get(ctx workflow.Context) (*kinesisvideo.UntagStreamOutput, error) {
	var output kinesisvideo.UntagStreamOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisvideoUpdateDataRetentionResult struct {
	Result workflow.Future
}

func (r *KinesisvideoUpdateDataRetentionResult) Get(ctx workflow.Context) (*kinesisvideo.UpdateDataRetentionOutput, error) {
	var output kinesisvideo.UpdateDataRetentionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisvideoUpdateSignalingChannelResult struct {
	Result workflow.Future
}

func (r *KinesisvideoUpdateSignalingChannelResult) Get(ctx workflow.Context) (*kinesisvideo.UpdateSignalingChannelOutput, error) {
	var output kinesisvideo.UpdateSignalingChannelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisvideoUpdateStreamResult struct {
	Result workflow.Future
}

func (r *KinesisvideoUpdateStreamResult) Get(ctx workflow.Context) (*kinesisvideo.UpdateStreamOutput, error) {
	var output kinesisvideo.UpdateStreamOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type KinesisVideoStub struct {
	activities awsactivities.KinesisVideoActivities
}

func NewKinesisVideoStub() KinesisVideoClient {
	return &KinesisVideoStub{}
}

func (a *KinesisVideoStub) CreateSignalingChannel(ctx workflow.Context, input *kinesisvideo.CreateSignalingChannelInput) (*kinesisvideo.CreateSignalingChannelOutput, error) {
	var output kinesisvideo.CreateSignalingChannelOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateSignalingChannel, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisVideoStub) CreateSignalingChannelAsync(ctx workflow.Context, input *kinesisvideo.CreateSignalingChannelInput) *KinesisvideoCreateSignalingChannelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateSignalingChannel, input)
	return &KinesisvideoCreateSignalingChannelResult{Result: future}
}

func (a *KinesisVideoStub) CreateStream(ctx workflow.Context, input *kinesisvideo.CreateStreamInput) (*kinesisvideo.CreateStreamOutput, error) {
	var output kinesisvideo.CreateStreamOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateStream, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisVideoStub) CreateStreamAsync(ctx workflow.Context, input *kinesisvideo.CreateStreamInput) *KinesisvideoCreateStreamResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateStream, input)
	return &KinesisvideoCreateStreamResult{Result: future}
}

func (a *KinesisVideoStub) DeleteSignalingChannel(ctx workflow.Context, input *kinesisvideo.DeleteSignalingChannelInput) (*kinesisvideo.DeleteSignalingChannelOutput, error) {
	var output kinesisvideo.DeleteSignalingChannelOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteSignalingChannel, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisVideoStub) DeleteSignalingChannelAsync(ctx workflow.Context, input *kinesisvideo.DeleteSignalingChannelInput) *KinesisvideoDeleteSignalingChannelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteSignalingChannel, input)
	return &KinesisvideoDeleteSignalingChannelResult{Result: future}
}

func (a *KinesisVideoStub) DeleteStream(ctx workflow.Context, input *kinesisvideo.DeleteStreamInput) (*kinesisvideo.DeleteStreamOutput, error) {
	var output kinesisvideo.DeleteStreamOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteStream, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisVideoStub) DeleteStreamAsync(ctx workflow.Context, input *kinesisvideo.DeleteStreamInput) *KinesisvideoDeleteStreamResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteStream, input)
	return &KinesisvideoDeleteStreamResult{Result: future}
}

func (a *KinesisVideoStub) DescribeSignalingChannel(ctx workflow.Context, input *kinesisvideo.DescribeSignalingChannelInput) (*kinesisvideo.DescribeSignalingChannelOutput, error) {
	var output kinesisvideo.DescribeSignalingChannelOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeSignalingChannel, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisVideoStub) DescribeSignalingChannelAsync(ctx workflow.Context, input *kinesisvideo.DescribeSignalingChannelInput) *KinesisvideoDescribeSignalingChannelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeSignalingChannel, input)
	return &KinesisvideoDescribeSignalingChannelResult{Result: future}
}

func (a *KinesisVideoStub) DescribeStream(ctx workflow.Context, input *kinesisvideo.DescribeStreamInput) (*kinesisvideo.DescribeStreamOutput, error) {
	var output kinesisvideo.DescribeStreamOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeStream, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisVideoStub) DescribeStreamAsync(ctx workflow.Context, input *kinesisvideo.DescribeStreamInput) *KinesisvideoDescribeStreamResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeStream, input)
	return &KinesisvideoDescribeStreamResult{Result: future}
}

func (a *KinesisVideoStub) GetDataEndpoint(ctx workflow.Context, input *kinesisvideo.GetDataEndpointInput) (*kinesisvideo.GetDataEndpointOutput, error) {
	var output kinesisvideo.GetDataEndpointOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDataEndpoint, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisVideoStub) GetDataEndpointAsync(ctx workflow.Context, input *kinesisvideo.GetDataEndpointInput) *KinesisvideoGetDataEndpointResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDataEndpoint, input)
	return &KinesisvideoGetDataEndpointResult{Result: future}
}

func (a *KinesisVideoStub) GetSignalingChannelEndpoint(ctx workflow.Context, input *kinesisvideo.GetSignalingChannelEndpointInput) (*kinesisvideo.GetSignalingChannelEndpointOutput, error) {
	var output kinesisvideo.GetSignalingChannelEndpointOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetSignalingChannelEndpoint, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisVideoStub) GetSignalingChannelEndpointAsync(ctx workflow.Context, input *kinesisvideo.GetSignalingChannelEndpointInput) *KinesisvideoGetSignalingChannelEndpointResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetSignalingChannelEndpoint, input)
	return &KinesisvideoGetSignalingChannelEndpointResult{Result: future}
}

func (a *KinesisVideoStub) ListSignalingChannels(ctx workflow.Context, input *kinesisvideo.ListSignalingChannelsInput) (*kinesisvideo.ListSignalingChannelsOutput, error) {
	var output kinesisvideo.ListSignalingChannelsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListSignalingChannels, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisVideoStub) ListSignalingChannelsAsync(ctx workflow.Context, input *kinesisvideo.ListSignalingChannelsInput) *KinesisvideoListSignalingChannelsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListSignalingChannels, input)
	return &KinesisvideoListSignalingChannelsResult{Result: future}
}

func (a *KinesisVideoStub) ListStreams(ctx workflow.Context, input *kinesisvideo.ListStreamsInput) (*kinesisvideo.ListStreamsOutput, error) {
	var output kinesisvideo.ListStreamsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListStreams, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisVideoStub) ListStreamsAsync(ctx workflow.Context, input *kinesisvideo.ListStreamsInput) *KinesisvideoListStreamsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListStreams, input)
	return &KinesisvideoListStreamsResult{Result: future}
}

func (a *KinesisVideoStub) ListTagsForResource(ctx workflow.Context, input *kinesisvideo.ListTagsForResourceInput) (*kinesisvideo.ListTagsForResourceOutput, error) {
	var output kinesisvideo.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisVideoStub) ListTagsForResourceAsync(ctx workflow.Context, input *kinesisvideo.ListTagsForResourceInput) *KinesisvideoListTagsForResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
	return &KinesisvideoListTagsForResourceResult{Result: future}
}

func (a *KinesisVideoStub) ListTagsForStream(ctx workflow.Context, input *kinesisvideo.ListTagsForStreamInput) (*kinesisvideo.ListTagsForStreamOutput, error) {
	var output kinesisvideo.ListTagsForStreamOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForStream, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisVideoStub) ListTagsForStreamAsync(ctx workflow.Context, input *kinesisvideo.ListTagsForStreamInput) *KinesisvideoListTagsForStreamResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForStream, input)
	return &KinesisvideoListTagsForStreamResult{Result: future}
}

func (a *KinesisVideoStub) TagResource(ctx workflow.Context, input *kinesisvideo.TagResourceInput) (*kinesisvideo.TagResourceOutput, error) {
	var output kinesisvideo.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisVideoStub) TagResourceAsync(ctx workflow.Context, input *kinesisvideo.TagResourceInput) *KinesisvideoTagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
	return &KinesisvideoTagResourceResult{Result: future}
}

func (a *KinesisVideoStub) TagStream(ctx workflow.Context, input *kinesisvideo.TagStreamInput) (*kinesisvideo.TagStreamOutput, error) {
	var output kinesisvideo.TagStreamOutput
	err := workflow.ExecuteActivity(ctx, a.activities.TagStream, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisVideoStub) TagStreamAsync(ctx workflow.Context, input *kinesisvideo.TagStreamInput) *KinesisvideoTagStreamResult {
	future := workflow.ExecuteActivity(ctx, a.activities.TagStream, input)
	return &KinesisvideoTagStreamResult{Result: future}
}

func (a *KinesisVideoStub) UntagResource(ctx workflow.Context, input *kinesisvideo.UntagResourceInput) (*kinesisvideo.UntagResourceOutput, error) {
	var output kinesisvideo.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisVideoStub) UntagResourceAsync(ctx workflow.Context, input *kinesisvideo.UntagResourceInput) *KinesisvideoUntagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
	return &KinesisvideoUntagResourceResult{Result: future}
}

func (a *KinesisVideoStub) UntagStream(ctx workflow.Context, input *kinesisvideo.UntagStreamInput) (*kinesisvideo.UntagStreamOutput, error) {
	var output kinesisvideo.UntagStreamOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UntagStream, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisVideoStub) UntagStreamAsync(ctx workflow.Context, input *kinesisvideo.UntagStreamInput) *KinesisvideoUntagStreamResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UntagStream, input)
	return &KinesisvideoUntagStreamResult{Result: future}
}

func (a *KinesisVideoStub) UpdateDataRetention(ctx workflow.Context, input *kinesisvideo.UpdateDataRetentionInput) (*kinesisvideo.UpdateDataRetentionOutput, error) {
	var output kinesisvideo.UpdateDataRetentionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateDataRetention, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisVideoStub) UpdateDataRetentionAsync(ctx workflow.Context, input *kinesisvideo.UpdateDataRetentionInput) *KinesisvideoUpdateDataRetentionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateDataRetention, input)
	return &KinesisvideoUpdateDataRetentionResult{Result: future}
}

func (a *KinesisVideoStub) UpdateSignalingChannel(ctx workflow.Context, input *kinesisvideo.UpdateSignalingChannelInput) (*kinesisvideo.UpdateSignalingChannelOutput, error) {
	var output kinesisvideo.UpdateSignalingChannelOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateSignalingChannel, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisVideoStub) UpdateSignalingChannelAsync(ctx workflow.Context, input *kinesisvideo.UpdateSignalingChannelInput) *KinesisvideoUpdateSignalingChannelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateSignalingChannel, input)
	return &KinesisvideoUpdateSignalingChannelResult{Result: future}
}

func (a *KinesisVideoStub) UpdateStream(ctx workflow.Context, input *kinesisvideo.UpdateStreamInput) (*kinesisvideo.UpdateStreamOutput, error) {
	var output kinesisvideo.UpdateStreamOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateStream, input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisVideoStub) UpdateStreamAsync(ctx workflow.Context, input *kinesisvideo.UpdateStreamInput) *KinesisvideoUpdateStreamResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateStream, input)
	return &KinesisvideoUpdateStreamResult{Result: future}
}
