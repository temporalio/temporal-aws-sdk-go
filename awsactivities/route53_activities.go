package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/route53/route53iface"
)

type Route53Activities struct {
	client route53iface.Route53API
}

func NewRoute53Activities(client route53iface.Route53API) *Route53Activities {
    return &Route53Activities{client: client}
}

func (a *Route53Activities) AssociateVPCWithHostedZone(input *route53.AssociateVPCWithHostedZoneInput) (*route53.AssociateVPCWithHostedZoneOutput, error) {
    return a.client.AssociateVPCWithHostedZone(input)
}

func (a *Route53Activities) ChangeResourceRecordSets(input *route53.ChangeResourceRecordSetsInput) (*route53.ChangeResourceRecordSetsOutput, error) {
    return a.client.ChangeResourceRecordSets(input)
}

func (a *Route53Activities) ChangeTagsForResource(input *route53.ChangeTagsForResourceInput) (*route53.ChangeTagsForResourceOutput, error) {
    return a.client.ChangeTagsForResource(input)
}

func (a *Route53Activities) CreateHealthCheck(input *route53.CreateHealthCheckInput) (*route53.CreateHealthCheckOutput, error) {
    return a.client.CreateHealthCheck(input)
}

func (a *Route53Activities) CreateHostedZone(input *route53.CreateHostedZoneInput) (*route53.CreateHostedZoneOutput, error) {
    return a.client.CreateHostedZone(input)
}

func (a *Route53Activities) CreateQueryLoggingConfig(input *route53.CreateQueryLoggingConfigInput) (*route53.CreateQueryLoggingConfigOutput, error) {
    return a.client.CreateQueryLoggingConfig(input)
}

func (a *Route53Activities) CreateReusableDelegationSet(input *route53.CreateReusableDelegationSetInput) (*route53.CreateReusableDelegationSetOutput, error) {
    return a.client.CreateReusableDelegationSet(input)
}

func (a *Route53Activities) CreateTrafficPolicy(input *route53.CreateTrafficPolicyInput) (*route53.CreateTrafficPolicyOutput, error) {
    return a.client.CreateTrafficPolicy(input)
}

func (a *Route53Activities) CreateTrafficPolicyInstance(input *route53.CreateTrafficPolicyInstanceInput) (*route53.CreateTrafficPolicyInstanceOutput, error) {
    return a.client.CreateTrafficPolicyInstance(input)
}

func (a *Route53Activities) CreateTrafficPolicyVersion(input *route53.CreateTrafficPolicyVersionInput) (*route53.CreateTrafficPolicyVersionOutput, error) {
    return a.client.CreateTrafficPolicyVersion(input)
}

func (a *Route53Activities) CreateVPCAssociationAuthorization(input *route53.CreateVPCAssociationAuthorizationInput) (*route53.CreateVPCAssociationAuthorizationOutput, error) {
    return a.client.CreateVPCAssociationAuthorization(input)
}

func (a *Route53Activities) DeleteHealthCheck(input *route53.DeleteHealthCheckInput) (*route53.DeleteHealthCheckOutput, error) {
    return a.client.DeleteHealthCheck(input)
}

func (a *Route53Activities) DeleteHostedZone(input *route53.DeleteHostedZoneInput) (*route53.DeleteHostedZoneOutput, error) {
    return a.client.DeleteHostedZone(input)
}

func (a *Route53Activities) DeleteQueryLoggingConfig(input *route53.DeleteQueryLoggingConfigInput) (*route53.DeleteQueryLoggingConfigOutput, error) {
    return a.client.DeleteQueryLoggingConfig(input)
}

func (a *Route53Activities) DeleteReusableDelegationSet(input *route53.DeleteReusableDelegationSetInput) (*route53.DeleteReusableDelegationSetOutput, error) {
    return a.client.DeleteReusableDelegationSet(input)
}

func (a *Route53Activities) DeleteTrafficPolicy(input *route53.DeleteTrafficPolicyInput) (*route53.DeleteTrafficPolicyOutput, error) {
    return a.client.DeleteTrafficPolicy(input)
}

func (a *Route53Activities) DeleteTrafficPolicyInstance(input *route53.DeleteTrafficPolicyInstanceInput) (*route53.DeleteTrafficPolicyInstanceOutput, error) {
    return a.client.DeleteTrafficPolicyInstance(input)
}

func (a *Route53Activities) DeleteVPCAssociationAuthorization(input *route53.DeleteVPCAssociationAuthorizationInput) (*route53.DeleteVPCAssociationAuthorizationOutput, error) {
    return a.client.DeleteVPCAssociationAuthorization(input)
}

func (a *Route53Activities) DisassociateVPCFromHostedZone(input *route53.DisassociateVPCFromHostedZoneInput) (*route53.DisassociateVPCFromHostedZoneOutput, error) {
    return a.client.DisassociateVPCFromHostedZone(input)
}

func (a *Route53Activities) GetAccountLimit(input *route53.GetAccountLimitInput) (*route53.GetAccountLimitOutput, error) {
    return a.client.GetAccountLimit(input)
}

func (a *Route53Activities) GetChange(input *route53.GetChangeInput) (*route53.GetChangeOutput, error) {
    return a.client.GetChange(input)
}

func (a *Route53Activities) GetCheckerIpRanges(input *route53.GetCheckerIpRangesInput) (*route53.GetCheckerIpRangesOutput, error) {
    return a.client.GetCheckerIpRanges(input)
}

func (a *Route53Activities) GetGeoLocation(input *route53.GetGeoLocationInput) (*route53.GetGeoLocationOutput, error) {
    return a.client.GetGeoLocation(input)
}

func (a *Route53Activities) GetHealthCheck(input *route53.GetHealthCheckInput) (*route53.GetHealthCheckOutput, error) {
    return a.client.GetHealthCheck(input)
}

func (a *Route53Activities) GetHealthCheckCount(input *route53.GetHealthCheckCountInput) (*route53.GetHealthCheckCountOutput, error) {
    return a.client.GetHealthCheckCount(input)
}

func (a *Route53Activities) GetHealthCheckLastFailureReason(input *route53.GetHealthCheckLastFailureReasonInput) (*route53.GetHealthCheckLastFailureReasonOutput, error) {
    return a.client.GetHealthCheckLastFailureReason(input)
}

func (a *Route53Activities) GetHealthCheckStatus(input *route53.GetHealthCheckStatusInput) (*route53.GetHealthCheckStatusOutput, error) {
    return a.client.GetHealthCheckStatus(input)
}

func (a *Route53Activities) GetHostedZone(input *route53.GetHostedZoneInput) (*route53.GetHostedZoneOutput, error) {
    return a.client.GetHostedZone(input)
}

func (a *Route53Activities) GetHostedZoneCount(input *route53.GetHostedZoneCountInput) (*route53.GetHostedZoneCountOutput, error) {
    return a.client.GetHostedZoneCount(input)
}

func (a *Route53Activities) GetHostedZoneLimit(input *route53.GetHostedZoneLimitInput) (*route53.GetHostedZoneLimitOutput, error) {
    return a.client.GetHostedZoneLimit(input)
}

func (a *Route53Activities) GetQueryLoggingConfig(input *route53.GetQueryLoggingConfigInput) (*route53.GetQueryLoggingConfigOutput, error) {
    return a.client.GetQueryLoggingConfig(input)
}

func (a *Route53Activities) GetReusableDelegationSet(input *route53.GetReusableDelegationSetInput) (*route53.GetReusableDelegationSetOutput, error) {
    return a.client.GetReusableDelegationSet(input)
}

func (a *Route53Activities) GetReusableDelegationSetLimit(input *route53.GetReusableDelegationSetLimitInput) (*route53.GetReusableDelegationSetLimitOutput, error) {
    return a.client.GetReusableDelegationSetLimit(input)
}

func (a *Route53Activities) GetTrafficPolicy(input *route53.GetTrafficPolicyInput) (*route53.GetTrafficPolicyOutput, error) {
    return a.client.GetTrafficPolicy(input)
}

func (a *Route53Activities) GetTrafficPolicyInstance(input *route53.GetTrafficPolicyInstanceInput) (*route53.GetTrafficPolicyInstanceOutput, error) {
    return a.client.GetTrafficPolicyInstance(input)
}

func (a *Route53Activities) GetTrafficPolicyInstanceCount(input *route53.GetTrafficPolicyInstanceCountInput) (*route53.GetTrafficPolicyInstanceCountOutput, error) {
    return a.client.GetTrafficPolicyInstanceCount(input)
}

func (a *Route53Activities) ListGeoLocations(input *route53.ListGeoLocationsInput) (*route53.ListGeoLocationsOutput, error) {
    return a.client.ListGeoLocations(input)
}

func (a *Route53Activities) ListHealthChecks(input *route53.ListHealthChecksInput) (*route53.ListHealthChecksOutput, error) {
    return a.client.ListHealthChecks(input)
}

func (a *Route53Activities) ListHostedZones(input *route53.ListHostedZonesInput) (*route53.ListHostedZonesOutput, error) {
    return a.client.ListHostedZones(input)
}

func (a *Route53Activities) ListHostedZonesByName(input *route53.ListHostedZonesByNameInput) (*route53.ListHostedZonesByNameOutput, error) {
    return a.client.ListHostedZonesByName(input)
}

func (a *Route53Activities) ListHostedZonesByVPC(input *route53.ListHostedZonesByVPCInput) (*route53.ListHostedZonesByVPCOutput, error) {
    return a.client.ListHostedZonesByVPC(input)
}

func (a *Route53Activities) ListQueryLoggingConfigs(input *route53.ListQueryLoggingConfigsInput) (*route53.ListQueryLoggingConfigsOutput, error) {
    return a.client.ListQueryLoggingConfigs(input)
}

func (a *Route53Activities) ListResourceRecordSets(input *route53.ListResourceRecordSetsInput) (*route53.ListResourceRecordSetsOutput, error) {
    return a.client.ListResourceRecordSets(input)
}

func (a *Route53Activities) ListReusableDelegationSets(input *route53.ListReusableDelegationSetsInput) (*route53.ListReusableDelegationSetsOutput, error) {
    return a.client.ListReusableDelegationSets(input)
}

func (a *Route53Activities) ListTagsForResource(input *route53.ListTagsForResourceInput) (*route53.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *Route53Activities) ListTagsForResources(input *route53.ListTagsForResourcesInput) (*route53.ListTagsForResourcesOutput, error) {
    return a.client.ListTagsForResources(input)
}

func (a *Route53Activities) ListTrafficPolicies(input *route53.ListTrafficPoliciesInput) (*route53.ListTrafficPoliciesOutput, error) {
    return a.client.ListTrafficPolicies(input)
}

func (a *Route53Activities) ListTrafficPolicyInstances(input *route53.ListTrafficPolicyInstancesInput) (*route53.ListTrafficPolicyInstancesOutput, error) {
    return a.client.ListTrafficPolicyInstances(input)
}

func (a *Route53Activities) ListTrafficPolicyInstancesByHostedZone(input *route53.ListTrafficPolicyInstancesByHostedZoneInput) (*route53.ListTrafficPolicyInstancesByHostedZoneOutput, error) {
    return a.client.ListTrafficPolicyInstancesByHostedZone(input)
}

func (a *Route53Activities) ListTrafficPolicyInstancesByPolicy(input *route53.ListTrafficPolicyInstancesByPolicyInput) (*route53.ListTrafficPolicyInstancesByPolicyOutput, error) {
    return a.client.ListTrafficPolicyInstancesByPolicy(input)
}

func (a *Route53Activities) ListTrafficPolicyVersions(input *route53.ListTrafficPolicyVersionsInput) (*route53.ListTrafficPolicyVersionsOutput, error) {
    return a.client.ListTrafficPolicyVersions(input)
}

func (a *Route53Activities) ListVPCAssociationAuthorizations(input *route53.ListVPCAssociationAuthorizationsInput) (*route53.ListVPCAssociationAuthorizationsOutput, error) {
    return a.client.ListVPCAssociationAuthorizations(input)
}

func (a *Route53Activities) TestDNSAnswer(input *route53.TestDNSAnswerInput) (*route53.TestDNSAnswerOutput, error) {
    return a.client.TestDNSAnswer(input)
}

func (a *Route53Activities) UpdateHealthCheck(input *route53.UpdateHealthCheckInput) (*route53.UpdateHealthCheckOutput, error) {
    return a.client.UpdateHealthCheck(input)
}

func (a *Route53Activities) UpdateHostedZoneComment(input *route53.UpdateHostedZoneCommentInput) (*route53.UpdateHostedZoneCommentOutput, error) {
    return a.client.UpdateHostedZoneComment(input)
}

func (a *Route53Activities) UpdateTrafficPolicyComment(input *route53.UpdateTrafficPolicyCommentInput) (*route53.UpdateTrafficPolicyCommentOutput, error) {
    return a.client.UpdateTrafficPolicyComment(input)
}

func (a *Route53Activities) UpdateTrafficPolicyInstance(input *route53.UpdateTrafficPolicyInstanceInput) (*route53.UpdateTrafficPolicyInstanceOutput, error) {
    return a.client.UpdateTrafficPolicyInstance(input)
}

func (a *Route53Activities) WaitUntilResourceRecordSetsChanged(input *route53.WaitUntilResourceRecordSetsChangedInput) (*route53.WaitUntilResourceRecordSetsChangedOutput, error) {
    return a.client.WaitUntilResourceRecordSetsChanged(input)
}
