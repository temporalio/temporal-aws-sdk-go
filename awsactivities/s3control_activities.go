package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/s3control"
	"github.com/aws/aws-sdk-go/service/s3control/s3controliface"
)

type S3ControlActivities struct {
	client s3controliface.S3ControlAPI
}

func NewS3ControlActivities(client s3controliface.S3ControlAPI) *S3ControlActivities {
    return &S3ControlActivities{client: client}
}

func (a *S3ControlActivities) CreateAccessPoint(input *s3control.CreateAccessPointInput) (*s3control.CreateAccessPointOutput, error) {
    return a.client.CreateAccessPoint(input)
}

func (a *S3ControlActivities) CreateJob(input *s3control.CreateJobInput) (*s3control.CreateJobOutput, error) {
    return a.client.CreateJob(input)
}

func (a *S3ControlActivities) DeleteAccessPoint(input *s3control.DeleteAccessPointInput) (*s3control.DeleteAccessPointOutput, error) {
    return a.client.DeleteAccessPoint(input)
}

func (a *S3ControlActivities) DeleteAccessPointPolicy(input *s3control.DeleteAccessPointPolicyInput) (*s3control.DeleteAccessPointPolicyOutput, error) {
    return a.client.DeleteAccessPointPolicy(input)
}

func (a *S3ControlActivities) DeleteJobTagging(input *s3control.DeleteJobTaggingInput) (*s3control.DeleteJobTaggingOutput, error) {
    return a.client.DeleteJobTagging(input)
}

func (a *S3ControlActivities) DeletePublicAccessBlock(input *s3control.DeletePublicAccessBlockInput) (*s3control.DeletePublicAccessBlockOutput, error) {
    return a.client.DeletePublicAccessBlock(input)
}

func (a *S3ControlActivities) DescribeJob(input *s3control.DescribeJobInput) (*s3control.DescribeJobOutput, error) {
    return a.client.DescribeJob(input)
}

func (a *S3ControlActivities) GetAccessPoint(input *s3control.GetAccessPointInput) (*s3control.GetAccessPointOutput, error) {
    return a.client.GetAccessPoint(input)
}

func (a *S3ControlActivities) GetAccessPointPolicy(input *s3control.GetAccessPointPolicyInput) (*s3control.GetAccessPointPolicyOutput, error) {
    return a.client.GetAccessPointPolicy(input)
}

func (a *S3ControlActivities) GetAccessPointPolicyStatus(input *s3control.GetAccessPointPolicyStatusInput) (*s3control.GetAccessPointPolicyStatusOutput, error) {
    return a.client.GetAccessPointPolicyStatus(input)
}

func (a *S3ControlActivities) GetJobTagging(input *s3control.GetJobTaggingInput) (*s3control.GetJobTaggingOutput, error) {
    return a.client.GetJobTagging(input)
}

func (a *S3ControlActivities) GetPublicAccessBlock(input *s3control.GetPublicAccessBlockInput) (*s3control.GetPublicAccessBlockOutput, error) {
    return a.client.GetPublicAccessBlock(input)
}

func (a *S3ControlActivities) ListAccessPoints(input *s3control.ListAccessPointsInput) (*s3control.ListAccessPointsOutput, error) {
    return a.client.ListAccessPoints(input)
}

func (a *S3ControlActivities) ListJobs(input *s3control.ListJobsInput) (*s3control.ListJobsOutput, error) {
    return a.client.ListJobs(input)
}

func (a *S3ControlActivities) PutAccessPointPolicy(input *s3control.PutAccessPointPolicyInput) (*s3control.PutAccessPointPolicyOutput, error) {
    return a.client.PutAccessPointPolicy(input)
}

func (a *S3ControlActivities) PutJobTagging(input *s3control.PutJobTaggingInput) (*s3control.PutJobTaggingOutput, error) {
    return a.client.PutJobTagging(input)
}

func (a *S3ControlActivities) PutPublicAccessBlock(input *s3control.PutPublicAccessBlockInput) (*s3control.PutPublicAccessBlockOutput, error) {
    return a.client.PutPublicAccessBlock(input)
}

func (a *S3ControlActivities) UpdateJobPriority(input *s3control.UpdateJobPriorityInput) (*s3control.UpdateJobPriorityOutput, error) {
    return a.client.UpdateJobPriority(input)
}

func (a *S3ControlActivities) UpdateJobStatus(input *s3control.UpdateJobStatusInput) (*s3control.UpdateJobStatusOutput, error) {
    return a.client.UpdateJobStatus(input)
}
