package awsclients

import (
	"github.com/aws/aws-sdk-go/service/cloud9"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type Cloud9Client interface {
	CreateEnvironmentEC2(ctx workflow.Context, input *cloud9.CreateEnvironmentEC2Input) (*cloud9.CreateEnvironmentEC2Output, error)
	CreateEnvironmentEC2Async(ctx workflow.Context, input *cloud9.CreateEnvironmentEC2Input) *Cloud9CreateEnvironmentEC2Result

	CreateEnvironmentMembership(ctx workflow.Context, input *cloud9.CreateEnvironmentMembershipInput) (*cloud9.CreateEnvironmentMembershipOutput, error)
	CreateEnvironmentMembershipAsync(ctx workflow.Context, input *cloud9.CreateEnvironmentMembershipInput) *Cloud9CreateEnvironmentMembershipResult

	DeleteEnvironment(ctx workflow.Context, input *cloud9.DeleteEnvironmentInput) (*cloud9.DeleteEnvironmentOutput, error)
	DeleteEnvironmentAsync(ctx workflow.Context, input *cloud9.DeleteEnvironmentInput) *Cloud9DeleteEnvironmentResult

	DeleteEnvironmentMembership(ctx workflow.Context, input *cloud9.DeleteEnvironmentMembershipInput) (*cloud9.DeleteEnvironmentMembershipOutput, error)
	DeleteEnvironmentMembershipAsync(ctx workflow.Context, input *cloud9.DeleteEnvironmentMembershipInput) *Cloud9DeleteEnvironmentMembershipResult

	DescribeEnvironmentMemberships(ctx workflow.Context, input *cloud9.DescribeEnvironmentMembershipsInput) (*cloud9.DescribeEnvironmentMembershipsOutput, error)
	DescribeEnvironmentMembershipsAsync(ctx workflow.Context, input *cloud9.DescribeEnvironmentMembershipsInput) *Cloud9DescribeEnvironmentMembershipsResult

	DescribeEnvironmentStatus(ctx workflow.Context, input *cloud9.DescribeEnvironmentStatusInput) (*cloud9.DescribeEnvironmentStatusOutput, error)
	DescribeEnvironmentStatusAsync(ctx workflow.Context, input *cloud9.DescribeEnvironmentStatusInput) *Cloud9DescribeEnvironmentStatusResult

	DescribeEnvironments(ctx workflow.Context, input *cloud9.DescribeEnvironmentsInput) (*cloud9.DescribeEnvironmentsOutput, error)
	DescribeEnvironmentsAsync(ctx workflow.Context, input *cloud9.DescribeEnvironmentsInput) *Cloud9DescribeEnvironmentsResult

	ListEnvironments(ctx workflow.Context, input *cloud9.ListEnvironmentsInput) (*cloud9.ListEnvironmentsOutput, error)
	ListEnvironmentsAsync(ctx workflow.Context, input *cloud9.ListEnvironmentsInput) *Cloud9ListEnvironmentsResult

	ListTagsForResource(ctx workflow.Context, input *cloud9.ListTagsForResourceInput) (*cloud9.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *cloud9.ListTagsForResourceInput) *Cloud9ListTagsForResourceResult

	TagResource(ctx workflow.Context, input *cloud9.TagResourceInput) (*cloud9.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *cloud9.TagResourceInput) *Cloud9TagResourceResult

	UntagResource(ctx workflow.Context, input *cloud9.UntagResourceInput) (*cloud9.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *cloud9.UntagResourceInput) *Cloud9UntagResourceResult

	UpdateEnvironment(ctx workflow.Context, input *cloud9.UpdateEnvironmentInput) (*cloud9.UpdateEnvironmentOutput, error)
	UpdateEnvironmentAsync(ctx workflow.Context, input *cloud9.UpdateEnvironmentInput) *Cloud9UpdateEnvironmentResult

	UpdateEnvironmentMembership(ctx workflow.Context, input *cloud9.UpdateEnvironmentMembershipInput) (*cloud9.UpdateEnvironmentMembershipOutput, error)
	UpdateEnvironmentMembershipAsync(ctx workflow.Context, input *cloud9.UpdateEnvironmentMembershipInput) *Cloud9UpdateEnvironmentMembershipResult
}

type Cloud9CreateEnvironmentEC2Result struct {
	Result workflow.Future
}

func (r *Cloud9CreateEnvironmentEC2Result) Get(ctx workflow.Context) (*cloud9.CreateEnvironmentEC2Output, error) {
	var output cloud9.CreateEnvironmentEC2Output
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Cloud9CreateEnvironmentMembershipResult struct {
	Result workflow.Future
}

func (r *Cloud9CreateEnvironmentMembershipResult) Get(ctx workflow.Context) (*cloud9.CreateEnvironmentMembershipOutput, error) {
	var output cloud9.CreateEnvironmentMembershipOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Cloud9DeleteEnvironmentResult struct {
	Result workflow.Future
}

func (r *Cloud9DeleteEnvironmentResult) Get(ctx workflow.Context) (*cloud9.DeleteEnvironmentOutput, error) {
	var output cloud9.DeleteEnvironmentOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Cloud9DeleteEnvironmentMembershipResult struct {
	Result workflow.Future
}

func (r *Cloud9DeleteEnvironmentMembershipResult) Get(ctx workflow.Context) (*cloud9.DeleteEnvironmentMembershipOutput, error) {
	var output cloud9.DeleteEnvironmentMembershipOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Cloud9DescribeEnvironmentMembershipsResult struct {
	Result workflow.Future
}

func (r *Cloud9DescribeEnvironmentMembershipsResult) Get(ctx workflow.Context) (*cloud9.DescribeEnvironmentMembershipsOutput, error) {
	var output cloud9.DescribeEnvironmentMembershipsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Cloud9DescribeEnvironmentStatusResult struct {
	Result workflow.Future
}

func (r *Cloud9DescribeEnvironmentStatusResult) Get(ctx workflow.Context) (*cloud9.DescribeEnvironmentStatusOutput, error) {
	var output cloud9.DescribeEnvironmentStatusOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Cloud9DescribeEnvironmentsResult struct {
	Result workflow.Future
}

func (r *Cloud9DescribeEnvironmentsResult) Get(ctx workflow.Context) (*cloud9.DescribeEnvironmentsOutput, error) {
	var output cloud9.DescribeEnvironmentsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Cloud9ListEnvironmentsResult struct {
	Result workflow.Future
}

func (r *Cloud9ListEnvironmentsResult) Get(ctx workflow.Context) (*cloud9.ListEnvironmentsOutput, error) {
	var output cloud9.ListEnvironmentsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Cloud9ListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *Cloud9ListTagsForResourceResult) Get(ctx workflow.Context) (*cloud9.ListTagsForResourceOutput, error) {
	var output cloud9.ListTagsForResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Cloud9TagResourceResult struct {
	Result workflow.Future
}

func (r *Cloud9TagResourceResult) Get(ctx workflow.Context) (*cloud9.TagResourceOutput, error) {
	var output cloud9.TagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Cloud9UntagResourceResult struct {
	Result workflow.Future
}

func (r *Cloud9UntagResourceResult) Get(ctx workflow.Context) (*cloud9.UntagResourceOutput, error) {
	var output cloud9.UntagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Cloud9UpdateEnvironmentResult struct {
	Result workflow.Future
}

func (r *Cloud9UpdateEnvironmentResult) Get(ctx workflow.Context) (*cloud9.UpdateEnvironmentOutput, error) {
	var output cloud9.UpdateEnvironmentOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Cloud9UpdateEnvironmentMembershipResult struct {
	Result workflow.Future
}

func (r *Cloud9UpdateEnvironmentMembershipResult) Get(ctx workflow.Context) (*cloud9.UpdateEnvironmentMembershipOutput, error) {
	var output cloud9.UpdateEnvironmentMembershipOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type Cloud9Stub struct {
	activities awsactivities.Cloud9Activities
}

func NewCloud9Stub() Cloud9Client {
	return &Cloud9Stub{}
}

func (a *Cloud9Stub) CreateEnvironmentEC2(ctx workflow.Context, input *cloud9.CreateEnvironmentEC2Input) (*cloud9.CreateEnvironmentEC2Output, error) {
	var output cloud9.CreateEnvironmentEC2Output
	err := workflow.ExecuteActivity(ctx, a.activities.CreateEnvironmentEC2, input).Get(ctx, &output)
	return &output, err
}

func (a *Cloud9Stub) CreateEnvironmentEC2Async(ctx workflow.Context, input *cloud9.CreateEnvironmentEC2Input) *Cloud9CreateEnvironmentEC2Result {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateEnvironmentEC2, input)
	return &Cloud9CreateEnvironmentEC2Result{Result: future}
}

func (a *Cloud9Stub) CreateEnvironmentMembership(ctx workflow.Context, input *cloud9.CreateEnvironmentMembershipInput) (*cloud9.CreateEnvironmentMembershipOutput, error) {
	var output cloud9.CreateEnvironmentMembershipOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateEnvironmentMembership, input).Get(ctx, &output)
	return &output, err
}

func (a *Cloud9Stub) CreateEnvironmentMembershipAsync(ctx workflow.Context, input *cloud9.CreateEnvironmentMembershipInput) *Cloud9CreateEnvironmentMembershipResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateEnvironmentMembership, input)
	return &Cloud9CreateEnvironmentMembershipResult{Result: future}
}

func (a *Cloud9Stub) DeleteEnvironment(ctx workflow.Context, input *cloud9.DeleteEnvironmentInput) (*cloud9.DeleteEnvironmentOutput, error) {
	var output cloud9.DeleteEnvironmentOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteEnvironment, input).Get(ctx, &output)
	return &output, err
}

func (a *Cloud9Stub) DeleteEnvironmentAsync(ctx workflow.Context, input *cloud9.DeleteEnvironmentInput) *Cloud9DeleteEnvironmentResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteEnvironment, input)
	return &Cloud9DeleteEnvironmentResult{Result: future}
}

func (a *Cloud9Stub) DeleteEnvironmentMembership(ctx workflow.Context, input *cloud9.DeleteEnvironmentMembershipInput) (*cloud9.DeleteEnvironmentMembershipOutput, error) {
	var output cloud9.DeleteEnvironmentMembershipOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteEnvironmentMembership, input).Get(ctx, &output)
	return &output, err
}

func (a *Cloud9Stub) DeleteEnvironmentMembershipAsync(ctx workflow.Context, input *cloud9.DeleteEnvironmentMembershipInput) *Cloud9DeleteEnvironmentMembershipResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteEnvironmentMembership, input)
	return &Cloud9DeleteEnvironmentMembershipResult{Result: future}
}

func (a *Cloud9Stub) DescribeEnvironmentMemberships(ctx workflow.Context, input *cloud9.DescribeEnvironmentMembershipsInput) (*cloud9.DescribeEnvironmentMembershipsOutput, error) {
	var output cloud9.DescribeEnvironmentMembershipsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeEnvironmentMemberships, input).Get(ctx, &output)
	return &output, err
}

func (a *Cloud9Stub) DescribeEnvironmentMembershipsAsync(ctx workflow.Context, input *cloud9.DescribeEnvironmentMembershipsInput) *Cloud9DescribeEnvironmentMembershipsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeEnvironmentMemberships, input)
	return &Cloud9DescribeEnvironmentMembershipsResult{Result: future}
}

func (a *Cloud9Stub) DescribeEnvironmentStatus(ctx workflow.Context, input *cloud9.DescribeEnvironmentStatusInput) (*cloud9.DescribeEnvironmentStatusOutput, error) {
	var output cloud9.DescribeEnvironmentStatusOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeEnvironmentStatus, input).Get(ctx, &output)
	return &output, err
}

func (a *Cloud9Stub) DescribeEnvironmentStatusAsync(ctx workflow.Context, input *cloud9.DescribeEnvironmentStatusInput) *Cloud9DescribeEnvironmentStatusResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeEnvironmentStatus, input)
	return &Cloud9DescribeEnvironmentStatusResult{Result: future}
}

func (a *Cloud9Stub) DescribeEnvironments(ctx workflow.Context, input *cloud9.DescribeEnvironmentsInput) (*cloud9.DescribeEnvironmentsOutput, error) {
	var output cloud9.DescribeEnvironmentsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeEnvironments, input).Get(ctx, &output)
	return &output, err
}

func (a *Cloud9Stub) DescribeEnvironmentsAsync(ctx workflow.Context, input *cloud9.DescribeEnvironmentsInput) *Cloud9DescribeEnvironmentsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeEnvironments, input)
	return &Cloud9DescribeEnvironmentsResult{Result: future}
}

func (a *Cloud9Stub) ListEnvironments(ctx workflow.Context, input *cloud9.ListEnvironmentsInput) (*cloud9.ListEnvironmentsOutput, error) {
	var output cloud9.ListEnvironmentsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListEnvironments, input).Get(ctx, &output)
	return &output, err
}

func (a *Cloud9Stub) ListEnvironmentsAsync(ctx workflow.Context, input *cloud9.ListEnvironmentsInput) *Cloud9ListEnvironmentsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListEnvironments, input)
	return &Cloud9ListEnvironmentsResult{Result: future}
}

func (a *Cloud9Stub) ListTagsForResource(ctx workflow.Context, input *cloud9.ListTagsForResourceInput) (*cloud9.ListTagsForResourceOutput, error) {
	var output cloud9.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
	return &output, err
}

func (a *Cloud9Stub) ListTagsForResourceAsync(ctx workflow.Context, input *cloud9.ListTagsForResourceInput) *Cloud9ListTagsForResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
	return &Cloud9ListTagsForResourceResult{Result: future}
}

func (a *Cloud9Stub) TagResource(ctx workflow.Context, input *cloud9.TagResourceInput) (*cloud9.TagResourceOutput, error) {
	var output cloud9.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *Cloud9Stub) TagResourceAsync(ctx workflow.Context, input *cloud9.TagResourceInput) *Cloud9TagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
	return &Cloud9TagResourceResult{Result: future}
}

func (a *Cloud9Stub) UntagResource(ctx workflow.Context, input *cloud9.UntagResourceInput) (*cloud9.UntagResourceOutput, error) {
	var output cloud9.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *Cloud9Stub) UntagResourceAsync(ctx workflow.Context, input *cloud9.UntagResourceInput) *Cloud9UntagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
	return &Cloud9UntagResourceResult{Result: future}
}

func (a *Cloud9Stub) UpdateEnvironment(ctx workflow.Context, input *cloud9.UpdateEnvironmentInput) (*cloud9.UpdateEnvironmentOutput, error) {
	var output cloud9.UpdateEnvironmentOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateEnvironment, input).Get(ctx, &output)
	return &output, err
}

func (a *Cloud9Stub) UpdateEnvironmentAsync(ctx workflow.Context, input *cloud9.UpdateEnvironmentInput) *Cloud9UpdateEnvironmentResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateEnvironment, input)
	return &Cloud9UpdateEnvironmentResult{Result: future}
}

func (a *Cloud9Stub) UpdateEnvironmentMembership(ctx workflow.Context, input *cloud9.UpdateEnvironmentMembershipInput) (*cloud9.UpdateEnvironmentMembershipOutput, error) {
	var output cloud9.UpdateEnvironmentMembershipOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateEnvironmentMembership, input).Get(ctx, &output)
	return &output, err
}

func (a *Cloud9Stub) UpdateEnvironmentMembershipAsync(ctx workflow.Context, input *cloud9.UpdateEnvironmentMembershipInput) *Cloud9UpdateEnvironmentMembershipResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateEnvironmentMembership, input)
	return &Cloud9UpdateEnvironmentMembershipResult{Result: future}
}
