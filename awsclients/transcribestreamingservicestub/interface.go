// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package transcribestreamingservicestub

import (
	"github.com/aws/aws-sdk-go/service/transcribestreamingservice"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	StartStreamTranscription(ctx workflow.Context, input *transcribestreamingservice.StartStreamTranscriptionInput) (*transcribestreamingservice.StartStreamTranscriptionOutput, error)
	StartStreamTranscriptionAsync(ctx workflow.Context, input *transcribestreamingservice.StartStreamTranscriptionInput) *TranscribeStreamingServiceStartStreamTranscriptionFuture
}

func NewClient() Client {
	return &stub{}
}