// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package migrationhubconfigstub

import (
	"github.com/aws/aws-sdk-go/service/migrationhubconfig"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type MigrationHubConfigCreateHomeRegionControlFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MigrationHubConfigCreateHomeRegionControlFuture) Get(ctx workflow.Context) (*migrationhubconfig.CreateHomeRegionControlOutput, error) {
	var output migrationhubconfig.CreateHomeRegionControlOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MigrationHubConfigDescribeHomeRegionControlsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MigrationHubConfigDescribeHomeRegionControlsFuture) Get(ctx workflow.Context) (*migrationhubconfig.DescribeHomeRegionControlsOutput, error) {
	var output migrationhubconfig.DescribeHomeRegionControlsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MigrationHubConfigGetHomeRegionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MigrationHubConfigGetHomeRegionFuture) Get(ctx workflow.Context) (*migrationhubconfig.GetHomeRegionOutput, error) {
	var output migrationhubconfig.GetHomeRegionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateHomeRegionControl(ctx workflow.Context, input *migrationhubconfig.CreateHomeRegionControlInput) (*migrationhubconfig.CreateHomeRegionControlOutput, error) {
	var output migrationhubconfig.CreateHomeRegionControlOutput
	err := workflow.ExecuteActivity(ctx, "aws.migrationhubconfig.CreateHomeRegionControl", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateHomeRegionControlAsync(ctx workflow.Context, input *migrationhubconfig.CreateHomeRegionControlInput) *MigrationHubConfigCreateHomeRegionControlFuture {
	future := workflow.ExecuteActivity(ctx, "aws.migrationhubconfig.CreateHomeRegionControl", input)
	return &MigrationHubConfigCreateHomeRegionControlFuture{Future: future}
}

func (a *stub) DescribeHomeRegionControls(ctx workflow.Context, input *migrationhubconfig.DescribeHomeRegionControlsInput) (*migrationhubconfig.DescribeHomeRegionControlsOutput, error) {
	var output migrationhubconfig.DescribeHomeRegionControlsOutput
	err := workflow.ExecuteActivity(ctx, "aws.migrationhubconfig.DescribeHomeRegionControls", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeHomeRegionControlsAsync(ctx workflow.Context, input *migrationhubconfig.DescribeHomeRegionControlsInput) *MigrationHubConfigDescribeHomeRegionControlsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.migrationhubconfig.DescribeHomeRegionControls", input)
	return &MigrationHubConfigDescribeHomeRegionControlsFuture{Future: future}
}

func (a *stub) GetHomeRegion(ctx workflow.Context, input *migrationhubconfig.GetHomeRegionInput) (*migrationhubconfig.GetHomeRegionOutput, error) {
	var output migrationhubconfig.GetHomeRegionOutput
	err := workflow.ExecuteActivity(ctx, "aws.migrationhubconfig.GetHomeRegion", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetHomeRegionAsync(ctx workflow.Context, input *migrationhubconfig.GetHomeRegionInput) *MigrationHubConfigGetHomeRegionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.migrationhubconfig.GetHomeRegion", input)
	return &MigrationHubConfigGetHomeRegionFuture{Future: future}
}
