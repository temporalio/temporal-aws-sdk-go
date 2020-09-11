package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/worklink"
	"github.com/aws/aws-sdk-go/service/worklink/worklinkiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type WorkLinkActivities struct {
	client worklinkiface.WorkLinkAPI
}

func NewWorkLinkActivities(session *session.Session, config ...*aws.Config) *WorkLinkActivities {
	client := worklink.New(session, config...)
	return &WorkLinkActivities{client: client}
}

func (a *WorkLinkActivities) AssociateDomain(ctx context.Context, input *worklink.AssociateDomainInput) (*worklink.AssociateDomainOutput, error) {
	return a.client.AssociateDomainWithContext(ctx, input)
}

func (a *WorkLinkActivities) AssociateWebsiteAuthorizationProvider(ctx context.Context, input *worklink.AssociateWebsiteAuthorizationProviderInput) (*worklink.AssociateWebsiteAuthorizationProviderOutput, error) {
	return a.client.AssociateWebsiteAuthorizationProviderWithContext(ctx, input)
}

func (a *WorkLinkActivities) AssociateWebsiteCertificateAuthority(ctx context.Context, input *worklink.AssociateWebsiteCertificateAuthorityInput) (*worklink.AssociateWebsiteCertificateAuthorityOutput, error) {
	return a.client.AssociateWebsiteCertificateAuthorityWithContext(ctx, input)
}

func (a *WorkLinkActivities) CreateFleet(ctx context.Context, input *worklink.CreateFleetInput) (*worklink.CreateFleetOutput, error) {
	return a.client.CreateFleetWithContext(ctx, input)
}

func (a *WorkLinkActivities) DeleteFleet(ctx context.Context, input *worklink.DeleteFleetInput) (*worklink.DeleteFleetOutput, error) {
	return a.client.DeleteFleetWithContext(ctx, input)
}

func (a *WorkLinkActivities) DescribeAuditStreamConfiguration(ctx context.Context, input *worklink.DescribeAuditStreamConfigurationInput) (*worklink.DescribeAuditStreamConfigurationOutput, error) {
	return a.client.DescribeAuditStreamConfigurationWithContext(ctx, input)
}

func (a *WorkLinkActivities) DescribeCompanyNetworkConfiguration(ctx context.Context, input *worklink.DescribeCompanyNetworkConfigurationInput) (*worklink.DescribeCompanyNetworkConfigurationOutput, error) {
	return a.client.DescribeCompanyNetworkConfigurationWithContext(ctx, input)
}

func (a *WorkLinkActivities) DescribeDevice(ctx context.Context, input *worklink.DescribeDeviceInput) (*worklink.DescribeDeviceOutput, error) {
	return a.client.DescribeDeviceWithContext(ctx, input)
}

func (a *WorkLinkActivities) DescribeDevicePolicyConfiguration(ctx context.Context, input *worklink.DescribeDevicePolicyConfigurationInput) (*worklink.DescribeDevicePolicyConfigurationOutput, error) {
	return a.client.DescribeDevicePolicyConfigurationWithContext(ctx, input)
}

func (a *WorkLinkActivities) DescribeDomain(ctx context.Context, input *worklink.DescribeDomainInput) (*worklink.DescribeDomainOutput, error) {
	return a.client.DescribeDomainWithContext(ctx, input)
}

func (a *WorkLinkActivities) DescribeFleetMetadata(ctx context.Context, input *worklink.DescribeFleetMetadataInput) (*worklink.DescribeFleetMetadataOutput, error) {
	return a.client.DescribeFleetMetadataWithContext(ctx, input)
}

func (a *WorkLinkActivities) DescribeIdentityProviderConfiguration(ctx context.Context, input *worklink.DescribeIdentityProviderConfigurationInput) (*worklink.DescribeIdentityProviderConfigurationOutput, error) {
	return a.client.DescribeIdentityProviderConfigurationWithContext(ctx, input)
}

func (a *WorkLinkActivities) DescribeWebsiteCertificateAuthority(ctx context.Context, input *worklink.DescribeWebsiteCertificateAuthorityInput) (*worklink.DescribeWebsiteCertificateAuthorityOutput, error) {
	return a.client.DescribeWebsiteCertificateAuthorityWithContext(ctx, input)
}

func (a *WorkLinkActivities) DisassociateDomain(ctx context.Context, input *worklink.DisassociateDomainInput) (*worklink.DisassociateDomainOutput, error) {
	return a.client.DisassociateDomainWithContext(ctx, input)
}

func (a *WorkLinkActivities) DisassociateWebsiteAuthorizationProvider(ctx context.Context, input *worklink.DisassociateWebsiteAuthorizationProviderInput) (*worklink.DisassociateWebsiteAuthorizationProviderOutput, error) {
	return a.client.DisassociateWebsiteAuthorizationProviderWithContext(ctx, input)
}

func (a *WorkLinkActivities) DisassociateWebsiteCertificateAuthority(ctx context.Context, input *worklink.DisassociateWebsiteCertificateAuthorityInput) (*worklink.DisassociateWebsiteCertificateAuthorityOutput, error) {
	return a.client.DisassociateWebsiteCertificateAuthorityWithContext(ctx, input)
}

func (a *WorkLinkActivities) ListDevices(ctx context.Context, input *worklink.ListDevicesInput) (*worklink.ListDevicesOutput, error) {
	return a.client.ListDevicesWithContext(ctx, input)
}

func (a *WorkLinkActivities) ListDomains(ctx context.Context, input *worklink.ListDomainsInput) (*worklink.ListDomainsOutput, error) {
	return a.client.ListDomainsWithContext(ctx, input)
}

func (a *WorkLinkActivities) ListFleets(ctx context.Context, input *worklink.ListFleetsInput) (*worklink.ListFleetsOutput, error) {
	return a.client.ListFleetsWithContext(ctx, input)
}

func (a *WorkLinkActivities) ListTagsForResource(ctx context.Context, input *worklink.ListTagsForResourceInput) (*worklink.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *WorkLinkActivities) ListWebsiteAuthorizationProviders(ctx context.Context, input *worklink.ListWebsiteAuthorizationProvidersInput) (*worklink.ListWebsiteAuthorizationProvidersOutput, error) {
	return a.client.ListWebsiteAuthorizationProvidersWithContext(ctx, input)
}

func (a *WorkLinkActivities) ListWebsiteCertificateAuthorities(ctx context.Context, input *worklink.ListWebsiteCertificateAuthoritiesInput) (*worklink.ListWebsiteCertificateAuthoritiesOutput, error) {
	return a.client.ListWebsiteCertificateAuthoritiesWithContext(ctx, input)
}

func (a *WorkLinkActivities) RestoreDomainAccess(ctx context.Context, input *worklink.RestoreDomainAccessInput) (*worklink.RestoreDomainAccessOutput, error) {
	return a.client.RestoreDomainAccessWithContext(ctx, input)
}

func (a *WorkLinkActivities) RevokeDomainAccess(ctx context.Context, input *worklink.RevokeDomainAccessInput) (*worklink.RevokeDomainAccessOutput, error) {
	return a.client.RevokeDomainAccessWithContext(ctx, input)
}

func (a *WorkLinkActivities) SignOutUser(ctx context.Context, input *worklink.SignOutUserInput) (*worklink.SignOutUserOutput, error) {
	return a.client.SignOutUserWithContext(ctx, input)
}

func (a *WorkLinkActivities) TagResource(ctx context.Context, input *worklink.TagResourceInput) (*worklink.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *WorkLinkActivities) UntagResource(ctx context.Context, input *worklink.UntagResourceInput) (*worklink.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *WorkLinkActivities) UpdateAuditStreamConfiguration(ctx context.Context, input *worklink.UpdateAuditStreamConfigurationInput) (*worklink.UpdateAuditStreamConfigurationOutput, error) {
	return a.client.UpdateAuditStreamConfigurationWithContext(ctx, input)
}

func (a *WorkLinkActivities) UpdateCompanyNetworkConfiguration(ctx context.Context, input *worklink.UpdateCompanyNetworkConfigurationInput) (*worklink.UpdateCompanyNetworkConfigurationOutput, error) {
	return a.client.UpdateCompanyNetworkConfigurationWithContext(ctx, input)
}

func (a *WorkLinkActivities) UpdateDevicePolicyConfiguration(ctx context.Context, input *worklink.UpdateDevicePolicyConfigurationInput) (*worklink.UpdateDevicePolicyConfigurationOutput, error) {
	return a.client.UpdateDevicePolicyConfigurationWithContext(ctx, input)
}

func (a *WorkLinkActivities) UpdateDomainMetadata(ctx context.Context, input *worklink.UpdateDomainMetadataInput) (*worklink.UpdateDomainMetadataOutput, error) {
	return a.client.UpdateDomainMetadataWithContext(ctx, input)
}

func (a *WorkLinkActivities) UpdateFleetMetadata(ctx context.Context, input *worklink.UpdateFleetMetadataInput) (*worklink.UpdateFleetMetadataOutput, error) {
	return a.client.UpdateFleetMetadataWithContext(ctx, input)
}

func (a *WorkLinkActivities) UpdateIdentityProviderConfiguration(ctx context.Context, input *worklink.UpdateIdentityProviderConfigurationInput) (*worklink.UpdateIdentityProviderConfigurationOutput, error) {
	return a.client.UpdateIdentityProviderConfigurationWithContext(ctx, input)
}
