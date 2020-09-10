package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/synthetics"
	"github.com/aws/aws-sdk-go/service/synthetics/syntheticsiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type SyntheticsActivities struct {
	client syntheticsiface.SyntheticsAPI
}

func NewSyntheticsActivities(session *session.Session, config ...*aws.Config) *SyntheticsActivities {
	client := synthetics.New(session, config...)
	return &SyntheticsActivities{client: client}
}

func (a *SyntheticsActivities) CreateCanary(ctx context.Context, input *synthetics.CreateCanaryInput) (*synthetics.CreateCanaryOutput, error) {
	return a.client.CreateCanaryWithContext(ctx, input)
}

func (a *SyntheticsActivities) DeleteCanary(ctx context.Context, input *synthetics.DeleteCanaryInput) (*synthetics.DeleteCanaryOutput, error) {
	return a.client.DeleteCanaryWithContext(ctx, input)
}

func (a *SyntheticsActivities) DescribeCanaries(ctx context.Context, input *synthetics.DescribeCanariesInput) (*synthetics.DescribeCanariesOutput, error) {
	return a.client.DescribeCanariesWithContext(ctx, input)
}

func (a *SyntheticsActivities) DescribeCanariesLastRun(ctx context.Context, input *synthetics.DescribeCanariesLastRunInput) (*synthetics.DescribeCanariesLastRunOutput, error) {
	return a.client.DescribeCanariesLastRunWithContext(ctx, input)
}

func (a *SyntheticsActivities) DescribeRuntimeVersions(ctx context.Context, input *synthetics.DescribeRuntimeVersionsInput) (*synthetics.DescribeRuntimeVersionsOutput, error) {
	return a.client.DescribeRuntimeVersionsWithContext(ctx, input)
}

func (a *SyntheticsActivities) GetCanary(ctx context.Context, input *synthetics.GetCanaryInput) (*synthetics.GetCanaryOutput, error) {
	return a.client.GetCanaryWithContext(ctx, input)
}

func (a *SyntheticsActivities) GetCanaryRuns(ctx context.Context, input *synthetics.GetCanaryRunsInput) (*synthetics.GetCanaryRunsOutput, error) {
	return a.client.GetCanaryRunsWithContext(ctx, input)
}

func (a *SyntheticsActivities) ListTagsForResource(ctx context.Context, input *synthetics.ListTagsForResourceInput) (*synthetics.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *SyntheticsActivities) StartCanary(ctx context.Context, input *synthetics.StartCanaryInput) (*synthetics.StartCanaryOutput, error) {
	return a.client.StartCanaryWithContext(ctx, input)
}

func (a *SyntheticsActivities) StopCanary(ctx context.Context, input *synthetics.StopCanaryInput) (*synthetics.StopCanaryOutput, error) {
	return a.client.StopCanaryWithContext(ctx, input)
}

func (a *SyntheticsActivities) TagResource(ctx context.Context, input *synthetics.TagResourceInput) (*synthetics.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *SyntheticsActivities) UntagResource(ctx context.Context, input *synthetics.UntagResourceInput) (*synthetics.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *SyntheticsActivities) UpdateCanary(ctx context.Context, input *synthetics.UpdateCanaryInput) (*synthetics.UpdateCanaryOutput, error) {
	return a.client.UpdateCanaryWithContext(ctx, input)
}
