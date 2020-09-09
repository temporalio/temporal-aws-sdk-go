
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/workmail"
	"github.com/aws/aws-sdk-go/service/workmail/workmailiface"
)

type WorkMailActivities struct {
	client workmailiface.WorkMailAPI
}

func NewWorkMailActivities(session *session.Session, config... *aws.Config) *WorkMailActivities {
    client := workmail.New(session, config...)
    return &WorkMailActivities{client: client}
}

func (a *WorkMailActivities) AssociateDelegateToResource(input *workmail.AssociateDelegateToResourceInput) (*workmail.AssociateDelegateToResourceOutput, error) {
    return a.client.AssociateDelegateToResource(input)
}

func (a *WorkMailActivities) AssociateMemberToGroup(input *workmail.AssociateMemberToGroupInput) (*workmail.AssociateMemberToGroupOutput, error) {
    return a.client.AssociateMemberToGroup(input)
}

func (a *WorkMailActivities) CreateAlias(input *workmail.CreateAliasInput) (*workmail.CreateAliasOutput, error) {
    return a.client.CreateAlias(input)
}

func (a *WorkMailActivities) CreateGroup(input *workmail.CreateGroupInput) (*workmail.CreateGroupOutput, error) {
    return a.client.CreateGroup(input)
}

func (a *WorkMailActivities) CreateResource(input *workmail.CreateResourceInput) (*workmail.CreateResourceOutput, error) {
    return a.client.CreateResource(input)
}

func (a *WorkMailActivities) CreateUser(input *workmail.CreateUserInput) (*workmail.CreateUserOutput, error) {
    return a.client.CreateUser(input)
}

func (a *WorkMailActivities) DeleteAccessControlRule(input *workmail.DeleteAccessControlRuleInput) (*workmail.DeleteAccessControlRuleOutput, error) {
    return a.client.DeleteAccessControlRule(input)
}

func (a *WorkMailActivities) DeleteAlias(input *workmail.DeleteAliasInput) (*workmail.DeleteAliasOutput, error) {
    return a.client.DeleteAlias(input)
}

func (a *WorkMailActivities) DeleteGroup(input *workmail.DeleteGroupInput) (*workmail.DeleteGroupOutput, error) {
    return a.client.DeleteGroup(input)
}

func (a *WorkMailActivities) DeleteMailboxPermissions(input *workmail.DeleteMailboxPermissionsInput) (*workmail.DeleteMailboxPermissionsOutput, error) {
    return a.client.DeleteMailboxPermissions(input)
}

func (a *WorkMailActivities) DeleteResource(input *workmail.DeleteResourceInput) (*workmail.DeleteResourceOutput, error) {
    return a.client.DeleteResource(input)
}

func (a *WorkMailActivities) DeleteRetentionPolicy(input *workmail.DeleteRetentionPolicyInput) (*workmail.DeleteRetentionPolicyOutput, error) {
    return a.client.DeleteRetentionPolicy(input)
}

func (a *WorkMailActivities) DeleteUser(input *workmail.DeleteUserInput) (*workmail.DeleteUserOutput, error) {
    return a.client.DeleteUser(input)
}

func (a *WorkMailActivities) DeregisterFromWorkMail(input *workmail.DeregisterFromWorkMailInput) (*workmail.DeregisterFromWorkMailOutput, error) {
    return a.client.DeregisterFromWorkMail(input)
}

func (a *WorkMailActivities) DescribeGroup(input *workmail.DescribeGroupInput) (*workmail.DescribeGroupOutput, error) {
    return a.client.DescribeGroup(input)
}

func (a *WorkMailActivities) DescribeOrganization(input *workmail.DescribeOrganizationInput) (*workmail.DescribeOrganizationOutput, error) {
    return a.client.DescribeOrganization(input)
}

func (a *WorkMailActivities) DescribeResource(input *workmail.DescribeResourceInput) (*workmail.DescribeResourceOutput, error) {
    return a.client.DescribeResource(input)
}

func (a *WorkMailActivities) DescribeUser(input *workmail.DescribeUserInput) (*workmail.DescribeUserOutput, error) {
    return a.client.DescribeUser(input)
}

func (a *WorkMailActivities) DisassociateDelegateFromResource(input *workmail.DisassociateDelegateFromResourceInput) (*workmail.DisassociateDelegateFromResourceOutput, error) {
    return a.client.DisassociateDelegateFromResource(input)
}

func (a *WorkMailActivities) DisassociateMemberFromGroup(input *workmail.DisassociateMemberFromGroupInput) (*workmail.DisassociateMemberFromGroupOutput, error) {
    return a.client.DisassociateMemberFromGroup(input)
}

func (a *WorkMailActivities) GetAccessControlEffect(input *workmail.GetAccessControlEffectInput) (*workmail.GetAccessControlEffectOutput, error) {
    return a.client.GetAccessControlEffect(input)
}

func (a *WorkMailActivities) GetDefaultRetentionPolicy(input *workmail.GetDefaultRetentionPolicyInput) (*workmail.GetDefaultRetentionPolicyOutput, error) {
    return a.client.GetDefaultRetentionPolicy(input)
}

func (a *WorkMailActivities) GetMailboxDetails(input *workmail.GetMailboxDetailsInput) (*workmail.GetMailboxDetailsOutput, error) {
    return a.client.GetMailboxDetails(input)
}

func (a *WorkMailActivities) ListAccessControlRules(input *workmail.ListAccessControlRulesInput) (*workmail.ListAccessControlRulesOutput, error) {
    return a.client.ListAccessControlRules(input)
}

func (a *WorkMailActivities) ListAliases(input *workmail.ListAliasesInput) (*workmail.ListAliasesOutput, error) {
    return a.client.ListAliases(input)
}

func (a *WorkMailActivities) ListGroupMembers(input *workmail.ListGroupMembersInput) (*workmail.ListGroupMembersOutput, error) {
    return a.client.ListGroupMembers(input)
}

func (a *WorkMailActivities) ListGroups(input *workmail.ListGroupsInput) (*workmail.ListGroupsOutput, error) {
    return a.client.ListGroups(input)
}

func (a *WorkMailActivities) ListMailboxPermissions(input *workmail.ListMailboxPermissionsInput) (*workmail.ListMailboxPermissionsOutput, error) {
    return a.client.ListMailboxPermissions(input)
}

func (a *WorkMailActivities) ListOrganizations(input *workmail.ListOrganizationsInput) (*workmail.ListOrganizationsOutput, error) {
    return a.client.ListOrganizations(input)
}

func (a *WorkMailActivities) ListResourceDelegates(input *workmail.ListResourceDelegatesInput) (*workmail.ListResourceDelegatesOutput, error) {
    return a.client.ListResourceDelegates(input)
}

func (a *WorkMailActivities) ListResources(input *workmail.ListResourcesInput) (*workmail.ListResourcesOutput, error) {
    return a.client.ListResources(input)
}

func (a *WorkMailActivities) ListTagsForResource(input *workmail.ListTagsForResourceInput) (*workmail.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *WorkMailActivities) ListUsers(input *workmail.ListUsersInput) (*workmail.ListUsersOutput, error) {
    return a.client.ListUsers(input)
}

func (a *WorkMailActivities) PutAccessControlRule(input *workmail.PutAccessControlRuleInput) (*workmail.PutAccessControlRuleOutput, error) {
    return a.client.PutAccessControlRule(input)
}

func (a *WorkMailActivities) PutMailboxPermissions(input *workmail.PutMailboxPermissionsInput) (*workmail.PutMailboxPermissionsOutput, error) {
    return a.client.PutMailboxPermissions(input)
}

func (a *WorkMailActivities) PutRetentionPolicy(input *workmail.PutRetentionPolicyInput) (*workmail.PutRetentionPolicyOutput, error) {
    return a.client.PutRetentionPolicy(input)
}

func (a *WorkMailActivities) RegisterToWorkMail(input *workmail.RegisterToWorkMailInput) (*workmail.RegisterToWorkMailOutput, error) {
    return a.client.RegisterToWorkMail(input)
}

func (a *WorkMailActivities) ResetPassword(input *workmail.ResetPasswordInput) (*workmail.ResetPasswordOutput, error) {
    return a.client.ResetPassword(input)
}

func (a *WorkMailActivities) TagResource(input *workmail.TagResourceInput) (*workmail.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *WorkMailActivities) UntagResource(input *workmail.UntagResourceInput) (*workmail.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *WorkMailActivities) UpdateMailboxQuota(input *workmail.UpdateMailboxQuotaInput) (*workmail.UpdateMailboxQuotaOutput, error) {
    return a.client.UpdateMailboxQuota(input)
}

func (a *WorkMailActivities) UpdatePrimaryEmailAddress(input *workmail.UpdatePrimaryEmailAddressInput) (*workmail.UpdatePrimaryEmailAddressOutput, error) {
    return a.client.UpdatePrimaryEmailAddress(input)
}

func (a *WorkMailActivities) UpdateResource(input *workmail.UpdateResourceInput) (*workmail.UpdateResourceOutput, error) {
    return a.client.UpdateResource(input)
}
