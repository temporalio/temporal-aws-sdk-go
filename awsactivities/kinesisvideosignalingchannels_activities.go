package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesisvideosignalingchannels"
	"github.com/aws/aws-sdk-go/service/kinesisvideosignalingchannels/kinesisvideosignalingchannelsiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type KinesisVideoSignalingChannelsActivities struct {
	client kinesisvideosignalingchannelsiface.KinesisVideoSignalingChannelsAPI
}

func NewKinesisVideoSignalingChannelsActivities(session *session.Session, config ...*aws.Config) *KinesisVideoSignalingChannelsActivities {
	client := kinesisvideosignalingchannels.New(session, config...)
	return &KinesisVideoSignalingChannelsActivities{client: client}
}

func (a *KinesisVideoSignalingChannelsActivities) GetIceServerConfig(ctx context.Context, input *kinesisvideosignalingchannels.GetIceServerConfigInput) (*kinesisvideosignalingchannels.GetIceServerConfigOutput, error) {
	return a.client.GetIceServerConfigWithContext(ctx, input)
}

func (a *KinesisVideoSignalingChannelsActivities) SendAlexaOfferToMaster(ctx context.Context, input *kinesisvideosignalingchannels.SendAlexaOfferToMasterInput) (*kinesisvideosignalingchannels.SendAlexaOfferToMasterOutput, error) {
	return a.client.SendAlexaOfferToMasterWithContext(ctx, input)
}
