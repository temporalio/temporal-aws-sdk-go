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

type Client interface {
	DeleteObject(ctx workflow.Context, input *mediastoredata.DeleteObjectInput) (*mediastoredata.DeleteObjectOutput, error)
	DeleteObjectAsync(ctx workflow.Context, input *mediastoredata.DeleteObjectInput) *MediaStoreDataDeleteObjectFuture

	DescribeObject(ctx workflow.Context, input *mediastoredata.DescribeObjectInput) (*mediastoredata.DescribeObjectOutput, error)
	DescribeObjectAsync(ctx workflow.Context, input *mediastoredata.DescribeObjectInput) *MediaStoreDataDescribeObjectFuture

	GetObject(ctx workflow.Context, input *mediastoredata.GetObjectInput) (*mediastoredata.GetObjectOutput, error)
	GetObjectAsync(ctx workflow.Context, input *mediastoredata.GetObjectInput) *MediaStoreDataGetObjectFuture

	ListItems(ctx workflow.Context, input *mediastoredata.ListItemsInput) (*mediastoredata.ListItemsOutput, error)
	ListItemsAsync(ctx workflow.Context, input *mediastoredata.ListItemsInput) *MediaStoreDataListItemsFuture

	PutObject(ctx workflow.Context, input *mediastoredata.PutObjectInput) (*mediastoredata.PutObjectOutput, error)
	PutObjectAsync(ctx workflow.Context, input *mediastoredata.PutObjectInput) *MediaStoreDataPutObjectFuture
}

func NewClient() Client {
	return &stub{}
}
