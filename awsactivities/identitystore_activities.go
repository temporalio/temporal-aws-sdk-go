package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/identitystore"
	"github.com/aws/aws-sdk-go/service/identitystore/identitystoreiface"
)

type IdentityStoreActivities struct {
	client identitystoreiface.IdentityStoreAPI
}

func NewIdentityStoreActivities(client identitystoreiface.IdentityStoreAPI) *IdentityStoreActivities {
    return &IdentityStoreActivities{client: client}
}

func (a *IdentityStoreActivities) DescribeGroup(input *identitystore.DescribeGroupInput) (*identitystore.DescribeGroupOutput, error) {
    return a.client.DescribeGroup(input)
}

func (a *IdentityStoreActivities) DescribeUser(input *identitystore.DescribeUserInput) (*identitystore.DescribeUserOutput, error) {
    return a.client.DescribeUser(input)
}

func (a *IdentityStoreActivities) ListGroups(input *identitystore.ListGroupsInput) (*identitystore.ListGroupsOutput, error) {
    return a.client.ListGroups(input)
}

func (a *IdentityStoreActivities) ListUsers(input *identitystore.ListUsersInput) (*identitystore.ListUsersOutput, error) {
    return a.client.ListUsers(input)
}
