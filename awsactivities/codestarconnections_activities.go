package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codestarconnections"
	"github.com/aws/aws-sdk-go/service/codestarconnections/codestarconnectionsiface"
)

type CodeStarConnectionsActivities struct {
	client codestarconnectionsiface.CodeStarConnectionsAPI
}

func NewCodeStarConnectionsActivities(session *session.Session, config ...*aws.Config) *CodeStarConnectionsActivities {
	client := codestarconnections.New(session, config...)
	return &CodeStarConnectionsActivities{client: client}
}

func (a *CodeStarConnectionsActivities) CreateConnection(input *codestarconnections.CreateConnectionInput) (*codestarconnections.CreateConnectionOutput, error) {
	return a.client.CreateConnection(input)
}

func (a *CodeStarConnectionsActivities) CreateHost(input *codestarconnections.CreateHostInput) (*codestarconnections.CreateHostOutput, error) {
	return a.client.CreateHost(input)
}

func (a *CodeStarConnectionsActivities) DeleteConnection(input *codestarconnections.DeleteConnectionInput) (*codestarconnections.DeleteConnectionOutput, error) {
	return a.client.DeleteConnection(input)
}

func (a *CodeStarConnectionsActivities) DeleteHost(input *codestarconnections.DeleteHostInput) (*codestarconnections.DeleteHostOutput, error) {
	return a.client.DeleteHost(input)
}

func (a *CodeStarConnectionsActivities) GetConnection(input *codestarconnections.GetConnectionInput) (*codestarconnections.GetConnectionOutput, error) {
	return a.client.GetConnection(input)
}

func (a *CodeStarConnectionsActivities) GetHost(input *codestarconnections.GetHostInput) (*codestarconnections.GetHostOutput, error) {
	return a.client.GetHost(input)
}

func (a *CodeStarConnectionsActivities) ListConnections(input *codestarconnections.ListConnectionsInput) (*codestarconnections.ListConnectionsOutput, error) {
	return a.client.ListConnections(input)
}

func (a *CodeStarConnectionsActivities) ListHosts(input *codestarconnections.ListHostsInput) (*codestarconnections.ListHostsOutput, error) {
	return a.client.ListHosts(input)
}

func (a *CodeStarConnectionsActivities) ListTagsForResource(input *codestarconnections.ListTagsForResourceInput) (*codestarconnections.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResource(input)
}

func (a *CodeStarConnectionsActivities) TagResource(input *codestarconnections.TagResourceInput) (*codestarconnections.TagResourceOutput, error) {
	return a.client.TagResource(input)
}

func (a *CodeStarConnectionsActivities) UntagResource(input *codestarconnections.UntagResourceInput) (*codestarconnections.UntagResourceOutput, error) {
	return a.client.UntagResource(input)
}
