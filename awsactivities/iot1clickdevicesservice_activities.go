package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iot1clickdevicesservice"
	"github.com/aws/aws-sdk-go/service/iot1clickdevicesservice/iot1clickdevicesserviceiface"
)

type IoT1ClickDevicesServiceActivities struct {
	client iot1clickdevicesserviceiface.IoT1ClickDevicesServiceAPI
}

func NewIoT1ClickDevicesServiceActivities(session *session.Session, config ...*aws.Config) *IoT1ClickDevicesServiceActivities {
	client := iot1clickdevicesservice.New(session, config...)
	return &IoT1ClickDevicesServiceActivities{client: client}
}

func (a *IoT1ClickDevicesServiceActivities) ClaimDevicesByClaimCode(input *iot1clickdevicesservice.ClaimDevicesByClaimCodeInput) (*iot1clickdevicesservice.ClaimDevicesByClaimCodeOutput, error) {
	return a.client.ClaimDevicesByClaimCode(input)
}

func (a *IoT1ClickDevicesServiceActivities) DescribeDevice(input *iot1clickdevicesservice.DescribeDeviceInput) (*iot1clickdevicesservice.DescribeDeviceOutput, error) {
	return a.client.DescribeDevice(input)
}

func (a *IoT1ClickDevicesServiceActivities) FinalizeDeviceClaim(input *iot1clickdevicesservice.FinalizeDeviceClaimInput) (*iot1clickdevicesservice.FinalizeDeviceClaimOutput, error) {
	return a.client.FinalizeDeviceClaim(input)
}

func (a *IoT1ClickDevicesServiceActivities) GetDeviceMethods(input *iot1clickdevicesservice.GetDeviceMethodsInput) (*iot1clickdevicesservice.GetDeviceMethodsOutput, error) {
	return a.client.GetDeviceMethods(input)
}

func (a *IoT1ClickDevicesServiceActivities) InitiateDeviceClaim(input *iot1clickdevicesservice.InitiateDeviceClaimInput) (*iot1clickdevicesservice.InitiateDeviceClaimOutput, error) {
	return a.client.InitiateDeviceClaim(input)
}

func (a *IoT1ClickDevicesServiceActivities) InvokeDeviceMethod(input *iot1clickdevicesservice.InvokeDeviceMethodInput) (*iot1clickdevicesservice.InvokeDeviceMethodOutput, error) {
	return a.client.InvokeDeviceMethod(input)
}

func (a *IoT1ClickDevicesServiceActivities) ListDeviceEvents(input *iot1clickdevicesservice.ListDeviceEventsInput) (*iot1clickdevicesservice.ListDeviceEventsOutput, error) {
	return a.client.ListDeviceEvents(input)
}

func (a *IoT1ClickDevicesServiceActivities) ListDevices(input *iot1clickdevicesservice.ListDevicesInput) (*iot1clickdevicesservice.ListDevicesOutput, error) {
	return a.client.ListDevices(input)
}

func (a *IoT1ClickDevicesServiceActivities) ListTagsForResource(input *iot1clickdevicesservice.ListTagsForResourceInput) (*iot1clickdevicesservice.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResource(input)
}

func (a *IoT1ClickDevicesServiceActivities) TagResource(input *iot1clickdevicesservice.TagResourceInput) (*iot1clickdevicesservice.TagResourceOutput, error) {
	return a.client.TagResource(input)
}

func (a *IoT1ClickDevicesServiceActivities) UnclaimDevice(input *iot1clickdevicesservice.UnclaimDeviceInput) (*iot1clickdevicesservice.UnclaimDeviceOutput, error) {
	return a.client.UnclaimDevice(input)
}

func (a *IoT1ClickDevicesServiceActivities) UntagResource(input *iot1clickdevicesservice.UntagResourceInput) (*iot1clickdevicesservice.UntagResourceOutput, error) {
	return a.client.UntagResource(input)
}

func (a *IoT1ClickDevicesServiceActivities) UpdateDeviceState(input *iot1clickdevicesservice.UpdateDeviceStateInput) (*iot1clickdevicesservice.UpdateDeviceStateOutput, error) {
	return a.client.UpdateDeviceState(input)
}
