// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package transferstub

import (
	"github.com/aws/aws-sdk-go/service/transfer"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CreateServerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateServerFuture) Get(ctx workflow.Context) (*transfer.CreateServerOutput, error) {
	var output transfer.CreateServerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateUserFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateUserFuture) Get(ctx workflow.Context) (*transfer.CreateUserOutput, error) {
	var output transfer.CreateUserOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteServerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteServerFuture) Get(ctx workflow.Context) (*transfer.DeleteServerOutput, error) {
	var output transfer.DeleteServerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteSshPublicKeyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteSshPublicKeyFuture) Get(ctx workflow.Context) (*transfer.DeleteSshPublicKeyOutput, error) {
	var output transfer.DeleteSshPublicKeyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteUserFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteUserFuture) Get(ctx workflow.Context) (*transfer.DeleteUserOutput, error) {
	var output transfer.DeleteUserOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeSecurityPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeSecurityPolicyFuture) Get(ctx workflow.Context) (*transfer.DescribeSecurityPolicyOutput, error) {
	var output transfer.DescribeSecurityPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeServerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeServerFuture) Get(ctx workflow.Context) (*transfer.DescribeServerOutput, error) {
	var output transfer.DescribeServerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeUserFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeUserFuture) Get(ctx workflow.Context) (*transfer.DescribeUserOutput, error) {
	var output transfer.DescribeUserOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ImportSshPublicKeyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ImportSshPublicKeyFuture) Get(ctx workflow.Context) (*transfer.ImportSshPublicKeyOutput, error) {
	var output transfer.ImportSshPublicKeyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListSecurityPoliciesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListSecurityPoliciesFuture) Get(ctx workflow.Context) (*transfer.ListSecurityPoliciesOutput, error) {
	var output transfer.ListSecurityPoliciesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListServersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListServersFuture) Get(ctx workflow.Context) (*transfer.ListServersOutput, error) {
	var output transfer.ListServersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*transfer.ListTagsForResourceOutput, error) {
	var output transfer.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListUsersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListUsersFuture) Get(ctx workflow.Context) (*transfer.ListUsersOutput, error) {
	var output transfer.ListUsersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StartServerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StartServerFuture) Get(ctx workflow.Context) (*transfer.StartServerOutput, error) {
	var output transfer.StartServerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StopServerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StopServerFuture) Get(ctx workflow.Context) (*transfer.StopServerOutput, error) {
	var output transfer.StopServerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*transfer.TagResourceOutput, error) {
	var output transfer.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TestIdentityProviderFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TestIdentityProviderFuture) Get(ctx workflow.Context) (*transfer.TestIdentityProviderOutput, error) {
	var output transfer.TestIdentityProviderOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*transfer.UntagResourceOutput, error) {
	var output transfer.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateServerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateServerFuture) Get(ctx workflow.Context) (*transfer.UpdateServerOutput, error) {
	var output transfer.UpdateServerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateUserFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateUserFuture) Get(ctx workflow.Context) (*transfer.UpdateUserOutput, error) {
	var output transfer.UpdateUserOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateServer(ctx workflow.Context, input *transfer.CreateServerInput) (*transfer.CreateServerOutput, error) {
	var output transfer.CreateServerOutput
	err := workflow.ExecuteActivity(ctx, "aws.transfer.CreateServer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateServerAsync(ctx workflow.Context, input *transfer.CreateServerInput) *CreateServerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transfer.CreateServer", input)
	return &CreateServerFuture{Future: future}
}

func (a *stub) CreateUser(ctx workflow.Context, input *transfer.CreateUserInput) (*transfer.CreateUserOutput, error) {
	var output transfer.CreateUserOutput
	err := workflow.ExecuteActivity(ctx, "aws.transfer.CreateUser", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateUserAsync(ctx workflow.Context, input *transfer.CreateUserInput) *CreateUserFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transfer.CreateUser", input)
	return &CreateUserFuture{Future: future}
}

func (a *stub) DeleteServer(ctx workflow.Context, input *transfer.DeleteServerInput) (*transfer.DeleteServerOutput, error) {
	var output transfer.DeleteServerOutput
	err := workflow.ExecuteActivity(ctx, "aws.transfer.DeleteServer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteServerAsync(ctx workflow.Context, input *transfer.DeleteServerInput) *DeleteServerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transfer.DeleteServer", input)
	return &DeleteServerFuture{Future: future}
}

func (a *stub) DeleteSshPublicKey(ctx workflow.Context, input *transfer.DeleteSshPublicKeyInput) (*transfer.DeleteSshPublicKeyOutput, error) {
	var output transfer.DeleteSshPublicKeyOutput
	err := workflow.ExecuteActivity(ctx, "aws.transfer.DeleteSshPublicKey", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteSshPublicKeyAsync(ctx workflow.Context, input *transfer.DeleteSshPublicKeyInput) *DeleteSshPublicKeyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transfer.DeleteSshPublicKey", input)
	return &DeleteSshPublicKeyFuture{Future: future}
}

func (a *stub) DeleteUser(ctx workflow.Context, input *transfer.DeleteUserInput) (*transfer.DeleteUserOutput, error) {
	var output transfer.DeleteUserOutput
	err := workflow.ExecuteActivity(ctx, "aws.transfer.DeleteUser", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteUserAsync(ctx workflow.Context, input *transfer.DeleteUserInput) *DeleteUserFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transfer.DeleteUser", input)
	return &DeleteUserFuture{Future: future}
}

func (a *stub) DescribeSecurityPolicy(ctx workflow.Context, input *transfer.DescribeSecurityPolicyInput) (*transfer.DescribeSecurityPolicyOutput, error) {
	var output transfer.DescribeSecurityPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.transfer.DescribeSecurityPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeSecurityPolicyAsync(ctx workflow.Context, input *transfer.DescribeSecurityPolicyInput) *DescribeSecurityPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transfer.DescribeSecurityPolicy", input)
	return &DescribeSecurityPolicyFuture{Future: future}
}

func (a *stub) DescribeServer(ctx workflow.Context, input *transfer.DescribeServerInput) (*transfer.DescribeServerOutput, error) {
	var output transfer.DescribeServerOutput
	err := workflow.ExecuteActivity(ctx, "aws.transfer.DescribeServer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeServerAsync(ctx workflow.Context, input *transfer.DescribeServerInput) *DescribeServerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transfer.DescribeServer", input)
	return &DescribeServerFuture{Future: future}
}

func (a *stub) DescribeUser(ctx workflow.Context, input *transfer.DescribeUserInput) (*transfer.DescribeUserOutput, error) {
	var output transfer.DescribeUserOutput
	err := workflow.ExecuteActivity(ctx, "aws.transfer.DescribeUser", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeUserAsync(ctx workflow.Context, input *transfer.DescribeUserInput) *DescribeUserFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transfer.DescribeUser", input)
	return &DescribeUserFuture{Future: future}
}

func (a *stub) ImportSshPublicKey(ctx workflow.Context, input *transfer.ImportSshPublicKeyInput) (*transfer.ImportSshPublicKeyOutput, error) {
	var output transfer.ImportSshPublicKeyOutput
	err := workflow.ExecuteActivity(ctx, "aws.transfer.ImportSshPublicKey", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ImportSshPublicKeyAsync(ctx workflow.Context, input *transfer.ImportSshPublicKeyInput) *ImportSshPublicKeyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transfer.ImportSshPublicKey", input)
	return &ImportSshPublicKeyFuture{Future: future}
}

func (a *stub) ListSecurityPolicies(ctx workflow.Context, input *transfer.ListSecurityPoliciesInput) (*transfer.ListSecurityPoliciesOutput, error) {
	var output transfer.ListSecurityPoliciesOutput
	err := workflow.ExecuteActivity(ctx, "aws.transfer.ListSecurityPolicies", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListSecurityPoliciesAsync(ctx workflow.Context, input *transfer.ListSecurityPoliciesInput) *ListSecurityPoliciesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transfer.ListSecurityPolicies", input)
	return &ListSecurityPoliciesFuture{Future: future}
}

func (a *stub) ListServers(ctx workflow.Context, input *transfer.ListServersInput) (*transfer.ListServersOutput, error) {
	var output transfer.ListServersOutput
	err := workflow.ExecuteActivity(ctx, "aws.transfer.ListServers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListServersAsync(ctx workflow.Context, input *transfer.ListServersInput) *ListServersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transfer.ListServers", input)
	return &ListServersFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *transfer.ListTagsForResourceInput) (*transfer.ListTagsForResourceOutput, error) {
	var output transfer.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.transfer.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *transfer.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transfer.ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) ListUsers(ctx workflow.Context, input *transfer.ListUsersInput) (*transfer.ListUsersOutput, error) {
	var output transfer.ListUsersOutput
	err := workflow.ExecuteActivity(ctx, "aws.transfer.ListUsers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListUsersAsync(ctx workflow.Context, input *transfer.ListUsersInput) *ListUsersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transfer.ListUsers", input)
	return &ListUsersFuture{Future: future}
}

func (a *stub) StartServer(ctx workflow.Context, input *transfer.StartServerInput) (*transfer.StartServerOutput, error) {
	var output transfer.StartServerOutput
	err := workflow.ExecuteActivity(ctx, "aws.transfer.StartServer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartServerAsync(ctx workflow.Context, input *transfer.StartServerInput) *StartServerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transfer.StartServer", input)
	return &StartServerFuture{Future: future}
}

func (a *stub) StopServer(ctx workflow.Context, input *transfer.StopServerInput) (*transfer.StopServerOutput, error) {
	var output transfer.StopServerOutput
	err := workflow.ExecuteActivity(ctx, "aws.transfer.StopServer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopServerAsync(ctx workflow.Context, input *transfer.StopServerInput) *StopServerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transfer.StopServer", input)
	return &StopServerFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *transfer.TagResourceInput) (*transfer.TagResourceOutput, error) {
	var output transfer.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.transfer.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *transfer.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transfer.TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) TestIdentityProvider(ctx workflow.Context, input *transfer.TestIdentityProviderInput) (*transfer.TestIdentityProviderOutput, error) {
	var output transfer.TestIdentityProviderOutput
	err := workflow.ExecuteActivity(ctx, "aws.transfer.TestIdentityProvider", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TestIdentityProviderAsync(ctx workflow.Context, input *transfer.TestIdentityProviderInput) *TestIdentityProviderFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transfer.TestIdentityProvider", input)
	return &TestIdentityProviderFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *transfer.UntagResourceInput) (*transfer.UntagResourceOutput, error) {
	var output transfer.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.transfer.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *transfer.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transfer.UntagResource", input)
	return &UntagResourceFuture{Future: future}
}

func (a *stub) UpdateServer(ctx workflow.Context, input *transfer.UpdateServerInput) (*transfer.UpdateServerOutput, error) {
	var output transfer.UpdateServerOutput
	err := workflow.ExecuteActivity(ctx, "aws.transfer.UpdateServer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateServerAsync(ctx workflow.Context, input *transfer.UpdateServerInput) *UpdateServerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transfer.UpdateServer", input)
	return &UpdateServerFuture{Future: future}
}

func (a *stub) UpdateUser(ctx workflow.Context, input *transfer.UpdateUserInput) (*transfer.UpdateUserOutput, error) {
	var output transfer.UpdateUserOutput
	err := workflow.ExecuteActivity(ctx, "aws.transfer.UpdateUser", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateUserAsync(ctx workflow.Context, input *transfer.UpdateUserInput) *UpdateUserFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transfer.UpdateUser", input)
	return &UpdateUserFuture{Future: future}
}
