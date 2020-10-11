// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package wafregionalstub

import (
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/aws/aws-sdk-go/service/wafregional"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	AssociateWebACL(ctx workflow.Context, input *wafregional.AssociateWebACLInput) (*wafregional.AssociateWebACLOutput, error)
	AssociateWebACLAsync(ctx workflow.Context, input *wafregional.AssociateWebACLInput) *WAFRegionalAssociateWebACLFuture

	CreateByteMatchSet(ctx workflow.Context, input *waf.CreateByteMatchSetInput) (*waf.CreateByteMatchSetOutput, error)
	CreateByteMatchSetAsync(ctx workflow.Context, input *waf.CreateByteMatchSetInput) *WAFRegionalCreateByteMatchSetFuture

	CreateGeoMatchSet(ctx workflow.Context, input *waf.CreateGeoMatchSetInput) (*waf.CreateGeoMatchSetOutput, error)
	CreateGeoMatchSetAsync(ctx workflow.Context, input *waf.CreateGeoMatchSetInput) *WAFRegionalCreateGeoMatchSetFuture

	CreateIPSet(ctx workflow.Context, input *waf.CreateIPSetInput) (*waf.CreateIPSetOutput, error)
	CreateIPSetAsync(ctx workflow.Context, input *waf.CreateIPSetInput) *WAFRegionalCreateIPSetFuture

	CreateRateBasedRule(ctx workflow.Context, input *waf.CreateRateBasedRuleInput) (*waf.CreateRateBasedRuleOutput, error)
	CreateRateBasedRuleAsync(ctx workflow.Context, input *waf.CreateRateBasedRuleInput) *WAFRegionalCreateRateBasedRuleFuture

	CreateRegexMatchSet(ctx workflow.Context, input *waf.CreateRegexMatchSetInput) (*waf.CreateRegexMatchSetOutput, error)
	CreateRegexMatchSetAsync(ctx workflow.Context, input *waf.CreateRegexMatchSetInput) *WAFRegionalCreateRegexMatchSetFuture

	CreateRegexPatternSet(ctx workflow.Context, input *waf.CreateRegexPatternSetInput) (*waf.CreateRegexPatternSetOutput, error)
	CreateRegexPatternSetAsync(ctx workflow.Context, input *waf.CreateRegexPatternSetInput) *WAFRegionalCreateRegexPatternSetFuture

	CreateRule(ctx workflow.Context, input *waf.CreateRuleInput) (*waf.CreateRuleOutput, error)
	CreateRuleAsync(ctx workflow.Context, input *waf.CreateRuleInput) *WAFRegionalCreateRuleFuture

	CreateRuleGroup(ctx workflow.Context, input *waf.CreateRuleGroupInput) (*waf.CreateRuleGroupOutput, error)
	CreateRuleGroupAsync(ctx workflow.Context, input *waf.CreateRuleGroupInput) *WAFRegionalCreateRuleGroupFuture

	CreateSizeConstraintSet(ctx workflow.Context, input *waf.CreateSizeConstraintSetInput) (*waf.CreateSizeConstraintSetOutput, error)
	CreateSizeConstraintSetAsync(ctx workflow.Context, input *waf.CreateSizeConstraintSetInput) *WAFRegionalCreateSizeConstraintSetFuture

	CreateSqlInjectionMatchSet(ctx workflow.Context, input *waf.CreateSqlInjectionMatchSetInput) (*waf.CreateSqlInjectionMatchSetOutput, error)
	CreateSqlInjectionMatchSetAsync(ctx workflow.Context, input *waf.CreateSqlInjectionMatchSetInput) *WAFRegionalCreateSqlInjectionMatchSetFuture

	CreateWebACL(ctx workflow.Context, input *waf.CreateWebACLInput) (*waf.CreateWebACLOutput, error)
	CreateWebACLAsync(ctx workflow.Context, input *waf.CreateWebACLInput) *WAFRegionalCreateWebACLFuture

	CreateWebACLMigrationStack(ctx workflow.Context, input *waf.CreateWebACLMigrationStackInput) (*waf.CreateWebACLMigrationStackOutput, error)
	CreateWebACLMigrationStackAsync(ctx workflow.Context, input *waf.CreateWebACLMigrationStackInput) *WAFRegionalCreateWebACLMigrationStackFuture

	CreateXssMatchSet(ctx workflow.Context, input *waf.CreateXssMatchSetInput) (*waf.CreateXssMatchSetOutput, error)
	CreateXssMatchSetAsync(ctx workflow.Context, input *waf.CreateXssMatchSetInput) *WAFRegionalCreateXssMatchSetFuture

	DeleteByteMatchSet(ctx workflow.Context, input *waf.DeleteByteMatchSetInput) (*waf.DeleteByteMatchSetOutput, error)
	DeleteByteMatchSetAsync(ctx workflow.Context, input *waf.DeleteByteMatchSetInput) *WAFRegionalDeleteByteMatchSetFuture

	DeleteGeoMatchSet(ctx workflow.Context, input *waf.DeleteGeoMatchSetInput) (*waf.DeleteGeoMatchSetOutput, error)
	DeleteGeoMatchSetAsync(ctx workflow.Context, input *waf.DeleteGeoMatchSetInput) *WAFRegionalDeleteGeoMatchSetFuture

	DeleteIPSet(ctx workflow.Context, input *waf.DeleteIPSetInput) (*waf.DeleteIPSetOutput, error)
	DeleteIPSetAsync(ctx workflow.Context, input *waf.DeleteIPSetInput) *WAFRegionalDeleteIPSetFuture

	DeleteLoggingConfiguration(ctx workflow.Context, input *waf.DeleteLoggingConfigurationInput) (*waf.DeleteLoggingConfigurationOutput, error)
	DeleteLoggingConfigurationAsync(ctx workflow.Context, input *waf.DeleteLoggingConfigurationInput) *WAFRegionalDeleteLoggingConfigurationFuture

	DeletePermissionPolicy(ctx workflow.Context, input *waf.DeletePermissionPolicyInput) (*waf.DeletePermissionPolicyOutput, error)
	DeletePermissionPolicyAsync(ctx workflow.Context, input *waf.DeletePermissionPolicyInput) *WAFRegionalDeletePermissionPolicyFuture

	DeleteRateBasedRule(ctx workflow.Context, input *waf.DeleteRateBasedRuleInput) (*waf.DeleteRateBasedRuleOutput, error)
	DeleteRateBasedRuleAsync(ctx workflow.Context, input *waf.DeleteRateBasedRuleInput) *WAFRegionalDeleteRateBasedRuleFuture

	DeleteRegexMatchSet(ctx workflow.Context, input *waf.DeleteRegexMatchSetInput) (*waf.DeleteRegexMatchSetOutput, error)
	DeleteRegexMatchSetAsync(ctx workflow.Context, input *waf.DeleteRegexMatchSetInput) *WAFRegionalDeleteRegexMatchSetFuture

	DeleteRegexPatternSet(ctx workflow.Context, input *waf.DeleteRegexPatternSetInput) (*waf.DeleteRegexPatternSetOutput, error)
	DeleteRegexPatternSetAsync(ctx workflow.Context, input *waf.DeleteRegexPatternSetInput) *WAFRegionalDeleteRegexPatternSetFuture

	DeleteRule(ctx workflow.Context, input *waf.DeleteRuleInput) (*waf.DeleteRuleOutput, error)
	DeleteRuleAsync(ctx workflow.Context, input *waf.DeleteRuleInput) *WAFRegionalDeleteRuleFuture

	DeleteRuleGroup(ctx workflow.Context, input *waf.DeleteRuleGroupInput) (*waf.DeleteRuleGroupOutput, error)
	DeleteRuleGroupAsync(ctx workflow.Context, input *waf.DeleteRuleGroupInput) *WAFRegionalDeleteRuleGroupFuture

	DeleteSizeConstraintSet(ctx workflow.Context, input *waf.DeleteSizeConstraintSetInput) (*waf.DeleteSizeConstraintSetOutput, error)
	DeleteSizeConstraintSetAsync(ctx workflow.Context, input *waf.DeleteSizeConstraintSetInput) *WAFRegionalDeleteSizeConstraintSetFuture

	DeleteSqlInjectionMatchSet(ctx workflow.Context, input *waf.DeleteSqlInjectionMatchSetInput) (*waf.DeleteSqlInjectionMatchSetOutput, error)
	DeleteSqlInjectionMatchSetAsync(ctx workflow.Context, input *waf.DeleteSqlInjectionMatchSetInput) *WAFRegionalDeleteSqlInjectionMatchSetFuture

	DeleteWebACL(ctx workflow.Context, input *waf.DeleteWebACLInput) (*waf.DeleteWebACLOutput, error)
	DeleteWebACLAsync(ctx workflow.Context, input *waf.DeleteWebACLInput) *WAFRegionalDeleteWebACLFuture

	DeleteXssMatchSet(ctx workflow.Context, input *waf.DeleteXssMatchSetInput) (*waf.DeleteXssMatchSetOutput, error)
	DeleteXssMatchSetAsync(ctx workflow.Context, input *waf.DeleteXssMatchSetInput) *WAFRegionalDeleteXssMatchSetFuture

	DisassociateWebACL(ctx workflow.Context, input *wafregional.DisassociateWebACLInput) (*wafregional.DisassociateWebACLOutput, error)
	DisassociateWebACLAsync(ctx workflow.Context, input *wafregional.DisassociateWebACLInput) *WAFRegionalDisassociateWebACLFuture

	GetByteMatchSet(ctx workflow.Context, input *waf.GetByteMatchSetInput) (*waf.GetByteMatchSetOutput, error)
	GetByteMatchSetAsync(ctx workflow.Context, input *waf.GetByteMatchSetInput) *WAFRegionalGetByteMatchSetFuture

	GetChangeToken(ctx workflow.Context, input *waf.GetChangeTokenInput) (*waf.GetChangeTokenOutput, error)
	GetChangeTokenAsync(ctx workflow.Context, input *waf.GetChangeTokenInput) *WAFRegionalGetChangeTokenFuture

	GetChangeTokenStatus(ctx workflow.Context, input *waf.GetChangeTokenStatusInput) (*waf.GetChangeTokenStatusOutput, error)
	GetChangeTokenStatusAsync(ctx workflow.Context, input *waf.GetChangeTokenStatusInput) *WAFRegionalGetChangeTokenStatusFuture

	GetGeoMatchSet(ctx workflow.Context, input *waf.GetGeoMatchSetInput) (*waf.GetGeoMatchSetOutput, error)
	GetGeoMatchSetAsync(ctx workflow.Context, input *waf.GetGeoMatchSetInput) *WAFRegionalGetGeoMatchSetFuture

	GetIPSet(ctx workflow.Context, input *waf.GetIPSetInput) (*waf.GetIPSetOutput, error)
	GetIPSetAsync(ctx workflow.Context, input *waf.GetIPSetInput) *WAFRegionalGetIPSetFuture

	GetLoggingConfiguration(ctx workflow.Context, input *waf.GetLoggingConfigurationInput) (*waf.GetLoggingConfigurationOutput, error)
	GetLoggingConfigurationAsync(ctx workflow.Context, input *waf.GetLoggingConfigurationInput) *WAFRegionalGetLoggingConfigurationFuture

	GetPermissionPolicy(ctx workflow.Context, input *waf.GetPermissionPolicyInput) (*waf.GetPermissionPolicyOutput, error)
	GetPermissionPolicyAsync(ctx workflow.Context, input *waf.GetPermissionPolicyInput) *WAFRegionalGetPermissionPolicyFuture

	GetRateBasedRule(ctx workflow.Context, input *waf.GetRateBasedRuleInput) (*waf.GetRateBasedRuleOutput, error)
	GetRateBasedRuleAsync(ctx workflow.Context, input *waf.GetRateBasedRuleInput) *WAFRegionalGetRateBasedRuleFuture

	GetRateBasedRuleManagedKeys(ctx workflow.Context, input *waf.GetRateBasedRuleManagedKeysInput) (*waf.GetRateBasedRuleManagedKeysOutput, error)
	GetRateBasedRuleManagedKeysAsync(ctx workflow.Context, input *waf.GetRateBasedRuleManagedKeysInput) *WAFRegionalGetRateBasedRuleManagedKeysFuture

	GetRegexMatchSet(ctx workflow.Context, input *waf.GetRegexMatchSetInput) (*waf.GetRegexMatchSetOutput, error)
	GetRegexMatchSetAsync(ctx workflow.Context, input *waf.GetRegexMatchSetInput) *WAFRegionalGetRegexMatchSetFuture

	GetRegexPatternSet(ctx workflow.Context, input *waf.GetRegexPatternSetInput) (*waf.GetRegexPatternSetOutput, error)
	GetRegexPatternSetAsync(ctx workflow.Context, input *waf.GetRegexPatternSetInput) *WAFRegionalGetRegexPatternSetFuture

	GetRule(ctx workflow.Context, input *waf.GetRuleInput) (*waf.GetRuleOutput, error)
	GetRuleAsync(ctx workflow.Context, input *waf.GetRuleInput) *WAFRegionalGetRuleFuture

	GetRuleGroup(ctx workflow.Context, input *waf.GetRuleGroupInput) (*waf.GetRuleGroupOutput, error)
	GetRuleGroupAsync(ctx workflow.Context, input *waf.GetRuleGroupInput) *WAFRegionalGetRuleGroupFuture

	GetSampledRequests(ctx workflow.Context, input *waf.GetSampledRequestsInput) (*waf.GetSampledRequestsOutput, error)
	GetSampledRequestsAsync(ctx workflow.Context, input *waf.GetSampledRequestsInput) *WAFRegionalGetSampledRequestsFuture

	GetSizeConstraintSet(ctx workflow.Context, input *waf.GetSizeConstraintSetInput) (*waf.GetSizeConstraintSetOutput, error)
	GetSizeConstraintSetAsync(ctx workflow.Context, input *waf.GetSizeConstraintSetInput) *WAFRegionalGetSizeConstraintSetFuture

	GetSqlInjectionMatchSet(ctx workflow.Context, input *waf.GetSqlInjectionMatchSetInput) (*waf.GetSqlInjectionMatchSetOutput, error)
	GetSqlInjectionMatchSetAsync(ctx workflow.Context, input *waf.GetSqlInjectionMatchSetInput) *WAFRegionalGetSqlInjectionMatchSetFuture

	GetWebACL(ctx workflow.Context, input *waf.GetWebACLInput) (*waf.GetWebACLOutput, error)
	GetWebACLAsync(ctx workflow.Context, input *waf.GetWebACLInput) *WAFRegionalGetWebACLFuture

	GetWebACLForResource(ctx workflow.Context, input *wafregional.GetWebACLForResourceInput) (*wafregional.GetWebACLForResourceOutput, error)
	GetWebACLForResourceAsync(ctx workflow.Context, input *wafregional.GetWebACLForResourceInput) *WAFRegionalGetWebACLForResourceFuture

	GetXssMatchSet(ctx workflow.Context, input *waf.GetXssMatchSetInput) (*waf.GetXssMatchSetOutput, error)
	GetXssMatchSetAsync(ctx workflow.Context, input *waf.GetXssMatchSetInput) *WAFRegionalGetXssMatchSetFuture

	ListActivatedRulesInRuleGroup(ctx workflow.Context, input *waf.ListActivatedRulesInRuleGroupInput) (*waf.ListActivatedRulesInRuleGroupOutput, error)
	ListActivatedRulesInRuleGroupAsync(ctx workflow.Context, input *waf.ListActivatedRulesInRuleGroupInput) *WAFRegionalListActivatedRulesInRuleGroupFuture

	ListByteMatchSets(ctx workflow.Context, input *waf.ListByteMatchSetsInput) (*waf.ListByteMatchSetsOutput, error)
	ListByteMatchSetsAsync(ctx workflow.Context, input *waf.ListByteMatchSetsInput) *WAFRegionalListByteMatchSetsFuture

	ListGeoMatchSets(ctx workflow.Context, input *waf.ListGeoMatchSetsInput) (*waf.ListGeoMatchSetsOutput, error)
	ListGeoMatchSetsAsync(ctx workflow.Context, input *waf.ListGeoMatchSetsInput) *WAFRegionalListGeoMatchSetsFuture

	ListIPSets(ctx workflow.Context, input *waf.ListIPSetsInput) (*waf.ListIPSetsOutput, error)
	ListIPSetsAsync(ctx workflow.Context, input *waf.ListIPSetsInput) *WAFRegionalListIPSetsFuture

	ListLoggingConfigurations(ctx workflow.Context, input *waf.ListLoggingConfigurationsInput) (*waf.ListLoggingConfigurationsOutput, error)
	ListLoggingConfigurationsAsync(ctx workflow.Context, input *waf.ListLoggingConfigurationsInput) *WAFRegionalListLoggingConfigurationsFuture

	ListRateBasedRules(ctx workflow.Context, input *waf.ListRateBasedRulesInput) (*waf.ListRateBasedRulesOutput, error)
	ListRateBasedRulesAsync(ctx workflow.Context, input *waf.ListRateBasedRulesInput) *WAFRegionalListRateBasedRulesFuture

	ListRegexMatchSets(ctx workflow.Context, input *waf.ListRegexMatchSetsInput) (*waf.ListRegexMatchSetsOutput, error)
	ListRegexMatchSetsAsync(ctx workflow.Context, input *waf.ListRegexMatchSetsInput) *WAFRegionalListRegexMatchSetsFuture

	ListRegexPatternSets(ctx workflow.Context, input *waf.ListRegexPatternSetsInput) (*waf.ListRegexPatternSetsOutput, error)
	ListRegexPatternSetsAsync(ctx workflow.Context, input *waf.ListRegexPatternSetsInput) *WAFRegionalListRegexPatternSetsFuture

	ListResourcesForWebACL(ctx workflow.Context, input *wafregional.ListResourcesForWebACLInput) (*wafregional.ListResourcesForWebACLOutput, error)
	ListResourcesForWebACLAsync(ctx workflow.Context, input *wafregional.ListResourcesForWebACLInput) *WAFRegionalListResourcesForWebACLFuture

	ListRuleGroups(ctx workflow.Context, input *waf.ListRuleGroupsInput) (*waf.ListRuleGroupsOutput, error)
	ListRuleGroupsAsync(ctx workflow.Context, input *waf.ListRuleGroupsInput) *WAFRegionalListRuleGroupsFuture

	ListRules(ctx workflow.Context, input *waf.ListRulesInput) (*waf.ListRulesOutput, error)
	ListRulesAsync(ctx workflow.Context, input *waf.ListRulesInput) *WAFRegionalListRulesFuture

	ListSizeConstraintSets(ctx workflow.Context, input *waf.ListSizeConstraintSetsInput) (*waf.ListSizeConstraintSetsOutput, error)
	ListSizeConstraintSetsAsync(ctx workflow.Context, input *waf.ListSizeConstraintSetsInput) *WAFRegionalListSizeConstraintSetsFuture

	ListSqlInjectionMatchSets(ctx workflow.Context, input *waf.ListSqlInjectionMatchSetsInput) (*waf.ListSqlInjectionMatchSetsOutput, error)
	ListSqlInjectionMatchSetsAsync(ctx workflow.Context, input *waf.ListSqlInjectionMatchSetsInput) *WAFRegionalListSqlInjectionMatchSetsFuture

	ListSubscribedRuleGroups(ctx workflow.Context, input *waf.ListSubscribedRuleGroupsInput) (*waf.ListSubscribedRuleGroupsOutput, error)
	ListSubscribedRuleGroupsAsync(ctx workflow.Context, input *waf.ListSubscribedRuleGroupsInput) *WAFRegionalListSubscribedRuleGroupsFuture

	ListTagsForResource(ctx workflow.Context, input *waf.ListTagsForResourceInput) (*waf.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *waf.ListTagsForResourceInput) *WAFRegionalListTagsForResourceFuture

	ListWebACLs(ctx workflow.Context, input *waf.ListWebACLsInput) (*waf.ListWebACLsOutput, error)
	ListWebACLsAsync(ctx workflow.Context, input *waf.ListWebACLsInput) *WAFRegionalListWebACLsFuture

	ListXssMatchSets(ctx workflow.Context, input *waf.ListXssMatchSetsInput) (*waf.ListXssMatchSetsOutput, error)
	ListXssMatchSetsAsync(ctx workflow.Context, input *waf.ListXssMatchSetsInput) *WAFRegionalListXssMatchSetsFuture

	PutLoggingConfiguration(ctx workflow.Context, input *waf.PutLoggingConfigurationInput) (*waf.PutLoggingConfigurationOutput, error)
	PutLoggingConfigurationAsync(ctx workflow.Context, input *waf.PutLoggingConfigurationInput) *WAFRegionalPutLoggingConfigurationFuture

	PutPermissionPolicy(ctx workflow.Context, input *waf.PutPermissionPolicyInput) (*waf.PutPermissionPolicyOutput, error)
	PutPermissionPolicyAsync(ctx workflow.Context, input *waf.PutPermissionPolicyInput) *WAFRegionalPutPermissionPolicyFuture

	TagResource(ctx workflow.Context, input *waf.TagResourceInput) (*waf.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *waf.TagResourceInput) *WAFRegionalTagResourceFuture

	UntagResource(ctx workflow.Context, input *waf.UntagResourceInput) (*waf.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *waf.UntagResourceInput) *WAFRegionalUntagResourceFuture

	UpdateByteMatchSet(ctx workflow.Context, input *waf.UpdateByteMatchSetInput) (*waf.UpdateByteMatchSetOutput, error)
	UpdateByteMatchSetAsync(ctx workflow.Context, input *waf.UpdateByteMatchSetInput) *WAFRegionalUpdateByteMatchSetFuture

	UpdateGeoMatchSet(ctx workflow.Context, input *waf.UpdateGeoMatchSetInput) (*waf.UpdateGeoMatchSetOutput, error)
	UpdateGeoMatchSetAsync(ctx workflow.Context, input *waf.UpdateGeoMatchSetInput) *WAFRegionalUpdateGeoMatchSetFuture

	UpdateIPSet(ctx workflow.Context, input *waf.UpdateIPSetInput) (*waf.UpdateIPSetOutput, error)
	UpdateIPSetAsync(ctx workflow.Context, input *waf.UpdateIPSetInput) *WAFRegionalUpdateIPSetFuture

	UpdateRateBasedRule(ctx workflow.Context, input *waf.UpdateRateBasedRuleInput) (*waf.UpdateRateBasedRuleOutput, error)
	UpdateRateBasedRuleAsync(ctx workflow.Context, input *waf.UpdateRateBasedRuleInput) *WAFRegionalUpdateRateBasedRuleFuture

	UpdateRegexMatchSet(ctx workflow.Context, input *waf.UpdateRegexMatchSetInput) (*waf.UpdateRegexMatchSetOutput, error)
	UpdateRegexMatchSetAsync(ctx workflow.Context, input *waf.UpdateRegexMatchSetInput) *WAFRegionalUpdateRegexMatchSetFuture

	UpdateRegexPatternSet(ctx workflow.Context, input *waf.UpdateRegexPatternSetInput) (*waf.UpdateRegexPatternSetOutput, error)
	UpdateRegexPatternSetAsync(ctx workflow.Context, input *waf.UpdateRegexPatternSetInput) *WAFRegionalUpdateRegexPatternSetFuture

	UpdateRule(ctx workflow.Context, input *waf.UpdateRuleInput) (*waf.UpdateRuleOutput, error)
	UpdateRuleAsync(ctx workflow.Context, input *waf.UpdateRuleInput) *WAFRegionalUpdateRuleFuture

	UpdateRuleGroup(ctx workflow.Context, input *waf.UpdateRuleGroupInput) (*waf.UpdateRuleGroupOutput, error)
	UpdateRuleGroupAsync(ctx workflow.Context, input *waf.UpdateRuleGroupInput) *WAFRegionalUpdateRuleGroupFuture

	UpdateSizeConstraintSet(ctx workflow.Context, input *waf.UpdateSizeConstraintSetInput) (*waf.UpdateSizeConstraintSetOutput, error)
	UpdateSizeConstraintSetAsync(ctx workflow.Context, input *waf.UpdateSizeConstraintSetInput) *WAFRegionalUpdateSizeConstraintSetFuture

	UpdateSqlInjectionMatchSet(ctx workflow.Context, input *waf.UpdateSqlInjectionMatchSetInput) (*waf.UpdateSqlInjectionMatchSetOutput, error)
	UpdateSqlInjectionMatchSetAsync(ctx workflow.Context, input *waf.UpdateSqlInjectionMatchSetInput) *WAFRegionalUpdateSqlInjectionMatchSetFuture

	UpdateWebACL(ctx workflow.Context, input *waf.UpdateWebACLInput) (*waf.UpdateWebACLOutput, error)
	UpdateWebACLAsync(ctx workflow.Context, input *waf.UpdateWebACLInput) *WAFRegionalUpdateWebACLFuture

	UpdateXssMatchSet(ctx workflow.Context, input *waf.UpdateXssMatchSetInput) (*waf.UpdateXssMatchSetOutput, error)
	UpdateXssMatchSetAsync(ctx workflow.Context, input *waf.UpdateXssMatchSetInput) *WAFRegionalUpdateXssMatchSetFuture
}

func NewClient() Client {
	return &stub{}
}
