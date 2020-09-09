package awsclients

import (
	"github.com/aws/aws-sdk-go/service/transfer"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type TransferClient interface {
       CreateServer(ctx workflow.Context, input *transfer.CreateServerInput) (*transfer.CreateServerOutput, error)
       CreateServerAsync(ctx workflow.Context, input *transfer.CreateServerInput) *TransferCreateServerResult

       CreateUser(ctx workflow.Context, input *transfer.CreateUserInput) (*transfer.CreateUserOutput, error)
       CreateUserAsync(ctx workflow.Context, input *transfer.CreateUserInput) *TransferCreateUserResult

       DeleteServer(ctx workflow.Context, input *transfer.DeleteServerInput) (*transfer.DeleteServerOutput, error)
       DeleteServerAsync(ctx workflow.Context, input *transfer.DeleteServerInput) *TransferDeleteServerResult

       DeleteSshPublicKey(ctx workflow.Context, input *transfer.DeleteSshPublicKeyInput) (*transfer.DeleteSshPublicKeyOutput, error)
       DeleteSshPublicKeyAsync(ctx workflow.Context, input *transfer.DeleteSshPublicKeyInput) *TransferDeleteSshPublicKeyResult

       DeleteUser(ctx workflow.Context, input *transfer.DeleteUserInput) (*transfer.DeleteUserOutput, error)
       DeleteUserAsync(ctx workflow.Context, input *transfer.DeleteUserInput) *TransferDeleteUserResult

       DescribeSecurityPolicy(ctx workflow.Context, input *transfer.DescribeSecurityPolicyInput) (*transfer.DescribeSecurityPolicyOutput, error)
       DescribeSecurityPolicyAsync(ctx workflow.Context, input *transfer.DescribeSecurityPolicyInput) *TransferDescribeSecurityPolicyResult

       DescribeServer(ctx workflow.Context, input *transfer.DescribeServerInput) (*transfer.DescribeServerOutput, error)
       DescribeServerAsync(ctx workflow.Context, input *transfer.DescribeServerInput) *TransferDescribeServerResult

       DescribeUser(ctx workflow.Context, input *transfer.DescribeUserInput) (*transfer.DescribeUserOutput, error)
       DescribeUserAsync(ctx workflow.Context, input *transfer.DescribeUserInput) *TransferDescribeUserResult

       ImportSshPublicKey(ctx workflow.Context, input *transfer.ImportSshPublicKeyInput) (*transfer.ImportSshPublicKeyOutput, error)
       ImportSshPublicKeyAsync(ctx workflow.Context, input *transfer.ImportSshPublicKeyInput) *TransferImportSshPublicKeyResult

       ListSecurityPolicies(ctx workflow.Context, input *transfer.ListSecurityPoliciesInput) (*transfer.ListSecurityPoliciesOutput, error)
       ListSecurityPoliciesAsync(ctx workflow.Context, input *transfer.ListSecurityPoliciesInput) *TransferListSecurityPoliciesResult

       ListServers(ctx workflow.Context, input *transfer.ListServersInput) (*transfer.ListServersOutput, error)
       ListServersAsync(ctx workflow.Context, input *transfer.ListServersInput) *TransferListServersResult

       ListTagsForResource(ctx workflow.Context, input *transfer.ListTagsForResourceInput) (*transfer.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *transfer.ListTagsForResourceInput) *TransferListTagsForResourceResult

       ListUsers(ctx workflow.Context, input *transfer.ListUsersInput) (*transfer.ListUsersOutput, error)
       ListUsersAsync(ctx workflow.Context, input *transfer.ListUsersInput) *TransferListUsersResult

       StartServer(ctx workflow.Context, input *transfer.StartServerInput) (*transfer.StartServerOutput, error)
       StartServerAsync(ctx workflow.Context, input *transfer.StartServerInput) *TransferStartServerResult

       StopServer(ctx workflow.Context, input *transfer.StopServerInput) (*transfer.StopServerOutput, error)
       StopServerAsync(ctx workflow.Context, input *transfer.StopServerInput) *TransferStopServerResult

       TagResource(ctx workflow.Context, input *transfer.TagResourceInput) (*transfer.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *transfer.TagResourceInput) *TransferTagResourceResult

       TestIdentityProvider(ctx workflow.Context, input *transfer.TestIdentityProviderInput) (*transfer.TestIdentityProviderOutput, error)
       TestIdentityProviderAsync(ctx workflow.Context, input *transfer.TestIdentityProviderInput) *TransferTestIdentityProviderResult

       UntagResource(ctx workflow.Context, input *transfer.UntagResourceInput) (*transfer.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *transfer.UntagResourceInput) *TransferUntagResourceResult

       UpdateServer(ctx workflow.Context, input *transfer.UpdateServerInput) (*transfer.UpdateServerOutput, error)
       UpdateServerAsync(ctx workflow.Context, input *transfer.UpdateServerInput) *TransferUpdateServerResult

       UpdateUser(ctx workflow.Context, input *transfer.UpdateUserInput) (*transfer.UpdateUserOutput, error)
       UpdateUserAsync(ctx workflow.Context, input *transfer.UpdateUserInput) *TransferUpdateUserResult
}

type TransferCreateServerResult struct {
	Result workflow.Future
}

func (r *TransferCreateServerResult) Get(ctx workflow.Context) (*transfer.CreateServerOutput, error) {
    var output transfer.CreateServerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TransferCreateUserResult struct {
	Result workflow.Future
}

func (r *TransferCreateUserResult) Get(ctx workflow.Context) (*transfer.CreateUserOutput, error) {
    var output transfer.CreateUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TransferDeleteServerResult struct {
	Result workflow.Future
}

func (r *TransferDeleteServerResult) Get(ctx workflow.Context) (*transfer.DeleteServerOutput, error) {
    var output transfer.DeleteServerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TransferDeleteSshPublicKeyResult struct {
	Result workflow.Future
}

func (r *TransferDeleteSshPublicKeyResult) Get(ctx workflow.Context) (*transfer.DeleteSshPublicKeyOutput, error) {
    var output transfer.DeleteSshPublicKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TransferDeleteUserResult struct {
	Result workflow.Future
}

func (r *TransferDeleteUserResult) Get(ctx workflow.Context) (*transfer.DeleteUserOutput, error) {
    var output transfer.DeleteUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TransferDescribeSecurityPolicyResult struct {
	Result workflow.Future
}

func (r *TransferDescribeSecurityPolicyResult) Get(ctx workflow.Context) (*transfer.DescribeSecurityPolicyOutput, error) {
    var output transfer.DescribeSecurityPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TransferDescribeServerResult struct {
	Result workflow.Future
}

func (r *TransferDescribeServerResult) Get(ctx workflow.Context) (*transfer.DescribeServerOutput, error) {
    var output transfer.DescribeServerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TransferDescribeUserResult struct {
	Result workflow.Future
}

func (r *TransferDescribeUserResult) Get(ctx workflow.Context) (*transfer.DescribeUserOutput, error) {
    var output transfer.DescribeUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TransferImportSshPublicKeyResult struct {
	Result workflow.Future
}

func (r *TransferImportSshPublicKeyResult) Get(ctx workflow.Context) (*transfer.ImportSshPublicKeyOutput, error) {
    var output transfer.ImportSshPublicKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TransferListSecurityPoliciesResult struct {
	Result workflow.Future
}

func (r *TransferListSecurityPoliciesResult) Get(ctx workflow.Context) (*transfer.ListSecurityPoliciesOutput, error) {
    var output transfer.ListSecurityPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TransferListServersResult struct {
	Result workflow.Future
}

func (r *TransferListServersResult) Get(ctx workflow.Context) (*transfer.ListServersOutput, error) {
    var output transfer.ListServersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TransferListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *TransferListTagsForResourceResult) Get(ctx workflow.Context) (*transfer.ListTagsForResourceOutput, error) {
    var output transfer.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TransferListUsersResult struct {
	Result workflow.Future
}

func (r *TransferListUsersResult) Get(ctx workflow.Context) (*transfer.ListUsersOutput, error) {
    var output transfer.ListUsersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TransferStartServerResult struct {
	Result workflow.Future
}

func (r *TransferStartServerResult) Get(ctx workflow.Context) (*transfer.StartServerOutput, error) {
    var output transfer.StartServerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TransferStopServerResult struct {
	Result workflow.Future
}

func (r *TransferStopServerResult) Get(ctx workflow.Context) (*transfer.StopServerOutput, error) {
    var output transfer.StopServerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TransferTagResourceResult struct {
	Result workflow.Future
}

func (r *TransferTagResourceResult) Get(ctx workflow.Context) (*transfer.TagResourceOutput, error) {
    var output transfer.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TransferTestIdentityProviderResult struct {
	Result workflow.Future
}

func (r *TransferTestIdentityProviderResult) Get(ctx workflow.Context) (*transfer.TestIdentityProviderOutput, error) {
    var output transfer.TestIdentityProviderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TransferUntagResourceResult struct {
	Result workflow.Future
}

func (r *TransferUntagResourceResult) Get(ctx workflow.Context) (*transfer.UntagResourceOutput, error) {
    var output transfer.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TransferUpdateServerResult struct {
	Result workflow.Future
}

func (r *TransferUpdateServerResult) Get(ctx workflow.Context) (*transfer.UpdateServerOutput, error) {
    var output transfer.UpdateServerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TransferUpdateUserResult struct {
	Result workflow.Future
}

func (r *TransferUpdateUserResult) Get(ctx workflow.Context) (*transfer.UpdateUserOutput, error) {
    var output transfer.UpdateUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TransferStub struct {
    activities awsactivities.TransferActivities
}

func NewTransferStub() TransferClient {
    return &TransferStub{}
}

func (a *TransferStub) CreateServer(ctx workflow.Context, input *transfer.CreateServerInput) (*transfer.CreateServerOutput, error) {
    var output transfer.CreateServerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateServer, input).Get(ctx, &output)
    return &output, err
}

func (a *TransferStub) CreateServerAsync(ctx workflow.Context, input *transfer.CreateServerInput) *TransferCreateServerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateServer, input)
    return &TransferCreateServerResult{Result: future}
}

func (a *TransferStub) CreateUser(ctx workflow.Context, input *transfer.CreateUserInput) (*transfer.CreateUserOutput, error) {
    var output transfer.CreateUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateUser, input).Get(ctx, &output)
    return &output, err
}

func (a *TransferStub) CreateUserAsync(ctx workflow.Context, input *transfer.CreateUserInput) *TransferCreateUserResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateUser, input)
    return &TransferCreateUserResult{Result: future}
}

func (a *TransferStub) DeleteServer(ctx workflow.Context, input *transfer.DeleteServerInput) (*transfer.DeleteServerOutput, error) {
    var output transfer.DeleteServerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteServer, input).Get(ctx, &output)
    return &output, err
}

func (a *TransferStub) DeleteServerAsync(ctx workflow.Context, input *transfer.DeleteServerInput) *TransferDeleteServerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteServer, input)
    return &TransferDeleteServerResult{Result: future}
}

func (a *TransferStub) DeleteSshPublicKey(ctx workflow.Context, input *transfer.DeleteSshPublicKeyInput) (*transfer.DeleteSshPublicKeyOutput, error) {
    var output transfer.DeleteSshPublicKeyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSshPublicKey, input).Get(ctx, &output)
    return &output, err
}

func (a *TransferStub) DeleteSshPublicKeyAsync(ctx workflow.Context, input *transfer.DeleteSshPublicKeyInput) *TransferDeleteSshPublicKeyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSshPublicKey, input)
    return &TransferDeleteSshPublicKeyResult{Result: future}
}

func (a *TransferStub) DeleteUser(ctx workflow.Context, input *transfer.DeleteUserInput) (*transfer.DeleteUserOutput, error) {
    var output transfer.DeleteUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteUser, input).Get(ctx, &output)
    return &output, err
}

func (a *TransferStub) DeleteUserAsync(ctx workflow.Context, input *transfer.DeleteUserInput) *TransferDeleteUserResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteUser, input)
    return &TransferDeleteUserResult{Result: future}
}

func (a *TransferStub) DescribeSecurityPolicy(ctx workflow.Context, input *transfer.DescribeSecurityPolicyInput) (*transfer.DescribeSecurityPolicyOutput, error) {
    var output transfer.DescribeSecurityPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSecurityPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *TransferStub) DescribeSecurityPolicyAsync(ctx workflow.Context, input *transfer.DescribeSecurityPolicyInput) *TransferDescribeSecurityPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSecurityPolicy, input)
    return &TransferDescribeSecurityPolicyResult{Result: future}
}

func (a *TransferStub) DescribeServer(ctx workflow.Context, input *transfer.DescribeServerInput) (*transfer.DescribeServerOutput, error) {
    var output transfer.DescribeServerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeServer, input).Get(ctx, &output)
    return &output, err
}

func (a *TransferStub) DescribeServerAsync(ctx workflow.Context, input *transfer.DescribeServerInput) *TransferDescribeServerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeServer, input)
    return &TransferDescribeServerResult{Result: future}
}

func (a *TransferStub) DescribeUser(ctx workflow.Context, input *transfer.DescribeUserInput) (*transfer.DescribeUserOutput, error) {
    var output transfer.DescribeUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeUser, input).Get(ctx, &output)
    return &output, err
}

func (a *TransferStub) DescribeUserAsync(ctx workflow.Context, input *transfer.DescribeUserInput) *TransferDescribeUserResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeUser, input)
    return &TransferDescribeUserResult{Result: future}
}

func (a *TransferStub) ImportSshPublicKey(ctx workflow.Context, input *transfer.ImportSshPublicKeyInput) (*transfer.ImportSshPublicKeyOutput, error) {
    var output transfer.ImportSshPublicKeyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ImportSshPublicKey, input).Get(ctx, &output)
    return &output, err
}

func (a *TransferStub) ImportSshPublicKeyAsync(ctx workflow.Context, input *transfer.ImportSshPublicKeyInput) *TransferImportSshPublicKeyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ImportSshPublicKey, input)
    return &TransferImportSshPublicKeyResult{Result: future}
}

func (a *TransferStub) ListSecurityPolicies(ctx workflow.Context, input *transfer.ListSecurityPoliciesInput) (*transfer.ListSecurityPoliciesOutput, error) {
    var output transfer.ListSecurityPoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSecurityPolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *TransferStub) ListSecurityPoliciesAsync(ctx workflow.Context, input *transfer.ListSecurityPoliciesInput) *TransferListSecurityPoliciesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListSecurityPolicies, input)
    return &TransferListSecurityPoliciesResult{Result: future}
}

func (a *TransferStub) ListServers(ctx workflow.Context, input *transfer.ListServersInput) (*transfer.ListServersOutput, error) {
    var output transfer.ListServersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListServers, input).Get(ctx, &output)
    return &output, err
}

func (a *TransferStub) ListServersAsync(ctx workflow.Context, input *transfer.ListServersInput) *TransferListServersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListServers, input)
    return &TransferListServersResult{Result: future}
}

func (a *TransferStub) ListTagsForResource(ctx workflow.Context, input *transfer.ListTagsForResourceInput) (*transfer.ListTagsForResourceOutput, error) {
    var output transfer.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *TransferStub) ListTagsForResourceAsync(ctx workflow.Context, input *transfer.ListTagsForResourceInput) *TransferListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &TransferListTagsForResourceResult{Result: future}
}

func (a *TransferStub) ListUsers(ctx workflow.Context, input *transfer.ListUsersInput) (*transfer.ListUsersOutput, error) {
    var output transfer.ListUsersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListUsers, input).Get(ctx, &output)
    return &output, err
}

func (a *TransferStub) ListUsersAsync(ctx workflow.Context, input *transfer.ListUsersInput) *TransferListUsersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListUsers, input)
    return &TransferListUsersResult{Result: future}
}

func (a *TransferStub) StartServer(ctx workflow.Context, input *transfer.StartServerInput) (*transfer.StartServerOutput, error) {
    var output transfer.StartServerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartServer, input).Get(ctx, &output)
    return &output, err
}

func (a *TransferStub) StartServerAsync(ctx workflow.Context, input *transfer.StartServerInput) *TransferStartServerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartServer, input)
    return &TransferStartServerResult{Result: future}
}

func (a *TransferStub) StopServer(ctx workflow.Context, input *transfer.StopServerInput) (*transfer.StopServerOutput, error) {
    var output transfer.StopServerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopServer, input).Get(ctx, &output)
    return &output, err
}

func (a *TransferStub) StopServerAsync(ctx workflow.Context, input *transfer.StopServerInput) *TransferStopServerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopServer, input)
    return &TransferStopServerResult{Result: future}
}

func (a *TransferStub) TagResource(ctx workflow.Context, input *transfer.TagResourceInput) (*transfer.TagResourceOutput, error) {
    var output transfer.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *TransferStub) TagResourceAsync(ctx workflow.Context, input *transfer.TagResourceInput) *TransferTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &TransferTagResourceResult{Result: future}
}

func (a *TransferStub) TestIdentityProvider(ctx workflow.Context, input *transfer.TestIdentityProviderInput) (*transfer.TestIdentityProviderOutput, error) {
    var output transfer.TestIdentityProviderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TestIdentityProvider, input).Get(ctx, &output)
    return &output, err
}

func (a *TransferStub) TestIdentityProviderAsync(ctx workflow.Context, input *transfer.TestIdentityProviderInput) *TransferTestIdentityProviderResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TestIdentityProvider, input)
    return &TransferTestIdentityProviderResult{Result: future}
}

func (a *TransferStub) UntagResource(ctx workflow.Context, input *transfer.UntagResourceInput) (*transfer.UntagResourceOutput, error) {
    var output transfer.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *TransferStub) UntagResourceAsync(ctx workflow.Context, input *transfer.UntagResourceInput) *TransferUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &TransferUntagResourceResult{Result: future}
}

func (a *TransferStub) UpdateServer(ctx workflow.Context, input *transfer.UpdateServerInput) (*transfer.UpdateServerOutput, error) {
    var output transfer.UpdateServerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateServer, input).Get(ctx, &output)
    return &output, err
}

func (a *TransferStub) UpdateServerAsync(ctx workflow.Context, input *transfer.UpdateServerInput) *TransferUpdateServerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateServer, input)
    return &TransferUpdateServerResult{Result: future}
}

func (a *TransferStub) UpdateUser(ctx workflow.Context, input *transfer.UpdateUserInput) (*transfer.UpdateUserOutput, error) {
    var output transfer.UpdateUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateUser, input).Get(ctx, &output)
    return &output, err
}

func (a *TransferStub) UpdateUserAsync(ctx workflow.Context, input *transfer.UpdateUserInput) *TransferUpdateUserResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateUser, input)
    return &TransferUpdateUserResult{Result: future}
}
