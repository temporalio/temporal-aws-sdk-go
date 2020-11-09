// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package kinesisvideomediastub

import (
	"github.com/aws/aws-sdk-go/service/kinesisvideomedia"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type GetMediaFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetMediaFuture) Get(ctx workflow.Context) (*kinesisvideomedia.GetMediaOutput, error) {
	var output kinesisvideomedia.GetMediaOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) GetMedia(ctx workflow.Context, input *kinesisvideomedia.GetMediaInput) (*kinesisvideomedia.GetMediaOutput, error) {
	var output kinesisvideomedia.GetMediaOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisvideomedia.GetMedia", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetMediaAsync(ctx workflow.Context, input *kinesisvideomedia.GetMediaInput) *GetMediaFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisvideomedia.GetMedia", input)
	return &GetMediaFuture{Future: future}
}
