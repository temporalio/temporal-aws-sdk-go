// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package ioteventsdatastub

import (
	"github.com/aws/aws-sdk-go/service/ioteventsdata"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

type Client interface {
	BatchPutMessage(ctx workflow.Context, input *ioteventsdata.BatchPutMessageInput) (*ioteventsdata.BatchPutMessageOutput, error)
	BatchPutMessageAsync(ctx workflow.Context, input *ioteventsdata.BatchPutMessageInput) *IoTEventsDataBatchPutMessageFuture

	BatchUpdateDetector(ctx workflow.Context, input *ioteventsdata.BatchUpdateDetectorInput) (*ioteventsdata.BatchUpdateDetectorOutput, error)
	BatchUpdateDetectorAsync(ctx workflow.Context, input *ioteventsdata.BatchUpdateDetectorInput) *IoTEventsDataBatchUpdateDetectorFuture

	DescribeDetector(ctx workflow.Context, input *ioteventsdata.DescribeDetectorInput) (*ioteventsdata.DescribeDetectorOutput, error)
	DescribeDetectorAsync(ctx workflow.Context, input *ioteventsdata.DescribeDetectorInput) *IoTEventsDataDescribeDetectorFuture

	ListDetectors(ctx workflow.Context, input *ioteventsdata.ListDetectorsInput) (*ioteventsdata.ListDetectorsOutput, error)
	ListDetectorsAsync(ctx workflow.Context, input *ioteventsdata.ListDetectorsInput) *IoTEventsDataListDetectorsFuture
}

func NewClient() Client {
	return &stub{}
}

