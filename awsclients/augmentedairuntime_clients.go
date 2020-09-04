package awsclients

import (
	"github.com/aws/aws-sdk-go/service/augmentedairuntime"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type AugmentedAIRuntimeClient interface {
    DeleteHumanLoop(ctx workflow.Context, input *augmentedairuntime.DeleteHumanLoopInput) (*augmentedairuntime.DeleteHumanLoopOutput, error)
    DeleteHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.DeleteHumanLoopInput) *AugmentedairuntimeDeleteHumanLoopResult

    DescribeHumanLoop(ctx workflow.Context, input *augmentedairuntime.DescribeHumanLoopInput) (*augmentedairuntime.DescribeHumanLoopOutput, error)
    DescribeHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.DescribeHumanLoopInput) *AugmentedairuntimeDescribeHumanLoopResult

    ListHumanLoops(ctx workflow.Context, input *augmentedairuntime.ListHumanLoopsInput) (*augmentedairuntime.ListHumanLoopsOutput, error)
    ListHumanLoopsAsync(ctx workflow.Context, input *augmentedairuntime.ListHumanLoopsInput) *AugmentedairuntimeListHumanLoopsResult

    StartHumanLoop(ctx workflow.Context, input *augmentedairuntime.StartHumanLoopInput) (*augmentedairuntime.StartHumanLoopOutput, error)
    StartHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.StartHumanLoopInput) *AugmentedairuntimeStartHumanLoopResult

    StopHumanLoop(ctx workflow.Context, input *augmentedairuntime.StopHumanLoopInput) (*augmentedairuntime.StopHumanLoopOutput, error)
    StopHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.StopHumanLoopInput) *AugmentedairuntimeStopHumanLoopResult
}
type AugmentedairuntimeDeleteHumanLoopResult struct {
	Result workflow.Future
}

func (r *AugmentedairuntimeDeleteHumanLoopResult) Get(ctx workflow.Context) (*augmentedairuntime.DeleteHumanLoopOutput, error) {
    var output augmentedairuntime.DeleteHumanLoopOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AugmentedairuntimeDescribeHumanLoopResult struct {
	Result workflow.Future
}

func (r *AugmentedairuntimeDescribeHumanLoopResult) Get(ctx workflow.Context) (*augmentedairuntime.DescribeHumanLoopOutput, error) {
    var output augmentedairuntime.DescribeHumanLoopOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AugmentedairuntimeListHumanLoopsResult struct {
	Result workflow.Future
}

func (r *AugmentedairuntimeListHumanLoopsResult) Get(ctx workflow.Context) (*augmentedairuntime.ListHumanLoopsOutput, error) {
    var output augmentedairuntime.ListHumanLoopsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AugmentedairuntimeStartHumanLoopResult struct {
	Result workflow.Future
}

func (r *AugmentedairuntimeStartHumanLoopResult) Get(ctx workflow.Context) (*augmentedairuntime.StartHumanLoopOutput, error) {
    var output augmentedairuntime.StartHumanLoopOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AugmentedairuntimeStopHumanLoopResult struct {
	Result workflow.Future
}

func (r *AugmentedairuntimeStopHumanLoopResult) Get(ctx workflow.Context) (*augmentedairuntime.StopHumanLoopOutput, error) {
    var output augmentedairuntime.StopHumanLoopOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type AugmentedAIRuntimeStub struct {
    activities awsactivities.AugmentedAIRuntimeActivities
}

func NewAugmentedAIRuntimeStub() AugmentedAIRuntimeClient {
    return &AugmentedAIRuntimeStub{}
}
func (a *AugmentedAIRuntimeStub) DeleteHumanLoop(ctx workflow.Context, input *augmentedairuntime.DeleteHumanLoopInput) (*augmentedairuntime.DeleteHumanLoopOutput, error) {
    var output augmentedairuntime.DeleteHumanLoopOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteHumanLoop, input).Get(ctx, &output)
    return &output, err
}

func (a *AugmentedAIRuntimeStub) DeleteHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.DeleteHumanLoopInput) *AugmentedairuntimeDeleteHumanLoopResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteHumanLoop, input)
    return &AugmentedairuntimeDeleteHumanLoopResult{Result: future}
}
func (a *AugmentedAIRuntimeStub) DescribeHumanLoop(ctx workflow.Context, input *augmentedairuntime.DescribeHumanLoopInput) (*augmentedairuntime.DescribeHumanLoopOutput, error) {
    var output augmentedairuntime.DescribeHumanLoopOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeHumanLoop, input).Get(ctx, &output)
    return &output, err
}

func (a *AugmentedAIRuntimeStub) DescribeHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.DescribeHumanLoopInput) *AugmentedairuntimeDescribeHumanLoopResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeHumanLoop, input)
    return &AugmentedairuntimeDescribeHumanLoopResult{Result: future}
}
func (a *AugmentedAIRuntimeStub) ListHumanLoops(ctx workflow.Context, input *augmentedairuntime.ListHumanLoopsInput) (*augmentedairuntime.ListHumanLoopsOutput, error) {
    var output augmentedairuntime.ListHumanLoopsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListHumanLoops, input).Get(ctx, &output)
    return &output, err
}

func (a *AugmentedAIRuntimeStub) ListHumanLoopsAsync(ctx workflow.Context, input *augmentedairuntime.ListHumanLoopsInput) *AugmentedairuntimeListHumanLoopsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListHumanLoops, input)
    return &AugmentedairuntimeListHumanLoopsResult{Result: future}
}
func (a *AugmentedAIRuntimeStub) StartHumanLoop(ctx workflow.Context, input *augmentedairuntime.StartHumanLoopInput) (*augmentedairuntime.StartHumanLoopOutput, error) {
    var output augmentedairuntime.StartHumanLoopOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartHumanLoop, input).Get(ctx, &output)
    return &output, err
}

func (a *AugmentedAIRuntimeStub) StartHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.StartHumanLoopInput) *AugmentedairuntimeStartHumanLoopResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartHumanLoop, input)
    return &AugmentedairuntimeStartHumanLoopResult{Result: future}
}
func (a *AugmentedAIRuntimeStub) StopHumanLoop(ctx workflow.Context, input *augmentedairuntime.StopHumanLoopInput) (*augmentedairuntime.StopHumanLoopOutput, error) {
    var output augmentedairuntime.StopHumanLoopOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopHumanLoop, input).Get(ctx, &output)
    return &output, err
}

func (a *AugmentedAIRuntimeStub) StopHumanLoopAsync(ctx workflow.Context, input *augmentedairuntime.StopHumanLoopInput) *AugmentedairuntimeStopHumanLoopResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopHumanLoop, input)
    return &AugmentedairuntimeStopHumanLoopResult{Result: future}
}
