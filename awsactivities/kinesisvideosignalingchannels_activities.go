package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/kinesisvideosignalingchannels"
	"github.com/aws/aws-sdk-go/service/kinesisvideosignalingchannels/kinesisvideosignalingchannelsiface"
)

type KinesisVideoSignalingChannelsActivities struct {
	client kinesisvideosignalingchannelsiface.KinesisVideoSignalingChannelsAPI
}

func NewKinesisVideoSignalingChannelsActivities(client kinesisvideosignalingchannelsiface.KinesisVideoSignalingChannelsAPI) *KinesisVideoSignalingChannelsActivities {
    return &KinesisVideoSignalingChannelsActivities{client: client}
}


func (a *KinesisVideoSignalingChannelsActivities) GetIceServerConfig(input *kinesisvideosignalingchannels.GetIceServerConfigInput) (*kinesisvideosignalingchannels.GetIceServerConfigOutput, error) {
    return a.client.GetIceServerConfig(input)
}



func (a *KinesisVideoSignalingChannelsActivities) SendAlexaOfferToMaster(input *kinesisvideosignalingchannels.SendAlexaOfferToMasterInput) (*kinesisvideosignalingchannels.SendAlexaOfferToMasterOutput, error) {
    return a.client.SendAlexaOfferToMaster(input)
}

