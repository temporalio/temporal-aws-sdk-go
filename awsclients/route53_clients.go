package awsclients

import (
	"github.com/aws/aws-sdk-go/service/route53"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type Route53Client interface {
    AssociateVPCWithHostedZone(ctx workflow.Context, input *route53.AssociateVPCWithHostedZoneInput) (*route53.AssociateVPCWithHostedZoneOutput, error)
    AssociateVPCWithHostedZoneAsync(ctx workflow.Context, input *route53.AssociateVPCWithHostedZoneInput) *Route53AssociateVPCWithHostedZoneResult

    ChangeResourceRecordSets(ctx workflow.Context, input *route53.ChangeResourceRecordSetsInput) (*route53.ChangeResourceRecordSetsOutput, error)
    ChangeResourceRecordSetsAsync(ctx workflow.Context, input *route53.ChangeResourceRecordSetsInput) *Route53ChangeResourceRecordSetsResult

    ChangeTagsForResource(ctx workflow.Context, input *route53.ChangeTagsForResourceInput) (*route53.ChangeTagsForResourceOutput, error)
    ChangeTagsForResourceAsync(ctx workflow.Context, input *route53.ChangeTagsForResourceInput) *Route53ChangeTagsForResourceResult

    CreateHealthCheck(ctx workflow.Context, input *route53.CreateHealthCheckInput) (*route53.CreateHealthCheckOutput, error)
    CreateHealthCheckAsync(ctx workflow.Context, input *route53.CreateHealthCheckInput) *Route53CreateHealthCheckResult

    CreateHostedZone(ctx workflow.Context, input *route53.CreateHostedZoneInput) (*route53.CreateHostedZoneOutput, error)
    CreateHostedZoneAsync(ctx workflow.Context, input *route53.CreateHostedZoneInput) *Route53CreateHostedZoneResult

    CreateQueryLoggingConfig(ctx workflow.Context, input *route53.CreateQueryLoggingConfigInput) (*route53.CreateQueryLoggingConfigOutput, error)
    CreateQueryLoggingConfigAsync(ctx workflow.Context, input *route53.CreateQueryLoggingConfigInput) *Route53CreateQueryLoggingConfigResult

    CreateReusableDelegationSet(ctx workflow.Context, input *route53.CreateReusableDelegationSetInput) (*route53.CreateReusableDelegationSetOutput, error)
    CreateReusableDelegationSetAsync(ctx workflow.Context, input *route53.CreateReusableDelegationSetInput) *Route53CreateReusableDelegationSetResult

    CreateTrafficPolicy(ctx workflow.Context, input *route53.CreateTrafficPolicyInput) (*route53.CreateTrafficPolicyOutput, error)
    CreateTrafficPolicyAsync(ctx workflow.Context, input *route53.CreateTrafficPolicyInput) *Route53CreateTrafficPolicyResult

    CreateTrafficPolicyInstance(ctx workflow.Context, input *route53.CreateTrafficPolicyInstanceInput) (*route53.CreateTrafficPolicyInstanceOutput, error)
    CreateTrafficPolicyInstanceAsync(ctx workflow.Context, input *route53.CreateTrafficPolicyInstanceInput) *Route53CreateTrafficPolicyInstanceResult

    CreateTrafficPolicyVersion(ctx workflow.Context, input *route53.CreateTrafficPolicyVersionInput) (*route53.CreateTrafficPolicyVersionOutput, error)
    CreateTrafficPolicyVersionAsync(ctx workflow.Context, input *route53.CreateTrafficPolicyVersionInput) *Route53CreateTrafficPolicyVersionResult

    CreateVPCAssociationAuthorization(ctx workflow.Context, input *route53.CreateVPCAssociationAuthorizationInput) (*route53.CreateVPCAssociationAuthorizationOutput, error)
    CreateVPCAssociationAuthorizationAsync(ctx workflow.Context, input *route53.CreateVPCAssociationAuthorizationInput) *Route53CreateVPCAssociationAuthorizationResult

    DeleteHealthCheck(ctx workflow.Context, input *route53.DeleteHealthCheckInput) (*route53.DeleteHealthCheckOutput, error)
    DeleteHealthCheckAsync(ctx workflow.Context, input *route53.DeleteHealthCheckInput) *Route53DeleteHealthCheckResult

    DeleteHostedZone(ctx workflow.Context, input *route53.DeleteHostedZoneInput) (*route53.DeleteHostedZoneOutput, error)
    DeleteHostedZoneAsync(ctx workflow.Context, input *route53.DeleteHostedZoneInput) *Route53DeleteHostedZoneResult

    DeleteQueryLoggingConfig(ctx workflow.Context, input *route53.DeleteQueryLoggingConfigInput) (*route53.DeleteQueryLoggingConfigOutput, error)
    DeleteQueryLoggingConfigAsync(ctx workflow.Context, input *route53.DeleteQueryLoggingConfigInput) *Route53DeleteQueryLoggingConfigResult

    DeleteReusableDelegationSet(ctx workflow.Context, input *route53.DeleteReusableDelegationSetInput) (*route53.DeleteReusableDelegationSetOutput, error)
    DeleteReusableDelegationSetAsync(ctx workflow.Context, input *route53.DeleteReusableDelegationSetInput) *Route53DeleteReusableDelegationSetResult

    DeleteTrafficPolicy(ctx workflow.Context, input *route53.DeleteTrafficPolicyInput) (*route53.DeleteTrafficPolicyOutput, error)
    DeleteTrafficPolicyAsync(ctx workflow.Context, input *route53.DeleteTrafficPolicyInput) *Route53DeleteTrafficPolicyResult

    DeleteTrafficPolicyInstance(ctx workflow.Context, input *route53.DeleteTrafficPolicyInstanceInput) (*route53.DeleteTrafficPolicyInstanceOutput, error)
    DeleteTrafficPolicyInstanceAsync(ctx workflow.Context, input *route53.DeleteTrafficPolicyInstanceInput) *Route53DeleteTrafficPolicyInstanceResult

    DeleteVPCAssociationAuthorization(ctx workflow.Context, input *route53.DeleteVPCAssociationAuthorizationInput) (*route53.DeleteVPCAssociationAuthorizationOutput, error)
    DeleteVPCAssociationAuthorizationAsync(ctx workflow.Context, input *route53.DeleteVPCAssociationAuthorizationInput) *Route53DeleteVPCAssociationAuthorizationResult

    DisassociateVPCFromHostedZone(ctx workflow.Context, input *route53.DisassociateVPCFromHostedZoneInput) (*route53.DisassociateVPCFromHostedZoneOutput, error)
    DisassociateVPCFromHostedZoneAsync(ctx workflow.Context, input *route53.DisassociateVPCFromHostedZoneInput) *Route53DisassociateVPCFromHostedZoneResult

    GetAccountLimit(ctx workflow.Context, input *route53.GetAccountLimitInput) (*route53.GetAccountLimitOutput, error)
    GetAccountLimitAsync(ctx workflow.Context, input *route53.GetAccountLimitInput) *Route53GetAccountLimitResult

    GetChange(ctx workflow.Context, input *route53.GetChangeInput) (*route53.GetChangeOutput, error)
    GetChangeAsync(ctx workflow.Context, input *route53.GetChangeInput) *Route53GetChangeResult

    GetCheckerIpRanges(ctx workflow.Context, input *route53.GetCheckerIpRangesInput) (*route53.GetCheckerIpRangesOutput, error)
    GetCheckerIpRangesAsync(ctx workflow.Context, input *route53.GetCheckerIpRangesInput) *Route53GetCheckerIpRangesResult

    GetGeoLocation(ctx workflow.Context, input *route53.GetGeoLocationInput) (*route53.GetGeoLocationOutput, error)
    GetGeoLocationAsync(ctx workflow.Context, input *route53.GetGeoLocationInput) *Route53GetGeoLocationResult

    GetHealthCheck(ctx workflow.Context, input *route53.GetHealthCheckInput) (*route53.GetHealthCheckOutput, error)
    GetHealthCheckAsync(ctx workflow.Context, input *route53.GetHealthCheckInput) *Route53GetHealthCheckResult

    GetHealthCheckCount(ctx workflow.Context, input *route53.GetHealthCheckCountInput) (*route53.GetHealthCheckCountOutput, error)
    GetHealthCheckCountAsync(ctx workflow.Context, input *route53.GetHealthCheckCountInput) *Route53GetHealthCheckCountResult

    GetHealthCheckLastFailureReason(ctx workflow.Context, input *route53.GetHealthCheckLastFailureReasonInput) (*route53.GetHealthCheckLastFailureReasonOutput, error)
    GetHealthCheckLastFailureReasonAsync(ctx workflow.Context, input *route53.GetHealthCheckLastFailureReasonInput) *Route53GetHealthCheckLastFailureReasonResult

    GetHealthCheckStatus(ctx workflow.Context, input *route53.GetHealthCheckStatusInput) (*route53.GetHealthCheckStatusOutput, error)
    GetHealthCheckStatusAsync(ctx workflow.Context, input *route53.GetHealthCheckStatusInput) *Route53GetHealthCheckStatusResult

    GetHostedZone(ctx workflow.Context, input *route53.GetHostedZoneInput) (*route53.GetHostedZoneOutput, error)
    GetHostedZoneAsync(ctx workflow.Context, input *route53.GetHostedZoneInput) *Route53GetHostedZoneResult

    GetHostedZoneCount(ctx workflow.Context, input *route53.GetHostedZoneCountInput) (*route53.GetHostedZoneCountOutput, error)
    GetHostedZoneCountAsync(ctx workflow.Context, input *route53.GetHostedZoneCountInput) *Route53GetHostedZoneCountResult

    GetHostedZoneLimit(ctx workflow.Context, input *route53.GetHostedZoneLimitInput) (*route53.GetHostedZoneLimitOutput, error)
    GetHostedZoneLimitAsync(ctx workflow.Context, input *route53.GetHostedZoneLimitInput) *Route53GetHostedZoneLimitResult

    GetQueryLoggingConfig(ctx workflow.Context, input *route53.GetQueryLoggingConfigInput) (*route53.GetQueryLoggingConfigOutput, error)
    GetQueryLoggingConfigAsync(ctx workflow.Context, input *route53.GetQueryLoggingConfigInput) *Route53GetQueryLoggingConfigResult

    GetReusableDelegationSet(ctx workflow.Context, input *route53.GetReusableDelegationSetInput) (*route53.GetReusableDelegationSetOutput, error)
    GetReusableDelegationSetAsync(ctx workflow.Context, input *route53.GetReusableDelegationSetInput) *Route53GetReusableDelegationSetResult

    GetReusableDelegationSetLimit(ctx workflow.Context, input *route53.GetReusableDelegationSetLimitInput) (*route53.GetReusableDelegationSetLimitOutput, error)
    GetReusableDelegationSetLimitAsync(ctx workflow.Context, input *route53.GetReusableDelegationSetLimitInput) *Route53GetReusableDelegationSetLimitResult

    GetTrafficPolicy(ctx workflow.Context, input *route53.GetTrafficPolicyInput) (*route53.GetTrafficPolicyOutput, error)
    GetTrafficPolicyAsync(ctx workflow.Context, input *route53.GetTrafficPolicyInput) *Route53GetTrafficPolicyResult

    GetTrafficPolicyInstance(ctx workflow.Context, input *route53.GetTrafficPolicyInstanceInput) (*route53.GetTrafficPolicyInstanceOutput, error)
    GetTrafficPolicyInstanceAsync(ctx workflow.Context, input *route53.GetTrafficPolicyInstanceInput) *Route53GetTrafficPolicyInstanceResult

    GetTrafficPolicyInstanceCount(ctx workflow.Context, input *route53.GetTrafficPolicyInstanceCountInput) (*route53.GetTrafficPolicyInstanceCountOutput, error)
    GetTrafficPolicyInstanceCountAsync(ctx workflow.Context, input *route53.GetTrafficPolicyInstanceCountInput) *Route53GetTrafficPolicyInstanceCountResult

    ListGeoLocations(ctx workflow.Context, input *route53.ListGeoLocationsInput) (*route53.ListGeoLocationsOutput, error)
    ListGeoLocationsAsync(ctx workflow.Context, input *route53.ListGeoLocationsInput) *Route53ListGeoLocationsResult

    ListHealthChecks(ctx workflow.Context, input *route53.ListHealthChecksInput) (*route53.ListHealthChecksOutput, error)
    ListHealthChecksAsync(ctx workflow.Context, input *route53.ListHealthChecksInput) *Route53ListHealthChecksResult

    ListHostedZones(ctx workflow.Context, input *route53.ListHostedZonesInput) (*route53.ListHostedZonesOutput, error)
    ListHostedZonesAsync(ctx workflow.Context, input *route53.ListHostedZonesInput) *Route53ListHostedZonesResult

    ListHostedZonesByName(ctx workflow.Context, input *route53.ListHostedZonesByNameInput) (*route53.ListHostedZonesByNameOutput, error)
    ListHostedZonesByNameAsync(ctx workflow.Context, input *route53.ListHostedZonesByNameInput) *Route53ListHostedZonesByNameResult

    ListHostedZonesByVPC(ctx workflow.Context, input *route53.ListHostedZonesByVPCInput) (*route53.ListHostedZonesByVPCOutput, error)
    ListHostedZonesByVPCAsync(ctx workflow.Context, input *route53.ListHostedZonesByVPCInput) *Route53ListHostedZonesByVPCResult

    ListQueryLoggingConfigs(ctx workflow.Context, input *route53.ListQueryLoggingConfigsInput) (*route53.ListQueryLoggingConfigsOutput, error)
    ListQueryLoggingConfigsAsync(ctx workflow.Context, input *route53.ListQueryLoggingConfigsInput) *Route53ListQueryLoggingConfigsResult

    ListResourceRecordSets(ctx workflow.Context, input *route53.ListResourceRecordSetsInput) (*route53.ListResourceRecordSetsOutput, error)
    ListResourceRecordSetsAsync(ctx workflow.Context, input *route53.ListResourceRecordSetsInput) *Route53ListResourceRecordSetsResult

    ListReusableDelegationSets(ctx workflow.Context, input *route53.ListReusableDelegationSetsInput) (*route53.ListReusableDelegationSetsOutput, error)
    ListReusableDelegationSetsAsync(ctx workflow.Context, input *route53.ListReusableDelegationSetsInput) *Route53ListReusableDelegationSetsResult

    ListTagsForResource(ctx workflow.Context, input *route53.ListTagsForResourceInput) (*route53.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *route53.ListTagsForResourceInput) *Route53ListTagsForResourceResult

    ListTagsForResources(ctx workflow.Context, input *route53.ListTagsForResourcesInput) (*route53.ListTagsForResourcesOutput, error)
    ListTagsForResourcesAsync(ctx workflow.Context, input *route53.ListTagsForResourcesInput) *Route53ListTagsForResourcesResult

    ListTrafficPolicies(ctx workflow.Context, input *route53.ListTrafficPoliciesInput) (*route53.ListTrafficPoliciesOutput, error)
    ListTrafficPoliciesAsync(ctx workflow.Context, input *route53.ListTrafficPoliciesInput) *Route53ListTrafficPoliciesResult

    ListTrafficPolicyInstances(ctx workflow.Context, input *route53.ListTrafficPolicyInstancesInput) (*route53.ListTrafficPolicyInstancesOutput, error)
    ListTrafficPolicyInstancesAsync(ctx workflow.Context, input *route53.ListTrafficPolicyInstancesInput) *Route53ListTrafficPolicyInstancesResult

    ListTrafficPolicyInstancesByHostedZone(ctx workflow.Context, input *route53.ListTrafficPolicyInstancesByHostedZoneInput) (*route53.ListTrafficPolicyInstancesByHostedZoneOutput, error)
    ListTrafficPolicyInstancesByHostedZoneAsync(ctx workflow.Context, input *route53.ListTrafficPolicyInstancesByHostedZoneInput) *Route53ListTrafficPolicyInstancesByHostedZoneResult

    ListTrafficPolicyInstancesByPolicy(ctx workflow.Context, input *route53.ListTrafficPolicyInstancesByPolicyInput) (*route53.ListTrafficPolicyInstancesByPolicyOutput, error)
    ListTrafficPolicyInstancesByPolicyAsync(ctx workflow.Context, input *route53.ListTrafficPolicyInstancesByPolicyInput) *Route53ListTrafficPolicyInstancesByPolicyResult

    ListTrafficPolicyVersions(ctx workflow.Context, input *route53.ListTrafficPolicyVersionsInput) (*route53.ListTrafficPolicyVersionsOutput, error)
    ListTrafficPolicyVersionsAsync(ctx workflow.Context, input *route53.ListTrafficPolicyVersionsInput) *Route53ListTrafficPolicyVersionsResult

    ListVPCAssociationAuthorizations(ctx workflow.Context, input *route53.ListVPCAssociationAuthorizationsInput) (*route53.ListVPCAssociationAuthorizationsOutput, error)
    ListVPCAssociationAuthorizationsAsync(ctx workflow.Context, input *route53.ListVPCAssociationAuthorizationsInput) *Route53ListVPCAssociationAuthorizationsResult

    TestDNSAnswer(ctx workflow.Context, input *route53.TestDNSAnswerInput) (*route53.TestDNSAnswerOutput, error)
    TestDNSAnswerAsync(ctx workflow.Context, input *route53.TestDNSAnswerInput) *Route53TestDNSAnswerResult

    UpdateHealthCheck(ctx workflow.Context, input *route53.UpdateHealthCheckInput) (*route53.UpdateHealthCheckOutput, error)
    UpdateHealthCheckAsync(ctx workflow.Context, input *route53.UpdateHealthCheckInput) *Route53UpdateHealthCheckResult

    UpdateHostedZoneComment(ctx workflow.Context, input *route53.UpdateHostedZoneCommentInput) (*route53.UpdateHostedZoneCommentOutput, error)
    UpdateHostedZoneCommentAsync(ctx workflow.Context, input *route53.UpdateHostedZoneCommentInput) *Route53UpdateHostedZoneCommentResult

    UpdateTrafficPolicyComment(ctx workflow.Context, input *route53.UpdateTrafficPolicyCommentInput) (*route53.UpdateTrafficPolicyCommentOutput, error)
    UpdateTrafficPolicyCommentAsync(ctx workflow.Context, input *route53.UpdateTrafficPolicyCommentInput) *Route53UpdateTrafficPolicyCommentResult

    UpdateTrafficPolicyInstance(ctx workflow.Context, input *route53.UpdateTrafficPolicyInstanceInput) (*route53.UpdateTrafficPolicyInstanceOutput, error)
    UpdateTrafficPolicyInstanceAsync(ctx workflow.Context, input *route53.UpdateTrafficPolicyInstanceInput) *Route53UpdateTrafficPolicyInstanceResult

    WaitUntilResourceRecordSetsChanged(ctx workflow.Context, input *route53.GetChangeInput) error}

type Route53AssociateVPCWithHostedZoneResult struct {
	Result workflow.Future
}

func (r *Route53AssociateVPCWithHostedZoneResult) Get(ctx workflow.Context) (*route53.AssociateVPCWithHostedZoneOutput, error) {
    var output route53.AssociateVPCWithHostedZoneOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53ChangeResourceRecordSetsResult struct {
	Result workflow.Future
}

func (r *Route53ChangeResourceRecordSetsResult) Get(ctx workflow.Context) (*route53.ChangeResourceRecordSetsOutput, error) {
    var output route53.ChangeResourceRecordSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53ChangeTagsForResourceResult struct {
	Result workflow.Future
}

func (r *Route53ChangeTagsForResourceResult) Get(ctx workflow.Context) (*route53.ChangeTagsForResourceOutput, error) {
    var output route53.ChangeTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53CreateHealthCheckResult struct {
	Result workflow.Future
}

func (r *Route53CreateHealthCheckResult) Get(ctx workflow.Context) (*route53.CreateHealthCheckOutput, error) {
    var output route53.CreateHealthCheckOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53CreateHostedZoneResult struct {
	Result workflow.Future
}

func (r *Route53CreateHostedZoneResult) Get(ctx workflow.Context) (*route53.CreateHostedZoneOutput, error) {
    var output route53.CreateHostedZoneOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53CreateQueryLoggingConfigResult struct {
	Result workflow.Future
}

func (r *Route53CreateQueryLoggingConfigResult) Get(ctx workflow.Context) (*route53.CreateQueryLoggingConfigOutput, error) {
    var output route53.CreateQueryLoggingConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53CreateReusableDelegationSetResult struct {
	Result workflow.Future
}

func (r *Route53CreateReusableDelegationSetResult) Get(ctx workflow.Context) (*route53.CreateReusableDelegationSetOutput, error) {
    var output route53.CreateReusableDelegationSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53CreateTrafficPolicyResult struct {
	Result workflow.Future
}

func (r *Route53CreateTrafficPolicyResult) Get(ctx workflow.Context) (*route53.CreateTrafficPolicyOutput, error) {
    var output route53.CreateTrafficPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53CreateTrafficPolicyInstanceResult struct {
	Result workflow.Future
}

func (r *Route53CreateTrafficPolicyInstanceResult) Get(ctx workflow.Context) (*route53.CreateTrafficPolicyInstanceOutput, error) {
    var output route53.CreateTrafficPolicyInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53CreateTrafficPolicyVersionResult struct {
	Result workflow.Future
}

func (r *Route53CreateTrafficPolicyVersionResult) Get(ctx workflow.Context) (*route53.CreateTrafficPolicyVersionOutput, error) {
    var output route53.CreateTrafficPolicyVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53CreateVPCAssociationAuthorizationResult struct {
	Result workflow.Future
}

func (r *Route53CreateVPCAssociationAuthorizationResult) Get(ctx workflow.Context) (*route53.CreateVPCAssociationAuthorizationOutput, error) {
    var output route53.CreateVPCAssociationAuthorizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53DeleteHealthCheckResult struct {
	Result workflow.Future
}

func (r *Route53DeleteHealthCheckResult) Get(ctx workflow.Context) (*route53.DeleteHealthCheckOutput, error) {
    var output route53.DeleteHealthCheckOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53DeleteHostedZoneResult struct {
	Result workflow.Future
}

func (r *Route53DeleteHostedZoneResult) Get(ctx workflow.Context) (*route53.DeleteHostedZoneOutput, error) {
    var output route53.DeleteHostedZoneOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53DeleteQueryLoggingConfigResult struct {
	Result workflow.Future
}

func (r *Route53DeleteQueryLoggingConfigResult) Get(ctx workflow.Context) (*route53.DeleteQueryLoggingConfigOutput, error) {
    var output route53.DeleteQueryLoggingConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53DeleteReusableDelegationSetResult struct {
	Result workflow.Future
}

func (r *Route53DeleteReusableDelegationSetResult) Get(ctx workflow.Context) (*route53.DeleteReusableDelegationSetOutput, error) {
    var output route53.DeleteReusableDelegationSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53DeleteTrafficPolicyResult struct {
	Result workflow.Future
}

func (r *Route53DeleteTrafficPolicyResult) Get(ctx workflow.Context) (*route53.DeleteTrafficPolicyOutput, error) {
    var output route53.DeleteTrafficPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53DeleteTrafficPolicyInstanceResult struct {
	Result workflow.Future
}

func (r *Route53DeleteTrafficPolicyInstanceResult) Get(ctx workflow.Context) (*route53.DeleteTrafficPolicyInstanceOutput, error) {
    var output route53.DeleteTrafficPolicyInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53DeleteVPCAssociationAuthorizationResult struct {
	Result workflow.Future
}

func (r *Route53DeleteVPCAssociationAuthorizationResult) Get(ctx workflow.Context) (*route53.DeleteVPCAssociationAuthorizationOutput, error) {
    var output route53.DeleteVPCAssociationAuthorizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53DisassociateVPCFromHostedZoneResult struct {
	Result workflow.Future
}

func (r *Route53DisassociateVPCFromHostedZoneResult) Get(ctx workflow.Context) (*route53.DisassociateVPCFromHostedZoneOutput, error) {
    var output route53.DisassociateVPCFromHostedZoneOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53GetAccountLimitResult struct {
	Result workflow.Future
}

func (r *Route53GetAccountLimitResult) Get(ctx workflow.Context) (*route53.GetAccountLimitOutput, error) {
    var output route53.GetAccountLimitOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53GetChangeResult struct {
	Result workflow.Future
}

func (r *Route53GetChangeResult) Get(ctx workflow.Context) (*route53.GetChangeOutput, error) {
    var output route53.GetChangeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53GetCheckerIpRangesResult struct {
	Result workflow.Future
}

func (r *Route53GetCheckerIpRangesResult) Get(ctx workflow.Context) (*route53.GetCheckerIpRangesOutput, error) {
    var output route53.GetCheckerIpRangesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53GetGeoLocationResult struct {
	Result workflow.Future
}

func (r *Route53GetGeoLocationResult) Get(ctx workflow.Context) (*route53.GetGeoLocationOutput, error) {
    var output route53.GetGeoLocationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53GetHealthCheckResult struct {
	Result workflow.Future
}

func (r *Route53GetHealthCheckResult) Get(ctx workflow.Context) (*route53.GetHealthCheckOutput, error) {
    var output route53.GetHealthCheckOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53GetHealthCheckCountResult struct {
	Result workflow.Future
}

func (r *Route53GetHealthCheckCountResult) Get(ctx workflow.Context) (*route53.GetHealthCheckCountOutput, error) {
    var output route53.GetHealthCheckCountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53GetHealthCheckLastFailureReasonResult struct {
	Result workflow.Future
}

func (r *Route53GetHealthCheckLastFailureReasonResult) Get(ctx workflow.Context) (*route53.GetHealthCheckLastFailureReasonOutput, error) {
    var output route53.GetHealthCheckLastFailureReasonOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53GetHealthCheckStatusResult struct {
	Result workflow.Future
}

func (r *Route53GetHealthCheckStatusResult) Get(ctx workflow.Context) (*route53.GetHealthCheckStatusOutput, error) {
    var output route53.GetHealthCheckStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53GetHostedZoneResult struct {
	Result workflow.Future
}

func (r *Route53GetHostedZoneResult) Get(ctx workflow.Context) (*route53.GetHostedZoneOutput, error) {
    var output route53.GetHostedZoneOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53GetHostedZoneCountResult struct {
	Result workflow.Future
}

func (r *Route53GetHostedZoneCountResult) Get(ctx workflow.Context) (*route53.GetHostedZoneCountOutput, error) {
    var output route53.GetHostedZoneCountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53GetHostedZoneLimitResult struct {
	Result workflow.Future
}

func (r *Route53GetHostedZoneLimitResult) Get(ctx workflow.Context) (*route53.GetHostedZoneLimitOutput, error) {
    var output route53.GetHostedZoneLimitOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53GetQueryLoggingConfigResult struct {
	Result workflow.Future
}

func (r *Route53GetQueryLoggingConfigResult) Get(ctx workflow.Context) (*route53.GetQueryLoggingConfigOutput, error) {
    var output route53.GetQueryLoggingConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53GetReusableDelegationSetResult struct {
	Result workflow.Future
}

func (r *Route53GetReusableDelegationSetResult) Get(ctx workflow.Context) (*route53.GetReusableDelegationSetOutput, error) {
    var output route53.GetReusableDelegationSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53GetReusableDelegationSetLimitResult struct {
	Result workflow.Future
}

func (r *Route53GetReusableDelegationSetLimitResult) Get(ctx workflow.Context) (*route53.GetReusableDelegationSetLimitOutput, error) {
    var output route53.GetReusableDelegationSetLimitOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53GetTrafficPolicyResult struct {
	Result workflow.Future
}

func (r *Route53GetTrafficPolicyResult) Get(ctx workflow.Context) (*route53.GetTrafficPolicyOutput, error) {
    var output route53.GetTrafficPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53GetTrafficPolicyInstanceResult struct {
	Result workflow.Future
}

func (r *Route53GetTrafficPolicyInstanceResult) Get(ctx workflow.Context) (*route53.GetTrafficPolicyInstanceOutput, error) {
    var output route53.GetTrafficPolicyInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53GetTrafficPolicyInstanceCountResult struct {
	Result workflow.Future
}

func (r *Route53GetTrafficPolicyInstanceCountResult) Get(ctx workflow.Context) (*route53.GetTrafficPolicyInstanceCountOutput, error) {
    var output route53.GetTrafficPolicyInstanceCountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53ListGeoLocationsResult struct {
	Result workflow.Future
}

func (r *Route53ListGeoLocationsResult) Get(ctx workflow.Context) (*route53.ListGeoLocationsOutput, error) {
    var output route53.ListGeoLocationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53ListHealthChecksResult struct {
	Result workflow.Future
}

func (r *Route53ListHealthChecksResult) Get(ctx workflow.Context) (*route53.ListHealthChecksOutput, error) {
    var output route53.ListHealthChecksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53ListHostedZonesResult struct {
	Result workflow.Future
}

func (r *Route53ListHostedZonesResult) Get(ctx workflow.Context) (*route53.ListHostedZonesOutput, error) {
    var output route53.ListHostedZonesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53ListHostedZonesByNameResult struct {
	Result workflow.Future
}

func (r *Route53ListHostedZonesByNameResult) Get(ctx workflow.Context) (*route53.ListHostedZonesByNameOutput, error) {
    var output route53.ListHostedZonesByNameOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53ListHostedZonesByVPCResult struct {
	Result workflow.Future
}

func (r *Route53ListHostedZonesByVPCResult) Get(ctx workflow.Context) (*route53.ListHostedZonesByVPCOutput, error) {
    var output route53.ListHostedZonesByVPCOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53ListQueryLoggingConfigsResult struct {
	Result workflow.Future
}

func (r *Route53ListQueryLoggingConfigsResult) Get(ctx workflow.Context) (*route53.ListQueryLoggingConfigsOutput, error) {
    var output route53.ListQueryLoggingConfigsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53ListResourceRecordSetsResult struct {
	Result workflow.Future
}

func (r *Route53ListResourceRecordSetsResult) Get(ctx workflow.Context) (*route53.ListResourceRecordSetsOutput, error) {
    var output route53.ListResourceRecordSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53ListReusableDelegationSetsResult struct {
	Result workflow.Future
}

func (r *Route53ListReusableDelegationSetsResult) Get(ctx workflow.Context) (*route53.ListReusableDelegationSetsOutput, error) {
    var output route53.ListReusableDelegationSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53ListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *Route53ListTagsForResourceResult) Get(ctx workflow.Context) (*route53.ListTagsForResourceOutput, error) {
    var output route53.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53ListTagsForResourcesResult struct {
	Result workflow.Future
}

func (r *Route53ListTagsForResourcesResult) Get(ctx workflow.Context) (*route53.ListTagsForResourcesOutput, error) {
    var output route53.ListTagsForResourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53ListTrafficPoliciesResult struct {
	Result workflow.Future
}

func (r *Route53ListTrafficPoliciesResult) Get(ctx workflow.Context) (*route53.ListTrafficPoliciesOutput, error) {
    var output route53.ListTrafficPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53ListTrafficPolicyInstancesResult struct {
	Result workflow.Future
}

func (r *Route53ListTrafficPolicyInstancesResult) Get(ctx workflow.Context) (*route53.ListTrafficPolicyInstancesOutput, error) {
    var output route53.ListTrafficPolicyInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53ListTrafficPolicyInstancesByHostedZoneResult struct {
	Result workflow.Future
}

func (r *Route53ListTrafficPolicyInstancesByHostedZoneResult) Get(ctx workflow.Context) (*route53.ListTrafficPolicyInstancesByHostedZoneOutput, error) {
    var output route53.ListTrafficPolicyInstancesByHostedZoneOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53ListTrafficPolicyInstancesByPolicyResult struct {
	Result workflow.Future
}

func (r *Route53ListTrafficPolicyInstancesByPolicyResult) Get(ctx workflow.Context) (*route53.ListTrafficPolicyInstancesByPolicyOutput, error) {
    var output route53.ListTrafficPolicyInstancesByPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53ListTrafficPolicyVersionsResult struct {
	Result workflow.Future
}

func (r *Route53ListTrafficPolicyVersionsResult) Get(ctx workflow.Context) (*route53.ListTrafficPolicyVersionsOutput, error) {
    var output route53.ListTrafficPolicyVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53ListVPCAssociationAuthorizationsResult struct {
	Result workflow.Future
}

func (r *Route53ListVPCAssociationAuthorizationsResult) Get(ctx workflow.Context) (*route53.ListVPCAssociationAuthorizationsOutput, error) {
    var output route53.ListVPCAssociationAuthorizationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53TestDNSAnswerResult struct {
	Result workflow.Future
}

func (r *Route53TestDNSAnswerResult) Get(ctx workflow.Context) (*route53.TestDNSAnswerOutput, error) {
    var output route53.TestDNSAnswerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53UpdateHealthCheckResult struct {
	Result workflow.Future
}

func (r *Route53UpdateHealthCheckResult) Get(ctx workflow.Context) (*route53.UpdateHealthCheckOutput, error) {
    var output route53.UpdateHealthCheckOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53UpdateHostedZoneCommentResult struct {
	Result workflow.Future
}

func (r *Route53UpdateHostedZoneCommentResult) Get(ctx workflow.Context) (*route53.UpdateHostedZoneCommentOutput, error) {
    var output route53.UpdateHostedZoneCommentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53UpdateTrafficPolicyCommentResult struct {
	Result workflow.Future
}

func (r *Route53UpdateTrafficPolicyCommentResult) Get(ctx workflow.Context) (*route53.UpdateTrafficPolicyCommentOutput, error) {
    var output route53.UpdateTrafficPolicyCommentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53UpdateTrafficPolicyInstanceResult struct {
	Result workflow.Future
}

func (r *Route53UpdateTrafficPolicyInstanceResult) Get(ctx workflow.Context) (*route53.UpdateTrafficPolicyInstanceOutput, error) {
    var output route53.UpdateTrafficPolicyInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53Stub struct {
    activities awsactivities.Route53Activities
}

func NewRoute53Stub() Route53Client {
    return &Route53Stub{}
}

func (a *Route53Stub) AssociateVPCWithHostedZone(ctx workflow.Context, input *route53.AssociateVPCWithHostedZoneInput) (*route53.AssociateVPCWithHostedZoneOutput, error) {
    var output route53.AssociateVPCWithHostedZoneOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateVPCWithHostedZone, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) AssociateVPCWithHostedZoneAsync(ctx workflow.Context, input *route53.AssociateVPCWithHostedZoneInput) *Route53AssociateVPCWithHostedZoneResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateVPCWithHostedZone, input)
    return &Route53AssociateVPCWithHostedZoneResult{Result: future}
}

func (a *Route53Stub) ChangeResourceRecordSets(ctx workflow.Context, input *route53.ChangeResourceRecordSetsInput) (*route53.ChangeResourceRecordSetsOutput, error) {
    var output route53.ChangeResourceRecordSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ChangeResourceRecordSets, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) ChangeResourceRecordSetsAsync(ctx workflow.Context, input *route53.ChangeResourceRecordSetsInput) *Route53ChangeResourceRecordSetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ChangeResourceRecordSets, input)
    return &Route53ChangeResourceRecordSetsResult{Result: future}
}

func (a *Route53Stub) ChangeTagsForResource(ctx workflow.Context, input *route53.ChangeTagsForResourceInput) (*route53.ChangeTagsForResourceOutput, error) {
    var output route53.ChangeTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ChangeTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) ChangeTagsForResourceAsync(ctx workflow.Context, input *route53.ChangeTagsForResourceInput) *Route53ChangeTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ChangeTagsForResource, input)
    return &Route53ChangeTagsForResourceResult{Result: future}
}

func (a *Route53Stub) CreateHealthCheck(ctx workflow.Context, input *route53.CreateHealthCheckInput) (*route53.CreateHealthCheckOutput, error) {
    var output route53.CreateHealthCheckOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateHealthCheck, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) CreateHealthCheckAsync(ctx workflow.Context, input *route53.CreateHealthCheckInput) *Route53CreateHealthCheckResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateHealthCheck, input)
    return &Route53CreateHealthCheckResult{Result: future}
}

func (a *Route53Stub) CreateHostedZone(ctx workflow.Context, input *route53.CreateHostedZoneInput) (*route53.CreateHostedZoneOutput, error) {
    var output route53.CreateHostedZoneOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateHostedZone, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) CreateHostedZoneAsync(ctx workflow.Context, input *route53.CreateHostedZoneInput) *Route53CreateHostedZoneResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateHostedZone, input)
    return &Route53CreateHostedZoneResult{Result: future}
}

func (a *Route53Stub) CreateQueryLoggingConfig(ctx workflow.Context, input *route53.CreateQueryLoggingConfigInput) (*route53.CreateQueryLoggingConfigOutput, error) {
    var output route53.CreateQueryLoggingConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateQueryLoggingConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) CreateQueryLoggingConfigAsync(ctx workflow.Context, input *route53.CreateQueryLoggingConfigInput) *Route53CreateQueryLoggingConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateQueryLoggingConfig, input)
    return &Route53CreateQueryLoggingConfigResult{Result: future}
}

func (a *Route53Stub) CreateReusableDelegationSet(ctx workflow.Context, input *route53.CreateReusableDelegationSetInput) (*route53.CreateReusableDelegationSetOutput, error) {
    var output route53.CreateReusableDelegationSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateReusableDelegationSet, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) CreateReusableDelegationSetAsync(ctx workflow.Context, input *route53.CreateReusableDelegationSetInput) *Route53CreateReusableDelegationSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateReusableDelegationSet, input)
    return &Route53CreateReusableDelegationSetResult{Result: future}
}

func (a *Route53Stub) CreateTrafficPolicy(ctx workflow.Context, input *route53.CreateTrafficPolicyInput) (*route53.CreateTrafficPolicyOutput, error) {
    var output route53.CreateTrafficPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTrafficPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) CreateTrafficPolicyAsync(ctx workflow.Context, input *route53.CreateTrafficPolicyInput) *Route53CreateTrafficPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTrafficPolicy, input)
    return &Route53CreateTrafficPolicyResult{Result: future}
}

func (a *Route53Stub) CreateTrafficPolicyInstance(ctx workflow.Context, input *route53.CreateTrafficPolicyInstanceInput) (*route53.CreateTrafficPolicyInstanceOutput, error) {
    var output route53.CreateTrafficPolicyInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTrafficPolicyInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) CreateTrafficPolicyInstanceAsync(ctx workflow.Context, input *route53.CreateTrafficPolicyInstanceInput) *Route53CreateTrafficPolicyInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTrafficPolicyInstance, input)
    return &Route53CreateTrafficPolicyInstanceResult{Result: future}
}

func (a *Route53Stub) CreateTrafficPolicyVersion(ctx workflow.Context, input *route53.CreateTrafficPolicyVersionInput) (*route53.CreateTrafficPolicyVersionOutput, error) {
    var output route53.CreateTrafficPolicyVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTrafficPolicyVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) CreateTrafficPolicyVersionAsync(ctx workflow.Context, input *route53.CreateTrafficPolicyVersionInput) *Route53CreateTrafficPolicyVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTrafficPolicyVersion, input)
    return &Route53CreateTrafficPolicyVersionResult{Result: future}
}

func (a *Route53Stub) CreateVPCAssociationAuthorization(ctx workflow.Context, input *route53.CreateVPCAssociationAuthorizationInput) (*route53.CreateVPCAssociationAuthorizationOutput, error) {
    var output route53.CreateVPCAssociationAuthorizationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateVPCAssociationAuthorization, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) CreateVPCAssociationAuthorizationAsync(ctx workflow.Context, input *route53.CreateVPCAssociationAuthorizationInput) *Route53CreateVPCAssociationAuthorizationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateVPCAssociationAuthorization, input)
    return &Route53CreateVPCAssociationAuthorizationResult{Result: future}
}

func (a *Route53Stub) DeleteHealthCheck(ctx workflow.Context, input *route53.DeleteHealthCheckInput) (*route53.DeleteHealthCheckOutput, error) {
    var output route53.DeleteHealthCheckOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteHealthCheck, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) DeleteHealthCheckAsync(ctx workflow.Context, input *route53.DeleteHealthCheckInput) *Route53DeleteHealthCheckResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteHealthCheck, input)
    return &Route53DeleteHealthCheckResult{Result: future}
}

func (a *Route53Stub) DeleteHostedZone(ctx workflow.Context, input *route53.DeleteHostedZoneInput) (*route53.DeleteHostedZoneOutput, error) {
    var output route53.DeleteHostedZoneOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteHostedZone, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) DeleteHostedZoneAsync(ctx workflow.Context, input *route53.DeleteHostedZoneInput) *Route53DeleteHostedZoneResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteHostedZone, input)
    return &Route53DeleteHostedZoneResult{Result: future}
}

func (a *Route53Stub) DeleteQueryLoggingConfig(ctx workflow.Context, input *route53.DeleteQueryLoggingConfigInput) (*route53.DeleteQueryLoggingConfigOutput, error) {
    var output route53.DeleteQueryLoggingConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteQueryLoggingConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) DeleteQueryLoggingConfigAsync(ctx workflow.Context, input *route53.DeleteQueryLoggingConfigInput) *Route53DeleteQueryLoggingConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteQueryLoggingConfig, input)
    return &Route53DeleteQueryLoggingConfigResult{Result: future}
}

func (a *Route53Stub) DeleteReusableDelegationSet(ctx workflow.Context, input *route53.DeleteReusableDelegationSetInput) (*route53.DeleteReusableDelegationSetOutput, error) {
    var output route53.DeleteReusableDelegationSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteReusableDelegationSet, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) DeleteReusableDelegationSetAsync(ctx workflow.Context, input *route53.DeleteReusableDelegationSetInput) *Route53DeleteReusableDelegationSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteReusableDelegationSet, input)
    return &Route53DeleteReusableDelegationSetResult{Result: future}
}

func (a *Route53Stub) DeleteTrafficPolicy(ctx workflow.Context, input *route53.DeleteTrafficPolicyInput) (*route53.DeleteTrafficPolicyOutput, error) {
    var output route53.DeleteTrafficPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTrafficPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) DeleteTrafficPolicyAsync(ctx workflow.Context, input *route53.DeleteTrafficPolicyInput) *Route53DeleteTrafficPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTrafficPolicy, input)
    return &Route53DeleteTrafficPolicyResult{Result: future}
}

func (a *Route53Stub) DeleteTrafficPolicyInstance(ctx workflow.Context, input *route53.DeleteTrafficPolicyInstanceInput) (*route53.DeleteTrafficPolicyInstanceOutput, error) {
    var output route53.DeleteTrafficPolicyInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTrafficPolicyInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) DeleteTrafficPolicyInstanceAsync(ctx workflow.Context, input *route53.DeleteTrafficPolicyInstanceInput) *Route53DeleteTrafficPolicyInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTrafficPolicyInstance, input)
    return &Route53DeleteTrafficPolicyInstanceResult{Result: future}
}

func (a *Route53Stub) DeleteVPCAssociationAuthorization(ctx workflow.Context, input *route53.DeleteVPCAssociationAuthorizationInput) (*route53.DeleteVPCAssociationAuthorizationOutput, error) {
    var output route53.DeleteVPCAssociationAuthorizationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVPCAssociationAuthorization, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) DeleteVPCAssociationAuthorizationAsync(ctx workflow.Context, input *route53.DeleteVPCAssociationAuthorizationInput) *Route53DeleteVPCAssociationAuthorizationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteVPCAssociationAuthorization, input)
    return &Route53DeleteVPCAssociationAuthorizationResult{Result: future}
}

func (a *Route53Stub) DisassociateVPCFromHostedZone(ctx workflow.Context, input *route53.DisassociateVPCFromHostedZoneInput) (*route53.DisassociateVPCFromHostedZoneOutput, error) {
    var output route53.DisassociateVPCFromHostedZoneOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateVPCFromHostedZone, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) DisassociateVPCFromHostedZoneAsync(ctx workflow.Context, input *route53.DisassociateVPCFromHostedZoneInput) *Route53DisassociateVPCFromHostedZoneResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateVPCFromHostedZone, input)
    return &Route53DisassociateVPCFromHostedZoneResult{Result: future}
}

func (a *Route53Stub) GetAccountLimit(ctx workflow.Context, input *route53.GetAccountLimitInput) (*route53.GetAccountLimitOutput, error) {
    var output route53.GetAccountLimitOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAccountLimit, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) GetAccountLimitAsync(ctx workflow.Context, input *route53.GetAccountLimitInput) *Route53GetAccountLimitResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAccountLimit, input)
    return &Route53GetAccountLimitResult{Result: future}
}

func (a *Route53Stub) GetChange(ctx workflow.Context, input *route53.GetChangeInput) (*route53.GetChangeOutput, error) {
    var output route53.GetChangeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetChange, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) GetChangeAsync(ctx workflow.Context, input *route53.GetChangeInput) *Route53GetChangeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetChange, input)
    return &Route53GetChangeResult{Result: future}
}

func (a *Route53Stub) GetCheckerIpRanges(ctx workflow.Context, input *route53.GetCheckerIpRangesInput) (*route53.GetCheckerIpRangesOutput, error) {
    var output route53.GetCheckerIpRangesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCheckerIpRanges, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) GetCheckerIpRangesAsync(ctx workflow.Context, input *route53.GetCheckerIpRangesInput) *Route53GetCheckerIpRangesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetCheckerIpRanges, input)
    return &Route53GetCheckerIpRangesResult{Result: future}
}

func (a *Route53Stub) GetGeoLocation(ctx workflow.Context, input *route53.GetGeoLocationInput) (*route53.GetGeoLocationOutput, error) {
    var output route53.GetGeoLocationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetGeoLocation, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) GetGeoLocationAsync(ctx workflow.Context, input *route53.GetGeoLocationInput) *Route53GetGeoLocationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetGeoLocation, input)
    return &Route53GetGeoLocationResult{Result: future}
}

func (a *Route53Stub) GetHealthCheck(ctx workflow.Context, input *route53.GetHealthCheckInput) (*route53.GetHealthCheckOutput, error) {
    var output route53.GetHealthCheckOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetHealthCheck, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) GetHealthCheckAsync(ctx workflow.Context, input *route53.GetHealthCheckInput) *Route53GetHealthCheckResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetHealthCheck, input)
    return &Route53GetHealthCheckResult{Result: future}
}

func (a *Route53Stub) GetHealthCheckCount(ctx workflow.Context, input *route53.GetHealthCheckCountInput) (*route53.GetHealthCheckCountOutput, error) {
    var output route53.GetHealthCheckCountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetHealthCheckCount, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) GetHealthCheckCountAsync(ctx workflow.Context, input *route53.GetHealthCheckCountInput) *Route53GetHealthCheckCountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetHealthCheckCount, input)
    return &Route53GetHealthCheckCountResult{Result: future}
}

func (a *Route53Stub) GetHealthCheckLastFailureReason(ctx workflow.Context, input *route53.GetHealthCheckLastFailureReasonInput) (*route53.GetHealthCheckLastFailureReasonOutput, error) {
    var output route53.GetHealthCheckLastFailureReasonOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetHealthCheckLastFailureReason, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) GetHealthCheckLastFailureReasonAsync(ctx workflow.Context, input *route53.GetHealthCheckLastFailureReasonInput) *Route53GetHealthCheckLastFailureReasonResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetHealthCheckLastFailureReason, input)
    return &Route53GetHealthCheckLastFailureReasonResult{Result: future}
}

func (a *Route53Stub) GetHealthCheckStatus(ctx workflow.Context, input *route53.GetHealthCheckStatusInput) (*route53.GetHealthCheckStatusOutput, error) {
    var output route53.GetHealthCheckStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetHealthCheckStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) GetHealthCheckStatusAsync(ctx workflow.Context, input *route53.GetHealthCheckStatusInput) *Route53GetHealthCheckStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetHealthCheckStatus, input)
    return &Route53GetHealthCheckStatusResult{Result: future}
}

func (a *Route53Stub) GetHostedZone(ctx workflow.Context, input *route53.GetHostedZoneInput) (*route53.GetHostedZoneOutput, error) {
    var output route53.GetHostedZoneOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetHostedZone, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) GetHostedZoneAsync(ctx workflow.Context, input *route53.GetHostedZoneInput) *Route53GetHostedZoneResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetHostedZone, input)
    return &Route53GetHostedZoneResult{Result: future}
}

func (a *Route53Stub) GetHostedZoneCount(ctx workflow.Context, input *route53.GetHostedZoneCountInput) (*route53.GetHostedZoneCountOutput, error) {
    var output route53.GetHostedZoneCountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetHostedZoneCount, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) GetHostedZoneCountAsync(ctx workflow.Context, input *route53.GetHostedZoneCountInput) *Route53GetHostedZoneCountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetHostedZoneCount, input)
    return &Route53GetHostedZoneCountResult{Result: future}
}

func (a *Route53Stub) GetHostedZoneLimit(ctx workflow.Context, input *route53.GetHostedZoneLimitInput) (*route53.GetHostedZoneLimitOutput, error) {
    var output route53.GetHostedZoneLimitOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetHostedZoneLimit, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) GetHostedZoneLimitAsync(ctx workflow.Context, input *route53.GetHostedZoneLimitInput) *Route53GetHostedZoneLimitResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetHostedZoneLimit, input)
    return &Route53GetHostedZoneLimitResult{Result: future}
}

func (a *Route53Stub) GetQueryLoggingConfig(ctx workflow.Context, input *route53.GetQueryLoggingConfigInput) (*route53.GetQueryLoggingConfigOutput, error) {
    var output route53.GetQueryLoggingConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetQueryLoggingConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) GetQueryLoggingConfigAsync(ctx workflow.Context, input *route53.GetQueryLoggingConfigInput) *Route53GetQueryLoggingConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetQueryLoggingConfig, input)
    return &Route53GetQueryLoggingConfigResult{Result: future}
}

func (a *Route53Stub) GetReusableDelegationSet(ctx workflow.Context, input *route53.GetReusableDelegationSetInput) (*route53.GetReusableDelegationSetOutput, error) {
    var output route53.GetReusableDelegationSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetReusableDelegationSet, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) GetReusableDelegationSetAsync(ctx workflow.Context, input *route53.GetReusableDelegationSetInput) *Route53GetReusableDelegationSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetReusableDelegationSet, input)
    return &Route53GetReusableDelegationSetResult{Result: future}
}

func (a *Route53Stub) GetReusableDelegationSetLimit(ctx workflow.Context, input *route53.GetReusableDelegationSetLimitInput) (*route53.GetReusableDelegationSetLimitOutput, error) {
    var output route53.GetReusableDelegationSetLimitOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetReusableDelegationSetLimit, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) GetReusableDelegationSetLimitAsync(ctx workflow.Context, input *route53.GetReusableDelegationSetLimitInput) *Route53GetReusableDelegationSetLimitResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetReusableDelegationSetLimit, input)
    return &Route53GetReusableDelegationSetLimitResult{Result: future}
}

func (a *Route53Stub) GetTrafficPolicy(ctx workflow.Context, input *route53.GetTrafficPolicyInput) (*route53.GetTrafficPolicyOutput, error) {
    var output route53.GetTrafficPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTrafficPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) GetTrafficPolicyAsync(ctx workflow.Context, input *route53.GetTrafficPolicyInput) *Route53GetTrafficPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTrafficPolicy, input)
    return &Route53GetTrafficPolicyResult{Result: future}
}

func (a *Route53Stub) GetTrafficPolicyInstance(ctx workflow.Context, input *route53.GetTrafficPolicyInstanceInput) (*route53.GetTrafficPolicyInstanceOutput, error) {
    var output route53.GetTrafficPolicyInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTrafficPolicyInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) GetTrafficPolicyInstanceAsync(ctx workflow.Context, input *route53.GetTrafficPolicyInstanceInput) *Route53GetTrafficPolicyInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTrafficPolicyInstance, input)
    return &Route53GetTrafficPolicyInstanceResult{Result: future}
}

func (a *Route53Stub) GetTrafficPolicyInstanceCount(ctx workflow.Context, input *route53.GetTrafficPolicyInstanceCountInput) (*route53.GetTrafficPolicyInstanceCountOutput, error) {
    var output route53.GetTrafficPolicyInstanceCountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTrafficPolicyInstanceCount, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) GetTrafficPolicyInstanceCountAsync(ctx workflow.Context, input *route53.GetTrafficPolicyInstanceCountInput) *Route53GetTrafficPolicyInstanceCountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTrafficPolicyInstanceCount, input)
    return &Route53GetTrafficPolicyInstanceCountResult{Result: future}
}

func (a *Route53Stub) ListGeoLocations(ctx workflow.Context, input *route53.ListGeoLocationsInput) (*route53.ListGeoLocationsOutput, error) {
    var output route53.ListGeoLocationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListGeoLocations, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) ListGeoLocationsAsync(ctx workflow.Context, input *route53.ListGeoLocationsInput) *Route53ListGeoLocationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListGeoLocations, input)
    return &Route53ListGeoLocationsResult{Result: future}
}

func (a *Route53Stub) ListHealthChecks(ctx workflow.Context, input *route53.ListHealthChecksInput) (*route53.ListHealthChecksOutput, error) {
    var output route53.ListHealthChecksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListHealthChecks, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) ListHealthChecksAsync(ctx workflow.Context, input *route53.ListHealthChecksInput) *Route53ListHealthChecksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListHealthChecks, input)
    return &Route53ListHealthChecksResult{Result: future}
}

func (a *Route53Stub) ListHostedZones(ctx workflow.Context, input *route53.ListHostedZonesInput) (*route53.ListHostedZonesOutput, error) {
    var output route53.ListHostedZonesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListHostedZones, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) ListHostedZonesAsync(ctx workflow.Context, input *route53.ListHostedZonesInput) *Route53ListHostedZonesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListHostedZones, input)
    return &Route53ListHostedZonesResult{Result: future}
}

func (a *Route53Stub) ListHostedZonesByName(ctx workflow.Context, input *route53.ListHostedZonesByNameInput) (*route53.ListHostedZonesByNameOutput, error) {
    var output route53.ListHostedZonesByNameOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListHostedZonesByName, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) ListHostedZonesByNameAsync(ctx workflow.Context, input *route53.ListHostedZonesByNameInput) *Route53ListHostedZonesByNameResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListHostedZonesByName, input)
    return &Route53ListHostedZonesByNameResult{Result: future}
}

func (a *Route53Stub) ListHostedZonesByVPC(ctx workflow.Context, input *route53.ListHostedZonesByVPCInput) (*route53.ListHostedZonesByVPCOutput, error) {
    var output route53.ListHostedZonesByVPCOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListHostedZonesByVPC, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) ListHostedZonesByVPCAsync(ctx workflow.Context, input *route53.ListHostedZonesByVPCInput) *Route53ListHostedZonesByVPCResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListHostedZonesByVPC, input)
    return &Route53ListHostedZonesByVPCResult{Result: future}
}

func (a *Route53Stub) ListQueryLoggingConfigs(ctx workflow.Context, input *route53.ListQueryLoggingConfigsInput) (*route53.ListQueryLoggingConfigsOutput, error) {
    var output route53.ListQueryLoggingConfigsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListQueryLoggingConfigs, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) ListQueryLoggingConfigsAsync(ctx workflow.Context, input *route53.ListQueryLoggingConfigsInput) *Route53ListQueryLoggingConfigsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListQueryLoggingConfigs, input)
    return &Route53ListQueryLoggingConfigsResult{Result: future}
}

func (a *Route53Stub) ListResourceRecordSets(ctx workflow.Context, input *route53.ListResourceRecordSetsInput) (*route53.ListResourceRecordSetsOutput, error) {
    var output route53.ListResourceRecordSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListResourceRecordSets, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) ListResourceRecordSetsAsync(ctx workflow.Context, input *route53.ListResourceRecordSetsInput) *Route53ListResourceRecordSetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListResourceRecordSets, input)
    return &Route53ListResourceRecordSetsResult{Result: future}
}

func (a *Route53Stub) ListReusableDelegationSets(ctx workflow.Context, input *route53.ListReusableDelegationSetsInput) (*route53.ListReusableDelegationSetsOutput, error) {
    var output route53.ListReusableDelegationSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListReusableDelegationSets, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) ListReusableDelegationSetsAsync(ctx workflow.Context, input *route53.ListReusableDelegationSetsInput) *Route53ListReusableDelegationSetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListReusableDelegationSets, input)
    return &Route53ListReusableDelegationSetsResult{Result: future}
}

func (a *Route53Stub) ListTagsForResource(ctx workflow.Context, input *route53.ListTagsForResourceInput) (*route53.ListTagsForResourceOutput, error) {
    var output route53.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) ListTagsForResourceAsync(ctx workflow.Context, input *route53.ListTagsForResourceInput) *Route53ListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &Route53ListTagsForResourceResult{Result: future}
}

func (a *Route53Stub) ListTagsForResources(ctx workflow.Context, input *route53.ListTagsForResourcesInput) (*route53.ListTagsForResourcesOutput, error) {
    var output route53.ListTagsForResourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResources, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) ListTagsForResourcesAsync(ctx workflow.Context, input *route53.ListTagsForResourcesInput) *Route53ListTagsForResourcesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResources, input)
    return &Route53ListTagsForResourcesResult{Result: future}
}

func (a *Route53Stub) ListTrafficPolicies(ctx workflow.Context, input *route53.ListTrafficPoliciesInput) (*route53.ListTrafficPoliciesOutput, error) {
    var output route53.ListTrafficPoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTrafficPolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) ListTrafficPoliciesAsync(ctx workflow.Context, input *route53.ListTrafficPoliciesInput) *Route53ListTrafficPoliciesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTrafficPolicies, input)
    return &Route53ListTrafficPoliciesResult{Result: future}
}

func (a *Route53Stub) ListTrafficPolicyInstances(ctx workflow.Context, input *route53.ListTrafficPolicyInstancesInput) (*route53.ListTrafficPolicyInstancesOutput, error) {
    var output route53.ListTrafficPolicyInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTrafficPolicyInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) ListTrafficPolicyInstancesAsync(ctx workflow.Context, input *route53.ListTrafficPolicyInstancesInput) *Route53ListTrafficPolicyInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTrafficPolicyInstances, input)
    return &Route53ListTrafficPolicyInstancesResult{Result: future}
}

func (a *Route53Stub) ListTrafficPolicyInstancesByHostedZone(ctx workflow.Context, input *route53.ListTrafficPolicyInstancesByHostedZoneInput) (*route53.ListTrafficPolicyInstancesByHostedZoneOutput, error) {
    var output route53.ListTrafficPolicyInstancesByHostedZoneOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTrafficPolicyInstancesByHostedZone, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) ListTrafficPolicyInstancesByHostedZoneAsync(ctx workflow.Context, input *route53.ListTrafficPolicyInstancesByHostedZoneInput) *Route53ListTrafficPolicyInstancesByHostedZoneResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTrafficPolicyInstancesByHostedZone, input)
    return &Route53ListTrafficPolicyInstancesByHostedZoneResult{Result: future}
}

func (a *Route53Stub) ListTrafficPolicyInstancesByPolicy(ctx workflow.Context, input *route53.ListTrafficPolicyInstancesByPolicyInput) (*route53.ListTrafficPolicyInstancesByPolicyOutput, error) {
    var output route53.ListTrafficPolicyInstancesByPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTrafficPolicyInstancesByPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) ListTrafficPolicyInstancesByPolicyAsync(ctx workflow.Context, input *route53.ListTrafficPolicyInstancesByPolicyInput) *Route53ListTrafficPolicyInstancesByPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTrafficPolicyInstancesByPolicy, input)
    return &Route53ListTrafficPolicyInstancesByPolicyResult{Result: future}
}

func (a *Route53Stub) ListTrafficPolicyVersions(ctx workflow.Context, input *route53.ListTrafficPolicyVersionsInput) (*route53.ListTrafficPolicyVersionsOutput, error) {
    var output route53.ListTrafficPolicyVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTrafficPolicyVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) ListTrafficPolicyVersionsAsync(ctx workflow.Context, input *route53.ListTrafficPolicyVersionsInput) *Route53ListTrafficPolicyVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTrafficPolicyVersions, input)
    return &Route53ListTrafficPolicyVersionsResult{Result: future}
}

func (a *Route53Stub) ListVPCAssociationAuthorizations(ctx workflow.Context, input *route53.ListVPCAssociationAuthorizationsInput) (*route53.ListVPCAssociationAuthorizationsOutput, error) {
    var output route53.ListVPCAssociationAuthorizationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListVPCAssociationAuthorizations, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) ListVPCAssociationAuthorizationsAsync(ctx workflow.Context, input *route53.ListVPCAssociationAuthorizationsInput) *Route53ListVPCAssociationAuthorizationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListVPCAssociationAuthorizations, input)
    return &Route53ListVPCAssociationAuthorizationsResult{Result: future}
}

func (a *Route53Stub) TestDNSAnswer(ctx workflow.Context, input *route53.TestDNSAnswerInput) (*route53.TestDNSAnswerOutput, error) {
    var output route53.TestDNSAnswerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TestDNSAnswer, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) TestDNSAnswerAsync(ctx workflow.Context, input *route53.TestDNSAnswerInput) *Route53TestDNSAnswerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TestDNSAnswer, input)
    return &Route53TestDNSAnswerResult{Result: future}
}

func (a *Route53Stub) UpdateHealthCheck(ctx workflow.Context, input *route53.UpdateHealthCheckInput) (*route53.UpdateHealthCheckOutput, error) {
    var output route53.UpdateHealthCheckOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateHealthCheck, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) UpdateHealthCheckAsync(ctx workflow.Context, input *route53.UpdateHealthCheckInput) *Route53UpdateHealthCheckResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateHealthCheck, input)
    return &Route53UpdateHealthCheckResult{Result: future}
}

func (a *Route53Stub) UpdateHostedZoneComment(ctx workflow.Context, input *route53.UpdateHostedZoneCommentInput) (*route53.UpdateHostedZoneCommentOutput, error) {
    var output route53.UpdateHostedZoneCommentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateHostedZoneComment, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) UpdateHostedZoneCommentAsync(ctx workflow.Context, input *route53.UpdateHostedZoneCommentInput) *Route53UpdateHostedZoneCommentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateHostedZoneComment, input)
    return &Route53UpdateHostedZoneCommentResult{Result: future}
}

func (a *Route53Stub) UpdateTrafficPolicyComment(ctx workflow.Context, input *route53.UpdateTrafficPolicyCommentInput) (*route53.UpdateTrafficPolicyCommentOutput, error) {
    var output route53.UpdateTrafficPolicyCommentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTrafficPolicyComment, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) UpdateTrafficPolicyCommentAsync(ctx workflow.Context, input *route53.UpdateTrafficPolicyCommentInput) *Route53UpdateTrafficPolicyCommentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateTrafficPolicyComment, input)
    return &Route53UpdateTrafficPolicyCommentResult{Result: future}
}

func (a *Route53Stub) UpdateTrafficPolicyInstance(ctx workflow.Context, input *route53.UpdateTrafficPolicyInstanceInput) (*route53.UpdateTrafficPolicyInstanceOutput, error) {
    var output route53.UpdateTrafficPolicyInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTrafficPolicyInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53Stub) UpdateTrafficPolicyInstanceAsync(ctx workflow.Context, input *route53.UpdateTrafficPolicyInstanceInput) *Route53UpdateTrafficPolicyInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateTrafficPolicyInstance, input)
    return &Route53UpdateTrafficPolicyInstanceResult{Result: future}
}

func (a *Route53Stub) WaitUntilResourceRecordSetsChanged(ctx workflow.Context, input *route53.GetChangeInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilResourceRecordSetsChanged, input).Get(ctx, nil)
}

func (a *Route53Stub) WaitUntilResourceRecordSetsChangedAsync(ctx workflow.Context, input *route53.GetChangeInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilResourceRecordSetsChanged, input)
}

