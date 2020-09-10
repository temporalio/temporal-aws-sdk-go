package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesisvideomedia"
	"github.com/aws/aws-sdk-go/service/kinesisvideomedia/kinesisvideomediaiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type KinesisVideoMediaActivities struct {
	client kinesisvideomediaiface.KinesisVideoMediaAPI
}

func NewKinesisVideoMediaActivities(session *session.Session, config ...*aws.Config) *KinesisVideoMediaActivities {
	client := kinesisvideomedia.New(session, config...)
	return &KinesisVideoMediaActivities{client: client}
}

func (a *KinesisVideoMediaActivities) GetMedia(ctx context.Context, input *kinesisvideomedia.GetMediaInput) (*kinesisvideomedia.GetMediaOutput, error) {
	return a.client.GetMediaWithContext(ctx, input)
}
