package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codedeploy"
	"github.com/aws/aws-sdk-go/service/codedeploy/codedeployiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type CodeDeployActivities struct {
	client codedeployiface.CodeDeployAPI
}

func NewCodeDeployActivities(session *session.Session, config ...*aws.Config) *CodeDeployActivities {
	client := codedeploy.New(session, config...)
	return &CodeDeployActivities{client: client}
}

func (a *CodeDeployActivities) AddTagsToOnPremisesInstances(ctx context.Context, input *codedeploy.AddTagsToOnPremisesInstancesInput) (*codedeploy.AddTagsToOnPremisesInstancesOutput, error) {
	return a.client.AddTagsToOnPremisesInstancesWithContext(ctx, input)
}

func (a *CodeDeployActivities) BatchGetApplicationRevisions(ctx context.Context, input *codedeploy.BatchGetApplicationRevisionsInput) (*codedeploy.BatchGetApplicationRevisionsOutput, error) {
	return a.client.BatchGetApplicationRevisionsWithContext(ctx, input)
}

func (a *CodeDeployActivities) BatchGetApplications(ctx context.Context, input *codedeploy.BatchGetApplicationsInput) (*codedeploy.BatchGetApplicationsOutput, error) {
	return a.client.BatchGetApplicationsWithContext(ctx, input)
}

func (a *CodeDeployActivities) BatchGetDeploymentGroups(ctx context.Context, input *codedeploy.BatchGetDeploymentGroupsInput) (*codedeploy.BatchGetDeploymentGroupsOutput, error) {
	return a.client.BatchGetDeploymentGroupsWithContext(ctx, input)
}

func (a *CodeDeployActivities) BatchGetDeploymentInstances(ctx context.Context, input *codedeploy.BatchGetDeploymentInstancesInput) (*codedeploy.BatchGetDeploymentInstancesOutput, error) {
	return a.client.BatchGetDeploymentInstancesWithContext(ctx, input)
}

func (a *CodeDeployActivities) BatchGetDeploymentTargets(ctx context.Context, input *codedeploy.BatchGetDeploymentTargetsInput) (*codedeploy.BatchGetDeploymentTargetsOutput, error) {
	return a.client.BatchGetDeploymentTargetsWithContext(ctx, input)
}

func (a *CodeDeployActivities) BatchGetDeployments(ctx context.Context, input *codedeploy.BatchGetDeploymentsInput) (*codedeploy.BatchGetDeploymentsOutput, error) {
	return a.client.BatchGetDeploymentsWithContext(ctx, input)
}

func (a *CodeDeployActivities) BatchGetOnPremisesInstances(ctx context.Context, input *codedeploy.BatchGetOnPremisesInstancesInput) (*codedeploy.BatchGetOnPremisesInstancesOutput, error) {
	return a.client.BatchGetOnPremisesInstancesWithContext(ctx, input)
}

func (a *CodeDeployActivities) ContinueDeployment(ctx context.Context, input *codedeploy.ContinueDeploymentInput) (*codedeploy.ContinueDeploymentOutput, error) {
	return a.client.ContinueDeploymentWithContext(ctx, input)
}

func (a *CodeDeployActivities) CreateApplication(ctx context.Context, input *codedeploy.CreateApplicationInput) (*codedeploy.CreateApplicationOutput, error) {
	return a.client.CreateApplicationWithContext(ctx, input)
}

func (a *CodeDeployActivities) CreateDeployment(ctx context.Context, input *codedeploy.CreateDeploymentInput) (*codedeploy.CreateDeploymentOutput, error) {
	return a.client.CreateDeploymentWithContext(ctx, input)
}

func (a *CodeDeployActivities) CreateDeploymentConfig(ctx context.Context, input *codedeploy.CreateDeploymentConfigInput) (*codedeploy.CreateDeploymentConfigOutput, error) {
	return a.client.CreateDeploymentConfigWithContext(ctx, input)
}

func (a *CodeDeployActivities) CreateDeploymentGroup(ctx context.Context, input *codedeploy.CreateDeploymentGroupInput) (*codedeploy.CreateDeploymentGroupOutput, error) {
	return a.client.CreateDeploymentGroupWithContext(ctx, input)
}

func (a *CodeDeployActivities) DeleteApplication(ctx context.Context, input *codedeploy.DeleteApplicationInput) (*codedeploy.DeleteApplicationOutput, error) {
	return a.client.DeleteApplicationWithContext(ctx, input)
}

func (a *CodeDeployActivities) DeleteDeploymentConfig(ctx context.Context, input *codedeploy.DeleteDeploymentConfigInput) (*codedeploy.DeleteDeploymentConfigOutput, error) {
	return a.client.DeleteDeploymentConfigWithContext(ctx, input)
}

func (a *CodeDeployActivities) DeleteDeploymentGroup(ctx context.Context, input *codedeploy.DeleteDeploymentGroupInput) (*codedeploy.DeleteDeploymentGroupOutput, error) {
	return a.client.DeleteDeploymentGroupWithContext(ctx, input)
}

func (a *CodeDeployActivities) DeleteGitHubAccountToken(ctx context.Context, input *codedeploy.DeleteGitHubAccountTokenInput) (*codedeploy.DeleteGitHubAccountTokenOutput, error) {
	return a.client.DeleteGitHubAccountTokenWithContext(ctx, input)
}

func (a *CodeDeployActivities) DeleteResourcesByExternalId(ctx context.Context, input *codedeploy.DeleteResourcesByExternalIdInput) (*codedeploy.DeleteResourcesByExternalIdOutput, error) {
	return a.client.DeleteResourcesByExternalIdWithContext(ctx, input)
}

func (a *CodeDeployActivities) DeregisterOnPremisesInstance(ctx context.Context, input *codedeploy.DeregisterOnPremisesInstanceInput) (*codedeploy.DeregisterOnPremisesInstanceOutput, error) {
	return a.client.DeregisterOnPremisesInstanceWithContext(ctx, input)
}

func (a *CodeDeployActivities) GetApplication(ctx context.Context, input *codedeploy.GetApplicationInput) (*codedeploy.GetApplicationOutput, error) {
	return a.client.GetApplicationWithContext(ctx, input)
}

func (a *CodeDeployActivities) GetApplicationRevision(ctx context.Context, input *codedeploy.GetApplicationRevisionInput) (*codedeploy.GetApplicationRevisionOutput, error) {
	return a.client.GetApplicationRevisionWithContext(ctx, input)
}

func (a *CodeDeployActivities) GetDeployment(ctx context.Context, input *codedeploy.GetDeploymentInput) (*codedeploy.GetDeploymentOutput, error) {
	return a.client.GetDeploymentWithContext(ctx, input)
}

func (a *CodeDeployActivities) GetDeploymentConfig(ctx context.Context, input *codedeploy.GetDeploymentConfigInput) (*codedeploy.GetDeploymentConfigOutput, error) {
	return a.client.GetDeploymentConfigWithContext(ctx, input)
}

func (a *CodeDeployActivities) GetDeploymentGroup(ctx context.Context, input *codedeploy.GetDeploymentGroupInput) (*codedeploy.GetDeploymentGroupOutput, error) {
	return a.client.GetDeploymentGroupWithContext(ctx, input)
}

func (a *CodeDeployActivities) GetDeploymentInstance(ctx context.Context, input *codedeploy.GetDeploymentInstanceInput) (*codedeploy.GetDeploymentInstanceOutput, error) {
	return a.client.GetDeploymentInstanceWithContext(ctx, input)
}

func (a *CodeDeployActivities) GetDeploymentTarget(ctx context.Context, input *codedeploy.GetDeploymentTargetInput) (*codedeploy.GetDeploymentTargetOutput, error) {
	return a.client.GetDeploymentTargetWithContext(ctx, input)
}

func (a *CodeDeployActivities) GetOnPremisesInstance(ctx context.Context, input *codedeploy.GetOnPremisesInstanceInput) (*codedeploy.GetOnPremisesInstanceOutput, error) {
	return a.client.GetOnPremisesInstanceWithContext(ctx, input)
}

func (a *CodeDeployActivities) ListApplicationRevisions(ctx context.Context, input *codedeploy.ListApplicationRevisionsInput) (*codedeploy.ListApplicationRevisionsOutput, error) {
	return a.client.ListApplicationRevisionsWithContext(ctx, input)
}

func (a *CodeDeployActivities) ListApplications(ctx context.Context, input *codedeploy.ListApplicationsInput) (*codedeploy.ListApplicationsOutput, error) {
	return a.client.ListApplicationsWithContext(ctx, input)
}

func (a *CodeDeployActivities) ListDeploymentConfigs(ctx context.Context, input *codedeploy.ListDeploymentConfigsInput) (*codedeploy.ListDeploymentConfigsOutput, error) {
	return a.client.ListDeploymentConfigsWithContext(ctx, input)
}

func (a *CodeDeployActivities) ListDeploymentGroups(ctx context.Context, input *codedeploy.ListDeploymentGroupsInput) (*codedeploy.ListDeploymentGroupsOutput, error) {
	return a.client.ListDeploymentGroupsWithContext(ctx, input)
}

func (a *CodeDeployActivities) ListDeploymentInstances(ctx context.Context, input *codedeploy.ListDeploymentInstancesInput) (*codedeploy.ListDeploymentInstancesOutput, error) {
	return a.client.ListDeploymentInstancesWithContext(ctx, input)
}

func (a *CodeDeployActivities) ListDeploymentTargets(ctx context.Context, input *codedeploy.ListDeploymentTargetsInput) (*codedeploy.ListDeploymentTargetsOutput, error) {
	return a.client.ListDeploymentTargetsWithContext(ctx, input)
}

func (a *CodeDeployActivities) ListDeployments(ctx context.Context, input *codedeploy.ListDeploymentsInput) (*codedeploy.ListDeploymentsOutput, error) {
	return a.client.ListDeploymentsWithContext(ctx, input)
}

func (a *CodeDeployActivities) ListGitHubAccountTokenNames(ctx context.Context, input *codedeploy.ListGitHubAccountTokenNamesInput) (*codedeploy.ListGitHubAccountTokenNamesOutput, error) {
	return a.client.ListGitHubAccountTokenNamesWithContext(ctx, input)
}

func (a *CodeDeployActivities) ListOnPremisesInstances(ctx context.Context, input *codedeploy.ListOnPremisesInstancesInput) (*codedeploy.ListOnPremisesInstancesOutput, error) {
	return a.client.ListOnPremisesInstancesWithContext(ctx, input)
}

func (a *CodeDeployActivities) ListTagsForResource(ctx context.Context, input *codedeploy.ListTagsForResourceInput) (*codedeploy.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *CodeDeployActivities) PutLifecycleEventHookExecutionStatus(ctx context.Context, input *codedeploy.PutLifecycleEventHookExecutionStatusInput) (*codedeploy.PutLifecycleEventHookExecutionStatusOutput, error) {
	return a.client.PutLifecycleEventHookExecutionStatusWithContext(ctx, input)
}

func (a *CodeDeployActivities) RegisterApplicationRevision(ctx context.Context, input *codedeploy.RegisterApplicationRevisionInput) (*codedeploy.RegisterApplicationRevisionOutput, error) {
	return a.client.RegisterApplicationRevisionWithContext(ctx, input)
}

func (a *CodeDeployActivities) RegisterOnPremisesInstance(ctx context.Context, input *codedeploy.RegisterOnPremisesInstanceInput) (*codedeploy.RegisterOnPremisesInstanceOutput, error) {
	return a.client.RegisterOnPremisesInstanceWithContext(ctx, input)
}

func (a *CodeDeployActivities) RemoveTagsFromOnPremisesInstances(ctx context.Context, input *codedeploy.RemoveTagsFromOnPremisesInstancesInput) (*codedeploy.RemoveTagsFromOnPremisesInstancesOutput, error) {
	return a.client.RemoveTagsFromOnPremisesInstancesWithContext(ctx, input)
}

func (a *CodeDeployActivities) SkipWaitTimeForInstanceTermination(ctx context.Context, input *codedeploy.SkipWaitTimeForInstanceTerminationInput) (*codedeploy.SkipWaitTimeForInstanceTerminationOutput, error) {
	return a.client.SkipWaitTimeForInstanceTerminationWithContext(ctx, input)
}

func (a *CodeDeployActivities) StopDeployment(ctx context.Context, input *codedeploy.StopDeploymentInput) (*codedeploy.StopDeploymentOutput, error) {
	return a.client.StopDeploymentWithContext(ctx, input)
}

func (a *CodeDeployActivities) TagResource(ctx context.Context, input *codedeploy.TagResourceInput) (*codedeploy.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *CodeDeployActivities) UntagResource(ctx context.Context, input *codedeploy.UntagResourceInput) (*codedeploy.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *CodeDeployActivities) UpdateApplication(ctx context.Context, input *codedeploy.UpdateApplicationInput) (*codedeploy.UpdateApplicationOutput, error) {
	return a.client.UpdateApplicationWithContext(ctx, input)
}

func (a *CodeDeployActivities) UpdateDeploymentGroup(ctx context.Context, input *codedeploy.UpdateDeploymentGroupInput) (*codedeploy.UpdateDeploymentGroupOutput, error) {
	return a.client.UpdateDeploymentGroupWithContext(ctx, input)
}

func (a *CodeDeployActivities) WaitUntilDeploymentSuccessful(ctx context.Context, input *codedeploy.GetDeploymentInput) error {
	return a.client.WaitUntilDeploymentSuccessfulWithContext(ctx, input)

}
