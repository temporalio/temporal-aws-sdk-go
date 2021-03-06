// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package pricingstub

import (
	"github.com/aws/aws-sdk-go/service/pricing"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type DescribeServicesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeServicesFuture) Get(ctx workflow.Context) (*pricing.DescribeServicesOutput, error) {
	var output pricing.DescribeServicesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetAttributeValuesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetAttributeValuesFuture) Get(ctx workflow.Context) (*pricing.GetAttributeValuesOutput, error) {
	var output pricing.GetAttributeValuesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetProductsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetProductsFuture) Get(ctx workflow.Context) (*pricing.GetProductsOutput, error) {
	var output pricing.GetProductsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeServices(ctx workflow.Context, input *pricing.DescribeServicesInput) (*pricing.DescribeServicesOutput, error) {
	var output pricing.DescribeServicesOutput
	err := workflow.ExecuteActivity(ctx, "aws.pricing.DescribeServices", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeServicesAsync(ctx workflow.Context, input *pricing.DescribeServicesInput) *DescribeServicesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.pricing.DescribeServices", input)
	return &DescribeServicesFuture{Future: future}
}

func (a *stub) GetAttributeValues(ctx workflow.Context, input *pricing.GetAttributeValuesInput) (*pricing.GetAttributeValuesOutput, error) {
	var output pricing.GetAttributeValuesOutput
	err := workflow.ExecuteActivity(ctx, "aws.pricing.GetAttributeValues", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetAttributeValuesAsync(ctx workflow.Context, input *pricing.GetAttributeValuesInput) *GetAttributeValuesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.pricing.GetAttributeValues", input)
	return &GetAttributeValuesFuture{Future: future}
}

func (a *stub) GetProducts(ctx workflow.Context, input *pricing.GetProductsInput) (*pricing.GetProductsOutput, error) {
	var output pricing.GetProductsOutput
	err := workflow.ExecuteActivity(ctx, "aws.pricing.GetProducts", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetProductsAsync(ctx workflow.Context, input *pricing.GetProductsInput) *GetProductsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.pricing.GetProducts", input)
	return &GetProductsFuture{Future: future}
}
