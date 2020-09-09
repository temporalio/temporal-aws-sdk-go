package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/aws/aws-sdk-go/service/wafregional"
	"github.com/aws/aws-sdk-go/service/wafregional/wafregionaliface"
)

type WAFRegionalActivities struct {
	client wafregionaliface.WAFRegionalAPI
}

func NewWAFRegionalActivities(session *session.Session, config ...*aws.Config) *WAFRegionalActivities {
	client := wafregional.New(session, config...)
	return &WAFRegionalActivities{client: client}
}

func (a *WAFRegionalActivities) AssociateWebACL(input *wafregional.AssociateWebACLInput) (*wafregional.AssociateWebACLOutput, error) {
	return a.client.AssociateWebACL(input)
}

func (a *WAFRegionalActivities) CreateByteMatchSet(input *waf.CreateByteMatchSetInput) (*waf.CreateByteMatchSetOutput, error) {
	return a.client.CreateByteMatchSet(input)
}

func (a *WAFRegionalActivities) CreateGeoMatchSet(input *waf.CreateGeoMatchSetInput) (*waf.CreateGeoMatchSetOutput, error) {
	return a.client.CreateGeoMatchSet(input)
}

func (a *WAFRegionalActivities) CreateIPSet(input *waf.CreateIPSetInput) (*waf.CreateIPSetOutput, error) {
	return a.client.CreateIPSet(input)
}

func (a *WAFRegionalActivities) CreateRateBasedRule(input *waf.CreateRateBasedRuleInput) (*waf.CreateRateBasedRuleOutput, error) {
	return a.client.CreateRateBasedRule(input)
}

func (a *WAFRegionalActivities) CreateRegexMatchSet(input *waf.CreateRegexMatchSetInput) (*waf.CreateRegexMatchSetOutput, error) {
	return a.client.CreateRegexMatchSet(input)
}

func (a *WAFRegionalActivities) CreateRegexPatternSet(input *waf.CreateRegexPatternSetInput) (*waf.CreateRegexPatternSetOutput, error) {
	return a.client.CreateRegexPatternSet(input)
}

func (a *WAFRegionalActivities) CreateRule(input *waf.CreateRuleInput) (*waf.CreateRuleOutput, error) {
	return a.client.CreateRule(input)
}

func (a *WAFRegionalActivities) CreateRuleGroup(input *waf.CreateRuleGroupInput) (*waf.CreateRuleGroupOutput, error) {
	return a.client.CreateRuleGroup(input)
}

func (a *WAFRegionalActivities) CreateSizeConstraintSet(input *waf.CreateSizeConstraintSetInput) (*waf.CreateSizeConstraintSetOutput, error) {
	return a.client.CreateSizeConstraintSet(input)
}

func (a *WAFRegionalActivities) CreateSqlInjectionMatchSet(input *waf.CreateSqlInjectionMatchSetInput) (*waf.CreateSqlInjectionMatchSetOutput, error) {
	return a.client.CreateSqlInjectionMatchSet(input)
}

func (a *WAFRegionalActivities) CreateWebACL(input *waf.CreateWebACLInput) (*waf.CreateWebACLOutput, error) {
	return a.client.CreateWebACL(input)
}

func (a *WAFRegionalActivities) CreateWebACLMigrationStack(input *waf.CreateWebACLMigrationStackInput) (*waf.CreateWebACLMigrationStackOutput, error) {
	return a.client.CreateWebACLMigrationStack(input)
}

func (a *WAFRegionalActivities) CreateXssMatchSet(input *waf.CreateXssMatchSetInput) (*waf.CreateXssMatchSetOutput, error) {
	return a.client.CreateXssMatchSet(input)
}

func (a *WAFRegionalActivities) DeleteByteMatchSet(input *waf.DeleteByteMatchSetInput) (*waf.DeleteByteMatchSetOutput, error) {
	return a.client.DeleteByteMatchSet(input)
}

func (a *WAFRegionalActivities) DeleteGeoMatchSet(input *waf.DeleteGeoMatchSetInput) (*waf.DeleteGeoMatchSetOutput, error) {
	return a.client.DeleteGeoMatchSet(input)
}

func (a *WAFRegionalActivities) DeleteIPSet(input *waf.DeleteIPSetInput) (*waf.DeleteIPSetOutput, error) {
	return a.client.DeleteIPSet(input)
}

func (a *WAFRegionalActivities) DeleteLoggingConfiguration(input *waf.DeleteLoggingConfigurationInput) (*waf.DeleteLoggingConfigurationOutput, error) {
	return a.client.DeleteLoggingConfiguration(input)
}

func (a *WAFRegionalActivities) DeletePermissionPolicy(input *waf.DeletePermissionPolicyInput) (*waf.DeletePermissionPolicyOutput, error) {
	return a.client.DeletePermissionPolicy(input)
}

func (a *WAFRegionalActivities) DeleteRateBasedRule(input *waf.DeleteRateBasedRuleInput) (*waf.DeleteRateBasedRuleOutput, error) {
	return a.client.DeleteRateBasedRule(input)
}

func (a *WAFRegionalActivities) DeleteRegexMatchSet(input *waf.DeleteRegexMatchSetInput) (*waf.DeleteRegexMatchSetOutput, error) {
	return a.client.DeleteRegexMatchSet(input)
}

func (a *WAFRegionalActivities) DeleteRegexPatternSet(input *waf.DeleteRegexPatternSetInput) (*waf.DeleteRegexPatternSetOutput, error) {
	return a.client.DeleteRegexPatternSet(input)
}

func (a *WAFRegionalActivities) DeleteRule(input *waf.DeleteRuleInput) (*waf.DeleteRuleOutput, error) {
	return a.client.DeleteRule(input)
}

func (a *WAFRegionalActivities) DeleteRuleGroup(input *waf.DeleteRuleGroupInput) (*waf.DeleteRuleGroupOutput, error) {
	return a.client.DeleteRuleGroup(input)
}

func (a *WAFRegionalActivities) DeleteSizeConstraintSet(input *waf.DeleteSizeConstraintSetInput) (*waf.DeleteSizeConstraintSetOutput, error) {
	return a.client.DeleteSizeConstraintSet(input)
}

func (a *WAFRegionalActivities) DeleteSqlInjectionMatchSet(input *waf.DeleteSqlInjectionMatchSetInput) (*waf.DeleteSqlInjectionMatchSetOutput, error) {
	return a.client.DeleteSqlInjectionMatchSet(input)
}

func (a *WAFRegionalActivities) DeleteWebACL(input *waf.DeleteWebACLInput) (*waf.DeleteWebACLOutput, error) {
	return a.client.DeleteWebACL(input)
}

func (a *WAFRegionalActivities) DeleteXssMatchSet(input *waf.DeleteXssMatchSetInput) (*waf.DeleteXssMatchSetOutput, error) {
	return a.client.DeleteXssMatchSet(input)
}

func (a *WAFRegionalActivities) DisassociateWebACL(input *wafregional.DisassociateWebACLInput) (*wafregional.DisassociateWebACLOutput, error) {
	return a.client.DisassociateWebACL(input)
}

func (a *WAFRegionalActivities) GetByteMatchSet(input *waf.GetByteMatchSetInput) (*waf.GetByteMatchSetOutput, error) {
	return a.client.GetByteMatchSet(input)
}

func (a *WAFRegionalActivities) GetChangeToken(input *waf.GetChangeTokenInput) (*waf.GetChangeTokenOutput, error) {
	return a.client.GetChangeToken(input)
}

func (a *WAFRegionalActivities) GetChangeTokenStatus(input *waf.GetChangeTokenStatusInput) (*waf.GetChangeTokenStatusOutput, error) {
	return a.client.GetChangeTokenStatus(input)
}

func (a *WAFRegionalActivities) GetGeoMatchSet(input *waf.GetGeoMatchSetInput) (*waf.GetGeoMatchSetOutput, error) {
	return a.client.GetGeoMatchSet(input)
}

func (a *WAFRegionalActivities) GetIPSet(input *waf.GetIPSetInput) (*waf.GetIPSetOutput, error) {
	return a.client.GetIPSet(input)
}

func (a *WAFRegionalActivities) GetLoggingConfiguration(input *waf.GetLoggingConfigurationInput) (*waf.GetLoggingConfigurationOutput, error) {
	return a.client.GetLoggingConfiguration(input)
}

func (a *WAFRegionalActivities) GetPermissionPolicy(input *waf.GetPermissionPolicyInput) (*waf.GetPermissionPolicyOutput, error) {
	return a.client.GetPermissionPolicy(input)
}

func (a *WAFRegionalActivities) GetRateBasedRule(input *waf.GetRateBasedRuleInput) (*waf.GetRateBasedRuleOutput, error) {
	return a.client.GetRateBasedRule(input)
}

func (a *WAFRegionalActivities) GetRateBasedRuleManagedKeys(input *waf.GetRateBasedRuleManagedKeysInput) (*waf.GetRateBasedRuleManagedKeysOutput, error) {
	return a.client.GetRateBasedRuleManagedKeys(input)
}

func (a *WAFRegionalActivities) GetRegexMatchSet(input *waf.GetRegexMatchSetInput) (*waf.GetRegexMatchSetOutput, error) {
	return a.client.GetRegexMatchSet(input)
}

func (a *WAFRegionalActivities) GetRegexPatternSet(input *waf.GetRegexPatternSetInput) (*waf.GetRegexPatternSetOutput, error) {
	return a.client.GetRegexPatternSet(input)
}

func (a *WAFRegionalActivities) GetRule(input *waf.GetRuleInput) (*waf.GetRuleOutput, error) {
	return a.client.GetRule(input)
}

func (a *WAFRegionalActivities) GetRuleGroup(input *waf.GetRuleGroupInput) (*waf.GetRuleGroupOutput, error) {
	return a.client.GetRuleGroup(input)
}

func (a *WAFRegionalActivities) GetSampledRequests(input *waf.GetSampledRequestsInput) (*waf.GetSampledRequestsOutput, error) {
	return a.client.GetSampledRequests(input)
}

func (a *WAFRegionalActivities) GetSizeConstraintSet(input *waf.GetSizeConstraintSetInput) (*waf.GetSizeConstraintSetOutput, error) {
	return a.client.GetSizeConstraintSet(input)
}

func (a *WAFRegionalActivities) GetSqlInjectionMatchSet(input *waf.GetSqlInjectionMatchSetInput) (*waf.GetSqlInjectionMatchSetOutput, error) {
	return a.client.GetSqlInjectionMatchSet(input)
}

func (a *WAFRegionalActivities) GetWebACL(input *waf.GetWebACLInput) (*waf.GetWebACLOutput, error) {
	return a.client.GetWebACL(input)
}

func (a *WAFRegionalActivities) GetWebACLForResource(input *wafregional.GetWebACLForResourceInput) (*wafregional.GetWebACLForResourceOutput, error) {
	return a.client.GetWebACLForResource(input)
}

func (a *WAFRegionalActivities) GetXssMatchSet(input *waf.GetXssMatchSetInput) (*waf.GetXssMatchSetOutput, error) {
	return a.client.GetXssMatchSet(input)
}

func (a *WAFRegionalActivities) ListActivatedRulesInRuleGroup(input *waf.ListActivatedRulesInRuleGroupInput) (*waf.ListActivatedRulesInRuleGroupOutput, error) {
	return a.client.ListActivatedRulesInRuleGroup(input)
}

func (a *WAFRegionalActivities) ListByteMatchSets(input *waf.ListByteMatchSetsInput) (*waf.ListByteMatchSetsOutput, error) {
	return a.client.ListByteMatchSets(input)
}

func (a *WAFRegionalActivities) ListGeoMatchSets(input *waf.ListGeoMatchSetsInput) (*waf.ListGeoMatchSetsOutput, error) {
	return a.client.ListGeoMatchSets(input)
}

func (a *WAFRegionalActivities) ListIPSets(input *waf.ListIPSetsInput) (*waf.ListIPSetsOutput, error) {
	return a.client.ListIPSets(input)
}

func (a *WAFRegionalActivities) ListLoggingConfigurations(input *waf.ListLoggingConfigurationsInput) (*waf.ListLoggingConfigurationsOutput, error) {
	return a.client.ListLoggingConfigurations(input)
}

func (a *WAFRegionalActivities) ListRateBasedRules(input *waf.ListRateBasedRulesInput) (*waf.ListRateBasedRulesOutput, error) {
	return a.client.ListRateBasedRules(input)
}

func (a *WAFRegionalActivities) ListRegexMatchSets(input *waf.ListRegexMatchSetsInput) (*waf.ListRegexMatchSetsOutput, error) {
	return a.client.ListRegexMatchSets(input)
}

func (a *WAFRegionalActivities) ListRegexPatternSets(input *waf.ListRegexPatternSetsInput) (*waf.ListRegexPatternSetsOutput, error) {
	return a.client.ListRegexPatternSets(input)
}

func (a *WAFRegionalActivities) ListResourcesForWebACL(input *wafregional.ListResourcesForWebACLInput) (*wafregional.ListResourcesForWebACLOutput, error) {
	return a.client.ListResourcesForWebACL(input)
}

func (a *WAFRegionalActivities) ListRuleGroups(input *waf.ListRuleGroupsInput) (*waf.ListRuleGroupsOutput, error) {
	return a.client.ListRuleGroups(input)
}

func (a *WAFRegionalActivities) ListRules(input *waf.ListRulesInput) (*waf.ListRulesOutput, error) {
	return a.client.ListRules(input)
}

func (a *WAFRegionalActivities) ListSizeConstraintSets(input *waf.ListSizeConstraintSetsInput) (*waf.ListSizeConstraintSetsOutput, error) {
	return a.client.ListSizeConstraintSets(input)
}

func (a *WAFRegionalActivities) ListSqlInjectionMatchSets(input *waf.ListSqlInjectionMatchSetsInput) (*waf.ListSqlInjectionMatchSetsOutput, error) {
	return a.client.ListSqlInjectionMatchSets(input)
}

func (a *WAFRegionalActivities) ListSubscribedRuleGroups(input *waf.ListSubscribedRuleGroupsInput) (*waf.ListSubscribedRuleGroupsOutput, error) {
	return a.client.ListSubscribedRuleGroups(input)
}

func (a *WAFRegionalActivities) ListTagsForResource(input *waf.ListTagsForResourceInput) (*waf.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResource(input)
}

func (a *WAFRegionalActivities) ListWebACLs(input *waf.ListWebACLsInput) (*waf.ListWebACLsOutput, error) {
	return a.client.ListWebACLs(input)
}

func (a *WAFRegionalActivities) ListXssMatchSets(input *waf.ListXssMatchSetsInput) (*waf.ListXssMatchSetsOutput, error) {
	return a.client.ListXssMatchSets(input)
}

func (a *WAFRegionalActivities) PutLoggingConfiguration(input *waf.PutLoggingConfigurationInput) (*waf.PutLoggingConfigurationOutput, error) {
	return a.client.PutLoggingConfiguration(input)
}

func (a *WAFRegionalActivities) PutPermissionPolicy(input *waf.PutPermissionPolicyInput) (*waf.PutPermissionPolicyOutput, error) {
	return a.client.PutPermissionPolicy(input)
}

func (a *WAFRegionalActivities) TagResource(input *waf.TagResourceInput) (*waf.TagResourceOutput, error) {
	return a.client.TagResource(input)
}

func (a *WAFRegionalActivities) UntagResource(input *waf.UntagResourceInput) (*waf.UntagResourceOutput, error) {
	return a.client.UntagResource(input)
}

func (a *WAFRegionalActivities) UpdateByteMatchSet(input *waf.UpdateByteMatchSetInput) (*waf.UpdateByteMatchSetOutput, error) {
	return a.client.UpdateByteMatchSet(input)
}

func (a *WAFRegionalActivities) UpdateGeoMatchSet(input *waf.UpdateGeoMatchSetInput) (*waf.UpdateGeoMatchSetOutput, error) {
	return a.client.UpdateGeoMatchSet(input)
}

func (a *WAFRegionalActivities) UpdateIPSet(input *waf.UpdateIPSetInput) (*waf.UpdateIPSetOutput, error) {
	return a.client.UpdateIPSet(input)
}

func (a *WAFRegionalActivities) UpdateRateBasedRule(input *waf.UpdateRateBasedRuleInput) (*waf.UpdateRateBasedRuleOutput, error) {
	return a.client.UpdateRateBasedRule(input)
}

func (a *WAFRegionalActivities) UpdateRegexMatchSet(input *waf.UpdateRegexMatchSetInput) (*waf.UpdateRegexMatchSetOutput, error) {
	return a.client.UpdateRegexMatchSet(input)
}

func (a *WAFRegionalActivities) UpdateRegexPatternSet(input *waf.UpdateRegexPatternSetInput) (*waf.UpdateRegexPatternSetOutput, error) {
	return a.client.UpdateRegexPatternSet(input)
}

func (a *WAFRegionalActivities) UpdateRule(input *waf.UpdateRuleInput) (*waf.UpdateRuleOutput, error) {
	return a.client.UpdateRule(input)
}

func (a *WAFRegionalActivities) UpdateRuleGroup(input *waf.UpdateRuleGroupInput) (*waf.UpdateRuleGroupOutput, error) {
	return a.client.UpdateRuleGroup(input)
}

func (a *WAFRegionalActivities) UpdateSizeConstraintSet(input *waf.UpdateSizeConstraintSetInput) (*waf.UpdateSizeConstraintSetOutput, error) {
	return a.client.UpdateSizeConstraintSet(input)
}

func (a *WAFRegionalActivities) UpdateSqlInjectionMatchSet(input *waf.UpdateSqlInjectionMatchSetInput) (*waf.UpdateSqlInjectionMatchSetOutput, error) {
	return a.client.UpdateSqlInjectionMatchSet(input)
}

func (a *WAFRegionalActivities) UpdateWebACL(input *waf.UpdateWebACLInput) (*waf.UpdateWebACLOutput, error) {
	return a.client.UpdateWebACL(input)
}

func (a *WAFRegionalActivities) UpdateXssMatchSet(input *waf.UpdateXssMatchSetInput) (*waf.UpdateXssMatchSetOutput, error) {
	return a.client.UpdateXssMatchSet(input)
}
