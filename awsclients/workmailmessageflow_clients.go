package awsclients

import (
	"github.com/aws/aws-sdk-go/service/workmailmessageflow"
	"go.temporal.io/sdk/workflow"
)

type WorkMailMessageFlowClient interface {
       GetRawMessageContent(ctx workflow.Context, input *workmailmessageflow.GetRawMessageContentInput) (*workmailmessageflow.GetRawMessageContentOutput, error)
       GetRawMessageContentAsync(ctx workflow.Context, input *workmailmessageflow.GetRawMessageContentInput) *WorkmailmessageflowGetRawMessageContentResult
}

type WorkMailMessageFlowStub struct {
}

func NewWorkMailMessageFlowStub() WorkMailMessageFlowClient {
    return &WorkMailMessageFlowStub{}
}

type WorkmailmessageflowGetRawMessageContentResult struct {
	Result workflow.Future
}

func (r *WorkmailmessageflowGetRawMessageContentResult) Get(ctx workflow.Context) (*workmailmessageflow.GetRawMessageContentOutput, error) {
    var output workmailmessageflow.GetRawMessageContentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *WorkMailMessageFlowStub) GetRawMessageContent(ctx workflow.Context, input *workmailmessageflow.GetRawMessageContentInput) (*workmailmessageflow.GetRawMessageContentOutput, error) {
    var output workmailmessageflow.GetRawMessageContentOutput
    err := workflow.ExecuteActivity(ctx, "WorkMailMessageFlow.GetRawMessageContent", input).Get(ctx, &output)
    return &output, err
}

func (a *WorkMailMessageFlowStub) GetRawMessageContentAsync(ctx workflow.Context, input *workmailmessageflow.GetRawMessageContentInput) *WorkmailmessageflowGetRawMessageContentResult {
    future := workflow.ExecuteActivity(ctx, "WorkMailMessageFlow.GetRawMessageContent", input)
    return &WorkmailmessageflowGetRawMessageContentResult{Result: future}
}
