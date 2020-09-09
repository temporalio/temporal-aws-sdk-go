package awsclients

import (
	"github.com/aws/aws-sdk-go/service/networkmanager"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type NetworkManagerClient interface {
       AssociateCustomerGateway(ctx workflow.Context, input *networkmanager.AssociateCustomerGatewayInput) (*networkmanager.AssociateCustomerGatewayOutput, error)
       AssociateCustomerGatewayAsync(ctx workflow.Context, input *networkmanager.AssociateCustomerGatewayInput) *NetworkmanagerAssociateCustomerGatewayResult

       AssociateLink(ctx workflow.Context, input *networkmanager.AssociateLinkInput) (*networkmanager.AssociateLinkOutput, error)
       AssociateLinkAsync(ctx workflow.Context, input *networkmanager.AssociateLinkInput) *NetworkmanagerAssociateLinkResult

       CreateDevice(ctx workflow.Context, input *networkmanager.CreateDeviceInput) (*networkmanager.CreateDeviceOutput, error)
       CreateDeviceAsync(ctx workflow.Context, input *networkmanager.CreateDeviceInput) *NetworkmanagerCreateDeviceResult

       CreateGlobalNetwork(ctx workflow.Context, input *networkmanager.CreateGlobalNetworkInput) (*networkmanager.CreateGlobalNetworkOutput, error)
       CreateGlobalNetworkAsync(ctx workflow.Context, input *networkmanager.CreateGlobalNetworkInput) *NetworkmanagerCreateGlobalNetworkResult

       CreateLink(ctx workflow.Context, input *networkmanager.CreateLinkInput) (*networkmanager.CreateLinkOutput, error)
       CreateLinkAsync(ctx workflow.Context, input *networkmanager.CreateLinkInput) *NetworkmanagerCreateLinkResult

       CreateSite(ctx workflow.Context, input *networkmanager.CreateSiteInput) (*networkmanager.CreateSiteOutput, error)
       CreateSiteAsync(ctx workflow.Context, input *networkmanager.CreateSiteInput) *NetworkmanagerCreateSiteResult

       DeleteDevice(ctx workflow.Context, input *networkmanager.DeleteDeviceInput) (*networkmanager.DeleteDeviceOutput, error)
       DeleteDeviceAsync(ctx workflow.Context, input *networkmanager.DeleteDeviceInput) *NetworkmanagerDeleteDeviceResult

       DeleteGlobalNetwork(ctx workflow.Context, input *networkmanager.DeleteGlobalNetworkInput) (*networkmanager.DeleteGlobalNetworkOutput, error)
       DeleteGlobalNetworkAsync(ctx workflow.Context, input *networkmanager.DeleteGlobalNetworkInput) *NetworkmanagerDeleteGlobalNetworkResult

       DeleteLink(ctx workflow.Context, input *networkmanager.DeleteLinkInput) (*networkmanager.DeleteLinkOutput, error)
       DeleteLinkAsync(ctx workflow.Context, input *networkmanager.DeleteLinkInput) *NetworkmanagerDeleteLinkResult

       DeleteSite(ctx workflow.Context, input *networkmanager.DeleteSiteInput) (*networkmanager.DeleteSiteOutput, error)
       DeleteSiteAsync(ctx workflow.Context, input *networkmanager.DeleteSiteInput) *NetworkmanagerDeleteSiteResult

       DeregisterTransitGateway(ctx workflow.Context, input *networkmanager.DeregisterTransitGatewayInput) (*networkmanager.DeregisterTransitGatewayOutput, error)
       DeregisterTransitGatewayAsync(ctx workflow.Context, input *networkmanager.DeregisterTransitGatewayInput) *NetworkmanagerDeregisterTransitGatewayResult

       DescribeGlobalNetworks(ctx workflow.Context, input *networkmanager.DescribeGlobalNetworksInput) (*networkmanager.DescribeGlobalNetworksOutput, error)
       DescribeGlobalNetworksAsync(ctx workflow.Context, input *networkmanager.DescribeGlobalNetworksInput) *NetworkmanagerDescribeGlobalNetworksResult

       DisassociateCustomerGateway(ctx workflow.Context, input *networkmanager.DisassociateCustomerGatewayInput) (*networkmanager.DisassociateCustomerGatewayOutput, error)
       DisassociateCustomerGatewayAsync(ctx workflow.Context, input *networkmanager.DisassociateCustomerGatewayInput) *NetworkmanagerDisassociateCustomerGatewayResult

       DisassociateLink(ctx workflow.Context, input *networkmanager.DisassociateLinkInput) (*networkmanager.DisassociateLinkOutput, error)
       DisassociateLinkAsync(ctx workflow.Context, input *networkmanager.DisassociateLinkInput) *NetworkmanagerDisassociateLinkResult

       GetCustomerGatewayAssociations(ctx workflow.Context, input *networkmanager.GetCustomerGatewayAssociationsInput) (*networkmanager.GetCustomerGatewayAssociationsOutput, error)
       GetCustomerGatewayAssociationsAsync(ctx workflow.Context, input *networkmanager.GetCustomerGatewayAssociationsInput) *NetworkmanagerGetCustomerGatewayAssociationsResult

       GetDevices(ctx workflow.Context, input *networkmanager.GetDevicesInput) (*networkmanager.GetDevicesOutput, error)
       GetDevicesAsync(ctx workflow.Context, input *networkmanager.GetDevicesInput) *NetworkmanagerGetDevicesResult

       GetLinkAssociations(ctx workflow.Context, input *networkmanager.GetLinkAssociationsInput) (*networkmanager.GetLinkAssociationsOutput, error)
       GetLinkAssociationsAsync(ctx workflow.Context, input *networkmanager.GetLinkAssociationsInput) *NetworkmanagerGetLinkAssociationsResult

       GetLinks(ctx workflow.Context, input *networkmanager.GetLinksInput) (*networkmanager.GetLinksOutput, error)
       GetLinksAsync(ctx workflow.Context, input *networkmanager.GetLinksInput) *NetworkmanagerGetLinksResult

       GetSites(ctx workflow.Context, input *networkmanager.GetSitesInput) (*networkmanager.GetSitesOutput, error)
       GetSitesAsync(ctx workflow.Context, input *networkmanager.GetSitesInput) *NetworkmanagerGetSitesResult

       GetTransitGatewayRegistrations(ctx workflow.Context, input *networkmanager.GetTransitGatewayRegistrationsInput) (*networkmanager.GetTransitGatewayRegistrationsOutput, error)
       GetTransitGatewayRegistrationsAsync(ctx workflow.Context, input *networkmanager.GetTransitGatewayRegistrationsInput) *NetworkmanagerGetTransitGatewayRegistrationsResult

       ListTagsForResource(ctx workflow.Context, input *networkmanager.ListTagsForResourceInput) (*networkmanager.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *networkmanager.ListTagsForResourceInput) *NetworkmanagerListTagsForResourceResult

       RegisterTransitGateway(ctx workflow.Context, input *networkmanager.RegisterTransitGatewayInput) (*networkmanager.RegisterTransitGatewayOutput, error)
       RegisterTransitGatewayAsync(ctx workflow.Context, input *networkmanager.RegisterTransitGatewayInput) *NetworkmanagerRegisterTransitGatewayResult

       TagResource(ctx workflow.Context, input *networkmanager.TagResourceInput) (*networkmanager.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *networkmanager.TagResourceInput) *NetworkmanagerTagResourceResult

       UntagResource(ctx workflow.Context, input *networkmanager.UntagResourceInput) (*networkmanager.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *networkmanager.UntagResourceInput) *NetworkmanagerUntagResourceResult

       UpdateDevice(ctx workflow.Context, input *networkmanager.UpdateDeviceInput) (*networkmanager.UpdateDeviceOutput, error)
       UpdateDeviceAsync(ctx workflow.Context, input *networkmanager.UpdateDeviceInput) *NetworkmanagerUpdateDeviceResult

       UpdateGlobalNetwork(ctx workflow.Context, input *networkmanager.UpdateGlobalNetworkInput) (*networkmanager.UpdateGlobalNetworkOutput, error)
       UpdateGlobalNetworkAsync(ctx workflow.Context, input *networkmanager.UpdateGlobalNetworkInput) *NetworkmanagerUpdateGlobalNetworkResult

       UpdateLink(ctx workflow.Context, input *networkmanager.UpdateLinkInput) (*networkmanager.UpdateLinkOutput, error)
       UpdateLinkAsync(ctx workflow.Context, input *networkmanager.UpdateLinkInput) *NetworkmanagerUpdateLinkResult

       UpdateSite(ctx workflow.Context, input *networkmanager.UpdateSiteInput) (*networkmanager.UpdateSiteOutput, error)
       UpdateSiteAsync(ctx workflow.Context, input *networkmanager.UpdateSiteInput) *NetworkmanagerUpdateSiteResult
}

type NetworkmanagerAssociateCustomerGatewayResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerAssociateCustomerGatewayResult) Get(ctx workflow.Context) (*networkmanager.AssociateCustomerGatewayOutput, error) {
    var output networkmanager.AssociateCustomerGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerAssociateLinkResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerAssociateLinkResult) Get(ctx workflow.Context) (*networkmanager.AssociateLinkOutput, error) {
    var output networkmanager.AssociateLinkOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerCreateDeviceResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerCreateDeviceResult) Get(ctx workflow.Context) (*networkmanager.CreateDeviceOutput, error) {
    var output networkmanager.CreateDeviceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerCreateGlobalNetworkResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerCreateGlobalNetworkResult) Get(ctx workflow.Context) (*networkmanager.CreateGlobalNetworkOutput, error) {
    var output networkmanager.CreateGlobalNetworkOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerCreateLinkResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerCreateLinkResult) Get(ctx workflow.Context) (*networkmanager.CreateLinkOutput, error) {
    var output networkmanager.CreateLinkOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerCreateSiteResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerCreateSiteResult) Get(ctx workflow.Context) (*networkmanager.CreateSiteOutput, error) {
    var output networkmanager.CreateSiteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerDeleteDeviceResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerDeleteDeviceResult) Get(ctx workflow.Context) (*networkmanager.DeleteDeviceOutput, error) {
    var output networkmanager.DeleteDeviceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerDeleteGlobalNetworkResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerDeleteGlobalNetworkResult) Get(ctx workflow.Context) (*networkmanager.DeleteGlobalNetworkOutput, error) {
    var output networkmanager.DeleteGlobalNetworkOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerDeleteLinkResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerDeleteLinkResult) Get(ctx workflow.Context) (*networkmanager.DeleteLinkOutput, error) {
    var output networkmanager.DeleteLinkOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerDeleteSiteResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerDeleteSiteResult) Get(ctx workflow.Context) (*networkmanager.DeleteSiteOutput, error) {
    var output networkmanager.DeleteSiteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerDeregisterTransitGatewayResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerDeregisterTransitGatewayResult) Get(ctx workflow.Context) (*networkmanager.DeregisterTransitGatewayOutput, error) {
    var output networkmanager.DeregisterTransitGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerDescribeGlobalNetworksResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerDescribeGlobalNetworksResult) Get(ctx workflow.Context) (*networkmanager.DescribeGlobalNetworksOutput, error) {
    var output networkmanager.DescribeGlobalNetworksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerDisassociateCustomerGatewayResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerDisassociateCustomerGatewayResult) Get(ctx workflow.Context) (*networkmanager.DisassociateCustomerGatewayOutput, error) {
    var output networkmanager.DisassociateCustomerGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerDisassociateLinkResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerDisassociateLinkResult) Get(ctx workflow.Context) (*networkmanager.DisassociateLinkOutput, error) {
    var output networkmanager.DisassociateLinkOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerGetCustomerGatewayAssociationsResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerGetCustomerGatewayAssociationsResult) Get(ctx workflow.Context) (*networkmanager.GetCustomerGatewayAssociationsOutput, error) {
    var output networkmanager.GetCustomerGatewayAssociationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerGetDevicesResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerGetDevicesResult) Get(ctx workflow.Context) (*networkmanager.GetDevicesOutput, error) {
    var output networkmanager.GetDevicesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerGetLinkAssociationsResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerGetLinkAssociationsResult) Get(ctx workflow.Context) (*networkmanager.GetLinkAssociationsOutput, error) {
    var output networkmanager.GetLinkAssociationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerGetLinksResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerGetLinksResult) Get(ctx workflow.Context) (*networkmanager.GetLinksOutput, error) {
    var output networkmanager.GetLinksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerGetSitesResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerGetSitesResult) Get(ctx workflow.Context) (*networkmanager.GetSitesOutput, error) {
    var output networkmanager.GetSitesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerGetTransitGatewayRegistrationsResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerGetTransitGatewayRegistrationsResult) Get(ctx workflow.Context) (*networkmanager.GetTransitGatewayRegistrationsOutput, error) {
    var output networkmanager.GetTransitGatewayRegistrationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerListTagsForResourceResult) Get(ctx workflow.Context) (*networkmanager.ListTagsForResourceOutput, error) {
    var output networkmanager.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerRegisterTransitGatewayResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerRegisterTransitGatewayResult) Get(ctx workflow.Context) (*networkmanager.RegisterTransitGatewayOutput, error) {
    var output networkmanager.RegisterTransitGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerTagResourceResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerTagResourceResult) Get(ctx workflow.Context) (*networkmanager.TagResourceOutput, error) {
    var output networkmanager.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerUntagResourceResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerUntagResourceResult) Get(ctx workflow.Context) (*networkmanager.UntagResourceOutput, error) {
    var output networkmanager.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerUpdateDeviceResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerUpdateDeviceResult) Get(ctx workflow.Context) (*networkmanager.UpdateDeviceOutput, error) {
    var output networkmanager.UpdateDeviceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerUpdateGlobalNetworkResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerUpdateGlobalNetworkResult) Get(ctx workflow.Context) (*networkmanager.UpdateGlobalNetworkOutput, error) {
    var output networkmanager.UpdateGlobalNetworkOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerUpdateLinkResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerUpdateLinkResult) Get(ctx workflow.Context) (*networkmanager.UpdateLinkOutput, error) {
    var output networkmanager.UpdateLinkOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkmanagerUpdateSiteResult struct {
	Result workflow.Future
}

func (r *NetworkmanagerUpdateSiteResult) Get(ctx workflow.Context) (*networkmanager.UpdateSiteOutput, error) {
    var output networkmanager.UpdateSiteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type NetworkManagerStub struct {
    activities awsactivities.NetworkManagerActivities
}

func NewNetworkManagerStub() NetworkManagerClient {
    return &NetworkManagerStub{}
}

func (a *NetworkManagerStub) AssociateCustomerGateway(ctx workflow.Context, input *networkmanager.AssociateCustomerGatewayInput) (*networkmanager.AssociateCustomerGatewayOutput, error) {
    var output networkmanager.AssociateCustomerGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateCustomerGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) AssociateCustomerGatewayAsync(ctx workflow.Context, input *networkmanager.AssociateCustomerGatewayInput) *NetworkmanagerAssociateCustomerGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateCustomerGateway, input)
    return &NetworkmanagerAssociateCustomerGatewayResult{Result: future}
}

func (a *NetworkManagerStub) AssociateLink(ctx workflow.Context, input *networkmanager.AssociateLinkInput) (*networkmanager.AssociateLinkOutput, error) {
    var output networkmanager.AssociateLinkOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateLink, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) AssociateLinkAsync(ctx workflow.Context, input *networkmanager.AssociateLinkInput) *NetworkmanagerAssociateLinkResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateLink, input)
    return &NetworkmanagerAssociateLinkResult{Result: future}
}

func (a *NetworkManagerStub) CreateDevice(ctx workflow.Context, input *networkmanager.CreateDeviceInput) (*networkmanager.CreateDeviceOutput, error) {
    var output networkmanager.CreateDeviceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDevice, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) CreateDeviceAsync(ctx workflow.Context, input *networkmanager.CreateDeviceInput) *NetworkmanagerCreateDeviceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDevice, input)
    return &NetworkmanagerCreateDeviceResult{Result: future}
}

func (a *NetworkManagerStub) CreateGlobalNetwork(ctx workflow.Context, input *networkmanager.CreateGlobalNetworkInput) (*networkmanager.CreateGlobalNetworkOutput, error) {
    var output networkmanager.CreateGlobalNetworkOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateGlobalNetwork, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) CreateGlobalNetworkAsync(ctx workflow.Context, input *networkmanager.CreateGlobalNetworkInput) *NetworkmanagerCreateGlobalNetworkResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateGlobalNetwork, input)
    return &NetworkmanagerCreateGlobalNetworkResult{Result: future}
}

func (a *NetworkManagerStub) CreateLink(ctx workflow.Context, input *networkmanager.CreateLinkInput) (*networkmanager.CreateLinkOutput, error) {
    var output networkmanager.CreateLinkOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLink, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) CreateLinkAsync(ctx workflow.Context, input *networkmanager.CreateLinkInput) *NetworkmanagerCreateLinkResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateLink, input)
    return &NetworkmanagerCreateLinkResult{Result: future}
}

func (a *NetworkManagerStub) CreateSite(ctx workflow.Context, input *networkmanager.CreateSiteInput) (*networkmanager.CreateSiteOutput, error) {
    var output networkmanager.CreateSiteOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSite, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) CreateSiteAsync(ctx workflow.Context, input *networkmanager.CreateSiteInput) *NetworkmanagerCreateSiteResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSite, input)
    return &NetworkmanagerCreateSiteResult{Result: future}
}

func (a *NetworkManagerStub) DeleteDevice(ctx workflow.Context, input *networkmanager.DeleteDeviceInput) (*networkmanager.DeleteDeviceOutput, error) {
    var output networkmanager.DeleteDeviceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDevice, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) DeleteDeviceAsync(ctx workflow.Context, input *networkmanager.DeleteDeviceInput) *NetworkmanagerDeleteDeviceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDevice, input)
    return &NetworkmanagerDeleteDeviceResult{Result: future}
}

func (a *NetworkManagerStub) DeleteGlobalNetwork(ctx workflow.Context, input *networkmanager.DeleteGlobalNetworkInput) (*networkmanager.DeleteGlobalNetworkOutput, error) {
    var output networkmanager.DeleteGlobalNetworkOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteGlobalNetwork, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) DeleteGlobalNetworkAsync(ctx workflow.Context, input *networkmanager.DeleteGlobalNetworkInput) *NetworkmanagerDeleteGlobalNetworkResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteGlobalNetwork, input)
    return &NetworkmanagerDeleteGlobalNetworkResult{Result: future}
}

func (a *NetworkManagerStub) DeleteLink(ctx workflow.Context, input *networkmanager.DeleteLinkInput) (*networkmanager.DeleteLinkOutput, error) {
    var output networkmanager.DeleteLinkOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteLink, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) DeleteLinkAsync(ctx workflow.Context, input *networkmanager.DeleteLinkInput) *NetworkmanagerDeleteLinkResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteLink, input)
    return &NetworkmanagerDeleteLinkResult{Result: future}
}

func (a *NetworkManagerStub) DeleteSite(ctx workflow.Context, input *networkmanager.DeleteSiteInput) (*networkmanager.DeleteSiteOutput, error) {
    var output networkmanager.DeleteSiteOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSite, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) DeleteSiteAsync(ctx workflow.Context, input *networkmanager.DeleteSiteInput) *NetworkmanagerDeleteSiteResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSite, input)
    return &NetworkmanagerDeleteSiteResult{Result: future}
}

func (a *NetworkManagerStub) DeregisterTransitGateway(ctx workflow.Context, input *networkmanager.DeregisterTransitGatewayInput) (*networkmanager.DeregisterTransitGatewayOutput, error) {
    var output networkmanager.DeregisterTransitGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterTransitGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) DeregisterTransitGatewayAsync(ctx workflow.Context, input *networkmanager.DeregisterTransitGatewayInput) *NetworkmanagerDeregisterTransitGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeregisterTransitGateway, input)
    return &NetworkmanagerDeregisterTransitGatewayResult{Result: future}
}

func (a *NetworkManagerStub) DescribeGlobalNetworks(ctx workflow.Context, input *networkmanager.DescribeGlobalNetworksInput) (*networkmanager.DescribeGlobalNetworksOutput, error) {
    var output networkmanager.DescribeGlobalNetworksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeGlobalNetworks, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) DescribeGlobalNetworksAsync(ctx workflow.Context, input *networkmanager.DescribeGlobalNetworksInput) *NetworkmanagerDescribeGlobalNetworksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeGlobalNetworks, input)
    return &NetworkmanagerDescribeGlobalNetworksResult{Result: future}
}

func (a *NetworkManagerStub) DisassociateCustomerGateway(ctx workflow.Context, input *networkmanager.DisassociateCustomerGatewayInput) (*networkmanager.DisassociateCustomerGatewayOutput, error) {
    var output networkmanager.DisassociateCustomerGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateCustomerGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) DisassociateCustomerGatewayAsync(ctx workflow.Context, input *networkmanager.DisassociateCustomerGatewayInput) *NetworkmanagerDisassociateCustomerGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateCustomerGateway, input)
    return &NetworkmanagerDisassociateCustomerGatewayResult{Result: future}
}

func (a *NetworkManagerStub) DisassociateLink(ctx workflow.Context, input *networkmanager.DisassociateLinkInput) (*networkmanager.DisassociateLinkOutput, error) {
    var output networkmanager.DisassociateLinkOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateLink, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) DisassociateLinkAsync(ctx workflow.Context, input *networkmanager.DisassociateLinkInput) *NetworkmanagerDisassociateLinkResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateLink, input)
    return &NetworkmanagerDisassociateLinkResult{Result: future}
}

func (a *NetworkManagerStub) GetCustomerGatewayAssociations(ctx workflow.Context, input *networkmanager.GetCustomerGatewayAssociationsInput) (*networkmanager.GetCustomerGatewayAssociationsOutput, error) {
    var output networkmanager.GetCustomerGatewayAssociationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCustomerGatewayAssociations, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) GetCustomerGatewayAssociationsAsync(ctx workflow.Context, input *networkmanager.GetCustomerGatewayAssociationsInput) *NetworkmanagerGetCustomerGatewayAssociationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetCustomerGatewayAssociations, input)
    return &NetworkmanagerGetCustomerGatewayAssociationsResult{Result: future}
}

func (a *NetworkManagerStub) GetDevices(ctx workflow.Context, input *networkmanager.GetDevicesInput) (*networkmanager.GetDevicesOutput, error) {
    var output networkmanager.GetDevicesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDevices, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) GetDevicesAsync(ctx workflow.Context, input *networkmanager.GetDevicesInput) *NetworkmanagerGetDevicesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDevices, input)
    return &NetworkmanagerGetDevicesResult{Result: future}
}

func (a *NetworkManagerStub) GetLinkAssociations(ctx workflow.Context, input *networkmanager.GetLinkAssociationsInput) (*networkmanager.GetLinkAssociationsOutput, error) {
    var output networkmanager.GetLinkAssociationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetLinkAssociations, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) GetLinkAssociationsAsync(ctx workflow.Context, input *networkmanager.GetLinkAssociationsInput) *NetworkmanagerGetLinkAssociationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetLinkAssociations, input)
    return &NetworkmanagerGetLinkAssociationsResult{Result: future}
}

func (a *NetworkManagerStub) GetLinks(ctx workflow.Context, input *networkmanager.GetLinksInput) (*networkmanager.GetLinksOutput, error) {
    var output networkmanager.GetLinksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetLinks, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) GetLinksAsync(ctx workflow.Context, input *networkmanager.GetLinksInput) *NetworkmanagerGetLinksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetLinks, input)
    return &NetworkmanagerGetLinksResult{Result: future}
}

func (a *NetworkManagerStub) GetSites(ctx workflow.Context, input *networkmanager.GetSitesInput) (*networkmanager.GetSitesOutput, error) {
    var output networkmanager.GetSitesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSites, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) GetSitesAsync(ctx workflow.Context, input *networkmanager.GetSitesInput) *NetworkmanagerGetSitesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSites, input)
    return &NetworkmanagerGetSitesResult{Result: future}
}

func (a *NetworkManagerStub) GetTransitGatewayRegistrations(ctx workflow.Context, input *networkmanager.GetTransitGatewayRegistrationsInput) (*networkmanager.GetTransitGatewayRegistrationsOutput, error) {
    var output networkmanager.GetTransitGatewayRegistrationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTransitGatewayRegistrations, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) GetTransitGatewayRegistrationsAsync(ctx workflow.Context, input *networkmanager.GetTransitGatewayRegistrationsInput) *NetworkmanagerGetTransitGatewayRegistrationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTransitGatewayRegistrations, input)
    return &NetworkmanagerGetTransitGatewayRegistrationsResult{Result: future}
}

func (a *NetworkManagerStub) ListTagsForResource(ctx workflow.Context, input *networkmanager.ListTagsForResourceInput) (*networkmanager.ListTagsForResourceOutput, error) {
    var output networkmanager.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) ListTagsForResourceAsync(ctx workflow.Context, input *networkmanager.ListTagsForResourceInput) *NetworkmanagerListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &NetworkmanagerListTagsForResourceResult{Result: future}
}

func (a *NetworkManagerStub) RegisterTransitGateway(ctx workflow.Context, input *networkmanager.RegisterTransitGatewayInput) (*networkmanager.RegisterTransitGatewayOutput, error) {
    var output networkmanager.RegisterTransitGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterTransitGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) RegisterTransitGatewayAsync(ctx workflow.Context, input *networkmanager.RegisterTransitGatewayInput) *NetworkmanagerRegisterTransitGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterTransitGateway, input)
    return &NetworkmanagerRegisterTransitGatewayResult{Result: future}
}

func (a *NetworkManagerStub) TagResource(ctx workflow.Context, input *networkmanager.TagResourceInput) (*networkmanager.TagResourceOutput, error) {
    var output networkmanager.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) TagResourceAsync(ctx workflow.Context, input *networkmanager.TagResourceInput) *NetworkmanagerTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &NetworkmanagerTagResourceResult{Result: future}
}

func (a *NetworkManagerStub) UntagResource(ctx workflow.Context, input *networkmanager.UntagResourceInput) (*networkmanager.UntagResourceOutput, error) {
    var output networkmanager.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) UntagResourceAsync(ctx workflow.Context, input *networkmanager.UntagResourceInput) *NetworkmanagerUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &NetworkmanagerUntagResourceResult{Result: future}
}

func (a *NetworkManagerStub) UpdateDevice(ctx workflow.Context, input *networkmanager.UpdateDeviceInput) (*networkmanager.UpdateDeviceOutput, error) {
    var output networkmanager.UpdateDeviceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDevice, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) UpdateDeviceAsync(ctx workflow.Context, input *networkmanager.UpdateDeviceInput) *NetworkmanagerUpdateDeviceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateDevice, input)
    return &NetworkmanagerUpdateDeviceResult{Result: future}
}

func (a *NetworkManagerStub) UpdateGlobalNetwork(ctx workflow.Context, input *networkmanager.UpdateGlobalNetworkInput) (*networkmanager.UpdateGlobalNetworkOutput, error) {
    var output networkmanager.UpdateGlobalNetworkOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateGlobalNetwork, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) UpdateGlobalNetworkAsync(ctx workflow.Context, input *networkmanager.UpdateGlobalNetworkInput) *NetworkmanagerUpdateGlobalNetworkResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateGlobalNetwork, input)
    return &NetworkmanagerUpdateGlobalNetworkResult{Result: future}
}

func (a *NetworkManagerStub) UpdateLink(ctx workflow.Context, input *networkmanager.UpdateLinkInput) (*networkmanager.UpdateLinkOutput, error) {
    var output networkmanager.UpdateLinkOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateLink, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) UpdateLinkAsync(ctx workflow.Context, input *networkmanager.UpdateLinkInput) *NetworkmanagerUpdateLinkResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateLink, input)
    return &NetworkmanagerUpdateLinkResult{Result: future}
}

func (a *NetworkManagerStub) UpdateSite(ctx workflow.Context, input *networkmanager.UpdateSiteInput) (*networkmanager.UpdateSiteOutput, error) {
    var output networkmanager.UpdateSiteOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSite, input).Get(ctx, &output)
    return &output, err
}

func (a *NetworkManagerStub) UpdateSiteAsync(ctx workflow.Context, input *networkmanager.UpdateSiteInput) *NetworkmanagerUpdateSiteResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateSite, input)
    return &NetworkmanagerUpdateSiteResult{Result: future}
}
