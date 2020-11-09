// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package ssooidcstub

import (
	"github.com/aws/aws-sdk-go/service/ssooidc"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CreateTokenFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateTokenFuture) Get(ctx workflow.Context) (*ssooidc.CreateTokenOutput, error) {
	var output ssooidc.CreateTokenOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RegisterClientFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RegisterClientFuture) Get(ctx workflow.Context) (*ssooidc.RegisterClientOutput, error) {
	var output ssooidc.RegisterClientOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StartDeviceAuthorizationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StartDeviceAuthorizationFuture) Get(ctx workflow.Context) (*ssooidc.StartDeviceAuthorizationOutput, error) {
	var output ssooidc.StartDeviceAuthorizationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateToken(ctx workflow.Context, input *ssooidc.CreateTokenInput) (*ssooidc.CreateTokenOutput, error) {
	var output ssooidc.CreateTokenOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssooidc.CreateToken", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateTokenAsync(ctx workflow.Context, input *ssooidc.CreateTokenInput) *CreateTokenFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssooidc.CreateToken", input)
	return &CreateTokenFuture{Future: future}
}

func (a *stub) RegisterClient(ctx workflow.Context, input *ssooidc.RegisterClientInput) (*ssooidc.RegisterClientOutput, error) {
	var output ssooidc.RegisterClientOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssooidc.RegisterClient", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RegisterClientAsync(ctx workflow.Context, input *ssooidc.RegisterClientInput) *RegisterClientFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssooidc.RegisterClient", input)
	return &RegisterClientFuture{Future: future}
}

func (a *stub) StartDeviceAuthorization(ctx workflow.Context, input *ssooidc.StartDeviceAuthorizationInput) (*ssooidc.StartDeviceAuthorizationOutput, error) {
	var output ssooidc.StartDeviceAuthorizationOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssooidc.StartDeviceAuthorization", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartDeviceAuthorizationAsync(ctx workflow.Context, input *ssooidc.StartDeviceAuthorizationInput) *StartDeviceAuthorizationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssooidc.StartDeviceAuthorization", input)
	return &StartDeviceAuthorizationFuture{Future: future}
}
