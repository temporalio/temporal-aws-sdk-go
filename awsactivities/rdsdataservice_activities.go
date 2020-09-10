package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rdsdataservice"
	"github.com/aws/aws-sdk-go/service/rdsdataservice/rdsdataserviceiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type RDSDataServiceActivities struct {
	client rdsdataserviceiface.RDSDataServiceAPI
}

func NewRDSDataServiceActivities(session *session.Session, config ...*aws.Config) *RDSDataServiceActivities {
	client := rdsdataservice.New(session, config...)
	return &RDSDataServiceActivities{client: client}
}

func (a *RDSDataServiceActivities) BatchExecuteStatement(ctx context.Context, input *rdsdataservice.BatchExecuteStatementInput) (*rdsdataservice.BatchExecuteStatementOutput, error) {
	return a.client.BatchExecuteStatementWithContext(ctx, input)
}

func (a *RDSDataServiceActivities) BeginTransaction(ctx context.Context, input *rdsdataservice.BeginTransactionInput) (*rdsdataservice.BeginTransactionOutput, error) {
	return a.client.BeginTransactionWithContext(ctx, input)
}

func (a *RDSDataServiceActivities) CommitTransaction(ctx context.Context, input *rdsdataservice.CommitTransactionInput) (*rdsdataservice.CommitTransactionOutput, error) {
	return a.client.CommitTransactionWithContext(ctx, input)
}

func (a *RDSDataServiceActivities) ExecuteSql(ctx context.Context, input *rdsdataservice.ExecuteSqlInput) (*rdsdataservice.ExecuteSqlOutput, error) {
	return a.client.ExecuteSqlWithContext(ctx, input)
}

func (a *RDSDataServiceActivities) ExecuteStatement(ctx context.Context, input *rdsdataservice.ExecuteStatementInput) (*rdsdataservice.ExecuteStatementOutput, error) {
	return a.client.ExecuteStatementWithContext(ctx, input)
}

func (a *RDSDataServiceActivities) RollbackTransaction(ctx context.Context, input *rdsdataservice.RollbackTransactionInput) (*rdsdataservice.RollbackTransactionOutput, error) {
	return a.client.RollbackTransactionWithContext(ctx, input)
}
