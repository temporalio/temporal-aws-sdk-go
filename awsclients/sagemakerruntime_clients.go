package awsclients

import (
	"github.com/aws/aws-sdk-go/service/sagemakerruntime"
	"go.temporal.io/sdk/workflow"
)

type SageMakerRuntimeClient interface {
       InvokeEndpoint(ctx workflow.Context, input *sagemakerruntime.InvokeEndpointInput) (*sagemakerruntime.InvokeEndpointOutput, error)
       InvokeEndpointAsync(ctx workflow.Context, input *sagemakerruntime.InvokeEndpointInput) *SagemakerruntimeInvokeEndpointResult
}

type SageMakerRuntimeStub struct {
}

func NewSageMakerRuntimeStub() SageMakerRuntimeClient {
    return &SageMakerRuntimeStub{}
}

type SagemakerruntimeInvokeEndpointResult struct {
	Result workflow.Future
}

func (r *SagemakerruntimeInvokeEndpointResult) Get(ctx workflow.Context) (*sagemakerruntime.InvokeEndpointOutput, error) {
    var output sagemakerruntime.InvokeEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *SageMakerRuntimeStub) InvokeEndpoint(ctx workflow.Context, input *sagemakerruntime.InvokeEndpointInput) (*sagemakerruntime.InvokeEndpointOutput, error) {
    var output sagemakerruntime.InvokeEndpointOutput
    err := workflow.ExecuteActivity(ctx, "SageMakerRuntime.InvokeEndpoint", input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerRuntimeStub) InvokeEndpointAsync(ctx workflow.Context, input *sagemakerruntime.InvokeEndpointInput) *SagemakerruntimeInvokeEndpointResult {
    future := workflow.ExecuteActivity(ctx, "SageMakerRuntime.InvokeEndpoint", input)
    return &SagemakerruntimeInvokeEndpointResult{Result: future}
}
