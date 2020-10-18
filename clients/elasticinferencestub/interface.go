// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package elasticinferencestub

import (
	"github.com/aws/aws-sdk-go/service/elasticinference"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	DescribeAcceleratorOfferings(ctx workflow.Context, input *elasticinference.DescribeAcceleratorOfferingsInput) (*elasticinference.DescribeAcceleratorOfferingsOutput, error)
	DescribeAcceleratorOfferingsAsync(ctx workflow.Context, input *elasticinference.DescribeAcceleratorOfferingsInput) *ElasticInferenceDescribeAcceleratorOfferingsFuture

	DescribeAcceleratorTypes(ctx workflow.Context, input *elasticinference.DescribeAcceleratorTypesInput) (*elasticinference.DescribeAcceleratorTypesOutput, error)
	DescribeAcceleratorTypesAsync(ctx workflow.Context, input *elasticinference.DescribeAcceleratorTypesInput) *ElasticInferenceDescribeAcceleratorTypesFuture

	DescribeAccelerators(ctx workflow.Context, input *elasticinference.DescribeAcceleratorsInput) (*elasticinference.DescribeAcceleratorsOutput, error)
	DescribeAcceleratorsAsync(ctx workflow.Context, input *elasticinference.DescribeAcceleratorsInput) *ElasticInferenceDescribeAcceleratorsFuture

	ListTagsForResource(ctx workflow.Context, input *elasticinference.ListTagsForResourceInput) (*elasticinference.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *elasticinference.ListTagsForResourceInput) *ElasticInferenceListTagsForResourceFuture

	TagResource(ctx workflow.Context, input *elasticinference.TagResourceInput) (*elasticinference.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *elasticinference.TagResourceInput) *ElasticInferenceTagResourceFuture

	UntagResource(ctx workflow.Context, input *elasticinference.UntagResourceInput) (*elasticinference.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *elasticinference.UntagResourceInput) *ElasticInferenceUntagResourceFuture
}

func NewClient() Client {
	return &stub{}
}
