// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package amplifystub

import (
	"github.com/aws/aws-sdk-go/service/amplify"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateApp(ctx workflow.Context, input *amplify.CreateAppInput) (*amplify.CreateAppOutput, error)
	CreateAppAsync(ctx workflow.Context, input *amplify.CreateAppInput) *CreateAppFuture

	CreateBackendEnvironment(ctx workflow.Context, input *amplify.CreateBackendEnvironmentInput) (*amplify.CreateBackendEnvironmentOutput, error)
	CreateBackendEnvironmentAsync(ctx workflow.Context, input *amplify.CreateBackendEnvironmentInput) *CreateBackendEnvironmentFuture

	CreateBranch(ctx workflow.Context, input *amplify.CreateBranchInput) (*amplify.CreateBranchOutput, error)
	CreateBranchAsync(ctx workflow.Context, input *amplify.CreateBranchInput) *CreateBranchFuture

	CreateDeployment(ctx workflow.Context, input *amplify.CreateDeploymentInput) (*amplify.CreateDeploymentOutput, error)
	CreateDeploymentAsync(ctx workflow.Context, input *amplify.CreateDeploymentInput) *CreateDeploymentFuture

	CreateDomainAssociation(ctx workflow.Context, input *amplify.CreateDomainAssociationInput) (*amplify.CreateDomainAssociationOutput, error)
	CreateDomainAssociationAsync(ctx workflow.Context, input *amplify.CreateDomainAssociationInput) *CreateDomainAssociationFuture

	CreateWebhook(ctx workflow.Context, input *amplify.CreateWebhookInput) (*amplify.CreateWebhookOutput, error)
	CreateWebhookAsync(ctx workflow.Context, input *amplify.CreateWebhookInput) *CreateWebhookFuture

	DeleteApp(ctx workflow.Context, input *amplify.DeleteAppInput) (*amplify.DeleteAppOutput, error)
	DeleteAppAsync(ctx workflow.Context, input *amplify.DeleteAppInput) *DeleteAppFuture

	DeleteBackendEnvironment(ctx workflow.Context, input *amplify.DeleteBackendEnvironmentInput) (*amplify.DeleteBackendEnvironmentOutput, error)
	DeleteBackendEnvironmentAsync(ctx workflow.Context, input *amplify.DeleteBackendEnvironmentInput) *DeleteBackendEnvironmentFuture

	DeleteBranch(ctx workflow.Context, input *amplify.DeleteBranchInput) (*amplify.DeleteBranchOutput, error)
	DeleteBranchAsync(ctx workflow.Context, input *amplify.DeleteBranchInput) *DeleteBranchFuture

	DeleteDomainAssociation(ctx workflow.Context, input *amplify.DeleteDomainAssociationInput) (*amplify.DeleteDomainAssociationOutput, error)
	DeleteDomainAssociationAsync(ctx workflow.Context, input *amplify.DeleteDomainAssociationInput) *DeleteDomainAssociationFuture

	DeleteJob(ctx workflow.Context, input *amplify.DeleteJobInput) (*amplify.DeleteJobOutput, error)
	DeleteJobAsync(ctx workflow.Context, input *amplify.DeleteJobInput) *DeleteJobFuture

	DeleteWebhook(ctx workflow.Context, input *amplify.DeleteWebhookInput) (*amplify.DeleteWebhookOutput, error)
	DeleteWebhookAsync(ctx workflow.Context, input *amplify.DeleteWebhookInput) *DeleteWebhookFuture

	GenerateAccessLogs(ctx workflow.Context, input *amplify.GenerateAccessLogsInput) (*amplify.GenerateAccessLogsOutput, error)
	GenerateAccessLogsAsync(ctx workflow.Context, input *amplify.GenerateAccessLogsInput) *GenerateAccessLogsFuture

	GetApp(ctx workflow.Context, input *amplify.GetAppInput) (*amplify.GetAppOutput, error)
	GetAppAsync(ctx workflow.Context, input *amplify.GetAppInput) *GetAppFuture

	GetArtifactUrl(ctx workflow.Context, input *amplify.GetArtifactUrlInput) (*amplify.GetArtifactUrlOutput, error)
	GetArtifactUrlAsync(ctx workflow.Context, input *amplify.GetArtifactUrlInput) *GetArtifactUrlFuture

	GetBackendEnvironment(ctx workflow.Context, input *amplify.GetBackendEnvironmentInput) (*amplify.GetBackendEnvironmentOutput, error)
	GetBackendEnvironmentAsync(ctx workflow.Context, input *amplify.GetBackendEnvironmentInput) *GetBackendEnvironmentFuture

	GetBranch(ctx workflow.Context, input *amplify.GetBranchInput) (*amplify.GetBranchOutput, error)
	GetBranchAsync(ctx workflow.Context, input *amplify.GetBranchInput) *GetBranchFuture

	GetDomainAssociation(ctx workflow.Context, input *amplify.GetDomainAssociationInput) (*amplify.GetDomainAssociationOutput, error)
	GetDomainAssociationAsync(ctx workflow.Context, input *amplify.GetDomainAssociationInput) *GetDomainAssociationFuture

	GetJob(ctx workflow.Context, input *amplify.GetJobInput) (*amplify.GetJobOutput, error)
	GetJobAsync(ctx workflow.Context, input *amplify.GetJobInput) *GetJobFuture

	GetWebhook(ctx workflow.Context, input *amplify.GetWebhookInput) (*amplify.GetWebhookOutput, error)
	GetWebhookAsync(ctx workflow.Context, input *amplify.GetWebhookInput) *GetWebhookFuture

	ListApps(ctx workflow.Context, input *amplify.ListAppsInput) (*amplify.ListAppsOutput, error)
	ListAppsAsync(ctx workflow.Context, input *amplify.ListAppsInput) *ListAppsFuture

	ListArtifacts(ctx workflow.Context, input *amplify.ListArtifactsInput) (*amplify.ListArtifactsOutput, error)
	ListArtifactsAsync(ctx workflow.Context, input *amplify.ListArtifactsInput) *ListArtifactsFuture

	ListBackendEnvironments(ctx workflow.Context, input *amplify.ListBackendEnvironmentsInput) (*amplify.ListBackendEnvironmentsOutput, error)
	ListBackendEnvironmentsAsync(ctx workflow.Context, input *amplify.ListBackendEnvironmentsInput) *ListBackendEnvironmentsFuture

	ListBranches(ctx workflow.Context, input *amplify.ListBranchesInput) (*amplify.ListBranchesOutput, error)
	ListBranchesAsync(ctx workflow.Context, input *amplify.ListBranchesInput) *ListBranchesFuture

	ListDomainAssociations(ctx workflow.Context, input *amplify.ListDomainAssociationsInput) (*amplify.ListDomainAssociationsOutput, error)
	ListDomainAssociationsAsync(ctx workflow.Context, input *amplify.ListDomainAssociationsInput) *ListDomainAssociationsFuture

	ListJobs(ctx workflow.Context, input *amplify.ListJobsInput) (*amplify.ListJobsOutput, error)
	ListJobsAsync(ctx workflow.Context, input *amplify.ListJobsInput) *ListJobsFuture

	ListTagsForResource(ctx workflow.Context, input *amplify.ListTagsForResourceInput) (*amplify.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *amplify.ListTagsForResourceInput) *ListTagsForResourceFuture

	ListWebhooks(ctx workflow.Context, input *amplify.ListWebhooksInput) (*amplify.ListWebhooksOutput, error)
	ListWebhooksAsync(ctx workflow.Context, input *amplify.ListWebhooksInput) *ListWebhooksFuture

	StartDeployment(ctx workflow.Context, input *amplify.StartDeploymentInput) (*amplify.StartDeploymentOutput, error)
	StartDeploymentAsync(ctx workflow.Context, input *amplify.StartDeploymentInput) *StartDeploymentFuture

	StartJob(ctx workflow.Context, input *amplify.StartJobInput) (*amplify.StartJobOutput, error)
	StartJobAsync(ctx workflow.Context, input *amplify.StartJobInput) *StartJobFuture

	StopJob(ctx workflow.Context, input *amplify.StopJobInput) (*amplify.StopJobOutput, error)
	StopJobAsync(ctx workflow.Context, input *amplify.StopJobInput) *StopJobFuture

	TagResource(ctx workflow.Context, input *amplify.TagResourceInput) (*amplify.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *amplify.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *amplify.UntagResourceInput) (*amplify.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *amplify.UntagResourceInput) *UntagResourceFuture

	UpdateApp(ctx workflow.Context, input *amplify.UpdateAppInput) (*amplify.UpdateAppOutput, error)
	UpdateAppAsync(ctx workflow.Context, input *amplify.UpdateAppInput) *UpdateAppFuture

	UpdateBranch(ctx workflow.Context, input *amplify.UpdateBranchInput) (*amplify.UpdateBranchOutput, error)
	UpdateBranchAsync(ctx workflow.Context, input *amplify.UpdateBranchInput) *UpdateBranchFuture

	UpdateDomainAssociation(ctx workflow.Context, input *amplify.UpdateDomainAssociationInput) (*amplify.UpdateDomainAssociationOutput, error)
	UpdateDomainAssociationAsync(ctx workflow.Context, input *amplify.UpdateDomainAssociationInput) *UpdateDomainAssociationFuture

	UpdateWebhook(ctx workflow.Context, input *amplify.UpdateWebhookInput) (*amplify.UpdateWebhookOutput, error)
	UpdateWebhookAsync(ctx workflow.Context, input *amplify.UpdateWebhookInput) *UpdateWebhookFuture
}

func NewClient() Client {
	return &stub{}
}
