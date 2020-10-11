// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/networkmanager"
	"go.temporal.io/sdk/workflow"
)

type NetworkManagerClient interface {
	AssociateCustomerGateway(ctx workflow.Context, input *networkmanager.AssociateCustomerGatewayInput) (*networkmanager.AssociateCustomerGatewayOutput, error)
	AssociateCustomerGatewayAsync(ctx workflow.Context, input *networkmanager.AssociateCustomerGatewayInput) *NetworkmanagerAssociateCustomerGatewayFuture

	AssociateLink(ctx workflow.Context, input *networkmanager.AssociateLinkInput) (*networkmanager.AssociateLinkOutput, error)
	AssociateLinkAsync(ctx workflow.Context, input *networkmanager.AssociateLinkInput) *NetworkmanagerAssociateLinkFuture

	CreateDevice(ctx workflow.Context, input *networkmanager.CreateDeviceInput) (*networkmanager.CreateDeviceOutput, error)
	CreateDeviceAsync(ctx workflow.Context, input *networkmanager.CreateDeviceInput) *NetworkmanagerCreateDeviceFuture

	CreateGlobalNetwork(ctx workflow.Context, input *networkmanager.CreateGlobalNetworkInput) (*networkmanager.CreateGlobalNetworkOutput, error)
	CreateGlobalNetworkAsync(ctx workflow.Context, input *networkmanager.CreateGlobalNetworkInput) *NetworkmanagerCreateGlobalNetworkFuture

	CreateLink(ctx workflow.Context, input *networkmanager.CreateLinkInput) (*networkmanager.CreateLinkOutput, error)
	CreateLinkAsync(ctx workflow.Context, input *networkmanager.CreateLinkInput) *NetworkmanagerCreateLinkFuture

	CreateSite(ctx workflow.Context, input *networkmanager.CreateSiteInput) (*networkmanager.CreateSiteOutput, error)
	CreateSiteAsync(ctx workflow.Context, input *networkmanager.CreateSiteInput) *NetworkmanagerCreateSiteFuture

	DeleteDevice(ctx workflow.Context, input *networkmanager.DeleteDeviceInput) (*networkmanager.DeleteDeviceOutput, error)
	DeleteDeviceAsync(ctx workflow.Context, input *networkmanager.DeleteDeviceInput) *NetworkmanagerDeleteDeviceFuture

	DeleteGlobalNetwork(ctx workflow.Context, input *networkmanager.DeleteGlobalNetworkInput) (*networkmanager.DeleteGlobalNetworkOutput, error)
	DeleteGlobalNetworkAsync(ctx workflow.Context, input *networkmanager.DeleteGlobalNetworkInput) *NetworkmanagerDeleteGlobalNetworkFuture

	DeleteLink(ctx workflow.Context, input *networkmanager.DeleteLinkInput) (*networkmanager.DeleteLinkOutput, error)
	DeleteLinkAsync(ctx workflow.Context, input *networkmanager.DeleteLinkInput) *NetworkmanagerDeleteLinkFuture

	DeleteSite(ctx workflow.Context, input *networkmanager.DeleteSiteInput) (*networkmanager.DeleteSiteOutput, error)
	DeleteSiteAsync(ctx workflow.Context, input *networkmanager.DeleteSiteInput) *NetworkmanagerDeleteSiteFuture

	DeregisterTransitGateway(ctx workflow.Context, input *networkmanager.DeregisterTransitGatewayInput) (*networkmanager.DeregisterTransitGatewayOutput, error)
	DeregisterTransitGatewayAsync(ctx workflow.Context, input *networkmanager.DeregisterTransitGatewayInput) *NetworkmanagerDeregisterTransitGatewayFuture

	DescribeGlobalNetworks(ctx workflow.Context, input *networkmanager.DescribeGlobalNetworksInput) (*networkmanager.DescribeGlobalNetworksOutput, error)
	DescribeGlobalNetworksAsync(ctx workflow.Context, input *networkmanager.DescribeGlobalNetworksInput) *NetworkmanagerDescribeGlobalNetworksFuture

	DisassociateCustomerGateway(ctx workflow.Context, input *networkmanager.DisassociateCustomerGatewayInput) (*networkmanager.DisassociateCustomerGatewayOutput, error)
	DisassociateCustomerGatewayAsync(ctx workflow.Context, input *networkmanager.DisassociateCustomerGatewayInput) *NetworkmanagerDisassociateCustomerGatewayFuture

	DisassociateLink(ctx workflow.Context, input *networkmanager.DisassociateLinkInput) (*networkmanager.DisassociateLinkOutput, error)
	DisassociateLinkAsync(ctx workflow.Context, input *networkmanager.DisassociateLinkInput) *NetworkmanagerDisassociateLinkFuture

	GetCustomerGatewayAssociations(ctx workflow.Context, input *networkmanager.GetCustomerGatewayAssociationsInput) (*networkmanager.GetCustomerGatewayAssociationsOutput, error)
	GetCustomerGatewayAssociationsAsync(ctx workflow.Context, input *networkmanager.GetCustomerGatewayAssociationsInput) *NetworkmanagerGetCustomerGatewayAssociationsFuture

	GetDevices(ctx workflow.Context, input *networkmanager.GetDevicesInput) (*networkmanager.GetDevicesOutput, error)
	GetDevicesAsync(ctx workflow.Context, input *networkmanager.GetDevicesInput) *NetworkmanagerGetDevicesFuture

	GetLinkAssociations(ctx workflow.Context, input *networkmanager.GetLinkAssociationsInput) (*networkmanager.GetLinkAssociationsOutput, error)
	GetLinkAssociationsAsync(ctx workflow.Context, input *networkmanager.GetLinkAssociationsInput) *NetworkmanagerGetLinkAssociationsFuture

	GetLinks(ctx workflow.Context, input *networkmanager.GetLinksInput) (*networkmanager.GetLinksOutput, error)
	GetLinksAsync(ctx workflow.Context, input *networkmanager.GetLinksInput) *NetworkmanagerGetLinksFuture

	GetSites(ctx workflow.Context, input *networkmanager.GetSitesInput) (*networkmanager.GetSitesOutput, error)
	GetSitesAsync(ctx workflow.Context, input *networkmanager.GetSitesInput) *NetworkmanagerGetSitesFuture

	GetTransitGatewayRegistrations(ctx workflow.Context, input *networkmanager.GetTransitGatewayRegistrationsInput) (*networkmanager.GetTransitGatewayRegistrationsOutput, error)
	GetTransitGatewayRegistrationsAsync(ctx workflow.Context, input *networkmanager.GetTransitGatewayRegistrationsInput) *NetworkmanagerGetTransitGatewayRegistrationsFuture

	ListTagsForResource(ctx workflow.Context, input *networkmanager.ListTagsForResourceInput) (*networkmanager.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *networkmanager.ListTagsForResourceInput) *NetworkmanagerListTagsForResourceFuture

	RegisterTransitGateway(ctx workflow.Context, input *networkmanager.RegisterTransitGatewayInput) (*networkmanager.RegisterTransitGatewayOutput, error)
	RegisterTransitGatewayAsync(ctx workflow.Context, input *networkmanager.RegisterTransitGatewayInput) *NetworkmanagerRegisterTransitGatewayFuture

	TagResource(ctx workflow.Context, input *networkmanager.TagResourceInput) (*networkmanager.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *networkmanager.TagResourceInput) *NetworkmanagerTagResourceFuture

	UntagResource(ctx workflow.Context, input *networkmanager.UntagResourceInput) (*networkmanager.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *networkmanager.UntagResourceInput) *NetworkmanagerUntagResourceFuture

	UpdateDevice(ctx workflow.Context, input *networkmanager.UpdateDeviceInput) (*networkmanager.UpdateDeviceOutput, error)
	UpdateDeviceAsync(ctx workflow.Context, input *networkmanager.UpdateDeviceInput) *NetworkmanagerUpdateDeviceFuture

	UpdateGlobalNetwork(ctx workflow.Context, input *networkmanager.UpdateGlobalNetworkInput) (*networkmanager.UpdateGlobalNetworkOutput, error)
	UpdateGlobalNetworkAsync(ctx workflow.Context, input *networkmanager.UpdateGlobalNetworkInput) *NetworkmanagerUpdateGlobalNetworkFuture

	UpdateLink(ctx workflow.Context, input *networkmanager.UpdateLinkInput) (*networkmanager.UpdateLinkOutput, error)
	UpdateLinkAsync(ctx workflow.Context, input *networkmanager.UpdateLinkInput) *NetworkmanagerUpdateLinkFuture

	UpdateSite(ctx workflow.Context, input *networkmanager.UpdateSiteInput) (*networkmanager.UpdateSiteOutput, error)
	UpdateSiteAsync(ctx workflow.Context, input *networkmanager.UpdateSiteInput) *NetworkmanagerUpdateSiteFuture
}

type NetworkManagerStub struct{}

func NewNetworkManagerStub() NetworkManagerClient {
	return &NetworkManagerStub{}
}

type NetworkmanagerAssociateCustomerGatewayFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerAssociateCustomerGatewayFuture) Get(ctx workflow.Context) (*networkmanager.AssociateCustomerGatewayOutput, error) {
	var output networkmanager.AssociateCustomerGatewayOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerAssociateLinkFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerAssociateLinkFuture) Get(ctx workflow.Context) (*networkmanager.AssociateLinkOutput, error) {
	var output networkmanager.AssociateLinkOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerCreateDeviceFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerCreateDeviceFuture) Get(ctx workflow.Context) (*networkmanager.CreateDeviceOutput, error) {
	var output networkmanager.CreateDeviceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerCreateGlobalNetworkFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerCreateGlobalNetworkFuture) Get(ctx workflow.Context) (*networkmanager.CreateGlobalNetworkOutput, error) {
	var output networkmanager.CreateGlobalNetworkOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerCreateLinkFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerCreateLinkFuture) Get(ctx workflow.Context) (*networkmanager.CreateLinkOutput, error) {
	var output networkmanager.CreateLinkOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerCreateSiteFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerCreateSiteFuture) Get(ctx workflow.Context) (*networkmanager.CreateSiteOutput, error) {
	var output networkmanager.CreateSiteOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerDeleteDeviceFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerDeleteDeviceFuture) Get(ctx workflow.Context) (*networkmanager.DeleteDeviceOutput, error) {
	var output networkmanager.DeleteDeviceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerDeleteGlobalNetworkFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerDeleteGlobalNetworkFuture) Get(ctx workflow.Context) (*networkmanager.DeleteGlobalNetworkOutput, error) {
	var output networkmanager.DeleteGlobalNetworkOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerDeleteLinkFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerDeleteLinkFuture) Get(ctx workflow.Context) (*networkmanager.DeleteLinkOutput, error) {
	var output networkmanager.DeleteLinkOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerDeleteSiteFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerDeleteSiteFuture) Get(ctx workflow.Context) (*networkmanager.DeleteSiteOutput, error) {
	var output networkmanager.DeleteSiteOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerDeregisterTransitGatewayFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerDeregisterTransitGatewayFuture) Get(ctx workflow.Context) (*networkmanager.DeregisterTransitGatewayOutput, error) {
	var output networkmanager.DeregisterTransitGatewayOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerDescribeGlobalNetworksFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerDescribeGlobalNetworksFuture) Get(ctx workflow.Context) (*networkmanager.DescribeGlobalNetworksOutput, error) {
	var output networkmanager.DescribeGlobalNetworksOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerDisassociateCustomerGatewayFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerDisassociateCustomerGatewayFuture) Get(ctx workflow.Context) (*networkmanager.DisassociateCustomerGatewayOutput, error) {
	var output networkmanager.DisassociateCustomerGatewayOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerDisassociateLinkFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerDisassociateLinkFuture) Get(ctx workflow.Context) (*networkmanager.DisassociateLinkOutput, error) {
	var output networkmanager.DisassociateLinkOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerGetCustomerGatewayAssociationsFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerGetCustomerGatewayAssociationsFuture) Get(ctx workflow.Context) (*networkmanager.GetCustomerGatewayAssociationsOutput, error) {
	var output networkmanager.GetCustomerGatewayAssociationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerGetDevicesFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerGetDevicesFuture) Get(ctx workflow.Context) (*networkmanager.GetDevicesOutput, error) {
	var output networkmanager.GetDevicesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerGetLinkAssociationsFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerGetLinkAssociationsFuture) Get(ctx workflow.Context) (*networkmanager.GetLinkAssociationsOutput, error) {
	var output networkmanager.GetLinkAssociationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerGetLinksFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerGetLinksFuture) Get(ctx workflow.Context) (*networkmanager.GetLinksOutput, error) {
	var output networkmanager.GetLinksOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerGetSitesFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerGetSitesFuture) Get(ctx workflow.Context) (*networkmanager.GetSitesOutput, error) {
	var output networkmanager.GetSitesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerGetTransitGatewayRegistrationsFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerGetTransitGatewayRegistrationsFuture) Get(ctx workflow.Context) (*networkmanager.GetTransitGatewayRegistrationsOutput, error) {
	var output networkmanager.GetTransitGatewayRegistrationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerListTagsForResourceFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerListTagsForResourceFuture) Get(ctx workflow.Context) (*networkmanager.ListTagsForResourceOutput, error) {
	var output networkmanager.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerRegisterTransitGatewayFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerRegisterTransitGatewayFuture) Get(ctx workflow.Context) (*networkmanager.RegisterTransitGatewayOutput, error) {
	var output networkmanager.RegisterTransitGatewayOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerTagResourceFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerTagResourceFuture) Get(ctx workflow.Context) (*networkmanager.TagResourceOutput, error) {
	var output networkmanager.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerUntagResourceFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerUntagResourceFuture) Get(ctx workflow.Context) (*networkmanager.UntagResourceOutput, error) {
	var output networkmanager.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerUpdateDeviceFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerUpdateDeviceFuture) Get(ctx workflow.Context) (*networkmanager.UpdateDeviceOutput, error) {
	var output networkmanager.UpdateDeviceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerUpdateGlobalNetworkFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerUpdateGlobalNetworkFuture) Get(ctx workflow.Context) (*networkmanager.UpdateGlobalNetworkOutput, error) {
	var output networkmanager.UpdateGlobalNetworkOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerUpdateLinkFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerUpdateLinkFuture) Get(ctx workflow.Context) (*networkmanager.UpdateLinkOutput, error) {
	var output networkmanager.UpdateLinkOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type NetworkmanagerUpdateSiteFuture struct {
	Future workflow.Future
}

func (r *NetworkmanagerUpdateSiteFuture) Get(ctx workflow.Context) (*networkmanager.UpdateSiteOutput, error) {
	var output networkmanager.UpdateSiteOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) AssociateCustomerGateway(ctx workflow.Context, input *networkmanager.AssociateCustomerGatewayInput) (*networkmanager.AssociateCustomerGatewayOutput, error) {
	var output networkmanager.AssociateCustomerGatewayOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.AssociateCustomerGateway", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) AssociateCustomerGatewayAsync(ctx workflow.Context, input *networkmanager.AssociateCustomerGatewayInput) *NetworkmanagerAssociateCustomerGatewayFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.AssociateCustomerGateway", input)
	return &NetworkmanagerAssociateCustomerGatewayFuture{Future: future}
}

func (a *NetworkManagerStub) AssociateLink(ctx workflow.Context, input *networkmanager.AssociateLinkInput) (*networkmanager.AssociateLinkOutput, error) {
	var output networkmanager.AssociateLinkOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.AssociateLink", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) AssociateLinkAsync(ctx workflow.Context, input *networkmanager.AssociateLinkInput) *NetworkmanagerAssociateLinkFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.AssociateLink", input)
	return &NetworkmanagerAssociateLinkFuture{Future: future}
}

func (a *NetworkManagerStub) CreateDevice(ctx workflow.Context, input *networkmanager.CreateDeviceInput) (*networkmanager.CreateDeviceOutput, error) {
	var output networkmanager.CreateDeviceOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.CreateDevice", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) CreateDeviceAsync(ctx workflow.Context, input *networkmanager.CreateDeviceInput) *NetworkmanagerCreateDeviceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.CreateDevice", input)
	return &NetworkmanagerCreateDeviceFuture{Future: future}
}

func (a *NetworkManagerStub) CreateGlobalNetwork(ctx workflow.Context, input *networkmanager.CreateGlobalNetworkInput) (*networkmanager.CreateGlobalNetworkOutput, error) {
	var output networkmanager.CreateGlobalNetworkOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.CreateGlobalNetwork", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) CreateGlobalNetworkAsync(ctx workflow.Context, input *networkmanager.CreateGlobalNetworkInput) *NetworkmanagerCreateGlobalNetworkFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.CreateGlobalNetwork", input)
	return &NetworkmanagerCreateGlobalNetworkFuture{Future: future}
}

func (a *NetworkManagerStub) CreateLink(ctx workflow.Context, input *networkmanager.CreateLinkInput) (*networkmanager.CreateLinkOutput, error) {
	var output networkmanager.CreateLinkOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.CreateLink", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) CreateLinkAsync(ctx workflow.Context, input *networkmanager.CreateLinkInput) *NetworkmanagerCreateLinkFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.CreateLink", input)
	return &NetworkmanagerCreateLinkFuture{Future: future}
}

func (a *NetworkManagerStub) CreateSite(ctx workflow.Context, input *networkmanager.CreateSiteInput) (*networkmanager.CreateSiteOutput, error) {
	var output networkmanager.CreateSiteOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.CreateSite", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) CreateSiteAsync(ctx workflow.Context, input *networkmanager.CreateSiteInput) *NetworkmanagerCreateSiteFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.CreateSite", input)
	return &NetworkmanagerCreateSiteFuture{Future: future}
}

func (a *NetworkManagerStub) DeleteDevice(ctx workflow.Context, input *networkmanager.DeleteDeviceInput) (*networkmanager.DeleteDeviceOutput, error) {
	var output networkmanager.DeleteDeviceOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.DeleteDevice", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) DeleteDeviceAsync(ctx workflow.Context, input *networkmanager.DeleteDeviceInput) *NetworkmanagerDeleteDeviceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.DeleteDevice", input)
	return &NetworkmanagerDeleteDeviceFuture{Future: future}
}

func (a *NetworkManagerStub) DeleteGlobalNetwork(ctx workflow.Context, input *networkmanager.DeleteGlobalNetworkInput) (*networkmanager.DeleteGlobalNetworkOutput, error) {
	var output networkmanager.DeleteGlobalNetworkOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.DeleteGlobalNetwork", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) DeleteGlobalNetworkAsync(ctx workflow.Context, input *networkmanager.DeleteGlobalNetworkInput) *NetworkmanagerDeleteGlobalNetworkFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.DeleteGlobalNetwork", input)
	return &NetworkmanagerDeleteGlobalNetworkFuture{Future: future}
}

func (a *NetworkManagerStub) DeleteLink(ctx workflow.Context, input *networkmanager.DeleteLinkInput) (*networkmanager.DeleteLinkOutput, error) {
	var output networkmanager.DeleteLinkOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.DeleteLink", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) DeleteLinkAsync(ctx workflow.Context, input *networkmanager.DeleteLinkInput) *NetworkmanagerDeleteLinkFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.DeleteLink", input)
	return &NetworkmanagerDeleteLinkFuture{Future: future}
}

func (a *NetworkManagerStub) DeleteSite(ctx workflow.Context, input *networkmanager.DeleteSiteInput) (*networkmanager.DeleteSiteOutput, error) {
	var output networkmanager.DeleteSiteOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.DeleteSite", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) DeleteSiteAsync(ctx workflow.Context, input *networkmanager.DeleteSiteInput) *NetworkmanagerDeleteSiteFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.DeleteSite", input)
	return &NetworkmanagerDeleteSiteFuture{Future: future}
}

func (a *NetworkManagerStub) DeregisterTransitGateway(ctx workflow.Context, input *networkmanager.DeregisterTransitGatewayInput) (*networkmanager.DeregisterTransitGatewayOutput, error) {
	var output networkmanager.DeregisterTransitGatewayOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.DeregisterTransitGateway", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) DeregisterTransitGatewayAsync(ctx workflow.Context, input *networkmanager.DeregisterTransitGatewayInput) *NetworkmanagerDeregisterTransitGatewayFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.DeregisterTransitGateway", input)
	return &NetworkmanagerDeregisterTransitGatewayFuture{Future: future}
}

func (a *NetworkManagerStub) DescribeGlobalNetworks(ctx workflow.Context, input *networkmanager.DescribeGlobalNetworksInput) (*networkmanager.DescribeGlobalNetworksOutput, error) {
	var output networkmanager.DescribeGlobalNetworksOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.DescribeGlobalNetworks", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) DescribeGlobalNetworksAsync(ctx workflow.Context, input *networkmanager.DescribeGlobalNetworksInput) *NetworkmanagerDescribeGlobalNetworksFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.DescribeGlobalNetworks", input)
	return &NetworkmanagerDescribeGlobalNetworksFuture{Future: future}
}

func (a *NetworkManagerStub) DisassociateCustomerGateway(ctx workflow.Context, input *networkmanager.DisassociateCustomerGatewayInput) (*networkmanager.DisassociateCustomerGatewayOutput, error) {
	var output networkmanager.DisassociateCustomerGatewayOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.DisassociateCustomerGateway", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) DisassociateCustomerGatewayAsync(ctx workflow.Context, input *networkmanager.DisassociateCustomerGatewayInput) *NetworkmanagerDisassociateCustomerGatewayFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.DisassociateCustomerGateway", input)
	return &NetworkmanagerDisassociateCustomerGatewayFuture{Future: future}
}

func (a *NetworkManagerStub) DisassociateLink(ctx workflow.Context, input *networkmanager.DisassociateLinkInput) (*networkmanager.DisassociateLinkOutput, error) {
	var output networkmanager.DisassociateLinkOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.DisassociateLink", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) DisassociateLinkAsync(ctx workflow.Context, input *networkmanager.DisassociateLinkInput) *NetworkmanagerDisassociateLinkFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.DisassociateLink", input)
	return &NetworkmanagerDisassociateLinkFuture{Future: future}
}

func (a *NetworkManagerStub) GetCustomerGatewayAssociations(ctx workflow.Context, input *networkmanager.GetCustomerGatewayAssociationsInput) (*networkmanager.GetCustomerGatewayAssociationsOutput, error) {
	var output networkmanager.GetCustomerGatewayAssociationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.GetCustomerGatewayAssociations", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) GetCustomerGatewayAssociationsAsync(ctx workflow.Context, input *networkmanager.GetCustomerGatewayAssociationsInput) *NetworkmanagerGetCustomerGatewayAssociationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.GetCustomerGatewayAssociations", input)
	return &NetworkmanagerGetCustomerGatewayAssociationsFuture{Future: future}
}

func (a *NetworkManagerStub) GetDevices(ctx workflow.Context, input *networkmanager.GetDevicesInput) (*networkmanager.GetDevicesOutput, error) {
	var output networkmanager.GetDevicesOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.GetDevices", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) GetDevicesAsync(ctx workflow.Context, input *networkmanager.GetDevicesInput) *NetworkmanagerGetDevicesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.GetDevices", input)
	return &NetworkmanagerGetDevicesFuture{Future: future}
}

func (a *NetworkManagerStub) GetLinkAssociations(ctx workflow.Context, input *networkmanager.GetLinkAssociationsInput) (*networkmanager.GetLinkAssociationsOutput, error) {
	var output networkmanager.GetLinkAssociationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.GetLinkAssociations", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) GetLinkAssociationsAsync(ctx workflow.Context, input *networkmanager.GetLinkAssociationsInput) *NetworkmanagerGetLinkAssociationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.GetLinkAssociations", input)
	return &NetworkmanagerGetLinkAssociationsFuture{Future: future}
}

func (a *NetworkManagerStub) GetLinks(ctx workflow.Context, input *networkmanager.GetLinksInput) (*networkmanager.GetLinksOutput, error) {
	var output networkmanager.GetLinksOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.GetLinks", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) GetLinksAsync(ctx workflow.Context, input *networkmanager.GetLinksInput) *NetworkmanagerGetLinksFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.GetLinks", input)
	return &NetworkmanagerGetLinksFuture{Future: future}
}

func (a *NetworkManagerStub) GetSites(ctx workflow.Context, input *networkmanager.GetSitesInput) (*networkmanager.GetSitesOutput, error) {
	var output networkmanager.GetSitesOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.GetSites", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) GetSitesAsync(ctx workflow.Context, input *networkmanager.GetSitesInput) *NetworkmanagerGetSitesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.GetSites", input)
	return &NetworkmanagerGetSitesFuture{Future: future}
}

func (a *NetworkManagerStub) GetTransitGatewayRegistrations(ctx workflow.Context, input *networkmanager.GetTransitGatewayRegistrationsInput) (*networkmanager.GetTransitGatewayRegistrationsOutput, error) {
	var output networkmanager.GetTransitGatewayRegistrationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.GetTransitGatewayRegistrations", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) GetTransitGatewayRegistrationsAsync(ctx workflow.Context, input *networkmanager.GetTransitGatewayRegistrationsInput) *NetworkmanagerGetTransitGatewayRegistrationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.GetTransitGatewayRegistrations", input)
	return &NetworkmanagerGetTransitGatewayRegistrationsFuture{Future: future}
}

func (a *NetworkManagerStub) ListTagsForResource(ctx workflow.Context, input *networkmanager.ListTagsForResourceInput) (*networkmanager.ListTagsForResourceOutput, error) {
	var output networkmanager.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) ListTagsForResourceAsync(ctx workflow.Context, input *networkmanager.ListTagsForResourceInput) *NetworkmanagerListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.ListTagsForResource", input)
	return &NetworkmanagerListTagsForResourceFuture{Future: future}
}

func (a *NetworkManagerStub) RegisterTransitGateway(ctx workflow.Context, input *networkmanager.RegisterTransitGatewayInput) (*networkmanager.RegisterTransitGatewayOutput, error) {
	var output networkmanager.RegisterTransitGatewayOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.RegisterTransitGateway", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) RegisterTransitGatewayAsync(ctx workflow.Context, input *networkmanager.RegisterTransitGatewayInput) *NetworkmanagerRegisterTransitGatewayFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.RegisterTransitGateway", input)
	return &NetworkmanagerRegisterTransitGatewayFuture{Future: future}
}

func (a *NetworkManagerStub) TagResource(ctx workflow.Context, input *networkmanager.TagResourceInput) (*networkmanager.TagResourceOutput, error) {
	var output networkmanager.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) TagResourceAsync(ctx workflow.Context, input *networkmanager.TagResourceInput) *NetworkmanagerTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.TagResource", input)
	return &NetworkmanagerTagResourceFuture{Future: future}
}

func (a *NetworkManagerStub) UntagResource(ctx workflow.Context, input *networkmanager.UntagResourceInput) (*networkmanager.UntagResourceOutput, error) {
	var output networkmanager.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) UntagResourceAsync(ctx workflow.Context, input *networkmanager.UntagResourceInput) *NetworkmanagerUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.UntagResource", input)
	return &NetworkmanagerUntagResourceFuture{Future: future}
}

func (a *NetworkManagerStub) UpdateDevice(ctx workflow.Context, input *networkmanager.UpdateDeviceInput) (*networkmanager.UpdateDeviceOutput, error) {
	var output networkmanager.UpdateDeviceOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.UpdateDevice", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) UpdateDeviceAsync(ctx workflow.Context, input *networkmanager.UpdateDeviceInput) *NetworkmanagerUpdateDeviceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.UpdateDevice", input)
	return &NetworkmanagerUpdateDeviceFuture{Future: future}
}

func (a *NetworkManagerStub) UpdateGlobalNetwork(ctx workflow.Context, input *networkmanager.UpdateGlobalNetworkInput) (*networkmanager.UpdateGlobalNetworkOutput, error) {
	var output networkmanager.UpdateGlobalNetworkOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.UpdateGlobalNetwork", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) UpdateGlobalNetworkAsync(ctx workflow.Context, input *networkmanager.UpdateGlobalNetworkInput) *NetworkmanagerUpdateGlobalNetworkFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.UpdateGlobalNetwork", input)
	return &NetworkmanagerUpdateGlobalNetworkFuture{Future: future}
}

func (a *NetworkManagerStub) UpdateLink(ctx workflow.Context, input *networkmanager.UpdateLinkInput) (*networkmanager.UpdateLinkOutput, error) {
	var output networkmanager.UpdateLinkOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.UpdateLink", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) UpdateLinkAsync(ctx workflow.Context, input *networkmanager.UpdateLinkInput) *NetworkmanagerUpdateLinkFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.UpdateLink", input)
	return &NetworkmanagerUpdateLinkFuture{Future: future}
}

func (a *NetworkManagerStub) UpdateSite(ctx workflow.Context, input *networkmanager.UpdateSiteInput) (*networkmanager.UpdateSiteOutput, error) {
	var output networkmanager.UpdateSiteOutput
	err := workflow.ExecuteActivity(ctx, "aws.networkmanager.UpdateSite", input).Get(ctx, &output)
	return &output, err
}

func (a *NetworkManagerStub) UpdateSiteAsync(ctx workflow.Context, input *networkmanager.UpdateSiteInput) *NetworkmanagerUpdateSiteFuture {
	future := workflow.ExecuteActivity(ctx, "aws.networkmanager.UpdateSite", input)
	return &NetworkmanagerUpdateSiteFuture{Future: future}
}
