package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/worklink"
	"go.temporal.io/sdk/workflow"
)

type WorkLinkClient interface {
    AssociateDomain(ctx workflow.Context, input *worklink.AssociateDomainInput) (*worklink.AssociateDomainOutput, error)
    AssociateDomainAsync(ctx workflow.Context, input *worklink.AssociateDomainInput) *WorklinkAssociateDomainResult

    AssociateWebsiteAuthorizationProvider(ctx workflow.Context, input *worklink.AssociateWebsiteAuthorizationProviderInput) (*worklink.AssociateWebsiteAuthorizationProviderOutput, error)
    AssociateWebsiteAuthorizationProviderAsync(ctx workflow.Context, input *worklink.AssociateWebsiteAuthorizationProviderInput) *WorklinkAssociateWebsiteAuthorizationProviderResult

    AssociateWebsiteCertificateAuthority(ctx workflow.Context, input *worklink.AssociateWebsiteCertificateAuthorityInput) (*worklink.AssociateWebsiteCertificateAuthorityOutput, error)
    AssociateWebsiteCertificateAuthorityAsync(ctx workflow.Context, input *worklink.AssociateWebsiteCertificateAuthorityInput) *WorklinkAssociateWebsiteCertificateAuthorityResult

    CreateFleet(ctx workflow.Context, input *worklink.CreateFleetInput) (*worklink.CreateFleetOutput, error)
    CreateFleetAsync(ctx workflow.Context, input *worklink.CreateFleetInput) *WorklinkCreateFleetResult

    DeleteFleet(ctx workflow.Context, input *worklink.DeleteFleetInput) (*worklink.DeleteFleetOutput, error)
    DeleteFleetAsync(ctx workflow.Context, input *worklink.DeleteFleetInput) *WorklinkDeleteFleetResult

    DescribeAuditStreamConfiguration(ctx workflow.Context, input *worklink.DescribeAuditStreamConfigurationInput) (*worklink.DescribeAuditStreamConfigurationOutput, error)
    DescribeAuditStreamConfigurationAsync(ctx workflow.Context, input *worklink.DescribeAuditStreamConfigurationInput) *WorklinkDescribeAuditStreamConfigurationResult

    DescribeCompanyNetworkConfiguration(ctx workflow.Context, input *worklink.DescribeCompanyNetworkConfigurationInput) (*worklink.DescribeCompanyNetworkConfigurationOutput, error)
    DescribeCompanyNetworkConfigurationAsync(ctx workflow.Context, input *worklink.DescribeCompanyNetworkConfigurationInput) *WorklinkDescribeCompanyNetworkConfigurationResult

    DescribeDevice(ctx workflow.Context, input *worklink.DescribeDeviceInput) (*worklink.DescribeDeviceOutput, error)
    DescribeDeviceAsync(ctx workflow.Context, input *worklink.DescribeDeviceInput) *WorklinkDescribeDeviceResult

    DescribeDevicePolicyConfiguration(ctx workflow.Context, input *worklink.DescribeDevicePolicyConfigurationInput) (*worklink.DescribeDevicePolicyConfigurationOutput, error)
    DescribeDevicePolicyConfigurationAsync(ctx workflow.Context, input *worklink.DescribeDevicePolicyConfigurationInput) *WorklinkDescribeDevicePolicyConfigurationResult

    DescribeDomain(ctx workflow.Context, input *worklink.DescribeDomainInput) (*worklink.DescribeDomainOutput, error)
    DescribeDomainAsync(ctx workflow.Context, input *worklink.DescribeDomainInput) *WorklinkDescribeDomainResult

    DescribeFleetMetadata(ctx workflow.Context, input *worklink.DescribeFleetMetadataInput) (*worklink.DescribeFleetMetadataOutput, error)
    DescribeFleetMetadataAsync(ctx workflow.Context, input *worklink.DescribeFleetMetadataInput) *WorklinkDescribeFleetMetadataResult

    DescribeIdentityProviderConfiguration(ctx workflow.Context, input *worklink.DescribeIdentityProviderConfigurationInput) (*worklink.DescribeIdentityProviderConfigurationOutput, error)
    DescribeIdentityProviderConfigurationAsync(ctx workflow.Context, input *worklink.DescribeIdentityProviderConfigurationInput) *WorklinkDescribeIdentityProviderConfigurationResult

    DescribeWebsiteCertificateAuthority(ctx workflow.Context, input *worklink.DescribeWebsiteCertificateAuthorityInput) (*worklink.DescribeWebsiteCertificateAuthorityOutput, error)
    DescribeWebsiteCertificateAuthorityAsync(ctx workflow.Context, input *worklink.DescribeWebsiteCertificateAuthorityInput) *WorklinkDescribeWebsiteCertificateAuthorityResult

    DisassociateDomain(ctx workflow.Context, input *worklink.DisassociateDomainInput) (*worklink.DisassociateDomainOutput, error)
    DisassociateDomainAsync(ctx workflow.Context, input *worklink.DisassociateDomainInput) *WorklinkDisassociateDomainResult

    DisassociateWebsiteAuthorizationProvider(ctx workflow.Context, input *worklink.DisassociateWebsiteAuthorizationProviderInput) (*worklink.DisassociateWebsiteAuthorizationProviderOutput, error)
    DisassociateWebsiteAuthorizationProviderAsync(ctx workflow.Context, input *worklink.DisassociateWebsiteAuthorizationProviderInput) *WorklinkDisassociateWebsiteAuthorizationProviderResult

    DisassociateWebsiteCertificateAuthority(ctx workflow.Context, input *worklink.DisassociateWebsiteCertificateAuthorityInput) (*worklink.DisassociateWebsiteCertificateAuthorityOutput, error)
    DisassociateWebsiteCertificateAuthorityAsync(ctx workflow.Context, input *worklink.DisassociateWebsiteCertificateAuthorityInput) *WorklinkDisassociateWebsiteCertificateAuthorityResult

    ListDevices(ctx workflow.Context, input *worklink.ListDevicesInput) (*worklink.ListDevicesOutput, error)
    ListDevicesAsync(ctx workflow.Context, input *worklink.ListDevicesInput) *WorklinkListDevicesResult

    ListDomains(ctx workflow.Context, input *worklink.ListDomainsInput) (*worklink.ListDomainsOutput, error)
    ListDomainsAsync(ctx workflow.Context, input *worklink.ListDomainsInput) *WorklinkListDomainsResult

    ListFleets(ctx workflow.Context, input *worklink.ListFleetsInput) (*worklink.ListFleetsOutput, error)
    ListFleetsAsync(ctx workflow.Context, input *worklink.ListFleetsInput) *WorklinkListFleetsResult

    ListTagsForResource(ctx workflow.Context, input *worklink.ListTagsForResourceInput) (*worklink.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *worklink.ListTagsForResourceInput) *WorklinkListTagsForResourceResult

    ListWebsiteAuthorizationProviders(ctx workflow.Context, input *worklink.ListWebsiteAuthorizationProvidersInput) (*worklink.ListWebsiteAuthorizationProvidersOutput, error)
    ListWebsiteAuthorizationProvidersAsync(ctx workflow.Context, input *worklink.ListWebsiteAuthorizationProvidersInput) *WorklinkListWebsiteAuthorizationProvidersResult

    ListWebsiteCertificateAuthorities(ctx workflow.Context, input *worklink.ListWebsiteCertificateAuthoritiesInput) (*worklink.ListWebsiteCertificateAuthoritiesOutput, error)
    ListWebsiteCertificateAuthoritiesAsync(ctx workflow.Context, input *worklink.ListWebsiteCertificateAuthoritiesInput) *WorklinkListWebsiteCertificateAuthoritiesResult

    RestoreDomainAccess(ctx workflow.Context, input *worklink.RestoreDomainAccessInput) (*worklink.RestoreDomainAccessOutput, error)
    RestoreDomainAccessAsync(ctx workflow.Context, input *worklink.RestoreDomainAccessInput) *WorklinkRestoreDomainAccessResult

    RevokeDomainAccess(ctx workflow.Context, input *worklink.RevokeDomainAccessInput) (*worklink.RevokeDomainAccessOutput, error)
    RevokeDomainAccessAsync(ctx workflow.Context, input *worklink.RevokeDomainAccessInput) *WorklinkRevokeDomainAccessResult

    SignOutUser(ctx workflow.Context, input *worklink.SignOutUserInput) (*worklink.SignOutUserOutput, error)
    SignOutUserAsync(ctx workflow.Context, input *worklink.SignOutUserInput) *WorklinkSignOutUserResult

    TagResource(ctx workflow.Context, input *worklink.TagResourceInput) (*worklink.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *worklink.TagResourceInput) *WorklinkTagResourceResult

    UntagResource(ctx workflow.Context, input *worklink.UntagResourceInput) (*worklink.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *worklink.UntagResourceInput) *WorklinkUntagResourceResult

    UpdateAuditStreamConfiguration(ctx workflow.Context, input *worklink.UpdateAuditStreamConfigurationInput) (*worklink.UpdateAuditStreamConfigurationOutput, error)
    UpdateAuditStreamConfigurationAsync(ctx workflow.Context, input *worklink.UpdateAuditStreamConfigurationInput) *WorklinkUpdateAuditStreamConfigurationResult

    UpdateCompanyNetworkConfiguration(ctx workflow.Context, input *worklink.UpdateCompanyNetworkConfigurationInput) (*worklink.UpdateCompanyNetworkConfigurationOutput, error)
    UpdateCompanyNetworkConfigurationAsync(ctx workflow.Context, input *worklink.UpdateCompanyNetworkConfigurationInput) *WorklinkUpdateCompanyNetworkConfigurationResult

    UpdateDevicePolicyConfiguration(ctx workflow.Context, input *worklink.UpdateDevicePolicyConfigurationInput) (*worklink.UpdateDevicePolicyConfigurationOutput, error)
    UpdateDevicePolicyConfigurationAsync(ctx workflow.Context, input *worklink.UpdateDevicePolicyConfigurationInput) *WorklinkUpdateDevicePolicyConfigurationResult

    UpdateDomainMetadata(ctx workflow.Context, input *worklink.UpdateDomainMetadataInput) (*worklink.UpdateDomainMetadataOutput, error)
    UpdateDomainMetadataAsync(ctx workflow.Context, input *worklink.UpdateDomainMetadataInput) *WorklinkUpdateDomainMetadataResult

    UpdateFleetMetadata(ctx workflow.Context, input *worklink.UpdateFleetMetadataInput) (*worklink.UpdateFleetMetadataOutput, error)
    UpdateFleetMetadataAsync(ctx workflow.Context, input *worklink.UpdateFleetMetadataInput) *WorklinkUpdateFleetMetadataResult

    UpdateIdentityProviderConfiguration(ctx workflow.Context, input *worklink.UpdateIdentityProviderConfigurationInput) (*worklink.UpdateIdentityProviderConfigurationOutput, error)
    UpdateIdentityProviderConfigurationAsync(ctx workflow.Context, input *worklink.UpdateIdentityProviderConfigurationInput) *WorklinkUpdateIdentityProviderConfigurationResult
}
type WorklinkAssociateDomainResult struct {
	Result workflow.Future
}

func (r *WorklinkAssociateDomainResult) Get(ctx workflow.Context) (*worklink.AssociateDomainOutput, error) {
    var output worklink.AssociateDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkAssociateWebsiteAuthorizationProviderResult struct {
	Result workflow.Future
}

func (r *WorklinkAssociateWebsiteAuthorizationProviderResult) Get(ctx workflow.Context) (*worklink.AssociateWebsiteAuthorizationProviderOutput, error) {
    var output worklink.AssociateWebsiteAuthorizationProviderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkAssociateWebsiteCertificateAuthorityResult struct {
	Result workflow.Future
}

func (r *WorklinkAssociateWebsiteCertificateAuthorityResult) Get(ctx workflow.Context) (*worklink.AssociateWebsiteCertificateAuthorityOutput, error) {
    var output worklink.AssociateWebsiteCertificateAuthorityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkCreateFleetResult struct {
	Result workflow.Future
}

func (r *WorklinkCreateFleetResult) Get(ctx workflow.Context) (*worklink.CreateFleetOutput, error) {
    var output worklink.CreateFleetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkDeleteFleetResult struct {
	Result workflow.Future
}

func (r *WorklinkDeleteFleetResult) Get(ctx workflow.Context) (*worklink.DeleteFleetOutput, error) {
    var output worklink.DeleteFleetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkDescribeAuditStreamConfigurationResult struct {
	Result workflow.Future
}

func (r *WorklinkDescribeAuditStreamConfigurationResult) Get(ctx workflow.Context) (*worklink.DescribeAuditStreamConfigurationOutput, error) {
    var output worklink.DescribeAuditStreamConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkDescribeCompanyNetworkConfigurationResult struct {
	Result workflow.Future
}

func (r *WorklinkDescribeCompanyNetworkConfigurationResult) Get(ctx workflow.Context) (*worklink.DescribeCompanyNetworkConfigurationOutput, error) {
    var output worklink.DescribeCompanyNetworkConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkDescribeDeviceResult struct {
	Result workflow.Future
}

func (r *WorklinkDescribeDeviceResult) Get(ctx workflow.Context) (*worklink.DescribeDeviceOutput, error) {
    var output worklink.DescribeDeviceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkDescribeDevicePolicyConfigurationResult struct {
	Result workflow.Future
}

func (r *WorklinkDescribeDevicePolicyConfigurationResult) Get(ctx workflow.Context) (*worklink.DescribeDevicePolicyConfigurationOutput, error) {
    var output worklink.DescribeDevicePolicyConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkDescribeDomainResult struct {
	Result workflow.Future
}

func (r *WorklinkDescribeDomainResult) Get(ctx workflow.Context) (*worklink.DescribeDomainOutput, error) {
    var output worklink.DescribeDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkDescribeFleetMetadataResult struct {
	Result workflow.Future
}

func (r *WorklinkDescribeFleetMetadataResult) Get(ctx workflow.Context) (*worklink.DescribeFleetMetadataOutput, error) {
    var output worklink.DescribeFleetMetadataOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkDescribeIdentityProviderConfigurationResult struct {
	Result workflow.Future
}

func (r *WorklinkDescribeIdentityProviderConfigurationResult) Get(ctx workflow.Context) (*worklink.DescribeIdentityProviderConfigurationOutput, error) {
    var output worklink.DescribeIdentityProviderConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkDescribeWebsiteCertificateAuthorityResult struct {
	Result workflow.Future
}

func (r *WorklinkDescribeWebsiteCertificateAuthorityResult) Get(ctx workflow.Context) (*worklink.DescribeWebsiteCertificateAuthorityOutput, error) {
    var output worklink.DescribeWebsiteCertificateAuthorityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkDisassociateDomainResult struct {
	Result workflow.Future
}

func (r *WorklinkDisassociateDomainResult) Get(ctx workflow.Context) (*worklink.DisassociateDomainOutput, error) {
    var output worklink.DisassociateDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkDisassociateWebsiteAuthorizationProviderResult struct {
	Result workflow.Future
}

func (r *WorklinkDisassociateWebsiteAuthorizationProviderResult) Get(ctx workflow.Context) (*worklink.DisassociateWebsiteAuthorizationProviderOutput, error) {
    var output worklink.DisassociateWebsiteAuthorizationProviderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkDisassociateWebsiteCertificateAuthorityResult struct {
	Result workflow.Future
}

func (r *WorklinkDisassociateWebsiteCertificateAuthorityResult) Get(ctx workflow.Context) (*worklink.DisassociateWebsiteCertificateAuthorityOutput, error) {
    var output worklink.DisassociateWebsiteCertificateAuthorityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkListDevicesResult struct {
	Result workflow.Future
}

func (r *WorklinkListDevicesResult) Get(ctx workflow.Context) (*worklink.ListDevicesOutput, error) {
    var output worklink.ListDevicesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkListDomainsResult struct {
	Result workflow.Future
}

func (r *WorklinkListDomainsResult) Get(ctx workflow.Context) (*worklink.ListDomainsOutput, error) {
    var output worklink.ListDomainsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkListFleetsResult struct {
	Result workflow.Future
}

func (r *WorklinkListFleetsResult) Get(ctx workflow.Context) (*worklink.ListFleetsOutput, error) {
    var output worklink.ListFleetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *WorklinkListTagsForResourceResult) Get(ctx workflow.Context) (*worklink.ListTagsForResourceOutput, error) {
    var output worklink.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkListWebsiteAuthorizationProvidersResult struct {
	Result workflow.Future
}

func (r *WorklinkListWebsiteAuthorizationProvidersResult) Get(ctx workflow.Context) (*worklink.ListWebsiteAuthorizationProvidersOutput, error) {
    var output worklink.ListWebsiteAuthorizationProvidersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkListWebsiteCertificateAuthoritiesResult struct {
	Result workflow.Future
}

func (r *WorklinkListWebsiteCertificateAuthoritiesResult) Get(ctx workflow.Context) (*worklink.ListWebsiteCertificateAuthoritiesOutput, error) {
    var output worklink.ListWebsiteCertificateAuthoritiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkRestoreDomainAccessResult struct {
	Result workflow.Future
}

func (r *WorklinkRestoreDomainAccessResult) Get(ctx workflow.Context) (*worklink.RestoreDomainAccessOutput, error) {
    var output worklink.RestoreDomainAccessOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkRevokeDomainAccessResult struct {
	Result workflow.Future
}

func (r *WorklinkRevokeDomainAccessResult) Get(ctx workflow.Context) (*worklink.RevokeDomainAccessOutput, error) {
    var output worklink.RevokeDomainAccessOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkSignOutUserResult struct {
	Result workflow.Future
}

func (r *WorklinkSignOutUserResult) Get(ctx workflow.Context) (*worklink.SignOutUserOutput, error) {
    var output worklink.SignOutUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkTagResourceResult struct {
	Result workflow.Future
}

func (r *WorklinkTagResourceResult) Get(ctx workflow.Context) (*worklink.TagResourceOutput, error) {
    var output worklink.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkUntagResourceResult struct {
	Result workflow.Future
}

func (r *WorklinkUntagResourceResult) Get(ctx workflow.Context) (*worklink.UntagResourceOutput, error) {
    var output worklink.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkUpdateAuditStreamConfigurationResult struct {
	Result workflow.Future
}

func (r *WorklinkUpdateAuditStreamConfigurationResult) Get(ctx workflow.Context) (*worklink.UpdateAuditStreamConfigurationOutput, error) {
    var output worklink.UpdateAuditStreamConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkUpdateCompanyNetworkConfigurationResult struct {
	Result workflow.Future
}

func (r *WorklinkUpdateCompanyNetworkConfigurationResult) Get(ctx workflow.Context) (*worklink.UpdateCompanyNetworkConfigurationOutput, error) {
    var output worklink.UpdateCompanyNetworkConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkUpdateDevicePolicyConfigurationResult struct {
	Result workflow.Future
}

func (r *WorklinkUpdateDevicePolicyConfigurationResult) Get(ctx workflow.Context) (*worklink.UpdateDevicePolicyConfigurationOutput, error) {
    var output worklink.UpdateDevicePolicyConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkUpdateDomainMetadataResult struct {
	Result workflow.Future
}

func (r *WorklinkUpdateDomainMetadataResult) Get(ctx workflow.Context) (*worklink.UpdateDomainMetadataOutput, error) {
    var output worklink.UpdateDomainMetadataOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkUpdateFleetMetadataResult struct {
	Result workflow.Future
}

func (r *WorklinkUpdateFleetMetadataResult) Get(ctx workflow.Context) (*worklink.UpdateFleetMetadataOutput, error) {
    var output worklink.UpdateFleetMetadataOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorklinkUpdateIdentityProviderConfigurationResult struct {
	Result workflow.Future
}

func (r *WorklinkUpdateIdentityProviderConfigurationResult) Get(ctx workflow.Context) (*worklink.UpdateIdentityProviderConfigurationOutput, error) {
    var output worklink.UpdateIdentityProviderConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type WorkLinkStub struct {
    activities WorkLinkClient
}

func NewWorkLinkStub() WorkLinkClient {
    return &WorkLinkStub{}
}

func (a *WorkLinkStub) AssociateDomain(ctx workflow.Context, input *worklink.AssociateDomainInput) (*worklink.AssociateDomainOutput, error) {
    var output worklink.AssociateDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) AssociateWebsiteAuthorizationProvider(ctx workflow.Context, input *worklink.AssociateWebsiteAuthorizationProviderInput) (*worklink.AssociateWebsiteAuthorizationProviderOutput, error) {
    var output worklink.AssociateWebsiteAuthorizationProviderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateWebsiteAuthorizationProvider, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) AssociateWebsiteCertificateAuthority(ctx workflow.Context, input *worklink.AssociateWebsiteCertificateAuthorityInput) (*worklink.AssociateWebsiteCertificateAuthorityOutput, error) {
    var output worklink.AssociateWebsiteCertificateAuthorityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateWebsiteCertificateAuthority, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) CreateFleet(ctx workflow.Context, input *worklink.CreateFleetInput) (*worklink.CreateFleetOutput, error) {
    var output worklink.CreateFleetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFleet, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) DeleteFleet(ctx workflow.Context, input *worklink.DeleteFleetInput) (*worklink.DeleteFleetOutput, error) {
    var output worklink.DeleteFleetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFleet, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) DescribeAuditStreamConfiguration(ctx workflow.Context, input *worklink.DescribeAuditStreamConfigurationInput) (*worklink.DescribeAuditStreamConfigurationOutput, error) {
    var output worklink.DescribeAuditStreamConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAuditStreamConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) DescribeCompanyNetworkConfiguration(ctx workflow.Context, input *worklink.DescribeCompanyNetworkConfigurationInput) (*worklink.DescribeCompanyNetworkConfigurationOutput, error) {
    var output worklink.DescribeCompanyNetworkConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCompanyNetworkConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) DescribeDevice(ctx workflow.Context, input *worklink.DescribeDeviceInput) (*worklink.DescribeDeviceOutput, error) {
    var output worklink.DescribeDeviceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDevice, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) DescribeDevicePolicyConfiguration(ctx workflow.Context, input *worklink.DescribeDevicePolicyConfigurationInput) (*worklink.DescribeDevicePolicyConfigurationOutput, error) {
    var output worklink.DescribeDevicePolicyConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDevicePolicyConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) DescribeDomain(ctx workflow.Context, input *worklink.DescribeDomainInput) (*worklink.DescribeDomainOutput, error) {
    var output worklink.DescribeDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) DescribeFleetMetadata(ctx workflow.Context, input *worklink.DescribeFleetMetadataInput) (*worklink.DescribeFleetMetadataOutput, error) {
    var output worklink.DescribeFleetMetadataOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFleetMetadata, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) DescribeIdentityProviderConfiguration(ctx workflow.Context, input *worklink.DescribeIdentityProviderConfigurationInput) (*worklink.DescribeIdentityProviderConfigurationOutput, error) {
    var output worklink.DescribeIdentityProviderConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeIdentityProviderConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) DescribeWebsiteCertificateAuthority(ctx workflow.Context, input *worklink.DescribeWebsiteCertificateAuthorityInput) (*worklink.DescribeWebsiteCertificateAuthorityOutput, error) {
    var output worklink.DescribeWebsiteCertificateAuthorityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeWebsiteCertificateAuthority, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) DisassociateDomain(ctx workflow.Context, input *worklink.DisassociateDomainInput) (*worklink.DisassociateDomainOutput, error) {
    var output worklink.DisassociateDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) DisassociateWebsiteAuthorizationProvider(ctx workflow.Context, input *worklink.DisassociateWebsiteAuthorizationProviderInput) (*worklink.DisassociateWebsiteAuthorizationProviderOutput, error) {
    var output worklink.DisassociateWebsiteAuthorizationProviderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateWebsiteAuthorizationProvider, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) DisassociateWebsiteCertificateAuthority(ctx workflow.Context, input *worklink.DisassociateWebsiteCertificateAuthorityInput) (*worklink.DisassociateWebsiteCertificateAuthorityOutput, error) {
    var output worklink.DisassociateWebsiteCertificateAuthorityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateWebsiteCertificateAuthority, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) ListDevices(ctx workflow.Context, input *worklink.ListDevicesInput) (*worklink.ListDevicesOutput, error) {
    var output worklink.ListDevicesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDevices, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) ListDomains(ctx workflow.Context, input *worklink.ListDomainsInput) (*worklink.ListDomainsOutput, error) {
    var output worklink.ListDomainsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDomains, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) ListFleets(ctx workflow.Context, input *worklink.ListFleetsInput) (*worklink.ListFleetsOutput, error) {
    var output worklink.ListFleetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFleets, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) ListTagsForResource(ctx workflow.Context, input *worklink.ListTagsForResourceInput) (*worklink.ListTagsForResourceOutput, error) {
    var output worklink.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) ListWebsiteAuthorizationProviders(ctx workflow.Context, input *worklink.ListWebsiteAuthorizationProvidersInput) (*worklink.ListWebsiteAuthorizationProvidersOutput, error) {
    var output worklink.ListWebsiteAuthorizationProvidersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListWebsiteAuthorizationProviders, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) ListWebsiteCertificateAuthorities(ctx workflow.Context, input *worklink.ListWebsiteCertificateAuthoritiesInput) (*worklink.ListWebsiteCertificateAuthoritiesOutput, error) {
    var output worklink.ListWebsiteCertificateAuthoritiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListWebsiteCertificateAuthorities, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) RestoreDomainAccess(ctx workflow.Context, input *worklink.RestoreDomainAccessInput) (*worklink.RestoreDomainAccessOutput, error) {
    var output worklink.RestoreDomainAccessOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RestoreDomainAccess, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) RevokeDomainAccess(ctx workflow.Context, input *worklink.RevokeDomainAccessInput) (*worklink.RevokeDomainAccessOutput, error) {
    var output worklink.RevokeDomainAccessOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RevokeDomainAccess, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) SignOutUser(ctx workflow.Context, input *worklink.SignOutUserInput) (*worklink.SignOutUserOutput, error) {
    var output worklink.SignOutUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SignOutUser, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) TagResource(ctx workflow.Context, input *worklink.TagResourceInput) (*worklink.TagResourceOutput, error) {
    var output worklink.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) UntagResource(ctx workflow.Context, input *worklink.UntagResourceInput) (*worklink.UntagResourceOutput, error) {
    var output worklink.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) UpdateAuditStreamConfiguration(ctx workflow.Context, input *worklink.UpdateAuditStreamConfigurationInput) (*worklink.UpdateAuditStreamConfigurationOutput, error) {
    var output worklink.UpdateAuditStreamConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAuditStreamConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) UpdateCompanyNetworkConfiguration(ctx workflow.Context, input *worklink.UpdateCompanyNetworkConfigurationInput) (*worklink.UpdateCompanyNetworkConfigurationOutput, error) {
    var output worklink.UpdateCompanyNetworkConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateCompanyNetworkConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) UpdateDevicePolicyConfiguration(ctx workflow.Context, input *worklink.UpdateDevicePolicyConfigurationInput) (*worklink.UpdateDevicePolicyConfigurationOutput, error) {
    var output worklink.UpdateDevicePolicyConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDevicePolicyConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) UpdateDomainMetadata(ctx workflow.Context, input *worklink.UpdateDomainMetadataInput) (*worklink.UpdateDomainMetadataOutput, error) {
    var output worklink.UpdateDomainMetadataOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDomainMetadata, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) UpdateFleetMetadata(ctx workflow.Context, input *worklink.UpdateFleetMetadataInput) (*worklink.UpdateFleetMetadataOutput, error) {
    var output worklink.UpdateFleetMetadataOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFleetMetadata, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkLinkStub) UpdateIdentityProviderConfiguration(ctx workflow.Context, input *worklink.UpdateIdentityProviderConfigurationInput) (*worklink.UpdateIdentityProviderConfigurationOutput, error) {
    var output worklink.UpdateIdentityProviderConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateIdentityProviderConfiguration, input).Get(ctx, &output)
    return &output, err
}
