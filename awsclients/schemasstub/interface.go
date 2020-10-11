// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package schemasstub

import (
	"github.com/aws/aws-sdk-go/service/schemas"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	CreateDiscoverer(ctx workflow.Context, input *schemas.CreateDiscovererInput) (*schemas.CreateDiscovererOutput, error)
	CreateDiscovererAsync(ctx workflow.Context, input *schemas.CreateDiscovererInput) *SchemasCreateDiscovererFuture

	CreateRegistry(ctx workflow.Context, input *schemas.CreateRegistryInput) (*schemas.CreateRegistryOutput, error)
	CreateRegistryAsync(ctx workflow.Context, input *schemas.CreateRegistryInput) *SchemasCreateRegistryFuture

	CreateSchema(ctx workflow.Context, input *schemas.CreateSchemaInput) (*schemas.CreateSchemaOutput, error)
	CreateSchemaAsync(ctx workflow.Context, input *schemas.CreateSchemaInput) *SchemasCreateSchemaFuture

	DeleteDiscoverer(ctx workflow.Context, input *schemas.DeleteDiscovererInput) (*schemas.DeleteDiscovererOutput, error)
	DeleteDiscovererAsync(ctx workflow.Context, input *schemas.DeleteDiscovererInput) *SchemasDeleteDiscovererFuture

	DeleteRegistry(ctx workflow.Context, input *schemas.DeleteRegistryInput) (*schemas.DeleteRegistryOutput, error)
	DeleteRegistryAsync(ctx workflow.Context, input *schemas.DeleteRegistryInput) *SchemasDeleteRegistryFuture

	DeleteResourcePolicy(ctx workflow.Context, input *schemas.DeleteResourcePolicyInput) (*schemas.DeleteResourcePolicyOutput, error)
	DeleteResourcePolicyAsync(ctx workflow.Context, input *schemas.DeleteResourcePolicyInput) *SchemasDeleteResourcePolicyFuture

	DeleteSchema(ctx workflow.Context, input *schemas.DeleteSchemaInput) (*schemas.DeleteSchemaOutput, error)
	DeleteSchemaAsync(ctx workflow.Context, input *schemas.DeleteSchemaInput) *SchemasDeleteSchemaFuture

	DeleteSchemaVersion(ctx workflow.Context, input *schemas.DeleteSchemaVersionInput) (*schemas.DeleteSchemaVersionOutput, error)
	DeleteSchemaVersionAsync(ctx workflow.Context, input *schemas.DeleteSchemaVersionInput) *SchemasDeleteSchemaVersionFuture

	DescribeCodeBinding(ctx workflow.Context, input *schemas.DescribeCodeBindingInput) (*schemas.DescribeCodeBindingOutput, error)
	DescribeCodeBindingAsync(ctx workflow.Context, input *schemas.DescribeCodeBindingInput) *SchemasDescribeCodeBindingFuture

	DescribeDiscoverer(ctx workflow.Context, input *schemas.DescribeDiscovererInput) (*schemas.DescribeDiscovererOutput, error)
	DescribeDiscovererAsync(ctx workflow.Context, input *schemas.DescribeDiscovererInput) *SchemasDescribeDiscovererFuture

	DescribeRegistry(ctx workflow.Context, input *schemas.DescribeRegistryInput) (*schemas.DescribeRegistryOutput, error)
	DescribeRegistryAsync(ctx workflow.Context, input *schemas.DescribeRegistryInput) *SchemasDescribeRegistryFuture

	DescribeSchema(ctx workflow.Context, input *schemas.DescribeSchemaInput) (*schemas.DescribeSchemaOutput, error)
	DescribeSchemaAsync(ctx workflow.Context, input *schemas.DescribeSchemaInput) *SchemasDescribeSchemaFuture

	ExportSchema(ctx workflow.Context, input *schemas.ExportSchemaInput) (*schemas.ExportSchemaOutput, error)
	ExportSchemaAsync(ctx workflow.Context, input *schemas.ExportSchemaInput) *SchemasExportSchemaFuture

	GetCodeBindingSource(ctx workflow.Context, input *schemas.GetCodeBindingSourceInput) (*schemas.GetCodeBindingSourceOutput, error)
	GetCodeBindingSourceAsync(ctx workflow.Context, input *schemas.GetCodeBindingSourceInput) *SchemasGetCodeBindingSourceFuture

	GetDiscoveredSchema(ctx workflow.Context, input *schemas.GetDiscoveredSchemaInput) (*schemas.GetDiscoveredSchemaOutput, error)
	GetDiscoveredSchemaAsync(ctx workflow.Context, input *schemas.GetDiscoveredSchemaInput) *SchemasGetDiscoveredSchemaFuture

	GetResourcePolicy(ctx workflow.Context, input *schemas.GetResourcePolicyInput) (*schemas.GetResourcePolicyOutput, error)
	GetResourcePolicyAsync(ctx workflow.Context, input *schemas.GetResourcePolicyInput) *SchemasGetResourcePolicyFuture

	ListDiscoverers(ctx workflow.Context, input *schemas.ListDiscoverersInput) (*schemas.ListDiscoverersOutput, error)
	ListDiscoverersAsync(ctx workflow.Context, input *schemas.ListDiscoverersInput) *SchemasListDiscoverersFuture

	ListRegistries(ctx workflow.Context, input *schemas.ListRegistriesInput) (*schemas.ListRegistriesOutput, error)
	ListRegistriesAsync(ctx workflow.Context, input *schemas.ListRegistriesInput) *SchemasListRegistriesFuture

	ListSchemaVersions(ctx workflow.Context, input *schemas.ListSchemaVersionsInput) (*schemas.ListSchemaVersionsOutput, error)
	ListSchemaVersionsAsync(ctx workflow.Context, input *schemas.ListSchemaVersionsInput) *SchemasListSchemaVersionsFuture

	ListSchemas(ctx workflow.Context, input *schemas.ListSchemasInput) (*schemas.ListSchemasOutput, error)
	ListSchemasAsync(ctx workflow.Context, input *schemas.ListSchemasInput) *SchemasListSchemasFuture

	ListTagsForResource(ctx workflow.Context, input *schemas.ListTagsForResourceInput) (*schemas.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *schemas.ListTagsForResourceInput) *SchemasListTagsForResourceFuture

	PutCodeBinding(ctx workflow.Context, input *schemas.PutCodeBindingInput) (*schemas.PutCodeBindingOutput, error)
	PutCodeBindingAsync(ctx workflow.Context, input *schemas.PutCodeBindingInput) *SchemasPutCodeBindingFuture

	PutResourcePolicy(ctx workflow.Context, input *schemas.PutResourcePolicyInput) (*schemas.PutResourcePolicyOutput, error)
	PutResourcePolicyAsync(ctx workflow.Context, input *schemas.PutResourcePolicyInput) *SchemasPutResourcePolicyFuture

	SearchSchemas(ctx workflow.Context, input *schemas.SearchSchemasInput) (*schemas.SearchSchemasOutput, error)
	SearchSchemasAsync(ctx workflow.Context, input *schemas.SearchSchemasInput) *SchemasSearchSchemasFuture

	StartDiscoverer(ctx workflow.Context, input *schemas.StartDiscovererInput) (*schemas.StartDiscovererOutput, error)
	StartDiscovererAsync(ctx workflow.Context, input *schemas.StartDiscovererInput) *SchemasStartDiscovererFuture

	StopDiscoverer(ctx workflow.Context, input *schemas.StopDiscovererInput) (*schemas.StopDiscovererOutput, error)
	StopDiscovererAsync(ctx workflow.Context, input *schemas.StopDiscovererInput) *SchemasStopDiscovererFuture

	TagResource(ctx workflow.Context, input *schemas.TagResourceInput) (*schemas.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *schemas.TagResourceInput) *SchemasTagResourceFuture

	UntagResource(ctx workflow.Context, input *schemas.UntagResourceInput) (*schemas.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *schemas.UntagResourceInput) *SchemasUntagResourceFuture

	UpdateDiscoverer(ctx workflow.Context, input *schemas.UpdateDiscovererInput) (*schemas.UpdateDiscovererOutput, error)
	UpdateDiscovererAsync(ctx workflow.Context, input *schemas.UpdateDiscovererInput) *SchemasUpdateDiscovererFuture

	UpdateRegistry(ctx workflow.Context, input *schemas.UpdateRegistryInput) (*schemas.UpdateRegistryOutput, error)
	UpdateRegistryAsync(ctx workflow.Context, input *schemas.UpdateRegistryInput) *SchemasUpdateRegistryFuture

	UpdateSchema(ctx workflow.Context, input *schemas.UpdateSchemaInput) (*schemas.UpdateSchemaOutput, error)
	UpdateSchemaAsync(ctx workflow.Context, input *schemas.UpdateSchemaInput) *SchemasUpdateSchemaFuture

	WaitUntilCodeBindingExists(ctx workflow.Context, input *schemas.DescribeCodeBindingInput) error
	WaitUntilCodeBindingExistsAsync(ctx workflow.Context, input *schemas.DescribeCodeBindingInput) *awsclients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
