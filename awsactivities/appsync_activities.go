package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/appsync"
	"github.com/aws/aws-sdk-go/service/appsync/appsynciface"
)

type AppSyncActivities struct {
	client appsynciface.AppSyncAPI
}

func NewAppSyncActivities(client appsynciface.AppSyncAPI) *AppSyncActivities {
    return &AppSyncActivities{client: client}
}

func (a *AppSyncActivities) CreateApiCache(input *appsync.CreateApiCacheInput) (*appsync.CreateApiCacheOutput, error) {
    return a.client.CreateApiCache(input)
}

func (a *AppSyncActivities) CreateApiKey(input *appsync.CreateApiKeyInput) (*appsync.CreateApiKeyOutput, error) {
    return a.client.CreateApiKey(input)
}

func (a *AppSyncActivities) CreateDataSource(input *appsync.CreateDataSourceInput) (*appsync.CreateDataSourceOutput, error) {
    return a.client.CreateDataSource(input)
}

func (a *AppSyncActivities) CreateFunction(input *appsync.CreateFunctionInput) (*appsync.CreateFunctionOutput, error) {
    return a.client.CreateFunction(input)
}

func (a *AppSyncActivities) CreateGraphqlApi(input *appsync.CreateGraphqlApiInput) (*appsync.CreateGraphqlApiOutput, error) {
    return a.client.CreateGraphqlApi(input)
}

func (a *AppSyncActivities) CreateResolver(input *appsync.CreateResolverInput) (*appsync.CreateResolverOutput, error) {
    return a.client.CreateResolver(input)
}

func (a *AppSyncActivities) CreateType(input *appsync.CreateTypeInput) (*appsync.CreateTypeOutput, error) {
    return a.client.CreateType(input)
}

func (a *AppSyncActivities) DeleteApiCache(input *appsync.DeleteApiCacheInput) (*appsync.DeleteApiCacheOutput, error) {
    return a.client.DeleteApiCache(input)
}

func (a *AppSyncActivities) DeleteApiKey(input *appsync.DeleteApiKeyInput) (*appsync.DeleteApiKeyOutput, error) {
    return a.client.DeleteApiKey(input)
}

func (a *AppSyncActivities) DeleteDataSource(input *appsync.DeleteDataSourceInput) (*appsync.DeleteDataSourceOutput, error) {
    return a.client.DeleteDataSource(input)
}

func (a *AppSyncActivities) DeleteFunction(input *appsync.DeleteFunctionInput) (*appsync.DeleteFunctionOutput, error) {
    return a.client.DeleteFunction(input)
}

func (a *AppSyncActivities) DeleteGraphqlApi(input *appsync.DeleteGraphqlApiInput) (*appsync.DeleteGraphqlApiOutput, error) {
    return a.client.DeleteGraphqlApi(input)
}

func (a *AppSyncActivities) DeleteResolver(input *appsync.DeleteResolverInput) (*appsync.DeleteResolverOutput, error) {
    return a.client.DeleteResolver(input)
}

func (a *AppSyncActivities) DeleteType(input *appsync.DeleteTypeInput) (*appsync.DeleteTypeOutput, error) {
    return a.client.DeleteType(input)
}

func (a *AppSyncActivities) FlushApiCache(input *appsync.FlushApiCacheInput) (*appsync.FlushApiCacheOutput, error) {
    return a.client.FlushApiCache(input)
}

func (a *AppSyncActivities) GetApiCache(input *appsync.GetApiCacheInput) (*appsync.GetApiCacheOutput, error) {
    return a.client.GetApiCache(input)
}

func (a *AppSyncActivities) GetDataSource(input *appsync.GetDataSourceInput) (*appsync.GetDataSourceOutput, error) {
    return a.client.GetDataSource(input)
}

func (a *AppSyncActivities) GetFunction(input *appsync.GetFunctionInput) (*appsync.GetFunctionOutput, error) {
    return a.client.GetFunction(input)
}

func (a *AppSyncActivities) GetGraphqlApi(input *appsync.GetGraphqlApiInput) (*appsync.GetGraphqlApiOutput, error) {
    return a.client.GetGraphqlApi(input)
}

func (a *AppSyncActivities) GetIntrospectionSchema(input *appsync.GetIntrospectionSchemaInput) (*appsync.GetIntrospectionSchemaOutput, error) {
    return a.client.GetIntrospectionSchema(input)
}

func (a *AppSyncActivities) GetResolver(input *appsync.GetResolverInput) (*appsync.GetResolverOutput, error) {
    return a.client.GetResolver(input)
}

func (a *AppSyncActivities) GetSchemaCreationStatus(input *appsync.GetSchemaCreationStatusInput) (*appsync.GetSchemaCreationStatusOutput, error) {
    return a.client.GetSchemaCreationStatus(input)
}

func (a *AppSyncActivities) GetType(input *appsync.GetTypeInput) (*appsync.GetTypeOutput, error) {
    return a.client.GetType(input)
}

func (a *AppSyncActivities) ListApiKeys(input *appsync.ListApiKeysInput) (*appsync.ListApiKeysOutput, error) {
    return a.client.ListApiKeys(input)
}

func (a *AppSyncActivities) ListDataSources(input *appsync.ListDataSourcesInput) (*appsync.ListDataSourcesOutput, error) {
    return a.client.ListDataSources(input)
}

func (a *AppSyncActivities) ListFunctions(input *appsync.ListFunctionsInput) (*appsync.ListFunctionsOutput, error) {
    return a.client.ListFunctions(input)
}

func (a *AppSyncActivities) ListGraphqlApis(input *appsync.ListGraphqlApisInput) (*appsync.ListGraphqlApisOutput, error) {
    return a.client.ListGraphqlApis(input)
}

func (a *AppSyncActivities) ListResolvers(input *appsync.ListResolversInput) (*appsync.ListResolversOutput, error) {
    return a.client.ListResolvers(input)
}

func (a *AppSyncActivities) ListResolversByFunction(input *appsync.ListResolversByFunctionInput) (*appsync.ListResolversByFunctionOutput, error) {
    return a.client.ListResolversByFunction(input)
}

func (a *AppSyncActivities) ListTagsForResource(input *appsync.ListTagsForResourceInput) (*appsync.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *AppSyncActivities) ListTypes(input *appsync.ListTypesInput) (*appsync.ListTypesOutput, error) {
    return a.client.ListTypes(input)
}

func (a *AppSyncActivities) StartSchemaCreation(input *appsync.StartSchemaCreationInput) (*appsync.StartSchemaCreationOutput, error) {
    return a.client.StartSchemaCreation(input)
}

func (a *AppSyncActivities) TagResource(input *appsync.TagResourceInput) (*appsync.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *AppSyncActivities) UntagResource(input *appsync.UntagResourceInput) (*appsync.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *AppSyncActivities) UpdateApiCache(input *appsync.UpdateApiCacheInput) (*appsync.UpdateApiCacheOutput, error) {
    return a.client.UpdateApiCache(input)
}

func (a *AppSyncActivities) UpdateApiKey(input *appsync.UpdateApiKeyInput) (*appsync.UpdateApiKeyOutput, error) {
    return a.client.UpdateApiKey(input)
}

func (a *AppSyncActivities) UpdateDataSource(input *appsync.UpdateDataSourceInput) (*appsync.UpdateDataSourceOutput, error) {
    return a.client.UpdateDataSource(input)
}

func (a *AppSyncActivities) UpdateFunction(input *appsync.UpdateFunctionInput) (*appsync.UpdateFunctionOutput, error) {
    return a.client.UpdateFunction(input)
}

func (a *AppSyncActivities) UpdateGraphqlApi(input *appsync.UpdateGraphqlApiInput) (*appsync.UpdateGraphqlApiOutput, error) {
    return a.client.UpdateGraphqlApi(input)
}

func (a *AppSyncActivities) UpdateResolver(input *appsync.UpdateResolverInput) (*appsync.UpdateResolverOutput, error) {
    return a.client.UpdateResolver(input)
}

func (a *AppSyncActivities) UpdateType(input *appsync.UpdateTypeInput) (*appsync.UpdateTypeOutput, error) {
    return a.client.UpdateType(input)
}
