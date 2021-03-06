// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package appflowstub

import (
	"github.com/aws/aws-sdk-go/service/appflow"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateConnectorProfile(ctx workflow.Context, input *appflow.CreateConnectorProfileInput) (*appflow.CreateConnectorProfileOutput, error)
	CreateConnectorProfileAsync(ctx workflow.Context, input *appflow.CreateConnectorProfileInput) *CreateConnectorProfileFuture

	CreateFlow(ctx workflow.Context, input *appflow.CreateFlowInput) (*appflow.CreateFlowOutput, error)
	CreateFlowAsync(ctx workflow.Context, input *appflow.CreateFlowInput) *CreateFlowFuture

	DeleteConnectorProfile(ctx workflow.Context, input *appflow.DeleteConnectorProfileInput) (*appflow.DeleteConnectorProfileOutput, error)
	DeleteConnectorProfileAsync(ctx workflow.Context, input *appflow.DeleteConnectorProfileInput) *DeleteConnectorProfileFuture

	DeleteFlow(ctx workflow.Context, input *appflow.DeleteFlowInput) (*appflow.DeleteFlowOutput, error)
	DeleteFlowAsync(ctx workflow.Context, input *appflow.DeleteFlowInput) *DeleteFlowFuture

	DescribeConnectorEntity(ctx workflow.Context, input *appflow.DescribeConnectorEntityInput) (*appflow.DescribeConnectorEntityOutput, error)
	DescribeConnectorEntityAsync(ctx workflow.Context, input *appflow.DescribeConnectorEntityInput) *DescribeConnectorEntityFuture

	DescribeConnectorProfiles(ctx workflow.Context, input *appflow.DescribeConnectorProfilesInput) (*appflow.DescribeConnectorProfilesOutput, error)
	DescribeConnectorProfilesAsync(ctx workflow.Context, input *appflow.DescribeConnectorProfilesInput) *DescribeConnectorProfilesFuture

	DescribeConnectors(ctx workflow.Context, input *appflow.DescribeConnectorsInput) (*appflow.DescribeConnectorsOutput, error)
	DescribeConnectorsAsync(ctx workflow.Context, input *appflow.DescribeConnectorsInput) *DescribeConnectorsFuture

	DescribeFlow(ctx workflow.Context, input *appflow.DescribeFlowInput) (*appflow.DescribeFlowOutput, error)
	DescribeFlowAsync(ctx workflow.Context, input *appflow.DescribeFlowInput) *DescribeFlowFuture

	DescribeFlowExecutionRecords(ctx workflow.Context, input *appflow.DescribeFlowExecutionRecordsInput) (*appflow.DescribeFlowExecutionRecordsOutput, error)
	DescribeFlowExecutionRecordsAsync(ctx workflow.Context, input *appflow.DescribeFlowExecutionRecordsInput) *DescribeFlowExecutionRecordsFuture

	ListConnectorEntities(ctx workflow.Context, input *appflow.ListConnectorEntitiesInput) (*appflow.ListConnectorEntitiesOutput, error)
	ListConnectorEntitiesAsync(ctx workflow.Context, input *appflow.ListConnectorEntitiesInput) *ListConnectorEntitiesFuture

	ListFlows(ctx workflow.Context, input *appflow.ListFlowsInput) (*appflow.ListFlowsOutput, error)
	ListFlowsAsync(ctx workflow.Context, input *appflow.ListFlowsInput) *ListFlowsFuture

	ListTagsForResource(ctx workflow.Context, input *appflow.ListTagsForResourceInput) (*appflow.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *appflow.ListTagsForResourceInput) *ListTagsForResourceFuture

	StartFlow(ctx workflow.Context, input *appflow.StartFlowInput) (*appflow.StartFlowOutput, error)
	StartFlowAsync(ctx workflow.Context, input *appflow.StartFlowInput) *StartFlowFuture

	StopFlow(ctx workflow.Context, input *appflow.StopFlowInput) (*appflow.StopFlowOutput, error)
	StopFlowAsync(ctx workflow.Context, input *appflow.StopFlowInput) *StopFlowFuture

	TagResource(ctx workflow.Context, input *appflow.TagResourceInput) (*appflow.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *appflow.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *appflow.UntagResourceInput) (*appflow.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *appflow.UntagResourceInput) *UntagResourceFuture

	UpdateConnectorProfile(ctx workflow.Context, input *appflow.UpdateConnectorProfileInput) (*appflow.UpdateConnectorProfileOutput, error)
	UpdateConnectorProfileAsync(ctx workflow.Context, input *appflow.UpdateConnectorProfileInput) *UpdateConnectorProfileFuture

	UpdateFlow(ctx workflow.Context, input *appflow.UpdateFlowInput) (*appflow.UpdateFlowOutput, error)
	UpdateFlowAsync(ctx workflow.Context, input *appflow.UpdateFlowInput) *UpdateFlowFuture
}

func NewClient() Client {
	return &stub{}
}
