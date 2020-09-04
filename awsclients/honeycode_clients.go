package awsclients

import (
	"github.com/aws/aws-sdk-go/service/honeycode"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type HoneycodeClient interface {
    GetScreenData(ctx workflow.Context, input *honeycode.GetScreenDataInput) (*honeycode.GetScreenDataOutput, error)
    GetScreenDataAsync(ctx workflow.Context, input *honeycode.GetScreenDataInput) *HoneycodeGetScreenDataResult

    InvokeScreenAutomation(ctx workflow.Context, input *honeycode.InvokeScreenAutomationInput) (*honeycode.InvokeScreenAutomationOutput, error)
    InvokeScreenAutomationAsync(ctx workflow.Context, input *honeycode.InvokeScreenAutomationInput) *HoneycodeInvokeScreenAutomationResult
}

type HoneycodeGetScreenDataResult struct {
	Result workflow.Future
}

func (r *HoneycodeGetScreenDataResult) Get(ctx workflow.Context) (*honeycode.GetScreenDataOutput, error) {
    var output honeycode.GetScreenDataOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type HoneycodeInvokeScreenAutomationResult struct {
	Result workflow.Future
}

func (r *HoneycodeInvokeScreenAutomationResult) Get(ctx workflow.Context) (*honeycode.InvokeScreenAutomationOutput, error) {
    var output honeycode.InvokeScreenAutomationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type HoneycodeStub struct {
    activities awsactivities.HoneycodeActivities
}

func NewHoneycodeStub() HoneycodeClient {
    return &HoneycodeStub{}
}

func (a *HoneycodeStub) GetScreenData(ctx workflow.Context, input *honeycode.GetScreenDataInput) (*honeycode.GetScreenDataOutput, error) {
    var output honeycode.GetScreenDataOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetScreenData, input).Get(ctx, &output)
    return &output, err
}

func (a *HoneycodeStub) GetScreenDataAsync(ctx workflow.Context, input *honeycode.GetScreenDataInput) *HoneycodeGetScreenDataResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetScreenData, input)
    return &HoneycodeGetScreenDataResult{Result: future}
}

func (a *HoneycodeStub) InvokeScreenAutomation(ctx workflow.Context, input *honeycode.InvokeScreenAutomationInput) (*honeycode.InvokeScreenAutomationOutput, error) {
    var output honeycode.InvokeScreenAutomationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.InvokeScreenAutomation, input).Get(ctx, &output)
    return &output, err
}

func (a *HoneycodeStub) InvokeScreenAutomationAsync(ctx workflow.Context, input *honeycode.InvokeScreenAutomationInput) *HoneycodeInvokeScreenAutomationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.InvokeScreenAutomation, input)
    return &HoneycodeInvokeScreenAutomationResult{Result: future}
}
