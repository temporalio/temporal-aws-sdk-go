package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/cloudfront/cloudfrontiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type CloudFrontActivities struct {
	client cloudfrontiface.CloudFrontAPI
}

func NewCloudFrontActivities(session *session.Session, config ...*aws.Config) *CloudFrontActivities {
	client := cloudfront.New(session, config...)
	return &CloudFrontActivities{client: client}
}

func (a *CloudFrontActivities) CreateCachePolicy(ctx context.Context, input *cloudfront.CreateCachePolicyInput) (*cloudfront.CreateCachePolicyOutput, error) {
	return a.client.CreateCachePolicyWithContext(ctx, input)
}

func (a *CloudFrontActivities) CreateCloudFrontOriginAccessIdentity(ctx context.Context, input *cloudfront.CreateCloudFrontOriginAccessIdentityInput) (*cloudfront.CreateCloudFrontOriginAccessIdentityOutput, error) {
	return a.client.CreateCloudFrontOriginAccessIdentityWithContext(ctx, input)
}

func (a *CloudFrontActivities) CreateDistribution(ctx context.Context, input *cloudfront.CreateDistributionInput) (*cloudfront.CreateDistributionOutput, error) {
	return a.client.CreateDistributionWithContext(ctx, input)
}

func (a *CloudFrontActivities) CreateDistributionWithTags(ctx context.Context, input *cloudfront.CreateDistributionWithTagsInput) (*cloudfront.CreateDistributionWithTagsOutput, error) {
	return a.client.CreateDistributionWithTagsWithContext(ctx, input)
}

func (a *CloudFrontActivities) CreateFieldLevelEncryptionConfig(ctx context.Context, input *cloudfront.CreateFieldLevelEncryptionConfigInput) (*cloudfront.CreateFieldLevelEncryptionConfigOutput, error) {
	return a.client.CreateFieldLevelEncryptionConfigWithContext(ctx, input)
}

func (a *CloudFrontActivities) CreateFieldLevelEncryptionProfile(ctx context.Context, input *cloudfront.CreateFieldLevelEncryptionProfileInput) (*cloudfront.CreateFieldLevelEncryptionProfileOutput, error) {
	return a.client.CreateFieldLevelEncryptionProfileWithContext(ctx, input)
}

func (a *CloudFrontActivities) CreateInvalidation(ctx context.Context, input *cloudfront.CreateInvalidationInput) (*cloudfront.CreateInvalidationOutput, error) {
	return a.client.CreateInvalidationWithContext(ctx, input)
}

func (a *CloudFrontActivities) CreateMonitoringSubscription(ctx context.Context, input *cloudfront.CreateMonitoringSubscriptionInput) (*cloudfront.CreateMonitoringSubscriptionOutput, error) {
	return a.client.CreateMonitoringSubscriptionWithContext(ctx, input)
}

func (a *CloudFrontActivities) CreateOriginRequestPolicy(ctx context.Context, input *cloudfront.CreateOriginRequestPolicyInput) (*cloudfront.CreateOriginRequestPolicyOutput, error) {
	return a.client.CreateOriginRequestPolicyWithContext(ctx, input)
}

func (a *CloudFrontActivities) CreatePublicKey(ctx context.Context, input *cloudfront.CreatePublicKeyInput) (*cloudfront.CreatePublicKeyOutput, error) {
	return a.client.CreatePublicKeyWithContext(ctx, input)
}

func (a *CloudFrontActivities) CreateRealtimeLogConfig(ctx context.Context, input *cloudfront.CreateRealtimeLogConfigInput) (*cloudfront.CreateRealtimeLogConfigOutput, error) {
	return a.client.CreateRealtimeLogConfigWithContext(ctx, input)
}

func (a *CloudFrontActivities) CreateStreamingDistribution(ctx context.Context, input *cloudfront.CreateStreamingDistributionInput) (*cloudfront.CreateStreamingDistributionOutput, error) {
	return a.client.CreateStreamingDistributionWithContext(ctx, input)
}

func (a *CloudFrontActivities) CreateStreamingDistributionWithTags(ctx context.Context, input *cloudfront.CreateStreamingDistributionWithTagsInput) (*cloudfront.CreateStreamingDistributionWithTagsOutput, error) {
	return a.client.CreateStreamingDistributionWithTagsWithContext(ctx, input)
}

func (a *CloudFrontActivities) DeleteCachePolicy(ctx context.Context, input *cloudfront.DeleteCachePolicyInput) (*cloudfront.DeleteCachePolicyOutput, error) {
	return a.client.DeleteCachePolicyWithContext(ctx, input)
}

func (a *CloudFrontActivities) DeleteCloudFrontOriginAccessIdentity(ctx context.Context, input *cloudfront.DeleteCloudFrontOriginAccessIdentityInput) (*cloudfront.DeleteCloudFrontOriginAccessIdentityOutput, error) {
	return a.client.DeleteCloudFrontOriginAccessIdentityWithContext(ctx, input)
}

func (a *CloudFrontActivities) DeleteDistribution(ctx context.Context, input *cloudfront.DeleteDistributionInput) (*cloudfront.DeleteDistributionOutput, error) {
	return a.client.DeleteDistributionWithContext(ctx, input)
}

func (a *CloudFrontActivities) DeleteFieldLevelEncryptionConfig(ctx context.Context, input *cloudfront.DeleteFieldLevelEncryptionConfigInput) (*cloudfront.DeleteFieldLevelEncryptionConfigOutput, error) {
	return a.client.DeleteFieldLevelEncryptionConfigWithContext(ctx, input)
}

func (a *CloudFrontActivities) DeleteFieldLevelEncryptionProfile(ctx context.Context, input *cloudfront.DeleteFieldLevelEncryptionProfileInput) (*cloudfront.DeleteFieldLevelEncryptionProfileOutput, error) {
	return a.client.DeleteFieldLevelEncryptionProfileWithContext(ctx, input)
}

func (a *CloudFrontActivities) DeleteMonitoringSubscription(ctx context.Context, input *cloudfront.DeleteMonitoringSubscriptionInput) (*cloudfront.DeleteMonitoringSubscriptionOutput, error) {
	return a.client.DeleteMonitoringSubscriptionWithContext(ctx, input)
}

func (a *CloudFrontActivities) DeleteOriginRequestPolicy(ctx context.Context, input *cloudfront.DeleteOriginRequestPolicyInput) (*cloudfront.DeleteOriginRequestPolicyOutput, error) {
	return a.client.DeleteOriginRequestPolicyWithContext(ctx, input)
}

func (a *CloudFrontActivities) DeletePublicKey(ctx context.Context, input *cloudfront.DeletePublicKeyInput) (*cloudfront.DeletePublicKeyOutput, error) {
	return a.client.DeletePublicKeyWithContext(ctx, input)
}

func (a *CloudFrontActivities) DeleteRealtimeLogConfig(ctx context.Context, input *cloudfront.DeleteRealtimeLogConfigInput) (*cloudfront.DeleteRealtimeLogConfigOutput, error) {
	return a.client.DeleteRealtimeLogConfigWithContext(ctx, input)
}

func (a *CloudFrontActivities) DeleteStreamingDistribution(ctx context.Context, input *cloudfront.DeleteStreamingDistributionInput) (*cloudfront.DeleteStreamingDistributionOutput, error) {
	return a.client.DeleteStreamingDistributionWithContext(ctx, input)
}

func (a *CloudFrontActivities) GetCachePolicy(ctx context.Context, input *cloudfront.GetCachePolicyInput) (*cloudfront.GetCachePolicyOutput, error) {
	return a.client.GetCachePolicyWithContext(ctx, input)
}

func (a *CloudFrontActivities) GetCachePolicyConfig(ctx context.Context, input *cloudfront.GetCachePolicyConfigInput) (*cloudfront.GetCachePolicyConfigOutput, error) {
	return a.client.GetCachePolicyConfigWithContext(ctx, input)
}

func (a *CloudFrontActivities) GetCloudFrontOriginAccessIdentity(ctx context.Context, input *cloudfront.GetCloudFrontOriginAccessIdentityInput) (*cloudfront.GetCloudFrontOriginAccessIdentityOutput, error) {
	return a.client.GetCloudFrontOriginAccessIdentityWithContext(ctx, input)
}

func (a *CloudFrontActivities) GetCloudFrontOriginAccessIdentityConfig(ctx context.Context, input *cloudfront.GetCloudFrontOriginAccessIdentityConfigInput) (*cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput, error) {
	return a.client.GetCloudFrontOriginAccessIdentityConfigWithContext(ctx, input)
}

func (a *CloudFrontActivities) GetDistribution(ctx context.Context, input *cloudfront.GetDistributionInput) (*cloudfront.GetDistributionOutput, error) {
	return a.client.GetDistributionWithContext(ctx, input)
}

func (a *CloudFrontActivities) GetDistributionConfig(ctx context.Context, input *cloudfront.GetDistributionConfigInput) (*cloudfront.GetDistributionConfigOutput, error) {
	return a.client.GetDistributionConfigWithContext(ctx, input)
}

func (a *CloudFrontActivities) GetFieldLevelEncryption(ctx context.Context, input *cloudfront.GetFieldLevelEncryptionInput) (*cloudfront.GetFieldLevelEncryptionOutput, error) {
	return a.client.GetFieldLevelEncryptionWithContext(ctx, input)
}

func (a *CloudFrontActivities) GetFieldLevelEncryptionConfig(ctx context.Context, input *cloudfront.GetFieldLevelEncryptionConfigInput) (*cloudfront.GetFieldLevelEncryptionConfigOutput, error) {
	return a.client.GetFieldLevelEncryptionConfigWithContext(ctx, input)
}

func (a *CloudFrontActivities) GetFieldLevelEncryptionProfile(ctx context.Context, input *cloudfront.GetFieldLevelEncryptionProfileInput) (*cloudfront.GetFieldLevelEncryptionProfileOutput, error) {
	return a.client.GetFieldLevelEncryptionProfileWithContext(ctx, input)
}

func (a *CloudFrontActivities) GetFieldLevelEncryptionProfileConfig(ctx context.Context, input *cloudfront.GetFieldLevelEncryptionProfileConfigInput) (*cloudfront.GetFieldLevelEncryptionProfileConfigOutput, error) {
	return a.client.GetFieldLevelEncryptionProfileConfigWithContext(ctx, input)
}

func (a *CloudFrontActivities) GetInvalidation(ctx context.Context, input *cloudfront.GetInvalidationInput) (*cloudfront.GetInvalidationOutput, error) {
	return a.client.GetInvalidationWithContext(ctx, input)
}

func (a *CloudFrontActivities) GetMonitoringSubscription(ctx context.Context, input *cloudfront.GetMonitoringSubscriptionInput) (*cloudfront.GetMonitoringSubscriptionOutput, error) {
	return a.client.GetMonitoringSubscriptionWithContext(ctx, input)
}

func (a *CloudFrontActivities) GetOriginRequestPolicy(ctx context.Context, input *cloudfront.GetOriginRequestPolicyInput) (*cloudfront.GetOriginRequestPolicyOutput, error) {
	return a.client.GetOriginRequestPolicyWithContext(ctx, input)
}

func (a *CloudFrontActivities) GetOriginRequestPolicyConfig(ctx context.Context, input *cloudfront.GetOriginRequestPolicyConfigInput) (*cloudfront.GetOriginRequestPolicyConfigOutput, error) {
	return a.client.GetOriginRequestPolicyConfigWithContext(ctx, input)
}

func (a *CloudFrontActivities) GetPublicKey(ctx context.Context, input *cloudfront.GetPublicKeyInput) (*cloudfront.GetPublicKeyOutput, error) {
	return a.client.GetPublicKeyWithContext(ctx, input)
}

func (a *CloudFrontActivities) GetPublicKeyConfig(ctx context.Context, input *cloudfront.GetPublicKeyConfigInput) (*cloudfront.GetPublicKeyConfigOutput, error) {
	return a.client.GetPublicKeyConfigWithContext(ctx, input)
}

func (a *CloudFrontActivities) GetRealtimeLogConfig(ctx context.Context, input *cloudfront.GetRealtimeLogConfigInput) (*cloudfront.GetRealtimeLogConfigOutput, error) {
	return a.client.GetRealtimeLogConfigWithContext(ctx, input)
}

func (a *CloudFrontActivities) GetStreamingDistribution(ctx context.Context, input *cloudfront.GetStreamingDistributionInput) (*cloudfront.GetStreamingDistributionOutput, error) {
	return a.client.GetStreamingDistributionWithContext(ctx, input)
}

func (a *CloudFrontActivities) GetStreamingDistributionConfig(ctx context.Context, input *cloudfront.GetStreamingDistributionConfigInput) (*cloudfront.GetStreamingDistributionConfigOutput, error) {
	return a.client.GetStreamingDistributionConfigWithContext(ctx, input)
}

func (a *CloudFrontActivities) ListCachePolicies(ctx context.Context, input *cloudfront.ListCachePoliciesInput) (*cloudfront.ListCachePoliciesOutput, error) {
	return a.client.ListCachePoliciesWithContext(ctx, input)
}

func (a *CloudFrontActivities) ListCloudFrontOriginAccessIdentities(ctx context.Context, input *cloudfront.ListCloudFrontOriginAccessIdentitiesInput) (*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, error) {
	return a.client.ListCloudFrontOriginAccessIdentitiesWithContext(ctx, input)
}

func (a *CloudFrontActivities) ListDistributions(ctx context.Context, input *cloudfront.ListDistributionsInput) (*cloudfront.ListDistributionsOutput, error) {
	return a.client.ListDistributionsWithContext(ctx, input)
}

func (a *CloudFrontActivities) ListDistributionsByCachePolicyId(ctx context.Context, input *cloudfront.ListDistributionsByCachePolicyIdInput) (*cloudfront.ListDistributionsByCachePolicyIdOutput, error) {
	return a.client.ListDistributionsByCachePolicyIdWithContext(ctx, input)
}

func (a *CloudFrontActivities) ListDistributionsByOriginRequestPolicyId(ctx context.Context, input *cloudfront.ListDistributionsByOriginRequestPolicyIdInput) (*cloudfront.ListDistributionsByOriginRequestPolicyIdOutput, error) {
	return a.client.ListDistributionsByOriginRequestPolicyIdWithContext(ctx, input)
}

func (a *CloudFrontActivities) ListDistributionsByRealtimeLogConfig(ctx context.Context, input *cloudfront.ListDistributionsByRealtimeLogConfigInput) (*cloudfront.ListDistributionsByRealtimeLogConfigOutput, error) {
	return a.client.ListDistributionsByRealtimeLogConfigWithContext(ctx, input)
}

func (a *CloudFrontActivities) ListDistributionsByWebACLId(ctx context.Context, input *cloudfront.ListDistributionsByWebACLIdInput) (*cloudfront.ListDistributionsByWebACLIdOutput, error) {
	return a.client.ListDistributionsByWebACLIdWithContext(ctx, input)
}

func (a *CloudFrontActivities) ListFieldLevelEncryptionConfigs(ctx context.Context, input *cloudfront.ListFieldLevelEncryptionConfigsInput) (*cloudfront.ListFieldLevelEncryptionConfigsOutput, error) {
	return a.client.ListFieldLevelEncryptionConfigsWithContext(ctx, input)
}

func (a *CloudFrontActivities) ListFieldLevelEncryptionProfiles(ctx context.Context, input *cloudfront.ListFieldLevelEncryptionProfilesInput) (*cloudfront.ListFieldLevelEncryptionProfilesOutput, error) {
	return a.client.ListFieldLevelEncryptionProfilesWithContext(ctx, input)
}

func (a *CloudFrontActivities) ListInvalidations(ctx context.Context, input *cloudfront.ListInvalidationsInput) (*cloudfront.ListInvalidationsOutput, error) {
	return a.client.ListInvalidationsWithContext(ctx, input)
}

func (a *CloudFrontActivities) ListOriginRequestPolicies(ctx context.Context, input *cloudfront.ListOriginRequestPoliciesInput) (*cloudfront.ListOriginRequestPoliciesOutput, error) {
	return a.client.ListOriginRequestPoliciesWithContext(ctx, input)
}

func (a *CloudFrontActivities) ListPublicKeys(ctx context.Context, input *cloudfront.ListPublicKeysInput) (*cloudfront.ListPublicKeysOutput, error) {
	return a.client.ListPublicKeysWithContext(ctx, input)
}

func (a *CloudFrontActivities) ListRealtimeLogConfigs(ctx context.Context, input *cloudfront.ListRealtimeLogConfigsInput) (*cloudfront.ListRealtimeLogConfigsOutput, error) {
	return a.client.ListRealtimeLogConfigsWithContext(ctx, input)
}

func (a *CloudFrontActivities) ListStreamingDistributions(ctx context.Context, input *cloudfront.ListStreamingDistributionsInput) (*cloudfront.ListStreamingDistributionsOutput, error) {
	return a.client.ListStreamingDistributionsWithContext(ctx, input)
}

func (a *CloudFrontActivities) ListTagsForResource(ctx context.Context, input *cloudfront.ListTagsForResourceInput) (*cloudfront.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *CloudFrontActivities) TagResource(ctx context.Context, input *cloudfront.TagResourceInput) (*cloudfront.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *CloudFrontActivities) UntagResource(ctx context.Context, input *cloudfront.UntagResourceInput) (*cloudfront.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *CloudFrontActivities) UpdateCachePolicy(ctx context.Context, input *cloudfront.UpdateCachePolicyInput) (*cloudfront.UpdateCachePolicyOutput, error) {
	return a.client.UpdateCachePolicyWithContext(ctx, input)
}

func (a *CloudFrontActivities) UpdateCloudFrontOriginAccessIdentity(ctx context.Context, input *cloudfront.UpdateCloudFrontOriginAccessIdentityInput) (*cloudfront.UpdateCloudFrontOriginAccessIdentityOutput, error) {
	return a.client.UpdateCloudFrontOriginAccessIdentityWithContext(ctx, input)
}

func (a *CloudFrontActivities) UpdateDistribution(ctx context.Context, input *cloudfront.UpdateDistributionInput) (*cloudfront.UpdateDistributionOutput, error) {
	return a.client.UpdateDistributionWithContext(ctx, input)
}

func (a *CloudFrontActivities) UpdateFieldLevelEncryptionConfig(ctx context.Context, input *cloudfront.UpdateFieldLevelEncryptionConfigInput) (*cloudfront.UpdateFieldLevelEncryptionConfigOutput, error) {
	return a.client.UpdateFieldLevelEncryptionConfigWithContext(ctx, input)
}

func (a *CloudFrontActivities) UpdateFieldLevelEncryptionProfile(ctx context.Context, input *cloudfront.UpdateFieldLevelEncryptionProfileInput) (*cloudfront.UpdateFieldLevelEncryptionProfileOutput, error) {
	return a.client.UpdateFieldLevelEncryptionProfileWithContext(ctx, input)
}

func (a *CloudFrontActivities) UpdateOriginRequestPolicy(ctx context.Context, input *cloudfront.UpdateOriginRequestPolicyInput) (*cloudfront.UpdateOriginRequestPolicyOutput, error) {
	return a.client.UpdateOriginRequestPolicyWithContext(ctx, input)
}

func (a *CloudFrontActivities) UpdatePublicKey(ctx context.Context, input *cloudfront.UpdatePublicKeyInput) (*cloudfront.UpdatePublicKeyOutput, error) {
	return a.client.UpdatePublicKeyWithContext(ctx, input)
}

func (a *CloudFrontActivities) UpdateRealtimeLogConfig(ctx context.Context, input *cloudfront.UpdateRealtimeLogConfigInput) (*cloudfront.UpdateRealtimeLogConfigOutput, error) {
	return a.client.UpdateRealtimeLogConfigWithContext(ctx, input)
}

func (a *CloudFrontActivities) UpdateStreamingDistribution(ctx context.Context, input *cloudfront.UpdateStreamingDistributionInput) (*cloudfront.UpdateStreamingDistributionOutput, error) {
	return a.client.UpdateStreamingDistributionWithContext(ctx, input)
}

func (a *CloudFrontActivities) WaitUntilDistributionDeployed(ctx context.Context, input *cloudfront.GetDistributionInput) error {
	return a.client.WaitUntilDistributionDeployedWithContext(ctx, input)

}

func (a *CloudFrontActivities) WaitUntilInvalidationCompleted(ctx context.Context, input *cloudfront.GetInvalidationInput) error {
	return a.client.WaitUntilInvalidationCompletedWithContext(ctx, input)

}

func (a *CloudFrontActivities) WaitUntilStreamingDistributionDeployed(ctx context.Context, input *cloudfront.GetStreamingDistributionInput) error {
	return a.client.WaitUntilStreamingDistributionDeployedWithContext(ctx, input)

}
