package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53resolver"
	"github.com/aws/aws-sdk-go/service/route53resolver/route53resolveriface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type Route53ResolverActivities struct {
	client route53resolveriface.Route53ResolverAPI
}

func NewRoute53ResolverActivities(session *session.Session, config ...*aws.Config) *Route53ResolverActivities {
	client := route53resolver.New(session, config...)
	return &Route53ResolverActivities{client: client}
}

func (a *Route53ResolverActivities) AssociateResolverEndpointIpAddress(ctx context.Context, input *route53resolver.AssociateResolverEndpointIpAddressInput) (*route53resolver.AssociateResolverEndpointIpAddressOutput, error) {
	return a.client.AssociateResolverEndpointIpAddressWithContext(ctx, input)
}

func (a *Route53ResolverActivities) AssociateResolverQueryLogConfig(ctx context.Context, input *route53resolver.AssociateResolverQueryLogConfigInput) (*route53resolver.AssociateResolverQueryLogConfigOutput, error) {
	return a.client.AssociateResolverQueryLogConfigWithContext(ctx, input)
}

func (a *Route53ResolverActivities) AssociateResolverRule(ctx context.Context, input *route53resolver.AssociateResolverRuleInput) (*route53resolver.AssociateResolverRuleOutput, error) {
	return a.client.AssociateResolverRuleWithContext(ctx, input)
}

func (a *Route53ResolverActivities) CreateResolverEndpoint(ctx context.Context, input *route53resolver.CreateResolverEndpointInput) (*route53resolver.CreateResolverEndpointOutput, error) {
	return a.client.CreateResolverEndpointWithContext(ctx, input)
}

func (a *Route53ResolverActivities) CreateResolverQueryLogConfig(ctx context.Context, input *route53resolver.CreateResolverQueryLogConfigInput) (*route53resolver.CreateResolverQueryLogConfigOutput, error) {
	return a.client.CreateResolverQueryLogConfigWithContext(ctx, input)
}

func (a *Route53ResolverActivities) CreateResolverRule(ctx context.Context, input *route53resolver.CreateResolverRuleInput) (*route53resolver.CreateResolverRuleOutput, error) {
	return a.client.CreateResolverRuleWithContext(ctx, input)
}

func (a *Route53ResolverActivities) DeleteResolverEndpoint(ctx context.Context, input *route53resolver.DeleteResolverEndpointInput) (*route53resolver.DeleteResolverEndpointOutput, error) {
	return a.client.DeleteResolverEndpointWithContext(ctx, input)
}

func (a *Route53ResolverActivities) DeleteResolverQueryLogConfig(ctx context.Context, input *route53resolver.DeleteResolverQueryLogConfigInput) (*route53resolver.DeleteResolverQueryLogConfigOutput, error) {
	return a.client.DeleteResolverQueryLogConfigWithContext(ctx, input)
}

func (a *Route53ResolverActivities) DeleteResolverRule(ctx context.Context, input *route53resolver.DeleteResolverRuleInput) (*route53resolver.DeleteResolverRuleOutput, error) {
	return a.client.DeleteResolverRuleWithContext(ctx, input)
}

func (a *Route53ResolverActivities) DisassociateResolverEndpointIpAddress(ctx context.Context, input *route53resolver.DisassociateResolverEndpointIpAddressInput) (*route53resolver.DisassociateResolverEndpointIpAddressOutput, error) {
	return a.client.DisassociateResolverEndpointIpAddressWithContext(ctx, input)
}

func (a *Route53ResolverActivities) DisassociateResolverQueryLogConfig(ctx context.Context, input *route53resolver.DisassociateResolverQueryLogConfigInput) (*route53resolver.DisassociateResolverQueryLogConfigOutput, error) {
	return a.client.DisassociateResolverQueryLogConfigWithContext(ctx, input)
}

func (a *Route53ResolverActivities) DisassociateResolverRule(ctx context.Context, input *route53resolver.DisassociateResolverRuleInput) (*route53resolver.DisassociateResolverRuleOutput, error) {
	return a.client.DisassociateResolverRuleWithContext(ctx, input)
}

func (a *Route53ResolverActivities) GetResolverEndpoint(ctx context.Context, input *route53resolver.GetResolverEndpointInput) (*route53resolver.GetResolverEndpointOutput, error) {
	return a.client.GetResolverEndpointWithContext(ctx, input)
}

func (a *Route53ResolverActivities) GetResolverQueryLogConfig(ctx context.Context, input *route53resolver.GetResolverQueryLogConfigInput) (*route53resolver.GetResolverQueryLogConfigOutput, error) {
	return a.client.GetResolverQueryLogConfigWithContext(ctx, input)
}

func (a *Route53ResolverActivities) GetResolverQueryLogConfigAssociation(ctx context.Context, input *route53resolver.GetResolverQueryLogConfigAssociationInput) (*route53resolver.GetResolverQueryLogConfigAssociationOutput, error) {
	return a.client.GetResolverQueryLogConfigAssociationWithContext(ctx, input)
}

func (a *Route53ResolverActivities) GetResolverQueryLogConfigPolicy(ctx context.Context, input *route53resolver.GetResolverQueryLogConfigPolicyInput) (*route53resolver.GetResolverQueryLogConfigPolicyOutput, error) {
	return a.client.GetResolverQueryLogConfigPolicyWithContext(ctx, input)
}

func (a *Route53ResolverActivities) GetResolverRule(ctx context.Context, input *route53resolver.GetResolverRuleInput) (*route53resolver.GetResolverRuleOutput, error) {
	return a.client.GetResolverRuleWithContext(ctx, input)
}

func (a *Route53ResolverActivities) GetResolverRuleAssociation(ctx context.Context, input *route53resolver.GetResolverRuleAssociationInput) (*route53resolver.GetResolverRuleAssociationOutput, error) {
	return a.client.GetResolverRuleAssociationWithContext(ctx, input)
}

func (a *Route53ResolverActivities) GetResolverRulePolicy(ctx context.Context, input *route53resolver.GetResolverRulePolicyInput) (*route53resolver.GetResolverRulePolicyOutput, error) {
	return a.client.GetResolverRulePolicyWithContext(ctx, input)
}

func (a *Route53ResolverActivities) ListResolverEndpointIpAddresses(ctx context.Context, input *route53resolver.ListResolverEndpointIpAddressesInput) (*route53resolver.ListResolverEndpointIpAddressesOutput, error) {
	return a.client.ListResolverEndpointIpAddressesWithContext(ctx, input)
}

func (a *Route53ResolverActivities) ListResolverEndpoints(ctx context.Context, input *route53resolver.ListResolverEndpointsInput) (*route53resolver.ListResolverEndpointsOutput, error) {
	return a.client.ListResolverEndpointsWithContext(ctx, input)
}

func (a *Route53ResolverActivities) ListResolverQueryLogConfigAssociations(ctx context.Context, input *route53resolver.ListResolverQueryLogConfigAssociationsInput) (*route53resolver.ListResolverQueryLogConfigAssociationsOutput, error) {
	return a.client.ListResolverQueryLogConfigAssociationsWithContext(ctx, input)
}

func (a *Route53ResolverActivities) ListResolverQueryLogConfigs(ctx context.Context, input *route53resolver.ListResolverQueryLogConfigsInput) (*route53resolver.ListResolverQueryLogConfigsOutput, error) {
	return a.client.ListResolverQueryLogConfigsWithContext(ctx, input)
}

func (a *Route53ResolverActivities) ListResolverRuleAssociations(ctx context.Context, input *route53resolver.ListResolverRuleAssociationsInput) (*route53resolver.ListResolverRuleAssociationsOutput, error) {
	return a.client.ListResolverRuleAssociationsWithContext(ctx, input)
}

func (a *Route53ResolverActivities) ListResolverRules(ctx context.Context, input *route53resolver.ListResolverRulesInput) (*route53resolver.ListResolverRulesOutput, error) {
	return a.client.ListResolverRulesWithContext(ctx, input)
}

func (a *Route53ResolverActivities) ListTagsForResource(ctx context.Context, input *route53resolver.ListTagsForResourceInput) (*route53resolver.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *Route53ResolverActivities) PutResolverQueryLogConfigPolicy(ctx context.Context, input *route53resolver.PutResolverQueryLogConfigPolicyInput) (*route53resolver.PutResolverQueryLogConfigPolicyOutput, error) {
	return a.client.PutResolverQueryLogConfigPolicyWithContext(ctx, input)
}

func (a *Route53ResolverActivities) PutResolverRulePolicy(ctx context.Context, input *route53resolver.PutResolverRulePolicyInput) (*route53resolver.PutResolverRulePolicyOutput, error) {
	return a.client.PutResolverRulePolicyWithContext(ctx, input)
}

func (a *Route53ResolverActivities) TagResource(ctx context.Context, input *route53resolver.TagResourceInput) (*route53resolver.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *Route53ResolverActivities) UntagResource(ctx context.Context, input *route53resolver.UntagResourceInput) (*route53resolver.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *Route53ResolverActivities) UpdateResolverEndpoint(ctx context.Context, input *route53resolver.UpdateResolverEndpointInput) (*route53resolver.UpdateResolverEndpointOutput, error) {
	return a.client.UpdateResolverEndpointWithContext(ctx, input)
}

func (a *Route53ResolverActivities) UpdateResolverRule(ctx context.Context, input *route53resolver.UpdateResolverRuleInput) (*route53resolver.UpdateResolverRuleOutput, error) {
	return a.client.UpdateResolverRuleWithContext(ctx, input)
}
