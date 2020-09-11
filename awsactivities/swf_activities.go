package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/swf"
	"github.com/aws/aws-sdk-go/service/swf/swfiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type SWFActivities struct {
	client swfiface.SWFAPI
}

func NewSWFActivities(session *session.Session, config ...*aws.Config) *SWFActivities {
	client := swf.New(session, config...)
	return &SWFActivities{client: client}
}

func (a *SWFActivities) CountClosedWorkflowExecutions(ctx context.Context, input *swf.CountClosedWorkflowExecutionsInput) (*swf.WorkflowExecutionCount, error) {
	return a.client.CountClosedWorkflowExecutionsWithContext(ctx, input)
}

func (a *SWFActivities) CountOpenWorkflowExecutions(ctx context.Context, input *swf.CountOpenWorkflowExecutionsInput) (*swf.WorkflowExecutionCount, error) {
	return a.client.CountOpenWorkflowExecutionsWithContext(ctx, input)
}

func (a *SWFActivities) CountPendingActivityTasks(ctx context.Context, input *swf.CountPendingActivityTasksInput) (*swf.PendingTaskCount, error) {
	return a.client.CountPendingActivityTasksWithContext(ctx, input)
}

func (a *SWFActivities) CountPendingDecisionTasks(ctx context.Context, input *swf.CountPendingDecisionTasksInput) (*swf.PendingTaskCount, error) {
	return a.client.CountPendingDecisionTasksWithContext(ctx, input)
}

func (a *SWFActivities) DeprecateActivityType(ctx context.Context, input *swf.DeprecateActivityTypeInput) (*swf.DeprecateActivityTypeOutput, error) {
	return a.client.DeprecateActivityTypeWithContext(ctx, input)
}

func (a *SWFActivities) DeprecateDomain(ctx context.Context, input *swf.DeprecateDomainInput) (*swf.DeprecateDomainOutput, error) {
	return a.client.DeprecateDomainWithContext(ctx, input)
}

func (a *SWFActivities) DeprecateWorkflowType(ctx context.Context, input *swf.DeprecateWorkflowTypeInput) (*swf.DeprecateWorkflowTypeOutput, error) {
	return a.client.DeprecateWorkflowTypeWithContext(ctx, input)
}

func (a *SWFActivities) DescribeActivityType(ctx context.Context, input *swf.DescribeActivityTypeInput) (*swf.DescribeActivityTypeOutput, error) {
	return a.client.DescribeActivityTypeWithContext(ctx, input)
}

func (a *SWFActivities) DescribeDomain(ctx context.Context, input *swf.DescribeDomainInput) (*swf.DescribeDomainOutput, error) {
	return a.client.DescribeDomainWithContext(ctx, input)
}

func (a *SWFActivities) DescribeWorkflowExecution(ctx context.Context, input *swf.DescribeWorkflowExecutionInput) (*swf.DescribeWorkflowExecutionOutput, error) {
	return a.client.DescribeWorkflowExecutionWithContext(ctx, input)
}

func (a *SWFActivities) DescribeWorkflowType(ctx context.Context, input *swf.DescribeWorkflowTypeInput) (*swf.DescribeWorkflowTypeOutput, error) {
	return a.client.DescribeWorkflowTypeWithContext(ctx, input)
}

func (a *SWFActivities) GetWorkflowExecutionHistory(ctx context.Context, input *swf.GetWorkflowExecutionHistoryInput) (*swf.GetWorkflowExecutionHistoryOutput, error) {
	return a.client.GetWorkflowExecutionHistoryWithContext(ctx, input)
}

func (a *SWFActivities) ListActivityTypes(ctx context.Context, input *swf.ListActivityTypesInput) (*swf.ListActivityTypesOutput, error) {
	return a.client.ListActivityTypesWithContext(ctx, input)
}

func (a *SWFActivities) ListClosedWorkflowExecutions(ctx context.Context, input *swf.ListClosedWorkflowExecutionsInput) (*swf.WorkflowExecutionInfos, error) {
	return a.client.ListClosedWorkflowExecutionsWithContext(ctx, input)
}

func (a *SWFActivities) ListDomains(ctx context.Context, input *swf.ListDomainsInput) (*swf.ListDomainsOutput, error) {
	return a.client.ListDomainsWithContext(ctx, input)
}

func (a *SWFActivities) ListOpenWorkflowExecutions(ctx context.Context, input *swf.ListOpenWorkflowExecutionsInput) (*swf.WorkflowExecutionInfos, error) {
	return a.client.ListOpenWorkflowExecutionsWithContext(ctx, input)
}

func (a *SWFActivities) ListTagsForResource(ctx context.Context, input *swf.ListTagsForResourceInput) (*swf.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *SWFActivities) ListWorkflowTypes(ctx context.Context, input *swf.ListWorkflowTypesInput) (*swf.ListWorkflowTypesOutput, error) {
	return a.client.ListWorkflowTypesWithContext(ctx, input)
}

func (a *SWFActivities) PollForActivityTask(ctx context.Context, input *swf.PollForActivityTaskInput) (*swf.PollForActivityTaskOutput, error) {
	return a.client.PollForActivityTaskWithContext(ctx, input)
}

func (a *SWFActivities) PollForDecisionTask(ctx context.Context, input *swf.PollForDecisionTaskInput) (*swf.PollForDecisionTaskOutput, error) {
	return a.client.PollForDecisionTaskWithContext(ctx, input)
}

func (a *SWFActivities) RecordActivityTaskHeartbeat(ctx context.Context, input *swf.RecordActivityTaskHeartbeatInput) (*swf.RecordActivityTaskHeartbeatOutput, error) {
	return a.client.RecordActivityTaskHeartbeatWithContext(ctx, input)
}

func (a *SWFActivities) RegisterActivityType(ctx context.Context, input *swf.RegisterActivityTypeInput) (*swf.RegisterActivityTypeOutput, error) {
	return a.client.RegisterActivityTypeWithContext(ctx, input)
}

func (a *SWFActivities) RegisterDomain(ctx context.Context, input *swf.RegisterDomainInput) (*swf.RegisterDomainOutput, error) {
	return a.client.RegisterDomainWithContext(ctx, input)
}

func (a *SWFActivities) RegisterWorkflowType(ctx context.Context, input *swf.RegisterWorkflowTypeInput) (*swf.RegisterWorkflowTypeOutput, error) {
	return a.client.RegisterWorkflowTypeWithContext(ctx, input)
}

func (a *SWFActivities) RequestCancelWorkflowExecution(ctx context.Context, input *swf.RequestCancelWorkflowExecutionInput) (*swf.RequestCancelWorkflowExecutionOutput, error) {
	return a.client.RequestCancelWorkflowExecutionWithContext(ctx, input)
}

func (a *SWFActivities) RespondActivityTaskCanceled(ctx context.Context, input *swf.RespondActivityTaskCanceledInput) (*swf.RespondActivityTaskCanceledOutput, error) {
	return a.client.RespondActivityTaskCanceledWithContext(ctx, input)
}

func (a *SWFActivities) RespondActivityTaskCompleted(ctx context.Context, input *swf.RespondActivityTaskCompletedInput) (*swf.RespondActivityTaskCompletedOutput, error) {
	return a.client.RespondActivityTaskCompletedWithContext(ctx, input)
}

func (a *SWFActivities) RespondActivityTaskFailed(ctx context.Context, input *swf.RespondActivityTaskFailedInput) (*swf.RespondActivityTaskFailedOutput, error) {
	return a.client.RespondActivityTaskFailedWithContext(ctx, input)
}

func (a *SWFActivities) RespondDecisionTaskCompleted(ctx context.Context, input *swf.RespondDecisionTaskCompletedInput) (*swf.RespondDecisionTaskCompletedOutput, error) {
	return a.client.RespondDecisionTaskCompletedWithContext(ctx, input)
}

func (a *SWFActivities) SignalWorkflowExecution(ctx context.Context, input *swf.SignalWorkflowExecutionInput) (*swf.SignalWorkflowExecutionOutput, error) {
	return a.client.SignalWorkflowExecutionWithContext(ctx, input)
}

func (a *SWFActivities) StartWorkflowExecution(ctx context.Context, input *swf.StartWorkflowExecutionInput) (*swf.StartWorkflowExecutionOutput, error) {
	return a.client.StartWorkflowExecutionWithContext(ctx, input)
}

func (a *SWFActivities) TagResource(ctx context.Context, input *swf.TagResourceInput) (*swf.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *SWFActivities) TerminateWorkflowExecution(ctx context.Context, input *swf.TerminateWorkflowExecutionInput) (*swf.TerminateWorkflowExecutionOutput, error) {
	return a.client.TerminateWorkflowExecutionWithContext(ctx, input)
}

func (a *SWFActivities) UndeprecateActivityType(ctx context.Context, input *swf.UndeprecateActivityTypeInput) (*swf.UndeprecateActivityTypeOutput, error) {
	return a.client.UndeprecateActivityTypeWithContext(ctx, input)
}

func (a *SWFActivities) UndeprecateDomain(ctx context.Context, input *swf.UndeprecateDomainInput) (*swf.UndeprecateDomainOutput, error) {
	return a.client.UndeprecateDomainWithContext(ctx, input)
}

func (a *SWFActivities) UndeprecateWorkflowType(ctx context.Context, input *swf.UndeprecateWorkflowTypeInput) (*swf.UndeprecateWorkflowTypeOutput, error) {
	return a.client.UndeprecateWorkflowTypeWithContext(ctx, input)
}

func (a *SWFActivities) UntagResource(ctx context.Context, input *swf.UntagResourceInput) (*swf.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}
