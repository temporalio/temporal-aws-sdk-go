package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesisvideoarchivedmedia"
	"github.com/aws/aws-sdk-go/service/kinesisvideoarchivedmedia/kinesisvideoarchivedmediaiface"
)

type KinesisVideoArchivedMediaActivities struct {
	client kinesisvideoarchivedmediaiface.KinesisVideoArchivedMediaAPI
}

func NewKinesisVideoArchivedMediaActivities(session *session.Session, config ...*aws.Config) *KinesisVideoArchivedMediaActivities {
	client := kinesisvideoarchivedmedia.New(session, config...)
	return &KinesisVideoArchivedMediaActivities{client: client}
}

func (a *KinesisVideoArchivedMediaActivities) GetClip(input *kinesisvideoarchivedmedia.GetClipInput) (*kinesisvideoarchivedmedia.GetClipOutput, error) {
	return a.client.GetClip(input)
}

func (a *KinesisVideoArchivedMediaActivities) GetDASHStreamingSessionURL(input *kinesisvideoarchivedmedia.GetDASHStreamingSessionURLInput) (*kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput, error) {
	return a.client.GetDASHStreamingSessionURL(input)
}

func (a *KinesisVideoArchivedMediaActivities) GetHLSStreamingSessionURL(input *kinesisvideoarchivedmedia.GetHLSStreamingSessionURLInput) (*kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput, error) {
	return a.client.GetHLSStreamingSessionURL(input)
}

func (a *KinesisVideoArchivedMediaActivities) GetMediaForFragmentList(input *kinesisvideoarchivedmedia.GetMediaForFragmentListInput) (*kinesisvideoarchivedmedia.GetMediaForFragmentListOutput, error) {
	return a.client.GetMediaForFragmentList(input)
}

func (a *KinesisVideoArchivedMediaActivities) ListFragments(input *kinesisvideoarchivedmedia.ListFragmentsInput) (*kinesisvideoarchivedmedia.ListFragmentsOutput, error) {
	return a.client.ListFragments(input)
}
