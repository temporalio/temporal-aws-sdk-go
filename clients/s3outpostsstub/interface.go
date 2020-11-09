// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package s3outpostsstub

import (
	"github.com/aws/aws-sdk-go/service/s3outposts"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateEndpoint(ctx workflow.Context, input *s3outposts.CreateEndpointInput) (*s3outposts.CreateEndpointOutput, error)
	CreateEndpointAsync(ctx workflow.Context, input *s3outposts.CreateEndpointInput) *CreateEndpointFuture

	DeleteEndpoint(ctx workflow.Context, input *s3outposts.DeleteEndpointInput) (*s3outposts.DeleteEndpointOutput, error)
	DeleteEndpointAsync(ctx workflow.Context, input *s3outposts.DeleteEndpointInput) *DeleteEndpointFuture

	ListEndpoints(ctx workflow.Context, input *s3outposts.ListEndpointsInput) (*s3outposts.ListEndpointsOutput, error)
	ListEndpointsAsync(ctx workflow.Context, input *s3outposts.ListEndpointsInput) *ListEndpointsFuture
}

func NewClient() Client {
	return &stub{}
}
