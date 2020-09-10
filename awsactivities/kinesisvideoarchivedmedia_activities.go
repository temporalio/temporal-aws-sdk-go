package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesisvideoarchivedmedia"
	"github.com/aws/aws-sdk-go/service/kinesisvideoarchivedmedia/kinesisvideoarchivedmediaiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type KinesisVideoArchivedMediaActivities struct {
	client kinesisvideoarchivedmediaiface.KinesisVideoArchivedMediaAPI
}

func NewKinesisVideoArchivedMediaActivities(session *session.Session, config ...*aws.Config) *KinesisVideoArchivedMediaActivities {
	client := kinesisvideoarchivedmedia.New(session, config...)
	return &KinesisVideoArchivedMediaActivities{client: client}
}

func (a *KinesisVideoArchivedMediaActivities) GetClip(ctx context.Context, input *kinesisvideoarchivedmedia.GetClipInput) (*kinesisvideoarchivedmedia.GetClipOutput, error) {
	return a.client.GetClipWithContext(ctx, input)
}

func (a *KinesisVideoArchivedMediaActivities) GetDASHStreamingSessionURL(ctx context.Context, input *kinesisvideoarchivedmedia.GetDASHStreamingSessionURLInput) (*kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput, error) {
	return a.client.GetDASHStreamingSessionURLWithContext(ctx, input)
}

func (a *KinesisVideoArchivedMediaActivities) GetHLSStreamingSessionURL(ctx context.Context, input *kinesisvideoarchivedmedia.GetHLSStreamingSessionURLInput) (*kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput, error) {
	return a.client.GetHLSStreamingSessionURLWithContext(ctx, input)
}

func (a *KinesisVideoArchivedMediaActivities) GetMediaForFragmentList(ctx context.Context, input *kinesisvideoarchivedmedia.GetMediaForFragmentListInput) (*kinesisvideoarchivedmedia.GetMediaForFragmentListOutput, error) {
	return a.client.GetMediaForFragmentListWithContext(ctx, input)
}

func (a *KinesisVideoArchivedMediaActivities) ListFragments(ctx context.Context, input *kinesisvideoarchivedmedia.ListFragmentsInput) (*kinesisvideoarchivedmedia.ListFragmentsOutput, error) {
	return a.client.ListFragmentsWithContext(ctx, input)
}
