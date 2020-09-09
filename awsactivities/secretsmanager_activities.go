
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
)

type SecretsManagerActivities struct {
    client secretsmanageriface.SecretsManagerAPI
}

func NewSecretsManagerActivities(session *session.Session, config ...*aws.Config) *SecretsManagerActivities {
    client := secretsmanager.New(session, config...)
    return &SecretsManagerActivities{client: client}
}

func (a *SecretsManagerActivities) CancelRotateSecret(input *secretsmanager.CancelRotateSecretInput) (*secretsmanager.CancelRotateSecretOutput, error) {
    return a.client.CancelRotateSecret(input)
}

func (a *SecretsManagerActivities) CreateSecret(input *secretsmanager.CreateSecretInput) (*secretsmanager.CreateSecretOutput, error) {
    return a.client.CreateSecret(input)
}

func (a *SecretsManagerActivities) DeleteResourcePolicy(input *secretsmanager.DeleteResourcePolicyInput) (*secretsmanager.DeleteResourcePolicyOutput, error) {
    return a.client.DeleteResourcePolicy(input)
}

func (a *SecretsManagerActivities) DeleteSecret(input *secretsmanager.DeleteSecretInput) (*secretsmanager.DeleteSecretOutput, error) {
    return a.client.DeleteSecret(input)
}

func (a *SecretsManagerActivities) DescribeSecret(input *secretsmanager.DescribeSecretInput) (*secretsmanager.DescribeSecretOutput, error) {
    return a.client.DescribeSecret(input)
}

func (a *SecretsManagerActivities) GetRandomPassword(input *secretsmanager.GetRandomPasswordInput) (*secretsmanager.GetRandomPasswordOutput, error) {
    return a.client.GetRandomPassword(input)
}

func (a *SecretsManagerActivities) GetResourcePolicy(input *secretsmanager.GetResourcePolicyInput) (*secretsmanager.GetResourcePolicyOutput, error) {
    return a.client.GetResourcePolicy(input)
}

func (a *SecretsManagerActivities) GetSecretValue(input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error) {
    return a.client.GetSecretValue(input)
}

func (a *SecretsManagerActivities) ListSecretVersionIds(input *secretsmanager.ListSecretVersionIdsInput) (*secretsmanager.ListSecretVersionIdsOutput, error) {
    return a.client.ListSecretVersionIds(input)
}

func (a *SecretsManagerActivities) ListSecrets(input *secretsmanager.ListSecretsInput) (*secretsmanager.ListSecretsOutput, error) {
    return a.client.ListSecrets(input)
}

func (a *SecretsManagerActivities) PutResourcePolicy(input *secretsmanager.PutResourcePolicyInput) (*secretsmanager.PutResourcePolicyOutput, error) {
    return a.client.PutResourcePolicy(input)
}

func (a *SecretsManagerActivities) PutSecretValue(input *secretsmanager.PutSecretValueInput) (*secretsmanager.PutSecretValueOutput, error) {
    return a.client.PutSecretValue(input)
}

func (a *SecretsManagerActivities) RestoreSecret(input *secretsmanager.RestoreSecretInput) (*secretsmanager.RestoreSecretOutput, error) {
    return a.client.RestoreSecret(input)
}

func (a *SecretsManagerActivities) RotateSecret(input *secretsmanager.RotateSecretInput) (*secretsmanager.RotateSecretOutput, error) {
    return a.client.RotateSecret(input)
}

func (a *SecretsManagerActivities) TagResource(input *secretsmanager.TagResourceInput) (*secretsmanager.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *SecretsManagerActivities) UntagResource(input *secretsmanager.UntagResourceInput) (*secretsmanager.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *SecretsManagerActivities) UpdateSecret(input *secretsmanager.UpdateSecretInput) (*secretsmanager.UpdateSecretOutput, error) {
    return a.client.UpdateSecret(input)
}

func (a *SecretsManagerActivities) UpdateSecretVersionStage(input *secretsmanager.UpdateSecretVersionStageInput) (*secretsmanager.UpdateSecretVersionStageOutput, error) {
    return a.client.UpdateSecretVersionStage(input)
}

func (a *SecretsManagerActivities) ValidateResourcePolicy(input *secretsmanager.ValidateResourcePolicyInput) (*secretsmanager.ValidateResourcePolicyOutput, error) {
    return a.client.ValidateResourcePolicy(input)
}
