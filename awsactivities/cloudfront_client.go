package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"go.temporal.io/sdk/workflow"
)

type CloudFrontClient interface {
    CreateCachePolicy(ctx workflow.Context, input *cloudfront.CreateCachePolicyInput) (*cloudfront.CreateCachePolicyOutput, error)
    CreateCachePolicyAsync(ctx workflow.Context, input *cloudfront.CreateCachePolicyInput) *CloudfrontCreateCachePolicyResult

    CreateCloudFrontOriginAccessIdentity(ctx workflow.Context, input *cloudfront.CreateCloudFrontOriginAccessIdentityInput) (*cloudfront.CreateCloudFrontOriginAccessIdentityOutput, error)
    CreateCloudFrontOriginAccessIdentityAsync(ctx workflow.Context, input *cloudfront.CreateCloudFrontOriginAccessIdentityInput) *CloudfrontCreateCloudFrontOriginAccessIdentityResult

    CreateDistribution(ctx workflow.Context, input *cloudfront.CreateDistributionInput) (*cloudfront.CreateDistributionOutput, error)
    CreateDistributionAsync(ctx workflow.Context, input *cloudfront.CreateDistributionInput) *CloudfrontCreateDistributionResult

    CreateDistributionWithTags(ctx workflow.Context, input *cloudfront.CreateDistributionWithTagsInput) (*cloudfront.CreateDistributionWithTagsOutput, error)
    CreateDistributionWithTagsAsync(ctx workflow.Context, input *cloudfront.CreateDistributionWithTagsInput) *CloudfrontCreateDistributionWithTagsResult

    CreateFieldLevelEncryptionConfig(ctx workflow.Context, input *cloudfront.CreateFieldLevelEncryptionConfigInput) (*cloudfront.CreateFieldLevelEncryptionConfigOutput, error)
    CreateFieldLevelEncryptionConfigAsync(ctx workflow.Context, input *cloudfront.CreateFieldLevelEncryptionConfigInput) *CloudfrontCreateFieldLevelEncryptionConfigResult

    CreateFieldLevelEncryptionProfile(ctx workflow.Context, input *cloudfront.CreateFieldLevelEncryptionProfileInput) (*cloudfront.CreateFieldLevelEncryptionProfileOutput, error)
    CreateFieldLevelEncryptionProfileAsync(ctx workflow.Context, input *cloudfront.CreateFieldLevelEncryptionProfileInput) *CloudfrontCreateFieldLevelEncryptionProfileResult

    CreateInvalidation(ctx workflow.Context, input *cloudfront.CreateInvalidationInput) (*cloudfront.CreateInvalidationOutput, error)
    CreateInvalidationAsync(ctx workflow.Context, input *cloudfront.CreateInvalidationInput) *CloudfrontCreateInvalidationResult

    CreateMonitoringSubscription(ctx workflow.Context, input *cloudfront.CreateMonitoringSubscriptionInput) (*cloudfront.CreateMonitoringSubscriptionOutput, error)
    CreateMonitoringSubscriptionAsync(ctx workflow.Context, input *cloudfront.CreateMonitoringSubscriptionInput) *CloudfrontCreateMonitoringSubscriptionResult

    CreateOriginRequestPolicy(ctx workflow.Context, input *cloudfront.CreateOriginRequestPolicyInput) (*cloudfront.CreateOriginRequestPolicyOutput, error)
    CreateOriginRequestPolicyAsync(ctx workflow.Context, input *cloudfront.CreateOriginRequestPolicyInput) *CloudfrontCreateOriginRequestPolicyResult

    CreatePublicKey(ctx workflow.Context, input *cloudfront.CreatePublicKeyInput) (*cloudfront.CreatePublicKeyOutput, error)
    CreatePublicKeyAsync(ctx workflow.Context, input *cloudfront.CreatePublicKeyInput) *CloudfrontCreatePublicKeyResult

    CreateStreamingDistribution(ctx workflow.Context, input *cloudfront.CreateStreamingDistributionInput) (*cloudfront.CreateStreamingDistributionOutput, error)
    CreateStreamingDistributionAsync(ctx workflow.Context, input *cloudfront.CreateStreamingDistributionInput) *CloudfrontCreateStreamingDistributionResult

    CreateStreamingDistributionWithTags(ctx workflow.Context, input *cloudfront.CreateStreamingDistributionWithTagsInput) (*cloudfront.CreateStreamingDistributionWithTagsOutput, error)
    CreateStreamingDistributionWithTagsAsync(ctx workflow.Context, input *cloudfront.CreateStreamingDistributionWithTagsInput) *CloudfrontCreateStreamingDistributionWithTagsResult

    DeleteCachePolicy(ctx workflow.Context, input *cloudfront.DeleteCachePolicyInput) (*cloudfront.DeleteCachePolicyOutput, error)
    DeleteCachePolicyAsync(ctx workflow.Context, input *cloudfront.DeleteCachePolicyInput) *CloudfrontDeleteCachePolicyResult

    DeleteCloudFrontOriginAccessIdentity(ctx workflow.Context, input *cloudfront.DeleteCloudFrontOriginAccessIdentityInput) (*cloudfront.DeleteCloudFrontOriginAccessIdentityOutput, error)
    DeleteCloudFrontOriginAccessIdentityAsync(ctx workflow.Context, input *cloudfront.DeleteCloudFrontOriginAccessIdentityInput) *CloudfrontDeleteCloudFrontOriginAccessIdentityResult

    DeleteDistribution(ctx workflow.Context, input *cloudfront.DeleteDistributionInput) (*cloudfront.DeleteDistributionOutput, error)
    DeleteDistributionAsync(ctx workflow.Context, input *cloudfront.DeleteDistributionInput) *CloudfrontDeleteDistributionResult

    DeleteFieldLevelEncryptionConfig(ctx workflow.Context, input *cloudfront.DeleteFieldLevelEncryptionConfigInput) (*cloudfront.DeleteFieldLevelEncryptionConfigOutput, error)
    DeleteFieldLevelEncryptionConfigAsync(ctx workflow.Context, input *cloudfront.DeleteFieldLevelEncryptionConfigInput) *CloudfrontDeleteFieldLevelEncryptionConfigResult

    DeleteFieldLevelEncryptionProfile(ctx workflow.Context, input *cloudfront.DeleteFieldLevelEncryptionProfileInput) (*cloudfront.DeleteFieldLevelEncryptionProfileOutput, error)
    DeleteFieldLevelEncryptionProfileAsync(ctx workflow.Context, input *cloudfront.DeleteFieldLevelEncryptionProfileInput) *CloudfrontDeleteFieldLevelEncryptionProfileResult

    DeleteMonitoringSubscription(ctx workflow.Context, input *cloudfront.DeleteMonitoringSubscriptionInput) (*cloudfront.DeleteMonitoringSubscriptionOutput, error)
    DeleteMonitoringSubscriptionAsync(ctx workflow.Context, input *cloudfront.DeleteMonitoringSubscriptionInput) *CloudfrontDeleteMonitoringSubscriptionResult

    DeleteOriginRequestPolicy(ctx workflow.Context, input *cloudfront.DeleteOriginRequestPolicyInput) (*cloudfront.DeleteOriginRequestPolicyOutput, error)
    DeleteOriginRequestPolicyAsync(ctx workflow.Context, input *cloudfront.DeleteOriginRequestPolicyInput) *CloudfrontDeleteOriginRequestPolicyResult

    DeletePublicKey(ctx workflow.Context, input *cloudfront.DeletePublicKeyInput) (*cloudfront.DeletePublicKeyOutput, error)
    DeletePublicKeyAsync(ctx workflow.Context, input *cloudfront.DeletePublicKeyInput) *CloudfrontDeletePublicKeyResult

    DeleteStreamingDistribution(ctx workflow.Context, input *cloudfront.DeleteStreamingDistributionInput) (*cloudfront.DeleteStreamingDistributionOutput, error)
    DeleteStreamingDistributionAsync(ctx workflow.Context, input *cloudfront.DeleteStreamingDistributionInput) *CloudfrontDeleteStreamingDistributionResult

    GetCachePolicy(ctx workflow.Context, input *cloudfront.GetCachePolicyInput) (*cloudfront.GetCachePolicyOutput, error)
    GetCachePolicyAsync(ctx workflow.Context, input *cloudfront.GetCachePolicyInput) *CloudfrontGetCachePolicyResult

    GetCachePolicyConfig(ctx workflow.Context, input *cloudfront.GetCachePolicyConfigInput) (*cloudfront.GetCachePolicyConfigOutput, error)
    GetCachePolicyConfigAsync(ctx workflow.Context, input *cloudfront.GetCachePolicyConfigInput) *CloudfrontGetCachePolicyConfigResult

    GetCloudFrontOriginAccessIdentity(ctx workflow.Context, input *cloudfront.GetCloudFrontOriginAccessIdentityInput) (*cloudfront.GetCloudFrontOriginAccessIdentityOutput, error)
    GetCloudFrontOriginAccessIdentityAsync(ctx workflow.Context, input *cloudfront.GetCloudFrontOriginAccessIdentityInput) *CloudfrontGetCloudFrontOriginAccessIdentityResult

    GetCloudFrontOriginAccessIdentityConfig(ctx workflow.Context, input *cloudfront.GetCloudFrontOriginAccessIdentityConfigInput) (*cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput, error)
    GetCloudFrontOriginAccessIdentityConfigAsync(ctx workflow.Context, input *cloudfront.GetCloudFrontOriginAccessIdentityConfigInput) *CloudfrontGetCloudFrontOriginAccessIdentityConfigResult

    GetDistribution(ctx workflow.Context, input *cloudfront.GetDistributionInput) (*cloudfront.GetDistributionOutput, error)
    GetDistributionAsync(ctx workflow.Context, input *cloudfront.GetDistributionInput) *CloudfrontGetDistributionResult

    GetDistributionConfig(ctx workflow.Context, input *cloudfront.GetDistributionConfigInput) (*cloudfront.GetDistributionConfigOutput, error)
    GetDistributionConfigAsync(ctx workflow.Context, input *cloudfront.GetDistributionConfigInput) *CloudfrontGetDistributionConfigResult

    GetFieldLevelEncryption(ctx workflow.Context, input *cloudfront.GetFieldLevelEncryptionInput) (*cloudfront.GetFieldLevelEncryptionOutput, error)
    GetFieldLevelEncryptionAsync(ctx workflow.Context, input *cloudfront.GetFieldLevelEncryptionInput) *CloudfrontGetFieldLevelEncryptionResult

    GetFieldLevelEncryptionConfig(ctx workflow.Context, input *cloudfront.GetFieldLevelEncryptionConfigInput) (*cloudfront.GetFieldLevelEncryptionConfigOutput, error)
    GetFieldLevelEncryptionConfigAsync(ctx workflow.Context, input *cloudfront.GetFieldLevelEncryptionConfigInput) *CloudfrontGetFieldLevelEncryptionConfigResult

    GetFieldLevelEncryptionProfile(ctx workflow.Context, input *cloudfront.GetFieldLevelEncryptionProfileInput) (*cloudfront.GetFieldLevelEncryptionProfileOutput, error)
    GetFieldLevelEncryptionProfileAsync(ctx workflow.Context, input *cloudfront.GetFieldLevelEncryptionProfileInput) *CloudfrontGetFieldLevelEncryptionProfileResult

    GetFieldLevelEncryptionProfileConfig(ctx workflow.Context, input *cloudfront.GetFieldLevelEncryptionProfileConfigInput) (*cloudfront.GetFieldLevelEncryptionProfileConfigOutput, error)
    GetFieldLevelEncryptionProfileConfigAsync(ctx workflow.Context, input *cloudfront.GetFieldLevelEncryptionProfileConfigInput) *CloudfrontGetFieldLevelEncryptionProfileConfigResult

    GetInvalidation(ctx workflow.Context, input *cloudfront.GetInvalidationInput) (*cloudfront.GetInvalidationOutput, error)
    GetInvalidationAsync(ctx workflow.Context, input *cloudfront.GetInvalidationInput) *CloudfrontGetInvalidationResult

    GetMonitoringSubscription(ctx workflow.Context, input *cloudfront.GetMonitoringSubscriptionInput) (*cloudfront.GetMonitoringSubscriptionOutput, error)
    GetMonitoringSubscriptionAsync(ctx workflow.Context, input *cloudfront.GetMonitoringSubscriptionInput) *CloudfrontGetMonitoringSubscriptionResult

    GetOriginRequestPolicy(ctx workflow.Context, input *cloudfront.GetOriginRequestPolicyInput) (*cloudfront.GetOriginRequestPolicyOutput, error)
    GetOriginRequestPolicyAsync(ctx workflow.Context, input *cloudfront.GetOriginRequestPolicyInput) *CloudfrontGetOriginRequestPolicyResult

    GetOriginRequestPolicyConfig(ctx workflow.Context, input *cloudfront.GetOriginRequestPolicyConfigInput) (*cloudfront.GetOriginRequestPolicyConfigOutput, error)
    GetOriginRequestPolicyConfigAsync(ctx workflow.Context, input *cloudfront.GetOriginRequestPolicyConfigInput) *CloudfrontGetOriginRequestPolicyConfigResult

    GetPublicKey(ctx workflow.Context, input *cloudfront.GetPublicKeyInput) (*cloudfront.GetPublicKeyOutput, error)
    GetPublicKeyAsync(ctx workflow.Context, input *cloudfront.GetPublicKeyInput) *CloudfrontGetPublicKeyResult

    GetPublicKeyConfig(ctx workflow.Context, input *cloudfront.GetPublicKeyConfigInput) (*cloudfront.GetPublicKeyConfigOutput, error)
    GetPublicKeyConfigAsync(ctx workflow.Context, input *cloudfront.GetPublicKeyConfigInput) *CloudfrontGetPublicKeyConfigResult

    GetStreamingDistribution(ctx workflow.Context, input *cloudfront.GetStreamingDistributionInput) (*cloudfront.GetStreamingDistributionOutput, error)
    GetStreamingDistributionAsync(ctx workflow.Context, input *cloudfront.GetStreamingDistributionInput) *CloudfrontGetStreamingDistributionResult

    GetStreamingDistributionConfig(ctx workflow.Context, input *cloudfront.GetStreamingDistributionConfigInput) (*cloudfront.GetStreamingDistributionConfigOutput, error)
    GetStreamingDistributionConfigAsync(ctx workflow.Context, input *cloudfront.GetStreamingDistributionConfigInput) *CloudfrontGetStreamingDistributionConfigResult

    ListCachePolicies(ctx workflow.Context, input *cloudfront.ListCachePoliciesInput) (*cloudfront.ListCachePoliciesOutput, error)
    ListCachePoliciesAsync(ctx workflow.Context, input *cloudfront.ListCachePoliciesInput) *CloudfrontListCachePoliciesResult

    ListCloudFrontOriginAccessIdentities(ctx workflow.Context, input *cloudfront.ListCloudFrontOriginAccessIdentitiesInput) (*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, error)
    ListCloudFrontOriginAccessIdentitiesAsync(ctx workflow.Context, input *cloudfront.ListCloudFrontOriginAccessIdentitiesInput) *CloudfrontListCloudFrontOriginAccessIdentitiesResult

    ListDistributions(ctx workflow.Context, input *cloudfront.ListDistributionsInput) (*cloudfront.ListDistributionsOutput, error)
    ListDistributionsAsync(ctx workflow.Context, input *cloudfront.ListDistributionsInput) *CloudfrontListDistributionsResult

    ListDistributionsByCachePolicyId(ctx workflow.Context, input *cloudfront.ListDistributionsByCachePolicyIdInput) (*cloudfront.ListDistributionsByCachePolicyIdOutput, error)
    ListDistributionsByCachePolicyIdAsync(ctx workflow.Context, input *cloudfront.ListDistributionsByCachePolicyIdInput) *CloudfrontListDistributionsByCachePolicyIdResult

    ListDistributionsByOriginRequestPolicyId(ctx workflow.Context, input *cloudfront.ListDistributionsByOriginRequestPolicyIdInput) (*cloudfront.ListDistributionsByOriginRequestPolicyIdOutput, error)
    ListDistributionsByOriginRequestPolicyIdAsync(ctx workflow.Context, input *cloudfront.ListDistributionsByOriginRequestPolicyIdInput) *CloudfrontListDistributionsByOriginRequestPolicyIdResult

    ListDistributionsByWebACLId(ctx workflow.Context, input *cloudfront.ListDistributionsByWebACLIdInput) (*cloudfront.ListDistributionsByWebACLIdOutput, error)
    ListDistributionsByWebACLIdAsync(ctx workflow.Context, input *cloudfront.ListDistributionsByWebACLIdInput) *CloudfrontListDistributionsByWebACLIdResult

    ListFieldLevelEncryptionConfigs(ctx workflow.Context, input *cloudfront.ListFieldLevelEncryptionConfigsInput) (*cloudfront.ListFieldLevelEncryptionConfigsOutput, error)
    ListFieldLevelEncryptionConfigsAsync(ctx workflow.Context, input *cloudfront.ListFieldLevelEncryptionConfigsInput) *CloudfrontListFieldLevelEncryptionConfigsResult

    ListFieldLevelEncryptionProfiles(ctx workflow.Context, input *cloudfront.ListFieldLevelEncryptionProfilesInput) (*cloudfront.ListFieldLevelEncryptionProfilesOutput, error)
    ListFieldLevelEncryptionProfilesAsync(ctx workflow.Context, input *cloudfront.ListFieldLevelEncryptionProfilesInput) *CloudfrontListFieldLevelEncryptionProfilesResult

    ListInvalidations(ctx workflow.Context, input *cloudfront.ListInvalidationsInput) (*cloudfront.ListInvalidationsOutput, error)
    ListInvalidationsAsync(ctx workflow.Context, input *cloudfront.ListInvalidationsInput) *CloudfrontListInvalidationsResult

    ListOriginRequestPolicies(ctx workflow.Context, input *cloudfront.ListOriginRequestPoliciesInput) (*cloudfront.ListOriginRequestPoliciesOutput, error)
    ListOriginRequestPoliciesAsync(ctx workflow.Context, input *cloudfront.ListOriginRequestPoliciesInput) *CloudfrontListOriginRequestPoliciesResult

    ListPublicKeys(ctx workflow.Context, input *cloudfront.ListPublicKeysInput) (*cloudfront.ListPublicKeysOutput, error)
    ListPublicKeysAsync(ctx workflow.Context, input *cloudfront.ListPublicKeysInput) *CloudfrontListPublicKeysResult

    ListStreamingDistributions(ctx workflow.Context, input *cloudfront.ListStreamingDistributionsInput) (*cloudfront.ListStreamingDistributionsOutput, error)
    ListStreamingDistributionsAsync(ctx workflow.Context, input *cloudfront.ListStreamingDistributionsInput) *CloudfrontListStreamingDistributionsResult

    ListTagsForResource(ctx workflow.Context, input *cloudfront.ListTagsForResourceInput) (*cloudfront.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *cloudfront.ListTagsForResourceInput) *CloudfrontListTagsForResourceResult

    TagResource(ctx workflow.Context, input *cloudfront.TagResourceInput) (*cloudfront.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *cloudfront.TagResourceInput) *CloudfrontTagResourceResult

    UntagResource(ctx workflow.Context, input *cloudfront.UntagResourceInput) (*cloudfront.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *cloudfront.UntagResourceInput) *CloudfrontUntagResourceResult

    UpdateCachePolicy(ctx workflow.Context, input *cloudfront.UpdateCachePolicyInput) (*cloudfront.UpdateCachePolicyOutput, error)
    UpdateCachePolicyAsync(ctx workflow.Context, input *cloudfront.UpdateCachePolicyInput) *CloudfrontUpdateCachePolicyResult

    UpdateCloudFrontOriginAccessIdentity(ctx workflow.Context, input *cloudfront.UpdateCloudFrontOriginAccessIdentityInput) (*cloudfront.UpdateCloudFrontOriginAccessIdentityOutput, error)
    UpdateCloudFrontOriginAccessIdentityAsync(ctx workflow.Context, input *cloudfront.UpdateCloudFrontOriginAccessIdentityInput) *CloudfrontUpdateCloudFrontOriginAccessIdentityResult

    UpdateDistribution(ctx workflow.Context, input *cloudfront.UpdateDistributionInput) (*cloudfront.UpdateDistributionOutput, error)
    UpdateDistributionAsync(ctx workflow.Context, input *cloudfront.UpdateDistributionInput) *CloudfrontUpdateDistributionResult

    UpdateFieldLevelEncryptionConfig(ctx workflow.Context, input *cloudfront.UpdateFieldLevelEncryptionConfigInput) (*cloudfront.UpdateFieldLevelEncryptionConfigOutput, error)
    UpdateFieldLevelEncryptionConfigAsync(ctx workflow.Context, input *cloudfront.UpdateFieldLevelEncryptionConfigInput) *CloudfrontUpdateFieldLevelEncryptionConfigResult

    UpdateFieldLevelEncryptionProfile(ctx workflow.Context, input *cloudfront.UpdateFieldLevelEncryptionProfileInput) (*cloudfront.UpdateFieldLevelEncryptionProfileOutput, error)
    UpdateFieldLevelEncryptionProfileAsync(ctx workflow.Context, input *cloudfront.UpdateFieldLevelEncryptionProfileInput) *CloudfrontUpdateFieldLevelEncryptionProfileResult

    UpdateOriginRequestPolicy(ctx workflow.Context, input *cloudfront.UpdateOriginRequestPolicyInput) (*cloudfront.UpdateOriginRequestPolicyOutput, error)
    UpdateOriginRequestPolicyAsync(ctx workflow.Context, input *cloudfront.UpdateOriginRequestPolicyInput) *CloudfrontUpdateOriginRequestPolicyResult

    UpdatePublicKey(ctx workflow.Context, input *cloudfront.UpdatePublicKeyInput) (*cloudfront.UpdatePublicKeyOutput, error)
    UpdatePublicKeyAsync(ctx workflow.Context, input *cloudfront.UpdatePublicKeyInput) *CloudfrontUpdatePublicKeyResult

    UpdateStreamingDistribution(ctx workflow.Context, input *cloudfront.UpdateStreamingDistributionInput) (*cloudfront.UpdateStreamingDistributionOutput, error)
    UpdateStreamingDistributionAsync(ctx workflow.Context, input *cloudfront.UpdateStreamingDistributionInput) *CloudfrontUpdateStreamingDistributionResult

    WaitUntilDistributionDeployed(ctx workflow.Context, input *cloudfront.GetDistributionInput) error
    WaitUntilInvalidationCompleted(ctx workflow.Context, input *cloudfront.GetInvalidationInput) error
    WaitUntilStreamingDistributionDeployed(ctx workflow.Context, input *cloudfront.GetStreamingDistributionInput) error}
type CloudfrontCreateCachePolicyResult struct {
	Result workflow.Future
}

func (r *CloudfrontCreateCachePolicyResult) Get(ctx workflow.Context) (*cloudfront.CreateCachePolicyOutput, error) {
    var output cloudfront.CreateCachePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontCreateCloudFrontOriginAccessIdentityResult struct {
	Result workflow.Future
}

func (r *CloudfrontCreateCloudFrontOriginAccessIdentityResult) Get(ctx workflow.Context) (*cloudfront.CreateCloudFrontOriginAccessIdentityOutput, error) {
    var output cloudfront.CreateCloudFrontOriginAccessIdentityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontCreateDistributionResult struct {
	Result workflow.Future
}

func (r *CloudfrontCreateDistributionResult) Get(ctx workflow.Context) (*cloudfront.CreateDistributionOutput, error) {
    var output cloudfront.CreateDistributionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontCreateDistributionWithTagsResult struct {
	Result workflow.Future
}

func (r *CloudfrontCreateDistributionWithTagsResult) Get(ctx workflow.Context) (*cloudfront.CreateDistributionWithTagsOutput, error) {
    var output cloudfront.CreateDistributionWithTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontCreateFieldLevelEncryptionConfigResult struct {
	Result workflow.Future
}

func (r *CloudfrontCreateFieldLevelEncryptionConfigResult) Get(ctx workflow.Context) (*cloudfront.CreateFieldLevelEncryptionConfigOutput, error) {
    var output cloudfront.CreateFieldLevelEncryptionConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontCreateFieldLevelEncryptionProfileResult struct {
	Result workflow.Future
}

func (r *CloudfrontCreateFieldLevelEncryptionProfileResult) Get(ctx workflow.Context) (*cloudfront.CreateFieldLevelEncryptionProfileOutput, error) {
    var output cloudfront.CreateFieldLevelEncryptionProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontCreateInvalidationResult struct {
	Result workflow.Future
}

func (r *CloudfrontCreateInvalidationResult) Get(ctx workflow.Context) (*cloudfront.CreateInvalidationOutput, error) {
    var output cloudfront.CreateInvalidationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontCreateMonitoringSubscriptionResult struct {
	Result workflow.Future
}

func (r *CloudfrontCreateMonitoringSubscriptionResult) Get(ctx workflow.Context) (*cloudfront.CreateMonitoringSubscriptionOutput, error) {
    var output cloudfront.CreateMonitoringSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontCreateOriginRequestPolicyResult struct {
	Result workflow.Future
}

func (r *CloudfrontCreateOriginRequestPolicyResult) Get(ctx workflow.Context) (*cloudfront.CreateOriginRequestPolicyOutput, error) {
    var output cloudfront.CreateOriginRequestPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontCreatePublicKeyResult struct {
	Result workflow.Future
}

func (r *CloudfrontCreatePublicKeyResult) Get(ctx workflow.Context) (*cloudfront.CreatePublicKeyOutput, error) {
    var output cloudfront.CreatePublicKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontCreateStreamingDistributionResult struct {
	Result workflow.Future
}

func (r *CloudfrontCreateStreamingDistributionResult) Get(ctx workflow.Context) (*cloudfront.CreateStreamingDistributionOutput, error) {
    var output cloudfront.CreateStreamingDistributionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontCreateStreamingDistributionWithTagsResult struct {
	Result workflow.Future
}

func (r *CloudfrontCreateStreamingDistributionWithTagsResult) Get(ctx workflow.Context) (*cloudfront.CreateStreamingDistributionWithTagsOutput, error) {
    var output cloudfront.CreateStreamingDistributionWithTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontDeleteCachePolicyResult struct {
	Result workflow.Future
}

func (r *CloudfrontDeleteCachePolicyResult) Get(ctx workflow.Context) (*cloudfront.DeleteCachePolicyOutput, error) {
    var output cloudfront.DeleteCachePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontDeleteCloudFrontOriginAccessIdentityResult struct {
	Result workflow.Future
}

func (r *CloudfrontDeleteCloudFrontOriginAccessIdentityResult) Get(ctx workflow.Context) (*cloudfront.DeleteCloudFrontOriginAccessIdentityOutput, error) {
    var output cloudfront.DeleteCloudFrontOriginAccessIdentityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontDeleteDistributionResult struct {
	Result workflow.Future
}

func (r *CloudfrontDeleteDistributionResult) Get(ctx workflow.Context) (*cloudfront.DeleteDistributionOutput, error) {
    var output cloudfront.DeleteDistributionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontDeleteFieldLevelEncryptionConfigResult struct {
	Result workflow.Future
}

func (r *CloudfrontDeleteFieldLevelEncryptionConfigResult) Get(ctx workflow.Context) (*cloudfront.DeleteFieldLevelEncryptionConfigOutput, error) {
    var output cloudfront.DeleteFieldLevelEncryptionConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontDeleteFieldLevelEncryptionProfileResult struct {
	Result workflow.Future
}

func (r *CloudfrontDeleteFieldLevelEncryptionProfileResult) Get(ctx workflow.Context) (*cloudfront.DeleteFieldLevelEncryptionProfileOutput, error) {
    var output cloudfront.DeleteFieldLevelEncryptionProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontDeleteMonitoringSubscriptionResult struct {
	Result workflow.Future
}

func (r *CloudfrontDeleteMonitoringSubscriptionResult) Get(ctx workflow.Context) (*cloudfront.DeleteMonitoringSubscriptionOutput, error) {
    var output cloudfront.DeleteMonitoringSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontDeleteOriginRequestPolicyResult struct {
	Result workflow.Future
}

func (r *CloudfrontDeleteOriginRequestPolicyResult) Get(ctx workflow.Context) (*cloudfront.DeleteOriginRequestPolicyOutput, error) {
    var output cloudfront.DeleteOriginRequestPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontDeletePublicKeyResult struct {
	Result workflow.Future
}

func (r *CloudfrontDeletePublicKeyResult) Get(ctx workflow.Context) (*cloudfront.DeletePublicKeyOutput, error) {
    var output cloudfront.DeletePublicKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontDeleteStreamingDistributionResult struct {
	Result workflow.Future
}

func (r *CloudfrontDeleteStreamingDistributionResult) Get(ctx workflow.Context) (*cloudfront.DeleteStreamingDistributionOutput, error) {
    var output cloudfront.DeleteStreamingDistributionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontGetCachePolicyResult struct {
	Result workflow.Future
}

func (r *CloudfrontGetCachePolicyResult) Get(ctx workflow.Context) (*cloudfront.GetCachePolicyOutput, error) {
    var output cloudfront.GetCachePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontGetCachePolicyConfigResult struct {
	Result workflow.Future
}

func (r *CloudfrontGetCachePolicyConfigResult) Get(ctx workflow.Context) (*cloudfront.GetCachePolicyConfigOutput, error) {
    var output cloudfront.GetCachePolicyConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontGetCloudFrontOriginAccessIdentityResult struct {
	Result workflow.Future
}

func (r *CloudfrontGetCloudFrontOriginAccessIdentityResult) Get(ctx workflow.Context) (*cloudfront.GetCloudFrontOriginAccessIdentityOutput, error) {
    var output cloudfront.GetCloudFrontOriginAccessIdentityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontGetCloudFrontOriginAccessIdentityConfigResult struct {
	Result workflow.Future
}

func (r *CloudfrontGetCloudFrontOriginAccessIdentityConfigResult) Get(ctx workflow.Context) (*cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput, error) {
    var output cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontGetDistributionResult struct {
	Result workflow.Future
}

func (r *CloudfrontGetDistributionResult) Get(ctx workflow.Context) (*cloudfront.GetDistributionOutput, error) {
    var output cloudfront.GetDistributionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontGetDistributionConfigResult struct {
	Result workflow.Future
}

func (r *CloudfrontGetDistributionConfigResult) Get(ctx workflow.Context) (*cloudfront.GetDistributionConfigOutput, error) {
    var output cloudfront.GetDistributionConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontGetFieldLevelEncryptionResult struct {
	Result workflow.Future
}

func (r *CloudfrontGetFieldLevelEncryptionResult) Get(ctx workflow.Context) (*cloudfront.GetFieldLevelEncryptionOutput, error) {
    var output cloudfront.GetFieldLevelEncryptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontGetFieldLevelEncryptionConfigResult struct {
	Result workflow.Future
}

func (r *CloudfrontGetFieldLevelEncryptionConfigResult) Get(ctx workflow.Context) (*cloudfront.GetFieldLevelEncryptionConfigOutput, error) {
    var output cloudfront.GetFieldLevelEncryptionConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontGetFieldLevelEncryptionProfileResult struct {
	Result workflow.Future
}

func (r *CloudfrontGetFieldLevelEncryptionProfileResult) Get(ctx workflow.Context) (*cloudfront.GetFieldLevelEncryptionProfileOutput, error) {
    var output cloudfront.GetFieldLevelEncryptionProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontGetFieldLevelEncryptionProfileConfigResult struct {
	Result workflow.Future
}

func (r *CloudfrontGetFieldLevelEncryptionProfileConfigResult) Get(ctx workflow.Context) (*cloudfront.GetFieldLevelEncryptionProfileConfigOutput, error) {
    var output cloudfront.GetFieldLevelEncryptionProfileConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontGetInvalidationResult struct {
	Result workflow.Future
}

func (r *CloudfrontGetInvalidationResult) Get(ctx workflow.Context) (*cloudfront.GetInvalidationOutput, error) {
    var output cloudfront.GetInvalidationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontGetMonitoringSubscriptionResult struct {
	Result workflow.Future
}

func (r *CloudfrontGetMonitoringSubscriptionResult) Get(ctx workflow.Context) (*cloudfront.GetMonitoringSubscriptionOutput, error) {
    var output cloudfront.GetMonitoringSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontGetOriginRequestPolicyResult struct {
	Result workflow.Future
}

func (r *CloudfrontGetOriginRequestPolicyResult) Get(ctx workflow.Context) (*cloudfront.GetOriginRequestPolicyOutput, error) {
    var output cloudfront.GetOriginRequestPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontGetOriginRequestPolicyConfigResult struct {
	Result workflow.Future
}

func (r *CloudfrontGetOriginRequestPolicyConfigResult) Get(ctx workflow.Context) (*cloudfront.GetOriginRequestPolicyConfigOutput, error) {
    var output cloudfront.GetOriginRequestPolicyConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontGetPublicKeyResult struct {
	Result workflow.Future
}

func (r *CloudfrontGetPublicKeyResult) Get(ctx workflow.Context) (*cloudfront.GetPublicKeyOutput, error) {
    var output cloudfront.GetPublicKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontGetPublicKeyConfigResult struct {
	Result workflow.Future
}

func (r *CloudfrontGetPublicKeyConfigResult) Get(ctx workflow.Context) (*cloudfront.GetPublicKeyConfigOutput, error) {
    var output cloudfront.GetPublicKeyConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontGetStreamingDistributionResult struct {
	Result workflow.Future
}

func (r *CloudfrontGetStreamingDistributionResult) Get(ctx workflow.Context) (*cloudfront.GetStreamingDistributionOutput, error) {
    var output cloudfront.GetStreamingDistributionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontGetStreamingDistributionConfigResult struct {
	Result workflow.Future
}

func (r *CloudfrontGetStreamingDistributionConfigResult) Get(ctx workflow.Context) (*cloudfront.GetStreamingDistributionConfigOutput, error) {
    var output cloudfront.GetStreamingDistributionConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontListCachePoliciesResult struct {
	Result workflow.Future
}

func (r *CloudfrontListCachePoliciesResult) Get(ctx workflow.Context) (*cloudfront.ListCachePoliciesOutput, error) {
    var output cloudfront.ListCachePoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontListCloudFrontOriginAccessIdentitiesResult struct {
	Result workflow.Future
}

func (r *CloudfrontListCloudFrontOriginAccessIdentitiesResult) Get(ctx workflow.Context) (*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, error) {
    var output cloudfront.ListCloudFrontOriginAccessIdentitiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontListDistributionsResult struct {
	Result workflow.Future
}

func (r *CloudfrontListDistributionsResult) Get(ctx workflow.Context) (*cloudfront.ListDistributionsOutput, error) {
    var output cloudfront.ListDistributionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontListDistributionsByCachePolicyIdResult struct {
	Result workflow.Future
}

func (r *CloudfrontListDistributionsByCachePolicyIdResult) Get(ctx workflow.Context) (*cloudfront.ListDistributionsByCachePolicyIdOutput, error) {
    var output cloudfront.ListDistributionsByCachePolicyIdOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontListDistributionsByOriginRequestPolicyIdResult struct {
	Result workflow.Future
}

func (r *CloudfrontListDistributionsByOriginRequestPolicyIdResult) Get(ctx workflow.Context) (*cloudfront.ListDistributionsByOriginRequestPolicyIdOutput, error) {
    var output cloudfront.ListDistributionsByOriginRequestPolicyIdOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontListDistributionsByWebACLIdResult struct {
	Result workflow.Future
}

func (r *CloudfrontListDistributionsByWebACLIdResult) Get(ctx workflow.Context) (*cloudfront.ListDistributionsByWebACLIdOutput, error) {
    var output cloudfront.ListDistributionsByWebACLIdOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontListFieldLevelEncryptionConfigsResult struct {
	Result workflow.Future
}

func (r *CloudfrontListFieldLevelEncryptionConfigsResult) Get(ctx workflow.Context) (*cloudfront.ListFieldLevelEncryptionConfigsOutput, error) {
    var output cloudfront.ListFieldLevelEncryptionConfigsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontListFieldLevelEncryptionProfilesResult struct {
	Result workflow.Future
}

func (r *CloudfrontListFieldLevelEncryptionProfilesResult) Get(ctx workflow.Context) (*cloudfront.ListFieldLevelEncryptionProfilesOutput, error) {
    var output cloudfront.ListFieldLevelEncryptionProfilesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontListInvalidationsResult struct {
	Result workflow.Future
}

func (r *CloudfrontListInvalidationsResult) Get(ctx workflow.Context) (*cloudfront.ListInvalidationsOutput, error) {
    var output cloudfront.ListInvalidationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontListOriginRequestPoliciesResult struct {
	Result workflow.Future
}

func (r *CloudfrontListOriginRequestPoliciesResult) Get(ctx workflow.Context) (*cloudfront.ListOriginRequestPoliciesOutput, error) {
    var output cloudfront.ListOriginRequestPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontListPublicKeysResult struct {
	Result workflow.Future
}

func (r *CloudfrontListPublicKeysResult) Get(ctx workflow.Context) (*cloudfront.ListPublicKeysOutput, error) {
    var output cloudfront.ListPublicKeysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontListStreamingDistributionsResult struct {
	Result workflow.Future
}

func (r *CloudfrontListStreamingDistributionsResult) Get(ctx workflow.Context) (*cloudfront.ListStreamingDistributionsOutput, error) {
    var output cloudfront.ListStreamingDistributionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *CloudfrontListTagsForResourceResult) Get(ctx workflow.Context) (*cloudfront.ListTagsForResourceOutput, error) {
    var output cloudfront.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontTagResourceResult struct {
	Result workflow.Future
}

func (r *CloudfrontTagResourceResult) Get(ctx workflow.Context) (*cloudfront.TagResourceOutput, error) {
    var output cloudfront.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontUntagResourceResult struct {
	Result workflow.Future
}

func (r *CloudfrontUntagResourceResult) Get(ctx workflow.Context) (*cloudfront.UntagResourceOutput, error) {
    var output cloudfront.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontUpdateCachePolicyResult struct {
	Result workflow.Future
}

func (r *CloudfrontUpdateCachePolicyResult) Get(ctx workflow.Context) (*cloudfront.UpdateCachePolicyOutput, error) {
    var output cloudfront.UpdateCachePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontUpdateCloudFrontOriginAccessIdentityResult struct {
	Result workflow.Future
}

func (r *CloudfrontUpdateCloudFrontOriginAccessIdentityResult) Get(ctx workflow.Context) (*cloudfront.UpdateCloudFrontOriginAccessIdentityOutput, error) {
    var output cloudfront.UpdateCloudFrontOriginAccessIdentityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontUpdateDistributionResult struct {
	Result workflow.Future
}

func (r *CloudfrontUpdateDistributionResult) Get(ctx workflow.Context) (*cloudfront.UpdateDistributionOutput, error) {
    var output cloudfront.UpdateDistributionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontUpdateFieldLevelEncryptionConfigResult struct {
	Result workflow.Future
}

func (r *CloudfrontUpdateFieldLevelEncryptionConfigResult) Get(ctx workflow.Context) (*cloudfront.UpdateFieldLevelEncryptionConfigOutput, error) {
    var output cloudfront.UpdateFieldLevelEncryptionConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontUpdateFieldLevelEncryptionProfileResult struct {
	Result workflow.Future
}

func (r *CloudfrontUpdateFieldLevelEncryptionProfileResult) Get(ctx workflow.Context) (*cloudfront.UpdateFieldLevelEncryptionProfileOutput, error) {
    var output cloudfront.UpdateFieldLevelEncryptionProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontUpdateOriginRequestPolicyResult struct {
	Result workflow.Future
}

func (r *CloudfrontUpdateOriginRequestPolicyResult) Get(ctx workflow.Context) (*cloudfront.UpdateOriginRequestPolicyOutput, error) {
    var output cloudfront.UpdateOriginRequestPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontUpdatePublicKeyResult struct {
	Result workflow.Future
}

func (r *CloudfrontUpdatePublicKeyResult) Get(ctx workflow.Context) (*cloudfront.UpdatePublicKeyOutput, error) {
    var output cloudfront.UpdatePublicKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudfrontUpdateStreamingDistributionResult struct {
	Result workflow.Future
}

func (r *CloudfrontUpdateStreamingDistributionResult) Get(ctx workflow.Context) (*cloudfront.UpdateStreamingDistributionOutput, error) {
    var output cloudfront.UpdateStreamingDistributionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type CloudFrontStub struct {
    activities CloudFrontClient
}

func NewCloudFrontStub() CloudFrontClient {
    return &CloudFrontStub{}
}

func (a *CloudFrontStub) CreateCachePolicy(ctx workflow.Context, input *cloudfront.CreateCachePolicyInput) (*cloudfront.CreateCachePolicyOutput, error) {
    var output cloudfront.CreateCachePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCachePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) CreateCloudFrontOriginAccessIdentity(ctx workflow.Context, input *cloudfront.CreateCloudFrontOriginAccessIdentityInput) (*cloudfront.CreateCloudFrontOriginAccessIdentityOutput, error) {
    var output cloudfront.CreateCloudFrontOriginAccessIdentityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCloudFrontOriginAccessIdentity, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) CreateDistribution(ctx workflow.Context, input *cloudfront.CreateDistributionInput) (*cloudfront.CreateDistributionOutput, error) {
    var output cloudfront.CreateDistributionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDistribution, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) CreateDistributionWithTags(ctx workflow.Context, input *cloudfront.CreateDistributionWithTagsInput) (*cloudfront.CreateDistributionWithTagsOutput, error) {
    var output cloudfront.CreateDistributionWithTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDistributionWithTags, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) CreateFieldLevelEncryptionConfig(ctx workflow.Context, input *cloudfront.CreateFieldLevelEncryptionConfigInput) (*cloudfront.CreateFieldLevelEncryptionConfigOutput, error) {
    var output cloudfront.CreateFieldLevelEncryptionConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFieldLevelEncryptionConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) CreateFieldLevelEncryptionProfile(ctx workflow.Context, input *cloudfront.CreateFieldLevelEncryptionProfileInput) (*cloudfront.CreateFieldLevelEncryptionProfileOutput, error) {
    var output cloudfront.CreateFieldLevelEncryptionProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFieldLevelEncryptionProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) CreateInvalidation(ctx workflow.Context, input *cloudfront.CreateInvalidationInput) (*cloudfront.CreateInvalidationOutput, error) {
    var output cloudfront.CreateInvalidationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateInvalidation, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) CreateMonitoringSubscription(ctx workflow.Context, input *cloudfront.CreateMonitoringSubscriptionInput) (*cloudfront.CreateMonitoringSubscriptionOutput, error) {
    var output cloudfront.CreateMonitoringSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateMonitoringSubscription, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) CreateOriginRequestPolicy(ctx workflow.Context, input *cloudfront.CreateOriginRequestPolicyInput) (*cloudfront.CreateOriginRequestPolicyOutput, error) {
    var output cloudfront.CreateOriginRequestPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateOriginRequestPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) CreatePublicKey(ctx workflow.Context, input *cloudfront.CreatePublicKeyInput) (*cloudfront.CreatePublicKeyOutput, error) {
    var output cloudfront.CreatePublicKeyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePublicKey, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) CreateStreamingDistribution(ctx workflow.Context, input *cloudfront.CreateStreamingDistributionInput) (*cloudfront.CreateStreamingDistributionOutput, error) {
    var output cloudfront.CreateStreamingDistributionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateStreamingDistribution, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) CreateStreamingDistributionWithTags(ctx workflow.Context, input *cloudfront.CreateStreamingDistributionWithTagsInput) (*cloudfront.CreateStreamingDistributionWithTagsOutput, error) {
    var output cloudfront.CreateStreamingDistributionWithTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateStreamingDistributionWithTags, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) DeleteCachePolicy(ctx workflow.Context, input *cloudfront.DeleteCachePolicyInput) (*cloudfront.DeleteCachePolicyOutput, error) {
    var output cloudfront.DeleteCachePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCachePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) DeleteCloudFrontOriginAccessIdentity(ctx workflow.Context, input *cloudfront.DeleteCloudFrontOriginAccessIdentityInput) (*cloudfront.DeleteCloudFrontOriginAccessIdentityOutput, error) {
    var output cloudfront.DeleteCloudFrontOriginAccessIdentityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCloudFrontOriginAccessIdentity, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) DeleteDistribution(ctx workflow.Context, input *cloudfront.DeleteDistributionInput) (*cloudfront.DeleteDistributionOutput, error) {
    var output cloudfront.DeleteDistributionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDistribution, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) DeleteFieldLevelEncryptionConfig(ctx workflow.Context, input *cloudfront.DeleteFieldLevelEncryptionConfigInput) (*cloudfront.DeleteFieldLevelEncryptionConfigOutput, error) {
    var output cloudfront.DeleteFieldLevelEncryptionConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFieldLevelEncryptionConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) DeleteFieldLevelEncryptionProfile(ctx workflow.Context, input *cloudfront.DeleteFieldLevelEncryptionProfileInput) (*cloudfront.DeleteFieldLevelEncryptionProfileOutput, error) {
    var output cloudfront.DeleteFieldLevelEncryptionProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFieldLevelEncryptionProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) DeleteMonitoringSubscription(ctx workflow.Context, input *cloudfront.DeleteMonitoringSubscriptionInput) (*cloudfront.DeleteMonitoringSubscriptionOutput, error) {
    var output cloudfront.DeleteMonitoringSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteMonitoringSubscription, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) DeleteOriginRequestPolicy(ctx workflow.Context, input *cloudfront.DeleteOriginRequestPolicyInput) (*cloudfront.DeleteOriginRequestPolicyOutput, error) {
    var output cloudfront.DeleteOriginRequestPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteOriginRequestPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) DeletePublicKey(ctx workflow.Context, input *cloudfront.DeletePublicKeyInput) (*cloudfront.DeletePublicKeyOutput, error) {
    var output cloudfront.DeletePublicKeyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePublicKey, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) DeleteStreamingDistribution(ctx workflow.Context, input *cloudfront.DeleteStreamingDistributionInput) (*cloudfront.DeleteStreamingDistributionOutput, error) {
    var output cloudfront.DeleteStreamingDistributionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteStreamingDistribution, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) GetCachePolicy(ctx workflow.Context, input *cloudfront.GetCachePolicyInput) (*cloudfront.GetCachePolicyOutput, error) {
    var output cloudfront.GetCachePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCachePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) GetCachePolicyConfig(ctx workflow.Context, input *cloudfront.GetCachePolicyConfigInput) (*cloudfront.GetCachePolicyConfigOutput, error) {
    var output cloudfront.GetCachePolicyConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCachePolicyConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) GetCloudFrontOriginAccessIdentity(ctx workflow.Context, input *cloudfront.GetCloudFrontOriginAccessIdentityInput) (*cloudfront.GetCloudFrontOriginAccessIdentityOutput, error) {
    var output cloudfront.GetCloudFrontOriginAccessIdentityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCloudFrontOriginAccessIdentity, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) GetCloudFrontOriginAccessIdentityConfig(ctx workflow.Context, input *cloudfront.GetCloudFrontOriginAccessIdentityConfigInput) (*cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput, error) {
    var output cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCloudFrontOriginAccessIdentityConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) GetDistribution(ctx workflow.Context, input *cloudfront.GetDistributionInput) (*cloudfront.GetDistributionOutput, error) {
    var output cloudfront.GetDistributionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDistribution, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) GetDistributionConfig(ctx workflow.Context, input *cloudfront.GetDistributionConfigInput) (*cloudfront.GetDistributionConfigOutput, error) {
    var output cloudfront.GetDistributionConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDistributionConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) GetFieldLevelEncryption(ctx workflow.Context, input *cloudfront.GetFieldLevelEncryptionInput) (*cloudfront.GetFieldLevelEncryptionOutput, error) {
    var output cloudfront.GetFieldLevelEncryptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFieldLevelEncryption, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) GetFieldLevelEncryptionConfig(ctx workflow.Context, input *cloudfront.GetFieldLevelEncryptionConfigInput) (*cloudfront.GetFieldLevelEncryptionConfigOutput, error) {
    var output cloudfront.GetFieldLevelEncryptionConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFieldLevelEncryptionConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) GetFieldLevelEncryptionProfile(ctx workflow.Context, input *cloudfront.GetFieldLevelEncryptionProfileInput) (*cloudfront.GetFieldLevelEncryptionProfileOutput, error) {
    var output cloudfront.GetFieldLevelEncryptionProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFieldLevelEncryptionProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) GetFieldLevelEncryptionProfileConfig(ctx workflow.Context, input *cloudfront.GetFieldLevelEncryptionProfileConfigInput) (*cloudfront.GetFieldLevelEncryptionProfileConfigOutput, error) {
    var output cloudfront.GetFieldLevelEncryptionProfileConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFieldLevelEncryptionProfileConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) GetInvalidation(ctx workflow.Context, input *cloudfront.GetInvalidationInput) (*cloudfront.GetInvalidationOutput, error) {
    var output cloudfront.GetInvalidationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetInvalidation, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) GetMonitoringSubscription(ctx workflow.Context, input *cloudfront.GetMonitoringSubscriptionInput) (*cloudfront.GetMonitoringSubscriptionOutput, error) {
    var output cloudfront.GetMonitoringSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMonitoringSubscription, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) GetOriginRequestPolicy(ctx workflow.Context, input *cloudfront.GetOriginRequestPolicyInput) (*cloudfront.GetOriginRequestPolicyOutput, error) {
    var output cloudfront.GetOriginRequestPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetOriginRequestPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) GetOriginRequestPolicyConfig(ctx workflow.Context, input *cloudfront.GetOriginRequestPolicyConfigInput) (*cloudfront.GetOriginRequestPolicyConfigOutput, error) {
    var output cloudfront.GetOriginRequestPolicyConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetOriginRequestPolicyConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) GetPublicKey(ctx workflow.Context, input *cloudfront.GetPublicKeyInput) (*cloudfront.GetPublicKeyOutput, error) {
    var output cloudfront.GetPublicKeyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPublicKey, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) GetPublicKeyConfig(ctx workflow.Context, input *cloudfront.GetPublicKeyConfigInput) (*cloudfront.GetPublicKeyConfigOutput, error) {
    var output cloudfront.GetPublicKeyConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPublicKeyConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) GetStreamingDistribution(ctx workflow.Context, input *cloudfront.GetStreamingDistributionInput) (*cloudfront.GetStreamingDistributionOutput, error) {
    var output cloudfront.GetStreamingDistributionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetStreamingDistribution, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) GetStreamingDistributionConfig(ctx workflow.Context, input *cloudfront.GetStreamingDistributionConfigInput) (*cloudfront.GetStreamingDistributionConfigOutput, error) {
    var output cloudfront.GetStreamingDistributionConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetStreamingDistributionConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) ListCachePolicies(ctx workflow.Context, input *cloudfront.ListCachePoliciesInput) (*cloudfront.ListCachePoliciesOutput, error) {
    var output cloudfront.ListCachePoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListCachePolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) ListCloudFrontOriginAccessIdentities(ctx workflow.Context, input *cloudfront.ListCloudFrontOriginAccessIdentitiesInput) (*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, error) {
    var output cloudfront.ListCloudFrontOriginAccessIdentitiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListCloudFrontOriginAccessIdentities, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) ListDistributions(ctx workflow.Context, input *cloudfront.ListDistributionsInput) (*cloudfront.ListDistributionsOutput, error) {
    var output cloudfront.ListDistributionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDistributions, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) ListDistributionsByCachePolicyId(ctx workflow.Context, input *cloudfront.ListDistributionsByCachePolicyIdInput) (*cloudfront.ListDistributionsByCachePolicyIdOutput, error) {
    var output cloudfront.ListDistributionsByCachePolicyIdOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDistributionsByCachePolicyId, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) ListDistributionsByOriginRequestPolicyId(ctx workflow.Context, input *cloudfront.ListDistributionsByOriginRequestPolicyIdInput) (*cloudfront.ListDistributionsByOriginRequestPolicyIdOutput, error) {
    var output cloudfront.ListDistributionsByOriginRequestPolicyIdOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDistributionsByOriginRequestPolicyId, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) ListDistributionsByWebACLId(ctx workflow.Context, input *cloudfront.ListDistributionsByWebACLIdInput) (*cloudfront.ListDistributionsByWebACLIdOutput, error) {
    var output cloudfront.ListDistributionsByWebACLIdOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDistributionsByWebACLId, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) ListFieldLevelEncryptionConfigs(ctx workflow.Context, input *cloudfront.ListFieldLevelEncryptionConfigsInput) (*cloudfront.ListFieldLevelEncryptionConfigsOutput, error) {
    var output cloudfront.ListFieldLevelEncryptionConfigsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFieldLevelEncryptionConfigs, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) ListFieldLevelEncryptionProfiles(ctx workflow.Context, input *cloudfront.ListFieldLevelEncryptionProfilesInput) (*cloudfront.ListFieldLevelEncryptionProfilesOutput, error) {
    var output cloudfront.ListFieldLevelEncryptionProfilesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFieldLevelEncryptionProfiles, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) ListInvalidations(ctx workflow.Context, input *cloudfront.ListInvalidationsInput) (*cloudfront.ListInvalidationsOutput, error) {
    var output cloudfront.ListInvalidationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListInvalidations, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) ListOriginRequestPolicies(ctx workflow.Context, input *cloudfront.ListOriginRequestPoliciesInput) (*cloudfront.ListOriginRequestPoliciesOutput, error) {
    var output cloudfront.ListOriginRequestPoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListOriginRequestPolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) ListPublicKeys(ctx workflow.Context, input *cloudfront.ListPublicKeysInput) (*cloudfront.ListPublicKeysOutput, error) {
    var output cloudfront.ListPublicKeysOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPublicKeys, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) ListStreamingDistributions(ctx workflow.Context, input *cloudfront.ListStreamingDistributionsInput) (*cloudfront.ListStreamingDistributionsOutput, error) {
    var output cloudfront.ListStreamingDistributionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListStreamingDistributions, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) ListTagsForResource(ctx workflow.Context, input *cloudfront.ListTagsForResourceInput) (*cloudfront.ListTagsForResourceOutput, error) {
    var output cloudfront.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) TagResource(ctx workflow.Context, input *cloudfront.TagResourceInput) (*cloudfront.TagResourceOutput, error) {
    var output cloudfront.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) UntagResource(ctx workflow.Context, input *cloudfront.UntagResourceInput) (*cloudfront.UntagResourceOutput, error) {
    var output cloudfront.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) UpdateCachePolicy(ctx workflow.Context, input *cloudfront.UpdateCachePolicyInput) (*cloudfront.UpdateCachePolicyOutput, error) {
    var output cloudfront.UpdateCachePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateCachePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) UpdateCloudFrontOriginAccessIdentity(ctx workflow.Context, input *cloudfront.UpdateCloudFrontOriginAccessIdentityInput) (*cloudfront.UpdateCloudFrontOriginAccessIdentityOutput, error) {
    var output cloudfront.UpdateCloudFrontOriginAccessIdentityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateCloudFrontOriginAccessIdentity, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) UpdateDistribution(ctx workflow.Context, input *cloudfront.UpdateDistributionInput) (*cloudfront.UpdateDistributionOutput, error) {
    var output cloudfront.UpdateDistributionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDistribution, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) UpdateFieldLevelEncryptionConfig(ctx workflow.Context, input *cloudfront.UpdateFieldLevelEncryptionConfigInput) (*cloudfront.UpdateFieldLevelEncryptionConfigOutput, error) {
    var output cloudfront.UpdateFieldLevelEncryptionConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFieldLevelEncryptionConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) UpdateFieldLevelEncryptionProfile(ctx workflow.Context, input *cloudfront.UpdateFieldLevelEncryptionProfileInput) (*cloudfront.UpdateFieldLevelEncryptionProfileOutput, error) {
    var output cloudfront.UpdateFieldLevelEncryptionProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFieldLevelEncryptionProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) UpdateOriginRequestPolicy(ctx workflow.Context, input *cloudfront.UpdateOriginRequestPolicyInput) (*cloudfront.UpdateOriginRequestPolicyOutput, error) {
    var output cloudfront.UpdateOriginRequestPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateOriginRequestPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) UpdatePublicKey(ctx workflow.Context, input *cloudfront.UpdatePublicKeyInput) (*cloudfront.UpdatePublicKeyOutput, error) {
    var output cloudfront.UpdatePublicKeyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdatePublicKey, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFrontStub) UpdateStreamingDistribution(ctx workflow.Context, input *cloudfront.UpdateStreamingDistributionInput) (*cloudfront.UpdateStreamingDistributionOutput, error) {
    var output cloudfront.UpdateStreamingDistributionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateStreamingDistribution, input).Get(ctx, &output)
    return &output, err
}


func (a *CloudFrontStub) WaitUntilDistributionDeployed(ctx workflow.Context, input *cloudfront.GetDistributionInput) error {
    return a.client.WaitUntilDistributionDeployed(input)
}


func (a *CloudFrontStub) WaitUntilInvalidationCompleted(ctx workflow.Context, input *cloudfront.GetInvalidationInput) error {
    return a.client.WaitUntilInvalidationCompleted(input)
}


func (a *CloudFrontStub) WaitUntilStreamingDistributionDeployed(ctx workflow.Context, input *cloudfront.GetStreamingDistributionInput) error {
    return a.client.WaitUntilStreamingDistributionDeployed(input)
}
