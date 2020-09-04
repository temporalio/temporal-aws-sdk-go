
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/cloud9"
	"github.com/aws/aws-sdk-go/service/cloud9/cloud9iface"
)

type Cloud9Activities struct {
	client cloud9iface.Cloud9API
}

func NewCloud9Activities(client cloud9iface.Cloud9API) *Cloud9Activities {
    return &Cloud9Activities{client: client}
}

func (a *Cloud9Activities) CreateEnvironmentEC2(input *cloud9.CreateEnvironmentEC2Input) (*cloud9.CreateEnvironmentEC2Output, error) {
    return a.client.CreateEnvironmentEC2(input)
}

func (a *Cloud9Activities) CreateEnvironmentMembership(input *cloud9.CreateEnvironmentMembershipInput) (*cloud9.CreateEnvironmentMembershipOutput, error) {
    return a.client.CreateEnvironmentMembership(input)
}

func (a *Cloud9Activities) DeleteEnvironment(input *cloud9.DeleteEnvironmentInput) (*cloud9.DeleteEnvironmentOutput, error) {
    return a.client.DeleteEnvironment(input)
}

func (a *Cloud9Activities) DeleteEnvironmentMembership(input *cloud9.DeleteEnvironmentMembershipInput) (*cloud9.DeleteEnvironmentMembershipOutput, error) {
    return a.client.DeleteEnvironmentMembership(input)
}

func (a *Cloud9Activities) DescribeEnvironmentMemberships(input *cloud9.DescribeEnvironmentMembershipsInput) (*cloud9.DescribeEnvironmentMembershipsOutput, error) {
    return a.client.DescribeEnvironmentMemberships(input)
}

func (a *Cloud9Activities) DescribeEnvironmentStatus(input *cloud9.DescribeEnvironmentStatusInput) (*cloud9.DescribeEnvironmentStatusOutput, error) {
    return a.client.DescribeEnvironmentStatus(input)
}

func (a *Cloud9Activities) DescribeEnvironments(input *cloud9.DescribeEnvironmentsInput) (*cloud9.DescribeEnvironmentsOutput, error) {
    return a.client.DescribeEnvironments(input)
}

func (a *Cloud9Activities) ListEnvironments(input *cloud9.ListEnvironmentsInput) (*cloud9.ListEnvironmentsOutput, error) {
    return a.client.ListEnvironments(input)
}

func (a *Cloud9Activities) ListTagsForResource(input *cloud9.ListTagsForResourceInput) (*cloud9.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *Cloud9Activities) TagResource(input *cloud9.TagResourceInput) (*cloud9.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *Cloud9Activities) UntagResource(input *cloud9.UntagResourceInput) (*cloud9.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *Cloud9Activities) UpdateEnvironment(input *cloud9.UpdateEnvironmentInput) (*cloud9.UpdateEnvironmentOutput, error) {
    return a.client.UpdateEnvironment(input)
}

func (a *Cloud9Activities) UpdateEnvironmentMembership(input *cloud9.UpdateEnvironmentMembershipInput) (*cloud9.UpdateEnvironmentMembershipOutput, error) {
    return a.client.UpdateEnvironmentMembership(input)
}
