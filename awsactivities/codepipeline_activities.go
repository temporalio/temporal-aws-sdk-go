package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codepipeline"
	"github.com/aws/aws-sdk-go/service/codepipeline/codepipelineiface"
)

type CodePipelineActivities struct {
	client codepipelineiface.CodePipelineAPI
}

func NewCodePipelineActivities(session *session.Session, config ...*aws.Config) *CodePipelineActivities {
	client := codepipeline.New(session, config...)
	return &CodePipelineActivities{client: client}
}

func (a *CodePipelineActivities) AcknowledgeJob(input *codepipeline.AcknowledgeJobInput) (*codepipeline.AcknowledgeJobOutput, error) {
	return a.client.AcknowledgeJob(input)
}

func (a *CodePipelineActivities) AcknowledgeThirdPartyJob(input *codepipeline.AcknowledgeThirdPartyJobInput) (*codepipeline.AcknowledgeThirdPartyJobOutput, error) {
	return a.client.AcknowledgeThirdPartyJob(input)
}

func (a *CodePipelineActivities) CreateCustomActionType(input *codepipeline.CreateCustomActionTypeInput) (*codepipeline.CreateCustomActionTypeOutput, error) {
	return a.client.CreateCustomActionType(input)
}

func (a *CodePipelineActivities) CreatePipeline(input *codepipeline.CreatePipelineInput) (*codepipeline.CreatePipelineOutput, error) {
	return a.client.CreatePipeline(input)
}

func (a *CodePipelineActivities) DeleteCustomActionType(input *codepipeline.DeleteCustomActionTypeInput) (*codepipeline.DeleteCustomActionTypeOutput, error) {
	return a.client.DeleteCustomActionType(input)
}

func (a *CodePipelineActivities) DeletePipeline(input *codepipeline.DeletePipelineInput) (*codepipeline.DeletePipelineOutput, error) {
	return a.client.DeletePipeline(input)
}

func (a *CodePipelineActivities) DeleteWebhook(input *codepipeline.DeleteWebhookInput) (*codepipeline.DeleteWebhookOutput, error) {
	return a.client.DeleteWebhook(input)
}

func (a *CodePipelineActivities) DeregisterWebhookWithThirdParty(input *codepipeline.DeregisterWebhookWithThirdPartyInput) (*codepipeline.DeregisterWebhookWithThirdPartyOutput, error) {
	return a.client.DeregisterWebhookWithThirdParty(input)
}

func (a *CodePipelineActivities) DisableStageTransition(input *codepipeline.DisableStageTransitionInput) (*codepipeline.DisableStageTransitionOutput, error) {
	return a.client.DisableStageTransition(input)
}

func (a *CodePipelineActivities) EnableStageTransition(input *codepipeline.EnableStageTransitionInput) (*codepipeline.EnableStageTransitionOutput, error) {
	return a.client.EnableStageTransition(input)
}

func (a *CodePipelineActivities) GetJobDetails(input *codepipeline.GetJobDetailsInput) (*codepipeline.GetJobDetailsOutput, error) {
	return a.client.GetJobDetails(input)
}

func (a *CodePipelineActivities) GetPipeline(input *codepipeline.GetPipelineInput) (*codepipeline.GetPipelineOutput, error) {
	return a.client.GetPipeline(input)
}

func (a *CodePipelineActivities) GetPipelineExecution(input *codepipeline.GetPipelineExecutionInput) (*codepipeline.GetPipelineExecutionOutput, error) {
	return a.client.GetPipelineExecution(input)
}

func (a *CodePipelineActivities) GetPipelineState(input *codepipeline.GetPipelineStateInput) (*codepipeline.GetPipelineStateOutput, error) {
	return a.client.GetPipelineState(input)
}

func (a *CodePipelineActivities) GetThirdPartyJobDetails(input *codepipeline.GetThirdPartyJobDetailsInput) (*codepipeline.GetThirdPartyJobDetailsOutput, error) {
	return a.client.GetThirdPartyJobDetails(input)
}

func (a *CodePipelineActivities) ListActionExecutions(input *codepipeline.ListActionExecutionsInput) (*codepipeline.ListActionExecutionsOutput, error) {
	return a.client.ListActionExecutions(input)
}

func (a *CodePipelineActivities) ListActionTypes(input *codepipeline.ListActionTypesInput) (*codepipeline.ListActionTypesOutput, error) {
	return a.client.ListActionTypes(input)
}

func (a *CodePipelineActivities) ListPipelineExecutions(input *codepipeline.ListPipelineExecutionsInput) (*codepipeline.ListPipelineExecutionsOutput, error) {
	return a.client.ListPipelineExecutions(input)
}

func (a *CodePipelineActivities) ListPipelines(input *codepipeline.ListPipelinesInput) (*codepipeline.ListPipelinesOutput, error) {
	return a.client.ListPipelines(input)
}

func (a *CodePipelineActivities) ListTagsForResource(input *codepipeline.ListTagsForResourceInput) (*codepipeline.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResource(input)
}

func (a *CodePipelineActivities) ListWebhooks(input *codepipeline.ListWebhooksInput) (*codepipeline.ListWebhooksOutput, error) {
	return a.client.ListWebhooks(input)
}

func (a *CodePipelineActivities) PollForJobs(input *codepipeline.PollForJobsInput) (*codepipeline.PollForJobsOutput, error) {
	return a.client.PollForJobs(input)
}

func (a *CodePipelineActivities) PollForThirdPartyJobs(input *codepipeline.PollForThirdPartyJobsInput) (*codepipeline.PollForThirdPartyJobsOutput, error) {
	return a.client.PollForThirdPartyJobs(input)
}

func (a *CodePipelineActivities) PutActionRevision(input *codepipeline.PutActionRevisionInput) (*codepipeline.PutActionRevisionOutput, error) {
	return a.client.PutActionRevision(input)
}

func (a *CodePipelineActivities) PutApprovalResult(input *codepipeline.PutApprovalResultInput) (*codepipeline.PutApprovalResultOutput, error) {
	return a.client.PutApprovalResult(input)
}

func (a *CodePipelineActivities) PutJobFailureResult(input *codepipeline.PutJobFailureResultInput) (*codepipeline.PutJobFailureResultOutput, error) {
	return a.client.PutJobFailureResult(input)
}

func (a *CodePipelineActivities) PutJobSuccessResult(input *codepipeline.PutJobSuccessResultInput) (*codepipeline.PutJobSuccessResultOutput, error) {
	return a.client.PutJobSuccessResult(input)
}

func (a *CodePipelineActivities) PutThirdPartyJobFailureResult(input *codepipeline.PutThirdPartyJobFailureResultInput) (*codepipeline.PutThirdPartyJobFailureResultOutput, error) {
	return a.client.PutThirdPartyJobFailureResult(input)
}

func (a *CodePipelineActivities) PutThirdPartyJobSuccessResult(input *codepipeline.PutThirdPartyJobSuccessResultInput) (*codepipeline.PutThirdPartyJobSuccessResultOutput, error) {
	return a.client.PutThirdPartyJobSuccessResult(input)
}

func (a *CodePipelineActivities) PutWebhook(input *codepipeline.PutWebhookInput) (*codepipeline.PutWebhookOutput, error) {
	return a.client.PutWebhook(input)
}

func (a *CodePipelineActivities) RegisterWebhookWithThirdParty(input *codepipeline.RegisterWebhookWithThirdPartyInput) (*codepipeline.RegisterWebhookWithThirdPartyOutput, error) {
	return a.client.RegisterWebhookWithThirdParty(input)
}

func (a *CodePipelineActivities) RetryStageExecution(input *codepipeline.RetryStageExecutionInput) (*codepipeline.RetryStageExecutionOutput, error) {
	return a.client.RetryStageExecution(input)
}

func (a *CodePipelineActivities) StartPipelineExecution(input *codepipeline.StartPipelineExecutionInput) (*codepipeline.StartPipelineExecutionOutput, error) {
	return a.client.StartPipelineExecution(input)
}

func (a *CodePipelineActivities) StopPipelineExecution(input *codepipeline.StopPipelineExecutionInput) (*codepipeline.StopPipelineExecutionOutput, error) {
	return a.client.StopPipelineExecution(input)
}

func (a *CodePipelineActivities) TagResource(input *codepipeline.TagResourceInput) (*codepipeline.TagResourceOutput, error) {
	return a.client.TagResource(input)
}

func (a *CodePipelineActivities) UntagResource(input *codepipeline.UntagResourceInput) (*codepipeline.UntagResourceOutput, error) {
	return a.client.UntagResource(input)
}

func (a *CodePipelineActivities) UpdatePipeline(input *codepipeline.UpdatePipelineInput) (*codepipeline.UpdatePipelineOutput, error) {
	return a.client.UpdatePipeline(input)
}
