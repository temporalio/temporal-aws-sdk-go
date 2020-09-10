package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/aws/aws-sdk-go/service/waf/wafiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type WAFActivities struct {
	client wafiface.WAFAPI
}

func NewWAFActivities(session *session.Session, config ...*aws.Config) *WAFActivities {
	client := waf.New(session, config...)
	return &WAFActivities{client: client}
}

func (a *WAFActivities) CreateByteMatchSet(ctx context.Context, input *waf.CreateByteMatchSetInput) (*waf.CreateByteMatchSetOutput, error) {
	return a.client.CreateByteMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) CreateGeoMatchSet(ctx context.Context, input *waf.CreateGeoMatchSetInput) (*waf.CreateGeoMatchSetOutput, error) {
	return a.client.CreateGeoMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) CreateIPSet(ctx context.Context, input *waf.CreateIPSetInput) (*waf.CreateIPSetOutput, error) {
	return a.client.CreateIPSetWithContext(ctx, input)
}

func (a *WAFActivities) CreateRateBasedRule(ctx context.Context, input *waf.CreateRateBasedRuleInput) (*waf.CreateRateBasedRuleOutput, error) {
	return a.client.CreateRateBasedRuleWithContext(ctx, input)
}

func (a *WAFActivities) CreateRegexMatchSet(ctx context.Context, input *waf.CreateRegexMatchSetInput) (*waf.CreateRegexMatchSetOutput, error) {
	return a.client.CreateRegexMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) CreateRegexPatternSet(ctx context.Context, input *waf.CreateRegexPatternSetInput) (*waf.CreateRegexPatternSetOutput, error) {
	return a.client.CreateRegexPatternSetWithContext(ctx, input)
}

func (a *WAFActivities) CreateRule(ctx context.Context, input *waf.CreateRuleInput) (*waf.CreateRuleOutput, error) {
	return a.client.CreateRuleWithContext(ctx, input)
}

func (a *WAFActivities) CreateRuleGroup(ctx context.Context, input *waf.CreateRuleGroupInput) (*waf.CreateRuleGroupOutput, error) {
	return a.client.CreateRuleGroupWithContext(ctx, input)
}

func (a *WAFActivities) CreateSizeConstraintSet(ctx context.Context, input *waf.CreateSizeConstraintSetInput) (*waf.CreateSizeConstraintSetOutput, error) {
	return a.client.CreateSizeConstraintSetWithContext(ctx, input)
}

func (a *WAFActivities) CreateSqlInjectionMatchSet(ctx context.Context, input *waf.CreateSqlInjectionMatchSetInput) (*waf.CreateSqlInjectionMatchSetOutput, error) {
	return a.client.CreateSqlInjectionMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) CreateWebACL(ctx context.Context, input *waf.CreateWebACLInput) (*waf.CreateWebACLOutput, error) {
	return a.client.CreateWebACLWithContext(ctx, input)
}

func (a *WAFActivities) CreateWebACLMigrationStack(ctx context.Context, input *waf.CreateWebACLMigrationStackInput) (*waf.CreateWebACLMigrationStackOutput, error) {
	return a.client.CreateWebACLMigrationStackWithContext(ctx, input)
}

func (a *WAFActivities) CreateXssMatchSet(ctx context.Context, input *waf.CreateXssMatchSetInput) (*waf.CreateXssMatchSetOutput, error) {
	return a.client.CreateXssMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) DeleteByteMatchSet(ctx context.Context, input *waf.DeleteByteMatchSetInput) (*waf.DeleteByteMatchSetOutput, error) {
	return a.client.DeleteByteMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) DeleteGeoMatchSet(ctx context.Context, input *waf.DeleteGeoMatchSetInput) (*waf.DeleteGeoMatchSetOutput, error) {
	return a.client.DeleteGeoMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) DeleteIPSet(ctx context.Context, input *waf.DeleteIPSetInput) (*waf.DeleteIPSetOutput, error) {
	return a.client.DeleteIPSetWithContext(ctx, input)
}

func (a *WAFActivities) DeleteLoggingConfiguration(ctx context.Context, input *waf.DeleteLoggingConfigurationInput) (*waf.DeleteLoggingConfigurationOutput, error) {
	return a.client.DeleteLoggingConfigurationWithContext(ctx, input)
}

func (a *WAFActivities) DeletePermissionPolicy(ctx context.Context, input *waf.DeletePermissionPolicyInput) (*waf.DeletePermissionPolicyOutput, error) {
	return a.client.DeletePermissionPolicyWithContext(ctx, input)
}

func (a *WAFActivities) DeleteRateBasedRule(ctx context.Context, input *waf.DeleteRateBasedRuleInput) (*waf.DeleteRateBasedRuleOutput, error) {
	return a.client.DeleteRateBasedRuleWithContext(ctx, input)
}

func (a *WAFActivities) DeleteRegexMatchSet(ctx context.Context, input *waf.DeleteRegexMatchSetInput) (*waf.DeleteRegexMatchSetOutput, error) {
	return a.client.DeleteRegexMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) DeleteRegexPatternSet(ctx context.Context, input *waf.DeleteRegexPatternSetInput) (*waf.DeleteRegexPatternSetOutput, error) {
	return a.client.DeleteRegexPatternSetWithContext(ctx, input)
}

func (a *WAFActivities) DeleteRule(ctx context.Context, input *waf.DeleteRuleInput) (*waf.DeleteRuleOutput, error) {
	return a.client.DeleteRuleWithContext(ctx, input)
}

func (a *WAFActivities) DeleteRuleGroup(ctx context.Context, input *waf.DeleteRuleGroupInput) (*waf.DeleteRuleGroupOutput, error) {
	return a.client.DeleteRuleGroupWithContext(ctx, input)
}

func (a *WAFActivities) DeleteSizeConstraintSet(ctx context.Context, input *waf.DeleteSizeConstraintSetInput) (*waf.DeleteSizeConstraintSetOutput, error) {
	return a.client.DeleteSizeConstraintSetWithContext(ctx, input)
}

func (a *WAFActivities) DeleteSqlInjectionMatchSet(ctx context.Context, input *waf.DeleteSqlInjectionMatchSetInput) (*waf.DeleteSqlInjectionMatchSetOutput, error) {
	return a.client.DeleteSqlInjectionMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) DeleteWebACL(ctx context.Context, input *waf.DeleteWebACLInput) (*waf.DeleteWebACLOutput, error) {
	return a.client.DeleteWebACLWithContext(ctx, input)
}

func (a *WAFActivities) DeleteXssMatchSet(ctx context.Context, input *waf.DeleteXssMatchSetInput) (*waf.DeleteXssMatchSetOutput, error) {
	return a.client.DeleteXssMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) GetByteMatchSet(ctx context.Context, input *waf.GetByteMatchSetInput) (*waf.GetByteMatchSetOutput, error) {
	return a.client.GetByteMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) GetChangeToken(ctx context.Context, input *waf.GetChangeTokenInput) (*waf.GetChangeTokenOutput, error) {
	return a.client.GetChangeTokenWithContext(ctx, input)
}

func (a *WAFActivities) GetChangeTokenStatus(ctx context.Context, input *waf.GetChangeTokenStatusInput) (*waf.GetChangeTokenStatusOutput, error) {
	return a.client.GetChangeTokenStatusWithContext(ctx, input)
}

func (a *WAFActivities) GetGeoMatchSet(ctx context.Context, input *waf.GetGeoMatchSetInput) (*waf.GetGeoMatchSetOutput, error) {
	return a.client.GetGeoMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) GetIPSet(ctx context.Context, input *waf.GetIPSetInput) (*waf.GetIPSetOutput, error) {
	return a.client.GetIPSetWithContext(ctx, input)
}

func (a *WAFActivities) GetLoggingConfiguration(ctx context.Context, input *waf.GetLoggingConfigurationInput) (*waf.GetLoggingConfigurationOutput, error) {
	return a.client.GetLoggingConfigurationWithContext(ctx, input)
}

func (a *WAFActivities) GetPermissionPolicy(ctx context.Context, input *waf.GetPermissionPolicyInput) (*waf.GetPermissionPolicyOutput, error) {
	return a.client.GetPermissionPolicyWithContext(ctx, input)
}

func (a *WAFActivities) GetRateBasedRule(ctx context.Context, input *waf.GetRateBasedRuleInput) (*waf.GetRateBasedRuleOutput, error) {
	return a.client.GetRateBasedRuleWithContext(ctx, input)
}

func (a *WAFActivities) GetRateBasedRuleManagedKeys(ctx context.Context, input *waf.GetRateBasedRuleManagedKeysInput) (*waf.GetRateBasedRuleManagedKeysOutput, error) {
	return a.client.GetRateBasedRuleManagedKeysWithContext(ctx, input)
}

func (a *WAFActivities) GetRegexMatchSet(ctx context.Context, input *waf.GetRegexMatchSetInput) (*waf.GetRegexMatchSetOutput, error) {
	return a.client.GetRegexMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) GetRegexPatternSet(ctx context.Context, input *waf.GetRegexPatternSetInput) (*waf.GetRegexPatternSetOutput, error) {
	return a.client.GetRegexPatternSetWithContext(ctx, input)
}

func (a *WAFActivities) GetRule(ctx context.Context, input *waf.GetRuleInput) (*waf.GetRuleOutput, error) {
	return a.client.GetRuleWithContext(ctx, input)
}

func (a *WAFActivities) GetRuleGroup(ctx context.Context, input *waf.GetRuleGroupInput) (*waf.GetRuleGroupOutput, error) {
	return a.client.GetRuleGroupWithContext(ctx, input)
}

func (a *WAFActivities) GetSampledRequests(ctx context.Context, input *waf.GetSampledRequestsInput) (*waf.GetSampledRequestsOutput, error) {
	return a.client.GetSampledRequestsWithContext(ctx, input)
}

func (a *WAFActivities) GetSizeConstraintSet(ctx context.Context, input *waf.GetSizeConstraintSetInput) (*waf.GetSizeConstraintSetOutput, error) {
	return a.client.GetSizeConstraintSetWithContext(ctx, input)
}

func (a *WAFActivities) GetSqlInjectionMatchSet(ctx context.Context, input *waf.GetSqlInjectionMatchSetInput) (*waf.GetSqlInjectionMatchSetOutput, error) {
	return a.client.GetSqlInjectionMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) GetWebACL(ctx context.Context, input *waf.GetWebACLInput) (*waf.GetWebACLOutput, error) {
	return a.client.GetWebACLWithContext(ctx, input)
}

func (a *WAFActivities) GetXssMatchSet(ctx context.Context, input *waf.GetXssMatchSetInput) (*waf.GetXssMatchSetOutput, error) {
	return a.client.GetXssMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) ListActivatedRulesInRuleGroup(ctx context.Context, input *waf.ListActivatedRulesInRuleGroupInput) (*waf.ListActivatedRulesInRuleGroupOutput, error) {
	return a.client.ListActivatedRulesInRuleGroupWithContext(ctx, input)
}

func (a *WAFActivities) ListByteMatchSets(ctx context.Context, input *waf.ListByteMatchSetsInput) (*waf.ListByteMatchSetsOutput, error) {
	return a.client.ListByteMatchSetsWithContext(ctx, input)
}

func (a *WAFActivities) ListGeoMatchSets(ctx context.Context, input *waf.ListGeoMatchSetsInput) (*waf.ListGeoMatchSetsOutput, error) {
	return a.client.ListGeoMatchSetsWithContext(ctx, input)
}

func (a *WAFActivities) ListIPSets(ctx context.Context, input *waf.ListIPSetsInput) (*waf.ListIPSetsOutput, error) {
	return a.client.ListIPSetsWithContext(ctx, input)
}

func (a *WAFActivities) ListLoggingConfigurations(ctx context.Context, input *waf.ListLoggingConfigurationsInput) (*waf.ListLoggingConfigurationsOutput, error) {
	return a.client.ListLoggingConfigurationsWithContext(ctx, input)
}

func (a *WAFActivities) ListRateBasedRules(ctx context.Context, input *waf.ListRateBasedRulesInput) (*waf.ListRateBasedRulesOutput, error) {
	return a.client.ListRateBasedRulesWithContext(ctx, input)
}

func (a *WAFActivities) ListRegexMatchSets(ctx context.Context, input *waf.ListRegexMatchSetsInput) (*waf.ListRegexMatchSetsOutput, error) {
	return a.client.ListRegexMatchSetsWithContext(ctx, input)
}

func (a *WAFActivities) ListRegexPatternSets(ctx context.Context, input *waf.ListRegexPatternSetsInput) (*waf.ListRegexPatternSetsOutput, error) {
	return a.client.ListRegexPatternSetsWithContext(ctx, input)
}

func (a *WAFActivities) ListRuleGroups(ctx context.Context, input *waf.ListRuleGroupsInput) (*waf.ListRuleGroupsOutput, error) {
	return a.client.ListRuleGroupsWithContext(ctx, input)
}

func (a *WAFActivities) ListRules(ctx context.Context, input *waf.ListRulesInput) (*waf.ListRulesOutput, error) {
	return a.client.ListRulesWithContext(ctx, input)
}

func (a *WAFActivities) ListSizeConstraintSets(ctx context.Context, input *waf.ListSizeConstraintSetsInput) (*waf.ListSizeConstraintSetsOutput, error) {
	return a.client.ListSizeConstraintSetsWithContext(ctx, input)
}

func (a *WAFActivities) ListSqlInjectionMatchSets(ctx context.Context, input *waf.ListSqlInjectionMatchSetsInput) (*waf.ListSqlInjectionMatchSetsOutput, error) {
	return a.client.ListSqlInjectionMatchSetsWithContext(ctx, input)
}

func (a *WAFActivities) ListSubscribedRuleGroups(ctx context.Context, input *waf.ListSubscribedRuleGroupsInput) (*waf.ListSubscribedRuleGroupsOutput, error) {
	return a.client.ListSubscribedRuleGroupsWithContext(ctx, input)
}

func (a *WAFActivities) ListTagsForResource(ctx context.Context, input *waf.ListTagsForResourceInput) (*waf.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *WAFActivities) ListWebACLs(ctx context.Context, input *waf.ListWebACLsInput) (*waf.ListWebACLsOutput, error) {
	return a.client.ListWebACLsWithContext(ctx, input)
}

func (a *WAFActivities) ListXssMatchSets(ctx context.Context, input *waf.ListXssMatchSetsInput) (*waf.ListXssMatchSetsOutput, error) {
	return a.client.ListXssMatchSetsWithContext(ctx, input)
}

func (a *WAFActivities) PutLoggingConfiguration(ctx context.Context, input *waf.PutLoggingConfigurationInput) (*waf.PutLoggingConfigurationOutput, error) {
	return a.client.PutLoggingConfigurationWithContext(ctx, input)
}

func (a *WAFActivities) PutPermissionPolicy(ctx context.Context, input *waf.PutPermissionPolicyInput) (*waf.PutPermissionPolicyOutput, error) {
	return a.client.PutPermissionPolicyWithContext(ctx, input)
}

func (a *WAFActivities) TagResource(ctx context.Context, input *waf.TagResourceInput) (*waf.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *WAFActivities) UntagResource(ctx context.Context, input *waf.UntagResourceInput) (*waf.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *WAFActivities) UpdateByteMatchSet(ctx context.Context, input *waf.UpdateByteMatchSetInput) (*waf.UpdateByteMatchSetOutput, error) {
	return a.client.UpdateByteMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) UpdateGeoMatchSet(ctx context.Context, input *waf.UpdateGeoMatchSetInput) (*waf.UpdateGeoMatchSetOutput, error) {
	return a.client.UpdateGeoMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) UpdateIPSet(ctx context.Context, input *waf.UpdateIPSetInput) (*waf.UpdateIPSetOutput, error) {
	return a.client.UpdateIPSetWithContext(ctx, input)
}

func (a *WAFActivities) UpdateRateBasedRule(ctx context.Context, input *waf.UpdateRateBasedRuleInput) (*waf.UpdateRateBasedRuleOutput, error) {
	return a.client.UpdateRateBasedRuleWithContext(ctx, input)
}

func (a *WAFActivities) UpdateRegexMatchSet(ctx context.Context, input *waf.UpdateRegexMatchSetInput) (*waf.UpdateRegexMatchSetOutput, error) {
	return a.client.UpdateRegexMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) UpdateRegexPatternSet(ctx context.Context, input *waf.UpdateRegexPatternSetInput) (*waf.UpdateRegexPatternSetOutput, error) {
	return a.client.UpdateRegexPatternSetWithContext(ctx, input)
}

func (a *WAFActivities) UpdateRule(ctx context.Context, input *waf.UpdateRuleInput) (*waf.UpdateRuleOutput, error) {
	return a.client.UpdateRuleWithContext(ctx, input)
}

func (a *WAFActivities) UpdateRuleGroup(ctx context.Context, input *waf.UpdateRuleGroupInput) (*waf.UpdateRuleGroupOutput, error) {
	return a.client.UpdateRuleGroupWithContext(ctx, input)
}

func (a *WAFActivities) UpdateSizeConstraintSet(ctx context.Context, input *waf.UpdateSizeConstraintSetInput) (*waf.UpdateSizeConstraintSetOutput, error) {
	return a.client.UpdateSizeConstraintSetWithContext(ctx, input)
}

func (a *WAFActivities) UpdateSqlInjectionMatchSet(ctx context.Context, input *waf.UpdateSqlInjectionMatchSetInput) (*waf.UpdateSqlInjectionMatchSetOutput, error) {
	return a.client.UpdateSqlInjectionMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) UpdateWebACL(ctx context.Context, input *waf.UpdateWebACLInput) (*waf.UpdateWebACLOutput, error) {
	return a.client.UpdateWebACLWithContext(ctx, input)
}

func (a *WAFActivities) UpdateXssMatchSet(ctx context.Context, input *waf.UpdateXssMatchSetInput) (*waf.UpdateXssMatchSetOutput, error) {
	return a.client.UpdateXssMatchSetWithContext(ctx, input)
}
