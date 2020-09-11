package awsclients

import (
	"github.com/aws/aws-sdk-go/service/appsync"
	"go.temporal.io/sdk/workflow"
)

type AppSyncClient interface {
       CreateApiCache(ctx workflow.Context, input *appsync.CreateApiCacheInput) (*appsync.CreateApiCacheOutput, error)
       CreateApiCacheAsync(ctx workflow.Context, input *appsync.CreateApiCacheInput) *AppsyncCreateApiCacheResult

       CreateApiKey(ctx workflow.Context, input *appsync.CreateApiKeyInput) (*appsync.CreateApiKeyOutput, error)
       CreateApiKeyAsync(ctx workflow.Context, input *appsync.CreateApiKeyInput) *AppsyncCreateApiKeyResult

       CreateDataSource(ctx workflow.Context, input *appsync.CreateDataSourceInput) (*appsync.CreateDataSourceOutput, error)
       CreateDataSourceAsync(ctx workflow.Context, input *appsync.CreateDataSourceInput) *AppsyncCreateDataSourceResult

       CreateFunction(ctx workflow.Context, input *appsync.CreateFunctionInput) (*appsync.CreateFunctionOutput, error)
       CreateFunctionAsync(ctx workflow.Context, input *appsync.CreateFunctionInput) *AppsyncCreateFunctionResult

       CreateGraphqlApi(ctx workflow.Context, input *appsync.CreateGraphqlApiInput) (*appsync.CreateGraphqlApiOutput, error)
       CreateGraphqlApiAsync(ctx workflow.Context, input *appsync.CreateGraphqlApiInput) *AppsyncCreateGraphqlApiResult

       CreateResolver(ctx workflow.Context, input *appsync.CreateResolverInput) (*appsync.CreateResolverOutput, error)
       CreateResolverAsync(ctx workflow.Context, input *appsync.CreateResolverInput) *AppsyncCreateResolverResult

       CreateType(ctx workflow.Context, input *appsync.CreateTypeInput) (*appsync.CreateTypeOutput, error)
       CreateTypeAsync(ctx workflow.Context, input *appsync.CreateTypeInput) *AppsyncCreateTypeResult

       DeleteApiCache(ctx workflow.Context, input *appsync.DeleteApiCacheInput) (*appsync.DeleteApiCacheOutput, error)
       DeleteApiCacheAsync(ctx workflow.Context, input *appsync.DeleteApiCacheInput) *AppsyncDeleteApiCacheResult

       DeleteApiKey(ctx workflow.Context, input *appsync.DeleteApiKeyInput) (*appsync.DeleteApiKeyOutput, error)
       DeleteApiKeyAsync(ctx workflow.Context, input *appsync.DeleteApiKeyInput) *AppsyncDeleteApiKeyResult

       DeleteDataSource(ctx workflow.Context, input *appsync.DeleteDataSourceInput) (*appsync.DeleteDataSourceOutput, error)
       DeleteDataSourceAsync(ctx workflow.Context, input *appsync.DeleteDataSourceInput) *AppsyncDeleteDataSourceResult

       DeleteFunction(ctx workflow.Context, input *appsync.DeleteFunctionInput) (*appsync.DeleteFunctionOutput, error)
       DeleteFunctionAsync(ctx workflow.Context, input *appsync.DeleteFunctionInput) *AppsyncDeleteFunctionResult

       DeleteGraphqlApi(ctx workflow.Context, input *appsync.DeleteGraphqlApiInput) (*appsync.DeleteGraphqlApiOutput, error)
       DeleteGraphqlApiAsync(ctx workflow.Context, input *appsync.DeleteGraphqlApiInput) *AppsyncDeleteGraphqlApiResult

       DeleteResolver(ctx workflow.Context, input *appsync.DeleteResolverInput) (*appsync.DeleteResolverOutput, error)
       DeleteResolverAsync(ctx workflow.Context, input *appsync.DeleteResolverInput) *AppsyncDeleteResolverResult

       DeleteType(ctx workflow.Context, input *appsync.DeleteTypeInput) (*appsync.DeleteTypeOutput, error)
       DeleteTypeAsync(ctx workflow.Context, input *appsync.DeleteTypeInput) *AppsyncDeleteTypeResult

       FlushApiCache(ctx workflow.Context, input *appsync.FlushApiCacheInput) (*appsync.FlushApiCacheOutput, error)
       FlushApiCacheAsync(ctx workflow.Context, input *appsync.FlushApiCacheInput) *AppsyncFlushApiCacheResult

       GetApiCache(ctx workflow.Context, input *appsync.GetApiCacheInput) (*appsync.GetApiCacheOutput, error)
       GetApiCacheAsync(ctx workflow.Context, input *appsync.GetApiCacheInput) *AppsyncGetApiCacheResult

       GetDataSource(ctx workflow.Context, input *appsync.GetDataSourceInput) (*appsync.GetDataSourceOutput, error)
       GetDataSourceAsync(ctx workflow.Context, input *appsync.GetDataSourceInput) *AppsyncGetDataSourceResult

       GetFunction(ctx workflow.Context, input *appsync.GetFunctionInput) (*appsync.GetFunctionOutput, error)
       GetFunctionAsync(ctx workflow.Context, input *appsync.GetFunctionInput) *AppsyncGetFunctionResult

       GetGraphqlApi(ctx workflow.Context, input *appsync.GetGraphqlApiInput) (*appsync.GetGraphqlApiOutput, error)
       GetGraphqlApiAsync(ctx workflow.Context, input *appsync.GetGraphqlApiInput) *AppsyncGetGraphqlApiResult

       GetIntrospectionSchema(ctx workflow.Context, input *appsync.GetIntrospectionSchemaInput) (*appsync.GetIntrospectionSchemaOutput, error)
       GetIntrospectionSchemaAsync(ctx workflow.Context, input *appsync.GetIntrospectionSchemaInput) *AppsyncGetIntrospectionSchemaResult

       GetResolver(ctx workflow.Context, input *appsync.GetResolverInput) (*appsync.GetResolverOutput, error)
       GetResolverAsync(ctx workflow.Context, input *appsync.GetResolverInput) *AppsyncGetResolverResult

       GetSchemaCreationStatus(ctx workflow.Context, input *appsync.GetSchemaCreationStatusInput) (*appsync.GetSchemaCreationStatusOutput, error)
       GetSchemaCreationStatusAsync(ctx workflow.Context, input *appsync.GetSchemaCreationStatusInput) *AppsyncGetSchemaCreationStatusResult

       GetType(ctx workflow.Context, input *appsync.GetTypeInput) (*appsync.GetTypeOutput, error)
       GetTypeAsync(ctx workflow.Context, input *appsync.GetTypeInput) *AppsyncGetTypeResult

       ListApiKeys(ctx workflow.Context, input *appsync.ListApiKeysInput) (*appsync.ListApiKeysOutput, error)
       ListApiKeysAsync(ctx workflow.Context, input *appsync.ListApiKeysInput) *AppsyncListApiKeysResult

       ListDataSources(ctx workflow.Context, input *appsync.ListDataSourcesInput) (*appsync.ListDataSourcesOutput, error)
       ListDataSourcesAsync(ctx workflow.Context, input *appsync.ListDataSourcesInput) *AppsyncListDataSourcesResult

       ListFunctions(ctx workflow.Context, input *appsync.ListFunctionsInput) (*appsync.ListFunctionsOutput, error)
       ListFunctionsAsync(ctx workflow.Context, input *appsync.ListFunctionsInput) *AppsyncListFunctionsResult

       ListGraphqlApis(ctx workflow.Context, input *appsync.ListGraphqlApisInput) (*appsync.ListGraphqlApisOutput, error)
       ListGraphqlApisAsync(ctx workflow.Context, input *appsync.ListGraphqlApisInput) *AppsyncListGraphqlApisResult

       ListResolvers(ctx workflow.Context, input *appsync.ListResolversInput) (*appsync.ListResolversOutput, error)
       ListResolversAsync(ctx workflow.Context, input *appsync.ListResolversInput) *AppsyncListResolversResult

       ListResolversByFunction(ctx workflow.Context, input *appsync.ListResolversByFunctionInput) (*appsync.ListResolversByFunctionOutput, error)
       ListResolversByFunctionAsync(ctx workflow.Context, input *appsync.ListResolversByFunctionInput) *AppsyncListResolversByFunctionResult

       ListTagsForResource(ctx workflow.Context, input *appsync.ListTagsForResourceInput) (*appsync.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *appsync.ListTagsForResourceInput) *AppsyncListTagsForResourceResult

       ListTypes(ctx workflow.Context, input *appsync.ListTypesInput) (*appsync.ListTypesOutput, error)
       ListTypesAsync(ctx workflow.Context, input *appsync.ListTypesInput) *AppsyncListTypesResult

       StartSchemaCreation(ctx workflow.Context, input *appsync.StartSchemaCreationInput) (*appsync.StartSchemaCreationOutput, error)
       StartSchemaCreationAsync(ctx workflow.Context, input *appsync.StartSchemaCreationInput) *AppsyncStartSchemaCreationResult

       TagResource(ctx workflow.Context, input *appsync.TagResourceInput) (*appsync.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *appsync.TagResourceInput) *AppsyncTagResourceResult

       UntagResource(ctx workflow.Context, input *appsync.UntagResourceInput) (*appsync.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *appsync.UntagResourceInput) *AppsyncUntagResourceResult

       UpdateApiCache(ctx workflow.Context, input *appsync.UpdateApiCacheInput) (*appsync.UpdateApiCacheOutput, error)
       UpdateApiCacheAsync(ctx workflow.Context, input *appsync.UpdateApiCacheInput) *AppsyncUpdateApiCacheResult

       UpdateApiKey(ctx workflow.Context, input *appsync.UpdateApiKeyInput) (*appsync.UpdateApiKeyOutput, error)
       UpdateApiKeyAsync(ctx workflow.Context, input *appsync.UpdateApiKeyInput) *AppsyncUpdateApiKeyResult

       UpdateDataSource(ctx workflow.Context, input *appsync.UpdateDataSourceInput) (*appsync.UpdateDataSourceOutput, error)
       UpdateDataSourceAsync(ctx workflow.Context, input *appsync.UpdateDataSourceInput) *AppsyncUpdateDataSourceResult

       UpdateFunction(ctx workflow.Context, input *appsync.UpdateFunctionInput) (*appsync.UpdateFunctionOutput, error)
       UpdateFunctionAsync(ctx workflow.Context, input *appsync.UpdateFunctionInput) *AppsyncUpdateFunctionResult

       UpdateGraphqlApi(ctx workflow.Context, input *appsync.UpdateGraphqlApiInput) (*appsync.UpdateGraphqlApiOutput, error)
       UpdateGraphqlApiAsync(ctx workflow.Context, input *appsync.UpdateGraphqlApiInput) *AppsyncUpdateGraphqlApiResult

       UpdateResolver(ctx workflow.Context, input *appsync.UpdateResolverInput) (*appsync.UpdateResolverOutput, error)
       UpdateResolverAsync(ctx workflow.Context, input *appsync.UpdateResolverInput) *AppsyncUpdateResolverResult

       UpdateType(ctx workflow.Context, input *appsync.UpdateTypeInput) (*appsync.UpdateTypeOutput, error)
       UpdateTypeAsync(ctx workflow.Context, input *appsync.UpdateTypeInput) *AppsyncUpdateTypeResult
}

type AppSyncStub struct {
}

func NewAppSyncStub() AppSyncClient {
    return &AppSyncStub{}
}

type AppsyncCreateApiCacheResult struct {
	Result workflow.Future
}

func (r *AppsyncCreateApiCacheResult) Get(ctx workflow.Context) (*appsync.CreateApiCacheOutput, error) {
    var output appsync.CreateApiCacheOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncCreateApiKeyResult struct {
	Result workflow.Future
}

func (r *AppsyncCreateApiKeyResult) Get(ctx workflow.Context) (*appsync.CreateApiKeyOutput, error) {
    var output appsync.CreateApiKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncCreateDataSourceResult struct {
	Result workflow.Future
}

func (r *AppsyncCreateDataSourceResult) Get(ctx workflow.Context) (*appsync.CreateDataSourceOutput, error) {
    var output appsync.CreateDataSourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncCreateFunctionResult struct {
	Result workflow.Future
}

func (r *AppsyncCreateFunctionResult) Get(ctx workflow.Context) (*appsync.CreateFunctionOutput, error) {
    var output appsync.CreateFunctionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncCreateGraphqlApiResult struct {
	Result workflow.Future
}

func (r *AppsyncCreateGraphqlApiResult) Get(ctx workflow.Context) (*appsync.CreateGraphqlApiOutput, error) {
    var output appsync.CreateGraphqlApiOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncCreateResolverResult struct {
	Result workflow.Future
}

func (r *AppsyncCreateResolverResult) Get(ctx workflow.Context) (*appsync.CreateResolverOutput, error) {
    var output appsync.CreateResolverOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncCreateTypeResult struct {
	Result workflow.Future
}

func (r *AppsyncCreateTypeResult) Get(ctx workflow.Context) (*appsync.CreateTypeOutput, error) {
    var output appsync.CreateTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncDeleteApiCacheResult struct {
	Result workflow.Future
}

func (r *AppsyncDeleteApiCacheResult) Get(ctx workflow.Context) (*appsync.DeleteApiCacheOutput, error) {
    var output appsync.DeleteApiCacheOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncDeleteApiKeyResult struct {
	Result workflow.Future
}

func (r *AppsyncDeleteApiKeyResult) Get(ctx workflow.Context) (*appsync.DeleteApiKeyOutput, error) {
    var output appsync.DeleteApiKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncDeleteDataSourceResult struct {
	Result workflow.Future
}

func (r *AppsyncDeleteDataSourceResult) Get(ctx workflow.Context) (*appsync.DeleteDataSourceOutput, error) {
    var output appsync.DeleteDataSourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncDeleteFunctionResult struct {
	Result workflow.Future
}

func (r *AppsyncDeleteFunctionResult) Get(ctx workflow.Context) (*appsync.DeleteFunctionOutput, error) {
    var output appsync.DeleteFunctionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncDeleteGraphqlApiResult struct {
	Result workflow.Future
}

func (r *AppsyncDeleteGraphqlApiResult) Get(ctx workflow.Context) (*appsync.DeleteGraphqlApiOutput, error) {
    var output appsync.DeleteGraphqlApiOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncDeleteResolverResult struct {
	Result workflow.Future
}

func (r *AppsyncDeleteResolverResult) Get(ctx workflow.Context) (*appsync.DeleteResolverOutput, error) {
    var output appsync.DeleteResolverOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncDeleteTypeResult struct {
	Result workflow.Future
}

func (r *AppsyncDeleteTypeResult) Get(ctx workflow.Context) (*appsync.DeleteTypeOutput, error) {
    var output appsync.DeleteTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncFlushApiCacheResult struct {
	Result workflow.Future
}

func (r *AppsyncFlushApiCacheResult) Get(ctx workflow.Context) (*appsync.FlushApiCacheOutput, error) {
    var output appsync.FlushApiCacheOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncGetApiCacheResult struct {
	Result workflow.Future
}

func (r *AppsyncGetApiCacheResult) Get(ctx workflow.Context) (*appsync.GetApiCacheOutput, error) {
    var output appsync.GetApiCacheOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncGetDataSourceResult struct {
	Result workflow.Future
}

func (r *AppsyncGetDataSourceResult) Get(ctx workflow.Context) (*appsync.GetDataSourceOutput, error) {
    var output appsync.GetDataSourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncGetFunctionResult struct {
	Result workflow.Future
}

func (r *AppsyncGetFunctionResult) Get(ctx workflow.Context) (*appsync.GetFunctionOutput, error) {
    var output appsync.GetFunctionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncGetGraphqlApiResult struct {
	Result workflow.Future
}

func (r *AppsyncGetGraphqlApiResult) Get(ctx workflow.Context) (*appsync.GetGraphqlApiOutput, error) {
    var output appsync.GetGraphqlApiOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncGetIntrospectionSchemaResult struct {
	Result workflow.Future
}

func (r *AppsyncGetIntrospectionSchemaResult) Get(ctx workflow.Context) (*appsync.GetIntrospectionSchemaOutput, error) {
    var output appsync.GetIntrospectionSchemaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncGetResolverResult struct {
	Result workflow.Future
}

func (r *AppsyncGetResolverResult) Get(ctx workflow.Context) (*appsync.GetResolverOutput, error) {
    var output appsync.GetResolverOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncGetSchemaCreationStatusResult struct {
	Result workflow.Future
}

func (r *AppsyncGetSchemaCreationStatusResult) Get(ctx workflow.Context) (*appsync.GetSchemaCreationStatusOutput, error) {
    var output appsync.GetSchemaCreationStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncGetTypeResult struct {
	Result workflow.Future
}

func (r *AppsyncGetTypeResult) Get(ctx workflow.Context) (*appsync.GetTypeOutput, error) {
    var output appsync.GetTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncListApiKeysResult struct {
	Result workflow.Future
}

func (r *AppsyncListApiKeysResult) Get(ctx workflow.Context) (*appsync.ListApiKeysOutput, error) {
    var output appsync.ListApiKeysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncListDataSourcesResult struct {
	Result workflow.Future
}

func (r *AppsyncListDataSourcesResult) Get(ctx workflow.Context) (*appsync.ListDataSourcesOutput, error) {
    var output appsync.ListDataSourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncListFunctionsResult struct {
	Result workflow.Future
}

func (r *AppsyncListFunctionsResult) Get(ctx workflow.Context) (*appsync.ListFunctionsOutput, error) {
    var output appsync.ListFunctionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncListGraphqlApisResult struct {
	Result workflow.Future
}

func (r *AppsyncListGraphqlApisResult) Get(ctx workflow.Context) (*appsync.ListGraphqlApisOutput, error) {
    var output appsync.ListGraphqlApisOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncListResolversResult struct {
	Result workflow.Future
}

func (r *AppsyncListResolversResult) Get(ctx workflow.Context) (*appsync.ListResolversOutput, error) {
    var output appsync.ListResolversOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncListResolversByFunctionResult struct {
	Result workflow.Future
}

func (r *AppsyncListResolversByFunctionResult) Get(ctx workflow.Context) (*appsync.ListResolversByFunctionOutput, error) {
    var output appsync.ListResolversByFunctionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *AppsyncListTagsForResourceResult) Get(ctx workflow.Context) (*appsync.ListTagsForResourceOutput, error) {
    var output appsync.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncListTypesResult struct {
	Result workflow.Future
}

func (r *AppsyncListTypesResult) Get(ctx workflow.Context) (*appsync.ListTypesOutput, error) {
    var output appsync.ListTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncStartSchemaCreationResult struct {
	Result workflow.Future
}

func (r *AppsyncStartSchemaCreationResult) Get(ctx workflow.Context) (*appsync.StartSchemaCreationOutput, error) {
    var output appsync.StartSchemaCreationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncTagResourceResult struct {
	Result workflow.Future
}

func (r *AppsyncTagResourceResult) Get(ctx workflow.Context) (*appsync.TagResourceOutput, error) {
    var output appsync.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncUntagResourceResult struct {
	Result workflow.Future
}

func (r *AppsyncUntagResourceResult) Get(ctx workflow.Context) (*appsync.UntagResourceOutput, error) {
    var output appsync.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncUpdateApiCacheResult struct {
	Result workflow.Future
}

func (r *AppsyncUpdateApiCacheResult) Get(ctx workflow.Context) (*appsync.UpdateApiCacheOutput, error) {
    var output appsync.UpdateApiCacheOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncUpdateApiKeyResult struct {
	Result workflow.Future
}

func (r *AppsyncUpdateApiKeyResult) Get(ctx workflow.Context) (*appsync.UpdateApiKeyOutput, error) {
    var output appsync.UpdateApiKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncUpdateDataSourceResult struct {
	Result workflow.Future
}

func (r *AppsyncUpdateDataSourceResult) Get(ctx workflow.Context) (*appsync.UpdateDataSourceOutput, error) {
    var output appsync.UpdateDataSourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncUpdateFunctionResult struct {
	Result workflow.Future
}

func (r *AppsyncUpdateFunctionResult) Get(ctx workflow.Context) (*appsync.UpdateFunctionOutput, error) {
    var output appsync.UpdateFunctionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncUpdateGraphqlApiResult struct {
	Result workflow.Future
}

func (r *AppsyncUpdateGraphqlApiResult) Get(ctx workflow.Context) (*appsync.UpdateGraphqlApiOutput, error) {
    var output appsync.UpdateGraphqlApiOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncUpdateResolverResult struct {
	Result workflow.Future
}

func (r *AppsyncUpdateResolverResult) Get(ctx workflow.Context) (*appsync.UpdateResolverOutput, error) {
    var output appsync.UpdateResolverOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type AppsyncUpdateTypeResult struct {
	Result workflow.Future
}

func (r *AppsyncUpdateTypeResult) Get(ctx workflow.Context) (*appsync.UpdateTypeOutput, error) {
    var output appsync.UpdateTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) CreateApiCache(ctx workflow.Context, input *appsync.CreateApiCacheInput) (*appsync.CreateApiCacheOutput, error) {
    var output appsync.CreateApiCacheOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.CreateApiCache", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) CreateApiCacheAsync(ctx workflow.Context, input *appsync.CreateApiCacheInput) *AppsyncCreateApiCacheResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.CreateApiCache", input)
    return &AppsyncCreateApiCacheResult{Result: future}
}

func (a *AppSyncStub) CreateApiKey(ctx workflow.Context, input *appsync.CreateApiKeyInput) (*appsync.CreateApiKeyOutput, error) {
    var output appsync.CreateApiKeyOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.CreateApiKey", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) CreateApiKeyAsync(ctx workflow.Context, input *appsync.CreateApiKeyInput) *AppsyncCreateApiKeyResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.CreateApiKey", input)
    return &AppsyncCreateApiKeyResult{Result: future}
}

func (a *AppSyncStub) CreateDataSource(ctx workflow.Context, input *appsync.CreateDataSourceInput) (*appsync.CreateDataSourceOutput, error) {
    var output appsync.CreateDataSourceOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.CreateDataSource", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) CreateDataSourceAsync(ctx workflow.Context, input *appsync.CreateDataSourceInput) *AppsyncCreateDataSourceResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.CreateDataSource", input)
    return &AppsyncCreateDataSourceResult{Result: future}
}

func (a *AppSyncStub) CreateFunction(ctx workflow.Context, input *appsync.CreateFunctionInput) (*appsync.CreateFunctionOutput, error) {
    var output appsync.CreateFunctionOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.CreateFunction", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) CreateFunctionAsync(ctx workflow.Context, input *appsync.CreateFunctionInput) *AppsyncCreateFunctionResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.CreateFunction", input)
    return &AppsyncCreateFunctionResult{Result: future}
}

func (a *AppSyncStub) CreateGraphqlApi(ctx workflow.Context, input *appsync.CreateGraphqlApiInput) (*appsync.CreateGraphqlApiOutput, error) {
    var output appsync.CreateGraphqlApiOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.CreateGraphqlApi", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) CreateGraphqlApiAsync(ctx workflow.Context, input *appsync.CreateGraphqlApiInput) *AppsyncCreateGraphqlApiResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.CreateGraphqlApi", input)
    return &AppsyncCreateGraphqlApiResult{Result: future}
}

func (a *AppSyncStub) CreateResolver(ctx workflow.Context, input *appsync.CreateResolverInput) (*appsync.CreateResolverOutput, error) {
    var output appsync.CreateResolverOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.CreateResolver", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) CreateResolverAsync(ctx workflow.Context, input *appsync.CreateResolverInput) *AppsyncCreateResolverResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.CreateResolver", input)
    return &AppsyncCreateResolverResult{Result: future}
}

func (a *AppSyncStub) CreateType(ctx workflow.Context, input *appsync.CreateTypeInput) (*appsync.CreateTypeOutput, error) {
    var output appsync.CreateTypeOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.CreateType", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) CreateTypeAsync(ctx workflow.Context, input *appsync.CreateTypeInput) *AppsyncCreateTypeResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.CreateType", input)
    return &AppsyncCreateTypeResult{Result: future}
}

func (a *AppSyncStub) DeleteApiCache(ctx workflow.Context, input *appsync.DeleteApiCacheInput) (*appsync.DeleteApiCacheOutput, error) {
    var output appsync.DeleteApiCacheOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.DeleteApiCache", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) DeleteApiCacheAsync(ctx workflow.Context, input *appsync.DeleteApiCacheInput) *AppsyncDeleteApiCacheResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.DeleteApiCache", input)
    return &AppsyncDeleteApiCacheResult{Result: future}
}

func (a *AppSyncStub) DeleteApiKey(ctx workflow.Context, input *appsync.DeleteApiKeyInput) (*appsync.DeleteApiKeyOutput, error) {
    var output appsync.DeleteApiKeyOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.DeleteApiKey", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) DeleteApiKeyAsync(ctx workflow.Context, input *appsync.DeleteApiKeyInput) *AppsyncDeleteApiKeyResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.DeleteApiKey", input)
    return &AppsyncDeleteApiKeyResult{Result: future}
}

func (a *AppSyncStub) DeleteDataSource(ctx workflow.Context, input *appsync.DeleteDataSourceInput) (*appsync.DeleteDataSourceOutput, error) {
    var output appsync.DeleteDataSourceOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.DeleteDataSource", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) DeleteDataSourceAsync(ctx workflow.Context, input *appsync.DeleteDataSourceInput) *AppsyncDeleteDataSourceResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.DeleteDataSource", input)
    return &AppsyncDeleteDataSourceResult{Result: future}
}

func (a *AppSyncStub) DeleteFunction(ctx workflow.Context, input *appsync.DeleteFunctionInput) (*appsync.DeleteFunctionOutput, error) {
    var output appsync.DeleteFunctionOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.DeleteFunction", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) DeleteFunctionAsync(ctx workflow.Context, input *appsync.DeleteFunctionInput) *AppsyncDeleteFunctionResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.DeleteFunction", input)
    return &AppsyncDeleteFunctionResult{Result: future}
}

func (a *AppSyncStub) DeleteGraphqlApi(ctx workflow.Context, input *appsync.DeleteGraphqlApiInput) (*appsync.DeleteGraphqlApiOutput, error) {
    var output appsync.DeleteGraphqlApiOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.DeleteGraphqlApi", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) DeleteGraphqlApiAsync(ctx workflow.Context, input *appsync.DeleteGraphqlApiInput) *AppsyncDeleteGraphqlApiResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.DeleteGraphqlApi", input)
    return &AppsyncDeleteGraphqlApiResult{Result: future}
}

func (a *AppSyncStub) DeleteResolver(ctx workflow.Context, input *appsync.DeleteResolverInput) (*appsync.DeleteResolverOutput, error) {
    var output appsync.DeleteResolverOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.DeleteResolver", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) DeleteResolverAsync(ctx workflow.Context, input *appsync.DeleteResolverInput) *AppsyncDeleteResolverResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.DeleteResolver", input)
    return &AppsyncDeleteResolverResult{Result: future}
}

func (a *AppSyncStub) DeleteType(ctx workflow.Context, input *appsync.DeleteTypeInput) (*appsync.DeleteTypeOutput, error) {
    var output appsync.DeleteTypeOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.DeleteType", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) DeleteTypeAsync(ctx workflow.Context, input *appsync.DeleteTypeInput) *AppsyncDeleteTypeResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.DeleteType", input)
    return &AppsyncDeleteTypeResult{Result: future}
}

func (a *AppSyncStub) FlushApiCache(ctx workflow.Context, input *appsync.FlushApiCacheInput) (*appsync.FlushApiCacheOutput, error) {
    var output appsync.FlushApiCacheOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.FlushApiCache", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) FlushApiCacheAsync(ctx workflow.Context, input *appsync.FlushApiCacheInput) *AppsyncFlushApiCacheResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.FlushApiCache", input)
    return &AppsyncFlushApiCacheResult{Result: future}
}

func (a *AppSyncStub) GetApiCache(ctx workflow.Context, input *appsync.GetApiCacheInput) (*appsync.GetApiCacheOutput, error) {
    var output appsync.GetApiCacheOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.GetApiCache", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) GetApiCacheAsync(ctx workflow.Context, input *appsync.GetApiCacheInput) *AppsyncGetApiCacheResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.GetApiCache", input)
    return &AppsyncGetApiCacheResult{Result: future}
}

func (a *AppSyncStub) GetDataSource(ctx workflow.Context, input *appsync.GetDataSourceInput) (*appsync.GetDataSourceOutput, error) {
    var output appsync.GetDataSourceOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.GetDataSource", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) GetDataSourceAsync(ctx workflow.Context, input *appsync.GetDataSourceInput) *AppsyncGetDataSourceResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.GetDataSource", input)
    return &AppsyncGetDataSourceResult{Result: future}
}

func (a *AppSyncStub) GetFunction(ctx workflow.Context, input *appsync.GetFunctionInput) (*appsync.GetFunctionOutput, error) {
    var output appsync.GetFunctionOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.GetFunction", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) GetFunctionAsync(ctx workflow.Context, input *appsync.GetFunctionInput) *AppsyncGetFunctionResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.GetFunction", input)
    return &AppsyncGetFunctionResult{Result: future}
}

func (a *AppSyncStub) GetGraphqlApi(ctx workflow.Context, input *appsync.GetGraphqlApiInput) (*appsync.GetGraphqlApiOutput, error) {
    var output appsync.GetGraphqlApiOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.GetGraphqlApi", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) GetGraphqlApiAsync(ctx workflow.Context, input *appsync.GetGraphqlApiInput) *AppsyncGetGraphqlApiResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.GetGraphqlApi", input)
    return &AppsyncGetGraphqlApiResult{Result: future}
}

func (a *AppSyncStub) GetIntrospectionSchema(ctx workflow.Context, input *appsync.GetIntrospectionSchemaInput) (*appsync.GetIntrospectionSchemaOutput, error) {
    var output appsync.GetIntrospectionSchemaOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.GetIntrospectionSchema", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) GetIntrospectionSchemaAsync(ctx workflow.Context, input *appsync.GetIntrospectionSchemaInput) *AppsyncGetIntrospectionSchemaResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.GetIntrospectionSchema", input)
    return &AppsyncGetIntrospectionSchemaResult{Result: future}
}

func (a *AppSyncStub) GetResolver(ctx workflow.Context, input *appsync.GetResolverInput) (*appsync.GetResolverOutput, error) {
    var output appsync.GetResolverOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.GetResolver", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) GetResolverAsync(ctx workflow.Context, input *appsync.GetResolverInput) *AppsyncGetResolverResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.GetResolver", input)
    return &AppsyncGetResolverResult{Result: future}
}

func (a *AppSyncStub) GetSchemaCreationStatus(ctx workflow.Context, input *appsync.GetSchemaCreationStatusInput) (*appsync.GetSchemaCreationStatusOutput, error) {
    var output appsync.GetSchemaCreationStatusOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.GetSchemaCreationStatus", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) GetSchemaCreationStatusAsync(ctx workflow.Context, input *appsync.GetSchemaCreationStatusInput) *AppsyncGetSchemaCreationStatusResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.GetSchemaCreationStatus", input)
    return &AppsyncGetSchemaCreationStatusResult{Result: future}
}

func (a *AppSyncStub) GetType(ctx workflow.Context, input *appsync.GetTypeInput) (*appsync.GetTypeOutput, error) {
    var output appsync.GetTypeOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.GetType", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) GetTypeAsync(ctx workflow.Context, input *appsync.GetTypeInput) *AppsyncGetTypeResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.GetType", input)
    return &AppsyncGetTypeResult{Result: future}
}

func (a *AppSyncStub) ListApiKeys(ctx workflow.Context, input *appsync.ListApiKeysInput) (*appsync.ListApiKeysOutput, error) {
    var output appsync.ListApiKeysOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.ListApiKeys", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) ListApiKeysAsync(ctx workflow.Context, input *appsync.ListApiKeysInput) *AppsyncListApiKeysResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.ListApiKeys", input)
    return &AppsyncListApiKeysResult{Result: future}
}

func (a *AppSyncStub) ListDataSources(ctx workflow.Context, input *appsync.ListDataSourcesInput) (*appsync.ListDataSourcesOutput, error) {
    var output appsync.ListDataSourcesOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.ListDataSources", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) ListDataSourcesAsync(ctx workflow.Context, input *appsync.ListDataSourcesInput) *AppsyncListDataSourcesResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.ListDataSources", input)
    return &AppsyncListDataSourcesResult{Result: future}
}

func (a *AppSyncStub) ListFunctions(ctx workflow.Context, input *appsync.ListFunctionsInput) (*appsync.ListFunctionsOutput, error) {
    var output appsync.ListFunctionsOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.ListFunctions", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) ListFunctionsAsync(ctx workflow.Context, input *appsync.ListFunctionsInput) *AppsyncListFunctionsResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.ListFunctions", input)
    return &AppsyncListFunctionsResult{Result: future}
}

func (a *AppSyncStub) ListGraphqlApis(ctx workflow.Context, input *appsync.ListGraphqlApisInput) (*appsync.ListGraphqlApisOutput, error) {
    var output appsync.ListGraphqlApisOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.ListGraphqlApis", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) ListGraphqlApisAsync(ctx workflow.Context, input *appsync.ListGraphqlApisInput) *AppsyncListGraphqlApisResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.ListGraphqlApis", input)
    return &AppsyncListGraphqlApisResult{Result: future}
}

func (a *AppSyncStub) ListResolvers(ctx workflow.Context, input *appsync.ListResolversInput) (*appsync.ListResolversOutput, error) {
    var output appsync.ListResolversOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.ListResolvers", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) ListResolversAsync(ctx workflow.Context, input *appsync.ListResolversInput) *AppsyncListResolversResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.ListResolvers", input)
    return &AppsyncListResolversResult{Result: future}
}

func (a *AppSyncStub) ListResolversByFunction(ctx workflow.Context, input *appsync.ListResolversByFunctionInput) (*appsync.ListResolversByFunctionOutput, error) {
    var output appsync.ListResolversByFunctionOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.ListResolversByFunction", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) ListResolversByFunctionAsync(ctx workflow.Context, input *appsync.ListResolversByFunctionInput) *AppsyncListResolversByFunctionResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.ListResolversByFunction", input)
    return &AppsyncListResolversByFunctionResult{Result: future}
}

func (a *AppSyncStub) ListTagsForResource(ctx workflow.Context, input *appsync.ListTagsForResourceInput) (*appsync.ListTagsForResourceOutput, error) {
    var output appsync.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.ListTagsForResource", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) ListTagsForResourceAsync(ctx workflow.Context, input *appsync.ListTagsForResourceInput) *AppsyncListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.ListTagsForResource", input)
    return &AppsyncListTagsForResourceResult{Result: future}
}

func (a *AppSyncStub) ListTypes(ctx workflow.Context, input *appsync.ListTypesInput) (*appsync.ListTypesOutput, error) {
    var output appsync.ListTypesOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.ListTypes", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) ListTypesAsync(ctx workflow.Context, input *appsync.ListTypesInput) *AppsyncListTypesResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.ListTypes", input)
    return &AppsyncListTypesResult{Result: future}
}

func (a *AppSyncStub) StartSchemaCreation(ctx workflow.Context, input *appsync.StartSchemaCreationInput) (*appsync.StartSchemaCreationOutput, error) {
    var output appsync.StartSchemaCreationOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.StartSchemaCreation", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) StartSchemaCreationAsync(ctx workflow.Context, input *appsync.StartSchemaCreationInput) *AppsyncStartSchemaCreationResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.StartSchemaCreation", input)
    return &AppsyncStartSchemaCreationResult{Result: future}
}

func (a *AppSyncStub) TagResource(ctx workflow.Context, input *appsync.TagResourceInput) (*appsync.TagResourceOutput, error) {
    var output appsync.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.TagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) TagResourceAsync(ctx workflow.Context, input *appsync.TagResourceInput) *AppsyncTagResourceResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.TagResource", input)
    return &AppsyncTagResourceResult{Result: future}
}

func (a *AppSyncStub) UntagResource(ctx workflow.Context, input *appsync.UntagResourceInput) (*appsync.UntagResourceOutput, error) {
    var output appsync.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.UntagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) UntagResourceAsync(ctx workflow.Context, input *appsync.UntagResourceInput) *AppsyncUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.UntagResource", input)
    return &AppsyncUntagResourceResult{Result: future}
}

func (a *AppSyncStub) UpdateApiCache(ctx workflow.Context, input *appsync.UpdateApiCacheInput) (*appsync.UpdateApiCacheOutput, error) {
    var output appsync.UpdateApiCacheOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.UpdateApiCache", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) UpdateApiCacheAsync(ctx workflow.Context, input *appsync.UpdateApiCacheInput) *AppsyncUpdateApiCacheResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.UpdateApiCache", input)
    return &AppsyncUpdateApiCacheResult{Result: future}
}

func (a *AppSyncStub) UpdateApiKey(ctx workflow.Context, input *appsync.UpdateApiKeyInput) (*appsync.UpdateApiKeyOutput, error) {
    var output appsync.UpdateApiKeyOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.UpdateApiKey", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) UpdateApiKeyAsync(ctx workflow.Context, input *appsync.UpdateApiKeyInput) *AppsyncUpdateApiKeyResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.UpdateApiKey", input)
    return &AppsyncUpdateApiKeyResult{Result: future}
}

func (a *AppSyncStub) UpdateDataSource(ctx workflow.Context, input *appsync.UpdateDataSourceInput) (*appsync.UpdateDataSourceOutput, error) {
    var output appsync.UpdateDataSourceOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.UpdateDataSource", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) UpdateDataSourceAsync(ctx workflow.Context, input *appsync.UpdateDataSourceInput) *AppsyncUpdateDataSourceResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.UpdateDataSource", input)
    return &AppsyncUpdateDataSourceResult{Result: future}
}

func (a *AppSyncStub) UpdateFunction(ctx workflow.Context, input *appsync.UpdateFunctionInput) (*appsync.UpdateFunctionOutput, error) {
    var output appsync.UpdateFunctionOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.UpdateFunction", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) UpdateFunctionAsync(ctx workflow.Context, input *appsync.UpdateFunctionInput) *AppsyncUpdateFunctionResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.UpdateFunction", input)
    return &AppsyncUpdateFunctionResult{Result: future}
}

func (a *AppSyncStub) UpdateGraphqlApi(ctx workflow.Context, input *appsync.UpdateGraphqlApiInput) (*appsync.UpdateGraphqlApiOutput, error) {
    var output appsync.UpdateGraphqlApiOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.UpdateGraphqlApi", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) UpdateGraphqlApiAsync(ctx workflow.Context, input *appsync.UpdateGraphqlApiInput) *AppsyncUpdateGraphqlApiResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.UpdateGraphqlApi", input)
    return &AppsyncUpdateGraphqlApiResult{Result: future}
}

func (a *AppSyncStub) UpdateResolver(ctx workflow.Context, input *appsync.UpdateResolverInput) (*appsync.UpdateResolverOutput, error) {
    var output appsync.UpdateResolverOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.UpdateResolver", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) UpdateResolverAsync(ctx workflow.Context, input *appsync.UpdateResolverInput) *AppsyncUpdateResolverResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.UpdateResolver", input)
    return &AppsyncUpdateResolverResult{Result: future}
}

func (a *AppSyncStub) UpdateType(ctx workflow.Context, input *appsync.UpdateTypeInput) (*appsync.UpdateTypeOutput, error) {
    var output appsync.UpdateTypeOutput
    err := workflow.ExecuteActivity(ctx, "AppSync.UpdateType", input).Get(ctx, &output)
    return &output, err
}

func (a *AppSyncStub) UpdateTypeAsync(ctx workflow.Context, input *appsync.UpdateTypeInput) *AppsyncUpdateTypeResult {
    future := workflow.ExecuteActivity(ctx, "AppSync.UpdateType", input)
    return &AppsyncUpdateTypeResult{Result: future}
}
