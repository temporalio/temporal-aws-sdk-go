package awsclients

import (
	"github.com/aws/aws-sdk-go/service/sso"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type SSOClient interface {
	GetRoleCredentials(ctx workflow.Context, input *sso.GetRoleCredentialsInput) (*sso.GetRoleCredentialsOutput, error)
	GetRoleCredentialsAsync(ctx workflow.Context, input *sso.GetRoleCredentialsInput) *SsoGetRoleCredentialsResult

	ListAccountRoles(ctx workflow.Context, input *sso.ListAccountRolesInput) (*sso.ListAccountRolesOutput, error)
	ListAccountRolesAsync(ctx workflow.Context, input *sso.ListAccountRolesInput) *SsoListAccountRolesResult

	ListAccounts(ctx workflow.Context, input *sso.ListAccountsInput) (*sso.ListAccountsOutput, error)
	ListAccountsAsync(ctx workflow.Context, input *sso.ListAccountsInput) *SsoListAccountsResult

	Logout(ctx workflow.Context, input *sso.LogoutInput) (*sso.LogoutOutput, error)
	LogoutAsync(ctx workflow.Context, input *sso.LogoutInput) *SsoLogoutResult
}

type SsoGetRoleCredentialsResult struct {
	Result workflow.Future
}

func (r *SsoGetRoleCredentialsResult) Get(ctx workflow.Context) (*sso.GetRoleCredentialsOutput, error) {
	var output sso.GetRoleCredentialsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SsoListAccountRolesResult struct {
	Result workflow.Future
}

func (r *SsoListAccountRolesResult) Get(ctx workflow.Context) (*sso.ListAccountRolesOutput, error) {
	var output sso.ListAccountRolesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SsoListAccountsResult struct {
	Result workflow.Future
}

func (r *SsoListAccountsResult) Get(ctx workflow.Context) (*sso.ListAccountsOutput, error) {
	var output sso.ListAccountsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SsoLogoutResult struct {
	Result workflow.Future
}

func (r *SsoLogoutResult) Get(ctx workflow.Context) (*sso.LogoutOutput, error) {
	var output sso.LogoutOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SSOStub struct {
	activities awsactivities.SSOActivities
}

func NewSSOStub() SSOClient {
	return &SSOStub{}
}

func (a *SSOStub) GetRoleCredentials(ctx workflow.Context, input *sso.GetRoleCredentialsInput) (*sso.GetRoleCredentialsOutput, error) {
	var output sso.GetRoleCredentialsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRoleCredentials, input).Get(ctx, &output)
	return &output, err
}

func (a *SSOStub) GetRoleCredentialsAsync(ctx workflow.Context, input *sso.GetRoleCredentialsInput) *SsoGetRoleCredentialsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRoleCredentials, input)
	return &SsoGetRoleCredentialsResult{Result: future}
}

func (a *SSOStub) ListAccountRoles(ctx workflow.Context, input *sso.ListAccountRolesInput) (*sso.ListAccountRolesOutput, error) {
	var output sso.ListAccountRolesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListAccountRoles, input).Get(ctx, &output)
	return &output, err
}

func (a *SSOStub) ListAccountRolesAsync(ctx workflow.Context, input *sso.ListAccountRolesInput) *SsoListAccountRolesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListAccountRoles, input)
	return &SsoListAccountRolesResult{Result: future}
}

func (a *SSOStub) ListAccounts(ctx workflow.Context, input *sso.ListAccountsInput) (*sso.ListAccountsOutput, error) {
	var output sso.ListAccountsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListAccounts, input).Get(ctx, &output)
	return &output, err
}

func (a *SSOStub) ListAccountsAsync(ctx workflow.Context, input *sso.ListAccountsInput) *SsoListAccountsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListAccounts, input)
	return &SsoListAccountsResult{Result: future}
}

func (a *SSOStub) Logout(ctx workflow.Context, input *sso.LogoutInput) (*sso.LogoutOutput, error) {
	var output sso.LogoutOutput
	err := workflow.ExecuteActivity(ctx, a.activities.Logout, input).Get(ctx, &output)
	return &output, err
}

func (a *SSOStub) LogoutAsync(ctx workflow.Context, input *sso.LogoutInput) *SsoLogoutResult {
	future := workflow.ExecuteActivity(ctx, a.activities.Logout, input)
	return &SsoLogoutResult{Result: future}
}
