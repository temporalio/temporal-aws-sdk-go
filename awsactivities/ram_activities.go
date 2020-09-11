
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ram"
	"github.com/aws/aws-sdk-go/service/ram/ramiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type RAMActivities struct {
    client ramiface.RAMAPI
}

func NewRAMActivities(session *session.Session, config ...*aws.Config) *RAMActivities {
    client := ram.New(session, config...)
    return &RAMActivities{client: client}
}

func (a *RAMActivities) AcceptResourceShareInvitation(ctx context.Context, input *ram.AcceptResourceShareInvitationInput) (*ram.AcceptResourceShareInvitationOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.AcceptResourceShareInvitationWithContext(ctx, input)
}

func (a *RAMActivities) AssociateResourceShare(ctx context.Context, input *ram.AssociateResourceShareInput) (*ram.AssociateResourceShareOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.AssociateResourceShareWithContext(ctx, input)
}

func (a *RAMActivities) AssociateResourceSharePermission(ctx context.Context, input *ram.AssociateResourceSharePermissionInput) (*ram.AssociateResourceSharePermissionOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.AssociateResourceSharePermissionWithContext(ctx, input)
}

func (a *RAMActivities) CreateResourceShare(ctx context.Context, input *ram.CreateResourceShareInput) (*ram.CreateResourceShareOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.CreateResourceShareWithContext(ctx, input)
}

func (a *RAMActivities) DeleteResourceShare(ctx context.Context, input *ram.DeleteResourceShareInput) (*ram.DeleteResourceShareOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.DeleteResourceShareWithContext(ctx, input)
}

func (a *RAMActivities) DisassociateResourceShare(ctx context.Context, input *ram.DisassociateResourceShareInput) (*ram.DisassociateResourceShareOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.DisassociateResourceShareWithContext(ctx, input)
}

func (a *RAMActivities) DisassociateResourceSharePermission(ctx context.Context, input *ram.DisassociateResourceSharePermissionInput) (*ram.DisassociateResourceSharePermissionOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.DisassociateResourceSharePermissionWithContext(ctx, input)
}

func (a *RAMActivities) EnableSharingWithAwsOrganization(ctx context.Context, input *ram.EnableSharingWithAwsOrganizationInput) (*ram.EnableSharingWithAwsOrganizationOutput, error) {
    return a.client.EnableSharingWithAwsOrganizationWithContext(ctx, input)
}

func (a *RAMActivities) GetPermission(ctx context.Context, input *ram.GetPermissionInput) (*ram.GetPermissionOutput, error) {
    return a.client.GetPermissionWithContext(ctx, input)
}

func (a *RAMActivities) GetResourcePolicies(ctx context.Context, input *ram.GetResourcePoliciesInput) (*ram.GetResourcePoliciesOutput, error) {
    return a.client.GetResourcePoliciesWithContext(ctx, input)
}

func (a *RAMActivities) GetResourceShareAssociations(ctx context.Context, input *ram.GetResourceShareAssociationsInput) (*ram.GetResourceShareAssociationsOutput, error) {
    return a.client.GetResourceShareAssociationsWithContext(ctx, input)
}

func (a *RAMActivities) GetResourceShareInvitations(ctx context.Context, input *ram.GetResourceShareInvitationsInput) (*ram.GetResourceShareInvitationsOutput, error) {
    return a.client.GetResourceShareInvitationsWithContext(ctx, input)
}

func (a *RAMActivities) GetResourceShares(ctx context.Context, input *ram.GetResourceSharesInput) (*ram.GetResourceSharesOutput, error) {
    return a.client.GetResourceSharesWithContext(ctx, input)
}

func (a *RAMActivities) ListPendingInvitationResources(ctx context.Context, input *ram.ListPendingInvitationResourcesInput) (*ram.ListPendingInvitationResourcesOutput, error) {
    return a.client.ListPendingInvitationResourcesWithContext(ctx, input)
}

func (a *RAMActivities) ListPermissions(ctx context.Context, input *ram.ListPermissionsInput) (*ram.ListPermissionsOutput, error) {
    return a.client.ListPermissionsWithContext(ctx, input)
}

func (a *RAMActivities) ListPrincipals(ctx context.Context, input *ram.ListPrincipalsInput) (*ram.ListPrincipalsOutput, error) {
    return a.client.ListPrincipalsWithContext(ctx, input)
}

func (a *RAMActivities) ListResourceSharePermissions(ctx context.Context, input *ram.ListResourceSharePermissionsInput) (*ram.ListResourceSharePermissionsOutput, error) {
    return a.client.ListResourceSharePermissionsWithContext(ctx, input)
}

func (a *RAMActivities) ListResourceTypes(ctx context.Context, input *ram.ListResourceTypesInput) (*ram.ListResourceTypesOutput, error) {
    return a.client.ListResourceTypesWithContext(ctx, input)
}

func (a *RAMActivities) ListResources(ctx context.Context, input *ram.ListResourcesInput) (*ram.ListResourcesOutput, error) {
    return a.client.ListResourcesWithContext(ctx, input)
}

func (a *RAMActivities) PromoteResourceShareCreatedFromPolicy(ctx context.Context, input *ram.PromoteResourceShareCreatedFromPolicyInput) (*ram.PromoteResourceShareCreatedFromPolicyOutput, error) {
    return a.client.PromoteResourceShareCreatedFromPolicyWithContext(ctx, input)
}

func (a *RAMActivities) RejectResourceShareInvitation(ctx context.Context, input *ram.RejectResourceShareInvitationInput) (*ram.RejectResourceShareInvitationOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.RejectResourceShareInvitationWithContext(ctx, input)
}

func (a *RAMActivities) TagResource(ctx context.Context, input *ram.TagResourceInput) (*ram.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *RAMActivities) UntagResource(ctx context.Context, input *ram.UntagResourceInput) (*ram.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *RAMActivities) UpdateResourceShare(ctx context.Context, input *ram.UpdateResourceShareInput) (*ram.UpdateResourceShareOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.UpdateResourceShareWithContext(ctx, input)
}
