package awsclients

import (
	"github.com/aws/aws-sdk-go/service/pinpointsmsvoice"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type PinpointSMSVoiceClient interface {
	CreateConfigurationSet(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetInput) (*pinpointsmsvoice.CreateConfigurationSetOutput, error)
	CreateConfigurationSetAsync(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetInput) *PinpointsmsvoiceCreateConfigurationSetResult

	CreateConfigurationSetEventDestination(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetEventDestinationInput) (*pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput, error)
	CreateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetEventDestinationInput) *PinpointsmsvoiceCreateConfigurationSetEventDestinationResult

	DeleteConfigurationSet(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetInput) (*pinpointsmsvoice.DeleteConfigurationSetOutput, error)
	DeleteConfigurationSetAsync(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetInput) *PinpointsmsvoiceDeleteConfigurationSetResult

	DeleteConfigurationSetEventDestination(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetEventDestinationInput) (*pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput, error)
	DeleteConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetEventDestinationInput) *PinpointsmsvoiceDeleteConfigurationSetEventDestinationResult

	GetConfigurationSetEventDestinations(ctx workflow.Context, input *pinpointsmsvoice.GetConfigurationSetEventDestinationsInput) (*pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput, error)
	GetConfigurationSetEventDestinationsAsync(ctx workflow.Context, input *pinpointsmsvoice.GetConfigurationSetEventDestinationsInput) *PinpointsmsvoiceGetConfigurationSetEventDestinationsResult

	ListConfigurationSets(ctx workflow.Context, input *pinpointsmsvoice.ListConfigurationSetsInput) (*pinpointsmsvoice.ListConfigurationSetsOutput, error)
	ListConfigurationSetsAsync(ctx workflow.Context, input *pinpointsmsvoice.ListConfigurationSetsInput) *PinpointsmsvoiceListConfigurationSetsResult

	SendVoiceMessage(ctx workflow.Context, input *pinpointsmsvoice.SendVoiceMessageInput) (*pinpointsmsvoice.SendVoiceMessageOutput, error)
	SendVoiceMessageAsync(ctx workflow.Context, input *pinpointsmsvoice.SendVoiceMessageInput) *PinpointsmsvoiceSendVoiceMessageResult

	UpdateConfigurationSetEventDestination(ctx workflow.Context, input *pinpointsmsvoice.UpdateConfigurationSetEventDestinationInput) (*pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput, error)
	UpdateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointsmsvoice.UpdateConfigurationSetEventDestinationInput) *PinpointsmsvoiceUpdateConfigurationSetEventDestinationResult
}

type PinpointsmsvoiceCreateConfigurationSetResult struct {
	Result workflow.Future
}

func (r *PinpointsmsvoiceCreateConfigurationSetResult) Get(ctx workflow.Context) (*pinpointsmsvoice.CreateConfigurationSetOutput, error) {
	var output pinpointsmsvoice.CreateConfigurationSetOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type PinpointsmsvoiceCreateConfigurationSetEventDestinationResult struct {
	Result workflow.Future
}

func (r *PinpointsmsvoiceCreateConfigurationSetEventDestinationResult) Get(ctx workflow.Context) (*pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type PinpointsmsvoiceDeleteConfigurationSetResult struct {
	Result workflow.Future
}

func (r *PinpointsmsvoiceDeleteConfigurationSetResult) Get(ctx workflow.Context) (*pinpointsmsvoice.DeleteConfigurationSetOutput, error) {
	var output pinpointsmsvoice.DeleteConfigurationSetOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type PinpointsmsvoiceDeleteConfigurationSetEventDestinationResult struct {
	Result workflow.Future
}

func (r *PinpointsmsvoiceDeleteConfigurationSetEventDestinationResult) Get(ctx workflow.Context) (*pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type PinpointsmsvoiceGetConfigurationSetEventDestinationsResult struct {
	Result workflow.Future
}

func (r *PinpointsmsvoiceGetConfigurationSetEventDestinationsResult) Get(ctx workflow.Context) (*pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput, error) {
	var output pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type PinpointsmsvoiceListConfigurationSetsResult struct {
	Result workflow.Future
}

func (r *PinpointsmsvoiceListConfigurationSetsResult) Get(ctx workflow.Context) (*pinpointsmsvoice.ListConfigurationSetsOutput, error) {
	var output pinpointsmsvoice.ListConfigurationSetsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type PinpointsmsvoiceSendVoiceMessageResult struct {
	Result workflow.Future
}

func (r *PinpointsmsvoiceSendVoiceMessageResult) Get(ctx workflow.Context) (*pinpointsmsvoice.SendVoiceMessageOutput, error) {
	var output pinpointsmsvoice.SendVoiceMessageOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type PinpointsmsvoiceUpdateConfigurationSetEventDestinationResult struct {
	Result workflow.Future
}

func (r *PinpointsmsvoiceUpdateConfigurationSetEventDestinationResult) Get(ctx workflow.Context) (*pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type PinpointSMSVoiceStub struct {
	activities awsactivities.PinpointSMSVoiceActivities
}

func NewPinpointSMSVoiceStub() PinpointSMSVoiceClient {
	return &PinpointSMSVoiceStub{}
}

func (a *PinpointSMSVoiceStub) CreateConfigurationSet(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetInput) (*pinpointsmsvoice.CreateConfigurationSetOutput, error) {
	var output pinpointsmsvoice.CreateConfigurationSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateConfigurationSet, input).Get(ctx, &output)
	return &output, err
}

func (a *PinpointSMSVoiceStub) CreateConfigurationSetAsync(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetInput) *PinpointsmsvoiceCreateConfigurationSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateConfigurationSet, input)
	return &PinpointsmsvoiceCreateConfigurationSetResult{Result: future}
}

func (a *PinpointSMSVoiceStub) CreateConfigurationSetEventDestination(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetEventDestinationInput) (*pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateConfigurationSetEventDestination, input).Get(ctx, &output)
	return &output, err
}

func (a *PinpointSMSVoiceStub) CreateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetEventDestinationInput) *PinpointsmsvoiceCreateConfigurationSetEventDestinationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateConfigurationSetEventDestination, input)
	return &PinpointsmsvoiceCreateConfigurationSetEventDestinationResult{Result: future}
}

func (a *PinpointSMSVoiceStub) DeleteConfigurationSet(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetInput) (*pinpointsmsvoice.DeleteConfigurationSetOutput, error) {
	var output pinpointsmsvoice.DeleteConfigurationSetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteConfigurationSet, input).Get(ctx, &output)
	return &output, err
}

func (a *PinpointSMSVoiceStub) DeleteConfigurationSetAsync(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetInput) *PinpointsmsvoiceDeleteConfigurationSetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteConfigurationSet, input)
	return &PinpointsmsvoiceDeleteConfigurationSetResult{Result: future}
}

func (a *PinpointSMSVoiceStub) DeleteConfigurationSetEventDestination(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetEventDestinationInput) (*pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteConfigurationSetEventDestination, input).Get(ctx, &output)
	return &output, err
}

func (a *PinpointSMSVoiceStub) DeleteConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetEventDestinationInput) *PinpointsmsvoiceDeleteConfigurationSetEventDestinationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteConfigurationSetEventDestination, input)
	return &PinpointsmsvoiceDeleteConfigurationSetEventDestinationResult{Result: future}
}

func (a *PinpointSMSVoiceStub) GetConfigurationSetEventDestinations(ctx workflow.Context, input *pinpointsmsvoice.GetConfigurationSetEventDestinationsInput) (*pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput, error) {
	var output pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetConfigurationSetEventDestinations, input).Get(ctx, &output)
	return &output, err
}

func (a *PinpointSMSVoiceStub) GetConfigurationSetEventDestinationsAsync(ctx workflow.Context, input *pinpointsmsvoice.GetConfigurationSetEventDestinationsInput) *PinpointsmsvoiceGetConfigurationSetEventDestinationsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetConfigurationSetEventDestinations, input)
	return &PinpointsmsvoiceGetConfigurationSetEventDestinationsResult{Result: future}
}

func (a *PinpointSMSVoiceStub) ListConfigurationSets(ctx workflow.Context, input *pinpointsmsvoice.ListConfigurationSetsInput) (*pinpointsmsvoice.ListConfigurationSetsOutput, error) {
	var output pinpointsmsvoice.ListConfigurationSetsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListConfigurationSets, input).Get(ctx, &output)
	return &output, err
}

func (a *PinpointSMSVoiceStub) ListConfigurationSetsAsync(ctx workflow.Context, input *pinpointsmsvoice.ListConfigurationSetsInput) *PinpointsmsvoiceListConfigurationSetsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListConfigurationSets, input)
	return &PinpointsmsvoiceListConfigurationSetsResult{Result: future}
}

func (a *PinpointSMSVoiceStub) SendVoiceMessage(ctx workflow.Context, input *pinpointsmsvoice.SendVoiceMessageInput) (*pinpointsmsvoice.SendVoiceMessageOutput, error) {
	var output pinpointsmsvoice.SendVoiceMessageOutput
	err := workflow.ExecuteActivity(ctx, a.activities.SendVoiceMessage, input).Get(ctx, &output)
	return &output, err
}

func (a *PinpointSMSVoiceStub) SendVoiceMessageAsync(ctx workflow.Context, input *pinpointsmsvoice.SendVoiceMessageInput) *PinpointsmsvoiceSendVoiceMessageResult {
	future := workflow.ExecuteActivity(ctx, a.activities.SendVoiceMessage, input)
	return &PinpointsmsvoiceSendVoiceMessageResult{Result: future}
}

func (a *PinpointSMSVoiceStub) UpdateConfigurationSetEventDestination(ctx workflow.Context, input *pinpointsmsvoice.UpdateConfigurationSetEventDestinationInput) (*pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateConfigurationSetEventDestination, input).Get(ctx, &output)
	return &output, err
}

func (a *PinpointSMSVoiceStub) UpdateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointsmsvoice.UpdateConfigurationSetEventDestinationInput) *PinpointsmsvoiceUpdateConfigurationSetEventDestinationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateConfigurationSetEventDestination, input)
	return &PinpointsmsvoiceUpdateConfigurationSetEventDestinationResult{Result: future}
}
