// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package cognitoidentitystub

import (
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateIdentityPool(ctx workflow.Context, input *cognitoidentity.CreateIdentityPoolInput) (*cognitoidentity.IdentityPool, error)
	CreateIdentityPoolAsync(ctx workflow.Context, input *cognitoidentity.CreateIdentityPoolInput) *CreateIdentityPoolFuture

	DeleteIdentities(ctx workflow.Context, input *cognitoidentity.DeleteIdentitiesInput) (*cognitoidentity.DeleteIdentitiesOutput, error)
	DeleteIdentitiesAsync(ctx workflow.Context, input *cognitoidentity.DeleteIdentitiesInput) *DeleteIdentitiesFuture

	DeleteIdentityPool(ctx workflow.Context, input *cognitoidentity.DeleteIdentityPoolInput) (*cognitoidentity.DeleteIdentityPoolOutput, error)
	DeleteIdentityPoolAsync(ctx workflow.Context, input *cognitoidentity.DeleteIdentityPoolInput) *DeleteIdentityPoolFuture

	DescribeIdentity(ctx workflow.Context, input *cognitoidentity.DescribeIdentityInput) (*cognitoidentity.IdentityDescription, error)
	DescribeIdentityAsync(ctx workflow.Context, input *cognitoidentity.DescribeIdentityInput) *DescribeIdentityFuture

	DescribeIdentityPool(ctx workflow.Context, input *cognitoidentity.DescribeIdentityPoolInput) (*cognitoidentity.IdentityPool, error)
	DescribeIdentityPoolAsync(ctx workflow.Context, input *cognitoidentity.DescribeIdentityPoolInput) *DescribeIdentityPoolFuture

	GetCredentialsForIdentity(ctx workflow.Context, input *cognitoidentity.GetCredentialsForIdentityInput) (*cognitoidentity.GetCredentialsForIdentityOutput, error)
	GetCredentialsForIdentityAsync(ctx workflow.Context, input *cognitoidentity.GetCredentialsForIdentityInput) *GetCredentialsForIdentityFuture

	GetId(ctx workflow.Context, input *cognitoidentity.GetIdInput) (*cognitoidentity.GetIdOutput, error)
	GetIdAsync(ctx workflow.Context, input *cognitoidentity.GetIdInput) *GetIdFuture

	GetIdentityPoolRoles(ctx workflow.Context, input *cognitoidentity.GetIdentityPoolRolesInput) (*cognitoidentity.GetIdentityPoolRolesOutput, error)
	GetIdentityPoolRolesAsync(ctx workflow.Context, input *cognitoidentity.GetIdentityPoolRolesInput) *GetIdentityPoolRolesFuture

	GetOpenIdToken(ctx workflow.Context, input *cognitoidentity.GetOpenIdTokenInput) (*cognitoidentity.GetOpenIdTokenOutput, error)
	GetOpenIdTokenAsync(ctx workflow.Context, input *cognitoidentity.GetOpenIdTokenInput) *GetOpenIdTokenFuture

	GetOpenIdTokenForDeveloperIdentity(ctx workflow.Context, input *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput) (*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput, error)
	GetOpenIdTokenForDeveloperIdentityAsync(ctx workflow.Context, input *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput) *GetOpenIdTokenForDeveloperIdentityFuture

	ListIdentities(ctx workflow.Context, input *cognitoidentity.ListIdentitiesInput) (*cognitoidentity.ListIdentitiesOutput, error)
	ListIdentitiesAsync(ctx workflow.Context, input *cognitoidentity.ListIdentitiesInput) *ListIdentitiesFuture

	ListIdentityPools(ctx workflow.Context, input *cognitoidentity.ListIdentityPoolsInput) (*cognitoidentity.ListIdentityPoolsOutput, error)
	ListIdentityPoolsAsync(ctx workflow.Context, input *cognitoidentity.ListIdentityPoolsInput) *ListIdentityPoolsFuture

	ListTagsForResource(ctx workflow.Context, input *cognitoidentity.ListTagsForResourceInput) (*cognitoidentity.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *cognitoidentity.ListTagsForResourceInput) *ListTagsForResourceFuture

	LookupDeveloperIdentity(ctx workflow.Context, input *cognitoidentity.LookupDeveloperIdentityInput) (*cognitoidentity.LookupDeveloperIdentityOutput, error)
	LookupDeveloperIdentityAsync(ctx workflow.Context, input *cognitoidentity.LookupDeveloperIdentityInput) *LookupDeveloperIdentityFuture

	MergeDeveloperIdentities(ctx workflow.Context, input *cognitoidentity.MergeDeveloperIdentitiesInput) (*cognitoidentity.MergeDeveloperIdentitiesOutput, error)
	MergeDeveloperIdentitiesAsync(ctx workflow.Context, input *cognitoidentity.MergeDeveloperIdentitiesInput) *MergeDeveloperIdentitiesFuture

	SetIdentityPoolRoles(ctx workflow.Context, input *cognitoidentity.SetIdentityPoolRolesInput) (*cognitoidentity.SetIdentityPoolRolesOutput, error)
	SetIdentityPoolRolesAsync(ctx workflow.Context, input *cognitoidentity.SetIdentityPoolRolesInput) *SetIdentityPoolRolesFuture

	TagResource(ctx workflow.Context, input *cognitoidentity.TagResourceInput) (*cognitoidentity.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *cognitoidentity.TagResourceInput) *TagResourceFuture

	UnlinkDeveloperIdentity(ctx workflow.Context, input *cognitoidentity.UnlinkDeveloperIdentityInput) (*cognitoidentity.UnlinkDeveloperIdentityOutput, error)
	UnlinkDeveloperIdentityAsync(ctx workflow.Context, input *cognitoidentity.UnlinkDeveloperIdentityInput) *UnlinkDeveloperIdentityFuture

	UnlinkIdentity(ctx workflow.Context, input *cognitoidentity.UnlinkIdentityInput) (*cognitoidentity.UnlinkIdentityOutput, error)
	UnlinkIdentityAsync(ctx workflow.Context, input *cognitoidentity.UnlinkIdentityInput) *UnlinkIdentityFuture

	UntagResource(ctx workflow.Context, input *cognitoidentity.UntagResourceInput) (*cognitoidentity.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *cognitoidentity.UntagResourceInput) *UntagResourceFuture

	UpdateIdentityPool(ctx workflow.Context, input *cognitoidentity.IdentityPool) (*cognitoidentity.IdentityPool, error)
	UpdateIdentityPoolAsync(ctx workflow.Context, input *cognitoidentity.IdentityPool) *UpdateIdentityPoolFuture
}

func NewClient() Client {
	return &stub{}
}
