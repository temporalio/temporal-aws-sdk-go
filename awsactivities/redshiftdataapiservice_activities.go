package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/redshiftdataapiservice"
	"github.com/aws/aws-sdk-go/service/redshiftdataapiservice/redshiftdataapiserviceiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type RedshiftDataAPIServiceActivities struct {
	client redshiftdataapiserviceiface.RedshiftDataAPIServiceAPI
}

func NewRedshiftDataAPIServiceActivities(session *session.Session, config ...*aws.Config) *RedshiftDataAPIServiceActivities {
	client := redshiftdataapiservice.New(session, config...)
	return &RedshiftDataAPIServiceActivities{client: client}
}

func (a *RedshiftDataAPIServiceActivities) CancelStatement(ctx context.Context, input *redshiftdataapiservice.CancelStatementInput) (*redshiftdataapiservice.CancelStatementOutput, error) {
	return a.client.CancelStatementWithContext(ctx, input)
}

func (a *RedshiftDataAPIServiceActivities) DescribeStatement(ctx context.Context, input *redshiftdataapiservice.DescribeStatementInput) (*redshiftdataapiservice.DescribeStatementOutput, error) {
	return a.client.DescribeStatementWithContext(ctx, input)
}

func (a *RedshiftDataAPIServiceActivities) DescribeTable(ctx context.Context, input *redshiftdataapiservice.DescribeTableInput) (*redshiftdataapiservice.DescribeTableOutput, error) {
	return a.client.DescribeTableWithContext(ctx, input)
}

func (a *RedshiftDataAPIServiceActivities) ExecuteStatement(ctx context.Context, input *redshiftdataapiservice.ExecuteStatementInput) (*redshiftdataapiservice.ExecuteStatementOutput, error) {
	return a.client.ExecuteStatementWithContext(ctx, input)
}

func (a *RedshiftDataAPIServiceActivities) GetStatementResult(ctx context.Context, input *redshiftdataapiservice.GetStatementResultInput) (*redshiftdataapiservice.GetStatementResultOutput, error) {
	return a.client.GetStatementResultWithContext(ctx, input)
}

func (a *RedshiftDataAPIServiceActivities) ListDatabases(ctx context.Context, input *redshiftdataapiservice.ListDatabasesInput) (*redshiftdataapiservice.ListDatabasesOutput, error) {
	return a.client.ListDatabasesWithContext(ctx, input)
}

func (a *RedshiftDataAPIServiceActivities) ListSchemas(ctx context.Context, input *redshiftdataapiservice.ListSchemasInput) (*redshiftdataapiservice.ListSchemasOutput, error) {
	return a.client.ListSchemasWithContext(ctx, input)
}

func (a *RedshiftDataAPIServiceActivities) ListStatements(ctx context.Context, input *redshiftdataapiservice.ListStatementsInput) (*redshiftdataapiservice.ListStatementsOutput, error) {
	return a.client.ListStatementsWithContext(ctx, input)
}

func (a *RedshiftDataAPIServiceActivities) ListTables(ctx context.Context, input *redshiftdataapiservice.ListTablesInput) (*redshiftdataapiservice.ListTablesOutput, error) {
	return a.client.ListTablesWithContext(ctx, input)
}
