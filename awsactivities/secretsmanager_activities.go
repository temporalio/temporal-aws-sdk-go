// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type SecretsManagerActivities struct {
	client secretsmanageriface.SecretsManagerAPI

	sessionFactory SessionFactory
}

func NewSecretsManagerActivities(sess *session.Session, config ...*aws.Config) *SecretsManagerActivities {
	client := secretsmanager.New(sess, config...)
	return &SecretsManagerActivities{client: client}
}

func NewSecretsManagerActivitiesWithSessionFactory(sessionFactory SessionFactory) *SecretsManagerActivities {
	return &SecretsManagerActivities{sessionFactory: sessionFactory}
}

func (a *SecretsManagerActivities) getClient(ctx context.Context) (secretsmanageriface.SecretsManagerAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return secretsmanager.New(sess), nil
}

func (a *SecretsManagerActivities) CancelRotateSecret(ctx context.Context, input *secretsmanager.CancelRotateSecretInput) (*secretsmanager.CancelRotateSecretOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CancelRotateSecretWithContext(ctx, input)
}

func (a *SecretsManagerActivities) CreateSecret(ctx context.Context, input *secretsmanager.CreateSecretInput) (*secretsmanager.CreateSecretOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateSecretWithContext(ctx, input)
}

func (a *SecretsManagerActivities) DeleteResourcePolicy(ctx context.Context, input *secretsmanager.DeleteResourcePolicyInput) (*secretsmanager.DeleteResourcePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteResourcePolicyWithContext(ctx, input)
}

func (a *SecretsManagerActivities) DeleteSecret(ctx context.Context, input *secretsmanager.DeleteSecretInput) (*secretsmanager.DeleteSecretOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteSecretWithContext(ctx, input)
}

func (a *SecretsManagerActivities) DescribeSecret(ctx context.Context, input *secretsmanager.DescribeSecretInput) (*secretsmanager.DescribeSecretOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeSecretWithContext(ctx, input)
}

func (a *SecretsManagerActivities) GetRandomPassword(ctx context.Context, input *secretsmanager.GetRandomPasswordInput) (*secretsmanager.GetRandomPasswordOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRandomPasswordWithContext(ctx, input)
}

func (a *SecretsManagerActivities) GetResourcePolicy(ctx context.Context, input *secretsmanager.GetResourcePolicyInput) (*secretsmanager.GetResourcePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetResourcePolicyWithContext(ctx, input)
}

func (a *SecretsManagerActivities) GetSecretValue(ctx context.Context, input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSecretValueWithContext(ctx, input)
}

func (a *SecretsManagerActivities) ListSecretVersionIds(ctx context.Context, input *secretsmanager.ListSecretVersionIdsInput) (*secretsmanager.ListSecretVersionIdsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListSecretVersionIdsWithContext(ctx, input)
}

func (a *SecretsManagerActivities) ListSecrets(ctx context.Context, input *secretsmanager.ListSecretsInput) (*secretsmanager.ListSecretsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListSecretsWithContext(ctx, input)
}

func (a *SecretsManagerActivities) PutResourcePolicy(ctx context.Context, input *secretsmanager.PutResourcePolicyInput) (*secretsmanager.PutResourcePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutResourcePolicyWithContext(ctx, input)
}

func (a *SecretsManagerActivities) PutSecretValue(ctx context.Context, input *secretsmanager.PutSecretValueInput) (*secretsmanager.PutSecretValueOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutSecretValueWithContext(ctx, input)
}

func (a *SecretsManagerActivities) RestoreSecret(ctx context.Context, input *secretsmanager.RestoreSecretInput) (*secretsmanager.RestoreSecretOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RestoreSecretWithContext(ctx, input)
}

func (a *SecretsManagerActivities) RotateSecret(ctx context.Context, input *secretsmanager.RotateSecretInput) (*secretsmanager.RotateSecretOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RotateSecretWithContext(ctx, input)
}

func (a *SecretsManagerActivities) TagResource(ctx context.Context, input *secretsmanager.TagResourceInput) (*secretsmanager.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *SecretsManagerActivities) UntagResource(ctx context.Context, input *secretsmanager.UntagResourceInput) (*secretsmanager.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *SecretsManagerActivities) UpdateSecret(ctx context.Context, input *secretsmanager.UpdateSecretInput) (*secretsmanager.UpdateSecretOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateSecretWithContext(ctx, input)
}

func (a *SecretsManagerActivities) UpdateSecretVersionStage(ctx context.Context, input *secretsmanager.UpdateSecretVersionStageInput) (*secretsmanager.UpdateSecretVersionStageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateSecretVersionStageWithContext(ctx, input)
}

func (a *SecretsManagerActivities) ValidateResourcePolicy(ctx context.Context, input *secretsmanager.ValidateResourcePolicyInput) (*secretsmanager.ValidateResourcePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ValidateResourcePolicyWithContext(ctx, input)
}
