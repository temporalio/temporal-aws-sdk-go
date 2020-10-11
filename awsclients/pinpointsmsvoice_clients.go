// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/pinpointsmsvoice"
	"go.temporal.io/sdk/workflow"
)

type PinpointSMSVoiceClient interface {
	CreateConfigurationSet(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetInput) (*pinpointsmsvoice.CreateConfigurationSetOutput, error)
	CreateConfigurationSetAsync(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetInput) *PinpointsmsvoiceCreateConfigurationSetFuture

	CreateConfigurationSetEventDestination(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetEventDestinationInput) (*pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput, error)
	CreateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetEventDestinationInput) *PinpointsmsvoiceCreateConfigurationSetEventDestinationFuture

	DeleteConfigurationSet(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetInput) (*pinpointsmsvoice.DeleteConfigurationSetOutput, error)
	DeleteConfigurationSetAsync(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetInput) *PinpointsmsvoiceDeleteConfigurationSetFuture

	DeleteConfigurationSetEventDestination(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetEventDestinationInput) (*pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput, error)
	DeleteConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetEventDestinationInput) *PinpointsmsvoiceDeleteConfigurationSetEventDestinationFuture

	GetConfigurationSetEventDestinations(ctx workflow.Context, input *pinpointsmsvoice.GetConfigurationSetEventDestinationsInput) (*pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput, error)
	GetConfigurationSetEventDestinationsAsync(ctx workflow.Context, input *pinpointsmsvoice.GetConfigurationSetEventDestinationsInput) *PinpointsmsvoiceGetConfigurationSetEventDestinationsFuture

	ListConfigurationSets(ctx workflow.Context, input *pinpointsmsvoice.ListConfigurationSetsInput) (*pinpointsmsvoice.ListConfigurationSetsOutput, error)
	ListConfigurationSetsAsync(ctx workflow.Context, input *pinpointsmsvoice.ListConfigurationSetsInput) *PinpointsmsvoiceListConfigurationSetsFuture

	SendVoiceMessage(ctx workflow.Context, input *pinpointsmsvoice.SendVoiceMessageInput) (*pinpointsmsvoice.SendVoiceMessageOutput, error)
	SendVoiceMessageAsync(ctx workflow.Context, input *pinpointsmsvoice.SendVoiceMessageInput) *PinpointsmsvoiceSendVoiceMessageFuture

	UpdateConfigurationSetEventDestination(ctx workflow.Context, input *pinpointsmsvoice.UpdateConfigurationSetEventDestinationInput) (*pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput, error)
	UpdateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointsmsvoice.UpdateConfigurationSetEventDestinationInput) *PinpointsmsvoiceUpdateConfigurationSetEventDestinationFuture
}

type PinpointSMSVoiceStub struct{}

func NewPinpointSMSVoiceStub() PinpointSMSVoiceClient {
	return &PinpointSMSVoiceStub{}
}

type PinpointsmsvoiceCreateConfigurationSetFuture struct {
	Future workflow.Future
}

func (r *PinpointsmsvoiceCreateConfigurationSetFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.CreateConfigurationSetOutput, error) {
	var output pinpointsmsvoice.CreateConfigurationSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PinpointsmsvoiceCreateConfigurationSetEventDestinationFuture struct {
	Future workflow.Future
}

func (r *PinpointsmsvoiceCreateConfigurationSetEventDestinationFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PinpointsmsvoiceDeleteConfigurationSetFuture struct {
	Future workflow.Future
}

func (r *PinpointsmsvoiceDeleteConfigurationSetFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.DeleteConfigurationSetOutput, error) {
	var output pinpointsmsvoice.DeleteConfigurationSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PinpointsmsvoiceDeleteConfigurationSetEventDestinationFuture struct {
	Future workflow.Future
}

func (r *PinpointsmsvoiceDeleteConfigurationSetEventDestinationFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PinpointsmsvoiceGetConfigurationSetEventDestinationsFuture struct {
	Future workflow.Future
}

func (r *PinpointsmsvoiceGetConfigurationSetEventDestinationsFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput, error) {
	var output pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PinpointsmsvoiceListConfigurationSetsFuture struct {
	Future workflow.Future
}

func (r *PinpointsmsvoiceListConfigurationSetsFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.ListConfigurationSetsOutput, error) {
	var output pinpointsmsvoice.ListConfigurationSetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PinpointsmsvoiceSendVoiceMessageFuture struct {
	Future workflow.Future
}

func (r *PinpointsmsvoiceSendVoiceMessageFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.SendVoiceMessageOutput, error) {
	var output pinpointsmsvoice.SendVoiceMessageOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PinpointsmsvoiceUpdateConfigurationSetEventDestinationFuture struct {
	Future workflow.Future
}

func (r *PinpointsmsvoiceUpdateConfigurationSetEventDestinationFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *PinpointSMSVoiceStub) CreateConfigurationSet(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetInput) (*pinpointsmsvoice.CreateConfigurationSetOutput, error) {
	var output pinpointsmsvoice.CreateConfigurationSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.CreateConfigurationSet", input).Get(ctx, &output)
	return &output, err
}

func (a *PinpointSMSVoiceStub) CreateConfigurationSetAsync(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetInput) *PinpointsmsvoiceCreateConfigurationSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.CreateConfigurationSet", input)
	return &PinpointsmsvoiceCreateConfigurationSetFuture{Future: future}
}

func (a *PinpointSMSVoiceStub) CreateConfigurationSetEventDestination(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetEventDestinationInput) (*pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput
	err := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.CreateConfigurationSetEventDestination", input).Get(ctx, &output)
	return &output, err
}

func (a *PinpointSMSVoiceStub) CreateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetEventDestinationInput) *PinpointsmsvoiceCreateConfigurationSetEventDestinationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.CreateConfigurationSetEventDestination", input)
	return &PinpointsmsvoiceCreateConfigurationSetEventDestinationFuture{Future: future}
}

func (a *PinpointSMSVoiceStub) DeleteConfigurationSet(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetInput) (*pinpointsmsvoice.DeleteConfigurationSetOutput, error) {
	var output pinpointsmsvoice.DeleteConfigurationSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.DeleteConfigurationSet", input).Get(ctx, &output)
	return &output, err
}

func (a *PinpointSMSVoiceStub) DeleteConfigurationSetAsync(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetInput) *PinpointsmsvoiceDeleteConfigurationSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.DeleteConfigurationSet", input)
	return &PinpointsmsvoiceDeleteConfigurationSetFuture{Future: future}
}

func (a *PinpointSMSVoiceStub) DeleteConfigurationSetEventDestination(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetEventDestinationInput) (*pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput
	err := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.DeleteConfigurationSetEventDestination", input).Get(ctx, &output)
	return &output, err
}

func (a *PinpointSMSVoiceStub) DeleteConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetEventDestinationInput) *PinpointsmsvoiceDeleteConfigurationSetEventDestinationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.DeleteConfigurationSetEventDestination", input)
	return &PinpointsmsvoiceDeleteConfigurationSetEventDestinationFuture{Future: future}
}

func (a *PinpointSMSVoiceStub) GetConfigurationSetEventDestinations(ctx workflow.Context, input *pinpointsmsvoice.GetConfigurationSetEventDestinationsInput) (*pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput, error) {
	var output pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.GetConfigurationSetEventDestinations", input).Get(ctx, &output)
	return &output, err
}

func (a *PinpointSMSVoiceStub) GetConfigurationSetEventDestinationsAsync(ctx workflow.Context, input *pinpointsmsvoice.GetConfigurationSetEventDestinationsInput) *PinpointsmsvoiceGetConfigurationSetEventDestinationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.GetConfigurationSetEventDestinations", input)
	return &PinpointsmsvoiceGetConfigurationSetEventDestinationsFuture{Future: future}
}

func (a *PinpointSMSVoiceStub) ListConfigurationSets(ctx workflow.Context, input *pinpointsmsvoice.ListConfigurationSetsInput) (*pinpointsmsvoice.ListConfigurationSetsOutput, error) {
	var output pinpointsmsvoice.ListConfigurationSetsOutput
	err := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.ListConfigurationSets", input).Get(ctx, &output)
	return &output, err
}

func (a *PinpointSMSVoiceStub) ListConfigurationSetsAsync(ctx workflow.Context, input *pinpointsmsvoice.ListConfigurationSetsInput) *PinpointsmsvoiceListConfigurationSetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.ListConfigurationSets", input)
	return &PinpointsmsvoiceListConfigurationSetsFuture{Future: future}
}

func (a *PinpointSMSVoiceStub) SendVoiceMessage(ctx workflow.Context, input *pinpointsmsvoice.SendVoiceMessageInput) (*pinpointsmsvoice.SendVoiceMessageOutput, error) {
	var output pinpointsmsvoice.SendVoiceMessageOutput
	err := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.SendVoiceMessage", input).Get(ctx, &output)
	return &output, err
}

func (a *PinpointSMSVoiceStub) SendVoiceMessageAsync(ctx workflow.Context, input *pinpointsmsvoice.SendVoiceMessageInput) *PinpointsmsvoiceSendVoiceMessageFuture {
	future := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.SendVoiceMessage", input)
	return &PinpointsmsvoiceSendVoiceMessageFuture{Future: future}
}

func (a *PinpointSMSVoiceStub) UpdateConfigurationSetEventDestination(ctx workflow.Context, input *pinpointsmsvoice.UpdateConfigurationSetEventDestinationInput) (*pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput
	err := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.UpdateConfigurationSetEventDestination", input).Get(ctx, &output)
	return &output, err
}

func (a *PinpointSMSVoiceStub) UpdateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointsmsvoice.UpdateConfigurationSetEventDestinationInput) *PinpointsmsvoiceUpdateConfigurationSetEventDestinationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.UpdateConfigurationSetEventDestination", input)
	return &PinpointsmsvoiceUpdateConfigurationSetEventDestinationFuture{Future: future}
}
