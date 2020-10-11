// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"go.temporal.io/sdk/workflow"
)

type SQSClient interface {
	AddPermission(ctx workflow.Context, input *sqs.AddPermissionInput) (*sqs.AddPermissionOutput, error)
	AddPermissionAsync(ctx workflow.Context, input *sqs.AddPermissionInput) *SqsAddPermissionFuture

	ChangeMessageVisibility(ctx workflow.Context, input *sqs.ChangeMessageVisibilityInput) (*sqs.ChangeMessageVisibilityOutput, error)
	ChangeMessageVisibilityAsync(ctx workflow.Context, input *sqs.ChangeMessageVisibilityInput) *SqsChangeMessageVisibilityFuture

	ChangeMessageVisibilityBatch(ctx workflow.Context, input *sqs.ChangeMessageVisibilityBatchInput) (*sqs.ChangeMessageVisibilityBatchOutput, error)
	ChangeMessageVisibilityBatchAsync(ctx workflow.Context, input *sqs.ChangeMessageVisibilityBatchInput) *SqsChangeMessageVisibilityBatchFuture

	CreateQueue(ctx workflow.Context, input *sqs.CreateQueueInput) (*sqs.CreateQueueOutput, error)
	CreateQueueAsync(ctx workflow.Context, input *sqs.CreateQueueInput) *SqsCreateQueueFuture

	DeleteMessage(ctx workflow.Context, input *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error)
	DeleteMessageAsync(ctx workflow.Context, input *sqs.DeleteMessageInput) *SqsDeleteMessageFuture

	DeleteMessageBatch(ctx workflow.Context, input *sqs.DeleteMessageBatchInput) (*sqs.DeleteMessageBatchOutput, error)
	DeleteMessageBatchAsync(ctx workflow.Context, input *sqs.DeleteMessageBatchInput) *SqsDeleteMessageBatchFuture

	DeleteQueue(ctx workflow.Context, input *sqs.DeleteQueueInput) (*sqs.DeleteQueueOutput, error)
	DeleteQueueAsync(ctx workflow.Context, input *sqs.DeleteQueueInput) *SqsDeleteQueueFuture

	GetQueueAttributes(ctx workflow.Context, input *sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error)
	GetQueueAttributesAsync(ctx workflow.Context, input *sqs.GetQueueAttributesInput) *SqsGetQueueAttributesFuture

	GetQueueUrl(ctx workflow.Context, input *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error)
	GetQueueUrlAsync(ctx workflow.Context, input *sqs.GetQueueUrlInput) *SqsGetQueueUrlFuture

	ListDeadLetterSourceQueues(ctx workflow.Context, input *sqs.ListDeadLetterSourceQueuesInput) (*sqs.ListDeadLetterSourceQueuesOutput, error)
	ListDeadLetterSourceQueuesAsync(ctx workflow.Context, input *sqs.ListDeadLetterSourceQueuesInput) *SqsListDeadLetterSourceQueuesFuture

	ListQueueTags(ctx workflow.Context, input *sqs.ListQueueTagsInput) (*sqs.ListQueueTagsOutput, error)
	ListQueueTagsAsync(ctx workflow.Context, input *sqs.ListQueueTagsInput) *SqsListQueueTagsFuture

	ListQueues(ctx workflow.Context, input *sqs.ListQueuesInput) (*sqs.ListQueuesOutput, error)
	ListQueuesAsync(ctx workflow.Context, input *sqs.ListQueuesInput) *SqsListQueuesFuture

	PurgeQueue(ctx workflow.Context, input *sqs.PurgeQueueInput) (*sqs.PurgeQueueOutput, error)
	PurgeQueueAsync(ctx workflow.Context, input *sqs.PurgeQueueInput) *SqsPurgeQueueFuture

	ReceiveMessage(ctx workflow.Context, input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error)
	ReceiveMessageAsync(ctx workflow.Context, input *sqs.ReceiveMessageInput) *SqsReceiveMessageFuture

	RemovePermission(ctx workflow.Context, input *sqs.RemovePermissionInput) (*sqs.RemovePermissionOutput, error)
	RemovePermissionAsync(ctx workflow.Context, input *sqs.RemovePermissionInput) *SqsRemovePermissionFuture

	SendMessage(ctx workflow.Context, input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error)
	SendMessageAsync(ctx workflow.Context, input *sqs.SendMessageInput) *SqsSendMessageFuture

	SendMessageBatch(ctx workflow.Context, input *sqs.SendMessageBatchInput) (*sqs.SendMessageBatchOutput, error)
	SendMessageBatchAsync(ctx workflow.Context, input *sqs.SendMessageBatchInput) *SqsSendMessageBatchFuture

	SetQueueAttributes(ctx workflow.Context, input *sqs.SetQueueAttributesInput) (*sqs.SetQueueAttributesOutput, error)
	SetQueueAttributesAsync(ctx workflow.Context, input *sqs.SetQueueAttributesInput) *SqsSetQueueAttributesFuture

	TagQueue(ctx workflow.Context, input *sqs.TagQueueInput) (*sqs.TagQueueOutput, error)
	TagQueueAsync(ctx workflow.Context, input *sqs.TagQueueInput) *SqsTagQueueFuture

	UntagQueue(ctx workflow.Context, input *sqs.UntagQueueInput) (*sqs.UntagQueueOutput, error)
	UntagQueueAsync(ctx workflow.Context, input *sqs.UntagQueueInput) *SqsUntagQueueFuture
}

type SQSStub struct{}

func NewSQSStub() SQSClient {
	return &SQSStub{}
}

type SqsAddPermissionFuture struct {
	Future workflow.Future
}

func (r *SqsAddPermissionFuture) Get(ctx workflow.Context) (*sqs.AddPermissionOutput, error) {
	var output sqs.AddPermissionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SqsChangeMessageVisibilityFuture struct {
	Future workflow.Future
}

func (r *SqsChangeMessageVisibilityFuture) Get(ctx workflow.Context) (*sqs.ChangeMessageVisibilityOutput, error) {
	var output sqs.ChangeMessageVisibilityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SqsChangeMessageVisibilityBatchFuture struct {
	Future workflow.Future
}

func (r *SqsChangeMessageVisibilityBatchFuture) Get(ctx workflow.Context) (*sqs.ChangeMessageVisibilityBatchOutput, error) {
	var output sqs.ChangeMessageVisibilityBatchOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SqsCreateQueueFuture struct {
	Future workflow.Future
}

func (r *SqsCreateQueueFuture) Get(ctx workflow.Context) (*sqs.CreateQueueOutput, error) {
	var output sqs.CreateQueueOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SqsDeleteMessageFuture struct {
	Future workflow.Future
}

func (r *SqsDeleteMessageFuture) Get(ctx workflow.Context) (*sqs.DeleteMessageOutput, error) {
	var output sqs.DeleteMessageOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SqsDeleteMessageBatchFuture struct {
	Future workflow.Future
}

func (r *SqsDeleteMessageBatchFuture) Get(ctx workflow.Context) (*sqs.DeleteMessageBatchOutput, error) {
	var output sqs.DeleteMessageBatchOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SqsDeleteQueueFuture struct {
	Future workflow.Future
}

func (r *SqsDeleteQueueFuture) Get(ctx workflow.Context) (*sqs.DeleteQueueOutput, error) {
	var output sqs.DeleteQueueOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SqsGetQueueAttributesFuture struct {
	Future workflow.Future
}

func (r *SqsGetQueueAttributesFuture) Get(ctx workflow.Context) (*sqs.GetQueueAttributesOutput, error) {
	var output sqs.GetQueueAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SqsGetQueueUrlFuture struct {
	Future workflow.Future
}

func (r *SqsGetQueueUrlFuture) Get(ctx workflow.Context) (*sqs.GetQueueUrlOutput, error) {
	var output sqs.GetQueueUrlOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SqsListDeadLetterSourceQueuesFuture struct {
	Future workflow.Future
}

func (r *SqsListDeadLetterSourceQueuesFuture) Get(ctx workflow.Context) (*sqs.ListDeadLetterSourceQueuesOutput, error) {
	var output sqs.ListDeadLetterSourceQueuesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SqsListQueueTagsFuture struct {
	Future workflow.Future
}

func (r *SqsListQueueTagsFuture) Get(ctx workflow.Context) (*sqs.ListQueueTagsOutput, error) {
	var output sqs.ListQueueTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SqsListQueuesFuture struct {
	Future workflow.Future
}

func (r *SqsListQueuesFuture) Get(ctx workflow.Context) (*sqs.ListQueuesOutput, error) {
	var output sqs.ListQueuesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SqsPurgeQueueFuture struct {
	Future workflow.Future
}

func (r *SqsPurgeQueueFuture) Get(ctx workflow.Context) (*sqs.PurgeQueueOutput, error) {
	var output sqs.PurgeQueueOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SqsReceiveMessageFuture struct {
	Future workflow.Future
}

func (r *SqsReceiveMessageFuture) Get(ctx workflow.Context) (*sqs.ReceiveMessageOutput, error) {
	var output sqs.ReceiveMessageOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SqsRemovePermissionFuture struct {
	Future workflow.Future
}

func (r *SqsRemovePermissionFuture) Get(ctx workflow.Context) (*sqs.RemovePermissionOutput, error) {
	var output sqs.RemovePermissionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SqsSendMessageFuture struct {
	Future workflow.Future
}

func (r *SqsSendMessageFuture) Get(ctx workflow.Context) (*sqs.SendMessageOutput, error) {
	var output sqs.SendMessageOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SqsSendMessageBatchFuture struct {
	Future workflow.Future
}

func (r *SqsSendMessageBatchFuture) Get(ctx workflow.Context) (*sqs.SendMessageBatchOutput, error) {
	var output sqs.SendMessageBatchOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SqsSetQueueAttributesFuture struct {
	Future workflow.Future
}

func (r *SqsSetQueueAttributesFuture) Get(ctx workflow.Context) (*sqs.SetQueueAttributesOutput, error) {
	var output sqs.SetQueueAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SqsTagQueueFuture struct {
	Future workflow.Future
}

func (r *SqsTagQueueFuture) Get(ctx workflow.Context) (*sqs.TagQueueOutput, error) {
	var output sqs.TagQueueOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SqsUntagQueueFuture struct {
	Future workflow.Future
}

func (r *SqsUntagQueueFuture) Get(ctx workflow.Context) (*sqs.UntagQueueOutput, error) {
	var output sqs.UntagQueueOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *SQSStub) AddPermission(ctx workflow.Context, input *sqs.AddPermissionInput) (*sqs.AddPermissionOutput, error) {
	var output sqs.AddPermissionOutput
	err := workflow.ExecuteActivity(ctx, "aws.sqs.AddPermission", input).Get(ctx, &output)
	return &output, err
}

func (a *SQSStub) AddPermissionAsync(ctx workflow.Context, input *sqs.AddPermissionInput) *SqsAddPermissionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sqs.AddPermission", input)
	return &SqsAddPermissionFuture{Future: future}
}

func (a *SQSStub) ChangeMessageVisibility(ctx workflow.Context, input *sqs.ChangeMessageVisibilityInput) (*sqs.ChangeMessageVisibilityOutput, error) {
	var output sqs.ChangeMessageVisibilityOutput
	err := workflow.ExecuteActivity(ctx, "aws.sqs.ChangeMessageVisibility", input).Get(ctx, &output)
	return &output, err
}

func (a *SQSStub) ChangeMessageVisibilityAsync(ctx workflow.Context, input *sqs.ChangeMessageVisibilityInput) *SqsChangeMessageVisibilityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sqs.ChangeMessageVisibility", input)
	return &SqsChangeMessageVisibilityFuture{Future: future}
}

func (a *SQSStub) ChangeMessageVisibilityBatch(ctx workflow.Context, input *sqs.ChangeMessageVisibilityBatchInput) (*sqs.ChangeMessageVisibilityBatchOutput, error) {
	var output sqs.ChangeMessageVisibilityBatchOutput
	err := workflow.ExecuteActivity(ctx, "aws.sqs.ChangeMessageVisibilityBatch", input).Get(ctx, &output)
	return &output, err
}

func (a *SQSStub) ChangeMessageVisibilityBatchAsync(ctx workflow.Context, input *sqs.ChangeMessageVisibilityBatchInput) *SqsChangeMessageVisibilityBatchFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sqs.ChangeMessageVisibilityBatch", input)
	return &SqsChangeMessageVisibilityBatchFuture{Future: future}
}

func (a *SQSStub) CreateQueue(ctx workflow.Context, input *sqs.CreateQueueInput) (*sqs.CreateQueueOutput, error) {
	var output sqs.CreateQueueOutput
	err := workflow.ExecuteActivity(ctx, "aws.sqs.CreateQueue", input).Get(ctx, &output)
	return &output, err
}

func (a *SQSStub) CreateQueueAsync(ctx workflow.Context, input *sqs.CreateQueueInput) *SqsCreateQueueFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sqs.CreateQueue", input)
	return &SqsCreateQueueFuture{Future: future}
}

func (a *SQSStub) DeleteMessage(ctx workflow.Context, input *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {
	var output sqs.DeleteMessageOutput
	err := workflow.ExecuteActivity(ctx, "aws.sqs.DeleteMessage", input).Get(ctx, &output)
	return &output, err
}

func (a *SQSStub) DeleteMessageAsync(ctx workflow.Context, input *sqs.DeleteMessageInput) *SqsDeleteMessageFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sqs.DeleteMessage", input)
	return &SqsDeleteMessageFuture{Future: future}
}

func (a *SQSStub) DeleteMessageBatch(ctx workflow.Context, input *sqs.DeleteMessageBatchInput) (*sqs.DeleteMessageBatchOutput, error) {
	var output sqs.DeleteMessageBatchOutput
	err := workflow.ExecuteActivity(ctx, "aws.sqs.DeleteMessageBatch", input).Get(ctx, &output)
	return &output, err
}

func (a *SQSStub) DeleteMessageBatchAsync(ctx workflow.Context, input *sqs.DeleteMessageBatchInput) *SqsDeleteMessageBatchFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sqs.DeleteMessageBatch", input)
	return &SqsDeleteMessageBatchFuture{Future: future}
}

func (a *SQSStub) DeleteQueue(ctx workflow.Context, input *sqs.DeleteQueueInput) (*sqs.DeleteQueueOutput, error) {
	var output sqs.DeleteQueueOutput
	err := workflow.ExecuteActivity(ctx, "aws.sqs.DeleteQueue", input).Get(ctx, &output)
	return &output, err
}

func (a *SQSStub) DeleteQueueAsync(ctx workflow.Context, input *sqs.DeleteQueueInput) *SqsDeleteQueueFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sqs.DeleteQueue", input)
	return &SqsDeleteQueueFuture{Future: future}
}

func (a *SQSStub) GetQueueAttributes(ctx workflow.Context, input *sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error) {
	var output sqs.GetQueueAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.sqs.GetQueueAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *SQSStub) GetQueueAttributesAsync(ctx workflow.Context, input *sqs.GetQueueAttributesInput) *SqsGetQueueAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sqs.GetQueueAttributes", input)
	return &SqsGetQueueAttributesFuture{Future: future}
}

func (a *SQSStub) GetQueueUrl(ctx workflow.Context, input *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
	var output sqs.GetQueueUrlOutput
	err := workflow.ExecuteActivity(ctx, "aws.sqs.GetQueueUrl", input).Get(ctx, &output)
	return &output, err
}

func (a *SQSStub) GetQueueUrlAsync(ctx workflow.Context, input *sqs.GetQueueUrlInput) *SqsGetQueueUrlFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sqs.GetQueueUrl", input)
	return &SqsGetQueueUrlFuture{Future: future}
}

func (a *SQSStub) ListDeadLetterSourceQueues(ctx workflow.Context, input *sqs.ListDeadLetterSourceQueuesInput) (*sqs.ListDeadLetterSourceQueuesOutput, error) {
	var output sqs.ListDeadLetterSourceQueuesOutput
	err := workflow.ExecuteActivity(ctx, "aws.sqs.ListDeadLetterSourceQueues", input).Get(ctx, &output)
	return &output, err
}

func (a *SQSStub) ListDeadLetterSourceQueuesAsync(ctx workflow.Context, input *sqs.ListDeadLetterSourceQueuesInput) *SqsListDeadLetterSourceQueuesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sqs.ListDeadLetterSourceQueues", input)
	return &SqsListDeadLetterSourceQueuesFuture{Future: future}
}

func (a *SQSStub) ListQueueTags(ctx workflow.Context, input *sqs.ListQueueTagsInput) (*sqs.ListQueueTagsOutput, error) {
	var output sqs.ListQueueTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.sqs.ListQueueTags", input).Get(ctx, &output)
	return &output, err
}

func (a *SQSStub) ListQueueTagsAsync(ctx workflow.Context, input *sqs.ListQueueTagsInput) *SqsListQueueTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sqs.ListQueueTags", input)
	return &SqsListQueueTagsFuture{Future: future}
}

func (a *SQSStub) ListQueues(ctx workflow.Context, input *sqs.ListQueuesInput) (*sqs.ListQueuesOutput, error) {
	var output sqs.ListQueuesOutput
	err := workflow.ExecuteActivity(ctx, "aws.sqs.ListQueues", input).Get(ctx, &output)
	return &output, err
}

func (a *SQSStub) ListQueuesAsync(ctx workflow.Context, input *sqs.ListQueuesInput) *SqsListQueuesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sqs.ListQueues", input)
	return &SqsListQueuesFuture{Future: future}
}

func (a *SQSStub) PurgeQueue(ctx workflow.Context, input *sqs.PurgeQueueInput) (*sqs.PurgeQueueOutput, error) {
	var output sqs.PurgeQueueOutput
	err := workflow.ExecuteActivity(ctx, "aws.sqs.PurgeQueue", input).Get(ctx, &output)
	return &output, err
}

func (a *SQSStub) PurgeQueueAsync(ctx workflow.Context, input *sqs.PurgeQueueInput) *SqsPurgeQueueFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sqs.PurgeQueue", input)
	return &SqsPurgeQueueFuture{Future: future}
}

func (a *SQSStub) ReceiveMessage(ctx workflow.Context, input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	var output sqs.ReceiveMessageOutput
	err := workflow.ExecuteActivity(ctx, "aws.sqs.ReceiveMessage", input).Get(ctx, &output)
	return &output, err
}

func (a *SQSStub) ReceiveMessageAsync(ctx workflow.Context, input *sqs.ReceiveMessageInput) *SqsReceiveMessageFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sqs.ReceiveMessage", input)
	return &SqsReceiveMessageFuture{Future: future}
}

func (a *SQSStub) RemovePermission(ctx workflow.Context, input *sqs.RemovePermissionInput) (*sqs.RemovePermissionOutput, error) {
	var output sqs.RemovePermissionOutput
	err := workflow.ExecuteActivity(ctx, "aws.sqs.RemovePermission", input).Get(ctx, &output)
	return &output, err
}

func (a *SQSStub) RemovePermissionAsync(ctx workflow.Context, input *sqs.RemovePermissionInput) *SqsRemovePermissionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sqs.RemovePermission", input)
	return &SqsRemovePermissionFuture{Future: future}
}

func (a *SQSStub) SendMessage(ctx workflow.Context, input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	var output sqs.SendMessageOutput
	err := workflow.ExecuteActivity(ctx, "aws.sqs.SendMessage", input).Get(ctx, &output)
	return &output, err
}

func (a *SQSStub) SendMessageAsync(ctx workflow.Context, input *sqs.SendMessageInput) *SqsSendMessageFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sqs.SendMessage", input)
	return &SqsSendMessageFuture{Future: future}
}

func (a *SQSStub) SendMessageBatch(ctx workflow.Context, input *sqs.SendMessageBatchInput) (*sqs.SendMessageBatchOutput, error) {
	var output sqs.SendMessageBatchOutput
	err := workflow.ExecuteActivity(ctx, "aws.sqs.SendMessageBatch", input).Get(ctx, &output)
	return &output, err
}

func (a *SQSStub) SendMessageBatchAsync(ctx workflow.Context, input *sqs.SendMessageBatchInput) *SqsSendMessageBatchFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sqs.SendMessageBatch", input)
	return &SqsSendMessageBatchFuture{Future: future}
}

func (a *SQSStub) SetQueueAttributes(ctx workflow.Context, input *sqs.SetQueueAttributesInput) (*sqs.SetQueueAttributesOutput, error) {
	var output sqs.SetQueueAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.sqs.SetQueueAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *SQSStub) SetQueueAttributesAsync(ctx workflow.Context, input *sqs.SetQueueAttributesInput) *SqsSetQueueAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sqs.SetQueueAttributes", input)
	return &SqsSetQueueAttributesFuture{Future: future}
}

func (a *SQSStub) TagQueue(ctx workflow.Context, input *sqs.TagQueueInput) (*sqs.TagQueueOutput, error) {
	var output sqs.TagQueueOutput
	err := workflow.ExecuteActivity(ctx, "aws.sqs.TagQueue", input).Get(ctx, &output)
	return &output, err
}

func (a *SQSStub) TagQueueAsync(ctx workflow.Context, input *sqs.TagQueueInput) *SqsTagQueueFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sqs.TagQueue", input)
	return &SqsTagQueueFuture{Future: future}
}

func (a *SQSStub) UntagQueue(ctx workflow.Context, input *sqs.UntagQueueInput) (*sqs.UntagQueueOutput, error) {
	var output sqs.UntagQueueOutput
	err := workflow.ExecuteActivity(ctx, "aws.sqs.UntagQueue", input).Get(ctx, &output)
	return &output, err
}

func (a *SQSStub) UntagQueueAsync(ctx workflow.Context, input *sqs.UntagQueueInput) *SqsUntagQueueFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sqs.UntagQueue", input)
	return &SqsUntagQueueFuture{Future: future}
}
