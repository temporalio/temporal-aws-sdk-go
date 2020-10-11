// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package s3controlstub

import (
	"github.com/aws/aws-sdk-go/service/s3control"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	CreateAccessPoint(ctx workflow.Context, input *s3control.CreateAccessPointInput) (*s3control.CreateAccessPointOutput, error)
	CreateAccessPointAsync(ctx workflow.Context, input *s3control.CreateAccessPointInput) *S3ControlCreateAccessPointFuture

	CreateBucket(ctx workflow.Context, input *s3control.CreateBucketInput) (*s3control.CreateBucketOutput, error)
	CreateBucketAsync(ctx workflow.Context, input *s3control.CreateBucketInput) *S3ControlCreateBucketFuture

	CreateJob(ctx workflow.Context, input *s3control.CreateJobInput) (*s3control.CreateJobOutput, error)
	CreateJobAsync(ctx workflow.Context, input *s3control.CreateJobInput) *S3ControlCreateJobFuture

	DeleteAccessPoint(ctx workflow.Context, input *s3control.DeleteAccessPointInput) (*s3control.DeleteAccessPointOutput, error)
	DeleteAccessPointAsync(ctx workflow.Context, input *s3control.DeleteAccessPointInput) *S3ControlDeleteAccessPointFuture

	DeleteAccessPointPolicy(ctx workflow.Context, input *s3control.DeleteAccessPointPolicyInput) (*s3control.DeleteAccessPointPolicyOutput, error)
	DeleteAccessPointPolicyAsync(ctx workflow.Context, input *s3control.DeleteAccessPointPolicyInput) *S3ControlDeleteAccessPointPolicyFuture

	DeleteBucket(ctx workflow.Context, input *s3control.DeleteBucketInput) (*s3control.DeleteBucketOutput, error)
	DeleteBucketAsync(ctx workflow.Context, input *s3control.DeleteBucketInput) *S3ControlDeleteBucketFuture

	DeleteBucketLifecycleConfiguration(ctx workflow.Context, input *s3control.DeleteBucketLifecycleConfigurationInput) (*s3control.DeleteBucketLifecycleConfigurationOutput, error)
	DeleteBucketLifecycleConfigurationAsync(ctx workflow.Context, input *s3control.DeleteBucketLifecycleConfigurationInput) *S3ControlDeleteBucketLifecycleConfigurationFuture

	DeleteBucketPolicy(ctx workflow.Context, input *s3control.DeleteBucketPolicyInput) (*s3control.DeleteBucketPolicyOutput, error)
	DeleteBucketPolicyAsync(ctx workflow.Context, input *s3control.DeleteBucketPolicyInput) *S3ControlDeleteBucketPolicyFuture

	DeleteBucketTagging(ctx workflow.Context, input *s3control.DeleteBucketTaggingInput) (*s3control.DeleteBucketTaggingOutput, error)
	DeleteBucketTaggingAsync(ctx workflow.Context, input *s3control.DeleteBucketTaggingInput) *S3ControlDeleteBucketTaggingFuture

	DeleteJobTagging(ctx workflow.Context, input *s3control.DeleteJobTaggingInput) (*s3control.DeleteJobTaggingOutput, error)
	DeleteJobTaggingAsync(ctx workflow.Context, input *s3control.DeleteJobTaggingInput) *S3ControlDeleteJobTaggingFuture

	DeletePublicAccessBlock(ctx workflow.Context, input *s3control.DeletePublicAccessBlockInput) (*s3control.DeletePublicAccessBlockOutput, error)
	DeletePublicAccessBlockAsync(ctx workflow.Context, input *s3control.DeletePublicAccessBlockInput) *S3ControlDeletePublicAccessBlockFuture

	DescribeJob(ctx workflow.Context, input *s3control.DescribeJobInput) (*s3control.DescribeJobOutput, error)
	DescribeJobAsync(ctx workflow.Context, input *s3control.DescribeJobInput) *S3ControlDescribeJobFuture

	GetAccessPoint(ctx workflow.Context, input *s3control.GetAccessPointInput) (*s3control.GetAccessPointOutput, error)
	GetAccessPointAsync(ctx workflow.Context, input *s3control.GetAccessPointInput) *S3ControlGetAccessPointFuture

	GetAccessPointPolicy(ctx workflow.Context, input *s3control.GetAccessPointPolicyInput) (*s3control.GetAccessPointPolicyOutput, error)
	GetAccessPointPolicyAsync(ctx workflow.Context, input *s3control.GetAccessPointPolicyInput) *S3ControlGetAccessPointPolicyFuture

	GetAccessPointPolicyStatus(ctx workflow.Context, input *s3control.GetAccessPointPolicyStatusInput) (*s3control.GetAccessPointPolicyStatusOutput, error)
	GetAccessPointPolicyStatusAsync(ctx workflow.Context, input *s3control.GetAccessPointPolicyStatusInput) *S3ControlGetAccessPointPolicyStatusFuture

	GetBucket(ctx workflow.Context, input *s3control.GetBucketInput) (*s3control.GetBucketOutput, error)
	GetBucketAsync(ctx workflow.Context, input *s3control.GetBucketInput) *S3ControlGetBucketFuture

	GetBucketLifecycleConfiguration(ctx workflow.Context, input *s3control.GetBucketLifecycleConfigurationInput) (*s3control.GetBucketLifecycleConfigurationOutput, error)
	GetBucketLifecycleConfigurationAsync(ctx workflow.Context, input *s3control.GetBucketLifecycleConfigurationInput) *S3ControlGetBucketLifecycleConfigurationFuture

	GetBucketPolicy(ctx workflow.Context, input *s3control.GetBucketPolicyInput) (*s3control.GetBucketPolicyOutput, error)
	GetBucketPolicyAsync(ctx workflow.Context, input *s3control.GetBucketPolicyInput) *S3ControlGetBucketPolicyFuture

	GetBucketTagging(ctx workflow.Context, input *s3control.GetBucketTaggingInput) (*s3control.GetBucketTaggingOutput, error)
	GetBucketTaggingAsync(ctx workflow.Context, input *s3control.GetBucketTaggingInput) *S3ControlGetBucketTaggingFuture

	GetJobTagging(ctx workflow.Context, input *s3control.GetJobTaggingInput) (*s3control.GetJobTaggingOutput, error)
	GetJobTaggingAsync(ctx workflow.Context, input *s3control.GetJobTaggingInput) *S3ControlGetJobTaggingFuture

	GetPublicAccessBlock(ctx workflow.Context, input *s3control.GetPublicAccessBlockInput) (*s3control.GetPublicAccessBlockOutput, error)
	GetPublicAccessBlockAsync(ctx workflow.Context, input *s3control.GetPublicAccessBlockInput) *S3ControlGetPublicAccessBlockFuture

	ListAccessPoints(ctx workflow.Context, input *s3control.ListAccessPointsInput) (*s3control.ListAccessPointsOutput, error)
	ListAccessPointsAsync(ctx workflow.Context, input *s3control.ListAccessPointsInput) *S3ControlListAccessPointsFuture

	ListJobs(ctx workflow.Context, input *s3control.ListJobsInput) (*s3control.ListJobsOutput, error)
	ListJobsAsync(ctx workflow.Context, input *s3control.ListJobsInput) *S3ControlListJobsFuture

	ListRegionalBuckets(ctx workflow.Context, input *s3control.ListRegionalBucketsInput) (*s3control.ListRegionalBucketsOutput, error)
	ListRegionalBucketsAsync(ctx workflow.Context, input *s3control.ListRegionalBucketsInput) *S3ControlListRegionalBucketsFuture

	PutAccessPointPolicy(ctx workflow.Context, input *s3control.PutAccessPointPolicyInput) (*s3control.PutAccessPointPolicyOutput, error)
	PutAccessPointPolicyAsync(ctx workflow.Context, input *s3control.PutAccessPointPolicyInput) *S3ControlPutAccessPointPolicyFuture

	PutBucketLifecycleConfiguration(ctx workflow.Context, input *s3control.PutBucketLifecycleConfigurationInput) (*s3control.PutBucketLifecycleConfigurationOutput, error)
	PutBucketLifecycleConfigurationAsync(ctx workflow.Context, input *s3control.PutBucketLifecycleConfigurationInput) *S3ControlPutBucketLifecycleConfigurationFuture

	PutBucketPolicy(ctx workflow.Context, input *s3control.PutBucketPolicyInput) (*s3control.PutBucketPolicyOutput, error)
	PutBucketPolicyAsync(ctx workflow.Context, input *s3control.PutBucketPolicyInput) *S3ControlPutBucketPolicyFuture

	PutBucketTagging(ctx workflow.Context, input *s3control.PutBucketTaggingInput) (*s3control.PutBucketTaggingOutput, error)
	PutBucketTaggingAsync(ctx workflow.Context, input *s3control.PutBucketTaggingInput) *S3ControlPutBucketTaggingFuture

	PutJobTagging(ctx workflow.Context, input *s3control.PutJobTaggingInput) (*s3control.PutJobTaggingOutput, error)
	PutJobTaggingAsync(ctx workflow.Context, input *s3control.PutJobTaggingInput) *S3ControlPutJobTaggingFuture

	PutPublicAccessBlock(ctx workflow.Context, input *s3control.PutPublicAccessBlockInput) (*s3control.PutPublicAccessBlockOutput, error)
	PutPublicAccessBlockAsync(ctx workflow.Context, input *s3control.PutPublicAccessBlockInput) *S3ControlPutPublicAccessBlockFuture

	UpdateJobPriority(ctx workflow.Context, input *s3control.UpdateJobPriorityInput) (*s3control.UpdateJobPriorityOutput, error)
	UpdateJobPriorityAsync(ctx workflow.Context, input *s3control.UpdateJobPriorityInput) *S3ControlUpdateJobPriorityFuture

	UpdateJobStatus(ctx workflow.Context, input *s3control.UpdateJobStatusInput) (*s3control.UpdateJobStatusOutput, error)
	UpdateJobStatusAsync(ctx workflow.Context, input *s3control.UpdateJobStatusInput) *S3ControlUpdateJobStatusFuture
}

func NewClient() Client {
	return &stub{}
}
