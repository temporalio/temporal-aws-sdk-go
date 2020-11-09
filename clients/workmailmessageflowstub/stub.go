// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package workmailmessageflowstub

import (
	"github.com/aws/aws-sdk-go/service/workmailmessageflow"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type GetRawMessageContentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetRawMessageContentFuture) Get(ctx workflow.Context) (*workmailmessageflow.GetRawMessageContentOutput, error) {
	var output workmailmessageflow.GetRawMessageContentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) GetRawMessageContent(ctx workflow.Context, input *workmailmessageflow.GetRawMessageContentInput) (*workmailmessageflow.GetRawMessageContentOutput, error) {
	var output workmailmessageflow.GetRawMessageContentOutput
	err := workflow.ExecuteActivity(ctx, "aws.workmailmessageflow.GetRawMessageContent", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetRawMessageContentAsync(ctx workflow.Context, input *workmailmessageflow.GetRawMessageContentInput) *GetRawMessageContentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workmailmessageflow.GetRawMessageContent", input)
	return &GetRawMessageContentFuture{Future: future}
}
