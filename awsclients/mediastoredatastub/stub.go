// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package mediastoredatastub

import (
	"github.com/aws/aws-sdk-go/service/mediastoredata"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type MediaStoreDataDeleteObjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStoreDataDeleteObjectFuture) Get(ctx workflow.Context) (*mediastoredata.DeleteObjectOutput, error) {
	var output mediastoredata.DeleteObjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStoreDataDescribeObjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStoreDataDescribeObjectFuture) Get(ctx workflow.Context) (*mediastoredata.DescribeObjectOutput, error) {
	var output mediastoredata.DescribeObjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStoreDataGetObjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStoreDataGetObjectFuture) Get(ctx workflow.Context) (*mediastoredata.GetObjectOutput, error) {
	var output mediastoredata.GetObjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStoreDataListItemsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStoreDataListItemsFuture) Get(ctx workflow.Context) (*mediastoredata.ListItemsOutput, error) {
	var output mediastoredata.ListItemsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaStoreDataPutObjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaStoreDataPutObjectFuture) Get(ctx workflow.Context) (*mediastoredata.PutObjectOutput, error) {
	var output mediastoredata.PutObjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteObject(ctx workflow.Context, input *mediastoredata.DeleteObjectInput) (*mediastoredata.DeleteObjectOutput, error) {
	var output mediastoredata.DeleteObjectOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastoredata.DeleteObject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteObjectAsync(ctx workflow.Context, input *mediastoredata.DeleteObjectInput) *MediaStoreDataDeleteObjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastoredata.DeleteObject", input)
	return &MediaStoreDataDeleteObjectFuture{Future: future}
}

func (a *stub) DescribeObject(ctx workflow.Context, input *mediastoredata.DescribeObjectInput) (*mediastoredata.DescribeObjectOutput, error) {
	var output mediastoredata.DescribeObjectOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastoredata.DescribeObject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeObjectAsync(ctx workflow.Context, input *mediastoredata.DescribeObjectInput) *MediaStoreDataDescribeObjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastoredata.DescribeObject", input)
	return &MediaStoreDataDescribeObjectFuture{Future: future}
}

func (a *stub) GetObject(ctx workflow.Context, input *mediastoredata.GetObjectInput) (*mediastoredata.GetObjectOutput, error) {
	var output mediastoredata.GetObjectOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastoredata.GetObject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetObjectAsync(ctx workflow.Context, input *mediastoredata.GetObjectInput) *MediaStoreDataGetObjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastoredata.GetObject", input)
	return &MediaStoreDataGetObjectFuture{Future: future}
}

func (a *stub) ListItems(ctx workflow.Context, input *mediastoredata.ListItemsInput) (*mediastoredata.ListItemsOutput, error) {
	var output mediastoredata.ListItemsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastoredata.ListItems", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListItemsAsync(ctx workflow.Context, input *mediastoredata.ListItemsInput) *MediaStoreDataListItemsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastoredata.ListItems", input)
	return &MediaStoreDataListItemsFuture{Future: future}
}

func (a *stub) PutObject(ctx workflow.Context, input *mediastoredata.PutObjectInput) (*mediastoredata.PutObjectOutput, error) {
	var output mediastoredata.PutObjectOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediastoredata.PutObject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutObjectAsync(ctx workflow.Context, input *mediastoredata.PutObjectInput) *MediaStoreDataPutObjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediastoredata.PutObject", input)
	return &MediaStoreDataPutObjectFuture{Future: future}
}
