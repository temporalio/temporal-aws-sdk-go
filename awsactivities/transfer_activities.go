
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/transfer"
	"github.com/aws/aws-sdk-go/service/transfer/transferiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type TransferActivities struct {
    client transferiface.TransferAPI
}

func NewTransferActivities(session *session.Session, config ...*aws.Config) *TransferActivities {
    client := transfer.New(session, config...)
    return &TransferActivities{client: client}
}

func (a *TransferActivities) CreateServer(ctx context.Context, input *transfer.CreateServerInput) (*transfer.CreateServerOutput, error) {
    return a.client.CreateServerWithContext(ctx, input)
}

func (a *TransferActivities) CreateUser(ctx context.Context, input *transfer.CreateUserInput) (*transfer.CreateUserOutput, error) {
    return a.client.CreateUserWithContext(ctx, input)
}

func (a *TransferActivities) DeleteServer(ctx context.Context, input *transfer.DeleteServerInput) (*transfer.DeleteServerOutput, error) {
    return a.client.DeleteServerWithContext(ctx, input)
}

func (a *TransferActivities) DeleteSshPublicKey(ctx context.Context, input *transfer.DeleteSshPublicKeyInput) (*transfer.DeleteSshPublicKeyOutput, error) {
    return a.client.DeleteSshPublicKeyWithContext(ctx, input)
}

func (a *TransferActivities) DeleteUser(ctx context.Context, input *transfer.DeleteUserInput) (*transfer.DeleteUserOutput, error) {
    return a.client.DeleteUserWithContext(ctx, input)
}

func (a *TransferActivities) DescribeSecurityPolicy(ctx context.Context, input *transfer.DescribeSecurityPolicyInput) (*transfer.DescribeSecurityPolicyOutput, error) {
    return a.client.DescribeSecurityPolicyWithContext(ctx, input)
}

func (a *TransferActivities) DescribeServer(ctx context.Context, input *transfer.DescribeServerInput) (*transfer.DescribeServerOutput, error) {
    return a.client.DescribeServerWithContext(ctx, input)
}

func (a *TransferActivities) DescribeUser(ctx context.Context, input *transfer.DescribeUserInput) (*transfer.DescribeUserOutput, error) {
    return a.client.DescribeUserWithContext(ctx, input)
}

func (a *TransferActivities) ImportSshPublicKey(ctx context.Context, input *transfer.ImportSshPublicKeyInput) (*transfer.ImportSshPublicKeyOutput, error) {
    return a.client.ImportSshPublicKeyWithContext(ctx, input)
}

func (a *TransferActivities) ListSecurityPolicies(ctx context.Context, input *transfer.ListSecurityPoliciesInput) (*transfer.ListSecurityPoliciesOutput, error) {
    return a.client.ListSecurityPoliciesWithContext(ctx, input)
}

func (a *TransferActivities) ListServers(ctx context.Context, input *transfer.ListServersInput) (*transfer.ListServersOutput, error) {
    return a.client.ListServersWithContext(ctx, input)
}

func (a *TransferActivities) ListTagsForResource(ctx context.Context, input *transfer.ListTagsForResourceInput) (*transfer.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *TransferActivities) ListUsers(ctx context.Context, input *transfer.ListUsersInput) (*transfer.ListUsersOutput, error) {
    return a.client.ListUsersWithContext(ctx, input)
}

func (a *TransferActivities) StartServer(ctx context.Context, input *transfer.StartServerInput) (*transfer.StartServerOutput, error) {
    return a.client.StartServerWithContext(ctx, input)
}

func (a *TransferActivities) StopServer(ctx context.Context, input *transfer.StopServerInput) (*transfer.StopServerOutput, error) {
    return a.client.StopServerWithContext(ctx, input)
}

func (a *TransferActivities) TagResource(ctx context.Context, input *transfer.TagResourceInput) (*transfer.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *TransferActivities) TestIdentityProvider(ctx context.Context, input *transfer.TestIdentityProviderInput) (*transfer.TestIdentityProviderOutput, error) {
    return a.client.TestIdentityProviderWithContext(ctx, input)
}

func (a *TransferActivities) UntagResource(ctx context.Context, input *transfer.UntagResourceInput) (*transfer.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *TransferActivities) UpdateServer(ctx context.Context, input *transfer.UpdateServerInput) (*transfer.UpdateServerOutput, error) {
    return a.client.UpdateServerWithContext(ctx, input)
}

func (a *TransferActivities) UpdateUser(ctx context.Context, input *transfer.UpdateUserInput) (*transfer.UpdateUserOutput, error) {
    return a.client.UpdateUserWithContext(ctx, input)
}
