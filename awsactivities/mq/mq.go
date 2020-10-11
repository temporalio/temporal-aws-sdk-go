// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package mq

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mq"
	"github.com/aws/aws-sdk-go/service/mq/mqiface"
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
	client mqiface.MQAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := mq.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (mqiface.MQAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return mq.New(sess), nil
}

func (a *Activities) CreateBroker(ctx context.Context, input *mq.CreateBrokerRequest) (*mq.CreateBrokerResponse, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateBrokerWithContext(ctx, input)
}

func (a *Activities) CreateConfiguration(ctx context.Context, input *mq.CreateConfigurationRequest) (*mq.CreateConfigurationResponse, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateConfigurationWithContext(ctx, input)
}

func (a *Activities) CreateTags(ctx context.Context, input *mq.CreateTagsInput) (*mq.CreateTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateTagsWithContext(ctx, input)
}

func (a *Activities) CreateUser(ctx context.Context, input *mq.CreateUserRequest) (*mq.CreateUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateUserWithContext(ctx, input)
}

func (a *Activities) DeleteBroker(ctx context.Context, input *mq.DeleteBrokerInput) (*mq.DeleteBrokerResponse, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBrokerWithContext(ctx, input)
}

func (a *Activities) DeleteTags(ctx context.Context, input *mq.DeleteTagsInput) (*mq.DeleteTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteTagsWithContext(ctx, input)
}

func (a *Activities) DeleteUser(ctx context.Context, input *mq.DeleteUserInput) (*mq.DeleteUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteUserWithContext(ctx, input)
}

func (a *Activities) DescribeBroker(ctx context.Context, input *mq.DescribeBrokerInput) (*mq.DescribeBrokerResponse, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeBrokerWithContext(ctx, input)
}

func (a *Activities) DescribeBrokerEngineTypes(ctx context.Context, input *mq.DescribeBrokerEngineTypesInput) (*mq.DescribeBrokerEngineTypesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeBrokerEngineTypesWithContext(ctx, input)
}

func (a *Activities) DescribeBrokerInstanceOptions(ctx context.Context, input *mq.DescribeBrokerInstanceOptionsInput) (*mq.DescribeBrokerInstanceOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeBrokerInstanceOptionsWithContext(ctx, input)
}

func (a *Activities) DescribeConfiguration(ctx context.Context, input *mq.DescribeConfigurationInput) (*mq.DescribeConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeConfigurationWithContext(ctx, input)
}

func (a *Activities) DescribeConfigurationRevision(ctx context.Context, input *mq.DescribeConfigurationRevisionInput) (*mq.DescribeConfigurationRevisionResponse, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeConfigurationRevisionWithContext(ctx, input)
}

func (a *Activities) DescribeUser(ctx context.Context, input *mq.DescribeUserInput) (*mq.DescribeUserResponse, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeUserWithContext(ctx, input)
}

func (a *Activities) ListBrokers(ctx context.Context, input *mq.ListBrokersInput) (*mq.ListBrokersResponse, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListBrokersWithContext(ctx, input)
}

func (a *Activities) ListConfigurationRevisions(ctx context.Context, input *mq.ListConfigurationRevisionsInput) (*mq.ListConfigurationRevisionsResponse, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListConfigurationRevisionsWithContext(ctx, input)
}

func (a *Activities) ListConfigurations(ctx context.Context, input *mq.ListConfigurationsInput) (*mq.ListConfigurationsResponse, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListConfigurationsWithContext(ctx, input)
}

func (a *Activities) ListTags(ctx context.Context, input *mq.ListTagsInput) (*mq.ListTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsWithContext(ctx, input)
}

func (a *Activities) ListUsers(ctx context.Context, input *mq.ListUsersInput) (*mq.ListUsersResponse, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListUsersWithContext(ctx, input)
}

func (a *Activities) RebootBroker(ctx context.Context, input *mq.RebootBrokerInput) (*mq.RebootBrokerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RebootBrokerWithContext(ctx, input)
}

func (a *Activities) UpdateBroker(ctx context.Context, input *mq.UpdateBrokerRequest) (*mq.UpdateBrokerResponse, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateBrokerWithContext(ctx, input)
}

func (a *Activities) UpdateConfiguration(ctx context.Context, input *mq.UpdateConfigurationRequest) (*mq.UpdateConfigurationResponse, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateConfigurationWithContext(ctx, input)
}

func (a *Activities) UpdateUser(ctx context.Context, input *mq.UpdateUserRequest) (*mq.UpdateUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateUserWithContext(ctx, input)
}
