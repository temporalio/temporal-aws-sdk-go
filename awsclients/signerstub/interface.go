// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package signerstub

import (
	"github.com/aws/aws-sdk-go/service/signer"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

type Client interface {
	CancelSigningProfile(ctx workflow.Context, input *signer.CancelSigningProfileInput) (*signer.CancelSigningProfileOutput, error)
	CancelSigningProfileAsync(ctx workflow.Context, input *signer.CancelSigningProfileInput) *SignerCancelSigningProfileFuture

	DescribeSigningJob(ctx workflow.Context, input *signer.DescribeSigningJobInput) (*signer.DescribeSigningJobOutput, error)
	DescribeSigningJobAsync(ctx workflow.Context, input *signer.DescribeSigningJobInput) *SignerDescribeSigningJobFuture

	GetSigningPlatform(ctx workflow.Context, input *signer.GetSigningPlatformInput) (*signer.GetSigningPlatformOutput, error)
	GetSigningPlatformAsync(ctx workflow.Context, input *signer.GetSigningPlatformInput) *SignerGetSigningPlatformFuture

	GetSigningProfile(ctx workflow.Context, input *signer.GetSigningProfileInput) (*signer.GetSigningProfileOutput, error)
	GetSigningProfileAsync(ctx workflow.Context, input *signer.GetSigningProfileInput) *SignerGetSigningProfileFuture

	ListSigningJobs(ctx workflow.Context, input *signer.ListSigningJobsInput) (*signer.ListSigningJobsOutput, error)
	ListSigningJobsAsync(ctx workflow.Context, input *signer.ListSigningJobsInput) *SignerListSigningJobsFuture

	ListSigningPlatforms(ctx workflow.Context, input *signer.ListSigningPlatformsInput) (*signer.ListSigningPlatformsOutput, error)
	ListSigningPlatformsAsync(ctx workflow.Context, input *signer.ListSigningPlatformsInput) *SignerListSigningPlatformsFuture

	ListSigningProfiles(ctx workflow.Context, input *signer.ListSigningProfilesInput) (*signer.ListSigningProfilesOutput, error)
	ListSigningProfilesAsync(ctx workflow.Context, input *signer.ListSigningProfilesInput) *SignerListSigningProfilesFuture

	ListTagsForResource(ctx workflow.Context, input *signer.ListTagsForResourceInput) (*signer.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *signer.ListTagsForResourceInput) *SignerListTagsForResourceFuture

	PutSigningProfile(ctx workflow.Context, input *signer.PutSigningProfileInput) (*signer.PutSigningProfileOutput, error)
	PutSigningProfileAsync(ctx workflow.Context, input *signer.PutSigningProfileInput) *SignerPutSigningProfileFuture

	StartSigningJob(ctx workflow.Context, input *signer.StartSigningJobInput) (*signer.StartSigningJobOutput, error)
	StartSigningJobAsync(ctx workflow.Context, input *signer.StartSigningJobInput) *SignerStartSigningJobFuture

	TagResource(ctx workflow.Context, input *signer.TagResourceInput) (*signer.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *signer.TagResourceInput) *SignerTagResourceFuture

	UntagResource(ctx workflow.Context, input *signer.UntagResourceInput) (*signer.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *signer.UntagResourceInput) *SignerUntagResourceFuture

	WaitUntilSuccessfulSigningJob(ctx workflow.Context, input *signer.DescribeSigningJobInput) error
	WaitUntilSuccessfulSigningJobAsync(ctx workflow.Context, input *signer.DescribeSigningJobInput) *awsclients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}

