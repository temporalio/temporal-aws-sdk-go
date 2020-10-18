// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package identitystorestub

import (
	"github.com/aws/aws-sdk-go/service/identitystore"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type IdentityStoreDescribeGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IdentityStoreDescribeGroupFuture) Get(ctx workflow.Context) (*identitystore.DescribeGroupOutput, error) {
	var output identitystore.DescribeGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IdentityStoreDescribeUserFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IdentityStoreDescribeUserFuture) Get(ctx workflow.Context) (*identitystore.DescribeUserOutput, error) {
	var output identitystore.DescribeUserOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IdentityStoreListGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IdentityStoreListGroupsFuture) Get(ctx workflow.Context) (*identitystore.ListGroupsOutput, error) {
	var output identitystore.ListGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type IdentityStoreListUsersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *IdentityStoreListUsersFuture) Get(ctx workflow.Context) (*identitystore.ListUsersOutput, error) {
	var output identitystore.ListUsersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeGroup(ctx workflow.Context, input *identitystore.DescribeGroupInput) (*identitystore.DescribeGroupOutput, error) {
	var output identitystore.DescribeGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.identitystore.DescribeGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeGroupAsync(ctx workflow.Context, input *identitystore.DescribeGroupInput) *IdentityStoreDescribeGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.identitystore.DescribeGroup", input)
	return &IdentityStoreDescribeGroupFuture{Future: future}
}

func (a *stub) DescribeUser(ctx workflow.Context, input *identitystore.DescribeUserInput) (*identitystore.DescribeUserOutput, error) {
	var output identitystore.DescribeUserOutput
	err := workflow.ExecuteActivity(ctx, "aws.identitystore.DescribeUser", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeUserAsync(ctx workflow.Context, input *identitystore.DescribeUserInput) *IdentityStoreDescribeUserFuture {
	future := workflow.ExecuteActivity(ctx, "aws.identitystore.DescribeUser", input)
	return &IdentityStoreDescribeUserFuture{Future: future}
}

func (a *stub) ListGroups(ctx workflow.Context, input *identitystore.ListGroupsInput) (*identitystore.ListGroupsOutput, error) {
	var output identitystore.ListGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws.identitystore.ListGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListGroupsAsync(ctx workflow.Context, input *identitystore.ListGroupsInput) *IdentityStoreListGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.identitystore.ListGroups", input)
	return &IdentityStoreListGroupsFuture{Future: future}
}

func (a *stub) ListUsers(ctx workflow.Context, input *identitystore.ListUsersInput) (*identitystore.ListUsersOutput, error) {
	var output identitystore.ListUsersOutput
	err := workflow.ExecuteActivity(ctx, "aws.identitystore.ListUsers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListUsersAsync(ctx workflow.Context, input *identitystore.ListUsersInput) *IdentityStoreListUsersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.identitystore.ListUsers", input)
	return &IdentityStoreListUsersFuture{Future: future}
}
