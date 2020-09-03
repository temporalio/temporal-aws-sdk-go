package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/ivs"
	"go.temporal.io/sdk/workflow"
)

type IVSClient interface {
    BatchGetChannel(ctx workflow.Context, input *ivs.BatchGetChannelInput) (*ivs.BatchGetChannelOutput, error)
    BatchGetChannelAsync(ctx workflow.Context, input *ivs.BatchGetChannelInput) *IvsBatchGetChannelResult

    BatchGetStreamKey(ctx workflow.Context, input *ivs.BatchGetStreamKeyInput) (*ivs.BatchGetStreamKeyOutput, error)
    BatchGetStreamKeyAsync(ctx workflow.Context, input *ivs.BatchGetStreamKeyInput) *IvsBatchGetStreamKeyResult

    CreateChannel(ctx workflow.Context, input *ivs.CreateChannelInput) (*ivs.CreateChannelOutput, error)
    CreateChannelAsync(ctx workflow.Context, input *ivs.CreateChannelInput) *IvsCreateChannelResult

    CreateStreamKey(ctx workflow.Context, input *ivs.CreateStreamKeyInput) (*ivs.CreateStreamKeyOutput, error)
    CreateStreamKeyAsync(ctx workflow.Context, input *ivs.CreateStreamKeyInput) *IvsCreateStreamKeyResult

    DeleteChannel(ctx workflow.Context, input *ivs.DeleteChannelInput) (*ivs.DeleteChannelOutput, error)
    DeleteChannelAsync(ctx workflow.Context, input *ivs.DeleteChannelInput) *IvsDeleteChannelResult

    DeletePlaybackKeyPair(ctx workflow.Context, input *ivs.DeletePlaybackKeyPairInput) (*ivs.DeletePlaybackKeyPairOutput, error)
    DeletePlaybackKeyPairAsync(ctx workflow.Context, input *ivs.DeletePlaybackKeyPairInput) *IvsDeletePlaybackKeyPairResult

    DeleteStreamKey(ctx workflow.Context, input *ivs.DeleteStreamKeyInput) (*ivs.DeleteStreamKeyOutput, error)
    DeleteStreamKeyAsync(ctx workflow.Context, input *ivs.DeleteStreamKeyInput) *IvsDeleteStreamKeyResult

    GetChannel(ctx workflow.Context, input *ivs.GetChannelInput) (*ivs.GetChannelOutput, error)
    GetChannelAsync(ctx workflow.Context, input *ivs.GetChannelInput) *IvsGetChannelResult

    GetPlaybackKeyPair(ctx workflow.Context, input *ivs.GetPlaybackKeyPairInput) (*ivs.GetPlaybackKeyPairOutput, error)
    GetPlaybackKeyPairAsync(ctx workflow.Context, input *ivs.GetPlaybackKeyPairInput) *IvsGetPlaybackKeyPairResult

    GetStream(ctx workflow.Context, input *ivs.GetStreamInput) (*ivs.GetStreamOutput, error)
    GetStreamAsync(ctx workflow.Context, input *ivs.GetStreamInput) *IvsGetStreamResult

    GetStreamKey(ctx workflow.Context, input *ivs.GetStreamKeyInput) (*ivs.GetStreamKeyOutput, error)
    GetStreamKeyAsync(ctx workflow.Context, input *ivs.GetStreamKeyInput) *IvsGetStreamKeyResult

    ImportPlaybackKeyPair(ctx workflow.Context, input *ivs.ImportPlaybackKeyPairInput) (*ivs.ImportPlaybackKeyPairOutput, error)
    ImportPlaybackKeyPairAsync(ctx workflow.Context, input *ivs.ImportPlaybackKeyPairInput) *IvsImportPlaybackKeyPairResult

    ListChannels(ctx workflow.Context, input *ivs.ListChannelsInput) (*ivs.ListChannelsOutput, error)
    ListChannelsAsync(ctx workflow.Context, input *ivs.ListChannelsInput) *IvsListChannelsResult

    ListPlaybackKeyPairs(ctx workflow.Context, input *ivs.ListPlaybackKeyPairsInput) (*ivs.ListPlaybackKeyPairsOutput, error)
    ListPlaybackKeyPairsAsync(ctx workflow.Context, input *ivs.ListPlaybackKeyPairsInput) *IvsListPlaybackKeyPairsResult

    ListStreamKeys(ctx workflow.Context, input *ivs.ListStreamKeysInput) (*ivs.ListStreamKeysOutput, error)
    ListStreamKeysAsync(ctx workflow.Context, input *ivs.ListStreamKeysInput) *IvsListStreamKeysResult

    ListStreams(ctx workflow.Context, input *ivs.ListStreamsInput) (*ivs.ListStreamsOutput, error)
    ListStreamsAsync(ctx workflow.Context, input *ivs.ListStreamsInput) *IvsListStreamsResult

    ListTagsForResource(ctx workflow.Context, input *ivs.ListTagsForResourceInput) (*ivs.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *ivs.ListTagsForResourceInput) *IvsListTagsForResourceResult

    PutMetadata(ctx workflow.Context, input *ivs.PutMetadataInput) (*ivs.PutMetadataOutput, error)
    PutMetadataAsync(ctx workflow.Context, input *ivs.PutMetadataInput) *IvsPutMetadataResult

    StopStream(ctx workflow.Context, input *ivs.StopStreamInput) (*ivs.StopStreamOutput, error)
    StopStreamAsync(ctx workflow.Context, input *ivs.StopStreamInput) *IvsStopStreamResult

    TagResource(ctx workflow.Context, input *ivs.TagResourceInput) (*ivs.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *ivs.TagResourceInput) *IvsTagResourceResult

    UntagResource(ctx workflow.Context, input *ivs.UntagResourceInput) (*ivs.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *ivs.UntagResourceInput) *IvsUntagResourceResult

    UpdateChannel(ctx workflow.Context, input *ivs.UpdateChannelInput) (*ivs.UpdateChannelOutput, error)
    UpdateChannelAsync(ctx workflow.Context, input *ivs.UpdateChannelInput) *IvsUpdateChannelResult
}
type IvsBatchGetChannelResult struct {
	Result workflow.Future
}

func (r *IvsBatchGetChannelResult) Get(ctx workflow.Context) (*ivs.BatchGetChannelOutput, error) {
    var output ivs.BatchGetChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IvsBatchGetStreamKeyResult struct {
	Result workflow.Future
}

func (r *IvsBatchGetStreamKeyResult) Get(ctx workflow.Context) (*ivs.BatchGetStreamKeyOutput, error) {
    var output ivs.BatchGetStreamKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IvsCreateChannelResult struct {
	Result workflow.Future
}

func (r *IvsCreateChannelResult) Get(ctx workflow.Context) (*ivs.CreateChannelOutput, error) {
    var output ivs.CreateChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IvsCreateStreamKeyResult struct {
	Result workflow.Future
}

func (r *IvsCreateStreamKeyResult) Get(ctx workflow.Context) (*ivs.CreateStreamKeyOutput, error) {
    var output ivs.CreateStreamKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IvsDeleteChannelResult struct {
	Result workflow.Future
}

func (r *IvsDeleteChannelResult) Get(ctx workflow.Context) (*ivs.DeleteChannelOutput, error) {
    var output ivs.DeleteChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IvsDeletePlaybackKeyPairResult struct {
	Result workflow.Future
}

func (r *IvsDeletePlaybackKeyPairResult) Get(ctx workflow.Context) (*ivs.DeletePlaybackKeyPairOutput, error) {
    var output ivs.DeletePlaybackKeyPairOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IvsDeleteStreamKeyResult struct {
	Result workflow.Future
}

func (r *IvsDeleteStreamKeyResult) Get(ctx workflow.Context) (*ivs.DeleteStreamKeyOutput, error) {
    var output ivs.DeleteStreamKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IvsGetChannelResult struct {
	Result workflow.Future
}

func (r *IvsGetChannelResult) Get(ctx workflow.Context) (*ivs.GetChannelOutput, error) {
    var output ivs.GetChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IvsGetPlaybackKeyPairResult struct {
	Result workflow.Future
}

func (r *IvsGetPlaybackKeyPairResult) Get(ctx workflow.Context) (*ivs.GetPlaybackKeyPairOutput, error) {
    var output ivs.GetPlaybackKeyPairOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IvsGetStreamResult struct {
	Result workflow.Future
}

func (r *IvsGetStreamResult) Get(ctx workflow.Context) (*ivs.GetStreamOutput, error) {
    var output ivs.GetStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IvsGetStreamKeyResult struct {
	Result workflow.Future
}

func (r *IvsGetStreamKeyResult) Get(ctx workflow.Context) (*ivs.GetStreamKeyOutput, error) {
    var output ivs.GetStreamKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IvsImportPlaybackKeyPairResult struct {
	Result workflow.Future
}

func (r *IvsImportPlaybackKeyPairResult) Get(ctx workflow.Context) (*ivs.ImportPlaybackKeyPairOutput, error) {
    var output ivs.ImportPlaybackKeyPairOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IvsListChannelsResult struct {
	Result workflow.Future
}

func (r *IvsListChannelsResult) Get(ctx workflow.Context) (*ivs.ListChannelsOutput, error) {
    var output ivs.ListChannelsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IvsListPlaybackKeyPairsResult struct {
	Result workflow.Future
}

func (r *IvsListPlaybackKeyPairsResult) Get(ctx workflow.Context) (*ivs.ListPlaybackKeyPairsOutput, error) {
    var output ivs.ListPlaybackKeyPairsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IvsListStreamKeysResult struct {
	Result workflow.Future
}

func (r *IvsListStreamKeysResult) Get(ctx workflow.Context) (*ivs.ListStreamKeysOutput, error) {
    var output ivs.ListStreamKeysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IvsListStreamsResult struct {
	Result workflow.Future
}

func (r *IvsListStreamsResult) Get(ctx workflow.Context) (*ivs.ListStreamsOutput, error) {
    var output ivs.ListStreamsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IvsListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *IvsListTagsForResourceResult) Get(ctx workflow.Context) (*ivs.ListTagsForResourceOutput, error) {
    var output ivs.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IvsPutMetadataResult struct {
	Result workflow.Future
}

func (r *IvsPutMetadataResult) Get(ctx workflow.Context) (*ivs.PutMetadataOutput, error) {
    var output ivs.PutMetadataOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IvsStopStreamResult struct {
	Result workflow.Future
}

func (r *IvsStopStreamResult) Get(ctx workflow.Context) (*ivs.StopStreamOutput, error) {
    var output ivs.StopStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IvsTagResourceResult struct {
	Result workflow.Future
}

func (r *IvsTagResourceResult) Get(ctx workflow.Context) (*ivs.TagResourceOutput, error) {
    var output ivs.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IvsUntagResourceResult struct {
	Result workflow.Future
}

func (r *IvsUntagResourceResult) Get(ctx workflow.Context) (*ivs.UntagResourceOutput, error) {
    var output ivs.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IvsUpdateChannelResult struct {
	Result workflow.Future
}

func (r *IvsUpdateChannelResult) Get(ctx workflow.Context) (*ivs.UpdateChannelOutput, error) {
    var output ivs.UpdateChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type IVSStub struct {
    activities IVSClient
}

func NewIVSStub() IVSClient {
    return &IVSStub{}
}

func (a *IVSStub) BatchGetChannel(ctx workflow.Context, input *ivs.BatchGetChannelInput) (*ivs.BatchGetChannelOutput, error) {
    var output ivs.BatchGetChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *IVSStub) BatchGetStreamKey(ctx workflow.Context, input *ivs.BatchGetStreamKeyInput) (*ivs.BatchGetStreamKeyOutput, error) {
    var output ivs.BatchGetStreamKeyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetStreamKey, input).Get(ctx, &output)
    return &output, err
}

func (a *IVSStub) CreateChannel(ctx workflow.Context, input *ivs.CreateChannelInput) (*ivs.CreateChannelOutput, error) {
    var output ivs.CreateChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *IVSStub) CreateStreamKey(ctx workflow.Context, input *ivs.CreateStreamKeyInput) (*ivs.CreateStreamKeyOutput, error) {
    var output ivs.CreateStreamKeyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateStreamKey, input).Get(ctx, &output)
    return &output, err
}

func (a *IVSStub) DeleteChannel(ctx workflow.Context, input *ivs.DeleteChannelInput) (*ivs.DeleteChannelOutput, error) {
    var output ivs.DeleteChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *IVSStub) DeletePlaybackKeyPair(ctx workflow.Context, input *ivs.DeletePlaybackKeyPairInput) (*ivs.DeletePlaybackKeyPairOutput, error) {
    var output ivs.DeletePlaybackKeyPairOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePlaybackKeyPair, input).Get(ctx, &output)
    return &output, err
}

func (a *IVSStub) DeleteStreamKey(ctx workflow.Context, input *ivs.DeleteStreamKeyInput) (*ivs.DeleteStreamKeyOutput, error) {
    var output ivs.DeleteStreamKeyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteStreamKey, input).Get(ctx, &output)
    return &output, err
}

func (a *IVSStub) GetChannel(ctx workflow.Context, input *ivs.GetChannelInput) (*ivs.GetChannelOutput, error) {
    var output ivs.GetChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *IVSStub) GetPlaybackKeyPair(ctx workflow.Context, input *ivs.GetPlaybackKeyPairInput) (*ivs.GetPlaybackKeyPairOutput, error) {
    var output ivs.GetPlaybackKeyPairOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPlaybackKeyPair, input).Get(ctx, &output)
    return &output, err
}

func (a *IVSStub) GetStream(ctx workflow.Context, input *ivs.GetStreamInput) (*ivs.GetStreamOutput, error) {
    var output ivs.GetStreamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetStream, input).Get(ctx, &output)
    return &output, err
}

func (a *IVSStub) GetStreamKey(ctx workflow.Context, input *ivs.GetStreamKeyInput) (*ivs.GetStreamKeyOutput, error) {
    var output ivs.GetStreamKeyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetStreamKey, input).Get(ctx, &output)
    return &output, err
}

func (a *IVSStub) ImportPlaybackKeyPair(ctx workflow.Context, input *ivs.ImportPlaybackKeyPairInput) (*ivs.ImportPlaybackKeyPairOutput, error) {
    var output ivs.ImportPlaybackKeyPairOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ImportPlaybackKeyPair, input).Get(ctx, &output)
    return &output, err
}

func (a *IVSStub) ListChannels(ctx workflow.Context, input *ivs.ListChannelsInput) (*ivs.ListChannelsOutput, error) {
    var output ivs.ListChannelsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListChannels, input).Get(ctx, &output)
    return &output, err
}

func (a *IVSStub) ListPlaybackKeyPairs(ctx workflow.Context, input *ivs.ListPlaybackKeyPairsInput) (*ivs.ListPlaybackKeyPairsOutput, error) {
    var output ivs.ListPlaybackKeyPairsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPlaybackKeyPairs, input).Get(ctx, &output)
    return &output, err
}

func (a *IVSStub) ListStreamKeys(ctx workflow.Context, input *ivs.ListStreamKeysInput) (*ivs.ListStreamKeysOutput, error) {
    var output ivs.ListStreamKeysOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListStreamKeys, input).Get(ctx, &output)
    return &output, err
}

func (a *IVSStub) ListStreams(ctx workflow.Context, input *ivs.ListStreamsInput) (*ivs.ListStreamsOutput, error) {
    var output ivs.ListStreamsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListStreams, input).Get(ctx, &output)
    return &output, err
}

func (a *IVSStub) ListTagsForResource(ctx workflow.Context, input *ivs.ListTagsForResourceInput) (*ivs.ListTagsForResourceOutput, error) {
    var output ivs.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *IVSStub) PutMetadata(ctx workflow.Context, input *ivs.PutMetadataInput) (*ivs.PutMetadataOutput, error) {
    var output ivs.PutMetadataOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutMetadata, input).Get(ctx, &output)
    return &output, err
}

func (a *IVSStub) StopStream(ctx workflow.Context, input *ivs.StopStreamInput) (*ivs.StopStreamOutput, error) {
    var output ivs.StopStreamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopStream, input).Get(ctx, &output)
    return &output, err
}

func (a *IVSStub) TagResource(ctx workflow.Context, input *ivs.TagResourceInput) (*ivs.TagResourceOutput, error) {
    var output ivs.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *IVSStub) UntagResource(ctx workflow.Context, input *ivs.UntagResourceInput) (*ivs.UntagResourceOutput, error) {
    var output ivs.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *IVSStub) UpdateChannel(ctx workflow.Context, input *ivs.UpdateChannelInput) (*ivs.UpdateChannelOutput, error) {
    var output ivs.UpdateChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateChannel, input).Get(ctx, &output)
    return &output, err
}
