package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dlm"
	"github.com/aws/aws-sdk-go/service/dlm/dlmiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type DLMActivities struct {
	client dlmiface.DLMAPI
}

func NewDLMActivities(session *session.Session, config ...*aws.Config) *DLMActivities {
	client := dlm.New(session, config...)
	return &DLMActivities{client: client}
}

func (a *DLMActivities) CreateLifecyclePolicy(ctx context.Context, input *dlm.CreateLifecyclePolicyInput) (*dlm.CreateLifecyclePolicyOutput, error) {
	return a.client.CreateLifecyclePolicyWithContext(ctx, input)
}

func (a *DLMActivities) DeleteLifecyclePolicy(ctx context.Context, input *dlm.DeleteLifecyclePolicyInput) (*dlm.DeleteLifecyclePolicyOutput, error) {
	return a.client.DeleteLifecyclePolicyWithContext(ctx, input)
}

func (a *DLMActivities) GetLifecyclePolicies(ctx context.Context, input *dlm.GetLifecyclePoliciesInput) (*dlm.GetLifecyclePoliciesOutput, error) {
	return a.client.GetLifecyclePoliciesWithContext(ctx, input)
}

func (a *DLMActivities) GetLifecyclePolicy(ctx context.Context, input *dlm.GetLifecyclePolicyInput) (*dlm.GetLifecyclePolicyOutput, error) {
	return a.client.GetLifecyclePolicyWithContext(ctx, input)
}

func (a *DLMActivities) ListTagsForResource(ctx context.Context, input *dlm.ListTagsForResourceInput) (*dlm.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *DLMActivities) TagResource(ctx context.Context, input *dlm.TagResourceInput) (*dlm.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *DLMActivities) UntagResource(ctx context.Context, input *dlm.UntagResourceInput) (*dlm.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *DLMActivities) UpdateLifecyclePolicy(ctx context.Context, input *dlm.UpdateLifecyclePolicyInput) (*dlm.UpdateLifecyclePolicyOutput, error) {
	return a.client.UpdateLifecyclePolicyWithContext(ctx, input)
}
