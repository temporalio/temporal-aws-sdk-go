// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package elasticsearchservicestub

import (
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	AcceptInboundCrossClusterSearchConnection(ctx workflow.Context, input *elasticsearchservice.AcceptInboundCrossClusterSearchConnectionInput) (*elasticsearchservice.AcceptInboundCrossClusterSearchConnectionOutput, error)
	AcceptInboundCrossClusterSearchConnectionAsync(ctx workflow.Context, input *elasticsearchservice.AcceptInboundCrossClusterSearchConnectionInput) *ElasticsearchServiceAcceptInboundCrossClusterSearchConnectionFuture

	AddTags(ctx workflow.Context, input *elasticsearchservice.AddTagsInput) (*elasticsearchservice.AddTagsOutput, error)
	AddTagsAsync(ctx workflow.Context, input *elasticsearchservice.AddTagsInput) *ElasticsearchServiceAddTagsFuture

	AssociatePackage(ctx workflow.Context, input *elasticsearchservice.AssociatePackageInput) (*elasticsearchservice.AssociatePackageOutput, error)
	AssociatePackageAsync(ctx workflow.Context, input *elasticsearchservice.AssociatePackageInput) *ElasticsearchServiceAssociatePackageFuture

	CancelElasticsearchServiceSoftwareUpdate(ctx workflow.Context, input *elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateInput) (*elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateOutput, error)
	CancelElasticsearchServiceSoftwareUpdateAsync(ctx workflow.Context, input *elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateInput) *ElasticsearchServiceCancelElasticsearchServiceSoftwareUpdateFuture

	CreateElasticsearchDomain(ctx workflow.Context, input *elasticsearchservice.CreateElasticsearchDomainInput) (*elasticsearchservice.CreateElasticsearchDomainOutput, error)
	CreateElasticsearchDomainAsync(ctx workflow.Context, input *elasticsearchservice.CreateElasticsearchDomainInput) *ElasticsearchServiceCreateElasticsearchDomainFuture

	CreateOutboundCrossClusterSearchConnection(ctx workflow.Context, input *elasticsearchservice.CreateOutboundCrossClusterSearchConnectionInput) (*elasticsearchservice.CreateOutboundCrossClusterSearchConnectionOutput, error)
	CreateOutboundCrossClusterSearchConnectionAsync(ctx workflow.Context, input *elasticsearchservice.CreateOutboundCrossClusterSearchConnectionInput) *ElasticsearchServiceCreateOutboundCrossClusterSearchConnectionFuture

	CreatePackage(ctx workflow.Context, input *elasticsearchservice.CreatePackageInput) (*elasticsearchservice.CreatePackageOutput, error)
	CreatePackageAsync(ctx workflow.Context, input *elasticsearchservice.CreatePackageInput) *ElasticsearchServiceCreatePackageFuture

	DeleteElasticsearchDomain(ctx workflow.Context, input *elasticsearchservice.DeleteElasticsearchDomainInput) (*elasticsearchservice.DeleteElasticsearchDomainOutput, error)
	DeleteElasticsearchDomainAsync(ctx workflow.Context, input *elasticsearchservice.DeleteElasticsearchDomainInput) *ElasticsearchServiceDeleteElasticsearchDomainFuture

	DeleteElasticsearchServiceRole(ctx workflow.Context, input *elasticsearchservice.DeleteElasticsearchServiceRoleInput) (*elasticsearchservice.DeleteElasticsearchServiceRoleOutput, error)
	DeleteElasticsearchServiceRoleAsync(ctx workflow.Context, input *elasticsearchservice.DeleteElasticsearchServiceRoleInput) *ElasticsearchServiceDeleteElasticsearchServiceRoleFuture

	DeleteInboundCrossClusterSearchConnection(ctx workflow.Context, input *elasticsearchservice.DeleteInboundCrossClusterSearchConnectionInput) (*elasticsearchservice.DeleteInboundCrossClusterSearchConnectionOutput, error)
	DeleteInboundCrossClusterSearchConnectionAsync(ctx workflow.Context, input *elasticsearchservice.DeleteInboundCrossClusterSearchConnectionInput) *ElasticsearchServiceDeleteInboundCrossClusterSearchConnectionFuture

	DeleteOutboundCrossClusterSearchConnection(ctx workflow.Context, input *elasticsearchservice.DeleteOutboundCrossClusterSearchConnectionInput) (*elasticsearchservice.DeleteOutboundCrossClusterSearchConnectionOutput, error)
	DeleteOutboundCrossClusterSearchConnectionAsync(ctx workflow.Context, input *elasticsearchservice.DeleteOutboundCrossClusterSearchConnectionInput) *ElasticsearchServiceDeleteOutboundCrossClusterSearchConnectionFuture

	DeletePackage(ctx workflow.Context, input *elasticsearchservice.DeletePackageInput) (*elasticsearchservice.DeletePackageOutput, error)
	DeletePackageAsync(ctx workflow.Context, input *elasticsearchservice.DeletePackageInput) *ElasticsearchServiceDeletePackageFuture

	DescribeElasticsearchDomain(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchDomainInput) (*elasticsearchservice.DescribeElasticsearchDomainOutput, error)
	DescribeElasticsearchDomainAsync(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchDomainInput) *ElasticsearchServiceDescribeElasticsearchDomainFuture

	DescribeElasticsearchDomainConfig(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchDomainConfigInput) (*elasticsearchservice.DescribeElasticsearchDomainConfigOutput, error)
	DescribeElasticsearchDomainConfigAsync(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchDomainConfigInput) *ElasticsearchServiceDescribeElasticsearchDomainConfigFuture

	DescribeElasticsearchDomains(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchDomainsInput) (*elasticsearchservice.DescribeElasticsearchDomainsOutput, error)
	DescribeElasticsearchDomainsAsync(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchDomainsInput) *ElasticsearchServiceDescribeElasticsearchDomainsFuture

	DescribeElasticsearchInstanceTypeLimits(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsInput) (*elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsOutput, error)
	DescribeElasticsearchInstanceTypeLimitsAsync(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsInput) *ElasticsearchServiceDescribeElasticsearchInstanceTypeLimitsFuture

	DescribeInboundCrossClusterSearchConnections(ctx workflow.Context, input *elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsInput) (*elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput, error)
	DescribeInboundCrossClusterSearchConnectionsAsync(ctx workflow.Context, input *elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsInput) *ElasticsearchServiceDescribeInboundCrossClusterSearchConnectionsFuture

	DescribeOutboundCrossClusterSearchConnections(ctx workflow.Context, input *elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsInput) (*elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput, error)
	DescribeOutboundCrossClusterSearchConnectionsAsync(ctx workflow.Context, input *elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsInput) *ElasticsearchServiceDescribeOutboundCrossClusterSearchConnectionsFuture

	DescribePackages(ctx workflow.Context, input *elasticsearchservice.DescribePackagesInput) (*elasticsearchservice.DescribePackagesOutput, error)
	DescribePackagesAsync(ctx workflow.Context, input *elasticsearchservice.DescribePackagesInput) *ElasticsearchServiceDescribePackagesFuture

	DescribeReservedElasticsearchInstanceOfferings(ctx workflow.Context, input *elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsInput) (*elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput, error)
	DescribeReservedElasticsearchInstanceOfferingsAsync(ctx workflow.Context, input *elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsInput) *ElasticsearchServiceDescribeReservedElasticsearchInstanceOfferingsFuture

	DescribeReservedElasticsearchInstances(ctx workflow.Context, input *elasticsearchservice.DescribeReservedElasticsearchInstancesInput) (*elasticsearchservice.DescribeReservedElasticsearchInstancesOutput, error)
	DescribeReservedElasticsearchInstancesAsync(ctx workflow.Context, input *elasticsearchservice.DescribeReservedElasticsearchInstancesInput) *ElasticsearchServiceDescribeReservedElasticsearchInstancesFuture

	DissociatePackage(ctx workflow.Context, input *elasticsearchservice.DissociatePackageInput) (*elasticsearchservice.DissociatePackageOutput, error)
	DissociatePackageAsync(ctx workflow.Context, input *elasticsearchservice.DissociatePackageInput) *ElasticsearchServiceDissociatePackageFuture

	GetCompatibleElasticsearchVersions(ctx workflow.Context, input *elasticsearchservice.GetCompatibleElasticsearchVersionsInput) (*elasticsearchservice.GetCompatibleElasticsearchVersionsOutput, error)
	GetCompatibleElasticsearchVersionsAsync(ctx workflow.Context, input *elasticsearchservice.GetCompatibleElasticsearchVersionsInput) *ElasticsearchServiceGetCompatibleElasticsearchVersionsFuture

	GetUpgradeHistory(ctx workflow.Context, input *elasticsearchservice.GetUpgradeHistoryInput) (*elasticsearchservice.GetUpgradeHistoryOutput, error)
	GetUpgradeHistoryAsync(ctx workflow.Context, input *elasticsearchservice.GetUpgradeHistoryInput) *ElasticsearchServiceGetUpgradeHistoryFuture

	GetUpgradeStatus(ctx workflow.Context, input *elasticsearchservice.GetUpgradeStatusInput) (*elasticsearchservice.GetUpgradeStatusOutput, error)
	GetUpgradeStatusAsync(ctx workflow.Context, input *elasticsearchservice.GetUpgradeStatusInput) *ElasticsearchServiceGetUpgradeStatusFuture

	ListDomainNames(ctx workflow.Context, input *elasticsearchservice.ListDomainNamesInput) (*elasticsearchservice.ListDomainNamesOutput, error)
	ListDomainNamesAsync(ctx workflow.Context, input *elasticsearchservice.ListDomainNamesInput) *ElasticsearchServiceListDomainNamesFuture

	ListDomainsForPackage(ctx workflow.Context, input *elasticsearchservice.ListDomainsForPackageInput) (*elasticsearchservice.ListDomainsForPackageOutput, error)
	ListDomainsForPackageAsync(ctx workflow.Context, input *elasticsearchservice.ListDomainsForPackageInput) *ElasticsearchServiceListDomainsForPackageFuture

	ListElasticsearchInstanceTypes(ctx workflow.Context, input *elasticsearchservice.ListElasticsearchInstanceTypesInput) (*elasticsearchservice.ListElasticsearchInstanceTypesOutput, error)
	ListElasticsearchInstanceTypesAsync(ctx workflow.Context, input *elasticsearchservice.ListElasticsearchInstanceTypesInput) *ElasticsearchServiceListElasticsearchInstanceTypesFuture

	ListElasticsearchVersions(ctx workflow.Context, input *elasticsearchservice.ListElasticsearchVersionsInput) (*elasticsearchservice.ListElasticsearchVersionsOutput, error)
	ListElasticsearchVersionsAsync(ctx workflow.Context, input *elasticsearchservice.ListElasticsearchVersionsInput) *ElasticsearchServiceListElasticsearchVersionsFuture

	ListPackagesForDomain(ctx workflow.Context, input *elasticsearchservice.ListPackagesForDomainInput) (*elasticsearchservice.ListPackagesForDomainOutput, error)
	ListPackagesForDomainAsync(ctx workflow.Context, input *elasticsearchservice.ListPackagesForDomainInput) *ElasticsearchServiceListPackagesForDomainFuture

	ListTags(ctx workflow.Context, input *elasticsearchservice.ListTagsInput) (*elasticsearchservice.ListTagsOutput, error)
	ListTagsAsync(ctx workflow.Context, input *elasticsearchservice.ListTagsInput) *ElasticsearchServiceListTagsFuture

	PurchaseReservedElasticsearchInstanceOffering(ctx workflow.Context, input *elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingInput) (*elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingOutput, error)
	PurchaseReservedElasticsearchInstanceOfferingAsync(ctx workflow.Context, input *elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingInput) *ElasticsearchServicePurchaseReservedElasticsearchInstanceOfferingFuture

	RejectInboundCrossClusterSearchConnection(ctx workflow.Context, input *elasticsearchservice.RejectInboundCrossClusterSearchConnectionInput) (*elasticsearchservice.RejectInboundCrossClusterSearchConnectionOutput, error)
	RejectInboundCrossClusterSearchConnectionAsync(ctx workflow.Context, input *elasticsearchservice.RejectInboundCrossClusterSearchConnectionInput) *ElasticsearchServiceRejectInboundCrossClusterSearchConnectionFuture

	RemoveTags(ctx workflow.Context, input *elasticsearchservice.RemoveTagsInput) (*elasticsearchservice.RemoveTagsOutput, error)
	RemoveTagsAsync(ctx workflow.Context, input *elasticsearchservice.RemoveTagsInput) *ElasticsearchServiceRemoveTagsFuture

	StartElasticsearchServiceSoftwareUpdate(ctx workflow.Context, input *elasticsearchservice.StartElasticsearchServiceSoftwareUpdateInput) (*elasticsearchservice.StartElasticsearchServiceSoftwareUpdateOutput, error)
	StartElasticsearchServiceSoftwareUpdateAsync(ctx workflow.Context, input *elasticsearchservice.StartElasticsearchServiceSoftwareUpdateInput) *ElasticsearchServiceStartElasticsearchServiceSoftwareUpdateFuture

	UpdateElasticsearchDomainConfig(ctx workflow.Context, input *elasticsearchservice.UpdateElasticsearchDomainConfigInput) (*elasticsearchservice.UpdateElasticsearchDomainConfigOutput, error)
	UpdateElasticsearchDomainConfigAsync(ctx workflow.Context, input *elasticsearchservice.UpdateElasticsearchDomainConfigInput) *ElasticsearchServiceUpdateElasticsearchDomainConfigFuture

	UpgradeElasticsearchDomain(ctx workflow.Context, input *elasticsearchservice.UpgradeElasticsearchDomainInput) (*elasticsearchservice.UpgradeElasticsearchDomainOutput, error)
	UpgradeElasticsearchDomainAsync(ctx workflow.Context, input *elasticsearchservice.UpgradeElasticsearchDomainInput) *ElasticsearchServiceUpgradeElasticsearchDomainFuture
}

func NewClient() Client {
	return &stub{}
}
