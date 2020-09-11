package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/pinpointsmsvoice"
	"github.com/aws/aws-sdk-go/service/pinpointsmsvoice/pinpointsmsvoiceiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type PinpointSMSVoiceActivities struct {
	client pinpointsmsvoiceiface.PinpointSMSVoiceAPI
}

func NewPinpointSMSVoiceActivities(session *session.Session, config ...*aws.Config) *PinpointSMSVoiceActivities {
	client := pinpointsmsvoice.New(session, config...)
	return &PinpointSMSVoiceActivities{client: client}
}

func (a *PinpointSMSVoiceActivities) CreateConfigurationSet(ctx context.Context, input *pinpointsmsvoice.CreateConfigurationSetInput) (*pinpointsmsvoice.CreateConfigurationSetOutput, error) {
	return a.client.CreateConfigurationSetWithContext(ctx, input)
}

func (a *PinpointSMSVoiceActivities) CreateConfigurationSetEventDestination(ctx context.Context, input *pinpointsmsvoice.CreateConfigurationSetEventDestinationInput) (*pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput, error) {
	return a.client.CreateConfigurationSetEventDestinationWithContext(ctx, input)
}

func (a *PinpointSMSVoiceActivities) DeleteConfigurationSet(ctx context.Context, input *pinpointsmsvoice.DeleteConfigurationSetInput) (*pinpointsmsvoice.DeleteConfigurationSetOutput, error) {
	return a.client.DeleteConfigurationSetWithContext(ctx, input)
}

func (a *PinpointSMSVoiceActivities) DeleteConfigurationSetEventDestination(ctx context.Context, input *pinpointsmsvoice.DeleteConfigurationSetEventDestinationInput) (*pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput, error) {
	return a.client.DeleteConfigurationSetEventDestinationWithContext(ctx, input)
}

func (a *PinpointSMSVoiceActivities) GetConfigurationSetEventDestinations(ctx context.Context, input *pinpointsmsvoice.GetConfigurationSetEventDestinationsInput) (*pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput, error) {
	return a.client.GetConfigurationSetEventDestinationsWithContext(ctx, input)
}

func (a *PinpointSMSVoiceActivities) ListConfigurationSets(ctx context.Context, input *pinpointsmsvoice.ListConfigurationSetsInput) (*pinpointsmsvoice.ListConfigurationSetsOutput, error) {
	return a.client.ListConfigurationSetsWithContext(ctx, input)
}

func (a *PinpointSMSVoiceActivities) SendVoiceMessage(ctx context.Context, input *pinpointsmsvoice.SendVoiceMessageInput) (*pinpointsmsvoice.SendVoiceMessageOutput, error) {
	return a.client.SendVoiceMessageWithContext(ctx, input)
}

func (a *PinpointSMSVoiceActivities) UpdateConfigurationSetEventDestination(ctx context.Context, input *pinpointsmsvoice.UpdateConfigurationSetEventDestinationInput) (*pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput, error) {
	return a.client.UpdateConfigurationSetEventDestinationWithContext(ctx, input)
}
