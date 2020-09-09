
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/cloudfront/cloudfrontiface"
)

type CloudFrontActivities struct {
    client cloudfrontiface.CloudFrontAPI
}

func NewCloudFrontActivities(session *session.Session, config ...*aws.Config) *CloudFrontActivities {
    client := cloudfront.New(session, config...)
    return &CloudFrontActivities{client: client}
}

func (a *CloudFrontActivities) CreateCachePolicy(input *cloudfront.CreateCachePolicyInput) (*cloudfront.CreateCachePolicyOutput, error) {
    return a.client.CreateCachePolicy(input)
}

func (a *CloudFrontActivities) CreateCloudFrontOriginAccessIdentity(input *cloudfront.CreateCloudFrontOriginAccessIdentityInput) (*cloudfront.CreateCloudFrontOriginAccessIdentityOutput, error) {
    return a.client.CreateCloudFrontOriginAccessIdentity(input)
}

func (a *CloudFrontActivities) CreateDistribution(input *cloudfront.CreateDistributionInput) (*cloudfront.CreateDistributionOutput, error) {
    return a.client.CreateDistribution(input)
}

func (a *CloudFrontActivities) CreateDistributionWithTags(input *cloudfront.CreateDistributionWithTagsInput) (*cloudfront.CreateDistributionWithTagsOutput, error) {
    return a.client.CreateDistributionWithTags(input)
}

func (a *CloudFrontActivities) CreateFieldLevelEncryptionConfig(input *cloudfront.CreateFieldLevelEncryptionConfigInput) (*cloudfront.CreateFieldLevelEncryptionConfigOutput, error) {
    return a.client.CreateFieldLevelEncryptionConfig(input)
}

func (a *CloudFrontActivities) CreateFieldLevelEncryptionProfile(input *cloudfront.CreateFieldLevelEncryptionProfileInput) (*cloudfront.CreateFieldLevelEncryptionProfileOutput, error) {
    return a.client.CreateFieldLevelEncryptionProfile(input)
}

func (a *CloudFrontActivities) CreateInvalidation(input *cloudfront.CreateInvalidationInput) (*cloudfront.CreateInvalidationOutput, error) {
    return a.client.CreateInvalidation(input)
}

func (a *CloudFrontActivities) CreateMonitoringSubscription(input *cloudfront.CreateMonitoringSubscriptionInput) (*cloudfront.CreateMonitoringSubscriptionOutput, error) {
    return a.client.CreateMonitoringSubscription(input)
}

func (a *CloudFrontActivities) CreateOriginRequestPolicy(input *cloudfront.CreateOriginRequestPolicyInput) (*cloudfront.CreateOriginRequestPolicyOutput, error) {
    return a.client.CreateOriginRequestPolicy(input)
}

func (a *CloudFrontActivities) CreatePublicKey(input *cloudfront.CreatePublicKeyInput) (*cloudfront.CreatePublicKeyOutput, error) {
    return a.client.CreatePublicKey(input)
}

func (a *CloudFrontActivities) CreateRealtimeLogConfig(input *cloudfront.CreateRealtimeLogConfigInput) (*cloudfront.CreateRealtimeLogConfigOutput, error) {
    return a.client.CreateRealtimeLogConfig(input)
}

func (a *CloudFrontActivities) CreateStreamingDistribution(input *cloudfront.CreateStreamingDistributionInput) (*cloudfront.CreateStreamingDistributionOutput, error) {
    return a.client.CreateStreamingDistribution(input)
}

func (a *CloudFrontActivities) CreateStreamingDistributionWithTags(input *cloudfront.CreateStreamingDistributionWithTagsInput) (*cloudfront.CreateStreamingDistributionWithTagsOutput, error) {
    return a.client.CreateStreamingDistributionWithTags(input)
}

func (a *CloudFrontActivities) DeleteCachePolicy(input *cloudfront.DeleteCachePolicyInput) (*cloudfront.DeleteCachePolicyOutput, error) {
    return a.client.DeleteCachePolicy(input)
}

func (a *CloudFrontActivities) DeleteCloudFrontOriginAccessIdentity(input *cloudfront.DeleteCloudFrontOriginAccessIdentityInput) (*cloudfront.DeleteCloudFrontOriginAccessIdentityOutput, error) {
    return a.client.DeleteCloudFrontOriginAccessIdentity(input)
}

func (a *CloudFrontActivities) DeleteDistribution(input *cloudfront.DeleteDistributionInput) (*cloudfront.DeleteDistributionOutput, error) {
    return a.client.DeleteDistribution(input)
}

func (a *CloudFrontActivities) DeleteFieldLevelEncryptionConfig(input *cloudfront.DeleteFieldLevelEncryptionConfigInput) (*cloudfront.DeleteFieldLevelEncryptionConfigOutput, error) {
    return a.client.DeleteFieldLevelEncryptionConfig(input)
}

func (a *CloudFrontActivities) DeleteFieldLevelEncryptionProfile(input *cloudfront.DeleteFieldLevelEncryptionProfileInput) (*cloudfront.DeleteFieldLevelEncryptionProfileOutput, error) {
    return a.client.DeleteFieldLevelEncryptionProfile(input)
}

func (a *CloudFrontActivities) DeleteMonitoringSubscription(input *cloudfront.DeleteMonitoringSubscriptionInput) (*cloudfront.DeleteMonitoringSubscriptionOutput, error) {
    return a.client.DeleteMonitoringSubscription(input)
}

func (a *CloudFrontActivities) DeleteOriginRequestPolicy(input *cloudfront.DeleteOriginRequestPolicyInput) (*cloudfront.DeleteOriginRequestPolicyOutput, error) {
    return a.client.DeleteOriginRequestPolicy(input)
}

func (a *CloudFrontActivities) DeletePublicKey(input *cloudfront.DeletePublicKeyInput) (*cloudfront.DeletePublicKeyOutput, error) {
    return a.client.DeletePublicKey(input)
}

func (a *CloudFrontActivities) DeleteRealtimeLogConfig(input *cloudfront.DeleteRealtimeLogConfigInput) (*cloudfront.DeleteRealtimeLogConfigOutput, error) {
    return a.client.DeleteRealtimeLogConfig(input)
}

func (a *CloudFrontActivities) DeleteStreamingDistribution(input *cloudfront.DeleteStreamingDistributionInput) (*cloudfront.DeleteStreamingDistributionOutput, error) {
    return a.client.DeleteStreamingDistribution(input)
}

func (a *CloudFrontActivities) GetCachePolicy(input *cloudfront.GetCachePolicyInput) (*cloudfront.GetCachePolicyOutput, error) {
    return a.client.GetCachePolicy(input)
}

func (a *CloudFrontActivities) GetCachePolicyConfig(input *cloudfront.GetCachePolicyConfigInput) (*cloudfront.GetCachePolicyConfigOutput, error) {
    return a.client.GetCachePolicyConfig(input)
}

func (a *CloudFrontActivities) GetCloudFrontOriginAccessIdentity(input *cloudfront.GetCloudFrontOriginAccessIdentityInput) (*cloudfront.GetCloudFrontOriginAccessIdentityOutput, error) {
    return a.client.GetCloudFrontOriginAccessIdentity(input)
}

func (a *CloudFrontActivities) GetCloudFrontOriginAccessIdentityConfig(input *cloudfront.GetCloudFrontOriginAccessIdentityConfigInput) (*cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput, error) {
    return a.client.GetCloudFrontOriginAccessIdentityConfig(input)
}

func (a *CloudFrontActivities) GetDistribution(input *cloudfront.GetDistributionInput) (*cloudfront.GetDistributionOutput, error) {
    return a.client.GetDistribution(input)
}

func (a *CloudFrontActivities) GetDistributionConfig(input *cloudfront.GetDistributionConfigInput) (*cloudfront.GetDistributionConfigOutput, error) {
    return a.client.GetDistributionConfig(input)
}

func (a *CloudFrontActivities) GetFieldLevelEncryption(input *cloudfront.GetFieldLevelEncryptionInput) (*cloudfront.GetFieldLevelEncryptionOutput, error) {
    return a.client.GetFieldLevelEncryption(input)
}

func (a *CloudFrontActivities) GetFieldLevelEncryptionConfig(input *cloudfront.GetFieldLevelEncryptionConfigInput) (*cloudfront.GetFieldLevelEncryptionConfigOutput, error) {
    return a.client.GetFieldLevelEncryptionConfig(input)
}

func (a *CloudFrontActivities) GetFieldLevelEncryptionProfile(input *cloudfront.GetFieldLevelEncryptionProfileInput) (*cloudfront.GetFieldLevelEncryptionProfileOutput, error) {
    return a.client.GetFieldLevelEncryptionProfile(input)
}

func (a *CloudFrontActivities) GetFieldLevelEncryptionProfileConfig(input *cloudfront.GetFieldLevelEncryptionProfileConfigInput) (*cloudfront.GetFieldLevelEncryptionProfileConfigOutput, error) {
    return a.client.GetFieldLevelEncryptionProfileConfig(input)
}

func (a *CloudFrontActivities) GetInvalidation(input *cloudfront.GetInvalidationInput) (*cloudfront.GetInvalidationOutput, error) {
    return a.client.GetInvalidation(input)
}

func (a *CloudFrontActivities) GetMonitoringSubscription(input *cloudfront.GetMonitoringSubscriptionInput) (*cloudfront.GetMonitoringSubscriptionOutput, error) {
    return a.client.GetMonitoringSubscription(input)
}

func (a *CloudFrontActivities) GetOriginRequestPolicy(input *cloudfront.GetOriginRequestPolicyInput) (*cloudfront.GetOriginRequestPolicyOutput, error) {
    return a.client.GetOriginRequestPolicy(input)
}

func (a *CloudFrontActivities) GetOriginRequestPolicyConfig(input *cloudfront.GetOriginRequestPolicyConfigInput) (*cloudfront.GetOriginRequestPolicyConfigOutput, error) {
    return a.client.GetOriginRequestPolicyConfig(input)
}

func (a *CloudFrontActivities) GetPublicKey(input *cloudfront.GetPublicKeyInput) (*cloudfront.GetPublicKeyOutput, error) {
    return a.client.GetPublicKey(input)
}

func (a *CloudFrontActivities) GetPublicKeyConfig(input *cloudfront.GetPublicKeyConfigInput) (*cloudfront.GetPublicKeyConfigOutput, error) {
    return a.client.GetPublicKeyConfig(input)
}

func (a *CloudFrontActivities) GetRealtimeLogConfig(input *cloudfront.GetRealtimeLogConfigInput) (*cloudfront.GetRealtimeLogConfigOutput, error) {
    return a.client.GetRealtimeLogConfig(input)
}

func (a *CloudFrontActivities) GetStreamingDistribution(input *cloudfront.GetStreamingDistributionInput) (*cloudfront.GetStreamingDistributionOutput, error) {
    return a.client.GetStreamingDistribution(input)
}

func (a *CloudFrontActivities) GetStreamingDistributionConfig(input *cloudfront.GetStreamingDistributionConfigInput) (*cloudfront.GetStreamingDistributionConfigOutput, error) {
    return a.client.GetStreamingDistributionConfig(input)
}

func (a *CloudFrontActivities) ListCachePolicies(input *cloudfront.ListCachePoliciesInput) (*cloudfront.ListCachePoliciesOutput, error) {
    return a.client.ListCachePolicies(input)
}

func (a *CloudFrontActivities) ListCloudFrontOriginAccessIdentities(input *cloudfront.ListCloudFrontOriginAccessIdentitiesInput) (*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, error) {
    return a.client.ListCloudFrontOriginAccessIdentities(input)
}

func (a *CloudFrontActivities) ListDistributions(input *cloudfront.ListDistributionsInput) (*cloudfront.ListDistributionsOutput, error) {
    return a.client.ListDistributions(input)
}

func (a *CloudFrontActivities) ListDistributionsByCachePolicyId(input *cloudfront.ListDistributionsByCachePolicyIdInput) (*cloudfront.ListDistributionsByCachePolicyIdOutput, error) {
    return a.client.ListDistributionsByCachePolicyId(input)
}

func (a *CloudFrontActivities) ListDistributionsByOriginRequestPolicyId(input *cloudfront.ListDistributionsByOriginRequestPolicyIdInput) (*cloudfront.ListDistributionsByOriginRequestPolicyIdOutput, error) {
    return a.client.ListDistributionsByOriginRequestPolicyId(input)
}

func (a *CloudFrontActivities) ListDistributionsByRealtimeLogConfig(input *cloudfront.ListDistributionsByRealtimeLogConfigInput) (*cloudfront.ListDistributionsByRealtimeLogConfigOutput, error) {
    return a.client.ListDistributionsByRealtimeLogConfig(input)
}

func (a *CloudFrontActivities) ListDistributionsByWebACLId(input *cloudfront.ListDistributionsByWebACLIdInput) (*cloudfront.ListDistributionsByWebACLIdOutput, error) {
    return a.client.ListDistributionsByWebACLId(input)
}

func (a *CloudFrontActivities) ListFieldLevelEncryptionConfigs(input *cloudfront.ListFieldLevelEncryptionConfigsInput) (*cloudfront.ListFieldLevelEncryptionConfigsOutput, error) {
    return a.client.ListFieldLevelEncryptionConfigs(input)
}

func (a *CloudFrontActivities) ListFieldLevelEncryptionProfiles(input *cloudfront.ListFieldLevelEncryptionProfilesInput) (*cloudfront.ListFieldLevelEncryptionProfilesOutput, error) {
    return a.client.ListFieldLevelEncryptionProfiles(input)
}

func (a *CloudFrontActivities) ListInvalidations(input *cloudfront.ListInvalidationsInput) (*cloudfront.ListInvalidationsOutput, error) {
    return a.client.ListInvalidations(input)
}

func (a *CloudFrontActivities) ListOriginRequestPolicies(input *cloudfront.ListOriginRequestPoliciesInput) (*cloudfront.ListOriginRequestPoliciesOutput, error) {
    return a.client.ListOriginRequestPolicies(input)
}

func (a *CloudFrontActivities) ListPublicKeys(input *cloudfront.ListPublicKeysInput) (*cloudfront.ListPublicKeysOutput, error) {
    return a.client.ListPublicKeys(input)
}

func (a *CloudFrontActivities) ListRealtimeLogConfigs(input *cloudfront.ListRealtimeLogConfigsInput) (*cloudfront.ListRealtimeLogConfigsOutput, error) {
    return a.client.ListRealtimeLogConfigs(input)
}

func (a *CloudFrontActivities) ListStreamingDistributions(input *cloudfront.ListStreamingDistributionsInput) (*cloudfront.ListStreamingDistributionsOutput, error) {
    return a.client.ListStreamingDistributions(input)
}

func (a *CloudFrontActivities) ListTagsForResource(input *cloudfront.ListTagsForResourceInput) (*cloudfront.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *CloudFrontActivities) TagResource(input *cloudfront.TagResourceInput) (*cloudfront.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *CloudFrontActivities) UntagResource(input *cloudfront.UntagResourceInput) (*cloudfront.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *CloudFrontActivities) UpdateCachePolicy(input *cloudfront.UpdateCachePolicyInput) (*cloudfront.UpdateCachePolicyOutput, error) {
    return a.client.UpdateCachePolicy(input)
}

func (a *CloudFrontActivities) UpdateCloudFrontOriginAccessIdentity(input *cloudfront.UpdateCloudFrontOriginAccessIdentityInput) (*cloudfront.UpdateCloudFrontOriginAccessIdentityOutput, error) {
    return a.client.UpdateCloudFrontOriginAccessIdentity(input)
}

func (a *CloudFrontActivities) UpdateDistribution(input *cloudfront.UpdateDistributionInput) (*cloudfront.UpdateDistributionOutput, error) {
    return a.client.UpdateDistribution(input)
}

func (a *CloudFrontActivities) UpdateFieldLevelEncryptionConfig(input *cloudfront.UpdateFieldLevelEncryptionConfigInput) (*cloudfront.UpdateFieldLevelEncryptionConfigOutput, error) {
    return a.client.UpdateFieldLevelEncryptionConfig(input)
}

func (a *CloudFrontActivities) UpdateFieldLevelEncryptionProfile(input *cloudfront.UpdateFieldLevelEncryptionProfileInput) (*cloudfront.UpdateFieldLevelEncryptionProfileOutput, error) {
    return a.client.UpdateFieldLevelEncryptionProfile(input)
}

func (a *CloudFrontActivities) UpdateOriginRequestPolicy(input *cloudfront.UpdateOriginRequestPolicyInput) (*cloudfront.UpdateOriginRequestPolicyOutput, error) {
    return a.client.UpdateOriginRequestPolicy(input)
}

func (a *CloudFrontActivities) UpdatePublicKey(input *cloudfront.UpdatePublicKeyInput) (*cloudfront.UpdatePublicKeyOutput, error) {
    return a.client.UpdatePublicKey(input)
}

func (a *CloudFrontActivities) UpdateRealtimeLogConfig(input *cloudfront.UpdateRealtimeLogConfigInput) (*cloudfront.UpdateRealtimeLogConfigOutput, error) {
    return a.client.UpdateRealtimeLogConfig(input)
}

func (a *CloudFrontActivities) UpdateStreamingDistribution(input *cloudfront.UpdateStreamingDistributionInput) (*cloudfront.UpdateStreamingDistributionOutput, error) {
    return a.client.UpdateStreamingDistribution(input)
}

func (a *CloudFrontActivities) WaitUntilDistributionDeployed(input *cloudfront.GetDistributionInput) error {
    return a.client.WaitUntilDistributionDeployed(input)
}

func (a *CloudFrontActivities) WaitUntilInvalidationCompleted(input *cloudfront.GetInvalidationInput) error {
    return a.client.WaitUntilInvalidationCompleted(input)
}

func (a *CloudFrontActivities) WaitUntilStreamingDistributionDeployed(input *cloudfront.GetStreamingDistributionInput) error {
    return a.client.WaitUntilStreamingDistributionDeployed(input)
}
