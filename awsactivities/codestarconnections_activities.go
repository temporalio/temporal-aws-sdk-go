package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codestarconnections"
	"github.com/aws/aws-sdk-go/service/codestarconnections/codestarconnectionsiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type CodeStarConnectionsActivities struct {
	client codestarconnectionsiface.CodeStarConnectionsAPI
}

func NewCodeStarConnectionsActivities(session *session.Session, config ...*aws.Config) *CodeStarConnectionsActivities {
	client := codestarconnections.New(session, config...)
	return &CodeStarConnectionsActivities{client: client}
}

func (a *CodeStarConnectionsActivities) CreateConnection(ctx context.Context, input *codestarconnections.CreateConnectionInput) (*codestarconnections.CreateConnectionOutput, error) {
	return a.client.CreateConnectionWithContext(ctx, input)
}

func (a *CodeStarConnectionsActivities) CreateHost(ctx context.Context, input *codestarconnections.CreateHostInput) (*codestarconnections.CreateHostOutput, error) {
	return a.client.CreateHostWithContext(ctx, input)
}

func (a *CodeStarConnectionsActivities) DeleteConnection(ctx context.Context, input *codestarconnections.DeleteConnectionInput) (*codestarconnections.DeleteConnectionOutput, error) {
	return a.client.DeleteConnectionWithContext(ctx, input)
}

func (a *CodeStarConnectionsActivities) DeleteHost(ctx context.Context, input *codestarconnections.DeleteHostInput) (*codestarconnections.DeleteHostOutput, error) {
	return a.client.DeleteHostWithContext(ctx, input)
}

func (a *CodeStarConnectionsActivities) GetConnection(ctx context.Context, input *codestarconnections.GetConnectionInput) (*codestarconnections.GetConnectionOutput, error) {
	return a.client.GetConnectionWithContext(ctx, input)
}

func (a *CodeStarConnectionsActivities) GetHost(ctx context.Context, input *codestarconnections.GetHostInput) (*codestarconnections.GetHostOutput, error) {
	return a.client.GetHostWithContext(ctx, input)
}

func (a *CodeStarConnectionsActivities) ListConnections(ctx context.Context, input *codestarconnections.ListConnectionsInput) (*codestarconnections.ListConnectionsOutput, error) {
	return a.client.ListConnectionsWithContext(ctx, input)
}

func (a *CodeStarConnectionsActivities) ListHosts(ctx context.Context, input *codestarconnections.ListHostsInput) (*codestarconnections.ListHostsOutput, error) {
	return a.client.ListHostsWithContext(ctx, input)
}

func (a *CodeStarConnectionsActivities) ListTagsForResource(ctx context.Context, input *codestarconnections.ListTagsForResourceInput) (*codestarconnections.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *CodeStarConnectionsActivities) TagResource(ctx context.Context, input *codestarconnections.TagResourceInput) (*codestarconnections.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *CodeStarConnectionsActivities) UntagResource(ctx context.Context, input *codestarconnections.UntagResourceInput) (*codestarconnections.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}
