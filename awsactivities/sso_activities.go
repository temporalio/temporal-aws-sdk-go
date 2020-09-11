
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sso"
	"github.com/aws/aws-sdk-go/service/sso/ssoiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type SSOActivities struct {
    client ssoiface.SSOAPI
}

func NewSSOActivities(session *session.Session, config ...*aws.Config) *SSOActivities {
    client := sso.New(session, config...)
    return &SSOActivities{client: client}
}

func (a *SSOActivities) GetRoleCredentials(ctx context.Context, input *sso.GetRoleCredentialsInput) (*sso.GetRoleCredentialsOutput, error) {
    return a.client.GetRoleCredentialsWithContext(ctx, input)
}

func (a *SSOActivities) ListAccountRoles(ctx context.Context, input *sso.ListAccountRolesInput) (*sso.ListAccountRolesOutput, error) {
    return a.client.ListAccountRolesWithContext(ctx, input)
}

func (a *SSOActivities) ListAccounts(ctx context.Context, input *sso.ListAccountsInput) (*sso.ListAccountsOutput, error) {
    return a.client.ListAccountsWithContext(ctx, input)
}

func (a *SSOActivities) Logout(ctx context.Context, input *sso.LogoutInput) (*sso.LogoutOutput, error) {
    return a.client.LogoutWithContext(ctx, input)
}
