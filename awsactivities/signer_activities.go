package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/signer"
	"github.com/aws/aws-sdk-go/service/signer/signeriface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type SignerActivities struct {
	client signeriface.SignerAPI
}

func NewSignerActivities(session *session.Session, config ...*aws.Config) *SignerActivities {
	client := signer.New(session, config...)
	return &SignerActivities{client: client}
}

func (a *SignerActivities) CancelSigningProfile(ctx context.Context, input *signer.CancelSigningProfileInput) (*signer.CancelSigningProfileOutput, error) {
	return a.client.CancelSigningProfileWithContext(ctx, input)
}

func (a *SignerActivities) DescribeSigningJob(ctx context.Context, input *signer.DescribeSigningJobInput) (*signer.DescribeSigningJobOutput, error) {
	return a.client.DescribeSigningJobWithContext(ctx, input)
}

func (a *SignerActivities) GetSigningPlatform(ctx context.Context, input *signer.GetSigningPlatformInput) (*signer.GetSigningPlatformOutput, error) {
	return a.client.GetSigningPlatformWithContext(ctx, input)
}

func (a *SignerActivities) GetSigningProfile(ctx context.Context, input *signer.GetSigningProfileInput) (*signer.GetSigningProfileOutput, error) {
	return a.client.GetSigningProfileWithContext(ctx, input)
}

func (a *SignerActivities) ListSigningJobs(ctx context.Context, input *signer.ListSigningJobsInput) (*signer.ListSigningJobsOutput, error) {
	return a.client.ListSigningJobsWithContext(ctx, input)
}

func (a *SignerActivities) ListSigningPlatforms(ctx context.Context, input *signer.ListSigningPlatformsInput) (*signer.ListSigningPlatformsOutput, error) {
	return a.client.ListSigningPlatformsWithContext(ctx, input)
}

func (a *SignerActivities) ListSigningProfiles(ctx context.Context, input *signer.ListSigningProfilesInput) (*signer.ListSigningProfilesOutput, error) {
	return a.client.ListSigningProfilesWithContext(ctx, input)
}

func (a *SignerActivities) ListTagsForResource(ctx context.Context, input *signer.ListTagsForResourceInput) (*signer.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *SignerActivities) PutSigningProfile(ctx context.Context, input *signer.PutSigningProfileInput) (*signer.PutSigningProfileOutput, error) {
	return a.client.PutSigningProfileWithContext(ctx, input)
}

func (a *SignerActivities) StartSigningJob(ctx context.Context, input *signer.StartSigningJobInput) (*signer.StartSigningJobOutput, error) {
	return a.client.StartSigningJobWithContext(ctx, input)
}

func (a *SignerActivities) TagResource(ctx context.Context, input *signer.TagResourceInput) (*signer.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *SignerActivities) UntagResource(ctx context.Context, input *signer.UntagResourceInput) (*signer.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *SignerActivities) WaitUntilSuccessfulSigningJob(ctx context.Context, input *signer.DescribeSigningJobInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilSuccessfulSigningJobWithContext(ctx, input, options...)
	})
}
