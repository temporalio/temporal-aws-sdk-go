// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package athenastub

import (
	"github.com/aws/aws-sdk-go/service/athena"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type AthenaBatchGetNamedQueryFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaBatchGetNamedQueryFuture) Get(ctx workflow.Context) (*athena.BatchGetNamedQueryOutput, error) {
	var output athena.BatchGetNamedQueryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaBatchGetQueryExecutionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaBatchGetQueryExecutionFuture) Get(ctx workflow.Context) (*athena.BatchGetQueryExecutionOutput, error) {
	var output athena.BatchGetQueryExecutionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaCreateDataCatalogFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaCreateDataCatalogFuture) Get(ctx workflow.Context) (*athena.CreateDataCatalogOutput, error) {
	var output athena.CreateDataCatalogOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaCreateNamedQueryFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaCreateNamedQueryFuture) Get(ctx workflow.Context) (*athena.CreateNamedQueryOutput, error) {
	var output athena.CreateNamedQueryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaCreateWorkGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaCreateWorkGroupFuture) Get(ctx workflow.Context) (*athena.CreateWorkGroupOutput, error) {
	var output athena.CreateWorkGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaDeleteDataCatalogFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaDeleteDataCatalogFuture) Get(ctx workflow.Context) (*athena.DeleteDataCatalogOutput, error) {
	var output athena.DeleteDataCatalogOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaDeleteNamedQueryFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaDeleteNamedQueryFuture) Get(ctx workflow.Context) (*athena.DeleteNamedQueryOutput, error) {
	var output athena.DeleteNamedQueryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaDeleteWorkGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaDeleteWorkGroupFuture) Get(ctx workflow.Context) (*athena.DeleteWorkGroupOutput, error) {
	var output athena.DeleteWorkGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaGetDataCatalogFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaGetDataCatalogFuture) Get(ctx workflow.Context) (*athena.GetDataCatalogOutput, error) {
	var output athena.GetDataCatalogOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaGetDatabaseFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaGetDatabaseFuture) Get(ctx workflow.Context) (*athena.GetDatabaseOutput, error) {
	var output athena.GetDatabaseOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaGetNamedQueryFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaGetNamedQueryFuture) Get(ctx workflow.Context) (*athena.GetNamedQueryOutput, error) {
	var output athena.GetNamedQueryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaGetQueryExecutionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaGetQueryExecutionFuture) Get(ctx workflow.Context) (*athena.GetQueryExecutionOutput, error) {
	var output athena.GetQueryExecutionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaGetQueryResultsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaGetQueryResultsFuture) Get(ctx workflow.Context) (*athena.GetQueryResultsOutput, error) {
	var output athena.GetQueryResultsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaGetTableMetadataFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaGetTableMetadataFuture) Get(ctx workflow.Context) (*athena.GetTableMetadataOutput, error) {
	var output athena.GetTableMetadataOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaGetWorkGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaGetWorkGroupFuture) Get(ctx workflow.Context) (*athena.GetWorkGroupOutput, error) {
	var output athena.GetWorkGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaListDataCatalogsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaListDataCatalogsFuture) Get(ctx workflow.Context) (*athena.ListDataCatalogsOutput, error) {
	var output athena.ListDataCatalogsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaListDatabasesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaListDatabasesFuture) Get(ctx workflow.Context) (*athena.ListDatabasesOutput, error) {
	var output athena.ListDatabasesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaListNamedQueriesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaListNamedQueriesFuture) Get(ctx workflow.Context) (*athena.ListNamedQueriesOutput, error) {
	var output athena.ListNamedQueriesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaListQueryExecutionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaListQueryExecutionsFuture) Get(ctx workflow.Context) (*athena.ListQueryExecutionsOutput, error) {
	var output athena.ListQueryExecutionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaListTableMetadataFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaListTableMetadataFuture) Get(ctx workflow.Context) (*athena.ListTableMetadataOutput, error) {
	var output athena.ListTableMetadataOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaListTagsForResourceFuture) Get(ctx workflow.Context) (*athena.ListTagsForResourceOutput, error) {
	var output athena.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaListWorkGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaListWorkGroupsFuture) Get(ctx workflow.Context) (*athena.ListWorkGroupsOutput, error) {
	var output athena.ListWorkGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaStartQueryExecutionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaStartQueryExecutionFuture) Get(ctx workflow.Context) (*athena.StartQueryExecutionOutput, error) {
	var output athena.StartQueryExecutionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaStopQueryExecutionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaStopQueryExecutionFuture) Get(ctx workflow.Context) (*athena.StopQueryExecutionOutput, error) {
	var output athena.StopQueryExecutionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaTagResourceFuture) Get(ctx workflow.Context) (*athena.TagResourceOutput, error) {
	var output athena.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaUntagResourceFuture) Get(ctx workflow.Context) (*athena.UntagResourceOutput, error) {
	var output athena.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaUpdateDataCatalogFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaUpdateDataCatalogFuture) Get(ctx workflow.Context) (*athena.UpdateDataCatalogOutput, error) {
	var output athena.UpdateDataCatalogOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AthenaUpdateWorkGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AthenaUpdateWorkGroupFuture) Get(ctx workflow.Context) (*athena.UpdateWorkGroupOutput, error) {
	var output athena.UpdateWorkGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchGetNamedQuery(ctx workflow.Context, input *athena.BatchGetNamedQueryInput) (*athena.BatchGetNamedQueryOutput, error) {
	var output athena.BatchGetNamedQueryOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.BatchGetNamedQuery", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchGetNamedQueryAsync(ctx workflow.Context, input *athena.BatchGetNamedQueryInput) *AthenaBatchGetNamedQueryFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.BatchGetNamedQuery", input)
	return &AthenaBatchGetNamedQueryFuture{Future: future}
}

func (a *stub) BatchGetQueryExecution(ctx workflow.Context, input *athena.BatchGetQueryExecutionInput) (*athena.BatchGetQueryExecutionOutput, error) {
	var output athena.BatchGetQueryExecutionOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.BatchGetQueryExecution", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchGetQueryExecutionAsync(ctx workflow.Context, input *athena.BatchGetQueryExecutionInput) *AthenaBatchGetQueryExecutionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.BatchGetQueryExecution", input)
	return &AthenaBatchGetQueryExecutionFuture{Future: future}
}

func (a *stub) CreateDataCatalog(ctx workflow.Context, input *athena.CreateDataCatalogInput) (*athena.CreateDataCatalogOutput, error) {
	var output athena.CreateDataCatalogOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.CreateDataCatalog", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDataCatalogAsync(ctx workflow.Context, input *athena.CreateDataCatalogInput) *AthenaCreateDataCatalogFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.CreateDataCatalog", input)
	return &AthenaCreateDataCatalogFuture{Future: future}
}

func (a *stub) CreateNamedQuery(ctx workflow.Context, input *athena.CreateNamedQueryInput) (*athena.CreateNamedQueryOutput, error) {
	var output athena.CreateNamedQueryOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.CreateNamedQuery", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateNamedQueryAsync(ctx workflow.Context, input *athena.CreateNamedQueryInput) *AthenaCreateNamedQueryFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.CreateNamedQuery", input)
	return &AthenaCreateNamedQueryFuture{Future: future}
}

func (a *stub) CreateWorkGroup(ctx workflow.Context, input *athena.CreateWorkGroupInput) (*athena.CreateWorkGroupOutput, error) {
	var output athena.CreateWorkGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.CreateWorkGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateWorkGroupAsync(ctx workflow.Context, input *athena.CreateWorkGroupInput) *AthenaCreateWorkGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.CreateWorkGroup", input)
	return &AthenaCreateWorkGroupFuture{Future: future}
}

func (a *stub) DeleteDataCatalog(ctx workflow.Context, input *athena.DeleteDataCatalogInput) (*athena.DeleteDataCatalogOutput, error) {
	var output athena.DeleteDataCatalogOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.DeleteDataCatalog", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDataCatalogAsync(ctx workflow.Context, input *athena.DeleteDataCatalogInput) *AthenaDeleteDataCatalogFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.DeleteDataCatalog", input)
	return &AthenaDeleteDataCatalogFuture{Future: future}
}

func (a *stub) DeleteNamedQuery(ctx workflow.Context, input *athena.DeleteNamedQueryInput) (*athena.DeleteNamedQueryOutput, error) {
	var output athena.DeleteNamedQueryOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.DeleteNamedQuery", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteNamedQueryAsync(ctx workflow.Context, input *athena.DeleteNamedQueryInput) *AthenaDeleteNamedQueryFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.DeleteNamedQuery", input)
	return &AthenaDeleteNamedQueryFuture{Future: future}
}

func (a *stub) DeleteWorkGroup(ctx workflow.Context, input *athena.DeleteWorkGroupInput) (*athena.DeleteWorkGroupOutput, error) {
	var output athena.DeleteWorkGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.DeleteWorkGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteWorkGroupAsync(ctx workflow.Context, input *athena.DeleteWorkGroupInput) *AthenaDeleteWorkGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.DeleteWorkGroup", input)
	return &AthenaDeleteWorkGroupFuture{Future: future}
}

func (a *stub) GetDataCatalog(ctx workflow.Context, input *athena.GetDataCatalogInput) (*athena.GetDataCatalogOutput, error) {
	var output athena.GetDataCatalogOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.GetDataCatalog", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDataCatalogAsync(ctx workflow.Context, input *athena.GetDataCatalogInput) *AthenaGetDataCatalogFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.GetDataCatalog", input)
	return &AthenaGetDataCatalogFuture{Future: future}
}

func (a *stub) GetDatabase(ctx workflow.Context, input *athena.GetDatabaseInput) (*athena.GetDatabaseOutput, error) {
	var output athena.GetDatabaseOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.GetDatabase", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDatabaseAsync(ctx workflow.Context, input *athena.GetDatabaseInput) *AthenaGetDatabaseFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.GetDatabase", input)
	return &AthenaGetDatabaseFuture{Future: future}
}

func (a *stub) GetNamedQuery(ctx workflow.Context, input *athena.GetNamedQueryInput) (*athena.GetNamedQueryOutput, error) {
	var output athena.GetNamedQueryOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.GetNamedQuery", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetNamedQueryAsync(ctx workflow.Context, input *athena.GetNamedQueryInput) *AthenaGetNamedQueryFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.GetNamedQuery", input)
	return &AthenaGetNamedQueryFuture{Future: future}
}

func (a *stub) GetQueryExecution(ctx workflow.Context, input *athena.GetQueryExecutionInput) (*athena.GetQueryExecutionOutput, error) {
	var output athena.GetQueryExecutionOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.GetQueryExecution", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetQueryExecutionAsync(ctx workflow.Context, input *athena.GetQueryExecutionInput) *AthenaGetQueryExecutionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.GetQueryExecution", input)
	return &AthenaGetQueryExecutionFuture{Future: future}
}

func (a *stub) GetQueryResults(ctx workflow.Context, input *athena.GetQueryResultsInput) (*athena.GetQueryResultsOutput, error) {
	var output athena.GetQueryResultsOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.GetQueryResults", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetQueryResultsAsync(ctx workflow.Context, input *athena.GetQueryResultsInput) *AthenaGetQueryResultsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.GetQueryResults", input)
	return &AthenaGetQueryResultsFuture{Future: future}
}

func (a *stub) GetTableMetadata(ctx workflow.Context, input *athena.GetTableMetadataInput) (*athena.GetTableMetadataOutput, error) {
	var output athena.GetTableMetadataOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.GetTableMetadata", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetTableMetadataAsync(ctx workflow.Context, input *athena.GetTableMetadataInput) *AthenaGetTableMetadataFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.GetTableMetadata", input)
	return &AthenaGetTableMetadataFuture{Future: future}
}

func (a *stub) GetWorkGroup(ctx workflow.Context, input *athena.GetWorkGroupInput) (*athena.GetWorkGroupOutput, error) {
	var output athena.GetWorkGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.GetWorkGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetWorkGroupAsync(ctx workflow.Context, input *athena.GetWorkGroupInput) *AthenaGetWorkGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.GetWorkGroup", input)
	return &AthenaGetWorkGroupFuture{Future: future}
}

func (a *stub) ListDataCatalogs(ctx workflow.Context, input *athena.ListDataCatalogsInput) (*athena.ListDataCatalogsOutput, error) {
	var output athena.ListDataCatalogsOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.ListDataCatalogs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDataCatalogsAsync(ctx workflow.Context, input *athena.ListDataCatalogsInput) *AthenaListDataCatalogsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.ListDataCatalogs", input)
	return &AthenaListDataCatalogsFuture{Future: future}
}

func (a *stub) ListDatabases(ctx workflow.Context, input *athena.ListDatabasesInput) (*athena.ListDatabasesOutput, error) {
	var output athena.ListDatabasesOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.ListDatabases", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDatabasesAsync(ctx workflow.Context, input *athena.ListDatabasesInput) *AthenaListDatabasesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.ListDatabases", input)
	return &AthenaListDatabasesFuture{Future: future}
}

func (a *stub) ListNamedQueries(ctx workflow.Context, input *athena.ListNamedQueriesInput) (*athena.ListNamedQueriesOutput, error) {
	var output athena.ListNamedQueriesOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.ListNamedQueries", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListNamedQueriesAsync(ctx workflow.Context, input *athena.ListNamedQueriesInput) *AthenaListNamedQueriesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.ListNamedQueries", input)
	return &AthenaListNamedQueriesFuture{Future: future}
}

func (a *stub) ListQueryExecutions(ctx workflow.Context, input *athena.ListQueryExecutionsInput) (*athena.ListQueryExecutionsOutput, error) {
	var output athena.ListQueryExecutionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.ListQueryExecutions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListQueryExecutionsAsync(ctx workflow.Context, input *athena.ListQueryExecutionsInput) *AthenaListQueryExecutionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.ListQueryExecutions", input)
	return &AthenaListQueryExecutionsFuture{Future: future}
}

func (a *stub) ListTableMetadata(ctx workflow.Context, input *athena.ListTableMetadataInput) (*athena.ListTableMetadataOutput, error) {
	var output athena.ListTableMetadataOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.ListTableMetadata", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTableMetadataAsync(ctx workflow.Context, input *athena.ListTableMetadataInput) *AthenaListTableMetadataFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.ListTableMetadata", input)
	return &AthenaListTableMetadataFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *athena.ListTagsForResourceInput) (*athena.ListTagsForResourceOutput, error) {
	var output athena.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *athena.ListTagsForResourceInput) *AthenaListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.ListTagsForResource", input)
	return &AthenaListTagsForResourceFuture{Future: future}
}

func (a *stub) ListWorkGroups(ctx workflow.Context, input *athena.ListWorkGroupsInput) (*athena.ListWorkGroupsOutput, error) {
	var output athena.ListWorkGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.ListWorkGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListWorkGroupsAsync(ctx workflow.Context, input *athena.ListWorkGroupsInput) *AthenaListWorkGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.ListWorkGroups", input)
	return &AthenaListWorkGroupsFuture{Future: future}
}

func (a *stub) StartQueryExecution(ctx workflow.Context, input *athena.StartQueryExecutionInput) (*athena.StartQueryExecutionOutput, error) {
	var output athena.StartQueryExecutionOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.StartQueryExecution", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartQueryExecutionAsync(ctx workflow.Context, input *athena.StartQueryExecutionInput) *AthenaStartQueryExecutionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.StartQueryExecution", input)
	return &AthenaStartQueryExecutionFuture{Future: future}
}

func (a *stub) StopQueryExecution(ctx workflow.Context, input *athena.StopQueryExecutionInput) (*athena.StopQueryExecutionOutput, error) {
	var output athena.StopQueryExecutionOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.StopQueryExecution", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopQueryExecutionAsync(ctx workflow.Context, input *athena.StopQueryExecutionInput) *AthenaStopQueryExecutionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.StopQueryExecution", input)
	return &AthenaStopQueryExecutionFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *athena.TagResourceInput) (*athena.TagResourceOutput, error) {
	var output athena.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *athena.TagResourceInput) *AthenaTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.TagResource", input)
	return &AthenaTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *athena.UntagResourceInput) (*athena.UntagResourceOutput, error) {
	var output athena.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *athena.UntagResourceInput) *AthenaUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.UntagResource", input)
	return &AthenaUntagResourceFuture{Future: future}
}

func (a *stub) UpdateDataCatalog(ctx workflow.Context, input *athena.UpdateDataCatalogInput) (*athena.UpdateDataCatalogOutput, error) {
	var output athena.UpdateDataCatalogOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.UpdateDataCatalog", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateDataCatalogAsync(ctx workflow.Context, input *athena.UpdateDataCatalogInput) *AthenaUpdateDataCatalogFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.UpdateDataCatalog", input)
	return &AthenaUpdateDataCatalogFuture{Future: future}
}

func (a *stub) UpdateWorkGroup(ctx workflow.Context, input *athena.UpdateWorkGroupInput) (*athena.UpdateWorkGroupOutput, error) {
	var output athena.UpdateWorkGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.athena.UpdateWorkGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateWorkGroupAsync(ctx workflow.Context, input *athena.UpdateWorkGroupInput) *AthenaUpdateWorkGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.athena.UpdateWorkGroup", input)
	return &AthenaUpdateWorkGroupFuture{Future: future}
}
