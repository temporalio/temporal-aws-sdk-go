package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/sso"
	"github.com/aws/aws-sdk-go/service/sso/ssoiface"
)

type SSOActivities struct {
	client ssoiface.SSOAPI
}

func NewSSOActivities(client ssoiface.SSOAPI) *SSOActivities {
    return &SSOActivities{client: client}
}

func (a *SSOActivities) GetRoleCredentials(input *sso.GetRoleCredentialsInput) (*sso.GetRoleCredentialsOutput, error) {
    return a.client.GetRoleCredentials(input)
}

func (a *SSOActivities) ListAccountRoles(input *sso.ListAccountRolesInput) (*sso.ListAccountRolesOutput, error) {
    return a.client.ListAccountRoles(input)
}

func (a *SSOActivities) ListAccounts(input *sso.ListAccountsInput) (*sso.ListAccountsOutput, error) {
    return a.client.ListAccounts(input)
}

func (a *SSOActivities) Logout(input *sso.LogoutInput) (*sso.LogoutOutput, error) {
    return a.client.Logout(input)
}
