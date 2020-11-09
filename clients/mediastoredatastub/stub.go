// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package mediastoredatastub

import (
	"github.com/aws/aws-sdk-go/service/mediastoredata"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type DeleteObjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteObjectFuture) Get(ctx workflow.Context) (*mediastoredata.DeleteObjectOutput, error) {
	var output mediastoredata.DeleteObjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeObjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeObjectFuture) Get(ctx workflow.Context) (*mediastoredata.DescribeObjectOutput, error) {
	var output mediastoredata.DescribeObjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetObjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetObjectFuture) Get(ctx workflow.Context) (*mediastoredata.GetObjectOutput, error) {
	var output mediastoredata.GetObjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListItemsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListItemsFuture) Get(ctx workflow.Context) (*mediastoredata.ListItemsOutput, error) {
	var output mediastoredata.ListItemsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutObjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutObjectFuture) Get(ctx workflow.Context) (*mediastoredata.PutObjectOutput, error) {
	var output mediastoredata.PutObjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteObject(ctx workflow.Context, input *mediastoredata.DeleteObjectInput) (*mediastoredata.DeleteObjectOutput, error) {
	var output mediastoredata.DeleteObjectOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastoredata.DeleteObject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteObjectAsync(ctx workflow.Context, input *mediastoredata.DeleteObjectInput) *DeleteObjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastoredata.DeleteObject", input)
	return &DeleteObjectFuture{Future: future}
}

func (a *stub) DescribeObject(ctx workflow.Context, input *mediastoredata.DescribeObjectInput) (*mediastoredata.DescribeObjectOutput, error) {
	var output mediastoredata.DescribeObjectOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastoredata.DescribeObject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeObjectAsync(ctx workflow.Context, input *mediastoredata.DescribeObjectInput) *DescribeObjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastoredata.DescribeObject", input)
	return &DescribeObjectFuture{Future: future}
}

func (a *stub) GetObject(ctx workflow.Context, input *mediastoredata.GetObjectInput) (*mediastoredata.GetObjectOutput, error) {
	var output mediastoredata.GetObjectOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastoredata.GetObject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetObjectAsync(ctx workflow.Context, input *mediastoredata.GetObjectInput) *GetObjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastoredata.GetObject", input)
	return &GetObjectFuture{Future: future}
}

func (a *stub) ListItems(ctx workflow.Context, input *mediastoredata.ListItemsInput) (*mediastoredata.ListItemsOutput, error) {
	var output mediastoredata.ListItemsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastoredata.ListItems", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListItemsAsync(ctx workflow.Context, input *mediastoredata.ListItemsInput) *ListItemsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastoredata.ListItems", input)
	return &ListItemsFuture{Future: future}
}

func (a *stub) PutObject(ctx workflow.Context, input *mediastoredata.PutObjectInput) (*mediastoredata.PutObjectOutput, error) {
	var output mediastoredata.PutObjectOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastoredata.PutObject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutObjectAsync(ctx workflow.Context, input *mediastoredata.PutObjectInput) *PutObjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastoredata.PutObject", input)
	return &PutObjectFuture{Future: future}
}
