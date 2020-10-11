// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package wafv2stub

import (
	"github.com/aws/aws-sdk-go/service/wafv2"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type WAFV2AssociateWebACLFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2AssociateWebACLFuture) Get(ctx workflow.Context) (*wafv2.AssociateWebACLOutput, error) {
	var output wafv2.AssociateWebACLOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2CheckCapacityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2CheckCapacityFuture) Get(ctx workflow.Context) (*wafv2.CheckCapacityOutput, error) {
	var output wafv2.CheckCapacityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2CreateIPSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2CreateIPSetFuture) Get(ctx workflow.Context) (*wafv2.CreateIPSetOutput, error) {
	var output wafv2.CreateIPSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2CreateRegexPatternSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2CreateRegexPatternSetFuture) Get(ctx workflow.Context) (*wafv2.CreateRegexPatternSetOutput, error) {
	var output wafv2.CreateRegexPatternSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2CreateRuleGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2CreateRuleGroupFuture) Get(ctx workflow.Context) (*wafv2.CreateRuleGroupOutput, error) {
	var output wafv2.CreateRuleGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2CreateWebACLFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2CreateWebACLFuture) Get(ctx workflow.Context) (*wafv2.CreateWebACLOutput, error) {
	var output wafv2.CreateWebACLOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2DeleteFirewallManagerRuleGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2DeleteFirewallManagerRuleGroupsFuture) Get(ctx workflow.Context) (*wafv2.DeleteFirewallManagerRuleGroupsOutput, error) {
	var output wafv2.DeleteFirewallManagerRuleGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2DeleteIPSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2DeleteIPSetFuture) Get(ctx workflow.Context) (*wafv2.DeleteIPSetOutput, error) {
	var output wafv2.DeleteIPSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2DeleteLoggingConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2DeleteLoggingConfigurationFuture) Get(ctx workflow.Context) (*wafv2.DeleteLoggingConfigurationOutput, error) {
	var output wafv2.DeleteLoggingConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2DeletePermissionPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2DeletePermissionPolicyFuture) Get(ctx workflow.Context) (*wafv2.DeletePermissionPolicyOutput, error) {
	var output wafv2.DeletePermissionPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2DeleteRegexPatternSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2DeleteRegexPatternSetFuture) Get(ctx workflow.Context) (*wafv2.DeleteRegexPatternSetOutput, error) {
	var output wafv2.DeleteRegexPatternSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2DeleteRuleGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2DeleteRuleGroupFuture) Get(ctx workflow.Context) (*wafv2.DeleteRuleGroupOutput, error) {
	var output wafv2.DeleteRuleGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2DeleteWebACLFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2DeleteWebACLFuture) Get(ctx workflow.Context) (*wafv2.DeleteWebACLOutput, error) {
	var output wafv2.DeleteWebACLOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2DescribeManagedRuleGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2DescribeManagedRuleGroupFuture) Get(ctx workflow.Context) (*wafv2.DescribeManagedRuleGroupOutput, error) {
	var output wafv2.DescribeManagedRuleGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2DisassociateWebACLFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2DisassociateWebACLFuture) Get(ctx workflow.Context) (*wafv2.DisassociateWebACLOutput, error) {
	var output wafv2.DisassociateWebACLOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2GetIPSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2GetIPSetFuture) Get(ctx workflow.Context) (*wafv2.GetIPSetOutput, error) {
	var output wafv2.GetIPSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2GetLoggingConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2GetLoggingConfigurationFuture) Get(ctx workflow.Context) (*wafv2.GetLoggingConfigurationOutput, error) {
	var output wafv2.GetLoggingConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2GetPermissionPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2GetPermissionPolicyFuture) Get(ctx workflow.Context) (*wafv2.GetPermissionPolicyOutput, error) {
	var output wafv2.GetPermissionPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2GetRateBasedStatementManagedKeysFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2GetRateBasedStatementManagedKeysFuture) Get(ctx workflow.Context) (*wafv2.GetRateBasedStatementManagedKeysOutput, error) {
	var output wafv2.GetRateBasedStatementManagedKeysOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2GetRegexPatternSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2GetRegexPatternSetFuture) Get(ctx workflow.Context) (*wafv2.GetRegexPatternSetOutput, error) {
	var output wafv2.GetRegexPatternSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2GetRuleGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2GetRuleGroupFuture) Get(ctx workflow.Context) (*wafv2.GetRuleGroupOutput, error) {
	var output wafv2.GetRuleGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2GetSampledRequestsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2GetSampledRequestsFuture) Get(ctx workflow.Context) (*wafv2.GetSampledRequestsOutput, error) {
	var output wafv2.GetSampledRequestsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2GetWebACLFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2GetWebACLFuture) Get(ctx workflow.Context) (*wafv2.GetWebACLOutput, error) {
	var output wafv2.GetWebACLOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2GetWebACLForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2GetWebACLForResourceFuture) Get(ctx workflow.Context) (*wafv2.GetWebACLForResourceOutput, error) {
	var output wafv2.GetWebACLForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2ListAvailableManagedRuleGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2ListAvailableManagedRuleGroupsFuture) Get(ctx workflow.Context) (*wafv2.ListAvailableManagedRuleGroupsOutput, error) {
	var output wafv2.ListAvailableManagedRuleGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2ListIPSetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2ListIPSetsFuture) Get(ctx workflow.Context) (*wafv2.ListIPSetsOutput, error) {
	var output wafv2.ListIPSetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2ListLoggingConfigurationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2ListLoggingConfigurationsFuture) Get(ctx workflow.Context) (*wafv2.ListLoggingConfigurationsOutput, error) {
	var output wafv2.ListLoggingConfigurationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2ListRegexPatternSetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2ListRegexPatternSetsFuture) Get(ctx workflow.Context) (*wafv2.ListRegexPatternSetsOutput, error) {
	var output wafv2.ListRegexPatternSetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2ListResourcesForWebACLFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2ListResourcesForWebACLFuture) Get(ctx workflow.Context) (*wafv2.ListResourcesForWebACLOutput, error) {
	var output wafv2.ListResourcesForWebACLOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2ListRuleGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2ListRuleGroupsFuture) Get(ctx workflow.Context) (*wafv2.ListRuleGroupsOutput, error) {
	var output wafv2.ListRuleGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2ListTagsForResourceFuture) Get(ctx workflow.Context) (*wafv2.ListTagsForResourceOutput, error) {
	var output wafv2.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2ListWebACLsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2ListWebACLsFuture) Get(ctx workflow.Context) (*wafv2.ListWebACLsOutput, error) {
	var output wafv2.ListWebACLsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2PutLoggingConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2PutLoggingConfigurationFuture) Get(ctx workflow.Context) (*wafv2.PutLoggingConfigurationOutput, error) {
	var output wafv2.PutLoggingConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2PutPermissionPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2PutPermissionPolicyFuture) Get(ctx workflow.Context) (*wafv2.PutPermissionPolicyOutput, error) {
	var output wafv2.PutPermissionPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2TagResourceFuture) Get(ctx workflow.Context) (*wafv2.TagResourceOutput, error) {
	var output wafv2.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2UntagResourceFuture) Get(ctx workflow.Context) (*wafv2.UntagResourceOutput, error) {
	var output wafv2.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2UpdateIPSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2UpdateIPSetFuture) Get(ctx workflow.Context) (*wafv2.UpdateIPSetOutput, error) {
	var output wafv2.UpdateIPSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2UpdateRegexPatternSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2UpdateRegexPatternSetFuture) Get(ctx workflow.Context) (*wafv2.UpdateRegexPatternSetOutput, error) {
	var output wafv2.UpdateRegexPatternSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2UpdateRuleGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2UpdateRuleGroupFuture) Get(ctx workflow.Context) (*wafv2.UpdateRuleGroupOutput, error) {
	var output wafv2.UpdateRuleGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WAFV2UpdateWebACLFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WAFV2UpdateWebACLFuture) Get(ctx workflow.Context) (*wafv2.UpdateWebACLOutput, error) {
	var output wafv2.UpdateWebACLOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateWebACL(ctx workflow.Context, input *wafv2.AssociateWebACLInput) (*wafv2.AssociateWebACLOutput, error) {
	var output wafv2.AssociateWebACLOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.AssociateWebACL", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateWebACLAsync(ctx workflow.Context, input *wafv2.AssociateWebACLInput) *WAFV2AssociateWebACLFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.AssociateWebACL", input)
	return &WAFV2AssociateWebACLFuture{Future: future}
}

func (a *stub) CheckCapacity(ctx workflow.Context, input *wafv2.CheckCapacityInput) (*wafv2.CheckCapacityOutput, error) {
	var output wafv2.CheckCapacityOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.CheckCapacity", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CheckCapacityAsync(ctx workflow.Context, input *wafv2.CheckCapacityInput) *WAFV2CheckCapacityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.CheckCapacity", input)
	return &WAFV2CheckCapacityFuture{Future: future}
}

func (a *stub) CreateIPSet(ctx workflow.Context, input *wafv2.CreateIPSetInput) (*wafv2.CreateIPSetOutput, error) {
	var output wafv2.CreateIPSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.CreateIPSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateIPSetAsync(ctx workflow.Context, input *wafv2.CreateIPSetInput) *WAFV2CreateIPSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.CreateIPSet", input)
	return &WAFV2CreateIPSetFuture{Future: future}
}

func (a *stub) CreateRegexPatternSet(ctx workflow.Context, input *wafv2.CreateRegexPatternSetInput) (*wafv2.CreateRegexPatternSetOutput, error) {
	var output wafv2.CreateRegexPatternSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.CreateRegexPatternSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateRegexPatternSetAsync(ctx workflow.Context, input *wafv2.CreateRegexPatternSetInput) *WAFV2CreateRegexPatternSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.CreateRegexPatternSet", input)
	return &WAFV2CreateRegexPatternSetFuture{Future: future}
}

func (a *stub) CreateRuleGroup(ctx workflow.Context, input *wafv2.CreateRuleGroupInput) (*wafv2.CreateRuleGroupOutput, error) {
	var output wafv2.CreateRuleGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.CreateRuleGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateRuleGroupAsync(ctx workflow.Context, input *wafv2.CreateRuleGroupInput) *WAFV2CreateRuleGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.CreateRuleGroup", input)
	return &WAFV2CreateRuleGroupFuture{Future: future}
}

func (a *stub) CreateWebACL(ctx workflow.Context, input *wafv2.CreateWebACLInput) (*wafv2.CreateWebACLOutput, error) {
	var output wafv2.CreateWebACLOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.CreateWebACL", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateWebACLAsync(ctx workflow.Context, input *wafv2.CreateWebACLInput) *WAFV2CreateWebACLFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.CreateWebACL", input)
	return &WAFV2CreateWebACLFuture{Future: future}
}

func (a *stub) DeleteFirewallManagerRuleGroups(ctx workflow.Context, input *wafv2.DeleteFirewallManagerRuleGroupsInput) (*wafv2.DeleteFirewallManagerRuleGroupsOutput, error) {
	var output wafv2.DeleteFirewallManagerRuleGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.DeleteFirewallManagerRuleGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteFirewallManagerRuleGroupsAsync(ctx workflow.Context, input *wafv2.DeleteFirewallManagerRuleGroupsInput) *WAFV2DeleteFirewallManagerRuleGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.DeleteFirewallManagerRuleGroups", input)
	return &WAFV2DeleteFirewallManagerRuleGroupsFuture{Future: future}
}

func (a *stub) DeleteIPSet(ctx workflow.Context, input *wafv2.DeleteIPSetInput) (*wafv2.DeleteIPSetOutput, error) {
	var output wafv2.DeleteIPSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.DeleteIPSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteIPSetAsync(ctx workflow.Context, input *wafv2.DeleteIPSetInput) *WAFV2DeleteIPSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.DeleteIPSet", input)
	return &WAFV2DeleteIPSetFuture{Future: future}
}

func (a *stub) DeleteLoggingConfiguration(ctx workflow.Context, input *wafv2.DeleteLoggingConfigurationInput) (*wafv2.DeleteLoggingConfigurationOutput, error) {
	var output wafv2.DeleteLoggingConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.DeleteLoggingConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteLoggingConfigurationAsync(ctx workflow.Context, input *wafv2.DeleteLoggingConfigurationInput) *WAFV2DeleteLoggingConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.DeleteLoggingConfiguration", input)
	return &WAFV2DeleteLoggingConfigurationFuture{Future: future}
}

func (a *stub) DeletePermissionPolicy(ctx workflow.Context, input *wafv2.DeletePermissionPolicyInput) (*wafv2.DeletePermissionPolicyOutput, error) {
	var output wafv2.DeletePermissionPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.DeletePermissionPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeletePermissionPolicyAsync(ctx workflow.Context, input *wafv2.DeletePermissionPolicyInput) *WAFV2DeletePermissionPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.DeletePermissionPolicy", input)
	return &WAFV2DeletePermissionPolicyFuture{Future: future}
}

func (a *stub) DeleteRegexPatternSet(ctx workflow.Context, input *wafv2.DeleteRegexPatternSetInput) (*wafv2.DeleteRegexPatternSetOutput, error) {
	var output wafv2.DeleteRegexPatternSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.DeleteRegexPatternSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteRegexPatternSetAsync(ctx workflow.Context, input *wafv2.DeleteRegexPatternSetInput) *WAFV2DeleteRegexPatternSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.DeleteRegexPatternSet", input)
	return &WAFV2DeleteRegexPatternSetFuture{Future: future}
}

func (a *stub) DeleteRuleGroup(ctx workflow.Context, input *wafv2.DeleteRuleGroupInput) (*wafv2.DeleteRuleGroupOutput, error) {
	var output wafv2.DeleteRuleGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.DeleteRuleGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteRuleGroupAsync(ctx workflow.Context, input *wafv2.DeleteRuleGroupInput) *WAFV2DeleteRuleGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.DeleteRuleGroup", input)
	return &WAFV2DeleteRuleGroupFuture{Future: future}
}

func (a *stub) DeleteWebACL(ctx workflow.Context, input *wafv2.DeleteWebACLInput) (*wafv2.DeleteWebACLOutput, error) {
	var output wafv2.DeleteWebACLOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.DeleteWebACL", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteWebACLAsync(ctx workflow.Context, input *wafv2.DeleteWebACLInput) *WAFV2DeleteWebACLFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.DeleteWebACL", input)
	return &WAFV2DeleteWebACLFuture{Future: future}
}

func (a *stub) DescribeManagedRuleGroup(ctx workflow.Context, input *wafv2.DescribeManagedRuleGroupInput) (*wafv2.DescribeManagedRuleGroupOutput, error) {
	var output wafv2.DescribeManagedRuleGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.DescribeManagedRuleGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeManagedRuleGroupAsync(ctx workflow.Context, input *wafv2.DescribeManagedRuleGroupInput) *WAFV2DescribeManagedRuleGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.DescribeManagedRuleGroup", input)
	return &WAFV2DescribeManagedRuleGroupFuture{Future: future}
}

func (a *stub) DisassociateWebACL(ctx workflow.Context, input *wafv2.DisassociateWebACLInput) (*wafv2.DisassociateWebACLOutput, error) {
	var output wafv2.DisassociateWebACLOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.DisassociateWebACL", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateWebACLAsync(ctx workflow.Context, input *wafv2.DisassociateWebACLInput) *WAFV2DisassociateWebACLFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.DisassociateWebACL", input)
	return &WAFV2DisassociateWebACLFuture{Future: future}
}

func (a *stub) GetIPSet(ctx workflow.Context, input *wafv2.GetIPSetInput) (*wafv2.GetIPSetOutput, error) {
	var output wafv2.GetIPSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.GetIPSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetIPSetAsync(ctx workflow.Context, input *wafv2.GetIPSetInput) *WAFV2GetIPSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.GetIPSet", input)
	return &WAFV2GetIPSetFuture{Future: future}
}

func (a *stub) GetLoggingConfiguration(ctx workflow.Context, input *wafv2.GetLoggingConfigurationInput) (*wafv2.GetLoggingConfigurationOutput, error) {
	var output wafv2.GetLoggingConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.GetLoggingConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetLoggingConfigurationAsync(ctx workflow.Context, input *wafv2.GetLoggingConfigurationInput) *WAFV2GetLoggingConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.GetLoggingConfiguration", input)
	return &WAFV2GetLoggingConfigurationFuture{Future: future}
}

func (a *stub) GetPermissionPolicy(ctx workflow.Context, input *wafv2.GetPermissionPolicyInput) (*wafv2.GetPermissionPolicyOutput, error) {
	var output wafv2.GetPermissionPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.GetPermissionPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetPermissionPolicyAsync(ctx workflow.Context, input *wafv2.GetPermissionPolicyInput) *WAFV2GetPermissionPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.GetPermissionPolicy", input)
	return &WAFV2GetPermissionPolicyFuture{Future: future}
}

func (a *stub) GetRateBasedStatementManagedKeys(ctx workflow.Context, input *wafv2.GetRateBasedStatementManagedKeysInput) (*wafv2.GetRateBasedStatementManagedKeysOutput, error) {
	var output wafv2.GetRateBasedStatementManagedKeysOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.GetRateBasedStatementManagedKeys", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetRateBasedStatementManagedKeysAsync(ctx workflow.Context, input *wafv2.GetRateBasedStatementManagedKeysInput) *WAFV2GetRateBasedStatementManagedKeysFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.GetRateBasedStatementManagedKeys", input)
	return &WAFV2GetRateBasedStatementManagedKeysFuture{Future: future}
}

func (a *stub) GetRegexPatternSet(ctx workflow.Context, input *wafv2.GetRegexPatternSetInput) (*wafv2.GetRegexPatternSetOutput, error) {
	var output wafv2.GetRegexPatternSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.GetRegexPatternSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetRegexPatternSetAsync(ctx workflow.Context, input *wafv2.GetRegexPatternSetInput) *WAFV2GetRegexPatternSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.GetRegexPatternSet", input)
	return &WAFV2GetRegexPatternSetFuture{Future: future}
}

func (a *stub) GetRuleGroup(ctx workflow.Context, input *wafv2.GetRuleGroupInput) (*wafv2.GetRuleGroupOutput, error) {
	var output wafv2.GetRuleGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.GetRuleGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetRuleGroupAsync(ctx workflow.Context, input *wafv2.GetRuleGroupInput) *WAFV2GetRuleGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.GetRuleGroup", input)
	return &WAFV2GetRuleGroupFuture{Future: future}
}

func (a *stub) GetSampledRequests(ctx workflow.Context, input *wafv2.GetSampledRequestsInput) (*wafv2.GetSampledRequestsOutput, error) {
	var output wafv2.GetSampledRequestsOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.GetSampledRequests", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetSampledRequestsAsync(ctx workflow.Context, input *wafv2.GetSampledRequestsInput) *WAFV2GetSampledRequestsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.GetSampledRequests", input)
	return &WAFV2GetSampledRequestsFuture{Future: future}
}

func (a *stub) GetWebACL(ctx workflow.Context, input *wafv2.GetWebACLInput) (*wafv2.GetWebACLOutput, error) {
	var output wafv2.GetWebACLOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.GetWebACL", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetWebACLAsync(ctx workflow.Context, input *wafv2.GetWebACLInput) *WAFV2GetWebACLFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.GetWebACL", input)
	return &WAFV2GetWebACLFuture{Future: future}
}

func (a *stub) GetWebACLForResource(ctx workflow.Context, input *wafv2.GetWebACLForResourceInput) (*wafv2.GetWebACLForResourceOutput, error) {
	var output wafv2.GetWebACLForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.GetWebACLForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetWebACLForResourceAsync(ctx workflow.Context, input *wafv2.GetWebACLForResourceInput) *WAFV2GetWebACLForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.GetWebACLForResource", input)
	return &WAFV2GetWebACLForResourceFuture{Future: future}
}

func (a *stub) ListAvailableManagedRuleGroups(ctx workflow.Context, input *wafv2.ListAvailableManagedRuleGroupsInput) (*wafv2.ListAvailableManagedRuleGroupsOutput, error) {
	var output wafv2.ListAvailableManagedRuleGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.ListAvailableManagedRuleGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAvailableManagedRuleGroupsAsync(ctx workflow.Context, input *wafv2.ListAvailableManagedRuleGroupsInput) *WAFV2ListAvailableManagedRuleGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.ListAvailableManagedRuleGroups", input)
	return &WAFV2ListAvailableManagedRuleGroupsFuture{Future: future}
}

func (a *stub) ListIPSets(ctx workflow.Context, input *wafv2.ListIPSetsInput) (*wafv2.ListIPSetsOutput, error) {
	var output wafv2.ListIPSetsOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.ListIPSets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListIPSetsAsync(ctx workflow.Context, input *wafv2.ListIPSetsInput) *WAFV2ListIPSetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.ListIPSets", input)
	return &WAFV2ListIPSetsFuture{Future: future}
}

func (a *stub) ListLoggingConfigurations(ctx workflow.Context, input *wafv2.ListLoggingConfigurationsInput) (*wafv2.ListLoggingConfigurationsOutput, error) {
	var output wafv2.ListLoggingConfigurationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.ListLoggingConfigurations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListLoggingConfigurationsAsync(ctx workflow.Context, input *wafv2.ListLoggingConfigurationsInput) *WAFV2ListLoggingConfigurationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.ListLoggingConfigurations", input)
	return &WAFV2ListLoggingConfigurationsFuture{Future: future}
}

func (a *stub) ListRegexPatternSets(ctx workflow.Context, input *wafv2.ListRegexPatternSetsInput) (*wafv2.ListRegexPatternSetsOutput, error) {
	var output wafv2.ListRegexPatternSetsOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.ListRegexPatternSets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListRegexPatternSetsAsync(ctx workflow.Context, input *wafv2.ListRegexPatternSetsInput) *WAFV2ListRegexPatternSetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.ListRegexPatternSets", input)
	return &WAFV2ListRegexPatternSetsFuture{Future: future}
}

func (a *stub) ListResourcesForWebACL(ctx workflow.Context, input *wafv2.ListResourcesForWebACLInput) (*wafv2.ListResourcesForWebACLOutput, error) {
	var output wafv2.ListResourcesForWebACLOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.ListResourcesForWebACL", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListResourcesForWebACLAsync(ctx workflow.Context, input *wafv2.ListResourcesForWebACLInput) *WAFV2ListResourcesForWebACLFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.ListResourcesForWebACL", input)
	return &WAFV2ListResourcesForWebACLFuture{Future: future}
}

func (a *stub) ListRuleGroups(ctx workflow.Context, input *wafv2.ListRuleGroupsInput) (*wafv2.ListRuleGroupsOutput, error) {
	var output wafv2.ListRuleGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.ListRuleGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListRuleGroupsAsync(ctx workflow.Context, input *wafv2.ListRuleGroupsInput) *WAFV2ListRuleGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.ListRuleGroups", input)
	return &WAFV2ListRuleGroupsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *wafv2.ListTagsForResourceInput) (*wafv2.ListTagsForResourceOutput, error) {
	var output wafv2.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *wafv2.ListTagsForResourceInput) *WAFV2ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.ListTagsForResource", input)
	return &WAFV2ListTagsForResourceFuture{Future: future}
}

func (a *stub) ListWebACLs(ctx workflow.Context, input *wafv2.ListWebACLsInput) (*wafv2.ListWebACLsOutput, error) {
	var output wafv2.ListWebACLsOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.ListWebACLs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListWebACLsAsync(ctx workflow.Context, input *wafv2.ListWebACLsInput) *WAFV2ListWebACLsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.ListWebACLs", input)
	return &WAFV2ListWebACLsFuture{Future: future}
}

func (a *stub) PutLoggingConfiguration(ctx workflow.Context, input *wafv2.PutLoggingConfigurationInput) (*wafv2.PutLoggingConfigurationOutput, error) {
	var output wafv2.PutLoggingConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.PutLoggingConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutLoggingConfigurationAsync(ctx workflow.Context, input *wafv2.PutLoggingConfigurationInput) *WAFV2PutLoggingConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.PutLoggingConfiguration", input)
	return &WAFV2PutLoggingConfigurationFuture{Future: future}
}

func (a *stub) PutPermissionPolicy(ctx workflow.Context, input *wafv2.PutPermissionPolicyInput) (*wafv2.PutPermissionPolicyOutput, error) {
	var output wafv2.PutPermissionPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.PutPermissionPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutPermissionPolicyAsync(ctx workflow.Context, input *wafv2.PutPermissionPolicyInput) *WAFV2PutPermissionPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.PutPermissionPolicy", input)
	return &WAFV2PutPermissionPolicyFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *wafv2.TagResourceInput) (*wafv2.TagResourceOutput, error) {
	var output wafv2.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *wafv2.TagResourceInput) *WAFV2TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.TagResource", input)
	return &WAFV2TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *wafv2.UntagResourceInput) (*wafv2.UntagResourceOutput, error) {
	var output wafv2.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *wafv2.UntagResourceInput) *WAFV2UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.UntagResource", input)
	return &WAFV2UntagResourceFuture{Future: future}
}

func (a *stub) UpdateIPSet(ctx workflow.Context, input *wafv2.UpdateIPSetInput) (*wafv2.UpdateIPSetOutput, error) {
	var output wafv2.UpdateIPSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.UpdateIPSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateIPSetAsync(ctx workflow.Context, input *wafv2.UpdateIPSetInput) *WAFV2UpdateIPSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.UpdateIPSet", input)
	return &WAFV2UpdateIPSetFuture{Future: future}
}

func (a *stub) UpdateRegexPatternSet(ctx workflow.Context, input *wafv2.UpdateRegexPatternSetInput) (*wafv2.UpdateRegexPatternSetOutput, error) {
	var output wafv2.UpdateRegexPatternSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.UpdateRegexPatternSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateRegexPatternSetAsync(ctx workflow.Context, input *wafv2.UpdateRegexPatternSetInput) *WAFV2UpdateRegexPatternSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.UpdateRegexPatternSet", input)
	return &WAFV2UpdateRegexPatternSetFuture{Future: future}
}

func (a *stub) UpdateRuleGroup(ctx workflow.Context, input *wafv2.UpdateRuleGroupInput) (*wafv2.UpdateRuleGroupOutput, error) {
	var output wafv2.UpdateRuleGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.UpdateRuleGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateRuleGroupAsync(ctx workflow.Context, input *wafv2.UpdateRuleGroupInput) *WAFV2UpdateRuleGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.UpdateRuleGroup", input)
	return &WAFV2UpdateRuleGroupFuture{Future: future}
}

func (a *stub) UpdateWebACL(ctx workflow.Context, input *wafv2.UpdateWebACLInput) (*wafv2.UpdateWebACLOutput, error) {
	var output wafv2.UpdateWebACLOutput
	err := workflow.ExecuteActivity(ctx, "aws.wafv2.UpdateWebACL", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateWebACLAsync(ctx workflow.Context, input *wafv2.UpdateWebACLInput) *WAFV2UpdateWebACLFuture {
	future := workflow.ExecuteActivity(ctx, "aws.wafv2.UpdateWebACL", input)
	return &WAFV2UpdateWebACLFuture{Future: future}
}
