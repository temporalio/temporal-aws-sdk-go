package awsclients

import (
	"github.com/aws/aws-sdk-go/service/ssooidc"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type SSOOIDCClient interface {
	CreateToken(ctx workflow.Context, input *ssooidc.CreateTokenInput) (*ssooidc.CreateTokenOutput, error)
	CreateTokenAsync(ctx workflow.Context, input *ssooidc.CreateTokenInput) *SsooidcCreateTokenResult

	RegisterClient(ctx workflow.Context, input *ssooidc.RegisterClientInput) (*ssooidc.RegisterClientOutput, error)
	RegisterClientAsync(ctx workflow.Context, input *ssooidc.RegisterClientInput) *SsooidcRegisterClientResult

	StartDeviceAuthorization(ctx workflow.Context, input *ssooidc.StartDeviceAuthorizationInput) (*ssooidc.StartDeviceAuthorizationOutput, error)
	StartDeviceAuthorizationAsync(ctx workflow.Context, input *ssooidc.StartDeviceAuthorizationInput) *SsooidcStartDeviceAuthorizationResult
}

type SsooidcCreateTokenResult struct {
	Result workflow.Future
}

func (r *SsooidcCreateTokenResult) Get(ctx workflow.Context) (*ssooidc.CreateTokenOutput, error) {
	var output ssooidc.CreateTokenOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SsooidcRegisterClientResult struct {
	Result workflow.Future
}

func (r *SsooidcRegisterClientResult) Get(ctx workflow.Context) (*ssooidc.RegisterClientOutput, error) {
	var output ssooidc.RegisterClientOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SsooidcStartDeviceAuthorizationResult struct {
	Result workflow.Future
}

func (r *SsooidcStartDeviceAuthorizationResult) Get(ctx workflow.Context) (*ssooidc.StartDeviceAuthorizationOutput, error) {
	var output ssooidc.StartDeviceAuthorizationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SSOOIDCStub struct {
	activities awsactivities.SSOOIDCActivities
}

func NewSSOOIDCStub() SSOOIDCClient {
	return &SSOOIDCStub{}
}

func (a *SSOOIDCStub) CreateToken(ctx workflow.Context, input *ssooidc.CreateTokenInput) (*ssooidc.CreateTokenOutput, error) {
	var output ssooidc.CreateTokenOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateToken, input).Get(ctx, &output)
	return &output, err
}

func (a *SSOOIDCStub) CreateTokenAsync(ctx workflow.Context, input *ssooidc.CreateTokenInput) *SsooidcCreateTokenResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateToken, input)
	return &SsooidcCreateTokenResult{Result: future}
}

func (a *SSOOIDCStub) RegisterClient(ctx workflow.Context, input *ssooidc.RegisterClientInput) (*ssooidc.RegisterClientOutput, error) {
	var output ssooidc.RegisterClientOutput
	err := workflow.ExecuteActivity(ctx, a.activities.RegisterClient, input).Get(ctx, &output)
	return &output, err
}

func (a *SSOOIDCStub) RegisterClientAsync(ctx workflow.Context, input *ssooidc.RegisterClientInput) *SsooidcRegisterClientResult {
	future := workflow.ExecuteActivity(ctx, a.activities.RegisterClient, input)
	return &SsooidcRegisterClientResult{Result: future}
}

func (a *SSOOIDCStub) StartDeviceAuthorization(ctx workflow.Context, input *ssooidc.StartDeviceAuthorizationInput) (*ssooidc.StartDeviceAuthorizationOutput, error) {
	var output ssooidc.StartDeviceAuthorizationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StartDeviceAuthorization, input).Get(ctx, &output)
	return &output, err
}

func (a *SSOOIDCStub) StartDeviceAuthorizationAsync(ctx workflow.Context, input *ssooidc.StartDeviceAuthorizationInput) *SsooidcStartDeviceAuthorizationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StartDeviceAuthorization, input)
	return &SsooidcStartDeviceAuthorizationResult{Result: future}
}
