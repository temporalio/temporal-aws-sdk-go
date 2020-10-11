// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package serverlessapplicationrepository

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/serverlessapplicationrepository"
	"github.com/aws/aws-sdk-go/service/serverlessapplicationrepository/serverlessapplicationrepositoryiface"
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
	client serverlessapplicationrepositoryiface.ServerlessApplicationRepositoryAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := serverlessapplicationrepository.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (serverlessapplicationrepositoryiface.ServerlessApplicationRepositoryAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return serverlessapplicationrepository.New(sess), nil
}

func (a *Activities) CreateApplication(ctx context.Context, input *serverlessapplicationrepository.CreateApplicationRequest) (*serverlessapplicationrepository.CreateApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateApplicationWithContext(ctx, input)
}

func (a *Activities) CreateApplicationVersion(ctx context.Context, input *serverlessapplicationrepository.CreateApplicationVersionRequest) (*serverlessapplicationrepository.CreateApplicationVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateApplicationVersionWithContext(ctx, input)
}

func (a *Activities) CreateCloudFormationChangeSet(ctx context.Context, input *serverlessapplicationrepository.CreateCloudFormationChangeSetRequest) (*serverlessapplicationrepository.CreateCloudFormationChangeSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateCloudFormationChangeSetWithContext(ctx, input)
}

func (a *Activities) CreateCloudFormationTemplate(ctx context.Context, input *serverlessapplicationrepository.CreateCloudFormationTemplateInput) (*serverlessapplicationrepository.CreateCloudFormationTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateCloudFormationTemplateWithContext(ctx, input)
}

func (a *Activities) DeleteApplication(ctx context.Context, input *serverlessapplicationrepository.DeleteApplicationInput) (*serverlessapplicationrepository.DeleteApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteApplicationWithContext(ctx, input)
}

func (a *Activities) GetApplication(ctx context.Context, input *serverlessapplicationrepository.GetApplicationInput) (*serverlessapplicationrepository.GetApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetApplicationWithContext(ctx, input)
}

func (a *Activities) GetApplicationPolicy(ctx context.Context, input *serverlessapplicationrepository.GetApplicationPolicyInput) (*serverlessapplicationrepository.GetApplicationPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetApplicationPolicyWithContext(ctx, input)
}

func (a *Activities) GetCloudFormationTemplate(ctx context.Context, input *serverlessapplicationrepository.GetCloudFormationTemplateInput) (*serverlessapplicationrepository.GetCloudFormationTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetCloudFormationTemplateWithContext(ctx, input)
}

func (a *Activities) ListApplicationDependencies(ctx context.Context, input *serverlessapplicationrepository.ListApplicationDependenciesInput) (*serverlessapplicationrepository.ListApplicationDependenciesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListApplicationDependenciesWithContext(ctx, input)
}

func (a *Activities) ListApplicationVersions(ctx context.Context, input *serverlessapplicationrepository.ListApplicationVersionsInput) (*serverlessapplicationrepository.ListApplicationVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListApplicationVersionsWithContext(ctx, input)
}

func (a *Activities) ListApplications(ctx context.Context, input *serverlessapplicationrepository.ListApplicationsInput) (*serverlessapplicationrepository.ListApplicationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListApplicationsWithContext(ctx, input)
}

func (a *Activities) PutApplicationPolicy(ctx context.Context, input *serverlessapplicationrepository.PutApplicationPolicyInput) (*serverlessapplicationrepository.PutApplicationPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutApplicationPolicyWithContext(ctx, input)
}

func (a *Activities) UnshareApplication(ctx context.Context, input *serverlessapplicationrepository.UnshareApplicationInput) (*serverlessapplicationrepository.UnshareApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UnshareApplicationWithContext(ctx, input)
}

func (a *Activities) UpdateApplication(ctx context.Context, input *serverlessapplicationrepository.UpdateApplicationRequest) (*serverlessapplicationrepository.UpdateApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateApplicationWithContext(ctx, input)
}
