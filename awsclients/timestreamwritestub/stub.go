// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package timestreamwritestub

import (
	"github.com/aws/aws-sdk-go/service/timestreamwrite"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type TimestreamWriteCreateDatabaseFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TimestreamWriteCreateDatabaseFuture) Get(ctx workflow.Context) (*timestreamwrite.CreateDatabaseOutput, error) {
	var output timestreamwrite.CreateDatabaseOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TimestreamWriteCreateTableFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TimestreamWriteCreateTableFuture) Get(ctx workflow.Context) (*timestreamwrite.CreateTableOutput, error) {
	var output timestreamwrite.CreateTableOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TimestreamWriteDeleteDatabaseFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TimestreamWriteDeleteDatabaseFuture) Get(ctx workflow.Context) (*timestreamwrite.DeleteDatabaseOutput, error) {
	var output timestreamwrite.DeleteDatabaseOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TimestreamWriteDeleteTableFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TimestreamWriteDeleteTableFuture) Get(ctx workflow.Context) (*timestreamwrite.DeleteTableOutput, error) {
	var output timestreamwrite.DeleteTableOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TimestreamWriteDescribeDatabaseFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TimestreamWriteDescribeDatabaseFuture) Get(ctx workflow.Context) (*timestreamwrite.DescribeDatabaseOutput, error) {
	var output timestreamwrite.DescribeDatabaseOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TimestreamWriteDescribeEndpointsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TimestreamWriteDescribeEndpointsFuture) Get(ctx workflow.Context) (*timestreamwrite.DescribeEndpointsOutput, error) {
	var output timestreamwrite.DescribeEndpointsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TimestreamWriteDescribeTableFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TimestreamWriteDescribeTableFuture) Get(ctx workflow.Context) (*timestreamwrite.DescribeTableOutput, error) {
	var output timestreamwrite.DescribeTableOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TimestreamWriteListDatabasesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TimestreamWriteListDatabasesFuture) Get(ctx workflow.Context) (*timestreamwrite.ListDatabasesOutput, error) {
	var output timestreamwrite.ListDatabasesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TimestreamWriteListTablesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TimestreamWriteListTablesFuture) Get(ctx workflow.Context) (*timestreamwrite.ListTablesOutput, error) {
	var output timestreamwrite.ListTablesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TimestreamWriteListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TimestreamWriteListTagsForResourceFuture) Get(ctx workflow.Context) (*timestreamwrite.ListTagsForResourceOutput, error) {
	var output timestreamwrite.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TimestreamWriteTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TimestreamWriteTagResourceFuture) Get(ctx workflow.Context) (*timestreamwrite.TagResourceOutput, error) {
	var output timestreamwrite.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TimestreamWriteUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TimestreamWriteUntagResourceFuture) Get(ctx workflow.Context) (*timestreamwrite.UntagResourceOutput, error) {
	var output timestreamwrite.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TimestreamWriteUpdateDatabaseFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TimestreamWriteUpdateDatabaseFuture) Get(ctx workflow.Context) (*timestreamwrite.UpdateDatabaseOutput, error) {
	var output timestreamwrite.UpdateDatabaseOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TimestreamWriteUpdateTableFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TimestreamWriteUpdateTableFuture) Get(ctx workflow.Context) (*timestreamwrite.UpdateTableOutput, error) {
	var output timestreamwrite.UpdateTableOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TimestreamWriteWriteRecordsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TimestreamWriteWriteRecordsFuture) Get(ctx workflow.Context) (*timestreamwrite.WriteRecordsOutput, error) {
	var output timestreamwrite.WriteRecordsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDatabase(ctx workflow.Context, input *timestreamwrite.CreateDatabaseInput) (*timestreamwrite.CreateDatabaseOutput, error) {
	var output timestreamwrite.CreateDatabaseOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.CreateDatabase", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDatabaseAsync(ctx workflow.Context, input *timestreamwrite.CreateDatabaseInput) *TimestreamWriteCreateDatabaseFuture {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.CreateDatabase", input)
	return &TimestreamWriteCreateDatabaseFuture{Future: future}
}

func (a *stub) CreateTable(ctx workflow.Context, input *timestreamwrite.CreateTableInput) (*timestreamwrite.CreateTableOutput, error) {
	var output timestreamwrite.CreateTableOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.CreateTable", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateTableAsync(ctx workflow.Context, input *timestreamwrite.CreateTableInput) *TimestreamWriteCreateTableFuture {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.CreateTable", input)
	return &TimestreamWriteCreateTableFuture{Future: future}
}

func (a *stub) DeleteDatabase(ctx workflow.Context, input *timestreamwrite.DeleteDatabaseInput) (*timestreamwrite.DeleteDatabaseOutput, error) {
	var output timestreamwrite.DeleteDatabaseOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.DeleteDatabase", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDatabaseAsync(ctx workflow.Context, input *timestreamwrite.DeleteDatabaseInput) *TimestreamWriteDeleteDatabaseFuture {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.DeleteDatabase", input)
	return &TimestreamWriteDeleteDatabaseFuture{Future: future}
}

func (a *stub) DeleteTable(ctx workflow.Context, input *timestreamwrite.DeleteTableInput) (*timestreamwrite.DeleteTableOutput, error) {
	var output timestreamwrite.DeleteTableOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.DeleteTable", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteTableAsync(ctx workflow.Context, input *timestreamwrite.DeleteTableInput) *TimestreamWriteDeleteTableFuture {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.DeleteTable", input)
	return &TimestreamWriteDeleteTableFuture{Future: future}
}

func (a *stub) DescribeDatabase(ctx workflow.Context, input *timestreamwrite.DescribeDatabaseInput) (*timestreamwrite.DescribeDatabaseOutput, error) {
	var output timestreamwrite.DescribeDatabaseOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.DescribeDatabase", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDatabaseAsync(ctx workflow.Context, input *timestreamwrite.DescribeDatabaseInput) *TimestreamWriteDescribeDatabaseFuture {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.DescribeDatabase", input)
	return &TimestreamWriteDescribeDatabaseFuture{Future: future}
}

func (a *stub) DescribeEndpoints(ctx workflow.Context, input *timestreamwrite.DescribeEndpointsInput) (*timestreamwrite.DescribeEndpointsOutput, error) {
	var output timestreamwrite.DescribeEndpointsOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.DescribeEndpoints", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEndpointsAsync(ctx workflow.Context, input *timestreamwrite.DescribeEndpointsInput) *TimestreamWriteDescribeEndpointsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.DescribeEndpoints", input)
	return &TimestreamWriteDescribeEndpointsFuture{Future: future}
}

func (a *stub) DescribeTable(ctx workflow.Context, input *timestreamwrite.DescribeTableInput) (*timestreamwrite.DescribeTableOutput, error) {
	var output timestreamwrite.DescribeTableOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.DescribeTable", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeTableAsync(ctx workflow.Context, input *timestreamwrite.DescribeTableInput) *TimestreamWriteDescribeTableFuture {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.DescribeTable", input)
	return &TimestreamWriteDescribeTableFuture{Future: future}
}

func (a *stub) ListDatabases(ctx workflow.Context, input *timestreamwrite.ListDatabasesInput) (*timestreamwrite.ListDatabasesOutput, error) {
	var output timestreamwrite.ListDatabasesOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.ListDatabases", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDatabasesAsync(ctx workflow.Context, input *timestreamwrite.ListDatabasesInput) *TimestreamWriteListDatabasesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.ListDatabases", input)
	return &TimestreamWriteListDatabasesFuture{Future: future}
}

func (a *stub) ListTables(ctx workflow.Context, input *timestreamwrite.ListTablesInput) (*timestreamwrite.ListTablesOutput, error) {
	var output timestreamwrite.ListTablesOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.ListTables", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTablesAsync(ctx workflow.Context, input *timestreamwrite.ListTablesInput) *TimestreamWriteListTablesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.ListTables", input)
	return &TimestreamWriteListTablesFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *timestreamwrite.ListTagsForResourceInput) (*timestreamwrite.ListTagsForResourceOutput, error) {
	var output timestreamwrite.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *timestreamwrite.ListTagsForResourceInput) *TimestreamWriteListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.ListTagsForResource", input)
	return &TimestreamWriteListTagsForResourceFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *timestreamwrite.TagResourceInput) (*timestreamwrite.TagResourceOutput, error) {
	var output timestreamwrite.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *timestreamwrite.TagResourceInput) *TimestreamWriteTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.TagResource", input)
	return &TimestreamWriteTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *timestreamwrite.UntagResourceInput) (*timestreamwrite.UntagResourceOutput, error) {
	var output timestreamwrite.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *timestreamwrite.UntagResourceInput) *TimestreamWriteUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.UntagResource", input)
	return &TimestreamWriteUntagResourceFuture{Future: future}
}

func (a *stub) UpdateDatabase(ctx workflow.Context, input *timestreamwrite.UpdateDatabaseInput) (*timestreamwrite.UpdateDatabaseOutput, error) {
	var output timestreamwrite.UpdateDatabaseOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.UpdateDatabase", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateDatabaseAsync(ctx workflow.Context, input *timestreamwrite.UpdateDatabaseInput) *TimestreamWriteUpdateDatabaseFuture {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.UpdateDatabase", input)
	return &TimestreamWriteUpdateDatabaseFuture{Future: future}
}

func (a *stub) UpdateTable(ctx workflow.Context, input *timestreamwrite.UpdateTableInput) (*timestreamwrite.UpdateTableOutput, error) {
	var output timestreamwrite.UpdateTableOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.UpdateTable", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateTableAsync(ctx workflow.Context, input *timestreamwrite.UpdateTableInput) *TimestreamWriteUpdateTableFuture {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.UpdateTable", input)
	return &TimestreamWriteUpdateTableFuture{Future: future}
}

func (a *stub) WriteRecords(ctx workflow.Context, input *timestreamwrite.WriteRecordsInput) (*timestreamwrite.WriteRecordsOutput, error) {
	var output timestreamwrite.WriteRecordsOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.WriteRecords", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) WriteRecordsAsync(ctx workflow.Context, input *timestreamwrite.WriteRecordsInput) *TimestreamWriteWriteRecordsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.WriteRecords", input)
	return &TimestreamWriteWriteRecordsFuture{Future: future}
}
