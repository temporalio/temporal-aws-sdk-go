package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/worklink"
	"github.com/aws/aws-sdk-go/service/worklink/worklinkiface"
)

type WorkLinkActivities struct {
	client worklinkiface.WorkLinkAPI
}

func NewWorkLinkActivities(client worklinkiface.WorkLinkAPI) *WorkLinkActivities {
    return &WorkLinkActivities{client: client}
}


func (a *WorkLinkActivities) AssociateDomain(input *worklink.AssociateDomainInput) (*worklink.AssociateDomainOutput, error) {
    return a.client.AssociateDomain(input)
}



func (a *WorkLinkActivities) AssociateWebsiteAuthorizationProvider(input *worklink.AssociateWebsiteAuthorizationProviderInput) (*worklink.AssociateWebsiteAuthorizationProviderOutput, error) {
    return a.client.AssociateWebsiteAuthorizationProvider(input)
}



func (a *WorkLinkActivities) AssociateWebsiteCertificateAuthority(input *worklink.AssociateWebsiteCertificateAuthorityInput) (*worklink.AssociateWebsiteCertificateAuthorityOutput, error) {
    return a.client.AssociateWebsiteCertificateAuthority(input)
}



func (a *WorkLinkActivities) CreateFleet(input *worklink.CreateFleetInput) (*worklink.CreateFleetOutput, error) {
    return a.client.CreateFleet(input)
}



func (a *WorkLinkActivities) DeleteFleet(input *worklink.DeleteFleetInput) (*worklink.DeleteFleetOutput, error) {
    return a.client.DeleteFleet(input)
}



func (a *WorkLinkActivities) DescribeAuditStreamConfiguration(input *worklink.DescribeAuditStreamConfigurationInput) (*worklink.DescribeAuditStreamConfigurationOutput, error) {
    return a.client.DescribeAuditStreamConfiguration(input)
}



func (a *WorkLinkActivities) DescribeCompanyNetworkConfiguration(input *worklink.DescribeCompanyNetworkConfigurationInput) (*worklink.DescribeCompanyNetworkConfigurationOutput, error) {
    return a.client.DescribeCompanyNetworkConfiguration(input)
}



func (a *WorkLinkActivities) DescribeDevice(input *worklink.DescribeDeviceInput) (*worklink.DescribeDeviceOutput, error) {
    return a.client.DescribeDevice(input)
}



func (a *WorkLinkActivities) DescribeDevicePolicyConfiguration(input *worklink.DescribeDevicePolicyConfigurationInput) (*worklink.DescribeDevicePolicyConfigurationOutput, error) {
    return a.client.DescribeDevicePolicyConfiguration(input)
}



func (a *WorkLinkActivities) DescribeDomain(input *worklink.DescribeDomainInput) (*worklink.DescribeDomainOutput, error) {
    return a.client.DescribeDomain(input)
}



func (a *WorkLinkActivities) DescribeFleetMetadata(input *worklink.DescribeFleetMetadataInput) (*worklink.DescribeFleetMetadataOutput, error) {
    return a.client.DescribeFleetMetadata(input)
}



func (a *WorkLinkActivities) DescribeIdentityProviderConfiguration(input *worklink.DescribeIdentityProviderConfigurationInput) (*worklink.DescribeIdentityProviderConfigurationOutput, error) {
    return a.client.DescribeIdentityProviderConfiguration(input)
}



func (a *WorkLinkActivities) DescribeWebsiteCertificateAuthority(input *worklink.DescribeWebsiteCertificateAuthorityInput) (*worklink.DescribeWebsiteCertificateAuthorityOutput, error) {
    return a.client.DescribeWebsiteCertificateAuthority(input)
}



func (a *WorkLinkActivities) DisassociateDomain(input *worklink.DisassociateDomainInput) (*worklink.DisassociateDomainOutput, error) {
    return a.client.DisassociateDomain(input)
}



func (a *WorkLinkActivities) DisassociateWebsiteAuthorizationProvider(input *worklink.DisassociateWebsiteAuthorizationProviderInput) (*worklink.DisassociateWebsiteAuthorizationProviderOutput, error) {
    return a.client.DisassociateWebsiteAuthorizationProvider(input)
}



func (a *WorkLinkActivities) DisassociateWebsiteCertificateAuthority(input *worklink.DisassociateWebsiteCertificateAuthorityInput) (*worklink.DisassociateWebsiteCertificateAuthorityOutput, error) {
    return a.client.DisassociateWebsiteCertificateAuthority(input)
}



func (a *WorkLinkActivities) ListDevices(input *worklink.ListDevicesInput) (*worklink.ListDevicesOutput, error) {
    return a.client.ListDevices(input)
}



func (a *WorkLinkActivities) ListDomains(input *worklink.ListDomainsInput) (*worklink.ListDomainsOutput, error) {
    return a.client.ListDomains(input)
}



func (a *WorkLinkActivities) ListFleets(input *worklink.ListFleetsInput) (*worklink.ListFleetsOutput, error) {
    return a.client.ListFleets(input)
}



func (a *WorkLinkActivities) ListTagsForResource(input *worklink.ListTagsForResourceInput) (*worklink.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}



func (a *WorkLinkActivities) ListWebsiteAuthorizationProviders(input *worklink.ListWebsiteAuthorizationProvidersInput) (*worklink.ListWebsiteAuthorizationProvidersOutput, error) {
    return a.client.ListWebsiteAuthorizationProviders(input)
}



func (a *WorkLinkActivities) ListWebsiteCertificateAuthorities(input *worklink.ListWebsiteCertificateAuthoritiesInput) (*worklink.ListWebsiteCertificateAuthoritiesOutput, error) {
    return a.client.ListWebsiteCertificateAuthorities(input)
}



func (a *WorkLinkActivities) RestoreDomainAccess(input *worklink.RestoreDomainAccessInput) (*worklink.RestoreDomainAccessOutput, error) {
    return a.client.RestoreDomainAccess(input)
}



func (a *WorkLinkActivities) RevokeDomainAccess(input *worklink.RevokeDomainAccessInput) (*worklink.RevokeDomainAccessOutput, error) {
    return a.client.RevokeDomainAccess(input)
}



func (a *WorkLinkActivities) SignOutUser(input *worklink.SignOutUserInput) (*worklink.SignOutUserOutput, error) {
    return a.client.SignOutUser(input)
}



func (a *WorkLinkActivities) TagResource(input *worklink.TagResourceInput) (*worklink.TagResourceOutput, error) {
    return a.client.TagResource(input)
}



func (a *WorkLinkActivities) UntagResource(input *worklink.UntagResourceInput) (*worklink.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}



func (a *WorkLinkActivities) UpdateAuditStreamConfiguration(input *worklink.UpdateAuditStreamConfigurationInput) (*worklink.UpdateAuditStreamConfigurationOutput, error) {
    return a.client.UpdateAuditStreamConfiguration(input)
}



func (a *WorkLinkActivities) UpdateCompanyNetworkConfiguration(input *worklink.UpdateCompanyNetworkConfigurationInput) (*worklink.UpdateCompanyNetworkConfigurationOutput, error) {
    return a.client.UpdateCompanyNetworkConfiguration(input)
}



func (a *WorkLinkActivities) UpdateDevicePolicyConfiguration(input *worklink.UpdateDevicePolicyConfigurationInput) (*worklink.UpdateDevicePolicyConfigurationOutput, error) {
    return a.client.UpdateDevicePolicyConfiguration(input)
}



func (a *WorkLinkActivities) UpdateDomainMetadata(input *worklink.UpdateDomainMetadataInput) (*worklink.UpdateDomainMetadataOutput, error) {
    return a.client.UpdateDomainMetadata(input)
}



func (a *WorkLinkActivities) UpdateFleetMetadata(input *worklink.UpdateFleetMetadataInput) (*worklink.UpdateFleetMetadataOutput, error) {
    return a.client.UpdateFleetMetadata(input)
}



func (a *WorkLinkActivities) UpdateIdentityProviderConfiguration(input *worklink.UpdateIdentityProviderConfigurationInput) (*worklink.UpdateIdentityProviderConfigurationOutput, error) {
    return a.client.UpdateIdentityProviderConfiguration(input)
}

