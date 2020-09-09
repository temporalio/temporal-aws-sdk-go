package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
)

type STSActivities struct {
	client stsiface.STSAPI
}

func NewSTSActivities(session *session.Session, config ...*aws.Config) *STSActivities {
	client := sts.New(session, config...)
	return &STSActivities{client: client}
}

func (a *STSActivities) AssumeRole(input *sts.AssumeRoleInput) (*sts.AssumeRoleOutput, error) {
	return a.client.AssumeRole(input)
}

func (a *STSActivities) AssumeRoleWithSAML(input *sts.AssumeRoleWithSAMLInput) (*sts.AssumeRoleWithSAMLOutput, error) {
	return a.client.AssumeRoleWithSAML(input)
}

func (a *STSActivities) AssumeRoleWithWebIdentity(input *sts.AssumeRoleWithWebIdentityInput) (*sts.AssumeRoleWithWebIdentityOutput, error) {
	return a.client.AssumeRoleWithWebIdentity(input)
}

func (a *STSActivities) DecodeAuthorizationMessage(input *sts.DecodeAuthorizationMessageInput) (*sts.DecodeAuthorizationMessageOutput, error) {
	return a.client.DecodeAuthorizationMessage(input)
}

func (a *STSActivities) GetAccessKeyInfo(input *sts.GetAccessKeyInfoInput) (*sts.GetAccessKeyInfoOutput, error) {
	return a.client.GetAccessKeyInfo(input)
}

func (a *STSActivities) GetCallerIdentity(input *sts.GetCallerIdentityInput) (*sts.GetCallerIdentityOutput, error) {
	return a.client.GetCallerIdentity(input)
}

func (a *STSActivities) GetFederationToken(input *sts.GetFederationTokenInput) (*sts.GetFederationTokenOutput, error) {
	return a.client.GetFederationToken(input)
}

func (a *STSActivities) GetSessionToken(input *sts.GetSessionTokenInput) (*sts.GetSessionTokenOutput, error) {
	return a.client.GetSessionToken(input)
}
