// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package qldbsessionstub

import (
	"github.com/aws/aws-sdk-go/service/qldbsession"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type QLDBSessionSendCommandFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *QLDBSessionSendCommandFuture) Get(ctx workflow.Context) (*qldbsession.SendCommandOutput, error) {
	var output qldbsession.SendCommandOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) SendCommand(ctx workflow.Context, input *qldbsession.SendCommandInput) (*qldbsession.SendCommandOutput, error) {
	var output qldbsession.SendCommandOutput
	err := workflow.ExecuteActivity(ctx, "aws.qldbsession.SendCommand", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SendCommandAsync(ctx workflow.Context, input *qldbsession.SendCommandInput) *QLDBSessionSendCommandFuture {
	future := workflow.ExecuteActivity(ctx, "aws.qldbsession.SendCommand", input)
	return &QLDBSessionSendCommandFuture{Future: future}
}
