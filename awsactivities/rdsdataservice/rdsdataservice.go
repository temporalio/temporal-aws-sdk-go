// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package rdsdataservice

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rdsdataservice"
	"github.com/aws/aws-sdk-go/service/rdsdataservice/rdsdataserviceiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client rdsdataserviceiface.RDSDataServiceAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := rdsdataservice.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (rdsdataserviceiface.RDSDataServiceAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return rdsdataservice.New(sess), nil
}

func (a *Activities) BatchExecuteStatement(ctx context.Context, input *rdsdataservice.BatchExecuteStatementInput) (*rdsdataservice.BatchExecuteStatementOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchExecuteStatementWithContext(ctx, input)
}

func (a *Activities) BeginTransaction(ctx context.Context, input *rdsdataservice.BeginTransactionInput) (*rdsdataservice.BeginTransactionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BeginTransactionWithContext(ctx, input)
}

func (a *Activities) CommitTransaction(ctx context.Context, input *rdsdataservice.CommitTransactionInput) (*rdsdataservice.CommitTransactionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CommitTransactionWithContext(ctx, input)
}

func (a *Activities) ExecuteSql(ctx context.Context, input *rdsdataservice.ExecuteSqlInput) (*rdsdataservice.ExecuteSqlOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ExecuteSqlWithContext(ctx, input)
}

func (a *Activities) ExecuteStatement(ctx context.Context, input *rdsdataservice.ExecuteStatementInput) (*rdsdataservice.ExecuteStatementOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ExecuteStatementWithContext(ctx, input)
}

func (a *Activities) RollbackTransaction(ctx context.Context, input *rdsdataservice.RollbackTransactionInput) (*rdsdataservice.RollbackTransactionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RollbackTransactionWithContext(ctx, input)
}