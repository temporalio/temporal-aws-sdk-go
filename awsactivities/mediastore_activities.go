package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mediastore"
	"github.com/aws/aws-sdk-go/service/mediastore/mediastoreiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type MediaStoreActivities struct {
	client mediastoreiface.MediaStoreAPI
}

func NewMediaStoreActivities(session *session.Session, config ...*aws.Config) *MediaStoreActivities {
	client := mediastore.New(session, config...)
	return &MediaStoreActivities{client: client}
}

func (a *MediaStoreActivities) CreateContainer(ctx context.Context, input *mediastore.CreateContainerInput) (*mediastore.CreateContainerOutput, error) {
	return a.client.CreateContainerWithContext(ctx, input)
}

func (a *MediaStoreActivities) DeleteContainer(ctx context.Context, input *mediastore.DeleteContainerInput) (*mediastore.DeleteContainerOutput, error) {
	return a.client.DeleteContainerWithContext(ctx, input)
}

func (a *MediaStoreActivities) DeleteContainerPolicy(ctx context.Context, input *mediastore.DeleteContainerPolicyInput) (*mediastore.DeleteContainerPolicyOutput, error) {
	return a.client.DeleteContainerPolicyWithContext(ctx, input)
}

func (a *MediaStoreActivities) DeleteCorsPolicy(ctx context.Context, input *mediastore.DeleteCorsPolicyInput) (*mediastore.DeleteCorsPolicyOutput, error) {
	return a.client.DeleteCorsPolicyWithContext(ctx, input)
}

func (a *MediaStoreActivities) DeleteLifecyclePolicy(ctx context.Context, input *mediastore.DeleteLifecyclePolicyInput) (*mediastore.DeleteLifecyclePolicyOutput, error) {
	return a.client.DeleteLifecyclePolicyWithContext(ctx, input)
}

func (a *MediaStoreActivities) DeleteMetricPolicy(ctx context.Context, input *mediastore.DeleteMetricPolicyInput) (*mediastore.DeleteMetricPolicyOutput, error) {
	return a.client.DeleteMetricPolicyWithContext(ctx, input)
}

func (a *MediaStoreActivities) DescribeContainer(ctx context.Context, input *mediastore.DescribeContainerInput) (*mediastore.DescribeContainerOutput, error) {
	return a.client.DescribeContainerWithContext(ctx, input)
}

func (a *MediaStoreActivities) GetContainerPolicy(ctx context.Context, input *mediastore.GetContainerPolicyInput) (*mediastore.GetContainerPolicyOutput, error) {
	return a.client.GetContainerPolicyWithContext(ctx, input)
}

func (a *MediaStoreActivities) GetCorsPolicy(ctx context.Context, input *mediastore.GetCorsPolicyInput) (*mediastore.GetCorsPolicyOutput, error) {
	return a.client.GetCorsPolicyWithContext(ctx, input)
}

func (a *MediaStoreActivities) GetLifecyclePolicy(ctx context.Context, input *mediastore.GetLifecyclePolicyInput) (*mediastore.GetLifecyclePolicyOutput, error) {
	return a.client.GetLifecyclePolicyWithContext(ctx, input)
}

func (a *MediaStoreActivities) GetMetricPolicy(ctx context.Context, input *mediastore.GetMetricPolicyInput) (*mediastore.GetMetricPolicyOutput, error) {
	return a.client.GetMetricPolicyWithContext(ctx, input)
}

func (a *MediaStoreActivities) ListContainers(ctx context.Context, input *mediastore.ListContainersInput) (*mediastore.ListContainersOutput, error) {
	return a.client.ListContainersWithContext(ctx, input)
}

func (a *MediaStoreActivities) ListTagsForResource(ctx context.Context, input *mediastore.ListTagsForResourceInput) (*mediastore.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *MediaStoreActivities) PutContainerPolicy(ctx context.Context, input *mediastore.PutContainerPolicyInput) (*mediastore.PutContainerPolicyOutput, error) {
	return a.client.PutContainerPolicyWithContext(ctx, input)
}

func (a *MediaStoreActivities) PutCorsPolicy(ctx context.Context, input *mediastore.PutCorsPolicyInput) (*mediastore.PutCorsPolicyOutput, error) {
	return a.client.PutCorsPolicyWithContext(ctx, input)
}

func (a *MediaStoreActivities) PutLifecyclePolicy(ctx context.Context, input *mediastore.PutLifecyclePolicyInput) (*mediastore.PutLifecyclePolicyOutput, error) {
	return a.client.PutLifecyclePolicyWithContext(ctx, input)
}

func (a *MediaStoreActivities) PutMetricPolicy(ctx context.Context, input *mediastore.PutMetricPolicyInput) (*mediastore.PutMetricPolicyOutput, error) {
	return a.client.PutMetricPolicyWithContext(ctx, input)
}

func (a *MediaStoreActivities) StartAccessLogging(ctx context.Context, input *mediastore.StartAccessLoggingInput) (*mediastore.StartAccessLoggingOutput, error) {
	return a.client.StartAccessLoggingWithContext(ctx, input)
}

func (a *MediaStoreActivities) StopAccessLogging(ctx context.Context, input *mediastore.StopAccessLoggingInput) (*mediastore.StopAccessLoggingOutput, error) {
	return a.client.StopAccessLoggingWithContext(ctx, input)
}

func (a *MediaStoreActivities) TagResource(ctx context.Context, input *mediastore.TagResourceInput) (*mediastore.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *MediaStoreActivities) UntagResource(ctx context.Context, input *mediastore.UntagResourceInput) (*mediastore.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}
