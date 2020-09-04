package awsclients

import (
	"github.com/aws/aws-sdk-go/service/identitystore"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type IdentityStoreClient interface {
    DescribeGroup(ctx workflow.Context, input *identitystore.DescribeGroupInput) (*identitystore.DescribeGroupOutput, error)
    DescribeGroupAsync(ctx workflow.Context, input *identitystore.DescribeGroupInput) *IdentitystoreDescribeGroupResult

    DescribeUser(ctx workflow.Context, input *identitystore.DescribeUserInput) (*identitystore.DescribeUserOutput, error)
    DescribeUserAsync(ctx workflow.Context, input *identitystore.DescribeUserInput) *IdentitystoreDescribeUserResult

    ListGroups(ctx workflow.Context, input *identitystore.ListGroupsInput) (*identitystore.ListGroupsOutput, error)
    ListGroupsAsync(ctx workflow.Context, input *identitystore.ListGroupsInput) *IdentitystoreListGroupsResult

    ListUsers(ctx workflow.Context, input *identitystore.ListUsersInput) (*identitystore.ListUsersOutput, error)
    ListUsersAsync(ctx workflow.Context, input *identitystore.ListUsersInput) *IdentitystoreListUsersResult
}
type IdentitystoreDescribeGroupResult struct {
	Result workflow.Future
}

func (r *IdentitystoreDescribeGroupResult) Get(ctx workflow.Context) (*identitystore.DescribeGroupOutput, error) {
    var output identitystore.DescribeGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IdentitystoreDescribeUserResult struct {
	Result workflow.Future
}

func (r *IdentitystoreDescribeUserResult) Get(ctx workflow.Context) (*identitystore.DescribeUserOutput, error) {
    var output identitystore.DescribeUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IdentitystoreListGroupsResult struct {
	Result workflow.Future
}

func (r *IdentitystoreListGroupsResult) Get(ctx workflow.Context) (*identitystore.ListGroupsOutput, error) {
    var output identitystore.ListGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IdentitystoreListUsersResult struct {
	Result workflow.Future
}

func (r *IdentitystoreListUsersResult) Get(ctx workflow.Context) (*identitystore.ListUsersOutput, error) {
    var output identitystore.ListUsersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type IdentityStoreStub struct {
    activities awsactivities.IdentityStoreActivities
}

func NewIdentityStoreStub() IdentityStoreClient {
    return &IdentityStoreStub{}
}
func (a *IdentityStoreStub) DescribeGroup(ctx workflow.Context, input *identitystore.DescribeGroupInput) (*identitystore.DescribeGroupOutput, error) {
    var output identitystore.DescribeGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IdentityStoreStub) DescribeGroupAsync(ctx workflow.Context, input *identitystore.DescribeGroupInput) *IdentitystoreDescribeGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeGroup, input)
    return &IdentitystoreDescribeGroupResult{Result: future}
}
func (a *IdentityStoreStub) DescribeUser(ctx workflow.Context, input *identitystore.DescribeUserInput) (*identitystore.DescribeUserOutput, error) {
    var output identitystore.DescribeUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeUser, input).Get(ctx, &output)
    return &output, err
}

func (a *IdentityStoreStub) DescribeUserAsync(ctx workflow.Context, input *identitystore.DescribeUserInput) *IdentitystoreDescribeUserResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeUser, input)
    return &IdentitystoreDescribeUserResult{Result: future}
}
func (a *IdentityStoreStub) ListGroups(ctx workflow.Context, input *identitystore.ListGroupsInput) (*identitystore.ListGroupsOutput, error) {
    var output identitystore.ListGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *IdentityStoreStub) ListGroupsAsync(ctx workflow.Context, input *identitystore.ListGroupsInput) *IdentitystoreListGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListGroups, input)
    return &IdentitystoreListGroupsResult{Result: future}
}
func (a *IdentityStoreStub) ListUsers(ctx workflow.Context, input *identitystore.ListUsersInput) (*identitystore.ListUsersOutput, error) {
    var output identitystore.ListUsersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListUsers, input).Get(ctx, &output)
    return &output, err
}

func (a *IdentityStoreStub) ListUsersAsync(ctx workflow.Context, input *identitystore.ListUsersInput) *IdentitystoreListUsersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListUsers, input)
    return &IdentitystoreListUsersResult{Result: future}
}
