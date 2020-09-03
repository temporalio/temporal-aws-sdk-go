package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/schemas"
	"github.com/aws/aws-sdk-go/service/schemas/schemasiface"
)

type SchemasActivities struct {
	client schemasiface.SchemasAPI
}

func NewSchemasActivities(client schemasiface.SchemasAPI) *SchemasActivities {
    return &SchemasActivities{client: client}
}

func (a *SchemasActivities) CreateDiscoverer(input *schemas.CreateDiscovererInput) (*schemas.CreateDiscovererOutput, error) {
    return a.client.CreateDiscoverer(input)
}

func (a *SchemasActivities) CreateRegistry(input *schemas.CreateRegistryInput) (*schemas.CreateRegistryOutput, error) {
    return a.client.CreateRegistry(input)
}

func (a *SchemasActivities) CreateSchema(input *schemas.CreateSchemaInput) (*schemas.CreateSchemaOutput, error) {
    return a.client.CreateSchema(input)
}

func (a *SchemasActivities) DeleteDiscoverer(input *schemas.DeleteDiscovererInput) (*schemas.DeleteDiscovererOutput, error) {
    return a.client.DeleteDiscoverer(input)
}

func (a *SchemasActivities) DeleteRegistry(input *schemas.DeleteRegistryInput) (*schemas.DeleteRegistryOutput, error) {
    return a.client.DeleteRegistry(input)
}

func (a *SchemasActivities) DeleteResourcePolicy(input *schemas.DeleteResourcePolicyInput) (*schemas.DeleteResourcePolicyOutput, error) {
    return a.client.DeleteResourcePolicy(input)
}

func (a *SchemasActivities) DeleteSchema(input *schemas.DeleteSchemaInput) (*schemas.DeleteSchemaOutput, error) {
    return a.client.DeleteSchema(input)
}

func (a *SchemasActivities) DeleteSchemaVersion(input *schemas.DeleteSchemaVersionInput) (*schemas.DeleteSchemaVersionOutput, error) {
    return a.client.DeleteSchemaVersion(input)
}

func (a *SchemasActivities) DescribeCodeBinding(input *schemas.DescribeCodeBindingInput) (*schemas.DescribeCodeBindingOutput, error) {
    return a.client.DescribeCodeBinding(input)
}

func (a *SchemasActivities) DescribeDiscoverer(input *schemas.DescribeDiscovererInput) (*schemas.DescribeDiscovererOutput, error) {
    return a.client.DescribeDiscoverer(input)
}

func (a *SchemasActivities) DescribeRegistry(input *schemas.DescribeRegistryInput) (*schemas.DescribeRegistryOutput, error) {
    return a.client.DescribeRegistry(input)
}

func (a *SchemasActivities) DescribeSchema(input *schemas.DescribeSchemaInput) (*schemas.DescribeSchemaOutput, error) {
    return a.client.DescribeSchema(input)
}

func (a *SchemasActivities) GetCodeBindingSource(input *schemas.GetCodeBindingSourceInput) (*schemas.GetCodeBindingSourceOutput, error) {
    return a.client.GetCodeBindingSource(input)
}

func (a *SchemasActivities) GetDiscoveredSchema(input *schemas.GetDiscoveredSchemaInput) (*schemas.GetDiscoveredSchemaOutput, error) {
    return a.client.GetDiscoveredSchema(input)
}

func (a *SchemasActivities) GetResourcePolicy(input *schemas.GetResourcePolicyInput) (*schemas.GetResourcePolicyOutput, error) {
    return a.client.GetResourcePolicy(input)
}

func (a *SchemasActivities) ListDiscoverers(input *schemas.ListDiscoverersInput) (*schemas.ListDiscoverersOutput, error) {
    return a.client.ListDiscoverers(input)
}

func (a *SchemasActivities) ListRegistries(input *schemas.ListRegistriesInput) (*schemas.ListRegistriesOutput, error) {
    return a.client.ListRegistries(input)
}

func (a *SchemasActivities) ListSchemaVersions(input *schemas.ListSchemaVersionsInput) (*schemas.ListSchemaVersionsOutput, error) {
    return a.client.ListSchemaVersions(input)
}

func (a *SchemasActivities) ListSchemas(input *schemas.ListSchemasInput) (*schemas.ListSchemasOutput, error) {
    return a.client.ListSchemas(input)
}

func (a *SchemasActivities) ListTagsForResource(input *schemas.ListTagsForResourceInput) (*schemas.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *SchemasActivities) PutCodeBinding(input *schemas.PutCodeBindingInput) (*schemas.PutCodeBindingOutput, error) {
    return a.client.PutCodeBinding(input)
}

func (a *SchemasActivities) PutResourcePolicy(input *schemas.PutResourcePolicyInput) (*schemas.PutResourcePolicyOutput, error) {
    return a.client.PutResourcePolicy(input)
}

func (a *SchemasActivities) SearchSchemas(input *schemas.SearchSchemasInput) (*schemas.SearchSchemasOutput, error) {
    return a.client.SearchSchemas(input)
}

func (a *SchemasActivities) StartDiscoverer(input *schemas.StartDiscovererInput) (*schemas.StartDiscovererOutput, error) {
    return a.client.StartDiscoverer(input)
}

func (a *SchemasActivities) StopDiscoverer(input *schemas.StopDiscovererInput) (*schemas.StopDiscovererOutput, error) {
    return a.client.StopDiscoverer(input)
}

func (a *SchemasActivities) TagResource(input *schemas.TagResourceInput) (*schemas.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *SchemasActivities) UntagResource(input *schemas.UntagResourceInput) (*schemas.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *SchemasActivities) UpdateDiscoverer(input *schemas.UpdateDiscovererInput) (*schemas.UpdateDiscovererOutput, error) {
    return a.client.UpdateDiscoverer(input)
}

func (a *SchemasActivities) UpdateRegistry(input *schemas.UpdateRegistryInput) (*schemas.UpdateRegistryOutput, error) {
    return a.client.UpdateRegistry(input)
}

func (a *SchemasActivities) UpdateSchema(input *schemas.UpdateSchemaInput) (*schemas.UpdateSchemaOutput, error) {
    return a.client.UpdateSchema(input)
}

func (a *SchemasActivities) WaitUntilCodeBindingExists(input *schemas.WaitUntilCodeBindingExistsInput) (*schemas.WaitUntilCodeBindingExistsOutput, error) {
    return a.client.WaitUntilCodeBindingExists(input)
}
