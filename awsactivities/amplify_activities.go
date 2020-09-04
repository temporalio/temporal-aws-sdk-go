
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/amplify"
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"
)

type AmplifyActivities struct {
	client amplifyiface.AmplifyAPI
}

func NewAmplifyActivities(client amplifyiface.AmplifyAPI) *AmplifyActivities {
    return &AmplifyActivities{client: client}
}

func (a *AmplifyActivities) CreateApp(input *amplify.CreateAppInput) (*amplify.CreateAppOutput, error) {
    return a.client.CreateApp(input)
}

func (a *AmplifyActivities) CreateBackendEnvironment(input *amplify.CreateBackendEnvironmentInput) (*amplify.CreateBackendEnvironmentOutput, error) {
    return a.client.CreateBackendEnvironment(input)
}

func (a *AmplifyActivities) CreateBranch(input *amplify.CreateBranchInput) (*amplify.CreateBranchOutput, error) {
    return a.client.CreateBranch(input)
}

func (a *AmplifyActivities) CreateDeployment(input *amplify.CreateDeploymentInput) (*amplify.CreateDeploymentOutput, error) {
    return a.client.CreateDeployment(input)
}

func (a *AmplifyActivities) CreateDomainAssociation(input *amplify.CreateDomainAssociationInput) (*amplify.CreateDomainAssociationOutput, error) {
    return a.client.CreateDomainAssociation(input)
}

func (a *AmplifyActivities) CreateWebhook(input *amplify.CreateWebhookInput) (*amplify.CreateWebhookOutput, error) {
    return a.client.CreateWebhook(input)
}

func (a *AmplifyActivities) DeleteApp(input *amplify.DeleteAppInput) (*amplify.DeleteAppOutput, error) {
    return a.client.DeleteApp(input)
}

func (a *AmplifyActivities) DeleteBackendEnvironment(input *amplify.DeleteBackendEnvironmentInput) (*amplify.DeleteBackendEnvironmentOutput, error) {
    return a.client.DeleteBackendEnvironment(input)
}

func (a *AmplifyActivities) DeleteBranch(input *amplify.DeleteBranchInput) (*amplify.DeleteBranchOutput, error) {
    return a.client.DeleteBranch(input)
}

func (a *AmplifyActivities) DeleteDomainAssociation(input *amplify.DeleteDomainAssociationInput) (*amplify.DeleteDomainAssociationOutput, error) {
    return a.client.DeleteDomainAssociation(input)
}

func (a *AmplifyActivities) DeleteJob(input *amplify.DeleteJobInput) (*amplify.DeleteJobOutput, error) {
    return a.client.DeleteJob(input)
}

func (a *AmplifyActivities) DeleteWebhook(input *amplify.DeleteWebhookInput) (*amplify.DeleteWebhookOutput, error) {
    return a.client.DeleteWebhook(input)
}

func (a *AmplifyActivities) GenerateAccessLogs(input *amplify.GenerateAccessLogsInput) (*amplify.GenerateAccessLogsOutput, error) {
    return a.client.GenerateAccessLogs(input)
}

func (a *AmplifyActivities) GetApp(input *amplify.GetAppInput) (*amplify.GetAppOutput, error) {
    return a.client.GetApp(input)
}

func (a *AmplifyActivities) GetArtifactUrl(input *amplify.GetArtifactUrlInput) (*amplify.GetArtifactUrlOutput, error) {
    return a.client.GetArtifactUrl(input)
}

func (a *AmplifyActivities) GetBackendEnvironment(input *amplify.GetBackendEnvironmentInput) (*amplify.GetBackendEnvironmentOutput, error) {
    return a.client.GetBackendEnvironment(input)
}

func (a *AmplifyActivities) GetBranch(input *amplify.GetBranchInput) (*amplify.GetBranchOutput, error) {
    return a.client.GetBranch(input)
}

func (a *AmplifyActivities) GetDomainAssociation(input *amplify.GetDomainAssociationInput) (*amplify.GetDomainAssociationOutput, error) {
    return a.client.GetDomainAssociation(input)
}

func (a *AmplifyActivities) GetJob(input *amplify.GetJobInput) (*amplify.GetJobOutput, error) {
    return a.client.GetJob(input)
}

func (a *AmplifyActivities) GetWebhook(input *amplify.GetWebhookInput) (*amplify.GetWebhookOutput, error) {
    return a.client.GetWebhook(input)
}

func (a *AmplifyActivities) ListApps(input *amplify.ListAppsInput) (*amplify.ListAppsOutput, error) {
    return a.client.ListApps(input)
}

func (a *AmplifyActivities) ListArtifacts(input *amplify.ListArtifactsInput) (*amplify.ListArtifactsOutput, error) {
    return a.client.ListArtifacts(input)
}

func (a *AmplifyActivities) ListBackendEnvironments(input *amplify.ListBackendEnvironmentsInput) (*amplify.ListBackendEnvironmentsOutput, error) {
    return a.client.ListBackendEnvironments(input)
}

func (a *AmplifyActivities) ListBranches(input *amplify.ListBranchesInput) (*amplify.ListBranchesOutput, error) {
    return a.client.ListBranches(input)
}

func (a *AmplifyActivities) ListDomainAssociations(input *amplify.ListDomainAssociationsInput) (*amplify.ListDomainAssociationsOutput, error) {
    return a.client.ListDomainAssociations(input)
}

func (a *AmplifyActivities) ListJobs(input *amplify.ListJobsInput) (*amplify.ListJobsOutput, error) {
    return a.client.ListJobs(input)
}

func (a *AmplifyActivities) ListTagsForResource(input *amplify.ListTagsForResourceInput) (*amplify.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *AmplifyActivities) ListWebhooks(input *amplify.ListWebhooksInput) (*amplify.ListWebhooksOutput, error) {
    return a.client.ListWebhooks(input)
}

func (a *AmplifyActivities) StartDeployment(input *amplify.StartDeploymentInput) (*amplify.StartDeploymentOutput, error) {
    return a.client.StartDeployment(input)
}

func (a *AmplifyActivities) StartJob(input *amplify.StartJobInput) (*amplify.StartJobOutput, error) {
    return a.client.StartJob(input)
}

func (a *AmplifyActivities) StopJob(input *amplify.StopJobInput) (*amplify.StopJobOutput, error) {
    return a.client.StopJob(input)
}

func (a *AmplifyActivities) TagResource(input *amplify.TagResourceInput) (*amplify.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *AmplifyActivities) UntagResource(input *amplify.UntagResourceInput) (*amplify.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *AmplifyActivities) UpdateApp(input *amplify.UpdateAppInput) (*amplify.UpdateAppOutput, error) {
    return a.client.UpdateApp(input)
}

func (a *AmplifyActivities) UpdateBranch(input *amplify.UpdateBranchInput) (*amplify.UpdateBranchOutput, error) {
    return a.client.UpdateBranch(input)
}

func (a *AmplifyActivities) UpdateDomainAssociation(input *amplify.UpdateDomainAssociationInput) (*amplify.UpdateDomainAssociationOutput, error) {
    return a.client.UpdateDomainAssociation(input)
}

func (a *AmplifyActivities) UpdateWebhook(input *amplify.UpdateWebhookInput) (*amplify.UpdateWebhookOutput, error) {
    return a.client.UpdateWebhook(input)
}
