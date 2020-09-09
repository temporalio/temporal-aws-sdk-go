package awsclients

import (
	"github.com/aws/aws-sdk-go/service/workmailmessageflow"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type WorkMailMessageFlowClient interface {
	GetRawMessageContent(ctx workflow.Context, input *workmailmessageflow.GetRawMessageContentInput) (*workmailmessageflow.GetRawMessageContentOutput, error)
	GetRawMessageContentAsync(ctx workflow.Context, input *workmailmessageflow.GetRawMessageContentInput) *WorkmailmessageflowGetRawMessageContentResult
}

type WorkmailmessageflowGetRawMessageContentResult struct {
	Result workflow.Future
}

func (r *WorkmailmessageflowGetRawMessageContentResult) Get(ctx workflow.Context) (*workmailmessageflow.GetRawMessageContentOutput, error) {
	var output workmailmessageflow.GetRawMessageContentOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type WorkMailMessageFlowStub struct {
	activities awsactivities.WorkMailMessageFlowActivities
}

func NewWorkMailMessageFlowStub() WorkMailMessageFlowClient {
	return &WorkMailMessageFlowStub{}
}

func (a *WorkMailMessageFlowStub) GetRawMessageContent(ctx workflow.Context, input *workmailmessageflow.GetRawMessageContentInput) (*workmailmessageflow.GetRawMessageContentOutput, error) {
	var output workmailmessageflow.GetRawMessageContentOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRawMessageContent, input).Get(ctx, &output)
	return &output, err
}

func (a *WorkMailMessageFlowStub) GetRawMessageContentAsync(ctx workflow.Context, input *workmailmessageflow.GetRawMessageContentInput) *WorkmailmessageflowGetRawMessageContentResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRawMessageContent, input)
	return &WorkmailmessageflowGetRawMessageContentResult{Result: future}
}
