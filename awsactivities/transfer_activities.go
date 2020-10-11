// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/transfer"
	"github.com/aws/aws-sdk-go/service/transfer/transferiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type TransferActivities struct {
	client transferiface.TransferAPI

	sessionFactory SessionFactory
}

func NewTransferActivities(sess *session.Session, config ...*aws.Config) *TransferActivities {
	client := transfer.New(sess, config...)
	return &TransferActivities{client: client}
}

func NewTransferActivitiesWithSessionFactory(sessionFactory SessionFactory) *TransferActivities {
	return &TransferActivities{sessionFactory: sessionFactory}
}

func (a *TransferActivities) getClient(ctx context.Context) (transferiface.TransferAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return transfer.New(sess), nil
}

func (a *TransferActivities) CreateServer(ctx context.Context, input *transfer.CreateServerInput) (*transfer.CreateServerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateServerWithContext(ctx, input)
}

func (a *TransferActivities) CreateUser(ctx context.Context, input *transfer.CreateUserInput) (*transfer.CreateUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateUserWithContext(ctx, input)
}

func (a *TransferActivities) DeleteServer(ctx context.Context, input *transfer.DeleteServerInput) (*transfer.DeleteServerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteServerWithContext(ctx, input)
}

func (a *TransferActivities) DeleteSshPublicKey(ctx context.Context, input *transfer.DeleteSshPublicKeyInput) (*transfer.DeleteSshPublicKeyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteSshPublicKeyWithContext(ctx, input)
}

func (a *TransferActivities) DeleteUser(ctx context.Context, input *transfer.DeleteUserInput) (*transfer.DeleteUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteUserWithContext(ctx, input)
}

func (a *TransferActivities) DescribeSecurityPolicy(ctx context.Context, input *transfer.DescribeSecurityPolicyInput) (*transfer.DescribeSecurityPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeSecurityPolicyWithContext(ctx, input)
}

func (a *TransferActivities) DescribeServer(ctx context.Context, input *transfer.DescribeServerInput) (*transfer.DescribeServerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeServerWithContext(ctx, input)
}

func (a *TransferActivities) DescribeUser(ctx context.Context, input *transfer.DescribeUserInput) (*transfer.DescribeUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeUserWithContext(ctx, input)
}

func (a *TransferActivities) ImportSshPublicKey(ctx context.Context, input *transfer.ImportSshPublicKeyInput) (*transfer.ImportSshPublicKeyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ImportSshPublicKeyWithContext(ctx, input)
}

func (a *TransferActivities) ListSecurityPolicies(ctx context.Context, input *transfer.ListSecurityPoliciesInput) (*transfer.ListSecurityPoliciesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListSecurityPoliciesWithContext(ctx, input)
}

func (a *TransferActivities) ListServers(ctx context.Context, input *transfer.ListServersInput) (*transfer.ListServersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListServersWithContext(ctx, input)
}

func (a *TransferActivities) ListTagsForResource(ctx context.Context, input *transfer.ListTagsForResourceInput) (*transfer.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *TransferActivities) ListUsers(ctx context.Context, input *transfer.ListUsersInput) (*transfer.ListUsersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListUsersWithContext(ctx, input)
}

func (a *TransferActivities) StartServer(ctx context.Context, input *transfer.StartServerInput) (*transfer.StartServerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartServerWithContext(ctx, input)
}

func (a *TransferActivities) StopServer(ctx context.Context, input *transfer.StopServerInput) (*transfer.StopServerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopServerWithContext(ctx, input)
}

func (a *TransferActivities) TagResource(ctx context.Context, input *transfer.TagResourceInput) (*transfer.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *TransferActivities) TestIdentityProvider(ctx context.Context, input *transfer.TestIdentityProviderInput) (*transfer.TestIdentityProviderOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TestIdentityProviderWithContext(ctx, input)
}

func (a *TransferActivities) UntagResource(ctx context.Context, input *transfer.UntagResourceInput) (*transfer.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *TransferActivities) UpdateServer(ctx context.Context, input *transfer.UpdateServerInput) (*transfer.UpdateServerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateServerWithContext(ctx, input)
}

func (a *TransferActivities) UpdateUser(ctx context.Context, input *transfer.UpdateUserInput) (*transfer.UpdateUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateUserWithContext(ctx, input)
}
