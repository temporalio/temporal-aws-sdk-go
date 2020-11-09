// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package cloudfrontstub

import (
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateCachePolicy(ctx workflow.Context, input *cloudfront.CreateCachePolicyInput) (*cloudfront.CreateCachePolicyOutput, error)
	CreateCachePolicyAsync(ctx workflow.Context, input *cloudfront.CreateCachePolicyInput) *CreateCachePolicyFuture

	CreateCloudFrontOriginAccessIdentity(ctx workflow.Context, input *cloudfront.CreateCloudFrontOriginAccessIdentityInput) (*cloudfront.CreateCloudFrontOriginAccessIdentityOutput, error)
	CreateCloudFrontOriginAccessIdentityAsync(ctx workflow.Context, input *cloudfront.CreateCloudFrontOriginAccessIdentityInput) *CreateCloudFrontOriginAccessIdentityFuture

	CreateDistribution(ctx workflow.Context, input *cloudfront.CreateDistributionInput) (*cloudfront.CreateDistributionOutput, error)
	CreateDistributionAsync(ctx workflow.Context, input *cloudfront.CreateDistributionInput) *CreateDistributionFuture

	CreateDistributionWithTags(ctx workflow.Context, input *cloudfront.CreateDistributionWithTagsInput) (*cloudfront.CreateDistributionWithTagsOutput, error)
	CreateDistributionWithTagsAsync(ctx workflow.Context, input *cloudfront.CreateDistributionWithTagsInput) *CreateDistributionWithTagsFuture

	CreateFieldLevelEncryptionConfig(ctx workflow.Context, input *cloudfront.CreateFieldLevelEncryptionConfigInput) (*cloudfront.CreateFieldLevelEncryptionConfigOutput, error)
	CreateFieldLevelEncryptionConfigAsync(ctx workflow.Context, input *cloudfront.CreateFieldLevelEncryptionConfigInput) *CreateFieldLevelEncryptionConfigFuture

	CreateFieldLevelEncryptionProfile(ctx workflow.Context, input *cloudfront.CreateFieldLevelEncryptionProfileInput) (*cloudfront.CreateFieldLevelEncryptionProfileOutput, error)
	CreateFieldLevelEncryptionProfileAsync(ctx workflow.Context, input *cloudfront.CreateFieldLevelEncryptionProfileInput) *CreateFieldLevelEncryptionProfileFuture

	CreateInvalidation(ctx workflow.Context, input *cloudfront.CreateInvalidationInput) (*cloudfront.CreateInvalidationOutput, error)
	CreateInvalidationAsync(ctx workflow.Context, input *cloudfront.CreateInvalidationInput) *CreateInvalidationFuture

	CreateMonitoringSubscription(ctx workflow.Context, input *cloudfront.CreateMonitoringSubscriptionInput) (*cloudfront.CreateMonitoringSubscriptionOutput, error)
	CreateMonitoringSubscriptionAsync(ctx workflow.Context, input *cloudfront.CreateMonitoringSubscriptionInput) *CreateMonitoringSubscriptionFuture

	CreateOriginRequestPolicy(ctx workflow.Context, input *cloudfront.CreateOriginRequestPolicyInput) (*cloudfront.CreateOriginRequestPolicyOutput, error)
	CreateOriginRequestPolicyAsync(ctx workflow.Context, input *cloudfront.CreateOriginRequestPolicyInput) *CreateOriginRequestPolicyFuture

	CreatePublicKey(ctx workflow.Context, input *cloudfront.CreatePublicKeyInput) (*cloudfront.CreatePublicKeyOutput, error)
	CreatePublicKeyAsync(ctx workflow.Context, input *cloudfront.CreatePublicKeyInput) *CreatePublicKeyFuture

	CreateRealtimeLogConfig(ctx workflow.Context, input *cloudfront.CreateRealtimeLogConfigInput) (*cloudfront.CreateRealtimeLogConfigOutput, error)
	CreateRealtimeLogConfigAsync(ctx workflow.Context, input *cloudfront.CreateRealtimeLogConfigInput) *CreateRealtimeLogConfigFuture

	CreateStreamingDistribution(ctx workflow.Context, input *cloudfront.CreateStreamingDistributionInput) (*cloudfront.CreateStreamingDistributionOutput, error)
	CreateStreamingDistributionAsync(ctx workflow.Context, input *cloudfront.CreateStreamingDistributionInput) *CreateStreamingDistributionFuture

	CreateStreamingDistributionWithTags(ctx workflow.Context, input *cloudfront.CreateStreamingDistributionWithTagsInput) (*cloudfront.CreateStreamingDistributionWithTagsOutput, error)
	CreateStreamingDistributionWithTagsAsync(ctx workflow.Context, input *cloudfront.CreateStreamingDistributionWithTagsInput) *CreateStreamingDistributionWithTagsFuture

	DeleteCachePolicy(ctx workflow.Context, input *cloudfront.DeleteCachePolicyInput) (*cloudfront.DeleteCachePolicyOutput, error)
	DeleteCachePolicyAsync(ctx workflow.Context, input *cloudfront.DeleteCachePolicyInput) *DeleteCachePolicyFuture

	DeleteCloudFrontOriginAccessIdentity(ctx workflow.Context, input *cloudfront.DeleteCloudFrontOriginAccessIdentityInput) (*cloudfront.DeleteCloudFrontOriginAccessIdentityOutput, error)
	DeleteCloudFrontOriginAccessIdentityAsync(ctx workflow.Context, input *cloudfront.DeleteCloudFrontOriginAccessIdentityInput) *DeleteCloudFrontOriginAccessIdentityFuture

	DeleteDistribution(ctx workflow.Context, input *cloudfront.DeleteDistributionInput) (*cloudfront.DeleteDistributionOutput, error)
	DeleteDistributionAsync(ctx workflow.Context, input *cloudfront.DeleteDistributionInput) *DeleteDistributionFuture

	DeleteFieldLevelEncryptionConfig(ctx workflow.Context, input *cloudfront.DeleteFieldLevelEncryptionConfigInput) (*cloudfront.DeleteFieldLevelEncryptionConfigOutput, error)
	DeleteFieldLevelEncryptionConfigAsync(ctx workflow.Context, input *cloudfront.DeleteFieldLevelEncryptionConfigInput) *DeleteFieldLevelEncryptionConfigFuture

	DeleteFieldLevelEncryptionProfile(ctx workflow.Context, input *cloudfront.DeleteFieldLevelEncryptionProfileInput) (*cloudfront.DeleteFieldLevelEncryptionProfileOutput, error)
	DeleteFieldLevelEncryptionProfileAsync(ctx workflow.Context, input *cloudfront.DeleteFieldLevelEncryptionProfileInput) *DeleteFieldLevelEncryptionProfileFuture

	DeleteMonitoringSubscription(ctx workflow.Context, input *cloudfront.DeleteMonitoringSubscriptionInput) (*cloudfront.DeleteMonitoringSubscriptionOutput, error)
	DeleteMonitoringSubscriptionAsync(ctx workflow.Context, input *cloudfront.DeleteMonitoringSubscriptionInput) *DeleteMonitoringSubscriptionFuture

	DeleteOriginRequestPolicy(ctx workflow.Context, input *cloudfront.DeleteOriginRequestPolicyInput) (*cloudfront.DeleteOriginRequestPolicyOutput, error)
	DeleteOriginRequestPolicyAsync(ctx workflow.Context, input *cloudfront.DeleteOriginRequestPolicyInput) *DeleteOriginRequestPolicyFuture

	DeletePublicKey(ctx workflow.Context, input *cloudfront.DeletePublicKeyInput) (*cloudfront.DeletePublicKeyOutput, error)
	DeletePublicKeyAsync(ctx workflow.Context, input *cloudfront.DeletePublicKeyInput) *DeletePublicKeyFuture

	DeleteRealtimeLogConfig(ctx workflow.Context, input *cloudfront.DeleteRealtimeLogConfigInput) (*cloudfront.DeleteRealtimeLogConfigOutput, error)
	DeleteRealtimeLogConfigAsync(ctx workflow.Context, input *cloudfront.DeleteRealtimeLogConfigInput) *DeleteRealtimeLogConfigFuture

	DeleteStreamingDistribution(ctx workflow.Context, input *cloudfront.DeleteStreamingDistributionInput) (*cloudfront.DeleteStreamingDistributionOutput, error)
	DeleteStreamingDistributionAsync(ctx workflow.Context, input *cloudfront.DeleteStreamingDistributionInput) *DeleteStreamingDistributionFuture

	GetCachePolicy(ctx workflow.Context, input *cloudfront.GetCachePolicyInput) (*cloudfront.GetCachePolicyOutput, error)
	GetCachePolicyAsync(ctx workflow.Context, input *cloudfront.GetCachePolicyInput) *GetCachePolicyFuture

	GetCachePolicyConfig(ctx workflow.Context, input *cloudfront.GetCachePolicyConfigInput) (*cloudfront.GetCachePolicyConfigOutput, error)
	GetCachePolicyConfigAsync(ctx workflow.Context, input *cloudfront.GetCachePolicyConfigInput) *GetCachePolicyConfigFuture

	GetCloudFrontOriginAccessIdentity(ctx workflow.Context, input *cloudfront.GetCloudFrontOriginAccessIdentityInput) (*cloudfront.GetCloudFrontOriginAccessIdentityOutput, error)
	GetCloudFrontOriginAccessIdentityAsync(ctx workflow.Context, input *cloudfront.GetCloudFrontOriginAccessIdentityInput) *GetCloudFrontOriginAccessIdentityFuture

	GetCloudFrontOriginAccessIdentityConfig(ctx workflow.Context, input *cloudfront.GetCloudFrontOriginAccessIdentityConfigInput) (*cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput, error)
	GetCloudFrontOriginAccessIdentityConfigAsync(ctx workflow.Context, input *cloudfront.GetCloudFrontOriginAccessIdentityConfigInput) *GetCloudFrontOriginAccessIdentityConfigFuture

	GetDistribution(ctx workflow.Context, input *cloudfront.GetDistributionInput) (*cloudfront.GetDistributionOutput, error)
	GetDistributionAsync(ctx workflow.Context, input *cloudfront.GetDistributionInput) *GetDistributionFuture

	GetDistributionConfig(ctx workflow.Context, input *cloudfront.GetDistributionConfigInput) (*cloudfront.GetDistributionConfigOutput, error)
	GetDistributionConfigAsync(ctx workflow.Context, input *cloudfront.GetDistributionConfigInput) *GetDistributionConfigFuture

	GetFieldLevelEncryption(ctx workflow.Context, input *cloudfront.GetFieldLevelEncryptionInput) (*cloudfront.GetFieldLevelEncryptionOutput, error)
	GetFieldLevelEncryptionAsync(ctx workflow.Context, input *cloudfront.GetFieldLevelEncryptionInput) *GetFieldLevelEncryptionFuture

	GetFieldLevelEncryptionConfig(ctx workflow.Context, input *cloudfront.GetFieldLevelEncryptionConfigInput) (*cloudfront.GetFieldLevelEncryptionConfigOutput, error)
	GetFieldLevelEncryptionConfigAsync(ctx workflow.Context, input *cloudfront.GetFieldLevelEncryptionConfigInput) *GetFieldLevelEncryptionConfigFuture

	GetFieldLevelEncryptionProfile(ctx workflow.Context, input *cloudfront.GetFieldLevelEncryptionProfileInput) (*cloudfront.GetFieldLevelEncryptionProfileOutput, error)
	GetFieldLevelEncryptionProfileAsync(ctx workflow.Context, input *cloudfront.GetFieldLevelEncryptionProfileInput) *GetFieldLevelEncryptionProfileFuture

	GetFieldLevelEncryptionProfileConfig(ctx workflow.Context, input *cloudfront.GetFieldLevelEncryptionProfileConfigInput) (*cloudfront.GetFieldLevelEncryptionProfileConfigOutput, error)
	GetFieldLevelEncryptionProfileConfigAsync(ctx workflow.Context, input *cloudfront.GetFieldLevelEncryptionProfileConfigInput) *GetFieldLevelEncryptionProfileConfigFuture

	GetInvalidation(ctx workflow.Context, input *cloudfront.GetInvalidationInput) (*cloudfront.GetInvalidationOutput, error)
	GetInvalidationAsync(ctx workflow.Context, input *cloudfront.GetInvalidationInput) *GetInvalidationFuture

	GetMonitoringSubscription(ctx workflow.Context, input *cloudfront.GetMonitoringSubscriptionInput) (*cloudfront.GetMonitoringSubscriptionOutput, error)
	GetMonitoringSubscriptionAsync(ctx workflow.Context, input *cloudfront.GetMonitoringSubscriptionInput) *GetMonitoringSubscriptionFuture

	GetOriginRequestPolicy(ctx workflow.Context, input *cloudfront.GetOriginRequestPolicyInput) (*cloudfront.GetOriginRequestPolicyOutput, error)
	GetOriginRequestPolicyAsync(ctx workflow.Context, input *cloudfront.GetOriginRequestPolicyInput) *GetOriginRequestPolicyFuture

	GetOriginRequestPolicyConfig(ctx workflow.Context, input *cloudfront.GetOriginRequestPolicyConfigInput) (*cloudfront.GetOriginRequestPolicyConfigOutput, error)
	GetOriginRequestPolicyConfigAsync(ctx workflow.Context, input *cloudfront.GetOriginRequestPolicyConfigInput) *GetOriginRequestPolicyConfigFuture

	GetPublicKey(ctx workflow.Context, input *cloudfront.GetPublicKeyInput) (*cloudfront.GetPublicKeyOutput, error)
	GetPublicKeyAsync(ctx workflow.Context, input *cloudfront.GetPublicKeyInput) *GetPublicKeyFuture

	GetPublicKeyConfig(ctx workflow.Context, input *cloudfront.GetPublicKeyConfigInput) (*cloudfront.GetPublicKeyConfigOutput, error)
	GetPublicKeyConfigAsync(ctx workflow.Context, input *cloudfront.GetPublicKeyConfigInput) *GetPublicKeyConfigFuture

	GetRealtimeLogConfig(ctx workflow.Context, input *cloudfront.GetRealtimeLogConfigInput) (*cloudfront.GetRealtimeLogConfigOutput, error)
	GetRealtimeLogConfigAsync(ctx workflow.Context, input *cloudfront.GetRealtimeLogConfigInput) *GetRealtimeLogConfigFuture

	GetStreamingDistribution(ctx workflow.Context, input *cloudfront.GetStreamingDistributionInput) (*cloudfront.GetStreamingDistributionOutput, error)
	GetStreamingDistributionAsync(ctx workflow.Context, input *cloudfront.GetStreamingDistributionInput) *GetStreamingDistributionFuture

	GetStreamingDistributionConfig(ctx workflow.Context, input *cloudfront.GetStreamingDistributionConfigInput) (*cloudfront.GetStreamingDistributionConfigOutput, error)
	GetStreamingDistributionConfigAsync(ctx workflow.Context, input *cloudfront.GetStreamingDistributionConfigInput) *GetStreamingDistributionConfigFuture

	ListCachePolicies(ctx workflow.Context, input *cloudfront.ListCachePoliciesInput) (*cloudfront.ListCachePoliciesOutput, error)
	ListCachePoliciesAsync(ctx workflow.Context, input *cloudfront.ListCachePoliciesInput) *ListCachePoliciesFuture

	ListCloudFrontOriginAccessIdentities(ctx workflow.Context, input *cloudfront.ListCloudFrontOriginAccessIdentitiesInput) (*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, error)
	ListCloudFrontOriginAccessIdentitiesAsync(ctx workflow.Context, input *cloudfront.ListCloudFrontOriginAccessIdentitiesInput) *ListCloudFrontOriginAccessIdentitiesFuture

	ListDistributions(ctx workflow.Context, input *cloudfront.ListDistributionsInput) (*cloudfront.ListDistributionsOutput, error)
	ListDistributionsAsync(ctx workflow.Context, input *cloudfront.ListDistributionsInput) *ListDistributionsFuture

	ListDistributionsByCachePolicyId(ctx workflow.Context, input *cloudfront.ListDistributionsByCachePolicyIdInput) (*cloudfront.ListDistributionsByCachePolicyIdOutput, error)
	ListDistributionsByCachePolicyIdAsync(ctx workflow.Context, input *cloudfront.ListDistributionsByCachePolicyIdInput) *ListDistributionsByCachePolicyIdFuture

	ListDistributionsByOriginRequestPolicyId(ctx workflow.Context, input *cloudfront.ListDistributionsByOriginRequestPolicyIdInput) (*cloudfront.ListDistributionsByOriginRequestPolicyIdOutput, error)
	ListDistributionsByOriginRequestPolicyIdAsync(ctx workflow.Context, input *cloudfront.ListDistributionsByOriginRequestPolicyIdInput) *ListDistributionsByOriginRequestPolicyIdFuture

	ListDistributionsByRealtimeLogConfig(ctx workflow.Context, input *cloudfront.ListDistributionsByRealtimeLogConfigInput) (*cloudfront.ListDistributionsByRealtimeLogConfigOutput, error)
	ListDistributionsByRealtimeLogConfigAsync(ctx workflow.Context, input *cloudfront.ListDistributionsByRealtimeLogConfigInput) *ListDistributionsByRealtimeLogConfigFuture

	ListDistributionsByWebACLId(ctx workflow.Context, input *cloudfront.ListDistributionsByWebACLIdInput) (*cloudfront.ListDistributionsByWebACLIdOutput, error)
	ListDistributionsByWebACLIdAsync(ctx workflow.Context, input *cloudfront.ListDistributionsByWebACLIdInput) *ListDistributionsByWebACLIdFuture

	ListFieldLevelEncryptionConfigs(ctx workflow.Context, input *cloudfront.ListFieldLevelEncryptionConfigsInput) (*cloudfront.ListFieldLevelEncryptionConfigsOutput, error)
	ListFieldLevelEncryptionConfigsAsync(ctx workflow.Context, input *cloudfront.ListFieldLevelEncryptionConfigsInput) *ListFieldLevelEncryptionConfigsFuture

	ListFieldLevelEncryptionProfiles(ctx workflow.Context, input *cloudfront.ListFieldLevelEncryptionProfilesInput) (*cloudfront.ListFieldLevelEncryptionProfilesOutput, error)
	ListFieldLevelEncryptionProfilesAsync(ctx workflow.Context, input *cloudfront.ListFieldLevelEncryptionProfilesInput) *ListFieldLevelEncryptionProfilesFuture

	ListInvalidations(ctx workflow.Context, input *cloudfront.ListInvalidationsInput) (*cloudfront.ListInvalidationsOutput, error)
	ListInvalidationsAsync(ctx workflow.Context, input *cloudfront.ListInvalidationsInput) *ListInvalidationsFuture

	ListOriginRequestPolicies(ctx workflow.Context, input *cloudfront.ListOriginRequestPoliciesInput) (*cloudfront.ListOriginRequestPoliciesOutput, error)
	ListOriginRequestPoliciesAsync(ctx workflow.Context, input *cloudfront.ListOriginRequestPoliciesInput) *ListOriginRequestPoliciesFuture

	ListPublicKeys(ctx workflow.Context, input *cloudfront.ListPublicKeysInput) (*cloudfront.ListPublicKeysOutput, error)
	ListPublicKeysAsync(ctx workflow.Context, input *cloudfront.ListPublicKeysInput) *ListPublicKeysFuture

	ListRealtimeLogConfigs(ctx workflow.Context, input *cloudfront.ListRealtimeLogConfigsInput) (*cloudfront.ListRealtimeLogConfigsOutput, error)
	ListRealtimeLogConfigsAsync(ctx workflow.Context, input *cloudfront.ListRealtimeLogConfigsInput) *ListRealtimeLogConfigsFuture

	ListStreamingDistributions(ctx workflow.Context, input *cloudfront.ListStreamingDistributionsInput) (*cloudfront.ListStreamingDistributionsOutput, error)
	ListStreamingDistributionsAsync(ctx workflow.Context, input *cloudfront.ListStreamingDistributionsInput) *ListStreamingDistributionsFuture

	ListTagsForResource(ctx workflow.Context, input *cloudfront.ListTagsForResourceInput) (*cloudfront.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *cloudfront.ListTagsForResourceInput) *ListTagsForResourceFuture

	TagResource(ctx workflow.Context, input *cloudfront.TagResourceInput) (*cloudfront.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *cloudfront.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *cloudfront.UntagResourceInput) (*cloudfront.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *cloudfront.UntagResourceInput) *UntagResourceFuture

	UpdateCachePolicy(ctx workflow.Context, input *cloudfront.UpdateCachePolicyInput) (*cloudfront.UpdateCachePolicyOutput, error)
	UpdateCachePolicyAsync(ctx workflow.Context, input *cloudfront.UpdateCachePolicyInput) *UpdateCachePolicyFuture

	UpdateCloudFrontOriginAccessIdentity(ctx workflow.Context, input *cloudfront.UpdateCloudFrontOriginAccessIdentityInput) (*cloudfront.UpdateCloudFrontOriginAccessIdentityOutput, error)
	UpdateCloudFrontOriginAccessIdentityAsync(ctx workflow.Context, input *cloudfront.UpdateCloudFrontOriginAccessIdentityInput) *UpdateCloudFrontOriginAccessIdentityFuture

	UpdateDistribution(ctx workflow.Context, input *cloudfront.UpdateDistributionInput) (*cloudfront.UpdateDistributionOutput, error)
	UpdateDistributionAsync(ctx workflow.Context, input *cloudfront.UpdateDistributionInput) *UpdateDistributionFuture

	UpdateFieldLevelEncryptionConfig(ctx workflow.Context, input *cloudfront.UpdateFieldLevelEncryptionConfigInput) (*cloudfront.UpdateFieldLevelEncryptionConfigOutput, error)
	UpdateFieldLevelEncryptionConfigAsync(ctx workflow.Context, input *cloudfront.UpdateFieldLevelEncryptionConfigInput) *UpdateFieldLevelEncryptionConfigFuture

	UpdateFieldLevelEncryptionProfile(ctx workflow.Context, input *cloudfront.UpdateFieldLevelEncryptionProfileInput) (*cloudfront.UpdateFieldLevelEncryptionProfileOutput, error)
	UpdateFieldLevelEncryptionProfileAsync(ctx workflow.Context, input *cloudfront.UpdateFieldLevelEncryptionProfileInput) *UpdateFieldLevelEncryptionProfileFuture

	UpdateOriginRequestPolicy(ctx workflow.Context, input *cloudfront.UpdateOriginRequestPolicyInput) (*cloudfront.UpdateOriginRequestPolicyOutput, error)
	UpdateOriginRequestPolicyAsync(ctx workflow.Context, input *cloudfront.UpdateOriginRequestPolicyInput) *UpdateOriginRequestPolicyFuture

	UpdatePublicKey(ctx workflow.Context, input *cloudfront.UpdatePublicKeyInput) (*cloudfront.UpdatePublicKeyOutput, error)
	UpdatePublicKeyAsync(ctx workflow.Context, input *cloudfront.UpdatePublicKeyInput) *UpdatePublicKeyFuture

	UpdateRealtimeLogConfig(ctx workflow.Context, input *cloudfront.UpdateRealtimeLogConfigInput) (*cloudfront.UpdateRealtimeLogConfigOutput, error)
	UpdateRealtimeLogConfigAsync(ctx workflow.Context, input *cloudfront.UpdateRealtimeLogConfigInput) *UpdateRealtimeLogConfigFuture

	UpdateStreamingDistribution(ctx workflow.Context, input *cloudfront.UpdateStreamingDistributionInput) (*cloudfront.UpdateStreamingDistributionOutput, error)
	UpdateStreamingDistributionAsync(ctx workflow.Context, input *cloudfront.UpdateStreamingDistributionInput) *UpdateStreamingDistributionFuture

	WaitUntilDistributionDeployed(ctx workflow.Context, input *cloudfront.GetDistributionInput) error
	WaitUntilDistributionDeployedAsync(ctx workflow.Context, input *cloudfront.GetDistributionInput) *clients.VoidFuture

	WaitUntilInvalidationCompleted(ctx workflow.Context, input *cloudfront.GetInvalidationInput) error
	WaitUntilInvalidationCompletedAsync(ctx workflow.Context, input *cloudfront.GetInvalidationInput) *clients.VoidFuture

	WaitUntilStreamingDistributionDeployed(ctx workflow.Context, input *cloudfront.GetStreamingDistributionInput) error
	WaitUntilStreamingDistributionDeployedAsync(ctx workflow.Context, input *cloudfront.GetStreamingDistributionInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
