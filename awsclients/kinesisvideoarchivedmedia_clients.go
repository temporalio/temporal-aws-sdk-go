package awsclients

import (
	"github.com/aws/aws-sdk-go/service/kinesisvideoarchivedmedia"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type KinesisVideoArchivedMediaClient interface {
    GetClip(ctx workflow.Context, input *kinesisvideoarchivedmedia.GetClipInput) (*kinesisvideoarchivedmedia.GetClipOutput, error)
    GetClipAsync(ctx workflow.Context, input *kinesisvideoarchivedmedia.GetClipInput) *KinesisvideoarchivedmediaGetClipResult

    GetDASHStreamingSessionURL(ctx workflow.Context, input *kinesisvideoarchivedmedia.GetDASHStreamingSessionURLInput) (*kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput, error)
    GetDASHStreamingSessionURLAsync(ctx workflow.Context, input *kinesisvideoarchivedmedia.GetDASHStreamingSessionURLInput) *KinesisvideoarchivedmediaGetDASHStreamingSessionURLResult

    GetHLSStreamingSessionURL(ctx workflow.Context, input *kinesisvideoarchivedmedia.GetHLSStreamingSessionURLInput) (*kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput, error)
    GetHLSStreamingSessionURLAsync(ctx workflow.Context, input *kinesisvideoarchivedmedia.GetHLSStreamingSessionURLInput) *KinesisvideoarchivedmediaGetHLSStreamingSessionURLResult

    GetMediaForFragmentList(ctx workflow.Context, input *kinesisvideoarchivedmedia.GetMediaForFragmentListInput) (*kinesisvideoarchivedmedia.GetMediaForFragmentListOutput, error)
    GetMediaForFragmentListAsync(ctx workflow.Context, input *kinesisvideoarchivedmedia.GetMediaForFragmentListInput) *KinesisvideoarchivedmediaGetMediaForFragmentListResult

    ListFragments(ctx workflow.Context, input *kinesisvideoarchivedmedia.ListFragmentsInput) (*kinesisvideoarchivedmedia.ListFragmentsOutput, error)
    ListFragmentsAsync(ctx workflow.Context, input *kinesisvideoarchivedmedia.ListFragmentsInput) *KinesisvideoarchivedmediaListFragmentsResult
}
type KinesisvideoarchivedmediaGetClipResult struct {
	Result workflow.Future
}

func (r *KinesisvideoarchivedmediaGetClipResult) Get(ctx workflow.Context) (*kinesisvideoarchivedmedia.GetClipOutput, error) {
    var output kinesisvideoarchivedmedia.GetClipOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisvideoarchivedmediaGetDASHStreamingSessionURLResult struct {
	Result workflow.Future
}

func (r *KinesisvideoarchivedmediaGetDASHStreamingSessionURLResult) Get(ctx workflow.Context) (*kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput, error) {
    var output kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisvideoarchivedmediaGetHLSStreamingSessionURLResult struct {
	Result workflow.Future
}

func (r *KinesisvideoarchivedmediaGetHLSStreamingSessionURLResult) Get(ctx workflow.Context) (*kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput, error) {
    var output kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisvideoarchivedmediaGetMediaForFragmentListResult struct {
	Result workflow.Future
}

func (r *KinesisvideoarchivedmediaGetMediaForFragmentListResult) Get(ctx workflow.Context) (*kinesisvideoarchivedmedia.GetMediaForFragmentListOutput, error) {
    var output kinesisvideoarchivedmedia.GetMediaForFragmentListOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisvideoarchivedmediaListFragmentsResult struct {
	Result workflow.Future
}

func (r *KinesisvideoarchivedmediaListFragmentsResult) Get(ctx workflow.Context) (*kinesisvideoarchivedmedia.ListFragmentsOutput, error) {
    var output kinesisvideoarchivedmedia.ListFragmentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type KinesisVideoArchivedMediaStub struct {
    activities awsactivities.KinesisVideoArchivedMediaActivities
}

func NewKinesisVideoArchivedMediaStub() KinesisVideoArchivedMediaClient {
    return &KinesisVideoArchivedMediaStub{}
}
func (a *KinesisVideoArchivedMediaStub) GetClip(ctx workflow.Context, input *kinesisvideoarchivedmedia.GetClipInput) (*kinesisvideoarchivedmedia.GetClipOutput, error) {
    var output kinesisvideoarchivedmedia.GetClipOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetClip, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisVideoArchivedMediaStub) GetClipAsync(ctx workflow.Context, input *kinesisvideoarchivedmedia.GetClipInput) *KinesisvideoarchivedmediaGetClipResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetClip, input)
    return &KinesisvideoarchivedmediaGetClipResult{Result: future}
}
func (a *KinesisVideoArchivedMediaStub) GetDASHStreamingSessionURL(ctx workflow.Context, input *kinesisvideoarchivedmedia.GetDASHStreamingSessionURLInput) (*kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput, error) {
    var output kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDASHStreamingSessionURL, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisVideoArchivedMediaStub) GetDASHStreamingSessionURLAsync(ctx workflow.Context, input *kinesisvideoarchivedmedia.GetDASHStreamingSessionURLInput) *KinesisvideoarchivedmediaGetDASHStreamingSessionURLResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDASHStreamingSessionURL, input)
    return &KinesisvideoarchivedmediaGetDASHStreamingSessionURLResult{Result: future}
}
func (a *KinesisVideoArchivedMediaStub) GetHLSStreamingSessionURL(ctx workflow.Context, input *kinesisvideoarchivedmedia.GetHLSStreamingSessionURLInput) (*kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput, error) {
    var output kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetHLSStreamingSessionURL, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisVideoArchivedMediaStub) GetHLSStreamingSessionURLAsync(ctx workflow.Context, input *kinesisvideoarchivedmedia.GetHLSStreamingSessionURLInput) *KinesisvideoarchivedmediaGetHLSStreamingSessionURLResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetHLSStreamingSessionURL, input)
    return &KinesisvideoarchivedmediaGetHLSStreamingSessionURLResult{Result: future}
}
func (a *KinesisVideoArchivedMediaStub) GetMediaForFragmentList(ctx workflow.Context, input *kinesisvideoarchivedmedia.GetMediaForFragmentListInput) (*kinesisvideoarchivedmedia.GetMediaForFragmentListOutput, error) {
    var output kinesisvideoarchivedmedia.GetMediaForFragmentListOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMediaForFragmentList, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisVideoArchivedMediaStub) GetMediaForFragmentListAsync(ctx workflow.Context, input *kinesisvideoarchivedmedia.GetMediaForFragmentListInput) *KinesisvideoarchivedmediaGetMediaForFragmentListResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetMediaForFragmentList, input)
    return &KinesisvideoarchivedmediaGetMediaForFragmentListResult{Result: future}
}
func (a *KinesisVideoArchivedMediaStub) ListFragments(ctx workflow.Context, input *kinesisvideoarchivedmedia.ListFragmentsInput) (*kinesisvideoarchivedmedia.ListFragmentsOutput, error) {
    var output kinesisvideoarchivedmedia.ListFragmentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFragments, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisVideoArchivedMediaStub) ListFragmentsAsync(ctx workflow.Context, input *kinesisvideoarchivedmedia.ListFragmentsInput) *KinesisvideoarchivedmediaListFragmentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListFragments, input)
    return &KinesisvideoarchivedmediaListFragmentsResult{Result: future}
}
