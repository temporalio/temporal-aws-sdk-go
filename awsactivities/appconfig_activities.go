
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/appconfig"
	"github.com/aws/aws-sdk-go/service/appconfig/appconfigiface"
)

type AppConfigActivities struct {
	client appconfigiface.AppConfigAPI
}

func NewAppConfigActivities(client appconfigiface.AppConfigAPI) *AppConfigActivities {
    return &AppConfigActivities{client: client}
}

func (a *AppConfigActivities) CreateApplication(input *appconfig.CreateApplicationInput) (*appconfig.CreateApplicationOutput, error) {
    return a.client.CreateApplication(input)
}

func (a *AppConfigActivities) CreateConfigurationProfile(input *appconfig.CreateConfigurationProfileInput) (*appconfig.CreateConfigurationProfileOutput, error) {
    return a.client.CreateConfigurationProfile(input)
}

func (a *AppConfigActivities) CreateDeploymentStrategy(input *appconfig.CreateDeploymentStrategyInput) (*appconfig.CreateDeploymentStrategyOutput, error) {
    return a.client.CreateDeploymentStrategy(input)
}

func (a *AppConfigActivities) CreateEnvironment(input *appconfig.CreateEnvironmentInput) (*appconfig.CreateEnvironmentOutput, error) {
    return a.client.CreateEnvironment(input)
}

func (a *AppConfigActivities) CreateHostedConfigurationVersion(input *appconfig.CreateHostedConfigurationVersionInput) (*appconfig.CreateHostedConfigurationVersionOutput, error) {
    return a.client.CreateHostedConfigurationVersion(input)
}

func (a *AppConfigActivities) DeleteApplication(input *appconfig.DeleteApplicationInput) (*appconfig.DeleteApplicationOutput, error) {
    return a.client.DeleteApplication(input)
}

func (a *AppConfigActivities) DeleteConfigurationProfile(input *appconfig.DeleteConfigurationProfileInput) (*appconfig.DeleteConfigurationProfileOutput, error) {
    return a.client.DeleteConfigurationProfile(input)
}

func (a *AppConfigActivities) DeleteDeploymentStrategy(input *appconfig.DeleteDeploymentStrategyInput) (*appconfig.DeleteDeploymentStrategyOutput, error) {
    return a.client.DeleteDeploymentStrategy(input)
}

func (a *AppConfigActivities) DeleteEnvironment(input *appconfig.DeleteEnvironmentInput) (*appconfig.DeleteEnvironmentOutput, error) {
    return a.client.DeleteEnvironment(input)
}

func (a *AppConfigActivities) DeleteHostedConfigurationVersion(input *appconfig.DeleteHostedConfigurationVersionInput) (*appconfig.DeleteHostedConfigurationVersionOutput, error) {
    return a.client.DeleteHostedConfigurationVersion(input)
}

func (a *AppConfigActivities) GetApplication(input *appconfig.GetApplicationInput) (*appconfig.GetApplicationOutput, error) {
    return a.client.GetApplication(input)
}

func (a *AppConfigActivities) GetConfiguration(input *appconfig.GetConfigurationInput) (*appconfig.GetConfigurationOutput, error) {
    return a.client.GetConfiguration(input)
}

func (a *AppConfigActivities) GetConfigurationProfile(input *appconfig.GetConfigurationProfileInput) (*appconfig.GetConfigurationProfileOutput, error) {
    return a.client.GetConfigurationProfile(input)
}

func (a *AppConfigActivities) GetDeployment(input *appconfig.GetDeploymentInput) (*appconfig.GetDeploymentOutput, error) {
    return a.client.GetDeployment(input)
}

func (a *AppConfigActivities) GetDeploymentStrategy(input *appconfig.GetDeploymentStrategyInput) (*appconfig.GetDeploymentStrategyOutput, error) {
    return a.client.GetDeploymentStrategy(input)
}

func (a *AppConfigActivities) GetEnvironment(input *appconfig.GetEnvironmentInput) (*appconfig.GetEnvironmentOutput, error) {
    return a.client.GetEnvironment(input)
}

func (a *AppConfigActivities) GetHostedConfigurationVersion(input *appconfig.GetHostedConfigurationVersionInput) (*appconfig.GetHostedConfigurationVersionOutput, error) {
    return a.client.GetHostedConfigurationVersion(input)
}

func (a *AppConfigActivities) ListApplications(input *appconfig.ListApplicationsInput) (*appconfig.ListApplicationsOutput, error) {
    return a.client.ListApplications(input)
}

func (a *AppConfigActivities) ListConfigurationProfiles(input *appconfig.ListConfigurationProfilesInput) (*appconfig.ListConfigurationProfilesOutput, error) {
    return a.client.ListConfigurationProfiles(input)
}

func (a *AppConfigActivities) ListDeploymentStrategies(input *appconfig.ListDeploymentStrategiesInput) (*appconfig.ListDeploymentStrategiesOutput, error) {
    return a.client.ListDeploymentStrategies(input)
}

func (a *AppConfigActivities) ListDeployments(input *appconfig.ListDeploymentsInput) (*appconfig.ListDeploymentsOutput, error) {
    return a.client.ListDeployments(input)
}

func (a *AppConfigActivities) ListEnvironments(input *appconfig.ListEnvironmentsInput) (*appconfig.ListEnvironmentsOutput, error) {
    return a.client.ListEnvironments(input)
}

func (a *AppConfigActivities) ListHostedConfigurationVersions(input *appconfig.ListHostedConfigurationVersionsInput) (*appconfig.ListHostedConfigurationVersionsOutput, error) {
    return a.client.ListHostedConfigurationVersions(input)
}

func (a *AppConfigActivities) ListTagsForResource(input *appconfig.ListTagsForResourceInput) (*appconfig.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *AppConfigActivities) StartDeployment(input *appconfig.StartDeploymentInput) (*appconfig.StartDeploymentOutput, error) {
    return a.client.StartDeployment(input)
}

func (a *AppConfigActivities) StopDeployment(input *appconfig.StopDeploymentInput) (*appconfig.StopDeploymentOutput, error) {
    return a.client.StopDeployment(input)
}

func (a *AppConfigActivities) TagResource(input *appconfig.TagResourceInput) (*appconfig.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *AppConfigActivities) UntagResource(input *appconfig.UntagResourceInput) (*appconfig.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *AppConfigActivities) UpdateApplication(input *appconfig.UpdateApplicationInput) (*appconfig.UpdateApplicationOutput, error) {
    return a.client.UpdateApplication(input)
}

func (a *AppConfigActivities) UpdateConfigurationProfile(input *appconfig.UpdateConfigurationProfileInput) (*appconfig.UpdateConfigurationProfileOutput, error) {
    return a.client.UpdateConfigurationProfile(input)
}

func (a *AppConfigActivities) UpdateDeploymentStrategy(input *appconfig.UpdateDeploymentStrategyInput) (*appconfig.UpdateDeploymentStrategyOutput, error) {
    return a.client.UpdateDeploymentStrategy(input)
}

func (a *AppConfigActivities) UpdateEnvironment(input *appconfig.UpdateEnvironmentInput) (*appconfig.UpdateEnvironmentOutput, error) {
    return a.client.UpdateEnvironment(input)
}

func (a *AppConfigActivities) ValidateConfiguration(input *appconfig.ValidateConfigurationInput) (*appconfig.ValidateConfigurationOutput, error) {
    return a.client.ValidateConfiguration(input)
}
