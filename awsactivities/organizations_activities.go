package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/organizations"
	"github.com/aws/aws-sdk-go/service/organizations/organizationsiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type OrganizationsActivities struct {
	client organizationsiface.OrganizationsAPI
}

func NewOrganizationsActivities(session *session.Session, config ...*aws.Config) *OrganizationsActivities {
	client := organizations.New(session, config...)
	return &OrganizationsActivities{client: client}
}

func (a *OrganizationsActivities) AcceptHandshake(ctx context.Context, input *organizations.AcceptHandshakeInput) (*organizations.AcceptHandshakeOutput, error) {
	return a.client.AcceptHandshakeWithContext(ctx, input)
}

func (a *OrganizationsActivities) AttachPolicy(ctx context.Context, input *organizations.AttachPolicyInput) (*organizations.AttachPolicyOutput, error) {
	return a.client.AttachPolicyWithContext(ctx, input)
}

func (a *OrganizationsActivities) CancelHandshake(ctx context.Context, input *organizations.CancelHandshakeInput) (*organizations.CancelHandshakeOutput, error) {
	return a.client.CancelHandshakeWithContext(ctx, input)
}

func (a *OrganizationsActivities) CreateAccount(ctx context.Context, input *organizations.CreateAccountInput) (*organizations.CreateAccountOutput, error) {
	return a.client.CreateAccountWithContext(ctx, input)
}

func (a *OrganizationsActivities) CreateGovCloudAccount(ctx context.Context, input *organizations.CreateGovCloudAccountInput) (*organizations.CreateGovCloudAccountOutput, error) {
	return a.client.CreateGovCloudAccountWithContext(ctx, input)
}

func (a *OrganizationsActivities) CreateOrganization(ctx context.Context, input *organizations.CreateOrganizationInput) (*organizations.CreateOrganizationOutput, error) {
	return a.client.CreateOrganizationWithContext(ctx, input)
}

func (a *OrganizationsActivities) CreateOrganizationalUnit(ctx context.Context, input *organizations.CreateOrganizationalUnitInput) (*organizations.CreateOrganizationalUnitOutput, error) {
	return a.client.CreateOrganizationalUnitWithContext(ctx, input)
}

func (a *OrganizationsActivities) CreatePolicy(ctx context.Context, input *organizations.CreatePolicyInput) (*organizations.CreatePolicyOutput, error) {
	return a.client.CreatePolicyWithContext(ctx, input)
}

func (a *OrganizationsActivities) DeclineHandshake(ctx context.Context, input *organizations.DeclineHandshakeInput) (*organizations.DeclineHandshakeOutput, error) {
	return a.client.DeclineHandshakeWithContext(ctx, input)
}

func (a *OrganizationsActivities) DeleteOrganization(ctx context.Context, input *organizations.DeleteOrganizationInput) (*organizations.DeleteOrganizationOutput, error) {
	return a.client.DeleteOrganizationWithContext(ctx, input)
}

func (a *OrganizationsActivities) DeleteOrganizationalUnit(ctx context.Context, input *organizations.DeleteOrganizationalUnitInput) (*organizations.DeleteOrganizationalUnitOutput, error) {
	return a.client.DeleteOrganizationalUnitWithContext(ctx, input)
}

func (a *OrganizationsActivities) DeletePolicy(ctx context.Context, input *organizations.DeletePolicyInput) (*organizations.DeletePolicyOutput, error) {
	return a.client.DeletePolicyWithContext(ctx, input)
}

func (a *OrganizationsActivities) DeregisterDelegatedAdministrator(ctx context.Context, input *organizations.DeregisterDelegatedAdministratorInput) (*organizations.DeregisterDelegatedAdministratorOutput, error) {
	return a.client.DeregisterDelegatedAdministratorWithContext(ctx, input)
}

func (a *OrganizationsActivities) DescribeAccount(ctx context.Context, input *organizations.DescribeAccountInput) (*organizations.DescribeAccountOutput, error) {
	return a.client.DescribeAccountWithContext(ctx, input)
}

func (a *OrganizationsActivities) DescribeCreateAccountStatus(ctx context.Context, input *organizations.DescribeCreateAccountStatusInput) (*organizations.DescribeCreateAccountStatusOutput, error) {
	return a.client.DescribeCreateAccountStatusWithContext(ctx, input)
}

func (a *OrganizationsActivities) DescribeEffectivePolicy(ctx context.Context, input *organizations.DescribeEffectivePolicyInput) (*organizations.DescribeEffectivePolicyOutput, error) {
	return a.client.DescribeEffectivePolicyWithContext(ctx, input)
}

func (a *OrganizationsActivities) DescribeHandshake(ctx context.Context, input *organizations.DescribeHandshakeInput) (*organizations.DescribeHandshakeOutput, error) {
	return a.client.DescribeHandshakeWithContext(ctx, input)
}

func (a *OrganizationsActivities) DescribeOrganization(ctx context.Context, input *organizations.DescribeOrganizationInput) (*organizations.DescribeOrganizationOutput, error) {
	return a.client.DescribeOrganizationWithContext(ctx, input)
}

func (a *OrganizationsActivities) DescribeOrganizationalUnit(ctx context.Context, input *organizations.DescribeOrganizationalUnitInput) (*organizations.DescribeOrganizationalUnitOutput, error) {
	return a.client.DescribeOrganizationalUnitWithContext(ctx, input)
}

func (a *OrganizationsActivities) DescribePolicy(ctx context.Context, input *organizations.DescribePolicyInput) (*organizations.DescribePolicyOutput, error) {
	return a.client.DescribePolicyWithContext(ctx, input)
}

func (a *OrganizationsActivities) DetachPolicy(ctx context.Context, input *organizations.DetachPolicyInput) (*organizations.DetachPolicyOutput, error) {
	return a.client.DetachPolicyWithContext(ctx, input)
}

func (a *OrganizationsActivities) DisableAWSServiceAccess(ctx context.Context, input *organizations.DisableAWSServiceAccessInput) (*organizations.DisableAWSServiceAccessOutput, error) {
	return a.client.DisableAWSServiceAccessWithContext(ctx, input)
}

func (a *OrganizationsActivities) DisablePolicyType(ctx context.Context, input *organizations.DisablePolicyTypeInput) (*organizations.DisablePolicyTypeOutput, error) {
	return a.client.DisablePolicyTypeWithContext(ctx, input)
}

func (a *OrganizationsActivities) EnableAWSServiceAccess(ctx context.Context, input *organizations.EnableAWSServiceAccessInput) (*organizations.EnableAWSServiceAccessOutput, error) {
	return a.client.EnableAWSServiceAccessWithContext(ctx, input)
}

func (a *OrganizationsActivities) EnableAllFeatures(ctx context.Context, input *organizations.EnableAllFeaturesInput) (*organizations.EnableAllFeaturesOutput, error) {
	return a.client.EnableAllFeaturesWithContext(ctx, input)
}

func (a *OrganizationsActivities) EnablePolicyType(ctx context.Context, input *organizations.EnablePolicyTypeInput) (*organizations.EnablePolicyTypeOutput, error) {
	return a.client.EnablePolicyTypeWithContext(ctx, input)
}

func (a *OrganizationsActivities) InviteAccountToOrganization(ctx context.Context, input *organizations.InviteAccountToOrganizationInput) (*organizations.InviteAccountToOrganizationOutput, error) {
	return a.client.InviteAccountToOrganizationWithContext(ctx, input)
}

func (a *OrganizationsActivities) LeaveOrganization(ctx context.Context, input *organizations.LeaveOrganizationInput) (*organizations.LeaveOrganizationOutput, error) {
	return a.client.LeaveOrganizationWithContext(ctx, input)
}

func (a *OrganizationsActivities) ListAWSServiceAccessForOrganization(ctx context.Context, input *organizations.ListAWSServiceAccessForOrganizationInput) (*organizations.ListAWSServiceAccessForOrganizationOutput, error) {
	return a.client.ListAWSServiceAccessForOrganizationWithContext(ctx, input)
}

func (a *OrganizationsActivities) ListAccounts(ctx context.Context, input *organizations.ListAccountsInput) (*organizations.ListAccountsOutput, error) {
	return a.client.ListAccountsWithContext(ctx, input)
}

func (a *OrganizationsActivities) ListAccountsForParent(ctx context.Context, input *organizations.ListAccountsForParentInput) (*organizations.ListAccountsForParentOutput, error) {
	return a.client.ListAccountsForParentWithContext(ctx, input)
}

func (a *OrganizationsActivities) ListChildren(ctx context.Context, input *organizations.ListChildrenInput) (*organizations.ListChildrenOutput, error) {
	return a.client.ListChildrenWithContext(ctx, input)
}

func (a *OrganizationsActivities) ListCreateAccountStatus(ctx context.Context, input *organizations.ListCreateAccountStatusInput) (*organizations.ListCreateAccountStatusOutput, error) {
	return a.client.ListCreateAccountStatusWithContext(ctx, input)
}

func (a *OrganizationsActivities) ListDelegatedAdministrators(ctx context.Context, input *organizations.ListDelegatedAdministratorsInput) (*organizations.ListDelegatedAdministratorsOutput, error) {
	return a.client.ListDelegatedAdministratorsWithContext(ctx, input)
}

func (a *OrganizationsActivities) ListDelegatedServicesForAccount(ctx context.Context, input *organizations.ListDelegatedServicesForAccountInput) (*organizations.ListDelegatedServicesForAccountOutput, error) {
	return a.client.ListDelegatedServicesForAccountWithContext(ctx, input)
}

func (a *OrganizationsActivities) ListHandshakesForAccount(ctx context.Context, input *organizations.ListHandshakesForAccountInput) (*organizations.ListHandshakesForAccountOutput, error) {
	return a.client.ListHandshakesForAccountWithContext(ctx, input)
}

func (a *OrganizationsActivities) ListHandshakesForOrganization(ctx context.Context, input *organizations.ListHandshakesForOrganizationInput) (*organizations.ListHandshakesForOrganizationOutput, error) {
	return a.client.ListHandshakesForOrganizationWithContext(ctx, input)
}

func (a *OrganizationsActivities) ListOrganizationalUnitsForParent(ctx context.Context, input *organizations.ListOrganizationalUnitsForParentInput) (*organizations.ListOrganizationalUnitsForParentOutput, error) {
	return a.client.ListOrganizationalUnitsForParentWithContext(ctx, input)
}

func (a *OrganizationsActivities) ListParents(ctx context.Context, input *organizations.ListParentsInput) (*organizations.ListParentsOutput, error) {
	return a.client.ListParentsWithContext(ctx, input)
}

func (a *OrganizationsActivities) ListPolicies(ctx context.Context, input *organizations.ListPoliciesInput) (*organizations.ListPoliciesOutput, error) {
	return a.client.ListPoliciesWithContext(ctx, input)
}

func (a *OrganizationsActivities) ListPoliciesForTarget(ctx context.Context, input *organizations.ListPoliciesForTargetInput) (*organizations.ListPoliciesForTargetOutput, error) {
	return a.client.ListPoliciesForTargetWithContext(ctx, input)
}

func (a *OrganizationsActivities) ListRoots(ctx context.Context, input *organizations.ListRootsInput) (*organizations.ListRootsOutput, error) {
	return a.client.ListRootsWithContext(ctx, input)
}

func (a *OrganizationsActivities) ListTagsForResource(ctx context.Context, input *organizations.ListTagsForResourceInput) (*organizations.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *OrganizationsActivities) ListTargetsForPolicy(ctx context.Context, input *organizations.ListTargetsForPolicyInput) (*organizations.ListTargetsForPolicyOutput, error) {
	return a.client.ListTargetsForPolicyWithContext(ctx, input)
}

func (a *OrganizationsActivities) MoveAccount(ctx context.Context, input *organizations.MoveAccountInput) (*organizations.MoveAccountOutput, error) {
	return a.client.MoveAccountWithContext(ctx, input)
}

func (a *OrganizationsActivities) RegisterDelegatedAdministrator(ctx context.Context, input *organizations.RegisterDelegatedAdministratorInput) (*organizations.RegisterDelegatedAdministratorOutput, error) {
	return a.client.RegisterDelegatedAdministratorWithContext(ctx, input)
}

func (a *OrganizationsActivities) RemoveAccountFromOrganization(ctx context.Context, input *organizations.RemoveAccountFromOrganizationInput) (*organizations.RemoveAccountFromOrganizationOutput, error) {
	return a.client.RemoveAccountFromOrganizationWithContext(ctx, input)
}

func (a *OrganizationsActivities) TagResource(ctx context.Context, input *organizations.TagResourceInput) (*organizations.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *OrganizationsActivities) UntagResource(ctx context.Context, input *organizations.UntagResourceInput) (*organizations.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *OrganizationsActivities) UpdateOrganizationalUnit(ctx context.Context, input *organizations.UpdateOrganizationalUnitInput) (*organizations.UpdateOrganizationalUnitOutput, error) {
	return a.client.UpdateOrganizationalUnitWithContext(ctx, input)
}

func (a *OrganizationsActivities) UpdatePolicy(ctx context.Context, input *organizations.UpdatePolicyInput) (*organizations.UpdatePolicyOutput, error) {
	return a.client.UpdatePolicyWithContext(ctx, input)
}
