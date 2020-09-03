package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/codedeploy"
	"github.com/aws/aws-sdk-go/service/codedeploy/codedeployiface"
)

type CodeDeployActivities struct {
	client codedeployiface.CodeDeployAPI
}

func NewCodeDeployActivities(client codedeployiface.CodeDeployAPI) *CodeDeployActivities {
    return &CodeDeployActivities{client: client}
}

func (a *CodeDeployActivities) AddTagsToOnPremisesInstances(input *codedeploy.AddTagsToOnPremisesInstancesInput) (*codedeploy.AddTagsToOnPremisesInstancesOutput, error) {
    return a.client.AddTagsToOnPremisesInstances(input)
}

func (a *CodeDeployActivities) BatchGetApplicationRevisions(input *codedeploy.BatchGetApplicationRevisionsInput) (*codedeploy.BatchGetApplicationRevisionsOutput, error) {
    return a.client.BatchGetApplicationRevisions(input)
}

func (a *CodeDeployActivities) BatchGetApplications(input *codedeploy.BatchGetApplicationsInput) (*codedeploy.BatchGetApplicationsOutput, error) {
    return a.client.BatchGetApplications(input)
}

func (a *CodeDeployActivities) BatchGetDeploymentGroups(input *codedeploy.BatchGetDeploymentGroupsInput) (*codedeploy.BatchGetDeploymentGroupsOutput, error) {
    return a.client.BatchGetDeploymentGroups(input)
}

func (a *CodeDeployActivities) BatchGetDeploymentInstances(input *codedeploy.BatchGetDeploymentInstancesInput) (*codedeploy.BatchGetDeploymentInstancesOutput, error) {
    return a.client.BatchGetDeploymentInstances(input)
}

func (a *CodeDeployActivities) BatchGetDeploymentTargets(input *codedeploy.BatchGetDeploymentTargetsInput) (*codedeploy.BatchGetDeploymentTargetsOutput, error) {
    return a.client.BatchGetDeploymentTargets(input)
}

func (a *CodeDeployActivities) BatchGetDeployments(input *codedeploy.BatchGetDeploymentsInput) (*codedeploy.BatchGetDeploymentsOutput, error) {
    return a.client.BatchGetDeployments(input)
}

func (a *CodeDeployActivities) BatchGetOnPremisesInstances(input *codedeploy.BatchGetOnPremisesInstancesInput) (*codedeploy.BatchGetOnPremisesInstancesOutput, error) {
    return a.client.BatchGetOnPremisesInstances(input)
}

func (a *CodeDeployActivities) ContinueDeployment(input *codedeploy.ContinueDeploymentInput) (*codedeploy.ContinueDeploymentOutput, error) {
    return a.client.ContinueDeployment(input)
}

func (a *CodeDeployActivities) CreateApplication(input *codedeploy.CreateApplicationInput) (*codedeploy.CreateApplicationOutput, error) {
    return a.client.CreateApplication(input)
}

func (a *CodeDeployActivities) CreateDeployment(input *codedeploy.CreateDeploymentInput) (*codedeploy.CreateDeploymentOutput, error) {
    return a.client.CreateDeployment(input)
}

func (a *CodeDeployActivities) CreateDeploymentConfig(input *codedeploy.CreateDeploymentConfigInput) (*codedeploy.CreateDeploymentConfigOutput, error) {
    return a.client.CreateDeploymentConfig(input)
}

func (a *CodeDeployActivities) CreateDeploymentGroup(input *codedeploy.CreateDeploymentGroupInput) (*codedeploy.CreateDeploymentGroupOutput, error) {
    return a.client.CreateDeploymentGroup(input)
}

func (a *CodeDeployActivities) DeleteApplication(input *codedeploy.DeleteApplicationInput) (*codedeploy.DeleteApplicationOutput, error) {
    return a.client.DeleteApplication(input)
}

func (a *CodeDeployActivities) DeleteDeploymentConfig(input *codedeploy.DeleteDeploymentConfigInput) (*codedeploy.DeleteDeploymentConfigOutput, error) {
    return a.client.DeleteDeploymentConfig(input)
}

func (a *CodeDeployActivities) DeleteDeploymentGroup(input *codedeploy.DeleteDeploymentGroupInput) (*codedeploy.DeleteDeploymentGroupOutput, error) {
    return a.client.DeleteDeploymentGroup(input)
}

func (a *CodeDeployActivities) DeleteGitHubAccountToken(input *codedeploy.DeleteGitHubAccountTokenInput) (*codedeploy.DeleteGitHubAccountTokenOutput, error) {
    return a.client.DeleteGitHubAccountToken(input)
}

func (a *CodeDeployActivities) DeleteResourcesByExternalId(input *codedeploy.DeleteResourcesByExternalIdInput) (*codedeploy.DeleteResourcesByExternalIdOutput, error) {
    return a.client.DeleteResourcesByExternalId(input)
}

func (a *CodeDeployActivities) DeregisterOnPremisesInstance(input *codedeploy.DeregisterOnPremisesInstanceInput) (*codedeploy.DeregisterOnPremisesInstanceOutput, error) {
    return a.client.DeregisterOnPremisesInstance(input)
}

func (a *CodeDeployActivities) GetApplication(input *codedeploy.GetApplicationInput) (*codedeploy.GetApplicationOutput, error) {
    return a.client.GetApplication(input)
}

func (a *CodeDeployActivities) GetApplicationRevision(input *codedeploy.GetApplicationRevisionInput) (*codedeploy.GetApplicationRevisionOutput, error) {
    return a.client.GetApplicationRevision(input)
}

func (a *CodeDeployActivities) GetDeployment(input *codedeploy.GetDeploymentInput) (*codedeploy.GetDeploymentOutput, error) {
    return a.client.GetDeployment(input)
}

func (a *CodeDeployActivities) GetDeploymentConfig(input *codedeploy.GetDeploymentConfigInput) (*codedeploy.GetDeploymentConfigOutput, error) {
    return a.client.GetDeploymentConfig(input)
}

func (a *CodeDeployActivities) GetDeploymentGroup(input *codedeploy.GetDeploymentGroupInput) (*codedeploy.GetDeploymentGroupOutput, error) {
    return a.client.GetDeploymentGroup(input)
}

func (a *CodeDeployActivities) GetDeploymentInstance(input *codedeploy.GetDeploymentInstanceInput) (*codedeploy.GetDeploymentInstanceOutput, error) {
    return a.client.GetDeploymentInstance(input)
}

func (a *CodeDeployActivities) GetDeploymentTarget(input *codedeploy.GetDeploymentTargetInput) (*codedeploy.GetDeploymentTargetOutput, error) {
    return a.client.GetDeploymentTarget(input)
}

func (a *CodeDeployActivities) GetOnPremisesInstance(input *codedeploy.GetOnPremisesInstanceInput) (*codedeploy.GetOnPremisesInstanceOutput, error) {
    return a.client.GetOnPremisesInstance(input)
}

func (a *CodeDeployActivities) ListApplicationRevisions(input *codedeploy.ListApplicationRevisionsInput) (*codedeploy.ListApplicationRevisionsOutput, error) {
    return a.client.ListApplicationRevisions(input)
}

func (a *CodeDeployActivities) ListApplications(input *codedeploy.ListApplicationsInput) (*codedeploy.ListApplicationsOutput, error) {
    return a.client.ListApplications(input)
}

func (a *CodeDeployActivities) ListDeploymentConfigs(input *codedeploy.ListDeploymentConfigsInput) (*codedeploy.ListDeploymentConfigsOutput, error) {
    return a.client.ListDeploymentConfigs(input)
}

func (a *CodeDeployActivities) ListDeploymentGroups(input *codedeploy.ListDeploymentGroupsInput) (*codedeploy.ListDeploymentGroupsOutput, error) {
    return a.client.ListDeploymentGroups(input)
}

func (a *CodeDeployActivities) ListDeploymentInstances(input *codedeploy.ListDeploymentInstancesInput) (*codedeploy.ListDeploymentInstancesOutput, error) {
    return a.client.ListDeploymentInstances(input)
}

func (a *CodeDeployActivities) ListDeploymentTargets(input *codedeploy.ListDeploymentTargetsInput) (*codedeploy.ListDeploymentTargetsOutput, error) {
    return a.client.ListDeploymentTargets(input)
}

func (a *CodeDeployActivities) ListDeployments(input *codedeploy.ListDeploymentsInput) (*codedeploy.ListDeploymentsOutput, error) {
    return a.client.ListDeployments(input)
}

func (a *CodeDeployActivities) ListGitHubAccountTokenNames(input *codedeploy.ListGitHubAccountTokenNamesInput) (*codedeploy.ListGitHubAccountTokenNamesOutput, error) {
    return a.client.ListGitHubAccountTokenNames(input)
}

func (a *CodeDeployActivities) ListOnPremisesInstances(input *codedeploy.ListOnPremisesInstancesInput) (*codedeploy.ListOnPremisesInstancesOutput, error) {
    return a.client.ListOnPremisesInstances(input)
}

func (a *CodeDeployActivities) ListTagsForResource(input *codedeploy.ListTagsForResourceInput) (*codedeploy.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *CodeDeployActivities) PutLifecycleEventHookExecutionStatus(input *codedeploy.PutLifecycleEventHookExecutionStatusInput) (*codedeploy.PutLifecycleEventHookExecutionStatusOutput, error) {
    return a.client.PutLifecycleEventHookExecutionStatus(input)
}

func (a *CodeDeployActivities) RegisterApplicationRevision(input *codedeploy.RegisterApplicationRevisionInput) (*codedeploy.RegisterApplicationRevisionOutput, error) {
    return a.client.RegisterApplicationRevision(input)
}

func (a *CodeDeployActivities) RegisterOnPremisesInstance(input *codedeploy.RegisterOnPremisesInstanceInput) (*codedeploy.RegisterOnPremisesInstanceOutput, error) {
    return a.client.RegisterOnPremisesInstance(input)
}

func (a *CodeDeployActivities) RemoveTagsFromOnPremisesInstances(input *codedeploy.RemoveTagsFromOnPremisesInstancesInput) (*codedeploy.RemoveTagsFromOnPremisesInstancesOutput, error) {
    return a.client.RemoveTagsFromOnPremisesInstances(input)
}

func (a *CodeDeployActivities) SkipWaitTimeForInstanceTermination(input *codedeploy.SkipWaitTimeForInstanceTerminationInput) (*codedeploy.SkipWaitTimeForInstanceTerminationOutput, error) {
    return a.client.SkipWaitTimeForInstanceTermination(input)
}

func (a *CodeDeployActivities) StopDeployment(input *codedeploy.StopDeploymentInput) (*codedeploy.StopDeploymentOutput, error) {
    return a.client.StopDeployment(input)
}

func (a *CodeDeployActivities) TagResource(input *codedeploy.TagResourceInput) (*codedeploy.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *CodeDeployActivities) UntagResource(input *codedeploy.UntagResourceInput) (*codedeploy.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *CodeDeployActivities) UpdateApplication(input *codedeploy.UpdateApplicationInput) (*codedeploy.UpdateApplicationOutput, error) {
    return a.client.UpdateApplication(input)
}

func (a *CodeDeployActivities) UpdateDeploymentGroup(input *codedeploy.UpdateDeploymentGroupInput) (*codedeploy.UpdateDeploymentGroupOutput, error) {
    return a.client.UpdateDeploymentGroup(input)
}

func (a *CodeDeployActivities) WaitUntilDeploymentSuccessful(input *codedeploy.WaitUntilDeploymentSuccessfulInput) (*codedeploy.WaitUntilDeploymentSuccessfulOutput, error) {
    return a.client.WaitUntilDeploymentSuccessful(input)
}
