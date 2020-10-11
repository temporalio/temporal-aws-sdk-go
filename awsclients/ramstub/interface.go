// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package ramstub

import (
	"github.com/aws/aws-sdk-go/service/ram"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	AcceptResourceShareInvitation(ctx workflow.Context, input *ram.AcceptResourceShareInvitationInput) (*ram.AcceptResourceShareInvitationOutput, error)
	AcceptResourceShareInvitationAsync(ctx workflow.Context, input *ram.AcceptResourceShareInvitationInput) *RAMAcceptResourceShareInvitationFuture

	AssociateResourceShare(ctx workflow.Context, input *ram.AssociateResourceShareInput) (*ram.AssociateResourceShareOutput, error)
	AssociateResourceShareAsync(ctx workflow.Context, input *ram.AssociateResourceShareInput) *RAMAssociateResourceShareFuture

	AssociateResourceSharePermission(ctx workflow.Context, input *ram.AssociateResourceSharePermissionInput) (*ram.AssociateResourceSharePermissionOutput, error)
	AssociateResourceSharePermissionAsync(ctx workflow.Context, input *ram.AssociateResourceSharePermissionInput) *RAMAssociateResourceSharePermissionFuture

	CreateResourceShare(ctx workflow.Context, input *ram.CreateResourceShareInput) (*ram.CreateResourceShareOutput, error)
	CreateResourceShareAsync(ctx workflow.Context, input *ram.CreateResourceShareInput) *RAMCreateResourceShareFuture

	DeleteResourceShare(ctx workflow.Context, input *ram.DeleteResourceShareInput) (*ram.DeleteResourceShareOutput, error)
	DeleteResourceShareAsync(ctx workflow.Context, input *ram.DeleteResourceShareInput) *RAMDeleteResourceShareFuture

	DisassociateResourceShare(ctx workflow.Context, input *ram.DisassociateResourceShareInput) (*ram.DisassociateResourceShareOutput, error)
	DisassociateResourceShareAsync(ctx workflow.Context, input *ram.DisassociateResourceShareInput) *RAMDisassociateResourceShareFuture

	DisassociateResourceSharePermission(ctx workflow.Context, input *ram.DisassociateResourceSharePermissionInput) (*ram.DisassociateResourceSharePermissionOutput, error)
	DisassociateResourceSharePermissionAsync(ctx workflow.Context, input *ram.DisassociateResourceSharePermissionInput) *RAMDisassociateResourceSharePermissionFuture

	EnableSharingWithAwsOrganization(ctx workflow.Context, input *ram.EnableSharingWithAwsOrganizationInput) (*ram.EnableSharingWithAwsOrganizationOutput, error)
	EnableSharingWithAwsOrganizationAsync(ctx workflow.Context, input *ram.EnableSharingWithAwsOrganizationInput) *RAMEnableSharingWithAwsOrganizationFuture

	GetPermission(ctx workflow.Context, input *ram.GetPermissionInput) (*ram.GetPermissionOutput, error)
	GetPermissionAsync(ctx workflow.Context, input *ram.GetPermissionInput) *RAMGetPermissionFuture

	GetResourcePolicies(ctx workflow.Context, input *ram.GetResourcePoliciesInput) (*ram.GetResourcePoliciesOutput, error)
	GetResourcePoliciesAsync(ctx workflow.Context, input *ram.GetResourcePoliciesInput) *RAMGetResourcePoliciesFuture

	GetResourceShareAssociations(ctx workflow.Context, input *ram.GetResourceShareAssociationsInput) (*ram.GetResourceShareAssociationsOutput, error)
	GetResourceShareAssociationsAsync(ctx workflow.Context, input *ram.GetResourceShareAssociationsInput) *RAMGetResourceShareAssociationsFuture

	GetResourceShareInvitations(ctx workflow.Context, input *ram.GetResourceShareInvitationsInput) (*ram.GetResourceShareInvitationsOutput, error)
	GetResourceShareInvitationsAsync(ctx workflow.Context, input *ram.GetResourceShareInvitationsInput) *RAMGetResourceShareInvitationsFuture

	GetResourceShares(ctx workflow.Context, input *ram.GetResourceSharesInput) (*ram.GetResourceSharesOutput, error)
	GetResourceSharesAsync(ctx workflow.Context, input *ram.GetResourceSharesInput) *RAMGetResourceSharesFuture

	ListPendingInvitationResources(ctx workflow.Context, input *ram.ListPendingInvitationResourcesInput) (*ram.ListPendingInvitationResourcesOutput, error)
	ListPendingInvitationResourcesAsync(ctx workflow.Context, input *ram.ListPendingInvitationResourcesInput) *RAMListPendingInvitationResourcesFuture

	ListPermissions(ctx workflow.Context, input *ram.ListPermissionsInput) (*ram.ListPermissionsOutput, error)
	ListPermissionsAsync(ctx workflow.Context, input *ram.ListPermissionsInput) *RAMListPermissionsFuture

	ListPrincipals(ctx workflow.Context, input *ram.ListPrincipalsInput) (*ram.ListPrincipalsOutput, error)
	ListPrincipalsAsync(ctx workflow.Context, input *ram.ListPrincipalsInput) *RAMListPrincipalsFuture

	ListResourceSharePermissions(ctx workflow.Context, input *ram.ListResourceSharePermissionsInput) (*ram.ListResourceSharePermissionsOutput, error)
	ListResourceSharePermissionsAsync(ctx workflow.Context, input *ram.ListResourceSharePermissionsInput) *RAMListResourceSharePermissionsFuture

	ListResourceTypes(ctx workflow.Context, input *ram.ListResourceTypesInput) (*ram.ListResourceTypesOutput, error)
	ListResourceTypesAsync(ctx workflow.Context, input *ram.ListResourceTypesInput) *RAMListResourceTypesFuture

	ListResources(ctx workflow.Context, input *ram.ListResourcesInput) (*ram.ListResourcesOutput, error)
	ListResourcesAsync(ctx workflow.Context, input *ram.ListResourcesInput) *RAMListResourcesFuture

	PromoteResourceShareCreatedFromPolicy(ctx workflow.Context, input *ram.PromoteResourceShareCreatedFromPolicyInput) (*ram.PromoteResourceShareCreatedFromPolicyOutput, error)
	PromoteResourceShareCreatedFromPolicyAsync(ctx workflow.Context, input *ram.PromoteResourceShareCreatedFromPolicyInput) *RAMPromoteResourceShareCreatedFromPolicyFuture

	RejectResourceShareInvitation(ctx workflow.Context, input *ram.RejectResourceShareInvitationInput) (*ram.RejectResourceShareInvitationOutput, error)
	RejectResourceShareInvitationAsync(ctx workflow.Context, input *ram.RejectResourceShareInvitationInput) *RAMRejectResourceShareInvitationFuture

	TagResource(ctx workflow.Context, input *ram.TagResourceInput) (*ram.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *ram.TagResourceInput) *RAMTagResourceFuture

	UntagResource(ctx workflow.Context, input *ram.UntagResourceInput) (*ram.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *ram.UntagResourceInput) *RAMUntagResourceFuture

	UpdateResourceShare(ctx workflow.Context, input *ram.UpdateResourceShareInput) (*ram.UpdateResourceShareOutput, error)
	UpdateResourceShareAsync(ctx workflow.Context, input *ram.UpdateResourceShareInput) *RAMUpdateResourceShareFuture
}

func NewClient() Client {
	return &stub{}
}
