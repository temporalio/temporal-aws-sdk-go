package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/aws/aws-sdk-go/service/waf/wafiface"
)

type WAFActivities struct {
	client wafiface.WAFAPI
}

func NewWAFActivities(session *session.Session, config ...*aws.Config) *WAFActivities {
	client := waf.New(session, config...)
	return &WAFActivities{client: client}
}

func (a *WAFActivities) CreateByteMatchSet(input *waf.CreateByteMatchSetInput) (*waf.CreateByteMatchSetOutput, error) {
	return a.client.CreateByteMatchSet(input)
}

func (a *WAFActivities) CreateGeoMatchSet(input *waf.CreateGeoMatchSetInput) (*waf.CreateGeoMatchSetOutput, error) {
	return a.client.CreateGeoMatchSet(input)
}

func (a *WAFActivities) CreateIPSet(input *waf.CreateIPSetInput) (*waf.CreateIPSetOutput, error) {
	return a.client.CreateIPSet(input)
}

func (a *WAFActivities) CreateRateBasedRule(input *waf.CreateRateBasedRuleInput) (*waf.CreateRateBasedRuleOutput, error) {
	return a.client.CreateRateBasedRule(input)
}

func (a *WAFActivities) CreateRegexMatchSet(input *waf.CreateRegexMatchSetInput) (*waf.CreateRegexMatchSetOutput, error) {
	return a.client.CreateRegexMatchSet(input)
}

func (a *WAFActivities) CreateRegexPatternSet(input *waf.CreateRegexPatternSetInput) (*waf.CreateRegexPatternSetOutput, error) {
	return a.client.CreateRegexPatternSet(input)
}

func (a *WAFActivities) CreateRule(input *waf.CreateRuleInput) (*waf.CreateRuleOutput, error) {
	return a.client.CreateRule(input)
}

func (a *WAFActivities) CreateRuleGroup(input *waf.CreateRuleGroupInput) (*waf.CreateRuleGroupOutput, error) {
	return a.client.CreateRuleGroup(input)
}

func (a *WAFActivities) CreateSizeConstraintSet(input *waf.CreateSizeConstraintSetInput) (*waf.CreateSizeConstraintSetOutput, error) {
	return a.client.CreateSizeConstraintSet(input)
}

func (a *WAFActivities) CreateSqlInjectionMatchSet(input *waf.CreateSqlInjectionMatchSetInput) (*waf.CreateSqlInjectionMatchSetOutput, error) {
	return a.client.CreateSqlInjectionMatchSet(input)
}

func (a *WAFActivities) CreateWebACL(input *waf.CreateWebACLInput) (*waf.CreateWebACLOutput, error) {
	return a.client.CreateWebACL(input)
}

func (a *WAFActivities) CreateWebACLMigrationStack(input *waf.CreateWebACLMigrationStackInput) (*waf.CreateWebACLMigrationStackOutput, error) {
	return a.client.CreateWebACLMigrationStack(input)
}

func (a *WAFActivities) CreateXssMatchSet(input *waf.CreateXssMatchSetInput) (*waf.CreateXssMatchSetOutput, error) {
	return a.client.CreateXssMatchSet(input)
}

func (a *WAFActivities) DeleteByteMatchSet(input *waf.DeleteByteMatchSetInput) (*waf.DeleteByteMatchSetOutput, error) {
	return a.client.DeleteByteMatchSet(input)
}

func (a *WAFActivities) DeleteGeoMatchSet(input *waf.DeleteGeoMatchSetInput) (*waf.DeleteGeoMatchSetOutput, error) {
	return a.client.DeleteGeoMatchSet(input)
}

func (a *WAFActivities) DeleteIPSet(input *waf.DeleteIPSetInput) (*waf.DeleteIPSetOutput, error) {
	return a.client.DeleteIPSet(input)
}

func (a *WAFActivities) DeleteLoggingConfiguration(input *waf.DeleteLoggingConfigurationInput) (*waf.DeleteLoggingConfigurationOutput, error) {
	return a.client.DeleteLoggingConfiguration(input)
}

func (a *WAFActivities) DeletePermissionPolicy(input *waf.DeletePermissionPolicyInput) (*waf.DeletePermissionPolicyOutput, error) {
	return a.client.DeletePermissionPolicy(input)
}

func (a *WAFActivities) DeleteRateBasedRule(input *waf.DeleteRateBasedRuleInput) (*waf.DeleteRateBasedRuleOutput, error) {
	return a.client.DeleteRateBasedRule(input)
}

func (a *WAFActivities) DeleteRegexMatchSet(input *waf.DeleteRegexMatchSetInput) (*waf.DeleteRegexMatchSetOutput, error) {
	return a.client.DeleteRegexMatchSet(input)
}

func (a *WAFActivities) DeleteRegexPatternSet(input *waf.DeleteRegexPatternSetInput) (*waf.DeleteRegexPatternSetOutput, error) {
	return a.client.DeleteRegexPatternSet(input)
}

func (a *WAFActivities) DeleteRule(input *waf.DeleteRuleInput) (*waf.DeleteRuleOutput, error) {
	return a.client.DeleteRule(input)
}

func (a *WAFActivities) DeleteRuleGroup(input *waf.DeleteRuleGroupInput) (*waf.DeleteRuleGroupOutput, error) {
	return a.client.DeleteRuleGroup(input)
}

func (a *WAFActivities) DeleteSizeConstraintSet(input *waf.DeleteSizeConstraintSetInput) (*waf.DeleteSizeConstraintSetOutput, error) {
	return a.client.DeleteSizeConstraintSet(input)
}

func (a *WAFActivities) DeleteSqlInjectionMatchSet(input *waf.DeleteSqlInjectionMatchSetInput) (*waf.DeleteSqlInjectionMatchSetOutput, error) {
	return a.client.DeleteSqlInjectionMatchSet(input)
}

func (a *WAFActivities) DeleteWebACL(input *waf.DeleteWebACLInput) (*waf.DeleteWebACLOutput, error) {
	return a.client.DeleteWebACL(input)
}

func (a *WAFActivities) DeleteXssMatchSet(input *waf.DeleteXssMatchSetInput) (*waf.DeleteXssMatchSetOutput, error) {
	return a.client.DeleteXssMatchSet(input)
}

func (a *WAFActivities) GetByteMatchSet(input *waf.GetByteMatchSetInput) (*waf.GetByteMatchSetOutput, error) {
	return a.client.GetByteMatchSet(input)
}

func (a *WAFActivities) GetChangeToken(input *waf.GetChangeTokenInput) (*waf.GetChangeTokenOutput, error) {
	return a.client.GetChangeToken(input)
}

func (a *WAFActivities) GetChangeTokenStatus(input *waf.GetChangeTokenStatusInput) (*waf.GetChangeTokenStatusOutput, error) {
	return a.client.GetChangeTokenStatus(input)
}

func (a *WAFActivities) GetGeoMatchSet(input *waf.GetGeoMatchSetInput) (*waf.GetGeoMatchSetOutput, error) {
	return a.client.GetGeoMatchSet(input)
}

func (a *WAFActivities) GetIPSet(input *waf.GetIPSetInput) (*waf.GetIPSetOutput, error) {
	return a.client.GetIPSet(input)
}

func (a *WAFActivities) GetLoggingConfiguration(input *waf.GetLoggingConfigurationInput) (*waf.GetLoggingConfigurationOutput, error) {
	return a.client.GetLoggingConfiguration(input)
}

func (a *WAFActivities) GetPermissionPolicy(input *waf.GetPermissionPolicyInput) (*waf.GetPermissionPolicyOutput, error) {
	return a.client.GetPermissionPolicy(input)
}

func (a *WAFActivities) GetRateBasedRule(input *waf.GetRateBasedRuleInput) (*waf.GetRateBasedRuleOutput, error) {
	return a.client.GetRateBasedRule(input)
}

func (a *WAFActivities) GetRateBasedRuleManagedKeys(input *waf.GetRateBasedRuleManagedKeysInput) (*waf.GetRateBasedRuleManagedKeysOutput, error) {
	return a.client.GetRateBasedRuleManagedKeys(input)
}

func (a *WAFActivities) GetRegexMatchSet(input *waf.GetRegexMatchSetInput) (*waf.GetRegexMatchSetOutput, error) {
	return a.client.GetRegexMatchSet(input)
}

func (a *WAFActivities) GetRegexPatternSet(input *waf.GetRegexPatternSetInput) (*waf.GetRegexPatternSetOutput, error) {
	return a.client.GetRegexPatternSet(input)
}

func (a *WAFActivities) GetRule(input *waf.GetRuleInput) (*waf.GetRuleOutput, error) {
	return a.client.GetRule(input)
}

func (a *WAFActivities) GetRuleGroup(input *waf.GetRuleGroupInput) (*waf.GetRuleGroupOutput, error) {
	return a.client.GetRuleGroup(input)
}

func (a *WAFActivities) GetSampledRequests(input *waf.GetSampledRequestsInput) (*waf.GetSampledRequestsOutput, error) {
	return a.client.GetSampledRequests(input)
}

func (a *WAFActivities) GetSizeConstraintSet(input *waf.GetSizeConstraintSetInput) (*waf.GetSizeConstraintSetOutput, error) {
	return a.client.GetSizeConstraintSet(input)
}

func (a *WAFActivities) GetSqlInjectionMatchSet(input *waf.GetSqlInjectionMatchSetInput) (*waf.GetSqlInjectionMatchSetOutput, error) {
	return a.client.GetSqlInjectionMatchSet(input)
}

func (a *WAFActivities) GetWebACL(input *waf.GetWebACLInput) (*waf.GetWebACLOutput, error) {
	return a.client.GetWebACL(input)
}

func (a *WAFActivities) GetXssMatchSet(input *waf.GetXssMatchSetInput) (*waf.GetXssMatchSetOutput, error) {
	return a.client.GetXssMatchSet(input)
}

func (a *WAFActivities) ListActivatedRulesInRuleGroup(input *waf.ListActivatedRulesInRuleGroupInput) (*waf.ListActivatedRulesInRuleGroupOutput, error) {
	return a.client.ListActivatedRulesInRuleGroup(input)
}

func (a *WAFActivities) ListByteMatchSets(input *waf.ListByteMatchSetsInput) (*waf.ListByteMatchSetsOutput, error) {
	return a.client.ListByteMatchSets(input)
}

func (a *WAFActivities) ListGeoMatchSets(input *waf.ListGeoMatchSetsInput) (*waf.ListGeoMatchSetsOutput, error) {
	return a.client.ListGeoMatchSets(input)
}

func (a *WAFActivities) ListIPSets(input *waf.ListIPSetsInput) (*waf.ListIPSetsOutput, error) {
	return a.client.ListIPSets(input)
}

func (a *WAFActivities) ListLoggingConfigurations(input *waf.ListLoggingConfigurationsInput) (*waf.ListLoggingConfigurationsOutput, error) {
	return a.client.ListLoggingConfigurations(input)
}

func (a *WAFActivities) ListRateBasedRules(input *waf.ListRateBasedRulesInput) (*waf.ListRateBasedRulesOutput, error) {
	return a.client.ListRateBasedRules(input)
}

func (a *WAFActivities) ListRegexMatchSets(input *waf.ListRegexMatchSetsInput) (*waf.ListRegexMatchSetsOutput, error) {
	return a.client.ListRegexMatchSets(input)
}

func (a *WAFActivities) ListRegexPatternSets(input *waf.ListRegexPatternSetsInput) (*waf.ListRegexPatternSetsOutput, error) {
	return a.client.ListRegexPatternSets(input)
}

func (a *WAFActivities) ListRuleGroups(input *waf.ListRuleGroupsInput) (*waf.ListRuleGroupsOutput, error) {
	return a.client.ListRuleGroups(input)
}

func (a *WAFActivities) ListRules(input *waf.ListRulesInput) (*waf.ListRulesOutput, error) {
	return a.client.ListRules(input)
}

func (a *WAFActivities) ListSizeConstraintSets(input *waf.ListSizeConstraintSetsInput) (*waf.ListSizeConstraintSetsOutput, error) {
	return a.client.ListSizeConstraintSets(input)
}

func (a *WAFActivities) ListSqlInjectionMatchSets(input *waf.ListSqlInjectionMatchSetsInput) (*waf.ListSqlInjectionMatchSetsOutput, error) {
	return a.client.ListSqlInjectionMatchSets(input)
}

func (a *WAFActivities) ListSubscribedRuleGroups(input *waf.ListSubscribedRuleGroupsInput) (*waf.ListSubscribedRuleGroupsOutput, error) {
	return a.client.ListSubscribedRuleGroups(input)
}

func (a *WAFActivities) ListTagsForResource(input *waf.ListTagsForResourceInput) (*waf.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResource(input)
}

func (a *WAFActivities) ListWebACLs(input *waf.ListWebACLsInput) (*waf.ListWebACLsOutput, error) {
	return a.client.ListWebACLs(input)
}

func (a *WAFActivities) ListXssMatchSets(input *waf.ListXssMatchSetsInput) (*waf.ListXssMatchSetsOutput, error) {
	return a.client.ListXssMatchSets(input)
}

func (a *WAFActivities) PutLoggingConfiguration(input *waf.PutLoggingConfigurationInput) (*waf.PutLoggingConfigurationOutput, error) {
	return a.client.PutLoggingConfiguration(input)
}

func (a *WAFActivities) PutPermissionPolicy(input *waf.PutPermissionPolicyInput) (*waf.PutPermissionPolicyOutput, error) {
	return a.client.PutPermissionPolicy(input)
}

func (a *WAFActivities) TagResource(input *waf.TagResourceInput) (*waf.TagResourceOutput, error) {
	return a.client.TagResource(input)
}

func (a *WAFActivities) UntagResource(input *waf.UntagResourceInput) (*waf.UntagResourceOutput, error) {
	return a.client.UntagResource(input)
}

func (a *WAFActivities) UpdateByteMatchSet(input *waf.UpdateByteMatchSetInput) (*waf.UpdateByteMatchSetOutput, error) {
	return a.client.UpdateByteMatchSet(input)
}

func (a *WAFActivities) UpdateGeoMatchSet(input *waf.UpdateGeoMatchSetInput) (*waf.UpdateGeoMatchSetOutput, error) {
	return a.client.UpdateGeoMatchSet(input)
}

func (a *WAFActivities) UpdateIPSet(input *waf.UpdateIPSetInput) (*waf.UpdateIPSetOutput, error) {
	return a.client.UpdateIPSet(input)
}

func (a *WAFActivities) UpdateRateBasedRule(input *waf.UpdateRateBasedRuleInput) (*waf.UpdateRateBasedRuleOutput, error) {
	return a.client.UpdateRateBasedRule(input)
}

func (a *WAFActivities) UpdateRegexMatchSet(input *waf.UpdateRegexMatchSetInput) (*waf.UpdateRegexMatchSetOutput, error) {
	return a.client.UpdateRegexMatchSet(input)
}

func (a *WAFActivities) UpdateRegexPatternSet(input *waf.UpdateRegexPatternSetInput) (*waf.UpdateRegexPatternSetOutput, error) {
	return a.client.UpdateRegexPatternSet(input)
}

func (a *WAFActivities) UpdateRule(input *waf.UpdateRuleInput) (*waf.UpdateRuleOutput, error) {
	return a.client.UpdateRule(input)
}

func (a *WAFActivities) UpdateRuleGroup(input *waf.UpdateRuleGroupInput) (*waf.UpdateRuleGroupOutput, error) {
	return a.client.UpdateRuleGroup(input)
}

func (a *WAFActivities) UpdateSizeConstraintSet(input *waf.UpdateSizeConstraintSetInput) (*waf.UpdateSizeConstraintSetOutput, error) {
	return a.client.UpdateSizeConstraintSet(input)
}

func (a *WAFActivities) UpdateSqlInjectionMatchSet(input *waf.UpdateSqlInjectionMatchSetInput) (*waf.UpdateSqlInjectionMatchSetOutput, error) {
	return a.client.UpdateSqlInjectionMatchSet(input)
}

func (a *WAFActivities) UpdateWebACL(input *waf.UpdateWebACLInput) (*waf.UpdateWebACLOutput, error) {
	return a.client.UpdateWebACL(input)
}

func (a *WAFActivities) UpdateXssMatchSet(input *waf.UpdateXssMatchSetInput) (*waf.UpdateXssMatchSetOutput, error) {
	return a.client.UpdateXssMatchSet(input)
}
