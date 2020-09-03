package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/ram"
	"github.com/aws/aws-sdk-go/service/ram/ramiface"
)

type RAMActivities struct {
	client ramiface.RAMAPI
}

func NewRAMActivities(client ramiface.RAMAPI) *RAMActivities {
    return &RAMActivities{client: client}
}

func (a *RAMActivities) AcceptResourceShareInvitation(input *ram.AcceptResourceShareInvitationInput) (*ram.AcceptResourceShareInvitationOutput, error) {
    return a.client.AcceptResourceShareInvitation(input)
}

func (a *RAMActivities) AssociateResourceShare(input *ram.AssociateResourceShareInput) (*ram.AssociateResourceShareOutput, error) {
    return a.client.AssociateResourceShare(input)
}

func (a *RAMActivities) AssociateResourceSharePermission(input *ram.AssociateResourceSharePermissionInput) (*ram.AssociateResourceSharePermissionOutput, error) {
    return a.client.AssociateResourceSharePermission(input)
}

func (a *RAMActivities) CreateResourceShare(input *ram.CreateResourceShareInput) (*ram.CreateResourceShareOutput, error) {
    return a.client.CreateResourceShare(input)
}

func (a *RAMActivities) DeleteResourceShare(input *ram.DeleteResourceShareInput) (*ram.DeleteResourceShareOutput, error) {
    return a.client.DeleteResourceShare(input)
}

func (a *RAMActivities) DisassociateResourceShare(input *ram.DisassociateResourceShareInput) (*ram.DisassociateResourceShareOutput, error) {
    return a.client.DisassociateResourceShare(input)
}

func (a *RAMActivities) DisassociateResourceSharePermission(input *ram.DisassociateResourceSharePermissionInput) (*ram.DisassociateResourceSharePermissionOutput, error) {
    return a.client.DisassociateResourceSharePermission(input)
}

func (a *RAMActivities) EnableSharingWithAwsOrganization(input *ram.EnableSharingWithAwsOrganizationInput) (*ram.EnableSharingWithAwsOrganizationOutput, error) {
    return a.client.EnableSharingWithAwsOrganization(input)
}

func (a *RAMActivities) GetPermission(input *ram.GetPermissionInput) (*ram.GetPermissionOutput, error) {
    return a.client.GetPermission(input)
}

func (a *RAMActivities) GetResourcePolicies(input *ram.GetResourcePoliciesInput) (*ram.GetResourcePoliciesOutput, error) {
    return a.client.GetResourcePolicies(input)
}

func (a *RAMActivities) GetResourceShareAssociations(input *ram.GetResourceShareAssociationsInput) (*ram.GetResourceShareAssociationsOutput, error) {
    return a.client.GetResourceShareAssociations(input)
}

func (a *RAMActivities) GetResourceShareInvitations(input *ram.GetResourceShareInvitationsInput) (*ram.GetResourceShareInvitationsOutput, error) {
    return a.client.GetResourceShareInvitations(input)
}

func (a *RAMActivities) GetResourceShares(input *ram.GetResourceSharesInput) (*ram.GetResourceSharesOutput, error) {
    return a.client.GetResourceShares(input)
}

func (a *RAMActivities) ListPendingInvitationResources(input *ram.ListPendingInvitationResourcesInput) (*ram.ListPendingInvitationResourcesOutput, error) {
    return a.client.ListPendingInvitationResources(input)
}

func (a *RAMActivities) ListPermissions(input *ram.ListPermissionsInput) (*ram.ListPermissionsOutput, error) {
    return a.client.ListPermissions(input)
}

func (a *RAMActivities) ListPrincipals(input *ram.ListPrincipalsInput) (*ram.ListPrincipalsOutput, error) {
    return a.client.ListPrincipals(input)
}

func (a *RAMActivities) ListResourceSharePermissions(input *ram.ListResourceSharePermissionsInput) (*ram.ListResourceSharePermissionsOutput, error) {
    return a.client.ListResourceSharePermissions(input)
}

func (a *RAMActivities) ListResourceTypes(input *ram.ListResourceTypesInput) (*ram.ListResourceTypesOutput, error) {
    return a.client.ListResourceTypes(input)
}

func (a *RAMActivities) ListResources(input *ram.ListResourcesInput) (*ram.ListResourcesOutput, error) {
    return a.client.ListResources(input)
}

func (a *RAMActivities) PromoteResourceShareCreatedFromPolicy(input *ram.PromoteResourceShareCreatedFromPolicyInput) (*ram.PromoteResourceShareCreatedFromPolicyOutput, error) {
    return a.client.PromoteResourceShareCreatedFromPolicy(input)
}

func (a *RAMActivities) RejectResourceShareInvitation(input *ram.RejectResourceShareInvitationInput) (*ram.RejectResourceShareInvitationOutput, error) {
    return a.client.RejectResourceShareInvitation(input)
}

func (a *RAMActivities) TagResource(input *ram.TagResourceInput) (*ram.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *RAMActivities) UntagResource(input *ram.UntagResourceInput) (*ram.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *RAMActivities) UpdateResourceShare(input *ram.UpdateResourceShareInput) (*ram.UpdateResourceShareOutput, error) {
    return a.client.UpdateResourceShare(input)
}
