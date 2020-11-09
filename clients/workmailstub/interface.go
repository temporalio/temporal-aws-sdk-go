// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package workmailstub

import (
	"github.com/aws/aws-sdk-go/service/workmail"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AssociateDelegateToResource(ctx workflow.Context, input *workmail.AssociateDelegateToResourceInput) (*workmail.AssociateDelegateToResourceOutput, error)
	AssociateDelegateToResourceAsync(ctx workflow.Context, input *workmail.AssociateDelegateToResourceInput) *AssociateDelegateToResourceFuture

	AssociateMemberToGroup(ctx workflow.Context, input *workmail.AssociateMemberToGroupInput) (*workmail.AssociateMemberToGroupOutput, error)
	AssociateMemberToGroupAsync(ctx workflow.Context, input *workmail.AssociateMemberToGroupInput) *AssociateMemberToGroupFuture

	CancelMailboxExportJob(ctx workflow.Context, input *workmail.CancelMailboxExportJobInput) (*workmail.CancelMailboxExportJobOutput, error)
	CancelMailboxExportJobAsync(ctx workflow.Context, input *workmail.CancelMailboxExportJobInput) *CancelMailboxExportJobFuture

	CreateAlias(ctx workflow.Context, input *workmail.CreateAliasInput) (*workmail.CreateAliasOutput, error)
	CreateAliasAsync(ctx workflow.Context, input *workmail.CreateAliasInput) *CreateAliasFuture

	CreateGroup(ctx workflow.Context, input *workmail.CreateGroupInput) (*workmail.CreateGroupOutput, error)
	CreateGroupAsync(ctx workflow.Context, input *workmail.CreateGroupInput) *CreateGroupFuture

	CreateResource(ctx workflow.Context, input *workmail.CreateResourceInput) (*workmail.CreateResourceOutput, error)
	CreateResourceAsync(ctx workflow.Context, input *workmail.CreateResourceInput) *CreateResourceFuture

	CreateUser(ctx workflow.Context, input *workmail.CreateUserInput) (*workmail.CreateUserOutput, error)
	CreateUserAsync(ctx workflow.Context, input *workmail.CreateUserInput) *CreateUserFuture

	DeleteAccessControlRule(ctx workflow.Context, input *workmail.DeleteAccessControlRuleInput) (*workmail.DeleteAccessControlRuleOutput, error)
	DeleteAccessControlRuleAsync(ctx workflow.Context, input *workmail.DeleteAccessControlRuleInput) *DeleteAccessControlRuleFuture

	DeleteAlias(ctx workflow.Context, input *workmail.DeleteAliasInput) (*workmail.DeleteAliasOutput, error)
	DeleteAliasAsync(ctx workflow.Context, input *workmail.DeleteAliasInput) *DeleteAliasFuture

	DeleteGroup(ctx workflow.Context, input *workmail.DeleteGroupInput) (*workmail.DeleteGroupOutput, error)
	DeleteGroupAsync(ctx workflow.Context, input *workmail.DeleteGroupInput) *DeleteGroupFuture

	DeleteMailboxPermissions(ctx workflow.Context, input *workmail.DeleteMailboxPermissionsInput) (*workmail.DeleteMailboxPermissionsOutput, error)
	DeleteMailboxPermissionsAsync(ctx workflow.Context, input *workmail.DeleteMailboxPermissionsInput) *DeleteMailboxPermissionsFuture

	DeleteResource(ctx workflow.Context, input *workmail.DeleteResourceInput) (*workmail.DeleteResourceOutput, error)
	DeleteResourceAsync(ctx workflow.Context, input *workmail.DeleteResourceInput) *DeleteResourceFuture

	DeleteRetentionPolicy(ctx workflow.Context, input *workmail.DeleteRetentionPolicyInput) (*workmail.DeleteRetentionPolicyOutput, error)
	DeleteRetentionPolicyAsync(ctx workflow.Context, input *workmail.DeleteRetentionPolicyInput) *DeleteRetentionPolicyFuture

	DeleteUser(ctx workflow.Context, input *workmail.DeleteUserInput) (*workmail.DeleteUserOutput, error)
	DeleteUserAsync(ctx workflow.Context, input *workmail.DeleteUserInput) *DeleteUserFuture

	DeregisterFromWorkMail(ctx workflow.Context, input *workmail.DeregisterFromWorkMailInput) (*workmail.DeregisterFromWorkMailOutput, error)
	DeregisterFromWorkMailAsync(ctx workflow.Context, input *workmail.DeregisterFromWorkMailInput) *DeregisterFromWorkMailFuture

	DescribeGroup(ctx workflow.Context, input *workmail.DescribeGroupInput) (*workmail.DescribeGroupOutput, error)
	DescribeGroupAsync(ctx workflow.Context, input *workmail.DescribeGroupInput) *DescribeGroupFuture

	DescribeMailboxExportJob(ctx workflow.Context, input *workmail.DescribeMailboxExportJobInput) (*workmail.DescribeMailboxExportJobOutput, error)
	DescribeMailboxExportJobAsync(ctx workflow.Context, input *workmail.DescribeMailboxExportJobInput) *DescribeMailboxExportJobFuture

	DescribeOrganization(ctx workflow.Context, input *workmail.DescribeOrganizationInput) (*workmail.DescribeOrganizationOutput, error)
	DescribeOrganizationAsync(ctx workflow.Context, input *workmail.DescribeOrganizationInput) *DescribeOrganizationFuture

	DescribeResource(ctx workflow.Context, input *workmail.DescribeResourceInput) (*workmail.DescribeResourceOutput, error)
	DescribeResourceAsync(ctx workflow.Context, input *workmail.DescribeResourceInput) *DescribeResourceFuture

	DescribeUser(ctx workflow.Context, input *workmail.DescribeUserInput) (*workmail.DescribeUserOutput, error)
	DescribeUserAsync(ctx workflow.Context, input *workmail.DescribeUserInput) *DescribeUserFuture

	DisassociateDelegateFromResource(ctx workflow.Context, input *workmail.DisassociateDelegateFromResourceInput) (*workmail.DisassociateDelegateFromResourceOutput, error)
	DisassociateDelegateFromResourceAsync(ctx workflow.Context, input *workmail.DisassociateDelegateFromResourceInput) *DisassociateDelegateFromResourceFuture

	DisassociateMemberFromGroup(ctx workflow.Context, input *workmail.DisassociateMemberFromGroupInput) (*workmail.DisassociateMemberFromGroupOutput, error)
	DisassociateMemberFromGroupAsync(ctx workflow.Context, input *workmail.DisassociateMemberFromGroupInput) *DisassociateMemberFromGroupFuture

	GetAccessControlEffect(ctx workflow.Context, input *workmail.GetAccessControlEffectInput) (*workmail.GetAccessControlEffectOutput, error)
	GetAccessControlEffectAsync(ctx workflow.Context, input *workmail.GetAccessControlEffectInput) *GetAccessControlEffectFuture

	GetDefaultRetentionPolicy(ctx workflow.Context, input *workmail.GetDefaultRetentionPolicyInput) (*workmail.GetDefaultRetentionPolicyOutput, error)
	GetDefaultRetentionPolicyAsync(ctx workflow.Context, input *workmail.GetDefaultRetentionPolicyInput) *GetDefaultRetentionPolicyFuture

	GetMailboxDetails(ctx workflow.Context, input *workmail.GetMailboxDetailsInput) (*workmail.GetMailboxDetailsOutput, error)
	GetMailboxDetailsAsync(ctx workflow.Context, input *workmail.GetMailboxDetailsInput) *GetMailboxDetailsFuture

	ListAccessControlRules(ctx workflow.Context, input *workmail.ListAccessControlRulesInput) (*workmail.ListAccessControlRulesOutput, error)
	ListAccessControlRulesAsync(ctx workflow.Context, input *workmail.ListAccessControlRulesInput) *ListAccessControlRulesFuture

	ListAliases(ctx workflow.Context, input *workmail.ListAliasesInput) (*workmail.ListAliasesOutput, error)
	ListAliasesAsync(ctx workflow.Context, input *workmail.ListAliasesInput) *ListAliasesFuture

	ListGroupMembers(ctx workflow.Context, input *workmail.ListGroupMembersInput) (*workmail.ListGroupMembersOutput, error)
	ListGroupMembersAsync(ctx workflow.Context, input *workmail.ListGroupMembersInput) *ListGroupMembersFuture

	ListGroups(ctx workflow.Context, input *workmail.ListGroupsInput) (*workmail.ListGroupsOutput, error)
	ListGroupsAsync(ctx workflow.Context, input *workmail.ListGroupsInput) *ListGroupsFuture

	ListMailboxExportJobs(ctx workflow.Context, input *workmail.ListMailboxExportJobsInput) (*workmail.ListMailboxExportJobsOutput, error)
	ListMailboxExportJobsAsync(ctx workflow.Context, input *workmail.ListMailboxExportJobsInput) *ListMailboxExportJobsFuture

	ListMailboxPermissions(ctx workflow.Context, input *workmail.ListMailboxPermissionsInput) (*workmail.ListMailboxPermissionsOutput, error)
	ListMailboxPermissionsAsync(ctx workflow.Context, input *workmail.ListMailboxPermissionsInput) *ListMailboxPermissionsFuture

	ListOrganizations(ctx workflow.Context, input *workmail.ListOrganizationsInput) (*workmail.ListOrganizationsOutput, error)
	ListOrganizationsAsync(ctx workflow.Context, input *workmail.ListOrganizationsInput) *ListOrganizationsFuture

	ListResourceDelegates(ctx workflow.Context, input *workmail.ListResourceDelegatesInput) (*workmail.ListResourceDelegatesOutput, error)
	ListResourceDelegatesAsync(ctx workflow.Context, input *workmail.ListResourceDelegatesInput) *ListResourceDelegatesFuture

	ListResources(ctx workflow.Context, input *workmail.ListResourcesInput) (*workmail.ListResourcesOutput, error)
	ListResourcesAsync(ctx workflow.Context, input *workmail.ListResourcesInput) *ListResourcesFuture

	ListTagsForResource(ctx workflow.Context, input *workmail.ListTagsForResourceInput) (*workmail.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *workmail.ListTagsForResourceInput) *ListTagsForResourceFuture

	ListUsers(ctx workflow.Context, input *workmail.ListUsersInput) (*workmail.ListUsersOutput, error)
	ListUsersAsync(ctx workflow.Context, input *workmail.ListUsersInput) *ListUsersFuture

	PutAccessControlRule(ctx workflow.Context, input *workmail.PutAccessControlRuleInput) (*workmail.PutAccessControlRuleOutput, error)
	PutAccessControlRuleAsync(ctx workflow.Context, input *workmail.PutAccessControlRuleInput) *PutAccessControlRuleFuture

	PutMailboxPermissions(ctx workflow.Context, input *workmail.PutMailboxPermissionsInput) (*workmail.PutMailboxPermissionsOutput, error)
	PutMailboxPermissionsAsync(ctx workflow.Context, input *workmail.PutMailboxPermissionsInput) *PutMailboxPermissionsFuture

	PutRetentionPolicy(ctx workflow.Context, input *workmail.PutRetentionPolicyInput) (*workmail.PutRetentionPolicyOutput, error)
	PutRetentionPolicyAsync(ctx workflow.Context, input *workmail.PutRetentionPolicyInput) *PutRetentionPolicyFuture

	RegisterToWorkMail(ctx workflow.Context, input *workmail.RegisterToWorkMailInput) (*workmail.RegisterToWorkMailOutput, error)
	RegisterToWorkMailAsync(ctx workflow.Context, input *workmail.RegisterToWorkMailInput) *RegisterToWorkMailFuture

	ResetPassword(ctx workflow.Context, input *workmail.ResetPasswordInput) (*workmail.ResetPasswordOutput, error)
	ResetPasswordAsync(ctx workflow.Context, input *workmail.ResetPasswordInput) *ResetPasswordFuture

	StartMailboxExportJob(ctx workflow.Context, input *workmail.StartMailboxExportJobInput) (*workmail.StartMailboxExportJobOutput, error)
	StartMailboxExportJobAsync(ctx workflow.Context, input *workmail.StartMailboxExportJobInput) *StartMailboxExportJobFuture

	TagResource(ctx workflow.Context, input *workmail.TagResourceInput) (*workmail.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *workmail.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *workmail.UntagResourceInput) (*workmail.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *workmail.UntagResourceInput) *UntagResourceFuture

	UpdateMailboxQuota(ctx workflow.Context, input *workmail.UpdateMailboxQuotaInput) (*workmail.UpdateMailboxQuotaOutput, error)
	UpdateMailboxQuotaAsync(ctx workflow.Context, input *workmail.UpdateMailboxQuotaInput) *UpdateMailboxQuotaFuture

	UpdatePrimaryEmailAddress(ctx workflow.Context, input *workmail.UpdatePrimaryEmailAddressInput) (*workmail.UpdatePrimaryEmailAddressOutput, error)
	UpdatePrimaryEmailAddressAsync(ctx workflow.Context, input *workmail.UpdatePrimaryEmailAddressInput) *UpdatePrimaryEmailAddressFuture

	UpdateResource(ctx workflow.Context, input *workmail.UpdateResourceInput) (*workmail.UpdateResourceOutput, error)
	UpdateResourceAsync(ctx workflow.Context, input *workmail.UpdateResourceInput) *UpdateResourceFuture
}

func NewClient() Client {
	return &stub{}
}
