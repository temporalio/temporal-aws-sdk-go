package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
	"github.com/aws/aws-sdk-go/service/cognitoidentity/cognitoidentityiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type CognitoIdentityActivities struct {
	client cognitoidentityiface.CognitoIdentityAPI
}

func NewCognitoIdentityActivities(session *session.Session, config ...*aws.Config) *CognitoIdentityActivities {
	client := cognitoidentity.New(session, config...)
	return &CognitoIdentityActivities{client: client}
}

func (a *CognitoIdentityActivities) CreateIdentityPool(ctx context.Context, input *cognitoidentity.CreateIdentityPoolInput) (*cognitoidentity.IdentityPool, error) {
	return a.client.CreateIdentityPoolWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) DeleteIdentities(ctx context.Context, input *cognitoidentity.DeleteIdentitiesInput) (*cognitoidentity.DeleteIdentitiesOutput, error) {
	return a.client.DeleteIdentitiesWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) DeleteIdentityPool(ctx context.Context, input *cognitoidentity.DeleteIdentityPoolInput) (*cognitoidentity.DeleteIdentityPoolOutput, error) {
	return a.client.DeleteIdentityPoolWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) DescribeIdentity(ctx context.Context, input *cognitoidentity.DescribeIdentityInput) (*cognitoidentity.IdentityDescription, error) {
	return a.client.DescribeIdentityWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) DescribeIdentityPool(ctx context.Context, input *cognitoidentity.DescribeIdentityPoolInput) (*cognitoidentity.IdentityPool, error) {
	return a.client.DescribeIdentityPoolWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) GetCredentialsForIdentity(ctx context.Context, input *cognitoidentity.GetCredentialsForIdentityInput) (*cognitoidentity.GetCredentialsForIdentityOutput, error) {
	return a.client.GetCredentialsForIdentityWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) GetId(ctx context.Context, input *cognitoidentity.GetIdInput) (*cognitoidentity.GetIdOutput, error) {
	return a.client.GetIdWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) GetIdentityPoolRoles(ctx context.Context, input *cognitoidentity.GetIdentityPoolRolesInput) (*cognitoidentity.GetIdentityPoolRolesOutput, error) {
	return a.client.GetIdentityPoolRolesWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) GetOpenIdToken(ctx context.Context, input *cognitoidentity.GetOpenIdTokenInput) (*cognitoidentity.GetOpenIdTokenOutput, error) {
	return a.client.GetOpenIdTokenWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) GetOpenIdTokenForDeveloperIdentity(ctx context.Context, input *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput) (*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput, error) {
	return a.client.GetOpenIdTokenForDeveloperIdentityWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) ListIdentities(ctx context.Context, input *cognitoidentity.ListIdentitiesInput) (*cognitoidentity.ListIdentitiesOutput, error) {
	return a.client.ListIdentitiesWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) ListIdentityPools(ctx context.Context, input *cognitoidentity.ListIdentityPoolsInput) (*cognitoidentity.ListIdentityPoolsOutput, error) {
	return a.client.ListIdentityPoolsWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) ListTagsForResource(ctx context.Context, input *cognitoidentity.ListTagsForResourceInput) (*cognitoidentity.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) LookupDeveloperIdentity(ctx context.Context, input *cognitoidentity.LookupDeveloperIdentityInput) (*cognitoidentity.LookupDeveloperIdentityOutput, error) {
	return a.client.LookupDeveloperIdentityWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) MergeDeveloperIdentities(ctx context.Context, input *cognitoidentity.MergeDeveloperIdentitiesInput) (*cognitoidentity.MergeDeveloperIdentitiesOutput, error) {
	return a.client.MergeDeveloperIdentitiesWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) SetIdentityPoolRoles(ctx context.Context, input *cognitoidentity.SetIdentityPoolRolesInput) (*cognitoidentity.SetIdentityPoolRolesOutput, error) {
	return a.client.SetIdentityPoolRolesWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) TagResource(ctx context.Context, input *cognitoidentity.TagResourceInput) (*cognitoidentity.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) UnlinkDeveloperIdentity(ctx context.Context, input *cognitoidentity.UnlinkDeveloperIdentityInput) (*cognitoidentity.UnlinkDeveloperIdentityOutput, error) {
	return a.client.UnlinkDeveloperIdentityWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) UnlinkIdentity(ctx context.Context, input *cognitoidentity.UnlinkIdentityInput) (*cognitoidentity.UnlinkIdentityOutput, error) {
	return a.client.UnlinkIdentityWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) UntagResource(ctx context.Context, input *cognitoidentity.UntagResourceInput) (*cognitoidentity.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) UpdateIdentityPool(ctx context.Context, input *cognitoidentity.IdentityPool) (*cognitoidentity.IdentityPool, error) {
	return a.client.UpdateIdentityPoolWithContext(ctx, input)
}
