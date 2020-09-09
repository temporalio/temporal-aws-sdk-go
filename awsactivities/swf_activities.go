
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/swf"
	"github.com/aws/aws-sdk-go/service/swf/swfiface"
)

type SWFActivities struct {
    client swfiface.SWFAPI
}

func NewSWFActivities(session *session.Session, config ...*aws.Config) *SWFActivities {
    client := swf.New(session, config...)
    return &SWFActivities{client: client}
}

func (a *SWFActivities) CountClosedWorkflowExecutions(input *swf.CountClosedWorkflowExecutionsInput) (*swf.WorkflowExecutionCount, error) {
    return a.client.CountClosedWorkflowExecutions(input)
}

func (a *SWFActivities) CountOpenWorkflowExecutions(input *swf.CountOpenWorkflowExecutionsInput) (*swf.WorkflowExecutionCount, error) {
    return a.client.CountOpenWorkflowExecutions(input)
}

func (a *SWFActivities) CountPendingActivityTasks(input *swf.CountPendingActivityTasksInput) (*swf.PendingTaskCount, error) {
    return a.client.CountPendingActivityTasks(input)
}

func (a *SWFActivities) CountPendingDecisionTasks(input *swf.CountPendingDecisionTasksInput) (*swf.PendingTaskCount, error) {
    return a.client.CountPendingDecisionTasks(input)
}

func (a *SWFActivities) DeprecateActivityType(input *swf.DeprecateActivityTypeInput) (*swf.DeprecateActivityTypeOutput, error) {
    return a.client.DeprecateActivityType(input)
}

func (a *SWFActivities) DeprecateDomain(input *swf.DeprecateDomainInput) (*swf.DeprecateDomainOutput, error) {
    return a.client.DeprecateDomain(input)
}

func (a *SWFActivities) DeprecateWorkflowType(input *swf.DeprecateWorkflowTypeInput) (*swf.DeprecateWorkflowTypeOutput, error) {
    return a.client.DeprecateWorkflowType(input)
}

func (a *SWFActivities) DescribeActivityType(input *swf.DescribeActivityTypeInput) (*swf.DescribeActivityTypeOutput, error) {
    return a.client.DescribeActivityType(input)
}

func (a *SWFActivities) DescribeDomain(input *swf.DescribeDomainInput) (*swf.DescribeDomainOutput, error) {
    return a.client.DescribeDomain(input)
}

func (a *SWFActivities) DescribeWorkflowExecution(input *swf.DescribeWorkflowExecutionInput) (*swf.DescribeWorkflowExecutionOutput, error) {
    return a.client.DescribeWorkflowExecution(input)
}

func (a *SWFActivities) DescribeWorkflowType(input *swf.DescribeWorkflowTypeInput) (*swf.DescribeWorkflowTypeOutput, error) {
    return a.client.DescribeWorkflowType(input)
}

func (a *SWFActivities) GetWorkflowExecutionHistory(input *swf.GetWorkflowExecutionHistoryInput) (*swf.GetWorkflowExecutionHistoryOutput, error) {
    return a.client.GetWorkflowExecutionHistory(input)
}

func (a *SWFActivities) ListActivityTypes(input *swf.ListActivityTypesInput) (*swf.ListActivityTypesOutput, error) {
    return a.client.ListActivityTypes(input)
}

func (a *SWFActivities) ListClosedWorkflowExecutions(input *swf.ListClosedWorkflowExecutionsInput) (*swf.WorkflowExecutionInfos, error) {
    return a.client.ListClosedWorkflowExecutions(input)
}

func (a *SWFActivities) ListDomains(input *swf.ListDomainsInput) (*swf.ListDomainsOutput, error) {
    return a.client.ListDomains(input)
}

func (a *SWFActivities) ListOpenWorkflowExecutions(input *swf.ListOpenWorkflowExecutionsInput) (*swf.WorkflowExecutionInfos, error) {
    return a.client.ListOpenWorkflowExecutions(input)
}

func (a *SWFActivities) ListTagsForResource(input *swf.ListTagsForResourceInput) (*swf.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *SWFActivities) ListWorkflowTypes(input *swf.ListWorkflowTypesInput) (*swf.ListWorkflowTypesOutput, error) {
    return a.client.ListWorkflowTypes(input)
}

func (a *SWFActivities) PollForActivityTask(input *swf.PollForActivityTaskInput) (*swf.PollForActivityTaskOutput, error) {
    return a.client.PollForActivityTask(input)
}

func (a *SWFActivities) PollForDecisionTask(input *swf.PollForDecisionTaskInput) (*swf.PollForDecisionTaskOutput, error) {
    return a.client.PollForDecisionTask(input)
}

func (a *SWFActivities) RecordActivityTaskHeartbeat(input *swf.RecordActivityTaskHeartbeatInput) (*swf.RecordActivityTaskHeartbeatOutput, error) {
    return a.client.RecordActivityTaskHeartbeat(input)
}

func (a *SWFActivities) RegisterActivityType(input *swf.RegisterActivityTypeInput) (*swf.RegisterActivityTypeOutput, error) {
    return a.client.RegisterActivityType(input)
}

func (a *SWFActivities) RegisterDomain(input *swf.RegisterDomainInput) (*swf.RegisterDomainOutput, error) {
    return a.client.RegisterDomain(input)
}

func (a *SWFActivities) RegisterWorkflowType(input *swf.RegisterWorkflowTypeInput) (*swf.RegisterWorkflowTypeOutput, error) {
    return a.client.RegisterWorkflowType(input)
}

func (a *SWFActivities) RequestCancelWorkflowExecution(input *swf.RequestCancelWorkflowExecutionInput) (*swf.RequestCancelWorkflowExecutionOutput, error) {
    return a.client.RequestCancelWorkflowExecution(input)
}

func (a *SWFActivities) RespondActivityTaskCanceled(input *swf.RespondActivityTaskCanceledInput) (*swf.RespondActivityTaskCanceledOutput, error) {
    return a.client.RespondActivityTaskCanceled(input)
}

func (a *SWFActivities) RespondActivityTaskCompleted(input *swf.RespondActivityTaskCompletedInput) (*swf.RespondActivityTaskCompletedOutput, error) {
    return a.client.RespondActivityTaskCompleted(input)
}

func (a *SWFActivities) RespondActivityTaskFailed(input *swf.RespondActivityTaskFailedInput) (*swf.RespondActivityTaskFailedOutput, error) {
    return a.client.RespondActivityTaskFailed(input)
}

func (a *SWFActivities) RespondDecisionTaskCompleted(input *swf.RespondDecisionTaskCompletedInput) (*swf.RespondDecisionTaskCompletedOutput, error) {
    return a.client.RespondDecisionTaskCompleted(input)
}

func (a *SWFActivities) SignalWorkflowExecution(input *swf.SignalWorkflowExecutionInput) (*swf.SignalWorkflowExecutionOutput, error) {
    return a.client.SignalWorkflowExecution(input)
}

func (a *SWFActivities) StartWorkflowExecution(input *swf.StartWorkflowExecutionInput) (*swf.StartWorkflowExecutionOutput, error) {
    return a.client.StartWorkflowExecution(input)
}

func (a *SWFActivities) TagResource(input *swf.TagResourceInput) (*swf.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *SWFActivities) TerminateWorkflowExecution(input *swf.TerminateWorkflowExecutionInput) (*swf.TerminateWorkflowExecutionOutput, error) {
    return a.client.TerminateWorkflowExecution(input)
}

func (a *SWFActivities) UndeprecateActivityType(input *swf.UndeprecateActivityTypeInput) (*swf.UndeprecateActivityTypeOutput, error) {
    return a.client.UndeprecateActivityType(input)
}

func (a *SWFActivities) UndeprecateDomain(input *swf.UndeprecateDomainInput) (*swf.UndeprecateDomainOutput, error) {
    return a.client.UndeprecateDomain(input)
}

func (a *SWFActivities) UndeprecateWorkflowType(input *swf.UndeprecateWorkflowTypeInput) (*swf.UndeprecateWorkflowTypeOutput, error) {
    return a.client.UndeprecateWorkflowType(input)
}

func (a *SWFActivities) UntagResource(input *swf.UntagResourceInput) (*swf.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}
