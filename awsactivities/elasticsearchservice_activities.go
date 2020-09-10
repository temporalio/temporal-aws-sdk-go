package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice/elasticsearchserviceiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type ElasticsearchServiceActivities struct {
	client elasticsearchserviceiface.ElasticsearchServiceAPI
}

func NewElasticsearchServiceActivities(session *session.Session, config ...*aws.Config) *ElasticsearchServiceActivities {
	client := elasticsearchservice.New(session, config...)
	return &ElasticsearchServiceActivities{client: client}
}

func (a *ElasticsearchServiceActivities) AcceptInboundCrossClusterSearchConnection(ctx context.Context, input *elasticsearchservice.AcceptInboundCrossClusterSearchConnectionInput) (*elasticsearchservice.AcceptInboundCrossClusterSearchConnectionOutput, error) {
	return a.client.AcceptInboundCrossClusterSearchConnectionWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) AddTags(ctx context.Context, input *elasticsearchservice.AddTagsInput) (*elasticsearchservice.AddTagsOutput, error) {
	return a.client.AddTagsWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) AssociatePackage(ctx context.Context, input *elasticsearchservice.AssociatePackageInput) (*elasticsearchservice.AssociatePackageOutput, error) {
	return a.client.AssociatePackageWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) CancelElasticsearchServiceSoftwareUpdate(ctx context.Context, input *elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateInput) (*elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateOutput, error) {
	return a.client.CancelElasticsearchServiceSoftwareUpdateWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) CreateElasticsearchDomain(ctx context.Context, input *elasticsearchservice.CreateElasticsearchDomainInput) (*elasticsearchservice.CreateElasticsearchDomainOutput, error) {
	return a.client.CreateElasticsearchDomainWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) CreateOutboundCrossClusterSearchConnection(ctx context.Context, input *elasticsearchservice.CreateOutboundCrossClusterSearchConnectionInput) (*elasticsearchservice.CreateOutboundCrossClusterSearchConnectionOutput, error) {
	return a.client.CreateOutboundCrossClusterSearchConnectionWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) CreatePackage(ctx context.Context, input *elasticsearchservice.CreatePackageInput) (*elasticsearchservice.CreatePackageOutput, error) {
	return a.client.CreatePackageWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) DeleteElasticsearchDomain(ctx context.Context, input *elasticsearchservice.DeleteElasticsearchDomainInput) (*elasticsearchservice.DeleteElasticsearchDomainOutput, error) {
	return a.client.DeleteElasticsearchDomainWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) DeleteElasticsearchServiceRole(ctx context.Context, input *elasticsearchservice.DeleteElasticsearchServiceRoleInput) (*elasticsearchservice.DeleteElasticsearchServiceRoleOutput, error) {
	return a.client.DeleteElasticsearchServiceRoleWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) DeleteInboundCrossClusterSearchConnection(ctx context.Context, input *elasticsearchservice.DeleteInboundCrossClusterSearchConnectionInput) (*elasticsearchservice.DeleteInboundCrossClusterSearchConnectionOutput, error) {
	return a.client.DeleteInboundCrossClusterSearchConnectionWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) DeleteOutboundCrossClusterSearchConnection(ctx context.Context, input *elasticsearchservice.DeleteOutboundCrossClusterSearchConnectionInput) (*elasticsearchservice.DeleteOutboundCrossClusterSearchConnectionOutput, error) {
	return a.client.DeleteOutboundCrossClusterSearchConnectionWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) DeletePackage(ctx context.Context, input *elasticsearchservice.DeletePackageInput) (*elasticsearchservice.DeletePackageOutput, error) {
	return a.client.DeletePackageWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) DescribeElasticsearchDomain(ctx context.Context, input *elasticsearchservice.DescribeElasticsearchDomainInput) (*elasticsearchservice.DescribeElasticsearchDomainOutput, error) {
	return a.client.DescribeElasticsearchDomainWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) DescribeElasticsearchDomainConfig(ctx context.Context, input *elasticsearchservice.DescribeElasticsearchDomainConfigInput) (*elasticsearchservice.DescribeElasticsearchDomainConfigOutput, error) {
	return a.client.DescribeElasticsearchDomainConfigWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) DescribeElasticsearchDomains(ctx context.Context, input *elasticsearchservice.DescribeElasticsearchDomainsInput) (*elasticsearchservice.DescribeElasticsearchDomainsOutput, error) {
	return a.client.DescribeElasticsearchDomainsWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) DescribeElasticsearchInstanceTypeLimits(ctx context.Context, input *elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsInput) (*elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsOutput, error) {
	return a.client.DescribeElasticsearchInstanceTypeLimitsWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) DescribeInboundCrossClusterSearchConnections(ctx context.Context, input *elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsInput) (*elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput, error) {
	return a.client.DescribeInboundCrossClusterSearchConnectionsWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) DescribeOutboundCrossClusterSearchConnections(ctx context.Context, input *elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsInput) (*elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput, error) {
	return a.client.DescribeOutboundCrossClusterSearchConnectionsWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) DescribePackages(ctx context.Context, input *elasticsearchservice.DescribePackagesInput) (*elasticsearchservice.DescribePackagesOutput, error) {
	return a.client.DescribePackagesWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) DescribeReservedElasticsearchInstanceOfferings(ctx context.Context, input *elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsInput) (*elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput, error) {
	return a.client.DescribeReservedElasticsearchInstanceOfferingsWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) DescribeReservedElasticsearchInstances(ctx context.Context, input *elasticsearchservice.DescribeReservedElasticsearchInstancesInput) (*elasticsearchservice.DescribeReservedElasticsearchInstancesOutput, error) {
	return a.client.DescribeReservedElasticsearchInstancesWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) DissociatePackage(ctx context.Context, input *elasticsearchservice.DissociatePackageInput) (*elasticsearchservice.DissociatePackageOutput, error) {
	return a.client.DissociatePackageWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) GetCompatibleElasticsearchVersions(ctx context.Context, input *elasticsearchservice.GetCompatibleElasticsearchVersionsInput) (*elasticsearchservice.GetCompatibleElasticsearchVersionsOutput, error) {
	return a.client.GetCompatibleElasticsearchVersionsWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) GetUpgradeHistory(ctx context.Context, input *elasticsearchservice.GetUpgradeHistoryInput) (*elasticsearchservice.GetUpgradeHistoryOutput, error) {
	return a.client.GetUpgradeHistoryWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) GetUpgradeStatus(ctx context.Context, input *elasticsearchservice.GetUpgradeStatusInput) (*elasticsearchservice.GetUpgradeStatusOutput, error) {
	return a.client.GetUpgradeStatusWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) ListDomainNames(ctx context.Context, input *elasticsearchservice.ListDomainNamesInput) (*elasticsearchservice.ListDomainNamesOutput, error) {
	return a.client.ListDomainNamesWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) ListDomainsForPackage(ctx context.Context, input *elasticsearchservice.ListDomainsForPackageInput) (*elasticsearchservice.ListDomainsForPackageOutput, error) {
	return a.client.ListDomainsForPackageWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) ListElasticsearchInstanceTypes(ctx context.Context, input *elasticsearchservice.ListElasticsearchInstanceTypesInput) (*elasticsearchservice.ListElasticsearchInstanceTypesOutput, error) {
	return a.client.ListElasticsearchInstanceTypesWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) ListElasticsearchVersions(ctx context.Context, input *elasticsearchservice.ListElasticsearchVersionsInput) (*elasticsearchservice.ListElasticsearchVersionsOutput, error) {
	return a.client.ListElasticsearchVersionsWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) ListPackagesForDomain(ctx context.Context, input *elasticsearchservice.ListPackagesForDomainInput) (*elasticsearchservice.ListPackagesForDomainOutput, error) {
	return a.client.ListPackagesForDomainWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) ListTags(ctx context.Context, input *elasticsearchservice.ListTagsInput) (*elasticsearchservice.ListTagsOutput, error) {
	return a.client.ListTagsWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) PurchaseReservedElasticsearchInstanceOffering(ctx context.Context, input *elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingInput) (*elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingOutput, error) {
	return a.client.PurchaseReservedElasticsearchInstanceOfferingWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) RejectInboundCrossClusterSearchConnection(ctx context.Context, input *elasticsearchservice.RejectInboundCrossClusterSearchConnectionInput) (*elasticsearchservice.RejectInboundCrossClusterSearchConnectionOutput, error) {
	return a.client.RejectInboundCrossClusterSearchConnectionWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) RemoveTags(ctx context.Context, input *elasticsearchservice.RemoveTagsInput) (*elasticsearchservice.RemoveTagsOutput, error) {
	return a.client.RemoveTagsWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) StartElasticsearchServiceSoftwareUpdate(ctx context.Context, input *elasticsearchservice.StartElasticsearchServiceSoftwareUpdateInput) (*elasticsearchservice.StartElasticsearchServiceSoftwareUpdateOutput, error) {
	return a.client.StartElasticsearchServiceSoftwareUpdateWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) UpdateElasticsearchDomainConfig(ctx context.Context, input *elasticsearchservice.UpdateElasticsearchDomainConfigInput) (*elasticsearchservice.UpdateElasticsearchDomainConfigOutput, error) {
	return a.client.UpdateElasticsearchDomainConfigWithContext(ctx, input)
}

func (a *ElasticsearchServiceActivities) UpgradeElasticsearchDomain(ctx context.Context, input *elasticsearchservice.UpgradeElasticsearchDomainInput) (*elasticsearchservice.UpgradeElasticsearchDomainOutput, error) {
	return a.client.UpgradeElasticsearchDomainWithContext(ctx, input)
}
