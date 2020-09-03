package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
	"github.com/aws/aws-sdk-go/service/cognitoidentity/cognitoidentityiface"
)

type CognitoIdentityActivities struct {
	client cognitoidentityiface.CognitoIdentityAPI
}

func NewCognitoIdentityActivities(client cognitoidentityiface.CognitoIdentityAPI) *CognitoIdentityActivities {
    return &CognitoIdentityActivities{client: client}
}

func (a *CognitoIdentityActivities) CreateIdentityPool(input *cognitoidentity.CreateIdentityPoolInput) (*cognitoidentity.CreateIdentityPoolOutput, error) {
    return a.client.CreateIdentityPool(input)
}

func (a *CognitoIdentityActivities) DeleteIdentities(input *cognitoidentity.DeleteIdentitiesInput) (*cognitoidentity.DeleteIdentitiesOutput, error) {
    return a.client.DeleteIdentities(input)
}

func (a *CognitoIdentityActivities) DeleteIdentityPool(input *cognitoidentity.DeleteIdentityPoolInput) (*cognitoidentity.DeleteIdentityPoolOutput, error) {
    return a.client.DeleteIdentityPool(input)
}

func (a *CognitoIdentityActivities) DescribeIdentity(input *cognitoidentity.DescribeIdentityInput) (*cognitoidentity.DescribeIdentityOutput, error) {
    return a.client.DescribeIdentity(input)
}

func (a *CognitoIdentityActivities) DescribeIdentityPool(input *cognitoidentity.DescribeIdentityPoolInput) (*cognitoidentity.DescribeIdentityPoolOutput, error) {
    return a.client.DescribeIdentityPool(input)
}

func (a *CognitoIdentityActivities) GetCredentialsForIdentity(input *cognitoidentity.GetCredentialsForIdentityInput) (*cognitoidentity.GetCredentialsForIdentityOutput, error) {
    return a.client.GetCredentialsForIdentity(input)
}

func (a *CognitoIdentityActivities) GetId(input *cognitoidentity.GetIdInput) (*cognitoidentity.GetIdOutput, error) {
    return a.client.GetId(input)
}

func (a *CognitoIdentityActivities) GetIdentityPoolRoles(input *cognitoidentity.GetIdentityPoolRolesInput) (*cognitoidentity.GetIdentityPoolRolesOutput, error) {
    return a.client.GetIdentityPoolRoles(input)
}

func (a *CognitoIdentityActivities) GetOpenIdToken(input *cognitoidentity.GetOpenIdTokenInput) (*cognitoidentity.GetOpenIdTokenOutput, error) {
    return a.client.GetOpenIdToken(input)
}

func (a *CognitoIdentityActivities) GetOpenIdTokenForDeveloperIdentity(input *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput) (*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput, error) {
    return a.client.GetOpenIdTokenForDeveloperIdentity(input)
}

func (a *CognitoIdentityActivities) ListIdentities(input *cognitoidentity.ListIdentitiesInput) (*cognitoidentity.ListIdentitiesOutput, error) {
    return a.client.ListIdentities(input)
}

func (a *CognitoIdentityActivities) ListIdentityPools(input *cognitoidentity.ListIdentityPoolsInput) (*cognitoidentity.ListIdentityPoolsOutput, error) {
    return a.client.ListIdentityPools(input)
}

func (a *CognitoIdentityActivities) ListTagsForResource(input *cognitoidentity.ListTagsForResourceInput) (*cognitoidentity.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *CognitoIdentityActivities) LookupDeveloperIdentity(input *cognitoidentity.LookupDeveloperIdentityInput) (*cognitoidentity.LookupDeveloperIdentityOutput, error) {
    return a.client.LookupDeveloperIdentity(input)
}

func (a *CognitoIdentityActivities) MergeDeveloperIdentities(input *cognitoidentity.MergeDeveloperIdentitiesInput) (*cognitoidentity.MergeDeveloperIdentitiesOutput, error) {
    return a.client.MergeDeveloperIdentities(input)
}

func (a *CognitoIdentityActivities) SetIdentityPoolRoles(input *cognitoidentity.SetIdentityPoolRolesInput) (*cognitoidentity.SetIdentityPoolRolesOutput, error) {
    return a.client.SetIdentityPoolRoles(input)
}

func (a *CognitoIdentityActivities) TagResource(input *cognitoidentity.TagResourceInput) (*cognitoidentity.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *CognitoIdentityActivities) UnlinkDeveloperIdentity(input *cognitoidentity.UnlinkDeveloperIdentityInput) (*cognitoidentity.UnlinkDeveloperIdentityOutput, error) {
    return a.client.UnlinkDeveloperIdentity(input)
}

func (a *CognitoIdentityActivities) UnlinkIdentity(input *cognitoidentity.UnlinkIdentityInput) (*cognitoidentity.UnlinkIdentityOutput, error) {
    return a.client.UnlinkIdentity(input)
}

func (a *CognitoIdentityActivities) UntagResource(input *cognitoidentity.UntagResourceInput) (*cognitoidentity.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *CognitoIdentityActivities) UpdateIdentityPool(input *cognitoidentity.UpdateIdentityPoolInput) (*cognitoidentity.UpdateIdentityPoolOutput, error) {
    return a.client.UpdateIdentityPool(input)
}
