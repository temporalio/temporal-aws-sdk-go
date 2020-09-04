package awsclients

import (
	"github.com/aws/aws-sdk-go/service/kinesisvideosignalingchannels"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type KinesisVideoSignalingChannelsClient interface {
    GetIceServerConfig(ctx workflow.Context, input *kinesisvideosignalingchannels.GetIceServerConfigInput) (*kinesisvideosignalingchannels.GetIceServerConfigOutput, error)
    GetIceServerConfigAsync(ctx workflow.Context, input *kinesisvideosignalingchannels.GetIceServerConfigInput) *KinesisvideosignalingchannelsGetIceServerConfigResult

    SendAlexaOfferToMaster(ctx workflow.Context, input *kinesisvideosignalingchannels.SendAlexaOfferToMasterInput) (*kinesisvideosignalingchannels.SendAlexaOfferToMasterOutput, error)
    SendAlexaOfferToMasterAsync(ctx workflow.Context, input *kinesisvideosignalingchannels.SendAlexaOfferToMasterInput) *KinesisvideosignalingchannelsSendAlexaOfferToMasterResult
}
type KinesisvideosignalingchannelsGetIceServerConfigResult struct {
	Result workflow.Future
}

func (r *KinesisvideosignalingchannelsGetIceServerConfigResult) Get(ctx workflow.Context) (*kinesisvideosignalingchannels.GetIceServerConfigOutput, error) {
    var output kinesisvideosignalingchannels.GetIceServerConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KinesisvideosignalingchannelsSendAlexaOfferToMasterResult struct {
	Result workflow.Future
}

func (r *KinesisvideosignalingchannelsSendAlexaOfferToMasterResult) Get(ctx workflow.Context) (*kinesisvideosignalingchannels.SendAlexaOfferToMasterOutput, error) {
    var output kinesisvideosignalingchannels.SendAlexaOfferToMasterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type KinesisVideoSignalingChannelsStub struct {
    activities awsactivities.KinesisVideoSignalingChannelsActivities
}

func NewKinesisVideoSignalingChannelsStub() KinesisVideoSignalingChannelsClient {
    return &KinesisVideoSignalingChannelsStub{}
}
func (a *KinesisVideoSignalingChannelsStub) GetIceServerConfig(ctx workflow.Context, input *kinesisvideosignalingchannels.GetIceServerConfigInput) (*kinesisvideosignalingchannels.GetIceServerConfigOutput, error) {
    var output kinesisvideosignalingchannels.GetIceServerConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetIceServerConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisVideoSignalingChannelsStub) GetIceServerConfigAsync(ctx workflow.Context, input *kinesisvideosignalingchannels.GetIceServerConfigInput) *KinesisvideosignalingchannelsGetIceServerConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetIceServerConfig, input)
    return &KinesisvideosignalingchannelsGetIceServerConfigResult{Result: future}
}
func (a *KinesisVideoSignalingChannelsStub) SendAlexaOfferToMaster(ctx workflow.Context, input *kinesisvideosignalingchannels.SendAlexaOfferToMasterInput) (*kinesisvideosignalingchannels.SendAlexaOfferToMasterOutput, error) {
    var output kinesisvideosignalingchannels.SendAlexaOfferToMasterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendAlexaOfferToMaster, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisVideoSignalingChannelsStub) SendAlexaOfferToMasterAsync(ctx workflow.Context, input *kinesisvideosignalingchannels.SendAlexaOfferToMasterInput) *KinesisvideosignalingchannelsSendAlexaOfferToMasterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SendAlexaOfferToMaster, input)
    return &KinesisvideosignalingchannelsSendAlexaOfferToMasterResult{Result: future}
}
