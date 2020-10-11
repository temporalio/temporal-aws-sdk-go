// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package timestreamquerystub

import (
	"github.com/aws/aws-sdk-go/service/timestreamquery"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

type Client interface {
	CancelQuery(ctx workflow.Context, input *timestreamquery.CancelQueryInput) (*timestreamquery.CancelQueryOutput, error)
	CancelQueryAsync(ctx workflow.Context, input *timestreamquery.CancelQueryInput) *TimestreamQueryCancelQueryFuture

	DescribeEndpoints(ctx workflow.Context, input *timestreamquery.DescribeEndpointsInput) (*timestreamquery.DescribeEndpointsOutput, error)
	DescribeEndpointsAsync(ctx workflow.Context, input *timestreamquery.DescribeEndpointsInput) *TimestreamQueryDescribeEndpointsFuture

	Query(ctx workflow.Context, input *timestreamquery.QueryInput) (*timestreamquery.QueryOutput, error)
	QueryAsync(ctx workflow.Context, input *timestreamquery.QueryInput) *TimestreamQueryQueryFuture
}

func NewClient() Client {
	return &stub{}
}

