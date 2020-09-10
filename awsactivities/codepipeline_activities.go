package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codepipeline"
	"github.com/aws/aws-sdk-go/service/codepipeline/codepipelineiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type CodePipelineActivities struct {
	client codepipelineiface.CodePipelineAPI
}

func NewCodePipelineActivities(session *session.Session, config ...*aws.Config) *CodePipelineActivities {
	client := codepipeline.New(session, config...)
	return &CodePipelineActivities{client: client}
}

func (a *CodePipelineActivities) AcknowledgeJob(ctx context.Context, input *codepipeline.AcknowledgeJobInput) (*codepipeline.AcknowledgeJobOutput, error) {
	return a.client.AcknowledgeJobWithContext(ctx, input)
}

func (a *CodePipelineActivities) AcknowledgeThirdPartyJob(ctx context.Context, input *codepipeline.AcknowledgeThirdPartyJobInput) (*codepipeline.AcknowledgeThirdPartyJobOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.AcknowledgeThirdPartyJobWithContext(ctx, input)
}

func (a *CodePipelineActivities) CreateCustomActionType(ctx context.Context, input *codepipeline.CreateCustomActionTypeInput) (*codepipeline.CreateCustomActionTypeOutput, error) {
	return a.client.CreateCustomActionTypeWithContext(ctx, input)
}

func (a *CodePipelineActivities) CreatePipeline(ctx context.Context, input *codepipeline.CreatePipelineInput) (*codepipeline.CreatePipelineOutput, error) {
	return a.client.CreatePipelineWithContext(ctx, input)
}

func (a *CodePipelineActivities) DeleteCustomActionType(ctx context.Context, input *codepipeline.DeleteCustomActionTypeInput) (*codepipeline.DeleteCustomActionTypeOutput, error) {
	return a.client.DeleteCustomActionTypeWithContext(ctx, input)
}

func (a *CodePipelineActivities) DeletePipeline(ctx context.Context, input *codepipeline.DeletePipelineInput) (*codepipeline.DeletePipelineOutput, error) {
	return a.client.DeletePipelineWithContext(ctx, input)
}

func (a *CodePipelineActivities) DeleteWebhook(ctx context.Context, input *codepipeline.DeleteWebhookInput) (*codepipeline.DeleteWebhookOutput, error) {
	return a.client.DeleteWebhookWithContext(ctx, input)
}

func (a *CodePipelineActivities) DeregisterWebhookWithThirdParty(ctx context.Context, input *codepipeline.DeregisterWebhookWithThirdPartyInput) (*codepipeline.DeregisterWebhookWithThirdPartyOutput, error) {
	return a.client.DeregisterWebhookWithThirdPartyWithContext(ctx, input)
}

func (a *CodePipelineActivities) DisableStageTransition(ctx context.Context, input *codepipeline.DisableStageTransitionInput) (*codepipeline.DisableStageTransitionOutput, error) {
	return a.client.DisableStageTransitionWithContext(ctx, input)
}

func (a *CodePipelineActivities) EnableStageTransition(ctx context.Context, input *codepipeline.EnableStageTransitionInput) (*codepipeline.EnableStageTransitionOutput, error) {
	return a.client.EnableStageTransitionWithContext(ctx, input)
}

func (a *CodePipelineActivities) GetJobDetails(ctx context.Context, input *codepipeline.GetJobDetailsInput) (*codepipeline.GetJobDetailsOutput, error) {
	return a.client.GetJobDetailsWithContext(ctx, input)
}

func (a *CodePipelineActivities) GetPipeline(ctx context.Context, input *codepipeline.GetPipelineInput) (*codepipeline.GetPipelineOutput, error) {
	return a.client.GetPipelineWithContext(ctx, input)
}

func (a *CodePipelineActivities) GetPipelineExecution(ctx context.Context, input *codepipeline.GetPipelineExecutionInput) (*codepipeline.GetPipelineExecutionOutput, error) {
	return a.client.GetPipelineExecutionWithContext(ctx, input)
}

func (a *CodePipelineActivities) GetPipelineState(ctx context.Context, input *codepipeline.GetPipelineStateInput) (*codepipeline.GetPipelineStateOutput, error) {
	return a.client.GetPipelineStateWithContext(ctx, input)
}

func (a *CodePipelineActivities) GetThirdPartyJobDetails(ctx context.Context, input *codepipeline.GetThirdPartyJobDetailsInput) (*codepipeline.GetThirdPartyJobDetailsOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.GetThirdPartyJobDetailsWithContext(ctx, input)
}

func (a *CodePipelineActivities) ListActionExecutions(ctx context.Context, input *codepipeline.ListActionExecutionsInput) (*codepipeline.ListActionExecutionsOutput, error) {
	return a.client.ListActionExecutionsWithContext(ctx, input)
}

func (a *CodePipelineActivities) ListActionTypes(ctx context.Context, input *codepipeline.ListActionTypesInput) (*codepipeline.ListActionTypesOutput, error) {
	return a.client.ListActionTypesWithContext(ctx, input)
}

func (a *CodePipelineActivities) ListPipelineExecutions(ctx context.Context, input *codepipeline.ListPipelineExecutionsInput) (*codepipeline.ListPipelineExecutionsOutput, error) {
	return a.client.ListPipelineExecutionsWithContext(ctx, input)
}

func (a *CodePipelineActivities) ListPipelines(ctx context.Context, input *codepipeline.ListPipelinesInput) (*codepipeline.ListPipelinesOutput, error) {
	return a.client.ListPipelinesWithContext(ctx, input)
}

func (a *CodePipelineActivities) ListTagsForResource(ctx context.Context, input *codepipeline.ListTagsForResourceInput) (*codepipeline.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *CodePipelineActivities) ListWebhooks(ctx context.Context, input *codepipeline.ListWebhooksInput) (*codepipeline.ListWebhooksOutput, error) {
	return a.client.ListWebhooksWithContext(ctx, input)
}

func (a *CodePipelineActivities) PollForJobs(ctx context.Context, input *codepipeline.PollForJobsInput) (*codepipeline.PollForJobsOutput, error) {
	return a.client.PollForJobsWithContext(ctx, input)
}

func (a *CodePipelineActivities) PollForThirdPartyJobs(ctx context.Context, input *codepipeline.PollForThirdPartyJobsInput) (*codepipeline.PollForThirdPartyJobsOutput, error) {
	return a.client.PollForThirdPartyJobsWithContext(ctx, input)
}

func (a *CodePipelineActivities) PutActionRevision(ctx context.Context, input *codepipeline.PutActionRevisionInput) (*codepipeline.PutActionRevisionOutput, error) {
	return a.client.PutActionRevisionWithContext(ctx, input)
}

func (a *CodePipelineActivities) PutApprovalResult(ctx context.Context, input *codepipeline.PutApprovalResultInput) (*codepipeline.PutApprovalResultOutput, error) {
	return a.client.PutApprovalResultWithContext(ctx, input)
}

func (a *CodePipelineActivities) PutJobFailureResult(ctx context.Context, input *codepipeline.PutJobFailureResultInput) (*codepipeline.PutJobFailureResultOutput, error) {
	return a.client.PutJobFailureResultWithContext(ctx, input)
}

func (a *CodePipelineActivities) PutJobSuccessResult(ctx context.Context, input *codepipeline.PutJobSuccessResultInput) (*codepipeline.PutJobSuccessResultOutput, error) {
	return a.client.PutJobSuccessResultWithContext(ctx, input)
}

func (a *CodePipelineActivities) PutThirdPartyJobFailureResult(ctx context.Context, input *codepipeline.PutThirdPartyJobFailureResultInput) (*codepipeline.PutThirdPartyJobFailureResultOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.PutThirdPartyJobFailureResultWithContext(ctx, input)
}

func (a *CodePipelineActivities) PutThirdPartyJobSuccessResult(ctx context.Context, input *codepipeline.PutThirdPartyJobSuccessResultInput) (*codepipeline.PutThirdPartyJobSuccessResultOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.PutThirdPartyJobSuccessResultWithContext(ctx, input)
}

func (a *CodePipelineActivities) PutWebhook(ctx context.Context, input *codepipeline.PutWebhookInput) (*codepipeline.PutWebhookOutput, error) {
	return a.client.PutWebhookWithContext(ctx, input)
}

func (a *CodePipelineActivities) RegisterWebhookWithThirdParty(ctx context.Context, input *codepipeline.RegisterWebhookWithThirdPartyInput) (*codepipeline.RegisterWebhookWithThirdPartyOutput, error) {
	return a.client.RegisterWebhookWithThirdPartyWithContext(ctx, input)
}

func (a *CodePipelineActivities) RetryStageExecution(ctx context.Context, input *codepipeline.RetryStageExecutionInput) (*codepipeline.RetryStageExecutionOutput, error) {
	return a.client.RetryStageExecutionWithContext(ctx, input)
}

func (a *CodePipelineActivities) StartPipelineExecution(ctx context.Context, input *codepipeline.StartPipelineExecutionInput) (*codepipeline.StartPipelineExecutionOutput, error) {
	return a.client.StartPipelineExecutionWithContext(ctx, input)
}

func (a *CodePipelineActivities) StopPipelineExecution(ctx context.Context, input *codepipeline.StopPipelineExecutionInput) (*codepipeline.StopPipelineExecutionOutput, error) {
	return a.client.StopPipelineExecutionWithContext(ctx, input)
}

func (a *CodePipelineActivities) TagResource(ctx context.Context, input *codepipeline.TagResourceInput) (*codepipeline.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *CodePipelineActivities) UntagResource(ctx context.Context, input *codepipeline.UntagResourceInput) (*codepipeline.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *CodePipelineActivities) UpdatePipeline(ctx context.Context, input *codepipeline.UpdatePipelineInput) (*codepipeline.UpdatePipelineOutput, error) {
	return a.client.UpdatePipelineWithContext(ctx, input)
}
