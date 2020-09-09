
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"
)

type ELBV2Activities struct {
    client elbv2iface.ELBV2API
}

func NewELBV2Activities(session *session.Session, config ...*aws.Config) *ELBV2Activities {
    client := elbv2.New(session, config...)
    return &ELBV2Activities{client: client}
}

func (a *ELBV2Activities) AddListenerCertificates(input *elbv2.AddListenerCertificatesInput) (*elbv2.AddListenerCertificatesOutput, error) {
    return a.client.AddListenerCertificates(input)
}

func (a *ELBV2Activities) AddTags(input *elbv2.AddTagsInput) (*elbv2.AddTagsOutput, error) {
    return a.client.AddTags(input)
}

func (a *ELBV2Activities) CreateListener(input *elbv2.CreateListenerInput) (*elbv2.CreateListenerOutput, error) {
    return a.client.CreateListener(input)
}

func (a *ELBV2Activities) CreateLoadBalancer(input *elbv2.CreateLoadBalancerInput) (*elbv2.CreateLoadBalancerOutput, error) {
    return a.client.CreateLoadBalancer(input)
}

func (a *ELBV2Activities) CreateRule(input *elbv2.CreateRuleInput) (*elbv2.CreateRuleOutput, error) {
    return a.client.CreateRule(input)
}

func (a *ELBV2Activities) CreateTargetGroup(input *elbv2.CreateTargetGroupInput) (*elbv2.CreateTargetGroupOutput, error) {
    return a.client.CreateTargetGroup(input)
}

func (a *ELBV2Activities) DeleteListener(input *elbv2.DeleteListenerInput) (*elbv2.DeleteListenerOutput, error) {
    return a.client.DeleteListener(input)
}

func (a *ELBV2Activities) DeleteLoadBalancer(input *elbv2.DeleteLoadBalancerInput) (*elbv2.DeleteLoadBalancerOutput, error) {
    return a.client.DeleteLoadBalancer(input)
}

func (a *ELBV2Activities) DeleteRule(input *elbv2.DeleteRuleInput) (*elbv2.DeleteRuleOutput, error) {
    return a.client.DeleteRule(input)
}

func (a *ELBV2Activities) DeleteTargetGroup(input *elbv2.DeleteTargetGroupInput) (*elbv2.DeleteTargetGroupOutput, error) {
    return a.client.DeleteTargetGroup(input)
}

func (a *ELBV2Activities) DeregisterTargets(input *elbv2.DeregisterTargetsInput) (*elbv2.DeregisterTargetsOutput, error) {
    return a.client.DeregisterTargets(input)
}

func (a *ELBV2Activities) DescribeAccountLimits(input *elbv2.DescribeAccountLimitsInput) (*elbv2.DescribeAccountLimitsOutput, error) {
    return a.client.DescribeAccountLimits(input)
}

func (a *ELBV2Activities) DescribeListenerCertificates(input *elbv2.DescribeListenerCertificatesInput) (*elbv2.DescribeListenerCertificatesOutput, error) {
    return a.client.DescribeListenerCertificates(input)
}

func (a *ELBV2Activities) DescribeListeners(input *elbv2.DescribeListenersInput) (*elbv2.DescribeListenersOutput, error) {
    return a.client.DescribeListeners(input)
}

func (a *ELBV2Activities) DescribeLoadBalancerAttributes(input *elbv2.DescribeLoadBalancerAttributesInput) (*elbv2.DescribeLoadBalancerAttributesOutput, error) {
    return a.client.DescribeLoadBalancerAttributes(input)
}

func (a *ELBV2Activities) DescribeLoadBalancers(input *elbv2.DescribeLoadBalancersInput) (*elbv2.DescribeLoadBalancersOutput, error) {
    return a.client.DescribeLoadBalancers(input)
}

func (a *ELBV2Activities) DescribeRules(input *elbv2.DescribeRulesInput) (*elbv2.DescribeRulesOutput, error) {
    return a.client.DescribeRules(input)
}

func (a *ELBV2Activities) DescribeSSLPolicies(input *elbv2.DescribeSSLPoliciesInput) (*elbv2.DescribeSSLPoliciesOutput, error) {
    return a.client.DescribeSSLPolicies(input)
}

func (a *ELBV2Activities) DescribeTags(input *elbv2.DescribeTagsInput) (*elbv2.DescribeTagsOutput, error) {
    return a.client.DescribeTags(input)
}

func (a *ELBV2Activities) DescribeTargetGroupAttributes(input *elbv2.DescribeTargetGroupAttributesInput) (*elbv2.DescribeTargetGroupAttributesOutput, error) {
    return a.client.DescribeTargetGroupAttributes(input)
}

func (a *ELBV2Activities) DescribeTargetGroups(input *elbv2.DescribeTargetGroupsInput) (*elbv2.DescribeTargetGroupsOutput, error) {
    return a.client.DescribeTargetGroups(input)
}

func (a *ELBV2Activities) DescribeTargetHealth(input *elbv2.DescribeTargetHealthInput) (*elbv2.DescribeTargetHealthOutput, error) {
    return a.client.DescribeTargetHealth(input)
}

func (a *ELBV2Activities) ModifyListener(input *elbv2.ModifyListenerInput) (*elbv2.ModifyListenerOutput, error) {
    return a.client.ModifyListener(input)
}

func (a *ELBV2Activities) ModifyLoadBalancerAttributes(input *elbv2.ModifyLoadBalancerAttributesInput) (*elbv2.ModifyLoadBalancerAttributesOutput, error) {
    return a.client.ModifyLoadBalancerAttributes(input)
}

func (a *ELBV2Activities) ModifyRule(input *elbv2.ModifyRuleInput) (*elbv2.ModifyRuleOutput, error) {
    return a.client.ModifyRule(input)
}

func (a *ELBV2Activities) ModifyTargetGroup(input *elbv2.ModifyTargetGroupInput) (*elbv2.ModifyTargetGroupOutput, error) {
    return a.client.ModifyTargetGroup(input)
}

func (a *ELBV2Activities) ModifyTargetGroupAttributes(input *elbv2.ModifyTargetGroupAttributesInput) (*elbv2.ModifyTargetGroupAttributesOutput, error) {
    return a.client.ModifyTargetGroupAttributes(input)
}

func (a *ELBV2Activities) RegisterTargets(input *elbv2.RegisterTargetsInput) (*elbv2.RegisterTargetsOutput, error) {
    return a.client.RegisterTargets(input)
}

func (a *ELBV2Activities) RemoveListenerCertificates(input *elbv2.RemoveListenerCertificatesInput) (*elbv2.RemoveListenerCertificatesOutput, error) {
    return a.client.RemoveListenerCertificates(input)
}

func (a *ELBV2Activities) RemoveTags(input *elbv2.RemoveTagsInput) (*elbv2.RemoveTagsOutput, error) {
    return a.client.RemoveTags(input)
}

func (a *ELBV2Activities) SetIpAddressType(input *elbv2.SetIpAddressTypeInput) (*elbv2.SetIpAddressTypeOutput, error) {
    return a.client.SetIpAddressType(input)
}

func (a *ELBV2Activities) SetRulePriorities(input *elbv2.SetRulePrioritiesInput) (*elbv2.SetRulePrioritiesOutput, error) {
    return a.client.SetRulePriorities(input)
}

func (a *ELBV2Activities) SetSecurityGroups(input *elbv2.SetSecurityGroupsInput) (*elbv2.SetSecurityGroupsOutput, error) {
    return a.client.SetSecurityGroups(input)
}

func (a *ELBV2Activities) SetSubnets(input *elbv2.SetSubnetsInput) (*elbv2.SetSubnetsOutput, error) {
    return a.client.SetSubnets(input)
}

func (a *ELBV2Activities) WaitUntilLoadBalancerAvailable(input *elbv2.DescribeLoadBalancersInput) error {
    return a.client.WaitUntilLoadBalancerAvailable(input)
}

func (a *ELBV2Activities) WaitUntilLoadBalancerExists(input *elbv2.DescribeLoadBalancersInput) error {
    return a.client.WaitUntilLoadBalancerExists(input)
}

func (a *ELBV2Activities) WaitUntilLoadBalancersDeleted(input *elbv2.DescribeLoadBalancersInput) error {
    return a.client.WaitUntilLoadBalancersDeleted(input)
}

func (a *ELBV2Activities) WaitUntilTargetDeregistered(input *elbv2.DescribeTargetHealthInput) error {
    return a.client.WaitUntilTargetDeregistered(input)
}

func (a *ELBV2Activities) WaitUntilTargetInService(input *elbv2.DescribeTargetHealthInput) error {
    return a.client.WaitUntilTargetInService(input)
}
