package awsclients

import (
	"github.com/aws/aws-sdk-go/service/kinesisvideomedia"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type KinesisVideoMediaClient interface {
       GetMedia(ctx workflow.Context, input *kinesisvideomedia.GetMediaInput) (*kinesisvideomedia.GetMediaOutput, error)
       GetMediaAsync(ctx workflow.Context, input *kinesisvideomedia.GetMediaInput) *KinesisvideomediaGetMediaResult
}

type KinesisvideomediaGetMediaResult struct {
	Result workflow.Future
}

func (r *KinesisvideomediaGetMediaResult) Get(ctx workflow.Context) (*kinesisvideomedia.GetMediaOutput, error) {
    var output kinesisvideomedia.GetMediaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisVideoMediaStub struct {
    activities awsactivities.KinesisVideoMediaActivities
}

func NewKinesisVideoMediaStub() KinesisVideoMediaClient {
    return &KinesisVideoMediaStub{}
}

func (a *KinesisVideoMediaStub) GetMedia(ctx workflow.Context, input *kinesisvideomedia.GetMediaInput) (*kinesisvideomedia.GetMediaOutput, error) {
    var output kinesisvideomedia.GetMediaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMedia, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisVideoMediaStub) GetMediaAsync(ctx workflow.Context, input *kinesisvideomedia.GetMediaInput) *KinesisvideomediaGetMediaResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetMedia, input)
    return &KinesisvideomediaGetMediaResult{Result: future}
}
