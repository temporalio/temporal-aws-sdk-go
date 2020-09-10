package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/appconfig"
	"github.com/aws/aws-sdk-go/service/appconfig/appconfigiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type AppConfigActivities struct {
	client appconfigiface.AppConfigAPI
}

func NewAppConfigActivities(session *session.Session, config ...*aws.Config) *AppConfigActivities {
	client := appconfig.New(session, config...)
	return &AppConfigActivities{client: client}
}

func (a *AppConfigActivities) CreateApplication(ctx context.Context, input *appconfig.CreateApplicationInput) (*appconfig.CreateApplicationOutput, error) {
	return a.client.CreateApplicationWithContext(ctx, input)
}

func (a *AppConfigActivities) CreateConfigurationProfile(ctx context.Context, input *appconfig.CreateConfigurationProfileInput) (*appconfig.CreateConfigurationProfileOutput, error) {
	return a.client.CreateConfigurationProfileWithContext(ctx, input)
}

func (a *AppConfigActivities) CreateDeploymentStrategy(ctx context.Context, input *appconfig.CreateDeploymentStrategyInput) (*appconfig.CreateDeploymentStrategyOutput, error) {
	return a.client.CreateDeploymentStrategyWithContext(ctx, input)
}

func (a *AppConfigActivities) CreateEnvironment(ctx context.Context, input *appconfig.CreateEnvironmentInput) (*appconfig.CreateEnvironmentOutput, error) {
	return a.client.CreateEnvironmentWithContext(ctx, input)
}

func (a *AppConfigActivities) CreateHostedConfigurationVersion(ctx context.Context, input *appconfig.CreateHostedConfigurationVersionInput) (*appconfig.CreateHostedConfigurationVersionOutput, error) {
	return a.client.CreateHostedConfigurationVersionWithContext(ctx, input)
}

func (a *AppConfigActivities) DeleteApplication(ctx context.Context, input *appconfig.DeleteApplicationInput) (*appconfig.DeleteApplicationOutput, error) {
	return a.client.DeleteApplicationWithContext(ctx, input)
}

func (a *AppConfigActivities) DeleteConfigurationProfile(ctx context.Context, input *appconfig.DeleteConfigurationProfileInput) (*appconfig.DeleteConfigurationProfileOutput, error) {
	return a.client.DeleteConfigurationProfileWithContext(ctx, input)
}

func (a *AppConfigActivities) DeleteDeploymentStrategy(ctx context.Context, input *appconfig.DeleteDeploymentStrategyInput) (*appconfig.DeleteDeploymentStrategyOutput, error) {
	return a.client.DeleteDeploymentStrategyWithContext(ctx, input)
}

func (a *AppConfigActivities) DeleteEnvironment(ctx context.Context, input *appconfig.DeleteEnvironmentInput) (*appconfig.DeleteEnvironmentOutput, error) {
	return a.client.DeleteEnvironmentWithContext(ctx, input)
}

func (a *AppConfigActivities) DeleteHostedConfigurationVersion(ctx context.Context, input *appconfig.DeleteHostedConfigurationVersionInput) (*appconfig.DeleteHostedConfigurationVersionOutput, error) {
	return a.client.DeleteHostedConfigurationVersionWithContext(ctx, input)
}

func (a *AppConfigActivities) GetApplication(ctx context.Context, input *appconfig.GetApplicationInput) (*appconfig.GetApplicationOutput, error) {
	return a.client.GetApplicationWithContext(ctx, input)
}

func (a *AppConfigActivities) GetConfiguration(ctx context.Context, input *appconfig.GetConfigurationInput) (*appconfig.GetConfigurationOutput, error) {
	return a.client.GetConfigurationWithContext(ctx, input)
}

func (a *AppConfigActivities) GetConfigurationProfile(ctx context.Context, input *appconfig.GetConfigurationProfileInput) (*appconfig.GetConfigurationProfileOutput, error) {
	return a.client.GetConfigurationProfileWithContext(ctx, input)
}

func (a *AppConfigActivities) GetDeployment(ctx context.Context, input *appconfig.GetDeploymentInput) (*appconfig.GetDeploymentOutput, error) {
	return a.client.GetDeploymentWithContext(ctx, input)
}

func (a *AppConfigActivities) GetDeploymentStrategy(ctx context.Context, input *appconfig.GetDeploymentStrategyInput) (*appconfig.GetDeploymentStrategyOutput, error) {
	return a.client.GetDeploymentStrategyWithContext(ctx, input)
}

func (a *AppConfigActivities) GetEnvironment(ctx context.Context, input *appconfig.GetEnvironmentInput) (*appconfig.GetEnvironmentOutput, error) {
	return a.client.GetEnvironmentWithContext(ctx, input)
}

func (a *AppConfigActivities) GetHostedConfigurationVersion(ctx context.Context, input *appconfig.GetHostedConfigurationVersionInput) (*appconfig.GetHostedConfigurationVersionOutput, error) {
	return a.client.GetHostedConfigurationVersionWithContext(ctx, input)
}

func (a *AppConfigActivities) ListApplications(ctx context.Context, input *appconfig.ListApplicationsInput) (*appconfig.ListApplicationsOutput, error) {
	return a.client.ListApplicationsWithContext(ctx, input)
}

func (a *AppConfigActivities) ListConfigurationProfiles(ctx context.Context, input *appconfig.ListConfigurationProfilesInput) (*appconfig.ListConfigurationProfilesOutput, error) {
	return a.client.ListConfigurationProfilesWithContext(ctx, input)
}

func (a *AppConfigActivities) ListDeploymentStrategies(ctx context.Context, input *appconfig.ListDeploymentStrategiesInput) (*appconfig.ListDeploymentStrategiesOutput, error) {
	return a.client.ListDeploymentStrategiesWithContext(ctx, input)
}

func (a *AppConfigActivities) ListDeployments(ctx context.Context, input *appconfig.ListDeploymentsInput) (*appconfig.ListDeploymentsOutput, error) {
	return a.client.ListDeploymentsWithContext(ctx, input)
}

func (a *AppConfigActivities) ListEnvironments(ctx context.Context, input *appconfig.ListEnvironmentsInput) (*appconfig.ListEnvironmentsOutput, error) {
	return a.client.ListEnvironmentsWithContext(ctx, input)
}

func (a *AppConfigActivities) ListHostedConfigurationVersions(ctx context.Context, input *appconfig.ListHostedConfigurationVersionsInput) (*appconfig.ListHostedConfigurationVersionsOutput, error) {
	return a.client.ListHostedConfigurationVersionsWithContext(ctx, input)
}

func (a *AppConfigActivities) ListTagsForResource(ctx context.Context, input *appconfig.ListTagsForResourceInput) (*appconfig.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *AppConfigActivities) StartDeployment(ctx context.Context, input *appconfig.StartDeploymentInput) (*appconfig.StartDeploymentOutput, error) {
	return a.client.StartDeploymentWithContext(ctx, input)
}

func (a *AppConfigActivities) StopDeployment(ctx context.Context, input *appconfig.StopDeploymentInput) (*appconfig.StopDeploymentOutput, error) {
	return a.client.StopDeploymentWithContext(ctx, input)
}

func (a *AppConfigActivities) TagResource(ctx context.Context, input *appconfig.TagResourceInput) (*appconfig.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *AppConfigActivities) UntagResource(ctx context.Context, input *appconfig.UntagResourceInput) (*appconfig.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *AppConfigActivities) UpdateApplication(ctx context.Context, input *appconfig.UpdateApplicationInput) (*appconfig.UpdateApplicationOutput, error) {
	return a.client.UpdateApplicationWithContext(ctx, input)
}

func (a *AppConfigActivities) UpdateConfigurationProfile(ctx context.Context, input *appconfig.UpdateConfigurationProfileInput) (*appconfig.UpdateConfigurationProfileOutput, error) {
	return a.client.UpdateConfigurationProfileWithContext(ctx, input)
}

func (a *AppConfigActivities) UpdateDeploymentStrategy(ctx context.Context, input *appconfig.UpdateDeploymentStrategyInput) (*appconfig.UpdateDeploymentStrategyOutput, error) {
	return a.client.UpdateDeploymentStrategyWithContext(ctx, input)
}

func (a *AppConfigActivities) UpdateEnvironment(ctx context.Context, input *appconfig.UpdateEnvironmentInput) (*appconfig.UpdateEnvironmentOutput, error) {
	return a.client.UpdateEnvironmentWithContext(ctx, input)
}

func (a *AppConfigActivities) ValidateConfiguration(ctx context.Context, input *appconfig.ValidateConfigurationInput) (*appconfig.ValidateConfigurationOutput, error) {
	return a.client.ValidateConfigurationWithContext(ctx, input)
}
