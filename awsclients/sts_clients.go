package awsclients

import (
	"github.com/aws/aws-sdk-go/service/sts"
	"go.temporal.io/sdk/workflow"
)

type STSClient interface {
       AssumeRole(ctx workflow.Context, input *sts.AssumeRoleInput) (*sts.AssumeRoleOutput, error)
       AssumeRoleAsync(ctx workflow.Context, input *sts.AssumeRoleInput) *StsAssumeRoleResult

       AssumeRoleWithSAML(ctx workflow.Context, input *sts.AssumeRoleWithSAMLInput) (*sts.AssumeRoleWithSAMLOutput, error)
       AssumeRoleWithSAMLAsync(ctx workflow.Context, input *sts.AssumeRoleWithSAMLInput) *StsAssumeRoleWithSAMLResult

       AssumeRoleWithWebIdentity(ctx workflow.Context, input *sts.AssumeRoleWithWebIdentityInput) (*sts.AssumeRoleWithWebIdentityOutput, error)
       AssumeRoleWithWebIdentityAsync(ctx workflow.Context, input *sts.AssumeRoleWithWebIdentityInput) *StsAssumeRoleWithWebIdentityResult

       DecodeAuthorizationMessage(ctx workflow.Context, input *sts.DecodeAuthorizationMessageInput) (*sts.DecodeAuthorizationMessageOutput, error)
       DecodeAuthorizationMessageAsync(ctx workflow.Context, input *sts.DecodeAuthorizationMessageInput) *StsDecodeAuthorizationMessageResult

       GetAccessKeyInfo(ctx workflow.Context, input *sts.GetAccessKeyInfoInput) (*sts.GetAccessKeyInfoOutput, error)
       GetAccessKeyInfoAsync(ctx workflow.Context, input *sts.GetAccessKeyInfoInput) *StsGetAccessKeyInfoResult

       GetCallerIdentity(ctx workflow.Context, input *sts.GetCallerIdentityInput) (*sts.GetCallerIdentityOutput, error)
       GetCallerIdentityAsync(ctx workflow.Context, input *sts.GetCallerIdentityInput) *StsGetCallerIdentityResult

       GetFederationToken(ctx workflow.Context, input *sts.GetFederationTokenInput) (*sts.GetFederationTokenOutput, error)
       GetFederationTokenAsync(ctx workflow.Context, input *sts.GetFederationTokenInput) *StsGetFederationTokenResult

       GetSessionToken(ctx workflow.Context, input *sts.GetSessionTokenInput) (*sts.GetSessionTokenOutput, error)
       GetSessionTokenAsync(ctx workflow.Context, input *sts.GetSessionTokenInput) *StsGetSessionTokenResult
}

type STSStub struct {
}

func NewSTSStub() STSClient {
    return &STSStub{}
}

type StsAssumeRoleResult struct {
	Result workflow.Future
}

func (r *StsAssumeRoleResult) Get(ctx workflow.Context) (*sts.AssumeRoleOutput, error) {
    var output sts.AssumeRoleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type StsAssumeRoleWithSAMLResult struct {
	Result workflow.Future
}

func (r *StsAssumeRoleWithSAMLResult) Get(ctx workflow.Context) (*sts.AssumeRoleWithSAMLOutput, error) {
    var output sts.AssumeRoleWithSAMLOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type StsAssumeRoleWithWebIdentityResult struct {
	Result workflow.Future
}

func (r *StsAssumeRoleWithWebIdentityResult) Get(ctx workflow.Context) (*sts.AssumeRoleWithWebIdentityOutput, error) {
    var output sts.AssumeRoleWithWebIdentityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type StsDecodeAuthorizationMessageResult struct {
	Result workflow.Future
}

func (r *StsDecodeAuthorizationMessageResult) Get(ctx workflow.Context) (*sts.DecodeAuthorizationMessageOutput, error) {
    var output sts.DecodeAuthorizationMessageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type StsGetAccessKeyInfoResult struct {
	Result workflow.Future
}

func (r *StsGetAccessKeyInfoResult) Get(ctx workflow.Context) (*sts.GetAccessKeyInfoOutput, error) {
    var output sts.GetAccessKeyInfoOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type StsGetCallerIdentityResult struct {
	Result workflow.Future
}

func (r *StsGetCallerIdentityResult) Get(ctx workflow.Context) (*sts.GetCallerIdentityOutput, error) {
    var output sts.GetCallerIdentityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type StsGetFederationTokenResult struct {
	Result workflow.Future
}

func (r *StsGetFederationTokenResult) Get(ctx workflow.Context) (*sts.GetFederationTokenOutput, error) {
    var output sts.GetFederationTokenOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type StsGetSessionTokenResult struct {
	Result workflow.Future
}

func (r *StsGetSessionTokenResult) Get(ctx workflow.Context) (*sts.GetSessionTokenOutput, error) {
    var output sts.GetSessionTokenOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *STSStub) AssumeRole(ctx workflow.Context, input *sts.AssumeRoleInput) (*sts.AssumeRoleOutput, error) {
    var output sts.AssumeRoleOutput
    err := workflow.ExecuteActivity(ctx, "STS.AssumeRole", input).Get(ctx, &output)
    return &output, err
}

func (a *STSStub) AssumeRoleAsync(ctx workflow.Context, input *sts.AssumeRoleInput) *StsAssumeRoleResult {
    future := workflow.ExecuteActivity(ctx, "STS.AssumeRole", input)
    return &StsAssumeRoleResult{Result: future}
}

func (a *STSStub) AssumeRoleWithSAML(ctx workflow.Context, input *sts.AssumeRoleWithSAMLInput) (*sts.AssumeRoleWithSAMLOutput, error) {
    var output sts.AssumeRoleWithSAMLOutput
    err := workflow.ExecuteActivity(ctx, "STS.AssumeRoleWithSAML", input).Get(ctx, &output)
    return &output, err
}

func (a *STSStub) AssumeRoleWithSAMLAsync(ctx workflow.Context, input *sts.AssumeRoleWithSAMLInput) *StsAssumeRoleWithSAMLResult {
    future := workflow.ExecuteActivity(ctx, "STS.AssumeRoleWithSAML", input)
    return &StsAssumeRoleWithSAMLResult{Result: future}
}

func (a *STSStub) AssumeRoleWithWebIdentity(ctx workflow.Context, input *sts.AssumeRoleWithWebIdentityInput) (*sts.AssumeRoleWithWebIdentityOutput, error) {
    var output sts.AssumeRoleWithWebIdentityOutput
    err := workflow.ExecuteActivity(ctx, "STS.AssumeRoleWithWebIdentity", input).Get(ctx, &output)
    return &output, err
}

func (a *STSStub) AssumeRoleWithWebIdentityAsync(ctx workflow.Context, input *sts.AssumeRoleWithWebIdentityInput) *StsAssumeRoleWithWebIdentityResult {
    future := workflow.ExecuteActivity(ctx, "STS.AssumeRoleWithWebIdentity", input)
    return &StsAssumeRoleWithWebIdentityResult{Result: future}
}

func (a *STSStub) DecodeAuthorizationMessage(ctx workflow.Context, input *sts.DecodeAuthorizationMessageInput) (*sts.DecodeAuthorizationMessageOutput, error) {
    var output sts.DecodeAuthorizationMessageOutput
    err := workflow.ExecuteActivity(ctx, "STS.DecodeAuthorizationMessage", input).Get(ctx, &output)
    return &output, err
}

func (a *STSStub) DecodeAuthorizationMessageAsync(ctx workflow.Context, input *sts.DecodeAuthorizationMessageInput) *StsDecodeAuthorizationMessageResult {
    future := workflow.ExecuteActivity(ctx, "STS.DecodeAuthorizationMessage", input)
    return &StsDecodeAuthorizationMessageResult{Result: future}
}

func (a *STSStub) GetAccessKeyInfo(ctx workflow.Context, input *sts.GetAccessKeyInfoInput) (*sts.GetAccessKeyInfoOutput, error) {
    var output sts.GetAccessKeyInfoOutput
    err := workflow.ExecuteActivity(ctx, "STS.GetAccessKeyInfo", input).Get(ctx, &output)
    return &output, err
}

func (a *STSStub) GetAccessKeyInfoAsync(ctx workflow.Context, input *sts.GetAccessKeyInfoInput) *StsGetAccessKeyInfoResult {
    future := workflow.ExecuteActivity(ctx, "STS.GetAccessKeyInfo", input)
    return &StsGetAccessKeyInfoResult{Result: future}
}

func (a *STSStub) GetCallerIdentity(ctx workflow.Context, input *sts.GetCallerIdentityInput) (*sts.GetCallerIdentityOutput, error) {
    var output sts.GetCallerIdentityOutput
    err := workflow.ExecuteActivity(ctx, "STS.GetCallerIdentity", input).Get(ctx, &output)
    return &output, err
}

func (a *STSStub) GetCallerIdentityAsync(ctx workflow.Context, input *sts.GetCallerIdentityInput) *StsGetCallerIdentityResult {
    future := workflow.ExecuteActivity(ctx, "STS.GetCallerIdentity", input)
    return &StsGetCallerIdentityResult{Result: future}
}

func (a *STSStub) GetFederationToken(ctx workflow.Context, input *sts.GetFederationTokenInput) (*sts.GetFederationTokenOutput, error) {
    var output sts.GetFederationTokenOutput
    err := workflow.ExecuteActivity(ctx, "STS.GetFederationToken", input).Get(ctx, &output)
    return &output, err
}

func (a *STSStub) GetFederationTokenAsync(ctx workflow.Context, input *sts.GetFederationTokenInput) *StsGetFederationTokenResult {
    future := workflow.ExecuteActivity(ctx, "STS.GetFederationToken", input)
    return &StsGetFederationTokenResult{Result: future}
}

func (a *STSStub) GetSessionToken(ctx workflow.Context, input *sts.GetSessionTokenInput) (*sts.GetSessionTokenOutput, error) {
    var output sts.GetSessionTokenOutput
    err := workflow.ExecuteActivity(ctx, "STS.GetSessionToken", input).Get(ctx, &output)
    return &output, err
}

func (a *STSStub) GetSessionTokenAsync(ctx workflow.Context, input *sts.GetSessionTokenInput) *StsGetSessionTokenResult {
    future := workflow.ExecuteActivity(ctx, "STS.GetSessionToken", input)
    return &StsGetSessionTokenResult{Result: future}
}
