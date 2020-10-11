// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package appconfig

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/appconfig"
	"github.com/aws/aws-sdk-go/service/appconfig/appconfigiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client appconfigiface.AppConfigAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := appconfig.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (appconfigiface.AppConfigAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return appconfig.New(sess), nil
}

func (a *Activities) CreateApplication(ctx context.Context, input *appconfig.CreateApplicationInput) (*appconfig.CreateApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateApplicationWithContext(ctx, input)
}

func (a *Activities) CreateConfigurationProfile(ctx context.Context, input *appconfig.CreateConfigurationProfileInput) (*appconfig.CreateConfigurationProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateConfigurationProfileWithContext(ctx, input)
}

func (a *Activities) CreateDeploymentStrategy(ctx context.Context, input *appconfig.CreateDeploymentStrategyInput) (*appconfig.CreateDeploymentStrategyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDeploymentStrategyWithContext(ctx, input)
}

func (a *Activities) CreateEnvironment(ctx context.Context, input *appconfig.CreateEnvironmentInput) (*appconfig.CreateEnvironmentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateEnvironmentWithContext(ctx, input)
}

func (a *Activities) CreateHostedConfigurationVersion(ctx context.Context, input *appconfig.CreateHostedConfigurationVersionInput) (*appconfig.CreateHostedConfigurationVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateHostedConfigurationVersionWithContext(ctx, input)
}

func (a *Activities) DeleteApplication(ctx context.Context, input *appconfig.DeleteApplicationInput) (*appconfig.DeleteApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteApplicationWithContext(ctx, input)
}

func (a *Activities) DeleteConfigurationProfile(ctx context.Context, input *appconfig.DeleteConfigurationProfileInput) (*appconfig.DeleteConfigurationProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteConfigurationProfileWithContext(ctx, input)
}

func (a *Activities) DeleteDeploymentStrategy(ctx context.Context, input *appconfig.DeleteDeploymentStrategyInput) (*appconfig.DeleteDeploymentStrategyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDeploymentStrategyWithContext(ctx, input)
}

func (a *Activities) DeleteEnvironment(ctx context.Context, input *appconfig.DeleteEnvironmentInput) (*appconfig.DeleteEnvironmentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteEnvironmentWithContext(ctx, input)
}

func (a *Activities) DeleteHostedConfigurationVersion(ctx context.Context, input *appconfig.DeleteHostedConfigurationVersionInput) (*appconfig.DeleteHostedConfigurationVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteHostedConfigurationVersionWithContext(ctx, input)
}

func (a *Activities) GetApplication(ctx context.Context, input *appconfig.GetApplicationInput) (*appconfig.GetApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetApplicationWithContext(ctx, input)
}

func (a *Activities) GetConfiguration(ctx context.Context, input *appconfig.GetConfigurationInput) (*appconfig.GetConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetConfigurationWithContext(ctx, input)
}

func (a *Activities) GetConfigurationProfile(ctx context.Context, input *appconfig.GetConfigurationProfileInput) (*appconfig.GetConfigurationProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetConfigurationProfileWithContext(ctx, input)
}

func (a *Activities) GetDeployment(ctx context.Context, input *appconfig.GetDeploymentInput) (*appconfig.GetDeploymentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDeploymentWithContext(ctx, input)
}

func (a *Activities) GetDeploymentStrategy(ctx context.Context, input *appconfig.GetDeploymentStrategyInput) (*appconfig.GetDeploymentStrategyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDeploymentStrategyWithContext(ctx, input)
}

func (a *Activities) GetEnvironment(ctx context.Context, input *appconfig.GetEnvironmentInput) (*appconfig.GetEnvironmentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetEnvironmentWithContext(ctx, input)
}

func (a *Activities) GetHostedConfigurationVersion(ctx context.Context, input *appconfig.GetHostedConfigurationVersionInput) (*appconfig.GetHostedConfigurationVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetHostedConfigurationVersionWithContext(ctx, input)
}

func (a *Activities) ListApplications(ctx context.Context, input *appconfig.ListApplicationsInput) (*appconfig.ListApplicationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListApplicationsWithContext(ctx, input)
}

func (a *Activities) ListConfigurationProfiles(ctx context.Context, input *appconfig.ListConfigurationProfilesInput) (*appconfig.ListConfigurationProfilesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListConfigurationProfilesWithContext(ctx, input)
}

func (a *Activities) ListDeploymentStrategies(ctx context.Context, input *appconfig.ListDeploymentStrategiesInput) (*appconfig.ListDeploymentStrategiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDeploymentStrategiesWithContext(ctx, input)
}

func (a *Activities) ListDeployments(ctx context.Context, input *appconfig.ListDeploymentsInput) (*appconfig.ListDeploymentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDeploymentsWithContext(ctx, input)
}

func (a *Activities) ListEnvironments(ctx context.Context, input *appconfig.ListEnvironmentsInput) (*appconfig.ListEnvironmentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListEnvironmentsWithContext(ctx, input)
}

func (a *Activities) ListHostedConfigurationVersions(ctx context.Context, input *appconfig.ListHostedConfigurationVersionsInput) (*appconfig.ListHostedConfigurationVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListHostedConfigurationVersionsWithContext(ctx, input)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *appconfig.ListTagsForResourceInput) (*appconfig.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *Activities) StartDeployment(ctx context.Context, input *appconfig.StartDeploymentInput) (*appconfig.StartDeploymentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartDeploymentWithContext(ctx, input)
}

func (a *Activities) StopDeployment(ctx context.Context, input *appconfig.StopDeploymentInput) (*appconfig.StopDeploymentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopDeploymentWithContext(ctx, input)
}

func (a *Activities) TagResource(ctx context.Context, input *appconfig.TagResourceInput) (*appconfig.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *Activities) UntagResource(ctx context.Context, input *appconfig.UntagResourceInput) (*appconfig.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *Activities) UpdateApplication(ctx context.Context, input *appconfig.UpdateApplicationInput) (*appconfig.UpdateApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateApplicationWithContext(ctx, input)
}

func (a *Activities) UpdateConfigurationProfile(ctx context.Context, input *appconfig.UpdateConfigurationProfileInput) (*appconfig.UpdateConfigurationProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateConfigurationProfileWithContext(ctx, input)
}

func (a *Activities) UpdateDeploymentStrategy(ctx context.Context, input *appconfig.UpdateDeploymentStrategyInput) (*appconfig.UpdateDeploymentStrategyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDeploymentStrategyWithContext(ctx, input)
}

func (a *Activities) UpdateEnvironment(ctx context.Context, input *appconfig.UpdateEnvironmentInput) (*appconfig.UpdateEnvironmentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateEnvironmentWithContext(ctx, input)
}

func (a *Activities) ValidateConfiguration(ctx context.Context, input *appconfig.ValidateConfigurationInput) (*appconfig.ValidateConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ValidateConfigurationWithContext(ctx, input)
}