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

type stub struct{}

type TranscribeStreamingServiceStartStreamTranscriptionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeStreamingServiceStartStreamTranscriptionFuture) Get(ctx workflow.Context) (*transcribestreamingservice.StartStreamTranscriptionOutput, error) {
	var output transcribestreamingservice.StartStreamTranscriptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) StartStreamTranscription(ctx workflow.Context, input *transcribestreamingservice.StartStreamTranscriptionInput) (*transcribestreamingservice.StartStreamTranscriptionOutput, error) {
	var output transcribestreamingservice.StartStreamTranscriptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribestreamingservice.StartStreamTranscription", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartStreamTranscriptionAsync(ctx workflow.Context, input *transcribestreamingservice.StartStreamTranscriptionInput) *TranscribeStreamingServiceStartStreamTranscriptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribestreamingservice.StartStreamTranscription", input)
	return &TranscribeStreamingServiceStartStreamTranscriptionFuture{Future: future}
}
