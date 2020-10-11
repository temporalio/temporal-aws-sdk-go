// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
	"github.com/aws/aws-sdk-go/service/cognitoidentity/cognitoidentityiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type CognitoIdentityActivities struct {
	client cognitoidentityiface.CognitoIdentityAPI

	sessionFactory SessionFactory
}

func NewCognitoIdentityActivities(sess *session.Session, config ...*aws.Config) *CognitoIdentityActivities {
	client := cognitoidentity.New(sess, config...)
	return &CognitoIdentityActivities{client: client}
}

func NewCognitoIdentityActivitiesWithSessionFactory(sessionFactory SessionFactory) *CognitoIdentityActivities {
	return &CognitoIdentityActivities{sessionFactory: sessionFactory}
}

func (a *CognitoIdentityActivities) getClient(ctx context.Context) (cognitoidentityiface.CognitoIdentityAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return cognitoidentity.New(sess), nil
}

func (a *CognitoIdentityActivities) CreateIdentityPool(ctx context.Context, input *cognitoidentity.CreateIdentityPoolInput) (*cognitoidentity.IdentityPool, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateIdentityPoolWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) DeleteIdentities(ctx context.Context, input *cognitoidentity.DeleteIdentitiesInput) (*cognitoidentity.DeleteIdentitiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteIdentitiesWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) DeleteIdentityPool(ctx context.Context, input *cognitoidentity.DeleteIdentityPoolInput) (*cognitoidentity.DeleteIdentityPoolOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteIdentityPoolWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) DescribeIdentity(ctx context.Context, input *cognitoidentity.DescribeIdentityInput) (*cognitoidentity.IdentityDescription, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeIdentityWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) DescribeIdentityPool(ctx context.Context, input *cognitoidentity.DescribeIdentityPoolInput) (*cognitoidentity.IdentityPool, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeIdentityPoolWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) GetCredentialsForIdentity(ctx context.Context, input *cognitoidentity.GetCredentialsForIdentityInput) (*cognitoidentity.GetCredentialsForIdentityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetCredentialsForIdentityWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) GetId(ctx context.Context, input *cognitoidentity.GetIdInput) (*cognitoidentity.GetIdOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetIdWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) GetIdentityPoolRoles(ctx context.Context, input *cognitoidentity.GetIdentityPoolRolesInput) (*cognitoidentity.GetIdentityPoolRolesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetIdentityPoolRolesWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) GetOpenIdToken(ctx context.Context, input *cognitoidentity.GetOpenIdTokenInput) (*cognitoidentity.GetOpenIdTokenOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetOpenIdTokenWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) GetOpenIdTokenForDeveloperIdentity(ctx context.Context, input *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput) (*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetOpenIdTokenForDeveloperIdentityWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) ListIdentities(ctx context.Context, input *cognitoidentity.ListIdentitiesInput) (*cognitoidentity.ListIdentitiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListIdentitiesWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) ListIdentityPools(ctx context.Context, input *cognitoidentity.ListIdentityPoolsInput) (*cognitoidentity.ListIdentityPoolsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListIdentityPoolsWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) ListTagsForResource(ctx context.Context, input *cognitoidentity.ListTagsForResourceInput) (*cognitoidentity.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) LookupDeveloperIdentity(ctx context.Context, input *cognitoidentity.LookupDeveloperIdentityInput) (*cognitoidentity.LookupDeveloperIdentityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.LookupDeveloperIdentityWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) MergeDeveloperIdentities(ctx context.Context, input *cognitoidentity.MergeDeveloperIdentitiesInput) (*cognitoidentity.MergeDeveloperIdentitiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.MergeDeveloperIdentitiesWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) SetIdentityPoolRoles(ctx context.Context, input *cognitoidentity.SetIdentityPoolRolesInput) (*cognitoidentity.SetIdentityPoolRolesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SetIdentityPoolRolesWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) TagResource(ctx context.Context, input *cognitoidentity.TagResourceInput) (*cognitoidentity.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) UnlinkDeveloperIdentity(ctx context.Context, input *cognitoidentity.UnlinkDeveloperIdentityInput) (*cognitoidentity.UnlinkDeveloperIdentityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UnlinkDeveloperIdentityWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) UnlinkIdentity(ctx context.Context, input *cognitoidentity.UnlinkIdentityInput) (*cognitoidentity.UnlinkIdentityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UnlinkIdentityWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) UntagResource(ctx context.Context, input *cognitoidentity.UntagResourceInput) (*cognitoidentity.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *CognitoIdentityActivities) UpdateIdentityPool(ctx context.Context, input *cognitoidentity.IdentityPool) (*cognitoidentity.IdentityPool, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateIdentityPoolWithContext(ctx, input)
}
