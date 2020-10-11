// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package wafv2stub

import (
	"github.com/aws/aws-sdk-go/service/wafv2"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

type Client interface {
	AssociateWebACL(ctx workflow.Context, input *wafv2.AssociateWebACLInput) (*wafv2.AssociateWebACLOutput, error)
	AssociateWebACLAsync(ctx workflow.Context, input *wafv2.AssociateWebACLInput) *WAFV2AssociateWebACLFuture

	CheckCapacity(ctx workflow.Context, input *wafv2.CheckCapacityInput) (*wafv2.CheckCapacityOutput, error)
	CheckCapacityAsync(ctx workflow.Context, input *wafv2.CheckCapacityInput) *WAFV2CheckCapacityFuture

	CreateIPSet(ctx workflow.Context, input *wafv2.CreateIPSetInput) (*wafv2.CreateIPSetOutput, error)
	CreateIPSetAsync(ctx workflow.Context, input *wafv2.CreateIPSetInput) *WAFV2CreateIPSetFuture

	CreateRegexPatternSet(ctx workflow.Context, input *wafv2.CreateRegexPatternSetInput) (*wafv2.CreateRegexPatternSetOutput, error)
	CreateRegexPatternSetAsync(ctx workflow.Context, input *wafv2.CreateRegexPatternSetInput) *WAFV2CreateRegexPatternSetFuture

	CreateRuleGroup(ctx workflow.Context, input *wafv2.CreateRuleGroupInput) (*wafv2.CreateRuleGroupOutput, error)
	CreateRuleGroupAsync(ctx workflow.Context, input *wafv2.CreateRuleGroupInput) *WAFV2CreateRuleGroupFuture

	CreateWebACL(ctx workflow.Context, input *wafv2.CreateWebACLInput) (*wafv2.CreateWebACLOutput, error)
	CreateWebACLAsync(ctx workflow.Context, input *wafv2.CreateWebACLInput) *WAFV2CreateWebACLFuture

	DeleteFirewallManagerRuleGroups(ctx workflow.Context, input *wafv2.DeleteFirewallManagerRuleGroupsInput) (*wafv2.DeleteFirewallManagerRuleGroupsOutput, error)
	DeleteFirewallManagerRuleGroupsAsync(ctx workflow.Context, input *wafv2.DeleteFirewallManagerRuleGroupsInput) *WAFV2DeleteFirewallManagerRuleGroupsFuture

	DeleteIPSet(ctx workflow.Context, input *wafv2.DeleteIPSetInput) (*wafv2.DeleteIPSetOutput, error)
	DeleteIPSetAsync(ctx workflow.Context, input *wafv2.DeleteIPSetInput) *WAFV2DeleteIPSetFuture

	DeleteLoggingConfiguration(ctx workflow.Context, input *wafv2.DeleteLoggingConfigurationInput) (*wafv2.DeleteLoggingConfigurationOutput, error)
	DeleteLoggingConfigurationAsync(ctx workflow.Context, input *wafv2.DeleteLoggingConfigurationInput) *WAFV2DeleteLoggingConfigurationFuture

	DeletePermissionPolicy(ctx workflow.Context, input *wafv2.DeletePermissionPolicyInput) (*wafv2.DeletePermissionPolicyOutput, error)
	DeletePermissionPolicyAsync(ctx workflow.Context, input *wafv2.DeletePermissionPolicyInput) *WAFV2DeletePermissionPolicyFuture

	DeleteRegexPatternSet(ctx workflow.Context, input *wafv2.DeleteRegexPatternSetInput) (*wafv2.DeleteRegexPatternSetOutput, error)
	DeleteRegexPatternSetAsync(ctx workflow.Context, input *wafv2.DeleteRegexPatternSetInput) *WAFV2DeleteRegexPatternSetFuture

	DeleteRuleGroup(ctx workflow.Context, input *wafv2.DeleteRuleGroupInput) (*wafv2.DeleteRuleGroupOutput, error)
	DeleteRuleGroupAsync(ctx workflow.Context, input *wafv2.DeleteRuleGroupInput) *WAFV2DeleteRuleGroupFuture

	DeleteWebACL(ctx workflow.Context, input *wafv2.DeleteWebACLInput) (*wafv2.DeleteWebACLOutput, error)
	DeleteWebACLAsync(ctx workflow.Context, input *wafv2.DeleteWebACLInput) *WAFV2DeleteWebACLFuture

	DescribeManagedRuleGroup(ctx workflow.Context, input *wafv2.DescribeManagedRuleGroupInput) (*wafv2.DescribeManagedRuleGroupOutput, error)
	DescribeManagedRuleGroupAsync(ctx workflow.Context, input *wafv2.DescribeManagedRuleGroupInput) *WAFV2DescribeManagedRuleGroupFuture

	DisassociateWebACL(ctx workflow.Context, input *wafv2.DisassociateWebACLInput) (*wafv2.DisassociateWebACLOutput, error)
	DisassociateWebACLAsync(ctx workflow.Context, input *wafv2.DisassociateWebACLInput) *WAFV2DisassociateWebACLFuture

	GetIPSet(ctx workflow.Context, input *wafv2.GetIPSetInput) (*wafv2.GetIPSetOutput, error)
	GetIPSetAsync(ctx workflow.Context, input *wafv2.GetIPSetInput) *WAFV2GetIPSetFuture

	GetLoggingConfiguration(ctx workflow.Context, input *wafv2.GetLoggingConfigurationInput) (*wafv2.GetLoggingConfigurationOutput, error)
	GetLoggingConfigurationAsync(ctx workflow.Context, input *wafv2.GetLoggingConfigurationInput) *WAFV2GetLoggingConfigurationFuture

	GetPermissionPolicy(ctx workflow.Context, input *wafv2.GetPermissionPolicyInput) (*wafv2.GetPermissionPolicyOutput, error)
	GetPermissionPolicyAsync(ctx workflow.Context, input *wafv2.GetPermissionPolicyInput) *WAFV2GetPermissionPolicyFuture

	GetRateBasedStatementManagedKeys(ctx workflow.Context, input *wafv2.GetRateBasedStatementManagedKeysInput) (*wafv2.GetRateBasedStatementManagedKeysOutput, error)
	GetRateBasedStatementManagedKeysAsync(ctx workflow.Context, input *wafv2.GetRateBasedStatementManagedKeysInput) *WAFV2GetRateBasedStatementManagedKeysFuture

	GetRegexPatternSet(ctx workflow.Context, input *wafv2.GetRegexPatternSetInput) (*wafv2.GetRegexPatternSetOutput, error)
	GetRegexPatternSetAsync(ctx workflow.Context, input *wafv2.GetRegexPatternSetInput) *WAFV2GetRegexPatternSetFuture

	GetRuleGroup(ctx workflow.Context, input *wafv2.GetRuleGroupInput) (*wafv2.GetRuleGroupOutput, error)
	GetRuleGroupAsync(ctx workflow.Context, input *wafv2.GetRuleGroupInput) *WAFV2GetRuleGroupFuture

	GetSampledRequests(ctx workflow.Context, input *wafv2.GetSampledRequestsInput) (*wafv2.GetSampledRequestsOutput, error)
	GetSampledRequestsAsync(ctx workflow.Context, input *wafv2.GetSampledRequestsInput) *WAFV2GetSampledRequestsFuture

	GetWebACL(ctx workflow.Context, input *wafv2.GetWebACLInput) (*wafv2.GetWebACLOutput, error)
	GetWebACLAsync(ctx workflow.Context, input *wafv2.GetWebACLInput) *WAFV2GetWebACLFuture

	GetWebACLForResource(ctx workflow.Context, input *wafv2.GetWebACLForResourceInput) (*wafv2.GetWebACLForResourceOutput, error)
	GetWebACLForResourceAsync(ctx workflow.Context, input *wafv2.GetWebACLForResourceInput) *WAFV2GetWebACLForResourceFuture

	ListAvailableManagedRuleGroups(ctx workflow.Context, input *wafv2.ListAvailableManagedRuleGroupsInput) (*wafv2.ListAvailableManagedRuleGroupsOutput, error)
	ListAvailableManagedRuleGroupsAsync(ctx workflow.Context, input *wafv2.ListAvailableManagedRuleGroupsInput) *WAFV2ListAvailableManagedRuleGroupsFuture

	ListIPSets(ctx workflow.Context, input *wafv2.ListIPSetsInput) (*wafv2.ListIPSetsOutput, error)
	ListIPSetsAsync(ctx workflow.Context, input *wafv2.ListIPSetsInput) *WAFV2ListIPSetsFuture

	ListLoggingConfigurations(ctx workflow.Context, input *wafv2.ListLoggingConfigurationsInput) (*wafv2.ListLoggingConfigurationsOutput, error)
	ListLoggingConfigurationsAsync(ctx workflow.Context, input *wafv2.ListLoggingConfigurationsInput) *WAFV2ListLoggingConfigurationsFuture

	ListRegexPatternSets(ctx workflow.Context, input *wafv2.ListRegexPatternSetsInput) (*wafv2.ListRegexPatternSetsOutput, error)
	ListRegexPatternSetsAsync(ctx workflow.Context, input *wafv2.ListRegexPatternSetsInput) *WAFV2ListRegexPatternSetsFuture

	ListResourcesForWebACL(ctx workflow.Context, input *wafv2.ListResourcesForWebACLInput) (*wafv2.ListResourcesForWebACLOutput, error)
	ListResourcesForWebACLAsync(ctx workflow.Context, input *wafv2.ListResourcesForWebACLInput) *WAFV2ListResourcesForWebACLFuture

	ListRuleGroups(ctx workflow.Context, input *wafv2.ListRuleGroupsInput) (*wafv2.ListRuleGroupsOutput, error)
	ListRuleGroupsAsync(ctx workflow.Context, input *wafv2.ListRuleGroupsInput) *WAFV2ListRuleGroupsFuture

	ListTagsForResource(ctx workflow.Context, input *wafv2.ListTagsForResourceInput) (*wafv2.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *wafv2.ListTagsForResourceInput) *WAFV2ListTagsForResourceFuture

	ListWebACLs(ctx workflow.Context, input *wafv2.ListWebACLsInput) (*wafv2.ListWebACLsOutput, error)
	ListWebACLsAsync(ctx workflow.Context, input *wafv2.ListWebACLsInput) *WAFV2ListWebACLsFuture

	PutLoggingConfiguration(ctx workflow.Context, input *wafv2.PutLoggingConfigurationInput) (*wafv2.PutLoggingConfigurationOutput, error)
	PutLoggingConfigurationAsync(ctx workflow.Context, input *wafv2.PutLoggingConfigurationInput) *WAFV2PutLoggingConfigurationFuture

	PutPermissionPolicy(ctx workflow.Context, input *wafv2.PutPermissionPolicyInput) (*wafv2.PutPermissionPolicyOutput, error)
	PutPermissionPolicyAsync(ctx workflow.Context, input *wafv2.PutPermissionPolicyInput) *WAFV2PutPermissionPolicyFuture

	TagResource(ctx workflow.Context, input *wafv2.TagResourceInput) (*wafv2.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *wafv2.TagResourceInput) *WAFV2TagResourceFuture

	UntagResource(ctx workflow.Context, input *wafv2.UntagResourceInput) (*wafv2.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *wafv2.UntagResourceInput) *WAFV2UntagResourceFuture

	UpdateIPSet(ctx workflow.Context, input *wafv2.UpdateIPSetInput) (*wafv2.UpdateIPSetOutput, error)
	UpdateIPSetAsync(ctx workflow.Context, input *wafv2.UpdateIPSetInput) *WAFV2UpdateIPSetFuture

	UpdateRegexPatternSet(ctx workflow.Context, input *wafv2.UpdateRegexPatternSetInput) (*wafv2.UpdateRegexPatternSetOutput, error)
	UpdateRegexPatternSetAsync(ctx workflow.Context, input *wafv2.UpdateRegexPatternSetInput) *WAFV2UpdateRegexPatternSetFuture

	UpdateRuleGroup(ctx workflow.Context, input *wafv2.UpdateRuleGroupInput) (*wafv2.UpdateRuleGroupOutput, error)
	UpdateRuleGroupAsync(ctx workflow.Context, input *wafv2.UpdateRuleGroupInput) *WAFV2UpdateRuleGroupFuture

	UpdateWebACL(ctx workflow.Context, input *wafv2.UpdateWebACLInput) (*wafv2.UpdateWebACLOutput, error)
	UpdateWebACLAsync(ctx workflow.Context, input *wafv2.UpdateWebACLInput) *WAFV2UpdateWebACLFuture
}

func NewClient() Client {
	return &stub{}
}

