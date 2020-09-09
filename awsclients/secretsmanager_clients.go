package awsclients

import (
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type SecretsManagerClient interface {
    CancelRotateSecret(ctx workflow.Context, input *secretsmanager.CancelRotateSecretInput) (*secretsmanager.CancelRotateSecretOutput, error)
    CancelRotateSecretAsync(ctx workflow.Context, input *secretsmanager.CancelRotateSecretInput) *SecretsmanagerCancelRotateSecretResult

    CreateSecret(ctx workflow.Context, input *secretsmanager.CreateSecretInput) (*secretsmanager.CreateSecretOutput, error)
    CreateSecretAsync(ctx workflow.Context, input *secretsmanager.CreateSecretInput) *SecretsmanagerCreateSecretResult

    DeleteResourcePolicy(ctx workflow.Context, input *secretsmanager.DeleteResourcePolicyInput) (*secretsmanager.DeleteResourcePolicyOutput, error)
    DeleteResourcePolicyAsync(ctx workflow.Context, input *secretsmanager.DeleteResourcePolicyInput) *SecretsmanagerDeleteResourcePolicyResult

    DeleteSecret(ctx workflow.Context, input *secretsmanager.DeleteSecretInput) (*secretsmanager.DeleteSecretOutput, error)
    DeleteSecretAsync(ctx workflow.Context, input *secretsmanager.DeleteSecretInput) *SecretsmanagerDeleteSecretResult

    DescribeSecret(ctx workflow.Context, input *secretsmanager.DescribeSecretInput) (*secretsmanager.DescribeSecretOutput, error)
    DescribeSecretAsync(ctx workflow.Context, input *secretsmanager.DescribeSecretInput) *SecretsmanagerDescribeSecretResult

    GetRandomPassword(ctx workflow.Context, input *secretsmanager.GetRandomPasswordInput) (*secretsmanager.GetRandomPasswordOutput, error)
    GetRandomPasswordAsync(ctx workflow.Context, input *secretsmanager.GetRandomPasswordInput) *SecretsmanagerGetRandomPasswordResult

    GetResourcePolicy(ctx workflow.Context, input *secretsmanager.GetResourcePolicyInput) (*secretsmanager.GetResourcePolicyOutput, error)
    GetResourcePolicyAsync(ctx workflow.Context, input *secretsmanager.GetResourcePolicyInput) *SecretsmanagerGetResourcePolicyResult

    GetSecretValue(ctx workflow.Context, input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error)
    GetSecretValueAsync(ctx workflow.Context, input *secretsmanager.GetSecretValueInput) *SecretsmanagerGetSecretValueResult

    ListSecretVersionIds(ctx workflow.Context, input *secretsmanager.ListSecretVersionIdsInput) (*secretsmanager.ListSecretVersionIdsOutput, error)
    ListSecretVersionIdsAsync(ctx workflow.Context, input *secretsmanager.ListSecretVersionIdsInput) *SecretsmanagerListSecretVersionIdsResult

    ListSecrets(ctx workflow.Context, input *secretsmanager.ListSecretsInput) (*secretsmanager.ListSecretsOutput, error)
    ListSecretsAsync(ctx workflow.Context, input *secretsmanager.ListSecretsInput) *SecretsmanagerListSecretsResult

    PutResourcePolicy(ctx workflow.Context, input *secretsmanager.PutResourcePolicyInput) (*secretsmanager.PutResourcePolicyOutput, error)
    PutResourcePolicyAsync(ctx workflow.Context, input *secretsmanager.PutResourcePolicyInput) *SecretsmanagerPutResourcePolicyResult

    PutSecretValue(ctx workflow.Context, input *secretsmanager.PutSecretValueInput) (*secretsmanager.PutSecretValueOutput, error)
    PutSecretValueAsync(ctx workflow.Context, input *secretsmanager.PutSecretValueInput) *SecretsmanagerPutSecretValueResult

    RestoreSecret(ctx workflow.Context, input *secretsmanager.RestoreSecretInput) (*secretsmanager.RestoreSecretOutput, error)
    RestoreSecretAsync(ctx workflow.Context, input *secretsmanager.RestoreSecretInput) *SecretsmanagerRestoreSecretResult

    RotateSecret(ctx workflow.Context, input *secretsmanager.RotateSecretInput) (*secretsmanager.RotateSecretOutput, error)
    RotateSecretAsync(ctx workflow.Context, input *secretsmanager.RotateSecretInput) *SecretsmanagerRotateSecretResult

    TagResource(ctx workflow.Context, input *secretsmanager.TagResourceInput) (*secretsmanager.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *secretsmanager.TagResourceInput) *SecretsmanagerTagResourceResult

    UntagResource(ctx workflow.Context, input *secretsmanager.UntagResourceInput) (*secretsmanager.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *secretsmanager.UntagResourceInput) *SecretsmanagerUntagResourceResult

    UpdateSecret(ctx workflow.Context, input *secretsmanager.UpdateSecretInput) (*secretsmanager.UpdateSecretOutput, error)
    UpdateSecretAsync(ctx workflow.Context, input *secretsmanager.UpdateSecretInput) *SecretsmanagerUpdateSecretResult

    UpdateSecretVersionStage(ctx workflow.Context, input *secretsmanager.UpdateSecretVersionStageInput) (*secretsmanager.UpdateSecretVersionStageOutput, error)
    UpdateSecretVersionStageAsync(ctx workflow.Context, input *secretsmanager.UpdateSecretVersionStageInput) *SecretsmanagerUpdateSecretVersionStageResult

    ValidateResourcePolicy(ctx workflow.Context, input *secretsmanager.ValidateResourcePolicyInput) (*secretsmanager.ValidateResourcePolicyOutput, error)
    ValidateResourcePolicyAsync(ctx workflow.Context, input *secretsmanager.ValidateResourcePolicyInput) *SecretsmanagerValidateResourcePolicyResult
}

type SecretsmanagerCancelRotateSecretResult struct {
	Result workflow.Future
}

func (r *SecretsmanagerCancelRotateSecretResult) Get(ctx workflow.Context) (*secretsmanager.CancelRotateSecretOutput, error) {
    var output secretsmanager.CancelRotateSecretOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecretsmanagerCreateSecretResult struct {
	Result workflow.Future
}

func (r *SecretsmanagerCreateSecretResult) Get(ctx workflow.Context) (*secretsmanager.CreateSecretOutput, error) {
    var output secretsmanager.CreateSecretOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecretsmanagerDeleteResourcePolicyResult struct {
	Result workflow.Future
}

func (r *SecretsmanagerDeleteResourcePolicyResult) Get(ctx workflow.Context) (*secretsmanager.DeleteResourcePolicyOutput, error) {
    var output secretsmanager.DeleteResourcePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecretsmanagerDeleteSecretResult struct {
	Result workflow.Future
}

func (r *SecretsmanagerDeleteSecretResult) Get(ctx workflow.Context) (*secretsmanager.DeleteSecretOutput, error) {
    var output secretsmanager.DeleteSecretOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecretsmanagerDescribeSecretResult struct {
	Result workflow.Future
}

func (r *SecretsmanagerDescribeSecretResult) Get(ctx workflow.Context) (*secretsmanager.DescribeSecretOutput, error) {
    var output secretsmanager.DescribeSecretOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecretsmanagerGetRandomPasswordResult struct {
	Result workflow.Future
}

func (r *SecretsmanagerGetRandomPasswordResult) Get(ctx workflow.Context) (*secretsmanager.GetRandomPasswordOutput, error) {
    var output secretsmanager.GetRandomPasswordOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecretsmanagerGetResourcePolicyResult struct {
	Result workflow.Future
}

func (r *SecretsmanagerGetResourcePolicyResult) Get(ctx workflow.Context) (*secretsmanager.GetResourcePolicyOutput, error) {
    var output secretsmanager.GetResourcePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecretsmanagerGetSecretValueResult struct {
	Result workflow.Future
}

func (r *SecretsmanagerGetSecretValueResult) Get(ctx workflow.Context) (*secretsmanager.GetSecretValueOutput, error) {
    var output secretsmanager.GetSecretValueOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecretsmanagerListSecretVersionIdsResult struct {
	Result workflow.Future
}

func (r *SecretsmanagerListSecretVersionIdsResult) Get(ctx workflow.Context) (*secretsmanager.ListSecretVersionIdsOutput, error) {
    var output secretsmanager.ListSecretVersionIdsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecretsmanagerListSecretsResult struct {
	Result workflow.Future
}

func (r *SecretsmanagerListSecretsResult) Get(ctx workflow.Context) (*secretsmanager.ListSecretsOutput, error) {
    var output secretsmanager.ListSecretsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecretsmanagerPutResourcePolicyResult struct {
	Result workflow.Future
}

func (r *SecretsmanagerPutResourcePolicyResult) Get(ctx workflow.Context) (*secretsmanager.PutResourcePolicyOutput, error) {
    var output secretsmanager.PutResourcePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecretsmanagerPutSecretValueResult struct {
	Result workflow.Future
}

func (r *SecretsmanagerPutSecretValueResult) Get(ctx workflow.Context) (*secretsmanager.PutSecretValueOutput, error) {
    var output secretsmanager.PutSecretValueOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecretsmanagerRestoreSecretResult struct {
	Result workflow.Future
}

func (r *SecretsmanagerRestoreSecretResult) Get(ctx workflow.Context) (*secretsmanager.RestoreSecretOutput, error) {
    var output secretsmanager.RestoreSecretOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecretsmanagerRotateSecretResult struct {
	Result workflow.Future
}

func (r *SecretsmanagerRotateSecretResult) Get(ctx workflow.Context) (*secretsmanager.RotateSecretOutput, error) {
    var output secretsmanager.RotateSecretOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecretsmanagerTagResourceResult struct {
	Result workflow.Future
}

func (r *SecretsmanagerTagResourceResult) Get(ctx workflow.Context) (*secretsmanager.TagResourceOutput, error) {
    var output secretsmanager.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecretsmanagerUntagResourceResult struct {
	Result workflow.Future
}

func (r *SecretsmanagerUntagResourceResult) Get(ctx workflow.Context) (*secretsmanager.UntagResourceOutput, error) {
    var output secretsmanager.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecretsmanagerUpdateSecretResult struct {
	Result workflow.Future
}

func (r *SecretsmanagerUpdateSecretResult) Get(ctx workflow.Context) (*secretsmanager.UpdateSecretOutput, error) {
    var output secretsmanager.UpdateSecretOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecretsmanagerUpdateSecretVersionStageResult struct {
	Result workflow.Future
}

func (r *SecretsmanagerUpdateSecretVersionStageResult) Get(ctx workflow.Context) (*secretsmanager.UpdateSecretVersionStageOutput, error) {
    var output secretsmanager.UpdateSecretVersionStageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecretsmanagerValidateResourcePolicyResult struct {
	Result workflow.Future
}

func (r *SecretsmanagerValidateResourcePolicyResult) Get(ctx workflow.Context) (*secretsmanager.ValidateResourcePolicyOutput, error) {
    var output secretsmanager.ValidateResourcePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SecretsManagerStub struct {
    activities awsactivities.SecretsManagerActivities
}

func NewSecretsManagerStub() SecretsManagerClient {
    return &SecretsManagerStub{}
}

func (a *SecretsManagerStub) CancelRotateSecret(ctx workflow.Context, input *secretsmanager.CancelRotateSecretInput) (*secretsmanager.CancelRotateSecretOutput, error) {
    var output secretsmanager.CancelRotateSecretOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelRotateSecret, input).Get(ctx, &output)
    return &output, err
}

func (a *SecretsManagerStub) CancelRotateSecretAsync(ctx workflow.Context, input *secretsmanager.CancelRotateSecretInput) *SecretsmanagerCancelRotateSecretResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelRotateSecret, input)
    return &SecretsmanagerCancelRotateSecretResult{Result: future}
}

func (a *SecretsManagerStub) CreateSecret(ctx workflow.Context, input *secretsmanager.CreateSecretInput) (*secretsmanager.CreateSecretOutput, error) {
    var output secretsmanager.CreateSecretOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSecret, input).Get(ctx, &output)
    return &output, err
}

func (a *SecretsManagerStub) CreateSecretAsync(ctx workflow.Context, input *secretsmanager.CreateSecretInput) *SecretsmanagerCreateSecretResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSecret, input)
    return &SecretsmanagerCreateSecretResult{Result: future}
}

func (a *SecretsManagerStub) DeleteResourcePolicy(ctx workflow.Context, input *secretsmanager.DeleteResourcePolicyInput) (*secretsmanager.DeleteResourcePolicyOutput, error) {
    var output secretsmanager.DeleteResourcePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteResourcePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *SecretsManagerStub) DeleteResourcePolicyAsync(ctx workflow.Context, input *secretsmanager.DeleteResourcePolicyInput) *SecretsmanagerDeleteResourcePolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteResourcePolicy, input)
    return &SecretsmanagerDeleteResourcePolicyResult{Result: future}
}

func (a *SecretsManagerStub) DeleteSecret(ctx workflow.Context, input *secretsmanager.DeleteSecretInput) (*secretsmanager.DeleteSecretOutput, error) {
    var output secretsmanager.DeleteSecretOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSecret, input).Get(ctx, &output)
    return &output, err
}

func (a *SecretsManagerStub) DeleteSecretAsync(ctx workflow.Context, input *secretsmanager.DeleteSecretInput) *SecretsmanagerDeleteSecretResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSecret, input)
    return &SecretsmanagerDeleteSecretResult{Result: future}
}

func (a *SecretsManagerStub) DescribeSecret(ctx workflow.Context, input *secretsmanager.DescribeSecretInput) (*secretsmanager.DescribeSecretOutput, error) {
    var output secretsmanager.DescribeSecretOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSecret, input).Get(ctx, &output)
    return &output, err
}

func (a *SecretsManagerStub) DescribeSecretAsync(ctx workflow.Context, input *secretsmanager.DescribeSecretInput) *SecretsmanagerDescribeSecretResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSecret, input)
    return &SecretsmanagerDescribeSecretResult{Result: future}
}

func (a *SecretsManagerStub) GetRandomPassword(ctx workflow.Context, input *secretsmanager.GetRandomPasswordInput) (*secretsmanager.GetRandomPasswordOutput, error) {
    var output secretsmanager.GetRandomPasswordOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRandomPassword, input).Get(ctx, &output)
    return &output, err
}

func (a *SecretsManagerStub) GetRandomPasswordAsync(ctx workflow.Context, input *secretsmanager.GetRandomPasswordInput) *SecretsmanagerGetRandomPasswordResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetRandomPassword, input)
    return &SecretsmanagerGetRandomPasswordResult{Result: future}
}

func (a *SecretsManagerStub) GetResourcePolicy(ctx workflow.Context, input *secretsmanager.GetResourcePolicyInput) (*secretsmanager.GetResourcePolicyOutput, error) {
    var output secretsmanager.GetResourcePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetResourcePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *SecretsManagerStub) GetResourcePolicyAsync(ctx workflow.Context, input *secretsmanager.GetResourcePolicyInput) *SecretsmanagerGetResourcePolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetResourcePolicy, input)
    return &SecretsmanagerGetResourcePolicyResult{Result: future}
}

func (a *SecretsManagerStub) GetSecretValue(ctx workflow.Context, input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error) {
    var output secretsmanager.GetSecretValueOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSecretValue, input).Get(ctx, &output)
    return &output, err
}

func (a *SecretsManagerStub) GetSecretValueAsync(ctx workflow.Context, input *secretsmanager.GetSecretValueInput) *SecretsmanagerGetSecretValueResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSecretValue, input)
    return &SecretsmanagerGetSecretValueResult{Result: future}
}

func (a *SecretsManagerStub) ListSecretVersionIds(ctx workflow.Context, input *secretsmanager.ListSecretVersionIdsInput) (*secretsmanager.ListSecretVersionIdsOutput, error) {
    var output secretsmanager.ListSecretVersionIdsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSecretVersionIds, input).Get(ctx, &output)
    return &output, err
}

func (a *SecretsManagerStub) ListSecretVersionIdsAsync(ctx workflow.Context, input *secretsmanager.ListSecretVersionIdsInput) *SecretsmanagerListSecretVersionIdsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListSecretVersionIds, input)
    return &SecretsmanagerListSecretVersionIdsResult{Result: future}
}

func (a *SecretsManagerStub) ListSecrets(ctx workflow.Context, input *secretsmanager.ListSecretsInput) (*secretsmanager.ListSecretsOutput, error) {
    var output secretsmanager.ListSecretsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSecrets, input).Get(ctx, &output)
    return &output, err
}

func (a *SecretsManagerStub) ListSecretsAsync(ctx workflow.Context, input *secretsmanager.ListSecretsInput) *SecretsmanagerListSecretsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListSecrets, input)
    return &SecretsmanagerListSecretsResult{Result: future}
}

func (a *SecretsManagerStub) PutResourcePolicy(ctx workflow.Context, input *secretsmanager.PutResourcePolicyInput) (*secretsmanager.PutResourcePolicyOutput, error) {
    var output secretsmanager.PutResourcePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutResourcePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *SecretsManagerStub) PutResourcePolicyAsync(ctx workflow.Context, input *secretsmanager.PutResourcePolicyInput) *SecretsmanagerPutResourcePolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutResourcePolicy, input)
    return &SecretsmanagerPutResourcePolicyResult{Result: future}
}

func (a *SecretsManagerStub) PutSecretValue(ctx workflow.Context, input *secretsmanager.PutSecretValueInput) (*secretsmanager.PutSecretValueOutput, error) {
    var output secretsmanager.PutSecretValueOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutSecretValue, input).Get(ctx, &output)
    return &output, err
}

func (a *SecretsManagerStub) PutSecretValueAsync(ctx workflow.Context, input *secretsmanager.PutSecretValueInput) *SecretsmanagerPutSecretValueResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutSecretValue, input)
    return &SecretsmanagerPutSecretValueResult{Result: future}
}

func (a *SecretsManagerStub) RestoreSecret(ctx workflow.Context, input *secretsmanager.RestoreSecretInput) (*secretsmanager.RestoreSecretOutput, error) {
    var output secretsmanager.RestoreSecretOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RestoreSecret, input).Get(ctx, &output)
    return &output, err
}

func (a *SecretsManagerStub) RestoreSecretAsync(ctx workflow.Context, input *secretsmanager.RestoreSecretInput) *SecretsmanagerRestoreSecretResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RestoreSecret, input)
    return &SecretsmanagerRestoreSecretResult{Result: future}
}

func (a *SecretsManagerStub) RotateSecret(ctx workflow.Context, input *secretsmanager.RotateSecretInput) (*secretsmanager.RotateSecretOutput, error) {
    var output secretsmanager.RotateSecretOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RotateSecret, input).Get(ctx, &output)
    return &output, err
}

func (a *SecretsManagerStub) RotateSecretAsync(ctx workflow.Context, input *secretsmanager.RotateSecretInput) *SecretsmanagerRotateSecretResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RotateSecret, input)
    return &SecretsmanagerRotateSecretResult{Result: future}
}

func (a *SecretsManagerStub) TagResource(ctx workflow.Context, input *secretsmanager.TagResourceInput) (*secretsmanager.TagResourceOutput, error) {
    var output secretsmanager.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SecretsManagerStub) TagResourceAsync(ctx workflow.Context, input *secretsmanager.TagResourceInput) *SecretsmanagerTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &SecretsmanagerTagResourceResult{Result: future}
}

func (a *SecretsManagerStub) UntagResource(ctx workflow.Context, input *secretsmanager.UntagResourceInput) (*secretsmanager.UntagResourceOutput, error) {
    var output secretsmanager.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SecretsManagerStub) UntagResourceAsync(ctx workflow.Context, input *secretsmanager.UntagResourceInput) *SecretsmanagerUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &SecretsmanagerUntagResourceResult{Result: future}
}

func (a *SecretsManagerStub) UpdateSecret(ctx workflow.Context, input *secretsmanager.UpdateSecretInput) (*secretsmanager.UpdateSecretOutput, error) {
    var output secretsmanager.UpdateSecretOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSecret, input).Get(ctx, &output)
    return &output, err
}

func (a *SecretsManagerStub) UpdateSecretAsync(ctx workflow.Context, input *secretsmanager.UpdateSecretInput) *SecretsmanagerUpdateSecretResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateSecret, input)
    return &SecretsmanagerUpdateSecretResult{Result: future}
}

func (a *SecretsManagerStub) UpdateSecretVersionStage(ctx workflow.Context, input *secretsmanager.UpdateSecretVersionStageInput) (*secretsmanager.UpdateSecretVersionStageOutput, error) {
    var output secretsmanager.UpdateSecretVersionStageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSecretVersionStage, input).Get(ctx, &output)
    return &output, err
}

func (a *SecretsManagerStub) UpdateSecretVersionStageAsync(ctx workflow.Context, input *secretsmanager.UpdateSecretVersionStageInput) *SecretsmanagerUpdateSecretVersionStageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateSecretVersionStage, input)
    return &SecretsmanagerUpdateSecretVersionStageResult{Result: future}
}

func (a *SecretsManagerStub) ValidateResourcePolicy(ctx workflow.Context, input *secretsmanager.ValidateResourcePolicyInput) (*secretsmanager.ValidateResourcePolicyOutput, error) {
    var output secretsmanager.ValidateResourcePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ValidateResourcePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *SecretsManagerStub) ValidateResourcePolicyAsync(ctx workflow.Context, input *secretsmanager.ValidateResourcePolicyInput) *SecretsmanagerValidateResourcePolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ValidateResourcePolicy, input)
    return &SecretsmanagerValidateResourcePolicyResult{Result: future}
}
