
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elb/elbiface"
)

type ELBActivities struct {
	client elbiface.ELBAPI
}

func NewELBActivities(session *session.Session, config... *aws.Config) *ELBActivities {
    client := elb.New(session, config...)
    return &ELBActivities{client: client}
}

func (a *ELBActivities) AddTags(input *elb.AddTagsInput) (*elb.AddTagsOutput, error) {
    return a.client.AddTags(input)
}

func (a *ELBActivities) ApplySecurityGroupsToLoadBalancer(input *elb.ApplySecurityGroupsToLoadBalancerInput) (*elb.ApplySecurityGroupsToLoadBalancerOutput, error) {
    return a.client.ApplySecurityGroupsToLoadBalancer(input)
}

func (a *ELBActivities) AttachLoadBalancerToSubnets(input *elb.AttachLoadBalancerToSubnetsInput) (*elb.AttachLoadBalancerToSubnetsOutput, error) {
    return a.client.AttachLoadBalancerToSubnets(input)
}

func (a *ELBActivities) ConfigureHealthCheck(input *elb.ConfigureHealthCheckInput) (*elb.ConfigureHealthCheckOutput, error) {
    return a.client.ConfigureHealthCheck(input)
}

func (a *ELBActivities) CreateAppCookieStickinessPolicy(input *elb.CreateAppCookieStickinessPolicyInput) (*elb.CreateAppCookieStickinessPolicyOutput, error) {
    return a.client.CreateAppCookieStickinessPolicy(input)
}

func (a *ELBActivities) CreateLBCookieStickinessPolicy(input *elb.CreateLBCookieStickinessPolicyInput) (*elb.CreateLBCookieStickinessPolicyOutput, error) {
    return a.client.CreateLBCookieStickinessPolicy(input)
}

func (a *ELBActivities) CreateLoadBalancer(input *elb.CreateLoadBalancerInput) (*elb.CreateLoadBalancerOutput, error) {
    return a.client.CreateLoadBalancer(input)
}

func (a *ELBActivities) CreateLoadBalancerListeners(input *elb.CreateLoadBalancerListenersInput) (*elb.CreateLoadBalancerListenersOutput, error) {
    return a.client.CreateLoadBalancerListeners(input)
}

func (a *ELBActivities) CreateLoadBalancerPolicy(input *elb.CreateLoadBalancerPolicyInput) (*elb.CreateLoadBalancerPolicyOutput, error) {
    return a.client.CreateLoadBalancerPolicy(input)
}

func (a *ELBActivities) DeleteLoadBalancer(input *elb.DeleteLoadBalancerInput) (*elb.DeleteLoadBalancerOutput, error) {
    return a.client.DeleteLoadBalancer(input)
}

func (a *ELBActivities) DeleteLoadBalancerListeners(input *elb.DeleteLoadBalancerListenersInput) (*elb.DeleteLoadBalancerListenersOutput, error) {
    return a.client.DeleteLoadBalancerListeners(input)
}

func (a *ELBActivities) DeleteLoadBalancerPolicy(input *elb.DeleteLoadBalancerPolicyInput) (*elb.DeleteLoadBalancerPolicyOutput, error) {
    return a.client.DeleteLoadBalancerPolicy(input)
}

func (a *ELBActivities) DeregisterInstancesFromLoadBalancer(input *elb.DeregisterInstancesFromLoadBalancerInput) (*elb.DeregisterInstancesFromLoadBalancerOutput, error) {
    return a.client.DeregisterInstancesFromLoadBalancer(input)
}

func (a *ELBActivities) DescribeAccountLimits(input *elb.DescribeAccountLimitsInput) (*elb.DescribeAccountLimitsOutput, error) {
    return a.client.DescribeAccountLimits(input)
}

func (a *ELBActivities) DescribeInstanceHealth(input *elb.DescribeInstanceHealthInput) (*elb.DescribeInstanceHealthOutput, error) {
    return a.client.DescribeInstanceHealth(input)
}

func (a *ELBActivities) DescribeLoadBalancerAttributes(input *elb.DescribeLoadBalancerAttributesInput) (*elb.DescribeLoadBalancerAttributesOutput, error) {
    return a.client.DescribeLoadBalancerAttributes(input)
}

func (a *ELBActivities) DescribeLoadBalancerPolicies(input *elb.DescribeLoadBalancerPoliciesInput) (*elb.DescribeLoadBalancerPoliciesOutput, error) {
    return a.client.DescribeLoadBalancerPolicies(input)
}

func (a *ELBActivities) DescribeLoadBalancerPolicyTypes(input *elb.DescribeLoadBalancerPolicyTypesInput) (*elb.DescribeLoadBalancerPolicyTypesOutput, error) {
    return a.client.DescribeLoadBalancerPolicyTypes(input)
}

func (a *ELBActivities) DescribeLoadBalancers(input *elb.DescribeLoadBalancersInput) (*elb.DescribeLoadBalancersOutput, error) {
    return a.client.DescribeLoadBalancers(input)
}

func (a *ELBActivities) DescribeTags(input *elb.DescribeTagsInput) (*elb.DescribeTagsOutput, error) {
    return a.client.DescribeTags(input)
}

func (a *ELBActivities) DetachLoadBalancerFromSubnets(input *elb.DetachLoadBalancerFromSubnetsInput) (*elb.DetachLoadBalancerFromSubnetsOutput, error) {
    return a.client.DetachLoadBalancerFromSubnets(input)
}

func (a *ELBActivities) DisableAvailabilityZonesForLoadBalancer(input *elb.DisableAvailabilityZonesForLoadBalancerInput) (*elb.DisableAvailabilityZonesForLoadBalancerOutput, error) {
    return a.client.DisableAvailabilityZonesForLoadBalancer(input)
}

func (a *ELBActivities) EnableAvailabilityZonesForLoadBalancer(input *elb.EnableAvailabilityZonesForLoadBalancerInput) (*elb.EnableAvailabilityZonesForLoadBalancerOutput, error) {
    return a.client.EnableAvailabilityZonesForLoadBalancer(input)
}

func (a *ELBActivities) ModifyLoadBalancerAttributes(input *elb.ModifyLoadBalancerAttributesInput) (*elb.ModifyLoadBalancerAttributesOutput, error) {
    return a.client.ModifyLoadBalancerAttributes(input)
}

func (a *ELBActivities) RegisterInstancesWithLoadBalancer(input *elb.RegisterInstancesWithLoadBalancerInput) (*elb.RegisterInstancesWithLoadBalancerOutput, error) {
    return a.client.RegisterInstancesWithLoadBalancer(input)
}

func (a *ELBActivities) RemoveTags(input *elb.RemoveTagsInput) (*elb.RemoveTagsOutput, error) {
    return a.client.RemoveTags(input)
}

func (a *ELBActivities) SetLoadBalancerListenerSSLCertificate(input *elb.SetLoadBalancerListenerSSLCertificateInput) (*elb.SetLoadBalancerListenerSSLCertificateOutput, error) {
    return a.client.SetLoadBalancerListenerSSLCertificate(input)
}

func (a *ELBActivities) SetLoadBalancerPoliciesForBackendServer(input *elb.SetLoadBalancerPoliciesForBackendServerInput) (*elb.SetLoadBalancerPoliciesForBackendServerOutput, error) {
    return a.client.SetLoadBalancerPoliciesForBackendServer(input)
}

func (a *ELBActivities) SetLoadBalancerPoliciesOfListener(input *elb.SetLoadBalancerPoliciesOfListenerInput) (*elb.SetLoadBalancerPoliciesOfListenerOutput, error) {
    return a.client.SetLoadBalancerPoliciesOfListener(input)
}

func (a *ELBActivities) WaitUntilAnyInstanceInService(input *elb.DescribeInstanceHealthInput) error {
    return a.client.WaitUntilAnyInstanceInService(input)
}

func (a *ELBActivities) WaitUntilInstanceDeregistered(input *elb.DescribeInstanceHealthInput) error {
    return a.client.WaitUntilInstanceDeregistered(input)
}

func (a *ELBActivities) WaitUntilInstanceInService(input *elb.DescribeInstanceHealthInput) error {
    return a.client.WaitUntilInstanceInService(input)
}
