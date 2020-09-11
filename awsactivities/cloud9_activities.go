package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloud9"
	"github.com/aws/aws-sdk-go/service/cloud9/cloud9iface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type Cloud9Activities struct {
	client cloud9iface.Cloud9API
}

func NewCloud9Activities(session *session.Session, config ...*aws.Config) *Cloud9Activities {
	client := cloud9.New(session, config...)
	return &Cloud9Activities{client: client}
}

func (a *Cloud9Activities) CreateEnvironmentEC2(ctx context.Context, input *cloud9.CreateEnvironmentEC2Input) (*cloud9.CreateEnvironmentEC2Output, error) {
	return a.client.CreateEnvironmentEC2WithContext(ctx, input)
}

func (a *Cloud9Activities) CreateEnvironmentMembership(ctx context.Context, input *cloud9.CreateEnvironmentMembershipInput) (*cloud9.CreateEnvironmentMembershipOutput, error) {
	return a.client.CreateEnvironmentMembershipWithContext(ctx, input)
}

func (a *Cloud9Activities) DeleteEnvironment(ctx context.Context, input *cloud9.DeleteEnvironmentInput) (*cloud9.DeleteEnvironmentOutput, error) {
	return a.client.DeleteEnvironmentWithContext(ctx, input)
}

func (a *Cloud9Activities) DeleteEnvironmentMembership(ctx context.Context, input *cloud9.DeleteEnvironmentMembershipInput) (*cloud9.DeleteEnvironmentMembershipOutput, error) {
	return a.client.DeleteEnvironmentMembershipWithContext(ctx, input)
}

func (a *Cloud9Activities) DescribeEnvironmentMemberships(ctx context.Context, input *cloud9.DescribeEnvironmentMembershipsInput) (*cloud9.DescribeEnvironmentMembershipsOutput, error) {
	return a.client.DescribeEnvironmentMembershipsWithContext(ctx, input)
}

func (a *Cloud9Activities) DescribeEnvironmentStatus(ctx context.Context, input *cloud9.DescribeEnvironmentStatusInput) (*cloud9.DescribeEnvironmentStatusOutput, error) {
	return a.client.DescribeEnvironmentStatusWithContext(ctx, input)
}

func (a *Cloud9Activities) DescribeEnvironments(ctx context.Context, input *cloud9.DescribeEnvironmentsInput) (*cloud9.DescribeEnvironmentsOutput, error) {
	return a.client.DescribeEnvironmentsWithContext(ctx, input)
}

func (a *Cloud9Activities) ListEnvironments(ctx context.Context, input *cloud9.ListEnvironmentsInput) (*cloud9.ListEnvironmentsOutput, error) {
	return a.client.ListEnvironmentsWithContext(ctx, input)
}

func (a *Cloud9Activities) ListTagsForResource(ctx context.Context, input *cloud9.ListTagsForResourceInput) (*cloud9.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *Cloud9Activities) TagResource(ctx context.Context, input *cloud9.TagResourceInput) (*cloud9.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *Cloud9Activities) UntagResource(ctx context.Context, input *cloud9.UntagResourceInput) (*cloud9.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *Cloud9Activities) UpdateEnvironment(ctx context.Context, input *cloud9.UpdateEnvironmentInput) (*cloud9.UpdateEnvironmentOutput, error) {
	return a.client.UpdateEnvironmentWithContext(ctx, input)
}

func (a *Cloud9Activities) UpdateEnvironmentMembership(ctx context.Context, input *cloud9.UpdateEnvironmentMembershipInput) (*cloud9.UpdateEnvironmentMembershipOutput, error) {
	return a.client.UpdateEnvironmentMembershipWithContext(ctx, input)
}
