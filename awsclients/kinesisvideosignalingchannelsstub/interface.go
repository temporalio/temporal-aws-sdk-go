// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package kinesisvideosignalingchannelsstub

import (
	"github.com/aws/aws-sdk-go/service/kinesisvideosignalingchannels"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

type Client interface {
	GetIceServerConfig(ctx workflow.Context, input *kinesisvideosignalingchannels.GetIceServerConfigInput) (*kinesisvideosignalingchannels.GetIceServerConfigOutput, error)
	GetIceServerConfigAsync(ctx workflow.Context, input *kinesisvideosignalingchannels.GetIceServerConfigInput) *KinesisVideoSignalingChannelsGetIceServerConfigFuture

	SendAlexaOfferToMaster(ctx workflow.Context, input *kinesisvideosignalingchannels.SendAlexaOfferToMasterInput) (*kinesisvideosignalingchannels.SendAlexaOfferToMasterOutput, error)
	SendAlexaOfferToMasterAsync(ctx workflow.Context, input *kinesisvideosignalingchannels.SendAlexaOfferToMasterInput) *KinesisVideoSignalingChannelsSendAlexaOfferToMasterFuture
}

func NewClient() Client {
	return &stub{}
}

