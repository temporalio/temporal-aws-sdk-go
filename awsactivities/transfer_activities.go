
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/transfer"
	"github.com/aws/aws-sdk-go/service/transfer/transferiface"
)

type TransferActivities struct {
    client transferiface.TransferAPI
}

func NewTransferActivities(session *session.Session, config ...*aws.Config) *TransferActivities {
    client := transfer.New(session, config...)
    return &TransferActivities{client: client}
}

func (a *TransferActivities) CreateServer(input *transfer.CreateServerInput) (*transfer.CreateServerOutput, error) {
    return a.client.CreateServer(input)
}

func (a *TransferActivities) CreateUser(input *transfer.CreateUserInput) (*transfer.CreateUserOutput, error) {
    return a.client.CreateUser(input)
}

func (a *TransferActivities) DeleteServer(input *transfer.DeleteServerInput) (*transfer.DeleteServerOutput, error) {
    return a.client.DeleteServer(input)
}

func (a *TransferActivities) DeleteSshPublicKey(input *transfer.DeleteSshPublicKeyInput) (*transfer.DeleteSshPublicKeyOutput, error) {
    return a.client.DeleteSshPublicKey(input)
}

func (a *TransferActivities) DeleteUser(input *transfer.DeleteUserInput) (*transfer.DeleteUserOutput, error) {
    return a.client.DeleteUser(input)
}

func (a *TransferActivities) DescribeSecurityPolicy(input *transfer.DescribeSecurityPolicyInput) (*transfer.DescribeSecurityPolicyOutput, error) {
    return a.client.DescribeSecurityPolicy(input)
}

func (a *TransferActivities) DescribeServer(input *transfer.DescribeServerInput) (*transfer.DescribeServerOutput, error) {
    return a.client.DescribeServer(input)
}

func (a *TransferActivities) DescribeUser(input *transfer.DescribeUserInput) (*transfer.DescribeUserOutput, error) {
    return a.client.DescribeUser(input)
}

func (a *TransferActivities) ImportSshPublicKey(input *transfer.ImportSshPublicKeyInput) (*transfer.ImportSshPublicKeyOutput, error) {
    return a.client.ImportSshPublicKey(input)
}

func (a *TransferActivities) ListSecurityPolicies(input *transfer.ListSecurityPoliciesInput) (*transfer.ListSecurityPoliciesOutput, error) {
    return a.client.ListSecurityPolicies(input)
}

func (a *TransferActivities) ListServers(input *transfer.ListServersInput) (*transfer.ListServersOutput, error) {
    return a.client.ListServers(input)
}

func (a *TransferActivities) ListTagsForResource(input *transfer.ListTagsForResourceInput) (*transfer.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *TransferActivities) ListUsers(input *transfer.ListUsersInput) (*transfer.ListUsersOutput, error) {
    return a.client.ListUsers(input)
}

func (a *TransferActivities) StartServer(input *transfer.StartServerInput) (*transfer.StartServerOutput, error) {
    return a.client.StartServer(input)
}

func (a *TransferActivities) StopServer(input *transfer.StopServerInput) (*transfer.StopServerOutput, error) {
    return a.client.StopServer(input)
}

func (a *TransferActivities) TagResource(input *transfer.TagResourceInput) (*transfer.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *TransferActivities) TestIdentityProvider(input *transfer.TestIdentityProviderInput) (*transfer.TestIdentityProviderOutput, error) {
    return a.client.TestIdentityProvider(input)
}

func (a *TransferActivities) UntagResource(input *transfer.UntagResourceInput) (*transfer.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *TransferActivities) UpdateServer(input *transfer.UpdateServerInput) (*transfer.UpdateServerOutput, error) {
    return a.client.UpdateServer(input)
}

func (a *TransferActivities) UpdateUser(input *transfer.UpdateUserInput) (*transfer.UpdateUserOutput, error) {
    return a.client.UpdateUser(input)
}
