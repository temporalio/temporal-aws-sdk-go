
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice/elasticsearchserviceiface"
)

type ElasticsearchServiceActivities struct {
	client elasticsearchserviceiface.ElasticsearchServiceAPI
}

func NewElasticsearchServiceActivities(session *session.Session, config... *aws.Config) *ElasticsearchServiceActivities {
    client := elasticsearchservice.New(session, config...)
    return &ElasticsearchServiceActivities{client: client}
}

func (a *ElasticsearchServiceActivities) AcceptInboundCrossClusterSearchConnection(input *elasticsearchservice.AcceptInboundCrossClusterSearchConnectionInput) (*elasticsearchservice.AcceptInboundCrossClusterSearchConnectionOutput, error) {
    return a.client.AcceptInboundCrossClusterSearchConnection(input)
}

func (a *ElasticsearchServiceActivities) AddTags(input *elasticsearchservice.AddTagsInput) (*elasticsearchservice.AddTagsOutput, error) {
    return a.client.AddTags(input)
}

func (a *ElasticsearchServiceActivities) AssociatePackage(input *elasticsearchservice.AssociatePackageInput) (*elasticsearchservice.AssociatePackageOutput, error) {
    return a.client.AssociatePackage(input)
}

func (a *ElasticsearchServiceActivities) CancelElasticsearchServiceSoftwareUpdate(input *elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateInput) (*elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateOutput, error) {
    return a.client.CancelElasticsearchServiceSoftwareUpdate(input)
}

func (a *ElasticsearchServiceActivities) CreateElasticsearchDomain(input *elasticsearchservice.CreateElasticsearchDomainInput) (*elasticsearchservice.CreateElasticsearchDomainOutput, error) {
    return a.client.CreateElasticsearchDomain(input)
}

func (a *ElasticsearchServiceActivities) CreateOutboundCrossClusterSearchConnection(input *elasticsearchservice.CreateOutboundCrossClusterSearchConnectionInput) (*elasticsearchservice.CreateOutboundCrossClusterSearchConnectionOutput, error) {
    return a.client.CreateOutboundCrossClusterSearchConnection(input)
}

func (a *ElasticsearchServiceActivities) CreatePackage(input *elasticsearchservice.CreatePackageInput) (*elasticsearchservice.CreatePackageOutput, error) {
    return a.client.CreatePackage(input)
}

func (a *ElasticsearchServiceActivities) DeleteElasticsearchDomain(input *elasticsearchservice.DeleteElasticsearchDomainInput) (*elasticsearchservice.DeleteElasticsearchDomainOutput, error) {
    return a.client.DeleteElasticsearchDomain(input)
}

func (a *ElasticsearchServiceActivities) DeleteElasticsearchServiceRole(input *elasticsearchservice.DeleteElasticsearchServiceRoleInput) (*elasticsearchservice.DeleteElasticsearchServiceRoleOutput, error) {
    return a.client.DeleteElasticsearchServiceRole(input)
}

func (a *ElasticsearchServiceActivities) DeleteInboundCrossClusterSearchConnection(input *elasticsearchservice.DeleteInboundCrossClusterSearchConnectionInput) (*elasticsearchservice.DeleteInboundCrossClusterSearchConnectionOutput, error) {
    return a.client.DeleteInboundCrossClusterSearchConnection(input)
}

func (a *ElasticsearchServiceActivities) DeleteOutboundCrossClusterSearchConnection(input *elasticsearchservice.DeleteOutboundCrossClusterSearchConnectionInput) (*elasticsearchservice.DeleteOutboundCrossClusterSearchConnectionOutput, error) {
    return a.client.DeleteOutboundCrossClusterSearchConnection(input)
}

func (a *ElasticsearchServiceActivities) DeletePackage(input *elasticsearchservice.DeletePackageInput) (*elasticsearchservice.DeletePackageOutput, error) {
    return a.client.DeletePackage(input)
}

func (a *ElasticsearchServiceActivities) DescribeElasticsearchDomain(input *elasticsearchservice.DescribeElasticsearchDomainInput) (*elasticsearchservice.DescribeElasticsearchDomainOutput, error) {
    return a.client.DescribeElasticsearchDomain(input)
}

func (a *ElasticsearchServiceActivities) DescribeElasticsearchDomainConfig(input *elasticsearchservice.DescribeElasticsearchDomainConfigInput) (*elasticsearchservice.DescribeElasticsearchDomainConfigOutput, error) {
    return a.client.DescribeElasticsearchDomainConfig(input)
}

func (a *ElasticsearchServiceActivities) DescribeElasticsearchDomains(input *elasticsearchservice.DescribeElasticsearchDomainsInput) (*elasticsearchservice.DescribeElasticsearchDomainsOutput, error) {
    return a.client.DescribeElasticsearchDomains(input)
}

func (a *ElasticsearchServiceActivities) DescribeElasticsearchInstanceTypeLimits(input *elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsInput) (*elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsOutput, error) {
    return a.client.DescribeElasticsearchInstanceTypeLimits(input)
}

func (a *ElasticsearchServiceActivities) DescribeInboundCrossClusterSearchConnections(input *elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsInput) (*elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput, error) {
    return a.client.DescribeInboundCrossClusterSearchConnections(input)
}

func (a *ElasticsearchServiceActivities) DescribeOutboundCrossClusterSearchConnections(input *elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsInput) (*elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput, error) {
    return a.client.DescribeOutboundCrossClusterSearchConnections(input)
}

func (a *ElasticsearchServiceActivities) DescribePackages(input *elasticsearchservice.DescribePackagesInput) (*elasticsearchservice.DescribePackagesOutput, error) {
    return a.client.DescribePackages(input)
}

func (a *ElasticsearchServiceActivities) DescribeReservedElasticsearchInstanceOfferings(input *elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsInput) (*elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput, error) {
    return a.client.DescribeReservedElasticsearchInstanceOfferings(input)
}

func (a *ElasticsearchServiceActivities) DescribeReservedElasticsearchInstances(input *elasticsearchservice.DescribeReservedElasticsearchInstancesInput) (*elasticsearchservice.DescribeReservedElasticsearchInstancesOutput, error) {
    return a.client.DescribeReservedElasticsearchInstances(input)
}

func (a *ElasticsearchServiceActivities) DissociatePackage(input *elasticsearchservice.DissociatePackageInput) (*elasticsearchservice.DissociatePackageOutput, error) {
    return a.client.DissociatePackage(input)
}

func (a *ElasticsearchServiceActivities) GetCompatibleElasticsearchVersions(input *elasticsearchservice.GetCompatibleElasticsearchVersionsInput) (*elasticsearchservice.GetCompatibleElasticsearchVersionsOutput, error) {
    return a.client.GetCompatibleElasticsearchVersions(input)
}

func (a *ElasticsearchServiceActivities) GetUpgradeHistory(input *elasticsearchservice.GetUpgradeHistoryInput) (*elasticsearchservice.GetUpgradeHistoryOutput, error) {
    return a.client.GetUpgradeHistory(input)
}

func (a *ElasticsearchServiceActivities) GetUpgradeStatus(input *elasticsearchservice.GetUpgradeStatusInput) (*elasticsearchservice.GetUpgradeStatusOutput, error) {
    return a.client.GetUpgradeStatus(input)
}

func (a *ElasticsearchServiceActivities) ListDomainNames(input *elasticsearchservice.ListDomainNamesInput) (*elasticsearchservice.ListDomainNamesOutput, error) {
    return a.client.ListDomainNames(input)
}

func (a *ElasticsearchServiceActivities) ListDomainsForPackage(input *elasticsearchservice.ListDomainsForPackageInput) (*elasticsearchservice.ListDomainsForPackageOutput, error) {
    return a.client.ListDomainsForPackage(input)
}

func (a *ElasticsearchServiceActivities) ListElasticsearchInstanceTypes(input *elasticsearchservice.ListElasticsearchInstanceTypesInput) (*elasticsearchservice.ListElasticsearchInstanceTypesOutput, error) {
    return a.client.ListElasticsearchInstanceTypes(input)
}

func (a *ElasticsearchServiceActivities) ListElasticsearchVersions(input *elasticsearchservice.ListElasticsearchVersionsInput) (*elasticsearchservice.ListElasticsearchVersionsOutput, error) {
    return a.client.ListElasticsearchVersions(input)
}

func (a *ElasticsearchServiceActivities) ListPackagesForDomain(input *elasticsearchservice.ListPackagesForDomainInput) (*elasticsearchservice.ListPackagesForDomainOutput, error) {
    return a.client.ListPackagesForDomain(input)
}

func (a *ElasticsearchServiceActivities) ListTags(input *elasticsearchservice.ListTagsInput) (*elasticsearchservice.ListTagsOutput, error) {
    return a.client.ListTags(input)
}

func (a *ElasticsearchServiceActivities) PurchaseReservedElasticsearchInstanceOffering(input *elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingInput) (*elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingOutput, error) {
    return a.client.PurchaseReservedElasticsearchInstanceOffering(input)
}

func (a *ElasticsearchServiceActivities) RejectInboundCrossClusterSearchConnection(input *elasticsearchservice.RejectInboundCrossClusterSearchConnectionInput) (*elasticsearchservice.RejectInboundCrossClusterSearchConnectionOutput, error) {
    return a.client.RejectInboundCrossClusterSearchConnection(input)
}

func (a *ElasticsearchServiceActivities) RemoveTags(input *elasticsearchservice.RemoveTagsInput) (*elasticsearchservice.RemoveTagsOutput, error) {
    return a.client.RemoveTags(input)
}

func (a *ElasticsearchServiceActivities) StartElasticsearchServiceSoftwareUpdate(input *elasticsearchservice.StartElasticsearchServiceSoftwareUpdateInput) (*elasticsearchservice.StartElasticsearchServiceSoftwareUpdateOutput, error) {
    return a.client.StartElasticsearchServiceSoftwareUpdate(input)
}

func (a *ElasticsearchServiceActivities) UpdateElasticsearchDomainConfig(input *elasticsearchservice.UpdateElasticsearchDomainConfigInput) (*elasticsearchservice.UpdateElasticsearchDomainConfigOutput, error) {
    return a.client.UpdateElasticsearchDomainConfig(input)
}

func (a *ElasticsearchServiceActivities) UpgradeElasticsearchDomain(input *elasticsearchservice.UpgradeElasticsearchDomainInput) (*elasticsearchservice.UpgradeElasticsearchDomainOutput, error) {
    return a.client.UpgradeElasticsearchDomain(input)
}
