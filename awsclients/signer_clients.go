// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/signer"
	"go.temporal.io/sdk/workflow"
)

type SignerClient interface {
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
	WaitUntilSuccessfulSigningJobAsync(ctx workflow.Context, input *signer.DescribeSigningJobInput) workflow.Future
}

type SignerStub struct{}

func NewSignerStub() SignerClient {
	return &SignerStub{}
}

type SignerCancelSigningProfileFuture struct {
	Future workflow.Future
}

func (r *SignerCancelSigningProfileFuture) Get(ctx workflow.Context) (*signer.CancelSigningProfileOutput, error) {
	var output signer.CancelSigningProfileOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SignerDescribeSigningJobFuture struct {
	Future workflow.Future
}

func (r *SignerDescribeSigningJobFuture) Get(ctx workflow.Context) (*signer.DescribeSigningJobOutput, error) {
	var output signer.DescribeSigningJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SignerGetSigningPlatformFuture struct {
	Future workflow.Future
}

func (r *SignerGetSigningPlatformFuture) Get(ctx workflow.Context) (*signer.GetSigningPlatformOutput, error) {
	var output signer.GetSigningPlatformOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SignerGetSigningProfileFuture struct {
	Future workflow.Future
}

func (r *SignerGetSigningProfileFuture) Get(ctx workflow.Context) (*signer.GetSigningProfileOutput, error) {
	var output signer.GetSigningProfileOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SignerListSigningJobsFuture struct {
	Future workflow.Future
}

func (r *SignerListSigningJobsFuture) Get(ctx workflow.Context) (*signer.ListSigningJobsOutput, error) {
	var output signer.ListSigningJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SignerListSigningPlatformsFuture struct {
	Future workflow.Future
}

func (r *SignerListSigningPlatformsFuture) Get(ctx workflow.Context) (*signer.ListSigningPlatformsOutput, error) {
	var output signer.ListSigningPlatformsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SignerListSigningProfilesFuture struct {
	Future workflow.Future
}

func (r *SignerListSigningProfilesFuture) Get(ctx workflow.Context) (*signer.ListSigningProfilesOutput, error) {
	var output signer.ListSigningProfilesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SignerListTagsForResourceFuture struct {
	Future workflow.Future
}

func (r *SignerListTagsForResourceFuture) Get(ctx workflow.Context) (*signer.ListTagsForResourceOutput, error) {
	var output signer.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SignerPutSigningProfileFuture struct {
	Future workflow.Future
}

func (r *SignerPutSigningProfileFuture) Get(ctx workflow.Context) (*signer.PutSigningProfileOutput, error) {
	var output signer.PutSigningProfileOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SignerStartSigningJobFuture struct {
	Future workflow.Future
}

func (r *SignerStartSigningJobFuture) Get(ctx workflow.Context) (*signer.StartSigningJobOutput, error) {
	var output signer.StartSigningJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SignerTagResourceFuture struct {
	Future workflow.Future
}

func (r *SignerTagResourceFuture) Get(ctx workflow.Context) (*signer.TagResourceOutput, error) {
	var output signer.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SignerUntagResourceFuture struct {
	Future workflow.Future
}

func (r *SignerUntagResourceFuture) Get(ctx workflow.Context) (*signer.UntagResourceOutput, error) {
	var output signer.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *SignerStub) CancelSigningProfile(ctx workflow.Context, input *signer.CancelSigningProfileInput) (*signer.CancelSigningProfileOutput, error) {
	var output signer.CancelSigningProfileOutput
	err := workflow.ExecuteActivity(ctx, "aws.signer.CancelSigningProfile", input).Get(ctx, &output)
	return &output, err
}

func (a *SignerStub) CancelSigningProfileAsync(ctx workflow.Context, input *signer.CancelSigningProfileInput) *SignerCancelSigningProfileFuture {
	future := workflow.ExecuteActivity(ctx, "aws.signer.CancelSigningProfile", input)
	return &SignerCancelSigningProfileFuture{Future: future}
}

func (a *SignerStub) DescribeSigningJob(ctx workflow.Context, input *signer.DescribeSigningJobInput) (*signer.DescribeSigningJobOutput, error) {
	var output signer.DescribeSigningJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.signer.DescribeSigningJob", input).Get(ctx, &output)
	return &output, err
}

func (a *SignerStub) DescribeSigningJobAsync(ctx workflow.Context, input *signer.DescribeSigningJobInput) *SignerDescribeSigningJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.signer.DescribeSigningJob", input)
	return &SignerDescribeSigningJobFuture{Future: future}
}

func (a *SignerStub) GetSigningPlatform(ctx workflow.Context, input *signer.GetSigningPlatformInput) (*signer.GetSigningPlatformOutput, error) {
	var output signer.GetSigningPlatformOutput
	err := workflow.ExecuteActivity(ctx, "aws.signer.GetSigningPlatform", input).Get(ctx, &output)
	return &output, err
}

func (a *SignerStub) GetSigningPlatformAsync(ctx workflow.Context, input *signer.GetSigningPlatformInput) *SignerGetSigningPlatformFuture {
	future := workflow.ExecuteActivity(ctx, "aws.signer.GetSigningPlatform", input)
	return &SignerGetSigningPlatformFuture{Future: future}
}

func (a *SignerStub) GetSigningProfile(ctx workflow.Context, input *signer.GetSigningProfileInput) (*signer.GetSigningProfileOutput, error) {
	var output signer.GetSigningProfileOutput
	err := workflow.ExecuteActivity(ctx, "aws.signer.GetSigningProfile", input).Get(ctx, &output)
	return &output, err
}

func (a *SignerStub) GetSigningProfileAsync(ctx workflow.Context, input *signer.GetSigningProfileInput) *SignerGetSigningProfileFuture {
	future := workflow.ExecuteActivity(ctx, "aws.signer.GetSigningProfile", input)
	return &SignerGetSigningProfileFuture{Future: future}
}

func (a *SignerStub) ListSigningJobs(ctx workflow.Context, input *signer.ListSigningJobsInput) (*signer.ListSigningJobsOutput, error) {
	var output signer.ListSigningJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws.signer.ListSigningJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *SignerStub) ListSigningJobsAsync(ctx workflow.Context, input *signer.ListSigningJobsInput) *SignerListSigningJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.signer.ListSigningJobs", input)
	return &SignerListSigningJobsFuture{Future: future}
}

func (a *SignerStub) ListSigningPlatforms(ctx workflow.Context, input *signer.ListSigningPlatformsInput) (*signer.ListSigningPlatformsOutput, error) {
	var output signer.ListSigningPlatformsOutput
	err := workflow.ExecuteActivity(ctx, "aws.signer.ListSigningPlatforms", input).Get(ctx, &output)
	return &output, err
}

func (a *SignerStub) ListSigningPlatformsAsync(ctx workflow.Context, input *signer.ListSigningPlatformsInput) *SignerListSigningPlatformsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.signer.ListSigningPlatforms", input)
	return &SignerListSigningPlatformsFuture{Future: future}
}

func (a *SignerStub) ListSigningProfiles(ctx workflow.Context, input *signer.ListSigningProfilesInput) (*signer.ListSigningProfilesOutput, error) {
	var output signer.ListSigningProfilesOutput
	err := workflow.ExecuteActivity(ctx, "aws.signer.ListSigningProfiles", input).Get(ctx, &output)
	return &output, err
}

func (a *SignerStub) ListSigningProfilesAsync(ctx workflow.Context, input *signer.ListSigningProfilesInput) *SignerListSigningProfilesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.signer.ListSigningProfiles", input)
	return &SignerListSigningProfilesFuture{Future: future}
}

func (a *SignerStub) ListTagsForResource(ctx workflow.Context, input *signer.ListTagsForResourceInput) (*signer.ListTagsForResourceOutput, error) {
	var output signer.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.signer.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *SignerStub) ListTagsForResourceAsync(ctx workflow.Context, input *signer.ListTagsForResourceInput) *SignerListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.signer.ListTagsForResource", input)
	return &SignerListTagsForResourceFuture{Future: future}
}

func (a *SignerStub) PutSigningProfile(ctx workflow.Context, input *signer.PutSigningProfileInput) (*signer.PutSigningProfileOutput, error) {
	var output signer.PutSigningProfileOutput
	err := workflow.ExecuteActivity(ctx, "aws.signer.PutSigningProfile", input).Get(ctx, &output)
	return &output, err
}

func (a *SignerStub) PutSigningProfileAsync(ctx workflow.Context, input *signer.PutSigningProfileInput) *SignerPutSigningProfileFuture {
	future := workflow.ExecuteActivity(ctx, "aws.signer.PutSigningProfile", input)
	return &SignerPutSigningProfileFuture{Future: future}
}

func (a *SignerStub) StartSigningJob(ctx workflow.Context, input *signer.StartSigningJobInput) (*signer.StartSigningJobOutput, error) {
	var output signer.StartSigningJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.signer.StartSigningJob", input).Get(ctx, &output)
	return &output, err
}

func (a *SignerStub) StartSigningJobAsync(ctx workflow.Context, input *signer.StartSigningJobInput) *SignerStartSigningJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.signer.StartSigningJob", input)
	return &SignerStartSigningJobFuture{Future: future}
}

func (a *SignerStub) TagResource(ctx workflow.Context, input *signer.TagResourceInput) (*signer.TagResourceOutput, error) {
	var output signer.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.signer.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *SignerStub) TagResourceAsync(ctx workflow.Context, input *signer.TagResourceInput) *SignerTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.signer.TagResource", input)
	return &SignerTagResourceFuture{Future: future}
}

func (a *SignerStub) UntagResource(ctx workflow.Context, input *signer.UntagResourceInput) (*signer.UntagResourceOutput, error) {
	var output signer.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.signer.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *SignerStub) UntagResourceAsync(ctx workflow.Context, input *signer.UntagResourceInput) *SignerUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.signer.UntagResource", input)
	return &SignerUntagResourceFuture{Future: future}
}

func (a *SignerStub) WaitUntilSuccessfulSigningJob(ctx workflow.Context, input *signer.DescribeSigningJobInput) error {
	return workflow.ExecuteActivity(ctx, "aws.signer.WaitUntilSuccessfulSigningJob", input).Get(ctx, nil)
}

func (a *SignerStub) WaitUntilSuccessfulSigningJobAsync(ctx workflow.Context, input *signer.DescribeSigningJobInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, "aws.signer.WaitUntilSuccessfulSigningJob", input)
}
