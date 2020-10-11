// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package migrationhubstub

import (
	"github.com/aws/aws-sdk-go/service/migrationhub"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	AssociateCreatedArtifact(ctx workflow.Context, input *migrationhub.AssociateCreatedArtifactInput) (*migrationhub.AssociateCreatedArtifactOutput, error)
	AssociateCreatedArtifactAsync(ctx workflow.Context, input *migrationhub.AssociateCreatedArtifactInput) *MigrationHubAssociateCreatedArtifactFuture

	AssociateDiscoveredResource(ctx workflow.Context, input *migrationhub.AssociateDiscoveredResourceInput) (*migrationhub.AssociateDiscoveredResourceOutput, error)
	AssociateDiscoveredResourceAsync(ctx workflow.Context, input *migrationhub.AssociateDiscoveredResourceInput) *MigrationHubAssociateDiscoveredResourceFuture

	CreateProgressUpdateStream(ctx workflow.Context, input *migrationhub.CreateProgressUpdateStreamInput) (*migrationhub.CreateProgressUpdateStreamOutput, error)
	CreateProgressUpdateStreamAsync(ctx workflow.Context, input *migrationhub.CreateProgressUpdateStreamInput) *MigrationHubCreateProgressUpdateStreamFuture

	DeleteProgressUpdateStream(ctx workflow.Context, input *migrationhub.DeleteProgressUpdateStreamInput) (*migrationhub.DeleteProgressUpdateStreamOutput, error)
	DeleteProgressUpdateStreamAsync(ctx workflow.Context, input *migrationhub.DeleteProgressUpdateStreamInput) *MigrationHubDeleteProgressUpdateStreamFuture

	DescribeApplicationState(ctx workflow.Context, input *migrationhub.DescribeApplicationStateInput) (*migrationhub.DescribeApplicationStateOutput, error)
	DescribeApplicationStateAsync(ctx workflow.Context, input *migrationhub.DescribeApplicationStateInput) *MigrationHubDescribeApplicationStateFuture

	DescribeMigrationTask(ctx workflow.Context, input *migrationhub.DescribeMigrationTaskInput) (*migrationhub.DescribeMigrationTaskOutput, error)
	DescribeMigrationTaskAsync(ctx workflow.Context, input *migrationhub.DescribeMigrationTaskInput) *MigrationHubDescribeMigrationTaskFuture

	DisassociateCreatedArtifact(ctx workflow.Context, input *migrationhub.DisassociateCreatedArtifactInput) (*migrationhub.DisassociateCreatedArtifactOutput, error)
	DisassociateCreatedArtifactAsync(ctx workflow.Context, input *migrationhub.DisassociateCreatedArtifactInput) *MigrationHubDisassociateCreatedArtifactFuture

	DisassociateDiscoveredResource(ctx workflow.Context, input *migrationhub.DisassociateDiscoveredResourceInput) (*migrationhub.DisassociateDiscoveredResourceOutput, error)
	DisassociateDiscoveredResourceAsync(ctx workflow.Context, input *migrationhub.DisassociateDiscoveredResourceInput) *MigrationHubDisassociateDiscoveredResourceFuture

	ImportMigrationTask(ctx workflow.Context, input *migrationhub.ImportMigrationTaskInput) (*migrationhub.ImportMigrationTaskOutput, error)
	ImportMigrationTaskAsync(ctx workflow.Context, input *migrationhub.ImportMigrationTaskInput) *MigrationHubImportMigrationTaskFuture

	ListApplicationStates(ctx workflow.Context, input *migrationhub.ListApplicationStatesInput) (*migrationhub.ListApplicationStatesOutput, error)
	ListApplicationStatesAsync(ctx workflow.Context, input *migrationhub.ListApplicationStatesInput) *MigrationHubListApplicationStatesFuture

	ListCreatedArtifacts(ctx workflow.Context, input *migrationhub.ListCreatedArtifactsInput) (*migrationhub.ListCreatedArtifactsOutput, error)
	ListCreatedArtifactsAsync(ctx workflow.Context, input *migrationhub.ListCreatedArtifactsInput) *MigrationHubListCreatedArtifactsFuture

	ListDiscoveredResources(ctx workflow.Context, input *migrationhub.ListDiscoveredResourcesInput) (*migrationhub.ListDiscoveredResourcesOutput, error)
	ListDiscoveredResourcesAsync(ctx workflow.Context, input *migrationhub.ListDiscoveredResourcesInput) *MigrationHubListDiscoveredResourcesFuture

	ListMigrationTasks(ctx workflow.Context, input *migrationhub.ListMigrationTasksInput) (*migrationhub.ListMigrationTasksOutput, error)
	ListMigrationTasksAsync(ctx workflow.Context, input *migrationhub.ListMigrationTasksInput) *MigrationHubListMigrationTasksFuture

	ListProgressUpdateStreams(ctx workflow.Context, input *migrationhub.ListProgressUpdateStreamsInput) (*migrationhub.ListProgressUpdateStreamsOutput, error)
	ListProgressUpdateStreamsAsync(ctx workflow.Context, input *migrationhub.ListProgressUpdateStreamsInput) *MigrationHubListProgressUpdateStreamsFuture

	NotifyApplicationState(ctx workflow.Context, input *migrationhub.NotifyApplicationStateInput) (*migrationhub.NotifyApplicationStateOutput, error)
	NotifyApplicationStateAsync(ctx workflow.Context, input *migrationhub.NotifyApplicationStateInput) *MigrationHubNotifyApplicationStateFuture

	NotifyMigrationTaskState(ctx workflow.Context, input *migrationhub.NotifyMigrationTaskStateInput) (*migrationhub.NotifyMigrationTaskStateOutput, error)
	NotifyMigrationTaskStateAsync(ctx workflow.Context, input *migrationhub.NotifyMigrationTaskStateInput) *MigrationHubNotifyMigrationTaskStateFuture

	PutResourceAttributes(ctx workflow.Context, input *migrationhub.PutResourceAttributesInput) (*migrationhub.PutResourceAttributesOutput, error)
	PutResourceAttributesAsync(ctx workflow.Context, input *migrationhub.PutResourceAttributesInput) *MigrationHubPutResourceAttributesFuture
}

func NewClient() Client {
	return &stub{}
}
