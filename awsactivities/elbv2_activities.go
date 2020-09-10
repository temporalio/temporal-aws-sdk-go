package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type ELBV2Activities struct {
	client elbv2iface.ELBV2API
}

func NewELBV2Activities(session *session.Session, config ...*aws.Config) *ELBV2Activities {
	client := elbv2.New(session, config...)
	return &ELBV2Activities{client: client}
}

func (a *ELBV2Activities) AddListenerCertificates(ctx context.Context, input *elbv2.AddListenerCertificatesInput) (*elbv2.AddListenerCertificatesOutput, error) {
	return a.client.AddListenerCertificatesWithContext(ctx, input)
}

func (a *ELBV2Activities) AddTags(ctx context.Context, input *elbv2.AddTagsInput) (*elbv2.AddTagsOutput, error) {
	return a.client.AddTagsWithContext(ctx, input)
}

func (a *ELBV2Activities) CreateListener(ctx context.Context, input *elbv2.CreateListenerInput) (*elbv2.CreateListenerOutput, error) {
	return a.client.CreateListenerWithContext(ctx, input)
}

func (a *ELBV2Activities) CreateLoadBalancer(ctx context.Context, input *elbv2.CreateLoadBalancerInput) (*elbv2.CreateLoadBalancerOutput, error) {
	return a.client.CreateLoadBalancerWithContext(ctx, input)
}

func (a *ELBV2Activities) CreateRule(ctx context.Context, input *elbv2.CreateRuleInput) (*elbv2.CreateRuleOutput, error) {
	return a.client.CreateRuleWithContext(ctx, input)
}

func (a *ELBV2Activities) CreateTargetGroup(ctx context.Context, input *elbv2.CreateTargetGroupInput) (*elbv2.CreateTargetGroupOutput, error) {
	return a.client.CreateTargetGroupWithContext(ctx, input)
}

func (a *ELBV2Activities) DeleteListener(ctx context.Context, input *elbv2.DeleteListenerInput) (*elbv2.DeleteListenerOutput, error) {
	return a.client.DeleteListenerWithContext(ctx, input)
}

func (a *ELBV2Activities) DeleteLoadBalancer(ctx context.Context, input *elbv2.DeleteLoadBalancerInput) (*elbv2.DeleteLoadBalancerOutput, error) {
	return a.client.DeleteLoadBalancerWithContext(ctx, input)
}

func (a *ELBV2Activities) DeleteRule(ctx context.Context, input *elbv2.DeleteRuleInput) (*elbv2.DeleteRuleOutput, error) {
	return a.client.DeleteRuleWithContext(ctx, input)
}

func (a *ELBV2Activities) DeleteTargetGroup(ctx context.Context, input *elbv2.DeleteTargetGroupInput) (*elbv2.DeleteTargetGroupOutput, error) {
	return a.client.DeleteTargetGroupWithContext(ctx, input)
}

func (a *ELBV2Activities) DeregisterTargets(ctx context.Context, input *elbv2.DeregisterTargetsInput) (*elbv2.DeregisterTargetsOutput, error) {
	return a.client.DeregisterTargetsWithContext(ctx, input)
}

func (a *ELBV2Activities) DescribeAccountLimits(ctx context.Context, input *elbv2.DescribeAccountLimitsInput) (*elbv2.DescribeAccountLimitsOutput, error) {
	return a.client.DescribeAccountLimitsWithContext(ctx, input)
}

func (a *ELBV2Activities) DescribeListenerCertificates(ctx context.Context, input *elbv2.DescribeListenerCertificatesInput) (*elbv2.DescribeListenerCertificatesOutput, error) {
	return a.client.DescribeListenerCertificatesWithContext(ctx, input)
}

func (a *ELBV2Activities) DescribeListeners(ctx context.Context, input *elbv2.DescribeListenersInput) (*elbv2.DescribeListenersOutput, error) {
	return a.client.DescribeListenersWithContext(ctx, input)
}

func (a *ELBV2Activities) DescribeLoadBalancerAttributes(ctx context.Context, input *elbv2.DescribeLoadBalancerAttributesInput) (*elbv2.DescribeLoadBalancerAttributesOutput, error) {
	return a.client.DescribeLoadBalancerAttributesWithContext(ctx, input)
}

func (a *ELBV2Activities) DescribeLoadBalancers(ctx context.Context, input *elbv2.DescribeLoadBalancersInput) (*elbv2.DescribeLoadBalancersOutput, error) {
	return a.client.DescribeLoadBalancersWithContext(ctx, input)
}

func (a *ELBV2Activities) DescribeRules(ctx context.Context, input *elbv2.DescribeRulesInput) (*elbv2.DescribeRulesOutput, error) {
	return a.client.DescribeRulesWithContext(ctx, input)
}

func (a *ELBV2Activities) DescribeSSLPolicies(ctx context.Context, input *elbv2.DescribeSSLPoliciesInput) (*elbv2.DescribeSSLPoliciesOutput, error) {
	return a.client.DescribeSSLPoliciesWithContext(ctx, input)
}

func (a *ELBV2Activities) DescribeTags(ctx context.Context, input *elbv2.DescribeTagsInput) (*elbv2.DescribeTagsOutput, error) {
	return a.client.DescribeTagsWithContext(ctx, input)
}

func (a *ELBV2Activities) DescribeTargetGroupAttributes(ctx context.Context, input *elbv2.DescribeTargetGroupAttributesInput) (*elbv2.DescribeTargetGroupAttributesOutput, error) {
	return a.client.DescribeTargetGroupAttributesWithContext(ctx, input)
}

func (a *ELBV2Activities) DescribeTargetGroups(ctx context.Context, input *elbv2.DescribeTargetGroupsInput) (*elbv2.DescribeTargetGroupsOutput, error) {
	return a.client.DescribeTargetGroupsWithContext(ctx, input)
}

func (a *ELBV2Activities) DescribeTargetHealth(ctx context.Context, input *elbv2.DescribeTargetHealthInput) (*elbv2.DescribeTargetHealthOutput, error) {
	return a.client.DescribeTargetHealthWithContext(ctx, input)
}

func (a *ELBV2Activities) ModifyListener(ctx context.Context, input *elbv2.ModifyListenerInput) (*elbv2.ModifyListenerOutput, error) {
	return a.client.ModifyListenerWithContext(ctx, input)
}

func (a *ELBV2Activities) ModifyLoadBalancerAttributes(ctx context.Context, input *elbv2.ModifyLoadBalancerAttributesInput) (*elbv2.ModifyLoadBalancerAttributesOutput, error) {
	return a.client.ModifyLoadBalancerAttributesWithContext(ctx, input)
}

func (a *ELBV2Activities) ModifyRule(ctx context.Context, input *elbv2.ModifyRuleInput) (*elbv2.ModifyRuleOutput, error) {
	return a.client.ModifyRuleWithContext(ctx, input)
}

func (a *ELBV2Activities) ModifyTargetGroup(ctx context.Context, input *elbv2.ModifyTargetGroupInput) (*elbv2.ModifyTargetGroupOutput, error) {
	return a.client.ModifyTargetGroupWithContext(ctx, input)
}

func (a *ELBV2Activities) ModifyTargetGroupAttributes(ctx context.Context, input *elbv2.ModifyTargetGroupAttributesInput) (*elbv2.ModifyTargetGroupAttributesOutput, error) {
	return a.client.ModifyTargetGroupAttributesWithContext(ctx, input)
}

func (a *ELBV2Activities) RegisterTargets(ctx context.Context, input *elbv2.RegisterTargetsInput) (*elbv2.RegisterTargetsOutput, error) {
	return a.client.RegisterTargetsWithContext(ctx, input)
}

func (a *ELBV2Activities) RemoveListenerCertificates(ctx context.Context, input *elbv2.RemoveListenerCertificatesInput) (*elbv2.RemoveListenerCertificatesOutput, error) {
	return a.client.RemoveListenerCertificatesWithContext(ctx, input)
}

func (a *ELBV2Activities) RemoveTags(ctx context.Context, input *elbv2.RemoveTagsInput) (*elbv2.RemoveTagsOutput, error) {
	return a.client.RemoveTagsWithContext(ctx, input)
}

func (a *ELBV2Activities) SetIpAddressType(ctx context.Context, input *elbv2.SetIpAddressTypeInput) (*elbv2.SetIpAddressTypeOutput, error) {
	return a.client.SetIpAddressTypeWithContext(ctx, input)
}

func (a *ELBV2Activities) SetRulePriorities(ctx context.Context, input *elbv2.SetRulePrioritiesInput) (*elbv2.SetRulePrioritiesOutput, error) {
	return a.client.SetRulePrioritiesWithContext(ctx, input)
}

func (a *ELBV2Activities) SetSecurityGroups(ctx context.Context, input *elbv2.SetSecurityGroupsInput) (*elbv2.SetSecurityGroupsOutput, error) {
	return a.client.SetSecurityGroupsWithContext(ctx, input)
}

func (a *ELBV2Activities) SetSubnets(ctx context.Context, input *elbv2.SetSubnetsInput) (*elbv2.SetSubnetsOutput, error) {
	return a.client.SetSubnetsWithContext(ctx, input)
}

func (a *ELBV2Activities) WaitUntilLoadBalancerAvailable(ctx context.Context, input *elbv2.DescribeLoadBalancersInput) error {
	return a.client.WaitUntilLoadBalancerAvailableWithContext(ctx, input)

}

func (a *ELBV2Activities) WaitUntilLoadBalancerExists(ctx context.Context, input *elbv2.DescribeLoadBalancersInput) error {
	return a.client.WaitUntilLoadBalancerExistsWithContext(ctx, input)

}

func (a *ELBV2Activities) WaitUntilLoadBalancersDeleted(ctx context.Context, input *elbv2.DescribeLoadBalancersInput) error {
	return a.client.WaitUntilLoadBalancersDeletedWithContext(ctx, input)

}

func (a *ELBV2Activities) WaitUntilTargetDeregistered(ctx context.Context, input *elbv2.DescribeTargetHealthInput) error {
	return a.client.WaitUntilTargetDeregisteredWithContext(ctx, input)

}

func (a *ELBV2Activities) WaitUntilTargetInService(ctx context.Context, input *elbv2.DescribeTargetHealthInput) error {
	return a.client.WaitUntilTargetInServiceWithContext(ctx, input)

}
