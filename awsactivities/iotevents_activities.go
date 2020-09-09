
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotevents"
	"github.com/aws/aws-sdk-go/service/iotevents/ioteventsiface"
)

type IoTEventsActivities struct {
    client ioteventsiface.IoTEventsAPI
}

func NewIoTEventsActivities(session *session.Session, config ...*aws.Config) *IoTEventsActivities {
    client := iotevents.New(session, config...)
    return &IoTEventsActivities{client: client}
}

func (a *IoTEventsActivities) CreateDetectorModel(input *iotevents.CreateDetectorModelInput) (*iotevents.CreateDetectorModelOutput, error) {
    return a.client.CreateDetectorModel(input)
}

func (a *IoTEventsActivities) CreateInput(input *iotevents.CreateInputInput) (*iotevents.CreateInputOutput, error) {
    return a.client.CreateInput(input)
}

func (a *IoTEventsActivities) DeleteDetectorModel(input *iotevents.DeleteDetectorModelInput) (*iotevents.DeleteDetectorModelOutput, error) {
    return a.client.DeleteDetectorModel(input)
}

func (a *IoTEventsActivities) DeleteInput(input *iotevents.DeleteInputInput) (*iotevents.DeleteInputOutput, error) {
    return a.client.DeleteInput(input)
}

func (a *IoTEventsActivities) DescribeDetectorModel(input *iotevents.DescribeDetectorModelInput) (*iotevents.DescribeDetectorModelOutput, error) {
    return a.client.DescribeDetectorModel(input)
}

func (a *IoTEventsActivities) DescribeInput(input *iotevents.DescribeInputInput) (*iotevents.DescribeInputOutput, error) {
    return a.client.DescribeInput(input)
}

func (a *IoTEventsActivities) DescribeLoggingOptions(input *iotevents.DescribeLoggingOptionsInput) (*iotevents.DescribeLoggingOptionsOutput, error) {
    return a.client.DescribeLoggingOptions(input)
}

func (a *IoTEventsActivities) ListDetectorModelVersions(input *iotevents.ListDetectorModelVersionsInput) (*iotevents.ListDetectorModelVersionsOutput, error) {
    return a.client.ListDetectorModelVersions(input)
}

func (a *IoTEventsActivities) ListDetectorModels(input *iotevents.ListDetectorModelsInput) (*iotevents.ListDetectorModelsOutput, error) {
    return a.client.ListDetectorModels(input)
}

func (a *IoTEventsActivities) ListInputs(input *iotevents.ListInputsInput) (*iotevents.ListInputsOutput, error) {
    return a.client.ListInputs(input)
}

func (a *IoTEventsActivities) ListTagsForResource(input *iotevents.ListTagsForResourceInput) (*iotevents.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *IoTEventsActivities) PutLoggingOptions(input *iotevents.PutLoggingOptionsInput) (*iotevents.PutLoggingOptionsOutput, error) {
    return a.client.PutLoggingOptions(input)
}

func (a *IoTEventsActivities) TagResource(input *iotevents.TagResourceInput) (*iotevents.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *IoTEventsActivities) UntagResource(input *iotevents.UntagResourceInput) (*iotevents.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *IoTEventsActivities) UpdateDetectorModel(input *iotevents.UpdateDetectorModelInput) (*iotevents.UpdateDetectorModelOutput, error) {
    return a.client.UpdateDetectorModel(input)
}

func (a *IoTEventsActivities) UpdateInput(input *iotevents.UpdateInputInput) (*iotevents.UpdateInputOutput, error) {
    return a.client.UpdateInput(input)
}
