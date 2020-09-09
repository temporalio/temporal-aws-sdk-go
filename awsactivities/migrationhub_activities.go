package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/migrationhub"
	"github.com/aws/aws-sdk-go/service/migrationhub/migrationhubiface"
)

type MigrationHubActivities struct {
	client migrationhubiface.MigrationHubAPI
}

func NewMigrationHubActivities(session *session.Session, config ...*aws.Config) *MigrationHubActivities {
	client := migrationhub.New(session, config...)
	return &MigrationHubActivities{client: client}
}

func (a *MigrationHubActivities) AssociateCreatedArtifact(input *migrationhub.AssociateCreatedArtifactInput) (*migrationhub.AssociateCreatedArtifactOutput, error) {
	return a.client.AssociateCreatedArtifact(input)
}

func (a *MigrationHubActivities) AssociateDiscoveredResource(input *migrationhub.AssociateDiscoveredResourceInput) (*migrationhub.AssociateDiscoveredResourceOutput, error) {
	return a.client.AssociateDiscoveredResource(input)
}

func (a *MigrationHubActivities) CreateProgressUpdateStream(input *migrationhub.CreateProgressUpdateStreamInput) (*migrationhub.CreateProgressUpdateStreamOutput, error) {
	return a.client.CreateProgressUpdateStream(input)
}

func (a *MigrationHubActivities) DeleteProgressUpdateStream(input *migrationhub.DeleteProgressUpdateStreamInput) (*migrationhub.DeleteProgressUpdateStreamOutput, error) {
	return a.client.DeleteProgressUpdateStream(input)
}

func (a *MigrationHubActivities) DescribeApplicationState(input *migrationhub.DescribeApplicationStateInput) (*migrationhub.DescribeApplicationStateOutput, error) {
	return a.client.DescribeApplicationState(input)
}

func (a *MigrationHubActivities) DescribeMigrationTask(input *migrationhub.DescribeMigrationTaskInput) (*migrationhub.DescribeMigrationTaskOutput, error) {
	return a.client.DescribeMigrationTask(input)
}

func (a *MigrationHubActivities) DisassociateCreatedArtifact(input *migrationhub.DisassociateCreatedArtifactInput) (*migrationhub.DisassociateCreatedArtifactOutput, error) {
	return a.client.DisassociateCreatedArtifact(input)
}

func (a *MigrationHubActivities) DisassociateDiscoveredResource(input *migrationhub.DisassociateDiscoveredResourceInput) (*migrationhub.DisassociateDiscoveredResourceOutput, error) {
	return a.client.DisassociateDiscoveredResource(input)
}

func (a *MigrationHubActivities) ImportMigrationTask(input *migrationhub.ImportMigrationTaskInput) (*migrationhub.ImportMigrationTaskOutput, error) {
	return a.client.ImportMigrationTask(input)
}

func (a *MigrationHubActivities) ListApplicationStates(input *migrationhub.ListApplicationStatesInput) (*migrationhub.ListApplicationStatesOutput, error) {
	return a.client.ListApplicationStates(input)
}

func (a *MigrationHubActivities) ListCreatedArtifacts(input *migrationhub.ListCreatedArtifactsInput) (*migrationhub.ListCreatedArtifactsOutput, error) {
	return a.client.ListCreatedArtifacts(input)
}

func (a *MigrationHubActivities) ListDiscoveredResources(input *migrationhub.ListDiscoveredResourcesInput) (*migrationhub.ListDiscoveredResourcesOutput, error) {
	return a.client.ListDiscoveredResources(input)
}

func (a *MigrationHubActivities) ListMigrationTasks(input *migrationhub.ListMigrationTasksInput) (*migrationhub.ListMigrationTasksOutput, error) {
	return a.client.ListMigrationTasks(input)
}

func (a *MigrationHubActivities) ListProgressUpdateStreams(input *migrationhub.ListProgressUpdateStreamsInput) (*migrationhub.ListProgressUpdateStreamsOutput, error) {
	return a.client.ListProgressUpdateStreams(input)
}

func (a *MigrationHubActivities) NotifyApplicationState(input *migrationhub.NotifyApplicationStateInput) (*migrationhub.NotifyApplicationStateOutput, error) {
	return a.client.NotifyApplicationState(input)
}

func (a *MigrationHubActivities) NotifyMigrationTaskState(input *migrationhub.NotifyMigrationTaskStateInput) (*migrationhub.NotifyMigrationTaskStateOutput, error) {
	return a.client.NotifyMigrationTaskState(input)
}

func (a *MigrationHubActivities) PutResourceAttributes(input *migrationhub.PutResourceAttributesInput) (*migrationhub.PutResourceAttributesOutput, error) {
	return a.client.PutResourceAttributes(input)
}
