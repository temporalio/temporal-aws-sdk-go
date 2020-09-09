package awsclients

import (
	"github.com/aws/aws-sdk-go/service/wafv2"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type WAFV2Client interface {
    AssociateWebACL(ctx workflow.Context, input *wafv2.AssociateWebACLInput) (*wafv2.AssociateWebACLOutput, error)
    AssociateWebACLAsync(ctx workflow.Context, input *wafv2.AssociateWebACLInput) *Wafv2AssociateWebACLResult

    CheckCapacity(ctx workflow.Context, input *wafv2.CheckCapacityInput) (*wafv2.CheckCapacityOutput, error)
    CheckCapacityAsync(ctx workflow.Context, input *wafv2.CheckCapacityInput) *Wafv2CheckCapacityResult

    CreateIPSet(ctx workflow.Context, input *wafv2.CreateIPSetInput) (*wafv2.CreateIPSetOutput, error)
    CreateIPSetAsync(ctx workflow.Context, input *wafv2.CreateIPSetInput) *Wafv2CreateIPSetResult

    CreateRegexPatternSet(ctx workflow.Context, input *wafv2.CreateRegexPatternSetInput) (*wafv2.CreateRegexPatternSetOutput, error)
    CreateRegexPatternSetAsync(ctx workflow.Context, input *wafv2.CreateRegexPatternSetInput) *Wafv2CreateRegexPatternSetResult

    CreateRuleGroup(ctx workflow.Context, input *wafv2.CreateRuleGroupInput) (*wafv2.CreateRuleGroupOutput, error)
    CreateRuleGroupAsync(ctx workflow.Context, input *wafv2.CreateRuleGroupInput) *Wafv2CreateRuleGroupResult

    CreateWebACL(ctx workflow.Context, input *wafv2.CreateWebACLInput) (*wafv2.CreateWebACLOutput, error)
    CreateWebACLAsync(ctx workflow.Context, input *wafv2.CreateWebACLInput) *Wafv2CreateWebACLResult

    DeleteFirewallManagerRuleGroups(ctx workflow.Context, input *wafv2.DeleteFirewallManagerRuleGroupsInput) (*wafv2.DeleteFirewallManagerRuleGroupsOutput, error)
    DeleteFirewallManagerRuleGroupsAsync(ctx workflow.Context, input *wafv2.DeleteFirewallManagerRuleGroupsInput) *Wafv2DeleteFirewallManagerRuleGroupsResult

    DeleteIPSet(ctx workflow.Context, input *wafv2.DeleteIPSetInput) (*wafv2.DeleteIPSetOutput, error)
    DeleteIPSetAsync(ctx workflow.Context, input *wafv2.DeleteIPSetInput) *Wafv2DeleteIPSetResult

    DeleteLoggingConfiguration(ctx workflow.Context, input *wafv2.DeleteLoggingConfigurationInput) (*wafv2.DeleteLoggingConfigurationOutput, error)
    DeleteLoggingConfigurationAsync(ctx workflow.Context, input *wafv2.DeleteLoggingConfigurationInput) *Wafv2DeleteLoggingConfigurationResult

    DeletePermissionPolicy(ctx workflow.Context, input *wafv2.DeletePermissionPolicyInput) (*wafv2.DeletePermissionPolicyOutput, error)
    DeletePermissionPolicyAsync(ctx workflow.Context, input *wafv2.DeletePermissionPolicyInput) *Wafv2DeletePermissionPolicyResult

    DeleteRegexPatternSet(ctx workflow.Context, input *wafv2.DeleteRegexPatternSetInput) (*wafv2.DeleteRegexPatternSetOutput, error)
    DeleteRegexPatternSetAsync(ctx workflow.Context, input *wafv2.DeleteRegexPatternSetInput) *Wafv2DeleteRegexPatternSetResult

    DeleteRuleGroup(ctx workflow.Context, input *wafv2.DeleteRuleGroupInput) (*wafv2.DeleteRuleGroupOutput, error)
    DeleteRuleGroupAsync(ctx workflow.Context, input *wafv2.DeleteRuleGroupInput) *Wafv2DeleteRuleGroupResult

    DeleteWebACL(ctx workflow.Context, input *wafv2.DeleteWebACLInput) (*wafv2.DeleteWebACLOutput, error)
    DeleteWebACLAsync(ctx workflow.Context, input *wafv2.DeleteWebACLInput) *Wafv2DeleteWebACLResult

    DescribeManagedRuleGroup(ctx workflow.Context, input *wafv2.DescribeManagedRuleGroupInput) (*wafv2.DescribeManagedRuleGroupOutput, error)
    DescribeManagedRuleGroupAsync(ctx workflow.Context, input *wafv2.DescribeManagedRuleGroupInput) *Wafv2DescribeManagedRuleGroupResult

    DisassociateWebACL(ctx workflow.Context, input *wafv2.DisassociateWebACLInput) (*wafv2.DisassociateWebACLOutput, error)
    DisassociateWebACLAsync(ctx workflow.Context, input *wafv2.DisassociateWebACLInput) *Wafv2DisassociateWebACLResult

    GetIPSet(ctx workflow.Context, input *wafv2.GetIPSetInput) (*wafv2.GetIPSetOutput, error)
    GetIPSetAsync(ctx workflow.Context, input *wafv2.GetIPSetInput) *Wafv2GetIPSetResult

    GetLoggingConfiguration(ctx workflow.Context, input *wafv2.GetLoggingConfigurationInput) (*wafv2.GetLoggingConfigurationOutput, error)
    GetLoggingConfigurationAsync(ctx workflow.Context, input *wafv2.GetLoggingConfigurationInput) *Wafv2GetLoggingConfigurationResult

    GetPermissionPolicy(ctx workflow.Context, input *wafv2.GetPermissionPolicyInput) (*wafv2.GetPermissionPolicyOutput, error)
    GetPermissionPolicyAsync(ctx workflow.Context, input *wafv2.GetPermissionPolicyInput) *Wafv2GetPermissionPolicyResult

    GetRateBasedStatementManagedKeys(ctx workflow.Context, input *wafv2.GetRateBasedStatementManagedKeysInput) (*wafv2.GetRateBasedStatementManagedKeysOutput, error)
    GetRateBasedStatementManagedKeysAsync(ctx workflow.Context, input *wafv2.GetRateBasedStatementManagedKeysInput) *Wafv2GetRateBasedStatementManagedKeysResult

    GetRegexPatternSet(ctx workflow.Context, input *wafv2.GetRegexPatternSetInput) (*wafv2.GetRegexPatternSetOutput, error)
    GetRegexPatternSetAsync(ctx workflow.Context, input *wafv2.GetRegexPatternSetInput) *Wafv2GetRegexPatternSetResult

    GetRuleGroup(ctx workflow.Context, input *wafv2.GetRuleGroupInput) (*wafv2.GetRuleGroupOutput, error)
    GetRuleGroupAsync(ctx workflow.Context, input *wafv2.GetRuleGroupInput) *Wafv2GetRuleGroupResult

    GetSampledRequests(ctx workflow.Context, input *wafv2.GetSampledRequestsInput) (*wafv2.GetSampledRequestsOutput, error)
    GetSampledRequestsAsync(ctx workflow.Context, input *wafv2.GetSampledRequestsInput) *Wafv2GetSampledRequestsResult

    GetWebACL(ctx workflow.Context, input *wafv2.GetWebACLInput) (*wafv2.GetWebACLOutput, error)
    GetWebACLAsync(ctx workflow.Context, input *wafv2.GetWebACLInput) *Wafv2GetWebACLResult

    GetWebACLForResource(ctx workflow.Context, input *wafv2.GetWebACLForResourceInput) (*wafv2.GetWebACLForResourceOutput, error)
    GetWebACLForResourceAsync(ctx workflow.Context, input *wafv2.GetWebACLForResourceInput) *Wafv2GetWebACLForResourceResult

    ListAvailableManagedRuleGroups(ctx workflow.Context, input *wafv2.ListAvailableManagedRuleGroupsInput) (*wafv2.ListAvailableManagedRuleGroupsOutput, error)
    ListAvailableManagedRuleGroupsAsync(ctx workflow.Context, input *wafv2.ListAvailableManagedRuleGroupsInput) *Wafv2ListAvailableManagedRuleGroupsResult

    ListIPSets(ctx workflow.Context, input *wafv2.ListIPSetsInput) (*wafv2.ListIPSetsOutput, error)
    ListIPSetsAsync(ctx workflow.Context, input *wafv2.ListIPSetsInput) *Wafv2ListIPSetsResult

    ListLoggingConfigurations(ctx workflow.Context, input *wafv2.ListLoggingConfigurationsInput) (*wafv2.ListLoggingConfigurationsOutput, error)
    ListLoggingConfigurationsAsync(ctx workflow.Context, input *wafv2.ListLoggingConfigurationsInput) *Wafv2ListLoggingConfigurationsResult

    ListRegexPatternSets(ctx workflow.Context, input *wafv2.ListRegexPatternSetsInput) (*wafv2.ListRegexPatternSetsOutput, error)
    ListRegexPatternSetsAsync(ctx workflow.Context, input *wafv2.ListRegexPatternSetsInput) *Wafv2ListRegexPatternSetsResult

    ListResourcesForWebACL(ctx workflow.Context, input *wafv2.ListResourcesForWebACLInput) (*wafv2.ListResourcesForWebACLOutput, error)
    ListResourcesForWebACLAsync(ctx workflow.Context, input *wafv2.ListResourcesForWebACLInput) *Wafv2ListResourcesForWebACLResult

    ListRuleGroups(ctx workflow.Context, input *wafv2.ListRuleGroupsInput) (*wafv2.ListRuleGroupsOutput, error)
    ListRuleGroupsAsync(ctx workflow.Context, input *wafv2.ListRuleGroupsInput) *Wafv2ListRuleGroupsResult

    ListTagsForResource(ctx workflow.Context, input *wafv2.ListTagsForResourceInput) (*wafv2.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *wafv2.ListTagsForResourceInput) *Wafv2ListTagsForResourceResult

    ListWebACLs(ctx workflow.Context, input *wafv2.ListWebACLsInput) (*wafv2.ListWebACLsOutput, error)
    ListWebACLsAsync(ctx workflow.Context, input *wafv2.ListWebACLsInput) *Wafv2ListWebACLsResult

    PutLoggingConfiguration(ctx workflow.Context, input *wafv2.PutLoggingConfigurationInput) (*wafv2.PutLoggingConfigurationOutput, error)
    PutLoggingConfigurationAsync(ctx workflow.Context, input *wafv2.PutLoggingConfigurationInput) *Wafv2PutLoggingConfigurationResult

    PutPermissionPolicy(ctx workflow.Context, input *wafv2.PutPermissionPolicyInput) (*wafv2.PutPermissionPolicyOutput, error)
    PutPermissionPolicyAsync(ctx workflow.Context, input *wafv2.PutPermissionPolicyInput) *Wafv2PutPermissionPolicyResult

    TagResource(ctx workflow.Context, input *wafv2.TagResourceInput) (*wafv2.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *wafv2.TagResourceInput) *Wafv2TagResourceResult

    UntagResource(ctx workflow.Context, input *wafv2.UntagResourceInput) (*wafv2.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *wafv2.UntagResourceInput) *Wafv2UntagResourceResult

    UpdateIPSet(ctx workflow.Context, input *wafv2.UpdateIPSetInput) (*wafv2.UpdateIPSetOutput, error)
    UpdateIPSetAsync(ctx workflow.Context, input *wafv2.UpdateIPSetInput) *Wafv2UpdateIPSetResult

    UpdateRegexPatternSet(ctx workflow.Context, input *wafv2.UpdateRegexPatternSetInput) (*wafv2.UpdateRegexPatternSetOutput, error)
    UpdateRegexPatternSetAsync(ctx workflow.Context, input *wafv2.UpdateRegexPatternSetInput) *Wafv2UpdateRegexPatternSetResult

    UpdateRuleGroup(ctx workflow.Context, input *wafv2.UpdateRuleGroupInput) (*wafv2.UpdateRuleGroupOutput, error)
    UpdateRuleGroupAsync(ctx workflow.Context, input *wafv2.UpdateRuleGroupInput) *Wafv2UpdateRuleGroupResult

    UpdateWebACL(ctx workflow.Context, input *wafv2.UpdateWebACLInput) (*wafv2.UpdateWebACLOutput, error)
    UpdateWebACLAsync(ctx workflow.Context, input *wafv2.UpdateWebACLInput) *Wafv2UpdateWebACLResult
}

type Wafv2AssociateWebACLResult struct {
	Result workflow.Future
}

func (r *Wafv2AssociateWebACLResult) Get(ctx workflow.Context) (*wafv2.AssociateWebACLOutput, error) {
    var output wafv2.AssociateWebACLOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2CheckCapacityResult struct {
	Result workflow.Future
}

func (r *Wafv2CheckCapacityResult) Get(ctx workflow.Context) (*wafv2.CheckCapacityOutput, error) {
    var output wafv2.CheckCapacityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2CreateIPSetResult struct {
	Result workflow.Future
}

func (r *Wafv2CreateIPSetResult) Get(ctx workflow.Context) (*wafv2.CreateIPSetOutput, error) {
    var output wafv2.CreateIPSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2CreateRegexPatternSetResult struct {
	Result workflow.Future
}

func (r *Wafv2CreateRegexPatternSetResult) Get(ctx workflow.Context) (*wafv2.CreateRegexPatternSetOutput, error) {
    var output wafv2.CreateRegexPatternSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2CreateRuleGroupResult struct {
	Result workflow.Future
}

func (r *Wafv2CreateRuleGroupResult) Get(ctx workflow.Context) (*wafv2.CreateRuleGroupOutput, error) {
    var output wafv2.CreateRuleGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2CreateWebACLResult struct {
	Result workflow.Future
}

func (r *Wafv2CreateWebACLResult) Get(ctx workflow.Context) (*wafv2.CreateWebACLOutput, error) {
    var output wafv2.CreateWebACLOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2DeleteFirewallManagerRuleGroupsResult struct {
	Result workflow.Future
}

func (r *Wafv2DeleteFirewallManagerRuleGroupsResult) Get(ctx workflow.Context) (*wafv2.DeleteFirewallManagerRuleGroupsOutput, error) {
    var output wafv2.DeleteFirewallManagerRuleGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2DeleteIPSetResult struct {
	Result workflow.Future
}

func (r *Wafv2DeleteIPSetResult) Get(ctx workflow.Context) (*wafv2.DeleteIPSetOutput, error) {
    var output wafv2.DeleteIPSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2DeleteLoggingConfigurationResult struct {
	Result workflow.Future
}

func (r *Wafv2DeleteLoggingConfigurationResult) Get(ctx workflow.Context) (*wafv2.DeleteLoggingConfigurationOutput, error) {
    var output wafv2.DeleteLoggingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2DeletePermissionPolicyResult struct {
	Result workflow.Future
}

func (r *Wafv2DeletePermissionPolicyResult) Get(ctx workflow.Context) (*wafv2.DeletePermissionPolicyOutput, error) {
    var output wafv2.DeletePermissionPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2DeleteRegexPatternSetResult struct {
	Result workflow.Future
}

func (r *Wafv2DeleteRegexPatternSetResult) Get(ctx workflow.Context) (*wafv2.DeleteRegexPatternSetOutput, error) {
    var output wafv2.DeleteRegexPatternSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2DeleteRuleGroupResult struct {
	Result workflow.Future
}

func (r *Wafv2DeleteRuleGroupResult) Get(ctx workflow.Context) (*wafv2.DeleteRuleGroupOutput, error) {
    var output wafv2.DeleteRuleGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2DeleteWebACLResult struct {
	Result workflow.Future
}

func (r *Wafv2DeleteWebACLResult) Get(ctx workflow.Context) (*wafv2.DeleteWebACLOutput, error) {
    var output wafv2.DeleteWebACLOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2DescribeManagedRuleGroupResult struct {
	Result workflow.Future
}

func (r *Wafv2DescribeManagedRuleGroupResult) Get(ctx workflow.Context) (*wafv2.DescribeManagedRuleGroupOutput, error) {
    var output wafv2.DescribeManagedRuleGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2DisassociateWebACLResult struct {
	Result workflow.Future
}

func (r *Wafv2DisassociateWebACLResult) Get(ctx workflow.Context) (*wafv2.DisassociateWebACLOutput, error) {
    var output wafv2.DisassociateWebACLOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2GetIPSetResult struct {
	Result workflow.Future
}

func (r *Wafv2GetIPSetResult) Get(ctx workflow.Context) (*wafv2.GetIPSetOutput, error) {
    var output wafv2.GetIPSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2GetLoggingConfigurationResult struct {
	Result workflow.Future
}

func (r *Wafv2GetLoggingConfigurationResult) Get(ctx workflow.Context) (*wafv2.GetLoggingConfigurationOutput, error) {
    var output wafv2.GetLoggingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2GetPermissionPolicyResult struct {
	Result workflow.Future
}

func (r *Wafv2GetPermissionPolicyResult) Get(ctx workflow.Context) (*wafv2.GetPermissionPolicyOutput, error) {
    var output wafv2.GetPermissionPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2GetRateBasedStatementManagedKeysResult struct {
	Result workflow.Future
}

func (r *Wafv2GetRateBasedStatementManagedKeysResult) Get(ctx workflow.Context) (*wafv2.GetRateBasedStatementManagedKeysOutput, error) {
    var output wafv2.GetRateBasedStatementManagedKeysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2GetRegexPatternSetResult struct {
	Result workflow.Future
}

func (r *Wafv2GetRegexPatternSetResult) Get(ctx workflow.Context) (*wafv2.GetRegexPatternSetOutput, error) {
    var output wafv2.GetRegexPatternSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2GetRuleGroupResult struct {
	Result workflow.Future
}

func (r *Wafv2GetRuleGroupResult) Get(ctx workflow.Context) (*wafv2.GetRuleGroupOutput, error) {
    var output wafv2.GetRuleGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2GetSampledRequestsResult struct {
	Result workflow.Future
}

func (r *Wafv2GetSampledRequestsResult) Get(ctx workflow.Context) (*wafv2.GetSampledRequestsOutput, error) {
    var output wafv2.GetSampledRequestsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2GetWebACLResult struct {
	Result workflow.Future
}

func (r *Wafv2GetWebACLResult) Get(ctx workflow.Context) (*wafv2.GetWebACLOutput, error) {
    var output wafv2.GetWebACLOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2GetWebACLForResourceResult struct {
	Result workflow.Future
}

func (r *Wafv2GetWebACLForResourceResult) Get(ctx workflow.Context) (*wafv2.GetWebACLForResourceOutput, error) {
    var output wafv2.GetWebACLForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2ListAvailableManagedRuleGroupsResult struct {
	Result workflow.Future
}

func (r *Wafv2ListAvailableManagedRuleGroupsResult) Get(ctx workflow.Context) (*wafv2.ListAvailableManagedRuleGroupsOutput, error) {
    var output wafv2.ListAvailableManagedRuleGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2ListIPSetsResult struct {
	Result workflow.Future
}

func (r *Wafv2ListIPSetsResult) Get(ctx workflow.Context) (*wafv2.ListIPSetsOutput, error) {
    var output wafv2.ListIPSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2ListLoggingConfigurationsResult struct {
	Result workflow.Future
}

func (r *Wafv2ListLoggingConfigurationsResult) Get(ctx workflow.Context) (*wafv2.ListLoggingConfigurationsOutput, error) {
    var output wafv2.ListLoggingConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2ListRegexPatternSetsResult struct {
	Result workflow.Future
}

func (r *Wafv2ListRegexPatternSetsResult) Get(ctx workflow.Context) (*wafv2.ListRegexPatternSetsOutput, error) {
    var output wafv2.ListRegexPatternSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2ListResourcesForWebACLResult struct {
	Result workflow.Future
}

func (r *Wafv2ListResourcesForWebACLResult) Get(ctx workflow.Context) (*wafv2.ListResourcesForWebACLOutput, error) {
    var output wafv2.ListResourcesForWebACLOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2ListRuleGroupsResult struct {
	Result workflow.Future
}

func (r *Wafv2ListRuleGroupsResult) Get(ctx workflow.Context) (*wafv2.ListRuleGroupsOutput, error) {
    var output wafv2.ListRuleGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2ListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *Wafv2ListTagsForResourceResult) Get(ctx workflow.Context) (*wafv2.ListTagsForResourceOutput, error) {
    var output wafv2.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2ListWebACLsResult struct {
	Result workflow.Future
}

func (r *Wafv2ListWebACLsResult) Get(ctx workflow.Context) (*wafv2.ListWebACLsOutput, error) {
    var output wafv2.ListWebACLsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2PutLoggingConfigurationResult struct {
	Result workflow.Future
}

func (r *Wafv2PutLoggingConfigurationResult) Get(ctx workflow.Context) (*wafv2.PutLoggingConfigurationOutput, error) {
    var output wafv2.PutLoggingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2PutPermissionPolicyResult struct {
	Result workflow.Future
}

func (r *Wafv2PutPermissionPolicyResult) Get(ctx workflow.Context) (*wafv2.PutPermissionPolicyOutput, error) {
    var output wafv2.PutPermissionPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2TagResourceResult struct {
	Result workflow.Future
}

func (r *Wafv2TagResourceResult) Get(ctx workflow.Context) (*wafv2.TagResourceOutput, error) {
    var output wafv2.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2UntagResourceResult struct {
	Result workflow.Future
}

func (r *Wafv2UntagResourceResult) Get(ctx workflow.Context) (*wafv2.UntagResourceOutput, error) {
    var output wafv2.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2UpdateIPSetResult struct {
	Result workflow.Future
}

func (r *Wafv2UpdateIPSetResult) Get(ctx workflow.Context) (*wafv2.UpdateIPSetOutput, error) {
    var output wafv2.UpdateIPSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2UpdateRegexPatternSetResult struct {
	Result workflow.Future
}

func (r *Wafv2UpdateRegexPatternSetResult) Get(ctx workflow.Context) (*wafv2.UpdateRegexPatternSetOutput, error) {
    var output wafv2.UpdateRegexPatternSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2UpdateRuleGroupResult struct {
	Result workflow.Future
}

func (r *Wafv2UpdateRuleGroupResult) Get(ctx workflow.Context) (*wafv2.UpdateRuleGroupOutput, error) {
    var output wafv2.UpdateRuleGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Wafv2UpdateWebACLResult struct {
	Result workflow.Future
}

func (r *Wafv2UpdateWebACLResult) Get(ctx workflow.Context) (*wafv2.UpdateWebACLOutput, error) {
    var output wafv2.UpdateWebACLOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WAFV2Stub struct {
    activities awsactivities.WAFV2Activities
}

func NewWAFV2Stub() WAFV2Client {
    return &WAFV2Stub{}
}

func (a *WAFV2Stub) AssociateWebACL(ctx workflow.Context, input *wafv2.AssociateWebACLInput) (*wafv2.AssociateWebACLOutput, error) {
    var output wafv2.AssociateWebACLOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateWebACL, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) AssociateWebACLAsync(ctx workflow.Context, input *wafv2.AssociateWebACLInput) *Wafv2AssociateWebACLResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateWebACL, input)
    return &Wafv2AssociateWebACLResult{Result: future}
}

func (a *WAFV2Stub) CheckCapacity(ctx workflow.Context, input *wafv2.CheckCapacityInput) (*wafv2.CheckCapacityOutput, error) {
    var output wafv2.CheckCapacityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CheckCapacity, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) CheckCapacityAsync(ctx workflow.Context, input *wafv2.CheckCapacityInput) *Wafv2CheckCapacityResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CheckCapacity, input)
    return &Wafv2CheckCapacityResult{Result: future}
}

func (a *WAFV2Stub) CreateIPSet(ctx workflow.Context, input *wafv2.CreateIPSetInput) (*wafv2.CreateIPSetOutput, error) {
    var output wafv2.CreateIPSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateIPSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) CreateIPSetAsync(ctx workflow.Context, input *wafv2.CreateIPSetInput) *Wafv2CreateIPSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateIPSet, input)
    return &Wafv2CreateIPSetResult{Result: future}
}

func (a *WAFV2Stub) CreateRegexPatternSet(ctx workflow.Context, input *wafv2.CreateRegexPatternSetInput) (*wafv2.CreateRegexPatternSetOutput, error) {
    var output wafv2.CreateRegexPatternSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateRegexPatternSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) CreateRegexPatternSetAsync(ctx workflow.Context, input *wafv2.CreateRegexPatternSetInput) *Wafv2CreateRegexPatternSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateRegexPatternSet, input)
    return &Wafv2CreateRegexPatternSetResult{Result: future}
}

func (a *WAFV2Stub) CreateRuleGroup(ctx workflow.Context, input *wafv2.CreateRuleGroupInput) (*wafv2.CreateRuleGroupOutput, error) {
    var output wafv2.CreateRuleGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateRuleGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) CreateRuleGroupAsync(ctx workflow.Context, input *wafv2.CreateRuleGroupInput) *Wafv2CreateRuleGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateRuleGroup, input)
    return &Wafv2CreateRuleGroupResult{Result: future}
}

func (a *WAFV2Stub) CreateWebACL(ctx workflow.Context, input *wafv2.CreateWebACLInput) (*wafv2.CreateWebACLOutput, error) {
    var output wafv2.CreateWebACLOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateWebACL, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) CreateWebACLAsync(ctx workflow.Context, input *wafv2.CreateWebACLInput) *Wafv2CreateWebACLResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateWebACL, input)
    return &Wafv2CreateWebACLResult{Result: future}
}

func (a *WAFV2Stub) DeleteFirewallManagerRuleGroups(ctx workflow.Context, input *wafv2.DeleteFirewallManagerRuleGroupsInput) (*wafv2.DeleteFirewallManagerRuleGroupsOutput, error) {
    var output wafv2.DeleteFirewallManagerRuleGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFirewallManagerRuleGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) DeleteFirewallManagerRuleGroupsAsync(ctx workflow.Context, input *wafv2.DeleteFirewallManagerRuleGroupsInput) *Wafv2DeleteFirewallManagerRuleGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteFirewallManagerRuleGroups, input)
    return &Wafv2DeleteFirewallManagerRuleGroupsResult{Result: future}
}

func (a *WAFV2Stub) DeleteIPSet(ctx workflow.Context, input *wafv2.DeleteIPSetInput) (*wafv2.DeleteIPSetOutput, error) {
    var output wafv2.DeleteIPSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteIPSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) DeleteIPSetAsync(ctx workflow.Context, input *wafv2.DeleteIPSetInput) *Wafv2DeleteIPSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteIPSet, input)
    return &Wafv2DeleteIPSetResult{Result: future}
}

func (a *WAFV2Stub) DeleteLoggingConfiguration(ctx workflow.Context, input *wafv2.DeleteLoggingConfigurationInput) (*wafv2.DeleteLoggingConfigurationOutput, error) {
    var output wafv2.DeleteLoggingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteLoggingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) DeleteLoggingConfigurationAsync(ctx workflow.Context, input *wafv2.DeleteLoggingConfigurationInput) *Wafv2DeleteLoggingConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteLoggingConfiguration, input)
    return &Wafv2DeleteLoggingConfigurationResult{Result: future}
}

func (a *WAFV2Stub) DeletePermissionPolicy(ctx workflow.Context, input *wafv2.DeletePermissionPolicyInput) (*wafv2.DeletePermissionPolicyOutput, error) {
    var output wafv2.DeletePermissionPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePermissionPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) DeletePermissionPolicyAsync(ctx workflow.Context, input *wafv2.DeletePermissionPolicyInput) *Wafv2DeletePermissionPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeletePermissionPolicy, input)
    return &Wafv2DeletePermissionPolicyResult{Result: future}
}

func (a *WAFV2Stub) DeleteRegexPatternSet(ctx workflow.Context, input *wafv2.DeleteRegexPatternSetInput) (*wafv2.DeleteRegexPatternSetOutput, error) {
    var output wafv2.DeleteRegexPatternSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRegexPatternSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) DeleteRegexPatternSetAsync(ctx workflow.Context, input *wafv2.DeleteRegexPatternSetInput) *Wafv2DeleteRegexPatternSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteRegexPatternSet, input)
    return &Wafv2DeleteRegexPatternSetResult{Result: future}
}

func (a *WAFV2Stub) DeleteRuleGroup(ctx workflow.Context, input *wafv2.DeleteRuleGroupInput) (*wafv2.DeleteRuleGroupOutput, error) {
    var output wafv2.DeleteRuleGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRuleGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) DeleteRuleGroupAsync(ctx workflow.Context, input *wafv2.DeleteRuleGroupInput) *Wafv2DeleteRuleGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteRuleGroup, input)
    return &Wafv2DeleteRuleGroupResult{Result: future}
}

func (a *WAFV2Stub) DeleteWebACL(ctx workflow.Context, input *wafv2.DeleteWebACLInput) (*wafv2.DeleteWebACLOutput, error) {
    var output wafv2.DeleteWebACLOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteWebACL, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) DeleteWebACLAsync(ctx workflow.Context, input *wafv2.DeleteWebACLInput) *Wafv2DeleteWebACLResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteWebACL, input)
    return &Wafv2DeleteWebACLResult{Result: future}
}

func (a *WAFV2Stub) DescribeManagedRuleGroup(ctx workflow.Context, input *wafv2.DescribeManagedRuleGroupInput) (*wafv2.DescribeManagedRuleGroupOutput, error) {
    var output wafv2.DescribeManagedRuleGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeManagedRuleGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) DescribeManagedRuleGroupAsync(ctx workflow.Context, input *wafv2.DescribeManagedRuleGroupInput) *Wafv2DescribeManagedRuleGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeManagedRuleGroup, input)
    return &Wafv2DescribeManagedRuleGroupResult{Result: future}
}

func (a *WAFV2Stub) DisassociateWebACL(ctx workflow.Context, input *wafv2.DisassociateWebACLInput) (*wafv2.DisassociateWebACLOutput, error) {
    var output wafv2.DisassociateWebACLOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateWebACL, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) DisassociateWebACLAsync(ctx workflow.Context, input *wafv2.DisassociateWebACLInput) *Wafv2DisassociateWebACLResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateWebACL, input)
    return &Wafv2DisassociateWebACLResult{Result: future}
}

func (a *WAFV2Stub) GetIPSet(ctx workflow.Context, input *wafv2.GetIPSetInput) (*wafv2.GetIPSetOutput, error) {
    var output wafv2.GetIPSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetIPSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) GetIPSetAsync(ctx workflow.Context, input *wafv2.GetIPSetInput) *Wafv2GetIPSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetIPSet, input)
    return &Wafv2GetIPSetResult{Result: future}
}

func (a *WAFV2Stub) GetLoggingConfiguration(ctx workflow.Context, input *wafv2.GetLoggingConfigurationInput) (*wafv2.GetLoggingConfigurationOutput, error) {
    var output wafv2.GetLoggingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetLoggingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) GetLoggingConfigurationAsync(ctx workflow.Context, input *wafv2.GetLoggingConfigurationInput) *Wafv2GetLoggingConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetLoggingConfiguration, input)
    return &Wafv2GetLoggingConfigurationResult{Result: future}
}

func (a *WAFV2Stub) GetPermissionPolicy(ctx workflow.Context, input *wafv2.GetPermissionPolicyInput) (*wafv2.GetPermissionPolicyOutput, error) {
    var output wafv2.GetPermissionPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPermissionPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) GetPermissionPolicyAsync(ctx workflow.Context, input *wafv2.GetPermissionPolicyInput) *Wafv2GetPermissionPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetPermissionPolicy, input)
    return &Wafv2GetPermissionPolicyResult{Result: future}
}

func (a *WAFV2Stub) GetRateBasedStatementManagedKeys(ctx workflow.Context, input *wafv2.GetRateBasedStatementManagedKeysInput) (*wafv2.GetRateBasedStatementManagedKeysOutput, error) {
    var output wafv2.GetRateBasedStatementManagedKeysOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRateBasedStatementManagedKeys, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) GetRateBasedStatementManagedKeysAsync(ctx workflow.Context, input *wafv2.GetRateBasedStatementManagedKeysInput) *Wafv2GetRateBasedStatementManagedKeysResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetRateBasedStatementManagedKeys, input)
    return &Wafv2GetRateBasedStatementManagedKeysResult{Result: future}
}

func (a *WAFV2Stub) GetRegexPatternSet(ctx workflow.Context, input *wafv2.GetRegexPatternSetInput) (*wafv2.GetRegexPatternSetOutput, error) {
    var output wafv2.GetRegexPatternSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRegexPatternSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) GetRegexPatternSetAsync(ctx workflow.Context, input *wafv2.GetRegexPatternSetInput) *Wafv2GetRegexPatternSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetRegexPatternSet, input)
    return &Wafv2GetRegexPatternSetResult{Result: future}
}

func (a *WAFV2Stub) GetRuleGroup(ctx workflow.Context, input *wafv2.GetRuleGroupInput) (*wafv2.GetRuleGroupOutput, error) {
    var output wafv2.GetRuleGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRuleGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) GetRuleGroupAsync(ctx workflow.Context, input *wafv2.GetRuleGroupInput) *Wafv2GetRuleGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetRuleGroup, input)
    return &Wafv2GetRuleGroupResult{Result: future}
}

func (a *WAFV2Stub) GetSampledRequests(ctx workflow.Context, input *wafv2.GetSampledRequestsInput) (*wafv2.GetSampledRequestsOutput, error) {
    var output wafv2.GetSampledRequestsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSampledRequests, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) GetSampledRequestsAsync(ctx workflow.Context, input *wafv2.GetSampledRequestsInput) *Wafv2GetSampledRequestsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSampledRequests, input)
    return &Wafv2GetSampledRequestsResult{Result: future}
}

func (a *WAFV2Stub) GetWebACL(ctx workflow.Context, input *wafv2.GetWebACLInput) (*wafv2.GetWebACLOutput, error) {
    var output wafv2.GetWebACLOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetWebACL, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) GetWebACLAsync(ctx workflow.Context, input *wafv2.GetWebACLInput) *Wafv2GetWebACLResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetWebACL, input)
    return &Wafv2GetWebACLResult{Result: future}
}

func (a *WAFV2Stub) GetWebACLForResource(ctx workflow.Context, input *wafv2.GetWebACLForResourceInput) (*wafv2.GetWebACLForResourceOutput, error) {
    var output wafv2.GetWebACLForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetWebACLForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) GetWebACLForResourceAsync(ctx workflow.Context, input *wafv2.GetWebACLForResourceInput) *Wafv2GetWebACLForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetWebACLForResource, input)
    return &Wafv2GetWebACLForResourceResult{Result: future}
}

func (a *WAFV2Stub) ListAvailableManagedRuleGroups(ctx workflow.Context, input *wafv2.ListAvailableManagedRuleGroupsInput) (*wafv2.ListAvailableManagedRuleGroupsOutput, error) {
    var output wafv2.ListAvailableManagedRuleGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAvailableManagedRuleGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) ListAvailableManagedRuleGroupsAsync(ctx workflow.Context, input *wafv2.ListAvailableManagedRuleGroupsInput) *Wafv2ListAvailableManagedRuleGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAvailableManagedRuleGroups, input)
    return &Wafv2ListAvailableManagedRuleGroupsResult{Result: future}
}

func (a *WAFV2Stub) ListIPSets(ctx workflow.Context, input *wafv2.ListIPSetsInput) (*wafv2.ListIPSetsOutput, error) {
    var output wafv2.ListIPSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListIPSets, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) ListIPSetsAsync(ctx workflow.Context, input *wafv2.ListIPSetsInput) *Wafv2ListIPSetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListIPSets, input)
    return &Wafv2ListIPSetsResult{Result: future}
}

func (a *WAFV2Stub) ListLoggingConfigurations(ctx workflow.Context, input *wafv2.ListLoggingConfigurationsInput) (*wafv2.ListLoggingConfigurationsOutput, error) {
    var output wafv2.ListLoggingConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListLoggingConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) ListLoggingConfigurationsAsync(ctx workflow.Context, input *wafv2.ListLoggingConfigurationsInput) *Wafv2ListLoggingConfigurationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListLoggingConfigurations, input)
    return &Wafv2ListLoggingConfigurationsResult{Result: future}
}

func (a *WAFV2Stub) ListRegexPatternSets(ctx workflow.Context, input *wafv2.ListRegexPatternSetsInput) (*wafv2.ListRegexPatternSetsOutput, error) {
    var output wafv2.ListRegexPatternSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRegexPatternSets, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) ListRegexPatternSetsAsync(ctx workflow.Context, input *wafv2.ListRegexPatternSetsInput) *Wafv2ListRegexPatternSetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListRegexPatternSets, input)
    return &Wafv2ListRegexPatternSetsResult{Result: future}
}

func (a *WAFV2Stub) ListResourcesForWebACL(ctx workflow.Context, input *wafv2.ListResourcesForWebACLInput) (*wafv2.ListResourcesForWebACLOutput, error) {
    var output wafv2.ListResourcesForWebACLOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListResourcesForWebACL, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) ListResourcesForWebACLAsync(ctx workflow.Context, input *wafv2.ListResourcesForWebACLInput) *Wafv2ListResourcesForWebACLResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListResourcesForWebACL, input)
    return &Wafv2ListResourcesForWebACLResult{Result: future}
}

func (a *WAFV2Stub) ListRuleGroups(ctx workflow.Context, input *wafv2.ListRuleGroupsInput) (*wafv2.ListRuleGroupsOutput, error) {
    var output wafv2.ListRuleGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRuleGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) ListRuleGroupsAsync(ctx workflow.Context, input *wafv2.ListRuleGroupsInput) *Wafv2ListRuleGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListRuleGroups, input)
    return &Wafv2ListRuleGroupsResult{Result: future}
}

func (a *WAFV2Stub) ListTagsForResource(ctx workflow.Context, input *wafv2.ListTagsForResourceInput) (*wafv2.ListTagsForResourceOutput, error) {
    var output wafv2.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) ListTagsForResourceAsync(ctx workflow.Context, input *wafv2.ListTagsForResourceInput) *Wafv2ListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &Wafv2ListTagsForResourceResult{Result: future}
}

func (a *WAFV2Stub) ListWebACLs(ctx workflow.Context, input *wafv2.ListWebACLsInput) (*wafv2.ListWebACLsOutput, error) {
    var output wafv2.ListWebACLsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListWebACLs, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) ListWebACLsAsync(ctx workflow.Context, input *wafv2.ListWebACLsInput) *Wafv2ListWebACLsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListWebACLs, input)
    return &Wafv2ListWebACLsResult{Result: future}
}

func (a *WAFV2Stub) PutLoggingConfiguration(ctx workflow.Context, input *wafv2.PutLoggingConfigurationInput) (*wafv2.PutLoggingConfigurationOutput, error) {
    var output wafv2.PutLoggingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutLoggingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) PutLoggingConfigurationAsync(ctx workflow.Context, input *wafv2.PutLoggingConfigurationInput) *Wafv2PutLoggingConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutLoggingConfiguration, input)
    return &Wafv2PutLoggingConfigurationResult{Result: future}
}

func (a *WAFV2Stub) PutPermissionPolicy(ctx workflow.Context, input *wafv2.PutPermissionPolicyInput) (*wafv2.PutPermissionPolicyOutput, error) {
    var output wafv2.PutPermissionPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutPermissionPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) PutPermissionPolicyAsync(ctx workflow.Context, input *wafv2.PutPermissionPolicyInput) *Wafv2PutPermissionPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutPermissionPolicy, input)
    return &Wafv2PutPermissionPolicyResult{Result: future}
}

func (a *WAFV2Stub) TagResource(ctx workflow.Context, input *wafv2.TagResourceInput) (*wafv2.TagResourceOutput, error) {
    var output wafv2.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) TagResourceAsync(ctx workflow.Context, input *wafv2.TagResourceInput) *Wafv2TagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &Wafv2TagResourceResult{Result: future}
}

func (a *WAFV2Stub) UntagResource(ctx workflow.Context, input *wafv2.UntagResourceInput) (*wafv2.UntagResourceOutput, error) {
    var output wafv2.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) UntagResourceAsync(ctx workflow.Context, input *wafv2.UntagResourceInput) *Wafv2UntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &Wafv2UntagResourceResult{Result: future}
}

func (a *WAFV2Stub) UpdateIPSet(ctx workflow.Context, input *wafv2.UpdateIPSetInput) (*wafv2.UpdateIPSetOutput, error) {
    var output wafv2.UpdateIPSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateIPSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) UpdateIPSetAsync(ctx workflow.Context, input *wafv2.UpdateIPSetInput) *Wafv2UpdateIPSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateIPSet, input)
    return &Wafv2UpdateIPSetResult{Result: future}
}

func (a *WAFV2Stub) UpdateRegexPatternSet(ctx workflow.Context, input *wafv2.UpdateRegexPatternSetInput) (*wafv2.UpdateRegexPatternSetOutput, error) {
    var output wafv2.UpdateRegexPatternSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateRegexPatternSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) UpdateRegexPatternSetAsync(ctx workflow.Context, input *wafv2.UpdateRegexPatternSetInput) *Wafv2UpdateRegexPatternSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateRegexPatternSet, input)
    return &Wafv2UpdateRegexPatternSetResult{Result: future}
}

func (a *WAFV2Stub) UpdateRuleGroup(ctx workflow.Context, input *wafv2.UpdateRuleGroupInput) (*wafv2.UpdateRuleGroupOutput, error) {
    var output wafv2.UpdateRuleGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateRuleGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) UpdateRuleGroupAsync(ctx workflow.Context, input *wafv2.UpdateRuleGroupInput) *Wafv2UpdateRuleGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateRuleGroup, input)
    return &Wafv2UpdateRuleGroupResult{Result: future}
}

func (a *WAFV2Stub) UpdateWebACL(ctx workflow.Context, input *wafv2.UpdateWebACLInput) (*wafv2.UpdateWebACLOutput, error) {
    var output wafv2.UpdateWebACLOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateWebACL, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFV2Stub) UpdateWebACLAsync(ctx workflow.Context, input *wafv2.UpdateWebACLInput) *Wafv2UpdateWebACLResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateWebACL, input)
    return &Wafv2UpdateWebACLResult{Result: future}
}
