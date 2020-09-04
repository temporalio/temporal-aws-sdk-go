package awsclients

import (
	"github.com/aws/aws-sdk-go/service/athena"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type AthenaClient interface {
    BatchGetNamedQuery(ctx workflow.Context, input *athena.BatchGetNamedQueryInput) (*athena.BatchGetNamedQueryOutput, error)
    BatchGetNamedQueryAsync(ctx workflow.Context, input *athena.BatchGetNamedQueryInput) *AthenaBatchGetNamedQueryResult

    BatchGetQueryExecution(ctx workflow.Context, input *athena.BatchGetQueryExecutionInput) (*athena.BatchGetQueryExecutionOutput, error)
    BatchGetQueryExecutionAsync(ctx workflow.Context, input *athena.BatchGetQueryExecutionInput) *AthenaBatchGetQueryExecutionResult

    CreateDataCatalog(ctx workflow.Context, input *athena.CreateDataCatalogInput) (*athena.CreateDataCatalogOutput, error)
    CreateDataCatalogAsync(ctx workflow.Context, input *athena.CreateDataCatalogInput) *AthenaCreateDataCatalogResult

    CreateNamedQuery(ctx workflow.Context, input *athena.CreateNamedQueryInput) (*athena.CreateNamedQueryOutput, error)
    CreateNamedQueryAsync(ctx workflow.Context, input *athena.CreateNamedQueryInput) *AthenaCreateNamedQueryResult

    CreateWorkGroup(ctx workflow.Context, input *athena.CreateWorkGroupInput) (*athena.CreateWorkGroupOutput, error)
    CreateWorkGroupAsync(ctx workflow.Context, input *athena.CreateWorkGroupInput) *AthenaCreateWorkGroupResult

    DeleteDataCatalog(ctx workflow.Context, input *athena.DeleteDataCatalogInput) (*athena.DeleteDataCatalogOutput, error)
    DeleteDataCatalogAsync(ctx workflow.Context, input *athena.DeleteDataCatalogInput) *AthenaDeleteDataCatalogResult

    DeleteNamedQuery(ctx workflow.Context, input *athena.DeleteNamedQueryInput) (*athena.DeleteNamedQueryOutput, error)
    DeleteNamedQueryAsync(ctx workflow.Context, input *athena.DeleteNamedQueryInput) *AthenaDeleteNamedQueryResult

    DeleteWorkGroup(ctx workflow.Context, input *athena.DeleteWorkGroupInput) (*athena.DeleteWorkGroupOutput, error)
    DeleteWorkGroupAsync(ctx workflow.Context, input *athena.DeleteWorkGroupInput) *AthenaDeleteWorkGroupResult

    GetDataCatalog(ctx workflow.Context, input *athena.GetDataCatalogInput) (*athena.GetDataCatalogOutput, error)
    GetDataCatalogAsync(ctx workflow.Context, input *athena.GetDataCatalogInput) *AthenaGetDataCatalogResult

    GetDatabase(ctx workflow.Context, input *athena.GetDatabaseInput) (*athena.GetDatabaseOutput, error)
    GetDatabaseAsync(ctx workflow.Context, input *athena.GetDatabaseInput) *AthenaGetDatabaseResult

    GetNamedQuery(ctx workflow.Context, input *athena.GetNamedQueryInput) (*athena.GetNamedQueryOutput, error)
    GetNamedQueryAsync(ctx workflow.Context, input *athena.GetNamedQueryInput) *AthenaGetNamedQueryResult

    GetQueryExecution(ctx workflow.Context, input *athena.GetQueryExecutionInput) (*athena.GetQueryExecutionOutput, error)
    GetQueryExecutionAsync(ctx workflow.Context, input *athena.GetQueryExecutionInput) *AthenaGetQueryExecutionResult

    GetQueryResults(ctx workflow.Context, input *athena.GetQueryResultsInput) (*athena.GetQueryResultsOutput, error)
    GetQueryResultsAsync(ctx workflow.Context, input *athena.GetQueryResultsInput) *AthenaGetQueryResultsResult

    GetTableMetadata(ctx workflow.Context, input *athena.GetTableMetadataInput) (*athena.GetTableMetadataOutput, error)
    GetTableMetadataAsync(ctx workflow.Context, input *athena.GetTableMetadataInput) *AthenaGetTableMetadataResult

    GetWorkGroup(ctx workflow.Context, input *athena.GetWorkGroupInput) (*athena.GetWorkGroupOutput, error)
    GetWorkGroupAsync(ctx workflow.Context, input *athena.GetWorkGroupInput) *AthenaGetWorkGroupResult

    ListDataCatalogs(ctx workflow.Context, input *athena.ListDataCatalogsInput) (*athena.ListDataCatalogsOutput, error)
    ListDataCatalogsAsync(ctx workflow.Context, input *athena.ListDataCatalogsInput) *AthenaListDataCatalogsResult

    ListDatabases(ctx workflow.Context, input *athena.ListDatabasesInput) (*athena.ListDatabasesOutput, error)
    ListDatabasesAsync(ctx workflow.Context, input *athena.ListDatabasesInput) *AthenaListDatabasesResult

    ListNamedQueries(ctx workflow.Context, input *athena.ListNamedQueriesInput) (*athena.ListNamedQueriesOutput, error)
    ListNamedQueriesAsync(ctx workflow.Context, input *athena.ListNamedQueriesInput) *AthenaListNamedQueriesResult

    ListQueryExecutions(ctx workflow.Context, input *athena.ListQueryExecutionsInput) (*athena.ListQueryExecutionsOutput, error)
    ListQueryExecutionsAsync(ctx workflow.Context, input *athena.ListQueryExecutionsInput) *AthenaListQueryExecutionsResult

    ListTableMetadata(ctx workflow.Context, input *athena.ListTableMetadataInput) (*athena.ListTableMetadataOutput, error)
    ListTableMetadataAsync(ctx workflow.Context, input *athena.ListTableMetadataInput) *AthenaListTableMetadataResult

    ListTagsForResource(ctx workflow.Context, input *athena.ListTagsForResourceInput) (*athena.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *athena.ListTagsForResourceInput) *AthenaListTagsForResourceResult

    ListWorkGroups(ctx workflow.Context, input *athena.ListWorkGroupsInput) (*athena.ListWorkGroupsOutput, error)
    ListWorkGroupsAsync(ctx workflow.Context, input *athena.ListWorkGroupsInput) *AthenaListWorkGroupsResult

    StartQueryExecution(ctx workflow.Context, input *athena.StartQueryExecutionInput) (*athena.StartQueryExecutionOutput, error)
    StartQueryExecutionAsync(ctx workflow.Context, input *athena.StartQueryExecutionInput) *AthenaStartQueryExecutionResult

    StopQueryExecution(ctx workflow.Context, input *athena.StopQueryExecutionInput) (*athena.StopQueryExecutionOutput, error)
    StopQueryExecutionAsync(ctx workflow.Context, input *athena.StopQueryExecutionInput) *AthenaStopQueryExecutionResult

    TagResource(ctx workflow.Context, input *athena.TagResourceInput) (*athena.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *athena.TagResourceInput) *AthenaTagResourceResult

    UntagResource(ctx workflow.Context, input *athena.UntagResourceInput) (*athena.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *athena.UntagResourceInput) *AthenaUntagResourceResult

    UpdateDataCatalog(ctx workflow.Context, input *athena.UpdateDataCatalogInput) (*athena.UpdateDataCatalogOutput, error)
    UpdateDataCatalogAsync(ctx workflow.Context, input *athena.UpdateDataCatalogInput) *AthenaUpdateDataCatalogResult

    UpdateWorkGroup(ctx workflow.Context, input *athena.UpdateWorkGroupInput) (*athena.UpdateWorkGroupOutput, error)
    UpdateWorkGroupAsync(ctx workflow.Context, input *athena.UpdateWorkGroupInput) *AthenaUpdateWorkGroupResult
}

type AthenaBatchGetNamedQueryResult struct {
	Result workflow.Future
}

func (r *AthenaBatchGetNamedQueryResult) Get(ctx workflow.Context) (*athena.BatchGetNamedQueryOutput, error) {
    var output athena.BatchGetNamedQueryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaBatchGetQueryExecutionResult struct {
	Result workflow.Future
}

func (r *AthenaBatchGetQueryExecutionResult) Get(ctx workflow.Context) (*athena.BatchGetQueryExecutionOutput, error) {
    var output athena.BatchGetQueryExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaCreateDataCatalogResult struct {
	Result workflow.Future
}

func (r *AthenaCreateDataCatalogResult) Get(ctx workflow.Context) (*athena.CreateDataCatalogOutput, error) {
    var output athena.CreateDataCatalogOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaCreateNamedQueryResult struct {
	Result workflow.Future
}

func (r *AthenaCreateNamedQueryResult) Get(ctx workflow.Context) (*athena.CreateNamedQueryOutput, error) {
    var output athena.CreateNamedQueryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaCreateWorkGroupResult struct {
	Result workflow.Future
}

func (r *AthenaCreateWorkGroupResult) Get(ctx workflow.Context) (*athena.CreateWorkGroupOutput, error) {
    var output athena.CreateWorkGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaDeleteDataCatalogResult struct {
	Result workflow.Future
}

func (r *AthenaDeleteDataCatalogResult) Get(ctx workflow.Context) (*athena.DeleteDataCatalogOutput, error) {
    var output athena.DeleteDataCatalogOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaDeleteNamedQueryResult struct {
	Result workflow.Future
}

func (r *AthenaDeleteNamedQueryResult) Get(ctx workflow.Context) (*athena.DeleteNamedQueryOutput, error) {
    var output athena.DeleteNamedQueryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaDeleteWorkGroupResult struct {
	Result workflow.Future
}

func (r *AthenaDeleteWorkGroupResult) Get(ctx workflow.Context) (*athena.DeleteWorkGroupOutput, error) {
    var output athena.DeleteWorkGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaGetDataCatalogResult struct {
	Result workflow.Future
}

func (r *AthenaGetDataCatalogResult) Get(ctx workflow.Context) (*athena.GetDataCatalogOutput, error) {
    var output athena.GetDataCatalogOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaGetDatabaseResult struct {
	Result workflow.Future
}

func (r *AthenaGetDatabaseResult) Get(ctx workflow.Context) (*athena.GetDatabaseOutput, error) {
    var output athena.GetDatabaseOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaGetNamedQueryResult struct {
	Result workflow.Future
}

func (r *AthenaGetNamedQueryResult) Get(ctx workflow.Context) (*athena.GetNamedQueryOutput, error) {
    var output athena.GetNamedQueryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaGetQueryExecutionResult struct {
	Result workflow.Future
}

func (r *AthenaGetQueryExecutionResult) Get(ctx workflow.Context) (*athena.GetQueryExecutionOutput, error) {
    var output athena.GetQueryExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaGetQueryResultsResult struct {
	Result workflow.Future
}

func (r *AthenaGetQueryResultsResult) Get(ctx workflow.Context) (*athena.GetQueryResultsOutput, error) {
    var output athena.GetQueryResultsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaGetTableMetadataResult struct {
	Result workflow.Future
}

func (r *AthenaGetTableMetadataResult) Get(ctx workflow.Context) (*athena.GetTableMetadataOutput, error) {
    var output athena.GetTableMetadataOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaGetWorkGroupResult struct {
	Result workflow.Future
}

func (r *AthenaGetWorkGroupResult) Get(ctx workflow.Context) (*athena.GetWorkGroupOutput, error) {
    var output athena.GetWorkGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaListDataCatalogsResult struct {
	Result workflow.Future
}

func (r *AthenaListDataCatalogsResult) Get(ctx workflow.Context) (*athena.ListDataCatalogsOutput, error) {
    var output athena.ListDataCatalogsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaListDatabasesResult struct {
	Result workflow.Future
}

func (r *AthenaListDatabasesResult) Get(ctx workflow.Context) (*athena.ListDatabasesOutput, error) {
    var output athena.ListDatabasesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaListNamedQueriesResult struct {
	Result workflow.Future
}

func (r *AthenaListNamedQueriesResult) Get(ctx workflow.Context) (*athena.ListNamedQueriesOutput, error) {
    var output athena.ListNamedQueriesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaListQueryExecutionsResult struct {
	Result workflow.Future
}

func (r *AthenaListQueryExecutionsResult) Get(ctx workflow.Context) (*athena.ListQueryExecutionsOutput, error) {
    var output athena.ListQueryExecutionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaListTableMetadataResult struct {
	Result workflow.Future
}

func (r *AthenaListTableMetadataResult) Get(ctx workflow.Context) (*athena.ListTableMetadataOutput, error) {
    var output athena.ListTableMetadataOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *AthenaListTagsForResourceResult) Get(ctx workflow.Context) (*athena.ListTagsForResourceOutput, error) {
    var output athena.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaListWorkGroupsResult struct {
	Result workflow.Future
}

func (r *AthenaListWorkGroupsResult) Get(ctx workflow.Context) (*athena.ListWorkGroupsOutput, error) {
    var output athena.ListWorkGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaStartQueryExecutionResult struct {
	Result workflow.Future
}

func (r *AthenaStartQueryExecutionResult) Get(ctx workflow.Context) (*athena.StartQueryExecutionOutput, error) {
    var output athena.StartQueryExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaStopQueryExecutionResult struct {
	Result workflow.Future
}

func (r *AthenaStopQueryExecutionResult) Get(ctx workflow.Context) (*athena.StopQueryExecutionOutput, error) {
    var output athena.StopQueryExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaTagResourceResult struct {
	Result workflow.Future
}

func (r *AthenaTagResourceResult) Get(ctx workflow.Context) (*athena.TagResourceOutput, error) {
    var output athena.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaUntagResourceResult struct {
	Result workflow.Future
}

func (r *AthenaUntagResourceResult) Get(ctx workflow.Context) (*athena.UntagResourceOutput, error) {
    var output athena.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaUpdateDataCatalogResult struct {
	Result workflow.Future
}

func (r *AthenaUpdateDataCatalogResult) Get(ctx workflow.Context) (*athena.UpdateDataCatalogOutput, error) {
    var output athena.UpdateDataCatalogOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaUpdateWorkGroupResult struct {
	Result workflow.Future
}

func (r *AthenaUpdateWorkGroupResult) Get(ctx workflow.Context) (*athena.UpdateWorkGroupOutput, error) {
    var output athena.UpdateWorkGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AthenaStub struct {
    activities awsactivities.AthenaActivities
}

func NewAthenaStub() AthenaClient {
    return &AthenaStub{}
}

func (a *AthenaStub) BatchGetNamedQuery(ctx workflow.Context, input *athena.BatchGetNamedQueryInput) (*athena.BatchGetNamedQueryOutput, error) {
    var output athena.BatchGetNamedQueryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetNamedQuery, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) BatchGetNamedQueryAsync(ctx workflow.Context, input *athena.BatchGetNamedQueryInput) *AthenaBatchGetNamedQueryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchGetNamedQuery, input)
    return &AthenaBatchGetNamedQueryResult{Result: future}
}

func (a *AthenaStub) BatchGetQueryExecution(ctx workflow.Context, input *athena.BatchGetQueryExecutionInput) (*athena.BatchGetQueryExecutionOutput, error) {
    var output athena.BatchGetQueryExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetQueryExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) BatchGetQueryExecutionAsync(ctx workflow.Context, input *athena.BatchGetQueryExecutionInput) *AthenaBatchGetQueryExecutionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchGetQueryExecution, input)
    return &AthenaBatchGetQueryExecutionResult{Result: future}
}

func (a *AthenaStub) CreateDataCatalog(ctx workflow.Context, input *athena.CreateDataCatalogInput) (*athena.CreateDataCatalogOutput, error) {
    var output athena.CreateDataCatalogOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDataCatalog, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) CreateDataCatalogAsync(ctx workflow.Context, input *athena.CreateDataCatalogInput) *AthenaCreateDataCatalogResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDataCatalog, input)
    return &AthenaCreateDataCatalogResult{Result: future}
}

func (a *AthenaStub) CreateNamedQuery(ctx workflow.Context, input *athena.CreateNamedQueryInput) (*athena.CreateNamedQueryOutput, error) {
    var output athena.CreateNamedQueryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateNamedQuery, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) CreateNamedQueryAsync(ctx workflow.Context, input *athena.CreateNamedQueryInput) *AthenaCreateNamedQueryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateNamedQuery, input)
    return &AthenaCreateNamedQueryResult{Result: future}
}

func (a *AthenaStub) CreateWorkGroup(ctx workflow.Context, input *athena.CreateWorkGroupInput) (*athena.CreateWorkGroupOutput, error) {
    var output athena.CreateWorkGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateWorkGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) CreateWorkGroupAsync(ctx workflow.Context, input *athena.CreateWorkGroupInput) *AthenaCreateWorkGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateWorkGroup, input)
    return &AthenaCreateWorkGroupResult{Result: future}
}

func (a *AthenaStub) DeleteDataCatalog(ctx workflow.Context, input *athena.DeleteDataCatalogInput) (*athena.DeleteDataCatalogOutput, error) {
    var output athena.DeleteDataCatalogOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDataCatalog, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) DeleteDataCatalogAsync(ctx workflow.Context, input *athena.DeleteDataCatalogInput) *AthenaDeleteDataCatalogResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDataCatalog, input)
    return &AthenaDeleteDataCatalogResult{Result: future}
}

func (a *AthenaStub) DeleteNamedQuery(ctx workflow.Context, input *athena.DeleteNamedQueryInput) (*athena.DeleteNamedQueryOutput, error) {
    var output athena.DeleteNamedQueryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteNamedQuery, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) DeleteNamedQueryAsync(ctx workflow.Context, input *athena.DeleteNamedQueryInput) *AthenaDeleteNamedQueryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteNamedQuery, input)
    return &AthenaDeleteNamedQueryResult{Result: future}
}

func (a *AthenaStub) DeleteWorkGroup(ctx workflow.Context, input *athena.DeleteWorkGroupInput) (*athena.DeleteWorkGroupOutput, error) {
    var output athena.DeleteWorkGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteWorkGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) DeleteWorkGroupAsync(ctx workflow.Context, input *athena.DeleteWorkGroupInput) *AthenaDeleteWorkGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteWorkGroup, input)
    return &AthenaDeleteWorkGroupResult{Result: future}
}

func (a *AthenaStub) GetDataCatalog(ctx workflow.Context, input *athena.GetDataCatalogInput) (*athena.GetDataCatalogOutput, error) {
    var output athena.GetDataCatalogOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDataCatalog, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) GetDataCatalogAsync(ctx workflow.Context, input *athena.GetDataCatalogInput) *AthenaGetDataCatalogResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDataCatalog, input)
    return &AthenaGetDataCatalogResult{Result: future}
}

func (a *AthenaStub) GetDatabase(ctx workflow.Context, input *athena.GetDatabaseInput) (*athena.GetDatabaseOutput, error) {
    var output athena.GetDatabaseOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDatabase, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) GetDatabaseAsync(ctx workflow.Context, input *athena.GetDatabaseInput) *AthenaGetDatabaseResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDatabase, input)
    return &AthenaGetDatabaseResult{Result: future}
}

func (a *AthenaStub) GetNamedQuery(ctx workflow.Context, input *athena.GetNamedQueryInput) (*athena.GetNamedQueryOutput, error) {
    var output athena.GetNamedQueryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetNamedQuery, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) GetNamedQueryAsync(ctx workflow.Context, input *athena.GetNamedQueryInput) *AthenaGetNamedQueryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetNamedQuery, input)
    return &AthenaGetNamedQueryResult{Result: future}
}

func (a *AthenaStub) GetQueryExecution(ctx workflow.Context, input *athena.GetQueryExecutionInput) (*athena.GetQueryExecutionOutput, error) {
    var output athena.GetQueryExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetQueryExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) GetQueryExecutionAsync(ctx workflow.Context, input *athena.GetQueryExecutionInput) *AthenaGetQueryExecutionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetQueryExecution, input)
    return &AthenaGetQueryExecutionResult{Result: future}
}

func (a *AthenaStub) GetQueryResults(ctx workflow.Context, input *athena.GetQueryResultsInput) (*athena.GetQueryResultsOutput, error) {
    var output athena.GetQueryResultsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetQueryResults, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) GetQueryResultsAsync(ctx workflow.Context, input *athena.GetQueryResultsInput) *AthenaGetQueryResultsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetQueryResults, input)
    return &AthenaGetQueryResultsResult{Result: future}
}

func (a *AthenaStub) GetTableMetadata(ctx workflow.Context, input *athena.GetTableMetadataInput) (*athena.GetTableMetadataOutput, error) {
    var output athena.GetTableMetadataOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTableMetadata, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) GetTableMetadataAsync(ctx workflow.Context, input *athena.GetTableMetadataInput) *AthenaGetTableMetadataResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTableMetadata, input)
    return &AthenaGetTableMetadataResult{Result: future}
}

func (a *AthenaStub) GetWorkGroup(ctx workflow.Context, input *athena.GetWorkGroupInput) (*athena.GetWorkGroupOutput, error) {
    var output athena.GetWorkGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetWorkGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) GetWorkGroupAsync(ctx workflow.Context, input *athena.GetWorkGroupInput) *AthenaGetWorkGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetWorkGroup, input)
    return &AthenaGetWorkGroupResult{Result: future}
}

func (a *AthenaStub) ListDataCatalogs(ctx workflow.Context, input *athena.ListDataCatalogsInput) (*athena.ListDataCatalogsOutput, error) {
    var output athena.ListDataCatalogsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDataCatalogs, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) ListDataCatalogsAsync(ctx workflow.Context, input *athena.ListDataCatalogsInput) *AthenaListDataCatalogsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDataCatalogs, input)
    return &AthenaListDataCatalogsResult{Result: future}
}

func (a *AthenaStub) ListDatabases(ctx workflow.Context, input *athena.ListDatabasesInput) (*athena.ListDatabasesOutput, error) {
    var output athena.ListDatabasesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDatabases, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) ListDatabasesAsync(ctx workflow.Context, input *athena.ListDatabasesInput) *AthenaListDatabasesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDatabases, input)
    return &AthenaListDatabasesResult{Result: future}
}

func (a *AthenaStub) ListNamedQueries(ctx workflow.Context, input *athena.ListNamedQueriesInput) (*athena.ListNamedQueriesOutput, error) {
    var output athena.ListNamedQueriesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListNamedQueries, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) ListNamedQueriesAsync(ctx workflow.Context, input *athena.ListNamedQueriesInput) *AthenaListNamedQueriesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListNamedQueries, input)
    return &AthenaListNamedQueriesResult{Result: future}
}

func (a *AthenaStub) ListQueryExecutions(ctx workflow.Context, input *athena.ListQueryExecutionsInput) (*athena.ListQueryExecutionsOutput, error) {
    var output athena.ListQueryExecutionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListQueryExecutions, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) ListQueryExecutionsAsync(ctx workflow.Context, input *athena.ListQueryExecutionsInput) *AthenaListQueryExecutionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListQueryExecutions, input)
    return &AthenaListQueryExecutionsResult{Result: future}
}

func (a *AthenaStub) ListTableMetadata(ctx workflow.Context, input *athena.ListTableMetadataInput) (*athena.ListTableMetadataOutput, error) {
    var output athena.ListTableMetadataOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTableMetadata, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) ListTableMetadataAsync(ctx workflow.Context, input *athena.ListTableMetadataInput) *AthenaListTableMetadataResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTableMetadata, input)
    return &AthenaListTableMetadataResult{Result: future}
}

func (a *AthenaStub) ListTagsForResource(ctx workflow.Context, input *athena.ListTagsForResourceInput) (*athena.ListTagsForResourceOutput, error) {
    var output athena.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) ListTagsForResourceAsync(ctx workflow.Context, input *athena.ListTagsForResourceInput) *AthenaListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &AthenaListTagsForResourceResult{Result: future}
}

func (a *AthenaStub) ListWorkGroups(ctx workflow.Context, input *athena.ListWorkGroupsInput) (*athena.ListWorkGroupsOutput, error) {
    var output athena.ListWorkGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListWorkGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) ListWorkGroupsAsync(ctx workflow.Context, input *athena.ListWorkGroupsInput) *AthenaListWorkGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListWorkGroups, input)
    return &AthenaListWorkGroupsResult{Result: future}
}

func (a *AthenaStub) StartQueryExecution(ctx workflow.Context, input *athena.StartQueryExecutionInput) (*athena.StartQueryExecutionOutput, error) {
    var output athena.StartQueryExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartQueryExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) StartQueryExecutionAsync(ctx workflow.Context, input *athena.StartQueryExecutionInput) *AthenaStartQueryExecutionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartQueryExecution, input)
    return &AthenaStartQueryExecutionResult{Result: future}
}

func (a *AthenaStub) StopQueryExecution(ctx workflow.Context, input *athena.StopQueryExecutionInput) (*athena.StopQueryExecutionOutput, error) {
    var output athena.StopQueryExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopQueryExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) StopQueryExecutionAsync(ctx workflow.Context, input *athena.StopQueryExecutionInput) *AthenaStopQueryExecutionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopQueryExecution, input)
    return &AthenaStopQueryExecutionResult{Result: future}
}

func (a *AthenaStub) TagResource(ctx workflow.Context, input *athena.TagResourceInput) (*athena.TagResourceOutput, error) {
    var output athena.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) TagResourceAsync(ctx workflow.Context, input *athena.TagResourceInput) *AthenaTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &AthenaTagResourceResult{Result: future}
}

func (a *AthenaStub) UntagResource(ctx workflow.Context, input *athena.UntagResourceInput) (*athena.UntagResourceOutput, error) {
    var output athena.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) UntagResourceAsync(ctx workflow.Context, input *athena.UntagResourceInput) *AthenaUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &AthenaUntagResourceResult{Result: future}
}

func (a *AthenaStub) UpdateDataCatalog(ctx workflow.Context, input *athena.UpdateDataCatalogInput) (*athena.UpdateDataCatalogOutput, error) {
    var output athena.UpdateDataCatalogOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDataCatalog, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) UpdateDataCatalogAsync(ctx workflow.Context, input *athena.UpdateDataCatalogInput) *AthenaUpdateDataCatalogResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateDataCatalog, input)
    return &AthenaUpdateDataCatalogResult{Result: future}
}

func (a *AthenaStub) UpdateWorkGroup(ctx workflow.Context, input *athena.UpdateWorkGroupInput) (*athena.UpdateWorkGroupOutput, error) {
    var output athena.UpdateWorkGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateWorkGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *AthenaStub) UpdateWorkGroupAsync(ctx workflow.Context, input *athena.UpdateWorkGroupInput) *AthenaUpdateWorkGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateWorkGroup, input)
    return &AthenaUpdateWorkGroupResult{Result: future}
}
