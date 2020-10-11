// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package swfstub

import (
	"github.com/aws/aws-sdk-go/service/swf"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

type Client interface {
	CountClosedWorkflowExecutions(ctx workflow.Context, input *swf.CountClosedWorkflowExecutionsInput) (*swf.WorkflowExecutionCount, error)
	CountClosedWorkflowExecutionsAsync(ctx workflow.Context, input *swf.CountClosedWorkflowExecutionsInput) *SWFCountClosedWorkflowExecutionsFuture

	CountOpenWorkflowExecutions(ctx workflow.Context, input *swf.CountOpenWorkflowExecutionsInput) (*swf.WorkflowExecutionCount, error)
	CountOpenWorkflowExecutionsAsync(ctx workflow.Context, input *swf.CountOpenWorkflowExecutionsInput) *SWFCountOpenWorkflowExecutionsFuture

	CountPendingActivityTasks(ctx workflow.Context, input *swf.CountPendingActivityTasksInput) (*swf.PendingTaskCount, error)
	CountPendingActivityTasksAsync(ctx workflow.Context, input *swf.CountPendingActivityTasksInput) *SWFCountPendingActivityTasksFuture

	CountPendingDecisionTasks(ctx workflow.Context, input *swf.CountPendingDecisionTasksInput) (*swf.PendingTaskCount, error)
	CountPendingDecisionTasksAsync(ctx workflow.Context, input *swf.CountPendingDecisionTasksInput) *SWFCountPendingDecisionTasksFuture

	DeprecateActivityType(ctx workflow.Context, input *swf.DeprecateActivityTypeInput) (*swf.DeprecateActivityTypeOutput, error)
	DeprecateActivityTypeAsync(ctx workflow.Context, input *swf.DeprecateActivityTypeInput) *SWFDeprecateActivityTypeFuture

	DeprecateDomain(ctx workflow.Context, input *swf.DeprecateDomainInput) (*swf.DeprecateDomainOutput, error)
	DeprecateDomainAsync(ctx workflow.Context, input *swf.DeprecateDomainInput) *SWFDeprecateDomainFuture

	DeprecateWorkflowType(ctx workflow.Context, input *swf.DeprecateWorkflowTypeInput) (*swf.DeprecateWorkflowTypeOutput, error)
	DeprecateWorkflowTypeAsync(ctx workflow.Context, input *swf.DeprecateWorkflowTypeInput) *SWFDeprecateWorkflowTypeFuture

	DescribeActivityType(ctx workflow.Context, input *swf.DescribeActivityTypeInput) (*swf.DescribeActivityTypeOutput, error)
	DescribeActivityTypeAsync(ctx workflow.Context, input *swf.DescribeActivityTypeInput) *SWFDescribeActivityTypeFuture

	DescribeDomain(ctx workflow.Context, input *swf.DescribeDomainInput) (*swf.DescribeDomainOutput, error)
	DescribeDomainAsync(ctx workflow.Context, input *swf.DescribeDomainInput) *SWFDescribeDomainFuture

	DescribeWorkflowExecution(ctx workflow.Context, input *swf.DescribeWorkflowExecutionInput) (*swf.DescribeWorkflowExecutionOutput, error)
	DescribeWorkflowExecutionAsync(ctx workflow.Context, input *swf.DescribeWorkflowExecutionInput) *SWFDescribeWorkflowExecutionFuture

	DescribeWorkflowType(ctx workflow.Context, input *swf.DescribeWorkflowTypeInput) (*swf.DescribeWorkflowTypeOutput, error)
	DescribeWorkflowTypeAsync(ctx workflow.Context, input *swf.DescribeWorkflowTypeInput) *SWFDescribeWorkflowTypeFuture

	GetWorkflowExecutionHistory(ctx workflow.Context, input *swf.GetWorkflowExecutionHistoryInput) (*swf.GetWorkflowExecutionHistoryOutput, error)
	GetWorkflowExecutionHistoryAsync(ctx workflow.Context, input *swf.GetWorkflowExecutionHistoryInput) *SWFGetWorkflowExecutionHistoryFuture

	ListActivityTypes(ctx workflow.Context, input *swf.ListActivityTypesInput) (*swf.ListActivityTypesOutput, error)
	ListActivityTypesAsync(ctx workflow.Context, input *swf.ListActivityTypesInput) *SWFListActivityTypesFuture

	ListClosedWorkflowExecutions(ctx workflow.Context, input *swf.ListClosedWorkflowExecutionsInput) (*swf.WorkflowExecutionInfos, error)
	ListClosedWorkflowExecutionsAsync(ctx workflow.Context, input *swf.ListClosedWorkflowExecutionsInput) *SWFListClosedWorkflowExecutionsFuture

	ListDomains(ctx workflow.Context, input *swf.ListDomainsInput) (*swf.ListDomainsOutput, error)
	ListDomainsAsync(ctx workflow.Context, input *swf.ListDomainsInput) *SWFListDomainsFuture

	ListOpenWorkflowExecutions(ctx workflow.Context, input *swf.ListOpenWorkflowExecutionsInput) (*swf.WorkflowExecutionInfos, error)
	ListOpenWorkflowExecutionsAsync(ctx workflow.Context, input *swf.ListOpenWorkflowExecutionsInput) *SWFListOpenWorkflowExecutionsFuture

	ListTagsForResource(ctx workflow.Context, input *swf.ListTagsForResourceInput) (*swf.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *swf.ListTagsForResourceInput) *SWFListTagsForResourceFuture

	ListWorkflowTypes(ctx workflow.Context, input *swf.ListWorkflowTypesInput) (*swf.ListWorkflowTypesOutput, error)
	ListWorkflowTypesAsync(ctx workflow.Context, input *swf.ListWorkflowTypesInput) *SWFListWorkflowTypesFuture

	PollForActivityTask(ctx workflow.Context, input *swf.PollForActivityTaskInput) (*swf.PollForActivityTaskOutput, error)
	PollForActivityTaskAsync(ctx workflow.Context, input *swf.PollForActivityTaskInput) *SWFPollForActivityTaskFuture

	PollForDecisionTask(ctx workflow.Context, input *swf.PollForDecisionTaskInput) (*swf.PollForDecisionTaskOutput, error)
	PollForDecisionTaskAsync(ctx workflow.Context, input *swf.PollForDecisionTaskInput) *SWFPollForDecisionTaskFuture

	RecordActivityTaskHeartbeat(ctx workflow.Context, input *swf.RecordActivityTaskHeartbeatInput) (*swf.RecordActivityTaskHeartbeatOutput, error)
	RecordActivityTaskHeartbeatAsync(ctx workflow.Context, input *swf.RecordActivityTaskHeartbeatInput) *SWFRecordActivityTaskHeartbeatFuture

	RegisterActivityType(ctx workflow.Context, input *swf.RegisterActivityTypeInput) (*swf.RegisterActivityTypeOutput, error)
	RegisterActivityTypeAsync(ctx workflow.Context, input *swf.RegisterActivityTypeInput) *SWFRegisterActivityTypeFuture

	RegisterDomain(ctx workflow.Context, input *swf.RegisterDomainInput) (*swf.RegisterDomainOutput, error)
	RegisterDomainAsync(ctx workflow.Context, input *swf.RegisterDomainInput) *SWFRegisterDomainFuture

	RegisterWorkflowType(ctx workflow.Context, input *swf.RegisterWorkflowTypeInput) (*swf.RegisterWorkflowTypeOutput, error)
	RegisterWorkflowTypeAsync(ctx workflow.Context, input *swf.RegisterWorkflowTypeInput) *SWFRegisterWorkflowTypeFuture

	RequestCancelWorkflowExecution(ctx workflow.Context, input *swf.RequestCancelWorkflowExecutionInput) (*swf.RequestCancelWorkflowExecutionOutput, error)
	RequestCancelWorkflowExecutionAsync(ctx workflow.Context, input *swf.RequestCancelWorkflowExecutionInput) *SWFRequestCancelWorkflowExecutionFuture

	RespondActivityTaskCanceled(ctx workflow.Context, input *swf.RespondActivityTaskCanceledInput) (*swf.RespondActivityTaskCanceledOutput, error)
	RespondActivityTaskCanceledAsync(ctx workflow.Context, input *swf.RespondActivityTaskCanceledInput) *SWFRespondActivityTaskCanceledFuture

	RespondActivityTaskCompleted(ctx workflow.Context, input *swf.RespondActivityTaskCompletedInput) (*swf.RespondActivityTaskCompletedOutput, error)
	RespondActivityTaskCompletedAsync(ctx workflow.Context, input *swf.RespondActivityTaskCompletedInput) *SWFRespondActivityTaskCompletedFuture

	RespondActivityTaskFailed(ctx workflow.Context, input *swf.RespondActivityTaskFailedInput) (*swf.RespondActivityTaskFailedOutput, error)
	RespondActivityTaskFailedAsync(ctx workflow.Context, input *swf.RespondActivityTaskFailedInput) *SWFRespondActivityTaskFailedFuture

	RespondDecisionTaskCompleted(ctx workflow.Context, input *swf.RespondDecisionTaskCompletedInput) (*swf.RespondDecisionTaskCompletedOutput, error)
	RespondDecisionTaskCompletedAsync(ctx workflow.Context, input *swf.RespondDecisionTaskCompletedInput) *SWFRespondDecisionTaskCompletedFuture

	SignalWorkflowExecution(ctx workflow.Context, input *swf.SignalWorkflowExecutionInput) (*swf.SignalWorkflowExecutionOutput, error)
	SignalWorkflowExecutionAsync(ctx workflow.Context, input *swf.SignalWorkflowExecutionInput) *SWFSignalWorkflowExecutionFuture

	StartWorkflowExecution(ctx workflow.Context, input *swf.StartWorkflowExecutionInput) (*swf.StartWorkflowExecutionOutput, error)
	StartWorkflowExecutionAsync(ctx workflow.Context, input *swf.StartWorkflowExecutionInput) *SWFStartWorkflowExecutionFuture

	TagResource(ctx workflow.Context, input *swf.TagResourceInput) (*swf.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *swf.TagResourceInput) *SWFTagResourceFuture

	TerminateWorkflowExecution(ctx workflow.Context, input *swf.TerminateWorkflowExecutionInput) (*swf.TerminateWorkflowExecutionOutput, error)
	TerminateWorkflowExecutionAsync(ctx workflow.Context, input *swf.TerminateWorkflowExecutionInput) *SWFTerminateWorkflowExecutionFuture

	UndeprecateActivityType(ctx workflow.Context, input *swf.UndeprecateActivityTypeInput) (*swf.UndeprecateActivityTypeOutput, error)
	UndeprecateActivityTypeAsync(ctx workflow.Context, input *swf.UndeprecateActivityTypeInput) *SWFUndeprecateActivityTypeFuture

	UndeprecateDomain(ctx workflow.Context, input *swf.UndeprecateDomainInput) (*swf.UndeprecateDomainOutput, error)
	UndeprecateDomainAsync(ctx workflow.Context, input *swf.UndeprecateDomainInput) *SWFUndeprecateDomainFuture

	UndeprecateWorkflowType(ctx workflow.Context, input *swf.UndeprecateWorkflowTypeInput) (*swf.UndeprecateWorkflowTypeOutput, error)
	UndeprecateWorkflowTypeAsync(ctx workflow.Context, input *swf.UndeprecateWorkflowTypeInput) *SWFUndeprecateWorkflowTypeFuture

	UntagResource(ctx workflow.Context, input *swf.UntagResourceInput) (*swf.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *swf.UntagResourceInput) *SWFUntagResourceFuture
}

func NewClient() Client {
	return &stub{}
}

