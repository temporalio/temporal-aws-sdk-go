// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package snsstub

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type AddPermissionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AddPermissionFuture) Get(ctx workflow.Context) (*sns.AddPermissionOutput, error) {
	var output sns.AddPermissionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CheckIfPhoneNumberIsOptedOutFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CheckIfPhoneNumberIsOptedOutFuture) Get(ctx workflow.Context) (*sns.CheckIfPhoneNumberIsOptedOutOutput, error) {
	var output sns.CheckIfPhoneNumberIsOptedOutOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ConfirmSubscriptionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ConfirmSubscriptionFuture) Get(ctx workflow.Context) (*sns.ConfirmSubscriptionOutput, error) {
	var output sns.ConfirmSubscriptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreatePlatformApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreatePlatformApplicationFuture) Get(ctx workflow.Context) (*sns.CreatePlatformApplicationOutput, error) {
	var output sns.CreatePlatformApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreatePlatformEndpointFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreatePlatformEndpointFuture) Get(ctx workflow.Context) (*sns.CreatePlatformEndpointOutput, error) {
	var output sns.CreatePlatformEndpointOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateTopicFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateTopicFuture) Get(ctx workflow.Context) (*sns.CreateTopicOutput, error) {
	var output sns.CreateTopicOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteEndpointFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteEndpointFuture) Get(ctx workflow.Context) (*sns.DeleteEndpointOutput, error) {
	var output sns.DeleteEndpointOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeletePlatformApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeletePlatformApplicationFuture) Get(ctx workflow.Context) (*sns.DeletePlatformApplicationOutput, error) {
	var output sns.DeletePlatformApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteTopicFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteTopicFuture) Get(ctx workflow.Context) (*sns.DeleteTopicOutput, error) {
	var output sns.DeleteTopicOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetEndpointAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetEndpointAttributesFuture) Get(ctx workflow.Context) (*sns.GetEndpointAttributesOutput, error) {
	var output sns.GetEndpointAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetPlatformApplicationAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetPlatformApplicationAttributesFuture) Get(ctx workflow.Context) (*sns.GetPlatformApplicationAttributesOutput, error) {
	var output sns.GetPlatformApplicationAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetSMSAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetSMSAttributesFuture) Get(ctx workflow.Context) (*sns.GetSMSAttributesOutput, error) {
	var output sns.GetSMSAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetSubscriptionAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetSubscriptionAttributesFuture) Get(ctx workflow.Context) (*sns.GetSubscriptionAttributesOutput, error) {
	var output sns.GetSubscriptionAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetTopicAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetTopicAttributesFuture) Get(ctx workflow.Context) (*sns.GetTopicAttributesOutput, error) {
	var output sns.GetTopicAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListEndpointsByPlatformApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListEndpointsByPlatformApplicationFuture) Get(ctx workflow.Context) (*sns.ListEndpointsByPlatformApplicationOutput, error) {
	var output sns.ListEndpointsByPlatformApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListPhoneNumbersOptedOutFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListPhoneNumbersOptedOutFuture) Get(ctx workflow.Context) (*sns.ListPhoneNumbersOptedOutOutput, error) {
	var output sns.ListPhoneNumbersOptedOutOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListPlatformApplicationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListPlatformApplicationsFuture) Get(ctx workflow.Context) (*sns.ListPlatformApplicationsOutput, error) {
	var output sns.ListPlatformApplicationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListSubscriptionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListSubscriptionsFuture) Get(ctx workflow.Context) (*sns.ListSubscriptionsOutput, error) {
	var output sns.ListSubscriptionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListSubscriptionsByTopicFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListSubscriptionsByTopicFuture) Get(ctx workflow.Context) (*sns.ListSubscriptionsByTopicOutput, error) {
	var output sns.ListSubscriptionsByTopicOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*sns.ListTagsForResourceOutput, error) {
	var output sns.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTopicsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTopicsFuture) Get(ctx workflow.Context) (*sns.ListTopicsOutput, error) {
	var output sns.ListTopicsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type OptInPhoneNumberFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *OptInPhoneNumberFuture) Get(ctx workflow.Context) (*sns.OptInPhoneNumberOutput, error) {
	var output sns.OptInPhoneNumberOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PublishFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PublishFuture) Get(ctx workflow.Context) (*sns.PublishOutput, error) {
	var output sns.PublishOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RemovePermissionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RemovePermissionFuture) Get(ctx workflow.Context) (*sns.RemovePermissionOutput, error) {
	var output sns.RemovePermissionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SetEndpointAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SetEndpointAttributesFuture) Get(ctx workflow.Context) (*sns.SetEndpointAttributesOutput, error) {
	var output sns.SetEndpointAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SetPlatformApplicationAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SetPlatformApplicationAttributesFuture) Get(ctx workflow.Context) (*sns.SetPlatformApplicationAttributesOutput, error) {
	var output sns.SetPlatformApplicationAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SetSMSAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SetSMSAttributesFuture) Get(ctx workflow.Context) (*sns.SetSMSAttributesOutput, error) {
	var output sns.SetSMSAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SetSubscriptionAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SetSubscriptionAttributesFuture) Get(ctx workflow.Context) (*sns.SetSubscriptionAttributesOutput, error) {
	var output sns.SetSubscriptionAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SetTopicAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SetTopicAttributesFuture) Get(ctx workflow.Context) (*sns.SetTopicAttributesOutput, error) {
	var output sns.SetTopicAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SubscribeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SubscribeFuture) Get(ctx workflow.Context) (*sns.SubscribeOutput, error) {
	var output sns.SubscribeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*sns.TagResourceOutput, error) {
	var output sns.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UnsubscribeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UnsubscribeFuture) Get(ctx workflow.Context) (*sns.UnsubscribeOutput, error) {
	var output sns.UnsubscribeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*sns.UntagResourceOutput, error) {
	var output sns.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AddPermission(ctx workflow.Context, input *sns.AddPermissionInput) (*sns.AddPermissionOutput, error) {
	var output sns.AddPermissionOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.AddPermission", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddPermissionAsync(ctx workflow.Context, input *sns.AddPermissionInput) *AddPermissionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.AddPermission", input)
	return &AddPermissionFuture{Future: future}
}

func (a *stub) CheckIfPhoneNumberIsOptedOut(ctx workflow.Context, input *sns.CheckIfPhoneNumberIsOptedOutInput) (*sns.CheckIfPhoneNumberIsOptedOutOutput, error) {
	var output sns.CheckIfPhoneNumberIsOptedOutOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.CheckIfPhoneNumberIsOptedOut", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CheckIfPhoneNumberIsOptedOutAsync(ctx workflow.Context, input *sns.CheckIfPhoneNumberIsOptedOutInput) *CheckIfPhoneNumberIsOptedOutFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.CheckIfPhoneNumberIsOptedOut", input)
	return &CheckIfPhoneNumberIsOptedOutFuture{Future: future}
}

func (a *stub) ConfirmSubscription(ctx workflow.Context, input *sns.ConfirmSubscriptionInput) (*sns.ConfirmSubscriptionOutput, error) {
	var output sns.ConfirmSubscriptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.ConfirmSubscription", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ConfirmSubscriptionAsync(ctx workflow.Context, input *sns.ConfirmSubscriptionInput) *ConfirmSubscriptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.ConfirmSubscription", input)
	return &ConfirmSubscriptionFuture{Future: future}
}

func (a *stub) CreatePlatformApplication(ctx workflow.Context, input *sns.CreatePlatformApplicationInput) (*sns.CreatePlatformApplicationOutput, error) {
	var output sns.CreatePlatformApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.CreatePlatformApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreatePlatformApplicationAsync(ctx workflow.Context, input *sns.CreatePlatformApplicationInput) *CreatePlatformApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.CreatePlatformApplication", input)
	return &CreatePlatformApplicationFuture{Future: future}
}

func (a *stub) CreatePlatformEndpoint(ctx workflow.Context, input *sns.CreatePlatformEndpointInput) (*sns.CreatePlatformEndpointOutput, error) {
	var output sns.CreatePlatformEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.CreatePlatformEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreatePlatformEndpointAsync(ctx workflow.Context, input *sns.CreatePlatformEndpointInput) *CreatePlatformEndpointFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.CreatePlatformEndpoint", input)
	return &CreatePlatformEndpointFuture{Future: future}
}

func (a *stub) CreateTopic(ctx workflow.Context, input *sns.CreateTopicInput) (*sns.CreateTopicOutput, error) {
	var output sns.CreateTopicOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.CreateTopic", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateTopicAsync(ctx workflow.Context, input *sns.CreateTopicInput) *CreateTopicFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.CreateTopic", input)
	return &CreateTopicFuture{Future: future}
}

func (a *stub) DeleteEndpoint(ctx workflow.Context, input *sns.DeleteEndpointInput) (*sns.DeleteEndpointOutput, error) {
	var output sns.DeleteEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.DeleteEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteEndpointAsync(ctx workflow.Context, input *sns.DeleteEndpointInput) *DeleteEndpointFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.DeleteEndpoint", input)
	return &DeleteEndpointFuture{Future: future}
}

func (a *stub) DeletePlatformApplication(ctx workflow.Context, input *sns.DeletePlatformApplicationInput) (*sns.DeletePlatformApplicationOutput, error) {
	var output sns.DeletePlatformApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.DeletePlatformApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeletePlatformApplicationAsync(ctx workflow.Context, input *sns.DeletePlatformApplicationInput) *DeletePlatformApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.DeletePlatformApplication", input)
	return &DeletePlatformApplicationFuture{Future: future}
}

func (a *stub) DeleteTopic(ctx workflow.Context, input *sns.DeleteTopicInput) (*sns.DeleteTopicOutput, error) {
	var output sns.DeleteTopicOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.DeleteTopic", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteTopicAsync(ctx workflow.Context, input *sns.DeleteTopicInput) *DeleteTopicFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.DeleteTopic", input)
	return &DeleteTopicFuture{Future: future}
}

func (a *stub) GetEndpointAttributes(ctx workflow.Context, input *sns.GetEndpointAttributesInput) (*sns.GetEndpointAttributesOutput, error) {
	var output sns.GetEndpointAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.GetEndpointAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetEndpointAttributesAsync(ctx workflow.Context, input *sns.GetEndpointAttributesInput) *GetEndpointAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.GetEndpointAttributes", input)
	return &GetEndpointAttributesFuture{Future: future}
}

func (a *stub) GetPlatformApplicationAttributes(ctx workflow.Context, input *sns.GetPlatformApplicationAttributesInput) (*sns.GetPlatformApplicationAttributesOutput, error) {
	var output sns.GetPlatformApplicationAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.GetPlatformApplicationAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetPlatformApplicationAttributesAsync(ctx workflow.Context, input *sns.GetPlatformApplicationAttributesInput) *GetPlatformApplicationAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.GetPlatformApplicationAttributes", input)
	return &GetPlatformApplicationAttributesFuture{Future: future}
}

func (a *stub) GetSMSAttributes(ctx workflow.Context, input *sns.GetSMSAttributesInput) (*sns.GetSMSAttributesOutput, error) {
	var output sns.GetSMSAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.GetSMSAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetSMSAttributesAsync(ctx workflow.Context, input *sns.GetSMSAttributesInput) *GetSMSAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.GetSMSAttributes", input)
	return &GetSMSAttributesFuture{Future: future}
}

func (a *stub) GetSubscriptionAttributes(ctx workflow.Context, input *sns.GetSubscriptionAttributesInput) (*sns.GetSubscriptionAttributesOutput, error) {
	var output sns.GetSubscriptionAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.GetSubscriptionAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetSubscriptionAttributesAsync(ctx workflow.Context, input *sns.GetSubscriptionAttributesInput) *GetSubscriptionAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.GetSubscriptionAttributes", input)
	return &GetSubscriptionAttributesFuture{Future: future}
}

func (a *stub) GetTopicAttributes(ctx workflow.Context, input *sns.GetTopicAttributesInput) (*sns.GetTopicAttributesOutput, error) {
	var output sns.GetTopicAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.GetTopicAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetTopicAttributesAsync(ctx workflow.Context, input *sns.GetTopicAttributesInput) *GetTopicAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.GetTopicAttributes", input)
	return &GetTopicAttributesFuture{Future: future}
}

func (a *stub) ListEndpointsByPlatformApplication(ctx workflow.Context, input *sns.ListEndpointsByPlatformApplicationInput) (*sns.ListEndpointsByPlatformApplicationOutput, error) {
	var output sns.ListEndpointsByPlatformApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.ListEndpointsByPlatformApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListEndpointsByPlatformApplicationAsync(ctx workflow.Context, input *sns.ListEndpointsByPlatformApplicationInput) *ListEndpointsByPlatformApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.ListEndpointsByPlatformApplication", input)
	return &ListEndpointsByPlatformApplicationFuture{Future: future}
}

func (a *stub) ListPhoneNumbersOptedOut(ctx workflow.Context, input *sns.ListPhoneNumbersOptedOutInput) (*sns.ListPhoneNumbersOptedOutOutput, error) {
	var output sns.ListPhoneNumbersOptedOutOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.ListPhoneNumbersOptedOut", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPhoneNumbersOptedOutAsync(ctx workflow.Context, input *sns.ListPhoneNumbersOptedOutInput) *ListPhoneNumbersOptedOutFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.ListPhoneNumbersOptedOut", input)
	return &ListPhoneNumbersOptedOutFuture{Future: future}
}

func (a *stub) ListPlatformApplications(ctx workflow.Context, input *sns.ListPlatformApplicationsInput) (*sns.ListPlatformApplicationsOutput, error) {
	var output sns.ListPlatformApplicationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.ListPlatformApplications", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPlatformApplicationsAsync(ctx workflow.Context, input *sns.ListPlatformApplicationsInput) *ListPlatformApplicationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.ListPlatformApplications", input)
	return &ListPlatformApplicationsFuture{Future: future}
}

func (a *stub) ListSubscriptions(ctx workflow.Context, input *sns.ListSubscriptionsInput) (*sns.ListSubscriptionsOutput, error) {
	var output sns.ListSubscriptionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.ListSubscriptions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListSubscriptionsAsync(ctx workflow.Context, input *sns.ListSubscriptionsInput) *ListSubscriptionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.ListSubscriptions", input)
	return &ListSubscriptionsFuture{Future: future}
}

func (a *stub) ListSubscriptionsByTopic(ctx workflow.Context, input *sns.ListSubscriptionsByTopicInput) (*sns.ListSubscriptionsByTopicOutput, error) {
	var output sns.ListSubscriptionsByTopicOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.ListSubscriptionsByTopic", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListSubscriptionsByTopicAsync(ctx workflow.Context, input *sns.ListSubscriptionsByTopicInput) *ListSubscriptionsByTopicFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.ListSubscriptionsByTopic", input)
	return &ListSubscriptionsByTopicFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *sns.ListTagsForResourceInput) (*sns.ListTagsForResourceOutput, error) {
	var output sns.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *sns.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) ListTopics(ctx workflow.Context, input *sns.ListTopicsInput) (*sns.ListTopicsOutput, error) {
	var output sns.ListTopicsOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.ListTopics", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTopicsAsync(ctx workflow.Context, input *sns.ListTopicsInput) *ListTopicsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.ListTopics", input)
	return &ListTopicsFuture{Future: future}
}

func (a *stub) OptInPhoneNumber(ctx workflow.Context, input *sns.OptInPhoneNumberInput) (*sns.OptInPhoneNumberOutput, error) {
	var output sns.OptInPhoneNumberOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.OptInPhoneNumber", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) OptInPhoneNumberAsync(ctx workflow.Context, input *sns.OptInPhoneNumberInput) *OptInPhoneNumberFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.OptInPhoneNumber", input)
	return &OptInPhoneNumberFuture{Future: future}
}

func (a *stub) Publish(ctx workflow.Context, input *sns.PublishInput) (*sns.PublishOutput, error) {
	var output sns.PublishOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.Publish", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PublishAsync(ctx workflow.Context, input *sns.PublishInput) *PublishFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.Publish", input)
	return &PublishFuture{Future: future}
}

func (a *stub) RemovePermission(ctx workflow.Context, input *sns.RemovePermissionInput) (*sns.RemovePermissionOutput, error) {
	var output sns.RemovePermissionOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.RemovePermission", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemovePermissionAsync(ctx workflow.Context, input *sns.RemovePermissionInput) *RemovePermissionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.RemovePermission", input)
	return &RemovePermissionFuture{Future: future}
}

func (a *stub) SetEndpointAttributes(ctx workflow.Context, input *sns.SetEndpointAttributesInput) (*sns.SetEndpointAttributesOutput, error) {
	var output sns.SetEndpointAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.SetEndpointAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetEndpointAttributesAsync(ctx workflow.Context, input *sns.SetEndpointAttributesInput) *SetEndpointAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.SetEndpointAttributes", input)
	return &SetEndpointAttributesFuture{Future: future}
}

func (a *stub) SetPlatformApplicationAttributes(ctx workflow.Context, input *sns.SetPlatformApplicationAttributesInput) (*sns.SetPlatformApplicationAttributesOutput, error) {
	var output sns.SetPlatformApplicationAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.SetPlatformApplicationAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetPlatformApplicationAttributesAsync(ctx workflow.Context, input *sns.SetPlatformApplicationAttributesInput) *SetPlatformApplicationAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.SetPlatformApplicationAttributes", input)
	return &SetPlatformApplicationAttributesFuture{Future: future}
}

func (a *stub) SetSMSAttributes(ctx workflow.Context, input *sns.SetSMSAttributesInput) (*sns.SetSMSAttributesOutput, error) {
	var output sns.SetSMSAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.SetSMSAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetSMSAttributesAsync(ctx workflow.Context, input *sns.SetSMSAttributesInput) *SetSMSAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.SetSMSAttributes", input)
	return &SetSMSAttributesFuture{Future: future}
}

func (a *stub) SetSubscriptionAttributes(ctx workflow.Context, input *sns.SetSubscriptionAttributesInput) (*sns.SetSubscriptionAttributesOutput, error) {
	var output sns.SetSubscriptionAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.SetSubscriptionAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetSubscriptionAttributesAsync(ctx workflow.Context, input *sns.SetSubscriptionAttributesInput) *SetSubscriptionAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.SetSubscriptionAttributes", input)
	return &SetSubscriptionAttributesFuture{Future: future}
}

func (a *stub) SetTopicAttributes(ctx workflow.Context, input *sns.SetTopicAttributesInput) (*sns.SetTopicAttributesOutput, error) {
	var output sns.SetTopicAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.SetTopicAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetTopicAttributesAsync(ctx workflow.Context, input *sns.SetTopicAttributesInput) *SetTopicAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.SetTopicAttributes", input)
	return &SetTopicAttributesFuture{Future: future}
}

func (a *stub) Subscribe(ctx workflow.Context, input *sns.SubscribeInput) (*sns.SubscribeOutput, error) {
	var output sns.SubscribeOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.Subscribe", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SubscribeAsync(ctx workflow.Context, input *sns.SubscribeInput) *SubscribeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.Subscribe", input)
	return &SubscribeFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *sns.TagResourceInput) (*sns.TagResourceOutput, error) {
	var output sns.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *sns.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) Unsubscribe(ctx workflow.Context, input *sns.UnsubscribeInput) (*sns.UnsubscribeOutput, error) {
	var output sns.UnsubscribeOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.Unsubscribe", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UnsubscribeAsync(ctx workflow.Context, input *sns.UnsubscribeInput) *UnsubscribeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.Unsubscribe", input)
	return &UnsubscribeFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *sns.UntagResourceInput) (*sns.UntagResourceOutput, error) {
	var output sns.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.sns.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *sns.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.sns.UntagResource", input)
	return &UntagResourceFuture{Future: future}
}
