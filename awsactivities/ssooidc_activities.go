package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/ssooidc"
	"github.com/aws/aws-sdk-go/service/ssooidc/ssooidciface"
)

type SSOOIDCActivities struct {
	client ssooidciface.SSOOIDCAPI
}

func NewSSOOIDCActivities(client ssooidciface.SSOOIDCAPI) *SSOOIDCActivities {
    return &SSOOIDCActivities{client: client}
}

func (a *SSOOIDCActivities) CreateToken(input *ssooidc.CreateTokenInput) (*ssooidc.CreateTokenOutput, error) {
    return a.client.CreateToken(input)
}

func (a *SSOOIDCActivities) RegisterClient(input *ssooidc.RegisterClientInput) (*ssooidc.RegisterClientOutput, error) {
    return a.client.RegisterClient(input)
}

func (a *SSOOIDCActivities) StartDeviceAuthorization(input *ssooidc.StartDeviceAuthorizationInput) (*ssooidc.StartDeviceAuthorizationOutput, error) {
    return a.client.StartDeviceAuthorization(input)
}
