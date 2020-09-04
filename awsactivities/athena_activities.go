
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/aws/aws-sdk-go/service/athena/athenaiface"
)

type AthenaActivities struct {
	client athenaiface.AthenaAPI
}

func NewAthenaActivities(client athenaiface.AthenaAPI) *AthenaActivities {
    return &AthenaActivities{client: client}
}

func (a *AthenaActivities) BatchGetNamedQuery(input *athena.BatchGetNamedQueryInput) (*athena.BatchGetNamedQueryOutput, error) {
    return a.client.BatchGetNamedQuery(input)
}

func (a *AthenaActivities) BatchGetQueryExecution(input *athena.BatchGetQueryExecutionInput) (*athena.BatchGetQueryExecutionOutput, error) {
    return a.client.BatchGetQueryExecution(input)
}

func (a *AthenaActivities) CreateDataCatalog(input *athena.CreateDataCatalogInput) (*athena.CreateDataCatalogOutput, error) {
    return a.client.CreateDataCatalog(input)
}

func (a *AthenaActivities) CreateNamedQuery(input *athena.CreateNamedQueryInput) (*athena.CreateNamedQueryOutput, error) {
    return a.client.CreateNamedQuery(input)
}

func (a *AthenaActivities) CreateWorkGroup(input *athena.CreateWorkGroupInput) (*athena.CreateWorkGroupOutput, error) {
    return a.client.CreateWorkGroup(input)
}

func (a *AthenaActivities) DeleteDataCatalog(input *athena.DeleteDataCatalogInput) (*athena.DeleteDataCatalogOutput, error) {
    return a.client.DeleteDataCatalog(input)
}

func (a *AthenaActivities) DeleteNamedQuery(input *athena.DeleteNamedQueryInput) (*athena.DeleteNamedQueryOutput, error) {
    return a.client.DeleteNamedQuery(input)
}

func (a *AthenaActivities) DeleteWorkGroup(input *athena.DeleteWorkGroupInput) (*athena.DeleteWorkGroupOutput, error) {
    return a.client.DeleteWorkGroup(input)
}

func (a *AthenaActivities) GetDataCatalog(input *athena.GetDataCatalogInput) (*athena.GetDataCatalogOutput, error) {
    return a.client.GetDataCatalog(input)
}

func (a *AthenaActivities) GetDatabase(input *athena.GetDatabaseInput) (*athena.GetDatabaseOutput, error) {
    return a.client.GetDatabase(input)
}

func (a *AthenaActivities) GetNamedQuery(input *athena.GetNamedQueryInput) (*athena.GetNamedQueryOutput, error) {
    return a.client.GetNamedQuery(input)
}

func (a *AthenaActivities) GetQueryExecution(input *athena.GetQueryExecutionInput) (*athena.GetQueryExecutionOutput, error) {
    return a.client.GetQueryExecution(input)
}

func (a *AthenaActivities) GetQueryResults(input *athena.GetQueryResultsInput) (*athena.GetQueryResultsOutput, error) {
    return a.client.GetQueryResults(input)
}

func (a *AthenaActivities) GetTableMetadata(input *athena.GetTableMetadataInput) (*athena.GetTableMetadataOutput, error) {
    return a.client.GetTableMetadata(input)
}

func (a *AthenaActivities) GetWorkGroup(input *athena.GetWorkGroupInput) (*athena.GetWorkGroupOutput, error) {
    return a.client.GetWorkGroup(input)
}

func (a *AthenaActivities) ListDataCatalogs(input *athena.ListDataCatalogsInput) (*athena.ListDataCatalogsOutput, error) {
    return a.client.ListDataCatalogs(input)
}

func (a *AthenaActivities) ListDatabases(input *athena.ListDatabasesInput) (*athena.ListDatabasesOutput, error) {
    return a.client.ListDatabases(input)
}

func (a *AthenaActivities) ListNamedQueries(input *athena.ListNamedQueriesInput) (*athena.ListNamedQueriesOutput, error) {
    return a.client.ListNamedQueries(input)
}

func (a *AthenaActivities) ListQueryExecutions(input *athena.ListQueryExecutionsInput) (*athena.ListQueryExecutionsOutput, error) {
    return a.client.ListQueryExecutions(input)
}

func (a *AthenaActivities) ListTableMetadata(input *athena.ListTableMetadataInput) (*athena.ListTableMetadataOutput, error) {
    return a.client.ListTableMetadata(input)
}

func (a *AthenaActivities) ListTagsForResource(input *athena.ListTagsForResourceInput) (*athena.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *AthenaActivities) ListWorkGroups(input *athena.ListWorkGroupsInput) (*athena.ListWorkGroupsOutput, error) {
    return a.client.ListWorkGroups(input)
}

func (a *AthenaActivities) StartQueryExecution(input *athena.StartQueryExecutionInput) (*athena.StartQueryExecutionOutput, error) {
    return a.client.StartQueryExecution(input)
}

func (a *AthenaActivities) StopQueryExecution(input *athena.StopQueryExecutionInput) (*athena.StopQueryExecutionOutput, error) {
    return a.client.StopQueryExecution(input)
}

func (a *AthenaActivities) TagResource(input *athena.TagResourceInput) (*athena.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *AthenaActivities) UntagResource(input *athena.UntagResourceInput) (*athena.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *AthenaActivities) UpdateDataCatalog(input *athena.UpdateDataCatalogInput) (*athena.UpdateDataCatalogOutput, error) {
    return a.client.UpdateDataCatalog(input)
}

func (a *AthenaActivities) UpdateWorkGroup(input *athena.UpdateWorkGroupInput) (*athena.UpdateWorkGroupOutput, error) {
    return a.client.UpdateWorkGroup(input)
}
