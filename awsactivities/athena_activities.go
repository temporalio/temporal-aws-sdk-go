
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/aws/aws-sdk-go/service/athena/athenaiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type AthenaActivities struct {
    client athenaiface.AthenaAPI
}

func NewAthenaActivities(session *session.Session, config ...*aws.Config) *AthenaActivities {
    client := athena.New(session, config...)
    return &AthenaActivities{client: client}
}

func (a *AthenaActivities) BatchGetNamedQuery(ctx context.Context, input *athena.BatchGetNamedQueryInput) (*athena.BatchGetNamedQueryOutput, error) {
    return a.client.BatchGetNamedQueryWithContext(ctx, input)
}

func (a *AthenaActivities) BatchGetQueryExecution(ctx context.Context, input *athena.BatchGetQueryExecutionInput) (*athena.BatchGetQueryExecutionOutput, error) {
    return a.client.BatchGetQueryExecutionWithContext(ctx, input)
}

func (a *AthenaActivities) CreateDataCatalog(ctx context.Context, input *athena.CreateDataCatalogInput) (*athena.CreateDataCatalogOutput, error) {
    return a.client.CreateDataCatalogWithContext(ctx, input)
}

func (a *AthenaActivities) CreateNamedQuery(ctx context.Context, input *athena.CreateNamedQueryInput) (*athena.CreateNamedQueryOutput, error) {
    return a.client.CreateNamedQueryWithContext(ctx, input)
}

func (a *AthenaActivities) CreateWorkGroup(ctx context.Context, input *athena.CreateWorkGroupInput) (*athena.CreateWorkGroupOutput, error) {
    return a.client.CreateWorkGroupWithContext(ctx, input)
}

func (a *AthenaActivities) DeleteDataCatalog(ctx context.Context, input *athena.DeleteDataCatalogInput) (*athena.DeleteDataCatalogOutput, error) {
    return a.client.DeleteDataCatalogWithContext(ctx, input)
}

func (a *AthenaActivities) DeleteNamedQuery(ctx context.Context, input *athena.DeleteNamedQueryInput) (*athena.DeleteNamedQueryOutput, error) {
    return a.client.DeleteNamedQueryWithContext(ctx, input)
}

func (a *AthenaActivities) DeleteWorkGroup(ctx context.Context, input *athena.DeleteWorkGroupInput) (*athena.DeleteWorkGroupOutput, error) {
    return a.client.DeleteWorkGroupWithContext(ctx, input)
}

func (a *AthenaActivities) GetDataCatalog(ctx context.Context, input *athena.GetDataCatalogInput) (*athena.GetDataCatalogOutput, error) {
    return a.client.GetDataCatalogWithContext(ctx, input)
}

func (a *AthenaActivities) GetDatabase(ctx context.Context, input *athena.GetDatabaseInput) (*athena.GetDatabaseOutput, error) {
    return a.client.GetDatabaseWithContext(ctx, input)
}

func (a *AthenaActivities) GetNamedQuery(ctx context.Context, input *athena.GetNamedQueryInput) (*athena.GetNamedQueryOutput, error) {
    return a.client.GetNamedQueryWithContext(ctx, input)
}

func (a *AthenaActivities) GetQueryExecution(ctx context.Context, input *athena.GetQueryExecutionInput) (*athena.GetQueryExecutionOutput, error) {
    return a.client.GetQueryExecutionWithContext(ctx, input)
}

func (a *AthenaActivities) GetQueryResults(ctx context.Context, input *athena.GetQueryResultsInput) (*athena.GetQueryResultsOutput, error) {
    return a.client.GetQueryResultsWithContext(ctx, input)
}

func (a *AthenaActivities) GetTableMetadata(ctx context.Context, input *athena.GetTableMetadataInput) (*athena.GetTableMetadataOutput, error) {
    return a.client.GetTableMetadataWithContext(ctx, input)
}

func (a *AthenaActivities) GetWorkGroup(ctx context.Context, input *athena.GetWorkGroupInput) (*athena.GetWorkGroupOutput, error) {
    return a.client.GetWorkGroupWithContext(ctx, input)
}

func (a *AthenaActivities) ListDataCatalogs(ctx context.Context, input *athena.ListDataCatalogsInput) (*athena.ListDataCatalogsOutput, error) {
    return a.client.ListDataCatalogsWithContext(ctx, input)
}

func (a *AthenaActivities) ListDatabases(ctx context.Context, input *athena.ListDatabasesInput) (*athena.ListDatabasesOutput, error) {
    return a.client.ListDatabasesWithContext(ctx, input)
}

func (a *AthenaActivities) ListNamedQueries(ctx context.Context, input *athena.ListNamedQueriesInput) (*athena.ListNamedQueriesOutput, error) {
    return a.client.ListNamedQueriesWithContext(ctx, input)
}

func (a *AthenaActivities) ListQueryExecutions(ctx context.Context, input *athena.ListQueryExecutionsInput) (*athena.ListQueryExecutionsOutput, error) {
    return a.client.ListQueryExecutionsWithContext(ctx, input)
}

func (a *AthenaActivities) ListTableMetadata(ctx context.Context, input *athena.ListTableMetadataInput) (*athena.ListTableMetadataOutput, error) {
    return a.client.ListTableMetadataWithContext(ctx, input)
}

func (a *AthenaActivities) ListTagsForResource(ctx context.Context, input *athena.ListTagsForResourceInput) (*athena.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *AthenaActivities) ListWorkGroups(ctx context.Context, input *athena.ListWorkGroupsInput) (*athena.ListWorkGroupsOutput, error) {
    return a.client.ListWorkGroupsWithContext(ctx, input)
}

func (a *AthenaActivities) StartQueryExecution(ctx context.Context, input *athena.StartQueryExecutionInput) (*athena.StartQueryExecutionOutput, error) {
    return a.client.StartQueryExecutionWithContext(ctx, input)
}

func (a *AthenaActivities) StopQueryExecution(ctx context.Context, input *athena.StopQueryExecutionInput) (*athena.StopQueryExecutionOutput, error) {
    return a.client.StopQueryExecutionWithContext(ctx, input)
}

func (a *AthenaActivities) TagResource(ctx context.Context, input *athena.TagResourceInput) (*athena.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *AthenaActivities) UntagResource(ctx context.Context, input *athena.UntagResourceInput) (*athena.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *AthenaActivities) UpdateDataCatalog(ctx context.Context, input *athena.UpdateDataCatalogInput) (*athena.UpdateDataCatalogOutput, error) {
    return a.client.UpdateDataCatalogWithContext(ctx, input)
}

func (a *AthenaActivities) UpdateWorkGroup(ctx context.Context, input *athena.UpdateWorkGroupInput) (*athena.UpdateWorkGroupOutput, error) {
    return a.client.UpdateWorkGroupWithContext(ctx, input)
}
