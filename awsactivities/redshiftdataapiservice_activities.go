// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/redshiftdataapiservice"
	"github.com/aws/aws-sdk-go/service/redshiftdataapiservice/redshiftdataapiserviceiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type RedshiftDataAPIServiceActivities struct {
	client redshiftdataapiserviceiface.RedshiftDataAPIServiceAPI

	sessionFactory SessionFactory
}

func NewRedshiftDataAPIServiceActivities(sess *session.Session, config ...*aws.Config) *RedshiftDataAPIServiceActivities {
	client := redshiftdataapiservice.New(sess, config...)
	return &RedshiftDataAPIServiceActivities{client: client}
}

func NewRedshiftDataAPIServiceActivitiesWithSessionFactory(sessionFactory SessionFactory) *RedshiftDataAPIServiceActivities {
	return &RedshiftDataAPIServiceActivities{sessionFactory: sessionFactory}
}

func (a *RedshiftDataAPIServiceActivities) getClient(ctx context.Context) (redshiftdataapiserviceiface.RedshiftDataAPIServiceAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return redshiftdataapiservice.New(sess), nil
}

func (a *RedshiftDataAPIServiceActivities) CancelStatement(ctx context.Context, input *redshiftdataapiservice.CancelStatementInput) (*redshiftdataapiservice.CancelStatementOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CancelStatementWithContext(ctx, input)
}

func (a *RedshiftDataAPIServiceActivities) DescribeStatement(ctx context.Context, input *redshiftdataapiservice.DescribeStatementInput) (*redshiftdataapiservice.DescribeStatementOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeStatementWithContext(ctx, input)
}

func (a *RedshiftDataAPIServiceActivities) DescribeTable(ctx context.Context, input *redshiftdataapiservice.DescribeTableInput) (*redshiftdataapiservice.DescribeTableOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeTableWithContext(ctx, input)
}

func (a *RedshiftDataAPIServiceActivities) ExecuteStatement(ctx context.Context, input *redshiftdataapiservice.ExecuteStatementInput) (*redshiftdataapiservice.ExecuteStatementOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ExecuteStatementWithContext(ctx, input)
}

func (a *RedshiftDataAPIServiceActivities) GetStatementResult(ctx context.Context, input *redshiftdataapiservice.GetStatementResultInput) (*redshiftdataapiservice.GetStatementResultOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetStatementResultWithContext(ctx, input)
}

func (a *RedshiftDataAPIServiceActivities) ListDatabases(ctx context.Context, input *redshiftdataapiservice.ListDatabasesInput) (*redshiftdataapiservice.ListDatabasesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDatabasesWithContext(ctx, input)
}

func (a *RedshiftDataAPIServiceActivities) ListSchemas(ctx context.Context, input *redshiftdataapiservice.ListSchemasInput) (*redshiftdataapiservice.ListSchemasOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListSchemasWithContext(ctx, input)
}

func (a *RedshiftDataAPIServiceActivities) ListStatements(ctx context.Context, input *redshiftdataapiservice.ListStatementsInput) (*redshiftdataapiservice.ListStatementsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListStatementsWithContext(ctx, input)
}

func (a *RedshiftDataAPIServiceActivities) ListTables(ctx context.Context, input *redshiftdataapiservice.ListTablesInput) (*redshiftdataapiservice.ListTablesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTablesWithContext(ctx, input)
}
