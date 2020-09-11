package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/schemas"
	"github.com/aws/aws-sdk-go/service/schemas/schemasiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type SchemasActivities struct {
	client schemasiface.SchemasAPI
}

func NewSchemasActivities(session *session.Session, config ...*aws.Config) *SchemasActivities {
	client := schemas.New(session, config...)
	return &SchemasActivities{client: client}
}

func (a *SchemasActivities) CreateDiscoverer(ctx context.Context, input *schemas.CreateDiscovererInput) (*schemas.CreateDiscovererOutput, error) {
	return a.client.CreateDiscovererWithContext(ctx, input)
}

func (a *SchemasActivities) CreateRegistry(ctx context.Context, input *schemas.CreateRegistryInput) (*schemas.CreateRegistryOutput, error) {
	return a.client.CreateRegistryWithContext(ctx, input)
}

func (a *SchemasActivities) CreateSchema(ctx context.Context, input *schemas.CreateSchemaInput) (*schemas.CreateSchemaOutput, error) {
	return a.client.CreateSchemaWithContext(ctx, input)
}

func (a *SchemasActivities) DeleteDiscoverer(ctx context.Context, input *schemas.DeleteDiscovererInput) (*schemas.DeleteDiscovererOutput, error) {
	return a.client.DeleteDiscovererWithContext(ctx, input)
}

func (a *SchemasActivities) DeleteRegistry(ctx context.Context, input *schemas.DeleteRegistryInput) (*schemas.DeleteRegistryOutput, error) {
	return a.client.DeleteRegistryWithContext(ctx, input)
}

func (a *SchemasActivities) DeleteResourcePolicy(ctx context.Context, input *schemas.DeleteResourcePolicyInput) (*schemas.DeleteResourcePolicyOutput, error) {
	return a.client.DeleteResourcePolicyWithContext(ctx, input)
}

func (a *SchemasActivities) DeleteSchema(ctx context.Context, input *schemas.DeleteSchemaInput) (*schemas.DeleteSchemaOutput, error) {
	return a.client.DeleteSchemaWithContext(ctx, input)
}

func (a *SchemasActivities) DeleteSchemaVersion(ctx context.Context, input *schemas.DeleteSchemaVersionInput) (*schemas.DeleteSchemaVersionOutput, error) {
	return a.client.DeleteSchemaVersionWithContext(ctx, input)
}

func (a *SchemasActivities) DescribeCodeBinding(ctx context.Context, input *schemas.DescribeCodeBindingInput) (*schemas.DescribeCodeBindingOutput, error) {
	return a.client.DescribeCodeBindingWithContext(ctx, input)
}

func (a *SchemasActivities) DescribeDiscoverer(ctx context.Context, input *schemas.DescribeDiscovererInput) (*schemas.DescribeDiscovererOutput, error) {
	return a.client.DescribeDiscovererWithContext(ctx, input)
}

func (a *SchemasActivities) DescribeRegistry(ctx context.Context, input *schemas.DescribeRegistryInput) (*schemas.DescribeRegistryOutput, error) {
	return a.client.DescribeRegistryWithContext(ctx, input)
}

func (a *SchemasActivities) DescribeSchema(ctx context.Context, input *schemas.DescribeSchemaInput) (*schemas.DescribeSchemaOutput, error) {
	return a.client.DescribeSchemaWithContext(ctx, input)
}

func (a *SchemasActivities) GetCodeBindingSource(ctx context.Context, input *schemas.GetCodeBindingSourceInput) (*schemas.GetCodeBindingSourceOutput, error) {
	return a.client.GetCodeBindingSourceWithContext(ctx, input)
}

func (a *SchemasActivities) GetDiscoveredSchema(ctx context.Context, input *schemas.GetDiscoveredSchemaInput) (*schemas.GetDiscoveredSchemaOutput, error) {
	return a.client.GetDiscoveredSchemaWithContext(ctx, input)
}

func (a *SchemasActivities) GetResourcePolicy(ctx context.Context, input *schemas.GetResourcePolicyInput) (*schemas.GetResourcePolicyOutput, error) {
	return a.client.GetResourcePolicyWithContext(ctx, input)
}

func (a *SchemasActivities) ListDiscoverers(ctx context.Context, input *schemas.ListDiscoverersInput) (*schemas.ListDiscoverersOutput, error) {
	return a.client.ListDiscoverersWithContext(ctx, input)
}

func (a *SchemasActivities) ListRegistries(ctx context.Context, input *schemas.ListRegistriesInput) (*schemas.ListRegistriesOutput, error) {
	return a.client.ListRegistriesWithContext(ctx, input)
}

func (a *SchemasActivities) ListSchemaVersions(ctx context.Context, input *schemas.ListSchemaVersionsInput) (*schemas.ListSchemaVersionsOutput, error) {
	return a.client.ListSchemaVersionsWithContext(ctx, input)
}

func (a *SchemasActivities) ListSchemas(ctx context.Context, input *schemas.ListSchemasInput) (*schemas.ListSchemasOutput, error) {
	return a.client.ListSchemasWithContext(ctx, input)
}

func (a *SchemasActivities) ListTagsForResource(ctx context.Context, input *schemas.ListTagsForResourceInput) (*schemas.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *SchemasActivities) PutCodeBinding(ctx context.Context, input *schemas.PutCodeBindingInput) (*schemas.PutCodeBindingOutput, error) {
	return a.client.PutCodeBindingWithContext(ctx, input)
}

func (a *SchemasActivities) PutResourcePolicy(ctx context.Context, input *schemas.PutResourcePolicyInput) (*schemas.PutResourcePolicyOutput, error) {
	return a.client.PutResourcePolicyWithContext(ctx, input)
}

func (a *SchemasActivities) SearchSchemas(ctx context.Context, input *schemas.SearchSchemasInput) (*schemas.SearchSchemasOutput, error) {
	return a.client.SearchSchemasWithContext(ctx, input)
}

func (a *SchemasActivities) StartDiscoverer(ctx context.Context, input *schemas.StartDiscovererInput) (*schemas.StartDiscovererOutput, error) {
	return a.client.StartDiscovererWithContext(ctx, input)
}

func (a *SchemasActivities) StopDiscoverer(ctx context.Context, input *schemas.StopDiscovererInput) (*schemas.StopDiscovererOutput, error) {
	return a.client.StopDiscovererWithContext(ctx, input)
}

func (a *SchemasActivities) TagResource(ctx context.Context, input *schemas.TagResourceInput) (*schemas.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *SchemasActivities) UntagResource(ctx context.Context, input *schemas.UntagResourceInput) (*schemas.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *SchemasActivities) UpdateDiscoverer(ctx context.Context, input *schemas.UpdateDiscovererInput) (*schemas.UpdateDiscovererOutput, error) {
	return a.client.UpdateDiscovererWithContext(ctx, input)
}

func (a *SchemasActivities) UpdateRegistry(ctx context.Context, input *schemas.UpdateRegistryInput) (*schemas.UpdateRegistryOutput, error) {
	return a.client.UpdateRegistryWithContext(ctx, input)
}

func (a *SchemasActivities) UpdateSchema(ctx context.Context, input *schemas.UpdateSchemaInput) (*schemas.UpdateSchemaOutput, error) {
	return a.client.UpdateSchemaWithContext(ctx, input)
}

func (a *SchemasActivities) WaitUntilCodeBindingExists(ctx context.Context, input *schemas.DescribeCodeBindingInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilCodeBindingExistsWithContext(ctx, input, options...)
	})
}
