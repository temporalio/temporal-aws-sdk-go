package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/organizations"
	"github.com/aws/aws-sdk-go/service/organizations/organizationsiface"
)

type OrganizationsActivities struct {
	client organizationsiface.OrganizationsAPI
}

func NewOrganizationsActivities(session *session.Session, config ...*aws.Config) *OrganizationsActivities {
	client := organizations.New(session, config...)
	return &OrganizationsActivities{client: client}
}

func (a *OrganizationsActivities) AcceptHandshake(input *organizations.AcceptHandshakeInput) (*organizations.AcceptHandshakeOutput, error) {
	return a.client.AcceptHandshake(input)
}

func (a *OrganizationsActivities) AttachPolicy(input *organizations.AttachPolicyInput) (*organizations.AttachPolicyOutput, error) {
	return a.client.AttachPolicy(input)
}

func (a *OrganizationsActivities) CancelHandshake(input *organizations.CancelHandshakeInput) (*organizations.CancelHandshakeOutput, error) {
	return a.client.CancelHandshake(input)
}

func (a *OrganizationsActivities) CreateAccount(input *organizations.CreateAccountInput) (*organizations.CreateAccountOutput, error) {
	return a.client.CreateAccount(input)
}

func (a *OrganizationsActivities) CreateGovCloudAccount(input *organizations.CreateGovCloudAccountInput) (*organizations.CreateGovCloudAccountOutput, error) {
	return a.client.CreateGovCloudAccount(input)
}

func (a *OrganizationsActivities) CreateOrganization(input *organizations.CreateOrganizationInput) (*organizations.CreateOrganizationOutput, error) {
	return a.client.CreateOrganization(input)
}

func (a *OrganizationsActivities) CreateOrganizationalUnit(input *organizations.CreateOrganizationalUnitInput) (*organizations.CreateOrganizationalUnitOutput, error) {
	return a.client.CreateOrganizationalUnit(input)
}

func (a *OrganizationsActivities) CreatePolicy(input *organizations.CreatePolicyInput) (*organizations.CreatePolicyOutput, error) {
	return a.client.CreatePolicy(input)
}

func (a *OrganizationsActivities) DeclineHandshake(input *organizations.DeclineHandshakeInput) (*organizations.DeclineHandshakeOutput, error) {
	return a.client.DeclineHandshake(input)
}

func (a *OrganizationsActivities) DeleteOrganization(input *organizations.DeleteOrganizationInput) (*organizations.DeleteOrganizationOutput, error) {
	return a.client.DeleteOrganization(input)
}

func (a *OrganizationsActivities) DeleteOrganizationalUnit(input *organizations.DeleteOrganizationalUnitInput) (*organizations.DeleteOrganizationalUnitOutput, error) {
	return a.client.DeleteOrganizationalUnit(input)
}

func (a *OrganizationsActivities) DeletePolicy(input *organizations.DeletePolicyInput) (*organizations.DeletePolicyOutput, error) {
	return a.client.DeletePolicy(input)
}

func (a *OrganizationsActivities) DeregisterDelegatedAdministrator(input *organizations.DeregisterDelegatedAdministratorInput) (*organizations.DeregisterDelegatedAdministratorOutput, error) {
	return a.client.DeregisterDelegatedAdministrator(input)
}

func (a *OrganizationsActivities) DescribeAccount(input *organizations.DescribeAccountInput) (*organizations.DescribeAccountOutput, error) {
	return a.client.DescribeAccount(input)
}

func (a *OrganizationsActivities) DescribeCreateAccountStatus(input *organizations.DescribeCreateAccountStatusInput) (*organizations.DescribeCreateAccountStatusOutput, error) {
	return a.client.DescribeCreateAccountStatus(input)
}

func (a *OrganizationsActivities) DescribeEffectivePolicy(input *organizations.DescribeEffectivePolicyInput) (*organizations.DescribeEffectivePolicyOutput, error) {
	return a.client.DescribeEffectivePolicy(input)
}

func (a *OrganizationsActivities) DescribeHandshake(input *organizations.DescribeHandshakeInput) (*organizations.DescribeHandshakeOutput, error) {
	return a.client.DescribeHandshake(input)
}

func (a *OrganizationsActivities) DescribeOrganization(input *organizations.DescribeOrganizationInput) (*organizations.DescribeOrganizationOutput, error) {
	return a.client.DescribeOrganization(input)
}

func (a *OrganizationsActivities) DescribeOrganizationalUnit(input *organizations.DescribeOrganizationalUnitInput) (*organizations.DescribeOrganizationalUnitOutput, error) {
	return a.client.DescribeOrganizationalUnit(input)
}

func (a *OrganizationsActivities) DescribePolicy(input *organizations.DescribePolicyInput) (*organizations.DescribePolicyOutput, error) {
	return a.client.DescribePolicy(input)
}

func (a *OrganizationsActivities) DetachPolicy(input *organizations.DetachPolicyInput) (*organizations.DetachPolicyOutput, error) {
	return a.client.DetachPolicy(input)
}

func (a *OrganizationsActivities) DisableAWSServiceAccess(input *organizations.DisableAWSServiceAccessInput) (*organizations.DisableAWSServiceAccessOutput, error) {
	return a.client.DisableAWSServiceAccess(input)
}

func (a *OrganizationsActivities) DisablePolicyType(input *organizations.DisablePolicyTypeInput) (*organizations.DisablePolicyTypeOutput, error) {
	return a.client.DisablePolicyType(input)
}

func (a *OrganizationsActivities) EnableAWSServiceAccess(input *organizations.EnableAWSServiceAccessInput) (*organizations.EnableAWSServiceAccessOutput, error) {
	return a.client.EnableAWSServiceAccess(input)
}

func (a *OrganizationsActivities) EnableAllFeatures(input *organizations.EnableAllFeaturesInput) (*organizations.EnableAllFeaturesOutput, error) {
	return a.client.EnableAllFeatures(input)
}

func (a *OrganizationsActivities) EnablePolicyType(input *organizations.EnablePolicyTypeInput) (*organizations.EnablePolicyTypeOutput, error) {
	return a.client.EnablePolicyType(input)
}

func (a *OrganizationsActivities) InviteAccountToOrganization(input *organizations.InviteAccountToOrganizationInput) (*organizations.InviteAccountToOrganizationOutput, error) {
	return a.client.InviteAccountToOrganization(input)
}

func (a *OrganizationsActivities) LeaveOrganization(input *organizations.LeaveOrganizationInput) (*organizations.LeaveOrganizationOutput, error) {
	return a.client.LeaveOrganization(input)
}

func (a *OrganizationsActivities) ListAWSServiceAccessForOrganization(input *organizations.ListAWSServiceAccessForOrganizationInput) (*organizations.ListAWSServiceAccessForOrganizationOutput, error) {
	return a.client.ListAWSServiceAccessForOrganization(input)
}

func (a *OrganizationsActivities) ListAccounts(input *organizations.ListAccountsInput) (*organizations.ListAccountsOutput, error) {
	return a.client.ListAccounts(input)
}

func (a *OrganizationsActivities) ListAccountsForParent(input *organizations.ListAccountsForParentInput) (*organizations.ListAccountsForParentOutput, error) {
	return a.client.ListAccountsForParent(input)
}

func (a *OrganizationsActivities) ListChildren(input *organizations.ListChildrenInput) (*organizations.ListChildrenOutput, error) {
	return a.client.ListChildren(input)
}

func (a *OrganizationsActivities) ListCreateAccountStatus(input *organizations.ListCreateAccountStatusInput) (*organizations.ListCreateAccountStatusOutput, error) {
	return a.client.ListCreateAccountStatus(input)
}

func (a *OrganizationsActivities) ListDelegatedAdministrators(input *organizations.ListDelegatedAdministratorsInput) (*organizations.ListDelegatedAdministratorsOutput, error) {
	return a.client.ListDelegatedAdministrators(input)
}

func (a *OrganizationsActivities) ListDelegatedServicesForAccount(input *organizations.ListDelegatedServicesForAccountInput) (*organizations.ListDelegatedServicesForAccountOutput, error) {
	return a.client.ListDelegatedServicesForAccount(input)
}

func (a *OrganizationsActivities) ListHandshakesForAccount(input *organizations.ListHandshakesForAccountInput) (*organizations.ListHandshakesForAccountOutput, error) {
	return a.client.ListHandshakesForAccount(input)
}

func (a *OrganizationsActivities) ListHandshakesForOrganization(input *organizations.ListHandshakesForOrganizationInput) (*organizations.ListHandshakesForOrganizationOutput, error) {
	return a.client.ListHandshakesForOrganization(input)
}

func (a *OrganizationsActivities) ListOrganizationalUnitsForParent(input *organizations.ListOrganizationalUnitsForParentInput) (*organizations.ListOrganizationalUnitsForParentOutput, error) {
	return a.client.ListOrganizationalUnitsForParent(input)
}

func (a *OrganizationsActivities) ListParents(input *organizations.ListParentsInput) (*organizations.ListParentsOutput, error) {
	return a.client.ListParents(input)
}

func (a *OrganizationsActivities) ListPolicies(input *organizations.ListPoliciesInput) (*organizations.ListPoliciesOutput, error) {
	return a.client.ListPolicies(input)
}

func (a *OrganizationsActivities) ListPoliciesForTarget(input *organizations.ListPoliciesForTargetInput) (*organizations.ListPoliciesForTargetOutput, error) {
	return a.client.ListPoliciesForTarget(input)
}

func (a *OrganizationsActivities) ListRoots(input *organizations.ListRootsInput) (*organizations.ListRootsOutput, error) {
	return a.client.ListRoots(input)
}

func (a *OrganizationsActivities) ListTagsForResource(input *organizations.ListTagsForResourceInput) (*organizations.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResource(input)
}

func (a *OrganizationsActivities) ListTargetsForPolicy(input *organizations.ListTargetsForPolicyInput) (*organizations.ListTargetsForPolicyOutput, error) {
	return a.client.ListTargetsForPolicy(input)
}

func (a *OrganizationsActivities) MoveAccount(input *organizations.MoveAccountInput) (*organizations.MoveAccountOutput, error) {
	return a.client.MoveAccount(input)
}

func (a *OrganizationsActivities) RegisterDelegatedAdministrator(input *organizations.RegisterDelegatedAdministratorInput) (*organizations.RegisterDelegatedAdministratorOutput, error) {
	return a.client.RegisterDelegatedAdministrator(input)
}

func (a *OrganizationsActivities) RemoveAccountFromOrganization(input *organizations.RemoveAccountFromOrganizationInput) (*organizations.RemoveAccountFromOrganizationOutput, error) {
	return a.client.RemoveAccountFromOrganization(input)
}

func (a *OrganizationsActivities) TagResource(input *organizations.TagResourceInput) (*organizations.TagResourceOutput, error) {
	return a.client.TagResource(input)
}

func (a *OrganizationsActivities) UntagResource(input *organizations.UntagResourceInput) (*organizations.UntagResourceOutput, error) {
	return a.client.UntagResource(input)
}

func (a *OrganizationsActivities) UpdateOrganizationalUnit(input *organizations.UpdateOrganizationalUnitInput) (*organizations.UpdateOrganizationalUnitOutput, error) {
	return a.client.UpdateOrganizationalUnit(input)
}

func (a *OrganizationsActivities) UpdatePolicy(input *organizations.UpdatePolicyInput) (*organizations.UpdatePolicyOutput, error) {
	return a.client.UpdatePolicy(input)
}
