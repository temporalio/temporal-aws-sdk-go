package awsclients

import (
	"github.com/aws/aws-sdk-go/service/synthetics"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type SyntheticsClient interface {
    CreateCanary(ctx workflow.Context, input *synthetics.CreateCanaryInput) (*synthetics.CreateCanaryOutput, error)
    CreateCanaryAsync(ctx workflow.Context, input *synthetics.CreateCanaryInput) *SyntheticsCreateCanaryResult

    DeleteCanary(ctx workflow.Context, input *synthetics.DeleteCanaryInput) (*synthetics.DeleteCanaryOutput, error)
    DeleteCanaryAsync(ctx workflow.Context, input *synthetics.DeleteCanaryInput) *SyntheticsDeleteCanaryResult

    DescribeCanaries(ctx workflow.Context, input *synthetics.DescribeCanariesInput) (*synthetics.DescribeCanariesOutput, error)
    DescribeCanariesAsync(ctx workflow.Context, input *synthetics.DescribeCanariesInput) *SyntheticsDescribeCanariesResult

    DescribeCanariesLastRun(ctx workflow.Context, input *synthetics.DescribeCanariesLastRunInput) (*synthetics.DescribeCanariesLastRunOutput, error)
    DescribeCanariesLastRunAsync(ctx workflow.Context, input *synthetics.DescribeCanariesLastRunInput) *SyntheticsDescribeCanariesLastRunResult

    DescribeRuntimeVersions(ctx workflow.Context, input *synthetics.DescribeRuntimeVersionsInput) (*synthetics.DescribeRuntimeVersionsOutput, error)
    DescribeRuntimeVersionsAsync(ctx workflow.Context, input *synthetics.DescribeRuntimeVersionsInput) *SyntheticsDescribeRuntimeVersionsResult

    GetCanary(ctx workflow.Context, input *synthetics.GetCanaryInput) (*synthetics.GetCanaryOutput, error)
    GetCanaryAsync(ctx workflow.Context, input *synthetics.GetCanaryInput) *SyntheticsGetCanaryResult

    GetCanaryRuns(ctx workflow.Context, input *synthetics.GetCanaryRunsInput) (*synthetics.GetCanaryRunsOutput, error)
    GetCanaryRunsAsync(ctx workflow.Context, input *synthetics.GetCanaryRunsInput) *SyntheticsGetCanaryRunsResult

    ListTagsForResource(ctx workflow.Context, input *synthetics.ListTagsForResourceInput) (*synthetics.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *synthetics.ListTagsForResourceInput) *SyntheticsListTagsForResourceResult

    StartCanary(ctx workflow.Context, input *synthetics.StartCanaryInput) (*synthetics.StartCanaryOutput, error)
    StartCanaryAsync(ctx workflow.Context, input *synthetics.StartCanaryInput) *SyntheticsStartCanaryResult

    StopCanary(ctx workflow.Context, input *synthetics.StopCanaryInput) (*synthetics.StopCanaryOutput, error)
    StopCanaryAsync(ctx workflow.Context, input *synthetics.StopCanaryInput) *SyntheticsStopCanaryResult

    TagResource(ctx workflow.Context, input *synthetics.TagResourceInput) (*synthetics.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *synthetics.TagResourceInput) *SyntheticsTagResourceResult

    UntagResource(ctx workflow.Context, input *synthetics.UntagResourceInput) (*synthetics.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *synthetics.UntagResourceInput) *SyntheticsUntagResourceResult

    UpdateCanary(ctx workflow.Context, input *synthetics.UpdateCanaryInput) (*synthetics.UpdateCanaryOutput, error)
    UpdateCanaryAsync(ctx workflow.Context, input *synthetics.UpdateCanaryInput) *SyntheticsUpdateCanaryResult
}
type SyntheticsCreateCanaryResult struct {
	Result workflow.Future
}

func (r *SyntheticsCreateCanaryResult) Get(ctx workflow.Context) (*synthetics.CreateCanaryOutput, error) {
    var output synthetics.CreateCanaryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SyntheticsDeleteCanaryResult struct {
	Result workflow.Future
}

func (r *SyntheticsDeleteCanaryResult) Get(ctx workflow.Context) (*synthetics.DeleteCanaryOutput, error) {
    var output synthetics.DeleteCanaryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SyntheticsDescribeCanariesResult struct {
	Result workflow.Future
}

func (r *SyntheticsDescribeCanariesResult) Get(ctx workflow.Context) (*synthetics.DescribeCanariesOutput, error) {
    var output synthetics.DescribeCanariesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SyntheticsDescribeCanariesLastRunResult struct {
	Result workflow.Future
}

func (r *SyntheticsDescribeCanariesLastRunResult) Get(ctx workflow.Context) (*synthetics.DescribeCanariesLastRunOutput, error) {
    var output synthetics.DescribeCanariesLastRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SyntheticsDescribeRuntimeVersionsResult struct {
	Result workflow.Future
}

func (r *SyntheticsDescribeRuntimeVersionsResult) Get(ctx workflow.Context) (*synthetics.DescribeRuntimeVersionsOutput, error) {
    var output synthetics.DescribeRuntimeVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SyntheticsGetCanaryResult struct {
	Result workflow.Future
}

func (r *SyntheticsGetCanaryResult) Get(ctx workflow.Context) (*synthetics.GetCanaryOutput, error) {
    var output synthetics.GetCanaryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SyntheticsGetCanaryRunsResult struct {
	Result workflow.Future
}

func (r *SyntheticsGetCanaryRunsResult) Get(ctx workflow.Context) (*synthetics.GetCanaryRunsOutput, error) {
    var output synthetics.GetCanaryRunsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SyntheticsListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *SyntheticsListTagsForResourceResult) Get(ctx workflow.Context) (*synthetics.ListTagsForResourceOutput, error) {
    var output synthetics.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SyntheticsStartCanaryResult struct {
	Result workflow.Future
}

func (r *SyntheticsStartCanaryResult) Get(ctx workflow.Context) (*synthetics.StartCanaryOutput, error) {
    var output synthetics.StartCanaryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SyntheticsStopCanaryResult struct {
	Result workflow.Future
}

func (r *SyntheticsStopCanaryResult) Get(ctx workflow.Context) (*synthetics.StopCanaryOutput, error) {
    var output synthetics.StopCanaryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SyntheticsTagResourceResult struct {
	Result workflow.Future
}

func (r *SyntheticsTagResourceResult) Get(ctx workflow.Context) (*synthetics.TagResourceOutput, error) {
    var output synthetics.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SyntheticsUntagResourceResult struct {
	Result workflow.Future
}

func (r *SyntheticsUntagResourceResult) Get(ctx workflow.Context) (*synthetics.UntagResourceOutput, error) {
    var output synthetics.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SyntheticsUpdateCanaryResult struct {
	Result workflow.Future
}

func (r *SyntheticsUpdateCanaryResult) Get(ctx workflow.Context) (*synthetics.UpdateCanaryOutput, error) {
    var output synthetics.UpdateCanaryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type SyntheticsStub struct {
    activities awsactivities.SyntheticsActivities
}

func NewSyntheticsStub() SyntheticsClient {
    return &SyntheticsStub{}
}
func (a *SyntheticsStub) CreateCanary(ctx workflow.Context, input *synthetics.CreateCanaryInput) (*synthetics.CreateCanaryOutput, error) {
    var output synthetics.CreateCanaryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCanary, input).Get(ctx, &output)
    return &output, err
}

func (a *SyntheticsStub) CreateCanaryAsync(ctx workflow.Context, input *synthetics.CreateCanaryInput) *SyntheticsCreateCanaryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateCanary, input)
    return &SyntheticsCreateCanaryResult{Result: future}
}
func (a *SyntheticsStub) DeleteCanary(ctx workflow.Context, input *synthetics.DeleteCanaryInput) (*synthetics.DeleteCanaryOutput, error) {
    var output synthetics.DeleteCanaryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCanary, input).Get(ctx, &output)
    return &output, err
}

func (a *SyntheticsStub) DeleteCanaryAsync(ctx workflow.Context, input *synthetics.DeleteCanaryInput) *SyntheticsDeleteCanaryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteCanary, input)
    return &SyntheticsDeleteCanaryResult{Result: future}
}
func (a *SyntheticsStub) DescribeCanaries(ctx workflow.Context, input *synthetics.DescribeCanariesInput) (*synthetics.DescribeCanariesOutput, error) {
    var output synthetics.DescribeCanariesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCanaries, input).Get(ctx, &output)
    return &output, err
}

func (a *SyntheticsStub) DescribeCanariesAsync(ctx workflow.Context, input *synthetics.DescribeCanariesInput) *SyntheticsDescribeCanariesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeCanaries, input)
    return &SyntheticsDescribeCanariesResult{Result: future}
}
func (a *SyntheticsStub) DescribeCanariesLastRun(ctx workflow.Context, input *synthetics.DescribeCanariesLastRunInput) (*synthetics.DescribeCanariesLastRunOutput, error) {
    var output synthetics.DescribeCanariesLastRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCanariesLastRun, input).Get(ctx, &output)
    return &output, err
}

func (a *SyntheticsStub) DescribeCanariesLastRunAsync(ctx workflow.Context, input *synthetics.DescribeCanariesLastRunInput) *SyntheticsDescribeCanariesLastRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeCanariesLastRun, input)
    return &SyntheticsDescribeCanariesLastRunResult{Result: future}
}
func (a *SyntheticsStub) DescribeRuntimeVersions(ctx workflow.Context, input *synthetics.DescribeRuntimeVersionsInput) (*synthetics.DescribeRuntimeVersionsOutput, error) {
    var output synthetics.DescribeRuntimeVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRuntimeVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *SyntheticsStub) DescribeRuntimeVersionsAsync(ctx workflow.Context, input *synthetics.DescribeRuntimeVersionsInput) *SyntheticsDescribeRuntimeVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeRuntimeVersions, input)
    return &SyntheticsDescribeRuntimeVersionsResult{Result: future}
}
func (a *SyntheticsStub) GetCanary(ctx workflow.Context, input *synthetics.GetCanaryInput) (*synthetics.GetCanaryOutput, error) {
    var output synthetics.GetCanaryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCanary, input).Get(ctx, &output)
    return &output, err
}

func (a *SyntheticsStub) GetCanaryAsync(ctx workflow.Context, input *synthetics.GetCanaryInput) *SyntheticsGetCanaryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetCanary, input)
    return &SyntheticsGetCanaryResult{Result: future}
}
func (a *SyntheticsStub) GetCanaryRuns(ctx workflow.Context, input *synthetics.GetCanaryRunsInput) (*synthetics.GetCanaryRunsOutput, error) {
    var output synthetics.GetCanaryRunsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCanaryRuns, input).Get(ctx, &output)
    return &output, err
}

func (a *SyntheticsStub) GetCanaryRunsAsync(ctx workflow.Context, input *synthetics.GetCanaryRunsInput) *SyntheticsGetCanaryRunsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetCanaryRuns, input)
    return &SyntheticsGetCanaryRunsResult{Result: future}
}
func (a *SyntheticsStub) ListTagsForResource(ctx workflow.Context, input *synthetics.ListTagsForResourceInput) (*synthetics.ListTagsForResourceOutput, error) {
    var output synthetics.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SyntheticsStub) ListTagsForResourceAsync(ctx workflow.Context, input *synthetics.ListTagsForResourceInput) *SyntheticsListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &SyntheticsListTagsForResourceResult{Result: future}
}
func (a *SyntheticsStub) StartCanary(ctx workflow.Context, input *synthetics.StartCanaryInput) (*synthetics.StartCanaryOutput, error) {
    var output synthetics.StartCanaryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartCanary, input).Get(ctx, &output)
    return &output, err
}

func (a *SyntheticsStub) StartCanaryAsync(ctx workflow.Context, input *synthetics.StartCanaryInput) *SyntheticsStartCanaryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartCanary, input)
    return &SyntheticsStartCanaryResult{Result: future}
}
func (a *SyntheticsStub) StopCanary(ctx workflow.Context, input *synthetics.StopCanaryInput) (*synthetics.StopCanaryOutput, error) {
    var output synthetics.StopCanaryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopCanary, input).Get(ctx, &output)
    return &output, err
}

func (a *SyntheticsStub) StopCanaryAsync(ctx workflow.Context, input *synthetics.StopCanaryInput) *SyntheticsStopCanaryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopCanary, input)
    return &SyntheticsStopCanaryResult{Result: future}
}
func (a *SyntheticsStub) TagResource(ctx workflow.Context, input *synthetics.TagResourceInput) (*synthetics.TagResourceOutput, error) {
    var output synthetics.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SyntheticsStub) TagResourceAsync(ctx workflow.Context, input *synthetics.TagResourceInput) *SyntheticsTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &SyntheticsTagResourceResult{Result: future}
}
func (a *SyntheticsStub) UntagResource(ctx workflow.Context, input *synthetics.UntagResourceInput) (*synthetics.UntagResourceOutput, error) {
    var output synthetics.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SyntheticsStub) UntagResourceAsync(ctx workflow.Context, input *synthetics.UntagResourceInput) *SyntheticsUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &SyntheticsUntagResourceResult{Result: future}
}
func (a *SyntheticsStub) UpdateCanary(ctx workflow.Context, input *synthetics.UpdateCanaryInput) (*synthetics.UpdateCanaryOutput, error) {
    var output synthetics.UpdateCanaryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateCanary, input).Get(ctx, &output)
    return &output, err
}

func (a *SyntheticsStub) UpdateCanaryAsync(ctx workflow.Context, input *synthetics.UpdateCanaryInput) *SyntheticsUpdateCanaryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateCanary, input)
    return &SyntheticsUpdateCanaryResult{Result: future}
}
