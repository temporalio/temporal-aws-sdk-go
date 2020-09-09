package awsclients

import (
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type CognitoIdentityClient interface {
	CreateIdentityPool(ctx workflow.Context, input *cognitoidentity.CreateIdentityPoolInput) (*cognitoidentity.IdentityPool, error)
	CreateIdentityPoolAsync(ctx workflow.Context, input *cognitoidentity.CreateIdentityPoolInput) *CognitoidentityCreateIdentityPoolResult

	DeleteIdentities(ctx workflow.Context, input *cognitoidentity.DeleteIdentitiesInput) (*cognitoidentity.DeleteIdentitiesOutput, error)
	DeleteIdentitiesAsync(ctx workflow.Context, input *cognitoidentity.DeleteIdentitiesInput) *CognitoidentityDeleteIdentitiesResult

	DeleteIdentityPool(ctx workflow.Context, input *cognitoidentity.DeleteIdentityPoolInput) (*cognitoidentity.DeleteIdentityPoolOutput, error)
	DeleteIdentityPoolAsync(ctx workflow.Context, input *cognitoidentity.DeleteIdentityPoolInput) *CognitoidentityDeleteIdentityPoolResult

	DescribeIdentity(ctx workflow.Context, input *cognitoidentity.DescribeIdentityInput) (*cognitoidentity.IdentityDescription, error)
	DescribeIdentityAsync(ctx workflow.Context, input *cognitoidentity.DescribeIdentityInput) *CognitoidentityDescribeIdentityResult

	DescribeIdentityPool(ctx workflow.Context, input *cognitoidentity.DescribeIdentityPoolInput) (*cognitoidentity.IdentityPool, error)
	DescribeIdentityPoolAsync(ctx workflow.Context, input *cognitoidentity.DescribeIdentityPoolInput) *CognitoidentityDescribeIdentityPoolResult

	GetCredentialsForIdentity(ctx workflow.Context, input *cognitoidentity.GetCredentialsForIdentityInput) (*cognitoidentity.GetCredentialsForIdentityOutput, error)
	GetCredentialsForIdentityAsync(ctx workflow.Context, input *cognitoidentity.GetCredentialsForIdentityInput) *CognitoidentityGetCredentialsForIdentityResult

	GetId(ctx workflow.Context, input *cognitoidentity.GetIdInput) (*cognitoidentity.GetIdOutput, error)
	GetIdAsync(ctx workflow.Context, input *cognitoidentity.GetIdInput) *CognitoidentityGetIdResult

	GetIdentityPoolRoles(ctx workflow.Context, input *cognitoidentity.GetIdentityPoolRolesInput) (*cognitoidentity.GetIdentityPoolRolesOutput, error)
	GetIdentityPoolRolesAsync(ctx workflow.Context, input *cognitoidentity.GetIdentityPoolRolesInput) *CognitoidentityGetIdentityPoolRolesResult

	GetOpenIdToken(ctx workflow.Context, input *cognitoidentity.GetOpenIdTokenInput) (*cognitoidentity.GetOpenIdTokenOutput, error)
	GetOpenIdTokenAsync(ctx workflow.Context, input *cognitoidentity.GetOpenIdTokenInput) *CognitoidentityGetOpenIdTokenResult

	GetOpenIdTokenForDeveloperIdentity(ctx workflow.Context, input *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput) (*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput, error)
	GetOpenIdTokenForDeveloperIdentityAsync(ctx workflow.Context, input *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput) *CognitoidentityGetOpenIdTokenForDeveloperIdentityResult

	ListIdentities(ctx workflow.Context, input *cognitoidentity.ListIdentitiesInput) (*cognitoidentity.ListIdentitiesOutput, error)
	ListIdentitiesAsync(ctx workflow.Context, input *cognitoidentity.ListIdentitiesInput) *CognitoidentityListIdentitiesResult

	ListIdentityPools(ctx workflow.Context, input *cognitoidentity.ListIdentityPoolsInput) (*cognitoidentity.ListIdentityPoolsOutput, error)
	ListIdentityPoolsAsync(ctx workflow.Context, input *cognitoidentity.ListIdentityPoolsInput) *CognitoidentityListIdentityPoolsResult

	ListTagsForResource(ctx workflow.Context, input *cognitoidentity.ListTagsForResourceInput) (*cognitoidentity.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *cognitoidentity.ListTagsForResourceInput) *CognitoidentityListTagsForResourceResult

	LookupDeveloperIdentity(ctx workflow.Context, input *cognitoidentity.LookupDeveloperIdentityInput) (*cognitoidentity.LookupDeveloperIdentityOutput, error)
	LookupDeveloperIdentityAsync(ctx workflow.Context, input *cognitoidentity.LookupDeveloperIdentityInput) *CognitoidentityLookupDeveloperIdentityResult

	MergeDeveloperIdentities(ctx workflow.Context, input *cognitoidentity.MergeDeveloperIdentitiesInput) (*cognitoidentity.MergeDeveloperIdentitiesOutput, error)
	MergeDeveloperIdentitiesAsync(ctx workflow.Context, input *cognitoidentity.MergeDeveloperIdentitiesInput) *CognitoidentityMergeDeveloperIdentitiesResult

	SetIdentityPoolRoles(ctx workflow.Context, input *cognitoidentity.SetIdentityPoolRolesInput) (*cognitoidentity.SetIdentityPoolRolesOutput, error)
	SetIdentityPoolRolesAsync(ctx workflow.Context, input *cognitoidentity.SetIdentityPoolRolesInput) *CognitoidentitySetIdentityPoolRolesResult

	TagResource(ctx workflow.Context, input *cognitoidentity.TagResourceInput) (*cognitoidentity.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *cognitoidentity.TagResourceInput) *CognitoidentityTagResourceResult

	UnlinkDeveloperIdentity(ctx workflow.Context, input *cognitoidentity.UnlinkDeveloperIdentityInput) (*cognitoidentity.UnlinkDeveloperIdentityOutput, error)
	UnlinkDeveloperIdentityAsync(ctx workflow.Context, input *cognitoidentity.UnlinkDeveloperIdentityInput) *CognitoidentityUnlinkDeveloperIdentityResult

	UnlinkIdentity(ctx workflow.Context, input *cognitoidentity.UnlinkIdentityInput) (*cognitoidentity.UnlinkIdentityOutput, error)
	UnlinkIdentityAsync(ctx workflow.Context, input *cognitoidentity.UnlinkIdentityInput) *CognitoidentityUnlinkIdentityResult

	UntagResource(ctx workflow.Context, input *cognitoidentity.UntagResourceInput) (*cognitoidentity.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *cognitoidentity.UntagResourceInput) *CognitoidentityUntagResourceResult

	UpdateIdentityPool(ctx workflow.Context, input *cognitoidentity.IdentityPool) (*cognitoidentity.IdentityPool, error)
	UpdateIdentityPoolAsync(ctx workflow.Context, input *cognitoidentity.IdentityPool) *CognitoidentityUpdateIdentityPoolResult
}

type CognitoidentityCreateIdentityPoolResult struct {
	Result workflow.Future
}

func (r *CognitoidentityCreateIdentityPoolResult) Get(ctx workflow.Context) (*cognitoidentity.IdentityPool, error) {
	var output cognitoidentity.IdentityPool
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CognitoidentityDeleteIdentitiesResult struct {
	Result workflow.Future
}

func (r *CognitoidentityDeleteIdentitiesResult) Get(ctx workflow.Context) (*cognitoidentity.DeleteIdentitiesOutput, error) {
	var output cognitoidentity.DeleteIdentitiesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CognitoidentityDeleteIdentityPoolResult struct {
	Result workflow.Future
}

func (r *CognitoidentityDeleteIdentityPoolResult) Get(ctx workflow.Context) (*cognitoidentity.DeleteIdentityPoolOutput, error) {
	var output cognitoidentity.DeleteIdentityPoolOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CognitoidentityDescribeIdentityResult struct {
	Result workflow.Future
}

func (r *CognitoidentityDescribeIdentityResult) Get(ctx workflow.Context) (*cognitoidentity.IdentityDescription, error) {
	var output cognitoidentity.IdentityDescription
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CognitoidentityDescribeIdentityPoolResult struct {
	Result workflow.Future
}

func (r *CognitoidentityDescribeIdentityPoolResult) Get(ctx workflow.Context) (*cognitoidentity.IdentityPool, error) {
	var output cognitoidentity.IdentityPool
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CognitoidentityGetCredentialsForIdentityResult struct {
	Result workflow.Future
}

func (r *CognitoidentityGetCredentialsForIdentityResult) Get(ctx workflow.Context) (*cognitoidentity.GetCredentialsForIdentityOutput, error) {
	var output cognitoidentity.GetCredentialsForIdentityOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CognitoidentityGetIdResult struct {
	Result workflow.Future
}

func (r *CognitoidentityGetIdResult) Get(ctx workflow.Context) (*cognitoidentity.GetIdOutput, error) {
	var output cognitoidentity.GetIdOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CognitoidentityGetIdentityPoolRolesResult struct {
	Result workflow.Future
}

func (r *CognitoidentityGetIdentityPoolRolesResult) Get(ctx workflow.Context) (*cognitoidentity.GetIdentityPoolRolesOutput, error) {
	var output cognitoidentity.GetIdentityPoolRolesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CognitoidentityGetOpenIdTokenResult struct {
	Result workflow.Future
}

func (r *CognitoidentityGetOpenIdTokenResult) Get(ctx workflow.Context) (*cognitoidentity.GetOpenIdTokenOutput, error) {
	var output cognitoidentity.GetOpenIdTokenOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CognitoidentityGetOpenIdTokenForDeveloperIdentityResult struct {
	Result workflow.Future
}

func (r *CognitoidentityGetOpenIdTokenForDeveloperIdentityResult) Get(ctx workflow.Context) (*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput, error) {
	var output cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CognitoidentityListIdentitiesResult struct {
	Result workflow.Future
}

func (r *CognitoidentityListIdentitiesResult) Get(ctx workflow.Context) (*cognitoidentity.ListIdentitiesOutput, error) {
	var output cognitoidentity.ListIdentitiesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CognitoidentityListIdentityPoolsResult struct {
	Result workflow.Future
}

func (r *CognitoidentityListIdentityPoolsResult) Get(ctx workflow.Context) (*cognitoidentity.ListIdentityPoolsOutput, error) {
	var output cognitoidentity.ListIdentityPoolsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CognitoidentityListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *CognitoidentityListTagsForResourceResult) Get(ctx workflow.Context) (*cognitoidentity.ListTagsForResourceOutput, error) {
	var output cognitoidentity.ListTagsForResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CognitoidentityLookupDeveloperIdentityResult struct {
	Result workflow.Future
}

func (r *CognitoidentityLookupDeveloperIdentityResult) Get(ctx workflow.Context) (*cognitoidentity.LookupDeveloperIdentityOutput, error) {
	var output cognitoidentity.LookupDeveloperIdentityOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CognitoidentityMergeDeveloperIdentitiesResult struct {
	Result workflow.Future
}

func (r *CognitoidentityMergeDeveloperIdentitiesResult) Get(ctx workflow.Context) (*cognitoidentity.MergeDeveloperIdentitiesOutput, error) {
	var output cognitoidentity.MergeDeveloperIdentitiesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CognitoidentitySetIdentityPoolRolesResult struct {
	Result workflow.Future
}

func (r *CognitoidentitySetIdentityPoolRolesResult) Get(ctx workflow.Context) (*cognitoidentity.SetIdentityPoolRolesOutput, error) {
	var output cognitoidentity.SetIdentityPoolRolesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CognitoidentityTagResourceResult struct {
	Result workflow.Future
}

func (r *CognitoidentityTagResourceResult) Get(ctx workflow.Context) (*cognitoidentity.TagResourceOutput, error) {
	var output cognitoidentity.TagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CognitoidentityUnlinkDeveloperIdentityResult struct {
	Result workflow.Future
}

func (r *CognitoidentityUnlinkDeveloperIdentityResult) Get(ctx workflow.Context) (*cognitoidentity.UnlinkDeveloperIdentityOutput, error) {
	var output cognitoidentity.UnlinkDeveloperIdentityOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CognitoidentityUnlinkIdentityResult struct {
	Result workflow.Future
}

func (r *CognitoidentityUnlinkIdentityResult) Get(ctx workflow.Context) (*cognitoidentity.UnlinkIdentityOutput, error) {
	var output cognitoidentity.UnlinkIdentityOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CognitoidentityUntagResourceResult struct {
	Result workflow.Future
}

func (r *CognitoidentityUntagResourceResult) Get(ctx workflow.Context) (*cognitoidentity.UntagResourceOutput, error) {
	var output cognitoidentity.UntagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CognitoidentityUpdateIdentityPoolResult struct {
	Result workflow.Future
}

func (r *CognitoidentityUpdateIdentityPoolResult) Get(ctx workflow.Context) (*cognitoidentity.IdentityPool, error) {
	var output cognitoidentity.IdentityPool
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type CognitoIdentityStub struct {
	activities awsactivities.CognitoIdentityActivities
}

func NewCognitoIdentityStub() CognitoIdentityClient {
	return &CognitoIdentityStub{}
}

func (a *CognitoIdentityStub) CreateIdentityPool(ctx workflow.Context, input *cognitoidentity.CreateIdentityPoolInput) (*cognitoidentity.IdentityPool, error) {
	var output cognitoidentity.IdentityPool
	err := workflow.ExecuteActivity(ctx, a.activities.CreateIdentityPool, input).Get(ctx, &output)
	return &output, err
}

func (a *CognitoIdentityStub) CreateIdentityPoolAsync(ctx workflow.Context, input *cognitoidentity.CreateIdentityPoolInput) *CognitoidentityCreateIdentityPoolResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateIdentityPool, input)
	return &CognitoidentityCreateIdentityPoolResult{Result: future}
}

func (a *CognitoIdentityStub) DeleteIdentities(ctx workflow.Context, input *cognitoidentity.DeleteIdentitiesInput) (*cognitoidentity.DeleteIdentitiesOutput, error) {
	var output cognitoidentity.DeleteIdentitiesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteIdentities, input).Get(ctx, &output)
	return &output, err
}

func (a *CognitoIdentityStub) DeleteIdentitiesAsync(ctx workflow.Context, input *cognitoidentity.DeleteIdentitiesInput) *CognitoidentityDeleteIdentitiesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteIdentities, input)
	return &CognitoidentityDeleteIdentitiesResult{Result: future}
}

func (a *CognitoIdentityStub) DeleteIdentityPool(ctx workflow.Context, input *cognitoidentity.DeleteIdentityPoolInput) (*cognitoidentity.DeleteIdentityPoolOutput, error) {
	var output cognitoidentity.DeleteIdentityPoolOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteIdentityPool, input).Get(ctx, &output)
	return &output, err
}

func (a *CognitoIdentityStub) DeleteIdentityPoolAsync(ctx workflow.Context, input *cognitoidentity.DeleteIdentityPoolInput) *CognitoidentityDeleteIdentityPoolResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteIdentityPool, input)
	return &CognitoidentityDeleteIdentityPoolResult{Result: future}
}

func (a *CognitoIdentityStub) DescribeIdentity(ctx workflow.Context, input *cognitoidentity.DescribeIdentityInput) (*cognitoidentity.IdentityDescription, error) {
	var output cognitoidentity.IdentityDescription
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeIdentity, input).Get(ctx, &output)
	return &output, err
}

func (a *CognitoIdentityStub) DescribeIdentityAsync(ctx workflow.Context, input *cognitoidentity.DescribeIdentityInput) *CognitoidentityDescribeIdentityResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeIdentity, input)
	return &CognitoidentityDescribeIdentityResult{Result: future}
}

func (a *CognitoIdentityStub) DescribeIdentityPool(ctx workflow.Context, input *cognitoidentity.DescribeIdentityPoolInput) (*cognitoidentity.IdentityPool, error) {
	var output cognitoidentity.IdentityPool
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeIdentityPool, input).Get(ctx, &output)
	return &output, err
}

func (a *CognitoIdentityStub) DescribeIdentityPoolAsync(ctx workflow.Context, input *cognitoidentity.DescribeIdentityPoolInput) *CognitoidentityDescribeIdentityPoolResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeIdentityPool, input)
	return &CognitoidentityDescribeIdentityPoolResult{Result: future}
}

func (a *CognitoIdentityStub) GetCredentialsForIdentity(ctx workflow.Context, input *cognitoidentity.GetCredentialsForIdentityInput) (*cognitoidentity.GetCredentialsForIdentityOutput, error) {
	var output cognitoidentity.GetCredentialsForIdentityOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetCredentialsForIdentity, input).Get(ctx, &output)
	return &output, err
}

func (a *CognitoIdentityStub) GetCredentialsForIdentityAsync(ctx workflow.Context, input *cognitoidentity.GetCredentialsForIdentityInput) *CognitoidentityGetCredentialsForIdentityResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetCredentialsForIdentity, input)
	return &CognitoidentityGetCredentialsForIdentityResult{Result: future}
}

func (a *CognitoIdentityStub) GetId(ctx workflow.Context, input *cognitoidentity.GetIdInput) (*cognitoidentity.GetIdOutput, error) {
	var output cognitoidentity.GetIdOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetId, input).Get(ctx, &output)
	return &output, err
}

func (a *CognitoIdentityStub) GetIdAsync(ctx workflow.Context, input *cognitoidentity.GetIdInput) *CognitoidentityGetIdResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetId, input)
	return &CognitoidentityGetIdResult{Result: future}
}

func (a *CognitoIdentityStub) GetIdentityPoolRoles(ctx workflow.Context, input *cognitoidentity.GetIdentityPoolRolesInput) (*cognitoidentity.GetIdentityPoolRolesOutput, error) {
	var output cognitoidentity.GetIdentityPoolRolesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetIdentityPoolRoles, input).Get(ctx, &output)
	return &output, err
}

func (a *CognitoIdentityStub) GetIdentityPoolRolesAsync(ctx workflow.Context, input *cognitoidentity.GetIdentityPoolRolesInput) *CognitoidentityGetIdentityPoolRolesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetIdentityPoolRoles, input)
	return &CognitoidentityGetIdentityPoolRolesResult{Result: future}
}

func (a *CognitoIdentityStub) GetOpenIdToken(ctx workflow.Context, input *cognitoidentity.GetOpenIdTokenInput) (*cognitoidentity.GetOpenIdTokenOutput, error) {
	var output cognitoidentity.GetOpenIdTokenOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetOpenIdToken, input).Get(ctx, &output)
	return &output, err
}

func (a *CognitoIdentityStub) GetOpenIdTokenAsync(ctx workflow.Context, input *cognitoidentity.GetOpenIdTokenInput) *CognitoidentityGetOpenIdTokenResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetOpenIdToken, input)
	return &CognitoidentityGetOpenIdTokenResult{Result: future}
}

func (a *CognitoIdentityStub) GetOpenIdTokenForDeveloperIdentity(ctx workflow.Context, input *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput) (*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput, error) {
	var output cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetOpenIdTokenForDeveloperIdentity, input).Get(ctx, &output)
	return &output, err
}

func (a *CognitoIdentityStub) GetOpenIdTokenForDeveloperIdentityAsync(ctx workflow.Context, input *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput) *CognitoidentityGetOpenIdTokenForDeveloperIdentityResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetOpenIdTokenForDeveloperIdentity, input)
	return &CognitoidentityGetOpenIdTokenForDeveloperIdentityResult{Result: future}
}

func (a *CognitoIdentityStub) ListIdentities(ctx workflow.Context, input *cognitoidentity.ListIdentitiesInput) (*cognitoidentity.ListIdentitiesOutput, error) {
	var output cognitoidentity.ListIdentitiesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListIdentities, input).Get(ctx, &output)
	return &output, err
}

func (a *CognitoIdentityStub) ListIdentitiesAsync(ctx workflow.Context, input *cognitoidentity.ListIdentitiesInput) *CognitoidentityListIdentitiesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListIdentities, input)
	return &CognitoidentityListIdentitiesResult{Result: future}
}

func (a *CognitoIdentityStub) ListIdentityPools(ctx workflow.Context, input *cognitoidentity.ListIdentityPoolsInput) (*cognitoidentity.ListIdentityPoolsOutput, error) {
	var output cognitoidentity.ListIdentityPoolsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListIdentityPools, input).Get(ctx, &output)
	return &output, err
}

func (a *CognitoIdentityStub) ListIdentityPoolsAsync(ctx workflow.Context, input *cognitoidentity.ListIdentityPoolsInput) *CognitoidentityListIdentityPoolsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListIdentityPools, input)
	return &CognitoidentityListIdentityPoolsResult{Result: future}
}

func (a *CognitoIdentityStub) ListTagsForResource(ctx workflow.Context, input *cognitoidentity.ListTagsForResourceInput) (*cognitoidentity.ListTagsForResourceOutput, error) {
	var output cognitoidentity.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
	return &output, err
}

func (a *CognitoIdentityStub) ListTagsForResourceAsync(ctx workflow.Context, input *cognitoidentity.ListTagsForResourceInput) *CognitoidentityListTagsForResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
	return &CognitoidentityListTagsForResourceResult{Result: future}
}

func (a *CognitoIdentityStub) LookupDeveloperIdentity(ctx workflow.Context, input *cognitoidentity.LookupDeveloperIdentityInput) (*cognitoidentity.LookupDeveloperIdentityOutput, error) {
	var output cognitoidentity.LookupDeveloperIdentityOutput
	err := workflow.ExecuteActivity(ctx, a.activities.LookupDeveloperIdentity, input).Get(ctx, &output)
	return &output, err
}

func (a *CognitoIdentityStub) LookupDeveloperIdentityAsync(ctx workflow.Context, input *cognitoidentity.LookupDeveloperIdentityInput) *CognitoidentityLookupDeveloperIdentityResult {
	future := workflow.ExecuteActivity(ctx, a.activities.LookupDeveloperIdentity, input)
	return &CognitoidentityLookupDeveloperIdentityResult{Result: future}
}

func (a *CognitoIdentityStub) MergeDeveloperIdentities(ctx workflow.Context, input *cognitoidentity.MergeDeveloperIdentitiesInput) (*cognitoidentity.MergeDeveloperIdentitiesOutput, error) {
	var output cognitoidentity.MergeDeveloperIdentitiesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.MergeDeveloperIdentities, input).Get(ctx, &output)
	return &output, err
}

func (a *CognitoIdentityStub) MergeDeveloperIdentitiesAsync(ctx workflow.Context, input *cognitoidentity.MergeDeveloperIdentitiesInput) *CognitoidentityMergeDeveloperIdentitiesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.MergeDeveloperIdentities, input)
	return &CognitoidentityMergeDeveloperIdentitiesResult{Result: future}
}

func (a *CognitoIdentityStub) SetIdentityPoolRoles(ctx workflow.Context, input *cognitoidentity.SetIdentityPoolRolesInput) (*cognitoidentity.SetIdentityPoolRolesOutput, error) {
	var output cognitoidentity.SetIdentityPoolRolesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.SetIdentityPoolRoles, input).Get(ctx, &output)
	return &output, err
}

func (a *CognitoIdentityStub) SetIdentityPoolRolesAsync(ctx workflow.Context, input *cognitoidentity.SetIdentityPoolRolesInput) *CognitoidentitySetIdentityPoolRolesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.SetIdentityPoolRoles, input)
	return &CognitoidentitySetIdentityPoolRolesResult{Result: future}
}

func (a *CognitoIdentityStub) TagResource(ctx workflow.Context, input *cognitoidentity.TagResourceInput) (*cognitoidentity.TagResourceOutput, error) {
	var output cognitoidentity.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *CognitoIdentityStub) TagResourceAsync(ctx workflow.Context, input *cognitoidentity.TagResourceInput) *CognitoidentityTagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
	return &CognitoidentityTagResourceResult{Result: future}
}

func (a *CognitoIdentityStub) UnlinkDeveloperIdentity(ctx workflow.Context, input *cognitoidentity.UnlinkDeveloperIdentityInput) (*cognitoidentity.UnlinkDeveloperIdentityOutput, error) {
	var output cognitoidentity.UnlinkDeveloperIdentityOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UnlinkDeveloperIdentity, input).Get(ctx, &output)
	return &output, err
}

func (a *CognitoIdentityStub) UnlinkDeveloperIdentityAsync(ctx workflow.Context, input *cognitoidentity.UnlinkDeveloperIdentityInput) *CognitoidentityUnlinkDeveloperIdentityResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UnlinkDeveloperIdentity, input)
	return &CognitoidentityUnlinkDeveloperIdentityResult{Result: future}
}

func (a *CognitoIdentityStub) UnlinkIdentity(ctx workflow.Context, input *cognitoidentity.UnlinkIdentityInput) (*cognitoidentity.UnlinkIdentityOutput, error) {
	var output cognitoidentity.UnlinkIdentityOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UnlinkIdentity, input).Get(ctx, &output)
	return &output, err
}

func (a *CognitoIdentityStub) UnlinkIdentityAsync(ctx workflow.Context, input *cognitoidentity.UnlinkIdentityInput) *CognitoidentityUnlinkIdentityResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UnlinkIdentity, input)
	return &CognitoidentityUnlinkIdentityResult{Result: future}
}

func (a *CognitoIdentityStub) UntagResource(ctx workflow.Context, input *cognitoidentity.UntagResourceInput) (*cognitoidentity.UntagResourceOutput, error) {
	var output cognitoidentity.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *CognitoIdentityStub) UntagResourceAsync(ctx workflow.Context, input *cognitoidentity.UntagResourceInput) *CognitoidentityUntagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
	return &CognitoidentityUntagResourceResult{Result: future}
}

func (a *CognitoIdentityStub) UpdateIdentityPool(ctx workflow.Context, input *cognitoidentity.IdentityPool) (*cognitoidentity.IdentityPool, error) {
	var output cognitoidentity.IdentityPool
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateIdentityPool, input).Get(ctx, &output)
	return &output, err
}

func (a *CognitoIdentityStub) UpdateIdentityPoolAsync(ctx workflow.Context, input *cognitoidentity.IdentityPool) *CognitoidentityUpdateIdentityPoolResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateIdentityPool, input)
	return &CognitoidentityUpdateIdentityPoolResult{Result: future}
}
