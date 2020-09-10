package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type SecretsManagerActivities struct {
	client secretsmanageriface.SecretsManagerAPI
}

func NewSecretsManagerActivities(session *session.Session, config ...*aws.Config) *SecretsManagerActivities {
	client := secretsmanager.New(session, config...)
	return &SecretsManagerActivities{client: client}
}

func (a *SecretsManagerActivities) CancelRotateSecret(ctx context.Context, input *secretsmanager.CancelRotateSecretInput) (*secretsmanager.CancelRotateSecretOutput, error) {
	return a.client.CancelRotateSecretWithContext(ctx, input)
}

func (a *SecretsManagerActivities) CreateSecret(ctx context.Context, input *secretsmanager.CreateSecretInput) (*secretsmanager.CreateSecretOutput, error) {
	return a.client.CreateSecretWithContext(ctx, input)
}

func (a *SecretsManagerActivities) DeleteResourcePolicy(ctx context.Context, input *secretsmanager.DeleteResourcePolicyInput) (*secretsmanager.DeleteResourcePolicyOutput, error) {
	return a.client.DeleteResourcePolicyWithContext(ctx, input)
}

func (a *SecretsManagerActivities) DeleteSecret(ctx context.Context, input *secretsmanager.DeleteSecretInput) (*secretsmanager.DeleteSecretOutput, error) {
	return a.client.DeleteSecretWithContext(ctx, input)
}

func (a *SecretsManagerActivities) DescribeSecret(ctx context.Context, input *secretsmanager.DescribeSecretInput) (*secretsmanager.DescribeSecretOutput, error) {
	return a.client.DescribeSecretWithContext(ctx, input)
}

func (a *SecretsManagerActivities) GetRandomPassword(ctx context.Context, input *secretsmanager.GetRandomPasswordInput) (*secretsmanager.GetRandomPasswordOutput, error) {
	return a.client.GetRandomPasswordWithContext(ctx, input)
}

func (a *SecretsManagerActivities) GetResourcePolicy(ctx context.Context, input *secretsmanager.GetResourcePolicyInput) (*secretsmanager.GetResourcePolicyOutput, error) {
	return a.client.GetResourcePolicyWithContext(ctx, input)
}

func (a *SecretsManagerActivities) GetSecretValue(ctx context.Context, input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error) {
	return a.client.GetSecretValueWithContext(ctx, input)
}

func (a *SecretsManagerActivities) ListSecretVersionIds(ctx context.Context, input *secretsmanager.ListSecretVersionIdsInput) (*secretsmanager.ListSecretVersionIdsOutput, error) {
	return a.client.ListSecretVersionIdsWithContext(ctx, input)
}

func (a *SecretsManagerActivities) ListSecrets(ctx context.Context, input *secretsmanager.ListSecretsInput) (*secretsmanager.ListSecretsOutput, error) {
	return a.client.ListSecretsWithContext(ctx, input)
}

func (a *SecretsManagerActivities) PutResourcePolicy(ctx context.Context, input *secretsmanager.PutResourcePolicyInput) (*secretsmanager.PutResourcePolicyOutput, error) {
	return a.client.PutResourcePolicyWithContext(ctx, input)
}

func (a *SecretsManagerActivities) PutSecretValue(ctx context.Context, input *secretsmanager.PutSecretValueInput) (*secretsmanager.PutSecretValueOutput, error) {
	return a.client.PutSecretValueWithContext(ctx, input)
}

func (a *SecretsManagerActivities) RestoreSecret(ctx context.Context, input *secretsmanager.RestoreSecretInput) (*secretsmanager.RestoreSecretOutput, error) {
	return a.client.RestoreSecretWithContext(ctx, input)
}

func (a *SecretsManagerActivities) RotateSecret(ctx context.Context, input *secretsmanager.RotateSecretInput) (*secretsmanager.RotateSecretOutput, error) {
	return a.client.RotateSecretWithContext(ctx, input)
}

func (a *SecretsManagerActivities) TagResource(ctx context.Context, input *secretsmanager.TagResourceInput) (*secretsmanager.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *SecretsManagerActivities) UntagResource(ctx context.Context, input *secretsmanager.UntagResourceInput) (*secretsmanager.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *SecretsManagerActivities) UpdateSecret(ctx context.Context, input *secretsmanager.UpdateSecretInput) (*secretsmanager.UpdateSecretOutput, error) {
	return a.client.UpdateSecretWithContext(ctx, input)
}

func (a *SecretsManagerActivities) UpdateSecretVersionStage(ctx context.Context, input *secretsmanager.UpdateSecretVersionStageInput) (*secretsmanager.UpdateSecretVersionStageOutput, error) {
	return a.client.UpdateSecretVersionStageWithContext(ctx, input)
}

func (a *SecretsManagerActivities) ValidateResourcePolicy(ctx context.Context, input *secretsmanager.ValidateResourcePolicyInput) (*secretsmanager.ValidateResourcePolicyOutput, error) {
	return a.client.ValidateResourcePolicyWithContext(ctx, input)
}
