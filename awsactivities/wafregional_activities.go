package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/wafregional"
	"github.com/aws/aws-sdk-go/service/wafregional/wafregionaliface"
)

type WAFRegionalActivities struct {
	client wafregionaliface.WAFRegionalAPI
}

func NewWAFRegionalActivities(client wafregionaliface.WAFRegionalAPI) *WAFRegionalActivities {
    return &WAFRegionalActivities{client: client}
}

func (a *WAFRegionalActivities) AssociateWebACL(input *wafregional.AssociateWebACLInput) (*wafregional.AssociateWebACLOutput, error) {
    return a.client.AssociateWebACL(input)
}

func (a *WAFRegionalActivities) CreateByteMatchSet(input *wafregional.CreateByteMatchSetInput) (*wafregional.CreateByteMatchSetOutput, error) {
    return a.client.CreateByteMatchSet(input)
}

func (a *WAFRegionalActivities) CreateGeoMatchSet(input *wafregional.CreateGeoMatchSetInput) (*wafregional.CreateGeoMatchSetOutput, error) {
    return a.client.CreateGeoMatchSet(input)
}

func (a *WAFRegionalActivities) CreateIPSet(input *wafregional.CreateIPSetInput) (*wafregional.CreateIPSetOutput, error) {
    return a.client.CreateIPSet(input)
}

func (a *WAFRegionalActivities) CreateRateBasedRule(input *wafregional.CreateRateBasedRuleInput) (*wafregional.CreateRateBasedRuleOutput, error) {
    return a.client.CreateRateBasedRule(input)
}

func (a *WAFRegionalActivities) CreateRegexMatchSet(input *wafregional.CreateRegexMatchSetInput) (*wafregional.CreateRegexMatchSetOutput, error) {
    return a.client.CreateRegexMatchSet(input)
}

func (a *WAFRegionalActivities) CreateRegexPatternSet(input *wafregional.CreateRegexPatternSetInput) (*wafregional.CreateRegexPatternSetOutput, error) {
    return a.client.CreateRegexPatternSet(input)
}

func (a *WAFRegionalActivities) CreateRule(input *wafregional.CreateRuleInput) (*wafregional.CreateRuleOutput, error) {
    return a.client.CreateRule(input)
}

func (a *WAFRegionalActivities) CreateRuleGroup(input *wafregional.CreateRuleGroupInput) (*wafregional.CreateRuleGroupOutput, error) {
    return a.client.CreateRuleGroup(input)
}

func (a *WAFRegionalActivities) CreateSizeConstraintSet(input *wafregional.CreateSizeConstraintSetInput) (*wafregional.CreateSizeConstraintSetOutput, error) {
    return a.client.CreateSizeConstraintSet(input)
}

func (a *WAFRegionalActivities) CreateSqlInjectionMatchSet(input *wafregional.CreateSqlInjectionMatchSetInput) (*wafregional.CreateSqlInjectionMatchSetOutput, error) {
    return a.client.CreateSqlInjectionMatchSet(input)
}

func (a *WAFRegionalActivities) CreateWebACL(input *wafregional.CreateWebACLInput) (*wafregional.CreateWebACLOutput, error) {
    return a.client.CreateWebACL(input)
}

func (a *WAFRegionalActivities) CreateWebACLMigrationStack(input *wafregional.CreateWebACLMigrationStackInput) (*wafregional.CreateWebACLMigrationStackOutput, error) {
    return a.client.CreateWebACLMigrationStack(input)
}

func (a *WAFRegionalActivities) CreateXssMatchSet(input *wafregional.CreateXssMatchSetInput) (*wafregional.CreateXssMatchSetOutput, error) {
    return a.client.CreateXssMatchSet(input)
}

func (a *WAFRegionalActivities) DeleteByteMatchSet(input *wafregional.DeleteByteMatchSetInput) (*wafregional.DeleteByteMatchSetOutput, error) {
    return a.client.DeleteByteMatchSet(input)
}

func (a *WAFRegionalActivities) DeleteGeoMatchSet(input *wafregional.DeleteGeoMatchSetInput) (*wafregional.DeleteGeoMatchSetOutput, error) {
    return a.client.DeleteGeoMatchSet(input)
}

func (a *WAFRegionalActivities) DeleteIPSet(input *wafregional.DeleteIPSetInput) (*wafregional.DeleteIPSetOutput, error) {
    return a.client.DeleteIPSet(input)
}

func (a *WAFRegionalActivities) DeleteLoggingConfiguration(input *wafregional.DeleteLoggingConfigurationInput) (*wafregional.DeleteLoggingConfigurationOutput, error) {
    return a.client.DeleteLoggingConfiguration(input)
}

func (a *WAFRegionalActivities) DeletePermissionPolicy(input *wafregional.DeletePermissionPolicyInput) (*wafregional.DeletePermissionPolicyOutput, error) {
    return a.client.DeletePermissionPolicy(input)
}

func (a *WAFRegionalActivities) DeleteRateBasedRule(input *wafregional.DeleteRateBasedRuleInput) (*wafregional.DeleteRateBasedRuleOutput, error) {
    return a.client.DeleteRateBasedRule(input)
}

func (a *WAFRegionalActivities) DeleteRegexMatchSet(input *wafregional.DeleteRegexMatchSetInput) (*wafregional.DeleteRegexMatchSetOutput, error) {
    return a.client.DeleteRegexMatchSet(input)
}

func (a *WAFRegionalActivities) DeleteRegexPatternSet(input *wafregional.DeleteRegexPatternSetInput) (*wafregional.DeleteRegexPatternSetOutput, error) {
    return a.client.DeleteRegexPatternSet(input)
}

func (a *WAFRegionalActivities) DeleteRule(input *wafregional.DeleteRuleInput) (*wafregional.DeleteRuleOutput, error) {
    return a.client.DeleteRule(input)
}

func (a *WAFRegionalActivities) DeleteRuleGroup(input *wafregional.DeleteRuleGroupInput) (*wafregional.DeleteRuleGroupOutput, error) {
    return a.client.DeleteRuleGroup(input)
}

func (a *WAFRegionalActivities) DeleteSizeConstraintSet(input *wafregional.DeleteSizeConstraintSetInput) (*wafregional.DeleteSizeConstraintSetOutput, error) {
    return a.client.DeleteSizeConstraintSet(input)
}

func (a *WAFRegionalActivities) DeleteSqlInjectionMatchSet(input *wafregional.DeleteSqlInjectionMatchSetInput) (*wafregional.DeleteSqlInjectionMatchSetOutput, error) {
    return a.client.DeleteSqlInjectionMatchSet(input)
}

func (a *WAFRegionalActivities) DeleteWebACL(input *wafregional.DeleteWebACLInput) (*wafregional.DeleteWebACLOutput, error) {
    return a.client.DeleteWebACL(input)
}

func (a *WAFRegionalActivities) DeleteXssMatchSet(input *wafregional.DeleteXssMatchSetInput) (*wafregional.DeleteXssMatchSetOutput, error) {
    return a.client.DeleteXssMatchSet(input)
}

func (a *WAFRegionalActivities) DisassociateWebACL(input *wafregional.DisassociateWebACLInput) (*wafregional.DisassociateWebACLOutput, error) {
    return a.client.DisassociateWebACL(input)
}

func (a *WAFRegionalActivities) GetByteMatchSet(input *wafregional.GetByteMatchSetInput) (*wafregional.GetByteMatchSetOutput, error) {
    return a.client.GetByteMatchSet(input)
}

func (a *WAFRegionalActivities) GetChangeToken(input *wafregional.GetChangeTokenInput) (*wafregional.GetChangeTokenOutput, error) {
    return a.client.GetChangeToken(input)
}

func (a *WAFRegionalActivities) GetChangeTokenStatus(input *wafregional.GetChangeTokenStatusInput) (*wafregional.GetChangeTokenStatusOutput, error) {
    return a.client.GetChangeTokenStatus(input)
}

func (a *WAFRegionalActivities) GetGeoMatchSet(input *wafregional.GetGeoMatchSetInput) (*wafregional.GetGeoMatchSetOutput, error) {
    return a.client.GetGeoMatchSet(input)
}

func (a *WAFRegionalActivities) GetIPSet(input *wafregional.GetIPSetInput) (*wafregional.GetIPSetOutput, error) {
    return a.client.GetIPSet(input)
}

func (a *WAFRegionalActivities) GetLoggingConfiguration(input *wafregional.GetLoggingConfigurationInput) (*wafregional.GetLoggingConfigurationOutput, error) {
    return a.client.GetLoggingConfiguration(input)
}

func (a *WAFRegionalActivities) GetPermissionPolicy(input *wafregional.GetPermissionPolicyInput) (*wafregional.GetPermissionPolicyOutput, error) {
    return a.client.GetPermissionPolicy(input)
}

func (a *WAFRegionalActivities) GetRateBasedRule(input *wafregional.GetRateBasedRuleInput) (*wafregional.GetRateBasedRuleOutput, error) {
    return a.client.GetRateBasedRule(input)
}

func (a *WAFRegionalActivities) GetRateBasedRuleManagedKeys(input *wafregional.GetRateBasedRuleManagedKeysInput) (*wafregional.GetRateBasedRuleManagedKeysOutput, error) {
    return a.client.GetRateBasedRuleManagedKeys(input)
}

func (a *WAFRegionalActivities) GetRegexMatchSet(input *wafregional.GetRegexMatchSetInput) (*wafregional.GetRegexMatchSetOutput, error) {
    return a.client.GetRegexMatchSet(input)
}

func (a *WAFRegionalActivities) GetRegexPatternSet(input *wafregional.GetRegexPatternSetInput) (*wafregional.GetRegexPatternSetOutput, error) {
    return a.client.GetRegexPatternSet(input)
}

func (a *WAFRegionalActivities) GetRule(input *wafregional.GetRuleInput) (*wafregional.GetRuleOutput, error) {
    return a.client.GetRule(input)
}

func (a *WAFRegionalActivities) GetRuleGroup(input *wafregional.GetRuleGroupInput) (*wafregional.GetRuleGroupOutput, error) {
    return a.client.GetRuleGroup(input)
}

func (a *WAFRegionalActivities) GetSampledRequests(input *wafregional.GetSampledRequestsInput) (*wafregional.GetSampledRequestsOutput, error) {
    return a.client.GetSampledRequests(input)
}

func (a *WAFRegionalActivities) GetSizeConstraintSet(input *wafregional.GetSizeConstraintSetInput) (*wafregional.GetSizeConstraintSetOutput, error) {
    return a.client.GetSizeConstraintSet(input)
}

func (a *WAFRegionalActivities) GetSqlInjectionMatchSet(input *wafregional.GetSqlInjectionMatchSetInput) (*wafregional.GetSqlInjectionMatchSetOutput, error) {
    return a.client.GetSqlInjectionMatchSet(input)
}

func (a *WAFRegionalActivities) GetWebACL(input *wafregional.GetWebACLInput) (*wafregional.GetWebACLOutput, error) {
    return a.client.GetWebACL(input)
}

func (a *WAFRegionalActivities) GetWebACLForResource(input *wafregional.GetWebACLForResourceInput) (*wafregional.GetWebACLForResourceOutput, error) {
    return a.client.GetWebACLForResource(input)
}

func (a *WAFRegionalActivities) GetXssMatchSet(input *wafregional.GetXssMatchSetInput) (*wafregional.GetXssMatchSetOutput, error) {
    return a.client.GetXssMatchSet(input)
}

func (a *WAFRegionalActivities) ListActivatedRulesInRuleGroup(input *wafregional.ListActivatedRulesInRuleGroupInput) (*wafregional.ListActivatedRulesInRuleGroupOutput, error) {
    return a.client.ListActivatedRulesInRuleGroup(input)
}

func (a *WAFRegionalActivities) ListByteMatchSets(input *wafregional.ListByteMatchSetsInput) (*wafregional.ListByteMatchSetsOutput, error) {
    return a.client.ListByteMatchSets(input)
}

func (a *WAFRegionalActivities) ListGeoMatchSets(input *wafregional.ListGeoMatchSetsInput) (*wafregional.ListGeoMatchSetsOutput, error) {
    return a.client.ListGeoMatchSets(input)
}

func (a *WAFRegionalActivities) ListIPSets(input *wafregional.ListIPSetsInput) (*wafregional.ListIPSetsOutput, error) {
    return a.client.ListIPSets(input)
}

func (a *WAFRegionalActivities) ListLoggingConfigurations(input *wafregional.ListLoggingConfigurationsInput) (*wafregional.ListLoggingConfigurationsOutput, error) {
    return a.client.ListLoggingConfigurations(input)
}

func (a *WAFRegionalActivities) ListRateBasedRules(input *wafregional.ListRateBasedRulesInput) (*wafregional.ListRateBasedRulesOutput, error) {
    return a.client.ListRateBasedRules(input)
}

func (a *WAFRegionalActivities) ListRegexMatchSets(input *wafregional.ListRegexMatchSetsInput) (*wafregional.ListRegexMatchSetsOutput, error) {
    return a.client.ListRegexMatchSets(input)
}

func (a *WAFRegionalActivities) ListRegexPatternSets(input *wafregional.ListRegexPatternSetsInput) (*wafregional.ListRegexPatternSetsOutput, error) {
    return a.client.ListRegexPatternSets(input)
}

func (a *WAFRegionalActivities) ListResourcesForWebACL(input *wafregional.ListResourcesForWebACLInput) (*wafregional.ListResourcesForWebACLOutput, error) {
    return a.client.ListResourcesForWebACL(input)
}

func (a *WAFRegionalActivities) ListRuleGroups(input *wafregional.ListRuleGroupsInput) (*wafregional.ListRuleGroupsOutput, error) {
    return a.client.ListRuleGroups(input)
}

func (a *WAFRegionalActivities) ListRules(input *wafregional.ListRulesInput) (*wafregional.ListRulesOutput, error) {
    return a.client.ListRules(input)
}

func (a *WAFRegionalActivities) ListSizeConstraintSets(input *wafregional.ListSizeConstraintSetsInput) (*wafregional.ListSizeConstraintSetsOutput, error) {
    return a.client.ListSizeConstraintSets(input)
}

func (a *WAFRegionalActivities) ListSqlInjectionMatchSets(input *wafregional.ListSqlInjectionMatchSetsInput) (*wafregional.ListSqlInjectionMatchSetsOutput, error) {
    return a.client.ListSqlInjectionMatchSets(input)
}

func (a *WAFRegionalActivities) ListSubscribedRuleGroups(input *wafregional.ListSubscribedRuleGroupsInput) (*wafregional.ListSubscribedRuleGroupsOutput, error) {
    return a.client.ListSubscribedRuleGroups(input)
}

func (a *WAFRegionalActivities) ListTagsForResource(input *wafregional.ListTagsForResourceInput) (*wafregional.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *WAFRegionalActivities) ListWebACLs(input *wafregional.ListWebACLsInput) (*wafregional.ListWebACLsOutput, error) {
    return a.client.ListWebACLs(input)
}

func (a *WAFRegionalActivities) ListXssMatchSets(input *wafregional.ListXssMatchSetsInput) (*wafregional.ListXssMatchSetsOutput, error) {
    return a.client.ListXssMatchSets(input)
}

func (a *WAFRegionalActivities) PutLoggingConfiguration(input *wafregional.PutLoggingConfigurationInput) (*wafregional.PutLoggingConfigurationOutput, error) {
    return a.client.PutLoggingConfiguration(input)
}

func (a *WAFRegionalActivities) PutPermissionPolicy(input *wafregional.PutPermissionPolicyInput) (*wafregional.PutPermissionPolicyOutput, error) {
    return a.client.PutPermissionPolicy(input)
}

func (a *WAFRegionalActivities) TagResource(input *wafregional.TagResourceInput) (*wafregional.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *WAFRegionalActivities) UntagResource(input *wafregional.UntagResourceInput) (*wafregional.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *WAFRegionalActivities) UpdateByteMatchSet(input *wafregional.UpdateByteMatchSetInput) (*wafregional.UpdateByteMatchSetOutput, error) {
    return a.client.UpdateByteMatchSet(input)
}

func (a *WAFRegionalActivities) UpdateGeoMatchSet(input *wafregional.UpdateGeoMatchSetInput) (*wafregional.UpdateGeoMatchSetOutput, error) {
    return a.client.UpdateGeoMatchSet(input)
}

func (a *WAFRegionalActivities) UpdateIPSet(input *wafregional.UpdateIPSetInput) (*wafregional.UpdateIPSetOutput, error) {
    return a.client.UpdateIPSet(input)
}

func (a *WAFRegionalActivities) UpdateRateBasedRule(input *wafregional.UpdateRateBasedRuleInput) (*wafregional.UpdateRateBasedRuleOutput, error) {
    return a.client.UpdateRateBasedRule(input)
}

func (a *WAFRegionalActivities) UpdateRegexMatchSet(input *wafregional.UpdateRegexMatchSetInput) (*wafregional.UpdateRegexMatchSetOutput, error) {
    return a.client.UpdateRegexMatchSet(input)
}

func (a *WAFRegionalActivities) UpdateRegexPatternSet(input *wafregional.UpdateRegexPatternSetInput) (*wafregional.UpdateRegexPatternSetOutput, error) {
    return a.client.UpdateRegexPatternSet(input)
}

func (a *WAFRegionalActivities) UpdateRule(input *wafregional.UpdateRuleInput) (*wafregional.UpdateRuleOutput, error) {
    return a.client.UpdateRule(input)
}

func (a *WAFRegionalActivities) UpdateRuleGroup(input *wafregional.UpdateRuleGroupInput) (*wafregional.UpdateRuleGroupOutput, error) {
    return a.client.UpdateRuleGroup(input)
}

func (a *WAFRegionalActivities) UpdateSizeConstraintSet(input *wafregional.UpdateSizeConstraintSetInput) (*wafregional.UpdateSizeConstraintSetOutput, error) {
    return a.client.UpdateSizeConstraintSet(input)
}

func (a *WAFRegionalActivities) UpdateSqlInjectionMatchSet(input *wafregional.UpdateSqlInjectionMatchSetInput) (*wafregional.UpdateSqlInjectionMatchSetOutput, error) {
    return a.client.UpdateSqlInjectionMatchSet(input)
}

func (a *WAFRegionalActivities) UpdateWebACL(input *wafregional.UpdateWebACLInput) (*wafregional.UpdateWebACLOutput, error) {
    return a.client.UpdateWebACL(input)
}

func (a *WAFRegionalActivities) UpdateXssMatchSet(input *wafregional.UpdateXssMatchSetInput) (*wafregional.UpdateXssMatchSetOutput, error) {
    return a.client.UpdateXssMatchSet(input)
}
