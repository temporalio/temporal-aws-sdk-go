
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/medialive"
	"github.com/aws/aws-sdk-go/service/medialive/medialiveiface"
)

type MediaLiveActivities struct {
    client medialiveiface.MediaLiveAPI
}

func NewMediaLiveActivities(session *session.Session, config ...*aws.Config) *MediaLiveActivities {
    client := medialive.New(session, config...)
    return &MediaLiveActivities{client: client}
}

func (a *MediaLiveActivities) BatchUpdateSchedule(input *medialive.BatchUpdateScheduleInput) (*medialive.BatchUpdateScheduleOutput, error) {
    return a.client.BatchUpdateSchedule(input)
}

func (a *MediaLiveActivities) CreateChannel(input *medialive.CreateChannelInput) (*medialive.CreateChannelOutput, error) {
    return a.client.CreateChannel(input)
}

func (a *MediaLiveActivities) CreateInput(input *medialive.CreateInputInput) (*medialive.CreateInputOutput, error) {
    return a.client.CreateInput(input)
}

func (a *MediaLiveActivities) CreateInputSecurityGroup(input *medialive.CreateInputSecurityGroupInput) (*medialive.CreateInputSecurityGroupOutput, error) {
    return a.client.CreateInputSecurityGroup(input)
}

func (a *MediaLiveActivities) CreateMultiplex(input *medialive.CreateMultiplexInput) (*medialive.CreateMultiplexOutput, error) {
    return a.client.CreateMultiplex(input)
}

func (a *MediaLiveActivities) CreateMultiplexProgram(input *medialive.CreateMultiplexProgramInput) (*medialive.CreateMultiplexProgramOutput, error) {
    return a.client.CreateMultiplexProgram(input)
}

func (a *MediaLiveActivities) CreateTags(input *medialive.CreateTagsInput) (*medialive.CreateTagsOutput, error) {
    return a.client.CreateTags(input)
}

func (a *MediaLiveActivities) DeleteChannel(input *medialive.DeleteChannelInput) (*medialive.DeleteChannelOutput, error) {
    return a.client.DeleteChannel(input)
}

func (a *MediaLiveActivities) DeleteInput(input *medialive.DeleteInputInput) (*medialive.DeleteInputOutput, error) {
    return a.client.DeleteInput(input)
}

func (a *MediaLiveActivities) DeleteInputSecurityGroup(input *medialive.DeleteInputSecurityGroupInput) (*medialive.DeleteInputSecurityGroupOutput, error) {
    return a.client.DeleteInputSecurityGroup(input)
}

func (a *MediaLiveActivities) DeleteMultiplex(input *medialive.DeleteMultiplexInput) (*medialive.DeleteMultiplexOutput, error) {
    return a.client.DeleteMultiplex(input)
}

func (a *MediaLiveActivities) DeleteMultiplexProgram(input *medialive.DeleteMultiplexProgramInput) (*medialive.DeleteMultiplexProgramOutput, error) {
    return a.client.DeleteMultiplexProgram(input)
}

func (a *MediaLiveActivities) DeleteReservation(input *medialive.DeleteReservationInput) (*medialive.DeleteReservationOutput, error) {
    return a.client.DeleteReservation(input)
}

func (a *MediaLiveActivities) DeleteSchedule(input *medialive.DeleteScheduleInput) (*medialive.DeleteScheduleOutput, error) {
    return a.client.DeleteSchedule(input)
}

func (a *MediaLiveActivities) DeleteTags(input *medialive.DeleteTagsInput) (*medialive.DeleteTagsOutput, error) {
    return a.client.DeleteTags(input)
}

func (a *MediaLiveActivities) DescribeChannel(input *medialive.DescribeChannelInput) (*medialive.DescribeChannelOutput, error) {
    return a.client.DescribeChannel(input)
}

func (a *MediaLiveActivities) DescribeInput(input *medialive.DescribeInputInput) (*medialive.DescribeInputOutput, error) {
    return a.client.DescribeInput(input)
}

func (a *MediaLiveActivities) DescribeInputDevice(input *medialive.DescribeInputDeviceInput) (*medialive.DescribeInputDeviceOutput, error) {
    return a.client.DescribeInputDevice(input)
}

func (a *MediaLiveActivities) DescribeInputDeviceThumbnail(input *medialive.DescribeInputDeviceThumbnailInput) (*medialive.DescribeInputDeviceThumbnailOutput, error) {
    return a.client.DescribeInputDeviceThumbnail(input)
}

func (a *MediaLiveActivities) DescribeInputSecurityGroup(input *medialive.DescribeInputSecurityGroupInput) (*medialive.DescribeInputSecurityGroupOutput, error) {
    return a.client.DescribeInputSecurityGroup(input)
}

func (a *MediaLiveActivities) DescribeMultiplex(input *medialive.DescribeMultiplexInput) (*medialive.DescribeMultiplexOutput, error) {
    return a.client.DescribeMultiplex(input)
}

func (a *MediaLiveActivities) DescribeMultiplexProgram(input *medialive.DescribeMultiplexProgramInput) (*medialive.DescribeMultiplexProgramOutput, error) {
    return a.client.DescribeMultiplexProgram(input)
}

func (a *MediaLiveActivities) DescribeOffering(input *medialive.DescribeOfferingInput) (*medialive.DescribeOfferingOutput, error) {
    return a.client.DescribeOffering(input)
}

func (a *MediaLiveActivities) DescribeReservation(input *medialive.DescribeReservationInput) (*medialive.DescribeReservationOutput, error) {
    return a.client.DescribeReservation(input)
}

func (a *MediaLiveActivities) DescribeSchedule(input *medialive.DescribeScheduleInput) (*medialive.DescribeScheduleOutput, error) {
    return a.client.DescribeSchedule(input)
}

func (a *MediaLiveActivities) ListChannels(input *medialive.ListChannelsInput) (*medialive.ListChannelsOutput, error) {
    return a.client.ListChannels(input)
}

func (a *MediaLiveActivities) ListInputDevices(input *medialive.ListInputDevicesInput) (*medialive.ListInputDevicesOutput, error) {
    return a.client.ListInputDevices(input)
}

func (a *MediaLiveActivities) ListInputSecurityGroups(input *medialive.ListInputSecurityGroupsInput) (*medialive.ListInputSecurityGroupsOutput, error) {
    return a.client.ListInputSecurityGroups(input)
}

func (a *MediaLiveActivities) ListInputs(input *medialive.ListInputsInput) (*medialive.ListInputsOutput, error) {
    return a.client.ListInputs(input)
}

func (a *MediaLiveActivities) ListMultiplexPrograms(input *medialive.ListMultiplexProgramsInput) (*medialive.ListMultiplexProgramsOutput, error) {
    return a.client.ListMultiplexPrograms(input)
}

func (a *MediaLiveActivities) ListMultiplexes(input *medialive.ListMultiplexesInput) (*medialive.ListMultiplexesOutput, error) {
    return a.client.ListMultiplexes(input)
}

func (a *MediaLiveActivities) ListOfferings(input *medialive.ListOfferingsInput) (*medialive.ListOfferingsOutput, error) {
    return a.client.ListOfferings(input)
}

func (a *MediaLiveActivities) ListReservations(input *medialive.ListReservationsInput) (*medialive.ListReservationsOutput, error) {
    return a.client.ListReservations(input)
}

func (a *MediaLiveActivities) ListTagsForResource(input *medialive.ListTagsForResourceInput) (*medialive.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *MediaLiveActivities) PurchaseOffering(input *medialive.PurchaseOfferingInput) (*medialive.PurchaseOfferingOutput, error) {
    return a.client.PurchaseOffering(input)
}

func (a *MediaLiveActivities) StartChannel(input *medialive.StartChannelInput) (*medialive.StartChannelOutput, error) {
    return a.client.StartChannel(input)
}

func (a *MediaLiveActivities) StartMultiplex(input *medialive.StartMultiplexInput) (*medialive.StartMultiplexOutput, error) {
    return a.client.StartMultiplex(input)
}

func (a *MediaLiveActivities) StopChannel(input *medialive.StopChannelInput) (*medialive.StopChannelOutput, error) {
    return a.client.StopChannel(input)
}

func (a *MediaLiveActivities) StopMultiplex(input *medialive.StopMultiplexInput) (*medialive.StopMultiplexOutput, error) {
    return a.client.StopMultiplex(input)
}

func (a *MediaLiveActivities) UpdateChannel(input *medialive.UpdateChannelInput) (*medialive.UpdateChannelOutput, error) {
    return a.client.UpdateChannel(input)
}

func (a *MediaLiveActivities) UpdateChannelClass(input *medialive.UpdateChannelClassInput) (*medialive.UpdateChannelClassOutput, error) {
    return a.client.UpdateChannelClass(input)
}

func (a *MediaLiveActivities) UpdateInput(input *medialive.UpdateInputInput) (*medialive.UpdateInputOutput, error) {
    return a.client.UpdateInput(input)
}

func (a *MediaLiveActivities) UpdateInputDevice(input *medialive.UpdateInputDeviceInput) (*medialive.UpdateInputDeviceOutput, error) {
    return a.client.UpdateInputDevice(input)
}

func (a *MediaLiveActivities) UpdateInputSecurityGroup(input *medialive.UpdateInputSecurityGroupInput) (*medialive.UpdateInputSecurityGroupOutput, error) {
    return a.client.UpdateInputSecurityGroup(input)
}

func (a *MediaLiveActivities) UpdateMultiplex(input *medialive.UpdateMultiplexInput) (*medialive.UpdateMultiplexOutput, error) {
    return a.client.UpdateMultiplex(input)
}

func (a *MediaLiveActivities) UpdateMultiplexProgram(input *medialive.UpdateMultiplexProgramInput) (*medialive.UpdateMultiplexProgramOutput, error) {
    return a.client.UpdateMultiplexProgram(input)
}

func (a *MediaLiveActivities) UpdateReservation(input *medialive.UpdateReservationInput) (*medialive.UpdateReservationOutput, error) {
    return a.client.UpdateReservation(input)
}

func (a *MediaLiveActivities) WaitUntilChannelCreated(input *medialive.DescribeChannelInput) error {
    return a.client.WaitUntilChannelCreated(input)
}

func (a *MediaLiveActivities) WaitUntilChannelDeleted(input *medialive.DescribeChannelInput) error {
    return a.client.WaitUntilChannelDeleted(input)
}

func (a *MediaLiveActivities) WaitUntilChannelRunning(input *medialive.DescribeChannelInput) error {
    return a.client.WaitUntilChannelRunning(input)
}

func (a *MediaLiveActivities) WaitUntilChannelStopped(input *medialive.DescribeChannelInput) error {
    return a.client.WaitUntilChannelStopped(input)
}

func (a *MediaLiveActivities) WaitUntilInputAttached(input *medialive.DescribeInputInput) error {
    return a.client.WaitUntilInputAttached(input)
}

func (a *MediaLiveActivities) WaitUntilInputDeleted(input *medialive.DescribeInputInput) error {
    return a.client.WaitUntilInputDeleted(input)
}

func (a *MediaLiveActivities) WaitUntilInputDetached(input *medialive.DescribeInputInput) error {
    return a.client.WaitUntilInputDetached(input)
}

func (a *MediaLiveActivities) WaitUntilMultiplexCreated(input *medialive.DescribeMultiplexInput) error {
    return a.client.WaitUntilMultiplexCreated(input)
}

func (a *MediaLiveActivities) WaitUntilMultiplexDeleted(input *medialive.DescribeMultiplexInput) error {
    return a.client.WaitUntilMultiplexDeleted(input)
}

func (a *MediaLiveActivities) WaitUntilMultiplexRunning(input *medialive.DescribeMultiplexInput) error {
    return a.client.WaitUntilMultiplexRunning(input)
}

func (a *MediaLiveActivities) WaitUntilMultiplexStopped(input *medialive.DescribeMultiplexInput) error {
    return a.client.WaitUntilMultiplexStopped(input)
}
