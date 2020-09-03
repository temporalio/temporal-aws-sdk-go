package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/signer"
	"github.com/aws/aws-sdk-go/service/signer/signeriface"
)

type SignerActivities struct {
	client signeriface.SignerAPI
}

func NewSignerActivities(client signeriface.SignerAPI) *SignerActivities {
    return &SignerActivities{client: client}
}

func (a *SignerActivities) CancelSigningProfile(input *signer.CancelSigningProfileInput) (*signer.CancelSigningProfileOutput, error) {
    return a.client.CancelSigningProfile(input)
}

func (a *SignerActivities) DescribeSigningJob(input *signer.DescribeSigningJobInput) (*signer.DescribeSigningJobOutput, error) {
    return a.client.DescribeSigningJob(input)
}

func (a *SignerActivities) GetSigningPlatform(input *signer.GetSigningPlatformInput) (*signer.GetSigningPlatformOutput, error) {
    return a.client.GetSigningPlatform(input)
}

func (a *SignerActivities) GetSigningProfile(input *signer.GetSigningProfileInput) (*signer.GetSigningProfileOutput, error) {
    return a.client.GetSigningProfile(input)
}

func (a *SignerActivities) ListSigningJobs(input *signer.ListSigningJobsInput) (*signer.ListSigningJobsOutput, error) {
    return a.client.ListSigningJobs(input)
}

func (a *SignerActivities) ListSigningPlatforms(input *signer.ListSigningPlatformsInput) (*signer.ListSigningPlatformsOutput, error) {
    return a.client.ListSigningPlatforms(input)
}

func (a *SignerActivities) ListSigningProfiles(input *signer.ListSigningProfilesInput) (*signer.ListSigningProfilesOutput, error) {
    return a.client.ListSigningProfiles(input)
}

func (a *SignerActivities) ListTagsForResource(input *signer.ListTagsForResourceInput) (*signer.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *SignerActivities) PutSigningProfile(input *signer.PutSigningProfileInput) (*signer.PutSigningProfileOutput, error) {
    return a.client.PutSigningProfile(input)
}

func (a *SignerActivities) StartSigningJob(input *signer.StartSigningJobInput) (*signer.StartSigningJobOutput, error) {
    return a.client.StartSigningJob(input)
}

func (a *SignerActivities) TagResource(input *signer.TagResourceInput) (*signer.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *SignerActivities) UntagResource(input *signer.UntagResourceInput) (*signer.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *SignerActivities) WaitUntilSuccessfulSigningJob(input *signer.WaitUntilSuccessfulSigningJobInput) (*signer.WaitUntilSuccessfulSigningJobOutput, error) {
    return a.client.WaitUntilSuccessfulSigningJob(input)
}
