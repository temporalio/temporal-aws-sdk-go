package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/schemas"
	"go.temporal.io/sdk/workflow"
)

type SchemasClient interface {
    CreateDiscoverer(ctx workflow.Context, input *schemas.CreateDiscovererInput) (*schemas.CreateDiscovererOutput, error)
    CreateDiscovererAsync(ctx workflow.Context, input *schemas.CreateDiscovererInput) *SchemasCreateDiscovererResult

    CreateRegistry(ctx workflow.Context, input *schemas.CreateRegistryInput) (*schemas.CreateRegistryOutput, error)
    CreateRegistryAsync(ctx workflow.Context, input *schemas.CreateRegistryInput) *SchemasCreateRegistryResult

    CreateSchema(ctx workflow.Context, input *schemas.CreateSchemaInput) (*schemas.CreateSchemaOutput, error)
    CreateSchemaAsync(ctx workflow.Context, input *schemas.CreateSchemaInput) *SchemasCreateSchemaResult

    DeleteDiscoverer(ctx workflow.Context, input *schemas.DeleteDiscovererInput) (*schemas.DeleteDiscovererOutput, error)
    DeleteDiscovererAsync(ctx workflow.Context, input *schemas.DeleteDiscovererInput) *SchemasDeleteDiscovererResult

    DeleteRegistry(ctx workflow.Context, input *schemas.DeleteRegistryInput) (*schemas.DeleteRegistryOutput, error)
    DeleteRegistryAsync(ctx workflow.Context, input *schemas.DeleteRegistryInput) *SchemasDeleteRegistryResult

    DeleteResourcePolicy(ctx workflow.Context, input *schemas.DeleteResourcePolicyInput) (*schemas.DeleteResourcePolicyOutput, error)
    DeleteResourcePolicyAsync(ctx workflow.Context, input *schemas.DeleteResourcePolicyInput) *SchemasDeleteResourcePolicyResult

    DeleteSchema(ctx workflow.Context, input *schemas.DeleteSchemaInput) (*schemas.DeleteSchemaOutput, error)
    DeleteSchemaAsync(ctx workflow.Context, input *schemas.DeleteSchemaInput) *SchemasDeleteSchemaResult

    DeleteSchemaVersion(ctx workflow.Context, input *schemas.DeleteSchemaVersionInput) (*schemas.DeleteSchemaVersionOutput, error)
    DeleteSchemaVersionAsync(ctx workflow.Context, input *schemas.DeleteSchemaVersionInput) *SchemasDeleteSchemaVersionResult

    DescribeCodeBinding(ctx workflow.Context, input *schemas.DescribeCodeBindingInput) (*schemas.DescribeCodeBindingOutput, error)
    DescribeCodeBindingAsync(ctx workflow.Context, input *schemas.DescribeCodeBindingInput) *SchemasDescribeCodeBindingResult

    DescribeDiscoverer(ctx workflow.Context, input *schemas.DescribeDiscovererInput) (*schemas.DescribeDiscovererOutput, error)
    DescribeDiscovererAsync(ctx workflow.Context, input *schemas.DescribeDiscovererInput) *SchemasDescribeDiscovererResult

    DescribeRegistry(ctx workflow.Context, input *schemas.DescribeRegistryInput) (*schemas.DescribeRegistryOutput, error)
    DescribeRegistryAsync(ctx workflow.Context, input *schemas.DescribeRegistryInput) *SchemasDescribeRegistryResult

    DescribeSchema(ctx workflow.Context, input *schemas.DescribeSchemaInput) (*schemas.DescribeSchemaOutput, error)
    DescribeSchemaAsync(ctx workflow.Context, input *schemas.DescribeSchemaInput) *SchemasDescribeSchemaResult

    GetCodeBindingSource(ctx workflow.Context, input *schemas.GetCodeBindingSourceInput) (*schemas.GetCodeBindingSourceOutput, error)
    GetCodeBindingSourceAsync(ctx workflow.Context, input *schemas.GetCodeBindingSourceInput) *SchemasGetCodeBindingSourceResult

    GetDiscoveredSchema(ctx workflow.Context, input *schemas.GetDiscoveredSchemaInput) (*schemas.GetDiscoveredSchemaOutput, error)
    GetDiscoveredSchemaAsync(ctx workflow.Context, input *schemas.GetDiscoveredSchemaInput) *SchemasGetDiscoveredSchemaResult

    GetResourcePolicy(ctx workflow.Context, input *schemas.GetResourcePolicyInput) (*schemas.GetResourcePolicyOutput, error)
    GetResourcePolicyAsync(ctx workflow.Context, input *schemas.GetResourcePolicyInput) *SchemasGetResourcePolicyResult

    ListDiscoverers(ctx workflow.Context, input *schemas.ListDiscoverersInput) (*schemas.ListDiscoverersOutput, error)
    ListDiscoverersAsync(ctx workflow.Context, input *schemas.ListDiscoverersInput) *SchemasListDiscoverersResult

    ListRegistries(ctx workflow.Context, input *schemas.ListRegistriesInput) (*schemas.ListRegistriesOutput, error)
    ListRegistriesAsync(ctx workflow.Context, input *schemas.ListRegistriesInput) *SchemasListRegistriesResult

    ListSchemaVersions(ctx workflow.Context, input *schemas.ListSchemaVersionsInput) (*schemas.ListSchemaVersionsOutput, error)
    ListSchemaVersionsAsync(ctx workflow.Context, input *schemas.ListSchemaVersionsInput) *SchemasListSchemaVersionsResult

    ListSchemas(ctx workflow.Context, input *schemas.ListSchemasInput) (*schemas.ListSchemasOutput, error)
    ListSchemasAsync(ctx workflow.Context, input *schemas.ListSchemasInput) *SchemasListSchemasResult

    ListTagsForResource(ctx workflow.Context, input *schemas.ListTagsForResourceInput) (*schemas.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *schemas.ListTagsForResourceInput) *SchemasListTagsForResourceResult

    PutCodeBinding(ctx workflow.Context, input *schemas.PutCodeBindingInput) (*schemas.PutCodeBindingOutput, error)
    PutCodeBindingAsync(ctx workflow.Context, input *schemas.PutCodeBindingInput) *SchemasPutCodeBindingResult

    PutResourcePolicy(ctx workflow.Context, input *schemas.PutResourcePolicyInput) (*schemas.PutResourcePolicyOutput, error)
    PutResourcePolicyAsync(ctx workflow.Context, input *schemas.PutResourcePolicyInput) *SchemasPutResourcePolicyResult

    SearchSchemas(ctx workflow.Context, input *schemas.SearchSchemasInput) (*schemas.SearchSchemasOutput, error)
    SearchSchemasAsync(ctx workflow.Context, input *schemas.SearchSchemasInput) *SchemasSearchSchemasResult

    StartDiscoverer(ctx workflow.Context, input *schemas.StartDiscovererInput) (*schemas.StartDiscovererOutput, error)
    StartDiscovererAsync(ctx workflow.Context, input *schemas.StartDiscovererInput) *SchemasStartDiscovererResult

    StopDiscoverer(ctx workflow.Context, input *schemas.StopDiscovererInput) (*schemas.StopDiscovererOutput, error)
    StopDiscovererAsync(ctx workflow.Context, input *schemas.StopDiscovererInput) *SchemasStopDiscovererResult

    TagResource(ctx workflow.Context, input *schemas.TagResourceInput) (*schemas.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *schemas.TagResourceInput) *SchemasTagResourceResult

    UntagResource(ctx workflow.Context, input *schemas.UntagResourceInput) (*schemas.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *schemas.UntagResourceInput) *SchemasUntagResourceResult

    UpdateDiscoverer(ctx workflow.Context, input *schemas.UpdateDiscovererInput) (*schemas.UpdateDiscovererOutput, error)
    UpdateDiscovererAsync(ctx workflow.Context, input *schemas.UpdateDiscovererInput) *SchemasUpdateDiscovererResult

    UpdateRegistry(ctx workflow.Context, input *schemas.UpdateRegistryInput) (*schemas.UpdateRegistryOutput, error)
    UpdateRegistryAsync(ctx workflow.Context, input *schemas.UpdateRegistryInput) *SchemasUpdateRegistryResult

    UpdateSchema(ctx workflow.Context, input *schemas.UpdateSchemaInput) (*schemas.UpdateSchemaOutput, error)
    UpdateSchemaAsync(ctx workflow.Context, input *schemas.UpdateSchemaInput) *SchemasUpdateSchemaResult

    WaitUntilCodeBindingExists(ctx workflow.Context, input *schemas.DescribeCodeBindingInput) error}
type SchemasCreateDiscovererResult struct {
	Result workflow.Future
}

func (r *SchemasCreateDiscovererResult) Get(ctx workflow.Context) (*schemas.CreateDiscovererOutput, error) {
    var output schemas.CreateDiscovererOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasCreateRegistryResult struct {
	Result workflow.Future
}

func (r *SchemasCreateRegistryResult) Get(ctx workflow.Context) (*schemas.CreateRegistryOutput, error) {
    var output schemas.CreateRegistryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasCreateSchemaResult struct {
	Result workflow.Future
}

func (r *SchemasCreateSchemaResult) Get(ctx workflow.Context) (*schemas.CreateSchemaOutput, error) {
    var output schemas.CreateSchemaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasDeleteDiscovererResult struct {
	Result workflow.Future
}

func (r *SchemasDeleteDiscovererResult) Get(ctx workflow.Context) (*schemas.DeleteDiscovererOutput, error) {
    var output schemas.DeleteDiscovererOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasDeleteRegistryResult struct {
	Result workflow.Future
}

func (r *SchemasDeleteRegistryResult) Get(ctx workflow.Context) (*schemas.DeleteRegistryOutput, error) {
    var output schemas.DeleteRegistryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasDeleteResourcePolicyResult struct {
	Result workflow.Future
}

func (r *SchemasDeleteResourcePolicyResult) Get(ctx workflow.Context) (*schemas.DeleteResourcePolicyOutput, error) {
    var output schemas.DeleteResourcePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasDeleteSchemaResult struct {
	Result workflow.Future
}

func (r *SchemasDeleteSchemaResult) Get(ctx workflow.Context) (*schemas.DeleteSchemaOutput, error) {
    var output schemas.DeleteSchemaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasDeleteSchemaVersionResult struct {
	Result workflow.Future
}

func (r *SchemasDeleteSchemaVersionResult) Get(ctx workflow.Context) (*schemas.DeleteSchemaVersionOutput, error) {
    var output schemas.DeleteSchemaVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasDescribeCodeBindingResult struct {
	Result workflow.Future
}

func (r *SchemasDescribeCodeBindingResult) Get(ctx workflow.Context) (*schemas.DescribeCodeBindingOutput, error) {
    var output schemas.DescribeCodeBindingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasDescribeDiscovererResult struct {
	Result workflow.Future
}

func (r *SchemasDescribeDiscovererResult) Get(ctx workflow.Context) (*schemas.DescribeDiscovererOutput, error) {
    var output schemas.DescribeDiscovererOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasDescribeRegistryResult struct {
	Result workflow.Future
}

func (r *SchemasDescribeRegistryResult) Get(ctx workflow.Context) (*schemas.DescribeRegistryOutput, error) {
    var output schemas.DescribeRegistryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasDescribeSchemaResult struct {
	Result workflow.Future
}

func (r *SchemasDescribeSchemaResult) Get(ctx workflow.Context) (*schemas.DescribeSchemaOutput, error) {
    var output schemas.DescribeSchemaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasGetCodeBindingSourceResult struct {
	Result workflow.Future
}

func (r *SchemasGetCodeBindingSourceResult) Get(ctx workflow.Context) (*schemas.GetCodeBindingSourceOutput, error) {
    var output schemas.GetCodeBindingSourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasGetDiscoveredSchemaResult struct {
	Result workflow.Future
}

func (r *SchemasGetDiscoveredSchemaResult) Get(ctx workflow.Context) (*schemas.GetDiscoveredSchemaOutput, error) {
    var output schemas.GetDiscoveredSchemaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasGetResourcePolicyResult struct {
	Result workflow.Future
}

func (r *SchemasGetResourcePolicyResult) Get(ctx workflow.Context) (*schemas.GetResourcePolicyOutput, error) {
    var output schemas.GetResourcePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasListDiscoverersResult struct {
	Result workflow.Future
}

func (r *SchemasListDiscoverersResult) Get(ctx workflow.Context) (*schemas.ListDiscoverersOutput, error) {
    var output schemas.ListDiscoverersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasListRegistriesResult struct {
	Result workflow.Future
}

func (r *SchemasListRegistriesResult) Get(ctx workflow.Context) (*schemas.ListRegistriesOutput, error) {
    var output schemas.ListRegistriesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasListSchemaVersionsResult struct {
	Result workflow.Future
}

func (r *SchemasListSchemaVersionsResult) Get(ctx workflow.Context) (*schemas.ListSchemaVersionsOutput, error) {
    var output schemas.ListSchemaVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasListSchemasResult struct {
	Result workflow.Future
}

func (r *SchemasListSchemasResult) Get(ctx workflow.Context) (*schemas.ListSchemasOutput, error) {
    var output schemas.ListSchemasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *SchemasListTagsForResourceResult) Get(ctx workflow.Context) (*schemas.ListTagsForResourceOutput, error) {
    var output schemas.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasPutCodeBindingResult struct {
	Result workflow.Future
}

func (r *SchemasPutCodeBindingResult) Get(ctx workflow.Context) (*schemas.PutCodeBindingOutput, error) {
    var output schemas.PutCodeBindingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasPutResourcePolicyResult struct {
	Result workflow.Future
}

func (r *SchemasPutResourcePolicyResult) Get(ctx workflow.Context) (*schemas.PutResourcePolicyOutput, error) {
    var output schemas.PutResourcePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasSearchSchemasResult struct {
	Result workflow.Future
}

func (r *SchemasSearchSchemasResult) Get(ctx workflow.Context) (*schemas.SearchSchemasOutput, error) {
    var output schemas.SearchSchemasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasStartDiscovererResult struct {
	Result workflow.Future
}

func (r *SchemasStartDiscovererResult) Get(ctx workflow.Context) (*schemas.StartDiscovererOutput, error) {
    var output schemas.StartDiscovererOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasStopDiscovererResult struct {
	Result workflow.Future
}

func (r *SchemasStopDiscovererResult) Get(ctx workflow.Context) (*schemas.StopDiscovererOutput, error) {
    var output schemas.StopDiscovererOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasTagResourceResult struct {
	Result workflow.Future
}

func (r *SchemasTagResourceResult) Get(ctx workflow.Context) (*schemas.TagResourceOutput, error) {
    var output schemas.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasUntagResourceResult struct {
	Result workflow.Future
}

func (r *SchemasUntagResourceResult) Get(ctx workflow.Context) (*schemas.UntagResourceOutput, error) {
    var output schemas.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasUpdateDiscovererResult struct {
	Result workflow.Future
}

func (r *SchemasUpdateDiscovererResult) Get(ctx workflow.Context) (*schemas.UpdateDiscovererOutput, error) {
    var output schemas.UpdateDiscovererOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasUpdateRegistryResult struct {
	Result workflow.Future
}

func (r *SchemasUpdateRegistryResult) Get(ctx workflow.Context) (*schemas.UpdateRegistryOutput, error) {
    var output schemas.UpdateRegistryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SchemasUpdateSchemaResult struct {
	Result workflow.Future
}

func (r *SchemasUpdateSchemaResult) Get(ctx workflow.Context) (*schemas.UpdateSchemaOutput, error) {
    var output schemas.UpdateSchemaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type SchemasStub struct {
    activities SchemasClient
}

func NewSchemasStub() SchemasClient {
    return &SchemasStub{}
}

func (a *SchemasStub) CreateDiscoverer(ctx workflow.Context, input *schemas.CreateDiscovererInput) (*schemas.CreateDiscovererOutput, error) {
    var output schemas.CreateDiscovererOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDiscoverer, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) CreateRegistry(ctx workflow.Context, input *schemas.CreateRegistryInput) (*schemas.CreateRegistryOutput, error) {
    var output schemas.CreateRegistryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateRegistry, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) CreateSchema(ctx workflow.Context, input *schemas.CreateSchemaInput) (*schemas.CreateSchemaOutput, error) {
    var output schemas.CreateSchemaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSchema, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) DeleteDiscoverer(ctx workflow.Context, input *schemas.DeleteDiscovererInput) (*schemas.DeleteDiscovererOutput, error) {
    var output schemas.DeleteDiscovererOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDiscoverer, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) DeleteRegistry(ctx workflow.Context, input *schemas.DeleteRegistryInput) (*schemas.DeleteRegistryOutput, error) {
    var output schemas.DeleteRegistryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRegistry, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) DeleteResourcePolicy(ctx workflow.Context, input *schemas.DeleteResourcePolicyInput) (*schemas.DeleteResourcePolicyOutput, error) {
    var output schemas.DeleteResourcePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteResourcePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) DeleteSchema(ctx workflow.Context, input *schemas.DeleteSchemaInput) (*schemas.DeleteSchemaOutput, error) {
    var output schemas.DeleteSchemaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSchema, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) DeleteSchemaVersion(ctx workflow.Context, input *schemas.DeleteSchemaVersionInput) (*schemas.DeleteSchemaVersionOutput, error) {
    var output schemas.DeleteSchemaVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSchemaVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) DescribeCodeBinding(ctx workflow.Context, input *schemas.DescribeCodeBindingInput) (*schemas.DescribeCodeBindingOutput, error) {
    var output schemas.DescribeCodeBindingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCodeBinding, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) DescribeDiscoverer(ctx workflow.Context, input *schemas.DescribeDiscovererInput) (*schemas.DescribeDiscovererOutput, error) {
    var output schemas.DescribeDiscovererOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDiscoverer, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) DescribeRegistry(ctx workflow.Context, input *schemas.DescribeRegistryInput) (*schemas.DescribeRegistryOutput, error) {
    var output schemas.DescribeRegistryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRegistry, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) DescribeSchema(ctx workflow.Context, input *schemas.DescribeSchemaInput) (*schemas.DescribeSchemaOutput, error) {
    var output schemas.DescribeSchemaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSchema, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) GetCodeBindingSource(ctx workflow.Context, input *schemas.GetCodeBindingSourceInput) (*schemas.GetCodeBindingSourceOutput, error) {
    var output schemas.GetCodeBindingSourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCodeBindingSource, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) GetDiscoveredSchema(ctx workflow.Context, input *schemas.GetDiscoveredSchemaInput) (*schemas.GetDiscoveredSchemaOutput, error) {
    var output schemas.GetDiscoveredSchemaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDiscoveredSchema, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) GetResourcePolicy(ctx workflow.Context, input *schemas.GetResourcePolicyInput) (*schemas.GetResourcePolicyOutput, error) {
    var output schemas.GetResourcePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetResourcePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) ListDiscoverers(ctx workflow.Context, input *schemas.ListDiscoverersInput) (*schemas.ListDiscoverersOutput, error) {
    var output schemas.ListDiscoverersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDiscoverers, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) ListRegistries(ctx workflow.Context, input *schemas.ListRegistriesInput) (*schemas.ListRegistriesOutput, error) {
    var output schemas.ListRegistriesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRegistries, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) ListSchemaVersions(ctx workflow.Context, input *schemas.ListSchemaVersionsInput) (*schemas.ListSchemaVersionsOutput, error) {
    var output schemas.ListSchemaVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSchemaVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) ListSchemas(ctx workflow.Context, input *schemas.ListSchemasInput) (*schemas.ListSchemasOutput, error) {
    var output schemas.ListSchemasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSchemas, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) ListTagsForResource(ctx workflow.Context, input *schemas.ListTagsForResourceInput) (*schemas.ListTagsForResourceOutput, error) {
    var output schemas.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) PutCodeBinding(ctx workflow.Context, input *schemas.PutCodeBindingInput) (*schemas.PutCodeBindingOutput, error) {
    var output schemas.PutCodeBindingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutCodeBinding, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) PutResourcePolicy(ctx workflow.Context, input *schemas.PutResourcePolicyInput) (*schemas.PutResourcePolicyOutput, error) {
    var output schemas.PutResourcePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutResourcePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) SearchSchemas(ctx workflow.Context, input *schemas.SearchSchemasInput) (*schemas.SearchSchemasOutput, error) {
    var output schemas.SearchSchemasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SearchSchemas, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) StartDiscoverer(ctx workflow.Context, input *schemas.StartDiscovererInput) (*schemas.StartDiscovererOutput, error) {
    var output schemas.StartDiscovererOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartDiscoverer, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) StopDiscoverer(ctx workflow.Context, input *schemas.StopDiscovererInput) (*schemas.StopDiscovererOutput, error) {
    var output schemas.StopDiscovererOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopDiscoverer, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) TagResource(ctx workflow.Context, input *schemas.TagResourceInput) (*schemas.TagResourceOutput, error) {
    var output schemas.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) UntagResource(ctx workflow.Context, input *schemas.UntagResourceInput) (*schemas.UntagResourceOutput, error) {
    var output schemas.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) UpdateDiscoverer(ctx workflow.Context, input *schemas.UpdateDiscovererInput) (*schemas.UpdateDiscovererOutput, error) {
    var output schemas.UpdateDiscovererOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDiscoverer, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) UpdateRegistry(ctx workflow.Context, input *schemas.UpdateRegistryInput) (*schemas.UpdateRegistryOutput, error) {
    var output schemas.UpdateRegistryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateRegistry, input).Get(ctx, &output)
    return &output, err
}

func (a *SchemasStub) UpdateSchema(ctx workflow.Context, input *schemas.UpdateSchemaInput) (*schemas.UpdateSchemaOutput, error) {
    var output schemas.UpdateSchemaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSchema, input).Get(ctx, &output)
    return &output, err
}


func (a *SchemasStub) WaitUntilCodeBindingExists(ctx workflow.Context, input *schemas.DescribeCodeBindingInput) error {
    return a.client.WaitUntilCodeBindingExists(input)
}
