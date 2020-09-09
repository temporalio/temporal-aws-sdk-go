
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/wafv2"
	"github.com/aws/aws-sdk-go/service/wafv2/wafv2iface"
)

type WAFV2Activities struct {
    client wafv2iface.WAFV2API
}

func NewWAFV2Activities(session *session.Session, config ...*aws.Config) *WAFV2Activities {
    client := wafv2.New(session, config...)
    return &WAFV2Activities{client: client}
}

func (a *WAFV2Activities) AssociateWebACL(input *wafv2.AssociateWebACLInput) (*wafv2.AssociateWebACLOutput, error) {
    return a.client.AssociateWebACL(input)
}

func (a *WAFV2Activities) CheckCapacity(input *wafv2.CheckCapacityInput) (*wafv2.CheckCapacityOutput, error) {
    return a.client.CheckCapacity(input)
}

func (a *WAFV2Activities) CreateIPSet(input *wafv2.CreateIPSetInput) (*wafv2.CreateIPSetOutput, error) {
    return a.client.CreateIPSet(input)
}

func (a *WAFV2Activities) CreateRegexPatternSet(input *wafv2.CreateRegexPatternSetInput) (*wafv2.CreateRegexPatternSetOutput, error) {
    return a.client.CreateRegexPatternSet(input)
}

func (a *WAFV2Activities) CreateRuleGroup(input *wafv2.CreateRuleGroupInput) (*wafv2.CreateRuleGroupOutput, error) {
    return a.client.CreateRuleGroup(input)
}

func (a *WAFV2Activities) CreateWebACL(input *wafv2.CreateWebACLInput) (*wafv2.CreateWebACLOutput, error) {
    return a.client.CreateWebACL(input)
}

func (a *WAFV2Activities) DeleteFirewallManagerRuleGroups(input *wafv2.DeleteFirewallManagerRuleGroupsInput) (*wafv2.DeleteFirewallManagerRuleGroupsOutput, error) {
    return a.client.DeleteFirewallManagerRuleGroups(input)
}

func (a *WAFV2Activities) DeleteIPSet(input *wafv2.DeleteIPSetInput) (*wafv2.DeleteIPSetOutput, error) {
    return a.client.DeleteIPSet(input)
}

func (a *WAFV2Activities) DeleteLoggingConfiguration(input *wafv2.DeleteLoggingConfigurationInput) (*wafv2.DeleteLoggingConfigurationOutput, error) {
    return a.client.DeleteLoggingConfiguration(input)
}

func (a *WAFV2Activities) DeletePermissionPolicy(input *wafv2.DeletePermissionPolicyInput) (*wafv2.DeletePermissionPolicyOutput, error) {
    return a.client.DeletePermissionPolicy(input)
}

func (a *WAFV2Activities) DeleteRegexPatternSet(input *wafv2.DeleteRegexPatternSetInput) (*wafv2.DeleteRegexPatternSetOutput, error) {
    return a.client.DeleteRegexPatternSet(input)
}

func (a *WAFV2Activities) DeleteRuleGroup(input *wafv2.DeleteRuleGroupInput) (*wafv2.DeleteRuleGroupOutput, error) {
    return a.client.DeleteRuleGroup(input)
}

func (a *WAFV2Activities) DeleteWebACL(input *wafv2.DeleteWebACLInput) (*wafv2.DeleteWebACLOutput, error) {
    return a.client.DeleteWebACL(input)
}

func (a *WAFV2Activities) DescribeManagedRuleGroup(input *wafv2.DescribeManagedRuleGroupInput) (*wafv2.DescribeManagedRuleGroupOutput, error) {
    return a.client.DescribeManagedRuleGroup(input)
}

func (a *WAFV2Activities) DisassociateWebACL(input *wafv2.DisassociateWebACLInput) (*wafv2.DisassociateWebACLOutput, error) {
    return a.client.DisassociateWebACL(input)
}

func (a *WAFV2Activities) GetIPSet(input *wafv2.GetIPSetInput) (*wafv2.GetIPSetOutput, error) {
    return a.client.GetIPSet(input)
}

func (a *WAFV2Activities) GetLoggingConfiguration(input *wafv2.GetLoggingConfigurationInput) (*wafv2.GetLoggingConfigurationOutput, error) {
    return a.client.GetLoggingConfiguration(input)
}

func (a *WAFV2Activities) GetPermissionPolicy(input *wafv2.GetPermissionPolicyInput) (*wafv2.GetPermissionPolicyOutput, error) {
    return a.client.GetPermissionPolicy(input)
}

func (a *WAFV2Activities) GetRateBasedStatementManagedKeys(input *wafv2.GetRateBasedStatementManagedKeysInput) (*wafv2.GetRateBasedStatementManagedKeysOutput, error) {
    return a.client.GetRateBasedStatementManagedKeys(input)
}

func (a *WAFV2Activities) GetRegexPatternSet(input *wafv2.GetRegexPatternSetInput) (*wafv2.GetRegexPatternSetOutput, error) {
    return a.client.GetRegexPatternSet(input)
}

func (a *WAFV2Activities) GetRuleGroup(input *wafv2.GetRuleGroupInput) (*wafv2.GetRuleGroupOutput, error) {
    return a.client.GetRuleGroup(input)
}

func (a *WAFV2Activities) GetSampledRequests(input *wafv2.GetSampledRequestsInput) (*wafv2.GetSampledRequestsOutput, error) {
    return a.client.GetSampledRequests(input)
}

func (a *WAFV2Activities) GetWebACL(input *wafv2.GetWebACLInput) (*wafv2.GetWebACLOutput, error) {
    return a.client.GetWebACL(input)
}

func (a *WAFV2Activities) GetWebACLForResource(input *wafv2.GetWebACLForResourceInput) (*wafv2.GetWebACLForResourceOutput, error) {
    return a.client.GetWebACLForResource(input)
}

func (a *WAFV2Activities) ListAvailableManagedRuleGroups(input *wafv2.ListAvailableManagedRuleGroupsInput) (*wafv2.ListAvailableManagedRuleGroupsOutput, error) {
    return a.client.ListAvailableManagedRuleGroups(input)
}

func (a *WAFV2Activities) ListIPSets(input *wafv2.ListIPSetsInput) (*wafv2.ListIPSetsOutput, error) {
    return a.client.ListIPSets(input)
}

func (a *WAFV2Activities) ListLoggingConfigurations(input *wafv2.ListLoggingConfigurationsInput) (*wafv2.ListLoggingConfigurationsOutput, error) {
    return a.client.ListLoggingConfigurations(input)
}

func (a *WAFV2Activities) ListRegexPatternSets(input *wafv2.ListRegexPatternSetsInput) (*wafv2.ListRegexPatternSetsOutput, error) {
    return a.client.ListRegexPatternSets(input)
}

func (a *WAFV2Activities) ListResourcesForWebACL(input *wafv2.ListResourcesForWebACLInput) (*wafv2.ListResourcesForWebACLOutput, error) {
    return a.client.ListResourcesForWebACL(input)
}

func (a *WAFV2Activities) ListRuleGroups(input *wafv2.ListRuleGroupsInput) (*wafv2.ListRuleGroupsOutput, error) {
    return a.client.ListRuleGroups(input)
}

func (a *WAFV2Activities) ListTagsForResource(input *wafv2.ListTagsForResourceInput) (*wafv2.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *WAFV2Activities) ListWebACLs(input *wafv2.ListWebACLsInput) (*wafv2.ListWebACLsOutput, error) {
    return a.client.ListWebACLs(input)
}

func (a *WAFV2Activities) PutLoggingConfiguration(input *wafv2.PutLoggingConfigurationInput) (*wafv2.PutLoggingConfigurationOutput, error) {
    return a.client.PutLoggingConfiguration(input)
}

func (a *WAFV2Activities) PutPermissionPolicy(input *wafv2.PutPermissionPolicyInput) (*wafv2.PutPermissionPolicyOutput, error) {
    return a.client.PutPermissionPolicy(input)
}

func (a *WAFV2Activities) TagResource(input *wafv2.TagResourceInput) (*wafv2.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *WAFV2Activities) UntagResource(input *wafv2.UntagResourceInput) (*wafv2.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *WAFV2Activities) UpdateIPSet(input *wafv2.UpdateIPSetInput) (*wafv2.UpdateIPSetOutput, error) {
    return a.client.UpdateIPSet(input)
}

func (a *WAFV2Activities) UpdateRegexPatternSet(input *wafv2.UpdateRegexPatternSetInput) (*wafv2.UpdateRegexPatternSetOutput, error) {
    return a.client.UpdateRegexPatternSet(input)
}

func (a *WAFV2Activities) UpdateRuleGroup(input *wafv2.UpdateRuleGroupInput) (*wafv2.UpdateRuleGroupOutput, error) {
    return a.client.UpdateRuleGroup(input)
}

func (a *WAFV2Activities) UpdateWebACL(input *wafv2.UpdateWebACLInput) (*wafv2.UpdateWebACLOutput, error) {
    return a.client.UpdateWebACL(input)
}
