// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/ram"
	"go.temporal.io/sdk/workflow"
)

type RAMClient interface {
	AcceptResourceShareInvitation(ctx workflow.Context, input *ram.AcceptResourceShareInvitationInput) (*ram.AcceptResourceShareInvitationOutput, error)
	AcceptResourceShareInvitationAsync(ctx workflow.Context, input *ram.AcceptResourceShareInvitationInput) *RamAcceptResourceShareInvitationFuture

	AssociateResourceShare(ctx workflow.Context, input *ram.AssociateResourceShareInput) (*ram.AssociateResourceShareOutput, error)
	AssociateResourceShareAsync(ctx workflow.Context, input *ram.AssociateResourceShareInput) *RamAssociateResourceShareFuture

	AssociateResourceSharePermission(ctx workflow.Context, input *ram.AssociateResourceSharePermissionInput) (*ram.AssociateResourceSharePermissionOutput, error)
	AssociateResourceSharePermissionAsync(ctx workflow.Context, input *ram.AssociateResourceSharePermissionInput) *RamAssociateResourceSharePermissionFuture

	CreateResourceShare(ctx workflow.Context, input *ram.CreateResourceShareInput) (*ram.CreateResourceShareOutput, error)
	CreateResourceShareAsync(ctx workflow.Context, input *ram.CreateResourceShareInput) *RamCreateResourceShareFuture

	DeleteResourceShare(ctx workflow.Context, input *ram.DeleteResourceShareInput) (*ram.DeleteResourceShareOutput, error)
	DeleteResourceShareAsync(ctx workflow.Context, input *ram.DeleteResourceShareInput) *RamDeleteResourceShareFuture

	DisassociateResourceShare(ctx workflow.Context, input *ram.DisassociateResourceShareInput) (*ram.DisassociateResourceShareOutput, error)
	DisassociateResourceShareAsync(ctx workflow.Context, input *ram.DisassociateResourceShareInput) *RamDisassociateResourceShareFuture

	DisassociateResourceSharePermission(ctx workflow.Context, input *ram.DisassociateResourceSharePermissionInput) (*ram.DisassociateResourceSharePermissionOutput, error)
	DisassociateResourceSharePermissionAsync(ctx workflow.Context, input *ram.DisassociateResourceSharePermissionInput) *RamDisassociateResourceSharePermissionFuture

	EnableSharingWithAwsOrganization(ctx workflow.Context, input *ram.EnableSharingWithAwsOrganizationInput) (*ram.EnableSharingWithAwsOrganizationOutput, error)
	EnableSharingWithAwsOrganizationAsync(ctx workflow.Context, input *ram.EnableSharingWithAwsOrganizationInput) *RamEnableSharingWithAwsOrganizationFuture

	GetPermission(ctx workflow.Context, input *ram.GetPermissionInput) (*ram.GetPermissionOutput, error)
	GetPermissionAsync(ctx workflow.Context, input *ram.GetPermissionInput) *RamGetPermissionFuture

	GetResourcePolicies(ctx workflow.Context, input *ram.GetResourcePoliciesInput) (*ram.GetResourcePoliciesOutput, error)
	GetResourcePoliciesAsync(ctx workflow.Context, input *ram.GetResourcePoliciesInput) *RamGetResourcePoliciesFuture

	GetResourceShareAssociations(ctx workflow.Context, input *ram.GetResourceShareAssociationsInput) (*ram.GetResourceShareAssociationsOutput, error)
	GetResourceShareAssociationsAsync(ctx workflow.Context, input *ram.GetResourceShareAssociationsInput) *RamGetResourceShareAssociationsFuture

	GetResourceShareInvitations(ctx workflow.Context, input *ram.GetResourceShareInvitationsInput) (*ram.GetResourceShareInvitationsOutput, error)
	GetResourceShareInvitationsAsync(ctx workflow.Context, input *ram.GetResourceShareInvitationsInput) *RamGetResourceShareInvitationsFuture

	GetResourceShares(ctx workflow.Context, input *ram.GetResourceSharesInput) (*ram.GetResourceSharesOutput, error)
	GetResourceSharesAsync(ctx workflow.Context, input *ram.GetResourceSharesInput) *RamGetResourceSharesFuture

	ListPendingInvitationResources(ctx workflow.Context, input *ram.ListPendingInvitationResourcesInput) (*ram.ListPendingInvitationResourcesOutput, error)
	ListPendingInvitationResourcesAsync(ctx workflow.Context, input *ram.ListPendingInvitationResourcesInput) *RamListPendingInvitationResourcesFuture

	ListPermissions(ctx workflow.Context, input *ram.ListPermissionsInput) (*ram.ListPermissionsOutput, error)
	ListPermissionsAsync(ctx workflow.Context, input *ram.ListPermissionsInput) *RamListPermissionsFuture

	ListPrincipals(ctx workflow.Context, input *ram.ListPrincipalsInput) (*ram.ListPrincipalsOutput, error)
	ListPrincipalsAsync(ctx workflow.Context, input *ram.ListPrincipalsInput) *RamListPrincipalsFuture

	ListResourceSharePermissions(ctx workflow.Context, input *ram.ListResourceSharePermissionsInput) (*ram.ListResourceSharePermissionsOutput, error)
	ListResourceSharePermissionsAsync(ctx workflow.Context, input *ram.ListResourceSharePermissionsInput) *RamListResourceSharePermissionsFuture

	ListResourceTypes(ctx workflow.Context, input *ram.ListResourceTypesInput) (*ram.ListResourceTypesOutput, error)
	ListResourceTypesAsync(ctx workflow.Context, input *ram.ListResourceTypesInput) *RamListResourceTypesFuture

	ListResources(ctx workflow.Context, input *ram.ListResourcesInput) (*ram.ListResourcesOutput, error)
	ListResourcesAsync(ctx workflow.Context, input *ram.ListResourcesInput) *RamListResourcesFuture

	PromoteResourceShareCreatedFromPolicy(ctx workflow.Context, input *ram.PromoteResourceShareCreatedFromPolicyInput) (*ram.PromoteResourceShareCreatedFromPolicyOutput, error)
	PromoteResourceShareCreatedFromPolicyAsync(ctx workflow.Context, input *ram.PromoteResourceShareCreatedFromPolicyInput) *RamPromoteResourceShareCreatedFromPolicyFuture

	RejectResourceShareInvitation(ctx workflow.Context, input *ram.RejectResourceShareInvitationInput) (*ram.RejectResourceShareInvitationOutput, error)
	RejectResourceShareInvitationAsync(ctx workflow.Context, input *ram.RejectResourceShareInvitationInput) *RamRejectResourceShareInvitationFuture

	TagResource(ctx workflow.Context, input *ram.TagResourceInput) (*ram.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *ram.TagResourceInput) *RamTagResourceFuture

	UntagResource(ctx workflow.Context, input *ram.UntagResourceInput) (*ram.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *ram.UntagResourceInput) *RamUntagResourceFuture

	UpdateResourceShare(ctx workflow.Context, input *ram.UpdateResourceShareInput) (*ram.UpdateResourceShareOutput, error)
	UpdateResourceShareAsync(ctx workflow.Context, input *ram.UpdateResourceShareInput) *RamUpdateResourceShareFuture
}

type RAMStub struct{}

func NewRAMStub() RAMClient {
	return &RAMStub{}
}

type RamAcceptResourceShareInvitationFuture struct {
	Future workflow.Future
}

func (r *RamAcceptResourceShareInvitationFuture) Get(ctx workflow.Context) (*ram.AcceptResourceShareInvitationOutput, error) {
	var output ram.AcceptResourceShareInvitationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RamAssociateResourceShareFuture struct {
	Future workflow.Future
}

func (r *RamAssociateResourceShareFuture) Get(ctx workflow.Context) (*ram.AssociateResourceShareOutput, error) {
	var output ram.AssociateResourceShareOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RamAssociateResourceSharePermissionFuture struct {
	Future workflow.Future
}

func (r *RamAssociateResourceSharePermissionFuture) Get(ctx workflow.Context) (*ram.AssociateResourceSharePermissionOutput, error) {
	var output ram.AssociateResourceSharePermissionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RamCreateResourceShareFuture struct {
	Future workflow.Future
}

func (r *RamCreateResourceShareFuture) Get(ctx workflow.Context) (*ram.CreateResourceShareOutput, error) {
	var output ram.CreateResourceShareOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RamDeleteResourceShareFuture struct {
	Future workflow.Future
}

func (r *RamDeleteResourceShareFuture) Get(ctx workflow.Context) (*ram.DeleteResourceShareOutput, error) {
	var output ram.DeleteResourceShareOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RamDisassociateResourceShareFuture struct {
	Future workflow.Future
}

func (r *RamDisassociateResourceShareFuture) Get(ctx workflow.Context) (*ram.DisassociateResourceShareOutput, error) {
	var output ram.DisassociateResourceShareOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RamDisassociateResourceSharePermissionFuture struct {
	Future workflow.Future
}

func (r *RamDisassociateResourceSharePermissionFuture) Get(ctx workflow.Context) (*ram.DisassociateResourceSharePermissionOutput, error) {
	var output ram.DisassociateResourceSharePermissionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RamEnableSharingWithAwsOrganizationFuture struct {
	Future workflow.Future
}

func (r *RamEnableSharingWithAwsOrganizationFuture) Get(ctx workflow.Context) (*ram.EnableSharingWithAwsOrganizationOutput, error) {
	var output ram.EnableSharingWithAwsOrganizationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RamGetPermissionFuture struct {
	Future workflow.Future
}

func (r *RamGetPermissionFuture) Get(ctx workflow.Context) (*ram.GetPermissionOutput, error) {
	var output ram.GetPermissionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RamGetResourcePoliciesFuture struct {
	Future workflow.Future
}

func (r *RamGetResourcePoliciesFuture) Get(ctx workflow.Context) (*ram.GetResourcePoliciesOutput, error) {
	var output ram.GetResourcePoliciesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RamGetResourceShareAssociationsFuture struct {
	Future workflow.Future
}

func (r *RamGetResourceShareAssociationsFuture) Get(ctx workflow.Context) (*ram.GetResourceShareAssociationsOutput, error) {
	var output ram.GetResourceShareAssociationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RamGetResourceShareInvitationsFuture struct {
	Future workflow.Future
}

func (r *RamGetResourceShareInvitationsFuture) Get(ctx workflow.Context) (*ram.GetResourceShareInvitationsOutput, error) {
	var output ram.GetResourceShareInvitationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RamGetResourceSharesFuture struct {
	Future workflow.Future
}

func (r *RamGetResourceSharesFuture) Get(ctx workflow.Context) (*ram.GetResourceSharesOutput, error) {
	var output ram.GetResourceSharesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RamListPendingInvitationResourcesFuture struct {
	Future workflow.Future
}

func (r *RamListPendingInvitationResourcesFuture) Get(ctx workflow.Context) (*ram.ListPendingInvitationResourcesOutput, error) {
	var output ram.ListPendingInvitationResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RamListPermissionsFuture struct {
	Future workflow.Future
}

func (r *RamListPermissionsFuture) Get(ctx workflow.Context) (*ram.ListPermissionsOutput, error) {
	var output ram.ListPermissionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RamListPrincipalsFuture struct {
	Future workflow.Future
}

func (r *RamListPrincipalsFuture) Get(ctx workflow.Context) (*ram.ListPrincipalsOutput, error) {
	var output ram.ListPrincipalsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RamListResourceSharePermissionsFuture struct {
	Future workflow.Future
}

func (r *RamListResourceSharePermissionsFuture) Get(ctx workflow.Context) (*ram.ListResourceSharePermissionsOutput, error) {
	var output ram.ListResourceSharePermissionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RamListResourceTypesFuture struct {
	Future workflow.Future
}

func (r *RamListResourceTypesFuture) Get(ctx workflow.Context) (*ram.ListResourceTypesOutput, error) {
	var output ram.ListResourceTypesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RamListResourcesFuture struct {
	Future workflow.Future
}

func (r *RamListResourcesFuture) Get(ctx workflow.Context) (*ram.ListResourcesOutput, error) {
	var output ram.ListResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RamPromoteResourceShareCreatedFromPolicyFuture struct {
	Future workflow.Future
}

func (r *RamPromoteResourceShareCreatedFromPolicyFuture) Get(ctx workflow.Context) (*ram.PromoteResourceShareCreatedFromPolicyOutput, error) {
	var output ram.PromoteResourceShareCreatedFromPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RamRejectResourceShareInvitationFuture struct {
	Future workflow.Future
}

func (r *RamRejectResourceShareInvitationFuture) Get(ctx workflow.Context) (*ram.RejectResourceShareInvitationOutput, error) {
	var output ram.RejectResourceShareInvitationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RamTagResourceFuture struct {
	Future workflow.Future
}

func (r *RamTagResourceFuture) Get(ctx workflow.Context) (*ram.TagResourceOutput, error) {
	var output ram.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RamUntagResourceFuture struct {
	Future workflow.Future
}

func (r *RamUntagResourceFuture) Get(ctx workflow.Context) (*ram.UntagResourceOutput, error) {
	var output ram.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RamUpdateResourceShareFuture struct {
	Future workflow.Future
}

func (r *RamUpdateResourceShareFuture) Get(ctx workflow.Context) (*ram.UpdateResourceShareOutput, error) {
	var output ram.UpdateResourceShareOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) AcceptResourceShareInvitation(ctx workflow.Context, input *ram.AcceptResourceShareInvitationInput) (*ram.AcceptResourceShareInvitationOutput, error) {
	var output ram.AcceptResourceShareInvitationOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.AcceptResourceShareInvitation", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) AcceptResourceShareInvitationAsync(ctx workflow.Context, input *ram.AcceptResourceShareInvitationInput) *RamAcceptResourceShareInvitationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.AcceptResourceShareInvitation", input)
	return &RamAcceptResourceShareInvitationFuture{Future: future}
}

func (a *RAMStub) AssociateResourceShare(ctx workflow.Context, input *ram.AssociateResourceShareInput) (*ram.AssociateResourceShareOutput, error) {
	var output ram.AssociateResourceShareOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.AssociateResourceShare", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) AssociateResourceShareAsync(ctx workflow.Context, input *ram.AssociateResourceShareInput) *RamAssociateResourceShareFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.AssociateResourceShare", input)
	return &RamAssociateResourceShareFuture{Future: future}
}

func (a *RAMStub) AssociateResourceSharePermission(ctx workflow.Context, input *ram.AssociateResourceSharePermissionInput) (*ram.AssociateResourceSharePermissionOutput, error) {
	var output ram.AssociateResourceSharePermissionOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.AssociateResourceSharePermission", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) AssociateResourceSharePermissionAsync(ctx workflow.Context, input *ram.AssociateResourceSharePermissionInput) *RamAssociateResourceSharePermissionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.AssociateResourceSharePermission", input)
	return &RamAssociateResourceSharePermissionFuture{Future: future}
}

func (a *RAMStub) CreateResourceShare(ctx workflow.Context, input *ram.CreateResourceShareInput) (*ram.CreateResourceShareOutput, error) {
	var output ram.CreateResourceShareOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.CreateResourceShare", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) CreateResourceShareAsync(ctx workflow.Context, input *ram.CreateResourceShareInput) *RamCreateResourceShareFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.CreateResourceShare", input)
	return &RamCreateResourceShareFuture{Future: future}
}

func (a *RAMStub) DeleteResourceShare(ctx workflow.Context, input *ram.DeleteResourceShareInput) (*ram.DeleteResourceShareOutput, error) {
	var output ram.DeleteResourceShareOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.DeleteResourceShare", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) DeleteResourceShareAsync(ctx workflow.Context, input *ram.DeleteResourceShareInput) *RamDeleteResourceShareFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.DeleteResourceShare", input)
	return &RamDeleteResourceShareFuture{Future: future}
}

func (a *RAMStub) DisassociateResourceShare(ctx workflow.Context, input *ram.DisassociateResourceShareInput) (*ram.DisassociateResourceShareOutput, error) {
	var output ram.DisassociateResourceShareOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.DisassociateResourceShare", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) DisassociateResourceShareAsync(ctx workflow.Context, input *ram.DisassociateResourceShareInput) *RamDisassociateResourceShareFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.DisassociateResourceShare", input)
	return &RamDisassociateResourceShareFuture{Future: future}
}

func (a *RAMStub) DisassociateResourceSharePermission(ctx workflow.Context, input *ram.DisassociateResourceSharePermissionInput) (*ram.DisassociateResourceSharePermissionOutput, error) {
	var output ram.DisassociateResourceSharePermissionOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.DisassociateResourceSharePermission", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) DisassociateResourceSharePermissionAsync(ctx workflow.Context, input *ram.DisassociateResourceSharePermissionInput) *RamDisassociateResourceSharePermissionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.DisassociateResourceSharePermission", input)
	return &RamDisassociateResourceSharePermissionFuture{Future: future}
}

func (a *RAMStub) EnableSharingWithAwsOrganization(ctx workflow.Context, input *ram.EnableSharingWithAwsOrganizationInput) (*ram.EnableSharingWithAwsOrganizationOutput, error) {
	var output ram.EnableSharingWithAwsOrganizationOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.EnableSharingWithAwsOrganization", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) EnableSharingWithAwsOrganizationAsync(ctx workflow.Context, input *ram.EnableSharingWithAwsOrganizationInput) *RamEnableSharingWithAwsOrganizationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.EnableSharingWithAwsOrganization", input)
	return &RamEnableSharingWithAwsOrganizationFuture{Future: future}
}

func (a *RAMStub) GetPermission(ctx workflow.Context, input *ram.GetPermissionInput) (*ram.GetPermissionOutput, error) {
	var output ram.GetPermissionOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.GetPermission", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) GetPermissionAsync(ctx workflow.Context, input *ram.GetPermissionInput) *RamGetPermissionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.GetPermission", input)
	return &RamGetPermissionFuture{Future: future}
}

func (a *RAMStub) GetResourcePolicies(ctx workflow.Context, input *ram.GetResourcePoliciesInput) (*ram.GetResourcePoliciesOutput, error) {
	var output ram.GetResourcePoliciesOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.GetResourcePolicies", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) GetResourcePoliciesAsync(ctx workflow.Context, input *ram.GetResourcePoliciesInput) *RamGetResourcePoliciesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.GetResourcePolicies", input)
	return &RamGetResourcePoliciesFuture{Future: future}
}

func (a *RAMStub) GetResourceShareAssociations(ctx workflow.Context, input *ram.GetResourceShareAssociationsInput) (*ram.GetResourceShareAssociationsOutput, error) {
	var output ram.GetResourceShareAssociationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.GetResourceShareAssociations", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) GetResourceShareAssociationsAsync(ctx workflow.Context, input *ram.GetResourceShareAssociationsInput) *RamGetResourceShareAssociationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.GetResourceShareAssociations", input)
	return &RamGetResourceShareAssociationsFuture{Future: future}
}

func (a *RAMStub) GetResourceShareInvitations(ctx workflow.Context, input *ram.GetResourceShareInvitationsInput) (*ram.GetResourceShareInvitationsOutput, error) {
	var output ram.GetResourceShareInvitationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.GetResourceShareInvitations", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) GetResourceShareInvitationsAsync(ctx workflow.Context, input *ram.GetResourceShareInvitationsInput) *RamGetResourceShareInvitationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.GetResourceShareInvitations", input)
	return &RamGetResourceShareInvitationsFuture{Future: future}
}

func (a *RAMStub) GetResourceShares(ctx workflow.Context, input *ram.GetResourceSharesInput) (*ram.GetResourceSharesOutput, error) {
	var output ram.GetResourceSharesOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.GetResourceShares", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) GetResourceSharesAsync(ctx workflow.Context, input *ram.GetResourceSharesInput) *RamGetResourceSharesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.GetResourceShares", input)
	return &RamGetResourceSharesFuture{Future: future}
}

func (a *RAMStub) ListPendingInvitationResources(ctx workflow.Context, input *ram.ListPendingInvitationResourcesInput) (*ram.ListPendingInvitationResourcesOutput, error) {
	var output ram.ListPendingInvitationResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.ListPendingInvitationResources", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) ListPendingInvitationResourcesAsync(ctx workflow.Context, input *ram.ListPendingInvitationResourcesInput) *RamListPendingInvitationResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.ListPendingInvitationResources", input)
	return &RamListPendingInvitationResourcesFuture{Future: future}
}

func (a *RAMStub) ListPermissions(ctx workflow.Context, input *ram.ListPermissionsInput) (*ram.ListPermissionsOutput, error) {
	var output ram.ListPermissionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.ListPermissions", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) ListPermissionsAsync(ctx workflow.Context, input *ram.ListPermissionsInput) *RamListPermissionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.ListPermissions", input)
	return &RamListPermissionsFuture{Future: future}
}

func (a *RAMStub) ListPrincipals(ctx workflow.Context, input *ram.ListPrincipalsInput) (*ram.ListPrincipalsOutput, error) {
	var output ram.ListPrincipalsOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.ListPrincipals", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) ListPrincipalsAsync(ctx workflow.Context, input *ram.ListPrincipalsInput) *RamListPrincipalsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.ListPrincipals", input)
	return &RamListPrincipalsFuture{Future: future}
}

func (a *RAMStub) ListResourceSharePermissions(ctx workflow.Context, input *ram.ListResourceSharePermissionsInput) (*ram.ListResourceSharePermissionsOutput, error) {
	var output ram.ListResourceSharePermissionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.ListResourceSharePermissions", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) ListResourceSharePermissionsAsync(ctx workflow.Context, input *ram.ListResourceSharePermissionsInput) *RamListResourceSharePermissionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.ListResourceSharePermissions", input)
	return &RamListResourceSharePermissionsFuture{Future: future}
}

func (a *RAMStub) ListResourceTypes(ctx workflow.Context, input *ram.ListResourceTypesInput) (*ram.ListResourceTypesOutput, error) {
	var output ram.ListResourceTypesOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.ListResourceTypes", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) ListResourceTypesAsync(ctx workflow.Context, input *ram.ListResourceTypesInput) *RamListResourceTypesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.ListResourceTypes", input)
	return &RamListResourceTypesFuture{Future: future}
}

func (a *RAMStub) ListResources(ctx workflow.Context, input *ram.ListResourcesInput) (*ram.ListResourcesOutput, error) {
	var output ram.ListResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.ListResources", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) ListResourcesAsync(ctx workflow.Context, input *ram.ListResourcesInput) *RamListResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.ListResources", input)
	return &RamListResourcesFuture{Future: future}
}

func (a *RAMStub) PromoteResourceShareCreatedFromPolicy(ctx workflow.Context, input *ram.PromoteResourceShareCreatedFromPolicyInput) (*ram.PromoteResourceShareCreatedFromPolicyOutput, error) {
	var output ram.PromoteResourceShareCreatedFromPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.PromoteResourceShareCreatedFromPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) PromoteResourceShareCreatedFromPolicyAsync(ctx workflow.Context, input *ram.PromoteResourceShareCreatedFromPolicyInput) *RamPromoteResourceShareCreatedFromPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.PromoteResourceShareCreatedFromPolicy", input)
	return &RamPromoteResourceShareCreatedFromPolicyFuture{Future: future}
}

func (a *RAMStub) RejectResourceShareInvitation(ctx workflow.Context, input *ram.RejectResourceShareInvitationInput) (*ram.RejectResourceShareInvitationOutput, error) {
	var output ram.RejectResourceShareInvitationOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.RejectResourceShareInvitation", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) RejectResourceShareInvitationAsync(ctx workflow.Context, input *ram.RejectResourceShareInvitationInput) *RamRejectResourceShareInvitationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.RejectResourceShareInvitation", input)
	return &RamRejectResourceShareInvitationFuture{Future: future}
}

func (a *RAMStub) TagResource(ctx workflow.Context, input *ram.TagResourceInput) (*ram.TagResourceOutput, error) {
	var output ram.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) TagResourceAsync(ctx workflow.Context, input *ram.TagResourceInput) *RamTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.TagResource", input)
	return &RamTagResourceFuture{Future: future}
}

func (a *RAMStub) UntagResource(ctx workflow.Context, input *ram.UntagResourceInput) (*ram.UntagResourceOutput, error) {
	var output ram.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) UntagResourceAsync(ctx workflow.Context, input *ram.UntagResourceInput) *RamUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.UntagResource", input)
	return &RamUntagResourceFuture{Future: future}
}

func (a *RAMStub) UpdateResourceShare(ctx workflow.Context, input *ram.UpdateResourceShareInput) (*ram.UpdateResourceShareOutput, error) {
	var output ram.UpdateResourceShareOutput
	err := workflow.ExecuteActivity(ctx, "aws.ram.UpdateResourceShare", input).Get(ctx, &output)
	return &output, err
}

func (a *RAMStub) UpdateResourceShareAsync(ctx workflow.Context, input *ram.UpdateResourceShareInput) *RamUpdateResourceShareFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ram.UpdateResourceShare", input)
	return &RamUpdateResourceShareFuture{Future: future}
}
