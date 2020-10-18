// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package apigatewaymanagementapistub

import (
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type ApiGatewayManagementApiDeleteConnectionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApiGatewayManagementApiDeleteConnectionFuture) Get(ctx workflow.Context) (*apigatewaymanagementapi.DeleteConnectionOutput, error) {
	var output apigatewaymanagementapi.DeleteConnectionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApiGatewayManagementApiGetConnectionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApiGatewayManagementApiGetConnectionFuture) Get(ctx workflow.Context) (*apigatewaymanagementapi.GetConnectionOutput, error) {
	var output apigatewaymanagementapi.GetConnectionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApiGatewayManagementApiPostToConnectionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApiGatewayManagementApiPostToConnectionFuture) Get(ctx workflow.Context) (*apigatewaymanagementapi.PostToConnectionOutput, error) {
	var output apigatewaymanagementapi.PostToConnectionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteConnection(ctx workflow.Context, input *apigatewaymanagementapi.DeleteConnectionInput) (*apigatewaymanagementapi.DeleteConnectionOutput, error) {
	var output apigatewaymanagementapi.DeleteConnectionOutput
	err := workflow.ExecuteActivity(ctx, "aws.apigatewaymanagementapi.DeleteConnection", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteConnectionAsync(ctx workflow.Context, input *apigatewaymanagementapi.DeleteConnectionInput) *ApiGatewayManagementApiDeleteConnectionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.apigatewaymanagementapi.DeleteConnection", input)
	return &ApiGatewayManagementApiDeleteConnectionFuture{Future: future}
}

func (a *stub) GetConnection(ctx workflow.Context, input *apigatewaymanagementapi.GetConnectionInput) (*apigatewaymanagementapi.GetConnectionOutput, error) {
	var output apigatewaymanagementapi.GetConnectionOutput
	err := workflow.ExecuteActivity(ctx, "aws.apigatewaymanagementapi.GetConnection", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetConnectionAsync(ctx workflow.Context, input *apigatewaymanagementapi.GetConnectionInput) *ApiGatewayManagementApiGetConnectionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.apigatewaymanagementapi.GetConnection", input)
	return &ApiGatewayManagementApiGetConnectionFuture{Future: future}
}

func (a *stub) PostToConnection(ctx workflow.Context, input *apigatewaymanagementapi.PostToConnectionInput) (*apigatewaymanagementapi.PostToConnectionOutput, error) {
	var output apigatewaymanagementapi.PostToConnectionOutput
	err := workflow.ExecuteActivity(ctx, "aws.apigatewaymanagementapi.PostToConnection", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PostToConnectionAsync(ctx workflow.Context, input *apigatewaymanagementapi.PostToConnectionInput) *ApiGatewayManagementApiPostToConnectionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.apigatewaymanagementapi.PostToConnection", input)
	return &ApiGatewayManagementApiPostToConnectionFuture{Future: future}
}
