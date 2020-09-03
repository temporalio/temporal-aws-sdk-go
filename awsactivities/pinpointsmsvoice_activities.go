package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/pinpointsmsvoice"
	"github.com/aws/aws-sdk-go/service/pinpointsmsvoice/pinpointsmsvoiceiface"
)

type PinpointSMSVoiceActivities struct {
	client pinpointsmsvoiceiface.PinpointSMSVoiceAPI
}

func NewPinpointSMSVoiceActivities(client pinpointsmsvoiceiface.PinpointSMSVoiceAPI) *PinpointSMSVoiceActivities {
    return &PinpointSMSVoiceActivities{client: client}
}


func (a *PinpointSMSVoiceActivities) CreateConfigurationSet(input *pinpointsmsvoice.CreateConfigurationSetInput) (*pinpointsmsvoice.CreateConfigurationSetOutput, error) {
    return a.client.CreateConfigurationSet(input)
}



func (a *PinpointSMSVoiceActivities) CreateConfigurationSetEventDestination(input *pinpointsmsvoice.CreateConfigurationSetEventDestinationInput) (*pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput, error) {
    return a.client.CreateConfigurationSetEventDestination(input)
}



func (a *PinpointSMSVoiceActivities) DeleteConfigurationSet(input *pinpointsmsvoice.DeleteConfigurationSetInput) (*pinpointsmsvoice.DeleteConfigurationSetOutput, error) {
    return a.client.DeleteConfigurationSet(input)
}



func (a *PinpointSMSVoiceActivities) DeleteConfigurationSetEventDestination(input *pinpointsmsvoice.DeleteConfigurationSetEventDestinationInput) (*pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput, error) {
    return a.client.DeleteConfigurationSetEventDestination(input)
}



func (a *PinpointSMSVoiceActivities) GetConfigurationSetEventDestinations(input *pinpointsmsvoice.GetConfigurationSetEventDestinationsInput) (*pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput, error) {
    return a.client.GetConfigurationSetEventDestinations(input)
}



func (a *PinpointSMSVoiceActivities) ListConfigurationSets(input *pinpointsmsvoice.ListConfigurationSetsInput) (*pinpointsmsvoice.ListConfigurationSetsOutput, error) {
    return a.client.ListConfigurationSets(input)
}



func (a *PinpointSMSVoiceActivities) SendVoiceMessage(input *pinpointsmsvoice.SendVoiceMessageInput) (*pinpointsmsvoice.SendVoiceMessageOutput, error) {
    return a.client.SendVoiceMessage(input)
}



func (a *PinpointSMSVoiceActivities) UpdateConfigurationSetEventDestination(input *pinpointsmsvoice.UpdateConfigurationSetEventDestinationInput) (*pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput, error) {
    return a.client.UpdateConfigurationSetEventDestination(input)
}

