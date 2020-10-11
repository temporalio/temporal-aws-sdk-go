// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package appsyncstub

import (
	"github.com/aws/aws-sdk-go/service/appsync"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	CreateApiCache(ctx workflow.Context, input *appsync.CreateApiCacheInput) (*appsync.CreateApiCacheOutput, error)
	CreateApiCacheAsync(ctx workflow.Context, input *appsync.CreateApiCacheInput) *AppSyncCreateApiCacheFuture

	CreateApiKey(ctx workflow.Context, input *appsync.CreateApiKeyInput) (*appsync.CreateApiKeyOutput, error)
	CreateApiKeyAsync(ctx workflow.Context, input *appsync.CreateApiKeyInput) *AppSyncCreateApiKeyFuture

	CreateDataSource(ctx workflow.Context, input *appsync.CreateDataSourceInput) (*appsync.CreateDataSourceOutput, error)
	CreateDataSourceAsync(ctx workflow.Context, input *appsync.CreateDataSourceInput) *AppSyncCreateDataSourceFuture

	CreateFunction(ctx workflow.Context, input *appsync.CreateFunctionInput) (*appsync.CreateFunctionOutput, error)
	CreateFunctionAsync(ctx workflow.Context, input *appsync.CreateFunctionInput) *AppSyncCreateFunctionFuture

	CreateGraphqlApi(ctx workflow.Context, input *appsync.CreateGraphqlApiInput) (*appsync.CreateGraphqlApiOutput, error)
	CreateGraphqlApiAsync(ctx workflow.Context, input *appsync.CreateGraphqlApiInput) *AppSyncCreateGraphqlApiFuture

	CreateResolver(ctx workflow.Context, input *appsync.CreateResolverInput) (*appsync.CreateResolverOutput, error)
	CreateResolverAsync(ctx workflow.Context, input *appsync.CreateResolverInput) *AppSyncCreateResolverFuture

	CreateType(ctx workflow.Context, input *appsync.CreateTypeInput) (*appsync.CreateTypeOutput, error)
	CreateTypeAsync(ctx workflow.Context, input *appsync.CreateTypeInput) *AppSyncCreateTypeFuture

	DeleteApiCache(ctx workflow.Context, input *appsync.DeleteApiCacheInput) (*appsync.DeleteApiCacheOutput, error)
	DeleteApiCacheAsync(ctx workflow.Context, input *appsync.DeleteApiCacheInput) *AppSyncDeleteApiCacheFuture

	DeleteApiKey(ctx workflow.Context, input *appsync.DeleteApiKeyInput) (*appsync.DeleteApiKeyOutput, error)
	DeleteApiKeyAsync(ctx workflow.Context, input *appsync.DeleteApiKeyInput) *AppSyncDeleteApiKeyFuture

	DeleteDataSource(ctx workflow.Context, input *appsync.DeleteDataSourceInput) (*appsync.DeleteDataSourceOutput, error)
	DeleteDataSourceAsync(ctx workflow.Context, input *appsync.DeleteDataSourceInput) *AppSyncDeleteDataSourceFuture

	DeleteFunction(ctx workflow.Context, input *appsync.DeleteFunctionInput) (*appsync.DeleteFunctionOutput, error)
	DeleteFunctionAsync(ctx workflow.Context, input *appsync.DeleteFunctionInput) *AppSyncDeleteFunctionFuture

	DeleteGraphqlApi(ctx workflow.Context, input *appsync.DeleteGraphqlApiInput) (*appsync.DeleteGraphqlApiOutput, error)
	DeleteGraphqlApiAsync(ctx workflow.Context, input *appsync.DeleteGraphqlApiInput) *AppSyncDeleteGraphqlApiFuture

	DeleteResolver(ctx workflow.Context, input *appsync.DeleteResolverInput) (*appsync.DeleteResolverOutput, error)
	DeleteResolverAsync(ctx workflow.Context, input *appsync.DeleteResolverInput) *AppSyncDeleteResolverFuture

	DeleteType(ctx workflow.Context, input *appsync.DeleteTypeInput) (*appsync.DeleteTypeOutput, error)
	DeleteTypeAsync(ctx workflow.Context, input *appsync.DeleteTypeInput) *AppSyncDeleteTypeFuture

	FlushApiCache(ctx workflow.Context, input *appsync.FlushApiCacheInput) (*appsync.FlushApiCacheOutput, error)
	FlushApiCacheAsync(ctx workflow.Context, input *appsync.FlushApiCacheInput) *AppSyncFlushApiCacheFuture

	GetApiCache(ctx workflow.Context, input *appsync.GetApiCacheInput) (*appsync.GetApiCacheOutput, error)
	GetApiCacheAsync(ctx workflow.Context, input *appsync.GetApiCacheInput) *AppSyncGetApiCacheFuture

	GetDataSource(ctx workflow.Context, input *appsync.GetDataSourceInput) (*appsync.GetDataSourceOutput, error)
	GetDataSourceAsync(ctx workflow.Context, input *appsync.GetDataSourceInput) *AppSyncGetDataSourceFuture

	GetFunction(ctx workflow.Context, input *appsync.GetFunctionInput) (*appsync.GetFunctionOutput, error)
	GetFunctionAsync(ctx workflow.Context, input *appsync.GetFunctionInput) *AppSyncGetFunctionFuture

	GetGraphqlApi(ctx workflow.Context, input *appsync.GetGraphqlApiInput) (*appsync.GetGraphqlApiOutput, error)
	GetGraphqlApiAsync(ctx workflow.Context, input *appsync.GetGraphqlApiInput) *AppSyncGetGraphqlApiFuture

	GetIntrospectionSchema(ctx workflow.Context, input *appsync.GetIntrospectionSchemaInput) (*appsync.GetIntrospectionSchemaOutput, error)
	GetIntrospectionSchemaAsync(ctx workflow.Context, input *appsync.GetIntrospectionSchemaInput) *AppSyncGetIntrospectionSchemaFuture

	GetResolver(ctx workflow.Context, input *appsync.GetResolverInput) (*appsync.GetResolverOutput, error)
	GetResolverAsync(ctx workflow.Context, input *appsync.GetResolverInput) *AppSyncGetResolverFuture

	GetSchemaCreationStatus(ctx workflow.Context, input *appsync.GetSchemaCreationStatusInput) (*appsync.GetSchemaCreationStatusOutput, error)
	GetSchemaCreationStatusAsync(ctx workflow.Context, input *appsync.GetSchemaCreationStatusInput) *AppSyncGetSchemaCreationStatusFuture

	GetType(ctx workflow.Context, input *appsync.GetTypeInput) (*appsync.GetTypeOutput, error)
	GetTypeAsync(ctx workflow.Context, input *appsync.GetTypeInput) *AppSyncGetTypeFuture

	ListApiKeys(ctx workflow.Context, input *appsync.ListApiKeysInput) (*appsync.ListApiKeysOutput, error)
	ListApiKeysAsync(ctx workflow.Context, input *appsync.ListApiKeysInput) *AppSyncListApiKeysFuture

	ListDataSources(ctx workflow.Context, input *appsync.ListDataSourcesInput) (*appsync.ListDataSourcesOutput, error)
	ListDataSourcesAsync(ctx workflow.Context, input *appsync.ListDataSourcesInput) *AppSyncListDataSourcesFuture

	ListFunctions(ctx workflow.Context, input *appsync.ListFunctionsInput) (*appsync.ListFunctionsOutput, error)
	ListFunctionsAsync(ctx workflow.Context, input *appsync.ListFunctionsInput) *AppSyncListFunctionsFuture

	ListGraphqlApis(ctx workflow.Context, input *appsync.ListGraphqlApisInput) (*appsync.ListGraphqlApisOutput, error)
	ListGraphqlApisAsync(ctx workflow.Context, input *appsync.ListGraphqlApisInput) *AppSyncListGraphqlApisFuture

	ListResolvers(ctx workflow.Context, input *appsync.ListResolversInput) (*appsync.ListResolversOutput, error)
	ListResolversAsync(ctx workflow.Context, input *appsync.ListResolversInput) *AppSyncListResolversFuture

	ListResolversByFunction(ctx workflow.Context, input *appsync.ListResolversByFunctionInput) (*appsync.ListResolversByFunctionOutput, error)
	ListResolversByFunctionAsync(ctx workflow.Context, input *appsync.ListResolversByFunctionInput) *AppSyncListResolversByFunctionFuture

	ListTagsForResource(ctx workflow.Context, input *appsync.ListTagsForResourceInput) (*appsync.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *appsync.ListTagsForResourceInput) *AppSyncListTagsForResourceFuture

	ListTypes(ctx workflow.Context, input *appsync.ListTypesInput) (*appsync.ListTypesOutput, error)
	ListTypesAsync(ctx workflow.Context, input *appsync.ListTypesInput) *AppSyncListTypesFuture

	StartSchemaCreation(ctx workflow.Context, input *appsync.StartSchemaCreationInput) (*appsync.StartSchemaCreationOutput, error)
	StartSchemaCreationAsync(ctx workflow.Context, input *appsync.StartSchemaCreationInput) *AppSyncStartSchemaCreationFuture

	TagResource(ctx workflow.Context, input *appsync.TagResourceInput) (*appsync.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *appsync.TagResourceInput) *AppSyncTagResourceFuture

	UntagResource(ctx workflow.Context, input *appsync.UntagResourceInput) (*appsync.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *appsync.UntagResourceInput) *AppSyncUntagResourceFuture

	UpdateApiCache(ctx workflow.Context, input *appsync.UpdateApiCacheInput) (*appsync.UpdateApiCacheOutput, error)
	UpdateApiCacheAsync(ctx workflow.Context, input *appsync.UpdateApiCacheInput) *AppSyncUpdateApiCacheFuture

	UpdateApiKey(ctx workflow.Context, input *appsync.UpdateApiKeyInput) (*appsync.UpdateApiKeyOutput, error)
	UpdateApiKeyAsync(ctx workflow.Context, input *appsync.UpdateApiKeyInput) *AppSyncUpdateApiKeyFuture

	UpdateDataSource(ctx workflow.Context, input *appsync.UpdateDataSourceInput) (*appsync.UpdateDataSourceOutput, error)
	UpdateDataSourceAsync(ctx workflow.Context, input *appsync.UpdateDataSourceInput) *AppSyncUpdateDataSourceFuture

	UpdateFunction(ctx workflow.Context, input *appsync.UpdateFunctionInput) (*appsync.UpdateFunctionOutput, error)
	UpdateFunctionAsync(ctx workflow.Context, input *appsync.UpdateFunctionInput) *AppSyncUpdateFunctionFuture

	UpdateGraphqlApi(ctx workflow.Context, input *appsync.UpdateGraphqlApiInput) (*appsync.UpdateGraphqlApiOutput, error)
	UpdateGraphqlApiAsync(ctx workflow.Context, input *appsync.UpdateGraphqlApiInput) *AppSyncUpdateGraphqlApiFuture

	UpdateResolver(ctx workflow.Context, input *appsync.UpdateResolverInput) (*appsync.UpdateResolverOutput, error)
	UpdateResolverAsync(ctx workflow.Context, input *appsync.UpdateResolverInput) *AppSyncUpdateResolverFuture

	UpdateType(ctx workflow.Context, input *appsync.UpdateTypeInput) (*appsync.UpdateTypeOutput, error)
	UpdateTypeAsync(ctx workflow.Context, input *appsync.UpdateTypeInput) *AppSyncUpdateTypeFuture
}

func NewClient() Client {
	return &stub{}
}
