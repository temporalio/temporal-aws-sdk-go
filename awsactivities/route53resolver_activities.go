
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/route53resolver"
	"github.com/aws/aws-sdk-go/service/route53resolver/route53resolveriface"
)

type Route53ResolverActivities struct {
	client route53resolveriface.Route53ResolverAPI
}

func NewRoute53ResolverActivities(client route53resolveriface.Route53ResolverAPI) *Route53ResolverActivities {
    return &Route53ResolverActivities{client: client}
}

func (a *Route53ResolverActivities) AssociateResolverEndpointIpAddress(input *route53resolver.AssociateResolverEndpointIpAddressInput) (*route53resolver.AssociateResolverEndpointIpAddressOutput, error) {
    return a.client.AssociateResolverEndpointIpAddress(input)
}

func (a *Route53ResolverActivities) AssociateResolverQueryLogConfig(input *route53resolver.AssociateResolverQueryLogConfigInput) (*route53resolver.AssociateResolverQueryLogConfigOutput, error) {
    return a.client.AssociateResolverQueryLogConfig(input)
}

func (a *Route53ResolverActivities) AssociateResolverRule(input *route53resolver.AssociateResolverRuleInput) (*route53resolver.AssociateResolverRuleOutput, error) {
    return a.client.AssociateResolverRule(input)
}

func (a *Route53ResolverActivities) CreateResolverEndpoint(input *route53resolver.CreateResolverEndpointInput) (*route53resolver.CreateResolverEndpointOutput, error) {
    return a.client.CreateResolverEndpoint(input)
}

func (a *Route53ResolverActivities) CreateResolverQueryLogConfig(input *route53resolver.CreateResolverQueryLogConfigInput) (*route53resolver.CreateResolverQueryLogConfigOutput, error) {
    return a.client.CreateResolverQueryLogConfig(input)
}

func (a *Route53ResolverActivities) CreateResolverRule(input *route53resolver.CreateResolverRuleInput) (*route53resolver.CreateResolverRuleOutput, error) {
    return a.client.CreateResolverRule(input)
}

func (a *Route53ResolverActivities) DeleteResolverEndpoint(input *route53resolver.DeleteResolverEndpointInput) (*route53resolver.DeleteResolverEndpointOutput, error) {
    return a.client.DeleteResolverEndpoint(input)
}

func (a *Route53ResolverActivities) DeleteResolverQueryLogConfig(input *route53resolver.DeleteResolverQueryLogConfigInput) (*route53resolver.DeleteResolverQueryLogConfigOutput, error) {
    return a.client.DeleteResolverQueryLogConfig(input)
}

func (a *Route53ResolverActivities) DeleteResolverRule(input *route53resolver.DeleteResolverRuleInput) (*route53resolver.DeleteResolverRuleOutput, error) {
    return a.client.DeleteResolverRule(input)
}

func (a *Route53ResolverActivities) DisassociateResolverEndpointIpAddress(input *route53resolver.DisassociateResolverEndpointIpAddressInput) (*route53resolver.DisassociateResolverEndpointIpAddressOutput, error) {
    return a.client.DisassociateResolverEndpointIpAddress(input)
}

func (a *Route53ResolverActivities) DisassociateResolverQueryLogConfig(input *route53resolver.DisassociateResolverQueryLogConfigInput) (*route53resolver.DisassociateResolverQueryLogConfigOutput, error) {
    return a.client.DisassociateResolverQueryLogConfig(input)
}

func (a *Route53ResolverActivities) DisassociateResolverRule(input *route53resolver.DisassociateResolverRuleInput) (*route53resolver.DisassociateResolverRuleOutput, error) {
    return a.client.DisassociateResolverRule(input)
}

func (a *Route53ResolverActivities) GetResolverEndpoint(input *route53resolver.GetResolverEndpointInput) (*route53resolver.GetResolverEndpointOutput, error) {
    return a.client.GetResolverEndpoint(input)
}

func (a *Route53ResolverActivities) GetResolverQueryLogConfig(input *route53resolver.GetResolverQueryLogConfigInput) (*route53resolver.GetResolverQueryLogConfigOutput, error) {
    return a.client.GetResolverQueryLogConfig(input)
}

func (a *Route53ResolverActivities) GetResolverQueryLogConfigAssociation(input *route53resolver.GetResolverQueryLogConfigAssociationInput) (*route53resolver.GetResolverQueryLogConfigAssociationOutput, error) {
    return a.client.GetResolverQueryLogConfigAssociation(input)
}

func (a *Route53ResolverActivities) GetResolverQueryLogConfigPolicy(input *route53resolver.GetResolverQueryLogConfigPolicyInput) (*route53resolver.GetResolverQueryLogConfigPolicyOutput, error) {
    return a.client.GetResolverQueryLogConfigPolicy(input)
}

func (a *Route53ResolverActivities) GetResolverRule(input *route53resolver.GetResolverRuleInput) (*route53resolver.GetResolverRuleOutput, error) {
    return a.client.GetResolverRule(input)
}

func (a *Route53ResolverActivities) GetResolverRuleAssociation(input *route53resolver.GetResolverRuleAssociationInput) (*route53resolver.GetResolverRuleAssociationOutput, error) {
    return a.client.GetResolverRuleAssociation(input)
}

func (a *Route53ResolverActivities) GetResolverRulePolicy(input *route53resolver.GetResolverRulePolicyInput) (*route53resolver.GetResolverRulePolicyOutput, error) {
    return a.client.GetResolverRulePolicy(input)
}

func (a *Route53ResolverActivities) ListResolverEndpointIpAddresses(input *route53resolver.ListResolverEndpointIpAddressesInput) (*route53resolver.ListResolverEndpointIpAddressesOutput, error) {
    return a.client.ListResolverEndpointIpAddresses(input)
}

func (a *Route53ResolverActivities) ListResolverEndpoints(input *route53resolver.ListResolverEndpointsInput) (*route53resolver.ListResolverEndpointsOutput, error) {
    return a.client.ListResolverEndpoints(input)
}

func (a *Route53ResolverActivities) ListResolverQueryLogConfigAssociations(input *route53resolver.ListResolverQueryLogConfigAssociationsInput) (*route53resolver.ListResolverQueryLogConfigAssociationsOutput, error) {
    return a.client.ListResolverQueryLogConfigAssociations(input)
}

func (a *Route53ResolverActivities) ListResolverQueryLogConfigs(input *route53resolver.ListResolverQueryLogConfigsInput) (*route53resolver.ListResolverQueryLogConfigsOutput, error) {
    return a.client.ListResolverQueryLogConfigs(input)
}

func (a *Route53ResolverActivities) ListResolverRuleAssociations(input *route53resolver.ListResolverRuleAssociationsInput) (*route53resolver.ListResolverRuleAssociationsOutput, error) {
    return a.client.ListResolverRuleAssociations(input)
}

func (a *Route53ResolverActivities) ListResolverRules(input *route53resolver.ListResolverRulesInput) (*route53resolver.ListResolverRulesOutput, error) {
    return a.client.ListResolverRules(input)
}

func (a *Route53ResolverActivities) ListTagsForResource(input *route53resolver.ListTagsForResourceInput) (*route53resolver.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *Route53ResolverActivities) PutResolverQueryLogConfigPolicy(input *route53resolver.PutResolverQueryLogConfigPolicyInput) (*route53resolver.PutResolverQueryLogConfigPolicyOutput, error) {
    return a.client.PutResolverQueryLogConfigPolicy(input)
}

func (a *Route53ResolverActivities) PutResolverRulePolicy(input *route53resolver.PutResolverRulePolicyInput) (*route53resolver.PutResolverRulePolicyOutput, error) {
    return a.client.PutResolverRulePolicy(input)
}

func (a *Route53ResolverActivities) TagResource(input *route53resolver.TagResourceInput) (*route53resolver.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *Route53ResolverActivities) UntagResource(input *route53resolver.UntagResourceInput) (*route53resolver.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *Route53ResolverActivities) UpdateResolverEndpoint(input *route53resolver.UpdateResolverEndpointInput) (*route53resolver.UpdateResolverEndpointOutput, error) {
    return a.client.UpdateResolverEndpoint(input)
}

func (a *Route53ResolverActivities) UpdateResolverRule(input *route53resolver.UpdateResolverRuleInput) (*route53resolver.UpdateResolverRuleOutput, error) {
    return a.client.UpdateResolverRule(input)
}
