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

type stub struct{}

type S3OutpostsCreateEndpointFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *S3OutpostsCreateEndpointFuture) Get(ctx workflow.Context) (*s3outposts.CreateEndpointOutput, error) {
	var output s3outposts.CreateEndpointOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type S3OutpostsDeleteEndpointFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *S3OutpostsDeleteEndpointFuture) Get(ctx workflow.Context) (*s3outposts.DeleteEndpointOutput, error) {
	var output s3outposts.DeleteEndpointOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type S3OutpostsListEndpointsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *S3OutpostsListEndpointsFuture) Get(ctx workflow.Context) (*s3outposts.ListEndpointsOutput, error) {
	var output s3outposts.ListEndpointsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateEndpoint(ctx workflow.Context, input *s3outposts.CreateEndpointInput) (*s3outposts.CreateEndpointOutput, error) {
	var output s3outposts.CreateEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws.s3outposts.CreateEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateEndpointAsync(ctx workflow.Context, input *s3outposts.CreateEndpointInput) *S3OutpostsCreateEndpointFuture {
	future := workflow.ExecuteActivity(ctx, "aws.s3outposts.CreateEndpoint", input)
	return &S3OutpostsCreateEndpointFuture{Future: future}
}

func (a *stub) DeleteEndpoint(ctx workflow.Context, input *s3outposts.DeleteEndpointInput) (*s3outposts.DeleteEndpointOutput, error) {
	var output s3outposts.DeleteEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws.s3outposts.DeleteEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteEndpointAsync(ctx workflow.Context, input *s3outposts.DeleteEndpointInput) *S3OutpostsDeleteEndpointFuture {
	future := workflow.ExecuteActivity(ctx, "aws.s3outposts.DeleteEndpoint", input)
	return &S3OutpostsDeleteEndpointFuture{Future: future}
}

func (a *stub) ListEndpoints(ctx workflow.Context, input *s3outposts.ListEndpointsInput) (*s3outposts.ListEndpointsOutput, error) {
	var output s3outposts.ListEndpointsOutput
	err := workflow.ExecuteActivity(ctx, "aws.s3outposts.ListEndpoints", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListEndpointsAsync(ctx workflow.Context, input *s3outposts.ListEndpointsInput) *S3OutpostsListEndpointsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.s3outposts.ListEndpoints", input)
	return &S3OutpostsListEndpointsFuture{Future: future}
}
