// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package ioteventsstub

import (
	"github.com/aws/aws-sdk-go/service/iotevents"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateDetectorModel(ctx workflow.Context, input *iotevents.CreateDetectorModelInput) (*iotevents.CreateDetectorModelOutput, error)
	CreateDetectorModelAsync(ctx workflow.Context, input *iotevents.CreateDetectorModelInput) *IoTEventsCreateDetectorModelFuture

	CreateInput(ctx workflow.Context, input *iotevents.CreateInputInput) (*iotevents.CreateInputOutput, error)
	CreateInputAsync(ctx workflow.Context, input *iotevents.CreateInputInput) *IoTEventsCreateInputFuture

	DeleteDetectorModel(ctx workflow.Context, input *iotevents.DeleteDetectorModelInput) (*iotevents.DeleteDetectorModelOutput, error)
	DeleteDetectorModelAsync(ctx workflow.Context, input *iotevents.DeleteDetectorModelInput) *IoTEventsDeleteDetectorModelFuture

	DeleteInput(ctx workflow.Context, input *iotevents.DeleteInputInput) (*iotevents.DeleteInputOutput, error)
	DeleteInputAsync(ctx workflow.Context, input *iotevents.DeleteInputInput) *IoTEventsDeleteInputFuture

	DescribeDetectorModel(ctx workflow.Context, input *iotevents.DescribeDetectorModelInput) (*iotevents.DescribeDetectorModelOutput, error)
	DescribeDetectorModelAsync(ctx workflow.Context, input *iotevents.DescribeDetectorModelInput) *IoTEventsDescribeDetectorModelFuture

	DescribeInput(ctx workflow.Context, input *iotevents.DescribeInputInput) (*iotevents.DescribeInputOutput, error)
	DescribeInputAsync(ctx workflow.Context, input *iotevents.DescribeInputInput) *IoTEventsDescribeInputFuture

	DescribeLoggingOptions(ctx workflow.Context, input *iotevents.DescribeLoggingOptionsInput) (*iotevents.DescribeLoggingOptionsOutput, error)
	DescribeLoggingOptionsAsync(ctx workflow.Context, input *iotevents.DescribeLoggingOptionsInput) *IoTEventsDescribeLoggingOptionsFuture

	ListDetectorModelVersions(ctx workflow.Context, input *iotevents.ListDetectorModelVersionsInput) (*iotevents.ListDetectorModelVersionsOutput, error)
	ListDetectorModelVersionsAsync(ctx workflow.Context, input *iotevents.ListDetectorModelVersionsInput) *IoTEventsListDetectorModelVersionsFuture

	ListDetectorModels(ctx workflow.Context, input *iotevents.ListDetectorModelsInput) (*iotevents.ListDetectorModelsOutput, error)
	ListDetectorModelsAsync(ctx workflow.Context, input *iotevents.ListDetectorModelsInput) *IoTEventsListDetectorModelsFuture

	ListInputs(ctx workflow.Context, input *iotevents.ListInputsInput) (*iotevents.ListInputsOutput, error)
	ListInputsAsync(ctx workflow.Context, input *iotevents.ListInputsInput) *IoTEventsListInputsFuture

	ListTagsForResource(ctx workflow.Context, input *iotevents.ListTagsForResourceInput) (*iotevents.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *iotevents.ListTagsForResourceInput) *IoTEventsListTagsForResourceFuture

	PutLoggingOptions(ctx workflow.Context, input *iotevents.PutLoggingOptionsInput) (*iotevents.PutLoggingOptionsOutput, error)
	PutLoggingOptionsAsync(ctx workflow.Context, input *iotevents.PutLoggingOptionsInput) *IoTEventsPutLoggingOptionsFuture

	TagResource(ctx workflow.Context, input *iotevents.TagResourceInput) (*iotevents.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *iotevents.TagResourceInput) *IoTEventsTagResourceFuture

	UntagResource(ctx workflow.Context, input *iotevents.UntagResourceInput) (*iotevents.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *iotevents.UntagResourceInput) *IoTEventsUntagResourceFuture

	UpdateDetectorModel(ctx workflow.Context, input *iotevents.UpdateDetectorModelInput) (*iotevents.UpdateDetectorModelOutput, error)
	UpdateDetectorModelAsync(ctx workflow.Context, input *iotevents.UpdateDetectorModelInput) *IoTEventsUpdateDetectorModelFuture

	UpdateInput(ctx workflow.Context, input *iotevents.UpdateInputInput) (*iotevents.UpdateInputOutput, error)
	UpdateInputAsync(ctx workflow.Context, input *iotevents.UpdateInputInput) *IoTEventsUpdateInputFuture
}

func NewClient() Client {
	return &stub{}
}
