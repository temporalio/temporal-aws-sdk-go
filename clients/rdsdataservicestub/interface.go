// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package rdsdataservicestub

import (
	"github.com/aws/aws-sdk-go/service/rdsdataservice"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	BatchExecuteStatement(ctx workflow.Context, input *rdsdataservice.BatchExecuteStatementInput) (*rdsdataservice.BatchExecuteStatementOutput, error)
	BatchExecuteStatementAsync(ctx workflow.Context, input *rdsdataservice.BatchExecuteStatementInput) *BatchExecuteStatementFuture

	BeginTransaction(ctx workflow.Context, input *rdsdataservice.BeginTransactionInput) (*rdsdataservice.BeginTransactionOutput, error)
	BeginTransactionAsync(ctx workflow.Context, input *rdsdataservice.BeginTransactionInput) *BeginTransactionFuture

	CommitTransaction(ctx workflow.Context, input *rdsdataservice.CommitTransactionInput) (*rdsdataservice.CommitTransactionOutput, error)
	CommitTransactionAsync(ctx workflow.Context, input *rdsdataservice.CommitTransactionInput) *CommitTransactionFuture

	ExecuteSql(ctx workflow.Context, input *rdsdataservice.ExecuteSqlInput) (*rdsdataservice.ExecuteSqlOutput, error)
	ExecuteSqlAsync(ctx workflow.Context, input *rdsdataservice.ExecuteSqlInput) *ExecuteSqlFuture

	ExecuteStatement(ctx workflow.Context, input *rdsdataservice.ExecuteStatementInput) (*rdsdataservice.ExecuteStatementOutput, error)
	ExecuteStatementAsync(ctx workflow.Context, input *rdsdataservice.ExecuteStatementInput) *ExecuteStatementFuture

	RollbackTransaction(ctx workflow.Context, input *rdsdataservice.RollbackTransactionInput) (*rdsdataservice.RollbackTransactionOutput, error)
	RollbackTransactionAsync(ctx workflow.Context, input *rdsdataservice.RollbackTransactionInput) *RollbackTransactionFuture
}

func NewClient() Client {
	return &stub{}
}
