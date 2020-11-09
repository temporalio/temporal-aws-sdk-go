// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package redshiftdataapiservicestub

import (
	"github.com/aws/aws-sdk-go/service/redshiftdataapiservice"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CancelStatement(ctx workflow.Context, input *redshiftdataapiservice.CancelStatementInput) (*redshiftdataapiservice.CancelStatementOutput, error)
	CancelStatementAsync(ctx workflow.Context, input *redshiftdataapiservice.CancelStatementInput) *CancelStatementFuture

	DescribeStatement(ctx workflow.Context, input *redshiftdataapiservice.DescribeStatementInput) (*redshiftdataapiservice.DescribeStatementOutput, error)
	DescribeStatementAsync(ctx workflow.Context, input *redshiftdataapiservice.DescribeStatementInput) *DescribeStatementFuture

	DescribeTable(ctx workflow.Context, input *redshiftdataapiservice.DescribeTableInput) (*redshiftdataapiservice.DescribeTableOutput, error)
	DescribeTableAsync(ctx workflow.Context, input *redshiftdataapiservice.DescribeTableInput) *DescribeTableFuture

	ExecuteStatement(ctx workflow.Context, input *redshiftdataapiservice.ExecuteStatementInput) (*redshiftdataapiservice.ExecuteStatementOutput, error)
	ExecuteStatementAsync(ctx workflow.Context, input *redshiftdataapiservice.ExecuteStatementInput) *ExecuteStatementFuture

	GetStatementResult(ctx workflow.Context, input *redshiftdataapiservice.GetStatementResultInput) (*redshiftdataapiservice.GetStatementResultOutput, error)
	GetStatementResultAsync(ctx workflow.Context, input *redshiftdataapiservice.GetStatementResultInput) *GetStatementResultFuture

	ListDatabases(ctx workflow.Context, input *redshiftdataapiservice.ListDatabasesInput) (*redshiftdataapiservice.ListDatabasesOutput, error)
	ListDatabasesAsync(ctx workflow.Context, input *redshiftdataapiservice.ListDatabasesInput) *ListDatabasesFuture

	ListSchemas(ctx workflow.Context, input *redshiftdataapiservice.ListSchemasInput) (*redshiftdataapiservice.ListSchemasOutput, error)
	ListSchemasAsync(ctx workflow.Context, input *redshiftdataapiservice.ListSchemasInput) *ListSchemasFuture

	ListStatements(ctx workflow.Context, input *redshiftdataapiservice.ListStatementsInput) (*redshiftdataapiservice.ListStatementsOutput, error)
	ListStatementsAsync(ctx workflow.Context, input *redshiftdataapiservice.ListStatementsInput) *ListStatementsFuture

	ListTables(ctx workflow.Context, input *redshiftdataapiservice.ListTablesInput) (*redshiftdataapiservice.ListTablesOutput, error)
	ListTablesAsync(ctx workflow.Context, input *redshiftdataapiservice.ListTablesInput) *ListTablesFuture
}

func NewClient() Client {
	return &stub{}
}
