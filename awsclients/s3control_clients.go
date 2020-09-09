package awsclients

import (
	"github.com/aws/aws-sdk-go/service/s3control"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type S3ControlClient interface {
    CreateAccessPoint(ctx workflow.Context, input *s3control.CreateAccessPointInput) (*s3control.CreateAccessPointOutput, error)
    CreateAccessPointAsync(ctx workflow.Context, input *s3control.CreateAccessPointInput) *S3controlCreateAccessPointResult

    CreateJob(ctx workflow.Context, input *s3control.CreateJobInput) (*s3control.CreateJobOutput, error)
    CreateJobAsync(ctx workflow.Context, input *s3control.CreateJobInput) *S3controlCreateJobResult

    DeleteAccessPoint(ctx workflow.Context, input *s3control.DeleteAccessPointInput) (*s3control.DeleteAccessPointOutput, error)
    DeleteAccessPointAsync(ctx workflow.Context, input *s3control.DeleteAccessPointInput) *S3controlDeleteAccessPointResult

    DeleteAccessPointPolicy(ctx workflow.Context, input *s3control.DeleteAccessPointPolicyInput) (*s3control.DeleteAccessPointPolicyOutput, error)
    DeleteAccessPointPolicyAsync(ctx workflow.Context, input *s3control.DeleteAccessPointPolicyInput) *S3controlDeleteAccessPointPolicyResult

    DeleteJobTagging(ctx workflow.Context, input *s3control.DeleteJobTaggingInput) (*s3control.DeleteJobTaggingOutput, error)
    DeleteJobTaggingAsync(ctx workflow.Context, input *s3control.DeleteJobTaggingInput) *S3controlDeleteJobTaggingResult

    DeletePublicAccessBlock(ctx workflow.Context, input *s3control.DeletePublicAccessBlockInput) (*s3control.DeletePublicAccessBlockOutput, error)
    DeletePublicAccessBlockAsync(ctx workflow.Context, input *s3control.DeletePublicAccessBlockInput) *S3controlDeletePublicAccessBlockResult

    DescribeJob(ctx workflow.Context, input *s3control.DescribeJobInput) (*s3control.DescribeJobOutput, error)
    DescribeJobAsync(ctx workflow.Context, input *s3control.DescribeJobInput) *S3controlDescribeJobResult

    GetAccessPoint(ctx workflow.Context, input *s3control.GetAccessPointInput) (*s3control.GetAccessPointOutput, error)
    GetAccessPointAsync(ctx workflow.Context, input *s3control.GetAccessPointInput) *S3controlGetAccessPointResult

    GetAccessPointPolicy(ctx workflow.Context, input *s3control.GetAccessPointPolicyInput) (*s3control.GetAccessPointPolicyOutput, error)
    GetAccessPointPolicyAsync(ctx workflow.Context, input *s3control.GetAccessPointPolicyInput) *S3controlGetAccessPointPolicyResult

    GetAccessPointPolicyStatus(ctx workflow.Context, input *s3control.GetAccessPointPolicyStatusInput) (*s3control.GetAccessPointPolicyStatusOutput, error)
    GetAccessPointPolicyStatusAsync(ctx workflow.Context, input *s3control.GetAccessPointPolicyStatusInput) *S3controlGetAccessPointPolicyStatusResult

    GetJobTagging(ctx workflow.Context, input *s3control.GetJobTaggingInput) (*s3control.GetJobTaggingOutput, error)
    GetJobTaggingAsync(ctx workflow.Context, input *s3control.GetJobTaggingInput) *S3controlGetJobTaggingResult

    GetPublicAccessBlock(ctx workflow.Context, input *s3control.GetPublicAccessBlockInput) (*s3control.GetPublicAccessBlockOutput, error)
    GetPublicAccessBlockAsync(ctx workflow.Context, input *s3control.GetPublicAccessBlockInput) *S3controlGetPublicAccessBlockResult

    ListAccessPoints(ctx workflow.Context, input *s3control.ListAccessPointsInput) (*s3control.ListAccessPointsOutput, error)
    ListAccessPointsAsync(ctx workflow.Context, input *s3control.ListAccessPointsInput) *S3controlListAccessPointsResult

    ListJobs(ctx workflow.Context, input *s3control.ListJobsInput) (*s3control.ListJobsOutput, error)
    ListJobsAsync(ctx workflow.Context, input *s3control.ListJobsInput) *S3controlListJobsResult

    PutAccessPointPolicy(ctx workflow.Context, input *s3control.PutAccessPointPolicyInput) (*s3control.PutAccessPointPolicyOutput, error)
    PutAccessPointPolicyAsync(ctx workflow.Context, input *s3control.PutAccessPointPolicyInput) *S3controlPutAccessPointPolicyResult

    PutJobTagging(ctx workflow.Context, input *s3control.PutJobTaggingInput) (*s3control.PutJobTaggingOutput, error)
    PutJobTaggingAsync(ctx workflow.Context, input *s3control.PutJobTaggingInput) *S3controlPutJobTaggingResult

    PutPublicAccessBlock(ctx workflow.Context, input *s3control.PutPublicAccessBlockInput) (*s3control.PutPublicAccessBlockOutput, error)
    PutPublicAccessBlockAsync(ctx workflow.Context, input *s3control.PutPublicAccessBlockInput) *S3controlPutPublicAccessBlockResult

    UpdateJobPriority(ctx workflow.Context, input *s3control.UpdateJobPriorityInput) (*s3control.UpdateJobPriorityOutput, error)
    UpdateJobPriorityAsync(ctx workflow.Context, input *s3control.UpdateJobPriorityInput) *S3controlUpdateJobPriorityResult

    UpdateJobStatus(ctx workflow.Context, input *s3control.UpdateJobStatusInput) (*s3control.UpdateJobStatusOutput, error)
    UpdateJobStatusAsync(ctx workflow.Context, input *s3control.UpdateJobStatusInput) *S3controlUpdateJobStatusResult
}

type S3controlCreateAccessPointResult struct {
	Result workflow.Future
}

func (r *S3controlCreateAccessPointResult) Get(ctx workflow.Context) (*s3control.CreateAccessPointOutput, error) {
    var output s3control.CreateAccessPointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type S3controlCreateJobResult struct {
	Result workflow.Future
}

func (r *S3controlCreateJobResult) Get(ctx workflow.Context) (*s3control.CreateJobOutput, error) {
    var output s3control.CreateJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type S3controlDeleteAccessPointResult struct {
	Result workflow.Future
}

func (r *S3controlDeleteAccessPointResult) Get(ctx workflow.Context) (*s3control.DeleteAccessPointOutput, error) {
    var output s3control.DeleteAccessPointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type S3controlDeleteAccessPointPolicyResult struct {
	Result workflow.Future
}

func (r *S3controlDeleteAccessPointPolicyResult) Get(ctx workflow.Context) (*s3control.DeleteAccessPointPolicyOutput, error) {
    var output s3control.DeleteAccessPointPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type S3controlDeleteJobTaggingResult struct {
	Result workflow.Future
}

func (r *S3controlDeleteJobTaggingResult) Get(ctx workflow.Context) (*s3control.DeleteJobTaggingOutput, error) {
    var output s3control.DeleteJobTaggingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type S3controlDeletePublicAccessBlockResult struct {
	Result workflow.Future
}

func (r *S3controlDeletePublicAccessBlockResult) Get(ctx workflow.Context) (*s3control.DeletePublicAccessBlockOutput, error) {
    var output s3control.DeletePublicAccessBlockOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type S3controlDescribeJobResult struct {
	Result workflow.Future
}

func (r *S3controlDescribeJobResult) Get(ctx workflow.Context) (*s3control.DescribeJobOutput, error) {
    var output s3control.DescribeJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type S3controlGetAccessPointResult struct {
	Result workflow.Future
}

func (r *S3controlGetAccessPointResult) Get(ctx workflow.Context) (*s3control.GetAccessPointOutput, error) {
    var output s3control.GetAccessPointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type S3controlGetAccessPointPolicyResult struct {
	Result workflow.Future
}

func (r *S3controlGetAccessPointPolicyResult) Get(ctx workflow.Context) (*s3control.GetAccessPointPolicyOutput, error) {
    var output s3control.GetAccessPointPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type S3controlGetAccessPointPolicyStatusResult struct {
	Result workflow.Future
}

func (r *S3controlGetAccessPointPolicyStatusResult) Get(ctx workflow.Context) (*s3control.GetAccessPointPolicyStatusOutput, error) {
    var output s3control.GetAccessPointPolicyStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type S3controlGetJobTaggingResult struct {
	Result workflow.Future
}

func (r *S3controlGetJobTaggingResult) Get(ctx workflow.Context) (*s3control.GetJobTaggingOutput, error) {
    var output s3control.GetJobTaggingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type S3controlGetPublicAccessBlockResult struct {
	Result workflow.Future
}

func (r *S3controlGetPublicAccessBlockResult) Get(ctx workflow.Context) (*s3control.GetPublicAccessBlockOutput, error) {
    var output s3control.GetPublicAccessBlockOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type S3controlListAccessPointsResult struct {
	Result workflow.Future
}

func (r *S3controlListAccessPointsResult) Get(ctx workflow.Context) (*s3control.ListAccessPointsOutput, error) {
    var output s3control.ListAccessPointsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type S3controlListJobsResult struct {
	Result workflow.Future
}

func (r *S3controlListJobsResult) Get(ctx workflow.Context) (*s3control.ListJobsOutput, error) {
    var output s3control.ListJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type S3controlPutAccessPointPolicyResult struct {
	Result workflow.Future
}

func (r *S3controlPutAccessPointPolicyResult) Get(ctx workflow.Context) (*s3control.PutAccessPointPolicyOutput, error) {
    var output s3control.PutAccessPointPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type S3controlPutJobTaggingResult struct {
	Result workflow.Future
}

func (r *S3controlPutJobTaggingResult) Get(ctx workflow.Context) (*s3control.PutJobTaggingOutput, error) {
    var output s3control.PutJobTaggingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type S3controlPutPublicAccessBlockResult struct {
	Result workflow.Future
}

func (r *S3controlPutPublicAccessBlockResult) Get(ctx workflow.Context) (*s3control.PutPublicAccessBlockOutput, error) {
    var output s3control.PutPublicAccessBlockOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type S3controlUpdateJobPriorityResult struct {
	Result workflow.Future
}

func (r *S3controlUpdateJobPriorityResult) Get(ctx workflow.Context) (*s3control.UpdateJobPriorityOutput, error) {
    var output s3control.UpdateJobPriorityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type S3controlUpdateJobStatusResult struct {
	Result workflow.Future
}

func (r *S3controlUpdateJobStatusResult) Get(ctx workflow.Context) (*s3control.UpdateJobStatusOutput, error) {
    var output s3control.UpdateJobStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type S3ControlStub struct {
    activities awsactivities.S3ControlActivities
}

func NewS3ControlStub() S3ControlClient {
    return &S3ControlStub{}
}

func (a *S3ControlStub) CreateAccessPoint(ctx workflow.Context, input *s3control.CreateAccessPointInput) (*s3control.CreateAccessPointOutput, error) {
    var output s3control.CreateAccessPointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAccessPoint, input).Get(ctx, &output)
    return &output, err
}

func (a *S3ControlStub) CreateAccessPointAsync(ctx workflow.Context, input *s3control.CreateAccessPointInput) *S3controlCreateAccessPointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateAccessPoint, input)
    return &S3controlCreateAccessPointResult{Result: future}
}

func (a *S3ControlStub) CreateJob(ctx workflow.Context, input *s3control.CreateJobInput) (*s3control.CreateJobOutput, error) {
    var output s3control.CreateJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateJob, input).Get(ctx, &output)
    return &output, err
}

func (a *S3ControlStub) CreateJobAsync(ctx workflow.Context, input *s3control.CreateJobInput) *S3controlCreateJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateJob, input)
    return &S3controlCreateJobResult{Result: future}
}

func (a *S3ControlStub) DeleteAccessPoint(ctx workflow.Context, input *s3control.DeleteAccessPointInput) (*s3control.DeleteAccessPointOutput, error) {
    var output s3control.DeleteAccessPointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAccessPoint, input).Get(ctx, &output)
    return &output, err
}

func (a *S3ControlStub) DeleteAccessPointAsync(ctx workflow.Context, input *s3control.DeleteAccessPointInput) *S3controlDeleteAccessPointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteAccessPoint, input)
    return &S3controlDeleteAccessPointResult{Result: future}
}

func (a *S3ControlStub) DeleteAccessPointPolicy(ctx workflow.Context, input *s3control.DeleteAccessPointPolicyInput) (*s3control.DeleteAccessPointPolicyOutput, error) {
    var output s3control.DeleteAccessPointPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAccessPointPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *S3ControlStub) DeleteAccessPointPolicyAsync(ctx workflow.Context, input *s3control.DeleteAccessPointPolicyInput) *S3controlDeleteAccessPointPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteAccessPointPolicy, input)
    return &S3controlDeleteAccessPointPolicyResult{Result: future}
}

func (a *S3ControlStub) DeleteJobTagging(ctx workflow.Context, input *s3control.DeleteJobTaggingInput) (*s3control.DeleteJobTaggingOutput, error) {
    var output s3control.DeleteJobTaggingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteJobTagging, input).Get(ctx, &output)
    return &output, err
}

func (a *S3ControlStub) DeleteJobTaggingAsync(ctx workflow.Context, input *s3control.DeleteJobTaggingInput) *S3controlDeleteJobTaggingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteJobTagging, input)
    return &S3controlDeleteJobTaggingResult{Result: future}
}

func (a *S3ControlStub) DeletePublicAccessBlock(ctx workflow.Context, input *s3control.DeletePublicAccessBlockInput) (*s3control.DeletePublicAccessBlockOutput, error) {
    var output s3control.DeletePublicAccessBlockOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePublicAccessBlock, input).Get(ctx, &output)
    return &output, err
}

func (a *S3ControlStub) DeletePublicAccessBlockAsync(ctx workflow.Context, input *s3control.DeletePublicAccessBlockInput) *S3controlDeletePublicAccessBlockResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeletePublicAccessBlock, input)
    return &S3controlDeletePublicAccessBlockResult{Result: future}
}

func (a *S3ControlStub) DescribeJob(ctx workflow.Context, input *s3control.DescribeJobInput) (*s3control.DescribeJobOutput, error) {
    var output s3control.DescribeJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeJob, input).Get(ctx, &output)
    return &output, err
}

func (a *S3ControlStub) DescribeJobAsync(ctx workflow.Context, input *s3control.DescribeJobInput) *S3controlDescribeJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeJob, input)
    return &S3controlDescribeJobResult{Result: future}
}

func (a *S3ControlStub) GetAccessPoint(ctx workflow.Context, input *s3control.GetAccessPointInput) (*s3control.GetAccessPointOutput, error) {
    var output s3control.GetAccessPointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAccessPoint, input).Get(ctx, &output)
    return &output, err
}

func (a *S3ControlStub) GetAccessPointAsync(ctx workflow.Context, input *s3control.GetAccessPointInput) *S3controlGetAccessPointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAccessPoint, input)
    return &S3controlGetAccessPointResult{Result: future}
}

func (a *S3ControlStub) GetAccessPointPolicy(ctx workflow.Context, input *s3control.GetAccessPointPolicyInput) (*s3control.GetAccessPointPolicyOutput, error) {
    var output s3control.GetAccessPointPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAccessPointPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *S3ControlStub) GetAccessPointPolicyAsync(ctx workflow.Context, input *s3control.GetAccessPointPolicyInput) *S3controlGetAccessPointPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAccessPointPolicy, input)
    return &S3controlGetAccessPointPolicyResult{Result: future}
}

func (a *S3ControlStub) GetAccessPointPolicyStatus(ctx workflow.Context, input *s3control.GetAccessPointPolicyStatusInput) (*s3control.GetAccessPointPolicyStatusOutput, error) {
    var output s3control.GetAccessPointPolicyStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAccessPointPolicyStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *S3ControlStub) GetAccessPointPolicyStatusAsync(ctx workflow.Context, input *s3control.GetAccessPointPolicyStatusInput) *S3controlGetAccessPointPolicyStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAccessPointPolicyStatus, input)
    return &S3controlGetAccessPointPolicyStatusResult{Result: future}
}

func (a *S3ControlStub) GetJobTagging(ctx workflow.Context, input *s3control.GetJobTaggingInput) (*s3control.GetJobTaggingOutput, error) {
    var output s3control.GetJobTaggingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetJobTagging, input).Get(ctx, &output)
    return &output, err
}

func (a *S3ControlStub) GetJobTaggingAsync(ctx workflow.Context, input *s3control.GetJobTaggingInput) *S3controlGetJobTaggingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetJobTagging, input)
    return &S3controlGetJobTaggingResult{Result: future}
}

func (a *S3ControlStub) GetPublicAccessBlock(ctx workflow.Context, input *s3control.GetPublicAccessBlockInput) (*s3control.GetPublicAccessBlockOutput, error) {
    var output s3control.GetPublicAccessBlockOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPublicAccessBlock, input).Get(ctx, &output)
    return &output, err
}

func (a *S3ControlStub) GetPublicAccessBlockAsync(ctx workflow.Context, input *s3control.GetPublicAccessBlockInput) *S3controlGetPublicAccessBlockResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetPublicAccessBlock, input)
    return &S3controlGetPublicAccessBlockResult{Result: future}
}

func (a *S3ControlStub) ListAccessPoints(ctx workflow.Context, input *s3control.ListAccessPointsInput) (*s3control.ListAccessPointsOutput, error) {
    var output s3control.ListAccessPointsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAccessPoints, input).Get(ctx, &output)
    return &output, err
}

func (a *S3ControlStub) ListAccessPointsAsync(ctx workflow.Context, input *s3control.ListAccessPointsInput) *S3controlListAccessPointsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAccessPoints, input)
    return &S3controlListAccessPointsResult{Result: future}
}

func (a *S3ControlStub) ListJobs(ctx workflow.Context, input *s3control.ListJobsInput) (*s3control.ListJobsOutput, error) {
    var output s3control.ListJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *S3ControlStub) ListJobsAsync(ctx workflow.Context, input *s3control.ListJobsInput) *S3controlListJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListJobs, input)
    return &S3controlListJobsResult{Result: future}
}

func (a *S3ControlStub) PutAccessPointPolicy(ctx workflow.Context, input *s3control.PutAccessPointPolicyInput) (*s3control.PutAccessPointPolicyOutput, error) {
    var output s3control.PutAccessPointPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutAccessPointPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *S3ControlStub) PutAccessPointPolicyAsync(ctx workflow.Context, input *s3control.PutAccessPointPolicyInput) *S3controlPutAccessPointPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutAccessPointPolicy, input)
    return &S3controlPutAccessPointPolicyResult{Result: future}
}

func (a *S3ControlStub) PutJobTagging(ctx workflow.Context, input *s3control.PutJobTaggingInput) (*s3control.PutJobTaggingOutput, error) {
    var output s3control.PutJobTaggingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutJobTagging, input).Get(ctx, &output)
    return &output, err
}

func (a *S3ControlStub) PutJobTaggingAsync(ctx workflow.Context, input *s3control.PutJobTaggingInput) *S3controlPutJobTaggingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutJobTagging, input)
    return &S3controlPutJobTaggingResult{Result: future}
}

func (a *S3ControlStub) PutPublicAccessBlock(ctx workflow.Context, input *s3control.PutPublicAccessBlockInput) (*s3control.PutPublicAccessBlockOutput, error) {
    var output s3control.PutPublicAccessBlockOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutPublicAccessBlock, input).Get(ctx, &output)
    return &output, err
}

func (a *S3ControlStub) PutPublicAccessBlockAsync(ctx workflow.Context, input *s3control.PutPublicAccessBlockInput) *S3controlPutPublicAccessBlockResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutPublicAccessBlock, input)
    return &S3controlPutPublicAccessBlockResult{Result: future}
}

func (a *S3ControlStub) UpdateJobPriority(ctx workflow.Context, input *s3control.UpdateJobPriorityInput) (*s3control.UpdateJobPriorityOutput, error) {
    var output s3control.UpdateJobPriorityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateJobPriority, input).Get(ctx, &output)
    return &output, err
}

func (a *S3ControlStub) UpdateJobPriorityAsync(ctx workflow.Context, input *s3control.UpdateJobPriorityInput) *S3controlUpdateJobPriorityResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateJobPriority, input)
    return &S3controlUpdateJobPriorityResult{Result: future}
}

func (a *S3ControlStub) UpdateJobStatus(ctx workflow.Context, input *s3control.UpdateJobStatusInput) (*s3control.UpdateJobStatusOutput, error) {
    var output s3control.UpdateJobStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateJobStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *S3ControlStub) UpdateJobStatusAsync(ctx workflow.Context, input *s3control.UpdateJobStatusInput) *S3controlUpdateJobStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateJobStatus, input)
    return &S3controlUpdateJobStatusResult{Result: future}
}
